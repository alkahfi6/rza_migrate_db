# Rebaseline Playbook (Draft)

Dokumen ini menyimpan rencana **re-baseline migration history** untuk repo `1 initial catalog = 1 repo`.

Status:
- `DRAFT`
- Belum dieksekusi.
- Fokus dokumen ini: menurunkan risiko saat jumlah object/migration sudah sangat besar.

## 1. Tujuan

1. Menetapkan ulang baseline history migration secara terkontrol.
2. Menjaga aturan `rules.md` tetap terpenuhi (deterministic, audited, fail-hard).
3. Mencegah praktik berisiko seperti hapus file migration acak.

## 2. Kapan Rebaseline Dipertimbangkan

Rebaseline dipertimbangkan bila salah satu kondisi berikut konsisten terjadi:
- `migrations/` tumbuh sangat besar dan memperlambat reconstruct/verify/gen.
- Waktu run harian meningkat signifikan dan mengganggu workflow tim.
- Proses onboarding/review migration history sudah tidak efisien.

Catatan:
- Rebaseline **bukan** aktivitas rutin.
- Rebaseline hanya dilakukan lewat change window terjadwal.

## 3. Prinsip Wajib

1. Satu repo tetap untuk satu initial catalog.
2. Tidak boleh menghapus migration lama secara acak.
3. Harus ada jejak audit penuh: tag, arsip, dan checklist persetujuan.
4. Seluruh environment harus diselaraskan ke titik yang sama sebelum cutover.
5. Setiap langkah wajib punya gate verifikasi otomatis.

## 4. Risiko Utama

- Salah sinkron version state antar environment.
- Kehilangan traceability bila arsip/tag tidak disiplin.
- Drift schema bila cutover tanpa validasi state final.
- Runner mismatch (`golang-migrate` version state vs baseline baru).

## 5. Prasyarat Sebelum Eksekusi

1. Freeze perubahan schema/data migration selama window.
2. Semua environment target berada pada version migration yang sama dan tidak dirty.
3. Backup database dan backup repository tersedia.
4. Approval dari owner aplikasi + DBA + release manager.
5. Rencana rollback tertulis dan disetujui.

## 6. Fase Eksekusi (Rencana)

## Fase 0 - Decision Gate

Output:
- Justifikasi rebaseline.
- Scope environment yang terdampak.
- Jadwal change window.

Gate:
- Persetujuan lintas fungsi (engineering, DBA, release).

## Fase 1 - Snapshot & Freeze

Langkah:
1. Freeze perubahan migration baru.
2. Jalankan extractor/validasi untuk snapshot latest desired state.
3. Pastikan source-of-truth bersih dan konsisten.

Gate minimal:
- `rza_get_object_db -dry-run`
- `rza_migrate_db verify`
- `go test ./...`

## Fase 2 - Arsip History Lama

Langkah:
1. Buat tag git sebelum rebaseline (contoh: `pre-rebaseline-YYYYMMDD`).
2. Simpan arsip migration lama sebagai artifact release.
3. Dokumentasikan mapping commit/tag untuk audit.

Gate:
- Arsip dan tag tervalidasi bisa direstore.

## Fase 3 - Build Baseline Baru

Strategi yang direkomendasikan:
- Bangun baseline baru secara terkontrol (branch/repo cutover), bukan edit acak.
- Tetap mematuhi kontrak `db_object/`, `patch/`, `migrations/` dan core audit pair.

Catatan:
- Pilihan strategi final (branch cutover vs repo baru) diputuskan saat eksekusi.
- Untuk environment existing, sinkronisasi state `golang-migrate` harus direncanakan eksplisit.

Gate:
- `rza_migrate_db verify` harus pass.
- `rza_migrate_db gen` pada baseline baru tidak memunculkan drift tak terduga.

## Fase 4 - Uji Cutover

Langkah:
1. Uji pada environment staging clone.
2. Validasi runner behavior end-to-end.
3. Validasi output artifact dan audit log contract.

Gate minimal:
- `go test ./...`
- `powershell -ExecutionPolicy Bypass -File scripts/run_final_hardening_gate.ps1`
- Verifikasi migration runner lock contract.

## Fase 5 - Production Cutover

Langkah:
1. Jalankan cutover sesuai window.
2. Pantau error dan stop condition.
3. Simpan evidence execution.

Gate:
- Post-cutover verify pass.
- Tidak ada dirty migration state.

## 7. Rollback Strategy (High Level)

Jika cutover gagal:
1. Stop perubahan lanjutan.
2. Kembali ke tag/artefak `pre-rebaseline`.
3. Restore pipeline dan chain migration lama.
4. Lakukan post-mortem sebelum window ulang.

## 8. Evidence Wajib

- Approval record.
- Snapshot version seluruh environment.
- Tag dan arsip history lama.
- Hasil command gate (`verify`, `gen`, test suite, hardening gate).
- Catatan hasil cutover dan rollback (jika terjadi).

## 9. Larangan Saat Rebaseline

- Hapus migration lama tanpa tag/arsip.
- Ubah banyak hal sekaligus di luar scope rebaseline.
- Cutover tanpa freeze.
- Cutover tanpa verifikasi environment version parity.

## 10. Backlog Otomasi (Untuk Dikerjakan Nanti)

1. Script precheck parity version antar environment.
2. Script manifest/hash snapshot schema sebelum dan sesudah baseline.
3. Script checklist runner cutover otomatis.
4. Template laporan cutover + evidence collector.

## 11. Status Implementasi Dokumen Ini

- Dokumen ini menyimpan rencana.
- Eksekusi teknis detail dan otomasi akan dikerjakan di fase berikutnya.
