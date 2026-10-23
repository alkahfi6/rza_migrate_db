# Retail UAT.dbo.ORISBNOrderMainSched_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 12 lokasi source, 9 procedure, 6 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 4 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-accounttype"></a>
## AccountType

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISBNOrderMainSched_TM.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005444 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateGagalBayarOri:72](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L72) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `main.AccountType` : `CHAR(1)` | `#tmp_data.AccountType` : `VARCHAR(3)` | `main.AccountType` → `VARCHAR(3)` | `#tmp_data.AccountType <- main.AccountType` |
| IC-005450 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateGagalBayarOri:159](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L159) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `main.AccountType` : `CHAR(1)` | `#tmp_data.AccountType` : `VARCHAR(3)` | `main.AccountType` → `VARCHAR(3)` | `#tmp_data.AccountType <- main.AccountType` |
| IC-005457 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateGagalBayarOri:271](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L271) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `main.AccountType` : `CHAR(1)` | `#tmp_data.AccountType` : `VARCHAR(3)` | `main.AccountType` → `VARCHAR(3)` | `#tmp_data.AccountType <- main.AccountType` |
| IC-005466 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateGagalBayarOri:395](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L395) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `main.AccountType` : `CHAR(1)` | `#tmp_data.AccountType` : `VARCHAR(3)` | `main.AccountType` → `VARCHAR(3)` | `#tmp_data.AccountType <- main.AccountType` |

<a id="kolom-guidprocess"></a>
## GuidProcess

**Tipe dasar:** `UNIQUEIDENTIFIER`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISBNOrderMainSched_TM.sql#L2). **Existing leading index:** ORISBNOrderMainSched_TM_CN1.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002861 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedProcessORIPaymentDirect:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessORIPaymentDirect.sql#L16) | WHERE / P13 / COMPARISON | `GuidProcess` : `UNIQUEIDENTIFIER` | `@puGuidProcess` : `varchar(100)` | `@puGuidProcess` → `UNIQUEIDENTIFIER` | `GuidProcess = @puGuidProcess` |

<a id="kolom-idreksb"></a>
## IdRekSB

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISBNOrderMainSched_TM.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000959 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.ORIOnlineReportDailyTransaction:364](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORIOnlineReportDailyTransaction.sql#L364) | JOIN_ON / P01 / COMPARISON | `tc.IdRekSB` : `INT` | `sb.IdRekSB` : `VARCHAR(20)` | `sb.IdRekSB` → `INT` | `tc.IdRekSB = sb.IdRekSB` |
| IC-008371 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportReportSettlementInstruction:619](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L619) | JOIN_ON / P01 / COMPARISON | `scd.IdRekSB` : `INT` | `sb.IdRekSB` : `VARCHAR(20)` | `sb.IdRekSB` → `INT` | `scd.IdRekSB = sb.IdRekSB` |
| IC-010343 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptOrderORIDetail:836](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOrderORIDetail.sql#L836) | JOIN_ON / P01 / COMPARISON | `tc.IdRekSB` : `INT` | `sb.IdRekSB` : `VARCHAR(20)` | `sb.IdRekSB` → `INT` | `tc.IdRekSB = sb.IdRekSB` |

<a id="kolom-idseri"></a>
## IdSeri

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISBNOrderMainSched_TM.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002811 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedPopulateORIOrder:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateORIOrder.sql#L110) | JOIN_ON / P01 / COMPARISON | `ms.IdSeri` : `INT` | `tm.IdSeri` : `VARCHAR(100)` | `tm.IdSeri` → `INT` | `ms.IdSeri = tm.IdSeri` |
| IC-002835 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedPopulateRegInvestorForOrder:83](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestorForOrder.sql#L83) | JOIN_ON / P01 / COMPARISON | `scd.IdSeri` : `INT` | `tm.IdSeri` : `VARCHAR(100)` | `tm.IdSeri` → `INT` | `scd.IdSeri = tm.IdSeri` |

<a id="kolom-insertdate"></a>
## InsertDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISBNOrderMainSched_TM.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002807 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedPopulateNotif:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateNotif.sql#L63) | WHERE / P04 / COMPARISON | `trm.InsertDate` : `DATETIME` | `CONVERT(VARCHAR(8),GETDATE(),112)` : `VARCHAR(8)` | `CONVERT(VARCHAR(8),GETDATE(),112)` → `DATETIME` | `trm.InsertDate >= CONVERT(VARCHAR(8),GETDATE(),112)` |

<a id="kolom-nominal"></a>
## Nominal

**Tipe dasar:** `DECIMAL(21,0)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISBNOrderMainSched_TM.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004821 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSORIOnlineNewTransaction:443](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlineNewTransaction.sql#L443) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.orisbnordermainsched_tm.Nominal` : `DECIMAL(21,0)` | `[OrderNominal]` : `money` | `[OrderNominal]` → `DECIMAL(21,0)` | `retail.dbo.orisbnordermainsched_tm.Nominal <- [OrderNominal]` |
