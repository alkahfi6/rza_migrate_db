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
DECLARE @file NVARCHAR(260) = N'20260303172244000_tables_dbo.TableA_v2.up.sql';
DECLARE @target NVARCHAR(256) = N'dbo.TableA';
DECLARE @action NVARCHAR(64) = N'alter_column';
DECLARE @direction NVARCHAR(8) = N'up';

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

  DECLARE @skip_alter_amount BIT = 0;

  DECLARE @bad_amount BIGINT;
  SELECT @bad_amount = COUNT_BIG(1)
  FROM dbo.TableA
  WHERE [amount] IS NOT NULL
    AND TRY_CONVERT(DECIMAL(10,2), [amount]) IS NULL;

  IF (@bad_amount > 0)
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
        @target = N'dbo.TableA.amount',
        @reason = N'REQUIRES_MANUAL',
        @message = N'[PRECHECK] TYPE change blocked (non-convertible)',
        @detail = CONCAT(N'bad_cnt=', @bad_amount),
        @error_number = NULL,
        @error_line = NULL,
        @error_procedure = NULL;
    END
    SET @skip_alter_amount = 1;
  END


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

    ALTER TABLE dbo.TableA ALTER COLUMN [amount] DECIMAL(10,2) NULL;

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

  DECLARE @skip_alter_code BIT = 0;

  DECLARE @null_cnt_code BIGINT;
  SELECT @null_cnt_code = COUNT_BIG(1)
  FROM dbo.TableA
  WHERE [code] IS NULL;

  IF (@null_cnt_code > 0)
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
        @reason = N'REQUIRES_MANUAL',
        @message = N'[PRECHECK] ALTER COLUMN NOT NULL blocked',
        @detail = CONCAT(N'null_cnt=', @null_cnt_code),
        @error_number = NULL,
        @error_line = NULL,
        @error_procedure = NULL;
    END
    SET @skip_alter_code = 1;
  END


  DECLARE @too_long_code BIGINT;
  SELECT @too_long_code = COUNT_BIG(1)
  FROM dbo.TableA
  WHERE [code] IS NOT NULL
    AND DATALENGTH([code]) > 5;

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
        @reason = N'NARROWING_WOULD_TRUNCATE',
        @message = N'[PRECHECK] SHRINK length blocked',
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

    ALTER TABLE dbo.TableA ALTER COLUMN [code] VARCHAR(5) NOT NULL;

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
