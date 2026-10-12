package engine

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	_ "github.com/microsoft/go-mssqldb"
)

// TestIntegration_SQLServerGeneratedWrapperCompileAndIntrospect validates the
// rendered SQL against a real SQL Server without retaining any sandbox state.
// Set RZA_MSSQL_TEST_DSN in CI to enable it.
func TestIntegration_SQLServerGeneratedWrapperCompileAndIntrospect(t *testing.T) {
	dsn := os.Getenv("RZA_MSSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("RZA_MSSQL_TEST_DSN is not configured")
	}
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()
	name := fmt.Sprintf("rza_sandbox_model_%d", time.Now().UnixNano())
	sqlText := renderTxWrapper("20991231235959000", "tables", "20991231235959000_tables_dbo."+name+"_v1.up.sql", "dbo."+name, true, "CREATE TABLE dbo."+name+" ([id] INT NOT NULL CONSTRAINT PK_"+name+" PRIMARY KEY);")
	if _, err := tx.ExecContext(ctx, sqlText); err != nil {
		t.Fatalf("generated SQL did not compile/apply: %v", err)
	}
	var count int
	if err := tx.QueryRowContext(ctx, "SELECT COUNT(1) FROM sys.columns WHERE object_id = OBJECT_ID(N'dbo."+name+"') AND name = N'id'").Scan(&count); err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatalf("introspection mismatch: id column count=%d", count)
	}
}

func TestIntegration_SQLServerDesiredStateRemovalAndEquivalentSkip(t *testing.T) {
	dsn := os.Getenv("RZA_MSSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("RZA_MSSQL_TEST_DSN is not configured")
	}
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	suffix := fmt.Sprintf("%d", time.Now().UnixNano())
	empty := "it_remove_empty_" + suffix
	used := "it_remove_used_" + suffix
	parent := "it_remove_parent_" + suffix
	child := "it_remove_child_" + suffix
	dynamic := "it_remove_dynamic_" + suffix
	dynamicProc := "it_remove_dynamic_proc_" + suffix
	fk := "FK_" + child
	targets := []string{child, parent, dynamic, used, empty}
	defer func() {
		_, _ = db.ExecContext(context.Background(), "DROP PROCEDURE IF EXISTS dbo.["+dynamicProc+"];")
		for _, name := range targets {
			_, _ = db.ExecContext(context.Background(), "DROP TABLE IF EXISTS dbo.["+name+"];")
		}
	}()

	plainCreate := "CREATE TABLE dbo.[" + empty + "] ([Id] BIGINT NOT NULL, [Code] NVARCHAR(40) NULL);"
	createWrapped := renderTxWrapper(
		"20991231235959100", "tables",
		"20991231235959100_tables_dbo."+empty+"_v1.up.sql",
		"dbo."+empty, true, plainCreate,
	)
	if _, err := db.ExecContext(ctx, createWrapped); err != nil {
		t.Fatalf("initial plain table create: %#v\n%s", err, createWrapped)
	}
	// Running the exact same initial migration again must prove catalog
	// equivalence and return without executing CREATE TABLE.
	if _, err := db.ExecContext(ctx, createWrapped); err != nil {
		t.Fatalf("equivalent table did not skip safely: %v", err)
	}

	indexBody := buildCreateIndexIfNotExistsStatement(
		"dbo."+empty,
		"IX_"+empty+"_Code",
		"CREATE NONCLUSTERED INDEX [IX_"+empty+"_Code] ON dbo.["+empty+"] ([Code] DESC) WITH (FILLFACTOR = 90);",
	)
	indexWrapped := renderTxWrapper(
		"20991231235959101", "indexes",
		"20991231235959101_indexes_dbo."+empty+"_v1.up.sql",
		"dbo."+empty, true, indexBody,
	)
	if _, err := db.ExecContext(ctx, indexWrapped); err != nil {
		t.Fatalf("initial index create: %v", err)
	}
	if _, err := db.ExecContext(ctx, indexWrapped); err != nil {
		t.Fatalf("equivalent index did not skip safely: %v", err)
	}

	// Remove the explicit index first, as the lifecycle planner orders it.
	if _, err := db.ExecContext(ctx, "DROP INDEX IF EXISTS [IX_"+empty+"_Code] ON dbo.["+empty+"];"); err != nil {
		t.Fatal(err)
	}
	dropEmpty := renderTxWrapper(
		"20991231235959102", "tables",
		"20991231235959102_tables_dbo."+empty+"_v2.up.sql",
		"dbo."+empty, true, buildDropTableSQL("dbo."+empty),
	)
	if _, err := db.ExecContext(ctx, dropEmpty); err != nil {
		t.Fatalf("empty unreferenced table removal: %v", err)
	}

	if _, err := db.ExecContext(ctx, "CREATE TABLE dbo.["+used+"] ([Id] INT NOT NULL); INSERT INTO dbo.["+used+"] VALUES (1);"); err != nil {
		t.Fatal(err)
	}
	dropUsed := renderTxWrapper(
		"20991231235959103", "tables",
		"20991231235959103_tables_dbo."+used+"_v2.up.sql",
		"dbo."+used, true, buildDropTableSQL("dbo."+used),
	)
	if _, err := db.ExecContext(ctx, dropUsed); err == nil {
		t.Fatal("table containing data must not be removed")
	}
	var usedRows int
	if err := db.QueryRowContext(ctx, "SELECT COUNT(1) FROM dbo.["+used+"]").Scan(&usedRows); err != nil || usedRows != 1 {
		t.Fatalf("blocked data table was changed: rows=%d err=%v", usedRows, err)
	}

	if _, err := db.ExecContext(ctx,
		"CREATE TABLE dbo.["+parent+"] ([Id] INT NOT NULL CONSTRAINT [PK_"+parent+"] PRIMARY KEY);"+
			"CREATE TABLE dbo.["+child+"] ([ParentId] INT NULL CONSTRAINT ["+fk+"] FOREIGN KEY REFERENCES dbo.["+parent+"] ([Id]));",
	); err != nil {
		t.Fatal(err)
	}
	dropReferenced := renderTxWrapper(
		"20991231235959104", "tables",
		"20991231235959104_tables_dbo."+parent+"_v2.up.sql",
		"dbo."+parent, true, buildDropTableSQL("dbo."+parent),
	)
	if _, err := db.ExecContext(ctx, dropReferenced); err == nil {
		t.Fatal("referenced table must not be removed")
	}
	var parentExists int
	if err := db.QueryRowContext(ctx, "SELECT COUNT(1) FROM sys.tables WHERE object_id=OBJECT_ID(@p1)", "dbo."+parent).Scan(&parentExists); err != nil || parentExists != 1 {
		t.Fatalf("blocked referenced table was changed: exists=%d err=%v", parentExists, err)
	}

	if _, err := db.ExecContext(ctx, "CREATE TABLE dbo.["+dynamic+"] ([Id] INT NULL);"); err != nil {
		t.Fatal(err)
	}
	if _, err := db.ExecContext(ctx, "CREATE PROCEDURE dbo.["+dynamicProc+"] AS BEGIN EXEC(N'SELECT * FROM dbo.["+dynamic+"]'); END;"); err != nil {
		t.Fatal(err)
	}
	dropDynamic := renderTxWrapper(
		"20991231235959105", "tables",
		"20991231235959105_tables_dbo."+dynamic+"_v2.up.sql",
		"dbo."+dynamic, true, buildDropTableSQL("dbo."+dynamic),
	)
	if _, err := db.ExecContext(ctx, dropDynamic); err == nil {
		t.Fatal("table referenced through dynamic SQL text must not be removed")
	}
}

