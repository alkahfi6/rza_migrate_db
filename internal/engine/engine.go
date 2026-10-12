package engine

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode"

	"rza_migrate_db/internal/buildinfo"
	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/diff"
	"rza_migrate_db/internal/fsx"
	"rza_migrate_db/internal/getobject"
	"rza_migrate_db/internal/history"
	"rza_migrate_db/internal/management"
	"rza_migrate_db/internal/output"
	"rza_migrate_db/internal/sqlfix"
	"rza_migrate_db/internal/sqlfmt"
	"rza_migrate_db/internal/sqlparse"
	"rza_migrate_db/internal/tsqlast"
	"rza_migrate_db/internal/verify"
)

const (
	coreAuditUpName   = "000000__core__mig__audit_init.up.sql"
	coreAuditDownName = "000000__core__mig__audit_init.down.sql"
)

type Engine struct {
	cfg   config.AppConfig
	state *runState
}

type genPrecheckHistoryCache struct {
	Ready         bool
	History       map[string]historyObject
	ExistingFiles map[string]struct{}
	MaxVersion    string
	DesiredReady  bool
	Desired       []desiredObject
}

type migrationMeta struct {
	FileName  string
	RelPath   string
	Version   string
	Category  string
	Object    string
	Direction string
	VK        int
	IsCore    bool
}

type parseIssue struct {
	Rule    string
	Message string
}

func New(cfg config.AppConfig) Engine {
	cfg.ValidationProfile = strings.ToLower(strings.TrimSpace(cfg.ValidationProfile))
	cfg.MigrationValidation = strings.ToLower(strings.TrimSpace(cfg.MigrationValidation))
	return Engine{cfg: cfg, state: &runState{}}
}

func (e Engine) Run(command string) (output.Report, error) {
	e.resetRepositorySnapshot()
	runStarted := time.Now()
	report := output.NewReport(e.cfg.OutputNamingPattern, e.cfg.OutputNamingLineEndings)
	report.Notes["command"] = command
	report.Notes["build"] = buildinfo.Current()
	report.Notes["bootstrap_mode"] = e.cfg.BootstrapMode
	report.Notes["patch_format_mode"] = effectivePatchFormatMode(e.cfg.PatchFormatMode)
	report.Notes["draft_mode"] = e.cfg.DraftMode
	report.Notes["fail_fast_enabled"] = e.cfg.EnableFailFast
	report.Notes["parallel_read_enabled"] = e.cfg.EnableParallelRead
	report.Notes["parallel_read_workers"] = e.cfg.ParallelWorkers
	report.Notes["existing_table_match_mode"] = effectiveExistingTableMatchMode(e.cfg)
	report.Notes["existing_table_required_checks"] = []string{
		"required_column_exists",
		"base_type_compatible",
		"identity_kind",
		"computed_kind",
	}
	report.Notes["existing_table_non_blocking_details"] = []string{
		"extra_columns",
		"column_order",
		"length",
		"precision_scale",
		"nullability",
		"collation",
		"default_constraints",
		"key_check_foreign_constraints",
		"constraint_runtime_state",
		"physical_index_options",
	}

	if err := e.cfg.Validate(); err != nil {
		report.AddFail(output.Item{
			Rule:     "CONFIG_HARD_LOCK_INVALID",
			Severity: 3,
			File:     "internal/config/app_config.go",
			Message:  err.Error(),
			Reason:   "REQUIRES_MANUAL",
			Action:   "FAIL",
		})
		if writeErr := report.WriteWithOptions(e.cfg.OutFilePath(), e.cfg.EnableOutputNotes); writeErr != nil {
			return report, writeErr
		}
		return report, nil
	}
	if e.cfg.BootstrapMode {
		if err := e.seedOptionalRepoFilesForBootstrap(&report); err != nil {
			return report, err
		}
		if err := e.seedCoreAuditForBootstrap(&report); err != nil {
			return report, err
		}
	}

	var runErr error
	switch command {
	case "fmt":
		runErr = e.runFmt(&report)
	case "check":
		runErr = e.runDeveloperCheck(&report)
	case "verify":
		runErr = e.runVerify(&report)
	case "verify_then_gen":
		runErr = e.runVerifyThenGen(&report)
	case "gen":
		runErr = e.runGen(&report)
	case "bootstrap":
		runErr = e.runVerifyThenGen(&report)
	case "draft":
		runErr = e.runGen(&report)
	case "seal":
		runErr = e.runSealDrafts(&report)
	default:
		return report, fmt.Errorf("unknown command %q", command)
	}

	if runErr != nil {
		report.AddFail(output.Item{
			Rule:     "INTERNAL_ERROR",
			Severity: 3,
			File:     ".",
			Message:  runErr.Error(),
			Reason:   "SQL_ERROR",
			Action:   "FAIL",
		})
	}
	e.recordPerf(&report, "run.total", runStarted)

	if err := report.WriteWithOptions(e.cfg.OutFilePath(), e.cfg.EnableOutputNotes); err != nil {
		return report, err
	}

	return report, nil
}

func effectivePatchFormatMode(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	if value == "" {
		return "off"
	}
	return value
}

func effectiveExistingTableMatchMode(cfg config.AppConfig) string {
	mode := strings.ToLower(strings.TrimSpace(cfg.ExistingTableMatchMode))
	if mode == "" {
		return config.ExistingTableMatchBootstrapCompatible
	}
	return mode
}

func (e Engine) runDeveloperCheck(report *output.Report) error {
	stageStarted := time.Now()
	defer e.recordPerf(report, "check.total", stageStarted)

	report.Notes["check_mode"] = "safe_object_fix_then_verify_no_generation"
	report.Notes["generation_enabled"] = false
	report.Notes["validation_profile"] = e.cfg.ValidationProfile
	report.Notes["migration_validation"] = e.cfg.MigrationValidation
	report.Notes["database_connection"] = false
	report.Notes["validation_scope"] = map[string]string{
		"db_object":  "full",
		"patch":      "full",
		"migrations": e.cfg.MigrationValidation,
	}
	report.SetNewMigrationFiles([]string{})

	fmtStarted := time.Now()
	if err := e.runFmt(report); err != nil {
		return err
	}
	e.recordPerf(report, "check.stage_fmt", fmtStarted)
	if report.Fail.Count > 0 {
		return nil
	}
	constraintFixStarted := time.Now()
	if err := e.runConstraintNameFix(report); err != nil {
		return err
	}
	e.recordPerf(report, "check.stage_constraint_name_fix", constraintFixStarted)
	if report.Fail.Count > 0 {
		return nil
	}

	// runFmt may update desired-state files. Force verify to build its immutable
	// snapshot from the post-format bytes.
	e.resetRepositorySnapshot()
	verifyStarted := time.Now()
	verifyErr := e.runVerifyChecks(report, e.cfg.MigrationValidation == config.MigrationValidationFull, nil)
	if attestErr := e.recordCheckAttestation(report); attestErr != nil {
		return attestErr
	}
	if verifyErr != nil {
		return verifyErr
	}
	e.recordPerf(report, "check.stage_verify", verifyStarted)
	if report.Fail.Count > 0 {
		return nil
	}

	report.AddPass(output.Item{
		Rule:     "MIGRATE_CHECK_OK",
		Severity: 0,
		File:     ".",
		Message:  "developer check completed: db_object formatted and repository verified without generating migrations",
		Reason:   "OBJECT_EXISTS",
		Action:   "ok",
	})
	report.AddPass(output.Item{
		Rule:     "EXISTING_TABLE_MATCH_PROFILE_OK",
		Severity: 0,
		File:     "internal/config/app_config.go",
		Message:  "generated CREATE TABLE wrappers use bootstrap-compatible matching; explicit schema mutations retain strict prechecks",
		Reason:   "OBJECT_EXISTS",
		Action:   "ok",
	})
	return nil
}

