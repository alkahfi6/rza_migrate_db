# Retail UAT.dbo.TRSLastCounterParameter_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 2 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-lastcounter"></a>
## LastCounter

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSLastCounterParameter_TR.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012028 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSSaveParamThreshold:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamThreshold.sql#L66) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `LastCounter` : `INT` | `@nCounter` : `bigint` | `LastCounter` → `bigint` | `@nCounter = LastCounter` |
| IC-012031 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSSaveParamThreshold:78](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamThreshold.sql#L78) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `LastCounter` : `INT` | `@nCounterSandi` : `bigint` | `LastCounter` → `bigint` | `@nCounterSandi = LastCounter` |
| IC-012030 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveParamThreshold:75](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamThreshold.sql#L75) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `LastCounter` : `INT` | `@nCounter` : `bigint` | `@nCounter` → `INT` | `LastCounter = @nCounter` |
| IC-012033 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveParamThreshold:87](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamThreshold.sql#L87) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `LastCounter` : `INT` | `@nCounterSandi` : `bigint` | `@nCounterSandi` → `INT` | `LastCounter = @nCounterSandi` |
