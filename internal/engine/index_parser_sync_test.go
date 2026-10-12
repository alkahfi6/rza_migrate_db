package engine

import "strings"
import "testing"

func TestExtractCreateIndexStatementsSupportsClusteredVariants(t *testing.T) {
	t.Parallel()

	sql := `CREATE NONCLUSTERED INDEX [IX_TableA__status] ON dbo.TableA ([status] ASC);
CREATE UNIQUE CLUSTERED INDEX [IX_TableA__id] ON [dbo].[TableA] ([id] ASC);
`

	stmts, err := extractCreateIndexStatements("dbo.TableA", sql)
	if err != nil {
		t.Fatalf("extractCreateIndexStatements failed: %v", err)
	}
	if len(stmts) != 2 {
		t.Fatalf("expected 2 index statements, got %d: %#v", len(stmts), stmts)
	}
}

func TestExtractCreateIndexStatementsSupportsOnWithMultipleSpaces(t *testing.T) {
	t.Parallel()

	sql := `CREATE NONCLUSTERED INDEX [IX_TableA__status] ON     dbo.TableA ([status] ASC);`

	stmts, err := extractCreateIndexStatements("dbo.TableA", sql)
	if err != nil {
		t.Fatalf("extractCreateIndexStatements failed: %v", err)
	}
	if len(stmts) != 1 {
		t.Fatalf("expected 1 index statement, got %d: %#v", len(stmts), stmts)
	}
}

func TestExtractCreateIndexStatementsAcceptsUnqualifiedDboTarget(t *testing.T) {
	t.Parallel()

	sql := `CREATE CLUSTERED INDEX IX_TableA__id ON TableA (id);`
	stmts, err := extractCreateIndexStatements("dbo.TableA", sql)
	if err != nil {
		t.Fatalf("extractCreateIndexStatements failed: %v", err)
	}
	if len(stmts) != 1 {
		t.Fatalf("expected 1 index statement, got %d: %#v", len(stmts), stmts)
	}
	if err := validateIndexSourceTargets("dbo.TableA", sql); err != nil {
		t.Fatalf("validateIndexSourceTargets rejected unqualified dbo target: %v", err)
	}
	state, err := parseIndexStateMap("dbo.TableA", sql)
	if err != nil {
		t.Fatalf("parseIndexStateMap failed: %v", err)
	}
	if !strings.Contains(state["IX_TableA__id"].Statement, "ON dbo.TableA") {
		t.Fatalf("expected generated index state to qualify dbo target: %#v", state)
	}
}

func TestExtractCreateIndexStatementsSupportsAdvancedFamilies(t *testing.T) {
	sqlText := `
CREATE PRIMARY XML INDEX [PXML_T_Doc] ON dbo.T ([Doc]);
CREATE XML INDEX [SXML_T_Doc_Path] ON dbo.T ([Doc]) USING XML INDEX [PXML_T_Doc] FOR PATH;
CREATE SPATIAL INDEX [SIX_T_Geo] ON dbo.T ([Geo]) USING GEOMETRY_AUTO_GRID;
CREATE NONCLUSTERED COLUMNSTORE INDEX [NCCI_T] ON dbo.T ([Amount]);
CREATE NONCLUSTERED INDEX [IX_T_Amount] ON dbo.T ([Amount]) WITH (DATA_COMPRESSION = PAGE);`
	statements, err := extractCreateIndexStatements("dbo.T", sqlText)
	if err != nil {
		t.Fatal(err)
	}
	if len(statements) != 5 {
		t.Fatalf("expected five advanced index statements, got %d: %#v", len(statements), statements)
	}
	for _, statement := range statements {
		if _, err := parseCreateIndexName(statement); err != nil {
			t.Fatalf("advanced index name parse failed for %q: %v", statement, err)
		}
	}
}

func TestXMLIndexLifecycleOrdersDependenciesForUpAndDown(t *testing.T) {
	state := `
CREATE XML INDEX [A_Secondary] ON dbo.T ([Doc]) USING XML INDEX [Z_Primary] FOR PATH;
CREATE PRIMARY XML INDEX [Z_Primary] ON dbo.T ([Doc]);`
	up, down, err := buildIndexLifecycleSQL("", state, "dbo.T")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Index(strings.ToUpper(up), "Z_PRIMARY") > strings.Index(strings.ToUpper(up), "A_SECONDARY") {
		t.Fatalf("primary XML index must be created first:\n%s", up)
	}
	if strings.Index(strings.ToUpper(down), "A_SECONDARY") > strings.Index(strings.ToUpper(down), "Z_PRIMARY") {
		t.Fatalf("secondary XML index must be dropped first:\n%s", down)
	}
}

func TestSemanticFromSourceIndexesSupportsFilteredIndex(t *testing.T) {
	t.Parallel()

	meta := migrationMeta{
		Category: "indexes",
		Object:   "dbo.TableA",
	}
	sql := "CREATE NONCLUSTERED INDEX [IX_TableA__status_active] ON dbo.TableA ([status] ASC) WHERE [status] = 'ACTIVE';\n"

	semantic, err := semanticFromSource(meta, sql)
	if err != nil {
		t.Fatalf("semanticFromSource failed: %v", err)
	}
	if semantic == "" {
		t.Fatalf("expected non-empty semantic SQL")
	}
	if !strings.Contains(strings.ToUpper(semantic), " WHERE ") {
		t.Fatalf("expected filtered predicate to be preserved in semantic SQL: %s", semantic)
	}
}

func TestExtractProgrammabilityStatementSupportsBracketedObject(t *testing.T) {
	t.Parallel()

	sql := "CREATE OR ALTER VIEW [dbo].[ViewA] AS SELECT 1 AS id;\n"
	stmt, err := extractProgrammabilityStatement("views", "dbo.ViewA", sql)
	if err != nil {
		t.Fatalf("extractProgrammabilityStatement failed: %v", err)
	}
	if stmt == "" {
		t.Fatalf("expected non-empty statement")
	}
}

func TestExtractProgrammabilityStatementSupportsFlexibleHeaderWhitespace(t *testing.T) {
	t.Parallel()

	sql := "CREATE OR ALTER   PROCEDURE [dbo].[ProcA]\nAS\nBEGIN\n  SELECT 1;\nEND\n"
	stmt, err := extractProgrammabilityStatement("procedures", "dbo.ProcA", sql)
	if err != nil {
		t.Fatalf("extractProgrammabilityStatement failed: %v", err)
	}
	if !strings.Contains(stmt, "PROCEDURE [dbo].[ProcA]") {
		t.Fatalf("expected procedure statement, got: %s", stmt)
	}
}