func TestIntegration_SQLServerInitialProgrammabilityDownRemovesObject(t *testing.T) {
	dsn := os.Getenv("RZA_MSSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("RZA_MSSQL_TEST_DSN is not configured")
	}
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	suffix := fmt.Sprintf("%d", time.Now().UnixNano())
	tests := []struct {
		category string
		object   string
		kind     string
		create   func(string) string
	}{
		{"views", "dbo.rza_CaseView_" + suffix, "V", func(name string) string {
			return "CREATE OR ALTER VIEW " + name + " AS SELECT 1 AS ExactCase;"
		}},
		{"functions", "dbo.rza_CaseFunction_" + suffix, "FN", func(name string) string {
			return "CREATE OR ALTER FUNCTION " + name + "() RETURNS INT AS BEGIN RETURN 1; END;"
		}},
		{"procedures", "dbo.rza_CaseProcedure_" + suffix, "P", func(name string) string {
			return "CREATE OR ALTER PROCEDURE " + name + " AS BEGIN SELECT 1 AS ExactCase; END;"
		}},
	}

	for _, tt := range tests {
		t.Run(tt.category, func(t *testing.T) {
			if _, err := tx.ExecContext(ctx, tt.create(tt.object)); err != nil {
				t.Fatalf("create %s: %v", tt.object, err)
			}
			var before int
			if err := tx.QueryRowContext(ctx, "SELECT COUNT(1) FROM sys.objects WHERE object_id = OBJECT_ID(@p1, @p2)", tt.object, tt.kind).Scan(&before); err != nil {
				t.Fatal(err)
			}
			if before != 1 {
				t.Fatalf("created object not found: %s", tt.object)
			}

			down, err := buildDropProgrammabilitySQL(tt.category, tt.object)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := tx.ExecContext(ctx, down); err != nil {
				t.Fatalf("generated down did not compile/apply for %s: %v\n%s", tt.object, err, down)
			}
			var after int
			if err := tx.QueryRowContext(ctx, "SELECT COUNT(1) FROM sys.objects WHERE object_id = OBJECT_ID(@p1, @p2)", tt.object, tt.kind).Scan(&after); err != nil {
				t.Fatal(err)
			}
			if after != 0 {
				t.Fatalf("generated down left object installed: %s", tt.object)
			}
		})
	}
}

func TestIntegration_SQLServerFunctionSubtypeDefinitions(t *testing.T) {
	dsn := os.Getenv("RZA_MSSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("RZA_MSSQL_TEST_DSN is not configured")
	}
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	suffix := fmt.Sprintf("%d", time.Now().UnixNano())
	tests := []struct {
		subtype string
		dbType  string
		sql     string
	}{
		{
			subtype: functionSubtypeScalar,
			dbType:  "FN",
			sql: "CREATE OR ALTER FUNCTION dbo.rza_Scalar_" + suffix +
				"(@value INT = 1) RETURNS INT WITH SCHEMABINDING AS BEGIN RETURN @value; END;",
		},
		{
			subtype: functionSubtypeInlineTable,
			dbType:  "IF",
			sql: "CREATE OR ALTER FUNCTION dbo.rza_Inline_" + suffix +
				"(@value INT) RETURNS TABLE WITH SCHEMABINDING AS RETURN (SELECT @value AS [Value]);",
		},
		{
			subtype: functionSubtypeMultiStatementTV,
			dbType:  "TF",
			sql: "CREATE OR ALTER FUNCTION dbo.rza_Multi_" + suffix +
				"(@value INT) RETURNS @result TABLE ([Id] INT NOT NULL PRIMARY KEY, [Value] INT NULL) WITH SCHEMABINDING AS BEGIN INSERT @result VALUES (1, @value); RETURN; END;",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.subtype, func(t *testing.T) {
			subtype, err := validateFunctionSource(test.sql)
			if err != nil {
				t.Fatalf("engine rejected valid %s function: %v", test.subtype, err)
			}
			if subtype != test.subtype {
				t.Fatalf("function subtype=%s, want %s", subtype, test.subtype)
			}
			if _, err := tx.ExecContext(ctx, test.sql); err != nil {
				t.Fatalf("SQL Server rejected %s function: %v\n%s", test.subtype, err, test.sql)
			}
			var actualType string
			if err := tx.QueryRowContext(ctx, `
SELECT o.type
FROM sys.objects o
WHERE o.object_id = OBJECT_ID(@p1);
`, functionObjectName(test.sql)).Scan(&actualType); err != nil {
				t.Fatal(err)
			}
			if actualType != test.dbType {
				t.Fatalf("SQL Server function type=%s, want %s", actualType, test.dbType)
			}
		})
	}
}

