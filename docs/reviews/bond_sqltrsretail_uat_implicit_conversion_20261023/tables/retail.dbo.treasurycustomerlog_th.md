# Retail UAT.dbo.TreasuryCustomerLog_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 29 lokasi source, 13 procedure, 13 kolom. **A:** 2 lokasi pada 1 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-jenisidentitas"></a>
## JenisIdentitas

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerLog_TH.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009588 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptCustomerLog:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerLog.sql#L107) | JOIN_ON / P01 / COMPARISON | `tclth.JenisIdentitas` : `CHAR(1)` | `ji.Id` : `INT` | `tclth.JenisIdentitas` → `INT` | `tclth.JenisIdentitas = ji.Id` |

<a id="kolom-jeniskelamin"></a>
## JenisKelamin

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerLog_TH.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009589 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptCustomerLog:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerLog.sql#L111) | JOIN_ON / P01 / COMPARISON | `tclth.JenisKelamin` : `CHAR(1)` | `jk.Id` : `INT` | `tclth.JenisKelamin` → `INT` | `tclth.JenisKelamin = jk.Id` |

<a id="kolom-citizenship"></a>
## Citizenship

**Tipe dasar:** `CHAR(3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerLog_TH.sql#L51). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013180 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:973](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L973) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.Citizenship` : `CHAR(3)` | `ttc.Citizenship` : `varchar(3)` | `ttc.Citizenship` → `CHAR(3)` | `retail.dbo.treasurycustomerlog_th.Citizenship <- ttc.Citizenship` |

<a id="kolom-idreviewbit"></a>
## IDReviewBit

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerLog_TH.sql#L79). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013181 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:973](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L973) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.IDReviewBit` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.treasurycustomerlog_th.IDReviewBit <- 0` |

<a id="kolom-insertedby"></a>
## InsertedBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerLog_TH.sql#L25). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013182 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:973](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L973) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.InsertedBy` : `INT` | `ttc.InsertedBy` : `varchar(20)` | `ttc.InsertedBy` → `INT` | `retail.dbo.treasurycustomerlog_th.InsertedBy <- ttc.InsertedBy` |

<a id="kolom-kitasexpdate"></a>
## KITASExpDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerLog_TH.sql#L81). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013183 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:973](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L973) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.KITASExpDate` : `DATETIME` | `ttc.[KITASExpDate]` : `varchar(10)` | `ttc.[KITASExpDate]` → `DATETIME` | `retail.dbo.treasurycustomerlog_th.KITASExpDate <- ttc.[KITASExpDate]` |

<a id="kolom-kitaslastupdatedate"></a>
## KITASLastUpdateDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerLog_TH.sql#L82). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013184 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:973](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L973) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.KITASLastUpdateDate` : `DATETIME` | `ttc.[KITASLastUpdateDate]` : `varchar(10)` | `ttc.[KITASLastUpdateDate]` → `DATETIME` | `retail.dbo.treasurycustomerlog_th.KITASLastUpdateDate <- ttc.[KITASLastUpdateDate]` |

