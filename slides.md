---
theme: seriph
background: false
title: Modern Database Migration & Deployment Architecture
info: |
  Panduan Arsitektur & Operasional Database RZA Suite
class: text-left
drawings:
  persist: false
transition: slide-left
mdc: true
---

# Modern Database Migration & Deployment Architecture

Meninggalkan Tradisi Manual Menuju Alur Rilis Database Deklaratif, Aman dari Tabrakan Git, dan Ter-Audit Otomatis

<div class="pt-8">
  <span class="text-sm font-semibold opacity-60">RZA Migration Suite · Enterprise Architecture Briefing</span>
</div>

---

# Masalah #1: Tiga Bom Waktu Eksekusi Manual via SSMS

Kebiasaan membuka server database live dan melakukan copy-paste query adalah sumber utama insiden sistem.

<div class="grid grid-cols-2 gap-8 pt-4">

<div class="border border-red-500/30 bg-red-500/5 p-4 rounded-lg">
  <h3 class="text-red-500 font-bold mb-2">Tragedi Patch Row 4.500</h3>
  <p class="text-sm opacity-80 leading-relaxed">
    Script patching update 10.000 data nasabah putus di tengah jalan karena timeout jaringan.
  </p>
  <p class="text-sm opacity-80 leading-relaxed mt-2">
    Tanpa transaksi atomik (<code>BEGIN TRAN</code>), database berstatus <strong>setengah matang</strong>—separuh data berubah, separuh masih data lama.
  </p>
  <div class="text-xs text-red-400 mt-4 border-t border-red-500/20 pt-2 font-mono">
    Dilema: Re-run takut duplikat, rollback tidak ada!
  </div>
</div>

<div class="space-y-4">
  <div class="border border-slate-700 p-3 rounded">
    <div class="font-bold text-sm">1. Backup SP ke Notepad Pribadi</div>
    <div class="text-xs opacity-75 mt-1">Script lama di-copy ke laptop pribadi. Saat ada bug di Production, versi aslinya hilang atau tercecer.</div>
  </div>
  <div class="border border-slate-700 p-3 rounded">
    <div class="font-bold text-sm">2. Human Error Urutan DDL</div>
    <div class="text-xs opacity-75 mt-1">Foreign Key dieksekusi sebelum tabelnya ada, atau script index terlewat tanpa disadari.</div>
  </div>
  <div class="border border-slate-700 p-3 rounded">
    <div class="font-bold text-sm">3. Salah Kamar Tab SSMS</div>
    <div class="text-xs opacity-75 mt-1">Buka tab Dev dan Prod bersamaan. Niat tes di Dev, ternyata kursor aktif mengeksekusi di tab Production.</div>
  </div>
</div>

</div>

---

# Pertanyaan Besar: "Kenapa Nggak Pakai Golang-Migrate Biasa Aja?"

Bukankah lebih sederhana jika setiap developer cukup membuat file `.up.sql` langsung di folder `migrations/`?

<div class="grid grid-cols-2 gap-8 pt-4">

<div class="border border-blue-500/30 bg-blue-500/5 p-4 rounded-lg">
  <h3 class="text-blue-400 font-bold mb-2">Ekspektasi Tim (Pola Sederhana)</h3>
  <ul class="text-sm space-y-2 opacity-80">
    <li>1. Developer membuat file di branch fiturnya:<br><code>20261101_add_col.up.sql</code></li>
    <li>2. Di-commit bersama kode program ke Git.</li>
    <li>3. Saat deploy, runner jalankan: <code>migrate up</code>.</li>
  </ul>
  <p class="text-xs opacity-60 mt-4 italic border-t border-blue-500/20 pt-2">
    Pola ini berhasil HANYA jika developernya cuma 1 orang dan jadwal deploy strictly satu per satu.
  </p>
</div>

