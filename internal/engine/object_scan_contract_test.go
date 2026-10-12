package engine

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"rza_migrate_db/internal/config"
)

func TestObjectScanPreservesCheckerAndGeneratorBoundaries(t *testing.T) {
	root := t.TempDir()
	source := filepath.Join(repoRootFromEnginePackageDir(t), "fixtures", "qa_rules", "07_vk_whitespace_only", "input")
	if err := copyDirForAutoIndentTest(source, root); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(root, "db_object", "procedures", "dbo.Precise.sql")
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	body := `CREATE OR ALTER PROCEDURE dbo.Precise AS
BEGIN
  UPDATE TOP (1) t SET status='READY'
  FROM dbo.TableA t WHERE t.id=1;
  SELECT id INTO #work FROM dbo.TableA;
END;
`
	if err := os.WriteFile(path, []byte(body), 0644); err != nil {
		t.Fatal(err)
	}
	if report, err := New(config.Default(root)).Run("fmt"); err != nil || report.Fail.Count != 0 {
		t.Fatalf("fmt: %v %+v", err, report)
	}
	migrations := hashSQLTree(t, filepath.Join(root, "migrations"))
	objects := hashSQLTree(t, filepath.Join(root, "db_object"))
	for _, plan := range []bool{true, false, true} {
		cfg := config.Default(root)
		cfg.EnableFmtCheck = plan
		report, err := New(cfg).Run("check")
		if err != nil || report.Fail.Count != 0 || len(report.NewMigrations.Files) != 0 {
			t.Fatalf("check plan=%v: err=%v report=%+v", plan, err, report)
		}
		if !reflect.DeepEqual(migrations, hashSQLTree(t, filepath.Join(root, "migrations"))) || !reflect.DeepEqual(objects, hashSQLTree(t, filepath.Join(root, "db_object"))) {
			t.Fatal("analysis changed authored SQL/history")
		}
	}
	report, err := New(config.Default(root)).Run("gen")
	if err != nil || report.Fail.Count != 0 || len(report.NewMigrations.Files) != 1 {
		t.Fatalf("gen: %v %+v", err, report)
	}
	entries, err := os.ReadDir(filepath.Join(root, "migrations"))
	if err != nil {
		t.Fatal(err)
	}
	foundUp, foundDown := false, false
	for _, e := range entries {
		if !strings.Contains(e.Name(), "procedures_dbo.Precise") {
			continue
		}
		raw, err := os.ReadFile(filepath.Join(root, "migrations", e.Name()))
		if err != nil {
			t.Fatal(err)
		}
		if strings.HasSuffix(e.Name(), ".up.sql") {
			foundUp = strings.Contains(string(raw), "UPDATE TOP (1) t") && strings.Contains(string(raw), "FROM dbo.TableA t") && strings.Contains(string(raw), "'READY'")
		}
		if strings.HasSuffix(e.Name(), ".down.sql") {
			foundDown = strings.Contains(string(raw), "DROP PROCEDURE")
		}
	}
	if !foundUp || !foundDown {
		t.Fatal("genuine module delta/up-down behavior was not preserved")
	}
	again, err := New(config.Default(root)).Run("gen")
	if err != nil || again.Fail.Count != 0 || len(again.NewMigrations.Files) != 0 {
		t.Fatalf("second gen not idempotent: %v %+v", err, again)
	}
}
