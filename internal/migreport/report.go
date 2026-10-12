package migreport

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Options holds runtime options for rza_migrate_report.
type Options struct {
	RepoRoot    string
	OutputDir   string
	DocsDir     string
	PauseOnExit bool
}

// Run executes the complete migration history reconciliation report workflow.
func Run(opts Options) (int, error) {
	repoRoot, err := DiscoverRepository(opts.RepoRoot)
	if err != nil {
		target := opts.RepoRoot
		if strings.TrimSpace(target) == "" {
			if cwd, cwdErr := os.Getwd(); cwdErr == nil && cwd != "" {
				target = cwd
			} else {
				target = "."
			}
		}
		fmt.Fprintf(os.Stderr, "[error] migrations directory not found: %s\n", target)
		fmt.Fprintf(os.Stderr, "[suggestion] place rza_migrate_report.exe beside the repository or run it from repository root\n")
		WaitForEnter(opts.PauseOnExit)
		return 1, err
	}
	repoRoot, err = filepath.Abs(repoRoot)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[error] failed resolving repository root: %v\n", err)
		WaitForEnter(opts.PauseOnExit)
		return 1, err
	}
	repoRoot = filepath.Clean(repoRoot)

	outDir := resolveReportDirectory(repoRoot, opts.OutputDir, "out", "report")
	if err := validateDirectoryOutsideProtected(repoRoot, outDir, "report output"); err != nil {
		fmt.Fprintf(os.Stderr, "[error] invalid report output directory: %v\n", err)
		WaitForEnter(opts.PauseOnExit)
		return 1, err
	}

	docsDir := resolveReportDirectory(repoRoot, opts.DocsDir, "docs", "database")
	if err := validateDirectoryOutsideProtected(repoRoot, docsDir, "documentation output"); err != nil {
		fmt.Fprintf(os.Stderr, "[error] invalid documentation output directory: %v\n", err)
		WaitForEnter(opts.PauseOnExit)
		return 1, err
	}

	pairs, allFiles, warnings, errors, scanErr := ScanMigrations(repoRoot)
	if scanErr != nil {
		fmt.Fprintf(os.Stderr, "[error] failed scanning migrations: %v\n", scanErr)
		WaitForEnter(opts.PauseOnExit)
		return 1, scanErr
	}

	// The canonical core audit bootstrap is migration infrastructure rather than
	// an application-schema change. Keep scanning it as a file, but do not let
	// it add mig_event/audit rows to the business migration history report.
	var excludedPairs int
	pairs, excludedPairs = excludeInitialCoreMigration(pairs)

	result, analyzeErr := Analyze(repoRoot, pairs, allFiles, warnings, errors)
	if analyzeErr != nil {
		fmt.Fprintf(os.Stderr, "[error] analysis failed: %v\n", analyzeErr)
		WaitForEnter(opts.PauseOnExit)
		return 1, analyzeErr
	}
	result.Summary.ExcludedPairs = excludedPairs
	result.JSONReport.ExcludedPairs = excludedPairs
	desiredStateDocs, docsErr := BuildDesiredStateDocs(repoRoot)
	if docsErr != nil {
		fmt.Fprintf(os.Stderr, "[error] desired-state documentation analysis failed: %v\n", docsErr)
		WaitForEnter(opts.PauseOnExit)
		return 1, docsErr
	}
	result.DesiredStateDocs = desiredStateDocs
	result.JSONReport.Reports = append(result.JSONReport.Reports, desiredStateReportFiles...)
	result.JSONReport.OutputDirectory = outDir
	result.JSONReport.DocumentationDirectory = docsDir
	result.JSONReport.Artifacts = reportArtifacts(result.DesiredStateDocs)

	if writeErr := WriteDesiredStateDocs(docsDir, result.DesiredStateDocs); writeErr != nil {
		fmt.Fprintf(os.Stderr, "[error] failed writing desired-state documentation: %v\n", writeErr)
		WaitForEnter(opts.PauseOnExit)
		return 1, writeErr
	}
	if writeErr := WriteReports(outDir, result); writeErr != nil {
		fmt.Fprintf(os.Stderr, "[error] failed writing reports: %v\n", writeErr)
		WaitForEnter(opts.PauseOnExit)
		return 1, writeErr
	}

	PrintConsoleSummary(result, outDir, docsDir)
	WaitForEnter(opts.PauseOnExit)

	if result.Summary.Status == StatusFailed {
		return 1, nil
	}
	return 0, nil
}

