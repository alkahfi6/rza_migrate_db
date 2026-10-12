package engine_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/engine"
	"rza_migrate_db/internal/output"
)

func TestGenEmitOrderScenario06(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "06_ordering_determinism")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("engine run failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected no fail items, got %d", report.Fail.Count)
	}

	expected := []string{"tables", "indexes", "functions", "views", "procedures", "patch"}
	if len(report.NewMigrations.Files) != len(expected) {
		t.Fatalf("unexpected generated up files count: got=%d expected=%d", len(report.NewMigrations.Files), len(expected))
	}

	for i, upFile := range report.NewMigrations.Files {
		if !strings.HasSuffix(strings.ToLower(upFile), ".up.sql") {
			t.Fatalf("generated file must be .up.sql, got %s", upFile)
		}

		category, err := parseCategoryFromMigrationUpFile(upFile)
		if err != nil {
			t.Fatalf("invalid generated file name %s: %v", upFile, err)
		}
		if category != expected[i] {
			t.Fatalf("category order mismatch at idx %d: got=%s expected=%s file=%s", i, category, expected[i], upFile)
		}

		downFile := strings.TrimSuffix(upFile, ".up.sql") + ".down.sql"
		if _, err := os.Stat(filepath.Join(workRoot, "migrations", downFile)); err != nil {
			t.Fatalf("missing down pair for %s: %v", upFile, err)
		}
	}
}

func TestGenCreatesForeignKeyParentFirstAndPreservesColumnComments(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "06_ordering_determinism")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	tableDir := filepath.Join(workRoot, "db_object", "schema", "tables")
	parentSource := `CREATE TABLE dbo.ZParent (
  -- primary event identifier
  EventId INT NOT NULL,
  CONSTRAINT [PK_ZParent] PRIMARY KEY ([EventId])
);
`
	childSource := `CREATE TABLE dbo.AChild (
  -- foreign-key event identifier
  EventId INT NOT NULL,
  -- readable event name
  Nama VARCHAR(100),
  CONSTRAINT [PK_AChild] PRIMARY KEY ([EventId]),
  CONSTRAINT [FK_AChild_ZParent] FOREIGN KEY(EventId) REFERENCES dbo.ZParent([EventId])
);
`
	if err := os.WriteFile(filepath.Join(tableDir, "dbo.ZParent.sql"), []byte(parentSource), 0o644); err != nil {
		t.Fatalf("write parent source: %v", err)
	}
	if err := os.WriteFile(filepath.Join(tableDir, "dbo.AChild.sql"), []byte(childSource), 0o644); err != nil {
		t.Fatalf("write child source: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("engine run failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected no generation or staged-verification failures, got: %+v", report.Fail.Items)
	}

	parentIndex, childIndex := -1, -1
	parentUp, childUp := "", ""
	for index, name := range report.NewMigrations.Files {
		switch {
		case strings.Contains(name, "_tables_dbo.ZParent_"):
			parentIndex, parentUp = index, name
		case strings.Contains(name, "_tables_dbo.AChild_"):
			childIndex, childUp = index, name
		}
	}
	if parentIndex < 0 || childIndex < 0 {
		t.Fatalf("expected parent and child migration files, got=%v", report.NewMigrations.Files)
	}
	if parentIndex >= childIndex || parentUp >= childUp {
		t.Fatalf("referenced table must have an earlier generated migration: parent=%s child=%s", parentUp, childUp)
	}

	parentRaw, err := os.ReadFile(filepath.Join(workRoot, "migrations", parentUp))
	if err != nil {
		t.Fatalf("read parent migration: %v", err)
	}
	childRaw, err := os.ReadFile(filepath.Join(workRoot, "migrations", childUp))
	if err != nil {
		t.Fatalf("read child migration: %v", err)
	}
	parentSQL, childSQL := string(parentRaw), string(childRaw)
	for _, check := range []struct {
		name string
		sql  string
		want []string
	}{
		{
			name: "parent",
			sql:  parentSQL,
			want: []string{"CREATE TABLE dbo.ZParent", "-- primary event identifier", "[EventId] INT NOT NULL"},
		},
		{
			name: "child",
			sql:  childSQL,
			want: []string{
				"CREATE TABLE dbo.AChild",
				"-- foreign-key event identifier",
				"-- readable event name",
				"[Nama] VARCHAR(100) NULL",
				"CONSTRAINT [FK_AChild_ZParent] FOREIGN KEY(EventId) REFERENCES dbo.ZParent([EventId])",
			},
		},
	} {
		for _, want := range check.want {
			if !strings.Contains(check.sql, want) {
				t.Fatalf("%s migration missing %q:\n%s", check.name, want, check.sql)
			}
		}
		if strings.Contains(check.sql, "[--]") || strings.Contains(check.sql, "[/*]") {
			t.Fatalf("%s migration contains a comment parsed as an identifier:\n%s", check.name, check.sql)
		}
	}

	if edges, ok := report.Notes["dependency_edges"].([]string); ok {
		found := false
		for _, edge := range edges {
			if edge == "dbo.ZParent -> dbo.AChild" {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected FK dependency edge in report notes, got=%v", edges)
		}
	}
}

func TestGenTableTypeChangeRecreatesFunctionAndProcedureDependents(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "06_ordering_determinism")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	typeDir := filepath.Join(workRoot, "db_object", "types")
	functionDir := filepath.Join(workRoot, "db_object", "schema", "functions")
	procedureDir := filepath.Join(workRoot, "db_object", "schema", "procedures")
	if err := os.MkdirAll(typeDir, 0o755); err != nil {
		t.Fatal(err)
	}
	typePath := filepath.Join(typeDir, "dbo.OrderItemType.sql")
	functionPath := filepath.Join(functionDir, "dbo.fnCountItems.sql")
	procedurePath := filepath.Join(procedureDir, "dbo.SaveItems.sql")

	oldType := "CREATE TYPE dbo.OrderItemType AS TABLE (\n  Id INT NOT NULL,\n  PRIMARY KEY CLUSTERED ([Id])\n);\n"
	newType := "CREATE TYPE dbo.OrderItemType AS TABLE (\n  Id INT NOT NULL,\n  Qty INT NULL,\n  PRIMARY KEY CLUSTERED ([Id])\n);\n"
	oldFunction := "CREATE OR ALTER FUNCTION dbo.fnCountItems(@items dbo.OrderItemType READONLY) RETURNS INT AS BEGIN RETURN (SELECT COUNT(*) FROM @items); END;\n"
	newFunction := "CREATE OR ALTER FUNCTION dbo.fnCountItems(@items dbo.OrderItemType READONLY) RETURNS BIGINT AS BEGIN RETURN (SELECT COUNT_BIG(*) FROM @items); END;\n"
	oldProcedure := "CREATE OR ALTER PROCEDURE dbo.SaveItems @items dbo.OrderItemType READONLY AS SELECT dbo.fnCountItems(@items) AS item_count;\n"
	newProcedure := "CREATE OR ALTER PROCEDURE dbo.SaveItems @items dbo.OrderItemType READONLY AS SELECT dbo.fnCountItems(@items) AS total_items;\n"

	for path, content := range map[string]string{
		typePath:      oldType,
		functionPath:  oldFunction,
		procedurePath: oldProcedure,
	} {
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			t.Fatalf("write initial source %s: %v", path, err)
		}
	}

	eng := engine.New(config.Default(workRoot))
	first, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("initial gen failed: %v", err)
	}
	if first.Fail.Count != 0 {
		t.Fatalf("initial gen validation failed: %+v", first.Fail.Items)
	}

	for path, content := range map[string]string{
		typePath:      newType,
		functionPath:  newFunction,
		procedurePath: newProcedure,
	} {
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			t.Fatalf("write changed source %s: %v", path, err)
		}
	}

	second, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("alter gen failed: %v", err)
	}
	if second.Fail.Count != 0 {
		t.Fatalf("alter gen validation failed: %+v", second.Fail.Items)
	}

	typeIndex, functionIndex, procedureIndex := -1, -1, -1
	typeUp := ""
	for index, name := range second.NewMigrations.Files {
		switch {
		case strings.Contains(name, "_types_dbo.OrderItemType_v2.up.sql"):
			typeIndex, typeUp = index, name
		case strings.Contains(name, "_functions_dbo.fnCountItems_v2.up.sql"):
			functionIndex = index
		case strings.Contains(name, "_procedures_dbo.SaveItems_v2.up.sql"):
			procedureIndex = index
		}
	}
	if typeIndex < 0 || functionIndex < 0 || procedureIndex < 0 {
		t.Fatalf("missing type/function/procedure v2 migrations: %v", second.NewMigrations.Files)
	}
	if typeIndex >= functionIndex || typeIndex >= procedureIndex {
		t.Fatalf("type migration must precede dependent modules: %v", second.NewMigrations.Files)
	}

	raw, err := os.ReadFile(filepath.Join(workRoot, "migrations", typeUp))
	if err != nil {
		t.Fatal(err)
	}
	typeMigration := string(raw)
	for _, want := range []string{
		"DROP PROCEDURE IF EXISTS dbo.SaveItems",
		"DROP FUNCTION IF EXISTS dbo.fnCountItems",
		"DROP TYPE IF EXISTS dbo.OrderItemType",
		"CREATE TYPE dbo.OrderItemType AS TABLE",
		"Qty INT NULL",
	} {
		if !strings.Contains(typeMigration, want) {
			t.Fatalf("type lifecycle migration missing %q:\n%s", want, typeMigration)
		}
	}
	if strings.Contains(typeMigration, "EXEC sys.sp_executesql") {
		t.Fatalf("changed dependents must be deferred to their ordered migrations:\n%s", typeMigration)
	}

	third, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("idempotence gen failed: %v", err)
	}
	if third.Fail.Count != 0 || len(third.NewMigrations.Files) != 0 {
		t.Fatalf("TVP lifecycle must be idempotent, failures=%+v files=%v", third.Fail.Items, third.NewMigrations.Files)
	}
}

