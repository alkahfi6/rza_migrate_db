package migrest

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlserver"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	_ "github.com/microsoft/go-mssqldb"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"

	"rza_migrate_db/internal/requestmeta"
)

const (
	coreAuditUpName   = "000000__core__mig__audit_init.up.sql"
	coreAuditDownName = "000000__core__mig__audit_init.down.sql"
)

// ServiceConfig controls migration API behavior.
type ServiceConfig struct {
	ServiceName             string
	DatabaseURL             string
	MigrationsDir           string
	AllowDown               bool
	AllowRunSpecificFile    bool
	AllowAcknowledgeManual  bool
	OperationTimeoutSeconds int
	DBLockEnabled           bool
	DBLockResource          string
	DBLockTimeoutMS         int
}

// Service orchestrates safe migration operations.
type Service struct {
	cfg            ServiceConfig
	tracerName     string
	migrator       MigratorFactory
	lock           LockManager
	journal        operationJournal
	closeResources func() error
	mu             sync.Mutex
}

// NewServiceConfigFromRuntime converts runtime YAML config to service config.
func NewServiceConfigFromRuntime(cfg Config) ServiceConfig {
	return ServiceConfig{
		ServiceName:             cfg.ServiceName,
		DatabaseURL:             cfg.Migration.DatabaseURL,
		MigrationsDir:           cfg.Migration.MigrationsDir,
		AllowDown:               cfg.Migration.AllowDown,
		AllowRunSpecificFile:    cfg.Migration.AllowRunSpecificFile,
		AllowAcknowledgeManual:  cfg.Migration.AllowAcknowledgeManual,
		OperationTimeoutSeconds: cfg.Migration.OperationTimeoutSeconds,
		DBLockEnabled:           cfg.Migration.DBLock.Enabled,
		DBLockResource:          cfg.Migration.DBLock.Resource,
		DBLockTimeoutMS:         cfg.Migration.DBLock.TimeoutMS,
	}
}

// Normalize applies defaults for empty service fields.
func (c *ServiceConfig) Normalize() {
	if c == nil {
		return
	}
	c.DatabaseURL = strings.TrimSpace(c.DatabaseURL)
	c.ServiceName = strings.TrimSpace(c.ServiceName)
	if c.ServiceName == "" {
		c.ServiceName = defaultServiceName
	}
	c.MigrationsDir = strings.TrimSpace(c.MigrationsDir)
	if c.OperationTimeoutSeconds <= 0 {
		c.OperationTimeoutSeconds = 300
	}
	if c.DBLockResource == "" {
		c.DBLockResource = "rza_migrate_db"
	}
	if c.DBLockTimeoutMS <= 0 {
		c.DBLockTimeoutMS = 60000
	}
}

// Validate enforces non-empty critical fields.
func (c ServiceConfig) Validate() error {
	var issues []string
	if strings.TrimSpace(c.DatabaseURL) == "" {
		issues = append(issues, "database url is required")
	}
	if strings.TrimSpace(c.MigrationsDir) == "" {
		issues = append(issues, "migrations dir is required")
	}
	if c.OperationTimeoutSeconds <= 0 {
		issues = append(issues, "operation timeout seconds must be > 0")
	}
	if c.DBLockEnabled {
		if strings.TrimSpace(c.DBLockResource) == "" {
			issues = append(issues, "db lock resource is required when lock is enabled")
		}
		if c.DBLockTimeoutMS <= 0 {
			issues = append(issues, "db lock timeout ms must be > 0 when lock is enabled")
		}
	}
	if len(issues) > 0 {
		return errors.New(strings.Join(issues, "; "))
	}
	return nil
}

// NewService builds production service with golang-migrate and optional DB lock.
func NewService(cfg ServiceConfig) (*Service, error) {
	cfg.Normalize()
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	if err := ensureMigrationsDir(cfg.MigrationsDir); err != nil {
		return nil, err
	}

	factory, err := NewGoMigrateFactory(cfg.MigrationsDir, cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlserver", cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("open sqlserver connection for migration control: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeoutDuration(cfg.OperationTimeoutSeconds))
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping sqlserver for migration control: %w", err)
	}

	lock := LockManager(NoopLockManager{})
	if cfg.DBLockEnabled {
		lock = &MSSQLSessionLockManager{
			DB:        db,
			Resource:  cfg.DBLockResource,
			TimeoutMS: cfg.DBLockTimeoutMS,
		}
	}
	service, err := NewServiceWithDependencies(cfg, factory, lock, db.Close)
	if err != nil {
		_ = db.Close()
		return nil, err
	}
	service.journal = mssqlOperationJournal{db: db}
	return service, nil
}

