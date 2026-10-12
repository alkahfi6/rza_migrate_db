package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	// ExistingTableMatchBootstrapCompatible is the production deployment
	// profile for a CREATE TABLE migration whose target table already exists.
	// Required columns, base types, identity-kind, and computed-kind must match;
	// DBA-managed details and extra database objects do not block bootstrap.
	ExistingTableMatchBootstrapCompatible = "bootstrap-compatible"

	ValidationProfileFull      = "full"
	ValidationProfileAuthoring = "authoring"

	MigrationValidationOff       = "off"
	MigrationValidationIntegrity = "integrity"
	MigrationValidationFull      = "full"
)

// AppConfig contains locked defaults and runtime paths for the engine.
type AppConfig struct {
	RepoRoot                                      string
	DbObjectDir                                   string
	MigrationsDir                                 string
	PatchDir                                      string
	OutDir                                        string
	ValidationProfile                             string
	MigrationValidation                           string
	LegacyAuditMode                               string
	EnableProgrammabilitySourceContractForNewOnly bool
	EnableGenAutoIndent                           bool
	PatchFormatMode                               string
	BootstrapMode                                 bool
	DraftMode                                     bool
	DraftDir                                      string
	ManagementManifestFile                        string
	EnableFmtCheck                                bool
	EnablePerfNotes                               bool
	EnableOutputNotes                             bool
	EnableLicence                                 bool
	LicenceFile                                   string
	EnableFailFast                                bool
	EnableStrictMode                              bool
	ExistingTableMatchMode                        string
	EnableParallelRead                            bool
	ParallelWorkers                               int
	PairSeqWidth                                  int
	MaxPairsPerRun                                int
	AllowedSeverities                             map[int]struct{}
	AllowedDecisions                              map[string]struct{}
	AllowedActions                                map[string]struct{}
	AllowedReasons                                map[string]struct{}
	AllowedCategories                             map[string]struct{}
	AllowedObjectTokenRunes                       map[rune]struct{}
	ForbiddenSQLTokens                            []string
	LineEnding                                    string
	TrimTrailingSpaces                            bool
	EnsureFinalNewline                            bool
	FailHardOnParseError                          bool
	NoRegexParsing                                bool
	GenerationOrder                               []string
	OutputNamingPattern                           string
	OutputNamingLineEndings                       string
}

var requiredDecisions = []string{
	"run",
	"ok",
	"skip",
	"warn",
	"fail",
}

var requiredActions = []string{
	"create_type",
	"recreate_type",
	"drop_type",
	"create_table",
	"drop_table",
	"add_column",
	"drop_column",
	"alter_column",
	"backfill",
	"add_default",
	"drop_default",
	"add_check",
	"drop_check",
	"add_unique",
	"drop_unique",
	"add_index",
	"drop_index",
	"add_fk",
	"drop_fk",
	"create_or_alter_view",
	"drop_view",
	"create_or_alter_function",
	"drop_function",
	"create_or_alter_procedure",
	"drop_procedure",
	"followup_required",
	"migration_failed",
}

var requiredReasons = []string{
	"TABLE_EXISTS",
	"TABLE_NOT_EXISTS",
	"COLUMN_EXISTS",
	"COLUMN_NOT_EXISTS",
	"INDEX_EXISTS",
	"INDEX_NOT_EXISTS",
	"DF_EXISTS",
	"DF_NOT_EXISTS",
	"FK_EXISTS",
	"FK_NOT_EXISTS",
	"CK_EXISTS",
	"CK_NOT_EXISTS",
	"UQ_EXISTS",
	"UQ_NOT_EXISTS",
	"OBJECT_EXISTS",
	"OBJECT_NOT_EXISTS",
	"CREATED",
	"DROPPED",
	"ALTERED",
	"BACKFILLED",
	"REPLACED",
	"DROPPED_IMPACTED_INDEX",
	"DROPPED_DEPENDENCY",
	"INDEX_RECREATE_REQUIRED",
	"REQUIRES_MANUAL",
	"NARROWING_WOULD_TRUNCATE",
	"DATA_TOO_LONG_FOR_TARGET",
	"FK_REFERENCED_BY_OTHERS",
	"DATA_RESTORE_NOT_GUARANTEED",
	"DRIFT_DETECTED",
	"INLINE_DEFAULT_NOT_ALLOWED",
	"SQL_ERROR",
}

