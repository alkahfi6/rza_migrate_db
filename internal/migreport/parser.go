package migreport

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"rza_migrate_db/internal/sqlfmt"
	"rza_migrate_db/internal/sqlparse"
)

// ParsedColumn represents a parsed column definition.
type ParsedColumn struct {
	Name        string
	RawType     string
	BaseType    string
	Nullable    bool
	Default     string
	DefaultName string
	Identity    string
	Computed    bool
	Expression  string
}

// ParsedConstraint represents a parsed constraint definition.
type ParsedConstraint struct {
	Name            string
	Type            string // "PRIMARY KEY", "UNIQUE", "FOREIGN KEY", "CHECK", "DEFAULT"
	Columns         []string
	TargetColumn    string // for DEFAULT FOR [col]
	Expression      string // for CHECK or DEFAULT
	ReferencedTable string // for FOREIGN KEY
	ReferencedCols  []string
	RawDefinition   string
}

// ParsedIndex represents a parsed index definition.
type ParsedIndex struct {
	IndexName      string
	TableName      string
	IndexType      string // "CLUSTERED", "NONCLUSTERED", "COLUMNSTORE", "XML", "SPATIAL"
	Columns        []string
	IncludeColumns []string
	Filter         string
	IsUnique       bool
	IsClustered    bool
	RawDefinition  string
	Position       int // byte offset in the cleaned migration SQL
}

// ParsedTableStatement represents table-level DDL actions.
type ParsedTableStatement struct {
	Kind           string // "CREATE_TABLE", "DROP_TABLE", "ADD_COLUMN", "DROP_COLUMN", "RENAME_COLUMN", "ALTER_COLUMN", "ADD_CONSTRAINT", "DROP_CONSTRAINT", "RENAME_CONSTRAINT", "ALTER_CONSTRAINT"
	TableName      string
	SchemaName     string
	TargetName     string
	OldValue       string
	NewValue       string
	DependencyName string
	Column         *ParsedColumn
	Constraint     *ParsedConstraint
	Reversible     string
	Confidence     string
	Notes          string
	Position       int // byte offset in the cleaned migration SQL
}

// ParsedProgrammableStatement represents procedure, view, or function DDL.
type ParsedProgrammableStatement struct {
	Category       string // "procedures", "views", "functions"
	SchemaName     string
	ObjectName     string
	Kind           string // "CREATE", "ALTER", "DROP", "REPLACE_DEFINITION"
	NormalizedSQL  string
	DefinitionHash string
	Parameters     string
	ReturnType     string
	Body           string
	RawSQL         string
	Confidence     string
	Notes          string
}

// ParsedTypeStatement represents type DDL.
type ParsedTypeStatement struct {
	SchemaName    string
	TypeName      string
	Kind          string // "CREATE_TYPE", "DROP_TYPE", "ALTER_TYPE", "REPLACE_TYPE"
	RawDefinition string
	Reversible    string
	Confidence    string
	Notes         string
}

// StripComments removes single-line (--) and block (/* ... */) comments outside strings.
func StripComments(sql string) string {
	var b strings.Builder
	runes := []rune(sql)
	n := len(runes)
	inString := false
	var stringChar rune

	for i := 0; i < n; i++ {
		r := runes[i]
		if inString {
			b.WriteRune(r)
			if r == stringChar {
				if i+1 < n && runes[i+1] == stringChar {
					b.WriteRune(runes[i+1])
					i++
				} else {
					inString = false
				}
			}
			continue
		}

		if r == '\'' || r == '"' {
			inString = true
			stringChar = r
			b.WriteRune(r)
			continue
		}

		// Single line comment --
		if r == '-' && i+1 < n && runes[i+1] == '-' {
			for i < n && runes[i] != '\n' {
				i++
			}
			if i < n {
				b.WriteRune('\n')
			}
			continue
		}

		// Block comment /* ... */
		if r == '/' && i+1 < n && runes[i+1] == '*' {
			i += 2
			for i+1 < n && !(runes[i] == '*' && runes[i+1] == '/') {
				i++
			}
			i++ // skip '/'
			b.WriteRune(' ')
			continue
		}

		b.WriteRune(r)
	}
	return b.String()
}

// CleanSQLForAnalysis strips comments, transaction wrappers, audit logs, and prechecks.
func CleanSQLForAnalysis(rawSQL string) string {
	noComments := StripComments(rawSQL)
	sql := sqlfmt.NormalizeSQLText(noComments, true, true)
	return sql
}

// ComputeNormalizedHash computes a deterministic SHA-256 of cleaned SQL ignoring comments, whitespace, and casing.
func ComputeNormalizedHash(sql string) string {
	noComments := StripComments(sql)
	cleaned := sqlfmt.NormalizeSQLText(noComments, true, true)
	words := strings.Fields(cleaned)
	for i, w := range words {
		if !strings.HasPrefix(w, "'") && !strings.HasPrefix(w, "N'") {
			words[i] = strings.ToUpper(w)
		}
	}
	normalized := strings.Join(words, " ")
	h := sha256.Sum256([]byte(normalized))
	return hex.EncodeToString(h[:])
}

// StripSurroundingBrackets removes enclosing square brackets e.g. [Customer] -> Customer
func StripSurroundingBrackets(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") && len(s) >= 2 {
		return s[1 : len(s)-1]
	}
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") && len(s) >= 2 {
		return s[1 : len(s)-1]
	}
	if strings.HasPrefix(s, "'") && strings.HasSuffix(s, "'") && len(s) >= 2 {
		return s[1 : len(s)-1]
	}
	return s
}

