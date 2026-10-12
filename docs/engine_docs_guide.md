# Engine Docs Guide

Dokumen ini adalah peta dokumen utama untuk pengembangan engine suite:
- `rza_get_object_db` (extractor source-of-truth)
- `rza_migrate_check` (checker developer: safe-fix + verify tanpa generation, serta complete console diagnostics)
- `rza_migrate_db` (formatter/verifier/generator migration)
- `rza_migrate_rest` (REST executor migration SQL Server)

## 1. Hierarki Kebenaran Dokumen

Urutan prioritas jika ada konflik:

1. `rules.md`
2. `spec.md`
3. `fsd.md`
4. `plan.md`
5. `docs/traceability/rules_traceability_matrix.md`
6. `docs/qa/qa_gate_engine_rules.md`

## 2. Fungsi Tiap Dokumen

| File | Fungsi Utama | Dipakai Saat |
| --- | --- | --- |
| `README.md` | Gambaran cepat engine suite, alur operasional harian, dan command utama | Onboarding cepat |
| `rules.md` | Aturan final yang wajib dipatuhi (hard locks) | Semua fase |
| `spec.md` | Requirement baseline engine | Desain & review requirement |
| `fsd.md` | Desain teknis implementasi modul/flow | Implementasi kode |
| `plan.md` | Roadmap fase kerja + gate delivery | Eksekusi project |
| `docs/walkthrough.md` | Gambaran operasional aktual empat executable + flow end-to-end | Onboarding, handover, troubleshooting |
| `docs/traceability/rules_traceability_matrix.md` | Mapping rules 1-25 ke artefak teknis | Review coverage |
| `docs/qa/qa_gate_engine_rules.md` | Checklist QA skenario rules | Validasi sebelum release |
| `docs/qa/sql_server_capability_matrix.md` | Matrix fitur SQL Server supported/fail/manual/skip | Scope parser/generator dan triage unsupported SQL |
| `docs/policy/production_readiness.md` | Checklist observability, packaging, CI split, benchmark, dan runner ops | Release readiness |
| `.gitlab/merge_request_templates/rza_migrate_db_engine_rules.md` | Gate checklist untuk MR | Code review/MR |
| `.gitlab/issue_templates/rza_migrate_db_rule_exception.md` | Mekanisme deviasi rule sementara | Risk/exception handling |

## 3. Dokumen Pendukung Fixture dan Test