func TestGenNoEmitOnWhitespaceOnlyScenario07(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "07_vk_whitespace_only")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("engine run failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected no fail items, got %d", report.Fail.Count)
	}
	if len(report.NewMigrations.Files) != 0 {
		t.Fatalf("expected no generated migrations for whitespace-only scenario, got %d", len(report.NewMigrations.Files))
	}
}

func TestGenUnsupportedHistorySubsetScenario08(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "08_reconstruction_unsupported_sql")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("engine run failed: %v", err)
	}
	if report.Fail.Count == 0 {
		t.Fatalf("expected fail for unsupported SQL subset in history")
	}
	hasUnsupported := false
	for _, item := range report.Fail.Items {
		if item.Rule == "UNSUPPORTED_SQL_SUBSET" {
			hasUnsupported = true
			break
		}
	}
	if !hasUnsupported {
		t.Fatalf("expected fail rule UNSUPPORTED_SQL_SUBSET, got fail items: %+v", report.Fail.Items)
	}
	if len(report.NewMigrations.Files) != 0 {
		t.Fatalf("expected no emitted migrations on failure, got %d", len(report.NewMigrations.Files))
	}
}

func TestGenDropTableGuardedRemovalScenario11(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "11_destructive_drop_table_manual")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("engine run failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected no fail items, got %d", report.Fail.Count)
	}
	if !hasPassReason(report, "DESTRUCTIVE_CHANGE") {
		t.Fatalf("expected guarded destructive-change warning in scenario 11")
	}
	if len(report.NewMigrations.Files) == 0 {
		t.Fatalf("expected guarded drop-table migration pair to be generated")
	}

	upName := ""
	for _, name := range report.NewMigrations.Files {
		if strings.Contains(name, "_tables_") {
			upName = name
			break
		}
	}
	if upName == "" {
		t.Fatalf("expected generated tables removal migration")
	}

	upPath := filepath.Join(workRoot, "migrations", upName)
	upRaw, err := os.ReadFile(upPath)
	if err != nil {
		t.Fatalf("read generated up migration failed: %v", err)
	}
	upSQL := strings.ToUpper(string(upRaw))
	for _, token := range []string{"DROP TABLE IF EXISTS", "TARGET CONTAINS DATA", "SYS.FOREIGN_KEYS", "SYS.SQL_EXPRESSION_DEPENDENCIES", "SYS.TRIGGERS"} {
		if !strings.Contains(upSQL, token) {
			t.Fatalf("guarded drop-table migration missing %q:\n%s", token, upSQL)
		}
	}

	downPath := filepath.Join(workRoot, "migrations", strings.TrimSuffix(upName, ".up.sql")+".down.sql")
	downRaw, err := os.ReadFile(downPath)
	if err != nil {
		t.Fatalf("read generated down migration failed: %v", err)
	}
	downSQL := strings.ToUpper(string(downRaw))
	if !strings.Contains(downSQL, "CREATE TABLE") {
		t.Fatalf("drop-table down migration must restore the previous table schema")
	}
}

func TestGenDropTableRemovalSecondRunNoReemit(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "11_destructive_drop_table_manual")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	first, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("first gen run failed: %v", err)
	}
	if first.Fail.Count != 0 {
		t.Fatalf("first gen expected no fail items, got %d: %+v", first.Fail.Count, first.Fail.Items)
	}
	if len(first.NewMigrations.Files) == 0 {
		t.Fatalf("first gen should emit guarded drop-table migration")
	}

	second, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("second gen run failed: %v", err)
	}
	if second.Fail.Count != 0 {
		t.Fatalf("second gen expected no fail items, got %d", second.Fail.Count)
	}
	if len(second.NewMigrations.Files) != 0 {
		t.Fatalf("second gen should not re-emit drop-table migration, got %d", len(second.NewMigrations.Files))
	}
	if !hasPassRule(second, "GEN_NO_CHANGES") {
		t.Fatalf("expected idempotent GEN_NO_CHANGES marker on second run")
	}
}

func TestGenDownRestoreWarningScenario14(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "14_down_data_restore_warning")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("engine run failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected no fail items, got %d", report.Fail.Count)
	}
	if !hasPassReason(report, "DATA_RESTORE_NOT_GUARANTEED") {
		t.Fatalf("expected pass reason DATA_RESTORE_NOT_GUARANTEED in scenario 14")
	}
	if !hasPassReason(report, "INDEX_RECREATE_REQUIRED") {
		t.Fatalf("expected pass reason INDEX_RECREATE_REQUIRED in scenario 14")
	}

	tableUp := ""
	for _, name := range report.NewMigrations.Files {
		if strings.Contains(name, "_tables_") {
			tableUp = name
			break
		}
	}
	if tableUp == "" {
		t.Fatalf("expected generated table migration in scenario 14")
	}

	upRaw, err := os.ReadFile(filepath.Join(workRoot, "migrations", tableUp))
	if err != nil {
		t.Fatalf("read generated up migration failed: %v", err)
	}
	upSQL := strings.ToUpper(string(upRaw))
	if strings.Contains(upSQL, "SP_EXECUTESQL") {
		t.Fatalf("drop-column up migration must not use dynamic SQL")
	}

	tableDown := strings.TrimSuffix(tableUp, ".up.sql") + ".down.sql"
	downRaw, err := os.ReadFile(filepath.Join(workRoot, "migrations", tableDown))
	if err != nil {
		t.Fatalf("read generated down migration failed: %v", err)
	}
	downSQL := strings.ToUpper(string(downRaw))
	if strings.Contains(downSQL, "CREATE TABLE DBO.") {
		t.Fatalf("drop-column down migration must not rebuild table using CREATE TABLE snapshot")
	}
	if !strings.Contains(downSQL, "ALTER TABLE DBO.TABLEA ADD [AMOUNT] DECIMAL(10,2) NULL;") {
		t.Fatalf("drop-column down migration must re-add dropped column using ALTER TABLE ADD")
	}
}

func TestGenReconstructDropColumnHistorySecondRunNoEmit(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "14_down_data_restore_warning")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	first, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("first gen run failed: %v", err)
	}
	if first.Fail.Count != 0 {
		t.Fatalf("first gen expected no fail items, got %d: %+v", first.Fail.Count, first.Fail.Items)
	}
	if len(first.NewMigrations.Files) == 0 {
		t.Fatalf("first gen should emit drop-column migration")
	}

	second, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("second gen run failed: %v", err)
	}
	if second.Fail.Count != 0 {
		t.Fatalf("second gen expected no fail items, got %d", second.Fail.Count)
	}
	if len(second.NewMigrations.Files) != 0 {
		t.Fatalf("second gen should not emit migrations after history reconstruction, got %d", len(second.NewMigrations.Files))
	}
}

