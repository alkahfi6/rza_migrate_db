# Legacy Wrapper Cutover Policy (Completed)

Dokumen ini mengunci transisi dari mode legacy wrapper audit ke strict-only.

## Cutover Status

- Effective 2026-02-19: pipeline CI wajib berjalan dengan `RZA_LEGACY_AUDIT_MODE=fail`.
- Effective 2026-02-19: mode legacy `warn` tidak lagi didukung oleh CLI/engine.
- Cleanup state: transisi selesai, strict-only menjadi baseline permanen.

## Enforcement Rules

1. CI wajib fail jika `RZA_LEGACY_AUDIT_MODE` tidak bernilai `fail`.
2. Repository migration baru wajib mengikuti wrapper generated-style (named audit args + deterministic metadata).
3. Mode `warn` dilarang untuk semua alur (lokal maupun CI), bukan hanya quality gate.

## Migration Job Lock Contract

Migration job aplikasi wajib acquire lock sebelum `migrate up`:

- Template acuan: `.gitlab/migration_jobs/migrate_with_lock.yml`
- Wajib memanggil `dbo.usp_mig_try_get_lock` sebelum `migrate -path ./migrations -database "$MSSQL_DSN" up`
