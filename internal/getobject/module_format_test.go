package getobject

import (
	"strings"
	"testing"
)

func TestStripSQLCommentsPreservesStringLiterals(t *testing.T) {
	t.Parallel()

	in := `
CREATE OR ALTER PROCEDURE dbo.ProcA
AS
BEGIN
  -- this comment should be removed
  SELECT 'text -- keep inside string' AS s, [--not_comment] AS c;
  /* block
     comment */
  SELECT 2 AS n;
END
`
	out := stripSQLComments(in)
	if strings.Contains(out, "this comment should be removed") {
		t.Fatalf("line comment was not removed")
	}
	if strings.Contains(out, "block\n") {
		t.Fatalf("block comment was not removed")
	}
	if !strings.Contains(out, "'text -- keep inside string'") {
		t.Fatalf("string literal content changed unexpectedly")
	}
	if !strings.Contains(out, "[--not_comment]") {
		t.Fatalf("bracket identifier changed unexpectedly")
	}
}

func TestReindentTSQLBeginEnd(t *testing.T) {
	t.Parallel()

	in := `
CREATE OR ALTER PROCEDURE dbo.ProcA
AS
BEGIN
IF EXISTS (SELECT 1)
BEGIN
SELECT 1;
END
END
`
	out := reindentTSQL(in, 2)
	wantContains := []string{
		"BEGIN\n  IF EXISTS (SELECT 1)\n  BEGIN\n    SELECT 1;\n  END\nEND",
	}
	for _, w := range wantContains {
		if !strings.Contains(out, w) {
			t.Fatalf("unexpected indentation result:\n%s", out)
		}
	}
}

func TestReindentTSQLIfWithoutBeginIndentsNextLine(t *testing.T) {
	t.Parallel()

	in := `
CREATE OR ALTER PROCEDURE dbo.ProcA
AS
BEGIN
IF EXISTS (SELECT 1)
SELECT 1;
ELSE IF EXISTS (SELECT 2)
SELECT 2;
ELSE
SELECT 3;
END
`
	out := reindentTSQL(in, 2)
	wantContains := []string{
		"IF EXISTS (SELECT 1)\n    SELECT 1;",
		"ELSE IF EXISTS (SELECT 2)\n    SELECT 2;",
		"ELSE\n    SELECT 3;",
	}
	for _, w := range wantContains {
		if !strings.Contains(out, w) {
			t.Fatalf("expected pattern not found:\n%s\n--- output ---\n%s", w, out)
		}
	}
}

func TestExpandTrailingControlBegin(t *testing.T) {
	t.Parallel()

	got := expandTrailingControlBegin("ELSE IF @a = 1 BEGIN")
	if len(got) != 2 || got[0] != "ELSE IF @a = 1" || got[1] != "BEGIN" {
		t.Fatalf("unexpected expand result: %#v", got)
	}
}

func TestFormatModuleDefinitionWithConfig(t *testing.T) {
	t.Parallel()

	cfg := DefaultAppConfig(".")
	cfg.ConnectionString = "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable"
	cfg.StripModuleComments = true
	cfg.ReindentModuleSQL = true
	cfg.IndentSize = 2

	in := "CREATE OR ALTER VIEW dbo.ViewA AS\n-- note\nSELECT 1 AS a;\n"
	out := formatModuleDefinition(in, cfg)
	if strings.Contains(out, "note") {
		t.Fatalf("comment should be removed")
	}
	if !strings.Contains(out, "SELECT 1 AS a;") {
		t.Fatalf("query should remain")
	}
}

func TestReindentTSQLCaseAlignment(t *testing.T) {
	t.Parallel()

	in := `
CREATE OR ALTER PROCEDURE dbo.ProcA
AS
BEGIN
SELECT CASE WHEN a=1 THEN 'A' ELSE 'B' END AS X;
END
`
	out := reindentTSQL(in, 2)
	want := []string{
		"CASE",
		"WHEN a=1 THEN 'A'",
		"ELSE 'B'",
		"END AS X;",
	}
	for _, w := range want {
		if !strings.Contains(out, w) {
			t.Fatalf("missing case aligned token %q in output:\n%s", w, out)
		}
	}
}