// NewServiceWithDependencies is used by tests/custom runtime wiring.
func NewServiceWithDependencies(cfg ServiceConfig, factory MigratorFactory, lock LockManager, closeFn func() error) (*Service, error) {
	cfg.Normalize()
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	if factory == nil {
		return nil, errors.New("migrator factory is required")
	}
	if lock == nil {
		lock = NoopLockManager{}
	}
	if closeFn == nil {
		closeFn = func() error { return nil }
	}
	return &Service{
		cfg:            cfg,
		tracerName:     buildServiceTracerName(cfg.ServiceName),
		migrator:       factory,
		lock:           lock,
		journal:        noopOperationJournal{},
		closeResources: closeFn,
	}, nil
}

func buildServiceTracerName(serviceName string) string {
	trimmed := strings.TrimSpace(serviceName)
	if trimmed == "" {
		trimmed = defaultServiceName
	}
	return trimmed + "/service"
}

// NewGoMigrateFactory creates migrator instances from file source + sqlserver URL.
func NewGoMigrateFactory(migrationsDir, databaseURL string) (MigratorFactory, error) {
	abs, err := filepath.Abs(migrationsDir)
	if err != nil {
		return nil, fmt.Errorf("resolve migrations dir: %w", err)
	}
	sourceURL := toFileSourceURL(abs)
	return func() (Migrator, error) {
		m, err := migrate.New(sourceURL, databaseURL)
		if err != nil {
			return nil, err
		}
		return m, nil
	}, nil
}

// Close releases service runtime resources.
func (s *Service) Close() error {
	if s == nil || s.closeResources == nil {
		return nil
	}
	return s.closeResources()
}

// CurrentVersion returns current migration version state.
func (s *Service) CurrentVersion(ctx context.Context) (VersionState, error) {
	if s == nil {
		return VersionState{}, errors.New("service is nil")
	}
	log := operationLogger(ctx, "current_version", "")
	ctx, span := otel.Tracer(s.tracerName).Start(ctx, "service.current_version")
	defer span.End()
	_, cancel := s.operationContext(ctx)
	defer cancel()

	s.mu.Lock()
	defer s.mu.Unlock()

	m, err := s.migrator()
	if err != nil {
		span.RecordError(err)
		log.Error("create_migrator_failed", zap.Error(err))
		return VersionState{}, newServiceError(ErrorKindInternal, "create migrator failed", err)
	}
	defer func() {
		srcErr, dbErr := m.Close()
		_ = srcErr
		_ = dbErr
	}()

	state, err := readVersionState(m)
	if err != nil {
		span.RecordError(err)
		log.Error("read_current_version_failed", zap.Error(err))
		return VersionState{}, newServiceError(ErrorKindInternal, "read current version failed", err)
	}
	if state.Dirty {
		span.AddEvent("state_dirty")
		log.Warn("migration_state_dirty")
		return state, newServiceError(ErrorKindConflict, "migration state is dirty", nil)
	}
	log.Debug("current_version_ok", appendVersionStateFields(nil, "current", state)...)
	return state, nil
}

// Pending returns current state plus pending up migration files in deterministic order.
func (s *Service) Pending(ctx context.Context) (PendingState, error) {
	if s == nil {
		return PendingState{}, errors.New("service is nil")
	}
	log := operationLogger(ctx, "pending", "")
	ctx, span := otel.Tracer(s.tracerName).Start(ctx, "service.pending")
	defer span.End()
	_, cancel := s.operationContext(ctx)
	defer cancel()

	s.mu.Lock()
	defer s.mu.Unlock()

	m, err := s.migrator()
	if err != nil {
		span.RecordError(err)
		log.Error("create_migrator_failed", zap.Error(err))
		return PendingState{}, newServiceError(ErrorKindInternal, "create migrator failed", err)
	}
	defer func() {
		srcErr, dbErr := m.Close()
		_ = srcErr
		_ = dbErr
	}()

	state, err := readVersionState(m)
	if err != nil {
		span.RecordError(err)
		log.Error("read_current_version_failed", zap.Error(err))
		return PendingState{}, newServiceError(ErrorKindInternal, "read current version failed", err)
	}

	files, err := s.listPendingMigrations(state.Version)
	if err != nil {
		span.RecordError(err)
		log.Error("list_pending_migrations_failed", zap.Error(err))
		return PendingState{}, err
	}
	log.Debug("pending_snapshot_ok",
		append(
			appendVersionStateFields(nil, "current", state),
			zap.Int("pending_count", len(files)),
		)...,
	)
	return PendingState{
		Current: state,
		Files:   files,
	}, nil
}

