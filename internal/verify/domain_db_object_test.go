package verify

import (
	"os"
	"path/filepath"
	"testing"

	"rza_migrate_db/internal/output"
)

func TestValidateDbObjectConsistencyLegacyLayoutWarning(t *testing.T) {
	root := t.TempDir()
	canonical := filepath.Join(root, "db_object", "tables", "dbo.TableA.sql")
	legacy := filepath.Join(root, "db_object", "schema", "tables", "dbo.TableA.sql")
	content := "CREATE TABLE dbo.TableA (\n  id INT NOT NULL\n);\n"
	if err := os.MkdirAll(filepath.Dir(canonical), 0o755); err != nil {
		t.Fatalf("mkdir canonical failed: %v", err)
	}
	if err := os.MkdirAll(filepath.Dir(legacy), 0o755); err != nil {
		t.Fatalf("mkdir legacy failed: %v", err)
	}
	if err := os.WriteFile(canonical, []byte(content), 0o644); err != nil {
		t.Fatalf("write canonical failed: %v", err)
	}
	if err := os.WriteFile(legacy, []byte(content), 0o644); err != nil {
		t.Fatalf("write legacy failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	err := ValidateDbObjectConsistency(&report, DbObjectConsistencyOptions{
		RepoPath:    func(rel string) string { return filepath.Join(root, rel) },
		DbObjectDir: "db_object",
		ToRelPath: func(abs string) string {
			rel, relErr := filepath.Rel(root, abs)
			if relErr != nil {
				return filepath.ToSlash(abs)
			}
			return filepath.ToSlash(rel)
		},
		TrimTrailingSpaces: true,
		EnsureFinalNewline: true,
		ExtractCreateTableStatement: func(objectName, sql string) (string, error) {
			return sql, nil
		},
		ExtractCreateIndexStatements: func(objectName, sql string) ([]string, error) {
			return []string{}, nil
		},
		ExtractProgrammabilityStmt: func(category, objectName, sql string) (string, error) {
			return sql, nil
		},
	})
	if err != nil {
		t.Fatalf("validate db_object consistency failed: %v", err)
	}
	if !hasPassRuleByName(report, "LEGACY_DB_OBJECT_LAYOUT_DETECTED") {
		t.Fatalf("expected LEGACY_DB_OBJECT_LAYOUT_DETECTED warning")
	}
	if !hasPassRuleByName(report, "LEGACY_DB_OBJECT_DUPLICATE_CANONICAL") {
		t.Fatalf("expected LEGACY_DB_OBJECT_DUPLICATE_CANONICAL warning")
	}
	if hasFailRule(report, "DB_OBJECT_DUAL_PATH_CONFLICT") {
		t.Fatalf("did not expect DB_OBJECT_DUAL_PATH_CONFLICT for same content")
	}
}

func TestValidateDbObjectConsistencyLegacyDualPathConflict(t *testing.T) {
	root := t.TempDir()
	canonical := filepath.Join(root, "db_object", "tables", "dbo.TableA.sql")
	legacy := filepath.Join(root, "db_object", "schema", "tables", "dbo.TableA.sql")
	if err := os.MkdirAll(filepath.Dir(canonical), 0o755); err != nil {
		t.Fatalf("mkdir canonical failed: %v", err)
	}
	if err := os.MkdirAll(filepath.Dir(legacy), 0o755); err != nil {
		t.Fatalf("mkdir legacy failed: %v", err)
	}
	if err := os.WriteFile(canonical, []byte("CREATE TABLE dbo.TableA (\n  id INT NOT NULL\n);\n"), 0o644); err != nil {
		t.Fatalf("write canonical failed: %v", err)
	}
	if err := os.WriteFile(legacy, []byte("CREATE TABLE dbo.TableA (\n  id INT NOT NULL,\n  code VARCHAR(10) NULL\n);\n"), 0o644); err != nil {
		t.Fatalf("write legacy failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	err := ValidateDbObjectConsistency(&report, DbObjectConsistencyOptions{
		RepoPath:    func(rel string) string { return filepath.Join(root, rel) },
		DbObjectDir: "db_object",
		ToRelPath: func(abs string) string {
			rel, relErr := filepath.Rel(root, abs)
			if relErr != nil {
				return filepath.ToSlash(abs)
			}
			return filepath.ToSlash(rel)
		},
		TrimTrailingSpaces: true,
		EnsureFinalNewline: true,
		ExtractCreateTableStatement: func(objectName, sql string) (string, error) {
			return sql, nil
		},
		ExtractCreateIndexStatements: func(objectName, sql string) ([]string, error) {
			return []string{}, nil
		},
		ExtractProgrammabilityStmt: func(category, objectName, sql string) (string, error) {
			return sql, nil
		},
	})
	if err != nil {
		t.Fatalf("validate db_object consistency failed: %v", err)
	}
	if !hasFailRule(report, "DB_OBJECT_DUAL_PATH_CONFLICT") {
		t.Fatalf("expected DB_OBJECT_DUAL_PATH_CONFLICT fail")
	}
}