func functionObjectName(sql string) string {
	upper := strings.ToUpper(sql)
	start := strings.Index(upper, "FUNCTION ")
	if start < 0 {
		return ""
	}
	tail := strings.TrimSpace(sql[start+len("FUNCTION "):])
	if end := strings.Index(tail, "("); end >= 0 {
		return strings.TrimSpace(tail[:end])
	}
	return ""
}

func TestIntegration_SQLServerTableTypeLifecycleRecreatesDependentFunctions(t *testing.T) {
	dsn := os.Getenv("RZA_MSSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("RZA_MSSQL_TEST_DSN is not configured")
	}
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	suffix := fmt.Sprintf("%d", time.Now().UnixNano())
	typeName := "dbo.rza_ItemType_" + suffix
	functionName := "dbo.rza_CountItems_" + suffix
	procedureName := "dbo.rza_SaveItems_" + suffix
	oldType := "CREATE TYPE " + typeName + " AS TABLE ([Id] INT NOT NULL PRIMARY KEY);"
	newType := "CREATE TYPE " + typeName + " AS TABLE ([Id] INT NOT NULL PRIMARY KEY, [Qty] INT NULL);"
	oldFunction := "CREATE OR ALTER FUNCTION " + functionName + "(@items " + typeName + " READONLY) RETURNS INT AS BEGIN RETURN (SELECT COUNT(*) FROM @items); END;"
	oldProcedure := "CREATE OR ALTER PROCEDURE " + procedureName + " @items " + typeName + " READONLY AS SELECT " + functionName + "(@items) AS item_count;"

	for _, stmt := range []string{oldType, oldFunction, oldProcedure} {
		if _, err := tx.ExecContext(ctx, stmt); err != nil {
			t.Fatalf("prepare TVP lifecycle sandbox: %v\n%s", err, stmt)
		}
	}

	history := map[string]historyObject{
		objectKey("types", typeName): {
			Category: "types", Object: typeName, SemanticSQL: oldType,
		},
		objectKey("functions", functionName): {
			Category: "functions", Object: functionName, SemanticSQL: oldFunction,
		},
		objectKey("procedures", procedureName): {
			Category: "procedures", Object: procedureName, SemanticSQL: oldProcedure,
		},
	}
	desired := map[string]desiredObject{
		objectKey("types", typeName): {
			Category: "types", Object: typeName, SemanticSQL: newType,
		},
		objectKey("functions", functionName): {
			Category: "functions", Object: functionName, SemanticSQL: oldFunction,
		},
		objectKey("procedures", procedureName): {
			Category: "procedures", Object: procedureName, SemanticSQL: oldProcedure,
		},
	}
	up, _, dependents, err := buildTableTypeLifecycleSQL(
		typeName,
		oldType,
		newType,
		history,
		desired,
		[]string{"types", "tables", "indexes", "functions", "views", "procedures", "patch"},
	)
	if err != nil {
		t.Fatal(err)
	}
	if len(dependents) != 2 {
		t.Fatalf("dependent module count=%d, want 2: %v", len(dependents), dependents)
	}
	wrapped := renderTxWrapper(
		"20991231235959002",
		"types",
		"20991231235959002_types_"+typeName+"_v2.up.sql",
		typeName,
		true,
		up,
	)
	if _, err := tx.ExecContext(ctx, wrapped); err != nil {
		t.Fatalf("generated TVP lifecycle did not compile/apply: %v\n%s", err, wrapped)
	}

	var qtyColumns int
	if err := tx.QueryRowContext(ctx, `
SELECT COUNT(1)
FROM sys.table_types tt
JOIN sys.schemas s ON s.schema_id = tt.schema_id
JOIN sys.columns c ON c.object_id = tt.type_table_object_id
WHERE s.name = N'dbo' AND tt.name = @p1 AND c.name = N'Qty';
`, strings.TrimPrefix(typeName, "dbo.")).Scan(&qtyColumns); err != nil {
		t.Fatal(err)
	}
	if qtyColumns != 1 {
		t.Fatalf("recreated TVP Qty column count=%d, want 1", qtyColumns)
	}
	for _, object := range []string{functionName, procedureName} {
		var exists int
		if err := tx.QueryRowContext(ctx, "SELECT COUNT(1) FROM sys.objects WHERE object_id = OBJECT_ID(@p1)", object).Scan(&exists); err != nil {
			t.Fatal(err)
		}
		if exists != 1 {
			t.Fatalf("dependent module was not recreated: %s", object)
		}
	}
}

