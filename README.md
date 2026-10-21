# RZA DB Engine Suite

Mulai memahami project dari [Engine Docs Guide](docs/engine_docs_guide.md).
Panduan tersebut menghubungkan [arsitektur](docs/architecture/overview.md),
[spesifikasi tiap engine](docs/engine_reference.md), [call tree](docs/architecture/call_tree.md),
[penjelasan setiap paket](docs/architecture/package_catalog.md), dan
[panduan agent/developer](docs/engineering/agent_onboarding.md).
Untuk mencoba alur tanpa database, ikuti [tutorial authoring offline](docs/tutorials/offline_authoring.md).

Untuk publikasi repository private, lihat [cakupan file GitHub](docs/github_repository.md).
Folder percobaan lokal dan kredensial dikecualikan; regression test, fixture QA,
serta release Windows/Linux tetap disertakan.

Suite ini berisi 5 aplikasi untuk workflow migration SQL Server berbasis source-of-truth:
- `rza_get_object_db`: extractor object DB -> `db_object/*`
- `rza_migrate_check`: checker developer (format + verify, tanpa generate migration)
- `rza_migrate_db`: formatter + verifier + generator migration
- `rza_migrate_rest`: REST trigger eksekusi migration SQL Server (`up/down/steps/run-file`)
- `rza_migrate_report`: generator report rekon histori perubahan database berbasis `migrations/` offline (zero DB access)

## Tujuan Aplikasi
1. Menjadikan `db_object/` sebagai desired state yang konsisten.
2. Menjaga `migrations/` tetap deterministic, audited, dan sesuai rules.
3. Memastikan perubahan schema bisa di-review di repo sebelum dieksekusi runner migration.
4. Menyediakan satu paket dokumentasi database di `docs/database/`, dengan current state dari `db_object/` dan histori dari `migrations/`, secara offline.

## Komponen Utama
- `cmd/rza_get_object_db/`
- `cmd/rza_migrate_check/`
- `cmd/rza_migrate_db/`
- `cmd/rza_migrate_rest/`
- `cmd/rza_migrate_report/`
- `internal/getobject/`
- `internal/engine/`
- `internal/verify/`
- `internal/migrest/`
- `internal/migreport/`
- `internal/handler/`
- `internal/history/`
- `internal/diff/`
- `internal/sqlparse/`

## Arsitektur Ringkas
1. `rza_get_object_db` membaca object `dbo` dari SQL Server.
2. Hasil extractor ditulis ke `db_object/types|tables|indexes|functions|views|procedures`.
3. `rza_migrate_check` memakai profile `authoring`: merapikan dan memvalidasi penuh `db_object`/`patch`, memeriksa integritas dasar `migrations`, dan menulis laporan tanpa mengubah migration.
4. `rza_migrate_db verify` memvalidasi kontrak folder, file, wrapper, pairing, dan rules SQL.
5. `rza_migrate_db gen` melakukan reconstruct history, diff state, dan emit migration up/down.
6. `rza_migrate_rest` menjalankan migration melalui REST API dengan API key + optional DB lock.

Pada diff table, parameter tipe kolom dibaca lengkap. Variasi spasi seperti
`DECIMAL (17, 5)` ekuivalen dengan `DECIMAL(17,5)`, tetapi perubahan menjadi
`DECIMAL(17,9)`, perubahan panjang character/binary, date-time precision, atau
base type tetap menghasilkan migration `ALTER COLUMN` sesuai risk guard. Nama
tipe bawaan yang didelimitasi extractor juga ekuivalen, misalnya
`[DECIMAL](25,2)` dan `DECIMAL(25,2)`; perubahan `[DECIMAL](25,2)` menjadi
`[DECIMAL](25,6)` menghasilkan migration guarded dengan precheck konversi
lossless, bukan `MANUAL_MIGRATION_REQUIRED`.

## Alur Operasional Harian
1. Jalankan `rza_get_object_db -dry-run` untuk cek delta tanpa write.
2. Jalankan `rza_get_object_db` untuk refresh snapshot source-of-truth ke `db_object/`.
3. Sebelum commit, double-click `tools\rza_migrate_check.exe` atau jalankan lewat terminal untuk format + validasi tanpa membuat migration.
4. Saat memang ingin membuat migration, jalankan `tools\rza_migrate_db.exe` (default `verify -> gen`).
5. Review `migrations/*.sql`, `out/output.json`, dan `out/get_object_summary.json` sebelum commit.
6. Jika diperlukan eksekusi migrate via service, jalankan `rza_migrate_rest -config config.yaml`.

