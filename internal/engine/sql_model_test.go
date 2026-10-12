package engine

import (
	"strings"
	"testing"
)

func TestBuildTableModelAndDiff(t *testing.T) {
	before, err := BuildTableModel("dbo.Account", `CREATE TABLE dbo.Account (
  [id] INT NOT NULL,
  [name] NVARCHAR(100) NULL
);`)
	if err != nil {
		t.Fatal(err)
	}
	after, err := BuildTableModel("dbo.Account", `CREATE TABLE dbo.Account (
  [id] INT NOT NULL,
  [name] NVARCHAR(200) NULL,
  [created_at] DATETIME2(7) NULL
);`)
	if err != nil {
		t.Fatal(err)
	}
	diff := DiffTableModels(before, after)
	if len(diff.Added) != 1 || diff.Added[0].Name != "created_at" {
		t.Fatalf("unexpected additions: %#v", diff.Added)
	}
	if len(diff.Changed) != 1 || diff.Changed[0].After.Name != "name" {
		t.Fatalf("unexpected changes: %#v", diff.Changed)
	}
}

func TestRiskClassifiesGuardedLossyTableTransitions(t *testing.T) {
	before := `CREATE TABLE dbo.Amounts ([amount] DECIMAL(18,4) NOT NULL, [note] NVARCHAR(MAX) NULL);`
	after := `CREATE TABLE dbo.Amounts ([amount] DECIMAL(10,2) NOT NULL, [note] NVARCHAR(50) NULL);`
	risk, err := assessTableTransition(before, after, "dbo.Amounts")
	if err != nil {
		t.Fatal(err)
	}
	if risk.Level != riskGuarded || len(risk.Reasons) != 2 {
		t.Fatalf("expected two guarded reasons, got %#v", risk)
	}
}

func TestRiskClassifiesBracketedDecimalScaleChangeAsGuarded(t *testing.T) {
	before := `CREATE TABLE dbo.Amounts ([near] [DECIMAL](25,2) NULL, [nonnear] [DECIMAL] (25, 2) NULL);`
	after := `CREATE TABLE dbo.Amounts ([near] [DECIMAL](25,6) NULL, [nonnear] [DECIMAL] (25, 6) NULL);`
	risk, err := assessTableTransition(before, after, "dbo.Amounts")
	if err != nil {
		t.Fatal(err)
	}
	if risk.Level != riskGuarded || len(risk.Reasons) != 2 {
		t.Fatalf("bracketed DECIMAL scale changes must use guarded automatic migration, got %#v", risk)
	}

	model, err := BuildTableModel("dbo.Amounts", after)
	if err != nil {
		t.Fatal(err)
	}
	if model.Columns["near"].Type.Base != "DECIMAL" || model.Columns["nonnear"].Type.Base != "DECIMAL" {
		t.Fatalf("bracketed built-in type was not canonicalized: %#v", model.Columns)
	}
}

func TestBracketedAndBareBuiltInTypesAreSemanticallyEquivalent(t *testing.T) {
	before, err := BuildTableModel("dbo.T", `CREATE TABLE dbo.T ([amount] DECIMAL(25,6) NULL);`)
	if err != nil {
		t.Fatal(err)
	}
	after, err := BuildTableModel("dbo.T", `CREATE TABLE dbo.T ([amount] [decimal] (25, 6) NULL);`)
	if err != nil {
		t.Fatal(err)
	}
	if diff := DiffTableModels(before, after); len(diff.Changed) != 0 {
		t.Fatalf("type identifier delimiters must be style-only, got %#v", diff.Changed)
	}
	changes, err := classifyTableColumnChanges(
		`CREATE TABLE dbo.T ([amount] DECIMAL(25,6) NULL);`,
		`CREATE TABLE dbo.T ([amount] [decimal] (25, 6) NULL);`,
	)
	if err != nil {
		t.Fatal(err)
	}
	if len(changes.Changed) != 0 {
		t.Fatalf("type identifier delimiters must not emit ALTER COLUMN, got %#v", changes.Changed)
	}
}

func TestBracketedBuiltinTypeFamiliesCanonicalize(t *testing.T) {
	model, err := BuildTableModel("dbo.T", `CREATE TABLE dbo.T (
  [amount] [numeric](19,4) NULL,
  [code] [varchar](24) NULL,
  [payload] [varbinary](16) NULL,
  [occurred_at] [datetime2](5) NULL,
  [counter] [bigint] NULL
);`)
	if err != nil {
		t.Fatal(err)
	}
	for column, want := range map[string]string{
		"amount": "NUMERIC", "code": "VARCHAR", "payload": "VARBINARY",
		"occurred_at": "DATETIME2", "counter": "BIGINT",
	} {
		if got := model.Columns[column].Type.Base; got != want {
			t.Errorf("%s base type = %q, want %q", column, got, want)
		}
	}
}

