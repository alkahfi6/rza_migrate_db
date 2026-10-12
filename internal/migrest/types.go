package migrest

import (
	"context"
)

// VersionState is current migration version state from golang-migrate.
type VersionState struct {
	Version *uint64 `json:"version"`
	Dirty   bool    `json:"dirty"`
}

// OperationResult captures before/after state for one migration API call.
type OperationResult struct {
	Operation      string
	OperationID    string
	FileName       string
	RequestedSteps int
	NoChange       bool
	Before         VersionState
	After          VersionState
	AppliedFiles   []AppliedMigrationFile
	ManifestSHA256 string
}

// AppliedMigrationFile is the immutable file-level evidence for one operation.
type AppliedMigrationFile struct {
	Version   uint64 `json:"version"`
	FileName  string `json:"file_name"`
	Direction string `json:"direction"`
	SHA256    string `json:"sha256"`
}

// PendingMigration describes one up migration file that has not been applied yet.
type PendingMigration struct {
	FileName string
	Version  uint64
	Category string
	Object   string
	VK       int
	IsCore   bool
}

// PendingState is read-only snapshot of current version and remaining up files.
type PendingState struct {
	Current VersionState
	Files   []PendingMigration
}

// Migrator abstracts golang-migrate runtime for testability.
type Migrator interface {
	Up() error
	Down() error
	Steps(n int) error
	Force(version int) error
	Version() (version uint, dirty bool, err error)
	Close() (sourceErr error, databaseErr error)
}

// MigratorFactory creates a fresh migrator per request.
type MigratorFactory func() (Migrator, error)

// LockManager acquires/release cross-process lock for migration operations.
type LockManager interface {
	Acquire(ctx context.Context) (release func() error, err error)
}
