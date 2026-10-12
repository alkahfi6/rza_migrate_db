package engine

import (
	"strings"
	"testing"
)

func TestColumnGenerationSafetyMatrixDoesNotSkipDDL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		sql   string
		wants []string
	}{
		{
			name: "add nullable column",
			sql: buildAddColumnIfMissingStatement("dbo.Customer", tableColumnDef{
				Name: "Note", BaseType: "NVARCHAR", RawType: "NVARCHAR(200)", Length: 200, IsUnicode: true, Nullable: true,
			}),
			wants: []string{"C.NAME COLLATE LATIN1_GENERAL_100_BIN2 = N'NOTE'", "ALTER TABLE DBO.CUSTOMER ADD [NOTE] NVARCHAR(200) NULL;", "SCHEMA DRIFT"},
		},
		{
			name: "add not null column with default",
			sql: buildAddNotNullColumnWithDefaultSQL(
				"dbo.Customer",
				tableColumnDef{Name: "Status", BaseType: "VARCHAR", RawType: "VARCHAR(20)", Length: 20, Nullable: false},
				tableDefaultConstraintDef{ConstraintName: "DF_Customer_Status", Column: "Status", Expression: "('NEW')"},
			),
			wants: []string{"C.NAME COLLATE LATIN1_GENERAL_100_BIN2 = N'STATUS'", "ALTER TABLE DBO.CUSTOMER ADD [STATUS] VARCHAR(20) NOT NULL", "WITH VALUES;", "SCHEMA DRIFT"},
		},
		{
			name: "alter widening column",
			sql: buildAlterColumnGuardedSQL(
				"dbo.Customer",
				tableColumnDef{Name: "Name", BaseType: "NVARCHAR", RawType: "NVARCHAR(100)", Length: 100, IsUnicode: true, Nullable: true},
				tableColumnDef{Name: "Name", BaseType: "NVARCHAR", RawType: "NVARCHAR(200)", Length: 200, IsUnicode: true, Nullable: true},
			),
			wants: []string{"C.NAME COLLATE LATIN1_GENERAL_100_BIN2 = N'NAME'", "THROW 51000", "ALTER TABLE DBO.CUSTOMER ALTER COLUMN [NAME] NVARCHAR(200) NULL;"},
		},
		{
			name:  "drop column",
			sql:   buildDropColumnIfExistsStatement("dbo.Customer", "LegacyCode"),
			wants: []string{"C.NAME COLLATE LATIN1_GENERAL_100_BIN2 = N'LEGACYCODE'", "WHERE [LEGACYCODE] IS NOT NULL", "PRECHECK BLOCKED: DROP COLUMN", "ALTER TABLE DBO.CUSTOMER DROP COLUMN [LEGACYCODE];"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			upper := strings.ToUpper(strings.TrimSpace(tt.sql))
			if upper == "" {
				t.Fatal("generator silently returned an empty SQL body")
			}
			for _, want := range tt.wants {
				if !strings.Contains(upper, want) {
					t.Fatalf("generated SQL missing %q:\n%s", want, tt.sql)
				}
			}
		})
	}
}

func TestDestructiveDownRequiresEmptyData(t *testing.T) {
	t.Parallel()

	sql := strings.ToUpper(buildDropTableSQL("dbo.TransientData"))
	for _, want := range []string{
		"OBJECT_ID(N'DBO.TRANSIENTDATA', N'U') IS NOT NULL",
		"EXISTS (SELECT 1 FROM DBO.TRANSIENTDATA WITH (TABLOCKX, HOLDLOCK))",
		"PRECHECK BLOCKED: DROP TABLE TARGET CONTAINS DATA",
		"DROP TABLE IF EXISTS DBO.TRANSIENTDATA;",
	} {
		if !strings.Contains(sql, want) {
			t.Fatalf("destructive down missing fail-closed guard %q:\n%s", want, sql)
		}
	}
}