| File | Fungsi |
| --- | --- |
| `fixtures/qa_rules/` | Fixture QA-01 s.d. QA-26 |
| `examples/rules_2x_learning/` | Data learning pass/fail + lifecycle `fmt/gen/verify` untuk rules prioritas |
| `docs/qa/learning_edge_case_backlog.md` | Backlog edge-case learning yang sudah/akan ditambahkan ke examples matrix |
| `scripts/verify_rules_2x_learning.ps1` | Validasi massal step-based (`fmt/gen/verify`) vs matrix, termasuk assert `fail.rule`, `pass.reason`, dan migration generated |
| `scripts/populate_qa_fixtures.ps1` | Generate/refresh input fixture |
| `scripts/generate_fixture_output_goldens.ps1` | Generate `output.golden.json` |
| `scripts/validate_qa_fixture_contract.ps1` | Validasi kontrak fixture |
| `scripts/validate_learning_examples_hygiene.ps1` | Memastikan `examples/**/out/output.json` tidak tersimpan di repo |
| `scripts/stress_reconstruct_perf.ps1` | Stress test rekonstruksi history besar (contoh 10k objects) + budget check |
| `scripts/validate_migration_lock_contract.ps1` | Validasi template migration job selalu acquire DB lock sebelum `migrate up` |
| `scripts/validate_runner_stop_behavior_contract.ps1` | Validasi template runner tidak menelan error `migrate up` dan lock dieksekusi sebelum migrate |
| `scripts/generate_rule_coverage_report.ps1` | Generate laporan coverage QA + residual risk checklist |
| `scripts/run_final_hardening_gate.ps1` | Runner otomatis gate final (test + fixture + learning + report) |
| `scripts/run_release_readiness_gate.ps1` | Runner release-readiness terpisah: fast unit, fixture, repeatability, benchmark, build |
| `.gitlab/migration_jobs/migrate_with_lock.yml` | Template opsional execution job: GitLab `resource_group` + session lock bawaan driver `golang-migrate`; tidak di-include oleh authoring gate |
| `.gitlab/rza_migrate_check.gitlab-ci.yml` | Gate authoring read-only pada MR, `sit`, `uat`, dan tag, diisolasi pada special stage `.pre` |
| `docs/policy/legacy_wrapper_cutover.md` | Kebijakan cutover mode legacy wrapper ke strict-only CI |
| `docs/policy/rebaseline_playbook.md` | Rencana terkontrol untuk re-baseline saat growth migration sangat besar |
| `docs/policy/parallel_safe_design.md` | Desain paralelisasi aman untuk `rza_migrate_db` (deterministic + fail-hard) |
| `docs/policy/rest_api_hardening.md` | Baseline hardening deployment `rza_migrate_rest` (dev/staging/prod) |
| `tests/fixtures/fixture_golden_contract_test.go` | Test kontrak golden fixture |
| `tests/engine/engine_command_smoke_test.go` | Smoke test command engine |
| `tests/engine/engine_fixture_expectation_test.go` | Test eksekusi engine QA-01..QA-26 terhadap `expectation.json` |
| `tests/engine/config_hard_lock_test.go` | Test fail-hard saat config lock invalid |
| `tests/engine/pair_integrity_test.go` | Test mismatch metadata pair up/down |
| `tests/engine/core_audit_contract_test.go` | Test kontrak core audit init dan lock readiness |
| `tests/engine/patch_immutability_test.go` | Test patch hash header dan deteksi mutasi source patch |
| `tests/engine/source_contract_strict_test.go` | Test strict source contract untuk table/index/programmability |
| `internal/verify/domain_patch_test.go` | Test patch source integrity dan parity actual-entry casing Windows/Linux |
| `internal/output/console_test.go` | Test seluruh field failure tampil tanpa truncation di console lokal/CI |

## 4. Urutan Baca Saat Mulai Coding Engine

1. Mulai dari `README.md` untuk konteks cepat engine suite.
2. Baca `rules.md` section yang terkait change.
3. Konfirmasi requirement di `spec.md`.
4. Ambil desain implementasi dari `fsd.md`.
5. Cek fase dan gate di `plan.md`.
6. Baca `docs/walkthrough.md` untuk flow operasional aktual (`rza_get_object_db` -> `rza_migrate_db` -> `rza_migrate_rest` opsional).
7. Pastikan coverage ada di `docs/traceability/rules_traceability_matrix.md`.
8. Tentukan skenario validasi dari `docs/qa/qa_gate_engine_rules.md`.

## 5. Workflow Saat Implementasi

1. Ubah kode sesuai `spec.md` + `fsd.md`.
2. Update fixture bila behavior berubah:
   - `scripts/populate_qa_fixtures.ps1`
   - `scripts/generate_fixture_output_goldens.ps1`
3. Jalankan:
   - `powershell -ExecutionPolicy Bypass -File scripts/validate_qa_fixture_contract.ps1`
   - `powershell -ExecutionPolicy Bypass -File scripts/verify_rules_2x_learning.ps1`
   - `powershell -ExecutionPolicy Bypass -File scripts/validate_learning_examples_hygiene.ps1 -Fix`
   - `powershell -ExecutionPolicy Bypass -File scripts/validate_learning_examples_hygiene.ps1`
   - `powershell -ExecutionPolicy Bypass -File scripts/validate_migration_lock_contract.ps1`
   - `powershell -ExecutionPolicy Bypass -File scripts/validate_runner_stop_behavior_contract.ps1`
   - `powershell -ExecutionPolicy Bypass -File scripts/validate_parallel_repeatability.ps1`
   - `powershell -ExecutionPolicy Bypass -File scripts/benchmark_parallel_read.ps1`
   - `powershell -ExecutionPolicy Bypass -File scripts/generate_rule_coverage_report.ps1`
   - `go test ./...`
