package cmd

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"rza_migrate_db/internal/license"
)

func TestRunMigrateCheckRejectsRenamedExecutable(t *testing.T) {
	t.Parallel()
	if code := RunMigrateCheck([]string{"renamed_checker.exe"}); code != 1 {
		t.Fatalf("exit code mismatch: got=%d want=1", code)
	}
}

func TestRunMigrateCheckBuildInfoDoesNotRequireRepositoryOrLicense(t *testing.T) {
	if code := RunMigrateCheck([]string{"rza_migrate_check", "-build-info"}); code != 0 {
		t.Fatalf("exit code mismatch: got=%d want=0", code)
	}
}

func TestResolveCheckerBool(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		flagValue    string
		envValue     string
		defaultValue bool
		want         bool
		wantErr      bool
	}{
		{name: "default", defaultValue: true, want: true},
		{name: "flag", flagValue: "false", defaultValue: true, want: false},
		{name: "env", envValue: "yes", want: true},
		{name: "flag precedence", flagValue: "off", envValue: "on", want: false},
		{name: "invalid", flagValue: "sometimes", wantErr: true},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := resolveCheckerBool(tc.flagValue, tc.envValue, tc.defaultValue, "test")
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got=%t", got)
				}
				return
			}
			if err != nil || got != tc.want {
				t.Fatalf("got=(%t, %v) want=(%t, nil)", got, err, tc.want)
			}
		})
	}
}

func TestResolveCheckerWorkers(t *testing.T) {
	t.Parallel()
	if got, err := resolveCheckerWorkers("4", "9"); err != nil || got != 4 {
		t.Fatalf("flag precedence mismatch: got=(%d, %v)", got, err)
	}
	if got, err := resolveCheckerWorkers("", "6"); err != nil || got != 6 {
		t.Fatalf("environment workers mismatch: got=(%d, %v)", got, err)
	}
	for _, invalid := range []string{"", "zero", "0", "-1"} {
		if _, err := resolveCheckerWorkers(invalid, ""); err == nil {
			t.Fatalf("expected invalid worker value %q to fail", invalid)
		}
	}
}

func TestResolveCheckerProfile(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name, profileFlag, profileEnv, migrationFlag, migrationEnv string
		wantProfile, wantMigration                                 string
		wantErr                                                    bool
	}{
		{name: "defaults to authoring integrity", wantProfile: "authoring", wantMigration: "integrity"},
		{name: "full defaults to full migrations", profileFlag: "full", wantProfile: "full", wantMigration: "full"},
		{name: "authoring allows off", migrationFlag: "off", wantProfile: "authoring", wantMigration: "off"},
		{name: "flag wins over environment", profileFlag: "authoring", profileEnv: "full", migrationFlag: "integrity", migrationEnv: "full", wantProfile: "authoring", wantMigration: "integrity"},
		{name: "invalid profile", profileFlag: "loose", wantErr: true},
		{name: "invalid migration mode", migrationFlag: "legacy", wantErr: true},
		{name: "full cannot weaken migrations", profileFlag: "full", migrationFlag: "integrity", wantErr: true},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			profile, migration, err := resolveCheckerProfile(tc.profileFlag, tc.profileEnv, tc.migrationFlag, tc.migrationEnv)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got profile=%q migration=%q", profile, migration)
				}
				return
			}
			if err != nil || profile != tc.wantProfile || migration != tc.wantMigration {
				t.Fatalf("got=(%q,%q,%v) want=(%q,%q,nil)", profile, migration, err, tc.wantProfile, tc.wantMigration)
			}
		})
	}
}

func TestRunMigrateCheckDiscoversRepoFromSubfolder(t *testing.T) {
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
	now := time.Now()
	licText, err := license.Generate("test", now, now.AddDate(0, 1, 0))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(toolsDir, "license.key"), []byte(licText), 0644); err != nil {
		t.Fatal(err)
	}

	// Run checker simulating execution from tools/rza_migrate_check.exe without explicit -repo
	fakeExeInTools := filepath.Join(toolsDir, "rza_migrate_check.exe")
	_ = RunMigrateCheck([]string{
		fakeExeInTools,
		"-pause-on-exit", "false",
	})
	if _, err := os.Stat(filepath.Join(repo, "out", "output.json")); err != nil {
		t.Fatalf("expected output.json in repo/out/, got err: %v", err)
	}
}