func (e Engine) runFmt(report *output.Report) error {
	stageStarted := time.Now()
	defer e.recordPerf(report, "fmt.total", stageStarted)

	failCountBefore := report.Fail.Count
	cache, cacheErr := e.loadFmtCache()
	if cacheErr != nil {
		report.Notes["fmt_cache_reset_reason"] = cacheErr.Error()
		cache = emptyFmtCache()
	}
	nextCache := emptyFmtCache()
	report.Notes["fmt_cache_enabled"] = true
	report.Notes["fmt_check_enabled"] = e.cfg.EnableFmtCheck
	renameCount, renamePaths, err := e.normalizeIndexSourceFileNames(report)
	if err != nil {
		return err
	}
	report.Notes["fmt_index_renames"] = renamePaths

	dbRoot := e.cfg.RepoPath(e.cfg.DbObjectDir)
	catalogManifest, err := e.loadRepoCatalogManifest()
	if err != nil {
		return fmt.Errorf("load rza_migrate.yaml: %w", err)
	}
	localObjects, err := e.desiredStateObjectNames(dbRoot)
	if err != nil {
		return err
	}
	report.Notes["catalog_self"] = catalogManifest.Catalog.Self
	report.Notes["catalog_collation"] = catalogManifest.Catalog.Collation
	files, err := listFilesRecursive(dbRoot)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			report.AddPass(output.Item{
				Rule:     "FMT_SKIPPED_DB_OBJECT_MISSING",
				Severity: 0,
				File:     e.cfg.DbObjectDir,
				Message:  "db_object directory not found, nothing to format",
				Reason:   "OBJECT_NOT_EXISTS",
				Action:   "skip",
			})
			return nil
		}
		return err
	}

	type fmtTask struct {
		Index int
		Abs   string
		Rel   string
	}
	type fmtResult struct {
		Index     int
		Rel       string
		Changed   bool
		CacheHit  bool
		Entry     fmtCacheEntry
		HasEntry  bool
		Fail      *output.Item
		Err       error
		Formatted []byte
		Fixes     []sqlfix.Change
	}

	processTask := func(task fmtTask) fmtResult {
		raw, err := os.ReadFile(task.Abs)
		if err != nil {
			return fmtResult{Index: task.Index, Rel: task.Rel, Err: err}
		}
		utf8Source, encodingFix, err := sqlfmt.NormalizeSQLToUTF8(raw)
		if err != nil {
			fail := output.Item{
				Rule:       "INVALID_SQL_ENCODING",
				Severity:   3,
				File:       task.Rel,
				Message:    err.Error(),
				Suggestion: "save the file as UTF-8, or use a supported UTF-16/UTF-32 Unicode encoding with a valid BOM",
				Reason:     "REQUIRES_MANUAL",
				Action:     "FAIL",
			}
			return fmtResult{Index: task.Index, Rel: task.Rel, Fail: &fail}
		}
		contentHash := sqlContentHash(raw)
		if !encodingFix.Changed {
			if prev, ok := cache.Entries[task.Rel]; catalogManifest.Catalog.Self == "" && ok && prev.ContentHash == contentHash {
				return fmtResult{
					Index:    task.Index,
					Rel:      task.Rel,
					CacheHit: true,
					Entry:    prev,
					HasEntry: true,
				}
			}
		}

		original := string(utf8Source)
		formatCategory := ""
		objectName := ""
		if tableObject, ok := tableObjectNameFromPath(task.Rel, e.cfg.DbObjectDir); ok {
			formatCategory = "tables"
			objectName = tableObject
		}
		if typeObject, ok := typeObjectNameFromPath(task.Rel, e.cfg.DbObjectDir); ok {
			formatCategory = "types"
			objectName = typeObject
		}
		if indexObject, ok := indexObjectNameFromPath(task.Rel, e.cfg.DbObjectDir); ok {
			formatCategory = "indexes"
			objectName = indexObject
		}
		if kind, ok := programmabilityKindFromPath(task.Rel, e.cfg.DbObjectDir); ok {
			formatCategory = strings.ToLower(kind) + "s"
			_, objectName = kind, strings.TrimSuffix(filepath.Base(task.Rel), filepath.Ext(task.Rel))
		}

		fixed, fixErr := sqlfix.RewriteSafe(original, sqlfix.Options{
			Category:     formatCategory,
			ObjectName:   objectName,
			SelfCatalog:  catalogManifest.Catalog.Self,
			LocalObjects: localObjects,
		})
		if fixErr != nil {
			fail := output.Item{
				Rule:       "SAFE_FIX_PARSE_BLOCKED",
				Severity:   3,
				File:       task.Rel,
				Message:    fixErr.Error(),
				Suggestion: "repair the reported structural ambiguity; the original file was preserved",
				Reason:     "REQUIRES_MANUAL",
				Action:     "FAIL",
			}
			return fmtResult{Index: task.Index, Rel: task.Rel, Fail: &fail}
		}

		normalized := fixed.SQL
		safetyBaseline := fixed.SQL
		if formatCategory == "tables" {
			normalized = canonicalizeTableSnapshotFormatting(normalized, objectName)
			safetyBaseline = canonicalizeTableSnapshotFormatting(safetyBaseline, objectName)
		}
		if formatCategory == "indexes" {
			// The file contract already binds this snapshot to objectName. Use the
			// same canonical target qualification on both sides of the safety
			// comparison so adding the known dbo schema is not a false positive.
			safetyBaseline = canonicalizeIndexSnapshotFormatting(safetyBaseline, objectName)
			normalized = canonicalizeIndexSnapshotFormatting(normalized, objectName)
		}

		issue := formattingSafetyIssue(formatCategory, safetyBaseline, normalized)
		if issue == "" {
			issue = formattingCommentLayoutIssue(original, normalized)
		}
		if issue != "" {
			fail := output.Item{
				Rule:       "FORMAT_SAFETY_BLOCKED",
				Severity:   3,
				File:       task.Rel,
				Message:    issue,
				Suggestion: "leave the source unchanged and extend the trivia-aware formatter with a focused regression test",
				Reason:     "REQUIRES_MANUAL",
				Action:     "FAIL",
			}
			return fmtResult{Index: task.Index, Rel: task.Rel, Fail: &fail}
		}

		changed := encodingFix.Changed || normalized != string(raw)
		fixes := fixed.Changes
		if encodingFix.Changed {
			fixes = append(fixes, sqlfix.Change{
				Rule:    "SAFE_SQL_ENCODING_UTF8",
				Message: "converted SQL source from " + encodingFix.SourceEncoding + " to UTF-8",
			})
		}

		return fmtResult{
			Index:     task.Index,
			Rel:       task.Rel,
			Changed:   changed,
			Entry:     fmtCacheEntry{Size: int64(len(normalized)), ContentHash: sqlContentHash([]byte(normalized))},
			HasEntry:  true,
			Formatted: []byte(normalized),
			Fixes:     fixes,
		}
	}

	tasks := make([]fmtTask, 0, len(files))
	for _, file := range files {
		if !strings.EqualFold(filepath.Ext(file), ".sql") {
			continue
		}
		tasks = append(tasks, fmtTask{
			Index: len(tasks),
			Abs:   file,
			Rel:   e.toRelPath(file),
		})
	}

	results := make([]fmtResult, len(tasks))
	workers := e.readParallelWorkers(len(tasks))
	if workers <= 1 {
		for _, task := range tasks {
			results[task.Index] = processTask(task)
		}
	} else {
		jobs := make(chan fmtTask)
		var wg sync.WaitGroup
		for i := 0; i < workers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for task := range jobs {
					results[task.Index] = processTask(task)
				}
			}()
		}
		for _, task := range tasks {
			jobs <- task
		}
		close(jobs)
		wg.Wait()
	}

	changedCount := renameCount
	changedFiles := append([]string{}, renamePaths...)
	cacheHits := 0
	safeFixCount := 0
	safeFixRules := map[string]int{}
	for _, res := range results {
		if res.Err != nil {
			return res.Err
		}
		if res.Fail != nil {
			if e.cfg.BootstrapMode && res.Fail.Rule == "FORMAT_SAFETY_BLOCKED" {
				item := *res.Fail
				item.Rule = "FORMAT_SAFETY_SKIPPED_UNSAFE"
				item.Severity = 2
				item.Message += "; bootstrap preserved the original source bytes"
				item.Suggestion = "review the warning after bootstrap; formatting is optional and generation continues"
				item.Action = "warn"
				report.AddPass(item)
			} else {
				report.AddFail(*res.Fail)
			}
			continue
		}
		if res.CacheHit {
			cacheHits++
		}
		if res.Changed {
			changedCount++
			changedFiles = append(changedFiles, res.Rel)
		}
		for _, fix := range res.Fixes {
			safeFixCount++
			safeFixRules[fix.Rule]++
		}
		if res.HasEntry {
			nextCache.Entries[res.Rel] = res.Entry
		}
	}
	sqlFileCount := len(tasks)

	patchWrites, patchFixes, err := e.collectSafePatchFixes()
	if err != nil {
		return err
	}
	for _, pending := range patchWrites {
		changedCount++
		changedFiles = append(changedFiles, pending.Rel)
	}
	for _, fix := range patchFixes {
		safeFixCount++
		safeFixRules[fix.Rule]++
	}
	sort.Strings(changedFiles)

	report.Notes["fmt_changed_files"] = changedCount
	report.Notes["fmt_changed_file_paths"] = changedFiles
	report.Notes["fmt_sql_files"] = sqlFileCount
	report.Notes["fmt_cache_hits"] = cacheHits
	report.Notes["fmt_cache_misses"] = sqlFileCount - cacheHits
	report.Notes["fmt_safe_fix_count"] = safeFixCount
	report.Notes["fmt_safe_fix_rules"] = safeFixRules
	report.Notes["fmt_patch_safe_fix_files"] = len(patchWrites)
	if report.Fail.Count == failCountBefore && !e.cfg.EnableFmtCheck {
		for index, res := range results {
			if !res.Changed {
				continue
			}
			if err := atomicWriteSQLFile(tasks[index].Abs, res.Formatted); err != nil {
				return err
			}
		}
		for _, pending := range patchWrites {
			if err := atomicWriteSQLFile(pending.Abs, pending.Content); err != nil {
				return err
			}
		}
	}
	if report.Fail.Count == failCountBefore {
		if err := e.saveFmtCache(nextCache); err != nil {
			report.Notes["fmt_cache_write_error"] = err.Error()
			report.Notes["fmt_cache_saved"] = false
		} else {
			report.Notes["fmt_cache_saved"] = true
		}
	} else {
		report.Notes["fmt_cache_saved"] = false
	}
	if e.cfg.EnableFmtCheck && changedCount > 0 {
		report.AddFail(output.Item{
			Rule:       "FMT_CHECK_CHANGES_REQUIRED",
			Severity:   2,
			File:       ".",
			Message:    fmt.Sprintf("%d db_object/patch file(s) require safe formatting; no source files were written", changedCount),
			Suggestion: "review fmt_changed_file_paths, then run fmt without -check to apply safe changes",
			Reason:     "REQUIRES_MANUAL",
			Action:     "FAIL",
		})
		return nil
	}
	if changedCount > 0 {
		report.AddPass(output.Item{
			Rule:     "FMT_APPLIED",
			Severity: 0,
			File:     ".",
			Message:  fmt.Sprintf("formatted %d file(s)", changedCount),
			Reason:   "CREATED",
			Action:   "ok",
		})
	} else {
		report.AddPass(output.Item{
			Rule:     "FMT_NO_CHANGES",
			Severity: 0,
			File:     ".",
			Message:  "no formatting changes required",
			Reason:   "OBJECT_EXISTS",
			Action:   "ok",
		})
	}

	return nil
}

