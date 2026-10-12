package migreport

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func precisionTables() map[string]*DocumentedTable {
	out := map[string]*DocumentedTable{}
	for _, name := range []string{"dbo.A", "dbo.B", "dbo.C"} {
		out[name] = &DocumentedTable{Name: name, Columns: []DocumentedColumn{{Name: "Id"}, {Name: "Code"}}}
	}
	return out
}

func TestObservedJoinPrecision(t *testing.T) {
	cases := []struct {
		name, sql string
		want      int
		join      string
	}{
		{"outer join", `SELECT * FROM dbo.A a LEFT OUTER JOIN dbo.B b ON a.Id=b.Id;`, 1, "LEFT"},
		{"merge join hint", `SELECT * FROM dbo.A a LEFT OUTER MERGE JOIN dbo.B b ON a.Id=b.Id;`, 1, "LEFT"},
		{"hash join hint", `SELECT * FROM dbo.A a FULL HASH JOIN dbo.B b ON a.Id=b.Id;`, 1, "FULL"},
		{"outer aliases survive subquery", `SELECT * FROM dbo.A a JOIN dbo.B b ON a.Id=b.Id AND EXISTS (SELECT 1 FROM dbo.C c WHERE c.Id=a.Id) JOIN dbo.C d ON a.Code=d.Code;`, 2, "INNER"},
		{"nested aliases do not escape", `SELECT * FROM dbo.A a JOIN dbo.B b ON EXISTS (SELECT 1 FROM dbo.C a WHERE a.Id=b.Id);`, 0, ""},
		{"CTE shadows physical table", `WITH A AS (SELECT Id FROM dbo.C) SELECT * FROM A a JOIN dbo.B b ON a.Id=b.Id;`, 0, ""},
		{"duplicate alias", `SELECT * FROM dbo.A x JOIN dbo.B x ON x.Id=c.Id JOIN dbo.C c ON x.Code=c.Code;`, 0, ""},
		{"quoted keyword alias", `SELECT * FROM dbo.A AS [WHERE] JOIN dbo.B AS [JOIN] ON [WHERE].Id=[JOIN].Id;`, 1, "INNER"},
		{"multi argument wrapper", `SELECT * FROM dbo.A a JOIN dbo.B b ON ISNULL(a.Code,'')=b.Code;`, 1, "INNER"},
		{"schema qualified unaliased", `SELECT * FROM dbo.A JOIN dbo.B ON dbo.A.Id=dbo.B.Id;`, 1, "INNER"},
		{"statement leakage", `SELECT * FROM dbo.A a, dbo.B b WHERE a.Id=b.Id SELECT * FROM dbo.C c WHERE a.Code=b.Code;`, 1, "IMPLICIT"},
		{"four part external", `SELECT * FROM srv.db.dbo.A a JOIN dbo.B b ON a.Id=b.Id;`, 1, "INNER"},
		{"schema omitted external", `SELECT * FROM db..A a JOIN dbo.B b ON a.Id=b.Id;`, 1, "INNER"},
		{"case canonicalization", `SELECT * FROM dbo.a a JOIN dbo.b b ON a.id=b.id;`, 1, "INNER"},
		{"hint and multiline", "SELECT * FROM dbo.A a WITH (NOLOCK) LEFT\nOUTER JOIN\ndbo.B b WITH(NOLOCK) ON a.Id=b.Id;", 1, "LEFT"},
		{"nested column wrapper", `SELECT * FROM dbo.A a JOIN dbo.B b ON dbo.Normalise(ISNULL(a.Code,''))=b.Code;`, 1, "INNER"},
		{"ambiguous expression", `SELECT * FROM dbo.A a JOIN dbo.B b ON a.Code+a.Id=b.Code;`, 0, ""},
		{"temporary source", `SELECT * FROM #A a JOIN dbo.B b ON a.Id=b.Id;`, 0, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := extractObservedJoins(moduleSource{Name: "dbo.p", Kind: "procedure", SQL: tc.sql}, precisionTables(), nil)
			if len(got) != tc.want {
				t.Fatalf("got %d relations, want %d: %+v", len(got), tc.want, got)
			}
			if tc.join != "" && len(got) > 0 && got[0].JoinTypes[0] != tc.join {
				t.Fatalf("join type: %+v", got)
			}
		})
	}
}

