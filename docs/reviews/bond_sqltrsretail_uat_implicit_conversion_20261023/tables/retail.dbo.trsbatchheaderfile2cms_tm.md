# Retail UAT.dbo.TRSBatchHeaderFile2CMS_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 38 lokasi source, 8 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-batchid"></a>
## BatchID

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBatchHeaderFile2CMS_TM.sql#L2). **Existing leading index:** IU_TRSBatchDetailFile2CMS_TM.

**Frekuensi source:** 14 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000305 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GenBatchTxtFile2CMS:54](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L54) | WHERE / P11 / COMPARISON | `BatchID` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `BatchID = @nBatchID` |
| IC-000306 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GenBatchTxtFile2CMS:84](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L84) | WHERE / P11 / COMPARISON | `tbh.BatchID` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `tbh.BatchID = @nBatchID` |
| IC-000310 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GenBatchTxtFile2CMS:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L127) | WHERE / P11 / COMPARISON | `tbh.BatchID` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `tbh.BatchID = @nBatchID` |
| IC-000313 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GenBatchTxtFile2CMS:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L168) | WHERE / P11 / COMPARISON | `tbh.BatchID` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `tbh.BatchID = @nBatchID` |
| IC-000345 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GenBatchTxtFile2CMS:303](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L303) | WHERE / P11 / COMPARISON | `tbh.BatchID` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `tbh.BatchID = @nBatchID` |
| IC-000366 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GenBatchTxtFile2CMS:390](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L390) | WHERE / P11 / COMPARISON | `tbh.BatchID` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `tbh.BatchID = @nBatchID` |
| IC-000390 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GenBatchTxtFile2CMS:472](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L472) | WHERE / P11 / COMPARISON | `BatchID` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `BatchID = @nBatchID` |
| IC-000391 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GenBatchTxtFile2CMS_Urgent:54](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L54) | WHERE / P11 / COMPARISON | `BatchID` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `BatchID = @nBatchID` |
| IC-000392 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GenBatchTxtFile2CMS_Urgent:84](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L84) | WHERE / P11 / COMPARISON | `tbh.BatchID` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `tbh.BatchID = @nBatchID` |
| IC-000396 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GenBatchTxtFile2CMS_Urgent:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L127) | WHERE / P11 / COMPARISON | `tbh.BatchID` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `tbh.BatchID = @nBatchID` |
| IC-000399 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GenBatchTxtFile2CMS_Urgent:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L168) | WHERE / P11 / COMPARISON | `tbh.BatchID` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `tbh.BatchID = @nBatchID` |
| IC-000431 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GenBatchTxtFile2CMS_Urgent:303](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L303) | WHERE / P11 / COMPARISON | `tbh.BatchID` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `tbh.BatchID = @nBatchID` |
| IC-000452 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GenBatchTxtFile2CMS_Urgent:390](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L390) | WHERE / P11 / COMPARISON | `tbh.BatchID` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `tbh.BatchID = @nBatchID` |
| IC-000476 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GenBatchTxtFile2CMS_Urgent:472](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L472) | WHERE / P11 / COMPARISON | `BatchID` : `BIGINT` | `@nBatchID` : `int` | `@nBatchID` → `BIGINT` | `BatchID = @nBatchID` |

<a id="kolom-bitprocess"></a>
## BitProcess

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBatchHeaderFile2CMS_TM.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000389 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS:471](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS.sql#L471) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `BitProcess` : `BIT` | `1` : `int` | `1` → `BIT` | `BitProcess = 1` |
| IC-000475 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GenBatchTxtFile2CMS_Urgent:471](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql#L471) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `BitProcess` : `BIT` | `1` : `int` | `1` → `BIT` | `BitProcess = 1` |
| IC-008986 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailORIAllotment:50](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L50) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchheaderfile2cms_tm.BitProcess` : `BIT` | `1` : `int` | `1` → `BIT` | `retail.dbo.trsbatchheaderfile2cms_tm.BitProcess <- 1` |
| IC-009008 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailORIAllotment:203](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L203) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchheaderfile2cms_tm.BitProcess` : `BIT` | `1` : `int` | `1` → `BIT` | `retail.dbo.trsbatchheaderfile2cms_tm.BitProcess <- 1` |