func (e Engine) runVerify(report *output.Report) error {
	stageStarted := time.Now()
	defer e.recordPerf(report, "verify.total", stageStarted)

	report.Notes["verify_mode"] = "bootstrap"
	verifyChecksStarted := time.Now()
	if err := e.runVerifyChecks(report, true, nil); err != nil {
		return err
	}
	e.recordPerf(report, "verify.checks", verifyChecksStarted)

	if report.Fail.Count == 0 {
		report.AddPass(output.Item{
			Rule:     "VERIFY_OK",
			Severity: 0,
			File:     ".",
			Message:  "verify completed with no fail items",
			Reason:   "OBJECT_EXISTS",
			Action:   "ok",
		})
	}

	return nil
}

func (e Engine) validateTableInlineDefault(report *output.Report) error {
	return verify.ValidateTableInlineDefault(report, verify.TableInlineDefaultOptions{
		RepoPath:    e.cfg.RepoPath,
		DbObjectDir: e.cfg.DbObjectDir,
		ToRelPath:   e.toRelPath,
		ReadFile:    e.readRepositoryFile,
	})
}

func (e Engine) runGen(report *output.Report) error {
	stageStarted := time.Now()
	defer e.recordPerf(report, "gen.total", stageStarted)

	report.Notes["gen_mode"] = "emit"
	report.Notes["gen_auto_indent_enabled"] = e.cfg.EnableGenAutoIndent

	fmtStarted := time.Now()
	if err := e.runFmt(report); err != nil {
		return err
	}
	e.recordPerf(report, "gen.stage_fmt", fmtStarted)

	precheckStarted := time.Now()
	var precheckCache genPrecheckHistoryCache
	if err := e.runGenPrecheck(report, &precheckCache); err != nil {
		return err
	}
	e.recordPerf(report, "gen.stage_precheck", precheckStarted)
	return e.runGenCore(report, &precheckCache)
}

func (e Engine) runVerifyThenGen(report *output.Report) error {
	stageStarted := time.Now()
	defer e.recordPerf(report, "verify_then_gen.total", stageStarted)

	report.Notes["gen_mode"] = "emit"
	report.Notes["verify_then_gen_mode"] = "combined"
	report.Notes["gen_auto_indent_enabled"] = e.cfg.EnableGenAutoIndent

	fmtStarted := time.Now()
	if err := e.runFmt(report); err != nil {
		return err
	}
	e.recordPerf(report, "verify_then_gen.stage_fmt", fmtStarted)

	precheckStarted := time.Now()
	var precheckCache genPrecheckHistoryCache
	if err := e.runGenPrecheck(report, &precheckCache); err != nil {
		return err
	}
	e.recordPerf(report, "verify_then_gen.stage_precheck", precheckStarted)
	return e.runGenCore(report, &precheckCache)
}

