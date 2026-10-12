package tsqlast

import (
	"strings"
	"testing"
)

func TestReindentTSQLPreservesBlockCommentLayoutAndIgnoresKeywords(t *testing.T) {
	t.Parallel()
	in := "CREATE OR ALTER PROCEDURE dbo.ProcA\nAS\nBEGIN\n/*\n               BEGIN data saya ...\n       CASE WHEN x = 1 THEN END\n*/\nSELECT 1;\nEND\n"
	out := ReindentTSQL(in, 2)
	wantComment := "/*\n               BEGIN data saya ...\n       CASE WHEN x = 1 THEN END\n*/"
	if !strings.Contains(out, wantComment) {
		t.Fatalf("block comment layout changed:\n%s", out)
	}
	if !strings.Contains(out, "*/\n  SELECT 1;\nEND") {
		t.Fatalf("keywords inside block comment affected SQL indentation:\n%s", out)
	}
}

func TestReindentTSQLPreservesLineCommentIndentAndIgnoresBeginEnd(t *testing.T) {
	t.Parallel()
	in := "BEGIN\n       -- history BEGIN CASE END\nSELECT 1;\nEND\n"
	out := ReindentTSQL(in, 2)
	if !strings.Contains(out, "\n       -- history BEGIN CASE END\n") {
		t.Fatalf("line comment indentation changed:\n%s", out)
	}
	if !strings.Contains(out, "\n  SELECT 1;\nEND\n") {
		t.Fatalf("line comment keywords affected SQL indentation:\n%s", out)
	}
}

func TestReindentTSQLPreservesMultilineDynamicSQLLiteral(t *testing.T) {
	t.Parallel()
	in := "BEGIN\nSET @sql = N'SELECT 1\n               -- literal text BEGIN\n    FROM dbo.T\n'\nEXEC dbo.sp_executesql @sql;\nEND\n"
	out := ReindentTSQL(in, 2)
	wantLiteral := "SET @sql = N'SELECT 1\n               -- literal text BEGIN\n    FROM dbo.T\n'"
	if !strings.Contains(out, wantLiteral) {
		t.Fatalf("multiline literal layout changed:\n%s", out)
	}
	if !strings.Contains(out, "\n  EXEC dbo.sp_executesql @sql;\nEND\n") {
		t.Fatalf("multiline literal content affected following SQL indentation:\n%s", out)
	}
}

func TestReindentTSQLTriviaAwareIsIdempotent(t *testing.T) {
	t.Parallel()
	in := "BEGIN\n/* outer\n     /* nested BEGIN */\n  END marker\n*/\nIF EXISTS (SELECT 1)\nSELECT N'x';\nEND\n"
	first := ReindentTSQL(in, 2)
	second := ReindentTSQL(first, 2)
	if first != second {
		t.Fatalf("trivia-aware formatting must be idempotent\nfirst:\n%s\nsecond:\n%s", first, second)
	}
}

func TestContainsSQLCommentIgnoresMarkersInsideLiteral(t *testing.T) {
	t.Parallel()
	if ContainsSQLComment("SELECT N'-- not a comment /* x */';") {
		t.Fatal("comment markers inside literal must not be classified as comments")
	}
	if !ContainsSQLComment("SELECT 1; /* real comment */") {
		t.Fatal("real block comment was not detected")
	}
}
