# Specification Baseline - rza_migrate_db Engine Suite

2026-10-12 report requirement: `table_unique_ids.md` menambahkan Suspected JOIN identifiers dari static observed equality JOIN, tanpa index/IDENTITY atau nama ID-like prerequisite dan tetap tampil pada table ber-PK lain. Evidence harus memuat canonical column mapping, source modules, usage, serta index/key dan IDENTITY. `OBSERVED_JOIN_COLUMN_SIGNAL` membedakan generic JOIN operand dari `OBSERVED_JOIN_ID_SIGNAL`; bukan unique row-ID claim. Single-column unfiltered unique key tidak diduplikasi; filtered/composite status dijelaskan. FK-only/inequality/runtime/ambiguous relations tidak menjadi JOIN evidence. Fitur report-only; extractor/checker apply/CI/generator tidak mendapat schema action atau migration baru; full release binary parity tetap wajib. QA-31 golden serta report rerun/no-write test menjadi acceptance.


Object-recognition refinement (2026-10-01): scanner shared wajib mempertahankan quoted/escaped identity, menangani malformed delimiter tanpa panic, DML TOP, MERGE target/USING tanpa phantom action, APPLY/comma sources, SELECT/OUTPUT INTO writes, dan static EXEC multiline/repeated. UPDATE/DELETE alias hanya dihubungkan pada statement serta query scope yang sama; collision tidak ditebak. Client-result classification mengecualikan cursor SELECT dan cabang compound INSERT/SELECT INTO serta mengenali OUTPUT tanpa sink. Mapping memakai source span statement dan lexical top-level clause; multirow atau TOP/DISTINCT/compound INSERT projection tetap tanpa exact mapping. QA-31 golden, precision regression, fuzz bounds, dan checker/generator boundary test menjadi acceptance. Extractor tidak mendapat object kind baru; interactive/CI checker tidak mendapat correction baru; generator hanya mendapat dependency recognition yang lebih tepat tanpa mengganti semantic SQL/up-down. Full Windows/Linux build dan GitLab checker smoke wajib setelah source final. Detail desain/batas di [review object scan](docs/review_object_scan_20261001.md).

Dokumen ini adalah baseline requirement pengembangan engine. Dokumen ini bersifat normatif dan menjadi acuan implementasi, review, dan acceptance.

Untuk lookup per aplikasi, lihat [engine reference](docs/engine_reference.md): input/output,
mode, efek samping, CLI/config, dan failure boundary. [Architecture overview](docs/architecture/overview.md)
menjelaskan desired/history/runtime state; [call tree](docs/architecture/call_tree.md)
menautkan requirement ke flow source. Panduan tambahan ini tidak mengubah requirement atau status capability.

## 1. Target Sistem

- DB target: Microsoft SQL Server.
- Schema target: tunggal `dbo`.
- Runner deployment: `golang-migrate`.
- Extractor: `rza_get_object_db` menarik snapshot source-of-truth ke `db_object/`.
- Generator/Verifier: `rza_migrate_db` berjalan offline untuk `fmt`, `verify`, `gen`.
- Sumber waktu audit SQL: `GETDATE()` (waktu lokal server SQL), bukan UTC.

## 2. Definisi Kunci

- Source-of-truth: `db_object/`.
- History: `migrations/`.
- Patch: perubahan one-time di `patch/<folder>/` yang digenerate menjadi migration.
- Semantic diff: perubahan bermakna pada object schema, bukan whitespace-only.
- State reconstruction: pembacaan history migration untuk menyusun state terakhir per object tanpa koneksi DB.
- Extractor summary: laporan run `rza_get_object_db` pada `out/get_object_summary.json`.

## 3. Hard Locks Global

1. Jika ada severity `3` atau `out/output.json.fail.count > 0`, engine wajib fail hard.
2. Fail hard berarti exit code `1`, stop proses, dan tidak boleh ada migration baru yang tertulis sebagai hasil final run.
3. Semua parsing filename/grammar utama wajib structured, tanpa regex sebagai parser utama.
4. `migrations/` wajib flat, tanpa subfolder.
5. Semua migration wajib berpasangan up/down pada version yang sama persis.
6. Semua perubahan schema yang dikelola engine wajib melalui repo; out-of-band DDL dianggap drift.

## 4. Kontrak Struktur Repo

## 4.1 `db_object/`

Struktur wajib:

- `db_object/types/dbo.<TableType>.sql`
- `db_object/tables/dbo.<Table>.sql`
- `db_object/indexes/dbo.<Table>.sql`
- `db_object/views/dbo.<View>.sql`
- `db_object/functions/dbo.<Function>.sql`
- `db_object/procedures/dbo.<Procedure>.sql`

Aturan wajib:

- Schema hanya `dbo`.
- Nama file harus sama dengan object yang didefinisikan.
- `indexes/dbo.<Table>.sql` hanya boleh berisi index untuk table tersebut.

Mode transisi layout:

- Engine masih membaca layout legacy `db_object/schema/<category>/` untuk kompatibilitas.
- Jika layout legacy terdeteksi, engine emit warning `LEGACY_DB_OBJECT_LAYOUT_DETECTED`.
- Jika object yang sama ada di canonical+legacy dengan isi berbeda, engine fail `DB_OBJECT_DUAL_PATH_CONFLICT`.
- Jika object yang sama ada di canonical+legacy dengan isi sama, engine emit warning `LEGACY_DB_OBJECT_DUPLICATE_CANONICAL`.

## 4.2 `patch/`

Struktur wajib per folder patch:

- `README.md`
- `up.sql`
- `down.sql`

Aturan wajib:

- Folder patch hanya boleh berisi tiga file tersebut.
- Patch source immutable setelah patch migration pertama dibuat.

## 4.3 `migrations/`

Aturan wajib:

- Semua file langsung berada di `migrations/`.
- Core audit init wajib ada:
  - `000000__core__mig__audit_init.up.sql`
  - `000000__core__mig__audit_init.down.sql`

Naming wajib:

