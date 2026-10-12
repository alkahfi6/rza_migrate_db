package getobject

import "testing"

func TestQuoteIdentifierPreservesCatalogWhitespace(t *testing.T) {
	t.Parallel()

	if got := quoteIdentifier(" "); got != "[ ]" {
		t.Fatalf("single-space identifier changed: %q", got)
	}
	if got := quoteIdentifier("a] b"); got != "[a]] b]" {
		t.Fatalf("escaped identifier changed: %q", got)
	}
}
