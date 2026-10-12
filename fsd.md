# Functional Specification Design (FSD) - rza_migrate_db Engine Suite

2026-10-12 JOIN identifier listing: `internal/migreport/identifier_join_review.go` memproyeksikan `DesiredStateDocs.Relationships` dengan evidence `OBSERVED_JOIN` ke column review rows pada `table_unique_ids.md`. Source scanner/relationship parsing tidak berubah. Rows hanya untuk local table columns tanpa individually enforced unfiltered unique key; metadata index/identity bukan prerequisite. Canonical mapping dan module provenance diurutkan/didedup deterministik, self-role tidak menggandakan occurrence, FK-only evidence tidak ikut Join usage. Filtered single-column key tetap conditional, composite key menjelaskan tuple-level uniqueness. Rendering ini tidak dipakai extractor/checker/planner untuk rewrite atau schema action; released report binary membawa output yang sama pada Windows/Linux.


Refinement object scan (2026-10-01): `internal/sqlparse/access.go` menjadi lexical source bersama untuk access, static calls, dan top-level clause offsets. Token quoted tetap identifier, string menjadi opaque barrier, invalid delimiter tidak diiris sebagai nama, dan `NameParts` mempertahankan separator vs titik di dalam quoted name. SQLAccess menyimpan statement bounds/query depth, CTE identity, dan ambiguity; resolusi alias sintaktis mendahului catalog lookup. `dependencies.go` memakai token yang sama, mengecualikan column-list/function-definition false positives serta external/temporary edges, dan tidak membuat routine dari target INSERT. `migreport` menjaga setiap source occurrence, memakai bounds untuk evidence/call arguments/write mapping, dan tidak membuang kandidat alias ambigu. Scanner ini belum full AST; migration schema diff/rewriting tetap milik parser/planner yang sudah ada. Component boundaries dan residual limitations ada di [review object scan](docs/review_object_scan_20261001.md).


Dokumen ini menjabarkan bagaimana requirement di `spec.md` dan `rules.md` diimplementasikan dalam engine.

## 1. Tujuan Desain

- Menentukan arsitektur, kontrak modul, dan alur data yang deterministic.
- Mengunci implementasi agar semua hard lock dapat diverifikasi otomatis.
- Menyediakan blueprint coding dan testing yang langsung executable.

## 2. Arsitektur Modul

Struktur code target:

- `cmd/rza_migrate_db/main.go`
- `cmd/rza_get_object_db/main.go`
- `cmd/rza_migrate_rest/main.go`
- `cmd/rza_migrate_report/main.go`
- `internal/cmd/rest.go`
- `internal/handler`
- `internal/migrest`
- `internal/migreport`
- `internal/config`
- `internal/fsx`
- `internal/sqlfmt`
- `internal/sqlparse`
- `internal/history`
- `internal/diff`
- `internal/emit`
- `internal/verify`
- `internal/output`
- `internal/getobject`
- `internal/tsqlast`

## 2.1 Status Modularisasi Package (2026-02-20)

| Package | Status | Implementasi aktif |
| --- | --- | --- |
| `internal/fsx` | Selesai | file discovery recursive + stable ordering |
| `internal/sqlfmt` | Selesai | normalizer text SQL |
| `internal/sqlparse` | Selesai | lexical/parser helper subset MVP |
| `internal/history` | Selesai | parser filename migration + helper history key/version |
| `internal/diff` | Selesai | drift comparator + sort comparator deterministik |
| `internal/emit` | Selesai | filename builder + allocator version pair |
| `internal/verify` | Selesai | orchestrator urutan validate + validator concrete domain (`migrations`, `forbidden_sql`, `repo_hygiene`, `patch`, `db_object`, `wrapper`, `table_inline_default`) |
| `internal/output` | Selesai | report writer `out/output.json` |
| `internal/getobject` | Selesai | extractor SQL Server -> `db_object/*` + summary `out/get_object_summary.json` |
| `internal/tsqlast` | Selesai (fallback mode) | formatter modul SQL backend `lite`, opsi `antlr` fallback otomatis |
| `internal/migrest` | Selesai | service eksekusi migrate SQL Server + lock policy + run-file deterministic guard |
| `internal/migreport` | Selesai | offline migration history reconstruction report generator (`docs/database/migrations/*`) |
| `internal/handler` | Selesai | HTTP handler REST API (`up/down/steps/run-file/version/healthz`) |

Kontrak tanggung jawab:

- `config`: hard-locked app config (`app_config.go`) dan enum whitelist.
- `fsx`: file discovery, stable sorting, path validation.
- `sqlfmt`: format canonical db_object dan patch body normalization.
- `sqlparse`: parser structured filename dan SQL subset MVP (tanpa regex parser utama).
- `history`: read migration history, build reconstructed state.
- `diff`: desired-vs-state semantic diff dan policy mapping.
- `emit`: generate pair up/down sesuai template wrapper.
- `verify`: orchestrator rule checks pada source dan migration output.
- `output`: builder `out/output.json` dan exit decision.
- `getobject`: export object DB (table/index/module) ke source-of-truth + summary run.
- `tsqlast`: formatter modul programmability shared (dipakai extractor dan auto-indent pada `gen`).
- `migrest`: runtime config YAML + eksekusi `golang-migrate` + DB lock + guard specific-file.
- `migreport`: scanner & analyzer read-only `migrations/` untuk rekonsiliasi histori, serta pembaca `db_object/` untuk dokumentasi identifier, observed relationship, module lineage, dan procedure impact profiles; output 7 CSV + JSON, root README/manifest, empat grouped Markdown, procedure index/catalog, dan satu generated profile per procedure.
- Desired-state analyzer mem-parsing PK/UNIQUE table-level maupun inline, unique/filtered/non-unique index, serta deklarasi `IDENTITY` berbatas token. Writer memisahkan enforcement, conditional enforcement, generation, dan ID review signal; tidak ada signal yang mengubah desired state atau migration planning.
- Relationship analyzer mempertahankan alias scope berdasarkan shared query depth dan statement bounds, lalu menerima static explicit `JOIN ... ON` dan legacy comma-join `WHERE` equality yang setiap sisinya mempunyai tepat satu resolved `alias.column`. Wrapper fungsi satu-kolom dapat diterima; inequality, multi-column expression, temp/table variable, dan dynamic SQL di-skip. FK declared dan observed JOIN masuk catalog/graph dengan evidence berbeda; semua module reference masuk lineage terpisah, sedangkan module tanpa resolved reference masuk coverage-review list.
- Markdown writer memakai narasi Bahasa Indonesia dan conventional English database terms, stable evidence codes, serta delimiter `;` untuk multi-value. Template risk/action per classification dipindahkan ke reference section. Non-self mappings di-collapse per object pair. Deterministic connected-edge traversal memulai map dari object dengan uncovered degree tertinggi, memprioritaskan edge yang menutup object yang sudah ada, lalu menambahkan actual edge yang menyentuh node dalam map; tidak ada artificial Hub, parent, atau edge. Relational Map dibentuk adaptif dengan hard limit 20 object, 24 relationship lines, 35 exact mappings, dan 9.000 karakter Mermaid. Dense pair dipecah deterministik menjadi numbered `Mapping batch` sebelum traversal. Node memuat TABLE/VIEW/EXTERNAL OBJECT dan status key/index bagi kolom relationship. Edge mempertahankan `left_column = right_column`; `[FK]`/`[JOIN]` hanya ditampilkan bila satu object pair memiliki mixed evidence. Occurrence/usage tetap tersedia di catalog dan tidak dipakai menggantikan informasi struktur. Seluruh Mermaid ditulis dalam satu `relational_erd.md`; catalog mengikuti Relationship Group/Relational Map yang sama dan mempertahankan setiap exact mapping sebagai satu row. Object index dan shared-object table menyediakan navigasi lokal, self-join dikeluarkan dari graph dan ditulis pada review section, dan writer mempertahankan unowned `relationship_diagrams/` dari format lama untuk manual review. Analyzer/writer hanya dipanggil oleh report path sehingga extractor, checker apply/plan, generator, dan pipeline authoring tetap tidak berubah.
- Procedure analyzer membangun catalog terpisah untuk table, view, function, dan procedure hanya dari file `db_object`. `internal/sqlparse.ExtractSQLAccesses` melakukan lexical scan shared dengan source offset, multipart identity, statement boundary, alias, serta CTE scope; `migreport` kemudian melakukan catalog resolution. Pemisahan ini mengikuti pola lexer/symbol scope yang dipelajari dari kahfigo tanpa membawa default metadata atau asumsi generatornya. Static reference di-resolve case-insensitive tetapi collision tetap `AMBIGUOUS`; unqualified name hanya diterima bila tepat satu kandidat catalog. DML alias dicari dalam boundary statement; target/read multiline tetap dikenali; CTE hanya `NOT_PERSISTED` pada statement aktif; schema-qualified temporary object dilabel `NOT_PERSISTED`; `sys.*` dilabel `SYSTEM_OBJECT`; `Database..Object` dilabel external; table-valued/built-in function dipisahkan dari data access; dan cursor `FETCH ... FROM` diabaikan sebagai object. `internal/sqlparse.HasClientResultSet` membedakan client SELECT dari assignment, `SELECT INTO`, subquery/EXISTS, dan INSERT SELECT. Literal dynamic SQL direkonstruksi serta diparse ulang; variable dynamic SQL menyimpan execution expression, runtime variables, dan assignment evidence tanpa menebak target. Bentuk `EXEC @return_code = schema.Procedure` dikenali sebagai static call dengan return assignment, sedangkan `EXEC @sql` tetap dynamic. Literal dari static caller dapat menambah kandidat `RESOLVED_FROM_CALLER` dengan provenance caller. Dependency lineage menggunakan identity case-insensitive agar kapitalisasi berbeda tidak menghasilkan row ganda. Writer menghasilkan contract, full source-ordered process sequence, branch-aware local execution flow, local object map, execution signals, direct access plus source guard, exact INSERT/UPDATE mapping, enforced-key/filter context, readable per-call blocks dengan child-parameter ← caller-expression mapping, cycle-safe recursive procedure maps/paths dan write impact, extraction notices, serta resolution glossary. Overview flow memakai `flowchart TD`, mengelompokkan condition reads, membentuk conditional routing/rejoin, merangkum transaction tanpa child `EXEC`, dan mengarahkan `GOTO ERROR` ke error handler; full evidence tidak dihapus dari tabel. Procedure-specific Mermaid formatter mengubah logical line break menjadi `<br/>` hanya di dalam Mermaid code fence agar viewer tidak menampilkan literal `\n`; tabel Markdown tidak memakai raw HTML. Link profile menggunakan URL path escaping sehingga karakter seperti `#` tidak dibaca sebagai fragment. Guard structure dibaca dari sanitized SQL tetapi text/literal diambil dari source asli. Local dan recursive Mermaid dipecah maksimal 20 node per map; seluruh edge label dikutip untuk menampung operator SQL, sementara detail dan recursive call edge tetap lengkap.
- `migreport.Run` mengubah repository root menjadi absolute canonical path dan menghitung relative `-out`/`-docs` dari root tersebut. Publisher menulis kandidat lengkap di temporary staging, memvalidasi semua target sebelum write, mengganti setiap file secara atomic, melakukan cleanup file yang terbukti owned, lalu menerbitkan `manifest.json` terakhir. JSON history mempertahankan schema/data fields; inventory path mengikuti layout V2 dan `reports[]` relatif terhadap documentation root. Manifest V2 memakai relative root_paths agar clone/OS relocation tetap dapat memverifikasi ownership, serta membawa exact-byte SHA-256 untuk semua artifact selain manifest sendiri; console memakai inventory konkret yang sama. Tidak ada transaksi atomic lintas seluruh folder: kegagalan filesystem saat publication dapat menyisakan partial update yang dilaporkan sebagai error.
- Integrasi lintas tool sengaja non-applicable untuk extractor/checker/generator karena capability ini read-only documentation pada command report. Source `db_object`, migrations, canonicalization, safe-fix, verify, dan generated migration tidak dimutasi; release Windows/Linux tetap dibangun atomik dari source revision yang sama.
- `migreport` memanggil parser filename strict `internal/history` dengan whitelist `internal/config`, menjaga satu kontrak metadata lintas checker/generator/report. File invalid tetap diinventarisasi sebagai warning melalui fallback identity, tetapi tidak lolos sebagai metadata canonical.
- Sesudah scan dan sebelum analyzer, report memfilter hanya pair bootstrap canonical `000000__core__mig__audit_init`; file tetap dihitung sebagai scan fisik, tetapi tidak menambah migration summary, object history, atau perubahan semantik aplikasi. Filter mengembalikan counter exclusion yang dipublikasikan sebagai `excluded_pairs` di summary CSV, `excluded_migration_pairs` di JSON, dan banner console bersama analyzed pair count.
- Analyzer report memisahkan statement yang dikenali dari operasi semantik yang benar-benar dipancarkan: no-op type/nullability dan rename constraint identik disuppress tanpa berubah menjadi `SQL_AMBIGUOUS`; file up kosong tidak membentuk state programmable/type semu.
- Validator round-trip mem-parsing up/down dan mencocokkan inverse table, column, constraint, index, programmable, dan type. `CREATE TABLE` + `DROP TABLE` mencakup child column/constraint/index; unsupported SQL menghasilkan `NOT_TESTED`, sedangkan inverse hilang menghasilkan `FAIL` + warning.
- Parser table memakai top-level comma splitter yang quote/parenthesis-aware untuk memisahkan mixed multi-column/multi-constraint dan composite FK/UQ. `WITH CHECK`/`WITH NOCHECK`, multiple drop, dan `IF EXISTS` dipetakan ke operasi individual.
- Scanner `CREATE TABLE` memajukan cursor dari offset relatif penuh (posisi header ditambah akhir body), sehingga wrapper audit/precheck yang mendahului migration body tidak dapat memancarkan operasi table/column/constraint yang sama lebih dari sekali. Parser column mengonsumsi identifier `[]`/`""` sebelum tokenisasi definition agar nama yang mengandung spasi tidak pecah menjadi nama dan tipe palsu.
- Setiap parsed table/index statement membawa source offset. Stable sort berdasarkan offset dijalankan sebelum reducer state sehingga add→rename→alter, drop→re-add, dan drop→recreate index diproses sesuai urutan SQL asli. Reducer menyatukan mechanical drop+replacement bernama sama menjadi satu `ALTER_*` row.
- `TableState` dan `IndexState` disimpan per schema+table sepanjang semua migration. Handler `DROP_TABLE` dan `CREATE_TABLE` menginisialisasi ulang map columns/constraints/dropped-constraints/indexes/dropped-indexes; ini memutus state incarnation lama sebelum lifecycle `REINTRODUCED`. Tes longitudinal v1→v7 memverifikasi type, nullability, default, dan index before-state pada v7 berasal dari v6.
- Writer me-resolve existing ancestor serta symlink/junction lalu menolak output path yang berada di dalam `migrations/` sebelum scan/write untuk mempertahankan boundary read-only pada salah konfigurasi CLI.
- `handler`: API boundary, auth `X-API-Key`, request/response contract, error mapping.

