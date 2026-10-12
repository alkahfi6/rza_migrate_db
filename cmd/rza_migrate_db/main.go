package main

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

func main() {
	exitCode := run(os.Args)
	os.Exit(exitCode)
}

func run(args []string) int {
	startedAt := time.Now()
	if len(args) > 0 {
		if err := config.ValidateExecutableName(args[0], config.DefaultMigrateExecutableNameLock()); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return 1
		}
	}
	if len(args) == 2 && (args[1] == "--build-info" || args[1] == "-build-info") {
		info := buildinfo.Current()
		fmt.Printf("version=%s\nbuild_time=%s\nsource_sha256=%s\n", info.Version, info.BuildTime, info.SourceSHA256)
		return 0
	}

	command := ""
	flagArgs := []string{}
	implicitDefaultMode := false
	if len(args) < 2 || strings.HasPrefix(args[1], "-") {
		// Double-click / no explicit command: run safe default pipeline.
		command = "verify_then_gen"
		implicitDefaultMode = true
		if len(args) >= 2 {
			flagArgs = args[1:]
		}
	} else {
		command = args[1]
		if len(args) >= 3 {
			flagArgs = args[2:]
		}
	}

	fs := flag.NewFlagSet(command, flag.ContinueOnError)
	repo := fs.String("repo", "", "repository root path")
	legacyAuditMode := fs.String("legacy-audit-mode", "", "legacy wrapper audit mode: fail (strict-only)")
	verifyProgNewOnly := fs.String("verify-programmability-new-only", "", "enable programmability source-contract check for new objects only: true|false")
	genAutoIndent := fs.String("gen-auto-indent", "", "enable auto-indent during gen: true|false")
	patchFormat := fs.String("format-patch", "", "patch formatting policy: off|safe|strict")
	fmtCheck := fs.Bool("check", false, "fmt only: report formatting changes without writing db_object files")
	perfNotes := fs.String("perf-notes", "", "enable performance notes in out/output.json: true|false")
	failFast := fs.String("fail-fast", "", "stop verify flow after first fail item: true|false")
	parallelRead := fs.String("parallel-read", "", "enable parallel-safe file reads for gen collect stage: true|false")
	parallelWorkers := fs.String("parallel-workers", "", "number of workers for parallel-read mode: integer >= 1")
	pauseOnExit := fs.String("pause-on-exit", "", "pause before process exit (auto true on double-click): true|false")
	autoBatch := fs.String("auto-batch", "", "auto-repeat gen while pending_pairs > 0 in verify_then_gen flow: true|false")
	licenseOwner := fs.String("license-owner", "", "license owner for `license_gen` command")
	licenseMonths := fs.String("license-months", "", "license validity in months for `license_gen` command")
	licenseOut := fs.String("license-out", "", "license output file path for `license_gen` command")
	fs.SetOutput(os.Stderr)
	if err := fs.Parse(flagArgs); err != nil {
		return 1
	}

	mode, err := resolveLegacyAuditMode(*legacyAuditMode, os.Getenv("RZA_LEGACY_AUDIT_MODE"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}
	verifyProgNewOnlyEnabled, err := resolveVerifyProgrammabilityNewOnly(*verifyProgNewOnly, os.Getenv("RZA_VERIFY_PROGRAMMABILITY_NEW_ONLY"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}
	autoIndentEnabled, err := resolveGenAutoIndent(*genAutoIndent, os.Getenv("RZA_GEN_AUTO_INDENT"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}
	perfNotesEnabled, err := resolvePerfNotes(*perfNotes, os.Getenv("RZA_PERF_NOTES"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}
	failFastEnabled, err := resolveFailFast(*failFast, os.Getenv("RZA_FAIL_FAST"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}
	parallelReadEnabled, err := resolveParallelRead(*parallelRead, os.Getenv("RZA_PARALLEL_READ"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}
	parallelWorkersValue, err := resolveParallelWorkers(*parallelWorkers, os.Getenv("RZA_PARALLEL_WORKERS"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}
	pauseEnabled, err := resolvePauseOnExit(*pauseOnExit, os.Getenv("RZA_PAUSE_ON_EXIT"), implicitDefaultMode && isInteractiveConsole())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}
	autoBatchEnabled, err := resolveAutoBatch(*autoBatch, os.Getenv("RZA_AUTO_BATCH"), implicitDefaultMode || command == "bootstrap")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}

	resolvedRepo := config.DiscoverRepoRoot(*repo, args[0])
	cfg := config.Default(resolvedRepo)
	cfg.LegacyAuditMode = mode
	cfg.EnableProgrammabilitySourceContractForNewOnly = verifyProgNewOnlyEnabled
	cfg.EnableGenAutoIndent = autoIndentEnabled
	patchFormatMode, err := resolvePatchFormatMode(*patchFormat, os.Getenv("RZA_FORMAT_PATCH"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}
	cfg.PatchFormatMode = patchFormatMode
	cfg.BootstrapMode = command == "bootstrap"
	cfg.DraftMode = command == "draft"
	cfg.EnableFmtCheck = *fmtCheck
	cfg.EnablePerfNotes = perfNotesEnabled
	// Build identity and selected safety policy are part of every machine-readable
	// report. Performance notes remain controlled by EnablePerfNotes.
	cfg.EnableOutputNotes = true
	cfg.EnableFailFast = failFastEnabled
	cfg.EnableParallelRead = parallelReadEnabled
	cfg.ParallelWorkers = parallelWorkersValue

	if command != "license_gen" && cfg.EnableLicence {
		if err := ensureLicenceValid(cfg, startedAt); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return 1
		}
	}
	if *fmtCheck && command != "fmt" {
		fmt.Fprintln(os.Stderr, "error: -check is only valid with the fmt command")
		return 1
	}
	eng := engine.New(cfg)

	build := buildinfo.Current()
	fmt.Printf("[start] command=%s repo=%s at=%s version=%s commit=%s build_time=%s\n", command, filepath.Clean(cfg.RepoRoot), startedAt.Format(time.RFC3339), build.Version, build.Commit, build.BuildTime)
	fmt.Printf("[policy] existing_table_match=%s\n", cfg.ExistingTableMatchMode)
	fmt.Printf("[policy] patch_format=%s\n", cfg.PatchFormatMode)
	if autoBatchEnabled && (command == "verify_then_gen" || command == "bootstrap") {
		fmt.Printf("[mode] auto-batch enabled (will continue `gen` until pending_pairs=0)\n")
	}
	if pauseEnabled {
		fmt.Printf("[mode] pause-on-exit enabled (press Enter before close)\n")
	}

	exitCode := 1

	switch command {
	case "fmt", "verify", "gen", "verify_then_gen", "bootstrap", "draft", "seal":
		exitCode = runCommandWithProgress(eng, command, cfg.OutFilePath(), autoBatchEnabled)
	case "license_gen":
		exitCode = runLicenseGen(cfg, *licenseOwner, *licenseMonths, *licenseOut, startedAt)
	default:
		printUsage()
		exitCode = 1
	}

	fmt.Printf("[done] command=%s exit_code=%d duration=%s\n", command, exitCode, time.Since(startedAt).Round(time.Millisecond))
	maybePauseOnExit(pauseEnabled)
	return exitCode
}

func runCommandWithProgress(eng engine.Engine, command, outPath string, autoBatch bool) int {
	currentCommand := command
	iteration := 1
	lastPending := -1
	maxIterations := 10000

	for {
		fmt.Printf("[run] iteration=%d command=%s\n", iteration, currentCommand)
		report, err := eng.Run(currentCommand)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[error] iteration=%d command=%s err=%v\n", iteration, currentCommand, err)
			fmt.Fprintf(os.Stderr, "[hint] see output: %s\n", filepath.Clean(outPath))
			return 1
		}

		failCount := report.Fail.Count
		passCount := report.Pass.Count
		newFiles := len(report.NewMigrations.Files)
		pending := noteAsInt(report.Notes, "pending_pairs")
		generatedPairs := noteAsInt(report.Notes, "generated_pairs")
		fmt.Printf("[result] iteration=%d fail=%d pass=%d generated_pairs=%d up_files=%d pending_pairs=%d output=%s\n",
			iteration, failCount, passCount, generatedPairs, newFiles, pending, filepath.Clean(outPath))
		printActionSummary(report)

		if failCount > 0 {
			return 1
		}

		if !autoBatch || (command != "verify_then_gen" && command != "bootstrap") || pending <= 0 {
			return 0
		}
		if generatedPairs <= 0 {
			fmt.Fprintf(os.Stderr, "[error] pending_pairs=%d but generated_pairs=0, abort auto-batch to avoid infinite loop\n", pending)
			return 1
		}
		if lastPending >= 0 && pending >= lastPending {
			fmt.Fprintf(os.Stderr, "[error] pending_pairs did not decrease (prev=%d now=%d), abort auto-batch\n", lastPending, pending)
			return 1
		}
		if iteration >= maxIterations {
			fmt.Fprintf(os.Stderr, "[error] auto-batch reached max iterations=%d\n", maxIterations)
			return 1
		}

		lastPending = pending
		currentCommand = "gen"
		iteration++
		fmt.Printf("[batch] continue with command=gen (remaining pending_pairs=%d)\n", pending)
	}
}

func printUsage() {
	fmt.Fprintln(os.Stderr, "Usage: rza_migrate_db [fmt|verify|gen|verify_then_gen|bootstrap|draft|seal|license_gen] [-repo <path>] [-check] [-legacy-audit-mode fail] [-verify-programmability-new-only true|false] [-gen-auto-indent true|false] [-format-patch off|safe|strict] [-perf-notes true|false] [-fail-fast true|false] [-parallel-read true|false] [-parallel-workers <n>] [-pause-on-exit true|false] [-auto-batch true|false]")
	fmt.Fprintln(os.Stderr, "license_gen extra flags: [-license-owner <owner>] [-license-months <n>] [-license-out <path>]")
	fmt.Fprintln(os.Stderr, "Default (no command): verify then gen")
	fmt.Fprintln(os.Stderr, "Env fallback: RZA_LEGACY_AUDIT_MODE=fail")
	fmt.Fprintln(os.Stderr, "Env fallback: RZA_VERIFY_PROGRAMMABILITY_NEW_ONLY=true|false")
	fmt.Fprintln(os.Stderr, "Env fallback: RZA_GEN_AUTO_INDENT=true|false")
	fmt.Fprintln(os.Stderr, "Env fallback: RZA_FORMAT_PATCH=off|safe|strict")
	fmt.Fprintln(os.Stderr, "Env fallback: RZA_PERF_NOTES=true|false")
	fmt.Fprintln(os.Stderr, "Env fallback: RZA_FAIL_FAST=true|false")
	fmt.Fprintln(os.Stderr, "Env fallback: RZA_PARALLEL_READ=true|false")
	fmt.Fprintln(os.Stderr, "Env fallback: RZA_PARALLEL_WORKERS=<n>")
	fmt.Fprintln(os.Stderr, "Env fallback: RZA_PAUSE_ON_EXIT=true|false")
	fmt.Fprintln(os.Stderr, "Env fallback: RZA_AUTO_BATCH=true|false")
}

func resolvePatchFormatMode(flagValue, envValue string) (string, error) {
	value := strings.ToLower(strings.TrimSpace(flagValue))
	if value == "" {
		value = strings.ToLower(strings.TrimSpace(envValue))
	}
	if value == "" {
		return "off", nil
	}
	switch value {
	case "off", "safe", "strict":
		return value, nil
	default:
		return "", fmt.Errorf("invalid format-patch value %q (allowed: off|safe|strict)", value)
	}
}

func printActionSummary(report output.Report) {
	// Use the same stream as action summaries so merged CI/terminal output keeps
	// each diagnostic intact and in deterministic order.
	output.PrintFailureDetails(os.Stdout, report)
	const maxItems = 5
	warnings := make([]output.Item, 0)
	for _, item := range report.Pass.Items {
		if item.Severity >= 2 || strings.EqualFold(item.Action, "warn") {
			warnings = append(warnings, item)
		}
	}
	for i, item := range warnings {
		if i >= maxItems {
			fmt.Printf("[warn] ... %d additional warning(s); see output.json\n", len(warnings)-maxItems)
			break
		}
		fmt.Printf("[warn] rule=%s file=%s message=%s\n", item.Rule, item.File, item.Message)
	}
}

func ensureLicenceValid(cfg config.AppConfig, now time.Time) error {
	licPath := cfg.ResolveLicencePath()
	info, err := license.ValidateFile(licPath, now)
	if err == nil {
		fmt.Printf("[license] valid owner=%s expires_at=%s\n", info.Owner, info.ExpiresAt.Format("2006-01-02"))
		return nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("license file not found: %s (application blocked)", filepath.Clean(licPath))
	}
	if errors.Is(err, license.ErrExpired) {
		return fmt.Errorf("license expired at %s (please renew license)", info.ExpiresAt.Format("2006-01-02"))
	}
	if errors.Is(err, license.ErrInvalidFormat) || errors.Is(err, license.ErrInvalidSignature) {
		return fmt.Errorf("license invalid at %s (please regenerate license)", filepath.Clean(licPath))
	}
	return fmt.Errorf("license check failed: %w", err)
}

func runLicenseGen(cfg config.AppConfig, owner, monthsValue, outPath string, now time.Time) int {
	owner = strings.TrimSpace(owner)
	if owner == "" {
		owner = "template_repo"
	}
	months, err := resolveLicenseMonths(monthsValue)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}
	if strings.TrimSpace(outPath) == "" {
		outPath = cfg.ResolveLicencePathForWrite()
	}

	expiresAt := now.AddDate(0, months, 0)
	text, err := license.Generate(owner, now, expiresAt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: generate license failed: %v\n", err)
		return 1
	}
	if err := os.WriteFile(outPath, []byte(text), 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "error: write license failed: %v\n", err)
		return 1
	}
	fmt.Printf("[license] generated file=%s owner=%s issued_at=%s expires_at=%s\n",
		filepath.Clean(outPath), owner, now.Format("2006-01-02"), expiresAt.Format("2006-01-02"))
	return 0
}

func resolveLicenseMonths(value string) (int, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return 3, nil
	}
	months, err := strconv.Atoi(value)
	if err != nil || months < 1 {
		return 0, fmt.Errorf("invalid license-months %q (must be integer >= 1)", value)
	}
	return months, nil
}

