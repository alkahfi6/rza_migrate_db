package engine_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/engine"
	"rza_migrate_db/internal/output"
)

func TestVerifyFailsWhenCoreAuditLogProcMissing(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "06_ordering_determinism")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	upPath := filepath.Join(workRoot, "migrations", "000000__core__mig__audit_init.up.sql")
	raw, err := os.ReadFile(upPath)
	if err != nil {
		t.Fatalf("read core up failed: %v", err)
	}
	broken := strings.ReplaceAll(string(raw), "dbo.usp_mig_log_event", "dbo.usp_mig_log_event_removed")
	if err := os.WriteFile(upPath, []byte(broken), 0o644); err != nil {
		t.Fatalf("write broken core up failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if report.Fail.Count == 0 {
		t.Fatalf("expected fail when core audit log proc is missing")
	}
	if !hasFailRule(report, "CORE_AUDIT_CONTRACT_MISSING") {
		t.Fatalf("expected CORE_AUDIT_CONTRACT_MISSING fail rule, got %+v", report.Fail.Items)
	}
}

func TestVerifyFailsWhenCoreLockProcNotPresent(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "06_ordering_determinism")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	upPath := filepath.Join(workRoot, "migrations", "000000__core__mig__audit_init.up.sql")
	raw, err := os.ReadFile(upPath)
	if err != nil {
		t.Fatalf("read core up failed: %v", err)
	}
	broken := strings.ReplaceAll(string(raw), "sp_getapplock", "sp_getapplock_removed")
	if err := os.WriteFile(upPath, []byte(broken), 0o644); err != nil {
		t.Fatalf("write broken core up failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if report.Fail.Count == 0 {
		t.Fatalf("verify should fail when core lock proc contract is missing")
	}
	if !hasFailRule(report, "CORE_AUDIT_CONTRACT_MISSING") {
		t.Fatalf("expected CORE_AUDIT_CONTRACT_MISSING fail rule")
	}
}

func hasFailRule(report output.Report, rule string) bool {
	for _, item := range report.Fail.Items {
		if item.Rule == rule {
			return true
		}
	}
	return false
}
