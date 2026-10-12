package sqlfix

import (
	"strings"
	"testing"
)

func TestRewriteSafeTableNamesDefaultsRemovesTrailingCommaAndGO(t *testing.T) {
	in := "CREATE TABLE dbo.Outbox\n(\n  process_guid UNIQUEIDENTIFIER NOT NULL DEFAULT NEWID(),\n  retry_count INT NOT NULL DEFAULT 0, -- keep\n);\nGO\n"
	got, err := RewriteSafe(in, Options{Category: "tables", ObjectName: "dbo.Outbox"})
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"CONSTRAINT DF_Outbox__process_guid DEFAULT NEWID()",
		"CONSTRAINT DF_Outbox__retry_count DEFAULT 0",
		"-- keep\n)",
	} {
		if !strings.Contains(got.SQL, want) {
			t.Fatalf("missing %q:\n%s", want, got.SQL)
		}
	}
	if strings.Contains(got.SQL, "\nGO\n") {
		t.Fatalf("GO was not removed:\n%s", got.SQL)
	}
	assertIdempotent(t, got.SQL, Options{Category: "tables", ObjectName: "dbo.Outbox"})
}

func TestRewriteSafeModulePreamblePreservesBannerAndBodyTrivia(t *testing.T) {
	in := "USE [SQL_TRSRETAIL]\nGO\n/****** banner USE x GO ******/\nSET ANSI_NULLS ON\nGO\nSET QUOTED_IDENTIFIER ON\nGO\nALTER     PROC [dbo].[P]\nAS\nBEGIN\n  SELECT N'USE X; GO; PROC'; -- PROC\nEND;"
	got, err := RewriteSafe(in, Options{Category: "procedures", ObjectName: "dbo.P"})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got.SQL, "/****** banner USE x GO ******/") || !strings.Contains(got.SQL, "N'USE X; GO; PROC'; -- PROC") {
		t.Fatalf("comments or literals changed:\n%s", got.SQL)
	}
	if !strings.Contains(got.SQL, "CREATE OR ALTER PROCEDURE [dbo].[P]") {
		t.Fatalf("procedure header not canonicalized:\n%s", got.SQL)
	}
	for _, forbidden := range []string{"USE [SQL_TRSRETAIL]", "SET ANSI_NULLS", "SET QUOTED_IDENTIFIER", "\nGO\n"} {
		if strings.Contains(got.SQL, forbidden) {
			t.Fatalf("preamble %q remains:\n%s", forbidden, got.SQL)
		}
	}
	if !strings.HasSuffix(got.SQL, "\n") {
		t.Fatal("expected final newline")
	}
	assertIdempotent(t, got.SQL, Options{Category: "procedures", ObjectName: "dbo.P"})
}

func TestRewriteSafePatchOnlyRemovesLeadingCatalogBatch(t *testing.T) {
	in := "-- deploy in target\nUSE SQL_REPLICATE;\nGO\nSELECT N'USE SQL_OTHER;';\nUSE SQL_OTHER;\nGO\n"
	got, err := RewriteSafe(in, Options{Category: "patch", Patch: true})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(got.SQL, "USE SQL_REPLICATE") {
		t.Fatalf("leading USE remains:\n%s", got.SQL)
	}
	if !strings.Contains(got.SQL, "USE SQL_OTHER;") {
		t.Fatalf("non-leading USE must be preserved for verifier:\n%s", got.SQL)
	}
	if !strings.Contains(got.SQL, "N'USE SQL_OTHER;'") {
		t.Fatalf("literal changed:\n%s", got.SQL)
	}
}

func TestRewriteSafeDoesNotRemoveProcedureArgumentComma(t *testing.T) {
	in := "CREATE OR ALTER PROCEDURE dbo.P AS\nBEGIN\n  EXEC dbo.Q @a = COALESCE((SELECT 1), 0), @b = 2;\nEND;\n"
	got, err := RewriteSafe(in, Options{Category: "procedures", ObjectName: "dbo.P"})
	if err != nil {
		t.Fatal(err)
	}
	if got.SQL != in {
		t.Fatalf("procedure body changed:\n%s", got.SQL)
	}
}

func TestRewriteSafeFoldsLegacyAlterDefault(t *testing.T) {
	in := "CREATE TABLE T (\n  id INT NOT NULL,\n  created_at DATETIME NULL\n);\nALTER TABLE [dbo].[T] ADD DEFAULT (getdate()) FOR [created_at]\n"
	got, err := RewriteSafe(in, Options{Category: "tables", ObjectName: "dbo.T"})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got.SQL, "created_at DATETIME NULL CONSTRAINT DF_T__created_at DEFAULT (getdate())") {
		t.Fatalf("default was not folded:\n%s", got.SQL)
	}
	if !strings.HasPrefix(got.SQL, "CREATE TABLE dbo.T") {
		t.Fatalf("table target was not repository-qualified:\n%s", got.SQL)
	}
	if strings.Contains(strings.ToUpper(got.SQL), "ALTER TABLE") {
		t.Fatalf("legacy ALTER TABLE remains:\n%s", got.SQL)
	}
	assertIdempotent(t, got.SQL, Options{Category: "tables", ObjectName: "dbo.T"})
}

