# Retail UAT.dbo.TRSHasilRevenue_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 73 lokasi source, 6 procedure, 16 kolom. **A:** 6 lokasi pada 2 procedure. **B:** 7 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-tanggaltransaksi"></a>
## TanggalTransaksi

**Tipe dasar:** `VARCHAR(8)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSHasilRevenue_TM.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 4 SP. **A:** 6 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P03 — Tanggal teks bertemu date/time:** Tentukan kontrak YYYYMMDD/hari/timestamp. Bentuk nilai pencarian dalam tipe kolom atau gunakan rentang native date; jangan konversi kolom seluruh baris tanpa alasan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004159 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:134](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L134) | WHERE / P03 / COMPARISON | `mm.[TanggalTransaksi]` : `VARCHAR(8)` | `@pdTglProses` : `DATETIME` | `mm.[TanggalTransaksi]` → `DATETIME` | `mm.[TanggalTransaksi] = @pdTglProses` |
| IC-004192 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:248](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L248) | WHERE / P03 / COMPARISON | `hr.[TanggalTransaksi]` : `VARCHAR(8)` | `@pdTglProses` : `DATETIME` | `hr.[TanggalTransaksi]` → `DATETIME` | `hr.[TanggalTransaksi] = @pdTglProses` |
| IC-004198 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:375](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L375) | WHERE / P03 / COMPARISON | `tm.[TanggalTransaksi]` : `VARCHAR(8)` | `@pdTglProses` : `DATETIME` | `tm.[TanggalTransaksi]` → `DATETIME` | `tm.[TanggalTransaksi] = @pdTglProses` |
| IC-004205 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:470](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L470) | WHERE / P03 / COMPARISON | `chr.[TanggalTransaksi]` : `VARCHAR(8)` | `@pdTglProses` : `DATETIME` | `chr.[TanggalTransaksi]` → `DATETIME` | `chr.[TanggalTransaksi] = @pdTglProses` |
| IC-012200 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSSentEmailHasilRevenueBondGagal:97](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailHasilRevenueBondGagal.sql#L97) | WHERE / P03 / COMPARISON | `tm.TanggalTransaksi` : `VARCHAR(8)` | `@pdTglProses` : `DATETIME` | `tm.TanggalTransaksi` → `DATETIME` | `tm.TanggalTransaksi >= @pdTglProses` |
| IC-012201 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSSentEmailHasilRevenueBondGagal:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailHasilRevenueBondGagal.sql#L98) | WHERE / P03 / COMPARISON | `tm.TanggalTransaksi` : `VARCHAR(8)` | `@dtNextDay` : `DATETIME` | `tm.TanggalTransaksi` → `DATETIME` | `tm.TanggalTransaksi < @dtNextDay` |
| IC-007918 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRPTRevenueOBLforTrops:60](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforTrops.sql#L60) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.TanggalTransaksi` : `VARCHAR(8)` | `#tmpDailyRevenue.TanggalTransaksi` : `DATETIME` | `tm.TanggalTransaksi` → `DATETIME` | `#tmpDailyRevenue.TanggalTransaksi <- tm.TanggalTransaksi` |
| IC-007950 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRPTRevenueOBLforWM:84](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforWM.sql#L84) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.TanggalTransaksi` : `VARCHAR(8)` | `#tmpOblRev.TanggalTransaksi` : `DATETIME` | `tm.TanggalTransaksi` → `DATETIME` | `#tmpOblRev.TanggalTransaksi <- tm.TanggalTransaksi` |

<a id="kolom-amount"></a>
## Amount

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSHasilRevenue_TM.sql#L37). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004188 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:196](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L196) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mast.Amount` : `MONEY` | `ttm.ProfitRCY` : `DECIMAL(20,2)` | `ttm.ProfitRCY` → `MONEY` | `mast.Amount = ttm.ProfitRCY` |
| IC-004260 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:263](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L263) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trshasilrevenue_tm.Amount` : `MONEY` | `dd.ProfitRCY` : `DECIMAL(20, 5)` | `dd.ProfitRCY` → `MONEY` | `retail.dbo.trshasilrevenue_tm.Amount <- dd.ProfitRCY` |
| IC-004279 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:312](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L312) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trshasilrevenue_tm.Amount` : `MONEY` | `dd.ProfitRCY` : `DECIMAL(20, 5)` | `dd.ProfitRCY` → `MONEY` | `retail.dbo.trshasilrevenue_tm.Amount <- dd.ProfitRCY` |

