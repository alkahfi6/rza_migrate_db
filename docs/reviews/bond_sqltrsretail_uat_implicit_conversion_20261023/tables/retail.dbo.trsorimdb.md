# Retail UAT.dbo.TRSORIMDB — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 5 lokasi source, 1 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-kodeintern"></a>
## KodeIntern

**Tipe dasar:** `VARCHAR(5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSORIMDB.sql#L24). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009082 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailPopulateORIOrder:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulateORIOrder.sql#L111) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsorimdb.KodeIntern` : `VARCHAR(5)` | `0` : `int` | `0` → `VARCHAR(5)` | `retail.dbo.trsorimdb.KodeIntern <- 0` |

<a id="kolom-ordernominal"></a>
## OrderNominal

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSORIMDB.sql#L19). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009086 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailPopulateORIOrder:167](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulateORIOrder.sql#L167) | CASE / P11 / COMPARISON | `OrderNominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `OrderNominal > 0` |
| IC-009088 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailPopulateORIOrder:178](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulateORIOrder.sql#L178) | CASE / P11 / COMPARISON | `OrderNominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `OrderNominal > 0` |
| IC-009096 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailPopulateORIOrder:434](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulateORIOrder.sql#L434) | WHERE / P11 / COMPARISON | `OrderNominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `OrderNominal > 0` |

<a id="kolom-priority"></a>
## Priority

**Tipe dasar:** `VARCHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSORIMDB.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009083 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailPopulateORIOrder:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulateORIOrder.sql#L111) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsorimdb.Priority` : `VARCHAR(1)` | `0` : `int` | `0` → `VARCHAR(1)` | `retail.dbo.trsorimdb.Priority <- 0` |
