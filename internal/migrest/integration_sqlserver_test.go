package migrest

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestIntegration_MSSQL_CurrentVersion(t *testing.T) {
	t.Parallel()

	dsn := os.Getenv("RZA_MSSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("skip integration test: RZA_MSSQL_TEST_DSN is empty")
	}

	migrationsDir := t.TempDir()
	writeSQLFixture(t, filepath.Join(migrationsDir, "000000__core__mig__audit_init.up.sql"), "SELECT 1;\n")
	writeSQLFixture(t, filepath.Join(migrationsDir, "000000__core__mig__audit_init.down.sql"), "SELECT 1;\n")

	svc, err := NewService(ServiceConfig{
		DatabaseURL:             dsn,
		MigrationsDir:           migrationsDir,
		AllowDown:               false,
		AllowRunSpecificFile:    false,
		OperationTimeoutSeconds: 30,
		DBLockEnabled:           false,
	})
	if err != nil {
		t.Fatalf("new service failed: %v", err)
	}
	defer func() { _ = svc.Close() }()

	state, err := svc.CurrentVersion(context.Background())
	if err != nil {
		t.Fatalf("current version failed: %v", err)
	}
	if state.Dirty {
		t.Fatalf("expected non-dirty state for integration baseline")
	}
}

func TestIntegration_MSSQL_DBLockAcquireRelease(t *testing.T) {
	t.Parallel()

	dsn := os.Getenv("RZA_MSSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("skip integration test: RZA_MSSQL_TEST_DSN is empty")
	}

	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		t.Fatalf("open db failed: %v", err)
	}
	defer func() { _ = db.Close() }()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		t.Fatalf("ping db failed: %v", err)
	}

	var lockProcID sql.NullInt64
	if err := db.QueryRowContext(ctx, "SELECT OBJECT_ID(N'dbo.usp_mig_try_get_lock', N'P')").Scan(&lockProcID); err != nil {
		t.Fatalf("query lock proc existence failed: %v", err)
	}
	if !lockProcID.Valid {
		t.Skip("skip integration test: dbo.usp_mig_try_get_lock is not installed in target DB")
	}

	manager := &MSSQLSessionLockManager{
		DB:        db,
		Resource:  "rza_migrate_db_integration_test",
		TimeoutMS: 5000,
	}

	release, err := manager.Acquire(ctx)
	if err != nil {
		t.Fatalf("acquire db lock failed: %v", err)
	}
	if release == nil {
		t.Fatalf("expected non-nil release callback")
	}
	if err := release(); err != nil {
		t.Fatalf("release db lock failed: %v", err)
	}
}

func TestIntegration_MSSQL_DBLockConflictTimeout(t *testing.T) {
	t.Parallel()

	dsn := os.Getenv("RZA_MSSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("skip integration test: RZA_MSSQL_TEST_DSN is empty")
	}

	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		t.Fatalf("open db failed: %v", err)
	}
	defer func() { _ = db.Close() }()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		t.Fatalf("ping db failed: %v", err)
	}

	var lockProcID sql.NullInt64
	if err := db.QueryRowContext(ctx, "SELECT OBJECT_ID(N'dbo.usp_mig_try_get_lock', N'P')").Scan(&lockProcID); err != nil {
		t.Fatalf("query lock proc existence failed: %v", err)
	}
	if !lockProcID.Valid {
		t.Skip("skip integration test: dbo.usp_mig_try_get_lock is not installed in target DB")
	}

	resource := "rza_migrate_db_integration_conflict_test"
	holder := &MSSQLSessionLockManager{
		DB:        db,
		Resource:  resource,
		TimeoutMS: 5000,
	}
	releaseHolder, err := holder.Acquire(ctx)
	if err != nil {
		t.Fatalf("initial lock acquire failed: %v", err)
	}
	defer func() { _ = releaseHolder() }()

	contender := &MSSQLSessionLockManager{
		DB:        db,
		Resource:  resource,
		TimeoutMS: 200,
	}
	_, err = contender.Acquire(ctx)
	if err == nil {
		t.Fatalf("expected lock conflict timeout error, got nil")
	}
}

func writeSQLFixture(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write fixture %s failed: %v", path, err)
	}
}
