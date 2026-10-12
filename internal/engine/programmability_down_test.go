package engine

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/output"
)

func TestInitialProgrammabilityMigrationHasExecutableDown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		category string
		object   string
		up       string
		wantDown string
	}{
		{
			category: "views",
			object:   "dbo.CaseSensitiveView",
			up:       "CREATE OR ALTER VIEW dbo.CaseSensitiveView AS SELECT 1 AS Id;\n",
			wantDown: "DROP VIEW IF EXISTS dbo.CaseSensitiveView;",
		},
		{
			category: "functions",
			object:   "dbo.CaseSensitiveFunction",
			up:       "CREATE OR ALTER FUNCTION dbo.CaseSensitiveFunction() RETURNS INT AS BEGIN RETURN 1; END;\n",
			wantDown: "DROP FUNCTION IF EXISTS dbo.CaseSensitiveFunction;",
		},
		{
			category: "procedures",
			object:   "dbo.Order Detail",
			up:       "CREATE OR ALTER PROCEDURE dbo.[Order Detail] AS BEGIN SELECT 1; END;\n",
			wantDown: "DROP PROCEDURE IF EXISTS dbo.[Order Detail];",
		},
	}

	for _, tt := range tests {
		t.Run(tt.category, func(t *testing.T) {
			cfg := config.Default(t.TempDir())
			engine := New(cfg)
			report := output.NewReport(cfg.OutputNamingPattern, cfg.OutputNamingLineEndings)
			ops, _, _, err := engine.planGenerationWithInputs(
				&report,
				map[string]historyObject{},
				map[string]struct{}{},
				"",
				[]desiredObject{{
					Category:       tt.category,
					Object:         tt.object,
					SourceFile:     "db_object/" + tt.category + "/source.sql",
					SemanticSQL:    tt.up,
					NormalizedHash: hashSemanticByCategory(tt.category, tt.up),
				}},
				nil,
			)
			if err != nil {
				t.Fatalf("plan initial %s: %v", tt.category, err)
			}
			if report.Fail.Count != 0 {
				t.Fatalf("unexpected planning failures: %+v", report.Fail.Items)
			}
			if len(ops) != 1 {
				t.Fatalf("planned operation count=%d, want=1", len(ops))
			}
			if strings.TrimSpace(ops[0].DownSemanticSQL) != tt.wantDown {
				t.Fatalf("down SQL mismatch\nwant: %s\ngot:  %s", tt.wantDown, ops[0].DownSemanticSQL)
			}

			rendered := renderPlannedMigration(ops[0], "20991231235959000", "migration.down.sql", false)
			if !strings.Contains(rendered, tt.wantDown) || strings.Contains(rendered, "Manual restore required") {
				t.Fatalf("rendered down is not executable:\n%s", rendered)
			}
		})
	}
}

func TestGenInitialProgrammabilityEmitsReversiblePairsAndIsIdempotent(t *testing.T) {
	t.Parallel()

	repo := setupMinimalGeneratorRepo(t)
	repositoryRoot := repoRootFromEnginePackageDir(t)
	for _, name := range []string{
		"000000__core__mig__audit_init.up.sql",
		"000000__core__mig__audit_init.down.sql",
		"10000000000000001_procedures_dbo.usp_mig_log_event_v1.up.sql",
		"10000000000000001_procedures_dbo.usp_mig_log_event_v1.down.sql",
		"10000000000000002_procedures_dbo.usp_mig_try_get_lock_v1.up.sql",
		"10000000000000002_procedures_dbo.usp_mig_try_get_lock_v1.down.sql",
	} {
		raw, err := os.ReadFile(filepath.Join(repositoryRoot, "migrations", name))
		if err != nil {
			t.Fatalf("read core migration %s: %v", name, err)
		}
		mustWriteFile(t, filepath.Join(repo, "migrations", name), string(raw))
	}
	sources := []struct {
		category string
		object   string
		body     string
		wantDown string
	}{
		{"views", "dbo.CaseView", "CREATE OR ALTER VIEW dbo.CaseView AS SELECT 1 AS ExactCase;\n", "DROP VIEW IF EXISTS dbo.CaseView;"},
		{"functions", "dbo.CaseFunction", "CREATE OR ALTER FUNCTION dbo.CaseFunction() RETURNS INT AS BEGIN RETURN 1; END;\n", "DROP FUNCTION IF EXISTS dbo.CaseFunction;"},
		{"procedures", "dbo.CaseProcedure", "CREATE OR ALTER PROCEDURE dbo.CaseProcedure AS BEGIN SELECT 1 AS ExactCase; END;\n", "DROP PROCEDURE IF EXISTS dbo.CaseProcedure;"},
	}
	for _, source := range sources {
		mustWriteFile(t, filepath.Join(repo, "db_object", source.category, source.object+".sql"), source.body)
	}

	cfg := config.Default(repo)
	report, err := New(cfg).Run("gen")
	if err != nil {
		t.Fatalf("generate reversible programmability pairs: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("unexpected generation failures: %+v", report.Fail.Items)
	}
	// output.json lists one logical entry per pair (the up filename); both
	// physical directions are asserted below.
	if len(report.NewMigrations.Files) != len(sources) {
		t.Fatalf("generated pair count=%d, want=%d: %v", len(report.NewMigrations.Files), len(sources), report.NewMigrations.Files)
	}

	for _, source := range sources {
		matches, err := filepath.Glob(filepath.Join(repo, "migrations", "*_"+source.category+"_"+source.object+"_v1.down.sql"))
		if err != nil || len(matches) != 1 {
			t.Fatalf("find %s down migration: matches=%v err=%v", source.object, matches, err)
		}
		raw, err := os.ReadFile(matches[0])
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(raw), source.wantDown) || strings.Contains(string(raw), "Manual restore required") {
			t.Fatalf("invalid generated down for %s:\n%s", source.object, raw)
		}
	}

	second, err := New(cfg).Run("gen")
	if err != nil {
		t.Fatalf("second generation: %v", err)
	}
	if second.Fail.Count != 0 || len(second.NewMigrations.Files) != 0 {
		t.Fatalf("second generation must be idempotent: fail=%d files=%v", second.Fail.Count, second.NewMigrations.Files)
	}
}

func TestBuildDropProgrammabilitySQLRejectsUnsupportedCategory(t *testing.T) {
	t.Parallel()
	if _, err := buildDropProgrammabilitySQL("triggers", "dbo.TriggerA"); err == nil {
		t.Fatal("unsupported programmability category must fail closed")
	}
}
