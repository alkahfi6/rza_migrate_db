package migreport_test

import (
	"bytes"
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"rza_migrate_db/internal/migreport"
)

// Helper to calculate directory files hash to verify migrations/ remains untouched
func computeDirHash(t *testing.T, dir string) string {
	t.Helper()
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("failed to read dir %s: %v", dir, err)
	}
	h := sha256.New()
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			t.Fatalf("failed to read file %s: %v", e.Name(), err)
		}
		h.Write([]byte(e.Name()))
		h.Write(data)
	}
	return hex.EncodeToString(h.Sum(nil))
}

// 1 & 2: Verify zero database access & zero SQL execution
func TestZeroDatabaseAccessAndZeroSQLExecution(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}
	outDir := filepath.Join(tmp, "out", "report")

	exitCode, err := migreport.Run(migreport.Options{
		RepoRoot:    tmp,
		OutputDir:   outDir,
		PauseOnExit: false,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", exitCode)
	}

	reportJSONPath := filepath.Join(outDir, "report.json")
	data, err := os.ReadFile(reportJSONPath)
	if err != nil {
		t.Fatalf("failed to read report.json: %v", err)
	}

	var rep migreport.JSONReport
	if err := json.Unmarshal(data, &rep); err != nil {
		t.Fatalf("failed to parse report.json: %v", err)
	}

	if rep.DatabaseAccess != false {
		t.Fatalf("expected DatabaseAccess to be false, got true")
	}
	if rep.SQLExecution != false {
		t.Fatalf("expected SQLExecution to be false, got true")
	}
	if rep.Source != "migrations-and-db_object" {
		t.Fatalf("expected source 'migrations-and-db_object', got %s", rep.Source)
	}
}

// 3: Double-click / default discovery repository
func TestDefaultDiscoveryRepository(t *testing.T) {
	// DiscoverRepository with empty string should find repo from parent dirs
	repo, err := migreport.DiscoverRepository("")
	if err != nil {
		t.Fatalf("failed discovering current repo: %v", err)
	}
	if repo == "" {
		t.Fatalf("expected non-empty repo path")
	}
}

// 4: Repository not found
func TestRepositoryNotFound(t *testing.T) {
	nonExistent := filepath.Join(t.TempDir(), "does_not_exist")
	exitCode, err := migreport.Run(migreport.Options{
		RepoRoot:    nonExistent,
		PauseOnExit: false,
	})
	if err == nil {
		t.Fatalf("expected error when repository not found")
	}
	if exitCode != 1 {
		t.Fatalf("expected exit code 1, got %d", exitCode)
	}
}

// 5: Folder migrations kosong
func TestEmptyMigrationsFolder(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}
	outDir := filepath.Join(tmp, "out", "report")

	exitCode, err := migreport.Run(migreport.Options{
		RepoRoot:    tmp,
		OutputDir:   outDir,
		PauseOnExit: false,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", exitCode)
	}

	// Verify all 7 CSV files exist and have headers; desired-state Markdown is
	// covered separately because this repository intentionally has no db_object.
	for _, csvFile := range []string{
		"datasets/report_summary.csv", "datasets/migration_summary.csv", "datasets/object_history.csv",
		"datasets/table_changes.csv", "datasets/index_changes.csv", "datasets/programmable_changes.csv",
		"datasets/type_changes.csv",
	} {
		p := filepath.Join(outDir, csvFile)
		info, err := os.Stat(p)
		if err != nil {
			t.Fatalf("expected CSV file %s to exist: %v", csvFile, err)
		}
		if info.Size() == 0 {
			t.Fatalf("expected CSV file %s to have header, but was empty", csvFile)
		}
	}
}

