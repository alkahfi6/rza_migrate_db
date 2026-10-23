# Retail UAT.dbo.SecurityCouponDate_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 32 lokasi source, 12 procedure, 3 kolom. **A:** 9 lokasi pada 7 procedure. **B:** 3 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-periode"></a>
## Periode

**Tipe dasar:** `VARCHAR(8)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityCouponDate_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 12 lokasi / 9 SP. **A:** 9 lokasi / 7 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P03 — Tanggal teks bertemu date/time:** Tentukan kontrak YYYYMMDD/hari/timestamp. Bentuk nilai pencarian dalam tipe kolom atau gunakan rentang native date; jangan konversi kolom seluruh baris tanpa alasan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007363 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1605](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1605) | WHERE / P03 / COMPARISON | `smptm.Periode` : `VARCHAR(8)` | `@nLastCouponDate` : `DATETIME` | `smptm.Periode` → `DATETIME` | `smptm.Periode > @nLastCouponDate` |
| IC-009484 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRetailUpdateCouponDate:95](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateCouponDate.sql#L95) | WHERE / P03 / COMPARISON | `Periode` : `VARCHAR(8)` | `@dNextCouponDate` : `smalldatetime` | `Periode` → `SMALLDATETIME` | `Periode = @dNextCouponDate` |
| IC-009488 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRetailUpdateCouponDate:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateCouponDate.sql#L145) | WHERE / P03 / COMPARISON | `Periode` : `VARCHAR(8)` | `@dNextCouponDate` : `smalldatetime` | `Periode` → `SMALLDATETIME` | `Periode = @dNextCouponDate` |
| IC-009494 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRetailUpdateCouponDate_Dummy:83](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateCouponDate_Dummy.sql#L83) | WHERE / P03 / COMPARISON | `Periode` : `VARCHAR(8)` | `@dNextCouponDate` : `smalldatetime` | `Periode` → `SMALLDATETIME` | `Periode = @dNextCouponDate` |
| IC-009498 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRetailUpdateCouponDate_Dummy:133](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateCouponDate_Dummy.sql#L133) | WHERE / P03 / COMPARISON | `Periode` : `VARCHAR(8)` | `@dNextCouponDate` : `smalldatetime` | `Periode` → `SMALLDATETIME` | `Periode = @dNextCouponDate` |
| IC-010093 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiBOND:503](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L503) | WHERE / P03 / COMPARISON | `Periode` : `VARCHAR(8)` | `@cNextCouponDateOri` : `datetime` | `Periode` → `DATETIME` | `Periode > @cNextCouponDateOri` |
| IC-010155 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiBONDNew:392](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBONDNew.sql#L392) | WHERE / P03 / COMPARISON | `Periode` : `VARCHAR(8)` | `@cNextCouponDateOri` : `datetime` | `Periode` → `DATETIME` | `Periode > @cNextCouponDateOri` |
| IC-010212 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiSRBI:298](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiSRBI.sql#L298) | WHERE / P03 / COMPARISON | `Periode` : `VARCHAR(8)` | `@cNextCouponDateOri` : `datetime` | `Periode` → `DATETIME` | `Periode > @cNextCouponDateOri` |
| IC-011035 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnesty:693](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L693) | WHERE / P03 / COMPARISON | `[Periode]` : `VARCHAR(8)` | `@dCurrentWorkingDate` : `datetime` | `[Periode]` → `DATETIME` | `[Periode] > @dCurrentWorkingDate` |
| IC-010806 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptPayCouponORIDetail:166](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L166) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[Periode]` : `VARCHAR(8)` | `@dFirstCouponDate` : `datetime` | `[Periode]` → `datetime` | `@dFirstCouponDate = [Periode]` |
| IC-010939 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptPayCouponORITaxAmnesty:161](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L161) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[Periode]` : `VARCHAR(8)` | `@dFirstCouponDate` : `datetime` | `[Periode]` → `datetime` | `@dFirstCouponDate = [Periode]` |
| IC-011135 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L130) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[Periode]` : `VARCHAR(8)` | `@dFirstCouponDate` : `DATETIME` | `[Periode]` → `DATETIME` | `@dFirstCouponDate = [Periode]` |

