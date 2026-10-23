# Retail UAT.dbo.PTOSTransaksi_Error — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 68 lokasi source, 2 procedure, 17 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-accruedinterest"></a>
## AccruedInterest

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Error.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001447 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.AccruedInterest` : `MONEY` | `AccruedInterest` : `FLOAT` | `AccruedInterest` → `MONEY` | `retail.dbo.ptostransaksi_error.AccruedInterest <- AccruedInterest` |
| IC-001473 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.AccruedInterest` : `MONEY` | `AccruedInterest` : `FLOAT` | `AccruedInterest` → `MONEY` | `retail.dbo.ptostransaksi_error.AccruedInterest <- AccruedInterest` |
| IC-001575 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.AccruedInterest` : `MONEY` | `AccruedInterest` : `FLOAT` | `AccruedInterest` → `MONEY` | `retail.dbo.ptostransaksi_error.AccruedInterest <- AccruedInterest` |
| IC-001601 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.AccruedInterest` : `MONEY` | `AccruedInterest` : `FLOAT` | `AccruedInterest` → `MONEY` | `retail.dbo.ptostransaksi_error.AccruedInterest <- AccruedInterest` |

<a id="kolom-ccy"></a>
## Ccy

**Tipe dasar:** `VARCHAR(3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Error.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001448 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.Ccy` : `VARCHAR(3)` | `Ccy` : `NVARCHAR(255)` | `Ccy` → `VARCHAR(3)` | `retail.dbo.ptostransaksi_error.Ccy <- Ccy` |
| IC-001474 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.Ccy` : `VARCHAR(3)` | `Ccy` : `NVARCHAR(255)` | `Ccy` → `VARCHAR(3)` | `retail.dbo.ptostransaksi_error.Ccy <- Ccy` |
| IC-001576 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.Ccy` : `VARCHAR(3)` | `Ccy` : `NVARCHAR(255)` | `Ccy` → `VARCHAR(3)` | `retail.dbo.ptostransaksi_error.Ccy <- Ccy` |
| IC-001602 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.Ccy` : `VARCHAR(3)` | `Ccy` : `NVARCHAR(255)` | `Ccy` → `VARCHAR(3)` | `retail.dbo.ptostransaksi_error.Ccy <- Ccy` |

<a id="kolom-dealid"></a>
## DealId

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Error.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001449 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.DealId` : `VARCHAR(20)` | `DealId` : `NVARCHAR(255)` | `DealId` → `VARCHAR(20)` | `retail.dbo.ptostransaksi_error.DealId <- DealId` |
| IC-001475 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.DealId` : `VARCHAR(20)` | `DealId` : `NVARCHAR(255)` | `DealId` → `VARCHAR(20)` | `retail.dbo.ptostransaksi_error.DealId <- DealId` |
| IC-001577 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.DealId` : `VARCHAR(20)` | `DealId` : `NVARCHAR(255)` | `DealId` → `VARCHAR(20)` | `retail.dbo.ptostransaksi_error.DealId <- DealId` |
| IC-001603 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.DealId` : `VARCHAR(20)` | `DealId` : `NVARCHAR(255)` | `DealId` → `VARCHAR(20)` | `retail.dbo.ptostransaksi_error.DealId <- DealId` |

<a id="kolom-dealprice"></a>
## DealPrice

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Error.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001450 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.DealPrice` : `DECIMAL(20,5)` | `DealPrice` : `FLOAT` | `DealPrice` → `DECIMAL(20,5)` | `retail.dbo.ptostransaksi_error.DealPrice <- DealPrice` |
| IC-001476 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.DealPrice` : `DECIMAL(20,5)` | `DealPrice` : `FLOAT` | `DealPrice` → `DECIMAL(20,5)` | `retail.dbo.ptostransaksi_error.DealPrice <- DealPrice` |
| IC-001578 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.DealPrice` : `DECIMAL(20,5)` | `DealPrice` : `FLOAT` | `DealPrice` → `DECIMAL(20,5)` | `retail.dbo.ptostransaksi_error.DealPrice <- DealPrice` |
| IC-001604 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.DealPrice` : `DECIMAL(20,5)` | `DealPrice` : `FLOAT` | `DealPrice` → `DECIMAL(20,5)` | `retail.dbo.ptostransaksi_error.DealPrice <- DealPrice` |

<a id="kolom-gainloss"></a>
## GainLoss

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Error.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001451 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.GainLoss` : `MONEY` | `GainLoss` : `FLOAT` | `GainLoss` → `MONEY` | `retail.dbo.ptostransaksi_error.GainLoss <- GainLoss` |
| IC-001477 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.GainLoss` : `MONEY` | `GainLoss` : `FLOAT` | `GainLoss` → `MONEY` | `retail.dbo.ptostransaksi_error.GainLoss <- GainLoss` |
| IC-001579 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.GainLoss` : `MONEY` | `GainLoss` : `FLOAT` | `GainLoss` → `MONEY` | `retail.dbo.ptostransaksi_error.GainLoss <- GainLoss` |
| IC-001605 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.GainLoss` : `MONEY` | `GainLoss` : `FLOAT` | `GainLoss` → `MONEY` | `retail.dbo.ptostransaksi_error.GainLoss <- GainLoss` |

<a id="kolom-hargaperolehan"></a>
## HargaPerolehan

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Error.sql#L19). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001452 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.HargaPerolehan` : `DECIMAL(20,5)` | `HargaPerolehan` : `FLOAT` | `HargaPerolehan` → `DECIMAL(20,5)` | `retail.dbo.ptostransaksi_error.HargaPerolehan <- HargaPerolehan` |
| IC-001478 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.HargaPerolehan` : `DECIMAL(20,5)` | `HargaPerolehan` : `FLOAT` | `HargaPerolehan` → `DECIMAL(20,5)` | `retail.dbo.ptostransaksi_error.HargaPerolehan <- HargaPerolehan` |
| IC-001580 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.HargaPerolehan` : `DECIMAL(20,5)` | `HargaPerolehan` : `FLOAT` | `HargaPerolehan` → `DECIMAL(20,5)` | `retail.dbo.ptostransaksi_error.HargaPerolehan <- HargaPerolehan` |
| IC-001606 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.HargaPerolehan` : `DECIMAL(20,5)` | `HargaPerolehan` : `FLOAT` | `HargaPerolehan` → `DECIMAL(20,5)` | `retail.dbo.ptostransaksi_error.HargaPerolehan <- HargaPerolehan` |

<a id="kolom-kodeproduk"></a>
## KodeProduk

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Error.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001453 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.KodeProduk` : `VARCHAR(20)` | `KodeProduk` : `NVARCHAR(255)` | `KodeProduk` → `VARCHAR(20)` | `retail.dbo.ptostransaksi_error.KodeProduk <- KodeProduk` |
| IC-001479 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.KodeProduk` : `VARCHAR(20)` | `KodeProduk` : `NVARCHAR(255)` | `KodeProduk` → `VARCHAR(20)` | `retail.dbo.ptostransaksi_error.KodeProduk <- KodeProduk` |
| IC-001581 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.KodeProduk` : `VARCHAR(20)` | `KodeProduk` : `NVARCHAR (expression; length unknown)` | `KodeProduk` → `VARCHAR(20)` | `retail.dbo.ptostransaksi_error.KodeProduk <- KodeProduk` |
| IC-001607 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.KodeProduk` : `VARCHAR(20)` | `KodeProduk` : `NVARCHAR (expression; length unknown)` | `KodeProduk` → `VARCHAR(20)` | `retail.dbo.ptostransaksi_error.KodeProduk <- KodeProduk` |

<a id="kolom-npwp"></a>
## NPWP

**Tipe dasar:** `VARCHAR(40)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Error.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001454 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.NPWP` : `VARCHAR(40)` | `NPWP` : `NVARCHAR(255)` | `NPWP` → `VARCHAR(40)` | `retail.dbo.ptostransaksi_error.NPWP <- NPWP` |
| IC-001480 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.NPWP` : `VARCHAR(40)` | `tmp.NPWP` : `NVARCHAR(255)` | `tmp.NPWP` → `VARCHAR(40)` | `retail.dbo.ptostransaksi_error.NPWP <- tmp.NPWP` |
| IC-001582 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.NPWP` : `VARCHAR(40)` | `NPWP` : `NVARCHAR(255)` | `NPWP` → `VARCHAR(40)` | `retail.dbo.ptostransaksi_error.NPWP <- NPWP` |
| IC-001608 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.NPWP` : `VARCHAR(40)` | `tmp.NPWP` : `NVARCHAR(255)` | `tmp.NPWP` → `VARCHAR(40)` | `retail.dbo.ptostransaksi_error.NPWP <- tmp.NPWP` |

