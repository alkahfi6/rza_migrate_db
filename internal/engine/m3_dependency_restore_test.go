package engine

import (
	"strings"
	"testing"
)

func TestBuildAlterColumnDependencyBlockList(t *testing.T) {
	t.Parallel()

	prev := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id),
  CONSTRAINT CK_TableA__status CHECK (status IN ('NEW','DONE'))
);
`
	history := map[string]historyObject{
		objectKey("tables", "dbo.TableA"): {
			Category:    "tables",
			Object:      "dbo.TableA",
			SemanticSQL: prev,
		},
		objectKey("indexes", "dbo.TableA"): {
			Category:    "indexes",
			Object:      "dbo.TableA",
			SemanticSQL: "CREATE INDEX IX_TableA__status ON dbo.TableA(status);",
		},
	}

	blocked, err := buildAlterColumnDependencyBlockList("dbo.TableA", []string{"status"}, prev, history)
	if err != nil {
		t.Fatalf("buildAlterColumnDependencyBlockList returned error: %v", err)
	}
	if len(blocked) == 0 {
		t.Fatalf("expected blocked dependencies for ALTER COLUMN")
	}
	joined := strings.ToUpper(strings.Join(blocked, ","))
	if !strings.Contains(joined, "CONSTRAINT:DBO.TABLEA.CK_TABLEA__STATUS") {
		t.Fatalf("missing expected constraint blocker: %v", blocked)
	}
	if !strings.Contains(joined, "INDEX:DBO.TABLEA.IX_TABLEA__STATUS") {
		t.Fatalf("missing expected index blocker: %v", blocked)
	}
}

func TestBuildDropColumnDownRestoreSQLRestoresDependencies(t *testing.T) {
	t.Parallel()

	prev := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id),
  CONSTRAINT CK_TableA__status CHECK (status IN ('NEW','DONE'))
);
`
	history := map[string]historyObject{
		objectKey("tables", "dbo.TableA"): {
			Category:    "tables",
			Object:      "dbo.TableA",
			SemanticSQL: prev,
		},
		objectKey("indexes", "dbo.TableA"): {
			Category:    "indexes",
			Object:      "dbo.TableA",
			SemanticSQL: "CREATE INDEX IX_TableA__status ON dbo.TableA(status);",
		},
	}
	depPlan := dropColumnDependencyPlan{
		Constraints: []dropConstraintTarget{
			{Table: "dbo.TableA", Name: "CK_TableA__status", Kind: "check"},
		},
		Indexes: []dropIndexTarget{
			{Table: "dbo.TableA", Name: "IX_TableA__status"},
		},
	}

	downSQL, err := buildDropColumnDownRestoreSQL(prev, "dbo.TableA", []string{"status"}, depPlan, history)
	if err != nil {
		t.Fatalf("buildDropColumnDownRestoreSQL returned error: %v", err)
	}
	if !strings.Contains(strings.ToUpper(downSQL), "EXEC SYS.SP_EXECUTESQL") {
		t.Fatalf("dependency restore after ADD COLUMN must use a later compilation batch: %s", downSQL)
	}
	batches := extractSpExecuteSQLBatches(downSQL)
	if len(batches) != 1 {
		t.Fatalf("expected one decoded dependency restore batch, got %d: %s", len(batches), downSQL)
	}
	upper := strings.ToUpper(downSQL + "\n" + batches[0])
	required := []string{
		"ALTER TABLE DBO.TABLEA ADD [STATUS] VARCHAR(10) NULL;",
		"ALTER TABLE DBO.TABLEA ADD CONSTRAINT CK_TABLEA__STATUS CHECK (STATUS IN ('NEW','DONE'));",
		"CREATE INDEX IX_TABLEA__STATUS ON DBO.TABLEA(STATUS);",
	}
	for _, token := range required {
		if !strings.Contains(upper, token) {
			t.Fatalf("down restore SQL missing token %q in SQL: %s", token, downSQL)
		}
	}
}

