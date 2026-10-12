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

const desiredCollectCacheVersion = 4

type desiredCollectCache struct {
	Version      int             `json:"version"`
	ParserSig    string          `json:"parser_signature,omitempty"`
	SnapshotHash string          `json:"snapshot_hash"`
	Objects      []desiredObject `json:"objects"`
}

func (e Engine) desiredCollectCachePath() string {
	return filepath.Join(e.cfg.RepoRoot, e.cfg.OutDir, "desired_collect_cache.json")
}

func (e Engine) loadDesiredCollectCache(snapshotHash string) ([]desiredObject, bool, error) {
	path := e.desiredCollectCachePath()
	raw, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, false, nil
		}
		return nil, false, err
	}

	var cache desiredCollectCache
	if err := json.Unmarshal(raw, &cache); err != nil {
		return nil, false, err
	}
	if cache.Version != desiredCollectCacheVersion {
		return nil, false, nil
	}
	if cache.ParserSig != parserSignature {
		return nil, false, nil
	}
	if cache.SnapshotHash != snapshotHash {
		return nil, false, nil
	}
	return cloneDesiredObjects(cache.Objects), true, nil
}

func (e Engine) saveDesiredCollectCache(snapshotHash string, objects []desiredObject) error {
	cache := desiredCollectCache{
		Version:      desiredCollectCacheVersion,
		ParserSig:    parserSignature,
		SnapshotHash: snapshotHash,
		Objects:      cloneDesiredObjects(objects),
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
	return os.WriteFile(e.desiredCollectCachePath(), payload, 0o644)
}

func buildDesiredSnapshotHash(tasks []desiredCollectTask) (string, error) {
	return buildDesiredSnapshotHashWithReader(tasks, os.ReadFile)
}

func buildDesiredSnapshotHashWithReader(tasks []desiredCollectTask, readFile func(string) ([]byte, error)) (string, error) {
	type contentEntry struct {
		rel  string
		hash string
	}
	entries := make([]contentEntry, 0, len(tasks))
	for _, t := range tasks {
		raw, err := readFile(t.AbsFile)
		if err != nil {
			return "", err
		}
		sum := sha256.Sum256(raw)
		entries = append(entries, contentEntry{rel: t.RelFile, hash: hex.EncodeToString(sum[:])})
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].rel < entries[j].rel })

	h := sha256.New()
	for _, entry := range entries {
		_, _ = h.Write([]byte(entry.rel))
		_, _ = h.Write([]byte{0})
		_, _ = h.Write([]byte(entry.hash))
		_, _ = h.Write([]byte{0})
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func cloneDesiredObjects(input []desiredObject) []desiredObject {
	if len(input) == 0 {
		return []desiredObject{}
	}
	out := make([]desiredObject, len(input))
	copy(out, input)
	return out
}
