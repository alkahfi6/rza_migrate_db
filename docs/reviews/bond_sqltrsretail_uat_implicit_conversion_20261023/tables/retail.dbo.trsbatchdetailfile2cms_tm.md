# Retail UAT.dbo.TRSBatchDetailFile2CMS_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 86 lokasi source, 3 procedure, 11 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-accruedinterest"></a>
## AccruedInterest

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBatchDetailFile2CMS_TM.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000319 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L228) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.AccruedInterest <- 0` |
| IC-000347 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L323) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.AccruedInterest <- 0` |
| IC-000368 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.AccruedInterest <- 0` |
| IC-000405 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L228) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.AccruedInterest <- 0` |
| IC-000433 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L323) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.AccruedInterest <- 0` |
| IC-000454 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.AccruedInterest <- 0` |
| IC-008996 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailORIAllotment:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L131) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.AccruedInterest <- 0` |
| IC-009016 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailORIAllotment:285](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L285) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.AccruedInterest <- 0` |

<a id="kolom-batchid"></a>
## BatchId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBatchDetailFile2CMS_TM.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 10 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000308 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L96) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.BatchId` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `retail.dbo.trsbatchdetailfile2cms_tm.BatchId <- @nBatchID` |
| IC-000311 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:139](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L139) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.BatchId` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `retail.dbo.trsbatchdetailfile2cms_tm.BatchId <- @nBatchID` |
| IC-000320 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L228) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.BatchId` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `retail.dbo.trsbatchdetailfile2cms_tm.BatchId <- @nBatchID` |
| IC-000348 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L323) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.BatchId` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `retail.dbo.trsbatchdetailfile2cms_tm.BatchId <- @nBatchID` |
| IC-000369 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.BatchId` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `retail.dbo.trsbatchdetailfile2cms_tm.BatchId <- @nBatchID` |
| IC-000394 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L96) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.BatchId` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `retail.dbo.trsbatchdetailfile2cms_tm.BatchId <- @nBatchID` |
| IC-000397 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:139](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L139) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.BatchId` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `retail.dbo.trsbatchdetailfile2cms_tm.BatchId <- @nBatchID` |
| IC-000406 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L228) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.BatchId` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `retail.dbo.trsbatchdetailfile2cms_tm.BatchId <- @nBatchID` |
| IC-000434 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L323) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.BatchId` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `retail.dbo.trsbatchdetailfile2cms_tm.BatchId <- @nBatchID` |
| IC-000455 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.BatchId` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `retail.dbo.trsbatchdetailfile2cms_tm.BatchId <- @nBatchID` |

<a id="kolom-d-k"></a>
## D/K

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBatchDetailFile2CMS_TM.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 12 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000309 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L96) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.D/K` : `CHAR(1)` | `'K'` : `varchar` | `'K'` → `CHAR(1)` | `retail.dbo.trsbatchdetailfile2cms_tm.D/K <- 'K'` |
| IC-000312 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:139](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L139) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.D/K` : `CHAR(1)` | `'D'` : `varchar` | `'D'` → `CHAR(1)` | `retail.dbo.trsbatchdetailfile2cms_tm.D/K <- 'D'` |
| IC-000321 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L228) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.D/K` : `CHAR(1)` | `'K'` : `varchar` | `'K'` → `CHAR(1)` | `retail.dbo.trsbatchdetailfile2cms_tm.D/K <- 'K'` |
| IC-000349 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L323) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.D/K` : `CHAR(1)` | `'K'` : `varchar` | `'K'` → `CHAR(1)` | `retail.dbo.trsbatchdetailfile2cms_tm.D/K <- 'K'` |
| IC-000370 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.D/K` : `CHAR(1)` | `'K'` : `varchar` | `'K'` → `CHAR(1)` | `retail.dbo.trsbatchdetailfile2cms_tm.D/K <- 'K'` |
| IC-000395 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L96) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.D/K` : `CHAR(1)` | `'K'` : `varchar` | `'K'` → `CHAR(1)` | `retail.dbo.trsbatchdetailfile2cms_tm.D/K <- 'K'` |
| IC-000398 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:139](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L139) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.D/K` : `CHAR(1)` | `'D'` : `varchar` | `'D'` → `CHAR(1)` | `retail.dbo.trsbatchdetailfile2cms_tm.D/K <- 'D'` |
| IC-000407 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L228) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.D/K` : `CHAR(1)` | `'K'` : `varchar` | `'K'` → `CHAR(1)` | `retail.dbo.trsbatchdetailfile2cms_tm.D/K <- 'K'` |
| IC-000435 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L323) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.D/K` : `CHAR(1)` | `'K'` : `varchar` | `'K'` → `CHAR(1)` | `retail.dbo.trsbatchdetailfile2cms_tm.D/K <- 'K'` |
| IC-000456 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.D/K` : `CHAR(1)` | `'K'` : `varchar` | `'K'` → `CHAR(1)` | `retail.dbo.trsbatchdetailfile2cms_tm.D/K <- 'K'` |
| IC-008997 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailORIAllotment:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L131) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.D/K` : `CHAR(1)` | `'D'` : `varchar` | `'D'` → `CHAR(1)` | `retail.dbo.trsbatchdetailfile2cms_tm.D/K <- 'D'` |
| IC-009017 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailORIAllotment:285](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L285) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.D/K` : `CHAR(1)` | `'D'` : `varchar` | `'D'` → `CHAR(1)` | `retail.dbo.trsbatchdetailfile2cms_tm.D/K <- 'D'` |

<a id="kolom-facevalue"></a>
## FaceValue

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBatchDetailFile2CMS_TM.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000371 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.FaceValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.FaceValue <- 0` |
| IC-000457 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.FaceValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.FaceValue <- 0` |

<a id="kolom-proceed"></a>
## Proceed

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBatchDetailFile2CMS_TM.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000322 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L228) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.Proceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.Proceed <- 0` |
| IC-000350 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L323) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.Proceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.Proceed <- 0` |
| IC-000372 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.Proceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.Proceed <- 0` |
| IC-000408 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L228) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.Proceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.Proceed <- 0` |
| IC-000436 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L323) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.Proceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.Proceed <- 0` |
| IC-000458 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.Proceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.Proceed <- 0` |

<a id="kolom-safekeepingfeeaftertax"></a>
## SafeKeepingFeeAfterTax

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBatchDetailFile2CMS_TM.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000323 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L228) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAfterTax` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAfterTax <- 0` |
| IC-000351 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L323) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAfterTax` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAfterTax <- 0` |
| IC-000373 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAfterTax` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAfterTax <- 0` |
| IC-000409 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L228) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAfterTax` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAfterTax <- 0` |
| IC-000437 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L323) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAfterTax` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAfterTax <- 0` |
| IC-000459 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAfterTax` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAfterTax <- 0` |
| IC-008998 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailORIAllotment:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L131) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAfterTax` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAfterTax <- 0` |
| IC-009018 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailORIAllotment:285](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L285) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAfterTax` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAfterTax <- 0` |

<a id="kolom-safekeepingfeeamount"></a>
## SafeKeepingFeeAmount

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBatchDetailFile2CMS_TM.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000324 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L228) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAmount <- 0` |
| IC-000352 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L323) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAmount <- 0` |
| IC-000374 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAmount <- 0` |
| IC-000410 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L228) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAmount <- 0` |
| IC-000438 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L323) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAmount <- 0` |
| IC-000460 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAmount <- 0` |
| IC-008999 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailORIAllotment:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L131) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAmount <- 0` |
| IC-009019 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailORIAllotment:285](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L285) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.SafeKeepingFeeAmount <- 0` |

<a id="kolom-taxamount"></a>
## TaxAmount

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBatchDetailFile2CMS_TM.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000325 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L228) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxAmount <- 0` |
| IC-000353 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L323) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxAmount <- 0` |
| IC-000375 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxAmount <- 0` |
| IC-000411 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L228) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxAmount <- 0` |
| IC-000439 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L323) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxAmount <- 0` |
| IC-000461 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxAmount <- 0` |
| IC-009000 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailORIAllotment:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L131) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxAmount <- 0` |
| IC-009020 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailORIAllotment:285](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L285) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxAmount <- 0` |

<a id="kolom-taxonaccrued"></a>
## TaxOnAccrued

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBatchDetailFile2CMS_TM.sql#L19). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000326 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L228) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnAccrued <- 0` |
| IC-000354 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L323) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnAccrued <- 0` |
| IC-000376 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnAccrued <- 0` |
| IC-000412 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L228) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnAccrued <- 0` |
| IC-000440 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L323) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnAccrued <- 0` |
| IC-000462 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnAccrued <- 0` |
| IC-009001 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailORIAllotment:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L131) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnAccrued <- 0` |
| IC-009021 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailORIAllotment:285](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L285) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnAccrued <- 0` |

<a id="kolom-taxoncapitalgainloss"></a>
## TaxOnCapitalGainLoss

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBatchDetailFile2CMS_TM.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000327 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L228) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnCapitalGainLoss <- 0` |
| IC-000355 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L323) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnCapitalGainLoss <- 0` |
| IC-000377 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnCapitalGainLoss <- 0` |
| IC-000413 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L228) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnCapitalGainLoss <- 0` |
| IC-000441 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L323) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnCapitalGainLoss <- 0` |
| IC-000463 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnCapitalGainLoss <- 0` |
| IC-009002 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailORIAllotment:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L131) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnCapitalGainLoss <- 0` |
| IC-009022 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailORIAllotment:285](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L285) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TaxOnCapitalGainLoss <- 0` |

<a id="kolom-transactionfeeamount"></a>
## TransactionFeeAmount

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBatchDetailFile2CMS_TM.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000328 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L228) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TransactionFeeAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TransactionFeeAmount <- 0` |
| IC-000356 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L323) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TransactionFeeAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TransactionFeeAmount <- 0` |
| IC-000378 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TransactionFeeAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TransactionFeeAmount <- 0` |
| IC-000414 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L228) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TransactionFeeAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TransactionFeeAmount <- 0` |
| IC-000442 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L323) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TransactionFeeAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TransactionFeeAmount <- 0` |
| IC-000464 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L410) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TransactionFeeAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TransactionFeeAmount <- 0` |
| IC-009003 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailORIAllotment:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L131) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TransactionFeeAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TransactionFeeAmount <- 0` |
| IC-009023 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailORIAllotment:285](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L285) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchdetailfile2cms_tm.TransactionFeeAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.trsbatchdetailfile2cms_tm.TransactionFeeAmount <- 0` |
