# Retail UAT.dbo.TRSParamIssuer_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-processcode"></a>
## ProcessCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamIssuer_TH.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005006 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamIssuerGetAuthDetail:7](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamIssuerGetAuthDetail.sql#L7) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ProcessCode` : `CHAR(1)` | `'I'` : `varchar` | `ProcessCode` → `VARCHAR` | `WHEN 'I' THEN 'INSERT'` |
| IC-005007 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamIssuerGetAuthDetail:8](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamIssuerGetAuthDetail.sql#L8) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ProcessCode` : `CHAR(1)` | `'O'` : `varchar` | `ProcessCode` → `VARCHAR` | `WHEN 'O' THEN 'OLD_DATA'` |
| IC-005008 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamIssuerGetAuthDetail:9](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamIssuerGetAuthDetail.sql#L9) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ProcessCode` : `CHAR(1)` | `'N'` : `varchar` | `ProcessCode` → `VARCHAR` | `WHEN 'N' THEN 'NEW_DATA'` |
| IC-005009 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamIssuerGetAuthDetail:10](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamIssuerGetAuthDetail.sql#L10) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ProcessCode` : `CHAR(1)` | `'D'` : `varchar` | `ProcessCode` → `VARCHAR` | `WHEN 'D' THEN 'DELETE'` |