<a id="kolom-namanasabah"></a>
## NamaNasabah

**Tipe dasar:** `VARCHAR(200)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Error.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001455 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.NamaNasabah` : `VARCHAR(200)` | `NamaNasabah` : `NVARCHAR(255)` | `NamaNasabah` → `VARCHAR(200)` | `retail.dbo.ptostransaksi_error.NamaNasabah <- NamaNasabah` |
| IC-001481 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.NamaNasabah` : `VARCHAR(200)` | `tmp.NamaNasabah` : `NVARCHAR(255)` | `tmp.NamaNasabah` → `VARCHAR(200)` | `retail.dbo.ptostransaksi_error.NamaNasabah <- tmp.NamaNasabah` |
| IC-001583 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.NamaNasabah` : `VARCHAR(200)` | `NamaNasabah` : `NVARCHAR(255)` | `NamaNasabah` → `VARCHAR(200)` | `retail.dbo.ptostransaksi_error.NamaNasabah <- NamaNasabah` |
| IC-001609 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.NamaNasabah` : `VARCHAR(200)` | `tmp.NamaNasabah` : `NVARCHAR(255)` | `tmp.NamaNasabah` → `VARCHAR(200)` | `retail.dbo.ptostransaksi_error.NamaNasabah <- tmp.NamaNasabah` |

<a id="kolom-namaproduk"></a>
## NamaProduk

**Tipe dasar:** `VARCHAR(200)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Error.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001456 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.NamaProduk` : `VARCHAR(200)` | `NamaProduk` : `NVARCHAR(255)` | `NamaProduk` → `VARCHAR(200)` | `retail.dbo.ptostransaksi_error.NamaProduk <- NamaProduk` |
| IC-001482 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.NamaProduk` : `VARCHAR(200)` | `NamaProduk` : `NVARCHAR(255)` | `NamaProduk` → `VARCHAR(200)` | `retail.dbo.ptostransaksi_error.NamaProduk <- NamaProduk` |
| IC-001584 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.NamaProduk` : `VARCHAR(200)` | `NamaProduk` : `NVARCHAR(255)` | `NamaProduk` → `VARCHAR(200)` | `retail.dbo.ptostransaksi_error.NamaProduk <- NamaProduk` |
| IC-001610 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.NamaProduk` : `VARCHAR(200)` | `NamaProduk` : `NVARCHAR(255)` | `NamaProduk` → `VARCHAR(200)` | `retail.dbo.ptostransaksi_error.NamaProduk <- NamaProduk` |

