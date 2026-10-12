package verify

import (
	"os"
	"path/filepath"
	"testing"

	"rza_migrate_db/internal/output"
)

func TestValidateMigrationPairing_FailMismatch(t *testing.T) {
	report := output.NewReport("p", "lf")
	metas := []MigrationMeta{
		{
			FileName:  "20260217073015000_tables_dbo.TableA_v1.up.sql",
			Version:   "20260217073015000",
			Category:  "tables",
			Object:    "dbo.TableA",
			Direction: "up",
			VK:        1,
		},
		{
			FileName:  "20260217073015000_tables_dbo.TableB_v1.down.sql",
			Version:   "20260217073015000",
			Category:  "tables",
			Object:    "dbo.TableB",
			Direction: "down",
			VK:        1,
		},
	}

	ValidateMigrationPairing(&report, "migrations", metas)
	if !hasFailRule(report, "MIGRATION_PAIR_MISMATCH") {
		t.Fatalf("expected MIGRATION_PAIR_MISMATCH rule, got %+v", report.Fail.Items)
	}
}

func TestValidateMigrations_FailFlatAndCorePair(t *testing.T) {
	root := t.TempDir()
	migrationsDir := filepath.Join(root, "migrations")
	if err := os.MkdirAll(filepath.Join(migrationsDir, "nested"), 0o755); err != nil {
		t.Fatalf("mkdir nested failed: %v", err)
	}
	if err := os.WriteFile(filepath.Join(migrationsDir, "notes.txt"), []byte("x"), 0o644); err != nil {
		t.Fatalf("write notes failed: %v", err)
	}
	if err := os.WriteFile(filepath.Join(migrationsDir, "20260217073015000_tables_dbo.TableA_v1.up.sql"), []byte("SELECT 1;\n"), 0o644); err != nil {
		t.Fatalf("write sql failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	_, err := ValidateMigrations(&report, MigrationsValidationOptions{
		RepoPath:      func(rel string) string { return filepath.Join(root, rel) },
		MigrationsDir: "migrations",
		ParseMigrationFileName: func(name string) (MigrationMeta, *MigrationParseIssue) {
			return MigrationMeta{
				FileName:  name,
				Version:   "20260217073015000",
				Category:  "tables",
				Object:    "dbo.TableA",
				Direction: "up",
				VK:        1,
			}, nil
		},
		CoreAuditUpName:   "000000__core__mig__audit_init.up.sql",
		CoreAuditDownName: "000000__core__mig__audit_init.down.sql",
	})
	if err != nil {
		t.Fatalf("validate migrations failed: %v", err)
	}

	for _, rule := range []string{
		"MIGRATIONS_NOT_FLAT",
		"NON_SQL_MIGRATION_FILE",
		"MISSING_MIGRATION_PAIR",
		"MISSING_CORE_AUDIT_INIT",
	} {
		if !hasFailRule(report, rule) {
			t.Fatalf("expected %s rule, got %+v", rule, report.Fail.Items)
		}
	}
}

