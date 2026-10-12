package sqlparse

import (
	"strings"
	"testing"
)

func TestDetectForbiddenTokensIgnoresCommentsAndStrings(t *testing.T) {
	t.Parallel()

	sql := `
-- USE master;
/* CREATE LOGIN app_login */
SELECT 'GO';
SELECT 'CREATE USER app_user';
`

	got := DetectForbiddenTokens(sql)
	if len(got) != 0 {
		t.Fatalf("expected no forbidden tokens, got: %#v", got)
	}
}

func TestDetectForbiddenTokensDetectsRealTokens(t *testing.T) {
	t.Parallel()

	sql := `
USE master;
CREATE USER app_user FOR LOGIN app_login;
GO
`

	got := DetectForbiddenTokens(sql)
	expected := map[string]bool{
		"USE":         true,
		"CREATE USER": true,
		"GO":          true,
	}
	if len(got) != len(expected) {
		t.Fatalf("unexpected token count: got=%d want=%d tokens=%#v", len(got), len(expected), got)
	}
	for _, token := range got {
		if !expected[token] {
			t.Fatalf("unexpected token detected: %s", token)
		}
	}
}

func TestExtractDeclaredIntegerLiteralRequiresLiteral(t *testing.T) {
	t.Parallel()

	literalSQL := strings.ToUpper(`
DECLARE @version BIGINT = 20260217073015000;
`)
	if got, ok := ExtractDeclaredIntegerLiteral(literalSQL, "@VERSION"); !ok || got != "20260217073015000" {
		t.Fatalf("expected integer literal extraction, got value=%q ok=%v", got, ok)
	}

	nonLiteralSQL := strings.ToUpper(`
DECLARE @version BIGINT = ABS(CHECKSUM(NEWID()));
`)
	if got, ok := ExtractDeclaredIntegerLiteral(nonLiteralSQL, "@VERSION"); ok {
		t.Fatalf("expected non-literal extraction to fail, got value=%q", got)
	}
}

func TestExtractDeclaredNVarcharLiteralAndNamedParams(t *testing.T) {
	t.Parallel()

	upperSQL := strings.ToUpper(`
DECLARE @category NVARCHAR(16) = N'tables';
EXEC dbo.usp_mig_log_event
  @decision = N'run',
  @severity = 1,
  @reason = N'OBJECT_EXISTS';
EXEC dbo.usp_mig_log_event
  @decision = N'ok',
  @severity = 3,
  @reason = N'SQL_ERROR';
`)

	if got, ok := ExtractDeclaredNVarcharLiteral(upperSQL, "@CATEGORY"); !ok || got != "TABLES" {
		t.Fatalf("expected declared NVARCHAR literal TABLES, got value=%q ok=%v", got, ok)
	}

	decisions := ExtractNamedNVarcharParamLiterals(upperSQL, "@DECISION")
	if len(decisions) != 2 || decisions[0] != "RUN" || decisions[1] != "OK" {
		t.Fatalf("unexpected decisions extraction: %#v", decisions)
	}

	severities := ExtractNamedIntParamLiterals(upperSQL, "@SEVERITY")
	if len(severities) != 2 || severities[0] != 1 || severities[1] != 3 {
		t.Fatalf("unexpected severities extraction: %#v", severities)
	}
}

func TestExtractDeclaredNVarcharLiteralPreservesExactCase(t *testing.T) {
	t.Parallel()
	sqlText := "declare @target nvarchar(256) = N'dbo.CaseSensitiveObject';\n"
	got, ok := ExtractDeclaredNVarcharLiteral(sqlText, "@TARGET")
	if !ok || got != "dbo.CaseSensitiveObject" {
		t.Fatalf("case-preserving literal mismatch: got=%q ok=%v", got, ok)
	}
}

func TestContainsStandaloneWordBoundary(t *testing.T) {
	t.Parallel()

	if !ContainsStandaloneWord(" USE MASTER ", "USE") {
		t.Fatalf("expected USE to be detected as standalone word")
	}
	if ContainsStandaloneWord("REUSE MASTER", "USE") {
		t.Fatalf("did not expect USE to match REUSE")
	}
}

func TestDetectForbiddenTokensWithList_CustomTokens(t *testing.T) {
	t.Parallel()

	sql := `
USE master;
CREATE USER app_user FOR LOGIN app_login;
`

	got := DetectForbiddenTokensWithList(sql, []string{"USE"})
	if len(got) != 1 || got[0] != "USE" {
		t.Fatalf("expected only USE token, got %#v", got)
	}
}

func TestFindStatementEndIgnoresCommentAndStringSemicolons(t *testing.T) {
	t.Parallel()

	sql := "CREATE INDEX IX_T__A ON dbo.T(a) -- ; hidden in line comment\nINCLUDE (b);\n"
	end := FindStatementEnd(sql, 0)
	want := strings.LastIndex(sql, ";") + 1
	if end != want {
		t.Fatalf("unexpected statement end: got=%d want=%d", end, want)
	}

	noSemi := "CREATE INDEX IX_T__B ON dbo.T(b)\n"
	endNoSemi := FindStatementEnd(noSemi, 0)
	if endNoSemi != len(noSemi) {
		t.Fatalf("expected end=len(sql) when semicolon is absent, got=%d len=%d", endNoSemi, len(noSemi))
	}
}
