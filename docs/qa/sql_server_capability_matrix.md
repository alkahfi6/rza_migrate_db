# SQL Server Capability Matrix

### Object scan refinement — 2026-10-01

| Area | Classification / cross-tool scope | Proof |
| --- | --- | --- |
| Quoted/escaped names, malformed delimiter, DML TOP, MERGE USING, APPLY/comma sources, SELECT/OUTPUT sinks, static EXEC source spans | Supported for offline access/dependency recognition; this does not add a new DDL migration object kind. Extractor catalog unchanged; checker apply/CI adds no correction; generator ordering uses shared references without rewriting business SQL. Full released binaries required. | `sqlparse/access_test.go`, `dependencies_test.go`, QA-31 golden |
| Statement/query-local DML aliases, repeated same-line statements, UPDATE mapping boundaries | Supported; ambiguous candidates remain AMBIGUOUS. | `migreport/procedure_analysis_test.go`, `engine/object_scan_contract_test.go` |
| CTE case variants with unknown collation | Ambiguous in report; exact-case possible physical dependency retained in engine. | `TestProcedureCTECaseDependsOnUnknownCollation`, existing case-distinct dependency test |
| INSERT multirow or TOP/DISTINCT/compound expression mapping; runtime SQL/branch constant propagation; full T-SQL grammar | Manual review / not claimed as full exact mapping. Access inventory remains available. | negative mapping regression; [residual review](../review_object_scan_20261001.md) |


Dokumen ini mengunci perilaku engine saat menemukan fitur SQL Server yang didukung, belum didukung, atau harus dikerjakan manual.

Status:
- `supported`: boleh diproses otomatis oleh `verify`/`gen`.
- `fail-hard`: `verify` atau `gen` wajib menambah fail item severity 3.
- `manual-only`: engine boleh emit migration manual/skip warning, tetapi tidak boleh mengarang DDL destruktif.
- `skipped-warning`: extractor boleh lanjut dan mencatat warning terstruktur.

## Table Snapshot (`db_object/tables`)

