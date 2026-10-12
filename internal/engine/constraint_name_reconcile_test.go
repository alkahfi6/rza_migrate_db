package engine

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"rza_migrate_db/internal/config"
)

func TestMigrateCheckPlansAndAppliesConstraintNameRepairWithoutGeneratingMigrations(t *testing.T) {
	repoRoot := repoRootFromEnginePackageDir(t)
	source := filepath.Join(repoRoot, "fixtures", "qa_rules", "07_vk_whitespace_only", "input")
	workRoot := t.TempDir()
	if err := copyDirForAutoIndentTest(source, workRoot); err != nil {
		t.Fatal(err)
	}
	cfg := config.Default(workRoot)
	if _, err := New(cfg).Run("fmt"); err != nil {
		t.Fatal(err)
	}
	tablePath := filepath.Join(workRoot, "db_object", "schema", "tables", "dbo.TableA.sql")
	raw, err := os.ReadFile(tablePath)
	if err != nil {
		t.Fatal(err)
	}
	drifted := strings.ReplaceAll(string(raw), "PK_TableA", "PK__TableA__generated")
	drifted = strings.ReplaceAll(drifted, "DF_TableA__status", "DF__TableA__generated")
	if err := os.WriteFile(tablePath, []byte(drifted), 0o644); err != nil {
		t.Fatal(err)
	}
	beforeMigrations, err := os.ReadDir(filepath.Join(workRoot, "migrations"))
	if err != nil {
		t.Fatal(err)
	}

	checkCfg := config.Default(workRoot)
	checkCfg.EnableFmtCheck = true
	report, err := New(checkCfg).Run("check")
	if err != nil {
		t.Fatal(err)
	}
	if !hasRule(report.Fail.Items, "CONSTRAINT_NAME_CHECK_CHANGES_REQUIRED") {
		t.Fatalf("CI check did not detect safe constraint repair: %+v", report.Fail.Items)
	}
	afterDryRun, err := os.ReadFile(tablePath)
	if err != nil {
		t.Fatal(err)
	}
	if string(afterDryRun) != drifted {
		t.Fatal("CI check modified the object source")
	}

	if _, err := New(config.Default(workRoot)).Run("check"); err != nil {
		t.Fatal(err)
	}
	afterApply, err := os.ReadFile(tablePath)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(afterApply), "PK_TableA") || !strings.Contains(string(afterApply), "DF_TableA__status") || strings.Contains(string(afterApply), "__generated") {
		t.Fatalf("interactive checker did not restore migration-history names:\n%s", afterApply)
	}
	afterMigrations, err := os.ReadDir(filepath.Join(workRoot, "migrations"))
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(entryNames(beforeMigrations), entryNames(afterMigrations)) {
		t.Fatal("rza_migrate_check changed the migrations directory")
	}
}

func entryNames(entries []os.DirEntry) []string {
	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		names = append(names, entry.Name())
	}
	return names
}

func TestEquivalentConstraintNameMappingMatchesOnlyIdenticalDefinitions(t *testing.T) {
	previous := `CREATE TABLE dbo.T (
  [id] INT NOT NULL CONSTRAINT [DF_T_id_old] DEFAULT ((0)),
  [parent_id] INT NULL,
  CONSTRAINT [PK_T_old] PRIMARY KEY CLUSTERED ([id] ASC),
  CONSTRAINT [FK_T_old] FOREIGN KEY ([parent_id]) REFERENCES dbo.Parent ([id])
);
`
	desired := `CREATE TABLE dbo.T (
  [id] INT NOT NULL CONSTRAINT [DF_T_id_new] DEFAULT ((0)),
  [parent_id] INT NULL,
  CONSTRAINT [PK_T_new] PRIMARY KEY CLUSTERED ([id] ASC),
  CONSTRAINT [FK_T_new] FOREIGN KEY ([parent_id]) REFERENCES dbo.Parent ([id])
);
`

	mapping, err := equivalentConstraintNameMapping(previous, desired)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{
		"DF_T_id_new": "DF_T_id_old",
		"PK_T_new":    "PK_T_old",
		"FK_T_new":    "FK_T_old",
	}
	if len(mapping) != len(want) {
		t.Fatalf("mapping count = %d, want %d: %#v", len(mapping), len(want), mapping)
	}
	for next, previous := range want {
		if mapping[next] != previous {
			t.Fatalf("mapping[%q] = %q, want %q", next, mapping[next], previous)
		}
	}
}

