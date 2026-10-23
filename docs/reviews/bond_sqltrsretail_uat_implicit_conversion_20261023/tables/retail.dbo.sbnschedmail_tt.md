# Retail UAT.dbo.SBNSchedMail_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 10 lokasi source, 7 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-flagsyariah"></a>
## FlagSyariah

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SBNSchedMail_TT.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002924 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedPushMailNotifOrderCompleted:683](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L683) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.sbnschedmail_tt.FlagSyariah` : `BIT` | `@nFlagSyariah` : `int` | `@nFlagSyariah` → `BIT` | `retail.dbo.sbnschedmail_tt.FlagSyariah <- @nFlagSyariah` |
| IC-002943 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedPushMailNotifOrderORI:457](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderORI.sql#L457) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.sbnschedmail_tt.FlagSyariah` : `BIT` | `@nFlagSyariah` : `int` | `@nFlagSyariah` → `BIT` | `retail.dbo.sbnschedmail_tt.FlagSyariah <- @nFlagSyariah` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SBNSchedMail_TT.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 7 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002800 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedPopulateNotif:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateNotif.sql#L22) | WHERE / P12 / COMPARISON | `[Status]` : `CHAR(1)` | `'P'` : `varchar` | `[Status]` → `VARCHAR` | `[Status] = 'P'` |
| IC-002801 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedPopulateNotif:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateNotif.sql#L23) | WHERE / P12 / COMPARISON | `[Status]` : `CHAR(1)` | `'F'` : `varchar` | `[Status]` → `VARCHAR` | `[Status] = 'F'` |
| IC-004928 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSObligasiHousekeepProcess:84](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSObligasiHousekeepProcess.sql#L84) | WHERE / P12 / COMPARISON | `[Status]` : `CHAR(1)` | `'P'` : `varchar` | `[Status]` → `VARCHAR` | `[Status] <> 'P'` |
| IC-002925 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedPushMailNotifOrderCompleted:683](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L683) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.sbnschedmail_tt.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.sbnschedmail_tt.Status <- 'P'` |
| IC-002944 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedPushMailNotifOrderORI:457](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderORI.sql#L457) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.sbnschedmail_tt.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.sbnschedmail_tt.Status <- 'P'` |
| IC-002957 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedPushMailNotifRegInvestor:278](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifRegInvestor.sql#L278) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.sbnschedmail_tt.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.sbnschedmail_tt.Status <- 'P'` |
| IC-002962 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedPushMailNotifSettleRedempt:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifSettleRedempt.sql#L165) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.sbnschedmail_tt.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.sbnschedmail_tt.Status <- 'P'` |
| IC-004753 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSNTIPushMailNotif:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSNTIPushMailNotif.sql#L107) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.sbnschedmail_tt.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.sbnschedmail_tt.Status <- 'P'` |
