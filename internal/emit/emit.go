package emit

import (
	"fmt"
	"strings"
)

const versionReservationPrefix = "\x00rza-version:"

// PlannedMigrationRef is minimal data needed for deterministic filename/version allocation.
type PlannedMigrationRef struct {
	Category string
	Object   string
	VK       int
}

// MigrationFileNames returns deterministic up/down migration filenames.
func MigrationFileNames(version string, ref PlannedMigrationRef) (string, string) {
	base := ""
	if ref.Category == "patch" {
		base = fmt.Sprintf("%s_patch_%s_v%d", version, ref.Object, ref.VK)
	} else {
		base = fmt.Sprintf("%s_%s_%s_v%d", version, ref.Category, ref.Object, ref.VK)
	}
	return base + ".up.sql", base + ".down.sql"
}

// AllocateVersion allocates next non-conflicting version by sequence suffix.
func AllocateVersion(baseTimestamp string, startSeq int, pairSeqWidth int, maxPairsPerRun int, reserved map[string]struct{}, ref PlannedMigrationRef) (string, int, error) {
	seq := startSeq
	for {
		if seq >= maxPairsPerRun {
			return "", seq, fmt.Errorf("pair sequence exceeded max %d", maxPairsPerRun)
		}

		version := fmt.Sprintf("%s%0*d", baseTimestamp, pairSeqWidth, seq)
		if versionIsReserved(reserved, version) {
			seq++
			continue
		}
		upName, downName := MigrationFileNames(version, ref)
		if _, exists := reserved[upName]; exists {
			seq++
			continue
		}
		if _, exists := reserved[downName]; exists {
			seq++
			continue
		}
		return version, seq + 1, nil
	}
}

// ReserveVersion records a version independently of object/category filenames.
// SQL migration versions are globally unique even when two files target
// different objects.
func ReserveVersion(reserved map[string]struct{}, version string) {
	if reserved == nil || strings.TrimSpace(version) == "" {
		return
	}
	reserved[versionReservationPrefix+version] = struct{}{}
}

func versionIsReserved(reserved map[string]struct{}, version string) bool {
	if _, exists := reserved[versionReservationPrefix+version]; exists {
		return true
	}
	prefix := version + "_"
	for name := range reserved {
		if strings.HasPrefix(name, prefix) {
			return true
		}
	}
	return false
}