func TestProcedureParameterKeywordsInsideLiteralAndOptions(t *testing.T) {
	params := parseProcedureParameters(`CREATE OR ALTER PROC dbo.p @message VARCHAR(30)='OUTPUT', @status INT=0 OUTPUT WITH EXECUTE AS OWNER AS SELECT 1;`)
	if len(params) != 2 || params[0].Default != "'OUTPUT'" || params[0].Direction != "INPUT" || params[1].Default != "0" || params[1].Direction != "INPUT/OUTPUT" || params[1].SQLType != "INT" {
		t.Fatalf("literal/option keywords corrupted parameters: %+v", params)
	}
}

func TestWriteMappingUnknownColumnIsEvidenceNotVerified(t *testing.T) {
	p := analyzeProcedure(moduleSource{Name: "dbo.p", Kind: "procedure", SQL: `UPDATE dbo.A SET MissingColumn=1, [Code]='OUTPUT';`}, newProcedureCatalog(precisionTables(), nil, nil, nil))
	if len(p.Mappings) != 2 || p.Mappings[0].Resolution != "COLUMN_NOT_FOUND" || p.Mappings[1].Resolution != ResolutionStatic || len(p.ExtractionNotices) == 0 {
		t.Fatalf("unknown column must be marked: %+v", p)
	}
}

func TestDocumentedInlineAndAlterForeignKeys(t *testing.T) {
	table := &DocumentedTable{Name: "dbo.A"}
	parseTableDefinition(`CREATE TABLE dbo.A ([Id] INT CONSTRAINT [PK_Id] PRIMARY KEY CONSTRAINT [FK_Id] REFERENCES dbo.B(Id), [Code] INT);
ALTER TABLE dbo.A WITH CHECK ADD CONSTRAINT [FK_Code] FOREIGN
KEY (Code) REFERENCES dbo.C(Code);`, table)
	if len(table.Keys) != 1 || table.Keys[0].Name != "PK_Id" || len(table.ForeignKeys) != 2 || table.ForeignKeys[0].ConstraintName != "FK_Id" || table.ForeignKeys[1].TargetTable != "dbo.C" {
		t.Fatalf("inline/ALTER FK metadata: %+v", table)
	}
}

func TestInsertQualifiedTargetDoesNotBecomeWrongColumnMapping(t *testing.T) {
	p := analyzeProcedure(moduleSource{Name: "dbo.p", Kind: "procedure", SQL: `INSERT dbo.A (Id,b.Code) SELECT b.Id,b.Code FROM dbo.B b;`}, newProcedureCatalog(precisionTables(), nil, nil, nil))
	if len(p.Mappings) != 0 || len(p.ExtractionNotices) == 0 || !strings.Contains(strings.Join(p.ExtractionNotices, " "), "daftar target column") {
		t.Fatalf("qualified target was guessed: %+v", p)
	}
}

func TestTableFilenameMismatchIsActionableAndReadOnly(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "db_object", "tables")
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, "dbo.A.sql")
	sql := []byte(`CREATE TABLE dbo.B (Id INT);`)
	if err := os.WriteFile(path, sql, 0644); err != nil {
		t.Fatal(err)
	}
	_, err := BuildDesiredStateDocs(root)
	if err == nil || !strings.Contains(err.Error(), "identity mismatch") {
		t.Fatalf("misidentified table accepted: %v", err)
	}
	after, _ := os.ReadFile(path)
	if string(after) != string(sql) {
		t.Fatal("source changed")
	}
}

func TestModuleFilenameMismatchIsNotMappedToWrongProcedure(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "db_object", "procedures")
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "dbo.A.sql"), []byte(`CREATE OR ALTER PROC dbo.B AS SELECT 1;`), 0644); err != nil {
		t.Fatal(err)
	}
	_, err := BuildDesiredStateDocs(root)
	if err == nil || !strings.Contains(err.Error(), "source identity mismatch") {
		t.Fatalf("procedure identity was guessed: %v", err)
	}
}

