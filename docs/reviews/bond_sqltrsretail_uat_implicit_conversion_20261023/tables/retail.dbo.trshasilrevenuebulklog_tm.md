# Retail UAT.dbo.TRSHasilRevenueBulkLog_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 2 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-bulkid"></a>
## BulkId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSHasilRevenueBulkLog_TM.sql#L2). **Existing leading index:** TRSHasilRevenueBulkLog_TM_CN1.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003187 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SettlementPopulateHasilRevenue:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SettlementPopulateHasilRevenue.sql#L27) | JOIN_ON / P01 / COMPARISON | `lg.BulkId` : `BIGINT` | `th.TranId` : `VARCHAR(20)` | `th.TranId` → `BIGINT` | `lg.BulkId = th.TranId` |
| IC-004209 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSHasilRevenueBondJurnalGL:651](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L651) | JOIN_ON / P01 / COMPARISON | `tl.[BulkId]` : `BIGINT` | `tt.[TranId]` : `VARCHAR(20)` | `tt.[TranId]` → `BIGINT` | `tl.[BulkId] = tt.[TranId]` |
| IC-004220 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSHasilRevenueBondJurnalGL:778](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L778) | JOIN_ON / P01 / COMPARISON | `tl.[BulkId]` : `BIGINT` | `th.[TranId]` : `VARCHAR(20)` | `th.[TranId]` → `BIGINT` | `tl.[BulkId] = th.[TranId]` |
| IC-004221 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSHasilRevenueBondJurnalGL:819](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L819) | JOIN_ON / P01 / COMPARISON | `tl.[BulkId]` : `BIGINT` | `tt.[TranId]` : `VARCHAR(20)` | `tt.[TranId]` → `BIGINT` | `tl.[BulkId] = tt.[TranId]` |
