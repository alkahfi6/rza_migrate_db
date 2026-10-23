# Retail UAT.dbo.CustodySfeeKupon_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 22 lokasi source, 4 procedure, 6 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 6 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-countretry"></a>
## CountRetry

**Tipe dasar:** `TINYINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CustodySfeeKupon_TR.sql#L30). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 10 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000256 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.CustodyUpdateEmailStatusCorporateAction:12](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyUpdateEmailStatusCorporateAction.sql#L12) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `CountRetry + 1` : `INT (combined result; facets unknown)` | `CountRetry` : `TINYINT` | `CountRetry + 1` → `TINYINT` | `CountRetry = CountRetry + 1` |
| IC-000260 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.CustodyUpdateEmailStatusCorporateAction:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyUpdateEmailStatusCorporateAction.sql#L24) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `CountRetry + 1` : `INT (combined result; facets unknown)` | `CountRetry` : `TINYINT` | `CountRetry + 1` → `TINYINT` | `CountRetry = CountRetry + 1` |
| IC-000131 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyPopulateEmailCorporateAction:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyPopulateEmailCorporateAction.sql#L34) | WHERE / P11 / COMPARISON | `CountRetry` : `TINYINT` | `3` : `int` | `CountRetry` → `INT` | `CountRetry < 3` |
| IC-000257 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyUpdateEmailStatusCorporateAction:12](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyUpdateEmailStatusCorporateAction.sql#L12) | DML_CONDITION / P05 / BINARY_Add | `CountRetry` : `TINYINT` | `1` : `int` | `CountRetry` → `INT` | `CountRetry + 1` |
| IC-000259 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyUpdateEmailStatusCorporateAction:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyUpdateEmailStatusCorporateAction.sql#L16) | WHERE / P11 / COMPARISON | `CountRetry` : `TINYINT` | `3` : `int` | `CountRetry` → `INT` | `CountRetry < 3` |
| IC-000261 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyUpdateEmailStatusCorporateAction:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyUpdateEmailStatusCorporateAction.sql#L24) | DML_CONDITION / P05 / BINARY_Add | `CountRetry` : `TINYINT` | `1` : `int` | `CountRetry` → `INT` | `CountRetry + 1` |
| IC-000263 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyUpdateEmailStatusCorporateAction:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyUpdateEmailStatusCorporateAction.sql#L27) | WHERE / P11 / COMPARISON | `CountRetry` : `TINYINT` | `3` : `int` | `CountRetry` → `INT` | `CountRetry < 3` |
| IC-006697 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprovedUploadSFeeKupon:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovedUploadSFeeKupon.sql#L53) | WHERE / P11 / COMPARISON | `dest.CountRetry` : `TINYINT` | `0` : `int` | `dest.CountRetry` → `INT` | `dest.CountRetry = 0` |
| IC-012855 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSValidateUploadSFeeKupon:197](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadSFeeKupon.sql#L197) | CASE / P11 / COMPARISON | `tc.CountRetry` : `TINYINT` | `0` : `int` | `tc.CountRetry` → `INT` | `tc.CountRetry = 0` |
| IC-012857 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSValidateUploadSFeeKupon:201](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadSFeeKupon.sql#L201) | CASE / P11 / COMPARISON | `tc.CountRetry` : `TINYINT` | `0` : `int` | `tc.CountRetry` → `INT` | `tc.CountRetry = 0` |
| IC-000256 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.CustodyUpdateEmailStatusCorporateAction:12](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyUpdateEmailStatusCorporateAction.sql#L12) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `CountRetry` : `TINYINT` | `CountRetry + 1` : `INT (combined result; facets unknown)` | `CountRetry + 1` → `TINYINT` | `CountRetry = CountRetry + 1` |
| IC-000260 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.CustodyUpdateEmailStatusCorporateAction:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyUpdateEmailStatusCorporateAction.sql#L24) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `CountRetry` : `TINYINT` | `CountRetry + 1` : `INT (combined result; facets unknown)` | `CountRetry + 1` → `TINYINT` | `CountRetry = CountRetry + 1` |

<a id="kolom-grossamount"></a>
## GrossAmount

**Tipe dasar:** `DECIMAL(20,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CustodySfeeKupon_TR.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT.

**Pola dan arah review:**

- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000125 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.CustodyPopulateEmailCorporateAction:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyPopulateEmailCorporateAction.sql#L21) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `GrossAmount` : `DECIMAL(20,2)` | `retail.dbo.fnformatstringthseparator.@cValue` : `varchar(100)` | `GrossAmount` → `varchar(100)` | `dbo.fnFormatStringThSeparator(GrossAmount, 2)` |

<a id="kolom-issent"></a>
## IsSent

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CustodySfeeKupon_TR.sql#L29). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000130 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyPopulateEmailCorporateAction:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyPopulateEmailCorporateAction.sql#L33) | WHERE / P11 / COMPARISON | `IsSent` : `BIT` | `0` : `int` | `IsSent` → `INT` | `IsSent = 0` |
| IC-000258 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyUpdateEmailStatusCorporateAction:15](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyUpdateEmailStatusCorporateAction.sql#L15) | WHERE / P11 / COMPARISON | `IsSent` : `BIT` | `0` : `int` | `IsSent` → `INT` | `IsSent = 0` |
| IC-000262 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyUpdateEmailStatusCorporateAction:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyUpdateEmailStatusCorporateAction.sql#L26) | WHERE / P11 / COMPARISON | `IsSent` : `BIT` | `0` : `int` | `IsSent` → `INT` | `IsSent = 0` |
| IC-006696 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprovedUploadSFeeKupon:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovedUploadSFeeKupon.sql#L52) | WHERE / P11 / COMPARISON | `dest.IsSent` : `BIT` | `0` : `int` | `dest.IsSent` → `INT` | `dest.IsSent = 0` |
| IC-012856 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSValidateUploadSFeeKupon:197](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadSFeeKupon.sql#L197) | CASE / P11 / COMPARISON | `tc.IsSent` : `BIT` | `0` : `int` | `tc.IsSent` → `INT` | `tc.IsSent = 0` |
| IC-012858 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSValidateUploadSFeeKupon:201](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadSFeeKupon.sql#L201) | CASE / P11 / COMPARISON | `tc.IsSent` : `BIT` | `0` : `int` | `tc.IsSent` → `INT` | `tc.IsSent = 0` |

<a id="kolom-netamount"></a>
## NetAmount

**Tipe dasar:** `DECIMAL(20,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CustodySfeeKupon_TR.sql#L26). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000129 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.CustodyPopulateEmailCorporateAction:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyPopulateEmailCorporateAction.sql#L24) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `ISNULL(NetAmount, 0)` : `DECIMAL(20,2)` | `retail.dbo.fnformatstringthseparator.@cValue` : `varchar(100)` | `ISNULL(NetAmount, 0)` → `varchar(100)` | `dbo.fnFormatStringThSeparator(ISNULL(NetAmount, 0), 2)` |
| IC-000128 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.CustodyPopulateEmailCorporateAction:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyPopulateEmailCorporateAction.sql#L24) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NetAmount` : `DECIMAL(20,2)` | `0` : `int` | `0` → `DECIMAL(20,2)` | `ISNULL(NetAmount, 0)` |

<a id="kolom-nominal"></a>
## Nominal

**Tipe dasar:** `DECIMAL(20,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CustodySfeeKupon_TR.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT.

**Pola dan arah review:**

- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000124 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.CustodyPopulateEmailCorporateAction:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyPopulateEmailCorporateAction.sql#L17) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `Nominal` : `DECIMAL(20,2)` | `retail.dbo.fnformatstringthseparator.@cValue` : `varchar(100)` | `Nominal` → `varchar(100)` | `dbo.fnFormatStringThSeparator(Nominal, 2)` |

<a id="kolom-taxoninterest"></a>
## TaxOnInterest

**Tipe dasar:** `DECIMAL(20,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CustodySfeeKupon_TR.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000127 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.CustodyPopulateEmailCorporateAction:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyPopulateEmailCorporateAction.sql#L22) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `ISNULL(TaxOnInterest, 0)` : `DECIMAL(20,2)` | `retail.dbo.fnformatstringthseparator.@cValue` : `varchar(100)` | `ISNULL(TaxOnInterest, 0)` → `varchar(100)` | `dbo.fnFormatStringThSeparator(ISNULL(TaxOnInterest, 0), 2)` |
| IC-000126 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.CustodyPopulateEmailCorporateAction:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyPopulateEmailCorporateAction.sql#L22) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TaxOnInterest` : `DECIMAL(20,2)` | `0` : `int` | `0` → `DECIMAL(20,2)` | `ISNULL(TaxOnInterest, 0)` |
