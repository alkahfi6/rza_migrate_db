package engine

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"rza_migrate_db/internal/sqlparse"
)

type columnRename struct{ Old, New string }

func parseColumnRenameDirectives(sql string) (map[string]string, error) {
	out := map[string]string{}
	newNames := map[string]string{}
	for lineNo, line := range strings.Split(strings.ReplaceAll(sql, "\r\n", "\n"), "\n") {
		trimmed := strings.TrimSpace(line)
		prefix := "-- rza:rename-column "
		if !strings.HasPrefix(strings.ToLower(trimmed), prefix) {
			continue
		}
		body := strings.TrimSpace(trimmed[len(prefix):])
		parts := strings.Split(body, "->")
		if len(parts) != 2 {
			return nil, fmt.Errorf("line %d: rename hint must contain exactly one ->", lineNo+1)
		}
		oldName := stripIdentifierBrackets(strings.TrimSpace(parts[0]))
		newName := stripIdentifierBrackets(strings.TrimSpace(parts[1]))
		if oldName == "" || newName == "" || strings.ContainsAny(oldName+newName, ".;'\"") {
			return nil, fmt.Errorf("line %d: rename hint contains an invalid column identifier", lineNo+1)
		}
		oldKey, newKey := oldName, newName
		if oldKey == newKey {
			return nil, fmt.Errorf("line %d: rename source and target are identical", lineNo+1)
		}
		if _, exists := out[oldKey]; exists {
			return nil, fmt.Errorf("line %d: duplicate rename source %s", lineNo+1, oldName)
		}
		if prior, exists := newNames[newKey]; exists {
			return nil, fmt.Errorf("line %d: rename target %s is already used by %s", lineNo+1, newName, prior)
		}
		out[oldKey] = newName
		newNames[newKey] = oldName
	}
	return out, nil
}

func planColumnRenames(previous, next, object string, hints map[string]string, history ...map[string]historyObject) (string, string, bool, error) {
	if len(hints) == 0 {
		return "", "", false, nil
	}
	prevCols, err := parseTableColumns(previous)
	if err != nil {
		return "", "", false, err
	}
	nextCols, err := parseTableColumns(next)
	if err != nil {
		return "", "", false, err
	}
	renames := make([]columnRename, 0, len(hints))
	for oldKey, newName := range hints {
		oldCol, oldOK := prevCols[oldKey]
		newCol, newOK := nextCols[newName]
		if !oldOK || !newOK {
			return "", "", false, fmt.Errorf("rename hint %s -> %s must reference an old column and a new desired column", oldKey, newName)
		}
		if _, collision := prevCols[newName]; collision {
			return "", "", false, fmt.Errorf("rename target already exists in previous table: %s", newName)
		}
		oldCol.Name, newCol.Name = "", ""
		if oldCol != newCol {
			return "", "", false, fmt.Errorf("rename %s -> %s also changes type/modifiers; split rename and alteration into separate migrations", prevCols[oldKey].Name, newName)
		}
		renames = append(renames, columnRename{Old: prevCols[oldKey].Name, New: newName})
	}
	changes, err := classifyTableColumnChanges(previous, next)
	if err != nil {
		return "", "", false, err
	}
	if len(changes.Changed) != 0 || len(changes.Removed) != len(renames) || len(changes.Added) != len(renames) {
		return "", "", false, fmt.Errorf("a rename-hint migration must contain only the declared rename pairs; split other table changes into a later migration")
	}
	sort.Slice(renames, func(i, j int) bool { return renames[i].Old < renames[j].Old })
	expected := previous
	for _, rename := range renames {
		expected = sqlparse.RewriteIdentifierExact(expected, rename.Old, rename.New)
	}
	if hashSemanticByCategory("tables", expected) != hashSemanticByCategory("tables", next) {
		return "", "", false, fmt.Errorf("a rename-hint migration changes semantics beyond exact-case identifier references; split default/check/constraint changes into a later migration")
	}
	up, down := make([]string, 0, len(renames)), make([]string, 0, len(renames))
	for _, rename := range renames {
		up = append(up, buildRenameColumnSQL(object, rename.Old, rename.New))
	}
	for i := len(renames) - 1; i >= 0; i-- {
		down = append(down, buildRenameColumnSQL(object, renames[i].New, renames[i].Old))
	}
	if len(history) > 0 && history[0] != nil {
		oldNames := make([]string, 0, len(renames))
		for _, rename := range renames {
			oldNames = append(oldNames, rename.Old)
		}
		depPlan, depErr := buildDropColumnDependencyPlan(object, oldNames, previous, history[0])
		if depErr != nil {
			return "", "", false, depErr
		}
		filteredConstraints := depPlan.Constraints[:0]
		for _, constraint := range depPlan.Constraints {
			// Named defaults remain attached to the same column_id during sp_rename
			// and don't block it; treating an inline default as a table item would
			// incorrectly attempt to add the column a second time.
			if constraint.Kind != "default" {
				filteredConstraints = append(filteredConstraints, constraint)
			}
		}
		depPlan.Constraints = filteredConstraints
		if len(depPlan.Constraints) > 0 || len(depPlan.Indexes) > 0 {
			upHistory := make(map[string]historyObject, len(history[0]))
			for key, state := range history[0] {
				copyState := state
				if state.Category == "indexes" && sameQualifiedObject(state.Object, object) {
					copyState.SemanticSQL = replaceColumnIdentifiers(state.SemanticSQL, renames)
				}
				upHistory[key] = copyState
			}
			upRestore, restoreErr := buildReAddDroppedDependenciesSQL(next, object, depPlan, upHistory)
			if restoreErr != nil {
				return "", "", false, restoreErr
			}
			downRestore, restoreErr := buildReAddDroppedDependenciesSQL(previous, object, depPlan, history[0])
			if restoreErr != nil {
				return "", "", false, restoreErr
			}
			dropSQL := buildDropDependencySQL(depPlan)
			up = []string{concatSQLSegments(dropSQL, strings.Join(up, "\n"), dynamicSQLBatch(upRestore))}
			down = []string{concatSQLSegments(dropSQL, strings.Join(down, "\n"), dynamicSQLBatch(downRestore))}
		}
	}
	return normalizeSQLText(strings.Join(up, "\n"), true, true), normalizeSQLText(strings.Join(down, "\n"), true, true), true, nil
}

