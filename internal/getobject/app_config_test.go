package getobject

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadAppConfigAndValidate(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	cfgPath := filepath.Join(root, "app_config.get_object.json")
	raw := `{
  "connection_string": "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable",
  "repo_root": ".",
  "db_object_dir": "db_object",
  "schema": "dbo",
  "connect_timeout_seconds": 10,
  "query_timeout_seconds": 60,
  "render_workers": 6,
  "write_workers": 3,
  "skip_unsupported_tables": true,
  "strip_module_comments": true,
  "reindent_module_sql": true,
  "indent_size": 4,
  "module_formatter_backend": "lite"
}
`
	if err := os.WriteFile(cfgPath, []byte(raw), 0o644); err != nil {
		t.Fatalf("write config failed: %v", err)
	}

	cfg, err := LoadAppConfig(cfgPath)
	if err != nil {
		t.Fatalf("load config failed: %v", err)
	}
	if cfg.Schema != "dbo" {
		t.Fatalf("unexpected schema: %s", cfg.Schema)
	}
	if cfg.ConnectTimeoutSeconds != 10 {
		t.Fatalf("unexpected connect timeout: %d", cfg.ConnectTimeoutSeconds)
	}
	if cfg.RenderWorkers != 6 {
		t.Fatalf("unexpected render workers: %d", cfg.RenderWorkers)
	}
	if cfg.WriteWorkers != 3 {
		t.Fatalf("unexpected write workers: %d", cfg.WriteWorkers)
	}
	if !cfg.SkipUnsupportedTables {
		t.Fatalf("unexpected skip_unsupported_tables: %v", cfg.SkipUnsupportedTables)
	}
	if !cfg.StripModuleComments {
		t.Fatalf("unexpected strip_module_comments: %v", cfg.StripModuleComments)
	}
	if !cfg.ReindentModuleSQL {
		t.Fatalf("unexpected reindent_module_sql: %v", cfg.ReindentModuleSQL)
	}
	if cfg.IndentSize != 4 {
		t.Fatalf("unexpected indent size: %d", cfg.IndentSize)
	}
	if cfg.ModuleFormatterBackend != "lite" {
		t.Fatalf("unexpected module formatter backend: %s", cfg.ModuleFormatterBackend)
	}
}

func TestValidateRejectsNonDBOSchema(t *testing.T) {
	t.Parallel()

	cfg := DefaultAppConfig(".")
	cfg.ConnectionString = "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable"
	cfg.Schema = "sales"
	if err := cfg.Validate(); err == nil {
		t.Fatalf("expected validate error for non-dbo schema")
	}
}

func TestNormalizeAppliesWorkerDefaults(t *testing.T) {
	t.Parallel()

	cfg := AppConfig{
		ConnectionString: "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable",
		RepoRoot:         ".",
		DbObjectDir:      "db_object",
		Schema:           "dbo",
	}
	cfg.Normalize()
	if cfg.RenderWorkers <= 0 {
		t.Fatalf("expected render workers default > 0")
	}
	if cfg.WriteWorkers <= 0 {
		t.Fatalf("expected write workers default > 0")
	}
	if cfg.IndentSize <= 0 {
		t.Fatalf("expected indent_size default > 0")
	}
	if cfg.ModuleFormatterBackend != "lite" {
		t.Fatalf("expected default module_formatter_backend=lite, got %s", cfg.ModuleFormatterBackend)
	}
}

func TestDefaultAppConfigUsesLiteBackend(t *testing.T) {
	t.Parallel()

	cfg := DefaultAppConfig(".")
	if cfg.ModuleFormatterBackend != "lite" {
		t.Fatalf("expected DefaultAppConfig module formatter backend lite, got %s", cfg.ModuleFormatterBackend)
	}
	if cfg.SkipUnsupportedTables {
		t.Fatalf("expected DefaultAppConfig skip_unsupported_tables=false for complete source-of-truth snapshots")
	}
}
