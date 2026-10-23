# Retail UAT.dbo.TRSMateraiTransaction_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 85 lokasi source, 20 procedure, 12 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 6 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-accountblockactype"></a>
## AccountBlockACTYPE

**Tipe dasar:** `CHAR(3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMateraiTransaction_TT.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002168 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.Rvmp_TRSGetSecurityTransaction_TT:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGetSecurityTransaction_TT.sql#L36) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `ISNULL(mt.AccountBlockACTYPE, '')` : `CHAR(3)` | `@cMateraiAccountBlockACTYPE` : `varchar(10)` | `ISNULL(mt.AccountBlockACTYPE, '')` → `varchar(10)` | `@cMateraiAccountBlockACTYPE = ISNULL(mt.AccountBlockACTYPE, '')` |
| IC-004148 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSGetSecurityTransaction_TT:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetSecurityTransaction_TT.sql#L37) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `ISNULL(mt.AccountBlockACTYPE, '')` : `CHAR(3)` | `@cMateraiAccountBlockACTYPE` : `varchar(10)` | `ISNULL(mt.AccountBlockACTYPE, '')` → `varchar(10)` | `@cMateraiAccountBlockACTYPE = ISNULL(mt.AccountBlockACTYPE, '')` |
| IC-002167 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSGetSecurityTransaction_TT:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGetSecurityTransaction_TT.sql#L36) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mt.AccountBlockACTYPE` : `CHAR(3)` | `''` : `varchar` | `''` → `CHAR(3)` | `ISNULL(mt.AccountBlockACTYPE, '')` |
| IC-003036 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SecurityPopulateDetailReplaceTrxLelang:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateDetailReplaceTrxLelang.sql#L33) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mt.AccountBlockACTYPE` : `CHAR(3)` | `''` : `varchar` | `''` → `CHAR(3)` | `ISNULL(mt.AccountBlockACTYPE, '')` |
| IC-004147 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSGetSecurityTransaction_TT:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetSecurityTransaction_TT.sql#L37) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mt.AccountBlockACTYPE` : `CHAR(3)` | `''` : `varchar` | `''` → `CHAR(3)` | `ISNULL(mt.AccountBlockACTYPE, '')` |
| IC-004872 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSORIOrderSummary:87](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOrderSummary.sql#L87) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mtr.AccountBlockACTYPE` : `CHAR(3)` | `''` : `varchar` | `''` → `CHAR(3)` | `ISNULL(mtr.AccountBlockACTYPE,'')` |
| IC-012578 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdatePostingResult:655](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L655) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_tt.AccountBlockACTYPE` : `CHAR(3)` | `''` : `varchar` | `''` → `CHAR(3)` | `retail.dbo.trsmateraitransaction_tt.AccountBlockACTYPE <- ''` |

<a id="kolom-accountblocksequence"></a>
## AccountBlockSequence

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMateraiTransaction_TT.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012579 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdatePostingResult:655](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L655) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_tt.AccountBlockSequence` : `INT` | `'0'` : `varchar` | `'0'` → `INT` | `retail.dbo.trsmateraitransaction_tt.AccountBlockSequence <- '0'` |

<a id="kolom-dealid"></a>
## DealId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMateraiTransaction_TT.sql#L6). **Existing leading index:** TRSMateraiTransaction_TT_NN1.