// Up runs full up (steps<=0) or up N steps.
func (s *Service) Up(ctx context.Context, steps int) (OperationResult, error) {
	operationLog := operationLogger(ctx, "up", "")
	operationLog.Info("up_requested", zap.Int("steps", steps))
	ctx, span := otel.Tracer(s.tracerName).Start(ctx, "service.up")
	span.SetAttributes(attribute.Int("migration.steps", steps))
	defer span.End()
	return s.runWithMigrationLockWithOptions(ctx, "up", "", runOptions{RequestedSteps: steps}, func(m Migrator, _ VersionState) error {
		if steps <= 0 {
			return m.Up()
		}
		return m.Steps(steps)
	})
}

// Down runs down N steps. Full down is intentionally disallowed.
func (s *Service) Down(ctx context.Context, steps int) (OperationResult, error) {
	operationLog := operationLogger(ctx, "down", "")
	operationLog.Info("down_requested", zap.Int("steps", steps))
	ctx, span := otel.Tracer(s.tracerName).Start(ctx, "service.down")
	span.SetAttributes(attribute.Int("migration.steps", steps))
	defer span.End()
	if !s.cfg.AllowDown {
		operationLog.Warn("down_blocked_by_config")
		return OperationResult{}, newServiceError(ErrorKindForbidden, "down migration is disabled by config", nil)
	}
	if steps <= 0 {
		operationLog.Warn("down_invalid_steps", zap.Int("steps", steps))
		return OperationResult{}, newServiceError(ErrorKindInvalidInput, "down steps must be > 0", nil)
	}
	return s.runWithMigrationLockWithOptions(ctx, "down", "", runOptions{RequestedSteps: steps}, func(m Migrator, _ VersionState) error {
		return m.Steps(-steps)
	})
}

// RunSpecificFile executes exactly one migration step bound to requested file.
// Safety rule:
// - *.up.sql: requested file must be next version from current state.
// - *.down.sql: requested file must match current version.
func (s *Service) RunSpecificFile(ctx context.Context, fileName string) (OperationResult, error) {
	operationLog := operationLogger(ctx, "run_file", fileName)
	operationLog.Info("run_specific_file_requested")
	ctx, span := otel.Tracer(s.tracerName).Start(ctx, "service.run_file")
	span.SetAttributes(attribute.String("migration.file", fileName))
	defer span.End()
	if !s.cfg.AllowRunSpecificFile {
		operationLog.Warn("run_specific_file_blocked_by_config")
		return OperationResult{}, newServiceError(ErrorKindForbidden, "run specific migration file is disabled by config", nil)
	}
	target, err := s.resolveRequestedFile(fileName)
	if err != nil {
		return OperationResult{}, err
	}

	return s.runWithMigrationLockWithOptions(ctx, "run_file", target.FileName, runOptions{RequestedSteps: 1}, func(m Migrator, before VersionState) error {
		upVersions, err := s.listUpVersions()
		if err != nil {
			return err
		}

		switch target.Direction {
		case "up":
			next, ok := nextUpVersion(upVersions, before.Version)
			if !ok {
				return newServiceError(ErrorKindConflict, "no next migration version is available", nil)
			}
			if next != target.Version {
				return newServiceError(
					ErrorKindConflict,
					fmt.Sprintf("requested up file %s is not next version (expected %d)", target.FileName, next),
					nil,
				)
			}
			return m.Steps(1)
		case "down":
			if before.Version == nil {
				return newServiceError(
					ErrorKindConflict,
					fmt.Sprintf("requested down file %s cannot run because current version is nil", target.FileName),
					nil,
				)
			}
			if *before.Version != target.Version {
				return newServiceError(
					ErrorKindConflict,
					fmt.Sprintf("requested down file %s must match current version %d", target.FileName, *before.Version),
					nil,
				)
			}
			return m.Steps(-1)
		default:
			return newServiceError(ErrorKindInvalidInput, "invalid migration file direction", nil)
		}
	})
}

