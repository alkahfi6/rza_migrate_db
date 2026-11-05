package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/microsoft/go-mssqldb/msdsn"
	"rza_migrate_db/internal/getobject"
)

func TestConfigPrecedenceAndOfflinePreview(t *testing.T) {
	t.Chdir(t.TempDir())
	t.Setenv("MSSQL_DSN", "primary-env")
	t.Setenv("RZA_GET_OBJECT_DSN", "alias-env")
	cfg, err := loadGetObjectConfig("", false)
	if err != nil || cfg.ConnectionString != "primary-env" {
		t.Fatal("primary environment fallback failed")
	}
	t.Setenv("MSSQL_DSN", "")
	cfg, err = loadGetObjectConfig("", false)
	if err != nil || cfg.ConnectionString != "alias-env" {
		t.Fatal("compatibility environment fallback failed")
	}
	for _, name := range []string{"config.yaml", "explicit.yaml"} {
		raw := "database:\n  host: localhost\n  user: test\n  password: 'p@ss#%/?'\n  initial_catalog: '" + name + "'\n"
		if err := os.WriteFile(name, []byte(raw), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	for _, path := range []string{"", "explicit.yaml"} {
		cfg, err = loadGetObjectConfig(path, false)
		if err != nil {
			t.Fatal(err)
		}
		parsed, err := msdsn.Parse(cfg.ConnectionString)
		want := path
		if want == "" {
			want = "config.yaml"
		}
		if err != nil || parsed.Database != want {
			t.Fatal("file precedence failed")
		}
	}
	if err := os.WriteFile("config.yaml", []byte("invalid: ["), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := loadGetObjectConfig("", false); err == nil {
		t.Fatal("invalid local YAML must not silently use environment")
	}
	cfg, err = loadGetObjectConfig("", true)
	if err != nil || cfg.ConnectionString != "alias-env" {
		t.Fatal("offline preview must not auto-load credentials")
	}
	if _, err := loadGetObjectConfig("missing.yaml", false); err == nil {
		t.Fatal("missing explicit config must fail")
	}
}

func TestRunYAMLConfigReachesExporterWithoutChangingMigrations(t *testing.T) {
	root := t.TempDir()
	t.Chdir(root)
	if err := os.Mkdir("migrations", 0o755); err != nil {
		t.Fatal(err)
	}
	const history = "-- immutable migration evidence\n"
	if err := os.WriteFile("migrations/history.up.sql", []byte(history), 0o644); err != nil {
		t.Fatal(err)
	}
	raw := "database:\n  host: localhost\n  user: 'user@name'\n  password: ' p@ss:#?/%40&+ '\n  initial_catalog: 'Sales & Ops'\n"
	if err := os.WriteFile("config.yaml", []byte(raw), 0o600); err != nil {
		t.Fatal(err)
	}
	prevExport := exportFromDatabase
	t.Cleanup(func() { exportFromDatabase = prevExport })
	calls := 0
	exportFromDatabase = func(cfg getobject.AppConfig, opts getobject.ExportOptions) (getobject.ExportSummary, error) {
		calls++
		parsed, err := msdsn.Parse(cfg.ConnectionString)
		if err != nil || parsed.Password != " p@ss:#?/%40&+ " || parsed.Database != "Sales & Ops" || parsed.User != "user@name" {
			t.Fatal("exporter did not receive literal YAML credentials")
		}
		return getobject.ExportSummary{DryRun: opts.DryRun}, nil
	}
	for _, flags := range [][]string{{}, {"-dry-run"}, {"-config", "config.yaml", "-dry-run"}} {
		if code := run(append([]string{"rza_get_object_db"}, flags...)); code != 0 {
			t.Fatalf("exit = %d", code)
		}
		summary, err := os.ReadFile(filepath.Join("out", "get_object_summary.json"))
		if err != nil || strings.Contains(string(summary), "p@ss") || strings.Contains(string(summary), "sqlserver://") {
			t.Fatal("missing summary or credential disclosure")
		}
	}
	if calls != 3 {
		t.Fatal("export was not called for every CLI mode")
	}
	entries, err := os.ReadDir("migrations")
	got, readErr := os.ReadFile("migrations/history.up.sql")
	if err != nil || readErr != nil || len(entries) != 1 || string(got) != history {
		t.Fatal("extractor changed migration history")
	}
	got, err = os.ReadFile("config.yaml")
	if err != nil || string(got) != raw {
		t.Fatal("extractor rewrote credentials")
	}
}