func TestEquivalentConstraintNameMappingDoesNotHideDefinitionChange(t *testing.T) {
	previous := `CREATE TABLE dbo.T (
  [id] INT NOT NULL CONSTRAINT [DF_T_old] DEFAULT ((0)),
  CONSTRAINT [PK_T_old] PRIMARY KEY CLUSTERED ([id] ASC)
);
`
	desired := `CREATE TABLE dbo.T (
  [id] INT NOT NULL CONSTRAINT [DF_T_new] DEFAULT ((1)),
  CONSTRAINT [PK_T_new] PRIMARY KEY NONCLUSTERED ([id] DESC)
);
`
	mapping, err := equivalentConstraintNameMapping(previous, desired)
	if err != nil {
		t.Fatal(err)
	}
	if len(mapping) != 0 {
		t.Fatalf("definition changes must not be reconciled by name: %#v", mapping)
	}
}

func TestReconcileDesiredTableConstraintNamesWritesHistoryNamesAndPreservesRealDelta(t *testing.T) {
	root := t.TempDir()
	sourceFile := filepath.Join("db_object", "tables", "dbo.T.sql")
	abs := filepath.Join(root, sourceFile)
	if err := os.MkdirAll(filepath.Dir(abs), 0o755); err != nil {
		t.Fatal(err)
	}
	previous := `CREATE TABLE dbo.T (
  [id] INT NOT NULL,
  CONSTRAINT [PK__T__old] PRIMARY KEY CLUSTERED ([id] ASC)
);
`
	desiredSQL := `CREATE TABLE dbo.T (
  [id] INT NOT NULL,
  [new_value] VARCHAR(50) NULL,
  CONSTRAINT [PK__T__new] PRIMARY KEY CLUSTERED ([id] ASC)
);
`
	if err := os.WriteFile(abs, []byte(desiredSQL), 0o644); err != nil {
		t.Fatal(err)
	}

	e := New(config.Default(root))
	history := map[string]historyObject{
		objectKey("tables", "dbo.T"): {
			Category: "tables", Object: "dbo.T", SemanticSQL: previous,
			NormalizedHash: hashSemanticByCategory("tables", previous),
		},
	}
	desired := []desiredObject{{
		Category: "tables", Object: "dbo.T", SourceFile: filepath.ToSlash(sourceFile),
		SemanticSQL: desiredSQL, NormalizedHash: hashSemanticByCategory("tables", desiredSQL),
	}}

	count, files, err := e.reconcileDesiredTableConstraintNames(history, desired)
	if err != nil {
		t.Fatal(err)
	}
	if count != 1 || len(files) != 1 || files[0] != filepath.ToSlash(sourceFile) {
		t.Fatalf("reconcile result count=%d files=%#v", count, files)
	}
	raw, err := os.ReadFile(abs)
	if err != nil {
		t.Fatal(err)
	}
	got := string(raw)
	if !strings.Contains(got, "CONSTRAINT [PK__T__old]") || strings.Contains(got, "PK__T__new") {
		t.Fatalf("history name was not restored:\n%s", got)
	}
	if !strings.Contains(got, "[new_value] VARCHAR(50) NULL") {
		t.Fatalf("real table delta was lost:\n%s", got)
	}
}

