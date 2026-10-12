package engine

import (
	"strings"
	"testing"
)

func TestExistingTableCompatibilityChecksRequiredNamesAndBaseTypesOnly(t *testing.T) {
	t.Parallel()

	condition, supported := buildExistingTableEquivalentCondition(
		"dbo.Example",
		"CREATE TABLE dbo.Example (code VARCHAR NOT NULL, amount DECIMAL NULL, happened_at DATETIME2 NULL);",
	)
	if !supported {
		t.Fatal("simple CREATE TABLE should support bootstrap compatibility")
	}
	for _, expected := range []string{
		"N'code'",
		"N'amount'",
		"N'happened_at'",
		"N'VARCHAR'",
		"IN (N'DECIMAL', N'NUMERIC')",
		"N'DATETIME2'",
		"c.is_computed = 0",
		"c.is_identity = 0",
	} {
		if !strings.Contains(condition, expected) {
			t.Fatalf("compatibility condition is missing %q:\n%s", expected, condition)
		}
	}
	for _, ignored := range []string{"c.column_id =", "c.max_length =", "c.precision =", "c.scale =", "c.is_nullable =", "collation_name"} {
		if strings.Contains(condition, ignored) {
			t.Fatalf("DBA-managed detail %q must not block bootstrap:\n%s", ignored, condition)
		}
	}
}

func TestExistingTableCompatibilityKeepsIdentityAndComputedKindButIgnoresConstraintDetails(t *testing.T) {
	t.Parallel()

	create := `CREATE TABLE dbo.Child (
  [Id] BIGINT IDENTITY(7,3) NOT NULL,
  [ParentId] BIGINT NULL,
  [Code] VARCHAR(20) COLLATE Latin1_General_BIN NOT NULL CONSTRAINT [DF_Child_Code] DEFAULT ('PENDING'),
  [CodeLen] AS (LEN([Code])) PERSISTED,
  CONSTRAINT [PK_Child] PRIMARY KEY CLUSTERED ([Id] ASC),
  CONSTRAINT [UQ_Child_Code] UNIQUE NONCLUSTERED ([Code] DESC),
  CONSTRAINT [CK_Child_Code] CHECK ([Code]<>('')),
  CONSTRAINT [FK_Child_Parent] FOREIGN KEY ([ParentId]) REFERENCES dbo.Parent ([Id]) ON DELETE CASCADE ON UPDATE NO ACTION
);`
	items, err := extractCreateTableItems(create)
	if err != nil {
		t.Fatalf("extract table items: %v", err)
	}
	for _, item := range items {
		upper := strings.ToUpper(strings.TrimSpace(item))
		if isConstraintLikeTableItem(upper) {
			continue
		}
		if _, err := parseTableColumnDefFragment(item); err != nil {
			t.Fatalf("parse column item %q: %v", item, err)
		}
	}
	if _, err := extractTableDefaultConstraintMap(create); err != nil {
		t.Fatalf("parse defaults: %v", err)
	}
	if _, ok := buildTableConstraintEquivalentConditions("dbo.Child", items); !ok {
		t.Fatal("parse table constraints")
	}
	condition, supported := buildExistingTableEquivalentCondition("dbo.Child", create)
	if !supported {
		t.Fatal("supported bootstrap-compatible table definition unexpectedly fell back")
	}
	for _, expected := range []string{
		"c.is_identity = 1",
		"c.is_computed = 1",
		"N'Id'",
		"N'ParentId'",
		"N'Code'",
		"N'CodeLen'",
	} {
		if !strings.Contains(condition, expected) {
			t.Fatalf("bootstrap compatibility condition is missing %q:\n%s", expected, condition)
		}
	}
	for _, ignored := range []string{
		"ic.seed_value",
		"sys.default_constraints",
		"sys.key_constraints",
		"sys.check_constraints",
		"sys.foreign_keys",
		"cc.definition",
	} {
		if strings.Contains(condition, ignored) {
			t.Fatalf("non-blocking table detail %q leaked into compatibility guard:\n%s", ignored, condition)
		}
	}
}

func TestColumnFragmentParserRetainsIdentityAndComputedMetadata(t *testing.T) {
	t.Parallel()

	identity, err := parseTableColumnDefFragment("[Id] BIGINT IDENTITY(9,2) NOT NULL")
	if err != nil {
		t.Fatalf("parse identity: %v", err)
	}
	if identity.IdentityDefinition != "IDENTITY(9,2)" || identity.Nullable {
		t.Fatalf("identity metadata was lost: %#v", identity)
	}

	computed, err := parseTableColumnDefFragment("[Total] AS ([Qty]*[Price]) PERSISTED")
	if err != nil {
		t.Fatalf("parse computed: %v", err)
	}
	if !computed.IsComputed || computed.ComputedDefinition != "AS ([Qty]*[Price]) PERSISTED" {
		t.Fatalf("computed metadata was lost: %#v", computed)
	}

	whitespace, err := parseTableColumnDefFragment("[ ] NVARCHAR(255) NULL")
	if err != nil {
		t.Fatalf("parse whitespace identifier: %v", err)
	}
	if whitespace.Name != " " {
		t.Fatalf("whitespace identifier was changed: %#v", whitespace)
	}
	if _, supported := buildExistingTableEquivalentCondition(
		"dbo.WhitespaceName",
		"CREATE TABLE dbo.WhitespaceName ([ ] NVARCHAR(255) NULL);",
	); !supported {
		t.Fatal("whitespace-only delimited identifier must support bootstrap compatibility")
	}
}