func TestIntegration_SQLServerHardenedTableAndIndexPrechecks(t *testing.T) {
	dsn := os.Getenv("RZA_MSSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("RZA_MSSQL_TEST_DSN is not configured")
	}
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	t.Run("XML to NVARCHAR rollback rejects invalid XML", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			t.Fatal(err)
		}
		defer tx.Rollback()
		name := fmt.Sprintf("rza_sandbox_xml_text_%d", time.Now().UnixNano())
		if _, err := tx.ExecContext(ctx, "CREATE TABLE dbo."+name+" ([payload] XML NULL); INSERT INTO dbo."+name+" VALUES (N'<root/>');"); err != nil {
			t.Fatal(err)
		}
		xmlDef := tableColumnDef{Name: "payload", BaseType: "XML", RawType: "XML", Nullable: true}
		textDef := tableColumnDef{Name: "payload", BaseType: "NVARCHAR", RawType: "NVARCHAR(MAX)", Length: -1, IsUnicode: true, Nullable: true}
		up := renderTxWrapper("20991231235959990", "tables", "20991231235959990_tables_dbo."+name+"_v2.up.sql", "dbo."+name, true, buildAlterColumnGuardedSQL("dbo."+name, xmlDef, textDef))
		if _, err := tx.ExecContext(ctx, up); err != nil {
			t.Fatalf("XML to NVARCHAR(MAX) did not compile/apply: %v", err)
		}
		// Plain text is valid XML content in SQL Server; use an actually
		// malformed fragment so TRY_CONVERT(XML, ...) must return NULL.
		if _, err := tx.ExecContext(ctx, "INSERT INTO dbo."+name+" ([payload]) VALUES (N'<not-xml');"); err != nil {
			t.Fatal(err)
		}
		down := renderTxWrapper("20991231235959990", "tables", "20991231235959990_tables_dbo."+name+"_v2.down.sql", "dbo."+name, false, buildAlterColumnGuardedSQL("dbo."+name, textDef, xmlDef))
		if _, err := tx.ExecContext(ctx, down); err == nil {
			t.Fatal("NVARCHAR(MAX) to XML rollback must reject invalid XML text")
		}
	})

	t.Run("UNIQUE INDEX rejects duplicate keys", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			t.Fatal(err)
		}
		defer tx.Rollback()
		name := fmt.Sprintf("rza_sandbox_unique_index_%d", time.Now().UnixNano())
		if _, err := tx.ExecContext(ctx, "CREATE TABLE dbo."+name+" ([code] INT NULL); INSERT INTO dbo."+name+" VALUES (1), (1);"); err != nil {
			t.Fatal(err)
		}
		indexName := "UX_" + name + "_code"
		body := buildCreateIndexIfNotExistsStatement("dbo."+name, indexName, "CREATE UNIQUE INDEX ["+indexName+"] ON dbo."+name+" ([code]);")
		wrapped := renderTxWrapper("20991231235959991", "indexes", "20991231235959991_indexes_dbo."+name+"_v1.up.sql", "dbo."+name, true, body)
		if _, err := tx.ExecContext(ctx, wrapped); err == nil {
			t.Fatal("UNIQUE INDEX precheck must reject duplicate keys")
		}
	})

	t.Run("destructive down rejects populated table", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			t.Fatal(err)
		}
		defer tx.Rollback()
		name := fmt.Sprintf("rza_sandbox_drop_table_%d", time.Now().UnixNano())
		if _, err := tx.ExecContext(ctx, "CREATE TABLE dbo."+name+" ([id] INT NOT NULL); INSERT INTO dbo."+name+" VALUES (1);"); err != nil {
			t.Fatal(err)
		}
		wrapped := renderTxWrapper("20991231235959992", "tables", "20991231235959992_tables_dbo."+name+"_v1.down.sql", "dbo."+name, false, buildDropTableSQL("dbo."+name))
		if _, err := tx.ExecContext(ctx, wrapped); err == nil {
			t.Fatal("destructive down must reject a populated table")
		}
	})
}

func TestIntegration_SQLServerGeneratedUniqueConstraintGuard(t *testing.T) {
	dsn := os.Getenv("RZA_MSSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("RZA_MSSQL_TEST_DSN is not configured")
	}
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()
	name := fmt.Sprintf("rza_sandbox_uq_%d", time.Now().UnixNano())
	if _, err := tx.ExecContext(ctx, "CREATE TABLE dbo."+name+" ([code] INT NULL);"); err != nil {
		t.Fatal(err)
	}
	constraint := "UQ_" + name + "_code"
	body, err := buildAddKeyConstraintWithPrecheckStatement("dbo."+name, constraint, "unique", "CONSTRAINT ["+constraint+"] UNIQUE ([code])")
	if err != nil {
		t.Fatal(err)
	}
	wrapped := renderTxWrapper("20991231235959001", "tables", "20991231235959001_tables_dbo."+name+"_v1.up.sql", "dbo."+name, true, body)
	if _, err := tx.ExecContext(ctx, wrapped); err != nil {
		t.Fatalf("unique guard did not compile/apply: %v", err)
	}
	var count int
	if err := tx.QueryRowContext(ctx, "SELECT COUNT(1) FROM sys.key_constraints WHERE parent_object_id = OBJECT_ID(N'dbo."+name+"') AND name = @p1", constraint).Scan(&count); err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatalf("unique constraint introspection mismatch: %d", count)
	}
}

func TestIntegration_SQLServerGeneratedUniqueConstraintGuardRejectsDuplicates(t *testing.T) {
	dsn := os.Getenv("RZA_MSSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("RZA_MSSQL_TEST_DSN is not configured")
	}
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()
	name := fmt.Sprintf("rza_sandbox_uq_dup_%d", time.Now().UnixNano())
	if _, err := tx.ExecContext(ctx, "CREATE TABLE dbo."+name+" ([code] INT NULL); INSERT INTO dbo."+name+" ([code]) VALUES (1), (1);"); err != nil {
		t.Fatal(err)
	}
	constraint := "UQ_" + name + "_code"
	body, err := buildAddKeyConstraintWithPrecheckStatement("dbo."+name, constraint, "unique", "CONSTRAINT ["+constraint+"] UNIQUE ([code])")
	if err != nil {
		t.Fatal(err)
	}
	wrapped := renderTxWrapper("20991231235959002", "tables", "20991231235959002_tables_dbo."+name+"_v1.up.sql", "dbo."+name, true, body)
	if _, err := tx.ExecContext(ctx, wrapped); err == nil {
		t.Fatal("expected duplicate-key precheck to reject generated unique constraint")
	}
}