func Default(repoRoot string) AppConfig {
	cleanRoot := filepath.Clean(repoRoot)
	if absoluteRoot, err := filepath.Abs(cleanRoot); err == nil {
		cleanRoot = absoluteRoot
	}
	return AppConfig{
		RepoRoot:            cleanRoot,
		DbObjectDir:         "db_object",
		MigrationsDir:       "migrations",
		PatchDir:            "patch",
		OutDir:              "out",
		ValidationProfile:   ValidationProfileFull,
		MigrationValidation: MigrationValidationFull,
		LegacyAuditMode:     "fail",
		EnableProgrammabilitySourceContractForNewOnly: false,
		EnableGenAutoIndent:                           true,
		PatchFormatMode:                               "off",
		BootstrapMode:                                 false,
		DraftMode:                                     false,
		DraftDir:                                      "out/draft_migrations",
		ManagementManifestFile:                        "out/object_management_manifest.json",
		EnableFmtCheck:                                false,
		EnablePerfNotes:                               false,
		EnableOutputNotes:                             false,
		EnableLicence:                                 true,
		LicenceFile:                                   "license.key",
		EnableFailFast:                                false,
		EnableStrictMode:                              true,
		ExistingTableMatchMode:                        ExistingTableMatchBootstrapCompatible,
		EnableParallelRead:                            true,
		ParallelWorkers:                               8,

		PairSeqWidth:   3,
		MaxPairsPerRun: 999,
		AllowedSeverities: map[int]struct{}{
			0: {},
			1: {},
			2: {},
			3: {},
		},
		AllowedDecisions: toSet(requiredDecisions),
		AllowedActions:   toSet(requiredActions),
		AllowedReasons:   toSet(requiredReasons),

		AllowedCategories: map[string]struct{}{
			"types":      {},
			"tables":     {},
			"indexes":    {},
			"views":      {},
			"functions":  {},
			"procedures": {},
			"patch":      {},
		},

		AllowedObjectTokenRunes: allowedObjectTokenRunes(),

		ForbiddenSQLTokens: []string{
			"GO",
			"USE",
			"CREATE DATABASE",
			"ALTER DATABASE",
			"CREATE LOGIN",
			"ALTER LOGIN",
			"CREATE USER",
		},

		LineEnding:           "LF",
		TrimTrailingSpaces:   true,
		EnsureFinalNewline:   true,
		FailHardOnParseError: true,
		NoRegexParsing:       true,
		GenerationOrder: []string{
			"types",
			"tables",
			"indexes",
			"functions",
			"views",
			"procedures",
			"patch",
		},
		OutputNamingPattern:     "YYYYMMDDHHMMSS<seq>_<category>_<object>_v<k>.(up|down).sql",
		OutputNamingLineEndings: "LF",
	}
}

