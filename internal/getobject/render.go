package getobject

import (
	"database/sql"
	"fmt"
	"strings"
)

func renderColumnClause(c columnMeta) string {
	var b strings.Builder
	b.WriteString(quoteIdentifier(c.Name))
	b.WriteString(" ")
	if c.IsComputed {
		b.WriteString("AS ")
		b.WriteString(strings.TrimSpace(c.ComputedDefinition.String))
		if c.IsPersisted {
			b.WriteString(" PERSISTED")
			if !c.IsNullable {
				b.WriteString(" NOT NULL")
			}
		}
		return b.String()
	}
	b.WriteString(formatDataType(c))

	if c.Collation.Valid && strings.TrimSpace(c.Collation.String) != "" && dataTypeSupportsCollation(c.TypeName) {
		b.WriteString(" COLLATE ")
		b.WriteString(strings.TrimSpace(c.Collation.String))
	}

	if c.IsIdentity {
		seed := mustIntString(c.SeedValue, 1)
		inc := mustIntString(c.IncValue, 1)
		b.WriteString(" IDENTITY(")
		b.WriteString(fmt.Sprintf("%d,%d", seed, inc))
		b.WriteString(")")
	}
	if c.IsRowGuidCol {
		b.WriteString(" ROWGUIDCOL")
	}
	if c.IsSparse {
		b.WriteString(" SPARSE")
	}

	if c.IsNullable {
		b.WriteString(" NULL")
	} else {
		b.WriteString(" NOT NULL")
	}
	return b.String()
}

func renderInlineDefaultConstraintFragment(d defaultConstraintMeta) string {
	expr := strings.TrimSpace(d.Definition)
	if expr == "" {
		return ""
	}

	var b strings.Builder
	if strings.TrimSpace(d.Name) != "" {
		b.WriteString(" CONSTRAINT ")
		b.WriteString(quoteIdentifier(d.Name))
	}
	b.WriteString(" DEFAULT ")
	b.WriteString(expr)
	return b.String()
}

func renderDefaultConstraintClause(d defaultConstraintMeta) string {
	return fmt.Sprintf(
		"CONSTRAINT %s DEFAULT %s FOR %s",
		quoteIdentifier(d.Name),
		strings.TrimSpace(d.Definition),
		quoteIdentifier(d.ColumnName),
	)
}

func defaultConstraintMapByColumn(defaults []defaultConstraintMeta) (map[string]defaultConstraintMeta, error) {
	out := make(map[string]defaultConstraintMeta, len(defaults))
	for _, d := range defaults {
		name := strings.TrimSpace(d.Name)
		if name == "" {
			return nil, fmt.Errorf("default constraint name must not be empty for column %s", strings.TrimSpace(d.ColumnName))
		}
		col := strings.TrimSpace(d.ColumnName)
		if col == "" {
			return nil, fmt.Errorf("default constraint %s has empty column name", strings.TrimSpace(d.Name))
		}
		if strings.TrimSpace(d.Definition) == "" {
			return nil, fmt.Errorf("default constraint %s on column %s has empty definition", strings.TrimSpace(d.Name), strings.TrimSpace(d.ColumnName))
		}
		if _, exists := out[col]; exists {
			return nil, fmt.Errorf("multiple default constraints detected for column %s", strings.TrimSpace(d.ColumnName))
		}
		out[col] = d
	}
	return out, nil
}

func renderCheckConstraintClause(c checkConstraintMeta) string {
	notForReplication := ""
	if c.IsNotForReplication {
		notForReplication = " NOT FOR REPLICATION"
	}
	return fmt.Sprintf(
		"CONSTRAINT %s CHECK%s %s",
		quoteIdentifier(c.Name),
		notForReplication,
		strings.TrimSpace(c.Definition),
	)
}

