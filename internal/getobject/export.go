package getobject

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "github.com/microsoft/go-mssqldb"

	migrateconfig "rza_migrate_db/internal/config"
	"rza_migrate_db/internal/history"
	"rza_migrate_db/internal/management"
	"rza_migrate_db/internal/sqlfmt"
)

var migrationObjectTokenRunes = func() map[rune]struct{} {
	cfg := migrateconfig.Default(".")
	return cfg.AllowedObjectTokenRunes
}()

type tableMeta struct {
	ObjectID int
	Name     string
}

type objectMeta struct {
	ObjectID int
	Name     string
}

type columnMeta struct {
	ColumnID                  int
	Name                      string
	TypeName                  string
	MaxLength                 int16
	Precision                 int
	Scale                     int
	IsNullable                bool
	IsIdentity                bool
	IdentityNotForReplication bool
	SeedValue                 sql.NullString
	IncValue                  sql.NullString
	Collation                 sql.NullString
	IsComputed                bool
	ComputedDefinition        sql.NullString
	IsPersisted               bool
	IsRowGuidCol              bool
	IsSparse                  bool
	IsFileStream              bool
}

type keyConstraintMeta struct {
	Name                string
	Type                string
	IndexID             int
	IndexType           string
	FillFactor          uint8
	IsPadded            bool
	IgnoreDupKey        bool
	AllowRowLocks       bool
	AllowPageLocks      bool
	IsDisabled          bool
	DataSpaceType       sql.NullString
	DataSpaceID         int
	CompressionDesc     sql.NullString
	HasMixedCompression bool
}

type checkConstraintMeta struct {
	Name                string
	Definition          string
	IsNotForReplication bool
	IsDisabled          bool
	IsNotTrusted        bool
}

type defaultConstraintMeta struct {
	Name       string
	ColumnName string
	Definition string
}

type foreignKeyMeta struct {
	ObjectID            int
	Name                string
	RefSchema           string
	RefTable            string
	DeleteAction        string
	UpdateAction        string
	IsNotForReplication bool
	IsDisabled          bool
	IsNotTrusted        bool
}

type indexMeta struct {
	Name                string
	IndexID             int
	TypeDesc            string
	IsUnique            bool
	FilterDefinition    sql.NullString
	HasFilter           bool
	CompressionDesc     sql.NullString
	HasMixedCompression bool
	XMLSecondaryType    sql.NullString
	XMLPrimaryName      sql.NullString
	SpatialScheme       sql.NullString
	BoundingXMin        sql.NullFloat64
	BoundingYMin        sql.NullFloat64
	BoundingXMax        sql.NullFloat64
	BoundingYMax        sql.NullFloat64
	CellsPerObject      sql.NullInt64
	Level1Grid          sql.NullString
	Level2Grid          sql.NullString
	Level3Grid          sql.NullString
	Level4Grid          sql.NullString
	FillFactor          uint8
	IsPadded            bool
	IgnoreDupKey        bool
	AllowRowLocks       bool
	AllowPageLocks      bool
	IsDisabled          bool
	DataSpaceType       sql.NullString
	DataSpaceID         int
}

type fkColumnMap struct {
	ParentColumn string
	RefColumn    string
}

type indexColumnMeta struct {
	ColumnName string
	Desc       bool
}

// ExportOptions controls extractor runtime behavior.
type ExportOptions struct {
	DryRun   bool
	Progress func(phase, message string)
}

// CategoryChange summarizes delta for one db_object category.
type CategoryChange struct {
	Dir       string
	Created   []string
	Updated   []string
	Deleted   []string
	Unchanged []string
}

// categoryWrite carries rendered files plus names that must be preserved. A
// preserved name belongs to an object that exists in SQL Server but could not
// be rendered safely; treating it as absent would delete the last good source
// snapshot and silently corrupt desired state.
type categoryWrite struct {
	Dir      string
	Files    map[string]string
	Preserve map[string]struct{}
}

// ExportWarning captures structured warning details produced during export.
type ExportWarning struct {
	Rule           string `json:"rule"`
	Category       string `json:"category"`
	Object         string `json:"object"`
	Constraint     string `json:"constraint,omitempty"`
	ConstraintType string `json:"constraint_type,omitempty"`
	Enabled        *bool  `json:"enabled,omitempty"`
	Trusted        *bool  `json:"trusted,omitempty"`
	Message        string `json:"message"`
}

// ExportSummary describes the export result.
type ExportSummary struct {
	DryRun         bool
	Categories     []CategoryChange
	ModuleWarnings []string
	TableWarnings  []ExportWarning
	ObjectWarnings []ExportWarning
}

// TotalCreated returns total created files across all categories.
func (s ExportSummary) TotalCreated() int {
	return totalByCategory(s.Categories, func(c CategoryChange) int { return len(c.Created) })
}

// TotalUpdated returns total updated files across all categories.
func (s ExportSummary) TotalUpdated() int {
	return totalByCategory(s.Categories, func(c CategoryChange) int { return len(c.Updated) })
}

// TotalDeleted returns total deleted files across all categories.
func (s ExportSummary) TotalDeleted() int {
	return totalByCategory(s.Categories, func(c CategoryChange) int { return len(c.Deleted) })
}

// TotalUnchanged returns total unchanged files across all categories.
func (s ExportSummary) TotalUnchanged() int {
	return totalByCategory(s.Categories, func(c CategoryChange) int { return len(c.Unchanged) })
}

// HasChanges returns true when create/update/delete exists.
func (s ExportSummary) HasChanges() bool {
	return s.TotalCreated()+s.TotalUpdated()+s.TotalDeleted() > 0
}

// ConstraintHealthFindingCount returns the number of observed disabled or
// untrusted constraints. These findings do not prevent logical snapshot
// generation, but callers may promote them to a CI/database-health failure.
func (s ExportSummary) ConstraintHealthFindingCount() int {
	count := 0
	for _, warning := range s.TableWarnings {
		if warning.Rule == "CONSTRAINT_STATE_REVIEW_REQUIRED" {
			count++
		}
	}
	return count
}

// WarningMessages returns flattened warning messages (module + table warnings).
func (s ExportSummary) WarningMessages() []string {
	out := make([]string, 0, len(s.ModuleWarnings)+len(s.TableWarnings)+len(s.ObjectWarnings))
	out = append(out, s.ModuleWarnings...)
	for _, w := range s.TableWarnings {
		object := strings.TrimSpace(w.Object)
		msg := strings.TrimSpace(w.Message)
		switch {
		case object != "" && msg != "":
			out = append(out, object+": "+msg)
		case msg != "":
			out = append(out, msg)
		case object != "":
			out = append(out, object)
		}
	}
	for _, w := range s.ObjectWarnings {
		if strings.TrimSpace(w.Object) != "" {
			out = append(out, strings.TrimSpace(w.Object)+": "+strings.TrimSpace(w.Message))
		}
	}
	return out
}

// ExportFromDatabase snapshots dbo objects from SQL Server into db_object.
func ExportFromDatabase(cfg AppConfig) error {
	_, err := ExportFromDatabaseWithOptions(cfg, ExportOptions{})
	return err
}

func emitProgress(opts ExportOptions, phase, format string, args ...interface{}) {
	if opts.Progress == nil {
		return
	}
	opts.Progress(phase, fmt.Sprintf(format, args...))
}

