package engine_test

import (
	"path/filepath"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/engine"
)

func TestEngineFailsWhenConfigHardLockInvalid(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "06_ordering_determinism")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	cfg := config.Default(workRoot)
	delete(cfg.AllowedDecisions, "run")

	eng := engine.New(cfg)
	report, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if report.Fail.Count == 0 {
		t.Fatalf("expected fail for invalid config hard lock")
	}

	found := false
	for _, item := range report.Fail.Items {
		if item.Rule == "CONFIG_HARD_LOCK_INVALID" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected CONFIG_HARD_LOCK_INVALID fail rule, got %+v", report.Fail.Items)
	}
}
