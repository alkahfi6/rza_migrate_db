package sqlparse

import (
	"reflect"
	"testing"
)

func TestExtractSQLAccessesPrecision(t *testing.T) {
	tests := []struct {
		name, sql string
		want      []string
	}{
		{"quoted keywords", `SELECT [FROM], [UPDATE], [SELECT] FROM [dbo].[Order Items];`, []string{"READ|dbo.Order Items"}},
		{"quoted optional keyword targets", `INSERT [INTO](Id) VALUES(1); DELETE [FROM] WHERE Id=1; MERGE [INTO] t USING dbo.Source s ON t.Id=s.Id WHEN MATCHED THEN DELETE;`, []string{"INSERT|INTO", "DELETE|FROM", "MERGE|INTO", "READ|dbo.Source"}},
		{"unfinished identifier", `SELECT 1 FROM [`, nil},
		{"top targets", `UPDATE TOP (2) dbo.T SET Id=1; DELETE TOP(1) FROM dbo.U; INSERT TOP(3) INTO dbo.V(Id) VALUES(1);`, []string{"UPDATE|dbo.T", "DELETE|dbo.U", "INSERT|dbo.V"}},
		{"merge", `MERGE dbo.Target AS t USING dbo.Source AS s ON t.Id=s.Id WHEN MATCHED THEN UPDATE SET Id=s.Id WHEN NOT MATCHED THEN INSERT(Id) VALUES(s.Id) WHEN NOT MATCHED BY SOURCE THEN DELETE;`, []string{"MERGE|dbo.Target", "READ|dbo.Source"}},
		{"select into", `SELECT Id INTO #work FROM dbo.T;`, []string{"INSERT|#work", "READ|dbo.T"}},
		{"comma and apply", `SELECT * FROM dbo.A a, dbo.B b CROSS APPLY dbo.Rows_FN(a.Id) f;`, []string{"READ|dbo.A", "READ|dbo.B", "READ|dbo.Rows_FN"}},
		{"literal barrier", `SELECT 'FROM dbo.Fake' FROM dbo.Real;`, []string{"READ|dbo.Real"}},
		{"cursor", `FETCH ABSOLUTE -1 FROM cursor1 INTO @id;`, nil},
		{"alias shadow", `UPDATE t SET Id=1 FROM dbo.Real t WHERE EXISTS(SELECT 1 FROM dbo.Other t);`, []string{"UPDATE|dbo.Real", "READ|dbo.Real", "READ|dbo.Other"}},
		{"hint boundary", `UPDATE t SET Id=1 FROM dbo.Real t WITH (NOLOCK);`, []string{"UPDATE|dbo.Real", "READ|dbo.Real"}},
		{"derived and sibling", `UPDATE d SET Id=1 FROM (SELECT Id FROM dbo.Real) d, dbo.Other o;`, []string{"UPDATE|d", "READ|d", "READ|dbo.Real", "READ|dbo.Other"}},
		{"output sink", `DELETE dbo.T OUTPUT deleted.Id INTO #audit;`, []string{"DELETE|dbo.T", "INSERT|#audit"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string
			for _, a := range ExtractSQLAccesses(tt.sql) {
				got = append(got, a.Operation+"|"+a.Name)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasClientResultSetCompoundStatements(t *testing.T) {
	for _, sql := range []string{
		`INSERT dbo.T SELECT Id FROM dbo.A UNION ALL SELECT Id FROM dbo.B;`,
		`SELECT Id INTO #x FROM dbo.A UNION ALL SELECT Id FROM dbo.B;`,
		`SELECT @x=Id FROM dbo.A UNION SELECT Id FROM dbo.B;`,
		`DECLARE c CURSOR FOR SELECT Id FROM dbo.A;`,
		`SELECT 1 WHERE EXISTS(SELECT [SELECT] FROM dbo.A);`,
	} {
		want := sql == `SELECT 1 WHERE EXISTS(SELECT [SELECT] FROM dbo.A);`
		if got := HasClientResultSet(sql); got != want {
			t.Errorf("%s: got %v want %v", sql, got, want)
		}
	}
}

func FuzzSQLAccessSourceBounds(f *testing.F) {
	for _, seed := range []string{`SELECT * FROM [`, `SELECT * FROM "`, `UPDATE t SET Id=1 FROM dbo.T t;`, `WITH x AS(SELECT 1) SELECT * FROM x`, `MERGE dbo.T USING dbo.U ON 1=1 WHEN MATCHED THEN DELETE;`, `SELECT 'it''s a FROM';`} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, sql string) {
		for _, a := range ExtractSQLAccesses(sql) {
			if a.Offset < 0 || a.End < a.Offset || a.End > len(sql) || a.StatementStart > a.Offset || a.StatementEnd < a.End || a.StatementEnd > len(sql) {
				t.Fatalf("invalid source bounds: %+v", a)
			}
		}
		for _, c := range ExtractSQLCalls(sql) {
			if c.Offset < 0 || c.TargetEnd < c.Offset || c.StatementEnd < c.TargetEnd || c.StatementEnd > len(sql) {
				t.Fatalf("invalid call bounds: %+v", c)
			}
		}
		HasClientResultSet(sql)
		ExtractObjectReferences(sql)
	})
}

func TestExtractSQLAccessesHandlesMultilineObjectsAndScopedCTE(t *testing.T) {
	sql := `WITH T AS (SELECT Id FROM dbo.Source)
SELECT * FROM T;
SELECT *
FROM
dbo.T;`
	got := ExtractSQLAccesses(sql)
	want := map[string]bool{
		"READ|dbo.Source|false": false,
		"READ|T|true":           false,
		"READ|dbo.T|false":      false,
	}
	for _, access := range got {
		key := access.Operation + "|" + access.Name + "|"
		if access.NonPersisted {
			key += "true"
		} else {
			key += "false"
		}
		if _, ok := want[key]; ok {
			want[key] = true
		}
	}
	for key, found := range want {
		if !found {
			t.Fatalf("missing %s in %+v", key, got)
		}
	}
}

func TestExtractSQLAccessesKeepsStatementAliasScope(t *testing.T) {
	sql := `UPDATE t SET Value = 1
FROM dbo.First AS t;
DELETE t
FROM dbo.Second t;`
	got := ExtractSQLAccesses(sql)
	if len(got) != 4 {
		t.Fatalf("expected four accesses, got %+v", got)
	}
	if got[1].Name != "dbo.First" || got[1].Alias != "t" || got[0].StatementEnd != got[1].StatementEnd {
		t.Fatalf("UPDATE alias scope mismatch: %+v", got[:2])
	}
	if got[3].Name != "dbo.Second" || got[3].Alias != "t" || got[2].StatementStart != got[3].StatementStart {
		t.Fatalf("DELETE alias scope mismatch: %+v", got[2:])
	}
}

func TestExtractSQLAccessesKeepsTemporaryAndTableVariableIdentity(t *testing.T) {
	got := ExtractSQLAccesses("INSERT INTO @work(Id) VALUES (1); SELECT * FROM #temp;")
	if len(got) != 2 || got[0].Name != "@work" || got[1].Name != "#temp" {
		t.Fatalf("temporary identities were not preserved: %+v", got)
	}
}

func TestHasClientResultSetUsesStatementSemantics(t *testing.T) {
	tests := []struct {
		name string
		sql  string
		want bool
	}{
		{"client", "SELECT Id FROM dbo.T;", true},
		{"select into", "SELECT Id INTO #work FROM dbo.T;", false},
		{"assignment", "SELECT @id = Id FROM dbo.T;", false},
		{"exists", "IF EXISTS (SELECT 1 FROM dbo.T) RETURN;", false},
		{"insert select", "INSERT dbo.Log(Id) SELECT Id FROM dbo.T;", false},
		{"cte client", "WITH x AS (SELECT Id FROM dbo.T) SELECT Id FROM x;", true},
		{"output client", "DELETE dbo.T OUTPUT deleted.Id;", true},
		{"output sink", "DELETE dbo.T OUTPUT deleted.Id INTO #audit;", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasClientResultSet(tt.sql); got != tt.want {
				t.Fatalf("HasClientResultSet()=%v want %v", got, tt.want)
			}
		})
	}
}
