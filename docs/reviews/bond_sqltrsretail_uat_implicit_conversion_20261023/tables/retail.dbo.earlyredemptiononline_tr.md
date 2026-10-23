# Retail UAT.dbo.EarlyRedemptionOnline_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 5 lokasi source, 2 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-nominal"></a>
## Nominal

**Tipe dasar:** `DECIMAL(21,0)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.EarlyRedemptionOnline_TR.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000975 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.ORISBNIPopulateEarlyRedemption:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateEarlyRedemption.sql#L165) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `SUM(ISNULL(er.Nominal,0))` : `DECIMAL (aggregate; facets unknown)` | `0` : `int` | `0` → `DECIMAL (aggregate; facets unknown)` | `ISNULL(SUM(ISNULL(er.Nominal,0)),0)` |
| IC-000976 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.ORISBNIPopulateEarlyRedemption:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateEarlyRedemption.sql#L165) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `er.Nominal` : `DECIMAL(21,0)` | `0` : `int` | `0` → `DECIMAL(21,0)` | `ISNULL(er.Nominal,0)` |
| IC-001009 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.ORISBNISubmitEarlyRedemption:293](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNISubmitEarlyRedemption.sql#L293) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `SUM(ISNULL(er.Nominal,0))` : `DECIMAL (aggregate; facets unknown)` | `0` : `int` | `0` → `DECIMAL (aggregate; facets unknown)` | `ISNULL(SUM(ISNULL(er.Nominal,0)),0)` |
| IC-001010 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.ORISBNISubmitEarlyRedemption:293](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNISubmitEarlyRedemption.sql#L293) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `er.Nominal` : `DECIMAL(21,0)` | `0` : `int` | `0` → `DECIMAL(21,0)` | `ISNULL(er.Nominal,0)` |
| IC-001021 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.ORISBNISubmitEarlyRedemption:491](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNISubmitEarlyRedemption.sql#L491) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.earlyredemptiononline_tr.Nominal` : `DECIMAL(21,0)` | `NominalRedemption` : `MONEY` | `NominalRedemption` → `DECIMAL(21,0)` | `retail.dbo.earlyredemptiononline_tr.Nominal <- NominalRedemption` |
