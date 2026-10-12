package engine

import (
	"strings"
	"testing"
)

func TestColumnRenameDirectivePlanAndSemanticRoundTrip(t *testing.T) {
	previous := `CREATE TABLE dbo.Customer ([Id] INT NOT NULL, [OldName] VARCHAR(40) NULL CONSTRAINT [DF_Customer_Name] DEFAULT ('x'), CONSTRAINT [CK_Customer_Name] CHECK ([OldName] <> 'bad'));`
	next := `CREATE TABLE dbo.Customer ([Id] INT NOT NULL, [NewName] VARCHAR(40) NULL CONSTRAINT [DF_Customer_Name] DEFAULT ('x'), CONSTRAINT [CK_Customer_Name] CHECK ([NewName] <> 'bad'));`
	hints, err := parseColumnRenameDirectives("-- rza:rename-column OldName -> NewName\n" + next)
	if err != nil {
		t.Fatal(err)
	}
	up, down, ok, err := planColumnRenames(previous, next, "dbo.Customer", hints)
	if err != nil || !ok {
		t.Fatalf("rename plan failed: ok=%v err=%v", ok, err)
	}
	if !strings.Contains(up, "sp_rename") || !strings.Contains(up, "OldName") || !strings.Contains(down, "NewName") {
		t.Fatalf("reversible rename SQL missing:\nUP:\n%s\nDOWN:\n%s", up, down)
	}
	reconstructed, err := semanticFromMigrationWithState(migrationMeta{Category: "tables", Object: "dbo.Customer"}, up, previous, true)
	if err != nil {
		t.Fatal(err)
	}
	if hashSemanticByCategory("tables", reconstructed) != hashSemanticByCategory("tables", next) {
		t.Fatalf("rename semantic round-trip mismatch:\n%s", reconstructed)
	}
	reversed, err := semanticFromMigrationWithState(migrationMeta{Category: "tables", Object: "dbo.Customer"}, down, reconstructed, true)
	if err != nil {
		t.Fatal(err)
	}
	if hashSemanticByCategory("tables", reversed) != hashSemanticByCategory("tables", previous) {
		t.Fatalf("rename down semantic mismatch:\n%s", reversed)
	}
}

func TestColumnRenameRejectsCombinedTypeChange(t *testing.T) {
	previous := `CREATE TABLE dbo.Customer ([OldName] VARCHAR(40) NULL);`
	next := `CREATE TABLE dbo.Customer ([NewName] VARCHAR(20) NULL);`
	_, _, _, err := planColumnRenames(previous, next, "dbo.Customer", map[string]string{"OldName": "NewName"})
	if err == nil {
		t.Fatal("rename plus type change must be split")
	}
}

func TestColumnRenameRejectsDefaultLiteralChangeAndPreservesLiteral(t *testing.T) {
	previous := `CREATE TABLE dbo.Customer ([OldName] VARCHAR(40) NULL CONSTRAINT [DF_Customer_Name] DEFAULT ('OldName'));`
	next := `CREATE TABLE dbo.Customer ([NewName] VARCHAR(40) NULL CONSTRAINT [DF_Customer_Name] DEFAULT ('NewName'));`
	_, _, _, err := planColumnRenames(previous, next, "dbo.Customer", map[string]string{"OldName": "NewName"})
	if err == nil {
		t.Fatal("rename plus default literal change must be rejected")
	}

	rewritten := replaceColumnIdentifiers(previous, []columnRename{{Old: "OldName", New: "NewName"}})
	if !strings.Contains(rewritten, "DEFAULT ('OldName')") || !strings.Contains(rewritten, "[NewName]") {
		t.Fatalf("identifier-aware rewrite corrupted literal: %s", rewritten)
	}
}