## Struktur Repo yang Diharapkan
- `db_object/types/` untuk user-defined table type/TVP
- `db_object/tables/`
- `db_object/indexes/`
- `db_object/views/`
- `db_object/functions/`
- `db_object/procedures/`
- `migrations/` (flat)
- `patch/`
- `tools/` (berisi executable Windows, `license.key`, dan `linux-amd64/`)
- `out/`

Urutan dasar generation adalah `types -> tables -> indexes -> functions -> views -> procedures -> patch`, lalu dependency graph dapat memperketat urutan antar-object. Function diklasifikasikan sebagai scalar (`FN`), inline TVF (`IF`), atau multi-statement TVF (`TF`). Saat struktur TVP berubah, migration menjatuhkan function/procedure pemakai secara terbalik menurut dependency, mengganti TVP, lalu membuat ulang module tersebut dalam batch SQL Server yang valid.

## CLI
`rza_migrate_check` (direkomendasikan untuk pre-commit developer):
```bash
tools\rza_migrate_check -repo <path> [-check] [-profile authoring|full] [-migration-validation off|integrity|full] [-perf-notes true|false] [-fail-fast true|false] [-parallel-read true|false] [-parallel-workers <n>] [-pause-on-exit true|false]
```

Tanpa argumen, checker mencari repo dari working directory lalu folder executable (atau parent jika berada di `tools/`), memakai `profile=authoring` dan `migration-validation=integrity`, kemudian menunggu Enter sebelum jendela double-click ditutup. Mode `integrity` hanya menjaga flatness, filename, version, pairing, core pair, isi non-kosong, dan encoding migration; kontrak semantik migration historis hanya dijalankan pada profile `full`. Checker tidak mempunyai code path `gen`, tidak membuat migration, dan `new_migrations.files` selalu kosong. Saat dijalankan tanpa `-check`, checker juga mengonversi source `db_object/` dan `patch/` dari UTF-16/UTF-32 ber-BOM valid (atau UTF-8 BOM) menjadi UTF-8 tanpa BOM. Invokasi dengan argumen tidak pause secara default sehingga aman untuk hook/script.

Checker tidak memvalidasi deklarasi variable yang hanya berbeda kapitalisasi karena sensitivitasnya bergantung pada server collation; error tersebut dibiarkan muncul saat SQL dijalankan pada target. `constraints_naming.md`, `.gitignore`, dan `.gitattributes` adalah helper opsional dan kekurangannya hanya menghasilkan warning. Bootstrap membuat helper yang belum ada tanpa menimpa file user. Constraint SQL tetap harus mempunyai nama stabil, tetapi tidak ada prefix wajib; safe-fix memberi nama deterministik hanya saat struktur constraint dapat dibuktikan aman.

Untuk membangun binary Linux statis yang dipakai gate GitLab, jalankan dari Windows:

```powershell
scripts\build_migrate_check_linux.bat
```

Script tersebut adalah compatibility wrapper ke `build_all.bat`: ia membangun paket lengkap tiga Windows/tiga Linux binary, license, checksum, dan identity; bukan checker-only build. Job `rza_migrate_check_gate` memverifikasi complete release checksum, embedded build identity, dan source digest sebelum menjalankan checker authoring/integrity dalam mode read-only pada MR, `sit`, `uat`, dan tag, lalu menyimpan `out/output.json` sebagai artifact. Jika validasi gagal, semua item fail dicetak lengkap ke job log (`rule`, severity, file, object, lokasi migration, reason, message, dan suggestion jika tersedia), sehingga diagnosis tidak bergantung pada akses download artifact. Job tidak mengunduh package tambahan dan tidak terhubung ke MSSQL.

Untuk membangun paket lengkap Windows + Linux sekaligus dan membuat lisensi baru:

```powershell
.\build_all.bat "nama-owner" 3
```

Hasilnya adalah `tools/rza_migrate_db.exe`, `tools/rza_migrate_check.exe`, `tools/rza_migrate_report.exe`, ketiga binary Linux di `tools/linux-amd64/`, manifest SHA-256 untuk setiap binary, dan `tools/license.key`. Argumen bulan bersifat opsional dengan default 3 bulan. Argumen ketiga dapat dipakai sebagai direktori output staging/test. Jika owner tidak diberikan, batch meminta owner secara interaktif sehingga double-click tetap dapat digunakan. Build dilakukan di staging dan baru dipublikasikan setelah seluruh binary serta lisensi lolos pemeriksaan dasar.