func TestAlterColumnWithNamedCheckAndIndexCanDropAndRestoreDependencies(t *testing.T) {
	t.Parallel()

	prev := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id),
  CONSTRAINT CK_TableA__status CHECK (status IN ('NEW','DONE'))
);
`
	next := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(20) NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id),
  CONSTRAINT CK_TableA__status CHECK (status IN ('NEW','DONE'))
);
`
	history := map[string]historyObject{
		objectKey("tables", "dbo.TableA"): {
			Category:    "tables",
			Object:      "dbo.TableA",
			SemanticSQL: prev,
		},
		objectKey("indexes", "dbo.TableA"): {
			Category:    "indexes",
			Object:      "dbo.TableA",
			SemanticSQL: "CREATE INDEX IX_TableA__status ON dbo.TableA(status);",
		},
	}
	changes, err := classifyTableColumnChanges(prev, next)
	if err != nil {
		t.Fatalf("classify changes failed: %v", err)
	}
	depPlan, err := buildDropColumnDependencyPlan("dbo.TableA", changes.Changed, prev, history)
	if err != nil {
		t.Fatalf("dependency plan failed: %v", err)
	}
	depPlan.Constraints = filterAutoManagedAlterColumnConstraints(depPlan.Constraints, "dbo.TableA")
	if hasUnsafeAlterColumnDependency(depPlan) {
		t.Fatalf("expected named CHECK + index dependency to be auto-manageable")
	}
	alterUp, _, _, err := buildTableAlterColumnsSQL(prev, next, "dbo.TableA", changes)
	if err != nil {
		t.Fatalf("build alter sql failed: %v", err)
	}
	restore, err := buildReAddDroppedDependenciesSQL(prev, "dbo.TableA", depPlan, history)
	if err != nil {
		t.Fatalf("restore dependency SQL failed: %v", err)
	}
	upSQL := concatSQLSegments(buildDropDependencySQL(depPlan), alterUp, restore)
	upper := strings.ToUpper(upSQL)
	required := []string{
		"ALTER TABLE DBO.TABLEA DROP CONSTRAINT [CK_TABLEA__STATUS];",
		"DROP INDEX IF EXISTS [IX_TABLEA__STATUS] ON DBO.TABLEA;",
		"ALTER TABLE DBO.TABLEA ALTER COLUMN [STATUS] VARCHAR(20) NULL;",
		"ALTER TABLE DBO.TABLEA ADD CONSTRAINT CK_TABLEA__STATUS CHECK (STATUS IN ('NEW','DONE'));",
		"CREATE INDEX IX_TABLEA__STATUS ON DBO.TABLEA(STATUS);",
	}
	for _, token := range required {
		if !strings.Contains(upper, token) {
			t.Fatalf("alter dependency SQL missing %q in SQL: %s", token, upSQL)
		}
	}
}

func TestAlterColumnWithNamedUniqueCanDropAndRestoreDependency(t *testing.T) {
	t.Parallel()

	prev := `CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  external_ref VARCHAR(40) NOT NULL,
  CONSTRAINT UQ_TableA__external_ref UNIQUE (external_ref)
);`
	next := `CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  external_ref VARCHAR(64) NOT NULL,
  CONSTRAINT UQ_TableA__external_ref UNIQUE (external_ref)
);`

	changes, err := classifyTableColumnChanges(prev, next)
	if err != nil {
		t.Fatalf("classify changes failed: %v", err)
	}
	depPlan, err := buildDropColumnDependencyPlan("dbo.TableA", changes.Changed, prev, map[string]historyObject{})
	if err != nil {
		t.Fatalf("dependency plan failed: %v", err)
	}
	depPlan.Constraints = filterAutoManagedAlterColumnConstraints(depPlan.Constraints, "dbo.TableA")
	if hasUnsafeAlterColumnDependency(depPlan) {
		t.Fatalf("expected named UNIQUE dependency to be auto-manageable: %+v", depPlan)
	}

	alterUp, _, _, err := buildTableAlterColumnsSQL(prev, next, "dbo.TableA", changes)
	if err != nil {
		t.Fatalf("build alter SQL failed: %v", err)
	}
	restoreUp, err := buildReAddDroppedDependenciesSQL(next, "dbo.TableA", depPlan, map[string]historyObject{})
	if err != nil {
		t.Fatalf("build restore SQL failed: %v", err)
	}
	upSQL := strings.ToUpper(concatSQLSegments(buildDropDependencySQL(depPlan), alterUp, restoreUp))
	requiredInOrder := []string{
		"ALTER TABLE DBO.TABLEA DROP CONSTRAINT [UQ_TABLEA__EXTERNAL_REF];",
		"ALTER TABLE DBO.TABLEA ALTER COLUMN [EXTERNAL_REF] VARCHAR(64) NOT NULL;",
		"ALTER TABLE DBO.TABLEA ADD CONSTRAINT UQ_TABLEA__EXTERNAL_REF UNIQUE (EXTERNAL_REF);",
	}
	last := -1
	for _, token := range requiredInOrder {
		pos := strings.Index(upSQL, token)
		if pos < 0 {
			t.Fatalf("alter UNIQUE dependency SQL missing %q in SQL: %s", token, upSQL)
		}
		if pos <= last {
			t.Fatalf("alter UNIQUE dependency SQL has unsafe order for %q: %s", token, upSQL)
		}
		last = pos
	}
	if !strings.Contains(upSQL, "UNIQUE CONSTRAINT UQ_TABLEA__EXTERNAL_REF CANNOT BE ADDED BECAUSE DUPLICATE KEY VALUES EXIST") {
		t.Fatalf("expected duplicate-data precheck before restoring UNIQUE constraint: %s", upSQL)
	}
}