<div class="border border-red-500/30 bg-red-500/5 p-4 rounded-lg">
  <h3 class="text-red-400 font-bold mb-2">Kenyataan Fatal Multi-Branch</h3>
  <ul class="text-sm space-y-2 opacity-80">
    <li><strong>1. Jebakan Silent Skip:</strong> Jika branch lain rilis duluan, versi database target naik. File Anda yang bertanggal lebih lama DILEWATI tanpa pesan error!</li>
    <li><strong>2. Database Dirty State:</strong> Jika urutan dipaksa, database terkunci status DIRTY dan menolak migrasi berikutnya.</li>
    <li><strong>3. Kucing-kucingan Timestamp:</strong> Developer lelah me-rename tanggal file migrasi tiap kali ada fitur lain yang nyalip deploy.</li>
  </ul>
</div>

</div>

---

# Simulasi Kasus: Kronologi Kerusakan Akibat "Silent Skip"

Bagaimana fitur yang rilis duluan menggeser version schema dan memicu kegagalan sistem.

<div class="grid grid-cols-4 gap-4 pt-6 text-xs">

<div class="border border-slate-700 p-3 rounded bg-slate-900/40">
  <div class="font-bold text-blue-400 text-sm mb-1">01. Kondisi Awal</div>
  <div class="opacity-60 mb-2">1 November</div>
  <p>DB di versi: <code>v20261001</code></p>
  <p class="mt-2">• Dev A (Pajak):<br><code>20261101_pajak.sql</code></p>
  <p class="mt-1">• Dev B (Kupon):<br><code>20261102_kupon.sql</code></p>
</div>

<div class="border border-emerald-700 p-3 rounded bg-emerald-950/20">
  <div class="font-bold text-emerald-400 text-sm mb-1">02. Kupon Nyalip</div>
  <div class="opacity-60 mb-2">3 November</div>
  <p>Fitur Kupon selesai testing duluan & dideploy ke Production.</p>
  <p class="mt-2 text-emerald-300">Versi DB naik menjadi: <strong>v20261102</strong>.</p>
  <p class="mt-1">Deploy Fitur B sukses.</p>
</div>

<div class="border border-amber-700 p-3 rounded bg-amber-950/20">
  <div class="font-bold text-amber-400 text-sm mb-1">03. Pajak Deploy</div>
  <div class="opacity-60 mb-2">5 November</div>
  <p>File Pajak bertanggal <code>20261101</code> (lebih tua dari versi DB <code>20261102</code>).</p>
  <p class="mt-2 text-amber-300 font-bold">Aturan golang-migrate: Versi lama DILEWATI (Silent Skip)!</p>
</div>

<div class="border border-red-700 p-3 rounded bg-red-950/40">
  <div class="font-bold text-red-400 text-sm mb-1">04. Incident Prod</div>
  <div class="opacity-60 mb-2">Produksi Crash</div>
  <p>Aplikasi live, tapi kolom <code>TaxAmount</code> TIDAK PERNAH DIBUAT di database.</p>
  <p class="mt-2 text-red-300 font-mono text-[11px]">Invalid column name TaxAmount</p>
  <p class="mt-1 font-bold">Nasabah gagal bayar.</p>
</div>

</div>

---

# Masalah #2: Fitur Batal Rilis & Repo Buta (Tanpa Desired State)

<div class="grid grid-cols-2 gap-8 pt-4">

<div>
  <h3 class="text-amber-400 font-bold mb-2">Dilema Fitur Batal Rilis di H-1</h3>
  <p class="text-sm opacity-80 leading-relaxed mb-3">
    Di plain golang-migrate, file migrasi sudah di-commit di branch rilis. Tiba-tiba bisnis menunda rilis fitur tersebut.
  </p>
  <ul class="text-xs space-y-2 opacity-80">
    <li>• <strong>Jika file dihapus:</strong> Urutan timestamp versi jadi bolong/loncat.</li>
    <li>• <strong>Jika dibiarkan:</strong> Tabel terlanjur ter-alter padahal aplikasinya belum rilis.</li>
    <li>• <strong>Jika dibuatkan script revert:</strong> Histori migrasi jadi kotor berantakan.</li>
  </ul>
</div>

