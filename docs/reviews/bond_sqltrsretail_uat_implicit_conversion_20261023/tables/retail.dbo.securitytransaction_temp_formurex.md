# Retail UAT.dbo.SecurityTransaction_TEMP_ForMurex — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 7 lokasi source, 2 procedure, 6 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-accstatus"></a>
## AccStatus

**Tipe dasar:** `TINYINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TEMP_ForMurex.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013814 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateStatusSecurityTransactionAfterMurex:288](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransactionAfterMurex.sql#L288) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `temp.[AccStatus]` : `TINYINT` | `3` : `int` | `3` → `TINYINT` | `temp.[AccStatus] = 3` |

<a id="kolom-dealid"></a>
## DealId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TEMP_ForMurex.sql#L4). **Existing leading index:** SecurityTransaction_TEMP_ForMurex_NN1.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008568 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSResponseBondFromMX:186](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L186) | JOIN_ON / P11 / COMPARISON | `fm.DealId` : `BIGINT` | `td.DealId` : `int` | `td.DealId` → `BIGINT` | `td.DealId = fm.DealId` |
| IC-008570 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSResponseBondFromMX:233](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L233) | JOIN_ON / P11 / COMPARISON | `fm.DealId` : `BIGINT` | `td.DealId` : `int` | `td.DealId` → `BIGINT` | `td.DealId = fm.DealId` |

<a id="kolom-lastupdatedate"></a>
## LastUpdateDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TEMP_ForMurex.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013808 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateStatusSecurityTransactionAfterMurex:254](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransactionAfterMurex.sql#L254) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_temp_formurex.LastUpdateDate` : `DATETIME` | `@dcurrent_working_date` : `smalldatetime` | `@dcurrent_working_date` → `DATETIME` | `retail.dbo.securitytransaction_temp_formurex.LastUpdateDate <- @dcurrent_working_date` |

<a id="kolom-skflogstatus"></a>
## SKFLogStatus

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TEMP_ForMurex.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013812 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateStatusSecurityTransactionAfterMurex:277](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransactionAfterMurex.sql#L277) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `SKFLogStatus` : `CHAR(1)` | `'S'` : `varchar` | `'S'` → `CHAR(1)` | `SKFLogStatus = 'S'` |

<a id="kolom-trxstatus"></a>
## TrxStatus

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TEMP_ForMurex.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013809 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateStatusSecurityTransactionAfterMurex:254](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransactionAfterMurex.sql#L254) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_temp_formurex.TrxStatus` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `retail.dbo.securitytransaction_temp_formurex.TrxStatus <- 3` |

<a id="kolom-trxtype"></a>
## TrxType

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TEMP_ForMurex.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013810 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateStatusSecurityTransactionAfterMurex:254](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransactionAfterMurex.sql#L254) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_temp_formurex.TrxType` : `SMALLINT` | `@nTrxType` : `int` | `@nTrxType` → `SMALLINT` | `retail.dbo.securitytransaction_temp_formurex.TrxType <- @nTrxType` |
