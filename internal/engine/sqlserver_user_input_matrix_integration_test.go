package engine

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"

	_ "github.com/microsoft/go-mssqldb"
)

// This matrix exercises user-authored db_object SQL through source parsing,
// migration rendering, SQL Server compilation, execution, and catalog lookup.
// Every object is created inside an outer transaction and rolled back.
func TestIntegration_SQLServerUserAuthoredSQLMatrix(t *testing.T) {
	dsn := os.Getenv("RZA_MSSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("RZA_MSSQL_TEST_DSN is not configured")
	}
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	suffix := fmt.Sprintf("%d", time.Now().UnixNano())
	table := "dbo.it_wild_table_" + suffix
	scalarFn := "dbo.it_wild_scalar_" + suffix
	inlineFn := "dbo.it_wild_inline_" + suffix
	multiFn := "dbo.it_wild_multi_" + suffix
	view := "dbo.it_wild_view_" + suffix
	procedure := "dbo.it_wild_proc_" + suffix

	tableSQL := fmt.Sprintf(`CREATE TABLE %s (
  [Id] BIGINT IDENTITY(10,5) NOT NULL,
  [Select] NVARCHAR(40) COLLATE Latin1_General_BIN NOT NULL,
  [Owner's Note] NVARCHAR(100) COLLATE Latin1_General_BIN NULL
    CONSTRAINT [DF_%s_note] DEFAULT (N'O''Brien'),
  [Amount] DECIMAL NULL,
  [OccurredAt] DATETIME2 NULL,
  [Payload] VARBINARY NULL,
  [IsActive] BIT NOT NULL CONSTRAINT [DF_%s_active] DEFAULT ((1)),
  [ComputedAmount] AS (ISNULL([Amount],(0))*(2)) PERSISTED,
  CONSTRAINT [PK_%s] PRIMARY KEY CLUSTERED ([Id] ASC),
  CONSTRAINT [UQ_%s_select] UNIQUE NONCLUSTERED ([Select] ASC),
  CONSTRAINT [CK_%s_amount] CHECK ([Amount] IS NULL OR [Amount] >= (0))
);`, table, suffix, suffix, suffix, suffix, suffix)
	applyAuthoredSource(t, ctx, tx, "tables", table, tableSQL, "U")

	indexSQL := fmt.Sprintf(`CREATE NONCLUSTERED INDEX [IX_%s_filter]
ON %s ([Select] ASC, [Id] DESC)
INCLUDE ([Amount], [OccurredAt])
WHERE [IsActive] = (1)
WITH (PAD_INDEX = ON, FILLFACTOR = 90, SORT_IN_TEMPDB = OFF);`, suffix, table)
	applyAuthoredSource(t, ctx, tx, "indexes", table, indexSQL, "")
	var indexCount int
	if err := tx.QueryRowContext(ctx, `
SELECT COUNT(*)
FROM sys.indexes
WHERE object_id=OBJECT_ID(@p1, N'U') AND name=@p2 AND has_filter=1;`,
		table, "IX_"+suffix+"_filter").Scan(&indexCount); err != nil || indexCount != 1 {
		t.Fatalf("filtered/include/descending index catalog mismatch: count=%d err=%v", indexCount, err)
	}

	scalarSQL := fmt.Sprintf(`CREATE OR ALTER FUNCTION %s
(
  @input NVARCHAR(100)
)
RETURNS NVARCHAR(200)
AS
BEGIN
  -- GO inside a comment is not a batch separator.
  RETURN CONCAT(N'GO:', N' O''Brien ', @input);
END;`, scalarFn)
	applyAuthoredSource(t, ctx, tx, "functions", scalarFn, scalarSQL, "FN")

	inlineSQL := fmt.Sprintf(`CREATE OR ALTER FUNCTION %s(@minimum DECIMAL(18,2))
RETURNS TABLE
AS
RETURN
(
  SELECT [Id], [Select], [Amount]
  FROM %s
  WHERE [Amount] >= @minimum
);`, inlineFn, table)
	applyAuthoredSource(t, ctx, tx, "functions", inlineFn, inlineSQL, "IF")

	multiSQL := fmt.Sprintf(`CREATE OR ALTER FUNCTION %s(@value INT)
RETURNS @result TABLE ([Value] INT NOT NULL, [Label] NVARCHAR(20) NOT NULL)
AS
BEGIN
  INSERT @result ([Value], [Label])
  VALUES (@value, CASE WHEN @value < 0 THEN N'negative' ELSE N'positive' END);
  RETURN;
END;`, multiFn)
	applyAuthoredSource(t, ctx, tx, "functions", multiFn, multiSQL, "TF")

	viewSQL := fmt.Sprintf(`CREATE OR ALTER VIEW %s
AS
  SELECT
    [Id],
    [Select],
    [Owner's Note],
    CAST(N'GO is data; -- is data' AS NVARCHAR(40)) AS [Literal]
  FROM %s
  WHERE [IsActive] = (1);`, view, table)
	applyAuthoredSource(t, ctx, tx, "views", view, viewSQL, "V")

	procedureSQL := fmt.Sprintf(`CREATE OR ALTER PROCEDURE %s
  @Name NVARCHAR(40),
  @ResultCount INT OUTPUT
AS
BEGIN
  SET NOCOUNT ON;
  CREATE TABLE #Matched ([Id] BIGINT NOT NULL, [Text] NVARCHAR(100) NULL);

  DECLARE @sql NVARCHAR(MAX) = N'
INSERT #Matched ([Id], [Text])
SELECT [Id], [Owner''s Note]
FROM %s
WHERE [Select] = @value
  AND N''GO'' = N''GO'';';

  BEGIN TRY
    EXEC sys.sp_executesql @sql, N'@value NVARCHAR(40)', @value=@Name;
    SELECT @ResultCount = COUNT(*) FROM #Matched;
  END TRY
  BEGIN CATCH
    THROW;
  END CATCH;
END;`, procedure, table)
	applyAuthoredSource(t, ctx, tx, "procedures", procedure, procedureSQL, "P")

	var resultCount int
	if _, err := tx.ExecContext(ctx,
		fmt.Sprintf("EXEC %s @Name=@p1, @ResultCount=@p2 OUTPUT;", procedure),
		sql.Named("p1", "nothing"),
		sql.Named("p2", sql.Out{Dest: &resultCount}),
	); err != nil {
		t.Fatalf("multiline dynamic SQL procedure failed: %v", err)
	}
}