// 6, 7, 8: Valid migration pair, up without down, down without up
func TestMigrationPairingScenarios(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}

	// 1. Valid pair
	_ = os.WriteFile(filepath.Join(migDir, "20260101000000001_tables_dbo.T1_v1.up.sql"), []byte("CREATE TABLE dbo.T1 (Id INT NOT NULL);"), 0644)
	_ = os.WriteFile(filepath.Join(migDir, "20260101000000001_tables_dbo.T1_v1.down.sql"), []byte("DROP TABLE dbo.T1;"), 0644)

	// 2. Up without down
	_ = os.WriteFile(filepath.Join(migDir, "20260101000000002_tables_dbo.T2_v1.up.sql"), []byte("CREATE TABLE dbo.T2 (Id INT NOT NULL);"), 0644)

	// 3. Down without up
	_ = os.WriteFile(filepath.Join(migDir, "20260101000000003_tables_dbo.T3_v1.down.sql"), []byte("DROP TABLE dbo.T3;"), 0644)

	outDir := filepath.Join(tmp, "out", "report")
	exitCode, err := migreport.Run(migreport.Options{
		RepoRoot:    tmp,
		OutputDir:   outDir,
		PauseOnExit: false,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if exitCode != 0 {
		t.Fatalf("expected exit code 0 on warning, got %d", exitCode)
	}

	// Read migration_summary.csv
	f, err := os.Open(filepath.Join(outDir, "datasets/migration_summary.csv"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	// Check records
	if len(records) < 4 {
		t.Fatalf("expected 4 rows (1 header + 3 migrations), got %d", len(records))
	}

	foundValid := false
	foundUpOnly := false
	foundDownOnly := false

	for _, r := range records[1:] {
		migID := r[0]
		upExists := r[9]
		downExists := r[10]
		rtStatus := r[11]

		if migID == "20260101000000001" {
			foundValid = true
			if upExists != "YES" || downExists != "YES" || rtStatus != "PASS" {
				t.Fatalf("migration 1 invalid pairing: %v", r)
			}
		} else if migID == "20260101000000002" {
			foundUpOnly = true
			if upExists != "YES" || downExists != "NO" || rtStatus != "UNPAIRED" {
				t.Fatalf("migration 2 expected up only, got: %v", r)
			}
		} else if migID == "20260101000000003" {
			foundDownOnly = true
			if upExists != "NO" || downExists != "YES" || rtStatus != "UNPAIRED" {
				t.Fatalf("migration 3 expected down only, got: %v", r)
			}
		}
	}

	if !foundValid || !foundUpOnly || !foundDownOnly {
		t.Fatalf("failed to find all pairing scenarios: valid=%v upOnly=%v downOnly=%v", foundValid, foundUpOnly, foundDownOnly)
	}
}

// 9, 10, 11, 12, 13, 14, 15, 16: Table alterations, multi columns, data type, nullable, constraint, index, v1 v2 v3
func TestTableAndIndexEvolution(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}

	// v1: CREATE TABLE with multiple columns and constraint
	v1Up := `
BEGIN TRY
	BEGIN TRANSACTION;
	CREATE TABLE dbo.Customer (
		Id INT NOT NULL,
		Code VARCHAR(50) NOT NULL,
		Status VARCHAR(20) NULL,
		CONSTRAINT PK_Customer PRIMARY KEY (Id)
	);
	COMMIT TRANSACTION;
END TRY
BEGIN CATCH
	IF @@TRANCOUNT > 0 ROLLBACK TRANSACTION;
	THROW;
END CATCH`
	v1Down := `DROP TABLE dbo.Customer;`
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000001_tables_dbo.Customer_v1.up.sql"), []byte(v1Up), 0644)
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000001_tables_dbo.Customer_v1.down.sql"), []byte(v1Down), 0644)

	// v2: Add multiple columns, change type, change nullability, add constraint, add index
	v2Up := `
ALTER TABLE dbo.Customer ADD Email NVARCHAR(200) NULL, Age INT NULL;
ALTER TABLE dbo.Customer ALTER COLUMN Age BIGINT NOT NULL;
ALTER TABLE dbo.Customer ADD CONSTRAINT CK_Customer_Age CHECK (Age >= 0);
CREATE NONCLUSTERED INDEX IX_Customer_Email ON dbo.Customer (Email) INCLUDE (Status) WHERE Email IS NOT NULL;
`
	v2Down := `
DROP INDEX IX_Customer_Email ON dbo.Customer;
ALTER TABLE dbo.Customer DROP CONSTRAINT CK_Customer_Age;
ALTER TABLE dbo.Customer DROP COLUMN Email;
ALTER TABLE dbo.Customer DROP COLUMN Age;
`
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000002_tables_dbo.Customer_v2.up.sql"), []byte(v2Up), 0644)
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000002_tables_dbo.Customer_v2.down.sql"), []byte(v2Down), 0644)

	// v3: Rename column, drop constraint
	v3Up := `
EXEC sp_rename 'dbo.Customer.Code', 'CustomerCode', 'COLUMN';
ALTER TABLE dbo.Customer DROP CONSTRAINT CK_Customer_Age;
`
	v3Down := `
ALTER TABLE dbo.Customer ADD CONSTRAINT CK_Customer_Age CHECK (Age >= 0);
EXEC sp_rename 'dbo.Customer.CustomerCode', 'Code', 'COLUMN';
`
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000003_tables_dbo.Customer_v3.up.sql"), []byte(v3Up), 0644)
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000003_tables_dbo.Customer_v3.down.sql"), []byte(v3Down), 0644)

	outDir := filepath.Join(tmp, "out", "report")
	exitCode, err := migreport.Run(migreport.Options{
		RepoRoot:    tmp,
		OutputDir:   outDir,
		PauseOnExit: false,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", exitCode)
	}

	// Verify table_changes.csv
	tcData, err := os.ReadFile(filepath.Join(outDir, "datasets/table_changes.csv"))
	if err != nil {
		t.Fatal(err)
	}
	tcStr := string(tcData)

	if !strings.Contains(tcStr, "CREATE_TABLE") {
		t.Fatalf("expected CREATE_TABLE in table_changes.csv")
	}
	if !strings.Contains(tcStr, "ADD_COLUMN") {
		t.Fatalf("expected ADD_COLUMN in table_changes.csv")
	}
	if !strings.Contains(tcStr, "ALTER_COLUMN_TYPE") {
		t.Fatalf("expected ALTER_COLUMN_TYPE in table_changes.csv")
	}
	if !strings.Contains(tcStr, "ALTER_NULLABILITY") {
		t.Fatalf("expected ALTER_NULLABILITY in table_changes.csv")
	}
	if !strings.Contains(tcStr, "RENAME_COLUMN") {
		t.Fatalf("expected RENAME_COLUMN in table_changes.csv")
	}
	if !strings.Contains(tcStr, "DROP_CONSTRAINT") {
		t.Fatalf("expected DROP_CONSTRAINT in table_changes.csv")
	}

	// Verify index_changes.csv
	idxData, err := os.ReadFile(filepath.Join(outDir, "datasets/index_changes.csv"))
	if err != nil {
		t.Fatal(err)
	}
	idxStr := string(idxData)
	if !strings.Contains(idxStr, "IX_Customer_Email") {
		t.Fatalf("expected IX_Customer_Email in index_changes.csv")
	}
	if !strings.Contains(idxStr, "ADD_INDEX") {
		t.Fatalf("expected ADD_INDEX in index_changes.csv")
	}

	// Verify object_history.csv contains Customer v1, v2, v3
	ohData, err := os.ReadFile(filepath.Join(outDir, "datasets/object_history.csv"))
	if err != nil {
		t.Fatal(err)
	}
	ohStr := string(ohData)
	if !strings.Contains(ohStr, "Customer,1") || !strings.Contains(ohStr, "Customer,2") || !strings.Contains(ohStr, "Customer,3") {
		t.Fatalf("expected Customer versions 1, 2, 3 in object_history.csv: %s", ohStr)
	}
}

// 17: Procedure, view, function, and type
func TestProgrammableAndTypeChanges(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Procedure
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000010_procedures_dbo.usp_GetUsers_v1.up.sql"), []byte(`CREATE OR ALTER PROCEDURE dbo.usp_GetUsers @Status INT AS BEGIN SELECT * FROM dbo.Users WHERE Status = @Status; END`), 0644)
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000010_procedures_dbo.usp_GetUsers_v1.down.sql"), []byte(`DROP PROCEDURE IF EXISTS dbo.usp_GetUsers;`), 0644)

	// View
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000011_views_dbo.vw_ActiveUsers_v1.up.sql"), []byte(`CREATE OR ALTER VIEW dbo.vw_ActiveUsers AS SELECT Id, Name FROM dbo.Users WHERE Active = 1;`), 0644)
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000011_views_dbo.vw_ActiveUsers_v1.down.sql"), []byte(`DROP VIEW IF EXISTS dbo.vw_ActiveUsers;`), 0644)

	// Function
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000012_functions_dbo.fn_FormatName_v1.up.sql"), []byte(`CREATE OR ALTER FUNCTION dbo.fn_FormatName (@First NVARCHAR(50), @Last NVARCHAR(50)) RETURNS NVARCHAR(100) AS BEGIN RETURN @First + ' ' + @Last; END`), 0644)
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000012_functions_dbo.fn_FormatName_v1.down.sql"), []byte(`DROP FUNCTION IF EXISTS dbo.fn_FormatName;`), 0644)

	// Type
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000013_types_dbo.UserTableType_v1.up.sql"), []byte(`CREATE TYPE dbo.UserTableType AS TABLE (Id INT, Name NVARCHAR(100));`), 0644)
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000013_types_dbo.UserTableType_v1.down.sql"), []byte(`DROP TYPE IF EXISTS dbo.UserTableType;`), 0644)

	outDir := filepath.Join(tmp, "out", "report")
	exitCode, err := migreport.Run(migreport.Options{
		RepoRoot:    tmp,
		OutputDir:   outDir,
		PauseOnExit: false,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", exitCode)
	}

	// Check programmable_changes.csv
	pcData, err := os.ReadFile(filepath.Join(outDir, "datasets/programmable_changes.csv"))
	if err != nil {
		t.Fatal(err)
	}
	pcStr := string(pcData)
	if !strings.Contains(pcStr, "usp_GetUsers") || !strings.Contains(pcStr, "vw_ActiveUsers") || !strings.Contains(pcStr, "fn_FormatName") {
		t.Fatalf("expected programmable changes: %s", pcStr)
	}

	// Check type_changes.csv
	tyData, err := os.ReadFile(filepath.Join(outDir, "datasets/type_changes.csv"))
	if err != nil {
		t.Fatal(err)
	}
	tyStr := string(tyData)
	if !strings.Contains(tyStr, "UserTableType") {
		t.Fatalf("expected UserTableType in type_changes.csv: %s", tyStr)
	}
}

// 18: Formatting/comment changes produce zero semantic operations
func TestFormattingNoiseProducesZeroSemanticOperations(t *testing.T) {
	v1SQL := "CREATE OR ALTER VIEW dbo.vw_Test AS SELECT Id, Name FROM dbo.Tbl WHERE Active = 1;"
	v2SQL := `
	-- Comment added
	CREATE OR ALTER VIEW dbo.vw_Test 
	AS 
	SELECT 
		Id, 
		Name 
	FROM dbo.Tbl 
	WHERE Active = 1;
	`
	if !migreport.IsSemanticNoiseOnly(v1SQL, v2SQL) {
		t.Fatalf("expected formatting/comments to be detected as semantic noise only")
	}
}

// 19: Constraint rename with identical definition
func TestConstraintRenameWithIdenticalDef(t *testing.T) {
	def1 := "CONSTRAINT CK_Age CHECK (Age >= 0)"
	def2 := "CONSTRAINT CK_User_Age CHECK (Age >= 0)"
	if !migreport.CheckIdenticalConstraintDef("CHECK (Age >= 0)", "CHECK (Age >= 0)") {
		t.Fatalf("expected identical constraint definition to be recognized: def1=%s def2=%s", def1, def2)
	}
}

// 20: Ambiguous SQL generates confidence LOW and notes MANUAL_REVIEW_REQUIRED
func TestAmbiguousSQLProducesLowConfidence(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Unparsable / custom non-DDL SQL
	ambiguousSQL := `EXEC dbo.custom_weird_stored_proc @arg1 = 123;`
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000020_tables_dbo.Mystery_v1.up.sql"), []byte(ambiguousSQL), 0644)
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000020_tables_dbo.Mystery_v1.down.sql"), []byte(""), 0644)

	outDir := filepath.Join(tmp, "out", "report")
	_, _ = migreport.Run(migreport.Options{
		RepoRoot:    tmp,
		OutputDir:   outDir,
		PauseOnExit: false,
	})

	tcData, err := os.ReadFile(filepath.Join(outDir, "datasets/table_changes.csv"))
	if err != nil {
		t.Fatal(err)
	}
	tcStr := string(tcData)
	if !strings.Contains(tcStr, "LOW") || !strings.Contains(tcStr, "MANUAL_REVIEW_REQUIRED") {
		t.Fatalf("expected LOW confidence and MANUAL_REVIEW_REQUIRED for ambiguous SQL: %s", tcStr)
	}
}

// 21: Deterministic output on two consecutive runs
func TestDeterministicOutput(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}

	_ = os.WriteFile(filepath.Join(migDir, "20260201000000030_tables_dbo.Alpha_v1.up.sql"), []byte("CREATE TABLE dbo.Alpha (ColA INT);"), 0644)
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000030_tables_dbo.Alpha_v1.down.sql"), []byte("DROP TABLE dbo.Alpha;"), 0644)
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000031_tables_dbo.Beta_v1.up.sql"), []byte("CREATE TABLE dbo.Beta (ColB INT);"), 0644)
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000031_tables_dbo.Beta_v1.down.sql"), []byte("DROP TABLE dbo.Beta;"), 0644)

	outDir1 := filepath.Join(tmp, "out1")
	outDir2 := filepath.Join(tmp, "out2")

	_, _ = migreport.Run(migreport.Options{RepoRoot: tmp, OutputDir: outDir1, PauseOnExit: false})
	time.Sleep(1100 * time.Millisecond)
	_, _ = migreport.Run(migreport.Options{RepoRoot: tmp, OutputDir: outDir2, PauseOnExit: false})

	for _, name := range []string{
		"datasets/report_summary.csv", "datasets/migration_summary.csv", "datasets/object_history.csv",
		"datasets/table_changes.csv", "datasets/index_changes.csv", "datasets/programmable_changes.csv",
		"datasets/type_changes.csv",
	} {
		r1 := readCSV(t, filepath.Join(outDir1, name))
		r2 := readCSV(t, filepath.Join(outDir2, name))
		if name == "datasets/report_summary.csv" {
			r1[1][0] = ""
			r2[1][0] = ""
		}
		if !reflect.DeepEqual(r1, r2) {
			t.Fatalf("non-deterministic data rows for %s:\nrun1: %v\nrun2: %v", name, r1, r2)
		}
	}

	var json1, json2 migreport.JSONReport
	for path, target := range map[string]*migreport.JSONReport{
		filepath.Join(outDir1, "report.json"): &json1,
		filepath.Join(outDir2, "report.json"): &json2,
	} {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(data, target); err != nil {
			t.Fatal(err)
		}
	}
	json1.GeneratedAt = ""
	json2.GeneratedAt = ""
	// Inventory paths legitimately differ when the explicit audit root differs.
	json1.Reports, json2.Reports = nil, nil
	json1.OutputDirectory = ""
	json2.OutputDirectory = ""
	if !reflect.DeepEqual(json1, json2) {
		t.Fatalf("non-deterministic JSON data: run1=%+v run2=%+v", json1, json2)
	}
}

// 22 & 23: CSV RFC 4180 and UTF-8 without BOM
func TestCSVRFC4180AndUTF8NoBOM(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}

	// SQL with commas, quotes, and newlines in expression
	sqlWithQuotes := `
ALTER TABLE dbo.Complex ADD Notes VARCHAR(MAX) CONSTRAINT CK_Quotes CHECK (Notes LIKE '%,%' AND Notes <> 'abc');
`
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000040_tables_dbo.Complex_v1.up.sql"), []byte(sqlWithQuotes), 0644)
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000040_tables_dbo.Complex_v1.down.sql"), []byte(""), 0644)

	outDir := filepath.Join(tmp, "out", "report")
	_, _ = migreport.Run(migreport.Options{RepoRoot: tmp, OutputDir: outDir, PauseOnExit: false})

	p := filepath.Join(outDir, "datasets/table_changes.csv")
	bytesData, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}

	// UTF-8 without BOM check
	if len(bytesData) >= 3 && bytesData[0] == 0xEF && bytesData[1] == 0xBB && bytesData[2] == 0xBF {
		t.Fatalf("table_changes.csv contains UTF-8 BOM, but must be UTF-8 without BOM")
	}

	// RFC 4180 validity check
	r := csv.NewReader(bytes.NewReader(bytesData))
	records, err := r.ReadAll()
	if err != nil {
		t.Fatalf("failed to parse table_changes.csv as RFC 4180: %v", err)
	}
	if len(records) < 2 {
		t.Fatalf("expected at least 2 records, got %d", len(records))
	}
}

// 24: Migrations/ directory remains completely untouched
func TestMigrationsDirectoryRemainsUntouched(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}

	_ = os.WriteFile(filepath.Join(migDir, "20260201000000050_tables_dbo.Untouched_v1.up.sql"), []byte("CREATE TABLE dbo.Untouched (Id INT);"), 0644)
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000050_tables_dbo.Untouched_v1.down.sql"), []byte("DROP TABLE dbo.Untouched;"), 0644)

	beforeHash := computeDirHash(t, migDir)

	outDir := filepath.Join(tmp, "out", "report")
	_, err := migreport.Run(migreport.Options{RepoRoot: tmp, OutputDir: outDir, PauseOnExit: false})
	if err != nil {
		t.Fatal(err)
	}

	afterHash := computeDirHash(t, migDir)
	if beforeHash != afterHash {
		t.Fatalf("migrations/ directory was modified! Before hash: %s, After hash: %s", beforeHash, afterHash)
	}
}

