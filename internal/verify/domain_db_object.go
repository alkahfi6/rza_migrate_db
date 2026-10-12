package verify

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"rza_migrate_db/internal/fsx"
	"rza_migrate_db/internal/output"
	"rza_migrate_db/internal/sqlfmt"
	"rza_migrate_db/internal/sqlparse"
)

// DbObjectConsistencyOptions holds runtime context for db_object consistency validation.
type DbObjectConsistencyOptions struct {
	RepoPath                     func(string) string
	DbObjectDir                  string
	ToRelPath                    func(string) string
	TrimTrailingSpaces           bool
	EnsureFinalNewline           bool
	ExtractCreateTableStatement  func(string, string) (string, error)
	ExtractCreateTypeStatement   func(string, string) (string, error)
	ExtractCreateIndexStatements func(string, string) ([]string, error)
	ExtractProgrammabilityStmt   func(string, string, string) (string, error)
	ValidateFunctionSource       func(string) (string, error)
	ListFilesRecursive           func(string) ([]string, error)
	ReadFile                     func(string) ([]byte, error)
}

// ProgrammabilityHeadersOptions holds runtime context for programmability header validation.
type ProgrammabilityHeadersOptions struct {
	RepoPath                       func(string) string
	DbObjectDir                    string
	ToRelPath                      func(string) string
	EnableSourceContractForNewOnly bool
	Metas                          []MigrationMeta
	ListFilesRecursive             func(string) ([]string, error)
	ReadFile                       func(string) ([]byte, error)
	SkipMigrationHeaders           bool
}