// AcknowledgeManualFile marks one UP migration as already handled manually.
// This is intended for controlled recovery when migration previously failed/dirty.
func (s *Service) AcknowledgeManualFile(ctx context.Context, fileName string, reason string) (OperationResult, error) {
	reason = strings.TrimSpace(reason)
	if !s.cfg.AllowAcknowledgeManual {
		return OperationResult{}, newServiceError(ErrorKindForbidden, "ack manual is disabled by config", nil)
	}
	if reason == "" {
		return OperationResult{}, newServiceError(ErrorKindInvalidInput, "reason is required", nil)
	}

	target, err := s.resolveRequestedFile(fileName)
	if err != nil {
		return OperationResult{}, err
	}
	if target.Direction != "up" {
		return OperationResult{}, newServiceError(ErrorKindInvalidInput, "ack manual only supports .up.sql migration file", nil)
	}

	operationLog := operationLogger(ctx, "ack_manual", target.FileName)
	operationLog.Warn("ack_manual_requested", zap.String("reason", reason))
	ctx, span := otel.Tracer(s.tracerName).Start(ctx, "service.ack_manual")
	span.SetAttributes(
		attribute.String("migration.file", target.FileName),
		attribute.String("migration.reason", reason),
	)
	defer span.End()

	return s.runWithMigrationLockWithOptions(
		ctx,
		"ack_manual",
		target.FileName,
		runOptions{AllowDirtyPreState: true, RequestedSteps: 1},
		func(m Migrator, before VersionState) error {
			upVersions, err := s.listUpVersions()
			if err != nil {
				return err
			}

			found := false
			for _, v := range upVersions {
				if v == target.Version {
					found = true
					break
				}
			}
			if !found {
				return newServiceError(
					ErrorKindInvalidInput,
					fmt.Sprintf("requested migration file %s is not registered in migrations directory", target.FileName),
					nil,
				)
			}

			if before.Dirty {
				if before.Version == nil {
					return newServiceError(
						ErrorKindConflict,
						"migration state is dirty but current version is nil",
						nil,
					)
				}
				if *before.Version != target.Version {
					return newServiceError(
						ErrorKindConflict,
						fmt.Sprintf("dirty version %d does not match requested file version %d", *before.Version, target.Version),
						nil,
					)
				}
			} else {
				next, ok := nextUpVersion(upVersions, before.Version)
				if !ok {
					return newServiceError(ErrorKindConflict, "no next migration version is available to acknowledge", nil)
				}
				if next != target.Version {
					return newServiceError(
						ErrorKindConflict,
						fmt.Sprintf("requested file %s is not next version (expected %d)", target.FileName, next),
						nil,
					)
				}
			}

			maxInt := uint64(^uint(0) >> 1)
			if target.Version > maxInt {
				return newServiceError(ErrorKindInvalidInput, "target version exceeds platform int range", nil)
			}
			return m.Force(int(target.Version))
		},
	)
}

type runOptions struct {
	AllowDirtyPreState bool
	RequestedSteps     int
}