<a id="kolom-done"></a>
## Done

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityCouponDate_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 17 lokasi / 12 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002751 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_trs_ListSecurityMaster_TM:339](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L339) | JOIN_ON / P11 / COMPARISON | `sc.Done` : `BIT` | `0` : `int` | `sc.Done` → `INT` | `sc.Done = 0` |
| IC-007362 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1604](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1604) | WHERE / P11 / COMPARISON | `smptm.Done` : `BIT` | `0` : `int` | `smptm.Done` → `INT` | `smptm.Done = 0` |
| IC-008828 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRetailCalculateFee:756](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L756) | CASE / P11 / COMPARISON | `ISNULL(Done,'')` : `BIT` | `0` : `int` | `ISNULL(Done,'')` → `INT` | `ISNULL(Done,'') = 0` |
| IC-010095 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiBOND:549](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L549) | WHERE / P11 / COMPARISON | `Done` : `BIT` | `0` : `int` | `Done` → `INT` | `Done = 0` |
| IC-010158 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiBONDNew:482](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBONDNew.sql#L482) | WHERE / P11 / COMPARISON | `Done` : `BIT` | `0` : `int` | `Done` → `INT` | `Done = 0` |
| IC-010214 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiSRBI:386](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiSRBI.sql#L386) | WHERE / P11 / COMPARISON | `Done` : `BIT` | `0` : `int` | `Done` → `INT` | `Done = 0` |
| IC-010807 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORIDetail:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L168) | WHERE / P11 / COMPARISON | `[Done]` : `BIT` | `0` : `int` | `[Done]` → `INT` | `[Done] = 0` |
| IC-010941 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnesty:163](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L163) | WHERE / P11 / COMPARISON | `[Done]` : `BIT` | `0` : `int` | `[Done]` → `INT` | `[Done] = 0` |
| IC-011137 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:134](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L134) | WHERE / P11 / COMPARISON | `[Done]` : `BIT` | `0` : `int` | `[Done]` → `INT` | `[Done] = 0` |
| IC-013369 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListSecurityMaster_TM:418](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L418) | JOIN_ON / P11 / COMPARISON | `sc.Done` : `BIT` | `0` : `int` | `sc.Done` → `INT` | `sc.Done = 0` |
| IC-007352 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1108) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitycoupondate_tr.Done` : `BIT` | `CASE WHEN cdp.Periode <= @nLastCouponDate THEN 1 ELSE 0 END` : `INT (combined result; facets unknown)` | `CASE WHEN cdp.Periode <= @nLastCouponDate THEN 1 ELSE 0 END` → `BIT` | `retail.dbo.securitycoupondate_tr.Done <- CASE WHEN cdp.Periode <= @nLastCouponDate THEN 1 ELSE 0 END` |
| IC-007364 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1608](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1608) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitycoupondate_tr.Done` : `BIT` | `CASE WHEN cdp.Periode <= @nLastCouponDate THEN 1 ELSE 0 END` : `INT (combined result; facets unknown)` | `CASE WHEN cdp.Periode <= @nLastCouponDate THEN 1 ELSE 0 END` → `BIT` | `retail.dbo.securitycoupondate_tr.Done <- CASE WHEN cdp.Periode <= @nLastCouponDate THEN 1 ELSE 0 END` |
| IC-008829 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRetailCalculateFee:756](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L756) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `Done` : `BIT` | `''` : `varchar` | `''` → `BIT` | `ISNULL(Done,'')` |
| IC-009487 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailUpdateCouponDate:142](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateCouponDate.sql#L142) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Done` : `BIT` | `1` : `int` | `1` → `BIT` | `Done = 1` |
| IC-009489 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailUpdateCouponDate:154](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateCouponDate.sql#L154) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitycoupondate_tr.Done` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.securitycoupondate_tr.Done <- 0` |
| IC-009497 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailUpdateCouponDate_Dummy:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateCouponDate_Dummy.sql#L130) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Done` : `BIT` | `1` : `int` | `1` → `BIT` | `Done = 1` |
| IC-009499 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailUpdateCouponDate_Dummy:142](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateCouponDate_Dummy.sql#L142) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitycoupondate_tr.Done` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.securitycoupondate_tr.Done <- 0` |

<a id="kolom-secid"></a>
## SecId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityCouponDate_TR.sql#L2). **Existing leading index:** SecurityCouponDate_TR_CN1.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-010940 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnesty:163](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L163) | WHERE / P11 / COMPARISON | `SecId` : `INT` | `@nSecId` : `bigint` | `SecId` → `BIGINT` | `SecId = @nSecId` |
| IC-011034 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnesty:693](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L693) | WHERE / P11 / COMPARISON | `SecId` : `INT` | `@nSecId` : `bigint` | `SecId` → `BIGINT` | `SecId = @nSecId` |
| IC-011136 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:132](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L132) | WHERE / P11 / COMPARISON | `SecId` : `INT` | `@nSecId` : `BIGINT` | `SecId` → `BIGINT` | `SecId = @nSecId` |