<div>
  <h3 class="text-blue-400 font-bold mb-2">Ketiadaan Desired State (Repo Buta)</h3>
  <p class="text-sm opacity-80 leading-relaxed mb-3">
    Di plain golang-migrate, tidak ada satu pun file di Git yang mencerminkan bentuk tabel saat ini.
  </p>
  <ul class="text-xs space-y-2 opacity-80">
    <li>• Yang ada di repo cuma tumpukan 200 file migrasi lama bertimpaan: <code>add_col</code>, <code>drop_col</code>, <code>alter_col</code>.</li>
    <li>• Developer yang mau tahu struktur tabel hari ini terpaksa buka SSMS langsung ke live server atau menghitung 200 file di kepala.</li>
    <li>• Git kehilangan fungsinya sebagai <em>Single Source of Truth</em> bentuk skema database.</li>
  </ul>
</div>

</div>

---

# Solusi rza_migrate_db: Pemisahan Desired State vs History

Developer cukup mendeklarasikan bentuk akhir objek. Engine yang mengurus hitungan transisi perubahannya.

<div class="grid grid-cols-2 gap-8 pt-4">

<div class="border border-blue-500/30 p-4 rounded-lg bg-blue-500/5">
  <span class="text-xs font-bold text-blue-400 uppercase tracking-wider">Tugas Developer di Laptop</span>
  <h3 class="text-lg font-bold mt-1 mb-3">Folder db_object/ (Desired State)</h3>
  <ul class="text-xs space-y-2.5 opacity-85">
    <li>• <strong>Deklaratif:</strong> Cukup nyatakan tabel dbo.Transaksi sekarang punya kolom TaxAmount.</li>
    <li>• <strong>1 Objek = 1 File:</strong> <code>db_object/tables/dbo.Transaksi.sql</code></li>
    <li>• <strong>Jelas di Merge Request:</strong> Reviewer melihat bentuk utuh tabel, bukan potongan ALTER acak.</li>
    <li>• <strong>Zero Migration:</strong> Dilarang membuat file di folder migrations/.</li>
  </ul>
</div>

<div class="border border-slate-700 p-4 rounded-lg bg-slate-800/40">
  <span class="text-xs font-bold text-slate-300 uppercase tracking-wider">Tugas Engine Saat Rilis</span>
  <h3 class="text-lg font-bold mt-1 mb-3">Folder migrations/ (History State)</h3>
  <ul class="text-xs space-y-2.5 opacity-85">
    <li>• <strong>Dihasilkan Otomatis oleh rza_migrate_db:</strong> Engine membandingkan desired state vs histori terakhir dan menghitung delta murni.</li>
    <li>• <strong>Timestamp Selalu Fresh:</strong> File migrasi baru dibuat saat mau rilis. Dijamin selalu lebih baru dari database target!</li>
    <li>• <strong>Selalu Berpasangan:</strong> Menyertakan .up.sql dan .down.sql rollback valid.</li>
    <li>• <strong>Transaksi Atomik:</strong> Dibungkus TRY/CATCH/THROW otomatis.</li>
  </ul>
</div>

</div>

---

# Mekanisme Git Branching: Solusi Mutlak Tabrakan Antar-Proyek

```mermaid
gitGraph
   commit id: "v1 (Baseline)"
   branch feat-pajak
   checkout feat-pajak
   commit id: "Ubah db_object/"
   checkout main
   branch feat-kupon
   checkout feat-kupon
   commit id: "Ubah db_object/"
   checkout main
   merge feat-kupon id: "Rilis Kupon Duluan (v2)"
   checkout feat-pajak
   merge main id: "git pull main"
   checkout main
   merge feat-pajak id: "Rilis Pajak (v3)"
```

<div class="grid grid-cols-3 gap-4 pt-4 text-xs">
  <div class="border border-slate-700 p-3 rounded">
    <div class="font-bold text-emerald-400 mb-1">1. Proyek Lain Rilis Duluan?</div>
    <div class="opacity-75">Tim A cukup <code>git pull origin main</code>. Delta migrasi Tim A baru dihitung saat Tim A siap rilis. Tidak ada timestamp basi.</div>
  </div>
  <div class="border border-slate-700 p-3 rounded">
    <div class="font-bold text-amber-400 mb-1">2. Fitur Batal Rilis di H-1?</div>
    <div class="opacity-75">Branch <code>feat/pajak</code> cukup JANGAN DI-MERGE ke main. Nol file migrasi yang perlu dihapus atau di-revert.</div>
  </div>
  <div class="border border-slate-700 p-3 rounded">
    <div class="font-bold text-blue-400 mb-1">3. Resolusi Konflik Bersih</div>
    <div class="opacity-75">Konflik diselesaikan di file deklaratif menggunakan standard merge conflict tool Git biasa.</div>
  </div>
