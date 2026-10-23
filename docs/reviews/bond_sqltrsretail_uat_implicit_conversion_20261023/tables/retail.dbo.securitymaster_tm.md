# Retail UAT.dbo.SecurityMaster_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 788 lokasi source, 176 procedure, 61 kolom. **A:** 25 lokasi pada 11 procedure. **B:** 107 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-idseri"></a>
## IdSeri

**Tipe dasar:** `VARCHAR(100)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L99). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 19 lokasi / 9 SP. **A:** 18 lokasi / 8 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000966 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.ORISBNIPopulateEarlyRedemption:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateEarlyRedemption.sql#L33) | WHERE / P01 / COMPARISON | `IdSeri` : `VARCHAR(100)` | `@pnSeriId` : `INT` | `IdSeri` → `INT` | `IdSeri = @pnSeriId` |
| IC-000967 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.ORISBNIPopulateEarlyRedemption:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateEarlyRedemption.sql#L61) | WHERE / P01 / COMPARISON | `IdSeri` : `VARCHAR(100)` | `@pnSeriId` : `INT` | `IdSeri` → `INT` | `IdSeri=@pnSeriId` |
| IC-000973 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.ORISBNIPopulateEarlyRedemption:117](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateEarlyRedemption.sql#L117) | WHERE / P01 / COMPARISON | `IdSeri` : `VARCHAR(100)` | `@pnSeriId` : `INT` | `IdSeri` → `INT` | `IdSeri = @pnSeriId` |
| IC-000996 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.ORISBNIPopulateRegisterInvestor:196](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateRegisterInvestor.sql#L196) | JOIN_ON / P01 / COMPARISON | `tm.IdSeri` : `VARCHAR(100)` | `iv.IdSeri` : `int` | `tm.IdSeri` → `INT` | `tm.IdSeri = iv.IdSeri` |
| IC-001026 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.ORISBNISubmitInvestorPending:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNISubmitInvestorPending.sql#L162) | JOIN_ON / P01 / COMPARISON | `a.IdSeri` : `VARCHAR(100)` | `b.IdSeri` : `int` | `a.IdSeri` → `INT` | `a.IdSeri = b.IdSeri` |
| IC-002811 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.SBNSchedPopulateORIOrder:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateORIOrder.sql#L110) | JOIN_ON / P01 / COMPARISON | `tm.IdSeri` : `VARCHAR(100)` | `ms.IdSeri` : `INT` | `tm.IdSeri` → `INT` | `ms.IdSeri = tm.IdSeri` |
| IC-002818 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.SBNSchedPopulateORIOrder:400](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateORIOrder.sql#L400) | JOIN_ON / P01 / COMPARISON | `sm.IdSeri` : `VARCHAR(100)` | `a.[IdSeri]` : `int` | `sm.IdSeri` → `INT` | `a.[IdSeri] = sm.IdSeri` |
| IC-002835 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.SBNSchedPopulateRegInvestorForOrder:83](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestorForOrder.sql#L83) | JOIN_ON / P01 / COMPARISON | `tm.IdSeri` : `VARCHAR(100)` | `scd.IdSeri` : `INT` | `tm.IdSeri` → `INT` | `scd.IdSeri = tm.IdSeri` |
| IC-004781 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSORICheckProductEligibility:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORICheckProductEligibility.sql#L64) | WHERE / P01 / COMPARISON | `IdSeri` : `VARCHAR(100)` | `@pnIdSeri` : `int` | `IdSeri` → `INT` | `IdSeri = @pnIdSeri` |
| IC-004782 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSORIGetProductDetail:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIGetProductDetail.sql#L43) | WHERE / P01 / COMPARISON | `[IdSeri]` : `VARCHAR(100)` | `@pnIdSeri` : `int` | `[IdSeri]` → `INT` | `[IdSeri] = @pnIdSeri` |
| IC-004783 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSORIGetProductDetail:54](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIGetProductDetail.sql#L54) | WHERE / P01 / COMPARISON | `[IdSeri]` : `VARCHAR(100)` | `@pnIdSeri` : `int` | `[IdSeri]` → `INT` | `[IdSeri] = @pnIdSeri` |
| IC-004784 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSORIGetProductDetail:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIGetProductDetail.sql#L77) | WHERE / P01 / COMPARISON | `[IdSeri]` : `VARCHAR(100)` | `@pnIdSeri` : `int` | `[IdSeri]` → `INT` | `[IdSeri] = @pnIdSeri` |
| IC-004793 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSORIGetProductDetail:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIGetProductDetail.sql#L130) | WHERE / P01 / COMPARISON | `[IdSeri]` : `VARCHAR(100)` | `@pnIdSeri` : `int` | `[IdSeri]` → `INT` | `[IdSeri] = @pnIdSeri` |
| IC-004797 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSORIOnlineNewTransaction:170](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlineNewTransaction.sql#L170) | JOIN_ON / P01 / COMPARISON | `sm.IdSeri` : `VARCHAR(100)` | `tmp.[IdSeri]` : `int` | `sm.IdSeri` → `INT` | `tmp.[IdSeri] = sm.IdSeri` |
| IC-004802 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSORIOnlineNewTransaction:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlineNewTransaction.sql#L192) | JOIN_ON / P01 / COMPARISON | `sm.IdSeri` : `VARCHAR(100)` | `tmp.[IdSeri]` : `int` | `sm.IdSeri` → `INT` | `tmp.[IdSeri] = sm.IdSeri` |
| IC-004803 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSORIOnlineNewTransaction:197](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlineNewTransaction.sql#L197) | WHERE / P01 / COMPARISON | `[IdSeri]` : `VARCHAR(100)` | `@nIdSeri` : `int` | `[IdSeri]` → `INT` | `[IdSeri] = @nIdSeri` |
| IC-004804 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSORIOnlineNewTransaction:219](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlineNewTransaction.sql#L219) | WHERE / P01 / COMPARISON | `[IdSeri]` : `VARCHAR(100)` | `@nIdSeri` : `int` | `[IdSeri]` → `INT` | `[IdSeri] = @nIdSeri` |
| IC-004816 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSORIOnlineNewTransaction:306](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlineNewTransaction.sql#L306) | WHERE / P01 / COMPARISON | `IdSeri` : `VARCHAR(100)` | `@nIdSeri` : `int` | `IdSeri` → `INT` | `IdSeri = @nIdSeri` |
| IC-004861 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSORIOnlinePopulateHist:287](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L287) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `sm.[IdSeri]` : `VARCHAR(100)` | `tmp.[IdSeri]` : `int` | `sm.[IdSeri]` → `int` | `tmp.[IdSeri] = sm.[IdSeri]` |

<a id="kolom-isincode"></a>
## ISINCode

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 6 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001880 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.RepoRONDataMigrate:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L71) | JOIN_ON / P02 / COMPARISON | `sm.ISINCode` : `VARCHAR(20)` | `rmt.ISINCode` : `NVARCHAR(15)` | `sm.ISINCode` → `NVARCHAR` | `rmt.ISINCode = sm.ISINCode` |
| IC-008278 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:159](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L159) | JOIN_ON / P02 / COMPARISON | `sttm.ISINCode` : `VARCHAR(20)` | `mxsec.M_SE_CODE` : `NVARCHAR(15)` | `sttm.ISINCode` → `NVARCHAR` | `sttm.ISINCode = mxsec.M_SE_CODE` |
| IC-008298 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:260](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L260) | JOIN_ON / P02 / COMPARISON | `sttm.ISINCode` : `VARCHAR(20)` | `mxsec.M_SE_CODE` : `NVARCHAR(15)` | `sttm.ISINCode` → `NVARCHAR` | `sttm.ISINCode = mxsec.M_SE_CODE` |
| IC-008318 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:362](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L362) | JOIN_ON / P02 / COMPARISON | `sttm.ISINCode` : `VARCHAR(20)` | `mxsec.M_SE_CODE` : `NVARCHAR(15)` | `sttm.ISINCode` → `NVARCHAR` | `sttm.ISINCode = mxsec.M_SE_CODE` |
| IC-008338 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:463](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L463) | JOIN_ON / P02 / COMPARISON | `sttm.ISINCode` : `VARCHAR(20)` | `mxsec.M_SE_CODE` : `NVARCHAR(15)` | `sttm.ISINCode` → `NVARCHAR` | `sttm.ISINCode = mxsec.M_SE_CODE` |
| IC-008373 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:627](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L627) | JOIN_ON / P02 / COMPARISON | `sttm.ISINCode` : `VARCHAR(20)` | `mxsec.M_SE_CODE` : `NVARCHAR(15)` | `sttm.ISINCode` → `NVARCHAR` | `sttm.ISINCode = mxsec.M_SE_CODE` |

<a id="kolom-securityno"></a>
## SecurityNo

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 21 lokasi / 2 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P06 — LIKE dan representasi karakter:** LIKE pada key numerik memerlukan representasi karakter. Equality numeric yang tervalidasi cocok untuk lookup tepat; prefix/substring membutuhkan requirement pencarian tersendiri, bukan cast kolom sebagai solusi index.
- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005739 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateMinimumRiskProfile:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMinimumRiskProfile.sql#L24) | WHERE / P06 / LIKE | `SecurityNo` : `VARCHAR(20)` | `'%' + @cSecurityNo + '%'` : `NVARCHAR (combined result; facets unknown)` | `SecurityNo` → `NVARCHAR` | `SecurityNo like '%' + @cSecurityNo + '%'` |
| IC-008266 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:121](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L121) | EXPRESSION / P07 / CASE_BRANCH | `sttm.SecurityNo` : `VARCHAR(20)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `sttm.SecurityNo` → `NVARCHAR` | `CASE WHEN sttm.IsCorporateBond = 1 AND sttm.SecCcy = 'IDR' THEN CASE WHEN (ISNULL(RTRIM(mxsec.M_CTP_BNDLBL),'') = '' or RTRIM(sttm.SecurityNo) = RTRIM(mxsec.M_CTP_BNDLBL)) THEN sttm.SecurityNo ELSE sttm.SecurityNo + ' / ' + RTRIM(mxsec.M_CTP_BNDLBL) END ELSE sttm.SecurityNo END` |
| IC-008268 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:124](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L124) | EXPRESSION / P07 / CASE_BRANCH | `sttm.SecurityNo` : `VARCHAR(20)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `sttm.SecurityNo` → `NVARCHAR` | `CASE WHEN (ISNULL(RTRIM(mxsec.M_CTP_BNDLBL),'') = '' or RTRIM(sttm.SecurityNo) = RTRIM(mxsec.M_CTP_BNDLBL)) THEN sttm.SecurityNo ELSE sttm.SecurityNo + ' / ' + RTRIM(mxsec.M_CTP_BNDLBL) END` |
| IC-008270 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSReportReportSettlementInstruction:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L125) | CASE / P04 / COMPARISON | `RTRIM(sttm.SecurityNo)` : `VARCHAR (expression; length unknown)` | `RTRIM(mxsec.M_CTP_BNDLBL)` : `NVARCHAR (expression; length unknown)` | `RTRIM(sttm.SecurityNo)` → `NVARCHAR` | `RTRIM(sttm.SecurityNo) = RTRIM(mxsec.M_CTP_BNDLBL)` |
| IC-008272 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSReportReportSettlementInstruction:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L127) | CASE / P05 / BINARY_Add | `sttm.SecurityNo + ' / '` : `VARCHAR (combined result; facets unknown)` | `RTRIM(mxsec.M_CTP_BNDLBL)` : `NVARCHAR (expression; length unknown)` | `sttm.SecurityNo + ' / '` → `NVARCHAR` | `sttm.SecurityNo + ' / ' + RTRIM(mxsec.M_CTP_BNDLBL)` |
| IC-008286 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:222](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L222) | EXPRESSION / P07 / CASE_BRANCH | `sttm.SecurityNo` : `VARCHAR(20)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `sttm.SecurityNo` → `NVARCHAR` | `CASE WHEN sttm.IsCorporateBond = 1 AND sttm.SecCcy = 'IDR' THEN CASE WHEN (ISNULL(RTRIM(mxsec.M_CTP_BNDLBL),'') = '' or RTRIM(sttm.SecurityNo) = RTRIM(mxsec.M_CTP_BNDLBL)) THEN sttm.SecurityNo ELSE sttm.SecurityNo + ' / ' + RTRIM(mxsec.M_CTP_BNDLBL) END ELSE sttm.SecurityNo END` |
| IC-008288 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L225) | EXPRESSION / P07 / CASE_BRANCH | `sttm.SecurityNo` : `VARCHAR(20)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `sttm.SecurityNo` → `NVARCHAR` | `CASE WHEN (ISNULL(RTRIM(mxsec.M_CTP_BNDLBL),'') = '' or RTRIM(sttm.SecurityNo) = RTRIM(mxsec.M_CTP_BNDLBL)) THEN sttm.SecurityNo ELSE sttm.SecurityNo + ' / ' + RTRIM(mxsec.M_CTP_BNDLBL) END` |
| IC-008290 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSReportReportSettlementInstruction:226](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L226) | CASE / P04 / COMPARISON | `RTRIM(sttm.SecurityNo)` : `VARCHAR (expression; length unknown)` | `RTRIM(mxsec.M_CTP_BNDLBL)` : `NVARCHAR (expression; length unknown)` | `RTRIM(sttm.SecurityNo)` → `NVARCHAR` | `RTRIM(sttm.SecurityNo) = RTRIM(mxsec.M_CTP_BNDLBL)` |
| IC-008292 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSReportReportSettlementInstruction:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L228) | CASE / P05 / BINARY_Add | `sttm.SecurityNo + ' / '` : `VARCHAR (combined result; facets unknown)` | `RTRIM(mxsec.M_CTP_BNDLBL)` : `NVARCHAR (expression; length unknown)` | `sttm.SecurityNo + ' / '` → `NVARCHAR` | `sttm.SecurityNo + ' / ' + RTRIM(mxsec.M_CTP_BNDLBL)` |
| IC-008306 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:324](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L324) | EXPRESSION / P07 / CASE_BRANCH | `sttm.SecurityNo` : `VARCHAR(20)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `sttm.SecurityNo` → `NVARCHAR` | `CASE WHEN sttm.IsCorporateBond = 1 AND sttm.SecCcy = 'IDR' THEN CASE WHEN (ISNULL(RTRIM(mxsec.M_CTP_BNDLBL),'') = '' or RTRIM(sttm.SecurityNo) = RTRIM(mxsec.M_CTP_BNDLBL)) THEN sttm.SecurityNo ELSE sttm.SecurityNo + ' / ' + RTRIM(mxsec.M_CTP_BNDLBL) END ELSE sttm.SecurityNo END` |
| IC-008308 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:327](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L327) | EXPRESSION / P07 / CASE_BRANCH | `sttm.SecurityNo` : `VARCHAR(20)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `sttm.SecurityNo` → `NVARCHAR` | `CASE WHEN (ISNULL(RTRIM(mxsec.M_CTP_BNDLBL),'') = '' or RTRIM(sttm.SecurityNo) = RTRIM(mxsec.M_CTP_BNDLBL)) THEN sttm.SecurityNo ELSE sttm.SecurityNo + ' / ' + RTRIM(mxsec.M_CTP_BNDLBL) END` |
| IC-008310 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSReportReportSettlementInstruction:328](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L328) | CASE / P04 / COMPARISON | `RTRIM(sttm.SecurityNo)` : `VARCHAR (expression; length unknown)` | `RTRIM(mxsec.M_CTP_BNDLBL)` : `NVARCHAR (expression; length unknown)` | `RTRIM(sttm.SecurityNo)` → `NVARCHAR` | `RTRIM(sttm.SecurityNo) = RTRIM(mxsec.M_CTP_BNDLBL)` |
| IC-008312 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSReportReportSettlementInstruction:330](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L330) | CASE / P05 / BINARY_Add | `sttm.SecurityNo + ' / '` : `VARCHAR (combined result; facets unknown)` | `RTRIM(mxsec.M_CTP_BNDLBL)` : `NVARCHAR (expression; length unknown)` | `sttm.SecurityNo + ' / '` → `NVARCHAR` | `sttm.SecurityNo + ' / ' + RTRIM(mxsec.M_CTP_BNDLBL)` |
| IC-008326 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:425](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L425) | EXPRESSION / P07 / CASE_BRANCH | `sttm.SecurityNo` : `VARCHAR(20)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `sttm.SecurityNo` → `NVARCHAR` | `CASE WHEN sttm.IsCorporateBond = 1 AND sttm.SecCcy = 'IDR' THEN CASE WHEN (ISNULL(RTRIM(mxsec.M_CTP_BNDLBL),'') = '' or RTRIM(sttm.SecurityNo) = RTRIM(mxsec.M_CTP_BNDLBL)) THEN sttm.SecurityNo ELSE sttm.SecurityNo + ' / ' + RTRIM(mxsec.M_CTP_BNDLBL) END ELSE sttm.SecurityNo END` |
| IC-008328 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:428](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L428) | EXPRESSION / P07 / CASE_BRANCH | `sttm.SecurityNo` : `VARCHAR(20)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `sttm.SecurityNo` → `NVARCHAR` | `CASE WHEN (ISNULL(RTRIM(mxsec.M_CTP_BNDLBL),'') = '' or RTRIM(sttm.SecurityNo) = RTRIM(mxsec.M_CTP_BNDLBL)) THEN sttm.SecurityNo ELSE sttm.SecurityNo + ' / ' + RTRIM(mxsec.M_CTP_BNDLBL) END` |
| IC-008330 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSReportReportSettlementInstruction:429](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L429) | CASE / P04 / COMPARISON | `RTRIM(sttm.SecurityNo)` : `VARCHAR (expression; length unknown)` | `RTRIM(mxsec.M_CTP_BNDLBL)` : `NVARCHAR (expression; length unknown)` | `RTRIM(sttm.SecurityNo)` → `NVARCHAR` | `RTRIM(sttm.SecurityNo) = RTRIM(mxsec.M_CTP_BNDLBL)` |
| IC-008332 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSReportReportSettlementInstruction:431](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L431) | CASE / P05 / BINARY_Add | `sttm.SecurityNo + ' / '` : `VARCHAR (combined result; facets unknown)` | `RTRIM(mxsec.M_CTP_BNDLBL)` : `NVARCHAR (expression; length unknown)` | `sttm.SecurityNo + ' / '` → `NVARCHAR` | `sttm.SecurityNo + ' / ' + RTRIM(mxsec.M_CTP_BNDLBL)` |
| IC-008363 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:575](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L575) | EXPRESSION / P07 / CASE_BRANCH | `sttm.SecurityNo` : `VARCHAR(20)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `sttm.SecurityNo` → `NVARCHAR` | `CASE WHEN sttm.IsCorporateBond = 1 AND sttm.SecCcy = 'IDR' THEN CASE WHEN (ISNULL(RTRIM(mxsec.M_CTP_BNDLBL),'') = '' or RTRIM(sttm.SecurityNo) = RTRIM(mxsec.M_CTP_BNDLBL)) THEN sttm.SecurityNo ELSE sttm.SecurityNo + ' / ' + RTRIM(mxsec.M_CTP_BNDLBL) END ELSE sttm.SecurityNo END` |
| IC-008365 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:578](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L578) | EXPRESSION / P07 / CASE_BRANCH | `sttm.SecurityNo` : `VARCHAR(20)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `sttm.SecurityNo` → `NVARCHAR` | `CASE WHEN (ISNULL(RTRIM(mxsec.M_CTP_BNDLBL),'') = '' or RTRIM(sttm.SecurityNo) = RTRIM(mxsec.M_CTP_BNDLBL)) THEN sttm.SecurityNo ELSE sttm.SecurityNo + ' / ' + RTRIM(mxsec.M_CTP_BNDLBL) END` |
| IC-008367 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSReportReportSettlementInstruction:579](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L579) | CASE / P04 / COMPARISON | `RTRIM(sttm.SecurityNo)` : `VARCHAR (expression; length unknown)` | `RTRIM(mxsec.M_CTP_BNDLBL)` : `NVARCHAR (expression; length unknown)` | `RTRIM(sttm.SecurityNo)` → `NVARCHAR` | `RTRIM(sttm.SecurityNo) = RTRIM(mxsec.M_CTP_BNDLBL)` |
| IC-008369 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSReportReportSettlementInstruction:581](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L581) | CASE / P05 / BINARY_Add | `sttm.SecurityNo + ' / '` : `VARCHAR (combined result; facets unknown)` | `RTRIM(mxsec.M_CTP_BNDLBL)` : `NVARCHAR (expression; length unknown)` | `sttm.SecurityNo + ' / '` → `NVARCHAR` | `sttm.SecurityNo + ' / ' + RTRIM(mxsec.M_CTP_BNDLBL)` |

<a id="kolom-couponeom"></a>
## CouponEOM

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L102). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012368 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSSubmitTransferAsset:178](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitTransferAsset.sql#L178) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CouponEOM` : `BIT` | `@curCouponEOM` : `int` | `CouponEOM` → `int` | `@curCouponEOM = CouponEOM` |
| IC-004131 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSGetMenuSettings:279](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetMenuSettings.sql#L279) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `CouponEOM` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(CouponEOM, 0)` |

<a id="kolom-couponperunit"></a>
## CouponPerUnit

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L30). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002703 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L96) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `CouponPerUnit` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(CouponPerUnit, 0)` |
| IC-010801 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORIDetail:143](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L143) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `CouponPerUnit` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(CouponPerUnit, 0)` |
| IC-010933 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnesty:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L137) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `CouponPerUnit` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(CouponPerUnit, 0)` |

<a id="kolom-couponperiode"></a>
## CouponPeriode

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L101). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004130 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSGetMenuSettings:278](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetMenuSettings.sql#L278) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `CouponPeriode` : `INT` | `''` : `varchar` | `''` → `INT` | `isnull(CouponPeriode, '')` |

<a id="kolom-couponrate"></a>
## CouponRate

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 16 lokasi / 12 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT, COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001277 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSProsesDataJobOutstanding:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L85) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tm.CouponRate` : `DECIMAL(20,5)` | `tmp.CouponRate` : `money` | `tm.CouponRate` → `money` | `tmp.CouponRate = tm.CouponRate` |
| IC-002908 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPushMailNotifOrderCompleted:155](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L155) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[CouponRate]` : `DECIMAL(20,5)` | `@cCOUPONRATE` : `varchar(30)` | `[CouponRate]` → `varchar(30)` | `@cCOUPONRATE = [CouponRate]` |
| IC-002931 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPushMailNotifOrderORI:138](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderORI.sql#L138) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[CouponRate]` : `DECIMAL(20,5)` | `@cCOUPONRATE` : `varchar(30)` | `[CouponRate]` → `varchar(30)` | `@cCOUPONRATE = [CouponRate]` |
| IC-010059 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSRptKonfirmasiBOND:369](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L369) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `isnull(sm.CouponRate,0)` : `DECIMAL(20,5)` | `retail.dbo.fnoblseparatorsign.@pcNominal` : `varchar(100)` | `isnull(sm.CouponRate,0)` → `varchar(100)` | `dbo.fnOBLSeparatorSign(isnull(sm.CouponRate,0))` |
| IC-010114 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSRptKonfirmasiBONDNew:249](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBONDNew.sql#L249) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `isnull(sm.CouponRate,0)` : `DECIMAL(20,5)` | `retail.dbo.fnoblseparatorsign.@pcNominal` : `varchar(100)` | `isnull(sm.CouponRate,0)` → `varchar(100)` | `dbo.fnOBLSeparatorSign(isnull(sm.CouponRate,0))` |
| IC-010176 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSRptKonfirmasiSRBI:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiSRBI.sql#L180) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `isnull(sm.CouponRate,0)` : `DECIMAL(20,5)` | `retail.dbo.fnoblseparatorsign.@pcNominal` : `varchar(100)` | `isnull(sm.CouponRate,0)` → `varchar(100)` | `dbo.fnOBLSeparatorSign(isnull(sm.CouponRate,0))` |
| IC-002430 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulateParamHargaModalNew:170](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateParamHargaModalNew.sql#L170) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.CouponRate` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(sm.CouponRate, 0)` |
| IC-005908 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateParamHargaModalNew:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamHargaModalNew.sql#L228) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.CouponRate` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(sm.CouponRate, 0)` |
| IC-006215 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBuyBack:164](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L164) | DML_CONDITION / P05 / BINARY_Multiply | `( smtm.CouponRate * 0.01 / 12.00)` : `DECIMAL (combined result; facets unknown)` | `dbo.fnOBLGetAccruedDays ( smtm.TanggalSettlementBuyBack,smtm.NextCouponDate,smtm.LastCouponDate,smtm.IsCorporateBond,smtm.SecCcy,isnull(acd.Adjustment,0) )` : `int` | `dbo.fnOBLGetAccruedDays ( smtm.TanggalSettlementBuyBack,smtm.NextCouponDate,smtm.LastCouponDate,smtm.IsCorporateBond,smtm.SecCcy,isnull(acd.Adjustment,0) )` → `DECIMAL` | `( smtm.CouponRate * 0.01 / 12.00) * dbo.fnOBLGetAccruedDays ( smtm.TanggalSettlementBuyBack,smtm.NextCouponDate,smtm.LastCouponDate,smtm.IsCorporateBond,smtm.SecCcy,isnull(acd.Adjustment,0) )` |
| IC-006510 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateTrxBuyBackPerdana:138](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L138) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.CouponRate` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(sm.CouponRate, 0)` |
| IC-007327 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSecurityMasterAuthorization:397](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L397) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.CouponRate` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(smtm.CouponRate, 0)` |
| IC-007331 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSecurityMasterAuthorization:426](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L426) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.CouponRate` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(smtm.CouponRate, 0)` |
| IC-010058 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptKonfirmasiBOND:369](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L369) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.CouponRate` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(sm.CouponRate,0)` |
| IC-010113 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptKonfirmasiBONDNew:249](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBONDNew.sql#L249) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.CouponRate` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(sm.CouponRate,0)` |
| IC-010175 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptKonfirmasiSRBI:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiSRBI.sql#L180) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.CouponRate` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(sm.CouponRate,0)` |
| IC-012766 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateUploadParamHargaModalNew:246](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L246) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.CouponRate` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(sm.CouponRate, 0)` |

