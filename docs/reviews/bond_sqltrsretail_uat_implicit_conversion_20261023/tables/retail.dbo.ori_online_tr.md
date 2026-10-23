# Retail UAT.dbo.ORI_Online_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 69 lokasi source, 24 procedure, 8 kolom. **A:** 18 lokasi pada 10 procedure. **B:** 29 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-kemenkeustatusid"></a>
## KemenkeuStatusId

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Online_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 44 lokasi / 18 SP. **A:** 18 lokasi / 10 SP. **Peran:** CONVERTED_ARGUMENT, CONVERTED_COLUMN, DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001790 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.RPAJurnalSettlementOriva:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RPAJurnalSettlementOriva.sql#L105) | WHERE / P01 / COMPARISON | `ori.KemenkeuStatusId` : `VARCHAR(20)` | `3` : `int` | `ori.KemenkeuStatusId` → `INT` | `ori.KemenkeuStatusId = 3` |
| IC-002780 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.SBNSchedCheckAlert:20](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedCheckAlert.sql#L20) | WHERE / P01 / COMPARISON | `a.KemenkeuStatusId` : `VARCHAR(20)` | `5` : `int` | `a.KemenkeuStatusId` → `INT` | `a.KemenkeuStatusId = 5` |
| IC-002782 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.SBNSchedCheckAlert:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedCheckAlert.sql#L34) | WHERE / P01 / COMPARISON | `a.KemenkeuStatusId` : `VARCHAR(20)` | `3` : `int` | `a.KemenkeuStatusId` → `INT` | `a.KemenkeuStatusId = 3` |
| IC-002805 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.SBNSchedPopulateNotif:57](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateNotif.sql#L57) | WHERE / P01 / IN_VALUE_TYPE | `tr.KemenkeuStatusId` : `VARCHAR(20)` | `4` : `int` | `tr.KemenkeuStatusId` → `INT` | `tr.KemenkeuStatusId IN (4,5)` |
| IC-002806 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.SBNSchedPopulateNotif:58](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateNotif.sql#L58) | WHERE / P01 / COMPARISON | `tr.KemenkeuStatusId` : `VARCHAR(20)` | `99` : `int` | `tr.KemenkeuStatusId` → `INT` | `tr.KemenkeuStatusId = 99` |
| IC-002812 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.SBNSchedPopulateORIOrder:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateORIOrder.sql#L228) | WHERE / P01 / COMPARISON | `ori.KemenkeuStatusId` : `VARCHAR(20)` | `3` : `int` | `ori.KemenkeuStatusId` → `INT` | `ori.KemenkeuStatusId = 3` |
| IC-002813 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.SBNSchedPopulateORIOrder:267](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateORIOrder.sql#L267) | WHERE / P01 / COMPARISON | `ori.KemenkeuStatusId` : `VARCHAR(20)` | `3` : `int` | `ori.KemenkeuStatusId` → `INT` | `ori.KemenkeuStatusId = 3` |
| IC-002912 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.SBNSchedPushMailNotifOrderCompleted:176](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L176) | WHERE / P01 / COMPARISON | `KemenkeuStatusId` : `VARCHAR(20)` | `4` : `int` | `KemenkeuStatusId` → `INT` | `KemenkeuStatusId = 4` |
| IC-002917 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.SBNSchedPushMailNotifOrderCompleted:332](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L332) | WHERE / P01 / COMPARISON | `KemenkeuStatusId` : `VARCHAR(20)` | `5` : `int` | `KemenkeuStatusId` → `INT` | `KemenkeuStatusId = 5` |
| IC-002936 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.SBNSchedPushMailNotifOrderORI:175](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderORI.sql#L175) | WHERE / P01 / COMPARISON | `KemenkeuStatusId` : `VARCHAR(20)` | `4` : `int` | `KemenkeuStatusId` → `INT` | `KemenkeuStatusId = 4` |
| IC-004785 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSORIGetProductDetail:83](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIGetProductDetail.sql#L83) | WHERE / P01 / IN_VALUE_TYPE | `KemenkeuStatusId` : `VARCHAR(20)` | `0` : `int` | `KemenkeuStatusId` → `INT` | `KemenkeuStatusId IN (0, 1, 3)` |
| IC-004873 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSORIOrderSummary:100](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOrderSummary.sql#L100) | WHERE / P01 / COMPARISON | `onl.KemenkeuStatusId` : `VARCHAR(20)` | `4` : `int` | `onl.KemenkeuStatusId` → `INT` | `onl.KemenkeuStatusId = 4` |
| IC-005449 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateGagalBayarOri:155](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L155) | WHERE / P01 / IN_VALUE_TYPE | `trx.KemenkeuStatusId` : `VARCHAR(20)` | `5` : `int` | `trx.KemenkeuStatusId` → `INT` | `trx.KemenkeuStatusId NOT IN (5)` |
| IC-005455 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateGagalBayarOri:245](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L245) | WHERE / P01 / IN_VALUE_TYPE | `trx.KemenkeuStatusId` : `VARCHAR(20)` | `5` : `int` | `trx.KemenkeuStatusId` → `INT` | `trx.KemenkeuStatusId NOT IN (5)` |
| IC-005462 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateGagalBayarOri:354](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L354) | WHERE / P01 / IN_VALUE_TYPE | `trx.KemenkeuStatusId` : `VARCHAR(20)` | `4` : `int` | `trx.KemenkeuStatusId` → `INT` | `trx.KemenkeuStatusId NOT IN (4,5)` |
| IC-005471 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateGagalBayarOri:478](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L478) | WHERE / P01 / IN_VALUE_TYPE | `trx.KemenkeuStatusId` : `VARCHAR(20)` | `3` : `int` | `trx.KemenkeuStatusId` → `INT` | `trx.KemenkeuStatusId NOT IN (3,4,5)` |
| IC-013868 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.trs_UpdateTreasuryCustomer_TM:638](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L638) | WHERE / P01 / IN_VALUE_TYPE | `orion.KemenkeuStatusId` : `VARCHAR(20)` | `5` : `int` | `orion.KemenkeuStatusId` → `INT` | `orion.KemenkeuStatusId NOT IN (5,99)` |
| IC-013889 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.trs_UpdateTreasuryCustomer_TM:755](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L755) | WHERE / P01 / IN_VALUE_TYPE | `orion.KemenkeuStatusId` : `VARCHAR(20)` | `5` : `int` | `orion.KemenkeuStatusId` → `INT` | `orion.KemenkeuStatusId NOT IN (5,99)` |
| IC-000944 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.ORIOnlineBrowseTrx:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORIOnlineBrowseTrx.sql#L98) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `trx.KemenkeuStatusId` : `VARCHAR(20)` | `3` : `int` | `trx.KemenkeuStatusId` → `INT` | `when 3 then 'Verified Order'` |
| IC-000945 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.ORIOnlineBrowseTrx:99](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORIOnlineBrowseTrx.sql#L99) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `trx.KemenkeuStatusId` : `VARCHAR(20)` | `4` : `int` | `trx.KemenkeuStatusId` → `INT` | `when 4 then 'Completed Order'` |
| IC-000946 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.ORIOnlineBrowseTrx:100](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORIOnlineBrowseTrx.sql#L100) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `trx.KemenkeuStatusId` : `VARCHAR(20)` | `5` : `int` | `trx.KemenkeuStatusId` → `INT` | `when 5 then 'Unpaid Order'` |
| IC-000948 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.ORIOnlineReleaseFailedUnblock:19](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORIOnlineReleaseFailedUnblock.sql#L19) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `KemenkeuStatusId` : `VARCHAR(20)` | `@nKemenkeuStatusId` : `int` | `KemenkeuStatusId` → `int` | `@nKemenkeuStatusId = KemenkeuStatusId` |
| IC-000950 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.ORIOnlineReportDailyTransaction:177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORIOnlineReportDailyTransaction.sql#L177) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `trx.KemenkeuStatusId` : `VARCHAR(20)` | `3` : `int` | `trx.KemenkeuStatusId` → `INT` | `when 3 then 'Verified Order'` |
| IC-000951 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.ORIOnlineReportDailyTransaction:178](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORIOnlineReportDailyTransaction.sql#L178) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `trx.KemenkeuStatusId` : `VARCHAR(20)` | `4` : `int` | `trx.KemenkeuStatusId` → `INT` | `when 4 then 'Completed Order'` |
| IC-000952 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.ORIOnlineReportDailyTransaction:179](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORIOnlineReportDailyTransaction.sql#L179) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `trx.KemenkeuStatusId` : `VARCHAR(20)` | `5` : `int` | `trx.KemenkeuStatusId` → `INT` | `when 5 then 'Unpaid Order'` |
| IC-000962 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.ORIOnlineReportDailyTrx:14](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORIOnlineReportDailyTrx.sql#L14) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `onl.KemenkeuStatusId` : `VARCHAR(20)` | `retail.dbo.fnoblgetoriorderstatus.@nKemenkeuStatusId` : `int` | `onl.KemenkeuStatusId` → `int` | `dbo.fnOBLGetORIOrderStatus(onl.KemenkeuStatusId)` |
| IC-002804 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.SBNSchedPopulateNotif:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateNotif.sql#L48) | CASE / P01 / COMPARISON | `KemenkeuStatusId` : `VARCHAR(20)` | `4` : `int` | `KemenkeuStatusId` → `INT` | `KemenkeuStatusId = 4` |
| IC-004853 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSORIOnlinePopulateHist:231](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L231) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `ol.[KemenkeuStatusId]` : `VARCHAR(20)` | `retail.dbo.fnoblgetoriorderstatus.@nKemenkeuStatusId` : `int` | `ol.[KemenkeuStatusId]` → `int` | `dbo.fnOBLGetORIOrderStatus(ol.[KemenkeuStatusId])` |
| IC-004854 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSORIOnlinePopulateHist:237](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L237) | CASE / P01 / COMPARISON | `ol.[KemenkeuStatusId]` : `VARCHAR(20)` | `4` : `int` | `ol.[KemenkeuStatusId]` → `INT` | `ol.[KemenkeuStatusId] = 4` |
| IC-005446 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateGagalBayarOri:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L130) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `trx.KemenkeuStatusId` : `VARCHAR(20)` | `3` : `int` | `trx.KemenkeuStatusId` → `INT` | `WHEN 3 THEN 'Verified Order'` |
| IC-005447 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateGagalBayarOri:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L131) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `trx.KemenkeuStatusId` : `VARCHAR(20)` | `4` : `int` | `trx.KemenkeuStatusId` → `INT` | `WHEN 4 THEN 'Completed Order'` |
| IC-005448 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateGagalBayarOri:132](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L132) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `trx.KemenkeuStatusId` : `VARCHAR(20)` | `5` : `int` | `trx.KemenkeuStatusId` → `INT` | `WHEN 5 THEN 'Unpaid Order'` |
| IC-005452 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateGagalBayarOri:217](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L217) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `trx.KemenkeuStatusId` : `VARCHAR(20)` | `3` : `int` | `trx.KemenkeuStatusId` → `INT` | `WHEN 3 THEN 'Verified Order'` |
| IC-005453 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateGagalBayarOri:218](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L218) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `trx.KemenkeuStatusId` : `VARCHAR(20)` | `4` : `int` | `trx.KemenkeuStatusId` → `INT` | `WHEN 4 THEN 'Completed Order'` |
| IC-005454 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateGagalBayarOri:219](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L219) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `trx.KemenkeuStatusId` : `VARCHAR(20)` | `5` : `int` | `trx.KemenkeuStatusId` → `INT` | `WHEN 5 THEN 'Unpaid Order'` |
| IC-005459 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateGagalBayarOri:329](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L329) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `trx.KemenkeuStatusId` : `VARCHAR(20)` | `3` : `int` | `trx.KemenkeuStatusId` → `INT` | `WHEN 3 THEN 'Verified Order'` |
| IC-005460 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateGagalBayarOri:330](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L330) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `trx.KemenkeuStatusId` : `VARCHAR(20)` | `4` : `int` | `trx.KemenkeuStatusId` → `INT` | `WHEN 4 THEN 'Completed Order'` |
| IC-005461 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateGagalBayarOri:331](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L331) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `trx.KemenkeuStatusId` : `VARCHAR(20)` | `5` : `int` | `trx.KemenkeuStatusId` → `INT` | `WHEN 5 THEN 'Unpaid Order'` |
| IC-005468 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateGagalBayarOri:453](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L453) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `trx.KemenkeuStatusId` : `VARCHAR(20)` | `3` : `int` | `trx.KemenkeuStatusId` → `INT` | `WHEN 3 THEN 'Verified Order'` |
| IC-005469 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateGagalBayarOri:454](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L454) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `trx.KemenkeuStatusId` : `VARCHAR(20)` | `4` : `int` | `trx.KemenkeuStatusId` → `INT` | `WHEN 4 THEN 'Completed Order'` |
| IC-005470 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateGagalBayarOri:455](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L455) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `trx.KemenkeuStatusId` : `VARCHAR(20)` | `5` : `int` | `trx.KemenkeuStatusId` → `INT` | `WHEN 5 THEN 'Unpaid Order'` |
| IC-002966 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedUpdateResponseOrderORI:203](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedUpdateResponseOrderORI.sql#L203) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `KemenkeuStatusId` : `VARCHAR(20)` | `99` : `int` | `99` → `VARCHAR(20)` | `KemenkeuStatusId = 99` |
| IC-004819 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSORIOnlineNewTransaction:394](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlineNewTransaction.sql#L394) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_online_tr.KemenkeuStatusId` : `VARCHAR(20)` | `0` : `int` | `0` → `VARCHAR(20)` | `retail.dbo.ori_online_tr.KemenkeuStatusId <- 0` |
| IC-004875 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSORIUpdateStatusPembayaranOnline:32](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIUpdateStatusPembayaranOnline.sql#L32) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `KemenkeuStatusId` : `VARCHAR(20)` | `4` : `int` | `4` → `VARCHAR(20)` | `KemenkeuStatusId = 4` |

<a id="kolom-cifid"></a>
## CIFId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Online_TR.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002905 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPushMailNotifOrderCompleted:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L115) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[CIFId]` : `BIGINT` | `@nCIFId` : `int` | `[CIFId]` → `int` | `@nCIFId = [CIFId]` |
| IC-002930 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPushMailNotifOrderORI:117](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderORI.sql#L117) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[CIFId]` : `BIGINT` | `@nCIFId` : `int` | `[CIFId]` → `int` | `@nCIFId = [CIFId]` |

<a id="kolom-orderid"></a>
## OrderId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Online_TR.sql#L2). **Existing leading index:** ORI_Online_TR_CU1&#124;ORI_Online_TR_NN2.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000964 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.ORISBNIDeleteTransactionData:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIDeleteTransactionData.sql#L25) | WHERE / P11 / COMPARISON | `OrderId` : `BIGINT` | `@nOrderId` : `int` | `@nOrderId` → `BIGINT` | `OrderId = @nOrderId` |

<a id="kolom-orivadebetflag"></a>
## OrivaDebetFlag

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Online_TR.sql#L26). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN_EXPRESSION, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001788 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.RPAJurnalSettlementOriva:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RPAJurnalSettlementOriva.sql#L104) | WHERE / P11 / COMPARISON | `ISNULL(ori.OrivaDebetFlag, 0)` : `BIT` | `0` : `int` | `ISNULL(ori.OrivaDebetFlag, 0)` → `INT` | `ISNULL(ori.OrivaDebetFlag, 0) = 0` |
| IC-001789 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.RPAJurnalSettlementOriva:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RPAJurnalSettlementOriva.sql#L104) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ori.OrivaDebetFlag` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(ori.OrivaDebetFlag, 0)` |
| IC-012560 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdatePostingResult:314](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L314) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `OrivaDebetFlag` : `BIT` | `1` : `int` | `1` → `BIT` | `OrivaDebetFlag = 1` |

