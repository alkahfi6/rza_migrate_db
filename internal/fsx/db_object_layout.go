package fsx

import (
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var dbObjectCategories = map[string]struct{}{
	"types":      {},
	"tables":     {},
	"indexes":    {},
	"views":      {},
	"functions":  {},
	"procedures": {},
}

// DbObjectCategoryCanonicalRel returns canonical category folder (new layout):
// db_object/<category>
func DbObjectCategoryCanonicalRel(dbObjectDir, category string) string {
	return filepath.Join(dbObjectDir, strings.ToLower(strings.TrimSpace(category)))
}

// DbObjectCategoryLegacyRel returns legacy category folder:
// db_object/schema/<category>
func DbObjectCategoryLegacyRel(dbObjectDir, category string) string {
	return filepath.Join(dbObjectDir, "schema", strings.ToLower(strings.TrimSpace(category)))
}

// DbObjectCategoryCandidatesRel returns lookup order for transition mode:
// 1) canonical path
// 2) legacy path
func DbObjectCategoryCandidatesRel(dbObjectDir, category string) []string {
	return []string{
		DbObjectCategoryCanonicalRel(dbObjectDir, category),
		DbObjectCategoryLegacyRel(dbObjectDir, category),
	}
}

// ResolveDbObjectCategoryFiles returns resolved file list for one category with
// transition compatibility:
// - canonical files have precedence over legacy files with the same filename.
// - legacy-only files are still included.
func ResolveDbObjectCategoryFiles(
	repoPath func(string) string,
	listFiles func(string) ([]string, error),
	dbObjectDir,
	category string,
) ([]string, error) {
	cat := strings.ToLower(strings.TrimSpace(category))
	if _, ok := dbObjectCategories[cat]; !ok {
		return nil, errors.New("unknown db_object category: " + category)
	}

	seenBase := make(map[string]struct{})
	out := make([]string, 0)
	for _, rel := range DbObjectCategoryCandidatesRel(dbObjectDir, cat) {
		files, err := listFiles(repoPath(rel))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return nil, err
		}
		for _, abs := range files {
			base := strings.ToLower(filepath.Base(abs))
			if _, exists := seenBase[base]; exists {
				continue
			}
			seenBase[base] = struct{}{}
			out = append(out, abs)
		}
	}
	sort.Strings(out)
	return out, nil
}