Binary Linux yang dijalankan pada MR **tetap memerlukan lisensi**. Job memastikan `tools/license.key` tersedia di repository, lalu `rza_migrate_check` sendiri memvalidasi format, signature, dan masa berlaku lisensi tersebut sebelum pemeriksaan authoring dimulai. Manifest SHA-256 dan lisensi mempunyai fungsi berbeda: checksum mendeteksi binary rusak/berubah, sedangkan lisensi mengizinkan binary berjalan sampai tanggal kedaluwarsa.

Checksum yang disimpan bersama binary melindungi dari korupsi, bukan dari MR yang sengaja mengganti binary dan manifest sekaligus. Untuk runner yang menerima MR tidak tepercaya, organisasi wajib menyediakan checker dari image/artifact terproteksi yang dipin digest melalui central pipeline policy; kontrol tersebut tidak dapat dibentuk secara aman hanya dari file yang dapat diubah oleh MR yang sama.

Pipeline utama mengikuti template kantor `ocbc/devops/pipeline` dan menambahkan satu local include checker. Checker menggunakan special stage `.pre`, mendefinisikan image dan `before_script` sendiri, serta tidak mendefinisikan `stages`, `workflow`, `default`, atau global variable. Migration execution/locking bukan tanggung jawab gate authoring ini.

`rza_migrate_db`:
```bash
rza_migrate_db
rza_migrate_db bootstrap -repo <path>
rza_migrate_db draft -repo <path>
rza_migrate_db seal -repo <path>
rza_migrate_db gen -repo <path> [-format-patch off|safe|strict]
```

`rza_migrate_test` menjalankan migration terhadap SQL Server nyata menggunakan file kredensial lokal `mssql_credential.nfm` (`database_type`, `host`, `username`, `password`, `initial_catalog`). File ini tidak boleh di-commit.

```powershell
rza_migrate_test status -repo .
rza_migrate_test up -repo .
rza_migrate_test down -repo . -steps 1
rza_migrate_test reset -repo . -dry-run
rza_migrate_test reset -repo . -confirm-reset <nama_database>
```

`reset` hanya menghapus object non-system pada schema `dbo`, termasuk `schema_migrations`, dan menolak database sistem. Jalankan `-dry-run` terlebih dahulu; eksekusi reset memerlukan nama database persis sebagai konfirmasi.

Keterangan:
- `rza_migrate_db` tanpa command: default `verify` lalu `gen`.
- `rza_migrate_db fmt`, `gen`, dan default flow menormalkan source `db_object/` dan `patch/` yang dapat diidentifikasi aman ke UTF-8 tanpa BOM sebelum diproses. `verify` tetap read-only agar tidak menyembunyikan pelanggaran yang sedang diaudit. Gunakan `fmt -check` bila hanya ingin mendeteksi perubahan tanpa menulis file.
- Saat `gen` atau default flow menemukan constraint dengan definisi identik tetapi nama berbeda dari state migration terakhir, nama pada `db_object/tables` otomatis dikembalikan ke nama history. Jenis, kolom/urutan, clustering, expression, dan reference tetap dibandingkan; perubahan definisi nyata tidak dinormalisasi dan tetap diproses sebagai migration.
- `-legacy-audit-mode` hanya menerima `fail` (strict-only).
- `-gen-auto-indent` default `true`; set `false` untuk mematikan auto-indent saat `gen`.
- Formatting patch terpisah dari formatting object dan default `off`. Mode `safe` melewati file yang tidak dapat dibuktikan aman tanpa menggagalkan generation; `strict` mempertahankan hard-fail lama.
- `bootstrap` menjalankan initial generation secara batch untuk repo baru. Gunakan extractor `-profile bootstrap` agar object yang tidak dapat direpresentasikan dipertahankan dan dicatat sebagai unmanaged.
- `draft` menulis calon migration ke `out/draft_migrations`; `seal` memverifikasi lalu mempublikasikan byte hasil review tanpa overwrite. Batas per batch tetap 999 pair; seal batch sebelum membuat draft berikutnya.
- Build version, commit, build time, policy aktif, dan ringkasan warning/failure dicatat di console serta `out/output.json`.
- Folder patch wajib berisi tepat satu logical `README.md`, `up.sql`, dan `down.sql`. Kapitalisasi nama file diterima secara case-insensitive dan file aktual dibuka berdasarkan hasil enumerasi directory agar perilaku Windows dan Linux sama; duplikat case-only atau file tambahan tetap gagal.
- `fmt -check` menghitung perubahan format tanpa menulis `db_object`; exit code `1` dengan rule `FMT_CHECK_CHANGES_REQUIRED` jika ada file yang perlu dirapikan.
- Formatter mempertahankan indentasi/layout comment dan multiline literal; keyword di dalam comment tidak ikut membentuk indentation stack.
- `-perf-notes` default `false`; set `true` untuk menulis telemetry durasi fase ke `out/output.json.notes.perf_ms`.
  - mencakup breakdown domain verify: `verify.domain.migrations|patch|db_object|wrapper`.
  - pada `gen`, ada flag note `gen_precheck_history_cache_hit=true|false` untuk indikasi reuse hasil reconstruct history dari precheck.
  - formatter memakai cache metadata `out/fmt_cache.json`; note terkait: `fmt_cache_hits`, `fmt_cache_misses`, `fmt_cache_saved`.
