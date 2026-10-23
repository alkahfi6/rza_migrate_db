# Retail UAT.dbo.TRSPendingRevenue_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 10 lokasi source, 2 procedure, 3 kolom. **A:** 1 lokasi pada 1 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-kodecabangtrx"></a>
## KodeCabangTRX

**Tipe dasar:** `VARCHAR(5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSPendingRevenue_TT.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004237 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L94) | JOIN_ON / P01 / COMPARISON | `s.KodeCabangTRX` : `VARCHAR(5)` | `b.[office_id]` : `INT` | `s.KodeCabangTRX` → `INT` | `s.KodeCabangTRX = b.[office_id]` |
| IC-008610 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSResponseBondFromMX:611](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L611) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `s.KodeCabangTRX` : `VARCHAR(5)` | `retail.dbo.trshasilrevenue_th.KodeCabangTRX` : `INT` | `s.KodeCabangTRX` → `INT` | `retail.dbo.trshasilrevenue_th.KodeCabangTRX <- s.KodeCabangTRX` |

<a id="kolom-secid"></a>
## SecId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSPendingRevenue_TT.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008607 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBondFromMX:598](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L598) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trspendingrevenue_tt.SecId` : `BIGINT` | `s.[SecId]` : `INT` | `s.[SecId]` → `BIGINT` | `retail.dbo.trspendingrevenue_tt.SecId <- s.[SecId]` |
| IC-008615 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBondFromMX:634](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L634) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trspendingrevenue_tt.SecId` : `BIGINT` | `s.[SecId]` : `INT` | `s.[SecId]` → `BIGINT` | `retail.dbo.trspendingrevenue_tt.SecId <- s.[SecId]` |

<a id="kolom-isdatacomplete"></a>
## isDataComplete

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSPendingRevenue_TT.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004238 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:97](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L97) | WHERE / P11 / COMPARISON | `s.isDataComplete` : `BIT` | `0` : `int` | `s.isDataComplete` → `INT` | `s.isDataComplete = 0` |
| IC-008608 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSResponseBondFromMX:609](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L609) | WHERE / P11 / COMPARISON | `rv.isDataComplete` : `BIT` | `1` : `int` | `rv.isDataComplete` → `INT` | `rv.isDataComplete = 1` |
| IC-008612 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSResponseBondFromMX:623](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L623) | WHERE / P11 / COMPARISON | `isDataComplete` : `BIT` | `0` : `int` | `isDataComplete` → `INT` | `isDataComplete = 0` |
| IC-008614 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSResponseBondFromMX:631](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L631) | WHERE / P11 / COMPARISON | `isDataComplete` : `BIT` | `0` : `int` | `isDataComplete` → `INT` | `isDataComplete = 0` |
| IC-004239 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:100](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L100) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `isDataComplete` : `BIT` | `1` : `int` | `1` → `BIT` | `isDataComplete = 1` |
| IC-008613 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBondFromMX:627](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L627) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `isDataComplete` : `BIT` | `1` : `int` | `1` → `BIT` | `isDataComplete = 1` |
