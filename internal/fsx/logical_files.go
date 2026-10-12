package fsx

import (
	"os"
	"path/filepath"
	"strings"
)

// ResolveLogicalFilesCaseInsensitive maps required logical file names to the
// actual directory entries. It returns ok=false for missing files, directory
// entries, or ambiguous case-only duplicates and never guesses between them.
func ResolveLogicalFilesCaseInsensitive(dir string, required ...string) (paths map[string]string, ok bool, err error) {
	wanted := make(map[string]struct{}, len(required))
	for _, name := range required {
		wanted[strings.ToLower(name)] = struct{}{}
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, false, err
	}
	paths = make(map[string]string, len(wanted))
	for _, entry := range entries {
		key := strings.ToLower(entry.Name())
		if _, required := wanted[key]; !required {
			continue
		}
		if entry.IsDir() {
			return nil, false, nil
		}
		if _, duplicate := paths[key]; duplicate {
			return nil, false, nil
		}
		paths[key] = filepath.Join(dir, entry.Name())
	}
	if len(paths) != len(wanted) {
		return paths, false, nil
	}
	return paths, true, nil
}
