# Guarded Automatic Migration Guide

`db_object/` adalah source of truth untuk desired schema. `migrations/` adalah source of truth untuk history perubahan. Engine tidak membuat manifest, state file, atau source SQL ketiga. Semua model/risk/cache yang muncul di `out/` hanyalah artefak turunan dan boleh dihapus.

## Penempatan file dan kesalahan input

Satu file hanya boleh mendefinisikan kategori object sesuai foldernya:

| Isi SQL | Lokasi yang benar | Contoh nama |
| --- | --- | --- |
| `CREATE TABLE` beserta named default/check/PK/UQ/FK | `db_object/tables/` | `dbo.Customer.sql` |
| Satu atau lebih `CREATE [UNIQUE] [CLUSTERED|NONCLUSTERED] INDEX` untuk satu table | `db_object/indexes/` | `dbo.Customer.sql` |
| `CREATE OR ALTER VIEW` | `db_object/views/` | `dbo.ActiveCustomer.sql` |
| `CREATE OR ALTER FUNCTION` | `db_object/functions/` | `dbo.NormalizeName.sql` |
| `CREATE OR ALTER PROCEDURE` | `db_object/procedures/` | `dbo.SaveCustomer.sql` |

Rule `DB_OBJECT_WRONG_CATEGORY` berarti header SQL dan folder berbeda. Contoh yang salah adalah `CREATE INDEX IX_Customer_Email ...` di `db_object/tables/dbo.Customer.sql`; pindahkan statement tersebut ke `db_object/indexes/dbo.Customer.sql`. Jangan menyalin statement ke dua folder karena akan menciptakan dua desired state yang saling bersaing.

File table harus berisi tepat satu snapshot `CREATE TABLE`, bukan `ALTER`, `DROP`, wrapper `IF`, atau index. File index hanya boleh berisi index yang seluruhnya menargetkan table pada nama file. File programmability harus berisi tepat satu object dengan header `CREATE OR ALTER`; DDL temporary table di body procedure tetap diperbolehkan.

Named default boleh ditulis dengan dua bentuk yang dinormalisasi ke model semantic yang sama:

```sql
-- Inline
[Status] VARCHAR(20) NOT NULL CONSTRAINT [DF_Customer_Status] DEFAULT ('ACTIVE')

-- Table-level
[Status] VARCHAR(20) NOT NULL,
CONSTRAINT [DF_Customer_Status] DEFAULT ('ACTIVE') FOR [Status]
```

Nama constraint wajib ada. Kolom pada `FOR` wajib ada dan satu column tidak boleh mempunyai lebih dari satu default.

Extractor tidak lagi menghapus snapshot lama bila object masih ada di SQL Server tetapi tidak dapat dirender, misalnya module terenkripsi atau table dengan fitur yang belum didukung. Warning tetap wajib diselesaikan: snapshot yang dipertahankan adalah last-known-good, bukan bukti bahwa source sudah sinkron dengan database.

## Cara membaca hasil validasi

| Rule | Arti | Tindakan user |
| --- | --- | --- |
| `GUARDED_AUTO_MIGRATION` | Engine dapat membuat pair otomatis, tetapi script akan memeriksa data saat `up` dan `down`. | Review SQL generated; jalankan di database copy/QA terlebih dahulu. Jika precheck gagal, perbaiki data atau buat pair manual di `migrations/`. |
| `MANUAL_MIGRATION_REQUIRED` | Engine tidak dapat membuktikan `up` dan `down` aman atau reversible. | Tetap ubah `db_object` ke target desired state, lalu tulis pair reviewed langsung di `migrations/`. |
| `UNSUPPORTED_SQL_FEATURE` | Fitur belum punya typed model + renderer + round-trip + SQL Server test. | Jangan menyederhanakan object secara paksa. Tulis pair manual di `migrations/`, atau perluas capability engine dengan fixture dan test. |
| `PLAN_ROUND_TRIP_FAILED` | Rekonstruksi semantic dari generated `up` tidak sama dengan `db_object`. | Jangan commit generated files. Perbaiki model/renderer atau gunakan pair manual. |
| `DB_OBJECT_WRONG_CATEGORY` | Jenis object SQL tidak sama dengan folder. | Pindahkan file/statement ke folder pada tabel penempatan di atas, lalu jalankan `verify_then_gen` kembali. |

Precheck yang menemukan data tidak kompatibel selalu memakai `THROW`; transaction wrapper melakukan rollback. Migration tidak boleh sukses hanya karena statement `ALTER COLUMN` di-skip.

## Perubahan table yang digenerate otomatis

