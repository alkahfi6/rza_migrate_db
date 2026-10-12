package sqlfmt

import "testing"

func TestNormalizeSQLTextReplacesNonBreakingSpaceSeparators(t *testing.T) {
	input := "tax_interest\u00A0\u00A0money NOT NULL\n"
	got := NormalizeSQLText(input, true, true)
	want := "tax_interest  money NOT NULL\n"
	if got != want {
		t.Fatalf("NormalizeSQLText NBSP mismatch: got %q want %q", got, want)
	}
}

func TestNormalizeSQLTextReplacesMojibakeNonBreakingSpace(t *testing.T) {
	input := "tax_interest\u00C2\u00A0money NOT NULL\n"
	got := NormalizeSQLText(input, true, true)
	want := "tax_interest money NOT NULL\n"
	if got != want {
		t.Fatalf("NormalizeSQLText mojibake NBSP mismatch: got %q want %q", got, want)
	}
}