// CleanIdentifier strips schema and brackets: "dbo.[Customer]" -> "Customer"
func CleanIdentifier(ident string) string {
	ident = strings.TrimSpace(ident)
	if dot := strings.LastIndex(ident, "."); dot >= 0 {
		ident = ident[dot+1:]
	}
	return StripSurroundingBrackets(ident)
}

// ExtractSchemaAndObject splits "dbo.Customer" -> ("dbo", "Customer")
func ExtractSchemaAndObject(ident string, defaultSchema string) (string, string) {
	ident = strings.TrimSpace(ident)
	if dot := strings.Index(ident, "."); dot >= 0 {
		schema := StripSurroundingBrackets(ident[:dot])
		obj := StripSurroundingBrackets(ident[dot+1:])
		if schema == "" {
			schema = defaultSchema
		}
		return schema, obj
	}
	return defaultSchema, StripSurroundingBrackets(ident)
}

// ParseStatementsFromSQL extracts semantic table operations from a migration script.
func ParseStatementsFromSQL(category, defaultSchema, defaultObject, sql string) ([]ParsedTableStatement, []ParsedIndex, *ParsedProgrammableStatement, *ParsedTypeStatement) {
	cleaned := CleanSQLForAnalysis(sql)
	if strings.TrimSpace(cleaned) == "" {
		return nil, nil, nil, nil
	}
	upper := strings.ToUpper(cleaned)

	var tableStmts []ParsedTableStatement
	var indexes []ParsedIndex
	var progStmt *ParsedProgrammableStatement
	var typeStmt *ParsedTypeStatement

	switch category {
	case "tables":
		tableStmts, indexes = parseTableAndIndexSQL(defaultSchema, defaultObject, cleaned, upper)
	case "indexes":
		indexes = parseIndexSQL(defaultSchema, defaultObject, cleaned)
	case "procedures", "views", "functions":
		progStmt = parseProgrammableSQL(category, defaultSchema, defaultObject, cleaned)
	case "types":
		typeStmt = parseTypeSQL(defaultSchema, defaultObject, cleaned)
	default:
		// Try parsing table, index, programmable or type
		if strings.Contains(upper, "CREATE TABLE") || strings.Contains(upper, "ALTER TABLE") || strings.Contains(upper, "DROP TABLE") {
			tableStmts, indexes = parseTableAndIndexSQL(defaultSchema, defaultObject, cleaned, upper)
		} else if strings.Contains(upper, "CREATE INDEX") || strings.Contains(upper, "CREATE UNIQUE INDEX") || strings.Contains(upper, "DROP INDEX") {
			indexes = parseIndexSQL(defaultSchema, defaultObject, cleaned)
		} else if strings.Contains(upper, "CREATE PROCEDURE") || strings.Contains(upper, "CREATE OR ALTER PROCEDURE") ||
			strings.Contains(upper, "CREATE VIEW") || strings.Contains(upper, "CREATE OR ALTER VIEW") ||
			strings.Contains(upper, "CREATE FUNCTION") || strings.Contains(upper, "CREATE OR ALTER FUNCTION") ||
			strings.Contains(upper, "DROP PROCEDURE") || strings.Contains(upper, "DROP PROC") ||
			strings.Contains(upper, "DROP VIEW") || strings.Contains(upper, "DROP FUNCTION") {
			progStmt = parseProgrammableSQL(category, defaultSchema, defaultObject, cleaned)
		} else if strings.Contains(upper, "CREATE TYPE") || strings.Contains(upper, "DROP TYPE") {
			typeStmt = parseTypeSQL(defaultSchema, defaultObject, cleaned)
		}
	}

	return tableStmts, indexes, progStmt, typeStmt
}

