# Retail UAT.dbo.TRSHasilRevenue_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 2 procedure, 4 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-bulkid"></a>
## BulkId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSHasilRevenue_TH.sql#L7). **Existing leading index:** TRSHasilRevenue_TH_NN3.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008609 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBondFromMX:611](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L611) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trshasilrevenue_th.BulkId` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `retail.dbo.trshasilrevenue_th.BulkId <- 0` |

<a id="kolom-kodecabangtrx"></a>
## KodeCabangTRX

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSHasilRevenue_TH.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008610 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBondFromMX:611](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L611) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trshasilrevenue_th.KodeCabangTRX` : `INT` | `s.KodeCabangTRX` : `VARCHAR(5)` | `s.KodeCabangTRX` → `INT` | `retail.dbo.trshasilrevenue_th.KodeCabangTRX <- s.KodeCabangTRX` |

<a id="kolom-masterid"></a>
## MasterId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSHasilRevenue_TH.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008611 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBondFromMX:611](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L611) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trshasilrevenue_th.MasterId` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `retail.dbo.trshasilrevenue_th.MasterId <- 0` |

<a id="kolom-trxstatus"></a>
## TrxStatus

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSHasilRevenue_TH.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004259 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:250](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L250) | WHERE / P11 / COMPARISON | `th.[TrxStatus]` : `SMALLINT` | `3` : `int` | `th.[TrxStatus]` → `INT` | `th.[TrxStatus] = 3` |