func TestRestoreInlineDefaultDependencyUsesTableLevelConstraint(t *testing.T) {
	t.Parallel()

	tableSQL := `CREATE TABLE dbo.TableA (
  Amount DECIMAL(19,4) NOT NULL CONSTRAINT DF_TableA_Amount DEFAULT (0)
);`
	plan, err := buildDropColumnDependencyPlan("dbo.TableA", []string{"Amount"}, tableSQL, nil)
	if err != nil {
		t.Fatal(err)
	}
	restore, err := buildReAddDroppedDependenciesSQL(tableSQL, "dbo.TableA", plan, nil)
	if err != nil {
		t.Fatal(err)
	}
	upper := strings.ToUpper(restore)
	if !strings.Contains(upper, "ALTER TABLE DBO.TABLEA ADD CONSTRAINT [DF_TABLEA_AMOUNT] DEFAULT (0) FOR [AMOUNT];") {
		t.Fatalf("default restore must use table-level DEFAULT syntax: %s", restore)
	}
	if strings.Contains(upper, "ADD [AMOUNT] DECIMAL") {
		t.Fatalf("default restore must not re-add the existing column: %s", restore)
	}
}

func TestNonDefaultLifecycleSkipsRemovedInlineDefault(t *testing.T) {
	t.Parallel()

	previous := `CREATE TABLE dbo.TableA (
  Amount DECIMAL(19,4) NOT NULL CONSTRAINT DF_TableA_Amount DEFAULT (0),
  CONSTRAINT CK_TableA_Amount CHECK (Amount >= 0)
);`
	next := `CREATE TABLE dbo.TableA (
  Amount DECIMAL(28,8) NOT NULL,
  CONSTRAINT CK_TableA_Amount CHECK (Amount >= 0)
);`

	up, down, _, err := buildNonDefaultConstraintLifecycleWithDependencies(previous, next, "dbo.TableA", nil)
	if err != nil {
		t.Fatal(err)
	}
	combined := strings.ToUpper(up + "\n" + down)
	if strings.Contains(combined, "DF_TABLEA_AMOUNT") || strings.Contains(combined, "ADD [AMOUNT] DECIMAL") {
		t.Fatalf("non-default lifecycle must not plan inline DEFAULT: %s", combined)
	}
}

func TestChangedDefaultIsNotRestoredAgainByDependencyPlan(t *testing.T) {
	t.Parallel()

	previous := `CREATE TABLE dbo.TableA (
  Amount DECIMAL(19,4) NOT NULL CONSTRAINT DF_TableA_Amount DEFAULT (0)
);`
	plan := dropColumnDependencyPlan{Constraints: []dropConstraintTarget{
		{Table: "dbo.TableA", Name: "DF_TableA_Amount", Kind: "default"},
		{Table: "dbo.TableA", Name: "CK_TableA_Amount", Kind: "check"},
	}}
	filtered, err := excludeColumnLifecycleDefaultsFromDependencyRestore(plan, previous, []string{"Amount"})
	if err != nil {
		t.Fatal(err)
	}
	if len(filtered.Constraints) != 1 || filtered.Constraints[0].Name != "CK_TableA_Amount" {
		t.Fatalf("changed default must be excluded while other dependencies remain: %+v", filtered.Constraints)
	}
}

