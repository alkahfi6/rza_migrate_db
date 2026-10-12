package engine

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rza_migrate_db/internal/config"
)

func TestRunFmtAppliesSharedSafeFixesAndNeverTouchesMigrations(t *testing.T) {
	repo := t.TempDir()
	tablePath := filepath.Join(repo, "db_object", "tables", "dbo.Outbox.sql")
	procPath := filepath.Join(repo, "db_object", "procedures", "dbo.Publish.sql")
	patchPath := filepath.Join(repo, "patch", "001_access", "up.sql")
	indexPath := filepath.Join(repo, "db_object", "indexes", "dbo.Wrong.sql")
	migrationPath := filepath.Join(repo, "migrations", "20260814000000000_patch_guard.up.sql")
	for _, path := range []string{tablePath, procPath, patchPath, indexPath, migrationPath} {
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	table := "CREATE TABLE dbo.Outbox (\n  id INT NOT NULL,\n  retry_count INT NOT NULL DEFAULT 0,\n  PRIMARY KEY (id),\n);\nGO"
	proc := "USE SQL_TEST;\nGO\n/* banner */\nSET ANSI_NULLS ON\nGO\nALTER PROC Publish AS SELECT 1;"
	patch := "USE SQL_TEST;\nGO\nCREATE PROC dbo.LegacyPatch AS SELECT 1;"
	index := "CREATE INDEX IX_Outbox_retry ON dbo.Outbox(retry_count);\n"
	migration := "USE SQL_MUST_REMAIN;\nGO\nCREATE PROC dbo.Sealed AS SELECT 1;\n"
	for path, content := range map[string]string{tablePath: table, procPath: proc, patchPath: patch, indexPath: index, migrationPath: migration} {
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	manifest := "catalog:\n  self: SQL_TEST\n  collation: Latin1_General_BIN\n"
	if err := os.WriteFile(filepath.Join(repo, "rza_migrate.yaml"), []byte(manifest), 0o644); err != nil {
		t.Fatal(err)
	}

	report, err := New(config.Default(repo)).Run("fmt")
	if err != nil {
		t.Fatal(err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("fmt failed: %+v", report.Fail.Items)
	}
	assertFileContains(t, tablePath, "CONSTRAINT DF_Outbox__retry_count DEFAULT 0")
	assertFileContains(t, tablePath, "CONSTRAINT PK_Outbox PRIMARY KEY (id)")
	assertFileNotContains(t, tablePath, "\nGO")
	assertFileContains(t, procPath, "/* banner */\nCREATE OR ALTER PROCEDURE dbo.Publish")
	assertFileContains(t, patchPath, "CREATE PROCEDURE dbo.LegacyPatch")
	assertFileNotContains(t, patchPath, "USE SQL_TEST")
	if _, err := os.Stat(filepath.Join(filepath.Dir(indexPath), "dbo.Outbox.sql")); err != nil {
		t.Fatalf("index source was not safely renamed: %v", err)
	}
	gotMigration, err := os.ReadFile(migrationPath)
	if err != nil {
		t.Fatal(err)
	}
	if string(gotMigration) != migration {
		t.Fatalf("sealed migration changed:\n%s", gotMigration)
	}
	temps, err := filepath.Glob(filepath.Join(repo, "**", ".rza-sqlfix-*.tmp"))
	if err != nil {
		t.Fatal(err)
	}
	if len(temps) > 0 {
		t.Fatalf("atomic write temp files leaked: %v", temps)
	}
}

func TestRunFmtCheckReportsAllSourcesWithoutWriting(t *testing.T) {
	repo := t.TempDir()
	tablePath := filepath.Join(repo, "db_object", "tables", "dbo.T.sql")
	patchPath := filepath.Join(repo, "patch", "001", "up.sql")
	for _, path := range []string{tablePath, patchPath} {
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	table := "CREATE TABLE dbo.T (id INT DEFAULT 0,);"
	patch := "USE SQL_TEST;\nGO\nSELECT 1;"
	if err := os.WriteFile(tablePath, []byte(table), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(patchPath, []byte(patch), 0o644); err != nil {
		t.Fatal(err)
	}

	cfg := config.Default(repo)
	cfg.EnableFmtCheck = true
	report, err := New(cfg).Run("fmt")
	if err != nil {
		t.Fatal(err)
	}
	if !hasRule(report.Fail.Items, "FMT_CHECK_CHANGES_REQUIRED") {
		t.Fatalf("missing check failure: %+v", report.Fail.Items)
	}
	for path, want := range map[string]string{tablePath: table, patchPath: patch} {
		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != want {
			t.Fatalf("check mode wrote %s", path)
		}
	}
}

func assertFileContains(t *testing.T, path, want string) {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), want) {
		t.Fatalf("%s missing %q:\n%s", path, want, raw)
	}
}

func assertFileNotContains(t *testing.T, path, unwanted string) {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(raw), unwanted) {
		t.Fatalf("%s unexpectedly contains %q:\n%s", path, unwanted, raw)
	}
}
