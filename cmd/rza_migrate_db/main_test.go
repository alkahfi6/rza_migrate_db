package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestRunInvalidLegacyAuditMode(t *testing.T) {
	t.Parallel()

	exit := run([]string{
		"rza_migrate_db",
		"verify",
		"-legacy-audit-mode",
		"invalid",
	})
	if exit != 1 {
		t.Fatalf("expected exit code 1 for invalid legacy-audit-mode, got %d", exit)
	}
}

func TestRunFailsWhenExecutableRenamed(t *testing.T) {
	t.Parallel()

	exit := run([]string{
		"renamed_migrate_tool.exe",
		"verify",
	})
	if exit != 1 {
		t.Fatalf("expected exit code 1 for renamed executable, got %d", exit)
	}
}

func TestRunBuildInfoDoesNotRequireRepositoryOrLicense(t *testing.T) {
	if exit := run([]string{"rza_migrate_db", "-build-info"}); exit != 0 {
		t.Fatalf("expected build info exit code 0, got %d", exit)
	}
}

func TestRunPerfNotesFlagWritesNotesToOutputJSON(t *testing.T) {
	repo := t.TempDir()
	t.Setenv("RZA_PERF_NOTES", "")
	t.Setenv("RZA_FAIL_FAST", "")
	t.Setenv("RZA_PARALLEL_READ", "")
	t.Setenv("RZA_PARALLEL_WORKERS", "")
	t.Setenv("RZA_GEN_AUTO_INDENT", "")
	t.Setenv("RZA_VERIFY_PROGRAMMABILITY_NEW_ONLY", "")
	t.Setenv("RZA_LEGACY_AUDIT_MODE", "")

	licenseExit := run([]string{
		"rza_migrate_db",
		"license_gen",
		"-repo",
		repo,
		"-license-owner",
		"test",
		"-license-months",
		"1",
	})
	if licenseExit != 0 {
		t.Fatalf("license_gen exit=%d", licenseExit)
	}

	_ = run([]string{
		"rza_migrate_db",
		"verify",
		"-repo",
		repo,
		"-perf-notes",
		"true",
	})

	raw, err := os.ReadFile(filepath.Join(repo, "out", "output.json"))
	if err != nil {
		t.Fatalf("read output.json: %v", err)
	}
	var payload struct {
		Notes map[string]any `json:"notes"`
	}
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("unmarshal output.json: %v", err)
	}
	if payload.Notes["command"] != "verify" {
		t.Fatalf("expected command note in output.json, got notes=%v", payload.Notes)
	}
	if _, ok := payload.Notes["perf_ms"]; !ok {
		t.Fatalf("expected perf_ms note in output.json, got notes=%v", payload.Notes)
	}
}

func TestResolveLegacyAuditMode(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name      string
		flagValue string
		envValue  string
		wantMode  string
		wantErr   bool
	}{
		{
			name:      "default_fail",
			flagValue: "",
			envValue:  "",
			wantMode:  "fail",
			wantErr:   false,
		},
		{
			name:      "env_used_when_flag_empty",
			flagValue: "",
			envValue:  "fail",
			wantMode:  "fail",
			wantErr:   false,
		},
		{
			name:      "flag_overrides_env",
			flagValue: "fail",
			envValue:  "fail",
			wantMode:  "fail",
			wantErr:   false,
		},
		{
			name:      "invalid_env",
			flagValue: "",
			envValue:  "invalid",
			wantMode:  "",
			wantErr:   true,
		},
		{
			name:      "invalid_flag",
			flagValue: "invalid",
			envValue:  "fail",
			wantMode:  "",
			wantErr:   true,
		},
		{
			name:      "warn_not_allowed",
			flagValue: "warn",
			envValue:  "",
			wantMode:  "",
			wantErr:   true,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := resolveLegacyAuditMode(tc.flagValue, tc.envValue)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got mode=%q", got)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.wantMode {
				t.Fatalf("mode mismatch: got=%q want=%q", got, tc.wantMode)
			}
		})
	}
}