<a id="kolom-orivaflag"></a>
## OrivaFlag

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Online_TR.sql#L25). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001040 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OrivaTransactionReport:90](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OrivaTransactionReport.sql#L90) | WHERE / P11 / COMPARISON | `ori.OrivaFlag` : `BIT` | `1` : `int` | `ori.OrivaFlag` → `INT` | `ori.OrivaFlag = 1` |
| IC-001787 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.RPAJurnalSettlementOriva:103](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RPAJurnalSettlementOriva.sql#L103) | WHERE / P11 / COMPARISON | `ori.OrivaFlag` : `BIT` | `1` : `int` | `ori.OrivaFlag` → `INT` | `ori.OrivaFlag = 1` |
| IC-002783 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedCheckAlert:35](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedCheckAlert.sql#L35) | WHERE / P11 / COMPARISON | `a.OrivaFlag` : `BIT` | `0` : `int` | `a.OrivaFlag` → `INT` | `a.OrivaFlag = 0` |
| IC-002788 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedGetDataStatusPemesanan:54](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedGetDataStatusPemesanan.sql#L54) | WHERE / P11 / COMPARISON | `ori.OrivaFlag` : `BIT` | `0` : `int` | `ori.OrivaFlag` → `INT` | `ori.OrivaFlag = 0` |
| IC-002789 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedGetDataStatusPemesanan:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedGetDataStatusPemesanan.sql#L70) | WHERE / P11 / COMPARISON | `onl.OrivaFlag` : `BIT` | `0` : `int` | `onl.OrivaFlag` → `INT` | `onl.OrivaFlag = 0` |
| IC-002790 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedGetDataStatusPemesanan:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedGetDataStatusPemesanan.sql#L89) | WHERE / P11 / COMPARISON | `ori.OrivaFlag` : `BIT` | `0` : `int` | `ori.OrivaFlag` → `INT` | `ori.OrivaFlag = 0` |
| IC-002791 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedGetDataStatusPemesanan:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedGetDataStatusPemesanan.sql#L105) | WHERE / P11 / COMPARISON | `onl.OrivaFlag` : `BIT` | `0` : `int` | `onl.OrivaFlag` → `INT` | `onl.OrivaFlag = 0` |
| IC-002808 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedPopulateNotif:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateNotif.sql#L64) | WHERE / P11 / COMPARISON | `tr.OrivaFlag` : `BIT` | `0` : `int` | `tr.OrivaFlag` → `INT` | `tr.OrivaFlag = 0` |

<a id="kolom-secid"></a>
## SecId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Online_TR.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005445 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateGagalBayarOri:72](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L72) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `trx.SecId` : `INT` | `#tmp_data.SecId` : `BIGINT` | `trx.SecId` → `BIGINT` | `#tmp_data.SecId <- trx.SecId` |
| IC-005451 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateGagalBayarOri:159](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L159) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `trx.SecId` : `INT` | `#tmp_data.SecId` : `BIGINT` | `trx.SecId` → `BIGINT` | `#tmp_data.SecId <- trx.SecId` |
| IC-005458 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateGagalBayarOri:271](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L271) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `trx.SecId` : `INT` | `#tmp_data.SecId` : `BIGINT` | `trx.SecId` → `BIGINT` | `#tmp_data.SecId <- trx.SecId` |
| IC-005467 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateGagalBayarOri:395](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L395) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `trx.SecId` : `INT` | `#tmp_data.SecId` : `BIGINT` | `trx.SecId` → `BIGINT` | `#tmp_data.SecId <- trx.SecId` |

