# Audit pembacaan object rza_migrate_report — 12 Oktober 2026

Audit memakai empat snapshot nyata dalam `all_object/`, bukan asumsi berdasarkan
nama kolom. Pembandingan dilakukan terhadap source revision `13e4eed` sebelum
perbaikan. Semua pemeriksaan bersifat offline; tidak ada koneksi atau eksekusi SQL.

## Corpus dan hasil observasi

| Snapshot | File object | Table | Module sebelum → sesudah | Observed mapping relasi sebelum → sesudah | Write mapping sebelum → sesudah |
| --- | ---: | ---: | ---: | ---: | ---: |
| bond/bond-sqltrsretail | 2.449 | 666 | 1.209 → 1.305 | 728 → 747 | 20.136 → 20.743 |
| bond/bond-sqlreplicate | 286 | 165 | 2 → 2 | 2 → 2 | 11 → 11 |
| reksa/reksa-sqlreksa2 | 1.654 | 468 | 902 → 965 | 647 → 675 | 13.463 → 13.643 |
| reksa/reksa-sqlreplicate | 123 | 65 | 2 → 2 | 0 → 0 | 0 → 0 |

Total: 4.512 file object, 1.364 table, 21.596 column, dan 2.274 module. Sebanyak
159 function sebelumnya tidak ikut consumer/lineage. Jumlah relasi adalah
mapping teragregasi dari source; kenaikan jumlah bukan bukti otomatis bahwa semua
hubungan valid secara bisnis. Perbaikan menghapus tebakan, mempertahankan scope,
dan menambah mapping yang terbukti pada regresi. Tidak ada klaim FK/cardinality,
row count, uniqueness data, atau bahwa procedure berhasil dijalankan di SQL Server.

## Bug yang diperbaiki dan bukti

| Temuan | Bukti nyata / reproduksi | Perbaikan |
| --- | --- | --- |
| OUTER JOIN salah menjadi INNER | `Rvmp_trs_ListSecurityTransactionForInsertEditBlockTrans.sql`, baris 60–72: LEFT OUTER JOIN pada DealId dan OrderId | Jenis JOIN melewati OUTER dan physical hint HASH/LOOP/MERGE/REMOTE; MERGE JOIN bukan statement DML. |
| Alias dan CTE salah dipetakan | Regresi nested query, CTE bernama sama dengan table, duplicate alias, query tanpa semicolon | Scope mengikuti shared access scanner, query depth, statement bounds dan catalog resolution; candidate ambigu tidak dipilih. Outer scope tidak hilang karena FROM nested. |
| Quoted keyword/escaped name dianggap keyword atau terpotong | Regresi `[WHERE]`, `[JOIN]`, `[a]]b]`, `[A.B]`, schema-qualified endpoint | Lexer bersama mempertahankan quoting dan multipart identity; table/column canonical sebelum agregasi. |
| Wrapper multi-argument kehilangan evidence | `ISNULL(a.Code,'')=b.Code` | Operand dibaca sebagai expression lengkap; literal/comment tetap opaque; operand multi-column atau subquery tidak ditebak. |
| Empat index hilang | `bond-sqlreplicate/db_object/indexes/dbo.JHCLDR.sql` dan `dbo.MXHolidayCurrency.sql`: ON pada baris berikutnya | Reader DDL token mengenali whitespace/newline, GO dan batas CREATE berikutnya; filter tidak bocor ke index sebelumnya. |
| PK/UQ/FK dan type dibaca dengan substring | Regresi comment berisi CREATE TABLE, quoted UNIQUE, precision/scale, inline IDENTITY/PK/FK, ALTER FK | `internal/sqlparse.ReadTableMetadata/ReadIndexMetadata` memakai lexer bersama, mempertahankan type declaration dan key columns/order. |
| UPDATE tanpa WHERE tidak mempunyai mapping | Tambahan mapping pada corpus Bond/Reksa dan regresi full-table UPDATE / UPDATE FROM | Assignment dipisahkan dari FROM/OUTPUT/OPTION; filter kosong tetap kosong; tidak menyembunyikan unfiltered write. |
| INSERT target qualifier menjadi kolom palsu | `TRSProcessSettlementBIS4.sql`, baris 222: `b.SellerPaymentAgent`, `b.SellerCashAccount`; `ReksaPopulateGenerateRDBCert.sql`, baris 39: `rsc.TranCode`, `rsc.ClientId` | Tidak mengambil identifier pertama sebagai kolom `b`/`rsc`. Seluruh exact positional mapping statement tersebut ditahan dan extraction notice menunjuk source line. |
| Target kolom yang tidak ada dianggap resolved | `TRSParamIssuerSave.sql`, baris 186: IssuerCode pada TRSParamIssuer_TH | Expression dipertahankan dengan `COLUMN_NOT_FOUND`; bukan mapping column terverifikasi. |
| Parameter PROC/default literal salah | Regresi default `'OUTPUT'`, OUTPUT parameter dan WITH EXECUTE AS | Header memakai token, mendukung PROC/PROCEDURE dan option boundary; keyword dalam literal tidak mengubah default/direction. |
| Dynamic SQL palsu/parsial | Regresi teks contoh EXEC di string dan literal + runtime suffix | Hanya token eksekusi nyata diproses. Complete literal diparse ulang; concatenation/runtime target tetap UNRESOLVED_DYNAMIC. Multiline EXEC/sys.sp_executesql tetap dikenali. |
| Encoding diam-diam menghilangkan object | Regresi UTF-16 desired source/history dan unsupported byte stream | Shared `NormalizeSQLToUTF8` hanya dipakai dalam memory. Encoding yang tidak dapat dikenali gagal dengan path; file tidak dinormalisasi oleh report. |
| Nama file dipercaya walau deklarasi berbeda | Regresi dbo.A.sql berisi CREATE TABLE/PROC dbo.B | Fail dengan source identity mismatch sebelum publication; tidak menerbitkan object/profile dengan nama salah. |