func TestValidateMigrationsRejectsEmptySQLInIntegrityContract(t *testing.T) {
	t.Parallel()
	root := t.TempDir()
	migrationsDir := filepath.Join(root, "migrations")
	if err := os.MkdirAll(migrationsDir, 0o755); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"000000__core__mig__audit_init.up.sql", "000000__core__mig__audit_init.down.sql"} {
		if err := os.WriteFile(filepath.Join(migrationsDir, name), []byte(" \n\t"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	report := output.NewReport("p", "lf")
	_, err := ValidateMigrations(&report, MigrationsValidationOptions{
		RepoPath:      func(rel string) string { return filepath.Join(root, rel) },
		MigrationsDir: "migrations",
		ParseMigrationFileName: func(name string) (MigrationMeta, *MigrationParseIssue) {
			direction := "up"
			if filepath.Base(name) == "000000__core__mig__audit_init.down.sql" {
				direction = "down"
			}
			return MigrationMeta{FileName: name, Version: "000000", Category: "core", Object: "mig", Direction: direction, IsCore: true}, nil
		},
		CoreAuditUpName:   "000000__core__mig__audit_init.up.sql",
		CoreAuditDownName: "000000__core__mig__audit_init.down.sql",
	})
	if err != nil {
		t.Fatal(err)
	}
	if !hasFailRule(report, "EMPTY_MIGRATION_FILE") {
		t.Fatalf("expected EMPTY_MIGRATION_FILE, got %+v", report.Fail.Items)
	}
}

func TestValidateCoreAuditContracts_FailMissingContract(t *testing.T) {
	root := t.TempDir()
	migrationsDir := filepath.Join(root, "migrations")
	if err := os.MkdirAll(migrationsDir, 0o755); err != nil {
		t.Fatalf("mkdir migrations failed: %v", err)
	}
	if err := os.WriteFile(filepath.Join(migrationsDir, "000000__core__mig__audit_init.up.sql"), []byte("SELECT 1;\n"), 0o644); err != nil {
		t.Fatalf("write up failed: %v", err)
	}
	if err := os.WriteFile(filepath.Join(migrationsDir, "000000__core__mig__audit_init.down.sql"), []byte("SELECT 1;\n"), 0o644); err != nil {
		t.Fatalf("write down failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	err := ValidateCoreAuditContracts(&report, CoreAuditContractsOptions{
		RepoPath:          func(rel string) string { return filepath.Join(root, rel) },
		MigrationsDir:     "migrations",
		CoreAuditUpName:   "000000__core__mig__audit_init.up.sql",
		CoreAuditDownName: "000000__core__mig__audit_init.down.sql",
	})
	if err != nil {
		t.Fatalf("validate core audit contracts failed: %v", err)
	}
	if !hasFailRule(report, "CORE_AUDIT_CONTRACT_MISSING") {
		t.Fatalf("expected CORE_AUDIT_CONTRACT_MISSING rule, got %+v", report.Fail.Items)
	}
}

func TestValidateMigrationGenerationOrderAllowsDependencyAwareCategoryOrder(t *testing.T) {
	report := output.NewReport("p", "lf")
	metas := []MigrationMeta{
		{
			FileName:  "20260217073015000_indexes_dbo.TableA_v1.up.sql",
			Version:   "20260217073015000",
			Category:  "indexes",
			Object:    "dbo.TableA",
			Direction: "up",
			VK:        1,
		},
		{
			FileName:  "20260217073015001_tables_dbo.TableA_v1.up.sql",
			Version:   "20260217073015001",
			Category:  "tables",
			Object:    "dbo.TableA",
			Direction: "up",
			VK:        1,
		},
	}

	ValidateMigrationGenerationOrder(&report, "migrations", metas, []string{
		"tables", "indexes", "views", "functions", "procedures", "patch",
	})
	if hasFailRule(report, "MIGRATION_ORDER_INVALID") {
		t.Fatalf("dependency-aware ordering must not be rejected, got %+v", report.Fail.Items)
	}
}

func TestValidateMigrationGenerationOrder_PassOrdered(t *testing.T) {
	report := output.NewReport("p", "lf")
	metas := []MigrationMeta{
		{
			FileName:  "20260217073015000_tables_dbo.TableA_v1.up.sql",
			Version:   "20260217073015000",
			Category:  "tables",
			Object:    "dbo.TableA",
			Direction: "up",
			VK:        1,
		},
		{
			FileName:  "20260217073015001_indexes_dbo.TableA_v1.up.sql",
			Version:   "20260217073015001",
			Category:  "indexes",
			Object:    "dbo.TableA",
			Direction: "up",
			VK:        1,
		},
	}

	ValidateMigrationGenerationOrder(&report, "migrations", metas, []string{
		"tables", "indexes", "views", "functions", "procedures", "patch",
	})
	if hasFailRule(report, "MIGRATION_ORDER_INVALID") {
		t.Fatalf("did not expect MIGRATION_ORDER_INVALID rule, got %+v", report.Fail.Items)
	}
}

func TestValidateMigrationGenerationOrder_PassWhenSequenceGapStartsNewBatch(t *testing.T) {
	report := output.NewReport("p", "lf")
	metas := []MigrationMeta{
		{
			FileName:  "20260217073015004_procedures_dbo.ProcA_v1.up.sql",
			Version:   "20260217073015004",
			Category:  "procedures",
			Object:    "dbo.ProcA",
			Direction: "up",
			VK:        1,
		},
		{
			FileName:  "20260217073015010_tables_dbo.TableB_v1.up.sql",
			Version:   "20260217073015010",
			Category:  "tables",
			Object:    "dbo.TableB",
			Direction: "up",
			VK:        1,
		},
	}

	ValidateMigrationGenerationOrder(&report, "migrations", metas, []string{
		"tables", "indexes", "views", "functions", "procedures", "patch",
	})
	if hasFailRule(report, "MIGRATION_ORDER_INVALID") {
		t.Fatalf("did not expect MIGRATION_ORDER_INVALID rule for sequence-gap split batch, got %+v", report.Fail.Items)
	}
}

func TestValidateMigrationsAllowsDependencyAwareCategoryOrder(t *testing.T) {
	root := t.TempDir()
	migrationsDir := filepath.Join(root, "migrations")
	if err := os.MkdirAll(migrationsDir, 0o755); err != nil {
		t.Fatalf("mkdir migrations failed: %v", err)
	}

	files := []string{
		"000000__core__mig__audit_init.up.sql",
		"000000__core__mig__audit_init.down.sql",
		"20260217073015000_indexes_dbo.TableA_v1.up.sql",
		"20260217073015000_indexes_dbo.TableA_v1.down.sql",
		"20260217073015001_tables_dbo.TableA_v1.up.sql",
		"20260217073015001_tables_dbo.TableA_v1.down.sql",
	}
	for _, name := range files {
		if err := os.WriteFile(filepath.Join(migrationsDir, name), []byte("SELECT 1;\n"), 0o644); err != nil {
			t.Fatalf("write fixture %s failed: %v", name, err)
		}
	}

	report := output.NewReport("p", "lf")
	_, err := ValidateMigrations(&report, MigrationsValidationOptions{
		RepoPath:      func(rel string) string { return filepath.Join(root, rel) },
		MigrationsDir: "migrations",
		GenerationOrder: []string{
			"tables", "indexes", "views", "functions", "procedures", "patch",
		},
		ParseMigrationFileName: func(name string) (MigrationMeta, *MigrationParseIssue) {
			switch name {
			case "000000__core__mig__audit_init.up.sql":
				return MigrationMeta{
					FileName:  name,
					Version:   "000000",
					Category:  "core",
					Direction: "up",
					VK:        0,
					IsCore:    true,
				}, nil
			case "000000__core__mig__audit_init.down.sql":
				return MigrationMeta{
					FileName:  name,
					Version:   "000000",
					Category:  "core",
					Direction: "down",
					VK:        0,
					IsCore:    true,
				}, nil
			case "20260217073015000_indexes_dbo.TableA_v1.up.sql":
				return MigrationMeta{
					FileName:  name,
					Version:   "20260217073015000",
					Category:  "indexes",
					Object:    "dbo.TableA",
					Direction: "up",
					VK:        1,
				}, nil
			case "20260217073015000_indexes_dbo.TableA_v1.down.sql":
				return MigrationMeta{
					FileName:  name,
					Version:   "20260217073015000",
					Category:  "indexes",
					Object:    "dbo.TableA",
					Direction: "down",
					VK:        1,
				}, nil
			case "20260217073015001_tables_dbo.TableA_v1.up.sql":
				return MigrationMeta{
					FileName:  name,
					Version:   "20260217073015001",
					Category:  "tables",
					Object:    "dbo.TableA",
					Direction: "up",
					VK:        1,
				}, nil
			case "20260217073015001_tables_dbo.TableA_v1.down.sql":
				return MigrationMeta{
					FileName:  name,
					Version:   "20260217073015001",
					Category:  "tables",
					Object:    "dbo.TableA",
					Direction: "down",
					VK:        1,
				}, nil
			default:
				return MigrationMeta{}, &MigrationParseIssue{
					Rule:    "INVALID_MIGRATION_FILENAME",
					Message: "unexpected migration fixture",
				}
			}
		},
		CoreAuditUpName:   "000000__core__mig__audit_init.up.sql",
		CoreAuditDownName: "000000__core__mig__audit_init.down.sql",
	})
	if err != nil {
		t.Fatalf("validate migrations failed: %v", err)
	}
	if hasFailRule(report, "MIGRATION_ORDER_INVALID") {
		t.Fatalf("dependency-aware ordering must not be rejected, got %+v", report.Fail.Items)
	}
}

func hasFailRule(report output.Report, rule string) bool {
	for _, item := range report.Fail.Items {
		if item.Rule == rule {
			return true
		}
	}
	return false
}