func dynamicSQLBatch(sql string) string {
	trimmed := strings.TrimSpace(sql)
	if trimmed == "" {
		return ""
	}
	escaped := strings.ReplaceAll(trimmed, "'", "''")
	return "EXEC sys.sp_executesql N'" + escaped + "', N'@run_key UNIQUEIDENTIFIER, @version BIGINT, @category NVARCHAR(16), @file NVARCHAR(260), @direction NVARCHAR(8)', @run_key=@run_key, @version=@version, @category=@category, @file=@file, @direction=@direction;"
}

func buildRenameColumnSQL(object, oldName, newName string) string {
	return normalizeSQLText(fmt.Sprintf(`IF COL_LENGTH(N'%s', N'%s') IS NOT NULL
BEGIN
  IF COL_LENGTH(N'%s', N'%s') IS NOT NULL
    THROW 51000, 'column rename target already exists: %s.%s', 1;
  EXEC sys.sp_rename N'%s.%s', N'%s', N'COLUMN';
END
ELSE IF COL_LENGTH(N'%s', N'%s') IS NULL
  THROW 51000, 'column rename source does not exist: %s.%s', 1;`, escapeSQLLiteral(object), escapeSQLLiteral(oldName), escapeSQLLiteral(object), escapeSQLLiteral(newName), escapeSQLLiteral(object), escapeSQLLiteral(newName), escapeSQLLiteral(object), escapeSQLLiteral(oldName), escapeSQLLiteral(newName), escapeSQLLiteral(object), escapeSQLLiteral(newName), escapeSQLLiteral(object), escapeSQLLiteral(oldName)), true, true)
}

var spRenameColumnPattern = regexp.MustCompile(`(?is)EXEC\s+(?:SYS\.)?SP_RENAME\s+N?'([^']+)\.([^'.]+)'\s*,\s*N?'([^']+)'\s*,\s*N?'COLUMN'`)

func extractColumnRenames(sql string) []columnRename {
	matches := spRenameColumnPattern.FindAllStringSubmatch(sql, -1)
	out := make([]columnRename, 0, len(matches))
	for _, match := range matches {
		out = append(out, columnRename{Old: match[2], New: match[3]})
	}
	return out
}

func applyColumnRenamesToCreateTable(createSQL string, renames []columnRename) (string, error) {
	updated := replaceColumnIdentifiers(createSQL, renames)
	if _, err := parseTableColumns(updated); err != nil {
		return "", err
	}
	return normalizeSQLText(updated, true, true), nil
}

func replaceColumnIdentifiers(sql string, renames []columnRename) string {
	updated := sql
	for _, rename := range renames {
		updated = sqlparse.RewriteIdentifierExact(updated, rename.Old, rename.New)
	}
	return updated
}
