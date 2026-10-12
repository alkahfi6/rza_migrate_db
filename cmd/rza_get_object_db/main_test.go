package main

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rza_migrate_db/internal/getobject"
)

func TestRunFailsWhenExecutableRenamed(t *testing.T) {
	exit := run([]string{
		"renamed_get_object_tool.exe",
	})
	if exit != 1 {
		t.Fatalf("expected renamed executable to fail, got exit %d", exit)
	}
}

func TestRunDryRunFlagAccepted(t *testing.T) {
	tmpRepo := t.TempDir()
	prevExport := exportFromDatabase
	prevPreview := previewFromDir
	prevCfg := compiledGetObjectConfig
	defer func() {
		exportFromDatabase = prevExport
		previewFromDir = prevPreview
		compiledGetObjectConfig = prevCfg
	}()

	compiledGetObjectConfig = getobject.AppConfig{
		ConnectionString:      "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable",
		RepoRoot:              tmpRepo,
		DbObjectDir:           "db_object",
		Schema:                "dbo",
		ConnectTimeoutSeconds: 10,
		QueryTimeoutSeconds:   20,
	}
	exportFromDatabase = func(cfg getobject.AppConfig, opts getobject.ExportOptions) (getobject.ExportSummary, error) {
		if !opts.DryRun {
			t.Fatalf("expected dry-run option true")
		}
		if opts.Progress == nil {
			t.Fatalf("expected progress callback to be set")
		}
		return getobject.ExportSummary{
			DryRun: true,
		}, nil
	}

	exit := run([]string{
		"rza_get_object_db",
		"-dry-run",
	})
	if exit != 0 {
		t.Fatalf("expected dry-run exit 0, got %d", exit)
	}

	summaryPath := filepath.Join(tmpRepo, "out", "get_object_summary.json")
	raw, err := os.ReadFile(summaryPath)
	if err != nil {
		t.Fatalf("read summary: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("parse summary json: %v", err)
	}
	if payload["status"] != "ok" {
		t.Fatalf("expected summary status=ok, got %v", payload["status"])
	}
}

func TestRunSkipUnsupportedTablesFlagOverridesCompiledConfig(t *testing.T) {
	tmpRepo := t.TempDir()
	prevExport := exportFromDatabase
	prevPreview := previewFromDir
	prevCfg := compiledGetObjectConfig
	defer func() {
		exportFromDatabase = prevExport
		previewFromDir = prevPreview
		compiledGetObjectConfig = prevCfg
	}()

	compiledGetObjectConfig = getobject.AppConfig{
		ConnectionString:      "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable",
		RepoRoot:              tmpRepo,
		DbObjectDir:           "db_object",
		Schema:                "dbo",
		ConnectTimeoutSeconds: 10,
		QueryTimeoutSeconds:   20,
		SkipUnsupportedTables: false,
	}
	exportFromDatabase = func(cfg getobject.AppConfig, opts getobject.ExportOptions) (getobject.ExportSummary, error) {
		if !cfg.SkipUnsupportedTables {
			t.Fatal("expected -skip-unsupported-tables to enable continuation mode")
		}
		return getobject.ExportSummary{}, nil
	}

	exit := run([]string{
		"rza_get_object_db",
		"-skip-unsupported-tables",
	})
	if exit != 0 {
		t.Fatalf("expected export exit 0, got %d", exit)
	}

	summaryPath := filepath.Join(tmpRepo, "out", "get_object_summary.json")
	raw, err := os.ReadFile(summaryPath)
	if err != nil {
		t.Fatalf("read summary: %v", err)
	}
	var payload struct {
		Config struct {
			SkipUnsupportedTables bool `json:"skip_unsupported_tables"`
		} `json:"config"`
	}
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("parse summary json: %v", err)
	}
	if !payload.Config.SkipUnsupportedTables {
		t.Fatal("expected summary to record skip_unsupported_tables=true")
	}
}