// ExportFromDatabaseWithOptions snapshots dbo objects from SQL Server into db_object.
// When DryRun=true, no files are written/deleted and summary is returned.
func ExportFromDatabaseWithOptions(cfg AppConfig, opts ExportOptions) (ExportSummary, error) {
	cfg.Normalize()
	if err := cfg.Validate(); err != nil {
		return ExportSummary{}, err
	}
	initialObjects, err := loadInitialMigrationObjects(cfg)
	if err != nil {
		return ExportSummary{}, err
	}
	cfg.initialMigrationObjects = initialObjects
	emitProgress(opts, "init", "start export (dry_run=%t, render_workers=%d, write_workers=%d)", opts.DryRun, cfg.RenderWorkers, cfg.WriteWorkers)
	emitProgress(opts, "init", "initial migration objects excluded: %d", len(initialObjects))

	emitProgress(opts, "connect", "open SQL Server connection")
	db, err := sql.Open("sqlserver", cfg.ConnectionString)
	if err != nil {
		return ExportSummary{}, fmt.Errorf("open sqlserver connection: %w", err)
	}
	defer db.Close()

	emitProgress(opts, "connect", "ping SQL Server")
	pingCtx, cancelPing := context.WithTimeout(context.Background(), time.Duration(cfg.ConnectTimeoutSeconds)*time.Second)
	defer cancelPing()
	if err := db.PingContext(pingCtx); err != nil {
		return ExportSummary{}, fmt.Errorf("ping sqlserver: %w", err)
	}
	emitProgress(opts, "connect", "ping success")

	emitProgress(opts, "load", "load table types")
	typeFiles, typeWarnings, typePreserve, err := loadTableTypeFiles(db, cfg)
	if err != nil {
		return ExportSummary{}, err
	}
	emitProgress(opts, "load", "table types loaded: %d (skipped=%d)", len(typeFiles), len(typeWarnings))

	emitProgress(opts, "load", "load table list")
	tables, err := loadTables(db, cfg)
	if err != nil {
		return ExportSummary{}, err
	}
	emitProgress(opts, "load", "tables discovered: %d", len(tables))

	emitProgress(opts, "load", "load table metadata cache")
	cache, err := loadTableSnapshotCache(db, cfg)
	if err != nil {
		return ExportSummary{}, err
	}
	emitProgress(opts, "load", "table metadata cache ready")

	emitProgress(opts, "render", "render table/index snapshot files")
	tableFiles, indexFiles, tableWarnings, err := buildTableAndIndexFiles(cfg, tables, cache)
	if err != nil {
		return ExportSummary{}, err
	}
	excludeInitialMigrationFiles(cfg, "indexes", indexFiles)
	emitProgress(opts, "render", "rendered tables=%d indexes=%d (table_warnings=%d)", len(tableFiles), len(indexFiles), len(tableWarnings))
	if len(tableWarnings) > 0 {
		skipped := 0
		for _, warning := range tableWarnings {
			if warning.Rule == "TABLE_EXTRACT_SKIPPED_UNSUPPORTED" {
				skipped++
			}
		}
		emitProgress(opts, "warn", "table health findings=%d (tables_skipped=%d)", len(tableWarnings), skipped)
	}

	emitProgress(opts, "load", "load views")
	viewFiles, viewWarnings, viewPreserve, err := loadProgrammabilityFiles(db, cfg, "V", "VIEW")
	if err != nil {
		return ExportSummary{}, err
	}
	emitProgress(opts, "load", "views loaded: %d (skipped=%d)", len(viewFiles), len(viewWarnings))
	if len(viewWarnings) > 0 {
		emitProgress(opts, "warn", "views skipped: %d", len(viewWarnings))
	}

	emitProgress(opts, "load", "load functions")
	functionFiles, functionWarnings, functionPreserve, err := loadProgrammabilityFiles(db, cfg, "FN,IF,TF", "FUNCTION")
	if err != nil {
		return ExportSummary{}, err
	}
	emitProgress(opts, "load", "functions loaded: %d (skipped=%d)", len(functionFiles), len(functionWarnings))
	if len(functionWarnings) > 0 {
		emitProgress(opts, "warn", "functions skipped: %d", len(functionWarnings))
	}

	emitProgress(opts, "load", "load procedures")
	procedureFiles, procedureWarnings, procedurePreserve, err := loadProgrammabilityFiles(db, cfg, "P", "PROCEDURE")
	if err != nil {
		return ExportSummary{}, err
	}
	emitProgress(opts, "load", "procedures loaded: %d (skipped=%d)", len(procedureFiles), len(procedureWarnings))
	if len(procedureWarnings) > 0 {
		emitProgress(opts, "warn", "procedures skipped: %d", len(procedureWarnings))
	}

	moduleWarnings := make([]string, 0, len(typeWarnings)+len(viewWarnings)+len(functionWarnings)+len(procedureWarnings))
	moduleWarnings = append(moduleWarnings, typeWarnings...)
	moduleWarnings = append(moduleWarnings, viewWarnings...)
	moduleWarnings = append(moduleWarnings, functionWarnings...)
	moduleWarnings = append(moduleWarnings, procedureWarnings...)
	sort.Strings(moduleWarnings)
	if !cfg.SkipUnsupportedTables && len(moduleWarnings) > 0 {
		return ExportSummary{}, fmt.Errorf(
			"extractor cannot produce a complete desired-state snapshot: %d module/type object(s) could not be rendered safely; no files were changed",
			len(moduleWarnings),
		)
	}

	tablePreserve := preserveNamesFromTableWarnings(tableWarnings)
	writes := []categoryWrite{
		{Dir: filepath.Join(cfg.DbObjectDir, "types"), Files: typeFiles, Preserve: typePreserve},
		{Dir: filepath.Join(cfg.DbObjectDir, "tables"), Files: tableFiles, Preserve: tablePreserve},
		{Dir: filepath.Join(cfg.DbObjectDir, "indexes"), Files: indexFiles, Preserve: tablePreserve},
		{Dir: filepath.Join(cfg.DbObjectDir, "views"), Files: viewFiles, Preserve: viewPreserve},
		{Dir: filepath.Join(cfg.DbObjectDir, "functions"), Files: functionFiles, Preserve: functionPreserve},
		{Dir: filepath.Join(cfg.DbObjectDir, "procedures"), Files: procedureFiles, Preserve: procedurePreserve},
	}
	objectWarnings := warningsFromPreserved("types", cfg.Schema, typePreserve)
	objectWarnings = append(objectWarnings, warningsFromPreserved("views", cfg.Schema, viewPreserve)...)
	objectWarnings = append(objectWarnings, warningsFromPreserved("functions", cfg.Schema, functionPreserve)...)
	objectWarnings = append(objectWarnings, warningsFromPreserved("procedures", cfg.Schema, procedurePreserve)...)
	summary := ExportSummary{
		DryRun:         opts.DryRun,
		Categories:     make([]CategoryChange, 0, len(writes)),
		ModuleWarnings: moduleWarnings,
		TableWarnings:  tableWarnings,
		ObjectWarnings: objectWarnings,
	}
	for _, item := range writes {
		emitProgress(opts, "apply", "apply category %s", filepath.ToSlash(item.Dir))
		change, err := applySQLCategoryPreserving(cfg, item.Dir, item.Files, item.Preserve, opts.DryRun)
		if err != nil {
			return ExportSummary{}, err
		}
		summary.Categories = append(summary.Categories, change)
		emitProgress(opts, "apply", "category %s done (create=%d update=%d delete=%d unchanged=%d)", filepath.ToSlash(item.Dir), len(change.Created), len(change.Updated), len(change.Deleted), len(change.Unchanged))
	}
	if !opts.DryRun {
		manifest := managementManifestFromSummary(cfg.Schema, summary)
		manifestPath := cfg.RepoPath(filepath.FromSlash("out/object_management_manifest.json"))
		if err := management.Write(manifestPath, manifest); err != nil {
			return ExportSummary{}, fmt.Errorf("write object management manifest: %w", err)
		}
	}

	emitProgress(opts, "done", "export finished (create=%d update=%d delete=%d unchanged=%d warnings=%d)", summary.TotalCreated(), summary.TotalUpdated(), summary.TotalDeleted(), summary.TotalUnchanged(), len(summary.WarningMessages()))
	return summary, nil
}

