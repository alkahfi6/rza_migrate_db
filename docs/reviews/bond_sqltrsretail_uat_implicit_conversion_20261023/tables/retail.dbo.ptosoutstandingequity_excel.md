# Retail UAT.dbo.PTOSOutstandingEquity_Excel — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 24 lokasi source, 2 procedure, 6 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-harga-perolehan-nasabah"></a>
## Harga Perolehan Nasabah

**Tipe dasar:** `FLOAT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstandingEquity_Excel.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001237 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobOutstanding:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L24) | JOIN_ON / P11 / COMPARISON | `tmp.[Harga Perolehan Nasabah]` : `FLOAT` | `potr.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `potr.HargaPerolehanNasabah` → `FLOAT` | `tmp.[Harga Perolehan Nasabah] = potr.HargaPerolehanNasabah` |

<a id="kolom-kode-produk"></a>
## Kode Produk

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstandingEquity_Excel.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001236 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobOutstanding:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L23) | JOIN_ON / P02 / COMPARISON | `tmp.[Kode Produk]` : `NVARCHAR(255)` | `potr.KodeProduk` : `VARCHAR(20)` | `potr.KodeProduk` → `NVARCHAR` | `tmp.[Kode Produk] = potr.KodeProduk` |
| IC-001322 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobOutstanding:157](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L157) | JOIN_ON / P02 / COMPARISON | `tmp.[Kode Produk]` : `NVARCHAR(255)` | `tr.KodeProduk` : `VARCHAR(20)` | `tr.KodeProduk` → `NVARCHAR` | `tr.KodeProduk = tmp.[Kode Produk]` |
| IC-012279 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitLogSentMailResult:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L59) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[Kode Produk]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([Kode Produk], '')` |

<a id="kolom-nama-nasabah"></a>
## Nama Nasabah

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstandingEquity_Excel.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012274 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSubmitLogSentMailResult:55](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L55) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '') +'</td> <td class="table-data">' + isnull([Nomor Rekening RDN], '') + '</td> <td class="table-data">' + isnull([Nama Nasabah], '')` : `NVARCHAR (combined result; facets unknown)` | `'</td> <td class="table-data">'` : `varchar` | `'</td> <td class="table-data">'` → `NVARCHAR` | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '') +'</td> <td class="table-data">' + isnull([Nomor Rekening RDN], '') + '</td> <td class="table-data">' + isnull([Nama Nasabah], '') + '</td> <td class="table-data">'` |
| IC-012278 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitLogSentMailResult:58](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L58) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[Nama Nasabah]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([Nama Nasabah], '')` |