func resolveLegacyAuditMode(flagValue, envValue string) (string, error) {
	flagMode := strings.ToLower(strings.TrimSpace(flagValue))
	envMode := strings.ToLower(strings.TrimSpace(envValue))

	mode := flagMode
	if mode == "" {
		mode = envMode
	}
	if mode == "" {
		mode = "fail"
	}

	if mode != "fail" {
		return "", fmt.Errorf("invalid legacy audit mode %q (allowed: fail only)", mode)
	}
	return mode, nil
}

func resolveGenAutoIndent(flagValue, envValue string) (bool, error) {
	value := strings.ToLower(strings.TrimSpace(flagValue))
	if value == "" {
		value = strings.ToLower(strings.TrimSpace(envValue))
	}
	if value == "" {
		return true, nil
	}
	switch value {
	case "1", "true", "on", "yes":
		return true, nil
	case "0", "false", "off", "no":
		return false, nil
	default:
		return false, fmt.Errorf("invalid gen auto-indent value %q (allowed: true|false)", value)
	}
}

func resolveVerifyProgrammabilityNewOnly(flagValue, envValue string) (bool, error) {
	value := strings.ToLower(strings.TrimSpace(flagValue))
	if value == "" {
		value = strings.ToLower(strings.TrimSpace(envValue))
	}
	if value == "" {
		return false, nil
	}
	switch value {
	case "1", "true", "on", "yes":
		return true, nil
	case "0", "false", "off", "no":
		return false, nil
	default:
		return false, fmt.Errorf("invalid verify-programmability-new-only value %q (allowed: true|false)", value)
	}
}

