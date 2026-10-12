package engine

import (
	"path/filepath"
	"reflect"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/output"
)

func TestReconstructHistoryCacheHitAndInvalidation(t *testing.T) {
	t.Parallel()

	repo := t.TempDir()
	mustWriteFile(t, filepath.Join(repo, "migrations", "20260217073015000_tables_dbo.TableA_v1.up.sql"), "CREATE TABLE dbo.TableA (\n  id INT NOT NULL,\n  CONSTRAINT PK_TableA PRIMARY KEY (id)\n);\n")
	mustWriteFile(t, filepath.Join(repo, "migrations", "20260217073015000_tables_dbo.TableA_v1.down.sql"), "DROP TABLE dbo.TableA;\n")

	e := New(config.Default(repo))

	report1 := output.NewReport(e.cfg.OutputNamingPattern, e.cfg.OutputNamingLineEndings)
	h1, ef1, mv1, err := e.reconstructHistory(&report1)
	if err != nil {
		t.Fatalf("first reconstructHistory error: %v", err)
	}
	if report1.Fail.Count > 0 {
		t.Fatalf("first reconstructHistory fail items: %#v", report1.Fail.Items)
	}
	if cacheHit, ok := noteBool(report1, "reconstruct_cache_hit"); !ok || cacheHit {
		t.Fatalf("expected first reconstruct cache miss, got hit=%v ok=%v", cacheHit, ok)
	}

	report2 := output.NewReport(e.cfg.OutputNamingPattern, e.cfg.OutputNamingLineEndings)
	h2, ef2, mv2, err := e.reconstructHistory(&report2)
	if err != nil {
		t.Fatalf("second reconstructHistory error: %v", err)
	}
	if report2.Fail.Count > 0 {
		t.Fatalf("second reconstructHistory fail items: %#v", report2.Fail.Items)
	}
	if cacheHit, ok := noteBool(report2, "reconstruct_cache_hit"); !ok || !cacheHit {
		t.Fatalf("expected second reconstruct cache hit, got hit=%v ok=%v", cacheHit, ok)
	}
	if !reflect.DeepEqual(h1, h2) || !reflect.DeepEqual(ef1, ef2) || mv1 != mv2 {
		t.Fatalf("cached reconstruct result mismatch")
	}

	mustWriteFile(t, filepath.Join(repo, "migrations", "20260217073015001_views_dbo.ViewA_v1.up.sql"), "CREATE OR ALTER VIEW dbo.ViewA AS SELECT 1 AS id;\n")
	mustWriteFile(t, filepath.Join(repo, "migrations", "20260217073015001_views_dbo.ViewA_v1.down.sql"), "CREATE OR ALTER VIEW dbo.ViewA AS SELECT 1 AS id;\n")
	report3 := output.NewReport(e.cfg.OutputNamingPattern, e.cfg.OutputNamingLineEndings)
	_, _, _, err = e.reconstructHistory(&report3)
	if err != nil {
		t.Fatalf("third reconstructHistory error: %v", err)
	}
	if cacheHit, ok := noteBool(report3, "reconstruct_cache_hit"); !ok || cacheHit {
		t.Fatalf("expected reconstruct cache miss after migration change, got hit=%v ok=%v", cacheHit, ok)
	}
}