## Temuan pada snapshot yang memerlukan review pemilik SQL

Empat target column pada Bond tidak cocok dengan snapshot table:

| Procedure / baris statement | Target yang tercatat |
| --- | --- |
| RepoSaveParamSecurityDirtyPrice / 301 | REPOParamMaturityDate_TR.IsAbleForTransaksi |
| trs_DeleteSecurityTransaction_TT / 124 | SecurityTransaction_TH.SecAccNo |
| TRSParamIssuerSave / 186 | TRSParamIssuer_TH.IssuerCode |
| TRSProcessSettlementEarlyRedem / 184 | TRSSettlementEarlyLog_TH.`TransactionBranch ` (trailing space) |

Ini menunjukkan mismatch snapshot/source, bukan bukti otomatis bahwa database
runtime kehilangan kolom. Equivalence identifier yang bergantung collation atau
trailing space tetap membutuhkan validasi SQL Server. Dua INSERT berqualifier
di tabel temuan juga perlu review source. Engine report tidak menulis ulang SQL
bisnis atau menambahkan kolom berdasarkan tebakan.

## Kontrak komponen dan compatibility

- Extractor tetap catalog/metadata based dan menghasilkan canonical UTF-8 source;
  report memverifikasi source identity dan memakai decoder yang sama di memory.
- Checker interactive tetap hanya mengaplikasikan safe correction yang sudah
  dimiliki engine. `-check` dan Linux MR gate tetap read-only dan block correction
  pending; fitur ini tidak menambah fixer atau migration.
- Generator mempertahankan parser/diff/up-down/history policy. Perubahan grammar
  shared yang spesifik adalah MERGE JOIN hint tidak diperlakukan sebagai target
  DML/batas statement, sehingga dependency consumers memperoleh source yang sama.
