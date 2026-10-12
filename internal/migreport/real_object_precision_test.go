package migreport

import (
	"crypto/sha256"
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func sourceTreeHashes(t *testing.T, root string) map[string][32]byte {
	t.Helper()
	out := map[string][32]byte{}
	for _, dir := range []string{"db_object", "migrations"} {
		path := filepath.Join(root, dir)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}
		if err := filepath.WalkDir(path, func(path string, entry fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if entry.IsDir() {
				return nil
			}
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			rel, _ := filepath.Rel(root, path)
			out[filepath.ToSlash(rel)] = sha256.Sum256(data)
			return nil
		}); err != nil {
			t.Fatal(err)
		}
	}
	return out
}

func TestRealObjectPrecisionFixtureReadOnlyAndIdempotent(t *testing.T) {
	fixture := filepath.Join("..", "..", "fixtures", "report_object_precision")
	root := t.TempDir()
	if err := filepath.WalkDir(filepath.Join(fixture, "input"), func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(filepath.Join(fixture, "input"), path)
		target := filepath.Join(root, rel)
		if entry.IsDir() {
			return os.MkdirAll(target, 0755)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, 0644)
	}); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(root, "migrations"), 0755); err != nil {
		t.Fatal(err)
	}
	before := sourceTreeHashes(t, root)
	docs, err := BuildDesiredStateDocs(root)
	if err != nil {
		t.Fatal(err)
	}
	indexes := 0
	for _, table := range docs.Tables {
		if table.Name == "dbo.JHCLDR" || table.Name == "dbo.MXHolidayCurrency" {
			indexes += len(table.Indexes)
		}
	}
	joinType := ""
	for _, r := range docs.Relationships {
		if r.LeftObject == "dbo.ORI_Order_TT" && r.RightObject == "dbo.TRSProcessBuyBackLog_TH" {
			joinType = strings.Join(r.JoinTypes, ",")
		}
	}
	consumer := ""
	for _, table := range docs.Tables {
		if table.Name == "dbo.ReksaProduct_TM" {
			for _, c := range table.Consumers {
				if c.Kind == "function" {
					consumer = c.Name
				}
			}
		}
	}
	unknown := ""
	for _, p := range docs.ProcedureProfiles {
		if p.Name == "dbo.TRSParamIssuerSave" {
			for _, m := range p.Mappings {
				if m.Column == "IssuerCode" {
					unknown = m.Resolution
				}
			}
		}
	}
	actual := map[string]any{"table_count": len(docs.Tables), "module_count": len(docs.Modules), "multiline_indexes": indexes, "outer_join": "dbo.ORI_Order_TT.OrderId = dbo.TRSProcessBuyBackLog_TH.OrderId", "join_type": joinType, "function_consumer": consumer, "function_table": "dbo.ReksaProduct_TM", "unknown_column": "dbo.TRSParamIssuer_TH.IssuerCode", "unknown_resolution": unknown}
	raw, err := json.Marshal(actual)
	if err != nil {
		t.Fatal(err)
	}
	normalized := map[string]any{}
	if err := json.Unmarshal(raw, &normalized); err != nil {
		t.Fatal(err)
	}
	goldenData, err := os.ReadFile(filepath.Join(fixture, "expected", "output.golden.json"))
	if err != nil {
		t.Fatal(err)
	}
	golden := map[string]any{}
	if err := json.Unmarshal(goldenData, &golden); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(normalized, golden) {
		t.Fatalf("real-source contract mismatch:\n%s\nwant:\n%s", raw, goldenData)
	}
	for run := 0; run < 2; run++ {
		if code, err := Run(Options{RepoRoot: root}); err != nil || code != 0 {
			t.Fatalf("run %d: code=%d err=%v", run, code, err)
		}
	}
	again, err := BuildDesiredStateDocs(root)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(docs, again) {
		t.Fatal("source interpretation changed on second run")
	}
	if !reflect.DeepEqual(before, sourceTreeHashes(t, root)) {
		t.Fatal("report modified source/history")
	}
	lineage, err := os.ReadFile(filepath.Join(root, "docs", "database", "relationships", "module_lineage.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(lineage), "Resolution") || !strings.Contains(string(lineage), "FUNCTION") {
		t.Fatal("lineage resolution/function not published")
	}
}

func TestReportRealSnapshotCorpus(t *testing.T) {
	base := filepath.Join("..", "..", "all_object")
	if _, err := os.Stat(base); os.IsNotExist(err) {
		t.Skip("optional local snapshots unavailable; committed real fixture still runs")
	}
	for _, name := range []string{"bond/bond-sqltrsretail", "bond/bond-sqlreplicate", "reksa/reksa-sqlreksa2", "reksa/reksa-sqlreplicate"} {
		t.Run(name, func(t *testing.T) {
			root := filepath.Join(base, filepath.FromSlash(name))
			if _, err := os.Stat(root); os.IsNotExist(err) {
				t.Skip("snapshot unavailable")
			}
			before := sourceTreeHashes(t, root)
			docs, err := BuildDesiredStateDocs(root)
			if err != nil {
				t.Fatal(err)
			}
			modules := 0
			for _, dir := range []string{"views", "functions", "procedures"} {
				entries, err := os.ReadDir(filepath.Join(root, "db_object", dir))
				if err != nil && !os.IsNotExist(err) {
					t.Fatal(err)
				}
				for _, e := range entries {
					if !e.IsDir() && strings.EqualFold(filepath.Ext(e.Name()), ".sql") {
						modules++
					}
				}
			}
			if modules != len(docs.Modules) {
				t.Fatalf("module coverage: %d/%d", len(docs.Modules), modules)
			}
			for _, table := range docs.Tables {
				if len(table.Columns) == 0 {
					t.Errorf("table has no recognized columns: %s", table.Name)
				}
			}
			if !reflect.DeepEqual(before, sourceTreeHashes(t, root)) {
				t.Fatal("corpus mutated")
			}
			mappings, unknown := 0, 0
			for _, p := range docs.ProcedureProfiles {
				for _, m := range p.Mappings {
					mappings++
					if m.Resolution == "COLUMN_NOT_FOUND" {
						unknown++
					}
				}
			}
			t.Logf("tables=%d modules=%d relationships=%d profiles=%d mappings=%d missing_column=%d unchanged_files=%d", len(docs.Tables), len(docs.Modules), len(docs.Relationships), len(docs.ProcedureProfiles), mappings, unknown, len(before))
		})
	}
}
