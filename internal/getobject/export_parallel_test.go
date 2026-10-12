package getobject

import (
	"database/sql"
	"strings"
	"testing"
)

func TestNormalizedWorkerCount(t *testing.T) {
	t.Parallel()

	if got := normalizedWorkerCount(0, 10); got != 1 {
		t.Fatalf("expected default worker 1 for requested=0, got %d", got)
	}
	if got := normalizedWorkerCount(3, 2); got != 2 {
		t.Fatalf("expected worker capped by total, got %d", got)
	}
	if got := normalizedWorkerCount(4, 4); got != 4 {
		t.Fatalf("expected same worker count, got %d", got)
	}
}

func TestBuildTableAndIndexFilesParallel(t *testing.T) {
	t.Parallel()

	cfg := DefaultAppConfig(".")
	cfg.ConnectionString = "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable"
	cfg.RenderWorkers = 4

	tables := []tableMeta{
		{ObjectID: 1, Name: "TableA"},
		{ObjectID: 2, Name: "TableB"},
	}

	cache := tableSnapshotCache{
		Columns: map[int][]columnMeta{
			1: {
				{ColumnID: 1, Name: "id", TypeName: "int", IsNullable: false},
			},
			2: {
				{ColumnID: 1, Name: "id", TypeName: "int", IsNullable: false},
			},
		},
		Defaults:            map[int][]defaultConstraintMeta{},
		Checks:              map[int][]checkConstraintMeta{},
		Keys:                map[int][]keyConstraintMeta{},
		KeyColumns:          map[objectIndexKey][]indexColumnMeta{},
		ForeignKeys:         map[int][]foreignKeyMeta{},
		ForeignKeyColumns:   map[int][]fkColumnMap{},
		Indexes:             map[int][]indexMeta{},
		IndexIncludeColumns: map[objectIndexKey][]indexColumnMeta{},
	}
	cache.Indexes[1] = []indexMeta{
		{Name: "IX_TableA__id", IndexID: 2, TypeDesc: "NONCLUSTERED", IsUnique: false},
	}
	cache.KeyColumns[objectIndexKey{ObjectID: 1, IndexID: 2}] = []indexColumnMeta{
		{ColumnName: "id", Desc: false},
	}

	tableFiles, indexFiles, warnings, err := buildTableAndIndexFiles(cfg, tables, cache)
	if err != nil {
		t.Fatalf("buildTableAndIndexFiles failed: %v", err)
	}
	if len(warnings) != 0 {
		t.Fatalf("expected no table warnings, got %d", len(warnings))
	}
	if len(tableFiles) != 2 {
		t.Fatalf("expected 2 table files, got %d", len(tableFiles))
	}
	if len(indexFiles) != 1 {
		t.Fatalf("expected 1 index file, got %d", len(indexFiles))
	}
	if !strings.Contains(tableFiles["dbo.TableA.sql"], "CREATE TABLE dbo.TableA") {
		t.Fatalf("unexpected table SQL for TableA")
	}
	if !strings.Contains(indexFiles["dbo.TableA.sql"], "CREATE NONCLUSTERED INDEX [IX_TableA__id] ON dbo.TableA") {
		t.Fatalf("unexpected index SQL for TableA")
	}
}

func TestBuildTableAndIndexFilesSkipUnsupportedTable(t *testing.T) {
	t.Parallel()

	cfg := DefaultAppConfig(".")
	cfg.ConnectionString = "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable"
	cfg.RenderWorkers = 2
	cfg.SkipUnsupportedTables = true

	tables := []tableMeta{
		{ObjectID: 1, Name: "TableA"},
	}
	cache := tableSnapshotCache{
		Columns: map[int][]columnMeta{
			1: {
				{ColumnID: 1, Name: "payload", TypeName: "varbinary", MaxLength: -1, IsNullable: true, IsFileStream: true},
			},
		},
		Defaults:            map[int][]defaultConstraintMeta{},
		Checks:              map[int][]checkConstraintMeta{},
		Keys:                map[int][]keyConstraintMeta{},
		KeyColumns:          map[objectIndexKey][]indexColumnMeta{},
		ForeignKeys:         map[int][]foreignKeyMeta{},
		ForeignKeyColumns:   map[int][]fkColumnMap{},
		Indexes:             map[int][]indexMeta{},
		IndexIncludeColumns: map[objectIndexKey][]indexColumnMeta{},
	}

	tableFiles, indexFiles, warnings, err := buildTableAndIndexFiles(cfg, tables, cache)
	if err != nil {
		t.Fatalf("buildTableAndIndexFiles failed: %v", err)
	}
	if len(tableFiles) != 0 || len(indexFiles) != 0 {
		t.Fatalf("expected unsupported table to be skipped")
	}
	if len(warnings) != 1 {
		t.Fatalf("expected 1 warning, got %d", len(warnings))
	}
	if warnings[0].Rule != "TABLE_EXTRACT_SKIPPED_UNSUPPORTED" {
		t.Fatalf("unexpected warning rule: %s", warnings[0].Rule)
	}
	if warnings[0].Object != "dbo.TableA" {
		t.Fatalf("unexpected warning object: %s", warnings[0].Object)
	}
}