4. Untuk final pass terkontrol, jalankan satu command:
   - `powershell -ExecutionPolicy Bypass -File scripts/run_final_hardening_gate.ps1`
5. Jika mengubah extractor (`rza_get_object_db`), verifikasi:
   - `out/get_object_summary.json` terisi konsisten
   - `export.warning_count`, `export.warnings[]`, dan `export.warning_items[]` konsisten (module + table skip)
   - `skip_unsupported_tables=true` menghasilkan warning terstruktur `TABLE_EXTRACT_SKIPPED_UNSUPPORTED` (non-fatal)
   - `skip_unsupported_tables=false` tetap fail-hard untuk unsupported table feature
   - backend formatter default tetap `lite` (opsi `antlr` fallback ke `lite`)
6. Jika melakukan tuning performa `rza_migrate_db`, jalankan dengan `-perf-notes true` (atau `RZA_PERF_NOTES=true`) dan evaluasi `out/output.json.notes.perf_ms`; untuk repo besar, uji juga `-parallel-read true -parallel-workers <n>` (effective worker akan di-cap CPU/task). Untuk debugging cepat lokal bisa pakai `-fail-fast true` (default CI disarankan tetap full-report).
7. Update traceability jika ada rule baru/berubah.
8. Gunakan template MR rules gate saat membuat MR.
9. Untuk perubahan error reporting, verifikasi `out/output.json` dan console memuat diagnosis yang sama; console wajib cukup informatif saat artifact GitLab tidak dapat dibuka.

## 6. Aturan Update Dokumen

- Perubahan behavior engine tanpa update dokumen terkait dianggap belum selesai.
- Jika aturan `rules.md` berubah, wajib cek dampak ke:
  - `spec.md`
  - `fsd.md`
  - `plan.md`
  - `docs/walkthrough.md`
  - `docs/traceability/rules_traceability_matrix.md`
  - fixture QA dan golden output.

## 7. Definition of Ready (Coding Task)

Task coding dianggap siap dikerjakan jika:

1. Rule target sudah jelas di `rules.md`.
2. Requirement sudah punya rujukan di `spec.md`.
3. Desain implementasi ada di `fsd.md`.
4. Flow operasional terdampak sudah dicek di `docs/walkthrough.md`.
5. Skenario QA tersedia atau sudah ditambahkan di `fixtures/qa_rules/`.

## 8. Catatan Enforcement Bertahap

- Wrapper migration `tables/indexes/patch` tetap wajib transaction contract (`TRY/CATCH`, `ROLLBACK`, `THROW`).
- Untuk wrapper format baru hasil generator (`DECLARE @file` + named audit args), verify menegakkan audit contract strict:
  - wajib ada `dbo.usp_mig_log_event`
  - wajib decision `run/ok/fail`
  - wajib action `migration_failed` dan reason `SQL_ERROR` pada CATCH
  - wajib value audit (`decision/action/reason/severity`) sesuai whitelist rules
  - wajib `@run_key` diinisialisasi `NEWID()`
- Wrapper legacy (umumnya positional audit call) tetap tervalidasi oleh contract dasar.
- Default mode engine sekarang strict (`fail`), sehingga wrapper legacy yang audit-nya belum lengkap akan fail-hard.
- Mode `warn` tidak didukung lagi.
- Pipeline/CI wajib strict-only (`RZA_LEGACY_AUDIT_MODE=fail`) sesuai `docs/policy/legacy_wrapper_cutover.md`.
- Prioritas resolusi mode: flag CLI > env var > default `fail`.

## Report layout V2

Report layout reading reference: [Database report output layout](report_output_layout.md) owns grouped output roots, path compatibility, manifest ownership/cleanup, and report-only cross-tool boundaries. It does not alter source-of-truth/history semantics.
