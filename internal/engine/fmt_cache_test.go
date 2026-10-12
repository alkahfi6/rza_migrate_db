package engine

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/output"
)

func TestRunFmtUsesCacheForUnchangedFiles(t *testing.T) {
	t.Parallel()

	repo := t.TempDir()
	tableDir := filepath.Join(repo, "db_object", "tables")
	if err := os.MkdirAll(tableDir, 0o755); err != nil {
		t.Fatalf("mkdir table dir: %v", err)
	}
	tablePath := filepath.Join(tableDir, "dbo.TableA.sql")
	sql := "CREATE TABLE dbo.TableA(\n  id INT NOT NULL,\nCONSTRAINT PK_TableA PRIMARY KEY (id)\n)\n"
	if err := os.WriteFile(tablePath, []byte(sql), 0o644); err != nil {
		t.Fatalf("write table: %v", err)
	}

	e := New(config.Default(repo))

	first := output.NewReport(e.cfg.OutputNamingPattern, e.cfg.OutputNamingLineEndings)
	if err := e.runFmt(&first); err != nil {
		t.Fatalf("first runFmt error: %v", err)
	}
	if first.Fail.Count != 0 {
		t.Fatalf("first runFmt unexpected fail count=%d", first.Fail.Count)
	}
	if got := noteInt(first.Notes, "fmt_cache_hits"); got != 0 {
		t.Fatalf("first fmt_cache_hits mismatch: got=%d want=0", got)
	}

	second := output.NewReport(e.cfg.OutputNamingPattern, e.cfg.OutputNamingLineEndings)
	if err := e.runFmt(&second); err != nil {
		t.Fatalf("second runFmt error: %v", err)
	}
	if second.Fail.Count != 0 {
		t.Fatalf("second runFmt unexpected fail count=%d", second.Fail.Count)
	}
	if got := noteInt(second.Notes, "fmt_cache_hits"); got != 1 {
		t.Fatalf("second fmt_cache_hits mismatch: got=%d want=1", got)
	}
	if got := noteInt(second.Notes, "fmt_cache_misses"); got != 0 {
		t.Fatalf("second fmt_cache_misses mismatch: got=%d want=0", got)
	}

	if _, err := os.Stat(filepath.Join(repo, "out", "fmt_cache.json")); err != nil {
		t.Fatalf("expected fmt cache file to exist: %v", err)
	}
}

func TestRunFmtCacheMissWhenFileChanged(t *testing.T) {
	t.Parallel()

	repo := t.TempDir()
	tableDir := filepath.Join(repo, "db_object", "tables")
	if err := os.MkdirAll(tableDir, 0o755); err != nil {
		t.Fatalf("mkdir table dir: %v", err)
	}
	tablePath := filepath.Join(tableDir, "dbo.TableA.sql")
	sql := "CREATE TABLE dbo.TableA(\n  id INT NOT NULL,\n  CONSTRAINT PK_TableA PRIMARY KEY (id)\n)\n"
	if err := os.WriteFile(tablePath, []byte(sql), 0o644); err != nil {
		t.Fatalf("write table: %v", err)
	}

	e := New(config.Default(repo))
	first := output.NewReport(e.cfg.OutputNamingPattern, e.cfg.OutputNamingLineEndings)
	if err := e.runFmt(&first); err != nil {
		t.Fatalf("first runFmt error: %v", err)
	}

	changed := "CREATE TABLE dbo.TableA(\n  id INT NOT NULL,\n  name VARCHAR(50) NULL,\n  CONSTRAINT PK_TableA PRIMARY KEY (id)\n)\n"
	if err := os.WriteFile(tablePath, []byte(changed), 0o644); err != nil {
		t.Fatalf("rewrite table: %v", err)
	}

	second := output.NewReport(e.cfg.OutputNamingPattern, e.cfg.OutputNamingLineEndings)
	if err := e.runFmt(&second); err != nil {
		t.Fatalf("second runFmt error: %v", err)
	}
	if got := noteInt(second.Notes, "fmt_cache_hits"); got != 0 {
		t.Fatalf("fmt_cache_hits mismatch after content change: got=%d want=0", got)
	}
	if got := noteInt(second.Notes, "fmt_cache_misses"); got != 1 {
		t.Fatalf("fmt_cache_misses mismatch after content change: got=%d want=1", got)
	}
}