func TestConstraintGenerationSafetyMatrixDoesNotSkipDDL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		kind       string
		constraint string
		definition string
		wants      []string
	}{
		{
			name: "default", kind: "default", constraint: "DF_Customer_Status",
			definition: "CONSTRAINT [DF_Customer_Status] DEFAULT ('NEW') FOR [Status]",
			wants:      []string{"OBJECT_ID(N'DBO.DF_CUSTOMER_STATUS', N'D') IS NOT NULL", "ALTER TABLE DBO.CUSTOMER ADD CONSTRAINT [DF_CUSTOMER_STATUS] DEFAULT ('NEW') FOR [STATUS];"},
		},
		{
			name: "check", kind: "check", constraint: "CK_Customer_Amount",
			definition: "CONSTRAINT [CK_Customer_Amount] CHECK ([Amount] IS NULL OR [Amount] \u003e= 0)",
			wants:      []string{"OBJECT_ID(N'DBO.CK_CUSTOMER_AMOUNT', N'C') IS NOT NULL", "WHERE NOT ([AMOUNT] IS NULL OR [AMOUNT] \u003e= 0)", "THROW 51000", "ALTER TABLE DBO.CUSTOMER ADD CONSTRAINT [CK_CUSTOMER_AMOUNT]"},
		},
		{
			name: "unique", kind: "unique", constraint: "UQ_Customer_Email",
			definition: "CONSTRAINT [UQ_Customer_Email] UNIQUE ([Email])",
			wants:      []string{"OBJECT_ID(N'DBO.UQ_CUSTOMER_EMAIL', N'UQ') IS NOT NULL", "GROUP BY [EMAIL]", "HAVING COUNT_BIG(*) \u003e 1", "ALTER TABLE DBO.CUSTOMER ADD CONSTRAINT [UQ_CUSTOMER_EMAIL]"},
		},
		{
			name: "primary key", kind: "primary_key", constraint: "PK_Customer",
			definition: "CONSTRAINT [PK_Customer] PRIMARY KEY ([Id])",
			wants:      []string{"OBJECT_ID(N'DBO.PK_CUSTOMER', N'PK') IS NOT NULL", "[ID] IS NULL", "HAVING COUNT_BIG(*) \u003e 1", "ALTER TABLE DBO.CUSTOMER ADD CONSTRAINT [PK_CUSTOMER]"},
		},
		{
			name: "foreign key", kind: "foreign_key", constraint: "FK_Customer_Tenant",
			definition: "CONSTRAINT [FK_Customer_Tenant] FOREIGN KEY ([TenantId]) REFERENCES dbo.Tenant ([Id])",
			wants:      []string{"OBJECT_ID(N'DBO.FK_CUSTOMER_TENANT', N'F') IS NOT NULL", "C.[TENANTID] IS NOT NULL", "NOT EXISTS (", "ORPHAN ROWS EXIST", "ALTER TABLE DBO.CUSTOMER ADD CONSTRAINT [FK_CUSTOMER_TENANT]"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql, err := buildAddConstraintWithDataPrecheckStatement("dbo.Customer", tt.constraint, tt.kind, tt.definition)
			if err != nil {
				t.Fatalf("guarded constraint generation failed: %v", err)
			}
			upper := strings.ToUpper(strings.TrimSpace(sql))
			if upper == "" {
				t.Fatal("generator silently returned an empty SQL body")
			}
			for _, want := range tt.wants {
				if !strings.Contains(upper, want) {
					t.Fatalf("generated SQL missing %q:\n%s", want, sql)
				}
			}
		})
	}
}

func TestConstraintLifecycleMatrixEmitsBothDirections(t *testing.T) {
	t.Parallel()

	base := `CREATE TABLE dbo.Customer ([Id] INT NOT NULL, [TenantId] INT NULL, [Amount] INT NULL, [Status] VARCHAR(20) NULL);`
	tests := []struct {
		name string
		next string
	}{
		{name: "default", next: `CREATE TABLE dbo.Customer ([Id] INT NOT NULL, [TenantId] INT NULL, [Amount] INT NULL, [Status] VARCHAR(20) NULL, CONSTRAINT [DF_Customer_Status] DEFAULT ('NEW') FOR [Status]);`},
		{name: "check", next: `CREATE TABLE dbo.Customer ([Id] INT NOT NULL, [TenantId] INT NULL, [Amount] INT NULL, [Status] VARCHAR(20) NULL, CONSTRAINT [CK_Customer_Amount] CHECK ([Amount] IS NULL OR [Amount] \u003e= 0));`},
		{name: "unique", next: `CREATE TABLE dbo.Customer ([Id] INT NOT NULL, [TenantId] INT NULL, [Amount] INT NULL, [Status] VARCHAR(20) NULL, CONSTRAINT [UQ_Customer_Status] UNIQUE ([Status]));`},
		{name: "primary key", next: `CREATE TABLE dbo.Customer ([Id] INT NOT NULL, [TenantId] INT NULL, [Amount] INT NULL, [Status] VARCHAR(20) NULL, CONSTRAINT [PK_Customer] PRIMARY KEY ([Id]));`},
		{name: "foreign key", next: `CREATE TABLE dbo.Customer ([Id] INT NOT NULL, [TenantId] INT NULL, [Amount] INT NULL, [Status] VARCHAR(20) NULL, CONSTRAINT [FK_Customer_Tenant] FOREIGN KEY ([TenantId]) REFERENCES dbo.Tenant ([Id]));`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			up, down, ok, err := buildCheckConstraintOnlyLifecycleSQL(base, tt.next, "dbo.Customer")
			if err != nil {
				t.Fatalf("constraint lifecycle failed: %v", err)
			}
			if !ok || strings.TrimSpace(up) == "" || strings.TrimSpace(down) == "" {
				t.Fatalf("constraint lifecycle was silently skipped: ok=%v\nUP:\n%s\nDOWN:\n%s", ok, up, down)
			}
			if !strings.Contains(strings.ToUpper(up), "ALTER TABLE") || !strings.Contains(strings.ToUpper(down), "ALTER TABLE") {
				t.Fatalf("both directions must contain guarded DDL:\nUP:\n%s\nDOWN:\n%s", up, down)
			}
		})
	}
}

