# Retail UAT.dbo.TRSCustomerRepoFunded_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 10 lokasi source, 4 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-gmra"></a>
## GMRA

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSCustomerRepoFunded_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 10 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN_EXPRESSION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001960 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.RepoUploadValidation:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoUploadValidation.sql#L165) | WHERE / P11 / COMPARISON | `ISNULL(tm.GMRA,0)` : `BIT` | `0` : `int` | `ISNULL(tm.GMRA,0)` → `INT` | `ISNULL(tm.GMRA,0) = 0` |
| IC-005400 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateEditCustomerAuthorization:455](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEditCustomerAuthorization.sql#L455) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ISNULL(crp.GMRA,0)` : `BIT` | `0` : `int` | `ISNULL(crp.GMRA,0)` → `INT` | `WHEN 0 THEN 'NO'` |
| IC-005401 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateEditCustomerAuthorization:456](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEditCustomerAuthorization.sql#L456) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ISNULL(crp.GMRA,0)` : `BIT` | `1` : `int` | `ISNULL(crp.GMRA,0)` → `INT` | `WHEN 1 THEN 'YES'` |
| IC-001961 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.RepoUploadValidation:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoUploadValidation.sql#L165) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.GMRA` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(tm.GMRA,0)` |
| IC-005399 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateEditCustomerAuthorization:454](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEditCustomerAuthorization.sql#L454) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `crp.GMRA` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(crp.GMRA,0)` |
| IC-013472 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusTreasuryCustomer_TM:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L69) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `crp.GMRA` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(crp.GMRA,0)` |
| IC-013483 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusTreasuryCustomer_TM:129](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L129) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `crp.GMRA` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(crp.GMRA,0)` |
| IC-013494 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusTreasuryCustomer_TM:184](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L184) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `crp.GMRA` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(crp.GMRA,0)` |
| IC-013560 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListTreasuryCustomer_TM_Original:275](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListTreasuryCustomer_TM_Original.sql#L275) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `crp.GMRA` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(crp.GMRA,0)` |
| IC-013566 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListTreasuryCustomer_TM_Original:449](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListTreasuryCustomer_TM_Original.sql#L449) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `crp.GMRA` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(crp.GMRA,0)` |
