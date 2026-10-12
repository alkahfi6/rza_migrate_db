# Production Readiness Playbook

Dokumen ini adalah checklist operasional sebelum engine dipakai untuk release rutin.

## Observability

- Jalankan `rza_migrate_db` dengan `-perf-notes true` pada benchmark/release candidate.
- Review `out/output.json.notes.perf_ms` untuk fase `fmt`, `verify`, `gen`, reconstruction, desired collect, patch collect, dan wrapper verify.
- `rza_migrate_rest` production harus memakai logging JSON, trace/request id middleware, dan API key aktif.
- Review `dbo.mig_operation_journal` untuk report operasi dan `dbo.mig_event` untuk keputusan per file. Query siap pakai tersedia di [migration monitoring runbook](../operations/migration_monitoring.md).
- Simpan `operation_id` dan `manifest_sha256` dari response REST sebagai bukti release.
- Jika telemetry diaktifkan, set `telemetry.sample_ratio` eksplisit dan validasi OTLP endpoint dari environment deployment.

## Release Packaging

Untuk authoring/MR tool release, jalankan `build_all.bat <license_owner> <license_months>`. Publish wajib atomik dan menghasilkan `tools/rza_migrate_db.exe`, `tools/rza_migrate_check.exe`, `tools/rza_migrate_report.exe`, ketiga binary Linux yang bersesuaian di `tools/linux-amd64/`, `tools/license.key`, individual checksum, build identity, source digest, dan complete release manifest. Jalankan gate Linux yang sama dengan GitLab setelah source final; binary/checksum yang dibangun dari source lama adalah stop condition. Untuk perubahan report, smoke-test binary Windows dan Linux pada source migration yang sama serta bandingkan enam dataset history (metadata waktu run dikecualikan).

Distribusi service/extractor yang lebih luas tetap memakai builder Windows berikut:

- Build Windows bundle via:
  - `powershell -ExecutionPolicy Bypass -File scripts/build_dist_windows.ps1`
- Untuk release candidate, jangan pakai `-SkipTests` kecuali test sudah dijalankan pada job CI terpisah.
- Paket wajib berisi:
  - `rza_migrate_db.exe`
  - `rza_get_object_db.exe`
  - `rza_migrate_rest.exe`
  - template repo dengan `migrations/`, config REST, Dockerfile, dan contoh GitLab CI REST.
- Setelah build, jalankan smoke command pada template repo:
  - `rza_migrate_db.exe verify -repo <template_repo>`
  - `rza_migrate_db.exe gen -repo <template_repo>`

## CI Split

Pipeline disarankan dipisah menjadi gate berikut:

| Gate | Command | Tujuan |
| --- | --- | --- |
| Docs structure | existing GitLab docs gates | Memastikan rule/docs/script wajib ada. |
| Fast unit | `go test ./internal/... ./cmd/... -count=1` | Feedback cepat untuk package internal dan command. |
| Fixture contract | `go test ./tests/fixtures -count=1` | Validasi shape golden fixture. |
| Engine fixtures | `go test ./tests/engine -count=1` | Validasi fixture QA-01..QA-26 dan contract tests tambahan. |
| Integration SQL Server | `go test ./internal/migrest -run Integration -count=1` | Opt-in dengan DB test environment. |
| Repeatability | `scripts/validate_parallel_repeatability.ps1` | Serial/parallel deterministic. |
| Benchmark | `scripts/benchmark_parallel_read.ps1` | Baseline performa gen/verify. |
| Final hardening | `scripts/run_final_hardening_gate.ps1` | Gate release/manual sebelum distribusi. |

Runner lokal untuk susunan gate release:

```powershell
powershell -ExecutionPolicy Bypass -File scripts/run_release_readiness_gate.ps1
```

## Benchmark Baseline

- Simpan hasil benchmark di artefak CI, bukan sebagai file source.
- Minimal ukur:
  - serial vs parallel read
  - worker count efektif
  - reconstruction cache hit/miss
  - desired collect cache hit/miss
  - jumlah migration history dan object desired
- Jika durasi naik signifikan, review `out/output.json.notes.perf_ms` sebelum mengubah algoritma.

## Migration Runner Operational Playbook

1. Pastikan `migrations/` sudah direview dan `rza_migrate_db verify` pass.
2. Pastikan core audit init dan `dbo.usp_mig_try_get_lock` tersedia di migration history.
3. Deploy `rza_migrate_rest` dengan:
   - `security.enable_api_key=true`
   - `migration.allow_down=false`
   - `migration.allow_run_specific_file=false`
   - `migration.db_lock.enabled=true` untuk environment shared.
4. Cek `GET /healthz`.
5. Cek `GET /api/v1/migrations/version`.
6. Jalankan `POST /api/v1/migrations/up` tanpa `steps` untuk full apply, atau dengan `steps` untuk controlled rollout.
7. Jika runner berhenti pada dirty version, jangan aktifkan acknowledge manual sebelum root cause SQL diperbaiki dan hasil DB diverifikasi.
8. Setelah migrate selesai, pastikan journal berstatus `success` atau `no_change`; simpan response version, `operation_id`, `manifest_sha256`, dan audit log DB sebagai bukti release.

## Release Stop Conditions

Stop release jika salah satu terjadi:
- `out/output.json.fail.count > 0`
- fixture engine merah
- `validate_parallel_repeatability.ps1` merah
- migration generated muncul dari perubahan whitespace-only
- REST config masih memakai placeholder DB/API key
- DB lock job template tidak mengeksekusi lock sebelum `migrate up`
- source digest binary Linux berbeda dari source MR atau complete release checksum gagal
- hasil dataset `rza_migrate_report` Windows dan Linux berbeda pada source migration yang sama
- failure checker tidak mencetak detail rule/file/message/suggestion yang tersedia ke job log
- patch source memberi hasil berbeda hanya karena casing filename pada Windows vs Linux

## Report layout V2

Report packaging gate: after changing layout publisher/tests, rebuild all Windows/Linux artifacts. Verify report default index/manifest and custom-root navigation from released binary, exact file checksums, embedded source identity, and unchanged Linux MR checker command. Checkout must preserve tools/** bytes. See [layout release contract](../report_output_layout.md).


## Audit object precision (2026-10-12)

Audit report object precision memakai complete six-binary Windows/Linux rebuild, embedded identity, source digest/checksums, released report smoke dan exact GitLab checker flags. Report fixture/corpus membuktikan no-write source/history. Tidak ada perubahan packaging, license contract atau pipeline flags.

Detail bukti: ../review_report_object_precision_20261012.md.