func validateKeyConstraintMetadata(k keyConstraintMeta) error {
	indexType := strings.ToUpper(strings.TrimSpace(k.IndexType))
	if indexType == "" {
		return fmt.Errorf("key constraint %s has empty backing index type", k.Name)
	}
	if indexType != "CLUSTERED" && indexType != "NONCLUSTERED" {
		return fmt.Errorf("key constraint %s uses unsupported backing index type %s", k.Name, k.IndexType)
	}
	return nil
}

func renderKeyConstraintClause(k keyConstraintMeta, cols []indexColumnMeta) string {
	keyType := "UNIQUE"
	if strings.EqualFold(k.Type, "PK") {
		keyType = "PRIMARY KEY"
	}

	indexType := strings.ToUpper(strings.TrimSpace(k.IndexType))
	if indexType != "CLUSTERED" && indexType != "NONCLUSTERED" {
		if strings.Contains(indexType, "CLUSTERED") {
			indexType = "CLUSTERED"
		} else {
			indexType = "NONCLUSTERED"
		}
	}

	return fmt.Sprintf(
		"CONSTRAINT %s %s %s (%s)",
		quoteIdentifier(k.Name),
		keyType,
		indexType,
		renderIndexColumnList(cols, true),
	)
}

func renderForeignKeyClause(schema string, fk foreignKeyMeta, maps []fkColumnMap) string {
	parentCols := make([]string, 0, len(maps))
	refCols := make([]string, 0, len(maps))
	for _, m := range maps {
		parentCols = append(parentCols, quoteIdentifier(m.ParentColumn))
		refCols = append(refCols, quoteIdentifier(m.RefColumn))
	}

	var b strings.Builder
	b.WriteString("CONSTRAINT ")
	b.WriteString(quoteIdentifier(fk.Name))
	b.WriteString(" FOREIGN KEY (")
	b.WriteString(strings.Join(parentCols, ", "))
	b.WriteString(") REFERENCES ")
	b.WriteString(qualifiedName(schema, fk.RefTable))
	b.WriteString(" (")
	b.WriteString(strings.Join(refCols, ", "))
	b.WriteString(")")

	deleteAction := normalizeFKAction(fk.DeleteAction)
	if deleteAction != "" {
		b.WriteString(" ON DELETE ")
		b.WriteString(deleteAction)
	}
	updateAction := normalizeFKAction(fk.UpdateAction)
	if updateAction != "" {
		b.WriteString(" ON UPDATE ")
		b.WriteString(updateAction)
	}
	if fk.IsNotForReplication {
		b.WriteString(" NOT FOR REPLICATION")
	}
	return b.String()
}