</div>

---

# Di Laptop Developer: Format & Validasi Mandiri (rza_migrate_check)

Developer memiliki linter otomatis di laptop: cukup double-click atau jalankan via terminal sebelum commit.

<div class="grid grid-cols-3 gap-6 pt-6">

<div class="border border-slate-700 p-4 rounded bg-slate-900/40">
  <div class="text-blue-400 font-bold text-xs uppercase mb-1">Langkah 01</div>
  <h3 class="font-bold text-base mb-2">Edit di db_object/</h3>
  <p class="text-xs opacity-80 leading-relaxed">
    Ubah tabel atau SP langsung pada file SQL terkait di dalam <code>db_object/</code>.
  </p>
  <p class="text-xs text-red-400 mt-3 font-semibold">
    Dilarang membuat file di migrations/.
  </p>
</div>

<div class="border border-emerald-700 p-4 rounded bg-emerald-950/20">
  <div class="text-emerald-400 font-bold text-xs uppercase mb-1">Langkah 02</div>
  <h3 class="font-bold text-base mb-2">Jalankan Checker</h3>
  <p class="text-xs opacity-80 leading-relaxed">
    Double-click <code>rza_migrate_check.exe</code> atau via terminal:
  </p>
  <ul class="text-xs opacity-80 mt-2 space-y-1">
    <li>• Normalisasi UTF-8 tanpa BOM</li>
    <li>• Rapikan line ending LF & spasi</li>
    <li>• Validasi penamaan constraint</li>
  </ul>
</div>

<div class="border border-slate-700 p-4 rounded bg-slate-900/40">
  <div class="text-slate-300 font-bold text-xs uppercase mb-1">Langkah 03</div>
  <h3 class="font-bold text-base mb-2">Commit & Buka MR</h3>
  <p class="text-xs opacity-80 leading-relaxed">
    Setelah status hijau (exit 0), commit dan push ke GitLab.
  </p>
  <p class="text-xs text-emerald-400 mt-3 font-semibold">
    Kode 100% rapi dan siap masuk CI.
  </p>
</div>

</div>

---

# Di GitLab Pipeline: Satpam Otomatis Pemblokir Pelanggaran

<div class="grid grid-cols-2 gap-8 pt-4">

<div class="border border-blue-500/30 p-4 rounded bg-blue-500/5">
  <h3 class="text-blue-400 font-bold mb-2">GitLab CI Stage: .pre</h3>
  <ul class="text-xs space-y-3 opacity-85">
    <li>• <strong>Mode Read-Only (-check):</strong> Menggunakan binary Linux resmi. Menolak memodifikasi file di repository runner.</li>
    <li>• <strong>Berjalan di Stage Terdepan:</strong> Gagal langsung menghentikan pipeline sebelum job build/test aplikasi yang mahal dijalankan.</li>
    <li>• <strong>Laporan di out/output.json:</strong> Kesalahan dicetak lengkap di console dan disimpan sebagai CI artifact.</li>
    <li>• <strong>Fail Hard (Exit 1):</strong> Tombol Merge terkunci otomatis hingga diperbaiki developer di branch fitur.</li>
  </ul>
</div>

<div class="border border-red-500/30 p-4 rounded bg-red-500/5">
  <h3 class="text-red-400 font-bold mb-2">Daftar Cekalan Satpam CI</h3>
  <ul class="text-xs space-y-3 opacity-85">
    <li>• <strong>Token Haram SQL:</strong> Menolak keras adanya perintah <code>GO</code>, <code>USE &lt;db&gt;</code>, <code>CREATE DATABASE</code>, atau <code>ALTER LOGIN</code>.</li>
    <li>• <strong>Encoding Ilegal:</strong> Menolak file UTF-16 atau UTF-8 dengan BOM yang menyebabkan runner Linux crash.</li>
    <li>• <strong>File Migrasi Ilegal:</strong> Menolak MR jika developer nekat menambahkan file <code>.up.sql</code> / <code>.down.sql</code> di branch fitur.</li>
    <li>• <strong>Sintaks Rusak:</strong> Memvalidasi integritas DDL dan nama constraint sebelum menyentuh main.</li>
  </ul>