func parseTableAndIndexSQL(defaultSchema, defaultObject, sql, upper string) ([]ParsedTableStatement, []ParsedIndex) {
	var tableStmts []ParsedTableStatement
	var indexes []ParsedIndex

	// Check DROP TABLE
	dropMatches := dropTableRegex.FindAllStringSubmatchIndex(sql, -1)
	for _, match := range dropMatches {
		if len(match) >= 4 {
			schema, tbl := ExtractSchemaAndObject(sql[match[2]:match[3]], defaultSchema)
			tableStmts = append(tableStmts, ParsedTableStatement{
				Kind:       OpDropTable,
				SchemaName: schema,
				TableName:  tbl,
				TargetName: tbl,
				Reversible: ReversibleNo,
				Confidence: ConfidenceHigh,
				Position:   match[0],
			})
		}
	}

	// Check all CREATE TABLE statements in the script
	searchUpper := upper
	searchOffset := 0
	for {
		idx := strings.Index(searchUpper, "CREATE TABLE ")
		if idx < 0 {
			break
		}
		actualIdx := searchOffset + idx
		stmtRest := sql[actualIdx:]
		openIdx := strings.Index(stmtRest, "(")
		if openIdx > 0 {
			header := stmtRest[:openIdx]
			headerTokens := strings.Fields(header)
			tableName := defaultObject
			schemaName := defaultSchema
			if len(headerTokens) >= 3 {
				schemaName, tableName = ExtractSchemaAndObject(headerTokens[2], defaultSchema)
			}

			closeIdx, ok := sqlparse.FindMatchingParen(stmtRest, openIdx)
			if ok {
				inside := stmtRest[openIdx+1 : closeIdx]
				tableStmts = append(tableStmts, ParsedTableStatement{
					Kind:       OpCreateTable,
					SchemaName: schemaName,
					TableName:  tableName,
					TargetName: tableName,
					Reversible: ReversibleYes,
					Confidence: ConfidenceHigh,
					Position:   actualIdx,
				})

				cols, constraints := parseTableBody(inside)
				for _, col := range cols {
					c := col
					tableStmts = append(tableStmts, ParsedTableStatement{
						Kind:       OpAddColumn,
						SchemaName: schemaName,
						TableName:  tableName,
						TargetName: c.Name,
						NewValue:   c.RawType,
						Column:     &c,
						Reversible: ReversibleYes,
						Confidence: ConfidenceHigh,
						Position:   actualIdx,
					})
				}
				for _, con := range constraints {
					cn := con
					tableStmts = append(tableStmts, ParsedTableStatement{
						Kind:           OpAddConstraint,
						SchemaName:     schemaName,
						TableName:      tableName,
						TargetName:     cn.Name,
						NewValue:       cn.RawDefinition,
						DependencyName: cn.ReferencedTable,
						Constraint:     &cn,
						Reversible:     ReversibleYes,
						Confidence:     ConfidenceHigh,
						Position:       actualIdx,
					})
				}

				// closeIdx is relative to stmtRest, whereas searchUpper starts at
				// searchOffset. Include idx so the next search starts after this
				// complete CREATE TABLE body rather than somewhere before its
				// header. Without it, a wrapper prefix makes the same statement
				// visible repeatedly and duplicates its report operations.
				advance := idx + closeIdx + 1
				if advance < len(searchUpper) {
					searchUpper = searchUpper[advance:]
					searchOffset += advance
					continue
				}
			}
		}

		advance := idx + len("CREATE TABLE ")
		if advance < len(searchUpper) {
			searchUpper = searchUpper[advance:]
			searchOffset += advance
		} else {
			break
		}
	}

	// Check ALTER TABLE statements
	alterMatches := alterTableRegex.FindAllStringSubmatchIndex(sql, -1)
	for _, m := range alterMatches {
		if len(m) < 8 {
			continue
		}
		position := m[0]
		tableIdent := sql[m[2]:m[3]]
		action := strings.Join(strings.Fields(strings.ToUpper(strings.TrimSpace(sql[m[4]:m[5]]))), " ")
		rest := strings.TrimSpace(sql[m[6]:m[7]])
		schemaName, tableName := ExtractSchemaAndObject(tableIdent, defaultSchema)

		switch {
		case strings.HasPrefix(action, "ADD"):
			addBody := rest
			if strings.HasPrefix(action, "ADD CONSTRAINT") {
				addBody = "CONSTRAINT " + addBody
			}
			for _, item := range splitTopLevelByComma(addBody) {
				item = strings.TrimSpace(item)
				if item == "" {
					continue
				}
				if isConstraintDefinition(item) {
					con := parseSingleConstraint(item)
					tableStmts = append(tableStmts, ParsedTableStatement{
						Kind:           OpAddConstraint,
						SchemaName:     schemaName,
						TableName:      tableName,
						TargetName:     con.Name,
						NewValue:       con.RawDefinition,
						DependencyName: con.ReferencedTable,
						Constraint:     &con,
						Reversible:     ReversibleYes,
						Confidence:     ConfidenceHigh,
						Position:       position,
					})
					continue
				}

				c := parseSingleColumn(item)
				if c.Name == "" {
					continue
				}
				tableStmts = append(tableStmts, ParsedTableStatement{
					Kind:       OpAddColumn,
					SchemaName: schemaName,
					TableName:  tableName,
					TargetName: c.Name,
					NewValue:   c.RawType,
					Column:     &c,
					Reversible: ReversibleYes,
					Confidence: ConfidenceHigh,
					Position:   position,
				})
				if c.DefaultName != "" {
					con := ParsedConstraint{
						Name:          c.DefaultName,
						Type:          "DEFAULT",
						TargetColumn:  c.Name,
						Expression:    c.Default,
						RawDefinition: fmt.Sprintf("DEFAULT %s FOR [%s]", c.Default, c.Name),
					}
					tableStmts = append(tableStmts, ParsedTableStatement{
						Kind:       OpAddConstraint,
						SchemaName: schemaName,
						TableName:  tableName,
						TargetName: c.DefaultName,
						NewValue:   con.RawDefinition,
						Constraint: &con,
						Reversible: ReversibleYes,
						Confidence: ConfidenceHigh,
						Position:   position,
					})
				}
			}
		case action == "DROP COLUMN":
			for _, col := range strings.Split(rest, ",") {
				colName := CleanIdentifier(col)
				if colName != "" {
					tableStmts = append(tableStmts, ParsedTableStatement{
						Kind:       OpDropColumn,
						SchemaName: schemaName,
						TableName:  tableName,
						TargetName: colName,
						Reversible: ReversibleYes,
						Confidence: ConfidenceHigh,
						Position:   position,
					})
				}
			}
		case action == "DROP CONSTRAINT":
			for _, item := range splitTopLevelByComma(rest) {
				item = strings.TrimSpace(item)
				if strings.HasPrefix(strings.ToUpper(item), "IF EXISTS ") {
					item = strings.TrimSpace(item[len("IF EXISTS "):])
				}
				conName := CleanIdentifier(item)
				if conName != "" {
					tableStmts = append(tableStmts, ParsedTableStatement{
						Kind:       OpDropConstraint,
						SchemaName: schemaName,
						TableName:  tableName,
						TargetName: conName,
						Reversible: ReversibleYes,
						Confidence: ConfidenceHigh,
						Position:   position,
					})
				}
			}
		case action == "ALTER COLUMN":
			col := parseSingleColumn(rest)
			tableStmts = append(tableStmts, ParsedTableStatement{
				Kind:       OpAlterColumnType,
				SchemaName: schemaName,
				TableName:  tableName,
				TargetName: col.Name,
				NewValue:   col.RawType,
				Column:     &col,
				Reversible: ReversibleYes,
				Confidence: ConfidenceHigh,
				Position:   position,
			})
			if col.Nullable {
				tableStmts = append(tableStmts, ParsedTableStatement{
					Kind:       OpAlterNullability,
					SchemaName: schemaName,
					TableName:  tableName,
					TargetName: col.Name,
					NewValue:   "NULL",
					Column:     &col,
					Reversible: ReversibleYes,
					Confidence: ConfidenceHigh,
					Position:   position,
				})
			} else {
				tableStmts = append(tableStmts, ParsedTableStatement{
					Kind:       OpAlterNullability,
					SchemaName: schemaName,
					TableName:  tableName,
					TargetName: col.Name,
					NewValue:   "NOT NULL",
					Column:     &col,
					Reversible: ReversibleYes,
					Confidence: ConfidenceHigh,
					Position:   position,
				})
			}
		}
	}

	// Check sp_rename statements (for column, index, or constraint rename)
	spRenameMatches := spRenameRegex.FindAllStringSubmatchIndex(sql, -1)
	for _, m := range spRenameMatches {
		if len(m) < 6 {
			continue
		}
		position := m[0]
		target := StripSurroundingBrackets(sql[m[2]:m[3]])
		newName := StripSurroundingBrackets(sql[m[4]:m[5]])
		renameType := ""
		if len(m) >= 8 && m[6] >= 0 {
			renameType = strings.ToUpper(StripSurroundingBrackets(sql[m[6]:m[7]]))
		}

		parts := strings.Split(target, ".")
		oldName := target
		tableName := defaultObject
		schemaName := defaultSchema
		if len(parts) >= 3 {
			schemaName = StripSurroundingBrackets(parts[0])
			tableName = StripSurroundingBrackets(parts[1])
			oldName = StripSurroundingBrackets(parts[2])
		} else if len(parts) == 2 {
			tableName = StripSurroundingBrackets(parts[0])
			oldName = StripSurroundingBrackets(parts[1])
		}

		if renameType == "INDEX" {
			indexes = append(indexes, ParsedIndex{
				IndexName:     newName,
				TableName:     tableName,
				IndexType:     "RENAME",
				RawDefinition: oldName,
				Position:      position,
			})
		} else if renameType == "COLUMN" || strings.Contains(strings.ToUpper(sql), "'COLUMN'") {
			tableStmts = append(tableStmts, ParsedTableStatement{
				Kind:       OpRenameColumn,
				SchemaName: schemaName,
				TableName:  tableName,
				TargetName: newName,
				OldValue:   oldName,
				NewValue:   newName,
				Reversible: ReversibleYes,
				Confidence: ConfidenceHigh,
				Position:   position,
			})
		} else {
			tableStmts = append(tableStmts, ParsedTableStatement{
				Kind:       OpRenameConstraint,
				SchemaName: schemaName,
				TableName:  tableName,
				TargetName: newName,
				OldValue:   oldName,
				NewValue:   newName,
				Reversible: ReversibleYes,
				Confidence: ConfidenceHigh,
				Position:   position,
			})
		}
	}

	// Also check any CREATE INDEX in table SQL
	idxList := parseIndexSQL(defaultSchema, defaultObject, sql)
	indexes = append(indexes, idxList...)
	sort.SliceStable(tableStmts, func(i, j int) bool {
		return tableStmts[i].Position < tableStmts[j].Position
	})
	sort.SliceStable(indexes, func(i, j int) bool {
		return indexes[i].Position < indexes[j].Position
	})

	return tableStmts, indexes
}

