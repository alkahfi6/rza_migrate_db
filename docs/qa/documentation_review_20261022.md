# Review Dokumentasi Engine — 2026-10-22

## Scope dan baseline

Baseline source: `1735ec494d8171d46ecc3ba11aa91984fd09d38c` (`Improve database object report precision`), branch `main`. Workspace tracked bersih pada awal review. Checkpoint percakapan sebelumnya tidak dipakai sebagai bukti keberhasilan; inventory dan validation diperiksa ulang dari checkout.

Tujuan perubahan: melengkapi pemahaman tujuan project, spesifikasi setiap aplikasi, call tree, tanggung jawab paket, onboarding agent, tutorial, serta pemeliharaan dokumentasi. Semua perubahan tracked berupa Markdown. Tidak ada perubahan behavior, capability SQL, parser/fixer, report schema, source SQL, fixture golden, migration, release script, atau binary.

## Hasil review

| Gap awal | Perbaikan |
| --- | --- |
| Rules dan desain tersedia tetapi belum ada rute belajar/source map yang ringkas | Guide menautkan overview, reference, call tree, package catalog, tutorial, agent workflow. |
| Guide masih menyebut empat executable, plan menyebut tiga | Diselaraskan menjadi lima aplikasi utama dan dua utilitas live test. |
| Deskripsi paket seolah seluruh reconstruction/diff/render sudah pindah | FSD dan katalog menjelaskan ownership aktual `history/diff/emit` vs `engine`. |
| Linux compatibility build dibaca sebagai checker-only build | README menjelaskan delegation ke complete release dan gate identity/source digest. |
| Contoh operasional belum berupa latihan end-to-end terisolasi | Tutorial baru dijalankan pada folder `out/docs_lab_<id>`, dengan checker no-write dan second-gen no-op assertions. |
| Panduan penulisan/maintenance belum menjelaskan cara mencegah stale docs | Maintenance guide memakai Diátaxis, C4 zoom, docs-as-code, source links, validation scope, dan update map. |

## Validation yang dijalankan

| Pemeriksaan | Hasil dan scope |
| --- | --- |
| `go test ./...` | Exit `0`, seluruh package test suite lulus. Live SQL Server integration tetap mengikuti opt-in; hasil ini tidak mengklaim eksekusi DB. |
| `powershell -NoProfile -ExecutionPolicy Bypass -File scripts/validate_qa_fixture_contract.ps1` | Exit `0`, fixture contract validation passed. Fixture/golden tidak diubah. |
| Seluruh enam PowerShell blocks tutorial, diekstrak berurutan ke script lokal | Exit `0`: bootstrap core + table pair, checker plan dengan hash source/history identik, verify, checker apply setelah edit, draft/seal, second gen `GEN_NO_CHANGES`, report status OK. |
| WSL Ubuntu 24.04 `sha256sum -c tools/linux-amd64/release.sha256` | Seluruh complete release entries OK, termasuk Windows/Linux binaries, individual checksum files, license, build identity, release manifest. |
| Linux `rza_migrate_check -build-info` dibanding `tools/linux-amd64/build-info.txt` | Identik, exit `0`. |
| WSL `sh scripts/verify_release_source.sh .` | Exit `0`; source digest `c4b48b2d5e2161a949e6617f76ae1a6d0f914c3675e070285fa5ff87e6a425d0`. |
| Linux checker command GitLab, di bawah | Exit `0`, `fail=0`, `pass=4`, `migrations_generated=0`. Ini smoke pada checkout project, bukan simulasi seluruh central pipeline. |
| Hash sebelum/sesudah Linux smoke pada `db_object`, `patch`, `migrations`, dan `tools` | Seluruh 22 file yang tersedia identik; authored source/history dan release bytes tidak berubah. |
| Pemeriksaan lokal Markdown/source | 14 Markdown files, 207 local links, dan 52 simbol call tree lolos; relative filename casing, heading fragments yang ada, dan fence closure diperiksa. |

Command Linux checker yang dieksekusi dari root repository:

```sh
tools/linux-amd64/rza_migrate_check -repo . -check -profile authoring -migration-validation integrity -pause-on-exit false -perf-notes true
```

Source path, link/fragment baru, fenced code, dan nama simbol call tree diperiksa lokal terhadap checkout; `git diff --check` dipakai sebelum commit. Tutorial juga mengassert history tetap setelah apply/draft, tepat satu draft pair, dan sealed bytes identik dengan draft yang direview. Bukti run sementara berada di `out/documentation_review/` dan folder latihan yang diabaikan Git, bukan artefak yang dipublikasikan.

