# QA Gate - rza_migrate_db Engine Rules

2026-10-12 suspected JOIN identifier acceptance: QA-31 JoinReference.ParentId dan BusinessCode harus tampil tanpa index/IDENTITY walau table mempunyai PK Id. Evidence mapping/source modules/usage/index/identity wajib; Id enforced tidak diulang, NotJoinedId FK-only/inequality tidak menjadi candidate. Unit cases mencakup composite/filtered key dan self-role count. Integration golden memverifikasi source/history tetap identik serta output rerun idempotent. Full Go/hardening/build/checksum/source-digest/Linux gate dan Windows/Linux report parity tetap wajib; tidak ada safe-fixer baru untuk apply/plan.


Object scan refinement gate (2026-10-01): execute `TestExtractSQLAccessesPrecision`, compound client-result cases, quoted/reference negative regressions, `TestProcedureScannerQA31Golden`, same-line mapping/call tests, ambiguity/subquery cases, and `TestObjectScanPreservesCheckerAndGeneratorBoundaries`. The latter hashes authored SQL/history in checker plan/apply modes and verifies one genuine module delta with up/down plus idempotent second generation. Fuzz `FuzzSQLAccessSourceBounds` for malformed input/source span invariants. Then run full Go/hardening gates, full release build, checksum/source identity verification, and the exact GitLab Linux checker. Do not claim full grammar, unknown collation, multirow/compound exact mapping, or runtime dynamic resolution from these gates.


Dokumen ini adalah checklist QA untuk memastikan implementasi patuh `rules.md`, `spec.md`, `fsd.md`, dan `plan.md`.

## 1. Precondition

- Workspace bersih dari perubahan yang tidak relevan.
- Fixture test tersedia untuk kasus valid dan invalid.
- Binary `rza_migrate_check`, `rza_migrate_db`, dan `rza_get_object_db` dapat dijalankan.

## 2. Command Baseline

Jalankan minimal:

1. `rza_migrate_db fmt`
2. `rza_migrate_db verify`
3. `rza_migrate_db gen`
4. `rza_get_object_db -dry-run` (untuk validasi contract summary extractor tanpa write source)

Catat:

- Exit code.
- Ringkasan `out/output.json`.
- Daftar file migration baru (jika ada).
- Ringkasan `out/get_object_summary.json` (khusus extractor).

Lokasi fixture acuan:

- `fixtures/qa_rules/`
- Script pembaruan fixture: `scripts/populate_qa_fixtures.ps1`
- Script generate golden output: `scripts/generate_fixture_output_goldens.ps1`
- Script validasi fixture: `scripts/validate_qa_fixture_contract.ps1`
- Script validasi hygiene learning examples: `scripts/validate_learning_examples_hygiene.ps1`
- Script validasi kontrak migration lock: `scripts/validate_migration_lock_contract.ps1`
- Script validasi stop-behavior runner migrate: `scripts/validate_runner_stop_behavior_contract.ps1`
- Script validasi repeatability serial-vs-parallel: `scripts/validate_parallel_repeatability.ps1`
- Script benchmark baseline serial-vs-parallel: `scripts/benchmark_parallel_read.ps1`
- Script generate coverage report: `scripts/generate_rule_coverage_report.ps1`
- Script runner gate final: `scripts/run_final_hardening_gate.ps1`
- Output coverage report: `docs/qa/qa_coverage_report.md`
- Go test harness fixture: `tests/fixtures/fixture_golden_contract_test.go`
- Go runtime harness fixture (expectation execution): `tests/engine/engine_fixture_expectation_test.go`

## 3. Checklist QA Wajib

## QA-01 Hard Fail Rule

- Skenario: masukkan satu pelanggaran severity=3.
- Ekspektasi: exit code `1`, `fail.count > 0`, dan tidak ada migration baru tersimpan.

## QA-02 Folder Contract

- Skenario: buat subfolder di `migrations/`.
- Ekspektasi: verify fail.

## QA-03 Forbidden SQL

- Skenario: sisipkan `GO` atau `USE`.
- Ekspektasi: fail severity=3.

## QA-04 Pairing

