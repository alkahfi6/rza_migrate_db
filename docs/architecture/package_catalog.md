# Katalog Source dan Tanggung Jawab Paket

Gunakan katalog ini setelah [overview](overview.md). Nama file adalah titik navigasi source saat ini; fungsi utama dan urutan eksekusi ada pada [call tree](call_tree.md). Tabel ini mencakup seluruh direktori langsung di `internal/`, termasuk paket yang tidak berhubungan dengan generation.

## Entrypoint

| Direktori | Peran | Mulai baca |
| --- | --- | --- |
| `cmd/rza_get_object_db` | CLI extractor dan config compiled/env/JSON | [main.go](../../cmd/rza_get_object_db/main.go), [config.go](../../cmd/rza_get_object_db/config.go) |
| `cmd/rza_migrate_check` | Adapter executable checker | [main.go](../../cmd/rza_migrate_check/main.go), lalu `internal/cmd/migrate_check.go` |
| `cmd/rza_migrate_db` | CLI authoring dan license generation | [main.go](../../cmd/rza_migrate_db/main.go) |
| `cmd/rza_migrate_report` | CLI report offline | [main.go](../../cmd/rza_migrate_report/main.go) |
| `cmd/rza_migrate_rest` | Adapter executable REST | [main.go](../../cmd/rza_migrate_rest/main.go), lalu `internal/cmd/rest.go` |
| `cmd/rza_migrate_test` | Runner test dengan credential `.nfm`, override catalog dan temporary catalog mapping | [main.go](../../cmd/rza_migrate_test/main.go) |
| `cmd/rza_mssql_test` | Utilitas status/up/down/reset dengan guard destructive tersendiri | [main.go](../../cmd/rza_mssql_test/main.go) |

## Paket internal

| Paket | Yang dimiliki | Source awal dan batas |
| --- | --- | --- |
| `buildinfo` | Build version, commit/time, source digest | [info.go](../../internal/buildinfo/info.go); metadata, bukan validator release lengkap. |
| `cmd` | Checker/REST CLI lifecycle dan logger | [migrate_check.go](../../internal/cmd/migrate_check.go), [rest.go](../../internal/cmd/rest.go); command authoring utama masih di `cmd/rza_migrate_db`. |
| `config` | Default/hard-lock authoring, path discovery, executable name lock | [app_config.go](../../internal/config/app_config.go), [executable_name_lock.go](../../internal/config/executable_name_lock.go); bukan config YAML REST. |
| `diff` | Drift record dan comparator ordering | [diff.go](../../internal/diff/diff.go); risk/policy serta SQL plan masih di `engine`. |
| `emit` | Filename pair dan version allocator | [emit.go](../../internal/emit/emit.go); tidak memiliki seluruh template/render/publish. |
| `engine` | Format/check/verify/gen, typed model, history reconstruction, dependency planning, round-trip, draft/seal | [engine.go](../../internal/engine/engine.go), [migration_plan.go](../../internal/engine/migration_plan.go), [gen_emit.go](../../internal/engine/gen_emit.go), [gen_emit_commit.go](../../internal/engine/gen_emit_commit.go). |
| `fsx` | Enumerasi stabil, layout canonical/legacy, logical filename resolution | [db_object_layout.go](../../internal/fsx/db_object_layout.go), [logical_files.go](../../internal/fsx/logical_files.go); mempertahankan file aktual untuk parity casing. |
| `getobject` | Akuisisi catalog, rendering snapshot, preservation unsupported, preview | [export.go](../../internal/getobject/export.go), [export_batch.go](../../internal/getobject/export_batch.go), [module_format.go](../../internal/getobject/module_format.go); tidak emit migration. |
| `handler` | Routing, auth, body validation, JSON/error mapping, middleware | [migration_handler.go](../../internal/handler/migration_handler.go), [swagger.yaml](../../internal/handler/swagger.yaml); hanya memanggil service interface. |
| `history` | Parse metadata filename, object key, pemilihan version | [meta.go](../../internal/history/meta.go); SQL replay/reconstruction berada di `engine`, audit analysis berada di `migreport`. |
| `license` | Generate/validate license yang dipakai tool terkait | [license.go](../../internal/license/license.go); checksum bukan license. |
| `management` | Manifest managed/unmanaged/skipped object | [manifest.go](../../internal/management/manifest.go); pengecualian eksplisit, bukan izin mengabaikan object hilang tanpa evidence. |
| `migreport` | Scan/analyze history dan desired, relationship/procedure evidence, output bundle ownership | [report.go](../../internal/migreport/report.go), [desired_state.go](../../internal/migreport/desired_state.go), [layout.go](../../internal/migreport/layout.go); tidak ikut menentukan migration delta. |
| `migrest` | Runtime config, migrator adapter, operation lock, file guard, journal/evidence | [config.go](../../internal/migrest/config.go), [service.go](../../internal/migrest/service.go), [journal.go](../../internal/migrest/journal.go); mengeksekusi migration yang sudah tersedia. |
| `output` | Report authoring dan diagnostic console | [report.go](../../internal/output/report.go), [console.go](../../internal/output/console.go); CLI menentukan exit dari error/fail bucket. |
| `requestmeta` | Context request/trace metadata | [context.go](../../internal/requestmeta/context.go); dipakai boundary service/logging. |
| `sqlfix` | Rewrite yang terbukti aman: wrapper catalog/batch, header, constraint | [rewrite.go](../../internal/sqlfix/rewrite.go); jangan menambahkan guessed business rewrite. |
| `sqlfmt` | Encoding, whitespace, canonical formatting dan safety | [encoding.go](../../internal/sqlfmt/encoding.go); semantic delta tetap harus terlihat. |
| `sqlparse` | Lexical token/span, semantic helpers, access/dependency, DDL documentation parsing | [scanner.go](../../internal/sqlparse/scanner.go), [access.go](../../internal/sqlparse/access.go), [dependencies.go](../../internal/sqlparse/dependencies.go); bukan parser grammar penuh. |
| `telemetry` | Setup OpenTelemetry/OTLP untuk REST | [telemetry.go](../../internal/telemetry/telemetry.go); tidak diperlukan authoring offline. |
| `tsqlast` | Module formatter backend abstraction | [format.go](../../internal/tsqlast/format.go), [lite_backend.go](../../internal/tsqlast/lite_backend.go); ANTLR unavailable pada build saat ini, consumer memakai lite/fallback. |
| `verify` | Orchestrator dan validator per domain | [verify.go](../../internal/verify/verify.go), [domain_migrations.go](../../internal/verify/domain_migrations.go), [domain_db_object.go](../../internal/verify/domain_db_object.go), [domain_wrapper.go](../../internal/verify/domain_wrapper.go); callback engine menghubungkan data reconstruction. |