func TestDynamicConstraintReconstructionDoesNotKeepEscapedQuotes(t *testing.T) {
	t.Parallel()

	previous := `CREATE TABLE dbo.TableA (
  Status VARCHAR(16) NOT NULL,
  CONSTRAINT CK_TableA_Status CHECK (Status IN ('NEW', 'DONE'))
);`
	body := `IF (OBJECT_ID(N'dbo.CK_TableA_Status', N'C') IS NOT NULL)
BEGIN
  ALTER TABLE dbo.TableA DROP CONSTRAINT [CK_TableA_Status];
END
ALTER TABLE dbo.TableA ADD CONSTRAINT CK_TableA_Status CHECK (Status IN ('PENDING', 'DONE'));`
	migration := dynamicSQLBatch(body)
	state, err := semanticFromMigrationWithState(migrationMeta{Category: "tables", Object: "dbo.TableA"}, migration, previous, true)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(state, "''PENDING''") {
		t.Fatalf("escaped dynamic-SQL quotes leaked into semantic state: %s", state)
	}
	if !strings.Contains(state, "Status IN ('PENDING', 'DONE')") {
		t.Fatalf("decoded CHECK definition missing from semantic state: %s", state)
	}
}

func TestBuildCheckConstraintOnlyLifecycleSQLWithPrecheck(t *testing.T) {
	t.Parallel()

	prev := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  amount DECIMAL(18,2) NOT NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id),
  CONSTRAINT CK_TableA__amount CHECK (amount >= 0)
);
`
	next := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  amount DECIMAL(18,2) NOT NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id),
  CONSTRAINT CK_TableA__amount CHECK (amount > 0)
);
`
	upSQL, downSQL, ok, err := buildCheckConstraintOnlyLifecycleSQL(prev, next, "dbo.TableA")
	if err != nil {
		t.Fatalf("build check lifecycle failed: %v", err)
	}
	if !ok {
		t.Fatalf("expected check constraint-only change to be auto-planned")
	}
	upUpper := strings.ToUpper(upSQL)
	for _, token := range []string{
		"WHERE NOT (AMOUNT > 0)",
		"THROW 51000, 'CHECK CONSTRAINT CK_TABLEA__AMOUNT WOULD FAIL EXISTING DATA', 1;",
		"ALTER TABLE DBO.TABLEA DROP CONSTRAINT [CK_TABLEA__AMOUNT];",
		"ALTER TABLE DBO.TABLEA ADD CONSTRAINT CK_TABLEA__AMOUNT CHECK (AMOUNT > 0);",
	} {
		if !strings.Contains(upUpper, token) {
			t.Fatalf("up SQL missing %q in SQL: %s", token, upSQL)
		}
	}
	if !strings.Contains(strings.ToUpper(downSQL), "WHERE NOT (AMOUNT >= 0)") {
		t.Fatalf("down SQL should precheck old CHECK before restore: %s", downSQL)
	}
}

func TestBuildConstraintOnlyLifecycleSQLDefaultEquivalentInlineVsTableLevel(t *testing.T) {
	t.Parallel()

	prev := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(20) NULL CONSTRAINT DF_TableA__status DEFAULT ('NEW')
);
`
	next := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(20) NULL,
  CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status
);
`
	upSQL, downSQL, ok, err := buildCheckConstraintOnlyLifecycleSQL(prev, next, "dbo.TableA")
	if err != nil {
		t.Fatalf("build default lifecycle failed: %v", err)
	}
	if ok || strings.TrimSpace(upSQL) != "" || strings.TrimSpace(downSQL) != "" {
		t.Fatalf("equivalent inline/table-level default should not emit SQL, ok=%v up=%q down=%q", ok, upSQL, downSQL)
	}
}

