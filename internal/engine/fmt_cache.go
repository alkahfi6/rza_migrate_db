package engine

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const fmtCacheVersion = 10

type fmtCache struct {
	Version int                      `json:"version"`
	Entries map[string]fmtCacheEntry `json:"entries"`
}

type fmtCacheEntry struct {
	Size        int64  `json:"size"`
	ModUnixNano int64  `json:"mod_unix_nano"`
	ContentHash string `json:"content_hash"`
}

func sqlContentHash(raw []byte) string {
	sum := sha256.Sum256(raw)
	return hex.EncodeToString(sum[:])
}

func emptyFmtCache() fmtCache {
	return fmtCache{
		Version: fmtCacheVersion,
		Entries: map[string]fmtCacheEntry{},
	}
}

func (e Engine) fmtCachePath() string {
	return filepath.Join(e.cfg.RepoRoot, e.cfg.OutDir, "fmt_cache.json")
}

func (e Engine) loadFmtCache() (fmtCache, error) {
	path := e.fmtCachePath()
	raw, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return emptyFmtCache(), nil
		}
		return emptyFmtCache(), err
	}

	var cache fmtCache
	if err := json.Unmarshal(raw, &cache); err != nil {
		return emptyFmtCache(), err
	}
	if cache.Version != fmtCacheVersion {
		return emptyFmtCache(), nil
	}
	if cache.Entries == nil {
		cache.Entries = map[string]fmtCacheEntry{}
	}
	return cache, nil
}

func (e Engine) saveFmtCache(cache fmtCache) error {
	if cache.Entries == nil {
		cache.Entries = map[string]fmtCacheEntry{}
	}
	cache.Version = fmtCacheVersion

	outDir := e.cfg.RepoPath(e.cfg.OutDir)
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return err
	}

	payload, err := json.MarshalIndent(cache, "", "  ")
	if err != nil {
		return err
	}
	payload = append(payload, '\n')
	return os.WriteFile(e.fmtCachePath(), payload, 0o644)
}