| Feature | Status | Rule/Reason | Catatan |
| --- | --- | --- | --- |
| `CREATE TABLE dbo.X (...)` sederhana | supported | - | Schema wajib `dbo`. |
| Kolom nullable/non-nullable tipe umum | supported | - | Canonical hash menormalkan noise style. |
| Inline named default: `<col> <type> CONSTRAINT <name> DEFAULT (...)` | supported | - | Bentuk default yang diwajibkan untuk source. |
| Table-level named default: `CONSTRAINT <name> DEFAULT (...) FOR <col>` | supported | semantic normalization | Setara dengan inline named default; duplicate default pada column yang sama tetap ditolak. |
| Constraint tanpa `CONSTRAINT <name>` | fail-hard | `UNNAMED_CONSTRAINT_NOT_ALLOWED` | Termasuk `PRIMARY KEY (...)`, inline `DEFAULT` tanpa nama, `CHECK`, `UNIQUE`, `REFERENCES`. |
| `ALTER TABLE`, `DROP TABLE`, wrapper `IF ... BEGIN ... END` di source table | fail-hard | source contract violation | Source table harus desired snapshot, bukan migration script. |
| Extract/add computed column (`AS <expression> [PERSISTED]`) | supported | extractor emits computed expression | Extractor now reads `sys.computed_columns.definition`; empty definition fails clearly. |
| Change expression computed non-persisted | supported-guarded | `GUARDED_AUTO_MIGRATION` | Drop/add memakai guard `sys.computed_columns`; dependency computed-ke-computed yang memerlukan urutan khusus ditolak untuk migrasi manual. |
| Change persistence/drop computed column | manual-only | `MANUAL_MIGRATION_REQUIRED` | Nilai dan dependency/index lifecycle belum selalu reversible. |
| `SPARSE NULL` | supported | typed column modifier | Extractor membaca `sys.columns.is_sparse`; `SPARSE NOT NULL` ditolak sebelum generation. |
| Filestream/temporal/ledger/graph/memory-optimized/partition/filegroup | fail-hard / skipped-warning di extractor | unsupported table subset / `TABLE_EXTRACT_SKIPPED_UNSUPPORTED` | Belum aman untuk canonical offline planning. Extractor juga fail-closed untuk typed XML/UDT, mask/encryption, hidden/generated/vector column, trigger, user statistics, CDC/change tracking/replication, legacy bound rule/default, dan heap compression. |
| Remove table dari desired state | manual-only | `REQUIRES_MANUAL`, `DATA_RESTORE_NOT_GUARANTEED` | Engine tidak auto emit `DROP TABLE`. |
| Add nullable column | supported | deterministic `ALTER TABLE ADD` | `down` schema-only drop column. |
| Add NOT NULL column dengan named default | supported | `WITH VALUES` | Default wajib named dan parseable. |
| `NULL` → `NOT NULL` | supported-guarded | `GUARDED_AUTO_MIGRATION` | Precheck count NULL; gagal memakai `THROW` dan rollback. |
| String narrowing dalam type/encoding sama | supported-guarded | `GUARDED_AUTO_MIGRATION` | `DATALENGTH` harus muat pada up/down. |
| Binary narrowing (`BINARY`/`VARBINARY`) dalam base type sama | supported-guarded | `GUARDED_AUTO_MIGRATION` | `DATALENGTH` harus muat pada up/down. |
| Parameterized column type (`DECIMAL/NUMERIC`, character/binary length, `TIME/DATETIME2/DATETIMEOFFSET`) | supported | typed column diff | Whitespace sekitar kurung/koma dan delimiter tunggal tipe bawaan (`[DECIMAL]`) dinormalisasi; perubahan nilai parameter tetap menghasilkan `ALTER COLUMN` reversible dengan guard sesuai risiko. |
| Numeric/date-time conversion | supported-guarded | `GUARDED_AUTO_MIGRATION` | `TRY_CONVERT` + round-trip conversion harus lossless. |
| Transisi tipe lintas keluarga tanpa pembuktian lossless | manual-only | `MANUAL_MIGRATION_REQUIRED` | Ditolak sebelum file migration ditulis; contoh `UNIQUEIDENTIFIER` ke `INT`. |
| Named `PRIMARY KEY` / `UNIQUE` lifecycle | supported-guarded | `GUARDED_AUTO_MIGRATION` | Precheck NULL/duplicate. Incoming named FK dari history sementara di-drop lalu direstore dengan orphan precheck. |
| Named `FOREIGN KEY` lifecycle | supported-guarded | `GUARDED_AUTO_MIGRATION` | Mendukung local/referenced composite columns dan raw cascade option; orphan precheck mengabaikan key yang mengandung NULL sesuai FK semantics. |
| `NOT FOR REPLICATION` pada named CHECK/FK | supported | catalog round-trip | Extractor mempertahankan flag dan renderer memakai posisi grammar MSSQL yang benar. |
| CHECK/FK disabled atau untrusted | supported logical snapshot + health warning | `CONSTRAINT_STATE_REVIEW_REQUIRED` | Definisi constraint tetap diekspor; observed enabled/trusted state dicatat terstruktur dan database tidak diubah. `-require-healthy-constraints` mempromosikan finding menjadi exit non-zero setelah export. |
| Opsi fisik non-default pada index backing PK/UQ | supported logical snapshot + warning | `KEY_PHYSICAL_METADATA_NOT_CAPTURED` | Definisi logical PK/UQ tetap diekspor. Compression, fill factor, pad/lock options, partition/filegroup dilaporkan sebagai metadata DBA yang tidak menjadi isi canonical table snapshot. |
| Identity `NOT FOR REPLICATION` | skipped-warning di extractor | preserve previous snapshot | Typed lifecycle identity belum menyimpan flag ini; extractor fail-closed. |
| Add `IDENTITY(seed, increment)` column | supported | deterministic `ALTER TABLE ADD` | Typed model menyimpan seed/increment; perubahan identity existing tetap manual-only. |
| Add `UNIQUEIDENTIFIER ROWGUIDCOL` column | supported | deterministic `ALTER TABLE ADD` | Verifier menolak `ROWGUIDCOL` pada tipe selain `UNIQUEIDENTIFIER`. |
| Change existing `IDENTITY` / `ROWGUIDCOL` | manual-only | `MANUAL_MIGRATION_REQUIRED` | SQL Server tidak mendukung alter in-place tanpa rebuild table. |
| Column masking/encryption | manual-only | `UNSUPPORTED_SQL_FEATURE` | Typed model dan capability server belum lengkap. |
| Collation-only change | manual-only | `MANUAL_MIGRATION_REQUIRED` | Perubahan comparison semantics dan dependency index tidak dianggap lossless secara otomatis. |
| Explicit column rename dengan directive | supported-guarded | `TYPED_COLUMN_RENAME` | `sp_rename` reversible; named CHECK/index dependency di-drop/recreate. Rename+alter harus dipisah. |
| Drop regular column | supported-guarded | `GUARDED_AUTO_MIGRATION` | Hanya dijalankan jika tidak ada nilai non-NULL dan seluruh named dependency dapat direkonstruksi; down menambah kembali schema. |
| Computed/non-computed transition | manual-only | `MANUAL_MIGRATION_REQUIRED` | Membutuhkan keputusan eksplisit tentang materialisasi atau kehilangan nilai. |
| Drop/alter column dengan named dependency | supported terbatas | deterministic precheck/drop dependency | Dependency harus diketahui dari history/source dan bernama. |
| Initial `CREATE TABLE` saat table sudah ada | supported bootstrap-compatible | required-column guard | Wajib cocok: nama kolom exact-case, base type, identity-kind, computed-kind. Extra column, length, precision/scale, nullability, collation, constraint, dan physical option tidak memblokir skip. |

