package engine

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/output"
)

func TestRepositorySnapshotInvalidatesSameSizeSameTimestampChange(t *testing.T) {
	root := t.TempDir()
	for _, dir := range []string{"db_object/tables", "migrations", "patch"} {
		if err := os.MkdirAll(filepath.Join(root, filepath.FromSlash(dir)), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	rel := "db_object/tables/dbo.CaseExact.sql"
	abs := filepath.Join(root, filepath.FromSlash(rel))
	before := []byte("CREATE TABLE [dbo].[CaseExact] ([Id] int);\n")
	after := []byte("CREATE TABLE [dbo].[CaseExact] ([ID] int);\n")
	if len(before) != len(after) {
		t.Fatal("test fixture must preserve byte length")
	}
	if err := os.WriteFile(abs, before, 0o644); err != nil {
		t.Fatal(err)
	}
	fixedTime := time.Unix(1_700_000_000, 0)
	if err := os.Chtimes(abs, fixedTime, fixedTime); err != nil {
		t.Fatal(err)
	}

	e := New(config.Default(root))
	report := output.NewReport("", "")
	if err := e.ensureRepositorySnapshot(&report); err != nil {
		t.Fatal(err)
	}
	firstHash, ok := e.snapshotHashForRelativePaths([]string{rel})
	if !ok {
		t.Fatal("expected first snapshot hash")
	}

	if err := os.WriteFile(abs, after, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Chtimes(abs, fixedTime, fixedTime); err != nil {
		t.Fatal(err)
	}
	e.resetRepositorySnapshot()
	if err := e.ensureRepositorySnapshot(&report); err != nil {
		t.Fatal(err)
	}
	secondHash, ok := e.snapshotHashForRelativePaths([]string{rel})
	if !ok {
		t.Fatal("expected second snapshot hash")
	}
	if firstHash == secondHash {
		t.Fatal("content-addressed snapshot did not detect same-size, same-timestamp mutation")
	}
	raw, err := e.readRepositoryFile(rel)
	if err != nil {
		t.Fatal(err)
	}
	if string(raw) != string(after) {
		t.Fatalf("snapshot returned stale bytes: %q", raw)
	}
}

func TestRepositorySnapshotLookupIsCaseSensitive(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "db_object", "tables"), 0o755); err != nil {
		t.Fatal(err)
	}
	rel := "db_object/tables/dbo.OrderItem.sql"
	if err := os.WriteFile(filepath.Join(root, filepath.FromSlash(rel)), []byte("SELECT 1;\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	e := New(config.Default(root))
	if err := e.ensureRepositorySnapshot(nil); err != nil {
		t.Fatal(err)
	}
	if _, err := e.readRepositoryFile("db_object/tables/dbo.orderitem.sql"); !os.IsNotExist(err) {
		t.Fatalf("case-mismatched snapshot lookup must fail, got %v", err)
	}
}