func TestGeneratedDDLRequiresTargetTablesAndColumns(t *testing.T) {
	t.Parallel()

	columnSQL := strings.ToUpper(buildAddColumnIfMissingStatement("audit.Customer", tableColumnDef{Name: "Code", BaseType: "INT", RawType: "INT", Nullable: true}))
	for _, want := range []string{"OBJECT_ID(N'AUDIT.CUSTOMER', N'U') IS NULL", "PRECHECK BLOCKED: ADD COLUMN", "C.NAME COLLATE LATIN1_GENERAL_100_BIN2 = N'CODE'"} {
		if !strings.Contains(columnSQL, want) {
			t.Fatalf("ADD COLUMN missing fail-fast precheck %q:\n%s", want, columnSQL)
		}
	}

	indexSQL := strings.ToUpper(buildCreateIndexIfNotExistsStatement("audit.Customer", "IX_Customer_Code", "CREATE NONCLUSTERED INDEX [IX_Customer_Code] ON audit.Customer ([Code]) INCLUDE ([DisplayName], [ExactCase]);"))
	for _, want := range []string{"OBJECT_ID(N'AUDIT.CUSTOMER', N'U') IS NULL", "PRECHECK BLOCKED: CREATE INDEX", "C.NAME COLLATE LATIN1_GENERAL_100_BIN2 = N'CODE'", "N'DISPLAYNAME'", "N'EXACTCASE'"} {
		if !strings.Contains(indexSQL, want) {
			t.Fatalf("CREATE INDEX missing fail-fast precheck %q:\n%s", want, indexSQL)
		}
	}

	pkSQL, err := buildAddKeyConstraintWithPrecheckStatement("audit.Customer", "PK_Customer", "primary_key", "CONSTRAINT [PK_Customer] PRIMARY KEY ([Id])")
	if err != nil {
		t.Fatal(err)
	}
	pkUpper := strings.ToUpper(pkSQL)
	for _, want := range []string{"OBJECT_ID(N'AUDIT.CUSTOMER', N'U') IS NULL", "C.NAME COLLATE LATIN1_GENERAL_100_BIN2 = N'ID'", "OBJECT_ID(N'AUDIT.PK_CUSTOMER', N'PK') IS NOT NULL"} {
		if !strings.Contains(pkUpper, want) {
			t.Fatalf("PRIMARY KEY missing schema-aware precheck %q:\n%s", want, pkSQL)
		}
	}
}

func TestCreateIndexIncludeParserRejectsCommaAndClauseErrors(t *testing.T) {
	t.Parallel()

	valid := `CREATE INDEX [IX_Complex] ON dbo.Customer ([TenantId] DESC, [Code] ASC) INCLUDE ([DisplayName], [ExactCase]) WHERE [IsActive] = 1 WITH (FILLFACTOR = 90);`
	columns, err := parseCreateIndexIncludeColumns(valid)
	if err != nil {
		t.Fatalf("valid INCLUDE rejected: %v", err)
	}
	if got, want := strings.Join(columns, "|"), "DisplayName|ExactCase"; got != want {
		t.Fatalf("INCLUDE columns = %q, want %q", got, want)
	}

	invalid := []string{
		`CREATE INDEX IX_A ON dbo.Customer ([Id]) INCLUDE ();`,
		`CREATE INDEX IX_A ON dbo.Customer ([Id]) INCLUDE ([Code],);`,
		`CREATE INDEX IX_A ON dbo.Customer ([Id]) INCLUDE [Code];`,
		`CREATE INDEX IX_A ON dbo.Customer ([Id]) INCLUDE ([Code] DESC);`,
		`CREATE INDEX IX_A ON dbo.Customer ([Id]) INCLUDE ([Code]) INCLUDE ([Name]);`,
	}
	for _, sqlText := range invalid {
		if _, err := parseIndexStateMap("dbo.Customer", sqlText); err == nil {
			t.Fatalf("malformed INCLUDE was accepted: %s", sqlText)
		}
	}
}

