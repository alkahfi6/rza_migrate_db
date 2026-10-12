# Review rza_migrate_report dan pola pengenalan object kahfigo

Tanggal: 27 September 2026. Scope: review implementasi lokal, tanpa perubahan engine, source SQL, migration, atau release binary. Tidak memakai koneksi database. Bukti checkpoint sebelumnya tidak dianggap sebagai hasil validasi sesi ini.

Status 30 September 2026: rekomendasi prioritas 1–3 telah diimplementasikan melalui scanner token shared `internal/sqlparse`, semantic client-result classification, repo-relative output roots, dan concrete artifact manifest. Bagian temuan di bawah dipertahankan sebagai bukti kondisi sebelum perbaikan; validasi rilis final dicatat pada hasil pekerjaan implementasi, bukan pada checkpoint review awal ini.

Audit ulang 1 Oktober menemukan residual bugs pada lexical delimiter, TOP/MERGE/APPLY/sinks, query alias, repeated source occurrence, static EXEC, dan exact mapping boundaries. Perbaikan serta batas capability dicatat pada [review object scan](review_object_scan_20261001.md); status implementasi sebelumnya tidak berarti grammar sudah sempurna.

## Temuan

### P1 — Scope CTE menutupi object permanen

Lokasi: `internal/migreport/procedure_analysis.go`, `resolveProcedureCTEs` (baris 799).

CTE dikumpulkan untuk seluruh procedure. Nama terakhir setiap access kemudian dibandingkan dengan kumpulan ini, termasuk access schema-qualified dan statement lain. Contoh dengan catalog `dbo.T`:

```sql
CREATE OR ALTER PROCEDURE dbo.P AS
WITH T AS (SELECT 1 AS Id)
SELECT * FROM T;
SELECT * FROM dbo.T;
```

Probe menghasilkan kedua access sebagai `Object= dbo.T`, `Resolution=NOT_PERSISTED`. Access terakhir seharusnya tetap resolved permanent table. Ini menghilangkan impact tabel dari laporan dan menyamakan identity CTE dengan physical object.

Usulan: resolve identifier sebelum canonical catalog lookup, memakai scope statement/query; CTE harus menyimpan identity sendiri dan kedaluwarsa setelah statement. Schema-qualified object tidak boleh tertimpa CTE unqualified.

### P1 — Object access bergantung pada line break

Lokasi: `internal/migreport/procedure_analysis.go`, `extractProcedureAccesses` (baris 265).

Regex berjalan per baris. Hanya INSERT mempunyai sambungan baris khusus. Probe berikut menghasilkan `accesses=[]`, meskipun `dbo.T` tersedia di catalog:

```sql
CREATE OR ALTER PROCEDURE dbo.P AS
SELECT *
FROM
dbo.T;
```

Usulan: tokenize seluruh SQL lalu parse statement, mempertahankan source span untuk evidence dan nomor baris. Tambahan regex multiline per operasi tidak menyelesaikan masalah scope alias, nested query, dan statement boundary.

### P2 — Output contract salah untuk SELECT INTO

Lokasi: `internal/migreport/procedure_analysis.go`, `detectDirectResultSet` (baris 813).

Probe `SELECT * INTO #work FROM dbo.T;` menghasilkan `HasDirectResult=true`. Statement tersebut mengisi temporary table, bukan menghasilkan client result set. `SELECT` assignment, SELECT INTO, INSERT SELECT, subquery, OUTPUT INTO, dan client SELECT perlu dibedakan oleh statement semantics.

### P2 — Path default dan override memakai basis berbeda

Lokasi: `internal/migreport/report.go`, `Run` (baris 36–58).

Default dihitung dari repo root. Override `-out reports -docs documentation` diteruskan sebagai relative path sehingga mengikuti working directory proses. Memanggil repo yang sama dari folder berbeda dapat menempatkan override output di lokasi berbeda. Temuan ini berasal dari tracing kode; belum dibuat integration probe terpisah untuk path override.

Pemisahan CSV/JSON ke `out/report` dan Markdown ke `docs/database` sendiri sudah konsisten dengan README/spec, bukan bug. Usulan: definisikan relative override terhadap repo root, canonicalize menjadi absolute path sebelum validasi/tulis, dan tampilkan resolved roots. Perubahan ini mengubah kontrak CLI dan harus didokumentasikan.

### P2 — JSON inventory tidak menunjukkan lokasi artefak sesungguhnya

Lokasi: `internal/migreport/report.go:84`, `desired_state_writer.go:12`, `model.go` field `JSONReport.Reports`.

`report.json` mencampur basename CSV dan Markdown dari dua root; `procedures/profiles/*.md` hanya wildcard. Consumer tidak dapat menentukan lokasi dokumen untuk custom `-docs` dari daftar ini saja. Console sudah memakai root terpisah; masalah utama ada pada manifest JSON.

Usulan: tambahkan manifest konkret `{kind, path}` dengan basis path eksplisit dan daftar semua generated profile, mempertahankan field lama selama transisi. Publikasikan JSON setelah dokumen berhasil ditulis agar manifest tidak menyatakan keberhasilan artefak yang belum tersedia. Saat ini audit output ditulis sebelum Markdown; kegagalan penulisan Markdown dapat meninggalkan hasil parsial.

## Pola kahfigo yang layak diserap