func TestSplitEndTrailingSuffix(t *testing.T) {
	t.Parallel()

	got := splitEndTrailingSuffix("END )")
	if len(got) != 2 || got[0] != "END" || got[1] != ")" {
		t.Fatalf("unexpected split result: %#v", got)
	}
}

func TestReindentTSQLJoinWhereAndTabContinuation(t *testing.T) {
	t.Parallel()

	in := `
SELECT a
FROM dbo.TableA t
JOIN dbo.TableB b
ON t.id = b.id
AND b.active = 1
WHERE t.flag = 1
AND t.status = 'A'
ORDER BY t.id
`
	out := reindentTSQL(in, 2)
	want := []string{
		"ON t.id = b.id",
		"\t\tAND b.active = 1",
		"WHERE t.flag = 1",
		"\tAND t.status = 'A'",
	}
	for _, w := range want {
		if !strings.Contains(out, w) {
			t.Fatalf("expected continuation style %q not found in output:\n%s", w, out)
		}
	}
}

func TestReindentTSQLSelectAndDeclareListIndent(t *testing.T) {
	t.Parallel()

	in := `
CREATE OR ALTER PROCEDURE dbo.ProcA
AS
BEGIN
DECLARE
@nOK int
,@cDateToday varchar(8)
SELECT
col1
,col2
FROM dbo.TableA
END
`
	out := reindentTSQL(in, 2)
	want := []string{
		"\t@nOK int",
		"\t, @cDateToday varchar(8)",
		"\tcol1",
		"\t,col2",
		"\n  FROM dbo.TableA",
	}
	for _, w := range want {
		if !strings.Contains(out, w) {
			t.Fatalf("expected list indentation %q not found in output:\n%s", w, out)
		}
	}
}

func TestReindentTSQLProcedureHeaderParamListIndentAndCompaction(t *testing.T) {
	t.Parallel()

	in := `
CREATE OR ALTER PROCEDURE dbo.ProcA

@pnBillId               int
,@pnNIK             int
, @pcGuid            varchar(50) = 'A  B'
as
SELECT 1
`
	out := reindentTSQL(in, 2)
	want := []string{
		"\t@pnBillId int",
		"\t, @pnNIK int",
		"\t, @pcGuid varchar(50) = 'A  B'",
		"\nas",
	}
	for _, w := range want {
		if !strings.Contains(out, w) {
			t.Fatalf("expected procedure header formatting %q not found in output:\n%s", w, out)
		}
	}
	if strings.Contains(out, "ProcA\n\n\t@pnBillId") {
		t.Fatalf("blank line between procedure header and first param should be compacted:\n%s", out)
	}
}

func TestReindentTSQLDeclareHeaderNormalizesAlignedTabs(t *testing.T) {
	t.Parallel()

	in := `
DECLARE @nKurs				decimal(19,2)
, @nTotalIDR			decimal (19,2)
`
	out := reindentTSQL(in, 2)
	want := []string{
		"DECLARE @nKurs decimal(19,2)",
		"\t, @nTotalIDR decimal (19,2)",
	}
	for _, w := range want {
		if !strings.Contains(out, w) {
			t.Fatalf("expected declare normalization %q not found in output:\n%s", w, out)
		}
	}
}

func TestReindentTSQLSelectIntoAndBracketIdentifier(t *testing.T) {
	t.Parallel()

	in := `
SELECT
ColA,
[SID],
ColC
INTO #tmpData
FROM dbo.TableA
`
	out := reindentTSQL(in, 2)
	want := []string{
		"\tColA,",
		"\t[SID],",
		"\tColC",
		"\nINTO #tmpData",
		"\nFROM dbo.TableA",
	}
	for _, w := range want {
		if !strings.Contains(out, w) {
			t.Fatalf("expected SELECT INTO formatting %q not found in output:\n%s", w, out)
		}
	}
}

