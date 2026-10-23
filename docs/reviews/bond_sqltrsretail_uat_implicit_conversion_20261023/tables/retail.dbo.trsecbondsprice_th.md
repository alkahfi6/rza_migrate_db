# Retail UAT.dbo.TRSECBondsPrice_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 5 lokasi source, 2 procedure, 5 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-bidytm"></a>
## Bidytm

**Tipe dasar:** `SMALLMONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECBondsPrice_TH.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003416 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsertEChannel:92](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsertEChannel.sql#L92) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsecbondsprice_th.Bidytm` : `SMALLMONEY` | `@Bidytm` : `money` | `@Bidytm` → `SMALLMONEY` | `retail.dbo.trsecbondsprice_th.Bidytm <- @Bidytm` |

<a id="kolom-changeby"></a>
## ChangeBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECBondsPrice_TH.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006654 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessApprParamSpread:340](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L340) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsecbondsprice_th.ChangeBy` : `INT` | `@pnNik` : `bigint` | `@pnNik` → `INT` | `retail.dbo.trsecbondsprice_th.ChangeBy <- @pnNik` |

<a id="kolom-maturitydate"></a>
## MaturityDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECBondsPrice_TH.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003417 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsertEChannel:92](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsertEChannel.sql#L92) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsecbondsprice_th.MaturityDate` : `DATETIME` | `@MaturityDate` : `varchar(8)` | `@MaturityDate` → `DATETIME` | `retail.dbo.trsecbondsprice_th.MaturityDate <- @MaturityDate` |

<a id="kolom-offerytm"></a>
## Offerytm

**Tipe dasar:** `SMALLMONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECBondsPrice_TH.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003418 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsertEChannel:92](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsertEChannel.sql#L92) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsecbondsprice_th.Offerytm` : `SMALLMONEY` | `@Offerytm` : `money` | `@Offerytm` → `SMALLMONEY` | `retail.dbo.trsecbondsprice_th.Offerytm <- @Offerytm` |

<a id="kolom-pricestatus"></a>
## PriceStatus

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECBondsPrice_TH.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003419 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsertEChannel:92](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsertEChannel.sql#L92) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsecbondsprice_th.PriceStatus` : `SMALLINT` | `0` : `int` | `0` → `SMALLINT` | `retail.dbo.trsecbondsprice_th.PriceStatus <- 0` |
