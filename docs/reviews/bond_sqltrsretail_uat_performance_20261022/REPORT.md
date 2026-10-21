# Analisis performa SQL — Bond SQLTRSRetail UAT

Tanggal analisis: **22 Oktober 2026**. Sumber: snapshot SQL, tanpa koneksi/eksekusi procedure ke database.

## Kesimpulan dan prioritas

Ada risiko konversi implisit di sisi kolom, pola query yang berpotensi membaca terlalu banyak data, dan celah search-key index dalam snapshot. **167 dari 1.063 procedure** mempunyai kandidat konversi kolom teks → numerik/tanggal atau ANSI → Unicode pada filter/join menurut aturan binding statis audit ini. Jumlah tersebut termasuk temporary table; bukan jumlah procedure yang sudah terbukti lambat.

Yang paling penting untuk kasus CFMAST: **`CFCIF` adalah `CHAR(19)`, sedangkan `CFCIF#` adalah `BIGINT`**. Jangan memperlakukan kedua kolom sebagai nama berbeda untuk satu tipe yang sama. Index untuk keduanya sudah tersedia; sejumlah query justru membandingkan representasi yang berbeda sehingga kolom teks dapat dikonversi.

Urutan tindakan yang disarankan:

1. Selaraskan join CFMAST, tipe staging PTOS, dan join identitas/rekening/pegawai. Konversi kolom dapat menghalangi seek, menambah CPU, dan gagal bila data tidak bisa dikonversi.
2. Dorong filter tanggal ke source laporan perubahan nasabah; hindari membangun seluruh histori sebelum membuang data di temporary table.
3. Uji kandidat index prioritas A: `SecurityMaster_TM.SecurityNo`, `SecurityTransaction_TT.DealNo`, `SecurityTransactionDetail_TT.DealId`, `PTOSTransaksi_TR(DealId,PTOSType)`, dan `TRSRetailCustBlock_TM.DealId`.
4. Uji kandidat lain berdasarkan logical reads, jumlah baris, frekuensi eksekusi, dan biaya INSERT/UPDATE. **18 kandidat bukan instruksi memasang 18 index sekaligus.**

Tidak ada source `db_object`, migration, binary, atau konfigurasi koneksi yang diubah. Semua SQL perbaikan dan index di paket ini adalah **usulan**.

## Isi paket dan cara membaca

| Artefak | Kegunaan |
|---|---|
| [Listing procedure](procedure_listing.md) | Nama seluruh 167 procedure dengan kandidat konversi kolom, jumlah menurut pola, tautan source, dan kandidat index terkait. |
| [Ringkasan 1.063 procedure](procedure_summary.csv) | Satu baris per procedure, termasuk yang tidak memiliki match aturan terbatas ini; jangan mengartikan nol match sebagai sertifikat aman. |
| [Listing konversi](implicit_conversions.csv) | 5.861 perbedaan tipe statis, termasuk INFO dan numeric family; lokasi, predicate, tipe dua sisi, deklarasi, arah konversi, context, dan existing leading index. |
| [Predicate CFMAST](cfmast_predicates.csv) | 171 referensi kolom CFMAST dalam comparison/IN/BETWEEN, tersebar di 33 module; lengkapi pencarian LIKE melalui query_patterns.csv. |
| [Pola query](query_patterns.csv) | Listing function-on-column, cursor, loop, wildcard, UDF, DISTINCT, UNION, optional OR, dan dynamic SQL dengan file/baris. Semuanya kandidat review. |
| [Rekomendasi index](index_candidates.md) | 18 kandidat gabungan, alasan, existing index pada tabel, prasyarat rewrite, contoh konsumen, dan CREATE INDEX. |
| [Pemetaan index ke workload](index_workload.csv) | 1.168 bukti predicate untuk key pertama kandidat, termasuk procedure dan baris. Angka bukan frekuensi runtime. |
| [SQL kandidat terpilih](recommended_indexes.sql) | Memilih **satu** kandidat melalui `@CandidateId`; default NULL menghentikan batch sebelum DDL. Belum dieksekusi. |
| [Index existing](existing_indexes.csv) / [kolom tabel](table_columns.csv) | Inventory index termasuk PK/UNIQUE dan tipe/kolasi kolom kedua snapshot. |
| [Referensi object](object_references.csv) | Resolusi tabel/view, temporary object, alias target DML, dan referensi yang tidak ditemukan/unsupported. |
| [Predicate belum terikat](unresolved_predicates.csv) | 1.422 pasangan comparison/IN/BETWEEN kompleks/ambigu atau tanpa metadata cukup; tidak dinyatakan bersih. |
| [Cakupan module](module_coverage.csv) / [ringkasan JSON](summary.json) | Rekonsiliasi scan, parser error, dan jumlah match per kategori. |
| [Validasi catalog live](catalog_validation.sql) | SELECT metadata/DMV untuk tindak lanjut DBA; tidak mengeksekusi procedure bisnis. |
| [Manifest source](source_manifest.csv) | SHA-256 setiap SQL input. Snapshot lokal, termasuk object yang belum masuk Git, menjadi bukti utama. |
| [Validasi audit](validation.json) | Hasil pemeriksaan source, referensi evidence, regression audit, syntax proposal SQL, dan batas verifikasi. |

