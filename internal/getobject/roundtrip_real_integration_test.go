package getobject

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

func TestIntegration_SQLServerFreshReplayMatchesDBObjectSnapshots(t *testing.T) {
	dsn := strings.TrimSpace(os.Getenv("RZA_MSSQL_TEST_DSN"))
	if dsn == "" {
		t.Skip("RZA_MSSQL_TEST_DSN is not configured")
	}
	repo, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	snapshotRoot := t.TempDir()
	cfg := DefaultAppConfig(snapshotRoot)
	cfg.ConnectionString = dsn
	cfg.RenderWorkers = 4
	cfg.WriteWorkers = 4
	if _, err := ExportFromDatabaseWithOptions(cfg, ExportOptions{}); err != nil {
		t.Fatalf("extract replayed database: %v", err)
	}

	var mismatches []string
	for _, category := range []string{"types", "tables", "indexes", "views", "functions", "procedures"} {
		wantDir := filepath.Join(repo, "db_object", category)
		gotDir := filepath.Join(snapshotRoot, "db_object", category)
		wantNames := sqlFileNames(t, wantDir)
		gotNames := sqlFileNames(t, gotDir)
		if strings.Join(wantNames, "\n") != strings.Join(gotNames, "\n") {
			mismatches = append(mismatches, category+": filename set differs")
			continue
		}
		for _, name := range wantNames {
			wantRaw, err := os.ReadFile(filepath.Join(wantDir, name))
			if err != nil {
				t.Fatal(err)
			}
			gotRaw, err := os.ReadFile(filepath.Join(gotDir, name))
			if err != nil {
				t.Fatal(err)
			}
			want := normalizeSQL(string(wantRaw))
			got := normalizeSQL(string(gotRaw))
			if want == got {
				continue
			}
			mismatches = append(mismatches, filepath.ToSlash(filepath.Join("db_object", category, name))+": "+firstLineDifference(want, got))
		}
	}
	if len(mismatches) > 0 {
		limit := len(mismatches)
		if limit > 20 {
			limit = 20
		}
		t.Fatalf("fresh replay is not extractor-stable: mismatches=%d\n%s", len(mismatches), strings.Join(mismatches[:limit], "\n"))
	}
}

func sqlFileNames(t *testing.T, dir string) []string {
	t.Helper()
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}
		}
		t.Fatal(err)
	}
	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() && strings.EqualFold(filepath.Ext(entry.Name()), ".sql") {
			names = append(names, entry.Name())
		}
	}
	sort.Strings(names)
	return names
}

func firstLineDifference(want, got string) string {
	wantLines := strings.Split(want, "\n")
	gotLines := strings.Split(got, "\n")
	count := len(wantLines)
	if len(gotLines) < count {
		count = len(gotLines)
	}
	for index := 0; index < count; index++ {
		if wantLines[index] != gotLines[index] {
			return "line " + itoa(index+1) + " want=" + quotePreview(wantLines[index]) + " got=" + quotePreview(gotLines[index])
		}
	}
	return "line count differs"
}

func itoa(value int) string {
	if value == 0 {
		return "0"
	}
	var digits [20]byte
	index := len(digits)
	for value > 0 {
		index--
		digits[index] = byte('0' + value%10)
		value /= 10
	}
	return string(digits[index:])
}

func quotePreview(value string) string {
	value = strings.TrimSpace(value)
	if len(value) > 160 {
		value = value[:157] + "..."
	}
	return `"` + value + `"`
}
