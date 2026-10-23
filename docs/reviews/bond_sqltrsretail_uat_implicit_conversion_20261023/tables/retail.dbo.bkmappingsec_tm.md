# Retail UAT.dbo.BKMappingSec_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 1 procedure, 1 kolom. **A:** 2 lokasi pada 1 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-murexlabel"></a>
## MurexLabel

**Tipe dasar:** `VARCHAR(40)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.BKMappingSec_TM.sql#L4). **Existing leading index:** BKMappingSec_TM_CN1.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 2 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000076 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.BKPopulateSummaryHolding:281](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BKPopulateSummaryHolding.sql#L281) | JOIN_ON / P02 / COMPARISON | `map.MurexLabel` : `VARCHAR(40)` | `mx.SECID` : `NVARCHAR(15)` | `map.MurexLabel` → `NVARCHAR` | `map.MurexLabel = mx.SECID` |
| IC-000079 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.BKPopulateSummaryHolding:291](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BKPopulateSummaryHolding.sql#L291) | JOIN_ON / P02 / COMPARISON | `map.MurexLabel` : `VARCHAR(40)` | `mx.SECID` : `NVARCHAR(15)` | `map.MurexLabel` → `NVARCHAR` | `map.MurexLabel = mx.SECID` |
