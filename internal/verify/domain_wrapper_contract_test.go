package verify

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rza_migrate_db/internal/output"
	"rza_migrate_db/internal/sqlparse"
)

func TestBootstrapCompatibleCreateTableGuardContract(t *testing.T) {
	t.Parallel()

	good := `
-- existing table match profile: bootstrap-compatible
IF EXISTS (
  SELECT 1
  FROM sys.columns AS c
  WHERE UPPER(TYPE_NAME(c.user_type_id)) = N'INT'
    AND c.is_identity = 0
    AND c.is_computed = 0
)
BEGIN
  CREATE TABLE dbo.T ([Id] INT NOT NULL);
END`
	report := output.NewReport("", "")
	validateBootstrapCompatibleCreateTableGuard(
		&report,
		"migrations/20260728000000000_tables_dbo.T_v1.up.sql",
		strings.ToUpper(good),
		strings.ToUpper(good),
		true,
	)
	if report.Fail.Count != 0 {
		t.Fatalf("valid bootstrap-compatible guard failed contract: %#v", report.Fail.Items)
	}

	bad := `
IF EXISTS (
  SELECT COUNT_BIG(*) FROM sys.columns AS c
  WHERE c.max_length = 20
)
BEGIN
  CREATE TABLE dbo.T ([Id] INT NOT NULL);
END`
	report = output.NewReport("", "")
	validateBootstrapCompatibleCreateTableGuard(
		&report,
		"migrations/20260728000000000_tables_dbo.T_v1.up.sql",
		strings.ToUpper(bad),
		strings.ToUpper(bad),
		true,
	)
	if report.Fail.Count == 0 {
		t.Fatal("strict checker accepted an exact/length-based CREATE TABLE guard")
	}

	withNoise := good + `
-- COUNT_BIG(*) FROM SYS.COLUMNS and C.MAX_LENGTH = are documentation only.
DECLARE @note NVARCHAR(200) = N'FROM SYS.FOREIGN_KEYS AS FK';`
	report = output.NewReport("", "")
	validateBootstrapCompatibleCreateTableGuard(
		&report,
		"migrations/20260728000000000_tables_dbo.T_v1.up.sql",
		strings.ToUpper(withNoise),
		strings.ToUpper(sqlparse.StripSQLStringsAndComments(withNoise)),
		true,
	)
	if report.Fail.Count != 0 {
		t.Fatalf("comments/literals caused false bootstrap guard findings: %#v", report.Fail.Items)
	}
}

