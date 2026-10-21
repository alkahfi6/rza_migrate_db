# Referensi Spesifikasi Tiap Engine

Referensi ini merangkum kontrak aplikasi yang sudah ada. Requirement normatif tetap dimiliki [rules](../rules.md) dan [spec](../spec.md); desain rinci ada di [FSD](../fsd.md). Untuk urutan fungsi gunakan [call tree](architecture/call_tree.md). Contoh berikut memakai PowerShell dari root source project dan path target eksplisit.

## Prasyarat bersama

- Go sesuai `go.mod` (`go 1.25`) untuk menjalankan source; dependency versions mengikuti `go.mod`/`go.sum`.
- Target authoring mempunyai `db_object/{types,tables,indexes,views,functions,procedures}`, `patch/`, dan `migrations/` flat. Object memakai schema `dbo`; file index memakai nama table pemiliknya.
- `rza_migrate_db` dan checker memvalidasi license di jalur yang diresolusikan config. Salin license sah ke `tools/license.key` pada repo latihan. Jangan mengasumsikan `go run` melewati license.
- Repo discovery tanpa `-repo` mencoba working directory dan lokasi executable. Untuk script/agent gunakan target eksplisit dan `-pause-on-exit false` pada aplikasi yang mendukungnya.
- Flag dengan nilai string boolean memakai `-nama true|false`; flag Go boolean seperti `-check`/`-dry-run` dapat ditulis sendiri. Posisi command dan flag berbeda antar-utilitas; jangan menyalin flag satu tool ke tool lain.

## `rza_get_object_db`: akuisisi dan canonical snapshot

| Kontrak | Penjelasan |
| --- | --- |
| Input | Catalog SQL Server dan definition module; config compiled/env atau JSON `-config`. |
| Kategori | Table type/TVP, table, index, view, function FN/IF/TF, procedure dalam batas capability matrix. |
| Output | Canonical `db_object/*`, `out/get_object_summary.json`; export apply juga menulis management manifest. |
| Mutasi | Bisa membuat/mengubah/menghapus snapshot; tidak mengubah metadata/data SQL Server dan tidak menghasilkan migration. |
| Failure | Invalid config, connection, unsupported strict snapshot, unsafe render; CLI mengembalikan nonzero. Healthy-constraint gate dapat gagal setelah output ditulis. |

```powershell
go run ./cmd/rza_get_object_db -config C:\local\extractor.json -dry-run
go run ./cmd/rza_get_object_db -config C:\local\extractor.json
go run ./cmd/rza_get_object_db -config C:\local\extractor.json -profile bootstrap
go run ./cmd/rza_get_object_db -preview-input C:\local\sql -preview-output C:\local\preview
```

`-dry-run` memerlukan akses DB dan hanya menahan snapshot writes; CLI summary tetap ditulis. Preview memakai file lokal tanpa DB. Config runtime, termasuk target repo, dimiliki extractor; tool ini tidak mempunyai flag `-repo` seperti checker. Sumber credential melalui JSON/env dijelaskan pada [bootstrap](bootstrap_four_repos.md); jangan memasukkan credential nyata ke dokumentasi.

Default `-profile strict` menuntut snapshot lengkap. `-skip-unsupported-tables` atau profile bootstrap mengizinkan skip/preserve dengan evidence unmanaged. `-require-healthy-constraints` mempromosikan disabled/untrusted CHECK/FK menjadi exit nonzero setelah export. Jangan mengartikan skipped object sebagai object yang boleh di-DROP.

## `rza_migrate_check`: authoring check tanpa generation

```powershell
go run ./cmd/rza_migrate_check -repo C:\local\repo -pause-on-exit false
go run ./cmd/rza_migrate_check -repo C:\local\repo -check -profile authoring -migration-validation integrity -pause-on-exit false
go run ./cmd/rza_migrate_check -repo C:\local\repo -check -profile full -pause-on-exit false
```

| Mode | Perilaku |
| --- | --- |
| Apply, default | Safe format/fix source, reconcile equivalent constraint names, verify, report. |
| `-check` | Plan correction tanpa menulis authored SQL; pending correction memblokir. Report/cache diagnostic tidak sama dengan authored source. |
| `-profile authoring` | Default; source/patch full checks, migration default integrity. |
| `-profile full` | Migration validation wajib full; mencakup semantic history/wrapper checks. |
| `-migration-validation off` | Opt-out migration domain yang tersedia pada authoring; tidak sesuai command MR gate. |

