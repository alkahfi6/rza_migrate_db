# Retail UAT.dbo.TRSMaturity_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 36 lokasi source, 7 procedure, 17 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 2 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-approvedby"></a>
## ApprovedBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMaturity_TT.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009675 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptDailyPaymentCouponOrMaturity:473](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L473) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tt.ApprovedBy` : `INT` | `UserApprovalMatching` : `varchar(50)` | `tt.ApprovedBy` → `varchar(50)` | `UserApprovalMatching = tt.ApprovedBy` |
| IC-009687 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptDailyPaymentCouponOrMaturity:499](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L499) | WHERE / P04 / COMPARISON | `isnull(tt.ApprovedBy,'')` : `INT` | `''` : `varchar` | `''` → `INT` | `isnull(tt.ApprovedBy,'') != ''` |
| IC-009688 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptDailyPaymentCouponOrMaturity:499](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L499) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tt.ApprovedBy` : `INT` | `''` : `varchar` | `''` → `INT` | `isnull(tt.ApprovedBy,'')` |

<a id="kolom-capitalgain"></a>
## CapitalGain

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMaturity_TT.sql#L28). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005691 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:1178](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L1178) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `CapitalGain` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(CapitalGain, 0)` |

<a id="kolom-capitalgainlossaftertax"></a>
## CapitalGainLossAfterTax

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMaturity_TT.sql#L25). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005693 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:1180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L1180) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `CapitalGainLossAfterTax` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(CapitalGainLossAfterTax, 0)` |

<a id="kolom-capitalgainlossaftertaxspt"></a>
## CapitalGainLossAfterTaxSPT

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMaturity_TT.sql#L31). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005696 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:1183](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L1183) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `CapitalGainLossAfterTaxSPT` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(CapitalGainLossAfterTaxSPT, 0)` |

<a id="kolom-capitalgainspt"></a>
## CapitalGainSPT

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMaturity_TT.sql#L30). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005695 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:1182](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L1182) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `CapitalGainSPT` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(CapitalGainSPT, 0)` |

<a id="kolom-couponaftertaxafter"></a>
## CouponAfterTaxAfter

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMaturity_TT.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005688 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:1175](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L1175) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `CouponAfterTaxAfter` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(CouponAfterTaxAfter, 0)` |

<a id="kolom-couponamount"></a>
## CouponAmount

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMaturity_TT.sql#L26). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005686 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:1173](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L1173) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `CouponAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(CouponAmount, 0)` |

<a id="kolom-insertedby"></a>
## InsertedBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMaturity_TT.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009665 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptDailyPaymentCouponOrMaturity:453](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L453) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tt.InsertedBy` : `INT` | `UserUploadMatching` : `varchar(50)` | `tt.InsertedBy` → `varchar(50)` | `UserUploadMatching = tt.InsertedBy` |

<a id="kolom-netprincipleamountafter"></a>
## NetPrincipleAmountAfter

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMaturity_TT.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005690 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:1177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L1177) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NetPrincipleAmountAfter` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(NetPrincipleAmountAfter, 0)` |

<a id="kolom-nominalafter"></a>
## NominalAfter

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMaturity_TT.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005685 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:1172](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L1172) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NominalAfter` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(NominalAfter, 0)` |

<a id="kolom-nominalbefore"></a>
## NominalBefore

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMaturity_TT.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-011539 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSettlementMaturity:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettlementMaturity.sql#L23) | EXPRESSION / P07 / CASE_BRANCH | `trsm.NominalBefore` : `MONEY` | `result type` : `DECIMAL (combined result; facets unknown)` | `trsm.NominalBefore` → `DECIMAL` | `case when pbblth.[Status] = 'S' then 0.0 else trsm.NominalBefore end` |

<a id="kolom-roundingafter"></a>
## RoundingAfter

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMaturity_TT.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005694 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:1181](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L1181) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `RoundingAfter` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(RoundingAfter, 0)` |

<a id="kolom-safekeepingaftertaxafter"></a>
## SafeKeepingAfterTaxAfter

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMaturity_TT.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005689 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:1176](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L1176) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `SafeKeepingAfterTaxAfter` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(SafeKeepingAfterTaxAfter, 0)` |