func resolveReportDirectory(repoRoot, configured string, defaultParts ...string) string {
	configured = strings.TrimSpace(configured)
	if configured == "" {
		parts := append([]string{repoRoot}, defaultParts...)
		return filepath.Clean(filepath.Join(parts...))
	}
	if filepath.IsAbs(configured) {
		return filepath.Clean(configured)
	}
	return filepath.Clean(filepath.Join(repoRoot, configured))
}

func reportArtifacts(docs DesiredStateDocs) []ReportArtifact {
	reportFiles := []string{
		"report_summary.csv", "migration_summary.csv", "object_history.csv", "table_changes.csv",
		"index_changes.csv", "programmable_changes.csv", "type_changes.csv", "report.json",
	}
	artifacts := make([]ReportArtifact, 0, len(reportFiles)+len(docs.ProcedureProfiles)+7)
	for _, path := range reportFiles {
		artifacts = append(artifacts, ReportArtifact{Kind: "audit", Root: "output_directory", Path: path})
	}
	for _, path := range desiredStateArtifactFiles(docs) {
		artifacts = append(artifacts, ReportArtifact{Kind: "documentation", Root: "documentation_directory", Path: filepath.ToSlash(path)})
	}
	return artifacts
}

func excludeInitialCoreMigration(pairs []MigrationPair) ([]MigrationPair, int) {
	filtered := make([]MigrationPair, 0, len(pairs))
	excluded := 0
	for _, pair := range pairs {
		if pair.Category == "core" && pair.MigrationID == "000000" && pair.ObjectName == "mig__audit_init" {
			excluded++
			continue
		}
		filtered = append(filtered, pair)
	}
	return filtered, excluded
}

func validateOutputDirectory(repoRoot, outputDir string) error {
	return validateDirectoryOutsideProtected(repoRoot, outputDir, "output")
}

func validateDirectoryOutsideProtected(repoRoot, targetDir, dirLabel string) error {
	migrationsDir, err := resolvePathForContainment(filepath.Join(repoRoot, "migrations"))
	if err != nil {
		return fmt.Errorf("resolve migrations directory: %w", err)
	}
	resolvedTarget, err := resolvePathForContainment(targetDir)
	if err != nil {
		return fmt.Errorf("resolve %s directory: %w", dirLabel, err)
	}
	relMig, err := filepath.Rel(migrationsDir, resolvedTarget)
	if err != nil {
		return fmt.Errorf("compare %s and migrations directories: %w", dirLabel, err)
	}
	if relMig == "." || (relMig != ".." && !strings.HasPrefix(relMig, ".."+string(filepath.Separator))) {
		return fmt.Errorf("%s directory must be outside migrations/: %s", dirLabel, targetDir)
	}

	dbObjDir, statErr := resolvePathForContainment(filepath.Join(repoRoot, "db_object"))
	if statErr == nil {
		relObj, err := filepath.Rel(dbObjDir, resolvedTarget)
		if err == nil {
			if relObj == "." || (relObj != ".." && !strings.HasPrefix(relObj, ".."+string(filepath.Separator))) {
				return fmt.Errorf("%s directory must be outside db_object/: %s", dirLabel, targetDir)
			}
		}
	}
	return nil
}

func resolvePathForContainment(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	existing := absPath
	var suffix []string
	for {
		_, statErr := os.Lstat(existing)
		if statErr == nil {
			break
		}
		if !os.IsNotExist(statErr) {
			return "", statErr
		}
		parent := filepath.Dir(existing)
		if parent == existing {
			return "", statErr
		}
		suffix = append([]string{filepath.Base(existing)}, suffix...)
		existing = parent
	}

	resolved, err := filepath.EvalSymlinks(existing)
	if err != nil {
		return "", err
	}
	for _, part := range suffix {
		resolved = filepath.Join(resolved, part)
	}
	return filepath.Clean(resolved), nil
}