| Perubahan | Mode | Jaminan |
| --- | --- | --- |
| Tambah nullable column | otomatis | `down` drop column, schema-only rollback. |
| Tambah NOT NULL column dengan named default | otomatis | `up` memakai `WITH VALUES`; `down` menghapus default lalu column. |
| Tambah computed column | otomatis | `ALTER TABLE ADD [col] AS (...)`; dependensi baru tetap harus valid pada SQL Server. |
| Tambah identity column | otomatis | Seed/increment dari source dipertahankan; perubahan identity existing ditolak dengan remediation. |
| Tambah `UNIQUEIDENTIFIER ROWGUIDCOL` | otomatis | Source salah tipe ditolak sebelum generator membuat SQL. |
| Ubah default named constraint | otomatis | `up/down` drop dan add definisi named constraint. |
| Ubah named CHECK constraint | otomatis terbatas | Definisi lengkap harus dapat diparse dan direkonstruksi. |
| Tambah/ubah/hapus named UNIQUE | guarded otomatis | Precheck duplicate tuple pada up dan down. |
| Tambah/ubah/hapus named PRIMARY KEY | guarded otomatis | Precheck NULL dan duplicate tuple pada up dan down. |
| Widen string dalam tipe/encoding yang sama | otomatis | Tidak memerlukan data-loss guard. |
| Shrink string | guarded otomatis | `DATALENGTH` harus muat target pada `up` dan kembali muat saat `down`. |
| Ubah precision/scale/length/date-time precision | otomatis atau guarded sesuai risiko | Seluruh parameter tipe dipertahankan walaupun source memakai spasi di sekitar kurung/koma; `up` memakai deklarasi target dan `down` memakai deklarasi history. |
| `NULL` ke `NOT NULL` | guarded otomatis | Tidak boleh ada nilai NULL. Jika ada, transaction rollback. |
| Numeric/date-time conversion | guarded otomatis | `TRY_CONVERT` dan round-trip conversion harus lossless. |
| Add/drop/ubah index subset supported | otomatis | Index direkonstruksi dari source index yang sama. |
| XML/spatial/columnstore dan uniform compression | otomatis terbatas | Definisi lama disimpan history untuk down; XML dependency diurutkan dan hasil extractor diuji melalui catalog introspection. |
| Explicit column rename | guarded otomatis | Gunakan `-- rza:rename-column OldName -> NewName`; engine membuat `sp_rename` up/down dan drop/recreate dependency bernama. |
| Drop/rename column/table, computed transition, perubahan FK/PK/UQ kompleks | manual | Tidak ada jaminan data-safe down atau dependency graph lengkap. |

## Contoh guarded automatic transition

Desired state di `db_object/tables/dbo.Customer.sql` berubah dari:

```sql
[Name] NVARCHAR(200) NULL
```

menjadi:

```sql
[Name] NVARCHAR(100) NOT NULL
```

Engine membuat migration dengan dua precheck:

```sql
SELECT @too_long = COUNT_BIG(1)
FROM dbo.Customer
WHERE [Name] IS NOT NULL AND DATALENGTH([Name]) > (100 * 2);

SELECT @null_cnt = COUNT_BIG(1)
FROM dbo.Customer
WHERE [Name] IS NULL;
```

Jika salah satu count lebih besar dari nol, engine menjalankan `THROW 51000`; wrapper rollback dan version migration tidak boleh dianggap applied.

Tindakan user bila gagal:

```sql
-- Contoh data remediation yang direview terpisah sebelum menjalankan migration.
UPDATE dbo.Customer
SET [Name] = N'UNKNOWN'
WHERE [Name] IS NULL;

SELECT [Name]
FROM dbo.Customer
WHERE DATALENGTH([Name]) > 200;
```

Jangan mengubah generated migration untuk menyembunyikan precheck.

## Contoh UNIQUE dan PRIMARY KEY guarded

Source berikut dapat dibuat otomatis bila constraint selalu diberi nama:

```sql
CREATE TABLE dbo.Customer
(
  Id INT NOT NULL,
  Email NVARCHAR(320) NULL,
  CONSTRAINT UQ_Customer_Email UNIQUE (Email)
);
```

Sebelum `ALTER TABLE ... ADD CONSTRAINT`, engine memeriksa:

```sql
SELECT 1
FROM dbo.Customer
GROUP BY [Email]
HAVING COUNT_BIG(*) > 1;
```

Untuk primary key, engine juga memeriksa setiap key column tidak bernilai `NULL`.
Jika ada duplicate atau NULL, migration melempar `THROW` dan wrapper rollback.