- Skenario: hapus pasangan `.down.sql` dari satu migration.
- Ekspektasi: fail hard missing pair.
- Skenario tambahan: ubah nama pasangan sehingga metadata up/down tidak match (`category/object/vk` berbeda pada version sama).
- Ekspektasi tambahan: fail hard `MIGRATION_PAIR_MISMATCH`.

## QA-05 Version Format

- Skenario: gunakan version bukan 17 digit.
- Ekspektasi: fail parse filename.

## QA-06 Ordering Determinism

- Skenario: perubahan multi kategori pada satu run.
- Ekspektasi: urutan emit tables -> indexes -> views -> functions -> procedures -> patch.

## QA-07 `v{k}` Rule

- Skenario: ubah whitespace-only.
- Ekspektasi: `v{k}` tidak naik, migration baru tidak dibuat.

## QA-08 Reconstruction Integrity

- Skenario: history berisi file unsupported SQL subset.
- Ekspektasi: fail reconstruction (bukan warn).

## QA-09 Data Master Contract

- Skenario: filename table mismatch dengan object SQL.
- Ekspektasi: fail (`FILE_OBJECT_MISMATCH`).
- Skenario tambahan:
  - table snapshot mengandung statement terlarang (mis. `ALTER TABLE`, `DROP TABLE`, `CREATE INDEX`, wrapper `IF ...`).
  - index source file membuat index ke table lain.
  - programmability source berisi lebih dari satu object.
- Ekspektasi tambahan:
  - fail `TABLE_SNAPSHOT_FORBIDDEN_STATEMENT`
  - fail `INDEX_SOURCE_TABLE_MISMATCH` / `INDEX_SOURCE_CONTRACT_VIOLATION`
  - fail `PROGRAMMABILITY_MULTI_OBJECT` / `PROGRAMMABILITY_SOURCE_CONTRACT_VIOLATION`

## QA-10 Inline Default

- Skenario: inline `DEFAULT(...)` di table snapshot.
- Ekspektasi: fail (`INLINE_DEFAULT_NOT_ALLOWED`).

## QA-11 Destructive Policy

- Skenario: remove table dari desired state.
- Ekspektasi:
  - tidak auto `DROP TABLE`
  - engine membuat migration pair explicit manual-only (tables category)
  - migration up mencatat skip/warn `REQUIRES_MANUAL`
  - migration down mencatat warning `DATA_RESTORE_NOT_GUARANTEED`

## QA-12 Transaction Strategy

- Skenario: paksa runtime SQL error pada tables/indexes/patch.
- Ekspektasi: rollback, log fail, THROW, runner stop.
- Catatan ketat: token `THROW` yang hanya muncul di komentar/string tidak dianggap valid.

## QA-13 Audit Contract

- Skenario: cek file wrapper generated.
- Ekspektasi: ada severity header, `@run_key=NEWID()`, metadata deterministik.
- Jika header severity dihapus, verify harus fail (`WRAPPER_SEVERITY_HEADER_MISSING`).
- Ekspektasi tambahan untuk wrapper format baru (`DECLARE @file` + named audit args):
  - ada call `EXEC dbo.usp_mig_log_event`
  - ada decision `run`, `ok`, `fail`
  - pada CATCH ada action `migration_failed` dan reason `SQL_ERROR`
  - value `decision/action/reason/severity` yang ditulis ke audit harus berada pada whitelist rules
  - `@run_key` harus diinisialisasi dengan `NEWID()`
  - `@version/@category/@file` harus berupa literal deterministik (bukan ekspresi dinamis)
- Catatan legacy wrapper (audit call positional, tanpa named args):
  - engine sekarang strict-only (`fail`), sehingga audit call legacy yang belum lengkap akan fail `WRAPPER_AUDIT_LEGACY_WEAK`
  - mode `warn` tidak didukung lagi
  - CI/MR gate tetap wajib strict-only (`RZA_LEGACY_AUDIT_MODE=fail`) sesuai `docs/policy/legacy_wrapper_cutover.md`

## QA-13B Core Audit Init Contract