func TestRiskGuardsDropColumnWhenDataMustBeEmpty(t *testing.T) {
	risk, err := assessTableTransition(`CREATE TABLE dbo.T ([id] INT NOT NULL, [obsolete] INT NULL);`, `CREATE TABLE dbo.T ([id] INT NOT NULL);`, "dbo.T")
	if err != nil {
		t.Fatal(err)
	}
	if risk.Level != riskGuarded {
		t.Fatalf("drop column should be guarded by an empty-data precheck, got %#v", risk)
	}
}

func TestRiskGuardsStringToExactNumericConversion(t *testing.T) {
	risk, err := assessTableTransition(`CREATE TABLE dbo.T ([amount] VARCHAR(20) COLLATE Latin1_General_BIN NULL);`, `CREATE TABLE dbo.T ([amount] DECIMAL(10,2) NULL);`, "dbo.T")
	if err != nil {
		t.Fatal(err)
	}
	if risk.Level != riskGuarded {
		t.Fatalf("string to exact numeric should use a round-trip guard, got %#v", risk)
	}
	previous := tableColumnDef{Name: "amount", BaseType: "VARCHAR", RawType: "VARCHAR(20) COLLATE Latin1_General_BIN", Length: 20, Nullable: true}
	next := tableColumnDef{Name: "amount", BaseType: "DECIMAL", RawType: "DECIMAL(10,2)", Nullable: true}
	sql := strings.ToUpper(buildAlterColumnGuardedSQL("dbo.T", previous, next))
	for _, want := range []string{"TRY_CONVERT(DECIMAL(10,2), [AMOUNT]) IS NULL", "TRY_CONVERT(VARCHAR(20), TRY_CONVERT(DECIMAL(10,2), [AMOUNT])) <> [AMOUNT]", "THROW 51000"} {
		if !strings.Contains(sql, want) {
			t.Fatalf("string/numeric guard missing %q:\n%s", want, sql)
		}
	}
}

func TestRiskKeepsFloatTransitionManual(t *testing.T) {
	risk, err := assessTableTransition(`CREATE TABLE dbo.T ([value] FLOAT NULL);`, `CREATE TABLE dbo.T ([value] DECIMAL(18,4) NULL);`, "dbo.T")
	if err != nil {
		t.Fatal(err)
	}
	if risk.Level != riskManual {
		t.Fatalf("float transition must remain manual, got %#v", risk)
	}
}

func TestRiskClassifiesXMLToNVarCharMaxAsGuarded(t *testing.T) {
	risk, err := assessTableTransition(
		`CREATE TABLE dbo.T ([payload] XML NULL);`,
		`CREATE TABLE dbo.T ([payload] NVARCHAR(MAX) COLLATE Latin1_General_BIN NULL);`,
		"dbo.T",
	)
	if err != nil {
		t.Fatal(err)
	}
	if risk.Level != riskGuarded {
		t.Fatalf("XML to NVARCHAR(MAX) should be guarded automatic, got %#v", risk)
	}
}

func TestRiskKeepsTypedXMLAndCollationChangesManual(t *testing.T) {
	tests := []struct {
		name   string
		before string
		after  string
	}{
		{
			name:   "typed XML requires schema-aware validation",
			before: `CREATE TABLE dbo.T ([payload] XML(dbo.PayloadSchema) NULL);`,
			after:  `CREATE TABLE dbo.T ([payload] NVARCHAR(MAX) NULL);`,
		},
		{
			name:   "varchar collation change may change code page",
			before: `CREATE TABLE dbo.T ([code] VARCHAR(50) COLLATE Latin1_General_BIN NULL);`,
			after:  `CREATE TABLE dbo.T ([code] VARCHAR(50) COLLATE SQL_Latin1_General_CP1_CI_AS NULL);`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			risk, err := assessTableTransition(tt.before, tt.after, "dbo.T")
			if err != nil {
				t.Fatal(err)
			}
			if risk.Level != riskManual {
				t.Fatalf("transition must remain manual, got %#v", risk)
			}
		})
	}
}

func TestRiskAllowsNonPersistedComputedExpressionRebuild(t *testing.T) {
	risk, err := assessTableTransition(`CREATE TABLE dbo.T ([base] INT NOT NULL, [value] AS ([base] * 2));`, `CREATE TABLE dbo.T ([base] INT NOT NULL, [value] AS ([base] * 3));`, "dbo.T")
	if err != nil {
		t.Fatal(err)
	}
	if risk.Level != riskGuarded {
		t.Fatalf("non-persisted computed change should be guarded, got %#v", risk)
	}
}