Integrity mencakup flatness, filename/version, pair/core pair, nonempty content, dan encoding. Ia tidak setara audit seluruh SQL historis. `new_migrations.files` checker selalu kosong; safe corrections ambigu harus dilaporkan untuk penanganan manual. Exit `0` berarti tidak ada failure, exit `1` berarti error atau `fail.count > 0`.

## `rza_migrate_db`: authoring migration offline

| Command | Input/aksi | Efek output |
| --- | --- | --- |
| Tanpa command / `verify_then_gen` | Format, combined precheck, reconstruct, plan, emit | Bisa mengoreksi source dan menambah pair; default CLI dapat auto-batch. |
| `fmt` | Safe encoding/rewrite/canonical formatting | Mengubah source yang aman; tidak membuat migration. |
| `fmt -check` | Rencana format read-only | `FMT_CHECK_CHANGES_REQUIRED` memblokir bila perlu perubahan. |
| `verify` | Audit source/history/wrapper | Tidak memperbaiki source atau history; menulis report. |
| `gen` | Format + precheck + diff dari history | Pair baru setelah plan round-trip dan staging verify. |
| `bootstrap` | Seed helper/core yang belum ada, initial generation | Untuk adopsi repo baru, bukan perbaikan otomatis applied history. |
| `draft` | Jalur generation ke direktori draft | `out/draft_migrations`; belum publish ke history. |
| `seal` | Verifikasi dan publish reviewed draft bytes | Exclusive creation di history, tanpa overwrite; draft dihapus setelah sukses. |
| `license_gen` | Generate license dengan owner/months | Operasi packaging terpisah, bukan workflow migration harian. |

```powershell
go run ./cmd/rza_migrate_db fmt -repo C:\local\repo -check -pause-on-exit false
go run ./cmd/rza_migrate_db verify -repo C:\local\repo -pause-on-exit false
go run ./cmd/rza_migrate_db draft -repo C:\local\repo -pause-on-exit false
go run ./cmd/rza_migrate_db seal -repo C:\local\repo -pause-on-exit false
```

Generation order dasar: `types -> tables -> indexes -> functions -> views -> procedures -> patch`; dependency planning memperketat urutan, removal membalik dependency. Pair version memakai timestamp plus sequence tiga digit; `v{k}` adalah version per object, bukan DB runtime version. Satu batch maksimum 999 pair. Deterministic berarti hasil semantic/order stabil untuk input yang sama; filename timestamp dan metadata waktu run dapat berbeda.

Patch format default `off`; `safe` melewati unsafe input dengan diagnostic, `strict` memblokir. Generation dapat reconcile constraint name hanya bila semantics identik. Parameter tipe, urutan key, expression, reference, clustering, dan delta nyata tetap diperhitungkan. `verify` dapat gagal karena drift sebelum migration dibuat; gunakan generation workflow yang sesuai untuk mengusulkan pair, jangan memaksa verify menjadi auto-fix.

## Flag dan env authoring yang sering dipakai

| Flag | Env fallback | Default/arti |
| --- | --- | --- |
| `-perf-notes` | `RZA_PERF_NOTES` | `false`; telemetry durasi di `notes.perf_ms`. |
| `-fail-fast` | `RZA_FAIL_FAST` | `false`; full report lebih berguna untuk CI. |
| `-parallel-read` | `RZA_PARALLEL_READ` | `true`; bounded worker dan stable merge. |
| `-parallel-workers` | `RZA_PARALLEL_WORKERS` | `8`; effective count dibatasi CPU/task. |
| `-pause-on-exit` | `RZA_PAUSE_ON_EXIT` | Automatic untuk interactive double-click; gunakan false di script. |
| `-gen-auto-indent` | `RZA_GEN_AUTO_INDENT` | Generator saja, `true`. |
| `-format-patch` | `RZA_FORMAT_PATCH` | Generator saja, `off|safe|strict`. |
| `-legacy-audit-mode` | `RZA_LEGACY_AUDIT_MODE` | Generator saja, hanya `fail`. |
| `-auto-batch` | `RZA_AUTO_BATCH` | Generator combined/bootstrap; default aktif pada implicit mode/bootstrap. |
| `-profile` | `RZA_VALIDATION_PROFILE` | Checker, `authoring|full`. |
| `-migration-validation` | `RZA_MIGRATION_VALIDATION` | Checker, `off|integrity|full`. |

Tabel ini bukan seluruh flag; signature lengkap ada pada CLI source yang ditautkan di [katalog](architecture/package_catalog.md). Flag eksplisit mendahului env pada resolver di atas.

