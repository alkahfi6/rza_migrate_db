// rza_migrate_test runs repository migrations against an explicitly configured SQL Server test database.
package main

import (
	"bufio"
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"rza_migrate_db/internal/migrest"
)

const defaultOperationTimeout = 5 * time.Minute

type nfmCredentials struct {
	DatabaseType string
	Host         string
	Username     string
	Password     string
	Database     string
}

type resetObject struct {
	Kind   string
	Schema string
	Name   string
}

type resetForeignKey struct {
	Schema     string
	Table      string
	Constraint string
}

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	if len(args) < 2 {
		printUsage()
		return 1
	}

	command := strings.ToLower(strings.TrimSpace(args[1]))
	fs := flag.NewFlagSet(command, flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	repo := fs.String("repo", ".", "repository root containing migrations/")
	configPath := fs.String("config", "mssql_credential.nfm", "path to .nfm SQL Server credentials")
	databaseOverride := fs.String("database", "", "override initial_catalog from credentials (recommended for isolated real tests)")
	collation := fs.String("collation", "", "create-db only: validated SQL Server collation for the isolated catalog")
	steps := fs.Int("steps", 0, "number of migration steps; required for down")
	timeout := fs.Duration("timeout", defaultOperationTimeout, "database operation timeout")
	dryRun := fs.Bool("dry-run", false, "show reset objects without dropping them")
	confirmReset := fs.String("confirm-reset", "", "required for reset: exact target database name")
	catalogMap := fs.String("catalog-map", "", "up/down only: comma-separated source=test catalog mappings applied to a temporary migration copy")
	if err := fs.Parse(args[2:]); err != nil {
		return 1
	}
	if *timeout <= 0 {
		fmt.Fprintln(os.Stderr, "error: timeout must be greater than zero")
		return 1
	}

	absRepo, err := filepath.Abs(*repo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: resolve repo: %v\n", err)
		return 1
	}
	absConfig := *configPath
	if !filepath.IsAbs(absConfig) {
		absConfig = filepath.Join(absRepo, absConfig)
	}
	creds, err := loadNFMCredentials(absConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: load credentials: %v\n", err)
		return 1
	}
	if strings.TrimSpace(*databaseOverride) != "" {
		creds.Database = strings.TrimSpace(*databaseOverride)
	}
	if err := validateDatabaseName(creds.Database); err != nil {
		fmt.Fprintf(os.Stderr, "error: target database: %v\n", err)
		return 1
	}
	if command == "create-db" {
		return runCreateDatabase(context.Background(), creds, *confirmReset, strings.TrimSpace(*collation), *timeout)
	}
	dsn, err := creds.databaseURL()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: build database connection: %v\n", err)
		return 1
	}
	migrationsDir := filepath.Join(absRepo, "migrations")
	if info, statErr := os.Stat(migrationsDir); statErr != nil || !info.IsDir() {
		fmt.Fprintf(os.Stderr, "error: migrations directory not found: %s\n", migrationsDir)
		return 1
	}
	if strings.TrimSpace(*catalogMap) != "" {
		if command != "up" && command != "down" {
			fmt.Fprintln(os.Stderr, "error: catalog-map is allowed only for up/down")
			return 1
		}
		mappedDir, cleanup, mapErr := prepareCatalogMappedMigrations(migrationsDir, *catalogMap)
		if mapErr != nil {
			fmt.Fprintf(os.Stderr, "error: catalog-map: %v\n", mapErr)
			return 1
		}
		defer cleanup()
		migrationsDir = mappedDir
		fmt.Printf("[catalog-map] mappings=%s source_immutable=true\n", strings.TrimSpace(*catalogMap))
	}

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	switch command {
	case "status":
		return runStatus(ctx, dsn, creds.Database)
	case "up":
		return runUp(ctx, dsn, migrationsDir, *steps)
	case "down":
		if *steps <= 0 {
			fmt.Fprintln(os.Stderr, "error: down requires -steps greater than zero")
			return 1
		}
		return runDown(ctx, dsn, migrationsDir, *steps)
	case "reset":
		return runReset(ctx, dsn, creds.Database, *confirmReset, *dryRun)
	default:
		printUsage()
		return 1
	}
}

