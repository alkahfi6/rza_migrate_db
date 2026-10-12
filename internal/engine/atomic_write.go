package engine

import (
	"fmt"
	"os"
	"path/filepath"
)

func atomicWriteSQLFile(path string, content []byte) error {
	dir := filepath.Dir(path)
	tmp, err := os.CreateTemp(dir, ".rza-sqlfix-*.tmp")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	committed := false
	defer func() {
		_ = tmp.Close()
		if !committed {
			_ = os.Remove(tmpName)
		}
	}()
	if err := tmp.Chmod(0o644); err != nil {
		return err
	}
	if _, err := tmp.Write(content); err != nil {
		return err
	}
	if err := tmp.Sync(); err != nil {
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	if err := replaceFileAtomic(tmpName, path); err != nil {
		return fmt.Errorf("atomic replace %s: %w", filepath.Clean(path), err)
	}
	committed = true
	return nil
}
