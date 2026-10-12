package engine

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"rza_migrate_db/internal/output"
)

// repositorySnapshot is an immutable, content-addressed view of SQL inputs for
// one engine run. Every SQL file is read and SHA-256 hashed exactly once here;
// validators may then safely share the bytes without trusting mtime or size.
type repositorySnapshot struct {
	byRel      map[string]repositorySnapshotEntry
	orderedRel []string
	totalBytes int64
}

type repositorySnapshotEntry struct {
	RelPath string
	Raw     []byte
	SHA256  string
}

type checkAttestationFile struct {
	Path   string `json:"path"`
	SHA256 string `json:"sha256"`
	Bytes  int    `json:"bytes"`
}

type checkAttestation struct {
	SchemaVersion      int                    `json:"schema_version"`
	Algorithm          string                 `json:"algorithm"`
	Complete           bool                   `json:"complete"`
	Profile            string                 `json:"profile"`
	CheckMode          bool                   `json:"check_mode"`
	GenerationEnabled  bool                   `json:"generation_enabled"`
	DatabaseConnection bool                   `json:"database_connection"`
	Validation         map[string]string      `json:"validation"`
	Scope              []string               `json:"scope"`
	AggregateSHA256    string                 `json:"aggregate_sha256"`
	FileCount          int                    `json:"file_count"`
	TotalBytes         int64                  `json:"total_bytes"`
	Files              []checkAttestationFile `json:"files"`
}

type runState struct {
	mu       sync.Mutex
	snapshot *repositorySnapshot
}

func (e Engine) resetRepositorySnapshot() {
	if e.state == nil {
		return
	}
	e.state.mu.Lock()
	e.state.snapshot = nil
	e.state.mu.Unlock()
}

func (e Engine) ensureRepositorySnapshot(report *output.Report) error {
	if e.state == nil {
		return nil
	}
	e.state.mu.Lock()
	defer e.state.mu.Unlock()
	if e.state.snapshot != nil {
		return nil
	}

	started := time.Now()
	snapshot, err := e.scanRepositorySnapshot()
	if err != nil {
		return err
	}
	e.state.snapshot = snapshot
	if report != nil {
		report.Notes["repository_snapshot_enabled"] = true
		report.Notes["repository_snapshot_files"] = len(snapshot.orderedRel)
		report.Notes["repository_snapshot_bytes"] = snapshot.totalBytes
		e.recordPerf(report, "repository_snapshot.scan", started)
	}
	return nil
}

func (e Engine) scanRepositorySnapshot() (*repositorySnapshot, error) {
	rels := make([]string, 0)
	for _, rootRel := range []string{e.cfg.DbObjectDir, e.cfg.MigrationsDir, e.cfg.PatchDir} {
		files, err := listFilesRecursive(e.cfg.RepoPath(rootRel))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return nil, err
		}
		for _, abs := range files {
			if !strings.EqualFold(filepath.Ext(abs), ".sql") {
				continue
			}
			rels = append(rels, e.toRelPath(abs))
		}
	}
	sort.Strings(rels)

	entries := make([]repositorySnapshotEntry, len(rels))
	type task struct{ index int }
	jobs := make(chan task)
	var wg sync.WaitGroup
	var firstErr error
	var errMu sync.Mutex
	workers := e.readParallelWorkers(len(rels))
	if workers < 1 {
		workers = 1
	}
	read := func(index int) {
		raw, err := os.ReadFile(e.cfg.RepoPath(rels[index]))
		if err != nil {
			errMu.Lock()
			if firstErr == nil {
				firstErr = err
			}
			errMu.Unlock()
			return
		}
		sum := sha256.Sum256(raw)
		entries[index] = repositorySnapshotEntry{
			RelPath: rels[index],
			Raw:     raw,
			SHA256:  hex.EncodeToString(sum[:]),
		}
	}
	if workers == 1 {
		for i := range rels {
			read(i)
		}
	} else {
		for i := 0; i < workers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for item := range jobs {
					read(item.index)
				}
			}()
		}
		for i := range rels {
			jobs <- task{index: i}
		}
		close(jobs)
		wg.Wait()
	}
	if firstErr != nil {
		return nil, firstErr
	}

	snapshot := &repositorySnapshot{
		byRel:      make(map[string]repositorySnapshotEntry, len(entries)),
		orderedRel: append([]string(nil), rels...),
	}
	for _, entry := range entries {
		snapshot.byRel[entry.RelPath] = entry
		snapshot.totalBytes += int64(len(entry.Raw))
	}
	return snapshot, nil
}

func (e Engine) readRepositoryFile(path string) ([]byte, error) {
	if e.state != nil && e.state.snapshot != nil {
		rel := filepath.ToSlash(path)
		if filepath.IsAbs(path) {
			if computed, err := filepath.Rel(e.cfg.RepoRoot, path); err == nil {
				rel = filepath.ToSlash(computed)
			}
		}
		if entry, ok := e.state.snapshot.byRel[rel]; ok {
			return entry.Raw, nil
		}
		if e.isManagedSQLRelativePath(rel) {
			return nil, &os.PathError{Op: "read snapshot", Path: rel, Err: os.ErrNotExist}
		}
	}
	if !filepath.IsAbs(path) {
		path = e.cfg.RepoPath(path)
	}
	return os.ReadFile(path)
}

