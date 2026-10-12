CREATE TABLE dbo.TRSProcessBuyBackLog_TH (
  [OrderId] BIGINT NULL,
  [DealId] BIGINT NULL,
  [CIFId] BIGINT NULL,
  [SecId] INT NULL,
  [NominalBuyBack] MONEY NULL,
  [DebitGLNoGiroBI] VARCHAR(12) COLLATE Latin1_General_BIN NULL,
  [DebitGLBrGiroBI] VARCHAR(5) COLLATE Latin1_General_BIN NULL,
  [CreditRekeningGNC] VARCHAR(20) COLLATE Latin1_General_BIN NULL,
  [ProcessUserNik] INT NULL,
  [ProcessDate] DATETIME NULL,
  [UniqueId] UNIQUEIDENTIFIER NULL,
  [Status] CHAR(1) COLLATE Latin1_General_BIN NULL
);
