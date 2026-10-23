# Retail UAT.dbo.TRSPostingGLBond_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 14 lokasi source, 4 procedure, 1 kolom. **A:** 14 lokasi pada 4 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-tranid"></a>
## TranId

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSPostingGLBond_TH.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 14 lokasi / 4 SP. **A:** 14 lokasi / 4 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003187 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.SettlementPopulateHasilRevenue:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SettlementPopulateHasilRevenue.sql#L27) | JOIN_ON / P01 / COMPARISON | `th.TranId` : `VARCHAR(20)` | `lg.BulkId` : `BIGINT` | `th.TranId` → `BIGINT` | `lg.BulkId = th.TranId` |
| IC-003189 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.SettlementPopulateHasilRevenue:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SettlementPopulateHasilRevenue.sql#L53) | JOIN_ON / P01 / COMPARISON | `th.TranId` : `VARCHAR(20)` | `c.DealId` : `BIGINT` | `th.TranId` → `BIGINT` | `c.DealId = th.TranId` |
| IC-004209 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:651](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L651) | JOIN_ON / P01 / COMPARISON | `tt.[TranId]` : `VARCHAR(20)` | `tl.[BulkId]` : `BIGINT` | `tt.[TranId]` → `BIGINT` | `tl.[BulkId] = tt.[TranId]` |
| IC-004210 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:685](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L685) | JOIN_ON / P01 / COMPARISON | `tt.[TranId]` : `VARCHAR(20)` | `ins.[DealId]` : `BIGINT` | `tt.[TranId]` → `BIGINT` | `ins.[DealId] = tt.[TranId]` |
| IC-004215 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:715](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L715) | JOIN_ON / P01 / COMPARISON | `h.[TranId]` : `VARCHAR(20)` | `d.[DealId]` : `BIGINT` | `h.[TranId]` → `BIGINT` | `d.[DealId] = h.[TranId]` |
| IC-004220 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:778](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L778) | JOIN_ON / P01 / COMPARISON | `th.[TranId]` : `VARCHAR(20)` | `tl.[BulkId]` : `BIGINT` | `th.[TranId]` → `BIGINT` | `tl.[BulkId] = th.[TranId]` |
| IC-004221 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:819](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L819) | JOIN_ON / P01 / COMPARISON | `tt.[TranId]` : `VARCHAR(20)` | `tl.[BulkId]` : `BIGINT` | `tt.[TranId]` → `BIGINT` | `tl.[BulkId] = tt.[TranId]` |
| IC-004222 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:853](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L853) | JOIN_ON / P01 / COMPARISON | `tt.[TranId]` : `VARCHAR(20)` | `ins.[DealId]` : `BIGINT` | `tt.[TranId]` → `BIGINT` | `ins.[DealId] = tt.[TranId]` |
| IC-007921 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRPTRevenueOBLforTrops:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforTrops.sql#L70) | JOIN_ON / P01 / COMPARISON | `th.TranId` : `VARCHAR(20)` | `tm.DealId` : `BIGINT` | `th.TranId` → `BIGINT` | `th.TranId = tm.DealId` |
| IC-007925 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRPTRevenueOBLforTrops:120](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforTrops.sql#L120) | JOIN_ON / P01 / COMPARISON | `h.TranId` : `VARCHAR(20)` | `m.DealId` : `BIGINT` | `h.TranId` → `BIGINT` | `m.DealId = h.TranId` |
| IC-007933 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRPTRevenueOBLforTrops:267](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforTrops.sql#L267) | JOIN_ON / P01 / COMPARISON | `h.TranId` : `VARCHAR(20)` | `m.DealId` : `BIGINT` | `h.TranId` → `BIGINT` | `m.DealId = h.TranId` |
| IC-007937 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRPTRevenueOBLforTrops:298](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforTrops.sql#L298) | JOIN_ON / P01 / COMPARISON | `h.TranId` : `VARCHAR(20)` | `m.DealId` : `BIGINT` | `h.TranId` → `BIGINT` | `h.TranId = m.DealId` |
| IC-007946 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRPTRevenueOBLforTrops:374](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforTrops.sql#L374) | JOIN_ON / P01 / COMPARISON | `h.TranId` : `VARCHAR(20)` | `m.DealId` : `BIGINT` | `h.TranId` → `BIGINT` | `m.DealId = h.TranId` |
| IC-007957 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRPTRevenueOBLforWM:126](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforWM.sql#L126) | JOIN_ON / P01 / COMPARISON | `th.TranId` : `VARCHAR(20)` | `tm.DealId` : `BIGINT` | `th.TranId` → `BIGINT` | `tm.DealId = th.TranId` |