<a id="kolom-nominal"></a>
## Nominal

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Error.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001457 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.Nominal` : `MONEY` | `Nominal` : `FLOAT` | `Nominal` → `MONEY` | `retail.dbo.ptostransaksi_error.Nominal <- Nominal` |
| IC-001483 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.Nominal` : `MONEY` | `Nominal` : `FLOAT` | `Nominal` → `MONEY` | `retail.dbo.ptostransaksi_error.Nominal <- Nominal` |
| IC-001585 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.Nominal` : `MONEY` | `Nominal` : `FLOAT` | `Nominal` → `MONEY` | `retail.dbo.ptostransaksi_error.Nominal <- Nominal` |
| IC-001611 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.Nominal` : `MONEY` | `Nominal` : `FLOAT` | `Nominal` → `MONEY` | `retail.dbo.ptostransaksi_error.Nominal <- Nominal` |

<a id="kolom-proceed"></a>
## Proceed

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Error.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001458 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.Proceed` : `MONEY` | `Proceed` : `FLOAT` | `Proceed` → `MONEY` | `retail.dbo.ptostransaksi_error.Proceed <- Proceed` |
| IC-001484 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.Proceed` : `MONEY` | `Proceed` : `FLOAT` | `Proceed` → `MONEY` | `retail.dbo.ptostransaksi_error.Proceed <- Proceed` |
| IC-001586 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.Proceed` : `MONEY` | `Proceed` : `FLOAT` | `Proceed` → `MONEY` | `retail.dbo.ptostransaksi_error.Proceed <- Proceed` |
| IC-001612 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.Proceed` : `MONEY` | `Proceed` : `FLOAT` | `Proceed` → `MONEY` | `retail.dbo.ptostransaksi_error.Proceed <- Proceed` |

<a id="kolom-sid"></a>
## SID