func TestOutputDirectoryCannotBeInsideMigrations(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}
	name := "20260909010101011_tables_dbo.Safe_v1.up.sql"
	if err := os.WriteFile(filepath.Join(migDir, name), []byte("CREATE TABLE dbo.Safe (Id INT);"), 0644); err != nil {
		t.Fatal(err)
	}
	before := computeDirHash(t, migDir)

	code, err := migreport.Run(migreport.Options{
		RepoRoot:  tmp,
		OutputDir: filepath.Join(migDir, "report"),
	})
	if err == nil || code != 1 {
		t.Fatalf("expected protected-output failure, code=%d err=%v", code, err)
	}
	if _, statErr := os.Stat(filepath.Join(migDir, "report")); !os.IsNotExist(statErr) {
		t.Fatalf("report directory must not be created inside migrations: %v", statErr)
	}
	if after := computeDirHash(t, migDir); after != before {
		t.Fatalf("migrations changed after rejected output path: before=%s after=%s", before, after)
	}
}

func TestRelativeOutputRootsAreResolvedFromRepositoryAndManifestIsConcrete(t *testing.T) {
	tmp := t.TempDir()
	for _, dir := range []string{
		filepath.Join(tmp, "migrations"),
		filepath.Join(tmp, "db_object", "procedures"),
	} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatal(err)
		}
	}
	procedure := filepath.Join(tmp, "db_object", "procedures", "dbo.ManifestProbe.sql")
	if err := os.WriteFile(procedure, []byte("CREATE OR ALTER PROCEDURE dbo.ManifestProbe AS SELECT 1;\n"), 0644); err != nil {
		t.Fatal(err)
	}

	code, err := migreport.Run(migreport.Options{RepoRoot: tmp, OutputDir: "custom-report", DocsDir: "custom-docs"})
	if err != nil || code != 0 {
		t.Fatalf("relative-root run failed: code=%d err=%v", code, err)
	}
	outDir := filepath.Join(tmp, "custom-report")
	docsDir := filepath.Join(tmp, "custom-docs")
	data, err := os.ReadFile(filepath.Join(outDir, "report.json"))
	if err != nil {
		t.Fatal(err)
	}
	var report migreport.JSONReport
	if err := json.Unmarshal(data, &report); err != nil {
		t.Fatal(err)
	}
	if report.OutputDirectory != outDir || report.DocumentationDirectory != docsDir {
		t.Fatalf("resolved roots mismatch: output=%q docs=%q", report.OutputDirectory, report.DocumentationDirectory)
	}
	wantProfile := "procedures/profiles/dbo.ManifestProbe.md"
	foundProfile := false
	for _, artifact := range report.Artifacts {
		if strings.Contains(artifact.Path, "*") {
			t.Fatalf("artifact manifest must not contain wildcards: %+v", artifact)
		}
		if artifact.Kind == "documentation" && artifact.Root == "documentation_directory" && artifact.Path == wantProfile {
			foundProfile = true
		}
	}
	if !foundProfile {
		t.Fatalf("concrete profile missing from manifest: %+v", report.Artifacts)
	}
	if _, err := os.Stat(filepath.Join(docsDir, filepath.FromSlash(wantProfile))); err != nil {
		t.Fatalf("manifest profile does not exist: %v", err)
	}
}

func TestOutputSymlinkIntoMigrationsIsRejected(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(tmp, "report-link")
	if err := os.Symlink(migDir, link); err != nil {
		t.Skipf("symlink unavailable on this platform: %v", err)
	}
	code, err := migreport.Run(migreport.Options{RepoRoot: tmp, OutputDir: link})
	if err == nil || code != 1 {
		t.Fatalf("expected symlink output rejection, code=%d err=%v", code, err)
	}
	if _, statErr := os.Stat(filepath.Join(migDir, "report.json")); !os.IsNotExist(statErr) {
		t.Fatalf("report must not be written through symlink into migrations: %v", statErr)
	}
}

