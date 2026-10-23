# Retail UAT.dbo.TreasuryStoreProcedureLog_TL — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 13 lokasi source, 6 procedure, 4 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-guid"></a>
## Guid

**Tipe dasar:** `VARCHAR(40)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryStoreProcedureLog_TL.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002849 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessMasterInvestor:157](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessMasterInvestor.sql#L157) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurystoreprocedurelog_tl.Guid` : `VARCHAR(40)` | `@puProcessGuid` : `uniqueidentifier` | `@puProcessGuid` → `VARCHAR(40)` | `retail.dbo.treasurystoreprocedurelog_tl.Guid <- @puProcessGuid` |
| IC-002926 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedPushMailNotifOrderCompleted:720](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L720) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurystoreprocedurelog_tl.Guid` : `VARCHAR(40)` | `@puProcessGuid` : `uniqueidentifier` | `@puProcessGuid` → `VARCHAR(40)` | `retail.dbo.treasurystoreprocedurelog_tl.Guid <- @puProcessGuid` |
| IC-002945 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedPushMailNotifOrderORI:494](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderORI.sql#L494) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurystoreprocedurelog_tl.Guid` : `VARCHAR(40)` | `@puProcessGuid` : `uniqueidentifier` | `@puProcessGuid` → `VARCHAR(40)` | `retail.dbo.treasurystoreprocedurelog_tl.Guid <- @puProcessGuid` |
| IC-004670 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSMergeFileUploadMaturity:529](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity.sql#L529) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurystoreprocedurelog_tl.Guid` : `VARCHAR(40)` | `@unGuid` : `uniqueidentifier` | `@unGuid` → `VARCHAR(40)` | `retail.dbo.treasurystoreprocedurelog_tl.Guid <- @unGuid` |
| IC-004671 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSMergeFileUploadMaturity:532](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity.sql#L532) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurystoreprocedurelog_tl.Guid` : `VARCHAR(40)` | `@unGuid` : `uniqueidentifier` | `@unGuid` → `VARCHAR(40)` | `retail.dbo.treasurystoreprocedurelog_tl.Guid <- @unGuid` |

<a id="kolom-inserteddate"></a>
## InsertedDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryStoreProcedureLog_TL.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002948 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedPushMailNotifOrderORI:498](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderORI.sql#L498) | WHERE / P04 / COMPARISON | `InsertedDate` : `DATETIME` | `CONVERT(varchar,DATEADD(day,-30,getdate()),112)` : `varchar` | `CONVERT(varchar,DATEADD(day,-30,getdate()),112)` → `DATETIME` | `InsertedDate < CONVERT(varchar,DATEADD(day,-30,getdate()),112)` |
| IC-003366 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSBTOMonitoringSISMONTAVAR:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBTOMonitoringSISMONTAVAR.sql#L21) | WHERE / P04 / COMPARISON | `InsertedDate` : `DATETIME` | `CONVERT(varchar,DATEADD(day,-30,getdate()),112)` : `varchar` | `CONVERT(varchar,DATEADD(day,-30,getdate()),112)` → `DATETIME` | `InsertedDate < CONVERT(varchar,DATEADD(day,-30,getdate()),112)` |
| IC-012705 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSValidateDocsLinkUrl:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateDocsLinkUrl.sql#L23) | WHERE / P04 / COMPARISON | `InsertedDate` : `DATETIME` | `CONVERT(varchar,DATEADD(day,-30,getdate()),112)` : `varchar` | `CONVERT(varchar,DATEADD(day,-30,getdate()),112)` → `DATETIME` | `InsertedDate < CONVERT(varchar,DATEADD(day,-30,getdate()),112)` |

<a id="kolom-parameter1"></a>
## Parameter1

**Tipe dasar:** `VARCHAR(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryStoreProcedureLog_TL.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002927 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedPushMailNotifOrderCompleted:720](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L720) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurystoreprocedurelog_tl.Parameter1` : `VARCHAR(MAX)` | `@bNotifExists` : `bit` | `@bNotifExists` → `VARCHAR(MAX)` | `retail.dbo.treasurystoreprocedurelog_tl.Parameter1 <- @bNotifExists` |
| IC-002946 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedPushMailNotifOrderORI:494](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderORI.sql#L494) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurystoreprocedurelog_tl.Parameter1` : `VARCHAR(MAX)` | `@bNotifExists` : `bit` | `@bNotifExists` → `VARCHAR(MAX)` | `retail.dbo.treasurystoreprocedurelog_tl.Parameter1 <- @bNotifExists` |
| IC-012706 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSValidateDocsLinkUrl:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateDocsLinkUrl.sql#L26) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurystoreprocedurelog_tl.Parameter1` : `VARCHAR(MAX)` | `@pcUrlInput` : `NVARCHAR(500)` | `@pcUrlInput` → `VARCHAR(MAX)` | `retail.dbo.treasurystoreprocedurelog_tl.Parameter1 <- @pcUrlInput` |

<a id="kolom-parameter2"></a>
## Parameter2

**Tipe dasar:** `VARCHAR(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryStoreProcedureLog_TL.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002928 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedPushMailNotifOrderCompleted:720](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L720) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurystoreprocedurelog_tl.Parameter2` : `VARCHAR(MAX)` | `@bMailExists` : `bit` | `@bMailExists` → `VARCHAR(MAX)` | `retail.dbo.treasurystoreprocedurelog_tl.Parameter2 <- @bMailExists` |
| IC-002947 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedPushMailNotifOrderORI:494](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderORI.sql#L494) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurystoreprocedurelog_tl.Parameter2` : `VARCHAR(MAX)` | `@bMailExists` : `bit` | `@bMailExists` → `VARCHAR(MAX)` | `retail.dbo.treasurystoreprocedurelog_tl.Parameter2 <- @bMailExists` |