type catalogMapping struct {
	Source string
	Target string
	Regex  *regexp.Regexp
}

func prepareCatalogMappedMigrations(sourceDir, rawMappings string) (string, func(), error) {
	mappings, err := parseCatalogMappings(rawMappings)
	if err != nil {
		return "", func() {}, err
	}
	tempDir, err := os.MkdirTemp("", "rza-migrate-test-catalog-map-")
	if err != nil {
		return "", func() {}, err
	}
	cleanup := func() { _ = os.RemoveAll(tempDir) }
	entries, err := os.ReadDir(sourceDir)
	if err != nil {
		cleanup()
		return "", func() {}, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		raw, err := os.ReadFile(filepath.Join(sourceDir, entry.Name()))
		if err != nil {
			cleanup()
			return "", func() {}, err
		}
		if strings.EqualFold(filepath.Ext(entry.Name()), ".sql") {
			text := string(raw)
			for _, mapping := range mappings {
				text = mapping.Regex.ReplaceAllString(text, mapping.Target)
			}
			raw = []byte(text)
		}
		if err := os.WriteFile(filepath.Join(tempDir, entry.Name()), raw, 0o600); err != nil {
			cleanup()
			return "", func() {}, err
		}
	}
	return tempDir, cleanup, nil
}

func parseCatalogMappings(raw string) ([]catalogMapping, error) {
	parts := strings.Split(raw, ",")
	out := make([]catalogMapping, 0, len(parts))
	seen := map[string]bool{}
	for _, part := range parts {
		pair := strings.SplitN(strings.TrimSpace(part), "=", 2)
		if len(pair) != 2 {
			return nil, fmt.Errorf("invalid mapping %q (expected source=target)", part)
		}
		source, target := strings.TrimSpace(pair[0]), strings.TrimSpace(pair[1])
		if err := validateDatabaseName(source); err != nil {
			return nil, fmt.Errorf("source %q: %w", source, err)
		}
		if err := validateDatabaseName(target); err != nil {
			return nil, fmt.Errorf("target %q: %w", target, err)
		}
		if err := validateMappingCatalogName(source); err != nil {
			return nil, fmt.Errorf("source %q: %w", source, err)
		}
		if err := validateMappingCatalogName(target); err != nil {
			return nil, fmt.Errorf("target %q: %w", target, err)
		}
		if isSystemDatabase(target) {
			return nil, fmt.Errorf("target mapping to system database %q is forbidden", target)
		}
		if strings.EqualFold(source, target) {
			return nil, fmt.Errorf("source and target must differ for %q", source)
		}
		key := strings.ToLower(source)
		if seen[key] {
			return nil, fmt.Errorf("duplicate source mapping %q", source)
		}
		seen[key] = true
		pattern := `(?i)(^|[^A-Za-z0-9_])` + regexp.QuoteMeta(source) + `([^A-Za-z0-9_]|$)`
		re, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
		// Preserve the delimiter captures so SQL_REPLICATE, [SQL_REPLICATE],
		// and literals containing the exact catalog token all map consistently.
		replacement := `${1}` + target + `${2}`
		out = append(out, catalogMapping{Source: source, Target: replacement, Regex: re})
	}
	return out, nil
}

func validateMappingCatalogName(name string) error {
	for _, ch := range name {
		if ch == '_' || ch >= '0' && ch <= '9' || ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' {
			continue
		}
		return errors.New("catalog-map names may contain only letters, digits, and underscore")
	}
	return nil
}