func renderCreateIndexClause(schema, table string, idx indexMeta, keyCols, includeCols []indexColumnMeta) (string, error) {
	var b strings.Builder
	typeDesc := strings.ToUpper(strings.TrimSpace(idx.TypeDesc))
	nameAndTarget := quoteIdentifier(idx.Name) + " ON " + qualifiedName(schema, table)
	if idx.IsDisabled {
		return "", fmt.Errorf("index %s is disabled and cannot be represented as an enabled CREATE INDEX snapshot", idx.Name)
	}
	if idx.DataSpaceType.Valid && strings.EqualFold(strings.TrimSpace(idx.DataSpaceType.String), "PARTITION_SCHEME") {
		return "", fmt.Errorf("index %s uses a partition scheme and cannot be represented without partition metadata", idx.Name)
	}
	if idx.DataSpaceID > 1 {
		return "", fmt.Errorf("index %s uses non-default data space %d and cannot be represented without filegroup metadata", idx.Name, idx.DataSpaceID)
	}
	if idx.HasMixedCompression {
		return "", fmt.Errorf("index %s uses mixed per-partition compression and cannot be represented losslessly", idx.Name)
	}
	if idx.HasFilter && (!idx.FilterDefinition.Valid || strings.TrimSpace(idx.FilterDefinition.String) == "") {
		return "", fmt.Errorf("index %s is filtered but its filter definition is not visible", idx.Name)
	}

	switch typeDesc {
	case "XML":
		compression := strings.ToUpper(strings.TrimSpace(idx.CompressionDesc.String))
		if (compression != "" && compression != "NONE") || idx.IgnoreDupKey {
			return "", fmt.Errorf("XML index %s has unsupported compression or IGNORE_DUP_KEY metadata", idx.Name)
		}
		if len(keyCols) != 1 {
			return "", fmt.Errorf("XML index %s must expose exactly one XML column", idx.Name)
		}
		if idx.XMLPrimaryName.Valid && strings.TrimSpace(idx.XMLPrimaryName.String) != "" {
			secondaryType := strings.ToUpper(strings.TrimSpace(idx.XMLSecondaryType.String))
			if secondaryType != "PATH" && secondaryType != "VALUE" && secondaryType != "PROPERTY" {
				return "", fmt.Errorf("XML index %s has unsupported secondary type %s", idx.Name, secondaryType)
			}
			b.WriteString(fmt.Sprintf("CREATE XML INDEX %s (%s) USING XML INDEX %s FOR %s", nameAndTarget, quoteIdentifier(keyCols[0].ColumnName), quoteIdentifier(idx.XMLPrimaryName.String), secondaryType))
			appendIndexOptions(&b, idx)
			b.WriteString(";")
			return b.String(), nil
		}
		b.WriteString(fmt.Sprintf("CREATE PRIMARY XML INDEX %s (%s)", nameAndTarget, quoteIdentifier(keyCols[0].ColumnName)))
		appendIndexOptions(&b, idx)
		b.WriteString(";")
		return b.String(), nil
	case "SPATIAL":
		if idx.IgnoreDupKey {
			return "", fmt.Errorf("spatial index %s unexpectedly has IGNORE_DUP_KEY enabled", idx.Name)
		}
		if len(keyCols) != 1 || !idx.SpatialScheme.Valid {
			return "", fmt.Errorf("spatial index %s metadata is incomplete", idx.Name)
		}
		options := make([]string, 0, 3)
		scheme := strings.ToUpper(strings.TrimSpace(idx.SpatialScheme.String))
		if strings.HasPrefix(scheme, "GEOMETRY_") && !(idx.BoundingXMin.Valid && idx.BoundingYMin.Valid && idx.BoundingXMax.Valid && idx.BoundingYMax.Valid) {
			return "", fmt.Errorf("spatial index %s is geometry but bounding-box metadata is incomplete", idx.Name)
		}
		if idx.BoundingXMin.Valid && idx.BoundingYMin.Valid && idx.BoundingXMax.Valid && idx.BoundingYMax.Valid {
			options = append(options, fmt.Sprintf("BOUNDING_BOX = (%g, %g, %g, %g)", idx.BoundingXMin.Float64, idx.BoundingYMin.Float64, idx.BoundingXMax.Float64, idx.BoundingYMax.Float64))
		}
		if strings.HasSuffix(scheme, "_GRID") && !strings.Contains(scheme, "AUTO") {
			levels := []sql.NullString{idx.Level1Grid, idx.Level2Grid, idx.Level3Grid, idx.Level4Grid}
			grid := make([]string, 0, 4)
			for i, level := range levels {
				value := strings.ToUpper(strings.TrimSpace(level.String))
				if !level.Valid || (value != "LOW" && value != "MEDIUM" && value != "HIGH") {
					return "", fmt.Errorf("spatial index %s has incomplete manual grid metadata", idx.Name)
				}
				grid = append(grid, fmt.Sprintf("LEVEL_%d = %s", i+1, value))
			}
			options = append(options, "GRIDS = ("+strings.Join(grid, ", ")+")")
		}
		if idx.CellsPerObject.Valid {
			options = append(options, fmt.Sprintf("CELLS_PER_OBJECT = %d", idx.CellsPerObject.Int64))
		}
		options = append(options, persistedRowstoreOptionParts(idx, false)...)
		optionSQL := ""
		if len(options) > 0 {
			optionSQL = " WITH (" + strings.Join(options, ", ") + ")"
		}
		return fmt.Sprintf("CREATE SPATIAL INDEX %s (%s) USING %s%s;", nameAndTarget, quoteIdentifier(keyCols[0].ColumnName), strings.TrimSpace(idx.SpatialScheme.String), optionSQL), nil
	case "CLUSTERED COLUMNSTORE":
		if hasNonDefaultIndexOptions(idx) {
			return "", fmt.Errorf("columnstore index %s has unsupported persisted rowstore options", idx.Name)
		}
		if idx.HasFilter {
			return "", fmt.Errorf("clustered columnstore index %s unexpectedly has a filter", idx.Name)
		}
		b.WriteString("CREATE CLUSTERED COLUMNSTORE INDEX ")
		b.WriteString(nameAndTarget)
		if err := appendColumnstoreOptions(&b, idx); err != nil {
			return "", err
		}
		b.WriteString(";")
		return b.String(), nil
	case "NONCLUSTERED COLUMNSTORE":
		if hasNonDefaultIndexOptions(idx) {
			return "", fmt.Errorf("columnstore index %s has unsupported persisted rowstore options", idx.Name)
		}
		if len(keyCols) == 0 {
			return "", fmt.Errorf("columnstore index %s has zero columns", idx.Name)
		}
		b.WriteString("CREATE NONCLUSTERED COLUMNSTORE INDEX ")
		b.WriteString(nameAndTarget)
		b.WriteString(" (")
		b.WriteString(renderIndexColumnList(keyCols, false))
		b.WriteString(")")
		if idx.HasFilter {
			b.WriteString(" WHERE ")
			b.WriteString(strings.TrimSpace(idx.FilterDefinition.String))
		}
		if err := appendColumnstoreOptions(&b, idx); err != nil {
			return "", err
		}
		b.WriteString(";")
		return b.String(), nil
	default:

		if typeDesc != "CLUSTERED" && typeDesc != "NONCLUSTERED" {
			return "", fmt.Errorf("index %s uses unsupported type %s", idx.Name, idx.TypeDesc)
		}

		b.WriteString("CREATE ")
		if idx.IsUnique {
			b.WriteString("UNIQUE ")
		}
		b.WriteString(typeDesc)
		b.WriteString(" INDEX ")
		b.WriteString(nameAndTarget)
		b.WriteString(" (")
		b.WriteString(renderIndexColumnList(keyCols, true))
		b.WriteString(")")

		if len(includeCols) > 0 {
			names := make([]string, 0, len(includeCols))
			for _, col := range includeCols {
				names = append(names, quoteIdentifier(col.ColumnName))
			}
			b.WriteString(" INCLUDE (")
			b.WriteString(strings.Join(names, ", "))
			b.WriteString(")")
		}
		if idx.FilterDefinition.Valid && strings.TrimSpace(idx.FilterDefinition.String) != "" {
			b.WriteString(" WHERE ")
			b.WriteString(strings.TrimSpace(idx.FilterDefinition.String))
		}
		appendIndexOptions(&b, idx)
		b.WriteString(";")
		return b.String(), nil
	}
}