func resolvePerfNotes(flagValue, envValue string) (bool, error) {
	value := strings.ToLower(strings.TrimSpace(flagValue))
	if value == "" {
		value = strings.ToLower(strings.TrimSpace(envValue))
	}
	if value == "" {
		return false, nil
	}
	switch value {
	case "1", "true", "on", "yes":
		return true, nil
	case "0", "false", "off", "no":
		return false, nil
	default:
		return false, fmt.Errorf("invalid perf-notes value %q (allowed: true|false)", value)
	}
}

func resolveFailFast(flagValue, envValue string) (bool, error) {
	value := strings.ToLower(strings.TrimSpace(flagValue))
	if value == "" {
		value = strings.ToLower(strings.TrimSpace(envValue))
	}
	if value == "" {
		return false, nil
	}
	switch value {
	case "1", "true", "on", "yes":
		return true, nil
	case "0", "false", "off", "no":
		return false, nil
	default:
		return false, fmt.Errorf("invalid fail-fast value %q (allowed: true|false)", value)
	}
}

func resolveParallelRead(flagValue, envValue string) (bool, error) {
	value := strings.ToLower(strings.TrimSpace(flagValue))
	if value == "" {
		value = strings.ToLower(strings.TrimSpace(envValue))
	}
	if value == "" {
		return true, nil
	}
	switch value {
	case "1", "true", "on", "yes":
		return true, nil
	case "0", "false", "off", "no":
		return false, nil
	default:
		return false, fmt.Errorf("invalid parallel-read value %q (allowed: true|false)", value)
	}
}

