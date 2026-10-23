# SQL_REPLICATE.dbo.MXCUSTODY — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 7 lokasi source, 2 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cost-center"></a>
## COST_CENTER

**Tipe dasar:** `NVARCHAR(5)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.MXCUSTODY.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000070 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.BKPopulateMXCustodian:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BKPopulateMXCustodian.sql#L24) | WHERE / P04 / COMPARISON | `COST_CENTER` : `NVARCHAR(5)` | `@cCostSyar` : `VARCHAR(20)` | `@cCostSyar` → `NVARCHAR` | `COST_CENTER = @cCostSyar` |
| IC-000071 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.BKPopulateMXCustodian:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BKPopulateMXCustodian.sql#L26) | WHERE / P04 / COMPARISON | `COST_CENTER` : `NVARCHAR(5)` | `@cCostSyar` : `VARCHAR(20)` | `@cCostSyar` → `NVARCHAR` | `COST_CENTER != @cCostSyar` |

<a id="kolom-secid"></a>
## SECID

**Tipe dasar:** `NVARCHAR(15)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.MXCUSTODY.sql#L3). **Existing leading index:** MXCUSTODY_CN1.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.
- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000076 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.BKPopulateSummaryHolding:281](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BKPopulateSummaryHolding.sql#L281) | JOIN_ON / P02 / COMPARISON | `mx.SECID` : `NVARCHAR(15)` | `map.MurexLabel` : `VARCHAR(40)` | `map.MurexLabel` → `NVARCHAR` | `map.MurexLabel = mx.SECID` |
| IC-000078 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.BKPopulateSummaryHolding:288](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BKPopulateSummaryHolding.sql#L288) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `mx.SECID` : `NVARCHAR(15)` | `@cMsg + '- '` : `VARCHAR (combined result; facets unknown)` | `@cMsg + '- '` → `NVARCHAR` | `@cMsg + '- '+mx.SECID` |
| IC-000079 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.BKPopulateSummaryHolding:291](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BKPopulateSummaryHolding.sql#L291) | JOIN_ON / P02 / COMPARISON | `mx.SECID` : `NVARCHAR(15)` | `map.MurexLabel` : `VARCHAR(40)` | `map.MurexLabel` → `NVARCHAR` | `map.MurexLabel = mx.SECID` |

<a id="kolom-tdprintamt"></a>
## TDPRINTAMT

**Tipe dasar:** `DECIMAL(25,2)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.MXCUSTODY.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000077 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.BKPopulateSummaryHolding:283](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BKPopulateSummaryHolding.sql#L283) | WHERE / P11 / COMPARISON | `TDPRINTAMT` : `DECIMAL(25,2)` | `0` : `int` | `0` → `DECIMAL` | `TDPRINTAMT > 0` |
| IC-000080 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.BKPopulateSummaryHolding:293](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BKPopulateSummaryHolding.sql#L293) | WHERE / P11 / COMPARISON | `TDPRINTAMT` : `DECIMAL(25,2)` | `0` : `int` | `0` → `DECIMAL` | `TDPRINTAMT > 0` |
