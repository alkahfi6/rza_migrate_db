package verify

import (
	"testing"

	"rza_migrate_db/internal/output"
)

func TestValidateProgrammabilitySingleObjectContract_AllowsTempTableDDL(t *testing.T) {
	report := output.NewReport("p", "lf")
	sql := `CREATE OR ALTER PROCEDURE dbo.ProcA
AS
BEGIN
  CREATE TABLE #TempA (id INT NOT NULL);
  CREATE UNIQUE NONCLUSTERED INDEX IX_TempA__id ON #TempA(id);
  ALTER TABLE #TempA ADD status VARCHAR(10) NULL;
  DROP INDEX IF EXISTS IX_TempA__id ON #TempA;
  DROP TABLE IF EXISTS #TempA;
END;
`

	ValidateProgrammabilitySingleObjectContract(&report, "db_object/procedures/dbo.ProcA.sql", sql, "procedures")

	if hasFailRule(report, "PROGRAMMABILITY_SOURCE_CONTRACT_VIOLATION") {
		t.Fatalf("did not expect PROGRAMMABILITY_SOURCE_CONTRACT_VIOLATION, got %+v", report.Fail.Items)
	}
	if hasFailRule(report, "PROGRAMMABILITY_MULTI_OBJECT") {
		t.Fatalf("did not expect PROGRAMMABILITY_MULTI_OBJECT, got %+v", report.Fail.Items)
	}
}

func TestValidateProgrammabilitySingleObjectContract_RejectsPersistentDDL(t *testing.T) {
	report := output.NewReport("p", "lf")
	sql := `CREATE OR ALTER PROCEDURE dbo.ProcA
AS
BEGIN
  CREATE TABLE dbo.BadLeak (id INT NOT NULL);
END;
`

	ValidateProgrammabilitySingleObjectContract(&report, "db_object/procedures/dbo.ProcA.sql", sql, "procedures")
	if !hasFailRule(report, "PROGRAMMABILITY_SOURCE_CONTRACT_VIOLATION") {
		t.Fatalf("expected PROGRAMMABILITY_SOURCE_CONTRACT_VIOLATION, got %+v", report.Fail.Items)
	}
}

func TestValidateProgrammabilitySingleObjectContract_RejectsPersistentCreateIndexVariants(t *testing.T) {
	report := output.NewReport("p", "lf")
	sql := `CREATE OR ALTER PROCEDURE dbo.ProcA
AS
BEGIN
  CREATE NONCLUSTERED INDEX IX_TableA__id ON dbo.TableA(id);
END;
`

	ValidateProgrammabilitySingleObjectContract(&report, "db_object/procedures/dbo.ProcA.sql", sql, "procedures")
	if !hasFailRule(report, "PROGRAMMABILITY_SOURCE_CONTRACT_VIOLATION") {
		t.Fatalf("expected PROGRAMMABILITY_SOURCE_CONTRACT_VIOLATION, got %+v", report.Fail.Items)
	}
}

func TestValidateProgrammabilitySingleObjectContract_AllowsTempDropIndexOldSyntax(t *testing.T) {
	report := output.NewReport("p", "lf")
	sql := `CREATE OR ALTER PROCEDURE dbo.ProcA
AS
BEGIN
  CREATE TABLE #TempA (id INT NOT NULL);
  CREATE INDEX IX_TempA__id ON #TempA(id);
  DROP INDEX #TempA.IX_TempA__id;
  DROP TABLE #TempA;
END;
`

	ValidateProgrammabilitySingleObjectContract(&report, "db_object/procedures/dbo.ProcA.sql", sql, "procedures")
	if hasFailRule(report, "PROGRAMMABILITY_SOURCE_CONTRACT_VIOLATION") {
		t.Fatalf("did not expect PROGRAMMABILITY_SOURCE_CONTRACT_VIOLATION, got %+v", report.Fail.Items)
	}
}

func TestValidateProgrammabilitySingleObjectContract_RejectsPersistentDropIndexOldSyntax(t *testing.T) {
	report := output.NewReport("p", "lf")
	sql := `CREATE OR ALTER PROCEDURE dbo.ProcA
AS
BEGIN
  DROP INDEX dbo.TableA.IX_TableA__id;
END;
`

	ValidateProgrammabilitySingleObjectContract(&report, "db_object/procedures/dbo.ProcA.sql", sql, "procedures")
	if !hasFailRule(report, "PROGRAMMABILITY_SOURCE_CONTRACT_VIOLATION") {
		t.Fatalf("expected PROGRAMMABILITY_SOURCE_CONTRACT_VIOLATION, got %+v", report.Fail.Items)
	}
}