func warningsFromPreserved(category, schema string, preserved map[string]struct{}) []ExportWarning {
	out := make([]ExportWarning, 0, len(preserved))
	for file := range preserved {
		object := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
		if !strings.Contains(object, ".") {
			object = strings.TrimSpace(schema) + "." + object
		}
		out = append(out, ExportWarning{Rule: "OBJECT_EXTRACT_SKIPPED_UNSUPPORTED", Category: category, Object: object, Message: "extractor could not represent this object losslessly; previous snapshot was preserved when available"})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Object < out[j].Object })
	return out
}

func managementManifestFromSummary(schema string, summary ExportSummary) management.Manifest {
	managedByKey := map[string]management.Object{}
	for _, category := range summary.Categories {
		name := strings.ToLower(filepath.Base(filepath.Clean(category.Dir)))
		files := append(append(append([]string{}, category.Created...), category.Updated...), category.Unchanged...)
		for _, file := range files {
			base := filepath.Base(file)
			object := strings.TrimSuffix(base, filepath.Ext(base))
			if object == "" {
				continue
			}
			item := management.Object{Category: name, Object: object}
			managedByKey[name+"|"+object] = item
		}
	}
	unmanagedByKey := map[string]management.Object{}
	for _, warning := range summary.TableWarnings {
		if warning.Rule != "TABLE_EXTRACT_SKIPPED_UNSUPPORTED" {
			continue
		}
		item := management.Object{Category: "tables", Object: strings.TrimSpace(warning.Object), Reason: strings.TrimSpace(warning.Message)}
		key := item.Category + "|" + item.Object
		unmanagedByKey[key] = item
		delete(managedByKey, key)
		indexItem := management.Object{Category: "indexes", Object: item.Object, Reason: item.Reason}
		unmanagedByKey[indexItem.Category+"|"+indexItem.Object] = indexItem
		delete(managedByKey, indexItem.Category+"|"+indexItem.Object)
	}
	for _, warning := range summary.ObjectWarnings {
		item := management.Object{Category: strings.ToLower(strings.TrimSpace(warning.Category)), Object: strings.TrimSpace(warning.Object), Reason: strings.TrimSpace(warning.Message)}
		if item.Category == "" || item.Object == "" {
			continue
		}
		key := item.Category + "|" + item.Object
		unmanagedByKey[key] = item
		delete(managedByKey, key)
	}
	manifest := management.Manifest{Schema: strings.TrimSpace(schema)}
	for _, item := range managedByKey {
		manifest.Managed = append(manifest.Managed, item)
	}
	for _, item := range unmanagedByKey {
		manifest.Unmanaged = append(manifest.Unmanaged, item)
	}
	return manifest
}

func loadTables(db *sql.DB, cfg AppConfig) ([]tableMeta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
SELECT t.object_id, t.name
FROM sys.tables t
JOIN sys.schemas s ON s.schema_id = t.schema_id
WHERE s.name = @schema
  AND t.is_ms_shipped = 0
ORDER BY t.name;
`, sql.Named("schema", cfg.Schema))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]tableMeta, 0)
	for rows.Next() {
		var t tableMeta
		if err := rows.Scan(&t.ObjectID, &t.Name); err != nil {
			return nil, err
		}
		if shouldExcludeExtractedObject(cfg, "tables", cfg.Schema, t.Name) {
			continue
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

func isExtractorInfrastructureTable(schema, name string) bool {
	return isExtractorInfrastructureObject("tables", schema, name)
}

func isExtractorInfrastructureObject(category, schema, name string) bool {
	category = strings.ToLower(strings.TrimSpace(category))
	schema = strings.TrimSpace(schema)
	name = strings.TrimSpace(name)
	if schema != "dbo" {
		return false
	}

	switch category {
	case "tables", "indexes":
		return name == "mig_event" ||
			name == "mig_operation_journal" ||
			name == "schema_migrations"
	case "procedures":
		return name == "usp_mig_log_event" ||
			name == "usp_mig_try_get_lock"
	default:
		return false
	}
}

func shouldExcludeExtractedObject(cfg AppConfig, category, schema, name string) bool {
	if isExtractorInfrastructureObject(category, schema, name) {
		return true
	}
	key := strings.ToLower(strings.TrimSpace(category)) + "|" +
		strings.TrimSpace(schema) + "." + strings.TrimSpace(name)
	_, excluded := cfg.initialMigrationObjects[key]
	return excluded
}

func excludeInitialMigrationFiles(cfg AppConfig, category string, files map[string]string) {
	category = strings.ToLower(strings.TrimSpace(category))
	for name := range files {
		object := strings.TrimSuffix(filepath.Base(name), filepath.Ext(name))
		if _, excluded := cfg.initialMigrationObjects[category+"|"+object]; excluded {
			delete(files, name)
		}
	}
}

// loadInitialMigrationObjects reads initial object migrations from the target
// repository. A v1 up migration is historical bootstrap state, so the
// extractor must not duplicate that object under db_object.
func loadInitialMigrationObjects(cfg AppConfig) (map[string]struct{}, error) {
	out := make(map[string]struct{})
	dir := cfg.RepoPath("migrations")
	entries, err := os.ReadDir(dir)
	if errors.Is(err, os.ErrNotExist) {
		return out, nil
	}
	if err != nil {
		return nil, fmt.Errorf("read initial migrations: %w", err)
	}

	migrationCfg := migrateconfig.Default(cfg.RepoRoot)
	parseOptions := history.ParseOptions{
		CoreAuditUpName:         "000000__core__mig__audit_init.up.sql",
		CoreAuditDownName:       "000000__core__mig__audit_init.down.sql",
		AllowedCategories:       migrationCfg.AllowedCategories,
		AllowedObjectTokenRunes: migrationCfg.AllowedObjectTokenRunes,
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		meta, issue := history.ParseMigrationFileName(entry.Name(), parseOptions)
		if issue != nil || meta.IsCore || meta.Direction != "up" || meta.VK != 1 || meta.Category == "patch" {
			continue
		}
		out[strings.ToLower(meta.Category)+"|"+meta.Object] = struct{}{}
	}
	return out, nil
}

func loadTableTypeFiles(db *sql.DB, cfg AppConfig) (map[string]string, []string, map[string]struct{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
SELECT tt.type_table_object_id, tt.name
FROM sys.table_types tt
JOIN sys.schemas s ON s.schema_id = tt.schema_id
WHERE s.name = @schema
  AND tt.is_user_defined = 1
ORDER BY tt.name;
`, sql.Named("schema", cfg.Schema))
	if err != nil {
		return nil, nil, nil, err
	}
	defer rows.Close()

	types := make([]tableMeta, 0)
	for rows.Next() {
		var item tableMeta
		if err := rows.Scan(&item.ObjectID, &item.Name); err != nil {
			return nil, nil, nil, err
		}
		types = append(types, item)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, nil, err
	}

	files := make(map[string]string, len(types))
	warnings := make([]string, 0)
	preserve := make(map[string]struct{})
	for _, item := range types {
		if shouldExcludeExtractedObject(cfg, "types", cfg.Schema, item.Name) {
			continue
		}
		fileName := fileNameForObject(item.Name)
		if err := validateObjectNameForRules(item.Name); err != nil {
			warnings = append(warnings, fmt.Sprintf("table type %s.%s skipped: invalid object name (%v)", cfg.Schema, item.Name, err))
			preserve[fileName] = struct{}{}
			continue
		}
		typeSQL, buildErr := buildTableTypeSnapshotSQL(db, cfg, item)
		if buildErr != nil {
			warnings = append(warnings, fmt.Sprintf("table type %s.%s skipped: %v", cfg.Schema, item.Name, buildErr))
			preserve[fileName] = struct{}{}
			continue
		}
		files[fileName] = normalizeSQL(typeSQL)
	}
	sort.Strings(warnings)
	return files, warnings, preserve, nil
}

