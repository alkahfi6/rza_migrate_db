IF OBJECT_ID(N'dbo.mig_event', N'U') IS NULL
BEGIN
  CREATE TABLE dbo.mig_event (
    event_id BIGINT IDENTITY(1,1) NOT NULL CONSTRAINT PK_mig_event PRIMARY KEY,
    run_key UNIQUEIDENTIFIER NOT NULL,
    operation_id UNIQUEIDENTIFIER NULL,
    version BIGINT NOT NULL,
    category NVARCHAR(16) COLLATE Latin1_General_BIN NOT NULL,
    file_name NVARCHAR(260) COLLATE Latin1_General_BIN NOT NULL,
    direction NVARCHAR(8) COLLATE Latin1_General_BIN NOT NULL,
    content_sha256 CHAR(64) COLLATE Latin1_General_BIN NULL,
    decision NVARCHAR(8) COLLATE Latin1_General_BIN NOT NULL,
    severity TINYINT NOT NULL,
    action NVARCHAR(64) COLLATE Latin1_General_BIN NOT NULL,
    target NVARCHAR(256) COLLATE Latin1_General_BIN NOT NULL,
    reason NVARCHAR(64) COLLATE Latin1_General_BIN NOT NULL,
    message NVARCHAR(4000) COLLATE Latin1_General_BIN NOT NULL,
    detail NVARCHAR(MAX) COLLATE Latin1_General_BIN NULL,
    error_number INT NULL,
    error_line INT NULL,
    error_procedure NVARCHAR(128) COLLATE Latin1_General_BIN NULL,
    started_at DATETIME2(7) NULL,
    duration_ms BIGINT NULL,
    created_at DATETIME2(7) NOT NULL CONSTRAINT DF_mig_event__created_at DEFAULT (SYSUTCDATETIME()),
    CONSTRAINT CK_mig_event__direction CHECK (direction IN (N'up', N'down')),
    CONSTRAINT CK_mig_event__decision CHECK (decision IN (N'run', N'ok', N'skip', N'fail')),
    CONSTRAINT CK_mig_event__severity CHECK (severity BETWEEN 0 AND 3),
    CONSTRAINT CK_mig_event__content_sha256 CHECK (
      content_sha256 IS NULL OR content_sha256 NOT LIKE '%[^0-9a-f]%' COLLATE Latin1_General_BIN
    )
  );

  CREATE INDEX IX_mig_event__run_key ON dbo.mig_event (run_key, created_at);
  CREATE INDEX IX_mig_event__version ON dbo.mig_event (version, created_at);
  CREATE INDEX IX_mig_event__operation_id ON dbo.mig_event (operation_id, created_at);
END

IF (
  SELECT COUNT_BIG(*) FROM sys.columns
  WHERE object_id = OBJECT_ID(N'dbo.mig_event', N'U')
) <> 21
OR NOT EXISTS (
  SELECT 1 FROM sys.columns
  WHERE object_id = OBJECT_ID(N'dbo.mig_event', N'U')
    AND name COLLATE Latin1_General_100_BIN2 = N'operation_id' COLLATE Latin1_General_100_BIN2
    AND TYPE_NAME(user_type_id) = N'uniqueidentifier'
    AND is_nullable = 1
)
OR NOT EXISTS (
  SELECT 1 FROM sys.columns
  WHERE object_id = OBJECT_ID(N'dbo.mig_event', N'U')
    AND name COLLATE Latin1_General_100_BIN2 = N'content_sha256' COLLATE Latin1_General_100_BIN2
    AND TYPE_NAME(user_type_id) = N'char'
    AND max_length = 64
)
OR NOT EXISTS (
  SELECT 1 FROM sys.indexes
  WHERE object_id = OBJECT_ID(N'dbo.mig_event', N'U')
    AND name COLLATE Latin1_General_100_BIN2 = N'IX_mig_event__operation_id' COLLATE Latin1_General_100_BIN2
)
BEGIN
  ;THROW 51000, N'Core precheck blocked: dbo.mig_event exists but does not match the mandatory audit contract', 1;
END