func TestBuildConstraintOnlyLifecycleSQLDefaultChange(t *testing.T) {
	t.Parallel()

	prev := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(20) NULL CONSTRAINT DF_TableA__status DEFAULT ('NEW')
);
`
	next := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(20) NULL CONSTRAINT DF_TableA__status DEFAULT ('DONE')
);
`
	upSQL, downSQL, ok, err := buildCheckConstraintOnlyLifecycleSQL(prev, next, "dbo.TableA")
	if err != nil {
		t.Fatalf("build default lifecycle failed: %v", err)
	}
	if !ok {
		t.Fatalf("expected default-only change to be auto-planned")
	}
	upUpper := strings.ToUpper(upSQL)
	for _, token := range []string{
		"ALTER TABLE DBO.TABLEA DROP CONSTRAINT [DF_TABLEA__STATUS];",
		"ALTER TABLE DBO.TABLEA ADD CONSTRAINT [DF_TABLEA__STATUS] DEFAULT ('DONE') FOR [STATUS];",
	} {
		if !strings.Contains(upUpper, token) {
			t.Fatalf("up SQL missing %q in SQL: %s", token, upSQL)
		}
	}
	if !strings.Contains(strings.ToUpper(downSQL), "DEFAULT ('NEW') FOR [STATUS]") {
		t.Fatalf("down SQL should restore old default: %s", downSQL)
	}
}

func TestSemanticFromMigrationAppliesCheckConstraintOnlyTransition(t *testing.T) {
	t.Parallel()

	current := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  amount DECIMAL(18,2) NULL
);
`
	sql := `
ALTER TABLE dbo.TableA ADD CONSTRAINT [CK_TableA__amount] CHECK ([amount] IS NULL OR [amount] >= 0);
`
	meta := migrationMeta{
		Category: "tables",
		Object:   "dbo.TableA",
	}
	next, err := semanticFromMigrationWithState(meta, sql, current, true)
	if err != nil {
		t.Fatalf("semantic apply failed: %v", err)
	}
	if !strings.Contains(strings.ToUpper(next), "CONSTRAINT [CK_TABLEA__AMOUNT] CHECK ([AMOUNT] IS NULL OR [AMOUNT] >= 0)") {
		t.Fatalf("expected CHECK constraint in reconstructed state: %s", next)
	}
}

func TestTableHashIgnoresTableConstraintOrder(t *testing.T) {
	t.Parallel()

	a := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  amount DECIMAL(18,2) NULL,
  status VARCHAR(20) NULL,
  CONSTRAINT CK_TableA__amount CHECK (amount >= 0),
  CONSTRAINT CK_TableA__status CHECK (status IS NULL OR status IN ('NEW','DONE'))
);
`
	b := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  amount DECIMAL(18,2) NULL,
  status VARCHAR(20) NULL,
  CONSTRAINT CK_TableA__status CHECK (status IS NULL OR status IN ('NEW','DONE')),
  CONSTRAINT CK_TableA__amount CHECK (amount >= 0)
);
`
	if hashSemanticByCategory("tables", a) != hashSemanticByCategory("tables", b) {
		t.Fatalf("table hash should ignore table-level constraint order")
	}
}

func TestApplyDependencyRestoreToHistoryStateKeepsIndexPresent(t *testing.T) {
	t.Parallel()

	tableSQL := `CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id),
  CONSTRAINT CK_TableA__status CHECK (status IN ('NEW','DONE'))
);`
	history := map[string]historyObject{
		objectKey("tables", "dbo.TableA"): {
			Category:    "tables",
			Object:      "dbo.TableA",
			SemanticSQL: tableSQL,
		},
		objectKey("indexes", "dbo.TableA"): {
			Category:    "indexes",
			Object:      "dbo.TableA",
			SemanticSQL: "CREATE INDEX IX_TableA__status ON dbo.TableA(status);",
		},
	}
	migrationSQL := `
IF (OBJECT_ID(N'dbo.CK_TableA__status', N'C') IS NOT NULL)
BEGIN
  ALTER TABLE dbo.TableA DROP CONSTRAINT [CK_TableA__status];
END
IF EXISTS (SELECT 1 FROM sys.indexes WHERE object_id = OBJECT_ID(N'dbo.TableA') AND name = N'IX_TableA__status')
BEGIN
  DROP INDEX IF EXISTS [IX_TableA__status] ON dbo.TableA;
END
ALTER TABLE dbo.TableA ALTER COLUMN [status] VARCHAR(20) NULL;
IF NOT (OBJECT_ID(N'dbo.CK_TableA__status', N'C') IS NOT NULL)
BEGIN
  ALTER TABLE dbo.TableA ADD CONSTRAINT CK_TableA__status CHECK (status IN ('NEW','DONE'));