**Frekuensi source:** 5 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003224 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSAmendSecurityTransaction_TT:987](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L987) | JOIN_ON / P01 / COMPARISON | `mt.DealId` : `BIGINT` | `tmp.DealNo` : `varchar(10)` | `tmp.DealNo` → `BIGINT` | `mt.DealId = tmp.DealNo` |
| IC-003226 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSAmendSecurityTransaction_TT:995](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L995) | JOIN_ON / P01 / COMPARISON | `mt.DealId` : `BIGINT` | `tmp.DealNo` : `varchar(10)` | `tmp.DealNo` → `BIGINT` | `mt.DealId = tmp.DealNo` |
| IC-003229 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSAmendSecurityTransaction_TT:1015](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L1015) | JOIN_ON / P01 / COMPARISON | `mt.DealId` : `BIGINT` | `tmp.DealNo` : `varchar(10)` | `tmp.DealNo` → `BIGINT` | `mt.DealId = tmp.DealNo` |
| IC-003230 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSAmendSecurityTransaction_TT:1078](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L1078) | JOIN_ON / P01 / COMPARISON | `mt.DealId` : `BIGINT` | `tmp.DealNo` : `varchar(10)` | `tmp.DealNo` → `BIGINT` | `mt.DealId = tmp.DealNo` |
| IC-008150 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportMateraiBuyBack:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportMateraiBuyBack.sql#L125) | JOIN_ON / P11 / COMPARISON | `mt.DealId` : `BIGINT` | `bb.BuyBackID` : `INT` | `bb.BuyBackID` → `BIGINT` | `bb.BuyBackID = mt.DealId` |