func TestBuildTableAndIndexFilesRendersComputedAndRowGuidColumns(t *testing.T) {
	cfg := DefaultAppConfig(t.TempDir())
	cfg.Schema = "dbo"
	tables := []tableMeta{{ObjectID: 1, Name: "TableA"}}
	cache := tableSnapshotCache{Columns: map[int][]columnMeta{1: {
		{ColumnID: 1, Name: "base", TypeName: "int", IsNullable: false},
		{ColumnID: 2, Name: "calc", TypeName: "int", IsComputed: true, ComputedDefinition: sql.NullString{String: "([base] * (2))", Valid: true}},
		{ColumnID: 3, Name: "guid", TypeName: "uniqueidentifier", IsNullable: true, IsRowGuidCol: true},
	}}, Defaults: map[int][]defaultConstraintMeta{}, Checks: map[int][]checkConstraintMeta{}, Keys: map[int][]keyConstraintMeta{}, KeyColumns: map[objectIndexKey][]indexColumnMeta{}, ForeignKeys: map[int][]foreignKeyMeta{}, ForeignKeyColumns: map[int][]fkColumnMap{}, Indexes: map[int][]indexMeta{}, IndexIncludeColumns: map[objectIndexKey][]indexColumnMeta{}}
	files, _, warnings, err := buildTableAndIndexFiles(cfg, tables, cache)
	if err != nil || len(warnings) != 0 {
		t.Fatalf("build failed: err=%v warnings=%v", err, warnings)
	}
	sql := files["dbo.TableA.sql"]
	if !strings.Contains(sql, "[calc] AS ([base] * (2))") || !strings.Contains(sql, "[guid] UNIQUEIDENTIFIER ROWGUIDCOL NULL") {
		t.Fatalf("missing supported modifier:\n%s", sql)
	}
}

func TestBuildTableAndIndexFilesUnsupportedTableFailsWhenSkipDisabled(t *testing.T) {
	t.Parallel()

	cfg := DefaultAppConfig(".")
	cfg.ConnectionString = "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable"
	cfg.SkipUnsupportedTables = false

	tables := []tableMeta{
		{ObjectID: 1, Name: "TableA"},
	}
	cache := tableSnapshotCache{
		Columns: map[int][]columnMeta{
			1: {
				{ColumnID: 1, Name: "id", TypeName: "int", IsNullable: false, IsComputed: true},
			},
		},
		Defaults:            map[int][]defaultConstraintMeta{},
		Checks:              map[int][]checkConstraintMeta{},
		Keys:                map[int][]keyConstraintMeta{},
		KeyColumns:          map[objectIndexKey][]indexColumnMeta{},
		ForeignKeys:         map[int][]foreignKeyMeta{},
		ForeignKeyColumns:   map[int][]fkColumnMap{},
		Indexes:             map[int][]indexMeta{},
		IndexIncludeColumns: map[objectIndexKey][]indexColumnMeta{},
	}

	_, _, _, err := buildTableAndIndexFiles(cfg, tables, cache)
	if err == nil {
		t.Fatalf("expected unsupported table feature to fail when skip_unsupported_tables=false")
	}
}

