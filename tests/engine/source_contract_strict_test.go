package engine_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/engine"
	"rza_migrate_db/internal/output"
)

func TestVerifyFailsWhenTableSnapshotContainsForbiddenStatement(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "06_ordering_determinism")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	tablePath := filepath.Join(workRoot, "db_object", "schema", "tables", "dbo.TableA.sql")
	raw, err := os.ReadFile(tablePath)
	if err != nil {
		t.Fatalf("read table source failed: %v", err)
	}
	broken := string(raw) + "\nALTER TABLE dbo.TableA ADD bad_col INT NULL;\n"
	if err := os.WriteFile(tablePath, []byte(broken), 0o644); err != nil {
		t.Fatalf("write table source failed: %v", err)
	}

	cfg := config.Default(workRoot)
	cfg.EnableProgrammabilitySourceContractForNewOnly = true
	eng := engine.New(cfg)
	report, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if report.Fail.Count == 0 {
		t.Fatalf("expected fail for forbidden table snapshot statement")
	}
	if !hasFailRuleByName(report.Fail.Items, "TABLE_SNAPSHOT_FORBIDDEN_STATEMENT") {
		t.Fatalf("expected TABLE_SNAPSHOT_FORBIDDEN_STATEMENT, got %+v", report.Fail.Items)
	}
}

func TestVerifyFailsWhenIndexSourceTargetsDifferentTable(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "06_ordering_determinism")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	indexPath := filepath.Join(workRoot, "db_object", "schema", "indexes", "dbo.TableA.sql")
	raw, err := os.ReadFile(indexPath)
	if err != nil {
		t.Fatalf("read index source failed: %v", err)
	}
	broken := strings.TrimSpace(string(raw)) + "\n\nCREATE INDEX IX_TableB__x ON dbo.TableB(id);\n"
	if err := os.WriteFile(indexPath, []byte(broken), 0o644); err != nil {
		t.Fatalf("write index source failed: %v", err)
	}

	cfg := config.Default(workRoot)
	cfg.EnableProgrammabilitySourceContractForNewOnly = true
	eng := engine.New(cfg)
	report, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if report.Fail.Count == 0 {
		t.Fatalf("expected fail for index table mismatch")
	}
	if !hasFailRuleByName(report.Fail.Items, "INDEX_SOURCE_TABLE_MISMATCH") {
		t.Fatalf("expected INDEX_SOURCE_TABLE_MISMATCH, got %+v", report.Fail.Items)
	}
}

func TestVerifyFailsWhenIndexSourceContainsNonIndexStatement(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "06_ordering_determinism")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	indexPath := filepath.Join(workRoot, "db_object", "schema", "indexes", "dbo.TableA.sql")
	broken := `CREATE INDEX IX_TableA__status ON dbo.TableA(status);
ALTER TABLE dbo.TableA ADD bogus INT NULL;
`
	if err := os.WriteFile(indexPath, []byte(broken), 0o644); err != nil {
		t.Fatalf("write index source failed: %v", err)
	}

	cfg := config.Default(workRoot)
	cfg.EnableProgrammabilitySourceContractForNewOnly = true
	eng := engine.New(cfg)
	report, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if report.Fail.Count == 0 {
		t.Fatalf("expected fail for non-index statement in index source")
	}
	if !hasFailRuleByName(report.Fail.Items, "INDEX_SOURCE_CONTRACT_VIOLATION") {
		t.Fatalf("expected INDEX_SOURCE_CONTRACT_VIOLATION, got %+v", report.Fail.Items)
	}
}

func TestVerifyFailsWhenProgrammabilityFileContainsMultipleObjects(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "06_ordering_determinism")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	procPath := filepath.Join(workRoot, "db_object", "schema", "procedures", "dbo.ProcNew.sql")
	broken := `CREATE OR ALTER PROCEDURE dbo.ProcNew
AS
BEGIN
  SELECT 1;
END;

CREATE OR ALTER PROCEDURE dbo.ProcA_Extra
AS
BEGIN
  SELECT 1;
END;
`
	if err := os.WriteFile(procPath, []byte(broken), 0o644); err != nil {
		t.Fatalf("write procedure source failed: %v", err)
	}

	cfg := config.Default(workRoot)
	cfg.EnableProgrammabilitySourceContractForNewOnly = true
	eng := engine.New(cfg)
	report, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if report.Fail.Count == 0 {
		t.Fatalf("expected fail for multi-object programmability file")
	}
	if !hasFailRuleByName(report.Fail.Items, "PROGRAMMABILITY_MULTI_OBJECT") {
		t.Fatalf("expected PROGRAMMABILITY_MULTI_OBJECT, got %+v", report.Fail.Items)
	}
}