func resolveParallelWorkers(flagValue, envValue string) (int, error) {
	value := strings.TrimSpace(flagValue)
	if value == "" {
		value = strings.TrimSpace(envValue)
	}
	if value == "" {
		return 8, nil
	}
	workers, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid parallel-workers value %q (must be integer >= 1)", value)
	}
	if workers < 1 {
		return 0, fmt.Errorf("invalid parallel-workers value %q (must be integer >= 1)", value)
	}
	return workers, nil
}

func resolvePauseOnExit(flagValue, envValue string, defaultValue bool) (bool, error) {
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
		return false, fmt.Errorf("invalid pause-on-exit value %q (allowed: true|false)", value)
	}
}

func resolveAutoBatch(flagValue, envValue string, defaultValue bool) (bool, error) {
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
		return false, fmt.Errorf("invalid auto-batch value %q (allowed: true|false)", value)
	}
}

func noteAsInt(notes map[string]interface{}, key string) int {
	if notes == nil {
		return 0
	}
	raw, ok := notes[key]
	if !ok || raw == nil {
		return 0
	}
	switch v := raw.(type) {
	case int:
		return v
	case int32:
		return int(v)
	case int64:
		return int(v)
	case float64:
		return int(v)
	case string:
		n, err := strconv.Atoi(strings.TrimSpace(v))
		if err == nil {
			return n
		}
	}
	return 0
}

func maybePauseOnExit(enabled bool) {
	if !enabled {
		return
	}
	fmt.Print("Press Enter to close...")
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
}

func isInteractiveConsole() bool {
	stdinInfo, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	stdoutInfo, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (stdinInfo.Mode()&os.ModeCharDevice) != 0 && (stdoutInfo.Mode()&os.ModeCharDevice) != 0
}
