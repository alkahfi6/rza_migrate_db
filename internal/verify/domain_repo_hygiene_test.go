package verify

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"rza_migrate_db/internal/output"
)

func TestValidateForbiddenSQL_DetectsToken(t *testing.T) {
	root := t.TempDir()
	sqlFile := filepath.Join(root, "db_object", "schema", "tables", "dbo.TableA.sql")
	if err := os.MkdirAll(filepath.Dir(sqlFile), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	if err := os.WriteFile(sqlFile, []byte("USE master;\nCREATE TABLE dbo.TableA(id INT);\n"), 0o644); err != nil {
		t.Fatalf("write sql failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	err := ValidateForbiddenSQL(&report, ForbiddenSQLOptions{
		RepoPath:      func(rel string) string { return filepath.Join(root, rel) },
		DbObjectDir:   "db_object",
		MigrationsDir: "migrations",
		PatchDir:      "patch",
		ToRelPath: func(abs string) string {
			rel, relErr := filepath.Rel(root, abs)
			if relErr != nil {
				return filepath.ToSlash(abs)
			}
			return filepath.ToSlash(rel)
		},
	})
	if err != nil {
		t.Fatalf("validate forbidden sql failed: %v", err)
	}
	if !hasFailRule(report, "FORBIDDEN_SQL_TOKEN") {
		t.Fatalf("expected FORBIDDEN_SQL_TOKEN rule, got %+v", report.Fail.Items)
	}
}

func TestValidateForbiddenSQL_DetectsInvalidEncoding(t *testing.T) {
	root := t.TempDir()
	sqlFile := filepath.Join(root, "db_object", "tables", "dbo.TableA.sql")
	if err := os.MkdirAll(filepath.Dir(sqlFile), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	// CP-1252 style byte 0x96 -> invalid UTF-8 sequence in this context.
	raw := []byte("CREATE TABLE dbo.TableA(id INT); -- bad byte: ")
	raw = append(raw, 0x96, '\n')
	if err := os.WriteFile(sqlFile, raw, 0o644); err != nil {
		t.Fatalf("write sql failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	err := ValidateForbiddenSQL(&report, ForbiddenSQLOptions{
		RepoPath:      func(rel string) string { return filepath.Join(root, rel) },
		DbObjectDir:   "db_object",
		MigrationsDir: "migrations",
		PatchDir:      "patch",
		ToRelPath: func(abs string) string {
			rel, relErr := filepath.Rel(root, abs)
			if relErr != nil {
				return filepath.ToSlash(abs)
			}
			return filepath.ToSlash(rel)
		},
	})
	if err != nil {
		t.Fatalf("validate forbidden sql failed: %v", err)
	}
	if !hasFailRule(report, "INVALID_SQL_ENCODING") {
		t.Fatalf("expected INVALID_SQL_ENCODING rule, got %+v", report.Fail.Items)
	}
}

func TestValidateRepoHygiene_WarnsForMissingOptionalFiles(t *testing.T) {
	root := t.TempDir()
	report := output.NewReport("p", "lf")

	err := ValidateRepoHygiene(&report, RepoHygieneOptions{
		RepoPath: func(rel string) string { return filepath.Join(root, rel) },
	})
	if err != nil {
		t.Fatalf("validate repo hygiene failed: %v", err)
	}

	if report.Fail.Count != 0 {
		t.Fatalf("optional repository files must not fail validation: %+v", report.Fail.Items)
	}
	if got := countPassRule(report, "OPTIONAL_REPO_FILE_MISSING"); got != 3 {
		t.Fatalf("expected three optional-file warnings, got %d: %+v", got, report.Pass.Items)
	}
}

func TestValidateRepoHygiene_WarnsForMissingRecommendedContent(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "constraints_naming.md"), []byte("# constraints\n"), 0o644); err != nil {
		t.Fatalf("write constraints file failed: %v", err)
	}
	if err := os.WriteFile(filepath.Join(root, ".gitignore"), []byte("tmp/**\n"), 0o644); err != nil {
		t.Fatalf("write gitignore failed: %v", err)
	}
	if err := os.WriteFile(filepath.Join(root, ".gitattributes"), []byte("* text=auto eol=lf\n"), 0o644); err != nil {
		t.Fatalf("write gitattributes failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	err := ValidateRepoHygiene(&report, RepoHygieneOptions{
		RepoPath: func(rel string) string { return filepath.Join(root, rel) },
	})
	if err != nil {
		t.Fatalf("validate repo hygiene failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("recommended repository content must not fail validation: %+v", report.Fail.Items)
	}
	if got := countPassRule(report, "OPTIONAL_REPO_FILE_CONTENT"); got != 5 {
		t.Fatalf("expected five optional-content warnings, got %d: %+v", got, report.Pass.Items)
	}
}

func countPassRule(report output.Report, rule string) int {
	count := 0
	for _, item := range report.Pass.Items {
		if item.Rule == rule {
			count++
		}
	}
	return count
}

func TestValidateForbiddenSQL_ParallelMatchesSerial(t *testing.T) {
	root := t.TempDir()

	writeSQL := func(rel, content string) {
		t.Helper()
		abs := filepath.Join(root, rel)
		if err := os.MkdirAll(filepath.Dir(abs), 0o755); err != nil {
			t.Fatalf("mkdir failed: %v", err)
		}
		if err := os.WriteFile(abs, []byte(content), 0o644); err != nil {
			t.Fatalf("write failed: %v", err)
		}
	}

	writeSQL("db_object/tables/dbo.TableA.sql", "CREATE TABLE dbo.TableA(id INT);\n")
	writeSQL("db_object/views/dbo.ViewA.sql", "USE master;\nCREATE OR ALTER VIEW dbo.ViewA AS SELECT 1 AS id;\n")
	writeSQL("migrations/20260217073015000_tables_dbo.TableA_v1.up.sql", "CREATE TABLE dbo.TableA(id INT);\n")
	writeSQL("migrations/20260217073015000_tables_dbo.TableA_v1.down.sql", "DROP TABLE dbo.TableA;\n")
	writeSQL("patch/Fix One/up.sql", "GO\nSELECT 1;\n")
	writeSQL("patch/Fix One/down.sql", "SELECT 1;\n")

	baseOpts := ForbiddenSQLOptions{
		RepoPath:      func(rel string) string { return filepath.Join(root, rel) },
		DbObjectDir:   "db_object",
		MigrationsDir: "migrations",
		PatchDir:      "patch",
		ToRelPath: func(abs string) string {
			rel, relErr := filepath.Rel(root, abs)
			if relErr != nil {
				return filepath.ToSlash(abs)
			}
			return filepath.ToSlash(rel)
		},
		ForbiddenSQLTokens: []string{"GO", "USE"},
	}

	serialReport := output.NewReport("p", "lf")
	serialOpts := baseOpts
	serialOpts.EnableParallelRead = false
	serialOpts.ParallelWorkers = 4
	if err := ValidateForbiddenSQL(&serialReport, serialOpts); err != nil {
		t.Fatalf("serial validate forbidden sql failed: %v", err)
	}

	parallelReport := output.NewReport("p", "lf")
	parallelOpts := baseOpts
	parallelOpts.EnableParallelRead = true
	parallelOpts.ParallelWorkers = 4
	if err := ValidateForbiddenSQL(&parallelReport, parallelOpts); err != nil {
		t.Fatalf("parallel validate forbidden sql failed: %v", err)
	}

	if !reflect.DeepEqual(serialReport.Fail.Items, parallelReport.Fail.Items) {
		t.Fatalf("forbidden fail items mismatch:\nserial=%#v\nparallel=%#v", serialReport.Fail.Items, parallelReport.Fail.Items)
	}
}

func TestForbiddenScanWorkersClamp(t *testing.T) {
	t.Parallel()

	if got := forbiddenScanWorkers(ForbiddenSQLOptions{EnableParallelRead: false, ParallelWorkers: 8}, 10); got != 1 {
		t.Fatalf("expected serial worker=1 when parallel disabled, got %d", got)
	}
	if got := forbiddenScanWorkers(ForbiddenSQLOptions{EnableParallelRead: true, ParallelWorkers: 0}, 10); got != 1 {
		t.Fatalf("expected clamp worker=1 for invalid value, got %d", got)
	}
	want := 3
	if cpu := runtime.GOMAXPROCS(0); cpu < want {
		want = cpu
	}
	if got := forbiddenScanWorkers(ForbiddenSQLOptions{EnableParallelRead: true, ParallelWorkers: 50}, 3); got != want {
		t.Fatalf("expected worker capped by cpu/task count, got=%d want=%d", got, want)
	}
}
