package engine

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"rza_migrate_db/internal/sqlfix"
	"rza_migrate_db/internal/sqlfmt"
)

type safePatchWrite struct {
	Rel     string
	Abs     string
	Content []byte
}

func (e Engine) collectSafePatchFixes() ([]safePatchWrite, []sqlfix.Change, error) {
	root := e.cfg.RepoPath(e.cfg.PatchDir)
	files, err := listFilesRecursive(root)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil, nil
		}
		return nil, nil, err
	}

	writes := make([]safePatchWrite, 0)
	changes := make([]sqlfix.Change, 0)
	for _, file := range files {
		if !strings.EqualFold(filepath.Ext(file), ".sql") {
			continue
		}
		raw, err := os.ReadFile(file)
		if err != nil {
			return nil, nil, err
		}
		utf8Source, encodingFix, err := sqlfmt.NormalizeSQLToUTF8(raw)
		if err != nil {
			return nil, nil, fmt.Errorf("normalize patch %s: %w", e.toRelPath(file), err)
		}
		fixed, err := sqlfix.RewriteSafe(string(utf8Source), sqlfix.Options{Category: "patch", Patch: true})
		if err != nil {
			return nil, nil, err
		}
		if fixed.SQL == string(raw) && !encodingFix.Changed {
			continue
		}
		writes = append(writes, safePatchWrite{Rel: e.toRelPath(file), Abs: file, Content: []byte(fixed.SQL)})
		changes = append(changes, fixed.Changes...)
		if encodingFix.Changed {
			changes = append(changes, sqlfix.Change{
				Rule:    "SAFE_SQL_ENCODING_UTF8",
				Message: "converted patch source from " + encodingFix.SourceEncoding + " to UTF-8",
			})
		}
	}
	return writes, changes, nil
}