func parseIndexSQL(defaultSchema, defaultObject, sql string) []ParsedIndex {
	var indexes []ParsedIndex
	locs := createIndexRegex.FindAllStringSubmatchIndex(sql, -1)
	for _, loc := range locs {
		isUnique := false
		if loc[2] >= 0 && loc[3] >= 0 && strings.TrimSpace(sql[loc[2]:loc[3]]) != "" {
			isUnique = true
		}
		clusteredStr := ""
		if loc[4] >= 0 && loc[5] >= 0 {
			clusteredStr = strings.ToUpper(strings.TrimSpace(sql[loc[4]:loc[5]]))
		}
		isClustered := clusteredStr == "CLUSTERED"
		indexType := "NONCLUSTERED"
		if isClustered {
			indexType = "CLUSTERED"
		}
		idxName := CleanIdentifier(sql[loc[6]:loc[7]])
		tblIdent := sql[loc[8]:loc[9]]
		colsInside := sql[loc[10]:loc[11]]

		_, tblName := ExtractSchemaAndObject(tblIdent, defaultSchema)
		if tblName == "" {
			tblName = defaultObject
		}

		var cols []string
		for _, c := range strings.Split(colsInside, ",") {
			cClean := cleanIndexColumn(c)
			if cClean != "" {
				cols = append(cols, cClean)
			}
		}

		var includeCols []string
		filter := "NULL"

		// Scoped tail immediately after this index declaration
		tail := sql[loc[1]:]
		if semiIdx := strings.Index(tail, ";"); semiIdx >= 0 {
			tail = tail[:semiIdx]
		} else if endIdx := strings.Index(strings.ToUpper(tail), "\nCREATE "); endIdx >= 0 {
			tail = tail[:endIdx]
		}

		if incMatch := includeRegex.FindStringSubmatch(tail); len(incMatch) > 1 {
			for _, c := range strings.Split(incMatch[1], ",") {
				cClean := CleanIdentifier(c)
				if cClean != "" {
					includeCols = append(includeCols, cClean)
				}
			}
		}

		if whereMatch := whereFilterRegex.FindStringSubmatch(tail); len(whereMatch) > 1 {
			filter = strings.TrimSpace(whereMatch[1])
		}

		indexes = append(indexes, ParsedIndex{
			IndexName:      idxName,
			TableName:      tblName,
			IndexType:      indexType,
			Columns:        cols,
			IncludeColumns: includeCols,
			Filter:         filter,
			IsUnique:       isUnique,
			IsClustered:    isClustered,
			RawDefinition:  sql[loc[0]:loc[1]] + tail,
			Position:       loc[0],
		})
	}

	// Check DROP INDEX
	dropLocs := dropIndexRegex.FindAllStringSubmatchIndex(sql, -1)
	for _, loc := range dropLocs {
		if len(loc) < 6 {
			continue
		}
		idxName := CleanIdentifier(sql[loc[2]:loc[3]])
		tblName := CleanIdentifier(sql[loc[4]:loc[5]])
		indexes = append(indexes, ParsedIndex{
			IndexName: idxName,
			TableName: tblName,
			IndexType: "DROP",
			Position:  loc[0],
		})
	}

	sort.SliceStable(indexes, func(i, j int) bool {
		return indexes[i].Position < indexes[j].Position
	})

	return indexes
}