func TestRunFailsWhenExportReturnsError(t *testing.T) {
	tmpRepo := t.TempDir()
	prevExport := exportFromDatabase
	prevPreview := previewFromDir
	prevCfg := compiledGetObjectConfig
	defer func() {
		exportFromDatabase = prevExport
		previewFromDir = prevPreview
		compiledGetObjectConfig = prevCfg
	}()

	compiledGetObjectConfig = getobject.AppConfig{
		ConnectionString:      "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable",
		RepoRoot:              tmpRepo,
		DbObjectDir:           "db_object",
		Schema:                "dbo",
		ConnectTimeoutSeconds: 10,
		QueryTimeoutSeconds:   20,
	}
	exportFromDatabase = func(cfg getobject.AppConfig, opts getobject.ExportOptions) (getobject.ExportSummary, error) {
		return getobject.ExportSummary{}, errors.New("forced error")
	}

	exit := run([]string{
		"rza_get_object_db",
	})
	if exit != 1 {
		t.Fatalf("expected exit 1 when export returns error, got %d", exit)
	}

	summaryPath := filepath.Join(tmpRepo, "out", "get_object_summary.json")
	raw, err := os.ReadFile(summaryPath)
	if err != nil {
		t.Fatalf("read summary: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("parse summary json: %v", err)
	}
	if payload["status"] != "fail" {
		t.Fatalf("expected summary status=fail, got %v", payload["status"])
	}
}

func TestRunUnsupportedTableErrorIncludesContinuationHint(t *testing.T) {
	tmpRepo := t.TempDir()
	prevExport := exportFromDatabase
	prevPreview := previewFromDir
	prevCfg := compiledGetObjectConfig
	defer func() {
		exportFromDatabase = prevExport
		previewFromDir = prevPreview
		compiledGetObjectConfig = prevCfg
	}()

	compiledGetObjectConfig = getobject.AppConfig{
		ConnectionString:      "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable",
		RepoRoot:              tmpRepo,
		DbObjectDir:           "db_object",
		Schema:                "dbo",
		ConnectTimeoutSeconds: 10,
		QueryTimeoutSeconds:   20,
	}
	exportFromDatabase = func(cfg getobject.AppConfig, opts getobject.ExportOptions) (getobject.ExportSummary, error) {
		return getobject.ExportSummary{}, errors.New("build table snapshot dbo.Child: foreign key FK_Child is disabled or untrusted and cannot be normalized safely")
	}

	exit := run([]string{"rza_get_object_db"})
	if exit != 1 {
		t.Fatalf("expected exit 1, got %d", exit)
	}

	raw, err := os.ReadFile(filepath.Join(tmpRepo, "out", "get_object_summary.json"))
	if err != nil {
		t.Fatalf("read summary: %v", err)
	}
	var payload struct {
		Error string `json:"error"`
	}
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("parse summary json: %v", err)
	}
	if !strings.Contains(payload.Error, "-skip-unsupported-tables") {
		t.Fatalf("expected actionable continuation hint, got %q", payload.Error)
	}
}