func TestGenDropColumnDeterministicDependencyDrop(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "07_vk_whitespace_only")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	historyTableUpPath := filepath.Join(workRoot, "migrations", "20260217073015000_tables_dbo.TableA_v1.up.sql")
	historyTableUpRaw, err := os.ReadFile(historyTableUpPath)
	if err != nil {
		t.Fatalf("read history table migration failed: %v", err)
	}
	historyTableUpSQL := strings.Replace(
		string(historyTableUpRaw),
		"    CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status",
		"    CONSTRAINT UQ_TableA__amount UNIQUE (amount),\n    CONSTRAINT CK_TableA__amount_nonneg CHECK (amount >= 0),\n    CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status",
		1,
	)
	if err := os.WriteFile(historyTableUpPath, []byte(historyTableUpSQL), 0o644); err != nil {
		t.Fatalf("write history table migration failed: %v", err)
	}

	historyIndexUpPath := filepath.Join(workRoot, "migrations", "20260217073015001_indexes_dbo.TableA_v1.up.sql")
	historyIndexUpRaw, err := os.ReadFile(historyIndexUpPath)
	if err != nil {
		t.Fatalf("read history index migration failed: %v", err)
	}
	historyIndexUpSQL := strings.Replace(
		string(historyIndexUpRaw),
		"  CREATE INDEX IX_TableA__status ON dbo.TableA(status);",
		"  CREATE INDEX IX_TableA__status ON dbo.TableA(status);\n  CREATE INDEX IX_TableA__amount ON dbo.TableA(amount);",
		1,
	)
	if err := os.WriteFile(historyIndexUpPath, []byte(historyIndexUpSQL), 0o644); err != nil {
		t.Fatalf("write history index migration failed: %v", err)
	}

	tableSourcePath := filepath.Join(workRoot, "db_object", "schema", "tables", "dbo.TableA.sql")
	tableSourceSQL := `CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id),
  CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status
);
`
	if err := os.WriteFile(tableSourcePath, []byte(tableSourceSQL), 0o644); err != nil {
		t.Fatalf("write source table failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("engine run failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected no fail items, got %d", report.Fail.Count)
	}

	tableUp := ""
	for _, name := range report.NewMigrations.Files {
		if strings.Contains(name, "_tables_") {
			tableUp = name
			break
		}
	}
	if tableUp == "" {
		t.Fatalf("expected generated table drop-column migration")
	}

	upRaw, err := os.ReadFile(filepath.Join(workRoot, "migrations", tableUp))
	if err != nil {
		t.Fatalf("read generated up migration failed: %v", err)
	}
	upSQL := strings.ToUpper(string(upRaw))
	if strings.Contains(upSQL, "SP_EXECUTESQL") {
		t.Fatalf("drop-column up migration must not use dynamic SQL")
	}
	if !strings.Contains(upSQL, "ALTER TABLE DBO.TABLEA DROP CONSTRAINT [CK_TABLEA__AMOUNT_NONNEG];") {
		t.Fatalf("drop-column up migration must drop impacted check constraint deterministically")
	}
	if !strings.Contains(upSQL, "ALTER TABLE DBO.TABLEA DROP CONSTRAINT [UQ_TABLEA__AMOUNT];") {
		t.Fatalf("drop-column up migration must drop impacted unique constraint deterministically")
	}
	if !strings.Contains(upSQL, "DROP INDEX IF EXISTS [IX_TABLEA__AMOUNT] ON DBO.TABLEA;") {
		t.Fatalf("drop-column up migration must drop impacted index deterministically")
	}
	if !strings.Contains(upSQL, "ALTER TABLE DBO.TABLEA DROP COLUMN [AMOUNT];") {
		t.Fatalf("drop-column up migration must drop target column")
	}
	for _, name := range report.NewMigrations.Files {
		if strings.Contains(name, "_indexes_") {
			t.Fatalf("drop-column with index side-effect should not emit redundant indexes migration in same run")
		}
	}
}

func TestGenDropColumnCrossTableFKDeterministicAndIdempotent(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "07_vk_whitespace_only")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	tableBHistoryUp := `/*
Severity mapping:
0=INFO (ok/skip normal)
1=NOTICE (schema change)
2=WARNING (follow-up/manual risk)
3=ERROR (stop)
*/
SET NOCOUNT ON;
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = 20260217073015005;
DECLARE @category NVARCHAR(16) = N'tables';
DECLARE @file_name NVARCHAR(260) = N'20260217073015005_tables_dbo.TableB_v1.up.sql';
DECLARE @target NVARCHAR(256) = N'dbo.TableB';
DECLARE @action NVARCHAR(64) = N'create_table';

BEGIN TRY
  BEGIN TRAN;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file_name,
      @decision = N'run',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'OBJECT_EXISTS',
      @message = N'migration run start',
      @detail = N'up',
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END

  CREATE TABLE dbo.TableB (
    id INT NOT NULL,
    tablea_amount DECIMAL(10,2) NULL,
    CONSTRAINT PK_TableB PRIMARY KEY (id),
    CONSTRAINT FK_TableB__TableA_amount FOREIGN KEY (tablea_amount) REFERENCES dbo.TableA(amount)
  );

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file_name,
      @decision = N'ok',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'CREATED',
      @message = N'migration step ok',
      @detail = N'up',
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END

  COMMIT;
END TRY
BEGIN CATCH
  IF @@TRANCOUNT > 0 ROLLBACK;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file_name,
      @decision = N'fail',
      @severity = 3,
      @action = N'migration_failed',
      @target = @target,
      @reason = N'SQL_ERROR',
      @message = N'migration failed',
      @detail = ERROR_MESSAGE(),
      @error_number = ERROR_NUMBER(),
      @error_line = ERROR_LINE(),
      @error_procedure = ERROR_PROCEDURE();
  END

  THROW;
END CATCH;
`
	tableBHistoryDown := `/*
Severity mapping:
0=INFO (ok/skip normal)
1=NOTICE (schema change)
2=WARNING (follow-up/manual risk)
3=ERROR (stop)
*/
SET NOCOUNT ON;
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = 20260217073015005;
DECLARE @category NVARCHAR(16) = N'tables';
DECLARE @file_name NVARCHAR(260) = N'20260217073015005_tables_dbo.TableB_v1.down.sql';
DECLARE @target NVARCHAR(256) = N'dbo.TableB';
DECLARE @action NVARCHAR(64) = N'drop_table';

BEGIN TRY
  BEGIN TRAN;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file_name,
      @decision = N'run',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'OBJECT_EXISTS',
      @message = N'migration run start',
      @detail = N'down',
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END

  DROP TABLE dbo.TableB;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file_name,
      @decision = N'ok',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'DROPPED',
      @message = N'migration step ok',
      @detail = N'down',
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END

  COMMIT;
END TRY
BEGIN CATCH
  IF @@TRANCOUNT > 0 ROLLBACK;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file_name,
      @decision = N'fail',
      @severity = 3,
      @action = N'migration_failed',
      @target = @target,
      @reason = N'SQL_ERROR',
      @message = N'migration failed',
      @detail = ERROR_MESSAGE(),
      @error_number = ERROR_NUMBER(),
      @error_line = ERROR_LINE(),
      @error_procedure = ERROR_PROCEDURE();
  END

  THROW;
END CATCH;
`
	if err := os.WriteFile(
		filepath.Join(workRoot, "migrations", "20260217073015005_tables_dbo.TableB_v1.up.sql"),
		[]byte(tableBHistoryUp),
		0o644,
	); err != nil {
		t.Fatalf("write tableB history up failed: %v", err)
	}
	if err := os.WriteFile(
		filepath.Join(workRoot, "migrations", "20260217073015005_tables_dbo.TableB_v1.down.sql"),
		[]byte(tableBHistoryDown),
		0o644,
	); err != nil {
		t.Fatalf("write tableB history down failed: %v", err)
	}

	tableASourcePath := filepath.Join(workRoot, "db_object", "schema", "tables", "dbo.TableA.sql")
	tableASource := `CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id),
  CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status
);
`
	if err := os.WriteFile(tableASourcePath, []byte(tableASource), 0o644); err != nil {
		t.Fatalf("write source tableA failed: %v", err)
	}

	tableBSourcePath := filepath.Join(workRoot, "db_object", "schema", "tables", "dbo.TableB.sql")
	tableBSource := `CREATE TABLE dbo.TableB (
  id INT NOT NULL,
  tablea_amount DECIMAL(10,2) NULL,
  CONSTRAINT PK_TableB PRIMARY KEY (id)
);
`
	if err := os.WriteFile(tableBSourcePath, []byte(tableBSource), 0o644); err != nil {
		t.Fatalf("write source tableB failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	first, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("first gen run failed: %v", err)
	}
	if first.Fail.Count != 0 {
		t.Fatalf("first gen expected no fail items, got %d: %+v", first.Fail.Count, first.Fail.Items)
	}
	if hasPassRule(first, "TABLE_CHANGE_REQUIRES_MANUAL") {
		t.Fatalf("cross-table FK drop should be auto planned, not manual-skipped")
	}

	tableUp := ""
	for _, name := range first.NewMigrations.Files {
		if strings.Contains(name, "_tables_") && strings.Contains(name, "dbo.TableA") {
			tableUp = name
			break
		}
	}
	if tableUp == "" {
		t.Fatalf("expected generated tableA drop-column migration")
	}

	upRaw, err := os.ReadFile(filepath.Join(workRoot, "migrations", tableUp))
	if err != nil {
		t.Fatalf("read generated up migration failed: %v", err)
	}
	upSQL := strings.ToUpper(string(upRaw))
	if !strings.Contains(upSQL, "ALTER TABLE DBO.TABLEB DROP CONSTRAINT [FK_TABLEB__TABLEA_AMOUNT];") {
		t.Fatalf("drop-column migration must drop incoming FK from referencing table deterministically")
	}
	if strings.Contains(upSQL, "SP_EXECUTESQL") {
		t.Fatalf("drop-column up migration must not use dynamic SQL")
	}

	second, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("second gen run failed: %v", err)
	}
	if second.Fail.Count != 0 {
		t.Fatalf("second gen expected no fail items, got %d", second.Fail.Count)
	}
	if len(second.NewMigrations.Files) != 0 {
		t.Fatalf("second gen should not emit migrations after reconstruction, got %d", len(second.NewMigrations.Files))
	}
	if hasPassRule(second, "TABLE_CHANGE_REQUIRES_MANUAL") {
		t.Fatalf("reconstruction should apply dropped FK side-effect and avoid manual drift warning")
	}
}

