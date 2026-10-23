# SQL_REPLICATE.dbo.MISCurrencyRateHistory — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-birateexchange"></a>
## BIRateExchange

**Tipe dasar:** `DECIMAL(28,4)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.MISCurrencyRateHistory.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004160 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSHasilRevenueBondJurnalGL:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L137) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ISNULL(sct.BiRateExchange, 1)` : `DECIMAL(28,4)` | `kss.[BIRate]` : `MONEY` | `ISNULL(sct.BiRateExchange, 1)` → `MONEY` | `kss.[BIRate] = ISNULL(sct.BiRateExchange, 1)` |
| IC-004161 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSHasilRevenueBondJurnalGL:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L137) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sct.BiRateExchange` : `DECIMAL(28,4)` | `1` : `int` | `1` → `DECIMAL(28,4)` | `ISNULL(sct.BiRateExchange, 1)` |