func TestDetectIndexTargetsIgnoresCommentsAndDynamicSQL(t *testing.T) {
	in := "-- CREATE INDEX fake ON dbo.Nope(id)\nCREATE CLUSTERED INDEX IX1 ON [dbo].[Target_TT](id);\nCREATE INDEX IX2 ON dbo.Target_TT(code);\nEXEC(N'CREATE INDEX fake ON dbo.Dynamic(id)');"
	got := DetectIndexTargets(in)
	if len(got) != 1 || got[0] != "dbo.Target_TT" {
		t.Fatalf("unexpected targets: %v", got)
	}
}

func TestRewriteSafeRemovesOnlyInvalidColumnCollation(t *testing.T) {
	in := "CREATE TABLE dbo.T (\n  payload XML COLLATE Latin1_General_BIN NULL,\n  code VARCHAR(20) COLLATE Latin1_General_BIN NULL\n);\n"
	got, err := RewriteSafe(in, Options{Category: "tables", ObjectName: "dbo.T"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(got.SQL, "XML COLLATE") {
		t.Fatalf("invalid XML collation remains:\n%s", got.SQL)
	}
	if !strings.Contains(got.SQL, "VARCHAR(20) COLLATE Latin1_General_BIN") {
		t.Fatalf("valid character collation was removed:\n%s", got.SQL)
	}
	assertIdempotent(t, got.SQL, Options{Category: "tables", ObjectName: "dbo.T"})
}

func TestRewriteSafeRemovesOnlyManifestProvenSelfCatalog(t *testing.T) {
	in := "CREATE OR ALTER VIEW dbo.V AS SELECT * FROM SQL_REPLICATE.dbo.LocalT l JOIN SQL_OTHER.dbo.ExternalT e ON 1=1;\n"
	got, err := RewriteSafe(in, Options{
		Category: "views", ObjectName: "dbo.V", SelfCatalog: "SQL_REPLICATE",
		LocalObjects: map[string]bool{"dbo.localt": true},
	})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(got.SQL, "SQL_REPLICATE.dbo.LocalT") || !strings.Contains(got.SQL, "dbo.LocalT") {
		t.Fatalf("self catalog was not removed:\n%s", got.SQL)
	}
	if !strings.Contains(got.SQL, "SQL_OTHER.dbo.ExternalT") {
		t.Fatalf("external catalog was changed:\n%s", got.SQL)
	}
}

func TestConstraintNameIsBoundedAndStable(t *testing.T) {
	name := constraintName("DF", strings.Repeat("table", 30), strings.Repeat("column", 30))
	if len(name) > 128 {
		t.Fatalf("constraint name length=%d", len(name))
	}
	if name != constraintName("DF", strings.Repeat("table", 30), strings.Repeat("column", 30)) {
		t.Fatal("constraint name is not deterministic")
	}
}

func TestRewriteSafeNamesProvenUnnamedConstraintsWithoutEnforcingUserPrefixes(t *testing.T) {
	in := `CREATE TABLE dbo.Child (
  id INT NOT NULL,
  parent_id INT NULL REFERENCES dbo.Parent(id),
  code VARCHAR(20) NULL UNIQUE,
  amount INT NULL CHECK (amount >= 0),
  PRIMARY KEY CLUSTERED (id)
);
`
	got, err := RewriteSafe(in, Options{Category: "tables", ObjectName: "dbo.Child"})
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"CONSTRAINT FK_Child__parent_id_Parent_",
		"CONSTRAINT UQ_Child__code UNIQUE",
		"CONSTRAINT CK_Child__amount_",
		"CONSTRAINT PK_Child PRIMARY KEY",
	} {
		if !strings.Contains(got.SQL, want) {
			t.Fatalf("missing %q:\n%s", want, got.SQL)
		}
	}
	assertIdempotent(t, got.SQL, Options{Category: "tables", ObjectName: "dbo.Child"})
}

func TestRewriteSafeNamesSecondUnnamedInlineConstraint(t *testing.T) {
	in := `CREATE TABLE dbo.T (
  amount INT CONSTRAINT LEGACY_DEFAULT DEFAULT 0 CHECK (amount >= 0)
);
`
	got, err := RewriteSafe(in, Options{Category: "tables", ObjectName: "dbo.T"})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got.SQL, "CONSTRAINT LEGACY_DEFAULT DEFAULT 0 CONSTRAINT CK_T__amount_") {
		t.Fatalf("second unnamed constraint was not named safely:\n%s", got.SQL)
	}
	assertIdempotent(t, got.SQL, Options{Category: "tables", ObjectName: "dbo.T"})
}

func TestRewriteSafePreservesAmbiguousOrExplicitLegacyConstraintNames(t *testing.T) {
	in := `CREATE TABLE dbo.T (
  id INT CONSTRAINT LEGACY_ANY_NAME PRIMARY KEY,
  amount INT DEFAULT 0 CHECK (amount >= 0),
  CONSTRAINT PRIMARY KEY (id)
);
`
	got, err := RewriteSafe(in, Options{Category: "tables", ObjectName: "dbo.T"})
	if err != nil {
		t.Fatal(err)
	}
	if got.SQL != in {
		t.Fatalf("ambiguous/explicit constraints must be preserved:\n%s", got.SQL)
	}
}

func assertIdempotent(t *testing.T, input string, opts Options) {
	t.Helper()
	second, err := RewriteSafe(input, opts)
	if err != nil {
		t.Fatal(err)
	}
	if second.SQL != input {
		t.Fatalf("rewrite is not idempotent:\n--- first\n%s\n--- second\n%s", input, second.SQL)
	}
}
