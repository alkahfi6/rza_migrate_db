# Retail UAT.dbo.TRSBondsPrice_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 8 lokasi source, 3 procedure, 5 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-bidsize"></a>
## Bidsize

**Tipe dasar:** `FLOAT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBondsPrice_TH.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003397 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceDelete:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceDelete.sql#L27) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbondsprice_th.Bidsize` : `FLOAT` | `0` : `int` | `0` → `FLOAT` | `retail.dbo.trsbondsprice_th.Bidsize <- 0` |
| IC-003405 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsert:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsert.sql#L85) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbondsprice_th.Bidsize` : `FLOAT` | `0` : `int` | `0` → `FLOAT` | `retail.dbo.trsbondsprice_th.Bidsize <- 0` |

<a id="kolom-changeby"></a>
## ChangeBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBondsPrice_TH.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006653 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessApprParamSpread:329](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L329) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbondsprice_th.ChangeBy` : `INT` | `@pnNik` : `bigint` | `@pnNik` → `INT` | `retail.dbo.trsbondsprice_th.ChangeBy <- @pnNik` |

<a id="kolom-offersize"></a>
## Offersize

**Tipe dasar:** `FLOAT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBondsPrice_TH.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003398 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceDelete:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceDelete.sql#L27) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbondsprice_th.Offersize` : `FLOAT` | `0` : `int` | `0` → `FLOAT` | `retail.dbo.trsbondsprice_th.Offersize <- 0` |
| IC-003406 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsert:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsert.sql#L85) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbondsprice_th.Offersize` : `FLOAT` | `0` : `int` | `0` → `FLOAT` | `retail.dbo.trsbondsprice_th.Offersize <- 0` |

<a id="kolom-pricestatus"></a>
## PriceStatus

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBondsPrice_TH.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003399 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceDelete:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceDelete.sql#L27) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbondsprice_th.PriceStatus` : `SMALLINT` | `'4'` : `varchar` | `'4'` → `SMALLINT` | `retail.dbo.trsbondsprice_th.PriceStatus <- '4'` |

<a id="kolom-recordingdate-date"></a>
## Recordingdate_date

**Tipe dasar:** `CHAR(25)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBondsPrice_TH.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003400 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceDelete:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceDelete.sql#L27) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbondsprice_th.Recordingdate_date` : `CHAR(25)` | `getdate()` : `datetime` | `getdate()` → `CHAR(25)` | `retail.dbo.trsbondsprice_th.Recordingdate_date <- getdate()` |
| IC-003407 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsert:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsert.sql#L85) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbondsprice_th.Recordingdate_date` : `CHAR(25)` | `getdate()` : `datetime` | `getdate()` → `CHAR(25)` | `retail.dbo.trsbondsprice_th.Recordingdate_date <- getdate()` |