func TestGenIndexLifecycleDiffUsesDropAndCreate(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "07_vk_whitespace_only")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	indexPath := filepath.Join(workRoot, "db_object", "schema", "indexes", "dbo.TableA.sql")
	indexSQL := `CREATE INDEX IX_TableA__amount ON dbo.TableA(amount);
`
	if err := os.WriteFile(indexPath, []byte(indexSQL), 0o644); err != nil {
		t.Fatalf("write index source failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("gen run failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected no fail items, got %d", report.Fail.Count)
	}

	indexUp := ""
	for _, name := range report.NewMigrations.Files {
		if strings.Contains(name, "_indexes_") {
			indexUp = name
			break
		}
	}
	if indexUp == "" {
		t.Fatalf("expected generated indexes migration")
	}

	upRaw, err := os.ReadFile(filepath.Join(workRoot, "migrations", indexUp))
	if err != nil {
		t.Fatalf("read indexes up migration failed: %v", err)
	}
	upSQL := strings.ToUpper(string(upRaw))
	if !strings.Contains(upSQL, "DROP INDEX IF EXISTS [IX_TABLEA__STATUS] ON DBO.TABLEA;") {
		t.Fatalf("indexes up migration must drop removed/changed index first")
	}
	if !strings.Contains(upSQL, "CREATE INDEX IX_TABLEA__AMOUNT ON DBO.TABLEA(AMOUNT);") {
		t.Fatalf("indexes up migration must create desired index")
	}
	if !strings.Contains(upSQL, "IF NOT EXISTS (") || !strings.Contains(upSQL, "FROM SYS.INDEXES") {
		t.Fatalf("indexes up migration create step must be guarded by per-index existence check")
	}

	indexDown := strings.TrimSuffix(indexUp, ".up.sql") + ".down.sql"
	downRaw, err := os.ReadFile(filepath.Join(workRoot, "migrations", indexDown))
	if err != nil {
		t.Fatalf("read indexes down migration failed: %v", err)
	}
	downSQL := strings.ToUpper(string(downRaw))
	if !strings.Contains(downSQL, "DROP INDEX IF EXISTS [IX_TABLEA__AMOUNT] ON DBO.TABLEA;") {
		t.Fatalf("indexes down migration must drop new index before restore")
	}
	if !strings.Contains(downSQL, "CREATE INDEX IX_TABLEA__STATUS ON DBO.TABLEA(STATUS);") {
		t.Fatalf("indexes down migration must recreate previous index")
	}
	if !strings.Contains(downSQL, "IF NOT EXISTS (") || !strings.Contains(downSQL, "FROM SYS.INDEXES") {
		t.Fatalf("indexes down migration recreate step must be guarded by per-index existence check")
	}
}

func TestGenIndexSourceRemovedEmitsDropOnce(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "07_vk_whitespace_only")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	indexPath := filepath.Join(workRoot, "db_object", "schema", "indexes", "dbo.TableA.sql")
	if err := os.Remove(indexPath); err != nil {
		t.Fatalf("remove index source failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	first, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("first gen run failed: %v", err)
	}
	if first.Fail.Count != 0 {
		t.Fatalf("first gen expected no fail items, got %d", first.Fail.Count)
	}

	indexUp := ""
	for _, name := range first.NewMigrations.Files {
		if strings.Contains(name, "_indexes_") {
			indexUp = name
			break
		}
	}
	if indexUp == "" {
		t.Fatalf("expected generated indexes drop migration when source file removed")
	}

	upRaw, err := os.ReadFile(filepath.Join(workRoot, "migrations", indexUp))
	if err != nil {
		t.Fatalf("read indexes up migration failed: %v", err)
	}
	upSQL := strings.ToUpper(string(upRaw))
	if !strings.Contains(upSQL, "DROP INDEX IF EXISTS [IX_TABLEA__STATUS] ON DBO.TABLEA;") {
		t.Fatalf("indexes drop migration must drop existing index deterministically")
	}

	second, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("second gen run failed: %v", err)
	}
	if second.Fail.Count != 0 {
		t.Fatalf("second gen expected no fail items, got %d", second.Fail.Count)
	}
	for _, name := range second.NewMigrations.Files {
		if strings.Contains(name, "_indexes_") {
			t.Fatalf("second gen should not re-emit indexes drop migration")
		}
	}
}

func TestGeneratedWrapperContainsAuditCalls(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "06_ordering_determinism")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("engine run failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected no fail items, got %d", report.Fail.Count)
	}
	if len(report.NewMigrations.Files) == 0 {
		t.Fatalf("expected generated files")
	}

	tableUp := ""
	for _, name := range report.NewMigrations.Files {
		if strings.Contains(name, "_tables_") {
			tableUp = name
			break
		}
	}
	if tableUp == "" {
		t.Fatalf("generated table migration not found")
	}

	content, err := os.ReadFile(filepath.Join(workRoot, "migrations", tableUp))
	if err != nil {
		t.Fatalf("read generated file failed: %v", err)
	}
	sql := strings.ToUpper(string(content))
	assertContains := func(token string) {
		t.Helper()
		if !strings.Contains(sql, token) {
			t.Fatalf("generated wrapper missing token: %s", token)
		}
	}

	assertContains("EXEC DBO.USP_MIG_LOG_EVENT")
	assertContains("@DECISION = N'RUN'")
	assertContains("@DECISION = N'OK'")
	assertContains("@DECISION = N'FAIL'")
	assertContains("@ACTION = N'MIGRATION_FAILED'")
	assertContains("@REASON = N'SQL_ERROR'")
}

