# Development Plan - rza_migrate_db Engine Suite

2026-10-22 documentation delivery: tersedia [arsitektur](docs/architecture/overview.md),
[call tree](docs/architecture/call_tree.md), [katalog paket](docs/architecture/package_catalog.md),
[referensi engine](docs/engine_reference.md), [panduan agent](docs/engineering/agent_onboarding.md),
dan [tutorial offline](docs/tutorials/offline_authoring.md) yang telah dijalankan.
[Maintenance guide](docs/documentation_maintenance.md) menjelaskan praktik dan aturan pemeliharaan;
[review record](docs/qa/documentation_review_20261022.md) mencatat baseline, validation, dan batasnya.
Scope hanya dokumentasi implementasi yang sudah ada; tidak menaikkan status dukungan SQL atau readiness engine.

2026-10-12 delivery scope: add report-only suspected JOIN identifiers with mapping/module/index/identity evidence, including unindexed/non-ID-named columns on keyed tables. Source parser, schema actions, auto-corrections, migration history, and cross-tool responsibilities stay unchanged. Acceptance: focused positive/negative tests, executable QA-31 golden, report no-write/rerun, full hardening/build and Windows/Linux parity. No data-level uniqueness certification is introduced.


2026-10-01 object scan refinement: precision and boundary fixes implemented with shared lexical/reference tests, executable QA-31 scan golden, fuzz coverage, and checker/generator integration. Delivery gate remains full tests/hardening, complete Windows/Linux release rebuild, source/checksum identity, and Linux MR checker smoke. Full AST, runtime constant propagation, synonym metadata, function-body lineage, and complex INSERT expression mapping remain outside this delivery; see [review and residual scope](docs/review_object_scan_20261001.md).


Dokumen ini adalah rencana eksekusi implementasi engine berdasarkan `rules.md`, `spec.md`, dan `fsd.md`.

## 1. Target Plan

- Menutup seluruh gap aturan yang pernah belum termapping.
- Menjadikan implementasi deterministic, verifiable, dan fail-hard sesuai kontrak.
- Menyiapkan urutan pengerjaan yang bisa langsung dipakai tim engineering.
- Menjaga sinkronisasi lima aplikasi utama: extractor -> checker -> generator -> pipeline/REST executor, dengan report offline sebagai pembaca desired/history; utilitas live test mempunyai boundary tersendiri.

## 2. Hasil Compare Ulang terhadap `rules.md`

Gap utama yang terdeteksi saat compare:

- G01: whitelist `action/reason/decision` belum eksplisit.
- G02: patch policy belum lengkap (canonicalize, duplication, source integrity).
- G03: locking `sp_getapplock` belum termapping.
- G04: MVP supported subset belum rinci.
- G05: ALTER COLUMN precheck map belum detail.
- G06: policy opsional `constraints_naming.md` dan safe-fix constraint belum dikunci.
- G07: repo hygiene `.gitignore/.gitattributes` non-blocking belum eksplisit.
- G08: engine config baseline `app_config.go` belum lengkap.
- G09: verify checklist belum mencakup semua lock.
- G10: output.json contract belum detail ke field wajib/recommended.
- G11: core audit init dan audit logging belum cukup rinci.
- G12: wrapper standard tables/indexes/patch belum preskriptif.
- G13: down migration policy belum lengkap.
- G14: skip-vs-stop semantics belum terkunci penuh.
- G15: traceability section 1-25 rules belum eksplisit.

Status sekarang:

- Semua gap G01-G15 sudah ditutup di dokumen revisi `spec.md` dan `fsd.md`.
- Plan ini memastikan gap tersebut tidak muncul lagi di implementasi.
- Refinement fase 5-6 sudah berjalan: orchestrator verify dipindah ke `internal/verify`, dan surface `internal/engine` diperkecil lewat bridge callback terpisah.

## 3. Prinsip Eksekusi

- Build validation dan parser dulu, baru diff dan emit.
- Setiap fase harus punya test dan acceptance gate.
- Tidak ada merge fase tanpa rule-trace evidence.
- Semua failure severity=3 wajib memblokir emit.

## 4. Workstreams

## WS-A: Foundation dan Config Locks

Cakupan:

- bootstrap command CLI
- executable-name lock untuk tiga binary
- `app_config.go` locked defaults
- output writer contract

Rules terkait:

- section 21, 22, 15, 23

## WS-B: Source Validation dan Formatter

Cakupan:

- scanner deterministic
- formatter safe mode
- validation `db_object` dan `patch`

Rules terkait:

- section 2, 4, 7, 14, 19, 20

## WS-C: Parser dan Reconstruction

Cakupan:

- parser filename tanpa regex
- SQL subset parser
- history reconstruction

Rules terkait:

- section 5, 6.1, 6.2, 17

## WS-D: Diff, Policy Engine, Emitter

