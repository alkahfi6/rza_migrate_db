CREATE TABLE dbo.SecurityTransactionLink_TR (
  [DealIdParent] BIGINT NULL,
  [DealIdChild] BIGINT NULL,
  [Nominal] MONEY NULL,
  [TransactionFeeAmount] MONEY NULL,
  [Proceed] MONEY NULL,
  [AccruedDays] INT NULL,
  [AccruedInterest] MONEY NULL,
  [TaxOnAccrued] MONEY NULL,
  [TaxOnCapitalGainLoss] MONEY NULL,
  [SafeKeepingFeeAfterTax] MONEY NULL,
  [TotalProceed] MONEY NULL,
  [SafeKeepingFeeAmount] MONEY NULL
);
