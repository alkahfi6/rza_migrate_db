CREATE OR ALTER PROCEDURE dbo.uspChild AS
SELECT c.Id
FROM dbo.Child c
JOIN dbo.Parent p
  ON c.ParentId >= 0
 AND c.ParentId = p.Id;
