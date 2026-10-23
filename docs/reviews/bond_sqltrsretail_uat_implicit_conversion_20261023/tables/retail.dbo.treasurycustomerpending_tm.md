# Retail UAT.dbo.TreasuryCustomerPending_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 23 lokasi source, 7 procedure, 11 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 3 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cifid"></a>
## CIFId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerPending_TM.sql#L2). **Existing leading index:** TreasuryCustomerPending_TM_CN1.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000897 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.ONEApprovalPrepareData:90](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareData.sql#L90) | JOIN_ON / P01 / COMPARISON | `obliEditMasterNasabah.CIFId` : `BIGINT` | `oap.SearchTerm` : `VARCHAR(100)` | `oap.SearchTerm` → `BIGINT` | `obliEditMasterNasabah.CIFId = oap.SearchTerm` |
| IC-000925 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.ONEApprovalPrepareDataBond:88](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareDataBond.sql#L88) | JOIN_ON / P01 / COMPARISON | `obliEditMasterNasabah.CIFId` : `BIGINT` | `oap.SearchTerm` : `VARCHAR(100)` | `oap.SearchTerm` → `BIGINT` | `obliEditMasterNasabah.CIFId = oap.SearchTerm` |

<a id="kolom-citizenship"></a>
## Citizenship

**Tipe dasar:** `CHAR(3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerPending_TM.sql#L44). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006831 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessEditCustomerAuthorization:213](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessEditCustomerAuthorization.sql#L213) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tcptm.Citizenship` : `CHAR(3)` | `Citizenship` : `VARCHAR(3)` | `tcptm.Citizenship` → `VARCHAR(3)` | `Citizenship = tcptm.Citizenship` |
| IC-005402 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateEditCustomerAuthorization:476](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEditCustomerAuthorization.sql#L476) | JOIN_ON / P12 / COMPARISON | `tcptm.Citizenship` : `CHAR(3)` | `jh.JHCCOC` : `VARCHAR(3)` | `tcptm.Citizenship` → `VARCHAR` | `jh.JHCCOC = tcptm.Citizenship` |
| IC-005564 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMasterNasabahForMurex:102](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMasterNasabahForMurex.sql#L102) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tcp.Citizenship` : `CHAR(3)` | `''` : `varchar` | `''` → `CHAR(3)` | `ISNULL(tcp.Citizenship, '')` |
| IC-013911 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateTreasuryCustomer_TM:851](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L851) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerpending_tm.Citizenship` : `CHAR(3)` | `temp.Citizenship` : `varchar(3)` | `temp.Citizenship` → `CHAR(3)` | `retail.dbo.treasurycustomerpending_tm.Citizenship <- temp.Citizenship` |

<a id="kolom-kitasexpdate"></a>
## KITASExpDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerPending_TM.sql#L73). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013912 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateTreasuryCustomer_TM:851](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L851) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerpending_tm.KITASExpDate` : `DATETIME` | `temp.[KITASExpDate]` : `varchar(10)` | `temp.[KITASExpDate]` → `DATETIME` | `retail.dbo.treasurycustomerpending_tm.KITASExpDate <- temp.[KITASExpDate]` |

<a id="kolom-kitaslastupdatedate"></a>
## KITASLastUpdateDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerPending_TM.sql#L74). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013913 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateTreasuryCustomer_TM:851](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L851) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerpending_tm.KITASLastUpdateDate` : `DATETIME` | `temp.[KITASLastUpdateDate]` : `varchar(10)` | `temp.[KITASLastUpdateDate]` → `DATETIME` | `retail.dbo.treasurycustomerpending_tm.KITASLastUpdateDate <- temp.[KITASLastUpdateDate]` |

