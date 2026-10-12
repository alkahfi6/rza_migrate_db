package getobject

import (
	"strings"
	"testing"
)

func TestFormatModuleSQLWithOptionsPreservesCommentsWhenDisabled(t *testing.T) {
	t.Parallel()

	in := `CREATE OR ALTER PROCEDURE dbo.ProcA
AS
BEGIN
-- keep this comment
IF EXISTS (SELECT 1)
SELECT 1;
END
`
	out := FormatModuleSQLWithOptions(in, ModuleFormatOptions{
		StripComments: false,
		Reindent:      true,
		IndentSize:    2,
		Backend:       "antlr",
	})

	if !strings.Contains(out, "-- keep this comment") {
		t.Fatalf("comment must be preserved when StripComments=false:\n%s", out)
	}
	if !strings.Contains(out, "IF EXISTS (SELECT 1)\n    SELECT 1;") {
		t.Fatalf("expected reindent for IF body:\n%s", out)
	}
}

func TestFormatModuleSQLWithOptionsStripsCommentsWhenEnabled(t *testing.T) {
	t.Parallel()

	in := `CREATE OR ALTER VIEW dbo.ViewA
AS
-- remove me
SELECT 1 AS id;
`
	out := FormatModuleSQLWithOptions(in, ModuleFormatOptions{
		StripComments: true,
		Reindent:      true,
		IndentSize:    2,
		Backend:       "antlr",
	})

	if strings.Contains(out, "remove me") {
		t.Fatalf("comment should be removed when StripComments=true:\n%s", out)
	}
}
