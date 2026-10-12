package migrest

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/golang-migrate/migrate/v4"
)

func TestRunSpecificFileUpRequiresNextVersion(t *testing.T) {
	t.Parallel()

	dir := setupMigrationDir(t)
	current := uint64(20260217073015000)
	fake := &fakeMigrator{
		version: &current,
		chain: []uint64{
			0,
			20260217073015000,
			20260217073015001,
		},
	}

	svc, err := NewServiceWithDependencies(
		ServiceConfig{
			DatabaseURL:             "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable",
			MigrationsDir:           dir,
			AllowDown:               true,
			AllowRunSpecificFile:    true,
			OperationTimeoutSeconds: 30,
			DBLockEnabled:           false,
		},
		func() (Migrator, error) { return fake, nil },
		NoopLockManager{},
		nil,
	)
	if err != nil {
		t.Fatalf("new service failed: %v", err)
	}

	_, err = svc.RunSpecificFile(context.Background(), "20260217073015002_views_dbo.ViewA_v1.up.sql")
	if err == nil {
		t.Fatalf("expected conflict error")
	}
	var svcErr *ServiceError
	if !errors.As(err, &svcErr) {
		t.Fatalf("expected ServiceError, got %T", err)
	}
	if svcErr.Kind != ErrorKindConflict {
		t.Fatalf("expected conflict kind, got %s", svcErr.Kind)
	}
	if fake.stepCalls != 0 {
		t.Fatalf("unexpected steps call count: got=%d want=0", fake.stepCalls)
	}
}

func TestRunSpecificFileDownRequiresCurrentVersion(t *testing.T) {
	t.Parallel()

	dir := setupMigrationDir(t)
	current := uint64(20260217073015001)
	fake := &fakeMigrator{
		version: &current,
		chain: []uint64{
			0,
			20260217073015000,
			20260217073015001,
		},
	}

	svc, err := NewServiceWithDependencies(
		ServiceConfig{
			DatabaseURL:             "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable",
			MigrationsDir:           dir,
			AllowDown:               true,
			AllowRunSpecificFile:    true,
			OperationTimeoutSeconds: 30,
			DBLockEnabled:           false,
		},
		func() (Migrator, error) { return fake, nil },
		NoopLockManager{},
		nil,
	)
	if err != nil {
		t.Fatalf("new service failed: %v", err)
	}

	_, err = svc.RunSpecificFile(context.Background(), "20260217073015000_tables_dbo.TableA_v1.down.sql")
	if err == nil {
		t.Fatalf("expected conflict error")
	}
	var svcErr *ServiceError
	if !errors.As(err, &svcErr) {
		t.Fatalf("expected ServiceError, got %T", err)
	}
	if svcErr.Kind != ErrorKindConflict {
		t.Fatalf("expected conflict kind, got %s", svcErr.Kind)
	}
	if fake.stepCalls != 0 {
		t.Fatalf("unexpected steps call count: got=%d want=0", fake.stepCalls)
	}
}

func TestRunSpecificFileExecutesSingleStepWhenValid(t *testing.T) {
	t.Parallel()

	dir := setupMigrationDir(t)
	current := uint64(20260217073015000)
	fake := &fakeMigrator{
		version: &current,
		chain: []uint64{
			0,
			20260217073015000,
			20260217073015001,
		},
	}

	svc, err := NewServiceWithDependencies(
		ServiceConfig{
			DatabaseURL:             "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable",
			MigrationsDir:           dir,
			AllowDown:               true,
			AllowRunSpecificFile:    true,
			OperationTimeoutSeconds: 30,
			DBLockEnabled:           false,
		},
		func() (Migrator, error) { return fake, nil },
		NoopLockManager{},
		nil,
	)
	if err != nil {
		t.Fatalf("new service failed: %v", err)
	}

	res, err := svc.RunSpecificFile(context.Background(), "20260217073015001_indexes_dbo.TableA_v1.up.sql")
	if err != nil {
		t.Fatalf("run specific file failed: %v", err)
	}
	if res.NoChange {
		t.Fatalf("expected real migration step, got no-change")
	}
	if fake.stepCalls != 1 {
		t.Fatalf("step call mismatch: got=%d want=1", fake.stepCalls)
	}
	if fake.lastSteps != 1 {
		t.Fatalf("step direction mismatch: got=%d want=1", fake.lastSteps)
	}
	if res.After.Version == nil || *res.After.Version != 20260217073015001 {
		t.Fatalf("after version mismatch: %+v", res.After)
	}
}

