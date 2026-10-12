package diff

import "testing"

func TestComputeDrift(t *testing.T) {
	t.Parallel()

	desired := []State{
		{Category: "tables", Object: "dbo.TableA", NormalizedHash: "A", SourceFile: "db_object/tables/dbo.TableA.sql"},
		{Category: "views", Object: "dbo.ViewA", NormalizedHash: "B", SourceFile: "db_object/views/dbo.ViewA.sql"},
	}
	history := []State{
		{Category: "tables", Object: "dbo.TableA", NormalizedHash: "X"},
		{Category: "functions", Object: "dbo.FuncA", NormalizedHash: "C"},
	}

	drifts := ComputeDrift(desired, history)
	if len(drifts) != 3 {
		t.Fatalf("expected 3 drift entries, got %d: %+v", len(drifts), drifts)
	}

	if drifts[0].Kind != DriftHashMismatch || drifts[0].Object != "dbo.TableA" {
		t.Fatalf("unexpected first drift: %+v", drifts[0])
	}
	if drifts[1].Kind != DriftMissingInHistory || drifts[1].Object != "dbo.ViewA" {
		t.Fatalf("unexpected second drift: %+v", drifts[1])
	}
	if drifts[2].Kind != DriftMissingInDesired || drifts[2].Object != "dbo.FuncA" {
		t.Fatalf("unexpected third drift: %+v", drifts[2])
	}
}

func TestLessHelpers(t *testing.T) {
	t.Parallel()

	order := BuildOrderIndex([]string{"tables", "indexes", "views"})
	if !LessCategoryObject(order, 3, "tables", "dbo.TableA", "views", "dbo.ViewA") {
		t.Fatalf("expected tables before views")
	}
	if !LessPlanned(order, 3, "tables", "dbo.TableA", 1, "tables", "dbo.TableA", 2) {
		t.Fatalf("expected lower vk first for same object/category")
	}
}
