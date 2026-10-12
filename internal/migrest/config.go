package migrest

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	defaultMigrationsDir = "migrations"
	defaultServiceName   = "migration_rest_service"
	placeholderDBURL     = "sqlserver://username:password@localhost:1433?database=YourDatabase&encrypt=disable"
	placeholderAPIKey    = "CHANGE_ME"
)

// Config is runtime config for migration REST service.
type Config struct {
	RepoRoot    string          `yaml:"repo_root"`
	ServiceName string          `yaml:"service_name"`
	Server      ServerConfig    `yaml:"server"`
	Security    SecurityConfig  `yaml:"security"`
	Logging     LoggingConfig   `yaml:"logging"`
	Telemetry   TelemetryConfig `yaml:"telemetry"`
	Migration   MigrationConfig `yaml:"migration"`
}

// ServerConfig controls HTTP listener/runtime timeouts.
type ServerConfig struct {
	Listen                 string `yaml:"listen"`
	ReadTimeoutSeconds     int    `yaml:"read_timeout_seconds"`
	WriteTimeoutSeconds    int    `yaml:"write_timeout_seconds"`
	IdleTimeoutSeconds     int    `yaml:"idle_timeout_seconds"`
	ShutdownTimeoutSeconds int    `yaml:"shutdown_timeout_seconds"`
}

// SecurityConfig controls API key auth for migration endpoints.
type SecurityConfig struct {
	EnableAPIKey bool   `yaml:"enable_api_key"`
	APIKey       string `yaml:"api_key"`
}

// LoggingConfig controls structured logger output.
type LoggingConfig struct {
	Level         string `yaml:"level"`
	Format        string `yaml:"format"`
	DisableCaller bool   `yaml:"disable_caller"`
}

// TelemetryConfig controls OpenTelemetry tracing export.
type TelemetryConfig struct {
	Enabled     bool    `yaml:"enabled"`
	Endpoint    string  `yaml:"otlp_endpoint"`
	Insecure    bool    `yaml:"insecure"`
	SampleRatio float64 `yaml:"sample_ratio"`
}

// MigrationConfig controls migration source and DB runtime.
type MigrationConfig struct {
	DatabaseURL             string       `yaml:"database_url"`
	Database                DatabaseDSN  `yaml:"database"`
	MigrationsDir           string       `yaml:"migrations_dir"`
	AllowDown               bool         `yaml:"allow_down"`
	AllowRunSpecificFile    bool         `yaml:"allow_run_specific_file"`
	AllowAcknowledgeManual  bool         `yaml:"allow_ack_manual"`
	OperationTimeoutSeconds int          `yaml:"operation_timeout_seconds"`
	DBLock                  DBLockConfig `yaml:"db_lock"`
}

