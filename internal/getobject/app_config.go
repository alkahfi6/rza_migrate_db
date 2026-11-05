package getobject

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// AppConfig controls runtime settings for rza_get_object_db.
type AppConfig struct {
	ConnectionString       string `json:"connection_string" yaml:"connection_string,omitempty"`
	RepoRoot               string `json:"repo_root" yaml:"repo_root"`
	DbObjectDir            string `json:"db_object_dir" yaml:"db_object_dir"`
	Schema                 string `json:"schema" yaml:"schema"`
	ConnectTimeoutSeconds  int    `json:"connect_timeout_seconds" yaml:"connect_timeout_seconds"`
	QueryTimeoutSeconds    int    `json:"query_timeout_seconds" yaml:"query_timeout_seconds"`
	RenderWorkers          int    `json:"render_workers" yaml:"render_workers"`
	WriteWorkers           int    `json:"write_workers" yaml:"write_workers"`
	SkipUnsupportedTables  bool   `json:"skip_unsupported_tables" yaml:"skip_unsupported_tables"`
	StripModuleComments    bool   `json:"strip_module_comments" yaml:"strip_module_comments"`
	ReindentModuleSQL      bool   `json:"reindent_module_sql" yaml:"reindent_module_sql"`
	IndentSize             int    `json:"indent_size" yaml:"indent_size"`
	ModuleFormatterBackend string `json:"module_formatter_backend" yaml:"module_formatter_backend"`

	// initialMigrationObjects is populated at export time from migrations/*_v1.up.sql.
	// It is intentionally runtime-only and is not part of the JSON config contract.
	initialMigrationObjects map[string]struct{}
}

// DefaultAppConfig returns a safe baseline config.
func DefaultAppConfig(repoRoot string) AppConfig {
	return AppConfig{
		ConnectionString:       "sqlserver://username:password@localhost:1433?database=YourDatabase&encrypt=disable",
		RepoRoot:               filepath.Clean(repoRoot),
		DbObjectDir:            "db_object",
		Schema:                 "dbo",
		ConnectTimeoutSeconds:  15,
		QueryTimeoutSeconds:    120,
		RenderWorkers:          4,
		WriteWorkers:           4,
		SkipUnsupportedTables:  true,
		StripModuleComments:    true,
		ReindentModuleSQL:      true,
		IndentSize:             2,
		ModuleFormatterBackend: "lite",
	}
}

// LoadAppConfig loads YAML or legacy JSON config and resolves database credentials
// into the DSN consumed by the exporter. Decoder errors never echo secret values.
func LoadAppConfig(path string) (AppConfig, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return AppConfig{}, err
	}

	file := struct {
		AppConfig `yaml:",inline"`
		Database  *DatabaseConfig `json:"database" yaml:"database"`
	}{AppConfig: DefaultAppConfig(".")}
	// A file must explicitly supply one credential source; do not confuse the
	// compiled placeholder with a connection_string actually present in the file.
	file.ConnectionString = ""
	if strings.EqualFold(filepath.Ext(path), ".json") {
		if err := json.Unmarshal(raw, &file); err != nil {
			return AppConfig{}, fmt.Errorf("parse app config %s: invalid JSON; check field names and value types", path)
		}
	} else {
		decoder := yaml.NewDecoder(strings.NewReader(string(raw)))
		decoder.KnownFields(true)
		if err := decoder.Decode(&file); err != nil {
			return AppConfig{}, fmt.Errorf("parse app config %s: invalid YAML; check field names, duplicates and value types", path)
		}
		var extra any
		if err := decoder.Decode(&extra); err != io.EOF {
			return AppConfig{}, fmt.Errorf("parse app config %s: expected one YAML document", path)
		}
	}
	cfg := file.AppConfig
	if file.Database != nil {
		if cfg.ConnectionString != "" {
			return AppConfig{}, errors.New("use either database or connection_string, not both")
		}
		cfg.ConnectionString, err = file.Database.ConnectionString()
		if err != nil {
			return AppConfig{}, err
		}
	}
	cfg.Normalize()
	return cfg, cfg.Validate()
}

