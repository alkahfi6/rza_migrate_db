# Retail UAT.dbo.TRSParamSpread_TL — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 18 lokasi source, 3 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-actionflag"></a>
## ActionFlag

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamSpread_TL.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 15 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005242 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprParamSpread:10](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamSpread.sql#L10) | WHERE / P12 / COMPARISON | `ActionFlag` : `CHAR(1)` | `ParamId` : `VARCHAR(20)` | `ActionFlag` → `VARCHAR` | `ParamId = ActionFlag` |
| IC-005244 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprParamSpread:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamSpread.sql#L27) | WHERE / P12 / IN_VALUE_TYPE | `ActionFlag` : `CHAR(1)` | `'O'` : `varchar` | `ActionFlag` → `VARCHAR` | `ActionFlag in ('O','D')` |
| IC-005245 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprParamSpread:35](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamSpread.sql#L35) | WHERE / P12 / COMPARISON | `ActionFlag` : `CHAR(1)` | `ParamId` : `VARCHAR(20)` | `ActionFlag` → `VARCHAR` | `ParamId = ActionFlag` |
| IC-005247 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprParamSpread:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamSpread.sql#L52) | WHERE / P12 / IN_VALUE_TYPE | `ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `ActionFlag` → `VARCHAR` | `ActionFlag in ('N','I')` |
| IC-006592 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprParamSpread:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L51) | WHERE / P12 / COMPARISON | `ActionFlag` : `CHAR(1)` | `'I'` : `varchar` | `ActionFlag` → `VARCHAR` | `ActionFlag = 'I'` |
| IC-006594 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprParamSpread:76](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L76) | WHERE / P12 / COMPARISON | `ActionFlag` : `CHAR(1)` | `'I'` : `varchar` | `ActionFlag` → `VARCHAR` | `ActionFlag = 'I'` |
| IC-006595 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprParamSpread:83](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L83) | WHERE / P12 / COMPARISON | `sl.ActionFlag` : `CHAR(1)` | `'D'` : `varchar` | `sl.ActionFlag` → `VARCHAR` | `sl.ActionFlag = 'D'` |
| IC-006596 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprParamSpread:91](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L91) | WHERE / P12 / COMPARISON | `sl.ActionFlag` : `CHAR(1)` | `'D'` : `varchar` | `sl.ActionFlag` → `VARCHAR` | `sl.ActionFlag = 'D'` |
| IC-006597 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprParamSpread:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L115) | WHERE / P12 / COMPARISON | `sl.ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `sl.ActionFlag` → `VARCHAR` | `sl.ActionFlag = 'N'` |
| IC-006598 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprParamSpread:134](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L134) | WHERE / P12 / COMPARISON | `sl.ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `sl.ActionFlag` → `VARCHAR` | `sl.ActionFlag = 'N'` |
| IC-006599 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprParamSpread:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L141) | WHERE / P12 / COMPARISON | `ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `ActionFlag` → `VARCHAR` | `ActionFlag = 'N'` |
| IC-009514 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParamSpread:9](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParamSpread.sql#L9) | CASE / P12 / COMPARISON | `tpsl.ActionFlag` : `CHAR(1)` | `'I'` : `varchar` | `tpsl.ActionFlag` → `VARCHAR` | `tpsl.ActionFlag = 'I'` |
| IC-009515 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParamSpread:10](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParamSpread.sql#L10) | CASE / P12 / COMPARISON | `tpsl.ActionFlag` : `CHAR(1)` | `'D'` : `varchar` | `tpsl.ActionFlag` → `VARCHAR` | `tpsl.ActionFlag = 'D'` |
| IC-009516 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParamSpread:14](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParamSpread.sql#L14) | CASE / P12 / COMPARISON | `tpsl.ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `tpsl.ActionFlag` → `VARCHAR` | `tpsl.ActionFlag = 'N'` |
| IC-009517 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParamSpread:15](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParamSpread.sql#L15) | CASE / P12 / COMPARISON | `tpsl.ActionFlag` : `CHAR(1)` | `'O'` : `varchar` | `tpsl.ActionFlag` → `VARCHAR` | `tpsl.ActionFlag = 'O'` |

<a id="kolom-indikasiharga"></a>
## IndikasiHarga

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamSpread_TL.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005243 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateApprParamSpread:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamSpread.sql#L25) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `IndikasiHarga` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(IndikasiHarga,0)` |
| IC-005246 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateApprParamSpread:50](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamSpread.sql#L50) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `IndikasiHarga` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(IndikasiHarga,0)` |
| IC-006593 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessApprParamSpread:73](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L73) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `IndikasiHarga` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(IndikasiHarga,0)` |
