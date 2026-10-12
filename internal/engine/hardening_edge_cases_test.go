package engine

import (
	"strings"
	"testing"
)

func TestTableSemanticHashIncludesPersistedColumnAttributes(t *testing.T) {
	base := `CREATE TABLE dbo.T ([Id] BIGINT IDENTITY(1,1) NOT NULL, [Code] VARCHAR(20) COLLATE Latin1_General_100_BIN2 NULL, [Guid] UNIQUEIDENTIFIER ROWGUIDCOL NULL, [Optional] INT SPARSE NULL, CONSTRAINT [PK_T] PRIMARY KEY ([Id]));`
	variants := []string{
		strings.Replace(base, "IDENTITY(1,1)", "IDENTITY(100,5)", 1),
		strings.Replace(base, "Latin1_General_100_BIN2", "SQL_Latin1_General_CP1_CI_AS", 1),
		strings.Replace(base, " ROWGUIDCOL", "", 1),
		strings.Replace(base, " SPARSE", "", 1),
		strings.Replace(base, "PRIMARY KEY ([Id])", "PRIMARY KEY ([Code])", 1),
	}
	want := hashSemanticByCategory("tables", base)
	for i, variant := range variants {
		if got := hashSemanticByCategory("tables", variant); got == want {
			t.Fatalf("variant %d lost a material table attribute in semantic hashing", i)
		}
	}
}

func TestComputedColumnDropDoesNotUseDataLossGuard(t *testing.T) {
	before := `CREATE TABLE dbo.T ([Base] INT NOT NULL, [Calc] AS ([Base] * 2));`
	after := `CREATE TABLE dbo.T ([Base] INT NOT NULL, [Calc] AS ([Base] * 3));`
	changes, err := classifyTableColumnChanges(before, after)
	if err != nil {
		t.Fatal(err)
	}
	up, _, _, err := buildTableAlterColumnsSQL(before, after, "dbo.T", changes)
	if err != nil {
		t.Fatal(err)
	}
	upper := strings.ToUpper(up)
	if strings.Contains(upper, "WHERE [CALC] IS NOT NULL") {
		t.Fatalf("computed value must not be treated as stored user data:\n%s", up)
	}
	if !strings.Contains(upper, "SYS.COMPUTED_COLUMNS") {
		t.Fatalf("computed-column identity guard missing:\n%s", up)
	}
}

func TestAddedIdentityColumnDoesNotRequireDefault(t *testing.T) {
	before := `CREATE TABLE dbo.T ([Code] INT NULL);`
	after := `CREATE TABLE dbo.T ([Code] INT NULL, [Id] BIGINT IDENTITY(1,1) NOT NULL);`
	changes, err := classifyTableColumnChanges(before, after)
	if err != nil {
		t.Fatal(err)
	}
	up, _, _, err := buildTableAlterColumnsSQL(before, after, "dbo.T", changes)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(strings.ToUpper(up), "[ID] BIGINT IDENTITY(1,1) NOT NULL") {
		t.Fatalf("identity definition missing:\n%s", up)
	}
}

func TestUnprovenTypeTransitionIsRejectedBeforeEmission(t *testing.T) {
	before := `CREATE TABLE dbo.T ([Value] UNIQUEIDENTIFIER NULL);`
	after := `CREATE TABLE dbo.T ([Value] INT NULL);`
	changes, err := classifyTableColumnChanges(before, after)
	if err != nil {
		t.Fatal(err)
	}
	if _, _, _, err := buildTableAlterColumnsSQL(before, after, "dbo.T", changes); err == nil || !strings.Contains(err.Error(), "unproven column type transition") {
		t.Fatalf("expected fail-closed type transition, got %v", err)
	}
}

func TestCheckConstraintExpressionAcceptsNotForReplication(t *testing.T) {
	got, err := extractCheckConstraintExpression(`CONSTRAINT [CK_T] CHECK NOT FOR REPLICATION ([Value] > 0)`)
	if err != nil {
		t.Fatal(err)
	}
	if got != "[Value] > 0" {
		t.Fatalf("unexpected CHECK expression: %q", got)
	}
}

func TestIndexParserIgnoresKeywordsInsideBracketedIdentifiers(t *testing.T) {
	stmt := `CREATE NONCLUSTERED INDEX [IX_ON_INCLUDE] ON [dbo].[ON] ([INCLUDE] ASC) INCLUDE ([WITH], [WHERE]);`
	target, err := parseCreateIndexTargetTable(stmt)
	if err != nil {
		t.Fatal(err)
	}
	if target != "dbo.ON" {
		t.Fatalf("unexpected target: %q", target)
	}
	include, err := parseCreateIndexIncludeColumns(stmt)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Join(include, ",") != "WITH,WHERE" {
		t.Fatalf("unexpected INCLUDE columns: %#v", include)
	}
	if err := validateIndexSubsetSupport(stmt); err != nil {
		t.Fatalf("bracketed keywords must remain valid identifiers: %v", err)
	}
}