func TestFormatModuleDefinitionRepresentativeProcedure(t *testing.T) {
	t.Parallel()

	cfg := DefaultAppConfig(".")
	cfg.StripModuleComments = true
	cfg.ReindentModuleSQL = true
	cfg.IndentSize = 2

	in := "CREATE PROCEDURE dbo.LoadRisk AS\nBEGIN\n-- obsolete comment\nSELECT Id, Score\nINTO #Risk\nFROM dbo.Risk;\nEND\n"
	out := formatModuleDefinition(in, cfg)
	for _, want := range []string{"CREATE PROCEDURE dbo.LoadRisk", "SELECT Id, Score", "INTO #Risk", "FROM dbo.Risk"} {
		if !strings.Contains(out, want) {
			t.Fatalf("expected formatted procedure to contain %q:\n%s", want, out)
		}
	}
	if strings.Contains(out, "obsolete comment") {
		t.Fatalf("expected comments to be stripped:\n%s", out)
	}
}

func TestReindentTSQLCreateTableAndInsertParenList(t *testing.T) {
	t.Parallel()

	in := `
CREATE TABLE #tmpError

(
Tanggal datetime,
[SID] varchar(100),
NoRekInvestorRDN varchar(20)
)

INSERT INTO #tmpPTOSCustmer_TM

(
Tanggal,
[SID],
NoRekInvestorRDN
)
SELECT 1,2,3
`
	out := reindentTSQL(in, 2)
	want := []string{
		"CREATE TABLE #tmpError\n(",
		"\tTanggal datetime,",
		"\t[SID] varchar(100),",
		"\tNoRekInvestorRDN varchar(20)",
		"INSERT INTO #tmpPTOSCustmer_TM\n(",
		"\tTanggal,",
		"\t[SID],",
		"\tNoRekInvestorRDN",
	}
	for _, w := range want {
		if !strings.Contains(out, w) {
			t.Fatalf("expected create/insert paren list formatting %q not found in output:\n%s", w, out)
		}
	}
	if strings.Contains(out, "INSERT INTO #tmpPTOSCustmer_TM\n\n(") {
		t.Fatalf("blank line before insert parenthesis should be removed:\n%s", out)
	}
	if strings.Contains(out, "CREATE TABLE #tmpError\n\n(") {
		t.Fatalf("blank line before create table parenthesis should be removed:\n%s", out)
	}
}

func TestReindentTSQLUpdateSetListIndent(t *testing.T) {
	t.Parallel()

	in := `
UPDATE tmp
SET ColA = src.ColA,
ColB = src.ColB,
ColC = src.ColC
FROM dbo.Tmp tmp
JOIN dbo.Src src ON src.Id = tmp.Id
`
	out := reindentTSQL(in, 2)
	want := []string{
		"SET ColA = src.ColA,",
		"\tColB = src.ColB,",
		"\tColC = src.ColC",
		"\nFROM dbo.Tmp tmp",
	}
	for _, w := range want {
		if !strings.Contains(out, w) {
			t.Fatalf("expected update set indentation %q not found in output:\n%s", w, out)
		}
	}
}

func TestReindentTSQLInsertInlineOpenParenListIndent(t *testing.T) {
	t.Parallel()

	in := `
INSERT INTO #tmpError (
Tanggal,
[SID],
WarningMessage
)
SELECT 1,2,3
`
	out := reindentTSQL(in, 2)
	want := []string{
		"INSERT INTO #tmpError (",
		"\tTanggal,",
		"\t[SID],",
		"\tWarningMessage",
	}
	for _, w := range want {
		if !strings.Contains(out, w) {
			t.Fatalf("expected inline insert list indentation %q not found in output:\n%s", w, out)
		}
	}
}

func TestReindentTSQLCompactsQueryBlankLines(t *testing.T) {
	t.Parallel()

	in := `
SELECT a,
	b

FROM dbo.T

INNER JOIN dbo.U u
ON u.id = t.id

DELETE t

FROM dbo.T t

INNER JOIN dbo.X x ON x.id = t.id
`
	out := reindentTSQL(in, 2)
	if strings.Contains(out, "b\n\nFROM dbo.T") {
		t.Fatalf("blank line between select list and from should be removed:\n%s", out)
	}
	if strings.Contains(out, "FROM dbo.T\n\nINNER JOIN dbo.U u") {
		t.Fatalf("blank line between from and join should be removed:\n%s", out)
	}
	if strings.Contains(out, "DELETE t\n\nFROM dbo.T t") {
		t.Fatalf("blank line between delete and from should be removed:\n%s", out)
	}
	if strings.Contains(out, "ON u.id = t.id\n\nAND u.active = 1") {
		t.Fatalf("blank line between on and and should be removed:\n%s", out)
	}
}