<a id="kolom-ecnominalkelipatan"></a>
## ECNominalKelipatan

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L97). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000652 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMGetListProductSekunder:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L156) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ECNominalKelipatan` : `MONEY` | `@tmp_Result.KelipatanNominal` : `decimal(21,2)` | `ECNominalKelipatan` → `decimal(21,2)` | `@tmp_Result.KelipatanNominal <- ECNominalKelipatan` |

<a id="kolom-ecnominalmax"></a>
## ECNominalMax

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L98). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000653 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMGetListProductSekunder:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L156) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ECNominalMax` : `MONEY` | `@tmp_Result.NominalMaksimum` : `decimal(21,2)` | `ECNominalMax` → `decimal(21,2)` | `@tmp_Result.NominalMaksimum <- ECNominalMax` |

<a id="kolom-ecnominalmin"></a>
## ECNominalMin

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L96). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000654 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMGetListProductSekunder:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L156) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ECNominalMin` : `MONEY` | `@tmp_Result.NominalMinimum` : `decimal(21,2)` | `ECNominalMin` → `decimal(21,2)` | `@tmp_Result.NominalMinimum <- ECNominalMin` |

<a id="kolom-feetransaksibankbeli"></a>
## FeeTransaksiBankBeli

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L34). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003197 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSAmendSecurityTransaction_TT:413](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L413) | CASE / P05 / BINARY_Multiply | `smtm.FeeTransaksiBankBeli` : `MONEY` | `(100 + smtm.FeeTransaksiBankBeliPPN)` : `DECIMAL (combined result; facets unknown)` | `smtm.FeeTransaksiBankBeli` → `DECIMAL` | `smtm.FeeTransaksiBankBeli * (100 + smtm.FeeTransaksiBankBeliPPN)` |
| IC-004304 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSInsertSecurityTransaction_TT_Original:736](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L736) | CASE / P05 / BINARY_Multiply | `smtm.FeeTransaksiBankBeli` : `MONEY` | `(100 + smtm.FeeTransaksiBankBeliPPN)` : `DECIMAL (combined result; facets unknown)` | `smtm.FeeTransaksiBankBeli` → `DECIMAL` | `smtm.FeeTransaksiBankBeli * (100 + smtm.FeeTransaksiBankBeliPPN)` |

<a id="kolom-feetransaksibankbelippn"></a>
## FeeTransaksiBankBeliPPN

**Tipe dasar:** `DECIMAL(6,3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L35). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003196 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSAmendSecurityTransaction_TT:413](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L413) | CASE / P05 / BINARY_Add | `smtm.FeeTransaksiBankBeliPPN` : `DECIMAL(6,3)` | `100` : `int` | `100` → `DECIMAL` | `100 + smtm.FeeTransaksiBankBeliPPN` |
| IC-003197 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSAmendSecurityTransaction_TT:413](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L413) | CASE / P05 / BINARY_Multiply | `(100 + smtm.FeeTransaksiBankBeliPPN)` : `DECIMAL (combined result; facets unknown)` | `smtm.FeeTransaksiBankBeli` : `MONEY` | `smtm.FeeTransaksiBankBeli` → `DECIMAL` | `smtm.FeeTransaksiBankBeli * (100 + smtm.FeeTransaksiBankBeliPPN)` |
| IC-004303 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSInsertSecurityTransaction_TT_Original:736](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L736) | CASE / P05 / BINARY_Add | `smtm.FeeTransaksiBankBeliPPN` : `DECIMAL(6,3)` | `100` : `int` | `100` → `DECIMAL` | `100 + smtm.FeeTransaksiBankBeliPPN` |
| IC-004304 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSInsertSecurityTransaction_TT_Original:736](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L736) | CASE / P05 / BINARY_Multiply | `(100 + smtm.FeeTransaksiBankBeliPPN)` : `DECIMAL (combined result; facets unknown)` | `smtm.FeeTransaksiBankBeli` : `MONEY` | `smtm.FeeTransaksiBankBeli` → `DECIMAL` | `smtm.FeeTransaksiBankBeli * (100 + smtm.FeeTransaksiBankBeliPPN)` |

<a id="kolom-feetransaksibankjual"></a>
## FeeTransaksiBankJual

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L36). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003200 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSAmendSecurityTransaction_TT:414](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L414) | CASE / P05 / BINARY_Multiply | `smtm.FeeTransaksiBankJual` : `MONEY` | `(100 + smtm.FeeTransaksiBankJualPPN)` : `DECIMAL (combined result; facets unknown)` | `smtm.FeeTransaksiBankJual` → `DECIMAL` | `smtm.FeeTransaksiBankJual * (100 + smtm.FeeTransaksiBankJualPPN)` |
| IC-004307 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSInsertSecurityTransaction_TT_Original:737](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L737) | CASE / P05 / BINARY_Multiply | `smtm.FeeTransaksiBankJual` : `MONEY` | `(100 + smtm.FeeTransaksiBankJualPPN)` : `DECIMAL (combined result; facets unknown)` | `smtm.FeeTransaksiBankJual` → `DECIMAL` | `smtm.FeeTransaksiBankJual * (100 + smtm.FeeTransaksiBankJualPPN)` |
| IC-007763 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessUploadTransaksiSuratBerharga:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessUploadTransaksiSuratBerharga.sql#L141) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `mst.FeeTransaksiBankJual` : `MONEY` | `(100 + mst.FeeTransaksiBankJualPPN)` : `DECIMAL (combined result; facets unknown)` | `mst.FeeTransaksiBankJual` → `DECIMAL` | `mst.FeeTransaksiBankJual * (100 + mst.FeeTransaksiBankJualPPN)` |

<a id="kolom-feetransaksibankjualppn"></a>
## FeeTransaksiBankJualPPN