- `-fail-fast` default `false`; set `true` untuk menghentikan verify flow setelah fail pertama (mode lokal cepat, CI disarankan tetap `false`).
- `-parallel-read` default `true`; set `false` jika ingin paksa mode serial.
- `-parallel-workers` default `8`; jumlah worker untuk `-parallel-read` (minimal `1`, effective worker di-cap oleh CPU dan jumlah task).
- env fallback: `RZA_GEN_AUTO_INDENT=true|false`.
- env fallback: `RZA_PERF_NOTES=true|false`.
- env fallback: `RZA_FAIL_FAST=true|false`.
- env fallback: `RZA_PARALLEL_READ=true|false`.
- env fallback: `RZA_PARALLEL_WORKERS=<n>`.

`rza_get_object_db`:
```bash
rza_get_object_db
rza_get_object_db -dry-run
rza_get_object_db -profile bootstrap
rza_get_object_db -config .\extractor.json -dry-run
rza_get_object_db -preview-input <module_sql_directory> -preview-output out/preview_modules
```

`rza_migrate_rest`:
```bash
rza_migrate_rest -config config.yaml
```

Endpoint utama:
- `GET /healthz`
- `GET /api/v1/migrations/version`
- `POST /api/v1/migrations/up` (body opsional `{"steps":<n>}`)
- `POST /api/v1/migrations/down` (body opsional `{"steps":<n>}`)
- `POST /api/v1/migrations/run-file` (body `{"file":"<migration_filename>"}`)
- Header auth: `X-API-Key: <security.api_key>`
- default hardening REST:
  - `migration.allow_down=false`
  - `migration.allow_run_specific_file=false`
  - timeout operasi selalu best-effort (dipakai untuk lock/preflight, bukan jaminan cancel migrate in-flight)
  - `db_lock` opsional (default nonaktif); jika diaktifkan gunakan `migration.db_lock.timeout_ms=0` agar otomatis ikut `migration.operation_timeout_seconds`

Saat berjalan, extractor akan menampilkan progress phase (`init/connect/load/render/apply/done`) di stdout.
Extractor juga menulis summary run ke `out/get_object_summary.json` agar hasil tetap bisa dicek saat aplikasi dijalankan via double-click.

## Konfigurasi Extractor (`rza_get_object_db`)
Konfigurasi sekarang compile-time (bukan JSON runtime):
- default compile-time config: `cmd/rza_get_object_db/config.go`
- optional runtime JSON config: `rza_get_object_db -config .\extractor.json`
- ubah nilai berikut lalu rebuild:
- `ConnectionString`
- `RepoRoot`
- `DbObjectDir`
- `Schema` (wajib `dbo`)
- `ConnectTimeoutSeconds`
- `QueryTimeoutSeconds`
- `RenderWorkers` (jumlah worker untuk render snapshot object)
- `WriteWorkers` (jumlah worker untuk write/delete file snapshot)
- `StripModuleComments` (hapus komentar `--` / `/* */` pada view/function/procedure hasil extract)
- `ReindentModuleSQL` (rapihkan indentasi blok `BEGIN ... END` untuk view/function/procedure)
- `IndentSize` (jumlah spasi per level indent, default `2`)
- `ModuleFormatterBackend` (`lite` atau `antlr`, default `lite`; `antlr` saat ini fallback otomatis ke `lite`)

Catatan export modul:
- Jika ada object view/function/procedure yang tidak bisa dinormalisasi (mis. definition kosong/encrypted/header invalid), extractor tidak menghentikan seluruh run.
- Object bermasalah akan di-skip dan dicatat di `out/get_object_summary.json` pada `export.warning_count` + `export.warnings[]`.
- Extractor menulis `out/object_management_manifest.json`. Engine memakai daftar unmanaged ini agar snapshot yang sengaja dipertahankan tidak dianggap sebagai permintaan `DROP`.

`rza_migrate_report`:

