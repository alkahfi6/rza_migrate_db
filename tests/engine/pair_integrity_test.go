package engine_test

import (
	"os"
	"path/filepath"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/engine"
)

func TestVerifyFailsWhenPairMetadataMismatch(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "01_hard_fail_rule")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	oldDown := filepath.Join(workRoot, "migrations", "20260217073015000_tables_dbo.TableA_v1.down.sql")
	newDown := filepath.Join(workRoot, "migrations", "20260217073015000_indexes_dbo.TableA_v1.down.sql")
	if err := os.Rename(oldDown, newDown); err != nil {
		t.Fatalf("rename migration down file failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if report.Fail.Count == 0 {
		t.Fatalf("expected fail for mismatched pair metadata")
	}

	found := false
	for _, item := range report.Fail.Items {
		if item.Rule == "MIGRATION_PAIR_MISMATCH" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected MIGRATION_PAIR_MISMATCH fail rule, got %+v", report.Fail.Items)
	}
}
