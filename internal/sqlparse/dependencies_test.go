package sqlparse

import (
	"reflect"
	"testing"
)

func TestExtractObjectReferencesDoesNotInventTargetFunctions(t *testing.T) {
	sql := `CREATE FUNCTION dbo.F(@id int) RETURNS int AS BEGIN RETURN @id; END;
INSERT dbo.T(Id) VALUES(1);
WITH rows(Id) AS (SELECT Id FROM dbo.Source) SELECT * FROM rows;
SELECT * FROM #work JOIN @items i ON 1=1;
EXEC OtherDB.dbo.Remote; EXEC OtherDB..Remote;
SELECT OtherDB.dbo.Remote_FN(1);`
	want := []ObjectReference{{"dbo.T", ObjectReferenceData}, {"dbo.Source", ObjectReferenceData}}
	if got := ExtractObjectReferences(sql); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v want %+v", got, want)
	}
}

func TestExtractObjectReferencesResolvesDMLAliasWithoutCatalogGuess(t *testing.T) {
	got := ExtractObjectReferences(`UPDATE t SET Id=1 FROM dbo.Real t WHERE EXISTS(SELECT 1 FROM dbo.Other t);`)
	want := []ObjectReference{{"dbo.Real", ObjectReferenceData}, {"dbo.Other", ObjectReferenceData}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v want %+v", got, want)
	}
}

func TestExtractObjectReferencesQuotedNames(t *testing.T) {
	got := ExtractObjectReferences(`SELECT "dbo"."My FN"(1) FROM [dbo].[Order Items] JOIN dbo.[a]]b] x ON 1=1;`)
	want := []ObjectReference{{"dbo.Order Items", ObjectReferenceData}, {"dbo.a]b", ObjectReferenceData}, {"dbo.My FN", ObjectReferenceFunction}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v want %+v", got, want)
	}
}

func TestExtractObjectReferencesIgnoresCommentsAndStrings(t *testing.T) {
	t.Parallel()

	sql := `
CREATE OR ALTER VIEW dbo.ViewA AS
SELECT dbo.fnReal(t.id)
FROM [dbo].[ViewB] t
JOIN dbo.TableC c ON c.id = t.id;
-- FROM dbo.NotAReference
SELECT 'EXEC dbo.AlsoNotAReference';
`

	got := ExtractObjectReferences(sql)
	want := map[string]ObjectReferenceKind{
		"dbo.fnReal": ObjectReferenceFunction,
		"dbo.ViewB":  ObjectReferenceData,
		"dbo.TableC": ObjectReferenceData,
	}
	if len(got) != len(want) {
		t.Fatalf("unexpected references: %#v", got)
	}
	for _, ref := range got {
		if want[ref.Name] != ref.Kind {
			t.Fatalf("unexpected reference: %#v", ref)
		}
	}
}

func TestExtractObjectReferencesFindsProcedureCalls(t *testing.T) {
	t.Parallel()

	got := ExtractObjectReferences(`EXEC @result = [dbo].[ProcB]; EXEC dbo.ProcC; EXEC ProcD; SELECT fnLocal(1);`)
	if len(got) != 4 || got[0].Kind != ObjectReferenceProcedure || got[1].Kind != ObjectReferenceProcedure || got[2].Kind != ObjectReferenceProcedure || got[3].Kind != ObjectReferenceFunction {
		t.Fatalf("unexpected procedure references: %#v", got)
	}
	if got[2].Name != "dbo.ProcD" || got[3].Name != "dbo.fnLocal" {
		t.Fatalf("unexpected unqualified reference normalization: %#v", got)
	}
}

func TestExtractObjectReferencesFindsUnqualifiedDataObjectsButSkipsCTE(t *testing.T) {
	t.Parallel()

	got := ExtractObjectReferences(`
WITH recent AS (SELECT id FROM dbo.TableA)
SELECT * FROM recent JOIN TreasuryCustomerSaldoTaxAmnesty_v t ON t.id = recent.id;
`)
	want := map[string]ObjectReferenceKind{
		"dbo.TableA":                            ObjectReferenceData,
		"dbo.TreasuryCustomerSaldoTaxAmnesty_v": ObjectReferenceData,
	}
	if len(got) != len(want) {
		t.Fatalf("unexpected data references: %#v", got)
	}
	for _, ref := range got {
		if want[ref.Name] != ref.Kind {
			t.Fatalf("unexpected data reference: %#v", ref)
		}
	}
}