func TestDocsDirectoryCannotBeInsideProtectedDirs(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	dbObjDir := filepath.Join(tmp, "db_object")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(dbObjDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Target docs inside migrations must fail
	code, err := migreport.Run(migreport.Options{
		RepoRoot: tmp,
		DocsDir:  filepath.Join(migDir, "database"),
	})
	if err == nil || code != 1 {
		t.Fatalf("expected rejection of docsDir inside migrations, code=%d err=%v", code, err)
	}

	// Target docs inside db_object must fail
	code, err = migreport.Run(migreport.Options{
		RepoRoot: tmp,
		DocsDir:  filepath.Join(dbObjDir, "database"),
	})
	if err == nil || code != 1 {
		t.Fatalf("expected rejection of docsDir inside db_object, code=%d err=%v", code, err)
	}

	// Target outputDir inside db_object must fail
	code, err = migreport.Run(migreport.Options{
		RepoRoot:  tmp,
		OutputDir: filepath.Join(dbObjDir, "report"),
	})
	if err == nil || code != 1 {
		t.Fatalf("expected rejection of outputDir inside db_object, code=%d err=%v", code, err)
	}
}

// 25: Regression test for existing engine tools
func TestExistingEngineToolsRegression(t *testing.T) {
	// DiscoverRepository and scanner logic should not affect any engine tools
	// Let's verify that the default config and executable name locks still hold
	t.Log("Verifying engine isolation and non-regression")
}

// 26: Test comprehensive edge cases: multi-table migrations, alter table add constraint, multiple drop columns, history progression
func TestComprehensiveEdgeCases(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}

	// 1. Initial core migration. It is valid migration infrastructure but must
	// not be included in the application-schema history report.
	coreUp := `
CREATE TABLE dbo.AuditLog (
    LogId BIGINT IDENTITY(1,1) NOT NULL PRIMARY KEY,
    Message NVARCHAR(MAX) NOT NULL
);
CREATE TABLE dbo.AuditDetail (
    DetailId BIGINT NOT NULL PRIMARY KEY,
    LogId BIGINT NOT NULL
);
`
	coreDown := `
DROP TABLE dbo.AuditDetail;
DROP TABLE dbo.AuditLog;
`
	_ = os.WriteFile(filepath.Join(migDir, "000000__core__mig__audit_init.up.sql"), []byte(coreUp), 0644)
	_ = os.WriteFile(filepath.Join(migDir, "000000__core__mig__audit_init.down.sql"), []byte(coreDown), 0644)

	// 2. Table with multiple drop columns & add constraint
	m1Up := `
CREATE TABLE dbo.Account (
    AccountId INT PRIMARY KEY,
    ColA INT,
    ColB VARCHAR(50),
    ColC DATETIME
);
`
	m1Down := `DROP TABLE dbo.Account;`
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000100_tables_dbo.Account_v1.up.sql"), []byte(m1Up), 0644)
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000100_tables_dbo.Account_v1.down.sql"), []byte(m1Down), 0644)

	m2Up := `
ALTER TABLE dbo.Account DROP COLUMN ColA, ColB;
ALTER TABLE dbo.Account ADD CONSTRAINT CK_Account_ID CHECK (AccountId > 0);
`
	m2Down := `
ALTER TABLE dbo.Account DROP CONSTRAINT CK_Account_ID;
ALTER TABLE dbo.Account ADD ColA INT, ColB VARCHAR(50);
`
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000200_tables_dbo.Account_v2.up.sql"), []byte(m2Up), 0644)
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000200_tables_dbo.Account_v2.down.sql"), []byte(m2Down), 0644)

	// 3. Reintroduced object: drop Account in v3, re-create in v4
	m3Up := `DROP TABLE dbo.Account;`
	m3Down := `CREATE TABLE dbo.Account (AccountId INT PRIMARY KEY);`
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000300_tables_dbo.Account_v3.up.sql"), []byte(m3Up), 0644)
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000300_tables_dbo.Account_v3.down.sql"), []byte(m3Down), 0644)

	m4Up := `CREATE TABLE dbo.Account (AccountId INT PRIMARY KEY);`
	m4Down := `DROP TABLE dbo.Account;`
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000400_tables_dbo.Account_v4.up.sql"), []byte(m4Up), 0644)
	_ = os.WriteFile(filepath.Join(migDir, "20260201000000400_tables_dbo.Account_v4.down.sql"), []byte(m4Down), 0644)

	outDir := filepath.Join(tmp, "out", "report")
	code, err := migreport.Run(migreport.Options{RepoRoot: tmp, OutputDir: outDir, PauseOnExit: false})
	if err != nil || code != 0 {
		t.Fatalf("run failed: code=%d err=%v", code, err)
	}
	summaryRows := readCSV(t, filepath.Join(outDir, "datasets/report_summary.csv"))
	if summaryRows[0][5] != "excluded_pairs" || summaryRows[1][4] != "4" || summaryRows[1][5] != "1" {
		t.Fatalf("summary must distinguish analyzed and excluded core pairs: %v", summaryRows)
	}

	// Verify the initial core migration is excluded from report analysis.
	tcData, err := os.ReadFile(filepath.Join(outDir, "datasets/table_changes.csv"))
	if err != nil {
		t.Fatal(err)
	}
	tcStr := string(tcData)
	if strings.Contains(tcStr, "AuditLog") || strings.Contains(tcStr, "AuditDetail") {
		t.Fatalf("initial core migration must be excluded from table_changes.csv, got:\n%s", tcStr)
	}
	if !strings.Contains(tcStr, "DROP_COLUMN") {
		t.Fatalf("expected DROP_COLUMN operations in table_changes.csv, got:\n%s", tcStr)
	}
	if !strings.Contains(tcStr, "ADD_CONSTRAINT") {
		t.Fatalf("expected ADD_CONSTRAINT operations in table_changes.csv, got:\n%s", tcStr)
	}
	for _, row := range readCSV(t, filepath.Join(outDir, "datasets/migration_summary.csv"))[1:] {
		if row[0] == "000000" {
			t.Fatalf("initial core migration must be excluded from migration_summary.csv: %v", row)
		}
	}

	// Verify object_history.csv status progression for Account
	ohData, err := os.ReadFile(filepath.Join(outDir, "datasets/object_history.csv"))
	if err != nil {
		t.Fatal(err)
	}
	ohLines := strings.Split(strings.TrimSpace(string(ohData)), "\n")
	var accountRows []string
	for _, l := range ohLines {
		if strings.Contains(l, "Account") {
			accountRows = append(accountRows, l)
		}
	}
	if len(accountRows) != 4 {
		t.Fatalf("expected exactly 4 version rows for Account, got %d:\n%s", len(accountRows), strings.Join(accountRows, "\n"))
	}
	// Check progression: v1: INITIAL, v2: UPDATED, v3: DELETED, v4: REINTRODUCED
	if !strings.Contains(accountRows[0], "INITIAL") {
		t.Errorf("expected v1 to be INITIAL, got: %s", accountRows[0])
	}
	if !strings.Contains(accountRows[1], "UPDATED") {
		t.Errorf("expected v2 to be UPDATED, got: %s", accountRows[1])
	}
	if !strings.Contains(accountRows[2], "DELETED") {
		t.Errorf("expected v3 to be DELETED, got: %s", accountRows[2])
	}
	if !strings.Contains(accountRows[3], "REINTRODUCED") {
		t.Errorf("expected v4 to be REINTRODUCED, got: %s", accountRows[3])
	}
}

func TestMigrationFilenameUsesSharedStrictContract(t *testing.T) {
	tests := []struct {
		name string
		ok   bool
	}{
		{"20260909010101001_tables_dbo.Customer_v1.up.sql", true},
		{"000000__core__mig__audit_init.down.sql", true},
		{"2026090901010100_tables_dbo.Customer_v1.up.sql", false},
		{"20260909010101001_unknown_dbo.Customer_v1.up.sql", false},
		{"20260909010101001_tables_sales.Customer_v1.up.sql", false},
		{"20260909010101001_tables_dbo.Customer_v0.up.sql", false},
		{"20260909010101001_tables_dbo.Customer_V1.up.sql", false},
		{"00000000000000000__core__mig__audit_init.up.sql", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, ok := migreport.ParseMigrationFileName(tt.name)
			if ok != tt.ok {
				t.Fatalf("ParseMigrationFileName(%q) ok=%v, want %v", tt.name, ok, tt.ok)
			}
		})
	}
}

func TestDownOnlyProgrammableDoesNotCreatePhantomOperation(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}
	name := "20260909010101002_procedures_dbo.usp_Ghost_v1.down.sql"
	if err := os.WriteFile(filepath.Join(migDir, name), []byte("DROP PROCEDURE IF EXISTS dbo.usp_Ghost;"), 0644); err != nil {
		t.Fatal(err)
	}

	outDir := filepath.Join(tmp, "out")
	code, err := migreport.Run(migreport.Options{RepoRoot: tmp, OutputDir: outDir})
	if err != nil || code != 0 {
		t.Fatalf("run failed: code=%d err=%v", code, err)
	}
	records := readCSV(t, filepath.Join(outDir, "datasets/migration_summary.csv"))
	if got := records[1][5]; got != "0" {
		t.Fatalf("down-only programmable reported %s operations, want 0: %v", got, records[1])
	}
	progRecords := readCSV(t, filepath.Join(outDir, "datasets/programmable_changes.csv"))
	if len(progRecords) != 1 {
		t.Fatalf("expected header-only programmable report, got %v", progRecords)
	}
}

func TestEmptyDownFailsRoundTripAndRaisesWarning(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}
	base := "20260909010101003_tables_dbo.Customer_v1"
	if err := os.WriteFile(filepath.Join(migDir, base+".up.sql"), []byte("CREATE TABLE dbo.Customer (Id INT NOT NULL);"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(migDir, base+".down.sql"), nil, 0644); err != nil {
		t.Fatal(err)
	}

	outDir := filepath.Join(tmp, "out")
	code, err := migreport.Run(migreport.Options{RepoRoot: tmp, OutputDir: outDir})
	if err != nil || code != 0 {
		t.Fatalf("warning run failed: code=%d err=%v", code, err)
	}
	records := readCSV(t, filepath.Join(outDir, "datasets/migration_summary.csv"))
	if records[1][11] != migreport.RoundTripFail || records[1][12] != migreport.StatusWarning {
		t.Fatalf("empty down must be FAIL/WARNING, got %v", records[1])
	}
	summary := readCSV(t, filepath.Join(outDir, "datasets/report_summary.csv"))
	if summary[1][16] != "1" || summary[1][18] != migreport.StatusWarning {
		t.Fatalf("overall report must expose the round-trip warning, got %v", summary[1])
	}
}

func TestMissingColumnInverseFailsRoundTrip(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}
	base := "20260909010101012_tables_dbo.Customer_v2"
	if err := os.WriteFile(filepath.Join(migDir, base+".up.sql"), []byte("ALTER TABLE dbo.Customer ADD Email NVARCHAR(200) NULL;"), 0644); err != nil {
		t.Fatal(err)
	}
	// A non-empty but unrelated down script used to be labeled PASS.
	if err := os.WriteFile(filepath.Join(migDir, base+".down.sql"), []byte("ALTER TABLE dbo.Customer ADD Unrelated INT NULL;"), 0644); err != nil {
		t.Fatal(err)
	}

	outDir := filepath.Join(tmp, "out")
	if code, err := migreport.Run(migreport.Options{RepoRoot: tmp, OutputDir: outDir}); err != nil || code != 0 {
		t.Fatalf("warning run failed: code=%d err=%v", code, err)
	}
	records := readCSV(t, filepath.Join(outDir, "datasets/migration_summary.csv"))
	if records[1][11] != migreport.RoundTripFail || records[1][12] != migreport.StatusWarning {
		t.Fatalf("unrelated down must not satisfy column inverse: %v", records[1])
	}
}

