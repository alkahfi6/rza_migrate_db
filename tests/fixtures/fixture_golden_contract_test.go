package fixtures_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

type expectation struct {
	QAID                string   `json:"qa_id"`
	Command             string   `json:"command"`
	ExpectedExitCode    int      `json:"expected_exit_code"`
	ExpectedFailRules   []string `json:"expected_fail_rules"`
	ExpectedWarnReasons []string `json:"expected_warn_reasons"`
	Summary             string   `json:"summary"`
}

type outputItem struct {
	Rule     string `json:"rule"`
	Severity int    `json:"severity"`
	File     string `json:"file"`
	Message  string `json:"message"`
	Reason   string `json:"reason"`
	Action   string `json:"action"`
}

type outputBucket struct {
	Count int          `json:"count"`
	Items []outputItem `json:"items"`
}

type outputNaming struct {
	Pattern     string `json:"pattern"`
	LineEndings string `json:"line_endings"`
}

type outputNewMigrations struct {
	Naming outputNaming `json:"naming"`
	Files  []string     `json:"files"`
}

type outputGolden struct {
	Fail          outputBucket        `json:"fail"`
	Pass          outputBucket        `json:"pass"`
	NewMigrations outputNewMigrations `json:"new_migrations"`
	Notes         map[string]any      `json:"notes"`
}

var scenarios = []string{
	"01_hard_fail_rule",
	"02_folder_contract",
	"03_forbidden_sql",
	"04_pairing_missing_down",
	"05_version_format_invalid",
	"06_ordering_determinism",
	"07_vk_whitespace_only",
	"08_reconstruction_unsupported_sql",
	"09_file_object_mismatch",
	"10_inline_default_not_allowed",
	"11_destructive_drop_table_manual",
	"12_transaction_error_throw",
	"13_audit_wrapper_contract",
	"14_down_data_restore_warning",
	"15_patch_canonicalization",
	"16_patch_duplication_once",
	"17_patch_source_integrity",
	"18_output_json_contract",
	"19_repo_hygiene_constraints",
	"20_cli_exit_code_contract",
	"21_wrapper_action_whitelist",
	"22_wrapper_reason_whitelist",
	"23_wrapper_run_key_newid",
	"24_constraint_unnamed",
	"25_wrapper_precheck_skip",
	"26_alter_column_precheck_map",
}

func TestFixtureGoldenContracts(t *testing.T) {
	root := repoRoot(t)
	fixtureRoot := filepath.Join(root, "fixtures", "qa_rules")
	if _, err := os.Stat(fixtureRoot); err != nil {
		t.Fatalf("missing fixture root %s: %v", fixtureRoot, err)
	}

	for _, scenario := range scenarios {
		scenario := scenario
		t.Run(scenario, func(t *testing.T) {
			base := filepath.Join(fixtureRoot, scenario)
			expectPath := filepath.Join(base, "expected", "expectation.json")
			goldenPath := filepath.Join(base, "expected", "output.golden.json")

			expect := readExpectation(t, expectPath)
			goldenRaw := readJSONRaw(t, goldenPath)
			assertRootKeys(t, goldenRaw, []string{"fail", "pass", "new_migrations", "notes"})

			var golden outputGolden
			if err := json.Unmarshal(mustReadFile(t, goldenPath), &golden); err != nil {
				t.Fatalf("invalid golden JSON at %s: %v", goldenPath, err)
			}

			if golden.Fail.Count != len(golden.Fail.Items) {
				t.Fatalf("fail.count mismatch fail.items length: %d != %d", golden.Fail.Count, len(golden.Fail.Items))
			}
			if golden.Pass.Count != len(golden.Pass.Items) {
				t.Fatalf("pass.count mismatch pass.items length: %d != %d", golden.Pass.Count, len(golden.Pass.Items))
			}

			if golden.NewMigrations.Naming.Pattern != "YYYYMMDDHHMMSS<seq>_<category>_<object>_v<k>.(up|down).sql" {
				t.Fatalf("unexpected naming.pattern: %s", golden.NewMigrations.Naming.Pattern)
			}
			if golden.NewMigrations.Naming.LineEndings != "LF" {
				t.Fatalf("unexpected naming.line_endings: %s", golden.NewMigrations.Naming.LineEndings)
			}

			notesQAID, ok := golden.Notes["qa_id"].(string)
			if !ok {
				t.Fatalf("notes.qa_id missing or not string")
			}
			if notesQAID != expect.QAID {
				t.Fatalf("notes.qa_id mismatch: %s != %s", notesQAID, expect.QAID)
			}

			actualFailRules := make([]string, 0, len(golden.Fail.Items))
			for _, item := range golden.Fail.Items {
				actualFailRules = append(actualFailRules, item.Rule)
			}
			for _, expectedRule := range expect.ExpectedFailRules {
				if !contains(actualFailRules, expectedRule) {
					t.Fatalf("missing expected fail rule %q in golden fail items", expectedRule)
				}
			}

			actualWarnReasons := make([]string, 0, len(golden.Pass.Items))
			for _, item := range golden.Pass.Items {
				actualWarnReasons = append(actualWarnReasons, item.Reason)
			}
			for _, expectedReason := range expect.ExpectedWarnReasons {
				if !contains(actualWarnReasons, expectedReason) {
					t.Fatalf("missing expected warn reason %q in golden pass items", expectedReason)
				}
			}

			if expect.ExpectedExitCode == 1 && golden.Fail.Count == 0 {
				t.Fatalf("expected exit code 1 requires fail.count > 0")
			}
			if expect.ExpectedExitCode == 0 && golden.Fail.Count > 0 {
				t.Fatalf("expected exit code 0 requires fail.count == 0")
			}
		})
	}
}

func repoRoot(t *testing.T) string {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get wd: %v", err)
	}

	curr := wd
	for {
		if _, err := os.Stat(filepath.Join(curr, "go.mod")); err == nil {
			return curr
		}
		parent := filepath.Dir(curr)
		if parent == curr {
			t.Fatalf("unable to locate repo root from %s", wd)
		}
		curr = parent
	}
}

func readExpectation(t *testing.T, path string) expectation {
	t.Helper()
	var e expectation
	if err := json.Unmarshal(mustReadFile(t, path), &e); err != nil {
		t.Fatalf("invalid expectation JSON at %s: %v", path, err)
	}
	return e
}

func readJSONRaw(t *testing.T, path string) map[string]json.RawMessage {
	t.Helper()
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(mustReadFile(t, path), &raw); err != nil {
		t.Fatalf("invalid JSON at %s: %v", path, err)
	}
	return raw
}

func assertRootKeys(t *testing.T, raw map[string]json.RawMessage, keys []string) {
	t.Helper()
	for _, k := range keys {
		if _, ok := raw[k]; !ok {
			t.Fatalf("missing root key %q", k)
		}
	}
}

func mustReadFile(t *testing.T, path string) []byte {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed reading %s: %v", path, err)
	}
	return b
}

func contains(list []string, want string) bool {
	for _, got := range list {
		if got == want {
			return true
		}
	}
	return false
}
