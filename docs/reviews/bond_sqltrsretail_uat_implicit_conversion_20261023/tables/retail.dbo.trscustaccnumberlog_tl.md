# Retail UAT.dbo.TRSCustAccNumberLog_TL — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 16 lokasi source, 3 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-isjoin"></a>
## IsJoin

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSCustAccNumberLog_TL.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000721 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.OMSubmitTransactionSekunder:222](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMSubmitTransactionSekunder.sql#L222) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumberlog_tl.IsJoin` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.trscustaccnumberlog_tl.IsJoin <- 'N'` |
| IC-000723 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.OMSubmitTransactionSekunder:230](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMSubmitTransactionSekunder.sql#L230) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumberlog_tl.IsJoin` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.trscustaccnumberlog_tl.IsJoin <- 'N'` |
| IC-000728 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.OMSubmitTransactionSekunder:257](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMSubmitTransactionSekunder.sql#L257) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumberlog_tl.IsJoin` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.trscustaccnumberlog_tl.IsJoin <- 'N'` |
| IC-013817 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateStatusSecurityTransactionAfterMurex:474](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransactionAfterMurex.sql#L474) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumberlog_tl.IsJoin` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.trscustaccnumberlog_tl.IsJoin <- 'N'` |
| IC-013819 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateStatusSecurityTransactionAfterMurex:482](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransactionAfterMurex.sql#L482) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumberlog_tl.IsJoin` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.trscustaccnumberlog_tl.IsJoin <- 'N'` |
| IC-013824 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateStatusSecurityTransactionAfterMurex:507](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransactionAfterMurex.sql#L507) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumberlog_tl.IsJoin` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.trscustaccnumberlog_tl.IsJoin <- 'N'` |

<a id="kolom-processcode"></a>
## ProcessCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSCustAccNumberLog_TL.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 10 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000722 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.OMSubmitTransactionSekunder:222](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMSubmitTransactionSekunder.sql#L222) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumberlog_tl.ProcessCode` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.trscustaccnumberlog_tl.ProcessCode <- 'O'` |
| IC-000724 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.OMSubmitTransactionSekunder:230](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMSubmitTransactionSekunder.sql#L230) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumberlog_tl.ProcessCode` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.trscustaccnumberlog_tl.ProcessCode <- 'N'` |
| IC-000727 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.OMSubmitTransactionSekunder:249](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMSubmitTransactionSekunder.sql#L249) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumberlog_tl.ProcessCode` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.trscustaccnumberlog_tl.ProcessCode <- 'O'` |
| IC-000729 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.OMSubmitTransactionSekunder:257](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMSubmitTransactionSekunder.sql#L257) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumberlog_tl.ProcessCode` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.trscustaccnumberlog_tl.ProcessCode <- 'N'` |
| IC-002873 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessRekDana:389](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L389) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumberlog_tl.ProcessCode` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.trscustaccnumberlog_tl.ProcessCode <- 'O'` |
| IC-002890 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessRekDana:852](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L852) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumberlog_tl.ProcessCode` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.trscustaccnumberlog_tl.ProcessCode <- 'N'` |
| IC-013818 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateStatusSecurityTransactionAfterMurex:474](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransactionAfterMurex.sql#L474) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumberlog_tl.ProcessCode` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.trscustaccnumberlog_tl.ProcessCode <- 'O'` |
| IC-013820 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateStatusSecurityTransactionAfterMurex:482](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransactionAfterMurex.sql#L482) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumberlog_tl.ProcessCode` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.trscustaccnumberlog_tl.ProcessCode <- 'N'` |
| IC-013823 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateStatusSecurityTransactionAfterMurex:499](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransactionAfterMurex.sql#L499) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumberlog_tl.ProcessCode` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.trscustaccnumberlog_tl.ProcessCode <- 'O'` |
| IC-013825 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateStatusSecurityTransactionAfterMurex:507](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransactionAfterMurex.sql#L507) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumberlog_tl.ProcessCode` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.trscustaccnumberlog_tl.ProcessCode <- 'N'` |