</div>

</div>

---

# Di Meja Rilis: Draft, Review, dan Seal

Dilakukan setelah seluruh branch fitur di-merge ke branch main oleh Release Manager atau CI Release Step.

<div class="grid grid-cols-3 gap-6 pt-6">

<div class="border border-blue-700/50 p-4 rounded bg-blue-950/20">
  <div class="text-blue-400 font-bold text-xs uppercase mb-1">Fase 01</div>
  <h3 class="font-bold text-base mb-2">rza_migrate_db draft</h3>
  <p class="text-xs opacity-80 leading-relaxed">
    Engine membandingkan seluruh histori migrations/ vs desired state terakhir db_object/.
  </p>
  <p class="text-xs opacity-75 mt-2">
    Menulis calon migrasi ke <code>out/draft_migrations/</code>. Folder resmi migrations/ belum tersentuh.
  </p>
</div>

<div class="border border-amber-700/50 p-4 rounded bg-amber-950/20">
  <div class="text-amber-400 font-bold text-xs uppercase mb-1">Fase 02</div>
  <h3 class="font-bold text-base mb-2">Review DBA & Arsitek</h3>
  <p class="text-xs opacity-80 leading-relaxed">
    Arsitek atau DBA meninjau script calon migrasi sebelum disahkan:
  </p>
  <ul class="text-xs opacity-80 mt-2 space-y-1">
    <li>• Cek lock tabel besar di jam sibuk</li>
    <li>• Cek kelayakan script rollback .down.sql</li>
    <li>• Cek precheck lossless konversi data</li>
  </ul>
</div>

<div class="border border-emerald-700/50 p-4 rounded bg-emerald-950/20">
  <div class="text-emerald-400 font-bold text-xs uppercase mb-1">Fase 03</div>
  <h3 class="font-bold text-base mb-2">rza_migrate_db seal</h3>
  <p class="text-xs opacity-80 leading-relaxed">
    Memindahkan file yang sudah di-review ke folder resmi <code>migrations/</code>.
  </p>
  <p class="text-xs opacity-75 mt-2">
    Timestamp 17-digit deterministik diberikan. File siap di-commit dan dideploy ke target UAT/Prod.
  </p>
</div>

</div>

---

# Operasional Server: Zero Direct SSMS Access

Tidak ada lagi user membuka SSMS ke server. Aplikasi berjalan di Pod OpenShift mengeksekusi migrasi secara terpusat.

<div class="grid grid-cols-2 gap-8 pt-4">

<div class="border border-blue-500/30 p-4 rounded bg-blue-500/5">
  <h3 class="text-blue-400 font-bold mb-1">Opsi 1: REST API (Port 8080)</h3>
  <div class="text-xs opacity-60 mb-3">Otomasi Pipeline CD & HTTP Trigger</div>
  <p class="text-xs font-semibold mb-2">Header Wajib: X-API-Key, X-Request-ID, X-Operator-NIK</p>
  <ul class="text-xs font-mono space-y-1.5 opacity-85">
    <li>GET  /api/v1/migrations/version</li>
    <li>GET  /api/v1/migrations/pending</li>
    <li>POST /api/v1/migrations/preflight</li>
    <li>POST /api/v1/migrations/up</li>
    <li>POST /api/v1/migrations/down</li>
    <li>POST /api/v1/migrations/recover-dirty</li>
  </ul>
</div>

