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