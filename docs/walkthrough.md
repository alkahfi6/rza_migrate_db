# Walkthrough Engine Suite

Dokumen ini menjelaskan cara kerja aktual lima executable pada repo ini:
- `rza_get_object_db` (extractor source-of-truth)
- `rza_migrate_check` (checker developer tanpa generation)
- `rza_migrate_db` (formatter + verifier + generator migration)
- `rza_migrate_report` (rekon histori perubahan database / declared object change report)
- `rza_migrate_rest` (REST executor migration SQL Server)

Dokumen acuan utama tetap:
1. `rules.md`
2. `spec.md`
3. `fsd.md`
4. `plan.md`

## 1. Tujuan

Engine suite ini dibuat untuk memastikan perubahan database SQL Server berjalan:
- deterministic (hasil repeatable)
- terkontrol (fail-hard saat rule dilanggar)
- audited (wrapper + audit contract)
- cocok untuk CI/CD (`golang-migrate` sebagai runner eksekusi)

## 2. Komponen dan Tanggung Jawab

### `rza_get_object_db`
- Menarik object `dbo` dari database SQL Server.
- Menulis snapshot ke:
  - `db_object/tables/`
  - `db_object/indexes/`
  - `db_object/views/`
  - `db_object/functions/`
  - `db_object/procedures/`
- Menormalkan header programmability menjadi `CREATE OR ALTER`.
- Menulis ringkasan run ke `out/get_object_summary.json`.

### `rza_migrate_db`
- `fmt`: format source SQL (`LF`, trim trailing space, final newline).
- `verify`: validasi folder contract, pairing, wrapper, audit, forbidden SQL, patch integrity.
- `gen`: rekonstruksi history, semantic diff, emit migration pair up/down, verify akhir.
- Menulis hasil ke `out/output.json`.

### `rza_migrate_check`
- Ditujukan untuk double-click atau hook sebelum commit.
- Default `profile=authoring`: merapikan dan memvalidasi penuh `db_object/` serta `patch/`, dengan pemeriksaan integrity-only pada `migrations/`.
- Profile `full` tersedia untuk audit semantik migration; history lama tidak diperiksa memakai rule semantik terbaru pada workflow harian.
- Menulis hasil ke `out/output.json` dengan `generation_enabled=false`.
- Tidak menjalankan diff/emit dan tidak mengubah file apa pun di `migrations/`.
- Pada failure, mencetak semua detail item ke console agar job log GitLab cukup untuk diagnosis tanpa membuka artifact.

### `rza_migrate_report`
- Menghasilkan laporan rekon histori perubahan database semata-mata dari `migrations/` sebagai satu-satunya source of truth.
- Zero DB connection, zero SQL execution, dan tidak membaca katalog sistem database atau `dbo.mig_event`.
- Tidak membaca `db_object/` dan tidak mengubah file migration apa pun (read-only).
- Mendukung auto-discovery repository root (exec dir -> cwd -> parent dirs).
- Menghasilkan 7 laporan CSV RFC 4180 (UTF-8 tanpa BOM) dan 1 JSON metadata di `docs/database/migrations/`, serta dokumentasi desired-state di `docs/database/`:
  - `docs/database/migrations/datasets/report_summary.csv`
  - `docs/database/migrations/datasets/migration_summary.csv`
  - `docs/database/migrations/datasets/object_history.csv`
  - `docs/database/migrations/datasets/table_changes.csv`
  - `docs/database/migrations/datasets/index_changes.csv`
  - `docs/database/migrations/datasets/programmable_changes.csv`
  - `docs/database/migrations/datasets/type_changes.csv`
  - `docs/database/migrations/report.json`
  - `docs/database/identifiers/table_unique_ids.md`
  - `docs/database/relationships/relational_erd.md`
  - `docs/database/relationships/relationship_catalog.md`
  - `docs/database/relationships/module_lineage.md`
  - `docs/database/README.md` dan `docs/database/manifest.json`
  - `docs/database/procedures/README.md`
  - `docs/database/procedures/procedure_impact_catalog.md`
  - `docs/database/procedures/profiles/dbo.<Procedure>.md`
- Menampilkan console summary interaktif dan pause Enter saat dijalankan via double-click (`-pause-on-exit false` untuk non-interactive/CI).