func buildTableTypeSnapshotSQL(db *sql.DB, cfg AppConfig, item tableMeta) (string, error) {
	cols, err := loadColumns(db, cfg, item.ObjectID)
	if err != nil {
		return "", err
	}
	if len(cols) == 0 {
		return "", errors.New("table type has zero columns")
	}

	defaults, err := loadDefaultConstraints(db, cfg, item.ObjectID)
	if err != nil {
		return "", err
	}
	defaultByColumn, err := defaultConstraintMapByColumn(defaults)
	if err != nil {
		return "", err
	}

	items := make([]string, 0, len(cols))
	for _, column := range cols {
		if column.IsComputed || column.IsSparse || column.IsFileStream || column.IdentityNotForReplication {
			return "", fmt.Errorf("table type column %s uses unsupported metadata", column.Name)
		}
		clause := renderColumnClause(column)
		if def, exists := defaultByColumn[strings.TrimSpace(column.Name)]; exists {
			clause += " DEFAULT " + strings.TrimSpace(def.Definition)
		}
		items = append(items, clause)
	}

	checks, err := loadCheckConstraints(db, cfg, item.ObjectID)
	if err != nil {
		return "", err
	}
	for _, check := range checks {
		if check.IsDisabled || check.IsNotTrusted || check.IsNotForReplication {
			return "", fmt.Errorf("table type check constraint uses unsupported state")
		}
		items = append(items, "CHECK "+strings.TrimSpace(check.Definition))
	}

	keys, err := loadKeyConstraints(db, cfg, item.ObjectID)
	if err != nil {
		return "", err
	}
	for _, key := range keys {
		if err := validateKeyConstraintMetadata(key); err != nil {
			return "", err
		}
		keyCols, err := loadKeyColumns(db, cfg, item.ObjectID, key.IndexID)
		if err != nil {
			return "", err
		}
		if len(keyCols) == 0 {
			return "", fmt.Errorf("table type key has zero columns")
		}
		keyType := "UNIQUE"
		if strings.EqualFold(key.Type, "PK") {
			keyType = "PRIMARY KEY"
		}
		indexType := strings.ToUpper(strings.TrimSpace(key.IndexType))
		if indexType != "CLUSTERED" && indexType != "NONCLUSTERED" {
			return "", fmt.Errorf("table type key uses unsupported index type %s", key.IndexType)
		}
		items = append(items, keyType+" "+indexType+" ("+renderIndexColumnList(keyCols, true)+")")
	}

	indexes, err := loadIndexes(db, cfg, item.ObjectID)
	if err != nil {
		return "", err
	}
	for _, index := range indexes {
		if index.FilterDefinition.Valid && strings.TrimSpace(index.FilterDefinition.String) != "" {
			return "", fmt.Errorf("table type index %s uses unsupported filter", index.Name)
		}
		indexType := strings.ToUpper(strings.TrimSpace(index.TypeDesc))
		if indexType != "CLUSTERED" && indexType != "NONCLUSTERED" {
			return "", fmt.Errorf("table type index %s uses unsupported type %s", index.Name, index.TypeDesc)
		}
		keyCols, err := loadIndexColumns(db, cfg, item.ObjectID, index.IndexID, false)
		if err != nil {
			return "", err
		}
		includeCols, err := loadIndexColumns(db, cfg, item.ObjectID, index.IndexID, true)
		if err != nil {
			return "", err
		}
		if len(keyCols) == 0 || len(includeCols) > 0 {
			return "", fmt.Errorf("table type index %s has unsupported key/include columns", index.Name)
		}
		unique := ""
		if index.IsUnique {
			unique = " UNIQUE"
		}
		items = append(items, "INDEX "+quoteIdentifier(index.Name)+" "+indexType+unique+" ("+renderIndexColumnList(keyCols, true)+")")
	}

	var rendered strings.Builder
	rendered.WriteString("CREATE TYPE ")
	rendered.WriteString(qualifiedName(cfg.Schema, item.Name))
	rendered.WriteString(" AS TABLE (\n")
	for index, clause := range items {
		rendered.WriteString("  ")
		rendered.WriteString(clause)
		if index < len(items)-1 {
			rendered.WriteString(",")
		}
		rendered.WriteString("\n")
	}
	rendered.WriteString(");\n")
	return rendered.String(), nil
}

func loadProgrammabilityFiles(db *sql.DB, cfg AppConfig, typeFilter, kind string) (map[string]string, []string, map[string]struct{}, error) {
	typeFilter = strings.TrimSpace(typeFilter)
	if typeFilter == "" {
		return map[string]string{}, []string{}, map[string]struct{}{}, nil
	}

	var typeList []string
	for _, t := range strings.Split(typeFilter, ",") {
		token := strings.TrimSpace(t)
		if token == "" {
			continue
		}
		typeList = append(typeList, "'"+token+"'")
	}
	if len(typeList) == 0 {
		return map[string]string{}, []string{}, map[string]struct{}{}, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	query := `
SELECT o.object_id, o.name, m.definition
FROM sys.objects o
JOIN sys.schemas s ON s.schema_id = o.schema_id
JOIN sys.sql_modules m ON m.object_id = o.object_id
WHERE s.name = @schema
  AND o.is_ms_shipped = 0
  AND o.type IN (` + strings.Join(typeList, ",") + `)
ORDER BY o.name;
`
	rows, err := db.QueryContext(ctx, query, sql.Named("schema", cfg.Schema))
	if err != nil {
		return nil, nil, nil, err
	}
	defer rows.Close()

	files := map[string]string{}
	warnings := make([]string, 0)
	preserve := map[string]struct{}{}
	for rows.Next() {
		var obj objectMeta
		var definition sql.NullString
		if err := rows.Scan(&obj.ObjectID, &obj.Name, &definition); err != nil {
			return nil, nil, nil, err
		}
		if shouldExcludeExtractedObject(cfg, strings.ToLower(kind)+"s", cfg.Schema, obj.Name) {
			continue
		}
		fileName := fileNameForObject(obj.Name)
		if err := validateObjectNameForRules(obj.Name); err != nil {
			warnings = append(warnings, fmt.Sprintf("%s %s.%s skipped: invalid object name (%v)", strings.ToLower(kind), cfg.Schema, obj.Name, err))
			preserve[fileName] = struct{}{}
			continue
		}
		if !definition.Valid || strings.TrimSpace(definition.String) == "" {
			warnings = append(warnings, fmt.Sprintf("%s %s.%s skipped: definition is empty or encrypted", strings.ToLower(kind), cfg.Schema, obj.Name))
			preserve[fileName] = struct{}{}
			continue
		}
		sqlText, err := ensureCreateOrAlter(definition.String, kind)
		if err != nil {
			warnings = append(warnings, fmt.Sprintf("%s %s.%s skipped: %v", strings.ToLower(kind), cfg.Schema, obj.Name, err))
			preserve[fileName] = struct{}{}
			continue
		}
		sqlText = formatModuleDefinition(sqlText, cfg)
		files[fileName] = normalizeSQL(sqlText)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, nil, err
	}
	sort.Strings(warnings)
	return files, warnings, preserve, nil
}

func preserveNamesFromTableWarnings(warnings []ExportWarning) map[string]struct{} {
	out := make(map[string]struct{}, len(warnings))
	for _, warning := range warnings {
		// Health findings describe a table that was rendered successfully.
		// Only an actual skip is allowed to preserve the previous table/index
		// files; preserving on every warning could retain a stale index file.
		if warning.Rule != "TABLE_EXTRACT_SKIPPED_UNSUPPORTED" {
			continue
		}
		name := strings.TrimSpace(warning.Object)
		if dot := strings.Index(name, "."); dot >= 0 {
			name = name[dot+1:]
		}
		if name != "" {
			out[fileNameForObject(strings.Trim(name, "[]"))] = struct{}{}
		}
	}
	return out
}

type tableBuildResult struct {
	TableFile string
	TableSQL  string
	IndexFile string
	IndexSQL  string
	Warnings  []ExportWarning
	Err       error
}

func buildTableAndIndexFiles(cfg AppConfig, tables []tableMeta, cache tableSnapshotCache) (map[string]string, map[string]string, []ExportWarning, error) {
	tableFiles := make(map[string]string, len(tables))
	indexFiles := make(map[string]string, len(tables))
	warnings := make([]ExportWarning, 0)
	if len(tables) == 0 {
		return tableFiles, indexFiles, warnings, nil
	}

	caseFoldedFiles := make(map[string]string, len(tables))
	for _, t := range tables {
		if err := validateObjectNameForRules(t.Name); err != nil {
			return nil, nil, nil, fmt.Errorf("table %q invalid for rules contract: %w", t.Name, err)
		}
		fileName := fileNameForObject(t.Name)
		folded := strings.ToLower(fileName)
		if existing, ok := caseFoldedFiles[folded]; ok && existing != fileName {
			return nil, nil, nil, fmt.Errorf("case-distinct table names cannot share a portable db_object filesystem: %s and %s", existing, fileName)
		}
		caseFoldedFiles[folded] = fileName
	}

	workers := normalizedWorkerCount(cfg.RenderWorkers, len(tables))
	jobs := make(chan int, len(tables))
	results := make([]tableBuildResult, len(tables))
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx := range jobs {
				t := tables[idx]
				result := tableBuildResult{
					TableFile: fileNameForObject(t.Name),
					IndexFile: fileNameForObject(t.Name),
					Warnings:  tableHealthWarningsFromCache(cfg, t, cache),
				}
				tableSQL, err := buildTableSnapshotSQLFromCache(cfg, t, cache)
				if err != nil {
					result.Err = fmt.Errorf("build table snapshot %s.%s: %w", cfg.Schema, t.Name, err)
					results[idx] = result
					continue
				}
				indexSQL, err := buildIndexSnapshotSQLFromCache(cfg, t, cache)
				if err != nil {
					result.Err = fmt.Errorf("build index snapshot %s.%s: %w", cfg.Schema, t.Name, err)
					results[idx] = result
					continue
				}
				result.TableSQL = normalizeSQL(tableSQL)
				result.IndexSQL = normalizeSQL(indexSQL)
				results[idx] = result
			}
		}()
	}

	for idx := range tables {
		jobs <- idx
	}
	close(jobs)
	wg.Wait()

	for _, r := range results {
		warnings = append(warnings, r.Warnings...)
		if r.Err != nil {
			if cfg.SkipUnsupportedTables && isUnsupportedTableBuildError(r.Err) {
				objectName := strings.TrimSuffix(strings.TrimPrefix(r.TableFile, "dbo."), ".sql")
				warnings = append(warnings, ExportWarning{
					Rule:     "TABLE_EXTRACT_SKIPPED_UNSUPPORTED",
					Category: "tables",
					Object:   qualifiedName(cfg.Schema, objectName),
					Message:  r.Err.Error(),
				})
				continue
			}
			return nil, nil, nil, r.Err
		}
		tableFiles[r.TableFile] = r.TableSQL
		if strings.TrimSpace(r.IndexSQL) != "" {
			indexFiles[r.IndexFile] = r.IndexSQL
		}
	}
	return tableFiles, indexFiles, warnings, nil
}