func cleanIndexColumn(input string) string {
	fields := strings.Fields(strings.TrimSpace(input))
	if len(fields) == 0 {
		return ""
	}
	name := CleanIdentifier(fields[0])
	if len(fields) > 1 && (strings.EqualFold(fields[1], "ASC") || strings.EqualFold(fields[1], "DESC")) {
		return name + " " + strings.ToUpper(fields[1])
	}
	return name
}

func parseProgrammableSQL(category, defaultSchema, defaultObject, sql string) *ParsedProgrammableStatement {
	upper := strings.ToUpper(sql)
	kind := OpProgAlter
	if strings.Contains(upper, "CREATE OR ALTER") {
		kind = OpProgAlter
	} else if strings.Contains(upper, "CREATE ") {
		kind = OpProgCreate
	} else if strings.Contains(upper, "DROP ") {
		kind = OpProgDrop
	}

	// Extract object name
	var objName = defaultObject
	var schemaName = defaultSchema

	if match := progHeaderRegex.FindStringSubmatch(sql); len(match) > 1 {
		schemaName, objName = ExtractSchemaAndObject(match[1], defaultSchema)
	}

	normHash := ComputeNormalizedHash(sql)
	body := extractProgrammableBody(sql)
	params := extractParameters(sql)

	return &ParsedProgrammableStatement{
		Category:       category,
		SchemaName:     schemaName,
		ObjectName:     objName,
		Kind:           kind,
		NormalizedSQL:  sql,
		DefinitionHash: normHash,
		Parameters:     params,
		Body:           body,
		RawSQL:         sql,
		Confidence:     ConfidenceHigh,
	}
}

func parseTypeSQL(defaultSchema, defaultObject, sql string) *ParsedTypeStatement {
	upper := strings.ToUpper(sql)
	kind := OpCreateType
	if strings.Contains(upper, "DROP TYPE") {
		kind = OpDropType
	} else if strings.Contains(upper, "CREATE TYPE") {
		kind = OpCreateType
	}

	typeName := defaultObject
	schemaName := defaultSchema
	if kind == OpDropType {
		if match := dropTypeRegex.FindStringSubmatch(sql); len(match) > 1 {
			schemaName, typeName = ExtractSchemaAndObject(match[1], defaultSchema)
		}
	} else {
		if match := createTypeRegex.FindStringSubmatch(sql); len(match) > 1 {
			schemaName, typeName = ExtractSchemaAndObject(match[1], defaultSchema)
		}
	}

	return &ParsedTypeStatement{
		SchemaName:    schemaName,
		TypeName:      typeName,
		Kind:          kind,
		RawDefinition: sql,
		Reversible:    ReversibleYes,
		Confidence:    ConfidenceHigh,
	}
}

// Helpers for column and constraint parsing inside CREATE TABLE or ALTER TABLE

