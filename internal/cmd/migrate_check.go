package cmd

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"rza_migrate_db/internal/buildinfo"
	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/engine"
	"rza_migrate_db/internal/license"
	"rza_migrate_db/internal/output"
)

// RunMigrateCheck runs the developer pre-commit workflow: apply safe db_object
// fixes, verify repository contracts, write output.json, and never generate a
// migration.
func RunMigrateCheck(args []string) int {
	startedAt := time.Now()
	if len(args) == 0 {
		args = []string{"rza_migrate_check"}
	}
	if err := config.ValidateExecutableName(args[0], config.DefaultMigrateCheckExecutableNameLock()); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}

	fs := flag.NewFlagSet("rza_migrate_check", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	repo := fs.String("repo", "", "repository root containing db_object/ and migrations/")
	perfNotes := fs.String("perf-notes", "", "include performance notes in out/output.json: true|false")
	failFast := fs.String("fail-fast", "", "stop verification after the first failing domain: true|false")
	parallelRead := fs.String("parallel-read", "", "enable bounded parallel validation: true|false")
	parallelWorkers := fs.String("parallel-workers", "", "number of validation workers: integer >= 1")
	fmtCheck := fs.Bool("check", false, "validate formatting without writing db_object (recommended for CI)")
	profile := fs.String("profile", "", "validation profile: authoring|full (default: authoring)")
	migrationValidation := fs.String("migration-validation", "", "migration validation: off|integrity|full (default: integrity for authoring)")
	pauseOnExit := fs.String("pause-on-exit", "", "pause before closing: true|false (automatic on interactive double-click)")
	showBuildInfo := fs.Bool("build-info", false, "print release build identity and exit")
	if err := fs.Parse(args[1:]); err != nil {
		return 1
	}
	if fs.NArg() != 0 {
		fmt.Fprintf(os.Stderr, "error: unexpected argument(s): %s\n", strings.Join(fs.Args(), " "))
		return 1
	}
	if *showBuildInfo {
		info := buildinfo.Current()
		fmt.Printf("version=%s\nbuild_time=%s\nsource_sha256=%s\n", info.Version, info.BuildTime, info.SourceSHA256)
		return 0
	}

	perfEnabled, err := resolveCheckerBool(*perfNotes, os.Getenv("RZA_PERF_NOTES"), false, "perf-notes")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}
	failFastEnabled, err := resolveCheckerBool(*failFast, os.Getenv("RZA_FAIL_FAST"), false, "fail-fast")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}
	parallelEnabled, err := resolveCheckerBool(*parallelRead, os.Getenv("RZA_PARALLEL_READ"), true, "parallel-read")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}
	// A launch with no arguments is the double-click workflow. Explicit CLI
	// invocations do not pause unless requested, so hooks and scripts stay non-interactive.
	pauseEnabled, err := resolveCheckerBool(*pauseOnExit, os.Getenv("RZA_PAUSE_ON_EXIT"), len(args) == 1 && checkerInteractiveConsole(), "pause-on-exit")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}

	resolvedRepo := config.DiscoverRepoRoot(*repo, args[0])
	cfg := config.Default(resolvedRepo)
	resolvedProfile, resolvedMigrationValidation, profileErr := resolveCheckerProfile(
		*profile,
		os.Getenv("RZA_VALIDATION_PROFILE"),
		*migrationValidation,
		os.Getenv("RZA_MIGRATION_VALIDATION"),
	)
	if profileErr != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", profileErr)
		return 1
	}
	cfg.ValidationProfile = resolvedProfile
	cfg.MigrationValidation = resolvedMigrationValidation
	cfg.EnablePerfNotes = perfEnabled
	// The checker always retains diagnostic notes such as formatted paths,
	// cache statistics, and the explicit no-generation mode.
	cfg.EnableOutputNotes = true
	cfg.EnableFmtCheck = *fmtCheck
	cfg.EnableFailFast = failFastEnabled
	cfg.EnableParallelRead = parallelEnabled
	if strings.TrimSpace(*parallelWorkers) != "" || strings.TrimSpace(os.Getenv("RZA_PARALLEL_WORKERS")) != "" {
		workers, workersErr := resolveCheckerWorkers(*parallelWorkers, os.Getenv("RZA_PARALLEL_WORKERS"))
		if workersErr != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", workersErr)
			return 1
		}
		cfg.ParallelWorkers = workers
	}

	if cfg.EnableLicence {
		if err := validateCheckerLicence(cfg, startedAt); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			checkerPause(pauseEnabled)
			return 1
		}
	}

	fmt.Printf("[start] app=rza_migrate_check repo=%s at=%s\n", filepath.Clean(cfg.RepoRoot), startedAt.Format(time.RFC3339))
	fmt.Printf("[mode] safe-fix db_object -> verify profile=%s migrations=%s -> output.json (migration generation disabled)\n", cfg.ValidationProfile, cfg.MigrationValidation)
	fmt.Printf("[policy] existing_table_match=%s\n", cfg.ExistingTableMatchMode)
	report, runErr := engine.New(cfg).Run("check")
	exitCode := 0
	if runErr != nil {
		fmt.Fprintf(os.Stderr, "[error] %v\n", runErr)
		exitCode = 1
	} else if report.Fail.Count > 0 {
		exitCode = 1
	}
	// Keep the structured result and its failure details on one stream. CI log
	// collectors may merge stdout/stderr out of order, which can corrupt an
	// otherwise actionable diagnostic line.
	output.PrintFailureDetails(os.Stdout, report)
	fmt.Printf("[result] fail=%d pass=%d migrations_generated=0 output=%s\n", report.Fail.Count, report.Pass.Count, filepath.Clean(cfg.OutFilePath()))
	fmt.Printf("[done] exit_code=%d duration=%s\n", exitCode, time.Since(startedAt).Round(time.Millisecond))
	checkerPause(pauseEnabled)
	return exitCode
}