### `rza_migrate_rest`
- Menjalankan migration (`golang-migrate`) via REST API.
- Membaca runtime `config.yaml`.
- Menyediakan endpoint:
  - `GET /healthz`
  - `GET /api/v1/migrations/version`
  - `POST /api/v1/migrations/up` (opsional `steps`)
  - `POST /api/v1/migrations/down` (opsional `steps`)
  - `POST /api/v1/migrations/run-file` (single-step deterministic)
- Menggunakan `X-API-Key` (default aktif) + DB lock (`dbo.usp_mig_try_get_lock`) agar aman.
- Default hardening:
  - `allow_down=false`
  - `allow_run_specific_file=false`
  - timeout operasi best-effort (hardcoded di service)

### Runner eksekusi
- `rza_migrate_rest` adalah wrapper eksekusi migrate berbasis API.
- Engine `rza_migrate_db` tetap tidak mengeksekusi migration ke DB.

## 3. Kontrak Folder

- `db_object/` = desired state
- `migrations/` = execution history (wajib flat)
- `patch/` = sumber patch one-time (tepat satu logical `README.md`, `up.sql`, `down.sql`). Kapitalisasi nama diterima case-insensitive dan resolver memakai nama entry aktual, sehingga checkout Windows dan Linux menghasilkan keputusan yang sama; file tambahan/duplikat tetap gagal.

Core audit migration wajib ada:
- `000000__core__mig__audit_init.up.sql`
- `000000__core__mig__audit_init.down.sql`

## 4. Flow `rza_get_object_db`

1. Validasi executable name lock.
2. Load config compile-time dari `cmd/rza_get_object_db/config.go`, atau JSON runtime melalui `-config <path>`.
3. Koneksi ke SQL Server.
4. Ambil table + index snapshot.
5. Ambil modul (`view/function/procedure`) dari `sys.sql_modules`.
6. Format modul dengan backend default `lite`.
7. Tulis file ke `db_object/*` (parallel render/write).
8. Tulis summary ke `out/get_object_summary.json`.

### Perilaku error
- Table/index extraction: computed dan `SPARSE NULL` direkonstruksi; metadata yang belum lossless (mis. FILESTREAM, temporal, trigger, partition/filegroup, advanced index catalog) fail-hard atau menjadi structured skipped-warning sesuai konfigurasi.
- Modul programmability: object bermasalah di-skip, run tetap lanjut, warning dicatat di:
  - `export.warning_count`
  - `export.warnings[]`

### Mode khusus
- `-dry-run`: hitung perubahan tanpa menulis file.
- `-preview-input/-preview-output`: format sample SQL lokal tanpa koneksi DB.

## 5. Flow `rza_migrate_db`

### Command routing
- `rza_migrate_db fmt`
- `rza_migrate_db verify`
- `rza_migrate_db gen`
- Tanpa command: default `verify` lalu `gen`.
- Optional profiling aman: `-perf-notes true` (atau `RZA_PERF_NOTES=true`) untuk menulis durasi per fase ke `notes.perf_ms`.
- Optional local fast stop: `-fail-fast true` (atau `RZA_FAIL_FAST=true`) untuk berhenti setelah fail pertama.
- Optional acceleration aman: `-parallel-read true -parallel-workers <n>` (atau env `RZA_PARALLEL_READ` / `RZA_PARALLEL_WORKERS`) untuk paralelisasi bounded di tahap collect source, history preload `gen`, dan scan forbidden SQL.

### Pipeline `gen`
1. `fmt` internal
2. precheck `verify`
3. optional auto-indent untuk source yang benar-benar berubah
4. reconstruct history dari `migrations/*.up.sql`
5. diff desired vs history; deklarasi tipe parameterized dibaca utuh dan
   whitespace di sekitar kurung/koma dinormalisasi. Karena itu perubahan
   `DECIMAL(17,5) -> DECIMAL(17,9)`, panjang character/binary, date-time
   precision, dan base type tidak dapat tereduksi menjadi no-op. Bentuk bawaan
   hasil extractor seperti `[DECIMAL](25,2)` memakai model yang sama dengan
   `DECIMAL(25,2)`, sehingga perubahan scale menghasilkan guarded migration.
6. emit pair migration
7. verify hasil emit
8. tulis `out/output.json`