<a id="kolom-expirydate"></a>
## ExpiryDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMateraiTransaction_TT.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007060 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessOverBookingMateraiBuyBack:330](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql#L330) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_tt.ExpiryDate` : `DATETIME` | `CONVERT(VARCHAR(8),GETDATE(),112)` : `VARCHAR(8)` | `CONVERT(VARCHAR(8),GETDATE(),112)` → `DATETIME` | `retail.dbo.trsmateraitransaction_tt.ExpiryDate <- CONVERT(VARCHAR(8),GETDATE(),112)` |
| IC-012580 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdatePostingResult:655](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L655) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_tt.ExpiryDate` : `DATETIME` | `CONVERT(VARCHAR(8),GETDATE(),112)` : `VARCHAR(8)` | `CONVERT(VARCHAR(8),GETDATE(),112)` → `DATETIME` | `retail.dbo.trsmateraitransaction_tt.ExpiryDate <- CONVERT(VARCHAR(8),GETDATE(),112)` |

<a id="kolom-insertedby"></a>
## InsertedBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMateraiTransaction_TT.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001939 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSubmitTransaction:360](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSubmitTransaction.sql#L360) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_tt.InsertedBy` : `INT` | `tmp.LastUpdateNIK` : `VARCHAR(10)` | `tmp.LastUpdateNIK` → `INT` | `retail.dbo.trsmateraitransaction_tt.InsertedBy <- tmp.LastUpdateNIK` |
| IC-007061 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessOverBookingMateraiBuyBack:330](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql#L330) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_tt.InsertedBy` : `INT` | `'77777'` : `varchar` | `'77777'` → `INT` | `retail.dbo.trsmateraitransaction_tt.InsertedBy <- '77777'` |
| IC-012581 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdatePostingResult:655](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L655) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_tt.InsertedBy` : `INT` | `'0'` : `varchar` | `'0'` → `INT` | `retail.dbo.trsmateraitransaction_tt.InsertedBy <- '0'` |

<a id="kolom-isabsorbedbybank"></a>
## IsAbsorbedByBank

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMateraiTransaction_TT.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 19 lokasi / 11 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008190 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportRekonMaterai:174](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L174) | CASE / P11 / COMPARISON | `mt.IsAbsorbedByBank` : `BIT` | `1` : `int` | `mt.IsAbsorbedByBank` → `INT` | `mt.IsAbsorbedByBank = 1` |
| IC-008191 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportRekonMaterai:175](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L175) | CASE / P11 / COMPARISON | `mt.IsAbsorbedByBank` : `BIT` | `0` : `int` | `mt.IsAbsorbedByBank` → `INT` | `mt.IsAbsorbedByBank = 0` |
| IC-008192 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportRekonMaterai:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L180) | CASE / P11 / COMPARISON | `mt.IsAbsorbedByBank` : `BIT` | `1` : `int` | `mt.IsAbsorbedByBank` → `INT` | `mt.IsAbsorbedByBank = 1` |
| IC-008193 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportRekonMaterai:181](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L181) | CASE / P11 / COMPARISON | `mt.IsAbsorbedByBank` : `BIT` | `0` : `int` | `mt.IsAbsorbedByBank` → `INT` | `mt.IsAbsorbedByBank = 0` |
| IC-010085 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiBOND:417](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L417) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `mt.IsAbsorbedByBank` : `BIT` | `0` : `int` | `mt.IsAbsorbedByBank` → `INT` | `WHEN 0 THEN CASE WHEN ISNULL(mt.MateraiCost, 0) > 0 THEN 1 ELSE 0 END` |
| IC-010089 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiBOND:425](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L425) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `mt.IsAbsorbedByBank` : `BIT` | `0` : `int` | `mt.IsAbsorbedByBank` → `INT` | `WHEN 0 THEN ISNULL(mt.MateraiCost, 0)` |
| IC-010140 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiBONDNew:294](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBONDNew.sql#L294) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `mt.IsAbsorbedByBank` : `BIT` | `0` : `int` | `mt.IsAbsorbedByBank` → `INT` | `WHEN 0 THEN CASE WHEN ISNULL(mt.MateraiCost, 0) > 0 THEN 1 ELSE 0 END` |
| IC-010144 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiBONDNew:303](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBONDNew.sql#L303) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `mt.IsAbsorbedByBank` : `BIT` | `0` : `int` | `mt.IsAbsorbedByBank` → `INT` | `WHEN 0 THEN ISNULL(mt.MateraiCost, 0)` |
| IC-010202 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiSRBI:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiSRBI.sql#L225) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `mt.IsAbsorbedByBank` : `BIT` | `0` : `int` | `mt.IsAbsorbedByBank` → `INT` | `WHEN 0 THEN CASE WHEN ISNULL(mt.MateraiCost, 0) > 0 THEN 1 ELSE 0 END` |
| IC-010206 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiSRBI:233](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiSRBI.sql#L233) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `mt.IsAbsorbedByBank` : `BIT` | `0` : `int` | `mt.IsAbsorbedByBank` → `INT` | `WHEN 0 THEN ISNULL(mt.MateraiCost, 0)` |
| IC-001772 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.REPOPopulateHapusTransaksi:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.REPOPopulateHapusTransaksi.sql#L51) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mt.IsAbsorbedByBank` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(mt.IsAbsorbedByBank, 0)` |
| IC-001775 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.REPOPopulateHapusTransaksi:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.REPOPopulateHapusTransaksi.sql#L89) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mt.IsAbsorbedByBank` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(mt.IsAbsorbedByBank, 0)` |
| IC-002166 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSGetSecurityTransaction_TT:35](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGetSecurityTransaction_TT.sql#L35) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mt.IsAbsorbedByBank` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(mt.IsAbsorbedByBank, 0)` |
| IC-003037 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SecurityPopulateDetailReplaceTrxLelang:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateDetailReplaceTrxLelang.sql#L36) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mt.IsAbsorbedByBank` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(mt.IsAbsorbedByBank, 0)` |
| IC-003096 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SecurityPopulateTrxLelang:144](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateTrxLelang.sql#L144) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mt.IsAbsorbedByBank` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(mt.IsAbsorbedByBank, 0)` |
| IC-003099 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SecurityPopulateTrxLelang:158](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateTrxLelang.sql#L158) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mt.IsAbsorbedByBank` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(mt.IsAbsorbedByBank, 0)` |
| IC-004146 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSGetSecurityTransaction_TT:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetSecurityTransaction_TT.sql#L36) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mt.IsAbsorbedByBank` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(mt.IsAbsorbedByBank, 0)` |
| IC-007062 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessOverBookingMateraiBuyBack:330](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql#L330) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_tt.IsAbsorbedByBank` : `BIT` | `case when dd.STATUS = 6 then 1 else 0 end` : `INT (combined result; facets unknown)` | `case when dd.STATUS = 6 then 1 else 0 end` → `BIT` | `retail.dbo.trsmateraitransaction_tt.IsAbsorbedByBank <- case when dd.STATUS = 6 then 1 else 0 end` |
| IC-012582 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdatePostingResult:655](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L655) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_tt.IsAbsorbedByBank` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.trsmateraitransaction_tt.IsAbsorbedByBank <- 0` |