func TestIntegration_SQLServerGeneratedComputedColumnAndPrimaryKeyGuard(t *testing.T) {
	dsn := os.Getenv("RZA_MSSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("RZA_MSSQL_TEST_DSN is not configured")
	}
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()
	name := fmt.Sprintf("rza_sandbox_pk_comp_%d", time.Now().UnixNano())
	if _, err := tx.ExecContext(ctx, "CREATE TABLE dbo."+name+" ([id] INT NOT NULL, [base] INT NOT NULL);"); err != nil {
		t.Fatal(err)
	}
	computed := tableColumnDef{Name: "doubled", IsComputed: true, ComputedDefinition: "AS ([base] * (2))"}
	computedBody := buildAddColumnIfMissingStatement("dbo."+name, computed)
	computedWrapped := renderTxWrapper("20991231235959003", "tables", "20991231235959003_tables_dbo."+name+"_v1.up.sql", "dbo."+name, true, computedBody)
	if _, err := tx.ExecContext(ctx, computedWrapped); err != nil {
		t.Fatalf("computed-column SQL did not compile/apply: %v", err)
	}
	var computedCount int
	if err := tx.QueryRowContext(ctx, "SELECT COUNT(1) FROM sys.columns WHERE object_id = OBJECT_ID(N'dbo."+name+"') AND name = N'doubled' AND is_computed = 1").Scan(&computedCount); err != nil {
		t.Fatal(err)
	}
	if computedCount != 1 {
		t.Fatalf("computed column introspection mismatch: %d", computedCount)
	}
	constraint := "PK_" + name
	pkBody, err := buildAddKeyConstraintWithPrecheckStatement("dbo."+name, constraint, "primary_key", "CONSTRAINT ["+constraint+"] PRIMARY KEY ([id])")
	if err != nil {
		t.Fatal(err)
	}
	pkWrapped := renderTxWrapper("20991231235959004", "tables", "20991231235959004_tables_dbo."+name+"_v2.up.sql", "dbo."+name, true, pkBody)
	if _, err := tx.ExecContext(ctx, pkWrapped); err != nil {
		t.Fatalf("primary-key guard did not compile/apply: %v", err)
	}
	var pkCount int
	if err := tx.QueryRowContext(ctx, "SELECT COUNT(1) FROM sys.key_constraints WHERE parent_object_id = OBJECT_ID(N'dbo."+name+"') AND name = @p1 AND type = 'PK'", constraint).Scan(&pkCount); err != nil {
		t.Fatal(err)
	}
	if pkCount != 1 {
		t.Fatalf("primary key introspection mismatch: %d", pkCount)
	}
}

func TestIntegration_SQLServerGeneratedForeignKeyGuard(t *testing.T) {
	dsn := os.Getenv("RZA_MSSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("RZA_MSSQL_TEST_DSN is not configured")
	}
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()
	parent := fmt.Sprintf("rza_sandbox_parent_%d", time.Now().UnixNano())
	child := fmt.Sprintf("rza_sandbox_child_%d", time.Now().UnixNano())
	if _, err := tx.ExecContext(ctx, "CREATE TABLE dbo."+parent+" ([id] INT NOT NULL CONSTRAINT PK_"+parent+" PRIMARY KEY); CREATE TABLE dbo."+child+" ([parent_id] INT NULL); INSERT INTO dbo."+parent+" ([id]) VALUES (1); INSERT INTO dbo."+child+" ([parent_id]) VALUES (1), (NULL);"); err != nil {
		t.Fatal(err)
	}
	name := "FK_" + child + "_parent"
	definition := "CONSTRAINT [" + name + "] FOREIGN KEY ([parent_id]) REFERENCES dbo." + parent + " ([id]) ON DELETE CASCADE"
	body, err := buildAddForeignKeyWithPrecheckStatement("dbo."+child, name, definition)
	if err != nil {
		t.Fatal(err)
	}
	wrapped := renderTxWrapper("20991231235959005", "tables", "20991231235959005_tables_dbo."+child+"_v1.up.sql", "dbo."+child, true, body)
	if _, err := tx.ExecContext(ctx, wrapped); err != nil {
		t.Fatalf("foreign-key guard did not compile/apply: %v", err)
	}
	var count int
	if err := tx.QueryRowContext(ctx, "SELECT COUNT(1) FROM sys.foreign_keys WHERE parent_object_id = OBJECT_ID(N'dbo."+child+"') AND name = @p1", name).Scan(&count); err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatalf("foreign key introspection mismatch: %d", count)
	}
}

func TestIntegration_SQLServerCommentedTablesExecuteInForeignKeyOrder(t *testing.T) {
	dsn := os.Getenv("RZA_MSSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("RZA_MSSQL_TEST_DSN is not configured")
	}
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	suffix := fmt.Sprintf("%d", time.Now().UnixNano())
	parent := "rza_comment_parent_" + suffix
	child := "rza_comment_child_" + suffix
	parentObject := "dbo." + parent
	childObject := "dbo." + child
	fkName := "FK_" + child + "_" + parent

	parentSource := fmt.Sprintf(`CREATE TABLE %s (
  -- primary event identifier
  EventId INT NOT NULL,
  CONSTRAINT [PK_%s] PRIMARY KEY ([EventId])
);`, parentObject, parent)
	childSource := fmt.Sprintf(`CREATE TABLE %s (
  -- foreign-key event identifier
  EventId INT NOT NULL,
  -- readable event name
  Nama VARCHAR(100),
  CONSTRAINT [PK_%s] PRIMARY KEY ([EventId]),
  CONSTRAINT [%s] FOREIGN KEY(EventId) REFERENCES %s([EventId])
);`, childObject, child, fkName, parentObject)

	parentSemantic, err := semanticFromSource(migrationMeta{Category: "tables", Object: parentObject}, parentSource)
	if err != nil {
		t.Fatalf("canonicalize parent source: %v", err)
	}
	childSemantic, err := semanticFromSource(migrationMeta{Category: "tables", Object: childObject}, childSource)
	if err != nil {
		t.Fatalf("canonicalize child source: %v", err)
	}

	ops := []plannedMigration{
		{Category: "tables", Object: childObject, UpSemanticSQL: childSemantic},
		{Category: "tables", Object: parentObject, UpSemanticSQL: parentSemantic},
	}
	desired := map[string]desiredObject{
		objectKey("tables", parentObject): {Category: "tables", Object: parentObject, SemanticSQL: parentSemantic},
		objectKey("tables", childObject):  {Category: "tables", Object: childObject, SemanticSQL: childSemantic},
	}
	ordered, err := orderPlannedMigrations(ops, desired, []string{"tables"})
	if err != nil {
		t.Fatalf("order commented FK tables: %v", err)
	}
	if len(ordered.Ops) != 2 || ordered.Ops[0].Object != parentObject || ordered.Ops[1].Object != childObject {
		t.Fatalf("unexpected live execution order: %+v", ordered.Ops)
	}

	for _, op := range ordered.Ops {
		if _, err := tx.ExecContext(ctx, op.UpSemanticSQL); err != nil {
			t.Fatalf("execute generated CREATE TABLE for %s: %v\n%s", op.Object, err, op.UpSemanticSQL)
		}
	}

	var parentColumns, childColumns, foreignKeys int
	if err := tx.QueryRowContext(ctx, "SELECT COUNT(1) FROM sys.columns WHERE object_id = OBJECT_ID(@p1)", parentObject).Scan(&parentColumns); err != nil {
		t.Fatal(err)
	}
	if err := tx.QueryRowContext(ctx, "SELECT COUNT(1) FROM sys.columns WHERE object_id = OBJECT_ID(@p1)", childObject).Scan(&childColumns); err != nil {
		t.Fatal(err)
	}
	if err := tx.QueryRowContext(ctx, "SELECT COUNT(1) FROM sys.foreign_keys WHERE parent_object_id = OBJECT_ID(@p1) AND referenced_object_id = OBJECT_ID(@p2) AND name = @p3", childObject, parentObject, fkName).Scan(&foreignKeys); err != nil {
		t.Fatal(err)
	}
	if parentColumns != 1 || childColumns != 2 || foreignKeys != 1 {
		t.Fatalf("live introspection mismatch: parent_columns=%d child_columns=%d foreign_keys=%d", parentColumns, childColumns, foreignKeys)
	}
}