func TestIntegration_SQLServerBootstrapCompatibleTableSkipAndCoreDrift(t *testing.T) {
	dsn := os.Getenv("RZA_MSSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("RZA_MSSQL_TEST_DSN is not configured")
	}
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	suffix := fmt.Sprintf("%d", time.Now().UnixNano())
	parent := "dbo.it_equiv_parent_" + suffix
	child := "dbo.it_equiv_child_" + suffix
	if _, err := tx.ExecContext(ctx, fmt.Sprintf(
		"CREATE TABLE %s ([Id] BIGINT NOT NULL CONSTRAINT [PK_equiv_parent_%s] PRIMARY KEY);",
		parent, suffix,
	)); err != nil {
		t.Fatalf("create equivalence parent: %v", err)
	}

	source := fmt.Sprintf(`CREATE TABLE %s (
  [Id] BIGINT IDENTITY(7,3) NOT NULL,
  [ParentId] BIGINT NULL,
  [Code] VARCHAR(20) COLLATE Latin1_General_BIN NOT NULL CONSTRAINT [DF_equiv_code_%s] DEFAULT ('PENDING'),
  [CreatedAt] DATETIME NOT NULL CONSTRAINT [DF_equiv_created_%s] DEFAULT (GETDATE()),
  [Amount] DECIMAL(18,2) NULL,
  [ComputedAmount] AS (ISNULL([Amount],(0))*(2)) PERSISTED,
  CONSTRAINT [PK_equiv_child_%s] PRIMARY KEY CLUSTERED ([Id] ASC),
  CONSTRAINT [UQ_equiv_code_%s] UNIQUE NONCLUSTERED ([Code] DESC),
  CONSTRAINT [CK_equiv_amount_%s] CHECK ([Amount] IS NULL OR [Amount]>=(0)),
  CONSTRAINT [FK_equiv_parent_%s] FOREIGN KEY ([ParentId]) REFERENCES %s ([Id]) ON DELETE CASCADE
);`, child, suffix, suffix, suffix, suffix, suffix, suffix, parent)
	semantic, err := semanticFromSource(migrationMeta{Category: "tables", Object: child}, source)
	if err != nil {
		t.Fatalf("bootstrap-compatible source rejected: %v", err)
	}
	wrapped := renderPlannedMigration(plannedMigration{
		Category:      "tables",
		Object:        child,
		VK:            1,
		UpSemanticSQL: semantic,
	}, "20991231235959997", "20991231235959997_tables_"+strings.ReplaceAll(child, ".", ".")+"_v1.up.sql", true)

	if _, err := tx.ExecContext(ctx, wrapped); err != nil {
		t.Fatalf("initial table create failed: %v\n%s", err, wrapped)
	}
	if _, err := tx.ExecContext(ctx, wrapped); err != nil {
		t.Fatalf("bootstrap-compatible table did not skip: %v\n%s", err, wrapped)
	}

	var identityOK, defaults, keys, checks, foreignKeys int
	if err := tx.QueryRowContext(ctx, fmt.Sprintf(`
SELECT
  CONVERT(int, CASE WHEN EXISTS (
    SELECT 1 FROM sys.identity_columns
    WHERE object_id=OBJECT_ID(N'%s',N'U') AND seed_value=CONVERT(sql_variant,7) AND increment_value=CONVERT(sql_variant,3)
  ) THEN 1 ELSE 0 END),
  (SELECT COUNT(*) FROM sys.default_constraints WHERE parent_object_id=OBJECT_ID(N'%s',N'U')),
  (SELECT COUNT(*) FROM sys.key_constraints WHERE parent_object_id=OBJECT_ID(N'%s',N'U')),
  (SELECT COUNT(*) FROM sys.check_constraints WHERE parent_object_id=OBJECT_ID(N'%s',N'U')),
  (SELECT COUNT(*) FROM sys.foreign_keys WHERE parent_object_id=OBJECT_ID(N'%s',N'U'));`,
		child, child, child, child, child,
	)).Scan(&identityOK, &defaults, &keys, &checks, &foreignKeys); err != nil {
		t.Fatalf("inspect table catalog: %v", err)
	}
	if identityOK != 1 || defaults != 2 || keys != 2 || checks != 1 || foreignKeys != 1 {
		t.Fatalf("table catalog mismatch: identity=%d defaults=%d keys=%d checks=%d foreign_keys=%d", identityOK, defaults, keys, checks, foreignKeys)
	}

	if _, err := tx.ExecContext(ctx, fmt.Sprintf(
		"ALTER TABLE %s DROP CONSTRAINT [DF_equiv_code_%s]; ALTER TABLE %s DROP CONSTRAINT [UQ_equiv_code_%s]; ALTER TABLE %s ADD CONSTRAINT [DF_equiv_code_%s] DEFAULT ('DRIFT') FOR [Code]; ALTER TABLE %s NOCHECK CONSTRAINT [FK_equiv_parent_%s]; ALTER TABLE %s ALTER COLUMN [Code] VARCHAR(50) COLLATE SQL_Latin1_General_CP1_CI_AS NULL;",
		child, suffix, child, suffix, child, suffix, child, suffix, child,
	)); err != nil {
		t.Fatalf("inject non-blocking DBA-managed drift: %v", err)
	}
	if _, err := tx.ExecContext(ctx, wrapped); err != nil {
		t.Fatalf("length/nullability/collation/default/FK-state drift should not block bootstrap: %v", err)
	}

	if _, err := tx.ExecContext(ctx, fmt.Sprintf(
		"ALTER TABLE %s DROP CONSTRAINT [DF_equiv_created_%s]; ALTER TABLE %s ALTER COLUMN [CreatedAt] DATE NOT NULL;",
		child, suffix, child,
	)); err != nil {
		t.Fatalf("inject incompatible base-type drift: %v", err)
	}
	if _, err := tx.ExecContext(ctx, wrapped); err == nil || !strings.Contains(err.Error(), "Schema drift") {
		t.Fatalf("incompatible base type was not blocked, err=%v", err)
	}
}