func TestVerifyFailsWhenProgrammabilityFileContainsNonProgrammabilityDDL(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "06_ordering_determinism")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	viewPath := filepath.Join(workRoot, "db_object", "schema", "views", "dbo.ViewNew.sql")
	broken := `CREATE OR ALTER VIEW dbo.ViewNew
AS
SELECT id, status, amount
FROM dbo.TableA;

CREATE TABLE dbo.BadLeak (
  id INT NOT NULL
);
`
	if err := os.WriteFile(viewPath, []byte(broken), 0o644); err != nil {
		t.Fatalf("write view source failed: %v", err)
	}

	cfg := config.Default(workRoot)
	cfg.EnableProgrammabilitySourceContractForNewOnly = true
	eng := engine.New(cfg)
	report, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if report.Fail.Count == 0 {
		t.Fatalf("expected fail for non-programmability DDL in programmability file")
	}
	if !hasFailRuleByName(report.Fail.Items, "PROGRAMMABILITY_SOURCE_CONTRACT_VIOLATION") {
		t.Fatalf("expected PROGRAMMABILITY_SOURCE_CONTRACT_VIOLATION, got %+v", report.Fail.Items)
	}
}

func TestVerifyFailsUnnamedTableConstraint(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "06_ordering_determinism")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	tablePath := filepath.Join(workRoot, "db_object", "schema", "tables", "dbo.TableA.sql")
	raw, err := os.ReadFile(tablePath)
	if err != nil {
		t.Fatalf("read table source failed: %v", err)
	}
	broken := strings.Replace(string(raw), "CONSTRAINT PK_TableA PRIMARY KEY (id),", "PRIMARY KEY (id),", 1)
	if err := os.WriteFile(tablePath, []byte(broken), 0o644); err != nil {
		t.Fatalf("write table source failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if !hasFailRuleByName(report.Fail.Items, "UNNAMED_CONSTRAINT_NOT_ALLOWED") {
		t.Fatalf("expected UNNAMED_CONSTRAINT_NOT_ALLOWED, got %+v", report.Fail.Items)
	}
}

func TestVerifyAllowsLegacyTableConstraintName(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "06_ordering_determinism")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	tablePath := filepath.Join(workRoot, "db_object", "schema", "tables", "dbo.TableA.sql")
	raw, err := os.ReadFile(tablePath)
	if err != nil {
		t.Fatalf("read table source failed: %v", err)
	}
	broken := strings.Replace(string(raw), "CONSTRAINT PK_TableA PRIMARY KEY (id),", "CONSTRAINT IDX_TableA PRIMARY KEY (id),", 1)
	if err := os.WriteFile(tablePath, []byte(broken), 0o644); err != nil {
		t.Fatalf("write table source failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if hasFailRuleByName(report.Fail.Items, "CONSTRAINT_NAMING_INVALID") {
		t.Fatalf("did not expect CONSTRAINT_NAMING_INVALID, got %+v", report.Fail.Items)
	}
}

func TestVerifyFailsWhenTableSourceUsesUnsupportedSubset(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "06_ordering_determinism")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	tablePath := filepath.Join(workRoot, "db_object", "schema", "tables", "dbo.TableA.sql")
	raw, err := os.ReadFile(tablePath)
	if err != nil {
		t.Fatalf("read table source failed: %v", err)
	}
	broken := strings.Replace(string(raw), "amount DECIMAL(10,2) NULL,", "amount VARBINARY(MAX) FILESTREAM NULL,", 1)
	if err := os.WriteFile(tablePath, []byte(broken), 0o644); err != nil {
		t.Fatalf("write table source failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if report.Fail.Count == 0 {
		t.Fatalf("expected fail for unsupported table subset")
	}
	if !hasFailRuleByName(report.Fail.Items, "UNSUPPORTED_SQL_FEATURE") {
		t.Fatalf("expected UNSUPPORTED_SQL_FEATURE, got %+v", report.Fail.Items)
	}
}

func TestVerifyPassesWhenIndexSourceUsesFilteredIndex(t *testing.T) {
	t.Parallel()

	repo := repoRoot(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "06_ordering_determinism")
	workRoot := t.TempDir()
	if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
		t.Fatalf("copy fixture input failed: %v", err)
	}

	indexPath := filepath.Join(workRoot, "db_object", "schema", "indexes", "dbo.TableA.sql")
	broken := "CREATE INDEX IX_TableA__status ON dbo.TableA(status) WHERE status IS NOT NULL;\n"
	if err := os.WriteFile(indexPath, []byte(broken), 0o644); err != nil {
		t.Fatalf("write index source failed: %v", err)
	}

	eng := engine.New(config.Default(workRoot))
	report, err := eng.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}
	if report.Fail.Count > 0 {
		t.Fatalf("expected filtered index source to be supported, got fail items: %+v", report.Fail.Items)
	}
}

func hasFailRuleByName(items []output.Item, rule string) bool {
	for _, item := range items {
		if item.Rule == rule {
			return true
		}
	}
	return false
}
