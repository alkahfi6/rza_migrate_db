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