<div class="border border-slate-700 p-4 rounded bg-slate-800/40">
  <h3 class="text-slate-200 font-bold mb-1">Opsi 2: In-Pod CLI (oc exec)</h3>
  <div class="text-xs opacity-60 mb-3">Eksekusi via Terminal Pod OpenShift</div>
  <ul class="text-xs font-mono space-y-2 opacity-85">
    <li>app migrate version</li>
    <li>app migrate pending</li>
    <li>app migrate preflight --all</li>
    <li>app migrate up --all --request-id REQ-01 --nik 12345</li>
    <li>app migrate rollback --target-request-id REQ-01 --nik 12345</li>
    <li>app migrate recover-dirty --reason "Restart"</li>
  </ul>
  <div class="mt-3 text-xs opacity-75 border-t border-slate-700 pt-2">
    Dilengkapi <code>app curl</code> untuk testing lokal tanpa butuh curl di container Alpine.
  </div>
</div>

</div>

---

# Sistem Pelacakan Berlapis: Dual-Level Audit & Concurrency Lock

<div class="space-y-4 pt-2 text-xs">

<div class="border border-blue-500/40 p-3 rounded bg-blue-500/5">
  <div class="font-bold text-blue-400 text-sm mb-1">Level 1: Audit Aplikasi & Operator · dbo.mig_operation_journal</div>
  <p class="opacity-80 leading-relaxed">
    Mencatat: <code>operation_id</code>, <code>request_id</code>, <code>nik</code> operator, <code>operation</code>, <code>status</code>, <code>before_version</code>, <code>after_version</code>, <code>applied_files</code> (JSON array).
  </p>
  <p class="text-[11px] opacity-60 mt-1 italic">
    Menjawab pertanyaan audit: "Siapa NIK operator yang mengeksekusi pada jam 02:00 pagi dan apa nomor tiketnya?"
  </p>
</div>

<div class="border border-amber-500/40 p-3 rounded bg-amber-500/5">
  <div class="font-bold text-amber-400 text-sm mb-1">Benteng Concurrency Lock · sys.sp_getapplock</div>
  <p class="opacity-80 leading-relaxed">
    Saat OpenShift menjalankan multi-pod bersamaan, engine mengambil session lock eksklusif di SQL Server. Pod kedua otomatis mengantre atau ditolak aman. Tidak akan pernah terjadi race condition atau deadlock migrasi.
  </p>
</div>

<div class="border border-slate-700 p-3 rounded bg-slate-800/40">
  <div class="font-bold text-slate-200 text-sm mb-1">Level 2: Audit Database & Statement DDL · dbo.mig_event (via dbo.usp_mig_log_event)</div>
  <p class="opacity-80 leading-relaxed">
    Mencatat detail per baris: <code>file_name</code>, <code>step_name</code>, <code>decision</code>, <code>action</code>, <code>run_key</code> (GUID unik batch).
  </p>
  <p class="text-[11px] opacity-60 mt-1 italic">
    Di dalam blok CATCH SQL Server: Merekam persis error_number, error_message, dan error_line saat terjadi kegagalan DDL.
  </p>
</div>

</div>

---

# Resiliensi Sistem: Preflight, LIFO Rollback, dan Dirty Recovery

<div class="grid grid-cols-3 gap-6 pt-6">

<div class="border border-blue-700/50 p-4 rounded bg-blue-950/20">
  <div class="text-blue-400 font-bold text-xs uppercase mb-1">Pilar 01</div>
  <h3 class="font-bold text-base mb-2">Preflight Dry-Run</h3>
  <p class="text-xs opacity-80 leading-relaxed">
    Simulasi pra-eksekusi tanpa efek samping. Memeriksa in-flight lock dan memvalidasi kecocokan file migrasi dengan target database sebelum DDL sungguhan dijalankan.
  </p>
</div>

<div class="border border-amber-700/50 p-4 rounded bg-amber-950/20">
  <div class="text-amber-400 font-bold text-xs uppercase mb-1">Pilar 02</div>
  <h3 class="font-bold text-base mb-2">LIFO Rollback</h3>
  <p class="text-xs opacity-80 leading-relaxed">
    Membatalkan perubahan satu paket rilis secara terbalik (Last-In, First-Out) berdasarkan target Request ID. Stored Procedure dan struktur tabel lama kembali ke titik awal.
  </p>
</div>