Extractor mengeluarkan bentuk source yang diperoleh secara kanonik tanpa
membuat migration. Checker apply dan `-check` membaca normalisasi tipe yang sama
tanpa menyentuh `migrations/`; hanya generator yang mengubah delta parameter
nyata menjadi pair up/down. Binary Linux pipeline membawa perilaku shared yang
sama setelah release artifact dibangun ulang.

### Fail-hard
- Jika ada severity `3` atau `fail.count > 0`, exit code `1`.
- Tidak boleh ada hasil migration final baru pada run yang fail.

## 5.1 Flow `rza_migrate_check`

1. Validasi executable name dan license yang sama dengan engine utama.
2. Format `db_object/**` dengan perlindungan comment dan multiline literal.
3. Bangun ulang repository snapshot dari hasil format.
4. Jalankan verify penuh terhadap `db_object`, `migrations`, `patch`, dan wrapper.
5. Tulis `out/output.json`; `new_migrations.files` selalu kosong.
6. Jika dijalankan tanpa argumen dari double-click, tunggu Enter agar hasil console dapat dibaca.

## 5.2 Flow `rza_migrate_report`

1. Validasi executable name lock `rza_migrate_report.exe`.
2. Deteksi repository root: flag `-repo`, direktori executable, current working directory, atau parent directories yang memiliki `migrations/`.
3. Scan seluruh file migration `*.up.sql` dan `*.down.sql` di `migrations/`, memvalidasi format nama file dan pairing up/down.
   Parser filename adalah parser strict shared engine (core canonical, 17 digit, category whitelist, `dbo.<object>`, `v{k} >= 1`).
   Pair bootstrap canonical `000000__core__mig__audit_init` tetap discan read-only, tetapi dikeluarkan sebelum analisa history aplikasi. Jumlah pair yang dianalisis dan yang dikecualikan dilaporkan terpisah.
4. Parsing isi SQL tiap migration menggunakan tokenizer/AST SQL, membuang noise komentar (`--`, `/* */`) dan whitespace formatting.
5. Analisis perubahan semantik:
   - Rekonstruksi siklus hidup objek (`INITIAL`, `UPDATED`, `DELETED`, `REINTRODUCED`).
   - Ekstraksi perubahan semantik table (kolom, constraint, default, tipe data).
   - Ekstraksi perubahan index (kolom, include, filter, unique, clustering).
   - Ekstraksi perubahan programmability dan komparasi hash definisi (`changed_sections`).
   - Ekstraksi perubahan tipe data user-defined.
   - Penilaian reversibilitas dan confidence level (`HIGH`, `MEDIUM`, `LOW` / `MANUAL_REVIEW_REQUIRED`).
   - Pencocokan inverse semantik up/down; inverse hilang menjadi `FAIL` + warning, dan SQL yang tak dapat dipastikan menjadi `NOT_TESTED`.
   - Suppression no-op `ALTER COLUMN` dan rename constraint dengan definisi identik; lifecycle programmable/type meliputi `DELETED` dan `REINTRODUCED`.
   - Ekspansi seluruh item dalam statement campuran: multi-column, multi-constraint, composite FK/UQ, multiple drop, `WITH CHECK`/`WITH NOCHECK`, rename, dan index drop+recreate.
   - Satu `CREATE TABLE` beserta child column/constraint dipancarkan tepat sekali; cursor scanner melewati header dan seluruh body walaupun migration didahului wrapper audit/precheck. Identifier column terdelimitasi `[]`/`""` tetap satu nama walaupun memuat spasi.
   - Reducer mengikuti source offset, bukan grouping jenis statement, sehingga perubahan berantai di migration yang sama memakai state before/after yang tepat.
   - Reducer berjalan sepanjang semua version object. `DROP TABLE`/create ulang mereset state table dan index, lalu version berikutnya menggunakan incarnation terbaru; skenario v1→v7 menguji deletion, reintroduction, dan perubahan campuran sesudahnya.
6. Tulis seluruh 7 file CSV (RFC 4180, UTF-8 tanpa BOM) dan `report.json` ke folder `docs/database/migrations/` (atau `-out`), serta root Markdown, procedure index/catalog, dan generated procedure profiles ke folder `docs/database/` (atau `-docs`). Summary CSV/JSON menyimpan explicit exclusion count; link profile di-escape sebagai URL path sehingga karakter khusus tetap membuka file yang benar. Lineage menghapus dependency identik case-insensitive, dan `EXEC @return_code = Procedure` tidak diklasifikasikan sebagai dynamic SQL.
   Output path yang sama dengan/berada di bawah `migrations/` atau `db_object/` ditolak sebelum write.
