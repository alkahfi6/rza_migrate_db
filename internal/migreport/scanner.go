package migreport

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/history"
)

const (
	coreAuditUpName   = "000000__core__mig__audit_init.up.sql"
	coreAuditDownName = "000000__core__mig__audit_init.down.sql"
)

// FileInfo holds parsed metadata from a migration filename.
type FileInfo struct {
	FileName      string
	RelPath       string
	MigrationID   string
	Category      string
	SchemaName    string
	ObjectName    string
	ObjectVersion string
	Direction     string // "up" or "down"
	IsCore        bool
	RawBytes      []byte
}

// MigrationPair groups up and down files for a single migration step.
type MigrationPair struct {
	MigrationID   string
	Category      string
	SchemaName    string
	ObjectName    string
	ObjectVersion string
	UpFile        *FileInfo
	DownFile      *FileInfo
	IsPaired      bool
}

// ScanMigrations scans the migrations/ folder within repoRoot deterministically.
func ScanMigrations(repoRoot string) ([]MigrationPair, []FileInfo, []string, []string, error) {
	migrationsDir := filepath.Join(repoRoot, "migrations")
	entries, err := os.ReadDir(migrationsDir)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to read migrations directory: %w", err)
	}

	var filenames []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		lower := strings.ToLower(name)
		if strings.HasSuffix(lower, ".up.sql") || strings.HasSuffix(lower, ".down.sql") {
			filenames = append(filenames, name)
		}
	}
	sort.Strings(filenames)

	var allFiles []FileInfo
	var warnings []string
	var errMsgs []string

	// map from pairKey -> *MigrationPair
	pairMap := make(map[string]*MigrationPair)
	var pairKeys []string

	for _, fname := range filenames {
		relPath := filepath.ToSlash(filepath.Join("migrations", fname))
		fullPath := filepath.Join(migrationsDir, fname)
		content, readErr := os.ReadFile(fullPath)
		if readErr != nil {
			errMsgs = append(errMsgs, fmt.Sprintf("failed to read %s: %v", relPath, readErr))
			continue
		}

		info, ok := ParseMigrationFileName(fname)
		if !ok {
			warnings = append(warnings, fmt.Sprintf("unrecognized migration filename format: %s", fname))
			// Create a fallback file info
			info = FileInfo{
				FileName:      fname,
				RelPath:       relPath,
				MigrationID:   extractVersionPrefix(fname),
				Category:      "unknown",
				SchemaName:    "dbo",
				ObjectName:    fname,
				ObjectVersion: "1",
				Direction:     extractDirection(fname),
			}
		}
		info.RelPath = relPath
		info.RawBytes = content
		allFiles = append(allFiles, info)

		// Key for pairing: migration_id + category + object_name + object_version
		key := strings.ToLower(fmt.Sprintf("%s|%s|%s|%s|%s", info.MigrationID, info.Category, info.SchemaName, info.ObjectName, info.ObjectVersion))
		if info.IsCore {
			key = strings.ToLower(fmt.Sprintf("core|%s|%s", info.MigrationID, info.ObjectName))
		}

		pair, exists := pairMap[key]
		if !exists {
			pair = &MigrationPair{
				MigrationID:   info.MigrationID,
				Category:      info.Category,
				SchemaName:    info.SchemaName,
				ObjectName:    info.ObjectName,
				ObjectVersion: info.ObjectVersion,
			}
			pairMap[key] = pair
			pairKeys = append(pairKeys, key)
		}

		if info.Direction == "up" {
			if pair.UpFile != nil {
				warnings = append(warnings, fmt.Sprintf("duplicate .up.sql file for migration %s: %s", key, fname))
				continue
			}
			pair.UpFile = &info
		} else if info.Direction == "down" {
			if pair.DownFile != nil {
				warnings = append(warnings, fmt.Sprintf("duplicate .down.sql file for migration %s: %s", key, fname))
				continue
			}
			pair.DownFile = &info
		}
	}

	var pairs []MigrationPair
	for _, key := range pairKeys {
		p := pairMap[key]
		p.IsPaired = (p.UpFile != nil && p.DownFile != nil)
		if !p.IsPaired {
			if p.UpFile != nil {
				warnings = append(warnings, fmt.Sprintf("migration %s has .up.sql without matching .down.sql (%s)", p.MigrationID, p.UpFile.RelPath))
			} else if p.DownFile != nil {
				warnings = append(warnings, fmt.Sprintf("migration %s has .down.sql without matching .up.sql (%s)", p.MigrationID, p.DownFile.RelPath))
			}
		}
		pairs = append(pairs, *p)
	}

	// Sort pairs deterministically: migration_id, category, schema_name, object_name
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].MigrationID != pairs[j].MigrationID {
			return pairs[i].MigrationID < pairs[j].MigrationID
		}
		if pairs[i].Category != pairs[j].Category {
			return pairs[i].Category < pairs[j].Category
		}
		if pairs[i].SchemaName != pairs[j].SchemaName {
			return pairs[i].SchemaName < pairs[j].SchemaName
		}
		return pairs[i].ObjectName < pairs[j].ObjectName
	})

	return pairs, allFiles, warnings, errMsgs, nil
}

// ParseMigrationFileName parses migration filenames into FileInfo.
func ParseMigrationFileName(name string) (FileInfo, bool) {
	cfg := config.Default(".")
	meta, issue := history.ParseMigrationFileName(name, history.ParseOptions{
		CoreAuditUpName:         coreAuditUpName,
		CoreAuditDownName:       coreAuditDownName,
		AllowedCategories:       cfg.AllowedCategories,
		AllowedObjectTokenRunes: cfg.AllowedObjectTokenRunes,
	})
	if issue != nil {
		return FileInfo{}, false
	}

	schemaName := "dbo"
	objName := meta.Object
	objVersion := fmt.Sprintf("%d", meta.VK)
	if meta.IsCore {
		objName = "mig__audit_init"
	} else if meta.Category != "patch" {
		schemaName, objName = ExtractSchemaAndObject(meta.Object, "dbo")
	}

	return FileInfo{
		FileName:      name,
		MigrationID:   meta.Version,
		Category:      meta.Category,
		SchemaName:    schemaName,
		ObjectName:    objName,
		ObjectVersion: objVersion,
		Direction:     meta.Direction,
		IsCore:        meta.IsCore,
	}, true
}

func extractVersionPrefix(name string) string {
	parts := strings.SplitN(name, "_", 2)
	if len(parts) > 0 {
		return parts[0]
	}
	return "000000"
}

func extractDirection(name string) string {
	lower := strings.ToLower(name)
	if strings.HasSuffix(lower, ".up.sql") {
		return "up"
	}
	if strings.HasSuffix(lower, ".down.sql") {
		return "down"
	}
	return "up"
}