// This optional live gate requires a database that has just completed the full
// repository up migration. It executes every generated table up script again
// inside a rollback-only transaction; every script must prove equivalence and
// take its skip path.
func TestIntegration_SQLServerAllGeneratedTableUpScriptsSkip(t *testing.T) {
	if os.Getenv("RZA_MSSQL_FULL_REPLAY_ASSERT") != "1" {
		t.Skip("RZA_MSSQL_FULL_REPLAY_ASSERT is not enabled")
	}
	dsn := os.Getenv("RZA_MSSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("RZA_MSSQL_TEST_DSN is not configured")
	}
	files, err := filepath.Glob(filepath.Join("..", "..", "migrations", "*_tables_*.up.sql"))
	if err != nil {
		t.Fatal(err)
	}
	sort.Strings(files)
	if len(files) == 0 {
		t.Fatal("no generated table up scripts found")
	}

	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	for index, file := range files {
		raw, err := os.ReadFile(file)
		if err != nil {
			t.Fatalf("read %s: %v", filepath.Base(file), err)
		}
		if _, err := tx.ExecContext(ctx, string(raw)); err != nil {
			t.Fatalf("existing-equivalent table script failed at %d/%d %s: %v", index+1, len(files), filepath.Base(file), err)
		}
	}
	t.Logf("all %d generated table scripts proved equivalence and skipped", len(files))
}