func TestRunPreviewModeAccepted(t *testing.T) {
	tmpRepo := t.TempDir()
	prevExport := exportFromDatabase
	prevPreview := previewFromDir
	prevCfg := compiledGetObjectConfig
	defer func() {
		exportFromDatabase = prevExport
		previewFromDir = prevPreview
		compiledGetObjectConfig = prevCfg
	}()

	compiledGetObjectConfig = getobject.AppConfig{
		ConnectionString:      "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable",
		RepoRoot:              tmpRepo,
		DbObjectDir:           "db_object",
		Schema:                "dbo",
		ConnectTimeoutSeconds: 10,
		QueryTimeoutSeconds:   20,
	}
	exportFromDatabase = func(cfg getobject.AppConfig, opts getobject.ExportOptions) (getobject.ExportSummary, error) {
		t.Fatalf("export must not be called in preview mode")
		return getobject.ExportSummary{}, nil
	}
	previewFromDir = func(cfg getobject.AppConfig, inputDir, outputDir string, progress func(phase, message string)) (getobject.PreviewSummary, error) {
		if inputDir != "test-modules" {
			t.Fatalf("unexpected preview input: %s", inputDir)
		}
		if outputDir != "out/preview_modules" {
			t.Fatalf("unexpected preview output: %s", outputDir)
		}
		if progress == nil {
			t.Fatalf("progress callback should not be nil")
		}
		return getobject.PreviewSummary{
			Processed: 1,
			Written:   1,
		}, nil
	}

	exit := run([]string{
		"rza_get_object_db",
		"-preview-input",
		"test-modules",
		"-preview-output",
		"out/preview_modules",
	})
	if exit != 0 {
		t.Fatalf("expected preview mode exit 0, got %d", exit)
	}

	summaryPath := filepath.Join(tmpRepo, "out", "get_object_summary.json")
	raw, err := os.ReadFile(summaryPath)
	if err != nil {
		t.Fatalf("read summary: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("parse summary json: %v", err)
	}
	if payload["command"] != "preview" {
		t.Fatalf("expected command=preview, got %v", payload["command"])
	}
}

func TestRunExportWarningsPersisted(t *testing.T) {
	tmpRepo := t.TempDir()
	prevExport := exportFromDatabase
	prevPreview := previewFromDir
	prevCfg := compiledGetObjectConfig
	defer func() {
		exportFromDatabase = prevExport
		previewFromDir = prevPreview
		compiledGetObjectConfig = prevCfg
	}()

	compiledGetObjectConfig = getobject.AppConfig{
		ConnectionString:      "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable",
		RepoRoot:              tmpRepo,
		DbObjectDir:           "db_object",
		Schema:                "dbo",
		ConnectTimeoutSeconds: 10,
		QueryTimeoutSeconds:   20,
	}
	exportFromDatabase = func(cfg getobject.AppConfig, opts getobject.ExportOptions) (getobject.ExportSummary, error) {
		return getobject.ExportSummary{
			ModuleWarnings: []string{
				"procedure dbo.ProcA skipped: definition is empty or encrypted",
				"view dbo.ViewA skipped: invalid object name",
			},
		}, nil
	}

	exit := run([]string{
		"rza_get_object_db",
	})
	if exit != 0 {
		t.Fatalf("expected exit 0 with module warnings, got %d", exit)
	}

	summaryPath := filepath.Join(tmpRepo, "out", "get_object_summary.json")
	raw, err := os.ReadFile(summaryPath)
	if err != nil {
		t.Fatalf("read summary: %v", err)
	}

	var payload struct {
		Export struct {
			WarningCount float64  `json:"warning_count"`
			Warnings     []string `json:"warnings"`
		} `json:"export"`
	}
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("parse summary json: %v", err)
	}
	if payload.Export.WarningCount != 2 {
		t.Fatalf("expected warning_count=2, got %v", payload.Export.WarningCount)
	}
	if len(payload.Export.Warnings) != 2 {
		t.Fatalf("expected 2 warnings, got %d", len(payload.Export.Warnings))
	}
}

func TestRunExportTableWarningsPersisted(t *testing.T) {
	tmpRepo := t.TempDir()
	prevExport := exportFromDatabase
	prevPreview := previewFromDir
	prevCfg := compiledGetObjectConfig
	defer func() {
		exportFromDatabase = prevExport
		previewFromDir = prevPreview
		compiledGetObjectConfig = prevCfg
	}()

	compiledGetObjectConfig = getobject.AppConfig{
		ConnectionString:      "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable",
		RepoRoot:              tmpRepo,
		DbObjectDir:           "db_object",
		Schema:                "dbo",
		ConnectTimeoutSeconds: 10,
		QueryTimeoutSeconds:   20,
	}
	exportFromDatabase = func(cfg getobject.AppConfig, opts getobject.ExportOptions) (getobject.ExportSummary, error) {
		return getobject.ExportSummary{
			TableWarnings: []getobject.ExportWarning{
				{
					Rule:     "TABLE_EXTRACT_SKIPPED_UNSUPPORTED",
					Category: "tables",
					Object:   "dbo.TableA",
					Message:  "build table snapshot dbo.TableA: computed column calc_col is not supported in extractor output",
				},
			},
		}, nil
	}

	exit := run([]string{
		"rza_get_object_db",
	})
	if exit != 0 {
		t.Fatalf("expected exit 0 with table warnings, got %d", exit)
	}

	summaryPath := filepath.Join(tmpRepo, "out", "get_object_summary.json")
	raw, err := os.ReadFile(summaryPath)
	if err != nil {
		t.Fatalf("read summary: %v", err)
	}

	var payload struct {
		Export struct {
			WarningCount float64                   `json:"warning_count"`
			Warnings     []string                  `json:"warnings"`
			WarningItems []getobject.ExportWarning `json:"warning_items"`
		} `json:"export"`
	}
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("parse summary json: %v", err)
	}
	if payload.Export.WarningCount != 1 {
		t.Fatalf("expected warning_count=1, got %v", payload.Export.WarningCount)
	}
	if len(payload.Export.Warnings) != 1 {
		t.Fatalf("expected 1 warning message, got %d", len(payload.Export.Warnings))
	}
	if len(payload.Export.WarningItems) != 1 {
		t.Fatalf("expected 1 warning item, got %d", len(payload.Export.WarningItems))
	}
	if payload.Export.WarningItems[0].Rule != "TABLE_EXTRACT_SKIPPED_UNSUPPORTED" {
		t.Fatalf("unexpected warning rule: %s", payload.Export.WarningItems[0].Rule)
	}
}

