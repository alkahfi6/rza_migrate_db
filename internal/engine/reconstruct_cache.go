package engine

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sort"
)

// Version 4 invalidates states produced before dynamic SQL constraint
// reconstruction masked string literals, decoded executable batches, and
// stopped cross-object propagation from reapplying local table transitions.
const reconstructCacheVersion = 4

type reconstructCache struct {
	Version      int                      `json:"version"`
	ParserSig    string                   `json:"parser_signature,omitempty"`
	SnapshotHash string                   `json:"snapshot_hash"`
	History      map[string]historyObject `json:"history"`
	Existing     []string                 `json:"existing_files"`
	MaxVersion   string                   `json:"max_version"`
}

func (e Engine) reconstructCachePath() string {
	return filepath.Join(e.cfg.RepoRoot, e.cfg.OutDir, "reconstruct_history_cache.json")
}

func (e Engine) loadReconstructCache(snapshotHash string) (reconstructCache, bool, error) {
	path := e.reconstructCachePath()
	raw, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return reconstructCache{}, false, nil
		}
		return reconstructCache{}, false, err
	}

	var cache reconstructCache
	if err := json.Unmarshal(raw, &cache); err != nil {
		return reconstructCache{}, false, err
	}
	if cache.Version != reconstructCacheVersion {
		return reconstructCache{}, false, nil
	}
	if cache.ParserSig != parserSignature {
		return reconstructCache{}, false, nil
	}
	if cache.SnapshotHash != snapshotHash {
		return reconstructCache{}, false, nil
	}
	if cache.History == nil {
		cache.History = map[string]historyObject{}
	}
	if cache.Existing == nil {
		cache.Existing = []string{}
	}
	return cache, true, nil
}

func (e Engine) saveReconstructCache(snapshotHash string, history map[string]historyObject, existingFiles map[string]struct{}, maxVersion string) error {
	if history == nil {
		history = map[string]historyObject{}
	}
	existing := make([]string, 0, len(existingFiles))
	for name := range existingFiles {
		existing = append(existing, name)
	}
	sort.Strings(existing)

	cache := reconstructCache{
		Version:      reconstructCacheVersion,
		ParserSig:    parserSignature,
		SnapshotHash: snapshotHash,
		History:      cloneHistoryStateMap(history),
		Existing:     existing,
		MaxVersion:   maxVersion,
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
	return os.WriteFile(e.reconstructCachePath(), payload, 0o644)
}

func buildMigrationsSnapshotHash(root string, entries []os.DirEntry) (string, []string, error) {
	return buildMigrationsSnapshotHashWithReader(root, entries, os.ReadFile)
}

func buildMigrationsSnapshotHashWithReader(root string, entries []os.DirEntry, readFile func(string) ([]byte, error)) (string, []string, error) {
	names := make([]string, 0, len(entries))
	type contentEntry struct {
		name string
		hash string
	}
	contents := make([]contentEntry, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		names = append(names, name)
		raw, err := readFile(filepath.Join(root, name))
		if err != nil {
			return "", nil, err
		}
		sum := sha256.Sum256(raw)
		contents = append(contents, contentEntry{name: name, hash: hex.EncodeToString(sum[:])})
	}
	sort.Strings(names)
	sort.Slice(contents, func(i, j int) bool { return contents[i].name < contents[j].name })

	h := sha256.New()
	for _, entry := range contents {
		_, _ = h.Write([]byte(entry.name))
		_, _ = h.Write([]byte{0})
		_, _ = h.Write([]byte(entry.hash))
		_, _ = h.Write([]byte{0})
	}
	return hex.EncodeToString(h.Sum(nil)), names, nil
}
