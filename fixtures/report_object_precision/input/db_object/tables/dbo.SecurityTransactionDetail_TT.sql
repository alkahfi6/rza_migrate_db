CREATE TABLE dbo.SecurityTransactionDetail_TT (
  [RowId] BIGINT IDENTITY(1,1) NOT NULL,
  [DealId] BIGINT NOT NULL,
  [CIFIdDebet] BIGINT NOT NULL,
  [CIFIdKredit] BIGINT NOT NULL,
  [TrxDetailType] SMALLINT NOT NULL,
  [D/K] VARCHAR(1) COLLATE Latin1_General_BIN NULL,
  [CustodyName] VARCHAR(50) COLLATE Latin1_General_BIN NULL,
  [InsertedBy] INT NULL,
  [InsertedDate] DATETIME NULL,
  [LastUpdateBy] INT NULL,
  [LastUpdateDate] DATETIME NULL,
  [CustNameDebet] VARCHAR(95) COLLATE Latin1_General_BIN NULL,
  [CustNameKredit] VARCHAR(95) COLLATE Latin1_General_BIN NULL,
  [AcquitionDate] DATETIME NULL
);