<a id="kolom-birate"></a>
## BIRate

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSHasilRevenue_TM.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004190 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSHasilRevenueBondJurnalGL:231](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L231) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `hr. BIRate` : `MONEY` | `#tmpBulkLog.Rate` : `DECIMAL(20, 5)` | `hr. BIRate` → `DECIMAL(20, 5)` | `#tmpBulkLog.Rate <- hr. BIRate` |
| IC-004202 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSHasilRevenueBondJurnalGL:448](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L448) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `chr.[BIRate]` : `MONEY` | `retail.dbo.trspostingglbond_tt.Rate` : `DECIMAL(25,5)` | `chr.[BIRate]` → `DECIMAL(25,5)` | `retail.dbo.trspostingglbond_tt.Rate <- chr.[BIRate]` |
| IC-012198 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSSentEmailHasilRevenueBondGagal:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailHasilRevenueBondGagal.sql#L64) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.BIRate` : `MONEY` | `#tmpRevenueGagal.BI Rate` : `DECIMAL(8,4)` | `tm.BIRate` → `DECIMAL(8,4)` | `#tmpRevenueGagal.BI Rate <- tm.BIRate` |
| IC-004195 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:342](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L342) | JOIN_ON / P11 / COMPARISON | `ins.[BIRate]` : `MONEY` | `tz.[Rate]` : `DECIMAL(20, 5)` | `ins.[BIRate]` → `DECIMAL` | `ins.[BIRate] = tz.[Rate]` |
| IC-004196 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:364](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L364) | JOIN_ON / P11 / COMPARISON | `tm.[BIRate]` : `MONEY` | `bl.[Rate]` : `DECIMAL(20, 5)` | `tm.[BIRate]` → `DECIMAL` | `tm.[BIRate] = bl.[Rate]` |
| IC-004201 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:430](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L430) | JOIN_ON / P11 / COMPARISON | `ins.[BIRate]` : `MONEY` | `tz.[Rate]` : `DECIMAL(20, 5)` | `ins.[BIRate]` → `DECIMAL` | `ins.[BIRate] = tz.[Rate]` |

<a id="kolom-branchcredit"></a>
## BranchCredit

