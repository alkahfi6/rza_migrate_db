# Retail UAT.dbo.SecurityResultForTA_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-isexecuted"></a>
## IsExecuted

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityResultForTA_TT.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003074 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SecurityPopulateResultLelangTA:99](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateResultLelangTA.sql#L99) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `tt.IsExecuted` : `BIT` | `1` : `int` | `tt.IsExecuted` → `INT` | `WHEN 1 THEN 0` |
| IC-003075 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SecurityPopulateResultLelangTA:100](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateResultLelangTA.sql#L100) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `tt.IsExecuted` : `BIT` | `0` : `int` | `tt.IsExecuted` → `INT` | `WHEN 0 THEN 1` |