- Skenario: cek `000000__core__mig__audit_init.(up|down).sql`.
- Ekspektasi:
  - up memuat signature kontrak object `dbo.mig_event` dan `dbo.usp_mig_log_event` (missing => fail `CORE_AUDIT_CONTRACT_MISSING`)
  - down memuat signature cleanup `DROP` untuk object yang sama (missing => fail `CORE_AUDIT_CONTRACT_MISSING`)
  - lock proc readiness (`dbo.usp_mig_try_get_lock` + `sp_getapplock` + `THROW`) minimal terdeteksi sebagai readiness signal
  - execution pipeline yang memilih template migration wajib memakai `resource_group`; SQL Server session lock selama migrate disediakan driver `golang-migrate`

## QA-14 Down Policy

- Skenario: operasi destruktif lalu cek down migration.
- Ekspektasi: ada warning `DATA_RESTORE_NOT_GUARANTEED` jika relevan.
- Ekspektasi tambahan (drop-column multi-step):
  - SQL generated tidak boleh mengandung runtime discovery `sys.*` / `sp_executesql`.
  - Dependency drop yang diketahui dari history/source (constraint/index table lokal) harus emitted explicit dan deterministic.
  - Incoming FK lintas-table yang mereferensi kolom target harus emitted explicit `ALTER TABLE <ref_table> DROP CONSTRAINT <fk_name>` secara deterministic.
  - Jika side-effect drop index sudah dilakukan oleh migration tables pada run yang sama, generator tidak boleh emit migration indexes redundan untuk index yang sama.
  - Jika index source berubah, migration indexes harus memakai lifecycle `DROP INDEX` -> `CREATE INDEX` (bukan create-only).
  - Statement drop index generated harus memakai precheck deterministic `DROP INDEX IF EXISTS`.

## QA-15 Patch Canonicalization

- Skenario: patch folder name dengan spasi/simbol.
- Ekspektasi: canonical patch_name sesuai algoritma.

## QA-16 Patch Duplication

- Skenario: patch_name sudah pernah digenerate.
- Ekspektasi: skip generate tanpa compare isi.

## QA-17 Patch Source Integrity

- Skenario: source patch hilang setelah migration patch ada.
- Ekspektasi: fail hard.

## QA-17B Patch Immutability

- Skenario: patch migration baru sudah berisi `patch_up_sha256`/`patch_down_sha256`, lalu source patch diubah.
- Ekspektasi: verify fail hard `PATCH_SOURCE_MUTATED`.
- Catatan: migration patch legacy tanpa hash hanya menghasilkan warning `PATCH_IMMUTABILITY_UNVERIFIED` (non-fatal) sampai di-regenerate.

## QA-18 Output JSON Contract

- Skenario: jalankan verify/gen normal.
- Ekspektasi: root `fail/pass/new_migrations/notes` selalu ada.

## QA-19 Repo Hygiene

- Skenario: cek keberadaan `constraints_naming.md`, `.gitignore`, `.gitattributes`.
- Ekspektasi: file hilang/konten rekomendasi tidak lengkap menghasilkan warning non-blocking; bootstrap membuat file yang belum ada tanpa overwrite.

## QA-20 CLI and Exit Code Contract

- Skenario: jalankan `fmt|verify|gen` di kasus pass dan fail.
- Ekspektasi: exit `0` saat `fail.count==0`, exit `1` saat `fail.count>0`.

## QA-21 Wrapper Action Whitelist

- Skenario: wrapper generated-style (named audit args) berisi `@action` di luar whitelist.
- Ekspektasi: fail `WRAPPER_AUDIT_ACTION_INVALID`.

## QA-22 Wrapper Reason Whitelist

- Skenario: wrapper generated-style berisi `@reason` di luar whitelist.
- Ekspektasi: fail `WRAPPER_AUDIT_REASON_INVALID`.

## QA-23 Wrapper Run Key NEWID

- Skenario: wrapper generated-style tidak menginisialisasi `@run_key` menggunakan `NEWID()`.
- Ekspektasi: fail `WRAPPER_RUN_KEY_INVALID`.

## QA-24 Constraint Naming Hard Fail

- Skenario: table snapshot memakai constraint tanpa nama eksplisit (mis. `PRIMARY KEY (...)` tanpa `CONSTRAINT <name>`).
- Ekspektasi: fail hard `UNNAMED_CONSTRAINT_NOT_ALLOWED`, karena naming constraint adalah gate deterministik.

