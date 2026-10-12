package config

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

// ExecutableNameLock configures allowed executable names.
// If Enforce=true and current executable name is outside AllowedNames,
// the process must stop.
type ExecutableNameLock struct {
	Enforce      bool
	AllowedNames []string
}

// DefaultMigrateExecutableNameLock returns strict lock for rza_migrate_db binary name.
func DefaultMigrateExecutableNameLock() ExecutableNameLock {
	return ExecutableNameLock{
		Enforce: true,
		AllowedNames: []string{
			"rza_migrate_db",
			"rza_migrate_db.exe",
		},
	}
}

// DefaultMigrateCheckExecutableNameLock returns the strict lock for the
// developer-facing format+verify checker.
func DefaultMigrateCheckExecutableNameLock() ExecutableNameLock {
	return ExecutableNameLock{
		Enforce: true,
		AllowedNames: []string{
			"rza_migrate_check",
			"rza_migrate_check.exe",
		},
	}
}

// DefaultGetObjectExecutableNameLock returns strict lock for rza_get_object_db binary name.
func DefaultGetObjectExecutableNameLock() ExecutableNameLock {
	return ExecutableNameLock{
		Enforce: true,
		AllowedNames: []string{
			"rza_get_object_db",
			"rza_get_object_db.exe",
		},
	}
}

// DefaultMigrateRestExecutableNameLock returns strict lock for rza_migrate_rest binary name.
func DefaultMigrateRestExecutableNameLock() ExecutableNameLock {
	return ExecutableNameLock{
		Enforce: true,
		AllowedNames: []string{
			"rza_migrate_rest",
			"rza_migrate_rest.exe",
		},
	}
}

// DefaultMigrateReportExecutableNameLock returns strict lock for rza_migrate_report binary name.
func DefaultMigrateReportExecutableNameLock() ExecutableNameLock {
	return ExecutableNameLock{
		Enforce: true,
		AllowedNames: []string{
			"rza_migrate_report",
			"rza_migrate_report.exe",
		},
	}
}

// ValidateExecutableName validates current executable name against lock policy.
func ValidateExecutableName(currentArg0 string, lock ExecutableNameLock) error {
	if !lock.Enforce {
		return nil
	}
	current := normalizeExecutableName(currentArg0)
	if current == "" {
		return fmt.Errorf("executable name lock violation: current executable name is empty")
	}

	allowedSet := map[string]struct{}{}
	for _, item := range lock.AllowedNames {
		normalized := normalizeExecutableName(item)
		if normalized == "" {
			continue
		}
		allowedSet[normalized] = struct{}{}
	}
	if len(allowedSet) == 0 {
		return fmt.Errorf("executable name lock violation: allowed names are empty")
	}
	if _, ok := allowedSet[current]; ok {
		return nil
	}

	allowedList := make([]string, 0, len(allowedSet))
	for name := range allowedSet {
		allowedList = append(allowedList, name+".exe")
	}
	sort.Strings(allowedList)
	return fmt.Errorf(
		"executable name lock violation: got %q, expected one of %s (rename is not allowed)",
		filepath.Base(strings.TrimSpace(currentArg0)),
		strings.Join(allowedList, ", "),
	)
}

func normalizeExecutableName(input string) string {
	name := strings.TrimSpace(input)
	if name == "" {
		return ""
	}
	name = strings.ToLower(filepath.Base(name))
	name = strings.TrimSuffix(name, ".exe")
	return strings.TrimSpace(name)
}
