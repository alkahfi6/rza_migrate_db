CREATE TABLE dbo.ReksaHolidayTable_TM (
  [HolidayId] INT IDENTITY(1,1) NOT NULL,
  [HolidayDesc] VARCHAR(40) COLLATE Latin1_General_BIN NULL,
  [ValueDate] DATETIME NULL,
  [LastUpdate] DATETIME NULL,
  [LastUser] INT NULL,
  CONSTRAINT [PK_ReksaHolidayTable_TM] PRIMARY KEY CLUSTERED ([HolidayId] ASC)
);