func runCreateDatabase(parent context.Context, creds nfmCredentials, confirmation, collation string, timeout time.Duration) int {
	target := strings.TrimSpace(creds.Database)
	if isSystemDatabase(target) {
		fmt.Fprintf(os.Stderr, "error: create-db is forbidden for system database %q\n", target)
		return 1
	}
	if !strings.EqualFold(strings.TrimSpace(confirmation), target) {
		fmt.Fprintf(os.Stderr, "error: create-db requires -confirm-reset %s\n", target)
		return 1
	}
	masterCreds := creds
	masterCreds.Database = "master"
	dsn, err := masterCreds.databaseURL()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: build master connection: %v\n", err)
		return 1
	}
	ctx, cancel := context.WithTimeout(parent, timeout)
	defer cancel()
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: open SQL Server master: %v\n", err)
		return 1
	}
	defer db.Close()
	if err := db.PingContext(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error: connect SQL Server master: %v\n", err)
		return 1
	}
	var exists int
	if err := db.QueryRowContext(ctx, "SELECT CASE WHEN DB_ID(@database) IS NULL THEN 0 ELSE 1 END;", sql.Named("database", target)).Scan(&exists); err != nil {
		fmt.Fprintf(os.Stderr, "error: inspect target database: %v\n", err)
		return 1
	}
	if exists == 1 {
		if collation != "" {
			var actual string
			if err := db.QueryRowContext(ctx, "SELECT CONVERT(sysname, DATABASEPROPERTYEX(@database, 'Collation'));", sql.Named("database", target)).Scan(&actual); err != nil {
				fmt.Fprintf(os.Stderr, "error: inspect target collation: %v\n", err)
				return 1
			}
			if !strings.EqualFold(actual, collation) {
				fmt.Fprintf(os.Stderr, "error: database %s already exists with collation %s, requested %s\n", target, actual, collation)
				return 1
			}
		}
		fmt.Printf("[create-db] database=%s already_exists=true\n", target)
		return 0
	}
	createSQL := "CREATE DATABASE " + quoteIdentifier(target)
	if collation != "" {
		if err := validateCollation(ctx, db, collation); err != nil {
			fmt.Fprintf(os.Stderr, "error: collation: %v\n", err)
			return 1
		}
		createSQL += " COLLATE " + collation
	}
	if _, err := db.ExecContext(ctx, createSQL+";"); err != nil {
		fmt.Fprintf(os.Stderr, "error: create database %s: %v\n", target, err)
		return 1
	}
	fmt.Printf("[create-db] database=%s created=true\n", target)
	return 0
}

func validateCollation(ctx context.Context, db *sql.DB, collation string) error {
	if collation == "" {
		return nil
	}
	for _, r := range collation {
		if !(r == '_' || r >= '0' && r <= '9' || r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z') {
			return fmt.Errorf("invalid collation name %q", collation)
		}
	}
	var exists int
	if err := db.QueryRowContext(ctx, "SELECT CASE WHEN EXISTS (SELECT 1 FROM sys.fn_helpcollations() WHERE name = @collation) THEN 1 ELSE 0 END;", sql.Named("collation", collation)).Scan(&exists); err != nil {
		return fmt.Errorf("validate with SQL Server: %w", err)
	}
	if exists == 0 {
		return fmt.Errorf("SQL Server does not recognize %q", collation)
	}
	return nil
}

func runStatus(ctx context.Context, dsn, database string) int {
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: open SQL Server: %v\n", err)
		return 1
	}
	defer db.Close()
	if err := db.PingContext(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error: connect SQL Server: %v\n", err)
		return 1
	}

	var exists int
	if err := db.QueryRowContext(ctx, "SELECT CASE WHEN OBJECT_ID(N'dbo.schema_migrations', N'U') IS NULL THEN 0 ELSE 1 END;").Scan(&exists); err != nil {
		fmt.Fprintf(os.Stderr, "error: check migration status: %v\n", err)
		return 1
	}
	if exists == 0 {
		fmt.Printf("[status] database=%s version=<none> dirty=false\n", database)
		return 0
	}
	var version int64
	var dirty bool
	err = db.QueryRowContext(ctx, "SELECT TOP (1) version, dirty FROM dbo.schema_migrations;").Scan(&version, &dirty)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("[status] database=%s version=<none> dirty=false\n", database)
		return 0
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: read migration status: %v\n", err)
		return 1
	}
	fmt.Printf("[status] database=%s version=%d dirty=%t\n", database, version, dirty)
	return 0
}