<a id="kolom-settlebit"></a>
## SettleBit

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Online_TR.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000968 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.ORISBNIPopulateEarlyRedemption:78](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateEarlyRedemption.sql#L78) | WHERE / P11 / COMPARISON | `b.SettleBit` : `BIT` | `1` : `int` | `b.SettleBit` → `INT` | `b.SettleBit = 1` |
| IC-000984 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.ORISBNIPopulateEarlyRedemption:226](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateEarlyRedemption.sql#L226) | WHERE / P11 / COMPARISON | `c.SettleBit` : `BIT` | `1` : `int` | `c.SettleBit` → `INT` | `c.SettleBit = 1` |
| IC-003353 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSAutoSettlementOnlineORI:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAutoSettlementOnlineORI.sql#L77) | WHERE / P11 / COMPARISON | `ISNULL(SettleBit, 0)` : `BIT` | `0` : `int` | `ISNULL(SettleBit, 0)` → `INT` | `ISNULL(SettleBit, 0) = 0` |
| IC-003354 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSAutoSettlementOnlineORI:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAutoSettlementOnlineORI.sql#L77) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `SettleBit` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(SettleBit, 0)` |
| IC-003364 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSAutoSettlementOnlineORI:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAutoSettlementOnlineORI.sql#L162) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `onl.[SettleBit]` : `BIT` | `1` : `int` | `1` → `BIT` | `onl.[SettleBit] = 1` |
| IC-004820 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSORIOnlineNewTransaction:394](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlineNewTransaction.sql#L394) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_online_tr.SettleBit` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.ori_online_tr.SettleBit <- 0` |

<a id="kolom-tglpembayaran"></a>
## TglPembayaran

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Online_TR.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004874 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSORIUpdateStatusPembayaranOnline:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIUpdateStatusPembayaranOnline.sql#L29) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `TglPembayaran` : `DATETIME` | `@pcTglPembayaran` : `varchar(30)` | `@pcTglPembayaran` → `DATETIME` | `TglPembayaran = @pcTglPembayaran` |
