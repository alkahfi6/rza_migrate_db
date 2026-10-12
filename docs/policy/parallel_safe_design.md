# Parallel-Safe Design for `rza_migrate_db` (Draft)

Dokumen ini mendesain strategi paralelisasi `rza_migrate_db` untuk repo besar (ribuan object + migration history panjang) **tanpa merusak determinism dan contract rules**.

Status:
- `DRAFT`
- Fokus: desain + rollout plan.
- Implementasi bertahap, default tetap aman.
- P0 telemetry baseline sudah tersedia (opt-in) via `rza_migrate_db -perf-notes true` atau `RZA_PERF_NOTES=true`.
- P1 collect-source parallel-safe sudah diimplementasikan (opt-in) via `-parallel-read true` + `-parallel-workers <n>` atau `RZA_PARALLEL_READ` + `RZA_PARALLEL_WORKERS`.
- P2 history preload parallel-safe sudah diimplementasikan (opt-in) pada `reconstructHistory`: preload filename parse + file read paralel, reduce state tetap serial.
- P3 forbidden-token scan parallel-safe sudah diimplementasikan (opt-in) dengan deterministic merge hasil scan per file.
- P4 production hardening gate sudah diimplementasikan:
  - `scripts/validate_parallel_repeatability.ps1` (determinism gate, blocking),
  - `scripts/benchmark_parallel_read.ps1` (baseline benchmark artifact).
- P7-lite sudah diimplementasikan: callback verify independen (`core_audit_contract` + `forbidden_sql`) dieksekusi paralel dengan sub-report terisolasi dan merge deterministic.
- P8 sudah diimplementasikan: effective parallel workers di-cap otomatis oleh `GOMAXPROCS` dan jumlah task.
- P9-lite sudah diimplementasikan: `fail-fast` opsional (`-fail-fast` / `RZA_FAIL_FAST`) untuk mode lokal cepat, default tetap full-report.

## 1. Objective

1. Mempercepat path `verify`/`gen` pada repository besar.
2. Menjaga output tetap deterministic (hasil identik untuk input identik).
3. Menjaga fail-hard behavior dan contract `rules.md`.

## 2. Constraint yang Tidak Boleh Dilanggar

1. Urutan emit migration tetap deterministic.
2. Pairing/version/check rule tidak berubah.
3. `out/output.json` contract tidak berubah.
4. Severity `3` tetap fail-hard.
5. Tidak ada race condition pada state/report.

## 3. Kondisi Saat Ini (As-Is)

`rza_migrate_db` saat ini mayoritas serial:
- `fmt`, `verify`, `gen` orchestration serial.
- reconstruction history serial.
- diff/planning serial.

`rza_get_object_db` sudah memakai worker pool pada beberapa tahap export, tetapi desain ini khusus `rza_migrate_db`.

## 4. Prinsip Parallel-Safe

Prinsip utama: **parallelize read/parse independent tasks, keep ordered reduce sequential**.

Artinya:
- Tahap yang independent per-file boleh diparalelkan.
- Tahap yang stateful (history reduce, planner mutate state, emit) tetap serial.
- Semua hasil worker harus di-merge menurut urutan deterministik (index/path sort), bukan urutan selesai worker.

## 5. Kandidat Paralelisasi

## 5.1 Aman Diparalelkan (Phase awal)

1. Read+normalize file source di `collectDesiredObjects`.
2. Read file patch (`README/up/down`) di `collectPatchSources`.
3. Forbidden token scan per file (jika dibutuhkan, dengan merge hasil terurut).

## 5.2 Hybrid (Parallel + Serial Reduce)

`reconstructHistory`:
1. Parallel stage:
   - baca file migration,
   - normalize text,
   - parse filename metadata.
2. Serial reduce stage (tetap urut filename):
   - apply semantic parser berbasis state object,
   - apply side effect drop constraint/index,
   - update `historyState`.

## 5.3 Tetap Serial (Jangan Diparalelkan dulu)

1. Diff planner yang memodifikasi `historyState` incremental.
2. Version allocator + emit file.
3. Update `output.Report` (`AddPass`/`AddFail`) langsung dari worker.

## 6. Desain Teknis Inti

## 6.1 Worker Pool Bounded

Gunakan worker pool bounded:
- input: daftar task dengan `task_index` deterministik.
- output: `result{task_index, data, err, diagnostics}`.

Workers tidak menulis shared map/report langsung.

## 6.2 Deterministic Merge

Setelah worker selesai:
1. kumpulkan seluruh result,
2. sort by `task_index`,
3. baru apply ke state/report pada main goroutine.

Ini menjaga deterministic ordering meski completion order acak.

## 6.3 Report Aggregation

Larangan:
- `report.AddPass`/`report.AddFail` dipanggil dari goroutine worker.

Aturan:
- worker hanya return diagnostics,
- main goroutine yang append ke report sesuai urutan deterministic.

## 6.4 Error Handling

1. Worker error dikumpulkan dan di-merge deterministic.
2. Jika ada fatal error -> fail-hard sesuai contract.
3. Tidak ada panic silent; semua error harus terlapor sebagai fail item / internal error.

## 7. Guard Configuration (Safe by Default)

Agar tidak merusak behavior existing:

1. Default mode: serial (current behavior).
2. Parallel mode: opt-in via config/flag/env.
3. Fallback:
   - jika parallel branch error tak terduga, proses fail-hard (tidak auto-skip ke mode lain tanpa jejak).

Rekomendasi naming (saat implementasi):
- `EnableParallelRead` (default `false`)
- `ParallelWorkers` (default `4`, validasi `>=1`)

## 8. Rollout Bertahap

## P0 - Measurement Baseline

Output:
- metrik waktu tiap fase (`fmt`, `verify`, `reconstruct`, `collectDesired`, `plan`, `emit`).

Tujuan:
- tahu bottleneck riil sebelum mengubah arsitektur.

## P1 - Parallel Collect Source

Scope:
- `collectDesiredObjects` dan `collectPatchSources`.

Gate:
- output dan diff hasil identik dengan serial.

## P2 - Parallel Preload History

Scope:
- preloading file+metadata history parallel,
- reduce tetap serial.

Gate:
- reconstruction output identik dengan serial.

## P3 - Optional Parallel Forbidden Scan

Scope:
- scan token forbidden per file parallel + deterministic merge.

Gate:
- daftar fail/pass identik urutan deterministik.

## 9. Test Strategy Wajib

1. `go test ./...`
2. `go test -race ./...` pada pipeline yang mendukung race detector.
3. Repeatability test:
   - jalankan command yang sama N kali (mis. 20x),
   - hash `out/output.json` dan generated files harus identik.
4. Existing QA gate tetap wajib pass:
   - `scripts/run_final_hardening_gate.ps1`

## 10. Risiko & Mitigasi

1. Risiko nondeterministic output.
   - Mitigasi: deterministic merge by index/path + stable sort.

2. Risiko race di shared state.
   - Mitigasi: no shared mutable state dari worker; reducer tunggal.

3. Risiko debugging makin sulit.
   - Mitigasi: phase metrics + trace note per stage.

4. Risiko regresi behavior.
   - Mitigasi: default tetap serial, parallel opt-in bertahap.

## 11. Definition of Done (Untuk Implementasi Nanti)

Parallel-safe dianggap siap produksi jika:
1. Semua gate test existing pass.
2. Race test pass.
3. Determinism repeatability pass.
4. Performa meningkat signifikan pada dataset besar (target minimal disepakati saat P0).
5. Dokumentasi (`walkthrough`, `fsd`, `plan`, `qa gate`) sudah sinkron.
