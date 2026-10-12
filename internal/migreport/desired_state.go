package migreport

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"rza_migrate_db/internal/sqlparse"
)

// DesiredStateDocs is a conservative documentation view of db_object. It never
// infers a table relationship merely from equal column names.
type DesiredStateDocs struct {
	Tables             []DocumentedTable
	Relationships      []ObservedRelationship
	ModuleDependencies []ModuleDependency
	DynamicSQLModules  []string
	Modules            []DocumentedModule
	ProcedureProfiles  []ProcedureProfile
}

type DocumentedTable struct {
	Name        string
	Keys        []DocumentedKey
	ForeignKeys []DocumentedForeignKey
	Consumers   []ModuleConsumer
	Columns     []DocumentedColumn
	Indexes     []DocumentedIndex
}

type DocumentedColumn struct {
	Name     string
	DataType string
	Nullable bool
	Identity bool
}

type DocumentedIndex struct {
	Name     string
	Columns  []string
	Unique   bool
	Filtered bool
}

type DocumentedKey struct {
	Columns  []string
	Evidence string
	Name     string
	Filtered bool
}

type DocumentedForeignKey struct {
	Columns        []string
	TargetTable    string
	TargetColumns  []string
	ConstraintName string
}

type ModuleConsumer struct {
	Name string
	Kind string
}

type ObservedRelationship struct {
	LeftObject   string
	LeftColumns  []string
	RightObject  string
	RightColumns []string
	Evidence     string
	JoinTypes    []string
	Modules      []string
	Occurrences  int
	Confidence   string
}

type ModuleDependency struct {
	Module string
	Kind   string
	Object string
}

type DocumentedModule struct {
	Name string
	Kind string
}

// BuildDesiredStateDocs reads only desired-state source files. Missing
// db_object is valid for a history-only repository and produces empty docs.
func BuildDesiredStateDocs(repoRoot string) (DesiredStateDocs, error) {
	root := filepath.Join(repoRoot, "db_object")
	if info, err := os.Stat(root); err != nil || !info.IsDir() {
		if os.IsNotExist(err) {
			return DesiredStateDocs{}, nil
		}
		return DesiredStateDocs{}, fmt.Errorf("read db_object: %w", err)
	}
	tables := map[string]*DocumentedTable{}
	if err := scanSQLFiles(filepath.Join(root, "tables"), func(path, sql string) {
		name := objectNameFromPath(path)
		if name == "" {
			return
		}
		t := &DocumentedTable{Name: name}
		parseTableDefinition(sql, t)
		tables[name] = t
	}); err != nil {
		return DesiredStateDocs{}, err
	}
	if err := scanSQLFiles(filepath.Join(root, "indexes"), func(_ string, sql string) {
		parseIndexes(sql, tables)
	}); err != nil {
		return DesiredStateDocs{}, err
	}
	views := objectNamesInDirectory(filepath.Join(root, "views"))
	functions := objectNamesInDirectory(filepath.Join(root, "functions"))
	procedures := objectNamesInDirectory(filepath.Join(root, "procedures"))
	modules := make([]moduleSource, 0)
	if err := scanModules(filepath.Join(root, "views"), "view", tables, &modules); err != nil {
		return DesiredStateDocs{}, err
	}
	if err := scanModules(filepath.Join(root, "procedures"), "procedure", tables, &modules); err != nil {
		return DesiredStateDocs{}, err
	}
	relationships, dependencies, dynamicModules := analyzeModuleRelationships(modules, tables, views)
	out := DesiredStateDocs{Tables: make([]DocumentedTable, 0, len(tables)), Relationships: relationships, ModuleDependencies: dependencies, DynamicSQLModules: dynamicModules}
	for _, module := range modules {
		out.Modules = append(out.Modules, DocumentedModule{Name: module.Name, Kind: module.Kind})
	}
	sort.Slice(out.Modules, func(i, j int) bool {
		return out.Modules[i].Kind+out.Modules[i].Name < out.Modules[j].Kind+out.Modules[j].Name
	})
	for _, table := range tables {
		sortDocumentedTable(table)
		out.Tables = append(out.Tables, *table)
	}
	sort.Slice(out.Tables, func(i, j int) bool { return out.Tables[i].Name < out.Tables[j].Name })
	out.ProcedureProfiles = buildProcedureProfiles(modules, tables, views, functions, procedures)
	return out, nil
}