<a id="kolom-secid"></a>
## SecId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMaturity_TT.sql#L3). **Existing leading index:** TRSMaturity_TT_CN1.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009623 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:220](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L220) | JOIN_ON / P11 / COMPARISON | `mtt.SecId` : `INT` | `tmp.SecId` : `bigint` | `mtt.SecId` → `BIGINT` | `tmp.SecId = mtt.SecId` |
| IC-009657 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:437](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L437) | JOIN_ON / P11 / COMPARISON | `tt.SecId` : `INT` | `tmp.SecId` : `bigint` | `tt.SecId` → `BIGINT` | `tmp.SecId = tt.SecId` |
| IC-009666 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:456](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L456) | JOIN_ON / P11 / COMPARISON | `tt.SecId` : `INT` | `tmp.SecId` : `bigint` | `tt.SecId` → `BIGINT` | `tmp.SecId = tt.SecId` |
| IC-009676 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:476](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L476) | JOIN_ON / P11 / COMPARISON | `tt.SecId` : `INT` | `tmp.SecId` : `bigint` | `tt.SecId` → `BIGINT` | `tmp.SecId = tt.SecId` |
| IC-009686 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:497](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L497) | JOIN_ON / P11 / COMPARISON | `tt.SecId` : `INT` | `tmp.SecId` : `bigint` | `tt.SecId` → `BIGINT` | `tmp.SecId = tt.SecId` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMaturity_TT.sql#L19). **Existing leading index:** TRSMaturity_TT_NN1.

**Frekuensi source:** 14 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002012 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSGenerateDumpOBL058:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGenerateDumpOBL058.sql#L216) | CASE / P12 / COMPARISON | `mtt.Status` : `CHAR(1)` | `'P'` : `varchar` | `mtt.Status` → `VARCHAR` | `mtt.Status = 'P'` |
| IC-002013 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSGenerateDumpOBL058:217](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGenerateDumpOBL058.sql#L217) | CASE / P12 / COMPARISON | `mtt.Status` : `CHAR(1)` | `'A'` : `varchar` | `mtt.Status` → `VARCHAR` | `mtt.Status = 'A'` |
| IC-002014 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSGenerateDumpOBL058:218](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGenerateDumpOBL058.sql#L218) | CASE / P12 / COMPARISON | `mtt.Status` : `CHAR(1)` | `'R'` : `varchar` | `mtt.Status` → `VARCHAR` | `mtt.Status = 'R'` |
| IC-002069 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSGenerateDumpOBL058_New:219](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGenerateDumpOBL058_New.sql#L219) | CASE / P12 / COMPARISON | `mtt.Status` : `CHAR(1)` | `'P'` : `varchar` | `mtt.Status` → `VARCHAR` | `mtt.Status = 'P'` |
| IC-002070 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSGenerateDumpOBL058_New:220](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGenerateDumpOBL058_New.sql#L220) | CASE / P12 / COMPARISON | `mtt.Status` : `CHAR(1)` | `'A'` : `varchar` | `mtt.Status` → `VARCHAR` | `mtt.Status = 'A'` |
| IC-002071 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSGenerateDumpOBL058_New:221](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGenerateDumpOBL058_New.sql#L221) | CASE / P12 / COMPARISON | `mtt.Status` : `CHAR(1)` | `'R'` : `varchar` | `mtt.Status` → `VARCHAR` | `mtt.Status = 'R'` |
| IC-002265 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSMaturityProcess:612](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityProcess.sql#L612) | WHERE / P12 / IN_VALUE_TYPE | `mtt.Status` : `CHAR(1)` | `'P'` : `varchar` | `mtt.Status` → `VARCHAR` | `mtt.Status in ('P','R')` |
| IC-005622 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateMaturityUpload:601](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L601) | CASE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'R'` : `varchar` | `Status` → `VARCHAR` | `Status = 'R'` |
| IC-005623 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateMaturityUpload:602](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L602) | CASE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'A'` : `varchar` | `Status` → `VARCHAR` | `Status = 'A'` |
| IC-005624 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateMaturityUpload:603](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L603) | CASE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'P'` : `varchar` | `Status` → `VARCHAR` | `Status = 'P'` |
| IC-005697 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateMaturityUpload:1186](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L1186) | WHERE / P12 / IN_VALUE_TYPE | `Status` : `CHAR(1)` | `'P'` : `varchar` | `Status` → `VARCHAR` | `Status IN ('P', 'A')` |
| IC-009677 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:478](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L478) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'A'` : `varchar` | `Status` → `VARCHAR` | `Status = 'A'` |
| IC-007025 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessMaturityUploadAuthorization:175](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessMaturityUploadAuthorization.sql#L175) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mtt.Status` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `mtt.Status = 'A'` |
| IC-007030 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessMaturityUploadAuthorization:296](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessMaturityUploadAuthorization.sql#L296) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mtt.Status` : `CHAR(1)` | `'R'` : `varchar` | `'R'` → `CHAR(1)` | `mtt.Status = 'R'` |

<a id="kolom-taxoncapital"></a>
## TaxOnCapital

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMaturity_TT.sql#L29). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005692 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:1179](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L1179) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TaxOnCapital` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(TaxOnCapital, 0)` |

<a id="kolom-taxoncoupon"></a>
## TaxOnCoupon

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMaturity_TT.sql#L27). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005687 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:1174](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L1174) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TaxOnCoupon` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(TaxOnCoupon, 0)` |