## `rza_migrate_report`: dokumentasi database offline

```powershell
go run ./cmd/rza_migrate_report -repo C:\local\repo -pause-on-exit false
go run ./cmd/rza_migrate_report -repo C:\local\repo -docs docs/database -out out/history-audit -pause-on-exit false
```

Input history menghasilkan rekonsiliasi perubahan; current `db_object` menghasilkan identifier, relationship, lineage, dan procedure impact documentation. Default docs root `docs/database`, audit root `<docs>/migrations`. `-docs` mengubah root dokumentasi, `-out` hanya audit root; keduanya diresolusikan relatif repo kecuali absolute. Output dilindungi agar tidak berada di dalam `db_object/` atau `migrations/`.

Layout, tujuh CSV, JSON, Markdown, ownership manifest, serta cleanup dijelaskan pada [report output layout](report_output_layout.md). Exit nonzero harus ditindaklanjuti meski sebagian report sudah ada. Output inference tidak membuktikan live database, uniqueness dari JOIN, atau runtime dynamic branch. `rza_migrate_report` saat ini tidak menjalankan license validation CLI; jangan menggeneralisasi policy checker ke seluruh binary.

## `rza_migrate_rest`: execution service

```powershell
go run ./cmd/rza_migrate_rest -config C:\local\config.yaml
```

Config YAML `migrest` terpisah dari authoring config. Service membaca directory migration dan database URL, membangun migrator, operation lock, journal, handler, serta telemetry. API schema lengkap dimiliki [swagger.yaml](../internal/handler/swagger.yaml).

| Endpoint | Aksi dan guard |
| --- | --- |
| `GET /healthz` | Liveness service; bukan bukti semantic equality DB dengan desired source. |
| `GET /api/v1/migrations/version` | Current version/dirty. |
| `GET /api/v1/migrations/pending` | Pending up files dengan urutan deterministik. |
| `POST /api/v1/migrations/up` | Seluruh pending atau steps positif. |
| `POST /api/v1/migrations/down` | Steps positif, hanya jika config mengizinkan. |
| `POST /api/v1/migrations/run-file` | Tepat next up/current down; bukan arbitrarily execute SQL. |
| `POST /api/v1/migrations/ack-manual` | Controlled recovery dengan config dan reason; baca runbook sebelum dipakai. |

API-key middleware melindungi route migration bila diaktifkan; production policy mewajibkannya. Mutex lokal dan configured DB lock mempunyai scope berbeda; external runner juga harus mematuhi locking contract. Applied-history evidence mismatch/dirty state bukan alasan mengedit checksum atau memaksa version. Lihat [REST hardening](policy/rest_api_hardening.md) dan [monitoring](operations/migration_monitoring.md).

## Utilitas live test dan release

`rza_migrate_test` menyediakan `create-db/status/up/down/reset` dengan `.nfm`, database override, dan optional catalog mapping pada temporary migration copy. `rza_mssql_test` menyediakan `status/up/down/reset`, default status, dengan guard `-allow-destructive` untuk full down/reset. Keduanya dapat mengubah SQL Server; tutorial onboarding tidak menggunakannya. Syntax dan exit contract tersedia pada [source masing-masing](architecture/package_catalog.md).

`build_all.bat` memublikasikan authoring/report release tiga Windows dan tiga Linux binaries berikut license/checksum/identity. Extractor/REST mempunyai jalur bundle tambahan melalui `scripts/build_dist_windows.ps1`; jangan menganggap complete six-binary release mencakup seluruh tujuh command source. Linux checker yang dipakai MR harus membawa source digest yang cocok, bukan implementasi terpisah. Detail gate ada pada [call tree release](architecture/call_tree.md) dan [production readiness](policy/production_readiness.md).

## Membaca diagnostic

`out/output.json` authoring memiliki `fail`, `pass`, `new_migrations`, dan `notes`. Baca `fail.items[].rule/file/object/message/suggestion` terlebih dahulu; warning dapat tampil di bucket pass sehingga `pass.count` bukan jumlah object yang terbukti benar. `new_migrations.files` mencatat up filenames; pasangan down tetap wajib. Console checker mencetak failure detail agar diagnosis tidak bergantung artifact download.

Ringkasan extractor, report JSON, response REST, DB `mig_event`, dan operation journal adalah kontrak berbeda. Jangan mencoba membaca semuanya sebagai `output.Report`. Simpan command, exit, input revision, dan batas validasi saat meneruskan hasil ke agent lain.