## QA-25 Wrapper Precheck Skip Deterministic

- Skenario: jalankan `gen` pada baseline objek baru (tables/indexes/patch).
- Ekspektasi:
  - wrapper generated untuk `tables/indexes` menyisipkan precheck+skip deterministik sebelum transaksi.
  - `tables up` memiliki precheck `OBJECT_ID(..., 'U') IS NOT NULL` + reason `TABLE_EXISTS`.
  - `tables down` memiliki precheck tolerant `OBJECT_ID(..., 'U') IS NULL` + reason `OBJECT_NOT_EXISTS`.
  - `indexes up` memiliki precheck `IF EXISTS (sys.indexes ...)` + reason `INDEX_EXISTS`.
  - `indexes down` memiliki precheck tolerant `IF NOT EXISTS (sys.indexes ...)` + reason `OBJECT_NOT_EXISTS`.
  - metadata wrapper tetap deterministik (`@version/@category/@file`) dan audit memakai named args (`@file_name = @file`).

## QA-26 ALTER COLUMN Precheck Deterministic

- Skenario: source table mengubah kolom dengan risiko data (`NULL -> NOT NULL`, shrink panjang string, type change ke numeric), parameter tipe dengan variasi whitespace, dan tipe bawaan ber-delimiter (`[DECIMAL](25,2) -> [DECIMAL](25,6)`, `VARCHAR (10) -> VARCHAR (24)`, binary/date-time precision).
- Ekspektasi:
  - `gen` tetap pass (tidak fail-hard) dan emit migration alter-column.
  - semantic hash/diff membedakan perubahan nilai precision/scale/length, tetapi menganggap perbedaan whitespace atau delimiter tunggal tipe bawaan saja sebagai ekuivalen.
  - setiap kolom yang berubah mendapat statement `ALTER COLUMN` target dan run `gen` kedua idempotent tanpa pair baru.
  - SQL emit menyisipkan precheck deterministic sebelum `ALTER COLUMN`:
    - null-check (`COUNT_BIG ... IS NULL`) untuk `NULL -> NOT NULL`
    - truncation-check (`DATALENGTH`) untuk shrink panjang string
    - conversion-check (`TRY_CONVERT`) untuk perubahan type ke numeric
  - saat precheck gagal, wrapper log `skip` severity `2` dengan reason sesuai risiko (`REQUIRES_MANUAL`/`NARROWING_WOULD_TRUNCATE`), melempar error, dan me-rollback transaksi.

## QA-GO-01 Extractor Warning Contract (`rza_get_object_db`)

- Skenario: jalankan extractor pada source yang memuat module bermasalah (mis. definition kosong/encrypted/header tidak valid).
- Ekspektasi:
  - run export tidak langsung fail-total hanya karena sebagian module bermasalah.
  - module bermasalah di-skip dan dicatat di `out/get_object_summary.json`.
  - field summary memuat:
    - `export.warning_count` > 0
    - `export.warnings[]` berisi detail object yang di-skip.
  - mode `-dry-run` tetap menghasilkan summary dengan struktur field konsisten.

## QA-27 Parallel Repeatability Determinism

- Skenario: jalankan `scripts/validate_parallel_repeatability.ps1`.
- Ekspektasi:
  - signature canonical output (`fail/pass/new_migrations`) stabil pada run berulang serial.
  - signature canonical output stabil pada run berulang parallel.
  - signature serial dan parallel identik (tidak ada drift behavior).

## QA-28 Parallel Benchmark Baseline

- Skenario: jalankan `scripts/benchmark_parallel_read.ps1`.
- Ekspektasi:
  - benchmark selesai tanpa fail runtime.
  - ringkasan tersimpan di `out/parallel_benchmark_summary.json`.
  - statistik serial/parallel (`avg/min/max`) tercatat untuk monitoring performa.

## QA-29 Cross-Platform Patch Source dan Console Failure