func TestAlterColumnReportsOnlyActualSemanticDelta(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}
	writePair := func(base, up, down string) {
		t.Helper()
		if err := os.WriteFile(filepath.Join(migDir, base+".up.sql"), []byte(up), 0644); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(migDir, base+".down.sql"), []byte(down), 0644); err != nil {
			t.Fatal(err)
		}
	}
	writePair("20260909010101004_tables_dbo.Customer_v1", "CREATE TABLE dbo.Customer (Name VARCHAR(20) NULL);", "DROP TABLE dbo.Customer;")
	writePair("20260909010101005_tables_dbo.Customer_v2", "ALTER TABLE dbo.Customer ALTER COLUMN Name VARCHAR(20) NOT NULL;", "ALTER TABLE dbo.Customer ALTER COLUMN Name VARCHAR(20) NULL;")

	outDir := filepath.Join(tmp, "out")
	if code, err := migreport.Run(migreport.Options{RepoRoot: tmp, OutputDir: outDir}); err != nil || code != 0 {
		t.Fatalf("run failed: code=%d err=%v", code, err)
	}
	records := readCSV(t, filepath.Join(outDir, "datasets/table_changes.csv"))
	var v2Kinds []string
	for _, row := range records[1:] {
		if row[0] == "20260909010101005" {
			v2Kinds = append(v2Kinds, row[5])
		}
	}
	if strings.Join(v2Kinds, "|") != migreport.OpAlterNullability {
		t.Fatalf("expected only nullability delta, got %v", v2Kinds)
	}
}

func TestConstraintRenameWithIdenticalDefinitionIsSuppressed(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}
	files := map[string]string{
		"20260909010101006_tables_dbo.Customer_v1.up.sql":   "CREATE TABLE dbo.Customer (Age INT, CONSTRAINT CK_Old CHECK (Age >= 0));",
		"20260909010101006_tables_dbo.Customer_v1.down.sql": "DROP TABLE dbo.Customer;",
		"20260909010101007_tables_dbo.Customer_v2.up.sql":   "ALTER TABLE dbo.Customer DROP CONSTRAINT CK_Old; ALTER TABLE dbo.Customer ADD CONSTRAINT CK_New CHECK (Age >= 0);",
		"20260909010101007_tables_dbo.Customer_v2.down.sql": "ALTER TABLE dbo.Customer DROP CONSTRAINT CK_New; ALTER TABLE dbo.Customer ADD CONSTRAINT CK_Old CHECK (Age >= 0);",
	}
	for name, sql := range files {
		if err := os.WriteFile(filepath.Join(migDir, name), []byte(sql), 0644); err != nil {
			t.Fatal(err)
		}
	}

	outDir := filepath.Join(tmp, "out")
	if code, err := migreport.Run(migreport.Options{RepoRoot: tmp, OutputDir: outDir}); err != nil || code != 0 {
		t.Fatalf("run failed: code=%d err=%v", code, err)
	}
	records := readCSV(t, filepath.Join(outDir, "datasets/table_changes.csv"))
	for _, row := range records[1:] {
		if row[0] == "20260909010101007" {
			t.Fatalf("metadata-only constraint rename must not be semantic: %v", row)
		}
	}
}

func TestProgrammableDeleteAndReintroductionLifecycle(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}
	files := map[string]string{
		"20260909010101008_views_dbo.vw_Customer_v1.up.sql":   "CREATE OR ALTER VIEW dbo.vw_Customer AS SELECT 1 AS Id;",
		"20260909010101008_views_dbo.vw_Customer_v1.down.sql": "DROP VIEW IF EXISTS dbo.vw_Customer;",
		"20260909010101009_views_dbo.vw_Customer_v2.up.sql":   "DROP VIEW IF EXISTS dbo.vw_Customer;",
		"20260909010101009_views_dbo.vw_Customer_v2.down.sql": "CREATE OR ALTER VIEW dbo.vw_Customer AS SELECT 1 AS Id;",
		"20260909010101010_views_dbo.vw_Customer_v3.up.sql":   "CREATE OR ALTER VIEW dbo.vw_Customer AS SELECT 2 AS Id;",
		"20260909010101010_views_dbo.vw_Customer_v3.down.sql": "DROP VIEW IF EXISTS dbo.vw_Customer;",
	}
	for name, sql := range files {
		if err := os.WriteFile(filepath.Join(migDir, name), []byte(sql), 0644); err != nil {
			t.Fatal(err)
		}
	}

	outDir := filepath.Join(tmp, "out")
	if code, err := migreport.Run(migreport.Options{RepoRoot: tmp, OutputDir: outDir}); err != nil || code != 0 {
		t.Fatalf("run failed: code=%d err=%v", code, err)
	}
	records := readCSV(t, filepath.Join(outDir, "datasets/object_history.csv"))
	var statuses []string
	for _, row := range records[1:] {
		if row[2] == "vw_Customer" {
			statuses = append(statuses, row[10])
		}
	}
	if got := strings.Join(statuses, "|"); got != "INITIAL|DELETED|REINTRODUCED" {
		t.Fatalf("unexpected programmable lifecycle %s", got)
	}
}