<a id="kolom-lastprocesscode"></a>
## LastProcessCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerPending_TM.sql#L40). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006834 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessEditCustomerAuthorization:335](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessEditCustomerAuthorization.sql#L335) | JOIN_ON / P12 / COMPARISON | `tcptm.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `tcptm.LastProcessCode` → `VARCHAR` | `tcptm.LastProcessCode = 'P'` |
| IC-013914 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateTreasuryCustomer_TM:851](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L851) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerpending_tm.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.treasurycustomerpending_tm.LastProcessCode <- 'P'` |

<a id="kolom-lastupdateby"></a>
## LastUpdateBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerPending_TM.sql#L27). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005405 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateEditCustomerAuthorization:482](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEditCustomerAuthorization.sql#L482) | JOIN_ON / P01 / COMPARISON | `tcptm.LastUpdateBy` : `INT` | `emp1.EMPLID` : `CHAR(11)` | `emp1.EMPLID` → `INT` | `emp1.EMPLID = tcptm.LastUpdateBy` |

<a id="kolom-nikreferentor"></a>
## NikReferentor

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerPending_TM.sql#L59). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005404 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateEditCustomerAuthorization:481](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEditCustomerAuthorization.sql#L481) | JOIN_ON / P01 / COMPARISON | `tcptm.NikReferentor` : `INT` | `emp.EMPLID` : `CHAR(11)` | `emp.EMPLID` → `INT` | `emp.EMPLID = tcptm.NikReferentor` |

<a id="kolom-operationcode"></a>
## OperationCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerPending_TM.sql#L38). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013915 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateTreasuryCustomer_TM:851](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L851) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerpending_tm.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `'U'` → `CHAR(1)` | `retail.dbo.treasurycustomerpending_tm.OperationCode <- 'U'` |

<a id="kolom-regviaonemobile"></a>
## RegViaONEMobile

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerPending_TM.sql#L71). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008534 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSResponseBONDCPYFromMX:215](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L215) | WHERE / P12 / COMPARISON | `isnull(RegViaONEMobile,'N')` : `CHAR(1)` | `'N'` : `varchar` | `isnull(RegViaONEMobile,'N')` → `VARCHAR` | `isnull(RegViaONEMobile,'N') = 'N'` |
| IC-008552 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSResponseBONDCPYFromMX:729](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L729) | WHERE / P12 / COMPARISON | `RegViaONEMobile` : `CHAR(1)` | `'Y'` : `varchar` | `RegViaONEMobile` → `VARCHAR` | `RegViaONEMobile = 'Y'` |
| IC-008558 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSResponseBONDCPYFromMX:988](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L988) | WHERE / P12 / COMPARISON | `isnull(RegViaONEMobile,'N')` : `CHAR(1)` | `'N'` : `varchar` | `isnull(RegViaONEMobile,'N')` → `VARCHAR` | `isnull(RegViaONEMobile,'N') = 'N'` |
| IC-008535 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSResponseBONDCPYFromMX:215](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L215) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `RegViaONEMobile` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `isnull(RegViaONEMobile,'N')` |
| IC-008559 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSResponseBONDCPYFromMX:988](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L988) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `RegViaONEMobile` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `isnull(RegViaONEMobile,'N')` |
| IC-013916 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateTreasuryCustomer_TM:851](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L851) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerpending_tm.RegViaONEMobile` : `CHAR(1)` | `temp.RegViaONEMobile` : `varchar(1)` | `temp.RegViaONEMobile` → `CHAR(1)` | `retail.dbo.treasurycustomerpending_tm.RegViaONEMobile <- temp.RegViaONEMobile` |

<a id="kolom-signature"></a>
## Signature

**Tipe dasar:** `VARBINARY(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerPending_TM.sql#L46). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006832 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessEditCustomerAuthorization:233](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessEditCustomerAuthorization.sql#L233) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tcptm.[Signature]` : `VARBINARY(MAX)` | `[Signature]` : `image` | `tcptm.[Signature]` → `image` | `[Signature] = tcptm.[Signature]` |
| IC-013917 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateTreasuryCustomer_TM:851](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L851) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerpending_tm.Signature` : `VARBINARY(MAX)` | `temp.[Signature]` : `image` | `temp.[Signature]` → `VARBINARY(MAX)` | `retail.dbo.treasurycustomerpending_tm.Signature <- temp.[Signature]` |

<a id="kolom-summary"></a>
## Summary

**Tipe dasar:** `VARBINARY(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerPending_TM.sql#L48). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006833 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessEditCustomerAuthorization:235](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessEditCustomerAuthorization.sql#L235) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tcptm.Summary` : `VARBINARY(MAX)` | `Summary` : `image` | `tcptm.Summary` → `image` | `Summary = tcptm.Summary` |
| IC-013918 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateTreasuryCustomer_TM:851](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L851) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerpending_tm.Summary` : `VARBINARY(MAX)` | `temp.Summary` : `image` | `temp.Summary` → `VARBINARY(MAX)` | `retail.dbo.treasurycustomerpending_tm.Summary <- temp.Summary` |