func parseTableBody(body string) ([]ParsedColumn, []ParsedConstraint) {
	var cols []ParsedColumn
	var constraints []ParsedConstraint

	items := splitTopLevelByComma(body)
	for _, item := range items {
		line := strings.TrimSpace(item)
		if line == "" {
			continue
		}
		upper := strings.ToUpper(line)
		if strings.HasPrefix(upper, "CONSTRAINT ") ||
			strings.HasPrefix(upper, "PRIMARY KEY") ||
			strings.HasPrefix(upper, "UNIQUE") ||
			strings.HasPrefix(upper, "FOREIGN KEY") ||
			strings.HasPrefix(upper, "CHECK") {
			con := parseSingleConstraint(line)
			constraints = append(constraints, con)
		} else {
			col := parseSingleColumn(line)
			if col.Name != "" {
				cols = append(cols, col)
				if col.DefaultName != "" {
					constraints = append(constraints, ParsedConstraint{
						Name:          col.DefaultName,
						Type:          "DEFAULT",
						TargetColumn:  col.Name,
						Expression:    col.Default,
						RawDefinition: fmt.Sprintf("CONSTRAINT [%s] DEFAULT %s", col.DefaultName, col.Default),
					})
				}
			}
		}
	}
	return cols, constraints
}

func isConstraintDefinition(input string) bool {
	upper := strings.ToUpper(strings.TrimSpace(input))
	return strings.HasPrefix(upper, "CONSTRAINT ") ||
		strings.HasPrefix(upper, "PRIMARY KEY") ||
		strings.HasPrefix(upper, "UNIQUE") ||
		strings.HasPrefix(upper, "FOREIGN KEY") ||
		strings.HasPrefix(upper, "CHECK") ||
		strings.HasPrefix(upper, "DEFAULT")
}

// splitLeadingIdentifier separates the first SQL identifier from its definition.
// A delimited identifier may legally contain spaces (for example, [Tax On
// Accrued]), so strings.Fields cannot be used before it has been consumed.
func splitLeadingIdentifier(input string) (string, string) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", ""
	}
	if input[0] == '[' {
		for i := 1; i < len(input); i++ {
			if input[i] != ']' {
				continue
			}
			// SQL Server escapes a literal ] inside a bracket identifier as ]].
			if i+1 < len(input) && input[i+1] == ']' {
				i++
				continue
			}
			return input[:i+1], strings.TrimSpace(input[i+1:])
		}
	}
	if input[0] == '"' {
		for i := 1; i < len(input); i++ {
			if input[i] != '"' {
				continue
			}
			if i+1 < len(input) && input[i+1] == '"' {
				i++
				continue
			}
			return input[:i+1], strings.TrimSpace(input[i+1:])
		}
	}
	fields := strings.Fields(input)
	if len(fields) == 0 {
		return "", ""
	}
	ident := fields[0]
	return ident, strings.TrimSpace(input[len(ident):])
}

func parseSingleColumn(line string) ParsedColumn {
	line = strings.TrimSpace(line)
	if line == "" {
		return ParsedColumn{}
	}

	ident, rest := splitLeadingIdentifier(line)
	if ident == "" {
		return ParsedColumn{}
	}

	colName := CleanIdentifier(ident)
	upperRest := strings.ToUpper(rest)

	// Check computed column AS (...)
	if strings.HasPrefix(upperRest, "AS ") || strings.HasPrefix(upperRest, "AS(") {
		return ParsedColumn{
			Name:       colName,
			BaseType:   "COMPUTED",
			RawType:    "COMPUTED",
			Nullable:   true,
			Computed:   true,
			Expression: strings.TrimSpace(rest[2:]),
		}
	}

	nullable := true
	if strings.Contains(upperRest, " NOT NULL") {
		nullable = false
	} else if strings.Contains(upperRest, " NULL") {
		nullable = true
	}

	// Extract type token
	rawType := ""
	restFields := strings.Fields(rest)
	if len(restFields) > 0 {
		rawType = restFields[0]
		// If type has parenthesis e.g. VARCHAR(100)
		if strings.Contains(rawType, "(") && !strings.Contains(rawType, ")") {
			// Include next fields until closing parenthesis
			for i := 1; i < len(restFields); i++ {
				rawType += " " + restFields[i]
				if strings.Contains(restFields[i], ")") {
					break
				}
			}
		}
	}

	// Check DEFAULT constraint
	defaultVal := ""
	defaultName := ""
	if dfIdx := strings.Index(upperRest, " DEFAULT"); dfIdx >= 0 {
		dfPart := strings.TrimSpace(rest[dfIdx+len(" DEFAULT"):])
		defaultVal = dfPart
		// Check if CONSTRAINT df_name DEFAULT
		beforeDf := rest[:dfIdx]
		if conIdx := strings.LastIndex(strings.ToUpper(beforeDf), "CONSTRAINT "); conIdx >= 0 {
			namePart := strings.TrimSpace(beforeDf[conIdx+len("CONSTRAINT "):])
			defaultName = CleanIdentifier(namePart)
		}
	}

	baseType := rawType
	if idx := strings.Index(baseType, "("); idx > 0 {
		baseType = strings.TrimSpace(baseType[:idx])
	}
	baseType = strings.ToUpper(baseType)

	return ParsedColumn{
		Name:        colName,
		RawType:     rawType,
		BaseType:    baseType,
		Nullable:    nullable,
		Default:     defaultVal,
		DefaultName: defaultName,
	}
}