func hasPassRuleByName(report output.Report, rule string) bool {
	for _, item := range report.Pass.Items {
		if item.Rule == rule {
			return true
		}
	}
	return false
}

func TestValidateDbObjectCategoryPlacement(t *testing.T) {
	tests := []struct {
		category string
		sql      string
	}{
		{"tables", "CREATE INDEX IX_A ON dbo.A(id);"},
		{"indexes", "CREATE TABLE dbo.A (id int NOT NULL);"},
		{"procedures", "CREATE OR ALTER FUNCTION dbo.A() RETURNS int AS BEGIN RETURN 1; END;"},
		{"functions", "CREATE OR ALTER VIEW dbo.A AS SELECT 1 AS id;"},
	}
	for _, tt := range tests {
		report := output.NewReport("p", "lf")
		validateDbObjectCategoryPlacement(&report, "db_object/"+tt.category+"/dbo.A.sql", tt.category, tt.sql)
		if !hasFailRule(report, "DB_OBJECT_WRONG_CATEGORY") {
			t.Fatalf("expected wrong-category failure for %s: %+v", tt.category, report.Fail.Items)
		}
	}
}

func TestValidateDbObjectCategoryPlacementIgnoresHeaderWhitespaceAndTempTables(t *testing.T) {
	report := output.NewReport("p", "lf")
	validateDbObjectCategoryPlacement(
		&report,
		"db_object/procedures/dbo.P.sql",
		"procedures",
		"CREATE OR ALTER   PROCEDURE dbo.P AS BEGIN CREATE TABLE #work(id int); END;",
	)
	if hasFailRule(report, "DB_OBJECT_WRONG_CATEGORY") {
		t.Fatalf("procedure with temp table was misclassified: %+v", report.Fail.Items)
	}
}