<a id="kolom-nomor-rekening-rdn"></a>
## Nomor Rekening RDN

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstandingEquity_Excel.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001242 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobOutstanding:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L34) | WHERE / P04 / COMPARISON | `isnull([Nomor Rekening RDN], '')` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR` | `isnull([Nomor Rekening RDN], '') = ''` |
| IC-001244 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.PTOSProsesDataJobOutstanding:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L34) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[Nomor Rekening RDN]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([Nomor Rekening RDN], '')` |
| IC-012277 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitLogSentMailResult:57](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L57) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[Nomor Rekening RDN]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([Nomor Rekening RDN], '')` |

<a id="kolom-result"></a>
## Result

**Tipe dasar:** `NVARCHAR(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstandingEquity_Excel.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P06 — LIKE dan representasi karakter:** LIKE pada key numerik memerlukan representasi karakter. Equality numeric yang tervalidasi cocok untuk lookup tepat; prefix/substring membutuhkan requirement pencarian tersendiri, bukan cast kolom sebagai solusi index.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001241 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L33) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Result` : `NVARCHAR(MAX)` | `'Failed - SID atau Nomor RDN kosong'` : `varchar` | `'Failed - SID atau Nomor RDN kosong'` → `NVARCHAR(MAX)` | `Result = 'Failed - SID atau Nomor RDN kosong'` |
| IC-001262 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L56) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tmp.Result` : `NVARCHAR(MAX)` | `'Failed - Customer tidak di temukan'` : `varchar` | `'Failed - Customer tidak di temukan'` → `NVARCHAR(MAX)` | `tmp.Result = 'Failed - Customer tidak di temukan'` |
| IC-001274 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobOutstanding:72](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L72) | WHERE / P06 / LIKE | `Result` : `NVARCHAR(MAX)` | `'Failed%'` : `varchar` | `'Failed%'` → `NVARCHAR` | `Result LIKE 'Failed%'` |
| IC-001276 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobOutstanding:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L80) | WHERE / P06 / LIKE | `Result` : `NVARCHAR(MAX)` | `'Failed%'` : `varchar` | `'Failed%'` → `NVARCHAR` | `Result LIKE 'Failed%'` |
| IC-001320 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:153](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L153) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tmp.Result` : `NVARCHAR(MAX)` | `'Failed - Insert to Table Master'` : `varchar` | `'Failed - Insert to Table Master'` → `NVARCHAR(MAX)` | `tmp.Result = 'Failed - Insert to Table Master'` |
| IC-012275 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSubmitLogSentMailResult:55](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L55) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '') +'</td> <td class="table-data">' + isnull([Nomor Rekening RDN], '') + '</td> <td class="table-data">' + isnull([Nama Nasabah], '') + '</td> <td class="table-data">' + isnull([Kode Produk], '') + '</td> <td class="table-data">' + cast([Nominal (Face Value) - Obligasi / Lembar Saham] as varchar) + '</td> <td class="table-data">' + [Result]` : `NVARCHAR (combined result; facets unknown)` | `'</td></tr>'` : `varchar` | `'</td></tr>'` → `NVARCHAR` | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '') +'</td> <td class="table-data">' + isnull([Nomor Rekening RDN], '') + '</td> <td class="table-data">' + isnull([Nama Nasabah], '') + '</td> <td class="table-data">' + isnull([Kode Produk], '') + '</td> <td class="table-data">' + cast([Nominal (Face Value) - Obligasi / Lembar Saham] as varchar) + '</td> <td class="table-data">' + [Result] + '</td></tr>'` |
| IC-012280 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSubmitLogSentMailResult:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L63) | WHERE / P04 / COMPARISON | `isnull(Result, '')` : `NVARCHAR(MAX)` | `''` : `varchar` | `''` → `NVARCHAR` | `isnull(Result, '') != ''` |
| IC-012281 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitLogSentMailResult:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L63) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `Result` : `NVARCHAR(MAX)` | `''` : `varchar` | `''` → `NVARCHAR(MAX)` | `isnull(Result, '')` |

<a id="kolom-sid"></a>
## SID

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstandingEquity_Excel.sql#L4). **Existing leading index:** PTOSOutstandingEquity_Excel_CN1.

**Frekuensi source:** 7 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.
- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001235 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobOutstanding:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L21) | JOIN_ON / P02 / COMPARISON | `tmp.[SID]` : `NVARCHAR(255)` | `potr.[SID]` : `VARCHAR(100)` | `potr.[SID]` → `NVARCHAR` | `tmp.[SID] = potr.[SID]` |
| IC-001243 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobOutstanding:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L34) | WHERE / P04 / COMPARISON | `isnull([SID], '')` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR` | `isnull([SID], '') = ''` |
| IC-001245 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.PTOSProsesDataJobOutstanding:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L34) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[SID]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([SID], '')` |
| IC-001263 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobOutstanding:58](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L58) | JOIN_ON / P02 / COMPARISON | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` : `VARCHAR(100)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001321 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobOutstanding:155](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L155) | JOIN_ON / P02 / COMPARISON | `tmp.[SID]` : `NVARCHAR(255)` | `tr.[SID]` : `VARCHAR(100)` | `tr.[SID]` → `NVARCHAR` | `tr.[SID] = tmp.[SID]` |
| IC-012273 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSubmitLogSentMailResult:55](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L55) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '')` : `NVARCHAR (combined result; facets unknown)` | `'</td> <td class="table-data">'` : `varchar` | `'</td> <td class="table-data">'` → `NVARCHAR` | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '') +'</td> <td class="table-data">'` |
| IC-012276 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitLogSentMailResult:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L56) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[SID]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([SID], '')` |
