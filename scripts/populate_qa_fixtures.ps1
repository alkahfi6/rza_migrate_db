$ErrorActionPreference = "Stop"

$root = Join-Path (Get-Location) "fixtures/qa_rules"
$enc = New-Object System.Text.UTF8Encoding($false)

function Write-LFFile {
  param(
    [string]$Path,
    [string]$Content
  )
  $dir = Split-Path -Parent $Path
  if ($dir -and -not (Test-Path $dir)) {
    New-Item -ItemType Directory -Force -Path $dir | Out-Null
  }
  $normalized = $Content -replace "`r`n", "`n"
  [System.IO.File]::WriteAllText($Path, $normalized, $enc)
}

function Baseline-Sql {
  return @{
    table_source = @'
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL,
  amount DECIMAL(10,2) NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id),
  CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status
);
'@
    index_source = @'
CREATE INDEX IX_TableA__status ON dbo.TableA(status);
'@
    view_source = @'
CREATE OR ALTER VIEW dbo.ViewA
AS
SELECT id, status, amount
FROM dbo.TableA;
'@
    function_source = @'
CREATE OR ALTER FUNCTION dbo.FuncA (@id INT)
RETURNS INT
AS
BEGIN
  RETURN @id;
END
'@
    procedure_source = @'
CREATE OR ALTER PROCEDURE dbo.ProcA
AS
BEGIN
  SET NOCOUNT ON;
  SELECT COUNT_BIG(1) AS total_rows FROM dbo.TableA;
END
'@
    core_up = @'
IF OBJECT_ID(N'dbo.mig_event', N'U') IS NULL
BEGIN
  CREATE TABLE dbo.mig_event (
    event_id BIGINT IDENTITY(1,1) NOT NULL CONSTRAINT PK_mig_event PRIMARY KEY,
    run_key UNIQUEIDENTIFIER NOT NULL,
    version BIGINT NOT NULL,
    category NVARCHAR(16) NOT NULL,
    file_name NVARCHAR(260) NOT NULL,
    decision NVARCHAR(8) NOT NULL,
    severity TINYINT NOT NULL,
    action NVARCHAR(64) NOT NULL,
    target NVARCHAR(256) NOT NULL,
    reason NVARCHAR(64) NOT NULL,
    message NVARCHAR(4000) NOT NULL,
    detail NVARCHAR(MAX) NULL,
    error_number INT NULL,
    error_line INT NULL,
    error_procedure NVARCHAR(128) NULL,
    created_at DATETIME2(7) NOT NULL CONSTRAINT DF_mig_event__created_at DEFAULT (GETDATE())
  );

  CREATE INDEX IX_mig_event__run_key ON dbo.mig_event (run_key, created_at);
  CREATE INDEX IX_mig_event__version ON dbo.mig_event (version, created_at);
END

IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NULL
BEGIN
  EXEC (N'CREATE PROCEDURE dbo.usp_mig_log_event AS BEGIN SET NOCOUNT ON; END');
END

ALTER PROCEDURE dbo.usp_mig_log_event
  @run_key UNIQUEIDENTIFIER,
  @version BIGINT,
  @category NVARCHAR(16),
  @file_name NVARCHAR(260),
  @decision NVARCHAR(8),
  @severity TINYINT,
  @action NVARCHAR(64),
  @target NVARCHAR(256),
  @reason NVARCHAR(64),
  @message NVARCHAR(4000),
  @detail NVARCHAR(MAX) = NULL,
  @error_number INT = NULL,
  @error_line INT = NULL,
  @error_procedure NVARCHAR(128) = NULL
AS
BEGIN
  SET NOCOUNT ON;
  INSERT INTO dbo.mig_event (
    run_key, version, category, file_name,
    decision, severity, action, target, reason, message, detail,
    error_number, error_line, error_procedure
  )
  VALUES (
    @run_key, @version, @category, @file_name,
    @decision, @severity, @action, @target, @reason, @message, @detail,
    @error_number, @error_line, @error_procedure
  );
END

IF OBJECT_ID(N'dbo.usp_mig_try_get_lock', N'P') IS NULL
BEGIN
  EXEC (N'CREATE PROCEDURE dbo.usp_mig_try_get_lock AS BEGIN SET NOCOUNT ON; END');
END

ALTER PROCEDURE dbo.usp_mig_try_get_lock
  @resource NVARCHAR(255),
  @timeout_ms INT = 60000
AS
BEGIN
  SET NOCOUNT ON;

  DECLARE @res INT;
  EXEC @res = sp_getapplock
    @Resource = @resource,
    @LockMode = 'Exclusive',
    @LockOwner = 'Session',
    @LockTimeout = @timeout_ms;

  IF (@res < 0)
  BEGIN
    THROW 51000, 'Failed to acquire migration lock (sp_getapplock).', 1;
  END
END
'@
    core_down = @'
IF OBJECT_ID(N'dbo.usp_mig_try_get_lock', N'P') IS NOT NULL
  DROP PROCEDURE dbo.usp_mig_try_get_lock;

IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  DROP PROCEDURE dbo.usp_mig_log_event;

IF OBJECT_ID(N'dbo.mig_event', N'U') IS NOT NULL
  DROP TABLE dbo.mig_event;
'@
    table_up = @'
/*
Severity mapping:
0=INFO (ok/skip normal)
1=NOTICE (schema change)
2=WARNING (follow-up/manual risk)
3=ERROR (stop)
*/
SET NOCOUNT ON;
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = 20260217073015000;
DECLARE @category NVARCHAR(16) = N'tables';
DECLARE @file NVARCHAR(260) = N'20260217073015000_tables_dbo.TableA_v1.up.sql';
DECLARE @target NVARCHAR(256) = N'dbo.TableA';
DECLARE @action NVARCHAR(64) = N'create_table';
DECLARE @direction NVARCHAR(8) = N'up';

IF OBJECT_ID(N'dbo.TableA', N'U') IS NOT NULL
BEGIN
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'skip',
      @severity = 0,
      @action = @action,
      @target = @target,
      @reason = N'TABLE_EXISTS',
      @message = N'[PRECHECK] table already exists',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END
  RETURN;