func TestTableModelPreservesIdentityAndRowGuid(t *testing.T) {
	model, err := BuildTableModel("dbo.T", `CREATE TABLE dbo.T ([Id] BIGINT IDENTITY(100, 5) NOT NULL, [Guid] UNIQUEIDENTIFIER ROWGUIDCOL NULL);`)
	if err != nil {
		t.Fatal(err)
	}
	if model.Columns["Id"].Identity != "IDENTITY(100, 5)" || !model.Columns["Guid"].RowGuidCol {
		t.Fatalf("modifiers not preserved: %#v", model.Columns)
	}
}

func TestCapabilityMatrixRejectsUnsupportedTableAndIndexFeatures(t *testing.T) {
	if issue := assessSourceCapability("tables", `CREATE TABLE dbo.T ([a] INT SPARSE NULL);`); issue != nil {
		t.Fatalf("SPARSE is supported, got: %v", issue)
	}
	if issue := assessSourceCapability("indexes", `CREATE NONCLUSTERED COLUMNSTORE INDEX IX ON dbo.T ([a]);`); issue != nil {
		t.Fatalf("columnstore is supported, got: %v", issue)
	}
	if issue := assessSourceCapability("indexes", `CREATE INDEX IX ON dbo.T ([a]) WITH (RESUMABLE = ON);`); issue == nil {
		t.Fatal("expected resumable rejection until non-transactional runner is active")
	}
	if issue := assessSourceCapability("tables", `CREATE TABLE dbo.T ([a] INT NOT NULL);`); issue != nil {
		t.Fatalf("unexpected issue: %v", issue)
	}
	if issue := assessSourceCapability("tables", `CREATE TABLE dbo.T ([a] NVARCHAR(MAX) NULL) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY];`); issue != nil {
		t.Fatalf("TEXTIMAGE_ON is valid pass-through create metadata: %v", issue)
	}
	if err := validateTableSubsetSupport(`CREATE TABLE dbo.T ([a] NVARCHAR(MAX) NULL) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY];`); err != nil {
		t.Fatalf("TEXTIMAGE_ON create should reconstruct losslessly: %v", err)
	}
}

func TestCanonicalizeIndexSnapshotDoesNotDuplicateSemicolonlessStatements(t *testing.T) {
	input := "CREATE CLUSTERED INDEX IX_A ON dbo.T ([A])\nCREATE NONCLUSTERED INDEX IX_B ON dbo.T ([B])\n"
	got := canonicalizeIndexSnapshotFormatting(input, "dbo.T")
	if strings.Count(got, "CREATE CLUSTERED INDEX IX_A") != 1 || strings.Count(got, "CREATE NONCLUSTERED INDEX IX_B") != 1 {
		t.Fatalf("index statements duplicated or lost:\n%s", got)
	}
}

func TestTableModelPreservesSparseAndRejectsNotNullSparse(t *testing.T) {
	model, err := BuildTableModel("dbo.T", `CREATE TABLE dbo.T ([Id] INT NOT NULL, [OptionalCode] VARCHAR(20) SPARSE NULL);`)
	if err != nil {
		t.Fatal(err)
	}
	if !model.Columns["OptionalCode"].Sparse {
		t.Fatalf("SPARSE modifier not preserved: %#v", model.Columns)
	}
	if _, err := BuildTableModel("dbo.T", `CREATE TABLE dbo.T ([Bad] INT SPARSE NOT NULL);`); err == nil {
		t.Fatal("SPARSE NOT NULL must be rejected before SQL generation")
	}
}

func TestTableModelKeepsCaseDistinctColumns(t *testing.T) {
	model, err := BuildTableModel("dbo.T", `CREATE TABLE dbo.T ([Code] INT NOT NULL, [code] INT NULL);`)
	if err != nil {
		t.Fatal(err)
	}
	if len(model.Columns) != 2 || model.Columns["Code"].Name != "Code" || model.Columns["code"].Name != "code" {
		t.Fatalf("case-distinct columns were collapsed: %#v", model.Columns)
	}
}

func TestTableModelRejectsCollationOnKnownNonCharacterType(t *testing.T) {
	_, err := BuildTableModel("dbo.T", `CREATE TABLE dbo.T ([Payload] XML COLLATE Latin1_General_BIN NULL);`)
	if err == nil || !strings.Contains(err.Error(), "COLLATE with non-character type XML") {
		t.Fatalf("expected actionable invalid COLLATE error, got %v", err)
	}
}
