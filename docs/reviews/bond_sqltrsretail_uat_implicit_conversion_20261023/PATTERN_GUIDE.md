# Panduan belajar 13 pola konversi implisit

Review source Bond Retail UAT, 23 Oktober 2026. Baca [ringkasan tabel/kolom](TABLE_COLUMN_SUMMARY.md) untuk peringkat dan [review utama](REVIEW.md) untuk cakupan. Semua contoh bertanda *aktual* berasal dari source; contoh bertanda *ilustrasi* membantu memahami aturan dan bukan patch siap deploy.

## Dua aturan yang harus dibedakan

**Comparison/expression biasa:** tipe dengan precedence lebih rendah dikonversi ke tipe lebih tinggi. Contoh VARCHAR vs INT: VARCHAR dikonversi ke INT. SQL Server tetap perlu mempunyai konversi yang didukung; bila tidak, expression bisa error. [Microsoft: precedence](https://learn.microsoft.com/en-us/sql/t-sql/data-types/data-type-precedence-transact-sql).

**Assignment/call boundary:** nilai source dikonversi ke tipe destination atau parameter yang dideklarasikan. Contoh memasukkan INT ke kolom VARCHAR: INT menjadi VARCHAR, meskipun INT mempunyai precedence lebih tinggi. [Microsoft: data type conversion](https://learn.microsoft.com/en-us/sql/t-sql/data-types/data-type-conversion-database-engine).

Dengan demikian, jangan membaca semua `kolom = nilai` sebagai aturan yang sama. Dalam WHERE itu comparison; dalam UPDATE SET itu assignment.

| Pertanyaan | Data yang dibaca |
|---|---|
| Apa tipe storage sebenarnya? | DDL tabel dasar, bukan nama variabel, alias view, atau dugaan bahwa semua ID pasti BIGINT. |
| Operator apa yang dipakai? | Comparison, LIKE, arithmetic, ISNULL, CASE, UNION, assignment, atau argument binding. |
| Bagian mana yang dikonversi? | Kolom fisik, expression berkolom, literal, parameter, source assignment, atau argument. |
| Apakah itu access predicate? | WHERE/JOIN bisa memengaruhi akses index; projection/assignment tidak otomatis demikian. |
| Apakah conversion merusak seek? | Tidak bisa dipastikan dari perbedaan tipe saja. Periksa actual plan, residual, reads, dan collation. |
| Apakah mengganti tipe tetap setara? | Uji leading zero, spasi, NULL/kosong, invalid numeric/date, overflow, Unicode, rounding, dan duplikasi hasil. |

Exact panjang/precision/scale hasil gabungan expression tidak diinferensi sebagai fakta. Tipe yang tertulis pada operand adalah deklarasi atau family hasil yang diketahui; `target_type` pada comparison/expression menyatakan family, bukan seluruh detail `CONVERT_IMPLICIT` yang akan tampil di plan.

## P01 — Kolom teks dibandingkan dengan numeric

### Bentuk yang sering muncul

```sql
-- Ilustrasi
WHERE text_column = @numeric_parameter
JOIN other_table n ON text_column = n.numeric_column
WHERE text_status IN (1, 2, 3)
```

Kolom CHAR/VARCHAR berada pada sisi precedence rendah dan dapat dikonversi per row. Dampaknya bisa berupa berubahnya access path, CPU tambahan, estimasi cardinality yang buruk, serta conversion error bila terdapat nilai nonnumeric. Angka dengan leading zero juga dapat menjadi sama ketika dibandingkan secara numeric.

### Contoh aktual dari beberapa domain

| Domain | Kolom | Procedure / bagian | Arah |
|---|---|---|---|
| Pegawai | [PPLMAPMasterWorkflow_TM.EMPLID](tables/sql_replicate.dbo.pplmapmasterworkflow_tm.md#kolom-emplid) | `OMSubmitTransactionSekunder:434`: `st.NIK_CS = un.EMPLID` | VARCHAR(11) EMPLID → INT. |
| Regional kantor | [MISOfficeInformationEODFull.regional_office_id](tables/sql_replicate.dbo.misofficeinformationeodfull.md#kolom-regional-office-id) | `TRSOBLExpiredKitasRPT:133`: `regional_office_id = office_id` | VARCHAR(5) regional key → INT. |
| Produk | [SecurityMaster_TM.IdSeri](tables/retail.dbo.securitymaster_tm.md#kolom-idseri) | `ORISBNIPopulateEarlyRedemption:33`: `IdSeri = @pnSeriId` | VARCHAR(100) → INT. |
| Status | [ORI_Online_TR.KemenkeuStatusId](tables/retail.dbo.ori_online_tr.md#kolom-kemenkeustatusid) | `SBNSchedPopulateNotif:57`: `IN (4,5)` | VARCHAR(20) → INT. |
| Rekening | [TRSCustAccNumber_TR.NoRekInvestor](tables/retail.dbo.trscustaccnumber_tr.md#kolom-norekinvestor) | `GetRekeningMasterNasabahTRS:76`: `a.NoRekInvestor = b.ACCTNO` | VARCHAR(20) → DECIMAL(19,0). |
| Posting | [TRSPostingGLBond_TH.TranId](tables/retail.dbo.trspostingglbond_th.md#kolom-tranid) | `SettlementPopulateHasilRevenue:27`: `lg.BulkId = th.TranId` | VARCHAR(20) → BIGINT. |
| Parameter generik | [TRSParameter_TR.Value](tables/retail.dbo.trsparameter_tr.md#kolom-value) | `SRBIProcessEmailNotifMaturityDateH1:61`: `a.SecType = b.Value` | VARCHAR(255) → SMALLINT. |

### Pilihan perbaikan

- Bila domain resminya teks, ubah signature/staging/literal pembanding agar sesuai tipe kolom. Jangan langsung mengubah kolom teks menjadi numeric hanya karena namanya ID.
- Bila domain numeric tetapi tersimpan teks, validasi dan normalisasi sekali pada ingest/staging atau rancang storage/typed key melalui perubahan schema yang ditinjau. Terapkan error policy yang eksplisit.
- Untuk tabel parameter generik, normalisasi hanya subset kode yang memang numeric. Value untuk kode lain dapat berisi tanggal, email, atau teks; jangan CAST semua isi tabel.
- Jangan mengandalkan urutan WHERE untuk menjamin validasi berjalan lebih dahulu. `ISNUMERIC(Value)=1 AND CAST(Value AS int)=...` tidak membuktikan CAST aman; valid numeric belum tentu valid INT. Gunakan validasi tipe tujuan yang tepat dan jangan diam-diam membuang nilai invalid.

Contoh status, **hanya jika kode teks `'4'` dan `'5'` adalah bentuk canonical**:

```sql
-- Semula
WHERE KemenkeuStatusId IN (4, 5);

-- Usulan bersyarat
WHERE KemenkeuStatusId IN ('4', '5');
```

Uji `'04'`, `'+4'`, `'4 '`, NULL, kosong, dan nilai nonnumeric. Bentuk string dapat menghasilkan matching berbeda dari numeric comparison. Bukan sekadar mengganti tanda kutip tanpa memeriksa data.

**Index bukan pengganti fixing tipe:** regional_office_id sudah memiliki index NC1; EMPLID juga memiliki index. Menambah index identik tidak menghapus conversion pada key.

## P02 — CHAR/VARCHAR bertemu NCHAR/NVARCHAR

Contoh aktual:

```sql
-- PTOSProsesDataJobNasabah:128
tmp.[SID] = tm.[SID]
-- tmp.SID mewarisi NVARCHAR(255) dari Excel.
-- PTOSCustomer_TM.SID adalah VARCHAR(100).
```

NVARCHAR memiliki precedence lebih tinggi daripada VARCHAR. Kolom ANSI dapat menjadi Unicode pada comparison. Efek seek tergantung collation dan bentuk plan, sehingga ini kandidat prioritas, bukan bukti pasti scan.

Kelompok lain: [PTOSOutstanding_TR.SID/KodeProduk](tables/retail.dbo.ptosoutstanding_tr.md), [PTOSTransaksi_TR.DealId](tables/retail.dbo.ptostransaksi_tr.md), [TRSMappingMurexBIS4.BIS4Code](tables/retail.dbo.trsmappingmurexbis4.md), [BKMappingSec_TM.MurexLabel](tables/retail.dbo.bkmappingsec_tm.md), [SecurityMaster_TM.ISINCode](tables/retail.dbo.securitymaster_tm.md#kolom-isincode).

Saran staging untuk domain SID yang memang terbukti bisa disimpan ANSI:

```sql
-- Ilustrasi typed staging; validasi karakter/panjang dilakukan sebelum pengisian.
CREATE TABLE #SIDStage
(
    SID varchar(100) COLLATE Latin1_General_BIN NOT NULL
);
```

Samakan tipe/kolasi di batas ingest, bukan `CAST(tm.SID AS nvarchar(255))` pada kolom tujuan. Uji panjang sesudah encoding, karakter yang tidak representable, trailing spaces, dan equality. Jika Unicode diperlukan, rancang storage Unicode yang konsisten; jangan memangkas karakter agar warning hilang.

## P03 — Tanggal disimpan teks, dibandingkan dengan date/time

Contoh aktual: [SecurityCouponDate_TR.Periode](tables/retail.dbo.securitycoupondate_tr.md#kolom-periode) VARCHAR(8) dibandingkan @dNextCouponDate SMALLDATETIME; [TRSHasilRevenue_TM.TanggalTransaksi](tables/retail.dbo.trshasilrevenue_tm.md#kolom-tanggaltransaksi) VARCHAR(8) dibandingkan @pdTglProses DATETIME.

Kolom teks dapat dikonversi menjadi date/time tiap comparison. Jika kontraknya YYYYMMDD dan pencariannya harian, bentuk key pada sisi parameter:

```sql
-- Ilustrasi untuk kontrak harian YYYYMMDD yang sudah dibuktikan.
DECLARE @DayKey char(8) = CONVERT(char(8), @pdTglProses, 112);
WHERE hr.TanggalTransaksi = @DayKey;
```

Jika storage sudah DATETIME dan kebutuhan satu hari, gunakan rentang native:

```sql
DECLARE @Day date = CONVERT(date, @pdTglProses);
WHERE t.EventTime >= @Day
  AND t.EventTime < DATEADD(day, 1, @Day);
```

Keduanya melayani kontrak berbeda. Membulatkan timestamp menjadi hari dapat mengubah hasil. Validasi tanggal invalid, locale, empty string, timezone yang relevan, serta upper bound. Migrasi storage ke DATE adalah pekerjaan schema terpisah dan tidak dijalankan oleh audit ini.

## P04 — Konversi pada parameter/literal/expression, bukan key pasangan

```sql
-- Ilustrasi
WHERE BigIntKey = @TextInput;
```

Parameter teks umumnya dikonversi ke BIGINT. Itu tidak sama dengan VARCHAR_column = BIGINT_parameter. Index BigIntKey masih mungkin dipakai; risiko tetap ada pada invalid input, overflow, estimasi, dan parameterization.

Di dokumen tiap tabel, peran `COUNTERPART_COLUMN_NOT_CONVERTED` menjaga perbedaan ini. Jangan menganggap semua tabel yang muncul dalam satu join mengalami conversion key. Ini sebabnya total tabel berkonteks conversion lebih besar daripada total tabel kategori A.

Perbaikan utama: validasi parameter ke tipe tujuan dengan error yang jelas sebelum query. Jangan mengubah key BIGINT menjadi VARCHAR hanya agar sama dengan input yang longgar.

## P05 — Aritmetika dan concatenation mencampur tipe

Contoh aktual yang relatif rutin: `CustodyUpdateEmailStatusCorporateAction:12,24` melakukan `CountRetry TINYINT + literal INT 1`; penghitungan menggunakan family numerik dan tidak otomatis membutuhkan tuning.

Contoh aktual text/Unicode: `BKPopulateSummaryHolding:288` menggabungkan `@cMsg + '- ' + mx.SECID NVARCHAR(15)`; hasil gabungan menggunakan Unicode. Panjang hasil gabungan tidak boleh disimpulkan hanya dari panjang satu operand.

Ilustrasi yang rentan salah makna:

```sql
-- Bila Code VARCHAR dan Seq INT, + tidak otomatis berarti concatenate.
SELECT Code + Seq;

-- Bila kontraknya benar-benar string:
SELECT Code + CONVERT(varchar(20), Seq);
```

Jika kontraknya perhitungan, validasi dan ubah source teks ke tipe numeric yang tepat. Jika kontraknya penyusunan string, pilih konversi eksplisit panjang/format yang memadai dan perhatikan NULL. Mengganti dengan CONCAT dapat mengubah perilaku NULL; jangan menyatakan ekuivalen tanpa pengujian.

Arithmetic date/time mempunyai aturan operator khusus dan tidak diinferensi dari precedence umum oleh extension review ini. Observasi tersebut ditandai terpisah pada unresolved_extensions.csv.

## P06 — LIKE pada numeric atau campuran Unicode

Contoh aktual di luar CIF:

| Kolom | Tipe | Procedure | Predicate |
|---|---|---|---|
| [DDMAST.BRANCH](tables/sql_replicate.dbo.ddmast.md#kolom-branch) | INT | `GetAccountTaxAmnesty:45` | `dd.BRANCH LIKE '70%'` |
| [MISOfficeInformationEODFull.office_id](tables/sql_replicate.dbo.misofficeinformationeodfull.md#kolom-office-id) | INT | `TRSSearchCostCenter:17,25,42` | `office_id LIKE '%' + @cCol1 + '%'` |
| [EarlyRedemption_TT.EarlyId](tables/retail.dbo.earlyredemption_tt.md#kolom-earlyid) | BIGINT | `TRSSearchSecIdEarlyRedem:44,60` | `EarlyId LIKE RTRIM(LTRIM(@cCol1)) + '%'` |
| [SecurityTransferAset_TM.DealId](tables/retail.dbo.securitytransferaset_tm.md#kolom-dealid) | INT | `TRSTransferAssetDealIdSrc:84,146` | `tm.DealId LIKE RTRIM(LTRIM(@cCol1)) + '%'` |

LIKE mencari pola karakter; operand numeric dapat dikonversi ke representasi karakter. **Aturannya bukan memaksa pattern `'70%'` menjadi INT seperti ordinary equality.** Campuran Unicode di LIKE juga menggunakan aturan karakter yang relevan. [Microsoft: LIKE](https://learn.microsoft.com/en-us/sql/t-sql/language-elements/like-transact-sql).

Tentukan requirement:

- Lookup nomor tepat: validasi input menjadi INT/BIGINT lalu gunakan `=`.
- Prefix kode yang sebenarnya teks: storage atau search representation terindeks perlu konsisten dengan domain.
- Substring arbitrer: B-tree key numeric tidak langsung menyelesaikannya; pertimbangkan strategi pencarian yang sesuai requirement.

Jangan mengganti `BRANCH LIKE '70%'` dengan rentang numeric tetap tanpa bukti lebar kode. Nilai 70, 700, dan 7000 sama-sama mempunyai prefix teks 70. Casting kolom secara eksplisit menjadi VARCHAR hanya membuat conversion terlihat; bukan memperbaiki sargability.

## P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION

Simple CASE mempunyai **dua proses tipe terpisah**. Input dibandingkan dengan setiap WHEN; THEN/ELSE menentukan tipe hasil. Contoh aktual `ORIOnlineBrowseTrx:98`: `CASE trx.KemenkeuStatusId WHEN 3 ...`, dengan KemenkeuStatusId VARCHAR(20). Input teks dibandingkan dengan INT walaupun semua hasil THEN/ELSE mungkin teks. `TRSReportSKFBankBeli:59` juga memakai `CASE tc.JenisIdentitas WHEN 1 ...` untuk kolom CHAR(1).

Untuk domain kode teks, pertimbangkan WHEN `'3'` atau `'1'` setelah menguji leading zero, spasi, NULL, dan kode nonnumeric. Equality numeric dan string dapat menghasilkan kelompok nilai berbeda. Pemeriksaan ulang menemukan 19 comparison teks-ke-numeric pada simple CASE; semuanya expression context B, bukan bukti akses index gagal. [Microsoft: CASE](https://learn.microsoft.com/en-us/sql/t-sql/language-elements/case-transact-sql).

Contoh aktual penting: `TRSPopulateCboParam:7`:

```sql
SELECT 0 AS ParamId, 'All' AS ParamValue
UNION ALL
SELECT ParamId, ParamValue
FROM dbo.TRSGeneralParam_TR
WHERE GroupId = 1 AND ParamId NOT IN (2,3);
```

TRSGeneralParam_TR.ParamId adalah VARCHAR(20). Angka `0` pada branch pertama dapat menjadikan output ParamId bertipe INT, sehingga kolom teks pada branch kedua dikonversi. Di procedure yang sama, NOT IN numerik juga menimbulkan conversion pada filter. Perbaiki **keduanya** bila domain ParamId memang teks:

```sql
-- Usulan bersyarat: output/caller dan domain ParamId telah disepakati teks.
SELECT CAST('0' AS varchar(20)) AS ParamId, 'All' AS ParamValue
UNION ALL
SELECT ParamId, ParamValue
FROM dbo.TRSGeneralParam_TR
WHERE GroupId = 1 AND ParamId NOT IN ('2','3');
```

Perubahan ini mengubah wire type output menjadi VARCHAR; periksa caller, bukan hanya plan query. CASE/COALESCE juga memilih family tipe hasil berdasarkan operandnya, bukan tipe branch yang kebetulan dipilih pada satu input. Ini tidak berarti setiap expression branch selalu dieksekusi.

Dalam UNION dengan dedup, perubahan representasi dapat mengubah kesetaraan/duplikasi, misalnya numeric normalization terhadap leading zero. Pertahankan ordinal, tipe, jumlah baris, dan aturan dedup. Bentuk yang ordinalnya tidak jelas atau mengandung SELECT * tetap unresolved; tidak dibuat mapping yang ditebak.

## P08 — INSERT/UPDATE/assignment menuju tipe tujuan

Contoh aktual:

- `BuybackLogInsert:73`: TRSSettlementBuyBackLog_TH.OrderId **BIGINT** diisi dari DealIdSekunder **VARCHAR(100)**.
- Statement yang sama mengisi ProcessUserNik **INT** dari SupervisorId **VARCHAR(10)**, dan UniqueId **UNIQUEIDENTIFIER** dari Remark3 **VARCHAR(40)**.
- `CustodySaveUploadOrchid:50`: NIKUploadOrichid BIGINT diisi @pnNik INT; ini pelebaran numeric yang mungkin memang wajar.
- `EBWInvestmentInquiry#ORI:32`: variabel DATETIME diisi dari kolom parameter Value VARCHAR(255).

Assignment mengikuti destination. Pada source INT → kolom VARCHAR, arah konversinya adalah INT menjadi VARCHAR, berlawanan dengan comparison VARCHAR = INT. Review utama memetakan source/destination terpisah sehingga destination tidak dilabeli key yang kehilangan seek.

Perbaikan: validasi source terhadap tipe/format destination sebelum write, laporkan invalid value, dan selaraskan staging. Uji overflow, GUID invalid, tanggal locale, rounding, truncation, dan Unicode loss. Jangan mengganti CAST/CONVERT dengan TRY_CONVERT lalu menyimpan NULL secara diam-diam jika error sebelumnya bagian kontrak.

INSERT hanya dipetakan bila explicit target column list dan urutan source diketahui. Contoh `INSERT T(B,A) SELECT X,Y` berarti X → B dan Y → A. Tidak boleh menyamakan kolom berdasarkan nama atau asal alfabet. INSERT EXEC, star, atau output ordinal ambigu tidak dianggap sudah terselesaikan.

## P09 — ISNULL dan NULLIF: return type sama acuannya, comparison berbeda

```sql
-- Ilustrasi, bila CodeText VARCHAR:
ISNULL(CodeText, 0)
```

Replacement angka 0 dikonversi ke tipe CodeText. Jika hasil tersebut lalu dibandingkan dengan numeric 0, terjadi **tahap conversion lain**:

```sql
WHERE ISNULL(CodeText, 0) <> 0;
-- tahap 1: replacement numeric → tipe teks argumen pertama;
-- tahap 2: hasil teks ISNULL → numeric untuk comparison.
```

Contoh rutin pada banyak tabel: `ISNULL(TaxOnInterest DECIMAL,0 INT)` mengubah replacement menjadi decimal. Itu lazim dan bukan temuan index kritis. Literal NULL sebagai argumen pertama mempunyai aturan return type tersendiri; jangan menerapkan asumsi tipe NULL.

Untuk tanggal, `ISNULL(DateColumn,'')` tidak berarti hasilnya string: replacement dapat menjadi nilai tanggal. Menggantinya dengan IS NULL saja bisa mengubah perilaku untuk sentinel yang sudah tersimpan. Tentukan NULL/empty/sentinel contract terlebih dahulu. [Microsoft: ISNULL](https://learn.microsoft.com/en-us/sql/t-sql/functions/isnull-transact-sql).

**NULLIF berbeda:** `NULLIF(CodeText, 0)` membandingkan argumen dengan aturan comparison, sehingga teks dapat dikonversi ke INT untuk pengujian equality; return type tetap tipe CodeText. Jangan memakai arah replacement ISNULL untuk comparison NULLIF. Dalam snapshot ini ada 10 pasangan NULLIF beda tipe, seluruhnya family numerik; bukan 10 temuan index kritis baru. [Microsoft: NULLIF](https://learn.microsoft.com/en-us/sql/t-sql/language-elements/nullif-transact-sql).

## P10 — Argumen function mengikuti signature

Contoh aktual: `CustodyPopulateEmailCorporateAction:17,21,22` memanggil `fnFormatStringThSeparator` dengan Nominal/GrossAmount/TaxOnInterest DECIMAL(20,2), sementara parameter @cValue pada signature adalah VARCHAR(100). Nilai numeric diubah menjadi teks saat masuk function.

Contoh lain: `CIFInquiryTaxAmnesty:16` melewatkan TCEXP7 DECIMAL(7,0) ke @nDate INT pada `fnJulianToDatetime`.

Konversi di sini mengikuti signature, tidak berarti kolom base di WHERE kehilangan seek. Periksa apakah boundary conversion memang bagian format output atau justru salah kontrak. Selaraskan signature dan pemanggil jika diperlukan, uji rounding/overflow/format, serta compatibility caller lain. Review hanya memakai function signature yang ada di snapshot; function eksternal/TVF/CLR atau argumen ambigu tidak ditebak.

## P11 — Family numerik berbeda, sering rutin

Contoh berulang: SecurityTransaction_TT.TrxType SMALLINT dibandingkan literal INT, flag BIT dibandingkan 0/1 INT, dan counter TINYINT ditambah INT. Family numeric berbeda dapat memerlukan conversion, tetapi ini bukan alasan otomatis mengubah seluruh parameter, menambah index, atau mengganti tipe storage.

Pelajari: nilai ekstrem, precision/scale, integer division, money/decimal rounding, dan actual plan. Jika mengubah cast arithmetic untuk membuat output decimal, verifikasi aturan hitung bisnis. Review ini tidak menginferensi exact decimal result precision/scale untuk expression gabungan.

Hasil function juga perlu diperiksa terpisah dari input. SUM/AVG pada SMALLINT menghasilkan INT, SMALLMONEY menghasilkan MONEY, REAL menghasilkan FLOAT; declared scalar UDF return type dapat berbeda dari parameter/body expression. Panjang, precision/scale, dan schema output gabungan harus diuji terhadap kontrak. [Microsoft SUM](https://learn.microsoft.com/en-us/sql/t-sql/functions/sum-transact-sql), [AVG](https://learn.microsoft.com/en-us/sql/t-sql/functions/avg-transact-sql).

## P12 — CHAR/VARCHAR biasa dan representasi teks

LEFT/RIGHT menghasilkan VARCHAR atau NVARCHAR, sehingga `RIGHT(numeric_expression, n)` tidak mengembalikan numeric. Jangan menandai outer comparison hasil RIGHT vs VARCHAR sebagai text-vs-number hanya berdasarkan tipe input. Namun operator `+` di dalam numeric_expression masih dapat mengonversi teks ke numeric; cek setiap level secara terpisah. Koreksi ini mencabut beberapa kandidat lama, tanpa menyatakan seluruh query bebas conversion. [Microsoft LEFT](https://learn.microsoft.com/en-us/sql/t-sql/functions/left-transact-sql).

REPLACE menghasilkan NVARCHAR bila ada argumen NVARCHAR, termasuk replacement. Unicode pada argumen kedua/ketiga tidak boleh diabaikan hanya karena argumen pertama VARCHAR. [Microsoft REPLACE](https://learn.microsoft.com/en-us/sql/t-sql/functions/replace-transact-sql).

CHAR(19) dibandingkan VARCHAR(19) bukan kasus yang sama dengan CHAR(19) dibandingkan BIGINT. Banyak mapping view/core memakai CHAR sementara staging/application memakai VARCHAR; kategorinya dipisahkan dari conversion numerik yang lebih berisiko.

Tetap periksa trailing spaces, padding, collation, panjang dan karakter. Jangan menyatakan family teks selalu aman atau selalu scan. Tipe string dan collation merupakan dimensi berbeda; merapikan type declaration saja belum tentu menyelesaikan collation conversion.

CFMAST menjadi salah satu contoh kecil: CFCIF CHAR(19) bertemu TreasuryCustomer_TM.CIFNo VARCHAR(19) berbeda dari join CFCIF CHAR dengan CIFNumber BIGINT. Pemisahan ini diterapkan pada semua tabel, bukan hanya CIF.

## P13 — Tipe lain dan ketidakcocokan yang memerlukan pembuktian

GUID vs teks, temporal-family yang berbeda, binary/text tertentu, atau user-defined type perlu aturan khusus. Precedence menentukan arah candidate reconciliation, tetapi tidak membuktikan implicit conversion diizinkan, sukses atas data yang ada, atau bebas efek storage/format.

Gunakan dokumentasi operator dan reproduction terisolasi jika diperlukan. Candidate OTHER_TYPE_DIFFERENCE bukan temuan pasti scan. Simpan error/ambiguity sebagai evidence, bukan membuat cast tebakan agar parser terlihat berhasil.

## Checklist saat memperbaiki satu pola

1. Pilih satu kolom dari rangkuman dan kumpulkan **seluruh SP terkait** pada dokumen tabelnya, termasuk cabang yang tampak serupa.
2. Tentukan domain canonical dan pemilik tipe di schema/interface. Jangan memakai nama variabel sebagai bukti tipe.
3. Pisahkan filter/join dari output/assignment. Arah konversinya dapat berbeda.
4. Uji data valid dan invalid: leading zero, spasi, empty, NULL, Unicode, overflow, tanggal, precision, serta row duplication.
5. Periksa plan sebelum/sesudah menggunakan parameter sempit, lebar, common, rare, dan NULL. Ukur logical reads/CPU/duration; jangan hanya mencari teks CONVERT_IMPLICIT.
6. Implementasikan konsisten pada semua procedure/interface yang memakai pola tersebut. Perbaikan satu query tidak menyelesaikan domain mismatch di seluruh workload.
7. Baru evaluasi index, write cost, dan migration baru sesuai repository. Audit ini tidak mengubah source atau menerapkan SQL.