func runUp(ctx context.Context, dsn, migrationsDir string, steps int) int {
	svc, err := newMigrationService(dsn, migrationsDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: create migrator: %v\n", err)
		return 1
	}
	defer svc.Close()

	result, err := svc.Up(ctx, steps)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: migration up failed: %v\n", err)
		return 1
	}
	printOperationResult(result)
	return 0
}

func runDown(ctx context.Context, dsn, migrationsDir string, steps int) int {
	svc, err := newMigrationService(dsn, migrationsDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: create migrator: %v\n", err)
		return 1
	}
	defer svc.Close()

	result, err := svc.Down(ctx, steps)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: migration down failed: %v\n", err)
		return 1
	}
	printOperationResult(result)
	return 0
}

func newMigrationService(dsn, migrationsDir string) (*migrest.Service, error) {
	return migrest.NewService(migrest.ServiceConfig{
		ServiceName:             "rza_migrate_test",
		DatabaseURL:             dsn,
		MigrationsDir:           migrationsDir,
		AllowDown:               true,
		OperationTimeoutSeconds: int(defaultOperationTimeout.Seconds()),
	})
}

func printOperationResult(result migrest.OperationResult) {
	before := "<none>"
	if result.Before.Version != nil {
		before = strconv.FormatUint(*result.Before.Version, 10)
	}
	after := "<none>"
	if result.After.Version != nil {
		after = strconv.FormatUint(*result.After.Version, 10)
	}
	fmt.Printf("[result] operation=%s before=%s after=%s no_change=%t\n", result.Operation, before, after, result.NoChange)
}

func runReset(ctx context.Context, dsn, database, confirmation string, dryRun bool) int {
	if isSystemDatabase(database) {
		fmt.Fprintf(os.Stderr, "error: reset is forbidden for system database %q\n", database)
		return 1
	}
	if !dryRun && !strings.EqualFold(strings.TrimSpace(confirmation), strings.TrimSpace(database)) {
		fmt.Fprintf(os.Stderr, "error: reset requires -confirm-reset %s\n", database)
		return 1
	}

	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: open SQL Server: %v\n", err)
		return 1
	}
	defer db.Close()
	if err := db.PingContext(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error: connect SQL Server: %v\n", err)
		return 1
	}

	objects, err := listResetObjects(ctx, db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: inspect dbo objects: %v\n", err)
		return 1
	}
	if dryRun {
		foreignKeys, fkErr := listResetForeignKeys(ctx, db)
		if fkErr != nil {
			fmt.Fprintf(os.Stderr, "error: inspect dbo foreign keys: %v\n", fkErr)
			return 1
		}
		fmt.Printf("[reset] database=%s dry_run=true foreign_keys=%d objects=%d\n", database, len(foreignKeys), len(objects))
		for _, fk := range foreignKeys {
			fmt.Printf("[reset] would_drop=FOREIGN_KEY %s.%s.%s\n", fk.Schema, fk.Table, fk.Constraint)
		}
		for _, object := range objects {
			fmt.Printf("[reset] would_drop=%s %s.%s\n", object.Kind, object.Schema, object.Name)
		}
		return 0
	}
	foreignKeys, err := listResetForeignKeys(ctx, db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: inspect dbo foreign keys: %v\n", err)
		return 1
	}
	for _, fk := range foreignKeys {
		statement := fmt.Sprintf("ALTER TABLE %s.%s DROP CONSTRAINT %s;", quoteIdentifier(fk.Schema), quoteIdentifier(fk.Table), quoteIdentifier(fk.Constraint))
		if _, execErr := db.ExecContext(ctx, statement); execErr != nil {
			fmt.Fprintf(os.Stderr, "error: drop FOREIGN_KEY %s.%s.%s: %v\n", fk.Schema, fk.Table, fk.Constraint, execErr)
			return 1
		}
	}

	for _, object := range objects {
		statement, buildErr := dropStatement(object)
		if buildErr != nil {
			fmt.Fprintf(os.Stderr, "error: build reset statement: %v\n", buildErr)
			return 1
		}
		if _, execErr := db.ExecContext(ctx, statement); execErr != nil {
			fmt.Fprintf(os.Stderr, "error: drop %s %s.%s: %v\n", object.Kind, object.Schema, object.Name, execErr)
			return 1
		}
	}
	fmt.Printf("[reset] database=%s dropped_foreign_keys=%d dropped_objects=%d schema=dbo\n", database, len(foreignKeys), len(objects))
	return 0
}

