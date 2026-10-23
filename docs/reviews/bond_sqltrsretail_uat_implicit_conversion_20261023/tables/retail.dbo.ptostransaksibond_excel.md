# Retail UAT.dbo.PTOSTransaksiBond_Excel — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 25 lokasi source, 2 procedure, 7 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-deal-id"></a>
## Deal ID

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksiBond_Excel.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001566 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobTransaksiBond:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L28) | JOIN_ON / P02 / COMPARISON | `tmp.[Deal ID]` : `NVARCHAR(255)` | `potr.DealId` : `VARCHAR(20)` | `potr.DealId` → `NVARCHAR` | `tmp.[Deal ID] = potr.DealId` |
| IC-001687 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobTransaksiBond:167](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L167) | JOIN_ON / P02 / COMPARISON | `tmp.[Deal ID]` : `NVARCHAR(255)` | `tr.DealId` : `VARCHAR(20)` | `tr.DealId` → `NVARCHAR` | `tr.DealId = tmp.[Deal ID]` |

<a id="kolom-jenis-transaksi-pe-beli-pe-jual"></a>
## Jenis Transaksi (PE Beli/PE Jual)

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksiBond_Excel.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001565 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobTransaksiBond:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L23) | CASE / P04 / COMPARISON | `[Jenis Transaksi (PE Beli/PE Jual)]` : `NVARCHAR(255)` | `'Beli'` : `varchar` | `'Beli'` → `NVARCHAR` | `[Jenis Transaksi (PE Beli/PE Jual)] = 'Beli'` |
| IC-012302 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitLogSentMailResult:95](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L95) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[Jenis Transaksi (PE Beli/PE Jual)]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([Jenis Transaksi (PE Beli/PE Jual)], '')` |

<a id="kolom-kode-produk"></a>
## Kode Produk

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksiBond_Excel.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012296 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSubmitLogSentMailResult:91](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L91) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '') +'</td> <td class="table-data">' + isnull([Nomor Rekening RDN], '') + '</td> <td class="table-data">' + isnull([Nama Nasabah], '') + '</td> <td class="table-data">' + isnull([Jenis Transaksi (PE Beli/PE Jual)], '') + '</td> <td class="table-data">' + isnull([Kode Produk], '')` : `NVARCHAR (combined result; facets unknown)` | `'</td> <td class="table-data">'` : `varchar` | `'</td> <td class="table-data">'` → `NVARCHAR` | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '') +'</td> <td class="table-data">' + isnull([Nomor Rekening RDN], '') + '</td> <td class="table-data">' + isnull([Nama Nasabah], '') + '</td> <td class="table-data">' + isnull([Jenis Transaksi (PE Beli/PE Jual)], '') + '</td> <td class="table-data">' + isnull([Kode Produk], '') + '</td> <td class="table-data">'` |
| IC-012297 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSubmitLogSentMailResult:91](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L91) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '') +'</td> <td class="table-data">' + isnull([Nomor Rekening RDN], '') + '</td> <td class="table-data">' + isnull([Nama Nasabah], '') + '</td> <td class="table-data">' + isnull([Jenis Transaksi (PE Beli/PE Jual)], '') + '</td> <td class="table-data">' + isnull([Kode Produk], '') + '</td> <td class="table-data">'` : `NVARCHAR (combined result; facets unknown)` | `cast([Nominal (Face Value) - Obligasi / Lembar Saham] as varchar)` : `varchar` | `cast([Nominal (Face Value) - Obligasi / Lembar Saham] as varchar)` → `NVARCHAR` | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '') +'</td> <td class="table-data">' + isnull([Nomor Rekening RDN], '') + '</td> <td class="table-data">' + isnull([Nama Nasabah], '') + '</td> <td class="table-data">' + isnull([Jenis Transaksi (PE Beli/PE Jual)], '') + '</td> <td class="table-data">' + isnull([Kode Produk], '') + '</td> <td class="table-data">' + cast([Nominal (Face Value) - Obligasi / Lembar Saham] as varchar)` |
| IC-012298 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSubmitLogSentMailResult:91](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L91) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '') +'</td> <td class="table-data">' + isnull([Nomor Rekening RDN], '') + '</td> <td class="table-data">' + isnull([Nama Nasabah], '') + '</td> <td class="table-data">' + isnull([Jenis Transaksi (PE Beli/PE Jual)], '') + '</td> <td class="table-data">' + isnull([Kode Produk], '') + '</td> <td class="table-data">' + cast([Nominal (Face Value) - Obligasi / Lembar Saham] as varchar)` : `NVARCHAR (combined result; facets unknown)` | `'</td> <td class="table-data">'` : `varchar` | `'</td> <td class="table-data">'` → `NVARCHAR` | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '') +'</td> <td class="table-data">' + isnull([Nomor Rekening RDN], '') + '</td> <td class="table-data">' + isnull([Nama Nasabah], '') + '</td> <td class="table-data">' + isnull([Jenis Transaksi (PE Beli/PE Jual)], '') + '</td> <td class="table-data">' + isnull([Kode Produk], '') + '</td> <td class="table-data">' + cast([Nominal (Face Value) - Obligasi / Lembar Saham] as varchar) + '</td> <td class="table-data">'` |
| IC-012303 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitLogSentMailResult:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L96) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[Kode Produk]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([Kode Produk], '')` |