## 3. Kontrak Data dan Enum Internal

## 3.1 Struct utama

- `AppConfig`
- `MigrationFileMeta`
- `ObjectKey`
- `ObjectState`
- `DesiredObject`
- `PlannedOperation`
- `OutputReport`
- `OutputItem`
- `getobject.AppConfig`
- `getobject.ExportSummary`
- `cmd/rza_get_object_db.cliRunSummary`

## 3.2 Enum terkunci

Enum harus disimpan sebagai constant set dan divalidasi ketat.

- `Severity`: 0,1,2,3
- `Decision`: run, ok, skip, warn, fail
- `Action`: seluruh whitelist action pada spec
- `Reason`: seluruh whitelist reason pada spec
- `Category`: types,tables,indexes,functions,views,procedures,patch

Rule: nilai di luar whitelist langsung fail.

## 3.3 `OutputReport`

`OutputReport` harus merepresentasikan format final:

- `fail.count`
- `fail.items[]`
- `pass.count`
- `pass.items[]`
- `new_migrations.naming.pattern`
- `new_migrations.naming.line_endings`
- `new_migrations.files[]`
- `notes`

## 4. Alur Command

## 4.1 `fmt`

1. Discover file `db_object/**` dengan whitelist path.
2. Validasi forbidden token scan (lexical-safe).
3. Jalankan canonicalizer in-place.
4. Validasi kontrak file-content minimal.
5. Tulis report ke `out/output.json`.
6. Exit berdasarkan `fail.count`.

