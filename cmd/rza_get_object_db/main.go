package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"rza_migrate_db/internal/buildinfo"
	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/getobject"
)

var exportFromDatabase = getobject.ExportFromDatabaseWithOptions
var previewFromDir = getobject.PreviewModulesFromDir

type cliRunSummary struct {
	Timestamp     string                   `json:"timestamp"`
	Executable    string                   `json:"executable"`
	Command       string                   `json:"command"`
	Status        string                   `json:"status"`
	ExitCode      int                      `json:"exit_code"`
	Error         string                   `json:"error,omitempty"`
	DryRun        bool                     `json:"dry_run,omitempty"`
	PreviewInput  string                   `json:"preview_input,omitempty"`
	PreviewOutput string                   `json:"preview_output,omitempty"`
	Export        *cliExportSummary        `json:"export,omitempty"`
	Preview       *cliPreviewSummary       `json:"preview,omitempty"`
	Config        *cliSummaryConfigSection `json:"config,omitempty"`
	Build         buildinfo.Info           `json:"build"`
}

type cliExportSummary struct {
	Created      int                        `json:"created"`
	Updated      int                        `json:"updated"`
	Deleted      int                        `json:"deleted"`
	Unchanged    int                        `json:"unchanged"`
	Categories   []getobject.CategoryChange `json:"categories,omitempty"`
	WarningCount int                        `json:"warning_count,omitempty"`
	Warnings     []string                   `json:"warnings,omitempty"`
	WarningItems []getobject.ExportWarning  `json:"warning_items,omitempty"`
}

type cliPreviewSummary struct {
	Processed  int      `json:"processed"`
	Written    int      `json:"written"`
	Failed     int      `json:"failed"`
	FailedList []string `json:"failed_list,omitempty"`
}