func TestRunFmtCorruptCacheNonFatal(t *testing.T) {
	t.Parallel()

	repo := t.TempDir()
	tableDir := filepath.Join(repo, "db_object", "tables")
	if err := os.MkdirAll(tableDir, 0o755); err != nil {
		t.Fatalf("mkdir table dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(tableDir, "dbo.TableA.sql"), []byte("CREATE TABLE dbo.TableA(id INT NOT NULL)\n"), 0o644); err != nil {
		t.Fatalf("write table: %v", err)
	}
	outDir := filepath.Join(repo, "out")
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		t.Fatalf("mkdir out dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(outDir, "fmt_cache.json"), []byte("{invalid-json"), 0o644); err != nil {
		t.Fatalf("write corrupt cache: %v", err)
	}

	e := New(config.Default(repo))
	report := output.NewReport(e.cfg.OutputNamingPattern, e.cfg.OutputNamingLineEndings)
	if err := e.runFmt(&report); err != nil {
		t.Fatalf("runFmt error with corrupt cache: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("unexpected fail count=%d", report.Fail.Count)
	}
	if _, ok := report.Notes["fmt_cache_reset_reason"]; !ok {
		t.Fatalf("expected fmt_cache_reset_reason note when cache is corrupt")
	}
	if saved, ok := report.Notes["fmt_cache_saved"].(bool); !ok || !saved {
		t.Fatalf("expected fmt_cache_saved=true after successful run, got=%v", report.Notes["fmt_cache_saved"])
	}
}

func TestRunFmtNormalizesProgrammabilityHeaderHabits(t *testing.T) {
	t.Parallel()

	repo := t.TempDir()
	procPath := filepath.Join(repo, "db_object", "procedures", "dbo.ProcA.sql")
	if err := os.MkdirAll(filepath.Dir(procPath), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	if err := os.WriteFile(procPath, []byte("create or alter proc ProcA\nas\nselect 1\n"), 0o644); err != nil {
		t.Fatalf("write proc failed: %v", err)
	}

	cfg := config.Default(repo)
	e := New(cfg)
	report, err := e.Run("fmt")
	if err != nil {
		t.Fatalf("fmt run failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected fmt pass, got %+v", report.Fail.Items)
	}
	got, err := os.ReadFile(procPath)
	if err != nil {
		t.Fatalf("read proc failed: %v", err)
	}
	if !strings.HasPrefix(string(got), "CREATE OR ALTER PROCEDURE dbo.ProcA") {
		t.Fatalf("expected canonical procedure header, got: %s", string(got))
	}
}

func TestRunFmtPreservesCommentAboveTableColumn(t *testing.T) {
	t.Parallel()

	repo := t.TempDir()
	tablePath := filepath.Join(repo, "db_object", "tables", "dbo.TableA.sql")
	if err := os.MkdirAll(filepath.Dir(tablePath), 0o755); err != nil {
		t.Fatal(err)
	}
	// Trailing whitespace on a comment is ordinary text normalization, not a
	// change to the comment's indentation or inline placement.
	original := "CREATE TABLE dbo.TableA (\n  -- field info   \n  fieldA VARCHAR(100) NULL\n);\n"
	if err := os.WriteFile(tablePath, []byte(original), 0o644); err != nil {
		t.Fatal(err)
	}

	cfg := config.Default(repo)
	e := New(cfg)
	report := output.NewReport(cfg.OutputNamingPattern, cfg.OutputNamingLineEndings)
	if err := e.runFmt(&report); err != nil {
		t.Fatal(err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("comment above a table column must not fail formatting: %#v", report.Fail.Items)
	}

	after, err := os.ReadFile(tablePath)
	if err != nil {
		t.Fatal(err)
	}
	want := "CREATE TABLE dbo.TableA (\n  -- field info\n  fieldA VARCHAR(100) NULL\n);\n"
	if string(after) != want {
		t.Fatalf("comment placement changed\nwant:\n%s\ngot:\n%s", want, after)
	}
}

func TestRunFmtCheckReportsWithoutWriting(t *testing.T) {
	t.Parallel()
	repo := t.TempDir()
	tablePath := filepath.Join(repo, "db_object", "tables", "dbo.TableA.sql")
	if err := os.MkdirAll(filepath.Dir(tablePath), 0o755); err != nil {
		t.Fatal(err)
	}
	original := "CREATE TABLE dbo.TableA(\nId INT NOT NULL,\nCONSTRAINT PK_TableA PRIMARY KEY (Id)\n)\n"
	if err := os.WriteFile(tablePath, []byte(original), 0o644); err != nil {
		t.Fatal(err)
	}
	cfg := config.Default(repo)
	cfg.EnableFmtCheck = true
	e := New(cfg)
	report := output.NewReport(cfg.OutputNamingPattern, cfg.OutputNamingLineEndings)
	if err := e.runFmt(&report); err != nil {
		t.Fatal(err)
	}
	if report.Fail.Count == 0 || report.Fail.Items[len(report.Fail.Items)-1].Rule != "FMT_CHECK_CHANGES_REQUIRED" {
		t.Fatalf("expected FMT_CHECK_CHANGES_REQUIRED, got %#v", report.Fail.Items)
	}
	after, err := os.ReadFile(tablePath)
	if err != nil {
		t.Fatal(err)
	}
	if string(after) != original {
		t.Fatalf("fmt -check wrote db_object\nwant:\n%s\ngot:\n%s", original, after)
	}
}

func noteInt(notes map[string]interface{}, key string) int {
	raw, ok := notes[key]
	if !ok || raw == nil {
		return 0
	}
	switch v := raw.(type) {
	case int:
		return v
	case int64:
		return int(v)
	case float64:
		return int(v)
	default:
		return 0
	}
}