// ValidateDbObjectConsistency validates db_object source contract by category.
func ValidateDbObjectConsistency(report *output.Report, opts DbObjectConsistencyOptions) error {
	readFile := opts.ReadFile
	if readFile == nil {
		readFile = os.ReadFile
	}
	type sourceDir struct {
		Category string
		RelPath  string
	}

	listFiles := opts.ListFilesRecursive
	if listFiles == nil {
		listFiles = fsx.ListFilesRecursive
	}

	dirs := []sourceDir{
		{Category: "types", RelPath: filepath.Join(opts.DbObjectDir, "types")},
		{Category: "tables", RelPath: filepath.Join(opts.DbObjectDir, "tables")},
		{Category: "indexes", RelPath: filepath.Join(opts.DbObjectDir, "indexes")},
		{Category: "views", RelPath: filepath.Join(opts.DbObjectDir, "views")},
		{Category: "functions", RelPath: filepath.Join(opts.DbObjectDir, "functions")},
		{Category: "procedures", RelPath: filepath.Join(opts.DbObjectDir, "procedures")},
	}

	for _, item := range dirs {
		canonicalFiles, err := listCategorySQLFiles(opts.RepoPath, listFiles, filepath.Join(opts.DbObjectDir, item.Category))
		if err != nil {
			return err
		}
		legacyFiles, err := listCategorySQLFiles(opts.RepoPath, listFiles, filepath.Join(opts.DbObjectDir, "schema", item.Category))
		if err != nil {
			return err
		}
		if len(legacyFiles) > 0 {
			report.AddPass(output.Item{
				Rule:       "LEGACY_DB_OBJECT_LAYOUT_DETECTED",
				Severity:   2,
				File:       filepath.ToSlash(filepath.Join(opts.DbObjectDir, "schema", item.Category)),
				Message:    "legacy db_object layout detected; migrate files to canonical db_object/" + item.Category,
				Suggestion: "move files from db_object/schema/" + item.Category + " to db_object/" + item.Category + " and remove legacy path",
				Reason:     "REQUIRES_MANUAL",
				Action:     "warn",
			})
		}
		for base, legacyAbs := range legacyFiles {
			canonicalAbs, exists := canonicalFiles[base]
			if !exists {
				continue
			}
			legacyRaw, readLegacyErr := readFile(legacyAbs)
			if readLegacyErr != nil {
				return readLegacyErr
			}
			canonicalRaw, readCanonicalErr := readFile(canonicalAbs)
			if readCanonicalErr != nil {
				return readCanonicalErr
			}
			legacyNorm := sqlfmt.NormalizeSQLText(string(legacyRaw), opts.TrimTrailingSpaces, opts.EnsureFinalNewline)
			canonicalNorm := sqlfmt.NormalizeSQLText(string(canonicalRaw), opts.TrimTrailingSpaces, opts.EnsureFinalNewline)
			objectName := strings.TrimSuffix(filepath.Base(canonicalAbs), filepath.Ext(canonicalAbs))
			if canonicalNorm != legacyNorm {
				report.AddFail(output.Item{
					Rule:       "DB_OBJECT_DUAL_PATH_CONFLICT",
					Severity:   3,
					File:       opts.ToRelPath(legacyAbs),
					Object:     objectName,
					Message:    "same object exists in canonical and legacy path with different content",
					Suggestion: "keep only canonical file and remove/merge legacy duplicate",
					Reason:     "REQUIRES_MANUAL",
					Action:     "FAIL",
				})
				continue
			}
			report.AddPass(output.Item{
				Rule:       "LEGACY_DB_OBJECT_DUPLICATE_CANONICAL",
				Severity:   2,
				File:       opts.ToRelPath(legacyAbs),
				Object:     objectName,
				Message:    "legacy duplicate file matches canonical content",
				Suggestion: "remove legacy duplicate file to avoid dual-path ambiguity",
				Reason:     "OBJECT_EXISTS",
				Action:     "warn",
			})
		}

		files, err := fsx.ResolveDbObjectCategoryFiles(opts.RepoPath, listFiles, opts.DbObjectDir, item.Category)
		if err != nil {
			return err
		}

		for _, file := range files {
			if !strings.EqualFold(filepath.Ext(file), ".sql") {
				continue
			}

			relFile := opts.ToRelPath(file)
			baseName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
			if !strings.HasPrefix(baseName, "dbo.") {
				report.AddFail(output.Item{
					Rule:     "FILE_OBJECT_MISMATCH",
					Severity: 3,
					File:     relFile,
					Message:  "db_object filename must use dbo.<Object>.sql",
					Reason:   "REQUIRES_MANUAL",
					Action:   "FAIL",
				})
				continue
			}

			raw, err := readFile(file)
			if err != nil {
				return err
			}
			normalized := sqlfmt.NormalizeSQLText(string(raw), opts.TrimTrailingSpaces, opts.EnsureFinalNewline)
			validateDbObjectCategoryPlacement(report, relFile, item.Category, normalized)

			switch item.Category {
			case "types":
				if opts.ExtractCreateTypeStatement == nil {
					return fmt.Errorf("type source validator is not configured")
				}
				if _, err := opts.ExtractCreateTypeStatement(baseName, normalized); err != nil {
					report.AddFail(output.Item{
						Rule:     "TYPE_SOURCE_INVALID",
						Severity: 3,
						File:     relFile,
						Object:   baseName,
						Message:  err.Error(),
						Reason:   "REQUIRES_MANUAL",
						Action:   "FAIL",
					})
				}
			case "tables":
				ValidateTableSnapshotContract(report, relFile, normalized)
				if _, err := opts.ExtractCreateTableStatement(baseName, normalized); err != nil {
					report.AddFail(output.Item{
						Rule:     "FILE_OBJECT_MISMATCH",
						Severity: 3,
						File:     relFile,
						Object:   baseName,
						Message:  "table definition does not match filename object",
						Reason:   "REQUIRES_MANUAL",
						Action:   "FAIL",
					})
				}
			case "indexes":
				ValidateIndexSnapshotContract(report, relFile, normalized, baseName)
				if _, err := opts.ExtractCreateIndexStatements(baseName, normalized); err != nil {
					report.AddFail(output.Item{
						Rule:     "FILE_OBJECT_MISMATCH",
						Severity: 3,
						File:     relFile,
						Object:   baseName,
						Message:  "index definition does not target table from filename",
						Reason:   "REQUIRES_MANUAL",
						Action:   "FAIL",
					})
				}
			case "views", "functions", "procedures":
				if _, err := opts.ExtractProgrammabilityStmt(item.Category, baseName, normalized); err != nil {
					report.AddFail(output.Item{
						Rule:     "FILE_OBJECT_MISMATCH",
						Severity: 3,
						File:     relFile,
						Object:   baseName,
						Message:  "programmability definition does not match filename object",
						Reason:   "REQUIRES_MANUAL",
						Action:   "FAIL",
					})
				}
				if item.Category == "functions" && opts.ValidateFunctionSource != nil {
					if subtype, err := opts.ValidateFunctionSource(normalized); err != nil {
						report.AddFail(output.Item{
							Rule:     "FUNCTION_SUBTYPE_INVALID",
							Severity: 3,
							File:     relFile,
							Object:   baseName,
							Message:  err.Error(),
							Reason:   "REQUIRES_MANUAL",
							Action:   "FAIL",
						})
					} else {
						report.AddPass(output.Item{
							Rule:     "FUNCTION_SUBTYPE_DETECTED",
							Severity: 0,
							File:     relFile,
							Object:   baseName,
							Message:  "function subtype: " + subtype,
							Reason:   "OBJECT_EXISTS",
							Action:   "ok",
						})
					}
				}
			}
		}
	}

	return nil
}

