rules-final.md — Database Migration Standards (SQL Server + golang-migrate + offline generator rza_migrate_db)

Target:

DB: Microsoft SQL Server (single schema dbo)

Runner: golang-migrate

Generator offline: rza_migrate_db

Time: gunakan GETDATE() (local time server SQL), tanpa UTC.


Dokumen ini mengunci standar agar:

migrasi deterministik & aman di skala besar

audit jelas (severity/decision/action/reason konsisten)

generator offline bisa diff & rebuild state history dengan tepat

pipeline GitLab CI stabil tanpa “out-of-band DDL”


> FAIL rule (explicit, hard lock): jika ada pelanggaran dengan severity=3 atau out/output.json.fail.count > 0, maka proses WAJIB FAIL HARD (exit code 1), tidak boleh menghasilkan/menulis migration baru, dan pipeline harus berhenti.

---

## Status Implementasi (Audit 2026-02-19)

Legend status:

- `SEMPURNA`: sudah terimplementasi dan tervalidasi end-to-end oleh gate/test saat ini.
- `KURANG_SEMPURNA`: sudah ada implementasi, namun coverage/ketegasan belum penuh.
- `BELUM_DILAKUKAN`: belum ada implementasi yang relevan.

| Point | Status | Catatan singkat |
| --- | --- | --- |
| 1. Istilah & Konsep | SEMPURNA | Terminologi dan alur source/history/patch/reconstruction sudah konsisten di engine + dokumen turunan. |
| 2. Foldering Repo | SEMPURNA | Contract `db_object/`, `patch/`, `migrations/` sudah divalidasi strict + fixture pass/fail lengkap. |
| 3. Tooling & Pipeline | SEMPURNA | Flow command + contract runner `migrate up` + stop-behavior gate sudah tervalidasi (`validate_runner_stop_behavior_contract.ps1`) dan terikat ke CI. |
| 4. Larangan & Safety Guard | SEMPURNA | Forbidden token fail-hard aktif, serta guard drift repo (`REPO_STATE_DRIFT`) memberi sinyal terkontrol untuk mismatch desired-vs-history. |
| 5. Konvensi File & Pairing | SEMPURNA | Pairing, version 17 digit, grammar filename, dan parse structured sudah ditegakkan. |
| 6. Generator Ordering | SEMPURNA | Ordering kategori + deterministic sorting + `v{k}` progression sudah berjalan dan diuji. |
| 7. Data Master Validation & Auto-Format | SEMPURNA | Auto-format dan source contract berjalan deterministik (LF/trim/newline/canonical table) dengan enforcement ketat pada semua kategori source MVP. |
| 8. Destructive Policy & Decision Semantics | SEMPURNA | Policy drop table manual-only, drop column semi-auto, whitelist decision/action/reason/severity sudah aktif. |
| 9. Transaction & Error Strategy | SEMPURNA | Wrapper TRY/TRAN/CATCH/ROLLBACK/THROW untuk tables/indexes/patch sudah diwajibkan dan dicek. |
| 10. Core Audit Init | SEMPURNA | Contract final core audit sudah ditegakkan (mig_event + index audit + log proc + lock proc + down cleanup konsisten). |
| 11. Audit Logging (DDL) | SEMPURNA | Shape audit logging final tervalidasi strict (parameter/kolom error, GETDATE, whitelist metadata, kontrak wrapper audit). |
| 12. Wrapper Standar | SEMPURNA | Wrapper tables/indexes/patch tervalidasi strict, metadata deterministik wajib, dan legacy mode `warn` sudah dihapus (strict-only). |
| 13. Rules Down Migration | SEMPURNA | Down wrapper, tolerant precheck, dan warning destruktif sudah dijalankan sesuai policy saat ini. |
| 14. Patch Policy | SEMPURNA | Canonicalize, duplication, immutability hash, source integrity, serta metadata header patch (`patch_folder/patch_name/readme`) sudah divalidasi konsisten. |
| 15. Output `out/output.json` | SEMPURNA | Struktur root + fail/pass/new_migrations/notes + exit behavior sudah konsisten dan dites. |
| 16. Locking 1 Run per DB | SEMPURNA | Contract lock tervalidasi lewat template migration job + CI include + contract check script (`usp_mig_try_get_lock` sebelum `migrate up`). |
| 17. MVP Supported Subset | SEMPURNA | Subset didukung jelas; SQL di luar subset fail-fast `UNSUPPORTED_SQL_SUBSET`. |
| 18. ALTER COLUMN Pre-check Map | SEMPURNA | Precheck deterministic (nullability/shrink/type conversion) sudah diemit dan diuji. |
| 19. `constraints_naming.md` | SEMPURNA | Dokumen hanya guideline opsional; constraint tetap harus bernama eksplisit, tanpa prefix wajib. |
| 20. Repo Hygiene | SEMPURNA | `.gitignore`/`.gitattributes` direkomendasikan dan dapat dibuat bootstrap, tetapi tidak memblokir SQL authoring. |
| 21. CLI Contract | SEMPURNA | `fmt|verify|gen` + exit code contract sudah sesuai dan dites. |
| 22. Engine Config & Verify Rules | SEMPURNA | Hard-lock config aktif; modular internal `fsx`, `sqlfmt`, `sqlparse`, `history`, `diff`, `emit`, dan `verify` sudah aktif dipakai engine. |
| 23. Final Locks / Gaps Closed | SEMPURNA | Lock/gap utama sudah ditutup dan tervalidasi oleh hardening gate final tanpa fail residual. |
| 24. Contoh Fail/Pass Output | SEMPURNA | Pattern output fail/pass sudah tersedia di docs + fixture/golden. |
| 25. Runner Cheatsheet | SEMPURNA | Cheatsheet sudah terdokumentasi dan konsisten dengan asumsi runner saat ini. |

Rekap:

- `SEMPURNA`: 25 point
- `KURANG_SEMPURNA`: 0 point
- `BELUM_DILAKUKAN`: 0 point




---

Table of Contents

1. Istilah & Konsep


2. Foldering Repo


3. Tooling & Pipeline


4. Larangan & Safety Guard


5. Konvensi File & Pairing golang-migrate


6. Generator Ordering 6.1 Versioning v{k} dan aturan naik versi 6.2 State reconstruction (history → state)


7. Data Master db_object/ — Validasi & Auto-Format


8. Policy Destructive & Decision Semantics


9. Transaction & Error Strategy (MSSQL)


10. Core Audit Init — DDL Final (wajib)


11. Audit Logging (DDL) — Final (Normalized)


12. Wrapper Standar (Tables/Indexes/Patch)


13. Rules Down Migration


14. Patch Policy


15. Output Generator out/output.json


16. Locking: 1 migrate run per DB


17. MVP Supported Subset


18. ALTER COLUMN Pre-check Templates (Generator Map)


19. constraints_naming.md (Opsional)


20. Repo Hygiene (.gitignore & .gitattributes)


21. CLI Contract (rza_migrate_db)


22. Engine Config & Codebase Skeleton 22.3 Verify rules (migrations folder integrity)


23. Final Locks / Gaps Closed


24. Contoh Fail/Pass Output & Error Patterns


25. Runner Cheatsheet (golang-migrate)




---

1) Istilah & Konsep

Source-of-truth: db_object/ (data master yang manusia edit). Ini yang menentukan “desired state”.
History: migrations/ (hasil generate + patch migrations) yang menjadi urutan deploy.
Patch: perubahan data sekali jalan, dibuat manual di patch/<folder>/ lalu di-wrap & digenerate menjadi migration.
Semantic diff: perubahan bermakna pada schema/object, bukan sekadar whitespace.
State reconstruction: proses engine membaca semua migration terdahulu untuk menyusun state terakhir (per object), agar bisa membangun down yang benar dan menghitung versi v{k} dengan tepat.


---

2) Foldering Repo

2.1 Source-of-truth (db_object/)

db_object/
  tables/
    dbo.<TableName>.sql
  indexes/
    dbo.<TableName>.sql        # 1 file per table, berisi banyak CREATE INDEX untuk table tsb
  views/
    dbo.<ViewName>.sql
  functions/
    dbo.<FunctionName>.sql
  procedures/
    dbo.<ProcedureName>.sql    # wajib CREATE OR ALTER

