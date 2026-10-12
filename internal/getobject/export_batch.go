package getobject

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

type objectIndexKey struct {
	ObjectID int
	IndexID  int
}

type tableSnapshotCache struct {
	Columns             map[int][]columnMeta
	Defaults            map[int][]defaultConstraintMeta
	Checks              map[int][]checkConstraintMeta
	Keys                map[int][]keyConstraintMeta
	KeyColumns          map[objectIndexKey][]indexColumnMeta
	ForeignKeys         map[int][]foreignKeyMeta
	ForeignKeyColumns   map[int][]fkColumnMap
	Indexes             map[int][]indexMeta
	IndexIncludeColumns map[objectIndexKey][]indexColumnMeta
	UnsupportedIndexes  map[int][]string
	UnsupportedTables   map[int][]string
}

func loadTableSnapshotCache(db *sql.DB, cfg AppConfig) (tableSnapshotCache, error) {
	columns, err := loadColumnsByTable(db, cfg)
	if err != nil {
		return tableSnapshotCache{}, err
	}
	defaults, err := loadDefaultConstraintsByTable(db, cfg)
	if err != nil {
		return tableSnapshotCache{}, err
	}
	checks, err := loadCheckConstraintsByTable(db, cfg)
	if err != nil {
		return tableSnapshotCache{}, err
	}
	keys, err := loadKeyConstraintsByTable(db, cfg)
	if err != nil {
		return tableSnapshotCache{}, err
	}
	keyCols, err := loadKeyColumnsByTableAndIndex(db, cfg)
	if err != nil {
		return tableSnapshotCache{}, err
	}
	fks, err := loadForeignKeysByTable(db, cfg)
	if err != nil {
		return tableSnapshotCache{}, err
	}
	fkCols, err := loadForeignKeyColumnsByFK(db, cfg)
	if err != nil {
		return tableSnapshotCache{}, err
	}
	indexes, err := loadIndexesByTable(db, cfg)
	if err != nil {
		return tableSnapshotCache{}, err
	}
	includeCols, err := loadIndexIncludeColumnsByTableAndIndex(db, cfg)
	if err != nil {
		return tableSnapshotCache{}, err
	}
	unsupportedIndexes, err := loadUnsupportedIndexFeaturesByTable(db, cfg)
	if err != nil {
		return tableSnapshotCache{}, err
	}
	unsupportedTables, err := loadUnsupportedTableFeaturesByTable(db, cfg)
	if err != nil {
		return tableSnapshotCache{}, err
	}

	return tableSnapshotCache{
		Columns:             columns,
		Defaults:            defaults,
		Checks:              checks,
		Keys:                keys,
		KeyColumns:          keyCols,
		ForeignKeys:         fks,
		ForeignKeyColumns:   fkCols,
		Indexes:             indexes,
		IndexIncludeColumns: includeCols,
		UnsupportedIndexes:  unsupportedIndexes,
		UnsupportedTables:   unsupportedTables,
	}, nil
}

