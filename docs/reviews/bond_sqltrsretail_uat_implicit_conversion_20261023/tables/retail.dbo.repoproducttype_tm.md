# Retail UAT.dbo.REPOProductType_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 5 lokasi source, 3 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-iddata"></a>
## IdData

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOProductType_TM.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005930 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateParamRepo:100](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamRepo.sql#L100) | SET_OUTPUT / P07 / UNION_OUTPUT | `right output #1` : `INT` | `left output #1` : `varchar` | `left output #1` → `INT` | `SELECT '' AS IdData, '' AS RepoTypeId, 'ALL' AS RepoTypeDesc, CAST(0 AS BIT) AS isUsingDeposito, CAST(0 AS BIT) AS isBookBuilding, CAST(0 AS BIT) AS isUsingHairCut UNION ALL SELECT IdData ,RepoTypeId ,RepoTypeDesc ,isUsingDeposito ,isBookBuilding ,isUsingHairCut FROM dbo.REPOProductType_TM` |

<a id="kolom-isusingdeposito"></a>
## isUsingDeposito

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOProductType_TM.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001965 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.RepoUploadValidation:207](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoUploadValidation.sql#L207) | WHERE / P11 / COMPARISON | `isUsingDeposito` : `BIT` | `0` : `int` | `isUsingDeposito` → `INT` | `isUsingDeposito = 0` |

<a id="kolom-isusinghaircut"></a>
## isUsingHairCut

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOProductType_TM.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005176 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamRepoValidate:644](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamRepoValidate.sql#L644) | WHERE / P11 / COMPARISON | `isUsingHairCut` : `BIT` | `1` : `int` | `isUsingHairCut` → `INT` | `isUsingHairCut = 1` |
| IC-005189 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamRepoValidate:712](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamRepoValidate.sql#L712) | WHERE / P11 / COMPARISON | `tm.isUsingHairCut` : `BIT` | `1` : `int` | `tm.isUsingHairCut` → `INT` | `tm.isUsingHairCut = 1` |
| IC-005190 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamRepoValidate:735](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamRepoValidate.sql#L735) | WHERE / P11 / COMPARISON | `tm.isUsingHairCut` : `BIT` | `1` : `int` | `tm.isUsingHairCut` → `INT` | `tm.isUsingHairCut = 1` |