func parseSingleConstraint(line string) ParsedConstraint {
	line = strings.TrimSpace(line)
	upper := strings.ToUpper(line)

	name := ""
	if strings.HasPrefix(upper, "CONSTRAINT ") {
		rest := strings.TrimSpace(line[len("CONSTRAINT "):])
		parts := strings.Fields(rest)
		if len(parts) > 0 {
			name = CleanIdentifier(parts[0])
			line = strings.TrimSpace(rest[len(parts[0]):])
			upper = strings.ToUpper(line)
		}
	}

	conType := "UNKNOWN"
	var cols []string
	var refTable string
	var refCols []string
	var expr string
	var targetCol string

	switch {
	case strings.HasPrefix(upper, "PRIMARY KEY"):
		conType = "PRIMARY KEY"
		cols = extractParenColumns(line)
	case strings.HasPrefix(upper, "UNIQUE"):
		conType = "UNIQUE"
		cols = extractParenColumns(line)
	case strings.HasPrefix(upper, "FOREIGN KEY"):
		conType = "FOREIGN KEY"
		cols = extractParenColumns(line)
		if refIdx := strings.Index(upper, "REFERENCES "); refIdx >= 0 {
			refPart := line[refIdx+len("REFERENCES "):]
			openIdx := strings.Index(refPart, "(")
			if openIdx > 0 {
				refTable = CleanIdentifier(refPart[:openIdx])
				refCols = extractParenColumns(refPart[openIdx:])
			}
		}
	case strings.HasPrefix(upper, "CHECK"):
		conType = "CHECK"
		if open := strings.Index(line, "("); open >= 0 {
			if closeIdx, ok := sqlparse.FindMatchingParen(line, open); ok {
				expr = strings.TrimSpace(line[open+1 : closeIdx])
			}
		}
	case strings.HasPrefix(upper, "DEFAULT"):
		conType = "DEFAULT"
		// DEFAULT (...) FOR [col]
		if forIdx := strings.Index(upper, " FOR "); forIdx >= 0 {
			expr = strings.TrimSpace(line[len("DEFAULT "):forIdx])
			targetCol = CleanIdentifier(line[forIdx+len(" FOR "):])
		} else {
			expr = strings.TrimSpace(line[len("DEFAULT "):])
		}
	}

	return ParsedConstraint{
		Name:            name,
		Type:            conType,
		Columns:         cols,
		TargetColumn:    targetCol,
		Expression:      expr,
		ReferencedTable: refTable,
		ReferencedCols:  refCols,
		RawDefinition:   line,
	}
}

func extractParenColumns(s string) []string {
	open := strings.Index(s, "(")
	if open < 0 {
		return nil
	}
	closeIdx, ok := sqlparse.FindMatchingParen(s, open)
	if !ok {
		return nil
	}
	inside := s[open+1 : closeIdx]
	var cols []string
	for _, c := range strings.Split(inside, ",") {
		clean := CleanIdentifier(c)
		if clean != "" {
			cols = append(cols, clean)
		}
	}
	return cols
}

func splitTopLevelByComma(s string) []string {
	var parts []string
	depth := 0
	start := 0
	inQuote := false
	var quoteChar rune

	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if inQuote {
			if r == quoteChar {
				inQuote = false
			}
			continue
		}
		if r == '\'' || r == '"' {
			inQuote = true
			quoteChar = r
			continue
		}
		if r == '(' || r == '[' {
			depth++
		} else if r == ')' || r == ']' {
			if depth > 0 {
				depth--
			}
		} else if r == ',' && depth == 0 {
			parts = append(parts, string(runes[start:i]))
			start = i + 1
		}
	}
	if start < len(runes) {
		parts = append(parts, string(runes[start:]))
	}
	return parts
}

func extractProgrammableBody(sql string) string {
	upper := strings.ToUpper(sql)
	asIdx := strings.Index(upper, " AS\n")
	if asIdx < 0 {
		asIdx = strings.Index(upper, " AS ")
	}
	if asIdx < 0 {
		asIdx = strings.Index(upper, "\nAS\n")
	}
	if asIdx >= 0 {
		return strings.TrimSpace(sql[asIdx+4:])
	}
	return sql
}

func extractParameters(sql string) string {
	upper := strings.ToUpper(sql)
	asIdx := strings.Index(upper, " AS")
	if asIdx > 0 {
		header := sql[:asIdx]
		if atIdx := strings.Index(header, "@"); atIdx > 0 {
			return strings.TrimSpace(header[atIdx:])
		}
	}
	return ""
}