func TestDynamicLiteralMustBeExecutedAndComplete(t *testing.T) {
	catalog := newProcedureCatalog(precisionTables(), nil, nil, nil)
	for _, sql := range []string{`DECLARE @sample NVARCHAR(MAX)=N'EXEC(N''SELECT * FROM dbo.A'')'; SELECT 1;`, `SELECT 'sp_executesql ''SELECT * FROM dbo.A''';`} {
		p := analyzeProcedure(moduleSource{Name: "dbo.p", Kind: "procedure", SQL: sql}, catalog)
		if len(p.DynamicSQL) != 0 {
			t.Fatalf("text inside literal became executable evidence: %+v", p.DynamicSQL)
		}
	}
	p := analyzeProcedure(moduleSource{Name: "dbo.p", Kind: "procedure", SQL: `EXEC sp_executesql N'SELECT * FROM dbo.A ' + @suffix;`}, catalog)
	if len(p.DynamicSQL) != 1 || p.DynamicSQL[0].Resolution != ResolutionUnresolved {
		t.Fatalf("partial runtime SQL claimed resolved: %+v", p.DynamicSQL)
	}
	p = analyzeProcedure(moduleSource{Name: "dbo.p", Kind: "procedure", SQL: "EXEC\n sys.sp_executesql\n N'SELECT * FROM dbo.A';"}, catalog)
	if len(p.DynamicSQL) != 1 || p.DynamicSQL[0].Resolution != ResolutionDynamic {
		t.Fatalf("multiline static literal not resolved: %+v", p.DynamicSQL)
	}
	if _, ok := completeDynamicLiteral(`N'SELECT ''`); ok {
		t.Fatal("unterminated literal claimed complete")
	}
	quotedTables := map[string]*DocumentedTable{"dbo.A.B": {Name: "dbo.A.B", Columns: []DocumentedColumn{{Name: "Id"}}}}
	if got, status := newProcedureCatalog(quotedTables, nil, nil, nil).resolveParts("A.B", []string{"A.B"}, "data"); got != "dbo.A.B" || status != ResolutionStatic {
		t.Fatalf("quoted dot lost during unqualified lookup: %s/%s", got, status)
	}
}

func TestDocumentedDDLPrecision(t *testing.T) {
	table := &DocumentedTable{Name: "dbo.A"}
	parseTableDefinition(`/* CREATE TABLE Fake (Wrong INT) */
 CREATE
 TABLE dbo.A ([Id] INT IDENTITY PRIMARY KEY, [Code] VARCHAR(30), [NotNullValue] INT NULL, [Unique] INT, CONSTRAINT [KEY_UQ] UNIQUE ([Code]));`, table)
	if len(table.Columns) != 4 || len(table.Keys) != 2 {
		t.Fatalf("columns/keys: %+v", table)
	}
	for _, c := range table.Columns {
		if c.Name == "Id" && (!c.Identity || c.Nullable) {
			t.Errorf("inline PK/identity: %+v", c)
		}
		if c.Name == "Code" && c.DataType != "VARCHAR(30)" {
			t.Errorf("type precision: %+v", c)
		}
		if c.Name == "NotNullValue" && !c.Nullable {
			t.Errorf("quoted name confused with nullability: %+v", c)
		}
	}
	parseIndexes("CREATE UNIQUE INDEX UX_A ON\n dbo.A (Code)\nGO\nCREATE INDEX IX_A ON dbo.A (Id) WHERE Id>0;", map[string]*DocumentedTable{"dbo.A": table})
	if len(table.Indexes) != 2 || table.Indexes[0].Filtered || !table.Indexes[1].Filtered {
		t.Fatalf("index statement boundaries: %+v", table.Indexes)
	}
}

func TestDocumentedTableLookupRejectsCaseCollision(t *testing.T) {
	tables := map[string]*DocumentedTable{"dbo.A": {Name: "dbo.A"}, "dbo.a": {Name: "dbo.a"}}
	if _, ok := lookupDocumentedTable(tables, "dbo.A"); ok {
		t.Fatal("colliding case identities must not be arbitrarily resolved")
	}
}