func TestExplicitIndexDataSpaceFailsClosed(t *testing.T) {
	t.Parallel()

	if !hasExplicitIndexDataSpaceClause("CREATE INDEX IX_Example ON dbo.Example(id) ON [PRIMARY];") {
		t.Fatal("explicit index data-space clause was not detected")
	}
	if hasExplicitIndexDataSpaceClause("CREATE INDEX IX_Example ON dbo.Example(id) INCLUDE (code) WITH (FILLFACTOR = 90);") {
		t.Fatal("ordinary index options were misclassified as a data-space clause")
	}
}

func TestProgrammabilityWrapperPreservesMultilineStringLiteralBytes(t *testing.T) {
	t.Parallel()

	module := "CREATE OR ALTER PROCEDURE dbo.MultilineLiteral\nAS\nBEGIN\n  DECLARE @sql NVARCHAR(MAX) = N'first line\nsecond line';\n  EXEC sys.sp_executesql @sql;\nEND;"
	wrapped := renderTxWrapper(
		"20991231235959999",
		"procedures",
		"20991231235959999_procedures_dbo.MultilineLiteral_v1.up.sql",
		"dbo.MultilineLiteral",
		true,
		renderProgrammabilityExecutableBody(module),
	)
	decoded, ok := extractSpExecuteSQLBatch(wrapped)
	if !ok {
		t.Fatal("wrapped programmable body could not be decoded")
	}
	want := strings.TrimSpace(compactGeneratedSQLWhitespace(module))
	if decoded != want {
		t.Fatalf("wrapper changed executable module bytes\nwant:\n%s\n\ngot:\n%s", want, decoded)
	}
}

func TestEngineInfrastructureNamesRemainCaseSensitive(t *testing.T) {
	t.Parallel()

	if !isEngineInfrastructureObject("tables", "dbo.schema_migrations") {
		t.Fatal("golang-migrate state table must be treated as engine infrastructure")
	}
	if isEngineInfrastructureObject("tables", "dbo.Schema_Migrations") {
		t.Fatal("case-distinct user table was incorrectly treated as engine infrastructure")
	}
	if isEngineInfrastructureObject("procedures", "dbo.USP_MIG_LOG_EVENT") {
		t.Fatal("case-distinct user procedure was incorrectly treated as engine infrastructure")
	}
}

func TestRenderProgrammabilityExecutableBodyHandlesLeadingBannerComment(t *testing.T) {
	body := "/****** Object: StoredProcedure dbo.P ******/\nCREATE OR ALTER   PROCEDURE dbo.P AS BEGIN SELECT 1; END;\n"
	rendered := renderProgrammabilityExecutableBody(body)
	if !strings.HasPrefix(rendered, "EXEC sys.sp_executesql N'") {
		t.Fatalf("comment-prefixed module was not isolated in a dynamic batch:\n%s", rendered)
	}
	if !strings.Contains(rendered, "PROCEDURE dbo.P") {
		t.Fatalf("module body was lost:\n%s", rendered)
	}
}

func TestIndexColumnParserPreservesBracketedIdentifiersWithSpaces(t *testing.T) {
	t.Parallel()

	statement := `CREATE INDEX IX_Account
ON dbo.Account ([Nomor] ASC, [SID] DESC, [Nomor Rekening RDN] ASC, [Kode Produk] ASC)
INCLUDE ([Owner's Note], [Occurred At]);`
	keys, _, err := parseCreateIndexKeysAndFilter(statement)
	if err != nil {
		t.Fatalf("parse index keys: %v", err)
	}
	includes, err := parseCreateIndexIncludeColumns(statement)
	if err != nil {
		t.Fatalf("parse index includes: %v", err)
	}
	if got, want := strings.Join(keys, "|"), "Nomor|SID|Nomor Rekening RDN|Kode Produk"; got != want {
		t.Fatalf("key columns=%q want=%q", got, want)
	}
	if got, want := strings.Join(includes, "|"), "Owner's Note|Occurred At"; got != want {
		t.Fatalf("include columns=%q want=%q", got, want)
	}
}
