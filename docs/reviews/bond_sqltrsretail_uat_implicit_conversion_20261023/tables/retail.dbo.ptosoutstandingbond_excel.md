# Retail UAT.dbo.PTOSOutstandingBond_Excel — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 26 lokasi source, 2 procedure, 6 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-harga-perolehan-nasabah"></a>
## Harga Perolehan Nasabah

**Tipe dasar:** `FLOAT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstandingBond_Excel.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001331 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobOutstandingBond:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L26) | JOIN_ON / P11 / COMPARISON | `tmp.[Harga Perolehan Nasabah]` : `FLOAT` | `potr.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `potr.HargaPerolehanNasabah` → `FLOAT` | `tmp.[Harga Perolehan Nasabah] = potr.HargaPerolehanNasabah` |

<a id="kolom-kode-produk"></a>
## Kode Produk

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstandingBond_Excel.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001330 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobOutstandingBond:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L25) | JOIN_ON / P02 / COMPARISON | `tmp.[Kode Produk]` : `NVARCHAR(255)` | `potr.KodeProduk` : `VARCHAR(20)` | `potr.KodeProduk` → `NVARCHAR` | `tmp.[Kode Produk] = potr.KodeProduk` |
| IC-001428 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobOutstandingBond:197](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L197) | JOIN_ON / P02 / COMPARISON | `tmp.[Kode Produk]` : `NVARCHAR(255)` | `tr.KodeProduk` : `VARCHAR(20)` | `tr.KodeProduk` → `NVARCHAR` | `tr.KodeProduk = tmp.[Kode Produk]` |
| IC-012270 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitLogSentMailResult:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L46) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[Kode Produk]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([Kode Produk], '')` |

<a id="kolom-nama-nasabah"></a>
## Nama Nasabah

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstandingBond_Excel.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012265 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSubmitLogSentMailResult:42](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L42) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '') +'</td> <td class="table-data">' + isnull([Nomor Rekening RDN], '') + '</td> <td class="table-data">' + isnull([Nama Nasabah], '')` : `NVARCHAR (combined result; facets unknown)` | `'</td> <td class="table-data">'` : `varchar` | `'</td> <td class="table-data">'` → `NVARCHAR` | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '') +'</td> <td class="table-data">' + isnull([Nomor Rekening RDN], '') + '</td> <td class="table-data">' + isnull([Nama Nasabah], '') + '</td> <td class="table-data">'` |
| IC-012269 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitLogSentMailResult:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L45) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[Nama Nasabah]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([Nama Nasabah], '')` |