Mode transisi kompatibel:

Engine tetap membaca path legacy `db_object/schema/<category>/` bila path baru belum tersedia.

Jika file object yang sama ada di kedua path, path baru `db_object/<category>/` menjadi prioritas.

Jika path legacy terdeteksi, engine emit warning `LEGACY_DB_OBJECT_LAYOUT_DETECTED` (non-fatal) agar repo segera dimigrasikan ke path canonical.

Jika object yang sama ada di canonical+legacy dengan isi berbeda => FAIL `DB_OBJECT_DUAL_PATH_CONFLICT`.

Jika object yang sama ada di canonical+legacy dengan isi sama => warn `LEGACY_DB_OBJECT_DUPLICATE_CANONICAL`.

Aturan wajib:

Schema hanya dbo.

Nama file adalah object name persis (dbo.X.sql), tanpa spasi.

Isi file harus mendefinisikan object yang sama dengan nama file (mismatch => FAIL).

indexes/dbo.<Table>.sql hanya boleh berisi index untuk table itu (campur table lain => FAIL).


2.2 Patch source (patch/)

patch/
  <foldername>/
    README.md
    up.sql
    down.sql

Aturan wajib:

Folder patch hanya boleh berisi 3 file di atas (file lain => FAIL).


2.3 Migration output (migrations/)

Wajib 1 folder flat: semua file migration harus berada langsung di migrations/ (tanpa subfolder) agar runner tidak skip.

Core init audit:

000000__core__mig__audit_init.up.sql

000000__core__mig__audit_init.down.sql


Schema migrations (generated):

{version}_{category}_{object}_v{k}.up.sql

{version}_{category}_{object}_v{k}.down.sql


Patch migrations (generated dari patch folder):

{version}_patch_{patch_name}_v{k}.up.sql

{version}_patch_{patch_name}_v{k}.down.sql



---

3) Tooling & Pipeline

Tool eksekusi: golang-migrate.

Trigger GitLab CI: pipeline bisa dipicu oleh merge atau tag. Sistem harus aman walau pipeline sering berjalan.

Rule:

job migration hanya menjalankan perubahan yang ada di migrations/ (perilaku standar runner).

generator bertanggung jawab memastikan migrations/ benar, berpasangan, terurut, dan deterministik.

Stop behavior: saat satu migration file error (THROW), runner stop dan tidak menjalankan file setelahnya.



---

4) Larangan & Safety Guard

4.1 No GO

GO dilarang pada semua file di migrations/, db_object/, patch/.
Jika ditemukan => FAIL.

4.2 Forbidden SQL

Dilarang:

USE <db>

CREATE DATABASE, ALTER DATABASE

CREATE LOGIN, ALTER LOGIN, CREATE USER (security/identity out-of-band)


Jika ditemukan => FAIL.

4.3 No out-of-band DDL

Semua perubahan schema yang dikelola sistem harus lewat repo (db_object/ + migrations/) dan dieksekusi via pipeline.
Perubahan manual di environment dianggap drift; generator boleh mengeluarkan REQUIRES_MANUAL bila berisiko.


---

5) Konvensi File & Pairing golang-migrate

5.1 Pairing rule (CRITICAL)

golang-migrate memasangkan up/down berdasarkan version yang sama (angka sebelum underscore pertama).

Setiap migration wajib punya pasangan:

{version}_... .up.sql

{version}_... .down.sql


Missing pair => FAIL hard.

5.2 Version format (pair-safe)

base_ts := YYYYMMDDHHMMSS

pair_seq := 000..999 (per pair)

version := base_ts + pair_seq (17 digit)


Contoh:

20260217073015000_tables_dbo.TableA_v2.up.sql

20260217073015000_tables_dbo.TableA_v2.down.sql


Jika dalam satu run menghasilkan >999 pair => FAIL.

5.3 Filename grammar (structured parsing; no regex)

Engine harus mem-parse filename tanpa regex.

Schema migration:

<version>_<category>_<object>_v<k>.(up|down).sql

category ∈ {types,tables,indexes,functions,views,procedures}

object token hanya: A-Z a-z 0-9 _ . # (tanpa spasi/bracket)

k >= 1


Patch migration:

<version>_patch_<patch_name>_v<k>.(up|down).sql


Jika parsing gagal => FAIL.


---

6) Generator Ordering

Jika ada perubahan multi kategori pada satu run gen, generator harus menghasilkan migration pair dalam urutan:

1. tables


2. indexes


3. views


4. functions


5. procedures


6. patch (selalu terakhir)



pair_seq harus meningkat mengikuti urutan di atas.

Tambahan (determinisme):

Di dalam satu kategori, urutan pair ditentukan oleh sort lexicographic atas object (case-sensitive sesuai filename), lalu v{k}.

Engine dilarang iterasi map tanpa sort saat emit.

Jika ada beberapa perubahan pada object yang sama pada 1 run, engine harus menggabungkan menjadi 1 pair (kecuali secara eksplisit dipisah oleh policy), supaya history tidak “pecah-pecah”.



---

6.1 Versioning v{k} dan aturan naik versi

Setiap object (mis. dbo.TableA di kategori tables) memiliki versi internal k.

k hanya naik saat ada semantic diff setelah fmt (whitespace-only tidak boleh menaikkan k).

Untuk programmability (`views/functions/procedures`), perbedaan style-only berupa bracket identifier (`[dbo].[X]` vs `dbo.X`) dan terminator `;` tidak dianggap semantic diff (tidak menaikkan `k`).

k dihitung dari history migrations/ melalui state reconstruction, bukan dari timestamp.

k untuk patch adalah per patch_name.



---

6.2 State reconstruction (history → state)

Tujuan: engine bisa “mengerti state terakhir” tanpa konek DB.

Aturan:

Engine mem-parse semua file migrations/*.up.sql dengan parser structured (subset MVP), menghasilkan state per object.

Bila ada migration tidak bisa diparse (di luar subset), engine FAIL (bukan warn), karena state tidak bisa dipercaya.

Engine wajib mendeteksi konflik history (contoh: dua file berbeda mengklaim version+category+object+v{k} yang sama) => FAIL.


Output minimal reconstruction (untuk debug di out/output.json / notes):

last_seen_version per object

last_seen_vk per object

hash konten normalized per object (untuk deteksi drift)



---

7) Data Master db_object/ — Validasi & Auto-Format

7.1 Auto-format in-place

db_object/** adalah data master, jadi engine wajib auto-rapihin sebelum parse/diff.

Aturan:

Normalize newline ke LF

Encoding file SQL disimpan sebagai UTF-8 tanpa BOM.

Saat fmt/check lokal, UTF-8 BOM serta UTF-16/UTF-32 yang memiliki BOM valid akan dikonversi lossless ke UTF-8. Pola UTF-16 tanpa BOM hanya dikonversi bila pola NUL byte mengidentifikasi endian dengan aman. Byte non-UTF8, UTF-16/UTF-32 rusak, dan encoding legacy yang tidak dapat diidentifikasi tetap => FAIL (`INVALID_SQL_ENCODING`); engine tidak menebak code page karena dapat mengubah identifier atau string literal. Mode CI `-check` tidak menulis checkout dan menghasilkan `FMT_CHECK_CHANGES_REQUIRED` sampai developer menjalankan check/fmt lokal.

Trim trailing spaces

Pastikan file diakhiri newline

Keyword casing tidak dipaksa (preserve)

Tidak boleh mengubah string literal & komentar

Nama constraint yang sudah tercatat pada migration history adalah nama stabil. Sebelum generation, jika constraint pada `db_object/tables` mempunyai definisi semantik identik tetapi nama berbeda, engine wajib mengembalikan nama source ke nama dari version terakhir. Kesetaraan harus mencakup jenis, kolom dan urutan, clustered/nonclustered, expression, serta reference. Definisi ambigu atau perubahan semantik tidak boleh ditebak atau disamarkan.


7.2 Table formatting canonical

Trailing comma (koma di akhir item), bukan leading comma.

Bentuk standar:

CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(5) NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id)
);

7.3 Table snapshot contract (MVP)

Di file table snapshot (`db_object/tables/*.sql`) dilarang:

IF ... BEGIN ... END wrapper

DROP TABLE, ALTER TABLE, CREATE INDEX

WITH (...) advanced table options (filegroup/partition/temporal)


Artinya table snapshot harus benar-benar desired CREATE TABLE saja.

7.3.1 Existing-table bootstrap compatibility

Untuk migration initial `CREATE TABLE` ketika target table sudah ada, wrapper
memakai profile `bootstrap-compatible`.

Wajib cocok:

- setiap kolom dalam source tersedia dengan nama exact-case;
- base type kompatibel (`DECIMAL`/`NUMERIC` dan `ROWVERSION`/`TIMESTAMP`
  diperlakukan sebagai synonym MSSQL);
- identity-kind cocok;
- computed-kind cocok.

Tidak memblokir skip initial `CREATE TABLE`:

- extra column dan column order;
- length, precision/scale, nullability, dan collation;
- default, PK/UQ/CHECK/FK, serta enabled/trusted state;
- physical index option, compression, filegroup, dan partition placement.

Pelonggaran ini hanya berlaku pada guard initial `CREATE TABLE`. Explicit
`ALTER`, `DROP`, `ADD/DROP CONSTRAINT`, dan `CREATE/DROP INDEX` tetap wajib
memakai precheck operation-specific yang ketat.

7.4 Programmability header wajib

Wajib:

Procedure: CREATE OR ALTER PROCEDURE dbo.X

Function:  CREATE OR ALTER FUNCTION dbo.X

View:      CREATE OR ALTER VIEW dbo.X


Jika tidak sesuai => FAIL.

7.5 Content ↔ filename consistency

`db_object/tables/dbo.TableA.sql` harus mendefinisikan dbo.TableA.

`db_object/indexes/dbo.TableA.sql` hanya boleh membuat index ON dbo.TableA.

Programmability (views/functions/procedures) juga harus match dbo.<Name>.


Mismatch/campuran => FAIL.

7.6 Index file contract (MVP)

`db_object/indexes/dbo.<Table>.sql`:

Hanya boleh berisi CREATE INDEX / CREATE UNIQUE INDEX untuk dbo.<Table>.

Dilarang DROP INDEX di data master.

Semua index harus punya nama deterministik.


7.7 Inline DEFAULT dilarang (FAIL)

Di `db_object/tables/*.sql`, inline default seperti:

status VARCHAR(6) NULL DEFAULT('NEW')

dilarang.

Default harus table-level constraint di dalam CREATE TABLE (bukan inline DEFAULT):

CONSTRAINT DF_TableA__status DEFAULT('NEW') FOR status

Jika ditemukan inline DEFAULT => FAIL (INLINE_DEFAULT_NOT_ALLOWED).

7.8 NULL/NOT NULL omission

Jika NULL/NOT NULL tidak ditulis pada kolom, engine treat sebagai NULL.

7.9 Whitespace-only changes

Jika setelah fmt tidak ada perubahan semantic, tidak membuat migration baru dan v{k} tidak naik.

Catatan programmability:

Perubahan style-only yang hanya mengubah bracket identifier dan/atau terminator `;` juga dianggap no-op (tidak emit migration baru).


---

8) Policy Destructive & Decision Semantics

8.1 Severity mapping (fixed)

0 = INFO (ok/skip normal)

1 = NOTICE (schema-impact)

2 = WARNING (follow-up / risk / manual)

3 = ERROR (fail)


8.2 Decision semantics (whitelist)

decision ∈ {run, ok, skip, warn, fail}

Catatan:

warn adalah event non-fatal yang tetap lanjut (contoh: follow-up required), bukan pengganti fail.


8.3 Action whitelist (final)

Tables:

create_table, drop_table


Columns:

add_column, drop_column, alter_column, backfill


Defaults/Checks/Uniques:

add_default, drop_default, add_check, drop_check, add_unique, drop_unique


Indexes:

add_index, drop_index


Foreign keys:

add_fk, drop_fk


Programmability:

create_or_alter_view, create_or_alter_function, create_or_alter_procedure


Meta/Audit:

followup_required, migration_failed


8.4 Reason whitelist (final)

Precheck/Skip:

TABLE_EXISTS, TABLE_NOT_EXISTS, COLUMN_EXISTS, COLUMN_NOT_EXISTS, INDEX_EXISTS, INDEX_NOT_EXISTS, DF_EXISTS, DF_NOT_EXISTS, FK_EXISTS, FK_NOT_EXISTS, CK_EXISTS, CK_NOT_EXISTS, UQ_EXISTS, UQ_NOT_EXISTS, OBJECT_EXISTS, OBJECT_NOT_EXISTS


Apply:

CREATED, DROPPED, ALTERED, BACKFILLED, REPLACED


Dependency:

DROPPED_IMPACTED_INDEX, DROPPED_DEPENDENCY


Risk/Follow-up:

INDEX_RECREATE_REQUIRED, REQUIRES_MANUAL, NARROWING_WOULD_TRUNCATE, FK_REFERENCED_BY_OTHERS, DATA_RESTORE_NOT_GUARANTEED, DRIFT_DETECTED


Spec violation:

INLINE_DEFAULT_NOT_ALLOWED


Error:

SQL_ERROR


8.5 Skip vs Stop (policy-only)

skip hanya boleh terjadi karena policy precheck (mis. unsafe alter/shrink/not-null) atau kondisi target sudah sesuai.

skip bukan mekanisme recovery error.

Jika terjadi SQL runtime error pada tables/indexes/patch => fail hard (log fail lalu THROW).


8.6 Destructive policy (detail)

DROP TABLE (manual-only)
Generator tidak boleh emit DROP TABLE sebagai tindakan otomatis.
Jika desired state menghilangkan table: engine harus menghasilkan migration yang skip + warn (severity=2) dengan reason REQUIRES_MANUAL, plus instruction jelas di message.
Down migration untuk kasus ini tidak dijamin (harus warn DATA_RESTORE_NOT_GUARANTEED).

DROP COLUMN (semi-automatic)
Generator boleh emit ALTER TABLE ... DROP COLUMN, tetapi wajib:

drop dependency yang langsung konflik (DF/CK/UQ/FK/index yang memakai kolom)

tidak auto-recreate index (follow-up required)

log warn dengan reason INDEX_RECREATE_REQUIRED


DROP INDEX
Generator boleh drop index jika index tidak ada di desired state.
Harus precheck existence; jika tidak ada => skip (INFO).


---

9) Transaction & Error Strategy (MSSQL)

9.1 Wrapper wajib untuk tables/indexes/patch

Semua migration category tables, indexes, dan patch wajib dibungkus TRY/TRAN/COMMIT/CATCH/ROLLBACK/THROW.

Rule error:

jika error: ROLLBACK dulu, log fail setelah rollback, lalu THROW.


9.2 Programmability (view/function/procedure)

Tidak wajib wrapper.
Wajib CREATE OR ALTER.

9.3 Stop behavior

THROW cukup untuk menghentikan golang-migrate (tidak perlu RAISERROR).

9.4 Multi-step definition

Multi-step = satu intent perubahan schema membutuhkan lebih dari satu statement (drop deps → action).

Multi-step harus:

deterministik (tanpa cari nama via sys.* pada runtime)

tercatat di audit

atomic via transaction untuk tables/indexes/patch



---

10) Core Audit Init — DDL Final (wajib)

File wajib dibuat paling awal:

000000__core__mig__audit_init.up.sql

000000__core__mig__audit_init.down.sql


10.1 Up (create audit objects)

Minimal wajib mencakup:

table dbo.mig_event

proc dbo.usp_mig_log_event

(opsional tapi direkomendasikan) proc lock dbo.usp_mig_try_get_lock


Catatan penting: audit object harus tanpa GO, tanpa USE, dan aman di-run sekali (precheck IF OBJECT_ID...).

Contoh DDL (final shape):

-- 000000__core__mig__audit_init.up.sql

IF OBJECT_ID(N'dbo.mig_event', N'U') IS NULL
BEGIN
  CREATE TABLE dbo.mig_event (
    event_id      BIGINT IDENTITY(1,1) NOT NULL CONSTRAINT PK_mig_event PRIMARY KEY,
    run_key       UNIQUEIDENTIFIER NOT NULL,
    version       BIGINT NOT NULL,
    category      NVARCHAR(16) NOT NULL,
    file_name     NVARCHAR(260) NOT NULL,
    decision      NVARCHAR(8) NOT NULL,
    severity      TINYINT NOT NULL,
    action        NVARCHAR(64) NOT NULL,
    target        NVARCHAR(256) NOT NULL,
    reason        NVARCHAR(64) NOT NULL,
    message       NVARCHAR(4000) NOT NULL,
    detail        NVARCHAR(MAX) NULL,

    error_number  INT NULL,
    error_line    INT NULL,
    error_procedure NVARCHAR(128) NULL,

    created_at    DATETIME2(7) NOT NULL
      CONSTRAINT DF_mig_event__created_at DEFAULT (GETDATE())
  );

  CREATE INDEX IX_mig_event__run_key
    ON dbo.mig_event (run_key, created_at);

  CREATE INDEX IX_mig_event__version
    ON dbo.mig_event (version, created_at);
END

IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NULL
BEGIN
  EXEC (N'CREATE PROCEDURE dbo.usp_mig_log_event AS BEGIN SET NOCOUNT ON; END');
END

ALTER PROCEDURE dbo.usp_mig_log_event
  @run_key     UNIQUEIDENTIFIER,
  @version     BIGINT,
  @category    NVARCHAR(16),
  @file_name   NVARCHAR(260),
  @decision    NVARCHAR(8),
  @severity    TINYINT,
  @action      NVARCHAR(64),
  @target      NVARCHAR(256),
  @reason      NVARCHAR(64),
  @message     NVARCHAR(4000),
  @detail      NVARCHAR(MAX) = NULL,
  @error_number INT = NULL,
  @error_line   INT = NULL,
  @error_procedure NVARCHAR(128) = NULL
AS
BEGIN
  SET NOCOUNT ON;

  INSERT INTO dbo.mig_event(
    run_key, version, category, file_name,
    decision, severity, action, target, reason,
    message, detail,
    error_number, error_line, error_procedure
  )
  VALUES(
    @run_key, @version, @category, @file_name,
    @decision, @severity, @action, @target, @reason,
    @message, @detail,
    @error_number, @error_line, @error_procedure
  );
END

IF OBJECT_ID(N'dbo.usp_mig_try_get_lock', N'P') IS NULL
BEGIN
  EXEC (N'CREATE PROCEDURE dbo.usp_mig_try_get_lock AS BEGIN SET NOCOUNT ON; END');
END

ALTER PROCEDURE dbo.usp_mig_try_get_lock
  @resource NVARCHAR(255),
  @timeout_ms INT = 60000
AS
BEGIN
  SET NOCOUNT ON;

  DECLARE @res INT;
  EXEC @res = sp_getapplock
    @Resource = @resource,
    @LockMode = 'Exclusive',
    @LockOwner = 'Session',
    @LockTimeout = @timeout_ms;

  IF (@res < 0)
  BEGIN
    THROW 51000, 'Failed to acquire migration lock (sp_getapplock).', 1;
  END
END

10.2 Down (drop audit objects)

Down harus tolerant (precheck), tetapi tetap aman.

-- 000000__core__mig__audit_init.down.sql

IF OBJECT_ID(N'dbo.usp_mig_try_get_lock', N'P') IS NOT NULL
  DROP PROCEDURE dbo.usp_mig_try_get_lock;

IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  DROP PROCEDURE dbo.usp_mig_log_event;

IF OBJECT_ID(N'dbo.mig_event', N'U') IS NOT NULL
  DROP TABLE dbo.mig_event;


---

11) Audit Logging (DDL) — Final (Normalized)

11.1 Tabel event: dbo.mig_event

Kolom minimal (final):

event_id (identity)

run_key (GUID)

version (bigint)

category (nvarchar)

file_name

decision (whitelist)

severity (0..3)

action (whitelist)

target

reason (whitelist)

message

detail (json/text)

created_at default GETDATE()


Tambahan strongly recommended:

error_number, error_line, error_procedure


11.2 SP logger: dbo.usp_mig_log_event

Rule:

created_at selalu default GETDATE().

Script migration tidak mengirim finished_at atau duration_ms.

Tidak ada kolom redundant (env/db_name/tool_*).


11.3 run_key (bulletproof)

Di setiap migration file category tables/indexes/patch:

DECLARE @run_key UNIQUEIDENTIFIER = NEWID();


11.4 Severity mapping header wajib

Generator menyisipkan header comment di setiap migration file (tables/indexes/patch):

/*
Severity mapping:
0=INFO (ok/skip normal)
1=NOTICE (schema change)
2=WARNING (follow-up/manual risk)
3=ERROR (stop)
*/


