CREATE TABLE dbo.TRSParamIssuer_TH (
  [ApprId] BIGINT NULL,
  [IssuerId] BIGINT NULL,
  [IssuerName] VARCHAR(250) COLLATE Latin1_General_BIN NULL,
  [Limit] MONEY NULL,
  [ProcessCode] CHAR(1) COLLATE Latin1_General_BIN NULL,
  [LastModifyDate] DATETIME NULL,
  [LastModifyNIK] INT NULL,
  [Notification] DECIMAL(12,2) NULL
);
