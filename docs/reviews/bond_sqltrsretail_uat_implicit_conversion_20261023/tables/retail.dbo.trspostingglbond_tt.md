# Retail UAT.dbo.TRSPostingGLBond_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 3 lokasi source, 1 procedure, 2 kolom. **A:** 1 lokasi pada 1 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-tranid"></a>
## TranId

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSPostingGLBond_TT.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004204 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:459](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L459) | JOIN_ON / P01 / COMPARISON | `tt.[TranId]` : `VARCHAR(20)` | `chr.[DealId]` : `BIGINT` | `tt.[TranId]` → `BIGINT` | `chr.[DealId] = tt.[TranId]` |
| IC-004203 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:448](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L448) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trspostingglbond_tt.TranId` : `VARCHAR(20)` | `chr.[DealId]` : `BIGINT` | `chr.[DealId]` → `VARCHAR(20)` | `retail.dbo.trspostingglbond_tt.TranId <- chr.[DealId]` |

<a id="kolom-rate"></a>
## Rate

**Tipe dasar:** `DECIMAL(25,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSPostingGLBond_TT.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004202 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:448](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L448) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trspostingglbond_tt.Rate` : `DECIMAL(25,5)` | `chr.[BIRate]` : `MONEY` | `chr.[BIRate]` → `DECIMAL(25,5)` | `retail.dbo.trspostingglbond_tt.Rate <- chr.[BIRate]` |
