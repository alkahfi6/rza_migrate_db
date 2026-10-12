CREATE OR ALTER VIEW dbo.vJoinCandidates AS
SELECT r.Id, r.ParentId, r.BusinessCode
FROM dbo.JoinReference r
JOIN dbo.Parent p ON r.ParentId = p.Id
  AND r.BusinessCode = p.Code
  AND r.NotJoinedId > p.Id;