func (c AppConfig) Validate() error {
	var issues []string

	if strings.TrimSpace(c.DbObjectDir) == "" {
		issues = append(issues, "DbObjectDir must not be empty")
	}
	if strings.TrimSpace(c.MigrationsDir) == "" {
		issues = append(issues, "MigrationsDir must not be empty")
	}
	if strings.TrimSpace(c.PatchDir) == "" {
		issues = append(issues, "PatchDir must not be empty")
	}
	if strings.TrimSpace(c.OutDir) == "" {
		issues = append(issues, "OutDir must not be empty")
	}
	profile := strings.ToLower(strings.TrimSpace(c.ValidationProfile))
	if profile != ValidationProfileFull && profile != ValidationProfileAuthoring {
		issues = append(issues, "ValidationProfile must be one of: full|authoring")
	}
	migrationValidation := strings.ToLower(strings.TrimSpace(c.MigrationValidation))
	if migrationValidation != MigrationValidationOff && migrationValidation != MigrationValidationIntegrity && migrationValidation != MigrationValidationFull {
		issues = append(issues, "MigrationValidation must be one of: off|integrity|full")
	}
	if profile == ValidationProfileFull && migrationValidation != MigrationValidationFull {
		issues = append(issues, "ValidationProfile full requires MigrationValidation full")
	}
	if strings.TrimSpace(c.LicenceFile) == "" {
		issues = append(issues, "LicenceFile must not be empty")
	}
	patchFormatMode := strings.ToLower(strings.TrimSpace(c.PatchFormatMode))
	if patchFormatMode == "" {
		patchFormatMode = "off"
	}
	if patchFormatMode != "off" && patchFormatMode != "safe" && patchFormatMode != "strict" {
		issues = append(issues, "PatchFormatMode must be one of: off|safe|strict")
	}
	if strings.TrimSpace(c.DraftDir) == "" {
		issues = append(issues, "DraftDir must not be empty")
	}
	if strings.TrimSpace(c.ManagementManifestFile) == "" {
		issues = append(issues, "ManagementManifestFile must not be empty")
	}

	if c.PairSeqWidth != 3 {
		issues = append(issues, "PairSeqWidth must be 3")
	}
	if c.MaxPairsPerRun != 999 {
		issues = append(issues, "MaxPairsPerRun must be 999")
	}
	if c.ParallelWorkers < 1 {
		issues = append(issues, "ParallelWorkers must be >= 1")
	}
	if strings.ToUpper(strings.TrimSpace(c.LineEnding)) != "LF" {
		issues = append(issues, "LineEnding must be LF")
	}
	if strings.ToLower(strings.TrimSpace(c.LegacyAuditMode)) != "fail" {
		issues = append(issues, "LegacyAuditMode must be fail (strict-only)")
	}
	matchMode := strings.ToLower(strings.TrimSpace(c.ExistingTableMatchMode))
	if matchMode == "" {
		matchMode = ExistingTableMatchBootstrapCompatible
	}
	if matchMode != ExistingTableMatchBootstrapCompatible {
		issues = append(issues, "ExistingTableMatchMode must be bootstrap-compatible")
	}
	if !c.TrimTrailingSpaces {
		issues = append(issues, "TrimTrailingSpaces must be true")
	}
	if !c.EnsureFinalNewline {
		issues = append(issues, "EnsureFinalNewline must be true")
	}
	if !c.FailHardOnParseError {
		issues = append(issues, "FailHardOnParseError must be true")
	}
	if !c.NoRegexParsing {
		issues = append(issues, "NoRegexParsing must be true")
	}

	requiredCategories := []string{"types", "tables", "indexes", "views", "functions", "procedures", "patch"}
	missingCategories := missingStringKeys(c.AllowedCategories, requiredCategories)
	if len(missingCategories) > 0 {
		issues = append(issues, "AllowedCategories missing: "+strings.Join(missingCategories, ","))
	}

	missingDecisions := missingStringKeys(c.AllowedDecisions, requiredDecisions)
	if len(missingDecisions) > 0 {
		issues = append(issues, "AllowedDecisions missing: "+strings.Join(missingDecisions, ","))
	}

	missingActions := missingStringKeys(c.AllowedActions, requiredActions)
	if len(missingActions) > 0 {
		issues = append(issues, "AllowedActions missing: "+strings.Join(missingActions, ","))
	}

	missingReasons := missingStringKeys(c.AllowedReasons, requiredReasons)
	if len(missingReasons) > 0 {
		issues = append(issues, "AllowedReasons missing: "+strings.Join(missingReasons, ","))
	}

	for _, sev := range []int{0, 1, 2, 3} {
		if _, ok := c.AllowedSeverities[sev]; !ok {
			issues = append(issues, fmt.Sprintf("AllowedSeverities missing: %d", sev))
		}
	}

	if len(issues) > 0 {
		return fmt.Errorf("config hard lock invalid: %s", strings.Join(issues, "; "))
	}
	return nil
}

func (c AppConfig) RepoPath(rel string) string {
	if filepath.IsAbs(rel) {
		return filepath.Clean(rel)
	}
	return filepath.Join(c.RepoRoot, filepath.FromSlash(rel))
}

// ResolveLicencePath locates the license file for reading.
// It checks tools/<LicenceFile>, <LicenceFile> at repo root, and beside the running executable.
func (c AppConfig) ResolveLicencePath() string {
	if filepath.IsAbs(c.LicenceFile) {
		return c.LicenceFile
	}
	toolsPath := filepath.Join(c.RepoRoot, "tools", filepath.FromSlash(c.LicenceFile))
	if info, err := os.Stat(toolsPath); err == nil && !info.IsDir() {
		return toolsPath
	}
	rootPath := c.RepoPath(c.LicenceFile)
	if info, err := os.Stat(rootPath); err == nil && !info.IsDir() {
		return rootPath
	}
	if execPath, err := os.Executable(); err == nil {
		besideExec := filepath.Join(filepath.Dir(execPath), filepath.FromSlash(c.LicenceFile))
		if info, err := os.Stat(besideExec); err == nil && !info.IsDir() {
			return besideExec
		}
	}
	toolsDir := filepath.Join(c.RepoRoot, "tools")
	if info, err := os.Stat(toolsDir); err == nil && info.IsDir() {
		return toolsPath
	}
	return toolsPath
}