func TestEquivalentConstraintNameMappingSkipsAmbiguousDuplicates(t *testing.T) {
	previous := `CREATE TABLE dbo.T (
  [id] INT NULL,
  CONSTRAINT [CK_T_old_a] CHECK ([id] > 0),
  CONSTRAINT [CK_T_old_b] CHECK ([id] > 0)
);
`
	desired := `CREATE TABLE dbo.T (
  [id] INT NULL,
  CONSTRAINT [CK_T_new_a] CHECK ([id] > 0),
  CONSTRAINT [CK_T_new_b] CHECK ([id] > 0)
);
`
	mapping, err := equivalentConstraintNameMapping(previous, desired)
	if err != nil {
		t.Fatal(err)
	}
	if len(mapping) != 0 {
		t.Fatalf("ambiguous duplicate definitions must not be guessed: %#v", mapping)
	}
}

func TestEquivalentConstraintNameMappingSkipsCaseInsensitiveNameCollision(t *testing.T) {
	previous := `CREATE TABLE dbo.T (
  [id] INT NULL,
  [other] INT NULL,
  CONSTRAINT [CK_T_old] CHECK ([id] > 0)
);
`
	desired := `CREATE TABLE dbo.T (
  [id] INT NULL,
  [other] INT NULL,
  CONSTRAINT [CK_T_new] CHECK ([id] > 0),
  CONSTRAINT [ck_t_old] CHECK ([other] > 0)
);
`
	mapping, err := equivalentConstraintNameMapping(previous, desired)
	if err != nil {
		t.Fatal(err)
	}
	if len(mapping) != 0 {
		t.Fatalf("case-insensitive constraint namespace collision must not be auto-fixed: %#v", mapping)
	}
}

func TestEquivalentConstraintNameMappingRestoresCaseOnlyDrift(t *testing.T) {
	previous := `CREATE TABLE dbo.T ([id] INT NULL, CONSTRAINT [CK_T_value] CHECK ([id] > 0));`
	desired := `CREATE TABLE dbo.T ([id] INT NULL, CONSTRAINT [ck_t_value] CHECK ([id] > 0));`
	mapping, err := equivalentConstraintNameMapping(previous, desired)
	if err != nil {
		t.Fatal(err)
	}
	if mapping["ck_t_value"] != "CK_T_value" {
		t.Fatalf("case-only constraint drift should be restored: %#v", mapping)
	}
}

func TestEquivalentConstraintNameMappingHandlesMultipleInlineConstraints(t *testing.T) {
	previous := `CREATE TABLE dbo.T (
  [value] INT NULL CONSTRAINT [DF_T_value_old] DEFAULT ((0)) CONSTRAINT [CK_T_value_old] CHECK ([value] >= 0)
);
`
	desired := `CREATE TABLE dbo.T (
  [value] INT NULL CONSTRAINT [DF_T_value_new] DEFAULT ((0)) CONSTRAINT [CK_T_value_new] CHECK ([value] >= 0)
);
`
	mapping, err := equivalentConstraintNameMapping(previous, desired)
	if err != nil {
		t.Fatal(err)
	}
	if mapping["DF_T_value_new"] != "DF_T_value_old" || mapping["CK_T_value_new"] != "CK_T_value_old" {
		t.Fatalf("multiple inline constraint mapping mismatch: %#v", mapping)
	}

	rewritten, count, err := replaceNamedConstraintsInTableSource(desired, "dbo.T", mapping)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 || !strings.Contains(rewritten, "[DF_T_value_old]") || !strings.Contains(rewritten, "[CK_T_value_old]") {
		t.Fatalf("multiple inline constraints were not reconciled:\n%s", rewritten)
	}
}

func TestConstraintNameOccurrencesIgnoresLiteralsCommentsAndQuotedIdentifiers(t *testing.T) {
	fragment := "[CONSTRAINT fake] NVARCHAR(100) CONSTRAINT\n[DF_real] DEFAULT (N'CONSTRAINT fake') /* CONSTRAINT ignored */"
	got, err := constraintNameOccurrences(fragment)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || got[0].Name != "DF_real" {
		t.Fatalf("constraint occurrences = %#v, want only DF_real", got)
	}
}
