package engine

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"rza_migrate_db/internal/output"
)

func (e Engine) emitPlannedMigrations(report *output.Report, ops []plannedMigration, existingFiles map[string]struct{}, maxVersion string) error {
	if len(ops) == 0 {
		report.SetNewMigrationFiles([]string{})
		return nil
	}

	baseTimestamp := time.Now().Format("20060102150405")
	if len(maxVersion) == 17 {
		lastBase := maxVersion[:14]
		if lastBase > baseTimestamp {
			baseTimestamp = lastBase
		}
	}

	reserved := make(map[string]struct{}, len(existingFiles))
	for name := range existingFiles {
		reserved[name] = struct{}{}
	}

	allocated, pendingCount, _, allocErr := e.allocateMigrationBatch(ops, baseTimestamp, 0, reserved)
	if allocErr != nil {
		report.AddFail(output.Item{
			Rule:     "PAIR_SEQ_EXCEEDED",
			Severity: 3,
			File:     e.cfg.MigrationsDir,
			Message:  allocErr.Error(),
			Reason:   "REQUIRES_MANUAL",
			Action:   "fail",
		})
		return nil
	}

	files, upFiles := e.renderAllocatedMigrationFiles(report, allocated)
	stagingDir, stagedPaths, err := e.writeEmitStaging(files)
	if err != nil {
		return err
	}
	report.Notes["emit_staging_dir"] = filepath.ToSlash(stagingDir)
	defer func() {
		_ = os.RemoveAll(stagingDir)
	}()

	tempVerify, verifyErr := e.verifyStagedMigrations(stagingDir)
	if verifyErr != nil {
		return verifyErr
	}
	if tempVerify.Fail.Count > 0 {
		if copies, copyErr := e.copyFailedEmitFiles(stagedPaths); copyErr != nil {
			report.Notes["gen_post_emit_failed_copy_error"] = copyErr.Error()
		} else if len(copies) > 0 {
			report.Notes["gen_post_emit_failed_files"] = copies
		}
		for _, item := range tempVerify.Fail.Items {
			report.AddFail(item)
		}
		report.SetNewMigrationFiles([]string{})
		return nil
	}

	commitDir := e.cfg.RepoPath(e.cfg.MigrationsDir)
	if e.cfg.DraftMode {
		commitDir = e.cfg.RepoPath(e.cfg.DraftDir)
	}
	committed, err := commitStagedFiles(stagedPaths, commitDir)
	if err != nil {
		_ = removeFiles(committed)
		return err
	}

	report.SetNewMigrationFiles(upFiles)
	if e.cfg.DraftMode {
		report.Notes["draft_directory"] = filepath.ToSlash(e.cfg.DraftDir)
		report.Notes["draft_pairs"] = len(allocated)
		report.AddPass(output.Item{
			Rule: "DRAFT_EMIT_OK", Severity: 0, File: e.cfg.DraftDir,
			Message: fmt.Sprintf("generated %d reviewed draft migration pair(s); run `rza_migrate_db seal` to publish them", len(allocated)),
			Reason:  "CREATED", Action: "ok",
		})
	}
	report.Notes["generated_pairs"] = len(allocated)
	if pendingCount > 0 {
		report.Notes["pending_pairs"] = pendingCount
		report.Notes["batch_limit_pairs"] = e.cfg.MaxPairsPerRun
		report.AddPass(output.Item{
			Rule:       "GEN_BATCH_LIMIT_REACHED",
			Severity:   2,
			File:       e.cfg.MigrationsDir,
			Message:    fmt.Sprintf("generated %d pair(s) in this run; %d pair(s) remain (re-run gen to continue)", len(allocated), pendingCount),
			Suggestion: "re-run `rza_migrate_db gen` until pending_pairs = 0",
			Reason:     "DRIFT_DETECTED",
			Action:     "warn",
		})
	}
	if !e.cfg.DraftMode {
		report.AddPass(output.Item{
			Rule:     "GEN_EMIT_OK",
			Severity: 0,
			File:     e.cfg.MigrationsDir,
			Message:  fmt.Sprintf("generated %d migration pair(s)", len(allocated)),
			Reason:   "CREATED",
			Action:   "ok",
		})
	}
	return nil
}

