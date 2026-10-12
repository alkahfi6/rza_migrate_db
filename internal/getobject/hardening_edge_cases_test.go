package getobject

import (
	"database/sql"
	"strings"
	"testing"
)

func TestDefaultConstraintMapPreservesCaseDistinctColumns(t *testing.T) {
	got, err := defaultConstraintMapByColumn([]defaultConstraintMeta{
		{Name: "DF_Code", ColumnName: "Code", Definition: "('A')"},
		{Name: "DF_code", ColumnName: "code", Definition: "('b')"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got["Code"].Name != "DF_Code" || got["code"].Name != "DF_code" {
		t.Fatalf("case-sensitive default mapping was collapsed: %#v", got)
	}
}

func TestRenderComputedPersistedNotNull(t *testing.T) {
	got := renderColumnClause(columnMeta{
		Name: "Calc", IsComputed: true, IsPersisted: true, IsNullable: false,
		ComputedDefinition: sql.NullString{String: "([Base] + 1)", Valid: true},
	})
	if got != "[Calc] AS ([Base] + 1) PERSISTED NOT NULL" {
		t.Fatalf("computed nullability was lost: %s", got)
	}
}

func TestRenderSpatialManualGridIsValidSQLShape(t *testing.T) {
	idx := indexMeta{
		Name: "SIX_Geo", TypeDesc: "SPATIAL",
		AllowRowLocks: true, AllowPageLocks: true,
		SpatialScheme: sql.NullString{String: "GEOMETRY_GRID", Valid: true},
		BoundingXMin:  sql.NullFloat64{Float64: -180, Valid: true}, BoundingYMin: sql.NullFloat64{Float64: -90, Valid: true},
		BoundingXMax: sql.NullFloat64{Float64: 180, Valid: true}, BoundingYMax: sql.NullFloat64{Float64: 90, Valid: true},
		CellsPerObject: sql.NullInt64{Int64: 16, Valid: true},
		Level1Grid:     sql.NullString{String: "LOW", Valid: true}, Level2Grid: sql.NullString{String: "MEDIUM", Valid: true},
		Level3Grid: sql.NullString{String: "HIGH", Valid: true}, Level4Grid: sql.NullString{String: "HIGH", Valid: true},
	}
	got, err := renderCreateIndexClause("dbo", "Places", idx, []indexColumnMeta{{ColumnName: "Geo"}}, nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"USING GEOMETRY_GRID WITH (BOUNDING_BOX = (-180, -90, 180, 90)", "GRIDS = (LEVEL_1 = LOW, LEVEL_2 = MEDIUM, LEVEL_3 = HIGH, LEVEL_4 = HIGH)", "CELLS_PER_OBJECT = 16);"} {
		if !strings.Contains(got, want) {
			t.Fatalf("spatial SQL missing %q:\n%s", want, got)
		}
	}
}

func TestIndexRendererRejectsLossyCatalogAndKeepsFilteredColumnstore(t *testing.T) {
	if _, err := renderCreateIndexClause("dbo", "T", indexMeta{Name: "IX", TypeDesc: "NONCLUSTERED", HasMixedCompression: true}, []indexColumnMeta{{ColumnName: "Id"}}, nil); err == nil {
		t.Fatal("mixed compression must fail closed")
	}
	if _, err := renderCreateIndexClause("dbo", "T", indexMeta{Name: "IX", TypeDesc: "NONCLUSTERED", HasFilter: true}, []indexColumnMeta{{ColumnName: "Id"}}, nil); err == nil {
		t.Fatal("invisible filter definition must fail closed")
	}
	got, err := renderCreateIndexClause("dbo", "T", indexMeta{
		Name: "NCCI", TypeDesc: "NONCLUSTERED COLUMNSTORE", HasFilter: true,
		FilterDefinition: sql.NullString{String: "[IsActive] = (1)", Valid: true},
	}, []indexColumnMeta{{ColumnName: "Payload"}}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "WHERE [IsActive] = (1)") {
		t.Fatalf("filtered NCCI predicate was lost: %s", got)
	}
	if strings.Contains(got, "ALLOW_ROW_LOCKS") || strings.Contains(got, "ALLOW_PAGE_LOCKS") {
		t.Fatalf("rowstore-only lock options leaked into columnstore SQL: %s", got)
	}
	archive, err := renderCreateIndexClause("dbo", "T", indexMeta{
		Name: "CCI", TypeDesc: "CLUSTERED COLUMNSTORE",
		CompressionDesc: sql.NullString{String: "COLUMNSTORE_ARCHIVE", Valid: true},
	}, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(archive, "WITH (DATA_COMPRESSION = COLUMNSTORE_ARCHIVE)") {
		t.Fatalf("columnstore archive compression was lost: %s", archive)
	}
}
