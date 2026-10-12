package migrest

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadConfigResolvesPathsRelativeToConfigFile(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	confDir := filepath.Join(root, "runtime")
	if err := os.MkdirAll(confDir, 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	configPath := filepath.Join(confDir, "config.yaml")
	content := `
repo_root: ../repo
server:
  listen: ":8090"
security:
  enable_api_key: true
  api_key: "test-key"
migration:
  database_url: "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable"
  migrations_dir: migrations
  allow_down: false
  allow_run_specific_file: false
  operation_timeout_seconds: 120
  db_lock:
    enabled: false
`
	if err := os.WriteFile(configPath, []byte(strings.TrimSpace(content)+"\n"), 0o644); err != nil {
		t.Fatalf("write config failed: %v", err)
	}

	cfg, err := LoadConfig(configPath)
	if err != nil {
		t.Fatalf("load config failed: %v", err)
	}
	wantRepo := filepath.Clean(filepath.Join(confDir, "../repo"))
	if cfg.RepoRoot != wantRepo {
		t.Fatalf("repo root mismatch: got=%s want=%s", cfg.RepoRoot, wantRepo)
	}
	wantMigrations := filepath.Join(wantRepo, "migrations")
	if cfg.Migration.MigrationsDir != wantMigrations {
		t.Fatalf("migrations dir mismatch: got=%s want=%s", cfg.Migration.MigrationsDir, wantMigrations)
	}
}

func TestLoadConfigSupportsSplitDatabaseFields(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	configPath := filepath.Join(root, "config.yaml")
	content := `
repo_root: .
server:
  listen: ":8080"
security:
  enable_api_key: true
  api_key: "test-key"
migration:
  migrations_dir: migrations
  allow_down: false
  allow_run_specific_file: false
  operation_timeout_seconds: 120
  database:
    server: "127.0.0.1"
    port: 1433
    username: "sa"
    password: "Password1!"
    database: "master"
    encrypt: "disable"
  db_lock:
    enabled: false
`
	if err := os.WriteFile(configPath, []byte(strings.TrimSpace(content)+"\n"), 0o644); err != nil {
		t.Fatalf("write config failed: %v", err)
	}

	cfg, err := LoadConfig(configPath)
	if err != nil {
		t.Fatalf("load config failed: %v", err)
	}
	if got := cfg.Migration.DatabaseURL; got == "" {
		t.Fatalf("expected generated database url, got empty")
	}
	wantPrefix := "sqlserver://sa:Password1%21@127.0.0.1:1433?"
	if !strings.HasPrefix(cfg.Migration.DatabaseURL, wantPrefix) {
		t.Fatalf("database url prefix mismatch: got=%s want-prefix=%s", cfg.Migration.DatabaseURL, wantPrefix)
	}
}

func TestLoadConfigPlaceholderURLCanBeOverriddenBySplitFields(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	configPath := filepath.Join(root, "config.yaml")
	content := `
repo_root: .
server:
  listen: ":8080"
security:
  enable_api_key: true
  api_key: "test-key"
migration:
  database_url: "sqlserver://username:password@localhost:1433?database=YourDatabase&encrypt=disable"
  migrations_dir: migrations
  allow_down: false
  allow_run_specific_file: false
  operation_timeout_seconds: 120
  database:
    server: "db.prod.local"
    port: 1433
    username: "svc_mig"
    password: "TopSecret1!"
    database: "MainDB"
    encrypt: "disable"
  db_lock:
    enabled: false
`
	if err := os.WriteFile(configPath, []byte(strings.TrimSpace(content)+"\n"), 0o644); err != nil {
		t.Fatalf("write config failed: %v", err)
	}

	cfg, err := LoadConfig(configPath)
	if err != nil {
		t.Fatalf("load config failed: %v", err)
	}
	if strings.Contains(cfg.Migration.DatabaseURL, "username:password@localhost") {
		t.Fatalf("expected placeholder URL to be replaced, got=%s", cfg.Migration.DatabaseURL)
	}
}

func TestLoadConfigRejectsPlaceholderValues(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	configPath := filepath.Join(root, "config.yaml")
	content := `
repo_root: .
server:
  listen: ":8080"
security:
  enable_api_key: true
  api_key: "CHANGE_ME"
migration:
  database_url: "sqlserver://username:password@localhost:1433?database=YourDatabase&encrypt=disable"
  migrations_dir: migrations
  allow_down: false
  allow_run_specific_file: false
  operation_timeout_seconds: 120
  db_lock:
    enabled: false
`
	if err := os.WriteFile(configPath, []byte(strings.TrimSpace(content)+"\n"), 0o644); err != nil {
		t.Fatalf("write config failed: %v", err)
	}

	_, err := LoadConfig(configPath)
	if err == nil {
		t.Fatalf("expected load config to fail for placeholder values")
	}
}

func TestLoadConfigSupportsServerInstance(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	configPath := filepath.Join(root, "config.yaml")
	content := `
repo_root: .
server:
  listen: ":8080"
security:
  enable_api_key: true
  api_key: "test-key"
migration:
  migrations_dir: migrations
  allow_down: false
  allow_run_specific_file: false
  operation_timeout_seconds: 120
  database:
    server: "dbserver\\SQLEXPRESS"
    username: "sa"
    password: "Password1!"
    database: "master"
    encrypt: "disable"
  db_lock:
    enabled: false
`
	if err := os.WriteFile(configPath, []byte(strings.TrimSpace(content)+"\n"), 0o644); err != nil {
		t.Fatalf("write config failed: %v", err)
	}

	cfg, err := LoadConfig(configPath)
	if err != nil {
		t.Fatalf("load config failed: %v", err)
	}
	if !strings.Contains(cfg.Migration.DatabaseURL, "instance=SQLEXPRESS") {
		t.Fatalf("expected generated url to contain instance query param, got=%s", cfg.Migration.DatabaseURL)
	}
}

func TestLoadConfigRejectsRemovedTimeoutModeField(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	configPath := filepath.Join(root, "config.yaml")
	content := `
repo_root: .
server:
  listen: ":8080"
security:
  enable_api_key: true
  api_key: "test-key"
migration:
  migrations_dir: migrations
  allow_down: false
  allow_run_specific_file: false
  operation_timeout_seconds: 120
  operation_timeout_mode: "best_effort"
  database:
    server: "localhost"
    port: 1433
    username: "sa"
    password: "Password1!"
    database: "master"
    encrypt: "disable"
  db_lock:
    enabled: false
`
	if err := os.WriteFile(configPath, []byte(strings.TrimSpace(content)+"\n"), 0o644); err != nil {
		t.Fatalf("write config failed: %v", err)
	}

	_, err := LoadConfig(configPath)
	if err == nil {
		t.Fatalf("expected load config to fail because operation_timeout_mode is removed")
	}
}

func TestLoadConfigSupportsTelemetry(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	configPath := filepath.Join(root, "config.yaml")
	content := `
repo_root: .
server:
  listen: ":8080"
security:
  enable_api_key: true
  api_key: "test-key"
telemetry:
  enabled: true
  otlp_endpoint: "http://otel-collector:4318"
  insecure: true
  sample_ratio: 1.0
migration:
  migrations_dir: migrations
  allow_down: false
  allow_run_specific_file: false
  operation_timeout_seconds: 120
  database:
    server: "localhost"
    port: 1433
    username: "sa"
    password: "Password1!"
    database: "master"
    encrypt: "disable"
  db_lock:
    enabled: false
`
	if err := os.WriteFile(configPath, []byte(strings.TrimSpace(content)+"\n"), 0o644); err != nil {
		t.Fatalf("write config failed: %v", err)
	}

	cfg, err := LoadConfig(configPath)
	if err != nil {
		t.Fatalf("load config failed: %v", err)
	}
	if !cfg.Telemetry.Enabled {
		t.Fatalf("expected telemetry enabled")
	}
	if cfg.Telemetry.Endpoint != "http://otel-collector:4318" {
		t.Fatalf("telemetry endpoint mismatch: got=%s", cfg.Telemetry.Endpoint)
	}
}

func TestLoadConfigRejectsInvalidTelemetrySampleRatio(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	configPath := filepath.Join(root, "config.yaml")
	content := `
repo_root: .
server:
  listen: ":8080"
security:
  enable_api_key: true
  api_key: "test-key"
telemetry:
  enabled: true
  otlp_endpoint: "localhost:4318"
  insecure: true
  sample_ratio: 1.5
migration:
  migrations_dir: migrations
  allow_down: false
  allow_run_specific_file: false
  operation_timeout_seconds: 120
  database:
    server: "localhost"
    port: 1433
    username: "sa"
    password: "Password1!"
    database: "master"
    encrypt: "disable"
  db_lock:
    enabled: false
`
	if err := os.WriteFile(configPath, []byte(strings.TrimSpace(content)+"\n"), 0o644); err != nil {
		t.Fatalf("write config failed: %v", err)
	}

	_, err := LoadConfig(configPath)
	if err == nil {
		t.Fatalf("expected config load to fail for invalid telemetry.sample_ratio")
	}
}

func TestLoadConfigSupportsLoggingConfig(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	configPath := filepath.Join(root, "config.yaml")
	content := `
repo_root: .
server:
  listen: ":8080"
security:
  enable_api_key: true
  api_key: "test-key"
logging:
  level: "debug"
  format: "console"
  disable_caller: true
migration:
  migrations_dir: migrations
  allow_down: false
  allow_run_specific_file: false
  operation_timeout_seconds: 120
  database:
    server: "localhost"
    port: 1433
    username: "sa"
    password: "Password1!"
    database: "master"
    encrypt: "disable"
  db_lock:
    enabled: false
`
	if err := os.WriteFile(configPath, []byte(strings.TrimSpace(content)+"\n"), 0o644); err != nil {
		t.Fatalf("write config failed: %v", err)
	}

	cfg, err := LoadConfig(configPath)
	if err != nil {
		t.Fatalf("load config failed: %v", err)
	}
	if cfg.Logging.Level != "debug" {
		t.Fatalf("logging.level mismatch: got=%s", cfg.Logging.Level)
	}
	if cfg.Logging.Format != "console" {
		t.Fatalf("logging.format mismatch: got=%s", cfg.Logging.Format)
	}
	if !cfg.Logging.DisableCaller {
		t.Fatalf("logging.disable_caller mismatch: got=%v", cfg.Logging.DisableCaller)
	}
}

func TestLoadConfigRejectsInvalidLoggingConfig(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	configPath := filepath.Join(root, "config.yaml")
	content := `
repo_root: .
server:
  listen: ":8080"
security:
  enable_api_key: true
  api_key: "test-key"
logging:
  level: "trace"
  format: "xml"
migration:
  migrations_dir: migrations
  allow_down: false
  allow_run_specific_file: false
  operation_timeout_seconds: 120
  database:
    server: "localhost"
    port: 1433
    username: "sa"
    password: "Password1!"
    database: "master"
    encrypt: "disable"
  db_lock:
    enabled: false
`
	if err := os.WriteFile(configPath, []byte(strings.TrimSpace(content)+"\n"), 0o644); err != nil {
		t.Fatalf("write config failed: %v", err)
	}

	_, err := LoadConfig(configPath)
	if err == nil {
		t.Fatalf("expected config load to fail for invalid logging config")
	}
}
