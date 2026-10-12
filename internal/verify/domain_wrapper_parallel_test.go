package verify

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"rza_migrate_db/internal/output"
)

func TestValidateWrapperSubsetParallelMatchesSerialDeterministically(t *testing.T) {
	root := t.TempDir()
	metas := []MigrationMeta{
		{RelPath: "migrations/001_tables_dbo.TableA_v1.up.sql", Category: "tables"},
		{RelPath: "migrations/002_indexes_dbo.TableA_v1.up.sql", Category: "indexes"},
		{RelPath: "migrations/003_patch_seed_v1.up.sql", Category: "patch"},
	}
	for index, meta := range metas {
		abs := filepath.Join(root, filepath.FromSlash(meta.RelPath))
		if err := os.MkdirAll(filepath.Dir(abs), 0o755); err != nil {
			t.Fatal(err)
		}
		content := []byte("-- deliberately incomplete wrapper\nSELECT " + string(rune('1'+index)) + ";\n")
		if err := os.WriteFile(abs, content, 0o644); err != nil {
			t.Fatal(err)
		}
	}
	base := WrapperSubsetOptions{
		RepoPath:        func(rel string) string { return filepath.Join(root, filepath.FromSlash(rel)) },
		Metas:           metas,
		ParallelWorkers: 3,
	}
	serial := output.NewReport("", "")
	if err := ValidateWrapperSubset(&serial, base); err != nil {
		t.Fatal(err)
	}
	parallelOpts := base
	parallelOpts.EnableParallelRead = true
	parallel := output.NewReport("", "")
	if err := ValidateWrapperSubset(&parallel, parallelOpts); err != nil {
		t.Fatal(err)
	}
	serialJSON, _ := json.Marshal(serial)
	parallelJSON, _ := json.Marshal(parallel)
	if string(serialJSON) != string(parallelJSON) {
		t.Fatalf("parallel wrapper report differs from serial\nserial: %s\nparallel: %s", serialJSON, parallelJSON)
	}
}