func tableHealthWarningsFromCache(cfg AppConfig, table tableMeta, cache tableSnapshotCache) []ExportWarning {
	object := qualifiedName(cfg.Schema, table.Name)
	warnings := make([]ExportWarning, 0)

	for _, check := range cache.Checks[table.ObjectID] {
		if !check.IsDisabled && !check.IsNotTrusted {
			continue
		}
		enabled := !check.IsDisabled
		trusted := !check.IsNotTrusted
		warnings = append(warnings, ExportWarning{
			Rule:           "CONSTRAINT_STATE_REVIEW_REQUIRED",
			Category:       "tables",
			Object:         object,
			Constraint:     check.Name,
			ConstraintType: "CHECK",
			Enabled:        boolPointer(enabled),
			Trusted:        boolPointer(trusted),
			Message:        fmt.Sprintf("check constraint %s is disabled or untrusted; its logical definition was exported without changing the database", check.Name),
		})
	}

	for _, key := range cache.Keys[table.ObjectID] {
		if key.IsDisabled {
			enabled := false
			warnings = append(warnings, ExportWarning{
				Rule:           "CONSTRAINT_STATE_REVIEW_REQUIRED",
				Category:       "tables",
				Object:         object,
				Constraint:     key.Name,
				ConstraintType: keyConstraintTypeName(key),
				Enabled:        boolPointer(enabled),
				Message:        fmt.Sprintf("key constraint %s has a disabled backing index; its logical definition was exported without changing the database", key.Name),
			})
		}
		if details := ignoredKeyPhysicalMetadata(key); len(details) > 0 {
			warnings = append(warnings, ExportWarning{
				Rule:           "KEY_PHYSICAL_METADATA_NOT_CAPTURED",
				Category:       "tables",
				Object:         object,
				Constraint:     key.Name,
				ConstraintType: keyConstraintTypeName(key),
				Message:        fmt.Sprintf("key constraint %s physical metadata is managed outside the logical table snapshot: %s", key.Name, strings.Join(details, ", ")),
			})
		}
	}

	for _, fk := range cache.ForeignKeys[table.ObjectID] {
		if !fk.IsDisabled && !fk.IsNotTrusted {
			continue
		}
		enabled := !fk.IsDisabled
		trusted := !fk.IsNotTrusted
		warnings = append(warnings, ExportWarning{
			Rule:           "CONSTRAINT_STATE_REVIEW_REQUIRED",
			Category:       "tables",
			Object:         object,
			Constraint:     fk.Name,
			ConstraintType: "FOREIGN_KEY",
			Enabled:        boolPointer(enabled),
			Trusted:        boolPointer(trusted),
			Message:        fmt.Sprintf("foreign key %s is disabled or untrusted; its logical definition was exported without changing the database", fk.Name),
		})
	}
	return warnings
}

func boolPointer(value bool) *bool {
	return &value
}

func keyConstraintTypeName(key keyConstraintMeta) string {
	if strings.EqualFold(strings.TrimSpace(key.Type), "PK") {
		return "PRIMARY_KEY"
	}
	return "UNIQUE"
}

func ignoredKeyPhysicalMetadata(key keyConstraintMeta) []string {
	details := make([]string, 0, 8)
	if key.FillFactor != 0 {
		details = append(details, fmt.Sprintf("FILLFACTOR=%d", key.FillFactor))
	}
	if key.IsPadded {
		details = append(details, "PAD_INDEX=ON")
	}
	if key.IgnoreDupKey {
		details = append(details, "IGNORE_DUP_KEY=ON")
	}
	if !key.AllowRowLocks {
		details = append(details, "ALLOW_ROW_LOCKS=OFF")
	}
	if !key.AllowPageLocks {
		details = append(details, "ALLOW_PAGE_LOCKS=OFF")
	}
	compression := strings.ToUpper(strings.TrimSpace(key.CompressionDesc.String))
	if key.HasMixedCompression {
		details = append(details, "DATA_COMPRESSION=MIXED")
	} else if compression != "" && compression != "NONE" {
		details = append(details, "DATA_COMPRESSION="+compression)
	}
	if key.DataSpaceType.Valid && strings.TrimSpace(key.DataSpaceType.String) != "" {
		dataSpaceType := strings.ToUpper(strings.TrimSpace(key.DataSpaceType.String))
		if dataSpaceType == "PARTITION_SCHEME" || key.DataSpaceID > 1 {
			details = append(details, fmt.Sprintf("DATA_SPACE=%s:%d", dataSpaceType, key.DataSpaceID))
		}
	} else if key.DataSpaceID > 1 {
		details = append(details, fmt.Sprintf("DATA_SPACE_ID=%d", key.DataSpaceID))
	}
	return details
}