func TestRunRequireHealthyConstraintsFailsAfterPersistingExportSummary(t *testing.T) {
	tmpRepo := t.TempDir()
	prevExport := exportFromDatabase
	prevPreview := previewFromDir
	prevCfg := compiledGetObjectConfig
	defer func() {
		exportFromDatabase = prevExport
		previewFromDir = prevPreview
		compiledGetObjectConfig = prevCfg
	}()

	compiledGetObjectConfig = getobject.AppConfig{
		ConnectionString:      "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable",
		RepoRoot:              tmpRepo,
		DbObjectDir:           "db_object",
		Schema:                "dbo",
		ConnectTimeoutSeconds: 10,
		QueryTimeoutSeconds:   20,
	}
	disabled := false
	untrusted := false
	exportFromDatabase = func(cfg getobject.AppConfig, opts getobject.ExportOptions) (getobject.ExportSummary, error) {
		return getobject.ExportSummary{
			Categories: []getobject.CategoryChange{{Dir: "db_object/tables", Created: []string{"dbo.Child.sql"}}},
			TableWarnings: []getobject.ExportWarning{{
				Rule:           "CONSTRAINT_STATE_REVIEW_REQUIRED",
				Category:       "tables",
				Object:         "dbo.Child",
				Constraint:     "FK_Child_Parent",
				ConstraintType: "FOREIGN_KEY",
				Enabled:        &disabled,
				Trusted:        &untrusted,
				Message:        "foreign key is disabled/untrusted",
			}},
		}, nil
	}

	exit := run([]string{"rza_get_object_db", "-require-healthy-constraints"})
	if exit != 1 {
		t.Fatalf("expected health gate exit 1, got %d", exit)
	}

	raw, err := os.ReadFile(filepath.Join(tmpRepo, "out", "get_object_summary.json"))
	if err != nil {
		t.Fatalf("read summary: %v", err)
	}
	var payload struct {
		Status string `json:"status"`
		Error  string `json:"error"`
		Export struct {
			Created      int                       `json:"created"`
			WarningCount int                       `json:"warning_count"`
			WarningItems []getobject.ExportWarning `json:"warning_items"`
		} `json:"export"`
		Config struct {
			RequireHealthyConstraints bool `json:"require_healthy_constraints"`
		} `json:"config"`
	}
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("parse summary: %v", err)
	}
	if payload.Status != "fail" || !strings.Contains(payload.Error, "health gate failed") {
		t.Fatalf("unexpected health gate result: status=%q error=%q", payload.Status, payload.Error)
	}
	if payload.Export.Created != 1 || payload.Export.WarningCount != 1 || len(payload.Export.WarningItems) != 1 {
		t.Fatalf("health gate discarded export evidence: %#v", payload.Export)
	}
	if !payload.Config.RequireHealthyConstraints {
		t.Fatal("health gate config was not recorded")
	}
}
