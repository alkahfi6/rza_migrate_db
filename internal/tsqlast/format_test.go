package tsqlast

import (
	"regexp"
	"strings"
	"testing"
)

func TestFormatModuleSQLANTLRFallbackToLite(t *testing.T) {
	t.Parallel()

	in := "CREATE PROC dbo.ProcA AS\nBEGIN\nSELECT 1;\nEND\n"
	out, report, err := FormatModuleSQL(in, FormatConfig{
		StripComments:  true,
		Reindent:       true,
		IndentSize:     2,
		Backend:        BackendANTLR,
		FallbackToLite: true,
	})
	if err != nil {
		t.Fatalf("format should fallback successfully: %v", err)
	}
	if !report.FallbackUsed {
		t.Fatalf("expected fallback to lite backend")
	}
	if report.Backend != BackendLite {
		t.Fatalf("expected report backend lite, got %s", report.Backend)
	}
	if !strings.Contains(out, "SELECT 1;") {
		t.Fatalf("unexpected formatter output: %s", out)
	}
}

func TestFormatModuleSQLEquivalenceANTLRFallbackVsLite(t *testing.T) {
	t.Parallel()

	samples := []struct {
		name  string
		input string
	}{
		{name: "procedure", input: "CREATE PROCEDURE dbo.UpdateRisk AS\nBEGIN\nUPDATE dbo.Risk SET Score = @score;\nEND\n"},
		{name: "view", input: "CREATE VIEW dbo.RiskSummary AS SELECT Id, Score FROM dbo.Risk;\n"},
		{name: "function", input: "CREATE FUNCTION dbo.RiskScore(@score INT) RETURNS INT AS BEGIN RETURN @score; END\n"},
	}

	for _, s := range samples {
		s := s
		t.Run(s.name, func(t *testing.T) {
			t.Parallel()

			liteOut, liteReport, err := FormatModuleSQL(s.input, FormatConfig{
				StripComments:  true,
				Reindent:       true,
				IndentSize:     2,
				Backend:        BackendLite,
				FallbackToLite: true,
			})
			if err != nil {
				t.Fatalf("lite backend failed: %v", err)
			}
			if liteReport.Backend != BackendLite {
				t.Fatalf("unexpected lite report backend: %s", liteReport.Backend)
			}

			fallbackOut, fallbackReport, err := FormatModuleSQL(s.input, FormatConfig{
				StripComments:  true,
				Reindent:       true,
				IndentSize:     2,
				Backend:        BackendANTLR,
				FallbackToLite: true,
			})
			if err != nil {
				t.Fatalf("antlr fallback failed: %v", err)
			}
			if !fallbackReport.FallbackUsed {
				t.Fatalf("expected antlr backend to fallback in current build")
			}
			if fallbackOut != liteOut {
				t.Fatalf("antlr fallback output must equal lite output for %s", s.name)
			}
		})
	}
}

func TestFormatModuleSQLControlFlowElseEndNotMisclassifiedAsSelectCase(t *testing.T) {
	t.Parallel()

	in := "if exists (select 1 from ReksaKuesionerTMP\n" +
		"  where NoCIF = @pnNoCIF and NoJawaban = @pnNoJawaban and KodeSoal = @pnKodeSoal)\n" +
		"begin\n" +
		"  select 'True' value\n" +
		"\tend\n" +
		"\telse\n" +
		"begin\n" +
		"  select 'False' value\n" +
		"\tend\n\n" +
		"\t--tambah info\n" +
		"\tselect 'hasil check','datas'\n"

	out, _, err := FormatModuleSQL(in, FormatConfig{
		StripComments:  false,
		Reindent:       true,
		IndentSize:     2,
		Backend:        BackendLite,
		FallbackToLite: true,
	})
	if err != nil {
		t.Fatalf("format failed: %v", err)
	}

	if strings.Contains(out, "\tend\n") || strings.Contains(out, "\telse\n") {
		t.Fatalf("control flow END/ELSE should not be tab-indented as select projection:\n%s", out)
	}
	if !strings.Contains(out, "end\nelse\nbegin\n  select 'False' value\nend\n\n\t--tambah info\nselect 'hasil check','datas'") {
		t.Fatalf("unexpected control-flow formatting:\n%s", out)
	}
}

func TestFormatModuleSQLSelectCaseProjectionStillIndented(t *testing.T) {
	t.Parallel()

	in := "SELECT\n" +
		"CASE\n" +
		"WHEN x = 1 THEN 'A'\n" +
		"ELSE 'B'\n" +
		"END as Remark,\n" +
		"col2\n" +
		"FROM dbo.T\n"

	out, _, err := FormatModuleSQL(in, FormatConfig{
		StripComments:  false,
		Reindent:       true,
		IndentSize:     2,
		Backend:        BackendLite,
		FallbackToLite: true,
	})
	if err != nil {
		t.Fatalf("format failed: %v", err)
	}

	if !strings.Contains(out, "\tCASE") || !strings.Contains(out, "\tEND as Remark,") {
		t.Fatalf("select CASE projection indentation regressed:\n%s", out)
	}
}