func (e Engine) runGenCore(report *output.Report, precheckCache *genPrecheckHistoryCache) error {
	if report.Fail.Count > 0 {
		report.SetNewMigrationFiles([]string{})
		return nil
	}

	var (
		historyState  map[string]historyObject
		existingFiles map[string]struct{}
		maxVersion    string
		err           error
	)
	reconstructStarted := time.Now()
	if precheckCache != nil && precheckCache.Ready {
		historyState = cloneHistoryStateMap(precheckCache.History)
		existingFiles = cloneStringSet(precheckCache.ExistingFiles)
		maxVersion = precheckCache.MaxVersion
		report.Notes["gen_precheck_history_cache_hit"] = true
	} else {
		historyState, existingFiles, maxVersion, err = e.reconstructHistory(report)
		if err != nil {
			return err
		}
		report.Notes["gen_precheck_history_cache_hit"] = false
	}
	e.recordPerf(report, "gen.reconstruct_history", reconstructStarted)
	if report.Fail.Count > 0 {
		report.SetNewMigrationFiles([]string{})
		return nil
	}

	desiredStarted := time.Now()
	var desiredObjects []desiredObject
	if precheckCache != nil && precheckCache.DesiredReady {
		desiredObjects = cloneDesiredObjects(precheckCache.Desired)
		report.Notes["gen_precheck_desired_cache_hit"] = true
	} else {
		desiredObjects, err = e.collectDesiredObjects(report)
		if err != nil {
			return err
		}
		report.Notes["gen_precheck_desired_cache_hit"] = false
	}
	e.recordPerf(report, "gen.collect_desired", desiredStarted)
	if report.Fail.Count > 0 {
		report.SetNewMigrationFiles([]string{})
		return nil
	}

	constraintReconcileStarted := time.Now()
	reconciledConstraintNames, reconciledConstraintFiles, err := e.reconcileDesiredTableConstraintNames(historyState, desiredObjects)
	if err != nil {
		return err
	}
	e.recordPerf(report, "gen.reconcile_constraint_names", constraintReconcileStarted)
	report.Notes["gen_reconciled_constraint_names"] = reconciledConstraintNames
	report.Notes["gen_reconciled_constraint_files"] = reconciledConstraintFiles
	if reconciledConstraintNames > 0 {
		e.resetRepositorySnapshot()
		recollectDesiredStarted := time.Now()
		desiredObjects, err = e.collectDesiredObjects(report)
		if err != nil {
			return err
		}
		e.recordPerf(report, "gen.recollect_desired_after_constraint_names", recollectDesiredStarted)
		if report.Fail.Count > 0 {
			report.SetNewMigrationFiles([]string{})
			return nil
		}
		report.AddPass(output.Item{
			Rule:     "CONSTRAINT_NAMES_RECONCILED",
			Severity: 0,
			File:     e.cfg.DbObjectDir,
			Message:  fmt.Sprintf("restored %d equivalent constraint name(s) from latest migration history across %d file(s)", reconciledConstraintNames, len(reconciledConstraintFiles)),
			Reason:   "ALTERED",
			Action:   "ok",
		})
	}

	patchesStarted := time.Now()
	patches, err := e.collectPatchSources(report)
	if err != nil {
		return err
	}
	e.recordPerf(report, "gen.collect_patch_sources", patchesStarted)
	if report.Fail.Count > 0 {
		report.SetNewMigrationFiles([]string{})
		return nil
	}

	if e.cfg.EnableGenAutoIndent {
		changedDesired := e.detectChangedDesiredObjectsFromHistory(historyState, desiredObjects)
		autoIndentModuleStarted := time.Now()
		autoIndentCount, autoIndentFiles, err := e.autoIndentChangedModuleSources(report, changedDesired)
		if err != nil {
			return err
		}
		e.recordPerf(report, "gen.auto_indent_modules", autoIndentModuleStarted)
		if report.Fail.Count > 0 {
			report.SetNewMigrationFiles([]string{})
			return nil
		}
		if autoIndentCount > 0 {
			e.resetRepositorySnapshot()
			report.Notes["gen_auto_indent_changed_files"] = autoIndentFiles
			report.AddPass(output.Item{
				Rule:     "GEN_AUTO_INDENT_APPLIED",
				Severity: 0,
				File:     e.cfg.DbObjectDir,
				Message:  fmt.Sprintf("auto-formatted %d changed source file(s) (tables/indexes/programmability) without stripping comments", autoIndentCount),
				Reason:   "ALTERED",
				Action:   "ok",
			})
			recollectDesiredStarted := time.Now()
			desiredObjects, err = e.collectDesiredObjects(report)
			if err != nil {
				return err
			}
			e.recordPerf(report, "gen.recollect_desired_after_indent", recollectDesiredStarted)
			if report.Fail.Count > 0 {
				report.SetNewMigrationFiles([]string{})
				return nil
			}
		}

		pendingPatches := e.detectPendingPatchSourcesFromHistory(historyState, patches)
		autoIndentPatchStarted := time.Now()
		patchIndentCount, patchIndentFiles, err := e.autoIndentPendingPatchSources(report, pendingPatches, effectivePatchFormatMode(e.cfg.PatchFormatMode))
		if err != nil {
			return err
		}
		e.recordPerf(report, "gen.auto_indent_patches", autoIndentPatchStarted)
		if report.Fail.Count > 0 {
			report.SetNewMigrationFiles([]string{})
			return nil
		}
		if patchIndentCount > 0 {
			e.resetRepositorySnapshot()
			report.Notes["gen_auto_indent_patch_files"] = patchIndentFiles
			report.AddPass(output.Item{
				Rule:     "GEN_AUTO_INDENT_PATCH_APPLIED",
				Severity: 0,
				File:     e.cfg.PatchDir,
				Message:  fmt.Sprintf("auto-indented %d pending patch sql file(s) without stripping comments", patchIndentCount),
				Reason:   "ALTERED",
				Action:   "ok",
			})
			recollectPatchStarted := time.Now()
			patches, err = e.collectPatchSources(report)
			if err != nil {
				return err
			}
			e.recordPerf(report, "gen.recollect_patch_sources_after_indent", recollectPatchStarted)
			if report.Fail.Count > 0 {
				report.SetNewMigrationFiles([]string{})
				return nil
			}
		}
	}

	planStarted := time.Now()
	ops, existingFiles, maxVersion, err := e.buildMigrationPlan(report, historyState, existingFiles, maxVersion, desiredObjects, patches)
	if err != nil {
		return err
	}
	e.recordPerf(report, "gen.plan_generation", planStarted)
	if report.Fail.Count > 0 {
		report.SetNewMigrationFiles([]string{})
		return nil
	}

	if len(ops) == 0 {
		report.SetNewMigrationFiles([]string{})
		report.AddPass(output.Item{
			Rule:     "GEN_NO_CHANGES",
			Severity: 0,
			File:     e.cfg.MigrationsDir,
			Message:  "no semantic changes detected, no migration emitted",
			Reason:   "OBJECT_EXISTS",
			Action:   "ok",
		})
		return nil
	}

	roundTripStarted := time.Now()
	if err := validatePlannedRoundTrip(historyState, desiredObjects, ops); err != nil {
		report.AddFail(output.Item{
			Rule: "PLAN_ROUND_TRIP_FAILED", Severity: 3, File: e.cfg.DbObjectDir,
			Message: err.Error(), Suggestion: "use a reviewed manual migration or extend the typed capability with tests",
			Reason: "REQUIRES_MANUAL", Action: "fail",
		})
		report.SetNewMigrationFiles([]string{})
		return nil
	}
	e.recordPerf(report, "gen.plan_round_trip", roundTripStarted)

	emitStarted := time.Now()
	if err := e.emitPlannedMigrations(report, ops, existingFiles, maxVersion); err != nil {
		return err
	}
	e.recordPerf(report, "gen.emit_migrations", emitStarted)
	return nil
}