END

BEGIN TRY
  BEGIN TRAN;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'run',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'OBJECT_EXISTS',
      @message = N'migration run start',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END

  CREATE TABLE dbo.TableA (
    id INT NOT NULL,
    status VARCHAR(10) NULL,
    amount DECIMAL(10,2) NULL,
    CONSTRAINT PK_TableA PRIMARY KEY (id),
    CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status
  );

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'ok',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'CREATED',
      @message = N'migration step ok',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END

  COMMIT;
END TRY
BEGIN CATCH
  IF XACT_STATE() <> 0 ROLLBACK;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
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
'@
table_down = @'
/*
Severity mapping:
0=INFO (ok/skip normal)
1=NOTICE (schema change)
2=WARNING (follow-up/manual risk)
3=ERROR (stop)
*/
SET NOCOUNT ON;
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = 20260217073015000;
DECLARE @category NVARCHAR(16) = N'tables';
DECLARE @file NVARCHAR(260) = N'20260217073015000_tables_dbo.TableA_v1.down.sql';
DECLARE @target NVARCHAR(256) = N'dbo.TableA';
DECLARE @action NVARCHAR(64) = N'drop_table';
DECLARE @direction NVARCHAR(8) = N'down';

IF OBJECT_ID(N'dbo.TableA', N'U') IS NULL
BEGIN
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'skip',
      @severity = 0,
      @action = @action,
      @target = @target,
      @reason = N'OBJECT_NOT_EXISTS',
      @message = N'[PRECHECK] table already absent',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END
  RETURN;
END

BEGIN TRY
  BEGIN TRAN;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'run',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'OBJECT_EXISTS',
      @message = N'migration run start',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END

  DROP TABLE dbo.TableA;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'ok',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'DROPPED',
      @message = N'migration step ok',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END

  COMMIT;
END TRY
BEGIN CATCH
  IF XACT_STATE() <> 0 ROLLBACK;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
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
'@
index_up = @'
/*
Severity mapping:
0=INFO (ok/skip normal)
1=NOTICE (schema change)
2=WARNING (follow-up/manual risk)
3=ERROR (stop)
*/
SET NOCOUNT ON;
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = 20260217073015001;
DECLARE @category NVARCHAR(16) = N'indexes';
DECLARE @file NVARCHAR(260) = N'20260217073015001_indexes_dbo.TableA_v1.up.sql';
DECLARE @target NVARCHAR(256) = N'dbo.TableA.IX_TableA__status';
DECLARE @action NVARCHAR(64) = N'add_index';
DECLARE @direction NVARCHAR(8) = N'up';

IF EXISTS (
  SELECT 1
  FROM sys.indexes
  WHERE object_id = OBJECT_ID(N'dbo.TableA')
    AND name = N'IX_TableA__status'
)
BEGIN
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'skip',
      @severity = 0,
      @action = @action,
      @target = @target,
      @reason = N'INDEX_EXISTS',
      @message = N'[PRECHECK] index already exists',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END
  RETURN;
END

BEGIN TRY
  BEGIN TRAN;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'run',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'OBJECT_EXISTS',
      @message = N'migration run start',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END

  CREATE INDEX IX_TableA__status ON dbo.TableA(status);

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'ok',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'CREATED',
      @message = N'migration step ok',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END

  COMMIT;
END TRY
BEGIN CATCH
  IF XACT_STATE() <> 0 ROLLBACK;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
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
'@
index_down = @'
/*
Severity mapping:
0=INFO (ok/skip normal)
1=NOTICE (schema change)
2=WARNING (follow-up/manual risk)
3=ERROR (stop)
*/
SET NOCOUNT ON;
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = 20260217073015001;
DECLARE @category NVARCHAR(16) = N'indexes';
DECLARE @file NVARCHAR(260) = N'20260217073015001_indexes_dbo.TableA_v1.down.sql';
DECLARE @target NVARCHAR(256) = N'dbo.TableA.IX_TableA__status';
DECLARE @action NVARCHAR(64) = N'drop_index';
DECLARE @direction NVARCHAR(8) = N'down';

IF NOT EXISTS (
  SELECT 1
  FROM sys.indexes
  WHERE object_id = OBJECT_ID(N'dbo.TableA')
    AND name = N'IX_TableA__status'
)
BEGIN
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'skip',
      @severity = 0,
      @action = @action,
      @target = @target,
      @reason = N'OBJECT_NOT_EXISTS',
      @message = N'[PRECHECK] index already absent',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END
  RETURN;
END

BEGIN TRY
  BEGIN TRAN;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'run',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'OBJECT_EXISTS',
      @message = N'migration run start',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END

  DROP INDEX IF EXISTS [IX_TableA__status] ON dbo.TableA;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'ok',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'DROPPED',
      @message = N'migration step ok',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END

  COMMIT;
END TRY
BEGIN CATCH
  IF XACT_STATE() <> 0 ROLLBACK;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
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
'@
    view_up = @'
CREATE OR ALTER VIEW dbo.ViewA
AS
SELECT id, status, amount
FROM dbo.TableA;
'@
    view_down = @'
CREATE OR ALTER VIEW dbo.ViewA
AS
SELECT id, status
FROM dbo.TableA;
'@
    function_up = @'
CREATE OR ALTER FUNCTION dbo.FuncA (@id INT)
RETURNS INT
AS
BEGIN
  RETURN @id;
END
'@
    function_down = @'
CREATE OR ALTER FUNCTION dbo.FuncA (@id INT)
RETURNS INT
AS
BEGIN
  RETURN @id;
END
'@
    procedure_up = @'
CREATE OR ALTER PROCEDURE dbo.ProcA
AS
BEGIN
  SET NOCOUNT ON;
  SELECT COUNT_BIG(1) AS total_rows FROM dbo.TableA;
END
'@
    procedure_down = @'