func TestReindentTSQLNoStrayTabOnNewStatementAfterQueryChain(t *testing.T) {
	t.Parallel()

	in := `
DELETE tmpCus
from #tmpPTOSCustmer_TM tmpCus
INNER JOIN #tmpError tmp
	ON tmpCus.NoRekInvestorRDN = tmp.NoRekInvestorRDN

	INSERT INTO #tmpError (
Tanggal
)
SELECT 1
`
	out := reindentTSQL(in, 2)
	if strings.Contains(out, "\n\tINSERT INTO #tmpError (") {
		t.Fatalf("insert statement should not keep stray tab indentation:\n%s", out)
	}
	if !strings.Contains(out, "\nINSERT INTO #tmpError (") {
		t.Fatalf("insert statement expected in output:\n%s", out)
	}
}

func TestReindentTSQLInlineParenClosedSameLineMustNotLeakContext(t *testing.T) {
	t.Parallel()

	in := `
INSERT INTO #tmpError (NoRekInvestorRDN)
SELECT NoRekInvestorRDN
FROM #tmpPTOSCustmer_TM
GROUP BY NoRekInvestorRDN
`
	out := reindentTSQL(in, 2)
	if strings.Contains(out, "\n\tFROM #tmpPTOSCustmer_TM") {
		t.Fatalf("from line should not get tab from stale paren context:\n%s", out)
	}
	if strings.Contains(out, "\n\tGROUP BY NoRekInvestorRDN") {
		t.Fatalf("group by line should not get tab from stale paren context:\n%s", out)
	}
}

func TestReindentTSQLCaseEndProjectionKeepsSelectListIndent(t *testing.T) {
	t.Parallel()

	in := `
SELECT
CASE
WHEN a = 1 THEN 'A'
ELSE 'B'
END as Remark,
ColA,
ColB
from dbo.T
`
	out := reindentTSQL(in, 2)
	want := []string{
		"\tEND as Remark,",
		"\tColA,",
		"\tColB",
		"\nfrom dbo.T",
	}
	for _, w := range want {
		if !strings.Contains(out, w) {
			t.Fatalf("expected CASE END projection list style %q not found in output:\n%s", w, out)
		}
	}
}

func TestReindentTSQLBlankLineAfterSingleLineIfDoesNotOverIndentNextSelect(t *testing.T) {
	t.Parallel()

	in := `
IF(@bDebug = 1) SELECT * FROM #tmpError

SELECT
ColA,
ColB
FROM dbo.T
`
	out := reindentTSQL(in, 2)
	if strings.Contains(out, "\n      SELECT") {
		t.Fatalf("next SELECT should not get extra indent after blank line:\n%s", out)
	}
	if !strings.Contains(out, "\nSELECT\n\tColA,") {
		t.Fatalf("expected select list formatting not found:\n%s", out)
	}
}

func TestReindentTSQLCompactsBlankLinesInsideSelectProjectionList(t *testing.T) {
	t.Parallel()

	in := `
SELECT
ColA,
ColB,

UPPER(ColC) as ColC,

ColD
FROM dbo.T
`
	out := reindentTSQL(in, 2)
	if strings.Contains(out, "ColB,\n\n\tUPPER(ColC)") {
		t.Fatalf("blank line between select projection items should be removed:\n%s", out)
	}
	if strings.Contains(out, "as ColC,\n\n\tColD") {
		t.Fatalf("blank line between select projection items should be removed:\n%s", out)
	}
}

