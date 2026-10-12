package engine

import (
	"os"
	"path/filepath"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/output"
)

func TestBootstrapSeedsMissingOptionalRepoFilesWithoutOverwrite(t *testing.T) {
	root := t.TempDir()
	existing := []byte("# keep custom content\n")
	if err := os.WriteFile(filepath.Join(root, ".gitignore"), existing, 0o644); err != nil {
		t.Fatal(err)
	}

	engine := New(config.Default(root))
	report := output.NewReport("p", "lf")
	if err := engine.seedOptionalRepoFilesForBootstrap(&report); err != nil {
		t.Fatal(err)
	}
	if report.Pass.Count != 2 {
		t.Fatalf("expected two created files, got %+v", report.Pass.Items)
	}
	got, err := os.ReadFile(filepath.Join(root, ".gitignore"))
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(existing) {
		t.Fatalf("existing helper file was overwritten: %q", got)
	}
	for _, rel := range []string{"constraints_naming.md", ".gitattributes"} {
		info, err := os.Stat(filepath.Join(root, rel))
		if err != nil || !info.Mode().IsRegular() || info.Size() == 0 {
			t.Fatalf("expected seeded %s: info=%v err=%v", rel, info, err)
		}
	}

	second := output.NewReport("p", "lf")
	if err := engine.seedOptionalRepoFilesForBootstrap(&second); err != nil {
		t.Fatal(err)
	}
	if second.Pass.Count != 0 {
		t.Fatalf("second seed must be idempotent: %+v", second.Pass.Items)
	}
}
