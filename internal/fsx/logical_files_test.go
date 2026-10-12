package fsx

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveLogicalFilesCaseInsensitiveReturnsActualEntryPaths(t *testing.T) {
	dir := t.TempDir()
	for _, name := range []string{"readme.md", "UP.SQL", "Down.Sql"} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(name), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	paths, ok, err := ResolveLogicalFilesCaseInsensitive(dir, "README.md", "up.sql", "down.sql")
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("case variants must resolve")
	}
	for logical, actual := range map[string]string{
		"readme.md": "readme.md",
		"up.sql":    "UP.SQL",
		"down.sql":  "Down.Sql",
	} {
		if filepath.Base(paths[logical]) != actual {
			t.Fatalf("%s resolved to %q, want %q", logical, paths[logical], actual)
		}
	}
}

func TestResolveLogicalFilesCaseInsensitiveRejectsMissingEntry(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "up.sql"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, ok, err := ResolveLogicalFilesCaseInsensitive(dir, "README.md", "up.sql", "down.sql"); err != nil || ok {
		t.Fatalf("missing logical entries must return ok=false, err=nil; ok=%t err=%v", ok, err)
	}
}
