package engine

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rza_migrate_db/internal/config"
)

func TestRunGenNoChangesForTableSemicolonAndCommaStyleOnly(t *testing.T) {
	t.Parallel()

	repo := repoRootFromEnginePackageDir(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "07_vk_whitespace_only", "input")
	workRoot := t.TempDir()
	if err := copyDirForAutoIndentTest(srcScenario, workRoot); err != nil {
		t.Fatalf("copy scenario failed: %v", err)
	}

	tablePath := filepath.Join(workRoot, "db_object", "schema", "tables", "dbo.TableA.sql")
	beforeRaw, err := os.ReadFile(tablePath)
	if err != nil {
		t.Fatalf("read source table failed: %v", err)
	}

	// Cosmetic-only rewrite: convert table body to leading-comma style
	// and remove terminal semicolon (semantics should remain identical).
	mutated := strings.ReplaceAll(string(beforeRaw), "\r\n", "\n")
	mutated = `CREATE TABLE dbo.TableA (
  id INT NOT NULL
  ,status VARCHAR(10) NULL CONSTRAINT DF_TableA__status DEFAULT ('NEW')
  ,amount DECIMAL(10,2) NULL
  ,CONSTRAINT PK_TableA PRIMARY KEY (id)
)
`
	if err := os.WriteFile(tablePath, []byte(mutated), 0o644); err != nil {
		t.Fatalf("write mutated table failed: %v", err)
	}

	cfg := config.Default(workRoot)
	cfg.EnableStrictMode = false
	eng := New(cfg)
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("gen failed: %v", err)
	}
	if report.Fail.Count > 0 {
		t.Fatalf("expected no fail, got: %+v", report.Fail.Items)
	}
	if len(report.NewMigrations.Files) != 0 {
		t.Fatalf("expected no new migrations for cosmetic-only change, got: %#v", report.NewMigrations.Files)
	}

	hasNoChanges := false
	hasDriftWarn := false
	for _, item := range report.Pass.Items {
		if item.Rule == "GEN_NO_CHANGES" {
			hasNoChanges = true
		}
		if item.Rule == "REPO_STATE_DRIFT" && item.Object == "dbo.TableA" {
			hasDriftWarn = true
		}
	}
	if !hasNoChanges {
		t.Fatalf("expected GEN_NO_CHANGES pass item, got: %+v", report.Pass.Items)
	}
	if hasDriftWarn {
		t.Fatalf("unexpected REPO_STATE_DRIFT for cosmetic-only table change")
	}
}

func TestRunGenNoChangesForIndexSpacingOnly(t *testing.T) {
	t.Parallel()

	repo := repoRootFromEnginePackageDir(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "07_vk_whitespace_only", "input")
	workRoot := t.TempDir()
	if err := copyDirForAutoIndentTest(srcScenario, workRoot); err != nil {
		t.Fatalf("copy scenario failed: %v", err)
	}

	indexPath := filepath.Join(workRoot, "db_object", "schema", "indexes", "dbo.TableA.sql")
	beforeRaw, err := os.ReadFile(indexPath)
	if err != nil {
		t.Fatalf("read source index failed: %v", err)
	}

	// Cosmetic-only rewrite: widen spacing around ON clause.
	mutated := strings.ReplaceAll(string(beforeRaw), " ON ", " ON    ")
	if err := os.WriteFile(indexPath, []byte(mutated), 0o644); err != nil {
		t.Fatalf("write mutated index failed: %v", err)
	}

	cfg := config.Default(workRoot)
	cfg.EnableStrictMode = false
	eng := New(cfg)
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("gen failed: %v", err)
	}
	if report.Fail.Count > 0 {
		t.Fatalf("expected no fail, got: %+v", report.Fail.Items)
	}
	if len(report.NewMigrations.Files) != 0 {
		t.Fatalf("expected no new migrations for cosmetic-only index spacing change, got: %#v", report.NewMigrations.Files)
	}

	hasNoChanges := false
	for _, item := range report.Pass.Items {
		if item.Rule == "GEN_NO_CHANGES" {
			hasNoChanges = true
		}
	}
	if !hasNoChanges {
		t.Fatalf("expected GEN_NO_CHANGES pass item, got: %+v", report.Pass.Items)
	}
}
