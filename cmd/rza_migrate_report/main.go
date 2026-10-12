package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"rza_migrate_db/internal/buildinfo"
	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/migreport"
)

func main() {
	exitCode := run(os.Args)
	os.Exit(exitCode)
}

func run(args []string) int {
	if len(args) > 0 {
		if err := config.ValidateExecutableName(args[0], config.DefaultMigrateReportExecutableNameLock()); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return 1
		}
	}
	if len(args) == 2 && (args[1] == "--build-info" || args[1] == "-build-info") {
		info := buildinfo.Current()
		fmt.Printf("version=%s\nbuild_time=%s\nsource_sha256=%s\n", info.Version, info.BuildTime, info.SourceSHA256)
		return 0
	}

	fs := flag.NewFlagSet("rza_migrate_report", flag.ContinueOnError)
	repo := fs.String("repo", "", "repository root containing migrations/ directory")
	outDir := fs.String("out", "", "migration audit directory override (default <docs>/migrations; CSV files in datasets/)")
	docsDir := fs.String("docs", "", "root directory for database documentation (default docs/database)")
	pauseOnExit := fs.String("pause-on-exit", "", "pause before process exit (auto true on interactive double-click): true|false")
	fs.SetOutput(os.Stderr)

	flagArgs := []string{}
	if len(args) > 1 {
		flagArgs = args[1:]
	}
	if err := fs.Parse(flagArgs); err != nil {
		return 1
	}
	if fs.NArg() != 0 {
		fmt.Fprintf(os.Stderr, "error: unexpected positional arguments: %s\n", strings.Join(fs.Args(), " "))
		return 1
	}

	isDoubleClicked := len(args) <= 1 && migreport.IsInteractiveConsole()
	pauseEnabled := isDoubleClicked
	if strings.TrimSpace(*pauseOnExit) != "" || strings.TrimSpace(os.Getenv("RZA_PAUSE_ON_EXIT")) != "" {
		val := strings.ToLower(strings.TrimSpace(*pauseOnExit))
		if val == "" {
			val = strings.ToLower(strings.TrimSpace(os.Getenv("RZA_PAUSE_ON_EXIT")))
		}
		switch val {
		case "1", "true", "yes", "on":
			pauseEnabled = true
		case "0", "false", "no", "off":
			pauseEnabled = false
		default:
			fmt.Fprintf(os.Stderr, "error: -pause-on-exit must be true or false, got %q\n", val)
			return 1
		}
	}

	arg0 := ""
	if len(args) > 0 {
		arg0 = args[0]
	}
	opts := migreport.Options{
		RepoRoot:    config.DiscoverRepoRoot(*repo, arg0),
		OutputDir:   *outDir,
		DocsDir:     *docsDir,
		PauseOnExit: pauseEnabled,
	}

	exitCode, err := migreport.Run(opts)
	if err != nil && exitCode == 0 {
		exitCode = 1
	}
	return exitCode
}