CSV memakai UTF-8 BOM agar mudah dibuka di Excel. CSV adalah data evidence; SQL query di dalam sel bukan perintah untuk dijalankan.

## Cakupan dan tingkat kepastian

| Sumber | Tabel | File index | Procedure | View | Function |
|---|---:|---:|---:|---:|---:|
| `bond-sqltrsretail-uat/db_object` | 676 | 485 | 1.063 | 150 | 96 |
| `all_object/bond/bond-sqlreplicate/db_object` | 165 | 119 | 2 | 0 | 0 |
| Total | 841 | 604 | 1.065 | 150 | 96 |

Seluruh **2.756 file SQL** dibaca. DDL memuat **12.675 kolom bertipe yang berhasil dipetakan** dan **1.050 definisi index/PK/UNIQUE**. Jumlah file index berbeda dari jumlah index karena satu file dapat memiliki banyak CREATE INDEX dan constraint index berada pada DDL tabel.

Parser: Microsoft ScriptDom NuGet **180.140.0**, `TSql170Parser`, `QUOTED_IDENTIFIER=true`. Pemilihan parser bukan bukti SQL Server live memakai versi/compatibility 170. DLL SHA-256: `5eb811e69fb56459db7acf3d32c3a59a0a725703be6d9ed9458f8f237e697c55`.

Binding terbatas memakai scope query/DML, deklarasi parameter/variabel, CREATE TABLE temporary, SELECT INTO sederhana, CTE dalam statement, derived SELECT sederhana, serta lineage kolom view yang dapat ditelusuri. Dari **21.865 pasangan operand berkolom** yang diperiksa, **20.443** dapat diberi tipe dan **1.422** belum dapat disimpulkan. IN dihitung per tipe nilai unik, BETWEEN per tipe batas unik; ini bukan jumlah seluruh predicate SQL maupun seluruh statement.

Satu procedure memiliki parser error: [`TRSPopulateHapusTransaksiSuratBerharga`](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateHapusTransaksiSuratBerharga.sql), baris **190**: `(unInput.fullname, '') as NameInputer`, terlihat kehilangan `ISNULL`. Parser juga melaporkan internal error pada baris 251 setelahnya. File ini **PARTIAL**, bukan lulus parse; hasil AST parsial tidak menjamin cakupan seluruh isinya.

Tingkat bukti dibedakan sebagai berikut:

- **Tipe source terbukti:** deklarasi kolom/parameter dan comparison dapat ditelusuri. Ini cukup untuk mengidentifikasi arah konversi menurut precedence SQL Server.
- **Dampak index berpotensi:** conversion di sisi kolom atau pola nonsargable dapat mengurangi kemampuan seek. Optimizer mungkin memilih seek dengan residual/dynamic seek, scan yang memang optimal, atau bentuk plan lain.
- **Missing index kandidat:** belum ada key prefix yang diusulkan dalam snapshot, atau existing prefix terlalu pendek. Ini tidak membuktikan index wajib/absen pada database live.
- **Runtime belum diverifikasi:** tidak tersedia actual plan, statistics histogram, row count, Query Store, parameter distribution, CPU/reads, write rate, ataupun versi/compatibility live. Tidak ada klaim persentase percepatan.

HEAD root saat awal inspeksi: `39836cf747e249b848e0c131c69283dac4c82164`; saat packaging teramati `23b10bb63474eedf2d490c653068bf2a2ac9b0ae`. HEAD berubah di luar aksi audit ini; tidak ada commit yang dijalankan audit. Karena snapshot UAT/untracked dan SQL_REPLICATE/ignored bukan jaminan isi commit root, rujukan input yang otoritatif adalah SHA-256 pada source_manifest.csv, bukan asumsi satu revision Git.

## 1. Konversi implisit: pola, listing, dan perbaikan

