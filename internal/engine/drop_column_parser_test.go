package engine

import "testing"

func TestExtractDroppedColumns(t *testing.T) {
	t.Parallel()

	sql := `
ALTER TABLE dbo.TableA DROP COLUMN amount;
ALTER TABLE dbo.TableA DROP COLUMN [legacy_code], status;
`

	cols, err := extractDroppedColumns("dbo.TableA", sql)
	if err != nil {
		t.Fatalf("extractDroppedColumns returned error: %v", err)
	}
	want := map[string]bool{
		"amount":      true,
		"legacy_code": true,
		"status":      true,
	}
	if len(cols) != len(want) {
		t.Fatalf("unexpected column count: got=%d want=%d cols=%#v", len(cols), len(want), cols)
	}
	for _, c := range cols {
		if !want[c] {
			t.Fatalf("unexpected dropped column %s", c)
		}
	}
}

func TestClassifyTableColumnChangesPreservesRemovedColumnCase(t *testing.T) {
	t.Parallel()

	prev := `CREATE TABLE dbo.TableA (
  [ID] INT NOT NULL,
  [RealCaseNote] VARCHAR(32) NULL
);`
	next := `CREATE TABLE dbo.TableA (
  [ID] INT NOT NULL
);`
	changes, err := classifyTableColumnChanges(prev, next)
	if err != nil {
		t.Fatalf("classifyTableColumnChanges failed: %v", err)
	}
	if len(changes.Removed) != 1 || changes.Removed[0] != "RealCaseNote" {
		t.Fatalf("expected removed column case to be preserved, got %#v", changes.Removed)
	}
}

func TestApplyDropColumnsToCreateTable(t *testing.T) {
	t.Parallel()

	current := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL,
  amount DECIMAL(10,2) NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id),
  CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status
);
`

	next, err := applyDropColumnsToCreateTable(current, []string{"amount", "status"})
	if err != nil {
		t.Fatalf("applyDropColumnsToCreateTable returned error: %v", err)
	}
	if containsColumnReference(next, "amount") {
		t.Fatalf("result still references removed column amount: %s", next)
	}
	if containsColumnReference(next, "status") {
		t.Fatalf("result still references removed column status: %s", next)
	}
	if !containsColumnReference(next, "id") {
		t.Fatalf("result should still contain id column: %s", next)
	}
}