---

12) Wrapper Standar (Tables/Indexes/Patch)

12.1 Prinsip

Wrapper wajib untuk tables/indexes/patch.

Precheck+skip boleh dilakukan sebelum transaksi.

Jika masuk transaksi dan error => rollback, log fail setelah rollback, THROW.

@version/@category/@file harus selalu terisi deterministik (sesuai filename).


12.2 Template wrapper (tables)

/* severity mapping ... */
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = 20260217073015000; -- engine inject
DECLARE @category NVARCHAR(16) = N'tables';
DECLARE @file NVARCHAR(260) = N'20260217073015000_tables_dbo.TableA_v2.up.sql';

IF OBJECT_ID(N'dbo.TableA', N'U') IS NOT NULL
BEGIN
  EXEC dbo.usp_mig_log_event
    @run_key, @version, @category, @file,
    N'skip', 0,
    N'create_table', N'dbo.TableA', N'TABLE_EXISTS',
    N'[PRECHECK] table already exists', NULL,
    NULL, NULL, NULL;
  RETURN;
END

BEGIN TRY
  BEGIN TRAN;

  EXEC dbo.usp_mig_log_event
    @run_key, @version, @category, @file,
    N'run', 1,
    N'create_table', N'dbo.TableA', N'CREATED',
    N'create table start', NULL,
    NULL, NULL, NULL;

  CREATE TABLE dbo.TableA (
    id INT NOT NULL,
    CONSTRAINT PK_TableA PRIMARY KEY (id)
  );

  EXEC dbo.usp_mig_log_event
    @run_key, @version, @category, @file,
    N'ok', 1,
    N'create_table', N'dbo.TableA', N'CREATED',
    N'create table ok', NULL,
    NULL, NULL, NULL;

  COMMIT;
END TRY
BEGIN CATCH
  IF @@TRANCOUNT > 0 ROLLBACK;

  EXEC dbo.usp_mig_log_event
    @run_key, @version, @category, @file,
    N'fail', 3,
    N'migration_failed', N'dbo.TableA', N'SQL_ERROR',
    N'migration failed',
    ERROR_MESSAGE(),
    ERROR_NUMBER(), ERROR_LINE(), ERROR_PROCEDURE();

  THROW;
END CATCH;

12.3 Template wrapper (indexes)

/* severity mapping ... */
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = 20260217073015001;
DECLARE @category NVARCHAR(16) = N'indexes';
DECLARE @file NVARCHAR(260) = N'20260217073015001_indexes_dbo.TableA_v1.up.sql';

IF EXISTS (
  SELECT 1
  FROM sys.indexes
  WHERE object_id = OBJECT_ID(N'dbo.TableA')
    AND name = N'IX_TableA__status'
)
BEGIN
  EXEC dbo.usp_mig_log_event
    @run_key, @version, @category, @file,
    N'skip', 0,
    N'add_index', N'dbo.TableA.IX_TableA__status', N'INDEX_EXISTS',
    N'[PRECHECK] index already exists', NULL,
    NULL, NULL, NULL;
  RETURN;
END

