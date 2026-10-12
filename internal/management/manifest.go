package management

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Object struct {
	Category string `json:"category"`
	Object   string `json:"object"`
	Reason   string `json:"reason,omitempty"`
}

type Manifest struct {
	Schema      string   `json:"schema"`
	GeneratedAt string   `json:"generated_at"`
	Managed     []Object `json:"managed_objects"`
	Unmanaged   []Object `json:"unmanaged_objects"`
}

func (m Manifest) UnmanagedSet() map[string]struct{} {
	out := make(map[string]struct{}, len(m.Unmanaged))
	for _, item := range m.Unmanaged {
		category := strings.ToLower(strings.TrimSpace(item.Category))
		object := strings.TrimSpace(item.Object)
		if category != "" && object != "" {
			out[category+"|"+object] = struct{}{}
		}
	}
	return out
}

func Load(path string) (Manifest, error) {
	raw, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		raw, err = os.ReadFile(path + ".bak")
		if errors.Is(err, os.ErrNotExist) {
			return Manifest{}, nil
		}
	}
	if err != nil {
		return Manifest{}, err
	}
	var manifest Manifest
	if err := json.Unmarshal(raw, &manifest); err != nil {
		return Manifest{}, err
	}
	return manifest, nil
}

func Write(path string, manifest Manifest) error {
	if strings.TrimSpace(manifest.GeneratedAt) == "" {
		manifest.GeneratedAt = time.Now().UTC().Format(time.RFC3339)
	}
	sortObjects(manifest.Managed)
	sortObjects(manifest.Unmanaged)
	raw, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(filepath.Dir(path), ".management-manifest-*")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName)
	if _, err := tmp.Write(append(raw, '\n')); err != nil {
		_ = tmp.Close()
		return err
	}
	if err := tmp.Sync(); err != nil {
		_ = tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	backup := path + ".bak"
	_ = os.Remove(backup)
	if _, err := os.Stat(path); err == nil {
		if err := os.Rename(path, backup); err != nil {
			return err
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}
	if err := os.Rename(tmpName, path); err != nil {
		_ = os.Rename(backup, path)
		return err
	}
	_ = os.Remove(backup)
	return nil
}

func sortObjects(items []Object) {
	sort.Slice(items, func(i, j int) bool {
		if items[i].Category != items[j].Category {
			return items[i].Category < items[j].Category
		}
		return items[i].Object < items[j].Object
	})
}
