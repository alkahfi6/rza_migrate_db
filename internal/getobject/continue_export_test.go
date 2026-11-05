package getobject

import (
	"os"
	"path/filepath"
	"testing"
)

func TestContinueExportsGoodTablePreservesFailedSnapshots(t *testing.T) {
	root := t.TempDir()
	cfg := DefaultAppConfig(root)
	tables := []tableMeta{{ObjectID: 1, Name: "Good"}, {ObjectID: 2, Name: "Broken"}}
	cache := tableSnapshotCache{Columns: map[int][]columnMeta{1: {{ColumnID: 1, Name: "Id", TypeName: "int"}}}}
	tableFiles, indexFiles, warnings, err := buildTableAndIndexFiles(cfg, tables, cache)
	if err != nil || len(tableFiles) != 1 || len(warnings) != 1 || warnings[0].Object != "dbo.Broken" {
		t.Fatalf("continuation failed: %v %v", err, warnings)
	}
	preserve := preserveNamesFromTableWarnings(warnings)
	for _, category := range []string{"tables", "indexes"} {
		dir := filepath.Join(root, "db_object", category)
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, "dbo.Broken.sql"), []byte("old snapshot"), 0600); err != nil {
			t.Fatal(err)
		}
	}
	history := filepath.Join(root, "migrations", "history.sql")
	if err := os.MkdirAll(filepath.Dir(history), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(history, []byte("immutable"), 0600); err != nil {
		t.Fatal(err)
	}
	for _, dryRun := range []bool{true, false, false} {
		for category, files := range map[string]map[string]string{"tables": tableFiles, "indexes": indexFiles} {
			if _, err := applySQLCategoryPreserving(cfg, "db_object/"+category, files, preserve, dryRun); err != nil {
				t.Fatal(err)
			}
			raw, err := os.ReadFile(filepath.Join(root, "db_object", category, "dbo.Broken.sql"))
			if err != nil || string(raw) != "old snapshot" {
				t.Fatal("failed object snapshot lost")
			}
		}
		_, err := os.Stat(filepath.Join(root, "db_object/tables/dbo.Good.sql"))
		if dryRun && !os.IsNotExist(err) {
			t.Fatal("dry-run wrote snapshot")
		}
		if !dryRun && err != nil {
			t.Fatal("good table was not extracted")
		}
	}
	raw, err := os.ReadFile(history)
	if err != nil || string(raw) != "immutable" {
		t.Fatal("migration history changed")
	}
	cfg.SkipUnsupportedTables = false
	if _, _, _, err := buildTableAndIndexFiles(cfg, tables, cache); err == nil {
		t.Fatal("explicit strict policy must still fail")
	}
}
