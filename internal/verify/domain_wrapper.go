package verify

import (
	"crypto/sha256"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"

	"rza_migrate_db/internal/output"
	"rza_migrate_db/internal/sqlparse"
)

func sqlWithoutMigrationBody(sqlText string) string {
	const beginMarker = "-- MIGRATION_BODY_BEGIN"
	const endMarker = "-- MIGRATION_BODY_END"
	begin := strings.Index(sqlText, beginMarker)
	if begin < 0 {
		return sqlText
	}
	bodyStart := begin + len(beginMarker)
	relEnd := strings.Index(sqlText[bodyStart:], endMarker)
	if relEnd < 0 {
		return sqlText
	}
	end := bodyStart + relEnd
	return sqlText[:bodyStart] + "\n" + sqlText[end:]
}

// WrapperSubsetOptions holds runtime context for wrapper subset validation.
type WrapperSubsetOptions struct {
	RepoPath           func(string) string
	Metas              []MigrationMeta
	AllowedDecisions   map[string]struct{}
	AllowedActions     map[string]struct{}
	AllowedReasons     map[string]struct{}
	AllowedSeverities  map[int]struct{}
	StrictMode         bool
	EnableParallelRead bool
	ParallelWorkers    int
	ReadFile           func(string) ([]byte, error)
}

