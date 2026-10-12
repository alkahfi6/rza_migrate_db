package migrest

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type operationJournalRecord struct {
	OperationID    string                 `json:"operation_id"`
	RequestID      string                 `json:"request_id"`
	Operation      string                 `json:"operation"`
	FileName       string                 `json:"file_name,omitempty"`
	RequestedSteps int                    `json:"requested_steps"`
	Before         VersionState           `json:"before"`
	After          VersionState           `json:"after"`
	Status         string                 `json:"status"`
	Direction      string                 `json:"direction"`
	AppliedFiles   []AppliedMigrationFile `json:"applied_files"`
	ManifestSHA256 string                 `json:"manifest_sha256,omitempty"`
	ErrorMessage   string                 `json:"error_message,omitempty"`
	NIK            string                 `json:"nik,omitempty"`
	TraceID        string                 `json:"trace_id,omitempty"`
	StartedAt      time.Time              `json:"started_at_utc"`
	FinishedAt     *time.Time             `json:"finished_at_utc,omitempty"`
}

type operationJournal interface {
	ValidateAppliedHistory(context.Context, string) error
	Start(context.Context, operationJournalRecord) error
	Finish(context.Context, operationJournalRecord) error
}

type noopOperationJournal struct{}

func (noopOperationJournal) ValidateAppliedHistory(context.Context, string) error { return nil }
func (noopOperationJournal) Start(context.Context, operationJournalRecord) error  { return nil }
func (noopOperationJournal) Finish(context.Context, operationJournalRecord) error { return nil }

type mssqlOperationJournal struct {
	db *sql.DB
}

type appliedHistoryEvidence struct {
	FileName string
	SHA256   string
}

func (j mssqlOperationJournal) ValidateAppliedHistory(ctx context.Context, migrationsDir string) error {
	if j.db == nil {
		return errors.New("operation journal database is unavailable")
	}
	exists, err := journalTableExists(ctx, j.db)
	if err != nil || !exists {
		return err
	}
	rows, err := j.db.QueryContext(ctx, `
SELECT evidence.file_name, evidence.sha256
FROM dbo.mig_operation_journal AS journal
CROSS APPLY OPENJSON(journal.applied_files)
WITH (
  file_name NVARCHAR(260) '$.file_name',
  sha256 CHAR(64) '$.sha256'
) AS evidence
WHERE journal.status IN (N'success', N'no_change')
  AND evidence.file_name IS NOT NULL
  AND evidence.sha256 IS NOT NULL;`)
	if err != nil {
		return fmt.Errorf("read applied migration checksum evidence: %w", err)
	}
	defer rows.Close()
	evidence := make([]appliedHistoryEvidence, 0)
	for rows.Next() {
		var item appliedHistoryEvidence
		if err := rows.Scan(&item.FileName, &item.SHA256); err != nil {
			return fmt.Errorf("scan applied migration checksum evidence: %w", err)
		}
		evidence = append(evidence, item)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("iterate applied migration checksum evidence: %w", err)
	}
	return validateAppliedHistoryFiles(migrationsDir, evidence)
}

func validateAppliedHistoryFiles(migrationsDir string, evidence []appliedHistoryEvidence) error {
	approved := make(map[string]string, len(evidence))
	for _, item := range evidence {
		name := strings.TrimSpace(item.FileName)
		expected := strings.ToLower(strings.TrimSpace(item.SHA256))
		if name == "" || filepath.Base(name) != name || strings.Contains(name, "..") {
			return fmt.Errorf("invalid migration filename in operation journal: %q", item.FileName)
		}
		if len(expected) != 64 {
			return fmt.Errorf("invalid SHA-256 evidence for migration %s", name)
		}
		if previous, exists := approved[name]; exists && previous != expected {
			return fmt.Errorf("conflicting SHA-256 evidence for migration %s", name)
		}
		approved[name] = expected
	}
	for name, expected := range approved {
		raw, err := os.ReadFile(filepath.Join(migrationsDir, name))
		if err != nil {
			return fmt.Errorf("read historically applied migration %s: %w", name, err)
		}
		actual := fmt.Sprintf("%x", sha256.Sum256(raw))
		if actual != expected {
			return fmt.Errorf("immutable migration history mismatch for %s: journal=%s file=%s", name, expected, actual)
		}
	}
	return nil
}