func (e Engine) detectChangedDesiredObjects(report *output.Report) ([]desiredObject, error) {
	historyState, _, _, err := e.reconstructHistory(report)
	if err != nil {
		return nil, err
	}
	if report.Fail.Count > 0 {
		return nil, nil
	}

	desiredObjects, err := e.collectDesiredObjects(report)
	if err != nil {
		return nil, err
	}
	if report.Fail.Count > 0 {
		return nil, nil
	}

	changed := make([]desiredObject, 0)
	for _, d := range desiredObjects {
		key := objectKey(d.Category, d.Object)
		last, hasHistory := historyState[key]
		if hasHistory && last.NormalizedHash == d.NormalizedHash {
			continue
		}
		changed = append(changed, d)
	}
	sortDesiredObjectsByOrder(changed, e.cfg.GenerationOrder)
	return changed, nil
}

func (e Engine) detectChangedDesiredObjectsFromHistory(historyState map[string]historyObject, desiredObjects []desiredObject) []desiredObject {
	changed := make([]desiredObject, 0)
	for _, d := range desiredObjects {
		key := objectKey(d.Category, d.Object)
		last, hasHistory := historyState[key]
		if hasHistory && last.NormalizedHash == d.NormalizedHash {
			continue
		}
		changed = append(changed, d)
	}
	sortDesiredObjectsByOrder(changed, e.cfg.GenerationOrder)
	return changed
}

func (e Engine) autoIndentChangedModuleSources(report *output.Report, desired []desiredObject) (int, []string, error) {
	opts := getobject.ModuleFormatOptions{
		StripComments: false,
		Reindent:      true,
		IndentSize:    2,
		Backend:       "antlr",
	}

	seen := make(map[string]struct{})
	type pendingFormatWrite struct {
		RelPath string
		AbsPath string
		Content []byte
	}
	pendingWrites := make([]pendingFormatWrite, 0)
	failBefore := report.Fail.Count
	for _, d := range desired {
		if !isAutoFormatCategory(d.Category) {
			continue
		}
		relPath := filepath.ToSlash(d.SourceFile)
		if relPath == "" {
			continue
		}
		if _, exists := seen[relPath]; exists {
			continue
		}
		seen[relPath] = struct{}{}

		absPath := e.cfg.RepoPath(relPath)
		raw, err := os.ReadFile(absPath)
		if err != nil {
			return 0, nil, err
		}

		formatted := string(raw)
		switch strings.ToLower(strings.TrimSpace(d.Category)) {
		case "types":
			formatted = normalizeSQLText(string(raw), true, true)
		case "views", "functions", "procedures":
			formatted = getobject.FormatModuleSQLWithOptions(string(raw), opts)
		case "tables":
			formatted = canonicalizeTableSnapshotFormatting(formatted, d.Object)
		case "indexes":
			formatted = canonicalizeIndexSnapshotFormatting(formatted, d.Object)
		default:
			continue
		}
		if formatted == string(raw) {
			continue
		}
		safetyBaseline := string(raw)
		if strings.EqualFold(d.Category, "tables") {
			safetyBaseline = repairTableSnapshotCommaBaseline(safetyBaseline, d.Object)
		} else if strings.EqualFold(d.Category, "indexes") {
			safetyBaseline = canonicalizeIndexSnapshotFormatting(safetyBaseline, d.Object)
		}
		if issue := formattingSafetyIssue(d.Category, safetyBaseline, formatted); issue != "" {
			item := output.Item{
				Rule: "FORMAT_SAFETY_BLOCKED", Severity: 3, File: relPath, Object: d.Object,
				Message: issue, Suggestion: "formatter left this file unchanged; add a focused regression test before broadening formatting",
				Reason: "REQUIRES_MANUAL", Action: "fail",
			}
			if e.cfg.BootstrapMode {
				item.Rule = "FORMAT_SAFETY_SKIPPED_UNSAFE"
				item.Severity = 2
				item.Message += "; bootstrap preserved the original source bytes"
				item.Suggestion = "review the warning after bootstrap; formatting is optional and generation continues"
				item.Action = "warn"
				report.AddPass(item)
			} else {
				report.AddFail(item)
			}
			continue
		}
		pendingWrites = append(pendingWrites, pendingFormatWrite{RelPath: relPath, AbsPath: absPath, Content: []byte(formatted)})
	}
	if report.Fail.Count > failBefore {
		return 0, []string{}, nil
	}
	changedFiles := make([]string, 0, len(pendingWrites))
	for _, pending := range pendingWrites {
		if err := os.WriteFile(pending.AbsPath, pending.Content, 0o644); err != nil {
			return 0, nil, err
		}
		changedFiles = append(changedFiles, pending.RelPath)
	}

	sort.Strings(changedFiles)
	return len(changedFiles), changedFiles, nil
}

func (e Engine) detectPendingPatchSources(report *output.Report) ([]patchSource, error) {
	historyState, _, _, err := e.reconstructHistory(report)
	if err != nil {
		return nil, err
	}
	if report.Fail.Count > 0 {
		return nil, nil
	}

	patches, err := e.collectPatchSources(report)
	if err != nil {
		return nil, err
	}
	if report.Fail.Count > 0 {
		return nil, nil
	}

	pending := make([]patchSource, 0)
	for _, patch := range patches {
		key := objectKey("patch", patch.CanonicalName)
		if _, exists := historyState[key]; exists {
			continue
		}
		pending = append(pending, patch)
	}
	sort.Slice(pending, func(i, j int) bool {
		return pending[i].CanonicalName < pending[j].CanonicalName
	})
	return pending, nil
}

func (e Engine) detectPendingPatchSourcesFromHistory(historyState map[string]historyObject, patches []patchSource) []patchSource {
	pending := make([]patchSource, 0)
	for _, patch := range patches {
		key := objectKey("patch", patch.CanonicalName)
		if _, exists := historyState[key]; exists {
			continue
		}
		pending = append(pending, patch)
	}
	sort.Slice(pending, func(i, j int) bool {
		return pending[i].CanonicalName < pending[j].CanonicalName
	})
	return pending
}

func (e Engine) autoIndentPendingPatchSources(report *output.Report, patches []patchSource, mode string) (int, []string, error) {
	mode = effectivePatchFormatMode(mode)
	if mode == "off" {
		report.Notes["patch_format_skipped"] = true
		return 0, []string{}, nil
	}
	opts := getobject.ModuleFormatOptions{
		StripComments: false,
		Reindent:      true,
		IndentSize:    2,
		Backend:       "antlr",
	}

	seen := make(map[string]struct{})
	type pendingPatchFormatWrite struct {
		RelPath string
		AbsPath string
		Content []byte
	}
	pendingWrites := make([]pendingPatchFormatWrite, 0)
	failBefore := report.Fail.Count
	for _, patch := range patches {
		upPath := patch.UpPath
		if strings.TrimSpace(upPath) == "" {
			upPath = filepath.ToSlash(filepath.Join(patch.RelDir, "up.sql"))
		}
		downPath := patch.DownPath
		if strings.TrimSpace(downPath) == "" {
			downPath = filepath.ToSlash(filepath.Join(patch.RelDir, "down.sql"))
		}
		targets := []string{upPath, downPath}
		for _, relPath := range targets {
			if _, exists := seen[relPath]; exists {
				continue
			}
			seen[relPath] = struct{}{}

			absPath := e.cfg.RepoPath(relPath)
			raw, err := os.ReadFile(absPath)
			if err != nil {
				return 0, nil, err
			}

			formatted := getobject.FormatModuleSQLWithOptions(string(raw), opts)
			if formatted == string(raw) {
				continue
			}
			if issue := formattingSafetyIssue("patch", string(raw), formatted); issue != "" {
				item := output.Item{
					Rule: "PATCH_FORMAT_SKIPPED_UNSAFE", Severity: 2, File: relPath, Object: patch.CanonicalName,
					Message:    issue + "; original patch content was preserved",
					Suggestion: "format the patch manually only if needed; generation can safely continue with the original SQL",
					Reason:     "REQUIRES_MANUAL", Action: "warn",
				}
				if mode == "strict" {
					item.Rule = "FORMAT_SAFETY_BLOCKED"
					item.Severity = 3
					item.Action = "fail"
					report.AddFail(item)
				} else {
					report.AddPass(item)
				}
				continue
			}
			pendingWrites = append(pendingWrites, pendingPatchFormatWrite{RelPath: relPath, AbsPath: absPath, Content: []byte(formatted)})
		}
	}
	if report.Fail.Count > failBefore {
		return 0, []string{}, nil
	}
	changedFiles := make([]string, 0, len(pendingWrites))
	for _, pending := range pendingWrites {
		if err := os.WriteFile(pending.AbsPath, pending.Content, 0o644); err != nil {
			return 0, nil, err
		}
		changedFiles = append(changedFiles, pending.RelPath)
	}

	sort.Strings(changedFiles)
	return len(changedFiles), changedFiles, nil
}

