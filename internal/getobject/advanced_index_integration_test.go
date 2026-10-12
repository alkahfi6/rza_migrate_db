package getobject

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	_ "github.com/microsoft/go-mssqldb"
)

// This test proves database -> db_object fidelity for advanced index families.
// Objects use unique names and are always removed from the shared test DB.
func TestIntegration_ExtractorAdvancedIndexFamilies(t *testing.T) {
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
	suffix := fmt.Sprintf("%d", time.Now().UnixNano())
	xmlTable := "rza_extract_xml_" + suffix
	spatialTable := "rza_extract_spatial_" + suffix
	columnstoreTable := "rza_extract_cs_" + suffix
	compressionTable := "rza_extract_cmp_" + suffix
	tables := []string{xmlTable, spatialTable, columnstoreTable, compressionTable}
	defer func() {
		for _, table := range tables {
			_, _ = db.ExecContext(context.Background(), "DROP TABLE IF EXISTS dbo."+table)
		}
	}()

	setup := []string{
		"CREATE TABLE dbo." + xmlTable + " ([id] INT NOT NULL CONSTRAINT [PK_" + xmlTable + "] PRIMARY KEY, [doc] XML NULL); CREATE PRIMARY XML INDEX [PXML_" + xmlTable + "] ON dbo." + xmlTable + " ([doc]); CREATE XML INDEX [SXML_" + xmlTable + "] ON dbo." + xmlTable + " ([doc]) USING XML INDEX [PXML_" + xmlTable + "] FOR VALUE;",
		"CREATE TABLE dbo." + spatialTable + " ([id] INT NOT NULL CONSTRAINT [PK_" + spatialTable + "] PRIMARY KEY CLUSTERED, [shape] geometry NULL); CREATE SPATIAL INDEX [SIX_" + spatialTable + "] ON dbo." + spatialTable + " ([shape]) USING GEOMETRY_AUTO_GRID WITH (BOUNDING_BOX=(-180,-90,180,90), CELLS_PER_OBJECT=12);",
		"CREATE TABLE dbo." + columnstoreTable + " ([id] BIGINT NOT NULL, [amount] DECIMAL(18,2) NULL); CREATE NONCLUSTERED COLUMNSTORE INDEX [NCCI_" + columnstoreTable + "] ON dbo." + columnstoreTable + " ([id],[amount]);",
		"CREATE TABLE dbo." + compressionTable + " ([id] BIGINT NOT NULL, [code] VARCHAR(50) NULL); CREATE INDEX [IX_" + compressionTable + "] ON dbo." + compressionTable + " ([id]) INCLUDE ([code]) WITH (DATA_COMPRESSION=PAGE, FILLFACTOR=80, PAD_INDEX=ON, ALLOW_ROW_LOCKS=OFF, ALLOW_PAGE_LOCKS=OFF);",
	}
	for _, statement := range setup {
		if _, err := db.ExecContext(ctx, statement); err != nil {
			t.Fatalf("advanced extractor setup failed: %v", err)
		}
	}

	root := t.TempDir()
	cfg := DefaultAppConfig(root)
	cfg.ConnectionString = dsn
	cfg.SkipUnsupportedTables = true
	if _, err := ExportFromDatabaseWithOptions(cfg, ExportOptions{}); err != nil {
		t.Fatalf("advanced index export failed: %v", err)
	}
	expectations := map[string][]string{
		xmlTable:         {"CREATE PRIMARY XML INDEX", "CREATE XML INDEX", "FOR VALUE"},
		spatialTable:     {"CREATE SPATIAL INDEX", "GEOMETRY_AUTO_GRID", "BOUNDING_BOX", "CELLS_PER_OBJECT = 12"},
		columnstoreTable: {"CREATE NONCLUSTERED COLUMNSTORE INDEX"},
		compressionTable: {"DATA_COMPRESSION = PAGE", "FILLFACTOR = 80", "PAD_INDEX = ON", "ALLOW_ROW_LOCKS = OFF", "ALLOW_PAGE_LOCKS = OFF"},
	}
	for table, tokens := range expectations {
		raw, err := os.ReadFile(filepath.Join(root, "db_object", "indexes", "dbo."+table+".sql"))
		if err != nil {
			t.Fatalf("extracted advanced index file missing for %s: %v", table, err)
		}
		upper := strings.ToUpper(string(raw))
		for _, token := range tokens {
			if !strings.Contains(upper, token) {
				t.Fatalf("extractor lost %s for %s:\n%s", token, table, raw)
			}
		}
	}
}
