package engine

import (
	"reflect"
	"strings"
	"testing"
)

func TestOrderPlannedMigrationsFollowsCrossCategoryDependencies(t *testing.T) {
	t.Parallel()
	ops := []plannedMigration{
		{Category: "views", Object: "dbo.ViewA"},
		{Category: "procedures", Object: "dbo.ProcA"},
		{Category: "views", Object: "dbo.ViewB"},
		{Category: "functions", Object: "dbo.fnValue"},
		{Category: "procedures", Object: "dbo.ProcB"},
	}
	desired := map[string]desiredObject{
		objectKey("views", "dbo.ViewA"):       {Category: "views", Object: "dbo.ViewA", SemanticSQL: "CREATE VIEW dbo.ViewA AS SELECT * FROM dbo.ViewB;"},
		objectKey("views", "dbo.ViewB"):       {Category: "views", Object: "dbo.ViewB", SemanticSQL: "CREATE VIEW dbo.ViewB AS SELECT dbo.fnValue(1) AS id;"},
		objectKey("functions", "dbo.fnValue"): {Category: "functions", Object: "dbo.fnValue", SemanticSQL: "CREATE FUNCTION dbo.fnValue(@id INT) RETURNS INT AS BEGIN RETURN @id; END"},
		objectKey("procedures", "dbo.ProcA"):  {Category: "procedures", Object: "dbo.ProcA", SemanticSQL: "CREATE PROCEDURE dbo.ProcA AS EXEC dbo.ProcB;"},
		objectKey("procedures", "dbo.ProcB"):  {Category: "procedures", Object: "dbo.ProcB", SemanticSQL: "CREATE PROCEDURE dbo.ProcB AS SELECT 1;"},
	}

	result, err := orderPlannedMigrations(ops, desired, []string{"tables", "indexes", "functions", "views", "procedures", "patch"})
	if err != nil {
		t.Fatalf("orderPlannedMigrations error: %v", err)
	}
	got := make([]string, 0, len(result.Ops))
	for _, op := range result.Ops {
		got = append(got, op.Object)
	}
	want := []string{"dbo.fnValue", "dbo.ViewB", "dbo.ViewA", "dbo.ProcB", "dbo.ProcA"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected dependency order: got=%v want=%v", got, want)
	}
}

func TestOrderPlannedMigrationsRejectsCycle(t *testing.T) {
	t.Parallel()
	ops := []plannedMigration{{Category: "views", Object: "dbo.ViewA"}, {Category: "views", Object: "dbo.ViewB"}}
	desired := map[string]desiredObject{
		objectKey("views", "dbo.ViewA"): {Category: "views", Object: "dbo.ViewA", SemanticSQL: "CREATE VIEW dbo.ViewA AS SELECT * FROM dbo.ViewB;"},
		objectKey("views", "dbo.ViewB"): {Category: "views", Object: "dbo.ViewB", SemanticSQL: "CREATE VIEW dbo.ViewB AS SELECT * FROM dbo.ViewA;"},
	}
	if _, err := orderPlannedMigrations(ops, desired, []string{"views"}); err == nil {
		t.Fatal("expected dependency cycle error")
	}
}