func listResetForeignKeys(ctx context.Context, db *sql.DB) ([]resetForeignKey, error) {
	const query = `
SELECT SCHEMA_NAME(t.schema_id), t.name, fk.name
FROM sys.foreign_keys fk
JOIN sys.tables t ON t.object_id = fk.parent_object_id
WHERE SCHEMA_NAME(t.schema_id) = N'dbo'
ORDER BY SCHEMA_NAME(t.schema_id), t.name, fk.name;`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := make([]resetForeignKey, 0)
	for rows.Next() {
		var item resetForeignKey
		if err := rows.Scan(&item.Schema, &item.Table, &item.Constraint); err != nil {
			return nil, err
		}
		out = append(out, item)
	}
	return out, rows.Err()
}

func listResetObjects(ctx context.Context, db *sql.DB) ([]resetObject, error) {
	const query = `
SELECT kind, schema_name, object_name
FROM (
  SELECT 10 AS sort_order, N'VIEW' AS kind, SCHEMA_NAME(v.schema_id) AS schema_name, v.name AS object_name
  FROM sys.views v WHERE SCHEMA_NAME(v.schema_id) = N'dbo'
  UNION ALL
  SELECT 20, N'PROCEDURE', SCHEMA_NAME(p.schema_id), p.name
  FROM sys.procedures p WHERE p.is_ms_shipped = 0 AND SCHEMA_NAME(p.schema_id) = N'dbo'
  UNION ALL
  SELECT 30, N'FUNCTION', SCHEMA_NAME(o.schema_id), o.name
  FROM sys.objects o WHERE o.type IN (N'FN', N'IF', N'TF', N'FS', N'FT') AND o.is_ms_shipped = 0 AND SCHEMA_NAME(o.schema_id) = N'dbo'
  UNION ALL
  SELECT 40, N'SYNONYM', SCHEMA_NAME(s.schema_id), s.name
  FROM sys.synonyms s WHERE SCHEMA_NAME(s.schema_id) = N'dbo'
  UNION ALL
  SELECT 50, N'TABLE', SCHEMA_NAME(t.schema_id), t.name
  FROM sys.tables t WHERE t.is_ms_shipped = 0 AND SCHEMA_NAME(t.schema_id) = N'dbo'
  UNION ALL
  SELECT 60, N'SEQUENCE', SCHEMA_NAME(s.schema_id), s.name
  FROM sys.sequences s WHERE SCHEMA_NAME(s.schema_id) = N'dbo'
  UNION ALL
  SELECT 70, N'TYPE', SCHEMA_NAME(t.schema_id), t.name
  FROM sys.types t WHERE t.is_user_defined = 1 AND SCHEMA_NAME(t.schema_id) = N'dbo'
) objects
ORDER BY sort_order, schema_name, object_name;`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	objects := make([]resetObject, 0)
	for rows.Next() {
		var object resetObject
		if err := rows.Scan(&object.Kind, &object.Schema, &object.Name); err != nil {
			return nil, err
		}
		objects = append(objects, object)
	}
	return objects, rows.Err()
}

func dropStatement(object resetObject) (string, error) {
	kind := strings.ToUpper(strings.TrimSpace(object.Kind))
	switch kind {
	case "VIEW", "PROCEDURE", "FUNCTION", "SYNONYM", "SEQUENCE", "TABLE", "TYPE":
		return fmt.Sprintf("DROP %s %s.%s;", kind, quoteIdentifier(object.Schema), quoteIdentifier(object.Name)), nil
	default:
		return "", fmt.Errorf("unsupported reset object kind %q", object.Kind)
	}
}

