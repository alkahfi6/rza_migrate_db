package migrest

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestValidateAppliedHistoryFilesRejectsMissingChangedAndConflictingEvidence(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	name := "10000000000000003_tables_dbo.Safe_v1.up.sql"
	body := []byte("SELECT 1;\n")
	if err := os.WriteFile(filepath.Join(dir, name), body, 0o644); err != nil {
		t.Fatal(err)
	}
	hash := fmt.Sprintf("%x", sha256.Sum256(body))
	if err := validateAppliedHistoryFiles(dir, []appliedHistoryEvidence{{FileName: name, SHA256: hash}}); err != nil {
		t.Fatalf("valid evidence rejected: %v", err)
	}

	if err := os.WriteFile(filepath.Join(dir, name), []byte("SELECT 2;\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := validateAppliedHistoryFiles(dir, []appliedHistoryEvidence{{FileName: name, SHA256: hash}}); err == nil ||
		!strings.Contains(err.Error(), "immutable migration history mismatch") {
		t.Fatalf("changed history was not rejected: %v", err)
	}

	missing := "10000000000000004_tables_dbo.Missing_v1.up.sql"
	if err := validateAppliedHistoryFiles(dir, []appliedHistoryEvidence{{FileName: missing, SHA256: hash}}); err == nil {
		t.Fatal("missing historical file was not rejected")
	}

	if err := validateAppliedHistoryFiles(dir, []appliedHistoryEvidence{
		{FileName: name, SHA256: hash},
		{FileName: name, SHA256: strings.Repeat("0", 64)},
	}); err == nil {
		t.Fatal("conflicting historical evidence was not rejected")
	}
}
