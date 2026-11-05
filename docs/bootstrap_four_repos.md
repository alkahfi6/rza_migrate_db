# Bootstrap Empat Repo Service

Gunakan alur yang sama secara terpisah untuk setiap pasangan server dan initial
catalog (`bond-sqlreplicate`, `bond-sqltrsretail`, `reksa-sqlreksa2`, dan
`reksa-sqlreplicate`). Jangan menyalin `migrations` atau manifest lintas catalog.

1. Build binary dengan version/commit/build time melalui
   `scripts/build_dist_windows.ps1`.
2. Di repo service tujuan, salin `config.example.yaml` menjadi `config.yaml`, isi `database.host`, `port`, `user`, `password`, dan `initial_catalog` untuk pasangan server/catalog yang tepat, lalu jalankan `rza_get_object_db -profile bootstrap`. Config default dibaca dari working directory; `-config <YAML|JSON>` memilih file lain. Credential literal dengan karakter khusus di-encode oleh `net/url`; jangan URL-encode manual. JSON/env DSN lama tetap didukung sesuai prioritas di README. Jangan commit `config.yaml` atau menyalin credential antar catalog.
3. Review warning dan `unmanaged_objects` dalam
   `out/object_management_manifest.json`. Warning bukan bukti bahwa object aman
   diabaikan; object tersebut tetap dikelola manual sampai extractor mendukungnya.
4. Jalankan `rza_migrate_db bootstrap -repo .`. Default patch formatting adalah
   `off`, dan batch diteruskan sampai `pending_pairs=0`.
5. Untuk review bertahap, gunakan `draft`, review maksimal 999 pair, lalu `seal`;
   ulangi sampai tidak ada drift.
6. Commit pasangan `.up.sql`/`.down.sql`, desired state, dan evidence review.
   Jangan commit credential atau artifact lokal `out/`.
7. Deploy seluruh immutable history ke App-DB. File applied yang hilang atau
   berubah tetap menghentikan operasi dengan HTTP 409 dan kode error spesifik.

Manifest harus tetap tersedia sejak selesai extract sampai generation selesai.
Jika manifest hilang atau rusak, hentikan bootstrap dan ulangi extractor; jangan
mengarang entri unmanaged untuk melewati removal.


## Extractor catalog names (2026-11-05)

`rza_get_object_db` no longer applies repository filename/authoring identifier rules to catalog object names, including table types, tables, modules, and FK reference names. Spaces, hyphens, leading digits, Unicode, dots and escaped brackets are preserved in SQL identifiers. Module header parsing reads delimited identifier parts without stripping their contents. Filename-only escaping uses `%HH` for percent, control bytes and Windows-reserved characters (`<>:"/\|?*`); SQL retains the catalog name. Case-colliding filenames still stop before writes to prevent loss on portable filesystems. Other extraction limitations (encrypted/inaccessible definitions, unsupported metadata, schema scope) remain in force.

This is an explicit extractor-only exception requested for catalog acquisition. Checker apply, checker `-check`/Linux CI, and generator retain authoring name validation: successful extraction is not a promise that unusual names can generate migrations. Encoded filenames are not decoded/renamed by these tools and must be reviewed before authoring. Extractor never repairs migration history. Existing conventional names keep their original filenames/output. Full authoring release rebuild is required for shared-source identity; extractor is separately compiled because `build_all.bat` does not package it.

Proof: `internal/getobject/catalog_names_test.go` covers real snapshot apply/read-only plan, idempotence, literal SQL identity/module headers, portable filename escaping, collision avoidance between slash and literal percent sequences, warning identity, and unchanged migration history. Existing checker/generator fixture tests remain the boundary control. No new SQL migration support or migration golden is introduced.


## Continue-on-object-error extractor default (2026-11-05)

Extractor executable defaults to `-profile continue`: individual table/index rendering failures and skipped module/type definitions do not stop other objects. Existing failed-object table/index/module/type snapshots are preserved; skipped objects remain recorded in the unmanaged manifest, preventing an extraction omission from becoming a DROP request. `out/get_object_summary.json` retains warnings and structured items, and `out/get_object_issues.txt` is refreshed every run with object names/reasons (or an explicit clean result). No extra flag is needed for `.exe` users. Database connection/global catalog reads, unsafe filename collisions, and output/report write errors remain fatal because safe acquisition/publication cannot be completed. Reports contain diagnostics, not invented SQL for inaccessible/encrypted definitions.

Explicit `-profile strict` retains the previous fail-on-incomplete policy; `bootstrap` and the legacy skip flag remain supported. Shared extractor defaults also enable continuation (`SkipUnsupportedTables=true`). Per-table render errors are all eligible for preserve/report, rather than only errors matching an unsupported-message whitelist. No migration files are generated or repaired by extractor. Checker apply/read-only CI and generator keep existing authoring policy and consume unmanaged evidence; no safe-fix, migration semantics, or up/down support is changed. The authoring Windows/Linux release is rebuilt; the separately compiled extractor is supplied at `tools/rza_get_object_db.exe` for direct execution.

Tests: `TestContinueExportsGoodTablePreservesFailedSnapshots` exercises mixed good/bad catalog metadata, apply/read-only plan, repeat execution, both table/index preservation, immutable migration history, and explicit strict behavior. `TestExeDefaultContinuesAndWritesReadableIssues` proves no-flag executable CLI defaults, JSON/text evidence, success-with-warnings, and removal of stale warnings after a clean run.
