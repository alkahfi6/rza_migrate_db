package migreport

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// DiscoverRepository locates the repository root containing migrations/ directory.
// Candidates checked in order:
// 1. Explicit repo if provided.
// 2. Executable directory.
// 3. Current working directory.
// 4. Parent directories of cwd and executable dir.
func DiscoverRepository(explicitRepo string) (string, error) {
	if strings.TrimSpace(explicitRepo) != "" {
		clean := filepath.Clean(explicitRepo)
		if hasMigrationsDir(clean) {
			return clean, nil
		}
		if clean == "." && hasMigrationsDir("..") {
			return "..", nil
		}
		return "", fmt.Errorf("migrations directory not found in specified repository: %s", clean)
	}

	// 1. Executable directory
	if execPath, err := os.Executable(); err == nil {
		execDir := filepath.Dir(filepath.Clean(execPath))
		if hasMigrationsDir(execDir) {
			return execDir, nil
		}
		// Check parent of execDir
		parent := filepath.Dir(execDir)
		if parent != execDir && hasMigrationsDir(parent) {
			return parent, nil
		}
	}

	// 2. Current working directory
	cwd, err := os.Getwd()
	if err == nil {
		cleanCwd := filepath.Clean(cwd)
		if hasMigrationsDir(cleanCwd) {
			return cleanCwd, nil
		}
		// Walk up parents of cwd up to 4 levels
		curr := cleanCwd
		for i := 0; i < 4; i++ {
			parent := filepath.Dir(curr)
			if parent == curr {
				break
			}
			if hasMigrationsDir(parent) {
				return parent, nil
			}
			curr = parent
		}
	}

	// Also check "."
	if hasMigrationsDir(".") {
		return ".", nil
	}

	searched := "."
	if cwd != "" {
		searched = cwd
	}
	return "", fmt.Errorf("migrations directory not found: %s", searched)
}

func hasMigrationsDir(dir string) bool {
	info, err := os.Stat(filepath.Join(dir, "migrations"))
	return err == nil && info.IsDir()
}

// IsInteractiveConsole checks if stdin and stdout are interactive character devices.
func IsInteractiveConsole() bool {
	stdinInfo, stdinErr := os.Stdin.Stat()
	stdoutInfo, stdoutErr := os.Stdout.Stat()
	return stdinErr == nil && stdoutErr == nil &&
		(stdinInfo.Mode()&os.ModeCharDevice) != 0 &&
		(stdoutInfo.Mode()&os.ModeCharDevice) != 0
}

// WaitForEnter pauses execution if enabled.
func WaitForEnter(enabled bool) {
	if !enabled {
		return
	}
	fmt.Print("[done] press Enter to exit\n")
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
}