<a id="kolom-lastupdateby"></a>
## LastUpdateBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerLog_TH.sql#L27). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008652 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBondFromMX:815](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L815) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.LastUpdateBy` : `INT` | `'77777'` : `varchar` | `'77777'` → `INT` | `retail.dbo.treasurycustomerlog_th.LastUpdateBy <- '77777'` |
| IC-013185 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:973](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L973) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.LastUpdateBy` : `INT` | `ttc.LastUpdateBy` : `varchar(20)` | `ttc.LastUpdateBy` → `INT` | `retail.dbo.treasurycustomerlog_th.LastUpdateBy <- ttc.LastUpdateBy` |

<a id="kolom-operationcode"></a>
## OperationCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerLog_TH.sql#L44). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 15 lokasi / 11 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008469 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTemplateSID:54](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSID.sql#L54) | WHERE / P12 / COMPARISON | `OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `OperationCode` → `VARCHAR` | `OperationCode = 'I'` |
| IC-008481 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTemplateSIDBySecAccNo:19](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSIDBySecAccNo.sql#L19) | WHERE / P12 / COMPARISON | `OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `OperationCode` → `VARCHAR` | `OperationCode = 'I'` |
| IC-000803 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.OMUpdateStatusTreasuryCust:14](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMUpdateStatusTreasuryCust.sql#L14) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.OperationCode` : `CHAR(1)` | `'F'` : `varchar` | `'F'` → `CHAR(1)` | `retail.dbo.treasurycustomerlog_th.OperationCode <- 'F'` |
| IC-002777 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_trs_updateStatusTreasuryCustomer_TM:99](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_updateStatusTreasuryCustomer_TM.sql#L99) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.treasurycustomerlog_th.OperationCode <- 'I'` |
| IC-002871 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessRekDana:244](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L244) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.OperationCode` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.treasurycustomerlog_th.OperationCode <- 'O'` |
| IC-002888 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessRekDana:707](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L707) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `'U'` → `CHAR(1)` | `retail.dbo.treasurycustomerlog_th.OperationCode <- 'U'` |
| IC-008533 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBONDCPYFromMX:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L48) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.treasurycustomerlog_th.OperationCode <- 'I'` |
| IC-008546 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBONDCPYFromMX:424](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L424) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `'U'` → `CHAR(1)` | `retail.dbo.treasurycustomerlog_th.OperationCode <- 'U'` |
| IC-008557 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBONDCPYFromMX:802](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L802) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `'U'` → `CHAR(1)` | `retail.dbo.treasurycustomerlog_th.OperationCode <- 'U'` |
| IC-008653 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBondFromMX:815](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L815) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `'U'` → `CHAR(1)` | `retail.dbo.treasurycustomerlog_th.OperationCode <- 'U'` |
| IC-009147 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailSyncDataCore:149](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L149) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `'U'` → `CHAR(1)` | `retail.dbo.treasurycustomerlog_th.OperationCode <- 'U'` |
| IC-009312 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailSyncDataCore:860](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L860) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `'U'` → `CHAR(1)` | `retail.dbo.treasurycustomerlog_th.OperationCode <- 'U'` |
| IC-009476 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:724](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L724) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `'U'` → `CHAR(1)` | `retail.dbo.treasurycustomerlog_th.OperationCode <- 'U'` |
| IC-012673 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateTaxAmnestyExpired:298](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateTaxAmnestyExpired.sql#L298) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `'U'` → `CHAR(1)` | `retail.dbo.treasurycustomerlog_th.OperationCode <- 'U'` |
| IC-013994 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_updateStatusTreasuryCustomer_TM:99](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateStatusTreasuryCustomer_TM.sql#L99) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.treasurycustomerlog_th.OperationCode <- 'I'` |

<a id="kolom-processby"></a>
## ProcessBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerLog_TH.sql#L48). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008654 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBondFromMX:815](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L815) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.ProcessBy` : `INT` | `'77777'` : `varchar` | `'77777'` → `INT` | `retail.dbo.treasurycustomerlog_th.ProcessBy <- '77777'` |
| IC-012674 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateTaxAmnestyExpired:298](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateTaxAmnestyExpired.sql#L298) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.ProcessBy` : `INT` | `'77777'` : `varchar` | `'77777'` → `INT` | `retail.dbo.treasurycustomerlog_th.ProcessBy <- '77777'` |

<a id="kolom-signature"></a>
## Signature

**Tipe dasar:** `VARBINARY(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerLog_TH.sql#L53). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013186 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:973](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L973) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.Signature` : `VARBINARY(MAX)` | `[Signature]` : `image` | `[Signature]` → `VARBINARY(MAX)` | `retail.dbo.treasurycustomerlog_th.Signature <- [Signature]` |

<a id="kolom-summary"></a>
## Summary

**Tipe dasar:** `VARBINARY(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerLog_TH.sql#L55). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013187 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:973](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L973) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.Summary` : `VARBINARY(MAX)` | `Summary` : `image` | `Summary` → `VARBINARY(MAX)` | `retail.dbo.treasurycustomerlog_th.Summary <- Summary` |

<a id="kolom-valasisjoin"></a>
## ValasIsJoin

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerLog_TH.sql#L41). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013188 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:973](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L973) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerlog_th.ValasIsJoin` : `CHAR(1)` | `tcan.STATUS_JOIN` : `varchar(1 )` | `tcan.STATUS_JOIN` → `CHAR(1)` | `retail.dbo.treasurycustomerlog_th.ValasIsJoin <- tcan.STATUS_JOIN` |
