package engine

import (
	"fmt"
	"strings"

	"rza_migrate_db/internal/sqlparse"
)

const (
	functionSubtypeScalar           = "scalar"
	functionSubtypeInlineTable      = "inline_table"
	functionSubtypeMultiStatementTV = "multi_statement_table"
)

func classifyFunctionSubtype(sql string) (string, error) {
	scanned := strings.ToUpper(sqlparse.StripSQLStringsAndComments(sql))
	returnsPositions := findAllKeywordPositions(scanned, "RETURNS")
	if len(returnsPositions) != 1 {
		return "", fmt.Errorf("function must contain exactly one RETURNS clause")
	}
	pos := returnsPositions[0] + len("RETURNS")
	returnTail := strings.TrimSpace(scanned[pos:])
	token := firstSQLToken(returnTail)
	if token == "" {
		return "", fmt.Errorf("function RETURNS clause has no return type")
	}
	if token == "TABLE" {
		if len(findAllKeywordPositions(returnTail[len(token):], "RETURN")) == 0 {
			return "", fmt.Errorf("inline table-valued function must contain RETURN")
		}
		return functionSubtypeInlineTable, nil
	}
	if strings.HasPrefix(token, "@") {
		tableToken := firstSQLToken(strings.TrimSpace(returnTail[len(token):]))
		if tableToken != "TABLE" {
			return "", fmt.Errorf("function table return variable must be followed by TABLE")
		}
		if len(findAllKeywordPositions(scanned, "BEGIN")) == 0 || len(findAllKeywordPositions(scanned, "END")) == 0 {
			return "", fmt.Errorf("multi-statement table-valued function must contain BEGIN/END")
		}
		return functionSubtypeMultiStatementTV, nil
	}
	return functionSubtypeScalar, nil
}

func validateFunctionSource(sql string) (string, error) {
	scanned := strings.ToUpper(sqlparse.StripSQLStringsAndComments(sql))
	if len(findAllKeywordPositions(scanned, "WITH ENCRYPTION")) > 0 ||
		(strings.Contains(scanned, "WITH") && strings.Contains(scanned, "ENCRYPTION")) {
		return "", fmt.Errorf("encrypted function is unsupported: source definition cannot be reconstructed safely")
	}
	return classifyFunctionSubtype(sql)
}

func extractCreateTypeStatement(objectName, sql string) (string, error) {
	normalized := normalizeSQLText(sql, true, true)
	scanned := strings.ToUpper(sqlparse.StripSQLStringsAndComments(normalized))
	for _, start := range findAllKeywordPositions(scanned, "CREATE TYPE") {
		afterKind := start + len("CREATE TYPE")
		rest := strings.TrimSpace(normalized[afterKind:])
		token, consumed := consumeLeadingObjectToken(rest)
		token = stripIdentifierBrackets(strings.TrimSpace(strings.TrimRight(token, ",;")))
		if token == "" || !sameQualifiedObject(token, objectName) {
			continue
		}

		objectOffset := strings.Index(normalized[afterKind:], rest)
		if objectOffset < 0 {
			continue
		}
		afterObject := afterKind + objectOffset + consumed
		asTableEnd := -1
		for _, asPos := range findAllKeywordPositions(scanned[afterObject:], "AS") {
			next, word := readNextWordUpper(scanned, afterObject+asPos+len("AS"))
			if word == "TABLE" {
				asTableEnd = next
				break
			}
		}
		if asTableEnd < 0 {
			return "", fmt.Errorf("CREATE TYPE %s must use AS TABLE", objectName)
		}
		openRel := strings.Index(scanned[asTableEnd:], "(")
		if openRel < 0 {
			return "", fmt.Errorf("CREATE TYPE %s AS TABLE has no opening parenthesis", objectName)
		}
		open := asTableEnd + openRel
		closeIndex, ok := sqlparse.FindMatchingParen(normalized, open)
		if !ok {
			return "", fmt.Errorf("CREATE TYPE %s AS TABLE has no closing parenthesis", objectName)
		}
		end := closeIndex + 1
		for end < len(normalized) && (normalized[end] == ' ' || normalized[end] == '\t' || normalized[end] == '\r' || normalized[end] == '\n') {
			end++
		}
		if end < len(normalized) && normalized[end] == ';' {
			end++
		}
		return normalizeSQLText(strings.TrimSpace(normalized[start:end]), true, true), nil
	}
	return "", fmt.Errorf("CREATE TYPE statement for %s not found", objectName)
}

