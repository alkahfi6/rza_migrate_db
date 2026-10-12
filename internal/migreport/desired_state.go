package migreport

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"rza_migrate_db/internal/sqlfmt"
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
	Module     string
	Kind       string
	Object     string
	Resolution string
}

type DocumentedModule struct {
	Name string
	Kind string
}

// BuildDesiredStateDocs reads only desired-state source files. Missing
// db_object is valid for a history-only repository and produces empty docs.
func BuildDesiredStateDocs(repoRoot string) (DesiredStateDocs, error) {
	root := filepath.Join(repoRoot, "db_object")
	info, statErr := os.Stat(root)
	if statErr != nil {
		if os.IsNotExist(statErr) {
			return DesiredStateDocs{}, nil
		}
		return DesiredStateDocs{}, fmt.Errorf("read db_object: %w", statErr)
	}
	if !info.IsDir() {
		return DesiredStateDocs{}, fmt.Errorf("db_object is not a directory: %s", root)
	}
	tables := map[string]*DocumentedTable{}
	var identityErr error
	if err := scanSQLFiles(filepath.Join(root, "tables"), func(path, sql string) {
		name := objectNameFromPath(path)
		if name == "" {
			identityErr = fmt.Errorf("table filename must identify schema.object: %s", path)
			return
		}
		t := &DocumentedTable{Name: name}
		declared := parseTableDefinition(sql, t)
		if !strings.Contains(declared, ".") && declared != "" {
			declared = "dbo." + declared
		}
		if !strings.EqualFold(declared, name) {
			identityErr = fmt.Errorf("table source identity mismatch in %s: filename=%s CREATE TABLE=%q; correct source identity before reporting", path, name, declared)
			return
		}
		tables[name] = t
	}); err != nil {
		return DesiredStateDocs{}, err
	}
	if identityErr != nil {
		return DesiredStateDocs{}, identityErr
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
	if err := scanModules(filepath.Join(root, "functions"), "function", tables, &modules); err != nil {
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
		decoded, _, decodeErr := sqlfmt.NormalizeSQLToUTF8(data)
		if decodeErr != nil {
			return fmt.Errorf("read SQL %s: %w", path, decodeErr)
		}
		consume(path, string(decoded))
	}
	return nil
}

type moduleSource struct{ Name, Kind, SQL string }

func scanModules(dir, kind string, tables map[string]*DocumentedTable, modules *[]moduleSource) error {
	catalog := newProcedureCatalog(tables, nil, nil, nil)
	var identityErr error
	err := scanSQLFiles(dir, func(path, sql string) {
		module := objectNameFromPath(path)
		declared := sqlparse.ModuleDeclarationName(sql, kind)
		if module == "" || !strings.EqualFold(module, declared) {
			identityErr = fmt.Errorf("%s source identity mismatch in %s: filename=%q declaration=%q; correct source identity before reporting", kind, path, module, declared)
			return
		}
		*modules = append(*modules, moduleSource{Name: module, Kind: kind, SQL: sql})
		seen := map[string]bool{}
		for _, ref := range sqlparse.ExtractSQLAccesses(sql) {
			if ref.NonPersisted || ref.Ambiguous || ref.CTECaseAmbiguous || ref.Function {
				continue
			}
			name, status := catalog.resolveParts(ref.Name, ref.NameParts, "data")
			if status != ResolutionStatic || seen[name] {
				continue
			}
			if table, ok := lookupDocumentedTable(tables, name); ok {
				table.Consumers = append(table.Consumers, ModuleConsumer{Name: module, Kind: kind})
				seen[name] = true
			}
		}
	})
	if err != nil {
		return err
	}
	return identityErr
}

func lookupDocumentedTable(tables map[string]*DocumentedTable, name string) (*DocumentedTable, bool) {
	var match *DocumentedTable
	for candidate, table := range tables {
		if strings.EqualFold(candidate, name) {
			if match != nil {
				return nil, false
			}
			match = table
		}
	}
	return match, match != nil
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

func parseTableDefinition(sql string, table *DocumentedTable) string {
	metadata := sqlparse.ReadTableMetadata(sql)
	for _, c := range metadata.Columns {
		table.Columns = append(table.Columns, DocumentedColumn{Name: c.Name, DataType: c.DataType, Nullable: c.Nullable, Identity: c.Identity})
	}
	for _, k := range metadata.Keys {
		table.Keys = append(table.Keys, DocumentedKey{Name: k.Name, Columns: k.Columns, Evidence: k.Evidence, Filtered: k.Filtered})
	}
	for _, fk := range metadata.ForeignKeys {
		table.ForeignKeys = append(table.ForeignKeys, DocumentedForeignKey{Columns: fk.Columns, TargetTable: fk.Target, TargetColumns: fk.TargetColumns, ConstraintName: fk.Name})
	}
	return metadata.Name
}

func isIdentifierByte(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || ch == '_' || ch == '#'
}

func parseIndexes(sql string, tables map[string]*DocumentedTable) {
	for _, index := range sqlparse.ReadIndexMetadata(sql) {
		table, ok := lookupDocumentedTable(tables, index.Table)
		if !ok {
			continue
		}
		table.Indexes = append(table.Indexes, DocumentedIndex{Name: index.Name, Columns: index.Columns, Unique: index.Unique, Filtered: index.Filtered})
		if index.Unique {
			evidence := "UNIQUE_INDEX"
			if index.Filtered {
				evidence = "FILTERED_UNIQUE_INDEX"
			}
			table.Keys = append(table.Keys, DocumentedKey{Name: index.Name, Columns: index.Columns, Evidence: evidence, Filtered: index.Filtered})
		}
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

func firstIdentifier(s string) string {
	name, _ := readIdentifier(s)
	return name
}

func readIdentifier(s string) (string, int) {
	tokens := sqlparse.SQLTokens(s)
	if len(tokens) == 0 || !tokens[0].Identifier {
		return "", 0
	}
	return tokens[0].Text, tokens[0].End
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
