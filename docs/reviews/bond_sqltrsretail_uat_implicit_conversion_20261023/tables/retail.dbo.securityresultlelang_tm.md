# Retail UAT.dbo.SecurityResultLelang_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 5 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-isexecuted"></a>
## IsExecuted

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityResultLelang_TM.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003058 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.SecurityPopulateResultLelangProduct:15](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateResultLelangProduct.sql#L15) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ISNULL(tt.IsExecuted,1)` : `BIT` | `1` : `int` | `ISNULL(tt.IsExecuted,1)` → `INT` | `WHEN 1 THEN 0` |
| IC-003059 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.SecurityPopulateResultLelangProduct:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateResultLelangProduct.sql#L16) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ISNULL(tt.IsExecuted,1)` : `BIT` | `0` : `int` | `ISNULL(tt.IsExecuted,1)` → `INT` | `WHEN 0 THEN 1` |
| IC-003061 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SecurityPopulateResultLelangProduct:67](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateResultLelangProduct.sql#L67) | WHERE / P11 / COMPARISON | `IsExecuted` : `BIT` | `1` : `int` | `IsExecuted` → `INT` | `IsExecuted = 1` |
| IC-003056 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SecurityPopulateResultLelangProduct:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateResultLelangProduct.sql#L13) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tt.IsExecuted` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(tt.IsExecuted,0)` |
| IC-003057 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SecurityPopulateResultLelangProduct:14](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateResultLelangProduct.sql#L14) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tt.IsExecuted` : `BIT` | `1` : `int` | `1` → `BIT` | `ISNULL(tt.IsExecuted,1)` |