CREATE OR ALTER PROCEDURE dbo.ProcA
AS
BEGIN
  SET NOCOUNT ON;
  SELECT TOP (1) id FROM dbo.TableA ORDER BY id;
END
'@
  }
}

function Build-PatchWrapperSql {
  param(
    [string]$Version,
    [string]$FileName,
    [string]$PatchName,
    [string]$PatchFolder,
    [string]$Direction,
    [string]$BodySql
  )

  $bodyNormalized = ($BodySql -replace "`r`n", "`n").Trim()
  if ([string]::IsNullOrWhiteSpace($bodyNormalized)) {
    $bodyNormalized = "SELECT 1;"
  }
  $bodyIndented = (($bodyNormalized -split "`n") | ForEach-Object { "  " + $_.TrimEnd() }) -join "`n"
  $readmePath = "patch/$PatchFolder/README.md"

  return @"
/*
patch_folder: $PatchFolder
patch_name: $PatchName
readme: $readmePath
*/
/*
Severity mapping:
0=INFO (ok/skip normal)
1=NOTICE (schema change)
2=WARNING (follow-up/manual risk)
3=ERROR (stop)
*/
SET NOCOUNT ON;
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = $Version;
DECLARE @category NVARCHAR(16) = N'patch';
DECLARE @file NVARCHAR(260) = N'$FileName';
DECLARE @target NVARCHAR(256) = N'$PatchName';
DECLARE @action NVARCHAR(64) = N'backfill';
DECLARE @direction NVARCHAR(8) = N'$Direction';

-- no precheck

BEGIN TRY
  BEGIN TRAN;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'run',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'OBJECT_EXISTS',
      @message = N'migration run start',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END

$bodyIndented

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'ok',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'CREATED',
      @message = N'migration step ok',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END

  COMMIT TRAN;
END TRY
BEGIN CATCH
  IF XACT_STATE() <> 0 ROLLBACK TRAN;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
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
"@
}

function Initialize-Scenario {
  param([string]$Scenario)
  $sql = Baseline-Sql

  $base = Join-Path $root $Scenario
  $input = Join-Path $base "input"
  $expected = Join-Path $base "expected"

  foreach ($d in @(
      "$input/db_object/tables",
      "$input/db_object/indexes",
      "$input/db_object/views",
      "$input/db_object/functions",
      "$input/db_object/procedures",
      "$input/patch",
      "$input/migrations",
      "$expected"
    )) {
    New-Item -ItemType Directory -Force -Path $d | Out-Null
  }

  Write-LFFile -Path (Join-Path $input ".gitignore") -Content "out/**`n*.output.json`n"
  Write-LFFile -Path (Join-Path $input ".gitattributes") -Content "* text=auto eol=lf`n*.sql text eol=lf`n*.md text eol=lf`n*.go text eol=lf`n"
  Write-LFFile -Path (Join-Path $input "constraints_naming.md") -Content "# Constraints naming`nAll constraints must be deterministic.`n"

  Write-LFFile -Path (Join-Path $input "db_object/tables/dbo.TableA.sql") -Content $sql.table_source
  Write-LFFile -Path (Join-Path $input "db_object/indexes/dbo.TableA.sql") -Content $sql.index_source
  Write-LFFile -Path (Join-Path $input "db_object/views/dbo.ViewA.sql") -Content $sql.view_source
  Write-LFFile -Path (Join-Path $input "db_object/functions/dbo.FuncA.sql") -Content $sql.function_source
  Write-LFFile -Path (Join-Path $input "db_object/procedures/dbo.ProcA.sql") -Content $sql.procedure_source

  Write-LFFile -Path (Join-Path $input "migrations/000000__core__mig__audit_init.up.sql") -Content $sql.core_up
  Write-LFFile -Path (Join-Path $input "migrations/000000__core__mig__audit_init.down.sql") -Content $sql.core_down
  Write-LFFile -Path (Join-Path $input "migrations/20260217073015000_tables_dbo.TableA_v1.up.sql") -Content $sql.table_up
  Write-LFFile -Path (Join-Path $input "migrations/20260217073015000_tables_dbo.TableA_v1.down.sql") -Content $sql.table_down
  Write-LFFile -Path (Join-Path $input "migrations/20260217073015001_indexes_dbo.TableA_v1.up.sql") -Content $sql.index_up
  Write-LFFile -Path (Join-Path $input "migrations/20260217073015001_indexes_dbo.TableA_v1.down.sql") -Content $sql.index_down
  Write-LFFile -Path (Join-Path $input "migrations/20260217073015002_views_dbo.ViewA_v1.up.sql") -Content $sql.view_up
  Write-LFFile -Path (Join-Path $input "migrations/20260217073015002_views_dbo.ViewA_v1.down.sql") -Content $sql.view_down
  Write-LFFile -Path (Join-Path $input "migrations/20260217073015003_functions_dbo.FuncA_v1.up.sql") -Content $sql.function_up
  Write-LFFile -Path (Join-Path $input "migrations/20260217073015003_functions_dbo.FuncA_v1.down.sql") -Content $sql.function_down
  Write-LFFile -Path (Join-Path $input "migrations/20260217073015004_procedures_dbo.ProcA_v1.up.sql") -Content $sql.procedure_up
  Write-LFFile -Path (Join-Path $input "migrations/20260217073015004_procedures_dbo.ProcA_v1.down.sql") -Content $sql.procedure_down
}

function Write-ScenarioReadme {
  param(
    [string]$Scenario,
    [string]$QaId,
    [string]$Title,
    [string]$Command,
    [string]$Expected
  )
  $content = @"
# $QaId - $Title

## Purpose

Validate rule behavior for scenario $QaId.

## Command Under Test

- $Command

## Fixture Layout

- `input/`: repository state before command run
- `expected/`: expected result baseline

## Expected Outcome

- $Expected
"@
  Write-LFFile -Path (Join-Path $root "$Scenario/README.md") -Content $content
}