func isUnsupportedTableBuildError(err error) bool {
	if err == nil {
		return false
	}
	msg := strings.ToLower(err.Error())
	switch {
	case strings.Contains(msg, "computed column") && strings.Contains(msg, "not supported"):
		return true
	case strings.Contains(msg, "sparse column") && strings.Contains(msg, "not supported"):
		return true
	case strings.Contains(msg, "filestream column") && strings.Contains(msg, "not supported"):
		return true
	case strings.Contains(msg, "is disabled and cannot be represented"):
		return true
	case strings.Contains(msg, "partition scheme") && strings.Contains(msg, "cannot be represented"):
		return true
	case strings.Contains(msg, "non-default data space") && strings.Contains(msg, "cannot be represented"):
		return true
	case strings.Contains(msg, "metadata is incomplete"):
		return true
	case strings.Contains(msg, "unsupported persisted"):
		return true
	case strings.Contains(msg, "uses unsupported type"):
		return true
	case strings.Contains(msg, "cannot be preserved safely"):
		return true
	case strings.Contains(msg, "cannot be normalized safely"):
		return true
	case strings.Contains(msg, "uses non-default index options"):
		return true
	case strings.Contains(msg, "uses data compression"):
		return true
	case strings.Contains(msg, "key constraint") && strings.Contains(msg, "partition scheme"):
		return true
	case strings.Contains(msg, "key constraint") && strings.Contains(msg, "non-default data space"):
		return true
	case strings.Contains(msg, "key constraint") && strings.Contains(msg, "is disabled"):
		return true
	case strings.Contains(msg, "unsupported index metadata"):
		return true
	case strings.Contains(msg, "unsupported table metadata"):
		return true
	default:
		return false
	}
}

// IsUnsupportedTableBuildError reports whether an export error can be handled
// by SkipUnsupportedTables without deleting the table's last known snapshots.
// It is exposed so CLI callers can provide an actionable remediation message
// instead of returning an opaque fail-closed catalog error.
func IsUnsupportedTableBuildError(err error) bool {
	return isUnsupportedTableBuildError(err)
}

func buildTableSnapshotSQL(db *sql.DB, cfg AppConfig, t tableMeta) (string, error) {
	cols, err := loadColumns(db, cfg, t.ObjectID)
	if err != nil {
		return "", err
	}
	if len(cols) == 0 {
		return "", errors.New("table has zero columns")
	}

	defaults, err := loadDefaultConstraints(db, cfg, t.ObjectID)
	if err != nil {
		return "", err
	}
	defaultByColumn, err := defaultConstraintMapByColumn(defaults)
	if err != nil {
		return "", err
	}

	items := make([]string, 0, len(cols))
	for _, c := range cols {
		if c.IsComputed {
			return "", fmt.Errorf("computed column %s is not supported in extractor output", c.Name)
		}
		if c.IsSparse {
			return "", fmt.Errorf("sparse column %s is not supported in extractor output", c.Name)
		}
		if c.IsFileStream {
			return "", fmt.Errorf("filestream column %s is not supported in extractor output", c.Name)
		}
		if c.IdentityNotForReplication {
			return "", fmt.Errorf("identity column %s uses NOT FOR REPLICATION, which cannot be preserved safely", c.Name)
		}
		clause := renderColumnClause(c)
		if d, ok := defaultByColumn[strings.TrimSpace(c.Name)]; ok {
			clause += renderInlineDefaultConstraintFragment(d)
		}
		items = append(items, clause)
	}

	checks, err := loadCheckConstraints(db, cfg, t.ObjectID)
	if err != nil {
		return "", err
	}
	for _, ck := range checks {
		if strings.TrimSpace(ck.Name) == "" {
			return "", fmt.Errorf("check constraint name must not be empty on table %s", t.Name)
		}
		items = append(items, renderCheckConstraintClause(ck))
	}

	keys, err := loadKeyConstraints(db, cfg, t.ObjectID)
	if err != nil {
		return "", err
	}
	for _, kc := range keys {
		if strings.TrimSpace(kc.Name) == "" {
			return "", fmt.Errorf("key constraint name must not be empty on table %s", t.Name)
		}
		if err := validateKeyConstraintMetadata(kc); err != nil {
			return "", err
		}
		keyCols, err := loadKeyColumns(db, cfg, t.ObjectID, kc.IndexID)
		if err != nil {
			return "", err
		}
		if len(keyCols) == 0 {
			return "", fmt.Errorf("key constraint %s has no key columns", kc.Name)
		}
		items = append(items, renderKeyConstraintClause(kc, keyCols))
	}

	fks, err := loadForeignKeys(db, cfg, t.ObjectID)
	if err != nil {
		return "", err
	}
	for _, fk := range fks {
		if strings.TrimSpace(fk.Name) == "" {
			return "", fmt.Errorf("foreign key name must not be empty on table %s", t.Name)
		}
		if strings.TrimSpace(fk.RefSchema) != strings.TrimSpace(cfg.Schema) {
			return "", fmt.Errorf("foreign key %s references non-dbo schema %s", fk.Name, fk.RefSchema)
		}
		if err := validateObjectNameForRules(fk.RefTable); err != nil {
			return "", fmt.Errorf("foreign key %s references invalid table %s: %w", fk.Name, fk.RefTable, err)
		}
		colMap, err := loadForeignKeyColumns(db, cfg, fk.ObjectID)
		if err != nil {
			return "", err
		}
		if len(colMap) == 0 {
			return "", fmt.Errorf("foreign key %s has no column mapping", fk.Name)
		}
		items = append(items, renderForeignKeyClause(cfg.Schema, fk, colMap))
	}

	var b strings.Builder
	b.WriteString("CREATE TABLE ")
	b.WriteString(qualifiedName(cfg.Schema, t.Name))
	b.WriteString(" (\n")
	for i, it := range items {
		b.WriteString("  ")
		b.WriteString(it)
		if i < len(items)-1 {
			b.WriteString(",")
		}
		b.WriteString("\n")
	}
	b.WriteString(");\n")
	return b.String(), nil
}

func buildIndexSnapshotSQL(db *sql.DB, cfg AppConfig, t tableMeta) (string, error) {
	indexes, err := loadIndexes(db, cfg, t.ObjectID)
	if err != nil {
		return "", err
	}
	if len(indexes) == 0 {
		return "", nil
	}

	stmts := make([]string, 0, len(indexes))
	for _, idx := range indexes {
		keyCols, err := loadIndexColumns(db, cfg, t.ObjectID, idx.IndexID, false)
		if err != nil {
			return "", err
		}
		if len(keyCols) == 0 {
			return "", fmt.Errorf("index %s has zero key columns", idx.Name)
		}

		includeCols, err := loadIndexColumns(db, cfg, t.ObjectID, idx.IndexID, true)
		if err != nil {
			return "", err
		}
		stmt, err := renderCreateIndexClause(cfg.Schema, t.Name, idx, keyCols, includeCols)
		if err != nil {
			return "", err
		}
		stmts = append(stmts, stmt)
	}

	return strings.Join(stmts, "\n\n") + "\n", nil
}