func TestBuildTableSnapshotSQLFromCache_DefaultConstraintRenderedInline(t *testing.T) {
	t.Parallel()

	cfg := DefaultAppConfig(".")
	table := tableMeta{ObjectID: 99, Name: "TableA"}
	cache := tableSnapshotCache{
		Columns: map[int][]columnMeta{
			99: {
				{ColumnID: 1, Name: "id", TypeName: "int", IsNullable: false},
				{ColumnID: 2, Name: "status", TypeName: "varchar", MaxLength: 10, IsNullable: true},
			},
		},
		Defaults: map[int][]defaultConstraintMeta{
			99: {
				{Name: "DF_TableA__status", ColumnName: "status", Definition: "('NEW')"},
			},
		},
		Checks:              map[int][]checkConstraintMeta{},
		Keys:                map[int][]keyConstraintMeta{},
		KeyColumns:          map[objectIndexKey][]indexColumnMeta{},
		ForeignKeys:         map[int][]foreignKeyMeta{},
		ForeignKeyColumns:   map[int][]fkColumnMap{},
		Indexes:             map[int][]indexMeta{},
		IndexIncludeColumns: map[objectIndexKey][]indexColumnMeta{},
	}

	sql, err := buildTableSnapshotSQLFromCache(cfg, table, cache)
	if err != nil {
		t.Fatalf("buildTableSnapshotSQLFromCache failed: %v", err)
	}
	if !strings.Contains(sql, "[status] VARCHAR(10) NULL CONSTRAINT [DF_TableA__status] DEFAULT ('NEW')") {
		t.Fatalf("expected inline default constraint in column definition, got:\n%s", sql)
	}
	if strings.Contains(strings.ToUpper(sql), "DEFAULT ('NEW') FOR [STATUS]") {
		t.Fatalf("did not expect table-level FOR default constraint, got:\n%s", sql)
	}
}

func TestBuildTableSnapshotSQLFromCache_RendersSparse(t *testing.T) {
	cfg := DefaultAppConfig(t.TempDir())
	table := tableMeta{ObjectID: 1, Name: "SparseExample"}
	cache := tableSnapshotCache{
		Columns: map[int][]columnMeta{1: {
			{ColumnID: 1, Name: "Id", TypeName: "int", IsNullable: false},
			{ColumnID: 2, Name: "OptionalCode", TypeName: "varchar", MaxLength: 20, IsNullable: true, IsSparse: true},
		}},
		Defaults: map[int][]defaultConstraintMeta{}, Checks: map[int][]checkConstraintMeta{}, Keys: map[int][]keyConstraintMeta{}, KeyColumns: map[objectIndexKey][]indexColumnMeta{}, ForeignKeys: map[int][]foreignKeyMeta{}, ForeignKeyColumns: map[int][]fkColumnMap{},
	}
	sqlText, err := buildTableSnapshotSQLFromCache(cfg, table, cache)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(strings.ToUpper(sqlText), "[OPTIONALCODE] VARCHAR(20) SPARSE NULL") {
		t.Fatalf("SPARSE modifier missing from extractor output:\n%s", sqlText)
	}
}

func TestBuildTableSnapshotSQLFromCacheRejectsLossyCatalogStates(t *testing.T) {
	t.Parallel()

	cfg := DefaultAppConfig(t.TempDir())
	table := tableMeta{ObjectID: 7, Name: "ExactCaseTable"}
	base := func() tableSnapshotCache {
		return tableSnapshotCache{
			Columns:             map[int][]columnMeta{7: {{ColumnID: 1, Name: "ExactCase", TypeName: "int"}}},
			Defaults:            map[int][]defaultConstraintMeta{},
			Checks:              map[int][]checkConstraintMeta{},
			Keys:                map[int][]keyConstraintMeta{},
			KeyColumns:          map[objectIndexKey][]indexColumnMeta{},
			ForeignKeys:         map[int][]foreignKeyMeta{},
			ForeignKeyColumns:   map[int][]fkColumnMap{},
			Indexes:             map[int][]indexMeta{},
			IndexIncludeColumns: map[objectIndexKey][]indexColumnMeta{},
			UnsupportedIndexes:  map[int][]string{},
		}
	}

	tests := []struct {
		name   string
		mutate func(*tableSnapshotCache)
	}{
		{name: "identity not for replication", mutate: func(c *tableSnapshotCache) { c.Columns[7][0].IdentityNotForReplication = true }},
		{name: "advanced index catalog", mutate: func(c *tableSnapshotCache) { c.UnsupportedIndexes[7] = []string{"FULLTEXT:(FULLTEXT)"} }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := base()
			tt.mutate(&cache)
			if _, err := buildTableSnapshotSQLFromCache(cfg, table, cache); err == nil {
				t.Fatal("lossy catalog state was silently accepted")
			}
		})
	}
}