func validateTableTypeSource(objectName, sql string) (string, error) {
	stmt, err := extractCreateTypeStatement(objectName, sql)
	if err != nil {
		return "", err
	}
	scanned := strings.ToUpper(sqlparse.StripSQLStringsAndComments(sql))
	if len(findAllKeywordPositions(scanned, "CONSTRAINT")) > 0 {
		return "", fmt.Errorf("table type constraints must use unnamed PRIMARY KEY, UNIQUE, CHECK, or DEFAULT syntax; SQL Server does not accept CONSTRAINT names in CREATE TYPE AS TABLE")
	}
	if len(findAllKeywordPositions(scanned, "FOREIGN KEY")) > 0 || len(findAllKeywordPositions(scanned, "REFERENCES")) > 0 {
		return "", fmt.Errorf("table type cannot define foreign keys")
	}
	if len(findAllKeywordPositions(scanned, "CREATE TYPE")) != 1 {
		return "", fmt.Errorf("table type source must contain exactly one CREATE TYPE statement")
	}
	sourceSemantic := sqlparse.CompactUpperSQL(sqlparse.StripSQLStringsAndComments(sql))
	stmtSemantic := sqlparse.CompactUpperSQL(sqlparse.StripSQLStringsAndComments(stmt))
	if strings.TrimSuffix(sourceSemantic, ";") != strings.TrimSuffix(stmtSemantic, ";") {
		return "", fmt.Errorf("table type source must contain only one CREATE TYPE ... AS TABLE statement")
	}
	return stmt, nil
}

func buildDropTypeSQL(objectName string) string {
	return normalizeSQLText("DROP TYPE IF EXISTS "+quoteQualifiedIdentifier(objectName)+";\n", true, true)
}