IF OBJECT_ID(N'dbo.mig_operation_journal', N'U') IS NULL
BEGIN
  CREATE TABLE dbo.mig_operation_journal (
    operation_id UNIQUEIDENTIFIER NOT NULL CONSTRAINT PK_mig_operation_journal PRIMARY KEY,
    request_id NVARCHAR(128) COLLATE Latin1_General_BIN NOT NULL,
    operation NVARCHAR(64) COLLATE Latin1_General_BIN NOT NULL,
    target_request_id NVARCHAR(128) COLLATE Latin1_General_BIN NULL,
    file_name NVARCHAR(260) COLLATE Latin1_General_BIN NULL,
    status NVARCHAR(16) COLLATE Latin1_General_BIN NOT NULL,
    requested_steps INT NOT NULL CONSTRAINT DF_mig_operation_journal__requested_steps DEFAULT (0),
    executed_steps INT NOT NULL CONSTRAINT DF_mig_operation_journal__executed_steps DEFAULT (0),
    before_version BIGINT NULL,
    before_dirty BIT NOT NULL CONSTRAINT DF_mig_operation_journal__before_dirty DEFAULT (0),
    after_version BIGINT NULL,
    after_dirty BIT NOT NULL CONSTRAINT DF_mig_operation_journal__after_dirty DEFAULT (0),
    result_direction NVARCHAR(8) COLLATE Latin1_General_BIN NOT NULL CONSTRAINT DF_mig_operation_journal__result_direction DEFAULT (N'none'),
    applied_files NVARCHAR(MAX) COLLATE Latin1_General_BIN NULL,
    error_message NVARCHAR(MAX) COLLATE Latin1_General_BIN NULL,
    nik NVARCHAR(64) COLLATE Latin1_General_BIN NULL,
    trace_id NVARCHAR(64) COLLATE Latin1_General_BIN NULL,
    manifest_sha256 CHAR(64) COLLATE Latin1_General_BIN NULL,
    report_json NVARCHAR(MAX) COLLATE Latin1_General_BIN NULL,
    rolled_back_by_request_id NVARCHAR(128) COLLATE Latin1_General_BIN NULL,
    rolled_back_at DATETIME2(7) NULL,
    created_at DATETIME2(7) NOT NULL CONSTRAINT DF_mig_operation_journal__created_at DEFAULT (SYSUTCDATETIME()),
    updated_at DATETIME2(7) NOT NULL CONSTRAINT DF_mig_operation_journal__updated_at DEFAULT (SYSUTCDATETIME()),
    finished_at DATETIME2(7) NULL,
    CONSTRAINT CK_mig_operation_journal__status CHECK (status IN (N'running', N'success', N'failed', N'no_change')),
    CONSTRAINT CK_mig_operation_journal__direction CHECK (result_direction IN (N'none', N'up', N'down', N'ack')),
    CONSTRAINT CK_mig_operation_journal__applied_files_json CHECK (applied_files IS NULL OR ISJSON(applied_files) = 1),
    CONSTRAINT CK_mig_operation_journal__report_json CHECK (report_json IS NULL OR ISJSON(report_json) = 1),
    CONSTRAINT CK_mig_operation_journal__manifest_sha256 CHECK (
      manifest_sha256 IS NULL OR manifest_sha256 NOT LIKE '%[^0-9a-f]%' COLLATE Latin1_General_BIN
    )
  );

  CREATE UNIQUE INDEX UX_mig_operation_journal__request_id ON dbo.mig_operation_journal(request_id);
  CREATE INDEX IX_mig_operation_journal__created_at ON dbo.mig_operation_journal(created_at DESC);
  CREATE INDEX IX_mig_operation_journal__rollback_lookup
    ON dbo.mig_operation_journal(status, result_direction, rolled_back_by_request_id, created_at DESC);
END

IF (
  SELECT COUNT_BIG(*) FROM sys.columns
  WHERE object_id = OBJECT_ID(N'dbo.mig_operation_journal', N'U')
) <> 24
OR NOT EXISTS (
  SELECT 1 FROM sys.columns
  WHERE object_id = OBJECT_ID(N'dbo.mig_operation_journal', N'U')
    AND name COLLATE Latin1_General_100_BIN2 = N'manifest_sha256' COLLATE Latin1_General_100_BIN2
    AND TYPE_NAME(user_type_id) = N'char'
    AND max_length = 64
)
OR NOT EXISTS (
  SELECT 1 FROM sys.columns
  WHERE object_id = OBJECT_ID(N'dbo.mig_operation_journal', N'U')
    AND name COLLATE Latin1_General_100_BIN2 = N'report_json' COLLATE Latin1_General_100_BIN2
    AND TYPE_NAME(user_type_id) = N'nvarchar'
    AND max_length = -1
)
OR NOT EXISTS (
  SELECT 1 FROM sys.indexes
  WHERE object_id = OBJECT_ID(N'dbo.mig_operation_journal', N'U')
    AND name COLLATE Latin1_General_100_BIN2 = N'UX_mig_operation_journal__request_id' COLLATE Latin1_General_100_BIN2
    AND is_unique = 1
)
BEGIN
  ;THROW 51000, N'Core precheck blocked: dbo.mig_operation_journal exists but does not match the mandatory journal contract', 1;
END

EXEC sys.sp_executesql N'CREATE OR ALTER PROCEDURE dbo.usp_mig_log_event
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
  SET @operation_id = COALESCE(@operation_id, TRY_CONVERT(UNIQUEIDENTIFIER, SESSION_CONTEXT(N''mig_operation_id'')));
  SET @direction = COALESCE(NULLIF(@direction, N''''), CASE WHEN @detail IN (N''up'', N''down'') THEN @detail ELSE N''unknown'' END);
  SET @direction = COALESCE(NULLIF(@direction, N''unknown''), TRY_CONVERT(NVARCHAR(8), SESSION_CONTEXT(N''mig_direction'')), N''unknown'');
  SET @content_sha256 = COALESCE(@content_sha256, TRY_CONVERT(CHAR(64), SESSION_CONTEXT(N''mig_content_sha256'')));
  SET @started_at = COALESCE(@started_at, TRY_CONVERT(DATETIME2(7), SESSION_CONTEXT(N''mig_started_at'')));
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
END';

EXEC sys.sp_executesql N'CREATE OR ALTER PROCEDURE dbo.usp_mig_try_get_lock
  @resource NVARCHAR(255),
  @timeout_ms INT = 60000
AS
BEGIN
  SET NOCOUNT ON;
  DECLARE @res INT;
  EXEC @res = sp_getapplock
    @Resource = @resource,
    @LockMode = ''Exclusive'',
    @LockOwner = ''Session'',
    @LockTimeout = @timeout_ms;
  IF (@res < 0)
    THROW 51000, ''Failed to acquire migration lock (sp_getapplock).'', 1;
END';
