package verify

import (
	"strings"
	"testing"

	"rza_migrate_db/internal/output"
	"rza_migrate_db/internal/sqlparse"
)

func TestHardenedGeneratedPrecheckRejectsUnguardedDDL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		category string
		sql      string
	}{
		{name: "add column", category: "tables", sql: "ALTER TABLE dbo.T ADD [code] INT NULL;"},
		{name: "drop column with data loss", category: "tables", sql: "IF COL_LENGTH(N'dbo.T', N'code') IS NOT NULL BEGIN ALTER TABLE dbo.T DROP COLUMN [code]; END"},
		{name: "drop table with data loss", category: "tables", sql: "IF OBJECT_ID(N'dbo.T', N'U') IS NOT NULL DROP TABLE dbo.T;"},
		{name: "create index missing table guard", category: "indexes", sql: "IF NOT EXISTS (SELECT 1 FROM sys.indexes WHERE object_id=OBJECT_ID(N'dbo.T') AND name=N'IX_T') CREATE INDEX IX_T ON dbo.T(code);"},
		{name: "unique index missing duplicate guard", category: "indexes", sql: "IF OBJECT_ID(N'dbo.T',N'U') IS NULL THROW 51000, N'Precheck blocked: CREATE INDEX target table does not exist', 1; CREATE UNIQUE INDEX UX_T ON dbo.T(code);"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			report := output.NewReport("pattern", "LF")
			upper := strings.ToUpper(sqlparse.StripSQLStringsAndComments(tt.sql))
			validateHardenedGeneratedPrechecks(&report, "migrations/20260726000000001_"+tt.category+"_dbo.T_v1.up.sql", tt.category, tt.sql, upper, true)
			if report.Fail.Count == 0 {
				t.Fatalf("expected hardened precheck failure for SQL:\n%s", tt.sql)
			}
		})
	}
}

func TestHardenedGeneratedPrecheckAcceptsGuardedUniqueIndex(t *testing.T) {
	t.Parallel()

	sql := `
IF OBJECT_ID(N'dbo.T', N'U') IS NULL
BEGIN
  THROW 51000, N'Precheck blocked: CREATE INDEX target table does not exist: dbo.T', 1;
END
IF NOT EXISTS (SELECT 1 FROM sys.indexes WHERE object_id=OBJECT_ID(N'dbo.T') AND name=N'UX_T')
BEGIN
  IF EXISTS (SELECT 1 FROM dbo.T GROUP BY [code] HAVING COUNT_BIG(*) > 1)
    THROW 51000, N'duplicate', 1;
  CREATE UNIQUE INDEX UX_T ON dbo.T(code);
END`
	report := output.NewReport("pattern", "LF")
	upper := strings.ToUpper(sqlparse.StripSQLStringsAndComments(sql))
	validateHardenedGeneratedPrechecks(&report, "migrations/20260726000000002_indexes_dbo.T_v1.up.sql", "indexes", sql, upper, true)
	if report.Fail.Count != 0 {
		t.Fatalf("unexpected hardened precheck failure: %#v", report.Fail.Items)
	}
}