## Index Snapshot (`db_object/indexes`)

| Feature | Status | Rule/Reason | Catatan |
| --- | --- | --- | --- |
| `CREATE INDEX` / `CREATE UNIQUE INDEX` untuk table target | supported | - | File `indexes/dbo.Table.sql` hanya boleh berisi index table tersebut. |
| Clustered/nonclustered variants | supported | - | Parser sudah punya contract test. |
| Filtered index (`WHERE ...`) | supported | - | Termasuk supported subset. |
| Composite key `ASC`/`DESC` + multi-column `INCLUDE` | supported | exact-case column precheck | Semua key dan INCLUDE column diverifikasi lewat `sys.columns` dengan collation BIN2 sebelum CREATE. Trailing comma/malformed INCLUDE ditolak. |
| Index source target table berbeda | fail-hard | `INDEX_SOURCE_TABLE_MISMATCH` | Mencegah drift lintas file. |
| `DROP INDEX` di source index | fail-hard | source contract violation | Drop index harus hasil diff/generation, bukan data master. |
| Primary + secondary XML index (`PATH`/`VALUE`/`PROPERTY`) | supported | dependency-ordered lifecycle | Primary dibuat sebelum secondary; secondary dihapus sebelum primary. Extractor memakai `sys.xml_indexes`. |
| Spatial geometry/geography index | supported terbatas | catalog round-trip | Extractor mempertahankan tessellation, bounding box, empat manual grid level, cells-per-object, compression/fill/pad/lock options. Metadata tidak lengkap tetap fail-closed. |
| Clustered/nonclustered columnstore | supported | raw-preserving lifecycle + catalog extraction | Filtered NCCI dan uniform `COLUMNSTORE_ARCHIVE` dipertahankan; opsi lock rowstore tidak ditulis ke columnstore. Ordered/data-clustered/compression-delay tetap fail-closed. |
| Clustered columnstore tanpa key column list | supported | keyless parser contract | Precheck hanya membutuhkan target table; nonclustered columnstore tetap memeriksa seluruh column. |
| Uniform index data compression (`ROW`/`PAGE`/`COLUMNSTORE`) | supported | catalog round-trip | Berlaku bila seluruh partition memakai compression yang sama. Per-partition compression masih fail-closed. |
| Online/resumable/partitioned index | fail-hard | `UNSUPPORTED_SQL_FEATURE` | Membutuhkan execution mode dan partition model khusus. |
| Full-text, selective XML, memory-optimized hash, ordered/data-clustered columnstore, `OPTIMIZE_FOR_SEQUENTIAL_KEY`, auto-created index, compression delay, XML compression | skipped-warning di extractor | preserve previous snapshot | Catalog dideteksi eksplisit, termasuk kolom catalog version-dependent melalui dynamic SQL; extractor tidak boleh menghasilkan snapshot lossy. |