func appendColumnstoreOptions(b *strings.Builder, idx indexMeta) error {
	compression := strings.ToUpper(strings.TrimSpace(idx.CompressionDesc.String))
	if compression == "" || compression == "NONE" || compression == "COLUMNSTORE" {
		return nil
	}
	// COLUMNSTORE_ARCHIVE is the only non-default compression state currently
	// reconstructed. Other catalog values fail closed instead of leaking a
	// rowstore-only option into CREATE COLUMNSTORE INDEX.
	if compression == "COLUMNSTORE_ARCHIVE" {
		b.WriteString(" WITH (DATA_COMPRESSION = COLUMNSTORE_ARCHIVE)")
		return nil
	}
	return fmt.Errorf("columnstore index %s has unsupported compression %s", idx.Name, compression)
}

func hasNonDefaultIndexOptions(idx indexMeta) bool {
	// SQL Server reports lock flags as OFF/NULL-like values for some index
	// families where those rowstore options do not apply (notably columnstore).
	return idx.FillFactor != 0 || idx.IsPadded || idx.IgnoreDupKey
}

func appendIndexOptions(b *strings.Builder, idx indexMeta) {
	options := persistedRowstoreOptionParts(idx, true)
	if len(options) == 0 {
		return
	}
	b.WriteString(" WITH (")
	b.WriteString(strings.Join(options, ", "))
	b.WriteString(")")
}

