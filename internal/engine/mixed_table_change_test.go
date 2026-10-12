package engine

import (
	"strings"
	"testing"
)

func TestAddedColumnAndCheckChangeRoundTripTogether(t *testing.T) {
	previous := `CREATE TABLE dbo.MixedChange (
  [Id] INT NOT NULL,
  [Status] CHAR(1) NOT NULL,
  CONSTRAINT [CK_MixedChange_Status] CHECK ([Status] IN ('N', 'P'))
);`
	next := `CREATE TABLE dbo.MixedChange (
  [Id] INT NOT NULL,
  [Status] CHAR(1) NOT NULL,
  [Note] VARCHAR(100) NOT NULL CONSTRAINT [DF_MixedChange_Note] DEFAULT (''),
  CONSTRAINT [CK_MixedChange_Status] CHECK ([Status] IN ('N', 'P', 'C'))
);`

	changes, err := classifyTableColumnChanges(previous, next)
	if err != nil {
		t.Fatal(err)
	}
	alterUp, alterDown, _, err := buildTableAlterColumnsSQL(previous, next, "dbo.MixedChange", changes)
	if err != nil {
		t.Fatal(err)
	}
	constraintUp, constraintDown, ok, err := buildNonDefaultConstraintLifecycleWithDependencies(previous, next, "dbo.MixedChange", nil)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("expected changed CHECK constraint lifecycle")
	}

	up := concatSQLSegments(alterUp, constraintUp)
	down := concatSQLSegments(constraintDown, alterDown)
	if strings.Count(up, "DF_MixedChange_Note") != 1 {
		t.Fatalf("default constraint must be emitted exactly once, got SQL:\n%s", up)
	}
	if !strings.Contains(up, "CK_MixedChange_Status") {
		t.Fatalf("changed CHECK constraint was not emitted:\n%s", up)
	}

	history := map[string]historyObject{
		objectKey("tables", "dbo.MixedChange"): {
			Category:       "tables",
			Object:         "dbo.MixedChange",
			SemanticSQL:    previous,
			NormalizedHash: hashSemanticByCategory("tables", previous),
		},
	}
	desired := []desiredObject{{
		Category:       "tables",
		Object:         "dbo.MixedChange",
		SemanticSQL:    next,
		NormalizedHash: hashSemanticByCategory("tables", next),
	}}
	ops := []plannedMigration{{
		Category:        "tables",
		Object:          "dbo.MixedChange",
		UpSemanticSQL:   up,
		DownSemanticSQL: down,
	}}
	if err := validatePlannedRoundTrip(history, desired, ops); err != nil {
		t.Fatalf("mixed table change must round-trip: %v\nup:\n%s\ndown:\n%s", err, up, down)
	}
}

func TestAlterUniqueColumnAddColumnsAndChangeCheckRoundTripTogether(t *testing.T) {
	previous := `CREATE TABLE dbo.MixedChange (
  [Id] INT NOT NULL,
  [ExternalRef] VARCHAR(40) NOT NULL,
  [Status] VARCHAR(16) NOT NULL,
  CONSTRAINT [UQ_MixedChange_ExternalRef] UNIQUE ([ExternalRef]),
  CONSTRAINT [CK_MixedChange_Status] CHECK ([Status] IN ('NEW', 'DONE'))
);`
	next := `CREATE TABLE dbo.MixedChange (
  [Id] INT NOT NULL,
  [ExternalRef] VARCHAR(64) NOT NULL,
  [Status] VARCHAR(16) NOT NULL,
  [Currency] CHAR(3) NOT NULL CONSTRAINT [DF_MixedChange_Currency] DEFAULT ('IDR'),
  [ModifiedAt] DATETIME2(3) NULL,
  CONSTRAINT [UQ_MixedChange_ExternalRef] UNIQUE ([ExternalRef]),
  CONSTRAINT [CK_MixedChange_Status] CHECK ([Status] IN ('NEW', 'DONE', 'CANCELLED'))
);`

	changes, err := classifyTableColumnChanges(previous, next)
	if err != nil {
		t.Fatal(err)
	}
	depPlan, err := buildDropColumnDependencyPlan("dbo.MixedChange", changes.Changed, previous, nil)
	if err != nil {
		t.Fatal(err)
	}
	depPlan.Constraints = filterAutoManagedAlterColumnConstraints(depPlan.Constraints, "dbo.MixedChange")
	if hasUnsafeAlterColumnDependency(depPlan) {
		t.Fatalf("expected dependency plan to be safe: %+v", depPlan)
	}
	alterUp, alterDown, _, err := buildTableAlterColumnsSQL(previous, next, "dbo.MixedChange", changes)
	if err != nil {
		t.Fatal(err)
	}
	constraintUp, constraintDown, _, err := buildNonDefaultConstraintLifecycleWithDependencies(previous, next, "dbo.MixedChange", nil)
	if err != nil {
		t.Fatal(err)
	}
	restoreUp, err := buildReAddDroppedDependenciesSQL(next, "dbo.MixedChange", depPlan, nil)
	if err != nil {
		t.Fatal(err)
	}
	restoreDown, err := buildReAddDroppedDependenciesSQL(previous, "dbo.MixedChange", depPlan, nil)
	if err != nil {
		t.Fatal(err)
	}
	up := concatSQLSegments(buildDropDependencySQL(depPlan), alterUp, constraintUp, restoreUp)
	down := concatSQLSegments(buildDropDependencySQL(depPlan), constraintDown, alterDown, restoreDown)

	history := map[string]historyObject{
		objectKey("tables", "dbo.MixedChange"): {
			Category:       "tables",
			Object:         "dbo.MixedChange",
			SemanticSQL:    previous,
			NormalizedHash: hashSemanticByCategory("tables", previous),
		},
	}
	desired := []desiredObject{{
		Category:       "tables",
		Object:         "dbo.MixedChange",
		SemanticSQL:    next,
		NormalizedHash: hashSemanticByCategory("tables", next),
	}}
	ops := []plannedMigration{{Category: "tables", Object: "dbo.MixedChange", UpSemanticSQL: up, DownSemanticSQL: down}}
	if err := validatePlannedRoundTrip(history, desired, ops); err != nil {
		reconstructed, applyErr := semanticFromMigrationWithState(migrationMeta{Category: "tables", Object: "dbo.MixedChange"}, up, previous, true)
		t.Fatalf("mixed ALTER/ADD/constraint change must round-trip: %v\napplyErr=%v\nreconstructed:\n%s\ndesired:\n%s\nup:\n%s", err, applyErr, reconstructed, next, up)
	}
}
