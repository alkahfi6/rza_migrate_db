# Retail UAT.dbo.upload_sid_sre_temp — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 7 lokasi source, 1 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cif-id"></a>
## cif_id

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.upload_sid_sre_temp.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002523 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulateUploadSRESID:19](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateUploadSRESID.sql#L19) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `usst.cif_id` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `ISNULL(usst.cif_id, 0)` |

<a id="kolom-jenis-identitas"></a>
## jenis_identitas

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.upload_sid_sre_temp.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002524 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateUploadSRESID:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateUploadSRESID.sql#L23) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `usst.jenis_identitas` : `CHAR(1)` | `'1'` : `varchar` | `usst.jenis_identitas` → `VARCHAR` | `WHEN '1' THEN 'KTP'` |
| IC-002525 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateUploadSRESID:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateUploadSRESID.sql#L24) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `usst.jenis_identitas` : `CHAR(1)` | `'2'` : `varchar` | `usst.jenis_identitas` → `VARCHAR` | `WHEN '2' THEN 'SIM'` |
| IC-002526 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateUploadSRESID:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateUploadSRESID.sql#L25) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `usst.jenis_identitas` : `CHAR(1)` | `'3'` : `varchar` | `usst.jenis_identitas` → `VARCHAR` | `WHEN '3' THEN 'Lain-lain'` |

<a id="kolom-jenis-kelamin"></a>
## jenis_kelamin

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.upload_sid_sre_temp.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002527 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateUploadSRESID:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateUploadSRESID.sql#L30) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `usst.jenis_kelamin` : `CHAR(1)` | `'1'` : `varchar` | `usst.jenis_kelamin` → `VARCHAR` | `WHEN '1' THEN 'Laki-laki'` |
| IC-002528 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateUploadSRESID:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateUploadSRESID.sql#L31) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `usst.jenis_kelamin` : `CHAR(1)` | `'2'` : `varchar` | `usst.jenis_kelamin` → `VARCHAR` | `WHEN '2' THEN 'Wanita'` |
| IC-002529 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateUploadSRESID:32](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateUploadSRESID.sql#L32) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `usst.jenis_kelamin` : `CHAR(1)` | `'3'` : `varchar` | `usst.jenis_kelamin` → `VARCHAR` | `WHEN '3' THEN 'Lain-lain'` |
