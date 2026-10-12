CREATE TABLE dbo.ReksaWindowPeriod_TM (
  [ProdId] INT NULL,
  [IsRange] BIT NULL,
  [RangePeriod] TINYINT NULL,
  [IsEffDate] BIT NULL,
  [MonthPeriod] TINYINT NULL,
  [DayStart] TINYINT NULL,
  [NDayBefore] TINYINT NULL,
  [IsBeforeHoliday] TINYINT NULL,
  [ReferDate] DATETIME NULL,
  [FixNAVDate] BIT NULL,
  [ReportNDBefore] TINYINT NULL,
  [PersenAUM] MONEY NULL
);
