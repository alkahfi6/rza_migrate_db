package engine

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"rza_migrate_db/internal/output"
	"rza_migrate_db/internal/sqlfix"
)

func (e Engine) normalizeIndexSourceFileNames(report *output.Report) (int, []string, error) {
	root := e.cfg.RepoPath(e.cfg.DbObjectDir)
	files, err := listFilesRecursive(root)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return 0, nil, nil
		}
		return 0, nil, err
	}
	count := 0
	paths := make([]string, 0)
	for _, source := range files {
		rel := strings.ToLower(strings.ReplaceAll(e.toRelPath(source), "\\", "/"))
		if !strings.EqualFold(filepath.Ext(source), ".sql") || !strings.Contains(rel, "/indexes/") {
			continue
		}
		raw, err := os.ReadFile(source)
		if err != nil {
			return count, paths, err
		}
		targets := sqlfix.DetectIndexTargets(string(raw))
		if len(targets) != 1 {
			continue
		}
		desiredBase := targets[0] + ".sql"
		if strings.EqualFold(filepath.Base(source), desiredBase) {
			continue
		}
		destination := filepath.Join(filepath.Dir(source), desiredBase)
		if _, err := os.Stat(destination); err == nil {
			report.AddFail(output.Item{
				Rule: "SAFE_INDEX_RENAME_COLLISION", Severity: 3, File: e.toRelPath(source),
				Message:    fmt.Sprintf("index source targets %s but destination file already exists", targets[0]),
				Suggestion: "merge or split the conflicting index snapshots manually; neither file was changed",
				Reason:     "REQUIRES_MANUAL", Action: "FAIL",
			})
			continue
		} else if !os.IsNotExist(err) {
			return count, paths, err
		}
		count++
		paths = append(paths, e.toRelPath(source)+" -> "+e.toRelPath(destination))
		if e.cfg.EnableFmtCheck {
			continue
		}
		if err := os.Rename(source, destination); err != nil {
			return count, paths, err
		}
	}
	return count, paths, nil
}