func (e Engine) renderAllocatedMigrationFiles(report *output.Report, allocated []allocatedMigration) ([]generatedFile, []string) {
	files := make([]generatedFile, 0, len(allocated)*2)
	upFiles := make([]string, 0, len(allocated))
	for _, item := range allocated {
		upBody := renderPlannedMigration(item.Op, item.Version, item.UpName, true)
		downBody := renderPlannedMigration(item.Op, item.Version, item.DownName, false)

		files = append(files, generatedFile{Name: item.UpName, Content: upBody})
		files = append(files, generatedFile{Name: item.DownName, Content: downBody})
		upFiles = append(upFiles, item.UpName)

		if item.Op.Category == "patch" {
			report.AddPass(output.Item{
				Rule:                       "PATCH_MAPPING",
				Severity:                   1,
				File:                       item.Op.SourceFile,
				Object:                     item.Op.PatchNameCanonical,
				Message:                    "patch mapped to generated migration",
				Reason:                     "CREATED",
				Action:                     "ok",
				PatchFolderOriginal:        item.Op.PatchFolderOriginal,
				PatchNameCanonical:         item.Op.PatchNameCanonical,
				GeneratedMigrationVersions: []string{item.Version},
			})
		}
	}
	return files, upFiles
}

func (e Engine) writeEmitStaging(files []generatedFile) (string, []string, error) {
	stagingRoot := filepath.Join(e.cfg.RepoPath(e.cfg.OutDir), "emit_staging")
	if err := os.MkdirAll(stagingRoot, 0o755); err != nil {
		return "", nil, err
	}
	stagingDir, err := os.MkdirTemp(stagingRoot, "run-*")
	if err != nil {
		return "", nil, err
	}

	written := make([]string, 0, len(files))
	for _, file := range files {
		if strings.TrimSpace(file.Name) == "" || filepath.Base(file.Name) != file.Name {
			_ = os.RemoveAll(stagingDir)
			return "", nil, fmt.Errorf("invalid generated migration filename for staging: %q", file.Name)
		}
		target := filepath.Join(stagingDir, file.Name)
		if err := os.WriteFile(target, []byte(file.Content), 0o644); err != nil {
			_ = os.RemoveAll(stagingDir)
			return "", nil, err
		}
		written = append(written, target)
	}
	return stagingDir, written, nil
}

func (e Engine) verifyStagedMigrations(stagingDir string) (output.Report, error) {
	verifyParent := filepath.Join(e.cfg.RepoPath(e.cfg.OutDir), "emit_verify")
	if err := os.MkdirAll(verifyParent, 0o755); err != nil {
		return output.Report{}, err
	}
	verifyRoot, err := os.MkdirTemp(verifyParent, "repo-*")
	if err != nil {
		return output.Report{}, err
	}
	verifyRoot, err = filepath.Abs(verifyRoot)
	if err != nil {
		return output.Report{}, err
	}
	defer func() {
		_ = os.RemoveAll(verifyRoot)
	}()

	if err := e.copyRepoForStagedVerify(verifyRoot); err != nil {
		return output.Report{}, err
	}
	if err := copyDirContents(stagingDir, filepath.Join(verifyRoot, e.cfg.MigrationsDir)); err != nil {
		return output.Report{}, err
	}

	verifyCfg := e.cfg
	verifyCfg.RepoRoot = verifyRoot
	verifyCfg.EnableOutputNotes = false
	tempVerify := output.NewReport(e.cfg.OutputNamingPattern, e.cfg.OutputNamingLineEndings)
	tempVerify.Notes["verify_mode"] = "gen_staging"
	if err := New(verifyCfg).runVerifyChecks(&tempVerify, true, nil); err != nil {
		return tempVerify, err
	}
	return tempVerify, nil
}

func (e Engine) copyRepoForStagedVerify(dstRoot string) error {
	for _, dir := range []string{e.cfg.DbObjectDir, e.cfg.MigrationsDir, e.cfg.PatchDir} {
		src := e.cfg.RepoPath(dir)
		dst := filepath.Join(dstRoot, dir)
		if err := copyDirIfExists(src, dst); err != nil {
			return err
		}
	}
	for _, file := range []string{"constraints_naming.md", ".gitignore", ".gitattributes"} {
		src := e.cfg.RepoPath(file)
		if err := copyFileIfExists(src, filepath.Join(dstRoot, file)); err != nil {
			return err
		}
	}
	return nil
}

func (e Engine) commitStagedMigrations(stagedPaths []string) ([]string, error) {
	return commitStagedFiles(stagedPaths, e.cfg.RepoPath(e.cfg.MigrationsDir))
}

func commitStagedFiles(stagedPaths []string, targetRoot string) ([]string, error) {
	if err := os.MkdirAll(targetRoot, 0o755); err != nil {
		return nil, err
	}

	committed := make([]string, 0, len(stagedPaths))
	for _, src := range stagedPaths {
		dst := filepath.Join(targetRoot, filepath.Base(src))
		if _, err := os.Stat(dst); err == nil {
			return committed, fmt.Errorf("migration already exists during staged commit: %s", filepath.Base(dst))
		} else if err != nil && !errors.Is(err, os.ErrNotExist) {
			return committed, err
		}
		if err := os.Rename(src, dst); err != nil {
			return committed, err
		}
		committed = append(committed, dst)
	}
	return committed, nil
}

