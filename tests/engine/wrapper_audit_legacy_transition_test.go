package engine_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/engine"
)

func TestVerifyLegacyWrapperAuditAlwaysFailsStrictOnly(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "13_audit_wrapper_contract")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	legacyFile := filepath.Join(workRoot, "migrations", "20260217073015000_tables_dbo.TableA_v1.up.sql")
	if err := os.WriteFile(legacyFile, []byte(strings.ReplaceAll(legacyWrapperSQL(), "\r\n", "\n")), 0o644); err != nil {
		t.Fatalf("write legacy wrapper fixture failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if report.Fail.Count == 0 {
		t.Fatalf("strict-only mode should fail when legacy wrapper audit is incomplete")
	}

	hasRule := false
	for _, item := range report.Fail.Items {
		if item.Rule == "WRAPPER_AUDIT_LEGACY_WEAK" {
			hasRule = true
			break
		}
	}
	if !hasRule {
		t.Fatalf("expected WRAPPER_AUDIT_LEGACY_WEAK fail item in strict-only mode")
	}
}

func TestVerifyGeneratedWrapperHasNoLegacyAuditWarning(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "06_ordering_determinism")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	genReport, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("gen run failed: %v", err)
	}
	if genReport.Fail.Count != 0 {
		t.Fatalf("gen expected no fail items, got %d", genReport.Fail.Count)
	}

	verifyReport, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if verifyReport.Fail.Count != 0 {
		t.Fatalf("verify expected no fail items, got %d", verifyReport.Fail.Count)
	}

	for _, item := range verifyReport.Pass.Items {
		if strings.EqualFold(item.Rule, "WRAPPER_AUDIT_LEGACY_WEAK") {
			t.Fatalf("generated wrappers should not produce WRAPPER_AUDIT_LEGACY_WEAK warning")
		}
	}
}

func legacyWrapperSQL() string {
	return `/*
Severity mapping:
0=INFO (ok/skip normal)
1=NOTICE (schema change)
2=WARNING (follow-up/manual risk)
3=ERROR (stop)
*/
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = 20260217073015000;
DECLARE @category NVARCHAR(16) = N'tables';
DECLARE @file NVARCHAR(260) = N'20260217073015000_tables_dbo.TableA_v1.up.sql';

BEGIN TRY
  BEGIN TRAN;
  CREATE TABLE dbo.TableA (
    id INT NOT NULL,
    status VARCHAR(10) NULL,
    amount DECIMAL(10,2) NULL,
    CONSTRAINT PK_TableA PRIMARY KEY (id),
    CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status
  );
  COMMIT;
END TRY
BEGIN CATCH
  IF @@TRANCOUNT > 0 ROLLBACK;
  THROW;
END CATCH;
`
}