Cakupan:

- semantic diff
- destructive policy
- wrapper emitter
- down emitter
- patch emitter

Rules terkait:

- section 8, 9, 10, 11, 12, 13, 14, 18

## WS-E: Verify Gate dan Hardening

Cakupan:

- final verify integrity migrations
- fail-hard no emit
- full integration tests

Rules terkait:

- section 3, 15, 22.3, 23, 24, 25

## 5. Milestone Implementasi

## 5.1 Status Modularisasi Package (Phase 0-6)

| Package | Status | Catatan |
| --- | --- | --- |
| `internal/fsx` | Selesai | dipakai scanner/formatter |
| `internal/sqlfmt` | Selesai | dipakai fmt dan emit normalizer |
| `internal/sqlparse` | Selesai | dipakai lexical scan + parser helper |
| `internal/history` | Selesai | dipakai parser filename + history helpers |
| `internal/diff` | Selesai | dipakai drift validation + deterministic sort |
| `internal/emit` | Selesai | dipakai allocator version + builder filename |
| `internal/verify` | Selesai | memegang orkestrasi urutan validasi + validator concrete domain (`migrations`, `forbidden_sql`, `repo_hygiene`, `patch`, `db_object`, `wrapper`, `table_inline_default`) |
| `internal/output` | Selesai | writer output.json |
| `internal/getobject` | Selesai | extractor SQL Server ke `db_object/*` + summary `out/get_object_summary.json` |
| `internal/tsqlast` | Selesai (fallback mode) | formatter modul SQL shared (`lite`, opsi `antlr` fallback) |

## M0 - Project Bootstrap

Deliverables:

- struktur folder Go sesuai FSD
- main command router (`fmt`, `verify`, `gen`, default `verify_then_gen`)
- command extractor (`rza_get_object_db`: `export`/`dry-run`/`preview`)
- output writer skeleton

Done criteria:

- compile sukses
- command tiga executable bisa dijalankan

## M1 - Config Hard Locks

Status per 2026-02-19: Completed.

Deliverables:

- `app_config.go` berisi semua default wajib
- enum whitelist severity/decision/action/reason
- whitelist category dan forbidden token

Done criteria:

- unit test validasi enum dan config pass
- config mismatch menghasilkan fail

Evidence:

- `internal/config/app_config.go`
- `internal/config/app_config_test.go`
- `tests/engine/config_hard_lock_test.go`

## M2 - Scanner dan Formatter

Status per 2026-02-19: Completed.

Deliverables:

- file discovery stable sorted
- formatter LF/trim/final newline
- preserve literal/comment

Done criteria:

- fixture formatter pass
- whitespace-only tidak memicu semantic change

Evidence:

- `internal/engine/engine.go` (`runFmt`, `normalizeSQLText`, `listFilesRecursive`)
- `tests/engine/gen_emit_contract_test.go` (`TestGenNoEmitOnWhitespaceOnlyScenario07`)

## M3 - Source Validator

Status per 2026-02-19: Completed.

Deliverables:

- validasi `db_object` kontrak lengkap
- validasi patch folder dan immutability
- validasi forbidden SQL token lexical-safe

Done criteria:

- fail cases (inline default, mismatch object, patch file extra) terdeteksi

Evidence:

- `internal/engine/engine.go` (`validateForbiddenSQL`, `validateTableInlineDefault`, `validatePatchFolders`, `validateDbObjectConsistency`)
- `tests/engine/engine_fixture_expectation_test.go`

## M4 - Filename Parser dan Pair Integrity

Status per 2026-02-19: Completed.

Deliverables:

- parser structured schema/patch filenames
- pair up/down checker
- duplicate version conflict checker

Done criteria:

- semua kasus pair-safe pass
- invalid grammar fail hard

Evidence:

- `internal/engine/engine.go` (`parseMigrationFileName`, `validatePairing`)
- `tests/engine/pair_integrity_test.go`

## M5 - SQL Parser MVP dan Reconstruction

Status per 2026-02-19: Completed.

Deliverables:

- parser subset tables/indexes/programmability/wrapper
- history reconstruction state map
- unsupported syntax fail-hard

Done criteria:

- notes berisi last_seen_version/vk/hash
- conflict history memicu fail

Evidence:

- `internal/engine/gen_emit.go` (`semanticFromMigrationWithState`, `reconstructHistory`)
- `tests/engine/gen_emit_contract_test.go` (`TestGenUnsupportedHistorySubsetScenario08`)

## M6 - Semantic Diff Engine

Status per 2026-02-19: Completed.

Deliverables:

- desired state builder
- AST semantic comparator
- policy mapping create/alter/drop

Done criteria:

- no-change tidak emit
- semantic-change menaikkan `vk`

Evidence:

- `internal/engine/gen_emit.go` (`planGeneration`, `collectDesiredObjects`, `classifyTableColumnChanges`)
- `tests/engine/gen_emit_contract_test.go`

