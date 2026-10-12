package engine

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const patchCollectCacheVersion = 2

type patchCollectCache struct {
	Version      int           `json:"version"`
	SnapshotHash string        `json:"snapshot_hash"`
	Sources      []patchSource `json:"sources"`
}

func (e Engine) patchCollectCachePath() string {
	return filepath.Join(e.cfg.RepoRoot, e.cfg.OutDir, "patch_collect_cache.json")
}

func (e Engine) loadPatchCollectCache(snapshotHash string) ([]patchSource, bool, error) {
	path := e.patchCollectCachePath()
	raw, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, false, nil
		}
		return nil, false, err
	}

	var cache patchCollectCache
	if err := json.Unmarshal(raw, &cache); err != nil {
		return nil, false, err
	}
	if cache.Version != patchCollectCacheVersion {
		return nil, false, nil
	}
	if cache.SnapshotHash != snapshotHash {
		return nil, false, nil
	}
	return clonePatchSources(cache.Sources), true, nil
}

func (e Engine) savePatchCollectCache(snapshotHash string, sources []patchSource) error {
	cache := patchCollectCache{
		Version:      patchCollectCacheVersion,
		SnapshotHash: snapshotHash,
		Sources:      clonePatchSources(sources),
	}

	outDir := e.cfg.RepoPath(e.cfg.OutDir)
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return err
	}

	payload, err := json.MarshalIndent(cache, "", "  ")
	if err != nil {
		return err
	}
	payload = append(payload, '\n')
	return os.WriteFile(e.patchCollectCachePath(), payload, 0o644)
}

func buildPatchCollectSnapshotHash(tasks []patchCollectTask) (string, error) {
	type statEntry struct {
		canonical string
		relDir    string
		readme    string
		up        string
		down      string
	}
	stats := make([]statEntry, 0, len(tasks))
	for _, task := range tasks {
		readmeSig, err := optionalFileStatSignature(task.ReadmeAbs)
		if err != nil {
			return "", err
		}
		upSig, err := optionalFileStatSignature(task.UpAbs)
		if err != nil {
			return "", err
		}
		downSig, err := optionalFileStatSignature(task.DownAbs)
		if err != nil {
			return "", err
		}
		stats = append(stats, statEntry{
			canonical: task.Canonical,
			relDir:    task.RelDir,
			readme:    readmeSig,
			up:        upSig,
			down:      downSig,
		})
	}
	sort.Slice(stats, func(i, j int) bool {
		if stats[i].canonical == stats[j].canonical {
			return stats[i].relDir < stats[j].relDir
		}
		return stats[i].canonical < stats[j].canonical
	})

	var b strings.Builder
	for _, s := range stats {
		_, _ = fmt.Fprintf(
			&b,
			"%s|%s|readme:%s|up:%s|down:%s\n",
			s.canonical,
			s.relDir,
			s.readme,
			s.up,
			s.down,
		)
	}
	sum := sha256.Sum256([]byte(b.String()))
	return hex.EncodeToString(sum[:]), nil
}

func optionalFileStatSignature(path string) (string, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "missing", nil
		}
		return "", err
	}
	sum := sha256.Sum256(raw)
	return hex.EncodeToString(sum[:]), nil
}

func clonePatchSources(input []patchSource) []patchSource {
	if len(input) == 0 {
		return []patchSource{}
	}
	out := make([]patchSource, len(input))
	copy(out, input)
	return out
}
