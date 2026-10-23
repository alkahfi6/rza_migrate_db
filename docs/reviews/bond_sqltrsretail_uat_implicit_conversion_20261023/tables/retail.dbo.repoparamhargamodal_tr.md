# Retail UAT.dbo.REPOParamHargaModal_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 10 lokasi source, 6 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-iseditdirtyprice"></a>
## IsEditDirtyPrice

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOParamHargaModal_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001829 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoMaintenanceParamHargaModal:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoMaintenanceParamHargaModal.sql#L89) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `IsEditDirtyPrice` : `BIT` | `0` : `int` | `0` → `BIT` | `IsEditDirtyPrice = 0` |
| IC-002602 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSSaveParamHargaModalNew_Normal:371](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSaveParamHargaModalNew_Normal.sql#L371) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `IsEditDirtyPrice` : `BIT` | `0` : `int` | `0` → `BIT` | `IsEditDirtyPrice = 0` |
| IC-006669 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessApprovalHargaModal:102](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalHargaModal.sql#L102) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `IsEditDirtyPrice` : `BIT` | `0` : `int` | `0` → `BIT` | `IsEditDirtyPrice = 0` |
| IC-011845 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveParamHargaModalNew:466](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L466) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `IsEditDirtyPrice` : `BIT` | `0` : `int` | `0` → `BIT` | `IsEditDirtyPrice = 0` |
| IC-011852 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveParamHargaModalNew:513](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L513) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `IsEditDirtyPrice` : `BIT` | `0` : `int` | `0` → `BIT` | `IsEditDirtyPrice = 0` |

<a id="kolom-iseditstartdate"></a>
## IsEditStartDate

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOParamHargaModal_TR.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001828 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoMaintenanceParamHargaModal:88](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoMaintenanceParamHargaModal.sql#L88) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `IsEditStartDate` : `BIT` | `0` : `int` | `0` → `BIT` | `IsEditStartDate = 0` |
| IC-011968 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveParamSettlementDate:495](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamSettlementDate.sql#L495) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `IsEditStartDate` : `BIT` | `0` : `int` | `0` → `BIT` | `IsEditStartDate = 0` |

<a id="kolom-lastupdatenik"></a>
## LastUpdateNIK

**Tipe dasar:** `VARCHAR(10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOParamHargaModal_TR.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001933 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.RepoSaveParamSecurityDirtyPrice:213](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSaveParamSecurityDirtyPrice.sql#L213) | EXPRESSION / P07 / CASE_BRANCH | `LastUpdateNIK` : `VARCHAR(10)` | `result type` : `INT (combined result; facets unknown)` | `LastUpdateNIK` → `INT` | `CASE WHEN @IsMaturityDateChanged = 1 AND (@IsStartDateChanged = 1 OR @IsDirtyPriceChanged = 1) THEN @pnNik ELSE LastUpdateNIK END` |
| IC-001931 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSaveParamSecurityDirtyPrice:161](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSaveParamSecurityDirtyPrice.sql#L161) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repoparamhargamodal_tr.LastUpdateNIK` : `VARCHAR(10)` | `@pnNik` : `INT` | `@pnNik` → `VARCHAR(10)` | `retail.dbo.repoparamhargamodal_tr.LastUpdateNIK <- @pnNik` |
| IC-001932 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSaveParamSecurityDirtyPrice:209](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSaveParamSecurityDirtyPrice.sql#L209) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `LastUpdateNIK` : `VARCHAR(10)` | `CASE @IsMaturityDateChanged WHEN 0 THEN @pnNik ELSE CASE WHEN @IsMaturityDateChanged = 1 AND (@IsStartDateChanged = 1 OR @IsDirtyPriceChanged = 1) THEN @pnNik ELSE LastUpdateNIK END END` : `INT (combined result; facets unknown)` | `CASE @IsMaturityDateChanged WHEN 0 THEN @pnNik ELSE CASE WHEN @IsMaturityDateChanged = 1 AND (@IsStartDateChanged = 1 OR @IsDirtyPriceChanged = 1) THEN @pnNik ELSE LastUpdateNIK END END` → `VARCHAR(10)` | `LastUpdateNIK = CASE @IsMaturityDateChanged WHEN 0 THEN @pnNik ELSE CASE WHEN @IsMaturityDateChanged = 1 AND (@IsStartDateChanged = 1 OR @IsDirtyPriceChanged = 1) THEN @pnNik ELSE LastUpdateNIK END END` |