func scanSQLFiles(dir string, consume func(string, string)) error {
	entries, err := os.ReadDir(dir)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() || !strings.EqualFold(filepath.Ext(entry.Name()), ".sql") {
			continue
		}
		path := filepath.Join(dir, entry.Name())
		data, readErr := os.ReadFile(path)
		if readErr != nil {
			return readErr
		}
		consume(path, string(data))
	}
	return nil
}

type moduleSource struct{ Name, Kind, SQL string }

func scanModules(dir, kind string, tables map[string]*DocumentedTable, modules *[]moduleSource) error {
	return scanSQLFiles(dir, func(path, sql string) {
		module := objectNameFromPath(path)
		*modules = append(*modules, moduleSource{Name: module, Kind: kind, SQL: sql})
		for _, ref := range sqlparse.ExtractObjectReferences(sql) {
			if ref.Kind != sqlparse.ObjectReferenceData {
				continue
			}
			if table, ok := lookupDocumentedTable(tables, ref.Name); ok {
				table.Consumers = append(table.Consumers, ModuleConsumer{Name: module, Kind: kind})
			}
		}
	})
}

func lookupDocumentedTable(tables map[string]*DocumentedTable, name string) (*DocumentedTable, bool) {
	if table, ok := tables[name]; ok {
		return table, true
	}
	for candidate, table := range tables {
		if strings.EqualFold(candidate, name) {
			return table, true
		}
	}
	return nil, false
}

func objectNamesInDirectory(dir string) map[string]struct{} {
	out := map[string]struct{}{}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return out
	}
	for _, entry := range entries {
		if !entry.IsDir() && strings.EqualFold(filepath.Ext(entry.Name()), ".sql") {
			out[strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))] = struct{}{}
		}
	}
	return out
}

func objectNameFromPath(path string) string {
	base := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	if !strings.Contains(base, ".") {
		return ""
	}
	return base
}

func parseTableDefinition(sql string, table *DocumentedTable) {
	upper := strings.ToUpper(sql)
	create := strings.Index(upper, "CREATE TABLE")
	if create < 0 {
		return
	}
	open := strings.Index(sql[create:], "(")
	if open < 0 {
		return
	}
	open += create
	close, ok := sqlparse.FindMatchingParen(sql, open)
	if !ok {
		return
	}
	for _, item := range sqlparse.SplitTopLevelComma(sql[open+1 : close]) {
		clean := strings.TrimSpace(sqlparse.StripSQLStringsAndComments(item))
		itemUpper := strings.ToUpper(clean)
		var column DocumentedColumn
		if !strings.HasPrefix(itemUpper, "CONSTRAINT") && !strings.HasPrefix(itemUpper, "PRIMARY") && !strings.HasPrefix(itemUpper, "UNIQUE") && !strings.HasPrefix(itemUpper, "FOREIGN") && !strings.HasPrefix(itemUpper, "CHECK") {
			column = parseDocumentedColumn(clean)
			if column.Name != "" {
				table.Columns = append(table.Columns, column)
			}
		}
		if containsSQLWord(itemUpper, "FOREIGN") && containsSQLWord(itemUpper, "KEY") {
			table.ForeignKeys = append(table.ForeignKeys, parseForeignKey(clean)...)
			continue
		}
		if containsSQLWord(itemUpper, "PRIMARY") && containsSQLWord(itemUpper, "KEY") {
			columns := keyColumns(clean)
			if column.Name != "" {
				columns = []string{column.Name}
			}
			table.Keys = append(table.Keys, DocumentedKey{Columns: columns, Evidence: "PRIMARY_KEY", Name: constraintName(clean)})
			continue
		}
		if containsSQLWord(itemUpper, "UNIQUE") && (column.Name != "" || strings.HasPrefix(itemUpper, "CONSTRAINT") || strings.HasPrefix(itemUpper, "UNIQUE")) {
			columns := keyColumns(clean)
			if column.Name != "" {
				columns = []string{column.Name}
			}
			table.Keys = append(table.Keys, DocumentedKey{Columns: columns, Evidence: "UNIQUE_CONSTRAINT", Name: constraintName(clean)})
			continue
		}
	}
}