func persistedRowstoreOptionParts(idx indexMeta, includeIgnoreDupKey bool) []string {
	options := make([]string, 0, 6)
	compression := strings.ToUpper(strings.TrimSpace(idx.CompressionDesc.String))
	if compression != "" && compression != "NONE" {
		options = append(options, "DATA_COMPRESSION = "+compression)
	}
	if idx.FillFactor != 0 {
		options = append(options, fmt.Sprintf("FILLFACTOR = %d", idx.FillFactor))
	}
	if idx.IsPadded {
		options = append(options, "PAD_INDEX = ON")
	}
	if includeIgnoreDupKey && idx.IgnoreDupKey {
		options = append(options, "IGNORE_DUP_KEY = ON")
	}
	if !idx.AllowRowLocks {
		options = append(options, "ALLOW_ROW_LOCKS = OFF")
	}
	if !idx.AllowPageLocks {
		options = append(options, "ALLOW_PAGE_LOCKS = OFF")
	}
	return options
}

func renderIndexColumnList(cols []indexColumnMeta, withDirection bool) string {
	parts := make([]string, 0, len(cols))
	for _, c := range cols {
		p := quoteIdentifier(c.ColumnName)
		if withDirection {
			if c.Desc {
				p += " DESC"
			} else {
				p += " ASC"
			}
		}
		parts = append(parts, p)
	}
	return strings.Join(parts, ", ")
}

func normalizeFKAction(action string) string {
	upper := strings.ToUpper(strings.TrimSpace(action))
	switch upper {
	case "", "NO_ACTION":
		return ""
	case "CASCADE":
		return "CASCADE"
	case "SET_NULL":
		return "SET NULL"
	case "SET_DEFAULT":
		return "SET DEFAULT"
	default:
		return strings.ReplaceAll(upper, "_", " ")
	}
}

func quoteIdentifier(name string) string {
	return "[" + strings.ReplaceAll(name, "]", "]]") + "]"
}

func qualifiedName(schema, object string) string {
	return strings.TrimSpace(schema) + "." + strings.TrimSpace(object)
}

func formatDataType(c columnMeta) string {
	base := strings.ToUpper(strings.TrimSpace(c.TypeName))
	switch base {
	case "VARCHAR", "CHAR", "VARBINARY", "BINARY":
		if c.MaxLength < 0 {
			return base + "(MAX)"
		}
		return fmt.Sprintf("%s(%d)", base, c.MaxLength)
	case "NVARCHAR", "NCHAR":
		if c.MaxLength < 0 {
			return base + "(MAX)"
		}
		return fmt.Sprintf("%s(%d)", base, c.MaxLength/2)
	case "DECIMAL", "NUMERIC":
		return fmt.Sprintf("%s(%d,%d)", base, c.Precision, c.Scale)
	case "TIME", "DATETIME2", "DATETIMEOFFSET":
		return fmt.Sprintf("%s(%d)", base, c.Scale)
	case "FLOAT":
		if c.Precision > 0 && c.Precision != 53 {
			return fmt.Sprintf("%s(%d)", base, c.Precision)
		}
		return base
	default:
		return base
	}
}

func dataTypeSupportsCollation(typeName string) bool {
	switch strings.ToUpper(strings.TrimSpace(typeName)) {
	case "CHAR", "VARCHAR", "TEXT", "NCHAR", "NVARCHAR", "NTEXT":
		return true
	default:
		return false
	}
}
