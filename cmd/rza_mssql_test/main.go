package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlserver"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/microsoft/go-mssqldb"
)

const defaultCredentialFile = "mssql_credential.nfm"

type credentialConfig struct {
	DatabaseType string
	Host         string
	Username     string
	Password     string
	Database     string
}

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	command := "status"
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		command = strings.ToLower(strings.TrimSpace(args[0]))
		args = args[1:]
	}

	fs := flag.NewFlagSet("rza_mssql_test", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	configPath := fs.String("config", defaultCredentialFile, "credential file in key=value format")
	migrationsDir := fs.String("migrations", "migrations", "migration directory")
	steps := fs.Int("steps", 0, "number of migration steps; up defaults to all, down defaults to 1")
	all := fs.Bool("all", false, "run all available steps (down requires -allow-destructive)")
	timeout := fs.Duration("timeout", 5*time.Minute, "database operation timeout")
	allowDestructive := fs.Bool("allow-destructive", false, "required for down --all and reset")
	confirmReset := fs.String("confirm-reset", "", "must exactly match the target database for reset")
	dryRun := fs.Bool("dry-run", false, "show reset target counts without changing the database")
	if err := fs.Parse(args); err != nil {
		return 2
	}

	if *timeout <= 0 {
		fmt.Fprintln(os.Stderr, "error: -timeout must be greater than zero")
		return 2
	}
	if command != "status" && command != "up" && command != "down" && command != "reset" {
		printUsage()
		return 2
	}
	if *dryRun && command != "reset" {
		fmt.Fprintln(os.Stderr, "error: -dry-run is only supported by reset")
		return 2
	}

	cfg, err := loadCredentialConfig(*configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: load credential config: %v\n", err)
		return 1
	}
	dsn, err := cfg.databaseURL()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: build database URL: %v\n", err)
		return 1
	}

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: open SQL Server connection: %v\n", err)
		return 1
	}
	defer db.Close()
	if err := db.PingContext(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error: ping SQL Server: %v\n", err)
		return 1
	}

	actualDatabase, err := currentDatabase(ctx, db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: read target database: %v\n", err)
		return 1
	}
	if !strings.EqualFold(actualDatabase, cfg.Database) {
		fmt.Fprintf(os.Stderr, "error: connected database %q does not match credential initial_catalog %q\n", actualDatabase, cfg.Database)
		return 1
	}

	switch command {
	case "status":
		return runStatus(ctx, dsn, *migrationsDir, actualDatabase)
	case "up":
		return runMigration(dsn, *migrationsDir, *steps, false, *all, *allowDestructive)
	case "down":
		return runMigration(dsn, *migrationsDir, *steps, true, *all, *allowDestructive)
	case "reset":
		return runReset(ctx, db, actualDatabase, *allowDestructive, *confirmReset, *dryRun)
	default:
		return 2
	}
}

func loadCredentialConfig(path string) (credentialConfig, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return credentialConfig{}, err
	}

	values := make(map[string]string)
	for lineNumber, rawLine := range strings.Split(string(raw), "\n") {
		line := strings.TrimSpace(strings.TrimSuffix(rawLine, "\r"))
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}
		separator := strings.IndexAny(line, "=:")
		if separator <= 0 {
			return credentialConfig{}, fmt.Errorf("invalid config line %d", lineNumber+1)
		}
		key := strings.ToLower(strings.TrimSpace(line[:separator]))
		value := strings.TrimSpace(line[separator+1:])
		value = strings.Trim(value, "\"'")
		values[key] = value
	}

	cfg := credentialConfig{
		DatabaseType: strings.ToLower(values["database_type"]),
		Host:         values["host"],
		Username:     values["username"],
		Password:     values["password"],
		Database:     values["initial_catalog"],
	}
	if cfg.DatabaseType == "" {
		cfg.DatabaseType = "mssql"
	}
	if cfg.DatabaseType != "mssql" && cfg.DatabaseType != "sqlserver" {
		return credentialConfig{}, fmt.Errorf("database_type must be mssql or sqlserver")
	}
	for key, value := range map[string]string{
		"host":            cfg.Host,
		"username":        cfg.Username,
		"password":        cfg.Password,
		"initial_catalog": cfg.Database,
	} {
		if strings.TrimSpace(value) == "" {
			return credentialConfig{}, fmt.Errorf("%s is required", key)
		}
	}
	return cfg, nil
}

func (c credentialConfig) databaseURL() (string, error) {
	host := strings.TrimSpace(c.Host)
	if strings.Contains(host, "://") || strings.ContainsAny(host, "/?#") {
		return "", errors.New("host must be a host name, host:port, or host\\instance")
	}
	instance := ""
	if strings.Contains(host, "\\") {
		parts := strings.Split(host, "\\")
		if len(parts) != 2 || strings.TrimSpace(parts[0]) == "" || strings.TrimSpace(parts[1]) == "" {
			return "", errors.New("host\\instance format is invalid")
		}
		host = strings.TrimSpace(parts[0])
		instance = strings.TrimSpace(parts[1])
	}
	query := url.Values{}
	query.Set("database", c.Database)
	query.Set("encrypt", "disable")
	if instance != "" {
		query.Set("instance", instance)
	}
	return (&url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(c.Username, c.Password),
		Host:     host,
		RawQuery: query.Encode(),
	}).String(), nil
}