func (j mssqlOperationJournal) Start(ctx context.Context, record operationJournalRecord) error {
	if j.db == nil {
		return errors.New("operation journal database is unavailable")
	}
	exists, err := journalTableExists(ctx, j.db)
	if err != nil {
		return err
	}
	if !exists {
		if record.Before.Version == nil {
			// The core migration creates the journal during the first deployment.
			return nil
		}
		return errors.New("mandatory dbo.mig_operation_journal table is unavailable")
	}
	_, err = j.db.ExecContext(ctx, `
INSERT INTO dbo.mig_operation_journal (
  operation_id, request_id, operation, file_name, status, requested_steps,
  before_version, before_dirty, result_direction, nik, trace_id,
  created_at, updated_at
) VALUES (
  @p1, @p2, @p3, @p4, N'running', @p5,
  @p6, @p7, @p8, @p9, @p10,
  @p11, @p11
);`,
		record.OperationID, record.RequestID, record.Operation, nullString(record.FileName),
		record.RequestedSteps, nullableVersion(record.Before), record.Before.Dirty,
		record.Direction, nullString(record.NIK), nullString(record.TraceID), record.StartedAt.UTC())
	if err != nil {
		return fmt.Errorf("start migration operation journal: %w", err)
	}
	return nil
}

func (j mssqlOperationJournal) Finish(ctx context.Context, record operationJournalRecord) error {
	if j.db == nil {
		return errors.New("operation journal database is unavailable")
	}
	exists, err := journalTableExists(ctx, j.db)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("mandatory dbo.mig_operation_journal table is unavailable after migration")
	}
	reportJSON, err := json.Marshal(record)
	if err != nil {
		return fmt.Errorf("encode migration operation report: %w", err)
	}
	appliedJSON, err := json.Marshal(record.AppliedFiles)
	if err != nil {
		return fmt.Errorf("encode applied migration files: %w", err)
	}
	finishedAt := time.Now().UTC()
	if record.FinishedAt != nil {
		finishedAt = record.FinishedAt.UTC()
	}
	result, err := j.db.ExecContext(ctx, `
UPDATE dbo.mig_operation_journal
SET status=@p2,
    executed_steps=@p3,
    after_version=@p4,
    after_dirty=@p5,
    result_direction=@p6,
    applied_files=@p7,
    error_message=@p8,
    manifest_sha256=@p9,
    report_json=@p10,
    updated_at=@p11,
    finished_at=@p11
WHERE operation_id=@p1;`,
		record.OperationID, record.Status, len(record.AppliedFiles), nullableVersion(record.After),
		record.After.Dirty, record.Direction, string(appliedJSON), nullString(record.ErrorMessage),
		nullString(record.ManifestSHA256), string(reportJSON), finishedAt)
	if err != nil {
		return fmt.Errorf("finish migration operation journal: %w", err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("read migration operation journal update count: %w", err)
	}
	if affected > 0 {
		return nil
	}
	_, err = j.db.ExecContext(ctx, `
INSERT INTO dbo.mig_operation_journal (
  operation_id, request_id, operation, file_name, status, requested_steps,
  executed_steps, before_version, before_dirty, after_version, after_dirty,
  result_direction, applied_files, error_message, nik, trace_id,
  manifest_sha256, report_json, created_at, updated_at, finished_at
) VALUES (
  @p1, @p2, @p3, @p4, @p5, @p6,
  @p7, @p8, @p9, @p10, @p11,
  @p12, @p13, @p14, @p15, @p16,
  @p17, @p18, @p19, @p20, @p20
);`,
		record.OperationID, record.RequestID, record.Operation, nullString(record.FileName),
		record.Status, record.RequestedSteps, len(record.AppliedFiles), nullableVersion(record.Before),
		record.Before.Dirty, nullableVersion(record.After), record.After.Dirty, record.Direction,
		string(appliedJSON), nullString(record.ErrorMessage), nullString(record.NIK),
		nullString(record.TraceID), nullString(record.ManifestSHA256), string(reportJSON),
		record.StartedAt.UTC(), finishedAt)
	if err != nil {
		return fmt.Errorf("insert completed migration operation journal: %w", err)
	}
	return nil
}

func journalTableExists(ctx context.Context, db *sql.DB) (bool, error) {
	var exists bool
	if err := db.QueryRowContext(ctx, "SELECT CONVERT(bit, CASE WHEN OBJECT_ID(N'dbo.mig_operation_journal', N'U') IS NULL THEN 0 ELSE 1 END);").Scan(&exists); err != nil {
		return false, fmt.Errorf("check migration operation journal: %w", err)
	}
	return exists, nil
}

func nullableVersion(state VersionState) any {
	if state.Version == nil {
		return nil
	}
	return int64(*state.Version)
}

func nullString(value string) any {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil
	}
	return value
}