func TestResolveGenAutoIndent(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name      string
		flagValue string
		envValue  string
		want      bool
		wantErr   bool
	}{
		{name: "default_true", flagValue: "", envValue: "", want: true, wantErr: false},
		{name: "flag_false", flagValue: "false", envValue: "", want: false, wantErr: false},
		{name: "flag_true", flagValue: "true", envValue: "", want: true, wantErr: false},
		{name: "env_false", flagValue: "", envValue: "off", want: false, wantErr: false},
		{name: "env_true", flagValue: "", envValue: "yes", want: true, wantErr: false},
		{name: "flag_overrides_env", flagValue: "false", envValue: "true", want: false, wantErr: false},
		{name: "invalid", flagValue: "maybe", envValue: "", want: false, wantErr: true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := resolveGenAutoIndent(tc.flagValue, tc.envValue)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got value=%v", got)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Fatalf("value mismatch: got=%v want=%v", got, tc.want)
			}
		})
	}
}

func TestResolveVerifyProgrammabilityNewOnly(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name      string
		flagValue string
		envValue  string
		want      bool
		wantErr   bool
	}{
		{name: "default_false", flagValue: "", envValue: "", want: false, wantErr: false},
		{name: "flag_true", flagValue: "true", envValue: "", want: true, wantErr: false},
		{name: "flag_false", flagValue: "false", envValue: "", want: false, wantErr: false},
		{name: "env_true", flagValue: "", envValue: "on", want: true, wantErr: false},
		{name: "env_false", flagValue: "", envValue: "off", want: false, wantErr: false},
		{name: "flag_overrides_env", flagValue: "false", envValue: "true", want: false, wantErr: false},
		{name: "invalid", flagValue: "maybe", envValue: "", want: false, wantErr: true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := resolveVerifyProgrammabilityNewOnly(tc.flagValue, tc.envValue)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got value=%v", got)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Fatalf("value mismatch: got=%v want=%v", got, tc.want)
			}
		})
	}
}

func TestResolvePerfNotes(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name      string
		flagValue string
		envValue  string
		want      bool
		wantErr   bool
	}{
		{name: "default_false", flagValue: "", envValue: "", want: false, wantErr: false},
		{name: "flag_true", flagValue: "true", envValue: "", want: true, wantErr: false},
		{name: "flag_false", flagValue: "false", envValue: "", want: false, wantErr: false},
		{name: "env_true", flagValue: "", envValue: "on", want: true, wantErr: false},
		{name: "env_false", flagValue: "", envValue: "off", want: false, wantErr: false},
		{name: "flag_overrides_env", flagValue: "false", envValue: "true", want: false, wantErr: false},
		{name: "invalid", flagValue: "maybe", envValue: "", want: false, wantErr: true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := resolvePerfNotes(tc.flagValue, tc.envValue)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got value=%v", got)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Fatalf("value mismatch: got=%v want=%v", got, tc.want)
			}
		})
	}
}

func TestResolveFailFast(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name      string
		flagValue string
		envValue  string
		want      bool
		wantErr   bool
	}{
		{name: "default_false", flagValue: "", envValue: "", want: false, wantErr: false},
		{name: "flag_true", flagValue: "true", envValue: "", want: true, wantErr: false},
		{name: "flag_false", flagValue: "false", envValue: "", want: false, wantErr: false},
		{name: "env_true", flagValue: "", envValue: "on", want: true, wantErr: false},
		{name: "env_false", flagValue: "", envValue: "off", want: false, wantErr: false},
		{name: "flag_overrides_env", flagValue: "false", envValue: "true", want: false, wantErr: false},
		{name: "invalid", flagValue: "maybe", envValue: "", want: false, wantErr: true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := resolveFailFast(tc.flagValue, tc.envValue)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got value=%v", got)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Fatalf("value mismatch: got=%v want=%v", got, tc.want)
			}
		})
	}
}