func loadColumns(db *sql.DB, cfg AppConfig, objectID int) ([]columnMeta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
SELECT
  c.column_id,
  c.name,
  ty.name AS type_name,
  c.max_length,
  c.precision,
  c.scale,
  c.is_nullable,
  c.is_identity,
  COALESCE(ic.is_not_for_replication, 0) AS identity_not_for_replication,
  CAST(ic.seed_value AS NVARCHAR(100)) AS seed_value,
  CAST(ic.increment_value AS NVARCHAR(100)) AS increment_value,
  c.collation_name,
  c.is_computed,
  c.is_sparse,
  c.is_filestream
FROM sys.columns c
JOIN sys.types ty
  ON ty.user_type_id = c.user_type_id
LEFT JOIN sys.identity_columns ic
  ON ic.object_id = c.object_id
 AND ic.column_id = c.column_id
WHERE c.object_id = @obj
ORDER BY c.column_id;
`, sql.Named("obj", objectID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols := make([]columnMeta, 0)
	for rows.Next() {
		var c columnMeta
		var precision uint8
		var scale uint8
		if err := rows.Scan(
			&c.ColumnID,
			&c.Name,
			&c.TypeName,
			&c.MaxLength,
			&precision,
			&scale,
			&c.IsNullable,
			&c.IsIdentity,
			&c.IdentityNotForReplication,
			&c.SeedValue,
			&c.IncValue,
			&c.Collation,
			&c.IsComputed,
			&c.IsSparse,
			&c.IsFileStream,
		); err != nil {
			return nil, err
		}
		c.Precision = int(precision)
		c.Scale = int(scale)
		cols = append(cols, c)
	}
	return cols, rows.Err()
}

func loadDefaultConstraints(db *sql.DB, cfg AppConfig, objectID int) ([]defaultConstraintMeta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
SELECT dc.name, c.name AS column_name, dc.definition
FROM sys.default_constraints dc
JOIN sys.columns c
  ON c.object_id = dc.parent_object_id
 AND c.column_id = dc.parent_column_id
WHERE dc.parent_object_id = @obj
ORDER BY dc.name;
`, sql.Named("obj", objectID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]defaultConstraintMeta, 0)
	for rows.Next() {
		var d defaultConstraintMeta
		if err := rows.Scan(&d.Name, &d.ColumnName, &d.Definition); err != nil {
			return nil, err
		}
		out = append(out, d)
	}
	return out, rows.Err()
}

func loadCheckConstraints(db *sql.DB, cfg AppConfig, objectID int) ([]checkConstraintMeta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
SELECT cc.name, cc.definition, cc.is_not_for_replication, cc.is_disabled, cc.is_not_trusted
FROM sys.check_constraints cc
WHERE cc.parent_object_id = @obj
ORDER BY cc.name;
`, sql.Named("obj", objectID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]checkConstraintMeta, 0)
	for rows.Next() {
		var c checkConstraintMeta
		if err := rows.Scan(&c.Name, &c.Definition, &c.IsNotForReplication, &c.IsDisabled, &c.IsNotTrusted); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

func loadKeyConstraints(db *sql.DB, cfg AppConfig, objectID int) ([]keyConstraintMeta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
SELECT
  kc.name,
  kc.type,
  i.index_id,
  i.type_desc,
  i.fill_factor,
  i.is_padded,
  i.ignore_dup_key,
  i.allow_row_locks,
  i.allow_page_locks,
  i.is_disabled,
  ds.type_desc AS data_space_type,
  i.data_space_id,
  CASE WHEN NOT EXISTS (
    SELECT 1 FROM sys.partitions p2
    WHERE p2.object_id=i.object_id AND p2.index_id=i.index_id
      AND p2.data_compression_desc <> (
        SELECT TOP (1) p3.data_compression_desc FROM sys.partitions p3
        WHERE p3.object_id=i.object_id AND p3.index_id=i.index_id ORDER BY p3.partition_number
      )
  ) THEN (
    SELECT TOP (1) p3.data_compression_desc FROM sys.partitions p3
    WHERE p3.object_id=i.object_id AND p3.index_id=i.index_id ORDER BY p3.partition_number
  ) END AS compression_desc,
  CAST(CASE WHEN EXISTS (
    SELECT 1 FROM sys.partitions p2
    WHERE p2.object_id=i.object_id AND p2.index_id=i.index_id
      AND p2.data_compression_desc <> (
        SELECT TOP (1) p3.data_compression_desc FROM sys.partitions p3
        WHERE p3.object_id=i.object_id AND p3.index_id=i.index_id ORDER BY p3.partition_number
      )
  ) THEN 1 ELSE 0 END AS bit) AS has_mixed_compression
FROM sys.key_constraints kc
JOIN sys.indexes i
  ON i.object_id = kc.parent_object_id
 AND i.index_id = kc.unique_index_id
LEFT JOIN sys.data_spaces ds
  ON ds.data_space_id=i.data_space_id
WHERE kc.parent_object_id = @obj
ORDER BY
  CASE kc.type WHEN 'PK' THEN 0 ELSE 1 END,
  kc.name;
`, sql.Named("obj", objectID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]keyConstraintMeta, 0)
	for rows.Next() {
		var k keyConstraintMeta
		if err := rows.Scan(&k.Name, &k.Type, &k.IndexID, &k.IndexType, &k.FillFactor, &k.IsPadded, &k.IgnoreDupKey, &k.AllowRowLocks, &k.AllowPageLocks, &k.IsDisabled, &k.DataSpaceType, &k.DataSpaceID, &k.CompressionDesc, &k.HasMixedCompression); err != nil {
			return nil, err
		}
		out = append(out, k)
	}
	return out, rows.Err()
}

func loadKeyColumns(db *sql.DB, cfg AppConfig, objectID, indexID int) ([]indexColumnMeta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
SELECT c.name, ic.is_descending_key
FROM sys.index_columns ic
JOIN sys.columns c
  ON c.object_id = ic.object_id
 AND c.column_id = ic.column_id
WHERE ic.object_id = @obj
  AND ic.index_id = @idx
  AND ic.key_ordinal > 0
ORDER BY ic.key_ordinal;
`, sql.Named("obj", objectID), sql.Named("idx", indexID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]indexColumnMeta, 0)
	for rows.Next() {
		var c indexColumnMeta
		if err := rows.Scan(&c.ColumnName, &c.Desc); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

func loadForeignKeys(db *sql.DB, cfg AppConfig, objectID int) ([]foreignKeyMeta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
SELECT
  fk.object_id,
  fk.name,
  rs.name AS ref_schema,
  rt.name AS ref_table,
  fk.delete_referential_action_desc,
  fk.update_referential_action_desc,
  fk.is_not_for_replication,
  fk.is_disabled,
  fk.is_not_trusted
FROM sys.foreign_keys fk
JOIN sys.tables rt ON rt.object_id = fk.referenced_object_id
JOIN sys.schemas rs ON rs.schema_id = rt.schema_id
WHERE fk.parent_object_id = @obj
ORDER BY fk.name;
`, sql.Named("obj", objectID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]foreignKeyMeta, 0)
	for rows.Next() {
		var fk foreignKeyMeta
		if err := rows.Scan(&fk.ObjectID, &fk.Name, &fk.RefSchema, &fk.RefTable, &fk.DeleteAction, &fk.UpdateAction, &fk.IsNotForReplication, &fk.IsDisabled, &fk.IsNotTrusted); err != nil {
			return nil, err
		}
		out = append(out, fk)
	}
	return out, rows.Err()
}

