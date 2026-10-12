package getobject

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadInitialMigrationObjectsUsesOnlyV1UpObjectMigrations(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	migrationsDir := filepath.Join(root, "migrations")
	if err := os.MkdirAll(migrationsDir, 0o755); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{
		"20260814000000001_tables_dbo.InitialTable_v1.up.sql",
		"20260814000000001_tables_dbo.InitialTable_v1.down.sql",
		"20260814000000002_procedures_dbo.InitialProc_v1.up.sql",
		"20260814000000003_views_dbo.ChangedView_v2.up.sql",
		"20260814000000004_patch_manual_fix_v1.up.sql",
		"not-a-migration.sql",
	} {
		if err := os.WriteFile(filepath.Join(migrationsDir, name), []byte("SELECT 1;\n"), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	cfg := DefaultAppConfig(root)
	objects, err := loadInitialMigrationObjects(cfg)
	if err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"tables|dbo.InitialTable", "procedures|dbo.InitialProc"} {
		if _, ok := objects[key]; !ok {
			t.Fatalf("expected initial migration object %q to be excluded: %#v", key, objects)
		}
	}
	for _, key := range []string{"views|dbo.ChangedView", "patch|manual_fix"} {
		if _, ok := objects[key]; ok {
			t.Fatalf("non-initial object %q must not be excluded", key)
		}
	}
}

func TestShouldExcludeExtractedObjectMatchesCategoryAndExactObjectCase(t *testing.T) {
	t.Parallel()

	cfg := DefaultAppConfig(".")
	cfg.initialMigrationObjects = map[string]struct{}{
		"tables|dbo.InitialTable": {},
	}
	if !shouldExcludeExtractedObject(cfg, "tables", "dbo", "InitialTable") {
		t.Fatal("initial table must be excluded")
	}
	if shouldExcludeExtractedObject(cfg, "views", "dbo", "InitialTable") {
		t.Fatal("a different category must not be excluded")
	}
	if shouldExcludeExtractedObject(cfg, "tables", "dbo", "initialtable") {
		t.Fatal("case-distinct object must not be excluded")
	}
}

func TestExcludeInitialMigrationFilesRemovesInitialIndexSnapshot(t *testing.T) {
	t.Parallel()

	cfg := DefaultAppConfig(".")
	cfg.initialMigrationObjects = map[string]struct{}{
		"indexes|dbo.InitialTable": {},
	}
	files := map[string]string{
		"dbo.InitialTable.sql": "CREATE INDEX IX_Initial ON dbo.InitialTable(id);",
		"dbo.NewTable.sql":     "CREATE INDEX IX_New ON dbo.NewTable(id);",
	}
	excludeInitialMigrationFiles(cfg, "indexes", files)
	if _, ok := files["dbo.InitialTable.sql"]; ok {
		t.Fatal("initial index snapshot must be excluded")
	}
	if _, ok := files["dbo.NewTable.sql"]; !ok {
		t.Fatal("non-initial index snapshot must remain")
	}
}

func TestLoadInitialMigrationObjectsAllowsMissingMigrationsDirectory(t *testing.T) {
	t.Parallel()

	objects, err := loadInitialMigrationObjects(DefaultAppConfig(t.TempDir()))
	if err != nil {
		t.Fatal(err)
	}
	if len(objects) != 0 {
		t.Fatalf("expected no exclusions, got %#v", objects)
	}
}