func checkerDefaultRepo(arg0 string) string {
	return config.DiscoverRepoRoot("", arg0)
}

func checkerLooksLikeRepo(root string) bool {
	return config.LooksLikeRepo(root)
}

func resolveCheckerBool(flagValue, envValue string, defaultValue bool, name string) (bool, error) {
	value := strings.ToLower(strings.TrimSpace(flagValue))
	if value == "" {
		value = strings.ToLower(strings.TrimSpace(envValue))
	}
	if value == "" {
		return defaultValue, nil
	}
	switch value {
	case "1", "true", "on", "yes":
		return true, nil
	case "0", "false", "off", "no":
		return false, nil
	default:
		return false, fmt.Errorf("invalid %s value %q (allowed: true|false)", name, value)
	}
}

func resolveCheckerWorkers(flagValue, envValue string) (int, error) {
	value := strings.TrimSpace(flagValue)
	if value == "" {
		value = strings.TrimSpace(envValue)
	}
	workers, err := strconv.Atoi(value)
	if err != nil || workers < 1 {
		return 0, fmt.Errorf("invalid parallel-workers value %q (must be integer >= 1)", value)
	}
	return workers, nil
}

func resolveCheckerProfile(profileFlag, profileEnv, migrationFlag, migrationEnv string) (string, string, error) {
	profile := strings.ToLower(strings.TrimSpace(profileFlag))
	if profile == "" {
		profile = strings.ToLower(strings.TrimSpace(profileEnv))
	}
	if profile == "" {
		profile = config.ValidationProfileAuthoring
	}
	if profile != config.ValidationProfileAuthoring && profile != config.ValidationProfileFull {
		return "", "", fmt.Errorf("invalid profile value %q (allowed: authoring|full)", profile)
	}

	migration := strings.ToLower(strings.TrimSpace(migrationFlag))
	if migration == "" {
		migration = strings.ToLower(strings.TrimSpace(migrationEnv))
	}
	if migration == "" {
		if profile == config.ValidationProfileFull {
			migration = config.MigrationValidationFull
		} else {
			migration = config.MigrationValidationIntegrity
		}
	}
	if migration != config.MigrationValidationOff && migration != config.MigrationValidationIntegrity && migration != config.MigrationValidationFull {
		return "", "", fmt.Errorf("invalid migration-validation value %q (allowed: off|integrity|full)", migration)
	}
	if profile == config.ValidationProfileFull && migration != config.MigrationValidationFull {
		return "", "", fmt.Errorf("profile full requires migration-validation full")
	}
	return profile, migration, nil
}

func validateCheckerLicence(cfg config.AppConfig, now time.Time) error {
	path := cfg.ResolveLicencePath()
	info, err := license.ValidateFile(path, now)
	if err == nil {
		fmt.Printf("[license] valid owner=%s expires_at=%s\n", info.Owner, info.ExpiresAt.Format("2006-01-02"))
		return nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("license file not found: %s (application blocked)", filepath.Clean(path))
	}
	if errors.Is(err, license.ErrExpired) {
		return fmt.Errorf("license expired at %s", info.ExpiresAt.Format("2006-01-02"))
	}
	return fmt.Errorf("license invalid at %s: %w", filepath.Clean(path), err)
}

func checkerPause(enabled bool) {
	if !enabled {
		return
	}
	fmt.Print("Press Enter to close...")
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
}

func checkerInteractiveConsole() bool {
	stdinInfo, stdinErr := os.Stdin.Stat()
	stdoutInfo, stdoutErr := os.Stdout.Stat()
	return stdinErr == nil && stdoutErr == nil &&
		(stdinInfo.Mode()&os.ModeCharDevice) != 0 &&
		(stdoutInfo.Mode()&os.ModeCharDevice) != 0
}
