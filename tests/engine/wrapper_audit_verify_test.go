package engine_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/engine"
)

func TestVerifyFailsWhenGeneratedWrapperAuditContractBroken(t *testing.T) {
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

	tableUp := ""
	for _, up := range genReport.NewMigrations.Files {
		if strings.Contains(up, "_tables_") {
			tableUp = up
			break
		}
	}
	if tableUp == "" {
		t.Fatalf("generated table up migration not found")
	}

	path := filepath.Join(workRoot, "migrations", tableUp)
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read generated migration failed: %v", err)
	}

	broken := strings.ReplaceAll(string(raw), "N'fail'", "N'broken'")
	if err := os.WriteFile(path, []byte(broken), 0o644); err != nil {
		t.Fatalf("write broken migration failed: %v", err)
	}

	verifyReport, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if verifyReport.Fail.Count == 0 {
		t.Fatalf("verify should fail when wrapper audit contract is broken")
	}

	hasRule := false
	for _, item := range verifyReport.Fail.Items {
		if item.Rule == "WRAPPER_AUDIT_CONTRACT_MISSING" {
			hasRule = true
			break
		}
	}
	if !hasRule {
		t.Fatalf("expected WRAPPER_AUDIT_CONTRACT_MISSING rule, got %+v", verifyReport.Fail.Items)
	}
}

func TestVerifyFailsWhenWrapperSeverityHeaderMissing(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "13_audit_wrapper_contract")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	path := filepath.Join(workRoot, "migrations", "20260217073015000_tables_dbo.TableA_v1.up.sql")
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read migration failed: %v", err)
	}

	broken := strings.ReplaceAll(string(raw), "Severity mapping:", "Mapping severity:")
	if err := os.WriteFile(path, []byte(broken), 0o644); err != nil {
		t.Fatalf("write broken migration failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if report.Fail.Count == 0 {
		t.Fatalf("verify should fail when wrapper severity header is missing")
	}

	hasRule := false
	for _, item := range report.Fail.Items {
		if item.Rule == "WRAPPER_SEVERITY_HEADER_MISSING" {
			hasRule = true
			break
		}
	}
	if !hasRule {
		t.Fatalf("expected WRAPPER_SEVERITY_HEADER_MISSING rule, got %+v", report.Fail.Items)
	}
}

func TestVerifyFailsWhenGeneratedWrapperActionOutsideWhitelist(t *testing.T) {
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

	tableUp := ""
	for _, up := range genReport.NewMigrations.Files {
		if strings.Contains(up, "_tables_") {
			tableUp = up
			break
		}
	}
	if tableUp == "" {
		t.Fatalf("generated table up migration not found")
	}

	path := filepath.Join(workRoot, "migrations", tableUp)
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read generated migration failed: %v", err)
	}

	broken := strings.Replace(string(raw), "DECLARE @action NVARCHAR(64) = N'create_table';", "DECLARE @action NVARCHAR(64) = N'unknown_action';", 1)
	if err := os.WriteFile(path, []byte(broken), 0o644); err != nil {
		t.Fatalf("write broken migration failed: %v", err)
	}

	verifyReport, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if verifyReport.Fail.Count == 0 {
		t.Fatalf("verify should fail when wrapper action is outside whitelist")
	}

	hasRule := false
	for _, item := range verifyReport.Fail.Items {
		if item.Rule == "WRAPPER_AUDIT_ACTION_INVALID" {
			hasRule = true
			break
		}
	}
	if !hasRule {
		t.Fatalf("expected WRAPPER_AUDIT_ACTION_INVALID rule, got %+v", verifyReport.Fail.Items)
	}
}

func TestVerifyFailsWhenGeneratedWrapperReasonOutsideWhitelist(t *testing.T) {
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

	tableUp := ""
	for _, up := range genReport.NewMigrations.Files {
		if strings.Contains(up, "_tables_") {
			tableUp = up
			break
		}
	}
	if tableUp == "" {
		t.Fatalf("generated table up migration not found")
	}

	path := filepath.Join(workRoot, "migrations", tableUp)
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read generated migration failed: %v", err)
	}

	broken := strings.Replace(string(raw), "@reason = N'OBJECT_EXISTS',", "@reason = N'UNKNOWN_REASON',", 1)
	if err := os.WriteFile(path, []byte(broken), 0o644); err != nil {
		t.Fatalf("write broken migration failed: %v", err)
	}

	verifyReport, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if verifyReport.Fail.Count == 0 {
		t.Fatalf("verify should fail when wrapper reason is outside whitelist")
	}

	hasRule := false
	for _, item := range verifyReport.Fail.Items {
		if item.Rule == "WRAPPER_AUDIT_REASON_INVALID" {
			hasRule = true
			break
		}
	}
	if !hasRule {
		t.Fatalf("expected WRAPPER_AUDIT_REASON_INVALID rule, got %+v", verifyReport.Fail.Items)
	}
}