END
IF NOT EXISTS (SELECT 1 FROM sys.indexes WHERE object_id = OBJECT_ID(N'dbo.TableA') AND name = N'IX_TableA__status')
BEGIN
  CREATE INDEX IX_TableA__status ON dbo.TableA(status);
END`

	drops := extractDroppedConstraintsFromMigrationSQL(migrationSQL)
	if err := applyConstraintDropsToHistoryState(history, drops); err != nil {
		t.Fatalf("apply constraint drops failed: %v", err)
	}
	adds := extractAddedConstraintsFromMigrationSQL(migrationSQL)
	if err := applyConstraintAddsToHistoryState(history, adds); err != nil {
		t.Fatalf("apply constraint adds failed: %v", err)
	}
	indexDrops := extractDropIndexTargets(migrationSQL)
	if err := applyIndexDropsToHistoryState(history, indexDrops); err != nil {
		t.Fatalf("apply index drops failed: %v", err)
	}
	indexCreates := extractCreateIndexStatementsAllowEmpty("dbo.TableA", migrationSQL)
	if err := applyIndexCreatesToHistoryState(history, "20260709000000000", "dbo.TableA", indexCreates); err != nil {
		t.Fatalf("apply index creates failed: %v", err)
	}

	tableState := history[objectKey("tables", "dbo.TableA")].SemanticSQL
	if !strings.Contains(strings.ToUpper(tableState), "CONSTRAINT CK_TABLEA__STATUS CHECK") {
		t.Fatalf("expected check constraint restored in table state: %s", tableState)
	}
	indexState := history[objectKey("indexes", "dbo.TableA")].SemanticSQL
	if !strings.Contains(strings.ToUpper(indexState), "CREATE INDEX IX_TABLEA__STATUS") {
		t.Fatalf("expected index restored in index state: %s", indexState)
	}
}

func TestDetectUnnamedConstraintDependencyBlocksLocal(t *testing.T) {
	t.Parallel()

	prev := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL,
  CHECK (status IN ('NEW','DONE'))
);
`
	history := map[string]historyObject{
		objectKey("tables", "dbo.TableA"): {
			Category:    "tables",
			Object:      "dbo.TableA",
			SemanticSQL: prev,
		},
	}

	blocked, err := detectUnnamedConstraintDependencyBlocks("dbo.TableA", []string{"status"}, prev, history)
	if err != nil {
		t.Fatalf("detectUnnamedConstraintDependencyBlocks returned error: %v", err)
	}
	if len(blocked) == 0 {
		t.Fatalf("expected unnamed dependency blockers")
	}
	joined := strings.ToUpper(strings.Join(blocked, ","))
	if !strings.Contains(joined, "CHECK:DBO.TABLEA.<UNNAMED:CHECK>") {
		t.Fatalf("expected unnamed check blocker, got: %v", blocked)
	}
}

func TestDetectUnnamedConstraintDependencyBlocksIncomingFK(t *testing.T) {
	t.Parallel()

	tableA := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id)
);
`
	tableB := `
CREATE TABLE dbo.TableB (
  id INT NOT NULL,
  status_ref VARCHAR(10) NULL,
  FOREIGN KEY (status_ref) REFERENCES dbo.TableA(status)
);
`
	history := map[string]historyObject{
		objectKey("tables", "dbo.TableA"): {
			Category:    "tables",
			Object:      "dbo.TableA",
			SemanticSQL: tableA,
		},
		objectKey("tables", "dbo.TableB"): {
			Category:    "tables",
			Object:      "dbo.TableB",
			SemanticSQL: tableB,
		},
	}

	blocked, err := detectUnnamedConstraintDependencyBlocks("dbo.TableA", []string{"status"}, tableA, history)
	if err != nil {
		t.Fatalf("detectUnnamedConstraintDependencyBlocks returned error: %v", err)
	}
	if len(blocked) == 0 {
		t.Fatalf("expected unnamed incoming fk blocker")
	}
	joined := strings.ToUpper(strings.Join(blocked, ","))
	if !strings.Contains(joined, "FOREIGN_KEY:DBO.TABLEB.<UNNAMED:FOREIGN_KEY>") {
		t.Fatalf("expected unnamed incoming fk blocker, got: %v", blocked)
	}
}
