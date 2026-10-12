package sqlparse

import "testing"

func TestDocumentationDDLUsesSharedIdentifierTokens(t *testing.T) {
	m := ReadTableMetadata(`/* CREATE TABLE wrong (Id int) */ CREATE TABLE [A.B] ([a]]b] DECIMAL(18, 4) NULL, CONSTRAINT [KEY_UQ] UNIQUE ([a]]b] DESC));`)
	if m.Name != "dbo.A.B" || len(m.Columns) != 1 || m.Columns[0].Name != "a]b" || m.Columns[0].DataType != "DECIMAL(18, 4)" || len(m.Keys) != 1 || m.Keys[0].Columns[0] != "a]b" {
		t.Fatalf("quoted identity/type lost: %+v", m)
	}
	indexes := ReadIndexMetadata(`CREATE UNIQUE INDEX [x] ON [A.B] ([a]]b]) WHERE [a]]b] IS NOT NULL; CREATE UNIQUE INDEX bad ON [A.B] ([a]]b]+Other);`)
	if len(indexes) != 1 || indexes[0].Table != "dbo.A.B" || !indexes[0].Filtered {
		t.Fatalf("index expression became a column list: %+v", indexes)
	}
}

func TestDocumentationDDLMultipleAlterForeignKeys(t *testing.T) {
	m := ReadTableMetadata(`CREATE TABLE dbo.A (Id INT, Code INT); ALTER TABLE dbo.A ADD CONSTRAINT FK_A FOREIGN KEY(Id) REFERENCES dbo.B(Id), CONSTRAINT FK_Code FOREIGN KEY(Code) REFERENCES dbo.C(Code);`)
	if len(m.ForeignKeys) != 2 {
		t.Fatalf("multi-constraint ALTER lost FK: %+v", m.ForeignKeys)
	}
}

func TestSQLTokensKeepQuotedKeywordsAndOpaqueStrings(t *testing.T) {
	tokens := SQLTokens(`SELECT [JOIN], "WHERE", N'FROM dbo.Fake' FROM dbo.Real;`)
	quoted, from := 0, 0
	for _, token := range tokens {
		if token.Quoted {
			quoted++
		}
		if !token.Quoted && token.Upper == "FROM" {
			from++
		}
	}
	if quoted != 2 || from != 1 {
		t.Fatalf("quoted/string keywords leaked: %+v", tokens)
	}
}

func TestMergeJoinHintIsNotDMLOrStatementBoundary(t *testing.T) {
	sql := `SELECT a.Id FROM dbo.A a LEFT OUTER MERGE JOIN dbo.B b ON a.Id=b.Id WHERE b.Id=1;`
	accesses := ExtractSQLAccesses(sql)
	if len(accesses) != 2 {
		t.Fatalf("join access lost: %+v", accesses)
	}
	for _, access := range accesses {
		if access.Operation != "READ" || access.StatementStart != 0 || access.StatementEnd != len(sql) {
			t.Fatalf("join hint split statement: %+v", access)
		}
	}
	refs := ExtractObjectReferences(sql)
	if len(refs) != 2 {
		t.Fatalf("cross-tool dependencies differ: %+v", refs)
	}
}
