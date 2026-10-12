package getobject

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// PreviewSummary describes preview formatting result from local SQL files.
type PreviewSummary struct {
	Processed  int
	Written    int
	Failed     int
	FailedList []string
}

// PreviewModulesFromDir normalizes module SQL files from inputDir into outputDir.
// It is intended for local learning/simulation without DB connection.
func PreviewModulesFromDir(cfg AppConfig, inputDir, outputDir string, progress func(phase, message string)) (PreviewSummary, error) {
	cfg.Normalize()
	if strings.TrimSpace(inputDir) == "" {
		return PreviewSummary{}, fmt.Errorf("preview input dir is required")
	}
	if strings.TrimSpace(outputDir) == "" {
		return PreviewSummary{}, fmt.Errorf("preview output dir is required")
	}

	inAbs, err := filepath.Abs(filepath.Clean(inputDir))
	if err != nil {
		return PreviewSummary{}, err
	}
	outAbs, err := filepath.Abs(filepath.Clean(outputDir))
	if err != nil {
		return PreviewSummary{}, err
	}

	if progress != nil {
		progress("preview", fmt.Sprintf("scan input directory %s", inAbs))
	}
	files, err := collectSQLFiles(inAbs)
	if err != nil {
		return PreviewSummary{}, err
	}
	if progress != nil {
		progress("preview", fmt.Sprintf("discovered %d sql files", len(files)))
	}

	if err := os.MkdirAll(outAbs, 0o755); err != nil {
		return PreviewSummary{}, err
	}

	summary := PreviewSummary{
		FailedList: []string{},
	}
	for _, file := range files {
		rel, err := filepath.Rel(inAbs, file)
		if err != nil {
			return summary, err
		}
		rel = filepath.ToSlash(rel)
		summary.Processed++

		raw, err := os.ReadFile(file)
		if err != nil {
			return summary, err
		}
		source := string(raw)

		kind, err := detectModuleKind(rel, source)
		if err != nil {
			summary.Failed++
			summary.FailedList = append(summary.FailedList, rel)
			continue
		}

		sqlText, err := ensureCreateOrAlter(source, kind)
		if err != nil {
			summary.Failed++
			summary.FailedList = append(summary.FailedList, rel)
			continue
		}
		formatted := formatModuleDefinition(sqlText, cfg)
		targetDir := filepath.Join(outAbs, moduleKindDir(kind))
		if err := os.MkdirAll(targetDir, 0o755); err != nil {
			return summary, err
		}

		baseName := filepath.Base(file)
		outPath := filepath.Join(targetDir, baseName)
		if err := os.WriteFile(outPath, []byte(formatted), 0o644); err != nil {
			return summary, err
		}
		summary.Written++
	}

	if progress != nil {
		progress("preview", fmt.Sprintf("done processed=%d written=%d failed=%d output=%s", summary.Processed, summary.Written, summary.Failed, outAbs))
	}
	return summary, nil
}

func collectSQLFiles(root string) ([]string, error) {
	out := make([]string, 0)
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.EqualFold(filepath.Ext(d.Name()), ".sql") {
			return nil
		}
		out = append(out, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Strings(out)
	return out, nil
}

func detectModuleKind(relPath, source string) (string, error) {
	lower := strings.ToLower(filepath.ToSlash(relPath))
	switch {
	case strings.Contains(lower, "/view/"), strings.Contains(lower, "/views/"), strings.HasPrefix(lower, "view/"):
		return "VIEW", nil
	case strings.Contains(lower, "/stored-procedure/"), strings.Contains(lower, "/procedure/"), strings.Contains(lower, "/procedures/"):
		return "PROCEDURE", nil
	case strings.Contains(lower, "/function/"), strings.Contains(lower, "/functions/"), strings.Contains(lower, "scalar_function"), strings.Contains(lower, "table_valued_function"):
		return "FUNCTION", nil
	}

	kinds := []string{"PROCEDURE", "FUNCTION", "VIEW"}
	for _, k := range kinds {
		if _, err := ensureCreateOrAlter(source, k); err == nil {
			return k, nil
		}
	}
	return "", fmt.Errorf("cannot detect module kind for %s", relPath)
}

func moduleKindDir(kind string) string {
	switch strings.ToUpper(strings.TrimSpace(kind)) {
	case "VIEW":
		return "views"
	case "PROCEDURE":
		return "procedures"
	default:
		return "functions"
	}
}