Catatan: capability harus mengikuti implementasi scanner aktual. Saat ini scanner tetap fail-closed untuk resumable dan `ON PARTITIONS`. Opsi operasional lain tidak boleh dianggap didukung hanya karena lolos token scanner; dukungan production baru sah setelah parser, round-trip, extractor, dan integration test tersedia.

## Programmability (`views/functions/procedures`)

| Feature | Status | Rule/Reason | Catatan |
| --- | --- | --- | --- |
| Satu object per file dengan `CREATE OR ALTER` | supported | - | Header harus match filename `dbo.<Name>.sql`. |
| Initial up/down programmability | supported | executable inverse | Up memakai `CREATE OR ALTER`; down memakai `DROP ... IF EXISTS` dengan identifier case-preserving. |
| `CREATE`, `ALTER`, atau missing `OR ALTER` di source | fail-hard | programmability source/header contract | Source harus canonical untuk repeatability. |
| Multiple persistent object dalam satu file | fail-hard | `PROGRAMMABILITY_MULTI_OBJECT` / source contract | Temp table DDL di body tetap boleh sesuai contract. |
| Komentar dan string literal | supported | - | Scanner forbidden SQL harus mengabaikan komentar/string. |
| Encrypted/definition kosong dari extractor | skipped-warning | extractor warning | Module di-skip dan dicatat di `out/get_object_summary.json`. |
| Formatter backend `antlr` | skipped-warning/fallback | fallback to `lite` | Saat ini ANTLR belum aktif; output harus sama dengan backend `lite`. |

## Patch (`patch/<name>`)

| Feature | Status | Rule/Reason | Catatan |
| --- | --- | --- | --- |
| Tepat satu logical `README.md`, `up.sql`, `down.sql` | supported | cross-platform case-insensitive resolver | Nama entry aktual boleh berbeda casing (`readme.md`, `UP.SQL`); Windows/Linux wajib memberi hasil identik. File tambahan dan duplicate case-only ditolak. |
| Patch name dengan spasi/simbol | supported | canonicalization | Nama migration memakai canonical patch token. |
| Source patch berubah setelah migration dibuat | fail-hard | `PATCH_SOURCE_MUTATED` | Berdasarkan hash header patch. |
| Source patch hilang setelah migration dibuat | fail-hard | `PATCH_SOURCE_MISSING` | Patch source immutable; console menunjukkan expected patch folder, migration asal, canonical patch name, dan saran Git/casing. |

## Runtime Migration Runner

| Feature | Status | Rule/Reason | Catatan |
| --- | --- | --- | --- |
| `up` via REST | supported | - | Serialized in-process; DB lock opsional. |
| `down` | manual-only by config | disabled default | `allow_down=false`. |
| `run-file` | manual-only by config | disabled default | Jika enabled, harus next/current version deterministic. |
| Dirty version acknowledgement | manual-only by config | disabled default | Hanya untuk operator yang memahami risiko. |

## Migration History Reconciliation & Report (`rza_migrate_report`)