- Schema migration: `{version}_{category}_{object}_v{k}.(up|down).sql`
- Patch migration: `{version}_patch_{patch_name}_v{k}.(up|down).sql`

## 5. Tooling dan Runtime Pipeline

- Pipeline lokal standar:
  - `rza_get_object_db` untuk refresh desired state (`db_object/`) dari SQL Server.
  - `rza_migrate_db` untuk `fmt/verify/gen` migration berdasarkan desired state.
- Runtime eksekusi migration: `golang-migrate`.
- Job migration hanya mengeksekusi isi `migrations/`.
- Jika satu file migration error (THROW), runner berhenti dan file setelahnya tidak dieksekusi.

## 6. Larangan dan Safety Guard

Token atau statement berikut dilarang di `migrations/`, `db_object/`, `patch/`:

- `GO`
- `USE <db>`
- `CREATE DATABASE`
- `ALTER DATABASE`
- `CREATE LOGIN`
- `ALTER LOGIN`
- `CREATE USER`

Jika ditemukan, wajib fail.

## 7. Pairing, Version, dan Filename Grammar

## 7.1 Pairing

- Setiap file `.up.sql` wajib punya pasangan `.down.sql` dengan version identik.
- Missing pair adalah fail hard.

## 7.2 Version format

- `base_ts = YYYYMMDDHHMMSS`
- `pair_seq = 000..999`
- `version = base_ts + pair_seq` (17 digit)
- Jika satu run menghasilkan lebih dari 999 pair, fail hard.

## 7.3 Grammar filename

Schema migration:

- `<version>_<category>_<object>_v<k>.(up|down).sql`
- category: `types|tables|indexes|functions|views|procedures`
- object token: `A-Z a-z 0-9 _ . #`
- `k >= 1`

Patch migration:

- `<version>_patch_<patch_name>_v<k>.(up|down).sql`

Jika parse gagal, fail.

## 8. Ordering, Version `v{k}`, dan Reconstruction

## 8.1 Ordering emit

Urutan emit per run wajib:

1. types
2. tables
3. indexes
4. functions
5. views
6. procedures
7. patch

Aturan tambahan:

- `pair_seq` naik sesuai urutan emit.
- Dalam kategori yang sama: sort lexicographic object (case-sensitive sesuai filename), lalu `v{k}`.
- Dilarang iterasi map tanpa sort saat emit.
- Banyak perubahan pada object yang sama dalam 1 run harus digabung jadi 1 pair kecuali policy meminta pemisahan.

## 8.2 Aturan naik `v{k}`

- `k` hanya naik jika ada semantic diff setelah format canonical.
- Whitespace-only change tidak boleh menaikkan `k`.
- `k` dihitung dari history reconstruction, bukan dari timestamp.
- `k` patch dihitung per `patch_name`.
- Untuk programmability (`views/functions/procedures`), perubahan style-only bracket identifier dan terminator `;` diperlakukan no-op (tidak menaikkan `k`).

## 8.3 State reconstruction

- Engine membaca seluruh `migrations/*.up.sql` dengan parser structured subset MVP.
- Jika ada file history tidak bisa diparse, fail hard.
- Konflik history key (`version+category+object+v{k}`) wajib fail.
- Output minimal reconstruction di notes:
  - `last_seen_version`
  - `last_seen_vk`
  - `normalized_hash`

## 9. Kontrak Validasi dan Format `db_object/`

## 9.1 Auto-format

Sebelum parse/diff, engine wajib:

- normalize line ending ke LF
- trim trailing spaces
- ensure final newline
- preserve keyword casing
- tidak mengubah string literal dan komentar

## 9.2 Canonical table snapshot

- Gaya koma: trailing comma, bukan leading comma.
- Isi file table snapshot harus representasi desired `CREATE TABLE` + named constraints.

## 9.3 Larangan table snapshot

Dilarang di `db_object/tables/*.sql`:

- wrapper `IF ... BEGIN ... END`
- `DROP TABLE`
- `ALTER TABLE`
- `CREATE INDEX`
- opsi advanced `WITH (...)` seperti filegroup/partition/temporal

## 9.4 Header programmability

Wajib:

- `CREATE OR ALTER PROCEDURE dbo.X`
- `CREATE OR ALTER FUNCTION dbo.X`
- `CREATE OR ALTER VIEW dbo.X`

## 9.5 Konsistensi file dan isi

- Nama file dan object definition harus sama.
- File index per table tidak boleh mencampur table lain.
- File programmability wajib match `dbo.<Name>` file.

## 9.6 Kontrak index snapshot

- Hanya `CREATE INDEX` atau `CREATE UNIQUE INDEX` untuk table target.
- `DROP INDEX` dilarang di data master.
- Nama index harus deterministik.

## 9.7 Inline default dilarang

- Inline `DEFAULT(...)` pada definisi kolom table snapshot adalah fail (`INLINE_DEFAULT_NOT_ALLOWED`).
- Default harus berupa named constraint seperti `CONSTRAINT DF_Table__col DEFAULT(...) FOR col`.

## 9.8 Nullability omission

- Kolom tanpa deklarasi `NULL/NOT NULL` diperlakukan sebagai `NULL`.

## 9.9 Whitespace-only

- Jika setelah format tidak ada semantic diff, tidak boleh ada migration baru dan `v{k}` tidak naik.

## 10. Severity, Decision, Action, Reason

## 10.1 Severity map (fixed)

- `0 = INFO`
- `1 = NOTICE`
- `2 = WARNING`
- `3 = ERROR`

## 10.2 Decision whitelist

- `run`
- `ok`
- `skip`
- `warn`
- `fail`

Catatan: `warn` adalah non-fatal event, bukan pengganti `fail`.

## 10.3 Action whitelist

- `create_table`
- `drop_table`
- `add_column`
- `drop_column`
- `alter_column`
- `backfill`
- `add_default`
- `drop_default`
- `add_check`
- `drop_check`
- `add_unique`
- `drop_unique`
- `add_index`
- `drop_index`
- `add_fk`
- `drop_fk`
- `create_or_alter_view`
- `create_or_alter_function`
- `create_or_alter_procedure`
- `followup_required`
- `migration_failed`