func TestUserAuthoredSQLInvalidInputMatrixFailsClosed(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		category string
		object   string
		sql      string
	}{
		{
			name: "table object mismatch", category: "tables", object: "dbo.Expected",
			sql: "CREATE TABLE dbo.Actual ([Id] INT NOT NULL);",
		},
		{
			name: "duplicate table column", category: "tables", object: "dbo.Bad",
			sql: "CREATE TABLE dbo.Bad ([Id] INT NOT NULL, [Id] BIGINT NULL);",
		},
		{
			name: "index targets another table", category: "indexes", object: "dbo.Expected",
			sql: "CREATE INDEX IX_Bad ON dbo.Other ([Id]);",
		},
		{
			name: "view object mismatch", category: "views", object: "dbo.Expected",
			sql: "CREATE OR ALTER VIEW dbo.Actual AS SELECT 1 AS [Id];",
		},
		{
			name: "procedure object mismatch", category: "procedures", object: "dbo.Expected",
			sql: "CREATE OR ALTER PROCEDURE dbo.Actual AS SELECT 1;",
		},
		{
			name: "function missing returns", category: "functions", object: "dbo.BadFn",
			sql: "CREATE OR ALTER FUNCTION dbo.BadFn(@id INT) AS BEGIN RETURN @id; END;",
		},
		{
			name: "unsupported category", category: "synonyms", object: "dbo.BadSynonym",
			sql: "CREATE SYNONYM dbo.BadSynonym FOR dbo.Target;",
		},
		{
			name: "table has no columns", category: "tables", object: "dbo.Empty",
			sql: "CREATE TABLE dbo.Empty ();",
		},
	}
	for _, test := range cases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := semanticFromSource(migrationMeta{Category: test.category, Object: test.object}, test.sql)
			if err == nil {
				t.Fatalf("unsafe/malformed source was accepted:\n%s", test.sql)
			}
		})
	}
}

func applyAuthoredSource(t *testing.T, ctx context.Context, tx *sql.Tx, category, object, source, objectType string) {
	t.Helper()
	semantic, err := semanticFromSource(migrationMeta{Category: category, Object: object}, source)
	if err != nil {
		t.Fatalf("%s source rejected: %v\n%s", category, err, source)
	}
	fileObject := strings.TrimPrefix(object, "dbo.")
	fileName := fmt.Sprintf("20991231235959998_%s_dbo.%s_v1.up.sql", category, fileObject)
	wrapped := renderPlannedMigration(plannedMigration{
		Category:      category,
		Object:        object,
		VK:            1,
		UpSemanticSQL: semantic,
	}, "20991231235959998", fileName, true)
	if _, err := tx.ExecContext(ctx, wrapped); err != nil {
		t.Fatalf("%s migration failed to compile/apply: %v\n%s", category, err, wrapped)
	}
	if objectType == "" {
		return
	}
	var exists int
	if err := tx.QueryRowContext(ctx,
		"SELECT CONVERT(int, CASE WHEN OBJECT_ID(@p1, @p2) IS NULL THEN 0 ELSE 1 END);",
		object, objectType,
	).Scan(&exists); err != nil {
		t.Fatalf("inspect %s: %v", object, err)
	}
	if exists != 1 {
		t.Fatalf("%s %s was not created", category, object)
	}
}
