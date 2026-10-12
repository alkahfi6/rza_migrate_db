package sqlparse

import (
	"reflect"
	"testing"
)

func TestFindMatchingParenIgnoresLiteralsCommentsAndBrackets(t *testing.T) {
	sql := "(DEFAULT (N')'), CHECK ([Odd]]Name] IN ('A,B', '(')), /* ) */ 1) tail"
	got, ok := FindMatchingParen(sql, 0)
	if !ok || got != len(sql)-len(" tail")-1 {
		t.Fatalf("unexpected closing paren: got=%d ok=%v", got, ok)
	}
}

func TestSplitTopLevelCommaIsSQLAware(t *testing.T) {
	input := "a INT, b VARCHAR(20) DEFAULT ('A,B'), c AS CONCAT('(', a), [x,y] INT /* , */"
	got := SplitTopLevelComma(input)
	want := []string{"a INT", " b VARCHAR(20) DEFAULT ('A,B')", " c AS CONCAT('(', a)", " [x,y] INT /* , */"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("split mismatch:\ngot=%#v\nwant=%#v", got, want)
	}
}

func TestRewriteIdentifierExactPreservesCaseLiteralsAndComments(t *testing.T) {
	input := "[OldName] + OldName + oldname + 'OldName' -- OldName\n/* OldName */ [Odd]]Name]"
	got := RewriteIdentifierExact(input, "OldName", "NewName")
	want := "[NewName] + NewName + oldname + 'OldName' -- OldName\n/* OldName */ [Odd]]Name]"
	if got != want {
		t.Fatalf("rewrite mismatch:\n%s", got)
	}
}