func quoteIdentifier(value string) string {
	return "[" + strings.ReplaceAll(value, "]", "]]") + "]"
}

func isSystemDatabase(name string) bool {
	switch strings.ToLower(strings.TrimSpace(name)) {
	case "master", "model", "msdb", "tempdb":
		return true
	default:
		return false
	}
}

func validateDatabaseName(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New("name must not be empty")
	}
	if len([]rune(name)) > 128 {
		return errors.New("name must be at most 128 characters")
	}
	for _, ch := range name {
		if ch == 0 || ch == '\r' || ch == '\n' {
			return errors.New("name contains a forbidden control character")
		}
	}
	return nil
}

func loadNFMCredentials(path string) (nfmCredentials, error) {
	file, err := os.Open(path)
	if err != nil {
		return nfmCredentials{}, err
	}
	defer file.Close()

	values := make(map[string]string)
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(strings.TrimPrefix(scanner.Text(), "\ufeff"))
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			return nfmCredentials{}, fmt.Errorf("invalid .nfm line %d: expected key=value", lineNumber)
		}
		key = strings.ToLower(strings.TrimSpace(key))
		if key == "" {
			return nfmCredentials{}, fmt.Errorf("invalid .nfm line %d: empty key", lineNumber)
		}
		values[key] = strings.TrimSpace(value)
	}
	if err := scanner.Err(); err != nil {
		return nfmCredentials{}, err
	}

	creds := nfmCredentials{
		DatabaseType: values["database_type"],
		Host:         values["host"],
		Username:     values["username"],
		Password:     values["password"],
		Database:     values["initial_catalog"],
	}
	if err := creds.validate(); err != nil {
		return nfmCredentials{}, err
	}
	return creds, nil
}

func (c nfmCredentials) validate() error {
	if databaseType := strings.ToLower(strings.TrimSpace(c.DatabaseType)); databaseType != "mssql" && databaseType != "sqlserver" && databaseType != "sql server" {
		return fmt.Errorf("database_type must be mssql or sqlserver")
	}
	missing := make([]string, 0, 4)
	for key, value := range map[string]string{
		"host":            c.Host,
		"username":        c.Username,
		"password":        c.Password,
		"initial_catalog": c.Database,
	} {
		if strings.TrimSpace(value) == "" {
			missing = append(missing, key)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing required .nfm keys: %s", strings.Join(missing, ", "))
	}
	return nil
}

func (c nfmCredentials) databaseURL() (string, error) {
	host := strings.TrimSpace(c.Host)
	port := 1433
	if comma := strings.LastIndex(host, ","); comma > 0 {
		parsedPort, err := strconv.Atoi(strings.TrimSpace(host[comma+1:]))
		if err != nil || parsedPort <= 0 {
			return "", fmt.Errorf("host port is invalid")
		}
		host = strings.TrimSpace(host[:comma])
		port = parsedPort
	}
	return (migrest.DatabaseDSN{
		Server:   host,
		Port:     port,
		Username: c.Username,
		Password: c.Password,
		Database: c.Database,
		Encrypt:  "disable",
	}).BuildURL()
}

func printUsage() {
	fmt.Fprintln(os.Stderr, "Usage: rza_migrate_test <create-db|status|up|down|reset> [-repo .] [-config mssql_credential.nfm] [-database <isolated_db>] [-timeout 5m] [-catalog-map source=test]")
	fmt.Fprintln(os.Stderr, "  create-db -database <db> -confirm-reset <db> [-collation <name>]  Create an isolated non-system test catalog")
	fmt.Fprintln(os.Stderr, "  up [-steps N]                 Apply all pending migrations, or N steps")
	fmt.Fprintln(os.Stderr, "  down -steps N                 Roll back exactly N migrations")
	fmt.Fprintln(os.Stderr, "  reset -dry-run                List dbo objects that reset would remove")
	fmt.Fprintln(os.Stderr, "  reset -confirm-reset <db>     Drop all non-system objects in dbo, including schema_migrations")
}
