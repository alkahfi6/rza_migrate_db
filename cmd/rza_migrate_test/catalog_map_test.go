package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPrepareCatalogMappedMigrationsUsesTemporaryCopy(t *testing.T) {
	source := t.TempDir()
	path := filepath.Join(source, "001_test.up.sql")
	original := "CREATE VIEW dbo.V AS SELECT * FROM SQL_REPLICATE..T; SELECT N'SQL_REPLICATE.dbo.X'; -- SQL_REPLICATE\n"
	if err := os.WriteFile(path, []byte(original), 0o644); err != nil {
		t.Fatal(err)
	}
	mapped, cleanup, err := prepareCatalogMappedMigrations(source, "SQL_REPLICATE=RZA_RT_REKSA_SQL_REPLICATE")
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()
	raw, err := os.ReadFile(filepath.Join(mapped, filepath.Base(path)))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(raw), "FROM SQL_REPLICATE") || strings.Contains(string(raw), "N'SQL_REPLICATE") || !strings.Contains(string(raw), "RZA_RT_REKSA_SQL_REPLICATE..T") {
		t.Fatalf("mapping not applied:\n%s", raw)
	}
	sourceRaw, err := os.ReadFile(path)
	if err != nil || string(sourceRaw) != original {
		t.Fatalf("source migration changed: err=%v content=%s", err, sourceRaw)
	}
}

func TestParseCatalogMappingsRejectsUnsafeValues(t *testing.T) {
	for _, input := range []string{"bad", "SQL_REPLICATE=master", "A=A", "A=B,A=C", "A;DROP=B"} {
		if _, err := parseCatalogMappings(input); err == nil {
			t.Fatalf("expected %q to fail", input)
		}
	}
}