func TestCollectDesiredObjectsCacheHitAndInvalidation(t *testing.T) {
	t.Parallel()

	repo := t.TempDir()
	mustWriteFile(t, filepath.Join(repo, "db_object", "tables", "dbo.TableA.sql"), "CREATE TABLE dbo.TableA (\n  id INT NOT NULL,\n  CONSTRAINT PK_TableA PRIMARY KEY (id)\n);\n")
	mustWriteFile(t, filepath.Join(repo, "db_object", "indexes", "dbo.TableA.sql"), "CREATE INDEX IX_TableA__id ON dbo.TableA(id);\n")
	mustWriteFile(t, filepath.Join(repo, "db_object", "views", "dbo.ViewA.sql"), "CREATE OR ALTER VIEW dbo.ViewA AS SELECT id FROM dbo.TableA;\n")

	e := New(config.Default(repo))

	report1 := output.NewReport(e.cfg.OutputNamingPattern, e.cfg.OutputNamingLineEndings)
	d1, err := e.collectDesiredObjects(&report1)
	if err != nil {
		t.Fatalf("first collectDesiredObjects error: %v", err)
	}
	if report1.Fail.Count > 0 {
		t.Fatalf("first collectDesiredObjects fail items: %#v", report1.Fail.Items)
	}
	if cacheHit, ok := noteBool(report1, "desired_collect_cache_hit"); !ok || cacheHit {
		t.Fatalf("expected first desired cache miss, got hit=%v ok=%v", cacheHit, ok)
	}

	report2 := output.NewReport(e.cfg.OutputNamingPattern, e.cfg.OutputNamingLineEndings)
	d2, err := e.collectDesiredObjects(&report2)
	if err != nil {
		t.Fatalf("second collectDesiredObjects error: %v", err)
	}
	if report2.Fail.Count > 0 {
		t.Fatalf("second collectDesiredObjects fail items: %#v", report2.Fail.Items)
	}
	if cacheHit, ok := noteBool(report2, "desired_collect_cache_hit"); !ok || !cacheHit {
		t.Fatalf("expected second desired cache hit, got hit=%v ok=%v", cacheHit, ok)
	}
	if !reflect.DeepEqual(d1, d2) {
		t.Fatalf("cached desired objects mismatch")
	}

	mustWriteFile(t, filepath.Join(repo, "db_object", "views", "dbo.ViewA.sql"), "CREATE OR ALTER VIEW dbo.ViewA AS SELECT id, 1 AS extra_col FROM dbo.TableA;\n")
	report3 := output.NewReport(e.cfg.OutputNamingPattern, e.cfg.OutputNamingLineEndings)
	_, err = e.collectDesiredObjects(&report3)
	if err != nil {
		t.Fatalf("third collectDesiredObjects error: %v", err)
	}
	if cacheHit, ok := noteBool(report3, "desired_collect_cache_hit"); !ok || cacheHit {
		t.Fatalf("expected desired cache miss after db_object change, got hit=%v ok=%v", cacheHit, ok)
	}
}

func TestCollectPatchSourcesCacheHitAndInvalidation(t *testing.T) {
	t.Parallel()

	repo := t.TempDir()
	mustWriteFile(t, filepath.Join(repo, "patch", "Fix Status", "README.md"), "# patch\n")
	mustWriteFile(t, filepath.Join(repo, "patch", "Fix Status", "up.sql"), "UPDATE dbo.TableA SET id = id;\n")
	mustWriteFile(t, filepath.Join(repo, "patch", "Fix Status", "down.sql"), "UPDATE dbo.TableA SET id = id;\n")

	e := New(config.Default(repo))

	report1 := output.NewReport(e.cfg.OutputNamingPattern, e.cfg.OutputNamingLineEndings)
	p1, err := e.collectPatchSources(&report1)
	if err != nil {
		t.Fatalf("first collectPatchSources error: %v", err)
	}
	if report1.Fail.Count > 0 {
		t.Fatalf("first collectPatchSources fail items: %#v", report1.Fail.Items)
	}
	if cacheHit, ok := noteBool(report1, "patch_collect_cache_hit"); !ok || cacheHit {
		t.Fatalf("expected first patch cache miss, got hit=%v ok=%v", cacheHit, ok)
	}

	report2 := output.NewReport(e.cfg.OutputNamingPattern, e.cfg.OutputNamingLineEndings)
	p2, err := e.collectPatchSources(&report2)
	if err != nil {
		t.Fatalf("second collectPatchSources error: %v", err)
	}
	if report2.Fail.Count > 0 {
		t.Fatalf("second collectPatchSources fail items: %#v", report2.Fail.Items)
	}
	if cacheHit, ok := noteBool(report2, "patch_collect_cache_hit"); !ok || !cacheHit {
		t.Fatalf("expected second patch cache hit, got hit=%v ok=%v", cacheHit, ok)
	}
	if !reflect.DeepEqual(p1, p2) {
		t.Fatalf("cached patch sources mismatch")
	}

	mustWriteFile(t, filepath.Join(repo, "patch", "Fix Status", "up.sql"), "UPDATE dbo.TableA SET id = id + 1;\n")
	report3 := output.NewReport(e.cfg.OutputNamingPattern, e.cfg.OutputNamingLineEndings)
	_, err = e.collectPatchSources(&report3)
	if err != nil {
		t.Fatalf("third collectPatchSources error: %v", err)
	}
	if cacheHit, ok := noteBool(report3, "patch_collect_cache_hit"); !ok || cacheHit {
		t.Fatalf("expected patch cache miss after patch file change, got hit=%v ok=%v", cacheHit, ok)
	}
}

func noteBool(report output.Report, key string) (bool, bool) {
	v, ok := report.Notes[key]
	if !ok {
		return false, false
	}
	b, ok := v.(bool)
	return b, ok
}