<a id="kolom-nomor-rekening-rdn"></a>
## Nomor Rekening RDN

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstandingBond_Excel.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.
- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001336 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobOutstandingBond:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L37) | WHERE / P04 / COMPARISON | `isnull([Nomor Rekening RDN], '')` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR` | `isnull([Nomor Rekening RDN], '') = ''` |
| IC-001338 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.PTOSProsesDataJobOutstandingBond:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L37) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[Nomor Rekening RDN]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([Nomor Rekening RDN], '')` |
| IC-001427 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobOutstandingBond:196](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L196) | JOIN_ON / P02 / COMPARISON | `tmp.[Nomor Rekening RDN]` : `NVARCHAR(255)` | `tr.NoRekInvestorRDN` : `VARCHAR(20)` | `tr.NoRekInvestorRDN` → `NVARCHAR` | `tr.NoRekInvestorRDN = tmp.[Nomor Rekening RDN]` |
| IC-012268 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitLogSentMailResult:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L44) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[Nomor Rekening RDN]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([Nomor Rekening RDN], '')` |

<a id="kolom-result"></a>
## Result

**Tipe dasar:** `NVARCHAR(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstandingBond_Excel.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P06 — LIKE dan representasi karakter:** LIKE pada key numerik memerlukan representasi karakter. Equality numeric yang tervalidasi cocok untuk lookup tepat; prefix/substring membutuhkan requirement pencarian tersendiri, bukan cast kolom sebagai solusi index.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001335 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L36) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Result` : `NVARCHAR(MAX)` | `'Failed - SID atau Nomor RDN kosong'` : `varchar` | `'Failed - SID atau Nomor RDN kosong'` → `NVARCHAR(MAX)` | `Result = 'Failed - SID atau Nomor RDN kosong'` |
| IC-001356 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L59) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tmp.Result` : `NVARCHAR(MAX)` | `'Failed - Customer tidak ditemukan'` : `varchar` | `'Failed - Customer tidak ditemukan'` → `NVARCHAR(MAX)` | `tmp.Result = 'Failed - Customer tidak ditemukan'` |
| IC-001369 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L86) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ex.Result` : `NVARCHAR(MAX)` | `'Failed - Kode Produk tidak ditemukan'` : `varchar` | `'Failed - Kode Produk tidak ditemukan'` → `NVARCHAR(MAX)` | `ex.Result = 'Failed - Kode Produk tidak ditemukan'` |
| IC-001379 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobOutstandingBond:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L104) | WHERE / P06 / LIKE | `Result` : `NVARCHAR(MAX)` | `'Failed%'` : `varchar` | `'Failed%'` → `NVARCHAR` | `Result LIKE 'Failed%'` |
| IC-001381 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobOutstandingBond:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L111) | WHERE / P06 / LIKE | `Result` : `NVARCHAR(MAX)` | `'Failed%'` : `varchar` | `'Failed%'` → `NVARCHAR` | `Result LIKE 'Failed%'` |
| IC-001425 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L193) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tmp.Result` : `NVARCHAR(MAX)` | `'Failed - Insert to Table Master'` : `varchar` | `'Failed - Insert to Table Master'` → `NVARCHAR(MAX)` | `tmp.Result = 'Failed - Insert to Table Master'` |
| IC-012266 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSubmitLogSentMailResult:42](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L42) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '') +'</td> <td class="table-data">' + isnull([Nomor Rekening RDN], '') + '</td> <td class="table-data">' + isnull([Nama Nasabah], '') + '</td> <td class="table-data">' + isnull([Kode Produk], '') + '</td> <td class="table-data">' + cast([Nominal (Face Value) - Obligasi / Lembar Saham] as varchar) + '</td> <td class="table-data">' + [Result]` : `NVARCHAR (combined result; facets unknown)` | `'</td></tr>'` : `varchar` | `'</td></tr>'` → `NVARCHAR` | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '') +'</td> <td class="table-data">' + isnull([Nomor Rekening RDN], '') + '</td> <td class="table-data">' + isnull([Nama Nasabah], '') + '</td> <td class="table-data">' + isnull([Kode Produk], '') + '</td> <td class="table-data">' + cast([Nominal (Face Value) - Obligasi / Lembar Saham] as varchar) + '</td> <td class="table-data">' + [Result] + '</td></tr>'` |
| IC-012271 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSubmitLogSentMailResult:50](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L50) | WHERE / P04 / COMPARISON | `isnull(Result, '')` : `NVARCHAR(MAX)` | `''` : `varchar` | `''` → `NVARCHAR` | `isnull(Result, '') != ''` |
| IC-012272 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitLogSentMailResult:50](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L50) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `Result` : `NVARCHAR(MAX)` | `''` : `varchar` | `''` → `NVARCHAR(MAX)` | `isnull(Result, '')` |

<a id="kolom-sid"></a>
## SID

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstandingBond_Excel.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.
- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001329 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobOutstandingBond:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L23) | JOIN_ON / P02 / COMPARISON | `tmp.[SID]` : `NVARCHAR(255)` | `potr.[SID]` : `VARCHAR(100)` | `potr.[SID]` → `NVARCHAR` | `tmp.[SID] = potr.[SID]` |
| IC-001337 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobOutstandingBond:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L37) | WHERE / P04 / COMPARISON | `isnull([SID], '')` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR` | `isnull([SID], '') = ''` |
| IC-001339 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.PTOSProsesDataJobOutstandingBond:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L37) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[SID]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([SID], '')` |
| IC-001357 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobOutstandingBond:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L61) | JOIN_ON / P02 / COMPARISON | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` : `VARCHAR(100)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001426 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobOutstandingBond:195](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L195) | JOIN_ON / P02 / COMPARISON | `tmp.[SID]` : `NVARCHAR(255)` | `tr.[SID]` : `VARCHAR(100)` | `tr.[SID]` → `NVARCHAR` | `tr.[SID] = tmp.[SID]` |
| IC-012264 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSubmitLogSentMailResult:42](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L42) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '')` : `NVARCHAR (combined result; facets unknown)` | `'</td> <td class="table-data">'` : `varchar` | `'</td> <td class="table-data">'` → `NVARCHAR` | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '') +'</td> <td class="table-data">'` |
| IC-012267 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitLogSentMailResult:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L43) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[SID]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([SID], '')` |
