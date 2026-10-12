package verify

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rza_migrate_db/internal/output"
)

func TestValidateTableInlineDefault_PassInlineDefaultNamedConstraint(t *testing.T) {
	root := t.TempDir()
	tablePath := filepath.Join(root, "db_object", "schema", "tables", "dbo.TableA.sql")
	if err := os.MkdirAll(filepath.Dir(tablePath), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	sql := "CREATE TABLE dbo.TableA (\n  id INT NOT NULL,\n  status VARCHAR(10) NOT NULL CONSTRAINT [DF_TableA__status] DEFAULT ('NEW')\n);\n"
	if err := os.WriteFile(tablePath, []byte(sql), 0o644); err != nil {
		t.Fatalf("write table file failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	err := ValidateTableInlineDefault(&report, TableInlineDefaultOptions{
		RepoPath:    func(rel string) string { return filepath.Join(root, rel) },
		DbObjectDir: "db_object",
		ToRelPath: func(abs string) string {
			rel, relErr := filepath.Rel(root, abs)
			if relErr != nil {
				return abs
			}
			return filepath.ToSlash(rel)
		},
	})
	if err != nil {
		t.Fatalf("validate failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected 0 fail, got %d", report.Fail.Count)
	}
}

func TestValidateTableInlineDefault_FailInlineDefaultWithoutConstraintName(t *testing.T) {
	root := t.TempDir()
	tablePath := filepath.Join(root, "db_object", "schema", "tables", "dbo.TableA.sql")
	if err := os.MkdirAll(filepath.Dir(tablePath), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	sql := "CREATE TABLE dbo.TableA (\n  id INT NOT NULL,\n  status VARCHAR(10) NOT NULL DEFAULT ('NEW')\n);\n"
	if err := os.WriteFile(tablePath, []byte(sql), 0o644); err != nil {
		t.Fatalf("write table file failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	err := ValidateTableInlineDefault(&report, TableInlineDefaultOptions{
		RepoPath:    func(rel string) string { return filepath.Join(root, rel) },
		DbObjectDir: "db_object",
		ToRelPath: func(abs string) string {
			rel, relErr := filepath.Rel(root, abs)
			if relErr != nil {
				return abs
			}
			return filepath.ToSlash(rel)
		},
	})
	if err != nil {
		t.Fatalf("validate failed: %v", err)
	}
	if report.Fail.Count == 0 {
		t.Fatalf("expected fail item, got 0")
	}
	var found bool
	for _, item := range report.Fail.Items {
		if item.Rule == "INLINE_DEFAULT_NOT_ALLOWED" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected INLINE_DEFAULT_NOT_ALLOWED, got: %+v", report.Fail.Items)
	}
}

func TestValidateTableInlineDefault_PassNamedTableLevelDefaultForConstraint(t *testing.T) {
	root := t.TempDir()
	tablePath := filepath.Join(root, "db_object", "schema", "tables", "dbo.TableA.sql")
	if err := os.MkdirAll(filepath.Dir(tablePath), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	sql := strings.Join([]string{
		"CREATE TABLE dbo.TableA (",
		"  id INT NOT NULL,",
		"  status VARCHAR(10) NULL,",
		"  CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status",
		");",
		"",
	}, "\n")
	if err := os.WriteFile(tablePath, []byte(sql), 0o644); err != nil {
		t.Fatalf("write table file failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	err := ValidateTableInlineDefault(&report, TableInlineDefaultOptions{
		RepoPath:    func(rel string) string { return filepath.Join(root, rel) },
		DbObjectDir: "db_object",
		ToRelPath:   func(abs string) string { return abs },
	})
	if err != nil {
		t.Fatalf("validate failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("named table-level DEFAULT must be accepted, got: %+v", report.Fail.Items)
	}
}

func TestValidateTableInlineDefault_PassColumnNameContainingDefault(t *testing.T) {
	root := t.TempDir()
	tablePath := filepath.Join(root, "db_object", "schema", "tables", "dbo.TableA.sql")
	if err := os.MkdirAll(filepath.Dir(tablePath), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	sql := strings.Join([]string{
		"CREATE TABLE dbo.TableA (",
		"  id INT NOT NULL,",
		"  PctSellingUpfrontDefault DECIMAL(25,13) NULL,",
		"  DefaultURLProspectusFundFactSheet VARCHAR(200) NULL",
		");",
		"",
	}, "\n")
	if err := os.WriteFile(tablePath, []byte(sql), 0o644); err != nil {
		t.Fatalf("write table file failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	err := ValidateTableInlineDefault(&report, TableInlineDefaultOptions{
		RepoPath:    func(rel string) string { return filepath.Join(root, rel) },
		DbObjectDir: "db_object",
		ToRelPath:   func(abs string) string { return abs },
	})
	if err != nil {
		t.Fatalf("validate failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected no fail, got %d: %+v", report.Fail.Count, report.Fail.Items)
	}
}

func TestValidateTableInlineDefault_FailUnnamedTableLevelConstraints(t *testing.T) {
	root := t.TempDir()
	tablePath := filepath.Join(root, "db_object", "schema", "tables", "dbo.TableA.sql")
	if err := os.MkdirAll(filepath.Dir(tablePath), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	sql := strings.Join([]string{
		"CREATE TABLE dbo.TableA (",
		"  id INT NOT NULL,",
		"  code VARCHAR(20) NULL,",
		"  PRIMARY KEY CLUSTERED (id),",
		"  UNIQUE (code),",
		"  CHECK (id > 0)",
		");",
		"",
	}, "\n")
	if err := os.WriteFile(tablePath, []byte(sql), 0o644); err != nil {
		t.Fatalf("write table file failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	err := ValidateTableInlineDefault(&report, TableInlineDefaultOptions{
		RepoPath:    func(rel string) string { return filepath.Join(root, rel) },
		DbObjectDir: "db_object",
		ToRelPath:   func(abs string) string { return abs },
	})
	if err != nil {
		t.Fatalf("validate failed: %v", err)
	}

	var found bool
	for _, item := range report.Fail.Items {
		if item.Rule == "UNNAMED_CONSTRAINT_NOT_ALLOWED" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected UNNAMED_CONSTRAINT_NOT_ALLOWED fail, got: %+v", report.Fail.Items)
	}
}

func TestValidateTableInlineDefault_FailUnnamedInlineReferenceConstraint(t *testing.T) {
	root := t.TempDir()
	tablePath := filepath.Join(root, "db_object", "schema", "tables", "dbo.TableChild.sql")
	if err := os.MkdirAll(filepath.Dir(tablePath), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	sql := strings.Join([]string{
		"CREATE TABLE dbo.TableChild (",
		"  id INT NOT NULL,",
		"  parent_id INT NULL REFERENCES dbo.TableParent(id),",
		"  CONSTRAINT [PK_TableChild] PRIMARY KEY CLUSTERED (id)",
		");",
		"",
	}, "\n")
	if err := os.WriteFile(tablePath, []byte(sql), 0o644); err != nil {
		t.Fatalf("write table file failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	err := ValidateTableInlineDefault(&report, TableInlineDefaultOptions{
		RepoPath:    func(rel string) string { return filepath.Join(root, rel) },
		DbObjectDir: "db_object",
		ToRelPath:   func(abs string) string { return abs },
	})
	if err != nil {
		t.Fatalf("validate failed: %v", err)
	}

	var found bool
	for _, item := range report.Fail.Items {
		if item.Rule == "UNNAMED_CONSTRAINT_NOT_ALLOWED" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected UNNAMED_CONSTRAINT_NOT_ALLOWED fail, got: %+v", report.Fail.Items)
	}
}

func TestValidateTableInlineDefault_PassNamedConstraintsAllKinds(t *testing.T) {
	root := t.TempDir()
	tablePath := filepath.Join(root, "db_object", "schema", "tables", "dbo.TableA.sql")
	if err := os.MkdirAll(filepath.Dir(tablePath), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	sql := strings.Join([]string{
		"CREATE TABLE dbo.TableA (",
		"  id INT NOT NULL,",
		"  parent_id INT NULL CONSTRAINT [FK_TableA__Parent] REFERENCES dbo.TableParent(id),",
		"  code VARCHAR(20) NULL CONSTRAINT [UQ_TableA__code] UNIQUE,",
		"  status VARCHAR(10) NULL CONSTRAINT [DF_TableA__status] DEFAULT ('NEW'),",
		"  CONSTRAINT [PK_TableA] PRIMARY KEY CLUSTERED (id),",
		"  CONSTRAINT [CK_TableA__code] CHECK (code <> '')",
		");",
		"",
	}, "\n")
	if err := os.WriteFile(tablePath, []byte(sql), 0o644); err != nil {
		t.Fatalf("write table file failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	err := ValidateTableInlineDefault(&report, TableInlineDefaultOptions{
		RepoPath:    func(rel string) string { return filepath.Join(root, rel) },
		DbObjectDir: "db_object",
		ToRelPath:   func(abs string) string { return abs },
	})
	if err != nil {
		t.Fatalf("validate failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected no fail, got %d: %+v", report.Fail.Count, report.Fail.Items)
	}
}

func TestValidateTableInlineDefault_FailsSecondUnnamedInlineConstraint(t *testing.T) {
	root := t.TempDir()
	tablePath := filepath.Join(root, "db_object", "tables", "dbo.T.sql")
	if err := os.MkdirAll(filepath.Dir(tablePath), 0o755); err != nil {
		t.Fatal(err)
	}
	sql := "CREATE TABLE dbo.T (amount INT CONSTRAINT LEGACY_DEFAULT DEFAULT 0 CHECK (amount >= 0));\n"
	if err := os.WriteFile(tablePath, []byte(sql), 0o644); err != nil {
		t.Fatal(err)
	}
	report := output.NewReport("p", "lf")
	if err := ValidateTableInlineDefault(&report, TableInlineDefaultOptions{
		RepoPath: func(rel string) string { return filepath.Join(root, rel) }, DbObjectDir: "db_object",
	}); err != nil {
		t.Fatal(err)
	}
	if !hasFailRule(report, "UNNAMED_CONSTRAINT_NOT_ALLOWED") {
		t.Fatalf("second unnamed constraint must be detected: %+v", report.Fail.Items)
	}
}

func TestValidateTableInlineDefault_PassColumnNamedCheckKeyword(t *testing.T) {
	root := t.TempDir()
	tablePath := filepath.Join(root, "db_object", "schema", "tables", "dbo.TableKeyword.sql")
	if err := os.MkdirAll(filepath.Dir(tablePath), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	sql := strings.Join([]string{
		"CREATE TABLE dbo.TableKeyword (",
		"  [CHECK] BIT NULL,",
		"  [DEFAULT] VARCHAR(10) NULL,",
		"  CONSTRAINT [PK_TableKeyword] PRIMARY KEY CLUSTERED ([CHECK])",
		");",
		"",
	}, "\n")
	if err := os.WriteFile(tablePath, []byte(sql), 0o644); err != nil {
		t.Fatalf("write table file failed: %v", err)
	}

	report := output.NewReport("p", "lf")
	err := ValidateTableInlineDefault(&report, TableInlineDefaultOptions{
		RepoPath:    func(rel string) string { return filepath.Join(root, rel) },
		DbObjectDir: "db_object",
		ToRelPath:   func(abs string) string { return abs },
	})
	if err != nil {
		t.Fatalf("validate failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected no fail, got %d: %+v", report.Fail.Count, report.Fail.Items)
	}
}
