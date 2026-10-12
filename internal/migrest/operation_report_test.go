package migrest

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestBuildAppliedMigrationFilesProducesExactDirectionalManifest(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	writePair := func(version uint64, object string) {
		t.Helper()
		base := fmt.Sprintf("%d_tables_dbo.%s_v1", version, object)
		for _, direction := range []string{"up", "down"} {
			path := filepath.Join(dir, base+"."+direction+".sql")
			if err := os.WriteFile(path, []byte(direction+" "+object+"\n"), 0o644); err != nil {
				t.Fatalf("write fixture %s: %v", path, err)
			}
		}
	}
	writePair(10000000000000001, "One")
	writePair(10000000000000002, "Two")
	writePair(10000000000000003, "Three")

	beforeValue := uint64(10000000000000001)
	afterValue := uint64(10000000000000003)
	files, manifest, direction, err := buildAppliedMigrationFiles(
		dir,
		VersionState{Version: &beforeValue},
		VersionState{Version: &afterValue},
	)
	if err != nil {
		t.Fatalf("build up report: %v", err)
	}
	if direction != "up" || len(files) != 2 {
		t.Fatalf("unexpected up range: direction=%s files=%#v", direction, files)
	}
	if files[0].Version != 10000000000000002 || files[1].Version != 10000000000000003 {
		t.Fatalf("up report is not ordered/ranged correctly: %#v", files)
	}
	expectedFileHash := sha256.Sum256([]byte("up Two\n"))
	if files[0].SHA256 != fmt.Sprintf("%x", expectedFileHash) {
		t.Fatalf("file checksum mismatch: got=%s", files[0].SHA256)
	}
	if len(manifest) != 64 {
		t.Fatalf("manifest checksum must be SHA-256: %q", manifest)
	}

	files, manifest, direction, err = buildAppliedMigrationFiles(
		dir,
		VersionState{Version: &afterValue},
		VersionState{Version: &beforeValue},
	)
	if err != nil {
		t.Fatalf("build down report: %v", err)
	}
	if direction != "down" || len(files) != 2 {
		t.Fatalf("unexpected down range: direction=%s files=%#v", direction, files)
	}
	if files[0].Version != 10000000000000003 || files[1].Version != 10000000000000002 {
		t.Fatalf("down report is not ordered/ranged correctly: %#v", files)
	}
	if manifest == "" {
		t.Fatal("down report manifest is empty")
	}
}
