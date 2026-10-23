# Retail UAT.dbo.ORIOnlineGeneralParam_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 31 lokasi source, 17 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 25 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-groupid"></a>
## GroupId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORIOnlineGeneralParam_TR.sql#L2). **Existing leading index:** ORIOnlineGeneralParam_TR_CN1.

**Frekuensi source:** 6 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000992 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.ORISBNIPopulateRegisterInvestor:139](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateRegisterInvestor.sql#L139) | JOIN_ON / P04 / COMPARISON | `b.GroupId` : `INT` | `'9000'` : `varchar` | `'9000'` → `INT` | `b.GroupId = '9000'` |
| IC-000993 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.ORISBNIPopulateRegisterInvestor:146](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateRegisterInvestor.sql#L146) | JOIN_ON / P04 / COMPARISON | `b.GroupId` : `INT` | `'7000'` : `varchar` | `'7000'` → `INT` | `b.GroupId = '7000'` |
| IC-002798 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedPopulateNotif:8](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateNotif.sql#L8) | WHERE / P04 / COMPARISON | `[GroupId]` : `INT` | `'8000'` : `varchar` | `'8000'` → `INT` | `[GroupId] = '8000'` |
| IC-002799 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedPopulateNotif:9](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateNotif.sql#L9) | WHERE / P04 / COMPARISON | `[GroupId]` : `INT` | `'8000'` : `varchar` | `'8000'` → `INT` | `[GroupId] = '8000'` |
| IC-004876 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSObliGetAPI:11](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSObliGetAPI.sql#L11) | WHERE / P04 / COMPARISON | `GroupId` : `INT` | `'2000'` : `varchar` | `'2000'` → `INT` | `GroupId = '2000'` |
| IC-004877 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSObliGetAPI:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSObliGetAPI.sql#L16) | WHERE / P04 / COMPARISON | `GroupId` : `INT` | `'2000'` : `varchar` | `'2000'` → `INT` | `GroupId = '2000'` |

<a id="kolom-valueid"></a>
## ValueId

**Tipe dasar:** `VARCHAR(100)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORIOnlineGeneralParam_TR.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 25 lokasi / 14 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002787 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedGetDataStatusPemesanan:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedGetDataStatusPemesanan.sql#L23) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `ValueId` : `VARCHAR(100)` | `@nDays` : `int` | `ValueId` → `int` | `@nDays = ValueId` |
| IC-002792 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedGetGeneralParam:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedGetGeneralParam.sql#L17) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[ValueId]` : `VARCHAR(100)` | `@nMaxRetry` : `int` | `[ValueId]` → `int` | `@nMaxRetry = [ValueId]` |
| IC-002793 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedGetGeneralParam:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedGetGeneralParam.sql#L21) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[ValueId]` : `VARCHAR(100)` | `@nIntervalInvstrRegis` : `int` | `[ValueId]` → `int` | `@nIntervalInvstrRegis = [ValueId]` |
| IC-002794 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedGetGeneralParam:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedGetGeneralParam.sql#L25) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[ValueId]` : `VARCHAR(100)` | `@nIntervalOrderStatus` : `int` | `[ValueId]` → `int` | `@nIntervalOrderStatus = [ValueId]` |
| IC-002795 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedGetGeneralParam:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedGetGeneralParam.sql#L29) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[ValueId]` : `VARCHAR(100)` | `@nIntervalOrderPerdana` : `int` | `[ValueId]` → `int` | `@nIntervalOrderPerdana = [ValueId]` |
| IC-002796 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedGetGeneralParam:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedGetGeneralParam.sql#L33) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[ValueId]` : `VARCHAR(100)` | `@nIntervalNotif` : `int` | `[ValueId]` → `int` | `@nIntervalNotif = [ValueId]` |
| IC-002797 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedGetGeneralParam:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedGetGeneralParam.sql#L37) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[ValueId]` : `VARCHAR(100)` | `@nIntervalLogging` : `int` | `[ValueId]` → `int` | `@nIntervalLogging = [ValueId]` |
| IC-002809 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPopulateORIOrder:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateORIOrder.sql#L69) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[ValueId]` : `VARCHAR(100)` | `@bAutoPay` : `bit` | `[ValueId]` → `bit` | `@bAutoPay = [ValueId]` |
| IC-002810 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPopulateORIOrder:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateORIOrder.sql#L77) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `ValueId` : `VARCHAR(100)` | `@nAPIPaymentSwTimeout` : `int` | `ValueId` → `int` | `@nAPIPaymentSwTimeout = ValueId` |
| IC-002904 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPushMailNotifOrderCompleted:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L70) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[ValueId]` : `VARCHAR(100)` | `@nMaxRetry` : `int` | `[ValueId]` → `int` | `@nMaxRetry = [ValueId]` |
| IC-002929 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPushMailNotifOrderORI:67](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderORI.sql#L67) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[ValueId]` : `VARCHAR(100)` | `@nMaxRetry` : `int` | `[ValueId]` → `int` | `@nMaxRetry = [ValueId]` |
| IC-002949 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPushMailNotifRegInvestor:55](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifRegInvestor.sql#L55) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[ValueId]` : `VARCHAR(100)` | `@nMaxRetry` : `int` | `[ValueId]` → `int` | `@nMaxRetry = [ValueId]` |
| IC-002958 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPushMailNotifSettleRedempt:42](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifSettleRedempt.sql#L42) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[ValueId]` : `VARCHAR(100)` | `@nMaxRetry` : `int` | `[ValueId]` → `int` | `@nMaxRetry = [ValueId]` |
| IC-002963 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedUpdateResponseInvestor:15](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedUpdateResponseInvestor.sql#L15) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[ValueId]` : `VARCHAR(100)` | `@nMaxRetry` : `int` | `[ValueId]` → `int` | `@nMaxRetry = [ValueId]` |
| IC-002964 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedUpdateResponseOrderORI:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedUpdateResponseOrderORI.sql#L23) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[ValueId]` : `VARCHAR(100)` | `@nMaxRetry` : `int` | `[ValueId]` → `int` | `@nMaxRetry = [ValueId]` |
| IC-002965 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedUpdateResponseOrderORI:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedUpdateResponseOrderORI.sql#L27) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[ValueId]` : `VARCHAR(100)` | `@bAutoPay` : `bit` | `[ValueId]` → `bit` | `@bAutoPay = [ValueId]` |
| IC-004752 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSNTIPushMailNotif:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSNTIPushMailNotif.sql#L37) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[ValueId]` : `VARCHAR(100)` | `@nMaxRetry` : `int` | `[ValueId]` → `int` | `@nMaxRetry = [ValueId]` |
| IC-004794 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSORIOnlineNewTransaction:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlineNewTransaction.sql#L108) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[ValueId]` : `VARCHAR(100)` | `@bAutoPay` : `bit` | `[ValueId]` → `bit` | `@bAutoPay = [ValueId]` |
| IC-004825 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSORIOnlinePopulateHist:62](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L62) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[ValueId]` : `VARCHAR(100)` | `@dHistFetchDt` : `datetime` | `[ValueId]` → `datetime` | `@dHistFetchDt = [ValueId]` |
| IC-004922 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSObligasiHousekeepProcess:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSObligasiHousekeepProcess.sql#L29) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[ValueId]` : `VARCHAR(100)` | `@nHKDaysLiveLog` : `int` | `[ValueId]` → `int` | `@nHKDaysLiveLog = [ValueId]` |
| IC-004923 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSObligasiHousekeepProcess:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSObligasiHousekeepProcess.sql#L30) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[ValueId]` : `VARCHAR(100)` | `@nHKDaysGeneralLog` : `int` | `[ValueId]` → `int` | `@nHKDaysGeneralLog = [ValueId]` |
| IC-004924 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSObligasiHousekeepProcess:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSObligasiHousekeepProcess.sql#L31) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[ValueId]` : `VARCHAR(100)` | `@nHKDaysModerateLog` : `int` | `[ValueId]` → `int` | `@nHKDaysModerateLog = [ValueId]` |
| IC-004925 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSObligasiHousekeepProcess:32](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSObligasiHousekeepProcess.sql#L32) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[ValueId]` : `VARCHAR(100)` | `@nHKWorkflow` : `int` | `[ValueId]` → `int` | `@nHKWorkflow = [ValueId]` |
| IC-004926 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSObligasiHousekeepProcess:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSObligasiHousekeepProcess.sql#L33) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[ValueId]` : `VARCHAR(100)` | `@nHKNotif` : `int` | `[ValueId]` → `int` | `@nHKNotif = [ValueId]` |
| IC-005443 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateGagalBayarOri:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L22) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `ValueId` : `VARCHAR(100)` | `@pParameterCutOffNotProcessed` : `INT` | `ValueId` → `INT` | `@pParameterCutOffNotProcessed = ValueId` |
