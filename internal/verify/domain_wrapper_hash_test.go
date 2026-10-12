package verify

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"testing"

	"rza_migrate_db/internal/output"
)

func TestValidateWrapperContentHashAcceptsLineEndingChangesAndRejectsTampering(t *testing.T) {
	t.Parallel()

	body := "IF OBJECT_ID(N'dbo.SafeTable', N'U') IS NULL\nBEGIN\n  CREATE TABLE dbo.SafeTable (id INT NOT NULL);\nEND"
	sum := sha256.Sum256([]byte(body))
	sqlText := fmt.Sprintf(
		"DECLARE @content_sha256 CHAR(64) = N'%x';\n-- MIGRATION_BODY_BEGIN\n%s\n-- MIGRATION_BODY_END\n",
		sum,
		body,
	)

	report := output.NewReport("pattern", "LF")
	validateWrapperContentHash(&report, "migrations/example.up.sql", strings.ReplaceAll(sqlText, "\n", "\r\n"))
	if report.Fail.Count != 0 {
		t.Fatalf("valid body rejected after line-ending conversion: %#v", report.Fail.Items)
	}

	tampered := strings.Replace(sqlText, "id INT", "id BIGINT", 1)
	report = output.NewReport("pattern", "LF")
	validateWrapperContentHash(&report, "migrations/example.up.sql", tampered)
	if report.Fail.Count != 1 || report.Fail.Items[0].Rule != "MIGRATION_CONTENT_HASH_MISMATCH" {
		t.Fatalf("tampering was not rejected: %#v", report.Fail.Items)
	}
}