func validateDbObjectCategoryPlacement(report *output.Report, relFile, actualCategory, sql string) {
	scanned := strings.ToUpper(sqlparse.StripSQLStringsAndComments(sql))
	// Category detection is lexical, so insignificant whitespace must not make
	// a valid CREATE OR ALTER header disappear and expose temp-table DDL later
	// in a procedure as a false table classification.
	scanned = strings.Join(strings.Fields(scanned), " ")
	detected := ""
	switch {
	case len(findAllKeywordPositions(scanned, "CREATE TYPE")) > 0:
		detected = "types"
	case len(findAllKeywordPositions(scanned, "CREATE OR ALTER VIEW")) > 0 || len(findAllKeywordPositions(scanned, "CREATE VIEW")) > 0:
		detected = "views"
	case len(findAllKeywordPositions(scanned, "CREATE OR ALTER FUNCTION")) > 0 || len(findAllKeywordPositions(scanned, "CREATE FUNCTION")) > 0:
		detected = "functions"
	case len(findAllKeywordPositions(scanned, "CREATE OR ALTER PROCEDURE")) > 0 || len(findAllKeywordPositions(scanned, "CREATE PROCEDURE")) > 0 || len(findAllKeywordPositions(scanned, "CREATE OR ALTER PROC")) > 0 || len(findAllKeywordPositions(scanned, "CREATE PROC")) > 0:
		detected = "procedures"
	case len(findAllKeywordPositions(scanned, "CREATE TABLE")) > 0:
		detected = "tables"
	case len(findCreateIndexStatementPositions(scanned)) > 0:
		detected = "indexes"
	}
	if detected == "" || detected == actualCategory {
		return
	}
	report.AddFail(output.Item{
		Rule:       "DB_OBJECT_WRONG_CATEGORY",
		Severity:   3,
		File:       relFile,
		Object:     actualCategory,
		Message:    fmt.Sprintf("SQL defines a %s object but file is stored under %s", detected, actualCategory),
		Suggestion: fmt.Sprintf("move the file to db_object/%s/ and keep the dbo.<Object>.sql filename contract", detected),
		Reason:     "INVALID_SOURCE_LAYOUT",
		Action:     "FAIL",
	})
}

func listCategorySQLFiles(repoPath func(string) string, listFiles func(string) ([]string, error), relDir string) (map[string]string, error) {
	files, err := listFiles(repoPath(relDir))
	if err != nil {
		if os.IsNotExist(err) {
			return map[string]string{}, nil
		}
		return nil, err
	}

	out := make(map[string]string)
	for _, file := range files {
		if !strings.EqualFold(filepath.Ext(file), ".sql") {
			continue
		}
		base := strings.ToLower(filepath.Base(file))
		if base == "" {
			continue
		}
		out[base] = file
	}
	return out, nil
}