## M7 - Migration Emitter

Status per 2026-02-19: Completed.

Deliverables:

- version allocator `YYYYMMDDHHMMSS+seq`
- deterministic ordering by category/object
- wrapper templates tables/indexes/patch
- programmability emitter `CREATE OR ALTER`
- down migration emitter

Done criteria:

- generated files valid, pair lengkap, seq <= 999

Evidence:

- `internal/engine/gen_emit.go` (`allocateVersion`, `renderMigrationContent`, `renderTxWrapper`)
- `tests/engine/gen_emit_contract_test.go`

## M8 - Patch Engine Complete

Status per 2026-02-19: Completed.

Deliverables:

- patch_name canonicalization
- patch duplication skip rule
- patch source integrity validator
- patch mapping ke `pass.items`

Done criteria:

- semua skenario patch policy pass

Evidence:

- `internal/engine/gen_emit.go` (`canonicalizePatchName`, `planPatchMigrations`, patch hash header emitter)
- `internal/engine/engine.go` (`validatePatchSourceIntegrity` dengan source integrity + immutability hash check)
- `tests/engine/patch_immutability_test.go`

## M9 - Verify Gate Final

Status per 2026-02-19: Completed.

Deliverables:

- full verify migrations integrity
- wrapper/audit call minimum validation
- final gate untuk `gen`

Done criteria:

- fail -> no emit
- pass -> commit emit

Evidence:

- `internal/verify/verify.go` (verify orchestrator)
- `internal/engine/verify_callbacks.go` (typed callback bridge)
- `internal/engine/gen_emit.go` (`gen_post_emit` full verify checks + rollback file write on fail)
- `tests/engine/engine_fixture_expectation_test.go`

## M10 - Core Audit and Locking Compliance

Status per 2026-02-19: Completed.

Deliverables:

- generator support untuk core audit init files
- template compliance `dbo.mig_event`, `dbo.usp_mig_log_event`, optional `dbo.usp_mig_try_get_lock`
- runtime lock guidance artifact

Done criteria:

- audit init pair tervalidasi
- lock failure semantics terdokumentasi

Evidence:

- `internal/engine/engine.go` (`validateCoreAuditInitPair`, `validateCoreAuditContracts`)
- `tests/engine/core_audit_contract_test.go`
- `docs/qa/qa_gate_engine_rules.md` (QA-13B)

## M11 - Test and Hardening

Status per 2026-02-19: Completed.

Deliverables:

- unit, golden, integration tests
- rule-trace coverage matrix
- CI gate checks

Done criteria:

- seluruh test pass
- regression suite tersedia untuk rule 1-25 dan QA-01 s.d. QA-26

Evidence:

- `go test ./...` pass
- `scripts/validate_qa_fixture_contract.ps1` pass
- `scripts/verify_rules_2x_learning.ps1` pass
- `scripts/run_final_hardening_gate.ps1` pass
- `docs/qa/qa_coverage_report.md` generated
- `tests/engine/*` + `tests/fixtures/fixture_golden_contract_test.go`

## M12 - Extractor-Migrator Sync Hardening

Status per 2026-02-26: Completed.

Deliverables:

- kontrak sinkronisasi `rza_get_object_db` vs `rza_migrate_db` terdokumentasi (`docs/walkthrough.md`, `docs/engine_docs_guide.md`).
- QA gate extractor warning contract (`QA-GO-01`) ditambahkan.
- command default `rza_migrate_db` (verify -> gen) dan flag `-gen-auto-indent` didokumentasikan.
- kontrak output extractor (`out/get_object_summary.json`) dipastikan masuk baseline dokumen.
- baseline telemetry performa fase `rza_migrate_db` tersedia via opsi `-perf-notes true` (`notes.perf_ms`).
- opsi parallel-safe read (`-parallel-read`, `-parallel-workers`) tersedia untuk akselerasi bounded tahap collect source pada `gen`.
- akselerasi parallel-safe juga mencakup preload history (`reconstructHistory`) dengan reduce serial agar deterministic.
- akselerasi parallel-safe juga mencakup scan forbidden SQL (verify/precheck) dengan deterministic merge.
- verify orchestrator mendukung `fail-fast` opsional untuk feedback lokal cepat; default tetap full-report agar kualitas laporan tidak turun.
- kontrak dual-path `db_object` diperketat: warning legacy layout + fail canonical-vs-legacy conflict.
- planner table menambah guard deterministic untuk dependency unnamed (`UNNAMED_CONSTRAINT_BLOCKS_DETERMINISTIC_PLAN`).
- semantic hash programmability mengabaikan style-only bracket/semicolon untuk mencegah false migration emit.
- extractor contract unsupported table disinkronkan (`skip_unsupported_tables` -> warning/fatal by config).
- tersedia stress script reconstruction history + perf budget (`scripts/stress_reconstruct_perf.ps1`).