- Skenario pass: patch source berisi tepat `readme.md`, `UP.SQL`, dan `Down.Sql`, sedangkan header migration menggunakan casing canonical `README.md`, `up.sql`, dan `down.sql`.
- Ekspektasi pass: integrity hash cocok pada Windows dan Linux; tidak ada false `PATCH_SOURCE_MISSING`.
- Skenario fail: source patch benar-benar tidak ada atau salah satu dari tiga file logical tidak ada.
- Ekspektasi fail: exit code non-zero dan job log menampilkan seluruh fail item pada satu ordered output stream, termasuk `rule`, expected patch folder pada `file`, canonical patch pada `object`, migration asal pada `location`, `message`, serta `suggestion` untuk memeriksa Git tracking dan casing.
- Skenario ambiguity: file tambahan atau duplicate case-only pada filesystem yang mendukungnya.
- Ekspektasi ambiguity: `PATCH_FOLDER_INVALID_CONTENT`; engine tidak memilih salah satu file secara diam-diam.

## QA-30 Migration History Reconstruction dan Report Contract (`rza_migrate_report`)

- Skenario valid scan: jalankan `rza_migrate_report` pada repository dengan folder `migrations/`.
- Ekspektasi valid scan:
  - 8 report terbuat di `docs/database/migrations/` (7 CSV dan `report.json`) serta desired-state docs terbuat di `docs/database/`: root README/manifest, empat grouped Markdown, `procedures/README.md`, `procedures/procedure_impact_catalog.md`, dan satu profile untuk setiap source procedure; tidak ada folder relational diagram tambahan.
  - CSV valid RFC 4180 dan UTF-8 tanpa BOM.
  - JSON metadata memuat `"database_access": false` dan `"sql_execution": false`.
  - Pair bootstrap `000000__core__mig__audit_init`, jika ada, tidak muncul pada `migration_summary.csv`, `object_history.csv`, atau semantic change CSV; `excluded_pairs`/`excluded_migration_pairs` bernilai sesuai dan analyzed pair count tetap dapat direkonsiliasi dengan jumlah file fisik.
- Skenario isolasi:
  - Tidak ada koneksi database dan tidak ada SQL execution.
  - Folder `migrations/` dan file-filenya tidak pernah berubah, dibuat, atau dihapus (immutable).
- Skenario edge cases:
  - Folder migrations kosong: tetap membuat seluruh file CSV hanya dengan header baris dan exit 0.
  - Repository/migrations tidak ditemukan: menampilkan petunjuk jelas dan exit code 1.
  - Up tanpa down atau down tanpa up: dicatat terstruktur sebagai warning/error tanpa mengubah file migration.
  - Komentar dan variasi formatting SQL: tidak dianggap sebagai perubahan semantik.
  - Renamed constraint dengan definisi identik: tidak dilaporkan sebagai operasi semantik.
  - SQL ambigu/unsupported: confidence `LOW` dan notes `MANUAL_REVIEW_REQUIRED`.
  - Desired-state docs: enforcement PK/UQ/unfiltered unique index dipisahkan dari filtered unique, identity, dan ID review signals; FK/observed JOIN/module lineage berlabel serta report berbeda; dua kolom bernama sama tanpa predicate tidak boleh membentuk edge.