## 10.4 Reason whitelist

- `TABLE_EXISTS`
- `TABLE_NOT_EXISTS`
- `COLUMN_EXISTS`
- `COLUMN_NOT_EXISTS`
- `INDEX_EXISTS`
- `INDEX_NOT_EXISTS`
- `DF_EXISTS`
- `DF_NOT_EXISTS`
- `FK_EXISTS`
- `FK_NOT_EXISTS`
- `CK_EXISTS`
- `CK_NOT_EXISTS`
- `UQ_EXISTS`
- `UQ_NOT_EXISTS`
- `OBJECT_EXISTS`
- `OBJECT_NOT_EXISTS`
- `CREATED`
- `DROPPED`
- `ALTERED`
- `BACKFILLED`
- `REPLACED`
- `DROPPED_IMPACTED_INDEX`
- `DROPPED_DEPENDENCY`
- `INDEX_RECREATE_REQUIRED`
- `REQUIRES_MANUAL`
- `NARROWING_WOULD_TRUNCATE`
- `FK_REFERENCED_BY_OTHERS`
- `DATA_RESTORE_NOT_GUARANTEED`
- `DRIFT_DETECTED`
- `INLINE_DEFAULT_NOT_ALLOWED`
- `SQL_ERROR`

## 10.5 Skip vs stop

- `skip` hanya untuk policy precheck atau target sudah sesuai.
- `skip` bukan recovery error.
- SQL runtime error pada tables/indexes/patch wajib fail hard lalu THROW.

## 11. Policy Destructive

- DROP TABLE: manual-only. Engine tidak auto-emit DROP TABLE.
- Jika table dihapus dari desired state: migration harus `skip + warn` severity=2 dengan reason `REQUIRES_MANUAL`.
- Down untuk kasus destruktif harus warning `DATA_RESTORE_NOT_GUARANTEED`.
- DROP COLUMN: boleh otomatis dengan syarat drop dependency konflik (DF/CK/UQ/FK/index terkait), tidak auto-recreate index, dan warning `INDEX_RECREATE_REQUIRED`.
- DROP INDEX: boleh bila hilang dari desired state, wajib precheck existence, jika sudah tidak ada maka `skip` INFO.

## 12. Transaction dan Error Strategy

- Kategori `tables`, `indexes`, `patch` wajib wrapper TRY/TRAN/COMMIT/CATCH/ROLLBACK/THROW.
- Urutan saat error: rollback, log fail, lalu THROW.
- `views/functions/procedures` tidak wajib wrapper, tetapi wajib `CREATE OR ALTER`.
- `THROW` adalah mekanisme stop default untuk golang-migrate.
- Multi-step operation wajib deterministik, tercatat audit, dan atomic.

## 13. Core Audit Init dan Audit Logging

## 13.1 Core audit init migration

Wajib ada file awal:

- `000000__core__mig__audit_init.up.sql`
- `000000__core__mig__audit_init.down.sql`

Up wajib minimal membuat:

- `dbo.mig_event`
- `dbo.mig_operation_journal`
- `dbo.usp_mig_log_event`
- `dbo.usp_mig_try_get_lock`

Aturan:

- tanpa `GO`
- tanpa `USE`
- tolerant rerun dengan precheck `OBJECT_ID`

## 13.2 Audit schema minimum

`dbo.mig_event` minimal berisi:

- `event_id`
- `run_key`
- `operation_id`
- `version`
- `category`
- `file_name`
- `direction`
- `content_sha256`
- `decision`
- `severity`
- `action`
- `target`
- `reason`
- `message`
- `detail`
- `started_at`
- `duration_ms`
- `created_at DEFAULT SYSUTCDATETIME()`

- `error_number`
- `error_line`
- `error_procedure`

`dbo.mig_operation_journal` wajib menyimpan minimal request/operation id,
before/after version, status, requested/executed steps, file report JSON,
manifest SHA-256, error, dan timestamp UTC.

## 13.3 Aturan logging

- `run_key` wajib `NEWID()` pada setiap file migration.
- Wrapper wajib mendeklarasikan SHA-256 body dan marker body yang dapat diverifikasi.
- Metadata audit tambahan dikirim melalui `SESSION_CONTEXT` agar call logger tetap kompatibel dengan signature lama.
- Tidak boleh ada kolom audit redundant seperti `env/db_name/tool_*`.
- Header severity mapping wajib disisipkan pada seluruh wrapper generated.

## 14. Kontrak Wrapper Migration

Untuk tables/indexes/views/functions/procedures:

- metadata `@version`, `@category`, `@file`, dan `@target` wajib deterministik serta exact-case sesuai filename/object. Target index boleh berbentuk `<table>.<index>` dengan prefix table exact-case.
- precheck dan skip boleh sebelum transaksi.
- audit event minimal mencatat run/ok/fail (atau skip saat precheck).
- pada CATCH harus log `migration_failed` dengan reason `SQL_ERROR` lalu THROW.

## 15. Kontrak Down Migration

- `.down.sql` wajib mengembalikan schema ke versi sebelumnya.
- Migration pertama untuk `views/functions/procedures` wajib mempunyai inverse executable `DROP VIEW|FUNCTION|PROCEDURE IF EXISTS`; down berupa komentar saja dilarang untuk output generator baru.
- Untuk operasi destruktif, wajib warning `DATA_RESTORE_NOT_GUARANTEED`.
- Down tables/indexes/patch wajib wrapper transaksi seperti up.
- Precheck down boleh skip jika object sudah sesuai, tetapi runtime SQL error tetap fail + THROW.

## 16. Patch Policy

## 16.1 Canonicalize `patch_name`

Algoritma wajib:

1. trim whitespace
2. whitespace menjadi `_`
3. lowercase
4. karakter di luar `[a-z0-9_-]` menjadi `_`
5. collapse underscore berulang
6. trim `_` dan `-` di ujung
7. hasil kosong -> fail

## 16.2 Generation

