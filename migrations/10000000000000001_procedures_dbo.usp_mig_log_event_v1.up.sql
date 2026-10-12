CREATE OR ALTER PROCEDURE dbo.usp_mig_log_event
  @run_key UNIQUEIDENTIFIER,
  @operation_id UNIQUEIDENTIFIER = NULL,
  @version BIGINT,
  @category NVARCHAR(16),
  @file_name NVARCHAR(260),
  @direction NVARCHAR(8) = NULL,
  @content_sha256 CHAR(64) = NULL,
  @decision NVARCHAR(8),
  @severity TINYINT,
  @action NVARCHAR(64),
  @target NVARCHAR(256),
  @reason NVARCHAR(64),
  @message NVARCHAR(4000),
  @detail NVARCHAR(MAX) = NULL,
  @error_number INT = NULL,
  @error_line INT = NULL,
  @error_procedure NVARCHAR(128) = NULL,
  @started_at DATETIME2(7) = NULL,
  @duration_ms BIGINT = NULL
AS
BEGIN
  SET NOCOUNT ON;
  SET @operation_id = COALESCE(@operation_id, TRY_CONVERT(UNIQUEIDENTIFIER, SESSION_CONTEXT(N'mig_operation_id')));
  SET @direction = COALESCE(NULLIF(@direction, N''), CASE WHEN @detail IN (N'up', N'down') THEN @detail ELSE N'unknown' END);
  SET @direction = COALESCE(NULLIF(@direction, N'unknown'), TRY_CONVERT(NVARCHAR(8), SESSION_CONTEXT(N'mig_direction')), N'unknown');
  SET @content_sha256 = COALESCE(@content_sha256, TRY_CONVERT(CHAR(64), SESSION_CONTEXT(N'mig_content_sha256')));
  SET @started_at = COALESCE(@started_at, TRY_CONVERT(DATETIME2(7), SESSION_CONTEXT(N'mig_started_at')));
  IF @duration_ms IS NULL AND @started_at IS NOT NULL
    SET @duration_ms = DATEDIFF_BIG(MILLISECOND, @started_at, SYSUTCDATETIME());
  INSERT INTO dbo.mig_event (
    run_key, operation_id, version, category, file_name, direction, content_sha256,
    decision, severity, action, target, reason, message, detail,
    error_number, error_line, error_procedure, started_at, duration_ms
  )
  VALUES (
    @run_key, @operation_id, @version, @category, @file_name, @direction, @content_sha256,
    @decision, @severity, @action, @target, @reason, @message, @detail,
    @error_number, @error_line, @error_procedure, @started_at, @duration_ms
  );
END