func buildTableTypeLifecycleSQL(
	objectName, previousType, nextType string,
	history map[string]historyObject,
	desired map[string]desiredObject,
	order []string,
) (string, string, []string, error) {
	historyDependents := make(map[string]historyObject)
	for key, state := range history {
		if state.Category != "functions" && state.Category != "procedures" {
			continue
		}
		if semanticReferencesType(state.SemanticSQL, objectName) {
			historyDependents[key] = state
		}
	}

	historyOrder, err := orderHistoryProgrammability(historyDependents, order)
	if err != nil {
		return "", "", nil, fmt.Errorf("order current TVP dependents: %w", err)
	}

	upDrop := make([]string, 0, len(historyOrder))
	for i := len(historyOrder) - 1; i >= 0; i-- {
		dropSQL, dropErr := buildDropProgrammabilitySQL(historyOrder[i].Category, historyOrder[i].Object)
		if dropErr != nil {
			return "", "", nil, dropErr
		}
		upDrop = append(upDrop, dropSQL)
	}

	upRecreateStates := make([]desiredObject, 0, len(historyDependents))
	for key, current := range historyDependents {
		next, exists := desired[key]
		if !exists {
			// The dependent is intentionally removed in desired state.
			continue
		}
		if hashSemanticByCategory(current.Category, current.SemanticSQL) != hashSemanticByCategory(next.Category, next.SemanticSQL) {
			// A changed dependent has its own later migration. Recreating its
			// desired definition here can reference table/view changes that have
			// not run yet; leave it dropped until that ordered migration executes.
			continue
		}
		upRecreateStates = append(upRecreateStates, desiredObject{
			Category: current.Category, Object: current.Object, SemanticSQL: current.SemanticSQL,
		})
	}
	upRecreate, err := orderDesiredProgrammability(upRecreateStates, order)
	if err != nil {
		return "", "", nil, fmt.Errorf("order desired TVP dependents: %w", err)
	}

	// When rolling back, later dependent migrations have already restored their
	// historical definitions. Drop every historical dependent before restoring
	// the old type, regardless of whether it was recreated during type-up.
	downDrop := make([]string, 0, len(historyOrder))
	for i := len(historyOrder) - 1; i >= 0; i-- {
		dropSQL, dropErr := buildDropProgrammabilitySQL(historyOrder[i].Category, historyOrder[i].Object)
		if dropErr != nil {
			return "", "", nil, dropErr
		}
		downDrop = append(downDrop, dropSQL)
	}

	upParts := append([]string{}, upDrop...)
	upParts = append(upParts, buildDropTypeSQL(objectName), nextType)
	for _, state := range upRecreate {
		upParts = append(upParts, buildDynamicDDLExecution(state.SemanticSQL))
	}

	downParts := append([]string{}, downDrop...)
	downParts = append(downParts, buildDropTypeSQL(objectName), previousType)
	for _, state := range historyOrder {
		downParts = append(downParts, buildDynamicDDLExecution(state.SemanticSQL))
	}

	dependentNames := make([]string, 0, len(historyDependents))
	for _, state := range historyOrder {
		dependentNames = append(dependentNames, state.Category+":"+state.Object)
	}
	return concatSQLSegments(upParts...), concatSQLSegments(downParts...), dependentNames, nil
}

// SQL Server requires CREATE/ALTER VIEW, FUNCTION, and PROCEDURE to be the
// first statement in their batch. A TVP lifecycle necessarily runs those DDL
// statements after DROP TYPE inside one transaction, so execute every module
// definition as its own dynamic batch.
func buildDynamicDDLExecution(sql string) string {
	sql = strings.TrimSpace(sql)
	return "EXEC sys.sp_executesql N'" + strings.ReplaceAll(sql, "'", "''") + "';"
}

func semanticReferencesType(sql, typeName string) bool {
	for _, ref := range sqlparse.ExtractObjectReferences(sql) {
		if ref.Kind == sqlparse.ObjectReferenceType && strings.EqualFold(ref.Name, typeName) {
			return true
		}
	}
	return false
}

func orderHistoryProgrammability(states map[string]historyObject, order []string) ([]historyObject, error) {
	ops := make([]plannedMigration, 0, len(states))
	desired := make(map[string]desiredObject, len(states))
	for key, state := range states {
		ops = append(ops, plannedMigration{Category: state.Category, Object: state.Object})
		desired[key] = desiredObject{Category: state.Category, Object: state.Object, SemanticSQL: state.SemanticSQL}
	}
	result, err := orderPlannedMigrations(ops, desired, order)
	if err != nil {
		return nil, err
	}
	out := make([]historyObject, 0, len(result.Ops))
	for _, op := range result.Ops {
		out = append(out, states[objectKey(op.Category, op.Object)])
	}
	return out, nil
}

func orderDesiredProgrammability(states []desiredObject, order []string) ([]desiredObject, error) {
	ops := make([]plannedMigration, 0, len(states))
	desired := make(map[string]desiredObject, len(states))
	for _, state := range states {
		ops = append(ops, plannedMigration{Category: state.Category, Object: state.Object})
		desired[objectKey(state.Category, state.Object)] = state
	}
	result, err := orderPlannedMigrations(ops, desired, order)
	if err != nil {
		return nil, err
	}
	out := make([]desiredObject, 0, len(result.Ops))
	for _, op := range result.Ops {
		out = append(out, desired[objectKey(op.Category, op.Object)])
	}
	return out, nil
}
