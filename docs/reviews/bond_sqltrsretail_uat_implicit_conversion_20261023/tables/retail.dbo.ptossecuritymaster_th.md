# Retail UAT.dbo.PTOSSecurityMaster_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 5 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-actionflag"></a>
## ActionFlag

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSSecurityMaster_TH.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007143 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSSecurityMasterAuthen:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSSecurityMasterAuthen.sql#L63) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptossecuritymaster_th.ActionFlag` : `CHAR(1)` | `CASE WHEN smptm.Status = 0 THEN 'I' WHEN smptm.Status = 1 THEN 'U' WHEN smptm.Status = 2 THEN 'D' END` : `VARCHAR (combined result; facets unknown)` | `CASE WHEN smptm.Status = 0 THEN 'I' WHEN smptm.Status = 1 THEN 'U' WHEN smptm.Status = 2 THEN 'D' END` → `CHAR(1)` | `retail.dbo.ptossecuritymaster_th.ActionFlag <- CASE WHEN smptm.Status = 0 THEN 'I' WHEN smptm.Status = 1 THEN 'U' WHEN smptm.Status = 2 THEN 'D' END` |
| IC-007153 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSSecurityMasterAuthen:146](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSSecurityMasterAuthen.sql#L146) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptossecuritymaster_th.ActionFlag` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.ptossecuritymaster_th.ActionFlag <- 'I'` |
| IC-007156 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSSecurityMasterAuthen:182](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSSecurityMasterAuthen.sql#L182) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptossecuritymaster_th.ActionFlag` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.ptossecuritymaster_th.ActionFlag <- 'O'` |
| IC-007158 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSSecurityMasterAuthen:224](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSSecurityMasterAuthen.sql#L224) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptossecuritymaster_th.ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.ptossecuritymaster_th.ActionFlag <- 'N'` |
| IC-007161 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSSecurityMasterAuthen:266](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSSecurityMasterAuthen.sql#L266) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptossecuritymaster_th.ActionFlag` : `CHAR(1)` | `'D'` : `varchar` | `'D'` → `CHAR(1)` | `retail.dbo.ptossecuritymaster_th.ActionFlag <- 'D'` |