Done criteria:

- dokumentasi utama (`spec.md`, `fsd.md`, `plan.md`) konsisten dengan behavior aktual code.
- QA docs memuat validasi run extractor (`-dry-run`) beserta warning summary.
- scenario learning tambahan untuk edge-case sinkronisasi sudah tervalidasi (`51` s.d. `54`).

Evidence:

- `docs/walkthrough.md`
- `docs/engine_docs_guide.md`
- `docs/qa/qa_gate_engine_rules.md` (`QA-GO-01`)
- `cmd/rza_get_object_db/main.go`
- `cmd/rza_migrate_db/main.go`
- `internal/verify/domain_db_object.go`
- `internal/engine/gen_emit.go`
- `internal/getobject/export.go`
- `scripts/stress_reconstruct_perf.ps1`

## 6. Rule Traceability Checklist (Implementasi)

Checklist ini wajib ditandai saat PR merge:

- RC-01: hard fail rule severity=3/no emit.
- RC-02: folder contracts db_object/patch/migrations.
- RC-03: no GO dan forbidden SQL.
- RC-04: pairing/version/filename grammar.
- RC-05: ordering emit dan `vk` rules.
- RC-06: reconstruction fail-hard.
- RC-07: data master contracts (table/index/programmability).
- RC-08: destructive policy dan skip-vs-stop.
- RC-09: transaction and throw strategy.
- RC-10: audit schema dan logger contract.
- RC-11: wrapper standards by category.
- RC-12: down policy.
- RC-13: patch policy end-to-end.
- RC-14: output.json schema and semantics.
- RC-15: locking readiness.
- RC-16: MVP subset restrictions.
- RC-17: ALTER COLUMN precheck templates.
- RC-18: constraints_naming and repo hygiene.
- RC-19: CLI contract and exit codes.
- RC-20: verify integrity final.

## 7. Quality Gates per Pull Request

Wajib sebelum merge:

1. test suite yang relevan lulus.
2. tidak ada pelanggaran RC-01 s.d. RC-20.
3. output deterministic untuk fixture yang sama.
4. `gen` tidak menulis file saat fail.
5. traceability update pada changelog internal.
6. kontrak extractor summary (`out/get_object_summary.json`) tervalidasi minimal via `rza_get_object_db -dry-run`.

## 8. Matrix Compare `rules.md` (1-25)

| Rule | Section `rules.md` | Referensi `spec.md` | Referensi `fsd.md` | Status |
| --- | --- | --- | --- | --- |
| 1 | Istilah dan Konsep | 2 | 1, 4, 9 | Closed |
| 2 | Foldering Repo | 4 | 5, 7 | Closed |
| 3 | Tooling dan Pipeline | 5 | 4 | Closed |
| 4 | Larangan dan Safety Guard | 6 | 7 | Closed |
| 5 | Pairing dan Filename | 7 | 6, 7.3 | Closed |
| 6 | Ordering, v{k}, Reconstruction | 8 | 8, 9, 10 | Closed |
| 7 | Data Master Validation dan Format | 9 | 6.4, 7.1 | Closed |
| 8 | Policy Destructive dan Semantics | 10, 11 | 9.3 | Closed |
| 9 | Transaction dan Error Strategy | 12 | 10.3, 15 | Closed |
| 10 | Core Audit Init | 13.1 | 10, 14 | Closed |
| 11 | Audit Logging | 13.2, 13.3 | 10.3, 12 | Closed |
| 12 | Wrapper Standar | 14 | 10.3, 11 | Closed |
| 13 | Rules Down Migration | 15 | 10.4 | Closed |
| 14 | Patch Policy | 16 | 7.2, 10.5 | Closed |
| 15 | Output Generator | 17 | 3.3, 12 | Closed |
| 16 | Locking 1 run per DB | 18 | 14 | Closed |
| 17 | MVP Supported Subset | 19 | 6.3, 16 | Closed |
| 18 | ALTER COLUMN Precheck Map | 20 | 9.3 | Closed |
| 19 | constraints_naming.md | 21 | 7.1 | Closed |
| 20 | Repo Hygiene | 21 | 7.1, 16.4 | Closed |
| 21 | CLI Contract | 22 | 4 | Closed |
| 22 | Engine Config dan Verify Integrity | 23, 24 | 2, 3, 11 | Closed |
| 23 | Final Locks | 3, 25, 26 | 15, 17 | Closed |
| 24 | Contoh Error Patterns | 27 | 12, 16.2 | Closed |
| 25 | Runner Cheatsheet | 27 | 4.3, 15 | Closed |

## 9. Risk Register dan Mitigasi

- Risiko parser subset terlalu sempit.
  Mitigasi: tambah fixture incremental dan fail-fast reason jelas.

