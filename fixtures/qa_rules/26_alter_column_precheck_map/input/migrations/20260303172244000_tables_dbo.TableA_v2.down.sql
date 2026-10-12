/*
Severity mapping:
0=INFO (ok/skip normal)
1=NOTICE (schema change)
2=WARNING (follow-up/manual risk)
3=ERROR (stop)
*/
SET NOCOUNT ON;
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = 20260303172244000;
DECLARE @category NVARCHAR(16) = N'tables';
DECLARE @file NVARCHAR(260) = N'20260303172244000_tables_dbo.TableA_v2.down.sql';
DECLARE @target NVARCHAR(256) = N'dbo.TableA';
DECLARE @action NVARCHAR(64) = N'alter_column';
DECLARE @direction NVARCHAR(8) = N'down';

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

  DECLARE @skip_alter_code BIT = 0;

  DECLARE @too_long_code BIGINT;
  SELECT @too_long_code = COUNT_BIG(1)
  FROM dbo.TableA
  WHERE [code] IS NOT NULL
    AND DATALENGTH([code]) > 20;

  IF (@too_long_code > 0)
  BEGIN
    IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
    BEGIN
      EXEC dbo.usp_mig_log_event
        @run_key = @run_key,
        @version = @version,
        @category = @category,
        @file_name = @file,
        @decision = N'skip',
        @severity = 2,
        @action = N'alter_column',
        @target = N'dbo.TableA.code',
        @reason = N'DATA_TOO_LONG_FOR_TARGET',
        @message = N'[PRECHECK] target length blocked',
        @detail = CONCAT(N'too_long=', @too_long_code),
        @error_number = NULL,
        @error_line = NULL,
        @error_procedure = NULL;
    END
    SET @skip_alter_code = 1;
  END


  IF (@skip_alter_code = 0)
  BEGIN
    IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
    BEGIN
      EXEC dbo.usp_mig_log_event
        @run_key = @run_key,
        @version = @version,
        @category = @category,
        @file_name = @file,
        @decision = N'run',
        @severity = 1,
        @action = N'alter_column',
        @target = N'dbo.TableA.code',
        @reason = N'OBJECT_EXISTS',
        @message = N'alter column start',
        @detail = @direction,
        @error_number = NULL,
        @error_line = NULL,
        @error_procedure = NULL;
    END

    ALTER TABLE dbo.TableA ALTER COLUMN [code] VARCHAR(20) NULL;

    IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
    BEGIN
      EXEC dbo.usp_mig_log_event
        @run_key = @run_key,
        @version = @version,
        @category = @category,
        @file_name = @file,
        @decision = N'ok',
        @severity = 1,
        @action = N'alter_column',
        @target = N'dbo.TableA.code',
        @reason = N'ALTERED',
        @message = N'alter column ok',
        @detail = @direction,
        @error_number = NULL,
        @error_line = NULL,
        @error_procedure = NULL;
    END
  END

  DECLARE @skip_alter_amount BIT = 0;

  IF (@skip_alter_amount = 0)
  BEGIN
    IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
    BEGIN
      EXEC dbo.usp_mig_log_event
        @run_key = @run_key,
        @version = @version,
        @category = @category,
        @file_name = @file,
        @decision = N'run',
        @severity = 1,
        @action = N'alter_column',
        @target = N'dbo.TableA.amount',
        @reason = N'OBJECT_EXISTS',
        @message = N'alter column start',
        @detail = @direction,
        @error_number = NULL,
        @error_line = NULL,
        @error_procedure = NULL;
    END

    ALTER TABLE dbo.TableA ALTER COLUMN [amount] VARCHAR(20) NULL;

    IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
    BEGIN
      EXEC dbo.usp_mig_log_event
        @run_key = @run_key,
        @version = @version,
        @category = @category,
        @file_name = @file,
        @decision = N'ok',
        @severity = 1,
        @action = N'alter_column',
        @target = N'dbo.TableA.amount',
        @reason = N'ALTERED',
        @message = N'alter column ok',
        @detail = @direction,
        @error_number = NULL,
        @error_line = NULL,
        @error_procedure = NULL;
    END
  END


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