**Tipe dasar:** `VARCHAR(10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSHasilRevenue_TM.sql#L34). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004261 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:263](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L263) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trshasilrevenue_tm.BranchCredit` : `VARCHAR(10)` | `CASE WHEN dd.TrxStatus = 3 THEN ISNULL(ta.[M_COSTCENTER], '-') ELSE ISNULL(trdr.[M_COSTCENTER], '-') END` : `NVARCHAR (combined result; facets unknown)` | `CASE WHEN dd.TrxStatus = 3 THEN ISNULL(ta.[M_COSTCENTER], '-') ELSE ISNULL(trdr.[M_COSTCENTER], '-') END` → `VARCHAR(10)` | `retail.dbo.trshasilrevenue_tm.BranchCredit <- CASE WHEN dd.TrxStatus = 3 THEN ISNULL(ta.[M_COSTCENTER], '-') ELSE ISNULL(trdr.[M_COSTCENTER], '-') END` |
| IC-004280 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:312](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L312) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trshasilrevenue_tm.BranchCredit` : `VARCHAR(10)` | `CASE WHEN dd.TrxStatus = 3 THEN CASE WHEN dd.KodeCabangRLS = '0' THEN '-' ELSE dd.KodeCabangRLS END ELSE ISNULL(ta.[M_COSTCENTER], '-') END` : `NVARCHAR (combined result; facets unknown)` | `CASE WHEN dd.TrxStatus = 3 THEN CASE WHEN dd.KodeCabangRLS = '0' THEN '-' ELSE dd.KodeCabangRLS END ELSE ISNULL(ta.[M_COSTCENTER], '-') END` → `VARCHAR(10)` | `retail.dbo.trshasilrevenue_tm.BranchCredit <- CASE WHEN dd.TrxStatus = 3 THEN CASE WHEN dd.KodeCabangRLS = '0' THEN '-' ELSE dd.KodeCabangRLS END ELSE ISNULL(ta.[M_COSTCENTER], '-') END` |

<a id="kolom-branchdebet"></a>
## BranchDebet

**Tipe dasar:** `VARCHAR(10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSHasilRevenue_TM.sql#L31). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004262 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:263](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L263) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trshasilrevenue_tm.BranchDebet` : `VARCHAR(10)` | `CASE WHEN dd.TrxStatus = 3 THEN ISNULL(trdr.[M_COSTCENTER], '-') ELSE ISNULL(ta.[M_COSTCENTER], '-') END` : `NVARCHAR (combined result; facets unknown)` | `CASE WHEN dd.TrxStatus = 3 THEN ISNULL(trdr.[M_COSTCENTER], '-') ELSE ISNULL(ta.[M_COSTCENTER], '-') END` → `VARCHAR(10)` | `retail.dbo.trshasilrevenue_tm.BranchDebet <- CASE WHEN dd.TrxStatus = 3 THEN ISNULL(trdr.[M_COSTCENTER], '-') ELSE ISNULL(ta.[M_COSTCENTER], '-') END` |
| IC-004281 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:312](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L312) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trshasilrevenue_tm.BranchDebet` : `VARCHAR(10)` | `CASE WHEN dd.TrxStatus = 3 THEN ISNULL(ta.[M_COSTCENTER], '-') ELSE CASE WHEN dd.KodeCabangRLS = '0' THEN '-' ELSE dd.KodeCabangRLS END END` : `NVARCHAR (combined result; facets unknown)` | `CASE WHEN dd.TrxStatus = 3 THEN ISNULL(ta.[M_COSTCENTER], '-') ELSE CASE WHEN dd.KodeCabangRLS = '0' THEN '-' ELSE dd.KodeCabangRLS END END` → `VARCHAR(10)` | `retail.dbo.trshasilrevenue_tm.BranchDebet <- CASE WHEN dd.TrxStatus = 3 THEN ISNULL(ta.[M_COSTCENTER], '-') ELSE CASE WHEN dd.KodeCabangRLS = '0' THEN '-' ELSE dd.KodeCabangRLS END END` |

<a id="kolom-bulkid"></a>
## BulkId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSHasilRevenue_TM.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003188 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SettlementPopulateHasilRevenue:35](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SettlementPopulateHasilRevenue.sql#L35) | WHERE / P11 / COMPARISON | `d.[BulkId]` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `d.[BulkId] <> 0` |
| IC-004191 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSHasilRevenueBondJurnalGL:242](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L242) | WHERE / P11 / COMPARISON | `hr.[BulkId]` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `hr.[BulkId] = 0` |
| IC-004197 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSHasilRevenueBondJurnalGL:369](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L369) | WHERE / P11 / COMPARISON | `tm.[BulkId]` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `tm.[BulkId] = 0` |
| IC-004208 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:579](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L579) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[BulkId]` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `[BulkId] = 0` |
| IC-004263 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:263](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L263) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trshasilrevenue_tm.BulkId` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `retail.dbo.trshasilrevenue_tm.BulkId <- 0` |
| IC-004282 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:312](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L312) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trshasilrevenue_tm.BulkId` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `retail.dbo.trshasilrevenue_tm.BulkId <- 0` |
| IC-012199 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSentEmailHasilRevenueBondGagal:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailHasilRevenueBondGagal.sql#L94) | WHERE / P11 / COMPARISON | `tm.BulkId` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `tm.BulkId <> 0` |

<a id="kolom-dealid"></a>
## DealId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSHasilRevenue_TM.sql#L5). **Existing leading index:** TRSHasilRevenue_TM_N2.