**Tipe dasar:** `DECIMAL(6,3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L37). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 10 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003199 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSAmendSecurityTransaction_TT:414](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L414) | CASE / P05 / BINARY_Add | `smtm.FeeTransaksiBankJualPPN` : `DECIMAL(6,3)` | `100` : `int` | `100` → `DECIMAL` | `100 + smtm.FeeTransaksiBankJualPPN` |
| IC-003200 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSAmendSecurityTransaction_TT:414](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L414) | CASE / P05 / BINARY_Multiply | `(100 + smtm.FeeTransaksiBankJualPPN)` : `DECIMAL (combined result; facets unknown)` | `smtm.FeeTransaksiBankJual` : `MONEY` | `smtm.FeeTransaksiBankJual` → `DECIMAL` | `smtm.FeeTransaksiBankJual * (100 + smtm.FeeTransaksiBankJualPPN)` |
| IC-004306 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSInsertSecurityTransaction_TT_Original:737](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L737) | CASE / P05 / BINARY_Add | `smtm.FeeTransaksiBankJualPPN` : `DECIMAL(6,3)` | `100` : `int` | `100` → `DECIMAL` | `100 + smtm.FeeTransaksiBankJualPPN` |
| IC-004307 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSInsertSecurityTransaction_TT_Original:737](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L737) | CASE / P05 / BINARY_Multiply | `(100 + smtm.FeeTransaksiBankJualPPN)` : `DECIMAL (combined result; facets unknown)` | `smtm.FeeTransaksiBankJual` : `MONEY` | `smtm.FeeTransaksiBankJual` → `DECIMAL` | `smtm.FeeTransaksiBankJual * (100 + smtm.FeeTransaksiBankJualPPN)` |
| IC-007762 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessUploadTransaksiSuratBerharga:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessUploadTransaksiSuratBerharga.sql#L141) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `mst.FeeTransaksiBankJualPPN` : `DECIMAL(6,3)` | `100` : `int` | `100` → `DECIMAL` | `100 + mst.FeeTransaksiBankJualPPN` |
| IC-007763 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessUploadTransaksiSuratBerharga:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessUploadTransaksiSuratBerharga.sql#L141) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `(100 + mst.FeeTransaksiBankJualPPN)` : `DECIMAL (combined result; facets unknown)` | `mst.FeeTransaksiBankJual` : `MONEY` | `mst.FeeTransaksiBankJual` → `DECIMAL` | `mst.FeeTransaksiBankJual * (100 + mst.FeeTransaksiBankJualPPN)` |
| IC-011021 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnesty:641](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L641) | CASE / P05 / BINARY_Multiply | `sm.FeeTransaksiBankJualPPN` : `DECIMAL(6,3)` | `tmp.Actual` : `money` | `tmp.Actual` → `DECIMAL` | `tmp.Actual * sm.FeeTransaksiBankJualPPN` |
| IC-011022 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnesty:642](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L642) | DML_CONDITION / P05 / BINARY_Multiply | `sm.FeeTransaksiBankJualPPN` : `DECIMAL(6,3)` | `tmp.Actual` : `money` | `tmp.Actual` → `DECIMAL` | `tmp.Actual * sm.FeeTransaksiBankJualPPN` |
| IC-011232 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:521](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L521) | DML_CONDITION / P05 / BINARY_Multiply | `sm.FeeTransaksiBankJualPPN` : `DECIMAL(6,3)` | `tmp.Actual` : `MONEY` | `tmp.Actual` → `DECIMAL` | `tmp.Actual * sm.FeeTransaksiBankJualPPN` |
| IC-011432 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnesty_Backdated:460](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty_Backdated.sql#L460) | DML_CONDITION / P05 / BINARY_Multiply | `sm.FeeTransaksiBankJualPPN` : `DECIMAL(6,3)` | `tmp.Actual` : `money` | `tmp.Actual` → `DECIMAL` | `tmp.Actual * sm.FeeTransaksiBankJualPPN` |

<a id="kolom-feetype"></a>
## FeeType

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L38). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008967 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailCalculateTransactionFee:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateTransactionFee.sql#L16) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `FeeType` : `INT` | `@nSecurityFeeType` : `smallint` | `FeeType` → `smallint` | `@nSecurityFeeType = FeeType` |

<a id="kolom-flagcashback"></a>
## FlagCashBack

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L85). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002709 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:140](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L140) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagCashBack` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(FlagCashBack, 0)` |
| IC-013339 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:93](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L93) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagCashBack` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(FlagCashBack,0)` |

<a id="kolom-flagechannel"></a>
## FlagEChannel

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L93). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 24 lokasi / 7 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000660 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMGetListProductSekunder:235](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L235) | WHERE / P11 / COMPARISON | `FlagEChannel` : `BIT` | `1` : `int` | `FlagEChannel` → `INT` | `FlagEChannel = 1` |
| IC-000972 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.ORISBNIPopulateEarlyRedemption:116](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateEarlyRedemption.sql#L116) | WHERE / P11 / COMPARISON | `FlagEChannel` : `BIT` | `1` : `int` | `FlagEChannel` → `INT` | `FlagEChannel = 1` |
| IC-002535 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSRetailSrcSecId:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L80) | WHERE / P11 / COMPARISON | `FlagEChannel` : `BIT` | `0` : `int` | `FlagEChannel` → `INT` | `FlagEChannel = 0` |
| IC-002540 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSRetailSrcSecId:112](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L112) | WHERE / P11 / COMPARISON | `FlagEChannel` : `BIT` | `0` : `int` | `FlagEChannel` → `INT` | `FlagEChannel = 0` |
| IC-002545 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSRetailSrcSecId:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L141) | WHERE / P11 / COMPARISON | `FlagEChannel` : `BIT` | `0` : `int` | `FlagEChannel` → `INT` | `FlagEChannel = 0` |
| IC-002550 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSRetailSrcSecId:171](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L171) | WHERE / P11 / COMPARISON | `FlagEChannel` : `BIT` | `0` : `int` | `FlagEChannel` → `INT` | `FlagEChannel = 0` |
| IC-002555 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSRetailSrcSecId:199](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L199) | WHERE / P11 / COMPARISON | `FlagEChannel` : `BIT` | `0` : `int` | `FlagEChannel` → `INT` | `FlagEChannel = 0` |
| IC-012128 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchECSecurityNoHargaORI:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchECSecurityNoHargaORI.sql#L46) | WHERE / P11 / COMPARISON | `a.FlagEChannel` : `BIT` | `1` : `int` | `a.FlagEChannel` → `INT` | `a.FlagEChannel = 1` |
| IC-012130 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchECSecurityNoHargaORI:60](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchECSecurityNoHargaORI.sql#L60) | WHERE / P11 / COMPARISON | `a.FlagEChannel` : `BIT` | `1` : `int` | `a.FlagEChannel` → `INT` | `a.FlagEChannel = 1` |
| IC-012133 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchECSecurityNoHargaORI:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchECSecurityNoHargaORI.sql#L86) | WHERE / P11 / COMPARISON | `a.FlagEChannel` : `BIT` | `1` : `int` | `a.FlagEChannel` → `INT` | `a.FlagEChannel = 1` |
| IC-012135 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchECSecurityNoHargaORI:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchECSecurityNoHargaORI.sql#L98) | WHERE / P11 / COMPARISON | `a.FlagEChannel` : `BIT` | `1` : `int` | `a.FlagEChannel` → `INT` | `a.FlagEChannel = 1` |
| IC-002531 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSRetailSrcSecId:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L70) | WHERE / P04 / COMPARISON | `ISNULL(FlagEChannel, 0)` : `BIT` | `@bEChannel` : `varchar(10)` | `@bEChannel` → `BIT` | `ISNULL(FlagEChannel, 0) = @bEChannel` |
| IC-002532 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSRetailSrcSecId:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L70) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagEChannel` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(FlagEChannel, 0)` |
| IC-002536 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSRetailSrcSecId:102](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L102) | WHERE / P04 / COMPARISON | `ISNULL(FlagEChannel, 0)` : `BIT` | `@bEChannel` : `varchar(10)` | `@bEChannel` → `BIT` | `ISNULL(FlagEChannel, 0) = @bEChannel` |
| IC-002537 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSRetailSrcSecId:102](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L102) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagEChannel` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(FlagEChannel, 0)` |
| IC-002541 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSRetailSrcSecId:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L131) | WHERE / P04 / COMPARISON | `ISNULL(FlagEChannel, 0)` : `BIT` | `@bEChannel` : `varchar(10)` | `@bEChannel` → `BIT` | `ISNULL(FlagEChannel, 0) = @bEChannel` |
| IC-002542 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSRetailSrcSecId:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L131) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagEChannel` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(FlagEChannel, 0)` |
| IC-002546 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSRetailSrcSecId:161](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L161) | WHERE / P04 / COMPARISON | `ISNULL(FlagEChannel, 0)` : `BIT` | `@bEChannel` : `varchar(10)` | `@bEChannel` → `BIT` | `ISNULL(FlagEChannel, 0) = @bEChannel` |
| IC-002547 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSRetailSrcSecId:161](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L161) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagEChannel` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(FlagEChannel, 0)` |
| IC-002551 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSRetailSrcSecId:189](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L189) | WHERE / P04 / COMPARISON | `ISNULL(FlagEChannel, 0)` : `BIT` | `@bEChannel` : `varchar(10)` | `@bEChannel` → `BIT` | `ISNULL(FlagEChannel, 0) = @bEChannel` |
| IC-002552 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSRetailSrcSecId:189](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L189) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagEChannel` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(FlagEChannel, 0)` |
| IC-002712 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:157](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L157) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagEChannel` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(FlagEChannel, 0)` |
| IC-004132 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSGetMenuSettings:281](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetMenuSettings.sql#L281) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagEChannel` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(FlagEChannel, 0)` |
| IC-013342 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L114) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagEChannel` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(FlagEChannel, 0)` |

<a id="kolom-flagperdana"></a>
## FlagPerdana

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L73). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 11 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN_EXPRESSION, COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001030 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.OrivaGetParameter:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OrivaGetParameter.sql#L22) | WHERE / P11 / COMPARISON | `ISNULL(FlagPerdana, CAST(0 AS BIT))` : `BIT` | `1` : `int` | `ISNULL(FlagPerdana, CAST(0 AS BIT))` → `INT` | `ISNULL(FlagPerdana, CAST(0 AS BIT)) = 1` |
| IC-002533 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSRetailSrcSecId:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L77) | WHERE / P04 / COMPARISON | `ISNULL(FlagPerdana, 0)` : `BIT` | `@bPerdana` : `varchar(10)` | `@bPerdana` → `BIT` | `ISNULL(FlagPerdana, 0) = @bPerdana` |
| IC-002534 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSRetailSrcSecId:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L77) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagPerdana` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(FlagPerdana, 0)` |
| IC-002538 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSRetailSrcSecId:109](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L109) | WHERE / P04 / COMPARISON | `ISNULL(FlagPerdana, 0)` : `BIT` | `@bPerdana` : `varchar(10)` | `@bPerdana` → `BIT` | `ISNULL(FlagPerdana, 0) = @bPerdana` |
| IC-002539 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSRetailSrcSecId:109](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L109) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagPerdana` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(FlagPerdana, 0)` |
| IC-002543 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSRetailSrcSecId:138](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L138) | WHERE / P04 / COMPARISON | `ISNULL(FlagPerdana, 0)` : `BIT` | `@bPerdana` : `varchar(10)` | `@bPerdana` → `BIT` | `ISNULL(FlagPerdana, 0) = @bPerdana` |
| IC-002544 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSRetailSrcSecId:138](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L138) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagPerdana` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(FlagPerdana, 0)` |
| IC-002548 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSRetailSrcSecId:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L168) | WHERE / P04 / COMPARISON | `ISNULL(FlagPerdana, 0)` : `BIT` | `@bPerdana` : `varchar(10)` | `@bPerdana` → `BIT` | `ISNULL(FlagPerdana, 0) = @bPerdana` |
| IC-002549 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSRetailSrcSecId:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L168) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagPerdana` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(FlagPerdana, 0)` |
| IC-002553 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSRetailSrcSecId:196](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L196) | WHERE / P04 / COMPARISON | `ISNULL(FlagPerdana, 0)` : `BIT` | `@bPerdana` : `varchar(10)` | `@bPerdana` → `BIT` | `ISNULL(FlagPerdana, 0) = @bPerdana` |
| IC-002554 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSRetailSrcSecId:196](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSRetailSrcSecId.sql#L196) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagPerdana` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(FlagPerdana, 0)` |

<a id="kolom-flagsyariah"></a>
## FlagSyariah

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L92). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 107 lokasi / 25 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000150 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.CustodyReportMaintainAcquisition:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportMaintainAcquisition.sql#L110) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `FlagSyariah` : `BIT` | `#tmp_SecurityMaster_TM.FlagSyariah` : `VARCHAR(5)` | `FlagSyariah` → `VARCHAR(5)` | `#tmp_SecurityMaster_TM.FlagSyariah <- FlagSyariah` |
| IC-002910 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPushMailNotifOrderCompleted:157](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L157) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `isnull(FlagSyariah,0)` : `BIT` | `@nFlagSyariah` : `int` | `isnull(FlagSyariah,0)` → `int` | `@nFlagSyariah = isnull(FlagSyariah,0)` |
| IC-002933 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPushMailNotifOrderORI:140](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderORI.sql#L140) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `isnull(FlagSyariah,0)` : `BIT` | `@nFlagSyariah` : `int` | `isnull(FlagSyariah,0)` → `int` | `@nFlagSyariah = isnull(FlagSyariah,0)` |
| IC-010023 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptKonfirmasiBOND:136](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L136) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sm.FlagSyariah` : `BIT` | `#tmp_data.FlagSyariah` : `int` | `sm.FlagSyariah` → `int` | `#tmp_data.FlagSyariah <- sm.FlagSyariah` |
| IC-010053 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptKonfirmasiBOND:272](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L272) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sm.FlagSyariah` : `BIT` | `#tmp_data.FlagSyariah` : `int` | `sm.FlagSyariah` → `int` | `#tmp_data.FlagSyariah <- sm.FlagSyariah` |
| IC-010108 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptKonfirmasiBONDNew:164](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBONDNew.sql#L164) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sm.FlagSyariah` : `BIT` | `#tmp_data.FlagSyariah` : `int` | `sm.FlagSyariah` → `int` | `#tmp_data.FlagSyariah <- sm.FlagSyariah` |
| IC-000189 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L111) | CASE / P11 / COMPARISON | `sttm.FlagSyariah` : `BIT` | `0` : `int` | `sttm.FlagSyariah` → `INT` | `sttm.FlagSyariah = 0` |
| IC-000191 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:112](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L112) | CASE / P11 / COMPARISON | `sttm.FlagSyariah` : `BIT` | `1` : `int` | `sttm.FlagSyariah` → `INT` | `sttm.FlagSyariah = 1` |
| IC-000193 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L113) | CASE / P11 / COMPARISON | `sttm.FlagSyariah` : `BIT` | `0` : `int` | `sttm.FlagSyariah` → `INT` | `sttm.FlagSyariah = 0` |
| IC-000196 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L114) | CASE / P11 / COMPARISON | `sttm.FlagSyariah` : `BIT` | `0` : `int` | `sttm.FlagSyariah` → `INT` | `sttm.FlagSyariah = 0` |
| IC-000199 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L115) | CASE / P11 / COMPARISON | `sttm.FlagSyariah` : `BIT` | `1` : `int` | `sttm.FlagSyariah` → `INT` | `sttm.FlagSyariah = 1` |
| IC-000216 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:197](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L197) | CASE / P11 / COMPARISON | `sttm.FlagSyariah` : `BIT` | `0` : `int` | `sttm.FlagSyariah` → `INT` | `sttm.FlagSyariah = 0` |
| IC-000218 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:198](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L198) | CASE / P11 / COMPARISON | `sttm.FlagSyariah` : `BIT` | `1` : `int` | `sttm.FlagSyariah` → `INT` | `sttm.FlagSyariah = 1` |
| IC-000220 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:199](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L199) | CASE / P11 / COMPARISON | `sttm.FlagSyariah` : `BIT` | `0` : `int` | `sttm.FlagSyariah` → `INT` | `sttm.FlagSyariah = 0` |
| IC-000222 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:200](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L200) | CASE / P11 / COMPARISON | `sttm.FlagSyariah` : `BIT` | `1` : `int` | `sttm.FlagSyariah` → `INT` | `sttm.FlagSyariah = 1` |
| IC-000226 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:262](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L262) | CASE / P11 / COMPARISON | `sttm.FlagSyariah` : `BIT` | `0` : `int` | `sttm.FlagSyariah` → `INT` | `sttm.FlagSyariah = 0` |
| IC-000227 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:263](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L263) | CASE / P11 / COMPARISON | `sttm.FlagSyariah` : `BIT` | `1` : `int` | `sttm.FlagSyariah` → `INT` | `sttm.FlagSyariah = 1` |
| IC-000229 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:264](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L264) | CASE / P11 / COMPARISON | `sttm.FlagSyariah` : `BIT` | `0` : `int` | `sttm.FlagSyariah` → `INT` | `sttm.FlagSyariah = 0` |
| IC-000231 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L265) | CASE / P11 / COMPARISON | `sttm.FlagSyariah` : `BIT` | `1` : `int` | `sttm.FlagSyariah` → `INT` | `sttm.FlagSyariah = 1` |
| IC-000237 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:336](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L336) | CASE / P11 / COMPARISON | `sm.FlagSyariah` : `BIT` | `0` : `int` | `sm.FlagSyariah` → `INT` | `sm.FlagSyariah = 0` |
| IC-000239 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L337) | CASE / P11 / COMPARISON | `sm.FlagSyariah` : `BIT` | `1` : `int` | `sm.FlagSyariah` → `INT` | `sm.FlagSyariah = 1` |
| IC-000241 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:342](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L342) | CASE / P11 / COMPARISON | `sm.FlagSyariah` : `BIT` | `0` : `int` | `sm.FlagSyariah` → `INT` | `sm.FlagSyariah = 0` |
| IC-000243 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:343](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L343) | CASE / P11 / COMPARISON | `sm.FlagSyariah` : `BIT` | `1` : `int` | `sm.FlagSyariah` → `INT` | `sm.FlagSyariah = 1` |
| IC-000246 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstructionTransferAsset:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstructionTransferAsset.sql#L43) | CASE / P11 / COMPARISON | `sm.FlagSyariah` : `BIT` | `0` : `int` | `sm.FlagSyariah` → `INT` | `sm.FlagSyariah = 0` |
| IC-000248 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstructionTransferAsset:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstructionTransferAsset.sql#L44) | CASE / P11 / COMPARISON | `sm.FlagSyariah` : `BIT` | `1` : `int` | `sm.FlagSyariah` → `INT` | `sm.FlagSyariah = 1` |
| IC-000250 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstructionTransferAsset:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstructionTransferAsset.sql#L45) | CASE / P11 / COMPARISON | `sm.FlagSyariah` : `BIT` | `0` : `int` | `sm.FlagSyariah` → `INT` | `sm.FlagSyariah = 0` |
| IC-000252 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstructionTransferAsset:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstructionTransferAsset.sql#L46) | CASE / P11 / COMPARISON | `sm.FlagSyariah` : `BIT` | `1` : `int` | `sm.FlagSyariah` → `INT` | `sm.FlagSyariah = 1` |
| IC-006298 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateSettlementMaturity:91](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementMaturity.sql#L91) | CASE / P11 / COMPARISON | `ISNULL(smtm.FlagSyariah, 0)` : `BIT` | `1` : `int` | `ISNULL(smtm.FlagSyariah, 0)` → `INT` | `ISNULL(smtm.FlagSyariah, 0) = 1` |
| IC-006300 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateSettlementMaturity:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementMaturity.sql#L98) | CASE / P11 / COMPARISON | `ISNULL(smtm.FlagSyariah, 0)` : `BIT` | `1` : `int` | `ISNULL(smtm.FlagSyariah, 0)` → `INT` | `ISNULL(smtm.FlagSyariah, 0) = 1` |
| IC-007493 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSProcessSettlementKupon:194](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementKupon.sql#L194) | WHERE / P11 / COMPARISON | `ISNULL(sc.FlagSyariah,0)` : `BIT` | `0` : `int` | `ISNULL(sc.FlagSyariah,0)` → `INT` | `ISNULL(sc.FlagSyariah,0) = 0` |
| IC-007496 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSProcessSettlementKupon:203](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementKupon.sql#L203) | WHERE / P11 / COMPARISON | `ISNULL(sc.FlagSyariah,0)` : `BIT` | `0` : `int` | `ISNULL(sc.FlagSyariah,0)` → `INT` | `ISNULL(sc.FlagSyariah,0) = 0` |
| IC-007502 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSProcessSettlementMaturity:116](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementMaturity.sql#L116) | CASE / P11 / COMPARISON | `isnull(smtm.FlagSyariah,0)` : `BIT` | `1` : `int` | `isnull(smtm.FlagSyariah,0)` → `INT` | `isnull(smtm.FlagSyariah,0) = 1` |
| IC-007504 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSProcessSettlementMaturity:122](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementMaturity.sql#L122) | CASE / P11 / COMPARISON | `isnull(smtm.FlagSyariah,0)` : `BIT` | `1` : `int` | `isnull(smtm.FlagSyariah,0)` → `INT` | `isnull(smtm.FlagSyariah,0) = 1` |
| IC-010459 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptOutstandingPerSecAccNo:188](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L188) | CASE / P11 / COMPARISON | `isnull(sm.FlagSyariah,0)` : `BIT` | `1` : `int` | `isnull(sm.FlagSyariah,0)` → `INT` | `isnull(sm.FlagSyariah,0) = 1` |
| IC-010461 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptOutstandingPerSecAccNo:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L193) | CASE / P11 / COMPARISON | `isnull(sm.FlagSyariah,0)` : `BIT` | `1` : `int` | `isnull(sm.FlagSyariah,0)` → `INT` | `isnull(sm.FlagSyariah,0) = 1` |
| IC-010480 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptOutstandingPerSecAccNo:294](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L294) | CASE / P11 / COMPARISON | `isnull(sm.FlagSyariah,0)` : `BIT` | `1` : `int` | `isnull(sm.FlagSyariah,0)` → `INT` | `isnull(sm.FlagSyariah,0) = 1` |
| IC-010482 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptOutstandingPerSecAccNo:299](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L299) | CASE / P11 / COMPARISON | `isnull(sm.FlagSyariah,0)` : `BIT` | `1` : `int` | `isnull(sm.FlagSyariah,0)` → `INT` | `isnull(sm.FlagSyariah,0) = 1` |
| IC-010501 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptOutstandingPerSecAccNo:432](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L432) | CASE / P11 / COMPARISON | `isnull(sm.FlagSyariah,0)` : `BIT` | `1` : `int` | `isnull(sm.FlagSyariah,0)` → `INT` | `isnull(sm.FlagSyariah,0) = 1` |
| IC-010503 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptOutstandingPerSecAccNo:437](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L437) | CASE / P11 / COMPARISON | `isnull(sm.FlagSyariah,0)` : `BIT` | `1` : `int` | `isnull(sm.FlagSyariah,0)` → `INT` | `isnull(sm.FlagSyariah,0) = 1` |
| IC-010522 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptOutstandingPerSecAccNo:534](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L534) | CASE / P11 / COMPARISON | `isnull(sm.FlagSyariah,0)` : `BIT` | `1` : `int` | `isnull(sm.FlagSyariah,0)` → `INT` | `isnull(sm.FlagSyariah,0) = 1` |
| IC-010524 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptOutstandingPerSecAccNo:539](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L539) | CASE / P11 / COMPARISON | `isnull(sm.FlagSyariah,0)` : `BIT` | `1` : `int` | `isnull(sm.FlagSyariah,0)` → `INT` | `isnull(sm.FlagSyariah,0) = 1` |
| IC-010892 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptPayCouponORIDetail:730](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L730) | CASE / P11 / COMPARISON | `isnull(sm.FlagSyariah,0)` : `BIT` | `1` : `int` | `isnull(sm.FlagSyariah,0)` → `INT` | `isnull(sm.FlagSyariah,0) = 1` |
| IC-010894 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptPayCouponORIDetail:735](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L735) | CASE / P11 / COMPARISON | `isnull(sm.FlagSyariah,0)` : `BIT` | `1` : `int` | `isnull(sm.FlagSyariah,0)` → `INT` | `isnull(sm.FlagSyariah,0) = 1` |
| IC-011025 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptPayCouponORITaxAmnesty:659](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L659) | CASE / P11 / COMPARISON | `isnull(sm.FlagSyariah,0)` : `BIT` | `1` : `int` | `isnull(sm.FlagSyariah,0)` → `INT` | `isnull(sm.FlagSyariah,0) = 1` |
| IC-011027 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptPayCouponORITaxAmnesty:665](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L665) | CASE / P11 / COMPARISON | `isnull(sm.FlagSyariah,0)` : `BIT` | `1` : `int` | `isnull(sm.FlagSyariah,0)` → `INT` | `isnull(sm.FlagSyariah,0) = 1` |
| IC-011235 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:537](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L537) | CASE / P11 / COMPARISON | `ISNULL(sm.FlagSyariah,0)` : `BIT` | `1` : `int` | `ISNULL(sm.FlagSyariah,0)` → `INT` | `ISNULL(sm.FlagSyariah,0) = 1` |
| IC-011237 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:543](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L543) | CASE / P11 / COMPARISON | `ISNULL(sm.FlagSyariah,0)` : `BIT` | `1` : `int` | `ISNULL(sm.FlagSyariah,0)` → `INT` | `ISNULL(sm.FlagSyariah,0) = 1` |
| IC-011435 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptPayCouponORITaxAmnesty_Backdated:476](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty_Backdated.sql#L476) | CASE / P11 / COMPARISON | `isnull(sm.FlagSyariah,0)` : `BIT` | `1` : `int` | `isnull(sm.FlagSyariah,0)` → `INT` | `isnull(sm.FlagSyariah,0) = 1` |
| IC-011437 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptPayCouponORITaxAmnesty_Backdated:482](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty_Backdated.sql#L482) | CASE / P11 / COMPARISON | `isnull(sm.FlagSyariah,0)` : `BIT` | `1` : `int` | `isnull(sm.FlagSyariah,0)` → `INT` | `isnull(sm.FlagSyariah,0) = 1` |
| IC-013287 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_ListOutstanding:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListOutstanding.sql#L47) | CASE / P11 / COMPARISON | `ISNULL(FlagSyariah,0)` : `BIT` | `0` : `int` | `ISNULL(FlagSyariah,0)` → `INT` | `ISNULL(FlagSyariah,0) = 0` |
| IC-013613 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_PopulateDataRekeningMaterai:73](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningMaterai.sql#L73) | WHERE / P11 / COMPARISON | `FlagSyariah` : `BIT` | `1` : `int` | `FlagSyariah` → `INT` | `FlagSyariah = 1` |
| IC-013634 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_PopulateDataRekeningObligasi:78](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningObligasi.sql#L78) | WHERE / P11 / COMPARISON | `FlagSyariah` : `BIT` | `1` : `int` | `FlagSyariah` → `INT` | `FlagSyariah = 1` |
| IC-013866 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_UpdateTreasuryCustomer_TM:635](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L635) | WHERE / P11 / COMPARISON | `isnull(smtm.FlagSyariah,0)` : `BIT` | `1` : `int` | `isnull(smtm.FlagSyariah,0)` → `INT` | `isnull(smtm.FlagSyariah,0) = 1` |
| IC-013870 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_UpdateTreasuryCustomer_TM:654](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L654) | WHERE / P11 / COMPARISON | `isnull(smtm.FlagSyariah,0)` : `BIT` | `1` : `int` | `isnull(smtm.FlagSyariah,0)` → `INT` | `isnull(smtm.FlagSyariah,0) = 1` |
| IC-013875 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_UpdateTreasuryCustomer_TM:671](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L671) | WHERE / P11 / COMPARISON | `isnull(smtm.FlagSyariah,0)` : `BIT` | `1` : `int` | `isnull(smtm.FlagSyariah,0)` → `INT` | `isnull(smtm.FlagSyariah,0) = 1` |
| IC-013878 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_UpdateTreasuryCustomer_TM:693](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L693) | WHERE / P11 / COMPARISON | `isnull(smtm.FlagSyariah,0)` : `BIT` | `1` : `int` | `isnull(smtm.FlagSyariah,0)` → `INT` | `isnull(smtm.FlagSyariah,0) = 1` |
| IC-013881 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_UpdateTreasuryCustomer_TM:710](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L710) | WHERE / P11 / COMPARISON | `isnull(smtm.FlagSyariah,0)` : `BIT` | `1` : `int` | `isnull(smtm.FlagSyariah,0)` → `INT` | `isnull(smtm.FlagSyariah,0) = 1` |
| IC-013886 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_UpdateTreasuryCustomer_TM:729](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L729) | WHERE / P11 / COMPARISON | `isnull(smtm.FlagSyariah,0)` : `BIT` | `1` : `int` | `isnull(smtm.FlagSyariah,0)` → `INT` | `isnull(smtm.FlagSyariah,0) = 1` |
| IC-013890 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_UpdateTreasuryCustomer_TM:757](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L757) | WHERE / P11 / COMPARISON | `isnull(smtm.FlagSyariah,0)` : `BIT` | `0` : `int` | `isnull(smtm.FlagSyariah,0)` → `INT` | `isnull(smtm.FlagSyariah,0) = 0` |
| IC-013893 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_UpdateTreasuryCustomer_TM:770](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L770) | WHERE / P11 / COMPARISON | `isnull(smtm.FlagSyariah,0)` : `BIT` | `0` : `int` | `isnull(smtm.FlagSyariah,0)` → `INT` | `isnull(smtm.FlagSyariah,0) = 0` |
| IC-013898 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_UpdateTreasuryCustomer_TM:785](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L785) | WHERE / P11 / COMPARISON | `isnull(smtm.FlagSyariah,0)` : `BIT` | `0` : `int` | `isnull(smtm.FlagSyariah,0)` → `INT` | `isnull(smtm.FlagSyariah,0) = 0` |
| IC-013901 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_UpdateTreasuryCustomer_TM:806](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L806) | WHERE / P11 / COMPARISON | `isnull(smtm.FlagSyariah,0)` : `BIT` | `0` : `int` | `isnull(smtm.FlagSyariah,0)` → `INT` | `isnull(smtm.FlagSyariah,0) = 0` |
| IC-013904 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_UpdateTreasuryCustomer_TM:821](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L821) | WHERE / P11 / COMPARISON | `isnull(smtm.FlagSyariah,0)` : `BIT` | `0` : `int` | `isnull(smtm.FlagSyariah,0)` → `INT` | `isnull(smtm.FlagSyariah,0) = 0` |
| IC-013909 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_UpdateTreasuryCustomer_TM:838](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L838) | WHERE / P11 / COMPARISON | `isnull(smtm.FlagSyariah,0)` : `BIT` | `0` : `int` | `isnull(smtm.FlagSyariah,0)` → `INT` | `isnull(smtm.FlagSyariah,0) = 0` |
| IC-000749 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.OMTrxSecPreBook:372](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMTrxSecPreBook.sql#L372) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `a.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(a.FlagSyariah,0)` |
| IC-002711 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L156) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(FlagSyariah, 0)` |
| IC-002909 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SBNSchedPushMailNotifOrderCompleted:157](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L157) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(FlagSyariah,0)` |
| IC-002932 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SBNSchedPushMailNotifOrderORI:140](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderORI.sql#L140) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(FlagSyariah,0)` |
| IC-004128 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSGetMenuSettings:273](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetMenuSettings.sql#L273) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(FlagSyariah, 0)` |
| IC-004316 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSInsertSecurityTransaction_TT_Original:791](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L791) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(sm.FlagSyariah, 0)` |
| IC-006299 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateSettlementMaturity:91](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementMaturity.sql#L91) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(smtm.FlagSyariah, 0)` |
| IC-006301 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateSettlementMaturity:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementMaturity.sql#L98) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(smtm.FlagSyariah, 0)` |
| IC-006427 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateTransactionBalance:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTransactionBalance.sql#L69) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(FlagSyariah,0)` |
| IC-007494 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSettlementKupon:194](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementKupon.sql#L194) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sc.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(sc.FlagSyariah,0)` |
| IC-007497 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSettlementKupon:203](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementKupon.sql#L203) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sc.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(sc.FlagSyariah,0)` |
| IC-007503 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSettlementMaturity:116](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementMaturity.sql#L116) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(smtm.FlagSyariah,0)` |
| IC-007505 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSettlementMaturity:122](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementMaturity.sql#L122) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(smtm.FlagSyariah,0)` |
| IC-010460 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:188](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L188) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(sm.FlagSyariah,0)` |
| IC-010462 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L193) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(sm.FlagSyariah,0)` |
| IC-010481 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:294](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L294) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(sm.FlagSyariah,0)` |
| IC-010483 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:299](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L299) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(sm.FlagSyariah,0)` |
| IC-010502 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:432](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L432) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(sm.FlagSyariah,0)` |
| IC-010504 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:437](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L437) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(sm.FlagSyariah,0)` |
| IC-010523 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:534](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L534) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(sm.FlagSyariah,0)` |
| IC-010525 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:539](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L539) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(sm.FlagSyariah,0)` |
| IC-010893 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORIDetail:730](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L730) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(sm.FlagSyariah,0)` |
| IC-010895 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORIDetail:735](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L735) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(sm.FlagSyariah,0)` |
| IC-011026 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnesty:659](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L659) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(sm.FlagSyariah,0)` |
| IC-011028 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnesty:665](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L665) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(sm.FlagSyariah,0)` |
| IC-011236 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:537](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L537) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(sm.FlagSyariah,0)` |
| IC-011238 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:543](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L543) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(sm.FlagSyariah,0)` |
| IC-011436 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnesty_Backdated:476](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty_Backdated.sql#L476) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(sm.FlagSyariah,0)` |
| IC-011438 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnesty_Backdated:482](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty_Backdated.sql#L482) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(sm.FlagSyariah,0)` |
| IC-013288 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListOutstanding:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListOutstanding.sql#L47) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(FlagSyariah,0)` |
| IC-013341 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L113) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(FlagSyariah, 0)` |
| IC-013867 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_UpdateTreasuryCustomer_TM:635](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L635) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(smtm.FlagSyariah,0)` |
| IC-013871 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_UpdateTreasuryCustomer_TM:654](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L654) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(smtm.FlagSyariah,0)` |
| IC-013876 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_UpdateTreasuryCustomer_TM:671](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L671) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(smtm.FlagSyariah,0)` |
| IC-013879 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_UpdateTreasuryCustomer_TM:693](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L693) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(smtm.FlagSyariah,0)` |
| IC-013882 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_UpdateTreasuryCustomer_TM:710](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L710) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(smtm.FlagSyariah,0)` |
| IC-013887 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_UpdateTreasuryCustomer_TM:729](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L729) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(smtm.FlagSyariah,0)` |
| IC-013891 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_UpdateTreasuryCustomer_TM:757](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L757) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(smtm.FlagSyariah,0)` |
| IC-013894 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_UpdateTreasuryCustomer_TM:770](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L770) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(smtm.FlagSyariah,0)` |
| IC-013899 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_UpdateTreasuryCustomer_TM:785](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L785) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(smtm.FlagSyariah,0)` |
| IC-013902 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_UpdateTreasuryCustomer_TM:806](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L806) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(smtm.FlagSyariah,0)` |
| IC-013905 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_UpdateTreasuryCustomer_TM:821](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L821) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(smtm.FlagSyariah,0)` |
| IC-013910 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_UpdateTreasuryCustomer_TM:838](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L838) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(smtm.FlagSyariah,0)` |

<a id="kolom-flagtransaksibeli"></a>
## FlagTransaksiBeli

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L84). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002708 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:139](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L139) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagTransaksiBeli` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(FlagTransaksiBeli, 0)` |
| IC-013338 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:92](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L92) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagTransaksiBeli` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(FlagTransaksiBeli,0)` |

<a id="kolom-flagtransaksibeliandcashback"></a>
## FlagTransaksiBeliAndCashBack

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L86). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002710 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L141) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagTransaksiBeliAndCashBack` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(FlagTransaksiBeliAndCashBack, 0)` |
| IC-013340 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L94) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagTransaksiBeliAndCashBack` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(FlagTransaksiBeliAndCashBack,0)` |

<a id="kolom-flagwaivefeetransaksiperdana"></a>
## FlagWaiveFeeTransaksiPerdana

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L72). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008769 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailCalculateFee:418](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L418) | WHERE / P11 / COMPARISON | `smtm.FlagWaiveFeeTransaksiPerdana` : `BIT` | `1` : `int` | `smtm.FlagWaiveFeeTransaksiPerdana` → `INT` | `smtm.FlagWaiveFeeTransaksiPerdana = 1` |

<a id="kolom-flagwaivesafekeepingfeenontransaksibeli"></a>
## FlagWaiveSafeKeepingFeeNonTransaksiBeli

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L71). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008865 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailCalculateFee:972](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L972) | WHERE / P11 / COMPARISON | `smtm.FlagWaiveSafeKeepingFeeNonTransaksiBeli` : `BIT` | `1` : `int` | `smtm.FlagWaiveSafeKeepingFeeNonTransaksiBeli` → `INT` | `smtm.FlagWaiveSafeKeepingFeeNonTransaksiBeli = 1` |

<a id="kolom-hargabuyback"></a>
## HargaBuyBack

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L39). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 14 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006198 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBuyBack:99](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L99) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `smtm.HargaBuyBack` : `MONEY` | `100.00` : `decimal` | `smtm.HargaBuyBack` → `DECIMAL` | `smtm.HargaBuyBack/100.00` |
| IC-006201 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBuyBack:100](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L100) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `smtm.HargaBuyBack` : `MONEY` | `100.00` : `decimal` | `smtm.HargaBuyBack` → `DECIMAL` | `smtm.HargaBuyBack/100.00` |
| IC-006203 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBuyBack:101](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L101) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `smtm.HargaBuyBack` : `MONEY` | `100.00` : `decimal` | `smtm.HargaBuyBack` → `DECIMAL` | `smtm.HargaBuyBack/100.00` |
| IC-006232 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBuyBack:194](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L194) | DML_CONDITION / P05 / BINARY_Divide | `smtm.HargaBuyBack` : `MONEY` | `100.00` : `decimal` | `smtm.HargaBuyBack` → `DECIMAL` | `smtm.HargaBuyBack/100.00` |
| IC-005255 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateApprovalBuyBackOM:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalBuyBackOM.sql#L85) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.HargaBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(sm.HargaBuyBack,0)` |
| IC-006199 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBuyBack:99](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L99) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `(smtm.HargaBuyBack/100.00)` : `DECIMAL (combined result; facets unknown)` | `oott.NominalBuyBack` : `MONEY` | `oott.NominalBuyBack` → `DECIMAL` | `(smtm.HargaBuyBack/100.00) * oott.NominalBuyBack` |
| IC-006202 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBuyBack:100](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L100) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `(smtm.HargaBuyBack/100.00)` : `DECIMAL (combined result; facets unknown)` | `oott.NominalBuyBack` : `MONEY` | `oott.NominalBuyBack` → `DECIMAL` | `(smtm.HargaBuyBack/100.00) * oott.NominalBuyBack` |
| IC-006204 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBuyBack:101](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L101) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `(smtm.HargaBuyBack/100.00)` : `DECIMAL (combined result; facets unknown)` | `oott.NominalBuyBack` : `MONEY` | `oott.NominalBuyBack` → `DECIMAL` | `(smtm.HargaBuyBack/100.00) * oott.NominalBuyBack` |
| IC-006206 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBuyBack:102](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L102) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `smtm.HargaBuyBack` : `MONEY` | `100` : `int` | `100` → `MONEY` | `smtm.HargaBuyBack/100` |
| IC-006233 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBuyBack:194](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L194) | DML_CONDITION / P05 / BINARY_Multiply | `(smtm.HargaBuyBack/100.00)` : `DECIMAL (combined result; facets unknown)` | `isnull(NominalBuyBack, 0)` : `money` | `isnull(NominalBuyBack, 0)` → `DECIMAL` | `isnull(NominalBuyBack, 0) * (smtm.HargaBuyBack/100.00)` |
| IC-006754 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessBuyBack:54](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBuyBack.sql#L54) | DML_CONDITION / P05 / BINARY_Divide | `smtm.HargaBuyBack` : `MONEY` | `100` : `int` | `100` → `MONEY` | `smtm.HargaBuyBack/100` |
| IC-007426 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessSettlementBuyBack:160](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBuyBack.sql#L160) | DML_CONDITION / P05 / BINARY_Divide | `smtm.HargaBuyBack` : `MONEY` | `100` : `int` | `100` → `MONEY` | `smtm.HargaBuyBack/100` |
| IC-007427 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessSettlementBuyBack:161](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBuyBack.sql#L161) | DML_CONDITION / P05 / BINARY_Divide | `smtm.HargaBuyBack` : `MONEY` | `100` : `int` | `100` → `MONEY` | `smtm.HargaBuyBack/100` |
| IC-007435 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessSettlementBuyBack:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBuyBack.sql#L168) | CASE / P05 / BINARY_Divide | `smtm.HargaBuyBack` : `MONEY` | `100` : `int` | `100` → `MONEY` | `smtm.HargaBuyBack/100` |

<a id="kolom-hargabuyback2"></a>
## HargaBuyBack2

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L55). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006507 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateTrxBuyBackPerdana:134](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L134) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.HargaBuyBack2` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sm.HargaBuyBack2, 0)` |