func (e Engine) isManagedSQLRelativePath(rel string) bool {
	if !strings.EqualFold(filepath.Ext(rel), ".sql") {
		return false
	}
	rel = filepath.ToSlash(strings.TrimPrefix(rel, "./"))
	for _, root := range []string{e.cfg.DbObjectDir, e.cfg.MigrationsDir, e.cfg.PatchDir} {
		root = strings.Trim(filepath.ToSlash(root), "/")
		if strings.EqualFold(rel, root) || strings.HasPrefix(strings.ToLower(rel), strings.ToLower(root)+"/") {
			return true
		}
	}
	return false
}

func (e Engine) snapshotHashForRelativePaths(paths []string) (string, bool) {
	if e.state == nil || e.state.snapshot == nil {
		return "", false
	}
	ordered := append([]string(nil), paths...)
	sort.Strings(ordered)
	h := sha256.New()
	for _, rel := range ordered {
		entry, ok := e.state.snapshot.byRel[filepath.ToSlash(rel)]
		if !ok {
			return "", false
		}
		_, _ = h.Write([]byte(filepath.ToSlash(rel)))
		_, _ = h.Write([]byte{0})
		_, _ = h.Write([]byte(entry.SHA256))
		_, _ = h.Write([]byte{0})
	}
	return hex.EncodeToString(h.Sum(nil)), true
}

// recordCheckAttestation exposes the exact immutable SQL snapshot verified by
// the developer checker. The aggregate covers both each case-sensitive path
// and its raw SHA-256, so additions, removals, renames, case changes, and byte
// changes all produce a different value.
func (e Engine) recordCheckAttestation(report *output.Report) error {
	if report == nil || e.state == nil || e.state.snapshot == nil {
		return errors.New("repository snapshot is unavailable for check attestation")
	}
	snapshot := e.state.snapshot
	aggregate, ok := e.snapshotHashForRelativePaths(snapshot.orderedRel)
	if !ok {
		return errors.New("repository snapshot is incomplete for check attestation")
	}
	files := make([]checkAttestationFile, 0, len(snapshot.orderedRel))
	for _, rel := range snapshot.orderedRel {
		entry, exists := snapshot.byRel[rel]
		if !exists || entry.RelPath != rel || len(entry.SHA256) != sha256.Size*2 {
			return errors.New("repository snapshot entry is incomplete for check attestation: " + rel)
		}
		files = append(files, checkAttestationFile{
			Path:   entry.RelPath,
			SHA256: entry.SHA256,
			Bytes:  len(entry.Raw),
		})
	}
	report.Notes["check_attestation"] = checkAttestation{
		SchemaVersion:      2,
		Algorithm:          "SHA-256",
		Complete:           true,
		Profile:            e.cfg.ValidationProfile,
		CheckMode:          e.cfg.EnableFmtCheck,
		GenerationEnabled:  false,
		DatabaseConnection: false,
		Validation: map[string]string{
			"db_object":  "full",
			"patch":      "full",
			"migrations": e.cfg.MigrationValidation,
		},
		Scope:           []string{e.cfg.DbObjectDir + "/**/*.sql", e.cfg.MigrationsDir + "/**/*.sql", e.cfg.PatchDir + "/**/*.sql"},
		AggregateSHA256: aggregate,
		FileCount:       len(files),
		TotalBytes:      snapshot.totalBytes,
		Files:           files,
	}
	return nil
}

func (e Engine) migrationsSnapshotHash(root string, entries []os.DirEntry) (string, []string, error) {
	if e.state == nil || e.state.snapshot == nil {
		return buildMigrationsSnapshotHash(root, entries)
	}
	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			names = append(names, entry.Name())
		}
	}
	sort.Strings(names)
	h := sha256.New()
	for _, name := range names {
		rel := filepath.ToSlash(filepath.Join(e.cfg.MigrationsDir, name))
		entry, ok := e.state.snapshot.byRel[rel]
		contentHash := ""
		if ok {
			contentHash = entry.SHA256
		} else {
			raw, err := os.ReadFile(filepath.Join(root, name))
			if err != nil {
				return "", nil, err
			}
			sum := sha256.Sum256(raw)
			contentHash = hex.EncodeToString(sum[:])
		}
		_, _ = h.Write([]byte(name))
		_, _ = h.Write([]byte{0})
		_, _ = h.Write([]byte(contentHash))
		_, _ = h.Write([]byte{0})
	}
	return hex.EncodeToString(h.Sum(nil)), names, nil
}

func (e Engine) desiredSnapshotHash(tasks []desiredCollectTask) (string, error) {
	if e.state == nil || e.state.snapshot == nil {
		return buildDesiredSnapshotHash(tasks)
	}
	paths := make([]string, 0, len(tasks))
	for _, task := range tasks {
		paths = append(paths, task.RelFile)
	}
	if hash, ok := e.snapshotHashForRelativePaths(paths); ok {
		return hash, nil
	}
	return buildDesiredSnapshotHashWithReader(tasks, e.readRepositoryFile)
}