**Frekuensi source:** 12 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004203 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSHasilRevenueBondJurnalGL:448](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L448) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `chr.[DealId]` : `BIGINT` | `retail.dbo.trspostingglbond_tt.TranId` : `VARCHAR(20)` | `chr.[DealId]` → `VARCHAR(20)` | `retail.dbo.trspostingglbond_tt.TranId <- chr.[DealId]` |
| IC-003189 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SettlementPopulateHasilRevenue:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SettlementPopulateHasilRevenue.sql#L53) | JOIN_ON / P01 / COMPARISON | `c.DealId` : `BIGINT` | `th.TranId` : `VARCHAR(20)` | `th.TranId` → `BIGINT` | `c.DealId = th.TranId` |
| IC-004204 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSHasilRevenueBondJurnalGL:459](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L459) | JOIN_ON / P01 / COMPARISON | `chr.[DealId]` : `BIGINT` | `tt.[TranId]` : `VARCHAR(20)` | `tt.[TranId]` → `BIGINT` | `chr.[DealId] = tt.[TranId]` |
| IC-004210 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSHasilRevenueBondJurnalGL:685](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L685) | JOIN_ON / P01 / COMPARISON | `ins.[DealId]` : `BIGINT` | `tt.[TranId]` : `VARCHAR(20)` | `tt.[TranId]` → `BIGINT` | `ins.[DealId] = tt.[TranId]` |
| IC-004215 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSHasilRevenueBondJurnalGL:715](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L715) | JOIN_ON / P01 / COMPARISON | `d.[DealId]` : `BIGINT` | `h.[TranId]` : `VARCHAR(20)` | `h.[TranId]` → `BIGINT` | `d.[DealId] = h.[TranId]` |
| IC-004222 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSHasilRevenueBondJurnalGL:853](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L853) | JOIN_ON / P01 / COMPARISON | `ins.[DealId]` : `BIGINT` | `tt.[TranId]` : `VARCHAR(20)` | `tt.[TranId]` → `BIGINT` | `ins.[DealId] = tt.[TranId]` |
| IC-007921 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRPTRevenueOBLforTrops:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforTrops.sql#L70) | JOIN_ON / P01 / COMPARISON | `tm.DealId` : `BIGINT` | `th.TranId` : `VARCHAR(20)` | `th.TranId` → `BIGINT` | `th.TranId = tm.DealId` |
| IC-007925 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRPTRevenueOBLforTrops:120](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforTrops.sql#L120) | JOIN_ON / P01 / COMPARISON | `m.DealId` : `BIGINT` | `h.TranId` : `VARCHAR(20)` | `h.TranId` → `BIGINT` | `m.DealId = h.TranId` |
| IC-007933 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRPTRevenueOBLforTrops:267](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforTrops.sql#L267) | JOIN_ON / P01 / COMPARISON | `m.DealId` : `BIGINT` | `h.TranId` : `VARCHAR(20)` | `h.TranId` → `BIGINT` | `m.DealId = h.TranId` |
| IC-007937 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRPTRevenueOBLforTrops:298](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforTrops.sql#L298) | JOIN_ON / P01 / COMPARISON | `m.DealId` : `BIGINT` | `h.TranId` : `VARCHAR(20)` | `h.TranId` → `BIGINT` | `h.TranId = m.DealId` |
| IC-007946 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRPTRevenueOBLforTrops:374](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforTrops.sql#L374) | JOIN_ON / P01 / COMPARISON | `m.DealId` : `BIGINT` | `h.TranId` : `VARCHAR(20)` | `h.TranId` → `BIGINT` | `m.DealId = h.TranId` |
| IC-007957 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRPTRevenueOBLforWM:126](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforWM.sql#L126) | JOIN_ON / P01 / COMPARISON | `tm.DealId` : `BIGINT` | `th.TranId` : `VARCHAR(20)` | `th.TranId` → `BIGINT` | `tm.DealId = th.TranId` |

<a id="kolom-idramount"></a>
## IDRAmount

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSHasilRevenue_TM.sql#L38). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004189 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:197](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L197) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mast.IDRAmount` : `MONEY` | `ttm.ProfitLCY` : `DECIMAL(20,2)` | `ttm.ProfitLCY` → `MONEY` | `mast.IDRAmount = ttm.ProfitLCY` |
| IC-004264 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:263](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L263) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trshasilrevenue_tm.IDRAmount` : `MONEY` | `dd.ProfitLCY` : `DECIMAL(20, 5)` | `dd.ProfitLCY` → `MONEY` | `retail.dbo.trshasilrevenue_tm.IDRAmount <- dd.ProfitLCY` |
| IC-004283 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:312](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L312) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trshasilrevenue_tm.IDRAmount` : `MONEY` | `dd.ProfitLCY` : `DECIMAL(20, 5)` | `dd.ProfitLCY` → `MONEY` | `retail.dbo.trshasilrevenue_tm.IDRAmount <- dd.ProfitLCY` |
| IC-007919 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRPTRevenueOBLforTrops:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforTrops.sql#L64) | CASE / P05 / BINARY_Multiply | `tm.IDRAmount` : `MONEY` | `-1` : `int` | `-1` → `MONEY` | `-1 * tm.IDRAmount` |
| IC-007923 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRPTRevenueOBLforTrops:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforTrops.sql#L113) | CASE / P05 / BINARY_Multiply | `IDRAmount` : `MONEY` | `-1` : `int` | `-1` → `MONEY` | `-1 * IDRAmount` |
| IC-007931 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRPTRevenueOBLforTrops:261](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforTrops.sql#L261) | CASE / P05 / BINARY_Multiply | `IDRAmount` : `MONEY` | `-1` : `int` | `-1` → `MONEY` | `-1 * IDRAmount` |
| IC-007935 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRPTRevenueOBLforTrops:291](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforTrops.sql#L291) | CASE / P05 / BINARY_Multiply | `IDRAmount` : `MONEY` | `-1` : `int` | `-1` → `MONEY` | `-1 * IDRAmount` |
| IC-007944 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRPTRevenueOBLforTrops:367](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforTrops.sql#L367) | CASE / P05 / BINARY_Multiply | `IDRAmount` : `MONEY` | `-1` : `int` | `-1` → `MONEY` | `-1 * IDRAmount` |

<a id="kolom-kodecabangrls"></a>
## KodeCabangRLS

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSHasilRevenue_TM.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004265 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:263](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L263) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trshasilrevenue_tm.KodeCabangRLS` : `INT` | `dd.KodeCabangRLS` : `VARCHAR(5)` | `dd.KodeCabangRLS` → `INT` | `retail.dbo.trshasilrevenue_tm.KodeCabangRLS <- dd.KodeCabangRLS` |
| IC-004284 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:312](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L312) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trshasilrevenue_tm.KodeCabangRLS` : `INT` | `dd.KodeCabangRLS` : `VARCHAR(5)` | `dd.KodeCabangRLS` → `INT` | `retail.dbo.trshasilrevenue_tm.KodeCabangRLS <- dd.KodeCabangRLS` |