func TestBuildTableAndIndexFilesExportsDisabledUntrustedConstraintsWithHealthWarnings(t *testing.T) {
	t.Parallel()

	cfg := DefaultAppConfig(t.TempDir())
	table := tableMeta{ObjectID: 7, Name: "ConstraintHealth"}
	cache := tableSnapshotCache{
		Columns: map[int][]columnMeta{7: {
			{ColumnID: 1, Name: "Id", TypeName: "int", IsNullable: false},
			{ColumnID: 2, Name: "ParentId", TypeName: "int", IsNullable: true},
		}},
		Defaults: map[int][]defaultConstraintMeta{},
		Checks: map[int][]checkConstraintMeta{7: {
			{Name: "CK_ConstraintHealth_Id", Definition: "([Id]>(0))", IsNotTrusted: true},
		}},
		Keys: map[int][]keyConstraintMeta{7: {
			{Name: "PK_ConstraintHealth", Type: "PK", IndexID: 1, IndexType: "CLUSTERED", FillFactor: 80, AllowRowLocks: true, AllowPageLocks: true},
		}},
		KeyColumns: map[objectIndexKey][]indexColumnMeta{
			{ObjectID: 7, IndexID: 1}: {{ColumnName: "Id"}},
		},
		ForeignKeys: map[int][]foreignKeyMeta{7: {
			{ObjectID: 70, Name: "FK_ConstraintHealth_Parent", RefSchema: "dbo", RefTable: "Parent", IsDisabled: true, IsNotTrusted: true},
		}},
		ForeignKeyColumns: map[int][]fkColumnMap{
			70: {{ParentColumn: "ParentId", RefColumn: "Id"}},
		},
		Indexes:             map[int][]indexMeta{},
		IndexIncludeColumns: map[objectIndexKey][]indexColumnMeta{},
		UnsupportedIndexes:  map[int][]string{},
		UnsupportedTables:   map[int][]string{},
	}

	tableFiles, _, warnings, err := buildTableAndIndexFiles(cfg, []tableMeta{table}, cache)
	if err != nil {
		t.Fatalf("logical constraint export failed: %v", err)
	}
	rendered := tableFiles["dbo.ConstraintHealth.sql"]
	for _, token := range []string{
		"CONSTRAINT [CK_ConstraintHealth_Id] CHECK ([Id]>(0))",
		"CONSTRAINT [PK_ConstraintHealth] PRIMARY KEY CLUSTERED ([Id] ASC)",
		"CONSTRAINT [FK_ConstraintHealth_Parent] FOREIGN KEY ([ParentId]) REFERENCES dbo.Parent ([Id])",
	} {
		if !strings.Contains(rendered, token) {
			t.Fatalf("logical table snapshot missing %q:\n%s", token, rendered)
		}
	}
	if len(warnings) != 3 {
		t.Fatalf("expected check, key physical metadata, and FK findings; got %#v", warnings)
	}
	if warnings[0].Rule != "CONSTRAINT_STATE_REVIEW_REQUIRED" ||
		warnings[1].Rule != "KEY_PHYSICAL_METADATA_NOT_CAPTURED" ||
		warnings[2].Rule != "CONSTRAINT_STATE_REVIEW_REQUIRED" {
		t.Fatalf("unexpected warning order/content: %#v", warnings)
	}
	if warnings[2].Enabled == nil || *warnings[2].Enabled || warnings[2].Trusted == nil || *warnings[2].Trusted {
		t.Fatalf("disabled/untrusted FK state was not captured structurally: %#v", warnings[2])
	}
}

func TestHealthWarningsDoNotPreserveStaleIndexSnapshots(t *testing.T) {
	t.Parallel()

	health := ExportWarning{
		Rule:     "CONSTRAINT_STATE_REVIEW_REQUIRED",
		Category: "tables",
		Object:   "dbo.TableA",
		Message:  "health finding",
	}
	if preserved := preserveNamesFromTableWarnings([]ExportWarning{health}); len(preserved) != 0 {
		t.Fatalf("health warning must not preserve stale files: %#v", preserved)
	}
	skipped := health
	skipped.Rule = "TABLE_EXTRACT_SKIPPED_UNSUPPORTED"
	if preserved := preserveNamesFromTableWarnings([]ExportWarning{skipped}); len(preserved) != 1 {
		t.Fatalf("actual skipped table must preserve prior files: %#v", preserved)
	}
}
