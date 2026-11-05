package main

import (
	"os"
	"path/filepath"
	"rza_migrate_db/internal/getobject"
	"strings"
	"testing"
)

func TestExeDefaultContinuesAndWritesReadableIssues(t *testing.T) {
	root := t.TempDir()
	t.Chdir(root)
	oldExport, oldConfig := exportFromDatabase, compiledGetObjectConfig
	t.Cleanup(func() { exportFromDatabase = oldExport; compiledGetObjectConfig = oldConfig })
	compiledGetObjectConfig = getobject.DefaultAppConfig(root)
	exportFromDatabase = func(cfg getobject.AppConfig, opts getobject.ExportOptions) (getobject.ExportSummary, error) {
		if !cfg.SkipUnsupportedTables {
			t.Fatal("exe without flags must continue")
		}
		return getobject.ExportSummary{ModuleWarnings: []string{"procedure dbo.Broken skipped: definition is empty or encrypted"}}, nil
	}
	if code := run([]string{"rza_get_object_db.exe"}); code != 0 {
		t.Fatalf("exit=%d", code)
	}
	report := filepath.Join(root, "out/get_object_issues.txt")
	raw, err := os.ReadFile(report)
	if err != nil || !strings.Contains(string(raw), "dbo.Broken") || !strings.Contains(string(raw), "encrypted") {
		t.Fatalf("missing readable diagnosis: %v", err)
	}
	raw, err = os.ReadFile(filepath.Join(root, "out/get_object_summary.json"))
	if err != nil || !strings.Contains(string(raw), "dbo.Broken") {
		t.Fatal("JSON evidence missing")
	}
	exportFromDatabase = func(getobject.AppConfig, getobject.ExportOptions) (getobject.ExportSummary, error) {
		return getobject.ExportSummary{}, nil
	}
	if code := run([]string{"rza_get_object_db.exe"}); code != 0 {
		t.Fatal(code)
	}
	raw, err = os.ReadFile(report)
	if err != nil || strings.Contains(string(raw), "dbo.Broken") {
		t.Fatal("stale issues survived clean run")
	}
}