- Emit pair: `{version}_patch_<patch_name>_v{k}.up.sql` dan `.down.sql`.
- Header patch wajib berisi `patch_folder`, `patch_name`, dan `readme`.
- Body patch disisipkan as-is, kecuali normalisasi LF dan trailing spaces.
- Wrapper audit patch ditambahkan otomatis.

## 16.3 Duplication dan integrity

- Folder patch wajib mempunyai tepat satu logical `README.md`, `up.sql`, dan `down.sql`, tanpa file/directory tambahan. Resolution nama file case-insensitive dan harus membuka nama entry aktual dari directory agar konsisten pada filesystem Windows maupun Linux; case-only duplicate bersifat ambigu dan ditolak.
- Jika patch migration dengan `patch_name` yang sama sudah ada, engine skip generate tanpa compare isi.
- Regenerate patch hanya boleh setelah migration patch lama dihapus manual.
- Jika migration patch sudah ada tetapi source patch hilang/tidak lengkap, fail hard.
- Output pass item wajib menyimpan mapping:
  - `patch_folder_original`
  - `patch_name_canonical`
  - `generated_migration_versions[]`

## 17. Kontrak `out/output.json`

- `out/output.json` bukan source-of-truth dan tidak boleh di-commit.
- CLI wajib mencetak seluruh fail item ke console secara deterministik dan menyertakan field diagnosis yang tersedia (`rule`, `severity`, `file`, `object`, `location`, `reason`, `action`, `message`, `suggestion`). Ringkasan hasil dan detail failure memakai satu output stream agar urutannya tidak rusak ketika GitLab menggabungkan stdout/stderr. Pipeline user harus dapat mendiagnosis failure tanpa membuka artifact JSON.
- Struktur root wajib:

```json
{
  "fail": {"count": 0, "items": []},
  "pass": {"count": 0, "items": []},
  "new_migrations": {
    "naming": {
      "pattern": "YYYYMMDDHHMMSS<seq>_<category>_<object>_v<k>.(up|down).sql",
      "line_endings": "LF"
    },
    "files": []
  },
  "notes": {}
}
```

Item minimal:

- `rule`
- `severity`
- `file`
- `message`

Item recommended:

- `object`
- `location`
- `suggestion`
- `reason`
- `action`

## 18. Locking Runtime

- Harus ada maksimal 1 run migration per DB.
- Preferred strategy: `sp_getapplock` melalui `dbo.usp_mig_try_get_lock`.
- Jika lock gagal, proses wajib fail + THROW.

## 19. MVP Supported Subset

Jika di luar subset ini, engine fail.

Tables supported:

- scalar types dasar: int/bigint/smallint/tinyint/bit
- decimal/numeric, money/smallmoney, float/real
- char/varchar/nchar/nvarchar
- date/datetime/datetime2/smalldatetime
- uniqueidentifier, varbinary, xml
- named constraints PK/UQ/DF/CK/FK
- computed columns in the form `<name> AS <expression> [PERSISTED]`

Tables not supported:

- sparse
- temporal
- partition/filegroup directives
- advanced features lain

Indexes supported:

- `CREATE [UNIQUE] INDEX ... ON dbo.Table(...)`
- optional `INCLUDE (...)`
- multiple statements per file index table

Indexes not supported:

- columnstore
- XML index
- spatial index
- indexed view
- data compression directives

Filtered index (`WHERE ...`) termasuk supported subset.

Programmability supported:

- snapshot text untuk view/function/procedure
- header wajib `CREATE OR ALTER`
- 1 file programmability hanya 1 object
- function scalar (`FN`), inline TVF (`IF`), dan multi-statement TVF (`TF`)
- parameter default, `SCHEMABINDING`, dan table return yang valid
- definition encrypted tidak boleh ditebak; extractor mempertahankan snapshot lama dan melaporkan warning

Table type/TVP supported:

- snapshot `CREATE TYPE dbo.<Name> AS TABLE (...)` di `db_object/types/`
- column, default, check, primary key, unique, dan inline rowstore index
- constraint TVP memakai grammar unnamed SQL Server; bentuk `CONSTRAINT <name>` ditolak
- perubahan struktur menjatuhkan dependent function/procedure, mengganti type, lalu membuat ulang dependent module di dalam transaksi

## 20. ALTER COLUMN Precheck Map

Jika unsafe, generator wajib `skip` severity=2 dengan suggestion manual.

Kasus wajib:

- parser/hash/diff wajib membaca seluruh parameter tipe walaupun ada spasi
  sebelum `(` atau sesudah `,`; whitespace-only tidak menghasilkan migration,
  sedangkan perubahan precision, scale, length, atau date-time precision wajib
  menghasilkan delta kolom.
- delimiter tipe bawaan hasil extractor (`[DECIMAL]`, `[VARCHAR]`, dan tipe
  bawaan lain) wajib dinormalisasi ke base type yang sama dengan bentuk tanpa
  delimiter. Perubahan delimiter saja adalah no-op; perubahan parameter tetap
  menghasilkan guarded migration bila transisinya dapat dibuktikan lossless.
- `NULL -> NOT NULL`: blok jika masih ada data NULL.
- shrink `VARCHAR/CHAR`: blok jika panjang data melebihi target.
- shrink `NVARCHAR/NCHAR`: blok jika `DATALENGTH > new_len*2`.
- type conversion ke numeric: blok jika `TRY_CONVERT` gagal.
- bila alter membutuhkan drop dependency untuk lolos, di MVP tidak boleh auto-drop dependency; wajib skip dengan reason `REQUIRES_MANUAL` dan detail dependency penghalang.
- jika dependency penghalang unnamed sehingga rencana drop/recreate tidak deterministik, engine wajib fail lebih awal lewat source validation (`UNNAMED_CONSTRAINT_NOT_ALLOWED`) sebelum masuk ke planning alter/drop column.

## 21. Dokumen dan Hygiene Repo Opsional