func TestDownRejectsZeroOrNegativeSteps(t *testing.T) {
	t.Parallel()

	dir := setupMigrationDir(t)
	current := uint64(20260217073015001)
	fake := &fakeMigrator{
		version: &current,
		chain: []uint64{
			0,
			20260217073015000,
			20260217073015001,
		},
	}

	svc, err := NewServiceWithDependencies(
		ServiceConfig{
			DatabaseURL:             "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable",
			MigrationsDir:           dir,
			AllowDown:               true,
			AllowRunSpecificFile:    true,
			OperationTimeoutSeconds: 30,
			DBLockEnabled:           false,
		},
		func() (Migrator, error) { return fake, nil },
		NoopLockManager{},
		nil,
	)
	if err != nil {
		t.Fatalf("new service failed: %v", err)
	}

	_, err = svc.Down(context.Background(), 0)
	if err == nil {
		t.Fatalf("expected invalid input error for steps=0")
	}
	var svcErr *ServiceError
	if !errors.As(err, &svcErr) {
		t.Fatalf("expected ServiceError, got %T", err)
	}
	if svcErr.Kind != ErrorKindInvalidInput {
		t.Fatalf("expected invalid input kind, got %s", svcErr.Kind)
	}
	if fake.stepCalls != 0 {
		t.Fatalf("unexpected steps call count for steps=0: got=%d want=0", fake.stepCalls)
	}

	_, err = svc.Down(context.Background(), -1)
	if err == nil {
		t.Fatalf("expected invalid input error for steps<0")
	}
	if !errors.As(err, &svcErr) {
		t.Fatalf("expected ServiceError, got %T", err)
	}
	if svcErr.Kind != ErrorKindInvalidInput {
		t.Fatalf("expected invalid input kind, got %s", svcErr.Kind)
	}
	if fake.stepCalls != 0 {
		t.Fatalf("unexpected steps call count for steps<0: got=%d want=0", fake.stepCalls)
	}
}

func TestAcknowledgeManualFileRejectsWhenDisabled(t *testing.T) {
	t.Parallel()

	dir := setupMigrationDir(t)
	current := uint64(20260217073015000)
	fake := &fakeMigrator{
		version: &current,
		chain: []uint64{
			0,
			20260217073015000,
			20260217073015001,
		},
	}

	svc, err := NewServiceWithDependencies(
		ServiceConfig{
			DatabaseURL:             "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable",
			MigrationsDir:           dir,
			AllowDown:               true,
			AllowRunSpecificFile:    true,
			AllowAcknowledgeManual:  false,
			OperationTimeoutSeconds: 30,
			DBLockEnabled:           false,
		},
		func() (Migrator, error) { return fake, nil },
		NoopLockManager{},
		nil,
	)
	if err != nil {
		t.Fatalf("new service failed: %v", err)
	}

	_, err = svc.AcknowledgeManualFile(context.Background(), "20260217073015001_indexes_dbo.TableA_v1.up.sql", "manual done")
	if err == nil {
		t.Fatalf("expected forbidden error")
	}
	var svcErr *ServiceError
	if !errors.As(err, &svcErr) {
		t.Fatalf("expected ServiceError, got %T", err)
	}
	if svcErr.Kind != ErrorKindForbidden {
		t.Fatalf("expected forbidden kind, got %s", svcErr.Kind)
	}
}