func isProgrammabilityCategory(category string) bool {
	switch category {
	case "views", "functions", "procedures":
		return true
	default:
		return false
	}
}

func isAutoFormatCategory(category string) bool {
	if isProgrammabilityCategory(category) {
		return true
	}
	switch strings.ToLower(strings.TrimSpace(category)) {
	case "types", "tables", "indexes":
		return true
	default:
		return false
	}
}

func (e Engine) runGenPrecheck(report *output.Report, cache *genPrecheckHistoryCache) error {
	report.Notes["verify_mode"] = "gen_precheck"
	stageStarted := time.Now()
	if err := e.runVerifyChecks(report, true, cache); err != nil {
		return err
	}
	e.recordPerf(report, "gen_precheck.verify_checks", stageStarted)

	if report.Fail.Count == 0 {
		report.AddPass(output.Item{
			Rule:     "GEN_PRECHECK_OK",
			Severity: 0,
			File:     ".",
			Message:  "gen precheck completed with no fail items",
			Reason:   "OBJECT_EXISTS",
			Action:   "ok",
		})
	}
	return nil
}

func (e Engine) recordPerf(report *output.Report, stage string, started time.Time) {
	if report == nil || !e.cfg.EnablePerfNotes {
		return
	}
	ms := time.Since(started).Milliseconds()
	entry, ok := report.Notes["perf_ms"]
	if !ok {
		perf := map[string]interface{}{
			stage: ms,
		}
		report.Notes["perf_ms"] = perf
		return
	}
	perf, ok := entry.(map[string]interface{})
	if !ok {
		perf = map[string]interface{}{}
		report.Notes["perf_ms"] = perf
	}
	perf[stage] = ms
}

func (e Engine) runVerifyChecks(report *output.Report, includeWrapper bool, cache *genPrecheckHistoryCache) error {
	if err := e.ensureRepositorySnapshot(report); err != nil {
		return err
	}
	orchestrator := verify.Orchestrator{
		IncludeWrapper:      includeWrapper,
		MigrationValidation: e.cfg.MigrationValidation,
		FailFast:            e.cfg.EnableFailFast,
		EnableParallelRun:   e.cfg.EnableParallelRead,
		ParallelWorkers:     e.cfg.ParallelWorkers,
		EnablePerfNotes:     e.cfg.EnablePerfNotes,
	}
	if err := orchestrator.Run(report, e.verifyCallbacks(cache)); err != nil {
		return err
	}
	if e.cfg.EnableFailFast && report.Fail.Count > 0 {
		return nil
	}
	if e.cfg.MigrationValidation != config.MigrationValidationFull {
		return e.validateForbiddenSQL(report)
	}
	return nil
}

func cloneHistoryStateMap(input map[string]historyObject) map[string]historyObject {
	if len(input) == 0 {
		return map[string]historyObject{}
	}
	out := make(map[string]historyObject, len(input))
	for k, v := range input {
		out[k] = v
	}
	return out
}

func cloneStringSet(input map[string]struct{}) map[string]struct{} {
	if len(input) == 0 {
		return map[string]struct{}{}
	}
	out := make(map[string]struct{}, len(input))
	for k := range input {
		out[k] = struct{}{}
	}
	return out
}

func (e Engine) validateMigrations(report *output.Report) ([]migrationMeta, error) {
	verifyMetas, err := verify.ValidateMigrations(report, verify.MigrationsValidationOptions{
		RepoPath:        e.cfg.RepoPath,
		MigrationsDir:   e.cfg.MigrationsDir,
		GenerationOrder: e.cfg.GenerationOrder,
		ParseMigrationFileName: func(name string) (verify.MigrationMeta, *verify.MigrationParseIssue) {
			meta, issue := e.parseMigrationFileName(name)
			if issue != nil {
				return verify.MigrationMeta{}, &verify.MigrationParseIssue{
					Rule:    issue.Rule,
					Message: issue.Message,
				}
			}
			return toVerifyMigrationMeta(meta), nil
		},
		CoreAuditUpName:   coreAuditUpName,
		CoreAuditDownName: coreAuditDownName,
		ReadFile:          e.readRepositoryFile,
	})
	if err != nil {
		return nil, err
	}
	return toEngineMigrationMetas(verifyMetas), nil
}

func (e Engine) parseMigrationFileName(name string) (migrationMeta, *parseIssue) {
	parsed, issue := history.ParseMigrationFileName(name, history.ParseOptions{
		CoreAuditUpName:         coreAuditUpName,
		CoreAuditDownName:       coreAuditDownName,
		AllowedCategories:       e.cfg.AllowedCategories,
		AllowedObjectTokenRunes: e.cfg.AllowedObjectTokenRunes,
	})
	if issue != nil {
		return migrationMeta{}, &parseIssue{
			Rule:    issue.Rule,
			Message: issue.Message,
		}
	}
	return migrationMeta{
		FileName:  parsed.FileName,
		Version:   parsed.Version,
		Category:  parsed.Category,
		Object:    parsed.Object,
		Direction: parsed.Direction,
		VK:        parsed.VK,
		IsCore:    parsed.IsCore,
	}, nil
}

func (e Engine) validateCoreAuditContracts(report *output.Report) error {
	return verify.ValidateCoreAuditContracts(report, verify.CoreAuditContractsOptions{
		RepoPath:          e.cfg.RepoPath,
		MigrationsDir:     e.cfg.MigrationsDir,
		CoreAuditUpName:   coreAuditUpName,
		CoreAuditDownName: coreAuditDownName,
	})
}

func (e Engine) validateForbiddenSQL(report *output.Report) error {
	return verify.ValidateForbiddenSQL(report, verify.ForbiddenSQLOptions{
		RepoPath:           e.cfg.RepoPath,
		DbObjectDir:        e.cfg.DbObjectDir,
		MigrationsDir:      e.cfg.MigrationsDir,
		PatchDir:           e.cfg.PatchDir,
		ForbiddenSQLTokens: e.cfg.ForbiddenSQLTokens,
		ToRelPath:          e.toRelPath,
		EnableParallelRead: e.cfg.EnableParallelRead,
		ParallelWorkers:    e.cfg.ParallelWorkers,
		ReadFile:           e.readRepositoryFile,
		SkipMigrations:     e.cfg.MigrationValidation != config.MigrationValidationFull,
	})
}

func (e Engine) validatePatchFolders(report *output.Report) error {
	return verify.ValidatePatchFolders(report, verify.PatchFoldersOptions{
		RepoPath: e.cfg.RepoPath,
		PatchDir: e.cfg.PatchDir,
	})
}

