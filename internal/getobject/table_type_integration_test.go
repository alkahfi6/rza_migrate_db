package getobject

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

func TestIntegration_SQLServerExtractTableType(t *testing.T) {
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
	suffix := fmt.Sprintf("%d", time.Now().UnixNano())
	name := "rza_ExtractType_" + suffix
	qualified := "dbo." + name
	createSQL := "CREATE TYPE " + qualified + ` AS TABLE (
  [Id] INT NOT NULL,
  [Code] VARCHAR(30) NOT NULL,
  [Qty] INT NULL DEFAULT ((0)),
  CHECK ([Qty] IS NULL OR [Qty] >= 0),
  PRIMARY KEY CLUSTERED ([Id]),
  UNIQUE NONCLUSTERED ([Code]),
  INDEX [IX_` + name + `_Qty] NONCLUSTERED ([Qty])
);`
	if _, err := db.ExecContext(ctx, createSQL); err != nil {
		t.Fatalf("create table type sandbox: %v\n%s", err, createSQL)
	}
	defer func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cleanupCancel()
		if _, err := db.ExecContext(cleanupCtx, "DROP TYPE IF EXISTS "+qualified+";"); err != nil {
			t.Errorf("cleanup table type sandbox: %v", err)
		}
	}()

	cfg := DefaultAppConfig(t.TempDir())
	cfg.ConnectionString = dsn
	cfg.Schema = "dbo"
	files, warnings, _, err := loadTableTypeFiles(db, cfg)
	if err != nil {
		t.Fatal(err)
	}
	fileName := fileNameForObject(name)
	got, exists := files[fileName]
	if !exists {
		t.Fatalf("extracted table type %s not found; warnings=%v", fileName, warnings)
	}
	for _, want := range []string{
		"CREATE TYPE dbo." + name + " AS TABLE",
		"[Id] INT NOT NULL",
		"[Code] VARCHAR(30)",
		"DEFAULT",
		"CHECK",
		"PRIMARY KEY CLUSTERED",
		"UNIQUE NONCLUSTERED",
		"INDEX [IX_" + name + "_Qty] NONCLUSTERED",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("extracted table type missing %q:\n%s", want, got)
		}
	}

	if _, err := db.ExecContext(ctx, "DROP TYPE "+qualified+";"); err != nil {
		t.Fatalf("drop original table type before replay: %v", err)
	}
	if _, err := db.ExecContext(ctx, got); err != nil {
		t.Fatalf("replay extracted table type: %v\n%s", err, got)
	}
	var columnCount, indexCount int
	if err := db.QueryRowContext(ctx, `
SELECT COUNT(1)
FROM sys.table_types tt
JOIN sys.schemas s ON s.schema_id = tt.schema_id
JOIN sys.columns c ON c.object_id = tt.type_table_object_id
WHERE s.name = N'dbo' AND tt.name = @p1;
`, name).Scan(&columnCount); err != nil {
		t.Fatal(err)
	}
	if err := db.QueryRowContext(ctx, `
SELECT COUNT(1)
FROM sys.table_types tt
JOIN sys.schemas s ON s.schema_id = tt.schema_id
JOIN sys.indexes i ON i.object_id = tt.type_table_object_id
WHERE s.name = N'dbo' AND tt.name = @p1 AND i.name = @p2;
`, name, "IX_"+name+"_Qty").Scan(&indexCount); err != nil {
		t.Fatal(err)
	}
	if columnCount != 3 || indexCount != 1 {
		t.Fatalf("replayed table type metadata mismatch: columns=%d indexes=%d", columnCount, indexCount)
	}
}

func TestIntegration_SQLServerExtractAllFunctionSubtypes(t *testing.T) {
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
	suffix := fmt.Sprintf("%d", time.Now().UnixNano())
	definitions := map[string]string{
		"rza_ExtractScalar_" + suffix: "(@value INT = 1) RETURNS INT AS BEGIN RETURN @value; END;",
		"rza_ExtractInline_" + suffix: "(@value INT) RETURNS TABLE AS RETURN (SELECT @value AS [Value]);",
		"rza_ExtractMulti_" + suffix:  "(@value INT) RETURNS @result TABLE ([Id] INT NOT NULL PRIMARY KEY, [Value] INT NULL) AS BEGIN INSERT @result VALUES (1, @value); RETURN; END;",
	}
	names := make([]string, 0, len(definitions))
	defer func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cleanupCancel()
		for _, name := range names {
			if _, err := db.ExecContext(cleanupCtx, "DROP FUNCTION IF EXISTS dbo."+name+";"); err != nil {
				t.Errorf("cleanup function %s: %v", name, err)
			}
		}
	}()
	for name, body := range definitions {
		names = append(names, name)
		if _, err := db.ExecContext(ctx, "CREATE OR ALTER FUNCTION dbo."+name+body); err != nil {
			t.Fatalf("create function %s: %v", name, err)
		}
	}

	cfg := DefaultAppConfig(t.TempDir())
	cfg.ConnectionString = dsn
	cfg.Schema = "dbo"
	files, warnings, _, err := loadProgrammabilityFiles(db, cfg, "FN,IF,TF", "FUNCTION")
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range names {
		fileName := fileNameForObject(name)
		got, exists := files[fileName]
		if !exists {
			t.Fatalf("function subtype %s was not extracted; warnings=%v", fileName, warnings)
		}
		if !strings.Contains(strings.ToUpper(got), "CREATE OR ALTER FUNCTION DBO."+strings.ToUpper(name)) {
			t.Fatalf("function extraction is not canonical:\n%s", got)
		}
	}
}