<a id="kolom-istrxcanceled"></a>
## IsTrxCanceled

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMateraiTransaction_TT.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001861 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.RepoProsesCancelTrx:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoProsesCancelTrx.sql#L337) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `mt.IsTrxCanceled` : `BIT` | `1` : `int` | `mt.IsTrxCanceled` → `INT` | `WHEN 1 THEN mt.IsAbsorbedByBank` |
| IC-003328 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApproveHapusTransaksiSuratBerharga:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveHapusTransaksiSuratBerharga.sql#L192) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `mt.IsTrxCanceled` : `BIT` | `1` : `int` | `mt.IsTrxCanceled` → `INT` | `WHEN 1 THEN mt.IsAbsorbedByBank` |
| IC-001862 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoProsesCancelTrx:340](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoProsesCancelTrx.sql#L340) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mt.IsTrxCanceled` : `BIT` | `1` : `int` | `1` → `BIT` | `mt.IsTrxCanceled = 1` |
| IC-003329 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApproveHapusTransaksiSuratBerharga:195](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveHapusTransaksiSuratBerharga.sql#L195) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mt.IsTrxCanceled` : `BIT` | `1` : `int` | `1` → `BIT` | `mt.IsTrxCanceled = 1` |
| IC-007063 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessOverBookingMateraiBuyBack:330](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql#L330) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_tt.IsTrxCanceled` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.trsmateraitransaction_tt.IsTrxCanceled <- 0` |

<a id="kolom-istrxdeleted"></a>
## IsTrxDeleted

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMateraiTransaction_TT.sql#L25). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001940 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSubmitTransaction:360](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSubmitTransaction.sql#L360) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_tt.IsTrxDeleted` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.trsmateraitransaction_tt.IsTrxDeleted <- 0` |
| IC-003018 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SecurityLelangTrxSubmit:399](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityLelangTrxSubmit.sql#L399) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `IsTrxDeleted` : `BIT` | `1` : `int` | `1` → `BIT` | `IsTrxDeleted = 1` |
| IC-003020 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SecurityLelangTrxSubmit:438](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityLelangTrxSubmit.sql#L438) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `IsTrxDeleted` : `BIT` | `0` : `int` | `0` → `BIT` | `IsTrxDeleted = 0` |
| IC-003023 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SecurityLelangTrxSubmit:513](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityLelangTrxSubmit.sql#L513) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `IsTrxDeleted` : `BIT` | `0` : `int` | `0` → `BIT` | `IsTrxDeleted = 0` |
| IC-003116 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SecurityProcessLelangToTransaction:463](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessLelangToTransaction.sql#L463) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mt.IsTrxDeleted` : `BIT` | `1` : `int` | `1` → `BIT` | `mt.IsTrxDeleted = 1` |
| IC-003225 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSAmendSecurityTransaction_TT:992](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L992) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `IsTrxDeleted` : `BIT` | `1` : `int` | `1` → `BIT` | `IsTrxDeleted = 1` |
| IC-003228 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSAmendSecurityTransaction_TT:1009](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L1009) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `IsTrxDeleted` : `BIT` | `0` : `int` | `0` → `BIT` | `IsTrxDeleted = 0` |

<a id="kolom-kursidr"></a>
## KursIDR

