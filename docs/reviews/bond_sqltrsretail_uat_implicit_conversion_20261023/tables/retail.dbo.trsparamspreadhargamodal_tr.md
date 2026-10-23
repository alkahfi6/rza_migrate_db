# Retail UAT.dbo.TRSParamSpreadHargaModal_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 11 lokasi source, 3 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-lastupdatenik"></a>
## LastUpdateNIK

**Tipe dasar:** `VARCHAR(10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamSpreadHargaModal_TR.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012339 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSubmitParamSpreadHargaModal:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitParamSpreadHargaModal.sql#L86) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamspreadhargamodal_tr.LastUpdateNIK` : `VARCHAR(10)` | `@pnNIK` : `INT` | `@pnNIK` → `VARCHAR(10)` | `retail.dbo.trsparamspreadhargamodal_tr.LastUpdateNIK <- @pnNIK` |

<a id="kolom-maxspread"></a>
## MaxSpread

**Tipe dasar:** `DECIMAL(10,4)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamSpreadHargaModal_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005937 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateParamSettingHargaModal:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSettingHargaModal.sql#L26) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `shm.MaxSpread` : `DECIMAL(10,4)` | `0` : `int` | `0` → `DECIMAL(10,4)` | `ISNULL(shm.MaxSpread,0)` |
| IC-012343 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitParamSpreadHargaModal:109](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitParamSpreadHargaModal.sql#L109) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `par.MaxSpread` : `DECIMAL(10,4)` | `0` : `int` | `0` → `DECIMAL(10,4)` | `ISNULL(par.MaxSpread,0)` |
| IC-012348 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitParamSpreadHargaModal:124](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitParamSpreadHargaModal.sql#L124) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `par.MaxSpread` : `DECIMAL(10,4)` | `0` : `int` | `0` → `DECIMAL(10,4)` | `ISNULL(par.MaxSpread,0)` |
| IC-012354 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitParamSpreadHargaModal:133](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitParamSpreadHargaModal.sql#L133) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `par.MaxSpread` : `DECIMAL(10,4)` | `0` : `int` | `0` → `DECIMAL(10,4)` | `ISNULL(par.MaxSpread,0)` |
| IC-012716 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateParamSpreadHargaModal:123](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateParamSpreadHargaModal.sql#L123) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `spr.MaxSpread` : `DECIMAL(10,4)` | `0` : `int` | `0` → `DECIMAL(10,4)` | `ISNULL(spr.MaxSpread,0)` |

<a id="kolom-minspread"></a>
## MinSpread

**Tipe dasar:** `DECIMAL(10,4)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamSpreadHargaModal_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005936 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateParamSettingHargaModal:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSettingHargaModal.sql#L25) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `shm.MinSpread` : `DECIMAL(10,4)` | `0` : `int` | `0` → `DECIMAL(10,4)` | `ISNULL(shm.MinSpread,0)` |
| IC-012341 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitParamSpreadHargaModal:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitParamSpreadHargaModal.sql#L108) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `par.MinSpread` : `DECIMAL(10,4)` | `0` : `int` | `0` → `DECIMAL(10,4)` | `ISNULL(par.MinSpread,0)` |
| IC-012346 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitParamSpreadHargaModal:123](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitParamSpreadHargaModal.sql#L123) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `par.MinSpread` : `DECIMAL(10,4)` | `0` : `int` | `0` → `DECIMAL(10,4)` | `ISNULL(par.MinSpread,0)` |
| IC-012352 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitParamSpreadHargaModal:132](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitParamSpreadHargaModal.sql#L132) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `par.MinSpread` : `DECIMAL(10,4)` | `0` : `int` | `0` → `DECIMAL(10,4)` | `ISNULL(par.MinSpread,0)` |
| IC-012715 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateParamSpreadHargaModal:122](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateParamSpreadHargaModal.sql#L122) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `spr.MinSpread` : `DECIMAL(10,4)` | `0` : `int` | `0` → `DECIMAL(10,4)` | `ISNULL(spr.MinSpread,0)` |