func buildTableSnapshotSQLFromCache(cfg AppConfig, t tableMeta, cache tableSnapshotCache) (string, error) {
	if features := cache.UnsupportedTables[t.ObjectID]; len(features) > 0 {
		return "", fmt.Errorf("table %s has unsupported table metadata: %s", t.Name, strings.Join(features, ", "))
	}
	if features := cache.UnsupportedIndexes[t.ObjectID]; len(features) > 0 {
		return "", fmt.Errorf("table %s has unsupported index metadata: %s", t.Name, strings.Join(features, ", "))
	}
	cols := cache.Columns[t.ObjectID]
	if len(cols) == 0 {
		return "", errors.New("table has zero columns")
	}

	defaultByColumn, err := defaultConstraintMapByColumn(cache.Defaults[t.ObjectID])
	if err != nil {
		return "", err
	}

	items := make([]string, 0, len(cols))
	for _, c := range cols {
		if c.IsComputed && (!c.ComputedDefinition.Valid || strings.TrimSpace(c.ComputedDefinition.String) == "") {
			return "", fmt.Errorf("computed column %s has empty definition", c.Name)
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
	for _, ck := range cache.Checks[t.ObjectID] {
		if strings.TrimSpace(ck.Name) == "" {
			return "", fmt.Errorf("check constraint name must not be empty on table %s", t.Name)
		}
		items = append(items, renderCheckConstraintClause(ck))
	}
	for _, kc := range cache.Keys[t.ObjectID] {
		if strings.TrimSpace(kc.Name) == "" {
			return "", fmt.Errorf("key constraint name must not be empty on table %s", t.Name)
		}
		if err := validateKeyConstraintMetadata(kc); err != nil {
			return "", err
		}
		keyCols := cache.KeyColumns[objectIndexKey{ObjectID: t.ObjectID, IndexID: kc.IndexID}]
		if len(keyCols) == 0 {
			return "", fmt.Errorf("key constraint %s has no key columns", kc.Name)
		}
		items = append(items, renderKeyConstraintClause(kc, keyCols))
	}

	for _, fk := range cache.ForeignKeys[t.ObjectID] {
		if strings.TrimSpace(fk.Name) == "" {
			return "", fmt.Errorf("foreign key name must not be empty on table %s", t.Name)
		}
		if strings.TrimSpace(fk.RefSchema) != strings.TrimSpace(cfg.Schema) {
			return "", fmt.Errorf("foreign key %s references non-dbo schema %s", fk.Name, fk.RefSchema)
		}
		if err := validateObjectNameForRules(fk.RefTable); err != nil {
			return "", fmt.Errorf("foreign key %s references invalid table %s: %w", fk.Name, fk.RefTable, err)
		}
		colMap := cache.ForeignKeyColumns[fk.ObjectID]
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

func loadUnsupportedTableFeaturesByTable(db *sql.DB, cfg AppConfig) (map[int][]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
SET NOCOUNT ON;
CREATE TABLE #rza_unsupported_tables (
  object_id INT NOT NULL,
  feature_name NVARCHAR(128) NOT NULL
);

INSERT INTO #rza_unsupported_tables
SELECT DISTINCT t.object_id, N'USER_DEFINED_OR_ALIAS_TYPE:' + ts.name + N'.' + ty.name
FROM sys.tables t
JOIN sys.schemas s ON s.schema_id=t.schema_id
JOIN sys.columns c ON c.object_id=t.object_id
JOIN sys.types ty ON ty.user_type_id=c.user_type_id
JOIN sys.schemas ts ON ts.schema_id=ty.schema_id
WHERE s.name=@schema AND t.is_ms_shipped=0 AND ty.is_user_defined=1;

INSERT INTO #rza_unsupported_tables
SELECT DISTINCT t.object_id, N'TYPED_XML_SCHEMA_COLLECTION'
FROM sys.tables t
JOIN sys.schemas s ON s.schema_id=t.schema_id
JOIN sys.columns c ON c.object_id=t.object_id
WHERE s.name=@schema AND t.is_ms_shipped=0 AND c.xml_collection_id<>0;

INSERT INTO #rza_unsupported_tables
SELECT DISTINCT t.object_id, N'COLUMN_SET_BOUND_RULE_OR_BOUND_DEFAULT'
FROM sys.tables t
JOIN sys.schemas s ON s.schema_id=t.schema_id
JOIN sys.columns c ON c.object_id=t.object_id
WHERE s.name=@schema AND t.is_ms_shipped=0 AND (
  c.is_column_set=1 OR c.rule_object_id<>0 OR
  (c.default_object_id<>0 AND NOT EXISTS (
    SELECT 1 FROM sys.default_constraints dc
    WHERE dc.object_id=c.default_object_id AND dc.parent_object_id=c.object_id AND dc.parent_column_id=c.column_id
  ))
);

INSERT INTO #rza_unsupported_tables
SELECT DISTINCT t.object_id, N'TABLE_TRIGGER:' + tr.name
FROM sys.tables t
JOIN sys.schemas s ON s.schema_id=t.schema_id
JOIN sys.triggers tr ON tr.parent_id=t.object_id AND tr.is_ms_shipped=0
WHERE s.name=@schema AND t.is_ms_shipped=0;

INSERT INTO #rza_unsupported_tables
SELECT DISTINCT t.object_id, N'USER_CREATED_STATISTICS:' + st.name
FROM sys.tables t
JOIN sys.schemas s ON s.schema_id=t.schema_id
JOIN sys.stats st ON st.object_id=t.object_id AND st.user_created=1
WHERE s.name=@schema AND t.is_ms_shipped=0
  AND NOT EXISTS (
    SELECT 1 FROM sys.indexes i
    WHERE i.object_id=st.object_id AND i.index_id=st.stats_id
  );

INSERT INTO #rza_unsupported_tables
SELECT DISTINCT t.object_id, N'REPLICATION_OR_CDC_TABLE'
FROM sys.tables t
JOIN sys.schemas s ON s.schema_id=t.schema_id
WHERE s.name=@schema AND t.is_ms_shipped=0
  AND (t.is_replicated=1 OR t.is_merge_published=1 OR t.is_sync_tran_subscribed=1 OR t.is_tracked_by_cdc=1);

IF OBJECT_ID(N'sys.change_tracking_tables') IS NOT NULL
  EXEC sys.sp_executesql N'INSERT INTO #rza_unsupported_tables SELECT t.object_id,N''CHANGE_TRACKING_TABLE'' FROM sys.tables t JOIN sys.schemas s ON s.schema_id=t.schema_id JOIN sys.change_tracking_tables ct ON ct.object_id=t.object_id WHERE s.name=@p AND t.is_ms_shipped=0', N'@p sysname', @p=@schema;

INSERT INTO #rza_unsupported_tables
SELECT DISTINCT t.object_id, N'HEAP_DATA_COMPRESSION'
FROM sys.tables t
JOIN sys.schemas s ON s.schema_id=t.schema_id
JOIN sys.partitions p ON p.object_id=t.object_id AND p.index_id=0
WHERE s.name=@schema AND t.is_ms_shipped=0 AND p.data_compression_desc<>N'NONE';

IF EXISTS (SELECT 1 FROM sys.all_columns WHERE object_id=OBJECT_ID(N'sys.tables') AND name=N'is_memory_optimized')
  EXEC sys.sp_executesql N'INSERT INTO #rza_unsupported_tables SELECT t.object_id,N''MEMORY_OPTIMIZED_TABLE'' FROM sys.tables t JOIN sys.schemas s ON s.schema_id=t.schema_id WHERE s.name=@p AND t.is_ms_shipped=0 AND t.is_memory_optimized=1', N'@p sysname', @p=@schema;
IF EXISTS (SELECT 1 FROM sys.all_columns WHERE object_id=OBJECT_ID(N'sys.tables') AND name=N'temporal_type')
  EXEC sys.sp_executesql N'INSERT INTO #rza_unsupported_tables SELECT t.object_id,N''TEMPORAL_OR_HISTORY_TABLE'' FROM sys.tables t JOIN sys.schemas s ON s.schema_id=t.schema_id WHERE s.name=@p AND t.is_ms_shipped=0 AND t.temporal_type<>0', N'@p sysname', @p=@schema;
IF EXISTS (SELECT 1 FROM sys.all_columns WHERE object_id=OBJECT_ID(N'sys.tables') AND name=N'is_node')
  EXEC sys.sp_executesql N'INSERT INTO #rza_unsupported_tables SELECT t.object_id,N''GRAPH_NODE_OR_EDGE_TABLE'' FROM sys.tables t JOIN sys.schemas s ON s.schema_id=t.schema_id WHERE s.name=@p AND t.is_ms_shipped=0 AND (t.is_node=1 OR t.is_edge=1)', N'@p sysname', @p=@schema;
IF EXISTS (SELECT 1 FROM sys.all_columns WHERE object_id=OBJECT_ID(N'sys.tables') AND name=N'ledger_type')
  EXEC sys.sp_executesql N'INSERT INTO #rza_unsupported_tables SELECT t.object_id,N''LEDGER_TABLE'' FROM sys.tables t JOIN sys.schemas s ON s.schema_id=t.schema_id WHERE s.name=@p AND t.is_ms_shipped=0 AND t.ledger_type<>0', N'@p sysname', @p=@schema;
IF EXISTS (SELECT 1 FROM sys.all_columns WHERE object_id=OBJECT_ID(N'sys.columns') AND name=N'is_hidden')
  EXEC sys.sp_executesql N'INSERT INTO #rza_unsupported_tables SELECT DISTINCT t.object_id,N''HIDDEN_OR_GENERATED_COLUMN'' FROM sys.tables t JOIN sys.schemas s ON s.schema_id=t.schema_id JOIN sys.columns c ON c.object_id=t.object_id WHERE s.name=@p AND t.is_ms_shipped=0 AND (c.is_hidden=1 OR c.generated_always_type<>0)', N'@p sysname', @p=@schema;
IF EXISTS (SELECT 1 FROM sys.all_columns WHERE object_id=OBJECT_ID(N'sys.columns') AND name=N'is_masked')
  EXEC sys.sp_executesql N'INSERT INTO #rza_unsupported_tables SELECT DISTINCT t.object_id,N''MASKED_COLUMN'' FROM sys.tables t JOIN sys.schemas s ON s.schema_id=t.schema_id JOIN sys.columns c ON c.object_id=t.object_id WHERE s.name=@p AND t.is_ms_shipped=0 AND c.is_masked=1', N'@p sysname', @p=@schema;
IF EXISTS (SELECT 1 FROM sys.all_columns WHERE object_id=OBJECT_ID(N'sys.columns') AND name=N'encryption_type')
  EXEC sys.sp_executesql N'INSERT INTO #rza_unsupported_tables SELECT DISTINCT t.object_id,N''ENCRYPTED_COLUMN'' FROM sys.tables t JOIN sys.schemas s ON s.schema_id=t.schema_id JOIN sys.columns c ON c.object_id=t.object_id WHERE s.name=@p AND t.is_ms_shipped=0 AND c.encryption_type IS NOT NULL', N'@p sysname', @p=@schema;
IF EXISTS (SELECT 1 FROM sys.all_columns WHERE object_id=OBJECT_ID(N'sys.columns') AND name=N'vector_dimensions')
  EXEC sys.sp_executesql N'INSERT INTO #rza_unsupported_tables SELECT DISTINCT t.object_id,N''VECTOR_COLUMN'' FROM sys.tables t JOIN sys.schemas s ON s.schema_id=t.schema_id JOIN sys.columns c ON c.object_id=t.object_id WHERE s.name=@p AND t.is_ms_shipped=0 AND c.vector_dimensions IS NOT NULL', N'@p sysname', @p=@schema;

SELECT DISTINCT object_id, feature_name
FROM #rza_unsupported_tables
ORDER BY object_id, feature_name;
`, sql.Named("schema", cfg.Schema))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := map[int][]string{}
	for rows.Next() {
		var objectID int
		var feature string
		if err := rows.Scan(&objectID, &feature); err != nil {
			return nil, err
		}
		out[objectID] = append(out[objectID], feature)
	}
	return out, rows.Err()
}

func buildIndexSnapshotSQLFromCache(cfg AppConfig, t tableMeta, cache tableSnapshotCache) (string, error) {
	if features := cache.UnsupportedIndexes[t.ObjectID]; len(features) > 0 {
		return "", fmt.Errorf("table %s has unsupported index metadata: %s", t.Name, strings.Join(features, ", "))
	}
	indexes := cache.Indexes[t.ObjectID]
	if len(indexes) == 0 {
		return "", nil
	}

	stmts := make([]string, 0, len(indexes))
	for _, idx := range indexes {
		key := objectIndexKey{ObjectID: t.ObjectID, IndexID: idx.IndexID}
		keyCols := cache.KeyColumns[key]
		if len(keyCols) == 0 && !strings.EqualFold(idx.TypeDesc, "CLUSTERED COLUMNSTORE") {
			return "", fmt.Errorf("index %s has zero key columns", idx.Name)
		}
		includeCols := cache.IndexIncludeColumns[key]

		stmt, err := renderCreateIndexClause(cfg.Schema, t.Name, idx, keyCols, includeCols)
		if err != nil {
			return "", err
		}
		stmts = append(stmts, stmt)
	}

	return strings.Join(stmts, "\n") + "\n", nil
}

func loadUnsupportedIndexFeaturesByTable(db *sql.DB, cfg AppConfig) (map[int][]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
SET NOCOUNT ON;
CREATE TABLE #rza_unsupported_indexes (
  object_id INT NOT NULL,
  index_name NVARCHAR(128) NOT NULL,
  feature_name NVARCHAR(64) NOT NULL
);

INSERT INTO #rza_unsupported_indexes (object_id, index_name, feature_name)
SELECT i.object_id, i.name, N'MEMORY_OPTIMIZED_HASH'
FROM sys.indexes i
JOIN sys.tables t ON t.object_id=i.object_id
JOIN sys.schemas s ON s.schema_id=t.schema_id
WHERE s.name=@schema AND t.is_ms_shipped=0 AND i.type_desc=N'NONCLUSTERED HASH';

INSERT INTO #rza_unsupported_indexes (object_id, index_name, feature_name)
SELECT fi.object_id, N'(FULLTEXT)', N'FULLTEXT'
FROM sys.fulltext_indexes fi
JOIN sys.tables t ON t.object_id=fi.object_id
JOIN sys.schemas s ON s.schema_id=t.schema_id
WHERE s.name=@schema AND t.is_ms_shipped=0;

IF EXISTS (SELECT 1 FROM sys.all_columns WHERE object_id=OBJECT_ID(N'sys.indexes') AND name=N'optimize_for_sequential_key')
BEGIN
  EXEC sys.sp_executesql N'
    INSERT INTO #rza_unsupported_indexes (object_id, index_name, feature_name)
    SELECT i.object_id, i.name, N''OPTIMIZE_FOR_SEQUENTIAL_KEY''
    FROM sys.indexes i
    JOIN sys.tables t ON t.object_id=i.object_id
    JOIN sys.schemas s ON s.schema_id=t.schema_id
    WHERE s.name=@inner_schema AND t.is_ms_shipped=0 AND i.optimize_for_sequential_key=1;',
    N'@inner_schema sysname', @inner_schema=@schema;
END;

IF EXISTS (SELECT 1 FROM sys.all_columns WHERE object_id=OBJECT_ID(N'sys.index_columns') AND name=N'column_store_order_ordinal')
BEGIN
  EXEC sys.sp_executesql N'
    INSERT INTO #rza_unsupported_indexes (object_id, index_name, feature_name)
    SELECT DISTINCT i.object_id, i.name, N''ORDERED_COLUMNSTORE''
    FROM sys.indexes i
    JOIN sys.index_columns ic ON ic.object_id=i.object_id AND ic.index_id=i.index_id
    JOIN sys.tables t ON t.object_id=i.object_id
    JOIN sys.schemas s ON s.schema_id=t.schema_id
    WHERE s.name=@inner_schema AND t.is_ms_shipped=0 AND ic.column_store_order_ordinal>0;',
    N'@inner_schema sysname', @inner_schema=@schema;
END;

IF EXISTS (SELECT 1 FROM sys.all_columns WHERE object_id=OBJECT_ID(N'sys.index_columns') AND name=N'data_clustering_ordinal')
BEGIN
  EXEC sys.sp_executesql N'
    INSERT INTO #rza_unsupported_indexes (object_id, index_name, feature_name)
    SELECT DISTINCT i.object_id, i.name, N''DATA_CLUSTERED_COLUMNSTORE''
    FROM sys.indexes i
    JOIN sys.index_columns ic ON ic.object_id=i.object_id AND ic.index_id=i.index_id
    JOIN sys.tables t ON t.object_id=i.object_id
    JOIN sys.schemas s ON s.schema_id=t.schema_id
    WHERE s.name=@inner_schema AND t.is_ms_shipped=0 AND ic.data_clustering_ordinal>0;',
    N'@inner_schema sysname', @inner_schema=@schema;
END;

IF EXISTS (SELECT 1 FROM sys.all_columns WHERE object_id=OBJECT_ID(N'sys.indexes') AND name=N'auto_created')
  EXEC sys.sp_executesql N'INSERT INTO #rza_unsupported_indexes SELECT i.object_id,i.name,N''AUTO_CREATED_INDEX'' FROM sys.indexes i JOIN sys.tables t ON t.object_id=i.object_id JOIN sys.schemas s ON s.schema_id=t.schema_id WHERE s.name=@p AND t.is_ms_shipped=0 AND i.auto_created=1', N'@p sysname', @p=@schema;
IF EXISTS (SELECT 1 FROM sys.all_columns WHERE object_id=OBJECT_ID(N'sys.indexes') AND name=N'suppress_dup_key_messages')
  EXEC sys.sp_executesql N'INSERT INTO #rza_unsupported_indexes SELECT i.object_id,i.name,N''SUPPRESS_DUP_KEY_MESSAGES'' FROM sys.indexes i JOIN sys.tables t ON t.object_id=i.object_id JOIN sys.schemas s ON s.schema_id=t.schema_id WHERE s.name=@p AND t.is_ms_shipped=0 AND i.suppress_dup_key_messages=1', N'@p sysname', @p=@schema;
IF EXISTS (SELECT 1 FROM sys.all_columns WHERE object_id=OBJECT_ID(N'sys.indexes') AND name=N'compression_delay')
  EXEC sys.sp_executesql N'INSERT INTO #rza_unsupported_indexes SELECT i.object_id,i.name,N''COLUMNSTORE_COMPRESSION_DELAY'' FROM sys.indexes i JOIN sys.tables t ON t.object_id=i.object_id JOIN sys.schemas s ON s.schema_id=t.schema_id WHERE s.name=@p AND t.is_ms_shipped=0 AND COALESCE(i.compression_delay,0)>0', N'@p sysname', @p=@schema;
IF EXISTS (SELECT 1 FROM sys.all_columns WHERE object_id=OBJECT_ID(N'sys.partitions') AND name=N'xml_compression')
  EXEC sys.sp_executesql N'INSERT INTO #rza_unsupported_indexes SELECT DISTINCT i.object_id,i.name,N''XML_COMPRESSION'' FROM sys.indexes i JOIN sys.partitions p ON p.object_id=i.object_id AND p.index_id=i.index_id JOIN sys.tables t ON t.object_id=i.object_id JOIN sys.schemas s ON s.schema_id=t.schema_id WHERE s.name=@p AND t.is_ms_shipped=0 AND p.xml_compression=1', N'@p sysname', @p=@schema;
IF OBJECT_ID(N'sys.selective_xml_indexes') IS NOT NULL
  EXEC sys.sp_executesql N'INSERT INTO #rza_unsupported_indexes SELECT i.object_id,i.name,N''SELECTIVE_XML_INDEX'' FROM sys.selective_xml_indexes i JOIN sys.tables t ON t.object_id=i.object_id JOIN sys.schemas s ON s.schema_id=t.schema_id WHERE s.name=@p AND t.is_ms_shipped=0', N'@p sysname', @p=@schema;

SELECT object_id, index_name, feature_name
FROM #rza_unsupported_indexes
ORDER BY object_id, feature_name, index_name;
`, sql.Named("schema", cfg.Schema))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := map[int][]string{}
	for rows.Next() {
		var objectID int
		var indexName, featureName string
		if err := rows.Scan(&objectID, &indexName, &featureName); err != nil {
			return nil, err
		}
		out[objectID] = append(out[objectID], featureName+":"+indexName)
	}
	return out, rows.Err()
}

func loadColumnsByTable(db *sql.DB, cfg AppConfig) (map[int][]columnMeta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
SELECT
  c.object_id,
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
  cc.definition AS computed_definition,
  COALESCE(cc.is_persisted, 0) AS is_persisted,
  c.is_rowguidcol,
  c.is_sparse,
  c.is_filestream
FROM sys.columns c
JOIN sys.tables t
  ON t.object_id = c.object_id
JOIN sys.schemas s
  ON s.schema_id = t.schema_id
JOIN sys.types ty
  ON ty.user_type_id = c.user_type_id
LEFT JOIN sys.identity_columns ic
  ON ic.object_id = c.object_id
 AND ic.column_id = c.column_id
LEFT JOIN sys.computed_columns cc
  ON cc.object_id = c.object_id
 AND cc.column_id = c.column_id
WHERE s.name = @schema
  AND t.is_ms_shipped = 0
ORDER BY c.object_id, c.column_id;
`, sql.Named("schema", cfg.Schema))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := map[int][]columnMeta{}
	for rows.Next() {
		var objectID int
		var c columnMeta
		var precision uint8
		var scale uint8
		if err := rows.Scan(
			&objectID,
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
			&c.ComputedDefinition,
			&c.IsPersisted,
			&c.IsRowGuidCol,
			&c.IsSparse,
			&c.IsFileStream,
		); err != nil {
			return nil, err
		}
		c.Precision = int(precision)
		c.Scale = int(scale)
		out[objectID] = append(out[objectID], c)
	}
	return out, rows.Err()
}

func loadDefaultConstraintsByTable(db *sql.DB, cfg AppConfig) (map[int][]defaultConstraintMeta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
SELECT
  dc.parent_object_id,
  dc.name,
  c.name AS column_name,
  dc.definition
FROM sys.default_constraints dc
JOIN sys.tables t
  ON t.object_id = dc.parent_object_id
JOIN sys.schemas s
  ON s.schema_id = t.schema_id
JOIN sys.columns c
  ON c.object_id = dc.parent_object_id
 AND c.column_id = dc.parent_column_id
WHERE s.name = @schema
  AND t.is_ms_shipped = 0
ORDER BY dc.parent_object_id, dc.name;
`, sql.Named("schema", cfg.Schema))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := map[int][]defaultConstraintMeta{}
	for rows.Next() {
		var objectID int
		var d defaultConstraintMeta
		if err := rows.Scan(&objectID, &d.Name, &d.ColumnName, &d.Definition); err != nil {
			return nil, err
		}
		out[objectID] = append(out[objectID], d)
	}
	return out, rows.Err()
}

func loadCheckConstraintsByTable(db *sql.DB, cfg AppConfig) (map[int][]checkConstraintMeta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
SELECT
  cc.parent_object_id,
  cc.name,
  cc.definition,
  cc.is_not_for_replication,
  cc.is_disabled,
  cc.is_not_trusted
FROM sys.check_constraints cc
JOIN sys.tables t
  ON t.object_id = cc.parent_object_id
JOIN sys.schemas s
  ON s.schema_id = t.schema_id
WHERE s.name = @schema
  AND t.is_ms_shipped = 0
ORDER BY cc.parent_object_id, cc.name;
`, sql.Named("schema", cfg.Schema))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := map[int][]checkConstraintMeta{}
	for rows.Next() {
		var objectID int
		var c checkConstraintMeta
		if err := rows.Scan(&objectID, &c.Name, &c.Definition, &c.IsNotForReplication, &c.IsDisabled, &c.IsNotTrusted); err != nil {
			return nil, err
		}
		out[objectID] = append(out[objectID], c)
	}
	return out, rows.Err()
}

func loadKeyConstraintsByTable(db *sql.DB, cfg AppConfig) (map[int][]keyConstraintMeta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
SELECT
  kc.parent_object_id,
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
JOIN sys.tables t
  ON t.object_id = kc.parent_object_id
JOIN sys.schemas s
  ON s.schema_id = t.schema_id
JOIN sys.indexes i
  ON i.object_id = kc.parent_object_id
 AND i.index_id = kc.unique_index_id
LEFT JOIN sys.data_spaces ds
  ON ds.data_space_id=i.data_space_id
WHERE s.name = @schema
  AND t.is_ms_shipped = 0
ORDER BY
  kc.parent_object_id,
  CASE kc.type WHEN 'PK' THEN 0 ELSE 1 END,
  kc.name;
`, sql.Named("schema", cfg.Schema))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := map[int][]keyConstraintMeta{}
	for rows.Next() {
		var objectID int
		var k keyConstraintMeta
		if err := rows.Scan(&objectID, &k.Name, &k.Type, &k.IndexID, &k.IndexType, &k.FillFactor, &k.IsPadded, &k.IgnoreDupKey, &k.AllowRowLocks, &k.AllowPageLocks, &k.IsDisabled, &k.DataSpaceType, &k.DataSpaceID, &k.CompressionDesc, &k.HasMixedCompression); err != nil {
			return nil, err
		}
		out[objectID] = append(out[objectID], k)
	}
	return out, rows.Err()
}

