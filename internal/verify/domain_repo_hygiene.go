package verify

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"rza_migrate_db/internal/fsx"
	"rza_migrate_db/internal/output"
	"rza_migrate_db/internal/sqlfmt"
	"rza_migrate_db/internal/sqlparse"
)

// ForbiddenSQLOptions holds runtime context for forbidden SQL token validation.
type ForbiddenSQLOptions struct {
	RepoPath           func(string) string
	DbObjectDir        string
	MigrationsDir      string
	PatchDir           string
	ForbiddenSQLTokens []string
	ToRelPath          func(string) string
	ListFilesRecursive func(string) ([]string, error)
	EnableParallelRead bool
	ParallelWorkers    int
	ReadFile           func(string) ([]byte, error)
	SkipMigrations     bool
}

// RepoHygieneOptions holds runtime context for repository hygiene validation.
type RepoHygieneOptions struct {
	RepoPath func(string) string
}

// ValidateForbiddenSQL scans db_object/migrations/patch SQL files for forbidden tokens.
func ValidateForbiddenSQL(report *output.Report, opts ForbiddenSQLOptions) error {
	listFiles := opts.ListFilesRecursive
	if listFiles == nil {
		listFiles = fsx.ListFilesRecursive
	}
	toRelPath := opts.ToRelPath
	if toRelPath == nil {
		toRelPath = func(absPath string) string { return filepath.ToSlash(absPath) }
	}

	scanRoots := []string{opts.RepoPath(opts.DbObjectDir), opts.RepoPath(opts.PatchDir)}
	if !opts.SkipMigrations {
		scanRoots = append(scanRoots, opts.RepoPath(opts.MigrationsDir))
	}

	tasks := make([]forbiddenScanTask, 0)
	for _, root := range scanRoots {
		files, err := listFiles(root)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return err
		}

		for _, file := range files {
			if !strings.EqualFold(filepath.Ext(file), ".sql") {
				continue
			}
			tasks = append(tasks, forbiddenScanTask{
				Index: len(tasks),
				File:  file,
			})
		}
	}

	results := runForbiddenScanTasks(opts, tasks)
	for _, task := range tasks {
		res := results[task.Index]
		if res.Err != nil {
			return res.Err
		}
		if res.EncodingIssue != "" {
			report.AddFail(output.Item{
				Rule:     "INVALID_SQL_ENCODING",
				Severity: 3,
				File:     toRelPath(task.File),
				Message:  res.EncodingIssue,
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
			continue
		}
		for _, token := range res.Tokens {
			report.AddFail(output.Item{
				Rule:     "FORBIDDEN_SQL_TOKEN",
				Severity: 3,
				File:     toRelPath(task.File),
				Message:  "forbidden SQL token detected: " + token,
				Reason:   "REQUIRES_MANUAL",
				Action:   "FAIL",
			})
		}
	}

	return nil
}

type forbiddenScanTask struct {
	Index int
	File  string
}

type forbiddenScanResult struct {
	Tokens        []string
	EncodingIssue string
	Err           error
}

func runForbiddenScanTasks(opts ForbiddenSQLOptions, tasks []forbiddenScanTask) []forbiddenScanResult {
	if len(tasks) == 0 {
		return []forbiddenScanResult{}
	}

	results := make([]forbiddenScanResult, len(tasks))
	workers := forbiddenScanWorkers(opts, len(tasks))
	if workers <= 1 {
		for _, task := range tasks {
			results[task.Index] = scanForbiddenTask(opts, task)
		}
		return results
	}

	jobs := make(chan forbiddenScanTask)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range jobs {
				results[task.Index] = scanForbiddenTask(opts, task)
			}
		}()
	}

	for _, task := range tasks {
		jobs <- task
	}
	close(jobs)
	wg.Wait()
	return results
}

func scanForbiddenTask(opts ForbiddenSQLOptions, task forbiddenScanTask) forbiddenScanResult {
	readFile := opts.ReadFile
	if readFile == nil {
		readFile = os.ReadFile
	}
	raw, err := readFile(task.File)
	if err != nil {
		return forbiddenScanResult{Err: err}
	}
	if msg, bad := sqlfmt.DetectSQLEncodingIssue(raw); bad {
		return forbiddenScanResult{EncodingIssue: msg}
	}
	tokens := sqlparse.DetectForbiddenTokensWithList(string(raw), opts.ForbiddenSQLTokens)
	return forbiddenScanResult{Tokens: tokens}
}

func forbiddenScanWorkers(opts ForbiddenSQLOptions, taskCount int) int {
	if taskCount <= 1 || !opts.EnableParallelRead {
		return 1
	}
	workers := opts.ParallelWorkers
	if workers < 1 {
		workers = 1
	}
	cpuCap := runtime.GOMAXPROCS(0)
	if cpuCap < 1 {
		cpuCap = 1
	}
	if workers > cpuCap {
		workers = cpuCap
	}
	if workers > taskCount {
		workers = taskCount
	}
	return workers
}

// ValidateRepoHygiene reports optional helper-file gaps without blocking SQL
// authoring. These files improve repository consistency, but are not part of
// the database object's executable contract.
func ValidateRepoHygiene(report *output.Report, opts RepoHygieneOptions) error {
	if err := reportOptionalFile(report, opts.RepoPath("constraints_naming.md"), "constraints_naming.md", nil); err != nil {
		return err
	}
	if err := reportOptionalFile(report, opts.RepoPath(".gitignore"), ".gitignore", []string{"out/**", "*.output.json"}); err != nil {
		return err
	}
	if err := reportOptionalFile(report, opts.RepoPath(".gitattributes"), ".gitattributes", []string{"* text=auto eol=lf", "*.sql text eol=lf", "*.md text eol=lf", "*.go text eol=lf"}); err != nil {
		return err
	}

	return nil
}

func reportOptionalFile(report *output.Report, absPath, relPath string, recommended []string) error {
	raw, err := os.ReadFile(absPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			report.AddPass(output.Item{
				Rule:     "OPTIONAL_REPO_FILE_MISSING",
				Severity: 1,
				File:     relPath,
				Message:  relPath + " is optional and missing; SQL validation continues",
				Reason:   "OPTIONAL_GUIDELINE",
				Action:   "WARN",
			})
			return nil
		}
		return err
	}
	content := string(raw)
	for _, token := range recommended {
		if !strings.Contains(content, token) {
			report.AddPass(output.Item{
				Rule:     "OPTIONAL_REPO_FILE_CONTENT",
				Severity: 1,
				File:     relPath,
				Message:  relPath + " is missing recommended entry: " + token,
				Reason:   "OPTIONAL_GUIDELINE",
				Action:   "WARN",
			})
		}
	}
	return nil
}