Jika validasi gagal, user dapat melakukan remediation eksplisit sebelum menjalankan migration:

```sql
SELECT Email, COUNT_BIG(*) AS DuplicateCount
FROM dbo.Customer
GROUP BY Email
HAVING COUNT_BIG(*) > 1;
```

Named FK sekarang guarded otomatis bila parser dapat membaca local-column dan referenced-column, termasuk composite key. Precheck memperlakukan tuple dengan minimal satu local `NULL` sebagai valid—sesuai semantics FK SQL Server—dan menolak orphan row lainnya. `ON DELETE`/`ON UPDATE` tetap direplay dari definition source.

Saat named PK/UQ yang direferensikan berubah, engine mencari FK incoming pada history, sementara drop FK tersebut, mengubah key, lalu menambahkan kembali FK dengan orphan precheck. Prosedur kebalikannya dibentuk untuk down.

## Contoh numeric guarded transition

Perubahan `DECIMAL(18,4)` menjadi `DECIMAL(10,2)` hanya dieksekusi bila konversi target dan konversi balik menghasilkan nilai yang sama. Nilai seperti `12.3456` akan membuat precheck gagal karena pembulatan menjadi `12.35` tidak lossless.

Bentuk extractor `[DECIMAL](25,2) -> [DECIMAL](25,6)` masuk jalur guarded yang
sama. Walaupun scale bertambah, jumlah digit integer berkurang dari 23 menjadi
19; karena itu precheck `TRY_CONVERT` dan round-trip tetap diperlukan sebelum
`ALTER COLUMN` dijalankan. Data yang tidak muat menyebabkan `THROW` dan rollback.

Jika pembulatan memang keputusan bisnis, buat pair manual di `migrations/`:

```sql
-- up: reviewed business transformation
UPDATE dbo.Invoice
SET Amount = ROUND(Amount, 2)
WHERE Amount <> ROUND(Amount, 2);

ALTER TABLE dbo.Invoice ALTER COLUMN Amount DECIMAL(10,2) NOT NULL;
```

`down` harus menjelaskan keterbatasan data: memperlebar kembali ke `DECIMAL(18,4)` tidak mengembalikan digit yang sudah dibulatkan.

## Contoh perubahan yang wajib manual

### Drop column

Tetap update `db_object` dengan menghapus column, lalu buat pair eksplisit di `migrations/` setelah backup/data-retention plan tersedia.

```sql
-- up: confirm retention policy before destructive action
ALTER TABLE dbo.Customer DROP COLUMN LegacyCode;

-- down: schema only; original values cannot be restored automatically
ALTER TABLE dbo.Customer ADD LegacyCode VARCHAR(30) NULL;
```

### Rename column

Rename tidak boleh ditebak sebagai drop+add. Tuliskan directive tepat satu kali di file table desired state:

```sql
-- rza:rename-column OldName -> NewName
CREATE TABLE dbo.Customer
(
  Id INT NOT NULL,
  NewName NVARCHAR(200) NULL
);
```

Migration rename harus hanya berisi pasangan rename. Perubahan tipe, nullability, default, atau penambahan column lain dipisahkan ke perubahan berikutnya. Engine memvalidasi bahwa definisi old/new sama, target belum ada, dependency bernama dapat dipulihkan, dan membentuk operasi kebalikan untuk down.

```sql
-- up
EXEC sp_rename N'dbo.Customer.OldName', N'NewName', N'COLUMN';

-- down
EXEC sp_rename N'dbo.Customer.NewName', N'OldName', N'COLUMN';
```

Directive boleh dihapus setelah migration rename sudah menjadi history dan `verify_then_gen` telah konvergen; nama final tetap berada di `db_object`.

## Syarat menambah capability baru

Sebuah fitur SQL Server baru hanya boleh menjadi otomatis setelah memiliki seluruh bukti berikut:

1. typed model dan parser fail-closed;
2. renderer `up` dan `down` deterministic;
3. unit test parser/diff;
4. model round-trip test;
5. fixture input `db_object`;
6. SQL Server sandbox compile + introspection test;
7. entri pada capability matrix.

## Repository besar dan konvergensi generator

Generator sengaja membatasi jumlah pair per iterasi. Selama output masih menunjukkan `pending_pairs > 0`, jalankan `verify_then_gen` lagi. Proses selesai hanya ketika `fail=0`, `generated_pairs=0`, dan `pending_pairs=0`. Ini tidak membuat source of truth baru: setiap iterasi tetap merekonstruksi state dari `migrations/` lalu membandingkannya dengan `db_object/`.