<a id="kolom-imbalhasil"></a>
## ImbalHasil

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L105). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 15 lokasi / 13 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003210 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSAmendSecurityTransaction_TT:637](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L637) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `sm.ImbalHasil` : `SMALLINT` | `@nImbalan` : `INT` | `sm.ImbalHasil` → `INT` | `@nImbalan = sm.ImbalHasil` |
| IC-004318 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSInsertSecurityTransaction_TT_Original:809](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L809) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `sm.ImbalHasil` : `SMALLINT` | `@nImbalan` : `INT` | `sm.ImbalHasil` → `INT` | `@nImbalan = sm.ImbalHasil` |
| IC-008734 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailCalculateFee:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L216) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `ImbalHasil` : `SMALLINT` | `@nImbalan` : `INT` | `ImbalHasil` → `INT` | `@nImbalan = ImbalHasil` |
| IC-010800 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptPayCouponORIDetail:138](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L138) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `ImbalHasil` : `SMALLINT` | `@nImbalan` : `INT` | `ImbalHasil` → `INT` | `@nImbalan = ImbalHasil` |
| IC-010930 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptPayCouponORITaxAmnesty:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L130) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `ImbalHasil` : `SMALLINT` | `@nImbalan` : `INT` | `ImbalHasil` → `INT` | `@nImbalan = ImbalHasil` |
| IC-011130 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L115) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `ImbalHasil` : `SMALLINT` | `@nImbalan` : `INT` | `ImbalHasil` → `INT` | `@nImbalan = ImbalHasil` |
| IC-011344 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptPayCouponORITaxAmnesty_Backdated:116](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty_Backdated.sql#L116) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `ImbalHasil` : `SMALLINT` | `@nImbalan` : `INT` | `ImbalHasil` → `INT` | `@nImbalan = ImbalHasil` |
| IC-011930 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSSaveParamSettlementDate:260](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamSettlementDate.sql#L260) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `sm.ImbalHasil` : `SMALLINT` | `ImbalHasil` : `INT` | `sm.ImbalHasil` → `INT` | `ImbalHasil = sm.ImbalHasil` |
| IC-003207 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSAmendSecurityTransaction_TT:627](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L627) | CASE / P11 / COMPARISON | `sm.ImbalHasil` : `SMALLINT` | `2` : `int` | `sm.ImbalHasil` → `INT` | `sm.ImbalHasil = 2` |
| IC-004317 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSInsertSecurityTransaction_TT_Original:805](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L805) | CASE / P11 / COMPARISON | `sm.ImbalHasil` : `SMALLINT` | `2` : `int` | `sm.ImbalHasil` → `INT` | `sm.ImbalHasil = 2` |
| IC-005407 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateEditHargaModalTrx:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEditHargaModalTrx.sql#L16) | CASE / P11 / COMPARISON | `sm.ImbalHasil` : `SMALLINT` | `2` : `int` | `sm.ImbalHasil` → `INT` | `sm.ImbalHasil = 2` |
| IC-011834 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveParamHargaModalNew:348](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L348) | WHERE / P11 / COMPARISON | `tm.ImbalHasil` : `SMALLINT` | `2` : `int` | `tm.ImbalHasil` → `INT` | `tm.ImbalHasil = 2` |
| IC-002713 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:174](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L174) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ImbalHasil` : `SMALLINT` | `1` : `int` | `1` → `SMALLINT` | `ISNULL(ImbalHasil, 1)` |
| IC-004133 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSGetMenuSettings:292](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetMenuSettings.sql#L292) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ImbalHasil` : `SMALLINT` | `1` : `int` | `1` → `SMALLINT` | `ISNULL(ImbalHasil, 1)` |
| IC-013343 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L137) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ImbalHasil` : `SMALLINT` | `1` : `int` | `1` → `SMALLINT` | `ISNULL(ImbalHasil, 1)` |

<a id="kolom-interesttypeid"></a>
## InterestTypeID

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004483 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSMaintenanceParamHargaModal:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L410) | WHERE / P11 / COMPARISON | `ms.InterestTypeID` : `SMALLINT` | `2` : `int` | `ms.InterestTypeID` → `INT` | `ms.InterestTypeID = 2` |
| IC-004484 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSMaintenanceParamHargaModal:421](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L421) | WHERE / P11 / COMPARISON | `ms.InterestTypeID` : `SMALLINT` | `2` : `int` | `ms.InterestTypeID` → `INT` | `ms.InterestTypeID = 2` |
| IC-012810 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSValidateUploadParamHargaModalNew:516](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L516) | WHERE / P11 / COMPARISON | `sm.InterestTypeID` : `SMALLINT` | `2` : `int` | `sm.InterestTypeID` → `INT` | `sm.InterestTypeID = 2` |
| IC-012841 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSValidateUploadParamHargaModalNew:887](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L887) | WHERE / P11 / COMPARISON | `tm.InterestTypeID` : `SMALLINT` | `2` : `int` | `tm.InterestTypeID` → `INT` | `tm.InterestTypeID = 2` |
| IC-013113 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertSecurityMaster_TM:93](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertSecurityMaster_TM.sql#L93) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitymaster_tm.InterestTypeID` : `SMALLINT` | `InterestTypeID` : `varchar(5)` | `InterestTypeID` → `SMALLINT` | `retail.dbo.securitymaster_tm.InterestTypeID <- InterestTypeID` |
| IC-013760 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateSecurityMaster_TM:87](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateSecurityMaster_TM.sql#L87) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Trx.InterestTypeID` : `SMALLINT` | `t.InterestTypeID` : `varchar(20)` | `t.InterestTypeID` → `SMALLINT` | `Trx.InterestTypeID=t.InterestTypeID` |

<a id="kolom-iscorporatebond"></a>
## IsCorporateBond

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L87). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 140 lokasi / 56 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000151 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.CustodyReportMaintainAcquisition:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportMaintainAcquisition.sql#L110) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `IsCorporateBond` : `BIT` | `#tmp_SecurityMaster_TM.IsCorporateBond` : `VARCHAR(5)` | `IsCorporateBond` → `VARCHAR(5)` | `#tmp_SecurityMaster_TM.IsCorporateBond <- IsCorporateBond` |
| IC-000190 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L111) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `1` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 1` |
| IC-000192 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:112](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L112) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `1` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 1` |
| IC-000194 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L113) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `0` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 0` |
| IC-000197 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L114) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `0` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 0` |
| IC-000200 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L115) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `0` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 0` |
| IC-000210 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:163](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L163) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `0` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 0` |
| IC-000217 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:197](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L197) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `1` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 1` |
| IC-000219 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:198](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L198) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `1` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 1` |
| IC-000221 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:199](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L199) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `0` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 0` |
| IC-000223 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:200](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L200) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `0` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 0` |
| IC-000224 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:231](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L231) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `0` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 0` |
| IC-000228 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:263](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L263) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `1` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 1` |
| IC-000230 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:264](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L264) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `0` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 0` |
| IC-000232 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L265) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `0` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 0` |
| IC-000236 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:296](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L296) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `0` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 0` |
| IC-000238 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:336](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L336) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-000240 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L337) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-000242 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:342](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L342) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-000244 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:343](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L343) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-000247 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstructionTransferAsset:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstructionTransferAsset.sql#L43) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-000249 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstructionTransferAsset:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstructionTransferAsset.sql#L44) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-000251 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstructionTransferAsset:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstructionTransferAsset.sql#L45) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-000253 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstructionTransferAsset:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstructionTransferAsset.sql#L46) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-000656 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMGetListProductSekunder:175](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L175) | CASE / P11 / COMPARISON | `IsCorporateBond` : `BIT` | `1` : `int` | `IsCorporateBond` → `INT` | `IsCorporateBond = 1` |
| IC-000657 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMGetListProductSekunder:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L192) | CASE / P11 / COMPARISON | `IsCorporateBond` : `BIT` | `1` : `int` | `IsCorporateBond` → `INT` | `IsCorporateBond = 1` |
| IC-000658 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMGetListProductSekunder:202](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L202) | CASE / P11 / COMPARISON | `IsCorporateBond` : `BIT` | `1` : `int` | `IsCorporateBond` → `INT` | `IsCorporateBond = 1` |
| IC-000659 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMGetListProductSekunder:226](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L226) | CASE / P11 / COMPARISON | `IsCorporateBond` : `BIT` | `1` : `int` | `IsCorporateBond` → `INT` | `IsCorporateBond = 1` |
| IC-000969 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.ORISBNIPopulateEarlyRedemption:87](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateEarlyRedemption.sql#L87) | CASE / P11 / COMPARISON | `IsCorporateBond` : `BIT` | `1` : `int` | `IsCorporateBond` → `INT` | `IsCorporateBond = 1` |
| IC-000970 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.ORISBNIPopulateEarlyRedemption:88](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateEarlyRedemption.sql#L88) | CASE / P11 / COMPARISON | `IsCorporateBond` : `BIT` | `0` : `int` | `IsCorporateBond` → `INT` | `IsCorporateBond = 0` |
| IC-001024 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.ORISBNISubmitEarlyRedemption:601](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNISubmitEarlyRedemption.sql#L601) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-001025 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.ORISBNISubmitEarlyRedemption:602](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNISubmitEarlyRedemption.sql#L602) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-002968 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedUpdateStatusTimeOut:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedUpdateStatusTimeOut.sql#L47) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-002969 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedUpdateStatusTimeOut:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedUpdateStatusTimeOut.sql#L48) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-003107 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SecurityProcessLelangToTransaction:303](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessLelangToTransaction.sql#L303) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-003109 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SecurityProcessLelangToTransaction:304](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessLelangToTransaction.sql#L304) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-003470 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCalculateCorporateBond:54](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateCorporateBond.sql#L54) | WHERE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-003474 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCalculateCorporateBond:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateCorporateBond.sql#L71) | WHERE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-003550 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCalculateTransactionSekunder:38](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateTransactionSekunder.sql#L38) | WHERE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-003559 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCalculateTransactionSekunder:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateTransactionSekunder.sql#L56) | WHERE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-003569 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCalculateTrxSecCorporate:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateTrxSecCorporate.sql#L29) | WHERE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-003574 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCalculateTrxSecCorporate:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateTrxSecCorporate.sql#L41) | WHERE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-004230 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSInsertEarlyRedemption_TT:149](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertEarlyRedemption_TT.sql#L149) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-004231 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSInsertEarlyRedemption_TT:150](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertEarlyRedemption_TT.sql#L150) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-004455 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSMaintenanceParamHargaModal:144](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L144) | WHERE / P11 / COMPARISON | `IsCorporateBond` : `BIT` | `1` : `int` | `IsCorporateBond` → `INT` | `IsCorporateBond = 1` |
| IC-004465 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSMaintenanceParamHargaModal:187](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L187) | CASE / P11 / COMPARISON | `mst.IsCorporateBond` : `BIT` | `1` : `int` | `mst.IsCorporateBond` → `INT` | `mst.IsCorporateBond = 1` |
| IC-004466 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSMaintenanceParamHargaModal:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L192) | CASE / P11 / COMPARISON | `mst.IsCorporateBond` : `BIT` | `1` : `int` | `mst.IsCorporateBond` → `INT` | `mst.IsCorporateBond = 1` |
| IC-004488 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSMaintenanceParamHargaModal:496](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L496) | WHERE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-004492 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSMaintenanceParamHargaModal:549](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L549) | WHERE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-004792 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSORIGetProductDetail:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIGetProductDetail.sql#L107) | CASE / P11 / COMPARISON | `IsCorporateBond` : `BIT` | `0` : `int` | `IsCorporateBond` → `INT` | `IsCorporateBond = 0` |
| IC-004959 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPTOSFilterlistProduk:11](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPTOSFilterlistProduk.sql#L11) | WHERE / P11 / COMPARISON | `isnull(IsCorporateBond,0)` : `BIT` | `1` : `int` | `isnull(IsCorporateBond,0)` → `INT` | `isnull(IsCorporateBond,0) = 1` |
| IC-005175 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamRepoValidate:633](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamRepoValidate.sql#L633) | WHERE / P11 / COMPARISON | `IsCorporateBond` : `BIT` | `0` : `int` | `IsCorporateBond` → `INT` | `IsCorporateBond <> 0` |
| IC-005349 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDataParamDeviasiTotalSpread:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDataParamDeviasiTotalSpread.sql#L25) | CASE / P11 / COMPARISON | `IsCorporateBond` : `BIT` | `1` : `int` | `IsCorporateBond` → `INT` | `IsCorporateBond = 1` |
| IC-005350 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDataParamDeviasiTotalSpread:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDataParamDeviasiTotalSpread.sql#L30) | CASE / P11 / COMPARISON | `IsCorporateBond` : `BIT` | `1` : `int` | `IsCorporateBond` → `INT` | `IsCorporateBond = 1` |
| IC-005351 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDataParamDeviasiTotalSpread:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDataParamDeviasiTotalSpread.sql#L41) | CASE / P11 / COMPARISON | `IsCorporateBond` : `BIT` | `1` : `int` | `IsCorporateBond` → `INT` | `IsCorporateBond = 1` |
| IC-005357 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDataTransaksiSuratBerharga:97](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDataTransaksiSuratBerharga.sql#L97) | WHERE / P11 / COMPARISON | `IsCorporateBond` : `BIT` | `1` : `int` | `IsCorporateBond` → `INT` | `IsCorporateBond = 1` |
| IC-005416 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateFOPTransferAsset:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateFOPTransferAsset.sql#L69) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-005419 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateFOPTransferAsset:73](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateFOPTransferAsset.sql#L73) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-005429 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateFilterCorporateBonds:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateFilterCorporateBonds.sql#L26) | WHERE / P11 / COMPARISON | `IsCorporateBond` : `BIT` | `1` : `int` | `IsCorporateBond` → `INT` | `IsCorporateBond = 1` |
| IC-005431 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateFilterCorporateBonds:35](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateFilterCorporateBonds.sql#L35) | WHERE / P11 / COMPARISON | `IsCorporateBond` : `BIT` | `1` : `int` | `IsCorporateBond` → `INT` | `IsCorporateBond = 1` |
| IC-005980 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateReportPTOSProduk:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateReportPTOSProduk.sql#L17) | WHERE / P11 / COMPARISON | `isnull(IsCorporateBond,0)` : `BIT` | `1` : `int` | `isnull(IsCorporateBond,0)` → `INT` | `isnull(IsCorporateBond,0) = 1` |
| IC-006396 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateStatusCorporateBonds:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateStatusCorporateBonds.sql#L66) | WHERE / P11 / COMPARISON | `tm.IsCorporateBond` : `BIT` | `1` : `int` | `tm.IsCorporateBond` → `INT` | `tm.IsCorporateBond = 1` |
| IC-006909 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessInsertORIOrderCommit:324](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertORIOrderCommit.sql#L324) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-006910 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessInsertORIOrderCommit:325](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertORIOrderCommit.sql#L325) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-006952 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessInsertSecurityTransactionCommit:722](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertSecurityTransactionCommit.sql#L722) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-006954 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessInsertSecurityTransactionCommit:723](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertSecurityTransactionCommit.sql#L723) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-007115 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSOuts:134](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSOuts.sql#L134) | JOIN_ON / P11 / COMPARISON | `smtm.IsCorporateBond` : `BIT` | `1` : `int` | `smtm.IsCorporateBond` → `INT` | `smtm.IsCorporateBond = 1` |
| IC-007168 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSTrx:164](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSTrx.sql#L164) | JOIN_ON / P11 / COMPARISON | `smtm.IsCorporateBond` : `BIT` | `1` : `int` | `smtm.IsCorporateBond` → `INT` | `smtm.IsCorporateBond = 1` |
| IC-007448 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSettlementBuyBack:309](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBuyBack.sql#L309) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-007449 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSettlementBuyBack:310](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBuyBack.sql#L310) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-007550 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSettlementOrder:447](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementOrder.sql#L447) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-007552 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSettlementOrder:448](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementOrder.sql#L448) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-007783 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessUploadTransaksiSuratBerharga:354](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessUploadTransaksiSuratBerharga.sql#L354) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-007785 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessUploadTransaksiSuratBerharga:355](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessUploadTransaksiSuratBerharga.sql#L355) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-007994 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductParamSrc:38](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductParamSrc.sql#L38) | WHERE / P11 / COMPARISON | `tm.IsCorporateBond` : `BIT` | `0` : `int` | `tm.IsCorporateBond` → `INT` | `tm.IsCorporateBond = 0` |
| IC-007996 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductParamSrc:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductParamSrc.sql#L41) | WHERE / P11 / COMPARISON | `IsCorporateBond` : `BIT` | `0` : `int` | `IsCorporateBond` → `INT` | `IsCorporateBond = 0` |
| IC-007997 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductParamSrc:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductParamSrc.sql#L51) | WHERE / P11 / COMPARISON | `tm.IsCorporateBond` : `BIT` | `0` : `int` | `tm.IsCorporateBond` → `INT` | `tm.IsCorporateBond = 0` |
| IC-007999 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductParamSrc:54](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductParamSrc.sql#L54) | WHERE / P11 / COMPARISON | `IsCorporateBond` : `BIT` | `0` : `int` | `IsCorporateBond` → `INT` | `IsCorporateBond = 0` |
| IC-008000 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductParamSrc:60](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductParamSrc.sql#L60) | WHERE / P11 / COMPARISON | `tm.IsCorporateBond` : `BIT` | `0` : `int` | `tm.IsCorporateBond` → `INT` | `tm.IsCorporateBond = 0` |
| IC-008002 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductParamSrc:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductParamSrc.sql#L63) | WHERE / P11 / COMPARISON | `IsCorporateBond` : `BIT` | `0` : `int` | `IsCorporateBond` → `INT` | `IsCorporateBond = 0` |
| IC-008003 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductParamSrc:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductParamSrc.sql#L69) | WHERE / P11 / COMPARISON | `tm.IsCorporateBond` : `BIT` | `0` : `int` | `tm.IsCorporateBond` → `INT` | `tm.IsCorporateBond = 0` |
| IC-008005 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductParamSrc:72](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductParamSrc.sql#L72) | WHERE / P11 / COMPARISON | `IsCorporateBond` : `BIT` | `0` : `int` | `IsCorporateBond` → `INT` | `IsCorporateBond = 0` |
| IC-008006 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductParamSrc:79](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductParamSrc.sql#L79) | WHERE / P11 / COMPARISON | `tm.IsCorporateBond` : `BIT` | `0` : `int` | `tm.IsCorporateBond` → `INT` | `tm.IsCorporateBond = 0` |
| IC-008008 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductParamSrc:82](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductParamSrc.sql#L82) | WHERE / P11 / COMPARISON | `IsCorporateBond` : `BIT` | `0` : `int` | `IsCorporateBond` → `INT` | `IsCorporateBond = 0` |
| IC-008267 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:122](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L122) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `1` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 1` |
| IC-008275 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L147) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `0` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 0` |
| IC-008287 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:223](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L223) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `1` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 1` |
| IC-008295 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:248](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L248) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `0` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 0` |
| IC-008307 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:325](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L325) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `1` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 1` |
| IC-008315 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:350](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L350) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `0` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 0` |
| IC-008327 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:426](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L426) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `1` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 1` |
| IC-008335 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:451](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L451) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `0` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 0` |
| IC-008351 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:514](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L514) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-008364 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:576](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L576) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `1` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 1` |
| IC-008370 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:604](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L604) | CASE / P11 / COMPARISON | `sttm.IsCorporateBond` : `BIT` | `0` : `int` | `sttm.IsCorporateBond` → `INT` | `sttm.IsCorporateBond = 0` |
| IC-008489 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTrxCorporateBOND:32](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L32) | JOIN_ON / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-008671 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailAutoRelBlock:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailAutoRelBlock.sql#L69) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-008673 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailAutoRelBlock:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailAutoRelBlock.sql#L70) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-008690 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailAutoRelBlock_urgent:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailAutoRelBlock_urgent.sql#L69) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-008692 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailAutoRelBlock_urgent:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailAutoRelBlock_urgent.sql#L70) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-008922 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailCalculateFee2:208](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee2.sql#L208) | CASE / P11 / COMPARISON | `mst.IsCorporateBond` : `BIT` | `1` : `int` | `mst.IsCorporateBond` → `INT` | `mst.IsCorporateBond = 1` |
| IC-008926 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailCalculateFee2:215](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee2.sql#L215) | CASE / P11 / COMPARISON | `mst.IsCorporateBond` : `BIT` | `1` : `int` | `mst.IsCorporateBond` → `INT` | `mst.IsCorporateBond = 1` |
| IC-009026 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailORIAllotment:307](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L307) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-009028 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailORIAllotment:308](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L308) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-010226 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptLaporanPajakBankBuy:76](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L76) | CASE / P11 / COMPARISON | `smtm.IsCorporateBond` : `BIT` | `1` : `int` | `smtm.IsCorporateBond` → `INT` | `smtm.IsCorporateBond = 1` |
| IC-010252 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptLaporanPajakBankBuy:188](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L188) | CASE / P11 / COMPARISON | `smtm.IsCorporateBond` : `BIT` | `1` : `int` | `smtm.IsCorporateBond` → `INT` | `smtm.IsCorporateBond = 1` |
| IC-010268 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptLaporanPajakBankBuy:305](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L305) | CASE / P11 / COMPARISON | `smtm.IsCorporateBond` : `BIT` | `1` : `int` | `smtm.IsCorporateBond` → `INT` | `smtm.IsCorporateBond = 1` |
| IC-010291 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptLaporanPajakBankBuy:384](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L384) | CASE / P11 / COMPARISON | `smtm.IsCorporateBond` : `BIT` | `1` : `int` | `smtm.IsCorporateBond` → `INT` | `smtm.IsCorporateBond = 1` |
| IC-010803 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORIDetail:149](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L149) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-010936 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnesty:143](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L143) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-012250 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSplitOrder:317](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSplitOrder.sql#L317) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-012251 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSplitOrder:318](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSplitOrder.sql#L318) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-012619 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdatePostingResult:1111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L1111) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-012621 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdatePostingResult:1112](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L1112) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-012731 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSValidateUploadParamHargaModal:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModal.sql#L107) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-012734 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSValidateUploadParamHargaModal:112](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModal.sql#L112) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-012759 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSValidateUploadParamHargaModalNew:229](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L229) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-012762 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSValidateUploadParamHargaModalNew:238](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L238) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-013109 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_InsertORI_Order_TT_Original:716](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORI_Order_TT_Original.sql#L716) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-013110 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_InsertORI_Order_TT_Original:717](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORI_Order_TT_Original.sql#L717) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-013142 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_InsertSecurityTransaction_TT:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertSecurityTransaction_TT.sql#L225) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-013144 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_InsertSecurityTransaction_TT:226](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertSecurityTransaction_TT.sql#L226) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-013279 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_InsertTrx2ndSecurityTransaction_TT:891](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTrx2ndSecurityTransaction_TT.sql#L891) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-013281 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_InsertTrx2ndSecurityTransaction_TT:892](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTrx2ndSecurityTransaction_TT.sql#L892) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-013987 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_updateStatusORI_Order_Temp:194](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateStatusORI_Order_Temp.sql#L194) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-013988 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_updateStatusORI_Order_Temp:195](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateStatusORI_Order_Temp.sql#L195) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-014009 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_updateTopUpORI_Order_Temp:140](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateTopUpORI_Order_Temp.sql#L140) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 0` |
| IC-014010 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_updateTopUpORI_Order_Temp:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateTopUpORI_Order_Temp.sql#L141) | CASE / P11 / COMPARISON | `sm.IsCorporateBond` : `BIT` | `1` : `int` | `sm.IsCorporateBond` → `INT` | `sm.IsCorporateBond = 1` |
| IC-002428 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulateParamHargaModalNew:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateParamHargaModalNew.sql#L168) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(sm.IsCorporateBond, 0)` |
| IC-002706 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:120](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L120) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `IsCorporateBond` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(IsCorporateBond, 0)` |
| IC-004960 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPTOSFilterlistProduk:11](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPTOSFilterlistProduk.sql#L11) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `IsCorporateBond` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(IsCorporateBond,0)` |
| IC-005906 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateParamHargaModalNew:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamHargaModalNew.sql#L225) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(sm.IsCorporateBond, 0)` |
| IC-005981 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateReportPTOSProduk:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateReportPTOSProduk.sql#L17) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `IsCorporateBond` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(IsCorporateBond,0)` |
| IC-006508 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateTrxBuyBackPerdana:136](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L136) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(sm.IsCorporateBond, 0)` |
| IC-006782 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessBuyBackOM:172](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBuyBackOM.sql#L172) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.IsCorporateBond` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(sm.IsCorporateBond,0)` |
| IC-008733 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRetailCalculateFee:207](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L207) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `IsCorporateBond` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(IsCorporateBond,0)` |
| IC-010799 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORIDetail:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L137) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `IsCorporateBond` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(IsCorporateBond,0)` |
| IC-010929 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnesty:129](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L129) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `IsCorporateBond` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(IsCorporateBond,0)` |
| IC-011129 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L114) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `IsCorporateBond` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(IsCorporateBond, 0)` |
| IC-011343 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnesty_Backdated:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty_Backdated.sql#L115) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `IsCorporateBond` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(IsCorporateBond,0)` |