// WriteAppConfigTemplate writes a starter config file.
func WriteAppConfigTemplate(path string, cfg AppConfig) error {
	cfg.Normalize()
	var raw []byte
	var err error
	if strings.EqualFold(filepath.Ext(path), ".json") {
		raw, err = json.MarshalIndent(cfg, "", "  ")
	} else {
		raw, err = yaml.Marshal(cfg)
	}
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, append(raw, '\n'), 0o600)
}

// Normalize applies empty-value defaults.
func (c *AppConfig) Normalize() {
	if c == nil {
		return
	}
	c.RepoRoot = strings.TrimSpace(c.RepoRoot)
	if c.RepoRoot == "" {
		c.RepoRoot = "."
	}
	c.RepoRoot = filepath.Clean(c.RepoRoot)

	c.DbObjectDir = strings.TrimSpace(c.DbObjectDir)
	if c.DbObjectDir == "" {
		c.DbObjectDir = "db_object"
	}
	c.DbObjectDir = filepath.Clean(c.DbObjectDir)

	c.Schema = strings.TrimSpace(c.Schema)
	if c.Schema == "" {
		c.Schema = "dbo"
	}

	if c.ConnectTimeoutSeconds <= 0 {
		c.ConnectTimeoutSeconds = 15
	}
	if c.QueryTimeoutSeconds <= 0 {
		c.QueryTimeoutSeconds = 120
	}
	if c.RenderWorkers <= 0 {
		c.RenderWorkers = 4
	}
	if c.WriteWorkers <= 0 {
		c.WriteWorkers = 4
	}
	if c.IndentSize <= 0 {
		c.IndentSize = 2
	}
	c.ModuleFormatterBackend = strings.ToLower(strings.TrimSpace(c.ModuleFormatterBackend))
	if c.ModuleFormatterBackend == "" {
		c.ModuleFormatterBackend = "lite"
	}
}

// Validate validates config against rules contract.
func (c AppConfig) Validate() error {
	var issues []string

	if strings.TrimSpace(c.ConnectionString) == "" {
		issues = append(issues, "connection_string is required")
	}
	if strings.EqualFold(strings.TrimSpace(c.ConnectionString), "sqlserver://username:password@localhost:1433?database=YourDatabase&encrypt=disable") {
		issues = append(issues, "connection_string must be replaced with actual MSSQL DSN")
	}
	if strings.TrimSpace(c.DbObjectDir) == "" {
		issues = append(issues, "db_object_dir is required")
	}
	if strings.TrimSpace(c.Schema) != "dbo" {
		issues = append(issues, "schema must be dbo to satisfy rules.md")
	}
	if c.ConnectTimeoutSeconds <= 0 {
		issues = append(issues, "connect_timeout_seconds must be > 0")
	}
	if c.QueryTimeoutSeconds <= 0 {
		issues = append(issues, "query_timeout_seconds must be > 0")
	}
	if c.RenderWorkers <= 0 {
		issues = append(issues, "render_workers must be > 0")
	}
	if c.WriteWorkers <= 0 {
		issues = append(issues, "write_workers must be > 0")
	}
	if c.IndentSize <= 0 {
		issues = append(issues, "indent_size must be > 0")
	}
	switch c.ModuleFormatterBackend {
	case "antlr", "lite":
	default:
		issues = append(issues, "module_formatter_backend must be one of: antlr|lite")
	}

	if len(issues) > 0 {
		return errors.New(strings.Join(issues, "; "))
	}
	return nil
}

// RepoPath resolves relative path against repo root.
func (c AppConfig) RepoPath(rel string) string {
	return filepath.Join(c.RepoRoot, rel)
}

// SchemaTreeDir returns db_object root (kept for backward compatibility).
func (c AppConfig) SchemaTreeDir() string {
	return c.DbObjectDir
}
