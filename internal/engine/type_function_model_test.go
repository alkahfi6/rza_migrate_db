package engine

import (
	"strings"
	"testing"
)

func TestClassifyFunctionSubtypeMatrix(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		sql  string
		want string
	}{
		{
			name: "scalar",
			sql:  `CREATE OR ALTER FUNCTION dbo.fnScalar(@id INT = 1) RETURNS INT WITH SCHEMABINDING AS BEGIN RETURN @id; END;`,
			want: functionSubtypeScalar,
		},
		{
			name: "inline_table",
			sql:  `CREATE OR ALTER FUNCTION dbo.fnInline(@id INT) RETURNS TABLE AS RETURN (SELECT @id AS id);`,
			want: functionSubtypeInlineTable,
		},
		{
			name: "multi_statement_table",
			sql: `CREATE OR ALTER FUNCTION dbo.fnMulti(@id INT)
RETURNS @result TABLE (id INT NOT NULL PRIMARY KEY)
AS
BEGIN
  INSERT @result(id) VALUES (@id);
  RETURN;
END;`,
			want: functionSubtypeMultiStatementTV,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := validateFunctionSource(tt.sql)
			if err != nil {
				t.Fatalf("validate function: %v", err)
			}
			if got != tt.want {
				t.Fatalf("subtype mismatch: got=%s want=%s", got, tt.want)
			}
		})
	}
}

func TestValidateFunctionSourceRejectsEncryptionAndMalformedTVF(t *testing.T) {
	t.Parallel()

	for name, sql := range map[string]string{
		"encrypted":       `CREATE OR ALTER FUNCTION dbo.fnSecret() RETURNS INT WITH ENCRYPTION AS BEGIN RETURN 1; END;`,
		"multi_no_table":  `CREATE OR ALTER FUNCTION dbo.fnBad() RETURNS @result INT AS BEGIN RETURN; END;`,
		"multi_no_block":  `CREATE OR ALTER FUNCTION dbo.fnBad() RETURNS @result TABLE (id INT) AS RETURN;`,
		"missing_returns": `CREATE OR ALTER FUNCTION dbo.fnBad() AS BEGIN RETURN 1; END;`,
	} {
		name, sql := name, sql
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if _, err := validateFunctionSource(sql); err == nil {
				t.Fatalf("expected invalid function source: %s", sql)
			}
		})
	}
}

func TestSemanticFromSourceSupportsTableType(t *testing.T) {
	t.Parallel()

	source := `-- order items passed by callers
CREATE TYPE dbo.OrderItemType AS TABLE (
  ProductId INT NOT NULL,
  Qty DECIMAL(18,4) NOT NULL,
  PRIMARY KEY CLUSTERED ([ProductId]),
  CHECK ([Qty] > 0)
);`
	got, err := semanticFromSource(migrationMeta{Category: "types", Object: "dbo.OrderItemType"}, source)
	if err != nil {
		t.Fatalf("semanticFromSource type failed: %v", err)
	}
	if !strings.Contains(got, "CREATE TYPE dbo.OrderItemType AS TABLE") ||
		!strings.Contains(got, "PRIMARY KEY CLUSTERED") {
		t.Fatalf("unexpected table type semantic:\n%s", got)
	}
	if _, err := semanticFromMigration(migrationMeta{Category: "types", Object: "dbo.OrderItemType"}, got); err != nil {
		t.Fatalf("type migration reconstruction failed: %v", err)
	}
}

func TestTableTypeValidationRejectsSQLServerInvalidNamedConstraint(t *testing.T) {
	t.Parallel()

	_, err := validateTableTypeSource(
		"dbo.OrderItemType",
		`CREATE TYPE dbo.OrderItemType AS TABLE (
  Id INT NOT NULL,
  CONSTRAINT PK_OrderItemType PRIMARY KEY (Id)
);`,
	)
	if err == nil || !strings.Contains(strings.ToLower(err.Error()), "constraint names") {
		t.Fatalf("expected named table-type constraint rejection, got %v", err)
	}
}

func TestBuildTableTypeLifecycleDropsAndRecreatesDependentModules(t *testing.T) {
	t.Parallel()

	oldType := `CREATE TYPE dbo.OrderItemType AS TABLE ([Id] INT NOT NULL);`
	newType := `CREATE TYPE dbo.OrderItemType AS TABLE ([Id] INT NOT NULL, [Qty] INT NULL);`
	oldFunction := `CREATE OR ALTER FUNCTION dbo.fnCountItems(@items dbo.OrderItemType READONLY) RETURNS INT AS BEGIN RETURN (SELECT COUNT(*) FROM @items); END;`
	newFunction := `CREATE OR ALTER FUNCTION dbo.fnCountItems(@items dbo.OrderItemType READONLY) RETURNS INT AS BEGIN RETURN (SELECT COUNT_BIG(*) FROM @items); END;`
	oldProcedure := `CREATE OR ALTER PROCEDURE dbo.SaveItems @items dbo.OrderItemType READONLY AS SELECT dbo.fnCountItems(@items);`
	newProcedure := `CREATE OR ALTER PROCEDURE dbo.SaveItems @items dbo.OrderItemType READONLY AS SELECT dbo.fnCountItems(@items) AS item_count;`

	history := map[string]historyObject{
		objectKey("functions", "dbo.fnCountItems"): {Category: "functions", Object: "dbo.fnCountItems", SemanticSQL: oldFunction},
		objectKey("procedures", "dbo.SaveItems"):   {Category: "procedures", Object: "dbo.SaveItems", SemanticSQL: oldProcedure},
	}
	desired := map[string]desiredObject{
		objectKey("types", "dbo.OrderItemType"):    {Category: "types", Object: "dbo.OrderItemType", SemanticSQL: newType},
		objectKey("functions", "dbo.fnCountItems"): {Category: "functions", Object: "dbo.fnCountItems", SemanticSQL: newFunction},
		objectKey("procedures", "dbo.SaveItems"):   {Category: "procedures", Object: "dbo.SaveItems", SemanticSQL: newProcedure},
	}

	up, down, dependents, err := buildTableTypeLifecycleSQL(
		"dbo.OrderItemType", oldType, newType, history, desired,
		[]string{"types", "tables", "indexes", "functions", "views", "procedures", "patch"},
	)
	if err != nil {
		t.Fatalf("build lifecycle: %v", err)
	}
	if len(dependents) != 2 {
		t.Fatalf("expected two dependents, got=%v", dependents)
	}

	assertOrder := func(label, sql string, tokens ...string) {
		t.Helper()
		last := -1
		for _, token := range tokens {
			pos := strings.Index(sql, token)
			if pos < 0 {
				t.Fatalf("%s missing %q:\n%s", label, token, sql)
			}
			if pos <= last {
				t.Fatalf("%s token order invalid at %q:\n%s", label, token, sql)
			}
			last = pos
		}
	}
	assertOrder("up", up,
		"DROP PROCEDURE IF EXISTS dbo.SaveItems",
		"DROP FUNCTION IF EXISTS dbo.fnCountItems",
		"DROP TYPE IF EXISTS dbo.OrderItemType",
		newType,
	)
	if strings.Contains(up, "EXEC sys.sp_executesql") {
		t.Fatalf("changed dependents must be deferred to their ordered migrations:\n%s", up)
	}
	assertOrder("down", down,
		"DROP PROCEDURE IF EXISTS dbo.SaveItems",
		"DROP FUNCTION IF EXISTS dbo.fnCountItems",
		"DROP TYPE IF EXISTS dbo.OrderItemType",
		oldType,
		oldFunction,
		oldProcedure,
	)
}