## 4.2 `verify`

1. Validasi struktur repo (`db_object`, `patch`, `migrations`).
2. Parse metadata filename migration.
3. Validasi pairing up/down.
4. Validasi forbidden SQL.
5. Validasi grammar SQL subset per kategori.
6. Validasi wrapper/audit requirements.
7. Validasi patch integrity dan duplication constraints.
8. Tulis report dan exit.

Catatan domain `db_object`:

- Verifier harus mendeteksi layout legacy `db_object/schema/*` sebagai warning `LEGACY_DB_OBJECT_LAYOUT_DETECTED`.
- Jika object sama ada di path canonical+legacy dengan isi berbeda, verifier fail `DB_OBJECT_DUAL_PATH_CONFLICT`.
- Jika isi sama, verifier warning `LEGACY_DB_OBJECT_DUPLICATE_CANONICAL`.

## 4.3 `gen`

1. Jalankan `fmt` fase internal.
2. Jalankan validate preconditions.
3. Reconstruct history dari `migrations/*.up.sql`.
4. Build desired state dari `db_object`.
5. Hitung semantic diff.
6. Build operation plan by category.
7. Emit migration pairs ke staging area.
8. Merge patch generation (selalu terakhir).
9. Verify full terhadap staging result.
10. Jika verify pass, commit write ke `migrations/`.
11. Tulis `out/output.json` final.
12. Jika fail, discard staging dan exit 1.
13. Jika `EnablePerfNotes=true`, tulis telemetry durasi fase ke `notes.perf_ms` (non-functional, tidak mengubah decision).
14. Jika `EnableParallelRead=true`, tahap collect source (`db_object` + `patch`), history preload (`migrations` parse/read), dan scan forbidden SQL per-file dieksekusi worker-pool bounded lalu di-merge by index/path; reducer state/report tetap serial untuk menjaga output deterministic.
15. Jika `EnableFailFast=true`, verify flow boleh berhenti setelah fail pertama untuk mempercepat feedback lokal; default tetap `false` agar laporan lengkap.