func TestIntegration_SQLServerGeneratedKeyTransitionRestoresIncomingForeignKey(t *testing.T) {
	dsn := os.Getenv("RZA_MSSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("RZA_MSSQL_TEST_DSN is not configured")
	}
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()
	parent := fmt.Sprintf("rza_sandbox_key_parent_%d", time.Now().UnixNano())
	child := fmt.Sprintf("rza_sandbox_key_child_%d", time.Now().UnixNano())
	uq := "UQ_" + parent + "_id"
	fk := "FK_" + child + "_parent"
	if _, err := tx.ExecContext(ctx, "CREATE TABLE dbo."+parent+" ([id] INT NOT NULL, CONSTRAINT ["+uq+"] UNIQUE ([id])); CREATE TABLE dbo."+child+" ([parent_id] INT NULL, CONSTRAINT ["+fk+"] FOREIGN KEY ([parent_id]) REFERENCES dbo."+parent+" ([id])); INSERT INTO dbo."+parent+" VALUES (1); INSERT INTO dbo."+child+" VALUES (1);"); err != nil {
		t.Fatal(err)
	}
	previous := "CREATE TABLE dbo." + parent + " ([id] INT NOT NULL, CONSTRAINT [" + uq + "] UNIQUE ([id]));"
	next := "CREATE TABLE dbo." + parent + " ([id] INT NOT NULL, CONSTRAINT [PK_" + parent + "] PRIMARY KEY ([id]));"
	history := map[string]historyObject{objectKey("tables", "dbo."+child): {Category: "tables", Object: "dbo." + child, SemanticSQL: "CREATE TABLE dbo." + child + " ([parent_id] INT NULL, CONSTRAINT [" + fk + "] FOREIGN KEY ([parent_id]) REFERENCES dbo." + parent + " ([id]));"}}
	body, _, ok, err := buildConstraintLifecycleWithDependencies(previous, next, "dbo."+parent, history)
	if err != nil || !ok {
		t.Fatalf("build key transition: ok=%v err=%v", ok, err)
	}
	wrapped := renderTxWrapper("20991231235959006", "tables", "20991231235959006_tables_dbo."+parent+"_v2.up.sql", "dbo."+parent, true, body)
	if _, err := tx.ExecContext(ctx, wrapped); err != nil {
		t.Fatalf("key transition did not compile/apply: %v", err)
	}
	var fkCount, pkCount int
	if err := tx.QueryRowContext(ctx, "SELECT COUNT(1) FROM sys.foreign_keys WHERE parent_object_id = OBJECT_ID(N'dbo."+child+"') AND name = @p1", fk).Scan(&fkCount); err != nil {
		t.Fatal(err)
	}
	if err := tx.QueryRowContext(ctx, "SELECT COUNT(1) FROM sys.key_constraints WHERE parent_object_id = OBJECT_ID(N'dbo."+parent+"') AND type = 'PK'").Scan(&pkCount); err != nil {
		t.Fatal(err)
	}
	if fkCount != 1 || pkCount != 1 {
		t.Fatalf("key/FK restoration mismatch: fk=%d pk=%d", fkCount, pkCount)
	}
}

