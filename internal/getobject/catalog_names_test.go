package getobject

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCatalogNamesExtractWithoutAuthoringNameRules(t *testing.T) {
	for _, name := range []string{"Customer Detail", "Customer-Detail", "2026Customer", "Customer.Archive", "客户", "a]b", " leading ", "a/b", `a\b`, "a%2Fb", "a:b?*"} {
		t.Run(name, func(t *testing.T) {
			root := t.TempDir()
			cfg := DefaultAppConfig(root)
			if err := os.Mkdir(filepath.Join(root, "migrations"), 0755); err != nil {
				t.Fatal(err)
			}
			history := filepath.Join(root, "migrations", "immutable.sql")
			if err := os.WriteFile(history, []byte("history"), 0600); err != nil {
				t.Fatal(err)
			}
			sql := "CREATE TABLE " + qualifiedName("dbo", name) + " ([Id] INT NULL);\n"
			if !strings.Contains(sql, quoteIdentifier(name)) {
				t.Fatal("catalog name was not delimited exactly")
			}
			file := fileNameForObject(name)
			if filepath.Base(file) != file || strings.ContainsAny(file, `/\:*?"<>|`) {
				t.Fatalf("unsafe filename %q", file)
			}
			files := map[string]string{file: sql}
			change, err := applySQLCategory(cfg, "db_object/tables", files, true)
			if err != nil || len(change.Created) != 1 {
				t.Fatalf("plan failed: %v", err)
			}
			if _, err := os.Stat(filepath.Join(root, "db_object")); !os.IsNotExist(err) {
				t.Fatal("plan wrote snapshots")
			}
			if _, err := applySQLCategory(cfg, "db_object/tables", files, false); err != nil {
				t.Fatal(err)
			}
			raw, err := os.ReadFile(filepath.Join(root, "db_object/tables", file))
			if err != nil || !strings.Contains(string(raw), qualifiedName("dbo", name)) {
				t.Fatal("snapshot lost SQL identity")
			}
			change, err = applySQLCategory(cfg, "db_object/tables", files, false)
			if err != nil || len(change.Unchanged) != 1 {
				t.Fatal("second extraction not idempotent")
			}
			raw, err = os.ReadFile(history)
			if err != nil || string(raw) != "history" {
				t.Fatal("migration history changed")
			}
			module := "CREATE PROCEDURE " + qualifiedName("dbo", name) + " AS SELECT 1;"
			normalized, err := ensureCreateOrAlter(module, "PROCEDURE")
			if err != nil || !strings.Contains(normalized, quoteIdentifier(name)) {
				t.Fatalf("module name changed: %s %v", normalized, err)
			}
		})
	}
	if fileNameForObject("a/b") == fileNameForObject("a%2Fb") {
		t.Fatal("encoded filenames collide")
	}
	warnings := warningsFromPreserved("procedures", "dbo", map[string]struct{}{fileNameForObject("a/b"): {}})
	if warnings[0].Object != "dbo.a/b" {
		t.Fatal("warning lost catalog identity")
	}
}
