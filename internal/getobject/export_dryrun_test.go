package getobject

import (
	"os"
	"path/filepath"
	"testing"
)

func TestApplySQLCategoryDryRunDoesNotMutateFiles(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	cfg := DefaultAppConfig(root)
	relDir := filepath.Join("db_object", "schema", "tables")
	dir := cfg.RepoPath(relDir)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}

	existingA := filepath.Join(dir, "dbo.TableA.sql")
	existingB := filepath.Join(dir, "dbo.TableB.sql")
	if err := os.WriteFile(existingA, []byte("CREATE TABLE dbo.TableA (\r\n  id INT NOT NULL\r\n);\r\n"), 0o644); err != nil {
		t.Fatalf("write existingA failed: %v", err)
	}
	if err := os.WriteFile(existingB, []byte("CREATE TABLE dbo.TableB (\n  id INT NOT NULL\n);\n"), 0o644); err != nil {
		t.Fatalf("write existingB failed: %v", err)
	}

	files := map[string]string{
		"dbo.TableA.sql": "CREATE TABLE dbo.TableA (\n  id BIGINT NOT NULL\n);\n",
		"dbo.TableC.sql": "CREATE TABLE dbo.TableC (\n  id INT NOT NULL\n);\n",
	}
	change, err := applySQLCategory(cfg, relDir, files, true)
	if err != nil {
		t.Fatalf("applySQLCategory dry-run failed: %v", err)
	}

	if len(change.Updated) != 1 || len(change.Created) != 1 || len(change.Deleted) != 1 {
		t.Fatalf("unexpected dry-run change summary: %+v", change)
	}

	afterA, err := os.ReadFile(existingA)
	if err != nil {
		t.Fatalf("read existingA failed: %v", err)
	}
	if string(afterA) != "CREATE TABLE dbo.TableA (\r\n  id INT NOT NULL\r\n);\r\n" {
		t.Fatalf("dry-run must not mutate existingA")
	}
	if _, err := os.Stat(existingB); err != nil {
		t.Fatalf("dry-run must not delete existingB: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "dbo.TableC.sql")); !os.IsNotExist(err) {
		t.Fatalf("dry-run must not create TableC file")
	}
}

func TestApplySQLCategoryApplyMutatesFiles(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	cfg := DefaultAppConfig(root)
	relDir := filepath.Join("db_object", "schema", "tables")
	dir := cfg.RepoPath(relDir)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}

	existingA := filepath.Join(dir, "dbo.TableA.sql")
	existingB := filepath.Join(dir, "dbo.TableB.sql")
	if err := os.WriteFile(existingA, []byte("CREATE TABLE dbo.TableA (\n  id INT NOT NULL\n);\n"), 0o644); err != nil {
		t.Fatalf("write existingA failed: %v", err)
	}
	if err := os.WriteFile(existingB, []byte("CREATE TABLE dbo.TableB (\n  id INT NOT NULL\n);\n"), 0o644); err != nil {
		t.Fatalf("write existingB failed: %v", err)
	}

	files := map[string]string{
		"dbo.TableA.sql": "CREATE TABLE dbo.TableA (\n  id BIGINT NOT NULL\n);\n",
		"dbo.TableC.sql": "CREATE TABLE dbo.TableC (\n  id INT NOT NULL\n);\n",
	}
	change, err := applySQLCategory(cfg, relDir, files, false)
	if err != nil {
		t.Fatalf("applySQLCategory apply failed: %v", err)
	}
	if len(change.Updated) != 1 || len(change.Created) != 1 || len(change.Deleted) != 1 {
		t.Fatalf("unexpected apply change summary: %+v", change)
	}

	afterA, err := os.ReadFile(existingA)
	if err != nil {
		t.Fatalf("read existingA failed: %v", err)
	}
	if string(afterA) != "CREATE TABLE dbo.TableA (\n  id BIGINT NOT NULL\n);\n" {
		t.Fatalf("apply must update existingA content")
	}
	if _, err := os.Stat(existingB); !os.IsNotExist(err) {
		t.Fatalf("apply must delete existingB")
	}
	if _, err := os.Stat(filepath.Join(dir, "dbo.TableC.sql")); err != nil {
		t.Fatalf("apply must create TableC: %v", err)
	}
}

func TestApplySQLCategoryPreservesUnrenderableDatabaseObject(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	cfg := DefaultAppConfig(root)
	relDir := filepath.Join("db_object", "tables")
	dir := cfg.RepoPath(relDir)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	preserved := filepath.Join(dir, "dbo.EncryptedOrUnsupported.sql")
	if err := os.WriteFile(preserved, []byte("CREATE TABLE dbo.EncryptedOrUnsupported (id int NOT NULL);\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	change, err := applySQLCategoryPreserving(cfg, relDir, map[string]string{}, map[string]struct{}{
		"dbo.EncryptedOrUnsupported.sql": {},
	}, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(change.Deleted) != 0 {
		t.Fatalf("preserved object must not be reported deleted: %+v", change)
	}
	if _, err := os.Stat(preserved); err != nil {
		t.Fatalf("last good source snapshot was deleted: %v", err)
	}
}