7. Cetak banner ringkasan ke console, dan tunggu tombol Enter jika dijalankan interaktif/double-click.

## 6. Sinkronisasi `rza_get_object_db` dan `rza_migrate_db`

Sinkronisasi kontrak yang sudah aktif:
- Object token charset: `A-Za-z0-9_.#`.
- Schema hanya `dbo`.
- Programmability source wajib `CREATE OR ALTER`.
- Index source mendukung:
  - clustered/nonclustered
  - unique
  - include
  - filtered predicate `WHERE ...`
- Satu file `indexes/dbo.<Table>.sql` boleh berisi banyak `CREATE INDEX` untuk table yang sama.

## 6.1 Sinkronisasi dengan `rza_migrate_rest`

- `rza_migrate_rest` mengeksekusi file di folder `migrations/` yang sama dengan hasil generator `rza_migrate_db`.
- Endpoint `run-file` hanya mengizinkan:
  - file `.up.sql` yang merupakan next version dari state current
  - file `.down.sql` yang tepat sama dengan current version
- Dengan kontrak ini, mode execute tetap konsisten dengan deterministic history yang dibangun engine.

## 7. Guard Penting di Verify

- Pairing up/down per version.
- Kontrak core audit init.
- Wrapper contract tables/indexes/patch.
- Forbidden SQL token.
- Patch immutability + source integrity.
- Deterministic migration ordering (runtime): validasi per chain `run timestamp + object + vk`.

## 8. Artifact Output

### `out/output.json`
- `fail`
- `pass`
- `new_migrations`
- `notes`

Saat `fail.count > 0`, isi setiap `fail.items[]` juga ditampilkan pada console dengan nomor item, rule, severity, file, object/location, reason/action, message, dan suggestion yang tersedia. `out/output.json` tetap berguna sebagai artifact terstruktur, tetapi bukan satu-satunya sumber diagnosis pipeline.

### `out/get_object_summary.json`
- metadata run extractor (`export`/`dry-run`/`preview`)
- count create/update/delete/unchanged
- warning modul yang di-skip (jika ada)

### `docs/database/migrations/*` (`rza_migrate_report` audit reports)
- `report_summary.csv`: ringkasan global status run, jumlah file, analyzed pair, excluded pair, objek, versi, operasi, warning, error, dan status.
- `migration_summary.csv`: ringkasan tiap pasangan migration, jumlah operasi, jenis operasi, dan status round-trip.
- `object_history.csv`: riwayat tiap objek dan versi (`INITIAL`, `UPDATED`, `DELETED`, `REINTRODUCED`).
- `table_changes.csv`: rincian operasi semantik table (kolom, nullability, default, constraint, tipe data).
- `index_changes.csv`: rincian operasi semantik index (tipe, kolom key, include, filter, unique, clustered).
- `programmable_changes.csv`: rincian operasi views, functions, procedures (hash before/after, changed sections).
- `type_changes.csv`: rincian operasi user-defined types.
- `report.json`: ringkasan metadata proses, menegaskan `database_access: false` dan `sql_execution: false`.

