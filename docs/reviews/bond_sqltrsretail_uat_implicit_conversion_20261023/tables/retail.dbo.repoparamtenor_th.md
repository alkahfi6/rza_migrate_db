# Retail UAT.dbo.REPOParamTenor_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 5 lokasi source, 3 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-active"></a>
## Active

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOParamTenor_TH.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001822 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoMaintenanceParamHargaModal:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoMaintenanceParamHargaModal.sql#L28) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repoparamtenor_th.Active` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.repoparamtenor_th.Active <- 0` |

<a id="kolom-lastupdatenik"></a>
## LastUpdateNIK

**Tipe dasar:** `VARCHAR(10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOParamTenor_TH.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001915 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSaveParamHargaModal:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSaveParamHargaModal.sql#L162) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repoparamtenor_th.LastUpdateNIK` : `VARCHAR(10)` | `@pnNik` : `INT` | `@pnNik` → `VARCHAR(10)` | `retail.dbo.repoparamtenor_th.LastUpdateNIK <- @pnNik` |

<a id="kolom-maturitydate"></a>
## MaturityDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOParamTenor_TH.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012196 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSSentEmailHargaRepo:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailHargaRepo.sql#L22) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `MaturityDate` : `DATETIME` | `@MATURITYDATEOLD` : `VARCHAR(100)` | `MaturityDate` → `VARCHAR(100)` | `@MATURITYDATEOLD = MaturityDate` |
| IC-001916 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSaveParamHargaModal:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSaveParamHargaModal.sql#L162) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repoparamtenor_th.MaturityDate` : `DATETIME` | `MaturityDate` : `varchar(200)` | `MaturityDate` → `DATETIME` | `retail.dbo.repoparamtenor_th.MaturityDate <- MaturityDate` |
| IC-001928 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSaveParamHargaModal:298](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSaveParamHargaModal.sql#L298) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repoparamtenor_th.MaturityDate` : `DATETIME` | `tmp.MaturityDate` : `varchar(200)` | `tmp.MaturityDate` → `DATETIME` | `retail.dbo.repoparamtenor_th.MaturityDate <- tmp.MaturityDate` |
