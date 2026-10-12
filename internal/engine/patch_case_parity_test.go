package engine

import (
	"os"
	"path/filepath"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/output"
)

func TestCollectPatchSourcesUsesActualCaseVariantPaths(t *testing.T) {
	repo := t.TempDir()
	patchDir := filepath.Join(repo, "patch", "Fix Status")
	if err := os.MkdirAll(patchDir, 0o755); err != nil {
		t.Fatal(err)
	}
	for name, content := range map[string]string{
		"readme.md": "# Fix Status\n",
		"UP.SQL":    "SELECT 1;\n",
		"Down.Sql":  "SELECT 0;\n",
	} {
		if err := os.WriteFile(filepath.Join(patchDir, name), []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	engine := New(config.Default(repo))
	report := output.NewReport("pattern", "lf")
	sources, err := engine.collectPatchSources(&report)
	if err != nil {
		t.Fatal(err)
	}
	if report.Fail.Count != 0 || len(sources) != 1 {
		t.Fatalf("mixed-case patch collection failed: fail=%+v sources=%+v", report.Fail.Items, sources)
	}
	source := sources[0]
	if filepath.Base(source.ReadmePath) != "readme.md" || filepath.Base(source.UpPath) != "UP.SQL" || filepath.Base(source.DownPath) != "Down.Sql" {
		t.Fatalf("collector did not retain actual entry paths: %+v", source)
	}
	if source.UpSQL != "SELECT 1;\n" || source.DownSQL != "SELECT 0;\n" {
		t.Fatalf("collector read unexpected patch content: %+v", source)
	}

	// The cache contract must retain actual paths too; otherwise a second run
	// could regress only after the cache is populated.
	secondReport := output.NewReport("pattern", "lf")
	second, err := engine.collectPatchSources(&secondReport)
	if err != nil {
		t.Fatal(err)
	}
	if secondReport.Fail.Count != 0 || len(second) != 1 || second[0].UpPath != source.UpPath || second[0].DownPath != source.DownPath {
		t.Fatalf("cached mixed-case patch collection drifted: fail=%+v sources=%+v", secondReport.Fail.Items, second)
	}
}
