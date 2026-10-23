# Retail UAT.dbo.SecurityFeeNasabah_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 11 lokasi source, 4 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-operationcode"></a>
## OperationCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityFeeNasabah_TH.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009536 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParameterFeeNasabah:249](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParameterFeeNasabah.sql#L249) | JOIN_ON / P12 / COMPARISON | `sfn.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `sfn.OperationCode` → `VARCHAR` | `sfn.OperationCode = 'U'` |
| IC-007269 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityFee:355](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityFee.sql#L355) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securityfeenasabah_th.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.securityfeenasabah_th.OperationCode <- 'I'` |
| IC-007275 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityFee:479](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityFee.sql#L479) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securityfeenasabah_th.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `'U'` → `CHAR(1)` | `retail.dbo.securityfeenasabah_th.OperationCode <- 'U'` |

<a id="kolom-processcode"></a>
## ProcessCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityFeeNasabah_TH.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005997 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSecurityFeeAuthorization:32](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSecurityFeeAuthorization.sql#L32) | WHERE / P12 / COMPARISON | `ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `ProcessCode` → `VARCHAR` | `ProcessCode = 'P'` |
| IC-009533 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParameterFeeNasabah:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParameterFeeNasabah.sql#L169) | WHERE / P12 / IN_VALUE_TYPE | `sfn.ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `sfn.ProcessCode` → `VARCHAR` | `sfn.ProcessCode in ('A','R')` |
| IC-009535 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParameterFeeNasabah:248](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParameterFeeNasabah.sql#L248) | JOIN_ON / P12 / COMPARISON | `sfn.ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `sfn.ProcessCode` → `VARCHAR` | `sfn.ProcessCode = 'P'` |
| IC-009538 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParameterFeeNasabah:262](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParameterFeeNasabah.sql#L262) | JOIN_ON / P12 / COMPARISON | `sfn.ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `sfn.ProcessCode` → `VARCHAR` | `sfn.ProcessCode = 'P'` |
| IC-007270 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityFee:355](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityFee.sql#L355) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securityfeenasabah_th.ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.securityfeenasabah_th.ProcessCode <- 'P'` |
| IC-007276 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityFee:479](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityFee.sql#L479) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securityfeenasabah_th.ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.securityfeenasabah_th.ProcessCode <- 'P'` |
| IC-007280 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityFeeAuthorization:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityFeeAuthorization.sql#L52) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securityfeenasabah_th.ProcessCode` : `CHAR(1)` | `'R'` : `varchar` | `'R'` → `CHAR(1)` | `retail.dbo.securityfeenasabah_th.ProcessCode <- 'R'` |
| IC-007284 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityFeeAuthorization:219](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityFeeAuthorization.sql#L219) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securityfeenasabah_th.ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.securityfeenasabah_th.ProcessCode <- 'A'` |
