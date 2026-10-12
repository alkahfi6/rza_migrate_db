package engine_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/engine"
)

func TestGenPatchWrapperContainsIntegrityHashes(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "15_patch_canonicalization")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("gen run failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("gen expected no fail items, got %d", report.Fail.Count)
	}

	patchUp := ""
	for _, up := range report.NewMigrations.Files {
		if strings.Contains(up, "_patch_") {
			patchUp = up
			break
		}
	}
	if patchUp == "" {
		t.Fatalf("generated patch up migration not found")
	}

	raw, err := os.ReadFile(filepath.Join(workRoot, "migrations", patchUp))
	if err != nil {
		t.Fatalf("read generated patch migration failed: %v", err)
	}
	upper := strings.ToUpper(string(raw))
	if !strings.Contains(upper, "PATCH_UP_SHA256:") {
		t.Fatalf("generated patch migration missing patch_up_sha256 header")
	}
	if !strings.Contains(upper, "PATCH_DOWN_SHA256:") {
		t.Fatalf("generated patch migration missing patch_down_sha256 header")
	}
}

func TestVerifyFailsWhenPatchSourceMutatedAfterGeneration(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "15_patch_canonicalization")
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

	patchUpSource := filepath.Join(workRoot, "patch", "Fix Status!!", "up.sql")
	raw, err := os.ReadFile(patchUpSource)
	if err != nil {
		t.Fatalf("read patch source failed: %v", err)
	}
	mutated := string(raw) + "\n-- mutated after generation\n"
	if err := os.WriteFile(patchUpSource, []byte(mutated), 0o644); err != nil {
		t.Fatalf("write patch source failed: %v", err)
	}

	verifyReport, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if verifyReport.Fail.Count == 0 {
		t.Fatalf("verify should fail for mutated patch source")
	}

	found := false
	for _, item := range verifyReport.Fail.Items {
		if item.Rule == "PATCH_SOURCE_MUTATED" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected PATCH_SOURCE_MUTATED fail rule, got %+v", verifyReport.Fail.Items)
	}
}

func TestVerifyFailsWhenPatchHeaderMetadataMismatch(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "15_patch_canonicalization")
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
	for _, up := range genReport.NewMigrations.Files {
		if strings.Contains(up, "_patch_") {
			patchUp = up
			break
		}
	}
	if patchUp == "" {
		t.Fatalf("generated patch up migration not found")
	}

	patchUpPath := filepath.Join(workRoot, "migrations", patchUp)
	raw, err := os.ReadFile(patchUpPath)
	if err != nil {
		t.Fatalf("read generated patch migration failed: %v", err)
	}
	broken := strings.Replace(string(raw), "patch_folder: Fix Status!!", "patch_folder: Wrong Folder", 1)
	if err := os.WriteFile(patchUpPath, []byte(broken), 0o644); err != nil {
		t.Fatalf("write broken patch migration failed: %v", err)
	}

	verifyReport, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if verifyReport.Fail.Count == 0 {
		t.Fatalf("verify should fail for patch header metadata mismatch")
	}

	found := false
	for _, item := range verifyReport.Fail.Items {
		if item.Rule == "PATCH_HEADER_METADATA_MISMATCH" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected PATCH_HEADER_METADATA_MISMATCH fail rule, got %+v", verifyReport.Fail.Items)
	}
}
