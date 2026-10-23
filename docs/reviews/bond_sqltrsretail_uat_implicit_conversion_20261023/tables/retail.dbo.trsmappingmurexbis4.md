# Retail UAT.dbo.TRSMappingMurexBIS4 — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 8 lokasi source, 1 procedure, 1 kolom. **A:** 8 lokasi pada 1 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-bis4code"></a>
## BIS4Code

**Tipe dasar:** `VARCHAR(50)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMappingMurexBIS4.sql#L5). **Existing leading index:** TRSMappingMurexBIS4_NN2.

**Frekuensi source:** 8 lokasi / 1 SP. **A:** 8 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006053 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:319](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L319) | JOIN_ON / P02 / COMPARISON | `MapBI.BIS4Code` : `VARCHAR(50)` | `ins.InstrumentCode` : `NVARCHAR(128)` | `MapBI.BIS4Code` → `NVARCHAR` | `ins.InstrumentCode = MapBI.BIS4Code` |
| IC-006082 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:486](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L486) | JOIN_ON / P02 / COMPARISON | `MapBI.BIS4Code` : `VARCHAR(50)` | `ins.InstrumentCode` : `NVARCHAR(128)` | `MapBI.BIS4Code` → `NVARCHAR` | `ins.InstrumentCode = MapBI.BIS4Code` |
| IC-006092 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:670](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L670) | JOIN_ON / P02 / COMPARISON | `MapBI.BIS4Code` : `VARCHAR(50)` | `mst.InstrumentCode` : `NVARCHAR(128)` | `MapBI.BIS4Code` → `NVARCHAR` | `mst.InstrumentCode= MapBI.BIS4Code` |
| IC-006102 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:849](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L849) | JOIN_ON / P02 / COMPARISON | `MapBI.BIS4Code` : `VARCHAR(50)` | `mst.InstrumentCode` : `NVARCHAR(128)` | `MapBI.BIS4Code` → `NVARCHAR` | `mst.InstrumentCode= MapBI.BIS4Code` |
| IC-006126 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:1010](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1010) | JOIN_ON / P02 / COMPARISON | `MapBI.BIS4Code` : `VARCHAR(50)` | `ins.InstrumentCode` : `NVARCHAR(128)` | `MapBI.BIS4Code` → `NVARCHAR` | `ins.InstrumentCode = MapBI.BIS4Code` |
| IC-006151 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:1131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1131) | JOIN_ON / P02 / COMPARISON | `MapBI.BIS4Code` : `VARCHAR(50)` | `ins.InstrumentCode` : `NVARCHAR(128)` | `MapBI.BIS4Code` → `NVARCHAR` | `ins.InstrumentCode = MapBI.BIS4Code` |
| IC-006161 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:1277](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1277) | JOIN_ON / P02 / COMPARISON | `MapBI.BIS4Code` : `VARCHAR(50)` | `mst.InstrumentCode` : `NVARCHAR(128)` | `MapBI.BIS4Code` → `NVARCHAR` | `mst.InstrumentCode= MapBI.BIS4Code` |
| IC-006172 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:1408](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1408) | JOIN_ON / P02 / COMPARISON | `MapBI.BIS4Code` : `VARCHAR(50)` | `mst.InstrumentCode` : `NVARCHAR(128)` | `MapBI.BIS4Code` → `NVARCHAR` | `mst.InstrumentCode= MapBI.BIS4Code` |