**Tipe dasar:** `DECIMAL(25,4)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMateraiTransaction_TT.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001941 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSubmitTransaction:360](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSubmitTransaction.sql#L360) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_tt.KursIDR` : `DECIMAL(25,4)` | `CASE WHEN sm.SecCcy = 'IDR' THEN '1' ELSE '0' END` : `VARCHAR (combined result; facets unknown)` | `CASE WHEN sm.SecCcy = 'IDR' THEN '1' ELSE '0' END` → `DECIMAL(25,4)` | `retail.dbo.trsmateraitransaction_tt.KursIDR <- CASE WHEN sm.SecCcy = 'IDR' THEN '1' ELSE '0' END` |
| IC-007064 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessOverBookingMateraiBuyBack:330](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql#L330) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_tt.KursIDR` : `DECIMAL(25,4)` | `1` : `int` | `1` → `DECIMAL(25,4)` | `retail.dbo.trsmateraitransaction_tt.KursIDR <- 1` |
| IC-012583 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdatePostingResult:655](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L655) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_tt.KursIDR` : `DECIMAL(25,4)` | `1` : `int` | `1` → `DECIMAL(25,4)` | `retail.dbo.trsmateraitransaction_tt.KursIDR <- 1` |

<a id="kolom-materaicost"></a>
## MateraiCost

**Tipe dasar:** `DECIMAL(25,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMateraiTransaction_TT.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 20 lokasi / 9 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002165 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.Rvmp_TRSGetSecurityTransaction_TT:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGetSecurityTransaction_TT.sql#L33) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `ISNULL(mt.MateraiCost, 0)` : `DECIMAL(25,2)` | `@nNominalMaterai` : `money` | `ISNULL(mt.MateraiCost, 0)` → `money` | `@nNominalMaterai = ISNULL(mt.MateraiCost, 0)` |
| IC-003094 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SecurityPopulateTrxLelang:140](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateTrxLelang.sql#L140) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ISNULL(mt.MateraiCost, 0)` : `DECIMAL(25,2)` | `NominalMaterai` : `MONEY` | `ISNULL(mt.MateraiCost, 0)` → `MONEY` | `NominalMaterai = ISNULL(mt.MateraiCost, 0)` |
| IC-003097 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SecurityPopulateTrxLelang:154](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateTrxLelang.sql#L154) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ISNULL(mt.MateraiCost, 0)` : `DECIMAL(25,2)` | `NominalMaterai` : `MONEY` | `ISNULL(mt.MateraiCost, 0)` → `MONEY` | `NominalMaterai = ISNULL(mt.MateraiCost, 0)` |
| IC-004145 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSGetSecurityTransaction_TT:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetSecurityTransaction_TT.sql#L34) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `ISNULL(mt.MateraiCost, 0)` : `DECIMAL(25,2)` | `@nNominalMaterai` : `money` | `ISNULL(mt.MateraiCost, 0)` → `money` | `@nNominalMaterai = ISNULL(mt.MateraiCost, 0)` |
| IC-002164 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSGetSecurityTransaction_TT:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGetSecurityTransaction_TT.sql#L33) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mt.MateraiCost` : `DECIMAL(25,2)` | `0` : `int` | `0` → `DECIMAL(25,2)` | `ISNULL(mt.MateraiCost, 0)` |
| IC-003035 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SecurityPopulateDetailReplaceTrxLelang:32](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateDetailReplaceTrxLelang.sql#L32) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mt.MateraiCost` : `DECIMAL(25,2)` | `0` : `int` | `0` → `DECIMAL(25,2)` | `ISNULL(mt.MateraiCost, 0)` |
| IC-003095 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SecurityPopulateTrxLelang:140](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateTrxLelang.sql#L140) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mt.MateraiCost` : `DECIMAL(25,2)` | `0` : `int` | `0` → `DECIMAL(25,2)` | `ISNULL(mt.MateraiCost, 0)` |
| IC-003098 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SecurityPopulateTrxLelang:154](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateTrxLelang.sql#L154) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mt.MateraiCost` : `DECIMAL(25,2)` | `0` : `int` | `0` → `DECIMAL(25,2)` | `ISNULL(mt.MateraiCost, 0)` |
| IC-004144 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSGetSecurityTransaction_TT:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetSecurityTransaction_TT.sql#L34) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mt.MateraiCost` : `DECIMAL(25,2)` | `0` : `int` | `0` → `DECIMAL(25,2)` | `ISNULL(mt.MateraiCost, 0)` |
| IC-007065 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessOverBookingMateraiBuyBack:330](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql#L330) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_tt.MateraiCost` : `DECIMAL(25,2)` | `bb.Materai` : `MONEY` | `bb.Materai` → `DECIMAL(25,2)` | `retail.dbo.trsmateraitransaction_tt.MateraiCost <- bb.Materai` |
| IC-010086 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptKonfirmasiBOND:419](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L419) | CASE / P11 / COMPARISON | `ISNULL(mt.MateraiCost, 0)` : `DECIMAL(25,2)` | `0` : `int` | `0` → `DECIMAL` | `ISNULL(mt.MateraiCost, 0) > 0` |
| IC-010087 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptKonfirmasiBOND:419](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L419) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mt.MateraiCost` : `DECIMAL(25,2)` | `0` : `int` | `0` → `DECIMAL(25,2)` | `ISNULL(mt.MateraiCost, 0)` |
| IC-010088 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptKonfirmasiBOND:425](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L425) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mt.MateraiCost` : `DECIMAL(25,2)` | `0` : `int` | `0` → `DECIMAL(25,2)` | `ISNULL(mt.MateraiCost, 0)` |
| IC-010141 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptKonfirmasiBONDNew:296](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBONDNew.sql#L296) | CASE / P11 / COMPARISON | `ISNULL(mt.MateraiCost, 0)` : `DECIMAL(25,2)` | `0` : `int` | `0` → `DECIMAL` | `ISNULL(mt.MateraiCost, 0) > 0` |
| IC-010142 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptKonfirmasiBONDNew:296](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBONDNew.sql#L296) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mt.MateraiCost` : `DECIMAL(25,2)` | `0` : `int` | `0` → `DECIMAL(25,2)` | `ISNULL(mt.MateraiCost, 0)` |
| IC-010143 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptKonfirmasiBONDNew:303](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBONDNew.sql#L303) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mt.MateraiCost` : `DECIMAL(25,2)` | `0` : `int` | `0` → `DECIMAL(25,2)` | `ISNULL(mt.MateraiCost, 0)` |
| IC-010203 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptKonfirmasiSRBI:227](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiSRBI.sql#L227) | CASE / P11 / COMPARISON | `ISNULL(mt.MateraiCost, 0)` : `DECIMAL(25,2)` | `0` : `int` | `0` → `DECIMAL` | `ISNULL(mt.MateraiCost, 0) > 0` |
| IC-010204 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptKonfirmasiSRBI:227](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiSRBI.sql#L227) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mt.MateraiCost` : `DECIMAL(25,2)` | `0` : `int` | `0` → `DECIMAL(25,2)` | `ISNULL(mt.MateraiCost, 0)` |
| IC-010205 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptKonfirmasiSRBI:233](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiSRBI.sql#L233) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mt.MateraiCost` : `DECIMAL(25,2)` | `0` : `int` | `0` → `DECIMAL(25,2)` | `ISNULL(mt.MateraiCost, 0)` |
| IC-012584 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdatePostingResult:655](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L655) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_tt.MateraiCost` : `DECIMAL(25,2)` | `tmp.Materai` : `MONEY` | `tmp.Materai` → `DECIMAL(25,2)` | `retail.dbo.trsmateraitransaction_tt.MateraiCost <- tmp.Materai` |

<a id="kolom-transactionamount"></a>
## TransactionAmount

**Tipe dasar:** `DECIMAL(25,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMateraiTransaction_TT.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003021 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SecurityLelangTrxSubmit:442](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityLelangTrxSubmit.sql#L442) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `TransactionAmount` : `DECIMAL(25,2)` | `tmp.TotalProceed` : `[money]` | `tmp.TotalProceed` → `DECIMAL(25,2)` | `TransactionAmount = tmp.TotalProceed` |
| IC-003024 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SecurityLelangTrxSubmit:517](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityLelangTrxSubmit.sql#L517) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `TransactionAmount` : `DECIMAL(25,2)` | `tmp.TotalProceed` : `[money]` | `tmp.TotalProceed` → `DECIMAL(25,2)` | `TransactionAmount = tmp.TotalProceed` |
| IC-003032 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SecurityLelangTrxSubmit:571](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityLelangTrxSubmit.sql#L571) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `TransactionAmount` : `DECIMAL(25,2)` | `tmp.TotalProceed` : `[money]` | `tmp.TotalProceed` → `DECIMAL(25,2)` | `TransactionAmount = tmp.TotalProceed` |
| IC-003115 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SecurityProcessLelangToTransaction:433](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessLelangToTransaction.sql#L433) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mt.TransactionAmount` : `DECIMAL(25,2)` | `tmp.TotalProceed` : `MONEY` | `tmp.TotalProceed` → `DECIMAL(25,2)` | `mt.TransactionAmount = tmp.TotalProceed` |
| IC-003117 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SecurityProcessLelangToTransaction:464](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessLelangToTransaction.sql#L464) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mt.TransactionAmount` : `DECIMAL(25,2)` | `tmp.TotalProceed` : `MONEY` | `tmp.TotalProceed` → `DECIMAL(25,2)` | `mt.TransactionAmount = tmp.TotalProceed` |
| IC-007066 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessOverBookingMateraiBuyBack:330](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql#L330) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_tt.TransactionAmount` : `DECIMAL(25,2)` | `tt.TotalProceed` : `MONEY` | `tt.TotalProceed` → `DECIMAL(25,2)` | `retail.dbo.trsmateraitransaction_tt.TransactionAmount <- tt.TotalProceed` |
| IC-012585 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdatePostingResult:655](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L655) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_tt.TransactionAmount` : `DECIMAL(25,2)` | `tt.TotalProceed` : `MONEY` | `tt.TotalProceed` → `DECIMAL(25,2)` | `retail.dbo.trsmateraitransaction_tt.TransactionAmount <- tt.TotalProceed` |

<a id="kolom-trxstatus"></a>
## TrxStatus

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMateraiTransaction_TT.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001864 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.RepoProsesCancelTrx:341](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoProsesCancelTrx.sql#L341) | EXPRESSION / P07 / CASE_BRANCH | `mt.TrxStatus` : `SMALLINT` | `result type` : `INT (combined result; facets unknown)` | `mt.TrxStatus` → `INT` | `CASE @pbIsAbsorbedByBank WHEN 1 THEN 0 ELSE mt.TrxStatus END` |
| IC-003331 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApproveHapusTransaksiSuratBerharga:196](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveHapusTransaksiSuratBerharga.sql#L196) | EXPRESSION / P07 / CASE_BRANCH | `mt.TrxStatus` : `SMALLINT` | `result type` : `INT (combined result; facets unknown)` | `mt.TrxStatus` → `INT` | `CASE @pbIsAbsorbedByBank WHEN 1 THEN 0 ELSE mt.TrxStatus END` |
| IC-001863 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoProsesCancelTrx:341](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoProsesCancelTrx.sql#L341) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mt.TrxStatus` : `SMALLINT` | `CASE @pbIsAbsorbedByBank WHEN 1 THEN 0 ELSE mt.TrxStatus END` : `INT (combined result; facets unknown)` | `CASE @pbIsAbsorbedByBank WHEN 1 THEN 0 ELSE mt.TrxStatus END` → `SMALLINT` | `mt.TrxStatus = CASE @pbIsAbsorbedByBank WHEN 1 THEN 0 ELSE mt.TrxStatus END` |
| IC-003330 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApproveHapusTransaksiSuratBerharga:196](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveHapusTransaksiSuratBerharga.sql#L196) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mt.TrxStatus` : `SMALLINT` | `CASE @pbIsAbsorbedByBank WHEN 1 THEN 0 ELSE mt.TrxStatus END` : `INT (combined result; facets unknown)` | `CASE @pbIsAbsorbedByBank WHEN 1 THEN 0 ELSE mt.TrxStatus END` → `SMALLINT` | `mt.TrxStatus = CASE @pbIsAbsorbedByBank WHEN 1 THEN 0 ELSE mt.TrxStatus END` |
| IC-007067 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessOverBookingMateraiBuyBack:330](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql#L330) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_tt.TrxStatus` : `SMALLINT` | `0` : `int` | `0` → `SMALLINT` | `retail.dbo.trsmateraitransaction_tt.TrxStatus <- 0` |
| IC-012649 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateStatusProcessTransaksi:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateStatusProcessTransaksi.sql#L169) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mt.TrxStatus` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `mt.TrxStatus = 3` |