func TestValidateWrapperSubsetWarnsWhenTablePrecheckMissing(t *testing.T) {
	t.Parallel()

	repo := t.TempDir()
	rel := "migrations/20260304000000000_tables_dbo.TableA_v2.up.sql"
	sql := `/*
Severity mapping:
0=INFO
1=NOTICE
2=WARNING
3=ERROR
*/
SET NOCOUNT ON;
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = 20260304000000000;
DECLARE @category NVARCHAR(16) = N'tables';
DECLARE @file_name NVARCHAR(260) = N'20260304000000000_tables_dbo.TableA_v2.up.sql';
DECLARE @target NVARCHAR(256) = N'dbo.TableA';
DECLARE @action NVARCHAR(64) = N'alter_column';
DECLARE @direction NVARCHAR(8) = N'up';
BEGIN TRY
  BEGIN TRAN;
  EXEC dbo.usp_mig_log_event
    @run_key=@run_key,@version=@version,@category=@category,@file_name=@file_name,
    @decision=N'run',@severity=1,@action=@action,@target=@target,@reason=N'OBJECT_EXISTS',@message=N'run',@detail=@direction,
    @error_number=NULL,@error_line=NULL,@error_procedure=NULL;
  ALTER TABLE dbo.TableA ADD col1 INT NULL;
  EXEC dbo.usp_mig_log_event
    @run_key=@run_key,@version=@version,@category=@category,@file_name=@file_name,
    @decision=N'ok',@severity=1,@action=@action,@target=@target,@reason=N'CREATED',@message=N'ok',@detail=@direction,
    @error_number=NULL,@error_line=NULL,@error_procedure=NULL;
  COMMIT TRAN;
END TRY
BEGIN CATCH
  IF XACT_STATE() <> 0 ROLLBACK;
  EXEC dbo.usp_mig_log_event
    @run_key=@run_key,@version=@version,@category=@category,@file_name=@file_name,
    @decision=N'fail',@severity=3,@action=N'migration_failed',@target=@target,@reason=N'SQL_ERROR',@message=N'fail',@detail=ERROR_MESSAGE(),
    @error_number=ERROR_NUMBER(),@error_line=ERROR_LINE(),@error_procedure=ERROR_PROCEDURE();
  THROW;
END CATCH;`

	abs := filepath.Join(repo, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(abs), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	if err := os.WriteFile(abs, []byte(sql), 0o644); err != nil {
		t.Fatalf("write failed: %v", err)
	}

	report := output.NewReport("pattern", "LF")
	err := ValidateWrapperSubset(&report, WrapperSubsetOptions{
		RepoPath: func(p string) string { return filepath.Join(repo, filepath.FromSlash(p)) },
		Metas: []MigrationMeta{{
			FileName:  filepath.Base(rel),
			Version:   "20260304000000000",
			Category:  "tables",
			Object:    "dbo.TableA",
			Direction: "up",
			RelPath:   rel,
			VK:        2,
		}},
		AllowedDecisions:  map[string]struct{}{"run": {}, "ok": {}, "fail": {}},
		AllowedActions:    map[string]struct{}{"alter_column": {}, "migration_failed": {}},
		AllowedReasons:    map[string]struct{}{"OBJECT_EXISTS": {}, "CREATED": {}, "SQL_ERROR": {}},
		AllowedSeverities: map[int]struct{}{0: {}, 1: {}, 2: {}, 3: {}},
	})
	if err != nil {
		t.Fatalf("validate wrapper failed: %v", err)
	}

	found := false
	for _, item := range report.Pass.Items {
		if item.Rule == "MIGRATION_PRECHECK_MISSING" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected MIGRATION_PRECHECK_MISSING warning, pass=%#v fail=%#v", report.Pass.Items, report.Fail.Items)
	}
}

func TestValidateWrapperSubsetFailsWhenTablePrecheckMissingInStrictMode(t *testing.T) {
	t.Parallel()

	repo := t.TempDir()
	rel := "migrations/20260304000000001_tables_dbo.TableA_v3.up.sql"
	sql := `/*
Severity mapping:
0=INFO
1=NOTICE
2=WARNING
3=ERROR
*/
SET NOCOUNT ON;
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = 20260304000000001;
DECLARE @category NVARCHAR(16) = N'tables';
DECLARE @file_name NVARCHAR(260) = N'20260304000000001_tables_dbo.TableA_v3.up.sql';
DECLARE @target NVARCHAR(256) = N'dbo.TableA';
DECLARE @action NVARCHAR(64) = N'alter_column';
DECLARE @direction NVARCHAR(8) = N'up';
BEGIN TRY
  BEGIN TRAN;
  EXEC dbo.usp_mig_log_event
    @run_key=@run_key,@version=@version,@category=@category,@file_name=@file_name,
    @decision=N'run',@severity=1,@action=@action,@target=@target,@reason=N'OBJECT_EXISTS',@message=N'run',@detail=@direction,
    @error_number=NULL,@error_line=NULL,@error_procedure=NULL;
  ALTER TABLE dbo.TableA ADD col2 INT NULL;
  EXEC dbo.usp_mig_log_event
    @run_key=@run_key,@version=@version,@category=@category,@file_name=@file_name,
    @decision=N'ok',@severity=1,@action=@action,@target=@target,@reason=N'CREATED',@message=N'ok',@detail=@direction,
    @error_number=NULL,@error_line=NULL,@error_procedure=NULL;
  COMMIT TRAN;
END TRY
BEGIN CATCH
  IF XACT_STATE() <> 0 ROLLBACK;
  EXEC dbo.usp_mig_log_event
    @run_key=@run_key,@version=@version,@category=@category,@file_name=@file_name,
    @decision=N'fail',@severity=3,@action=N'migration_failed',@target=@target,@reason=N'SQL_ERROR',@message=N'fail',@detail=ERROR_MESSAGE(),
    @error_number=ERROR_NUMBER(),@error_line=ERROR_LINE(),@error_procedure=ERROR_PROCEDURE();
  THROW;
END CATCH;`

	abs := filepath.Join(repo, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(abs), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	if err := os.WriteFile(abs, []byte(sql), 0o644); err != nil {
		t.Fatalf("write failed: %v", err)
	}

	report := output.NewReport("pattern", "LF")
	err := ValidateWrapperSubset(&report, WrapperSubsetOptions{
		RepoPath: func(p string) string { return filepath.Join(repo, filepath.FromSlash(p)) },
		Metas: []MigrationMeta{{
			FileName:  filepath.Base(rel),
			Version:   "20260304000000001",
			Category:  "tables",
			Object:    "dbo.TableA",
			Direction: "up",
			RelPath:   rel,
			VK:        3,
		}},
		AllowedDecisions:  map[string]struct{}{"run": {}, "ok": {}, "fail": {}},
		AllowedActions:    map[string]struct{}{"alter_column": {}, "migration_failed": {}},
		AllowedReasons:    map[string]struct{}{"OBJECT_EXISTS": {}, "CREATED": {}, "SQL_ERROR": {}},
		AllowedSeverities: map[int]struct{}{0: {}, 1: {}, 2: {}, 3: {}},
		StrictMode:        true,
	})
	if err != nil {
		t.Fatalf("validate wrapper failed: %v", err)
	}

	found := false
	for _, item := range report.Fail.Items {
		if item.Rule == "MIGRATION_PRECHECK_MISSING" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected MIGRATION_PRECHECK_MISSING fail in strict mode, pass=%#v fail=%#v", report.Pass.Items, report.Fail.Items)
	}
}

func TestValidateWrapperSubsetAcceptsIndexGuardWithReturnBeforeTry(t *testing.T) {
	t.Parallel()

	repo := t.TempDir()
	rel := "migrations/20260304000000002_indexes_dbo.TableA_v1.up.sql"
	sql := `/*
Severity mapping:
0=INFO
1=NOTICE
2=WARNING
3=ERROR
*/
SET NOCOUNT ON;
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = 20260304000000002;
DECLARE @category NVARCHAR(16) = N'indexes';
DECLARE @file_name NVARCHAR(260) = N'20260304000000002_indexes_dbo.TableA_v1.up.sql';
DECLARE @target NVARCHAR(256) = N'dbo.TableA';
DECLARE @action NVARCHAR(64) = N'add_index';
DECLARE @direction NVARCHAR(8) = N'up';

IF EXISTS (
  SELECT 1
  FROM sys.indexes
  WHERE object_id = OBJECT_ID(N'dbo.TableA')
    AND name = N'ix_TableA_A'
)
BEGIN
  RETURN;
END

BEGIN TRY
  BEGIN TRAN;
  CREATE NONCLUSTERED INDEX [ix_TableA_A] ON dbo.TableA ([A] ASC);
  COMMIT TRAN;
END TRY
BEGIN CATCH
  IF XACT_STATE() <> 0 ROLLBACK;
  THROW;
END CATCH;`

	abs := filepath.Join(repo, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(abs), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	if err := os.WriteFile(abs, []byte(sql), 0o644); err != nil {
		t.Fatalf("write failed: %v", err)
	}

	report := output.NewReport("pattern", "LF")
	err := ValidateWrapperSubset(&report, WrapperSubsetOptions{
		RepoPath: func(p string) string { return filepath.Join(repo, filepath.FromSlash(p)) },
		Metas: []MigrationMeta{{
			FileName:  filepath.Base(rel),
			Version:   "20260304000000002",
			Category:  "indexes",
			Object:    "dbo.TableA",
			Direction: "up",
			RelPath:   rel,
			VK:        1,
		}},
		AllowedDecisions:  map[string]struct{}{"run": {}, "ok": {}, "fail": {}, "skip": {}},
		AllowedActions:    map[string]struct{}{"add_index": {}, "migration_failed": {}},
		AllowedReasons:    map[string]struct{}{"OBJECT_EXISTS": {}, "CREATED": {}, "SQL_ERROR": {}, "INDEX_EXISTS": {}},
		AllowedSeverities: map[int]struct{}{0: {}, 1: {}, 2: {}, 3: {}},
		StrictMode:        true,
	})
	if err != nil {
		t.Fatalf("validate wrapper failed: %v", err)
	}

	for _, item := range report.Fail.Items {
		if item.Rule == "MIGRATION_PRECHECK_MISSING" {
			t.Fatalf("unexpected MIGRATION_PRECHECK_MISSING for valid index guard-return pattern: %#v", item)
		}
	}
}

func TestValidateWrapperSubsetLegacyNoPrecheckIsWarningInStrictMode(t *testing.T) {
	t.Parallel()

	repo := t.TempDir()
	rel := "migrations/20260304000000003_indexes_dbo.TableA_v1.up.sql"
	sql := `/*
Severity mapping:
0=INFO
1=NOTICE
2=WARNING
3=ERROR
*/
SET NOCOUNT ON;
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = 20260304000000003;
DECLARE @category NVARCHAR(16) = N'indexes';
DECLARE @file_name NVARCHAR(260) = N'20260304000000003_indexes_dbo.TableA_v1.up.sql';
DECLARE @target NVARCHAR(256) = N'dbo.TableA';
DECLARE @action NVARCHAR(64) = N'add_index';
DECLARE @direction NVARCHAR(8) = N'up';

-- no precheck

BEGIN TRY
  BEGIN TRAN;
  CREATE NONCLUSTERED INDEX [ix_TableA_A] ON dbo.TableA ([A] ASC);
  COMMIT TRAN;
END TRY
BEGIN CATCH
  IF XACT_STATE() <> 0 ROLLBACK;
  THROW;
END CATCH;`

	abs := filepath.Join(repo, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(abs), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	if err := os.WriteFile(abs, []byte(sql), 0o644); err != nil {
		t.Fatalf("write failed: %v", err)
	}

	report := output.NewReport("pattern", "LF")
	err := ValidateWrapperSubset(&report, WrapperSubsetOptions{
		RepoPath: func(p string) string { return filepath.Join(repo, filepath.FromSlash(p)) },
		Metas: []MigrationMeta{{
			FileName:  filepath.Base(rel),
			Version:   "20260304000000003",
			Category:  "indexes",
			Object:    "dbo.TableA",
			Direction: "up",
			RelPath:   rel,
			VK:        1,
		}},
		AllowedDecisions:  map[string]struct{}{"run": {}, "ok": {}, "fail": {}, "skip": {}},
		AllowedActions:    map[string]struct{}{"add_index": {}, "migration_failed": {}},
		AllowedReasons:    map[string]struct{}{"OBJECT_EXISTS": {}, "CREATED": {}, "SQL_ERROR": {}, "INDEX_EXISTS": {}},
		AllowedSeverities: map[int]struct{}{0: {}, 1: {}, 2: {}, 3: {}},
		StrictMode:        true,
	})
	if err != nil {
		t.Fatalf("validate wrapper failed: %v", err)
	}

	for _, item := range report.Fail.Items {
		if item.Rule == "MIGRATION_PRECHECK_MISSING" {
			t.Fatalf("legacy no-precheck must not fail in strict mode: %#v", item)
		}
	}
	for _, item := range report.Pass.Items {
		if item.Rule == "MIGRATION_PRECHECK_MISSING" {
			t.Fatalf("legacy no-precheck must be ignored (not warned): %#v", item)
		}
	}
}

func TestValidateWrapperSubsetAcceptsSkipVariablePrecheckPattern(t *testing.T) {
	t.Parallel()

	repo := t.TempDir()
	rel := "migrations/20260304000000004_tables_dbo.TableA_v2.up.sql"
	sql := `/*
Severity mapping:
0=INFO
1=NOTICE
2=WARNING
3=ERROR
*/
SET NOCOUNT ON;
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = 20260304000000004;
DECLARE @category NVARCHAR(16) = N'tables';
DECLARE @file_name NVARCHAR(260) = N'20260304000000004_tables_dbo.TableA_v2.up.sql';
DECLARE @target NVARCHAR(256) = N'dbo.TableA';
DECLARE @action NVARCHAR(64) = N'alter_column';
DECLARE @direction NVARCHAR(8) = N'up';
BEGIN TRY
  BEGIN TRAN;
  DECLARE @skip_alter_colA BIT = 0;
  IF COL_LENGTH(N'dbo.TableA', N'colA') IS NULL
  BEGIN
    SET @skip_alter_colA = 1;
  END
  IF (@skip_alter_colA = 0)
  BEGIN
    ALTER TABLE dbo.TableA ALTER COLUMN [colA] VARCHAR(20) NULL;
  END
  COMMIT TRAN;
END TRY
BEGIN CATCH
  IF XACT_STATE() <> 0 ROLLBACK;
  THROW;
END CATCH;`

	abs := filepath.Join(repo, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(abs), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	if err := os.WriteFile(abs, []byte(sql), 0o644); err != nil {
		t.Fatalf("write failed: %v", err)
	}

	report := output.NewReport("pattern", "LF")
	err := ValidateWrapperSubset(&report, WrapperSubsetOptions{
		RepoPath: func(p string) string { return filepath.Join(repo, filepath.FromSlash(p)) },
		Metas: []MigrationMeta{{
			FileName:  filepath.Base(rel),
			Version:   "20260304000000004",
			Category:  "tables",
			Object:    "dbo.TableA",
			Direction: "up",
			RelPath:   rel,
			VK:        2,
		}},
		AllowedDecisions:  map[string]struct{}{"run": {}, "ok": {}, "fail": {}, "skip": {}},
		AllowedActions:    map[string]struct{}{"alter_column": {}, "migration_failed": {}},
		AllowedReasons:    map[string]struct{}{"OBJECT_EXISTS": {}, "ALTERED": {}, "SQL_ERROR": {}, "COLUMN_NOT_EXISTS": {}},
		AllowedSeverities: map[int]struct{}{0: {}, 1: {}, 2: {}, 3: {}},
		StrictMode:        true,
	})
	if err != nil {
		t.Fatalf("validate wrapper failed: %v", err)
	}

	for _, item := range report.Fail.Items {
		if item.Rule == "MIGRATION_PRECHECK_MISSING" {
			t.Fatalf("skip-variable precheck pattern should be accepted: %#v", item)
		}
	}
}
