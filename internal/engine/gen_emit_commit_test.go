package engine

import (
	"os"
	"path/filepath"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/output"
)

func TestEmitPlannedMigrationsStagingVerifyFailureDoesNotTouchFinalMigrations(t *testing.T) {
	t.Parallel()

	repo := setupMinimalGeneratorRepo(t)
	cfg := config.Default(repo)
	e := New(cfg)

	op := plannedMigration{
		Category:      "views",
		Object:        "dbo.BadView",
		VK:            1,
		UpSemanticSQL: "CREATE VIEW dbo.BadView AS SELECT 1 AS id;\n",
	}
	report := output.NewReport(cfg.OutputNamingPattern, cfg.OutputNamingLineEndings)
	existing := map[string]struct{}{
		coreAuditUpName:   {},
		coreAuditDownName: {},
	}

	if err := e.emitPlannedMigrations(&report, []plannedMigration{op}, existing, ""); err != nil {
		t.Fatalf("emitPlannedMigrations returned error: %v", err)
	}
	if report.Fail.Count == 0 {
		t.Fatalf("expected staged post-emit verify failure")
	}
	if len(report.NewMigrations.Files) != 0 {
		t.Fatalf("expected no committed new migrations, got %v", report.NewMigrations.Files)
	}

	matches, err := filepath.Glob(filepath.Join(repo, "migrations", "*BadView*"))
	if err != nil {
		t.Fatalf("glob failed: %v", err)
	}
	if len(matches) != 0 {
		t.Fatalf("staged verify failure leaked files into final migrations: %v", matches)
	}

	failedCopies, err := filepath.Glob(filepath.Join(repo, "out", "failed_emit", "*BadView*"))
	if err != nil {
		t.Fatalf("glob failed: %v", err)
	}
	if len(failedCopies) == 0 {
		t.Fatalf("expected failed staged files to be copied for inspection")
	}
	if _, err := os.Stat(filepath.Join(repo, "out", "emit_staging")); err != nil {
		t.Fatalf("expected emit_staging parent to exist: %v", err)
	}
}
