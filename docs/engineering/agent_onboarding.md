# Panduan Kerja Agent dan Developer

Tujuan panduan ini adalah membuat agent berikutnya dapat menentukan entrypoint, kontrak yang terdampak, dan bukti yang diperlukan sebelum mengubah engine. Ini panduan kerja; [AGENTS.md](../../AGENTS.md) dan [rules.md](../../rules.md) tetap wajib dibaca.

## Rute baca awal

| Kebutuhan | Baca berurutan | Hasil yang harus dipahami |
| --- | --- | --- |
| Baru mengenal project | [README](../../README.md), [guide](../engine_docs_guide.md), [overview](../architecture/overview.md), [tutorial offline](../tutorials/offline_authoring.md) | Lima aplikasi utama, dua utilitas live test, dan batas desired/history/runtime. |
| Mulai coding | Rules terkait, [spec](../../spec.md), [FSD](../../fsd.md), [referensi engine](../engine_reference.md), [call tree](../architecture/call_tree.md) | Requirement, flow aktif dan side effects, bukan hanya nama paket. |
| Memilih lokasi perubahan | [Katalog paket](../architecture/package_catalog.md), tabel impact di bawah, tests di sekitar source | Satu shared control yang dapat diuji lintas consumer. |
| Menutup task | [QA gate](../qa/qa_gate_engine_rules.md), [traceability](../traceability/rules_traceability_matrix.md), [maintenance](../documentation_maintenance.md) | Evidence, release parity bila diperlukan, update dokumen yang tepat. |

## Orientasi workspace sebelum mengubah file

Jalankan dari root checkout:

```powershell
git status --short
git log -5 --oneline
git branch -vv
rg --files -g AGENTS.md
rg -n '^func ' internal/engine/engine.go internal/engine/gen_emit_commit.go
go test ./...
```

Catat perubahan yang sudah ada dan jangan masukkan ke commit task lain. Jika meneruskan task terputus, verifikasi file/evidence di workspace; narasi selesai pada chat lama bukan bukti bahwa validation/build/push berhasil. Go test tanpa live DSN tidak membuktikan SQL sudah dieksekusi di server. Jangan membuka credential lokal untuk onboarding offline.

## Peta impact berdasarkan jenis perubahan

| Jenis task | Implementasi yang diperiksa | Consumer dan bukti minimum |
| --- | --- | --- |
| Encoding, canonical SQL, safe correction | `sqlfmt`, `sqlfix`, `tsqlast`, `getobject`, `engine.runFmt` | Extractor canonical output; checker apply/plan; generator genuine delta; idempotence dan ambiguity negative. |
| Object identity/casing/path | `fsx`, `sqlparse`, `history`, `verify`, `engine` | Filename vs definition, escaped identifier, case-sensitive ambiguity, Windows/Linux actual-entry parity. |
| DDL/diff/capability baru | Typed model, risk, reconstruction, dependency order, renderer | Up/down round-trip, supported/manual/fail classification, source removal/reintroduction, fixture/golden. |
| Report relationships/procedure impact | `sqlparse/access`, `migreport` analyzer/writer/layout | Direct vs transitive, FK vs observed JOIN, unresolved dynamic, source/history no-write, rerun output stable. Audit shared parser consumer meski fitur report-only. |
| Checker profile/report code | `internal/cmd`, `engine`, `verify`, `output` | Exit code, apply vs `-check`, complete console diagnostic, migrations tetap identik, exact MR flags. |
| REST/execution | `internal/cmd/rest`, `handler`, `migrest`, telemetry | Mock unit + authorized live DB bila diperlukan; dirty/version guard, lock release, history evidence, journal, HTTP contract. |
| Parallel/cache/performance | Snapshot/cache, collect tasks, verify orchestrator | Input invalidation, stable merge, serial/parallel repeatability, benchmark, no hidden shared writes. |
| Build/release/CI | `build_all.bat`, release scripts, `.gitlab`, buildinfo/license | Complete publication, checksum, embedded identity/source digest, Linux checker smoke. |

## Checklist perubahan `db_object`

