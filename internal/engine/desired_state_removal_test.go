package engine

import (
	"path/filepath"
	"strings"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/management"
	"rza_migrate_db/internal/output"
)

func TestPlanDesiredStateRemovalCoversCurrentObjectCategories(t *testing.T) {
	cfg := config.Default(t.TempDir())
	engine := New(cfg)
	report := output.NewReport(cfg.OutputNamingPattern, cfg.OutputNamingLineEndings)
	history := map[string]historyObject{
		objectKey("tables", "dbo.EmptyLedger"): {
			Category: "tables", Object: "dbo.EmptyLedger", LastVersion: "20260101000000000", LastVK: 1,
			SemanticSQL: "CREATE TABLE dbo.EmptyLedger ([id] INT NOT NULL);\n",
		},
		objectKey("indexes", "dbo.EmptyLedger"): {
			Category: "indexes", Object: "dbo.EmptyLedger", LastVersion: "20260101000000001", LastVK: 1,
			SemanticSQL: "CREATE INDEX IX_EmptyLedger_id ON dbo.EmptyLedger ([id]);\n",
		},
		objectKey("types", "dbo.IdList"): {
			Category: "types", Object: "dbo.IdList", LastVersion: "20260101000000002", LastVK: 1,
			SemanticSQL: "CREATE TYPE dbo.IdList AS TABLE ([id] INT NOT NULL);\n",
		},
		objectKey("views", "dbo.EmptyLedgerView"): {
			Category: "views", Object: "dbo.EmptyLedgerView", LastVersion: "20260101000000003", LastVK: 1,
			SemanticSQL: "CREATE OR ALTER VIEW dbo.EmptyLedgerView AS SELECT 1 AS id;\n",
		},
		objectKey("functions", "dbo.EmptyLedgerFn"): {
			Category: "functions", Object: "dbo.EmptyLedgerFn", LastVersion: "20260101000000004", LastVK: 1,
			SemanticSQL: "CREATE OR ALTER FUNCTION dbo.EmptyLedgerFn() RETURNS INT AS BEGIN RETURN 1; END;\n",
		},
		objectKey("procedures", "dbo.EmptyLedgerProc"): {
			Category: "procedures", Object: "dbo.EmptyLedgerProc", LastVersion: "20260101000000005", LastVK: 1,
			SemanticSQL: "CREATE OR ALTER PROCEDURE dbo.EmptyLedgerProc AS BEGIN SELECT 1; END;\n",
		},
	}

	ops, _, _, err := engine.planGenerationWithInputs(&report, history, map[string]struct{}{}, "", nil, nil)
	if err != nil {
		t.Fatalf("plan removals: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("unexpected planning failures: %+v", report.Fail.Items)
	}
	if len(ops) != len(history) {
		t.Fatalf("planned operations=%d, want=%d: %#v", len(ops), len(history), ops)
	}

	position := map[string]int{}
	for i, op := range ops {
		position[objectKey(op.Category, op.Object)] = i
		if strings.TrimSpace(op.DownSemanticSQL) == "" {
			t.Fatalf("removal %s:%s has no restore script", op.Category, op.Object)
		}
	}
	if position[objectKey("indexes", "dbo.EmptyLedger")] >= position[objectKey("tables", "dbo.EmptyLedger")] {
		t.Fatal("index removal must be ordered before table removal")
	}

	tableSQL := ops[position[objectKey("tables", "dbo.EmptyLedger")]].UpSemanticSQL
	for _, required := range []string{
		"EXISTS (SELECT 1 FROM dbo.EmptyLedger WITH (TABLOCKX, HOLDLOCK))",
		"FROM sys.foreign_keys",
		"FROM sys.sql_expression_dependencies",
		"own.parent_object_id = OBJECT_ID(N'dbo.EmptyLedger')",
		"FROM sys.sql_modules",
		"dynamic SQL",
		"FROM sys.triggers",
		"DROP TABLE IF EXISTS dbo.EmptyLedger",
	} {
		if !strings.Contains(tableSQL, required) {
			t.Fatalf("guarded DROP TABLE missing %q:\n%s", required, tableSQL)
		}
	}
	if strings.Contains(strings.ToUpper(tableSQL), "RZA PRECHECK") {
		t.Fatalf("generated SQL contains verbose product prefix:\n%s", tableSQL)
	}
	if err := validatePlannedRoundTrip(history, nil, ops); err != nil {
		t.Fatalf("removal round trip: %v", err)
	}
}

func TestDeletedHistoryDoesNotReemitAndCanBeRecreated(t *testing.T) {
	drop := buildDropTableSQL("dbo.Reborn")
	semantic, err := semanticFromMigrationWithState(
		migrationMeta{Category: "tables", Object: "dbo.Reborn"},
		drop,
		"CREATE TABLE dbo.Reborn ([id] INT NOT NULL);\n",
		true,
	)
	if err != nil {
		t.Fatalf("reconstruct removal: %v", err)
	}
	if strings.TrimSpace(semantic) != "" {
		t.Fatalf("drop migration must reconstruct an empty state, got %q", semantic)
	}

	cfg := config.Default(t.TempDir())
	engine := New(cfg)
	report := output.NewReport(cfg.OutputNamingPattern, cfg.OutputNamingLineEndings)
	history := map[string]historyObject{
		objectKey("tables", "dbo.Reborn"): {
			Category: "tables", Object: "dbo.Reborn", LastVersion: "20260101000000000",
			LastVK: 2, Deleted: true,
		},
	}
	ops, _, _, err := engine.planGenerationWithInputs(&report, history, map[string]struct{}{}, "", nil, nil)
	if err != nil || report.Fail.Count != 0 || len(ops) != 0 {
		t.Fatalf("deleted state must be idempotent: ops=%#v fail=%+v err=%v", ops, report.Fail.Items, err)
	}

	create := "CREATE TABLE dbo.Reborn ([id] INT NOT NULL);\n"
	desired := desiredObject{
		Category: "tables", Object: "dbo.Reborn", SemanticSQL: create,
		NormalizedHash: hashSemanticByCategory("tables", create),
	}
	report = output.NewReport(cfg.OutputNamingPattern, cfg.OutputNamingLineEndings)
	ops, _, _, err = engine.planGenerationWithInputs(&report, history, map[string]struct{}{}, "", []desiredObject{desired}, nil)
	if err != nil || report.Fail.Count != 0 || len(ops) != 1 {
		t.Fatalf("recreate after deletion: ops=%#v fail=%+v err=%v", ops, report.Fail.Items, err)
	}
	if ops[0].VK != 3 || !strings.Contains(ops[0].UpSemanticSQL, "CREATE TABLE") {
		t.Fatalf("recreate must continue object version history: %#v", ops[0])
	}
}

func TestRoundTripFailsWhenHistoryOnlyObjectHasNoRemovalOperation(t *testing.T) {
	history := map[string]historyObject{
		objectKey("views", "dbo.Orphan"): {
			Category: "views", Object: "dbo.Orphan",
			SemanticSQL: "CREATE OR ALTER VIEW dbo.Orphan AS SELECT 1 AS id;\n",
		},
	}
	if err := validatePlannedRoundTrip(history, nil, nil); err == nil {
		t.Fatal("unresolved history-only object must fail round-trip coverage")
	}
}

func TestExplicitlyUnmanagedHistoryObjectIsNotDropped(t *testing.T) {
	root := t.TempDir()
	cfg := config.Default(root)
	if err := management.Write(filepath.Join(root, filepath.FromSlash(cfg.ManagementManifestFile)), management.Manifest{
		Schema:    "dbo",
		Unmanaged: []management.Object{{Category: "tables", Object: "dbo.Legacy", Reason: "unsupported table options"}},
	}); err != nil {
		t.Fatal(err)
	}
	engine := New(cfg)
	report := output.NewReport(cfg.OutputNamingPattern, cfg.OutputNamingLineEndings)
	history := map[string]historyObject{
		objectKey("tables", "dbo.Legacy"): {
			Category: "tables", Object: "dbo.Legacy", LastVersion: "20260101000000000", LastVK: 1,
			SemanticSQL: "CREATE TABLE dbo.Legacy ([id] INT NOT NULL);\n",
		},
	}
	ops, _, _, err := engine.planGenerationWithInputs(&report, history, map[string]struct{}{}, "", nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(ops) != 0 || report.Fail.Count != 0 {
		t.Fatalf("unmanaged object must not be dropped: ops=%#v fail=%#v", ops, report.Fail.Items)
	}
	if !hasRule(report.Pass.Items, "UNMANAGED_OBJECT_REMOVAL_SKIPPED") {
		t.Fatalf("missing explicit audit warning: %#v", report.Pass.Items)
	}
}

func TestExistingPlainTableSkipsWhenRequiredColumnsAreBootstrapCompatible(t *testing.T) {
	create := `CREATE TABLE dbo.ExactState (
  [Id] BIGINT NOT NULL,
  [Code] NVARCHAR(40) COLLATE Latin1_General_100_BIN2 NULL
);`
	precheck := buildTableWrapperPrecheck("dbo.ExactState", true, create)
	for _, token := range []string{
		"c.name COLLATE Latin1_General_100_BIN2 = N'Id'",
		"c.name COLLATE Latin1_General_100_BIN2 = N'Code'",
		"UPPER(TYPE_NAME(c.user_type_id)) = N'BIGINT'",
		"UPPER(TYPE_NAME(c.user_type_id)) = N'NVARCHAR'",
		"Precheck: bootstrap-compatible table already exists",
		"missing a required column or has an incompatible base type",
	} {
		if !strings.Contains(precheck, token) {
			t.Fatalf("bootstrap-compatible precheck missing %q:\n%s", token, precheck)
		}
	}
	for _, ignored := range []string{"COUNT_BIG(*) FROM sys.columns", "c.max_length =", "c.is_nullable =", "collation_name"} {
		if strings.Contains(precheck, ignored) {
			t.Fatalf("non-blocking detail %q leaked into CREATE TABLE guard:\n%s", ignored, precheck)
		}
	}

	withConstraint := `CREATE TABLE dbo.ExactState (
  [Id] BIGINT NOT NULL,
  CONSTRAINT [PK_ExactState] PRIMARY KEY ([Id])
);`
	compatible := buildTableWrapperPrecheck("dbo.ExactState", true, withConstraint)
	for _, token := range []string{
		"bootstrap-compatible table already exists",
		"c.name COLLATE Latin1_General_100_BIN2 = N'Id'",
	} {
		if !strings.Contains(compatible, token) {
			t.Fatalf("column compatibility precheck missing %q:\n%s", token, compatible)
		}
	}
	if strings.Contains(compatible, "sys.key_constraints") || strings.Contains(compatible, "PK_ExactState") {
		t.Fatalf("CREATE TABLE bootstrap guard must not compare key constraints:\n%s", compatible)
	}
}

func TestExistingRowstoreIndexEquivalenceCoversOrderIncludeFilterAndOptions(t *testing.T) {
	stmt := `CREATE UNIQUE NONCLUSTERED INDEX [UX_Exact]
ON dbo.ExactState ([TenantId] ASC, [Code] DESC)
INCLUDE ([DisplayName])
WHERE [TenantId] IS NOT NULL
WITH (FILLFACTOR = 90, PAD_INDEX = ON, ALLOW_PAGE_LOCKS = OFF);`
	predicate, ok := buildExistingIndexEquivalentPredicate("dbo.ExactState", "UX_Exact", stmt)
	if !ok {
		t.Fatal("supported rowstore index equivalence unexpectedly unavailable")
	}
	for _, token := range []string{
		"i.is_unique = 1",
		"i.fill_factor = 90",
		"i.is_padded = 1",
		"i.allow_page_locks = 0",
		"(1, N'TenantId', 0)",
		"(2, N'Code', 1)",
		"N'DisplayName'",
		"i.has_filter = 1",
	} {
		if !strings.Contains(predicate, token) {
			t.Fatalf("index equivalence missing %q:\n%s", token, predicate)
		}
	}
	if _, supported := buildExistingIndexEquivalentPredicate(
		"dbo.ExactState", "CCI_Exact",
		"CREATE CLUSTERED COLUMNSTORE INDEX [CCI_Exact] ON dbo.ExactState;",
	); supported {
		t.Fatal("advanced index must fail closed instead of using incomplete equivalence")
	}
}

func TestUnsupportedFeatureWordsInCommentsLiteralsAndIdentifiersDoNotBlock(t *testing.T) {
	sql := `-- FILESTREAM and SYSTEM_VERSIONING are documentation here.
CREATE TABLE dbo.TokenProbe (
  [FILESTREAM] NVARCHAR(40) NULL DEFAULT (N'LEDGER SYSTEM_VERSIONING')
);`
	if err := validateTableSubsetSupport(sql); err != nil {
		t.Fatalf("non-SQL feature words caused a false positive: %v", err)
	}
}

func TestGeneratedWrapperFormattingIsCompactAndProductPrefixFree(t *testing.T) {
	sql := renderTxWrapper(
		"20991231235959999",
		"tables",
		"20991231235959999_tables_dbo.Clean_v1.up.sql",
		"dbo.Clean",
		true,
		"CREATE TABLE dbo.Clean ([Id] INT NOT NULL);\n\n\n",
	)
	if strings.Contains(sql, "\n\n\n") {
		t.Fatalf("generated wrapper contains excessive blank lines:\n%s", sql)
	}
	for _, line := range strings.Split(sql, "\n") {
		if line != strings.TrimRight(line, " \t") {
			t.Fatalf("generated wrapper contains trailing whitespace: %q", line)
		}
	}
	upper := strings.ToUpper(sql)
	if strings.Contains(upper, "RZA PRECHECK") || strings.Contains(upper, "RZA SCHEMA DRIFT") {
		t.Fatalf("generated human-facing messages contain verbose product prefix:\n%s", sql)
	}
	if !strings.Contains(sql, "SET XACT_ABORT ON;") {
		t.Fatalf("generated wrapper must abort transactions on runtime errors:\n%s", sql)
	}
}
