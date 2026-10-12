package migrest

import (
	"context"
	"database/sql"
	"fmt"
)

// NoopLockManager is used when DB lock is disabled.
type NoopLockManager struct{}

func (NoopLockManager) Acquire(context.Context) (func() error, error) {
	return func() error { return nil }, nil
}

// MSSQLSessionLockManager acquires lock through dbo.usp_mig_try_get_lock.
// Lock owner is session, so connection is held until release() is called.
type MSSQLSessionLockManager struct {
	DB        *sql.DB
	Resource  string
	TimeoutMS int
}

func (m *MSSQLSessionLockManager) Acquire(ctx context.Context) (func() error, error) {
	if m == nil || m.DB == nil {
		return nil, fmt.Errorf("lock manager db is nil")
	}
	conn, err := m.DB.Conn(ctx)
	if err != nil {
		return nil, err
	}
	_, err = conn.ExecContext(
		ctx,
		"EXEC dbo.usp_mig_try_get_lock @resource=@resource, @timeout_ms=@timeout_ms",
		sql.Named("resource", m.Resource),
		sql.Named("timeout_ms", m.TimeoutMS),
	)
	if err != nil {
		_ = conn.Close()
		return nil, err
	}
	return conn.Close, nil
}