## 4.4 `rza_get_object_db` (`export`/`dry-run`/`preview`)

1. Validasi executable-name lock (`rza_get_object_db(.exe)`).
2. Load config compile-time (`cmd/rza_get_object_db/config.go`).
3. Jika mode `preview`, format sample SQL lokal tanpa koneksi DB.
4. Jika mode export:
   - query table/index source SQL Server
   - query module definition dari `sys.sql_modules`
   - normalisasi header programmability menjadi `CREATE OR ALTER`
   - format modul via `tsqlast` backend default `lite` (`antlr` fallback ke `lite`)
   - write hasil ke `db_object/*` (atau hitung delta saja untuk `-dry-run`)
5. Tulis ringkasan run ke `out/get_object_summary.json`.
6. Warning module parser/formatter bersifat non-fatal (dicatat di summary), error koneksi/query/io bersifat fatal (exit 1).

Kontrak unsupported table:

- Jika `SkipUnsupportedTables=true`, unsupported table feature (computed/sparse/filestream/type unsupported) dicatat sebagai warning `TABLE_EXTRACT_SKIPPED_UNSUPPORTED` dan export lanjut.
- Jika `SkipUnsupportedTables=false`, kasus yang sama menjadi fatal error.
- CHECK/FK disabled atau untrusted bukan alasan skip seluruh table: logical
  definition tetap ditulis dan observed state dicatat sebagai
  `CONSTRAINT_STATE_REVIEW_REQUIRED`.
- Opsi fisik non-default backing index PK/UQ dicatat sebagai
  `KEY_PHYSICAL_METADATA_NOT_CAPTURED`; canonical table snapshot tetap berisi
  logical PK/UQ.
- `-require-healthy-constraints` membuat disabled/untrusted finding menghasilkan
  exit non-zero setelah snapshot dan summary selesai ditulis; extractor tidak
  mengubah database.

Kontrak initial-table migration:

- Jika target `CREATE TABLE` belum ada, seluruh DDL dijalankan.
- Jika target sudah ada, profile `bootstrap-compatible` mewajibkan setiap kolom
  source tersedia dengan base type, identity-kind, dan computed-kind yang
  kompatibel.
- Extra column, length, precision/scale, nullability, collation, default,
  PK/UQ/CHECK/FK, runtime constraint state, dan physical index options tidak
  memblokir skip initial `CREATE TABLE`.
- Pelonggaran ini tidak berlaku pada explicit `ALTER`, `DROP`,
  `ADD/DROP CONSTRAINT`, atau `CREATE/DROP INDEX`.

## 4.5 `rza_migrate_rest`

1. Validasi executable-name lock (`rza_migrate_rest(.exe)`).
2. Load runtime config YAML (`config.yaml`).
3. Resolve path `repo_root` + `migration.migrations_dir`.
4. Inisialisasi service migrate:
   - factory `golang-migrate` (`file://...` + `sqlserver://...`)
   - DB lock manager (`dbo.usp_mig_try_get_lock`) jika enabled.
5. Expose endpoint:
   - `GET /healthz`
   - `GET /api/v1/migrations/version`
   - `POST /api/v1/migrations/up` (`steps` opsional)
   - `POST /api/v1/migrations/down` (`steps` opsional)
   - `POST /api/v1/migrations/run-file` (`file` wajib)
6. Enforce auth (`X-API-Key`) saat `security.enable_api_key=true`.
7. Enforce safety:
   - serialisasi operasi migrate (mutex in-process)
   - lock DB session
   - default config secure-by-default (`allow_down=false`, `allow_run_specific_file=false`)
   - timeout operasi best-effort (hardcoded; timeout context untuk lock/preflight)
   - `run-file` harus single-step deterministic:
     - `.up.sql` hanya jika file adalah next version dari current state
     - `.down.sql` hanya jika file match current version

## 5. Desain Scanner dan Ordering

- Semua file list wajib di-sort lexicographic sebelum diproses.
- Dilarang proses langsung dari map iteration.
- Scanner mengembalikan struktur:
  - absolute path
  - repo-relative path
  - category/source type
  - file bytes (opsional lazy load)

