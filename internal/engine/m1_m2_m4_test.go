package engine

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/output"
	"rza_migrate_db/internal/verify"
)

func TestParseMigrationFileNameRejectsNonDBOSchema(t *testing.T) {
	t.Parallel()

	e := New(config.Default(t.TempDir()))
	_, issue := e.parseMigrationFileName("20260217073015000_tables_abc.TableA_v1.up.sql")
	if issue == nil {
		t.Fatalf("expected parse issue for non-dbo migration object token")
	}
	if issue.Rule != "MIGRATION_OBJECT_SCHEMA_NOT_DBO" {
		t.Fatalf("unexpected rule: %s", issue.Rule)
	}
}

func TestVerifyIncludesReconstructionUnsupportedSubset(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	mustWrite := func(rel, content string) {
		t.Helper()
		path := filepath.Join(root, filepath.FromSlash(rel))
		if err := osMkdirAll(filepath.Dir(path)); err != nil {
			t.Fatalf("mkdir failed: %v", err)
		}
		if err := osWriteFile(path, content); err != nil {
			t.Fatalf("write failed: %v", err)
		}
	}

	mustWrite(".gitignore", "out/**\n*.output.json\n")
	mustWrite(".gitattributes", "* text=auto eol=lf\n*.sql text eol=lf\n*.md text eol=lf\n*.go text eol=lf\n")
	mustWrite("constraints_naming.md", "# constraints naming\n")

	mustWrite("migrations/000000__core__mig__audit_init.up.sql", `
IF OBJECT_ID(N'dbo.mig_event', N'U') IS NULL
BEGIN
  CREATE TABLE dbo.mig_event (
    event_id BIGINT IDENTITY(1,1) NOT NULL CONSTRAINT PK_mig_event PRIMARY KEY
  );
END
IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NULL
BEGIN
  EXEC (N'CREATE PROCEDURE dbo.usp_mig_log_event AS BEGIN SET NOCOUNT ON; END');
END
`)
	mustWrite("migrations/000000__core__mig__audit_init.down.sql", `
IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  DROP PROCEDURE dbo.usp_mig_log_event;
IF OBJECT_ID(N'dbo.mig_event', N'U') IS NOT NULL
  DROP TABLE dbo.mig_event;
`)
	mustWrite("migrations/20260217073015010_tables_dbo.Unsupported_v1.up.sql", `
CREATE TABLE dbo.Unsupported (
  id INT NOT NULL,
  payload VARBINARY(MAX) FILESTREAM NULL
);
`)
	mustWrite("migrations/20260217073015010_tables_dbo.Unsupported_v1.down.sql", `
DROP TABLE dbo.Unsupported;
`)

	e := New(config.Default(root))
	report, err := e.Run("verify")
	if err != nil {
		t.Fatalf("verify run failed: %v", err)
	}

	hasUnsupported := false
	for _, item := range report.Fail.Items {
		if item.Rule == "UNSUPPORTED_SQL_SUBSET" {
			hasUnsupported = true
			break
		}
	}
	if !hasUnsupported {
		t.Fatalf("expected UNSUPPORTED_SQL_SUBSET in verify fail items, got: %#v", report.Fail.Items)
	}
}

func TestHashSemanticByCategoryTableTreatsNullOmissionAsEquivalent(t *testing.T) {
	t.Parallel()

	withNull := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id)
);`
	withoutNull := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10),
  CONSTRAINT PK_TableA PRIMARY KEY (id)
);`

	h1 := hashSemanticByCategory("tables", withNull)
	h2 := hashSemanticByCategory("tables", withoutNull)
	if h1 != h2 {
		t.Fatalf("expected equal hash for NULL omission equivalence, got %s vs %s", h1, h2)
	}
}

func TestHashSemanticByCategoryIndexesTreatsOrderAsEquivalent(t *testing.T) {
	t.Parallel()

	firstOrder := `
CREATE NONCLUSTERED INDEX [IX_TableA_A] ON dbo.TableA ([ColA] ASC);
CREATE NONCLUSTERED INDEX [IX_TableA_B] ON dbo.TableA ([ColB] ASC) INCLUDE ([ColC]);
`
	secondOrder := `
CREATE NONCLUSTERED INDEX [IX_TableA_B] ON dbo.TableA ([ColB] ASC) INCLUDE ([ColC]);
CREATE NONCLUSTERED INDEX [IX_TableA_A] ON dbo.TableA ([ColA] ASC);
`

	h1 := hashSemanticByCategory("indexes", firstOrder)
	h2 := hashSemanticByCategory("indexes", secondOrder)
	if h1 != h2 {
		t.Fatalf("expected equal hash for index order equivalence, got %s vs %s", h1, h2)
	}
}

func TestCanonicalizeTableSnapshotFormattingLeadingComma(t *testing.T) {
	t.Parallel()

	raw := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL
  ,status VARCHAR(10) NULL
  ,CONSTRAINT PK_TableA PRIMARY KEY (id)
);
`
	got := canonicalizeTableSnapshotFormatting(raw, "dbo.TableA")
	if strings.Contains(got, "\n  ,") {
		t.Fatalf("expected trailing comma format, got: %s", got)
	}
	if !strings.Contains(got, "id INT NOT NULL,\n  status VARCHAR(10) NULL,\n  CONSTRAINT PK_TableA PRIMARY KEY (id)") {
		t.Fatalf("canonical format mismatch: %s", got)
	}
}

func TestCanonicalizeTableSnapshotFormattingTrailingComma(t *testing.T) {
	t.Parallel()

	raw := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL
);
`
	got := canonicalizeTableSnapshotFormatting(raw, "dbo.TableA")
	if strings.Contains(got, "\n  ,") {
		t.Fatalf("expected no leading comma in canonical output, got: %s", got)
	}
	if !strings.Contains(got, "id INT NOT NULL,\n  status VARCHAR(10) NULL") {
		t.Fatalf("canonical format mismatch: %s", got)
	}
}

