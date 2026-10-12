package verify

import (
	"strings"
	"testing"
)

func TestSQLWithoutMigrationBodyKeepsWrapperAroundQuoteHeavyPatch(t *testing.T) {
	sql := "BEGIN TRY\nBEGIN TRAN;\n-- MIGRATION_BODY_BEGIN\nUPDATE T SET C='select ''x'' from #t';\n-- MIGRATION_BODY_END\nCOMMIT TRAN;\nEND TRY\nBEGIN CATCH\nROLLBACK TRAN;\n;THROW;\nEND CATCH;\n"
	got := sqlWithoutMigrationBody(sql)
	if strings.Contains(got, "UPDATE T") {
		t.Fatalf("migration body was not excluded: %s", got)
	}
	for _, required := range []string{"BEGIN TRY", "BEGIN TRAN", "COMMIT TRAN", "BEGIN CATCH", "ROLLBACK TRAN", ";THROW"} {
		if !strings.Contains(got, required) {
			t.Fatalf("wrapper token %q was lost: %s", required, got)
		}
	}
}