func loadKeyColumnsByTableAndIndex(db *sql.DB, cfg AppConfig) (map[objectIndexKey][]indexColumnMeta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
SELECT
  ic.object_id,
  ic.index_id,
  c.name,
  ic.is_descending_key
FROM sys.index_columns ic
JOIN sys.indexes i
  ON i.object_id = ic.object_id
 AND i.index_id = ic.index_id
JOIN sys.columns c
  ON c.object_id = ic.object_id
 AND c.column_id = ic.column_id
JOIN sys.tables t
  ON t.object_id = ic.object_id
JOIN sys.schemas s
  ON s.schema_id = t.schema_id
WHERE s.name = @schema
  AND t.is_ms_shipped = 0
  AND (ic.key_ordinal > 0 OR i.type IN (3, 4, 5, 6))
ORDER BY ic.object_id, ic.index_id, ic.key_ordinal;
`, sql.Named("schema", cfg.Schema))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := map[objectIndexKey][]indexColumnMeta{}
	for rows.Next() {
		var objectID int
		var indexID int
		var c indexColumnMeta
		if err := rows.Scan(&objectID, &indexID, &c.ColumnName, &c.Desc); err != nil {
			return nil, err
		}
		key := objectIndexKey{ObjectID: objectID, IndexID: indexID}
		out[key] = append(out[key], c)
	}
	return out, rows.Err()
}

func loadForeignKeysByTable(db *sql.DB, cfg AppConfig) (map[int][]foreignKeyMeta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
SELECT
  fk.parent_object_id,
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
JOIN sys.tables pt
  ON pt.object_id = fk.parent_object_id
JOIN sys.schemas ps
  ON ps.schema_id = pt.schema_id
JOIN sys.tables rt
  ON rt.object_id = fk.referenced_object_id
JOIN sys.schemas rs
  ON rs.schema_id = rt.schema_id
WHERE ps.name = @schema
  AND pt.is_ms_shipped = 0
ORDER BY fk.parent_object_id, fk.name;
`, sql.Named("schema", cfg.Schema))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := map[int][]foreignKeyMeta{}
	for rows.Next() {
		var parentObjectID int
		var fk foreignKeyMeta
		if err := rows.Scan(
			&parentObjectID,
			&fk.ObjectID,
			&fk.Name,
			&fk.RefSchema,
			&fk.RefTable,
			&fk.DeleteAction,
			&fk.UpdateAction,
			&fk.IsNotForReplication,
			&fk.IsDisabled,
			&fk.IsNotTrusted,
		); err != nil {
			return nil, err
		}
		out[parentObjectID] = append(out[parentObjectID], fk)
	}
	return out, rows.Err()
}