func TestGeneratedWrapperContainsDeterministicPrecheckSkip(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "25_wrapper_precheck_skip")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("engine run failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected no fail items, got %d: %+v", report.Fail.Count, report.Fail.Items)
	}

	tableUp := ""
	indexUp := ""
	for _, name := range report.NewMigrations.Files {
		if strings.Contains(name, "_tables_") {
			tableUp = name
		}
		if strings.Contains(name, "_indexes_") {
			indexUp = name
		}
	}
	if tableUp == "" {
		t.Fatalf("expected generated table migration")
	}
	if indexUp == "" {
		t.Fatalf("expected generated indexes migration")
	}

	readUpper := func(path string) string {
		t.Helper()
		raw, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read migration failed: %v", err)
		}
		return strings.ToUpper(string(raw))
	}
	assertContainsAll := func(sql string, tokens ...string) {
		t.Helper()
		for _, token := range tokens {
			if !strings.Contains(sql, token) {
				t.Fatalf("generated wrapper missing token: %s", token)
			}
		}
	}

	tableUpSQL := readUpper(filepath.Join(workRoot, "migrations", tableUp))
	assertContainsAll(
		tableUpSQL,
		"DECLARE @FILE NVARCHAR(260)",
		"@FILE_NAME = @FILE",
		"IF OBJECT_ID(N'DBO.TABLEA', N'U') IS NOT NULL",
		"EXISTING TABLE MATCH PROFILE: BOOTSTRAP-COMPATIBLE",
		"FROM SYS.COLUMNS AS C",
		"TYPE_NAME(C.USER_TYPE_ID)",
		"C.IS_IDENTITY =",
		"C.IS_COMPUTED =",
		"N'PRECHECK: BOOTSTRAP-COMPATIBLE TABLE ALREADY EXISTS'",
		"SCHEMA DRIFT: EXISTING TABLE IS MISSING A REQUIRED COLUMN",
		"THROW 51000",
	)
	for _, ignored := range []string{"COUNT_BIG(*) FROM SYS.COLUMNS", "C.MAX_LENGTH =", "C.IS_NULLABLE =", "FROM SYS.FOREIGN_KEYS AS FK"} {
		if strings.Contains(tableUpSQL, ignored) {
			t.Fatalf("generated bootstrap-compatible wrapper contains blocking detail %s", ignored)
		}
	}

	tableDownSQL := readUpper(filepath.Join(workRoot, "migrations", strings.TrimSuffix(tableUp, ".up.sql")+".down.sql"))
	assertContainsAll(
		tableDownSQL,
		"IF OBJECT_ID(N'DBO.TABLEA', N'U') IS NULL",
		"@DECISION = N'SKIP'",
		"@REASON = N'OBJECT_NOT_EXISTS'",
		"N'PRECHECK: TABLE ALREADY ABSENT'",
		"RETURN;",
	)

	indexUpSQL := readUpper(filepath.Join(workRoot, "migrations", indexUp))
	assertContainsAll(
		indexUpSQL,
		"-- PRECHECK IS HANDLED PER STATEMENT IN THE MIGRATION BODY",
		"IF NOT EXISTS (",
		"FROM SYS.INDEXES",
		"OBJECT_ID = OBJECT_ID(N'DBO.TABLEA')",
		"NAME COLLATE LATIN1_GENERAL_100_BIN2 = N'IX_TABLEA__STATUS'",
		"SCHEMA DRIFT:",
	)

	indexDownSQL := readUpper(filepath.Join(workRoot, "migrations", strings.TrimSuffix(indexUp, ".up.sql")+".down.sql"))
	assertContainsAll(
		indexDownSQL,
		"-- PRECHECK IS HANDLED PER STATEMENT IN THE MIGRATION BODY",
		"IF EXISTS (",
		"FROM SYS.INDEXES",
		"OBJECT_ID = OBJECT_ID(N'DBO.TABLEA')",
		"NAME COLLATE LATIN1_GENERAL_100_BIN2 = N'IX_TABLEA__STATUS'",
		"@DECISION = N'SKIP'",
		"@REASON = N'INDEX_NOT_EXISTS'",
		"N'PRECHECK: INDEX ALREADY ABSENT'",
	)
}

func TestGeneratedWrapperContainsDeterministicPrecheckSkipMultiIndex(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "25_wrapper_precheck_skip")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	indexPath := filepath.Join(workRoot, "db_object", "schema", "indexes", "dbo.TableA.sql")
	indexSQL := `CREATE INDEX IX_TableA__status ON dbo.TableA(status);
CREATE INDEX IX_TableA__amount ON dbo.TableA(amount);
`
	if err := os.WriteFile(indexPath, []byte(indexSQL), 0o644); err != nil {
		t.Fatalf("write index source failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("engine run failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected no fail items, got %d: %+v", report.Fail.Count, report.Fail.Items)
	}

	indexUp := ""
	for _, name := range report.NewMigrations.Files {
		if strings.Contains(name, "_indexes_") {
			indexUp = name
			break
		}
	}
	if indexUp == "" {
		t.Fatalf("expected generated indexes migration")
	}

	readUpper := func(path string) string {
		t.Helper()
		raw, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read migration failed: %v", err)
		}
		return strings.ToUpper(string(raw))
	}

	indexUpSQL := readUpper(filepath.Join(workRoot, "migrations", indexUp))
	if !strings.Contains(indexUpSQL, "-- PRECHECK IS HANDLED PER STATEMENT IN THE MIGRATION BODY") {
		t.Fatalf("multi-index up migration must use per-index precheck and keep wrapper precheck disabled")
	}
	for _, token := range []string{
		"IF NOT EXISTS (",
		"FROM SYS.INDEXES",
		"NAME COLLATE LATIN1_GENERAL_100_BIN2 = N'IX_TABLEA__STATUS'",
		"NAME COLLATE LATIN1_GENERAL_100_BIN2 = N'IX_TABLEA__AMOUNT'",
		"SCHEMA DRIFT:",
		"IX_TABLEA__STATUS",
		"IX_TABLEA__AMOUNT",
	} {
		if !strings.Contains(indexUpSQL, token) {
			t.Fatalf("generated multi-index up wrapper missing token: %s", token)
		}
	}

	indexDownSQL := readUpper(filepath.Join(workRoot, "migrations", strings.TrimSuffix(indexUp, ".up.sql")+".down.sql"))
	if !strings.Contains(indexDownSQL, "-- PRECHECK IS HANDLED PER STATEMENT IN THE MIGRATION BODY") {
		t.Fatalf("multi-index down migration must use per-index precheck and keep wrapper precheck disabled")
	}
	for _, token := range []string{
		"IF EXISTS (",
		"FROM SYS.INDEXES",
		"NAME COLLATE LATIN1_GENERAL_100_BIN2 = N'IX_TABLEA__STATUS'",
		"NAME COLLATE LATIN1_GENERAL_100_BIN2 = N'IX_TABLEA__AMOUNT'",
		"@REASON = N'INDEX_NOT_EXISTS'",
		"N'PRECHECK: INDEX ALREADY ABSENT'",
		"IX_TABLEA__STATUS",
		"IX_TABLEA__AMOUNT",
	} {
		if !strings.Contains(indexDownSQL, token) {
			t.Fatalf("generated multi-index down wrapper missing token: %s", token)
		}
	}
}

func TestGenTableAddColumnUsesAlterAndSecondRunNoEmit(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "07_vk_whitespace_only")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	tablePath := filepath.Join(workRoot, "db_object", "schema", "tables", "dbo.TableA.sql")
	raw, err := os.ReadFile(tablePath)
	if err != nil {
		t.Fatalf("read table source failed: %v", err)
	}
	updated := strings.Replace(string(raw), "amount DECIMAL(10,2) NULL,", "amount DECIMAL(10,2) NULL,\n  note VARCHAR(20) NULL,", 1)
	if err := os.WriteFile(tablePath, []byte(updated), 0o644); err != nil {
		t.Fatalf("write table source failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	first, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("first gen run failed: %v", err)
	}
	if first.Fail.Count != 0 {
		t.Fatalf("first gen expected no fail items, got %d", first.Fail.Count)
	}
	if len(first.NewMigrations.Files) == 0 {
		t.Fatalf("first gen should emit table alter migration")
	}

	tableUp := ""
	for _, name := range first.NewMigrations.Files {
		if strings.Contains(name, "_tables_") {
			tableUp = name
			break
		}
	}
	if tableUp == "" {
		t.Fatalf("expected generated table migration")
	}

	upRaw, err := os.ReadFile(filepath.Join(workRoot, "migrations", tableUp))
	if err != nil {
		t.Fatalf("read generated up migration failed: %v", err)
	}
	upSQL := strings.ToUpper(string(upRaw))
	if strings.Contains(upSQL, "CREATE TABLE DBO.TABLEA") {
		t.Fatalf("table alter migration for existing table must not use CREATE TABLE")
	}
	if !strings.Contains(upSQL, "ALTER TABLE DBO.TABLEA ADD [NOTE] VARCHAR(20) NULL;") {
		t.Fatalf("table alter up migration must add new column via ALTER TABLE ADD")
	}
	if !strings.Contains(upSQL, "@REASON = N'COLUMN_EXISTS'") {
		t.Fatalf("table alter up migration must include add-column precheck reason")
	}
	if !strings.Contains(upSQL, "C.NAME COLLATE LATIN1_GENERAL_100_BIN2 = N'NOTE'") {
		t.Fatalf("table alter up migration must include add-column existence precheck")
	}
	if !strings.Contains(upSQL, "N'PRECHECK: TARGET COLUMN(S) ALREADY EXIST'") {
		t.Fatalf("table alter up migration must include add-column precheck message")
	}

	tableDown := strings.TrimSuffix(tableUp, ".up.sql") + ".down.sql"
	downRaw, err := os.ReadFile(filepath.Join(workRoot, "migrations", tableDown))
	if err != nil {
		t.Fatalf("read generated down migration failed: %v", err)
	}
	downSQL := strings.ToUpper(string(downRaw))
	if !strings.Contains(downSQL, "ALTER TABLE DBO.TABLEA DROP COLUMN [NOTE];") {
		t.Fatalf("table alter down migration must reverse added column using DROP COLUMN")
	}
	if !strings.Contains(downSQL, "@REASON = N'COLUMN_NOT_EXISTS'") {
		t.Fatalf("table alter down migration must include drop-column precheck reason")
	}
	if !strings.Contains(downSQL, "C.NAME COLLATE LATIN1_GENERAL_100_BIN2 = N'NOTE'") {
		t.Fatalf("table alter down migration must include drop-column existence precheck")
	}
	if !strings.Contains(downSQL, "N'PRECHECK: TARGET COLUMN(S) ALREADY ABSENT'") {
		t.Fatalf("table alter down migration must include drop-column precheck message")
	}

	second, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("second gen run failed: %v", err)
	}
	if second.Fail.Count != 0 {
		t.Fatalf("second gen expected no fail items, got %d", second.Fail.Count)
	}
	if len(second.NewMigrations.Files) != 0 {
		t.Fatalf("second gen should not emit migration after reconstruction, got %d", len(second.NewMigrations.Files))
	}
}