| Pola dan bukti kode kahfigo | Manfaat untuk engine | Batas adopsi |
| --- | --- | --- |
| `internal/sqlparser/lexer.go`: token dengan posisi; `ast.go`: statement/block/IF/EXEC | Parsing tidak tergantung line break; evidence mengacu source span | Perlu adaptor ke parser engine, bukan parser kedua yang hanya dipakai report |
| `internal/sqlparser/cte.go`: `parseWithStatement` menyimpan dan memulihkan scope CTE | Memisahkan CTE/derived source dari tabel permanen | Audit nested query, recursive CTE, column alias, dan DML alias; pola ini bukan bukti seluruh grammar lengkap |
| `internal/sqlparser/ast.go`: SymbolTable; `const_eval.go`: constant/nonconstant invalidation | Mengenali dynamic SQL yang literalnya terbukti, tanpa menebak runtime target | Assignment dari query/parameter/branch berbeda tetap unresolved atau candidate bersyarat |
| `internal/sqlparser/parser_dynamic_controlflow.go`: branch-aware variants | Guard dan pilihan dynamic target lebih akurat | Candidate call-site tidak boleh dianggap berlaku untuk semua invocation |
| `internal/sqlparser/access.go`: klasifikasi SelectStmt, IntoTable, IsAssignment, OutputStmt, ExecStmt | Membedakan client result, temporary population, dan write effects | Audit OUTPUT INTO dan INSERT EXEC; jangan menganggap access classifier sebagai migration planner |
| `internal/metadata/schema.go`: SchemaStore dan optional SynonymStore; `file_provider.go`: offline provider | Pisahkan catalog identity/type lookup dari parsing; snapshot dapat dipakai tanpa DB access | Identity harus lengkap database/schema/object, ambiguity eksplisit; report tetap offline |

`kahfigo` bukan pengganti drop-in: fokusnya output shape/code generation procedure, sementara engine ini harus menjaga DDL semantics, constraint attributes, migration history, dan valid up/down. Jangan menyalin default tipe atau fallback metadata menjadi kebenaran schema migration.

## Desain yang disarankan

Gunakan jalur bersama: SQL source → token/source span → statement/query AST → scoped symbol resolution + catalog → evidence model → consumer report/checker/generator. Mulai dengan memperluas `internal/sqlparse`/`internal/tsqlast` yang sudah ada. Evidence model mencatat object identity, kind, operation, resolution reason, source span, scope, dan guard. `NOT_FOUND`, `AMBIGUOUS`, dan runtime unresolved harus tetap terlihat.

Live extraction tetap memakai metadata SQL Server: `internal/getobject/export.go` sudah membaca `sys.objects`, `sys.schemas`, `sys.sql_modules`; `export_batch.go` membaca metadata table/key/index/FK. Menambah AST tidak menggantikan catalog extraction. Metadata tambahan/synonym snapshot bisa dikaji terpisah, tetapi bukan klaim fitur yang sudah didukung. Function saat ini diinventarisasi namanya di report, sementara source module yang dianalisis hanya views/procedures (`BuildDesiredStateDocs`); perluasan function-body lineage adalah scope baru.

Prioritas: (1) statement boundary, CTE/alias scope, dan multiline access; (2) result-set semantics; (3) path/manifest publication; (4) constant propagation dan function-body lineage. Untuk layout Mermaid, source sudah memakai top-down dan chunking; posisi visual antar-renderer belum diperiksa dalam review ini. Urutan deterministik tidak menjamin posisi node identik.

## Acceptance dan batas cross-tool

- Tambah kasus permanen: ketiga probe di atas, CTE bernama sama dengan tabel, alias dipakai ulang pada statement berbeda, quoted identifiers, whitespace variants, case collision, dan unresolved dynamic SQL.
- Path: uji cwd root/tools/folder luar dengan repo sama, default/relative/absolute override, manifest konkret, kegagalan write, dan rerun idempotent.
- Perubahan analisis report tidak boleh memperbaiki SQL atau memengaruhi schema diff. Shared parser change harus diuji pada extractor, checker apply dan read-only `-check`, fmt/verify/gen; hasil checker harus mempertahankan migration bytes.
- Saat implementasi selesai: unit/integration/fixture, hardening gate, build_all, identity/source digest/checksum, dan Linux smoke gate. Review ini tidak mengubah release dan tidak menyatakan binary saat ini sudah diverifikasi.

## Validasi sesi ini

- `go test ./...` pada rza_migrate_db: PASS (termasuk engine dan fixture tests; sebagian memakai Go cache).
- Temporary `TestReviewProbe`: PASS sebagai diagnostic probe, merekam tiga perilaku salah di atas; bukan test penerimaan correctness. File probe dihapus setelah dijalankan.
- `go test ./internal/sqlparser ./internal/metadata` pada kahfigo: PASS.
- Tidak menjalankan live DB, render Mermaid, build_all, hardening PowerShell, atau Linux binary smoke; tidak ada perubahan implementasi untuk dirilis.

Dokumen ditambahkan: dokumen review ini saja. Dokumen dibaca/review: engine docs guide, README, rules, spec, FSD, plan, walkthrough (rujukan report/output), capability matrix, traceability matrix. Dokumen normatif sengaja tidak diubah karena ini temuan/rekomendasi, bukan perilaku baru yang telah diimplementasikan. QA gate tercantum sebagai kebutuhan validasi lanjutan; generated coverage, policy, dan fixture tidak diperbarui karena tidak ada perubahan rule atau acceptance fixture.