## Bagian penting di dalam `engine`

| Area | File | Alasan perlu dibaca |
| --- | --- | --- |
| Shared safe corrections | [safe_fix_integration_test.go](../../internal/engine/safe_fix_integration_test.go), [constraint_name_reconcile.go](../../internal/engine/constraint_name_reconcile.go), [safe_index_rename.go](../../internal/engine/safe_index_rename.go) | Kontrak apply/plan, equivalence, ambiguity, idempotence. |
| Model dan capability | [sql_model.go](../../internal/engine/sql_model.go), [type_function_model.go](../../internal/engine/type_function_model.go), [capability.go](../../internal/engine/capability.go), [risk.go](../../internal/engine/risk.go) | Memisahkan parsed syntax, semantic identity, dan allowed transition. |
| State dan cache | [repository_snapshot.go](../../internal/engine/repository_snapshot.go), [reconstruct_cache.go](../../internal/engine/reconstruct_cache.go), [desired_cache.go](../../internal/engine/desired_cache.go), [fmt_cache.go](../../internal/engine/fmt_cache.go) | Cache harus terikat input yang benar; snapshot direset sesudah source write. |
| Reconstruct dan ordering | [semantic_reconstruction.go](../../internal/engine/semantic_reconstruction.go), [dependency_order.go](../../internal/engine/dependency_order.go), [column_rename.go](../../internal/engine/column_rename.go) | Urutan aman dan definisi sebelumnya menentukan up/down. |
| Render dan publication | [migration_render.go](../../internal/engine/migration_render.go), [roundtrip.go](../../internal/engine/roundtrip.go), [atomic_write.go](../../internal/engine/atomic_write.go), [gen_emit_commit.go](../../internal/engine/gen_emit_commit.go) | Staging verify, exclusive publish, rollback dan draft/seal. |

## Folder non-source

| Folder/file | Fungsi | Perlakuan |
| --- | --- | --- |
| `db_object/`, `patch/`, `migrations/` | Input authoring dan history | Ikuti [lifecycle](../desired_state_lifecycle.md); jangan jalankan generator pada repository asli sekadar untuk membaca arsitektur. |
| `fixtures/qa_rules/` | Input, expectations, golden QA | Kontrak fixture ada di [README fixture](../../fixtures/qa_rules/README.md). |
| `tests/engine/`, `tests/fixtures/` | Integrasi engine dan kontrak golden | Melengkapi unit test di tiap paket. |
| `examples/` | Contoh dan learning cases yang tersedia | Periksa keberadaan learning matrix; hardening script melewati langkah bila corpus opsional tidak ada. |
| `docs/database/` | Output dokumentasi database hasil report | Bedakan dari dokumentasi software di `docs/architecture/`; gunakan generator untuk memperbarui output miliknya. |
| `tools/` | Release Windows/Linux, license, checksums, identity | Release artifacts dari satu source; tidak diedit manual. |
| `scripts/`, `build_all.bat` | Validation, packaging, release source identity | Lihat [maintenance](../documentation_maintenance.md) dan production readiness. |
| `.gitlab/`, `.gitlab-ci.yml` | Gate authoring dan template deployment | Gate checker dan execution job mempunyai tanggung jawab berbeda. |
| `out/` | Diagnostic, cache, draft, staging lokal | Diabaikan Git; jangan memasukkan credential atau report lokal ke commit. |