func TestExtractObjectReferencesLimitsCTEToItsStatement(t *testing.T) {
	got := ExtractObjectReferences(`
WITH T AS (SELECT Id FROM dbo.Source)
SELECT * FROM T;
SELECT * FROM dbo.T;`)
	want := map[string]bool{"dbo.Source": false, "dbo.T": false}
	for _, ref := range got {
		if _, ok := want[ref.Name]; ok && ref.Kind == ObjectReferenceData {
			want[ref.Name] = true
		}
	}
	for name, found := range want {
		if !found {
			t.Fatalf("missing scoped dependency %s in %+v", name, got)
		}
	}
}

func TestExtractObjectReferencesPreservesCaseDistinctNames(t *testing.T) {
	t.Parallel()

	got := ExtractObjectReferences(`
WITH Recent AS (SELECT id FROM dbo.Source)
SELECT * FROM dbo.Code JOIN dbo.code ON 1 = 1 JOIN recent ON 1 = 1;
`)
	want := map[string]bool{"dbo.Source": true, "dbo.Code": true, "dbo.code": true, "dbo.recent": true}
	if len(got) != len(want) {
		t.Fatalf("case-distinct dependency set collapsed: %#v", got)
	}
	for _, ref := range got {
		if !want[ref.Name] {
			t.Fatalf("unexpected dependency: %#v", ref)
		}
	}
}

func TestExtractObjectReferencesClassifiesForeignKeyTargetsAsTables(t *testing.T) {
	t.Parallel()

	got := ExtractObjectReferences(`
CREATE TABLE dbo.Child (
  EventId INT,
  TenantId INT,
  CONSTRAINT [FK_Child_Parent] FOREIGN KEY(EventId)
    REFERENCES dbo.Parent([EventId]),
  CONSTRAINT [FK_Child_Tenant] FOREIGN KEY (TenantId)
    REFERENCES [dbo].[Tenant] ([TenantId])
);`)

	want := map[string]ObjectReferenceKind{
		"dbo.Parent": ObjectReferenceData,
		"dbo.Tenant": ObjectReferenceData,
	}
	if len(got) != len(want) {
		t.Fatalf("unexpected foreign-key references: %#v", got)
	}
	for _, ref := range got {
		if want[ref.Name] != ref.Kind {
			t.Fatalf("foreign-key target must be a data object: %#v", ref)
		}
	}
}

func TestExtractObjectReferencesStillClassifiesTableValuedFunctions(t *testing.T) {
	t.Parallel()

	got := ExtractObjectReferences(`SELECT * FROM dbo.fnRows(1) r JOIN dbo.TableA t ON t.Id = r.Id;`)
	if len(got) != 2 {
		t.Fatalf("unexpected references: %#v", got)
	}
	if got[0].Name != "dbo.fnRows" || got[0].Kind != ObjectReferenceFunction {
		t.Fatalf("FROM function must remain a function reference: %#v", got[0])
	}
	if got[1].Name != "dbo.TableA" || got[1].Kind != ObjectReferenceData {
		t.Fatalf("JOIN table must remain a data reference: %#v", got[1])
	}
}

func TestExtractObjectReferencesFindsTableValuedParameterTypes(t *testing.T) {
	t.Parallel()

	got := ExtractObjectReferences(`
CREATE OR ALTER PROCEDURE dbo.SaveOrder
  @items [dbo].[OrderItemType] READONLY,
  @code dbo.CodeType
AS
SELECT 1;
`)
	want := map[string]bool{
		"dbo.OrderItemType": true,
		"dbo.CodeType":      true,
	}
	found := map[string]bool{}
	for _, ref := range got {
		if ref.Kind == ObjectReferenceType {
			found[ref.Name] = true
		}
	}
	for name := range want {
		if !found[name] {
			t.Fatalf("missing type dependency %s in %#v", name, got)
		}
	}
}

func TestExtractObjectReferencesFindsUnqualifiedTableValuedParameterType(t *testing.T) {
	t.Parallel()

	got := ExtractObjectReferences(`
CREATE OR ALTER PROCEDURE dbo.SaveOrder
  @items [OrderItemType] READONLY
AS
SELECT COUNT(*) FROM @items;
`)
	for _, ref := range got {
		if ref.Kind == ObjectReferenceType && ref.Name == "dbo.OrderItemType" {
			return
		}
	}
	t.Fatalf("missing unqualified table type dependency in %#v", got)
}