- Skenario determinism: dua kali run pada source yang sama menghasilkan output data CSV dan JSON yang identik.
- Metadata determinism: normalisasi/abaikan hanya `run_at` dan `generated_at`; seluruh field dan row lain harus identik walau run dilakukan pada detik berbeda.
- Skenario strict filename: version bukan 17 digit, category di luar whitelist, schema non-`dbo`, `V1`, `v0`, dan core non-canonical harus menjadi warning dan tidak diterima sebagai metadata canonical.
- Skenario semantic round-trip: down kosong dan down non-empty tetapi tidak mengembalikan column/constraint/index yang sama harus `FAIL` + warning; SQL unsupported harus `NOT_TESTED`, bukan false `PASS`.
- Skenario no-op/lifecycle: perubahan nullability tanpa perubahan type hanya menghasilkan `ALTER_NULLABILITY`; rename constraint dengan definisi identik menghasilkan nol operasi; down-only programmable tidak menghasilkan `CREATE`; drop lalu create programmable/type menghasilkan `DELETED` lalu `REINTRODUCED`.
- Skenario dense table longitudinal: v1 membuat beberapa kolom/PK/UQ/CHECK/DEFAULT/index; v2 dalam satu script mencampur dua add column, CHECK/UQ/FK composite, dua drop column, dua alter type+nullability, replacement CHECK/DEFAULT, rename column, dan drop+recreate index. Assert exact operation count/kind/target, old/new value, dependency, confidence, serta `PASS` round-trip.
- Skenario long-distance v1→v7: jalankan tujuh pair berurutan dengan add/drop/alter/rename, constraint/default/index replacement, `DROP TABLE` di v5, reintroduksi dengan schema berbeda di v6, dan operasi campuran di v7. Assert chain `previous_version`, lifecycle `INITIAL|UPDATED|UPDATED|UPDATED|DELETED|REINTRODUCED|UPDATED`, nilai before v7 dari v6, exact operation count, dan round-trip `PASS`.
- Skenario stale-state negatif: atribut yang hanya ada sebelum `DROP TABLE` tidak boleh ditemukan setelah create ulang; alter terhadap atribut tersebut harus `LOW` + `MANUAL_REVIEW_REQUIRED`, bukan memakai old value pra-drop.
- Skenario order-sensitive: add B → rename B ke C → alter C dan drop A → re-add A pada migration yang sama harus tetap HIGH confidence dan mempertahankan old value. Parser tidak boleh mengurutkan ulang berdasarkan jenis statement.
- Skenario wrapper create-table: `CREATE TABLE` setelah audit/precheck wrapper harus menghasilkan satu `CREATE_TABLE` dan tepat satu row untuk setiap child column/constraint; tidak boleh ada pengulangan akibat cursor parser berhenti sebelum header statement. Nama child column terdelimitasi `[]`/`""` dengan spasi harus tetap utuh dan tipe datanya benar.
- Skenario partial inverse: up menambah dua kolom dan down hanya menghapus satu harus `FAIL` + warning.
- Skenario output guard: `-out` atau `-docs` menuju `migrations/` atau `db_object/` (termasuk child path dan symlink/junction) ditolak, exit 1, dan hash/daftar file migration/db_object tidak berubah.

## QA-31 Desired-state Unique-ID dan Relationship Documentation

- Fixture: `fixtures/qa_rules/QA-31-desired-state-documentation/`; proof executable ada pada `internal/migreport/desired_state_test.go`.
- Ekspektasi: `table_unique_ids.md` memuat `PRIMARY_KEY`, `UNIQUE_CONSTRAINT`, unfiltered/filtered unique index, `IDENTITY`, `OBSERVED_JOIN_ID_SIGNAL`, `NONUNIQUE_INDEX_ID_SIGNAL`, enforcement, serta risk/action yang tidak menyebut signal sebagai key candidate. Constraint inline wajib terdeteksi.
- Ekspektasi: `relationship_catalog.md` membedakan `DECLARED_FK` dan `OBSERVED_JOIN`, memuat exact column pair/source/frequency/confidence; `relational_erd.md` memuat semua diagram, menggambar tepat satu line per non-self object pair, dan menampilkan type serta key/index context bagi kolom relationship; `module_lineage.md` memisahkan dependency module.
- Ekspektasi usability: narasi Indonesia natural dengan conventional English database terms; guidance classification tidak diulang sebagai template per row; stable codes tetap; tidak ada raw `<br>`; setiap Relational Map connected dan tidak melampaui 20 object, 24 actual relationship lines, 35 exact mappings, atau 9.000 karakter Mermaid; tidak ada artificial Hub/parent; setiap exact mapping dipertahankan dan tidak diganti usage, sedangkan dense pair memakai bounded numbered Mapping batch; object index dan shared-object table menyediakan navigasi; setiap catalog table maksimal 35 row; numbering Group/Map konsisten; self-join hanya ada pada review section.
- Negatif wajib: nama kolom sama tanpa predicate, bagian `=` dari `>=`/`<=`/`!=`, ekspresi dengan lebih dari satu resolved column per sisi, temp/table variable, stale alias lintas query, dan dynamic SQL tidak diberi relationship.
- No-write: hash/daftar `migrations/` dan `db_object/` sebelum/sesudah run sama; seluruh audit report ditulis di `docs/database/migrations/` dan seluruh dokumen desired-state ditulis di `docs/database/` di luar `migrations/` dan `db_object/`.
- Jalankan dari repository root dan folder lain dengan relative `-out`/`-docs`; resolved target harus tetap berada di bawah repository root. `report.json.output_directory`, `documentation_directory`, dan setiap `artifacts[]` harus menunjuk file konkret yang ada; wildcard hanya diizinkan pada compatibility `reports[]`.