func runStatus(ctx context.Context, dsn, migrationsDir, database string) int {
	m, err := newMigrator(migrationsDir, dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: create migrator: %v\n", err)
		return 1
	}
	defer closeMigrator(m)

	version, dirty, err := m.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		fmt.Fprintf(os.Stderr, "error: read migration version: %v\n", err)
		return 1
	}
	if errors.Is(err, migrate.ErrNilVersion) {
		fmt.Printf("[status] database=%s version=<none> dirty=false\n", database)
		return 0
	}
	fmt.Printf("[status] database=%s version=%d dirty=%t\n", database, version, dirty)
	return 0
}

func runMigration(dsn, migrationsDir string, steps int, down, all, allowDestructive bool) int {
	if down {
		if all && !allowDestructive {
			fmt.Fprintln(os.Stderr, "error: down --all requires -allow-destructive")
			return 2
		}
		if !all && steps <= 0 {
			steps = 1
		}
	} else if !all && steps < 0 {
		fmt.Fprintln(os.Stderr, "error: -steps cannot be negative")
		return 2
	}

	m, err := newMigrator(migrationsDir, dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: create migrator: %v\n", err)
		return 1
	}
	defer closeMigrator(m)

	if down {
		if all {
			err = m.Down()
		} else {
			err = m.Steps(-steps)
		}
	} else if steps == 0 || all {
		err = m.Up()
	} else {
		err = m.Steps(steps)
	}
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		fmt.Fprintf(os.Stderr, "error: run migration: %v\n", err)
		return 1
	}
	if errors.Is(err, migrate.ErrNoChange) {
		fmt.Println("[migration] no change")
		return 0
	}
	direction := "up"
	if down {
		direction = "down"
	}
	fmt.Printf("[migration] direction=%s completed\n", direction)
	return 0
}

func runReset(ctx context.Context, db *sql.DB, database string, allowDestructive bool, confirmation string, dryRun bool) int {
	if !allowDestructive {
		fmt.Fprintln(os.Stderr, "error: reset requires -allow-destructive")
		return 2
	}
	if confirmation != database {
		fmt.Fprintf(os.Stderr, "error: reset requires -confirm-reset %q\n", database)
		return 2
	}

	count, err := countResetObjects(ctx, db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: count dbo objects: %v\n", err)
		return 1
	}
	if dryRun {
		fmt.Printf("[reset] database=%s dbo_objects=%d dry_run=true\n", database, count)
		return 0
	}
	if err := resetDBOSchema(ctx, db); err != nil {
		fmt.Fprintf(os.Stderr, "error: reset dbo schema: %v\n", err)
		return 1
	}
	remaining, err := countResetObjects(ctx, db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: verify dbo reset: %v\n", err)
		return 1
	}
	if remaining != 0 {
		fmt.Fprintf(os.Stderr, "error: reset incomplete; %d user object(s) remain in dbo\n", remaining)
		return 1
	}
	fmt.Printf("[reset] database=%s removed_objects=%d remaining=0\n", database, count)
	return 0
}

func currentDatabase(ctx context.Context, db *sql.DB) (string, error) {
	var database string
	if err := db.QueryRowContext(ctx, "SELECT DB_NAME()").Scan(&database); err != nil {
		return "", err
	}
	return database, nil
}

func countResetObjects(ctx context.Context, db *sql.DB) (int, error) {
	const query = `SELECT
  (SELECT COUNT(*)
   FROM sys.objects o
   JOIN sys.schemas s ON s.schema_id = o.schema_id
   WHERE s.name = N'dbo' AND o.is_ms_shipped = 0
     AND o.type IN ('U', 'V', 'P', 'PC', 'FN', 'IF', 'TF', 'FS', 'FT', 'SN', 'SO'))
  +
  (SELECT COUNT(*)
   FROM sys.types t
   JOIN sys.schemas s ON s.schema_id = t.schema_id
   WHERE s.name = N'dbo' AND t.is_user_defined = 1)`
	var count int
	err := db.QueryRowContext(ctx, query).Scan(&count)
	return count, err
}

