# Retail UAT.dbo.TRSFileFormat_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 10 lokasi source, 1 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-excludeinfile"></a>
## ExcludeInFile

**Tipe dasar:** `TINYINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSFileFormat_TM.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009608 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptCustomerPortfolio:274](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerPortfolio.sql#L274) | WHERE / P11 / COMPARISON | `ExcludeInFile` : `TINYINT` | `0` : `int` | `ExcludeInFile` → `INT` | `ExcludeInFile = 0` |

<a id="kolom-fieldname"></a>
## FieldName

**Tipe dasar:** `VARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSFileFormat_TM.sql#L4). **Existing leading index:** TRSFileFormat_TM_NN1.

**Frekuensi source:** 9 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009599 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSRptCustomerPortfolio:268](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerPortfolio.sql#L268) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `FieldName` : `VARCHAR(255)` | `@cCommand + ' dbo.fnGetFixedLengthValue('` : `NVARCHAR (combined result; facets unknown)` | `FieldName` → `NVARCHAR` | `@cCommand + ' dbo.fnGetFixedLengthValue(' + FieldName` |
| IC-009600 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptCustomerPortfolio:268](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerPortfolio.sql#L268) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `@cCommand + ' dbo.fnGetFixedLengthValue(' + FieldName` : `NVARCHAR (combined result; facets unknown)` | `', '` : `varchar` | `', '` → `NVARCHAR` | `@cCommand + ' dbo.fnGetFixedLengthValue(' + FieldName + ', '` |
| IC-009601 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptCustomerPortfolio:268](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerPortfolio.sql#L268) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `@cCommand + ' dbo.fnGetFixedLengthValue(' + FieldName + ', '` : `NVARCHAR (combined result; facets unknown)` | `CAST(FormatType AS VARCHAR(1))` : `VARCHAR(1)` | `CAST(FormatType AS VARCHAR(1))` → `NVARCHAR` | `@cCommand + ' dbo.fnGetFixedLengthValue(' + FieldName + ', ' + CAST(FormatType AS VARCHAR(1))` |
| IC-009602 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptCustomerPortfolio:268](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerPortfolio.sql#L268) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `@cCommand + ' dbo.fnGetFixedLengthValue(' + FieldName + ', ' + CAST(FormatType AS VARCHAR(1))` : `NVARCHAR (combined result; facets unknown)` | `', '` : `varchar` | `', '` → `NVARCHAR` | `@cCommand + ' dbo.fnGetFixedLengthValue(' + FieldName + ', ' + CAST(FormatType AS VARCHAR(1)) + ', '` |
| IC-009603 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptCustomerPortfolio:268](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerPortfolio.sql#L268) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `@cCommand + ' dbo.fnGetFixedLengthValue(' + FieldName + ', ' + CAST(FormatType AS VARCHAR(1)) + ', '` : `NVARCHAR (combined result; facets unknown)` | `CAST(FieldLength AS VARCHAR(5))` : `VARCHAR(5)` | `CAST(FieldLength AS VARCHAR(5))` → `NVARCHAR` | `@cCommand + ' dbo.fnGetFixedLengthValue(' + FieldName + ', ' + CAST(FormatType AS VARCHAR(1)) + ', ' + CAST(FieldLength AS VARCHAR(5))` |
| IC-009604 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptCustomerPortfolio:268](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerPortfolio.sql#L268) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `@cCommand + ' dbo.fnGetFixedLengthValue(' + FieldName + ', ' + CAST(FormatType AS VARCHAR(1)) + ', ' + CAST(FieldLength AS VARCHAR(5))` : `NVARCHAR (combined result; facets unknown)` | `', '` : `varchar` | `', '` → `NVARCHAR` | `@cCommand + ' dbo.fnGetFixedLengthValue(' + FieldName + ', ' + CAST(FormatType AS VARCHAR(1)) + ', ' + CAST(FieldLength AS VARCHAR(5)) + ', '` |
| IC-009605 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptCustomerPortfolio:268](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerPortfolio.sql#L268) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `@cCommand + ' dbo.fnGetFixedLengthValue(' + FieldName + ', ' + CAST(FormatType AS VARCHAR(1)) + ', ' + CAST(FieldLength AS VARCHAR(5)) + ', '` : `NVARCHAR (combined result; facets unknown)` | `CAST(DecimalPlaces AS VARCHAR(5))` : `VARCHAR(5)` | `CAST(DecimalPlaces AS VARCHAR(5))` → `NVARCHAR` | `@cCommand + ' dbo.fnGetFixedLengthValue(' + FieldName + ', ' + CAST(FormatType AS VARCHAR(1)) + ', ' + CAST(FieldLength AS VARCHAR(5)) + ', ' + CAST(DecimalPlaces AS VARCHAR(5))` |
| IC-009606 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptCustomerPortfolio:268](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerPortfolio.sql#L268) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `@cCommand + ' dbo.fnGetFixedLengthValue(' + FieldName + ', ' + CAST(FormatType AS VARCHAR(1)) + ', ' + CAST(FieldLength AS VARCHAR(5)) + ', ' + CAST(DecimalPlaces AS VARCHAR(5))` : `NVARCHAR (combined result; facets unknown)` | `')'` : `varchar` | `')'` → `NVARCHAR` | `@cCommand + ' dbo.fnGetFixedLengthValue(' + FieldName + ', ' + CAST(FormatType AS VARCHAR(1)) + ', ' + CAST(FieldLength AS VARCHAR(5)) + ', ' + CAST(DecimalPlaces AS VARCHAR(5)) + ')'` |
| IC-009607 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptCustomerPortfolio:268](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerPortfolio.sql#L268) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `@cCommand + ' dbo.fnGetFixedLengthValue(' + FieldName + ', ' + CAST(FormatType AS VARCHAR(1)) + ', ' + CAST(FieldLength AS VARCHAR(5)) + ', ' + CAST(DecimalPlaces AS VARCHAR(5)) + ')'` : `NVARCHAR (combined result; facets unknown)` | `' +'` : `varchar` | `' +'` → `NVARCHAR` | `@cCommand + ' dbo.fnGetFixedLengthValue(' + FieldName + ', ' + CAST(FormatType AS VARCHAR(1)) + ', ' + CAST(FieldLength AS VARCHAR(5)) + ', ' + CAST(DecimalPlaces AS VARCHAR(5)) + ')' + ' +'` |
