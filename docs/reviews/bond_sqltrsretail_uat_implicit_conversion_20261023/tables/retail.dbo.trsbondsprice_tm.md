# Retail UAT.dbo.TRSBondsPrice_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 14 lokasi source, 5 procedure, 6 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 2 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-bidprice"></a>
## Bidprice

**Tipe dasar:** `DECIMAL(20,3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBondsPrice_TM.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001382 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSProsesDataJobOutstandingBond:116](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L116) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tm.Bidprice` : `DECIMAL(20,3)` | `tmp.IndikasiHargaBid` : `FLOAT` | `tm.Bidprice` → `FLOAT` | `tmp.IndikasiHargaBid = tm.Bidprice` |
| IC-011486 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptRekSuratBerharga:151](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptRekSuratBerharga.sql#L151) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `b.[Bidprice]` : `DECIMAL(20,3)` | `0` : `int` | `0` → `DECIMAL(20,3)` | `isnull(b.[Bidprice], 0)` |
| IC-011488 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptRekSuratBerharga:152](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptRekSuratBerharga.sql#L152) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `b.[Bidprice]` : `DECIMAL(20,3)` | `0` : `int` | `0` → `DECIMAL(20,3)` | `isnull(b.[Bidprice],0)` |
| IC-013286 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListOutstanding:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListOutstanding.sql#L33) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tbp.Bidprice` : `DECIMAL(20,3)` | `0` : `int` | `0` → `DECIMAL(20,3)` | `isnull(tbp.Bidprice,0)` |

<a id="kolom-bidsize"></a>
## Bidsize

**Tipe dasar:** `FLOAT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBondsPrice_TM.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003401 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsert:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsert.sql#L59) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbondsprice_tm.Bidsize` : `FLOAT` | `0` : `int` | `0` → `FLOAT` | `retail.dbo.trsbondsprice_tm.Bidsize <- 0` |
| IC-003408 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsert:118](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsert.sql#L118) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbondsprice_tm.Bidsize` : `FLOAT` | `0` : `int` | `0` → `FLOAT` | `retail.dbo.trsbondsprice_tm.Bidsize <- 0` |

<a id="kolom-offerprice"></a>
## Offerprice

**Tipe dasar:** `DECIMAL(20,3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBondsPrice_TM.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007751 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessUploadTransaksiSuratBerharga:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessUploadTransaksiSuratBerharga.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `isnull(bp.Offerprice,0)` : `DECIMAL(20,3)` | `#tempSecurityTransaction_TT.HargaORI` : `MONEY` | `isnull(bp.Offerprice,0)` → `MONEY` | `#tempSecurityTransaction_TT.HargaORI <- isnull(bp.Offerprice,0)` |
| IC-007764 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessUploadTransaksiSuratBerharga:148](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessUploadTransaksiSuratBerharga.sql#L148) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bp.Offerprice` : `DECIMAL(20,3)` | `0` : `int` | `0` → `DECIMAL(20,3)` | `isnull(bp.Offerprice,0)` |

<a id="kolom-offersize"></a>
## Offersize

**Tipe dasar:** `FLOAT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBondsPrice_TM.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003402 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsert:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsert.sql#L59) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbondsprice_tm.Offersize` : `FLOAT` | `0` : `int` | `0` → `FLOAT` | `retail.dbo.trsbondsprice_tm.Offersize <- 0` |
| IC-003409 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsert:118](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsert.sql#L118) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbondsprice_tm.Offersize` : `FLOAT` | `0` : `int` | `0` → `FLOAT` | `retail.dbo.trsbondsprice_tm.Offersize <- 0` |

<a id="kolom-pricestatus"></a>
## PriceStatus

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBondsPrice_TM.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003403 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsert:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsert.sql#L59) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbondsprice_tm.PriceStatus` : `SMALLINT` | `0` : `int` | `0` → `SMALLINT` | `retail.dbo.trsbondsprice_tm.PriceStatus <- 0` |
| IC-003410 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsert:118](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsert.sql#L118) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbondsprice_tm.PriceStatus` : `SMALLINT` | `1` : `int` | `1` → `SMALLINT` | `retail.dbo.trsbondsprice_tm.PriceStatus <- 1` |

<a id="kolom-recordingdate-date"></a>
## Recordingdate_date

**Tipe dasar:** `CHAR(25)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBondsPrice_TM.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003404 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsert:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsert.sql#L59) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbondsprice_tm.Recordingdate_date` : `CHAR(25)` | `getdate()` : `datetime` | `getdate()` → `CHAR(25)` | `retail.dbo.trsbondsprice_tm.Recordingdate_date <- getdate()` |
| IC-003411 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsert:118](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsert.sql#L118) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbondsprice_tm.Recordingdate_date` : `CHAR(25)` | `getdate()` : `datetime` | `getdate()` → `CHAR(25)` | `retail.dbo.trsbondsprice_tm.Recordingdate_date <- getdate()` |
