# Retail UAT.dbo.EarlyRedemption_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 10 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-actionflag"></a>
## ActionFlag

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.EarlyRedemption_TH.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 10 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006821 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessEarlyRedemAuthorization:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessEarlyRedemAuthorization.sql#L113) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.earlyredemption_th.ActionFlag` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.earlyredemption_th.ActionFlag <- 'I'` |
| IC-006822 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessEarlyRedemAuthorization:163](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessEarlyRedemAuthorization.sql#L163) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.earlyredemption_th.ActionFlag` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.earlyredemption_th.ActionFlag <- 'O'` |
| IC-006823 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessEarlyRedemAuthorization:229](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessEarlyRedemAuthorization.sql#L229) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.earlyredemption_th.ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.earlyredemption_th.ActionFlag <- 'N'` |
| IC-006824 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessEarlyRedemAuthorization:292](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessEarlyRedemAuthorization.sql#L292) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.earlyredemption_th.ActionFlag` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.earlyredemption_th.ActionFlag <- 'O'` |
| IC-006825 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessEarlyRedemAuthorization:359](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessEarlyRedemAuthorization.sql#L359) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.earlyredemption_th.ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.earlyredemption_th.ActionFlag <- 'N'` |
| IC-006826 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessEarlyRedemAuthorization:427](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessEarlyRedemAuthorization.sql#L427) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.earlyredemption_th.ActionFlag` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.earlyredemption_th.ActionFlag <- 'I'` |
| IC-006827 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessEarlyRedemAuthorization:483](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessEarlyRedemAuthorization.sql#L483) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.earlyredemption_th.ActionFlag` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.earlyredemption_th.ActionFlag <- 'O'` |
| IC-006828 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessEarlyRedemAuthorization:537](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessEarlyRedemAuthorization.sql#L537) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.earlyredemption_th.ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.earlyredemption_th.ActionFlag <- 'N'` |
| IC-006829 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessEarlyRedemAuthorization:607](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessEarlyRedemAuthorization.sql#L607) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.earlyredemption_th.ActionFlag` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.earlyredemption_th.ActionFlag <- 'O'` |
| IC-006830 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessEarlyRedemAuthorization:661](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessEarlyRedemAuthorization.sql#L661) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.earlyredemption_th.ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.earlyredemption_th.ActionFlag <- 'N'` |
