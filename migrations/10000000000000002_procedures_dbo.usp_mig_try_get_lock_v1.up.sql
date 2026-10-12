CREATE OR ALTER PROCEDURE dbo.usp_mig_try_get_lock
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