<a id="kolom-kodecabangtrx"></a>
## KodeCabangTRX

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSHasilRevenue_TM.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004266 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:263](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L263) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trshasilrevenue_tm.KodeCabangTRX` : `INT` | `dd.KodeCabangTRX` : `VARCHAR(5)` | `dd.KodeCabangTRX` → `INT` | `retail.dbo.trshasilrevenue_tm.KodeCabangTRX <- dd.KodeCabangTRX` |
| IC-004285 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:312](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L312) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trshasilrevenue_tm.KodeCabangTRX` : `INT` | `dd.KodeCabangTRX` : `VARCHAR(5)` | `dd.KodeCabangTRX` → `INT` | `retail.dbo.trshasilrevenue_tm.KodeCabangTRX <- dd.KodeCabangTRX` |

<a id="kolom-profitlcy"></a>
## ProfitLCY

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSHasilRevenue_TM.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007955 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRPTRevenueOBLforWM:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforWM.sql#L108) | CASE / P05 / BINARY_Multiply | `tm.ProfitLCY` : `DECIMAL(20,5)` | `-1` : `int` | `-1` → `DECIMAL` | `-1 * tm.ProfitLCY` |

<a id="kolom-profitrcy"></a>
## ProfitRCY

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSHasilRevenue_TM.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007953 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRPTRevenueOBLforWM:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforWM.sql#L104) | CASE / P05 / BINARY_Multiply | `tm.ProfitRCY` : `DECIMAL(20,5)` | `-1` : `int` | `-1` → `DECIMAL` | `-1 * tm.ProfitRCY` |