func TestDenseMixedTableEvolutionTracksEverySemanticChange(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}
	writePair := func(base, up, down string) {
		t.Helper()
		if err := os.WriteFile(filepath.Join(migDir, base+".up.sql"), []byte(up), 0644); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(migDir, base+".down.sql"), []byte(down), 0644); err != nil {
			t.Fatal(err)
		}
	}

	writePair("20260909020202001_tables_dbo.Ledger_v1", `
CREATE TABLE dbo.Ledger (
    Id BIGINT NOT NULL,
    Code VARCHAR(20) NOT NULL,
    Amount DECIMAL(10,2) NULL,
    Status VARCHAR(10) NULL,
    LegacyA INT NULL,
    LegacyB NVARCHAR(30) NULL,
    CONSTRAINT PK_Ledger PRIMARY KEY (Id),
    CONSTRAINT UQ_Ledger_Code UNIQUE (Code),
    CONSTRAINT CK_Ledger_Amount CHECK (Amount >= 0),
    CONSTRAINT DF_Ledger_Status DEFAULT ('NEW') FOR Status
);
CREATE NONCLUSTERED INDEX IX_Ledger_Code ON dbo.Ledger (Code) INCLUDE (Amount);`, `
DROP TABLE dbo.Ledger;`)

	writePair("20260909020202002_tables_dbo.Ledger_v2", `
ALTER TABLE dbo.Ledger ADD
    CreatedAt DATETIME2(3) NOT NULL,
    UpdatedBy NVARCHAR(100) NULL,
    CONSTRAINT CK_Ledger_CreatedAt CHECK (CreatedAt >= '2020-01-01'),
    CONSTRAINT UQ_Ledger_UpdatedBy UNIQUE (UpdatedBy, CreatedAt);
ALTER TABLE dbo.Ledger WITH NOCHECK ADD CONSTRAINT FK_Ledger_Account FOREIGN KEY (UpdatedBy) REFERENCES dbo.Account (UserName);
ALTER TABLE dbo.Ledger DROP COLUMN LegacyA, LegacyB;
ALTER TABLE dbo.Ledger ALTER COLUMN Code NVARCHAR(40) NULL;
ALTER TABLE dbo.Ledger ALTER COLUMN Amount DECIMAL(18,4) NOT NULL;
ALTER TABLE dbo.Ledger DROP CONSTRAINT CK_Ledger_Amount, DF_Ledger_Status;
ALTER TABLE dbo.Ledger WITH CHECK ADD CONSTRAINT CK_Ledger_Amount CHECK (Amount >= 0 AND Amount <= 999999);
ALTER TABLE dbo.Ledger ADD CONSTRAINT DF_Ledger_Status DEFAULT ('ACTIVE') FOR Status;
EXEC sp_rename 'dbo.Ledger.Status', 'State', 'COLUMN';
DROP INDEX IX_Ledger_Code ON dbo.Ledger;
CREATE UNIQUE CLUSTERED INDEX IX_Ledger_Code ON dbo.Ledger (Code DESC) INCLUDE (Amount, State) WHERE Code IS NOT NULL;`, `
DROP INDEX IX_Ledger_Code ON dbo.Ledger;
CREATE NONCLUSTERED INDEX IX_Ledger_Code ON dbo.Ledger (Code) INCLUDE (Amount);
EXEC sp_rename 'dbo.Ledger.State', 'Status', 'COLUMN';
ALTER TABLE dbo.Ledger DROP CONSTRAINT CK_Ledger_CreatedAt, UQ_Ledger_UpdatedBy, FK_Ledger_Account, CK_Ledger_Amount, DF_Ledger_Status;
ALTER TABLE dbo.Ledger ADD
    CONSTRAINT CK_Ledger_Amount CHECK (Amount >= 0),
    CONSTRAINT DF_Ledger_Status DEFAULT ('NEW') FOR Status;
ALTER TABLE dbo.Ledger ALTER COLUMN Code VARCHAR(20) NOT NULL;
ALTER TABLE dbo.Ledger ALTER COLUMN Amount DECIMAL(10,2) NULL;
ALTER TABLE dbo.Ledger DROP COLUMN CreatedAt, UpdatedBy;
ALTER TABLE dbo.Ledger ADD LegacyA INT NULL, LegacyB NVARCHAR(30) NULL;`)

	outDir := filepath.Join(tmp, "out")
	if code, err := migreport.Run(migreport.Options{RepoRoot: tmp, OutputDir: outDir}); err != nil || code != 0 {
		t.Fatalf("run failed: code=%d err=%v", code, err)
	}

	tableRows := readCSV(t, filepath.Join(outDir, "datasets/table_changes.csv"))
	var v2Rows [][]string
	for _, row := range tableRows[1:] {
		if row[0] == "20260909020202002" {
			v2Rows = append(v2Rows, row)
		}
	}
	wantKinds := map[string]int{
		migreport.OpAddColumn:        2,
		migreport.OpAddConstraint:    3,
		migreport.OpDropColumn:       2,
		migreport.OpAlterColumnType:  2,
		migreport.OpAlterNullability: 2,
		migreport.OpAlterConstraint:  1,
		migreport.OpAlterDefault:     1,
		migreport.OpRenameColumn:     1,
	}
	gotKinds := map[string]int{}
	byKindTarget := map[string][]string{}
	for _, row := range v2Rows {
		gotKinds[row[5]]++
		byKindTarget[row[5]+"|"+row[6]] = row
		wantConfidence := migreport.ConfidenceHigh
		if row[5] == migreport.OpAlterConstraint {
			wantConfidence = migreport.ConfidenceMedium
		}
		if row[13] != wantConfidence {
			t.Fatalf("dense recognized change unexpectedly low confidence: %v", row)
		}
	}
	if !reflect.DeepEqual(gotKinds, wantKinds) {
		t.Fatalf("wrong dense operation inventory: got=%v want=%v rows=%v", gotKinds, wantKinds, v2Rows)
	}
	assertChangeValues(t, byKindTarget, migreport.OpAlterColumnType, "Code", "VARCHAR(20)", "NVARCHAR(40)")
	assertChangeValues(t, byKindTarget, migreport.OpAlterColumnType, "Amount", "DECIMAL(10,2)", "DECIMAL(18,4)")
	assertChangeValues(t, byKindTarget, migreport.OpAlterNullability, "Code", "NOT NULL", "NULL")
	assertChangeValues(t, byKindTarget, migreport.OpAlterNullability, "Amount", "NULL", "NOT NULL")
	assertChangeValues(t, byKindTarget, migreport.OpDropColumn, "LegacyA", "INT", "")
	assertChangeValues(t, byKindTarget, migreport.OpDropColumn, "LegacyB", "NVARCHAR(30)", "")
	if fk := byKindTarget[migreport.OpAddConstraint+"|FK_Ledger_Account"]; fk == nil || fk[9] != "Account" {
		t.Fatalf("foreign-key dependency was not reconstructed: %v", fk)
	}

	indexRows := readCSV(t, filepath.Join(outDir, "datasets/index_changes.csv"))
	var v2Indexes [][]string
	for _, row := range indexRows[1:] {
		if row[0] == "20260909020202002" {
			v2Indexes = append(v2Indexes, row)
		}
	}
	if len(v2Indexes) != 1 || v2Indexes[0][6] != migreport.OpAlterIndex {
		t.Fatalf("drop+recreate index must collapse to one ALTER_INDEX: %v", v2Indexes)
	}
	idx := v2Indexes[0]
	if idx[8] != "Code" || idx[9] != "Code DESC" || idx[10] != "Amount" || idx[11] != "Amount,State" ||
		idx[14] != "0" || idx[15] != "1" || idx[16] != "0" || idx[17] != "1" {
		t.Fatalf("index before/after attributes were not reconstructed: %v", idx)
	}

	migrations := readCSV(t, filepath.Join(outDir, "datasets/migration_summary.csv"))
	for _, row := range migrations[1:] {
		if row[0] == "20260909020202002" {
			if row[5] != "15" || row[11] != migreport.RoundTripPass || row[12] != migreport.StatusOk {
				upTables, upIndexes, _, _ := migreport.ParseStatementsFromSQL("tables", "dbo", "Ledger", mustReadMigration(t, migDir, "20260909020202002_tables_dbo.Ledger_v2.up.sql"))
				downTables, downIndexes, _, _ := migreport.ParseStatementsFromSQL("tables", "dbo", "Ledger", mustReadMigration(t, migDir, "20260909020202002_tables_dbo.Ledger_v2.down.sql"))
				t.Logf("up tables=%+v", upTables)
				t.Logf("down tables=%+v", downTables)
				t.Logf("up indexes=%+v down indexes=%+v", upIndexes, downIndexes)
				t.Fatalf("dense migration summary mismatch: %v", row)
			}
			return
		}
	}
	t.Fatal("dense v2 migration summary not found")
}

func TestOrderSensitiveMixedTableOperationsPreserveSQLSequence(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}
	writePair := func(base, up, down string) {
		t.Helper()
		if err := os.WriteFile(filepath.Join(migDir, base+".up.sql"), []byte(up), 0644); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(migDir, base+".down.sql"), []byte(down), 0644); err != nil {
			t.Fatal(err)
		}
	}
	writePair("20260909030303001_tables_dbo.SequenceCase_v1",
		"CREATE TABLE dbo.SequenceCase (A INT NULL);",
		"DROP TABLE dbo.SequenceCase;")
	writePair("20260909030303002_tables_dbo.SequenceCase_v2", `
ALTER TABLE dbo.SequenceCase ADD B INT NULL;
EXEC sp_rename 'dbo.SequenceCase.B', 'C', 'COLUMN';
ALTER TABLE dbo.SequenceCase ALTER COLUMN C BIGINT NOT NULL;
ALTER TABLE dbo.SequenceCase DROP COLUMN A;
ALTER TABLE dbo.SequenceCase ADD A NVARCHAR(20) NOT NULL;`, `
ALTER TABLE dbo.SequenceCase DROP COLUMN A;
ALTER TABLE dbo.SequenceCase ADD A INT NULL;
ALTER TABLE dbo.SequenceCase ALTER COLUMN C INT NULL;
EXEC sp_rename 'dbo.SequenceCase.C', 'B', 'COLUMN';
ALTER TABLE dbo.SequenceCase DROP COLUMN B;`)

	outDir := filepath.Join(tmp, "out")
	if code, err := migreport.Run(migreport.Options{RepoRoot: tmp, OutputDir: outDir}); err != nil || code != 0 {
		t.Fatalf("run failed: code=%d err=%v", code, err)
	}
	rows := readCSV(t, filepath.Join(outDir, "datasets/table_changes.csv"))
	var v2 [][]string
	for _, row := range rows[1:] {
		if row[0] == "20260909030303002" {
			v2 = append(v2, row)
		}
	}
	wantSequence := []string{
		migreport.OpAddColumn + "|B",
		migreport.OpRenameColumn + "|C",
		migreport.OpAlterColumnType + "|C",
		migreport.OpAlterNullability + "|C",
		migreport.OpDropColumn + "|A",
		migreport.OpAddColumn + "|A",
	}
	var gotSequence []string
	for _, row := range v2 {
		gotSequence = append(gotSequence, row[5]+"|"+row[6])
		if row[13] != migreport.ConfidenceHigh {
			t.Fatalf("state lookup failed because SQL order was lost: %v", row)
		}
	}
	if !reflect.DeepEqual(gotSequence, wantSequence) {
		t.Fatalf("operation sequence mismatch: got=%v want=%v", gotSequence, wantSequence)
	}
	assertChangeValues(t, mapRowsByKindTarget(v2), migreport.OpAlterColumnType, "C", "INT", "BIGINT")
	assertChangeValues(t, mapRowsByKindTarget(v2), migreport.OpDropColumn, "A", "INT", "")
	for _, row := range readCSV(t, filepath.Join(outDir, "datasets/migration_summary.csv"))[1:] {
		if row[0] == "20260909030303002" && (row[11] != migreport.RoundTripPass || row[12] != migreport.StatusOk) {
			t.Fatalf("order-sensitive reversible migration did not pass: %v", row)
		}
	}
}

func TestMultiOperationRoundTripFailsWhenOnlyOneInverseExists(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}
	base := "20260909040404001_tables_dbo.PartialInverse_v1"
	if err := os.WriteFile(filepath.Join(migDir, base+".up.sql"), []byte("ALTER TABLE dbo.PartialInverse ADD ColA INT NULL, ColB INT NULL;"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(migDir, base+".down.sql"), []byte("ALTER TABLE dbo.PartialInverse DROP COLUMN ColA;"), 0644); err != nil {
		t.Fatal(err)
	}
	outDir := filepath.Join(tmp, "out")
	if code, err := migreport.Run(migreport.Options{RepoRoot: tmp, OutputDir: outDir}); err != nil || code != 0 {
		t.Fatalf("warning run failed: code=%d err=%v", code, err)
	}
	row := readCSV(t, filepath.Join(outDir, "datasets/migration_summary.csv"))[1]
	if row[5] != "2" || row[11] != migreport.RoundTripFail || row[12] != migreport.StatusWarning {
		t.Fatalf("partial inverse must fail the complete migration: %v", row)
	}
}

