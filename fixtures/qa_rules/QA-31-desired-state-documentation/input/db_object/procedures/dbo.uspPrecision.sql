CREATE OR ALTER PROCEDURE dbo.uspPrecision AS
WITH ParentRows AS (SELECT Id FROM dbo.Parent)
SELECT Id INTO #work FROM ParentRows;
UPDATE TOP (1) c SET ParentId=1
FROM dbo.Child c WHERE c.Id=1;
MERGE dbo.Child AS t USING dbo.Parent AS s ON t.Id=s.Id
WHEN MATCHED THEN UPDATE SET ParentId=s.Id
WHEN NOT MATCHED THEN INSERT(Id, ParentId) VALUES(s.Id, s.Id);
SELECT c.Id FROM dbo.Child c, dbo.Parent p WHERE c.ParentId=p.Id;
EXEC
dbo.uspChild; EXEC dbo.uspChild;