func parseDocumentedColumn(item string) DocumentedColumn {
	name, consumed := readIdentifier(item)
	if name == "" {
		return DocumentedColumn{}
	}
	rest := strings.TrimSpace(item[consumed:])
	typeName := firstIdentifier(rest)
	upper := strings.ToUpper(rest)
	return DocumentedColumn{Name: name, DataType: strings.ToUpper(typeName), Nullable: !containsSQLWords(upper, "NOT", "NULL"), Identity: hasIdentityDeclaration(upper)}
}

// hasIdentityDeclaration accepts only the SQL IDENTITY keyword followed by its
// seed/increment parentheses. Identifiers such as M_IDENTITY are columns, not
// identity declarations.
func hasIdentityDeclaration(upper string) bool {
	for offset := 0; ; {
		pos := strings.Index(upper[offset:], "IDENTITY")
		if pos < 0 {
			return false
		}
		pos += offset
		beforeOK := pos == 0 || !isIdentifierByte(upper[pos-1])
		after := pos + len("IDENTITY")
		for after < len(upper) && (upper[after] == ' ' || upper[after] == '\t' || upper[after] == '\r' || upper[after] == '\n') {
			after++
		}
		if beforeOK && after < len(upper) && upper[after] == '(' {
			return true
		}
		offset = pos + len("IDENTITY")
	}
}

func isIdentifierByte(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || ch == '_' || ch == '#'
}

func parseIndexes(sql string, tables map[string]*DocumentedTable) {
	clean := sqlparse.StripSQLStringsAndComments(sql)
	upper := strings.ToUpper(clean)
	pos := 0
	for {
		next := strings.Index(upper[pos:], "CREATE ")
		if next < 0 {
			return
		}
		start := pos + next
		indexPos := strings.Index(upper[start:], "INDEX")
		onPos := strings.Index(upper[start:], " ON ")
		if indexPos < 0 || onPos < 0 || indexPos > onPos {
			pos = start + len("CREATE ")
			continue
		}
		indexPos += start
		onPos += start
		indexName := firstIdentifier(clean[indexPos+len("INDEX"):])
		target := strings.TrimSpace(clean[onPos+len(" ON "):])
		tableName := qualifiedIdentifier(target)
		open := strings.Index(target, "(")
		if table, ok := lookupDocumentedTable(tables, tableName); ok && open >= 0 {
			if close, matched := sqlparse.FindMatchingParen(target, open); matched {
				columns := identifiersInList(target[open+1 : close])
				unique := containsSQLWord(upper[start:indexPos], "UNIQUE")
				statementTail := strings.ToUpper(target[close+1:])
				if semi := strings.Index(statementTail, ";"); semi >= 0 {
					statementTail = statementTail[:semi]
				}
				filtered := containsSQLWord(statementTail, "WHERE")
				table.Indexes = append(table.Indexes, DocumentedIndex{Name: indexName, Columns: columns, Unique: unique, Filtered: filtered})
				if unique {
					evidence := "UNIQUE_INDEX"
					if filtered {
						evidence = "FILTERED_UNIQUE_INDEX"
					}
					table.Keys = append(table.Keys, DocumentedKey{Columns: columns, Evidence: evidence, Name: indexName, Filtered: filtered})
				}
			}
		}
		pos = onPos + len(" ON ")
	}
}

func containsSQLWord(upper, word string) bool {
	for offset := 0; ; {
		pos := strings.Index(upper[offset:], word)
		if pos < 0 {
			return false
		}
		pos += offset
		before := pos == 0 || !isIdentifierByte(upper[pos-1])
		afterPos := pos + len(word)
		after := afterPos >= len(upper) || !isIdentifierByte(upper[afterPos])
		if before && after {
			return true
		}
		offset = afterPos
	}
}

func containsSQLWords(upper string, words ...string) bool {
	pos := 0
	for _, word := range words {
		rel := strings.Index(upper[pos:], word)
		if rel < 0 {
			return false
		}
		pos += rel + len(word)
	}
	return true
}