function Write-Expectation {
  param(
    [string]$Scenario,
    [string]$QaId,
    [string]$Command,
    [int]$ExitCode,
    [string[]]$FailRules,
    [string[]]$WarnReasons,
    [string]$Summary
  )

  $obj = [ordered]@{
    qa_id = $QaId
    command = $Command
    expected_exit_code = $ExitCode
    expected_fail_rules = $FailRules
    expected_warn_reasons = $WarnReasons
    summary = $Summary
  }
  $json = $obj | ConvertTo-Json -Depth 8
  Write-LFFile -Path (Join-Path $root "$Scenario/expected/expectation.json") -Content ($json + "`n")
}

$scenarioDefs = @(
  @{name="01_hard_fail_rule"; qa="QA-01"; title="Hard Fail Rule"; command="verify"; expected="fail"; exit=1; fail=@("FORBIDDEN_SQL_TOKEN"); warn=@(); summary="Severity 3 must force exit 1 and no emit."},
  @{name="02_folder_contract"; qa="QA-02"; title="Folder Contract"; command="verify"; expected="fail"; exit=1; fail=@("MIGRATIONS_NOT_FLAT"); warn=@(); summary="Migrations folder must be flat."},
  @{name="03_forbidden_sql"; qa="QA-03"; title="Forbidden SQL"; command="verify"; expected="fail"; exit=1; fail=@("FORBIDDEN_SQL_TOKEN"); warn=@(); summary="Forbidden SQL token detection."},
  @{name="04_pairing_missing_down"; qa="QA-04"; title="Pairing Missing Down"; command="verify"; expected="fail"; exit=1; fail=@("MISSING_MIGRATION_PAIR"); warn=@(); summary="Each up migration must have down pair."},
  @{name="05_version_format_invalid"; qa="QA-05"; title="Version Format Invalid"; command="verify"; expected="fail"; exit=1; fail=@("INVALID_VERSION_FORMAT"); warn=@(); summary="Version must be 17 digits."},
  @{name="06_ordering_determinism"; qa="QA-06"; title="Ordering Determinism"; command="gen"; expected="pass"; exit=0; fail=@(); warn=@(); summary="Emit order must follow category sequence."},
  @{name="07_vk_whitespace_only"; qa="QA-07"; title="v{k} Whitespace Only"; command="gen"; expected="pass"; exit=0; fail=@(); warn=@(); summary="Whitespace-only change must not create migration."},
  @{name="08_reconstruction_unsupported_sql"; qa="QA-08"; title="Reconstruction Unsupported SQL"; command="gen"; expected="fail"; exit=1; fail=@("UNSUPPORTED_SQL_SUBSET"); warn=@(); summary="Unsupported SQL in history must fail reconstruction."},
  @{name="09_file_object_mismatch"; qa="QA-09"; title="File Object Mismatch"; command="verify"; expected="fail"; exit=1; fail=@("FILE_OBJECT_MISMATCH"); warn=@(); summary="Filename and SQL object name must match."},
  @{name="10_inline_default_not_allowed"; qa="QA-10"; title="Inline Default Not Allowed"; command="verify"; expected="fail"; exit=1; fail=@("INLINE_DEFAULT_NOT_ALLOWED"); warn=@(); summary="Inline default in table snapshot is forbidden."},
  @{name="11_destructive_drop_table_manual"; qa="QA-11"; title="Destructive Drop Table Manual"; command="gen"; expected="pass-with-warn"; exit=0; fail=@(); warn=@("REQUIRES_MANUAL","DATA_RESTORE_NOT_GUARANTEED"); summary="DROP TABLE must be manual-only with warning trace."},
  @{name="12_transaction_error_throw"; qa="QA-12"; title="Transaction Error Throw"; command="verify"; expected="fail"; exit=1; fail=@("WRAPPER_THROW_REQUIRED"); warn=@(); summary="Tables/indexes/patch wrapper catch must end with THROW."},
  @{name="13_audit_wrapper_contract"; qa="QA-13"; title="Audit Wrapper Contract"; command="verify"; expected="pass"; exit=0; fail=@(); warn=@(); summary="Wrapper and audit metadata must be present."},
  @{name="14_down_data_restore_warning"; qa="QA-14"; title="Down Data Restore Warning"; command="gen"; expected="pass-with-warn"; exit=0; fail=@(); warn=@("DATA_RESTORE_NOT_GUARANTEED"); summary="Destructive down path must carry data restore warning."},
  @{name="15_patch_canonicalization"; qa="QA-15"; title="Patch Canonicalization"; command="gen"; expected="pass"; exit=0; fail=@(); warn=@(); summary="Patch folder name must canonicalize to patch_name token."},
  @{name="16_patch_duplication_once"; qa="QA-16"; title="Patch Duplication Once"; command="gen"; expected="pass-skip"; exit=0; fail=@(); warn=@(); summary="Existing patch_name migration must skip duplicate generation."},
  @{name="17_patch_source_integrity"; qa="QA-17"; title="Patch Source Integrity"; command="verify"; expected="fail"; exit=1; fail=@("PATCH_SOURCE_MISSING"); warn=@(); summary="Patch source must not disappear if migration exists."},
  @{name="18_output_json_contract"; qa="QA-18"; title="Output JSON Contract"; command="verify"; expected="pass"; exit=0; fail=@(); warn=@(); summary="Output report must include fail/pass/new_migrations/notes roots."},
  @{name="19_repo_hygiene_constraints"; qa="QA-19"; title="Repo Hygiene and Constraints"; command="verify"; expected="pass"; exit=0; fail=@(); warn=@(); summary="constraints_naming and LF hygiene files must exist."},
  @{name="20_cli_exit_code_contract"; qa="QA-20"; title="CLI Exit Code Contract"; command="fmt|verify|gen"; expected="pass-fail-matrix"; exit=0; fail=@(); warn=@(); summary="Exit code 0/1 matrix must match fail.count."},
  @{name="21_wrapper_action_whitelist"; qa="QA-21"; title="Wrapper Action Whitelist"; command="verify"; expected="fail"; exit=1; fail=@("WRAPPER_AUDIT_ACTION_INVALID"); warn=@(); summary="Generated-style wrapper (DECLARE @file + @file_name=@file) action must be in whitelist."},
  @{name="22_wrapper_reason_whitelist"; qa="QA-22"; title="Wrapper Reason Whitelist"; command="verify"; expected="fail"; exit=1; fail=@("WRAPPER_AUDIT_REASON_INVALID"); warn=@(); summary="Generated-style wrapper (DECLARE @file + @file_name=@file) reason must be in whitelist."},
  @{name="23_wrapper_run_key_newid"; qa="QA-23"; title="Wrapper Run Key NEWID"; command="verify"; expected="fail"; exit=1; fail=@("WRAPPER_RUN_KEY_INVALID"); warn=@(); summary="Generated-style wrapper (DECLARE @file + @file_name=@file) must initialize @run_key with NEWID()."},
  @{name="24_constraint_unnamed"; qa="QA-24"; title="Constraint Naming Informational"; command="verify"; expected="pass"; exit=0; fail=@(); warn=@(); summary="Constraint naming is informational only; unnamed/legacy constraint names are allowed."},
  @{name="25_wrapper_precheck_skip"; qa="QA-25"; title="Wrapper Precheck Skip"; command="gen"; expected="pass"; exit=0; fail=@(); warn=@(); summary="Generated wrappers must include deterministic precheck+skip blocks."},
  @{name="26_alter_column_precheck_map"; qa="QA-26"; title="Alter Column Precheck Map"; command="gen"; expected="pass"; exit=0; fail=@(); warn=@(); summary="Alter column emit must include deterministic precheck mapping for nullability/length/type conversion risks."}
)

