package migrest

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func buildAppliedMigrationFiles(migrationsDir string, before, after VersionState) ([]AppliedMigrationFile, string, string, error) {
	if sameVersion(before.Version, after.Version) {
		return []AppliedMigrationFile{}, "", "none", nil
	}
	direction := "up"
	if versionGreater(before.Version, after.Version) {
		direction = "down"
	}
	entries, err := os.ReadDir(migrationsDir)
	if err != nil {
		return nil, "", direction, fmt.Errorf("read migrations for operation report: %w", err)
	}
	files := make([]AppliedMigrationFile, 0)
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(strings.ToLower(entry.Name()), ".up.sql") {
			continue
		}
		meta, issue := parseMigrationFileMeta(entry.Name())
		if issue != nil {
			continue
		}
		version, parseErr := strconv.ParseUint(meta.Version, 10, 64)
		if parseErr != nil || !versionWithinOperation(version, before.Version, after.Version, direction) {
			continue
		}
		name := entry.Name()
		if direction == "down" {
			name = strings.TrimSuffix(name, ".up.sql") + ".down.sql"
		}
		raw, readErr := os.ReadFile(filepath.Join(migrationsDir, name))
		if readErr != nil {
			return nil, "", direction, fmt.Errorf("read applied migration %s: %w", name, readErr)
		}
		sum := sha256.Sum256(raw)
		files = append(files, AppliedMigrationFile{
			Version:   version,
			FileName:  name,
			Direction: direction,
			SHA256:    fmt.Sprintf("%x", sum),
		})
	}
	sort.Slice(files, func(i, j int) bool {
		if direction == "down" {
			return files[i].Version > files[j].Version
		}
		return files[i].Version < files[j].Version
	})
	return files, buildMigrationManifest(files), direction, nil
}

func buildMigrationManifest(files []AppliedMigrationFile) string {
	manifest := sha256.New()
	for _, file := range files {
		_, _ = fmt.Fprintf(manifest, "%d|%s|%s|%s\n", file.Version, file.Direction, file.FileName, file.SHA256)
	}
	return fmt.Sprintf("%x", manifest.Sum(nil))
}

func versionWithinOperation(version uint64, before, after *uint64, direction string) bool {
	switch direction {
	case "down":
		return (before == nil || version <= *before) && (after == nil || version > *after)
	default:
		return (before == nil || version > *before) && (after == nil || version <= *after)
	}
}

func sameVersion(left, right *uint64) bool {
	if left == nil || right == nil {
		return left == nil && right == nil
	}
	return *left == *right
}

func versionGreater(left, right *uint64) bool {
	if left == nil {
		return false
	}
	if right == nil {
		return true
	}
	return *left > *right
}