1. Tulis contoh before/after beserta semantic delta yang harus tetap terlihat.
2. Telusuri extractor, checker apply, checker `-check`, generator `fmt/verify/gen`, report bila relevan, dan released Linux consumer.
3. Letakkan parsing/fixer bersama di `internal/`; adapter CLI hanya menentukan mode.
4. Uji kasus positif, no-op, ambiguous/semantic negative, casing, second-run idempotence. Hash authored SQL dan seluruh `migrations/` sebelum/sesudah checker plan maupun apply.
5. Buktikan generator hanya menghasilkan delta nyata dan up/down yang valid; jangan menyamakan formatting equivalence dengan business equivalence.
6. Update minimum documentation set dalam AGENTS beserta fixture/report code mapping.
7. Jalankan full tests dan gate relevan. Setelah source final, build seluruh release, verifikasi source digest/identity/checksum dan command Linux GitLab yang sama.

Tidak semua komponen harus melakukan aksi sama. Report-only perubahan boleh tidak menambah fixer; extractor tidak membuat migration; checker tidak mempunyai generation path; CI tidak memperbaiki authored SQL. Tulis alasan non-applicable secara eksplisit dan buktikan boundary dengan test ketika behavior berubah.

## Validation sesuai scope

| Scope | Validation |
| --- | --- |
| Dokumentasi saja | Link/path/symbol check, contoh command pada sandbox, `git diff --check`, tests terkait kontrak dokumentasi dan `go test ./...`; catat source/release tidak berubah. |
| Parser/fixer/generator | Focused tests, fixtures/goldens, `go test ./...`, full hardening, complete build + release/source verification + Linux smoke. |
| Performance/concurrency | Tambah repeatability/benchmark gate sesuai [parallel-safe policy](../policy/parallel_safe_design.md). |
| Live SQL behavior | Hanya target test yang diotorisasi; `RZA_MSSQL_TEST_DSN` dan optional corpus flags mengikuti test source. Catat skipped/inaccessible, jangan klaim pass live bila tidak dijalankan. |

Full hardening: `powershell -ExecutionPolicy Bypass -File scripts/run_final_hardening_gate.ps1`. Script dapat menulis generated coverage dan learning hygiene; review diff setelahnya. Dokumentasi saja di luar source digest tidak memerlukan penerbitan ulang lisensi/binary. Jika input digest berubah, complete build kembali wajib; jangan memperbaiki checksum secara terpisah.

## Triage kegagalan tanpa mengubah bukti

| Gejala | Langkah pertama |
| --- | --- |
| `FMT_CHECK_CHANGES_REQUIRED` atau pending safe fix | Jalankan checker apply pada source yang sedang dikerjakan, review diff, lalu ulangi `-check`. |
| Desired/history drift | Tentukan apakah desired edit memang change yang harus digenerate; baca lifecycle sebelum menghapus/menambah ulang object. |
| `MANUAL_MIGRATION_REQUIRED` / unsupported | Baca capability dan suggestion; jangan menghapus guard agar generation lewat. |
| Release digest/checksum mismatch | Periksa source/release mana yang stale; build lengkap setelah source final. |
| REST dirty/history hash mismatch | Ikuti monitoring/recovery runbook; jangan edit applied migration, checksum, atau force version tanpa workflow recovery eksplisit. |
| Report unresolved/ambiguous | Pertahankan evidence dan batas resolusi; jangan jadikan tebakan sebagai FK/uniqueness/schema action. |

## Handover yang dapat dilanjutkan

Gunakan template berikut pada deskripsi task/PR. Isi fakta aktual, termasuk hal yang belum dijalankan.

```text
Tujuan dan acceptance:
Source baseline / branch:
Perubahan dan entrypoint:
Kontrak lintas tool (extract/apply/plan/gen/report/CI):
Dokumen updated / reviewed unchanged + alasan:
Validation (command, exit, scope, skipped):
Release (dibangun atau N/A + alasan, source digest bila relevan):
Risiko/batas dan pekerjaan tersisa:
Commit dan status push:
```

Jangan menyimpan DSN, API key, local `out/`, atau real credential dalam handover/commit. Baca [publication scope](../github_repository.md). Struktur commit harus menjelaskan perubahan final dan bukti, lalu periksa remote status setelah push; push sukses belum berarti CI remote sudah lulus.