func TestGenAlterColumnNullToNotNullEmitsPrecheck(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "07_vk_whitespace_only")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	tablePath := filepath.Join(workRoot, "db_object", "schema", "tables", "dbo.TableA.sql")
	tableSQL := `CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL,
  amount DECIMAL(10,2) NOT NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id),
  CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status
);
`
	if err := os.WriteFile(tablePath, []byte(tableSQL), 0o644); err != nil {
		t.Fatalf("write table source failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("engine run failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected no fail items, got %d", report.Fail.Count)
	}
	if hasPassRule(report, "ALTER_COLUMN_REQUIRES_MANUAL") {
		t.Fatalf("NULL->NOT NULL alter should use generated precheck, not manual skip")
	}
	if len(report.NewMigrations.Files) == 0 {
		t.Fatalf("expected generated migration for ALTER COLUMN with precheck")
	}

	tableUp := ""
	for _, name := range report.NewMigrations.Files {
		if strings.Contains(name, "_tables_") {
			tableUp = name
			break
		}
	}
	if tableUp == "" {
		t.Fatalf("expected generated table migration")
	}

	upRaw, err := os.ReadFile(filepath.Join(workRoot, "migrations", tableUp))
	if err != nil {
		t.Fatalf("read generated up migration failed: %v", err)
	}
	upSQL := strings.ToUpper(string(upRaw))
	if !strings.Contains(upSQL, "DECLARE @NULL_CNT_COL_") {
		t.Fatalf("up migration must include NULL->NOT NULL precheck counter")
	}
	if !strings.Contains(upSQL, "N'PRECHECK: ALTER COLUMN NOT NULL BLOCKED'") {
		t.Fatalf("up migration must include ALTER COLUMN NOT NULL precheck message")
	}
	if !strings.Contains(upSQL, "ALTER TABLE DBO.TABLEA ALTER COLUMN [AMOUNT] DECIMAL(10,2) NOT NULL;") {
		t.Fatalf("up migration must apply ALTER COLUMN after precheck")
	}

	second, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("second gen run failed: %v", err)
	}
	if second.Fail.Count != 0 {
		t.Fatalf("second gen expected no fail items, got %d", second.Fail.Count)
	}
	if len(second.NewMigrations.Files) != 0 {
		t.Fatalf("second gen should not emit migration after reconstruction, got %d", len(second.NewMigrations.Files))
	}
}

func TestGenAlterColumnShrinkLengthEmitsPrecheck(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "07_vk_whitespace_only")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	historyUpPath := filepath.Join(workRoot, "migrations", "20260217073015000_tables_dbo.TableA_v1.up.sql")
	historyUpRaw, err := os.ReadFile(historyUpPath)
	if err != nil {
		t.Fatalf("read history table migration failed: %v", err)
	}
	historyUpSQL := strings.Replace(string(historyUpRaw), "amount DECIMAL(10,2) NULL,", "amount DECIMAL(10,2) NULL,\n    note VARCHAR(10) NULL,", 1)
	if err := os.WriteFile(historyUpPath, []byte(historyUpSQL), 0o644); err != nil {
		t.Fatalf("write history table migration failed: %v", err)
	}

	tablePath := filepath.Join(workRoot, "db_object", "schema", "tables", "dbo.TableA.sql")
	tableSQL := `CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL,
  amount DECIMAL(10,2) NULL,
  note VARCHAR(5) NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id),
  CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status
);
`
	if err := os.WriteFile(tablePath, []byte(tableSQL), 0o644); err != nil {
		t.Fatalf("write table source failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("gen run failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected no fail items, got %d", report.Fail.Count)
	}

	tableUp := ""
	for _, name := range report.NewMigrations.Files {
		if strings.Contains(name, "_tables_") {
			tableUp = name
			break
		}
	}
	if tableUp == "" {
		t.Fatalf("expected generated table migration")
	}

	upRaw, err := os.ReadFile(filepath.Join(workRoot, "migrations", tableUp))
	if err != nil {
		t.Fatalf("read generated up migration failed: %v", err)
	}
	upSQL := strings.ToUpper(string(upRaw))
	if !strings.Contains(upSQL, "DECLARE @TOO_LONG_COL_") {
		t.Fatalf("up migration must include shrink-length precheck counter")
	}
	if !strings.Contains(upSQL, "DATALENGTH([NOTE]) > 5;") {
		t.Fatalf("up migration must include shrink-length threshold precheck")
	}
	if !strings.Contains(upSQL, "N'NARROWING_WOULD_TRUNCATE'") {
		t.Fatalf("up migration must log NARROWING_WOULD_TRUNCATE on shrink block")
	}
	if !strings.Contains(upSQL, "ALTER TABLE DBO.TABLEA ALTER COLUMN [NOTE] VARCHAR(5) NULL;") {
		t.Fatalf("up migration must apply shrink ALTER COLUMN after precheck")
	}
}

func TestGenDetectsParameterizedColumnChangesWithWhitespace(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "07_vk_whitespace_only")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	historyUpPath := filepath.Join(workRoot, "migrations", "20260217073015000_tables_dbo.TableA_v1.up.sql")
	historyUpRaw, err := os.ReadFile(historyUpPath)
	if err != nil {
		t.Fatalf("read history table migration failed: %v", err)
	}
	historyUpSQL := strings.Replace(string(historyUpRaw), "amount DECIMAL(10,2) NULL,", `amount [DECIMAL](25, 2) NULL,
    code VARCHAR (10) NULL,
    payload VARBINARY (8) NULL,
    occurred_at DATETIME2 (2) NULL,
    counter INT NULL,`, 1)
	if err := os.WriteFile(historyUpPath, []byte(historyUpSQL), 0o644); err != nil {
		t.Fatalf("write history table migration failed: %v", err)
	}

	tablePath := filepath.Join(workRoot, "db_object", "schema", "tables", "dbo.TableA.sql")
	tableSQL := `CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL,
	  amount [DECIMAL](25, 6) NULL,
  code VARCHAR (24) NULL,
  payload VARBINARY (16) NULL,
  occurred_at DATETIME2 (5) NULL,
  counter BIGINT NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id),
  CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status
);
`
	if err := os.WriteFile(tablePath, []byte(tableSQL), 0o644); err != nil {
		t.Fatalf("write table source failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("gen run failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected no fail items, got: %+v", report.Fail.Items)
	}

	tableUp := ""
	for _, name := range report.NewMigrations.Files {
		if strings.Contains(name, "_tables_") {
			tableUp = name
			break
		}
	}
	if tableUp == "" {
		t.Fatal("expected generated table migration for parameterized type changes")
	}

	upRaw, err := os.ReadFile(filepath.Join(workRoot, "migrations", tableUp))
	if err != nil {
		t.Fatalf("read generated up migration failed: %v", err)
	}
	upSQL := strings.ToUpper(string(upRaw))
	for _, want := range []string{
		"ALTER TABLE DBO.TABLEA ALTER COLUMN [AMOUNT] [DECIMAL](25,6) NULL;",
		"ALTER TABLE DBO.TABLEA ALTER COLUMN [CODE] VARCHAR(24) NULL;",
		"ALTER TABLE DBO.TABLEA ALTER COLUMN [PAYLOAD] VARBINARY(16) NULL;",
		"ALTER TABLE DBO.TABLEA ALTER COLUMN [OCCURRED_AT] DATETIME2(5) NULL;",
		"ALTER TABLE DBO.TABLEA ALTER COLUMN [COUNTER] BIGINT NULL;",
	} {
		if !strings.Contains(upSQL, want) {
			t.Errorf("generated migration missing %q", want)
		}
	}
	for _, want := range []string{
		"TRY_CONVERT([DECIMAL](25,6), [AMOUNT]) IS NULL",
		"TRY_CONVERT([DECIMAL](25,2), TRY_CONVERT([DECIMAL](25,6), [AMOUNT])) <> [AMOUNT]",
	} {
		if !strings.Contains(upSQL, want) {
			t.Errorf("generated bracketed DECIMAL guard missing %q", want)
		}
	}

	second, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("second gen run failed: %v", err)
	}
	if second.Fail.Count != 0 || len(second.NewMigrations.Files) != 0 {
		t.Fatalf("second gen must be idempotent: fail=%d new=%v", second.Fail.Count, second.NewMigrations.Files)
	}
}