<a id="kolom-isusefortransaksi"></a>
## IsUseForTransaksi

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L88). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 36 lokasi / 18 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000661 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMGetListProductSekunder:237](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L237) | WHERE / P11 / COMPARISON | `IsUseForTransaksi` : `BIT` | `1` : `int` | `IsUseForTransaksi` → `INT` | `IsUseForTransaksi = 1` |
| IC-002972 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SRBIProcessEmailNotifMaturityDateH1:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH1.sql#L64) | WHERE / P11 / COMPARISON | `IsUseForTransaksi` : `BIT` | `1` : `int` | `IsUseForTransaksi` → `INT` | `IsUseForTransaksi = 1` |
| IC-002983 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SRBIProcessEmailNotifMaturityDateH3:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH3.sql#L64) | WHERE / P11 / COMPARISON | `IsUseForTransaksi` : `BIT` | `1` : `int` | `IsUseForTransaksi` → `INT` | `IsUseForTransaksi = 1` |
| IC-002994 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SRBIProcessEmailNotifMaturityDateH7:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH7.sql#L64) | WHERE / P11 / COMPARISON | `IsUseForTransaksi` : `BIT` | `1` : `int` | `IsUseForTransaksi` → `INT` | `IsUseForTransaksi = 1` |
| IC-005174 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamRepoValidate:619](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamRepoValidate.sql#L619) | WHERE / P11 / COMPARISON | `IsUseForTransaksi` : `BIT` | `1` : `int` | `IsUseForTransaksi` → `INT` | `IsUseForTransaksi <> 1` |
| IC-005423 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateFOPTransferAsset:82](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateFOPTransferAsset.sql#L82) | WHERE / P11 / COMPARISON | `sm.IsUseForTransaksi` : `BIT` | `1` : `int` | `sm.IsUseForTransaksi` → `INT` | `sm.IsUseForTransaksi = 1` |
| IC-005430 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateFilterCorporateBonds:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateFilterCorporateBonds.sql#L27) | WHERE / P11 / COMPARISON | `IsUseForTransaksi` : `BIT` | `1` : `int` | `IsUseForTransaksi` → `INT` | `IsUseForTransaksi = 1` |
| IC-005432 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateFilterCorporateBonds:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateFilterCorporateBonds.sql#L36) | WHERE / P11 / COMPARISON | `IsUseForTransaksi` : `BIT` | `1` : `int` | `IsUseForTransaksi` → `INT` | `IsUseForTransaksi = 1` |
| IC-006397 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateStatusCorporateBonds:67](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateStatusCorporateBonds.sql#L67) | WHERE / P11 / COMPARISON | `tm.IsUseForTransaksi` : `BIT` | `1` : `int` | `tm.IsUseForTransaksi` → `INT` | `tm.IsUseForTransaksi = 1` |
| IC-007995 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductParamSrc:40](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductParamSrc.sql#L40) | WHERE / P11 / COMPARISON | `IsUseForTransaksi` : `BIT` | `1` : `int` | `IsUseForTransaksi` → `INT` | `IsUseForTransaksi = 1` |
| IC-007998 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductParamSrc:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductParamSrc.sql#L53) | WHERE / P11 / COMPARISON | `IsUseForTransaksi` : `BIT` | `1` : `int` | `IsUseForTransaksi` → `INT` | `IsUseForTransaksi = 1` |
| IC-008001 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductParamSrc:62](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductParamSrc.sql#L62) | WHERE / P11 / COMPARISON | `IsUseForTransaksi` : `BIT` | `1` : `int` | `IsUseForTransaksi` → `INT` | `IsUseForTransaksi = 1` |
| IC-008004 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductParamSrc:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductParamSrc.sql#L71) | WHERE / P11 / COMPARISON | `IsUseForTransaksi` : `BIT` | `1` : `int` | `IsUseForTransaksi` → `INT` | `IsUseForTransaksi = 1` |
| IC-008007 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductParamSrc:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductParamSrc.sql#L81) | WHERE / P11 / COMPARISON | `IsUseForTransaksi` : `BIT` | `1` : `int` | `IsUseForTransaksi` → `INT` | `IsUseForTransaksi = 1` |
| IC-008010 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductSrc:76](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductSrc.sql#L76) | WHERE / P11 / COMPARISON | `tm.IsUseForTransaksi` : `BIT` | `1` : `int` | `tm.IsUseForTransaksi` → `INT` | `tm.IsUseForTransaksi = 1` |
| IC-008013 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductSrc:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductSrc.sql#L105) | WHERE / P11 / COMPARISON | `tm.IsUseForTransaksi` : `BIT` | `1` : `int` | `tm.IsUseForTransaksi` → `INT` | `tm.IsUseForTransaksi = 1` |
| IC-008016 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductSrc:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductSrc.sql#L130) | WHERE / P11 / COMPARISON | `tm.IsUseForTransaksi` : `BIT` | `1` : `int` | `tm.IsUseForTransaksi` → `INT` | `tm.IsUseForTransaksi = 1` |
| IC-008019 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductSrc:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductSrc.sql#L156) | WHERE / P11 / COMPARISON | `tm.IsUseForTransaksi` : `BIT` | `1` : `int` | `tm.IsUseForTransaksi` → `INT` | `tm.IsUseForTransaksi = 1` |
| IC-008022 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductSrc:182](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductSrc.sql#L182) | WHERE / P11 / COMPARISON | `tm.IsUseForTransaksi` : `BIT` | `1` : `int` | `tm.IsUseForTransaksi` → `INT` | `tm.IsUseForTransaksi = 1` |
| IC-008131 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoSecNoSrc:42](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSecNoSrc.sql#L42) | WHERE / P11 / COMPARISON | `tm.IsUseForTransaksi` : `BIT` | `1` : `int` | `tm.IsUseForTransaksi` → `INT` | `tm.IsUseForTransaksi = 1` |
| IC-008132 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoSecNoSrc:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSecNoSrc.sql#L56) | WHERE / P11 / COMPARISON | `tm.IsUseForTransaksi` : `BIT` | `1` : `int` | `tm.IsUseForTransaksi` → `INT` | `tm.IsUseForTransaksi = 1` |
| IC-008133 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoSecNoSrc:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSecNoSrc.sql#L66) | WHERE / P11 / COMPARISON | `tm.IsUseForTransaksi` : `BIT` | `1` : `int` | `tm.IsUseForTransaksi` → `INT` | `tm.IsUseForTransaksi = 1` |
| IC-008134 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoSecNoSrc:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSecNoSrc.sql#L77) | WHERE / P11 / COMPARISON | `tm.IsUseForTransaksi` : `BIT` | `1` : `int` | `tm.IsUseForTransaksi` → `INT` | `tm.IsUseForTransaksi = 1` |
| IC-008135 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoSecNoSrc:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSecNoSrc.sql#L86) | WHERE / P11 / COMPARISON | `tm.IsUseForTransaksi` : `BIT` | `1` : `int` | `tm.IsUseForTransaksi` → `INT` | `tm.IsUseForTransaksi = 1` |
| IC-011691 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveParamHargaModal:164](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModal.sql#L164) | WHERE / P11 / COMPARISON | `IsUseForTransaksi` : `BIT` | `1` : `int` | `IsUseForTransaksi` → `INT` | `IsUseForTransaksi = 1` |
| IC-011931 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveParamSettlementDate:267](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamSettlementDate.sql#L267) | WHERE / P11 / COMPARISON | `IsUseForTransaksi` : `BIT` | `1` : `int` | `IsUseForTransaksi` → `INT` | `IsUseForTransaksi = 1` |
| IC-011932 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveParamSettlementDate:276](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamSettlementDate.sql#L276) | WHERE / P11 / COMPARISON | `IsUseForTransaksi` : `BIT` | `1` : `int` | `IsUseForTransaksi` → `INT` | `IsUseForTransaksi = 1` |
| IC-012428 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSTransactionSrcSecId:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransactionSrcSecId.sql#L31) | WHERE / P11 / COMPARISON | `IsUseForTransaksi` : `BIT` | `1` : `int` | `IsUseForTransaksi` → `INT` | `IsUseForTransaksi = 1` |
| IC-012430 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSTransactionSrcSecId:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransactionSrcSecId.sql#L41) | WHERE / P11 / COMPARISON | `IsUseForTransaksi` : `BIT` | `1` : `int` | `IsUseForTransaksi` → `INT` | `IsUseForTransaksi = 1` |
| IC-012432 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSTransactionSrcSecId:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransactionSrcSecId.sql#L48) | WHERE / P11 / COMPARISON | `IsUseForTransaksi` : `BIT` | `1` : `int` | `IsUseForTransaksi` → `INT` | `IsUseForTransaksi = 1` |
| IC-012434 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSTransactionSrcSecId:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransactionSrcSecId.sql#L56) | WHERE / P11 / COMPARISON | `IsUseForTransaksi` : `BIT` | `1` : `int` | `IsUseForTransaksi` → `INT` | `IsUseForTransaksi = 1` |
| IC-012436 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSTransactionSrcSecId:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransactionSrcSecId.sql#L63) | WHERE / P11 / COMPARISON | `IsUseForTransaksi` : `BIT` | `1` : `int` | `IsUseForTransaksi` → `INT` | `IsUseForTransaksi = 1` |
| IC-002429 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulateParamHargaModalNew:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateParamHargaModalNew.sql#L169) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.IsUseForTransaksi` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(sm.IsUseForTransaksi, 0)` |
| IC-002707 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:121](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L121) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `IsUseForTransaksi` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(IsUseForTransaksi, 0)` |
| IC-005907 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateParamHargaModalNew:226](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamHargaModalNew.sql#L226) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.IsUseForTransaksi` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(sm.IsUseForTransaksi, 0)` |
| IC-013337 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L70) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `IsUseForTransaksi` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(IsUseForTransaksi,0)` |

<a id="kolom-jumlahunit"></a>
## JumlahUnit

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L54). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002704 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:117](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L117) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `JumlahUnit` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(JumlahUnit, 0)` |
| IC-010802 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORIDetail:144](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L144) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `JumlahUnit` : `MONEY` | `1` : `int` | `1` → `MONEY` | `isnull(JumlahUnit, 1)` |
| IC-010934 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnesty:138](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L138) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `nullif(JumlahUnit, 0)` : `MONEY` | `1` : `int` | `1` → `MONEY` | `isnull(nullif(JumlahUnit, 0), 1)` |
| IC-010935 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnesty:138](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L138) | EXPRESSION / P09 / NULLIF_COMPARE | `JumlahUnit` : `MONEY` | `0` : `int` | `0` → `MONEY` | `nullif(JumlahUnit, 0)` |
| IC-012773 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateUploadParamHargaModalNew:301](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L301) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.JumlahUnit` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sm.JumlahUnit, 0)` |
| IC-012775 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateUploadParamHargaModalNew:308](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L308) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.JumlahUnit` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sm.JumlahUnit, 0)` |
| IC-012777 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateUploadParamHargaModalNew:315](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L315) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.JumlahUnit` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sm.JumlahUnit, 0)` |
| IC-012779 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateUploadParamHargaModalNew:322](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L322) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.JumlahUnit` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sm.JumlahUnit, 0)` |
| IC-013335 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L64) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `JumlahUnit` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(JumlahUnit, 0)` |

<a id="kolom-kelipatannominal"></a>
## KelipatanNominal

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013924 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidateInsertBuyBackOM:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateInsertBuyBackOM.sql#L41) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `KelipatanNominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(KelipatanNominal,0)` |
| IC-013947 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidatePerubahanOrderSuratBerharga:38](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidatePerubahanOrderSuratBerharga.sql#L38) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `KelipatanNominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(KelipatanNominal,0)` |

<a id="kolom-kuotabuyback"></a>
## KuotaBuyBack

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L40). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002720 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:215](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L215) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `KuotaBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(KuotaBuyBack, 0)` |
| IC-006898 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessInsertORIOrderCommit:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertORIOrderCommit.sql#L114) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `KuotaBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(KuotaBuyBack,0)` |
| IC-013090 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_InsertORI_Order_TT_Original:183](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORI_Order_TT_Original.sql#L183) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `KuotaBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(KuotaBuyBack, 0)` |
| IC-013439 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusORI_Order_TT:215](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L215) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.KuotaBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(smtm.KuotaBuyBack, 0)` |
| IC-013930 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidateInsertBuyBackOM:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateInsertBuyBackOM.sql#L43) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `KuotaBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(KuotaBuyBack,0)` |
| IC-013953 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidatePerubahanOrderSuratBerharga:40](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidatePerubahanOrderSuratBerharga.sql#L40) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `KuotaBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(KuotaBuyBack,0)` |

<a id="kolom-kuotacashback"></a>
## KuotaCashBack

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L41). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002721 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L216) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `KuotaCashBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(KuotaCashBack, 0)` |
| IC-006899 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessInsertORIOrderCommit:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertORIOrderCommit.sql#L115) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `KuotaCashBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(KuotaCashBack,0)` |
| IC-013091 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_InsertORI_Order_TT_Original:184](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORI_Order_TT_Original.sql#L184) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `KuotaCashBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(KuotaCashBack, 0)` |
| IC-013440 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusORI_Order_TT:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L216) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.KuotaCashBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(smtm.KuotaCashBack, 0)` |
| IC-013928 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidateInsertBuyBackOM:42](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateInsertBuyBackOM.sql#L42) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `KuotaCashBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(KuotaCashBack,0)` |
| IC-013951 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidatePerubahanOrderSuratBerharga:39](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidatePerubahanOrderSuratBerharga.sql#L39) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `KuotaCashBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(KuotaCashBack,0)` |

<a id="kolom-lastcoupondate"></a>
## LastCouponDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009480 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailUpdateCouponDate:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateCouponDate.sql#L56) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `LastCouponDate` : `DATETIME` | `#CouponDate.LastCouponDate` : `smalldatetime` | `LastCouponDate` → `smalldatetime` | `#CouponDate.LastCouponDate <- LastCouponDate` |
| IC-009490 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailUpdateCouponDate_Dummy:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateCouponDate_Dummy.sql#L44) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `LastCouponDate` : `DATETIME` | `#CouponDate.LastCouponDate` : `smalldatetime` | `LastCouponDate` → `smalldatetime` | `#CouponDate.LastCouponDate <- LastCouponDate` |
| IC-009486 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailUpdateCouponDate:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateCouponDate.sql#L125) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `LastCouponDate` : `DATETIME` | `@dNextCouponDate` : `smalldatetime` | `@dNextCouponDate` → `DATETIME` | `LastCouponDate = @dNextCouponDate` |
| IC-009496 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailUpdateCouponDate_Dummy:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateCouponDate_Dummy.sql#L113) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `LastCouponDate` : `DATETIME` | `@dNextCouponDate` : `smalldatetime` | `@dNextCouponDate` → `DATETIME` | `LastCouponDate = @dNextCouponDate` |

<a id="kolom-lastupdateby"></a>
## LastUpdateBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L26). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013761 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateSecurityMaster_TM:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateSecurityMaster_TM.sql#L106) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Trx.LastUpdateBy` : `INT` | `t.LastUpdateBy` : `varchar(20)` | `t.LastUpdateBy` → `INT` | `Trx.LastUpdateBy=t.LastUpdateBy` |

<a id="kolom-maturitydate"></a>
## MaturityDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 13 lokasi / 11 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009481 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailUpdateCouponDate:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateCouponDate.sql#L56) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `MaturityDate` : `DATETIME` | `#CouponDate.MaturityDate` : `smalldatetime` | `MaturityDate` → `smalldatetime` | `#CouponDate.MaturityDate <- MaturityDate` |
| IC-009491 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailUpdateCouponDate_Dummy:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateCouponDate_Dummy.sql#L44) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `MaturityDate` : `DATETIME` | `#CouponDate.MaturityDate` : `smalldatetime` | `MaturityDate` → `smalldatetime` | `#CouponDate.MaturityDate <- MaturityDate` |
| IC-001988 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Repo_ValidateField:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Repo_ValidateField.sql#L77) | WHERE / P04 / COMPARISON | `MaturityDate` : `DATETIME` | `@pcParam1` : `varchar(50)` | `@pcParam1` → `DATETIME` | `MaturityDate <= @pcParam1` |
| IC-002971 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SRBIProcessEmailNotifMaturityDateH1:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH1.sql#L63) | WHERE / P04 / COMPARISON | `a.MaturityDate` : `DATETIME` | `@cJatuhTempo` : `VARCHAR(10)` | `@cJatuhTempo` → `DATETIME` | `a.MaturityDate > @cJatuhTempo` |
| IC-002982 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SRBIProcessEmailNotifMaturityDateH3:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH3.sql#L63) | WHERE / P04 / COMPARISON | `a.MaturityDate` : `DATETIME` | `@cJatuhTempo` : `VARCHAR(10)` | `@cJatuhTempo` → `DATETIME` | `a.MaturityDate > @cJatuhTempo` |
| IC-002993 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SRBIProcessEmailNotifMaturityDateH7:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH7.sql#L63) | WHERE / P04 / COMPARISON | `a.MaturityDate` : `DATETIME` | `@cJatuhTempo` : `VARCHAR(10)` | `@cJatuhTempo` → `DATETIME` | `a.MaturityDate > @cJatuhTempo` |
| IC-003588 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCheckRecordingDate:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCheckRecordingDate.sql#L28) | WHERE / P04 / COMPARISON | `MaturityDate` : `DATETIME` | `@cSettlementDate` : `varchar(8)` | `@cSettlementDate` → `DATETIME` | `MaturityDate = @cSettlementDate` |
| IC-003589 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCheckRecordingDate2:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCheckRecordingDate2.sql#L37) | WHERE / P04 / COMPARISON | `MaturityDate` : `DATETIME` | `@cSettlementDate` : `varchar(8)` | `@cSettlementDate` → `DATETIME` | `MaturityDate = @cSettlementDate` |
| IC-004451 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSMaintenanceParamHargaModal:116](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L116) | WHERE / P04 / COMPARISON | `mst.MaturityDate` : `DATETIME` | `convert(varchar(10), getdate(), 112)` : `varchar(10)` | `convert(varchar(10), getdate(), 112)` → `DATETIME` | `mst.MaturityDate < convert(varchar(10), getdate(), 112)` |
| IC-004453 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSMaintenanceParamHargaModal:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L127) | WHERE / P04 / COMPARISON | `mst.MaturityDate` : `DATETIME` | `convert(varchar(10), getdate(), 112)` : `varchar(10)` | `convert(varchar(10), getdate(), 112)` → `DATETIME` | `mst.MaturityDate < convert(varchar(10), getdate(), 112)` |
| IC-005929 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateParamRepo:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamRepo.sql#L59) | WHERE / P04 / COMPARISON | `sm.MaturityDate` : `DATETIME` | `CONVERT(VARCHAR(8),GETDATE(),112)` : `VARCHAR(8)` | `CONVERT(VARCHAR(8),GETDATE(),112)` → `DATETIME` | `sm.MaturityDate > CONVERT(VARCHAR(8),GETDATE(),112)` |
| IC-009632 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptDailyPaymentCouponOrMaturity:301](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L301) | WHERE / P13 / COMPARISON | `sm.MaturityDate` : `DATETIME` | `(select current_working_date from dbo.control_table where office_id = 'TRS')` : `SMALLDATETIME` | `(select current_working_date from dbo.control_table where office_id = 'TRS')` → `DATETIME` | `sm.MaturityDate >= (select current_working_date from dbo.control_table where office_id = 'TRS')` |
| IC-009643 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptDailyPaymentCouponOrMaturity:388](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L388) | WHERE / P13 / COMPARISON | `sm.MaturityDate` : `DATETIME` | `(select current_working_date from dbo.control_table where office_id = 'TRS')` : `SMALLDATETIME` | `(select current_working_date from dbo.control_table where office_id = 'TRS')` → `DATETIME` | `sm.MaturityDate >= (select current_working_date from dbo.control_table where office_id = 'TRS')` |

<a id="kolom-maxnominalredemption"></a>
## MaxNominalRedemption