Object scan juga mengenali target DML dengan `TOP`, source `MERGE USING`, `APPLY`, comma join, serta sink `SELECT INTO`/`OUTPUT INTO`. `MERGE` action bukan statement DML terpisah. Identifier quoted/escaped dan spasi di dalam nama dipertahankan; alias tidak bocor ke subquery. Akses dan static `EXEC` yang berulang pada baris yang sama tetap dicatat. `OUTPUT` tanpa sink adalah client result; SELECT pada cursor atau cabang UNION dari INSERT/SELECT INTO bukan client result tambahan. Mapping UPDATE berhenti sebelum klausa FROM dan tidak menyeberang statement; INSERT multirow, TOP/DISTINCT, atau compound query belum diberi exact expression mapping. CTE dengan kapitalisasi berbeda dari referensinya menjadi `AMBIGUOUS` karena offline snapshot tidak membuktikan collation. Parser tetap scanner konservatif, bukan seluruh grammar T-SQL. Detail review dan batas cross-tool ada di [review object scan](docs/review_object_scan_20261001.md).

Perbaikan lexical/dependency memakai `internal/sqlparse` bersama dengan generator: hanya urutan dependency baru/removal yang dapat terkoreksi, bukan semantic SQL atau migration history. Extractor tetap memakai catalog SQL Server, checker interactive/`-check` tetap tidak menghasilkan migration, dan CI tetap read-only. Semua binary Windows/Linux diperbarui lewat full release build.

```bash
rza_migrate_report [-repo <path>] [-out <output_directory>] [-docs <documentation_directory>] [-pause-on-exit true|false]
```
Generator laporan offline tanpa koneksi database atau eksekusi SQL. Histori perubahan direkonstruksi dari `migrations/`; dokumentasi desired-state dibaca dari `db_object/` tanpa memodifikasi sumber.
- Pair bootstrap `000000__core__mig__audit_init` tetap dipindai sebagai infrastruktur rilis, tetapi dikecualikan dari analisa history aplikasi dan seluruh row audit report agar tabel internal `mig_event` tidak tercampur dengan perubahan schema bisnis. `report_summary.csv`, `report.json`, dan banner console memisahkan jumlah pair yang dianalisis dari `excluded_pairs`, sehingga jumlah file fisik tetap dapat direkonsiliasi.
- Mode double-click: executable mencari repo secara otomatis (folder executable, cwd, parent folder ber-`migrations/`) dan menunggu Enter sebelum jendela ditutup.
- Mode CLI / CI: `-pause-on-exit false` secara default untuk script non-interaktif.
- Nama migration divalidasi dengan parser strict yang sama dengan engine (`17 digit`, kategori whitelist, `dbo.<object>`, dan `v{k}` dengan `k >= 1`); file non-canonical tetap terlihat sebagai warning dan tidak diam-diam dianggap migration valid.
- Pair yang memiliki isi `up.sql` tetapi `down.sql` kosong atau tidak memuat inverse operasi yang sesuai diberi `round_trip_status=FAIL`, status migration/report `WARNING`, dan exit tetap `0` sesuai kontrak warning.
- Satu `ALTER TABLE` dapat memuat banyak kolom dan constraint sekaligus. Report memecah setiap item top-level, mempertahankan urutan SQL lintas `ADD`/`DROP`/`ALTER`/`sp_rename`, serta mereduksi pola mekanis drop+recreate constraint/index menjadi satu perubahan semantik dengan nilai before/after.
- Setiap blok `CREATE TABLE` dipancarkan tepat satu kali, termasuk saat berada setelah wrapper audit/precheck migration; kolom dan constraint di dalamnya hanya muncul sebagai turunan dari blok tersebut sekali. Identifier kolom SQL yang dibatasi `[]` atau `""`, termasuk yang mengandung spasi, dipertahankan utuh.
- Rekonstruksi tidak dibatasi pada versi berdekatan: riwayat longitudinal v1 sampai v7 diuji, termasuk perubahan campuran, `DROP TABLE`, reintroduksi, lalu perubahan lanjutan. `CREATE TABLE`/`DROP TABLE` mereset state kolom, constraint, dan index agar nilai lama sebelum penghapusan tidak bocor ke versi baru.
- Direktori output audit (`-out`, default `docs/database/migrations/`) dan dokumentasi (`-docs`, default `docs/database/`) wajib berada di luar `migrations/` dan `db_object/`; target yang melanggar batasan ini ditolak sebelum file dibuat. Nilai `-out`/`-docs` relatif selalu dihitung dari repository root, sehingga hasil tidak berpindah ketika command dijalankan dari folder lain. `report.json` memuat resolved output roots dan manifest `artifacts[]` konkret; `reports[]` tetap tersedia sebagai inventory path relatif terhadap documentation root; `artifacts[]` memuat root setiap file secara eksplisit.
- Layout V2: `-docs` menjadi root seluruh paket; jika `-out` tidak diberikan, audit mengikuti `<docs>/migrations/`. CSV berada di `datasets/`. `out/` hanya menampung output sementara tool lain dan tidak menerima hasil report default. Lihat [kontrak layout dan upgrade](docs/report_output_layout.md).
- Publication memvalidasi seluruh target sebelum menulis; target bentrok, symlink pada managed child, file manual, atau output yang diedit manual ditolak secara actionable. Stale profile hanya dihapus jika path dan SHA-256 cocok dengan manifest run sebelumnya. Legacy file dihapus hanya jika setara dengan hasil regenerated; file yang tidak terbukti sama dipertahankan.
- Baris data diurutkan deterministik. `run_at` dan `generated_at` adalah metadata waktu eksekusi sehingga memang berubah antar-run.
- Output audit laporan tersimpan di `docs/database/migrations/` (atau `-out`):
  1. `datasets/report_summary.csv`
  2. `datasets/migration_summary.csv`
  3. `datasets/object_history.csv`
  4. `datasets/table_changes.csv`
  5. `datasets/index_changes.csv`
  6. `datasets/programmable_changes.csv`
  7. `datasets/type_changes.csv`
  8. `report.json`
  9. `README.md`: ringkasan, findings, tautan datasets, dan tautan kembali ke dokumentasi utama.