func TestResolveParallelRead(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name      string
		flagValue string
		envValue  string
		want      bool
		wantErr   bool
	}{
		{name: "default_true", flagValue: "", envValue: "", want: true, wantErr: false},
		{name: "flag_true", flagValue: "true", envValue: "", want: true, wantErr: false},
		{name: "flag_false", flagValue: "false", envValue: "", want: false, wantErr: false},
		{name: "env_true", flagValue: "", envValue: "on", want: true, wantErr: false},
		{name: "env_false", flagValue: "", envValue: "off", want: false, wantErr: false},
		{name: "flag_overrides_env", flagValue: "false", envValue: "true", want: false, wantErr: false},
		{name: "invalid", flagValue: "maybe", envValue: "", want: false, wantErr: true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := resolveParallelRead(tc.flagValue, tc.envValue)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got value=%v", got)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Fatalf("value mismatch: got=%v want=%v", got, tc.want)
			}
		})
	}
}

func TestResolveParallelWorkers(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name      string
		flagValue string
		envValue  string
		want      int
		wantErr   bool
	}{
		{name: "default_eight", flagValue: "", envValue: "", want: 8, wantErr: false},
		{name: "flag_value", flagValue: "8", envValue: "", want: 8, wantErr: false},
		{name: "env_value", flagValue: "", envValue: "6", want: 6, wantErr: false},
		{name: "flag_overrides_env", flagValue: "3", envValue: "9", want: 3, wantErr: false},
		{name: "invalid_text", flagValue: "abc", envValue: "", want: 0, wantErr: true},
		{name: "invalid_zero", flagValue: "0", envValue: "", want: 0, wantErr: true},
		{name: "invalid_negative", flagValue: "-2", envValue: "", want: 0, wantErr: true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := resolveParallelWorkers(tc.flagValue, tc.envValue)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got value=%d", got)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Fatalf("value mismatch: got=%d want=%d", got, tc.want)
			}
		})
	}
}

func TestResolvePatchFormatMode(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		flag, env, want string
		wantErr         bool
	}{
		{want: "off"},
		{flag: "SAFE", want: "safe"},
		{env: "strict", want: "strict"},
		{flag: "off", env: "strict", want: "off"},
		{flag: "unsafe", wantErr: true},
	} {
		got, err := resolvePatchFormatMode(tc.flag, tc.env)
		if tc.wantErr {
			if err == nil {
				t.Fatalf("expected error for %q", tc.flag)
			}
			continue
		}
		if err != nil || got != tc.want {
			t.Fatalf("resolve(%q,%q)=%q,%v want %q", tc.flag, tc.env, got, err, tc.want)
		}
	}
}

func TestRunDiscoversRepoFromSubfolder(t *testing.T) {
	repo := t.TempDir()
	dbObjDir := filepath.Join(repo, "db_object")
	migDir := filepath.Join(repo, "migrations")
	toolsDir := filepath.Join(repo, "tools")
	for _, d := range []string{dbObjDir, migDir, toolsDir} {
		if err := os.MkdirAll(d, 0755); err != nil {
			t.Fatal(err)
		}
	}

	// Generate license into tools/license.key
	licenseExit := run([]string{
		"rza_migrate_db",
		"license_gen",
		"-repo", repo,
		"-license-owner", "test",
		"-license-months", "1",
		"-license-out", filepath.Join(toolsDir, "license.key"),
	})
	if licenseExit != 0 {
		t.Fatalf("license_gen exit=%d", licenseExit)
	}

	// Run verify by simulating execution from tools/rza_migrate_db.exe without explicit -repo
	fakeExeInTools := filepath.Join(toolsDir, "rza_migrate_db.exe")
	_ = run([]string{
		fakeExeInTools,
		"verify",
		"-pause-on-exit", "false",
	})
	if _, err := os.Stat(filepath.Join(repo, "out", "output.json")); err != nil {
		t.Fatalf("expected output.json in repo/out/, got err: %v", err)
	}
}