type cliSummaryConfigSection struct {
	RepoRoot                  string `json:"repo_root"`
	Schema                    string `json:"schema"`
	SkipUnsupportedTables     bool   `json:"skip_unsupported_tables"`
	RequireHealthyConstraints bool   `json:"require_healthy_constraints"`
	Profile                   string `json:"profile"`
}

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	execName := ""
	if len(args) > 0 {
		execName = filepath.Base(strings.TrimSpace(args[0]))
	}
	summary := cliRunSummary{
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
		Executable: execName,
		Command:    "export",
		Status:     "fail",
		ExitCode:   1,
		Build:      buildinfo.Current(),
	}
	fmt.Printf("[start] version=%s commit=%s build_time=%s\n", summary.Build.Version, summary.Build.Commit, summary.Build.BuildTime)

	if len(args) > 0 {
		if err := config.ValidateExecutableName(args[0], config.DefaultGetObjectExecutableNameLock()); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			summary.Error = err.Error()
			return finishRun(".", summary, 1)
		}
	}

	fs := flag.NewFlagSet("rza_get_object_db", flag.ContinueOnError)
	profile := fs.String("profile", "continue", "extractor policy profile: continue (default)|strict|bootstrap")
	dryRun := fs.Bool("dry-run", false, "preview create/update/delete without writing files")
	skipUnsupportedTables := fs.Bool(
		"skip-unsupported-tables",
		false,
		"continue export with warnings when a table cannot be represented losslessly; preserve its existing snapshots",
	)
	requireHealthyConstraints := fs.Bool(
		"require-healthy-constraints",
		false,
		"return a non-zero exit after export when disabled or untrusted constraints are observed",
	)
	configPath := fs.String("config", "", "extractor YAML/JSON config (default: config.yaml in working directory, then DSN environment)")
	previewInput := fs.String("preview-input", "", "local folder containing .sql samples (no DB connection)")
	previewOutput := fs.String("preview-output", filepath.FromSlash("out/preview_modules"), "output folder for local preview formatting")
	fs.SetOutput(os.Stderr)
	if err := fs.Parse(args[1:]); err != nil {
		summary.Error = err.Error()
		return finishRun(".", summary, 1)
	}
	profileValue := strings.ToLower(strings.TrimSpace(*profile))
	if profileValue != "continue" && profileValue != "strict" && profileValue != "bootstrap" {
		summary.Error = fmt.Sprintf("invalid profile %q (allowed: continue|strict|bootstrap)", profileValue)
		fmt.Fprintln(os.Stderr, "error: "+summary.Error)
		return finishRun(".", summary, 1)
	}

	cfg, loadErr := loadGetObjectConfig(*configPath, strings.TrimSpace(*previewInput) != "")
	if loadErr != nil {
		fmt.Fprintf(os.Stderr, "error: load config: %v\n", loadErr)
		summary.Error = "load config: " + loadErr.Error()
		return finishRun(".", summary, 1)
	}
	if profileValue == "strict" {
		cfg.SkipUnsupportedTables = false
	}
	if *skipUnsupportedTables {
		cfg.SkipUnsupportedTables = true
	}
	if profileValue == "bootstrap" || profileValue == "continue" {
		cfg.SkipUnsupportedTables = true
	}
	repoRoot := cfg.RepoRoot
	summary.Config = &cliSummaryConfigSection{
		RepoRoot:                  cfg.RepoRoot,
		Schema:                    cfg.Schema,
		SkipUnsupportedTables:     cfg.SkipUnsupportedTables,
		RequireHealthyConstraints: *requireHealthyConstraints,
		Profile:                   profileValue,
	}
	if strings.TrimSpace(*previewInput) != "" {
		summary.Command = "preview"
		summary.PreviewInput = *previewInput
		summary.PreviewOutput = *previewOutput
		summary, err := previewFromDir(cfg, *previewInput, *previewOutput, func(phase, message string) {
			fmt.Fprintf(os.Stdout, "[%s] %s\n", phase, message)
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return finishRun(repoRoot, cliRunSummary{
				Timestamp:     time.Now().UTC().Format(time.RFC3339),
				Executable:    execName,
				Command:       "preview",
				Status:        "fail",
				ExitCode:      1,
				Error:         err.Error(),
				PreviewInput:  *previewInput,
				PreviewOutput: *previewOutput,
				Preview: &cliPreviewSummary{
					Processed:  summary.Processed,
					Written:    summary.Written,
					Failed:     summary.Failed,
					FailedList: summary.FailedList,
				},
				Config: &cliSummaryConfigSection{
					RepoRoot:                  cfg.RepoRoot,
					Schema:                    cfg.Schema,
					SkipUnsupportedTables:     cfg.SkipUnsupportedTables,
					RequireHealthyConstraints: *requireHealthyConstraints,
					Profile:                   profileValue,
				},
			}, 1)
		}
		fmt.Fprintf(os.Stdout, "preview completed: processed=%d written=%d failed=%d output=%s\n", summary.Processed, summary.Written, summary.Failed, *previewOutput)
		if summary.Failed > 0 {
			fmt.Fprintf(os.Stdout, "preview failed files:\n")
			for _, f := range summary.FailedList {
				fmt.Fprintf(os.Stdout, "- %s\n", f)
			}
			return finishRun(repoRoot, cliRunSummary{
				Timestamp:     time.Now().UTC().Format(time.RFC3339),
				Executable:    execName,
				Command:       "preview",
				Status:        "fail",
				ExitCode:      1,
				Error:         "preview contains failed files",
				PreviewInput:  *previewInput,
				PreviewOutput: *previewOutput,
				Preview: &cliPreviewSummary{
					Processed:  summary.Processed,
					Written:    summary.Written,
					Failed:     summary.Failed,
					FailedList: summary.FailedList,
				},
				Config: &cliSummaryConfigSection{
					RepoRoot: cfg.RepoRoot,
					Schema:   cfg.Schema,
				},
			}, 1)
		}
		return finishRun(repoRoot, cliRunSummary{
			Timestamp:     time.Now().UTC().Format(time.RFC3339),
			Executable:    execName,
			Command:       "preview",
			Status:        "ok",
			ExitCode:      0,
			PreviewInput:  *previewInput,
			PreviewOutput: *previewOutput,
			Preview: &cliPreviewSummary{
				Processed:  summary.Processed,
				Written:    summary.Written,
				Failed:     summary.Failed,
				FailedList: summary.FailedList,
			},
			Config: &cliSummaryConfigSection{
				RepoRoot:                  cfg.RepoRoot,
				Schema:                    cfg.Schema,
				SkipUnsupportedTables:     cfg.SkipUnsupportedTables,
				RequireHealthyConstraints: *requireHealthyConstraints,
				Profile:                   profileValue,
			},
		}, 0)
	}

	exportSummary, err := exportFromDatabase(cfg, getobject.ExportOptions{
		DryRun: *dryRun,
		Progress: func(phase, message string) {
			fmt.Fprintf(os.Stdout, "[%s] %s\n", phase, message)
		},
	})
	if err != nil {
		errorMessage := err.Error()
		if !cfg.SkipUnsupportedTables && getobject.IsUnsupportedTableBuildError(err) {
			errorMessage += "; rerun with -skip-unsupported-tables to continue safely while preserving existing snapshots for unsupported tables"
		}
		fmt.Fprintf(os.Stderr, "error: %s\n", errorMessage)
		return finishRun(repoRoot, cliRunSummary{
			Timestamp:  time.Now().UTC().Format(time.RFC3339),
			Executable: execName,
			Command:    "export",
			Status:     "fail",
			ExitCode:   1,
			Error:      errorMessage,
			DryRun:     *dryRun,
			Config: &cliSummaryConfigSection{
				RepoRoot:                  cfg.RepoRoot,
				Schema:                    cfg.Schema,
				SkipUnsupportedTables:     cfg.SkipUnsupportedTables,
				RequireHealthyConstraints: *requireHealthyConstraints,
				Profile:                   profileValue,
			},
		}, 1)
	}

	mode := "export"
	if exportSummary.DryRun {
		mode = "dry-run"
	}
	fmt.Fprintf(
		os.Stdout,
		"%s completed: create=%d update=%d delete=%d unchanged=%d\n",
		mode,
		exportSummary.TotalCreated(),
		exportSummary.TotalUpdated(),
		exportSummary.TotalDeleted(),
		exportSummary.TotalUnchanged(),
	)
	warningMessages := exportSummary.WarningMessages()
	if len(warningMessages) > 0 {
		fmt.Fprintf(os.Stdout, "warnings: %d\n", len(warningMessages))
		for _, warning := range warningMessages {
			fmt.Fprintf(os.Stdout, "- %s\n", warning)
		}
	}
	exportPayload := cliExportSummaryFrom(exportSummary)
	resultSummary := cliRunSummary{
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
		Executable: execName,
		Command:    mode,
		Status:     "ok",
		ExitCode:   0,
		DryRun:     exportSummary.DryRun,
		Export:     &exportPayload,
		Config: &cliSummaryConfigSection{
			RepoRoot:                  cfg.RepoRoot,
			Schema:                    cfg.Schema,
			SkipUnsupportedTables:     cfg.SkipUnsupportedTables,
			RequireHealthyConstraints: *requireHealthyConstraints,
			Profile:                   profileValue,
		},
	}
	if *requireHealthyConstraints && exportSummary.ConstraintHealthFindingCount() > 0 {
		resultSummary.Error = fmt.Sprintf(
			"constraint health gate failed: %d disabled or untrusted constraint(s) observed; export evidence was recorded and the database was not modified",
			exportSummary.ConstraintHealthFindingCount(),
		)
		fmt.Fprintf(os.Stderr, "error: %s\n", resultSummary.Error)
		return finishRun(repoRoot, resultSummary, 1)
	}
	return finishRun(repoRoot, resultSummary, 0)
}