// ResolveLicencePathForWrite returns the target path for license creation.
// If tools/ exists in RepoRoot, it prefers tools/<LicenceFile>, otherwise repo root.
func (c AppConfig) ResolveLicencePathForWrite() string {
	if filepath.IsAbs(c.LicenceFile) {
		return c.LicenceFile
	}
	toolsDir := filepath.Join(c.RepoRoot, "tools")
	if info, err := os.Stat(toolsDir); err == nil && info.IsDir() {
		return filepath.Join(toolsDir, filepath.FromSlash(c.LicenceFile))
	}
	return c.RepoPath(c.LicenceFile)
}

// LooksLikeRepo returns true if root contains database authoring directories (migrations or db_object).
func LooksLikeRepo(root string) bool {
	if strings.TrimSpace(root) == "" {
		return false
	}
	migInfo, err1 := os.Stat(filepath.Join(root, "migrations"))
	hasMig := err1 == nil && migInfo.IsDir()
	dbObjInfo, err2 := os.Stat(filepath.Join(root, "db_object"))
	hasDbObj := err2 == nil && dbObjInfo.IsDir()
	return hasMig || hasDbObj
}

// DiscoverRepoRoot resolves the repository root based on explicit flags, working directory,
// or executable location (e.g. tools/ subdirectory).
func DiscoverRepoRoot(explicit string, arg0 string) string {
	cleanExplicit := strings.TrimSpace(explicit)
	if cleanExplicit != "" && cleanExplicit != "." {
		return cleanExplicit
	}

	// 1. Current working directory
	if LooksLikeRepo(".") {
		return "."
	}

	// 2. Parent of current working directory (e.g. if cwd is tools/)
	if LooksLikeRepo("..") {
		return ".."
	}

	// 3. Executable directory and its parent (e.g. double-click from tools/)
	if execPath, err := os.Executable(); err == nil {
		execDir := filepath.Dir(filepath.Clean(execPath))
		if LooksLikeRepo(execDir) {
			return execDir
		}
		execParent := filepath.Dir(execDir)
		if execParent != execDir && LooksLikeRepo(execParent) {
			return execParent
		}
	}

	// 4. arg0 directory and its parent
	if strings.TrimSpace(arg0) != "" {
		argDir := filepath.Dir(filepath.Clean(arg0))
		if LooksLikeRepo(argDir) {
			return argDir
		}
		argParent := filepath.Dir(argDir)
		if argParent != argDir && LooksLikeRepo(argParent) {
			return argParent
		}
	}

	// 5. Walk up parents of cwd up to 4 levels
	if cwd, err := os.Getwd(); err == nil {
		curr := filepath.Clean(cwd)
		for i := 0; i < 4; i++ {
			parent := filepath.Dir(curr)
			if parent == curr {
				break
			}
			if LooksLikeRepo(parent) {
				return parent
			}
			curr = parent
		}
	}

	if cleanExplicit != "" {
		return cleanExplicit
	}
	return "."
}

func (c AppConfig) OutFilePath() string {
	return filepath.Join(c.RepoRoot, c.OutDir, "output.json")
}

func allowedObjectTokenRunes() map[rune]struct{} {
	m := make(map[rune]struct{})
	for ch := 'a'; ch <= 'z'; ch++ {
		m[ch] = struct{}{}
	}
	for ch := 'A'; ch <= 'Z'; ch++ {
		m[ch] = struct{}{}
	}
	for ch := '0'; ch <= '9'; ch++ {
		m[ch] = struct{}{}
	}
	m['_'] = struct{}{}
	m['#'] = struct{}{}
	m['.'] = struct{}{}
	return m
}

func toSet(values []string) map[string]struct{} {
	out := make(map[string]struct{}, len(values))
	for _, v := range values {
		out[v] = struct{}{}
	}
	return out
}

func missingStringKeys(set map[string]struct{}, required []string) []string {
	out := make([]string, 0)
	for _, key := range required {
		if _, ok := set[key]; !ok {
			out = append(out, key)
		}
	}
	return out
}
