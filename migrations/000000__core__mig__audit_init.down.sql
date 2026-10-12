IF OBJECT_ID(N'dbo.mig_event', N'U') IS NOT NULL
AND EXISTS (SELECT 1 FROM dbo.mig_event)
  THROW 51000, N'Core audit cleanup blocked: dbo.mig_event contains audit history', 1;

IF OBJECT_ID(N'dbo.mig_operation_journal', N'U') IS NOT NULL
AND EXISTS (SELECT 1 FROM dbo.mig_operation_journal)
  THROW 51000, N'Core audit cleanup blocked: dbo.mig_operation_journal contains deployment history', 1;

DROP TABLE IF EXISTS dbo.mig_event;
DROP TABLE IF EXISTS dbo.mig_operation_journal;
