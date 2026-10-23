# Retail UAT.dbo.TRSParamTradeSettleDate_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 1 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-isactive"></a>
## IsActive

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamTradeSettleDate_TR.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004468 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSMaintenanceParamHargaModal:207](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L207) | WHERE / P11 / COMPARISON | `sd.IsActive` : `BIT` | `1` : `int` | `sd.IsActive` → `INT` | `sd.IsActive = 1` |
| IC-004470 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSMaintenanceParamHargaModal:308](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L308) | WHERE / P11 / COMPARISON | `sd.IsActive` : `BIT` | `1` : `int` | `sd.IsActive` → `INT` | `sd.IsActive = 1` |

<a id="kolom-tradedate"></a>
## TradeDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamTradeSettleDate_TR.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004467 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSMaintenanceParamHargaModal:205](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L205) | JOIN_ON / P04 / COMPARISON | `sd.TradeDate` : `DATETIME` | `CONVERT(VARCHAR(8), ph.TransactionDate, 112)` : `VARCHAR(8)` | `CONVERT(VARCHAR(8), ph.TransactionDate, 112)` → `DATETIME` | `CONVERT(VARCHAR(8), ph.TransactionDate, 112) = sd.TradeDate` |
| IC-004469 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSMaintenanceParamHargaModal:306](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L306) | JOIN_ON / P04 / COMPARISON | `sd.TradeDate` : `DATETIME` | `CONVERT(VARCHAR(8), hm.TransactionDate, 112)` : `VARCHAR(8)` | `CONVERT(VARCHAR(8), hm.TransactionDate, 112)` → `DATETIME` | `CONVERT(VARCHAR(8), hm.TransactionDate, 112) = sd.TradeDate` |
