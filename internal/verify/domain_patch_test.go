package verify

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rza_migrate_db/internal/output"
)

func TestValidatePatchFolders_FailRootFileAndInvalidFolder(t *testing.T) {
	root := t.TempDir()
	patchRoot := filepath.Join(root, "patch")
	if err := os.MkdirAll(patchRoot, 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	if err := os.WriteFile(filepath.Join(patchRoot, "rogue.sql"), []byte("SELECT 1;"), 0o644); err != nil {
		t.Fatalf("write rogue file failed: %v", err)
	}
	invalidDir := filepath.Join(patchRoot, "Fix Status")
	if err := os.MkdirAll(invalidDir, 0o755); err != nil {
		t.Fatalf("mkdir invalid patch dir failed: %v", err)
	}
	if err := os.WriteFile(filepath.Join(invalidDir, "README.md"), []byte("# x\n"), 0o644); err != nil {
		t.Fatalf("write readme failed: %v", err)
	}
	if err := os.WriteFile(filepath.Join(invalidDir, "up.sql"), []byte("SELECT 1;\n"), 0o644); err != nil {
		t.Fatalf("write up failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	err := ValidatePatchFolders(&report, PatchFoldersOptions{
		RepoPath: func(rel string) string { return filepath.Join(root, rel) },
		PatchDir: "patch",
	})
	if err != nil {
		t.Fatalf("validate failed: %v", err)
	}
	if report.Fail.Count != 2 {
		t.Fatalf("expected 2 fail items, got %d", report.Fail.Count)
	}
}

func TestValidatePatchSourceIntegrityAcceptsCaseVariantsOnCaseSensitiveRunner(t *testing.T) {
	root := t.TempDir()
	patchDir := filepath.Join(root, "patch", "Fix Status")
	if err := os.MkdirAll(patchDir, 0o755); err != nil {
		t.Fatal(err)
	}
	upSQL := "INSERT INTO dbo.T(id) VALUES (1);\n"
	downSQL := "DELETE FROM dbo.T WHERE id = 1;\n"
	for name, content := range map[string]string{
		"readme.md": "# Fix Status\n",
		"UP.SQL":    upSQL,
		"Down.Sql":  downSQL,
	} {
		if err := os.WriteFile(filepath.Join(patchDir, name), []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	hash := func(input string) string {
		sum := sha256.Sum256([]byte(input))
		return hex.EncodeToString(sum[:])
	}
	migrationRel := "migrations/10000000000000003_patch_fix_status_v1.up.sql"
	migrationPath := filepath.Join(root, filepath.FromSlash(migrationRel))
	if err := os.MkdirAll(filepath.Dir(migrationPath), 0o755); err != nil {
		t.Fatal(err)
	}
	header := "/*\npatch_folder: Fix Status\npatch_name: fix_status\nreadme: patch/Fix Status/README.md\n" +
		"patch_up_sha256: " + hash(upSQL) + "\npatch_down_sha256: " + hash(downSQL) + "\n*/\nSELECT 1;\n"
	if err := os.WriteFile(migrationPath, []byte(header), 0o644); err != nil {
		t.Fatal(err)
	}

	report := output.NewReport("p", "lf")
	err := ValidatePatchSourceIntegrity(&report, PatchSourceIntegrityOptions{
		RepoPath: func(rel string) string { return filepath.Join(root, filepath.FromSlash(rel)) },
		PatchDir: "patch",
		Metas: []MigrationMeta{{
			Category: "patch", Direction: "up", Object: "fix_status", RelPath: migrationRel,
		}},
		CanonicalizePatchName: func(input string) string {
			return strings.ToLower(strings.ReplaceAll(strings.TrimSpace(input), " ", "_"))
		},
		HashStableContent: hash,
	})
	if err != nil {
		t.Fatal(err)
	}
	if report.Fail.Count != 0 || report.Pass.Count != 1 || report.Pass.Items[0].Rule != "PATCH_IMMUTABILITY_OK" {
		t.Fatalf("case-variant patch source must pass on every OS: fail=%+v pass=%+v", report.Fail.Items, report.Pass.Items)
	}
}

func TestValidatePatchFoldersRejectsExtraOrDuplicateContractEntry(t *testing.T) {
	root := t.TempDir()
	patchDir := filepath.Join(root, "patch", "Fix Status")
	if err := os.MkdirAll(patchDir, 0o755); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"readme.md", "UP.SQL", "down.sql", "notes.txt"} {
		if err := os.WriteFile(filepath.Join(patchDir, name), []byte("x\n"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	report := output.NewReport("p", "lf")
	if err := ValidatePatchFolders(&report, PatchFoldersOptions{
		RepoPath: func(rel string) string { return filepath.Join(root, filepath.FromSlash(rel)) },
		PatchDir: "patch",
	}); err != nil {
		t.Fatal(err)
	}
	if report.Fail.Count != 1 || report.Fail.Items[0].Rule != "PATCH_FOLDER_INVALID_CONTENT" {
		t.Fatalf("extra patch entry must fail strict three-file contract: %+v", report.Fail.Items)
	}
}

func TestValidatePatchSourceIntegrity_FailSourceMissing(t *testing.T) {
	root := t.TempDir()
	report := output.NewReport("p", "lf")

	err := ValidatePatchSourceIntegrity(&report, PatchSourceIntegrityOptions{
		RepoPath: func(rel string) string { return filepath.Join(root, rel) },
		PatchDir: "patch",
		Metas: []MigrationMeta{
			{
				Category:  "patch",
				Direction: "up",
				Object:    "fix_status",
				RelPath:   "migrations/20260217073015010_patch_fix_status_v1.up.sql",
			},
		},
		CanonicalizePatchName: func(input string) string { return input },
		HashStableContent:     func(input string) string { return input },
	})
	if err != nil {
		t.Fatalf("validate failed: %v", err)
	}
	if report.Fail.Count != 1 {
		t.Fatalf("expected 1 fail item, got %d", report.Fail.Count)
	}
	if report.Fail.Items[0].Rule != "PATCH_SOURCE_MISSING" {
		t.Fatalf("unexpected fail rule: %s", report.Fail.Items[0].Rule)
	}
}