<a id="kolom-nama-nasabah"></a>
## Nama Nasabah

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksiBond_Excel.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012295 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSubmitLogSentMailResult:91](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L91) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '') +'</td> <td class="table-data">' + isnull([Nomor Rekening RDN], '') + '</td> <td class="table-data">' + isnull([Nama Nasabah], '')` : `NVARCHAR (combined result; facets unknown)` | `'</td> <td class="table-data">'` : `varchar` | `'</td> <td class="table-data">'` → `NVARCHAR` | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '') +'</td> <td class="table-data">' + isnull([Nomor Rekening RDN], '') + '</td> <td class="table-data">' + isnull([Nama Nasabah], '') + '</td> <td class="table-data">'` |
| IC-012301 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitLogSentMailResult:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L94) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[Nama Nasabah]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([Nama Nasabah], '')` |

<a id="kolom-nomor-rekening-rdn"></a>
## Nomor Rekening RDN

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksiBond_Excel.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001571 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobTransaksiBond:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L41) | WHERE / P04 / COMPARISON | `isnull([Nomor Rekening RDN], '')` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR` | `isnull([Nomor Rekening RDN], '') = ''` |
| IC-001573 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.PTOSProsesDataJobTransaksiBond:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L41) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[Nomor Rekening RDN]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([Nomor Rekening RDN], '')` |
| IC-012300 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitLogSentMailResult:93](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L93) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[Nomor Rekening RDN]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([Nomor Rekening RDN], '')` |

<a id="kolom-result"></a>
## Result

**Tipe dasar:** `NVARCHAR(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksiBond_Excel.sql#L24). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P06 — LIKE dan representasi karakter:** LIKE pada key numerik memerlukan representasi karakter. Equality numeric yang tervalidasi cocok untuk lookup tepat; prefix/substring membutuhkan requirement pencarian tersendiri, bukan cast kolom sebagai solusi index.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001570 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:40](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L40) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Result` : `NVARCHAR(MAX)` | `'Failed - SID atau Nomor RDN kosong'` : `varchar` | `'Failed - SID atau Nomor RDN kosong'` → `NVARCHAR(MAX)` | `Result = 'Failed - SID atau Nomor RDN kosong'` |
| IC-001599 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L64) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tmp.Result` : `NVARCHAR(MAX)` | `'Failed - Customer tidak di temukan'` : `varchar` | `'Failed - Customer tidak di temukan'` → `NVARCHAR(MAX)` | `tmp.Result = 'Failed - Customer tidak di temukan'` |
| IC-001619 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobTransaksiBond:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L81) | WHERE / P06 / LIKE | `Result` : `NVARCHAR(MAX)` | `'Failed%'` : `varchar` | `'Failed%'` → `NVARCHAR` | `Result LIKE 'Failed%'` |
| IC-001621 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobTransaksiBond:88](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L88) | WHERE / P06 / LIKE | `Result` : `NVARCHAR(MAX)` | `'Failed%'` : `varchar` | `'Failed%'` → `NVARCHAR` | `Result LIKE 'Failed%'` |
| IC-001686 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L165) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tmp.Result` : `NVARCHAR(MAX)` | `'Failed - Insert to Table Master'` : `varchar` | `'Failed - Insert to Table Master'` → `NVARCHAR(MAX)` | `tmp.Result = 'Failed - Insert to Table Master'` |
| IC-012304 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSubmitLogSentMailResult:100](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L100) | WHERE / P04 / COMPARISON | `isnull(Result, '')` : `NVARCHAR(MAX)` | `''` : `varchar` | `''` → `NVARCHAR` | `isnull(Result, '') != ''` |
| IC-012305 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitLogSentMailResult:100](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L100) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `Result` : `NVARCHAR(MAX)` | `''` : `varchar` | `''` → `NVARCHAR(MAX)` | `isnull(Result, '')` |

<a id="kolom-sid"></a>
## SID

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksiBond_Excel.sql#L4). **Existing leading index:** PTOSTransaksiBond_Excel_CN1.

**Frekuensi source:** 5 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.
- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001572 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobTransaksiBond:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L41) | WHERE / P04 / COMPARISON | `isnull([SID], '')` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR` | `isnull([SID], '') = ''` |
| IC-001574 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.PTOSProsesDataJobTransaksiBond:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L41) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[SID]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([SID], '')` |
| IC-001600 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobTransaksiBond:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L66) | JOIN_ON / P02 / COMPARISON | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` : `VARCHAR(100)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-012294 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSubmitLogSentMailResult:91](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L91) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '')` : `NVARCHAR (combined result; facets unknown)` | `'</td> <td class="table-data">'` : `varchar` | `'</td> <td class="table-data">'` → `NVARCHAR` | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '') +'</td> <td class="table-data">'` |
| IC-012299 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitLogSentMailResult:92](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L92) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[SID]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([SID], '')` |
