# Retail UAT.dbo.TRSMatchingCoupon_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 44 lokasi source, 9 procedure, 9 kolom. **A:** 1 lokasi pada 1 procedure. **B:** 2 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-norekening"></a>
## NoRekening

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMatchingCoupon_TT.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009562 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptBatchTransaction:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptBatchTransaction.sql#L65) | JOIN_ON / P01 / COMPARISON | `tmct.NoRekening` : `VARCHAR(20)` | `dvc.ACCTNO` : `DECIMAL(19,0)` | `tmct.NoRekening` → `DECIMAL` | `tmct.NoRekening = dvc.ACCTNO` |

<a id="kolom-approvedby"></a>
## ApprovedBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMatchingCoupon_TT.sql#L30). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009669 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptDailyPaymentCouponOrMaturity:462](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L462) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tt.ApprovedBy` : `INT` | `UserApprovalMatching` : `varchar(50)` | `tt.ApprovedBy` → `varchar(50)` | `UserApprovalMatching = tt.ApprovedBy` |
| IC-009681 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptDailyPaymentCouponOrMaturity:488](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L488) | WHERE / P04 / COMPARISON | `isnull(tt.ApprovedBy,'')` : `INT` | `''` : `varchar` | `''` → `INT` | `isnull(tt.ApprovedBy,'') != ''` |
| IC-009682 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptDailyPaymentCouponOrMaturity:488](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L488) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tt.ApprovedBy` : `INT` | `''` : `varchar` | `''` → `INT` | `isnull(tt.ApprovedBy,'')` |

<a id="kolom-coupondate"></a>
## CouponDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMatchingCoupon_TT.sql#L33). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006268 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementKupon:231](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementKupon.sql#L231) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mc.CouponDate` : `DATETIME` | `CONVERT(DATE, sm.LastCouponDate)` : `DATE` | `CONVERT(DATE, sm.LastCouponDate)` → `DATETIME` | `mc.CouponDate = CONVERT(DATE, sm.LastCouponDate)` |

<a id="kolom-insertedby"></a>
## InsertedBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMatchingCoupon_TT.sql#L28). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009660 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptDailyPaymentCouponOrMaturity:443](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L443) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tt.InsertedBy` : `INT` | `UserUploadMatching` : `varchar(50)` | `tt.InsertedBy` → `varchar(50)` | `UserUploadMatching = tt.InsertedBy` |

<a id="kolom-ismigration"></a>
## IsMigration

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMatchingCoupon_TT.sql#L41). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002201 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSKuponProcess:260](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L260) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmatchingcoupon_tt.IsMigration` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.trsmatchingcoupon_tt.IsMigration <- 0` |

