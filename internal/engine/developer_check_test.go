package engine

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"unicode/utf16"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/output"
)

func TestFmtAndCheckAutoConvertUTF16SQLSources(t *testing.T) {
	repoRoot := repoRootFromEnginePackageDir(t)
	source := filepath.Join(repoRoot, "fixtures", "qa_rules", "07_vk_whitespace_only", "input")
	procedure := "CREATE OR ALTER PROCEDURE dbo.Utf16Source AS\nBEGIN\n  SELECT N'東京😀' AS message;\nEND;\n"

	newWorkRoot := func(t *testing.T) (string, string) {
		t.Helper()
		workRoot := t.TempDir()
		if err := copyDirForAutoIndentTest(source, workRoot); err != nil {
			t.Fatalf("copy scenario failed: %v", err)
		}
		path := filepath.Join(workRoot, "db_object", "procedures", "dbo.Utf16Source.sql")
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, utf16LEBOMForTest(procedure), 0o644); err != nil {
			t.Fatal(err)
		}
		return workRoot, path
	}

	t.Run("migrate check applies conversion", func(t *testing.T) {
		workRoot, path := newWorkRoot(t)
		cfg := config.Default(workRoot)
		cfg.EnableOutputNotes = true
		report, err := New(cfg).Run("check")
		if err != nil {
			t.Fatalf("check failed: %v", err)
		}
		if report.Fail.Count != 0 {
			t.Fatalf("check reported failures: %+v", report.Fail.Items)
		}
		if got := report.Notes["fmt_safe_fix_rules"]; got == nil {
			t.Fatal("check did not report safe encoding normalization")
		}
		raw, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		if string(raw) != procedure || len(raw) >= 2 && raw[0] == 0xFF && raw[1] == 0xFE {
			t.Fatalf("source was not written as UTF-8: % X", raw[:min(8, len(raw))])
		}
	})

	t.Run("migrate db fmt applies conversion", func(t *testing.T) {
		workRoot, path := newWorkRoot(t)
		report, err := New(config.Default(workRoot)).Run("fmt")
		if err != nil {
			t.Fatalf("fmt failed: %v", err)
		}
		if report.Fail.Count != 0 {
			t.Fatalf("fmt reported failures: %+v", report.Fail.Items)
		}
		raw, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		if string(raw) != procedure {
			t.Fatalf("fmt output = %q, want UTF-8 SQL %q", raw, procedure)
		}
	})

	t.Run("CI check reports but never writes", func(t *testing.T) {
		workRoot, path := newWorkRoot(t)
		before, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		cfg := config.Default(workRoot)
		cfg.EnableFmtCheck = true
		report, err := New(cfg).Run("check")
		if err != nil {
			t.Fatalf("CI check failed unexpectedly: %v", err)
		}
		if !hasRule(report.Fail.Items, "FMT_CHECK_CHANGES_REQUIRED") || hasRule(report.Fail.Items, "INVALID_SQL_ENCODING") {
			t.Fatalf("CI report must request a local safe fix, got: %+v", report.Fail.Items)
		}
		after, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(after, before) {
			t.Fatal("CI check modified UTF-16 source")
		}
	})

	t.Run("patch source is converted too", func(t *testing.T) {
		root := t.TempDir()
		tablePath := filepath.Join(root, "db_object", "tables", "dbo.TableA.sql")
		if err := os.MkdirAll(filepath.Dir(tablePath), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(tablePath, []byte("CREATE TABLE dbo.TableA (id INT NOT NULL);\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		patchPath := filepath.Join(root, "patch", "001_utf16.up.sql")
		if err := os.MkdirAll(filepath.Dir(patchPath), 0o755); err != nil {
			t.Fatal(err)
		}
		patchSQL := "UPDATE dbo.TableA SET id = id;\n"
		if err := os.WriteFile(patchPath, utf16LEBOMForTest(patchSQL), 0o644); err != nil {
			t.Fatal(err)
		}
		report, err := New(config.Default(root)).Run("fmt")
		if err != nil {
			t.Fatalf("fmt failed: %v", err)
		}
		if report.Fail.Count != 0 {
			t.Fatalf("fmt reported failures: %+v", report.Fail.Items)
		}
		raw, err := os.ReadFile(patchPath)
		if err != nil {
			t.Fatal(err)
		}
		if string(raw) != patchSQL {
			t.Fatalf("patch output = %q, want UTF-8 SQL %q", raw, patchSQL)
		}
	})
}

func utf16LEBOMForTest(text string) []byte {
	words := utf16.Encode([]rune(text))
	raw := make([]byte, 2+len(words)*2)
	raw[0], raw[1] = 0xFF, 0xFE
	for i, word := range words {
		binary.LittleEndian.PutUint16(raw[2+i*2:], word)
	}
	return raw
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestDeveloperCheckFormatsVerifiesAndNeverChangesMigrations(t *testing.T) {
	t.Parallel()

	repoRoot := repoRootFromEnginePackageDir(t)
	source := filepath.Join(repoRoot, "fixtures", "qa_rules", "07_vk_whitespace_only", "input")
	workRoot := t.TempDir()
	if err := copyDirForAutoIndentTest(source, workRoot); err != nil {
		t.Fatalf("copy scenario failed: %v", err)
	}

	tablePath := filepath.Join(workRoot, "db_object", "schema", "tables", "dbo.TableA.sql")
	beforeFormat := `CREATE TABLE dbo.TableA (
id INT NOT NULL
status VARCHAR(10) NULL
amount DECIMAL(10,2) NULL
CONSTRAINT PK_TableA PRIMARY KEY (id)
CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status
);
`
	if err := os.WriteFile(tablePath, []byte(beforeFormat), 0o644); err != nil {
		t.Fatalf("write unformatted table failed: %v", err)
	}
	typeDir := filepath.Join(workRoot, "db_object", "types")
	functionDir := filepath.Join(workRoot, "db_object", "functions")
	if err := os.MkdirAll(typeDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(functionDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(
		filepath.Join(typeDir, "dbo.OrderItemType.sql"),
		[]byte("CREATE TYPE dbo.OrderItemType AS TABLE (\nId INT NOT NULL,\nPRIMARY KEY CLUSTERED (Id)\n);\n"),
		0o644,
	); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(
		filepath.Join(functionDir, "dbo.fnOrderItems.sql"),
		[]byte("CREATE OR ALTER FUNCTION dbo.fnOrderItems(@items dbo.OrderItemType READONLY) RETURNS TABLE AS RETURN (SELECT Id FROM @items);\n"),
		0o644,
	); err != nil {
		t.Fatal(err)
	}

	migrationsBefore := hashSQLTree(t, filepath.Join(workRoot, "migrations"))
	cfg := config.Default(workRoot)
	cfg.EnableOutputNotes = true
	report, err := New(cfg).Run("check")
	if err != nil {
		t.Fatalf("check failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("expected successful check, got failures: %+v", report.Fail.Items)
	}
	if len(report.NewMigrations.Files) != 0 {
		t.Fatalf("checker must report zero generated migrations, got: %#v", report.NewMigrations.Files)
	}
	if got := report.Notes["generation_enabled"]; got != false {
		t.Fatalf("generation_enabled note mismatch: got=%#v want=false", got)
	}
	if got := report.Notes["existing_table_match_mode"]; got != config.ExistingTableMatchBootstrapCompatible {
		t.Fatalf("existing_table_match_mode note mismatch: got=%#v", got)
	}
	if !hasRule(report.Pass.Items, "MIGRATE_CHECK_OK") {
		t.Fatalf("expected MIGRATE_CHECK_OK, got: %+v", report.Pass.Items)
	}
	if !hasRule(report.Pass.Items, "EXISTING_TABLE_MATCH_PROFILE_OK") {
		t.Fatalf("expected EXISTING_TABLE_MATCH_PROFILE_OK, got: %+v", report.Pass.Items)
	}
	if !hasRule(report.Pass.Items, "FUNCTION_SUBTYPE_DETECTED") {
		t.Fatalf("expected table-valued function subtype validation, got: %+v", report.Pass.Items)
	}
	attestation, ok := report.Notes["check_attestation"].(checkAttestation)
	if !ok {
		t.Fatalf("check_attestation type mismatch: %#v", report.Notes["check_attestation"])
	}
	if !attestation.Complete || attestation.Algorithm != "SHA-256" || len(attestation.AggregateSHA256) != sha256.Size*2 {
		t.Fatalf("check attestation is incomplete: %#v", attestation)
	}
	if attestation.SchemaVersion != 2 || attestation.GenerationEnabled || attestation.DatabaseConnection {
		t.Fatalf("unexpected attestation safety contract: %#v", attestation)
	}
	if attestation.FileCount != len(attestation.Files) || attestation.FileCount == 0 {
		t.Fatalf("attestation file count mismatch: count=%d files=%d", attestation.FileCount, len(attestation.Files))
	}
	for _, file := range attestation.Files {
		if file.Path == "" || len(file.SHA256) != sha256.Size*2 {
			t.Fatalf("invalid attestation file: %#v", file)
		}
	}

	afterFormat, err := os.ReadFile(tablePath)
	if err != nil {
		t.Fatalf("read formatted table failed: %v", err)
	}
	if string(afterFormat) == beforeFormat {
		t.Fatal("expected checker to format db_object source")
	}
	if !strings.Contains(string(afterFormat), "  id INT NOT NULL,") {
		t.Fatalf("expected canonical indentation and repaired comma, got:\n%s", afterFormat)
	}

	migrationsAfter := hashSQLTree(t, filepath.Join(workRoot, "migrations"))
	if !reflect.DeepEqual(migrationsAfter, migrationsBefore) {
		t.Fatalf("checker changed migrations tree:\nbefore=%#v\nafter=%#v", migrationsBefore, migrationsAfter)
	}

	outPath := filepath.Join(workRoot, "out", "output.json")
	payload, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("read output.json failed: %v", err)
	}
	var written output.Report
	if err := json.Unmarshal(payload, &written); err != nil {
		t.Fatalf("decode output.json failed: %v", err)
	}
	if written.Fail.Count != 0 || len(written.NewMigrations.Files) != 0 {
		t.Fatalf("unexpected output.json contract: fail=%d migrations=%#v", written.Fail.Count, written.NewMigrations.Files)
	}
	writtenAttestation, ok := written.Notes["check_attestation"].(map[string]interface{})
	if !ok || writtenAttestation["complete"] != true {
		t.Fatalf("written check_attestation is missing or incomplete: %#v", written.Notes["check_attestation"])
	}

	formattedBeforeCI := hashSQLTree(t, filepath.Join(workRoot, "db_object"))
	ciCfg := cfg
	ciCfg.EnableFmtCheck = true
	ciReport, err := New(ciCfg).Run("check")
	if err != nil {
		t.Fatalf("CI-equivalent check failed: %v", err)
	}
	if ciReport.Fail.Count != 0 {
		t.Fatalf("CI must accept the state produced by local checker: %+v", ciReport.Fail.Items)
	}
	if formattedAfterCI := hashSQLTree(t, filepath.Join(workRoot, "db_object")); !reflect.DeepEqual(formattedAfterCI, formattedBeforeCI) {
		t.Fatalf("CI-equivalent check modified formatted db_object")
	}
}

func TestEngineNormalizesValidationProfileValues(t *testing.T) {
	t.Parallel()
	cfg := config.Default(t.TempDir())
	cfg.ValidationProfile = " AUTHORING "
	cfg.MigrationValidation = " INTEGRITY "
	e := New(cfg)
	if e.cfg.ValidationProfile != config.ValidationProfileAuthoring || e.cfg.MigrationValidation != config.MigrationValidationIntegrity {
		t.Fatalf("profile normalization mismatch: %q/%q", e.cfg.ValidationProfile, e.cfg.MigrationValidation)
	}
}

func TestAuthoringProfileDoesNotRevalidateLegacyMigrationSemantics(t *testing.T) {
	t.Parallel()
	repoRoot := repoRootFromEnginePackageDir(t)
	source := filepath.Join(repoRoot, "fixtures", "qa_rules", "07_vk_whitespace_only", "input")
	workRoot := t.TempDir()
	if err := copyDirForAutoIndentTest(source, workRoot); err != nil {
		t.Fatalf("copy scenario failed: %v", err)
	}

	upPath := filepath.Join(workRoot, "migrations", coreAuditUpName)
	upRaw, err := os.ReadFile(upPath)
	if err != nil {
		t.Fatal(err)
	}
	legacyUp := strings.Replace(string(upRaw), "DEFAULT (GETDATE())", "", 1)
	if err := os.WriteFile(upPath, []byte(legacyUp), 0o644); err != nil {
		t.Fatal(err)
	}
	downPath := filepath.Join(workRoot, "migrations", coreAuditDownName)
	downRaw, err := os.ReadFile(downPath)
	if err != nil {
		t.Fatal(err)
	}
	legacyDown := strings.Replace(string(downRaw), "IF OBJECT_ID(N'dbo.mig_event', N'U') IS NOT NULL\n  DROP TABLE dbo.mig_event;\n", "", 1)
	if err := os.WriteFile(downPath, []byte(legacyDown), 0o644); err != nil {
		t.Fatal(err)
	}

	authoring := config.Default(workRoot)
	authoring.ValidationProfile = config.ValidationProfileAuthoring
	authoring.MigrationValidation = config.MigrationValidationIntegrity
	authoring.EnableOutputNotes = true
	authoringReport, err := New(authoring).Run("check")
	if err != nil {
		t.Fatalf("authoring check failed: %v", err)
	}
	if authoringReport.Fail.Count != 0 {
		t.Fatalf("legacy migration semantics must not block authoring: %+v", authoringReport.Fail.Items)
	}
	if got := authoringReport.Notes["migration_validation"]; got != config.MigrationValidationIntegrity {
		t.Fatalf("migration validation note mismatch: %#v", got)
	}

	full := config.Default(workRoot)
	full.EnableFmtCheck = true
	fullReport, err := New(full).Run("check")
	if err != nil {
		t.Fatalf("full check failed: %v", err)
	}
	if !hasRule(fullReport.Fail.Items, "CORE_AUDIT_CONTRACT_MISSING") {
		t.Fatalf("full profile must retain semantic migration checks: %+v", fullReport.Fail.Items)
	}
}

func TestAuthoringIntegrityStillRejectsBrokenMigrationPair(t *testing.T) {
	t.Parallel()
	repoRoot := repoRootFromEnginePackageDir(t)
	source := filepath.Join(repoRoot, "fixtures", "qa_rules", "07_vk_whitespace_only", "input")
	workRoot := t.TempDir()
	if err := copyDirForAutoIndentTest(source, workRoot); err != nil {
		t.Fatalf("copy scenario failed: %v", err)
	}
	for _, entry := range mustReadDir(t, filepath.Join(workRoot, "migrations")) {
		if strings.HasSuffix(entry.Name(), ".down.sql") && !strings.HasPrefix(entry.Name(), "000000__core__") {
			if err := os.Remove(filepath.Join(workRoot, "migrations", entry.Name())); err != nil {
				t.Fatal(err)
			}
			break
		}
	}
	cfg := config.Default(workRoot)
	cfg.ValidationProfile = config.ValidationProfileAuthoring
	cfg.MigrationValidation = config.MigrationValidationIntegrity
	report, err := New(cfg).Run("check")
	if err != nil {
		t.Fatalf("authoring check failed: %v", err)
	}
	if !hasRule(report.Fail.Items, "MISSING_MIGRATION_PAIR") {
		t.Fatalf("integrity mode must reject broken pair: %+v", report.Fail.Items)
	}
}

func TestAuthoringOffAllowsRepositoryWithoutMigrations(t *testing.T) {
	t.Parallel()
	repoRoot := repoRootFromEnginePackageDir(t)
	source := filepath.Join(repoRoot, "fixtures", "qa_rules", "07_vk_whitespace_only", "input")
	workRoot := t.TempDir()
	if err := copyDirForAutoIndentTest(source, workRoot); err != nil {
		t.Fatalf("copy scenario failed: %v", err)
	}
	if err := os.RemoveAll(filepath.Join(workRoot, "migrations")); err != nil {
		t.Fatal(err)
	}
	cfg := config.Default(workRoot)
	cfg.ValidationProfile = config.ValidationProfileAuthoring
	cfg.MigrationValidation = config.MigrationValidationOff
	report, err := New(cfg).Run("check")
	if err != nil {
		t.Fatalf("authoring off failed: %v", err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("migration off must not require migration history: %+v", report.Fail.Items)
	}
}

func mustReadDir(t *testing.T, path string) []os.DirEntry {
	t.Helper()
	entries, err := os.ReadDir(path)
	if err != nil {
		t.Fatal(err)
	}
	return entries
}

func TestDeveloperCheckReadOnlyModeRejectsFormattingDrift(t *testing.T) {
	t.Parallel()

	repoRoot := repoRootFromEnginePackageDir(t)
	source := filepath.Join(repoRoot, "fixtures", "qa_rules", "07_vk_whitespace_only", "input")
	workRoot := t.TempDir()
	if err := copyDirForAutoIndentTest(source, workRoot); err != nil {
		t.Fatalf("copy scenario failed: %v", err)
	}
	tablePath := filepath.Join(workRoot, "db_object", "schema", "tables", "dbo.TableA.sql")
	unformatted := "CREATE TABLE dbo.TableA (\nid INT NOT NULL\n);\n"
	if err := os.WriteFile(tablePath, []byte(unformatted), 0o644); err != nil {
		t.Fatalf("write formatting drift: %v", err)
	}
	migrationsBefore := hashSQLTree(t, filepath.Join(workRoot, "migrations"))

	cfg := config.Default(workRoot)
	cfg.EnableOutputNotes = true
	cfg.EnableFmtCheck = true
	report, err := New(cfg).Run("check")
	if err != nil {
		t.Fatalf("read-only check failed: %v", err)
	}
	if !hasRule(report.Fail.Items, "FMT_CHECK_CHANGES_REQUIRED") {
		t.Fatalf("expected formatting drift failure, got: %+v", report.Fail.Items)
	}
	after, err := os.ReadFile(tablePath)
	if err != nil {
		t.Fatalf("read source after check: %v", err)
	}
	if string(after) != unformatted {
		t.Fatalf("read-only checker modified db_object:\n%s", after)
	}
	if migrationsAfter := hashSQLTree(t, filepath.Join(workRoot, "migrations")); !reflect.DeepEqual(migrationsAfter, migrationsBefore) {
		t.Fatalf("read-only checker changed migrations:\nbefore=%#v\nafter=%#v", migrationsBefore, migrationsAfter)
	}
}

func hashSQLTree(t *testing.T, root string) map[string]string {
	t.Helper()
	hashes := map[string]string{}
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.EqualFold(filepath.Ext(entry.Name()), ".sql") {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		sum := sha256.Sum256(data)
		hashes[filepath.ToSlash(rel)] = hex.EncodeToString(sum[:])
		return nil
	})
	if err != nil {
		t.Fatalf("hash SQL tree failed: %v", err)
	}
	return hashes
}

func hasRule(items []output.Item, rule string) bool {
	for _, item := range items {
		if item.Rule == rule {
			return true
		}
	}
	return false
}
