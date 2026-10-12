package output

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrintFailureDetailsIncludesEveryDiagnosticField(t *testing.T) {
	report := NewReport("pattern", "lf")
	report.AddFail(Item{
		Rule:       "PATCH_SOURCE_MISSING",
		Severity:   3,
		File:       "patch/Fix Status",
		Object:     "fix_status",
		Location:   "migrations/100_patch_fix_status_v1.up.sql",
		Message:    "patch source is missing",
		Suggestion: "restore and commit the patch source",
		Reason:     "DRIFT_DETECTED",
		Action:     "FAIL",
	})
	report.AddFail(Item{Rule: "SECOND_RULE", Severity: 2, File: "db_object/tables/dbo.T.sql", Message: "second failure"})

	var console bytes.Buffer
	PrintFailureDetails(&console, report)
	got := console.String()
	for _, want := range []string{
		"[fail-summary] total=2",
		"[fail 1/2]",
		"rule=PATCH_SOURCE_MISSING",
		`file="patch/Fix Status"`,
		`object="fix_status"`,
		`location="migrations/100_patch_fix_status_v1.up.sql"`,
		"reason=DRIFT_DETECTED",
		"action=FAIL",
		`message="patch source is missing"`,
		`suggestion="restore and commit the patch source"`,
		"[fail 2/2]",
		"rule=SECOND_RULE",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("console output is missing %q:\n%s", want, got)
		}
	}
}

func TestPrintFailureDetailsIsSilentWhenSuccessful(t *testing.T) {
	var console bytes.Buffer
	PrintFailureDetails(&console, NewReport("pattern", "lf"))
	if console.Len() != 0 {
		t.Fatalf("successful report produced failure output: %q", console.String())
	}
}
