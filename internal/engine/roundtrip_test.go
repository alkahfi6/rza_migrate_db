package engine

import "testing"

func TestValidatePlannedRoundTripCreateTable(t *testing.T) {
	desired := []desiredObject{{Category: "tables", Object: "dbo.T", SemanticSQL: "CREATE TABLE dbo.T ([id] INT NOT NULL);\n"}}
	desired[0].NormalizedHash = hashSemanticByCategory("tables", desired[0].SemanticSQL)
	ops := []plannedMigration{{Category: "tables", Object: "dbo.T", UpSemanticSQL: desired[0].SemanticSQL}}
	if err := validatePlannedRoundTrip(map[string]historyObject{}, desired, ops); err != nil {
		t.Fatalf("round trip failed: %v", err)
	}
}

func TestValidatePlannedRoundTripRejectsMismatch(t *testing.T) {
	desired := []desiredObject{{Category: "tables", Object: "dbo.T", SemanticSQL: "CREATE TABLE dbo.T ([id] INT NOT NULL);\n"}}
	desired[0].NormalizedHash = hashSemanticByCategory("tables", desired[0].SemanticSQL)
	ops := []plannedMigration{{Category: "tables", Object: "dbo.T", UpSemanticSQL: "CREATE TABLE dbo.T ([other] INT NOT NULL);\n"}}
	if err := validatePlannedRoundTrip(map[string]historyObject{}, desired, ops); err == nil {
		t.Fatal("expected mismatch")
	}
}
