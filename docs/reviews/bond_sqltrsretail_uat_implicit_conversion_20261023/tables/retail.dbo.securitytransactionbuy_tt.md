# Retail UAT.dbo.SecurityTransactionBuy_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 56 lokasi source, 8 procedure, 12 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-dealid"></a>
## DealId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionBuy_TT.sql#L3). **Existing leading index:** ix_SecurityTransactionBuy_TT.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004400 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertSecurityTransaction_TT_Original:2336](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L2336) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransactionbuy_tt.DealId` : `BIGINT` | `@nIdentity` : `int` | `@nIdentity` → `BIGINT` | `retail.dbo.securitytransactionbuy_tt.DealId <- @nIdentity` |

<a id="kolom-dealprice"></a>
## DealPrice

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionBuy_TT.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007057 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessOverBookingMateraiBuyBack:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql#L265) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransactionbuy_tt.DealPrice` : `DECIMAL(20,5)` | `BuyBackPrice` : `MONEY` | `BuyBackPrice` → `DECIMAL(20,5)` | `retail.dbo.securitytransactionbuy_tt.DealPrice <- BuyBackPrice` |
| IC-012575 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdatePostingResult:584](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L584) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransactionbuy_tt.DealPrice` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL(20,5)` | `retail.dbo.securitytransactionbuy_tt.DealPrice <- 100` |
| IC-012606 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdatePostingResult:952](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L952) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransactionbuy_tt.DealPrice` : `DECIMAL(20,5)` | `100` : `int` | `100` → `DECIMAL(20,5)` | `retail.dbo.securitytransactionbuy_tt.DealPrice <- 100` |

<a id="kolom-facevalue"></a>
## FaceValue

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionBuy_TT.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-010468 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:211](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L211) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.FaceValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(stb.FaceValue,0)` |
| IC-010489 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:317](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L317) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.FaceValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(stb.FaceValue,0)` |
| IC-010510 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:455](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L455) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.FaceValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(stb.FaceValue,0)` |
| IC-010531 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:557](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L557) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.FaceValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(stb.FaceValue,0)` |
| IC-010557 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo:118](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo.sql#L118) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.FaceValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(stb.FaceValue,0)` |
| IC-010580 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo:208](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo.sql#L208) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.FaceValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(stb.FaceValue,0)` |

<a id="kolom-safekeepingfeeaftertax"></a>
## SafeKeepingFeeAfterTax

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionBuy_TT.sql#L24). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007594 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSettlementTransaksiBankBeli:387](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L387) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sttt.SafeKeepingFeeAfterTax` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(sttt.SafeKeepingFeeAfterTax,0)` |
| IC-007614 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSettlementTransaksiBankBeli:469](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L469) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sttt.SafeKeepingFeeAfterTax` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(sttt.SafeKeepingFeeAfterTax,0)` |

<a id="kolom-signature"></a>
## Signature

**Tipe dasar:** `VARBINARY(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionBuy_TT.sql#L43). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004401 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertSecurityTransaction_TT_Original:2336](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L2336) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransactionbuy_tt.Signature` : `VARBINARY(MAX)` | `temp.[Signature]` : `image` | `temp.[Signature]` → `VARBINARY(MAX)` | `retail.dbo.securitytransactionbuy_tt.Signature <- temp.[Signature]` |

<a id="kolom-summary"></a>
## Summary

**Tipe dasar:** `VARBINARY(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionBuy_TT.sql#L45). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004402 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertSecurityTransaction_TT_Original:2336](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L2336) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransactionbuy_tt.Summary` : `VARBINARY(MAX)` | `temp.Summary` : `image` | `temp.Summary` → `VARBINARY(MAX)` | `retail.dbo.securitytransactionbuy_tt.Summary <- temp.Summary` |

<a id="kolom-taxonaccrued"></a>
## TaxOnAccrued

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionBuy_TT.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007579 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSettlementTransaksiBankBeli:344](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L344) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sttt.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(sttt.TaxOnAccrued,0)` |
| IC-007592 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSettlementTransaksiBankBeli:386](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L386) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sttt.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(sttt.TaxOnAccrued,0)` |
| IC-007599 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSettlementTransaksiBankBeli:426](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L426) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sttt.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(sttt.TaxOnAccrued,0)` |
| IC-007612 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSettlementTransaksiBankBeli:468](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L468) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sttt.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(sttt.TaxOnAccrued,0)` |

<a id="kolom-taxoncapitalgainloss"></a>
## TaxOnCapitalGainLoss

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionBuy_TT.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007580 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSettlementTransaksiBankBeli:344](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L344) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sttt.TaxOnCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(sttt.TaxOnCapitalGainLoss,0)` |
| IC-007593 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSettlementTransaksiBankBeli:386](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L386) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sttt.TaxOnCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(sttt.TaxOnCapitalGainLoss,0)` |
| IC-007600 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSettlementTransaksiBankBeli:426](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L426) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sttt.TaxOnCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(sttt.TaxOnCapitalGainLoss,0)` |
| IC-007613 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSettlementTransaksiBankBeli:468](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L468) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sttt.TaxOnCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(sttt.TaxOnCapitalGainLoss,0)` |