Path rules:

- `migrations/` tidak boleh punya subfolder.
- `patch/<folder>/` harus tepat tiga file.
- Nama file object harus `dbo.<Name>.sql` tanpa spasi.

## 6. Desain Parser

## 6.1 Filename parser (tanpa regex)

Strategi parser:

1. Validasi suffix `.up.sql` atau `.down.sql`.
2. Trim suffix, split token by `_`.
3. Parse token version dan validasi 17 digit.
4. Branch schema vs patch:
   - schema: token[1] category whitelist, token terakhir `v<k>`, object token charset `A-Za-z0-9_.#`
   - patch: token[1] wajib `patch`, token terakhir `v<k>`, patch_name non-empty
5. Return `MigrationFileMeta`.

Error parser adalah severity=3.

## 6.2 SQL lexical scanner

Scanner lexical dipakai untuk:

- forbidden token scan aman dari false positive di string/comment
- formatter safe mode
- parser subset tokenization

State machine minimal:

- `Normal`
- `SingleQuotedString`
- `LineComment`
- `BlockComment`

## 6.3 SQL parser subset MVP

Parser minimal yang wajib dikenali:

- table snapshot `CREATE TABLE dbo.X (...)`, termasuk computed column `<name> AS <expression> [PERSISTED]`
- constraints PK/UQ/DF/CK/FK (legacy naming allowed)
- index snapshot `CREATE [UNIQUE] INDEX ... ON dbo.X(...) [INCLUDE (...)]`
- programmability snapshot `CREATE OR ALTER VIEW/FUNCTION/PROCEDURE dbo.X`
- migration wrappers tables/indexes/patch

Parser wajib reject features out-of-scope (temporal, columnstore/XML/spatial index, dll) sebagai fail.

## 6.4 Formatter canonicalizer

Transformasi yang diperbolehkan:

- normalize LF
- trim trailing spaces
- ensure final newline

Transformasi yang dilarang:

- ubah isi literal string
- ubah komentar
- paksa casing keyword

Kontrak formatter modul shared (`internal/tsqlast`):

- default backend `lite`
- `antlr` bersifat best-effort dan fallback ke `lite` jika backend unavailable/error
- output formatter harus deterministik antara extractor (`rza_get_object_db`) dan auto-indent jalur `rza_migrate_db gen`

Table canonical contract:

- trailing comma style untuk item kolom/constraint.

## 7. Desain Validation Engine

Validation dibagi domain:

## 7.1 Source validation (`db_object`)

- schema wajib `dbo`
- file-content object match
- inline default forbidden
- index file only for owned table
- no `DROP INDEX` di data master
- programmability wajib `CREATE OR ALTER`
- programmability 1 file = 1 object
- source contract programmability dapat dijalankan mode `new-only` (`EnableSourceContractForNewOnly=true`)
- legacy layout warning dan canonical-vs-legacy conflict check wajib aktif
- `constraints_naming.md` opsional; source validation hanya fail hard untuk constraint tanpa nama yang tidak dapat diperbaiki deterministik

Catatan unnamed dependency:

- Constraint unnamed diperbaiki otomatis ketika struktur table/jenis/target terbukti; bentuk ambigu fail pada verify dengan rule `UNNAMED_CONSTRAINT_NOT_ALLOWED`.
- Karena unnamed constraint sudah ditolak di source validation, fase alter/drop column tidak boleh bergantung pada runtime discovery nama constraint.

## 7.2 Patch validation (`patch`)

- folder patch hanya tepat satu logical `README.md`, `up.sql`, `down.sql`; shared resolver `internal/fsx.ResolveLogicalFilesCaseInsensitive` mengindeks entry aktual secara case-insensitive untuk verifier, collector, cache, dan formatter agar konsisten Windows/Linux serta menolak duplicate case-only
- canonicalize patch_name sesuai algoritma
- patch source immutable setelah migration patch exist
- patch source must-not-disappear check

## 7.3 Migration validation (`migrations`)

- folder flat
- pair up/down lengkap
- version uniqueness pair-safe
- filename grammar valid
- category whitelist only
- forbidden SQL token check
- wrapper + audit call minimal tables/indexes/patch
- severity header comment ada pada tables/indexes/patch

## 7.4 Repo hygiene validation

- `constraints_naming.md`, `.gitignore`, dan `.gitattributes` adalah helper opsional.
- Bootstrap membuat helper yang belum ada tanpa menimpa file user.
- Checker melaporkan kekurangan helper sebagai warning non-blocking.

## 8. Desain State Reconstruction

Input:

- seluruh `migrations/*.up.sql` terurut filename.

