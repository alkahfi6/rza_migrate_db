package fsx

import (
	"os"
	"path/filepath"
	"sort"
	"testing"
)

func TestResolveDbObjectCategoryFilesPrefersCanonicalAndFallsBackLegacy(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	repoPath := func(rel string) string { return filepath.Join(root, rel) }

	canonicalDir := repoPath(filepath.Join("db_object", "tables"))
	legacyDir := repoPath(filepath.Join("db_object", "schema", "tables"))
	if err := os.MkdirAll(canonicalDir, 0o755); err != nil {
		t.Fatalf("mkdir canonical: %v", err)
	}
	if err := os.MkdirAll(legacyDir, 0o755); err != nil {
		t.Fatalf("mkdir legacy: %v", err)
	}

	// same filename in both locations -> canonical should win.
	if err := os.WriteFile(filepath.Join(canonicalDir, "dbo.TableA.sql"), []byte("A-new"), 0o644); err != nil {
		t.Fatalf("write canonical A: %v", err)
	}
	if err := os.WriteFile(filepath.Join(legacyDir, "dbo.TableA.sql"), []byte("A-old"), 0o644); err != nil {
		t.Fatalf("write legacy A: %v", err)
	}
	// legacy-only file should still be included.
	if err := os.WriteFile(filepath.Join(legacyDir, "dbo.TableB.sql"), []byte("B-old"), 0o644); err != nil {
		t.Fatalf("write legacy B: %v", err)
	}

	files, err := ResolveDbObjectCategoryFiles(repoPath, ListFilesRecursive, "db_object", "tables")
	if err != nil {
		t.Fatalf("resolve files: %v", err)
	}
	if len(files) != 2 {
		t.Fatalf("expected 2 files, got %d: %#v", len(files), files)
	}
	bases := []string{filepath.Base(files[0]), filepath.Base(files[1])}
	sort.Strings(bases)
	if bases[0] != "dbo.TableA.sql" || bases[1] != "dbo.TableB.sql" {
		t.Fatalf("unexpected files: %#v", files)
	}

	var resolvedA string
	for _, file := range files {
		if filepath.Base(file) == "dbo.TableA.sql" {
			resolvedA = file
			break
		}
	}
	if resolvedA == "" {
		t.Fatalf("resolved TableA file not found: %#v", files)
	}

	aContent, err := os.ReadFile(resolvedA)
	if err != nil {
		t.Fatalf("read resolved A: %v", err)
	}
	if string(aContent) != "A-new" {
		t.Fatalf("expected canonical file to win for TableA, got %q", string(aContent))
	}
}

func TestResolveDbObjectCategoryFilesLegacyOnly(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	repoPath := func(rel string) string { return filepath.Join(root, rel) }

	legacyDir := repoPath(filepath.Join("db_object", "schema", "procedures"))
	if err := os.MkdirAll(legacyDir, 0o755); err != nil {
		t.Fatalf("mkdir legacy: %v", err)
	}
	if err := os.WriteFile(filepath.Join(legacyDir, "dbo.ProcA.sql"), []byte("CREATE OR ALTER PROCEDURE dbo.ProcA AS SELECT 1;\n"), 0o644); err != nil {
		t.Fatalf("write file: %v", err)
	}

	files, err := ResolveDbObjectCategoryFiles(repoPath, ListFilesRecursive, "db_object", "procedures")
	if err != nil {
		t.Fatalf("resolve files: %v", err)
	}
	if len(files) != 1 {
		t.Fatalf("expected 1 file, got %d", len(files))
	}
	if filepath.Base(files[0]) != "dbo.ProcA.sql" {
		t.Fatalf("unexpected file: %s", files[0])
	}
}