BEGIN TRY
  BEGIN TRAN;

  EXEC dbo.usp_mig_log_event
    @run_key, @version, @category, @file,
    N'run', 1,
    N'add_index', N'dbo.TableA.IX_TableA__status', N'CREATED',
    N'create index start', NULL,
    NULL, NULL, NULL;

  CREATE INDEX IX_TableA__status ON dbo.TableA(status);

  EXEC dbo.usp_mig_log_event
    @run_key, @version, @category, @file,
    N'ok', 1,
    N'add_index', N'dbo.TableA.IX_TableA__status', N'CREATED',
    N'create index ok', NULL,
    NULL, NULL, NULL;

  COMMIT;
END TRY
BEGIN CATCH
  IF @@TRANCOUNT > 0 ROLLBACK;

  EXEC dbo.usp_mig_log_event
    @run_key, @version, @category, @file,
    N'fail', 3,
    N'migration_failed', N'dbo.TableA.IX_TableA__status', N'SQL_ERROR',
    N'migration failed',
    ERROR_MESSAGE(),
    ERROR_NUMBER(), ERROR_LINE(), ERROR_PROCEDURE();

  THROW;
END CATCH;

12.4 Template wrapper (patch)

/* severity mapping ... */
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @version BIGINT = 20260217073015010;
DECLARE @category NVARCHAR(16) = N'patch';
DECLARE @file NVARCHAR(260) = N'20260217073015010_patch_fix_status_v1.up.sql';

/*
patch_folder: Fix Status
patch_name: fix_status
readme: patch/Fix Status/README.md
*/

BEGIN TRY
  BEGIN TRAN;

  EXEC dbo.usp_mig_log_event
    @run_key, @version, @category, @file,
    N'run', 1,
    N'backfill', N'[PATCH] fix_status', N'BACKFILLED',
    N'[PATCH:fix_status] start', NULL,
    NULL, NULL, NULL;

  -- patch body (as-is)
  -- ...

  EXEC dbo.usp_mig_log_event
    @run_key, @version, @category, @file,
    N'ok', 1,
    N'backfill', N'[PATCH] fix_status', N'BACKFILLED',
    N'[PATCH:fix_status] ok', CONCAT(N'last_rowcount=', @@ROWCOUNT),
    NULL, NULL, NULL;

  COMMIT;
END TRY
BEGIN CATCH
  IF @@TRANCOUNT > 0 ROLLBACK;

  EXEC dbo.usp_mig_log_event
    @run_key, @version, @category, @file,
    N'fail', 3,
    N'migration_failed', N'[PATCH] fix_status', N'SQL_ERROR',
    N'[PATCH:fix_status] failed',
    ERROR_MESSAGE(),
    ERROR_NUMBER(), ERROR_LINE(), ERROR_PROCEDURE();

  THROW;
END CATCH;


---

13) Rules Down Migration

13.1 Prinsip

.down.sql wajib mengembalikan schema ke versi sebelumnya.
Data yang hilang tidak dijamin kembali; jika destruktif, down wajib emit warning DATA_RESTORE_NOT_GUARANTEED.

13.2 Down transaction

Down untuk tables/indexes/patch wajib wrapper transaksi seperti up.

13.3 Precheck down tolerant

Precheck down boleh skip (bukan fail) jika object sudah tidak ada / sudah sesuai.
Tetapi SQL runtime error tetap fail + THROW.


---

14) Patch Policy

14.1 Patch source rules

patch/<foldername>/ harus berisi tepat 3 file logical: README.md, up.sql, down.sql. Pencocokan nama file bersifat case-insensitive agar hasil identik pada Windows dan Linux, tetapi tiap nama logical wajib muncul tepat sekali; duplikat case-only dan file tambahan tetap fail-hard.
Patch tidak boleh diedit setelah sudah ada patch migration yang di-generate.

14.2 Canonicalize patch_name (untuk filename migration)

Algoritma canonicalize:

trim whitespace

whitespace => _

lowercase

karakter selain [a-z0-9_-] => _

collapse _ berturut-turut

trim _ dan -

hasil kosong => FAIL


14.3 Patch migration generation

Engine membuat:

{version}_patch_<patch_name>_v{k}.up.sql

{version}_patch_<patch_name>_v{k}.down.sql


Header wajib:

/*
patch_folder: <foldername_asli>
patch_name: <patch_name_canonical>
readme: <ringkas atau path>
*/

Patch body disisipkan as-is (kecuali normalize LF + trim trailing spaces), dan dibungkus audit wrapper otomatis.

14.4 Patch duplication rule (once)

Jika sudah ada patch migration dengan patch_name yang sama di migrations/, engine SKIP generate (tidak compare isi patch).
Untuk membuat ulang patch, developer harus menghapus patch migration lama secara eksplisit.

14.5 Patch source must not disappear (integrity)

Jika patch migration sudah ada, tetapi folder patch/<foldername_asli>/ hilang atau file kurang => FAIL hard.

Selain itu, engine wajib menulis mapping patch di out/output.json (pass item):

patch_folder_original

patch_name_canonical

generated_migration_versions[]



---

15) Output Generator out/output.json

15.1 Bukan source-of-truth

out/output.json tidak boleh di-commit. Ini hanya artifact bantu saat generate/verify.

Semua fail item juga wajib dicetak lengkap ke console CLI/CI (rule, severity, file, object/location jika ada, reason, action, message, dan suggestion jika ada). Ringkasan hasil dan detail fail wajib memakai stream output yang sama agar collector GitLab tidak mengacak atau menyisipkan baris saat menggabungkan stdout/stderr. Artifact JSON tetap canonical, tetapi akses artifact tidak boleh menjadi syarat untuk memahami penyebab pipeline gagal.

15.2 Format wajib

Root harus memisahkan fail dan pass:

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

Item minimal:

rule, severity (0..3), file, message


Recommended:

object, location, suggestion, reason, action


Jika fail.count > 0 => exit code non-zero (lihat FAIL rule).


---

16) Locking: 1 migrate run per DB

Harus ada maksimal 1 migrate run per DB.
Preferred: sp_getapplock.

Template lock (SQL) — disarankan dieksekusi sekali di awal job:

EXEC dbo.usp_mig_try_get_lock
  @resource = N'rza_migrate_db:' + DB_NAME(),
  @timeout_ms = 60000;

Jika lock tidak didapat => fail + THROW.


---

17) MVP Supported Subset

Jika di luar scope ini, engine FAIL (bukan warn/skip).

17.1 Tables (MVP)

Supported:

basic scalar types (int/bigint/smallint/tinyint/bit, decimal/numeric, money/smallmoney, float/real, char/varchar/nchar/nvarchar, date/datetime/datetime2/smalldatetime, uniqueidentifier, varbinary, xml)

NULL/NOT NULL (omission treated as NULL)

constraints: PK, UQ, DF, CK, FK

computed columns: `<name> AS <expression> [PERSISTED]` (perubahan definisi dibuat ulang dengan drop/add)


Not supported (FAIL):

sparse, temporal, partition/filegroup directives, advanced features.


17.2 Indexes (MVP)

Supported:

CREATE [UNIQUE] INDEX ... ON dbo.Table(...) optional INCLUDE (...)

multiple index statements per table file


Not supported (FAIL):

columnstore/XML/spatial/indexed view/data compression.

filtered index (WHERE ...) termasuk supported subset.


17.3 Programmability (MVP)

Supported:

snapshot text view/function/procedure

header wajib CREATE OR ALTER


Rule tambahan:

1 file programmability = 1 object saja. Tidak boleh mendefinisikan object lain.



---

18) ALTER COLUMN Pre-check Templates (Generator Map)

Jika unsafe => SKIP (bukan fail) + log severity=2 + suggestion manual.

Parser tipe kolom wajib mengonsumsi deklarasi parameter lengkap dan
menormalisasi whitespace non-semantik. `DECIMAL(17,5)`, `DECIMAL(17, 5)`, dan
`DECIMAL (17,5)` adalah representasi state yang sama; perubahan menjadi
`DECIMAL(17,9)` adalah delta schema nyata. Aturan yang sama berlaku untuk
length/precision pada tipe parameterized lain seperti `VARCHAR`, `NVARCHAR`,
`BINARY`, `VARBINARY`, `TIME`, `DATETIME2`, dan `DATETIMEOFFSET`. Delta tersebut
tidak boleh hilang pada semantic hash atau diff dan harus mencapai generator
`ALTER COLUMN` atau policy manual-only yang sesuai.