func TestValidateProgrammabilityHeaders_SourceInvalid(t *testing.T) {
	root := t.TempDir()
	sourceFile := filepath.Join(root, "db_object", "schema", "views", "dbo.ViewA.sql")
	if err := os.MkdirAll(filepath.Dir(sourceFile), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	if err := os.WriteFile(sourceFile, []byte("CREATE VIEW dbo.ViewA AS SELECT 1 AS id;\n"), 0o644); err != nil {
		t.Fatalf("write source file failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	err := ValidateProgrammabilityHeaders(&report, ProgrammabilityHeadersOptions{
		RepoPath:    func(rel string) string { return filepath.Join(root, rel) },
		DbObjectDir: "db_object",
		ToRelPath: func(abs string) string {
			rel, relErr := filepath.Rel(root, abs)
			if relErr != nil {
				return filepath.ToSlash(abs)
			}
			return filepath.ToSlash(rel)
		},
		Metas: nil,
	})
	if err != nil {
		t.Fatalf("validate programmability headers failed: %v", err)
	}
	if !hasFailRule(report, "INVALID_VIEW_HEADER") {
		t.Fatalf("expected INVALID_VIEW_HEADER rule, got %+v", report.Fail.Items)
	}
}

func TestValidateProgrammabilityHeaders_DoesNotValidateVariableCase(t *testing.T) {
	root := t.TempDir()
	sourceFile := filepath.Join(root, "db_object", "procedures", "dbo.P.sql")
	if err := os.MkdirAll(filepath.Dir(sourceFile), 0o755); err != nil {
		t.Fatal(err)
	}
	sql := "CREATE OR ALTER PROCEDURE dbo.P AS BEGIN DECLARE @Tenor VARCHAR(100); DECLARE @TENOR VARCHAR(10); SELECT 1; END;\n"
	if err := os.WriteFile(sourceFile, []byte(sql), 0o644); err != nil {
		t.Fatal(err)
	}
	report := output.NewReport("p", "lf")
	if err := ValidateProgrammabilityHeaders(&report, ProgrammabilityHeadersOptions{
		RepoPath: func(rel string) string { return filepath.Join(root, rel) }, DbObjectDir: "db_object",
		ToRelPath: func(abs string) string { return filepath.ToSlash(abs) },
	}); err != nil {
		t.Fatal(err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("variable casing must be left to target SQL execution: %+v", report.Fail.Items)
	}
}

func TestValidateProgrammabilityHeaders_MigrationInvalid(t *testing.T) {
	root := t.TempDir()
	relMigration := filepath.ToSlash(filepath.Join("migrations", "20260217073015010_views_dbo.ViewA_v1.up.sql"))
	migrationFile := filepath.Join(root, filepath.FromSlash(relMigration))
	if err := os.MkdirAll(filepath.Dir(migrationFile), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	if err := os.WriteFile(migrationFile, []byte("CREATE VIEW dbo.ViewA AS SELECT 1 AS id;\n"), 0o644); err != nil {
		t.Fatalf("write migration file failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	err := ValidateProgrammabilityHeaders(&report, ProgrammabilityHeadersOptions{
		RepoPath:    func(rel string) string { return filepath.Join(root, rel) },
		DbObjectDir: "db_object",
		ToRelPath: func(abs string) string {
			rel, relErr := filepath.Rel(root, abs)
			if relErr != nil {
				return filepath.ToSlash(abs)
			}
			return filepath.ToSlash(rel)
		},
		Metas: []MigrationMeta{
			{
				FileName:  filepath.Base(relMigration),
				RelPath:   relMigration,
				Version:   "20260217073015010",
				Category:  "views",
				Object:    "dbo.ViewA",
				Direction: "up",
				VK:        1,
				IsCore:    false,
			},
		},
	})
	if err != nil {
		t.Fatalf("validate programmability headers failed: %v", err)
	}
	if !hasFailRule(report, "INVALID_PROGRAMMABILITY_MIGRATION_HEADER") {
		t.Fatalf("expected INVALID_PROGRAMMABILITY_MIGRATION_HEADER rule, got %+v", report.Fail.Items)
	}
}

func TestValidateProgrammabilityHeaders_SourceBracketedHeaderValid(t *testing.T) {
	root := t.TempDir()
	sourceFile := filepath.Join(root, "db_object", "schema", "views", "dbo.ViewA.sql")
	if err := os.MkdirAll(filepath.Dir(sourceFile), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	if err := os.WriteFile(sourceFile, []byte("CREATE OR ALTER VIEW [dbo].[ViewA] AS SELECT 1 AS id;\n"), 0o644); err != nil {
		t.Fatalf("write source file failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	err := ValidateProgrammabilityHeaders(&report, ProgrammabilityHeadersOptions{
		RepoPath:    func(rel string) string { return filepath.Join(root, rel) },
		DbObjectDir: "db_object",
		ToRelPath: func(abs string) string {
			rel, relErr := filepath.Rel(root, abs)
			if relErr != nil {
				return filepath.ToSlash(abs)
			}
			return filepath.ToSlash(rel)
		},
		Metas: nil,
	})
	if err != nil {
		t.Fatalf("validate programmability headers failed: %v", err)
	}
	if hasFailRule(report, "INVALID_VIEW_HEADER") {
		t.Fatalf("did not expect INVALID_VIEW_HEADER rule, got %+v", report.Fail.Items)
	}
}

func TestValidateProgrammabilityHeaders_SourceFlexibleWhitespaceValid(t *testing.T) {
	root := t.TempDir()
	sourceFile := filepath.Join(root, "db_object", "schema", "procedures", "dbo.ProcA.sql")
	if err := os.MkdirAll(filepath.Dir(sourceFile), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	if err := os.WriteFile(sourceFile, []byte("CREATE OR ALTER   PROCEDURE [dbo].[ProcA] AS SELECT 1;\n"), 0o644); err != nil {
		t.Fatalf("write source file failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	err := ValidateProgrammabilityHeaders(&report, ProgrammabilityHeadersOptions{
		RepoPath:    func(rel string) string { return filepath.Join(root, rel) },
		DbObjectDir: "db_object",
		ToRelPath: func(abs string) string {
			rel, relErr := filepath.Rel(root, abs)
			if relErr != nil {
				return filepath.ToSlash(abs)
			}
			return filepath.ToSlash(rel)
		},
		Metas: nil,
	})
	if err != nil {
		t.Fatalf("validate programmability headers failed: %v", err)
	}
	if hasFailRule(report, "INVALID_PROCEDURE_HEADER") {
		t.Fatalf("did not expect INVALID_PROCEDURE_HEADER rule, got %+v", report.Fail.Items)
	}
}

func TestValidateProgrammabilityHeaders_SourceContractDisabled(t *testing.T) {
	root := t.TempDir()
	sourceFile := filepath.Join(root, "db_object", "schema", "views", "dbo.ViewA.sql")
	if err := os.MkdirAll(filepath.Dir(sourceFile), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	sql := `CREATE OR ALTER VIEW dbo.ViewA AS SELECT 1 AS id;

CREATE TABLE dbo.BadLeak (
  id INT NOT NULL
);
`
	if err := os.WriteFile(sourceFile, []byte(sql), 0o644); err != nil {
		t.Fatalf("write source file failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	err := ValidateProgrammabilityHeaders(&report, ProgrammabilityHeadersOptions{
		RepoPath:                       func(rel string) string { return filepath.Join(root, rel) },
		DbObjectDir:                    "db_object",
		ToRelPath:                      func(abs string) string { return filepath.ToSlash(abs) },
		Metas:                          nil,
		EnableSourceContractForNewOnly: false,
	})
	if err != nil {
		t.Fatalf("validate programmability headers failed: %v", err)
	}
	if hasFailRule(report, "PROGRAMMABILITY_SOURCE_CONTRACT_VIOLATION") {
		t.Fatalf("did not expect PROGRAMMABILITY_SOURCE_CONTRACT_VIOLATION, got %+v", report.Fail.Items)
	}
}

func TestValidateProgrammabilityHeaders_SourceContractEnabledForNewObject(t *testing.T) {
	root := t.TempDir()
	sourceFile := filepath.Join(root, "db_object", "schema", "views", "dbo.ViewA.sql")
	if err := os.MkdirAll(filepath.Dir(sourceFile), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	sql := `CREATE OR ALTER VIEW dbo.ViewA AS SELECT 1 AS id;

CREATE TABLE dbo.BadLeak (
  id INT NOT NULL
);
`
	if err := os.WriteFile(sourceFile, []byte(sql), 0o644); err != nil {
		t.Fatalf("write source file failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	err := ValidateProgrammabilityHeaders(&report, ProgrammabilityHeadersOptions{
		RepoPath:                       func(rel string) string { return filepath.Join(root, rel) },
		DbObjectDir:                    "db_object",
		ToRelPath:                      func(abs string) string { return filepath.ToSlash(abs) },
		Metas:                          nil,
		EnableSourceContractForNewOnly: true,
	})
	if err != nil {
		t.Fatalf("validate programmability headers failed: %v", err)
	}
	if !hasFailRule(report, "PROGRAMMABILITY_SOURCE_CONTRACT_VIOLATION") {
		t.Fatalf("expected PROGRAMMABILITY_SOURCE_CONTRACT_VIOLATION, got %+v", report.Fail.Items)
	}
}

func TestValidateProgrammabilityHeaders_SourceContractSkipsExistingObject(t *testing.T) {
	root := t.TempDir()
	sourceFile := filepath.Join(root, "db_object", "schema", "views", "dbo.ViewA.sql")
	if err := os.MkdirAll(filepath.Dir(sourceFile), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	sql := `CREATE OR ALTER VIEW dbo.ViewA AS SELECT 1 AS id;

CREATE TABLE dbo.BadLeak (
  id INT NOT NULL
);
`
	if err := os.WriteFile(sourceFile, []byte(sql), 0o644); err != nil {
		t.Fatalf("write source file failed: %v", err)
	}
	relMigration := filepath.ToSlash(filepath.Join("migrations", "20260217073015010_views_dbo.ViewA_v1.up.sql"))
	migrationFile := filepath.Join(root, filepath.FromSlash(relMigration))
	if err := os.MkdirAll(filepath.Dir(migrationFile), 0o755); err != nil {
		t.Fatalf("mkdir migration dir failed: %v", err)
	}
	if err := os.WriteFile(migrationFile, []byte("CREATE OR ALTER VIEW dbo.ViewA AS SELECT 1 AS id;\n"), 0o644); err != nil {
		t.Fatalf("write migration file failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	err := ValidateProgrammabilityHeaders(&report, ProgrammabilityHeadersOptions{
		RepoPath:    func(rel string) string { return filepath.Join(root, rel) },
		DbObjectDir: "db_object",
		ToRelPath:   func(abs string) string { return filepath.ToSlash(abs) },
		Metas: []MigrationMeta{
			{
				FileName:  "20260217073015010_views_dbo.ViewA_v1.up.sql",
				RelPath:   relMigration,
				Version:   "20260217073015010",
				Category:  "views",
				Object:    "dbo.ViewA",
				Direction: "up",
				VK:        1,
				IsCore:    false,
			},
		},
		EnableSourceContractForNewOnly: true,
	})
	if err != nil {
		t.Fatalf("validate programmability headers failed: %v", err)
	}
	if hasFailRule(report, "PROGRAMMABILITY_SOURCE_CONTRACT_VIOLATION") {
		t.Fatalf("did not expect PROGRAMMABILITY_SOURCE_CONTRACT_VIOLATION for existing object, got %+v", report.Fail.Items)
	}
}