// DatabaseDSN allows cleaner split config for MSSQL connection.
type DatabaseDSN struct {
	Server   string `yaml:"server"`
	Port     int    `yaml:"port"`
	Instance string `yaml:"instance"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Encrypt  string `yaml:"encrypt"`
}

// DBLockConfig controls optional lock acquisition via dbo.usp_mig_try_get_lock.
type DBLockConfig struct {
	Enabled   bool   `yaml:"enabled"`
	Resource  string `yaml:"resource"`
	TimeoutMS int    `yaml:"timeout_ms"`
}

// DefaultConfig returns safe baseline defaults.
func DefaultConfig() Config {
	return Config{
		RepoRoot:    ".",
		ServiceName: defaultServiceName,
		Server: ServerConfig{
			Listen:                 ":8080",
			ReadTimeoutSeconds:     15,
			WriteTimeoutSeconds:    120,
			IdleTimeoutSeconds:     120,
			ShutdownTimeoutSeconds: 10,
		},
		Security: SecurityConfig{
			EnableAPIKey: true,
			APIKey:       placeholderAPIKey,
		},
		Logging: LoggingConfig{
			Level:         "info",
			Format:        "json",
			DisableCaller: false,
		},
		Telemetry: TelemetryConfig{
			Enabled:     false,
			Endpoint:    "localhost:4318",
			Insecure:    true,
			SampleRatio: 1.0,
		},
		Migration: MigrationConfig{
			DatabaseURL:             "",
			Database:                defaultDatabaseDSN(),
			MigrationsDir:           defaultMigrationsDir,
			AllowDown:               false,
			AllowRunSpecificFile:    false,
			AllowAcknowledgeManual:  false,
			OperationTimeoutSeconds: 300,
			DBLock: DBLockConfig{
				Enabled:   true,
				Resource:  "rza_migrate_db",
				TimeoutMS: 0, // inherit from operation_timeout_seconds
			},
		},
	}
}

// LoadConfig loads and validates YAML config from disk.
func LoadConfig(path string) (Config, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	cfg := DefaultConfig()
	dec := yaml.NewDecoder(bytes.NewReader(raw))
	dec.KnownFields(true)
	if err := dec.Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("parse config yaml %s: %w", path, err)
	}
	cfg.Normalize()
	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}
	return cfg.ResolvePaths(path), nil
}

// Normalize applies default values for empty fields.
func (c *Config) Normalize() {
	if c == nil {
		return
	}
	c.RepoRoot = strings.TrimSpace(c.RepoRoot)
	if c.RepoRoot == "" {
		c.RepoRoot = "."
	}
	c.RepoRoot = filepath.Clean(c.RepoRoot)
	c.ServiceName = strings.TrimSpace(c.ServiceName)
	if c.ServiceName == "" {
		c.ServiceName = defaultServiceName
	}

	c.Server.Listen = strings.TrimSpace(c.Server.Listen)
	if c.Server.Listen == "" {
		c.Server.Listen = ":8080"
	}
	if c.Server.ReadTimeoutSeconds <= 0 {
		c.Server.ReadTimeoutSeconds = 15
	}
	if c.Server.WriteTimeoutSeconds <= 0 {
		c.Server.WriteTimeoutSeconds = 120
	}
	if c.Server.IdleTimeoutSeconds <= 0 {
		c.Server.IdleTimeoutSeconds = 120
	}
	if c.Server.ShutdownTimeoutSeconds <= 0 {
		c.Server.ShutdownTimeoutSeconds = 10
	}

	c.Migration.DatabaseURL = strings.TrimSpace(c.Migration.DatabaseURL)
	c.Migration.Database.Normalize()
	if c.Migration.DatabaseURL == "" || strings.EqualFold(c.Migration.DatabaseURL, placeholderDBURL) {
		if generated, err := c.Migration.Database.BuildURL(); err == nil {
			c.Migration.DatabaseURL = generated
		}
	}
	c.Migration.MigrationsDir = strings.TrimSpace(c.Migration.MigrationsDir)
	if c.Migration.MigrationsDir == "" {
		c.Migration.MigrationsDir = defaultMigrationsDir
	}
	c.Migration.MigrationsDir = filepath.Clean(c.Migration.MigrationsDir)
	if c.Migration.OperationTimeoutSeconds <= 0 {
		c.Migration.OperationTimeoutSeconds = 300
	}

	c.Migration.DBLock.Resource = strings.TrimSpace(c.Migration.DBLock.Resource)
	if c.Migration.DBLock.Resource == "" {
		c.Migration.DBLock.Resource = "rza_migrate_db"
	}
	if c.Migration.DBLock.TimeoutMS <= 0 {
		c.Migration.DBLock.TimeoutMS = c.Migration.OperationTimeoutSeconds * 1000
	}

	c.Security.APIKey = strings.TrimSpace(c.Security.APIKey)

	c.Logging.Level = strings.ToLower(strings.TrimSpace(c.Logging.Level))
	if c.Logging.Level == "" {
		c.Logging.Level = "info"
	}
	c.Logging.Format = strings.ToLower(strings.TrimSpace(c.Logging.Format))
	if c.Logging.Format == "" {
		c.Logging.Format = "json"
	}

	c.Telemetry.Endpoint = strings.TrimSpace(c.Telemetry.Endpoint)
	if c.Telemetry.Endpoint == "" {
		c.Telemetry.Endpoint = "localhost:4318"
	}
	if c.Telemetry.SampleRatio <= 0 {
		c.Telemetry.SampleRatio = 1.0
	}
}

// ResolvePaths resolves relative repo/migrations path from config file location.
func (c Config) ResolvePaths(configPath string) Config {
	configDir := filepath.Dir(filepath.Clean(configPath))

	repoRoot := c.RepoRoot
	if !filepath.IsAbs(repoRoot) {
		repoRoot = filepath.Join(configDir, repoRoot)
	}
	repoRoot = filepath.Clean(repoRoot)

	migrationsDir := c.Migration.MigrationsDir
	if !filepath.IsAbs(migrationsDir) {
		migrationsDir = filepath.Join(repoRoot, migrationsDir)
	}
	migrationsDir = filepath.Clean(migrationsDir)

	c.RepoRoot = repoRoot
	c.Migration.MigrationsDir = migrationsDir
	return c
}

// Validate validates config values with security-first defaults.
func (c Config) Validate() error {
	var issues []string

	if strings.TrimSpace(c.RepoRoot) == "" {
		issues = append(issues, "repo_root is required")
	}
	if strings.TrimSpace(c.ServiceName) == "" {
		issues = append(issues, "service_name is required")
	}
	if strings.TrimSpace(c.Server.Listen) == "" {
		issues = append(issues, "server.listen is required")
	}
	if c.Server.ReadTimeoutSeconds <= 0 {
		issues = append(issues, "server.read_timeout_seconds must be > 0")
	}
	if c.Server.WriteTimeoutSeconds <= 0 {
		issues = append(issues, "server.write_timeout_seconds must be > 0")
	}
	if c.Server.IdleTimeoutSeconds <= 0 {
		issues = append(issues, "server.idle_timeout_seconds must be > 0")
	}
	if c.Server.ShutdownTimeoutSeconds <= 0 {
		issues = append(issues, "server.shutdown_timeout_seconds must be > 0")
	}
	if strings.TrimSpace(c.Migration.DatabaseURL) == "" || strings.EqualFold(strings.TrimSpace(c.Migration.DatabaseURL), placeholderDBURL) {
		if _, err := c.Migration.Database.BuildURL(); err != nil {
			issues = append(issues, "provide migration.database_url or complete migration.database (server/username/password/database)")
			issues = append(issues, err.Error())
		}
	}
	if strings.TrimSpace(c.Migration.MigrationsDir) == "" {
		issues = append(issues, "migration.migrations_dir is required")
	}
	if c.Migration.OperationTimeoutSeconds <= 0 {
		issues = append(issues, "migration.operation_timeout_seconds must be > 0")
	}
	if c.Migration.DBLock.Enabled {
		if strings.TrimSpace(c.Migration.DBLock.Resource) == "" {
			issues = append(issues, "migration.db_lock.resource is required when lock is enabled")
		}
		if c.Migration.DBLock.TimeoutMS <= 0 {
			issues = append(issues, "migration.db_lock.timeout_ms must be > 0 when lock is enabled (or set 0 to inherit operation_timeout_seconds)")
		}
	}
	if c.Security.EnableAPIKey {
		if strings.TrimSpace(c.Security.APIKey) == "" {
			issues = append(issues, "security.api_key is required when security.enable_api_key=true")
		}
		if strings.EqualFold(strings.TrimSpace(c.Security.APIKey), placeholderAPIKey) {
			issues = append(issues, "security.api_key must be replaced from default placeholder")
		}
	}
	switch c.Logging.Level {
	case "debug", "info", "warn", "error":
	default:
		issues = append(issues, "logging.level must be one of: debug|info|warn|error")
	}
	switch c.Logging.Format {
	case "json", "console":
	default:
		issues = append(issues, "logging.format must be one of: json|console")
	}
	if c.Telemetry.SampleRatio < 0 || c.Telemetry.SampleRatio > 1 {
		issues = append(issues, "telemetry.sample_ratio must be between 0 and 1")
	}
	if c.Telemetry.Enabled {
		if strings.TrimSpace(c.Telemetry.Endpoint) == "" {
			issues = append(issues, "telemetry.otlp_endpoint is required when telemetry.enabled=true")
		}
	}

	if len(issues) > 0 {
		return errors.New(strings.Join(issues, "; "))
	}
	return nil
}

func defaultDatabaseDSN() DatabaseDSN {
	return DatabaseDSN{
		Server:   "localhost",
		Port:     1433,
		Instance: "",
		Username: "username",
		Password: "password",
		Database: "YourDatabase",
		Encrypt:  "disable",
	}
}

// Normalize applies defaults for split DB config.
func (d *DatabaseDSN) Normalize() {
	if d == nil {
		return
	}
	d.Server = strings.TrimSpace(d.Server)
	d.Instance = strings.TrimSpace(d.Instance)
	d.Username = strings.TrimSpace(d.Username)
	d.Password = strings.TrimSpace(d.Password)
	d.Database = strings.TrimSpace(d.Database)
	d.Encrypt = strings.TrimSpace(d.Encrypt)
	if d.Port <= 0 && d.Instance == "" {
		d.Port = 1433
	}
	if d.Encrypt == "" {
		d.Encrypt = "disable"
	}
}

// BuildURL builds SQL Server URL from split config.
func (d DatabaseDSN) BuildURL() (string, error) {
	d.Normalize()

	var missing []string
	if d.Server == "" {
		missing = append(missing, "migration.database.server")
	}
	if d.Username == "" {
		missing = append(missing, "migration.database.username")
	}
	if d.Password == "" {
		missing = append(missing, "migration.database.password")
	}
	if d.Database == "" {
		missing = append(missing, "migration.database.database")
	}
	if len(missing) > 0 {
		return "", fmt.Errorf("missing required fields: %s", strings.Join(missing, ", "))
	}
	if strings.EqualFold(d.Username, "username") && strings.EqualFold(d.Password, "password") && strings.EqualFold(d.Database, "yourdatabase") {
		return "", errors.New("migration.database must replace default placeholder username/password/database values")
	}

	// Support SQL Server named instance from:
	// - explicit `instance` field
	// - `server` value in form `host\\instance`
	if strings.Contains(d.Server, "\\") {
		parts := strings.SplitN(d.Server, "\\", 2)
		serverHost := strings.TrimSpace(parts[0])
		serverInstance := strings.TrimSpace(parts[1])
		if serverHost == "" || serverInstance == "" {
			return "", errors.New("migration.database.server host\\\\instance format is invalid")
		}
		if d.Instance != "" && !strings.EqualFold(d.Instance, serverInstance) {
			return "", errors.New("migration.database.instance conflicts with server host\\\\instance")
		}
		d.Server = serverHost
		d.Instance = serverInstance
	}

	host := d.Server
	if host == "" {
		return "", errors.New("migration.database.server is required")
	}
	if d.Instance == "" && d.Port > 0 && !strings.Contains(host, ":") {
		host = host + ":" + strconv.Itoa(d.Port)
	}

	query := url.Values{}
	query.Set("database", d.Database)
	query.Set("encrypt", d.Encrypt)
	if d.Instance != "" {
		query.Set("instance", d.Instance)
	}

	uri := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(d.Username, d.Password),
		Host:     host,
		RawQuery: query.Encode(),
	}
	return uri.String(), nil
}