- Risiko output nondeterministic.
  Mitigasi: enforce sorting global dan golden diff check.

- Risiko patch drift tidak terdeteksi.
  Mitigasi: integrity check patch source vs migration metadata.

- Risiko regression pada down migration.
  Mitigasi: dedicated down fixtures dan policy assertions.

- Risiko dokumen dan implementasi tidak sinkron.
  Mitigasi: gate review wajib terhadap `spec.md` dan `fsd.md`.

## 10. Definition of Done

Implementasi dianggap selesai jika:

1. `fmt`, `verify`, `gen` memenuhi CLI contract.
2. seluruh RC-01 s.d. RC-20 lulus otomatis.
3. seluruh gap G01-G15 tetap tertutup pada kode.
4. dokumentasi tetap sinkron dengan perilaku engine aktual.

## 11. Enforcement Artifacts

Artefak yang wajib dipakai selama development:

1. Template MR GitLab: `.gitlab/merge_request_templates/rza_migrate_db_engine_rules.md`
2. Checklist QA gate: `docs/qa/qa_gate_engine_rules.md`
3. Matrix traceability rules: `docs/traceability/rules_traceability_matrix.md`
4. Template issue exception: `.gitlab/issue_templates/rza_migrate_db_rule_exception.md`
5. CI docs/fixture gate: `.gitlab-ci.yml`
6. QA fixture scaffold: `fixtures/qa_rules/`
7. Fixture population script: `scripts/populate_qa_fixtures.ps1`
8. Fixture validation script: `scripts/validate_qa_fixture_contract.ps1`
9. Fixture golden generator: `scripts/generate_fixture_output_goldens.ps1`
10. Go fixture test harness: `tests/fixtures/fixture_golden_contract_test.go`
11. Dokumen peta acuan coding: `docs/engine_docs_guide.md`
12. Coverage report generator: `scripts/generate_rule_coverage_report.ps1`
13. Runner gate final: `scripts/run_final_hardening_gate.ps1`
14. Walkthrough operasional tiga executable: `docs/walkthrough.md`
15. Summary extractor contract artifact: `out/get_object_summary.json` (runtime output, tidak di-commit)

Aturan penggunaan:

1. Setiap MR engine wajib memakai template MR rules gate.
2. Setiap kandidat release wajib lulus checklist QA gate.
3. Setiap perubahan aturan harus update matrix traceability.
4. Setiap request exception rule wajib memakai issue template exception.
5. Pipeline wajib meloloskan `docs_gate` dan `fixture_gate`.

## 12. Gap Re-Assessment (2026-02-19)

Refinement M1-M5 sudah menutup gap behavior prioritas yang sempat terbuka:

- G16 closed: table alter existing tidak lagi fallback ke snapshot `CREATE TABLE`; jalur supported memakai DDL alter executable.
- G17 closed: down migration drop-column sudah restore kolom + dependency (constraint/index) secara deterministic pada subset supported.
- G18 closed: `verify` kini menjalankan reconstruction parser subset sehingga parity terhadap `gen` terjaga.
- G19 closed: dependency plan tetap deterministic tanpa runtime discovery `sys.*`/dynamic SQL.
- G20 closed: drop-table manual-only tetap sebagai pair explicit `skip+warn`.

Update lanjutan sinkronisasi:

- G21 closed: dual-path legacy `db_object/schema/*` kini punya sinyal eksplisit (`LEGACY_DB_OBJECT_LAYOUT_DETECTED`) dan fail conflict (`DB_OBJECT_DUAL_PATH_CONFLICT`).
- G22 closed: dependency unnamed pada alter/drop column tidak lagi jatuh ke false-fail acak; diperlakukan deterministic `skip+warn` (`UNNAMED_CONSTRAINT_BLOCKS_DETERMINISTIC_PLAN`).
- G23 closed: style-only churn programmability (`;`, bracket identifier) tidak lagi memicu emit migration.
- G24 closed: unsupported table extractor bisa dimitigasi operasional via `skip_unsupported_tables=true` dengan warning terstruktur.
- G25 closed: tersedia harness stress untuk budget rekonstruksi history besar.

Catatan operasional:

- SQL di luar subset MVP tetap fail-fast (`UNSUPPORTED_SQL_SUBSET`) dan wajib manual migration.
- Wrapper legacy tidak lagi memiliki mode transisi; `LegacyAuditMode` dikunci strict-only (`fail`).
- Mulai 2026-02-19 pipeline CI dikunci strict-only (`RZA_LEGACY_AUDIT_MODE=fail`) dan cutover governance terdokumentasi di `docs/policy/legacy_wrapper_cutover.md`.
- Kontrak migration job lock sebelum `migrate up` dibakukan melalui `.gitlab/migration_jobs/migrate_with_lock.yml` dan diverifikasi oleh `scripts/validate_migration_lock_contract.ps1`.