- Report menambah function module/consumer, kolom `Resolution` di module lineage,
  dan `COLUMN_NOT_FOUND` pada write mapping. CSV/JSON migration-history schema
  tetap sama. Layout V2, ownership manifest dan no-write source tetap berlaku.
- Paket Windows/Linux wajib dibangun lengkap dari source digest final yang sama;
  source-only pass tidak menyatakan binary lama sudah membawa perubahan.

## Proof dan batas kemampuan

`fixtures/report_object_precision` berisi 18 SQL source files nyata tanpa modifikasi,
README, expectation dan golden. `TestRealObjectPrecisionFixtureReadOnlyAndIdempotent`
menjalankan report dua kali dan memverifikasi golden serta no-write source/history.
`TestReportRealSnapshotCorpus` memeriksa seluruh corpus lokal, termasuk semua
module dan hash 9.419 file object/history. Snapshot lokal opsional pada CI;
committed real fixture selalu berjalan. Negative tests berada di
`object_precision_test.go` dan `documentation_ddl_test.go`.

Perintah validasi: `go test ./...`, final hardening gate, `build_all.bat 9 3`,
checksum/source identity verification, released Windows/Linux report smoke, dan
Linux checker dengan flag yang sama dengan MR. Receipt/log disimpan di
`out/report-object-*`; file tersebut adalah output lokal, bukan source Git.

Residual limitations: scanner ini belum full T-SQL AST atau SQL Server runtime.
Projection lineage CTE/derived table/TVF, correlated cross-scope equality,
schema resolution yang ambigu, inferred cardinality, dynamic metadata lookup,
dan runtime-dependent SQL tetap tidak ditebak. Kolom view/external tidak divalidasi
terhadap physical schema yang tidak ada pada snapshot. WRITE melalui view dan
MERGE tetap direct impact/syntax evidence; MERGE branch-to-column mapping belum
dimodelkan. Trigger/synonym dan detail desired-state type bukan inventory baru
pada audit ini; history types yang sudah didukung tetap dilaporkan.

## Documentation review

Updated: README, rules, spec, FSD, plan, walkthrough, capability matrix,
traceability, QA gate, edge-case backlog, engine docs guide, production readiness,
join-identifier guide, generated QA coverage, dan fixture/audit ini.

Reviewed unchanged: desired-state lifecycle, guarded migrations, bootstrap four
repos, constraints naming, parallel-safe design, migration monitoring, REST
hardening, rebaseline, legacy wrapper cutover, GitLab checker job/MR templates,
QA-31 golden dan report output layout. Tidak ada perubahan pada desired-state
lifecycle, migration SQL/reversibility, extraction/bootstrap, naming, runtime/REST,
parallelism, pipeline flags, atau layout/ownership contract tersebut.


## Hasil final gate

- `go test ./...`: PASS; final hardening gate: PASS (learning examples yang tidak tersedia mengikuti existing SKIP policy).
- Complete `build_all.bat 9 3`: PASS. Version `20261012.133931`, source digest `c4b48b2d5e2161a949e6617f76ae1a6d0f914c3675e070285fa5ff87e6a425d0`. Windows checksums dan Linux release.sha256, embedded checker identity serta source digest: PASS.
- Exact Linux MR checker: fail=0, pass=4, migrations_generated=0. Linux regression/real corpus: PASS; 9.419 authored object/history files unchanged.
- Released Windows/Linux report fixture: status OK, warnings=0, errors=0; 10 Markdown documents identik byte-for-byte. Released report pada Bond retail lengkap: PASS.
- Tiga snapshot lain tidak memiliki migrations directory; CLI tetap menolak missing history sesuai kontrak. Desired-state reader semua empat snapshot lulus corpus test. Tidak membuat migrations directory pada snapshot asli.

Report remains conservative: empat missing-column findings dan dua invalid INSERT lists perlu review source/schema secara manual. Audit ini tidak menyatakan full runtime SQL Server compatibility.