func loadForeignKeyColumnsByFK(db *sql.DB, cfg AppConfig) (map[int][]fkColumnMap, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
SELECT
  fkc.constraint_object_id,
  pc.name AS parent_col,
  rc.name AS ref_col
FROM sys.foreign_key_columns fkc
JOIN sys.foreign_keys fk
  ON fk.object_id = fkc.constraint_object_id
JOIN sys.tables pt
  ON pt.object_id = fk.parent_object_id
JOIN sys.schemas ps
  ON ps.schema_id = pt.schema_id
JOIN sys.columns pc
  ON pc.object_id = fkc.parent_object_id
 AND pc.column_id = fkc.parent_column_id
JOIN sys.columns rc
  ON rc.object_id = fkc.referenced_object_id
 AND rc.column_id = fkc.referenced_column_id
WHERE ps.name = @schema
  AND pt.is_ms_shipped = 0
ORDER BY fkc.constraint_object_id, fkc.constraint_column_id;
`, sql.Named("schema", cfg.Schema))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := map[int][]fkColumnMap{}
	for rows.Next() {
		var fkObjectID int
		var m fkColumnMap
		if err := rows.Scan(&fkObjectID, &m.ParentColumn, &m.RefColumn); err != nil {
			return nil, err
		}
		out[fkObjectID] = append(out[fkObjectID], m)
	}
	return out, rows.Err()
}

func loadIndexesByTable(db *sql.DB, cfg AppConfig) (map[int][]indexMeta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
SELECT
  i.object_id,
  i.name,
  i.index_id,
  i.type_desc,
  i.is_unique,
  i.filter_definition,
  i.has_filter,
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
  ) THEN 1 ELSE 0 END AS bit) AS has_mixed_compression,
  xi.secondary_type_desc,
  pxi.name AS primary_xml_index_name,
  sit.tessellation_scheme,
  sit.bounding_box_xmin,
  sit.bounding_box_ymin,
  sit.bounding_box_xmax,
  sit.bounding_box_ymax,
  sit.cells_per_object,
  sit.level_1_grid_desc,
  sit.level_2_grid_desc,
  sit.level_3_grid_desc,
  sit.level_4_grid_desc,
  i.fill_factor,
  i.is_padded,
  i.ignore_dup_key,
  i.allow_row_locks,
  i.allow_page_locks,
  i.is_disabled,
  ds.type_desc AS data_space_type,
  i.data_space_id
FROM sys.indexes i
JOIN sys.tables t
  ON t.object_id = i.object_id
JOIN sys.schemas s
  ON s.schema_id = t.schema_id
LEFT JOIN sys.xml_indexes xi
  ON xi.object_id=i.object_id AND xi.index_id=i.index_id
LEFT JOIN sys.indexes pxi
  ON pxi.object_id=xi.object_id AND pxi.index_id=xi.using_xml_index_id
LEFT JOIN sys.spatial_index_tessellations sit
  ON sit.object_id=i.object_id AND sit.index_id=i.index_id
LEFT JOIN sys.data_spaces ds
  ON ds.data_space_id=i.data_space_id
WHERE s.name = @schema
  AND t.is_ms_shipped = 0
  AND i.is_hypothetical = 0
  AND i.name IS NOT NULL
  AND i.is_primary_key = 0
  AND i.is_unique_constraint = 0
ORDER BY i.object_id, i.name;
`, sql.Named("schema", cfg.Schema))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := map[int][]indexMeta{}
	for rows.Next() {
		var objectID int
		var idx indexMeta
		if err := rows.Scan(&objectID, &idx.Name, &idx.IndexID, &idx.TypeDesc, &idx.IsUnique, &idx.FilterDefinition, &idx.HasFilter, &idx.CompressionDesc, &idx.HasMixedCompression, &idx.XMLSecondaryType, &idx.XMLPrimaryName, &idx.SpatialScheme, &idx.BoundingXMin, &idx.BoundingYMin, &idx.BoundingXMax, &idx.BoundingYMax, &idx.CellsPerObject, &idx.Level1Grid, &idx.Level2Grid, &idx.Level3Grid, &idx.Level4Grid, &idx.FillFactor, &idx.IsPadded, &idx.IgnoreDupKey, &idx.AllowRowLocks, &idx.AllowPageLocks, &idx.IsDisabled, &idx.DataSpaceType, &idx.DataSpaceID); err != nil {
			return nil, err
		}
		out[objectID] = append(out[objectID], idx)
	}
	return out, rows.Err()
}

func loadIndexIncludeColumnsByTableAndIndex(db *sql.DB, cfg AppConfig) (map[objectIndexKey][]indexColumnMeta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
SELECT
  ic.object_id,
  ic.index_id,
  c.name,
  ic.is_descending_key
FROM sys.index_columns ic
JOIN sys.columns c
  ON c.object_id = ic.object_id
 AND c.column_id = ic.column_id
JOIN sys.tables t
  ON t.object_id = ic.object_id
JOIN sys.schemas s
  ON s.schema_id = t.schema_id
WHERE s.name = @schema
  AND t.is_ms_shipped = 0
  AND ic.is_included_column = 1
ORDER BY ic.object_id, ic.index_id, ic.index_column_id;
`, sql.Named("schema", cfg.Schema))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := map[objectIndexKey][]indexColumnMeta{}
	for rows.Next() {
		var objectID int
		var indexID int
		var c indexColumnMeta
		if err := rows.Scan(&objectID, &indexID, &c.ColumnName, &c.Desc); err != nil {
			return nil, err
		}
		key := objectIndexKey{ObjectID: objectID, IndexID: indexID}
		out[key] = append(out[key], c)
	}
	return out, rows.Err()
}