func (e Engine) validatePatchSourceIntegrity(report *output.Report, metas []migrationMeta) error {
	return verify.ValidatePatchSourceIntegrity(report, verify.PatchSourceIntegrityOptions{
		RepoPath:              e.cfg.RepoPath,
		PatchDir:              e.cfg.PatchDir,
		Metas:                 toVerifyMigrationMetas(metas),
		CanonicalizePatchName: canonicalizePatchName,
		HashStableContent:     hashStableContent,
	})
}

func (e Engine) validateRepoHygiene(report *output.Report) error {
	return verify.ValidateRepoHygiene(report, verify.RepoHygieneOptions{
		RepoPath: e.cfg.RepoPath,
	})
}

func (e Engine) validateProgrammabilityHeaders(report *output.Report, metas []migrationMeta) error {
	return verify.ValidateProgrammabilityHeaders(report, verify.ProgrammabilityHeadersOptions{
		RepoPath:                       e.cfg.RepoPath,
		DbObjectDir:                    e.cfg.DbObjectDir,
		ToRelPath:                      e.toRelPath,
		EnableSourceContractForNewOnly: e.cfg.EnableProgrammabilitySourceContractForNewOnly,
		Metas:                          toVerifyMigrationMetas(metas),
		ReadFile:                       e.readRepositoryFile,
		SkipMigrationHeaders:           e.cfg.MigrationValidation != config.MigrationValidationFull,
	})
}

func (e Engine) validateWrapperSubset(report *output.Report, metas []migrationMeta) error {
	return verify.ValidateWrapperSubset(report, verify.WrapperSubsetOptions{
		RepoPath:           e.cfg.RepoPath,
		Metas:              toVerifyMigrationMetas(metas),
		AllowedDecisions:   e.cfg.AllowedDecisions,
		AllowedActions:     e.cfg.AllowedActions,
		AllowedReasons:     e.cfg.AllowedReasons,
		AllowedSeverities:  e.cfg.AllowedSeverities,
		StrictMode:         e.cfg.EnableStrictMode,
		EnableParallelRead: e.cfg.EnableParallelRead,
		ParallelWorkers:    e.cfg.ParallelWorkers,
		ReadFile:           e.readRepositoryFile,
	})
}

func compactUpperSQL(input string) string {
	return sqlparse.CompactUpperSQL(input)
}

func (e Engine) validateDbObjectConsistency(report *output.Report) error {
	return verify.ValidateDbObjectConsistency(report, verify.DbObjectConsistencyOptions{
		RepoPath:                     e.cfg.RepoPath,
		DbObjectDir:                  e.cfg.DbObjectDir,
		ToRelPath:                    e.toRelPath,
		TrimTrailingSpaces:           e.cfg.TrimTrailingSpaces,
		EnsureFinalNewline:           e.cfg.EnsureFinalNewline,
		ExtractCreateTypeStatement:   validateTableTypeSource,
		ExtractCreateTableStatement:  extractCreateTableStatement,
		ExtractCreateIndexStatements: extractCreateIndexStatements,
		ExtractProgrammabilityStmt:   extractProgrammabilityStatement,
		ValidateFunctionSource:       validateFunctionSource,
		ReadFile:                     e.readRepositoryFile,
	})
}

func (e Engine) validateDesiredHistoryDrift(report *output.Report, history map[string]historyObject, desired []desiredObject) error {
	if report.Fail.Count > 0 {
		return nil
	}
	if history == nil {
		history = map[string]historyObject{}
	}

	desiredStates := make([]diff.State, 0, len(desired))
	for _, d := range desired {
		if d.Category == "patch" || isEngineInfrastructureObject(d.Category, d.Object) {
			continue
		}
		desiredStates = append(desiredStates, diff.State{
			Category:       d.Category,
			Object:         d.Object,
			NormalizedHash: d.NormalizedHash,
			SourceFile:     d.SourceFile,
		})
	}

	unmanaged, manifestErr := e.loadUnmanagedObjectSet()
	if manifestErr != nil {
		return fmt.Errorf("load object management manifest: %w", manifestErr)
	}
	historyStates := make([]diff.State, 0, len(history))
	for _, h := range history {
		if h.Category == "patch" || h.Deleted || isEngineInfrastructureObject(h.Category, h.Object) {
			continue
		}
		if _, ignored := unmanaged[objectKey(h.Category, h.Object)]; ignored {
			continue
		}
		historyStates = append(historyStates, diff.State{
			Category:       h.Category,
			Object:         h.Object,
			NormalizedHash: h.NormalizedHash,
		})
	}
	sort.Slice(historyStates, func(i, j int) bool {
		if historyStates[i].Category != historyStates[j].Category {
			return historyStates[i].Category < historyStates[j].Category
		}
		return historyStates[i].Object < historyStates[j].Object
	})

	drifts := diff.ComputeDrift(desiredStates, historyStates)
	if len(drifts) == 0 {
		report.AddPass(output.Item{
			Rule:     "REPO_STATE_SYNCED",
			Severity: 0,
			File:     ".",
			Message:  "db_object and migration history are semantically aligned",
			Reason:   "OBJECT_EXISTS",
			Action:   "ok",
		})
		return nil
	}

	if e.isGenPrecheckMode(report) {
		report.AddPass(output.Item{
			Rule:     "GEN_PRECHECK_DIFF_DETECTED",
			Severity: 0,
			File:     ".",
			Message:  fmt.Sprintf("semantic differences detected (%d object(s)); generator planning will handle emit/skip policy", len(drifts)),
			Reason:   "OBJECT_EXISTS",
			Action:   "ok",
		})
		return nil
	}

	for _, d := range drifts {
		switch d.Kind {
		case diff.DriftMissingInHistory:
			report.AddPass(output.Item{
				Rule:       "REPO_STATE_DRIFT",
				Severity:   2,
				File:       d.SourceFile,
				Object:     d.Object,
				Message:    "object exists in db_object but not yet represented in migration history",
				Suggestion: "run `rza_migrate_db gen` to emit migration pair",
				Reason:     "DRIFT_DETECTED",
				Action:     "warn",
			})
		case diff.DriftHashMismatch:
			report.AddPass(output.Item{
				Rule:       "REPO_STATE_DRIFT",
				Severity:   2,
				File:       d.SourceFile,
				Object:     d.Object,
				Message:    "db_object semantic state differs from last migration history snapshot",
				Suggestion: "run `rza_migrate_db gen` and review generated migration",
				Reason:     "DRIFT_DETECTED",
				Action:     "warn",
			})
		case diff.DriftMissingInDesired:
			report.AddPass(output.Item{
				Rule:       "REPO_STATE_DRIFT",
				Severity:   2,
				File:       e.cfg.DbObjectDir,
				Object:     d.Object,
				Message:    "object exists in migration history but missing from db_object desired state",
				Suggestion: "run `rza_migrate_db gen` and review the guarded removal migration",
				Reason:     "DRIFT_DETECTED",
				Action:     "warn",
			})
		}
	}

	return nil
}

func (e Engine) loadUnmanagedObjectSet() (map[string]struct{}, error) {
	manifest, err := management.Load(e.cfg.RepoPath(e.cfg.ManagementManifestFile))
	if err != nil {
		return nil, err
	}
	return manifest.UnmanagedSet(), nil
}

func (e Engine) isGenPrecheckMode(report *output.Report) bool {
	if report == nil || report.Notes == nil {
		return false
	}
	rawMode, ok := report.Notes["verify_mode"]
	if !ok {
		return false
	}
	mode, ok := rawMode.(string)
	if !ok {
		return false
	}
	return strings.EqualFold(strings.TrimSpace(mode), "gen_precheck")
}

func normalizeSQLText(input string, trimTrailingSpaces bool, ensureFinalNewline bool) string {
	return sqlfmt.NormalizeSQLText(input, trimTrailingSpaces, ensureFinalNewline)
}

