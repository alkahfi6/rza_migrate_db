package engine

import (
	"strings"
	"testing"
)

func TestGuardedAlterColumnThrowsInsteadOfSilentlySkipping(t *testing.T) {
	previous := tableColumnDef{Name: "amount", BaseType: "DECIMAL", RawType: "DECIMAL(18,4)", Nullable: true}
	next := tableColumnDef{Name: "amount", BaseType: "DECIMAL", RawType: "DECIMAL(10,2)", Nullable: false}
	sql := buildAlterColumnGuardedSQL("dbo.Amount", previous, next)
	for _, want := range []string{"TRY_CONVERT(DECIMAL(18,4), TRY_CONVERT(DECIMAL(10,2), [amount])) <> [amount]", "THROW 51000", "WHERE [amount] IS NULL"} {
		if !strings.Contains(sql, want) {
			t.Fatalf("generated guard missing %q:\n%s", want, sql)
		}
	}
}

func TestGuardedAlterColumnAlwaysChecksTargetColumnExists(t *testing.T) {
	previous := tableColumnDef{Name: "payload", BaseType: "NVARCHAR", RawType: "NVARCHAR(100)", Length: 100, IsUnicode: true, Nullable: true}
	next := tableColumnDef{Name: "payload", BaseType: "NVARCHAR", RawType: "NVARCHAR(MAX)", Length: -1, IsUnicode: true, Nullable: true}

	sql := buildAlterColumnGuardedSQL("dbo.Events", previous, next)
	for _, want := range []string{
		"c.name COLLATE Latin1_General_100_BIN2 = N'payload' COLLATE Latin1_General_100_BIN2",
		"@reason = N'COLUMN_NOT_EXISTS'",
		"THROW 51000",
		"ALTER TABLE dbo.Events ALTER COLUMN [payload] NVARCHAR(MAX) NULL;",
	} {
		if !strings.Contains(sql, want) {
			t.Fatalf("generated baseline guard missing %q:\n%s", want, sql)
		}
	}
}

func TestGuardedXMLNVarCharMaxTransitionChecksRollbackData(t *testing.T) {
	xml := tableColumnDef{Name: "payload", BaseType: "XML", RawType: "XML", Length: 0, Nullable: true}
	text := tableColumnDef{Name: "payload", BaseType: "NVARCHAR", RawType: "NVARCHAR(MAX) COLLATE Latin1_General_BIN", Length: -1, IsUnicode: true, Nullable: true}

	up := strings.ToUpper(buildAlterColumnGuardedSQL("dbo.Events", xml, text))
	for _, want := range []string{
		"C.NAME COLLATE LATIN1_GENERAL_100_BIN2 = N'PAYLOAD' COLLATE LATIN1_GENERAL_100_BIN2",
		"ALTER TABLE DBO.EVENTS ALTER COLUMN [PAYLOAD] NVARCHAR(MAX) COLLATE LATIN1_GENERAL_BIN NULL;",
	} {
		if !strings.Contains(up, want) {
			t.Fatalf("XML to NVARCHAR(MAX) up migration missing %q:\n%s", want, up)
		}
	}

	down := strings.ToUpper(buildAlterColumnGuardedSQL("dbo.Events", text, xml))
	for _, want := range []string{
		"TRY_CONVERT(XML, [PAYLOAD]) IS NULL",
		"INVALID_XML_CNT=",
		"THROW 51000",
		"ALTER TABLE DBO.EVENTS ALTER COLUMN [PAYLOAD] XML NULL;",
	} {
		if !strings.Contains(down, want) {
			t.Fatalf("NVARCHAR(MAX) to XML down migration missing %q:\n%s", want, down)
		}
	}
}

func TestComputedExpressionLifecycleRebuildsBothDirections(t *testing.T) {
	before := `CREATE TABLE dbo.T ([base] INT NOT NULL, [value] AS ([base] * 2));`
	after := `CREATE TABLE dbo.T ([base] INT NOT NULL, [value] AS ([base] * 3));`
	changes, err := classifyTableColumnChanges(before, after)
	if err != nil {
		t.Fatal(err)
	}
	up, down, _, err := buildTableAlterColumnsSQL(before, after, "dbo.T", changes)
	if err != nil {
		t.Fatal(err)
	}
	for _, sql := range []string{up, down} {
		if !strings.Contains(strings.ToUpper(sql), "DROP COLUMN") || !strings.Contains(strings.ToUpper(sql), "ADD [VALUE] AS") {
			t.Fatalf("computed lifecycle not generated:\n%s", sql)
		}
	}
}
