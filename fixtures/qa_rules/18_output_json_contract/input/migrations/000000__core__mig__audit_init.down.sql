IF OBJECT_ID(N'dbo.usp_mig_try_get_lock', N'P') IS NOT NULL
  DROP PROCEDURE dbo.usp_mig_try_get_lock;

IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  DROP PROCEDURE dbo.usp_mig_log_event;

IF OBJECT_ID(N'dbo.mig_event', N'U') IS NOT NULL
  DROP TABLE dbo.mig_event;