## 13. Refinement Milestones (R1-R7)

### R1 - Verify/Gen Rule Parity

Deliverables:

- Tambah validator subset MVP (`tables/indexes`) ke flow `verify`.
- Pastikan skenario unsupported source fail di `verify` (bukan hanya di `gen`).

Acceptance:

- Skenario source computed column/columnstore index gagal di `verify` dengan `UNSUPPORTED_SQL_SUBSET`.

### R2 - Executable Table Alter Emitter

Deliverables:

- Ubah planner table diff agar tidak emit `CREATE TABLE` untuk object existing.
- Tambah builder DDL alter untuk:
  - add column
  - alter safe column
  - drop column (tetap policy-compliant)

Acceptance:

- Migration up untuk table existing selalu executable pada table existing.
- Tidak ada `CREATE TABLE` pada migration alter table non-create.

### R3 - Deterministic Dependency Plan (No Runtime sys.* Discovery)

Deliverables:

- Ganti builder drop-column runtime discovery menjadi dependency plan berbasis state parse/reconstruction.
- Hindari dynamic `sys.*` lookup untuk mencari nama dependency saat runtime migration.

Acceptance:

- SQL emit drop dependency bersifat deterministic dari input repo/history.
- Tidak ada query discovery `sys.*` untuk resolve nama dependency saat execution path normal.

Status update 2026-02-19:

- Runtime lookup `sys.*` + `sp_executesql` sudah dihapus dari jalur drop-column.
- Emit dependency drop sekarang deterministic dari reconstruction state untuk constraint named (table local), impacted index di table yang sama, dan incoming FK antar-table.
- Pola FK yang tidak dapat diparse deterministic (mis. unnamed/unsupported form) diperlakukan fail-fast (`REQUIRES_MANUAL`), bukan runtime discovery.
- Lifecycle index pada jalur schema diff sudah diseragamkan: index diff emit `DROP INDEX` + `CREATE INDEX` deterministic, source index yang dihapus emit drop migration sekali, dan reconstruction menyerap side-effect drop index dari migration tables.
- Precheck existence deterministic untuk drop index diterapkan via `DROP INDEX IF EXISTS ... ON ...` agar lebih tahan drift tanpa runtime discovery `sys.*`.

### R4 - Valid Down Migration Strategy

Deliverables:

- Builder down khusus table-diff:
  - add-column up -> drop-column down
  - drop-column up -> add-column down (dengan warning data restore)
  - alter-column up -> reverse alter/down policy sesuai safety

Acceptance:

- Down migration mengembalikan schema ke state sebelumnya sejauh didukung policy.
- Warning `DATA_RESTORE_NOT_GUARANTEED` tetap muncul untuk operasi destruktif.

### R5 - Drop Table Manual-only Explicit Pair

Deliverables:

- Emit migration pair explicit saat table dihapus dari desired state:
  - up: precheck + skip + warn `REQUIRES_MANUAL`
  - down: warning `DATA_RESTORE_NOT_GUARANTEED` + instruction manual

Acceptance:

- Ada audit trail migration file untuk event manual-only drop table.
- Tidak ada auto `DROP TABLE` statement.

### R6 - QA Fixture Expansion

Deliverables:

- Tambah QA skenario baru:
  - table alter add column executable
  - verify fails unsupported source subset
  - down migration executable contract
  - drop table explicit skip-warn migration pair

Acceptance:

- Fixture expectation + golden + tests diperbarui dan pass.

### R7 - Dokumentasi dan Traceability Realignment

Deliverables:

- Update status matrix dari `Closed` menjadi `Open` untuk rule terdampak sampai fix merged.
- Tambahkan section "Known Gaps" dan "Target Fix Version" pada dokumen utama.

Acceptance:

- Tidak ada mismatch antara claim dokumen dan perilaku engine aktual.

## Production compatibility alignment (2026-07-28)

Locked decisions:

- Initial `CREATE TABLE` wrappers use `bootstrap-compatible` matching when the
  target table already exists.
- Blocking facts are limited to required-column existence, compatible base
  type, identity-kind, and computed-kind.
- Length, precision/scale, nullability, collation, defaults, constraints,
  constraint runtime state, extra columns, and physical index options do not
  block the initial-table skip path.
- Explicit `ALTER`, `DROP`, `ADD/DROP CONSTRAINT`, and `CREATE/DROP INDEX`
  operations retain strict operation-specific guards.
- `rza_get_object_db` emits logical CHECK/FK/PK/UQ definitions even when
  observed constraints are disabled/untrusted or PK/UQ backing indexes use
  non-default physical metadata.
- Extractor health findings are structured and never mutate MSSQL.
- `-require-healthy-constraints` is the optional CI gate for promoting
  disabled/untrusted findings to a non-zero exit after evidence is written.