| Feature | Status | Rule/Reason | Catatan |
| --- | --- | --- | --- |
| History source `migrations/` + desired-state docs `db_object/` | supported | separated read-only sources | `migrations/` tetap satu-satunya source rekonstruksi history; `db_object/` hanya dibaca untuk current-state docs dan procedure profiles, serta tidak memengaruhi history. |
| Core audit bootstrap exclusion | supported | application-history boundary | Pair canonical `000000__core__mig__audit_init` dipindai read-only tetapi tidak dipancarkan ke history aplikasi; analyzed/excluded pair count dilaporkan terpisah pada CSV, JSON, dan console. |
| Desired-state unique-ID documentation | supported, conservative | enforcement vs signal separation | PK/UQ/unfiltered unique index enforce; filtered unique conditional; identity generates; JOIN/non-unique index/ID-like name hanya review signal. Inline PK/UQ didukung. |
| Suspected JOIN identifier listing | supported, report-only | observed equality column mapping + provenance | Index/IDENTITY/name-ID tidak wajib; PK pada kolom lain tidak menyembunyikan reference column. Evidence source modules/usage/index/key/identity; generic JOIN operands bukan proven row ID. FK-only/inequality/ambiguous/runtime tidak menjadi observed JOIN; composite/filtered uniqueness tetap terbatas. Extractor/checker/generator/CI tidak mendapat schema action; full released report binary parity diwajibkan. QA-31 golden dan no-write/rerun test. |
| Desired-state relationship documentation | supported, conservative | declared FK + static equality JOIN | Exact alias-column mapping dari `JOIN ... ON`; same-name columns, inequality, ambiguous expressions, temp variables, stale aliases, dan dynamic SQL tidak membuat edge. |
| Desired-state module lineage | supported, conservative | static object references | View/procedure dependency dipisahkan dari table relationship; identity dependency di-deduplicate case-insensitive. `EXEC @return_code = Procedure` adalah static call, sedangkan true variable execution masuk manual-review list. |
| Procedure impact profiles | supported, conservative | source-positioned token scan + catalog-backed static analysis | Satu profile per procedure memuat contract, top-down branch-aware local execution flow, local object map, source-preserved guards, exact writes, direct/recursive calls, recursive write impact, key context, functions/views, dan dynamic SQL. Object multiline dikenali; CTE dan alias dibatasi statement; schema-qualified table tidak tertimpa nama CTE; `SELECT INTO`/assignment/subquery/INSERT SELECT bukan client result. Flow merangkum transaction tanpa child call, condition/rejoin, procedure routing, dan error transfer; full source sequence tetap tersedia. Node/edge procedure memakai Mermaid-rendered line break dan tidak menampilkan literal `\n`; link profile URL-safe untuk identifier berkarakter khusus. Resolution code membedakan static, reconstructed literal, caller literal, runtime unresolved, external, temp, missing, ambiguous, dan system object; catalog tidak digunakan untuk menebak runtime target. Mermaid dibatasi 20 node per map tetapi map tambahan mempertahankan seluruh recursive call edge. |
| Report output roots and manifest | supported | repository-relative deterministic paths | Relative `-out`/`-docs` selalu berbasis resolved repository root. JSON memuat kedua absolute root dan concrete `artifacts[]`; compatibility `reports[]` tetap tersedia. Dokumentasi ditulis sebelum audit JSON dipublikasikan. |
| Markdown report usability | supported | single-file adaptive relational maps | Narasi Indonesia + conventional English database terms; seluruh Mermaid satu file; connected Relational Map maksimum 20 object, 24 relationship lines, 35 mappings, dan 9.000 karakter tanpa artificial Hub; node memuat type dan key/index context; setiap non-self pair dan exact mapping tampil, dense pair memakai bounded Mapping batch; exact evidence tetap satu row di catalog; self-join dipisahkan; tanpa raw HTML `<br>`. |
| Zero database connection & zero SQL execution | supported | strict offline security boundary | Tidak pernah membuka koneksi SQL Server atau membaca metadata katalog/sys/journal. |
| Immutability of migration files | supported | read-only contract | Tidak membuat, mengedit, menghapus, atau merename file migration apa pun. |
| Semantic Table Operations | supported | SQL AST / token parser | Mendukung deteksi `CREATE_TABLE`, `DROP_TABLE`, `ADD_COLUMN`, `DROP_COLUMN`, `RENAME_COLUMN`, `ALTER_COLUMN_TYPE`, `ALTER_NULLABILITY`, `ALTER_DEFAULT`, `ADD_CONSTRAINT`, `DROP_CONSTRAINT`, `RENAME_CONSTRAINT`, `ALTER_CONSTRAINT`; satu `CREATE TABLE` dalam migration wrapper dipancarkan tepat sekali dan nama column `[]`/`""` berspasi dipertahankan utuh. |
| Semantic Index Operations | supported | SQL AST / token parser | Mendukung deteksi `ADD_INDEX`, `DROP_INDEX`, `ALTER_INDEX`, `RENAME_INDEX` beserta atribut columns, include, filter, unique, clustering. |
| Programmable Operations | supported | SQL AST / definition hashing | Mendukung `CREATE`, `ALTER`, `DROP`, `REPLACE_DEFINITION` dengan perbandingan bagian yang berubah (`changed_sections`). |
| Type Operations | supported | SQL AST / token parser | Mendukung ekstraksi user-defined type definition changes. |
| Noise Filter (Comments & Formatting) | supported | AST canonical normalization | Komentar (`--`, `/* */`) dan variasi whitespace tidak dilaporkan sebagai operasi semantik. |
| Unnamed or Renamed Constraints with Identical Definition | supported | semantic constraint equality | Perubahan nama constraint dengan definisi identik tidak dianggap sebagai perubahan semantik. |
| Ambiguous / Complex SQL Constructs | supported with LOW confidence | `MANUAL_REVIEW_REQUIRED` | SQL yang tidak dapat dipetakan secara pasti oleh AST diberi confidence `LOW` dan catatan review manual. |
| Auto-Discovery Repository Root | supported | interactive/CLI usability | Deteksi cerdas: argumen `-repo`, direktori binary, current working dir, atau parent dirs. |
| RFC 4180 CSV, JSON, & Markdown outputs | supported | strict output contract | 7 CSV ber-header valid (UTF-8 tanpa BOM) dan 1 JSON history report di `docs/database/migrations/`, ditambah root README/manifest, empat grouped Markdown, procedure index/catalog, dan generated profiles di `docs/database/`. |
| Shared Strict Filename Contract | supported | cross-tool linearity | Memakai `internal/history`; filename invalid dilaporkan sebagai warning, bukan diterima sebagai migration canonical. |
| Semantic Up/Down Reconciliation | supported | inverse-operation validation | Down kosong/inverse tidak cocok = `FAIL` + warning; ambiguous unsupported = `NOT_TESTED`. |
| Dense Mixed `ALTER TABLE` | supported | ordered state reduction | Multi-column, multi-constraint, mixed add, composite FK/UQ, multiple drop, add→rename→alter, dan drop→re-add diproses per item sesuai urutan source. |
| Constraint/Index Drop-Recreate Reduction | supported | semantic before/after | Replacement bernama sama menjadi satu `ALTER_CONSTRAINT`, `ALTER_DEFAULT`, atau `ALTER_INDEX`; seluruh atribut lama/baru dipertahankan. |
| Protected Report Output Path | supported | migration immutability | `-out` dan `-docs` yang sama dengan atau berada di bawah `migrations/` atau `db_object/` ditolak sebelum write. |
| Programmable/Type Delete-Reintroduce Lifecycle | supported | reconstructed object state | Drop mengosongkan state; create berikutnya dilabeli `REINTRODUCED`. |
| Table History v1→v7 + Delete/Reintroduce | supported | longitudinal reconstructed state | State kolom/constraint/index dilacak lintas version dan di-reset pada drop/create; before-value pasca-reintroduksi berasal dari incarnation terbaru. |

## Report layout V2

Report layout V2: supported setelah targeted/full suite dan release gates; single root, grouped data/documents, hash manifest, escaped navigation, preflight no-write manual collision, dan checksum-only stale cleanup. Legacy output yang tidak terbukti equivalent dipertahankan, bukan guessed deletion. Source/history contract lintas tool tetap unchanged.


## Audit object precision (2026-10-12)

Supported observational report: multiline token DDL, inline/ALTER FK, full datatype, scoped OUTER/hinted JOIN, quoted identifier, function lineage, UPDATE tanpa WHERE, COLUMN_NOT_FOUND review, Unicode memory-only decode dan declared identity validation. Dynamic literal supported hanya jika actual EXEC dan literal lengkap; runtime-dependent target, ambiguous alias/schema, CTE/derived projection dan MERGE branch column mapping tetap manual/unresolved. Migration support classification tidak berubah.

Detail bukti: ../review_report_object_precision_20261012.md.