// ValidateProgrammabilityHeaders validates source and migration programmability header contracts.
func ValidateProgrammabilityHeaders(report *output.Report, opts ProgrammabilityHeadersOptions) error {
	readFile := opts.ReadFile
	if readFile == nil {
		readFile = os.ReadFile
	}
	existingProgrammabilityObjects := collectExistingProgrammabilityObjects(opts.Metas)

	programmabilityCategories := []struct {
		Category string
		Rule     string
	}{
		{Category: "views", Rule: "INVALID_VIEW_HEADER"},
		{Category: "functions", Rule: "INVALID_FUNCTION_HEADER"},
		{Category: "procedures", Rule: "INVALID_PROCEDURE_HEADER"},
	}

	listFiles := opts.ListFilesRecursive
	if listFiles == nil {
		listFiles = fsx.ListFilesRecursive
	}

	for _, item := range programmabilityCategories {
		files, err := fsx.ResolveDbObjectCategoryFiles(opts.RepoPath, listFiles, opts.DbObjectDir, item.Category)
		if err != nil {
			return err
		}
		for _, file := range files {
			if !strings.EqualFold(filepath.Ext(file), ".sql") {
				continue
			}
			raw, err := readFile(file)
			if err != nil {
				return err
			}
			sqlText := string(raw)
			baseName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
			if shouldValidateProgrammabilitySourceContract(opts, existingProgrammabilityObjects, item.Category, baseName) {
				ValidateProgrammabilitySingleObjectContract(report, opts.ToRelPath(file), sqlText, item.Category)
			}
			if !hasCanonicalProgrammabilityHeader(sqlText, item.Category) {
				report.AddFail(output.Item{
					Rule:     item.Rule,
					Severity: 3,
					File:     opts.ToRelPath(file),
					Message:  "programmability header must use CREATE OR ALTER",
					Reason:   "REQUIRES_MANUAL",
					Action:   "FAIL",
				})
			}
		}
	}

	if opts.SkipMigrationHeaders {
		return nil
	}
	for _, meta := range opts.Metas {
		if meta.IsCore {
			continue
		}
		if meta.Direction != "up" {
			continue
		}
		if meta.Category != "views" && meta.Category != "functions" && meta.Category != "procedures" {
			continue
		}
		filePath := opts.RepoPath(meta.RelPath)
		raw, err := readFile(filePath)
		if err != nil {
			return err
		}
		sqlText := string(raw)
		if !hasCanonicalProgrammabilityHeader(sqlText, meta.Category) &&
			!hasCanonicalProgrammabilityDrop(sqlText, meta.Category, meta.Object) {
			report.AddFail(output.Item{
				Rule:     "INVALID_PROGRAMMABILITY_MIGRATION_HEADER",
				Severity: 3,
				File:     meta.RelPath,
				Message:  "programmability migration must use CREATE OR ALTER or a matching DROP ... IF EXISTS removal",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
	}

	return nil
}

func hasCanonicalProgrammabilityDrop(sqlText, category, object string) bool {
	upper := strings.ToUpper(sqlparse.StripSQLStringsAndComments(sqlText))
	upper = normalizeIdentifierBracketsUpper(upper)
	upper = strings.Join(strings.Fields(upper), " ")
	kind := map[string]string{
		"views":      "VIEW",
		"functions":  "FUNCTION",
		"procedures": "PROCEDURE",
	}[category]
	target := strings.ToUpper(strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(object, "[", ""), "]", "")))
	if kind == "" || target == "" {
		return false
	}
	return strings.Contains(upper, "DROP "+kind+" IF EXISTS "+target+";") ||
		strings.HasSuffix(upper, "DROP "+kind+" IF EXISTS "+target)
}

func collectExistingProgrammabilityObjects(metas []MigrationMeta) map[string]struct{} {
	out := make(map[string]struct{})
	for _, meta := range metas {
		if meta.IsCore || meta.Direction != "up" {
			continue
		}
		if meta.Category != "views" && meta.Category != "functions" && meta.Category != "procedures" {
			continue
		}
		out[programmabilityObjectKey(meta.Category, meta.Object)] = struct{}{}
	}
	return out
}

func shouldValidateProgrammabilitySourceContract(opts ProgrammabilityHeadersOptions, existing map[string]struct{}, category, object string) bool {
	if !opts.EnableSourceContractForNewOnly {
		return false
	}
	_, exists := existing[programmabilityObjectKey(category, object)]
	return !exists
}

func programmabilityObjectKey(category, object string) string {
	return strings.ToLower(strings.TrimSpace(category)) + "|" + normalizeIdentifierBracketsUpper(strings.TrimSpace(object))
}

// ValidateTableSnapshotContract validates table source snapshot contract.
func ValidateTableSnapshotContract(report *output.Report, relFile, sql string) {
	scanned := strings.ToUpper(sqlparse.StripSQLStringsAndComments(sql))

	if len(findAllKeywordPositions(scanned, "CREATE TABLE")) != 1 {
		report.AddFail(output.Item{
			Rule:     "TABLE_SNAPSHOT_FORBIDDEN_STATEMENT",
			Severity: 3,
			File:     relFile,
			Message:  "table snapshot must contain exactly one CREATE TABLE statement",
			Reason:   "REQUIRES_MANUAL",
			Action:   "FAIL",
		})
	}

	for _, token := range []string{"ALTER TABLE", "DROP TABLE", "CREATE INDEX", "CREATE UNIQUE INDEX"} {
		if strings.Contains(scanned, token) {
			report.AddFail(output.Item{
				Rule:     "TABLE_SNAPSHOT_FORBIDDEN_STATEMENT",
				Severity: 3,
				File:     relFile,
				Message:  "table snapshot contains forbidden statement: " + token,
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
	}

	if strings.Contains(scanned, "IF OBJECT_ID(") || strings.Contains(scanned, "IF EXISTS(") || strings.Contains(scanned, "IF NOT EXISTS(") {
		report.AddFail(output.Item{
			Rule:     "TABLE_SNAPSHOT_FORBIDDEN_STATEMENT",
			Severity: 3,
			File:     relFile,
			Message:  "table snapshot must not use IF wrapper logic",
			Reason:   "REQUIRES_MANUAL",
			Action:   "FAIL",
		})
	}
}

// ValidateIndexSnapshotContract validates index source snapshot contract.
func ValidateIndexSnapshotContract(report *output.Report, relFile, sql, object string) {
	normalizedSQL := sqlfmt.NormalizeSQLText(sql, true, true)
	scanned := strings.ToUpper(sqlparse.StripSQLStringsAndComments(normalizedSQL))
	if strings.Contains(scanned, "DROP INDEX") {
		report.AddFail(output.Item{
			Rule:     "INDEX_SOURCE_CONTRACT_VIOLATION",
			Severity: 3,
			File:     relFile,
			Message:  "index source file must not contain DROP INDEX",
			Reason:   "REQUIRES_MANUAL",
			Action:   "FAIL",
		})
	}

	statements := extractAllCreateIndexStatements(normalizedSQL)
	if len(statements) == 0 {
		report.AddFail(output.Item{
			Rule:     "INDEX_SOURCE_CONTRACT_VIOLATION",
			Severity: 3,
			File:     relFile,
			Message:  "index source file must contain CREATE INDEX statements only",
			Reason:   "REQUIRES_MANUAL",
			Action:   "FAIL",
		})
		return
	}

	normalizedObject := strings.ToUpper(object)
	for _, stmt := range statements {
		if !indexStatementTargetsObject(stmt, normalizedObject) {
			report.AddFail(output.Item{
				Rule:     "INDEX_SOURCE_TABLE_MISMATCH",
				Severity: 3,
				File:     relFile,
				Object:   object,
				Message:  "index statement targets table outside filename contract",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
	}

	ranges := extractAllCreateIndexStatementRanges(normalizedSQL)
	if hasNonWhitespaceSQLOutsideRanges(normalizedSQL, ranges) {
		report.AddFail(output.Item{
			Rule:     "INDEX_SOURCE_CONTRACT_VIOLATION",
			Severity: 3,
			File:     relFile,
			Message:  "index source file must contain only CREATE [UNIQUE] INDEX statements",
			Reason:   "REQUIRES_MANUAL",
			Action:   "FAIL",
		})
	}

	if sqlparse.ContainsStandaloneWord(scanned, "CREATE") {
		allowedCreate := len(findCreateIndexStatementPositions(scanned))
		totalCreate := len(findAllKeywordPositions(scanned, "CREATE"))
		if totalCreate > allowedCreate {
			report.AddFail(output.Item{
				Rule:     "INDEX_SOURCE_CONTRACT_VIOLATION",
				Severity: 3,
				File:     relFile,
				Message:  "index source file must not contain CREATE statement besides CREATE [UNIQUE] INDEX",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
	}
}

func indexStatementTargetsObject(stmt, object string) bool {
	canonical := normalizeForSemanticContains(normalizeIdentifierBracketsUpper(stmt))
	if containsIndexTarget(canonical, object) {
		return true
	}
	if !strings.HasPrefix(object, "DBO.") {
		return false
	}
	shortName := strings.TrimPrefix(object, "DBO.")
	return containsIndexTarget(canonical, shortName)
}

func containsIndexTarget(canonical, target string) bool {
	needle := " ON " + target
	start := 0
	for {
		rel := strings.Index(canonical[start:], needle)
		if rel < 0 {
			return false
		}
		end := start + rel + len(needle)
		if end == len(canonical) || canonical[end] == '(' || canonical[end] == ' ' {
			return true
		}
		start = end
	}
}

// ValidateProgrammabilitySingleObjectContract validates single-object programmability source contract.
func ValidateProgrammabilitySingleObjectContract(report *output.Report, relFile, sql, category string) {
	scanned := strings.ToUpper(sqlparse.StripSQLStringsAndComments(sql))
	viewCount := len(findAllKeywordPositions(scanned, "CREATE OR ALTER VIEW"))
	functionCount := len(findAllKeywordPositions(scanned, "CREATE OR ALTER FUNCTION"))
	procedureCount := len(findAllKeywordPositions(scanned, "CREATE OR ALTER PROCEDURE"))
	total := viewCount + functionCount + procedureCount
	if total != 1 {
		report.AddFail(output.Item{
			Rule:     "PROGRAMMABILITY_MULTI_OBJECT",
			Severity: 3,
			File:     relFile,
			Object:   category,
			Message:  "programmability file must define exactly one object",
			Reason:   "REQUIRES_MANUAL",
			Action:   "FAIL",
		})
	}

	for _, token := range []string{"CREATE TABLE", "ALTER TABLE", "DROP TABLE"} {
		if hasPersistentTableDDL(scanned, token) {
			addProgrammabilityContractDDLFail(report, relFile, category, token)
		}
	}
	if hasPersistentCreateIndexDDL(scanned) {
		addProgrammabilityContractDDLFail(report, relFile, category, "CREATE INDEX")
	}
	for _, token := range []string{"ALTER INDEX", "DROP INDEX"} {
		if hasPersistentIndexDDL(scanned, token) {
			addProgrammabilityContractDDLFail(report, relFile, category, token)
		}
	}
	for _, token := range []string{"DROP VIEW", "DROP FUNCTION", "DROP PROCEDURE"} {
		if strings.Contains(scanned, token) {
			addProgrammabilityContractDDLFail(report, relFile, category, token)
		}
	}
}

func addProgrammabilityContractDDLFail(report *output.Report, relFile, category, token string) {
	report.AddFail(output.Item{
		Rule:     "PROGRAMMABILITY_SOURCE_CONTRACT_VIOLATION",
		Severity: 3,
		File:     relFile,
		Object:   category,
		Message:  "programmability file must not contain non-programmability DDL: " + token,
		Reason:   "REQUIRES_MANUAL",
		Action:   "FAIL",
	})
}

func hasPersistentTableDDL(scannedUpper, keyword string) bool {
	for _, pos := range findAllKeywordPositions(scannedUpper, keyword) {
		end := sqlparse.FindStatementEnd(scannedUpper, pos)
		if end <= pos {
			end = len(scannedUpper)
		}
		stmt := scannedUpper[pos:end]
		target := extractTableTargetToken(stmt, keyword)
		if !isTempTableTargetToken(target) {
			return true
		}
	}
	return false
}

func hasPersistentCreateIndexDDL(scannedUpper string) bool {
	for _, pos := range findCreateIndexStatementPositions(scannedUpper) {
		end := sqlparse.FindStatementEnd(scannedUpper, pos)
		if end <= pos {
			end = len(scannedUpper)
		}
		stmt := scannedUpper[pos:end]
		target := extractIndexOnTargetToken(stmt)
		if !isTempTableTargetToken(target) {
			return true
		}
	}
	return false
}

func hasPersistentIndexDDL(scannedUpper, keyword string) bool {
	for _, pos := range findAllKeywordPositions(scannedUpper, keyword) {
		end := sqlparse.FindStatementEnd(scannedUpper, pos)
		if end <= pos {
			end = len(scannedUpper)
		}
		stmt := scannedUpper[pos:end]
		target := ""
		if keyword == "DROP INDEX" {
			target = extractDropIndexTargetToken(stmt)
			if target == "" {
				target = extractIndexOnTargetToken(stmt)
			}
		} else {
			target = extractIndexOnTargetToken(stmt)
		}
		if !isTempTableTargetToken(target) {
			return true
		}
	}
	return false
}

func extractTableTargetToken(statementUpper, keyword string) string {
	if len(statementUpper) < len(keyword) {
		return ""
	}
	rest := strings.TrimSpace(statementUpper[len(keyword):])
	restUpper := strings.ToUpper(rest)
	if strings.HasPrefix(restUpper, "IF EXISTS") {
		rest = strings.TrimSpace(rest[len("IF EXISTS"):])
		restUpper = strings.ToUpper(rest)
	}
	if strings.HasPrefix(restUpper, "IF NOT EXISTS") {
		rest = strings.TrimSpace(rest[len("IF NOT EXISTS"):])
	}
	return firstSQLToken(rest)
}

func extractIndexOnTargetToken(statementUpper string) string {
	fields := strings.Fields(statementUpper)
	for i := 0; i+1 < len(fields); i++ {
		if fields[i] != "ON" {
			continue
		}
		return firstSQLToken(fields[i+1])
	}
	return ""
}

func extractDropIndexTargetToken(statementUpper string) string {
	const keyword = "DROP INDEX"
	if len(statementUpper) < len(keyword) {
		return ""
	}
	rest := strings.TrimSpace(statementUpper[len(keyword):])
	restUpper := strings.ToUpper(rest)
	if strings.HasPrefix(restUpper, "IF EXISTS") {
		rest = strings.TrimSpace(rest[len("IF EXISTS"):])
	}

	token := firstSQLToken(rest)
	token = strings.TrimSpace(strings.Trim(strings.Trim(token, ";"), ","))
	if token == "" {
		return ""
	}

	token = strings.ReplaceAll(strings.ReplaceAll(token, "[", ""), "]", "")
	token = strings.ReplaceAll(token, "\"", "")
	if !strings.Contains(token, ".") {
		return ""
	}

	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return ""
	}
	return strings.Join(parts[:len(parts)-1], ".")
}

func isTempTableTargetToken(token string) bool {
	target := strings.TrimSpace(token)
	target = strings.Trim(target, ";")
	target = strings.Trim(target, ",")
	if target == "" {
		return false
	}
	target = strings.ReplaceAll(strings.ReplaceAll(target, "[", ""), "]", "")
	target = strings.ReplaceAll(target, "\"", "")
	targetUpper := strings.ToUpper(target)
	if strings.HasPrefix(targetUpper, "#") {
		return true
	}
	if strings.HasPrefix(targetUpper, "TEMPDB..#") || strings.Contains(targetUpper, "..#") || strings.Contains(targetUpper, ".#") {
		return true
	}
	return false
}

func findAllKeywordPositions(inputUpper, keywordUpper string) []int {
	if keywordUpper == "" || len(inputUpper) < len(keywordUpper) {
		return []int{}
	}
	pos := make([]int, 0)
	limit := len(inputUpper) - len(keywordUpper)
	for i := 0; i <= limit; i++ {
		if inputUpper[i:i+len(keywordUpper)] != keywordUpper {
			continue
		}
		leftOK := i == 0 || !isWordChar(rune(inputUpper[i-1]))
		rightIdx := i + len(keywordUpper)
		rightOK := rightIdx >= len(inputUpper) || !isWordChar(rune(inputUpper[rightIdx]))
		if leftOK && rightOK {
			pos = append(pos, i)
		}
	}
	return pos
}

func extractAllCreateIndexStatements(sql string) []string {
	normalized := sqlfmt.NormalizeSQLText(sql, true, true)
	upper := strings.ToUpper(normalized)
	positions := findCreateIndexStatementPositions(upper)

	out := make([]string, 0, len(positions))
	seen := map[int]struct{}{}
	for _, pos := range positions {
		if _, ok := seen[pos]; ok {
			continue
		}
		seen[pos] = struct{}{}
		end := sqlparse.FindStatementEnd(normalized, pos)
		stmt := strings.TrimSpace(normalized[pos:end])
		if stmt == "" {
			continue
		}
		out = append(out, stmt)
	}
	return out
}

type sqlRange struct {
	Start int
	End   int
}

func extractAllCreateIndexStatementRanges(sql string) []sqlRange {
	normalized := sqlfmt.NormalizeSQLText(sql, true, true)
	upper := strings.ToUpper(normalized)
	positions := findCreateIndexStatementPositions(upper)

	ranges := make([]sqlRange, 0, len(positions))
	seen := map[int]struct{}{}
	for _, pos := range positions {
		if _, ok := seen[pos]; ok {
			continue
		}
		seen[pos] = struct{}{}
		end := sqlparse.FindStatementEnd(normalized, pos)
		ranges = append(ranges, sqlRange{Start: pos, End: end})
	}
	return ranges
}

func findCreateIndexStatementPositions(inputUpper string) []int {
	positions := make([]int, 0)
	seen := make(map[int]struct{})

	for i := 0; i < len(inputUpper); i++ {
		if !hasKeywordAtWordBoundary(inputUpper, i, "CREATE") {
			continue
		}

		next, tok := readNextWordUpper(inputUpper, i+len("CREATE"))
		if tok == "" {
			continue
		}

		matched := false
		switch tok {
		case "INDEX":
			matched = true
		case "UNIQUE":
			next2, tok2 := readNextWordUpper(inputUpper, next)
			switch tok2 {
			case "INDEX":
				matched = true
			case "NONCLUSTERED", "CLUSTERED":
				_, tok3 := readNextWordUpper(inputUpper, next2)
				if tok3 == "INDEX" {
					matched = true
				}
			}
		case "NONCLUSTERED", "CLUSTERED":
			next2, tok2 := readNextWordUpper(inputUpper, next)
			if tok2 == "INDEX" {
				matched = true
			} else if tok2 == "COLUMNSTORE" {
				_, tok3 := readNextWordUpper(inputUpper, next2)
				matched = tok3 == "INDEX"
			}
		case "COLUMNSTORE", "XML", "SPATIAL":
			_, tok2 := readNextWordUpper(inputUpper, next)
			matched = tok2 == "INDEX"
		case "PRIMARY":
			next2, tok2 := readNextWordUpper(inputUpper, next)
			if tok2 == "XML" {
				_, tok3 := readNextWordUpper(inputUpper, next2)
				matched = tok3 == "INDEX"
			}
		}
		if matched {
			if _, ok := seen[i]; ok {
				continue
			}
			seen[i] = struct{}{}
			positions = append(positions, i)
		}
	}
	sort.Ints(positions)
	return positions
}

func hasKeywordAtWordBoundary(inputUpper string, pos int, keyword string) bool {
	if pos < 0 || pos+len(keyword) > len(inputUpper) {
		return false
	}
	if inputUpper[pos:pos+len(keyword)] != keyword {
		return false
	}
	leftOK := pos == 0 || !isWordChar(rune(inputUpper[pos-1]))
	rightIdx := pos + len(keyword)
	rightOK := rightIdx >= len(inputUpper) || !isWordChar(rune(inputUpper[rightIdx]))
	return leftOK && rightOK
}

func readNextWordUpper(inputUpper string, start int) (int, string) {
	i := start
	for i < len(inputUpper) {
		switch inputUpper[i] {
		case ' ', '\t', '\r', '\n':
			i++
		default:
			goto tokenStart
		}
	}
	return i, ""

tokenStart:
	j := i
	for j < len(inputUpper) && isWordChar(rune(inputUpper[j])) {
		j++
	}
	if j == i {
		return i + 1, ""
	}
	return j, inputUpper[i:j]
}

func hasCanonicalProgrammabilityHeader(sqlText, category string) bool {
	if dynamicBatch, ok := extractSpExecuteSQLBatch(sqlText); ok {
		sqlText = dynamicBatch
	}
	upper := strings.ToUpper(sqlparse.StripSQLStringsAndComments(sqlText))
	upper = normalizeIdentifierBracketsUpper(upper)
	upper = strings.Join(strings.Fields(upper), " ")
	required := map[string]string{
		"views":      "CREATE OR ALTER VIEW DBO.",
		"functions":  "CREATE OR ALTER FUNCTION DBO.",
		"procedures": "CREATE OR ALTER PROCEDURE DBO.",
	}[category]
	if required == "" {
		return false
	}
	return strings.Contains(upper, required)
}

func extractSpExecuteSQLBatch(sqlText string) (string, bool) {
	upper := strings.ToUpper(sqlText)
	marker := "EXEC SYS.SP_EXECUTESQL N'"
	pos := strings.Index(upper, marker)
	if pos < 0 {
		return "", false
	}
	cursor := pos + len(marker)
	var decoded strings.Builder
	for cursor < len(sqlText) {
		if sqlText[cursor] != '\'' {
			decoded.WriteByte(sqlText[cursor])
			cursor++
			continue
		}
		if cursor+1 < len(sqlText) && sqlText[cursor+1] == '\'' {
			decoded.WriteByte('\'')
			cursor += 2
			continue
		}
		return strings.TrimSpace(decoded.String()), decoded.Len() > 0
	}
	return "", false
}

func hasNonWhitespaceSQLOutsideRanges(sql string, ranges []sqlRange) bool {
	normalized := sqlfmt.NormalizeSQLText(sql, true, true)
	if len(ranges) == 0 {
		return strings.TrimSpace(sqlparse.StripSQLStringsAndComments(normalized)) != ""
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	cursor := 0
	for _, rg := range ranges {
		if rg.Start > len(normalized) {
			break
		}
		if rg.Start > cursor {
			segment := normalized[cursor:rg.Start]
			if strings.TrimSpace(sqlparse.StripSQLStringsAndComments(segment)) != "" {
				return true
			}
		}
		if rg.End > cursor {
			cursor = rg.End
		}
		if cursor > len(normalized) {
			cursor = len(normalized)
		}
	}
	if cursor < len(normalized) {
		segment := normalized[cursor:]
		if strings.TrimSpace(sqlparse.StripSQLStringsAndComments(segment)) != "" {
			return true
		}
	}
	return false
}

func normalizeIdentifierBracketsUpper(input string) string {
	withoutBrackets := strings.ReplaceAll(strings.ReplaceAll(input, "[", ""), "]", "")
	return strings.ToUpper(withoutBrackets)
}

func normalizeForSemanticContains(input string) string {
	upper := strings.ToUpper(input)
	parts := strings.Fields(upper)
	if len(parts) == 0 {
		return ""
	}
	return strings.Join(parts, " ")
}

func firstSQLToken(input string) string {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return ""
	}

	for idx, ch := range trimmed {
		if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' || ch == '(' || ch == ')' || ch == ',' {
			return trimmed[:idx]
		}
	}
	return trimmed
}

func isWordChar(ch rune) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') || ch == '_'
}