### Procedure profile acceptance

- Fixture QA-31 dan test `internal/migreport/procedure_analysis_test.go` wajib membuktikan static catalog match, exact INSERT/UPDATE mapping, key/filter coverage, alias resolution per statement, source-preserved nested IF/ELSE guard, top-down branch/rejoin/error-aware local execution flow, transaction summarization yang tidak menyembunyikan child call, local object map, Mermaid-rendered line break tanpa literal `\n`, quoted Mermaid edge labels untuk operator SQL, readable per-call blocks, literal dynamic SQL, caller-literal provenance, unresolved runtime dynamic SQL, `EXEC @return_code = Procedure` sebagai static call, case-insensitive lineage deduplication, external/temp/system/missing/ambiguous status, cycle-safe recursion, complete call-edge retention saat recursive diagram dipecah pada batas 20 node, URL-safe profile links untuk identifier berkarakter khusus, dan no-write boundary.
- Negatif wajib: built-in function bukan UDF, table-valued function bukan table `NOT_FOUND`, cursor `FETCH ... FROM` bukan data object, CTE/schema-qualified temp bukan persisted object, CTE unqualified tidak menutupi schema-qualified table pada statement berikutnya, `sys.*` bukan missing application object, multiline `FROM`/`JOIN`/DML target tidak kehilangan object, INSERT tanpa target-column list tidak mengarang exact mapping, alias tidak bocor lintas statement, `Database..Object` tetap external, `EXEC @return = procedure` bukan dynamic SQL, serta assignment/`SELECT INTO`/subquery/`EXISTS`/`INSERT ... SELECT` bukan direct result set.
- Real-object smoke memakai `all_object/bond/bond-sqltrsretail/db_object`; jumlah profile harus sama dengan jumlah file procedure dan sample profile wajib direkonsiliasi dengan source line, direct calls, exact writes, serta key/index files.

## 4. Artefak Bukti QA

- Salinan `out/output.json` untuk skenario pass/fail.
- Salinan `docs/database/migrations/` dan `docs/database/` untuk skenario migration history report.
- Daftar command dan exit code.
- Daftar file migration baru yang dihasilkan.
- Catatan deviasi dan issue follow-up (jika ada).

## 5. Kriteria Kelulusan QA

Lulus jika:

1. Seluruh QA-01 s.d. QA-31 pass.
2. Tidak ada bypass fail-hard.
3. Tidak ada output non-deterministic pada fixture identik.

## Report layout V2

Layout V2 scenarios: run default tanpa out/; -docs saja memindahkan audit juga; explicit -out + -docs memiliki valid backlink; nested file/directory symlink dan root overlap gagal sebelum publication; existing manual collision tidak mengubah file; stale generated profile dihapus hanya jika hash owned cocok; edited stale/manual notes dipertahankan; equivalent legacy flat docs/CSV dipindahkan dengan conservative cleanup. Jalankan TestUnifiedDocumentationLayoutManifestAndLinks, TestReportCleanupPreservesManualFilesAndModifiedStaleProfiles, TestManualCollisionAndEditedCurrentOutputBlockBeforeWriting, TestLegacyLayoutRemovesOnlyEquivalentGeneratedFiles, TestOutputOverlapAndNestedSymlinkFailWithoutSourceWrites, TestPreviousManifestCannotDeleteAnUnlistedManualPath; QA-31 listing juga membuktikan layout fixture.

Layout relocation: TestDocumentationOwnershipSurvivesRepositoryClone membuktikan cloned output menerima genuine source delta tanpa menyentuh checkout lama; TestOutputFileSymlinkAndAliasedRootsCannotEscapePreflight membuktikan file symlink/root alias tidak menghasilkan partial output.