- `rza_migrate_check` records the active profile and rejects newly generated
  initial-table wrappers that reintroduce exact-count, length, nullability, or
  constraint-catalog blockers.

Acceptance:

- Historical migrations remain byte-immutable.
- Newly generated initial-table wrappers pass the profile contract beginning
  at version `20260728000000000`.
- Full tests, race tests, repository verify, repeatability, benchmark, and the
  final hardening gate pass.

## Cross-platform patch diagnostics hardening (2026-09-08)

Status: completed.

Update 2026-09-10: report menghasilkan lima desired-state Markdown documents. `table_unique_ids.md` memisahkan uniqueness enforcement, conditional index, `IDENTITY`, dan ID review signal. Relasi dibagi menjadi `relational_erd.md`, exact-evidence `relationship_catalog.md`, dan `module_lineage.md`; nama legacy `table_relationship_tree.md` menjadi compatibility pointer. Observed edge hanya berasal dari static equality JOIN dengan exact alias-column mapping, bukan dari nama kolom yang sama.

Usability update 2026-09-10: seluruh Mermaid berada dalam satu `relational_erd.md`, disusun sebagai Relationship Group → connected Relational Map. Map dibentuk adaptif dengan hard limit 20 object, 24 actual relationship lines, 35 exact mappings, dan 9.000 karakter Mermaid. Artificial Hub/parent dihapus; node menampilkan object type serta key/index context bagi relationship columns, sedangkan edge mempertahankan exact mappings tanpa statistik usage. `[FK]`/`[JOIN]` hanya dipakai untuk mixed-evidence pair. Dense pair memakai bounded numbered `Mapping batch`; object index dan shared-object table menjaga navigasi saat object muncul pada beberapa map. Catalog mempertahankan semua exact mappings dengan Group/Map yang sama; self-join masuk review section.

Documentation routing update 2026-09-10: output dokumentasi desired-state dialihkan dari `docs/database/migrations/` ke `docs/database/` (dapat dikonfigurasi via flag `-docs`) agar terdokumentasi rapi di version control repository, sedangkan audit report tabular (7 CSV dan `report.json`) tetap tersimpan di `docs/database/migrations/` (dikonfigurasi via `-out`). Guard diperluas untuk memastikan `-out` dan `-docs` tidak berada di dalam `migrations/` maupun `db_object/`.

Procedure impact update 2026-09-11: report menghasilkan `procedures/README.md`, global impact catalog, dan satu profile per procedure. Profile mempertahankan full source sequence, top-down branch-aware local execution flow dengan transaction/condition/error pattern dan Mermaid-rendered line breaks, local object map, source-preserved IF/ELSE guards, direct serta cycle-safe recursive calls, complete split recursive diagrams, exact INSERT/UPDATE mapping, key/index context, functions/views, transaction/error/result signals, dan dynamic SQL evidence. Semua persisted object resolution divalidasi terhadap `db_object`; ambiguous/runtime-dependent target tidak ditebak. Real-object validation menggunakan 1.058 procedure pada snapshot `all_object/bond/bond-sqltrsretail`.

Procedure recognition hardening 2026-09-30: line-oriented access detection diganti shared source-positioned token scanner. Multiline object, statement-scoped CTE, DML alias scope, temporary identity, dan direct result-set semantics mempunyai regression coverage. Relative report roots kini repo-based dan `report.json` memuat concrete artifact manifest setelah documentation write sukses. Extractor/checker/generator tetap tidak memakai hasil inference report; release gate tetap mewajibkan rebuild penuh.

- Patch source filenames are resolved from actual directory entries using the
  same case-insensitive logical contract on Windows and Linux.
- Exactly one `README.md`, `up.sql`, and `down.sql` logical entry is required;
  extras and case-only ambiguity remain fail-hard.
- Checker and generator share complete, non-truncated console failure details,
  including actionable `PATCH_SOURCE_MISSING` path and migration context.
- Acceptance requires targeted resolver/output tests, full test suite, rebuilt
  Windows/Linux artifacts, source digest verification, and GitLab-equivalent
  Linux smoke test.

## Migration history report and reconstruction tool (2026-09-08)

Status: completed.

- Executable `rza_migrate_report.exe` and Linux binary `tools/linux-amd64/rza_migrate_report`
  provide migration history reconstruction solely from `migrations/`.
- Absolute read-only boundaries: zero DB connections, zero SQL execution, zero access to
  `dbo.mig_event` or system catalog tables, and zero mutation/creation/deletion in `migrations/`.
- Independent command isolated in `cmd/rza_migrate_report/` and shared logic in `internal/migreport/`.
- Emits 7 RFC 4180 CSV reports and 1 metadata JSON report to `docs/database/migrations/`:
  `report_summary.csv`, `migration_summary.csv`, `object_history.csv`, `table_changes.csv`,
  `index_changes.csv`, `programmable_changes.csv`, `type_changes.csv`, `report.json`.
