package getobject

import (
	"database/sql"
	"strings"
	"testing"
)

func TestEnsureCreateOrAlterProcedureProcAlias(t *testing.T) {
	t.Parallel()

	in := "CREATE PROC dbo.ProcA AS\nSELECT 1;\n"
	out, err := ensureCreateOrAlter(in, "PROCEDURE")
	if err != nil {
		t.Fatalf("ensureCreateOrAlter failed: %v", err)
	}
	if out[:25] != "CREATE OR ALTER PROCEDURE" {
		t.Fatalf("unexpected prefix: %q", out)
	}
}

func TestFormatDataType(t *testing.T) {
	t.Parallel()

	col := columnMeta{TypeName: "nvarchar", MaxLength: 100}
	if got := formatDataType(col); got != "NVARCHAR(50)" {
		t.Fatalf("unexpected nvarchar render: %s", got)
	}
	col = columnMeta{TypeName: "decimal", Precision: 18, Scale: 4}
	if got := formatDataType(col); got != "DECIMAL(18,4)" {
		t.Fatalf("unexpected decimal render: %s", got)
	}
}

func TestRenderColumnClauseIdentityAndNullability(t *testing.T) {
	t.Parallel()

	col := columnMeta{
		Name:       "id",
		TypeName:   "int",
		IsNullable: false,
		IsIdentity: true,
		SeedValue:  sql.NullString{String: "1", Valid: true},
		IncValue:   sql.NullString{String: "1", Valid: true},
	}
	got := renderColumnClause(col)
	want := "[id] INT IDENTITY(1,1) NOT NULL"
	if got != want {
		t.Fatalf("unexpected column clause:\nwant: %s\ngot:  %s", want, got)
	}
}

func TestRenderCreateIndexClauseVariants(t *testing.T) {
	t.Parallel()

	idx := indexMeta{
		Name:             "IX_TableA__status_active",
		TypeDesc:         "NONCLUSTERED",
		IsUnique:         true,
		FilterDefinition: sql.NullString{String: "[status] = 'ACTIVE'", Valid: true},
	}
	keyCols := []indexColumnMeta{
		{ColumnName: "status", Desc: false},
	}
	includeCols := []indexColumnMeta{
		{ColumnName: "amount"},
	}

	stmt, err := renderCreateIndexClause("dbo", "TableA", idx, keyCols, includeCols)
	if err != nil {
		t.Fatalf("renderCreateIndexClause failed: %v", err)
	}

	expect := []string{
		"CREATE UNIQUE NONCLUSTERED INDEX [IX_TableA__status_active] ON dbo.TableA ([status] ASC)",
		"INCLUDE ([amount])",
		"WHERE [status] = 'ACTIVE'",
	}
	for _, needle := range expect {
		if !strings.Contains(stmt, needle) {
			t.Fatalf("index statement missing %q:\n%s", needle, stmt)
		}
	}
	if !strings.HasSuffix(stmt, ";") {
		t.Fatalf("index statement must end with semicolon: %s", stmt)
	}
}

func TestRenderCreateIndexClausePreservesRowstoreOptionsAndRejectsDisabled(t *testing.T) {
	idx := indexMeta{Name: "IX_T", TypeDesc: "NONCLUSTERED", FillFactor: 80, IsPadded: true, AllowRowLocks: false, AllowPageLocks: false}
	sqlText, err := renderCreateIndexClause("dbo", "T", idx, []indexColumnMeta{{ColumnName: "id"}}, nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, token := range []string{"FILLFACTOR = 80", "PAD_INDEX = ON", "ALLOW_ROW_LOCKS = OFF", "ALLOW_PAGE_LOCKS = OFF"} {
		if !strings.Contains(sqlText, token) {
			t.Fatalf("missing option %s in %s", token, sqlText)
		}
	}
	idx.IsDisabled = true
	if _, err := renderCreateIndexClause("dbo", "T", idx, []indexColumnMeta{{ColumnName: "id"}}, nil); err == nil {
		t.Fatal("disabled index must be fail-closed")
	}
}

func TestConstraintMetadataRendersLogicalDefinitionsAndRejectsUnknownKeyType(t *testing.T) {
	t.Parallel()

	checkSQL := renderCheckConstraintClause(checkConstraintMeta{
		Name:                "CK_ExactCase",
		Definition:          "([ExactCase] > 0)",
		IsNotForReplication: true,
	})
	if checkSQL != "CONSTRAINT [CK_ExactCase] CHECK NOT FOR REPLICATION ([ExactCase] > 0)" {
		t.Fatalf("unexpected CHECK rendering: %s", checkSQL)
	}

	fkSQL := renderForeignKeyClause("dbo", foreignKeyMeta{
		Name:                "FK_Child_Parent",
		RefTable:            "Parent",
		DeleteAction:        "CASCADE",
		IsNotForReplication: true,
	}, []fkColumnMap{{ParentColumn: "ParentId", RefColumn: "Id"}})
	for _, token := range []string{"ON DELETE CASCADE", "NOT FOR REPLICATION"} {
		if !strings.Contains(fkSQL, token) {
			t.Fatalf("foreign key lost %q: %s", token, fkSQL)
		}
	}

	logicalKeys := []keyConstraintMeta{
		{Name: "PK_Fill", IndexType: "CLUSTERED", FillFactor: 80},
		{Name: "PK_Compressed", IndexType: "CLUSTERED", CompressionDesc: sql.NullString{String: "PAGE", Valid: true}},
		{Name: "PK_Partitioned", IndexType: "CLUSTERED", DataSpaceType: sql.NullString{String: "PARTITION_SCHEME", Valid: true}},
		{Name: "PK_Disabled", IndexType: "CLUSTERED", IsDisabled: true},
	}
	for _, key := range logicalKeys {
		if err := validateKeyConstraintMetadata(key); err != nil {
			t.Fatalf("logical key definition rejected because of physical metadata: %+v: %v", key, err)
		}
	}
	if err := validateKeyConstraintMetadata(keyConstraintMeta{Name: "PK_Hash", IndexType: "NONCLUSTERED HASH"}); err == nil {
		t.Fatal("unknown key backing index type must remain unsupported")
	}
	if err := validateKeyConstraintMetadata(keyConstraintMeta{Name: "PK_MissingType"}); err == nil {
		t.Fatal("empty key backing index type must fail closed")
	}
}

func TestValidateObjectNameForRules(t *testing.T) {
	t.Parallel()

	if err := validateObjectNameForRules("TableA_1"); err != nil {
		t.Fatalf("expected valid object name: %v", err)
	}
	if err := validateObjectNameForRules("tmp_#data_param"); err != nil {
		t.Fatalf("expected valid object name with hash in middle: %v", err)
	}
	if err := validateObjectNameForRules("Table$A"); err == nil {
		t.Fatalf("expected invalid object name with unsupported symbol")
	}
	if err := validateObjectNameForRules("Table A"); err == nil {
		t.Fatalf("expected invalid object name with space")
	}
}
