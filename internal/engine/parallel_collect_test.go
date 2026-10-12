package engine

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/output"
)

func TestCollectDesiredObjectsParallelMatchesSerial(t *testing.T) {
	t.Parallel()

	repo := t.TempDir()
	mustWriteFile(t, filepath.Join(repo, "db_object", "schema", "tables", "dbo.TableA.sql"), "CREATE TABLE dbo.TableA (\n  id INT NOT NULL\n);\n")
	mustWriteFile(t, filepath.Join(repo, "db_object", "schema", "tables", "bad.TableB.sql"), "CREATE TABLE dbo.TableB (\n  id INT NOT NULL\n);\n")
	mustWriteFile(t, filepath.Join(repo, "db_object", "schema", "indexes", "dbo.TableA.sql"), "CREATE INDEX IX_TableA__id ON dbo.TableA(id);\n")
	mustWriteFile(t, filepath.Join(repo, "db_object", "schema", "views", "dbo.ViewA.sql"), "CREATE OR ALTER VIEW dbo.ViewA AS SELECT 1 AS id;\n")
	mustWriteFile(t, filepath.Join(repo, "db_object", "schema", "functions", "dbo.FuncA.sql"), "CREATE OR ALTER FUNCTION dbo.FuncA() RETURNS INT AS BEGIN RETURN 1; END;\n")
	mustWriteFile(t, filepath.Join(repo, "db_object", "schema", "procedures", "dbo.ProcA.sql"), "CREATE OR ALTER PROCEDURE dbo.ProcA AS BEGIN SELECT 1; END;\n")

	serialCfg := config.Default(repo)
	serialCfg.EnableParallelRead = false
	serialCfg.ParallelWorkers = 4
	serialEngine := New(serialCfg)
	serialReport := output.NewReport(serialCfg.OutputNamingPattern, serialCfg.OutputNamingLineEndings)
	serialObjects, serialErr := serialEngine.collectDesiredObjects(&serialReport)
	if serialErr != nil {
		t.Fatalf("serial collectDesiredObjects error: %v", serialErr)
	}

	parallelCfg := config.Default(repo)
	parallelCfg.EnableParallelRead = true
	parallelCfg.ParallelWorkers = 4
	parallelEngine := New(parallelCfg)
	parallelReport := output.NewReport(parallelCfg.OutputNamingPattern, parallelCfg.OutputNamingLineEndings)
	parallelObjects, parallelErr := parallelEngine.collectDesiredObjects(&parallelReport)
	if parallelErr != nil {
		t.Fatalf("parallel collectDesiredObjects error: %v", parallelErr)
	}

	if !reflect.DeepEqual(serialObjects, parallelObjects) {
		t.Fatalf("desired objects mismatch:\nserial=%#v\nparallel=%#v", serialObjects, parallelObjects)
	}
	if !reflect.DeepEqual(serialReport.Fail.Items, parallelReport.Fail.Items) {
		t.Fatalf("fail items mismatch:\nserial=%#v\nparallel=%#v", serialReport.Fail.Items, parallelReport.Fail.Items)
	}
}

func TestCollectPatchSourcesParallelMatchesSerial(t *testing.T) {
	t.Parallel()

	repo := t.TempDir()
	mustWriteFile(t, filepath.Join(repo, "patch", "Fix Status", "README.md"), "# patch\n")
	mustWriteFile(t, filepath.Join(repo, "patch", "Fix Status", "up.sql"), "UPDATE dbo.TableA SET id = id;\n")
	mustWriteFile(t, filepath.Join(repo, "patch", "Fix Status", "down.sql"), "UPDATE dbo.TableA SET id = id;\n")

	// duplicate canonical name (fix_status)
	mustWriteFile(t, filepath.Join(repo, "patch", "fix_status", "README.md"), "# duplicate\n")
	mustWriteFile(t, filepath.Join(repo, "patch", "fix_status", "up.sql"), "SELECT 1;\n")
	mustWriteFile(t, filepath.Join(repo, "patch", "fix_status", "down.sql"), "SELECT 1;\n")

	// missing down.sql
	mustWriteFile(t, filepath.Join(repo, "patch", "Backfill Missing Down", "README.md"), "# missing down\n")
	mustWriteFile(t, filepath.Join(repo, "patch", "Backfill Missing Down", "up.sql"), "SELECT 1;\n")

	serialCfg := config.Default(repo)
	serialCfg.EnableParallelRead = false
	serialCfg.ParallelWorkers = 4
	serialEngine := New(serialCfg)
	serialReport := output.NewReport(serialCfg.OutputNamingPattern, serialCfg.OutputNamingLineEndings)
	serialPatches, serialErr := serialEngine.collectPatchSources(&serialReport)
	if serialErr != nil {
		t.Fatalf("serial collectPatchSources error: %v", serialErr)
	}

	parallelCfg := config.Default(repo)
	parallelCfg.EnableParallelRead = true
	parallelCfg.ParallelWorkers = 4
	parallelEngine := New(parallelCfg)
	parallelReport := output.NewReport(parallelCfg.OutputNamingPattern, parallelCfg.OutputNamingLineEndings)
	parallelPatches, parallelErr := parallelEngine.collectPatchSources(&parallelReport)
	if parallelErr != nil {
		t.Fatalf("parallel collectPatchSources error: %v", parallelErr)
	}

	if !reflect.DeepEqual(serialPatches, parallelPatches) {
		t.Fatalf("patch sources mismatch:\nserial=%#v\nparallel=%#v", serialPatches, parallelPatches)
	}
	if !reflect.DeepEqual(serialReport.Fail.Items, parallelReport.Fail.Items) {
		t.Fatalf("fail items mismatch:\nserial=%#v\nparallel=%#v", serialReport.Fail.Items, parallelReport.Fail.Items)
	}
}

func TestReadParallelWorkersCapsByCPUAndTaskCount(t *testing.T) {
	t.Parallel()

	cfg := config.Default(".")
	cfg.EnableParallelRead = true
	cfg.ParallelWorkers = 999
	e := New(cfg)

	taskCount := 5
	want := taskCount
	if cpu := runtime.GOMAXPROCS(0); cpu < want {
		want = cpu
	}
	if got := e.readParallelWorkers(taskCount); got != want {
		t.Fatalf("unexpected workers cap: got=%d want=%d", got, want)
	}
}

func mustWriteFile(t *testing.T, path string, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir failed for %s: %v", path, err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write failed for %s: %v", path, err)
	}
}
