# Retail UAT.dbo.TRSParamBookBld_TL — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 7 lokasi source, 2 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-actionflag"></a>
## ActionFlag

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamBookBld_TL.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005228 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprParamBookbld:6](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamBookbld.sql#L6) | WHERE / P12 / COMPARISON | `ActionFlag` : `CHAR(1)` | `ParamId` : `VARCHAR(20)` | `ActionFlag` → `VARCHAR` | `ParamId = ActionFlag` |
| IC-005229 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprParamBookbld:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamBookbld.sql#L13) | WHERE / P12 / IN_VALUE_TYPE | `ActionFlag` : `CHAR(1)` | `'O'` : `varchar` | `ActionFlag` → `VARCHAR` | `ActionFlag in ('O','D')` |
| IC-005230 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprParamBookbld:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamBookbld.sql#L16) | WHERE / P12 / COMPARISON | `ActionFlag` : `CHAR(1)` | `ParamId` : `VARCHAR(20)` | `ActionFlag` → `VARCHAR` | `ParamId = ActionFlag` |
| IC-005231 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprParamBookbld:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamBookbld.sql#L23) | WHERE / P12 / IN_VALUE_TYPE | `ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `ActionFlag` → `VARCHAR` | `ActionFlag in ('N','I')` |
| IC-006552 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprParamBookbld:11](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamBookbld.sql#L11) | WHERE / P12 / COMPARISON | `ActionFlag` : `CHAR(1)` | `'I'` : `varchar` | `ActionFlag` → `VARCHAR` | `ActionFlag = 'I'` |
| IC-006553 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprParamBookbld:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamBookbld.sql#L18) | WHERE / P12 / COMPARISON | `bl.ActionFlag` : `CHAR(1)` | `'D'` : `varchar` | `bl.ActionFlag` → `VARCHAR` | `bl.ActionFlag = 'D'` |
| IC-006554 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprParamBookbld:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamBookbld.sql#L28) | WHERE / P12 / COMPARISON | `bl.ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `bl.ActionFlag` → `VARCHAR` | `bl.ActionFlag = 'N'` |
