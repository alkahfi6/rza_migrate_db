package engine

import (
	"path/filepath"
	"reflect"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/output"
)

func TestReconstructHistoryParallelMatchesSerial(t *testing.T) {
	t.Parallel()

	repo := t.TempDir()
	mustWriteFile(t, filepath.Join(repo, "migrations", "000000__core__mig__audit_init.up.sql"), "CREATE OR ALTER PROCEDURE dbo.usp_mig_log_event AS BEGIN SELECT 1; END;\n")
	mustWriteFile(t, filepath.Join(repo, "migrations", "000000__core__mig__audit_init.down.sql"), "CREATE OR ALTER PROCEDURE dbo.usp_mig_log_event AS BEGIN SELECT 1; END;\n")

	mustWriteFile(t, filepath.Join(repo, "migrations", "20260217073015000_tables_dbo.TableA_v1.up.sql"), "CREATE TABLE dbo.TableA (\n  id INT NOT NULL,\n  CONSTRAINT PK_TableA PRIMARY KEY (id)\n);\n")
	mustWriteFile(t, filepath.Join(repo, "migrations", "20260217073015000_tables_dbo.TableA_v1.down.sql"), "DROP TABLE dbo.TableA;\n")
	mustWriteFile(t, filepath.Join(repo, "migrations", "20260217073015001_indexes_dbo.TableA_v1.up.sql"), "CREATE INDEX IX_TableA__id ON dbo.TableA(id);\n")
	mustWriteFile(t, filepath.Join(repo, "migrations", "20260217073015001_indexes_dbo.TableA_v1.down.sql"), "DROP INDEX IX_TableA__id ON dbo.TableA;\n")
	mustWriteFile(t, filepath.Join(repo, "migrations", "20260217073015002_views_dbo.ViewA_v1.up.sql"), "CREATE OR ALTER VIEW dbo.ViewA AS SELECT id FROM dbo.TableA;\n")
	mustWriteFile(t, filepath.Join(repo, "migrations", "20260217073015002_views_dbo.ViewA_v1.down.sql"), "CREATE OR ALTER VIEW dbo.ViewA AS SELECT id FROM dbo.TableA;\n")

	// invalid filename to assert deterministic fail behavior in both modes
	mustWriteFile(t, filepath.Join(repo, "migrations", "bad_name.sql"), "SELECT 1;\n")

	serialCfg := config.Default(repo)
	serialCfg.EnableParallelRead = false
	serialCfg.ParallelWorkers = 4
	serialEngine := New(serialCfg)
	serialReport := output.NewReport(serialCfg.OutputNamingPattern, serialCfg.OutputNamingLineEndings)
	serialHistory, serialExistingFiles, serialMaxVersion, serialErr := serialEngine.reconstructHistory(&serialReport)
	if serialErr != nil {
		t.Fatalf("serial reconstructHistory error: %v", serialErr)
	}

	parallelCfg := config.Default(repo)
	parallelCfg.EnableParallelRead = true
	parallelCfg.ParallelWorkers = 4
	parallelEngine := New(parallelCfg)
	parallelReport := output.NewReport(parallelCfg.OutputNamingPattern, parallelCfg.OutputNamingLineEndings)
	parallelHistory, parallelExistingFiles, parallelMaxVersion, parallelErr := parallelEngine.reconstructHistory(&parallelReport)
	if parallelErr != nil {
		t.Fatalf("parallel reconstructHistory error: %v", parallelErr)
	}

	if !reflect.DeepEqual(serialHistory, parallelHistory) {
		t.Fatalf("history mismatch:\nserial=%#v\nparallel=%#v", serialHistory, parallelHistory)
	}
	if !reflect.DeepEqual(serialExistingFiles, parallelExistingFiles) {
		t.Fatalf("existing files mismatch:\nserial=%#v\nparallel=%#v", serialExistingFiles, parallelExistingFiles)
	}
	if serialMaxVersion != parallelMaxVersion {
		t.Fatalf("max version mismatch: serial=%s parallel=%s", serialMaxVersion, parallelMaxVersion)
	}
	if !reflect.DeepEqual(serialReport.Fail.Items, parallelReport.Fail.Items) {
		t.Fatalf("fail items mismatch:\nserial=%#v\nparallel=%#v", serialReport.Fail.Items, parallelReport.Fail.Items)
	}
	if !reflect.DeepEqual(serialReport.Notes["reconstruction"], parallelReport.Notes["reconstruction"]) {
		t.Fatalf("reconstruction notes mismatch")
	}
}