// DetectChangedSections compares two programmable versions and identifies changed sections.
func DetectChangedSections(oldSQL, newSQL string) string {
	if oldSQL == "" {
		return SectionBody
	}
	oldUpper := strings.ToUpper(oldSQL)
	newUpper := strings.ToUpper(newSQL)

	var sections []string

	// Check PARAMETERS
	oldParams := extractParameters(oldSQL)
	newParams := extractParameters(newSQL)
	if ComputeNormalizedHash(oldParams) != ComputeNormalizedHash(newParams) {
		sections = append(sections, SectionParameters)
	}

	// Check DECLARE
	if (strings.Contains(oldUpper, "DECLARE ") || strings.Contains(newUpper, "DECLARE ")) &&
		ComputeNormalizedHash(extractSection(oldSQL, "DECLARE", ";")) != ComputeNormalizedHash(extractSection(newSQL, "DECLARE", ";")) {
		sections = append(sections, SectionDeclare)
	}

	// Check JOIN
	if (strings.Contains(oldUpper, "JOIN ") || strings.Contains(newUpper, "JOIN ")) &&
		strings.Count(oldUpper, "JOIN ") != strings.Count(newUpper, "JOIN ") {
		sections = append(sections, SectionJoin)
	}

	// Check WHERE
	if (strings.Contains(oldUpper, "WHERE ") || strings.Contains(newUpper, "WHERE ")) &&
		ComputeNormalizedHash(extractSection(oldSQL, "WHERE", "ORDER BY")) != ComputeNormalizedHash(extractSection(newSQL, "WHERE", "ORDER BY")) {
		sections = append(sections, SectionWhere)
	}

	// Check DML
	hasOldDML := strings.Contains(oldUpper, "INSERT ") || strings.Contains(oldUpper, "UPDATE ") || strings.Contains(oldUpper, "DELETE ") || strings.Contains(oldUpper, "MERGE ")
	hasNewDML := strings.Contains(newUpper, "INSERT ") || strings.Contains(newUpper, "UPDATE ") || strings.Contains(newUpper, "DELETE ") || strings.Contains(newUpper, "MERGE ")
	if hasOldDML || hasNewDML {
		sections = append(sections, SectionDML)
	}

	// Check RETURN_TYPE for functions
	if strings.Contains(oldUpper, "RETURNS ") || strings.Contains(newUpper, "RETURNS ") {
		oldRet := extractSection(oldSQL, "RETURNS", "AS")
		newRet := extractSection(newSQL, "RETURNS", "AS")
		if oldRet != "" && newRet != "" && ComputeNormalizedHash(oldRet) != ComputeNormalizedHash(newRet) {
			sections = append(sections, SectionReturnType)
		}
	}

	// Check QUERY (SELECT)
	if strings.Contains(newUpper, "SELECT ") {
		sections = append(sections, SectionQuery)
	}

	// If no specific section found or changes are broad, fallback to BODY
	if len(sections) == 0 {
		sections = append(sections, SectionBody)
	}

	return strings.Join(sections, "|")
}

func extractSection(sql, startKeyword, endKeyword string) string {
	upper := strings.ToUpper(sql)
	sIdx := strings.Index(upper, startKeyword)
	if sIdx < 0 {
		return ""
	}
	rest := sql[sIdx:]
	if endKeyword != "" {
		eIdx := strings.Index(strings.ToUpper(rest[len(startKeyword):]), endKeyword)
		if eIdx > 0 {
			return rest[:len(startKeyword)+eIdx]
		}
	}
	return rest
}

// Regex patterns for SQL scanning
var (
	dropTableRegex   = regexp.MustCompile(`(?i)DROP\s+TABLE\s+(?:IF\s+EXISTS\s+)?([\[\]\w\.]+)`)
	alterTableRegex  = regexp.MustCompile(`(?i)ALTER\s+TABLE\s+([\[\]\w\.]+)\s+(?:WITH\s+(?:NO)?CHECK\s+)?(ADD\s+CONSTRAINT|ADD|DROP\s+COLUMN|DROP\s+CONSTRAINT|ALTER\s+COLUMN)\s+([^;]+)`)
	spRenameRegex    = regexp.MustCompile(`(?i)(?:EXEC|EXECUTE)\s+sp_rename\s+['"]([^'"]+)['"]\s*,\s*['"]([^'"]+)['"](?:\s*,\s*['"]([^'"]+)['"])?`)
	createIndexRegex = regexp.MustCompile(`(?i)CREATE\s+(UNIQUE\s+)?(CLUSTERED\s+|NONCLUSTERED\s+)?INDEX\s+([\[\]\w]+)\s+ON\s+([\[\]\w\.]+)\s*\(([^)]+)\)`)
	dropIndexRegex   = regexp.MustCompile(`(?i)DROP\s+INDEX\s+(?:IF\s+EXISTS\s+)?([\[\]\w]+)\s+ON\s+([\[\]\w\.]+)`)
	includeRegex     = regexp.MustCompile(`(?i)INCLUDE\s*\(([^)]+)\)`)
	whereFilterRegex = regexp.MustCompile(`(?i)WHERE\s+([^;]+)`)
	progHeaderRegex  = regexp.MustCompile(`(?i)(?:CREATE\s+OR\s+ALTER|CREATE|ALTER)\s+(?:PROCEDURE|PROC|VIEW|FUNCTION)\s+([\[\]\w\.]+)`)
	createTypeRegex  = regexp.MustCompile(`(?i)CREATE\s+TYPE\s+([\[\]\w\.]+)`)
	dropTypeRegex    = regexp.MustCompile(`(?i)DROP\s+TYPE\s+(?:IF\s+EXISTS\s+)?([\[\]\w\.]+)`)
)

// IsSemanticNoiseOnly returns true if oldSQL and newSQL have identical normalized tokens.
func IsSemanticNoiseOnly(oldSQL, newSQL string) bool {
	return ComputeNormalizedHash(oldSQL) == ComputeNormalizedHash(newSQL)
}

// CheckIdenticalConstraintDef returns true if two constraint definitions are semantically identical.
func CheckIdenticalConstraintDef(def1, def2 string) bool {
	return ComputeNormalizedHash(def1) == ComputeNormalizedHash(def2)
}