func cliExportSummaryFrom(summary getobject.ExportSummary) cliExportSummary {
	warnings := summary.WarningMessages()
	return cliExportSummary{
		Created:      summary.TotalCreated(),
		Updated:      summary.TotalUpdated(),
		Deleted:      summary.TotalDeleted(),
		Unchanged:    summary.TotalUnchanged(),
		Categories:   append([]getobject.CategoryChange(nil), summary.Categories...),
		WarningCount: len(warnings),
		Warnings:     append([]string(nil), warnings...),
		WarningItems: append(append([]getobject.ExportWarning(nil), summary.TableWarnings...), summary.ObjectWarnings...),
	}
}

func finishRun(repoRoot string, summary cliRunSummary, exitCode int) int {
	summary.Build = buildinfo.Current()
	summary.Timestamp = time.Now().UTC().Format(time.RFC3339)
	summary.ExitCode = exitCode
	if exitCode == 0 {
		summary.Status = "ok"
	} else {
		summary.Status = "fail"
	}

	path, err := writeRunSummary(repoRoot, summary)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: write extraction reports failed: %v\n", err)
		return 1
	}
	fmt.Fprintf(os.Stdout, "summary written: %s\n", filepath.ToSlash(path))
	fmt.Fprintf(os.Stdout, "object issues written: %s\n", filepath.ToSlash(filepath.Join(filepath.Dir(path), "get_object_issues.txt")))
	return exitCode
}

func writeRunSummary(repoRoot string, summary cliRunSummary) (string, error) {
	root := strings.TrimSpace(repoRoot)
	if root == "" {
		root = "."
	}
	summaryPath := filepath.Join(root, "out", "get_object_summary.json")
	if err := os.MkdirAll(filepath.Dir(summaryPath), 0o755); err != nil {
		return "", err
	}
	raw, err := json.MarshalIndent(summary, "", "  ")
	if err != nil {
		return "", err
	}
	if err := os.WriteFile(summaryPath, append(raw, '\n'), 0o644); err != nil {
		return "", err
	}
	issuesPath := filepath.Join(root, "out", "get_object_issues.txt")
	var issues strings.Builder
	issues.WriteString("Extractor object issues\n")
	if summary.Error != "" {
		fmt.Fprintf(&issues, "Run error: %s\n", summary.Error)
	}
	if summary.Export != nil {
		fmt.Fprintf(&issues, "Warnings: %d\n", summary.Export.WarningCount)
		for _, warning := range summary.Export.Warnings {
			fmt.Fprintf(&issues, "- %s\n", warning)
		}
	}
	if summary.Error == "" && (summary.Export == nil || summary.Export.WarningCount == 0) {
		issues.WriteString("No object issues reported.\n")
	}
	if err := os.WriteFile(issuesPath, []byte(issues.String()), 0o644); err != nil {
		return "", err
	}
	return summaryPath, nil
}