func TestExtractCreateTableStatementBoundsRejectsCurlyBraces(t *testing.T) {
	t.Parallel()

	raw := `
CREATE TABLE dbo.TableA {
  col1 varchar(2),
  col2 varchar(2)
}
`
	_, _, _, err := extractCreateTableStatementBounds("dbo.TableA", raw)
	if err == nil {
		t.Fatalf("expected invalid CREATE TABLE error for curly braces")
	}
	if !strings.Contains(strings.ToLower(err.Error()), "invalid create table statement") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestExtractCreateTableStatementSupportsBracketedQualifiedObject(t *testing.T) {
	t.Parallel()

	raw := `
CREATE TABLE [dbo].[TableA](
  [id] INT NOT NULL
)
`
	stmt, err := extractCreateTableStatement("dbo.TableA", raw)
	if err != nil {
		t.Fatalf("extractCreateTableStatement failed: %v", err)
	}
	if !strings.Contains(stmt, "[dbo].[TableA]") {
		t.Fatalf("expected bracketed table statement, got: %s", stmt)
	}
}

func TestExtractCreateTableItemsFailsOnMissingCommaBetweenColumns(t *testing.T) {
	t.Parallel()

	stmt := `
CREATE TABLE dbo.TableA (
  [id] INT NOT NULL
  [name] VARCHAR(50) NULL
);`

	_, err := extractCreateTableItems(stmt)
	if err == nil {
		t.Fatalf("expected missing comma error")
	}
	if !strings.Contains(strings.ToLower(err.Error()), "missing comma between table items") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestExtractCreateTableRequiresExactObjectCase(t *testing.T) {
	t.Parallel()

	sql := `CREATE TABLE [dbo].[CaseTable] ([Code] INT NOT NULL);`
	if _, err := extractCreateTableStatement("dbo.CaseTable", sql); err != nil {
		t.Fatalf("exact object spelling should match: %v", err)
	}
	if _, err := extractCreateTableStatement("dbo.casetable", sql); err == nil {
		t.Fatal("case-mismatched object must not bind to the source statement")
	}
}

func TestExtractCreateTableItemsFailsOnMissingCommaBeforeConstraint(t *testing.T) {
	t.Parallel()

	stmt := `
CREATE TABLE dbo.TableA (
  [id] INT NOT NULL
  CONSTRAINT [PK_TableA] PRIMARY KEY ([id])
);`

	_, err := extractCreateTableItems(stmt)
	if err == nil {
		t.Fatalf("expected missing comma error")
	}
	if !strings.Contains(strings.ToLower(err.Error()), "missing comma between table items") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSemanticFromSourceCanonicalizesTableColumnsAndConstraintPosition(t *testing.T) {
	t.Parallel()

	meta := migrationMeta{
		Category: "tables",
		Object:   "dbo.control_table",
	}
	src := `
CREATE TABLE dbo.control_table (
  id INT NOT NULL,
  columnBaru varchar(11),
  CONSTRAINT [DF_control_table_columnBaru] DEFAULT ('-') FOR [columnBaru],
  columnBaru3 varchar(24),
  CONSTRAINT [PK_control_table] PRIMARY KEY CLUSTERED ([id] ASC)
);`

	got, err := semanticFromSource(meta, src)
	if err != nil {
		t.Fatalf("semanticFromSource failed: %v", err)
	}

	if !strings.Contains(got, "[columnBaru] varchar(11) NULL") {
		t.Fatalf("expected bracket + NULL canonicalization, got: %s", got)
	}
	if !strings.Contains(got, "[columnBaru3] varchar(24) NULL") {
		t.Fatalf("expected bracket + NULL canonicalization for second column, got: %s", got)
	}

	idxCol3 := strings.Index(got, "[columnBaru3] varchar(24) NULL")
	idxDf := strings.Index(got, "CONSTRAINT [DF_control_table_columnBaru] DEFAULT ('-') FOR [columnBaru]")
	if idxCol3 < 0 || idxDf < 0 {
		t.Fatalf("expected both columnBaru3 and default constraint in output, got: %s", got)
	}
	if idxDf < idxCol3 {
		t.Fatalf("expected default constraint placed after column list, got: %s", got)
	}
}

func TestCanonicalizeTableColumnItemStyleHandlesBracketNameWithSpace(t *testing.T) {
	t.Parallel()

	in := "[Transaction Date] datetime"
	got := canonicalizeTableColumnItemStyle(in)
	if !strings.Contains(got, "[Transaction Date] datetime NULL") {
		t.Fatalf("unexpected canonicalized column item: %s", got)
	}
}

func TestCanonicalizeTableColumnItemStylePreservesSingleSpaceIdentifier(t *testing.T) {
	t.Parallel()

	got := canonicalizeTableColumnItemStyle("[ ] NVARCHAR(255) NULL")
	if got != "[ ] NVARCHAR(255) NULL" {
		t.Fatalf("single-space identifier must be preserved, got: %q", got)
	}
}

func TestCanonicalizeTableColumnItemStyleDoesNotDuplicateNullBeforeComma(t *testing.T) {
	t.Parallel()

	got := canonicalizeTableColumnItemStyle("[nullable_col] VARCHAR(100) null,")
	if got != "[nullable_col] VARCHAR(100) null," {
		t.Fatalf("explicit NULL before comma must not be duplicated, got: %q", got)
	}
}

func TestCanonicalizeTableColumnItemStylePreservesComputedColumn(t *testing.T) {
	t.Parallel()

	got := canonicalizeTableColumnItemStyle("[NameLower] AS LOWER(Name) PERSISTED")
	if got != "[NameLower] AS LOWER(Name) PERSISTED" {
		t.Fatalf("computed column must not receive NULL/NOT NULL suffix, got: %s", got)
	}
}

func TestSemanticFromSourcePreservesCommentsBeforeTableColumns(t *testing.T) {
	t.Parallel()

	meta := migrationMeta{Category: "tables", Object: "dbo.TableA"}
	source := `CREATE TABLE dbo.TableA(
--column untuk id
id int,
--column untuk nama
nama varchar(100)
);`

	got, err := semanticFromSource(meta, source)
	if err != nil {
		t.Fatalf("semanticFromSource failed: %v", err)
	}
	for _, want := range []string{
		"--column untuk id\n  [id] int NULL",
		"--column untuk nama\n  [nama] varchar(100) NULL",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("expected comment-bearing column %q:\n%s", want, got)
		}
	}
	if strings.Contains(got, "[--]") {
		t.Fatalf("comment marker must not become a column identifier:\n%s", got)
	}
}

func TestSemanticFromSourceSupportsMultipleAndBlockCommentsBeforeColumns(t *testing.T) {
	t.Parallel()

	meta := migrationMeta{Category: "tables", Object: "dbo.TableA"}
	source := `CREATE TABLE dbo.TableA (
  -- first line
  -- second line mentions NULL and varchar(100)
  id INT NOT NULL,
  /* customer
     display name */
  nama VARCHAR(100),
  amount AS (id * 10) PERSISTED,
  CONSTRAINT [PK_TableA] PRIMARY KEY ([id])
);`

	got, err := semanticFromSource(meta, source)
	if err != nil {
		t.Fatalf("semanticFromSource failed: %v", err)
	}
	for _, want := range []string{
		"-- first line",
		"-- second line mentions NULL and varchar(100)",
		"[id] INT NOT NULL",
		"/* customer",
		"display name */",
		"[nama] VARCHAR(100) NULL",
		"[amount] AS (id * 10) PERSISTED",
		"CONSTRAINT [PK_TableA] PRIMARY KEY ([id])",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("expected %q in canonical table:\n%s", want, got)
		}
	}
	if strings.Contains(got, "[--]") || strings.Contains(got, "[/*]") {
		t.Fatalf("SQL comments must never become column identifiers:\n%s", got)
	}
}

func TestSemanticFromSourcePreservesCommentAfterPreviousComma(t *testing.T) {
	t.Parallel()

	meta := migrationMeta{Category: "tables", Object: "dbo.TableA"}
	source := `CREATE TABLE dbo.TableA (
  id INT NOT NULL, -- comment attached to the next item by comma splitting
  nama VARCHAR(100)
);`

	got, err := semanticFromSource(meta, source)
	if err != nil {
		t.Fatalf("semanticFromSource failed: %v", err)
	}
	if !strings.Contains(got, "-- comment attached to the next item by comma splitting") ||
		!strings.Contains(got, "[nama] VARCHAR(100) NULL") ||
		strings.Contains(got, "[--]") {
		t.Fatalf("comment after a comma was not preserved safely:\n%s", got)
	}
}

func TestSemanticFromSourceDoesNotAppendNullToComputedColumn(t *testing.T) {
	t.Parallel()

	meta := migrationMeta{Category: "tables", Object: "dbo.TableA"}
	source := `CREATE TABLE dbo.TableA (
  ID INT NOT NULL,
  NameLower AS LOWER(Name) PERSISTED
);`
	semantic, err := semanticFromSource(meta, source)
	if err != nil {
		t.Fatalf("semanticFromSource failed: %v", err)
	}
	upper := strings.ToUpper(semantic)
	if !strings.Contains(upper, "[NAMELOWER] AS LOWER(NAME) PERSISTED") || strings.Contains(upper, "PERSISTED NULL") {
		t.Fatalf("computed column semantic is invalid: %s", semantic)
	}
}

func TestValidateTableSubsetSupportRejectsPersistedNullComputedColumn(t *testing.T) {
	t.Parallel()

	stmt := `CREATE TABLE dbo.TableA (
  ID INT NOT NULL,
  NameLower AS LOWER(Name) PERSISTED NULL
);`
	if err := validateTableSubsetSupport(stmt); err == nil || !strings.Contains(err.Error(), "PERSISTED NULL") {
		t.Fatalf("expected PERSISTED NULL validation error, got: %v", err)
	}
}

func TestBuildAlterColumnGuardedSQLUsesDetailVariableForPrecheck(t *testing.T) {
	t.Parallel()

	previous := tableColumnDef{Name: "Series", BaseType: "VARCHAR", RawType: "VARCHAR(255)", Length: 255, Nullable: true}
	next := tableColumnDef{Name: "Series", BaseType: "VARCHAR", RawType: "VARCHAR(100)", Length: 100, Nullable: true}
	sql := buildAlterColumnGuardedSQL("dbo.TableA", previous, next)
	if strings.Contains(sql, "@detail = CONCAT(") {
		t.Fatalf("procedure parameter must not receive CONCAT expression: %s", sql)
	}
	suffix := safeSQLVarSuffix("Series")
	if !strings.Contains(sql, "DECLARE @detail_too_long_"+suffix+" NVARCHAR(128);") ||
		!strings.Contains(sql, "@detail = @detail_too_long_"+suffix) {
		t.Fatalf("length precheck must assign and pass a detail variable: %s", sql)
	}
}

func TestSemanticFromSourceTreatsNonBreakingSpaceAsColumnTypeSeparator(t *testing.T) {
	t.Parallel()

	meta := migrationMeta{Category: "tables", Object: "dbo.TableA"}
	source := "CREATE TABLE dbo.TableA (\n  tax_interest\u00A0\u00A0money NOT NULL\n);\n"
	semantic, err := semanticFromSource(meta, source)
	if err != nil {
		t.Fatalf("semanticFromSource failed: %v", err)
	}
	upper := strings.ToUpper(semantic)
	if !strings.Contains(upper, "[TAX_INTEREST] MONEY NOT NULL") || strings.Contains(upper, "TAX_INTEREST MONEY]") {
		t.Fatalf("NBSP separator must produce a valid column definition: %s", semantic)
	}
}

func TestExtractCreateTableItemsFailsOnTrailingComma(t *testing.T) {
	t.Parallel()

	stmt := `
CREATE TABLE dbo.TableA (
  [id] INT NOT NULL,
  [name] VARCHAR(50) NULL,
);`

	_, err := extractCreateTableItems(stmt)
	if err == nil {
		t.Fatalf("expected trailing comma error")
	}
	if !strings.Contains(strings.ToLower(err.Error()), "trailing comma") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRepairCreateTableCommaSyntaxHandlesMissingDuplicateAndTrailing(t *testing.T) {
	t.Parallel()

	input := `CREATE TABLE dbo.TableA (
  [Code] NVARCHAR(30) NOT NULL
  [code] NVARCHAR(30) NULL,,
  [Marker] NVARCHAR(10) CONSTRAINT [DF_TableA_Marker] DEFAULT (N')') NULL,
  CONSTRAINT [UQ_TableA_Code] UNIQUE ([Code]),
);`
	repaired, changed := repairCreateTableCommaSyntax(input)
	if !changed {
		t.Fatal("expected safe comma recovery")
	}
	items, err := extractCreateTableItems(repaired)
	if err != nil {
		t.Fatalf("repaired CREATE TABLE is invalid: %v\n%s", err, repaired)
	}
	if len(items) != 4 {
		t.Fatalf("expected four exact-case items, got %d: %#v", len(items), items)
	}
	if !strings.Contains(repaired, "[Code] NVARCHAR(30) NOT NULL,") ||
		!strings.Contains(repaired, "DEFAULT (N')') NULL,") {
		t.Fatalf("repair damaged identifiers or parenthesis literal:\n%s", repaired)
	}
}

func TestRepairCreateTableCommaSyntaxDoesNotTreatColumnRemarkAsItem(t *testing.T) {
	t.Parallel()

	input := `CREATE TABLE dbo.TableA (
  -- field info
  fieldA VARCHAR(100) NULL
);`
	repaired, changed := repairCreateTableCommaSyntax(input)
	if changed {
		t.Fatalf("valid comment-bearing table must not be repaired:\n%s", repaired)
	}
	if repaired != input {
		t.Fatalf("column remark placement changed\nwant:\n%s\ngot:\n%s", input, repaired)
	}
}

func TestParseTableColumnsFailsOnDuplicateColumnName(t *testing.T) {
	t.Parallel()

	stmt := `
CREATE TABLE dbo.TableA (
  [id] INT NOT NULL,
  [id] VARCHAR(10) NULL
);`

	_, err := parseTableColumns(stmt)
	if err == nil {
		t.Fatalf("expected duplicate column definition error")
	}
	if !strings.Contains(strings.ToLower(err.Error()), "duplicate column definition") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestExtractNamedConstraintDefinitionMapFailsOnDuplicateConstraintName(t *testing.T) {
	t.Parallel()

	stmt := `
CREATE TABLE dbo.TableA (
  [id] INT NOT NULL,
  CONSTRAINT [DF_TableA_id] DEFAULT (0) FOR [id],
  CONSTRAINT [DF_TableA_id] DEFAULT (1) FOR [id]
);`

	_, err := extractNamedConstraintDefinitionMap(stmt)
	if err == nil {
		t.Fatalf("expected duplicate constraint definition error")
	}
	if !strings.Contains(strings.ToLower(err.Error()), "duplicate constraint definition") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateIndexSourceTargetsFailsOnWrongTargetTable(t *testing.T) {
	t.Parallel()

	sql := `
CREATE NONCLUSTERED INDEX IX_TableB_id ON dbo.TableB(id);
`
	err := validateIndexSourceTargets("dbo.TableA", sql)
	if err == nil {
		t.Fatalf("expected target mismatch error")
	}
	if !strings.Contains(strings.ToLower(err.Error()), "targets dbo.tableb") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateIndexSourceTargetsFailsOnMissingOnTarget(t *testing.T) {
	t.Parallel()

	sql := `
CREATE NONCLUSTERED INDEX IX_TableA_id (id);
`
	err := validateIndexSourceTargets("dbo.TableA", sql)
	if err == nil {
		t.Fatalf("expected missing ON target error")
	}
	if !strings.Contains(strings.ToLower(err.Error()), "missing on target") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCanonicalizeTableSnapshotFormattingIsIdempotentWithTrailingBlankLines(t *testing.T) {
	t.Parallel()

	raw := "CREATE TABLE dbo.TableA (\n  id INT NOT NULL\n);\n\n\n\n"
	first := canonicalizeTableSnapshotFormatting(raw, "dbo.TableA")
	second := canonicalizeTableSnapshotFormatting(first, "dbo.TableA")
	if first != second {
		t.Fatalf("expected idempotent canonicalization, got diff\nfirst=%q\nsecond=%q", first, second)
	}
	if !strings.HasSuffix(first, "\n") {
		t.Fatalf("expected canonicalized SQL to end with newline")
	}
	if strings.HasSuffix(first, "\n\n") {
		t.Fatalf("expected canonicalized SQL to avoid extra blank lines at end: %q", first)
	}
}

func TestValidateWrapperDeterministicMetadata(t *testing.T) {
	t.Parallel()

	report := output.NewReport("pattern", "LF")
	meta := verify.MigrationMeta{
		FileName: "20260217073015000_tables_dbo.TableA_v1.up.sql",
		RelPath:  "migrations/20260217073015000_tables_dbo.TableA_v1.up.sql",
		Version:  "20260217073015000",
		Category: "tables",
	}

	sqlText := `
DECLARE @version BIGINT = 20260217073015001;
DECLARE @category NVARCHAR(16) = N'indexes';
DECLARE @file NVARCHAR(260) = N'20260217073015000_tables_dbo.TableA_v1.up.sql';
`
	semanticUpper := strings.ToUpper(stripSQLStringsAndComments(sqlText))
	verify.ValidateWrapperDeterministicMetadata(&report, meta, sqlText, semanticUpper)
	if report.Fail.Count == 0 {
		t.Fatalf("expected metadata mismatch failures")
	}
}

func TestValidateWrapperMetadataRequiresExactCaseAndTarget(t *testing.T) {
	t.Parallel()

	report := output.NewReport("pattern", "LF")
	meta := verify.MigrationMeta{
		FileName: "20260217073015000_tables_dbo.CaseTable_v1.up.sql",
		RelPath:  "migrations/20260217073015000_tables_dbo.CaseTable_v1.up.sql",
		Version:  "20260217073015000",
		Category: "tables",
		Object:   "dbo.CaseTable",
	}
	sqlText := `
DECLARE @version BIGINT = 20260217073015000;
DECLARE @category NVARCHAR(16) = N'TABLES';
DECLARE @file NVARCHAR(260) = N'20260217073015000_tables_dbo.casetable_v1.up.sql';
DECLARE @target NVARCHAR(256) = N'dbo.casetable';
`
	semanticUpper := strings.ToUpper(stripSQLStringsAndComments(sqlText))
	verify.ValidateWrapperDeterministicMetadata(&report, meta, sqlText, semanticUpper)
	if report.Fail.Count != 3 {
		t.Fatalf("expected exact-case failures for category/file/target, got %d: %+v", report.Fail.Count, report.Fail.Items)
	}
}

func TestAllocateVersionHonorsMaxPairsPerRun(t *testing.T) {
	t.Parallel()

	cfg := config.Default(t.TempDir())
	cfg.MaxPairsPerRun = 3
	e := New(cfg)
	op := plannedMigration{Category: "tables", Object: "dbo.TableA", VK: 1}

	if _, _, err := e.allocateVersion("20260217073015", 3, map[string]struct{}{}, op); err == nil {
		t.Fatalf("expected sequence overflow at max boundary")
	}
	if _, _, err := e.allocateVersion("20260217073015", 2, map[string]struct{}{}, op); err != nil {
		t.Fatalf("unexpected allocation error before max boundary: %v", err)
	}
}

func TestAllocateMigrationBatchReturnsPartialWhenPairLimitReached(t *testing.T) {
	t.Parallel()

	cfg := config.Default(t.TempDir())
	cfg.MaxPairsPerRun = 2
	e := New(cfg)
	reserved := map[string]struct{}{}

	ops := []plannedMigration{
		{Category: "tables", Object: "dbo.TableA", VK: 1},
		{Category: "tables", Object: "dbo.TableB", VK: 1},
		{Category: "tables", Object: "dbo.TableC", VK: 1},
	}

	allocated, pending, _, err := e.allocateMigrationBatch(ops, "20260217073015", 0, reserved)
	if err != nil {
		t.Fatalf("unexpected allocation error: %v", err)
	}
	if len(allocated) != 2 {
		t.Fatalf("expected 2 allocated pairs, got %d", len(allocated))
	}
	if pending != 1 {
		t.Fatalf("expected 1 pending pair, got %d", pending)
	}
}

func TestBuildTableAlterColumnsSQLAddNotNullWithDefaultConstraint(t *testing.T) {
	t.Parallel()

	prev := `CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id)
);`
	next := `CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  amount DECIMAL(18,2) NOT NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id),
  CONSTRAINT DF_TableA__amount DEFAULT ((0)) FOR amount
);`

	changes, err := classifyTableColumnChanges(prev, next)
	if err != nil {
		t.Fatalf("classify changes failed: %v", err)
	}
	upSQL, downSQL, hasDestructiveDown, err := buildTableAlterColumnsSQL(prev, next, "dbo.TableA", changes)
	if err != nil {
		t.Fatalf("build alter sql failed: %v", err)
	}
	if !hasDestructiveDown {
		t.Fatalf("expected destructive down for added column")
	}

	upUpper := strings.ToUpper(upSQL)
	if !strings.Contains(upUpper, "ALTER TABLE DBO.TABLEA ADD [AMOUNT] DECIMAL(18,2) NOT NULL CONSTRAINT [DF_TABLEA__AMOUNT] DEFAULT ((0)) WITH VALUES;") {
		t.Fatalf("unexpected up SQL: %s", upSQL)
	}
	downUpper := strings.ToUpper(downSQL)
	if !strings.Contains(downUpper, "ALTER TABLE DBO.TABLEA DROP CONSTRAINT [DF_TABLEA__AMOUNT];") {
		t.Fatalf("expected down SQL to drop default constraint before column: %s", downSQL)
	}
	if !strings.Contains(downUpper, "ALTER TABLE DBO.TABLEA DROP COLUMN [AMOUNT];") {
		t.Fatalf("unexpected down SQL: %s", downSQL)
	}
	if strings.Index(downUpper, "DROP CONSTRAINT [DF_TABLEA__AMOUNT]") > strings.Index(downUpper, "DROP COLUMN [AMOUNT]") {
		t.Fatalf("expected default constraint drop before column drop: %s", downSQL)
	}
}

func TestBuildTableAlterColumnsSQLAddNotNullWithoutDefaultConstraint(t *testing.T) {
	t.Parallel()

	prev := `CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id)
);`
	next := `CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  amount DECIMAL(18,2) NOT NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id)
);`

	changes, err := classifyTableColumnChanges(prev, next)
	if err != nil {
		t.Fatalf("classify changes failed: %v", err)
	}
	_, _, _, err = buildTableAlterColumnsSQL(prev, next, "dbo.TableA", changes)
	if err == nil {
		t.Fatalf("expected manual-required error for NOT NULL add without default")
	}
	if !strings.Contains(strings.ToUpper(err.Error()), "REQUIRES DEFAULT CONSTRAINT") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestBuildTableAlterColumnsSQLAddNotNullWithInlineDefaultConstraint(t *testing.T) {
	t.Parallel()

	prev := `CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id)
);`
	next := `CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NOT NULL CONSTRAINT [DF_TableA__status] DEFAULT ('NEW'),
  CONSTRAINT PK_TableA PRIMARY KEY (id)
);`

	changes, err := classifyTableColumnChanges(prev, next)
	if err != nil {
		t.Fatalf("classify changes failed: %v", err)
	}
	upSQL, _, _, err := buildTableAlterColumnsSQL(prev, next, "dbo.TableA", changes)
	if err != nil {
		t.Fatalf("build alter sql failed: %v", err)
	}
	upUpper := strings.ToUpper(upSQL)
	if !strings.Contains(upUpper, "ALTER TABLE DBO.TABLEA ADD [STATUS] VARCHAR(10) NOT NULL CONSTRAINT [DF_TABLEA__STATUS] DEFAULT ('NEW') WITH VALUES;") {
		t.Fatalf("unexpected up SQL: %s", upSQL)
	}
}

func TestBuildTableAlterColumnsSQLAddNullableWithInlineDefaultDropsDefaultOnDown(t *testing.T) {
	t.Parallel()

	prev := `CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id)
);`
	next := `CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL CONSTRAINT [DF_TableA__status] DEFAULT ('NEW'),
  CONSTRAINT PK_TableA PRIMARY KEY (id)
);`

	changes, err := classifyTableColumnChanges(prev, next)
	if err != nil {
		t.Fatalf("classify changes failed: %v", err)
	}
	upSQL, downSQL, hasDestructiveDown, err := buildTableAlterColumnsSQL(prev, next, "dbo.TableA", changes)
	if err != nil {
		t.Fatalf("build alter sql failed: %v", err)
	}
	if !hasDestructiveDown {
		t.Fatalf("expected destructive down for added column")
	}

	upUpper := strings.ToUpper(upSQL)
	if !strings.Contains(upUpper, "ALTER TABLE DBO.TABLEA ADD [STATUS] VARCHAR(10) NULL;") {
		t.Fatalf("unexpected up SQL: %s", upSQL)
	}
	if !strings.Contains(upUpper, "ALTER TABLE DBO.TABLEA ADD CONSTRAINT [DF_TABLEA__STATUS] DEFAULT ('NEW') FOR [STATUS];") {
		t.Fatalf("expected separate default add in up SQL: %s", upSQL)
	}
	downUpper := strings.ToUpper(downSQL)
	if !strings.Contains(downUpper, "ALTER TABLE DBO.TABLEA DROP CONSTRAINT [DF_TABLEA__STATUS];") {
		t.Fatalf("expected down SQL to drop default constraint before column: %s", downSQL)
	}
	if !strings.Contains(downUpper, "ALTER TABLE DBO.TABLEA DROP COLUMN [STATUS];") {
		t.Fatalf("unexpected down SQL: %s", downSQL)
	}
	if strings.Index(downUpper, "DROP CONSTRAINT [DF_TABLEA__STATUS]") > strings.Index(downUpper, "DROP COLUMN [STATUS]") {
		t.Fatalf("expected default constraint drop before column drop: %s", downSQL)
	}
}

func TestParseTableColumnsSupportsBracketedSpaceColumnName(t *testing.T) {
	t.Parallel()

	stmt := `CREATE TABLE dbo.CTPLIST (
  [ ] NVARCHAR(255) COLLATE Latin1_General_BIN NULL,
  [M_NAME] NVARCHAR(255) COLLATE Latin1_General_BIN NULL
);`

	cols, err := parseTableColumns(stmt)
	if err != nil {
		t.Fatalf("parseTableColumns failed: %v", err)
	}
	def, ok := cols[" "]
	if !ok {
		t.Fatalf("expected single-space column key, got keys: %#v", cols)
	}
	if def.Name != " " || def.RawType != "NVARCHAR(255) COLLATE Latin1_General_BIN" || !def.Nullable {
		t.Fatalf("unexpected parsed column definition: %#v", def)
	}

	changes := tableColumnChanges{Added: []string{" "}}
	upSQL, _, _, err := buildTableAlterColumnsSQL("CREATE TABLE dbo.CTPLIST ([M_NAME] NVARCHAR(255) NULL);", stmt, "dbo.CTPLIST", changes)
	if err != nil {
		t.Fatalf("buildTableAlterColumnsSQL failed: %v", err)
	}
	if !strings.Contains(upSQL, "ALTER TABLE dbo.CTPLIST ADD [ ] NVARCHAR(255) COLLATE Latin1_General_BIN NULL;") {
		t.Fatalf("expected bracketed space column add, got: %s", upSQL)
	}
}

func TestBuildTableAlterColumnsSQLNumericPrecisionChangeHasPrecheck(t *testing.T) {
	t.Parallel()

	prev := `CREATE TABLE dbo.TableA (
  ID INT NOT NULL,
  Amount DECIMAL(18,2) NOT NULL CONSTRAINT DF_TableA_Amount DEFAULT ((0))
);`
	next := `CREATE TABLE dbo.TableA (
  ID INT NOT NULL,
  Amount DECIMAL(19,4) NOT NULL CONSTRAINT DF_TableA_Amount DEFAULT ((0))
);`
	changes, err := classifyTableColumnChanges(prev, next)
	if err != nil {
		t.Fatalf("classify changes failed: %v", err)
	}

	upSQL, _, _, err := buildTableAlterColumnsSQL(prev, next, "dbo.TableA", changes)
	if err != nil {
		t.Fatalf("buildTableAlterColumnsSQL failed: %v", err)
	}
	if !strings.Contains(upSQL, "TRY_CONVERT(DECIMAL(19,4), [Amount])") {
		t.Fatalf("expected TRY_CONVERT precheck for decimal precision change, got:\n%s", upSQL)
	}
	if !strings.Contains(upSQL, "SET @skip_alter_"+safeSQLVarSuffix("Amount")+" = 1") {
		t.Fatalf("expected numeric precision precheck to set skip variable, got:\n%s", upSQL)
	}
}

func TestColumnTypeParametersWithWhitespaceRemainSemantic(t *testing.T) {
	t.Parallel()

	prev := `CREATE TABLE dbo.TableA (
  Amount DECIMAL(17, 5) NULL,
  Code VARCHAR (10) NULL,
  Payload VARBINARY (8) NULL,
  OccurredAt DATETIME2 (2) NULL
);`
	next := `CREATE TABLE dbo.TableA (
  Amount DECIMAL(17, 9) NULL,
  Code VARCHAR (24) NULL,
  Payload VARBINARY (16) NULL,
  OccurredAt DATETIME2 (5) NULL
);`

	changes, err := classifyTableColumnChanges(prev, next)
	if err != nil {
		t.Fatalf("classify changes failed: %v", err)
	}
	if got, want := strings.Join(changes.Changed, ","), "Amount,Code,OccurredAt,Payload"; got != want {
		t.Fatalf("parameterized type changes mismatch: got=%q want=%q", got, want)
	}
	if hashSemanticByCategory("tables", prev) == hashSemanticByCategory("tables", next) {
		t.Fatal("precision, scale, and length changes must change the table semantic hash")
	}
	spacingOnly := `CREATE TABLE dbo.TableA (
  Amount DECIMAL (17,9) NULL,
  Code VARCHAR( 24 ) NULL,
  Payload VARBINARY(16) NULL,
  OccurredAt DATETIME2( 5 ) NULL
);`
	if hashSemanticByCategory("tables", next) != hashSemanticByCategory("tables", spacingOnly) {
		t.Fatal("whitespace around type parameters must not change the table semantic hash")
	}

	cols, err := parseTableColumns(next)
	if err != nil {
		t.Fatalf("parse next table failed: %v", err)
	}
	for name, want := range map[string]string{
		"Amount":     "DECIMAL(17,9)",
		"Code":       "VARCHAR(24)",
		"Payload":    "VARBINARY(16)",
		"OccurredAt": "DATETIME2(5)",
	} {
		if got := cols[name].RawType; got != want {
			t.Errorf("%s raw type mismatch: got=%q want=%q", name, got, want)
		}
	}
}

func TestBuildTableAlterColumnsSQLDateTimePrecisionChangeHasPrecheck(t *testing.T) {
	t.Parallel()

	prev := `CREATE TABLE dbo.TableA (
  ID INT NOT NULL,
  LastCheckedAt DATETIME2(0) NULL
);`
	next := `CREATE TABLE dbo.TableA (
  ID INT NOT NULL,
  LastCheckedAt DATETIME2(3) NULL
);`
	changes, err := classifyTableColumnChanges(prev, next)
	if err != nil {
		t.Fatalf("classify changes failed: %v", err)
	}

	upSQL, _, _, err := buildTableAlterColumnsSQL(prev, next, "dbo.TableA", changes)
	if err != nil {
		t.Fatalf("buildTableAlterColumnsSQL failed: %v", err)
	}
	if !strings.Contains(upSQL, "TRY_CONVERT(DATETIME2(3), [LastCheckedAt])") {
		t.Fatalf("expected TRY_CONVERT precheck for datetime precision change, got:\n%s", upSQL)
	}
	if !strings.Contains(upSQL, "SET @skip_alter_"+safeSQLVarSuffix("LastCheckedAt")+" = 1") {
		t.Fatalf("expected datetime precision precheck to set skip variable, got:\n%s", upSQL)
	}
}

func TestAlterColumnDependencyRestoreUsesNextCheckDefinitionForUp(t *testing.T) {
	t.Parallel()

	prev := `CREATE TABLE dbo.TableA (
  ID INT NOT NULL,
  Amount DECIMAL(18,2) NOT NULL,
  CONSTRAINT CK_TableA_Amount CHECK (Amount >= 0 AND Amount <= 999999.99)
);`
	next := `CREATE TABLE dbo.TableA (
  ID INT NOT NULL,
  Amount DECIMAL(19,4) NOT NULL,
  CONSTRAINT CK_TableA_Amount CHECK (Amount >= 0 AND Amount <= 9999999.9999)
);`

	depPlan, err := buildDropColumnDependencyPlan("dbo.TableA", []string{"Amount"}, prev, map[string]historyObject{})
	if err != nil {
		t.Fatalf("buildDropColumnDependencyPlan failed: %v", err)
	}
	depPlan.Constraints = filterAutoManagedAlterColumnConstraints(depPlan.Constraints, "dbo.TableA")
	restoreUp, err := buildReAddDroppedDependenciesSQL(next, "dbo.TableA", depPlan, map[string]historyObject{})
	if err != nil {
		t.Fatalf("buildReAddDroppedDependenciesSQL failed: %v", err)
	}
	if !strings.Contains(restoreUp, "9999999.9999") {
		t.Fatalf("expected up restore to use next check definition, got:\n%s", restoreUp)
	}
	if strings.Contains(restoreUp, "Amount <= 999999.99)") {
		t.Fatalf("up restore must not use previous check definition, got:\n%s", restoreUp)
	}
}

func TestValidateTableSubsetSupportAcceptsComputedAndRejectsTemporal(t *testing.T) {
	t.Parallel()

	computed := `CREATE TABLE dbo.TableA (
  ID INT NOT NULL,
  Amount DECIMAL(19,4) NOT NULL,
  AmountText AS CONVERT(VARCHAR(32), Amount) PERSISTED,
  CONSTRAINT PK_TableA PRIMARY KEY (ID)
);`
	if err := validateTableSubsetSupport(computed); err != nil {
		t.Fatalf("expected computed column to be supported: %v", err)
	}

	temporal := `CREATE TABLE dbo.TableA (
  ID INT NOT NULL,
  ValidFrom DATETIME2 GENERATED ALWAYS AS ROW START NOT NULL,
  ValidTo DATETIME2 GENERATED ALWAYS AS ROW END NOT NULL,
  PERIOD FOR SYSTEM_TIME (ValidFrom, ValidTo),
  CONSTRAINT PK_TableA PRIMARY KEY (ID)
);`
	if err := validateTableSubsetSupport(temporal); err == nil {
		t.Fatalf("expected unsupported temporal table subset")
	}
}

func TestBuildTableAlterColumnsSQLRecreatesComputedColumn(t *testing.T) {
	t.Parallel()

	previous := `CREATE TABLE dbo.TableA (
  ID INT NOT NULL,
  NameLower AS LOWER(Name) PERSISTED
);`
	next := `CREATE TABLE dbo.TableA (
  ID INT NOT NULL,
  NameLower AS LOWER(LTRIM(Name)) PERSISTED
);`
	changes, err := classifyTableColumnChanges(previous, next)
	if err != nil {
		t.Fatalf("classify changes failed: %v", err)
	}
	if len(changes.Changed) != 1 || changes.Changed[0] != "NameLower" {
		t.Fatalf("expected computed column change, got %+v", changes)
	}
	upSQL, downSQL, _, err := buildTableAlterColumnsSQL(previous, next, "dbo.TableA", changes)
	if err != nil {
		t.Fatalf("buildTableAlterColumnsSQL failed: %v", err)
	}
	upUpper := strings.ToUpper(upSQL)
	if !strings.Contains(upUpper, "DROP COLUMN [NAMELOWER]") || !strings.Contains(upUpper, "ADD [NAMELOWER] AS LOWER(LTRIM(NAME)) PERSISTED") {
		t.Fatalf("computed column change must drop and recreate column, got:\n%s", upSQL)
	}
	if strings.Contains(upUpper, "ALTER COLUMN [NAMELOWER]") {
		t.Fatalf("computed column change must not use ALTER COLUMN, got:\n%s", upSQL)
	}
	if !strings.Contains(strings.ToUpper(downSQL), "ADD [NAMELOWER] AS LOWER(NAME) PERSISTED") {
		t.Fatalf("computed column down migration must restore prior definition, got:\n%s", downSQL)
	}
}

func TestBuildTableAlterColumnsSQLAddMixedCaseNotNullWithInlineDefaultConstraint(t *testing.T) {
	t.Parallel()

	prev := `CREATE TABLE dbo.TableA (
  [ID] INT NOT NULL,
  CONSTRAINT [PK_TableA] PRIMARY KEY ([ID])
);`
	next := `CREATE TABLE dbo.TableA (
  [ID] INT NOT NULL,
  [CreatedAt] DATETIME2(0) CONSTRAINT [DF_TableA_CreatedAt] DEFAULT (SYSUTCDATETIME()) NOT NULL,
  CONSTRAINT [PK_TableA] PRIMARY KEY ([ID])
);`

	changes, err := classifyTableColumnChanges(prev, next)
	if err != nil {
		t.Fatalf("classify changes failed: %v", err)
	}
	upSQL, downSQL, hasDestructiveDown, err := buildTableAlterColumnsSQL(prev, next, "dbo.TableA", changes)
	if err != nil {
		t.Fatalf("build alter sql failed: %v", err)
	}
	if !hasDestructiveDown {
		t.Fatalf("expected destructive down for added column")
	}

	upUpper := strings.ToUpper(upSQL)
	if !strings.Contains(upUpper, "ALTER TABLE DBO.TABLEA ADD [CREATEDAT] DATETIME2(0) NOT NULL CONSTRAINT [DF_TABLEA_CREATEDAT] DEFAULT (SYSUTCDATETIME()) WITH VALUES;") {
		t.Fatalf("unexpected up SQL: %s", upSQL)
	}
	downUpper := strings.ToUpper(downSQL)
	if !strings.Contains(downUpper, "ALTER TABLE DBO.TABLEA DROP CONSTRAINT [DF_TABLEA_CREATEDAT];") {
		t.Fatalf("expected down SQL to drop default constraint before column: %s", downSQL)
	}
	if !strings.Contains(downUpper, "ALTER TABLE DBO.TABLEA DROP COLUMN [CREATEDAT];") {
		t.Fatalf("unexpected down SQL: %s", downSQL)
	}
	if strings.Index(downUpper, "DROP CONSTRAINT [DF_TABLEA_CREATEDAT]") > strings.Index(downUpper, "DROP COLUMN [CREATEDAT]") {
		t.Fatalf("expected default constraint drop before column drop: %s", downSQL)
	}
}

func TestRemoveNamedConstraintFromCreateTable_InlineDefaultKeepsColumn(t *testing.T) {
	t.Parallel()

	stmt := `CREATE TABLE dbo.TableA (
  [id] INT NOT NULL,
  [status] VARCHAR(10) NULL CONSTRAINT [DF_TableA__status] DEFAULT ('NEW'),
  CONSTRAINT [PK_TableA] PRIMARY KEY ([id])
);`

	got, err := removeNamedConstraintFromCreateTable(stmt, "DF_TableA__status")
	if err != nil {
		t.Fatalf("removeNamedConstraintFromCreateTable failed: %v", err)
	}
	upper := strings.ToUpper(got)
	if strings.Contains(upper, "DF_TABLEA__STATUS") {
		t.Fatalf("expected default constraint removed, got: %s", got)
	}
	if !strings.Contains(upper, "[STATUS] VARCHAR(10) NULL") {
		t.Fatalf("expected column to remain, got: %s", got)
	}
	if _, parseErr := parseTableColumns(got); parseErr != nil {
		t.Fatalf("result must remain parseable, got error: %v\nsql=%s", parseErr, got)
	}
}

func TestSemanticFromMigrationWithState_DropInlineDefaultKeepsColumn(t *testing.T) {
	t.Parallel()

	meta := migrationMeta{
		Category: "tables",
		Object:   "dbo.TableA",
	}
	current := `CREATE TABLE dbo.TableA (
  [id] INT NOT NULL,
  [status] VARCHAR(10) NULL CONSTRAINT [DF_TableA__status] DEFAULT ('NEW'),
  CONSTRAINT [PK_TableA] PRIMARY KEY ([id])
);`
	migration := `ALTER TABLE dbo.TableA DROP CONSTRAINT [DF_TableA__status];`

	got, err := semanticFromMigrationWithState(meta, migration, current, true)
	if err != nil {
		t.Fatalf("semanticFromMigrationWithState failed: %v", err)
	}
	upper := strings.ToUpper(got)
	if !strings.Contains(upper, "[STATUS] VARCHAR(10) NULL") {
		t.Fatalf("expected status column to remain after dropping default, got: %s", got)
	}
	if strings.Contains(upper, "DF_TABLEA__STATUS") {
		t.Fatalf("expected default constraint removed, got: %s", got)
	}
}

func TestRunFmtDetectsInvalidSQLEncoding(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	tableFile := filepath.Join(root, "db_object", "tables", "dbo.TableA.sql")
	if err := os.MkdirAll(filepath.Dir(tableFile), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	raw := []byte("CREATE TABLE dbo.TableA (id INT NOT NULL);\n-- bad byte: ")
	raw = append(raw, 0x96, '\n')
	if err := os.WriteFile(tableFile, raw, 0o644); err != nil {
		t.Fatalf("write table file failed: %v", err)
	}

	e := New(config.Default(root))
	report, err := e.Run("fmt")
	if err != nil {
		t.Fatalf("run fmt failed: %v", err)
	}
	hasEncodingFail := false
	for _, item := range report.Fail.Items {
		if item.Rule == "INVALID_SQL_ENCODING" {
			hasEncodingFail = true
			break
		}
	}
	if !hasEncodingFail {
		t.Fatalf("expected INVALID_SQL_ENCODING fail, got %+v", report.Fail.Items)
	}
}

func osMkdirAll(path string) error {
	return os.MkdirAll(path, 0o755)
}

func osWriteFile(path, content string) error {
	return os.WriteFile(path, []byte(strings.ReplaceAll(content, "\r\n", "\n")), 0o644)
}