func TestComplexFilteredIncludeIndexLifecycleIsReversibleAndCaseExact(t *testing.T) {
	t.Parallel()

	previous := `CREATE UNIQUE NONCLUSTERED INDEX [IX_Order_Complex] ON dbo.OrderLine ([TenantID] DESC, [OrderCode] ASC) INCLUDE ([DisplayName]) WHERE [IsDeleted] = 0 WITH (FILLFACTOR = 85, PAD_INDEX = ON);`
	next := `CREATE UNIQUE NONCLUSTERED INDEX [IX_Order_Complex] ON dbo.OrderLine ([TenantID] DESC, [OrderCode] ASC) INCLUDE ([DisplayName], [ExactCasePayload]) WHERE [IsDeleted] = 0 WITH (FILLFACTOR = 90, PAD_INDEX = ON, ALLOW_PAGE_LOCKS = OFF);`

	up, down, err := buildIndexLifecycleSQL(previous, next, "dbo.OrderLine")
	if err != nil {
		t.Fatalf("complex lifecycle failed: %v", err)
	}
	for direction, sqlText := range map[string]string{"up": up, "down": down} {
		for _, token := range []string{
			"DROP INDEX IF EXISTS [IX_Order_Complex] ON dbo.OrderLine",
			"CREATE UNIQUE NONCLUSTERED INDEX [IX_Order_Complex]",
			"name COLLATE Latin1_General_100_BIN2 = N'TenantID' COLLATE Latin1_General_100_BIN2",
			"name COLLATE Latin1_General_100_BIN2 = N'OrderCode' COLLATE Latin1_General_100_BIN2",
			"name COLLATE Latin1_General_100_BIN2 = N'DisplayName' COLLATE Latin1_General_100_BIN2",
			"GROUP BY [TenantID], [OrderCode]",
			"WHERE [IsDeleted] = 0",
		} {
			if !strings.Contains(sqlText, token) {
				t.Fatalf("%s SQL missing %q:\n%s", direction, token, sqlText)
			}
		}
	}
	if !strings.Contains(up, "N'ExactCasePayload'") || !strings.Contains(up, "FILLFACTOR = 90") {
		t.Fatalf("up SQL lost next-state INCLUDE/options:\n%s", up)
	}
	if strings.Contains(down, "N'ExactCasePayload'") || !strings.Contains(down, "FILLFACTOR = 85") {
		t.Fatalf("down SQL did not restore exact previous state:\n%s", down)
	}
}

func TestUniqueIndexGenerationChecksDuplicateData(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		indexName string
		statement string
		wants     []string
	}{
		{
			name:      "composite",
			indexName: "UX_Order",
			statement: "CREATE UNIQUE NONCLUSTERED INDEX [UX_Order] ON dbo.Orders ([TenantId] ASC, [OrderNo] DESC);",
			wants:     []string{"GROUP BY [TENANTID], [ORDERNO]", "HAVING COUNT_BIG(*) > 1", "DUPLICATE KEY VALUES"},
		},
		{
			name:      "filtered",
			indexName: "UX_Order_Active",
			statement: "CREATE UNIQUE NONCLUSTERED INDEX [UX_Order_Active] ON dbo.Orders ([TenantId], [OrderNo]) WHERE [IsActive] = 1 WITH (FILLFACTOR = 90);",
			wants:     []string{"WHERE [ISACTIVE] = 1", "GROUP BY [TENANTID], [ORDERNO]", "HAVING COUNT_BIG(*) > 1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql := strings.ToUpper(buildCreateIndexIfNotExistsStatement("dbo.Orders", tt.indexName, tt.statement))
			if strings.TrimSpace(sql) == "" {
				t.Fatal("unique index generator returned empty SQL")
			}
			for _, want := range tt.wants {
				if !strings.Contains(sql, want) {
					t.Fatalf("UNIQUE INDEX missing precheck %q:\n%s", want, sql)
				}
			}
		})
	}
}

