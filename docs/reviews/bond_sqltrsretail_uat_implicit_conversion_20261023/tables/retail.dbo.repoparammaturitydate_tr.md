# Retail UAT.dbo.REPOParamMaturityDate_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 8 lokasi source, 3 procedure, 4 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-isablefortransaksi"></a>
## IsAbleForTransaksi

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOParamMaturityDate_TR.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001935 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSaveParamSecurityDirtyPrice:301](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSaveParamSecurityDirtyPrice.sql#L301) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repoparammaturitydate_tr.IsAbleForTransaksi` : `BIT` | `IsAbleForTransaksi` : `int` | `IsAbleForTransaksi` → `BIT` | `retail.dbo.repoparammaturitydate_tr.IsAbleForTransaksi <- IsAbleForTransaksi` |

<a id="kolom-iseditmaturitydate"></a>
## IsEditMaturityDate

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOParamMaturityDate_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001851 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.RepoPopulateParamHargaModal:112](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoPopulateParamHargaModal.sql#L112) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `IsEditMaturityDate` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(IsEditMaturityDate,0)` |
| IC-001922 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSaveParamHargaModal:208](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSaveParamHargaModal.sql#L208) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `IsEditMaturityDate` : `BIT` | `0` : `int` | `0` → `BIT` | `IsEditMaturityDate = 0` |
| IC-001936 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSaveParamSecurityDirtyPrice:301](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSaveParamSecurityDirtyPrice.sql#L301) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repoparammaturitydate_tr.IsEditMaturityDate` : `BIT` | `IsEditMaturityDate` : `int` | `IsEditMaturityDate` → `BIT` | `retail.dbo.repoparammaturitydate_tr.IsEditMaturityDate <- IsEditMaturityDate` |

<a id="kolom-lastupdatenik"></a>
## LastUpdateNIK

**Tipe dasar:** `VARCHAR(10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOParamMaturityDate_TR.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001934 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.RepoSaveParamSecurityDirtyPrice:266](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSaveParamSecurityDirtyPrice.sql#L266) | EXPRESSION / P07 / CASE_BRANCH | `md.LastUpdateNIK` : `VARCHAR(10)` | `result type` : `INT (combined result; facets unknown)` | `md.LastUpdateNIK` → `INT` | `CASE WHEN @IsMaturityDateChanged = 1 AND @IsStartDateChanged = 0 AND @IsDirtyPriceChanged = 0 THEN @pnNik ELSE md.LastUpdateNIK END` |
| IC-001937 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSaveParamSecurityDirtyPrice:301](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSaveParamSecurityDirtyPrice.sql#L301) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repoparammaturitydate_tr.LastUpdateNIK` : `VARCHAR(10)` | `LastUpdateNIK` : `INT (combined result; facets unknown)` | `LastUpdateNIK` → `VARCHAR(10)` | `retail.dbo.repoparammaturitydate_tr.LastUpdateNIK <- LastUpdateNIK` |

<a id="kolom-maturitydate"></a>
## MaturityDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOParamMaturityDate_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P03 — Tanggal teks bertemu date/time:** Tentukan kontrak YYYYMMDD/hari/timestamp. Bentuk nilai pencarian dalam tipe kolom atau gunakan rentang native date; jangan konversi kolom seluruh baris tanpa alasan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001921 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSaveParamHargaModal:207](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSaveParamHargaModal.sql#L207) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `MaturityDate` : `DATETIME` | `tmp.MaturityDate` : `varchar(200)` | `tmp.MaturityDate` → `DATETIME` | `MaturityDate = tmp.MaturityDate` |
| IC-001923 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RepoSaveParamHargaModal:212](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSaveParamHargaModal.sql#L212) | WHERE / P03 / COMPARISON | `pm.MaturityDate` : `DATETIME` | `tmp.MaturityDate` : `varchar(200)` | `tmp.MaturityDate` → `DATETIME` | `pm.MaturityDate != tmp.MaturityDate` |
