# Retail UAT.dbo.SecurityNoBankCustody_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 11 lokasi source, 3 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-apprstatus"></a>
## ApprStatus

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityNoBankCustody_TH.sql#L19). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009552 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailSecurityNoBankCustody:100](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailSecurityNoBankCustody.sql#L100) | WHERE / P12 / IN_VALUE_TYPE | `ApprStatus` : `CHAR(1)` | `'A'` : `varchar` | `ApprStatus` → `VARCHAR` | `ApprStatus in ('A','R')` |
| IC-009553 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailSecurityNoBankCustody:144](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailSecurityNoBankCustody.sql#L144) | JOIN_ON / P12 / COMPARISON | `snbc.ApprStatus` : `CHAR(1)` | `'P'` : `varchar` | `snbc.ApprStatus` → `VARCHAR` | `snbc.ApprStatus = 'P'` |
| IC-009554 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailSecurityNoBankCustody:154](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailSecurityNoBankCustody.sql#L154) | JOIN_ON / P12 / COMPARISON | `snbc.ApprStatus` : `CHAR(1)` | `'P'` : `varchar` | `snbc.ApprStatus` → `VARCHAR` | `snbc.ApprStatus = 'P'` |
| IC-013638 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_ProcessSecurityNoBankCustody:116](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ProcessSecurityNoBankCustody.sql#L116) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitynobankcustody_th.ApprStatus` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.securitynobankcustody_th.ApprStatus <- 'P'` |
| IC-013640 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_ProcessSecurityNoBankCustody:206](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ProcessSecurityNoBankCustody.sql#L206) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitynobankcustody_th.ApprStatus` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.securitynobankcustody_th.ApprStatus <- 'P'` |
| IC-013642 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_ProcessSecurityNoBankCustody:295](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ProcessSecurityNoBankCustody.sql#L295) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitynobankcustody_th.ApprStatus` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.securitynobankcustody_th.ApprStatus <- 'P'` |

<a id="kolom-detailupdate"></a>
## DetailUpdate

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityNoBankCustody_TH.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013639 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_ProcessSecurityNoBankCustody:116](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ProcessSecurityNoBankCustody.sql#L116) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitynobankcustody_th.DetailUpdate` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.securitynobankcustody_th.DetailUpdate <- 'N'` |
| IC-013641 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_ProcessSecurityNoBankCustody:206](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ProcessSecurityNoBankCustody.sql#L206) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitynobankcustody_th.DetailUpdate` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.securitynobankcustody_th.DetailUpdate <- 'O'` |
| IC-013643 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_ProcessSecurityNoBankCustody:295](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ProcessSecurityNoBankCustody.sql#L295) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitynobankcustody_th.DetailUpdate` : `CHAR(1)` | `'D'` : `varchar` | `'D'` → `CHAR(1)` | `retail.dbo.securitynobankcustody_th.DetailUpdate <- 'D'` |
| IC-013644 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_ProcessSecurityNoBankCustodyAuthorization:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ProcessSecurityNoBankCustodyAuthorization.sql#L106) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitynobankcustody_th.DetailUpdate` : `CHAR(1)` | `case when snbcTemp.Status = 2 then 'D' else 'N' end` : `VARCHAR (combined result; facets unknown)` | `case when snbcTemp.Status = 2 then 'D' else 'N' end` → `CHAR(1)` | `retail.dbo.securitynobankcustody_th.DetailUpdate <- case when snbcTemp.Status = 2 then 'D' else 'N' end` |
| IC-013645 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_ProcessSecurityNoBankCustodyAuthorization:172](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ProcessSecurityNoBankCustodyAuthorization.sql#L172) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitynobankcustody_th.DetailUpdate` : `CHAR(1)` | `case when snbcTemp.Status = 2 then 'D' else 'N' end` : `VARCHAR (combined result; facets unknown)` | `case when snbcTemp.Status = 2 then 'D' else 'N' end` → `CHAR(1)` | `retail.dbo.securitynobankcustody_th.DetailUpdate <- case when snbcTemp.Status = 2 then 'D' else 'N' end` |