func TestGeneratedIndexPrechecksPreserveIdentifierCase(t *testing.T) {
	t.Parallel()

	sql := buildCreateIndexIfNotExistsStatement(
		"dbo.TRSDebugFlag_TR",
		"TRSDebugFlag_TR_CU1",
		"CREATE UNIQUE CLUSTERED INDEX [TRSDebugFlag_TR_CU1] ON dbo.TRSDebugFlag_TR ([SPName] ASC);",
	)
	for _, want := range []string{
		"c.name COLLATE Latin1_General_100_BIN2 = N'SPName'",
		"GROUP BY [SPName]",
		"ON dbo.TRSDebugFlag_TR ([SPName] ASC)",
	} {
		if !strings.Contains(sql, want) {
			t.Fatalf("case-sensitive index SQL lost source identifier %q:\n%s", want, sql)
		}
	}
	for _, unwanted := range []string{"N'spname'", "[spname]"} {
		if strings.Contains(sql, unwanted) {
			t.Fatalf("case-sensitive index SQL contains lowercased identifier %q:\n%s", unwanted, sql)
		}
	}
}

func TestGeneratedConstraintPrechecksPreserveIdentifierCase(t *testing.T) {
	t.Parallel()

	pkSQL, err := buildAddKeyConstraintWithPrecheckStatement(
		"dbo.TRSParameterTraderMX_TH",
		"PK_TRSParameterTraderMX_TH",
		"primary_key",
		"CONSTRAINT [PK_TRSParameterTraderMX_TH] PRIMARY KEY ([HistId], [Code])",
	)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"N'HistId'",
		"N'Code'",
		"GROUP BY [HistId], [Code]",
	} {
		if !strings.Contains(pkSQL, want) {
			t.Fatalf("case-sensitive constraint SQL lost source identifier %q:\n%s", want, pkSQL)
		}
	}

	fkSQL, err := buildAddForeignKeyWithPrecheckStatement(
		"audit.ChildRow",
		"FK_ChildRow_ParentRow",
		"CONSTRAINT [FK_ChildRow_ParentRow] FOREIGN KEY ([ParentId]) REFERENCES [core].[ParentRow] ([ParentId])",
	)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"N'ParentId'",
		"p.[ParentId] = c.[ParentId]",
	} {
		if !strings.Contains(fkSQL, want) {
			t.Fatalf("case-sensitive foreign-key SQL lost source identifier %q:\n%s", want, fkSQL)
		}
	}
}

func TestDDLQuotesUnsafeAndReservedObjectNames(t *testing.T) {
	t.Parallel()

	unsafe := buildAddColumnIfMissingStatement("dbo.Order Detail", tableColumnDef{Name: "select", BaseType: "INT", RawType: "INT", Nullable: true})
	if !strings.Contains(unsafe, "ALTER TABLE dbo.[Order Detail] ADD [select] INT NULL;") {
		t.Fatalf("unsafe object/column identifier was not quoted safely:\n%s", unsafe)
	}
	reserved := buildDropTableSQL("dbo.Order")
	if !strings.Contains(reserved, "DROP TABLE IF EXISTS dbo.[Order];") {
		t.Fatalf("reserved object identifier was not quoted safely:\n%s", reserved)
	}
}

func TestDuplicateIndexNamesAreRejected(t *testing.T) {
	t.Parallel()

	_, err := parseIndexStateMap("dbo.T", `
CREATE INDEX [IX_T_Code] ON dbo.T ([Code]);
CREATE INDEX [IX_T_Code] ON dbo.T ([OtherCode]);`)
	if err == nil || !strings.Contains(strings.ToLower(err.Error()), "duplicate index name") {
		t.Fatalf("expected duplicate index name rejection, got %v", err)
	}
}

func TestAdvancedIndexDependencyBlocksAutomaticAlterColumn(t *testing.T) {
	t.Parallel()

	plan := dropColumnDependencyPlan{Indexes: []dropIndexTarget{{Table: "dbo.Documents", Name: "PXML_Documents_Payload"}}}
	history := map[string]historyObject{
		objectKey("indexes", "dbo.Documents"): {
			Category: "indexes",
			Object:   "dbo.Documents",
			SemanticSQL: `CREATE PRIMARY XML INDEX [PXML_Documents_Payload]
ON dbo.Documents ([Payload]);`,
		},
	}
	blocked, err := findUnsafeAdvancedAlterIndexDependencies(plan, history)
	if err != nil {
		t.Fatal(err)
	}
	if len(blocked) != 1 || blocked[0] != "dbo.Documents.PXML_Documents_Payload" {
		t.Fatalf("advanced index dependency must block automatic ALTER COLUMN: %#v", blocked)
	}
}