foreach ($s in $scenarioDefs) {
  Initialize-Scenario -Scenario $s.name
  Write-ScenarioReadme -Scenario $s.name -QaId $s.qa -Title $s.title -Command $s.command -Expected $s.expected
  Write-Expectation -Scenario $s.name -QaId $s.qa -Command $s.command -ExitCode $s.exit -FailRules $s.fail -WarnReasons $s.warn -Summary $s.summary

  $inputRoot = Join-Path $root "$($s.name)/input"
  $expectedRoot = Join-Path $root "$($s.name)/expected"

  switch ($s.name) {
    "01_hard_fail_rule" {
      $p = Join-Path $inputRoot "db_object/tables/dbo.TableA.sql"
      Write-LFFile -Path $p -Content ((Get-Content $p -Raw) + "`nGO`n")
    }
    "02_folder_contract" {
      New-Item -ItemType Directory -Force -Path (Join-Path $inputRoot "migrations/archive") | Out-Null
      Write-LFFile -Path (Join-Path $inputRoot "migrations/archive/20260217073015999_tables_dbo.Bad_v1.up.sql") -Content "SELECT 1;`n"
    }
    "03_forbidden_sql" {
      New-Item -ItemType Directory -Force -Path (Join-Path $inputRoot "patch/Forbidden SQL") | Out-Null
      Write-LFFile -Path (Join-Path $inputRoot "patch/Forbidden SQL/README.md") -Content "Forbidden SQL patch fixture.`n"
      Write-LFFile -Path (Join-Path $inputRoot "patch/Forbidden SQL/up.sql") -Content "USE master;`nSELECT 1;`n"
      Write-LFFile -Path (Join-Path $inputRoot "patch/Forbidden SQL/down.sql") -Content "SELECT 1;`n"
    }
    "04_pairing_missing_down" {
      Remove-Item -Force (Join-Path $inputRoot "migrations/20260217073015001_indexes_dbo.TableA_v1.down.sql")
    }
    "05_version_format_invalid" {
      Write-LFFile -Path (Join-Path $inputRoot "migrations/2026021707301500_tables_dbo.Bad_v1.up.sql") -Content "SELECT 1;`n"
      Write-LFFile -Path (Join-Path $inputRoot "migrations/2026021707301500_tables_dbo.Bad_v1.down.sql") -Content "SELECT 1;`n"
    }
    "06_ordering_determinism" {
      Get-ChildItem -Path (Join-Path $inputRoot "migrations") -File | Where-Object { $_.Name -notlike "000000__core__mig__audit_init.*.sql" } | Remove-Item -Force
      New-Item -ItemType Directory -Force -Path (Join-Path $inputRoot "patch/Fix Status") | Out-Null
      Write-LFFile -Path (Join-Path $inputRoot "patch/Fix Status/README.md") -Content "Patch for ordering scenario.`n"
      Write-LFFile -Path (Join-Path $inputRoot "patch/Fix Status/up.sql") -Content "UPDATE dbo.TableA SET status = 'DONE' WHERE status = 'NEW';`n"
      Write-LFFile -Path (Join-Path $inputRoot "patch/Fix Status/down.sql") -Content "UPDATE dbo.TableA SET status = 'NEW' WHERE status = 'DONE';`n"
      Write-LFFile -Path (Join-Path $expectedRoot "generated_order.txt") -Content @'
<ts>000_tables_dbo.TableA_v1.up.sql
<ts>001_indexes_dbo.TableA_v1.up.sql
<ts>002_views_dbo.ViewA_v1.up.sql
<ts>003_functions_dbo.FuncA_v1.up.sql
<ts>004_procedures_dbo.ProcA_v1.up.sql
<ts>005_patch_fix_status_v1.up.sql
'@
    }
    "07_vk_whitespace_only" {
      Write-LFFile -Path (Join-Path $inputRoot "db_object/tables/dbo.TableA.sql") -Content @'
CREATE TABLE dbo.TableA (
    id INT NOT NULL,
    status VARCHAR(10) NULL,
    amount DECIMAL(10,2) NULL,
    CONSTRAINT PK_TableA PRIMARY KEY (id),
    CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status
);
'@
      Write-LFFile -Path (Join-Path $expectedRoot "expected_new_migrations.txt") -Content "0`n"
    }
    "08_reconstruction_unsupported_sql" {
      Write-LFFile -Path (Join-Path $inputRoot "migrations/20260217073015010_tables_dbo.Unsupported_v1.up.sql") -Content @'
CREATE TABLE dbo.Unsupported (
  id INT NOT NULL,
  computed_col AS (id + 1)
);
'@
      Write-LFFile -Path (Join-Path $inputRoot "migrations/20260217073015010_tables_dbo.Unsupported_v1.down.sql") -Content "DROP TABLE dbo.Unsupported;`n"
    }
    "09_file_object_mismatch" {
      Write-LFFile -Path (Join-Path $inputRoot "db_object/tables/dbo.TableA.sql") -Content @'
CREATE TABLE dbo.TableB (
  id INT NOT NULL,
  CONSTRAINT PK_TableB PRIMARY KEY (id)
);
'@
    }
    "10_inline_default_not_allowed" {
      Write-LFFile -Path (Join-Path $inputRoot "db_object/tables/dbo.TableA.sql") -Content @'
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL DEFAULT ('NEW'),
  CONSTRAINT PK_TableA PRIMARY KEY (id)
);
'@
    }
    "11_destructive_drop_table_manual" {
      Remove-Item -Force (Join-Path $inputRoot "db_object/tables/dbo.TableA.sql")
      Remove-Item -Force (Join-Path $inputRoot "db_object/indexes/dbo.TableA.sql")
      Write-LFFile -Path (Join-Path $expectedRoot "expected_warn_reasons.txt") -Content "REQUIRES_MANUAL`nDATA_RESTORE_NOT_GUARANTEED`n"
    }
    "12_transaction_error_throw" {
      $p = Join-Path $inputRoot "migrations/20260217073015000_tables_dbo.TableA_v1.up.sql"
      $t = Get-Content $p -Raw
      $t = $t -replace "`n\s*THROW;", "`n  -- THROW intentionally removed"
      Write-LFFile -Path $p -Content $t
    }
    "13_audit_wrapper_contract" {
      Write-LFFile -Path (Join-Path $expectedRoot "checks.txt") -Content "severity_header`nrun_key`nversion_category_file`naudit_call`n"
    }
    "14_down_data_restore_warning" {
      Write-LFFile -Path (Join-Path $inputRoot "db_object/tables/dbo.TableA.sql") -Content @'
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id),
  CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status
);
'@
      Write-LFFile -Path (Join-Path $expectedRoot "expected_warn_reasons.txt") -Content "DATA_RESTORE_NOT_GUARANTEED`n"
    }
    "15_patch_canonicalization" {
      New-Item -ItemType Directory -Force -Path (Join-Path $inputRoot "patch/Fix Status!!") | Out-Null
      Write-LFFile -Path (Join-Path $inputRoot "patch/Fix Status!!/README.md") -Content "Patch canonicalization fixture.`n"
      Write-LFFile -Path (Join-Path $inputRoot "patch/Fix Status!!/up.sql") -Content "UPDATE dbo.TableA SET status = 'DONE' WHERE status = 'NEW';`n"
      Write-LFFile -Path (Join-Path $inputRoot "patch/Fix Status!!/down.sql") -Content "UPDATE dbo.TableA SET status = 'NEW' WHERE status = 'DONE';`n"
      Write-LFFile -Path (Join-Path $expectedRoot "expected_patch_name.txt") -Content "fix_status`n"
    }
    "16_patch_duplication_once" {
      New-Item -ItemType Directory -Force -Path (Join-Path $inputRoot "patch/Fix Status") | Out-Null
      Write-LFFile -Path (Join-Path $inputRoot "patch/Fix Status/README.md") -Content "Existing patch should be skipped on generate.`n"
      Write-LFFile -Path (Join-Path $inputRoot "patch/Fix Status/up.sql") -Content "UPDATE dbo.TableA SET status = 'DONE' WHERE status = 'NEW';`n"
      Write-LFFile -Path (Join-Path $inputRoot "patch/Fix Status/down.sql") -Content "UPDATE dbo.TableA SET status = 'NEW' WHERE status = 'DONE';`n"
      $patchUp = Build-PatchWrapperSql `
        -Version "20260217073015010" `
        -FileName "20260217073015010_patch_fix_status_v1.up.sql" `
        -PatchName "fix_status" `
        -PatchFolder "Fix Status" `
        -Direction "up" `
        -BodySql "UPDATE dbo.TableA SET status = 'DONE' WHERE status = 'NEW';"
      $patchDown = Build-PatchWrapperSql `
        -Version "20260217073015010" `
        -FileName "20260217073015010_patch_fix_status_v1.down.sql" `
        -PatchName "fix_status" `
        -PatchFolder "Fix Status" `
        -Direction "down" `
        -BodySql "UPDATE dbo.TableA SET status = 'NEW' WHERE status = 'DONE';"
      Write-LFFile -Path (Join-Path $inputRoot "migrations/20260217073015010_patch_fix_status_v1.up.sql") -Content $patchUp
      Write-LFFile -Path (Join-Path $inputRoot "migrations/20260217073015010_patch_fix_status_v1.down.sql") -Content $patchDown
    }
    "17_patch_source_integrity" {
      $patchUp = Build-PatchWrapperSql `
        -Version "20260217073015011" `
        -FileName "20260217073015011_patch_fix_status_v1.up.sql" `
        -PatchName "fix_status" `
        -PatchFolder "Fix Status" `
        -Direction "up" `
        -BodySql "SELECT 1;"
      $patchDown = Build-PatchWrapperSql `
        -Version "20260217073015011" `
        -FileName "20260217073015011_patch_fix_status_v1.down.sql" `
        -PatchName "fix_status" `
        -PatchFolder "Fix Status" `
        -Direction "down" `
        -BodySql "SELECT 1;"
      Write-LFFile -Path (Join-Path $inputRoot "migrations/20260217073015011_patch_fix_status_v1.up.sql") -Content $patchUp
      Write-LFFile -Path (Join-Path $inputRoot "migrations/20260217073015011_patch_fix_status_v1.down.sql") -Content $patchDown
    }
    "18_output_json_contract" {
      Write-LFFile -Path (Join-Path $expectedRoot "output_contract_keys.txt") -Content "fail`npass`nnew_migrations`nnotes`n"
    }
    "19_repo_hygiene_constraints" {
      Write-LFFile -Path (Join-Path $expectedRoot "required_files.txt") -Content "constraints_naming.md`n.gitignore`n.gitattributes`n"
    }
    "20_cli_exit_code_contract" {
      New-Item -ItemType Directory -Force -Path (Join-Path $inputRoot "cases/pass"), (Join-Path $inputRoot "cases/fail") | Out-Null
      Write-LFFile -Path (Join-Path $inputRoot "cases/pass/README.md") -Content "Reference pass case from QA-18.`n"
      Write-LFFile -Path (Join-Path $inputRoot "cases/fail/README.md") -Content "Reference fail case from QA-10.`n"
      $matrix = [ordered]@{
        fmt = @{ pass_exit = 0; fail_exit = 1 }
        verify = @{ pass_exit = 0; fail_exit = 1 }
        gen = @{ pass_exit = 0; fail_exit = 1 }
      } | ConvertTo-Json -Depth 8
      Write-LFFile -Path (Join-Path $expectedRoot "cli_exit_matrix.json") -Content ($matrix + "`n")
    }
    "21_wrapper_action_whitelist" {
      Write-LFFile -Path (Join-Path $inputRoot "migrations/20260217073015100_tables_dbo.TableB_v1.up.sql") -Content @'
/*
Severity mapping:
0=INFO (ok/skip normal)
1=NOTICE (schema change)
2=WARNING (follow-up/manual risk)
3=ERROR (stop)
*/
SET NOCOUNT ON;
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = 20260217073015100;
DECLARE @category NVARCHAR(16) = N'tables';
DECLARE @file NVARCHAR(260) = N'20260217073015100_tables_dbo.TableB_v1.up.sql';
DECLARE @target NVARCHAR(256) = N'dbo.TableB';
DECLARE @action NVARCHAR(64) = N'UNKNOWN_ACTION';
DECLARE @direction NVARCHAR(8) = N'up';

BEGIN TRY
  BEGIN TRAN;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'run',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'OBJECT_EXISTS',
      @message = N'migration run start',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END

  SELECT 1;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'ok',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'CREATED',
      @message = N'migration step ok',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END

  COMMIT TRAN;
END TRY
BEGIN CATCH
  IF XACT_STATE() <> 0 ROLLBACK TRAN;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
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
'@
      Write-LFFile -Path (Join-Path $inputRoot "migrations/20260217073015100_tables_dbo.TableB_v1.down.sql") -Content @'
/*
Severity mapping:
0=INFO (ok/skip normal)
1=NOTICE (schema change)
2=WARNING (follow-up/manual risk)
3=ERROR (stop)
*/
SET NOCOUNT ON;
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = 20260217073015100;
DECLARE @category NVARCHAR(16) = N'tables';
DECLARE @file NVARCHAR(260) = N'20260217073015100_tables_dbo.TableB_v1.down.sql';
DECLARE @target NVARCHAR(256) = N'dbo.TableB';
DECLARE @action NVARCHAR(64) = N'UNKNOWN_ACTION';
DECLARE @direction NVARCHAR(8) = N'down';

BEGIN TRY
  BEGIN TRAN;
  SELECT 1;
  COMMIT TRAN;
END TRY
BEGIN CATCH
  IF XACT_STATE() <> 0 ROLLBACK TRAN;
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
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
'@
    }
    "22_wrapper_reason_whitelist" {
      Write-LFFile -Path (Join-Path $inputRoot "migrations/20260217073015101_tables_dbo.TableB_v1.up.sql") -Content @'
/*
Severity mapping:
0=INFO (ok/skip normal)
1=NOTICE (schema change)
2=WARNING (follow-up/manual risk)
3=ERROR (stop)
*/
SET NOCOUNT ON;
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = 20260217073015101;
DECLARE @category NVARCHAR(16) = N'tables';
DECLARE @file NVARCHAR(260) = N'20260217073015101_tables_dbo.TableB_v1.up.sql';
DECLARE @target NVARCHAR(256) = N'dbo.TableB';
DECLARE @action NVARCHAR(64) = N'create_table';
DECLARE @direction NVARCHAR(8) = N'up';

BEGIN TRY
  BEGIN TRAN;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'run',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'UNKNOWN_REASON',
      @message = N'migration run start',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END

  SELECT 1;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'ok',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'CREATED',
      @message = N'migration step ok',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END

  COMMIT TRAN;
END TRY
BEGIN CATCH
  IF XACT_STATE() <> 0 ROLLBACK TRAN;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
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
'@
      Write-LFFile -Path (Join-Path $inputRoot "migrations/20260217073015101_tables_dbo.TableB_v1.down.sql") -Content @'
/*
Severity mapping:
0=INFO (ok/skip normal)
1=NOTICE (schema change)
2=WARNING (follow-up/manual risk)
3=ERROR (stop)
*/
SET NOCOUNT ON;
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = 20260217073015101;
DECLARE @category NVARCHAR(16) = N'tables';
DECLARE @file NVARCHAR(260) = N'20260217073015101_tables_dbo.TableB_v1.down.sql';
DECLARE @target NVARCHAR(256) = N'dbo.TableB';
DECLARE @action NVARCHAR(64) = N'create_table';
DECLARE @direction NVARCHAR(8) = N'down';

BEGIN TRY
  BEGIN TRAN;
  SELECT 1;
  COMMIT TRAN;
END TRY
BEGIN CATCH
  IF XACT_STATE() <> 0 ROLLBACK TRAN;
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
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
'@
    }
    "23_wrapper_run_key_newid" {
      Write-LFFile -Path (Join-Path $inputRoot "migrations/20260217073015102_tables_dbo.TableB_v1.up.sql") -Content @'
/*
Severity mapping:
0=INFO (ok/skip normal)
1=NOTICE (schema change)
2=WARNING (follow-up/manual risk)
3=ERROR (stop)
*/
SET NOCOUNT ON;
DECLARE @run_key UNIQUEIDENTIFIER = NULL;
DECLARE @version BIGINT = 20260217073015102;
DECLARE @category NVARCHAR(16) = N'tables';
DECLARE @file NVARCHAR(260) = N'20260217073015102_tables_dbo.TableB_v1.up.sql';
DECLARE @target NVARCHAR(256) = N'dbo.TableB';
DECLARE @action NVARCHAR(64) = N'create_table';
DECLARE @direction NVARCHAR(8) = N'up';

BEGIN TRY
  BEGIN TRAN;
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'run',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'OBJECT_EXISTS',
      @message = N'migration run start',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END
  SELECT 1;
  COMMIT TRAN;
END TRY
BEGIN CATCH
  IF XACT_STATE() <> 0 ROLLBACK TRAN;
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
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
'@
      Write-LFFile -Path (Join-Path $inputRoot "migrations/20260217073015102_tables_dbo.TableB_v1.down.sql") -Content @'
/*
Severity mapping:
0=INFO (ok/skip normal)
1=NOTICE (schema change)
2=WARNING (follow-up/manual risk)
3=ERROR (stop)
*/
SET NOCOUNT ON;
DECLARE @run_key UNIQUEIDENTIFIER = NULL;
DECLARE @version BIGINT = 20260217073015102;
DECLARE @category NVARCHAR(16) = N'tables';
DECLARE @file NVARCHAR(260) = N'20260217073015102_tables_dbo.TableB_v1.down.sql';
DECLARE @target NVARCHAR(256) = N'dbo.TableB';
DECLARE @action NVARCHAR(64) = N'create_table';
DECLARE @direction NVARCHAR(8) = N'down';

BEGIN TRY
  BEGIN TRAN;
  SELECT 1;
  COMMIT TRAN;
END TRY
BEGIN CATCH
  IF XACT_STATE() <> 0 ROLLBACK TRAN;
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
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
'@
    }
    "24_constraint_unnamed" {
      Write-LFFile -Path (Join-Path $inputRoot "db_object/tables/dbo.TableA.sql") -Content @'
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL,
  amount DECIMAL(10,2) NULL,
  PRIMARY KEY (id),
  CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status
);
'@
    }
    "25_wrapper_precheck_skip" {
      Get-ChildItem -Path (Join-Path $inputRoot "migrations") -File | Where-Object { $_.Name -notlike "000000__core__mig__audit_init.*.sql" } | Remove-Item -Force
      New-Item -ItemType Directory -Force -Path (Join-Path $inputRoot "patch/Fix Status") | Out-Null
      Write-LFFile -Path (Join-Path $inputRoot "patch/Fix Status/README.md") -Content "Patch for wrapper precheck scenario.`n"
      Write-LFFile -Path (Join-Path $inputRoot "patch/Fix Status/up.sql") -Content "UPDATE dbo.TableA SET status = 'DONE' WHERE status = 'NEW';`n"
      Write-LFFile -Path (Join-Path $inputRoot "patch/Fix Status/down.sql") -Content "UPDATE dbo.TableA SET status = 'NEW' WHERE status = 'DONE';`n"
    }
    "26_alter_column_precheck_map" {
      Write-LFFile -Path (Join-Path $inputRoot "db_object/tables/dbo.TableA.sql") -Content @'
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL,
  amount DECIMAL(10,2) NULL,
  code VARCHAR(5) NOT NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id),
  CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status
);
'@

      Write-LFFile -Path (Join-Path $inputRoot "migrations/20260217073015000_tables_dbo.TableA_v1.up.sql") -Content @'
