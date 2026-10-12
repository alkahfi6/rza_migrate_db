package engine_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/engine"
)

type expectedContract struct {
	QAID             string `json:"qa_id"`
	Command          string `json:"command"`
	ExpectedExitCode int    `json:"expected_exit_code"`
}

func TestEngineCommandsAgainstFixtureSubset(t *testing.T) {
	t.Parallel()

	cases := []struct {
		scenario string
		command  string
	}{
		{scenario: "01_hard_fail_rule", command: "verify"},
		{scenario: "06_ordering_determinism", command: "gen"},
		{scenario: "09_file_object_mismatch", command: "verify"},
		{scenario: "10_inline_default_not_allowed", command: "verify"},
		{scenario: "12_transaction_error_throw", command: "verify"},
		{scenario: "13_audit_wrapper_contract", command: "verify"},
		{scenario: "25_wrapper_precheck_skip", command: "gen"},
		{scenario: "26_alter_column_precheck_map", command: "gen"},
	}

	repoRoot := repoRoot(t)
	for _, tc := range cases {
		tc := tc
		t.Run(tc.scenario+"_"+tc.command, func(t *testing.T) {
			srcScenario := filepath.Join(repoRoot, "fixtures", "qa_rules", tc.scenario)
			expect := readExpectedContract(t, filepath.Join(srcScenario, "expected", "expectation.json"))
			if expect.Command != tc.command {
				t.Fatalf("scenario command mismatch: expectation=%s test=%s", expect.Command, tc.command)
			}

			workRoot := t.TempDir()
			if err := copyDir(filepath.Join(srcScenario, "input"), workRoot); err != nil {
				t.Fatalf("copy fixture input failed: %v", err)
			}

			cfg := config.Default(workRoot)
			eng := engine.New(cfg)
			report, err := eng.Run(tc.command)
			if err != nil {
				t.Fatalf("engine run failed: %v", err)
			}

			gotExit := 0
			if report.Fail.Count > 0 {
				gotExit = 1
			}
			if gotExit != expect.ExpectedExitCode {
				t.Fatalf("exit mismatch: got=%d expect=%d fail_count=%d", gotExit, expect.ExpectedExitCode, report.Fail.Count)
			}

			if _, err := os.Stat(filepath.Join(workRoot, "out", "output.json")); err != nil {
				t.Fatalf("out/output.json was not written: %v", err)
			}
		})
	}
}

func readExpectedContract(t *testing.T, path string) expectedContract {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read expectation failed: %v", err)
	}
	var c expectedContract
	if err := json.Unmarshal(data, &c); err != nil {
		t.Fatalf("parse expectation failed: %v", err)
	}
	return c
}

func copyDir(src, dst string) error {
	return filepath.WalkDir(src, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)

		if d.IsDir() {
			return os.MkdirAll(target, 0o755)
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, 0o644)
	})
}

func repoRoot(t *testing.T) string {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd failed: %v", err)
	}

	curr := wd
	for {
		if _, err := os.Stat(filepath.Join(curr, "go.mod")); err == nil {
			return curr
		}
		next := filepath.Dir(curr)
		if next == curr {
			t.Fatalf("go.mod not found from %s", wd)
		}
		curr = next
	}
}
