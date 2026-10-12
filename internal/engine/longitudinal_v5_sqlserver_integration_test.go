package engine

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/output"
)

func TestIntegration_LongitudinalV5ArtifactsReconstructDesiredState(t *testing.T) {
	dir := strings.TrimSpace(os.Getenv("RZA_V5_MIGRATIONS_DIR"))
	if dir == "" {
		t.Skip("RZA_V5_MIGRATIONS_DIR is required")
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	paths := make([]string, 0, 5)
	for _, entry := range entries {
		if !entry.IsDir() && strings.Contains(entry.Name(), "_tables_dbo.RzaV5Order_v") && strings.HasSuffix(entry.Name(), ".up.sql") {
			paths = append(paths, filepath.Join(dir, entry.Name()))
		}
	}
	sort.Strings(paths)
	if len(paths) != 5 {
		t.Fatalf("expected five RzaV5Order up migrations, got %d", len(paths))
	}
	meta := migrationMeta{Category: "tables", Object: "dbo.RzaV5Order"}
	state := ""
	for i, path := range paths {
		raw, readErr := os.ReadFile(path)
		if readErr != nil {
			t.Fatal(readErr)
		}
		state, err = semanticFromMigrationWithState(meta, string(raw), state, i > 0)
		if err != nil {
			t.Fatalf("reconstruct %s: %v", filepath.Base(path), err)
		}
	}
	desiredPath := filepath.Join(filepath.Dir(dir), "db_object", "tables", "dbo.RzaV5Order.sql")
	desiredRaw, err := os.ReadFile(desiredPath)
	if err != nil {
		t.Fatal(err)
	}
	desired, err := semanticFromSource(meta, string(desiredRaw))
	if err != nil {
		t.Fatal(err)
	}
	if hashSemanticByCategory("tables", state) != hashSemanticByCategory("tables", desired) {
		t.Fatalf("v1-v5 reconstruction differs from desired\nreconstructed:\n%s\ndesired:\n%s", state, desired)
	}
	engine := New(config.Default(filepath.Dir(dir)))
	report := output.NewReport(engine.cfg.OutputNamingPattern, engine.cfg.OutputNamingLineEndings)
	history, _, _, err := engine.reconstructHistory(&report)
	if err != nil || report.Fail.Count != 0 {
		t.Fatalf("full history reconstruction failed: err=%v fail=%+v", err, report.Fail.Items)
	}
	fullState := history[objectKey("tables", "dbo.RzaV5Order")].SemanticSQL
	if hashSemanticByCategory("tables", fullState) != hashSemanticByCategory("tables", desired) {
		t.Fatalf("full-history v1-v5 reconstruction differs from desired\nreconstructed:\n%s\ndesired:\n%s", fullState, desired)
	}
}

// TestIntegration_SQLServerLongitudinalV1ToV5 executes an externally generated
// RzaV5 migration family without touching schema_migrations. The outer
// transaction is always rolled back, so this is safe for an explicitly
// configured test catalog. The generated core audit migration is installed in
// the same transaction before the object migrations are executed.
func TestIntegration_SQLServerLongitudinalV1ToV5(t *testing.T) {
	dsn := strings.TrimSpace(os.Getenv("RZA_MSSQL_TEST_DSN"))
	dir := strings.TrimSpace(os.Getenv("RZA_V5_MIGRATIONS_DIR"))
	if dsn == "" || dir == "" {
		t.Skip("RZA_MSSQL_TEST_DSN and RZA_V5_MIGRATIONS_DIR are required")
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	up, down := make([]string, 0, 30), make([]string, 0, 30)
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.Contains(name, "_dbo.RzaV5") {
			continue
		}
		switch {
		case strings.HasSuffix(name, ".up.sql"):
			up = append(up, filepath.Join(dir, name))
		case strings.HasSuffix(name, ".down.sql"):
			down = append(down, filepath.Join(dir, name))
		}
	}
	sort.Strings(up)
	sort.Sort(sort.Reverse(sort.StringSlice(down)))
	if len(up) != 30 || len(down) != 30 {
		t.Fatalf("expected 30 up and 30 down RzaV5 files, got up=%d down=%d", len(up), len(down))
	}

	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	coreAuditPath := filepath.Join(dir, coreAuditUpName)
	coreAuditSQL, err := os.ReadFile(coreAuditPath)
	if err != nil {
		t.Fatalf("read core audit migration: %v", err)
	}
	if _, err := tx.ExecContext(ctx, string(coreAuditSQL)); err != nil {
		t.Fatalf("install core audit migration: %v", err)
	}

	execFiles := func(paths []string) {
		t.Helper()
		for _, path := range paths {
			raw, readErr := os.ReadFile(path)
			if readErr != nil {
				t.Fatal(readErr)
			}
			if _, execErr := tx.ExecContext(ctx, string(raw)); execErr != nil {
				t.Fatalf("execute %s: %v", filepath.Base(path), execErr)
			}
		}
	}
	execFiles(up)

	var objects, priorityColumns, indexes, constraints int
	if err := tx.QueryRowContext(ctx, `
SELECT
  (SELECT COUNT(*) FROM sys.objects WHERE object_id IN (
    OBJECT_ID(N'dbo.RzaV5Order'), OBJECT_ID(N'dbo.RzaV5OrderSummary'),
    OBJECT_ID(N'dbo.RzaV5OrderTotal'), OBJECT_ID(N'dbo.RzaV5UpsertOrder'))),
  (SELECT COUNT(*) FROM sys.columns WHERE object_id=OBJECT_ID(N'dbo.RzaV5Order') AND name=N'Priority'),
  (SELECT COUNT(*) FROM sys.indexes WHERE object_id=OBJECT_ID(N'dbo.RzaV5Order') AND name IN (N'IX_RzaV5Order_Customer_Status',N'IX_RzaV5Order_Ready',N'UX_RzaV5Order_ActiveExternalRef')),
  (SELECT COUNT(*) FROM sys.objects WHERE parent_object_id=OBJECT_ID(N'dbo.RzaV5Order') AND name IN (N'DF_RzaV5Order_Amount',N'DF_RzaV5Order_Status',N'UQ_RzaV5Order_ExternalRef',N'CK_RzaV5Order_Priority'));
`).Scan(&objects, &priorityColumns, &indexes, &constraints); err != nil {
		t.Fatal(err)
	}
	if objects != 4 || priorityColumns != 1 || indexes != 3 || constraints != 4 {
		t.Fatalf("unexpected v5 catalog state: objects=%d priority=%d indexes=%d constraints=%d", objects, priorityColumns, indexes, constraints)
	}
	var typeExists int
	if err := tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM sys.table_types tt JOIN sys.schemas s ON s.schema_id=tt.schema_id WHERE s.name=N'dbo' AND tt.name=N'RzaV5IdList';`).Scan(&typeExists); err != nil {
		t.Fatal(err)
	}
	if typeExists != 1 {
		t.Fatalf("RzaV5IdList missing after v5 up")
	}

	execFiles(down)
	var remaining int
	if err := tx.QueryRowContext(ctx, `
SELECT
  (SELECT COUNT(*) FROM sys.objects WHERE object_id IN (
    OBJECT_ID(N'dbo.RzaV5Order'), OBJECT_ID(N'dbo.RzaV5OrderSummary'),
    OBJECT_ID(N'dbo.RzaV5OrderTotal'), OBJECT_ID(N'dbo.RzaV5UpsertOrder')))
  + (SELECT COUNT(*) FROM sys.table_types tt JOIN sys.schemas s ON s.schema_id=tt.schema_id WHERE s.name=N'dbo' AND tt.name=N'RzaV5IdList');
`).Scan(&remaining); err != nil {
		t.Fatal(err)
	}
	if remaining != 0 {
		t.Fatalf("RzaV5 objects remain after complete down: %d", remaining)
	}
}
