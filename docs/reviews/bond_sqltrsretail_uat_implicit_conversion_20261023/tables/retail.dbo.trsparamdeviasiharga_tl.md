# Retail UAT.dbo.TRSParamDeviasiHarga_TL — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 14 lokasi source, 3 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-actionflag"></a>
## ActionFlag

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamDeviasiHarga_TL.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 11 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005232 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprParamDeviasiHarga:6](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamDeviasiHarga.sql#L6) | WHERE / P12 / COMPARISON | `ActionFlag` : `CHAR(1)` | `ParamId` : `VARCHAR(20)` | `ActionFlag` → `VARCHAR` | `ParamId = ActionFlag` |
| IC-005233 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprParamDeviasiHarga:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamDeviasiHarga.sql#L17) | WHERE / P12 / IN_VALUE_TYPE | `ActionFlag` : `CHAR(1)` | `'O'` : `varchar` | `ActionFlag` → `VARCHAR` | `ActionFlag in ('O','D')` |
| IC-005234 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprParamDeviasiHarga:20](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamDeviasiHarga.sql#L20) | WHERE / P12 / COMPARISON | `ActionFlag` : `CHAR(1)` | `ParamId` : `VARCHAR(20)` | `ActionFlag` → `VARCHAR` | `ParamId = ActionFlag` |
| IC-005235 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprParamDeviasiHarga:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamDeviasiHarga.sql#L31) | WHERE / P12 / IN_VALUE_TYPE | `ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `ActionFlag` → `VARCHAR` | `ActionFlag in ('N','I')` |
| IC-006555 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprParamDeviasiHarga:11](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamDeviasiHarga.sql#L11) | WHERE / P12 / COMPARISON | `ActionFlag` : `CHAR(1)` | `'I'` : `varchar` | `ActionFlag` → `VARCHAR` | `ActionFlag = 'I'` |
| IC-006556 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprParamDeviasiHarga:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamDeviasiHarga.sql#L18) | WHERE / P12 / COMPARISON | `pl.ActionFlag` : `CHAR(1)` | `'D'` : `varchar` | `pl.ActionFlag` → `VARCHAR` | `pl.ActionFlag = 'D'` |
| IC-006557 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprParamDeviasiHarga:35](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamDeviasiHarga.sql#L35) | WHERE / P12 / COMPARISON | `pl.ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `pl.ActionFlag` → `VARCHAR` | `pl.ActionFlag = 'N'` |
| IC-009520 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParamWewenangDeviasi:10](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParamWewenangDeviasi.sql#L10) | CASE / P12 / COMPARISON | `tpdh.ActionFlag` : `CHAR(1)` | `'I'` : `varchar` | `tpdh.ActionFlag` → `VARCHAR` | `tpdh.ActionFlag = 'I'` |
| IC-009521 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParamWewenangDeviasi:11](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParamWewenangDeviasi.sql#L11) | CASE / P12 / COMPARISON | `tpdh.ActionFlag` : `CHAR(1)` | `'D'` : `varchar` | `tpdh.ActionFlag` → `VARCHAR` | `tpdh.ActionFlag = 'D'` |
| IC-009522 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParamWewenangDeviasi:15](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParamWewenangDeviasi.sql#L15) | CASE / P12 / COMPARISON | `tpdh.ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `tpdh.ActionFlag` → `VARCHAR` | `tpdh.ActionFlag = 'N'` |
| IC-009523 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParamWewenangDeviasi:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParamWewenangDeviasi.sql#L16) | CASE / P12 / COMPARISON | `tpdh.ActionFlag` : `CHAR(1)` | `'O'` : `varchar` | `tpdh.ActionFlag` → `VARCHAR` | `tpdh.ActionFlag = 'O'` |

<a id="kolom-trxtype"></a>
## TrxType

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamDeviasiHarga_TL.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009524 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParamWewenangDeviasi:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParamWewenangDeviasi.sql#L25) | EXPRESSION / P07 / CASE_BRANCH | `tpdh.TrxType` : `CHAR(1)` | `result type` : `VARCHAR (combined result; facets unknown)` | `tpdh.TrxType` → `VARCHAR` | `CASE WHEN tpdh.TrxType = 'S' THEN 'SELL' WHEN tpdh.TrxType = 'B' THEN 'BUY' ELSE tpdh.TrxType END` |
| IC-009525 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParamWewenangDeviasi:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParamWewenangDeviasi.sql#L26) | CASE / P12 / COMPARISON | `tpdh.TrxType` : `CHAR(1)` | `'S'` : `varchar` | `tpdh.TrxType` → `VARCHAR` | `tpdh.TrxType = 'S'` |
| IC-009526 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParamWewenangDeviasi:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParamWewenangDeviasi.sql#L27) | CASE / P12 / COMPARISON | `tpdh.TrxType` : `CHAR(1)` | `'B'` : `varchar` | `tpdh.TrxType` → `VARCHAR` | `tpdh.TrxType = 'B'` |
