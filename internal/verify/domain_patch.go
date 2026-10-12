package verify

import (
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"rza_migrate_db/internal/fsx"
	"rza_migrate_db/internal/output"
	"rza_migrate_db/internal/sqlfmt"
)

// MigrationMeta is minimal migration metadata used by verify validators.
type MigrationMeta struct {
	FileName  string
	Version   string
	IsCore    bool
	Category  string
	Direction string
	Object    string
	RelPath   string
	VK        int
}

// PatchFoldersOptions holds runtime context for patch folder validation.
type PatchFoldersOptions struct {
	RepoPath func(string) string
	PatchDir string
}

// PatchSourceIntegrityOptions holds runtime context for patch source integrity validation.
type PatchSourceIntegrityOptions struct {
	RepoPath              func(string) string
	PatchDir              string
	Metas                 []MigrationMeta
	CanonicalizePatchName func(string) string
	HashStableContent     func(string) string
}

// ValidatePatchFolders checks patch root and per-folder strict contract.
func ValidatePatchFolders(report *output.Report, opts PatchFoldersOptions) error {
	patchRoot := opts.RepoPath(opts.PatchDir)
	entries, err := os.ReadDir(patchRoot)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			report.AddFail(output.Item{
				Rule:     "PATCH_ROOT_ONLY_DIRECTORIES",
				Severity: 3,
				File:     filepath.ToSlash(filepath.Join(opts.PatchDir, entry.Name())),
				Message:  "patch root must contain directories only",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
			continue
		}

		patchDir := filepath.Join(patchRoot, entry.Name())
		children, err := os.ReadDir(patchDir)
		if err != nil {
			return err
		}

		found := map[string]int{}
		allowed := map[string]bool{
			"readme.md": true,
			"up.sql":    true,
			"down.sql":  true,
		}

		valid := len(children) == len(allowed)
		for _, child := range children {
			if child.IsDir() {
				valid = false
				continue
			}
			lower := strings.ToLower(child.Name())
			found[lower]++
			if !allowed[lower] {
				valid = false
			}
		}

		for req := range allowed {
			if found[req] != 1 {
				valid = false
			}
		}

		if !valid {
			report.AddFail(output.Item{
				Rule:     "PATCH_FOLDER_INVALID_CONTENT",
				Severity: 3,
				File:     filepath.ToSlash(filepath.Join(opts.PatchDir, entry.Name())),
				Message:  "patch folder must contain exactly README.md, up.sql, down.sql",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
	}

	return nil
}

// ValidatePatchSourceIntegrity verifies patch source/header/hash consistency for existing patch migrations.
func ValidatePatchSourceIntegrity(report *output.Report, opts PatchSourceIntegrityOptions) error {
	patchInHistory := map[string]MigrationMeta{}
	for _, meta := range opts.Metas {
		if meta.IsCore {
			continue
		}
		if meta.Category != "patch" {
			continue
		}
		if meta.Direction != "up" {
			continue
		}
		patchInHistory[meta.Object] = meta
	}
	if len(patchInHistory) == 0 {
		return nil
	}

	type patchSourceInfo struct {
		RelDir     string
		ReadmePath string
		UpHash     string
		DownHash   string
	}

	availablePatchCanonical := map[string]patchSourceInfo{}
	duplicatePatchCanonical := map[string]bool{}
	patchRoot := opts.RepoPath(opts.PatchDir)
	entries, err := os.ReadDir(patchRoot)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
		entries = []os.DirEntry{}
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		patchDir := filepath.Join(patchRoot, entry.Name())
		resolved, ok, err := fsx.ResolveLogicalFilesCaseInsensitive(patchDir, "README.md", "up.sql", "down.sql")
		if err != nil {
			return err
		}
		if !ok {
			continue
		}
		readmePath := resolved["readme.md"]
		upPath := resolved["up.sql"]
		downPath := resolved["down.sql"]
		upBytes, err := os.ReadFile(upPath)
		if err != nil {
			continue
		}
		downBytes, err := os.ReadFile(downPath)
		if err != nil {
			continue
		}

		canonical := opts.CanonicalizePatchName(entry.Name())
		if canonical == "" {
			continue
		}
		if _, exists := availablePatchCanonical[canonical]; exists {
			duplicatePatchCanonical[canonical] = true
			continue
		}
		relDir := filepath.ToSlash(filepath.Join(opts.PatchDir, entry.Name()))
		availablePatchCanonical[canonical] = patchSourceInfo{
			RelDir:     relDir,
			ReadmePath: filepath.ToSlash(filepath.Join(relDir, filepath.Base(readmePath))),
			UpHash:     opts.HashStableContent(string(upBytes)),
			DownHash:   opts.HashStableContent(string(downBytes)),
		}
	}

	patchNames := make([]string, 0, len(patchInHistory))
	for patchName := range patchInHistory {
		patchNames = append(patchNames, patchName)
	}
	sort.Strings(patchNames)
	for _, patchName := range patchNames {
		historyMeta := patchInHistory[patchName]
		if duplicatePatchCanonical[patchName] {
			report.AddPass(output.Item{
				Rule:       "PATCH_SOURCE_DUPLICATE_PENDING",
				Severity:   2,
				File:       opts.PatchDir,
				Object:     patchName,
				Message:    "multiple patch folders map to same canonical name; integrity check deferred to patch planning",
				Suggestion: "remove duplicate patch folder names before gen",
				Reason:     "REQUIRES_MANUAL",
				Action:     "warn",
			})
			continue
		}

		hdr, err := readPatchIntegrityHeader(opts.RepoPath(historyMeta.RelPath))
		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return err
			}
			hdr = patchIntegrityHeader{}
		}

		sourceInfo, ok := availablePatchCanonical[patchName]
		if !ok {
			expectedSource := opts.PatchDir
			if strings.TrimSpace(hdr.PatchFolder) != "" {
				expectedSource = filepath.ToSlash(filepath.Join(opts.PatchDir, hdr.PatchFolder))
			}
			report.AddFail(output.Item{
				Rule:       "PATCH_SOURCE_MISSING",
				Severity:   3,
				File:       expectedSource,
				Object:     patchName,
				Location:   historyMeta.RelPath,
				Message:    "patch source folder or its required files are missing for an existing patch migration",
				Suggestion: "restore and commit exactly one README.md, up.sql, and down.sql under the reported patch folder; verify Git tracking and filename casing",
				Reason:     "DRIFT_DETECTED",
				Action:     "FAIL",
			})
			continue
		}

		if strings.TrimSpace(hdr.PatchFolder) == "" || strings.TrimSpace(hdr.PatchName) == "" || strings.TrimSpace(hdr.Readme) == "" {
			report.AddFail(output.Item{
				Rule:     "PATCH_HEADER_METADATA_MISSING",
				Severity: 3,
				File:     historyMeta.RelPath,
				Object:   patchName,
				Message:  "patch migration header must include patch_folder, patch_name, and readme",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
			continue
		}
		expectedCanonical := opts.CanonicalizePatchName(hdr.PatchFolder)
		expectedReadme := filepath.ToSlash(filepath.Join(opts.PatchDir, hdr.PatchFolder, "README.md"))
		if !strings.EqualFold(hdr.PatchName, patchName) ||
			expectedCanonical == "" ||
			!strings.EqualFold(expectedCanonical, patchName) ||
			!strings.EqualFold(filepath.ToSlash(filepath.Join(opts.PatchDir, hdr.PatchFolder)), sourceInfo.RelDir) ||
			!strings.EqualFold(hdr.Readme, expectedReadme) ||
			!strings.EqualFold(hdr.Readme, sourceInfo.ReadmePath) {
			report.AddFail(output.Item{
				Rule:     "PATCH_HEADER_METADATA_MISMATCH",
				Severity: 3,
				File:     historyMeta.RelPath,
				Object:   patchName,
				Message:  "patch migration header metadata is not aligned with patch source folder",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
			continue
		}
		if hdr.UpHash == "" || hdr.DownHash == "" {
			report.AddPass(output.Item{
				Rule:       "PATCH_IMMUTABILITY_UNVERIFIED",
				Severity:   2,
				File:       historyMeta.RelPath,
				Object:     patchName,
				Message:    "patch migration does not contain integrity hashes",
				Suggestion: "regenerate patch migration using current engine format to enforce source immutability",
				Reason:     "REQUIRES_MANUAL",
				Action:     "warn",
			})
			continue
		}

		if !isHexDigest64(hdr.UpHash) || !isHexDigest64(hdr.DownHash) {
			report.AddFail(output.Item{
				Rule:     "PATCH_HEADER_HASH_INVALID",
				Severity: 3,
				File:     historyMeta.RelPath,
				Object:   patchName,
				Message:  "patch migration integrity hash format is invalid",
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
			continue
		}

		if !strings.EqualFold(hdr.UpHash, sourceInfo.UpHash) || !strings.EqualFold(hdr.DownHash, sourceInfo.DownHash) {
			report.AddFail(output.Item{
				Rule:     "PATCH_SOURCE_MUTATED",
				Severity: 3,
				File:     sourceInfo.RelDir,
				Object:   patchName,
				Message:  "patch source content changed after migration generation",
				Reason:   "DRIFT_DETECTED",
				Action:   "FAIL",
			})
			continue
		}

		report.AddPass(output.Item{
			Rule:     "PATCH_IMMUTABILITY_OK",
			Severity: 1,
			File:     historyMeta.RelPath,
			Object:   patchName,
			Message:  "patch source integrity hash matches generated migration header",
			Reason:   "OBJECT_EXISTS",
			Action:   "ok",
		})
	}

	return nil
}

type patchIntegrityHeader struct {
	PatchFolder string
	PatchName   string
	Readme      string
	UpHash      string
	DownHash    string
}

func readPatchIntegrityHeader(path string) (patchIntegrityHeader, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return patchIntegrityHeader{}, err
	}
	return parsePatchIntegrityHeader(string(raw)), nil
}

func parsePatchIntegrityHeader(sql string) patchIntegrityHeader {
	normalized := sqlfmt.NormalizeSQLText(sql, true, true)
	trimmed := strings.TrimSpace(normalized)
	if !strings.HasPrefix(trimmed, "/*") {
		return patchIntegrityHeader{}
	}

	end := strings.Index(trimmed, "*/")
	if end < 0 {
		return patchIntegrityHeader{}
	}
	block := trimmed[2:end]

	out := patchIntegrityHeader{}
	lines := strings.Split(block, "\n")
	for _, line := range lines {
		token := strings.TrimSpace(line)
		if token == "" {
			continue
		}
		parts := strings.SplitN(token, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.ToLower(strings.TrimSpace(parts[0]))
		valueRaw := strings.TrimSpace(parts[1])
		valueLower := strings.ToLower(valueRaw)
		switch key {
		case "patch_folder":
			out.PatchFolder = valueRaw
		case "patch_name":
			out.PatchName = valueRaw
		case "readme":
			out.Readme = filepath.ToSlash(valueRaw)
		case "patch_up_sha256":
			out.UpHash = valueLower
		case "patch_down_sha256":
			out.DownHash = valueLower
		}
	}
	return out
}

func isHexDigest64(value string) bool {
	if len(value) != 64 {
		return false
	}
	for _, ch := range value {
		switch {
		case ch >= '0' && ch <= '9':
		case ch >= 'a' && ch <= 'f':
		case ch >= 'A' && ch <= 'F':
		default:
			return false
		}
	}
	return true
}
