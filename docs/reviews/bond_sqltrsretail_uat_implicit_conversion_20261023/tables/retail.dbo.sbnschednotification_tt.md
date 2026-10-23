# Retail UAT.dbo.SBNSchedNotification_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 10 lokasi source, 7 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-flagsyariah"></a>
## FlagSyariah

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SBNSchedNotification_TT.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002922 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedPushMailNotifOrderCompleted:636](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L636) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.sbnschednotification_tt.FlagSyariah` : `BIT` | `@nFlagSyariah` : `int` | `@nFlagSyariah` → `BIT` | `retail.dbo.sbnschednotification_tt.FlagSyariah <- @nFlagSyariah` |
| IC-002941 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedPushMailNotifOrderORI:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderORI.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.sbnschednotification_tt.FlagSyariah` : `BIT` | `@nFlagSyariah` : `int` | `@nFlagSyariah` → `BIT` | `retail.dbo.sbnschednotification_tt.FlagSyariah <- @nFlagSyariah` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SBNSchedNotification_TT.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 7 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002802 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedPopulateNotif:40](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateNotif.sql#L40) | WHERE / P12 / COMPARISON | `[Status]` : `CHAR(1)` | `'P'` : `varchar` | `[Status]` → `VARCHAR` | `[Status] = 'P'` |
| IC-002803 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedPopulateNotif:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateNotif.sql#L41) | WHERE / P12 / COMPARISON | `[Status]` : `CHAR(1)` | `'F'` : `varchar` | `[Status]` → `VARCHAR` | `[Status] = 'F'` |
| IC-004927 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSObligasiHousekeepProcess:83](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSObligasiHousekeepProcess.sql#L83) | WHERE / P12 / COMPARISON | `[Status]` : `CHAR(1)` | `'P'` : `varchar` | `[Status]` → `VARCHAR` | `[Status] <> 'P'` |
| IC-002923 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedPushMailNotifOrderCompleted:636](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L636) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.sbnschednotification_tt.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.sbnschednotification_tt.Status <- 'P'` |
| IC-002942 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedPushMailNotifOrderORI:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderORI.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.sbnschednotification_tt.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.sbnschednotification_tt.Status <- 'P'` |
| IC-002956 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedPushMailNotifRegInvestor:237](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifRegInvestor.sql#L237) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.sbnschednotification_tt.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.sbnschednotification_tt.Status <- 'P'` |
| IC-002961 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedPushMailNotifSettleRedempt:124](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifSettleRedempt.sql#L124) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.sbnschednotification_tt.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.sbnschednotification_tt.Status <- 'P'` |
| IC-004754 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSNTIPushMailNotif:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSNTIPushMailNotif.sql#L115) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.sbnschednotification_tt.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.sbnschednotification_tt.Status <- 'P'` |