## Dokumen updated

- `README.md`, `docs/engine_docs_guide.md`: jalur masuk dokumentasi dan keterangan release wrapper.
- `spec.md`, `fsd.md`, `plan.md`, `docs/walkthrough.md`: navigasi dan penjelasan implementasi/lingkup dokumentasi; tidak mengubah aturan normatif.
- `docs/architecture/overview.md`, `call_tree.md`, `package_catalog.md`: tujuan, flow, ownership, source links, batas statis/runtime.
- `docs/engine_reference.md`: spesifikasi operasional lima aplikasi dan dua utilitas, mode/flag/config, input/output, side effects, diagnostic.
- `docs/engineering/agent_onboarding.md`: urutan baca, impact map, validation dan handover.
- `docs/tutorials/offline_authoring.md`: contoh runnable terisolasi.
- `docs/documentation_maintenance.md`: praktik, sumber eksternal, update/validation map.
- Dokumen review ini: baseline, scope, hasil, dan batas validation.

## Reviewed dan intentionally unchanged

| Dokumen | Alasan tidak diubah |
| --- | --- |
| `rules.md` | Tidak ada perubahan safety rule, hard lock, atau enum. |
| `docs/desired_state_lifecycle.md` | Desired/history/removal/reintroduction dan draft/seal tetap. |
| `docs/engine_guarded_migrations.md` | Tidak mengubah SQL yang dihasilkan atau up/down policy. Capability rinci tetap dirujuk ke matrix. |
| `docs/qa/sql_server_capability_matrix.md` | Tidak menambah fitur/object kind atau mengubah supported/manual/fail classification. |
| `docs/traceability/rules_traceability_matrix.md` | Tidak menambah/mengubah rule, report code, control, atau QA coverage mapping. |
| `docs/qa/qa_gate_engine_rules.md`, `fixtures/qa_rules/README.md` | Acceptance engine dan fixture shape tetap; tutorial dokumentasi bukan QA rule baru. |
| `docs/qa/qa_coverage_report.md` | Generated facts tidak berubah; tidak hand-edit/regenerate untuk memberi kesan coverage baru. |
| `docs/qa/learning_edge_case_backlog.md` | Tidak ada edge case engine baru yang ditemukan/diselesaikan dalam scope dokumentasi ini. |
| `docs/bootstrap_four_repos.md`, `constraints_naming.md` | Bootstrap/extraction/catalog policy dan optional naming guideline tetap. |
| `docs/policy/production_readiness.md` | Release contract tetap; hanya memperjelas deskripsinya pada README/call tree. |
| `.gitlab/rza_migrate_check.gitlab-ci.yml` dan MR templates | Job flags, gate dan checklist tidak berubah. Pipeline YAML tidak diedit. |
| `docs/policy/rest_api_hardening.md`, `docs/operations/migration_monitoring.md` | Tidak mengubah execution/recovery/API/observability behavior. |
| `docs/policy/parallel_safe_design.md`, `rebaseline_playbook.md`, `legacy_wrapper_cutover.md` | Tidak mengubah concurrency/cache, history repair, atau audit cutover. |
| `docs/report_output_layout.md`, `docs/database/*` | Tidak mengubah report publisher/layout/inference; output database project tidak diregenerate hanya untuk menjelaskan software. |

## Batas validasi dan release

Tidak menjalankan extraction/live migration/reset atau menghubungi database. Full hardening/performance suite tidak diulang karena tidak ada perubahan engine/rule/concurrency; full Go suite, fixture contract, tutorial, dan released Linux smoke dipakai untuk scope ini. Go tests dapat skip live tests bila opt-in belum tersedia.

`build_all.bat` tidak dijalankan: hanya Markdown di luar input source digest yang berubah. Source digest release terbukti masih cocok dan release bytes tetap. Ini tidak memberikan pengecualian untuk task source berikutnya; setelah digest input berubah, complete rebuild serta Linux gate wajib diulang.

Skill `mssql-object-engineering` tidak diubah: ownership dan batas yang diverifikasi sudah sesuai reference yang tersedia; tidak ada pelajaran parser/SQL baru yang memerlukan update skill. Tidak ada klaim bahwa seluruh grammar, data rollback, atau production runtime telah divalidasi.

Commit/push dilakukan hanya untuk permintaan dokumentasi yang telah direview ini. Status CI GitHub/GitLab remote harus diperiksa terpisah dari keberhasilan push lokal.
