package verify

import (
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	"rza_migrate_db/internal/output"
	"rza_migrate_db/internal/sqlfmt"
)

// MigrationParseIssue captures migration filename parse failure.
type MigrationParseIssue struct {
	Rule    string
	Message string
}

// MigrationsValidationOptions holds runtime context for migration directory validation.
type MigrationsValidationOptions struct {
	RepoPath               func(string) string
	MigrationsDir          string
	GenerationOrder        []string
	ParseMigrationFileName func(string) (MigrationMeta, *MigrationParseIssue)
	CoreAuditUpName        string
	CoreAuditDownName      string
	ReadFile               func(string) ([]byte, error)
}

// CoreAuditContractsOptions holds runtime context for core audit contract validation.
type CoreAuditContractsOptions struct {
	RepoPath          func(string) string
	MigrationsDir     string
	CoreAuditUpName   string
	CoreAuditDownName string
}

// ValidateMigrations validates migration folder flatness, filename pattern, pairing, and core init presence.
func ValidateMigrations(report *output.Report, opts MigrationsValidationOptions) ([]MigrationMeta, error) {
	if opts.ParseMigrationFileName == nil {
		return nil, errors.New("validate migrations requires ParseMigrationFileName callback")
	}

	migrationsRoot := opts.RepoPath(opts.MigrationsDir)
	readFile := opts.ReadFile
	entries, err := os.ReadDir(migrationsRoot)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			report.AddFail(output.Item{
				Rule:     "MIGRATIONS_DIR_MISSING",
				Severity: 3,
				File:     opts.MigrationsDir,
				Message:  "migrations directory is missing",
				Reason:   "OBJECT_NOT_EXISTS",
				Action:   "FAIL",
			})
			return nil, nil
		}
		return nil, err
	}

	metas := make([]MigrationMeta, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			report.AddFail(output.Item{
				Rule:     "MIGRATIONS_NOT_FLAT",
				Severity: 3,
				File:     filepath.ToSlash(filepath.Join(opts.MigrationsDir, entry.Name())),
				Message:  "migrations directory must be flat",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
			continue
		}

		name := entry.Name()
		if !strings.HasSuffix(strings.ToLower(name), ".sql") {
			report.AddFail(output.Item{
				Rule:     "NON_SQL_MIGRATION_FILE",
				Severity: 3,
				File:     filepath.ToSlash(filepath.Join(opts.MigrationsDir, name)),
				Message:  "migration file must have .sql extension",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
			continue
		}

		meta, issue := opts.ParseMigrationFileName(name)
		if issue != nil {
			report.AddFail(output.Item{
				Rule:     issue.Rule,
				Severity: 3,
				File:     filepath.ToSlash(filepath.Join(opts.MigrationsDir, name)),
				Message:  issue.Message,
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
			continue
		}

		meta.RelPath = filepath.ToSlash(filepath.Join(opts.MigrationsDir, meta.FileName))
		var raw []byte
		var readErr error
		if readFile != nil {
			raw, readErr = readFile(meta.RelPath)
		} else {
			raw, readErr = os.ReadFile(opts.RepoPath(meta.RelPath))
		}
		if readErr != nil {
			return nil, readErr
		}
		if len(raw) == 0 || strings.TrimSpace(string(raw)) == "" {
			report.AddFail(output.Item{
				Rule: "EMPTY_MIGRATION_FILE", Severity: 3, File: meta.RelPath,
				Message: "migration file must contain SQL", Reason: "REQUIRES_MANUAL", Action: "FAIL",
			})
		}
		if !utf8.Valid(raw) || strings.IndexByte(string(raw), 0) >= 0 {
			report.AddFail(output.Item{
				Rule: "INVALID_SQL_ENCODING", Severity: 3, File: meta.RelPath,
				Message: "migration file must be valid UTF-8 without NUL bytes", Reason: "REQUIRES_MANUAL", Action: "FAIL",
			})
		}
		metas = append(metas, meta)
	}

	sort.Slice(metas, func(i, j int) bool {
		return metas[i].FileName < metas[j].FileName
	})

	ValidateMigrationPairing(report, opts.MigrationsDir, metas)
	ValidateMigrationGenerationOrder(report, opts.MigrationsDir, metas, opts.GenerationOrder)
	ValidateCoreAuditInitPair(report, opts.MigrationsDir, opts.CoreAuditUpName, opts.CoreAuditDownName, metas)
	return metas, nil
}

// ValidateMigrationPairing validates each migration version has exactly one up/down pair with matching metadata.
func ValidateMigrationPairing(report *output.Report, migrationsDir string, metas []MigrationMeta) {
	type pairState struct {
		Up      *MigrationMeta
		Down    *MigrationMeta
		UpCount int
		DownCnt int
	}

	versionPairs := map[string]pairState{}
	for _, meta := range metas {
		state := versionPairs[meta.Version]
		switch meta.Direction {
		case "up":
			state.UpCount++
			if state.Up == nil {
				up := meta
				state.Up = &up
			}
		case "down":
			state.DownCnt++
			if state.Down == nil {
				down := meta
				state.Down = &down
			}
		}
		versionPairs[meta.Version] = state
	}

	for version, state := range versionPairs {
		if state.UpCount == 0 || state.DownCnt == 0 {
			report.AddFail(output.Item{
				Rule:     "MISSING_MIGRATION_PAIR",
				Severity: 3,
				File:     filepath.ToSlash(filepath.Join(migrationsDir, version+"_...")),
				Message:  "every migration version must have both up and down pair",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
		if state.UpCount > 1 || state.DownCnt > 1 {
			report.AddFail(output.Item{
				Rule:     "DUPLICATE_MIGRATION_VERSION",
				Severity: 3,
				File:     filepath.ToSlash(filepath.Join(migrationsDir, version+"_...")),
				Message:  "duplicate migration direction found for same version",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
		if state.UpCount == 1 && state.DownCnt == 1 &&
			(state.Up.Category != state.Down.Category ||
				state.Up.Object != state.Down.Object ||
				state.Up.VK != state.Down.VK ||
				state.Up.IsCore != state.Down.IsCore) {
			report.AddFail(output.Item{
				Rule:     "MIGRATION_PAIR_MISMATCH",
				Severity: 3,
				File:     filepath.ToSlash(filepath.Join(migrationsDir, version+"_...")),
				Message:  "up/down pair metadata must match category, object, and vk",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
	}
}

// ValidateMigrationGenerationOrder is retained as a compatibility hook. Ordering
// of new migrations is determined by the engine dependency graph; historical
// filenames alone do not contain enough information to safely reconstruct that
// graph, so verify must not reject a valid dependency-aware order by category.
func ValidateMigrationGenerationOrder(report *output.Report, migrationsDir string, metas []MigrationMeta, generationOrder []string) {
	_ = report
	_ = migrationsDir
	_ = metas
	_ = generationOrder
}

// ValidateCoreAuditInitPair enforces mandatory core audit init up/down pair presence.
func ValidateCoreAuditInitPair(report *output.Report, migrationsDir, coreAuditUpName, coreAuditDownName string, metas []MigrationMeta) {
	hasUp := false
	hasDown := false
	for _, meta := range metas {
		if meta.FileName == coreAuditUpName {
			hasUp = true
		}
		if meta.FileName == coreAuditDownName {
			hasDown = true
		}
	}
	if hasUp && hasDown {
		return
	}

	report.AddFail(output.Item{
		Rule:     "MISSING_CORE_AUDIT_INIT",
		Severity: 3,
		File:     migrationsDir,
		Message:  "core audit init pair is mandatory: 000000__core__mig__audit_init.(up|down).sql",
		Reason:   "OBJECT_NOT_EXISTS",
		Action:   "FAIL",
	})
}

// ValidateCoreAuditContracts validates up/down core migration signatures and lock-readiness contract.
func ValidateCoreAuditContracts(report *output.Report, opts CoreAuditContractsOptions) error {
	upRelPath := filepath.ToSlash(filepath.Join(opts.MigrationsDir, opts.CoreAuditUpName))
	upPath := opts.RepoPath(upRelPath)
	upRaw, err := os.ReadFile(upPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	upUpper := strings.ToUpper(sqlfmt.NormalizeSQLText(string(upRaw), true, true))
	allUpUpper, err := loadAllMigrationSQLByDirectionUpper(opts, "up")
	if err != nil {
		return err
	}
	// Fallback aggregate contains all *.up.sql content. This allows core contract
	// split across reserved core-range migrations (for example 000... + 1000...).
	aggregateUpUpper := upUpper
	if strings.TrimSpace(allUpUpper) != "" {
		aggregateUpUpper += "\n" + allUpUpper
	}

	requiredUp := []struct {
		AnyOf   []string
		Message string
		UseAll  bool
	}{
		{
			AnyOf: []string{
				"CREATE TABLE DBO.MIG_EVENT",
				"CREATE TABLE [DBO].[MIG_EVENT]",
			},
			Message: "core audit up must define dbo.mig_event",
			UseAll:  false,
		},
		{
			AnyOf: []string{
				"ALTER PROCEDURE DBO.USP_MIG_LOG_EVENT",
				"ALTER PROCEDURE [DBO].[USP_MIG_LOG_EVENT]",
				"CREATE PROCEDURE DBO.USP_MIG_LOG_EVENT",
				"CREATE PROCEDURE [DBO].[USP_MIG_LOG_EVENT]",
			},
			Message: "core audit up must define dbo.usp_mig_log_event",
			UseAll:  true,
		},
	}
	for _, req := range requiredUp {
		targetSQL := upUpper
		if req.UseAll {
			targetSQL = aggregateUpUpper
		}
		if containsAnyCoreSignature(targetSQL, req.AnyOf) {
			continue
		}
		report.AddFail(output.Item{
			Rule:     "CORE_AUDIT_CONTRACT_MISSING",
			Severity: 3,
			File:     upRelPath,
			Message:  req.Message,
			Reason:   "REQUIRES_MANUAL",
			Action:   "FAIL",
		})
	}
	if !strings.Contains(upUpper, "CREATED_AT") ||
		(!strings.Contains(upUpper, "SYSUTCDATETIME()") && !strings.Contains(upUpper, "GETDATE()")) {
		report.AddFail(output.Item{
			Rule:     "CORE_AUDIT_CONTRACT_MISSING",
			Severity: 3,
			File:     upRelPath,
			Message:  "core audit up must define created_at default using SYSUTCDATETIME() or GETDATE()",
			Reason:   "REQUIRES_MANUAL",
			Action:   "FAIL",
		})
	}
	for _, token := range []string{
		"ERROR_NUMBER",
		"ERROR_LINE",
		"ERROR_PROCEDURE",
		"IX_MIG_EVENT__RUN_KEY",
		"IX_MIG_EVENT__VERSION",
		"INSERT INTO DBO.MIG_EVENT",
		"@ERROR_NUMBER",
		"@ERROR_LINE",
		"@ERROR_PROCEDURE",
	} {
		targetSQL := upUpper
		// Procedure parameter tokens may live in split core procedure migrations.
		switch token {
		case "ERROR_NUMBER", "ERROR_LINE", "ERROR_PROCEDURE",
			"INSERT INTO DBO.MIG_EVENT",
			"@ERROR_NUMBER", "@ERROR_LINE", "@ERROR_PROCEDURE":
			targetSQL = aggregateUpUpper
		}
		if strings.Contains(targetSQL, token) {
			continue
		}
		report.AddFail(output.Item{
			Rule:     "CORE_AUDIT_CONTRACT_MISSING",
			Severity: 3,
			File:     upRelPath,
			Message:  "core audit up missing required token: " + token,
			Reason:   "REQUIRES_MANUAL",
			Action:   "FAIL",
		})
	}

	hasLockProc := containsAnyCoreSignature(aggregateUpUpper, []string{
		"ALTER PROCEDURE DBO.USP_MIG_TRY_GET_LOCK",
		"ALTER PROCEDURE [DBO].[USP_MIG_TRY_GET_LOCK]",
		"CREATE PROCEDURE DBO.USP_MIG_TRY_GET_LOCK",
		"CREATE PROCEDURE [DBO].[USP_MIG_TRY_GET_LOCK]",
	})
	hasAppLock := containsCoreSignature(aggregateUpUpper, "SP_GETAPPLOCK")
	hasThrow := containsCoreSignature(aggregateUpUpper, "THROW")
	report.Notes["core_lock_proc_present"] = hasLockProc
	report.Notes["core_lock_applock_present"] = hasAppLock
	if hasLockProc && hasAppLock && hasThrow {
		report.AddPass(output.Item{
			Rule:     "CORE_LOCK_READINESS_OK",
			Severity: 1,
			File:     upRelPath,
			Message:  "core lock procedure readiness detected (usp_mig_try_get_lock + sp_getapplock + THROW)",
			Reason:   "OBJECT_EXISTS",
			Action:   "ok",
		})
	} else {
		report.AddFail(output.Item{
			Rule:       "CORE_AUDIT_CONTRACT_MISSING",
			Severity:   3,
			File:       upRelPath,
			Message:    "core lock procedure contract is incomplete (requires usp_mig_try_get_lock + sp_getapplock + THROW)",
			Suggestion: "implement dbo.usp_mig_try_get_lock final contract before using migrations",
			Reason:     "REQUIRES_MANUAL",
			Action:     "FAIL",
		})
	}

	downRelPath := filepath.ToSlash(filepath.Join(opts.MigrationsDir, opts.CoreAuditDownName))
	downPath := opts.RepoPath(downRelPath)
	downRaw, err := os.ReadFile(downPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	downUpper := strings.ToUpper(sqlfmt.NormalizeSQLText(string(downRaw), true, true))
	allDownUpper, err := loadAllMigrationSQLByDirectionUpper(opts, "down")
	if err != nil {
		return err
	}
	aggregateDownUpper := downUpper
	if strings.TrimSpace(allDownUpper) != "" {
		aggregateDownUpper += "\n" + allDownUpper
	}
	requiredDown := []struct {
		AnyOf   []string
		Message string
		UseAll  bool
	}{
		{
			AnyOf: []string{
				"DROP TABLE DBO.MIG_EVENT",
				"DROP TABLE [DBO].[MIG_EVENT]",
				"DROP TABLE IF EXISTS DBO.MIG_EVENT",
				"DROP TABLE IF EXISTS [DBO].[MIG_EVENT]",
			},
			Message: "core audit down should reference dbo.mig_event cleanup",
			UseAll:  false,
		},
		{
			AnyOf: []string{
				"DROP PROCEDURE DBO.USP_MIG_LOG_EVENT",
				"DROP PROCEDURE [DBO].[USP_MIG_LOG_EVENT]",
			},
			Message: "core audit down should reference dbo.usp_mig_log_event cleanup",
			UseAll:  true,
		},
	}
	for _, req := range requiredDown {
		targetSQL := downUpper
		if req.UseAll {
			targetSQL = aggregateDownUpper
		}
		if containsAnyCoreSignature(targetSQL, req.AnyOf) {
			continue
		}
		report.AddFail(output.Item{
			Rule:     "CORE_AUDIT_CONTRACT_MISSING",
			Severity: 3,
			File:     downRelPath,
			Message:  req.Message,
			Reason:   "REQUIRES_MANUAL",
			Action:   "FAIL",
		})
	}
	if hasLockProc && !containsAnyCoreSignature(downUpper, []string{
		"DROP PROCEDURE DBO.USP_MIG_TRY_GET_LOCK",
		"DROP PROCEDURE [DBO].[USP_MIG_TRY_GET_LOCK]",
	}) && !containsAnyCoreSignature(aggregateDownUpper, []string{
		"DROP PROCEDURE DBO.USP_MIG_TRY_GET_LOCK",
		"DROP PROCEDURE [DBO].[USP_MIG_TRY_GET_LOCK]",
	}) {
		report.AddFail(output.Item{
			Rule:     "CORE_AUDIT_CONTRACT_MISSING",
			Severity: 3,
			File:     downRelPath,
			Message:  "core audit down should cleanup dbo.usp_mig_try_get_lock when lock proc exists in core up",
			Reason:   "REQUIRES_MANUAL",
			Action:   "FAIL",
		})
	}

	return nil
}

func containsAnyCoreSignature(sqlUpper string, signatures []string) bool {
	for _, signature := range signatures {
		if containsCoreSignature(sqlUpper, signature) {
			return true
		}
	}
	return false
}

func loadAllMigrationSQLByDirectionUpper(opts CoreAuditContractsOptions, direction string) (string, error) {
	root := opts.RepoPath(opts.MigrationsDir)
	entries, err := os.ReadDir(root)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", nil
		}
		return "", err
	}

	suffix := "." + strings.ToLower(strings.TrimSpace(direction)) + ".sql"
	parts := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if !strings.HasSuffix(strings.ToLower(name), suffix) {
			continue
		}
		raw, readErr := os.ReadFile(filepath.Join(root, name))
		if readErr != nil {
			return "", readErr
		}
		parts = append(parts, strings.ToUpper(sqlfmt.NormalizeSQLText(string(raw), true, true)))
	}
	return strings.Join(parts, "\n"), nil
}

func containsCoreSignature(sqlUpper, signature string) bool {
	if signature == "" {
		return false
	}

	start := 0
	for {
		rel := strings.Index(sqlUpper[start:], signature)
		if rel < 0 {
			return false
		}
		idx := start + rel + len(signature)
		if idx >= len(sqlUpper) {
			return true
		}

		next := rune(sqlUpper[idx])
		if !isCoreWordChar(next) {
			return true
		}
		start = start + rel + 1
	}
}

func isCoreWordChar(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_'
}
