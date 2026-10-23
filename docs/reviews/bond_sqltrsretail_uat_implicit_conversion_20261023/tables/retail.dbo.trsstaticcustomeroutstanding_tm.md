# Retail UAT.dbo.TRSStaticCustomerOutstanding_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 7 lokasi source, 2 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-outstanding"></a>
## Outstanding

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSStaticCustomerOutstanding_TM.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-011480 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptRekSuratBerharga:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptRekSuratBerharga.sql#L127) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `Outstanding` : `MONEY` | `@tmpData.Quantity` : `decimal(20,5)` | `Outstanding` → `decimal(20,5)` | `@tmpData.Quantity <- Outstanding` |
| IC-009595 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptCustomerPortfolio:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerPortfolio.sql#L106) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `Outstanding` : `MONEY` | `(SecMktPrice / 100)` : `DECIMAL (combined result; facets unknown)` | `Outstanding` → `DECIMAL` | `Outstanding * (SecMktPrice / 100)` |
| IC-009596 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptCustomerPortfolio:186](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerPortfolio.sql#L186) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `isnull(scotm.Outstanding,0)` : `MONEY` | `(isnull(SecMktPrice,0) / 100.0)` : `DECIMAL (combined result; facets unknown)` | `isnull(scotm.Outstanding,0)` → `DECIMAL` | `isnull(scotm.Outstanding,0) * (isnull(SecMktPrice,0) / 100.0)` |
| IC-009598 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptCustomerPortfolio:186](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerPortfolio.sql#L186) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `scotm.Outstanding` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(scotm.Outstanding,0)` |

<a id="kolom-period"></a>
## Period

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSStaticCustomerOutstanding_TM.sql#L2). **Existing leading index:** TRSStaticCustomerOutstanding_TM_CN1.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-011485 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptRekSuratBerharga:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptRekSuratBerharga.sql#L145) | WHERE / P04 / COMPARISON | `[Period]` : `DATETIME` | `convert(varchar, @pnPeriod)` : `varchar` | `convert(varchar, @pnPeriod)` → `DATETIME` | `[Period] = convert(varchar, @pnPeriod)` |

<a id="kolom-secmktprice"></a>
## SecMktPrice

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSStaticCustomerOutstanding_TM.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009594 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptCustomerPortfolio:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerPortfolio.sql#L106) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `SecMktPrice` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `SecMktPrice / 100` |
| IC-009595 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptCustomerPortfolio:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerPortfolio.sql#L106) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `(SecMktPrice / 100)` : `DECIMAL (combined result; facets unknown)` | `Outstanding` : `MONEY` | `Outstanding` → `DECIMAL` | `Outstanding * (SecMktPrice / 100)` |
| IC-009596 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptCustomerPortfolio:186](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerPortfolio.sql#L186) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `(isnull(SecMktPrice,0) / 100.0)` : `DECIMAL (combined result; facets unknown)` | `isnull(scotm.Outstanding,0)` : `MONEY` | `isnull(scotm.Outstanding,0)` → `DECIMAL` | `isnull(scotm.Outstanding,0) * (isnull(SecMktPrice,0) / 100.0)` |
| IC-009597 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptCustomerPortfolio:186](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerPortfolio.sql#L186) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `SecMktPrice` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(SecMktPrice,0)` |