func TestIntegration_SQLServerGeneratedIdentityAndRowGuidAdd(t *testing.T) {
	dsn := os.Getenv("RZA_MSSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("RZA_MSSQL_TEST_DSN is not configured")
	}
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()
	name := fmt.Sprintf("rza_sandbox_modifiers_%d", time.Now().UnixNano())
	if _, err := tx.ExecContext(ctx, "CREATE TABLE dbo."+name+" ([base] INT NOT NULL);"); err != nil {
		t.Fatal(err)
	}
	identity := tableColumnDef{Name: "seq", BaseType: "BIGINT", RawType: "BIGINT", Nullable: false, IdentityDefinition: "IDENTITY(10, 2)"}
	guid := tableColumnDef{Name: "row_guid", BaseType: "UNIQUEIDENTIFIER", RawType: "UNIQUEIDENTIFIER", Nullable: true, RowGuidCol: true}
	body := concatSQLSegments(buildAddColumnIfMissingStatement("dbo."+name, identity), buildAddColumnIfMissingStatement("dbo."+name, guid))
	wrapped := renderTxWrapper("20991231235959007", "tables", "20991231235959007_tables_dbo."+name+"_v1.up.sql", "dbo."+name, true, body)
	if _, err := tx.ExecContext(ctx, wrapped); err != nil {
		t.Fatalf("modifier add did not compile/apply: %v", err)
	}
	var identityCount, guidCount int
	if err := tx.QueryRowContext(ctx, "SELECT COUNT(1) FROM sys.identity_columns WHERE object_id = OBJECT_ID(N'dbo."+name+"') AND name = N'seq'").Scan(&identityCount); err != nil {
		t.Fatal(err)
	}
	if err := tx.QueryRowContext(ctx, "SELECT COUNT(1) FROM sys.columns WHERE object_id = OBJECT_ID(N'dbo."+name+"') AND name = N'row_guid' AND is_rowguidcol = 1").Scan(&guidCount); err != nil {
		t.Fatal(err)
	}
	if identityCount != 1 || guidCount != 1 {
		t.Fatalf("modifier introspection mismatch: identity=%d rowguid=%d", identityCount, guidCount)
	}
}

func TestIntegration_SQLServerGeneratedSparseAddAndReverse(t *testing.T) {
	dsn := os.Getenv("RZA_MSSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("RZA_MSSQL_TEST_DSN is not configured")
	}
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()
	name := fmt.Sprintf("rza_sandbox_sparse_%d", time.Now().UnixNano())
	if _, err := tx.ExecContext(ctx, "CREATE TABLE dbo."+name+" ([id] INT NOT NULL);"); err != nil {
		t.Fatal(err)
	}
	def := tableColumnDef{Name: "optional_code", BaseType: "VARCHAR", RawType: "VARCHAR(50)", Nullable: true, Sparse: true}
	up := renderTxWrapper("20991231235959008", "tables", "20991231235959008_tables_dbo."+name+"_v1.up.sql", "dbo."+name, true, buildAddColumnIfMissingStatement("dbo."+name, def))
	if _, err := tx.ExecContext(ctx, up); err != nil {
		t.Fatalf("SPARSE up SQL did not compile: %v", err)
	}
	var count int
	if err := tx.QueryRowContext(ctx, "SELECT COUNT(1) FROM sys.columns WHERE object_id=OBJECT_ID(N'dbo."+name+"') AND name=N'optional_code' AND is_sparse=1").Scan(&count); err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatalf("SPARSE introspection mismatch: %d", count)
	}
	down := renderTxWrapper("20991231235959008", "tables", "20991231235959008_tables_dbo."+name+"_v1.down.sql", "dbo."+name, false, buildDropColumnIfExistsStatement("dbo."+name, def.Name))
	if _, err := tx.ExecContext(ctx, down); err != nil {
		t.Fatalf("SPARSE down SQL did not compile: %v", err)
	}
}

func TestIntegration_SQLServerAdvancedIndexFamiliesAndCompression(t *testing.T) {
	dsn := os.Getenv("RZA_MSSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("RZA_MSSQL_TEST_DSN is not configured")
	}
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()
	suffix := fmt.Sprintf("%d", time.Now().UnixNano())
	execGenerated := func(version, table, index, statement string) {
		t.Helper()
		body := buildCreateIndexIfNotExistsStatement("dbo."+table, index, statement)
		wrapped := renderTxWrapper(version, "indexes", version+"_indexes_dbo."+table+"_v1.up.sql", "dbo."+table, true, body)
		if _, err := tx.ExecContext(ctx, wrapped); err != nil {
			t.Fatalf("advanced generated index SQL failed for %s: %v", index, err)
		}
	}

	xmlTable := "rza_xml_" + suffix
	if _, err := tx.ExecContext(ctx, "CREATE TABLE dbo."+xmlTable+" ([id] INT NOT NULL CONSTRAINT [PK_"+xmlTable+"] PRIMARY KEY, [doc] XML NULL);"); err != nil {
		t.Fatal(err)
	}
	primaryXML := "PXML_" + xmlTable
	secondaryXML := "SXML_" + xmlTable
	execGenerated("20991231235959020", xmlTable, primaryXML, "CREATE PRIMARY XML INDEX ["+primaryXML+"] ON dbo."+xmlTable+" ([doc]);")
	execGenerated("20991231235959021", xmlTable, secondaryXML, "CREATE XML INDEX ["+secondaryXML+"] ON dbo."+xmlTable+" ([doc]) USING XML INDEX ["+primaryXML+"] FOR PATH;")

	spatialTable := "rza_spatial_" + suffix
	if _, err := tx.ExecContext(ctx, "CREATE TABLE dbo."+spatialTable+" ([id] INT NOT NULL CONSTRAINT [PK_"+spatialTable+"] PRIMARY KEY CLUSTERED, [shape] geometry NULL);"); err != nil {
		t.Fatal(err)
	}
	spatial := "SIX_" + spatialTable
	execGenerated("20991231235959022", spatialTable, spatial, "CREATE SPATIAL INDEX ["+spatial+"] ON dbo."+spatialTable+" ([shape]) USING GEOMETRY_AUTO_GRID WITH (BOUNDING_BOX = (-180, -90, 180, 90));")

	columnstoreTable := "rza_columnstore_" + suffix
	if _, err := tx.ExecContext(ctx, "CREATE TABLE dbo."+columnstoreTable+" ([id] BIGINT NOT NULL, [amount] DECIMAL(18,2) NULL);"); err != nil {
		t.Fatal(err)
	}
	columnstore := "NCCI_" + columnstoreTable
	execGenerated("20991231235959023", columnstoreTable, columnstore, "CREATE NONCLUSTERED COLUMNSTORE INDEX ["+columnstore+"] ON dbo."+columnstoreTable+" ([id], [amount]);")

	compressionTable := "rza_compression_" + suffix
	if _, err := tx.ExecContext(ctx, "CREATE TABLE dbo."+compressionTable+" ([id] BIGINT NOT NULL, [code] VARCHAR(50) NULL);"); err != nil {
		t.Fatal(err)
	}
	compressed := "IX_" + compressionTable
	execGenerated("20991231235959024", compressionTable, compressed, "CREATE NONCLUSTERED INDEX ["+compressed+"] ON dbo."+compressionTable+" ([id]) INCLUDE ([code]) WITH (DATA_COMPRESSION = PAGE);")

	checks := []struct {
		table string
		name  string
		typeD string
	}{{xmlTable, primaryXML, "XML"}, {xmlTable, secondaryXML, "XML"}, {spatialTable, spatial, "SPATIAL"}, {columnstoreTable, columnstore, "COLUMNSTORE"}, {compressionTable, compressed, "NONCLUSTERED"}}
	for _, check := range checks {
		var count int
		if err := tx.QueryRowContext(ctx, "SELECT COUNT(1) FROM sys.indexes WHERE object_id=OBJECT_ID(@p1) AND name=@p2 AND type_desc LIKE '%'+@p3+'%'", "dbo."+check.table, check.name, check.typeD).Scan(&count); err != nil {
			t.Fatal(err)
		}
		if count != 1 {
			t.Fatalf("advanced index introspection mismatch: %s.%s type=%s", check.table, check.name, check.typeD)
		}
	}
	var compressedPartitions int
	if err := tx.QueryRowContext(ctx, "SELECT COUNT(1) FROM sys.partitions p JOIN sys.indexes i ON i.object_id=p.object_id AND i.index_id=p.index_id WHERE i.object_id=OBJECT_ID(@p1) AND i.name=@p2 AND p.data_compression_desc='PAGE'", "dbo."+compressionTable, compressed).Scan(&compressedPartitions); err != nil {
		t.Fatal(err)
	}
	if compressedPartitions == 0 {
		t.Fatal("PAGE compression was not applied")
	}
}