/*
Severity mapping:
0=INFO (ok/skip normal)
1=NOTICE (schema change)
2=WARNING (follow-up/manual risk)
3=ERROR (stop)
*/
SET NOCOUNT ON;
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = 20260217073015000;
DECLARE @category NVARCHAR(16) = N'tables';
DECLARE @file NVARCHAR(260) = N'20260217073015000_tables_dbo.TableA_v1.up.sql';
DECLARE @target NVARCHAR(256) = N'dbo.TableA';
DECLARE @action NVARCHAR(64) = N'create_table';
DECLARE @direction NVARCHAR(8) = N'up';

IF OBJECT_ID(N'dbo.TableA', N'U') IS NOT NULL
BEGIN
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'skip',
      @severity = 0,
      @action = @action,
      @target = @target,
      @reason = N'TABLE_EXISTS',
      @message = N'[PRECHECK] table already exists',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END
  RETURN;
END

BEGIN TRY
  BEGIN TRAN;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'run',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'OBJECT_EXISTS',
      @message = N'migration run start',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END

  CREATE TABLE dbo.TableA (
    id INT NOT NULL,
    status VARCHAR(10) NULL,
    amount VARCHAR(20) NULL,
    code VARCHAR(20) NULL,
    CONSTRAINT PK_TableA PRIMARY KEY (id),
    CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status
  );

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'ok',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'CREATED',
      @message = N'migration step ok',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END

  COMMIT;
END TRY
BEGIN CATCH
  IF XACT_STATE() <> 0 ROLLBACK;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
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
'@
    }
  }
}

# Keep top-level index updated.
$indexLines = @(
  "# QA Fixtures - Engine Rules",
  "",
  'Direktori ini berisi scaffold fixture untuk validasi QA-01 s.d. QA-26 sesuai `docs/qa/qa_gate_engine_rules.md`.',
  'Catatan: QA-17B (Patch Immutability) diuji terpisah oleh `tests/engine/patch_immutability_test.go`.',
  "",
  "| QA | Scenario Folder | Expected Outcome |",
  "| --- | --- | --- |"
)

foreach ($s in $scenarioDefs) {
  $indexLines += "| $($s.qa) | ``$($s.name)`` | $($s.expected) |"
}

Write-LFFile -Path (Join-Path $root "README.md") -Content (($indexLines -join "`n") + "`n")
