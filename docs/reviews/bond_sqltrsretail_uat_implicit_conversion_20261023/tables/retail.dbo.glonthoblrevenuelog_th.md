# Retail UAT.dbo.GLONTHOblRevenueLog_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 8 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-logid"></a>
## LogId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.GLONTHOblRevenueLog_TH.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004211 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSHasilRevenueBondJurnalGL:704](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L704) | CASE / P04 / COMPARISON | `ISNULL(lg.[LogId], '')` : `BIGINT` | `''` : `varchar` | `''` → `BIGINT` | `ISNULL(lg.[LogId], '') = ''` |
| IC-004212 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSHasilRevenueBondJurnalGL:704](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L704) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `lg.[LogId]` : `BIGINT` | `''` : `varchar` | `''` → `BIGINT` | `ISNULL(lg.[LogId], '')` |
| IC-004213 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSHasilRevenueBondJurnalGL:709](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L709) | CASE / P04 / COMPARISON | `ISNULL(lg.[LogId], '')` : `BIGINT` | `''` : `varchar` | `''` → `BIGINT` | `ISNULL(lg.[LogId], '') = ''` |
| IC-004214 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSHasilRevenueBondJurnalGL:709](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L709) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `lg.[LogId]` : `BIGINT` | `''` : `varchar` | `''` → `BIGINT` | `ISNULL(lg.[LogId], '')` |
| IC-004216 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSHasilRevenueBondJurnalGL:737](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L737) | CASE / P04 / COMPARISON | `ISNULL(lg.[LogId], '')` : `BIGINT` | `''` : `varchar` | `''` → `BIGINT` | `ISNULL(lg.[LogId], '') = ''` |
| IC-004217 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSHasilRevenueBondJurnalGL:737](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L737) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `lg.[LogId]` : `BIGINT` | `''` : `varchar` | `''` → `BIGINT` | `ISNULL(lg.[LogId], '')` |
| IC-004218 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSHasilRevenueBondJurnalGL:742](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L742) | CASE / P04 / COMPARISON | `ISNULL(lg.[LogId], '')` : `BIGINT` | `''` : `varchar` | `''` → `BIGINT` | `ISNULL(lg.[LogId], '') = ''` |
| IC-004219 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSHasilRevenueBondJurnalGL:742](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L742) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `lg.[LogId]` : `BIGINT` | `''` : `varchar` | `''` → `BIGINT` | `ISNULL(lg.[LogId], '')` |
