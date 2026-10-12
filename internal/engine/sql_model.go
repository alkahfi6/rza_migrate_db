package engine

import (
	"fmt"
	"sort"
	"strings"
)

// SQLType is the normalized type portion of a column declaration. It is an
// in-memory compiler model derived exclusively from db_object/migrations; it
// is deliberately not persisted as another source of truth.
type SQLType struct {
	Base        string
	Declaration string
	Length      int
	Unicode     bool
}

// ColumnModel is the subset of a SQL Server column for which the engine can
// reason about a schema transition without guessing.
type ColumnModel struct {
	Name       string
	Type       SQLType
	Nullable   bool
	Computed   bool
	Expression string
	Identity   string
	RowGuidCol bool
	Sparse     bool
}

// TableModel is a typed, ephemeral projection of CREATE TABLE source.
type TableModel struct {
	Name    string
	Columns map[string]ColumnModel // lower-case column name -> definition
}

// IndexModel is the typed state used by index lifecycle planning.
type IndexModel struct {
	Table   string
	Indexes map[string]string // lower-case index name -> normalized CREATE INDEX
}

// BuildTableModel parses a supported CREATE TABLE into an in-memory model.
// Callers must retain the original SQL as the authoritative representation.
func BuildTableModel(object, createTableSQL string) (TableModel, error) {
	columns, err := parseTableColumns(createTableSQL)
	if err != nil {
		return TableModel{}, err
	}
	model := TableModel{Name: object, Columns: make(map[string]ColumnModel, len(columns))}
	for key, col := range columns {
		if strings.TrimSpace(col.Name) == "" {
			return TableModel{}, fmt.Errorf("table %s has unnamed column", object)
		}
		model.Columns[key] = ColumnModel{
			Name:       col.Name,
			Type:       SQLType{Base: normalizeColumnBaseType(col.BaseType), Declaration: normalizeColumnTypeDeclaration(col.RawType), Length: col.Length, Unicode: col.IsUnicode},
			Nullable:   col.Nullable,
			Computed:   col.IsComputed,
			Expression: col.ComputedDefinition,
			Identity:   col.IdentityDefinition,
			RowGuidCol: col.RowGuidCol,
			Sparse:     col.Sparse,
		}
	}
	return model, nil
}

// BuildIndexModel projects supported CREATE INDEX statements into typed state.
func BuildIndexModel(object, sql string) (IndexModel, error) {
	state, err := parseIndexStateMap(object, sql)
	if err != nil {
		return IndexModel{}, err
	}
	model := IndexModel{Table: object, Indexes: make(map[string]string, len(state))}
	for name, def := range state {
		model.Indexes[name] = def.Statement
	}
	return model, nil
}

// TypedTableChange is the loss-aware view consumed by the risk engine.
type TypedTableChange struct {
	Added   []ColumnModel
	Removed []ColumnModel
	Changed []ColumnTransition
}

type ColumnTransition struct {
	Before ColumnModel
	After  ColumnModel
}

func DiffTableModels(before, after TableModel) TypedTableChange {
	result := TypedTableChange{}
	for key, old := range before.Columns {
		current, ok := after.Columns[key]
		if !ok {
			result.Removed = append(result.Removed, old)
			continue
		}
		if old != current {
			result.Changed = append(result.Changed, ColumnTransition{Before: old, After: current})
		}
	}
	for key, current := range after.Columns {
		if _, ok := before.Columns[key]; !ok {
			result.Added = append(result.Added, current)
		}
	}
	sort.Slice(result.Added, func(i, j int) bool { return result.Added[i].Name < result.Added[j].Name })
	sort.Slice(result.Removed, func(i, j int) bool { return result.Removed[i].Name < result.Removed[j].Name })
	sort.Slice(result.Changed, func(i, j int) bool { return result.Changed[i].After.Name < result.Changed[j].After.Name })
	return result
}