Delimiter identifier tunggal pada tipe bawaan bersifat style-only:
`[DECIMAL](25,2)` dan `DECIMAL(25,2)` wajib membentuk base type semantik yang
sama. Delimiter tidak boleh menyebabkan false manual classification atau false
migration, sedangkan perubahan parameternya tetap merupakan delta nyata dan
wajib melalui guarded `ALTER COLUMN` bila keluarga transisinya didukung.

Placeholder: <table> <col> <new_type> <target_type> <nullability> <new_len> <p_new> <s_new>

18.1 NULL -> NOT NULL

DECLARE @null_cnt BIGINT;
SELECT @null_cnt = COUNT_BIG(1)
FROM dbo.<table>
WHERE <col> IS NULL;

IF (@null_cnt > 0)
BEGIN
  EXEC dbo.usp_mig_log_event @run_key, <version>, N'tables', N'<file>', N'skip', 2,
    N'alter_column', N'dbo.<table>.<col>', N'REQUIRES_MANUAL',
    N'[PRECHECK] ALTER COLUMN NOT NULL blocked',
    CONCAT(N'null_cnt=', @null_cnt),
    NULL, NULL, NULL;
  RETURN;
END

ALTER TABLE dbo.<table>
ALTER COLUMN <col> <new_type> NOT NULL;

18.2 Shrink length (VARCHAR/CHAR)

DECLARE @too_long BIGINT;
SELECT @too_long = COUNT_BIG(1)
FROM dbo.<table>
WHERE <col> IS NOT NULL
  AND DATALENGTH(<col>) > <new_len>;

IF (@too_long > 0)
BEGIN
  EXEC dbo.usp_mig_log_event @run_key, <version>, N'tables', N'<file>', N'skip', 2,
    N'alter_column', N'dbo.<table>.<col>', N'NARROWING_WOULD_TRUNCATE',
    N'[PRECHECK] SHRINK length blocked',
    CONCAT(N'too_long=', @too_long),
    NULL, NULL, NULL;
  RETURN;
END

ALTER TABLE dbo.<table>
ALTER COLUMN <col> <new_type> <nullability>;

18.3 Shrink length (NVARCHAR/NCHAR)

DECLARE @too_long BIGINT;
SELECT @too_long = COUNT_BIG(1)
FROM dbo.<table>
WHERE <col> IS NOT NULL
  AND DATALENGTH(<col>) > (<new_len> * 2);

IF (@too_long > 0)
BEGIN
  EXEC dbo.usp_mig_log_event @run_key, <version>, N'tables', N'<file>', N'skip', 2,
    N'alter_column', N'dbo.<table>.<col>', N'NARROWING_WOULD_TRUNCATE',
    N'[PRECHECK] SHRINK length blocked',
    CONCAT(N'too_long=', @too_long),
    NULL, NULL, NULL;
  RETURN;
END

ALTER TABLE dbo.<table>
ALTER COLUMN <col> <new_type> <nullability>;

18.4 Type -> numeric (TRY_CONVERT)

DECLARE @bad BIGINT;
SELECT @bad = COUNT_BIG(1)
FROM dbo.<table>
WHERE <col> IS NOT NULL
  AND TRY_CONVERT(<target_type>, <col>) IS NULL;

IF (@bad > 0)
BEGIN
  EXEC dbo.usp_mig_log_event @run_key, <version>, N'tables', N'<file>', N'skip', 2,
    N'alter_column', N'dbo.<table>.<col>', N'REQUIRES_MANUAL',
    N'[PRECHECK] TYPE change blocked (non-convertible)',
    CONCAT(N'bad_cnt=', @bad),
    NULL, NULL, NULL;
  RETURN;
END

ALTER TABLE dbo.<table>
ALTER COLUMN <col> <target_type> <nullability>;

18.5 Dependency blocking rule for ALTER COLUMN (MVP)

Jika ALTER COLUMN memerlukan drop index/constraint untuk bisa jalan:

generator tidak boleh auto-drop dependency untuk meloloskan alter (di MVP)

generator harus SKIP + warn reason REQUIRES_MANUAL

detail harus mencantumkan dependency yang menghalangi.

Jika dependency penghalang tidak bernama (unnamed) sehingga drop/recreate plan tidak deterministik, engine wajib fail lebih awal pada source validation dengan rule `UNNAMED_CONSTRAINT_NOT_ALLOWED` dan reason `REQUIRES_MANUAL`.



---

19) constraints_naming.md (Opsional)

Dokumen ini menjadi guideline penamaan constraint lintas tim.
Semua constraint pada `db_object/tables/*.sql` wajib diberi nama eksplisit memakai `CONSTRAINT <name>`.
Engine memberi nama deterministik secara otomatis saat table, jenis constraint, dan target dapat dibuktikan aman. Bentuk ambigu tetap fail hard dengan rule `UNNAMED_CONSTRAINT_NOT_ALLOWED` karena drop/recreate plan tidak deterministik.
Tidak ada prefix nama wajib; `PK_`, `DF_`, `UQ_`, `CK_`, dan `FK_` hanya pola auto-fix/rekomendasi.
File `constraints_naming.md` opsional. Bootstrap boleh membuat guideline default, sedangkan checker hanya memberi warning jika file tidak ada.


---

20) Repo Hygiene (.gitignore & .gitattributes)

File repo hygiene direkomendasikan, dapat dibuat otomatis oleh bootstrap jika belum ada, dan tidak menjadi blocking rule bagi validitas object SQL.

20.1 .gitignore

Wajib ignore:

out/**

*.output.json


20.2 .gitattributes (LF)

Wajib menormalkan EOL ke LF:

* text=auto eol=lf
*.sql text eol=lf
*.md  text eol=lf
*.go  text eol=lf


---

21) CLI Contract (rza_migrate_db)

Commands minimum:

rza_migrate_db fmt

auto-format db_object/** in-place

comment dan multiline literal adalah protected spans; posisi/indentasi comment tidak boleh diubah

rza_migrate_db fmt -check

preview perubahan format tanpa menulis db_object/**

tulis out/output.json


rza_migrate_db gen

flow: fmt -> validate -> reconstruct history(migrations/) -> diff -> emit migrations -> verify -> out/output.json


rza_migrate_db verify

validate only, no emit

tulis out/output.json



Exit codes:

0: success (fail.count==0)

1: failed (fail.count>0)



---

22) Engine Config & Codebase Skeleton

22.1 app_config.go (semua config terkunci di sini)

Minimum config:

Paths:

DbObjectDir (default db_object)

MigrationsDir (default migrations)

PatchDir (default patch)

OutDir (default out)


Filename/versioning:

PairSeqWidth (=3)

MaxPairsPerRun (=999)

AllowedCategories (=types,tables,indexes,functions,views,procedures,patch)

AllowedObjectTokenCharset (=A-Za-z0-9_.# untuk filename object)


Formatting:

LineEnding (=LF)

TrimTrailingSpaces (=true)

EnsureFinalNewline (=true)


Parsing:

FailHardOnParseError (=true)

NoRegexParsing (=true; structured/AST untuk parser utama rza_migrate_db)


Safety:

ForbiddenSqlTokens (GO, USE, CREATE DATABASE, ALTER DATABASE, CREATE LOGIN, ALTER LOGIN, CREATE USER)

Strict mode:

EnableStrictMode (=false default)

Jika false: finding `MIGRATION_PRECHECK_MISSING` dicatat sebagai warning (non-fatal).

Jika true: finding `MIGRATION_PRECHECK_MISSING` dinaikkan menjadi fail-hard (pipeline stop).


Ordering:

GenerationOrder (=types,tables,indexes,functions,views,procedures,patch)

Programmability source contract mode:

EnableProgrammabilitySourceContractForNewOnly (=false default)

Jika true: validasi source contract programmability yang ketat diterapkan untuk object baru; object existing dari history tetap diverify header `CREATE OR ALTER` tetapi tidak di-hard-fail oleh source-contract parser tambahan.



22.2 Skeleton foldering code (disarankan)

cmd/rza_migrate_db/
  main.go
internal/
  config/     # load/app_config.go, defaults
  fsx/        # file walk, stable ordering
  sqlfmt/     # formatter canonicalization
  sqlparse/   # structured parse (subset MVP)
  history/    # parse migrations, reconstruct state
  diff/       # semantic diff (db_object vs state)
  emit/       # write new migration pairs
  output/     # output.json writer
  verify/     # validate migrations folder, pairs, content match

22.3 Verify rules (migrations folder integrity)

Engine verify wajib memastikan:

Semua migration punya pasangan up/down.

Tidak ada duplikat version untuk dua pair yang berbeda (pair-safe).

Tidak ada file dengan category di luar whitelist.

Tidak ada token forbidden (GO, USE, security DDL) di migrations.

Setiap file tables/indexes/patch memiliki wrapper + audit call minimal.

Untuk programmability: wajib CREATE OR ALTER, dan 1 file = 1 object.

migrations/ wajib flat (tanpa subfolder).



---

23) Final Locks / Gaps Closed

1. Schema hanya dbo.


2. db_object/ source-of-truth; migrations/ history eksekusi.


3. No GO.


4. Forbidden SQL: USE/DB/security DDL.


5. No regex: parser utama rza_migrate_db (filename + verify/diff/emit) wajib structured/AST.

   Catatan: rza_get_object_db saat ini masih regex-assisted khusus normalisasi header sys.sql_modules.


6. Up/down harus berbagi version yang sama persis.


7. Patch: only 3 files; wrapper audit auto; immutable; source patch tidak boleh hilang.


8. Ordering generate: types → tables → indexes → functions → views → procedures → patch; dependency graph dapat mendahului urutan dasar.


9. DROP TABLE manual-only: generator buat request trace, runtime skip + warn.


10. DROP COLUMN auto: drop dependency dulu; tidak auto recreate index; follow-up required.


11. ALTER COLUMN: precheck; aman auto; tidak aman skip + warn.


12. Tables/indexes/patch wajib TRY/TRAN; programmability cukup CREATE OR ALTER.


13. Audit schema normalized: tidak ada env/db_name/tool_*; semua timestamp GETDATE(); run_key=NEWID() per file.


14. out/output.json tidak di-commit.


15. MVP subset: di luar scope => FAIL.


16. Skip adalah policy-only (bukan recovery runtime error).


17. ALTER COLUMN tidak auto-drop dependency untuk meloloskan alter (MVP).


18. Programmability file = 1 object only.


19. migrations/ wajib 1 folder flat (tanpa subfolder).


20. FAIL rule explicit: severity=3 / fail.count>0 => exit 1, stop, no emit.




---

24) Contoh Fail/Pass Output & Error Patterns

24.1 Contoh fail item: file vs content mismatch

Kasus:

file: db_object/tables/dbo.TableA.sql

isi: CREATE TABLE dbo.TableB (...)


Output item:

{
  "rule": "FILE_OBJECT_MISMATCH",
  "severity": 3,
  "file": "db_object/tables/dbo.TableA.sql",
  "object": "dbo.TableA",
  "message": "Filename defines dbo.TableA but SQL defines dbo.TableB",
  "suggestion": "Rename file to dbo.TableB.sql or fix CREATE TABLE header to dbo.TableA",
  "reason": "REQUIRES_MANUAL",
  "action": "FAIL"
}

24.2 Contoh fail item: inline DEFAULT

{
  "rule": "INLINE_DEFAULT_NOT_ALLOWED",
  "severity": 3,
  "file": "db_object/tables/dbo.TableA.sql",
  "object": "dbo.TableA.status",
  "message": "Inline DEFAULT(...) is forbidden in db_object table snapshot",
  "suggestion": "Replace with table-level default constraint (non-inline), example: CONSTRAINT DF_TableA__status DEFAULT('NEW') FOR status",
  "reason": "INLINE_DEFAULT_NOT_ALLOWED",
  "action": "FAIL"
}

24.3 Contoh warn item: drop column dropped indexes (follow-up)

{
  "rule": "DROP_COLUMN_DROPS_INDEXES",
  "severity": 2,
  "file": "migrations/20260217073015022_tables_dbo.TableA_v3.up.sql",
  "object": "dbo.TableA.amount",
  "message": "Dropping column requires dropping impacted indexes (not auto recreated)",
  "suggestion": "Define new indexes explicitly in db_object/indexes/dbo.TableA.sql",
  "reason": "INDEX_RECREATE_REQUIRED",
  "action": "warn"
}


---

25) Runner Cheatsheet (golang-migrate)

Contoh (sesuaikan driver/DSN):

# Up semua yang belum jalan
migrate -path ./migrations -database "$MSSQL_DSN" up

# Down 1 step
migrate -path ./migrations -database "$MSSQL_DSN" down 1

# Down sampai version tertentu (goto = set version persis)
migrate -path ./migrations -database "$MSSQL_DSN" goto 20260217073015000

# Force version (HATI-HATI: hanya untuk memperbaiki dirty state)
migrate -path ./migrations -database "$MSSQL_DSN" force 20260217073015000

# Check version / dirty
migrate -path ./migrations -database "$MSSQL_DSN" version

Catatan:

golang-migrate berhenti saat ada error SQL (THROW), sehingga urutan berikutnya tidak dieksekusi.

Pairing up/down bergantung pada version yang sama persis.

---

26) Migration History Reconstruction Report

`rza_migrate_report` hanya boleh membaca `migrations/` dan source desired-state `db_object/`; tool dilarang mengakses database, mengeksekusi SQL, atau membuat/mengubah/menghapus file di dalam `migrations/` maupun `db_object/`. Dokumentasi relasi antardata-object hanya boleh memakai FK yang dideklarasikan atau static equality predicate dari explicit `JOIN ... ON` maupun legacy `FROM A, B WHERE ...` pada view/procedure yang memetakan tepat satu resolved `alias.column` pada setiap sisi. Evidence observed wajib dibedakan dari FK dan tidak boleh mengklaim cardinality atau referential integrity. Nama kolom sama, operator non-equality, ekspresi ambigu, temporary/table variable, alias lintas query scope, dan dynamic SQL tidak boleh ditebak sebagai relasi. Dependency module-ke-object wajib dilaporkan terpisah sebagai lineage, termasuk daftar module tanpa resolved static reference untuk audit coverage.

Pengenalan object procedure wajib berbasis token/source position dan statement scope, bukan physical line. Whitespace atau line break antara keyword dan object tidak boleh mengubah hasil. CTE unqualified hanya berstatus `NOT_PERSISTED` di statement yang mendeklarasikannya dan tidak boleh menutupi object schema-qualified pada statement berikutnya. Alias DML hanya boleh di-resolve dari source dalam statement yang sama. `SELECT INTO`, assignment `SELECT`, subquery/`EXISTS`, dan `INSERT ... SELECT` bukan direct client result set.

Alias DML juga wajib berada dalam query scope yang sama; alias nested subquery atau kandidat ambigu tidak boleh dipilih. Quoted identifier bukan keyword. Unterminated identifier tidak boleh memicu panic atau ditebak sebagai nama lengkap. Setiap occurrence akses/static EXEC mempertahankan source offset, termasuk occurrence pada baris yang sama. `MERGE WHEN ... THEN` action bukan target DML baru; SELECT/OUTPUT INTO adalah sink write, OUTPUT tanpa INTO adalah client result. CTE case variant dengan collation yang belum terbukti tetap ambigu pada report; dependency engine mempertahankan kemungkinan physical reference secara exact-case. Exact write mapping dilarang menyeberang statement atau mengandung FROM sebagai assignment expression; multirow/compound projection yang belum dimodelkan tidak boleh diberi mapping yang seolah lengkap.

Procedure impact profile wajib memakai inventory table/view/function/procedure dari `db_object` sebagai validation catalog, bukan sebagai daftar kandidat untuk menebak nilai runtime. Static, reconstructed literal dynamic, literal-from-caller, external database, unresolved dynamic, temporary/non-persisted, missing, ambiguous, dan system object wajib mempunyai resolution code berbeda. Case/schema collision tidak boleh dipilih otomatis. Cursor, built-in function, dan internal `SELECT` assignment/`EXISTS`/`INSERT ... SELECT` tidak boleh dilaporkan sebagai table/UDF/direct-result-set yang salah. Bentuk `EXEC @return_code = schema.Procedure` wajib diperlakukan sebagai static procedure call, bukan variable dynamic SQL; `EXEC @sql` tanpa target procedure static tetap masuk dynamic review. Dependency lineage identik wajib di-deduplicate case-insensitive, sementara collision catalog yang benar-benar ambigu tetap dilaporkan. Local execution flow wajib memakai orientasi top-down dan hanya boleh menyatakan pola source yang terobservasi: guarded event wajib bercabang dan bergabung kembali, `GOTO ERROR` wajib dipisahkan dari normal path, serta transaction hanya boleh diringkas bila tidak menyembunyikan child procedure call. Local object map wajib menjelaskan jenis akses dan tidak boleh diklaim sebagai execution order atau relasi FK. Guard `IF`/`ELSE` wajib mempertahankan literal dari source asli pada audit table; versi diagram boleh dipersingkat tanpa mengubah outcome branch. Line break di dalam node/edge procedure wajib memakai sintaks yang dirender Mermaid dan tidak boleh muncul sebagai teks literal `\n`. Mermaid edge label yang memuat SQL wajib dikutip agar operator SQL tidak merusak rendering. Recursive procedure traversal wajib cycle-safe, mempertahankan semua direct call edge, memecah diagram besar tanpa menghapus edge, dan mencegah ekspansi subgraph identik berulang. Exact write mapping dan filter hanya boleh ditampilkan bila target column/expression dapat dipisahkan secara konservatif; key/index context tidak boleh mengklaim single-row update tanpa seluruh enforced unique-key columns pada filter static. Link antar-generated Markdown wajib URL-safe tanpa mengubah filename sumber.

Dokumentasi identifier wajib membedakan enforcement uniqueness dari generation dan review signal: PK, UNIQUE constraint, dan unfiltered unique index adalah enforcement; filtered unique index bersifat conditional; `IDENTITY` hanya menghasilkan nilai; nama ID-like, pemakaian JOIN, dan non-unique index hanya review signal. Report dilarang menyebut signal tersebut sebagai unique-key candidate atau merekomendasikan perubahan schema tanpa validasi duplicate data, nullability, business scope, dan kebutuhan composite key.

Listing suspected JOIN identifiers wajib menerima kolom dari observed static equality JOIN tanpa mensyaratkan nama ID-like, index, atau IDENTITY, dan tidak boleh tersembunyi karena table memiliki PK lain. Evidence wajib menyatakan mapping kolom, source modules, occurrence, index/key serta IDENTITY; canonical mapping bukan salinan expression SQL. Individually unfiltered unique columns tetap di inventory enforcement; filtered/composite key tidak boleh diperlakukan sebagai global single-column uniqueness. Declared FK saja tidak boleh meningkatkan Join usage; per occurrence self-join pada kolom yang sama dihitung sekali. Listing tidak mengubah source, rule checker, migration planning, atau history.

Narasi dan panduan Markdown desired-state wajib memakai Bahasa Indonesia yang natural; istilah/header database yang lebih jelas dalam English (`Classification`, `Uniqueness`, `Source`, `Join usage`, `Evidence`, `Confidence`, `Relationship Group`, `Relational Map`) tidak boleh diterjemahkan secara literal. Stable machine/audit codes tetap dipertahankan. Guidance yang hanya bergantung pada classification wajib ditulis sekali pada reference section dan tidak boleh diulang seolah-olah merupakan risk assessment unik per table. Raw HTML seperti `<br>` dilarang. Seluruh Mermaid wajib berada dalam satu `relational_erd.md` dan disusun sebagai Relationship Group → Relational Map. Setiap map wajib connected, dibentuk adaptif, dan dibatasi maksimum 20 object, 24 relationship lines, 35 exact mappings, serta 9.000 karakter Mermaid. Map hanya boleh menggambar topology yang benar-benar ditemukan; writer dilarang memaksakan satu object sebagai Hub/parent. Node wajib menampilkan type object dan status key/index bagi kolom relationship. Banyak exact column mappings pada object pair yang sama digabung pada label garis sebagai `left_column = right_column`; `[FK]`/`[JOIN]` hanya wajib pada pasangan yang mempunyai mixed evidence. Statistik occurrence/usage tidak boleh menggantikan mapping tersebut. Setiap non-self object pair wajib tampil dan setiap exact mapping wajib muncul tepat sekali pada edge label. Dense pair yang melampaui mapping/text budget wajib dipecah menjadi numbered `Mapping batch`, bukan dipotong atau disembunyikan. Object index dan shared object navigation boleh mengulang nama object untuk navigasi, tetapi tidak boleh dinyatakan sebagai atribut database. Seluruh exact mapping, termasuk self-join, wajib tersedia tepat sekali sebagai data row pada `relationship_catalog.md` dengan Group/Map yang sama dan maksimal 35 row per tabel. Self-join wajib dikeluarkan dari diagram dan dilaporkan pada `Self-join review`. Generator wajib membersihkan stale `relationship_diagrams/` dari format lama.

Parser filename wajib memakai kontrak shared `internal/history`: core pair canonical, version migration 17 digit, category whitelist, object non-patch `dbo.<ObjectName>`, dan token `v{k}` dengan `k >= 1`.

Output report wajib berada di luar `migrations/`. Pair kosong/searah atau down yang tidak memuat inverse semantik harus terlihat pada `round_trip_status`; engine tidak boleh memberi `PASS` hanya karena file pasangan ada. SQL yang tidak dapat dipastikan tetap dilaporkan sebagai `NOT_TESTED`/`LOW` dan `MANUAL_REVIEW_REQUIRED`, bukan ditebak.

Path relatif `-out` dan `-docs` wajib di-resolve terhadap repository root. `report.json` wajib menyatakan kedua resolved root dan setiap generated artifact konkret; wildcard hanya boleh dipertahankan pada field compatibility lama. JSON audit baru dipublikasikan setelah desired-state documentation selesai ditulis agar manifest tidak mengklaim dokumen yang gagal dibuat.

Pair bootstrap canonical `000000__core__mig__audit_init` adalah infrastruktur audit dan dikecualikan dari analisa history aplikasi (`migration_summary`, object history, dan semantic change rows). File tetap dipindai read-only sebagai bagian release, dan tidak boleh dihapus atau dimutasi oleh report. Ringkasan CSV/JSON dan console wajib melaporkan jumlah pair yang dikecualikan secara eksplisit agar `total_files`, analyzed pairs, dan exclusion dapat direkonsiliasi tanpa asumsi tersembunyi.

Perubahan metadata saja (komentar, whitespace, dan rename constraint dengan definisi identik) tidak boleh menambah operasi semantik. Metadata waktu eksekusi boleh berubah, tetapi ordering dan seluruh data hasil rekonstruksi harus deterministik untuk source yang sama.

Parser wajib memproses seluruh item top-level dalam satu statement `ALTER TABLE`, termasuk multi-column, multi-constraint, campuran column+constraint, composite key/FK, `WITH CHECK`/`WITH NOCHECK`, serta multiple drop. Urutan statement di file adalah urutan perubahan state; parser dilarang mengelompokkan CREATE/DROP/ALTER berdasarkan jenis jika itu mengubah hasil rekonstruksi. Drop+recreate constraint/index bernama sama harus dilaporkan sebagai satu transisi `ALTER_*` dengan before/after, sedangkan inverse down harus lengkap untuk setiap operasi dalam migration.

State rekonstruksi wajib berlanjut melintasi semua version object, bukan hanya membandingkan dua file terakhir. `DROP TABLE` dan `CREATE TABLE` sesudah penghapusan wajib mengosongkan state kolom, constraint, dan index terdahulu; versi setelah reintroduksi hanya boleh memakai state hasil create terbaru. Kontrak ini harus dibuktikan minimal dengan rantai v1 sampai v7 dan tes negatif referensi atribut pra-drop.




