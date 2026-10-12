package sqlparse

import "testing"

func TestCanonicalSemanticSQLPreservesIdentifierCase(t *testing.T) {
	a := CanonicalSemanticSQL("create table [dbo].[Thing] ([Code] int);")
	b := CanonicalSemanticSQL("CREATE TABLE dbo.Thing (Code INT)")
	c := CanonicalSemanticSQL("CREATE TABLE dbo.Thing (code INT)")
	if a != b {
		t.Fatalf("keyword/bracket formatting should be equivalent:\n%s\n%s", a, b)
	}
	if a == c {
		t.Fatal("identifier case must be semantically significant")
	}
}

func TestCanonicalSemanticSQLIgnoresCommentsButPreservesLiterals(t *testing.T) {
	a := CanonicalSemanticSQL("SELECT N'Code' /* ) */ FROM dbo.T -- x\nWHERE [Code] = 1")
	b := CanonicalSemanticSQL("select N'Code' from [dbo].[T] where Code=1;")
	c := CanonicalSemanticSQL("select N'code' from dbo.T where Code=1")
	if a != b {
		t.Fatalf("comments and spacing should be ignored:\n%s\n%s", a, b)
	}
	if a == c {
		t.Fatal("literal case must remain significant")
	}
}
