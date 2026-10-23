# Retail UAT.dbo.TRSCustomerRepoFundedPending_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 6 lokasi source, 2 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-gmra"></a>
## GMRA

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSCustomerRepoFundedPending_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005394 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateEditCustomerAuthorization:359](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEditCustomerAuthorization.sql#L359) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ISNULL(crp.GMRA,0)` : `BIT` | `0` : `int` | `ISNULL(crp.GMRA,0)` → `INT` | `WHEN 0 THEN 'NO'` |
| IC-005395 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateEditCustomerAuthorization:360](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEditCustomerAuthorization.sql#L360) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ISNULL(crp.GMRA,0)` : `BIT` | `1` : `int` | `ISNULL(crp.GMRA,0)` → `INT` | `WHEN 1 THEN 'YES'` |
| IC-008547 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSResponseBONDCPYFromMX:652](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L652) | WHERE / P11 / COMPARISON | `tmp.GMRA` : `BIT` | `1` : `int` | `tmp.GMRA` → `INT` | `tmp.GMRA = 1` |
| IC-008548 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSResponseBONDCPYFromMX:661](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L661) | WHERE / P11 / COMPARISON | `tmp.GMRA` : `BIT` | `0` : `int` | `tmp.GMRA` → `INT` | `tmp.GMRA = 0` |
| IC-008549 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSResponseBONDCPYFromMX:681](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L681) | WHERE / P11 / COMPARISON | `tmp.GMRA` : `BIT` | `1` : `int` | `tmp.GMRA` → `INT` | `tmp.GMRA = 1` |
| IC-005393 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateEditCustomerAuthorization:358](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEditCustomerAuthorization.sql#L358) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `crp.GMRA` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(crp.GMRA,0)` |
