package engine

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestDesiredSnapshotHashUsesContentNotOnlyStat(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "dbo.T.sql")
	stamp := time.Unix(1700000000, 0)
	writeSameStat := func(body string) {
		t.Helper()
		if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
		if err := os.Chtimes(path, stamp, stamp); err != nil {
			t.Fatal(err)
		}
	}
	writeSameStat("SELECT Code;")
	tasks := []desiredCollectTask{{AbsFile: path, RelFile: "db_object/views/dbo.T.sql"}}
	first, err := buildDesiredSnapshotHash(tasks)
	if err != nil {
		t.Fatal(err)
	}
	writeSameStat("SELECT code;") // same byte length and timestamp, different identifier case
	second, err := buildDesiredSnapshotHash(tasks)
	if err != nil {
		t.Fatal(err)
	}
	if first == second {
		t.Fatal("content-changing, stat-preserving edit must invalidate desired cache")
	}
}

func TestMigrationSnapshotHashUsesContentNotOnlyStat(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "20260101000000000_views_dbo.T_v1.up.sql")
	stamp := time.Unix(1700000000, 0)
	writeSameStat := func(body string) []os.DirEntry {
		t.Helper()
		if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
		if err := os.Chtimes(path, stamp, stamp); err != nil {
			t.Fatal(err)
		}
		entries, err := os.ReadDir(dir)
		if err != nil {
			t.Fatal(err)
		}
		return entries
	}
	firstEntries := writeSameStat("SELECT Code;")
	first, _, err := buildMigrationsSnapshotHash(dir, firstEntries)
	if err != nil {
		t.Fatal(err)
	}
	secondEntries := writeSameStat("SELECT code;")
	second, _, err := buildMigrationsSnapshotHash(dir, secondEntries)
	if err != nil {
		t.Fatal(err)
	}
	if first == second {
		t.Fatal("content-changing, stat-preserving edit must invalidate reconstruction cache")
	}
}