func TestGenAlterColumnTypeToNumericEmitsPrecheck(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "07_vk_whitespace_only")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	historyUpPath := filepath.Join(workRoot, "migrations", "20260217073015000_tables_dbo.TableA_v1.up.sql")
	historyUpRaw, err := os.ReadFile(historyUpPath)
	if err != nil {
		t.Fatalf("read history table migration failed: %v", err)
	}
	historyUpSQL := strings.Replace(string(historyUpRaw), "amount DECIMAL(10,2) NULL,", "amount DECIMAL(10,2) NULL,\n    note VARCHAR(10) NULL,", 1)
	if err := os.WriteFile(historyUpPath, []byte(historyUpSQL), 0o644); err != nil {
		t.Fatalf("write history table migration failed: %v", err)
	}

	tablePath := filepath.Join(workRoot, "db_object", "schema", "tables", "dbo.TableA.sql")
	tableSQL := `CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL,
  amount DECIMAL(10,2) NULL,
  note DECIMAL(10,2) NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id),
  CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status
);
`
	if err := os.WriteFile(tablePath, []byte(tableSQL), 0o644); err != nil {
		t.Fatalf("write table source failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("gen run failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected no fail items, got %d", report.Fail.Count)
	}

	tableUp := ""
	for _, name := range report.NewMigrations.Files {
		if strings.Contains(name, "_tables_") {
			tableUp = name
			break
		}
	}
	if tableUp == "" {
		t.Fatalf("expected generated table migration")
	}

	upRaw, err := os.ReadFile(filepath.Join(workRoot, "migrations", tableUp))
	if err != nil {
		t.Fatalf("read generated up migration failed: %v", err)
	}
	upSQL := strings.ToUpper(string(upRaw))
	if !strings.Contains(upSQL, "DECLARE @BAD_COL_") {
		t.Fatalf("up migration must include type-conversion precheck counter")
	}
	if !strings.Contains(upSQL, "TRY_CONVERT(DECIMAL(10,2), [NOTE]) IS NULL") {
		t.Fatalf("up migration must include TRY_CONVERT precheck for numeric target type")
	}
	if !strings.Contains(upSQL, "N'PRECHECK: TYPE CHANGE BLOCKED (NON-CONVERTIBLE)'") {
		t.Fatalf("up migration must include type-change precheck message")
	}
	if !strings.Contains(upSQL, "ALTER TABLE DBO.TABLEA ALTER COLUMN [NOTE] DECIMAL(10,2) NULL;") {
		t.Fatalf("up migration must apply type-change ALTER COLUMN after precheck")
	}
}

func TestGenAlterColumnWidenEmitsTargetLengthGuard(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "07_vk_whitespace_only")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	historyUpPath := filepath.Join(workRoot, "migrations", "20260217073015000_tables_dbo.TableA_v1.up.sql")
	historyUpRaw, err := os.ReadFile(historyUpPath)
	if err != nil {
		t.Fatalf("read history table migration failed: %v", err)
	}
	historyUpSQL := strings.Replace(string(historyUpRaw), "amount DECIMAL(10,2) NULL,", "amount DECIMAL(10,2) NULL,\n    note VARCHAR(10) NULL,", 1)
	if err := os.WriteFile(historyUpPath, []byte(historyUpSQL), 0o644); err != nil {
		t.Fatalf("write history table migration failed: %v", err)
	}

	tablePath := filepath.Join(workRoot, "db_object", "schema", "tables", "dbo.TableA.sql")
	tableSQL := `CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL,
  amount DECIMAL(10,2) NULL,
  note VARCHAR(24) NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id),
  CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status
);
`
	if err := os.WriteFile(tablePath, []byte(tableSQL), 0o644); err != nil {
		t.Fatalf("write table source failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("gen run failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected no fail items, got %d", report.Fail.Count)
	}

	tableUp := ""
	for _, name := range report.NewMigrations.Files {
		if strings.Contains(name, "_tables_") {
			tableUp = name
			break
		}
	}
	if tableUp == "" {
		t.Fatalf("expected generated table migration")
	}

	upRaw, err := os.ReadFile(filepath.Join(workRoot, "migrations", tableUp))
	if err != nil {
		t.Fatalf("read generated up migration failed: %v", err)
	}
	upSQL := strings.ToUpper(string(upRaw))
	if !strings.Contains(upSQL, "DATALENGTH([NOTE]) > 24;") {
		t.Fatalf("up migration must include target-length precheck for widen alter")
	}
	if !strings.Contains(upSQL, "N'DATA_TOO_LONG_FOR_TARGET'") {
		t.Fatalf("up migration must emit DATA_TOO_LONG_FOR_TARGET when target length guard blocks alter")
	}
}

func TestGenAlterColumnPreservesSourceColumnCase(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "07_vk_whitespace_only")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	historyUpPath := filepath.Join(workRoot, "migrations", "20260217073015000_tables_dbo.TableA_v1.up.sql")
	historyUpRaw, err := os.ReadFile(historyUpPath)
	if err != nil {
		t.Fatalf("read history table migration failed: %v", err)
	}
	historyUpSQL := strings.Replace(string(historyUpRaw), "amount DECIMAL(10,2) NULL,", "amount DECIMAL(10,2) NULL,\n    columnBaru3 VARCHAR(10) NULL,", 1)
	if err := os.WriteFile(historyUpPath, []byte(historyUpSQL), 0o644); err != nil {
		t.Fatalf("write history table migration failed: %v", err)
	}

	tablePath := filepath.Join(workRoot, "db_object", "schema", "tables", "dbo.TableA.sql")
	tableSQL := `CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL,
  amount DECIMAL(10,2) NULL,
  columnBaru3 VARCHAR(24) NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id),
  CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status
);
`
	if err := os.WriteFile(tablePath, []byte(tableSQL), 0o644); err != nil {
		t.Fatalf("write table source failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("gen run failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected no fail items, got %d", report.Fail.Count)
	}

	tableUp := ""
	for _, name := range report.NewMigrations.Files {
		if strings.Contains(name, "_tables_") {
			tableUp = name
			break
		}
	}
	if tableUp == "" {
		t.Fatalf("expected generated table migration")
	}

	upRaw, err := os.ReadFile(filepath.Join(workRoot, "migrations", tableUp))
	if err != nil {
		t.Fatalf("read generated up migration failed: %v", err)
	}
	upSQL := string(upRaw)
	if !strings.Contains(upSQL, "ALTER TABLE dbo.TableA ALTER COLUMN [columnBaru3] VARCHAR(24) NULL;") {
		t.Fatalf("generated alter column must preserve source casing for column name")
	}
	if !strings.Contains(upSQL, "@target = N'dbo.TableA.columnBaru3'") {
		t.Fatalf("generated log target must preserve source casing for column name")
	}
}

