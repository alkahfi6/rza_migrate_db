package engine

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/output"
)

func TestVerifyThenGenFailsOnMalformedTableSource(t *testing.T) {
	t.Parallel()

	repo := setupMinimalGeneratorRepo(t)
	mustWriteFile(t, filepath.Join(repo, "db_object", "tables", "dbo.BadTable.sql"), `
CREATE TABLE dbo.BadTable {
  id INT NOT NULL
}
`)

	e := New(config.Default(repo))
	report, err := e.Run("verify_then_gen")
	if err != nil {
		t.Fatalf("verify_then_gen run failed: %v", err)
	}
	if !hasFailRuleWithMessage(report.Fail.Items, "TABLE_SOURCE_SYNTAX_INVALID", "invalid create table statement") {
		t.Fatalf("expected TABLE_SOURCE_SYNTAX_INVALID (malformed create table), got fail items: %#v", report.Fail.Items)
	}
}

func TestVerifyThenGenRepairsMissingCommaBetweenColumns(t *testing.T) {
	t.Parallel()

	repo := setupMinimalGeneratorRepo(t)
	mustWriteFile(t, filepath.Join(repo, "db_object", "tables", "dbo.TableA.sql"), `
CREATE TABLE dbo.TableA (
  [id] INT NOT NULL
  [name] VARCHAR(50) NULL
);
`)

	e := New(config.Default(repo))
	report, err := e.Run("verify_then_gen")
	if err != nil {
		t.Fatalf("verify_then_gen run failed: %v", err)
	}
	if hasFailRuleWithMessage(report.Fail.Items, "TABLE_SOURCE_SYNTAX_INVALID", "missing comma between table items") {
		t.Fatalf("missing comma should be repaired before validation, got fail items: %#v", report.Fail.Items)
	}
	repaired, err := os.ReadFile(filepath.Join(repo, "db_object", "tables", "dbo.TableA.sql"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(repaired), "[id] INT NOT NULL,\n") {
		t.Fatalf("expected source to contain repaired comma, got:\n%s", repaired)
	}
}

func TestVerifyThenGenStrictModeFailsWhenPrecheckMissing(t *testing.T) {
	t.Parallel()

	repo := setupMinimalGeneratorRepo(t)
	upName := "20260305000000001_tables_dbo.TableA_v1.up.sql"
	downName := "20260305000000001_tables_dbo.TableA_v1.down.sql"
	sql := `
/*
Severity mapping:
0=INFO
1=NOTICE
2=WARNING
3=ERROR
*/
SET NOCOUNT ON;
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = 20260305000000001;
DECLARE @category NVARCHAR(16) = N'tables';
DECLARE @file_name NVARCHAR(260) = N'20260305000000001_tables_dbo.TableA_v1.up.sql';
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
	mustWriteFile(t, filepath.Join(repo, "migrations", upName), sql)
	mustWriteFile(t, filepath.Join(repo, "migrations", downName), strings.ReplaceAll(sql, upName, downName))

	cfg := config.Default(repo)
	cfg.EnableStrictMode = true
	e := New(cfg)

	report, err := e.Run("verify_then_gen")
	if err != nil {
		t.Fatalf("verify_then_gen run failed: %v", err)
	}
	if !hasFailRuleWithMessage(report.Fail.Items, "MIGRATION_PRECHECK_MISSING", "table migration statement may run without explicit precheck guard") {
		t.Fatalf("expected strict precheck fail, got fail items: %#v", report.Fail.Items)
	}
}

func setupMinimalGeneratorRepo(t *testing.T) string {
	t.Helper()

	repo := t.TempDir()
	mustWriteFile(t, filepath.Join(repo, ".gitignore"), "out/**\n*.output.json\n")
	mustWriteFile(t, filepath.Join(repo, ".gitattributes"), "* text=auto eol=lf\n*.sql text eol=lf\n*.md text eol=lf\n*.go text eol=lf\n")
	mustWriteFile(t, filepath.Join(repo, "constraints_naming.md"), "# constraints naming\n")

	mustWriteFile(t, filepath.Join(repo, "migrations", "000000__core__mig__audit_init.up.sql"), `
IF OBJECT_ID(N'dbo.mig_event', N'U') IS NULL
BEGIN
  CREATE TABLE dbo.mig_event (
    event_id BIGINT IDENTITY(1,1) NOT NULL CONSTRAINT PK_mig_event PRIMARY KEY
  );
END
CREATE OR ALTER PROCEDURE dbo.usp_mig_log_event
AS
BEGIN
  SET NOCOUNT ON;
  SELECT 1;
END
CREATE OR ALTER PROCEDURE dbo.usp_mig_try_get_lock
AS
BEGIN
  SET NOCOUNT ON;
  SELECT 1;
END
`)
	mustWriteFile(t, filepath.Join(repo, "migrations", "000000__core__mig__audit_init.down.sql"), `
IF OBJECT_ID(N'dbo.usp_mig_try_get_lock', N'P') IS NOT NULL
  DROP PROCEDURE dbo.usp_mig_try_get_lock;
IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  DROP PROCEDURE dbo.usp_mig_log_event;
IF OBJECT_ID(N'dbo.mig_event', N'U') IS NOT NULL
  DROP TABLE dbo.mig_event;
`)

	return repo
}

func hasFailRuleWithMessage(items []output.Item, rule, message string) bool {
	message = strings.ToLower(strings.TrimSpace(message))
	for _, item := range items {
		if item.Rule != rule {
			continue
		}
		if message == "" || strings.Contains(strings.ToLower(item.Message), message) {
			return true
		}
	}
	return false
}