<a id="kolom-issettle"></a>
## IsSettle

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMatchingCoupon_TT.sql#L34). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006270 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateSettlementKupon:237](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementKupon.sql#L237) | WHERE / P11 / COMPARISON | `ISNULL(mc.IsSettle, 0)` : `BIT` | `0` : `int` | `ISNULL(mc.IsSettle, 0)` → `INT` | `ISNULL(mc.IsSettle, 0) = 0` |
| IC-006283 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateSettlementKupon:550](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementKupon.sql#L550) | WHERE / P11 / COMPARISON | `ISNULL(mc.IsSettle, 0)` : `BIT` | `0` : `int` | `ISNULL(mc.IsSettle, 0)` → `INT` | `ISNULL(mc.IsSettle, 0) = 0` |
| IC-009563 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptBatchTransaction:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptBatchTransaction.sql#L66) | WHERE / P11 / COMPARISON | `tmct.IsSettle` : `BIT` | `1` : `int` | `tmct.IsSettle` → `INT` | `tmct.IsSettle = 1` |
| IC-002202 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSKuponProcess:260](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L260) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmatchingcoupon_tt.IsSettle` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.trsmatchingcoupon_tt.IsSettle <- 0` |
| IC-006271 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateSettlementKupon:237](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementKupon.sql#L237) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mc.IsSettle` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(mc.IsSettle, 0)` |
| IC-006284 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateSettlementKupon:550](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementKupon.sql#L550) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mc.IsSettle` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(mc.IsSettle, 0)` |
| IC-007499 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementKupon:274](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementKupon.sql#L274) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mc.IsSettle` : `BIT` | `1` : `int` | `1` → `BIT` | `mc.IsSettle = 1` |

<a id="kolom-secid"></a>
## SecId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMatchingCoupon_TT.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002196 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSKuponProcess:245](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L245) | WHERE / P11 / COMPARISON | `mc.SecId` : `INT` | `@nSecId` : `BIGINT` | `mc.SecId` → `BIGINT` | `mc.SecId = @nSecId` |
| IC-009649 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:405](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L405) | JOIN_ON / P11 / COMPARISON | `mtt.SecId` : `INT` | `tmp.SecId` : `bigint` | `mtt.SecId` → `BIGINT` | `tmp.SecId = mtt.SecId` |
| IC-009653 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:427](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L427) | JOIN_ON / P11 / COMPARISON | `tt.SecId` : `INT` | `tmp.SecId` : `bigint` | `tt.SecId` → `BIGINT` | `tmp.SecId = tt.SecId` |
| IC-009661 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:446](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L446) | JOIN_ON / P11 / COMPARISON | `tt.SecId` : `INT` | `tmp.SecId` : `bigint` | `tt.SecId` → `BIGINT` | `tmp.SecId = tt.SecId` |
| IC-009670 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:465](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L465) | JOIN_ON / P11 / COMPARISON | `tt.SecId` : `INT` | `tmp.SecId` : `bigint` | `tt.SecId` → `BIGINT` | `tmp.SecId = tt.SecId` |
| IC-009680 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:486](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L486) | JOIN_ON / P11 / COMPARISON | `tt.SecId` : `INT` | `tmp.SecId` : `bigint` | `tt.SecId` → `BIGINT` | `tmp.SecId = tt.SecId` |
| IC-002203 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSKuponProcess:260](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L260) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmatchingcoupon_tt.SecId` : `INT` | `ps.sec_id` : `BIGINT` | `ps.sec_id` → `INT` | `retail.dbo.trsmatchingcoupon_tt.SecId <- ps.sec_id` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMatchingCoupon_TT.sql#L27). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 22 lokasi / 7 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002146 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSGenerateDumpOBL084:271](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGenerateDumpOBL084.sql#L271) | CASE / P12 / COMPARISON | `mtt.[Status]` : `CHAR(1)` | `'P'` : `varchar` | `mtt.[Status]` → `VARCHAR` | `mtt.[Status] = 'P'` |
| IC-002147 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSGenerateDumpOBL084:272](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGenerateDumpOBL084.sql#L272) | CASE / P12 / COMPARISON | `mtt.[Status]` : `CHAR(1)` | `'A'` : `varchar` | `mtt.[Status]` → `VARCHAR` | `mtt.[Status] = 'A'` |
| IC-002148 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSGenerateDumpOBL084:273](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGenerateDumpOBL084.sql#L273) | CASE / P12 / COMPARISON | `mtt.[Status]` : `CHAR(1)` | `'R'` : `varchar` | `mtt.[Status]` → `VARCHAR` | `mtt.[Status] = 'R'` |
| IC-002197 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSKuponProcess:247](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L247) | WHERE / P12 / IN_VALUE_TYPE | `mc.[Status]` : `CHAR(1)` | `'P'` : `varchar` | `mc.[Status]` → `VARCHAR` | `mc.[Status] in ('P','R')` |
| IC-005475 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateKuponUpload:367](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUpload.sql#L367) | CASE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'R'` : `varchar` | `Status` → `VARCHAR` | `Status = 'R'` |
| IC-005476 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateKuponUpload:368](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUpload.sql#L368) | CASE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'A'` : `varchar` | `Status` → `VARCHAR` | `Status = 'A'` |
| IC-005477 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateKuponUpload:369](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUpload.sql#L369) | CASE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'P'` : `varchar` | `Status` → `VARCHAR` | `Status = 'P'` |
| IC-005496 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateKuponUpload:574](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUpload.sql#L574) | CASE / P12 / COMPARISON | `mtt.Status` : `CHAR(1)` | `'A'` : `varchar` | `mtt.Status` → `VARCHAR` | `mtt.Status = 'A'` |
| IC-005497 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateKuponUpload:577](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUpload.sql#L577) | CASE / P12 / COMPARISON | `mtt.Status` : `CHAR(1)` | `'R'` : `varchar` | `mtt.Status` → `VARCHAR` | `mtt.Status = 'R'` |
| IC-005498 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateKuponUpload:580](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUpload.sql#L580) | CASE / P12 / COMPARISON | `mtt.Status` : `CHAR(1)` | `'P'` : `varchar` | `mtt.Status` → `VARCHAR` | `mtt.Status = 'P'` |
| IC-005556 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateKuponUploadAuthorization:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUploadAuthorization.sql#L45) | WHERE / P12 / COMPARISON | `mtt.Status` : `CHAR(1)` | `'P'` : `varchar` | `mtt.Status` → `VARCHAR` | `mtt.Status = 'P'` |
| IC-006269 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementKupon:236](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementKupon.sql#L236) | WHERE / P12 / COMPARISON | `mc.Status` : `CHAR(1)` | `'A'` : `varchar` | `mc.Status` → `VARCHAR` | `mc.Status = 'A'` |
| IC-006272 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementKupon:269](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementKupon.sql#L269) | WHERE / P12 / COMPARISON | `mc.Status` : `CHAR(1)` | `'A'` : `varchar` | `mc.Status` → `VARCHAR` | `mc.Status = 'A'` |
| IC-006282 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementKupon:546](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementKupon.sql#L546) | WHERE / P12 / COMPARISON | `mc.Status` : `CHAR(1)` | `'A'` : `varchar` | `mc.Status` → `VARCHAR` | `mc.Status = 'A'` |
| IC-006968 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessKuponUploadAuthorization:93](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessKuponUploadAuthorization.sql#L93) | WHERE / P12 / COMPARISON | `mtt.Status` : `CHAR(1)` | `'P'` : `varchar` | `mtt.Status` → `VARCHAR` | `mtt.Status = 'P'` |
| IC-006969 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessKuponUploadAuthorization:109](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessKuponUploadAuthorization.sql#L109) | JOIN_ON / P12 / COMPARISON | `mtt.Status` : `CHAR(1)` | `'A'` : `varchar` | `mtt.Status` → `VARCHAR` | `mtt.Status = 'A'` |
| IC-006972 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessKuponUploadAuthorization:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessKuponUploadAuthorization.sql#L130) | JOIN_ON / P12 / COMPARISON | `mtt.Status` : `CHAR(1)` | `'P'` : `varchar` | `mtt.Status` → `VARCHAR` | `mtt.Status = 'P'` |
| IC-006977 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessKuponUploadAuthorization:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessKuponUploadAuthorization.sql#L225) | JOIN_ON / P12 / COMPARISON | `mtt.Status` : `CHAR(1)` | `'P'` : `varchar` | `mtt.Status` → `VARCHAR` | `mtt.Status = 'P'` |
| IC-009671 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:467](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L467) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'A'` : `varchar` | `Status` → `VARCHAR` | `Status = 'A'` |
| IC-002204 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSKuponProcess:260](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L260) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmatchingcoupon_tt.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trsmatchingcoupon_tt.Status <- 'P'` |
| IC-006971 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessKuponUploadAuthorization:122](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessKuponUploadAuthorization.sql#L122) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Status` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `Status = 'A'` |
| IC-006976 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessKuponUploadAuthorization:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessKuponUploadAuthorization.sql#L216) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Status` : `CHAR(1)` | `'R'` : `varchar` | `'R'` → `CHAR(1)` | `Status = 'R'` |

<a id="kolom-templatetype"></a>
## TemplateType

**Tipe dasar:** `VARCHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMatchingCoupon_TT.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002205 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSKuponProcess:260](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L260) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmatchingcoupon_tt.TemplateType` : `VARCHAR(1)` | `ps.TemplateType` : `CHAR(1)` | `ps.TemplateType` → `VARCHAR(1)` | `retail.dbo.trsmatchingcoupon_tt.TemplateType <- ps.TemplateType` |
