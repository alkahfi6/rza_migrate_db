# Retail UAT.dbo.TRSGeneralParam_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 15 lokasi source, 8 procedure, 2 kolom. **A:** 2 lokasi pada 2 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-paramid"></a>
## ParamId

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSGeneralParam_TR.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 11 lokasi / 6 SP. **A:** 2 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN, COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005341 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateCboParam:12](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateCboParam.sql#L12) | WHERE / P01 / IN_VALUE_TYPE | `ParamId` : `VARCHAR(20)` | `2` : `int` | `ParamId` → `INT` | `ParamId not in (2,3)` |
| IC-005557 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateMainApprovalParam:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMainApprovalParam.sql#L44) | JOIN_ON / P01 / COMPARISON | `gp.ParamId` : `VARCHAR(20)` | `td.ApprTypeId` : `int` | `gp.ParamId` → `INT` | `td.ApprTypeId = gp.ParamId` |
| IC-005340 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateCboParam:7](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateCboParam.sql#L7) | SET_OUTPUT / P07 / UNION_OUTPUT | `right output #1` : `VARCHAR(20)` | `left output #1` : `int` | `right output #1` → `INT` | `select 0 as ParamId, 'All' as ParamValue union all select ParamId, ParamValue from dbo.TRSGeneralParam_TR where GroupId = 1 and ParamId not in (2,3)` |
| IC-005228 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateApprParamBookbld:6](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamBookbld.sql#L6) | WHERE / P12 / COMPARISON | `ParamId` : `VARCHAR(20)` | `ActionFlag` : `CHAR(1)` | `ActionFlag` → `VARCHAR` | `ParamId = ActionFlag` |
| IC-005230 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateApprParamBookbld:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamBookbld.sql#L16) | WHERE / P12 / COMPARISON | `ParamId` : `VARCHAR(20)` | `ActionFlag` : `CHAR(1)` | `ActionFlag` → `VARCHAR` | `ParamId = ActionFlag` |
| IC-005232 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateApprParamDeviasiHarga:6](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamDeviasiHarga.sql#L6) | WHERE / P12 / COMPARISON | `ParamId` : `VARCHAR(20)` | `ActionFlag` : `CHAR(1)` | `ActionFlag` → `VARCHAR` | `ParamId = ActionFlag` |
| IC-005234 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateApprParamDeviasiHarga:20](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamDeviasiHarga.sql#L20) | WHERE / P12 / COMPARISON | `ParamId` : `VARCHAR(20)` | `ActionFlag` : `CHAR(1)` | `ActionFlag` → `VARCHAR` | `ParamId = ActionFlag` |
| IC-005238 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateApprParamHargaModal:6](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamHargaModal.sql#L6) | WHERE / P12 / COMPARISON | `ParamId` : `VARCHAR(20)` | `ActionFlag` : `CHAR(1)` | `ActionFlag` → `VARCHAR` | `ParamId = ActionFlag` |
| IC-005240 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateApprParamHargaModal:19](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamHargaModal.sql#L19) | WHERE / P12 / COMPARISON | `ParamId` : `VARCHAR(20)` | `ActionFlag` : `CHAR(1)` | `ActionFlag` → `VARCHAR` | `ParamId = ActionFlag` |
| IC-005242 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateApprParamSpread:10](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamSpread.sql#L10) | WHERE / P12 / COMPARISON | `ParamId` : `VARCHAR(20)` | `ActionFlag` : `CHAR(1)` | `ActionFlag` → `VARCHAR` | `ParamId = ActionFlag` |
| IC-005245 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateApprParamSpread:35](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamSpread.sql#L35) | WHERE / P12 / COMPARISON | `ParamId` : `VARCHAR(20)` | `ActionFlag` : `CHAR(1)` | `ActionFlag` → `VARCHAR` | `ParamId = ActionFlag` |

<a id="kolom-groupid"></a>
## GroupId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSGeneralParam_TR.sql#L2). **Existing leading index:** TRSGeneralParam_TR_CU1.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012214 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSentEmailHasilRevenueBondGagal:157](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailHasilRevenueBondGagal.sql#L157) | JOIN_ON / P01 / COMPARISON | `gp.GroupId` : `INT` | `np.Value` : `VARCHAR(MAX)` | `np.Value` → `INT` | `gp.GroupId = np.Value` |
| IC-012218 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSentEmailTransferAsset:92](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailTransferAsset.sql#L92) | JOIN_ON / P01 / COMPARISON | `gp.GroupId` : `INT` | `np.Value` : `VARCHAR(MAX)` | `np.Value` → `INT` | `gp.GroupId = np.Value` |
| IC-012219 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSentEmailTransferAsset:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailTransferAsset.sql#L106) | JOIN_ON / P01 / COMPARISON | `gp.GroupId` : `INT` | `np.Value` : `VARCHAR(MAX)` | `np.Value` → `INT` | `gp.GroupId = np.Value` |
| IC-012220 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSentEmailTransferAsset:126](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailTransferAsset.sql#L126) | JOIN_ON / P01 / COMPARISON | `gp.GroupId` : `INT` | `np.Value` : `VARCHAR(MAX)` | `np.Value` → `INT` | `gp.GroupId = np.Value` |