### `docs/database/*` (`rza_migrate_report` desired-state docs)
- `identifiers/table_unique_ids.md`: inventory uniqueness enforcement, filtered uniqueness, `IDENTITY`, dan ID review signals. JOIN/non-unique index/nama tidak dianggap bukti uniqueness.
- Section `Suspected JOIN identifiers` menampilkan kolom tanpa index/IDENTITY, termasuk generic name seperti Code dan reference column pada table dengan PK lain. Evidence memuat canonical mapping dari observed equality JOIN, source view/procedure, usage, index/key, IDENTITY, serta composite-key context. FK-only tidak dihitung sebagai observed JOIN. Fitur ini hanya output report: extractor tetap catalog-based, checker apply/`-check` tidak menambah correction/migration, generator tetap memakai delta schema, pipeline tetap read-only; report tersedia lewat full release binaries.
- `relationships/relational_erd.md`: seluruh Mermaid dalam satu file, disusun per Relationship Group dan connected Relational Map. Setiap map dibentuk adaptif dengan maksimum 20 object, 24 relationship lines, 35 exact mappings, dan 9.000 karakter Mermaid. Node menunjukkan type serta key/index context bagi kolom relationship; solid line untuk pure declared FK, dashed undirected line untuk observed JOIN, dan setiap label mempertahankan exact column mapping.
- `relationships/relationship_catalog.md`: exact object-column mapping, evidence, jenis JOIN, occurrence, source module, dan confidence.
- `relationships/module_lineage.md`: static dependency view/procedure ke data object dan dynamic-SQL manual-review list; bukan klaim relasi antartable.
- `procedures/README.md`: entry point dan index seluruh procedure.
- `procedures/procedure_impact_catalog.md`: ringkasan direct reads/writes, calls, functions, views, dynamic SQL, dan unresolved references.
- `procedures/profiles/dbo.<Procedure>.md`: profile source-backed untuk contract, branch-aware local execution flow, local object map, exact writes, key context, guarded direct calls, recursive procedure map/write impact, serta dynamic SQL. Semua object divalidasi ke `db_object`; runtime target tidak ditebak.
- `README.md` dan `manifest.json`: navigasi utama dan inventory hash; legacy tree pointer tidak lagi dibuat.

Observed JOIN scanner membuka alias scope pada setiap `FROM`, membaca explicit `JOIN ... ON` dan legacy comma-join `WHERE`, hanya menerima equality dengan satu resolved `alias.column` per sisi, dan menolak operator non-equality, ekspresi ambigu, temp/table variable, atau dynamic SQL. Module tanpa resolved static reference dicantumkan untuk audit coverage. Karena capability ini report-only, extractor, checker apply/`-check`, generator, dan GitLab checker tidak mengubah source maupun migration; binary report tersedia setelah full release build.

Markdown desired-state memakai narasi Bahasa Indonesia yang natural dan conventional English database terms. Guidance classification berada pada reference section, bukan kolom template per row. Relationship Group berarti connected component, bukan urutan eksekusi. Seluruh diagram tetap di satu file; setiap Relational Map tetap connected serta dibatasi 20 object, 24 actual relationship lines, 35 exact mappings, dan 9.000 karakter Mermaid tanpa artificial Hub/parent. Node memuat object type dan key/index context bagi relationship columns. Setiap non-self object pair tampil dengan seluruh exact mappings; `[FK]`/`[JOIN]` hanya diperlukan pada mixed-evidence pair. Dense pair memakai numbered `Mapping batch` untuk menjaga text-size budget. Object index dan shared object navigation menghubungkan map yang mengulang object; occurrence tetap tersedia di catalog dengan Group/Map yang sama. Self-join hanya tampil pada review section. Daftar module menggunakan titik koma dan tidak memakai raw HTML `<br>`. Flow ini report-only: extractor tidak mengubah output, checker apply dan `-check` tidak merencanakan correction, generator tidak menghasilkan migration, pipeline authoring tetap read-only, dan binary report mendapat behavior yang sama setelah full release build.

Procedure flow membaca source `db_object/procedures` lalu lookup target ke inventory `db_object/tables|views|functions|procedures`. `RESOLVED_STATIC` dan `RESOLVED_DYNAMIC` hanya diberikan setelah catalog match; caller literal dicatat sebagai `RESOLVED_FROM_CALLER` untuk call site tersebut. Runtime variable/configuration tetap `UNRESOLVED_DYNAMIC`. Direct access/call detail selalu menyimpan guard, arguments/evidence, dan source line. Local execution flow memakai orientasi top-down serta merangkum transaction, condition/read checks, procedure routing, rejoin, dan error transfer yang terlihat pada source; full source sequence tetap tersedia di tabel. Logical line break pada node/edge dikonversi menjadi line break Mermaid sehingga tidak tampil sebagai teks `\n`. Local object map menyatakan jenis akses dan tidak boleh dibaca sebagai urutan eksekusi atau foreign key. Recursive section menampilkan arah caller → callee, guard, cycle, already-mapped branch, serta write impact child. Diagram dengan lebih dari 20 node dipecah menjadi beberapa map tanpa menghilangkan call edge; tabel audit tidak dipotong.

