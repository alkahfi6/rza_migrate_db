package engine

import (
	"path/filepath"
	"testing"

	"rza_migrate_db/internal/config"
)

func TestVerifyThenGenUsesPrecheckCaches(t *testing.T) {
	t.Parallel()

	repo := t.TempDir()
	mustWriteFile(t, filepath.Join(repo, ".gitignore"), "out/**\n*.output.json\n")
	mustWriteFile(t, filepath.Join(repo, ".gitattributes"), "* text=auto eol=lf\n*.sql text eol=lf\n*.md text eol=lf\n*.go text eol=lf\n")
	mustWriteFile(t, filepath.Join(repo, "constraints_naming.md"), "# constraints naming\n")

	mustWriteFile(t, filepath.Join(repo, "migrations", "000000__core__mig__audit_init.up.sql"), `
IF OBJECT_ID(N'dbo.mig_event', N'U') IS NULL
BEGIN
  CREATE TABLE dbo.mig_event (
    event_id BIGINT IDENTITY(1,1) NOT NULL CONSTRAINT PK_mig_event PRIMARY KEY,
    run_key UNIQUEIDENTIFIER NULL,
    version BIGINT NULL,
    created_at DATETIME NOT NULL CONSTRAINT DF_mig_event__created_at DEFAULT (GETDATE()),
    error_number INT NULL,
    error_line INT NULL,
    error_procedure NVARCHAR(128) NULL
  );
  CREATE INDEX IX_mig_event__run_key ON dbo.mig_event(run_key);
  CREATE INDEX IX_mig_event__version ON dbo.mig_event(version);
END
CREATE OR ALTER PROCEDURE dbo.usp_mig_log_event
  @error_number INT = NULL,
  @error_line INT = NULL,
  @error_procedure NVARCHAR(128) = NULL
AS
BEGIN
  SET NOCOUNT ON;
  INSERT INTO dbo.mig_event(run_key, version, created_at, error_number, error_line, error_procedure)
  VALUES (NEWID(), 1, GETDATE(), @error_number, @error_line, @error_procedure);
END
CREATE OR ALTER PROCEDURE dbo.usp_mig_try_get_lock
AS
BEGIN
  DECLARE @result INT;
  EXEC @result = sp_getapplock @Resource = N'rza_migrate_db', @LockMode = N'Exclusive', @LockOwner = N'Session', @LockTimeout = 1000;
  IF (@result < 0)
  BEGIN
    THROW 50001, 'lock failed', 1;
  END
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

	mustWriteFile(t, filepath.Join(repo, "migrations", "20260303010101000_procedures_dbo.ProcA_v1.up.sql"), `
CREATE OR ALTER PROCEDURE dbo.ProcA
AS
BEGIN
  SET NOCOUNT ON;
  SELECT 1 AS id;
END
`)
	mustWriteFile(t, filepath.Join(repo, "migrations", "20260303010101000_procedures_dbo.ProcA_v1.down.sql"), `
CREATE OR ALTER PROCEDURE dbo.ProcA
AS
BEGIN
  SET NOCOUNT ON;
  SELECT 0 AS id;
END
`)
	mustWriteFile(t, filepath.Join(repo, "db_object", "procedures", "dbo.ProcA.sql"), `
CREATE OR ALTER PROCEDURE dbo.ProcA
AS
BEGIN
  SET NOCOUNT ON;
  SELECT 1 AS id;
END
`)

	cfg := config.Default(repo)
	e := New(cfg)

	report1, err := e.Run("verify_then_gen")
	if err != nil {
		t.Fatalf("first verify_then_gen run failed: %v", err)
	}
	if report1.Fail.Count != 0 {
		t.Fatalf("first verify_then_gen expected zero fail, got: %#v", report1.Fail.Items)
	}
	if hit, ok := noteBool(report1, "gen_precheck_history_cache_hit"); !ok || !hit {
		t.Fatalf("expected gen_precheck_history_cache_hit=true on first run, got hit=%v ok=%v", hit, ok)
	}
	if hit, ok := noteBool(report1, "gen_precheck_desired_cache_hit"); !ok || !hit {
		t.Fatalf("expected gen_precheck_desired_cache_hit=true on first run, got hit=%v ok=%v", hit, ok)
	}

	report2, err := e.Run("verify_then_gen")
	if err != nil {
		t.Fatalf("second verify_then_gen run failed: %v", err)
	}
	if report2.Fail.Count != 0 {
		t.Fatalf("second verify_then_gen expected zero fail, got: %#v", report2.Fail.Items)
	}
	if hit, ok := noteBool(report2, "gen_precheck_history_cache_hit"); !ok || !hit {
		t.Fatalf("expected gen_precheck_history_cache_hit=true on second run, got hit=%v ok=%v", hit, ok)
	}
	if hit, ok := noteBool(report2, "gen_precheck_desired_cache_hit"); !ok || !hit {
		t.Fatalf("expected gen_precheck_desired_cache_hit=true on second run, got hit=%v ok=%v", hit, ok)
	}
}