- `constraints_naming.md` adalah guideline opsional; bootstrap boleh membuat template tanpa menimpa file yang sudah ada.
- Semua constraint pada `db_object/tables/*.sql` wajib diberi nama eksplisit dengan `CONSTRAINT <name>`.
- Constraint unnamed yang dapat dibuktikan dari struktur `CREATE TABLE` diberi nama deterministik oleh safe-fix. Bentuk ambigu wajib fail hard saat verify dengan rule `UNNAMED_CONSTRAINT_NOT_ALLOWED`.
- Prefix constraint tidak diwajibkan; nama eksplisit legacy tetap diterima.
- `.gitignore` direkomendasikan mengandung:
  - `out/**`
  - `*.output.json`
- `.gitattributes` direkomendasikan mengandung normalisasi LF minimal:
  - `* text=auto eol=lf`
  - `*.sql text eol=lf`
  - `*.md text eol=lf`
  - `*.go text eol=lf`

## 22. CLI Contract

Command minimum:

- `rza_migrate_check`
- `rza_migrate_db fmt`
- `rza_migrate_db verify`
- `rza_migrate_db gen`
- `rza_get_object_db`
- `rza_migrate_rest -config <path>`
- `rza_migrate_report [-repo <path>] [-out <path>] [-docs <path>] [-pause-on-exit true|false]`

Perintah:

- `rza_migrate_report`: rekonstruksi histori perubahan berbasis `migrations/`, serta dokumentasi desired-state read-only dari `db_object/`; zero DB connection, zero SQL execution, zero write/modify ke source. Menghasilkan 7 CSV (RFC 4180, UTF-8 tanpa BOM) dan `report.json` di `docs/database/migrations/`, serta root README/manifest, empat grouped Markdown, procedure index/catalog, dan generated procedure profiles di `docs/database/`.
- `rza_migrate_check`: workflow developer `fmt db_object -> verify authoring -> out/output.json`; tidak pernah menjalankan generator dan wajib mempertahankan byte seluruh file `migrations/`.
- `rza_migrate_check` tanpa argumen: mencari repo dari working directory/folder executable dan pause saat exit untuk penggunaan double-click.
- `rza_migrate_check` dengan argumen: non-interaktif secara default; menerima `-repo`, `-check`, `-profile authoring|full`, `-migration-validation off|integrity|full`, `-perf-notes`, `-fail-fast`, `-parallel-read`, `-parallel-workers`, dan `-pause-on-exit`.
- `rza_migrate_check -check`: gate read-only untuk CI; formatting drift menghasilkan `FMT_CHECK_CHANGES_REQUIRED` tanpa menulis `db_object`.
- `fmt`: auto-format aman `db_object/**` dan tulis `out/output.json`; comment layout dan multiline literal merupakan protected spans.
- `fmt -check`: preview daftar file yang perlu diformat tanpa menulis `db_object`; perubahan menghasilkan `FMT_CHECK_CHANGES_REQUIRED`.
- `verify`: validate only, no emit, dan tulis `out/output.json`.
- `gen`: `fmt -> validate -> reconstruct history -> diff -> emit -> verify -> out/output.json`.
- `rza_migrate_db` tanpa command: default `verify` lalu `gen`.
- `rza_migrate_db -gen-auto-indent true|false`: override auto-indent saat `gen` (flag > env `RZA_GEN_AUTO_INDENT` > default true).
- `rza_migrate_db -perf-notes true|false`: menulis telemetry durasi fase ke `out/output.json.notes.perf_ms` untuk baseline performance (flag > env `RZA_PERF_NOTES` > default false).
- `rza_migrate_db -fail-fast true|false`: hentikan verify flow setelah fail pertama untuk mode lokal cepat (flag > env `RZA_FAIL_FAST` > default false).
- `rza_migrate_db -parallel-read true|false`: aktifkan paralelisasi aman tahap collect source + history preload saat `gen`, serta forbidden SQL scan saat verify/precheck (flag > env `RZA_PARALLEL_READ` > default true).
- `rza_migrate_db -parallel-workers <n>`: jumlah worker untuk mode parallel-read (flag > env `RZA_PARALLEL_WORKERS` > default 8, minimum 1).
- `rza_get_object_db`: export object DB ke `db_object/*` dan tulis `out/get_object_summary.json`.
- `rza_get_object_db -dry-run`: hitung create/update/delete tanpa write file.
- `rza_get_object_db -require-healthy-constraints`: export tetap selesai, tetapi
  return `1` jika observed CHECK/FK/key state disabled atau untrusted; database
  tidak diubah.
- `rza_get_object_db -preview-input <dir> -preview-output <dir>`: preview format modul SQL lokal tanpa koneksi DB.
- `rza_migrate_rest -config <path>`: jalankan REST API untuk trigger migrate SQL Server (up/down/steps/run-file) dengan source `migrations/`.
  - hardening default runtime REST: `allow_down=false`, `allow_run_specific_file=false`, timeout operasi best-effort (hardcoded).

Exit code:

- `0` jika `fail.count == 0`
- `1` jika `fail.count > 0`
- `rza_migrate_check` menulis `new_migrations.files=[]` dan note `generation_enabled=false` pada setiap run yang mencapai engine.
- Default checker adalah `profile=authoring` dengan `migration-validation=integrity`; mode ini tidak merekonstruksi history, tidak memeriksa core semantic contract/wrapper migration, tetapi tetap menolak folder tidak flat, nama/version invalid, pair rusak, core pair hilang, SQL kosong, serta encoding invalid.
- `rza_migrate_check` menulis `notes.check_attestation` schema v2 berisi profile, check/read-only mode, generation/database flags, matriks validation, manifest path case-sensitive, raw SHA-256 per SQL, ukuran, jumlah file, total byte, dan aggregate SHA-256.
- Gate GitLab wajib memverifikasi checksum binary, menjalankan `-check -profile authoring -migration-validation integrity`, mengandalkan exit code, dan menyimpan `out/output.json`; gate berlaku untuk MR, `sit`, `uat`, dan tag.
- Integrasi pipeline kantor wajib melalui local include `/.gitlab/rza_migrate_check.gitlab-ci.yml` pada special stage `.pre`; include checker dilarang mendefinisikan global `stages`, `workflow`, `default`, atau variables yang dapat menimpa template pusat.
- Gate memakai `rza_migrate_check -check` sehingga checkout CI tetap read-only; tidak memerlukan `git`, `jq`, package installation, source build, maupun koneksi MSSQL.
- `rza_get_object_db` juga `1` untuk error hard (koneksi/query/io), namun warning modul programmability tetap non-fatal dan dicatat di summary.
- Jika `skip_unsupported_tables=true`, tabel unsupported di extractor tidak fail-hard; item tersebut dicatat sebagai warning terstruktur `TABLE_EXTRACT_SKIPPED_UNSUPPORTED`.
- Jika `skip_unsupported_tables=false`, tabel unsupported menyebabkan run extractor fail-hard.
- CHECK/FK disabled atau untrusted tetap dihasilkan sebagai logical definition
  dan dicatat dengan warning `CONSTRAINT_STATE_REVIEW_REQUIRED`; warning ini
  non-fatal kecuali `-require-healthy-constraints` digunakan.