func (s *Service) runWithMigrationLockWithOptions(ctx context.Context, operation string, fileName string, opts runOptions, run func(m Migrator, before VersionState) error) (res OperationResult, err error) {
	if s == nil {
		return OperationResult{}, errors.New("service is nil")
	}
	log := operationLogger(ctx, operation, fileName)
	ctx, span := otel.Tracer(s.tracerName).Start(ctx, "service.run_with_lock")
	span.SetAttributes(
		attribute.String("migration.operation", operation),
		attribute.String("migration.file", fileName),
	)
	defer span.End()
	opCtx, cancel := s.operationContext(ctx)
	defer cancel()
	startedAt := time.Now()
	operationID := uuid.NewString()

	res = OperationResult{
		Operation:      operation,
		OperationID:    operationID,
		FileName:       fileName,
		RequestedSteps: opts.RequestedSteps,
	}
	log.Info("migration_operation_started",
		zap.Int("timeout_seconds", s.cfg.OperationTimeoutSeconds),
	)

	s.mu.Lock()
	defer s.mu.Unlock()

	release, err := s.lock.Acquire(opCtx)
	if err != nil {
		span.RecordError(err)
		log.Error("acquire_migration_lock_failed", zap.Error(err))
		if errors.Is(err, context.DeadlineExceeded) {
			return res, newServiceError(
				ErrorKindConflict,
				fmt.Sprintf("failed to acquire migration lock within operation_timeout_seconds=%d", s.cfg.OperationTimeoutSeconds),
				err,
			)
		}
		return res, newServiceError(ErrorKindConflict, "failed to acquire migration lock", err)
	}
	defer func() {
		if release == nil {
			return
		}
		releaseErr := release()
		if releaseErr != nil && err == nil {
			log.Error("release_migration_lock_failed", zap.Error(releaseErr))
			err = newServiceError(ErrorKindInternal, "release migration lock failed", releaseErr)
		}
	}()

	m, err := s.migrator()
	if err != nil {
		span.RecordError(err)
		log.Error("create_migrator_failed", zap.Error(err))
		return res, newServiceError(ErrorKindInternal, "create migrator failed", err)
	}
	defer func() {
		srcErr, dbErr := m.Close()
		if err == nil && (srcErr != nil || dbErr != nil) {
			log.Error("close_migrator_failed", zap.Error(joinErrors(srcErr, dbErr)))
			err = newServiceError(ErrorKindInternal, "close migrator failed", joinErrors(srcErr, dbErr))
		}
	}()

	before, err := readVersionState(m)
	if err != nil {
		span.RecordError(err)
		log.Error("read_pre_migration_version_failed", zap.Error(err))
		return res, newServiceError(ErrorKindInternal, "read pre-migration version failed", err)
	}
	res.Before = before
	if before.Dirty && !opts.AllowDirtyPreState {
		span.AddEvent("state_dirty_pre_migration")
		log.Warn("pre_migration_state_dirty")
		return res, newServiceError(ErrorKindConflict, "migration state is dirty; manual recovery is required", nil)
	}
	log.Debug("migration_pre_state", appendVersionStateFields(nil, "before", before)...)
	if err := s.journal.ValidateAppliedHistory(opCtx, s.cfg.MigrationsDir); err != nil {
		span.RecordError(err)
		log.Error("validate_applied_history_failed", zap.Error(err))
		return res, newServiceError(ErrorKindConflict, "immutable migration history validation failed", err)
	}
	direction := "none"
	if operation == "up" || operation == "ack_manual" {
		direction = "up"
	} else if operation == "down" {
		direction = "down"
	}
	requestID := strings.TrimSpace(requestmeta.RequestID(ctx))
	if requestID == "" {
		requestID = operationID
	}
	journalRecord := operationJournalRecord{
		OperationID:    operationID,
		RequestID:      requestID,
		Operation:      operation,
		FileName:       fileName,
		RequestedSteps: opts.RequestedSteps,
		Before:         before,
		Status:         "running",
		Direction:      direction,
		NIK:            requestmeta.NIK(ctx),
		TraceID:        requestmeta.TraceID(ctx),
		StartedAt:      startedAt.UTC(),
	}
	if err := s.journal.Start(opCtx, journalRecord); err != nil {
		span.RecordError(err)
		log.Error("start_operation_journal_failed", zap.Error(err))
		return res, newServiceError(ErrorKindInternal, "mandatory migration operation journal failed", err)
	}
	finishJournal := func(status string, after VersionState, operationErr error) error {
		files, manifest, resolvedDirection, reportErr := buildAppliedMigrationFiles(s.cfg.MigrationsDir, before, after)
		if reportErr != nil {
			return reportErr
		}
		if operation == "ack_manual" {
			for index := range files {
				files[index].Direction = "ack"
			}
			manifest = buildMigrationManifest(files)
			resolvedDirection = "ack"
		}
		if resolvedDirection != "none" {
			journalRecord.Direction = resolvedDirection
		}
		journalRecord.Status = status
		journalRecord.After = after
		journalRecord.AppliedFiles = files
		journalRecord.ManifestSHA256 = manifest
		if operationErr != nil {
			journalRecord.ErrorMessage = operationErr.Error()
		}
		finishedAt := time.Now().UTC()
		journalRecord.FinishedAt = &finishedAt
		res.AppliedFiles = files
		res.ManifestSHA256 = manifest
		journalCtx, journalCancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer journalCancel()
		return s.journal.Finish(journalCtx, journalRecord)
	}

	runErr := run(m, before)
	if runErr != nil {
		if errors.Is(runErr, migrate.ErrNoChange) {
			span.AddEvent("no_change")
			res.NoChange = true
			res.After = before
			if journalErr := finishJournal("no_change", before, nil); journalErr != nil {
				return res, newServiceError(ErrorKindInternal, "finish migration operation journal failed", journalErr)
			}
			log.Info("migration_operation_no_change",
				append(appendVersionStateFields(nil, "before", before),
					appendVersionStateFields(nil, "after", before)...,
				)...,
			)
			return res, nil
		}
		span.RecordError(runErr)
		log.Error("migration_execution_failed", zap.Error(runErr))
		failedAfter := before
		if state, stateErr := readVersionState(m); stateErr == nil {
			failedAfter = state
			res.After = state
		}
		if journalErr := finishJournal("failed", failedAfter, runErr); journalErr != nil {
			runErr = errors.Join(runErr, journalErr)
			log.Error("finish_failed_operation_journal_failed", zap.Error(journalErr))
		}
		var dirtyErr migrate.ErrDirty
		if errors.As(runErr, &dirtyErr) {
			return res, newServiceError(
				ErrorKindConflict,
				fmt.Sprintf("migration state is dirty at version %d", dirtyErr.Version),
				runErr,
			)
		}
		var svcErr *ServiceError
		if errors.As(runErr, &svcErr) {
			return res, runErr
		}
		return res, newServiceError(ErrorKindInternal, "migration execution failed", runErr)
	}

	after, err := readVersionState(m)
	if err != nil {
		span.RecordError(err)
		log.Error("read_post_migration_version_failed", zap.Error(err))
		return res, newServiceError(ErrorKindInternal, "read post-migration version failed", err)
	}
	res.After = after
	if after.Dirty {
		span.AddEvent("state_dirty_post_migration")
		log.Warn("post_migration_state_dirty")
		dirtyErr := errors.New("migration completed but DB state is dirty")
		if journalErr := finishJournal("failed", after, dirtyErr); journalErr != nil {
			log.Error("finish_dirty_operation_journal_failed", zap.Error(journalErr))
		}
		return res, newServiceError(ErrorKindConflict, "migration completed but DB state is dirty", nil)
	}
	if journalErr := finishJournal("success", after, nil); journalErr != nil {
		span.RecordError(journalErr)
		log.Error("finish_operation_journal_failed", zap.Error(journalErr))
		return res, newServiceError(ErrorKindInternal, "migration completed but mandatory operation journal failed", journalErr)
	}
	successFields := appendVersionStateFields(nil, "before", before)
	successFields = appendVersionStateFields(successFields, "after", after)
	successFields = append(successFields, zap.Duration("duration", time.Since(startedAt)))
	log.Info("migration_operation_completed", successFields...)
	return res, nil
}

