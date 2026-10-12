package migreport_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rza_migrate_db/internal/migreport"
)

func TestSuspectedJoinIdentifiersQA31GoldenReadOnlyAndIdempotent(t *testing.T) {
	fixture := filepath.Join("..", "..", "fixtures", "qa_rules", "QA-31-desired-state-documentation")
	var expected struct {
		Candidates struct {
			Rows      []string `json:"required_rows"`
			Evidence  []string `json:"required_evidence"`
			Forbidden []string `json:"forbidden_rows"`
		} `json:"suspected_join_identifiers"`
	}
	raw, err := os.ReadFile(filepath.Join(fixture, "expected", "output.golden.json"))
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(raw, &expected); err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	source := filepath.Join(fixture, "input")
	if err := filepath.Walk(source, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		rel, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		dest := filepath.Join(root, rel)
		if info.IsDir() {
			return os.MkdirAll(dest, 0755)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(dest, data, 0644)
	}); err != nil {
		t.Fatal(err)
	}
	migrations := filepath.Join(root, "migrations")
	if err := os.MkdirAll(migrations, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(migrations, "history-marker.txt"), []byte("immutable history"), 0644); err != nil {
		t.Fatal(err)
	}
	beforeSource, beforeHistory := computeDirHash(t, filepath.Join(root, "db_object")), computeDirHash(t, migrations)
	var first string
	for i := 0; i < 2; i++ {
		if code, err := migreport.Run(migreport.Options{RepoRoot: root, OutputDir: "audit", DocsDir: "documentation"}); err != nil || code != 0 {
			t.Fatalf("report code=%d err=%v", code, err)
		}
		text := readDesiredFixture(t, filepath.Join(root, "documentation", "table_unique_ids.md"))
		start := strings.Index(text, "## Suspected JOIN identifiers")
		if start < 0 {
			t.Fatal("candidate section missing")
		}
		section := text[start:]
		for _, want := range append(expected.Candidates.Rows, expected.Candidates.Evidence...) {
			if !strings.Contains(section, want) {
				t.Fatalf("missing golden %q: %s", want, section)
			}
		}
		for _, bad := range expected.Candidates.Forbidden {
			if strings.Contains(section, bad) {
				t.Fatalf("false candidate %q: %s", bad, section)
			}
		}
		if i > 0 && text != first {
			t.Fatal("second report not deterministic")
		}
		first = text
		if beforeSource != computeDirHash(t, filepath.Join(root, "db_object")) || beforeHistory != computeDirHash(t, migrations) {
			t.Fatal("report mutated authored SQL/history")
		}
	}
}