func tableObjectNameFromPath(relPath, dbObjectRoot string) (string, bool) {
	normalized := strings.ToLower(strings.ReplaceAll(relPath, "\\", "/"))
	base := strings.ToLower(strings.TrimSuffix(strings.ReplaceAll(dbObjectRoot, "\\", "/"), "/"))
	newPrefix := base + "/tables/"
	legacyPrefix := base + "/schema/tables/"
	if !strings.HasPrefix(normalized, newPrefix) && !strings.HasPrefix(normalized, legacyPrefix) {
		return "", false
	}
	fileBase := filepath.Base(relPath)
	if !strings.EqualFold(filepath.Ext(fileBase), ".sql") {
		return "", false
	}
	object := strings.TrimSuffix(fileBase, filepath.Ext(fileBase))
	if !strings.HasPrefix(object, "dbo.") {
		return "", false
	}
	return object, true
}

func typeObjectNameFromPath(relPath, dbObjectRoot string) (string, bool) {
	normalized := strings.ToLower(strings.ReplaceAll(relPath, "\\", "/"))
	base := strings.ToLower(strings.TrimSuffix(strings.ReplaceAll(dbObjectRoot, "\\", "/"), "/"))
	newPrefix := base + "/types/"
	legacyPrefix := base + "/schema/types/"
	if !strings.HasPrefix(normalized, newPrefix) && !strings.HasPrefix(normalized, legacyPrefix) {
		return "", false
	}
	fileBase := filepath.Base(relPath)
	if !strings.EqualFold(filepath.Ext(fileBase), ".sql") {
		return "", false
	}
	object := strings.TrimSuffix(fileBase, filepath.Ext(fileBase))
	if !strings.HasPrefix(object, "dbo.") {
		return "", false
	}
	return object, true
}

func indexObjectNameFromPath(relPath, dbObjectRoot string) (string, bool) {
	normalized := strings.ToLower(strings.ReplaceAll(relPath, "\\", "/"))
	base := strings.ToLower(strings.TrimSuffix(strings.ReplaceAll(dbObjectRoot, "\\", "/"), "/"))
	newPrefix := base + "/indexes/"
	legacyPrefix := base + "/schema/indexes/"
	if !strings.HasPrefix(normalized, newPrefix) && !strings.HasPrefix(normalized, legacyPrefix) {
		return "", false
	}
	fileBase := filepath.Base(relPath)
	if !strings.EqualFold(filepath.Ext(fileBase), ".sql") {
		return "", false
	}
	object := strings.TrimSuffix(fileBase, filepath.Ext(fileBase))
	if !strings.HasPrefix(object, "dbo.") {
		return "", false
	}
	return object, true
}

func programmabilityKindFromPath(relPath, dbObjectRoot string) (string, bool) {
	normalized := strings.ToLower(strings.ReplaceAll(relPath, "\\", "/"))
	base := strings.ToLower(strings.TrimSuffix(strings.ReplaceAll(dbObjectRoot, "\\", "/"), "/"))
	kinds := map[string]string{
		"views":      "VIEW",
		"functions":  "FUNCTION",
		"procedures": "PROCEDURE",
	}
	for category, kind := range kinds {
		newPrefix := base + "/" + category + "/"
		legacyPrefix := base + "/schema/" + category + "/"
		if !strings.HasPrefix(normalized, newPrefix) && !strings.HasPrefix(normalized, legacyPrefix) {
			continue
		}
		fileBase := filepath.Base(relPath)
		if !strings.EqualFold(filepath.Ext(fileBase), ".sql") {
			return "", false
		}
		object := strings.TrimSuffix(fileBase, filepath.Ext(fileBase))
		if !strings.HasPrefix(object, "dbo.") {
			return "", false
		}
		return kind, true
	}
	return "", false
}

func canonicalizeTableSnapshotFormatting(sql, objectName string) string {
	if tsqlast.ContainsSQLComment(sql) {
		return normalizeSQLText(sql, true, true)
	}
	stmt, start, end, err := extractCreateTableStatementBounds(objectName, sql)
	if err != nil {
		return sql
	}
	stmt, _ = repairCreateTableCommaSyntax(stmt)
	items, err := extractCreateTableItems(stmt)
	if err != nil || len(items) == 0 {
		return sql
	}

	normalizedItems := make([]string, 0, len(items))
	for _, item := range items {
		trimmed := strings.TrimSpace(item)
		if trimmed == "" {
			continue
		}
		normalizedItems = append(normalizedItems, trimmed)
	}
	if len(normalizedItems) == 0 {
		return sql
	}

	rebuilt, err := rebuildCreateTableStatementWithItems(stmt, normalizedItems)
	if err != nil {
		return sql
	}
	normalized := normalizeSQLText(sql, true, true)
	if start < 0 || end > len(normalized) || start >= end {
		return normalized
	}
	return normalizeSQLText(normalized[:start]+rebuilt+normalized[end:], true, true)
}

func canonicalizeIndexSnapshotFormatting(sql, objectName string) string {
	if tsqlast.ContainsSQLComment(sql) {
		return normalizeSQLText(sql, true, true)
	}
	if err := validateIndexSubsetSupport(sql); err != nil {
		return normalizeSQLText(sql, true, true)
	}
	indexes, err := parseIndexStateMap(objectName, sql)
	if err != nil || len(indexes) == 0 {
		return normalizeSQLText(sql, true, true)
	}
	return serializeIndexState(indexes)
}

func repairTableSnapshotCommaBaseline(sql, objectName string) string {
	normalized := normalizeSQLText(sql, true, true)
	stmt, start, end, err := extractCreateTableStatementBounds(objectName, normalized)
	if err != nil {
		return normalized
	}
	repaired, _ := repairCreateTableCommaSyntax(stmt)
	if start < 0 || end > len(normalized) || start >= end {
		return normalized
	}
	return normalizeSQLText(normalized[:start]+repaired+normalized[end:], true, true)
}

func formattingSafetyIssue(category, before, after string) string {
	semanticEqual := sqlparse.CanonicalSemanticSQL(before) == sqlparse.CanonicalSemanticSQL(after)
	if category == "tables" || category == "indexes" {
		semanticEqual = hashSemanticByCategory(category, before) == hashSemanticByCategory(category, after)
	}
	if !semanticEqual {
		return "formatter output changed SQL tokens, identifier case, or literal content"
	}
	return formattingCommentLayoutIssue(before, after)
}

func formattingCommentLayoutIssue(before, after string) string {
	beforeComments := tsqlast.SQLCommentLines(before)
	afterComments := tsqlast.SQLCommentLines(after)
	if len(beforeComments) != len(afterComments) {
		return "formatter output added or removed SQL comment lines"
	}
	for i := range beforeComments {
		// Repository normalization intentionally removes trailing horizontal
		// whitespace. That does not alter a comment's indentation or whether it
		// is inline with SQL, so exclude only that suffix from the layout
		// contract. Leading whitespace and all non-whitespace content remain
		// byte-for-byte protected.
		beforeLine := strings.TrimRight(beforeComments[i], " \t")
		afterLine := strings.TrimRight(afterComments[i], " \t")
		if beforeLine != afterLine {
			return "formatter output changed SQL comment indentation or inline placement"
		}
	}
	return ""
}

func listFilesRecursive(root string) ([]string, error) {
	return fsx.ListFilesRecursive(root)
}

func (e Engine) toRelPath(absPath string) string {
	rel, err := filepath.Rel(e.cfg.RepoRoot, absPath)
	if err != nil {
		return filepath.ToSlash(absPath)
	}
	return filepath.ToSlash(rel)
}

func stripSQLStringsAndComments(input string) string {
	return sqlparse.StripSQLStringsAndComments(input)
}

func isWordChar(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_'
}