**Tipe dasar:** `DECIMAL(18,0)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L79). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001013 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.ORISBNISubmitEarlyRedemption:318](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNISubmitEarlyRedemption.sql#L318) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `smtm.MaxNominalRedemption` : `DECIMAL(18,0)` | `@maxRedemption` : `MONEY` | `smtm.MaxNominalRedemption` → `MONEY` | `@maxRedemption = smtm.MaxNominalRedemption` |
| IC-004226 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSInsertEarlyRedemption_TT:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertEarlyRedemption_TT.sql#L81) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `smtm.MaxNominalRedemption` : `DECIMAL(18,0)` | `100` : `int` | `100` → `DECIMAL` | `smtm.MaxNominalRedemption / 100` |
| IC-004227 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSInsertEarlyRedemption_TT:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertEarlyRedemption_TT.sql#L81) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `(smtm.MaxNominalRedemption / 100)` : `DECIMAL (combined result; facets unknown)` | `tmp.Outstanding` : `money` | `tmp.Outstanding` → `DECIMAL` | `tmp.Outstanding * (smtm.MaxNominalRedemption / 100)` |
| IC-012506 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSUpdateEarlyRedemption_TT:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateEarlyRedemption_TT.sql#L85) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `smtm.MaxNominalRedemption` : `DECIMAL(18,0)` | `100` : `int` | `100` → `DECIMAL` | `smtm.MaxNominalRedemption/100` |
| IC-012507 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSUpdateEarlyRedemption_TT:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateEarlyRedemption_TT.sql#L85) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `(smtm.MaxNominalRedemption/100)` : `DECIMAL (combined result; facets unknown)` | `tmp.Outstanding` : `money` | `tmp.Outstanding` → `DECIMAL` | `tmp.Outstanding * (smtm.MaxNominalRedemption/100)` |

<a id="kolom-minsafekeepingfee"></a>
## MinSafeKeepingFee

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-011498 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptSafeKeepingFeeCustodySummary:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSafeKeepingFeeCustodySummary.sql#L29) | CASE / P11 / COMPARISON | `smtm.MinSafeKeepingFee` : `DECIMAL(20,5)` | `skfcth.SKFProduct` : `MONEY` | `skfcth.SKFProduct` → `DECIMAL` | `skfcth.SKFProduct < smtm.MinSafeKeepingFee` |
| IC-011499 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptSafeKeepingFeeCustodySummary:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSafeKeepingFeeCustodySummary.sql#L30) | CASE / P11 / COMPARISON | `smtm.MinSafeKeepingFee` : `DECIMAL(20,5)` | `skfcth.SKFProduct` : `MONEY` | `skfcth.SKFProduct` → `DECIMAL` | `skfcth.SKFProduct > smtm.MinSafeKeepingFee` |

<a id="kolom-minimumbuyback"></a>
## MinimumBuyBack

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L42). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013931 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidateInsertBuyBackOM:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateInsertBuyBackOM.sql#L43) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `MinimumBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(MinimumBuyBack,0)` |
| IC-013954 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidatePerubahanOrderSuratBerharga:40](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidatePerubahanOrderSuratBerharga.sql#L40) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `MinimumBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(MinimumBuyBack,0)` |

<a id="kolom-minimumcashback"></a>
## MinimumCashBack

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L43). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013929 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidateInsertBuyBackOM:42](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateInsertBuyBackOM.sql#L42) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `MinimumCashBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(MinimumCashBack,0)` |
| IC-013952 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidatePerubahanOrderSuratBerharga:39](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidatePerubahanOrderSuratBerharga.sql#L39) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `MinimumCashBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(MinimumCashBack,0)` |

<a id="kolom-nextcoupondate"></a>
## NextCouponDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009482 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailUpdateCouponDate:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateCouponDate.sql#L56) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `NextCouponDate` : `DATETIME` | `#CouponDate.NextCouponDate` : `smalldatetime` | `NextCouponDate` → `smalldatetime` | `#CouponDate.NextCouponDate <- NextCouponDate` |
| IC-009492 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailUpdateCouponDate_Dummy:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateCouponDate_Dummy.sql#L44) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `NextCouponDate` : `DATETIME` | `#CouponDate.NextCouponDate` : `smalldatetime` | `NextCouponDate` → `smalldatetime` | `#CouponDate.NextCouponDate <- NextCouponDate` |
| IC-009483 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailUpdateCouponDate:60](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateCouponDate.sql#L60) | WHERE / P13 / COMPARISON | `NextCouponDate` : `DATETIME` | `@dcurrent_working_date` : `smalldatetime` | `@dcurrent_working_date` → `DATETIME` | `NextCouponDate <= @dcurrent_working_date` |
| IC-009485 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailUpdateCouponDate:119](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateCouponDate.sql#L119) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `NextCouponDate` : `DATETIME` | `CASE WHEN @dPeriodNextCoupon IS NULL THEN dateadd(mm, @nDiff, @dNextCouponDate) ELSE @dPeriodNextCoupon END` : `SMALLDATETIME (combined result; facets unknown)` | `CASE WHEN @dPeriodNextCoupon IS NULL THEN dateadd(mm, @nDiff, @dNextCouponDate) ELSE @dPeriodNextCoupon END` → `DATETIME` | `NextCouponDate = CASE WHEN @dPeriodNextCoupon IS NULL THEN dateadd(mm, @nDiff, @dNextCouponDate) ELSE @dPeriodNextCoupon END` |
| IC-009493 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailUpdateCouponDate_Dummy:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateCouponDate_Dummy.sql#L48) | WHERE / P13 / COMPARISON | `NextCouponDate` : `DATETIME` | `@dcurrent_working_date` : `smalldatetime` | `@dcurrent_working_date` → `DATETIME` | `NextCouponDate <= @dcurrent_working_date` |
| IC-009495 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailUpdateCouponDate_Dummy:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateCouponDate_Dummy.sql#L107) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `NextCouponDate` : `DATETIME` | `CASE WHEN @dPeriodNextCoupon IS NULL THEN dateadd(mm, @nDiff, @dNextCouponDate) ELSE @dPeriodNextCoupon END` : `SMALLDATETIME (combined result; facets unknown)` | `CASE WHEN @dPeriodNextCoupon IS NULL THEN dateadd(mm, @nDiff, @dNextCouponDate) ELSE @dPeriodNextCoupon END` → `DATETIME` | `NextCouponDate = CASE WHEN @dPeriodNextCoupon IS NULL THEN dateadd(mm, @nDiff, @dNextCouponDate) ELSE @dPeriodNextCoupon END` |

<a id="kolom-nominalmaksimum"></a>
## NominalMaksimum

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013922 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidateInsertBuyBackOM:40](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateInsertBuyBackOM.sql#L40) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NominalMaksimum` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(NominalMaksimum,0)` |
| IC-013945 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidatePerubahanOrderSuratBerharga:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidatePerubahanOrderSuratBerharga.sql#L37) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NominalMaksimum` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(NominalMaksimum,0)` |

<a id="kolom-nominalminimum"></a>
## NominalMinimum

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013923 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidateInsertBuyBackOM:40](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateInsertBuyBackOM.sql#L40) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NominalMinimum` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(NominalMinimum,0)` |
| IC-013946 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidatePerubahanOrderSuratBerharga:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidatePerubahanOrderSuratBerharga.sql#L37) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NominalMinimum` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(NominalMinimum,0)` |

<a id="kolom-onlywni"></a>
## OnlyWNI

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L91). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004129 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSGetMenuSettings:275](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetMenuSettings.sql#L275) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `OnlyWNI` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(OnlyWNI, 0)` |
| IC-004319 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSInsertSecurityTransaction_TT_Original:833](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L833) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `OnlyWNI` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(OnlyWNI, 0)` |

<a id="kolom-persentasecashback"></a>
## PersentaseCashBack

**Tipe dasar:** `DECIMAL(6,3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L46). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013998 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_updateTopUpORI_Order_Temp:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateTopUpORI_Order_Temp.sql#L80) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `(mst.NominalCashBackValue + Trx.NominalCashBack) * sm.PersentaseCashBack / 100` : `DECIMAL (combined result; facets unknown)` | `mst.NominalCashBack` : `MONEY` | `(mst.NominalCashBackValue + Trx.NominalCashBack) * sm.PersentaseCashBack / 100` → `MONEY` | `mst.NominalCashBack = (mst.NominalCashBackValue + Trx.NominalCashBack) * sm.PersentaseCashBack / 100` |
| IC-006544 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessAllotment:95](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessAllotment.sql#L95) | DML_CONDITION / P05 / BINARY_Multiply | `mst.PersentaseCashBack` : `DECIMAL(6,3)` | `(temp.OrderNominalBaru - (CASE WHEN ISNULL(oott.NominalBuyBack, 0) <= temp.OrderNominalBaru THEN oott.NominalBuyBack ELSE temp.OrderNominalBaru END ))` : `MONEY (combined result; facets unknown)` | `(temp.OrderNominalBaru - (CASE WHEN ISNULL(oott.NominalBuyBack, 0) <= temp.OrderNominalBaru THEN oott.NominalBuyBack ELSE temp.OrderNominalBaru END ))` → `DECIMAL` | `(temp.OrderNominalBaru - (CASE WHEN ISNULL(oott.NominalBuyBack, 0) <= temp.OrderNominalBaru THEN oott.NominalBuyBack ELSE temp.OrderNominalBaru END )) * mst.PersentaseCashBack` |
| IC-012240 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSplitOrder:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSplitOrder.sql#L44) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `PersentaseCashBack` : `DECIMAL(6,3)` | `0` : `int` | `0` → `DECIMAL(6,3)` | `ISNULL(PersentaseCashBack,0)` |
| IC-013089 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_InsertORI_Order_TT_Original:177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORI_Order_TT_Original.sql#L177) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `PersentaseCashBack` : `DECIMAL(6,3)` | `0` : `int` | `0` → `DECIMAL(6,3)` | `ISNULL(PersentaseCashBack, 0)` |
| IC-013999 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_updateTopUpORI_Order_Temp:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateTopUpORI_Order_Temp.sql#L80) | DML_CONDITION / P05 / BINARY_Divide | `(mst.NominalCashBackValue + Trx.NominalCashBack) * sm.PersentaseCashBack` : `DECIMAL (combined result; facets unknown)` | `100` : `int` | `100` → `DECIMAL` | `(mst.NominalCashBackValue + Trx.NominalCashBack) * sm.PersentaseCashBack / 100` |
| IC-014000 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_updateTopUpORI_Order_Temp:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateTopUpORI_Order_Temp.sql#L80) | DML_CONDITION / P05 / BINARY_Multiply | `sm.PersentaseCashBack` : `DECIMAL(6,3)` | `(mst.NominalCashBackValue + Trx.NominalCashBack)` : `MONEY (combined result; facets unknown)` | `(mst.NominalCashBackValue + Trx.NominalCashBack)` → `DECIMAL` | `(mst.NominalCashBackValue + Trx.NominalCashBack) * sm.PersentaseCashBack` |

<a id="kolom-safekeepingfeetarif"></a>
## SafeKeepingFeeTarif

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 14 lokasi / 12 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006780 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessBuyBackOM:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBuyBackOM.sql#L169) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `sm.SafeKeepingFeeTarif` : `DECIMAL(20,5)` | `dt.SafeKeepingFee` : `MONEY` | `sm.SafeKeepingFeeTarif` → `MONEY` | `dt.SafeKeepingFee = sm.SafeKeepingFeeTarif` |
| IC-008723 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailCalculateCoupon:461](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateCoupon.sql#L461) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `SafeKeepingFeeTarif` : `DECIMAL(20,5)` | `@nCalculatedSafeKeepingFee` : `MONEY` | `SafeKeepingFeeTarif` → `MONEY` | `@nCalculatedSafeKeepingFee = SafeKeepingFeeTarif` |
| IC-008724 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailCalculateCoupon:480](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateCoupon.sql#L480) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `SafeKeepingFeeTarif` : `DECIMAL(20,5)` | `@nCalculatedSafeKeepingFee` : `MONEY` | `SafeKeepingFeeTarif` → `MONEY` | `@nCalculatedSafeKeepingFee = SafeKeepingFeeTarif` |
| IC-008725 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailCalculateCoupon:492](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateCoupon.sql#L492) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `SafeKeepingFeeTarif` : `DECIMAL(20,5)` | `@nCalculatedSafeKeepingFee` : `MONEY` | `SafeKeepingFeeTarif` → `MONEY` | `@nCalculatedSafeKeepingFee = SafeKeepingFeeTarif` |
| IC-002702 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_trs_ListSecurityMaster_TM:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L80) | CASE / P11 / COMPARISON | `SafeKeepingFeeTarif` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `SafeKeepingFeeTarif = 0` |
| IC-010622 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPajakKupon:100](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPajakKupon.sql#L100) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `smtm.SafeKeepingFeeTarif` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `smtm.SafeKeepingFeeTarif / 100` |
| IC-010710 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORI:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORI.sql#L106) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `SafeKeepingFeeTarif` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `SafeKeepingFeeTarif / 100` |
| IC-010796 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORIDetail:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L127) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `SafeKeepingFeeTarif` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `SafeKeepingFeeTarif / 100` |
| IC-010927 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnesty:121](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L121) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `SafeKeepingFeeTarif` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `SafeKeepingFeeTarif / 100` |
| IC-011041 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnestyEOM:78](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyEOM.sql#L78) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `SafeKeepingFeeTarif` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `SafeKeepingFeeTarif / 100` |
| IC-011127 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L106) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `SafeKeepingFeeTarif` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `SafeKeepingFeeTarif / 100` |
| IC-011253 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnestySPT:82](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestySPT.sql#L82) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `SafeKeepingFeeTarif` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `SafeKeepingFeeTarif / 100` |
| IC-011341 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnesty_Backdated:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty_Backdated.sql#L107) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `SafeKeepingFeeTarif` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `SafeKeepingFeeTarif / 100` |
| IC-013334 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_ListSecurityMaster_TM:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L59) | CASE / P11 / COMPARISON | `SafeKeepingFeeTarif` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `SafeKeepingFeeTarif = 0` |

<a id="kolom-safekeepingfeetaxtarif"></a>
## SafeKeepingFeeTaxTarif

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L19). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 10 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006781 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessBuyBackOM:170](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBuyBackOM.sql#L170) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `sm.SafeKeepingFeeTaxTarif` : `DECIMAL(20,5)` | `dt.TaxOnSafeKeepingFee` : `MONEY` | `sm.SafeKeepingFeeTaxTarif` → `MONEY` | `dt.TaxOnSafeKeepingFee = sm.SafeKeepingFeeTaxTarif` |
| IC-006965 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessKuponUpload:281](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessKuponUpload.sql#L281) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sm.SafeKeepingFeeTaxTarif` : `DECIMAL(20,5)` | `#tFromDatabase.TaxSafeKepping` : `MONEY` | `sm.SafeKeepingFeeTaxTarif` → `MONEY` | `#tFromDatabase.TaxSafeKepping <- sm.SafeKeepingFeeTaxTarif` |
| IC-008378 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportSKFBankBeli:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeli.sql#L64) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `sm.SafeKeepingFeeTaxTarif/100.0` : `DECIMAL (combined result; facets unknown)` | `1` : `int` | `1` → `DECIMAL` | `1 + sm.SafeKeepingFeeTaxTarif/100.0` |
| IC-008379 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportSKFBankBeli:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeli.sql#L64) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `(1 + sm.SafeKeepingFeeTaxTarif/100.0)` : `DECIMAL (combined result; facets unknown)` | `st.SafeKeepingFeeAfterTax` : `MONEY` | `st.SafeKeepingFeeAfterTax` → `DECIMAL` | `st.SafeKeepingFeeAfterTax / (1 + sm.SafeKeepingFeeTaxTarif/100.0)` |
| IC-008380 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportSKFBankBeli:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeli.sql#L65) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `sm.SafeKeepingFeeTaxTarif/100.0` : `DECIMAL (combined result; facets unknown)` | `1` : `int` | `1` → `DECIMAL` | `1 + sm.SafeKeepingFeeTaxTarif/100.0` |
| IC-008381 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportSKFBankBeli:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeli.sql#L65) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `(1 + sm.SafeKeepingFeeTaxTarif/100.0)` : `DECIMAL (combined result; facets unknown)` | `st.SafeKeepingFeeAfterTax` : `MONEY` | `st.SafeKeepingFeeAfterTax` → `DECIMAL` | `st.SafeKeepingFeeAfterTax / (1 + sm.SafeKeepingFeeTaxTarif/100.0)` |
| IC-008398 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportSKFBankBeliBKON:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeliBKON.sql#L107) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `sm.SafeKeepingFeeTaxTarif/100.0` : `DECIMAL (combined result; facets unknown)` | `1` : `int` | `1` → `DECIMAL` | `1 + sm.SafeKeepingFeeTaxTarif/100.0` |
| IC-008399 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportSKFBankBeliBKON:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeliBKON.sql#L107) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `(1 + sm.SafeKeepingFeeTaxTarif/100.0)` : `DECIMAL (combined result; facets unknown)` | `tl.SafeKeepingFeeAfterTax` : `MONEY` | `tl.SafeKeepingFeeAfterTax` → `DECIMAL` | `tl.SafeKeepingFeeAfterTax / (1 + sm.SafeKeepingFeeTaxTarif/100.0)` |
| IC-008400 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportSKFBankBeliBKON:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeliBKON.sql#L108) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `sm.SafeKeepingFeeTaxTarif/100.0` : `DECIMAL (combined result; facets unknown)` | `1` : `int` | `1` → `DECIMAL` | `1 + sm.SafeKeepingFeeTaxTarif/100.0` |
| IC-008401 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportSKFBankBeliBKON:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeliBKON.sql#L108) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `(1 + sm.SafeKeepingFeeTaxTarif/100.0)` : `DECIMAL (combined result; facets unknown)` | `tl.SafeKeepingFeeAfterTax` : `MONEY` | `tl.SafeKeepingFeeAfterTax` → `DECIMAL` | `tl.SafeKeepingFeeAfterTax / (1 + sm.SafeKeepingFeeTaxTarif/100.0)` |

<a id="kolom-safekeepingfeetype"></a>
## SafeKeepingFeeType

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L47). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008867 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailCalculateFee:983](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L983) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `SafeKeepingFeeType` : `INT` | `@nSecuritySafeKeepingFeeType` : `smallint` | `SafeKeepingFeeType` → `smallint` | `@nSecuritySafeKeepingFeeType = SafeKeepingFeeType` |

<a id="kolom-secccy"></a>
## SecCcy

**Tipe dasar:** `VARCHAR(3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 18 lokasi / 7 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT.

**Pola dan arah review:**

- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001900 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.RepoReportNoticePayment:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoReportNoticePayment.sql#L22) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `sm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetterbilang.@pcCCY` : `[nvarchar](4000)` | `sm.SecCcy` → `[nvarchar](4000)` | `dbo.fnGetTerbilang(Repo2ndLeg, sm.SecCcy)` |
| IC-001902 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.RepoReportNoticePayment:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoReportNoticePayment.sql#L23) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `sm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetterbilangenglish.@pcCCY` : `[nvarchar](4000)` | `sm.SecCcy` → `[nvarchar](4000)` | `dbo.fnGetTerbilangEnglish(Repo2ndLeg, sm.SecCcy)` |
| IC-001942 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.RepoSubmitTransaction:377](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSubmitTransaction.sql#L377) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `sm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetmxworkingdate.@pcCcy` : `char(3)` | `sm.SecCcy` → `char(3)` | `dbo.fnGetMxWorkingDate('F',tmp.TrxDate,1,sm.SecCcy)` |
| IC-003015 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.SecurityLelangTrxSubmit:306](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityLelangTrxSubmit.sql#L306) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `sm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetmxworkingdate.@pcCcy` : `char(3)` | `sm.SecCcy` → `char(3)` | `dbo.fnGetMxWorkingDate('F',tmp.TrxDate,1,sm.SecCcy)` |
| IC-003030 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.SecurityLelangTrxSubmit:539](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityLelangTrxSubmit.sql#L539) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `sm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetmxworkingdate.@pcCcy` : `char(3)` | `sm.SecCcy` → `char(3)` | `dbo.fnGetMxWorkingDate('F',CONVERT(VARCHAR, tmp.TrxDate, 23) + ' ' + CONVERT(VARCHAR, GETDATE(), 14),1,sm.SecCcy)` |
| IC-005417 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSPopulateFOPTransferAsset:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateFOPTransferAsset.sql#L69) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `sm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetmxworkingdate.@pcCcy` : `char(3)` | `sm.SecCcy` → `char(3)` | `dbo.fnGetMxWorkingDate('B',sm.NextCouponDate,4,sm.SecCcy)` |
| IC-005418 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSPopulateFOPTransferAsset:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateFOPTransferAsset.sql#L70) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `sm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetmxworkingdate.@pcCcy` : `char(3)` | `sm.SecCcy` → `char(3)` | `dbo.fnGetMxWorkingDate('B',sm.NextCouponDate,2,sm.SecCcy)` |
| IC-005420 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSPopulateFOPTransferAsset:73](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateFOPTransferAsset.sql#L73) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `sm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetmxworkingdate.@pcCcy` : `char(3)` | `sm.SecCcy` → `char(3)` | `dbo.fnGetMxWorkingDate('B',sm.NextCouponDate,3,sm.SecCcy)` |
| IC-005421 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSPopulateFOPTransferAsset:74](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateFOPTransferAsset.sql#L74) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `sm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetmxworkingdate.@pcCcy` : `char(3)` | `sm.SecCcy` → `char(3)` | `dbo.fnGetMxWorkingDate('B',sm.NextCouponDate,1,sm.SecCcy)` |
| IC-009511 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSRpt2ndBondTrx:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRpt2ndBondTrx.sql#L106) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `sm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetterbilang.@pcCCY` : `[nvarchar](4000)` | `sm.SecCcy` → `[nvarchar](4000)` | `dbo.fnGetTerbilang(cast(st.TotalProceed as varchar), sm.SecCcy)` |
| IC-012732 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSValidateUploadParamHargaModal:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModal.sql#L107) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `sm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetmxworkingdate.@pcCcy` : `char(3)` | `sm.SecCcy` → `char(3)` | `dbo.fnGetMxWorkingDate('B',sm.NextCouponDate,4,sm.SecCcy)` |
| IC-012733 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSValidateUploadParamHargaModal:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModal.sql#L108) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `sm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetmxworkingdate.@pcCcy` : `char(3)` | `sm.SecCcy` → `char(3)` | `dbo.fnGetMxWorkingDate('B',sm.NextCouponDate,2,sm.SecCcy)` |
| IC-012735 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSValidateUploadParamHargaModal:112](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModal.sql#L112) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `sm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetmxworkingdate.@pcCcy` : `char(3)` | `sm.SecCcy` → `char(3)` | `dbo.fnGetMxWorkingDate('B',sm.NextCouponDate,3,sm.SecCcy)` |
| IC-012736 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSValidateUploadParamHargaModal:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModal.sql#L113) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `sm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetmxworkingdate.@pcCcy` : `char(3)` | `sm.SecCcy` → `char(3)` | `dbo.fnGetMxWorkingDate('B',sm.NextCouponDate,1,sm.SecCcy)` |
| IC-012760 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSValidateUploadParamHargaModalNew:229](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L229) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `sm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetmxworkingdatevalasincusd.@pcCcy` : `char(3)` | `sm.SecCcy` → `char(3)` | `dbo.fnGetMxWorkingDateValasIncUSD('B',sm.NextCouponDate,4,sm.SecCcy)` |
| IC-012761 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSValidateUploadParamHargaModalNew:230](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L230) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `sm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetmxworkingdatevalasincusd.@pcCcy` : `char(3)` | `sm.SecCcy` → `char(3)` | `dbo.fnGetMxWorkingDateValasIncUSD('B',sm.NextCouponDate,2,sm.SecCcy)` |
| IC-012763 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSValidateUploadParamHargaModalNew:238](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L238) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `sm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetmxworkingdatevalasincusd.@pcCcy` : `char(3)` | `sm.SecCcy` → `char(3)` | `dbo.fnGetMxWorkingDateValasIncUSD('B',sm.NextCouponDate,3,sm.SecCcy)` |
| IC-012764 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSValidateUploadParamHargaModalNew:239](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L239) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `sm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetmxworkingdatevalasincusd.@pcCcy` : `char(3)` | `sm.SecCcy` → `char(3)` | `dbo.fnGetMxWorkingDateValasIncUSD('B',sm.NextCouponDate,1,sm.SecCcy)` |

<a id="kolom-secdescr"></a>
## SecDescr

**Tipe dasar:** `VARCHAR(70)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005554 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateKuponUploadAuthorization:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUploadAuthorization.sql#L30) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `'Kupon ' + smtm.SecDescr + ' '` : `VARCHAR (combined result; facets unknown)` | `Convert(char(3),mtt.InsertedDate )` : `char(3)` | `Convert(char(3),mtt.InsertedDate )` → `VARCHAR` | `'Kupon ' + smtm.SecDescr + ' ' + Convert(char(3),mtt.InsertedDate )` |
| IC-005555 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateKuponUploadAuthorization:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUploadAuthorization.sql#L30) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `'Kupon ' + smtm.SecDescr + ' ' + Convert(char(3),mtt.InsertedDate ) + ' '` : `VARCHAR (combined result; facets unknown)` | `Convert(char(4),Year(mtt.InsertedDate))` : `char(4)` | `Convert(char(4),Year(mtt.InsertedDate))` → `VARCHAR` | `'Kupon ' + smtm.SecDescr + ' ' + Convert(char(3),mtt.InsertedDate ) + ' ' + Convert(char(4),Year(mtt.InsertedDate))` |
| IC-009559 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptBatchTransaction:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptBatchTransaction.sql#L47) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `'Kupon ' + smtm.SecDescr + ' '` : `VARCHAR (combined result; facets unknown)` | `convert(char(3),tmct.CouponDate)` : `char(3)` | `convert(char(3),tmct.CouponDate)` → `VARCHAR` | `'Kupon ' + smtm.SecDescr + ' ' + convert(char(3),tmct.CouponDate)` |
| IC-009560 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptBatchTransaction:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptBatchTransaction.sql#L47) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `'Kupon ' + smtm.SecDescr + ' ' + convert(char(3),tmct.CouponDate) + ' '` : `VARCHAR (combined result; facets unknown)` | `convert(char(4),year(tmct.InsertedDate))` : `char(4)` | `convert(char(4),year(tmct.InsertedDate))` → `VARCHAR` | `'Kupon ' + smtm.SecDescr + ' ' + convert(char(3),tmct.CouponDate) + ' ' + convert(char(4),year(tmct.InsertedDate))` |

<a id="kolom-secid"></a>
## SecId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L2). **Existing leading index:** IU_SecurityMaster_TM&#124;PK__SecurityMaster_T__7E37BEF6.

