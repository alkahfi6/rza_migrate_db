# Retail UAT.dbo.TRSSeller_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 32 lokasi source, 5 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-operationcode"></a>
## OperationCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSeller_TH.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 12 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008411 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L34) | WHERE / P12 / COMPARISON | `tsrth.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `tsrth.OperationCode` → `VARCHAR` | `tsrth.OperationCode = 'I'` |
| IC-008416 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L59) | WHERE / P12 / COMPARISON | `tsrth.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `tsrth.OperationCode` → `VARCHAR` | `tsrth.OperationCode = 'I'` |
| IC-008421 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:84](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L84) | WHERE / P12 / COMPARISON | `tsrth.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `tsrth.OperationCode` → `VARCHAR` | `tsrth.OperationCode = 'I'` |
| IC-008426 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L110) | WHERE / P12 / COMPARISON | `tsrth.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `tsrth.OperationCode` → `VARCHAR` | `tsrth.OperationCode = 'I'` |
| IC-008431 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:136](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L136) | WHERE / P12 / COMPARISON | `tsrth.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `tsrth.OperationCode` → `VARCHAR` | `tsrth.OperationCode = 'I'` |
| IC-008436 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:206](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L206) | WHERE / P12 / COMPARISON | `tsrth.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `tsrth.OperationCode` → `VARCHAR` | `tsrth.OperationCode = 'I'` |
| IC-008441 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L225) | WHERE / P12 / COMPARISON | `tsrth.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `tsrth.OperationCode` → `VARCHAR` | `tsrth.OperationCode = 'I'` |
| IC-008446 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:244](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L244) | WHERE / P12 / COMPARISON | `tsrth.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `tsrth.OperationCode` → `VARCHAR` | `tsrth.OperationCode = 'I'` |
| IC-008451 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:264](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L264) | WHERE / P12 / COMPARISON | `tsrth.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `tsrth.OperationCode` → `VARCHAR` | `tsrth.OperationCode = 'I'` |
| IC-008456 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:284](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L284) | WHERE / P12 / COMPARISON | `tsrth.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `tsrth.OperationCode` → `VARCHAR` | `tsrth.OperationCode = 'I'` |
| IC-009550 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParameterSeller:140](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParameterSeller.sql#L140) | JOIN_ON / P12 / COMPARISON | `tsth.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `tsth.OperationCode` → `VARCHAR` | `tsth.OperationCode = 'U'` |
| IC-005208 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSParamSellerSync:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamSellerSync.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsseller_th.OperationCode` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.trsseller_th.OperationCode <- 'A'` |

<a id="kolom-processcode"></a>
## ProcessCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSeller_TH.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 20 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006019 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSeller:99](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSeller.sql#L99) | WHERE / P12 / COMPARISON | `ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `ProcessCode` → `VARCHAR` | `ProcessCode = 'P'` |
| IC-007385 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSeller:223](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSeller.sql#L223) | WHERE / P12 / COMPARISON | `sth.ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `sth.ProcessCode` → `VARCHAR` | `sth.ProcessCode = 'P'` |
| IC-008412 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L34) | WHERE / P12 / COMPARISON | `tsrth.ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `tsrth.ProcessCode` → `VARCHAR` | `tsrth.ProcessCode = 'A'` |
| IC-008417 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L59) | WHERE / P12 / COMPARISON | `tsrth.ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `tsrth.ProcessCode` → `VARCHAR` | `tsrth.ProcessCode = 'A'` |
| IC-008422 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:84](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L84) | WHERE / P12 / COMPARISON | `tsrth.ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `tsrth.ProcessCode` → `VARCHAR` | `tsrth.ProcessCode = 'A'` |
| IC-008427 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L110) | WHERE / P12 / COMPARISON | `tsrth.ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `tsrth.ProcessCode` → `VARCHAR` | `tsrth.ProcessCode = 'A'` |
| IC-008432 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:136](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L136) | WHERE / P12 / COMPARISON | `tsrth.ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `tsrth.ProcessCode` → `VARCHAR` | `tsrth.ProcessCode = 'A'` |
| IC-008437 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:206](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L206) | WHERE / P12 / COMPARISON | `tsrth.ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `tsrth.ProcessCode` → `VARCHAR` | `tsrth.ProcessCode = 'A'` |
| IC-008442 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L225) | WHERE / P12 / COMPARISON | `tsrth.ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `tsrth.ProcessCode` → `VARCHAR` | `tsrth.ProcessCode = 'A'` |
| IC-008447 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:244](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L244) | WHERE / P12 / COMPARISON | `tsrth.ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `tsrth.ProcessCode` → `VARCHAR` | `tsrth.ProcessCode = 'A'` |
| IC-008452 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:264](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L264) | WHERE / P12 / COMPARISON | `tsrth.ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `tsrth.ProcessCode` → `VARCHAR` | `tsrth.ProcessCode = 'A'` |
| IC-008457 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:284](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L284) | WHERE / P12 / COMPARISON | `tsrth.ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `tsrth.ProcessCode` → `VARCHAR` | `tsrth.ProcessCode = 'A'` |
| IC-009548 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParameterSeller:95](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParameterSeller.sql#L95) | WHERE / P12 / IN_VALUE_TYPE | `ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `ProcessCode` → `VARCHAR` | `ProcessCode in ('A','R')` |
| IC-009549 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParameterSeller:139](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParameterSeller.sql#L139) | JOIN_ON / P12 / COMPARISON | `tsth.ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `tsth.ProcessCode` → `VARCHAR` | `tsth.ProcessCode = 'P'` |
| IC-009551 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParameterSeller:149](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParameterSeller.sql#L149) | JOIN_ON / P12 / COMPARISON | `tsth.ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `tsth.ProcessCode` → `VARCHAR` | `tsth.ProcessCode = 'P'` |
| IC-005209 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSParamSellerSync:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamSellerSync.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsseller_th.ProcessCode` : `CHAR(1)` | `'T'` : `varchar` | `'T'` → `CHAR(1)` | `retail.dbo.trsseller_th.ProcessCode <- 'T'` |
| IC-007391 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSeller:286](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSeller.sql#L286) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsseller_th.ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trsseller_th.ProcessCode <- 'P'` |
| IC-007393 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSeller:328](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSeller.sql#L328) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsseller_th.ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trsseller_th.ProcessCode <- 'P'` |
| IC-007399 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSeller:433](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSeller.sql#L433) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsseller_th.ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.trsseller_th.ProcessCode <- 'A'` |
| IC-007401 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSeller:485](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSeller.sql#L485) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsseller_th.ProcessCode` : `CHAR(1)` | `'R'` : `varchar` | `'R'` → `CHAR(1)` | `retail.dbo.trsseller_th.ProcessCode <- 'R'` |