Proses:

1. Parse metadata filename.
2. Parse SQL object/resulting intent.
3. Update `ObjectState` per `ObjectKey`.
4. Track conflict key `version+category+object+vk`.
5. Simpan hash normalized definition per object.

Khusus programmability hash:

- Normalizer hash mengabaikan style-only `;` dan bracket identifier di luar string/comment agar tidak terjadi false-drift.

Output minimal notes:

- `last_seen_version`
- `last_seen_vk`
- `normalized_hash`

Failure conditions:

- parse fail history
- duplicate conflict key
- category/object invalid

Semua failure di atas severity=3.

## 9. Desain Semantic Diff

## 9.1 Build desired state

- Parse seluruh source object setelah canonicalization.
- Build AST normalized per object.

## 9.2 Compare strategy

- Compare AST normalized, bukan raw text.
- Whitespace-only change diabaikan.
- Tentukan create/alter/drop per object.
- Untuk programmability, perubahan style-only bracket/semicolon diperlakukan no-op.

## 9.3 Policy mapping

- `DROP TABLE`: tidak generate DROP; generate precheck skip+warn `REQUIRES_MANUAL`.
- `DROP COLUMN`: drop direct dependency (DF/CK/UQ/FK/index), lalu drop column, plus warning `INDEX_RECREATE_REQUIRED`.
- `DROP INDEX`: precheck; jika not exists maka skip INFO.
- `ALTER COLUMN unsafe`: generate precheck block, skip+warn.
- Jika ALTER perlu drop dependency untuk lolos (MVP), jangan auto-drop; skip+warn `REQUIRES_MANUAL`.

## 10. Desain Emit Engine

Sebelum semantic hash dan typed diff, parser kolom membaca nama tipe (termasuk
schema-qualified user type), optional parameter list yang parenthesis-aware,
serta optional `COLLATE`. Parameter list dikanonisasi tanpa spasi di sekitar
koma/kurung, sehingga bentuk `DECIMAL (17, 5)` dan `DECIMAL(17,5)` identik,
namun `DECIMAL(17,5)` dan `DECIMAL(17,9)` memiliki hash dan `RawType` berbeda.
Base type berupa satu delimited identifier juga dikanonisasi (`[DECIMAL]` ke
`DECIMAL`) sebelum hash, typed diff, risk classification, dan catalog guard;
schema-qualified alias/CLR type tidak dilepas qualification-nya dan tetap
memakai policy konservatif.
Model yang sama dipakai risk assessment, dependency planning, renderer
`ALTER COLUMN`, dan rekonstruksi history agar precision/scale/length tidak
hilang di salah satu tahap.

## 10.1 Version allocator

- `base_ts` format `YYYYMMDDHHMMSS`.
- seq `000..999` untuk pair.
- `version = base_ts + seq`.
- jika pair > 999, fail hard.

## 10.2 Ordering emit

- category base order fixed: types,tables,indexes,functions,views,procedures,patch; dependency edges take precedence.
- within category: lexicographic object name, lalu `vk`.
- object same-run changes merged into one pair unless explicit policy split.

## 10.3 Template emitter per kategori

Tables/Indexes/Patch wajib:

- severity header comment
- deklarasi `@run_key = NEWID()`
- metadata `@version`, `@category`, `@file`
- precheck skip policy
- TRY/TRAN/COMMIT/CATCH/ROLLBACK/THROW
- audit call run/ok/fail

Programmability wajib:

- `CREATE OR ALTER`
- tanpa wrapper wajib

## 10.4 Down emitter

- Down mengarah ke state sebelumnya dari reconstruction.
- Down destruktif menulis warning `DATA_RESTORE_NOT_GUARANTEED`.
- Tables/indexes/patch down menggunakan wrapper transaksi.

## 10.5 Patch emitter

- Filename patch menggunakan canonical patch_name.
- Header patch:
  - `patch_folder`
  - `patch_name`
  - `readme`
- Body patch inserted as-is setelah LF normalize + trim trailing spaces.
- Patch duplication rule: patch_name existing -> skip generation.
- Write mapping patch ke `pass.items` output.

## 11. Desain Verify Integritas Final

Checklist verify final setelah emit:

- pairing lengkap
- version conflict tidak ada
- forbidden token tidak ada
- wrapper + audit call minimal valid
- programmability valid
- migrations flat
- output naming line ending LF

Jika verify final gagal:

- staging files dibuang
- tidak ada file baru yang committed ke `migrations/`
- report fail ditulis

## 12. Desain Output Writer