**Frekuensi source:** 36 lokasi / 28 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000148 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.CustodyReportMaintainAcquisition:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportMaintainAcquisition.sql#L104) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `SecId` : `INT` | `#tmp_TreasuryCustomerBalance_v.SecId` : `VARCHAR(30)` | `SecId` → `VARCHAR(30)` | `#tmp_TreasuryCustomerBalance_v.SecId <- SecId` |
| IC-000152 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.CustodyReportMaintainAcquisition:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportMaintainAcquisition.sql#L110) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `SecId` : `INT` | `#tmp_SecurityMaster_TM.SecId` : `VARCHAR(30)` | `SecId` → `VARCHAR(30)` | `#tmp_SecurityMaster_TM.SecId <- SecId` |
| IC-002128 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.Rvmp_TRSGenerateDumpOBL084:99](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGenerateDumpOBL084.sql#L99) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `SecId` : `INT` | `@nSecId` : `bigint` | `SecId` → `bigint` | `@nSecId = SecId` |
| IC-002176 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.Rvmp_TRSKuponCopyFromCustody:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponCopyFromCustody.sql#L85) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `SecId` : `INT` | `@nSecId` : `bigint` | `SecId` → `bigint` | `@nSecId = SecId` |
| IC-002184 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.Rvmp_TRSKuponProcess:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L86) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `sm.SecId` : `INT` | `@nSecId` : `BIGINT` | `sm.SecId` → `BIGINT` | `@nSecId = sm.SecId` |
| IC-002235 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.Rvmp_TRSMaturityDumpOBL058:68](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityDumpOBL058.sql#L68) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `SecId` : `INT` | `@nSecId` : `BIGINT` | `SecId` → `BIGINT` | `@nSecId = SecId` |
| IC-002254 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.Rvmp_TRSMaturityProcess:186](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityProcess.sql#L186) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `sm.SecId` : `INT` | `@nSecId` : `BIGINT` | `sm.SecId` → `BIGINT` | `@nSecId = sm.SecId` |
| IC-005977 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateReport008:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateReport008.sql#L180) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `SecId` : `INT` | `@nSecId` : `bigint` | `SecId` → `bigint` | `@nSecId = SecId` |
| IC-006957 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessKuponUpload:122](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessKuponUpload.sql#L122) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `SecId` : `INT` | `@nSecId` : `BIGINT` | `SecId` → `BIGINT` | `@nSecId = SecId` |
| IC-006981 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessMaturityUpload:101](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessMaturityUpload.sql#L101) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `SecId` : `INT` | `@nSecId` : `BIGINT` | `SecId` → `BIGINT` | `@nSecId = SecId` |
| IC-009140 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailSrcSecAccNo:38](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSrcSecAccNo.sql#L38) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `SecId` : `INT` | `@nSecId` : `bigint` | `SecId` → `bigint` | `@nSecId = SecId` |
| IC-009611 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptDailyPaymentCouponOrMaturity:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L77) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sm.SecId` : `INT` | `#tmpResult.SecId` : `bigint` | `sm.SecId` → `bigint` | `#tmpResult.SecId <- sm.SecId` |
| IC-009616 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptDailyPaymentCouponOrMaturity:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L141) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sm.SecId` : `INT` | `#tmpResult.SecId` : `bigint` | `sm.SecId` → `bigint` | `#tmpResult.SecId <- sm.SecId` |
| IC-009626 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptDailyPaymentCouponOrMaturity:222](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L222) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sm.SecId` : `INT` | `#tmpResult.SecId` : `bigint` | `sm.SecId` → `bigint` | `#tmpResult.SecId <- sm.SecId` |
| IC-009637 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptDailyPaymentCouponOrMaturity:309](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L309) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sm.SecId` : `INT` | `#tmpResult.SecId` : `bigint` | `sm.SecId` → `bigint` | `#tmpResult.SecId <- sm.SecId` |
| IC-010932 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptPayCouponORITaxAmnesty:135](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L135) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `SecId` : `INT` | `@nSecId` : `bigint` | `SecId` → `bigint` | `@nSecId = SecId` |
| IC-011132 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:119](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L119) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `SecId` : `INT` | `@nSecId` : `BIGINT` | `SecId` → `BIGINT` | `@nSecId = SecId` |
| IC-011505 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptSettleKuponParamProduct:9](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettleKuponParamProduct.sql#L9) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `SecId` : `INT` | `@tmp_data.SecId` : `bigint` | `SecId` → `bigint` | `@tmp_data.SecId <- SecId` |
| IC-011530 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptSettlementKupon:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettlementKupon.sql#L46) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[SecId]` : `INT` | `@nSecId` : `bigint` | `[SecId]` → `bigint` | `@nSecId = [SecId]` |
| IC-011661 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSSaveBalanceHistory:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveBalanceHistory.sql#L29) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tcb.SecId` : `INT` | `retail.dbo.treasurycustomerbalance_th.SecId` : `BIGINT` | `tcb.SecId` → `BIGINT` | `retail.dbo.treasurycustomerbalance_th.SecId <- tcb.SecId` |
| IC-001029 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OrivaGetParameter:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OrivaGetParameter.sql#L16) | SET_OUTPUT / P07 / UNION_OUTPUT | `right output #1` : `INT` | `left output #1` : `BIGINT` | `right output #1` → `BIGINT` | `SELECT CAST(0 AS BIGINT) AS SecId, 'SELECT ALL' AS SecurityNo UNION ALL SELECT SecId, SecurityNo FROM dbo.SecurityMaster_TM WHERE 1 = 1 AND SecCcy = 'USD' AND ISNULL(FlagPerdana, CAST(0 AS BIT)) = 1` |
| IC-001038 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OrivaTransactionReport:76](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OrivaTransactionReport.sql#L76) | JOIN_ON / P11 / COMPARISON | `sec.SecId` : `INT` | `tr.sec_id` : `BIGINT` | `sec.SecId` → `BIGINT` | `sec.SecId = tr.sec_id` |
| IC-001041 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OrivaTransactionReport:93](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OrivaTransactionReport.sql#L93) | WHERE / P11 / COMPARISON | `sec.SecId` : `INT` | `@pcSecurityNo` : `BIGINT` | `sec.SecId` → `BIGINT` | `sec.SecId= @pcSecurityNo` |
| IC-001781 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.RPAJurnalSettlementOriva:83](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RPAJurnalSettlementOriva.sql#L83) | JOIN_ON / P11 / COMPARISON | `sec.SecId` : `INT` | `tr.sec_id` : `BIGINT` | `sec.SecId` → `BIGINT` | `sec.SecId = tr.sec_id` |
| IC-001820 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.RepoGetXMLToMurex:312](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoGetXMLToMurex.sql#L312) | WHERE / P11 / COMPARISON | `SecId` : `INT` | `@nSecId` : `BIGINT` | `SecId` → `BIGINT` | `SecId = @nSecId` |
| IC-002223 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSKuponProcess:462](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L462) | JOIN_ON / P11 / COMPARISON | `tcb.SecId` : `INT` | `ps.sec_id` : `BIGINT` | `tcb.SecId` → `BIGINT` | `tcb.SecId = ps.sec_id` |
| IC-004243 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:136](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L136) | JOIN_ON / P11 / COMPARISON | `m.[SecId]` : `INT` | `scc.[SecId]` : `BIGINT` | `m.[SecId]` → `BIGINT` | `scc.[SecId] = m.[SecId]` |
| IC-004461 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSMaintenanceParamHargaModal:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L165) | JOIN_ON / P11 / COMPARISON | `mst.SecId` : `INT` | `excl.SecId` : `BIGINT` | `mst.SecId` → `BIGINT` | `mst.SecId = excl.SecId` |
| IC-005456 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateGagalBayarOri:266](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L266) | JOIN_ON / P11 / COMPARISON | `tm.SecId` : `INT` | `tmp.SecId` : `BIGINT` | `tm.SecId` → `BIGINT` | `tmp.SecId = tm.SecId` |
| IC-005463 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateGagalBayarOri:374](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L374) | JOIN_ON / P11 / COMPARISON | `tm.SecId` : `INT` | `tmp.SecId` : `BIGINT` | `tm.SecId` → `BIGINT` | `tmp.SecId = tm.SecId` |
| IC-005473 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateGagalBayarOri:500](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L500) | JOIN_ON / P11 / COMPARISON | `tm.SecId` : `INT` | `tmp.SecId` : `BIGINT` | `tm.SecId` → `BIGINT` | `tmp.SecId = tm.SecId` |
| IC-006013 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSecurityParameter:10](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSecurityParameter.sql#L10) | JOIN_ON / P11 / COMPARISON | `tm.SecId` : `INT` | `ex.SecId` : `BIGINT` | `tm.SecId` → `BIGINT` | `tm.SecId = ex.SecId` |
| IC-010451 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptOutstandingNasabahBySettDate:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingNasabahBySettDate.sql#L61) | JOIN_ON / P11 / COMPARISON | `sm.SecId` : `INT` | `tcb.SecId` : `BIGINT` | `sm.SecId` → `BIGINT` | `sm.SecId = tcb.SecId` |
| IC-002715 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_trs_ListSecurityMaster_TM:188](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L188) | WHERE / P01 / COMPARISON | `tm.SecId` : `INT` | `g.process_key` : `VARCHAR(100)` | `g.process_key` → `INT` | `g.process_key = tm.SecId` |
| IC-008463 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportSettlementBankBeli:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSettlementBankBeli.sql#L53) | WHERE / P04 / COMPARISON | `smtm.SecId` : `INT` | `@pcSecurityNo` : `varchar(20)` | `@pcSecurityNo` → `INT` | `smtm.SecId = @pcSecurityNo` |
| IC-013762 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_UpdateSecurityMaster_TM:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateSecurityMaster_TM.sql#L113) | JOIN_ON / P01 / COMPARISON | `Trx.SecId` : `INT` | `t.SecId` : `varchar(10)` | `t.SecId` → `INT` | `Trx.SecId=t.SecId` |

<a id="kolom-sectype"></a>
## SecType

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 91 lokasi / 40 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012767 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSValidateUploadParamHargaModalNew:249](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L249) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `sm.SecType` : `SMALLINT` | `td.SecurityType` : `INT` | `sm.SecType` → `INT` | `td.SecurityType = sm.SecType` |
| IC-000195 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L113) | CASE / P11 / COMPARISON | `sttm.SecType` : `SMALLINT` | `3` : `int` | `sttm.SecType` → `INT` | `sttm.SecType != 3` |
| IC-000198 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L114) | CASE / P11 / COMPARISON | `sttm.SecType` : `SMALLINT` | `3` : `int` | `sttm.SecType` → `INT` | `sttm.SecType = 3` |
| IC-000209 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:150](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L150) | CASE / P11 / COMPARISON | `SecType` : `SMALLINT` | `3` : `int` | `SecType` → `INT` | `SecType = 3` |
| IC-002426 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateParamHargaModalNew:159](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateParamHargaModalNew.sql#L159) | WHERE / P11 / COMPARISON | `sm.SecType` : `SMALLINT` | `@pnSecurityType` : `INT` | `sm.SecType` → `INT` | `sm.SecType = @pnSecurityType` |
| IC-002610 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSSaveParamHargaModalNew_Normal:455](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSaveParamHargaModalNew_Normal.sql#L455) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-002615 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSSaveParamHargaModalNew_Normal:462](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSaveParamHargaModalNew_Normal.sql#L462) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-002624 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSSaveParamHargaModalNew_Normal:556](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSaveParamHargaModalNew_Normal.sql#L556) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-002632 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSSaveParamHargaModalNew_Normal:578](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSaveParamHargaModalNew_Normal.sql#L578) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-003208 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSAmendSecurityTransaction_TT:630](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L630) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `sm.SecType` : `SMALLINT` | `1` : `int` | `sm.SecType` → `INT` | `WHEN 1 THEN 'BOND'` |
| IC-003209 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSAmendSecurityTransaction_TT:631](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L631) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `sm.SecType` : `SMALLINT` | `3` : `int` | `sm.SecType` → `INT` | `WHEN 3 THEN 'SRBI'` |
| IC-003324 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApproveHapusTransaksiSuratBerharga:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveHapusTransaksiSuratBerharga.sql#L71) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `sm.SecType` : `SMALLINT` | `1` : `int` | `sm.SecType` → `INT` | `WHEN 1 THEN 'BOND'` |
| IC-003325 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApproveHapusTransaksiSuratBerharga:72](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveHapusTransaksiSuratBerharga.sql#L72) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `sm.SecType` : `SMALLINT` | `3` : `int` | `sm.SecType` → `INT` | `WHEN 3 THEN 'SRBI'` |
| IC-003428 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSBondsPricePopulate:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPricePopulate.sql#L27) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-003435 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSBondsPricePopulate:42](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPricePopulate.sql#L42) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-004419 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSJobCalcPerformanceHargaSuratBerharga:218](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSJobCalcPerformanceHargaSuratBerharga.sql#L218) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-004421 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSJobCalcPerformanceHargaSuratBerharga:227](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSJobCalcPerformanceHargaSuratBerharga.sql#L227) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-004471 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSMaintenanceParamHargaModal:335](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L335) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-004477 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSMaintenanceParamHargaModal:357](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L357) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-004892 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSObligasiBlotterDaily:78](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSObligasiBlotterDaily.sql#L78) | WHERE / P11 / COMPARISON | `sm.SecType` : `SMALLINT` | `1` : `int` | `sm.SecType` → `INT` | `sm.SecType = 1` |
| IC-004908 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSObligasiBlotterDaily:116](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSObligasiBlotterDaily.sql#L116) | WHERE / P11 / COMPARISON | `sm.SecType` : `SMALLINT` | `1` : `int` | `sm.SecType` → `INT` | `sm.SecType = 1` |
| IC-005346 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDataParamDeviasiTotalSpread:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDataParamDeviasiTotalSpread.sql#L16) | CASE / P11 / COMPARISON | `SecType` : `SMALLINT` | `1` : `int` | `SecType` → `INT` | `SecType = 1` |
| IC-005347 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDataParamDeviasiTotalSpread:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDataParamDeviasiTotalSpread.sql#L17) | CASE / P11 / COMPARISON | `SecType` : `SMALLINT` | `2` : `int` | `SecType` → `INT` | `SecType = 2` |
| IC-005348 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDataParamDeviasiTotalSpread:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDataParamDeviasiTotalSpread.sql#L18) | CASE / P11 / COMPARISON | `SecType` : `SMALLINT` | `3` : `int` | `SecType` → `INT` | `SecType = 3` |
| IC-005904 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateParamHargaModalNew:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamHargaModalNew.sql#L216) | WHERE / P11 / COMPARISON | `sm.SecType` : `SMALLINT` | `@pnSecurityType` : `INT` | `sm.SecType` → `INT` | `sm.SecType = @pnSecurityType` |
| IC-005921 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateParamHargaModalNew:372](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamHargaModalNew.sql#L372) | WHERE / P11 / COMPARISON | `sm.SecType` : `SMALLINT` | `@pnSecurityType` : `INT` | `sm.SecType` → `INT` | `sm.SecType = @pnSecurityType` |
| IC-006577 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprParamHargaModal:100](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamHargaModal.sql#L100) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-006584 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprParamHargaModal:118](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamHargaModal.sql#L118) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-006607 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprParamSpread:204](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L204) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-006618 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprParamSpread:222](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L222) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-006629 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprParamSpread:271](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L271) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-006641 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprParamSpread:302](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L302) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-006670 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprovalHargaModal:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalHargaModal.sql#L111) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-006675 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprovalHargaModal:119](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalHargaModal.sql#L119) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-006680 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprovalHargaModal:164](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalHargaModal.sql#L164) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-006686 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprovalHargaModal:186](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalHargaModal.sql#L186) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-007794 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSREPOReportBiayaMaterai:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L71) | CASE / P11 / COMPARISON | `sm.SecType` : `SMALLINT` | `1` : `int` | `sm.SecType` → `INT` | `sm.SecType = 1` |
| IC-007795 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSREPOReportBiayaMaterai:72](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L72) | CASE / P11 / COMPARISON | `sm.SecType` : `SMALLINT` | `3` : `int` | `sm.SecType` → `INT` | `sm.SecType = 3` |
| IC-007813 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSREPOReportBiayaMaterai:154](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L154) | CASE / P11 / COMPARISON | `sm.SecType` : `SMALLINT` | `1` : `int` | `sm.SecType` → `INT` | `sm.SecType = 1` |
| IC-007814 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSREPOReportBiayaMaterai:155](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L155) | CASE / P11 / COMPARISON | `sm.SecType` : `SMALLINT` | `3` : `int` | `sm.SecType` → `INT` | `sm.SecType = 3` |
| IC-008167 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportRekonMaterai:118](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L118) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `sm.SecType` : `SMALLINT` | `1` : `int` | `sm.SecType` → `INT` | `WHEN 1 THEN 'BOND'` |
| IC-008168 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportRekonMaterai:119](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L119) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `sm.SecType` : `SMALLINT` | `3` : `int` | `sm.SecType` → `INT` | `WHEN 3 THEN 'SRBI'` |
| IC-008183 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportRekonMaterai:150](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L150) | WHERE / P11 / COMPARISON | `sm.SecType` : `SMALLINT` | `1` : `int` | `sm.SecType` → `INT` | `sm.SecType = 1` |
| IC-008184 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportRekonMaterai:152](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L152) | WHERE / P11 / COMPARISON | `sm.SecType` : `SMALLINT` | `3` : `int` | `sm.SecType` → `INT` | `sm.SecType = 3` |
| IC-008189 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportRekonMaterai:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L156) | WHERE / P11 / IN_VALUE_TYPE | `sm.SecType` : `SMALLINT` | `1` : `int` | `sm.SecType` → `INT` | `sm.SecType IN (1, 3)` |
| IC-008532 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSResetTimeTrxSecCabang:38](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResetTimeTrxSecCabang.sql#L38) | WHERE / P11 / COMPARISON | `SecType` : `SMALLINT` | `1` : `int` | `SecType` → `INT` | `SecType = 1` |
| IC-009037 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailPopulateCI:68](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulateCI.sql#L68) | CASE / P11 / COMPARISON | `SecType` : `SMALLINT` | `1` : `int` | `SecType` → `INT` | `SecType = 1` |
| IC-009047 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailPopulateLC:166](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulateLC.sql#L166) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `c.SecType` : `SMALLINT` | `1` : `int` | `c.SecType` → `INT` | `when 1 then b.DealPrice` |
| IC-009071 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailPopulateLCNew:211](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulateLCNew.sql#L211) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `b.SecType` : `SMALLINT` | `1` : `int` | `b.SecType` → `INT` | `when 1 then a.DealPrice` |
| IC-009509 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRpt2ndBondTrx:92](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRpt2ndBondTrx.sql#L92) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `sm.SecType` : `SMALLINT` | `1` : `int` | `sm.SecType` → `INT` | `when 1 then 'Obligasi'` |
| IC-009993 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptFeeBranch:91](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFeeBranch.sql#L91) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `sm.SecType` : `SMALLINT` | `1` : `int` | `sm.SecType` → `INT` | `when 1 then ( case trt.TrxDesc when 'Allotment' then 0.5 * Nominal else 0.1/100 * Nominal end )` |
| IC-009996 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptFeeBranch:109](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFeeBranch.sql#L109) | WHERE / P11 / IN_VALUE_TYPE | `sm.SecType` : `SMALLINT` | `1` : `int` | `sm.SecType` → `INT` | `sm.SecType in (1,2)` |
| IC-010097 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiBONDNew:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBONDNew.sql#L98) | WHERE / P11 / COMPARISON | `sm.SecType` : `SMALLINT` | `1` : `int` | `sm.SecType` → `INT` | `sm.SecType = 1` |
| IC-010101 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiBONDNew:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBONDNew.sql#L114) | WHERE / P11 / COMPARISON | `sm.SecType` : `SMALLINT` | `1` : `int` | `sm.SecType` → `INT` | `sm.SecType = 1` |
| IC-010104 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiBONDNew:133](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBONDNew.sql#L133) | WHERE / P11 / COMPARISON | `sm.SecType` : `SMALLINT` | `1` : `int` | `sm.SecType` → `INT` | `sm.SecType = 1` |
| IC-010107 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiBONDNew:152](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBONDNew.sql#L152) | WHERE / P11 / COMPARISON | `sm.SecType` : `SMALLINT` | `1` : `int` | `sm.SecType` → `INT` | `sm.SecType = 1` |
| IC-010170 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiSRBI:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiSRBI.sql#L86) | WHERE / P11 / COMPARISON | `sm.SecType` : `SMALLINT` | `3` : `int` | `sm.SecType` → `INT` | `sm.SecType = 3` |
| IC-011737 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveParamHargaModal:456](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModal.sql#L456) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-011742 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveParamHargaModal:465](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModal.sql#L465) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-011747 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveParamHargaModal:508](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModal.sql#L508) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-011753 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveParamHargaModal:531](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModal.sql#L531) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-011860 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveParamHargaModalNew:582](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L582) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-011865 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveParamHargaModalNew:591](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L591) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-011870 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveParamHargaModalNew:636](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L636) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-011876 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveParamHargaModalNew:659](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L659) | CASE / P11 / COMPARISON | `sec.SecType` : `SMALLINT` | `3` : `int` | `sec.SecType` → `INT` | `sec.SecType = 3` |
| IC-012127 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchECSecurityNoHargaORI:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchECSecurityNoHargaORI.sql#L43) | WHERE / P11 / IN_VALUE_TYPE | `a.SecType` : `SMALLINT` | `1` : `int` | `a.SecType` → `INT` | `a.SecType IN (1, 3)` |
| IC-012129 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchECSecurityNoHargaORI:57](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchECSecurityNoHargaORI.sql#L57) | WHERE / P11 / IN_VALUE_TYPE | `a.SecType` : `SMALLINT` | `1` : `int` | `a.SecType` → `INT` | `a.SecType IN (1, 3)` |
| IC-012131 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchECSecurityNoHargaORI:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchECSecurityNoHargaORI.sql#L70) | WHERE / P11 / IN_VALUE_TYPE | `a.SecType` : `SMALLINT` | `1` : `int` | `a.SecType` → `INT` | `a.SecType IN (1, 3)` |
| IC-012132 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchECSecurityNoHargaORI:82](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchECSecurityNoHargaORI.sql#L82) | WHERE / P11 / IN_VALUE_TYPE | `a.SecType` : `SMALLINT` | `1` : `int` | `a.SecType` → `INT` | `a.SecType IN (1, 3)` |
| IC-012134 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchECSecurityNoHargaORI:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchECSecurityNoHargaORI.sql#L96) | WHERE / P11 / IN_VALUE_TYPE | `a.SecType` : `SMALLINT` | `1` : `int` | `a.SecType` → `INT` | `a.SecType IN (1, 3)` |
| IC-012156 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchSRBIBlotterSummary:57](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchSRBIBlotterSummary.sql#L57) | WHERE / P11 / COMPARISON | `sm.SecType` : `SMALLINT` | `3` : `int` | `sm.SecType` → `INT` | `sm.SecType = 3` |
| IC-012164 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchSRBIBlotterSummary:79](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchSRBIBlotterSummary.sql#L79) | WHERE / P11 / COMPARISON | `sm.SecType` : `SMALLINT` | `3` : `int` | `sm.SecType` → `INT` | `sm.SecType = 3` |
| IC-012175 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchSecurityNo:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchSecurityNo.sql#L34) | WHERE / P11 / COMPARISON | `a.SecType` : `SMALLINT` | `1` : `int` | `a.SecType` → `INT` | `a.SecType = 1` |
| IC-012176 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchSecurityNo:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchSecurityNo.sql#L44) | WHERE / P11 / COMPARISON | `a.SecType` : `SMALLINT` | `1` : `int` | `a.SecType` → `INT` | `a.SecType = 1` |
| IC-012177 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchSecurityNo:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchSecurityNo.sql#L53) | WHERE / P11 / COMPARISON | `a.SecType` : `SMALLINT` | `1` : `int` | `a.SecType` → `INT` | `a.SecType = 1` |
| IC-012178 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchSecurityNo:62](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchSecurityNo.sql#L62) | WHERE / P11 / COMPARISON | `a.SecType` : `SMALLINT` | `1` : `int` | `a.SecType` → `INT` | `a.SecType = 1` |
| IC-012179 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchSecurityNo:72](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchSecurityNo.sql#L72) | WHERE / P11 / COMPARISON | `a.SecType` : `SMALLINT` | `1` : `int` | `a.SecType` → `INT` | `a.SecType = 1` |
| IC-012180 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchSecurityNoHargaORI:35](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchSecurityNoHargaORI.sql#L35) | WHERE / P11 / IN_VALUE_TYPE | `a.SecType` : `SMALLINT` | `1` : `int` | `a.SecType` → `INT` | `a.SecType IN (1, 3)` |
| IC-012181 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchSecurityNoHargaORI:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchSecurityNoHargaORI.sql#L48) | WHERE / P11 / IN_VALUE_TYPE | `a.SecType` : `SMALLINT` | `1` : `int` | `a.SecType` → `INT` | `a.SecType IN (1, 3)` |
| IC-012182 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchSecurityNoHargaORI:60](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchSecurityNoHargaORI.sql#L60) | WHERE / P11 / IN_VALUE_TYPE | `a.SecType` : `SMALLINT` | `1` : `int` | `a.SecType` → `INT` | `a.SecType IN (1, 3)` |
| IC-012183 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchSecurityNoHargaORI:72](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchSecurityNoHargaORI.sql#L72) | WHERE / P11 / IN_VALUE_TYPE | `a.SecType` : `SMALLINT` | `1` : `int` | `a.SecType` → `INT` | `a.SecType IN (1, 3)` |
| IC-012184 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchSecurityNoHargaORI:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchSecurityNoHargaORI.sql#L85) | WHERE / P11 / IN_VALUE_TYPE | `a.SecType` : `SMALLINT` | `1` : `int` | `a.SecType` → `INT` | `a.SecType IN (1, 3)` |
| IC-013687 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_StatusBondMaturity:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_StatusBondMaturity.sql#L63) | JOIN_ON / P11 / COMPARISON | `sm.SecType` : `SMALLINT` | `1` : `int` | `sm.SecType` → `INT` | `sm.SecType = 1` |
| IC-013697 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_StatusSBIMaturity:62](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_StatusSBIMaturity.sql#L62) | JOIN_ON / P11 / COMPARISON | `sm.SecType` : `SMALLINT` | `2` : `int` | `sm.SecType` → `INT` | `sm.SecType = 2` |
| IC-002970 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SRBIProcessEmailNotifMaturityDateH1:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH1.sql#L61) | JOIN_ON / P01 / COMPARISON | `a.SecType` : `SMALLINT` | `b.Value` : `VARCHAR(255)` | `b.Value` → `SMALLINT` | `a.SecType = b.Value` |
| IC-002981 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SRBIProcessEmailNotifMaturityDateH3:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH3.sql#L61) | JOIN_ON / P01 / COMPARISON | `a.SecType` : `SMALLINT` | `b.Value` : `VARCHAR(255)` | `b.Value` → `SMALLINT` | `a.SecType = b.Value` |
| IC-002992 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SRBIProcessEmailNotifMaturityDateH7:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH7.sql#L61) | JOIN_ON / P01 / COMPARISON | `a.SecType` : `SMALLINT` | `b.Value` : `VARCHAR(255)` | `b.Value` → `SMALLINT` | `a.SecType = b.Value` |
| IC-010798 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORIDetail:135](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L135) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `SecType` : `SMALLINT` | `0` : `int` | `0` → `SMALLINT` | `isnull(SecType,0)` |
| IC-010931 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnesty:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L131) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `SecType` : `SMALLINT` | `0` : `int` | `0` → `SMALLINT` | `isnull(SecType,0)` |
| IC-011131 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:116](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L116) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `SecType` : `SMALLINT` | `0` : `int` | `0` → `SMALLINT` | `ISNULL(SecType, 0)` |
| IC-011345 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnesty_Backdated:117](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty_Backdated.sql#L117) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `SecType` : `SMALLINT` | `0` : `int` | `0` → `SMALLINT` | `isnull(SecType,0)` |

<a id="kolom-secondmarketdate"></a>
## SecondMarketDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008983 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailORIAllotment:32](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L32) | WHERE / P04 / COMPARISON | `sm.SecondMarketDate` : `DATETIME` | `convert(varchar(8), @dCurrentDate, 112)` : `varchar(8)` | `convert(varchar(8), @dCurrentDate, 112)` → `DATETIME` | `sm.SecondMarketDate > convert(varchar(8), @dCurrentDate, 112)` |
| IC-008989 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailORIAllotment:73](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L73) | WHERE / P04 / COMPARISON | `sm.SecondMarketDate` : `DATETIME` | `convert(varchar(8), @dCurrentDate, 112)` : `varchar(8)` | `convert(varchar(8), @dCurrentDate, 112)` → `DATETIME` | `sm.SecondMarketDate > convert(varchar(8), @dCurrentDate, 112)` |
| IC-008995 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailORIAllotment:119](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L119) | WHERE / P04 / COMPARISON | `sm.SecondMarketDate` : `DATETIME` | `convert(varchar(8), @dCurrentDate, 112)` : `varchar(8)` | `convert(varchar(8), @dCurrentDate, 112)` → `DATETIME` | `sm.SecondMarketDate > convert(varchar(8), @dCurrentDate, 112)` |
| IC-009004 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailORIAllotment:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailORIAllotment.sql#L168) | WHERE / P04 / COMPARISON | `sm.SecondMarketDate` : `DATETIME` | `convert(varchar(8), @dCurrentDate, 112)` : `varchar(8)` | `convert(varchar(8), @dCurrentDate, 112)` → `DATETIME` | `sm.SecondMarketDate > convert(varchar(8), @dCurrentDate, 112)` |