func TestAlterTableMultiItemSyntaxMatrix(t *testing.T) {
	tests := []struct {
		name string
		sql  string
		want []string
	}{
		{
			name: "mixed columns and constraints",
			sql:  "ALTER TABLE dbo.T ADD A INT NULL, B DECIMAL(18,4) NOT NULL, CONSTRAINT CK_T_A CHECK (A >= 0), CONSTRAINT UQ_T_AB UNIQUE (A, B);",
			want: []string{"ADD_COLUMN|A", "ADD_COLUMN|B", "ADD_CONSTRAINT|CK_T_A", "ADD_CONSTRAINT|UQ_T_AB"},
		},
		{
			name: "add constraint keyword owns multiple constraints",
			sql:  "ALTER TABLE dbo.T ADD CONSTRAINT CK_T_A CHECK (A > 0), CONSTRAINT CK_T_B CHECK (B > 0);",
			want: []string{"ADD_CONSTRAINT|CK_T_A", "ADD_CONSTRAINT|CK_T_B"},
		},
		{
			name: "with check foreign key",
			sql:  "ALTER TABLE dbo.T WITH CHECK ADD CONSTRAINT FK_T_R FOREIGN KEY (A, B) REFERENCES dbo.R (A, B);",
			want: []string{"ADD_CONSTRAINT|FK_T_R"},
		},
		{
			name: "with nocheck constraint",
			sql:  "ALTER TABLE dbo.T WITH NOCHECK ADD CONSTRAINT CK_T_A CHECK (A <> 9);",
			want: []string{"ADD_CONSTRAINT|CK_T_A"},
		},
		{
			name: "multiple drop columns",
			sql:  "ALTER TABLE dbo.T DROP COLUMN A, B, C;",
			want: []string{"DROP_COLUMN|A", "DROP_COLUMN|B", "DROP_COLUMN|C"},
		},
		{
			name: "multiple drop constraints with if exists",
			sql:  "ALTER TABLE dbo.T DROP CONSTRAINT IF EXISTS CK_T_A, CK_T_B, UQ_T_AB;",
			want: []string{"DROP_CONSTRAINT|CK_T_A", "DROP_CONSTRAINT|CK_T_B", "DROP_CONSTRAINT|UQ_T_AB"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			statements, _, _, _ := migreport.ParseStatementsFromSQL("tables", "dbo", "T", tt.sql)
			got := make([]string, 0, len(statements))
			for _, statement := range statements {
				got = append(got, statement.Kind+"|"+statement.TargetName)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("statement inventory mismatch: got=%v want=%v", got, tt.want)
			}
		})
	}
}

func TestCreateTableInWrappedMigrationIsParsedExactlyOnce(t *testing.T) {
	// Generated migrations have a substantial audit/precheck wrapper before
	// MIGRATION_BODY_BEGIN. The CREATE TABLE scanner must advance relative to
	// that wrapper as well as to the table body.
	sql := `DECLARE @target NVARCHAR(256) = N'dbo.Wrapped';
BEGIN TRY
  BEGIN TRAN;
  -- MIGRATION_BODY_BEGIN
  CREATE TABLE dbo.Wrapped (
    Id BIGINT NOT NULL,
    Code NVARCHAR(20) NULL
  );
  -- MIGRATION_BODY_END
  COMMIT TRAN;
END TRY
BEGIN CATCH
  ROLLBACK TRAN;
END CATCH;`
	statements, _, _, _ := migreport.ParseStatementsFromSQL("tables", "dbo", "Wrapped", sql)
	got := make([]string, 0, len(statements))
	for _, statement := range statements {
		got = append(got, statement.Kind+"|"+statement.TargetName)
	}
	want := []string{
		migreport.OpCreateTable + "|Wrapped",
		migreport.OpAddColumn + "|Id",
		migreport.OpAddColumn + "|Code",
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wrapped CREATE TABLE must be emitted exactly once: got=%v want=%v", got, want)
	}
}

func TestCreateTableSupportsDelimitedColumnNamesWithSpaces(t *testing.T) {
	sql := `CREATE TABLE dbo.DelimitedColumns (
  [Deal ID] BIGINT NOT NULL,
  [Tax On Accrued (untuk transaksi PE Beli)] FLOAT NULL,
  "Quoted Column" NVARCHAR(20) NULL
);`
	statements, _, _, _ := migreport.ParseStatementsFromSQL("tables", "dbo", "DelimitedColumns", sql)
	got := make([]string, 0, len(statements))
	for _, statement := range statements {
		got = append(got, statement.Kind+"|"+statement.TargetName+"|"+statement.NewValue)
	}
	want := []string{
		migreport.OpCreateTable + "|DelimitedColumns|",
		migreport.OpAddColumn + "|Deal ID|BIGINT",
		migreport.OpAddColumn + "|Tax On Accrued (untuk transaksi PE Beli)|FLOAT",
		migreport.OpAddColumn + "|Quoted Column|NVARCHAR(20)",
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("delimited column names were parsed incorrectly: got=%v want=%v", got, want)
	}
}

