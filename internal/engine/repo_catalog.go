package engine

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type repoCatalogManifest struct {
	Catalog struct {
		Self      string   `yaml:"self"`
		Collation string   `yaml:"collation"`
		External  []string `yaml:"external"`
	} `yaml:"catalog"`
}

func (e Engine) loadRepoCatalogManifest() (repoCatalogManifest, error) {
	var manifest repoCatalogManifest
	path := e.cfg.RepoPath("rza_migrate.yaml")
	raw, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return manifest, nil
		}
		return manifest, err
	}
	if err := yaml.Unmarshal(raw, &manifest); err != nil {
		return manifest, err
	}
	manifest.Catalog.Self = strings.TrimSpace(manifest.Catalog.Self)
	return manifest, nil
}

func (e Engine) desiredStateObjectNames(root string) (map[string]bool, error) {
	files, err := listFilesRecursive(root)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return map[string]bool{}, nil
		}
		return nil, err
	}
	objects := map[string]bool{}
	for _, file := range files {
		if !strings.EqualFold(filepath.Ext(file), ".sql") {
			continue
		}
		rel := strings.ToLower(strings.ReplaceAll(e.toRelPath(file), "\\", "/"))
		if !strings.Contains(rel, "/tables/") && !strings.Contains(rel, "/views/") &&
			!strings.Contains(rel, "/functions/") && !strings.Contains(rel, "/procedures/") &&
			!strings.Contains(rel, "/types/") {
			continue
		}
		objects[strings.ToLower(strings.TrimSuffix(filepath.Base(file), filepath.Ext(file)))] = true
	}
	return objects, nil
}