func TestOrderPlannedMigrationsCreatesReferencedTableBeforeForeignKeyTable(t *testing.T) {
	t.Parallel()

	// The dependent table deliberately sorts before the referenced table.
	// Dependency ordering must override alphabetical ordering.
	ops := []plannedMigration{
		{Category: "tables", Object: "dbo.AChild"},
		{Category: "tables", Object: "dbo.ZParent"},
		{Category: "tables", Object: "dbo.BGrandChild"},
	}
	desired := map[string]desiredObject{
		objectKey("tables", "dbo.AChild"): {
			Category: "tables",
			Object:   "dbo.AChild",
			SemanticSQL: `CREATE TABLE dbo.AChild (
  EventId INT NOT NULL,
  CONSTRAINT [FK_AChild_ZParent] FOREIGN KEY(EventId) REFERENCES dbo.ZParent([EventId])
);`,
		},
		objectKey("tables", "dbo.ZParent"): {
			Category:    "tables",
			Object:      "dbo.ZParent",
			SemanticSQL: `CREATE TABLE dbo.ZParent ([EventId] INT NOT NULL);`,
		},
		objectKey("tables", "dbo.BGrandChild"): {
			Category: "tables",
			Object:   "dbo.BGrandChild",
			SemanticSQL: `CREATE TABLE dbo.BGrandChild (
  ChildId INT NOT NULL,
  CONSTRAINT [FK_BGrandChild_AChild] FOREIGN KEY (ChildId) REFERENCES [dbo].[AChild] ([EventId])
);`,
		},
	}

	result, err := orderPlannedMigrations(ops, desired, []string{"tables"})
	if err != nil {
		t.Fatalf("orderPlannedMigrations error: %v", err)
	}
	got := make([]string, 0, len(result.Ops))
	for _, op := range result.Ops {
		got = append(got, op.Object)
	}
	want := []string{"dbo.ZParent", "dbo.AChild", "dbo.BGrandChild"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected table dependency order: got=%v want=%v edges=%v", got, want, result.Edges)
	}
	if result.ResolvedReferences != 2 {
		t.Fatalf("expected two resolved FK references, got=%d edges=%v", result.ResolvedReferences, result.Edges)
	}
}

func TestOrderPlannedMigrationsRejectsMutualForeignKeyCycle(t *testing.T) {
	t.Parallel()

	ops := []plannedMigration{
		{Category: "tables", Object: "dbo.TableA"},
		{Category: "tables", Object: "dbo.TableB"},
	}
	desired := map[string]desiredObject{
		objectKey("tables", "dbo.TableA"): {
			Category:    "tables",
			Object:      "dbo.TableA",
			SemanticSQL: `CREATE TABLE dbo.TableA (BId INT, CONSTRAINT [FK_TableA_TableB] FOREIGN KEY (BId) REFERENCES dbo.TableB([Id]));`,
		},
		objectKey("tables", "dbo.TableB"): {
			Category:    "tables",
			Object:      "dbo.TableB",
			SemanticSQL: `CREATE TABLE dbo.TableB (Id INT, AId INT, CONSTRAINT [FK_TableB_TableA] FOREIGN KEY (AId) REFERENCES dbo.TableA([BId]));`,
		},
	}

	if _, err := orderPlannedMigrations(ops, desired, []string{"tables"}); err == nil ||
		!strings.Contains(err.Error(), "dependency cycle detected") {
		t.Fatalf("mutual inline foreign keys must fail as an explicit dependency cycle, got=%v", err)
	}
}

func TestOrderPlannedMigrationsAllowsSelfReferencingForeignKey(t *testing.T) {
	t.Parallel()

	ops := []plannedMigration{
		{Category: "tables", Object: "dbo.Category"},
		{Category: "tables", Object: "dbo.Other"},
	}
	desired := map[string]desiredObject{
		objectKey("tables", "dbo.Category"): {
			Category: "tables",
			Object:   "dbo.Category",
			SemanticSQL: `CREATE TABLE dbo.Category (
  Id INT NOT NULL,
  ParentId INT NULL,
  CONSTRAINT [FK_Category_Category] FOREIGN KEY (ParentId) REFERENCES dbo.Category([Id])
);`,
		},
		objectKey("tables", "dbo.Other"): {
			Category:    "tables",
			Object:      "dbo.Other",
			SemanticSQL: `CREATE TABLE dbo.Other (Id INT NOT NULL);`,
		},
	}

	result, err := orderPlannedMigrations(ops, desired, []string{"tables"})
	if err != nil {
		t.Fatalf("self-reference must not create a graph cycle: %v", err)
	}
	if result.ResolvedReferences != 0 || len(result.Edges) != 0 {
		t.Fatalf("self-reference must not create an inter-migration edge: %+v", result)
	}
}