func (e Engine) runSealDrafts(report *output.Report) error {
	draftRoot := e.cfg.RepoPath(e.cfg.DraftDir)
	entries, err := os.ReadDir(draftRoot)
	if errors.Is(err, os.ErrNotExist) {
		report.AddFail(output.Item{Rule: "DRAFT_NOT_FOUND", Severity: 3, File: e.cfg.DraftDir, Message: "draft migration directory does not exist", Suggestion: "run `rza_migrate_db draft` first", Reason: "REQUIRES_MANUAL", Action: "fail"})
		return nil
	}
	if err != nil {
		return err
	}
	paths := make([]string, 0, len(entries))
	upNames := make([]string, 0, len(entries)/2)
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(strings.ToLower(entry.Name()), ".sql") || filepath.Base(entry.Name()) != entry.Name() {
			report.AddFail(output.Item{Rule: "DRAFT_CONTENT_INVALID", Severity: 3, File: filepath.ToSlash(filepath.Join(e.cfg.DraftDir, entry.Name())), Message: "draft directory must contain only flat .sql migration files", Reason: "REQUIRES_MANUAL", Action: "fail"})
			continue
		}
		if _, statErr := os.Stat(filepath.Join(e.cfg.RepoPath(e.cfg.MigrationsDir), entry.Name())); statErr == nil {
			report.AddFail(output.Item{Rule: "DRAFT_VERSION_COLLISION", Severity: 3, File: filepath.ToSlash(filepath.Join(e.cfg.DraftDir, entry.Name())), Message: "a migration with the same filename already exists", Reason: "REQUIRES_MANUAL", Action: "fail"})
			continue
		} else if !errors.Is(statErr, os.ErrNotExist) {
			return statErr
		}
		paths = append(paths, filepath.Join(draftRoot, entry.Name()))
		if strings.HasSuffix(strings.ToLower(entry.Name()), ".up.sql") {
			upNames = append(upNames, entry.Name())
		}
	}
	if report.Fail.Count > 0 {
		return nil
	}
	if len(paths) == 0 {
		report.AddFail(output.Item{Rule: "DRAFT_EMPTY", Severity: 3, File: e.cfg.DraftDir, Message: "draft directory contains no migration files", Suggestion: "run `rza_migrate_db draft` first", Reason: "REQUIRES_MANUAL", Action: "fail"})
		return nil
	}
	tempVerify, err := e.verifyStagedMigrations(draftRoot)
	if err != nil {
		return err
	}
	if tempVerify.Fail.Count > 0 {
		for _, item := range tempVerify.Fail.Items {
			report.AddFail(item)
		}
		return nil
	}
	committed, err := publishDraftFiles(paths, e.cfg.RepoPath(e.cfg.MigrationsDir))
	if err != nil {
		_ = removeFiles(committed)
		return err
	}
	for _, path := range paths {
		if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
			report.Notes["draft_cleanup_warning"] = err.Error()
		}
	}
	sort.Strings(upNames)
	report.SetNewMigrationFiles(upNames)
	report.Notes["sealed_pairs"] = len(upNames)
	report.AddPass(output.Item{Rule: "DRAFT_SEAL_OK", Severity: 0, File: e.cfg.MigrationsDir, Message: fmt.Sprintf("sealed %d draft migration pair(s) without changing their reviewed bytes", len(upNames)), Reason: "CREATED", Action: "ok"})
	return nil
}

func publishDraftFiles(draftPaths []string, targetRoot string) ([]string, error) {
	if err := os.MkdirAll(targetRoot, 0o755); err != nil {
		return nil, err
	}
	published := make([]string, 0, len(draftPaths))
	for _, source := range draftPaths {
		raw, err := os.ReadFile(source)
		if err != nil {
			return published, err
		}
		target := filepath.Join(targetRoot, filepath.Base(source))
		file, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
		if err != nil {
			return published, err
		}
		if _, err := file.Write(raw); err != nil {
			_ = file.Close()
			_ = os.Remove(target)
			return published, err
		}
		if err := file.Sync(); err != nil {
			_ = file.Close()
			_ = os.Remove(target)
			return published, err
		}
		if err := file.Close(); err != nil {
			_ = os.Remove(target)
			return published, err
		}
		published = append(published, target)
	}
	return published, nil
}

func copyDirIfExists(src, dst string) error {
	info, err := os.Stat(src)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("expected directory: %s", src)
	}
	return copyDirContents(src, dst)
}

func copyDirContents(src, dst string) error {
	if err := os.MkdirAll(dst, 0o755); err != nil {
		return err
	}
	return filepath.WalkDir(src, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}
		target := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		return copyFile(path, target)
	})
}

func copyFileIfExists(src, dst string) error {
	info, err := os.Stat(src)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	if info.IsDir() {
		return fmt.Errorf("expected file: %s", src)
	}
	return copyFile(src, dst)
}

func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0o644)
}