- Output dokumentasi desired-state tersimpan di `docs/database/` (atau `-docs`):

  `table_unique_ids.md` juga memuat **Suspected JOIN identifiers** untuk kolom yang dipakai dalam static equality JOIN meskipun tanpa index/`IDENTITY`, termasuk pada table dengan PK lain. Evidence memuat canonical column mapping, source view/procedure, occurrence count, index/key dan IDENTITY. Nama non-ID seperti `Code` diberi `OBSERVED_JOIN_COLUMN_SIGNAL`; nama mirip ID/key memakai `OBSERVED_JOIN_ID_SIGNAL`. Status ini tidak membuktikan unique row ID atau merekomendasikan key baru. Filtered unique tetap `CONDITIONAL`; composite key tidak membuktikan uniqueness per kolom. Declared FK saja tidak dihitung sebagai observed JOIN. Fitur ini hanya report: extractor/checker/generator/CI tidak mendapat rewrite atau migration baru; binary diperbarui lewat full release build.

  1. `identifiers/table_unique_ids.md`: inventaris enforcement unik (`PRIMARY_KEY`, `UNIQUE_CONSTRAINT`, unfiltered `UNIQUE_INDEX`), uniqueness bersyarat, `IDENTITY` tanpa key, dan ID review signals dari JOIN/non-unique index. Signal bukan bukti uniqueness atau rekomendasi otomatis membuat key.
  2. `relationships/relational_erd.md`: seluruh diagram dalam satu file, disusun sebagai `Relationship Group` → connected `Relational Map`. Map dibentuk adaptif dengan batas maksimum 20 object, 24 relationship lines, 35 exact column mappings, dan 9.000 karakter Mermaid. Setiap node menunjukkan type object serta status key/index hanya untuk kolom yang ikut relationship; setiap garis mempertahankan exact column mapping. Posisi kiri/kanan tidak menyatakan arah, parent/child, atau cardinality. Dense pair otomatis memakai numbered `Mapping batch` agar tidak melewati text-size budget.
  3. `relationships/relationship_catalog.md`: audit trail exact `object.column = object.column`, jenis JOIN, jumlah kemunculan, source module, dan confidence. Nomor Relationship Group dan Relational Map sama dengan diagram.
  4. `relationships/module_lineage.md`: dependency module ke data object serta daftar module dynamic SQL yang wajib direview manual; dependency ini bukan relasi antartable.
  5. `README.md` dan `manifest.json`: halaman masuk dan inventory artifact/hash.
  6. `procedures/README.md`: index untuk seluruh procedure profile.
  7. `procedures/procedure_impact_catalog.md`: ringkasan lintas procedure untuk direct read/write, procedure call, function, view, dynamic SQL, dan unresolved reference.
  8. `procedures/profiles/dbo.<Procedure>.md`: satu profile per procedure berisi parameter/output contract, source-ordered process sequence, execution signals, branch-aware local execution flow, local object map, direct access beserta guard, exact INSERT/UPDATE mapping, key/index context, child-parameter mapping, direct call guard dalam blok per call, recursive procedure map, recursive write impact, dan dynamic-SQL review. Local execution flow memakai orientasi top-down agar procedure panjang tetap terbaca; transaction tanpa child call diringkas sebagai scope, sedangkan procedure routing dan error transfer tetap terlihat. Diagram besar dipecah maksimal 20 object/procedure per map tanpa menghapus detail dari dokumentasi; line break node memakai sintaks Mermaid yang renderable dan edge label selalu dikutip agar operator SQL tidak merusak rendering. Link profile memakai path URL-safe, termasuk procedure yang namanya memuat `#`.

