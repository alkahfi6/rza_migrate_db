package output

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteWithOptionsOmitsNotesWhenDisabled(t *testing.T) {
	t.Parallel()

	report := NewReport("p", "LF")
	report.Notes["debug_key"] = "debug_value"

	path := filepath.Join(t.TempDir(), "output.json")
	if err := report.WriteWithOptions(path, false); err != nil {
		t.Fatalf("WriteWithOptions should not fail: %v", err)
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read output file failed: %v", err)
	}
	if strings.Contains(string(raw), `"notes"`) {
		t.Fatalf("notes key must be omitted when disabled, got: %s", string(raw))
	}
}

func TestWriteWithOptionsIncludesNotesWhenEnabled(t *testing.T) {
	t.Parallel()

	report := NewReport("p", "LF")
	report.Notes["debug_key"] = "debug_value"

	path := filepath.Join(t.TempDir(), "output.json")
	if err := report.WriteWithOptions(path, true); err != nil {
		t.Fatalf("WriteWithOptions should not fail: %v", err)
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read output file failed: %v", err)
	}
	if !strings.Contains(string(raw), `"notes"`) {
		t.Fatalf("notes key must be present when enabled, got: %s", string(raw))
	}
}
