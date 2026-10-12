package engine

import (
	"strings"
	"testing"
)

func TestBuildAddKeyConstraintWithPrecheckUniqueAndPrimaryKey(t *testing.T) {
	unique, err := buildAddKeyConstraintWithPrecheckStatement("dbo.Customer", "UQ_Customer_Email", "unique", "CONSTRAINT [UQ_Customer_Email] UNIQUE ([Email])")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"GROUP BY [EMAIL]", "HAVING COUNT_BIG(*) > 1", "ALTER TABLE DBO.CUSTOMER ADD CONSTRAINT [UQ_CUSTOMER_EMAIL] UNIQUE ([EMAIL])"} {
		if !strings.Contains(strings.ToUpper(unique), want) {
			t.Fatalf("unique guard missing %q:\n%s", want, unique)
		}
	}

	primary, err := buildAddKeyConstraintWithPrecheckStatement("dbo.Customer", "PK_Customer", "primary_key", "CONSTRAINT [PK_Customer] PRIMARY KEY CLUSTERED ([Id], [TenantId])")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"[ID] IS NULL OR [TENANTID] IS NULL", "GROUP BY [ID], [TENANTID]", "PRIMARY KEY PK_CUSTOMER CANNOT BE ADDED"} {
		if !strings.Contains(strings.ToUpper(primary), want) {
			t.Fatalf("primary key guard missing %q:\n%s", want, primary)
		}
	}
}

func TestConstraintLifecycleSupportsNamedUnique(t *testing.T) {
	before := `CREATE TABLE dbo.Customer ([Id] INT NOT NULL);`
	after := `CREATE TABLE dbo.Customer ([Id] INT NOT NULL, CONSTRAINT [UQ_Customer_Id] UNIQUE ([Id]));`
	up, down, ok, err := buildCheckConstraintOnlyLifecycleSQL(before, after, "dbo.Customer")
	if err != nil || !ok {
		t.Fatalf("unique lifecycle failed: ok=%v err=%v", ok, err)
	}
	if !strings.Contains(strings.ToUpper(up), "GROUP BY [ID]") || !strings.Contains(strings.ToUpper(down), "DROP CONSTRAINT [UQ_CUSTOMER_ID]") {
		t.Fatalf("unexpected lifecycle:\nUP:\n%s\nDOWN:\n%s", up, down)
	}
}

func TestBuildAddForeignKeyWithCompositeOrphanPrecheck(t *testing.T) {
	definition := "CONSTRAINT [FK_Child_Parent] FOREIGN KEY ([ParentId], [TenantId]) REFERENCES dbo.Parent ([Id], [TenantId]) ON DELETE CASCADE"
	sql, err := buildAddForeignKeyWithPrecheckStatement("dbo.Child", "FK_Child_Parent", definition)
	if err != nil {
		t.Fatal(err)
	}
	upper := strings.ToUpper(sql)
	for _, want := range []string{"C.[PARENTID] IS NOT NULL AND C.[TENANTID] IS NOT NULL", "P.[ID] = C.[PARENTID] AND P.[TENANTID] = C.[TENANTID]", "ORPHAN ROWS EXIST", "ON DELETE CASCADE"} {
		if !strings.Contains(upper, want) {
			t.Fatalf("foreign key guard missing %q:\n%s", want, sql)
		}
	}
}

func TestConstraintLifecycleSupportsNamedForeignKey(t *testing.T) {
	before := `CREATE TABLE dbo.Child ([ParentId] INT NULL);`
	after := `CREATE TABLE dbo.Child ([ParentId] INT NULL, CONSTRAINT [FK_Child_Parent] FOREIGN KEY ([ParentId]) REFERENCES dbo.Parent ([Id]));`
	up, down, ok, err := buildCheckConstraintOnlyLifecycleSQL(before, after, "dbo.Child")
	if err != nil || !ok {
		t.Fatalf("foreign key lifecycle failed: ok=%v err=%v", ok, err)
	}
	if !strings.Contains(strings.ToUpper(up), "ORPHAN ROWS EXIST") || !strings.Contains(strings.ToUpper(down), "DROP CONSTRAINT [FK_CHILD_PARENT]") {
		t.Fatalf("unexpected FK lifecycle:\nUP:\n%s\nDOWN:\n%s", up, down)
	}
}

func TestKeyConstraintLifecycleTemporarilyRestoresIncomingForeignKeys(t *testing.T) {
	previous := `CREATE TABLE dbo.Parent ([Id] INT NOT NULL, CONSTRAINT [UQ_Parent_Id] UNIQUE ([Id]));`
	next := `CREATE TABLE dbo.Parent ([Id] INT NOT NULL, CONSTRAINT [PK_Parent] PRIMARY KEY ([Id]));`
	history := map[string]historyObject{
		objectKey("tables", "dbo.Child"): {Category: "tables", Object: "dbo.Child", SemanticSQL: `CREATE TABLE dbo.Child ([ParentId] INT NULL, CONSTRAINT [FK_Child_Parent] FOREIGN KEY ([ParentId]) REFERENCES dbo.Parent ([Id]));`},
	}
	up, down, ok, err := buildConstraintLifecycleWithDependencies(previous, next, "dbo.Parent", history)
	if err != nil || !ok {
		t.Fatalf("key lifecycle with FK failed: ok=%v err=%v", ok, err)
	}
	for _, sql := range []string{up, down} {
		upper := strings.ToUpper(sql)
		if !strings.Contains(upper, "ALTER TABLE DBO.CHILD DROP CONSTRAINT [FK_CHILD_PARENT]") || !strings.Contains(upper, "ORPHAN ROWS EXIST") {
			t.Fatalf("incoming FK was not safely cycled:\n%s", sql)
		}
	}
}