- Automatic repository discovery (exec dir -> cwd -> parents containing `migrations/`),
  interactive double-click console pause (`[done] press Enter to exit`), and non-interactive CI mode.
- Strict cross-tool linearity and release artifact parity via `build_all.bat`.
- Desired-state documentation hardening selesai: inline PK/UQ, filtered/non-unique indexes, unenforced identity risk, exact JOIN column evidence, operator/alias-scope false-positive guards, dynamic-SQL review list, dan real-snapshot audit empat repository.
- Markdown usability hardening selesai: narasi Indonesia yang natural dipadukan dengan conventional English database terms, template risk/action dipindahkan ke reference section, tanpa raw HTML `<br>`, diagram memakai adaptive connected Relational Maps (maksimum 20 object/24 lines/35 mappings/9.000 karakter) tanpa artificial Hub dalam satu file, node memuat key/index context, exact mapping tampil pada edge label, catalog mengikuti Group/Map yang sama, dan self-join dipisahkan dari diagram.

Hardening update (2026-09-09):

- report scanner sekarang memakai strict filename parser shared `internal/history`;
- output di dalam `migrations/` ditolak sebelum write;
- round-trip memvalidasi inverse operasi semantik dan tidak lagi meluluskan down kosong/unrelated;
- no-op `ALTER COLUMN`, metadata-only constraint rename, phantom empty-up programmable/type, serta lifecycle delete/reintroduce diperbaiki;
- determinism test berjalan lintas detik dan membandingkan seluruh data setelah mengecualikan hanya metadata waktu run.
- dense mixed-table hardening menambah ordered source-offset parsing, mixed multi-item `ALTER TABLE`, composite constraint coverage, exact before/after assertions, partial-inverse rejection, dan semantic collapse untuk constraint/index replacement.
- 2026-09-11 report recovery: scanner `CREATE TABLE` yang berada setelah wrapper audit/precheck sekarang memajukan cursor melewati header dan body penuh; operasi create serta child column/constraint dipancarkan tepat sekali. Parser juga mempertahankan identifier `[]`/`""` berspasi; keduanya diuji terhadap snapshot Bond SQLTRSRetail.
- 2026-09-11 report boundary update: bootstrap audit canonical `000000__core__mig__audit_init` dikeluarkan dari history aplikasi pada report, tanpa mengubah file migration atau validasi release-nya.
- 2026-09-23 ALTER COLUMN hardening: parser/hash/diff parameterized type kini mengonsumsi precision, scale, length, serta date-time precision walaupun source memiliki spasi di sekitar kurung/koma. Regression end-to-end membuktikan emit `ALTER COLUMN` untuk decimal, character, binary, datetime, dan base numeric type serta idempotence run kedua.
- 2026-09-24 bracketed built-in hardening: base type extractor seperti `[DECIMAL]` dinormalisasi lintas typed model, risk classifier, diff, dan runtime guard. Regression kasus `(25,2) -> (25,6)` membuktikan guarded migration dibuat, perubahan delimiter saja no-op, serta alias schema-qualified tetap konservatif.
- 2026-09-11 report audit hardening: summary memisahkan analyzed/excluded migration pair; generated profile links meng-escape karakter `#`; dependency lineage di-deduplicate case-insensitive; dan `EXEC @return_code = Procedure` tidak lagi menjadi false-positive dynamic SQL. Unit/CLI tests dan snapshot Bond SQLTRSRetail menjadi bukti penerimaan.
- longitudinal state hardening menambah bukti v1→v7 dengan drop di v5, reintroduksi di v6, perubahan campuran di v7, serta reset state table/index agar incarnation pra-drop tidak bocor.

## Report layout V2

Report layout V2: grouped publication, navigation, conservative ownership cleanup, QA-31 coverage dan CLI/integration proof. Release gate: full Go suite, hardening/fixture gates, full build_all.bat, source/checksum identity, dan smoke Linux MR/report. Tidak ada perubahan authoring/diff planning. Lihat [delivery contract](docs/report_output_layout.md).

Report layout V2 delivery complete (2026-10-12): full tests/hardening, complete release build 20261012.020616, Windows/Linux shared-checkout report regeneration, manifest/link/hash validation, and Linux MR smoke passed. Learning examples are the existing missing-directory skip; no new parser/schema/migration residual scope. Details in docs/report_output_layout.md.


## Audit object precision (2026-10-12)

Audit object precision selesai pada implementation dan focused regression; completion gate tetap full Go/hardening, complete Windows/Linux rebuild, source/checksum identity dan released checker/report smoke. Residual scope: full T-SQL AST, CTE projection, correlated scope, runtime dynamic metadata dan MERGE branch mapping tetap manual; lihat audit bertanggal untuk bukti dan hasil gate.

Detail bukti: docs/review_report_object_precision_20261012.md.