Observed relationship hanya diambil dari predicate statis explicit `JOIN ... ON` atau legacy `FROM A, B WHERE ...` yang memetakan tepat satu resolved `alias.column` pada masing-masing sisi `=`. Nama kolom sama, operator non-equality, ekspresi ambigu, temporary/table variable, dan dynamic SQL tidak ditebak sebagai relasi. Module tanpa static reference yang berhasil di-resolve tetap dicantumkan sebagai coverage-review list.

Narasi dan panduan Markdown memakai Bahasa Indonesia yang natural, sedangkan istilah database dan header teknis memakai English yang umum (`Classification`, `Uniqueness`, `Source`, `Confidence`). Kode seperti `OBSERVED_JOIN` dan `IDENTITY` dipertahankan agar stabil untuk pencarian dan audit. Panduan classification diletakkan sekali di bagian bawah, bukan diulang sebagai template pada setiap baris. Daftar multi-value memakai pemisah titik koma, bukan HTML `<br>`.

Pada diagram relasi, garis solid berarti pasangan object hanya memiliki `DECLARED_FK`; garis putus-putus tanpa panah berarti mempunyai `OBSERVED_JOIN`. Label `left_column = right_column` menunjukkan kolom yang benar-benar terhubung; `[FK]`/`[JOIN]` hanya ditambahkan bila satu pasangan object memiliki kedua evidence agar mapping-nya tidak ambigu. Diagram tidak menyatakan cardinality bila source tidak membuktikannya. Shared object navigation dan object index membantu berpindah antarmap tanpa menyatakan ownership. Self-join tidak digambar sebagai garis dan ditempatkan pada `Self-join review`. Unowned `relationship_diagrams/` lama dipertahankan untuk manual review; tidak ada recursive cleanup terhadap folder dokumentasi pengguna. Fitur report ini hanya membaca `db_object`: extractor, checker apply/`-check`, migration generator, dan pipeline authoring tidak mengubah source atau migration karena dokumentasi tersebut.

Procedure profile juga bersifat konservatif dan catalog-backed. Scanner token shared mempertahankan source position dan statement boundary, sehingga `FROM`/`JOIN`/target DML tetap dikenali walaupun keyword dan object terpisah baris. CTE hanya berlaku dalam statement yang mendeklarasikannya; CTE unqualified tidak dapat menutupi table schema-qualified pada statement berikutnya. Nama static baru berstatus `RESOLVED_STATIC` bila cocok dengan table/view/function/procedure pada `db_object`; unqualified atau case-colliding name yang mempunyai lebih dari satu kandidat menjadi `AMBIGUOUS`. Literal dynamic SQL diparse ulang sebelum `RESOLVED_DYNAMIC`; target yang bergantung pada parameter, configuration, table, atau runtime metadata tetap `UNRESOLVED_DYNAMIC`. `SELECT INTO`, assignment `SELECT`, subquery/`EXISTS`, dan `INSERT ... SELECT` tidak dinyatakan sebagai direct client result set. `EXEC @return_code = dbo.Procedure` dikenali sebagai static call dengan return assignment, bukan dynamic SQL; `EXEC @sql` tetap masuk review dynamic. `RESOLVED_FROM_CALLER` hanya mencatat kandidat pada static call site dengan literal yang terbukti dan tidak menyatakan seluruh invocation runtime memakai kandidat tersebut. Temporary object, system object, external database, dan nama yang tidak ada di snapshot dilabel terpisah. Cursor dan built-in function tidak boleh menjadi false table/UDF dependency. Lineage menghapus dependency identik secara case-insensitive agar variasi kapitalisasi object tidak menggandakan row. Recursive traversal mempertahankan semua direct call edge, menghentikan cycle, dan tidak menduplikasi subgraph yang sudah dipetakan.

## Exit Code Contract
- `rza_migrate_check`: `0` jika format + verify sukses; `1` jika validasi gagal atau terjadi error. `out/output.json` tetap menjadi laporan utama.
- `rza_migrate_db`: `0` jika `out/output.json.fail.count == 0`, selain itu `1`.
- `rza_get_object_db`: `0` jika run sukses (warning module tetap non-fatal), `1` jika error hard (config/koneksi/query/io).
- `rza_migrate_rest`: `0` saat server berhenti normal, `1` saat startup/runtime fatal.
- `rza_migrate_report`: `0` jika status `OK` atau `WARNING`, `1` jika terjadi error atau kegagalan fatal.

