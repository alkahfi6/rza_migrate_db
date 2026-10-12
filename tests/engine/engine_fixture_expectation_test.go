package engine_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/engine"
	"rza_migrate_db/internal/output"
)

type fixtureExpectation struct {
	QAID                string   `json:"qa_id"`
	Command             string   `json:"command"`
	ExpectedExitCode    int      `json:"expected_exit_code"`
	ExpectedFailRules   []string `json:"expected_fail_rules"`
	ExpectedWarnReasons []string `json:"expected_warn_reasons"`
}

func TestEngineAllFixturesAgainstExpectations(t *testing.T) {
	t.Parallel()

	scenarios := []string{
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

	repo := repoRoot(t)
	for _, scenario := range scenarios {
		scenario := scenario
		t.Run(scenario, func(t *testing.T) {
			srcScenario := filepath.Join(repo, "fixtures", "qa_rules", scenario)
			exp := readFixtureExpectation(t, filepath.Join(srcScenario, "expected", "expectation.json"))

			workRoot := t.TempDir()
			if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
				t.Fatalf("copy fixture input failed: %v", err)
			}

			cfg := config.Default(workRoot)
			eng := engine.New(cfg)

			commands := splitCommands(exp.Command)
			if len(commands) == 0 {
				t.Fatalf("expectation command is empty")
			}

			var lastFailRules []string
			var lastWarnReasons []string

			for _, cmd := range commands {
				report, err := eng.Run(cmd)
				if err != nil {
					t.Fatalf("engine run failed command=%s: %v", cmd, err)
				}

				gotExit := 0
				if report.Fail.Count > 0 {
					gotExit = 1
				}
				if gotExit != exp.ExpectedExitCode {
					t.Fatalf("exit mismatch command=%s got=%d expect=%d fail_count=%d", cmd, gotExit, exp.ExpectedExitCode, report.Fail.Count)
				}

				if _, err := os.Stat(filepath.Join(workRoot, "out", "output.json")); err != nil {
					t.Fatalf("output.json missing after command %s: %v", cmd, err)
				}

				lastFailRules = collectFailRules(report)
				lastWarnReasons = collectWarnReasons(report)
			}

			for _, want := range exp.ExpectedFailRules {
				if !containsString(lastFailRules, want) {
					t.Fatalf("missing expected fail rule %q in scenario %s (got: %v)", want, scenario, lastFailRules)
				}
			}
			for _, want := range exp.ExpectedWarnReasons {
				if !containsString(lastWarnReasons, want) {
					t.Fatalf("missing expected warn reason %q in scenario %s (got: %v)", want, scenario, lastWarnReasons)
				}
			}
		})
	}
}

func readFixtureExpectation(t *testing.T, path string) fixtureExpectation {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read expectation failed: %v", err)
	}
	var exp fixtureExpectation
	if err := json.Unmarshal(data, &exp); err != nil {
		t.Fatalf("parse expectation failed: %v", err)
	}
	return exp
}

func splitCommands(commandField string) []string {
	parts := strings.Split(commandField, "|")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		cmd := strings.TrimSpace(p)
		if cmd != "" {
			out = append(out, cmd)
		}
	}
	return out
}

func collectFailRules(report output.Report) []string {
	out := make([]string, 0, len(report.Fail.Items))
	for _, item := range report.Fail.Items {
		out = append(out, item.Rule)
	}
	return out
}

func collectWarnReasons(report output.Report) []string {
	out := make([]string, 0, len(report.Pass.Items))
	for _, item := range report.Pass.Items {
		out = append(out, item.Reason)
	}
	return out
}

func containsString(list []string, want string) bool {
	for _, got := range list {
		if got == want {
			return true
		}
	}
	return false
}
