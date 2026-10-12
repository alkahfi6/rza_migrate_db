CREATE TABLE dbo.JoinReference (
  Id INT NOT NULL,
  ParentId INT NULL,
  BusinessCode NVARCHAR(20) NULL,
  NotJoinedId INT NULL,
  CONSTRAINT PK_JoinReference PRIMARY KEY (Id),
  CONSTRAINT FK_JoinReference_NotJoined FOREIGN KEY (NotJoinedId) REFERENCES dbo.Parent (Id)
);