func (s *Service) operationContext(parent context.Context) (context.Context, context.CancelFunc) {
	if parent == nil {
		parent = context.Background()
	}
	timeout := operationTimeoutDuration(s.cfg.OperationTimeoutSeconds)
	return context.WithTimeout(parent, timeout)
}

func operationTimeoutDuration(seconds int) time.Duration {
	if seconds <= 0 {
		return 300 * time.Second
	}
	return time.Duration(seconds) * time.Second
}

func readVersionState(m Migrator) (VersionState, error) {
	version, dirty, err := m.Version()
	if err != nil {
		if errors.Is(err, migrate.ErrNilVersion) {
			return VersionState{
				Version: nil,
				Dirty:   false,
			}, nil
		}
		return VersionState{}, err
	}
	v := uint64(version)
	return VersionState{
		Version: &v,
		Dirty:   dirty,
	}, nil
}

type requestedFile struct {
	FileName  string
	Direction string
	Version   uint64
}

func (s *Service) resolveRequestedFile(raw string) (requestedFile, error) {
	fileName := strings.TrimSpace(raw)
	if fileName == "" {
		return requestedFile{}, newServiceError(ErrorKindInvalidInput, "file name is required", nil)
	}
	base := filepath.Base(fileName)
	if base != fileName {
		return requestedFile{}, newServiceError(ErrorKindInvalidInput, "file name must not contain path separator", nil)
	}
	if strings.Contains(base, "..") {
		return requestedFile{}, newServiceError(ErrorKindInvalidInput, "file name must not contain parent traversal", nil)
	}

	fullPath := filepath.Join(s.cfg.MigrationsDir, base)
	info, err := os.Stat(fullPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return requestedFile{}, newServiceError(ErrorKindInvalidInput, "requested migration file does not exist", err)
		}
		return requestedFile{}, newServiceError(ErrorKindInternal, "check migration file failed", err)
	}
	if info.IsDir() {
		return requestedFile{}, newServiceError(ErrorKindInvalidInput, "requested migration file must not be a directory", nil)
	}

	meta, issue := parseMigrationFileMeta(base)
	if issue != nil {
		return requestedFile{}, newServiceError(ErrorKindInvalidInput, issue.Message, nil)
	}
	version, err := strconv.ParseUint(meta.Version, 10, 64)
	if err != nil {
		return requestedFile{}, newServiceError(ErrorKindInvalidInput, "migration version is not numeric", err)
	}
	return requestedFile{
		FileName:  base,
		Direction: meta.Direction,
		Version:   version,
	}, nil
}