func resetDBOSchema(ctx context.Context, db *sql.DB) error {
	// Drop schema-bound module dependents iteratively before removing tables.
	// SQL Server rejects a parent function/view while another module still binds to it.
	const resetSQL = `SET NOCOUNT ON;
DECLARE @sql NVARCHAR(MAX) = N'';

SELECT @sql += N'ALTER TABLE ' + QUOTENAME(SCHEMA_NAME(parent.schema_id)) + N'.' + QUOTENAME(parent.name) + N' DROP CONSTRAINT ' + QUOTENAME(fk.name) + N';' + CHAR(10)
FROM sys.foreign_keys fk
JOIN sys.objects parent ON parent.object_id = fk.parent_object_id
JOIN sys.objects referenced ON referenced.object_id = fk.referenced_object_id
WHERE SCHEMA_NAME(parent.schema_id) = N'dbo' OR SCHEMA_NAME(referenced.schema_id) = N'dbo';
IF @sql <> N'' EXEC sys.sp_executesql @sql;

DECLARE @pass INT = 0;
DECLARE @dropped INT;
WHILE @pass < 2000
BEGIN
  SET @dropped = 0;
  SET @sql = N'';
  SELECT @sql += N'BEGIN TRY DROP ' +
    CASE o.type
      WHEN 'V' THEN N'VIEW '
      WHEN 'P' THEN N'PROCEDURE '
      WHEN 'PC' THEN N'PROCEDURE '
      ELSE N'FUNCTION '
    END +
    QUOTENAME(s.name) + N'.' + QUOTENAME(o.name) +
    N'; SET @dropped = @dropped + 1; END TRY BEGIN CATCH END CATCH;' + CHAR(10)
  FROM sys.objects o
  JOIN sys.schemas s ON s.schema_id = o.schema_id
  WHERE s.name = N'dbo' AND o.is_ms_shipped = 0
    AND o.type IN ('V', 'P', 'PC', 'FN', 'IF', 'TF', 'FS', 'FT');

  IF @sql = N'' BREAK;
  EXEC sys.sp_executesql @sql, N'@dropped INT OUTPUT', @dropped = @dropped OUTPUT;
  IF @dropped = 0 BREAK;
  SET @pass = @pass + 1;
END;

IF EXISTS (
  SELECT 1
  FROM sys.objects o
  JOIN sys.schemas s ON s.schema_id = o.schema_id
  WHERE s.name = N'dbo' AND o.is_ms_shipped = 0
    AND o.type IN ('V', 'P', 'PC', 'FN', 'IF', 'TF', 'FS', 'FT')
)
  THROW 50000, 'reset could not drop all dbo modules because an external schema-bound dependency remains', 1;

SET @sql = N'';
SELECT @sql += N'DROP SYNONYM ' + QUOTENAME(s.name) + N'.' + QUOTENAME(sn.name) + N';' + CHAR(10)
FROM sys.synonyms sn JOIN sys.schemas s ON s.schema_id = sn.schema_id
WHERE s.name = N'dbo';
IF @sql <> N'' EXEC sys.sp_executesql @sql;

SET @sql = N'';
SELECT @sql += N'DROP TABLE ' + QUOTENAME(s.name) + N'.' + QUOTENAME(o.name) + N';' + CHAR(10)
FROM sys.objects o JOIN sys.schemas s ON s.schema_id = o.schema_id
WHERE s.name = N'dbo' AND o.is_ms_shipped = 0 AND o.type = 'U';
IF @sql <> N'' EXEC sys.sp_executesql @sql;

SET @sql = N'';
SELECT @sql += N'DROP SEQUENCE ' + QUOTENAME(s.name) + N'.' + QUOTENAME(seq.name) + N';' + CHAR(10)
FROM sys.sequences seq JOIN sys.schemas s ON s.schema_id = seq.schema_id
WHERE s.name = N'dbo' AND seq.is_ms_shipped = 0;
IF @sql <> N'' EXEC sys.sp_executesql @sql;

SET @sql = N'';
SELECT @sql += N'DROP TYPE ' + QUOTENAME(s.name) + N'.' + QUOTENAME(t.name) + N';' + CHAR(10)
FROM sys.types t JOIN sys.schemas s ON s.schema_id = t.schema_id
WHERE s.name = N'dbo' AND t.is_user_defined = 1;
IF @sql <> N'' EXEC sys.sp_executesql @sql;`
	_, err := db.ExecContext(ctx, resetSQL)
	return err
}

func newMigrator(migrationsDir, dsn string) (*migrate.Migrate, error) {
	abs, err := filepath.Abs(migrationsDir)
	if err != nil {
		return nil, err
	}
	info, err := os.Stat(abs)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("migrations path is not a directory: %s", abs)
	}
	return migrate.New(fileSourceURL(abs), dsn)
}

func closeMigrator(m *migrate.Migrate) {
	if m == nil {
		return
	}
	_, _ = m.Close()
}

func fileSourceURL(abs string) string {
	path := filepath.ToSlash(abs)
	if runtime.GOOS == "windows" && len(path) >= 2 && path[1] == ':' {
		return "file://" + path
	}
	if strings.HasPrefix(path, "/") {
		return "file://" + path
	}
	return "file:///" + path
}

func printUsage() {
	fmt.Fprintln(os.Stderr, "Usage: rza_mssql_test [status|up|down|reset] [-config mssql_credential.nfm] [-migrations migrations] [-steps N] [-all] [-timeout 5m]")
	fmt.Fprintln(os.Stderr, "reset requires -allow-destructive -confirm-reset <initial_catalog>; use -dry-run to inspect first")
	fmt.Fprintln(os.Stderr, "down --all requires -allow-destructive; down without -steps runs one migration")
}