- `rza_migrate_rest` return `1` untuk startup/runtime fatal.

## 23. Engine Config Baseline (`app_config.go`)

Minimum config wajib:

- `DbObjectDir = db_object`
- `MigrationsDir = migrations`
- `PatchDir = patch`
- `OutDir = out`
- `LegacyAuditMode = fail`
- `ExistingTableMatchMode = bootstrap-compatible`
- `EnableGenAutoIndent = true`
- `EnablePerfNotes = false`
- `EnableFailFast = false`
- `EnableParallelRead = true`
- `ParallelWorkers = 8`
- `PairSeqWidth = 3`
- `MaxPairsPerRun = 999`
- `AllowedCategories = types,tables,indexes,functions,views,procedures,patch`
- `AllowedObjectTokenCharset = A-Za-z0-9_.#`
- `LineEnding = LF`
- `TrimTrailingSpaces = true`
- `EnsureFinalNewline = true`
- `FailHardOnParseError = true`
- `NoRegexParsing = true` (parser utama `rza_migrate_db`; `rza_get_object_db` header-normalizer masih regex-assisted)
- `EnableProgrammabilitySourceContractForNewOnly = false` (default)
- jika `EnableProgrammabilitySourceContractForNewOnly = true`, source-contract programmability ketat diterapkan untuk object baru; object existing tetap diverify header `CREATE OR ALTER`.
- `ForbiddenSqlTokens = GO, USE, CREATE DATABASE, ALTER DATABASE, CREATE LOGIN, ALTER LOGIN, CREATE USER`
- `GenerationOrder = types,tables,indexes,functions,views,procedures,patch`
- `ExecutableNameLock`:
  - `rza_migrate_check(.exe)` untuk checker developer tanpa generation
  - `rza_migrate_db(.exe)` untuk generator/verifier
  - `rza_get_object_db(.exe)` untuk extractor
  - `rza_migrate_rest(.exe)` untuk REST executor

## 24. Verify Checklist Integritas `migrations/`

`verify` wajib memastikan:

- semua migration berpasangan up/down
- tidak ada duplikat version untuk pair berbeda
- category tidak keluar whitelist
- tidak ada forbidden token
- tables/indexes/patch punya wrapper + audit call minimal
- programmability wajib `CREATE OR ALTER`
- programmability 1 file hanya 1 object
- `migrations/` flat

## 25. Acceptance Criteria

1. Semua command patuh CLI contract dan exit code contract.
2. Fail rule hard lock selalu ditegakkan.
3. Generator deterministik untuk input yang sama.
4. Whitespace-only tidak menghasilkan migration baru.
5. Semua keputusan penting tercermin di audit SQL atau `out/output.json`.
6. Policy destructive, patch, dan ALTER precheck berjalan persis sesuai aturan.
7. Extractor summary contract konsisten: warning modul tidak menggagalkan run total dan tercatat di `out/get_object_summary.json` (`export.warning_count`, `export.warnings[]`).

## 26. Traceability ke `rules.md`

Cakupan section `rules.md` ke dokumen ini:

- `1` sampai `3`: tercover pada Section 1, 2, 5.
- `4` sampai `7`: tercover pada Section 6, 7, 8, 9.
- `8` sampai `13`: tercover pada Section 10, 11, 12, 13, 14, 15.
- `14` sampai `18`: tercover pada Section 16, 17, 18, 19, 20.
- `19` sampai `22`: tercover pada Section 21, 22, 23, 24.
- `23` sampai `25`: tercover pada Section 3, 17, 25.

## 27. Output Pattern dan Runner Compatibility

Contoh pattern output untuk quality check:

- fail item mismatch object filename-content wajib severity 3 (`FILE_OBJECT_MISMATCH`).
- fail item inline default wajib severity 3 (`INLINE_DEFAULT_NOT_ALLOWED`).
- warn item drop column impacted indexes wajib severity 2 (`INDEX_RECREATE_REQUIRED`).
- extractor summary menyimpan warning partial module parse di `export.warning_count` dan `export.warnings[]`.

Runner compatibility notes (`golang-migrate`):

- Up: `migrate -path ./migrations -database \"$MSSQL_DSN\" up`
- Down 1 step: `migrate -path ./migrations -database \"$MSSQL_DSN\" down 1`
- Goto version: `migrate -path ./migrations -database \"$MSSQL_DSN\" goto <version>`
- Force dirty repair only: `migrate -path ./migrations -database \"$MSSQL_DSN\" force <version>`
- Version check: `migrate -path ./migrations -database \"$MSSQL_DSN\" version`

Aturan runtime: ketika migration melempar `THROW`, runner berhenti dan file berikutnya tidak dijalankan.

## 26. Migration History Reconstruction Report (`rza_migrate_report`)