func TestIntegration_SQLServerTypedColumnRenameUpAndDown(t *testing.T) {
	dsn := os.Getenv("RZA_MSSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("RZA_MSSQL_TEST_DSN is not configured")
	}
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()
	name := fmt.Sprintf("rza_rename_%d", time.Now().UnixNano())
	setup := "CREATE TABLE dbo." + name + " ([id] INT NOT NULL CONSTRAINT [PK_" + name + "] PRIMARY KEY, [old_name] VARCHAR(40) NULL CONSTRAINT [DF_" + name + "_name] DEFAULT ('x'), CONSTRAINT [CK_" + name + "_name] CHECK ([old_name] <> 'bad')); CREATE INDEX [IX_" + name + "_name] ON dbo." + name + " ([old_name]); INSERT INTO dbo." + name + " ([id],[old_name]) VALUES (1,'kept');"
	if _, err := tx.ExecContext(ctx, setup); err != nil {
		t.Fatal(err)
	}
	previous := "CREATE TABLE dbo." + name + " ([id] INT NOT NULL CONSTRAINT [PK_" + name + "] PRIMARY KEY, [old_name] VARCHAR(40) NULL CONSTRAINT [DF_" + name + "_name] DEFAULT ('x'), CONSTRAINT [CK_" + name + "_name] CHECK ([old_name] <> 'bad'));"
	next := "CREATE TABLE dbo." + name + " ([id] INT NOT NULL CONSTRAINT [PK_" + name + "] PRIMARY KEY, [new_name] VARCHAR(40) NULL CONSTRAINT [DF_" + name + "_name] DEFAULT ('x'), CONSTRAINT [CK_" + name + "_name] CHECK ([new_name] <> 'bad'));"
	history := map[string]historyObject{
		objectKey("tables", "dbo."+name):  {Category: "tables", Object: "dbo." + name, SemanticSQL: previous},
		objectKey("indexes", "dbo."+name): {Category: "indexes", Object: "dbo." + name, SemanticSQL: "CREATE INDEX [IX_" + name + "_name] ON dbo." + name + " ([old_name]);"},
	}
	up, down, ok, err := planColumnRenames(previous, next, "dbo."+name, map[string]string{"old_name": "new_name"}, history)
	if err != nil || !ok {
		t.Fatalf("rename plan: ok=%v err=%v", ok, err)
	}
	wrappedUp := renderTxWrapper("20991231235959030", "tables", "20991231235959030_tables_dbo."+name+"_v2.up.sql", "dbo."+name, true, up)
	if _, err := tx.ExecContext(ctx, wrappedUp); err != nil {
		t.Fatalf("rename up failed: %v", err)
	}
	var value string
	if err := tx.QueryRowContext(ctx, "SELECT [new_name] FROM dbo."+name+" WHERE [id]=1").Scan(&value); err != nil || value != "kept" {
		t.Fatalf("rename lost data: value=%q err=%v", value, err)
	}
	var dependencyCount int
	if err := tx.QueryRowContext(ctx, "SELECT COUNT(1) FROM sys.index_columns ic JOIN sys.columns c ON c.object_id=ic.object_id AND c.column_id=ic.column_id JOIN sys.indexes i ON i.object_id=ic.object_id AND i.index_id=ic.index_id WHERE i.object_id=OBJECT_ID(@p1) AND i.name=@p2 AND c.name='new_name'", "dbo."+name, "IX_"+name+"_name").Scan(&dependencyCount); err != nil || dependencyCount != 1 {
		t.Fatalf("index dependency did not follow rename: count=%d err=%v", dependencyCount, err)
	}
	wrappedDown := renderTxWrapper("20991231235959030", "tables", "20991231235959030_tables_dbo."+name+"_v2.down.sql", "dbo."+name, false, down)
	if _, err := tx.ExecContext(ctx, wrappedDown); err != nil {
		t.Fatalf("rename down failed: %v", err)
	}
	if err := tx.QueryRowContext(ctx, "SELECT [old_name] FROM dbo."+name+" WHERE [id]=1").Scan(&value); err != nil || value != "kept" {
		t.Fatalf("rename down lost data: value=%q err=%v", value, err)
	}
}
