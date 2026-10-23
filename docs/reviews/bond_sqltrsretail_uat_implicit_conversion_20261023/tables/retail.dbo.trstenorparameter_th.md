# Retail UAT.dbo.TRSTenorParameter_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 5 lokasi source, 1 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-creatornik"></a>
## CreatorNIK

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSTenorParameter_TH.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005206 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSParamSaveTenor:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamSaveTenor.sql#L108) | JOIN_ON / P04 / COMPARISON | `isnull(th1.CreatorNIK,'')` : `BIGINT` | `''` : `varchar` | `''` → `BIGINT` | `isnull(th1.CreatorNIK,'') = ''` |
| IC-005207 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamSaveTenor:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamSaveTenor.sql#L108) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `th1.CreatorNIK` : `BIGINT` | `''` : `varchar` | `''` → `BIGINT` | `isnull(th1.CreatorNIK,'')` |

<a id="kolom-statusparam"></a>
## StatusParam

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSTenorParameter_TH.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005203 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSParamSaveTenor:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamSaveTenor.sql#L56) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trstenorparameter_th.StatusParam` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.trstenorparameter_th.StatusParam <- 'A'` |
| IC-005204 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSParamSaveTenor:72](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamSaveTenor.sql#L72) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trstenorparameter_th.StatusParam` : `CHAR(1)` | `'E'` : `varchar` | `'E'` → `CHAR(1)` | `retail.dbo.trstenorparameter_th.StatusParam <- 'E'` |
| IC-005205 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSParamSaveTenor:88](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamSaveTenor.sql#L88) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trstenorparameter_th.StatusParam` : `CHAR(1)` | `'D'` : `varchar` | `'D'` → `CHAR(1)` | `retail.dbo.trstenorparameter_th.StatusParam <- 'D'` |