func TestAcknowledgeManualFileClearsDirtyWhenVersionMatches(t *testing.T) {
	t.Parallel()

	dir := setupMigrationDir(t)
	current := uint64(20260217073015001)
	fake := &fakeMigrator{
		version: &current,
		dirty:   true,
		chain: []uint64{
			0,
			20260217073015000,
			20260217073015001,
		},
	}

	svc, err := NewServiceWithDependencies(
		ServiceConfig{
			DatabaseURL:             "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable",
			MigrationsDir:           dir,
			AllowDown:               true,
			AllowRunSpecificFile:    true,
			AllowAcknowledgeManual:  true,
			OperationTimeoutSeconds: 30,
			DBLockEnabled:           false,
		},
		func() (Migrator, error) { return fake, nil },
		NoopLockManager{},
		nil,
	)
	if err != nil {
		t.Fatalf("new service failed: %v", err)
	}

	res, err := svc.AcknowledgeManualFile(context.Background(), "20260217073015001_indexes_dbo.TableA_v1.up.sql", "manual recovery")
	if err != nil {
		t.Fatalf("ack manual failed: %v", err)
	}
	if res.After.Version == nil || *res.After.Version != 20260217073015001 {
		t.Fatalf("after version mismatch: %+v", res.After)
	}
	if res.After.Dirty {
		t.Fatalf("expected clean state after force")
	}
	if fake.forceCalls != 1 {
		t.Fatalf("force call mismatch: got=%d want=1", fake.forceCalls)
	}
	if fake.lastForceVersion != int(20260217073015001) {
		t.Fatalf("force version mismatch: got=%d", fake.lastForceVersion)
	}
}

func TestAcknowledgeManualFileRejectsDirtyVersionMismatch(t *testing.T) {
	t.Parallel()

	dir := setupMigrationDir(t)
	current := uint64(20260217073015000)
	fake := &fakeMigrator{
		version: &current,
		dirty:   true,
		chain: []uint64{
			0,
			20260217073015000,
			20260217073015001,
		},
	}

	svc, err := NewServiceWithDependencies(
		ServiceConfig{
			DatabaseURL:             "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable",
			MigrationsDir:           dir,
			AllowDown:               true,
			AllowRunSpecificFile:    true,
			AllowAcknowledgeManual:  true,
			OperationTimeoutSeconds: 30,
			DBLockEnabled:           false,
		},
		func() (Migrator, error) { return fake, nil },
		NoopLockManager{},
		nil,
	)
	if err != nil {
		t.Fatalf("new service failed: %v", err)
	}

	_, err = svc.AcknowledgeManualFile(context.Background(), "20260217073015001_indexes_dbo.TableA_v1.up.sql", "manual recovery")
	if err == nil {
		t.Fatalf("expected conflict error")
	}
	var svcErr *ServiceError
	if !errors.As(err, &svcErr) {
		t.Fatalf("expected ServiceError, got %T", err)
	}
	if svcErr.Kind != ErrorKindConflict {
		t.Fatalf("expected conflict kind, got %s", svcErr.Kind)
	}
	if fake.forceCalls != 0 {
		t.Fatalf("unexpected force calls: got=%d want=0", fake.forceCalls)
	}
}

func TestAcknowledgeManualFileRejectsCleanSkipBeyondNextVersion(t *testing.T) {
	t.Parallel()

	dir := setupMigrationDir(t)
	current := uint64(20260217073015000)
	fake := &fakeMigrator{
		version: &current,
		dirty:   false,
		chain: []uint64{
			0,
			20260217073015000,
			20260217073015001,
			20260217073015002,
		},
	}

	svc, err := NewServiceWithDependencies(
		ServiceConfig{
			DatabaseURL:             "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable",
			MigrationsDir:           dir,
			AllowDown:               true,
			AllowRunSpecificFile:    true,
			AllowAcknowledgeManual:  true,
			OperationTimeoutSeconds: 30,
			DBLockEnabled:           false,
		},
		func() (Migrator, error) { return fake, nil },
		NoopLockManager{},
		nil,
	)
	if err != nil {
		t.Fatalf("new service failed: %v", err)
	}

	_, err = svc.AcknowledgeManualFile(context.Background(), "20260217073015002_views_dbo.ViewA_v1.up.sql", "manual recovery")
	if err == nil {
		t.Fatalf("expected conflict error")
	}
	var svcErr *ServiceError
	if !errors.As(err, &svcErr) {
		t.Fatalf("expected ServiceError, got %T", err)
	}
	if svcErr.Kind != ErrorKindConflict {
		t.Fatalf("expected conflict kind, got %s", svcErr.Kind)
	}
	if fake.forceCalls != 0 {
		t.Fatalf("unexpected force calls: got=%d want=0", fake.forceCalls)
	}
}