func loadForeignKeyColumns(db *sql.DB, cfg AppConfig, fkObjectID int) ([]fkColumnMap, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
SELECT pc.name AS parent_col, rc.name AS ref_col
FROM sys.foreign_key_columns fkc
JOIN sys.columns pc
  ON pc.object_id = fkc.parent_object_id
 AND pc.column_id = fkc.parent_column_id
JOIN sys.columns rc
  ON rc.object_id = fkc.referenced_object_id
 AND rc.column_id = fkc.referenced_column_id
WHERE fkc.constraint_object_id = @fk
ORDER BY fkc.constraint_column_id;
`, sql.Named("fk", fkObjectID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]fkColumnMap, 0)
	for rows.Next() {
		var m fkColumnMap
		if err := rows.Scan(&m.ParentColumn, &m.RefColumn); err != nil {
			return nil, err
		}
		out = append(out, m)
	}
	return out, rows.Err()
}

func loadIndexes(db *sql.DB, cfg AppConfig, objectID int) ([]indexMeta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
SELECT
  i.name,
  i.index_id,
  i.type_desc,
  i.is_unique,
  i.filter_definition
FROM sys.indexes i
WHERE i.object_id = @obj
  AND i.is_hypothetical = 0
  AND i.name IS NOT NULL
  AND i.is_primary_key = 0
  AND i.is_unique_constraint = 0
ORDER BY i.name;
`, sql.Named("obj", objectID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]indexMeta, 0)
	for rows.Next() {
		var idx indexMeta
		if err := rows.Scan(&idx.Name, &idx.IndexID, &idx.TypeDesc, &idx.IsUnique, &idx.FilterDefinition); err != nil {
			return nil, err
		}
		out = append(out, idx)
	}
	return out, rows.Err()
}

func loadIndexColumns(db *sql.DB, cfg AppConfig, objectID, indexID int, includeColumns bool) ([]indexColumnMeta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	var includeBit int
	orderClause := "ORDER BY ic.key_ordinal"
	if includeColumns {
		includeBit = 1
		orderClause = "ORDER BY ic.index_column_id"
	}

	query := `
SELECT c.name, ic.is_descending_key
FROM sys.index_columns ic
JOIN sys.columns c
  ON c.object_id = ic.object_id
 AND c.column_id = ic.column_id
WHERE ic.object_id = @obj
  AND ic.index_id = @idx
  AND ic.is_included_column = @included
`
	query += orderClause + ";"

	rows, err := db.QueryContext(ctx, query,
		sql.Named("obj", objectID),
		sql.Named("idx", indexID),
		sql.Named("included", includeBit),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]indexColumnMeta, 0)
	for rows.Next() {
		var c indexColumnMeta
		if err := rows.Scan(&c.ColumnName, &c.Desc); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

func fileNameForObject(objectName string) string {
	return "dbo." + objectName + ".sql"
}

func normalizeSQL(input string) string {
	return sqlfmt.NormalizeSQLText(input, true, true)
}

func applySQLCategory(cfg AppConfig, relDir string, files map[string]string, dryRun bool) (CategoryChange, error) {
	return applySQLCategoryPreserving(cfg, relDir, files, nil, dryRun)
}

func applySQLCategoryPreserving(cfg AppConfig, relDir string, files map[string]string, preserve map[string]struct{}, dryRun bool) (CategoryChange, error) {
	change := CategoryChange{
		Dir:       relDir,
		Created:   []string{},
		Updated:   []string{},
		Deleted:   []string{},
		Unchanged: []string{},
	}

	dir := cfg.RepoPath(relDir)

	existing := map[string]string{}
	entries, readErr := os.ReadDir(dir)
	if readErr != nil && !errors.Is(readErr, os.ErrNotExist) {
		return change, readErr
	}
	if readErr == nil {
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			name := entry.Name()
			if !strings.EqualFold(filepath.Ext(name), ".sql") {
				continue
			}
			raw, err := os.ReadFile(filepath.Join(dir, name))
			if err != nil {
				return change, err
			}
			existing[name] = normalizeSQL(string(raw))
		}
	}

	allowed := make(map[string]struct{}, len(files))
	for name, raw := range files {
		allowed[name] = struct{}{}
		targetRel := filepath.ToSlash(filepath.Join(relDir, name))
		normalizedRaw := normalizeSQL(raw)

		existingRaw, exists := existing[name]
		if !exists {
			change.Created = append(change.Created, targetRel)
		} else if existingRaw != normalizedRaw {
			change.Updated = append(change.Updated, targetRel)
		} else {
			change.Unchanged = append(change.Unchanged, targetRel)
		}
	}
	for name := range preserve {
		allowed[name] = struct{}{}
	}

	for name := range existing {
		if _, ok := allowed[name]; ok {
			continue
		}
		change.Deleted = append(change.Deleted, filepath.ToSlash(filepath.Join(relDir, name)))
	}

	sort.Strings(change.Created)
	sort.Strings(change.Updated)
	sort.Strings(change.Unchanged)
	sort.Strings(change.Deleted)

	if dryRun {
		return change, nil
	}

	if err := os.MkdirAll(dir, 0o755); err != nil {
		return change, err
	}

	if err := writeFilesInParallel(dir, files, cfg.WriteWorkers); err != nil {
		return change, err
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return change, err
	}
	deleteNames := make([]string, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if !strings.EqualFold(filepath.Ext(name), ".sql") {
			continue
		}
		if _, ok := allowed[name]; ok {
			continue
		}
		deleteNames = append(deleteNames, name)
	}
	sort.Strings(deleteNames)
	if err := deleteFilesInParallel(dir, deleteNames, cfg.WriteWorkers); err != nil {
		return change, err
	}
	return change, nil
}

func validateObjectNameForRules(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New("object name is empty")
	}
	if strings.Contains(name, ".") {
		return errors.New("object name must not contain dot")
	}
	if strings.ContainsAny(name, " \t\r\n/\\") {
		return errors.New("object name must not contain spaces or path separators")
	}
	for _, r := range name {
		if r == '.' {
			return errors.New("object name must not contain dot")
		}
		if _, ok := migrationObjectTokenRunes[r]; ok {
			continue
		}
		return fmt.Errorf("object name contains unsupported character %q", r)
	}
	if name[0] >= '0' && name[0] <= '9' {
		return errors.New("object name must not start with digit")
	}
	return nil
}

func mustIntString(v sql.NullString, fallback int) int {
	if !v.Valid {
		return fallback
	}
	n, err := strconv.Atoi(strings.TrimSpace(v.String))
	if err != nil {
		return fallback
	}
	return n
}

func sortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func totalByCategory(items []CategoryChange, f func(CategoryChange) int) int {
	total := 0
	for _, it := range items {
		total += f(it)
	}
	return total
}

func normalizedWorkerCount(requested, total int) int {
	if total <= 0 {
		return 1
	}
	if requested <= 0 {
		requested = 1
	}
	if requested > total {
		return total
	}
	return requested
}

func writeFilesInParallel(dir string, files map[string]string, workers int) error {
	names := sortedKeys(files)
	if len(names) == 0 {
		return nil
	}
	workers = normalizedWorkerCount(workers, len(names))
	jobs := make(chan int, len(names))
	errs := make([]error, len(names))
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx := range jobs {
				name := names[idx]
				target := filepath.Join(dir, name)
				errs[idx] = atomicReplaceSQLFile(target, []byte(normalizeSQL(files[name])))
			}
		}()
	}

	for i := range names {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func atomicReplaceSQLFile(target string, data []byte) error {
	dir := filepath.Dir(target)
	tmp, err := os.CreateTemp(dir, "."+filepath.Base(target)+".tmp-*")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	keepTemp := true
	defer func() {
		_ = tmp.Close()
		if keepTemp {
			_ = os.Remove(tmpName)
		}
	}()
	if err := tmp.Chmod(0o644); err != nil {
		return err
	}
	if _, err := tmp.Write(data); err != nil {
		return err
	}
	if err := tmp.Sync(); err != nil {
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	if err := os.Rename(tmpName, target); err != nil {
		return fmt.Errorf("atomically replace %s: %w", filepath.Base(target), err)
	}
	keepTemp = false
	return nil
}

func deleteFilesInParallel(dir string, names []string, workers int) error {
	if len(names) == 0 {
		return nil
	}
	workers = normalizedWorkerCount(workers, len(names))
	jobs := make(chan int, len(names))
	errs := make([]error, len(names))
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx := range jobs {
				errs[idx] = os.Remove(filepath.Join(dir, names[idx]))
			}
		}()
	}

	for i := range names {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}
