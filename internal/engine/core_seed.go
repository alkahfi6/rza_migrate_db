package engine

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"rza_migrate_db/internal/output"
)

//go:embed assets/000000__core__mig__audit_init.up.sql
var embeddedCoreAuditUp []byte

//go:embed assets/000000__core__mig__audit_init.down.sql
var embeddedCoreAuditDown []byte

func (e Engine) seedCoreAuditForBootstrap(report *output.Report) error {
	migrationsRoot := e.cfg.RepoPath(e.cfg.MigrationsDir)
	upPath := filepath.Join(migrationsRoot, coreAuditUpName)
	downPath := filepath.Join(migrationsRoot, coreAuditDownName)
	upExists, err := regularFileExists(upPath)
	if err != nil {
		return err
	}
	downExists, err := regularFileExists(downPath)
	if err != nil {
		return err
	}
	if upExists && downExists {
		return nil
	}
	if upExists != downExists {
		report.AddFail(output.Item{
			Rule: "CORE_AUDIT_PARTIAL", Severity: 3, File: e.cfg.MigrationsDir,
			Message:    "bootstrap found only one core audit direction; existing files were not changed",
			Suggestion: "restore the reviewed core up/down pair or remove the never-applied partial file and retry bootstrap",
			Reason:     "REQUIRES_MANUAL", Action: "fail",
		})
		return nil
	}
	if err := os.MkdirAll(migrationsRoot, 0o755); err != nil {
		return err
	}
	created := make([]string, 0, 2)
	for _, item := range []struct {
		path string
		body []byte
	}{{upPath, embeddedCoreAuditUp}, {downPath, embeddedCoreAuditDown}} {
		file, err := os.OpenFile(item.path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
		if err != nil {
			_ = removeFiles(created)
			return err
		}
		if _, err := file.Write(item.body); err != nil {
			_ = file.Close()
			_ = os.Remove(item.path)
			_ = removeFiles(created)
			return err
		}
		if err := file.Sync(); err != nil {
			_ = file.Close()
			_ = os.Remove(item.path)
			_ = removeFiles(created)
			return err
		}
		if err := file.Close(); err != nil {
			_ = os.Remove(item.path)
			_ = removeFiles(created)
			return err
		}
		created = append(created, item.path)
	}
	report.AddPass(output.Item{
		Rule: "CORE_AUDIT_SEEDED", Severity: 0, File: e.cfg.MigrationsDir,
		Message: "bootstrap seeded the deterministic core audit up/down pair",
		Reason:  "CREATED", Action: "ok",
	})
	return nil
}

func regularFileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if !info.Mode().IsRegular() {
		return false, fmt.Errorf("expected regular file: %s", path)
	}
	return true, nil
}