<a id="kolom-tambahkuota1"></a>
## TambahKuota1

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L48). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 10 lokasi / 9 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002717 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:212](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L212) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TambahKuota1` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(TambahKuota1, 0)` |
| IC-006895 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessInsertORIOrderCommit:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertORIOrderCommit.sql#L113) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TambahKuota1` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(TambahKuota1,0)` |
| IC-007299 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSecurityMaster:1079](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMaster.sql#L1079) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.TambahKuota1` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(smtm.TambahKuota1, 0)` |
| IC-007303 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSecurityMaster:1086](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMaster.sql#L1086) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.TambahKuota1` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(smtm.TambahKuota1, 0)` |
| IC-012241 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSplitOrder:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSplitOrder.sql#L45) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TambahKuota1` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(TambahKuota1, 0)` |
| IC-012508 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUpdateOrderPriority:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateOrderPriority.sql#L28) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TambahKuota1` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(TambahKuota1, 0)` |
| IC-013092 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_InsertORI_Order_TT_Original:185](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORI_Order_TT_Original.sql#L185) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TambahKuota1` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(TambahKuota1,0)` |
| IC-013436 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusORI_Order_TT:212](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L212) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.TambahKuota1` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(smtm.TambahKuota1, 0)` |
| IC-013925 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidateInsertBuyBackOM:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateInsertBuyBackOM.sql#L41) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TambahKuota1` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(TambahKuota1,0)` |
| IC-013948 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidatePerubahanOrderSuratBerharga:38](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidatePerubahanOrderSuratBerharga.sql#L38) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TambahKuota1` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(TambahKuota1,0)` |

<a id="kolom-tambahkuota2"></a>
## TambahKuota2

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L49). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 12 lokasi / 9 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002718 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:213](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L213) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TambahKuota2` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(TambahKuota2, 0)` |
| IC-002719 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:214](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L214) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TotalKuotaOrder + TambahKuota1 + TambahKuota2` : `MONEY (combined result; facets unknown)` | `0` : `int` | `0` → `MONEY (combined result; facets unknown)` | `ISNULL(TotalKuotaOrder + TambahKuota1 + TambahKuota2, 0)` |
| IC-006896 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessInsertORIOrderCommit:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertORIOrderCommit.sql#L113) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TambahKuota2` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(TambahKuota2,0)` |
| IC-007300 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSecurityMaster:1080](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMaster.sql#L1080) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.TambahKuota2` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(smtm.TambahKuota2, 0)` |
| IC-007305 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSecurityMaster:1090](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMaster.sql#L1090) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.TambahKuota2` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(smtm.TambahKuota2, 0)` |
| IC-012242 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSplitOrder:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSplitOrder.sql#L45) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TambahKuota2` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(TambahKuota2, 0)` |
| IC-012509 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUpdateOrderPriority:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateOrderPriority.sql#L28) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TambahKuota2` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(TambahKuota2, 0)` |
| IC-013093 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_InsertORI_Order_TT_Original:185](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORI_Order_TT_Original.sql#L185) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TambahKuota2` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(TambahKuota2,0)` |
| IC-013437 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusORI_Order_TT:213](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L213) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.TambahKuota2` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(smtm.TambahKuota2, 0)` |
| IC-013438 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusORI_Order_TT:214](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L214) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.TotalKuotaOrder + smtm.TambahKuota1 + smtm.TambahKuota2` : `MONEY (combined result; facets unknown)` | `0` : `int` | `0` → `MONEY (combined result; facets unknown)` | `ISNULL(smtm.TotalKuotaOrder + smtm.TambahKuota1 + smtm.TambahKuota2, 0)` |
| IC-013926 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidateInsertBuyBackOM:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateInsertBuyBackOM.sql#L41) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TambahKuota2` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(TambahKuota2,0)` |
| IC-013949 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidatePerubahanOrderSuratBerharga:38](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidatePerubahanOrderSuratBerharga.sql#L38) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TambahKuota2` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(TambahKuota2,0)` |

<a id="kolom-tanggalbuyback"></a>
## TanggalBuyBack

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L50). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-010020 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptKonfirmasiBOND:124](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L124) | WHERE / P04 / COMPARISON | `sm.TanggalBuyBack` : `DATETIME` | `convert(varchar(10), getdate(), 112)` : `varchar(10)` | `convert(varchar(10), getdate(), 112)` → `DATETIME` | `sm.TanggalBuyBack > convert(varchar(10), getdate(), 112)` |
| IC-010106 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptKonfirmasiBONDNew:150](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBONDNew.sql#L150) | WHERE / P04 / COMPARISON | `sm.TanggalBuyBack` : `DATETIME` | `convert(varchar(10), getdate(), 112)` : `varchar(10)` | `convert(varchar(10), getdate(), 112)` → `DATETIME` | `sm.TanggalBuyBack > convert(varchar(10), getdate(), 112)` |

<a id="kolom-tax"></a>
## Tax

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 30 lokasi / 12 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006207 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBuyBack:102](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L102) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `smtm.Tax` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `smtm.Tax / 100` |
| IC-006208 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBuyBack:102](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L102) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `(smtm.Tax / 100)` : `DECIMAL (combined result; facets unknown)` | `( ((smtm.HargaBuyBack/100) * oott.NominalBuyBack) - oott.NominalBuyBack )` : `MONEY (combined result; facets unknown)` | `( ((smtm.HargaBuyBack/100) * oott.NominalBuyBack) - oott.NominalBuyBack )` → `DECIMAL` | `(smtm.Tax / 100) * ( ((smtm.HargaBuyBack/100) * oott.NominalBuyBack) - oott.NominalBuyBack )` |
| IC-006220 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBuyBack:176](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L176) | DML_CONDITION / P05 / BINARY_Multiply | `( isnull(smtm.Tax, 0) / 100.00 )` : `DECIMAL (combined result; facets unknown)` | `isnull(tmp.AccruedInterest, 0)` : `money` | `isnull(tmp.AccruedInterest, 0)` → `DECIMAL` | `isnull(tmp.AccruedInterest, 0) * ( isnull(smtm.Tax, 0) / 100.00 )` |
| IC-006221 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateSettlementBuyBack:176](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L176) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.Tax` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(smtm.Tax, 0)` |
| IC-006224 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBuyBack:182](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L182) | DML_CONDITION / P05 / BINARY_Multiply | `( isnull(smtm.Tax, 0) / 100.00 )` : `DECIMAL (combined result; facets unknown)` | `isnull(tmp.AccruedInterest, 0)` : `money` | `isnull(tmp.AccruedInterest, 0)` → `DECIMAL` | `isnull(tmp.AccruedInterest, 0) * ( isnull(smtm.Tax, 0) / 100.00 )` |
| IC-006225 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateSettlementBuyBack:182](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L182) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.Tax` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(smtm.Tax, 0)` |
| IC-006228 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBuyBack:188](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L188) | DML_CONDITION / P05 / BINARY_Multiply | `( isnull(smtm.Tax, 0) / 100.00 )` : `DECIMAL (combined result; facets unknown)` | `isnull(tmp.AccruedInterest, 0)` : `money` | `isnull(tmp.AccruedInterest, 0)` → `DECIMAL` | `isnull(tmp.AccruedInterest, 0) * ( isnull(smtm.Tax, 0) / 100.00 )` |
| IC-006229 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateSettlementBuyBack:188](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L188) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.Tax` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(smtm.Tax, 0)` |
| IC-006496 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateTransaksiBuyBack:153](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTransaksiBuyBack.sql#L153) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `Tax` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `Tax / 100` |
| IC-006497 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateTransaksiBuyBack:153](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTransaksiBuyBack.sql#L153) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `(Tax / 100)` : `DECIMAL (combined result; facets unknown)` | `(((HargaBuyBack-100)/100) * sum(NominalBuyBack))` : `MONEY (combined result; facets unknown)` | `(((HargaBuyBack-100)/100) * sum(NominalBuyBack))` → `DECIMAL` | `(((HargaBuyBack-100)/100) * sum(NominalBuyBack)) * (Tax / 100)` |
| IC-006498 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateTransaksiBuyBack:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTransaksiBuyBack.sql#L156) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `Tax` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `Tax / 100` |
| IC-006509 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateTrxBuyBackPerdana:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L137) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.Tax` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(sm.Tax, 0)` |
| IC-007428 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessSettlementBuyBack:161](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBuyBack.sql#L161) | DML_CONDITION / P05 / BINARY_Divide | `smtm.Tax` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `smtm.Tax / 100` |
| IC-007429 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessSettlementBuyBack:161](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBuyBack.sql#L161) | DML_CONDITION / P05 / BINARY_Multiply | `(smtm.Tax / 100)` : `DECIMAL (combined result; facets unknown)` | `(((smtm.HargaBuyBack/100) * oott.NominalBuyBack) - oott.NominalBuyBack)` : `MONEY (combined result; facets unknown)` | `(((smtm.HargaBuyBack/100) * oott.NominalBuyBack) - oott.NominalBuyBack)` → `DECIMAL` | `(smtm.Tax / 100) * (((smtm.HargaBuyBack/100) * oott.NominalBuyBack) - oott.NominalBuyBack)` |
| IC-007436 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessSettlementBuyBack:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBuyBack.sql#L168) | CASE / P05 / BINARY_Divide | `smtm.Tax` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `smtm.Tax / 100` |
| IC-007437 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessSettlementBuyBack:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBuyBack.sql#L168) | CASE / P05 / BINARY_Multiply | `(smtm.Tax / 100)` : `DECIMAL (combined result; facets unknown)` | `( ((smtm.HargaBuyBack/100) * oott.NominalBuyBack) - oott.NominalBuyBack )` : `MONEY (combined result; facets unknown)` | `( ((smtm.HargaBuyBack/100) * oott.NominalBuyBack) - oott.NominalBuyBack )` → `DECIMAL` | `(smtm.Tax / 100) * ( ((smtm.HargaBuyBack/100) * oott.NominalBuyBack) - oott.NominalBuyBack )` |
| IC-010621 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPajakKupon:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPajakKupon.sql#L98) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `smtm.Tax` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `smtm.Tax/100` |
| IC-010709 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORI:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORI.sql#L105) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `Tax` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `Tax / 100` |
| IC-010794 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORIDetail:123](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L123) | CASE / P11 / COMPARISON | `Tax` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `Tax > 100` |
| IC-010795 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORIDetail:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L125) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `Tax` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `Tax / 100` |
| IC-010925 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnesty:117](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L117) | CASE / P11 / COMPARISON | `Tax` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `Tax > 100` |
| IC-010926 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnesty:119](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L119) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `Tax` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `Tax / 100` |
| IC-011039 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnestyEOM:74](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyEOM.sql#L74) | CASE / P11 / COMPARISON | `Tax` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `Tax > 100` |
| IC-011040 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnestyEOM:76](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyEOM.sql#L76) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `Tax` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `Tax / 100` |
| IC-011125 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:102](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L102) | CASE / P11 / COMPARISON | `Tax` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `Tax > 100` |
| IC-011126 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L104) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `Tax` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `Tax / 100` |
| IC-011251 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnestySPT:78](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestySPT.sql#L78) | CASE / P11 / COMPARISON | `Tax` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `Tax > 100` |
| IC-011252 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnestySPT:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestySPT.sql#L80) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `Tax` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `Tax / 100` |
| IC-011339 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnesty_Backdated:103](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty_Backdated.sql#L103) | CASE / P11 / COMPARISON | `Tax` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `Tax > 100` |
| IC-011340 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnesty_Backdated:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty_Backdated.sql#L105) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `Tax` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `Tax / 100` |

<a id="kolom-taxwnawithoutnpwp"></a>
## TaxWNAWithoutNPWP

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L57). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 11 lokasi / 11 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002705 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:118](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L118) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TaxWNAWithoutNPWP` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(TaxWNAWithoutNPWP, 0)` |
| IC-008732 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRetailCalculateFee:206](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L206) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TaxWNAWithoutNPWP` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(TaxWNAWithoutNPWP,0)` |
| IC-010623 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPajakKupon:109](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPajakKupon.sql#L109) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `isnull(smtm.TaxWNAWithoutNPWP,0.0)` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `isnull(smtm.TaxWNAWithoutNPWP,0.0) / 100` |
| IC-010711 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORI:112](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORI.sql#L112) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `isnull(TaxWNAWithoutNPWP,0.0)` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `isnull(TaxWNAWithoutNPWP,0.0) /100` |
| IC-010797 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORIDetail:134](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L134) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `isnull(TaxWNAWithoutNPWP,0.0)` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `isnull(TaxWNAWithoutNPWP,0.0) / 100` |
| IC-010928 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnesty:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L127) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `isnull(TaxWNAWithoutNPWP,0.0)` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `isnull(TaxWNAWithoutNPWP,0.0) / 100` |
| IC-011042 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnestyEOM:84](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyEOM.sql#L84) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `isnull(TaxWNAWithoutNPWP,0.0)` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `isnull(TaxWNAWithoutNPWP,0.0) / 100` |
| IC-011128 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:112](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L112) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `ISNULL(TaxWNAWithoutNPWP, 0.0)` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `ISNULL(TaxWNAWithoutNPWP, 0.0) / 100` |
| IC-011254 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnestySPT:88](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestySPT.sql#L88) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `isnull(TaxWNAWithoutNPWP,0.0)` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `isnull(TaxWNAWithoutNPWP,0.0) / 100` |
| IC-011342 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnesty_Backdated:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty_Backdated.sql#L113) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `isnull(TaxWNAWithoutNPWP,0.0)` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL` | `isnull(TaxWNAWithoutNPWP,0.0) / 100` |
| IC-013336 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L66) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TaxWNAWithoutNPWP` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(TaxWNAWithoutNPWP,0)` |

<a id="kolom-totalkuotaorder"></a>
## TotalKuotaOrder

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L53). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 11 lokasi / 10 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004788 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSORIGetProductDetail:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIGetProductDetail.sql#L85) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `TotalKuotaOrder` : `MONEY` | `@tmp_Data.KuotaSeri` : `decimal(21,2)` | `TotalKuotaOrder` → `decimal(21,2)` | `@tmp_Data.KuotaSeri <- TotalKuotaOrder` |
| IC-002716 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:211](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L211) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TotalKuotaOrder` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(TotalKuotaOrder, 0)` |
| IC-006897 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessInsertORIOrderCommit:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertORIOrderCommit.sql#L113) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TotalKuotaOrder` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(TotalKuotaOrder,0)` |
| IC-007298 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSecurityMaster:1078](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMaster.sql#L1078) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.TotalKuotaOrder` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(smtm.TotalKuotaOrder, 0)` |
| IC-007301 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSecurityMaster:1082](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMaster.sql#L1082) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.TotalKuotaOrder` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(smtm.TotalKuotaOrder, 0)` |
| IC-012243 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSplitOrder:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSplitOrder.sql#L45) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TotalKuotaOrder` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(TotalKuotaOrder, 0)` |
| IC-012510 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUpdateOrderPriority:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateOrderPriority.sql#L28) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TotalKuotaOrder` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(TotalKuotaOrder, 0)` |
| IC-013094 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_InsertORI_Order_TT_Original:185](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORI_Order_TT_Original.sql#L185) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TotalKuotaOrder` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(TotalKuotaOrder,0)` |
| IC-013435 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusORI_Order_TT:211](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L211) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smtm.TotalKuotaOrder` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(smtm.TotalKuotaOrder, 0)` |
| IC-013927 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidateInsertBuyBackOM:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateInsertBuyBackOM.sql#L41) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TotalKuotaOrder` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(TotalKuotaOrder,0)` |
| IC-013950 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidatePerubahanOrderSuratBerharga:38](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidatePerubahanOrderSuratBerharga.sql#L38) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TotalKuotaOrder` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(TotalKuotaOrder,0)` |

<a id="kolom-transactionfee"></a>
## TransactionFee

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008971 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailCalculateTransactionFee:42](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateTransactionFee.sql#L42) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `TransactionFee` : `MONEY` | `@mTransactionFee` : `decimal (20, 5)` | `TransactionFee` → `decimal (20, 5)` | `@mTransactionFee = TransactionFee` |
| IC-008972 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailCalculateTransactionFee:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateTransactionFee.sql#L56) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `TransactionFee` : `MONEY` | `@mTransactionFee` : `decimal (20, 5)` | `TransactionFee` → `decimal (20, 5)` | `@mTransactionFee = TransactionFee` |
| IC-008973 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailCalculateTransactionFee:75](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateTransactionFee.sql#L75) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `TransactionFee` : `MONEY` | `@mTransactionFee` : `decimal (20, 5)` | `TransactionFee` → `decimal (20, 5)` | `@mTransactionFee = TransactionFee` |
| IC-008974 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailCalculateTransactionFee:100](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateTransactionFee.sql#L100) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `TransactionFee` : `MONEY` | `@mTransactionFee` : `decimal (20, 5)` | `TransactionFee` → `decimal (20, 5)` | `@mTransactionFee = TransactionFee` |
| IC-008975 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailCalculateTransactionFee:109](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateTransactionFee.sql#L109) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `TransactionFee` : `MONEY` | `@mTransactionFee` : `decimal (20, 5)` | `TransactionFee` → `decimal (20, 5)` | `@mTransactionFee = TransactionFee` |

<a id="kolom-usingearlyredemption"></a>
## UsingEarlyRedemption

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMaster_TM.sql#L76). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 11 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000971 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.ORISBNIPopulateEarlyRedemption:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateEarlyRedemption.sql#L113) | WHERE / P11 / COMPARISON | `UsingEarlyRedemption` : `BIT` | `1` : `int` | `UsingEarlyRedemption` → `INT` | `UsingEarlyRedemption = 1` |
| IC-009142 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSrcSecIdEarlyRedem:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSrcSecIdEarlyRedem.sql#L31) | WHERE / P11 / COMPARISON | `UsingEarlyRedemption` : `BIT` | `1` : `int` | `UsingEarlyRedemption` → `INT` | `UsingEarlyRedemption = 1` |
| IC-009143 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSrcSecIdEarlyRedem:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSrcSecIdEarlyRedem.sql#L41) | WHERE / P11 / COMPARISON | `UsingEarlyRedemption` : `BIT` | `1` : `int` | `UsingEarlyRedemption` → `INT` | `UsingEarlyRedemption = 1` |
| IC-009144 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSrcSecIdEarlyRedem:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSrcSecIdEarlyRedem.sql#L49) | WHERE / P11 / COMPARISON | `UsingEarlyRedemption` : `BIT` | `1` : `int` | `UsingEarlyRedemption` → `INT` | `UsingEarlyRedemption = 1` |
| IC-009145 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSrcSecIdEarlyRedem:58](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSrcSecIdEarlyRedem.sql#L58) | WHERE / P11 / COMPARISON | `UsingEarlyRedemption` : `BIT` | `1` : `int` | `UsingEarlyRedemption` → `INT` | `UsingEarlyRedemption = 1` |
| IC-009146 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSrcSecIdEarlyRedem:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSrcSecIdEarlyRedem.sql#L65) | WHERE / P11 / COMPARISON | `UsingEarlyRedemption` : `BIT` | `1` : `int` | `UsingEarlyRedemption` → `INT` | `UsingEarlyRedemption = 1` |
| IC-012429 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSTransactionSrcSecId:32](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransactionSrcSecId.sql#L32) | WHERE / P11 / COMPARISON | `UsingEarlyRedemption` : `BIT` | `0` : `int` | `UsingEarlyRedemption` → `INT` | `UsingEarlyRedemption = 0` |
| IC-012431 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSTransactionSrcSecId:42](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransactionSrcSecId.sql#L42) | WHERE / P11 / COMPARISON | `UsingEarlyRedemption` : `BIT` | `0` : `int` | `UsingEarlyRedemption` → `INT` | `UsingEarlyRedemption = 0` |
| IC-012433 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSTransactionSrcSecId:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransactionSrcSecId.sql#L49) | WHERE / P11 / COMPARISON | `UsingEarlyRedemption` : `BIT` | `0` : `int` | `UsingEarlyRedemption` → `INT` | `UsingEarlyRedemption = 0` |
| IC-012435 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSTransactionSrcSecId:57](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransactionSrcSecId.sql#L57) | WHERE / P11 / COMPARISON | `UsingEarlyRedemption` : `BIT` | `0` : `int` | `UsingEarlyRedemption` → `INT` | `UsingEarlyRedemption = 0` |
| IC-012437 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSTransactionSrcSecId:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransactionSrcSecId.sql#L64) | WHERE / P11 / COMPARISON | `UsingEarlyRedemption` : `BIT` | `0` : `int` | `UsingEarlyRedemption` → `INT` | `UsingEarlyRedemption = 0` |