func TestPendingReturnsSortedPendingFiles(t *testing.T) {
	t.Parallel()

	dir := setupMigrationDir(t)
	current := uint64(20260217073015000)
	fake := &fakeMigrator{
		version: &current,
		dirty:   false,
		chain: []uint64{
			0,
			20260217073015000,
			20260217073015001,
			20260217073015002,
		},
	}

	svc, err := NewServiceWithDependencies(
		ServiceConfig{
			DatabaseURL:             "sqlserver://sa:Password1@localhost:1433?database=master&encrypt=disable",
			MigrationsDir:           dir,
			AllowDown:               true,
			AllowRunSpecificFile:    true,
			AllowAcknowledgeManual:  true,
			OperationTimeoutSeconds: 30,
			DBLockEnabled:           false,
		},
		func() (Migrator, error) { return fake, nil },
		NoopLockManager{},
		nil,
	)
	if err != nil {
		t.Fatalf("new service failed: %v", err)
	}

	pending, err := svc.Pending(context.Background())
	if err != nil {
		t.Fatalf("pending failed: %v", err)
	}
	if len(pending.Files) != 2 {
		t.Fatalf("pending file count mismatch: got=%d want=2", len(pending.Files))
	}
	if pending.Files[0].Version != 20260217073015001 {
		t.Fatalf("first pending version mismatch: got=%d", pending.Files[0].Version)
	}
	if pending.Files[1].Version != 20260217073015002 {
		t.Fatalf("second pending version mismatch: got=%d", pending.Files[1].Version)
	}
}

func setupMigrationDir(t *testing.T) string {
	t.Helper()

	dir := t.TempDir()
	files := []string{
		"000000__core__mig__audit_init.up.sql",
		"000000__core__mig__audit_init.down.sql",
		"20260217073015000_tables_dbo.TableA_v1.up.sql",
		"20260217073015000_tables_dbo.TableA_v1.down.sql",
		"20260217073015001_indexes_dbo.TableA_v1.up.sql",
		"20260217073015001_indexes_dbo.TableA_v1.down.sql",
		"20260217073015002_views_dbo.ViewA_v1.up.sql",
		"20260217073015002_views_dbo.ViewA_v1.down.sql",
	}
	for _, name := range files {
		path := filepath.Join(dir, name)
		if err := os.WriteFile(path, []byte("SELECT 1;\n"), 0o644); err != nil {
			t.Fatalf("write fixture %s failed: %v", name, err)
		}
	}
	return dir
}

type fakeMigrator struct {
	version          *uint64
	dirty            bool
	chain            []uint64
	stepCalls        int
	lastSteps        int
	forceCalls       int
	lastForceVersion int
}

func (f *fakeMigrator) Up() error {
	return f.Steps(9999)
}

func (f *fakeMigrator) Down() error {
	return f.Steps(-9999)
}

func (f *fakeMigrator) Steps(n int) error {
	f.stepCalls++
	f.lastSteps = n
	if len(f.chain) == 0 {
		return migrate.ErrNoChange
	}
	if n > 0 {
		next, ok := nextUpVersion(f.chain, f.version)
		if !ok {
			return migrate.ErrNoChange
		}
		f.version = &next
		return nil
	}
	if n < 0 {
		if f.version == nil {
			return migrate.ErrNoChange
		}
		cur := *f.version
		idx := -1
		for i := range f.chain {
			if f.chain[i] == cur {
				idx = i
				break
			}
		}
		if idx <= 0 {
			f.version = nil
			return nil
		}
		prev := f.chain[idx-1]
		f.version = &prev
		return nil
	}
	return migrate.ErrNoChange
}

func (f *fakeMigrator) Force(version int) error {
	f.forceCalls++
	f.lastForceVersion = version
	v := uint64(version)
	f.version = &v
	f.dirty = false
	return nil
}

func (f *fakeMigrator) Version() (uint, bool, error) {
	if f.version == nil {
		return 0, false, migrate.ErrNilVersion
	}
	return uint(*f.version), f.dirty, nil
}

func (f *fakeMigrator) Close() (error, error) {
	return nil, nil
}