<a id="kolom-transactionfeeamount"></a>
## TransactionFeeAmount

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionBuy_TT.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007595 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSettlementTransaksiBankBeli:387](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L387) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sttt.TransactionFeeAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(sttt.TransactionFeeAmount,0)` |
| IC-007615 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSettlementTransaksiBankBeli:469](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L469) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sttt.TransactionFeeAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(sttt.TransactionFeeAmount,0)` |

<a id="kolom-trxstatus"></a>
## TrxStatus

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionBuy_TT.sql#L26). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007058 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessOverBookingMateraiBuyBack:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql#L265) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransactionbuy_tt.TrxStatus` : `SMALLINT` | `0` : `int` | `0` → `SMALLINT` | `retail.dbo.securitytransactionbuy_tt.TrxStatus <- 0` |
| IC-012576 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdatePostingResult:584](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L584) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransactionbuy_tt.TrxStatus` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `retail.dbo.securitytransactionbuy_tt.TrxStatus <- 3` |
| IC-012607 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdatePostingResult:952](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L952) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransactionbuy_tt.TrxStatus` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `retail.dbo.securitytransactionbuy_tt.TrxStatus <- 3` |

<a id="kolom-trxtype"></a>
## TrxType

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionBuy_TT.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007059 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessOverBookingMateraiBuyBack:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql#L265) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransactionbuy_tt.TrxType` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `retail.dbo.securitytransactionbuy_tt.TrxType <- 3` |
| IC-012577 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdatePostingResult:584](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L584) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransactionbuy_tt.TrxType` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `retail.dbo.securitytransactionbuy_tt.TrxType <- 3` |
| IC-012608 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdatePostingResult:952](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L952) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransactionbuy_tt.TrxType` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `retail.dbo.securitytransactionbuy_tt.TrxType <- 3` |

<a id="kolom-istaxamnesty"></a>
## isTaxAmnesty

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionBuy_TT.sql#L42). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 26 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007581 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSProcessSettlementTransaksiBankBeli:351](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L351) | CASE / P11 / COMPARISON | `isnull(sttt.isTaxAmnesty,0)` : `BIT` | `1` : `int` | `isnull(sttt.isTaxAmnesty,0)` → `INT` | `isnull(sttt.isTaxAmnesty,0) = 1` |
| IC-007601 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSProcessSettlementTransaksiBankBeli:433](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L433) | CASE / P11 / COMPARISON | `isnull(sttt.isTaxAmnesty,0)` : `BIT` | `1` : `int` | `isnull(sttt.isTaxAmnesty,0)` → `INT` | `isnull(sttt.isTaxAmnesty,0) = 1` |
| IC-010570 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo.sql#L156) | WHERE / P11 / COMPARISON | `stb.isTaxAmnesty` : `BIT` | `1` : `int` | `stb.isTaxAmnesty` → `INT` | `stb.isTaxAmnesty = 1` |
| IC-010593 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo:247](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo.sql#L247) | WHERE / P11 / COMPARISON | `stb.isTaxAmnesty` : `BIT` | `1` : `int` | `stb.isTaxAmnesty` → `INT` | `stb.isTaxAmnesty = 1` |
| IC-007582 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSettlementTransaksiBankBeli:351](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L351) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sttt.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(sttt.isTaxAmnesty,0)` |
| IC-007602 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSettlementTransaksiBankBeli:433](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L433) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sttt.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(sttt.isTaxAmnesty,0)` |
| IC-007788 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProrateTrxTaxAmesty:55](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProrateTrxTaxAmesty.sql#L55) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `st3.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(st3.isTaxAmnesty,0)` |
| IC-010458 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:179](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L179) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(stb.isTaxAmnesty,0)` |
| IC-010473 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:219](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L219) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(stb.isTaxAmnesty,0)` |
| IC-010477 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:253](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L253) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(stb.isTaxAmnesty,0)` |
| IC-010479 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:285](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L285) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(stb.isTaxAmnesty,0)` |
| IC-010494 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:325](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L325) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(stb.isTaxAmnesty,0)` |
| IC-010498 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:359](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L359) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(stb.isTaxAmnesty,0)` |
| IC-010500 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:423](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L423) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(stb.isTaxAmnesty,0)` |
| IC-010515 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:463](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L463) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(stb.isTaxAmnesty,0)` |
| IC-010519 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:494](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L494) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(stb.isTaxAmnesty,0)` |
| IC-010521 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:525](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L525) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(stb.isTaxAmnesty,0)` |
| IC-010536 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:565](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L565) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(stb.isTaxAmnesty,0)` |
| IC-010540 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingPerSecAccNo:596](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql#L596) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(stb.isTaxAmnesty,0)` |
| IC-010551 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo:95](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo.sql#L95) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(stb.isTaxAmnesty,0)` |
| IC-010564 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo.sql#L131) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(stb.isTaxAmnesty,0)` |
| IC-010572 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo.sql#L162) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(stb.isTaxAmnesty,0)` |
| IC-010574 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo:185](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo.sql#L185) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(stb.isTaxAmnesty,0)` |
| IC-010587 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo:221](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo.sql#L221) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(stb.isTaxAmnesty,0)` |
| IC-010595 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo:253](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo.sql#L253) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stb.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(stb.isTaxAmnesty,0)` |
| IC-012660 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateTaxAmnestyExpired:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateTaxAmnestyExpired.sql#L41) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `isTaxAmnesty = 0` |
