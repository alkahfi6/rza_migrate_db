# SQL_REPLICATE.dbo.BIS4_INSTRUMENT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 22 lokasi source, 2 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 12 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-code"></a>
## code

**Tipe dasar:** `NVARCHAR(128)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.BIS4_INSTRUMENT.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 13 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.
- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006087 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateSettlementBIS4:502](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L502) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `mst.InstrumentCode` : `NVARCHAR(128)` | `retail.dbo.trssettlementbis4_temp.CodeInstrument` : `VARCHAR(500)` | `mst.InstrumentCode` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.CodeInstrument <- mst.InstrumentCode` |
| IC-006097 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateSettlementBIS4:682](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L682) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `mst.InstrumentCode` : `NVARCHAR(128)` | `retail.dbo.trssettlementbis4_temp.CodeInstrument` : `VARCHAR(500)` | `mst.InstrumentCode` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.CodeInstrument <- mst.InstrumentCode` |
| IC-006156 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateSettlementBIS4:1151](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1151) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `mst.InstrumentCode` : `NVARCHAR(128)` | `#tmpDispTrx.CodeInstrument` : `[varchar](500)` | `mst.InstrumentCode` → `[varchar](500)` | `#tmpDispTrx.CodeInstrument <- mst.InstrumentCode` |
| IC-006166 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateSettlementBIS4:1289](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1289) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `mst.InstrumentCode` : `NVARCHAR(128)` | `#tmpDispTrx.CodeInstrument` : `[varchar](500)` | `mst.InstrumentCode` → `[varchar](500)` | `#tmpDispTrx.CodeInstrument <- mst.InstrumentCode` |
| IC-006053 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBIS4:319](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L319) | JOIN_ON / P02 / COMPARISON | `ins.InstrumentCode` : `NVARCHAR(128)` | `MapBI.BIS4Code` : `VARCHAR(50)` | `MapBI.BIS4Code` → `NVARCHAR` | `ins.InstrumentCode = MapBI.BIS4Code` |
| IC-006082 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBIS4:486](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L486) | JOIN_ON / P02 / COMPARISON | `ins.InstrumentCode` : `NVARCHAR(128)` | `MapBI.BIS4Code` : `VARCHAR(50)` | `MapBI.BIS4Code` → `NVARCHAR` | `ins.InstrumentCode = MapBI.BIS4Code` |
| IC-006092 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBIS4:670](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L670) | JOIN_ON / P02 / COMPARISON | `mst.InstrumentCode` : `NVARCHAR(128)` | `MapBI.BIS4Code` : `VARCHAR(50)` | `MapBI.BIS4Code` → `NVARCHAR` | `mst.InstrumentCode= MapBI.BIS4Code` |
| IC-006102 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBIS4:849](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L849) | JOIN_ON / P02 / COMPARISON | `mst.InstrumentCode` : `NVARCHAR(128)` | `MapBI.BIS4Code` : `VARCHAR(50)` | `MapBI.BIS4Code` → `NVARCHAR` | `mst.InstrumentCode= MapBI.BIS4Code` |
| IC-006126 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBIS4:1010](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1010) | JOIN_ON / P02 / COMPARISON | `ins.InstrumentCode` : `NVARCHAR(128)` | `MapBI.BIS4Code` : `VARCHAR(50)` | `MapBI.BIS4Code` → `NVARCHAR` | `ins.InstrumentCode = MapBI.BIS4Code` |
| IC-006151 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBIS4:1131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1131) | JOIN_ON / P02 / COMPARISON | `ins.InstrumentCode` : `NVARCHAR(128)` | `MapBI.BIS4Code` : `VARCHAR(50)` | `MapBI.BIS4Code` → `NVARCHAR` | `ins.InstrumentCode = MapBI.BIS4Code` |
| IC-006161 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBIS4:1277](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1277) | JOIN_ON / P02 / COMPARISON | `mst.InstrumentCode` : `NVARCHAR(128)` | `MapBI.BIS4Code` : `VARCHAR(50)` | `MapBI.BIS4Code` → `NVARCHAR` | `mst.InstrumentCode= MapBI.BIS4Code` |
| IC-006172 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBIS4:1408](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1408) | JOIN_ON / P02 / COMPARISON | `mst.InstrumentCode` : `NVARCHAR(128)` | `MapBI.BIS4Code` : `VARCHAR(50)` | `MapBI.BIS4Code` → `NVARCHAR` | `mst.InstrumentCode= MapBI.BIS4Code` |
| IC-006423 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSubInstrumentParam:2](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSubInstrumentParam.sql#L2) | SET_OUTPUT / P07 / UNION_OUTPUT | `right output #1` : `NVARCHAR(128)` | `left output #1` : `varchar` | `left output #1` → `NVARCHAR` | `select '' as 'Value', '' as 'Display' union all select InstrumentCode as 'Value' ,InstrumentName as 'Display' from dbo.BIS4Instrument_v order by Value` |

<a id="kolom-name"></a>
## name

**Tipe dasar:** `NVARCHAR(100)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.BIS4_INSTRUMENT.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006424 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSubInstrumentParam:2](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSubInstrumentParam.sql#L2) | SET_OUTPUT / P07 / UNION_OUTPUT | `right output #2` : `NVARCHAR(100)` | `left output #2` : `varchar` | `left output #2` → `NVARCHAR` | `select '' as 'Value', '' as 'Display' union all select InstrumentCode as 'Value' ,InstrumentName as 'Display' from dbo.BIS4Instrument_v order by Value` |

<a id="kolom-type-code"></a>
## type_code

**Tipe dasar:** `NVARCHAR(100)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.BIS4_INSTRUMENT.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006039 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateSettlementBIS4:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L165) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ins.TypeCode` : `NVARCHAR(100)` | `retail.dbo.trssettlementbis4_temp.ShortCodeInstrument` : `VARCHAR(500)` | `ins.TypeCode` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.ShortCodeInstrument <- ins.TypeCode` |
| IC-006067 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateSettlementBIS4:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L337) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ins.TypeCode` : `NVARCHAR(100)` | `retail.dbo.trssettlementbis4_temp.ShortCodeInstrument` : `VARCHAR(500)` | `ins.TypeCode` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.ShortCodeInstrument <- ins.TypeCode` |
| IC-006089 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateSettlementBIS4:502](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L502) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `mst.TypeCode` : `NVARCHAR(100)` | `retail.dbo.trssettlementbis4_temp.ShortCodeInstrument` : `VARCHAR(500)` | `mst.TypeCode` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.ShortCodeInstrument <- mst.TypeCode` |
| IC-006099 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateSettlementBIS4:682](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L682) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `mst.TypeCode` : `NVARCHAR(100)` | `retail.dbo.trssettlementbis4_temp.ShortCodeInstrument` : `VARCHAR(500)` | `mst.TypeCode` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.ShortCodeInstrument <- mst.TypeCode` |
| IC-006114 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateSettlementBIS4:899](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L899) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ins.TypeCode` : `NVARCHAR(100)` | `#tmpDispTrx.ShortCodeInstrument` : `[varchar](500)` | `ins.TypeCode` → `[varchar](500)` | `#tmpDispTrx.ShortCodeInstrument <- ins.TypeCode` |
| IC-006139 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateSettlementBIS4:1028](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1028) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ins.TypeCode` : `NVARCHAR(100)` | `#tmpDispTrx.ShortCodeInstrument` : `[varchar](500)` | `ins.TypeCode` → `[varchar](500)` | `#tmpDispTrx.ShortCodeInstrument <- ins.TypeCode` |
| IC-006158 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateSettlementBIS4:1151](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1151) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `mst.TypeCode` : `NVARCHAR(100)` | `#tmpDispTrx.ShortCodeInstrument` : `[varchar](500)` | `mst.TypeCode` → `[varchar](500)` | `#tmpDispTrx.ShortCodeInstrument <- mst.TypeCode` |
| IC-006169 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateSettlementBIS4:1289](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1289) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `mst.TypeCode` : `NVARCHAR(100)` | `#tmpDispTrx.ShortCodeInstrument` : `[varchar](500)` | `mst.TypeCode` → `[varchar](500)` | `#tmpDispTrx.ShortCodeInstrument <- mst.TypeCode` |