Alat pelaporan offline untuk merekonstruksi histori perubahan objek database dari folder `migrations/` dan mendokumentasikan current desired state dari `db_object/`:
- **Source Separation**: Histori hanya membaca `migrations/*.up.sql` dan `migrations/*.down.sql`; desired-state/impact docs hanya membaca object source di `db_object/`.
- **Zero Database Access**: Tidak pernah membuka koneksi TCP/IP ke SQL Server atau database lain (`database_access = false`).
- **Zero SQL Execution**: Tidak mengeksekusi SQL apapun (`sql_execution = false`).
- **Immutability Protection**: Tidak pernah membuat, mengubah, menghapus, atau me-rename berkas migrasi sumber.
- **Core Bootstrap Exclusion**: Pair canonical `000000__core__mig__audit_init` dipindai sebagai infrastruktur tetapi tidak membentuk histori/operasi schema aplikasi di audit report. Output summary memuat analyzed pair count dan excluded pair count secara terpisah pada CSV, JSON, dan console.
- **Output Standard**:
  - Audit laporan rekonsiliasi (`docs/database/migrations/` atau `-out`):
    - `docs/database/migrations/datasets/report_summary.csv`
    - `docs/database/migrations/datasets/migration_summary.csv`
    - `docs/database/migrations/datasets/object_history.csv`
    - `docs/database/migrations/datasets/table_changes.csv`
    - `docs/database/migrations/datasets/index_changes.csv`
    - `docs/database/migrations/datasets/programmable_changes.csv`
    - `docs/database/migrations/datasets/type_changes.csv`
    - `docs/database/migrations/report.json`
  - Dokumentasi desired-state (`docs/database/` atau `-docs`):
    - `docs/database/identifiers/table_unique_ids.md` untuk enforcement uniqueness, conditional unique index, `IDENTITY`, dan ID review signals. `IDENTITY`, JOIN, nama ID-like, dan non-unique index tidak boleh dinyatakan sebagai bukti uniqueness.
    - `docs/database/relationships/relational_erd.md` untuk seluruh Mermaid graph dalam satu file, disusun sebagai Relationship Group → connected Relational Map. Setiap map dibentuk adaptif dengan batas maksimum 20 object, 24 relationship lines, 35 exact mappings, dan 9.000 karakter Mermaid, tanpa artificial Hub/parent. Node memuat type object serta key/index context untuk kolom relationship; edge memuat exact column mapping dan tidak mengklaim arah/cardinality yang tidak dibuktikan source.
    - `docs/database/relationships/relationship_catalog.md` untuk evidence, exact column mapping, join type, occurrence, source module, dan confidence.
    - `docs/database/relationships/module_lineage.md` untuk static view/procedure-to-object dependency dan dynamic-SQL manual-review list.
    - `docs/database/README.md` menjadi halaman masuk; `manifest.json` memuat semua artifact konkret dan SHA-256. Legacy `table_relationship_tree.md` tidak lagi dibuat; navigasi memakai index dan relationships/.
    - `docs/database/procedures/README.md`, `procedure_impact_catalog.md`, dan `profiles/dbo.<Procedure>.md` untuk procedure contract, branch-aware local execution flow, local object map, direct/recursive dependency map, exact write mapping, key/index context, execution/error signal, serta dynamic-SQL review.