Menurut [Microsoft: data type precedence](https://learn.microsoft.com/en-us/sql/t-sql/data-types/data-type-precedence-transact-sql), tipe dengan precedence lebih rendah dikonversi ke tipe yang lebih tinggi. Karena itu, perbandingan `BIGINT_column = VARCHAR_parameter` biasanya mengonversi **parameter**, sementara `VARCHAR_column = BIGINT_parameter` biasanya mengonversi **kolom**. Keberadaan conversion sendiri tidak membuktikan index gagal dipakai.

Untuk procedure UAT pada context WHERE/JOIN/DML, kategori langsung yang perlu diprioritaskan adalah:

| Pola di sisi kolom | Match operand/predicate | Procedure berbeda | Implikasi |
|---|---:|---:|---|
| Teks → numerik | 440 | 150 | Kolom teks dibandingkan dengan numeric; potensi scan/CPU/error konversi dan hilangnya perbedaan leading zero. |
| ANSI → Unicode | 87 | 14 | CHAR/VARCHAR dibandingkan dengan NCHAR/NVARCHAR; dampak seek bergantung collation dan plan. |
| Teks → tanggal | 22 | 11 | Kolom tanggal tersimpan sebagai VARCHAR dibandingkan dengan DATETIME; konversi per row dan risiko format data. |

Procedure dapat masuk lebih dari satu kategori. Match termasuk temporary table; 167 adalah union nama procedure, bukan penjumlahan tiga kolom di atas. Perbandingan di CASE/projection tetap ada dalam CSV dengan context berbeda, tetapi tidak otomatis dianggap access predicate yang menghalangi index.

`TEXT_FAMILY_ONLY` (misalnya CHAR dengan VARCHAR) adalah **INFO**, bukan temuan berbahaya otomatis. `NUMERIC_FAMILY` juga perlu dibaca hati-hati: BIT/TINYINT dibandingkan literal INT lazim terjadi dan tidak otomatis membutuhkan rewrite. Perbedaan panjang string saja tidak dicatat sebagai conversion tipe baru; kolasi dan panjang tetap perlu diselaraskan pada interface.

### 1.1 CFMAST: penelusuran ke tabel dasar

[`CFMAST_v`](../../../bond-sqltrsretail-uat/db_object/views/dbo.CFMAST_v.sql) membaca `SQL_REPLICATE.dbo.CFMAST`. Kolom di bawah diteruskan secara langsung; `CFBIR6` pada view justru merupakan expression dan tidak boleh diasumsikan sama dengan kolom dasar.

| Kolom dasar | Tipe | Existing search key |
|---|---|---|
| `CFCIF` | `CHAR(19) COLLATE Latin1_General_BIN` | `idx_cfmast_cfcif (CFCIF)` |
| `CFCIF#` | `BIGINT` | clustered `idx_cfmast_cfcif# (CFCIF#)` |
| `CFSSNO` | `CHAR(40)` | `idx_cfmast_cfssno (CFSSNO)` |
| `CFSSCD` | `CHAR(4)` | belum menjadi leading key tersendiri; kandidat IX17 memakai identitas dahulu |
| `CFBIRD` | `CHAR(7)` | belum ada index khusus birthday; IX17 menempatkannya sesudah identitas/jenis |
| `CFSNME` | `CHAR(20)` | `idx_cfmast_cfsnme (CFSNME)` |

Bukti tipe dan index: [DDL CFMAST](../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFMAST.sql), [index CFMAST](../../../all_object/bond/bond-sqlreplicate/db_object/indexes/dbo.CFMAST.sql). Tidak perlu menambah index CFCIF/CFCIF# yang sudah ada.

| Procedure / baris | Predicate aktual | Analisis dan arah perbaikan |
|---|---|---|
| `TRSRPTMasterNasabahPerubahanObligasi2:45,57,69` | `c.CFCIF=a.CIFNumber` | `CHAR(19)` vs `BIGINT` dari ProCIFChangeMaster_TH; CFCIF dikonversi numerik. Gunakan CFCIF# pada jalur numeric **setelah** konsistensi data dua representasi dibuktikan, atau normalisasi sisi histori menjadi CHAR(19) padded. |
| `TRSUpdateDataMasterNasabah:121` | `t.CFCIF = c.CFCIF` | Temporary CFCIF BIGINT bertemu CFMAST.CFCIF CHAR(19). Pilih jalur numeric `CFCIF#` atau text `t.CIFNo`; perbaiki juga schema/populasi #tmpCFMAST dan join lanjut baris 132 agar mismatch tidak berpindah ke staging. |
| `TRSSyncMasterRiskProfileData:24` | `CFMAST.CFCIF# = so.CIFNumber` | BIGINT vs ProCIFRiskProfile_TM.CIFNumber VARCHAR(19); yang dikonversi adalah kolom **risk-profile**, bukan CFMAST.CFCIF#. Index CIFNumber sudah ada. Selaraskan representasi risk-profile atau materialisasi staging numeric tervalidasi. |
| `PTOSProsesDataJobNasabah:94` | `tmp.[Nomor KTP] = v.CFSSNO` | NVARCHAR(255) Excel vs CHAR(40) CFMAST. Konversi ke Unicode di CFSSNO; normalisasi typed staging dengan pemeriksaan panjang/karakter terlebih dahulu. |
| `Rvmp_TRSPTOSProcessUpload:108` | `dt.CFBIRD = tmpCus.TglLahirJul` | CHAR(7) vs INT; birthday CFMAST dikonversi INT. Bentuk Julian CHAR(7) sekali di staging lalu bandingkan tipe yang sesuai. |
| `TRSDataNasabahRetailInstitusi:334` | `ISNULL(cfmast.CFISD6,0) <> 0` | Hasil ISNULL tetap CHAR(8), lalu dibandingkan INT. Ini conversion pada expression berkolom dan tercatat juga sebagai function-on-column. Tentukan dahulu makna kosong, spasi, NULL, dan string nol; gunakan predicate teks hanya bila ekuivalen. |
| `TRSSyncMasterIdAndFunc:16,32` | `ISNULL(mast.CFIEX7,0) != 0` | Conversion sejenis, tetapi berada pada CASE. Risiko CPU/error/nilai tetap ada; jangan melaporkannya sebagai bukti filter kehilangan seek. |

Contoh perubahan untuk laporan histori, **bersyarat kesetaraan data**:

```sql
-- Semula: kolom CHAR dikonversi BIGINT
ON c.CFCIF = a.CIFNumber

-- Pilihan jalur numeric, bila CFCIF# konsisten dengan representasi CFCIF
ON c.[CFCIF#] = a.CIFNumber
```

Normalisasi alternatif padded text juga harus memeriksa domain nonnegatif, panjang maksimum 19, dan leading zero; `RIGHT(...,19)` tanpa validasi dapat memotong nilai. Jangan menaruh `CAST(c.CFCIF AS BIGINT)` pada kolom sebagai "fix" karena tetap membungkus search key.

Contoh yang **bukan** otomatis berbahaya: `TRSGetJenisNasabahCIF:6` memakai `CFCIF = RIGHT(REPLICATE('0',19)+@pcCIF,19)` dengan parameter VARCHAR(19). CFMAST.CFCIF memang teks. Begitu pula banyak join `TreasuryCustomer_TM.CIFNo` VARCHAR(19) dengan CFMAST.CFCIF CHAR(19): bukan kasus BIGINT vs VARCHAR. Tetap cocokkan collation dan actual plan.

### 1.2 Mismatch penting di luar CFMAST

| Kelompok | Lokasi representatif | Bukti dan usulan |
|---|---|---|
| Kode seri surat berharga | `ORISBNIPopulateEarlyRedemption:33,61,117`; `ORISBNIPopulateRegisterInvestor:196`; `SBNSchedPopulateORIOrder:110` | SecurityMaster_TM.IdSeri VARCHAR(100) dibandingkan INT. Parameter/temp/source harus memakai representasi seri yang disepakati. IX02 tidak menyelesaikan conversion kolom sebelum rewrite. |
| Pegawai/NIK | `OMSubmitTransactionSekunder:434,436`; `ONEApprovalPrepareData:167`; `Rvmp_TRSPopulatePTOSUploadAuthen:51,150` | PPLMAPMasterWorkflow_TM.EMPLID VARCHAR(11), PS_NDM_EMPLOYEES.EMPLID CHAR(11), dibandingkan NIK/InsertedBy INT. Index EMPLID sudah ada. Validasi apakah leading zero signifikan; gunakan mapping canonical terindeks, bukan menambah index EMPLID identik. |
| Rekening nasabah | `GetRekeningMasterNasabahTRS:76,78,205,207`; `TRSListEarlyRedemption:109` | NoRekInvestor VARCHAR(20) dibandingkan DDMAST/DDTNEW.ACCTNO DECIMAL(19,0). Kolom rekening teks dikonversi. Buat typed staging numeric bila itu domain resmi; jangan mengubah identifier menjadi numeric tanpa memastikan leading zero. |
| Posting GL | `SettlementPopulateHasilRevenue:27,53`; `TRSHasilRevenueBondJurnalGL:459,651,685,715` | TRSPostingGLBond_TH/TT.TranId VARCHAR(20) dibandingkan DealId/BulkId BIGINT. Selaraskan di batas posting; existing clustered Product,TranId akan kurang berguna untuk comparison numerik pada TranId. |
| SearchTerm approval | `ONEApprovalPrepareData:60,90,127,174` dan versi Bond | SearchTerm VARCHAR(100) dibandingkan CIFId/DealId/OrderId BIGINT. Karena SearchTerm polymorphic, pisahkan typed key menurut jenis approval, validasi sekali, lalu index temporary result bila besar. Jangan mengubah semua SearchTerm menjadi BIGINT. |
| Kode status online | `SBNSchedPopulateNotif:57,58`; `TRSPopulateGagalBayarOri:155,245,354,478`; `trs_UpdateTreasuryCustomer_TM:638,755` | ORI_Online_TR.KemenkeuStatusId VARCHAR(20) dibandingkan literal INT/IN numerik. Gunakan literal teks setelah domain kode dibuktikan (misalnya apakah '04' ekuivalen '4'); jangan menutupi nilai invalid dengan cast kolom. |
| Revenue date | `TRSHasilRevenueBondJurnalGL:134,248,375,470` | TanggalTransaksi VARCHAR(8) dibandingkan @pdTglProses DATETIME. Jika source benar-benar YYYYMMDD, bentuk `CHAR(8)` di sisi parameter sekali; jangan cast seluruh kolom ke datetime setiap comparison. |
| Coupon period | `TRSProcessSecurityMasterAuthorization:1121,1605,1621`; `TRSRetailUpdateCouponDate_Dummy:83` | Periode VARCHAR(8) dibandingkan DATETIME/SMALLDATETIME. Selaraskan parameter YYYYMMDD atau desain ulang storage DATE melalui perubahan schema terpisah. |
| PTOS Excel | `PTOSProsesDataJobNasabah`, `PTOSProsesDataJobOutstanding`, `PTOSProsesDataJobTransaksi` | SID/produk/Deal ID NVARCHAR(255) bertemu VARCHAR(100/20) pada tabel utama. Staging berkontrak lebih baik daripada cast kolom target. Detail seluruh lokasi ada pada CSV. |
| Murex dan BIS4 | `BKPopulateSummaryHolding:281,291`; module settlement terkait pada CSV | MurexLabel VARCHAR(40) vs SECID NVARCHAR(15), dan BIS4Code VARCHAR(50) vs NVARCHAR(128). Selaraskan key staging tanpa menghilangkan karakter Unicode yang sah. |

Contoh date key yang tidak membungkus kolom, dengan asumsi kontrak harian YYYYMMDD:

```sql
DECLARE @cTanggal char(8) = CONVERT(char(8), @pdTglProses, 112);
-- Gunakan predicate ini pada kolom VARCHAR(8) dengan nilai YYYYMMDD valid.
WHERE hr.TanggalTransaksi = @cTanggal;
```

Penggantian comparison identifier numeric dengan string dapat mengubah hasil untuk '0001', '+1', spasi, scientific notation, string kosong, NULL, dan data invalid. Semua rewrite harus diuji atas kasus ini. `ISNUMERIC` tidak menjamin nilai cocok untuk BIGINT; gunakan validasi tipe tujuan, misalnya TRY_CONVERT, dengan kebijakan reject/report yang eksplisit. [Microsoft: ISNUMERIC](https://learn.microsoft.com/en-us/sql/t-sql/functions/isnumeric-transact-sql), [TRY_CONVERT](https://learn.microsoft.com/en-us/sql/t-sql/functions/try-convert-transact-sql).

## 2. Query yang perlu ditinjau dan saran fixing

### 2.1 Laporan perubahan membangun histori sebelum filter tanggal

**Procedure:** `TRSRPTMasterNasabahPerubahanObligasi2`, cabang UNION mulai baris 37; filter tanggal baru pada baris 250–251.

Setiap cabang membaca ProCIFChangeMaster_TH + Detail_TH, CFMAST_v, dan TreasuryCustomer_TM. Banyak cabang hanya berbeda kategori/kolom. Seluruh hasil dimasukkan ke #tmpMAIN, baru kemudian `DELETE` berdasarkan `CONVERT(varchar,ChangeDate,112)`. Selain tiga mismatch CFMAST tadi, bentuk ini berpotensi membaca/sort/hash banyak histori di luar tanggal yang diminta.

Usulan:

1. Dorong rentang ChangeDate ke **setiap cabang** dan cabang log alamat, sebelum UNION/materialisasi.
2. Bentuk scope histori yang sudah terfilter sekali, lalu klasifikasikan kategori dengan mapping atau APPLY yang tetap mempertahankan satu-ke-banyak kategori bila itu memang kontrak.
3. Tinjau `UNION ALL` hanya setelah duplikasi antar/dalam kategori dan kontrak output diuji. UNION sekarang menghapus baris duplikat; penggantian buta akan mengubah output.
4. Periksa IX18 di SQL_REPLICATE setelah rewrite, bukan memasangnya sambil filter tetap di temporary table.

```sql
-- Jika parameter bermakna hari kalender inclusive:
DECLARE @Start datetime = CONVERT(date, @pdStartDate);
DECLARE @EndExclusive datetime = DATEADD(day, 1, CONVERT(date, @pdEndDate));

-- Terapkan di source setiap cabang, bukan DELETE setelah SELECT INTO.
WHERE a.ChangeDate >= @Start
  AND a.ChangeDate < @EndExclusive;
```

Jika parameter bermakna timestamp persis, jangan melakukan pembulatan date di atas. Perilaku lama juga perlu diuji untuk parameter berisi jam; usulan rentang harian bukan klaim ekuivalensi untuk semua input datetime.

### 2.2 Tipe dan pemrosesan batch PTOS

**Procedure:** `PTOSProsesDataJobNasabah:10–17,89–99,128 dst.`, `PTOSProsesDataJobOutstanding`, `PTOSProsesDataJobTransaksi`, `Rvmp_TRSPTOSProcessUpload:75–111`.

- SELECT INTO dari Excel mewariskan NVARCHAR(255). Buat typed staging dengan panjang SID/produk/Deal ID/rekening yang sesuai tabel tujuan dan kolasi target. Validasi panjang serta karakter sebelum konversi; cast sempit langsung berisiko truncation/data loss.
- Parse tanggal/nominal/Julian sekali pada tahap ingest dan simpan hasil plus flag error. Hindari CONVERT/ISDATE/ISNUMERIC dan UDF yang berulang untuk input yang sama.
- Filter staging yang valid **sesudah** baris invalid diberi hasil error yang diperlukan; jangan menghilangkan audit kegagalan upload.
- `Rvmp_TRSPTOSProcessUpload:75–90` mengulangi pengecekan existing/waiting dengan JOIN dan WHERE yang tampak sama. Tinjau pemisahan Status yang dimaksud; ini potensi aturan validasi yang redundant/keliru, bukan alasan langsung menghapus salah satunya.
- Uji IX10–IX12 setelah tipe join diperbaiki, dan IX17 bila pencarian KTP masih mahal.

### 2.3 Tanggal sebagai expression pada kolom

**Contoh:** `PopulateBatchTxtFile2CMS:68`, `Rvmp_TRSPopulateSettlementTransaksiBankBeli:134` dan WHERE sejenis, `Rvmp_TRSPopulateSettlementTransaksiBankJual` pada daftar CSV.

Pola: `CONVERT(varchar(10),SettlementDate,120)` atau date-string-date round trip dalam predicate. Pilihan awal yang lebih mudah dievaluasi index adalah rentang native datetime.

```sql
DECLARE @Day date = CONVERT(date, @pdTanggal);
WHERE st.SettlementDate >= @Day
  AND st.SettlementDate < DATEADD(day, 1, @Day);
```

Sesuaikan nama parameter aktual. `CAST(datetime AS date)` kadang dapat memakai dynamic seek, sehingga jangan menyatakan semua CAST(date) pasti scan. Ukur residual predicate dan jumlah row yang dibaca. Pada procedure BankBeli, CONVERT(date) juga banyak muncul dalam **CASE projection**; itu bukan alasan menambah index atau mengganti CASE sebagai filter.

### 2.4 Pencarian nama dan leading wildcard

**Procedure:** `trs_SearchingCIFIdByName:20–22`, `RepoPopulateParamHargaModal:95`, `BuybackLogInsert:44,136`, dan seluruh lokasi `LEADING_WILDCARD` pada CSV.

SearchingCIFIdByName melakukan tiga predicate `%LOWER(@name)%`, `%UPPER(@name)%`, dan `%@name%`, lalu ROW_NUMBER/ORDER BY. Dengan kolasi binary, tiga varian tersebut juga bukan case-insensitive search yang lengkap untuk mixed case.

Saran: bila kebutuhan bisnis mendukung, ubah ke prefix search dengan normalisasi kolasi yang disepakati dan key normalisasi terindeks. Jika harus contains, B-tree CFSNME yang ada tidak memberi seek selektif berdasarkan substring, meskipun index scan/order masih mungkin berguna. Full-text hanya cocok jika pencarian berbasis kata/token memenuhi requirement; bukan pengganti substring arbitrer. Tentukan kontrak pencarian sebelum mengganti semantik.

### 2.5 Scalar UDF yang membaca tabel

**Contoh:** `TRSRPTMasterNasabahObligasi:219` memanggil `fnGetOBLAccByPriority(CIFId)`; function tersebut melakukan beberapa lookup dan TOP(1) ORDER BY RowId DESC pada rekening fallback. `fnOBLIsRiskProfileExpired:27` juga memiliki mismatch CIFNumber VARCHAR vs BIGINT yang tersembunyi di function.

Saran: untuk hasil besar, bandingkan bentuk OUTER APPLY/JOIN berbasis set atau materialisasi mapping per CIFId unik. Pertahankan prioritas rekening, NULL/empty, TOP order, dan satu baris hasil per nasabah. IX14 menargetkan lookup fallback RowId; index tidak menggantikan kebutuhan menghindari pengulangan function berbiaya besar.

Audit menemukan 707 call site UDF di 213 procedure, bukan 707 function yang pasti lambat. Sebagian scalar UDF dapat di-inline oleh SQL Server yang memenuhi syarat; versi, compatibility, function body, dan execution context harus diperiksa. [Microsoft: intelligent query processing](https://learn.microsoft.com/en-us/sql/relational-databases/performance/intelligent-query-processing-details?view=sql-server-ver16).

### 2.6 Cursor, loop, dan lookup per row

**Contoh:** `GenBatchTxtFile2CMS` dan `_Urgent` memiliki cursor pada baris 76,119,162,298,385; lookup SecurityTransaction_TT berulang termasuk baris sekitar 193–201 dan 244 dst. `OMValidateBalance:98` dan `TRSAmendSecurityTransaction_TT:811` memiliki cursor dengan urutan settlement/deal yang bisa penting bagi logika saldo.

Saran: pisahkan penghitungan kandidat secara set-based dari operasi per baris yang memang harus berurutan. Aggregate lookup transaksi sekali berdasarkan CIFId/SecId/date, lalu gunakan hasilnya saat batch diproses. Jangan mengganti transaksi, generator nomor deal, posting, atau kalkulasi saldo berurutan dengan UPDATE massal sebelum urutan/side effect terbukti ekuivalen. IX04 mungkin membantu lookup transaksi, tetapi pengurangan round trip per row tetap perlu ditinjau.

### 2.7 Optional OR, fan-out JOIN, DISTINCT, dan SELECT *

**Contoh:** `trs_ListTreasuryCustomer_TM_Original:385–415` memiliki banyak LEFT JOIN rekening dan optional `@SecAccNo IS NULL OR ...`. Bila beberapa tabel memiliki beberapa rekening per CIFId, kombinasi join dapat memperbanyak baris. Dari source saja tidak bisa dibuktikan berapa banyak atau apakah itu dikehendaki.

Saran: buktikan cardinality setiap join dan pilih satu rekening melalui aturan prioritas yang deterministik bila output memang satu nasabah satu baris. Jangan menutup join yang salah dengan DISTINCT. Untuk filter optional yang sangat berbeda selektivitasnya, uji branch terpisah, parameterized dynamic SQL, atau OPTION(RECOMPILE) pada statement yang layak; ukur biaya compile dan jangan memasang hint global.

SELECT * pada result besar memperlebar transfer/grant dan menyulitkan covering index; gunakan projection yang memang dibutuhkan. SELECT * di EXISTS bukan temuan performa sendiri karena optimizer tidak membutuhkan seluruh projection.

### 2.8 Ringkasan seluruh pattern inventory procedure

| Pattern | Lokasi AST | Procedure berbeda | Catatan |
|---|---:|---:|---|
| Function pada kolom di predicate | 2.438 | 362 | Nested function dapat dihitung lebih dari sekali; CASE projection dikecualikan dari pola ini. |
| Scalar UDF call | 707 | 213 | Perlu cek inlining dan isi function; tidak selalu bottleneck. |
| Leading wildcard pada operand berkolom | 142 | 42 | Full-text tidak selalu ekuivalen substring. |
| Cursor | 40 | 32 | Urutan/side effect bisa bagian kontrak bisnis. |
| WHILE | 69 | 58 | Loop workflow/retry tidak otomatis perlu rewrite. |
| DISTINCT | 169 | 80 | Cek fan-out sebelum menghapus. |
| UNION dengan dedup | 759 | 41 | Banyak cabang di procedure besar; jangan jumlahkan sebagai procedure. |
| OR | 1.914 | 522 | Termasuk logika bisnis/CASE; bukan semuanya optional predicate. |
| Dynamic SQL | 75 | 53 | Generated SQL tidak di-resolve ulang audit ini. |
| SELECT * | 381 | 230 | Termasuk EXISTS dan temporary query; review berdasarkan context. |

## 3. Rekomendasi index dari seluruh workload

Daftar lengkap dengan CREATE INDEX ada pada [index_candidates.md](index_candidates.md). Ringkasan kandidat prioritas awal:

| Kandidat | Search key | Procedure dengan equality key pertama | Mengapa existing belum sama |
|---|---|---:|---|
| IX01 | SecurityMaster_TM(SecurityNo) | 148 | Unique existing berawalan SecId; SecurityNo bukan key pertama. |
| IX03 | SecurityTransaction_TT(DealNo) | 22 | Unique existing berawalan DealId; DealNo bukan key pertama. |
| IX06 | SecurityTransactionDetail_TT(DealId) | 12 | Existing berawalan RowId. |
| IX12 | PTOSTransaksi_TR(DealId,PTOSType) | 10 | Existing hanya berawalan SID; perbaiki Unicode staging dahulu. |
| IX15 | TRSRetailCustBlock_TM(DealId) | 13 | Existing berawalan CollateralId. |

Contoh usulan IX01:

```sql
CREATE NONCLUSTERED INDEX [IX_Audit20261022_SecurityMaster_TM_SecurityNo]
ON dbo.SecurityMaster_TM (SecurityNo)
INCLUDE (SecCcy, SecType, ISINCode, InterestTypeID);
```

Tidak dijadikan UNIQUE: source query tidak membuktikan uniqueness bisnis/data. SecId sudah termasuk row locator dari clustered PK dan tidak perlu ditambahkan ke INCLUDE untuk alasan itu saja.

Prinsip penyaringan rekomendasi:

- Cocokkan **ordered leading keys**, bukan hanya apakah kolom muncul di suatu index. INCLUDE dan key urutan belakang tidak sama dengan search key pertama.
- Pertahankan index yang melayani workload berbeda. IX04, IX11, IX13, IX14, dan IX17 overlap existing prefix dan merupakan bahan evaluasi konsolidasi, bukan mandat menambah lalu membiarkan seluruh duplikat.
- Jangan menambah index satu kolom status/flag hanya karena sering muncul. Selektivitas rendah dan write cost dapat membuat scan lebih murah; pilih gabungan dari query yang nyata.
- Parameter table kecil seperti harga/spread/settlement bisa lebih baik discan; kategori B/C menandai perlunya pengukuran.
- Index dasar CFMAST/risk-profile/pegawai di SQL_REPLICATE sudah ada. Prioritasnya mengoreksi tipe/representasi yang dibandingkan, bukan menambah index yang sama.
- Index untuk CFMAST_v ditempatkan pada **tabel dasar di SQL_REPLICATE**. Jangan mencoba CREATE INDEX pada view lintas database ini sebagai shortcut.
- Setiap kandidat harus lolos pemeriksaan live: index aktif/filtered/direction/INCLUDE, table size, histogram, DML rate, ruang disk/log, dan batas perubahan replication. Tidak ada saran drop otomatis.

Microsoft juga menempatkan [missing-index suggestions sebagai salah satu sumber bukti, bukan keputusan final](https://learn.microsoft.com/sql/relational-databases/indexes/tune-nonclustered-missing-index-suggestions). Audit source ini memiliki keterbatasan lebih besar karena tidak melihat estimasi optimizer atau runtime.

## 4. Validasi sebelum dan sesudah perbaikan

1. Jalankan SELECT metadata yang relevan dari catalog_validation.sql pada UAT dan SQL_REPLICATE. Rekonsiliasi tipe/kolasi/index snapshot dengan live. Catat keterbatasan permission; hasil kosong bukan bukti object tidak ada.
2. Pilih procedure/statement yang mahal dari Query Store atau workload aktual. Gunakan kombinasi parameter sempit/lebar, common/rare, NULL, dan edge case domain.
3. Ambil actual plan dan STATISTICS IO/TIME sebelum perubahan. Periksa CONVERT_IMPLICIT di seek/filter/join, residual, row read vs row returned, estimated vs actual rows, sort/hash spill, dan lookup.
4. Uji **satu** rewrite atau index. Bandingkan output termasuk jumlah/duplikasi/NULL/leading zero, urutan yang menjadi kontrak, dan error behavior. Untuk DML, gunakan fixture UAT yang disetujui dan periksa efek bisnis/audit, bukan menjalankan procedure produksi sembarang.
5. Ukur logical reads/CPU/duration, p95 menurut distribusi parameter, serta INSERT/UPDATE/delete cost. Jangan menganggap index scan selalu salah; untuk result besar bisa optimal.
6. Untuk index baru, inventarisasi overlap dan strategi rollback DROP kandidat **yang baru dibuat**. Jangan menghapus index existing berdasarkan kemiripan saja.
7. Setelah perubahan SQL dipilih untuk implementasi, lakukan melalui desired state dan migration baru yang sesuai repository. Audit ini tidak memodifikasi migration history.

## 5. Batas cakupan yang sengaja tidak ditebak

- Dynamic SQL: 75 lokasi AST di 53 procedure; string SQL tidak dianggap sudah dieksekusi atau fully bound. Periksa generated text, tipe sp_executesql, dan actual plan per jalur.
- UNION/recursive CTE/derived output kompleks, temp table dengan schema berubah menurut branch, XML/TVF/OPENROWSET/rowset yang tak didukung, computed column/UDT kompleks, dan expression aritmetika campuran tidak diinferensi secara penuh. Unresolved tetap dilaporkan.
- Pemetaan nama menggunakan casefold untuk lookup snapshot; spelling asli source tetap dapat diperiksa lewat file/baris. Database identifier collation live tidak tersedia; case-only object collision dan kebutuhan quoted identifier/case harus dikonfirmasi.
- Default database untuk nama satu/dua bagian diasumsikan snapshot asal. Nama tiga bagian dipertahankan; `SQL_TRSRetail.*` tidak otomatis dianggap sama dengan snapshot UAT tanpa bukti. Referensi database lain yang tidak diberikan tetap external/not-found.
- Statistik index, disabled state live, filtered selectivity, partitioning fisik, plan cache, dan data distribusi tidak dapat disimpulkan dari DDL snapshot. Parser lulus bukan bukti object berhasil di-deploy atau query memilih index.
- Ranking kandidat terutama memakai equality/range pada kolom fisik yang dapat di-bind, ditambah review manual kasus sort/TOP dan date pushdown. Ini bukan optimasi global seluruh kombinasi JOIN/GROUP BY/ORDER BY; tidak mengklaim bahwa 18 kandidat merupakan himpunan index optimal atau mencakup setiap peluang tuning.
- CSV conversion memeriksa comparison, IN literal per tipe, dan BETWEEN. Ia bukan binder penuh seluruh operator LIKE/UNION/CASE return/assignment/UDF argument. Risiko tambahan operator tersebut ditelusuri melalui pattern inventory dan temuan manual, bukan dinyatakan tidak ada.
- Tidak ada connection string atau secret yang dibaca/dibutuhkan. Procedure bisnis, DDL, dan query DMV tidak dijalankan terhadap SQL Server.

## 6. Reproduksi, evidence, dan dokumen repository

Tooling audit mandiri ada di direktori laporan ini: `extract_ast.ps1`, `analyze.py`, `package_report.py`, dan `validate_audit.py`. ScriptDom merupakan dependency analisis lokal, bukan dependency baru engine. AST mentah disimpan di temporary directory, tidak dimasukkan sebagai source engine. Hash input divalidasi sebelum AST dipakai kembali.

Contoh urutan (gunakan PowerShell 7 dan Python 3, ganti path DLL/temporary directory yang sesuai):

```powershell
./extract_ast.ps1 -SourceRoot <retail-db_object> -ScriptDomDll <ScriptDom.dll> -OutputFile <temp>/ast.jsonl
./extract_ast.ps1 -SourceRoot <replicate-db_object> -ScriptDomDll <ScriptDom.dll> -OutputFile <temp>/replicate_ast.jsonl
python analyze.py --workspace <repo-root> --ast <temp>/ast.jsonl --replicate-ast <temp>/replicate_ast.jsonl
python package_report.py
python validate_audit.py --workspace <repo-root> --dll <ScriptDom.dll>
```

**Dokumen ditambahkan:** REPORT.md, procedure_listing.md, index_candidates.md, beserta data evidence dan SQL usulan pada direktori ini.

**Dokumen ditinjau dan sengaja tidak diubah:** Engine Docs Guide, README, rules, spec, FSD, plan, walkthrough, capability matrix, traceability matrix, dan QA gate. Alasannya: ini analisis satu snapshot, bukan penambahan fitur/rule engine, safe-fixer, dukungan parser produksi, atau perubahan flow extractor/checker/generator/pipeline. Hasil audit tidak dipromosikan menjadi klaim kemampuan engine.

**Boundary lintas tool:** extractor, checker apply, checker -check, generator, pipeline, dan binary release tidak berubah. Tidak ada migration dihasilkan. Karena tidak ada source engine/release yang diubah, go test ./..., hardening gate, build_all.bat, checksum release, dan Linux checker smoke tidak menjadi bukti yang relevan untuk kesimpulan tuning ini dan tidak dijalankan sebagai pengganti pengukuran SQL. Validasi yang dijalankan khusus untuk ketepatan artefak audit tercatat pada validation.json.

Tidak ada perubahan skill MSSQL yang diperlukan: aturan precedence, lineage view, dan batas kesimpulan statis yang dipakai sudah tercakup dalam skill; temuan di sini merupakan kondisi khusus snapshot aplikasi, bukan perilaku SQL Server baru yang perlu dipromosikan menjadi aturan skill.