func TestFormatModuleDefinitionEdgeCaseBatch(t *testing.T) {
	t.Parallel()

	cfg := DefaultAppConfig(".")
	cfg.StripModuleComments = true
	cfg.ReindentModuleSQL = true
	cfg.IndentSize = 2

	in := `CREATE PROCEDURE dbo.SyncRisk AS
BEGIN
MERGE dbo.Target AS target
USING dbo.Source AS src
ON target.Id = src.Id
WHEN MATCHED THEN UPDATE SET target.Amount = src.Amount,
Status = src.Status
WHEN NOT MATCHED THEN INSERT (Id, Amount) VALUES (src.Id, src.Amount);
END`
	out := formatModuleDefinition(in, cfg)
	for _, want := range []string{"MERGE dbo.Target AS target", "SET target.Amount = src.Amount,", "Status = src.Status", "WHEN NOT MATCHED THEN"} {
		if !strings.Contains(out, want) {
			t.Fatalf("expected pattern %q not found:\n%s", want, out)
		}
	}
	if strings.Contains(out, ",\n\n\t") {
		t.Fatalf("found blank line inside comma list:\n%s", out)
	}
}

func TestReindentTSQLMergeSetDoesNotLeakIntoWhenClause(t *testing.T) {
	t.Parallel()

	in := `
MERGE dbo.Target AS t
USING dbo.Source AS s
ON t.Id = s.Id
WHEN MATCHED THEN
UPDATE
SET t.ColA = s.ColA,
t.ColB = s.ColB
WHEN NOT MATCHED THEN
INSERT (Id, ColA, ColB)
VALUES (s.Id, s.ColA, s.ColB)
`
	out := reindentTSQL(in, 2)
	if strings.Contains(out, "\n\tWHEN NOT MATCHED THEN") {
		t.Fatalf("merge when clause should not be indented as update-set continuation:\n%s", out)
	}
	if !strings.Contains(out, "\nWHEN NOT MATCHED THEN") {
		t.Fatalf("expected merge when clause not found:\n%s", out)
	}
}

func TestReindentTSQLApplyOnClauseNoExtraSpaces(t *testing.T) {
	t.Parallel()

	in := `
SELECT p.Id
FROM dbo.P p
OUTER APPLY (
SELECT x.Id
FROM dbo.X x
JOIN dbo.Y y
ON y.Id = x.YId
WHERE x.PId = p.Id
) a
`
	out := reindentTSQL(in, 2)
	if strings.Contains(out, "\t  ON y.Id = x.YId") {
		t.Fatalf("on clause in apply block contains extra spaces:\n%s", out)
	}
	if !strings.Contains(out, "ON y.Id = x.YId") {
		t.Fatalf("expected on clause in apply block not found:\n%s", out)
	}
}

func TestReindentTSQLMergeInsertColumnListIndented(t *testing.T) {
	t.Parallel()

	in := `
MERGE dbo.Target AS t
USING dbo.Source AS s
ON t.Id = s.Id
WHEN NOT MATCHED THEN
INSERT (
Id,
Amount,
UpdatedAt
)
VALUES (
s.Id,
s.Amount,
SYSUTCDATETIME()
)
`
	out := reindentTSQL(in, 2)
	want := "INSERT (\n\tId,\n\tAmount,\n\tUpdatedAt\n)"
	if !strings.Contains(out, want) {
		t.Fatalf("expected merge insert list indentation %q not found:\n%s", want, out)
	}
}

func TestReindentTSQLDeclareTableVariableColumnListIndented(t *testing.T) {
	t.Parallel()

	in := `
DECLARE @tmpPTOSOutstanding1_TR TABLE
(
Tanggal datetime,
[SID] varchar(100),
NoRekInvestorRDN varchar(20)
)
`
	out := reindentTSQL(in, 2)
	want := "DECLARE @tmpPTOSOutstanding1_TR TABLE\n(\n\tTanggal datetime,\n\t[SID] varchar(100),\n\tNoRekInvestorRDN varchar(20)\n)"
	if !strings.Contains(out, want) {
		t.Fatalf("expected declare table list indentation %q not found:\n%s", want, out)
	}
}

func TestReindentTSQLValuesListIndented(t *testing.T) {
	t.Parallel()

	in := `
INSERT INTO dbo.T (
Id,
Amount
)
VALUES (
s.Id,
s.Amount
)
`
	out := reindentTSQL(in, 2)
	want := "VALUES (\n\ts.Id,\n\ts.Amount\n)"
	if !strings.Contains(out, want) {
		t.Fatalf("expected values list indentation %q not found:\n%s", want, out)
	}
}
