# Review konversi implisit seluruh procedure dan tabel — Bond Retail UAT

Tanggal: **23 Oktober 2026**. Review ini mengelompokkan bukti berdasarkan **tabel → kolom → procedure → bagian query**, dilengkapi panduan pola. CFMAST diperlakukan sama seperti tabel lain dan tidak menjadi pusat pengelompokan.

**Diperbarui setelah pemeriksaan ulang.** Ada tambahan 16 lokasi predicate pada 7 SP, comparison simple CASE yang sebelumnya terlewat, dan koreksi tipe hasil function. Baca [RECHECK.md](RECHECK.md) untuk seluruh tambahan, pasangan yang dicabut, perbandingan baseline, dan gap tersisa. Angka di dokumen ini adalah hasil terbaru.

## Mulai membaca

1. **[Rangkuman seluruh tabel dan kolom](TABLE_COLUMN_SUMMARY.md)** — daftar semua 81 kolom yang memiliki kandidat risiko akses index, peringkat berdasarkan banyaknya procedure, lalu seluruh 412 tabel yang muncul dalam konteks beda tipe.
2. **[Panduan 13 pola konversi](PATTERN_GUIDE.md)** — cara membaca arah konversi, contoh nyata dari tabel berbeda, dan pilihan perbaikan beserta syarat kesetaraan hasil.
3. Klik nama tabel pada rangkuman untuk membuka **dokumen detail tabel**. Setiap kolom memuat seluruh procedure/baris yang terpetakan, query atau mapping assignment, tipe kedua sisi, operand yang dikonversi, dan index yang sudah ada.

Dokumen detail bukan hanya kumpulan contoh terpilih. Seluruh occurrence yang berhasil dipetakan ke kolom fisik pada procedure UAT dicantumkan. Data temp/staging dan module pendukung tetap tersedia di CSV dan appendix agar tidak dicampur dengan tabel fisik.

## Temuan utama

Pemindaian dilakukan ulang terhadap **1.063 procedure UAT**, dengan metadata dari **676 tabel UAT dan 165 tabel SQL_REPLICATE**, serta 150 view dan 96 function sebagai konteks. Total input: **2.756 file SQL**, semuanya lolos pencocokan hash terhadap AST yang dipakai.

Hasil statis yang perlu dibedakan:

| Cakupan hasil | Jumlah | Makna |
|---|---:|---|
| Tabel fisik dalam seluruh konteks beda tipe di procedure | 412 | Termasuk kolom yang dikonversi, pasangan yang tidak dikonversi, dan destination assignment. |
| Kolom fisik dalam konteks tersebut | 1.580 | Dirinci melalui column_summary.csv dan dokumen tiap tabel. |
| Tabel dengan kandidat conversion kolom pada predicate | 57 | Prioritas A; dampak plan belum diamati. |
| Kolom dengan kandidat tersebut | 81 | Seluruhnya muncul di tabel pertama TABLE_COLUMN_SUMMARY.md. |
| Site pasangan operand kategori A | 456 | Comparison, IN, BETWEEN, atau LIKE pada predicate yang dapat di-bind. |
| Procedure berbeda dengan kategori A pada tabel fisik | 149 | Temp table dan expression-only tidak dijumlahkan ke angka ini. |
| Seluruh site beda tipe, termasuk family rutin dan boundary | 14.552 | Bukan 14.552 query berbahaya. Satu statement dapat mempunyai beberapa site, misalnya tiap ordinal INSERT. |

**Sering** pada laporan ini berarti banyaknya **procedure berbeda** dan **site source**, bukan execution count atau waktu yang terpakai di server. Procedure yang dipanggil sekali per bulan dapat memiliki banyak lokasi; procedure pendek yang dipanggil ribuan kali per menit bisa lebih penting secara operasional. Prioritas akhir membutuhkan actual plan dan workload.

### Kolom berulang di luar CFMAST