<a id="kolom-trxtypeid"></a>
## TrxTypeID

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBatchHeaderFile2CMS_TM.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 20 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001732 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PopulateBatchTxtFile2CMS:39](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PopulateBatchTxtFile2CMS.sql#L39) | JOIN_ON / P11 / COMPARISON | `tb.TrxTypeID` : `SMALLINT` | `5` : `int` | `tb.TrxTypeID` → `INT` | `tb.TrxTypeID = 5` |
| IC-001735 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PopulateBatchTxtFile2CMS:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PopulateBatchTxtFile2CMS.sql#L56) | JOIN_ON / P11 / IN_VALUE_TYPE | `tb.TrxTypeID` : `SMALLINT` | `7` : `int` | `tb.TrxTypeID` → `INT` | `tb.TrxTypeID in (7,8)` |
| IC-004001 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSGenerateDataForOBL009:133](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateDataForOBL009.sql#L133) | JOIN_ON / P11 / COMPARISON | `bhf.TrxTypeID` : `SMALLINT` | `5` : `int` | `bhf.TrxTypeID` → `INT` | `bhf.TrxTypeID = 5` |
| IC-008984 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailORIAllotment:35](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L35) | WHERE / P11 / COMPARISON | `TrxTypeID` : `SMALLINT` | `6` : `int` | `TrxTypeID` → `INT` | `TrxTypeID = 6` |
| IC-008985 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailORIAllotment:39](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L39) | WHERE / P11 / COMPARISON | `TrxTypeID` : `SMALLINT` | `6` : `int` | `TrxTypeID` → `INT` | `TrxTypeID = 6` |
| IC-008988 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailORIAllotment:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L63) | WHERE / P11 / COMPARISON | `TrxTypeID` : `SMALLINT` | `6` : `int` | `TrxTypeID` → `INT` | `TrxTypeID = 6` |
| IC-009005 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailORIAllotment:183](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L183) | WHERE / P11 / COMPARISON | `TrxTypeID` : `SMALLINT` | `6` : `int` | `TrxTypeID` → `INT` | `TrxTypeID = 6` |
| IC-009006 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailORIAllotment:186](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L186) | WHERE / P11 / COMPARISON | `TrxTypeID` : `SMALLINT` | `6` : `int` | `TrxTypeID` → `INT` | `TrxTypeID = 6` |
| IC-009007 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailORIAllotment:191](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L191) | WHERE / P11 / COMPARISON | `TrxTypeID` : `SMALLINT` | `6` : `int` | `TrxTypeID` → `INT` | `TrxTypeID = 6` |
| IC-009010 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailORIAllotment:217](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L217) | WHERE / P11 / COMPARISON | `TrxTypeID` : `SMALLINT` | `6` : `int` | `TrxTypeID` → `INT` | `TrxTypeID = 6` |
| IC-009109 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailRptUploadCoupon:83](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailRptUploadCoupon.sql#L83) | WHERE / P11 / COMPARISON | `bhc.TrxTypeID` : `SMALLINT` | `5` : `int` | `bhc.TrxTypeID` → `INT` | `bhc.TrxTypeID = 5` |
| IC-009586 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptCoupon:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCoupon.sql#L37) | WHERE / P11 / COMPARISON | `bhf.TrxTypeID` : `SMALLINT` | `5` : `int` | `bhf.TrxTypeID` → `INT` | `bhf.TrxTypeID = 5` |
| IC-011647 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptUploadCMS:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptUploadCMS.sql#L65) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `bhc.TrxTypeID` : `SMALLINT` | `3` : `int` | `bhc.TrxTypeID` → `INT` | `when 3 then 'Buy ' + bdc.SecurityNo` |
| IC-011648 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptUploadCMS:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptUploadCMS.sql#L66) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `bhc.TrxTypeID` : `SMALLINT` | `4` : `int` | `bhc.TrxTypeID` → `INT` | `when 4 then 'Sell ' + bdc.SecurityNo` |
| IC-011649 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptUploadCMS:67](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptUploadCMS.sql#L67) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `bhc.TrxTypeID` : `SMALLINT` | `6` : `int` | `bhc.TrxTypeID` → `INT` | `when 6 then 'Allotment ' + bdc.SecurityNo` |
| IC-011650 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptUploadCMS:68](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptUploadCMS.sql#L68) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `bhc.TrxTypeID` : `SMALLINT` | `7` : `int` | `bhc.TrxTypeID` → `INT` | `when 7 then 'SBI Due ' + bdc.SecurityNo` |
| IC-011651 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptUploadCMS:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptUploadCMS.sql#L69) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `bhc.TrxTypeID` : `SMALLINT` | `8` : `int` | `bhc.TrxTypeID` → `INT` | `when 8 then 'Bond Due ' + bdc.SecurityNo` |
| IC-011652 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptUploadCMS:76](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptUploadCMS.sql#L76) | JOIN_ON / P11 / IN_VALUE_TYPE | `bhc.TrxTypeID` : `SMALLINT` | `3` : `int` | `bhc.TrxTypeID` → `INT` | `bhc.TrxTypeID in ( 3,4,6,7,8)` |
| IC-008987 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailORIAllotment:50](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L50) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchheaderfile2cms_tm.TrxTypeID` : `SMALLINT` | `6` : `int` | `6` → `SMALLINT` | `retail.dbo.trsbatchheaderfile2cms_tm.TrxTypeID <- 6` |
| IC-009009 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailORIAllotment:203](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L203) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchheaderfile2cms_tm.TrxTypeID` : `SMALLINT` | `6` : `int` | `6` → `SMALLINT` | `retail.dbo.trsbatchheaderfile2cms_tm.TrxTypeID <- 6` |