func TestFormatModuleSQLSetThenRaiserrorNoCarryIndent(t *testing.T) {
	t.Parallel()

	in := "if isnull(@cErrMsg,'') = ''\n" +
		"  set @cErrMsg = 'Unknown Error'\n\n" +
		"\traiserror (@cErrMsg, 16, 1)\n"

	out, _, err := FormatModuleSQL(in, FormatConfig{
		StripComments:  false,
		Reindent:       true,
		IndentSize:     2,
		Backend:        BackendLite,
		FallbackToLite: true,
	})
	if err != nil {
		t.Fatalf("format failed: %v", err)
	}

	if strings.Contains(out, "\n\traiserror") {
		t.Fatalf("raiserror should not inherit continuation indent from previous SET:\n%s", out)
	}
	if !strings.Contains(out, "\nraiserror (@cErrMsg, 16, 1)\n") {
		t.Fatalf("raiserror line missing or malformed:\n%s", out)
	}
}

func TestFormatModuleSQLRaiserrorThrowInsideBeginStayIndented(t *testing.T) {
	t.Parallel()

	in := "IF (@x = 1)\n" +
		"BEGIN\n" +
		"RAISERROR ('err', 16, 1)\n" +
		"THROW 50000, 'boom', 1\n" +
		"END\n"

	out, _, err := FormatModuleSQL(in, FormatConfig{
		StripComments:  false,
		Reindent:       true,
		IndentSize:     2,
		Backend:        BackendLite,
		FallbackToLite: true,
	})
	if err != nil {
		t.Fatalf("format failed: %v", err)
	}

	if !strings.Contains(out, "\n  RAISERROR ('err', 16, 1)\n") {
		t.Fatalf("RAISERROR should stay indented inside BEGIN block:\n%s", out)
	}
	if !strings.Contains(out, "\n  THROW 50000, 'boom', 1\n") {
		t.Fatalf("THROW should stay indented inside BEGIN block:\n%s", out)
	}
}

func TestFormatModuleSQLSelectCaseProjectionNoMixedTabSpaceAndCompactBlank(t *testing.T) {
	t.Parallel()

	in := "SELECT\n" +
		"(\n" +
		"\n" +
		"CASE\n" +
		"WHEN isnull(a.NonNISPAccId,'') != ''\n" +
		"THEN a.NonNISPAccId\n" +
		"ELSE a.NISPAccId\n" +
		"END\n" +
		"\n" +
		") as BeneficiaryAccount,\n" +
		"col2\n" +
		"FROM dbo.T\n"

	out, _, err := FormatModuleSQL(in, FormatConfig{
		StripComments:  false,
		Reindent:       true,
		IndentSize:     2,
		Backend:        BackendLite,
		FallbackToLite: true,
	})
	if err != nil {
		t.Fatalf("format failed: %v", err)
	}

	if strings.Contains(out, "\n\n\tCASE") || strings.Contains(out, "\n\tEND\n\n\t)") {
		t.Fatalf("blank lines around CASE projection paren block should be compacted:\n%s", out)
	}
	mixed := regexp.MustCompile(`(?m)^ +\t(?:CASE|WHEN|THEN|ELSE|END)\b`)
	if mixed.MatchString(out) {
		t.Fatalf("CASE projection keyword lines should not mix leading spaces+tab:\n%s", out)
	}
	if !strings.Contains(out, "\tCASE") || !strings.Contains(out, "\t\tWHEN") || !strings.Contains(out, "\tEND") {
		t.Fatalf("CASE projection alignment regressed:\n%s", out)
	}
}

func TestFormatModuleSQLSelectCasePayloadIndentedDeeperThanWhenElse(t *testing.T) {
	t.Parallel()

	in := "SELECT\n" +
		"case\n" +
		"when isnull(g.CIFConfAddress1, '') != '' then\n" +
		"isnull(g.CIFConfAddress1 + ' ', '') +\n" +
		"isnull(g.CIFConfAddress2 + ' ', '')\n" +
		"else\n" +
		"isnull(g.CIFAddress1 + ' ', '')\n" +
		"end as AddrBuilding,\n" +
		"col2\n" +
		"FROM dbo.T\n"

	out, _, err := FormatModuleSQL(in, FormatConfig{
		StripComments:  false,
		Reindent:       true,
		IndentSize:     2,
		Backend:        BackendLite,
		FallbackToLite: true,
	})
	if err != nil {
		t.Fatalf("format failed: %v", err)
	}

	if !strings.Contains(out, "\t\twhen isnull(g.CIFConfAddress1, '') != '' then") {
		t.Fatalf("WHEN should be indented one level under CASE:\n%s", out)
	}
	if !strings.Contains(out, "\t\t\tisnull(g.CIFConfAddress1 + ' ', '') +") {
		t.Fatalf("CASE payload should be one level deeper than WHEN:\n%s", out)
	}
	if !strings.Contains(out, "\t\telse") {
		t.Fatalf("ELSE should be indented one level under CASE:\n%s", out)
	}
	if !strings.Contains(out, "\t\t\tisnull(g.CIFAddress1 + ' ', '')") {
		t.Fatalf("ELSE payload should be one level deeper than ELSE:\n%s", out)
	}
}
