package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDefaultConfigValidatePasses(t *testing.T) {
	t.Parallel()

	cfg := Default(".")
	if !cfg.EnableGenAutoIndent {
		t.Fatalf("EnableGenAutoIndent must default to true")
	}
	if cfg.EnablePerfNotes {
		t.Fatalf("EnablePerfNotes must default to false")
	}
	if cfg.EnableOutputNotes {
		t.Fatalf("EnableOutputNotes must default to false")
	}
	if !cfg.EnableLicence {
		t.Fatalf("EnableLicence must default to true")
	}
	if cfg.LicenceFile != "license.key" {
		t.Fatalf("LicenceFile must default to license.key, got %q", cfg.LicenceFile)
	}
	if cfg.EnableFailFast {
		t.Fatalf("EnableFailFast must default to false")
	}
	if !cfg.EnableParallelRead {
		t.Fatalf("EnableParallelRead must default to true")
	}
	if cfg.ParallelWorkers != 8 {
		t.Fatalf("ParallelWorkers must default to 8, got %d", cfg.ParallelWorkers)
	}
	if cfg.ValidationProfile != ValidationProfileFull || cfg.MigrationValidation != MigrationValidationFull {
		t.Fatalf("default engine profile must remain full/full, got %s/%s", cfg.ValidationProfile, cfg.MigrationValidation)
	}
	if err := cfg.Validate(); err != nil {
		t.Fatalf("default config should be valid, got error: %v", err)
	}
}

func TestDefaultNormalizesRepoRootToAbsoluteAndRepoPathAcceptsAbsoluteInput(t *testing.T) {
	t.Parallel()
	cfg := Default(filepath.Join("relative", "repo"))
	if !filepath.IsAbs(cfg.RepoRoot) {
		t.Fatalf("RepoRoot must be absolute, got %q", cfg.RepoRoot)
	}
	abs := filepath.Join(cfg.RepoRoot, "db_object", "tables", "dbo.A.sql")
	if got := cfg.RepoPath(abs); got != filepath.Clean(abs) {
		t.Fatalf("RepoPath duplicated an absolute root: got=%q want=%q", got, filepath.Clean(abs))
	}
}

func TestValidationProfileHardLocks(t *testing.T) {
	t.Parallel()
	cfg := Default(".")
	cfg.ValidationProfile = ValidationProfileAuthoring
	cfg.MigrationValidation = MigrationValidationIntegrity
	if err := cfg.Validate(); err != nil {
		t.Fatalf("authoring/integrity must validate: %v", err)
	}
	cfg.ValidationProfile = ValidationProfileFull
	if err := cfg.Validate(); err == nil || !strings.Contains(err.Error(), "requires MigrationValidation full") {
		t.Fatalf("full/integrity must fail hard lock, got: %v", err)
	}
}

func TestValidateFailsWhenDecisionWhitelistBroken(t *testing.T) {
	t.Parallel()

	cfg := Default(".")
	delete(cfg.AllowedDecisions, "run")

	err := cfg.Validate()
	if err == nil {
		t.Fatalf("expected validation error when required decision is missing")
	}
	if !strings.Contains(err.Error(), "AllowedDecisions missing: run") {
		t.Fatalf("unexpected validation error: %v", err)
	}
}

func TestValidateFailsWhenParallelWorkersInvalid(t *testing.T) {
	t.Parallel()

	cfg := Default(".")
	cfg.ParallelWorkers = 0

	err := cfg.Validate()
	if err == nil {
		t.Fatalf("expected validation error when ParallelWorkers < 1")
	}
	if !strings.Contains(err.Error(), "ParallelWorkers must be >= 1") {
		t.Fatalf("unexpected validation error: %v", err)
	}
}

func TestDefaultExistingTableMatchModeIsBootstrapCompatible(t *testing.T) {
	t.Parallel()

	cfg := Default(".")
	if cfg.ExistingTableMatchMode != ExistingTableMatchBootstrapCompatible {
		t.Fatalf("unexpected existing table match mode: %q", cfg.ExistingTableMatchMode)
	}
	if err := cfg.Validate(); err != nil {
		t.Fatalf("default bootstrap-compatible config must validate: %v", err)
	}

	cfg.ExistingTableMatchMode = "ignore-everything"
	if err := cfg.Validate(); err == nil {
		t.Fatal("unsafe existing table match mode must fail the config hard lock")
	}
}

func TestRepoDiscoveryAndLicencePathResolution(t *testing.T) {
	tmp := t.TempDir()

	// Initially, empty dir does not look like repo
	if LooksLikeRepo(tmp) {
		t.Fatalf("empty directory should not look like repo")
	}

	// Create db_object and migrations
	dbObjDir := filepath.Join(tmp, "db_object")
	migDir := filepath.Join(tmp, "migrations")
	toolsDir := filepath.Join(tmp, "tools")
	if err := os.MkdirAll(dbObjDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(toolsDir, 0755); err != nil {
		t.Fatal(err)
	}

	if !LooksLikeRepo(tmp) {
		t.Fatalf("expected tmp to look like repo")
	}
	if LooksLikeRepo(toolsDir) {
		t.Fatalf("tools directory alone should not look like repo")
	}

	// DiscoverRepoRoot with explicit path
	if got := DiscoverRepoRoot(tmp, ""); got != tmp {
		t.Fatalf("expected explicit repo %q, got %q", tmp, got)
	}

	// DiscoverRepoRoot from within tools/ directory
	fakeExeInTools := filepath.Join(toolsDir, "rza_migrate_check.exe")
	if got := DiscoverRepoRoot("", fakeExeInTools); got != tmp {
		t.Fatalf("expected DiscoverRepoRoot from tools exe to find parent %q, got %q", tmp, got)
	}

	// Test ResolveLicencePath
	cfg := Default(tmp)

	// Case 1: license.key in tools/
	toolsLic := filepath.Join(toolsDir, "license.key")
	if err := os.WriteFile(toolsLic, []byte("tools-license"), 0644); err != nil {
		t.Fatal(err)
	}
	if got := cfg.ResolveLicencePath(); got != toolsLic {
		t.Fatalf("expected ResolveLicencePath to prefer tools/license.key, got %q want %q", got, toolsLic)
	}

	// Case 2: ResolveLicencePathForWrite when tools/ exists
	if got := cfg.ResolveLicencePathForWrite(); got != toolsLic {
		t.Fatalf("expected ResolveLicencePathForWrite to write to tools/license.key, got %q want %q", got, toolsLic)
	}

	// Case 3: license.key at root when tools/license.key doesn't exist
	if err := os.Remove(toolsLic); err != nil {
		t.Fatal(err)
	}
	rootLic := filepath.Join(tmp, "license.key")
	if err := os.WriteFile(rootLic, []byte("root-license"), 0644); err != nil {
		t.Fatal(err)
	}
	if got := cfg.ResolveLicencePath(); got != rootLic {
		t.Fatalf("expected ResolveLicencePath to find root license.key, got %q want %q", got, rootLic)
	}
}
