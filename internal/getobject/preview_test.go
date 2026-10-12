package getobject

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPreviewModulesFromDir(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	inDir := filepath.Join(root, "input")
	outDir := filepath.Join(root, "output")
	if err := os.MkdirAll(filepath.Join(inDir, "view"), 0o755); err != nil {
		t.Fatalf("mkdir input: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(inDir, "stored-procedure"), 0o755); err != nil {
		t.Fatalf("mkdir input: %v", err)
	}

	viewSQL := "-- note\nCreate VIEW [dbo].[ViewA]\nAS\nSELECT 1 AS a;\n"
	if err := os.WriteFile(filepath.Join(inDir, "view", "ViewA.sql"), []byte(viewSQL), 0o644); err != nil {
		t.Fatalf("write view: %v", err)
	}
	procSQL := "CREATE  proc [dbo].[ProcA]\nAS\nBEGIN\nSELECT 1;\nEND\n"
	if err := os.WriteFile(filepath.Join(inDir, "stored-procedure", "ProcA.sql"), []byte(procSQL), 0o644); err != nil {
		t.Fatalf("write proc: %v", err)
	}
	invalidSQL := "SELECT * FROM sys.sql_modules;\n"
	if err := os.WriteFile(filepath.Join(inDir, "view", "query.sql"), []byte(invalidSQL), 0o644); err != nil {
		t.Fatalf("write invalid: %v", err)
	}

	cfg := DefaultAppConfig(".")
	cfg.ConnectionString = "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable"
	summary, err := PreviewModulesFromDir(cfg, inDir, outDir, nil)
	if err != nil {
		t.Fatalf("preview failed: %v", err)
	}
	if summary.Processed != 3 {
		t.Fatalf("unexpected processed: %d", summary.Processed)
	}
	if summary.Written != 2 {
		t.Fatalf("unexpected written: %d", summary.Written)
	}
	if summary.Failed != 1 {
		t.Fatalf("unexpected failed: %d", summary.Failed)
	}

	viewOut, err := os.ReadFile(filepath.Join(outDir, "views", "ViewA.sql"))
	if err != nil {
		t.Fatalf("read view output: %v", err)
	}
	if strings.Contains(string(viewOut), "-- note") {
		t.Fatalf("comment should be removed from view output")
	}
	if !strings.HasPrefix(string(viewOut), "CREATE OR ALTER VIEW") {
		t.Fatalf("view output should be normalized to CREATE OR ALTER")
	}

	procOut, err := os.ReadFile(filepath.Join(outDir, "procedures", "ProcA.sql"))
	if err != nil {
		t.Fatalf("read proc output: %v", err)
	}
	if !strings.Contains(string(procOut), "BEGIN\n  SELECT 1;\nEND") {
		t.Fatalf("procedure output should be reindented")
	}
}