**Tipe dasar:** `VARCHAR(100)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Error.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001459 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.SID` : `VARCHAR(100)` | `[SID]` : `NVARCHAR(255)` | `[SID]` → `VARCHAR(100)` | `retail.dbo.ptostransaksi_error.SID <- [SID]` |
| IC-001485 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.SID` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tmp.[SID]` → `VARCHAR(100)` | `retail.dbo.ptostransaksi_error.SID <- tmp.[SID]` |
| IC-001587 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.SID` : `VARCHAR(100)` | `[SID]` : `NVARCHAR (expression; length unknown)` | `[SID]` → `VARCHAR(100)` | `retail.dbo.ptostransaksi_error.SID <- [SID]` |
| IC-001613 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.SID` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR (expression; length unknown)` | `tmp.[SID]` → `VARCHAR(100)` | `retail.dbo.ptostransaksi_error.SID <- tmp.[SID]` |

<a id="kolom-taxcapitalgainloss"></a>
## TaxCapitalGainLoss

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Error.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001460 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.TaxCapitalGainLoss` : `MONEY` | `TaxCapitalGainLoss` : `FLOAT` | `TaxCapitalGainLoss` → `MONEY` | `retail.dbo.ptostransaksi_error.TaxCapitalGainLoss <- TaxCapitalGainLoss` |
| IC-001486 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.TaxCapitalGainLoss` : `MONEY` | `TaxCapitalGainLoss` : `FLOAT` | `TaxCapitalGainLoss` → `MONEY` | `retail.dbo.ptostransaksi_error.TaxCapitalGainLoss <- TaxCapitalGainLoss` |
| IC-001588 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.TaxCapitalGainLoss` : `MONEY` | `TaxCapitalGainLoss` : `FLOAT` | `TaxCapitalGainLoss` → `MONEY` | `retail.dbo.ptostransaksi_error.TaxCapitalGainLoss <- TaxCapitalGainLoss` |
| IC-001614 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.TaxCapitalGainLoss` : `MONEY` | `TaxCapitalGainLoss` : `FLOAT` | `TaxCapitalGainLoss` → `MONEY` | `retail.dbo.ptostransaksi_error.TaxCapitalGainLoss <- TaxCapitalGainLoss` |

<a id="kolom-taxonaccrued"></a>
## TaxOnAccrued

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Error.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001461 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.TaxOnAccrued` : `MONEY` | `TaxOnAccrued` : `FLOAT` | `TaxOnAccrued` → `MONEY` | `retail.dbo.ptostransaksi_error.TaxOnAccrued <- TaxOnAccrued` |
| IC-001487 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.TaxOnAccrued` : `MONEY` | `TaxOnAccrued` : `FLOAT` | `TaxOnAccrued` → `MONEY` | `retail.dbo.ptostransaksi_error.TaxOnAccrued <- TaxOnAccrued` |
| IC-001589 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.TaxOnAccrued` : `MONEY` | `TaxOnAccrued` : `FLOAT` | `TaxOnAccrued` → `MONEY` | `retail.dbo.ptostransaksi_error.TaxOnAccrued <- TaxOnAccrued` |
| IC-001615 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.TaxOnAccrued` : `MONEY` | `TaxOnAccrued` : `FLOAT` | `TaxOnAccrued` → `MONEY` | `retail.dbo.ptostransaksi_error.TaxOnAccrued <- TaxOnAccrued` |

<a id="kolom-totalproceed"></a>
## TotalProceed

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Error.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001462 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.TotalProceed` : `MONEY` | `TotalProceed` : `FLOAT` | `TotalProceed` → `MONEY` | `retail.dbo.ptostransaksi_error.TotalProceed <- TotalProceed` |
| IC-001488 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.TotalProceed` : `MONEY` | `TotalProceed` : `FLOAT` | `TotalProceed` → `MONEY` | `retail.dbo.ptostransaksi_error.TotalProceed <- TotalProceed` |
| IC-001590 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.TotalProceed` : `MONEY` | `TotalProceed` : `FLOAT` | `TotalProceed` → `MONEY` | `retail.dbo.ptostransaksi_error.TotalProceed <- TotalProceed` |
| IC-001616 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.TotalProceed` : `MONEY` | `TotalProceed` : `FLOAT` | `TotalProceed` → `MONEY` | `retail.dbo.ptostransaksi_error.TotalProceed <- TotalProceed` |

<a id="kolom-trxtype"></a>
## TrxType

**Tipe dasar:** `VARCHAR(8)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Error.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001463 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.TrxType` : `VARCHAR(8)` | `TrxType` : `NVARCHAR(255)` | `TrxType` → `VARCHAR(8)` | `retail.dbo.ptostransaksi_error.TrxType <- TrxType` |
| IC-001489 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.TrxType` : `VARCHAR(8)` | `TrxType` : `NVARCHAR(255)` | `TrxType` → `VARCHAR(8)` | `retail.dbo.ptostransaksi_error.TrxType <- TrxType` |
| IC-001591 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.TrxType` : `VARCHAR(8)` | `TrxType` : `NVARCHAR(255)` | `TrxType` → `VARCHAR(8)` | `retail.dbo.ptostransaksi_error.TrxType <- TrxType` |
| IC-001617 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_error.TrxType` : `VARCHAR(8)` | `TrxType` : `NVARCHAR(255)` | `TrxType` → `VARCHAR(8)` | `retail.dbo.ptostransaksi_error.TrxType <- TrxType` |