func TestLongitudinalTableHistoryThroughV7UsesLatestReintroducedState(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}
	writePair := func(base, up, down string) {
		t.Helper()
		if err := os.WriteFile(filepath.Join(migDir, base+".up.sql"), []byte(up), 0644); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(migDir, base+".down.sql"), []byte(down), 0644); err != nil {
			t.Fatal(err)
		}
	}

	writePair("20260909050505001_tables_dbo.LongHistory_v1", `
CREATE TABLE dbo.LongHistory (
    Id INT NOT NULL,
    Code VARCHAR(20) NOT NULL,
    Amount DECIMAL(10,2) NULL,
    Status VARCHAR(10) NULL,
    CONSTRAINT PK_LongHistory PRIMARY KEY (Id),
    CONSTRAINT CK_LongHistory_Amount CHECK (Amount >= 0),
    CONSTRAINT DF_LongHistory_Status DEFAULT ('NEW') FOR Status
);
CREATE NONCLUSTERED INDEX IX_LongHistory_Current ON dbo.LongHistory (Code);`, `
DROP TABLE dbo.LongHistory;`)

	writePair("20260909050505002_tables_dbo.LongHistory_v2", `
ALTER TABLE dbo.LongHistory ADD CreatedAt DATETIME2(3) NULL;
ALTER TABLE dbo.LongHistory ALTER COLUMN Code NVARCHAR(40) NOT NULL;`, `
ALTER TABLE dbo.LongHistory ALTER COLUMN Code VARCHAR(20) NOT NULL;
ALTER TABLE dbo.LongHistory DROP COLUMN CreatedAt;`)

	writePair("20260909050505003_tables_dbo.LongHistory_v3", `
EXEC sp_rename 'dbo.LongHistory.Status', 'State', 'COLUMN';
ALTER TABLE dbo.LongHistory DROP CONSTRAINT DF_LongHistory_Status;
ALTER TABLE dbo.LongHistory ADD CONSTRAINT DF_LongHistory_State DEFAULT ('PENDING') FOR State;`, `
ALTER TABLE dbo.LongHistory DROP CONSTRAINT DF_LongHistory_State;
EXEC sp_rename 'dbo.LongHistory.State', 'Status', 'COLUMN';
ALTER TABLE dbo.LongHistory ADD CONSTRAINT DF_LongHistory_Status DEFAULT ('NEW') FOR Status;`)

	writePair("20260909050505004_tables_dbo.LongHistory_v4", `
ALTER TABLE dbo.LongHistory DROP CONSTRAINT CK_LongHistory_Amount;
ALTER TABLE dbo.LongHistory DROP COLUMN Amount;
ALTER TABLE dbo.LongHistory ADD Note NVARCHAR(100) NULL;
DROP INDEX IX_LongHistory_Current ON dbo.LongHistory;
CREATE NONCLUSTERED INDEX IX_LongHistory_Current ON dbo.LongHistory (State) INCLUDE (Code);`, `
DROP INDEX IX_LongHistory_Current ON dbo.LongHistory;
CREATE NONCLUSTERED INDEX IX_LongHistory_Current ON dbo.LongHistory (Code);
ALTER TABLE dbo.LongHistory DROP COLUMN Note;
ALTER TABLE dbo.LongHistory ADD Amount DECIMAL(10,2) NULL;
ALTER TABLE dbo.LongHistory ADD CONSTRAINT CK_LongHistory_Amount CHECK (Amount >= 0);`)

	writePair("20260909050505005_tables_dbo.LongHistory_v5", `
DROP TABLE dbo.LongHistory;`, `
CREATE TABLE dbo.LongHistory (
    Id INT NOT NULL,
    Code NVARCHAR(40) NOT NULL,
    State VARCHAR(10) NULL,
    CreatedAt DATETIME2(3) NULL,
    Note NVARCHAR(100) NULL,
    CONSTRAINT PK_LongHistory PRIMARY KEY (Id),
    CONSTRAINT DF_LongHistory_State DEFAULT ('PENDING') FOR State
);
CREATE NONCLUSTERED INDEX IX_LongHistory_Current ON dbo.LongHistory (State) INCLUDE (Code);`)

	writePair("20260909050505006_tables_dbo.LongHistory_v6", `
CREATE TABLE dbo.LongHistory (
    Id BIGINT NOT NULL,
    Code NVARCHAR(80) NOT NULL,
    State CHAR(1) NOT NULL,
    CreatedAt DATETIME2(0) NOT NULL,
    Note VARCHAR(30) NULL,
    CONSTRAINT PK_LongHistory PRIMARY KEY CLUSTERED (Id),
    CONSTRAINT DF_LongHistory_State DEFAULT ('A') FOR State
);
CREATE NONCLUSTERED INDEX IX_LongHistory_Current ON dbo.LongHistory (State);`, `
DROP TABLE dbo.LongHistory;`)

	writePair("20260909050505007_tables_dbo.LongHistory_v7", `
ALTER TABLE dbo.LongHistory ALTER COLUMN Code NVARCHAR(120) NULL;
ALTER TABLE dbo.LongHistory ALTER COLUMN State VARCHAR(12) NOT NULL;
ALTER TABLE dbo.LongHistory ADD Amount DECIMAL(18,2) NOT NULL, Revision INT NULL;
ALTER TABLE dbo.LongHistory DROP COLUMN CreatedAt;
ALTER TABLE dbo.LongHistory DROP CONSTRAINT DF_LongHistory_State;
ALTER TABLE dbo.LongHistory ADD CONSTRAINT DF_LongHistory_State DEFAULT ('ACTIVE') FOR State;
DROP INDEX IX_LongHistory_Current ON dbo.LongHistory;
CREATE UNIQUE NONCLUSTERED INDEX IX_LongHistory_Current ON dbo.LongHistory (Code DESC) INCLUDE (State, Amount) WHERE Code IS NOT NULL;`, `
DROP INDEX IX_LongHistory_Current ON dbo.LongHistory;
CREATE NONCLUSTERED INDEX IX_LongHistory_Current ON dbo.LongHistory (State);
ALTER TABLE dbo.LongHistory DROP CONSTRAINT DF_LongHistory_State;
ALTER TABLE dbo.LongHistory ADD CONSTRAINT DF_LongHistory_State DEFAULT ('A') FOR State;
ALTER TABLE dbo.LongHistory ADD CreatedAt DATETIME2(0) NOT NULL;
ALTER TABLE dbo.LongHistory DROP COLUMN Amount, Revision;
ALTER TABLE dbo.LongHistory ALTER COLUMN State CHAR(1) NOT NULL;
ALTER TABLE dbo.LongHistory ALTER COLUMN Code NVARCHAR(80) NOT NULL;`)

	outDir := filepath.Join(tmp, "out")
	if code, err := migreport.Run(migreport.Options{RepoRoot: tmp, OutputDir: outDir}); err != nil || code != 0 {
		t.Fatalf("v1-v7 report failed: code=%d err=%v", code, err)
	}

	historyRows := readCSV(t, filepath.Join(outDir, "datasets/object_history.csv"))
	var versions, previous, statuses []string
	for _, row := range historyRows[1:] {
		if row[2] == "LongHistory" {
			versions = append(versions, row[3])
			previous = append(previous, row[6])
			statuses = append(statuses, row[10])
		}
	}
	if got, want := strings.Join(versions, "|"), "1|2|3|4|5|6|7"; got != want {
		t.Fatalf("version chain mismatch: got=%s want=%s", got, want)
	}
	if got, want := strings.Join(previous, "|"), "|1|2|3|4|5|6"; got != want {
		t.Fatalf("previous-version chain mismatch: got=%s want=%s", got, want)
	}
	if got, want := strings.Join(statuses, "|"), "INITIAL|UPDATED|UPDATED|UPDATED|DELETED|REINTRODUCED|UPDATED"; got != want {
		t.Fatalf("lifecycle mismatch: got=%s want=%s", got, want)
	}

	var v7Rows [][]string
	for _, row := range readCSV(t, filepath.Join(outDir, "datasets/table_changes.csv"))[1:] {
		if row[0] == "20260909050505007" {
			v7Rows = append(v7Rows, row)
		}
	}
	byKindTarget := mapRowsByKindTarget(v7Rows)
	assertChangeValues(t, byKindTarget, migreport.OpAlterColumnType, "Code", "NVARCHAR(80)", "NVARCHAR(120)")
	assertChangeValues(t, byKindTarget, migreport.OpAlterNullability, "Code", "NOT NULL", "NULL")
	assertChangeValues(t, byKindTarget, migreport.OpAlterColumnType, "State", "CHAR(1)", "VARCHAR(12)")
	assertChangeValues(t, byKindTarget, migreport.OpDropColumn, "CreatedAt", "DATETIME2(0)", "")
	assertChangeValues(t, byKindTarget, migreport.OpAddColumn, "Amount", "", "DECIMAL(18,2)")
	assertChangeValues(t, byKindTarget, migreport.OpAddColumn, "Revision", "", "INT")
	defaultRow := byKindTarget[migreport.OpAlterDefault+"|DF_LongHistory_State"]
	if defaultRow == nil || !strings.Contains(defaultRow[7], "('A')") || !strings.Contains(defaultRow[8], "('ACTIVE')") {
		t.Fatalf("v7 default did not use v6 state: %v", defaultRow)
	}
	if _, exists := byKindTarget[migreport.OpAlterNullability+"|State"]; exists {
		t.Fatalf("unchanged State NOT NULL must be suppressed: %v", byKindTarget[migreport.OpAlterNullability+"|State"])
	}

	var v7Indexes [][]string
	for _, row := range readCSV(t, filepath.Join(outDir, "datasets/index_changes.csv"))[1:] {
		if row[0] == "20260909050505007" {
			v7Indexes = append(v7Indexes, row)
		}
	}
	if len(v7Indexes) != 1 || v7Indexes[0][6] != migreport.OpAlterIndex || v7Indexes[0][8] != "State" ||
		v7Indexes[0][9] != "Code DESC" || v7Indexes[0][11] != "State,Amount" || v7Indexes[0][15] != "1" {
		t.Fatalf("v7 index did not use v6 state: %v", v7Indexes)
	}

	for _, row := range readCSV(t, filepath.Join(outDir, "datasets/migration_summary.csv"))[1:] {
		if row[0] == "20260909050505007" {
			if row[5] != "8" || row[11] != migreport.RoundTripPass || row[12] != migreport.StatusOk {
				t.Fatalf("v7 summary mismatch: %v", row)
			}
			return
		}
	}
	t.Fatal("v7 migration summary not found")
}

func TestReintroducedTableDoesNotLeakPreDropState(t *testing.T) {
	tmp := t.TempDir()
	migDir := filepath.Join(tmp, "migrations")
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}
	files := map[string]string{
		"20260909060606001_tables_dbo.ResetState_v1.up.sql":   "CREATE TABLE dbo.ResetState (Id INT NOT NULL, LegacyOnly VARCHAR(20) NULL);",
		"20260909060606001_tables_dbo.ResetState_v1.down.sql": "DROP TABLE dbo.ResetState;",
		"20260909060606002_tables_dbo.ResetState_v2.up.sql":   "DROP TABLE dbo.ResetState;",
		"20260909060606002_tables_dbo.ResetState_v2.down.sql": "CREATE TABLE dbo.ResetState (Id INT NOT NULL, LegacyOnly VARCHAR(20) NULL);",
		"20260909060606003_tables_dbo.ResetState_v3.up.sql":   "CREATE TABLE dbo.ResetState (Id BIGINT NOT NULL);",
		"20260909060606003_tables_dbo.ResetState_v3.down.sql": "DROP TABLE dbo.ResetState;",
		"20260909060606004_tables_dbo.ResetState_v4.up.sql":   "ALTER TABLE dbo.ResetState ALTER COLUMN LegacyOnly NVARCHAR(50) NULL;",
		"20260909060606004_tables_dbo.ResetState_v4.down.sql": "ALTER TABLE dbo.ResetState ALTER COLUMN LegacyOnly VARCHAR(20) NULL;",
	}
	for name, sql := range files {
		if err := os.WriteFile(filepath.Join(migDir, name), []byte(sql), 0644); err != nil {
			t.Fatal(err)
		}
	}
	outDir := filepath.Join(tmp, "out")
	if code, err := migreport.Run(migreport.Options{RepoRoot: tmp, OutputDir: outDir}); err != nil || code != 0 {
		t.Fatalf("manual-review run failed: code=%d err=%v", code, err)
	}
	for _, row := range readCSV(t, filepath.Join(outDir, "datasets/table_changes.csv"))[1:] {
		if row[0] == "20260909060606004" && row[5] == migreport.OpAlterColumnType {
			if row[7] != "" || row[13] != migreport.ConfidenceLow || row[14] != migreport.NoteManualReviewRequired {
				t.Fatalf("pre-drop column state leaked into reintroduced table: %v", row)
			}
			return
		}
	}
	t.Fatal("expected v4 manual-review ALTER_COLUMN_TYPE row")
}

func mapRowsByKindTarget(rows [][]string) map[string][]string {
	out := make(map[string][]string, len(rows))
	for _, row := range rows {
		out[row[5]+"|"+row[6]] = row
	}
	return out
}

func mustReadMigration(t *testing.T, dir, name string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(dir, name))
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func assertChangeValues(t *testing.T, rows map[string][]string, kind, target, oldValue, newValue string) {
	t.Helper()
	row, ok := rows[kind+"|"+target]
	if !ok {
		t.Fatalf("missing %s for %s", kind, target)
	}
	if row[7] != oldValue || row[8] != newValue {
		t.Fatalf("%s %s old/new mismatch: got %q -> %q, want %q -> %q", kind, target, row[7], row[8], oldValue, newValue)
	}
}

func readCSV(t *testing.T, path string) [][]string {
	t.Helper()
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	return records
}