## Real MSSQL Migration Test

`rza_mssql_test` menjalankan migration file ke SQL Server memakai `mssql_credential.nfm` (`database_type`, `host`, `username`, `password`, `initial_catalog`). File credential diabaikan oleh Git.

```powershell
go build -o .\rza_mssql_test.exe .\cmd\rza_mssql_test
.\rza_mssql_test.exe status
.\rza_mssql_test.exe up
.\rza_mssql_test.exe down -steps 1
```

Untuk mengulang pengujian dari database kosong, lihat dulu targetnya lalu gunakan konfirmasi eksplisit. `reset` hanya menghapus object user di schema `dbo`, termasuk metadata `schema_migrations`.

```powershell
.\rza_mssql_test.exe reset -allow-destructive -confirm-reset <nama_database> -dry-run
.\rza_mssql_test.exe reset -allow-destructive -confirm-reset <nama_database>
.\rza_mssql_test.exe up
```

`down -all` juga membutuhkan `-allow-destructive`.

## Cara Engine Mendeteksi Perubahan (`normalized_hash`)
Engine tidak membandingkan file mentah baris-per-baris. Engine membandingkan hash semantik object.

- Hash yang dipakai: `SHA-256` (hex string 64 karakter).
- Nilai hash disimpan sebagai `normalized_hash` di `out/output.json.notes.reconstruction`.
- Alur compare:
1. SQL dinormalisasi dulu (line ending/space/case/comment noise dibersihkan untuk compare semantik).
2. Untuk `tables`, ada canonicalization tambahan khusus snapshot kolom.
3. Untuk `indexes`, statement index dinormalisasi + diurutkan by index name sebelum hashing.

Konsekuensi:
- Perubahan kosmetik saja sering tidak memicu migration baru.
- Perubahan semantik object akan mengubah `normalized_hash` dan memicu plan `gen`.

## Executable Name Lock
- Binary wajib bernama:
  - `rza_migrate_db` / `rza_migrate_db.exe`
  - `rza_migrate_check` / `rza_migrate_check.exe`
  - `rza_get_object_db` / `rza_get_object_db.exe`
  - `rza_migrate_rest` / `rza_migrate_rest.exe`
  - `rza_migrate_report` / `rza_migrate_report.exe`
- Jika file executable di-rename, aplikasi akan fail saat start.

## Dokumentasi
- Peta dokumen dan urutan baca: `docs/engine_docs_guide.md`
- Rules utama: `rules.md`
- Requirement baseline: `spec.md`
- Functional design: `fsd.md`
- Development plan: `plan.md`
- Traceability: `docs/traceability/rules_traceability_matrix.md`
- QA gate: `docs/qa/qa_gate_engine_rules.md`
- Walkthrough modul: `docs/walkthrough.md`

## Distribusi Windows
Panduan penggunaan paket distribusi ada di:
- `dist/engine_windows_amd64/README.md`

## Validasi Kualitas
- Unit + integration test:
```bash
go test ./...
```
- Final hardening gate:
```powershell
powershell -ExecutionPolicy Bypass -File scripts/run_final_hardening_gate.ps1
```
- Parallel determinism gate:
```powershell
powershell -ExecutionPolicy Bypass -File scripts/validate_parallel_repeatability.ps1
```
- Parallel benchmark baseline:
```powershell
powershell -ExecutionPolicy Bypass -File scripts/benchmark_parallel_read.ps1
```

## Audit pembacaan object report (12 Oktober 2026)

Report kini membaca consumer/lineage dari view, procedure **dan function**, dengan
kolom `Resolution` pada `module_lineage.md`. JOIN mengikuti query/statement scope,
CTE, quoted identifiers, canonical object/column dan outer/hinted JOIN. DDL reader
menggunakan shared tokens untuk multiline index, key/FK dan type declaration.
UPDATE tanpa WHERE tetap mempunyai write mapping; target column yang tidak ada
pada table snapshot memakai `COLUMN_NOT_FOUND`. Target INSERT berqualifier tidak
boleh berubah menjadi kolom palsu. Dynamic SQL hanya menjadi resolved evidence
bila execution dan complete literal terbukti. Encoding Unicode dibaca di memory;
source identity mismatch/encoding tak dikenal gagal tanpa mengubah SQL.
Bukti corpus, temuan source, fixture, dan batas kemampuan ada pada
[audit object report](docs/review_report_object_precision_20261012.md).
