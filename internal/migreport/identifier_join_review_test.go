package migreport

import (
	"strings"
	"testing"
)

func TestSuspectedJoinIdentifiersIncludeUnindexedColumnsOnKeyedTable(t *testing.T) {
	docs := DesiredStateDocs{
		Tables: []DocumentedTable{{Name: "dbo.Child", Columns: []DocumentedColumn{{Name: "Id"}, {Name: "ParentId"}, {Name: "Code"}, {Name: "UnusedId"}}, Keys: []DocumentedKey{{Name: "PK_Child", Evidence: "PRIMARY_KEY", Columns: []string{"Id"}}}}},
		Relationships: []ObservedRelationship{
			{LeftObject: "dbo.Child", LeftColumns: []string{"ParentId"}, RightObject: "dbo.Parent", RightColumns: []string{"Id"}, Evidence: "OBSERVED_JOIN", Occurrences: 2, Modules: []string{"dbo.P2", "dbo.P1"}},
			{LeftObject: "dbo.Child", LeftColumns: []string{"Code"}, RightObject: "dbo.Lookup", RightColumns: []string{"Code"}, Evidence: "OBSERVED_JOIN", Occurrences: 1, Modules: []string{"dbo.P1"}},
			{LeftObject: "dbo.Child", LeftColumns: []string{"UnusedId"}, RightObject: "dbo.Parent", RightColumns: []string{"Id"}, Evidence: "DECLARED_FK", Occurrences: 1, Modules: []string{"FK_Child_Parent"}},
		},
	}
	report := renderSuspectedJoinIdentifiers(docs)
	for _, want := range []string{"| `dbo.Child` | `ParentId` | `OBSERVED_JOIN_ID_SIGNAL` | `NOT_ENFORCED` | 2 |", "| `dbo.Child` | `Code` | `OBSERVED_JOIN_COLUMN_SIGNAL` | `NOT_ENFORCED` | 1 |", "dbo.Child.ParentId = dbo.Parent.Id", "Source modules: dbo.P1; dbo.P2", "Index/key: tidak ditemukan. IDENTITY: tidak."} {
		if !strings.Contains(report, want) {
			t.Fatalf("missing %q: %s", want, report)
		}
	}
	if strings.Contains(report, "| `dbo.Child` | `Id` |") || strings.Contains(report, "UnusedId") || strings.Contains(report, "FK_Child_Parent") {
		t.Fatalf("invented JOIN evidence/duplicate enforced key: %s", report)
	}
	if got := joinUsage(docs, "dbo.Child", []string{"UnusedId"}); got != 0 {
		t.Fatalf("declared FK inflated JOIN count: %d", got)
	}
}

func TestSuspectedJoinIdentifiersExplainCompositeFilteredAndSelfJoin(t *testing.T) {
	docs := DesiredStateDocs{Tables: []DocumentedTable{{Name: "dbo.T", Columns: []DocumentedColumn{{Name: "Id", Identity: true}, {Name: "Scope"}, {Name: "Code"}}, Keys: []DocumentedKey{{Name: "PK_T", Evidence: "PRIMARY_KEY", Columns: []string{"Id", "Scope"}}, {Name: "UX_Code", Evidence: "FILTERED_UNIQUE_INDEX", Filtered: true, Columns: []string{"Code"}}}, Indexes: []DocumentedIndex{{Name: "UX_Code", Unique: true, Filtered: true, Columns: []string{"Code"}}}}}, Relationships: []ObservedRelationship{
		{LeftObject: "dbo.T", LeftColumns: []string{"Id"}, RightObject: "dbo.T", RightColumns: []string{"Id"}, Evidence: "OBSERVED_JOIN", Occurrences: 3, Modules: []string{"dbo.P"}},
		{LeftObject: "dbo.T", LeftColumns: []string{"Code"}, RightObject: "dbo.Other", RightColumns: []string{"Code"}, Evidence: "OBSERVED_JOIN", Occurrences: 1, Modules: []string{"dbo.P"}},
	}}
	report := renderSuspectedJoinIdentifiers(docs)
	for _, want := range []string{"| `dbo.T` | `Id` | `OBSERVED_JOIN_ID_SIGNAL` | `NOT_ENFORCED` | 3 |", "IDENTITY: ya", "composite key (Id, Scope)", "| `dbo.T` | `Code` | `OBSERVED_JOIN_COLUMN_SIGNAL` | `CONDITIONAL` | 1 |", "Index/key: UX_Code"} {
		if !strings.Contains(report, want) {
			t.Fatalf("missing %q: %s", want, report)
		}
	}
	if strings.Contains(report, "UX_Code; UX_Code") || joinUsage(docs, "dbo.T", []string{"Id"}) != 3 {
		t.Fatalf("duplicate source/role count: %s", report)
	}
}

func TestObservedJoinWithoutIDNameDoesNotClaimNoIdentifierEvidence(t *testing.T) {
	docs := DesiredStateDocs{Tables: []DocumentedTable{{Name: "dbo.T", Columns: []DocumentedColumn{{Name: "Code"}}}}, Relationships: []ObservedRelationship{{LeftObject: "dbo.T", LeftColumns: []string{"Code"}, RightObject: "dbo.Lookup", RightColumns: []string{"Code"}, Evidence: "OBSERVED_JOIN", Occurrences: 1, Modules: []string{"dbo.P"}}}}
	report := renderUniqueIDs(docs)
	if strings.Contains(report, "| `dbo.T` | — | `NO_IDENTIFIER_EVIDENCE`") || !strings.Contains(report, "OBSERVED_JOIN_COLUMN_SIGNAL") || !strings.Contains(report, "Table tanpa key, `IDENTITY`, atau sinyal ID dari implementasi: **0**") {
		t.Fatalf("contradictory inventory/summary: %s", report)
	}
}