func TestGenAlterColumnAddsDefaultConstraintOnExistingColumn(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "07_vk_whitespace_only")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	historyUpPath := filepath.Join(workRoot, "migrations", "20260217073015000_tables_dbo.TableA_v1.up.sql")
	historyUpRaw, err := os.ReadFile(historyUpPath)
	if err != nil {
		t.Fatalf("read history table migration failed: %v", err)
	}
	historyUpSQL := strings.Replace(string(historyUpRaw), "amount DECIMAL(10,2) NULL,", "amount DECIMAL(10,2) NULL,\n    note VARCHAR(14) NULL,", 1)
	if err := os.WriteFile(historyUpPath, []byte(historyUpSQL), 0o644); err != nil {
		t.Fatalf("write history table migration failed: %v", err)
	}

	tablePath := filepath.Join(workRoot, "db_object", "schema", "tables", "dbo.TableA.sql")
	tableSQL := `CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL,
  amount DECIMAL(10,2) NULL,
  note VARCHAR(14) NOT NULL,
  CONSTRAINT DF_TableA__note DEFAULT ('-') FOR [note],
  CONSTRAINT PK_TableA PRIMARY KEY (id),
  CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status
);
`
	if err := os.WriteFile(tablePath, []byte(tableSQL), 0o644); err != nil {
		t.Fatalf("write table source failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	first, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("first gen run failed: %v", err)
	}
	if first.Fail.Count != 0 {
		t.Fatalf("first gen expected no fail items, got %d", first.Fail.Count)
	}
	if len(first.NewMigrations.Files) == 0 {
		t.Fatalf("expected generated migration for default-constraint delta")
	}

	tableUp := ""
	for _, name := range first.NewMigrations.Files {
		if strings.Contains(name, "_tables_") {
			tableUp = name
			break
		}
	}
	if tableUp == "" {
		t.Fatalf("expected generated table migration")
	}

	upRaw, err := os.ReadFile(filepath.Join(workRoot, "migrations", tableUp))
	if err != nil {
		t.Fatalf("read generated up migration failed: %v", err)
	}
	upSQL := strings.ToUpper(string(upRaw))
	if !strings.Contains(upSQL, "ALTER TABLE DBO.TABLEA ALTER COLUMN [NOTE] VARCHAR(14) NOT NULL;") {
		t.Fatalf("up migration must alter NOTE column to NOT NULL")
	}
	if !strings.Contains(upSQL, "ALTER TABLE DBO.TABLEA ADD CONSTRAINT [DF_TABLEA__NOTE] DEFAULT ('-') FOR [NOTE];") {
		t.Fatalf("up migration must add named default constraint for existing column")
	}
	if !strings.Contains(upSQL, "@ACTION = N'ADD_DEFAULT'") {
		t.Fatalf("up migration must log add_default action")
	}

	second, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("second gen run failed: %v", err)
	}
	if second.Fail.Count != 0 {
		t.Fatalf("second gen expected no fail items, got %d", second.Fail.Count)
	}
	if len(second.NewMigrations.Files) != 0 {
		t.Fatalf("second gen should not re-emit after default-constraint reconstruction")
	}
}

func TestGenMixedDropAddTableChangeAutoPlanned(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "07_vk_whitespace_only")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	tablePath := filepath.Join(workRoot, "db_object", "tables", "dbo.TableA.sql")
	if _, err := os.Stat(tablePath); err != nil {
		tablePath = filepath.Join(workRoot, "db_object", "schema", "tables", "dbo.TableA.sql")
	}
	raw, err := os.ReadFile(tablePath)
	if err != nil {
		t.Fatalf("read table source failed: %v", err)
	}

	// Mixed evolution in one change set: drop amount + add note.
	updated := strings.Replace(string(raw), "amount DECIMAL(10,2) NULL,", "note VARCHAR(20) NULL,", 1)
	if err := os.WriteFile(tablePath, []byte(updated), 0o644); err != nil {
		t.Fatalf("write table source failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("gen run failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected no fail items, got %d", report.Fail.Count)
	}
	if len(report.NewMigrations.Files) == 0 {
		t.Fatalf("mixed drop+add table change should auto-emit migration")
	}
	if !hasPassRule(report, "GEN_PRECHECK_DIFF_DETECTED") {
		t.Fatalf("expected GEN_PRECHECK_DIFF_DETECTED summary in gen precheck mode")
	}
	if hasPassRule(report, "MIXED_TABLE_CHANGE_REQUIRES_MANUAL") {
		t.Fatalf("mixed drop+add table change should not be manual-only")
	}
	if hasPassRule(report, "REPO_STATE_DRIFT") {
		t.Fatalf("REPO_STATE_DRIFT should be suppressed in gen precheck mode")
	}

	tableUp := ""
	for _, name := range report.NewMigrations.Files {
		if strings.Contains(name, "_tables_") {
			tableUp = name
			break
		}
	}
	if tableUp == "" {
		t.Fatalf("expected generated table migration")
	}

	upRaw, err := os.ReadFile(filepath.Join(workRoot, "migrations", tableUp))
	if err != nil {
		t.Fatalf("read generated up migration failed: %v", err)
	}
	upSQL := strings.ToUpper(string(upRaw))
	if !strings.Contains(upSQL, "ALTER TABLE DBO.TABLEA DROP COLUMN [AMOUNT];") {
		t.Fatalf("up migration must include DROP COLUMN for removed column")
	}
	if !strings.Contains(upSQL, "ALTER TABLE DBO.TABLEA ADD [NOTE] VARCHAR(20) NULL;") {
		t.Fatalf("up migration must include ADD COLUMN for new column")
	}
	if !strings.Contains(upSQL, "C.NAME COLLATE LATIN1_GENERAL_100_BIN2 = N'AMOUNT'") {
		t.Fatalf("up migration must guard DROP COLUMN per statement")
	}
	if !strings.Contains(upSQL, "C.NAME COLLATE LATIN1_GENERAL_100_BIN2 = N'NOTE'") {
		t.Fatalf("up migration must guard ADD COLUMN per statement")
	}
	if !strings.Contains(upSQL, "-- PRECHECK IS HANDLED PER STATEMENT IN THE MIGRATION BODY") {
		t.Fatalf("up migration must keep wrapper precheck disabled for mixed column lifecycle")
	}

	tableDown := strings.TrimSuffix(tableUp, ".up.sql") + ".down.sql"
	downRaw, err := os.ReadFile(filepath.Join(workRoot, "migrations", tableDown))
	if err != nil {
		t.Fatalf("read generated down migration failed: %v", err)
	}
	downSQL := strings.ToUpper(string(downRaw))
	if !strings.Contains(downSQL, "ALTER TABLE DBO.TABLEA DROP COLUMN [NOTE];") {
		t.Fatalf("down migration must drop added column first")
	}
	if !strings.Contains(downSQL, "ALTER TABLE DBO.TABLEA ADD [AMOUNT] DECIMAL(10,2) NULL;") {
		t.Fatalf("down migration must restore removed column")
	}
	if !strings.Contains(downSQL, "C.NAME COLLATE LATIN1_GENERAL_100_BIN2 = N'NOTE'") {
		t.Fatalf("down migration must guard DROP COLUMN per statement")
	}
	if !strings.Contains(downSQL, "C.NAME COLLATE LATIN1_GENERAL_100_BIN2 = N'AMOUNT'") {
		t.Fatalf("down migration must guard ADD COLUMN per statement")
	}

	second, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("second gen run failed: %v", err)
	}
	if second.Fail.Count != 0 {
		t.Fatalf("second gen expected no fail items, got %d", second.Fail.Count)
	}
	if len(second.NewMigrations.Files) != 0 {
		t.Fatalf("second gen should not re-emit migration after mixed drop+add applied")
	}
	if !hasPassRule(second, "GEN_NO_CHANGES") {
		t.Fatalf("expected GEN_NO_CHANGES on second run")
	}
}

func parseCategoryFromMigrationUpFile(name string) (string, error) {
	base := filepath.Base(name)
	base = strings.TrimSuffix(base, ".up.sql")
	parts := strings.Split(base, "_")
	if len(parts) < 4 {
		return "", os.ErrInvalid
	}
	if parts[1] == "patch" {
		return "patch", nil
	}
	return parts[1], nil
}

func hasPassReason(report output.Report, reason string) bool {
	for _, item := range report.Pass.Items {
		if item.Reason == reason {
			return true
		}
	}
	return false
}

func hasPassRule(report output.Report, rule string) bool {
	for _, item := range report.Pass.Items {
		if item.Rule == rule {
			return true
		}
	}
	return false
}