- Selalu tulis `out/output.json` untuk setiap command.
- Writer menghitung `fail.count` dan `pass.count` otomatis.
- `new_migrations.files` diisi hanya jika emit final committed.
- `notes` menampung reconstruction summary.
- Untuk extractor, selalu tulis `out/get_object_summary.json` (status, mode run, statistik delta, warning parsial module).
- `internal/output.PrintFailureDetails` menjadi formatter console bersama untuk checker dan generator. Seluruh fail item dicetak tanpa truncation, berurutan sesuai report, dan membawa field diagnosis yang tersedia. Pemanggil menulis ringkasan hasil dan detail failure ke `stdout` yang sama agar collector CI tidak mengacak dua stream; JSON tetap laporan canonical.

Contoh item wajib minimal:

- `rule`
- `severity`
- `file`
- `message`

## 13. Atomic Write Strategy

Untuk `gen`:

1. Emit file baru ke direktori staging sementara di `out/emit_staging/run-*`.
2. Bangun repo verifikasi sementara di `out/emit_verify/repo-*` yang berisi copy `db_object/`, `patch/`, hygiene files, existing `migrations/`, lalu overlay staged migration.
3. Jalankan verify terhadap repo verifikasi sementara, bukan terhadap folder final.
4. Jika verify pass, pindahkan staged files ke `migrations/` dengan `os.Rename` per file dan rollback file yang sudah committed jika commit gagal.
5. Jika verify fail, folder final `migrations/` tidak disentuh; staged files disalin ke `out/failed_emit/` untuk inspeksi lalu staging dihapus.

Tujuan: menegakkan rule no-emit-on-fail.

## 14. Locking Integration

Engine generator tetap offline, namun desain migration core audit wajib menyediakan `dbo.usp_mig_try_get_lock`.

Runtime recommendation di pipeline:

- jalankan lock procedure sekali di awal run migration.
- failure lock harus membuat job berhenti.

## 15. Error Handling dan Exit Contract

- Semua rule violation diklasifikasikan severity.
- Severity 3 selalu masuk `fail.items`.
- Exit code:
  - `0` bila `fail.count == 0`
  - `1` bila `fail.count > 0`
- `skip` hanya policy event, bukan mekanisme recovery SQL runtime error.

## 16. Test Design dan Coverage

## 16.1 Unit tests

- filename parser schema/patch
- patch canonicalization
- lexical scanner state machine
- SQL parser subset categories
- formatter safe transform
- version allocator limit

## 16.2 Golden tests

- deterministic emit ordering
- whitespace-only no emit
- drop column impacts index warning
- drop table manual-only flow
- patch duplication and integrity

## 16.3 Integration tests

- command `fmt`, `verify`, `gen`
- fail hard no emit on invalid input
- output.json schema and counts
- migration verify gate pass/fail fixtures
- `rza_get_object_db` mode `export`, `dry-run`, dan `preview` + kontrak `out/get_object_summary.json`

## 16.4 Rule-trace tests

Setiap section utama `rules.md` harus punya minimal satu test assertion:

- safety guards
- pairing/versioning
- db_object contracts
- destructive policy
- audit wrapper
- patch policy
- locking contract readiness
- repo hygiene and config baseline

## 17. Traceability Ringkas ke `rules.md`

- Rules 1-7: modul `fsx`, `sqlfmt`, `sqlparse`, `history`.
- Rules 8-13: modul `diff`, `emit`, `verify`.
- Rules 14-18: modul `emit`, `output`, policy engine.
- Rules 19-22: modul `config`, repo checks, CLI.
- Rules 23-25: hard lock enforcement + output patterns + runner assumptions.

## Report layout V2

Report publication memakai internal/migreport/layout.go: stage lengkap, root-aware inventory dan URL-escaped navigation, validate semua artifact, atomic replacement per file, hash-guarded stale cleanup, conservative legacy equivalence cleanup, manifest terakhir. Tidak ada recursive cleanup pada output pengguna atau transaksi atomic seluruh folder. Lihat [layout design](docs/report_output_layout.md).


## Audit object precision (2026-10-12)

SQLTokens mengekspos existing lexer; ReadTableMetadata/ReadIndexMetadata dan ModuleDeclarationName adalah shared observational projection, bukan pengganti migration AST. Report menggunakan catalog dan query depth/statement bounds, collision-aware aliases, function consumers dan lineage Resolution. NormalizeSQLToUTF8 dipakai in-memory pada source/history. Procedure header/EXEC memakai token offsets; target list ambigu ditolak dan missing columns diberi COLUMN_NOT_FOUND. Shared access grammar membedakan MERGE JOIN hint dari DML.

Detail bukti: docs/review_report_object_precision_20261012.md.
