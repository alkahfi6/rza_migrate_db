# Retail UAT.dbo.TRSOnlineDocsLinkAppDetail_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 11 lokasi source, 2 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-processcd"></a>
## ProcessCd

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSOnlineDocsLinkAppDetail_TH.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 11 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005368 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDetailAppUploadDocsLink:11](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDetailAppUploadDocsLink.sql#L11) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ProcessCd` : `CHAR(1)` | `'I'` : `varchar` | `ProcessCd` → `VARCHAR` | `WHEN 'I' THEN 'INSERT'` |
| IC-005369 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDetailAppUploadDocsLink:12](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDetailAppUploadDocsLink.sql#L12) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ProcessCd` : `CHAR(1)` | `'U'` : `varchar` | `ProcessCd` → `VARCHAR` | `WHEN 'U' THEN 'UPDATE DATA'` |
| IC-005370 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDetailAppUploadDocsLink:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDetailAppUploadDocsLink.sql#L13) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ProcessCd` : `CHAR(1)` | `'O'` : `varchar` | `ProcessCd` → `VARCHAR` | `WHEN 'O' THEN 'OLD DATA'` |
| IC-005371 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDetailAppUploadDocsLink:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDetailAppUploadDocsLink.sql#L18) | WHERE / P12 / IN_VALUE_TYPE | `ProcessCd` : `CHAR(1)` | `'O'` : `varchar` | `ProcessCd` → `VARCHAR` | `ProcessCd IN ('O', 'D')` |
| IC-005372 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDetailAppUploadDocsLink:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDetailAppUploadDocsLink.sql#L27) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ProcessCd` : `CHAR(1)` | `'I'` : `varchar` | `ProcessCd` → `VARCHAR` | `WHEN 'I' THEN 'INSERT'` |
| IC-005373 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDetailAppUploadDocsLink:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDetailAppUploadDocsLink.sql#L28) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ProcessCd` : `CHAR(1)` | `'U'` : `varchar` | `ProcessCd` → `VARCHAR` | `WHEN 'U' THEN 'UPDATE DATA'` |
| IC-005374 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDetailAppUploadDocsLink:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDetailAppUploadDocsLink.sql#L29) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ProcessCd` : `CHAR(1)` | `'O'` : `varchar` | `ProcessCd` → `VARCHAR` | `WHEN 'O' THEN 'OLD DATA'` |
| IC-005375 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDetailAppUploadDocsLink:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDetailAppUploadDocsLink.sql#L34) | WHERE / P12 / IN_VALUE_TYPE | `ProcessCd` : `CHAR(1)` | `'I'` : `varchar` | `ProcessCd` → `VARCHAR` | `ProcessCd IN ('I', 'U')` |
| IC-006655 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprovalDocsLink:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalDocsLink.sql#L26) | WHERE / P12 / COMPARISON | `ProcessCd` : `CHAR(1)` | `'I'` : `varchar` | `ProcessCd` → `VARCHAR` | `ProcessCd = 'I'` |
| IC-006656 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprovalDocsLink:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalDocsLink.sql#L36) | WHERE / P12 / COMPARISON | `th.ProcessCd` : `CHAR(1)` | `'U'` : `varchar` | `th.ProcessCd` → `VARCHAR` | `th.ProcessCd = 'U'` |
| IC-006657 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprovalDocsLink:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalDocsLink.sql#L44) | WHERE / P12 / COMPARISON | `th.ProcessCd` : `CHAR(1)` | `'D'` : `varchar` | `th.ProcessCd` → `VARCHAR` | `th.ProcessCd = 'D'` |
