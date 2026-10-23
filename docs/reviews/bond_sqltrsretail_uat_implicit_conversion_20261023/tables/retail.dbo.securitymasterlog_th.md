# Retail UAT.dbo.SecurityMasterLog_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 23 lokasi source, 8 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-operationcode"></a>
## OperationCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterLog_TH.sql#L51). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009546 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParameterORI:480](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParameterORI.sql#L480) | JOIN_ON / P12 / COMPARISON | `smlth.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `smlth.OperationCode` → `VARCHAR` | `smlth.OperationCode = 'U'` |
| IC-010808 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORIDetail:199](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L199) | WHERE / P12 / COMPARISON | `OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `OperationCode` → `VARCHAR` | `OperationCode = 'U'` |
| IC-010942 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnesty:196](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L196) | WHERE / P12 / COMPARISON | `OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `OperationCode` → `VARCHAR` | `OperationCode = 'U'` |
| IC-011138 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:166](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L166) | WHERE / P12 / COMPARISON | `OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `OperationCode` → `VARCHAR` | `OperationCode = 'U'` |
| IC-011348 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnesty_Backdated:151](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty_Backdated.sql#L151) | WHERE / P12 / COMPARISON | `OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `OperationCode` → `VARCHAR` | `OperationCode = 'U'` |
| IC-007293 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityMaster:746](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMaster.sql#L746) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitymasterlog_th.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.securitymasterlog_th.OperationCode <- 'I'` |
| IC-007317 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityMaster:1364](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMaster.sql#L1364) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitymasterlog_th.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `'U'` → `CHAR(1)` | `retail.dbo.securitymasterlog_th.OperationCode <- 'U'` |
| IC-007321 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityMaster:1922](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMaster.sql#L1922) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitymasterlog_th.OperationCode` : `CHAR(1)` | `'D'` : `varchar` | `'D'` → `CHAR(1)` | `retail.dbo.securitymasterlog_th.OperationCode <- 'D'` |

<a id="kolom-processcode"></a>
## ProcessCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterLog_TH.sql#L53). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 15 lokasi / 8 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006012 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSecurityMasterAuthorization:196](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSecurityMasterAuthorization.sql#L196) | WHERE / P12 / COMPARISON | `ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `ProcessCode` → `VARCHAR` | `ProcessCode = 'P'` |
| IC-009544 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParameterORI:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParameterORI.sql#L323) | WHERE / P12 / IN_VALUE_TYPE | `ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `ProcessCode` → `VARCHAR` | `ProcessCode in ('A','R')` |
| IC-009545 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParameterORI:479](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParameterORI.sql#L479) | JOIN_ON / P12 / COMPARISON | `smlth.ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `smlth.ProcessCode` → `VARCHAR` | `smlth.ProcessCode = 'P'` |
| IC-009547 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParameterORI:489](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParameterORI.sql#L489) | JOIN_ON / P12 / COMPARISON | `tsth.ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `tsth.ProcessCode` → `VARCHAR` | `tsth.ProcessCode = 'P'` |
| IC-010809 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORIDetail:199](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L199) | WHERE / P12 / COMPARISON | `ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `ProcessCode` → `VARCHAR` | `ProcessCode = 'A'` |
| IC-010943 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnesty:196](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L196) | WHERE / P12 / COMPARISON | `ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `ProcessCode` → `VARCHAR` | `ProcessCode = 'A'` |
| IC-011139 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:166](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L166) | WHERE / P12 / COMPARISON | `ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `ProcessCode` → `VARCHAR` | `ProcessCode = 'A'` |
| IC-011349 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnesty_Backdated:151](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty_Backdated.sql#L151) | WHERE / P12 / COMPARISON | `ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `ProcessCode` → `VARCHAR` | `ProcessCode = 'A'` |
| IC-007294 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityMaster:746](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMaster.sql#L746) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitymasterlog_th.ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.securitymasterlog_th.ProcessCode <- 'P'` |
| IC-007318 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityMaster:1364](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMaster.sql#L1364) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitymasterlog_th.ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.securitymasterlog_th.ProcessCode <- 'P'` |
| IC-007322 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityMaster:1922](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMaster.sql#L1922) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitymasterlog_th.ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.securitymasterlog_th.ProcessCode <- 'P'` |
| IC-007323 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:68](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L68) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitymasterlog_th.ProcessCode` : `CHAR(1)` | `'R'` : `varchar` | `'R'` → `CHAR(1)` | `retail.dbo.securitymasterlog_th.ProcessCode <- 'R'` |
| IC-007344 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:767](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L767) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitymasterlog_th.ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.securitymasterlog_th.ProcessCode <- 'A'` |
| IC-007359 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1346](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1346) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitymasterlog_th.ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.securitymasterlog_th.ProcessCode <- 'A'` |
| IC-007366 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1659](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1659) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitymasterlog_th.ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.securitymasterlog_th.ProcessCode <- 'A'` |