func parseForeignKey(item string) []DocumentedForeignKey {
	upper := strings.ToUpper(item)
	fk := strings.Index(upper, "FOREIGN KEY")
	ref := strings.Index(upper, "REFERENCES")
	if fk < 0 || ref < 0 {
		return nil
	}
	local := parenListAfter(item, fk+len("FOREIGN KEY"))
	targetPart := strings.TrimSpace(item[ref+len("REFERENCES"):])
	target := qualifiedIdentifier(targetPart)
	open := strings.Index(targetPart, "(")
	var targetColumns []string
	if open >= 0 {
		if close, ok := sqlparse.FindMatchingParen(targetPart, open); ok {
			targetColumns = identifiersInList(targetPart[open+1 : close])
		}
	}
	if target == "" {
		return nil
	}
	return []DocumentedForeignKey{{Columns: identifiersInList(local), TargetTable: target, TargetColumns: targetColumns, ConstraintName: constraintName(item)}}
}

func keyColumns(item string) []string {
	upper := strings.ToUpper(item)
	pos := strings.Index(upper, "KEY")
	if pos < 0 {
		pos = strings.Index(upper, "UNIQUE") + len("UNIQUE")
	} else {
		pos += len("KEY")
	}
	return identifiersInList(parenListAfter(item, pos))
}

func parenListAfter(s string, start int) string {
	open := strings.Index(s[start:], "(")
	if open < 0 {
		return ""
	}
	open += start
	close, ok := sqlparse.FindMatchingParen(s, open)
	if !ok {
		return ""
	}
	return s[open+1 : close]
}

func constraintName(s string) string {
	upper := strings.ToUpper(strings.TrimSpace(s))
	if !strings.HasPrefix(upper, "CONSTRAINT") {
		return ""
	}
	return firstIdentifier(strings.TrimSpace(s[len("CONSTRAINT"):]))
}

func qualifiedIdentifier(s string) string {
	first, consumed := readIdentifier(s)
	if first == "" {
		return ""
	}
	rest := strings.TrimSpace(s[consumed:])
	if !strings.HasPrefix(rest, ".") {
		return "dbo." + first
	}
	second, _ := readIdentifier(strings.TrimSpace(rest[1:]))
	if second == "" {
		return ""
	}
	return first + "." + second
}

func firstIdentifier(s string) string {
	name, _ := readIdentifier(s)
	return name
}

func readIdentifier(s string) (string, int) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", 0
	}
	if s[0] == '[' {
		if end := strings.Index(s, "]"); end > 0 {
			return strings.ReplaceAll(s[1:end], "]]", "]"), end + 1
		}
		return "", 0
	}
	end := 0
	for end < len(s) && ((s[end] >= 'a' && s[end] <= 'z') || (s[end] >= 'A' && s[end] <= 'Z') || (s[end] >= '0' && s[end] <= '9') || s[end] == '_' || s[end] == '#') {
		end++
	}
	return s[:end], end
}

func identifiersInList(s string) []string {
	var out []string
	for _, part := range sqlparse.SplitTopLevelComma(s) {
		if name := firstIdentifier(strings.TrimSpace(part)); name != "" {
			out = append(out, name)
		}
	}
	return out
}

func sortDocumentedTable(t *DocumentedTable) {
	sort.Slice(t.Keys, func(i, j int) bool {
		return t.Keys[i].Evidence+t.Keys[i].Name+strings.Join(t.Keys[i].Columns, ",") < t.Keys[j].Evidence+t.Keys[j].Name+strings.Join(t.Keys[j].Columns, ",")
	})
	sort.Slice(t.ForeignKeys, func(i, j int) bool {
		return t.ForeignKeys[i].TargetTable+t.ForeignKeys[i].ConstraintName < t.ForeignKeys[j].TargetTable+t.ForeignKeys[j].ConstraintName
	})
	sort.Slice(t.Consumers, func(i, j int) bool {
		return t.Consumers[i].Kind+t.Consumers[i].Name < t.Consumers[j].Kind+t.Consumers[j].Name
	})
	sort.Slice(t.Columns, func(i, j int) bool { return t.Columns[i].Name < t.Columns[j].Name })
	sort.Slice(t.Indexes, func(i, j int) bool { return t.Indexes[i].Name < t.Indexes[j].Name })
}
