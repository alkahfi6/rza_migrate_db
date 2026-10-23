# Retail UAT.dbo.TRSProcessTransaksi_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 7 lokasi source, 4 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-dealid"></a>
## DealId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSProcessTransaksi_TT.sql#L8). **Existing leading index:** TRSProcessTransaksi_TT_CN1.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000081 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.BuybackLogInsert:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BuybackLogInsert.sql#L43) | JOIN_ON / P11 / COMPARISON | `ptt.DealId` : `BIGINT` | `bb.BuyBackID` : `INT` | `bb.BuyBackID` → `BIGINT` | `ptt.DealId = bb.BuyBackID` |
| IC-011458 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptProcessedBuyBack:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptProcessedBuyBack.sql#L43) | JOIN_ON / P11 / COMPARISON | `ptt.DealId` : `BIGINT` | `bb.BuyBackID` : `INT` | `bb.BuyBackID` → `BIGINT` | `ptt.DealId = bb.BuyBackID` |
| IC-011509 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptSettledBuyBack:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettledBuyBack.sql#L69) | JOIN_ON / P11 / COMPARISON | `ptt.DealId` : `BIGINT` | `bb.BuyBackID` : `INT` | `bb.BuyBackID` → `BIGINT` | `ptt.DealId = bb.BuyBackID` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSProcessTransaksi_TT.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-011455 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptProcessedBuyBack:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptProcessedBuyBack.sql#L28) | CASE / P04 / COMPARISON | `ptt.Status` : `INT` | `'1'` : `varchar` | `'1'` → `INT` | `ptt.Status = '1'` |
| IC-011456 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptProcessedBuyBack:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptProcessedBuyBack.sql#L29) | CASE / P04 / COMPARISON | `ptt.Status` : `INT` | `'3'` : `varchar` | `'3'` → `INT` | `ptt.Status = '3'` |
| IC-011457 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptProcessedBuyBack:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptProcessedBuyBack.sql#L30) | CASE / P04 / COMPARISON | `ptt.Status` : `INT` | `'2'` : `varchar` | `'2'` → `INT` | `ptt.Status = '2'` |

<a id="kolom-updatednik"></a>
## UpdatedNIK

**Tipe dasar:** `VARCHAR(10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSProcessTransaksi_TT.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004232 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertProcessTransaksi:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertProcessTransaksi.sql#L46) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsprocesstransaksi_tt.UpdatedNIK` : `VARCHAR(10)` | `@pnNIK` : `INT` | `@pnNIK` → `VARCHAR(10)` | `retail.dbo.trsprocesstransaksi_tt.UpdatedNIK <- @pnNIK` |