func (s *Service) listUpVersions() ([]uint64, error) {
	entries, err := os.ReadDir(s.cfg.MigrationsDir)
	if err != nil {
		return nil, newServiceError(ErrorKindInternal, "read migrations directory failed", err)
	}
	versionSet := map[uint64]struct{}{}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if !strings.HasSuffix(strings.ToLower(name), ".up.sql") {
			continue
		}
		meta, issue := parseMigrationFileMeta(name)
		if issue != nil {
			return nil, newServiceError(
				ErrorKindConflict,
				fmt.Sprintf("invalid migration filename detected: %s (%s)", name, issue.Message),
				nil,
			)
		}
		version, err := strconv.ParseUint(meta.Version, 10, 64)
		if err != nil {
			return nil, newServiceError(
				ErrorKindConflict,
				fmt.Sprintf("invalid numeric version on migration filename: %s", name),
				err,
			)
		}
		versionSet[version] = struct{}{}
	}

	versions := make([]uint64, 0, len(versionSet))
	for v := range versionSet {
		versions = append(versions, v)
	}
	sort.Slice(versions, func(i, j int) bool {
		return versions[i] < versions[j]
	})
	return versions, nil
}

func (s *Service) listPendingMigrations(current *uint64) ([]PendingMigration, error) {
	entries, err := os.ReadDir(s.cfg.MigrationsDir)
	if err != nil {
		return nil, newServiceError(ErrorKindInternal, "read migrations directory failed", err)
	}

	out := make([]PendingMigration, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if !strings.HasSuffix(strings.ToLower(name), ".up.sql") {
			continue
		}

		meta, issue := parseMigrationFileMeta(name)
		if issue != nil {
			return nil, newServiceError(
				ErrorKindConflict,
				fmt.Sprintf("invalid migration filename detected: %s (%s)", name, issue.Message),
				nil,
			)
		}
		version, parseErr := strconv.ParseUint(meta.Version, 10, 64)
		if parseErr != nil {
			return nil, newServiceError(
				ErrorKindConflict,
				fmt.Sprintf("invalid numeric version on migration filename: %s", name),
				parseErr,
			)
		}

		if current != nil && version <= *current {
			continue
		}

		out = append(out, PendingMigration{
			FileName: meta.FileName,
			Version:  version,
			Category: meta.Category,
			Object:   meta.Object,
			VK:       meta.VK,
			IsCore:   meta.IsCore,
		})
	}

	sort.Slice(out, func(i, j int) bool {
		if out[i].Version != out[j].Version {
			return out[i].Version < out[j].Version
		}
		return out[i].FileName < out[j].FileName
	})
	return out, nil
}

func ensureMigrationsDir(dir string) error {
	info, err := os.Stat(dir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("migrations directory does not exist: %s", filepath.ToSlash(dir))
		}
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("migrations path is not a directory: %s", filepath.ToSlash(dir))
	}
	return nil
}

func nextUpVersion(versions []uint64, current *uint64) (uint64, bool) {
	if len(versions) == 0 {
		return 0, false
	}
	if current == nil {
		return versions[0], true
	}
	for _, v := range versions {
		if v > *current {
			return v, true
		}
	}
	return 0, false
}

func joinErrors(left, right error) error {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return fmt.Errorf("%v; %v", left, right)
}

func toFileSourceURL(absPath string) string {
	path := filepath.ToSlash(absPath)
	if runtime.GOOS == "windows" && len(path) >= 2 && path[1] == ':' {
		// golang-migrate/source/file on Windows expects file://C:/...
		// (file:///C:/... becomes /C:/... and fails in os.DirFS)
		return "file://" + path
	}
	if strings.HasPrefix(path, "/") {
		return "file://" + path
	}
	return "file:///" + path
}