// ValidateWrapperSubset validates the generated wrapper contract for every
// migration category except the reserved core bootstrap files.
func ValidateWrapperSubset(report *output.Report, opts WrapperSubsetOptions) error {
	if opts.EnableParallelRead && len(opts.Metas) > 1 {
		return validateWrapperSubsetParallel(report, opts)
	}
	readFile := opts.ReadFile
	if readFile == nil {
		readFile = os.ReadFile
	}
	for _, meta := range opts.Metas {
		if meta.IsCore {
			continue
		}
		if meta.Category == "procedures" &&
			(strings.EqualFold(meta.Object, "dbo.usp_mig_log_event") ||
				strings.EqualFold(meta.Object, "dbo.usp_mig_try_get_lock")) {
			continue
		}
		if meta.Category != "types" && meta.Category != "tables" && meta.Category != "indexes" &&
			meta.Category != "views" && meta.Category != "functions" && meta.Category != "procedures" &&
			meta.Category != "patch" {
			continue
		}
		if (meta.Category == "views" || meta.Category == "functions" || meta.Category == "procedures") &&
			!isHardenedGeneratedMigration(meta.RelPath) {
			continue
		}

		raw, err := readFile(opts.RepoPath(meta.RelPath))
		if err != nil {
			return err
		}
		sqlText := string(raw)
		wrapperSQLText := sqlWithoutMigrationBody(sqlText)
		upper := strings.ToUpper(wrapperSQLText)
		semanticUpper := strings.ToUpper(sqlparse.StripSQLStringsAndComments(wrapperSQLText))
		hasFileName := strings.Contains(semanticUpper, "DECLARE @FILE_NAME")
		hasLegacyFile := strings.Contains(semanticUpper, "DECLARE @FILE ")
		hasNamedFileArg := strings.Contains(semanticUpper, "@FILE_NAME =")
		hasStrictAuditStyle := hasFileName || hasNamedFileArg

		if meta.Category == "patch" && !hasFileName && !hasLegacyFile {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_CONTRACT_MISSING",
				Severity: 3,
				File:     meta.RelPath,
				Message:  "patch migration must declare @file or @file_name",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
			continue
		}

		if !strings.Contains(upper, "SEVERITY MAPPING") {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_SEVERITY_HEADER_MISSING",
				Severity: 3,
				File:     meta.RelPath,
				Message:  "wrapper must include severity mapping header comment",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}

		if !strings.Contains(semanticUpper, "DECLARE @RUN_KEY") {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_CONTRACT_MISSING",
				Severity: 3,
				File:     meta.RelPath,
				Message:  "wrapper must declare @run_key",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
		if !strings.Contains(sqlparse.CompactUpperSQL(semanticUpper), "DECLARE@RUN_KEYUNIQUEIDENTIFIER=NEWID();") {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_RUN_KEY_INVALID",
				Severity: 3,
				File:     meta.RelPath,
				Message:  "wrapper must initialize @run_key using NEWID()",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
		if !strings.Contains(semanticUpper, "DECLARE @VERSION") {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_CONTRACT_MISSING",
				Severity: 3,
				File:     meta.RelPath,
				Message:  "wrapper must declare @version",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
		if !strings.Contains(semanticUpper, "DECLARE @CATEGORY") {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_CONTRACT_MISSING",
				Severity: 3,
				File:     meta.RelPath,
				Message:  "wrapper must declare @category",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
		if !strings.Contains(semanticUpper, "DECLARE @FILE ") && !strings.Contains(semanticUpper, "DECLARE @FILE_NAME ") {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_CONTRACT_MISSING",
				Severity: 3,
				File:     meta.RelPath,
				Message:  "wrapper must declare @file or @file_name",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
		ValidateWrapperDeterministicMetadata(report, meta, sqlText, semanticUpper)
		if isHardenedGeneratedMigration(meta.RelPath) {
			validateWrapperContentHash(report, meta.RelPath, sqlText)
			for _, required := range []string{
				"IF OBJECT_ID(N'DBO.USP_MIG_LOG_EVENT', N'P') IS NULL",
				"MIGRATION AUDIT INFRASTRUCTURE IS UNAVAILABLE",
				"DECLARE @OPERATION_ID",
				"DECLARE @CONTENT_SHA256",
				"DECLARE @STARTED_AT",
			} {
				if strings.Contains(upper, required) {
					continue
				}
				report.AddFail(output.Item{
					Rule:     "WRAPPER_MONITORING_CONTRACT_MISSING",
					Severity: 3,
					File:     meta.RelPath,
					Message:  "wrapper monitoring contract missing: " + required,
					Reason:   "REQUIRES_MANUAL",
					Action:   "FAIL",
				})
			}
		}
		if !strings.Contains(semanticUpper, "BEGIN TRY") || !strings.Contains(semanticUpper, "BEGIN CATCH") {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_CONTRACT_MISSING",
				Severity: 3,
				File:     meta.RelPath,
				Message:  "wrapper must have BEGIN TRY/BEGIN CATCH blocks",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
		if !strings.Contains(semanticUpper, "BEGIN TRAN") {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_CONTRACT_MISSING",
				Severity: 3,
				File:     meta.RelPath,
				Message:  "wrapper must open transaction",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
		if !strings.Contains(semanticUpper, "COMMIT") {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_CONTRACT_MISSING",
				Severity: 3,
				File:     meta.RelPath,
				Message:  "wrapper must commit transaction",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
		if !strings.Contains(semanticUpper, "ROLLBACK") {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_CONTRACT_MISSING",
				Severity: 3,
				File:     meta.RelPath,
				Message:  "wrapper must rollback on CATCH",
				Reason:   "SQL_ERROR",
				Action:   "FAIL",
			})
		}
		if !hasSemanticThrowInCatch(wrapperSQLText) {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_THROW_REQUIRED",
				Severity: 3,
				File:     meta.RelPath,
				Message:  "tables/indexes migration must throw on CATCH",
				Reason:   "SQL_ERROR",
				Action:   "FAIL",
			})
		}

		if hasStrictAuditStyle {
			missing := missingAuditContractTokens(upper)
			for _, token := range missing {
				report.AddFail(output.Item{
					Rule:     "WRAPPER_AUDIT_CONTRACT_MISSING",
					Severity: 3,
					File:     meta.RelPath,
					Message:  "wrapper audit contract missing: " + token,
					Reason:   "REQUIRES_MANUAL",
					Action:   "FAIL",
				})
			}
			ValidateWrapperAuditWhitelist(report, meta.RelPath, upper, opts.AllowedDecisions, opts.AllowedActions, opts.AllowedReasons, opts.AllowedSeverities)
		} else if hasLegacyFile {
			missing := missingAuditContractTokens(upper)
			if len(missing) > 0 {
				report.AddFail(output.Item{
					Rule:     "WRAPPER_AUDIT_LEGACY_WEAK",
					Severity: 3,
					File:     meta.RelPath,
					Message:  "legacy wrapper audit is incomplete (strict-only): " + strings.Join(missing, ", "),
					Reason:   "REQUIRES_MANUAL",
					Action:   "FAIL",
				})
			}
		}

		validatePrecheckGuardCoverage(report, meta.RelPath, meta.Category, sqlText, opts.StrictMode)
	}

	return nil
}

func validateWrapperContentHash(report *output.Report, relPath, sqlText string) {
	const (
		declaration = "DECLARE @content_sha256 CHAR(64) = N'"
		beginMarker = "-- MIGRATION_BODY_BEGIN"
		endMarker   = "-- MIGRATION_BODY_END"
	)
	lower := strings.ToLower(sqlText)
	declPos := strings.Index(lower, strings.ToLower(declaration))
	if declPos < 0 {
		report.AddFail(output.Item{Rule: "MIGRATION_CONTENT_HASH_MISSING", Severity: 3, File: relPath, Message: "migration content hash declaration is missing", Reason: "REQUIRES_MANUAL", Action: "FAIL"})
		return
	}
	hashStart := declPos + len(declaration)
	hashEndRel := strings.Index(sqlText[hashStart:], "'")
	if hashEndRel != 64 {
		report.AddFail(output.Item{Rule: "MIGRATION_CONTENT_HASH_INVALID", Severity: 3, File: relPath, Message: "migration content hash must be exactly 64 hexadecimal characters", Reason: "REQUIRES_MANUAL", Action: "FAIL"})
		return
	}
	declared := strings.ToLower(sqlText[hashStart : hashStart+hashEndRel])
	for _, ch := range declared {
		if !strings.ContainsRune("0123456789abcdef", ch) {
			report.AddFail(output.Item{Rule: "MIGRATION_CONTENT_HASH_INVALID", Severity: 3, File: relPath, Message: "migration content hash contains a non-hexadecimal character", Reason: "REQUIRES_MANUAL", Action: "FAIL"})
			return
		}
	}
	begin := strings.Index(sqlText, beginMarker)
	end := strings.Index(sqlText, endMarker)
	if begin < 0 || end <= begin {
		report.AddFail(output.Item{Rule: "MIGRATION_CONTENT_HASH_INVALID", Severity: 3, File: relPath, Message: "migration body hash markers are missing or out of order", Reason: "REQUIRES_MANUAL", Action: "FAIL"})
		return
	}
	body := canonicalHashedMigrationBody(sqlText[begin+len(beginMarker) : end])
	sum := sha256.Sum256([]byte(body))
	actual := fmt.Sprintf("%x", sum)
	if actual != declared {
		report.AddFail(output.Item{
			Rule:       "MIGRATION_CONTENT_HASH_MISMATCH",
			Severity:   3,
			File:       relPath,
			Message:    fmt.Sprintf("migration body checksum mismatch: declared=%s actual=%s", declared, actual),
			Suggestion: "do not rewrite the checksum header; restore the reviewed file. If it has never been applied, remove the complete unapproved pair and regenerate it from db_object/patch, preferably through draft then seal",
			Reason:     "REQUIRES_MANUAL",
			Action:     "FAIL",
		})
	}
}

// canonicalHashedMigrationBody keeps executable text byte-stable while
// tolerating CRLF/LF conversion by source control or deployment.
func canonicalHashedMigrationBody(body string) string {
	body = strings.ReplaceAll(body, "\r\n", "\n")
	return strings.TrimSpace(body)
}

func validateWrapperSubsetParallel(report *output.Report, opts WrapperSubsetOptions) error {
	type result struct {
		report output.Report
		err    error
	}
	results := make([]result, len(opts.Metas))
	workers := opts.ParallelWorkers
	if workers < 1 {
		workers = 1
	}
	if cpu := runtime.GOMAXPROCS(0); workers > cpu {
		workers = cpu
	}
	if workers > len(opts.Metas) {
		workers = len(opts.Metas)
	}
	jobs := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for index := range jobs {
				sub := output.NewReport("", "")
				single := opts
				single.Metas = []MigrationMeta{opts.Metas[index]}
				single.EnableParallelRead = false
				results[index] = result{report: sub, err: ValidateWrapperSubset(&sub, single)}
			}
		}()
	}
	for index := range opts.Metas {
		jobs <- index
	}
	close(jobs)
	wg.Wait()
	for _, item := range results {
		mergeReport(report, item.report)
		if item.err != nil {
			return item.err
		}
	}
	return nil
}

func validatePrecheckGuardCoverage(report *output.Report, relPath, category, sqlText string, strict bool) {
	if isLegacyNoPrecheckMigration(sqlText) {
		return
	}
	strictForFile := strict
	upper := strings.ToUpper(sqlparse.StripSQLStringsAndComments(sqlText))
	switch category {
	case "tables":
		validateTablePrecheckGuards(report, relPath, upper, strictForFile)
	case "indexes":
		validateIndexPrecheckGuards(report, relPath, upper, strictForFile)
	}
	validateHardenedGeneratedPrechecks(report, relPath, category, sqlText, upper, strictForFile)
}

// Historical migrations are immutable. Strict generated-precheck enforcement
// starts after the last acknowledged legacy batch; every newly generated file
// is validated against the hardened contract.
const hardenedPrecheckVersion = "20260726000000000"
const bootstrapCompatibleTablePrecheckVersion = "20260728000000000"

func validateHardenedGeneratedPrechecks(report *output.Report, relPath, category, rawSQL, semanticUpper string, strict bool) {
	if !isHardenedGeneratedMigration(relPath) {
		return
	}
	rawUpper := strings.ToUpper(rawSQL)
	require := func(hasDDL bool, marker, message string) {
		if hasDDL && !strings.Contains(rawUpper, marker) {
			addHardenedPrecheckFinding(report, relPath, message, strict)
		}
	}

	switch category {
	case "tables":
		hasAlterAdd := strings.Contains(semanticUpper, "ALTER TABLE ") && strings.Contains(semanticUpper, " ADD ")
		hasAlterColumn := strings.Contains(semanticUpper, "ALTER TABLE ") && strings.Contains(semanticUpper, " ALTER COLUMN ")
		require(hasAlterAdd || hasAlterColumn, "TARGET TABLE DOES NOT EXIST", "generated table ADD/ALTER must fail fast when target table is missing")
		require(strings.Contains(semanticUpper, " DROP COLUMN "), "PRECHECK BLOCKED: DROP COLUMN TARGET CONTAINS DATA", "generated DROP COLUMN must block destructive rollback when data exists")
		require(strings.Contains(semanticUpper, "DROP TABLE "), "PRECHECK BLOCKED: DROP TABLE TARGET CONTAINS DATA", "generated DROP TABLE must block destructive rollback when data exists")
		validateBootstrapCompatibleCreateTableGuard(report, relPath, rawUpper, semanticUpper, strict)
	case "indexes":
		hasCreate := containsCreateIndexDDL(semanticUpper)
		require(hasCreate, "CREATE INDEX TARGET TABLE DOES NOT EXIST", "generated CREATE INDEX must fail fast when target table is missing")
		require(strings.Contains(semanticUpper, "CREATE UNIQUE ") && strings.Contains(semanticUpper, " INDEX "), "HAVING COUNT_BIG(*) > 1", "generated UNIQUE INDEX must precheck duplicate key values")
	}
}

func validateBootstrapCompatibleCreateTableGuard(report *output.Report, relPath, rawUpper, semanticUpper string, strict bool) {
	if migrationVersionFromPath(relPath) < bootstrapCompatibleTablePrecheckVersion {
		return
	}
	if !strings.Contains(semanticUpper, "CREATE TABLE ") {
		return
	}
	if !strings.Contains(rawUpper, "EXISTING TABLE MATCH PROFILE: BOOTSTRAP-COMPATIBLE") {
		addHardenedPrecheckFinding(
			report,
			relPath,
			"generated CREATE TABLE wrapper must declare the bootstrap-compatible profile",
			strict,
		)
	}
	requiredSemantic := []struct {
		token   string
		message string
	}{
		{"FROM SYS.COLUMNS AS C", "generated CREATE TABLE wrapper must inspect required columns"},
		{"TYPE_NAME(C.USER_TYPE_ID)", "generated CREATE TABLE wrapper must validate required base types"},
		{"C.IS_IDENTITY =", "generated CREATE TABLE wrapper must validate identity-kind"},
		{"C.IS_COMPUTED =", "generated CREATE TABLE wrapper must validate computed-kind"},
	}
	for _, item := range requiredSemantic {
		if strings.Contains(semanticUpper, item.token) {
			continue
		}
		addHardenedPrecheckFinding(report, relPath, item.message, strict)
	}
	for _, forbidden := range []string{
		"COUNT_BIG(*) FROM SYS.COLUMNS",
		"C.MAX_LENGTH =",
		"C.PRECISION =",
		"C.SCALE =",
		"C.IS_NULLABLE =",
		"FROM SYS.FOREIGN_KEYS AS FK",
		"FROM SYS.CHECK_CONSTRAINTS AS CC",
		"FROM SYS.KEY_CONSTRAINTS AS KC",
	} {
		if !strings.Contains(semanticUpper, forbidden) {
			continue
		}
		addHardenedPrecheckFinding(
			report,
			relPath,
			"generated CREATE TABLE bootstrap guard contains a blocking DBA-managed detail: "+forbidden,
			strict,
		)
	}
}

func migrationVersionFromPath(relPath string) string {
	base := relPath
	if sep := strings.LastIndexAny(base, `/\\`); sep >= 0 {
		base = base[sep+1:]
	}
	if len(base) < 17 {
		return ""
	}
	version := base[:17]
	for _, ch := range version {
		if ch < '0' || ch > '9' {
			return ""
		}
	}
	return version
}

func isHardenedGeneratedMigration(relPath string) bool {
	version := migrationVersionFromPath(relPath)
	return version >= hardenedPrecheckVersion
}

func containsCreateIndexDDL(semanticUpper string) bool {
	for _, anchor := range []string{"CREATE INDEX ", "CREATE UNIQUE INDEX ", "CREATE UNIQUE CLUSTERED INDEX ", "CREATE UNIQUE NONCLUSTERED INDEX ", "CREATE CLUSTERED INDEX ", "CREATE NONCLUSTERED INDEX ", "CREATE PRIMARY XML INDEX ", "CREATE XML INDEX ", "CREATE SPATIAL INDEX ", "CREATE CLUSTERED COLUMNSTORE INDEX ", "CREATE NONCLUSTERED COLUMNSTORE INDEX "} {
		if strings.Contains(semanticUpper, anchor) {
			return true
		}
	}
	return false
}

func addHardenedPrecheckFinding(report *output.Report, relPath, message string, strict bool) {
	item := output.Item{
		Rule:       "MIGRATION_HARDENED_PRECHECK_MISSING",
		Severity:   2,
		File:       relPath,
		Message:    message,
		Suggestion: "regenerate the migration with the hardened generator; do not deploy an unguarded DDL",
		Reason:     "REQUIRES_MANUAL",
	}
	if strict {
		item.Action = "FAIL"
		report.AddFail(item)
		return
	}
	item.Action = "warn"
	report.AddPass(item)
}

func isLegacyNoPrecheckMigration(sqlText string) bool {
	upperRaw := strings.ToUpper(sqlText)
	return strings.Contains(upperRaw, "-- NO PRECHECK")
}

func validateTablePrecheckGuards(report *output.Report, relPath, upperSQL string, strict bool) {
	const anchor = "ALTER TABLE "
	cursor := 0
	for {
		rel := strings.Index(upperSQL[cursor:], anchor)
		if rel < 0 {
			return
		}
		pos := cursor + rel
		end := sqlparse.FindStatementEnd(upperSQL, pos)
		stmt := upperSQL[pos:end]
		stmtHasRisk := strings.Contains(stmt, " ADD ") || strings.Contains(stmt, " DROP COLUMN") || strings.Contains(stmt, " ALTER COLUMN")
		if stmtHasRisk && !hasGuardBeforeRiskStatement(upperSQL, pos, []string{"COL_LENGTH(", "EXISTS (", "OBJECT_ID("}) {
			addPrecheckGuardFinding(report, relPath, "table migration statement may run without explicit precheck guard", "add IF COL_LENGTH/IF EXISTS guard before ALTER TABLE statement", strict)
			return
		}
		cursor = end
	}
}

func validateIndexPrecheckGuards(report *output.Report, relPath, upperSQL string, strict bool) {
	indexAnchors := []string{"CREATE INDEX ", "CREATE UNIQUE INDEX ", "CREATE NONCLUSTERED INDEX ", "CREATE CLUSTERED INDEX ", "DROP INDEX "}
	for _, anchor := range indexAnchors {
		cursor := 0
		for {
			rel := strings.Index(upperSQL[cursor:], anchor)
			if rel < 0 {
				break
			}
			pos := cursor + rel
			if !hasGuardBeforeRiskStatement(upperSQL, pos, []string{"EXISTS ("}) {
				addPrecheckGuardFinding(report, relPath, "index migration statement may run without explicit precheck guard", "add IF EXISTS/IF NOT EXISTS guard before CREATE/DROP INDEX statement", strict)
				return
			}
			cursor = sqlparse.FindStatementEnd(upperSQL, pos)
		}
	}
}

func addPrecheckGuardFinding(report *output.Report, relPath, message, suggestion string, strict bool) {
	item := output.Item{
		Rule:       "MIGRATION_PRECHECK_MISSING",
		Severity:   2,
		File:       relPath,
		Message:    message,
		Suggestion: suggestion,
		Reason:     "REQUIRES_MANUAL",
	}
	if strict {
		item.Action = "FAIL"
		report.AddFail(item)
		return
	}
	item.Action = "warn"
	report.AddPass(item)
}

func hasGuardBeforeRiskStatement(upperSQL string, stmtPos int, anchors []string) bool {
	if hasEnclosingIfBeginGuard(upperSQL, stmtPos, anchors) {
		return true
	}
	if hasPriorGuardReturnBlock(upperSQL, stmtPos, anchors) {
		return true
	}
	stmtStart := 0
	if prevSemi := strings.LastIndex(upperSQL[:stmtPos], ";"); prevSemi >= 0 {
		stmtStart = prevSemi + 1
	}
	prefix := upperSQL[stmtStart:stmtPos]
	if !containsStandaloneWord(prefix, "IF") {
		return false
	}
	for _, anchor := range anchors {
		if strings.Contains(prefix, anchor) {
			return true
		}
	}
	return false
}

func hasPriorGuardReturnBlock(upperSQL string, stmtPos int, anchors []string) bool {
	prefix := upperSQL[:stmtPos]
	cursor := 0
	for {
		ifPos := findNextStandaloneKeyword(prefix, "IF", cursor)
		if ifPos < 0 {
			return false
		}
		beginPos := findNextStandaloneKeyword(prefix, "BEGIN", ifPos)
		if beginPos < 0 {
			return false
		}
		if beginPos >= stmtPos {
			return false
		}
		guard := prefix[ifPos:beginPos]
		matchedAnchor := false
		for _, anchor := range anchors {
			if strings.Contains(guard, anchor) {
				matchedAnchor = true
				break
			}
		}
		if matchedAnchor {
			endPos := findMatchingEndForBegin(prefix, beginPos)
			if endPos > beginPos && endPos <= stmtPos {
				block := prefix[beginPos:endPos]
				if containsStandaloneWord(block, "RETURN") {
					return true
				}
			}
		}
		cursor = ifPos + len("IF")
	}
}

func hasEnclosingIfBeginGuard(upperSQL string, stmtPos int, anchors []string) bool {
	searchLimit := stmtPos
	for searchLimit > 0 {
		beginPos := findPreviousStandaloneKeyword(upperSQL, "BEGIN", searchLimit)
		if beginPos < 0 {
			return false
		}
		endPos := findMatchingEndForBegin(upperSQL, beginPos)
		if endPos > beginPos && stmtPos < endPos {
			ifPos := findPreviousStandaloneKeyword(upperSQL, "IF", beginPos)
			if ifPos >= 0 {
				guard := upperSQL[ifPos:beginPos]
				for _, anchor := range anchors {
					if strings.Contains(guard, anchor) {
						return true
					}
				}
				if hasSkipVariableGuardWithPrecheck(upperSQL, guard, stmtPos) {
					return true
				}
			}
		}
		searchLimit = beginPos
	}
	return false
}

func hasSkipVariableGuardWithPrecheck(upperSQL, guard string, stmtPos int) bool {
	guardUpper := strings.ToUpper(guard)
	if !strings.Contains(guardUpper, "@SKIP_") || !strings.Contains(guardUpper, "= 0") {
		return false
	}
	varName := extractFirstVarName(guardUpper)
	if varName == "" || !strings.HasPrefix(varName, "@SKIP_") {
		return false
	}
	prefix := upperSQL[:stmtPos]
	if !strings.Contains(prefix, "SET "+varName+" = 1") {
		return false
	}
	return strings.Contains(prefix, "COL_LENGTH(") ||
		strings.Contains(prefix, "EXISTS (") ||
		strings.Contains(prefix, "DATALENGTH(") ||
		strings.Contains(prefix, "TRY_CONVERT(")
}

func extractFirstVarName(upperSQL string) string {
	for i := 0; i < len(upperSQL); i++ {
		if upperSQL[i] != '@' {
			continue
		}
		j := i + 1
		for j < len(upperSQL) && isSQLWordChar(rune(upperSQL[j])) {
			j++
		}
		if j > i+1 {
			return upperSQL[i:j]
		}
	}
	return ""
}

func findMatchingEndForBegin(upperSQL string, beginPos int) int {
	depth := 0
	cursor := beginPos
	for {
		nextBegin := findNextStandaloneKeyword(upperSQL, "BEGIN", cursor)
		nextEnd := findNextStandaloneKeyword(upperSQL, "END", cursor)
		if nextBegin < 0 && nextEnd < 0 {
			return -1
		}
		nextPos := nextBegin
		nextWord := "BEGIN"
		if nextPos < 0 || (nextEnd >= 0 && nextEnd < nextPos) {
			nextPos = nextEnd
			nextWord = "END"
		}
		if nextWord == "BEGIN" {
			depth++
			cursor = nextPos + len("BEGIN")
			continue
		}
		depth--
		cursor = nextPos + len("END")
		if depth == 0 {
			return cursor
		}
	}
}

func findPreviousStandaloneKeyword(upperSQL, keyword string, limit int) int {
	if limit > len(upperSQL) {
		limit = len(upperSQL)
	}
	for i := limit - len(keyword); i >= 0; i-- {
		if hasStandaloneKeywordAt(upperSQL, keyword, i) {
			return i
		}
	}
	return -1
}

func findNextStandaloneKeyword(upperSQL, keyword string, start int) int {
	if start < 0 {
		start = 0
	}
	for i := start; i+len(keyword) <= len(upperSQL); i++ {
		if hasStandaloneKeywordAt(upperSQL, keyword, i) {
			return i
		}
	}
	return -1
}

func hasStandaloneKeywordAt(upperSQL, keyword string, pos int) bool {
	if pos < 0 || pos+len(keyword) > len(upperSQL) {
		return false
	}
	if upperSQL[pos:pos+len(keyword)] != keyword {
		return false
	}
	leftOK := pos == 0 || !isSQLWordChar(rune(upperSQL[pos-1]))
	rightPos := pos + len(keyword)
	rightOK := rightPos >= len(upperSQL) || !isSQLWordChar(rune(upperSQL[rightPos]))
	return leftOK && rightOK
}

func containsStandaloneWord(upperSQL, word string) bool {
	for i := 0; i+len(word) <= len(upperSQL); i++ {
		if hasStandaloneKeywordAt(upperSQL, word, i) {
			return true
		}
	}
	return false
}

func isSQLWordChar(ch rune) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || ch == '_' || ch == '@' || ch == '#'
}

// ValidateWrapperDeterministicMetadata validates wrapper deterministic @version/@category/@file metadata.
func ValidateWrapperDeterministicMetadata(report *output.Report, meta MigrationMeta, sqlText, semanticUpperSQL string) {
	if hasDeclaredAssignment(semanticUpperSQL, "@VERSION") {
		declaredVersion, ok := sqlparse.ExtractDeclaredIntegerLiteral(sqlText, "@VERSION")
		if !ok {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_METADATA_NON_DETERMINISTIC",
				Severity: 3,
				File:     meta.RelPath,
				Message:  "wrapper @version must be deterministic integer literal matching filename",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		} else if declaredVersion != meta.Version {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_METADATA_MISMATCH",
				Severity: 3,
				File:     meta.RelPath,
				Message:  fmt.Sprintf("wrapper @version must match filename version (%s)", meta.Version),
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
	}

	if hasDeclaredAssignment(semanticUpperSQL, "@CATEGORY") {
		declaredCategory, ok := sqlparse.ExtractDeclaredNVarcharLiteral(sqlText, "@CATEGORY")
		if !ok {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_METADATA_NON_DETERMINISTIC",
				Severity: 3,
				File:     meta.RelPath,
				Message:  "wrapper @category must be deterministic NVARCHAR literal matching filename category",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		} else if declaredCategory != meta.Category {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_METADATA_MISMATCH",
				Severity: 3,
				File:     meta.RelPath,
				Message:  "wrapper @category must match filename category",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
	}

	if hasDeclaredAssignment(semanticUpperSQL, "@FILE_NAME") {
		declaredFile, ok := sqlparse.ExtractDeclaredNVarcharLiteral(sqlText, "@FILE_NAME")
		if !ok {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_METADATA_NON_DETERMINISTIC",
				Severity: 3,
				File:     meta.RelPath,
				Message:  "wrapper @file_name must be deterministic NVARCHAR literal matching migration filename",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		} else if declaredFile != meta.FileName {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_METADATA_MISMATCH",
				Severity: 3,
				File:     meta.RelPath,
				Message:  "wrapper @file_name must match migration filename",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
	} else if hasDeclaredAssignment(semanticUpperSQL, "@FILE") {
		declaredFile, ok := sqlparse.ExtractDeclaredNVarcharLiteral(sqlText, "@FILE")
		if !ok {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_METADATA_NON_DETERMINISTIC",
				Severity: 3,
				File:     meta.RelPath,
				Message:  "wrapper @file must be deterministic NVARCHAR literal matching migration filename",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		} else if declaredFile != meta.FileName {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_METADATA_MISMATCH",
				Severity: 3,
				File:     meta.RelPath,
				Message:  "wrapper @file must match migration filename",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
	}

	if hasDeclaredAssignment(semanticUpperSQL, "@TARGET") {
		declaredTarget, ok := sqlparse.ExtractDeclaredNVarcharLiteral(sqlText, "@TARGET")
		if !ok {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_METADATA_NON_DETERMINISTIC",
				Severity: 3,
				File:     meta.RelPath,
				Message:  "wrapper @target must be deterministic NVARCHAR literal matching filename object",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		} else if !wrapperTargetMatchesMetadata(meta, declaredTarget) {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_METADATA_MISMATCH",
				Severity: 3,
				File:     meta.RelPath,
				Message:  "wrapper @target must exactly match filename object, including case",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
	}
}

func wrapperTargetMatchesMetadata(meta MigrationMeta, declaredTarget string) bool {
	if declaredTarget == meta.Object {
		return true
	}
	// An index wrapper may identify a specific index as table.index. The
	// table prefix remains an exact-case contract with the filename object.
	return meta.Category == "indexes" &&
		strings.HasPrefix(declaredTarget, meta.Object+".") &&
		len(declaredTarget) > len(meta.Object)+1
}

// ValidateWrapperAuditWhitelist validates wrapper decision/action/reason/severity against whitelist.
func ValidateWrapperAuditWhitelist(
	report *output.Report,
	relPath string,
	upperSQL string,
	allowedDecisions map[string]struct{},
	allowedActions map[string]struct{},
	allowedReasons map[string]struct{},
	allowedSeverities map[int]struct{},
) {
	decisions := sqlparse.ExtractNamedNVarcharParamLiterals(upperSQL, "@DECISION")
	for _, decision := range decisions {
		if _, ok := allowedDecisions[strings.ToLower(decision)]; ok {
			continue
		}
		report.AddFail(output.Item{
			Rule:     "WRAPPER_AUDIT_DECISION_INVALID",
			Severity: 3,
			File:     relPath,
			Message:  "wrapper decision is outside whitelist: " + decision,
			Reason:   "REQUIRES_MANUAL",
			Action:   "FAIL",
		})
	}

	actions := sqlparse.ExtractNamedNVarcharParamLiterals(upperSQL, "@ACTION")
	for _, action := range actions {
		if _, ok := allowedActions[strings.ToLower(action)]; ok {
			continue
		}
		report.AddFail(output.Item{
			Rule:     "WRAPPER_AUDIT_ACTION_INVALID",
			Severity: 3,
			File:     relPath,
			Message:  "wrapper action is outside whitelist: " + action,
			Reason:   "REQUIRES_MANUAL",
			Action:   "FAIL",
		})
	}

	if declaredAction, ok := sqlparse.ExtractDeclaredNVarcharLiteral(upperSQL, "@ACTION"); ok {
		if _, valid := allowedActions[strings.ToLower(declaredAction)]; !valid {
			report.AddFail(output.Item{
				Rule:     "WRAPPER_AUDIT_ACTION_INVALID",
				Severity: 3,
				File:     relPath,
				Message:  "declared @action is outside whitelist: " + declaredAction,
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
	}

	reasons := sqlparse.ExtractNamedNVarcharParamLiterals(upperSQL, "@REASON")
	for _, reason := range reasons {
		if _, ok := allowedReasons[strings.ToUpper(reason)]; ok {
			continue
		}
		report.AddFail(output.Item{
			Rule:     "WRAPPER_AUDIT_REASON_INVALID",
			Severity: 3,
			File:     relPath,
			Message:  "wrapper reason is outside whitelist: " + reason,
			Reason:   "REQUIRES_MANUAL",
			Action:   "FAIL",
		})
	}

	severities := sqlparse.ExtractNamedIntParamLiterals(upperSQL, "@SEVERITY")
	for _, sev := range severities {
		if _, ok := allowedSeverities[sev]; ok {
			continue
		}
		report.AddFail(output.Item{
			Rule:     "WRAPPER_AUDIT_SEVERITY_INVALID",
			Severity: 3,
			File:     relPath,
			Message:  fmt.Sprintf("wrapper severity is outside whitelist: %d", sev),
			Reason:   "REQUIRES_MANUAL",
			Action:   "FAIL",
		})
	}
}

func hasDeclaredAssignment(upperSQL, varName string) bool {
	anchor := "DECLARE " + varName
	rel := strings.Index(upperSQL, anchor)
	if rel < 0 {
		return false
	}
	start := rel + len(anchor)
	return strings.Index(upperSQL[start:], "=") >= 0
}

func hasSemanticThrowInCatch(sql string) bool {
	upper := strings.ToUpper(sqlparse.StripSQLStringsAndComments(sql))
	cursor := 0
	for {
		rel := strings.Index(upper[cursor:], "BEGIN CATCH")
		if rel < 0 {
			return false
		}

		start := cursor + rel
		endRel := strings.Index(upper[start:], "END CATCH")
		end := len(upper)
		if endRel >= 0 {
			end = start + endRel + len("END CATCH")
		}

		if sqlparse.ContainsStandaloneWord(upper[start:end], "THROW") {
			return true
		}
		if end >= len(upper) {
			return false
		}
		cursor = end
	}
}

func missingAuditContractTokens(upperSQL string) []string {
	missing := make([]string, 0, 6)
	if !strings.Contains(upperSQL, "EXEC DBO.USP_MIG_LOG_EVENT") {
		missing = append(missing, "call_usp_mig_log_event")
	}
	if !strings.Contains(upperSQL, "N'RUN'") {
		missing = append(missing, "decision_run")
	}
	if !strings.Contains(upperSQL, "N'OK'") {
		missing = append(missing, "decision_ok")
	}
	if !strings.Contains(upperSQL, "N'FAIL'") {
		missing = append(missing, "decision_fail")
	}
	if !strings.Contains(upperSQL, "N'MIGRATION_FAILED'") {
		missing = append(missing, "action_migration_failed")
	}
	if !strings.Contains(upperSQL, "N'SQL_ERROR'") {
		missing = append(missing, "reason_sql_error")
	}
	return missing
}