<a id="kolom-rekeningrelasi"></a>
## RekeningRelasi

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSHasilRevenue_TM.sql#L24). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004267 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:263](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L263) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trshasilrevenue_tm.RekeningRelasi` : `VARCHAR(20)` | `dd.ACCTNO` : `DECIMAL(19,0)` | `dd.ACCTNO` → `VARCHAR(20)` | `retail.dbo.trshasilrevenue_tm.RekeningRelasi <- dd.ACCTNO` |
| IC-004286 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:312](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L312) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trshasilrevenue_tm.RekeningRelasi` : `VARCHAR(20)` | `dd.ACCTNO` : `DECIMAL(19,0)` | `dd.ACCTNO` → `VARCHAR(20)` | `retail.dbo.trshasilrevenue_tm.RekeningRelasi <- dd.ACCTNO` |

<a id="kolom-tanggalsettlement"></a>
## TanggalSettlement

**Tipe dasar:** `VARCHAR(8)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSHasilRevenue_TM.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007949 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRPTRevenueOBLforWM:84](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforWM.sql#L84) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.TanggalSettlement` : `VARCHAR(8)` | `#tmpOblRev.TanggalSettlement` : `DATETIME` | `tm.TanggalSettlement` → `DATETIME` | `#tmpOblRev.TanggalSettlement <- tm.TanggalSettlement` |

<a id="kolom-trxstatus"></a>
## TrxStatus

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSHasilRevenue_TM.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 13 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004194 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:262](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L262) | WHERE / P11 / COMPARISON | `hr.TrxStatus` : `SMALLINT` | `11` : `int` | `hr.TrxStatus` → `INT` | `hr.TrxStatus = 11` |
| IC-004200 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:389](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L389) | WHERE / P11 / COMPARISON | `tm.TrxStatus` : `SMALLINT` | `11` : `int` | `tm.TrxStatus` → `INT` | `tm.TrxStatus = 11` |
| IC-004207 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:484](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L484) | WHERE / P11 / COMPARISON | `chr.TrxStatus` : `SMALLINT` | `11` : `int` | `chr.TrxStatus` → `INT` | `chr.TrxStatus = 11` |
| IC-004252 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:190](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L190) | WHERE / P11 / COMPARISON | `m.TrxStatus` : `SMALLINT` | `3` : `int` | `m.TrxStatus` → `INT` | `m.TrxStatus = 3` |
| IC-004255 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:224](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L224) | WHERE / P11 / COMPARISON | `ins.[TrxStatus]` : `SMALLINT` | `3` : `int` | `ins.[TrxStatus]` → `INT` | `ins.[TrxStatus] = 3` |
| IC-004257 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:238](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L238) | WHERE / P11 / COMPARISON | `ins.[TrxStatus]` : `SMALLINT` | `3` : `int` | `ins.[TrxStatus]` → `INT` | `ins.[TrxStatus] = 3` |
| IC-007920 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTRevenueOBLforTrops:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforTrops.sql#L64) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `tm.TrxStatus` : `SMALLINT` | `11` : `int` | `tm.TrxStatus` → `INT` | `WHEN 11 THEN -1 * tm.IDRAmount` |
| IC-007924 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTRevenueOBLforTrops:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforTrops.sql#L113) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `TrxStatus` : `SMALLINT` | `11` : `int` | `TrxStatus` → `INT` | `WHEN 11 THEN -1 * IDRAmount` |
| IC-007932 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTRevenueOBLforTrops:261](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforTrops.sql#L261) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `TrxStatus` : `SMALLINT` | `11` : `int` | `TrxStatus` → `INT` | `WHEN 11 THEN -1 * IDRAmount` |
| IC-007936 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTRevenueOBLforTrops:291](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforTrops.sql#L291) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `TrxStatus` : `SMALLINT` | `11` : `int` | `TrxStatus` → `INT` | `WHEN 11 THEN -1 * IDRAmount` |
| IC-007945 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTRevenueOBLforTrops:367](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforTrops.sql#L367) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `TrxStatus` : `SMALLINT` | `11` : `int` | `TrxStatus` → `INT` | `WHEN 11 THEN -1 * IDRAmount` |
| IC-007954 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTRevenueOBLforWM:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforWM.sql#L104) | CASE / P11 / COMPARISON | `tm.TrxStatus` : `SMALLINT` | `11` : `int` | `tm.TrxStatus` → `INT` | `tm.TrxStatus = 11` |
| IC-007956 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTRevenueOBLforWM:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforWM.sql#L108) | CASE / P11 / COMPARISON | `tm.TrxStatus` : `SMALLINT` | `11` : `int` | `tm.TrxStatus` → `INT` | `tm.TrxStatus = 11` |

<a id="kolom-isbackdate"></a>
## isBackdate

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSHasilRevenue_TM.sql#L40). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004193 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:259](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L259) | WHERE / P11 / COMPARISON | `hr.[isBackdate]` : `BIT` | `1` : `int` | `hr.[isBackdate]` → `INT` | `hr.[isBackdate] = 1` |
| IC-004199 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:386](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L386) | WHERE / P11 / COMPARISON | `tm.[isBackdate]` : `BIT` | `1` : `int` | `tm.[isBackdate]` → `INT` | `tm.[isBackdate] = 1` |
| IC-004206 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSHasilRevenueBondJurnalGL:481](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql#L481) | WHERE / P11 / COMPARISON | `chr.[isBackdate]` : `BIT` | `1` : `int` | `chr.[isBackdate]` → `INT` | `chr.[isBackdate] = 1` |