func TestGenFailsWhenLegacyWrapperAuditBrokenEvenWithoutDiff(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "13_audit_wrapper_contract")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	legacy := `/*
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
	path := filepath.Join(workRoot, "migrations", "20260217073015000_tables_dbo.TableA_v1.up.sql")
	if err := os.WriteFile(path, []byte(strings.ReplaceAll(legacy, "\r\n", "\n")), 0o644); err != nil {
		t.Fatalf("write legacy migration failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("gen run failed: %v", err)
	}
	if report.Fail.Count == 0 {
		t.Fatalf("gen should fail when wrapper audit contract is broken")
	}
	if len(report.NewMigrations.Files) != 0 {
		t.Fatalf("gen must not emit migration files when precheck fails")
	}

	hasRule := false
	for _, item := range report.Fail.Items {
		if item.Rule == "WRAPPER_AUDIT_LEGACY_WEAK" {
			hasRule = true
			break
		}
	}
	if !hasRule {
		t.Fatalf("expected WRAPPER_AUDIT_LEGACY_WEAK fail in gen precheck, got %+v", report.Fail.Items)
	}
}

func TestVerifyFailsWhenPatchMigrationHasNoWrapperContract(t *testing.T) {
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

	patchUp := ""
	patchDown := ""
	for _, up := range genReport.NewMigrations.Files {
		if strings.Contains(up, "_patch_") {
			patchUp = up
			patchDown = strings.TrimSuffix(up, ".up.sql") + ".down.sql"
			break
		}
	}
	if patchUp == "" {
		t.Fatalf("generated patch migration not found")
	}

	if err := os.WriteFile(filepath.Join(workRoot, "migrations", patchUp), []byte("SELECT 1;\n"), 0o644); err != nil {
		t.Fatalf("write broken patch up failed: %v", err)
	}
	if err := os.WriteFile(filepath.Join(workRoot, "migrations", patchDown), []byte("SELECT 1;\n"), 0o644); err != nil {
		t.Fatalf("write broken patch down failed: %v", err)
	}

	verifyReport, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if verifyReport.Fail.Count == 0 {
		t.Fatalf("verify should fail when patch wrapper contract is missing")
	}

	hasRule := false
	for _, item := range verifyReport.Fail.Items {
		if item.Rule == "WRAPPER_CONTRACT_MISSING" {
			hasRule = true
			break
		}
	}
	if !hasRule {
		t.Fatalf("expected WRAPPER_CONTRACT_MISSING for patch migration, got %+v", verifyReport.Fail.Items)
	}
}

func TestVerifyFailsWhenWrapperMetadataIsNonDeterministic(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "13_audit_wrapper_contract")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	path := filepath.Join(workRoot, "migrations", "20260217073015000_tables_dbo.TableA_v1.up.sql")
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read migration failed: %v", err)
	}

	broken := strings.Replace(string(raw), "DECLARE @version BIGINT = 20260217073015000;", "DECLARE @version BIGINT = ABS(CHECKSUM(NEWID()));", 1)
	if broken == string(raw) {
		t.Fatalf("failed to mutate @version declaration for metadata deterministic test")
	}
	if err := os.WriteFile(path, []byte(broken), 0o644); err != nil {
		t.Fatalf("write broken migration failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if report.Fail.Count == 0 {
		t.Fatalf("verify should fail when wrapper metadata is non-deterministic")
	}

	hasRule := false
	for _, item := range report.Fail.Items {
		if item.Rule == "WRAPPER_METADATA_NON_DETERMINISTIC" {
			hasRule = true
			break
		}
	}
	if !hasRule {
		t.Fatalf("expected WRAPPER_METADATA_NON_DETERMINISTIC rule, got %+v", report.Fail.Items)
	}
}

func TestVerifyFailsWhenThrowOnlyInComment(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "13_audit_wrapper_contract")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	path := filepath.Join(workRoot, "migrations", "20260217073015000_tables_dbo.TableA_v1.up.sql")
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read migration failed: %v", err)
	}

	broken := strings.Replace(string(raw), "THROW;", "-- THROW;", 1)
	if broken == string(raw) {
		t.Fatalf("failed to mutate THROW statement for semantic throw test")
	}
	if err := os.WriteFile(path, []byte(broken), 0o644); err != nil {
		t.Fatalf("write broken migration failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if report.Fail.Count == 0 {
		t.Fatalf("verify should fail when THROW exists only in comment")
	}

	hasRule := false
	for _, item := range report.Fail.Items {
		if item.Rule == "WRAPPER_THROW_REQUIRED" {
			hasRule = true
			break
		}
	}
	if !hasRule {
		t.Fatalf("expected WRAPPER_THROW_REQUIRED rule, got %+v", report.Fail.Items)
	}
}