- Observed relationship hanya valid untuk static explicit `JOIN ... ON` atau legacy comma-join `WHERE` equality dengan tepat satu resolved `alias.column` pada masing-masing sisi. Same-name columns, inequality, ambiguous/multi-column expression, temp/table variable, stale alias scope, dan dynamic SQL wajib di-skip konservatif. Module tanpa resolved static dependency wajib tetap terlihat pada coverage-review list.
- Desired-state Markdown memakai narasi Bahasa Indonesia yang natural dan istilah/header database English yang lazim. Guidance per classification ditulis satu kali pada reference section, bukan sebagai kolom template per row. Stable codes tidak diterjemahkan dan multi-value tidak memakai raw HTML. Relationship Group adalah connected component, bukan urutan/prioritas. Semua diagram berada pada satu ERD document; setiap connected Relational Map dibentuk adaptif dengan batas maksimum 20 object, 24 lines, 35 mappings, dan 9.000 karakter Mermaid tanpa artificial Hub/parent. Node memuat type serta key/index context untuk kolom relationship. Setiap non-self object pair wajib tampil, dengan setiap exact column mapping tepat sekali pada edge; `[FK]`/`[JOIN]` digunakan hanya untuk mixed-evidence pair, sedangkan occurrence/usage hanya menjadi audit detail di catalog. Dense pair yang melampaui mapping/text budget dipecah menjadi numbered `Mapping batch`. Exact column mapping, termasuk self-join, muncul tepat sekali di catalog bernomor Group/Map sama, maksimal 35 row per tabel. Self-join dipisahkan ke review section khusus. Capability ini report-only: extractor, checker apply/`-check`, migration generator, pipeline, dan migration history tidak berubah.
- Fitur ini report-only. Extractor, checker apply/`-check`, migration formatter/verifier/generator, dan migration history tidak mengubah atau menafsirkan schema berdasarkan observed relationship/ID signals; pipeline hanya membawa binary report yang sama melalui release publication.
- Procedure dependency wajib divalidasi terhadap object inventory yang benar-benar ada di `db_object`; availability di catalog tidak boleh dipakai untuk menebak nilai runtime. Scanner shared memakai token dengan source position dan statement scope: object multiline tetap dikenali, CTE unqualified hanya `NOT_PERSISTED` dalam statement-nya, object schema-qualified tidak tertimpa nama CTE, dan DML alias tidak boleh bocor lintas statement. Resolution code wajib membedakan `RESOLVED_STATIC`, `RESOLVED_DYNAMIC`, `RESOLVED_FROM_CALLER`, `EXTERNAL_DATABASE`, `UNRESOLVED_DYNAMIC`, `NOT_PERSISTED`, `NOT_FOUND`, `AMBIGUOUS`, dan `SYSTEM_OBJECT`. Unqualified/case-colliding name dengan lebih dari satu kandidat tidak boleh dipilih otomatis. Cursor, built-in function, assignment `SELECT`, `SELECT INTO`, subquery/`EXISTS`, dan `INSERT ... SELECT` tidak boleh menjadi false dependency/result-set signal. Guard `IF`/`ELSE` wajib berasal dari source asli agar literal condition tidak hilang; Mermaid edge label wajib dikutip supaya condition dengan `<`, `>`, dan karakter SQL lain tetap renderable. Local execution flow wajib top-down dan tidak boleh menampilkan `\n` sebagai teks node; local object map wajib membedakan jenis akses dari urutan eksekusi. Direct call detail wajib memakai layout vertikal yang tetap terbaca untuk guard/argument panjang. Recursive map wajib cycle-safe, mempertahankan semua direct call edge, memecah diagram di batas 20 procedure tanpa membuang edge, dan mencegah ekspansi subgraph berulang tanpa menyembunyikan child procedure atau recursive write impact.
- `-out` dan `-docs` yang relatif dihitung dari resolved repository root. `report.json` mempertahankan `reports[]` compatibility serta menambah `output_directory`, `documentation_directory`, dan `artifacts[]` berisi `{kind, root, path}` untuk setiap file konkret, termasuk setiap generated procedure profile. Desired-state docs harus selesai ditulis sebelum audit JSON dipublikasikan.
- Seluruh berkas CSV mematuhi format RFC 4180 dan di-encode menggunakan UTF-8 tanpa BOM. Jika suatu kategori tidak memiliki perubahan, berkas CSV tetap dibuat dengan baris header.
- Pengurutan data deterministik: `migration_id, category, schema_name, object_name, operation_no`.
- Determinisme berlaku untuk dataset hasil rekonstruksi; `run_at` dan `generated_at` merekam waktu eksekusi aktual dan dikecualikan dari perbandingan byte-for-byte antar-run.
- Parser filename menggunakan kontrak shared engine: core canonical, version 17 digit, category whitelist, schema `dbo`, dan `v{k}` minimal 1. Nama non-canonical menghasilkan warning terstruktur dan tidak diperlakukan sebagai filename valid.
- Output directory audit (`-out`) maupun dokumentasi (`-docs`) (termasuk target symlink/junction yang dapat di-resolve) ditolak jika sama dengan atau berada di bawah `migrations/` atau `db_object/` untuk menjaga immutability bahkan ketika path salah konfigurasi.
- Round-trip diperiksa dari inverse operasi semantik. Down kosong atau inverse yang tidak cocok menghasilkan `FAIL` dan warning; SQL ambigu yang tidak dapat dibuktikan menghasilkan `NOT_TESTED`.
- File `up` kosong tidak boleh menghasilkan operasi programmable/type semu. Lifecycle procedure/view/function/type wajib membedakan `DELETED` dan `REINTRODUCED`.
- Setiap statement `CREATE TABLE` yang dikenali menghasilkan tepat satu operasi create dan satu operasi per child column/constraint; wrapper audit/precheck sebelum migration body tidak boleh menyebabkan statement yang sama dipindai ulang. Identifier kolom terdelimitasi dengan `[]` atau `""` harus tetap utuh walaupun memuat spasi.
- `ALTER COLUMN` hanya melaporkan atribut yang benar-benar berubah; tipe atau nullability yang identik tidak menambah operasi.
- `ALTER TABLE ... ADD` mendukung banyak kolom, banyak constraint, atau campuran keduanya dalam satu statement; comma di dalam type/expression/composite key tidak boleh memecah item secara salah. Multiple `DROP COLUMN` dan `DROP CONSTRAINT [IF EXISTS]` juga diekstrak per target.
- Rekonstruksi mempertahankan urutan source lintas CREATE/DROP/ALTER/rename. Ini wajib agar pola add→rename→alter serta drop→re-add dalam migration yang sama memperoleh state lama yang benar.
- Drop+recreate constraint/index dengan identity sama direduksi menjadi satu `ALTER_CONSTRAINT`, `ALTER_DEFAULT`, atau `ALTER_INDEX`; row menyimpan old/new definition/attributes. Round-trip `PASS` hanya jika semua operasi dalam migration mempunyai inverse, bukan hanya salah satunya.
- Reducer mempertahankan state sepanjang seluruh rantai version. Saat table di-drop atau di-create kembali, cache kolom/constraint/index di-reset sehingga before-value versi berikutnya berasal dari incarnation terbaru. Acceptance test wajib mencakup v1→v7 dengan `DELETED`/`REINTRODUCED`, perubahan campuran di v7, serta kasus negatif yang tidak boleh memakai state pra-drop.
- Perubahan komentar atau whitespace semata tidak dianggap sebagai perubahan semantik objek.
- SQL yang ambigu menghasilkan baris operasi dengan `confidence = LOW` dan `notes = MANUAL_REVIEW_REQUIRED`.

## Report layout V2

Report layout V2 acceptance: default docs/database berisi README/manifest, identifiers, relationships, procedures, dan migrations/datasets. -docs mengendalikan root; default audit mengikuti root/migrations; explicit -out tetap didukung dengan datasets child. Root index/backlinks, concrete hash manifest, no out default, repeatable rerun, manual collision preflight/no-write, ownership-only stale cleanup, nested symlink/overlap rejection wajib teruji. CSV/history semantics tidak berubah; extractor/checker apply/CI/generator tidak mendapat schema action. Lihat [layout contract](docs/report_output_layout.md).


## Audit object precision (2026-10-12)

Acceptance report object precision: token-based multiline DDL, inline/ALTER FK, complete datatype, function consumers, scoped JOIN dengan OUTER/hint dan quoted identifier, canonical case, unfiltered UPDATE assignments, COLUMN_NOT_FOUND, strict INSERT targets, parameter literal preservation, actual complete dynamic literal, Unicode no-write, serta actionable filename identity failure. Real corpus dan committed fixture membuktikan no-write/idempotence. Extractor/checker/generator tidak mendapat schema correction baru; shared MERGE JOIN grammar berlaku pada semua consumers.

Detail bukti: docs/review_report_object_precision_20261012.md.