func TestUpdateMappingWithoutWhere(t *testing.T) {
	tables := precisionTables()
	catalog := newProcedureCatalog(tables, nil, nil, nil)
	p := analyzeProcedure(moduleSource{Name: "dbo.p", Kind: "procedure", SQL: `CREATE PROCEDURE dbo.p AS UPDATE dbo.A SET Code='ready'; UPDATE a SET Code=b.Code FROM dbo.A a JOIN dbo.B b ON a.Id=b.Id;`}, catalog)
	if len(p.Mappings) != 2 {
		t.Fatalf("unfiltered writes must retain mappings: %+v", p.Mappings)
	}
	for _, m := range p.Mappings {
		if strings.Contains(m.Expression, "FROM") || m.Filter != "" {
			t.Fatalf("statement leakage: %+v", m)
		}
	}
}

func TestReportReadsUnicodeAndFunctionModulesWithoutWriting(t *testing.T) {
	root := t.TempDir()
	for _, dir := range []string{"tables", "functions"} {
		if err := os.MkdirAll(filepath.Join(root, "db_object", dir), 0755); err != nil {
			t.Fatal(err)
		}
	}
	path := filepath.Join(root, "db_object", "tables", "dbo.A.sql")
	sql := `CREATE TABLE dbo.A (Id INT PRIMARY KEY);`
	raw := []byte{0xff, 0xfe}
	for _, b := range []byte(sql) {
		raw = append(raw, b, 0)
	}
	if err := os.WriteFile(path, raw, 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "db_object", "functions", "dbo.f.sql"), []byte(`CREATE FUNCTION dbo.f() RETURNS INT AS BEGIN RETURN (SELECT Id FROM dbo.A); END;`), 0644); err != nil {
		t.Fatal(err)
	}
	docs, err := BuildDesiredStateDocs(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(docs.Tables) != 1 || len(docs.Tables[0].Columns) != 1 || len(docs.Modules) != 1 || len(docs.ModuleDependencies) != 1 || len(docs.Tables[0].Consumers) != 1 {
		t.Fatalf("unicode/function coverage: %+v", docs)
	}
	after, _ := os.ReadFile(path)
	if string(after) != string(raw) {
		t.Fatal("report normalized authored bytes on disk")
	}
}

func TestMigrationUnicodeIsDecodedOnlyInMemory(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "migrations")
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	sql := `CREATE TABLE dbo.A (Id INT);`
	raw := []byte{0xff, 0xfe}
	for _, b := range []byte(sql) {
		raw = append(raw, b, 0)
	}
	path := filepath.Join(dir, "20261012000000000_tables_dbo.A_v1.up.sql")
	if err := os.WriteFile(path, raw, 0644); err != nil {
		t.Fatal(err)
	}
	pairs, files, warnings, errors, err := ScanMigrations(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 1 || string(files[0].RawBytes) != sql {
		t.Fatalf("migration was not decoded: %+v", files)
	}
	result, err := Analyze(root, pairs, files, warnings, errors)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.TableChanges) == 0 {
		t.Fatal("unicode DDL silently disappeared")
	}
	after, _ := os.ReadFile(path)
	if string(after) != string(raw) {
		t.Fatal("immutable migration bytes rewritten")
	}
}

func TestReportRejectsUnidentifiableEncoding(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "db_object", "tables")
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "dbo.A.sql"), []byte{0xc0, 0x80}, 0644); err != nil {
		t.Fatal(err)
	}
	_, err := BuildDesiredStateDocs(root)
	if err == nil || !strings.Contains(err.Error(), "encoding") {
		t.Fatalf("unsafe encoding accepted: %v", err)
	}
}

func TestDynamicReturnAssignmentPrecision(t *testing.T) {
	if !moduleUsesDynamicSQL("EXEC @rc = @procedureName") {
		t.Fatal("runtime procedure target must remain dynamic")
	}
	if moduleUsesDynamicSQL("EXEC @rc = dbo.StaticProcedure") {
		t.Fatal("static return assignment must remain static")
	}
}