<div class="border border-emerald-700/50 p-4 rounded bg-emerald-950/20">
  <div class="text-emerald-400 font-bold text-xs uppercase mb-1">Pilar 03</div>
  <h3 class="font-bold text-base mb-2">Dirty State Recovery</h3>
  <p class="text-xs opacity-80 leading-relaxed">
    Jika server mati di tengah DDL dan database terkunci status DIRTY: panggil <code>app migrate recover-dirty</code>. Sistem pulih mandiri tanpa perlu manipulasi manual tabel sistem.
  </p>
</div>

</div>

---

# Dokumentasi Database Otomatis: Zero Database Access

Tool <code>rza_migrate_report</code> membaca Git dan migrations/ secara offline untuk menghasilkan dokumentasi lengkap tanpa perlu terhubung ke database produksi.

<div class="grid grid-cols-3 gap-6 pt-6">

<div class="border border-slate-700 p-4 rounded bg-slate-900/40">
  <div class="text-blue-400 font-bold text-xs uppercase mb-1">Output 01</div>
  <h3 class="font-bold text-base mb-1">relational_erd.md</h3>
  <div class="text-xs opacity-60 mb-3">Diagram ERD Visual (Mermaid)</div>
  <p class="text-xs opacity-80 leading-relaxed">
    Memetakan seluruh tabel, primary key, foreign key, dan observed JOIN langsung render visual di GitLab tanpa software gambar manual.
  </p>
</div>

<div class="border border-slate-700 p-4 rounded bg-slate-900/40">
  <div class="text-emerald-400 font-bold text-xs uppercase mb-1">Output 02</div>
  <h3 class="font-bold text-base mb-1">module_lineage.md</h3>
  <div class="text-xs opacity-60 mb-3">Silsilah Dependensi Modul</div>
  <p class="text-xs opacity-80 leading-relaxed">
    Menampilkan SP, Function, dan View mana saja yang terdampak saat suatu tabel diubah. Memungkinkan <em>impact analysis</em> instan sebelum rilis.
  </p>
</div>

<div class="border border-slate-700 p-4 rounded bg-slate-900/40">
  <div class="text-slate-300 font-bold text-xs uppercase mb-1">Output 03</div>
  <h3 class="font-bold text-base mb-1">7 File CSV Audit</h3>
  <div class="text-xs opacity-60 mb-3">RFC-4180 Format</div>
  <p class="text-xs opacity-80 leading-relaxed">
    Ekspor riwayat perubahan tabel, indeks, SP, tipe data, dan migration summary yang siap diolah di Excel atau sistem audit kepatuhan.
  </p>
</div>

</div>

---

# Perbandingan Tiga Generasi

| Kriteria Evaluasi | 1. Manual SSMS | 2. Plain Golang-Migrate | 3. RZA Engine Suite |
| :--- | :--- | :--- | :--- |
| **Jika Proyek Lain Rilis Duluan** | Saling timpa manual di database | Bentrok timestamp & Silent Skip | Otomatis delta fresh saat rilis (Zero Bentrok) |
| **Jika Fitur Batal di H-1** | Pusing rollback query di server | File migrasi nyangkut di rilis | Cukup jangan merge branch fitur ke main |
| **Single Source of Truth** | Nihil (harus koneksi ke DB live) | Nihil (harus baca 200 file migrasi) | Jelas di db_object/ (Desired State Utuh) |
| **Transaksi Atomik** | Sering lupa BEGIN TRAN (data korup) | Tergantung disiplin developer | 100% TRY/CATCH/THROW otomatis |
| **Audit Kepatuhan & Tracking** | Logbook manual / nihil | Hanya versi di schema_migrations | Dual-Level: NIK, Request ID & DDL Event |

<div class="mt-8 border border-blue-500/30 bg-blue-500/5 p-4 rounded text-xs leading-relaxed">
  <strong class="text-blue-400">Prinsip Arsitektur Kunci:</strong><br>
  Menggunakan rza_migrate_db bukan untuk menambah birokrasi coding, melainkan membebaskan tim dari perangkap timestamp basi dan risiko lembur darurat di tengah malam. Menjaga branch fitur bersih dari file migrasi adalah asuransi ketenangan seluruh tim saat deployment produksi.
</div>