| Tabel / kolom | Tipe kolom → tipe pembanding | Site A | SP A | Bagian query representatif |
|---|---|---:|---:|---|
| [PS_NDM_EMPLOYEES.EMPLID](tables/sql_replicate.dbo.ps_ndm_employees.md#kolom-emplid) | CHAR(11) → INT/BIGINT/DECIMAL | 59 | 20 | `Rvmp_TRSPopulatePTOSUploadAuthen:51`: `ptcm.InsertedBy = ndme.EMPLID`. |
| [PPLMAPMasterWorkflow_TM.EMPLID](tables/sql_replicate.dbo.pplmapmasterworkflow_tm.md#kolom-emplid) | VARCHAR(11) → INT | 40 | 18 | `OMSubmitTransactionSekunder:434,436`: NIK_CS/InsertedBy bertemu EMPLID. |
| [MISOfficeInformationEODFull.regional_office_id](tables/sql_replicate.dbo.misofficeinformationeodfull.md#kolom-regional-office-id) | VARCHAR(5) → INT | 41 | 11 | `TRSOBLExpiredKitasRPT:133`: `MOI1.regional_office_id = MOI2.office_id`. |
| [ORI_Online_TR.KemenkeuStatusId](tables/retail.dbo.ori_online_tr.md#kolom-kemenkeustatusid) | VARCHAR(20) → INT | 18 | 10 | `SBNSchedPopulateNotif:57`: `IN (4,5)`; `TRSPopulateGagalBayarOri` memakai NOT IN numerik. |
| [SecurityMaster_TM.IdSeri](tables/retail.dbo.securitymaster_tm.md#kolom-idseri) | VARCHAR(100) → INT | 18 | 8 | `ORISBNIPopulateEarlyRedemption:33,61,117`: `IdSeri = @pnSeriId`. |
| [CFPAR3.CP3UIC](tables/sql_replicate.dbo.cfpar3.md#kolom-cp3uic) | VARCHAR(1) → INT | 10 | 7 | Beberapa procedure data nasabah memakai `CP3UIC = 3/8`. |
| [SecurityCouponDate_TR.Periode](tables/retail.dbo.securitycoupondate_tr.md#kolom-periode) | VARCHAR(8) → DATETIME/SMALLDATETIME | 9 | 7 | `TRSRetailUpdateCouponDate:95`: `Periode = @dNextCouponDate`. |
| [TRSParameter_TR.Value](tables/retail.dbo.trsparameter_tr.md#kolom-value) | VARCHAR(255) → SMALLINT/INT | 7 | 7 | `SRBIProcessEmailNotifMaturityDateH1:61`: `a.SecType = b.Value`. |
| [PTOSCustomer_TM.SID](tables/retail.dbo.ptoscustomer_tm.md#kolom-sid) | VARCHAR(100) → NVARCHAR | 21 | 5 | `PTOSProsesDataJobNasabah:128,142 dst.`: SID dari Excel/staging. |
| [ProCIFRiskProfile_TM.CIFNumber](tables/sql_replicate.dbo.procifriskprofile_tm.md#kolom-cifnumber) | VARCHAR(19) → BIGINT | 6 | 5 | `TRSPopulateRiskProfileNasabah:28`: `CIFNumber = CAST(@pcCIF AS bigint)`. |
| [TRSPostingGLBond_TH.TranId](tables/retail.dbo.trspostingglbond_th.md#kolom-tranid) | VARCHAR(20) → BIGINT | 14 | 4 | `SettlementPopulateHasilRevenue:27,53` serta jurnal revenue. |
| [SecuritySysParameter_TR.ParamValue](tables/retail.dbo.securitysysparameter_tr.md#kolom-paramvalue) | VARCHAR(30) → INT | 14 | 4 | `TRSPopulateApprovalFOPTransferAsset:79,144`: `tm.Purpose = par1.ParamValue`. |
| [TRSCustAccNumber_TR.NoRekInvestor](tables/retail.dbo.trscustaccnumber_tr.md#kolom-norekinvestor) | VARCHAR(20) → DECIMAL(19,0) | 6 | 3 | `GetRekeningMasterNasabahTRS:76,78`: rekening lokal dibandingkan DDMAST/DDTNEW.ACCTNO. |

Daftar ini hanya orientasi awal. Tabel lain, seperti collateral, buyback, BIS4, repo, transfer asset, rekening batch, log nasabah, serta kode cabang tetap dicantumkan di rangkuman dan detailnya. CFMAST.CFCIF sendiri memiliki 4 site A pada 2 SP; keberadaannya tidak menutupi kelompok yang jauh lebih berulang.

## Cara membaca bukti per tabel

Contoh: buka [PPLMAPMasterWorkflow_TM](tables/sql_replicate.dbo.pplmapmasterworkflow_tm.md), lalu kolom EMPLID.

```sql
-- Procedure: OMSubmitTransactionSekunder, baris 434
st.NIK_CS = un.EMPLID
-- st.NIK_CS : INT
-- un.EMPLID : VARCHAR(11), dari SQL_REPLICATE.dbo.PPLMAPMasterWorkflow_TM
```

INT memiliki precedence lebih tinggi sehingga arah konversi adalah **EMPLID VARCHAR → INT**. Dua index EMPLID sudah ada dalam snapshot; masalah ini tidak otomatis diselesaikan dengan membuat index EMPLID yang sama lagi. Pelajari domain pegawai: apakah seluruh nilai numeric, apakah leading zero penting, dan apakah satu ID mempunyai beberapa versi effective date. Samakan key pada batas sistem atau staging yang tervalidasi, lalu ukur query hasilnya.

Untuk setiap site, dokumen mencantumkan:

- Procedure dan nomor baris, dengan tautan ke source lokal.
- Predicate asli atau mapping kolom tujuan ← source untuk INSERT.
- Tipe storage kolom, tipe expression/parameter pasangan, dan deklarasi asal.
- Operand yang dikonversi dan tipe tujuannya.
- Peran kolom: converted column, expression input, source assignment, destination, atau pasangan yang tidak dikonversi.
- Context: WHERE, JOIN_ON, CASE, expression, function call, variable assignment, write assignment, atau set output.
- Existing index yang menjadikan kolom tersebut key pertama; ini metadata snapshot, bukan bukti bahwa index aktif atau dipilih live.

Satu join dapat muncul pada dua tabel detail. Itu adalah **dua endpoint dari satu site_id**, bukan dua query independen. Misalnya join `INT_column = VARCHAR_column`: kolom INT tetap dicatat sebagai pasangan, tetapi **hanya kolom VARCHAR** yang diprediksi mengalami conversion oleh aturan precedence.

## Prioritas belajar dan perbaikan

| Kategori | Kondisi | Cara menindaklanjuti |
|---|---|---|
| A_COLUMN_PREDICATE | Kolom fisik menjadi operand yang dikonversi ke numeric/date/Unicode atau representasi LIKE pada predicate. | Ambil actual plan; periksa seek predicate, residual/join conversion, reads, dan distribusi parameter. Selaraskan tipe sebelum menambah index. |
| B_EXPRESSION_OR_BOUNDARY | Konversi pada expression, source assignment, argumen function, atau hasil branch. | Utamakan kontrak tipe/nilai, overflow, truncation, format, CPU per row, dan error behavior. Bukan bukti langsung index gagal. |
| C_FAMILY_OR_ROUTINE | Perbedaan family numerik atau CHAR/VARCHAR biasa. | Jangan menganggap semuanya berbahaya. Tinjau plan atau data correctness bila ada gejala, precision/scale, collation, atau rounding. |
| PARTNER_OR_DESTINATION | Kolom adalah pasangan yang tidak dikonversi, type anchor, atau destination assignment. | Dipertahankan untuk lineage. Jangan menghitungnya sebagai conversion pada key kolom itu. |

Dalam assignment, `VARCHAR_destination = INT_source` mengonversi nilai INT menjadi VARCHAR tujuan. Itu berbeda dari comparison `VARCHAR_column = INT_parameter`, yang mengikuti precedence dan mengonversi VARCHAR ke INT. Kedua konteks dipisahkan secara eksplisit dalam review ini.

## Cakupan operator diperluas

Selain comparison, IN literal, dan BETWEEN pada audit sebelumnya, review ini membaca:

- IN subquery dengan satu kolom output yang dapat diikat.
- LIKE pada kolom numerik serta campuran ANSI/Unicode.
- Aritmetika dan concatenation dengan dua operand yang diketahui tipenya.
- CASE input/WHEN, CASE/COALESCE result branch, dan output UNION/EXCEPT/INTERSECT yang ordinal dan tipenya dapat diselesaikan.
- ISNULL replacement terhadap tipe argumen pertama; NULLIF comparison terpisah dari return type argumen pertama.
- Argumen dan declared return scalar function, schema TVF deklaratif/inline, SUM/AVG, VALUES-derived table, serta propagasi schema view/CTE/SELECT INTO pada bentuk yang didukung.
- SELECT/SET assignment ke variabel, UPDATE assignment sederhana, INSERT SELECT/VALUES dengan daftar kolom eksplisit dan urutan source yang diketahui.

Ekspansi ini menjelaskan mengapa jumlah site lebih besar daripada laporan sebelumnya. Angka 149 SP kategori A di sini **hanya kolom tabel fisik**; angka 167 pada laporan 20261022 juga mencakup temporary table dan memakai cakupan/aturan inference sebelumnya. Keduanya tidak boleh dibandingkan sebagai kenaikan/penurunan performa.

## Inventory dan lampiran

| File | Isi |
|---|---|
| [TABLE_COLUMN_SUMMARY.md](TABLE_COLUMN_SUMMARY.md) | Peringkat kolom A serta seluruh tabel berkonteks conversion; navigasi detail. |
| [PATTERN_GUIDE.md](PATTERN_GUIDE.md) | 13 pola dengan penjelasan, contoh, dan fixing bersyarat. |
| [RECHECK.md](RECHECK.md) | Tambahan temuan, koreksi kandidat lama, baseline dan gap pemeriksaan ulang. |
| `tables/*.md` | 412 dokumen tabel, seluruh occurrence procedure yang berhasil dipetakan. |
| `table_summary.csv` | Inventory seluruh **841 tabel**, termasuk NO_BOUND_MATCH_NOT_CERTIFIED. |
| `column_summary.csv` | Ringkasan 1.580 kolom yang terlibat, jumlah SP/site A dan semua context, roles, dan tipe pasangan. |
| `column_occurrences.csv` | Setiap endpoint kolom fisik, termasuk pendukung function/view. Filter module_database=retail dan file procedures/ untuk scope utama. |
| `conversion_sites.csv` | Semua 14.552 site beda tipe; termasuk temporary object, supporting module, INFO/family, dan boundary. |
| `pattern_summary.csv` | Rekonsiliasi jumlah site/SP/tabel per pola. |
| `procedure_coverage.csv` | Tepat satu baris untuk masing-masing 1.063 procedure UAT, termasuk parse failure. |
| [SUPPORTING_MODULES.md](SUPPORTING_MODULES.md) | Konversi di body view/function, dipisahkan dari lokasi langsung pada SP. |
| `unresolved_predicates.csv` | 591 pasangan predicate kompleks/ambigu/tidak cukup metadata setelah inference yang diperluas. |
| `unresolved_extensions.csv` | Observasi extension yang tipenya belum diketahui/shape tidak didukung; mencakup NULL/context dan expression nonkolom, bukan daftar temuan pasti. |
| `source_manifest.csv` | Hash 2.756 file input; AST dengan hash stale ditolak. |
| `table_columns.csv`, `existing_indexes.csv`, `object_references.csv` | Metadata dan dependency evidence untuk menelusuri tipe/index/table dasar. |
| `summary.json`, `validation.json` | Angka cakupan dan hasil pemeriksaan audit terbaru. |

CSV menggunakan UTF-8 BOM. File .gitattributes lokal mempertahankan byte artefak laporan agar normalisasi newline Git tidak membatalkan hash pada validation.json. Source snapshot aplikasi tetap berada di lokasi lokal yang diberikan pengguna; laporan tidak menyalin seluruh snapshot tersebut ke Git. Tautan source dapat dipakai selama folder snapshot tersedia.

## Batas validasi yang tetap terlihat

**Seluruh file procedure dibaca, tetapi tidak semua jalur SQL dapat dibuktikan penuh dari source ini.** Jangan menyimpulkan nol match sebagai bebas conversion.

- Satu procedure, `TRSPopulateHapusTransaksiSuratBerharga`, gagal parse penuh di baris 190 dan melaporkan internal parser error setelahnya pada baris 251. Hasilnya PARTIAL_PARSE_ERROR, bukan lulus.
- Dynamic SQL: 75 execution site di 53 SP belum diikat penuh; ledger tersendiri ada di remaining_dynamic_sql.csv. Referensi database lain yang tidak diberikan, TVF/XML/OPENROWSET tertentu, UDT, hasil SELECT * atau INSERT tanpa explicit list, INSERT EXEC, bentuk UNION/CTE kompleks, dan mapping ambigu tetap dicatat unresolved. EXEC procedure dengan parameter variable serta kontrak application/driver tidak di-bind interprocedural secara penuh. Internal coercion setiap built-in juga belum dicakup penuh.
- Parser yang dipakai adalah ScriptDom 180.140.0/TSql170Parser. Itu tidak menyatakan versi/compatibility server live.
- Lookup identifier mewarisi casefold dari binder audit; collation database live dan case-only identifier collision belum menjadi oracle. Tipe/kolasi kolom yang tersedia dicatat pada metadata.
- Temporary-table schema lintas branch/redefinition dan expression lineage majemuk tidak dimodelkan penuh. Temp object tidak digabungkan sebagai satu tabel fisik lintas procedure.
- Beda panjang VARCHAR, panjang Unicode, decimal precision/scale, nullability, atau collation pada base type yang sama tidak otomatis dihitung sebagai conversion tipe baru. Itu tetap risiko kontrak tersendiri.
- OTHER_TYPE_DIFFERENCE dapat memerlukan pembuktian apakah implicit conversion diizinkan atau justru error. Tidak semua pasangan tipe dapat dikonversi.
- Tidak ada database runtime, business procedure, DDL, index baru, atau migrasi yang dijalankan. Scan tidak membuktikan kehilangan seek, actual scan, error pada data tertentu, atau persentase peningkatan performa.

## Validasi dan reproduksi

`review_conversions.py` membaca AST baru seluruh input, memakai binder audit yang sudah ada, lalu menghasilkan evidence bertipe dan mapping endpoint. `render_review.py` membentuk ringkasan dan dokumen tabel. `validate_review.py` memeriksa hash source, lokasi evidence, rekonsiliasi jumlah, link dokumen, serta regression synthetic untuk arah assignment, LIKE, ISNULL, CASE, UDF, dan ordinal INSERT.

**Hasil validasi:** jumlah dan hasil pemeriksaan terbaru dicatat pada [validation.json](validation.json), termasuk rekonsiliasi seluruh occurrence ke 412 dokumen tabel, tautan/anchor, seluruh source span, regression semantik tambahan, dan perlindungan 2.756 source + 5.100 file migration. Bukti berada pada validation.json. Ini adalah validasi artefak dan inference statis, bukan pengujian actual execution plan.

Urutan reproduksi menggunakan PowerShell 7, Python 3, dan DLL ScriptDom yang sama dengan audit sebelumnya:

```powershell
# Gunakan extract_ast.ps1 dari folder review 20261022.
# Output AST berada di temporary directory, bukan db_object atau migrations.
python review_conversions.py --workspace <repo> --retail-ast <retail_ast.jsonl> --replicate-ast <replicate_ast.jsonl>
python render_review.py
python validate_review.py --workspace <repo> --dll <Microsoft.SqlServer.TransactSql.ScriptDom.dll>
```

Dokumen baru berada dalam direktori review 20261023 ini. Laporan 20261022 dipertahankan sebagai catatan historis dan menjadi dependency helper, bukan ditimpa agar angka baru terlihat sama.

Engine Docs Guide, rules, README, spec, FSD, plan, walkthrough, capability matrix, traceability, serta QA gate ditinjau dalam konteks boundary audit; sengaja tidak diubah karena tidak ada fitur/rule parser produksi, safe-fixer, extractor, checker, generator, pipeline, ataupun release binary yang berubah. go test/build_all tidak dipakai sebagai pengganti bukti SQL runtime. Source object dan migration dipertahankan, dengan hasil hash pada validation.json. Tidak ada commit/push otomatis untuk dokumen lanjutan ini.

Skill MSSQL tidak diubah: generalisasi yang dipakai sudah tercakup; yang ditambahkan adalah bukti dan pengorganisasian source aplikasi, bukan aturan SQL Server baru.

Rujukan semantik: [Microsoft data type precedence](https://learn.microsoft.com/en-us/sql/t-sql/data-types/data-type-precedence-transact-sql), [conversion dan assignment](https://learn.microsoft.com/en-us/sql/t-sql/data-types/data-type-conversion-database-engine), [LIKE](https://learn.microsoft.com/en-us/sql/t-sql/language-elements/like-transact-sql), [ISNULL](https://learn.microsoft.com/en-us/sql/t-sql/functions/isnull-transact-sql). Diakses 23 Oktober 2026. Kesimpulan workload/table di atas berasal dari source lokal, bukan contoh pada dokumentasi Microsoft.
