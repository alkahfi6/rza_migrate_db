package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCLIBasicRuns(t *testing.T) {
	// Test --build-info
	exitCode := run([]string{"rza_migrate_report", "--build-info"})
	if exitCode != 0 {
		t.Fatalf("expected --build-info exit code 0, got %d", exitCode)
	}

	// Test invalid executable name
	exitCode = run([]string{"invalid_tool_name.exe"})
	if exitCode != 1 {
		t.Fatalf("expected invalid executable name exit code 1, got %d", exitCode)
	}

	// Test with explicit -repo pointing to a temp repo
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}
	outDir := filepath.Join(tmp, "out", "report")

	exitCode = run([]string{
		"rza_migrate_report.exe",
		"-repo", tmp,
		"-out", outDir,
		"-pause-on-exit", "false",
	})
	if exitCode != 0 {
		t.Fatalf("expected run with -repo exit code 0, got %d", exitCode)
	}

	// Verify report_summary.csv was generated in outDir
	if _, err := os.Stat(filepath.Join(outDir, "report_summary.csv")); err != nil {
		t.Fatalf("expected report_summary.csv to exist: %v", err)
	}

	// Verify desired-state docs were generated in docs/database, NOT in outDir
	defaultDocsDir := filepath.Join(tmp, "docs", "database")
	if _, err := os.Stat(filepath.Join(defaultDocsDir, "table_unique_ids.md")); err != nil {
		t.Fatalf("expected table_unique_ids.md to exist in docs/database: %v", err)
	}
	for _, rel := range []string{filepath.Join("procedures", "README.md"), filepath.Join("procedures", "procedure_impact_catalog.md")} {
		if _, err := os.Stat(filepath.Join(defaultDocsDir, rel)); err != nil {
			t.Fatalf("expected procedure document %s: %v", rel, err)
		}
	}
	if _, err := os.Stat(filepath.Join(outDir, "table_unique_ids.md")); !os.IsNotExist(err) {
		t.Fatalf("table_unique_ids.md must not exist in out/: %v", err)
	}

	// Test explicit -docs flag
	customDocsDir := filepath.Join(tmp, "custom_docs")
	exitCode = run([]string{
		"rza_migrate_report.exe",
		"-repo", tmp,
		"-out", outDir,
		"-docs", customDocsDir,
		"-pause-on-exit", "false",
	})
	if exitCode != 0 {
		t.Fatalf("expected run with -docs exit code 0, got %d", exitCode)
	}
	if _, err := os.Stat(filepath.Join(customDocsDir, "table_unique_ids.md")); err != nil {
		t.Fatalf("expected table_unique_ids.md to exist in custom_docs: %v", err)
	}
}

func TestCLIRejectsInvalidArguments(t *testing.T) {
	if code := run([]string{"rza_migrate_report.exe", "unexpected"}); code != 1 {
		t.Fatalf("expected positional argument failure, got %d", code)
	}
	if code := run([]string{"rza_migrate_report.exe", "-pause-on-exit", "sometimes"}); code != 1 {
		t.Fatalf("expected invalid pause value failure, got %d", code)
	}
}

func TestCLIDiscoversRepoFromSubfolder(t *testing.T) {
	tmp := t.TempDir()
	dbObjDir := filepath.Join(tmp, "db_object")
	migDir := filepath.Join(tmp, "migrations")
	toolsDir := filepath.Join(tmp, "tools")
	for _, d := range []string{dbObjDir, migDir, toolsDir} {
		if err := os.MkdirAll(d, 0755); err != nil {
			t.Fatal(err)
		}
	}

	fakeExeInTools := filepath.Join(toolsDir, "rza_migrate_report.exe")
	exitCode := run([]string{
		fakeExeInTools,
		"-pause-on-exit", "false",
	})
	if exitCode != 0 {
		t.Fatalf("expected run exit code 0, got %d", exitCode)
	}

	defaultOutDir := filepath.Join(tmp, "out", "report")
	if _, err := os.Stat(filepath.Join(defaultOutDir, "report_summary.csv")); err != nil {
		t.Fatalf("expected report_summary.csv in default out dir: %v", err)
	}
	defaultDocsDir := filepath.Join(tmp, "docs", "database")
	if _, err := os.Stat(filepath.Join(defaultDocsDir, "table_unique_ids.md")); err != nil {
		t.Fatalf("expected table_unique_ids.md in default docs dir: %v", err)
	}
}