Access scan memakai token yang membawa source position. Keyword `FROM`/`JOIN`/DML dan object boleh berada pada baris berbeda tanpa kehilangan identity. Scope CTE berakhir bersama statement-nya, sehingga `WITH T ...; SELECT ... FROM dbo.T` tetap mengenali `dbo.T` sebagai persisted table. Alias update/delete hanya dihubungkan ke source pada statement yang sama. Result contract mengecualikan variable assignment, `SELECT INTO`, subquery/`EXISTS`, dan `INSERT ... SELECT` dari client result set.

Refinement 2026-10-01 membatasi alias juga pada query scope, mempertahankan quoted identity serta setiap occurrence di satu baris, dan membaca TOP/MERGE USING/APPLY/comma sources, SELECT/OUTPUT sink, serta static EXEC multiline dari scanner shared. OUTPUT tanpa sink menghasilkan client result; cursor SELECT dan compound query branch tidak menghasilkan signal tambahan. Write mapping memakai statement offsets, bukan window 80/160 baris. Snapshot/collation/runtime ambiguity tetap terlihat. Extractor catalog tidak berubah; checker apply dan `-check` tidak mendapat fixer baru atau migration baru. Generator memakai dependency parser yang sama untuk ordering create/removal tanpa mengubah semantic SQL; pipeline tetap read-only dan membawa binary baru sesudah build lengkap.

Default output tetap `out/report` dan `docs/database`. Override relatif `-out`/`-docs` dihitung dari repository root yang ditemukan, bukan current working directory. Setelah dokumentasi selesai ditulis, `report.json` dipublikasikan dengan `output_directory`, `documentation_directory`, dan `artifacts[]` konkret. `reports[]` lama dipertahankan sebagai compatibility inventory.

Ordering dan dataset report deterministik untuk source yang sama. Kolom `run_at` dan field `generated_at` tetap merekam waktu run aktual dan dikecualikan dari perbandingan byte-identik.

## 9. Residual Risk Aktif

- Backend AST ANTLR belum diaktifkan penuh; saat ini formatter berjalan stabil dengan backend `lite` (opsi `antlr` tetap fallback ke `lite`).

## 10. Command Validasi Rekomendasi

- `go test ./...`
- `powershell -ExecutionPolicy Bypass -File scripts/validate_qa_fixture_contract.ps1`
- `powershell -ExecutionPolicy Bypass -File scripts/verify_rules_2x_learning.ps1`
- `powershell -ExecutionPolicy Bypass -File scripts/validate_learning_examples_hygiene.ps1`
- `powershell -ExecutionPolicy Bypass -File scripts/validate_migration_lock_contract.ps1`
- `powershell -ExecutionPolicy Bypass -File scripts/validate_runner_stop_behavior_contract.ps1`

## 11. Catatan Distribusi

- Lokasi standar file authoring Windows dan lisensi dipublikasikan ke `tools/`:
  - `tools/rza_migrate_db.exe`
  - `tools/rza_migrate_check.exe`
  - `tools/rza_migrate_report.exe`
  - `tools/license.key`
  - `tools/linux-amd64/`
- Nama executable di-lock (baik di root maupun di folder `tools/`):
  - `rza_get_object_db.exe`
  - `rza_migrate_check.exe`
  - `rza_migrate_db.exe`
  - `rza_migrate_report.exe`
  - `rza_migrate_rest.exe`
- Jika file di-rename, aplikasi akan fail saat start.

## Report layout V2

Report layout V2 workflow: jalankan rza_migrate_report, buka docs/database/README.md, lalu identifiers/relationships/procedures atau migrations/README.md. Default report tidak menulis out/. -docs memindahkan seluruh root; -out opsional memindahkan audit dengan datasets child. File manual tetap aman; checksum collision memblokir sebelum write. Extractor → checker apply/CI → generator → runner tetap seperti sebelumnya, dengan released report binary yang dibangun ulang. Lihat [layout walkthrough](report_output_layout.md).


## Audit object precision (2026-10-12)

Extractor tetap menghasilkan canonical source. Report decode source/history di memory, memeriksa declared identity, membangun DDL catalog dan scoped references, lalu mempublikasikan evidence/resolution termasuk function dan missing columns. Interactive checker/-check serta generator tidak memperoleh safe correction atau schema action baru. Shared MERGE JOIN lexer berlaku pada dependency consumers; pipeline flags dan history immutability tetap sama. Complete release membawa perilaku yang sama ke Windows/Linux.

Detail bukti: review_report_object_precision_20261012.md.
