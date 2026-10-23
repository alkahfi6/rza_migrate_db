# Retail UAT.dbo.SecurityTransactionLink_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 54 lokasi source, 14 procedure, 6 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 6 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-accruedinterest"></a>
## AccruedInterest

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionLink_TR.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008392 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportSKFBankBeliBKON:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeliBKON.sql#L44) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tl.AccruedInterest` : `MONEY` | `#tempResult.AccruedInterest` : `DECIMAL(18,2)` | `tl.AccruedInterest` → `DECIMAL(18,2)` | `#tempResult.AccruedInterest <- tl.AccruedInterest` |
| IC-010224 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptLaporanPajakBankBuy:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L41) | EXPRESSION / P07 / CASE_BRANCH | `(stl.AccruedInterest)` : `MONEY` | `result type` : `DECIMAL (combined result; facets unknown)` | `(stl.AccruedInterest)` → `DECIMAL` | `case when stttSell.SettlementDate < dbo.fnOffsetLastCouponDate(smtm.SecId, stttSell.SettlementDate) then (stl.AccruedInterest) else ( convert (decimal (20,2), '0.00') ) END` |
| IC-010266 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptLaporanPajakBankBuy:270](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L270) | EXPRESSION / P07 / CASE_BRANCH | `(stl.AccruedInterest)` : `MONEY` | `result type` : `DECIMAL (combined result; facets unknown)` | `(stl.AccruedInterest)` → `DECIMAL` | `case when stttSell.SettlementDate < dbo.fnOffsetLastCouponDate(smtm.SecId, stttSell.SettlementDate) then (stl.AccruedInterest) else ( convert (decimal (20,2), '0.00') ) END` |

<a id="kolom-dealidchild"></a>
## DealIdChild

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionLink_TR.sql#L3). **Existing leading index:** SecurityTransactionLink_TR_NN2.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000155 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.CustodyReportMaintainAcquisition:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportMaintainAcquisition.sql#L125) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `DealIdChild` : `BIGINT` | `#tmp_Sell.SellFrom` : `VARCHAR(50)` | `DealIdChild` → `VARCHAR(50)` | `#tmp_Sell.SellFrom <- DealIdChild` |

<a id="kolom-nominal"></a>
## Nominal

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionLink_TR.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 33 lokasi / 12 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000154 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.CustodyReportMaintainAcquisition:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportMaintainAcquisition.sql#L125) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sum(Nominal)` : `MONEY (aggregate; facets unknown)` | `#tmp_Sell.Nominal` : `DECIMAL(30,2)` | `sum(Nominal)` → `DECIMAL(30,2)` | `#tmp_Sell.Nominal <- sum(Nominal)` |
| IC-008395 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportSKFBankBeliBKON:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeliBKON.sql#L44) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tl.Nominal` : `MONEY` | `#tempResult.Nominal` : `DECIMAL(18,2)` | `tl.Nominal` → `DECIMAL(18,2)` | `#tempResult.Nominal <- tl.Nominal` |
| IC-009106 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRetailReCalculateAccrued:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailReCalculateAccrued.sql#L53) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `isnull(stl.Nominal,0)` : `MONEY` | `(stttParent.DealPrice - stttChild.DealPrice) / 100.00` : `DECIMAL (combined result; facets unknown)` | `isnull(stl.Nominal,0)` → `DECIMAL` | `(stttParent.DealPrice - stttChild.DealPrice) / 100.00 * isnull(stl.Nominal,0)` |
| IC-010220 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptLaporanPajakBankBuy:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L31) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `stl.Nominal` : `MONEY` | `stttSell.DealPrice` : `DECIMAL(20,5)` | `stl.Nominal` → `DECIMAL` | `stl.Nominal * stttSell.DealPrice` |
| IC-010221 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptLaporanPajakBankBuy:35](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L35) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `stl.Nominal` : `MONEY` | `stttBuy.DealPrice` : `DECIMAL(20,5)` | `stl.Nominal` → `DECIMAL` | `stl.Nominal * stttBuy.DealPrice` |
| IC-010222 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptLaporanPajakBankBuy:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L36) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `stl.Nominal` : `MONEY` | `stttBuy.DealPrice` : `DECIMAL(20,5)` | `stl.Nominal` → `DECIMAL` | `stl.Nominal * stttBuy.DealPrice` |
| IC-010223 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptLaporanPajakBankBuy:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L36) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `stl.Nominal` : `MONEY` | `stttSell.DealPrice` : `DECIMAL(20,5)` | `stl.Nominal` → `DECIMAL` | `stl.Nominal * stttSell.DealPrice` |
| IC-010262 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptLaporanPajakBankBuy:260](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L260) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `stl.Nominal` : `MONEY` | `stttSell.DealPrice` : `DECIMAL(20,5)` | `stl.Nominal` → `DECIMAL` | `stl.Nominal * stttSell.DealPrice` |
| IC-010263 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptLaporanPajakBankBuy:264](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L264) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `stl.Nominal` : `MONEY` | `stttBuy.DealPrice` : `DECIMAL(20,5)` | `stl.Nominal` → `DECIMAL` | `stl.Nominal * stttBuy.DealPrice` |
| IC-010264 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptLaporanPajakBankBuy:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L265) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `stl.Nominal` : `MONEY` | `stttBuy.DealPrice` : `DECIMAL(20,5)` | `stl.Nominal` → `DECIMAL` | `stl.Nominal * stttBuy.DealPrice` |
| IC-010265 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptLaporanPajakBankBuy:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L265) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `stl.Nominal` : `MONEY` | `stttSell.DealPrice` : `DECIMAL(20,5)` | `stl.Nominal` → `DECIMAL` | `stl.Nominal * stttSell.DealPrice` |
| IC-011461 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptProductWM:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptProductWM.sql#L34) | CASE / P05 / BINARY_Multiply | `isnull(lnk.Nominal,st.FaceValue)` : `MONEY` | `(st.DealPrice - st1.DealPrice)` : `DECIMAL (combined result; facets unknown)` | `isnull(lnk.Nominal,st.FaceValue)` → `DECIMAL` | `isnull(lnk.Nominal,st.FaceValue) * (st.DealPrice - st1.DealPrice)` |
| IC-011464 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptProductWM:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptProductWM.sql#L41) | CASE / P05 / BINARY_Multiply | `isnull(lnk.Nominal,st.FaceValue)` : `MONEY` | `(st1.DealPrice - st.DealPrice + bond.Bidprice)` : `DECIMAL (combined result; facets unknown)` | `isnull(lnk.Nominal,st.FaceValue)` → `DECIMAL` | `(st1.DealPrice - st.DealPrice + bond.Bidprice) * isnull(lnk.Nominal,st.FaceValue)` |
| IC-011465 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptProductWM:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptProductWM.sql#L43) | CASE / P05 / BINARY_Multiply | `isnull(lnk.Nominal,st.FaceValue)` : `MONEY` | `(st1.DealPrice - st.DealPrice + bond.Bidprice)` : `DECIMAL (combined result; facets unknown)` | `isnull(lnk.Nominal,st.FaceValue)` → `DECIMAL` | `(st1.DealPrice - st.DealPrice + bond.Bidprice) * isnull(lnk.Nominal,st.FaceValue)` |
| IC-000796 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.OMTrxSecPreBook:849](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMTrxSecPreBook.sql#L849) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `Nominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(Nominal, 0)` |
| IC-004090 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSGenerateSafeKeepingFeeNasabah:264](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeNasabah.sql#L264) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stltr.Nominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(stltr.Nominal, 0)` |
| IC-004091 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSGenerateSafeKeepingFeeNasabah:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeNasabah.sql#L265) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stltr.Nominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(stltr.Nominal, 0)` |
| IC-004301 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSInsertSecurityTransaction_TT_Original:512](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L512) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `Nominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(Nominal, 0)` |
| IC-006460 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateTransactionBalanceForUpdate:79](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTransactionBalanceForUpdate.sql#L79) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `resultTable.SumNominal` : `MONEY (aggregate; facets unknown)` | `0` : `int` | `0` → `MONEY (aggregate; facets unknown)` | `isnull(resultTable.SumNominal,0)` |
| IC-006463 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateTransactionBalanceForUpdate:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTransactionBalanceForUpdate.sql#L80) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `resultTable.SumNominal` : `MONEY (aggregate; facets unknown)` | `0` : `int` | `0` → `MONEY (aggregate; facets unknown)` | `isnull(resultTable.SumNominal,0)` |
| IC-009107 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRetailReCalculateAccrued:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailReCalculateAccrued.sql#L53) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stl.Nominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(stl.Nominal,0)` |
| IC-010217 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptLaporanPajakBankBuy:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L23) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stl.Nominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(stl.Nominal,0)` |
| IC-010218 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptLaporanPajakBankBuy:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L29) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `stl.Nominal` : `MONEY` | `1000000` : `int` | `1000000` → `MONEY` | `stl.Nominal/1000000` |
| IC-010219 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptLaporanPajakBankBuy:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L29) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `(stl.Nominal/1000000)` : `MONEY (combined result; facets unknown)` | `0` : `int` | `0` → `MONEY (combined result; facets unknown)` | `isnull((stl.Nominal/1000000),0)` |
| IC-010259 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptLaporanPajakBankBuy:252](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L252) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stl.Nominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(stl.Nominal,0)` |
| IC-010260 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptLaporanPajakBankBuy:258](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L258) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `stl.Nominal` : `MONEY` | `1000000` : `int` | `1000000` → `MONEY` | `stl.Nominal/1000000` |
| IC-010261 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptLaporanPajakBankBuy:258](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L258) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `(stl.Nominal/1000000)` : `MONEY (combined result; facets unknown)` | `0` : `int` | `0` → `MONEY (combined result; facets unknown)` | `isnull((stl.Nominal/1000000),0)` |
| IC-013216 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_InsertTrx2ndSecurityTransaction_TT:367](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTrx2ndSecurityTransaction_TT.sql#L367) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `resultTable.SumNominal` : `MONEY (aggregate; facets unknown)` | `0` : `int` | `0` → `MONEY (aggregate; facets unknown)` | `isnull(resultTable.SumNominal,0)` |
| IC-013219 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_InsertTrx2ndSecurityTransaction_TT:368](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTrx2ndSecurityTransaction_TT.sql#L368) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `resultTable.SumNominal` : `MONEY (aggregate; facets unknown)` | `0` : `int` | `0` → `MONEY (aggregate; facets unknown)` | `isnull(resultTable.SumNominal,0)` |
| IC-013506 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListTRSRetailCustBlockForRelease_TM:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListTRSRetailCustBlockForRelease_TM.sql#L48) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `resultTable.SumNominal` : `MONEY (aggregate; facets unknown)` | `0` : `int` | `0` → `MONEY (aggregate; facets unknown)` | `isnull(resultTable.SumNominal,0)` |
| IC-013509 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListTRSRetailCustBlockForRelease_TM:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListTRSRetailCustBlockForRelease_TM.sql#L49) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `resultTable.SumNominal` : `MONEY (aggregate; facets unknown)` | `0` : `int` | `0` → `MONEY (aggregate; facets unknown)` | `isnull(resultTable.SumNominal,0)` |
| IC-013535 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListTRSRetailCustBlock_TM:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListTRSRetailCustBlock_TM.sql#L48) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `resultTable.SumNominal` : `MONEY (aggregate; facets unknown)` | `0` : `int` | `0` → `MONEY (aggregate; facets unknown)` | `isnull(resultTable.SumNominal,0)` |
| IC-013538 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListTRSRetailCustBlock_TM:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListTRSRetailCustBlock_TM.sql#L49) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `resultTable.SumNominal` : `MONEY (aggregate; facets unknown)` | `0` : `int` | `0` → `MONEY (aggregate; facets unknown)` | `isnull(resultTable.SumNominal,0)` |

<a id="kolom-safekeepingfeeaftertax"></a>
## SafeKeepingFeeAfterTax

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionLink_TR.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008394 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportSKFBankBeliBKON:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeliBKON.sql#L44) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tl.SafeKeepingFeeAfterTax` : `MONEY` | `#tempResult.NetAmount` : `DECIMAL(18,2)` | `tl.SafeKeepingFeeAfterTax` → `DECIMAL(18,2)` | `#tempResult.NetAmount <- tl.SafeKeepingFeeAfterTax` |
| IC-008399 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSKFBankBeliBKON:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeliBKON.sql#L107) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `tl.SafeKeepingFeeAfterTax` : `MONEY` | `(1 + sm.SafeKeepingFeeTaxTarif/100.0)` : `DECIMAL (combined result; facets unknown)` | `tl.SafeKeepingFeeAfterTax` → `DECIMAL` | `tl.SafeKeepingFeeAfterTax / (1 + sm.SafeKeepingFeeTaxTarif/100.0)` |
| IC-008401 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSKFBankBeliBKON:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeliBKON.sql#L108) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `tl.SafeKeepingFeeAfterTax` : `MONEY` | `(1 + sm.SafeKeepingFeeTaxTarif/100.0)` : `DECIMAL (combined result; facets unknown)` | `tl.SafeKeepingFeeAfterTax` → `DECIMAL` | `tl.SafeKeepingFeeAfterTax / (1 + sm.SafeKeepingFeeTaxTarif/100.0)` |
| IC-008402 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSKFBankBeliBKON:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeliBKON.sql#L108) | EXPRESSION_OR_CONTROL / P05 / BINARY_Subtract | `tl.SafeKeepingFeeAfterTax` : `MONEY` | `CAST((tl.SafeKeepingFeeAfterTax / (1 + sm.SafeKeepingFeeTaxTarif/100.0)) AS DECIMAL(25,2))` : `DECIMAL(25,2)` | `tl.SafeKeepingFeeAfterTax` → `DECIMAL` | `tl.SafeKeepingFeeAfterTax - CAST((tl.SafeKeepingFeeAfterTax / (1 + sm.SafeKeepingFeeTaxTarif/100.0)) AS DECIMAL(25,2))` |

<a id="kolom-taxonaccrued"></a>
## TaxOnAccrued

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionLink_TR.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008396 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportSKFBankBeliBKON:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeliBKON.sql#L44) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tl.TaxOnAccrued` : `MONEY` | `#tempResult.TaxOnInterest` : `DECIMAL(18,2)` | `tl.TaxOnAccrued` → `DECIMAL(18,2)` | `#tempResult.TaxOnInterest <- tl.TaxOnAccrued` |
| IC-000177 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.CustodyReportSettlementInstruction:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L22) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tr.TaxOnAccrued,0)` |
| IC-000179 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.CustodyReportSettlementInstruction:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L24) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tr.TaxOnAccrued,0)` |
| IC-000181 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.CustodyReportSettlementInstruction:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L29) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tr.TaxOnAccrued,0)` |
| IC-008252 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSReportReportSettlementInstruction:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L52) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tr.TaxOnAccrued,0)` |
| IC-008254 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSReportReportSettlementInstruction:54](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L54) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tr.TaxOnAccrued,0)` |
| IC-008256 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSReportReportSettlementInstruction:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L59) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tr.TaxOnAccrued,0)` |

<a id="kolom-taxoncapitalgainloss"></a>
## TaxOnCapitalGainLoss

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionLink_TR.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000178 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.CustodyReportSettlementInstruction:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L22) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.TaxOnCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tr.TaxOnCapitalGainLoss,0)` |
| IC-000182 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.CustodyReportSettlementInstruction:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L29) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.TaxOnCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tr.TaxOnCapitalGainLoss,0)` |
| IC-000183 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.CustodyReportSettlementInstruction:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L31) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.TaxOnCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tr.TaxOnCapitalGainLoss,0)` |
| IC-008253 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSReportReportSettlementInstruction:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L52) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.TaxOnCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tr.TaxOnCapitalGainLoss,0)` |
| IC-008257 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSReportReportSettlementInstruction:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L59) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.TaxOnCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tr.TaxOnCapitalGainLoss,0)` |
| IC-008258 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSReportReportSettlementInstruction:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L61) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.TaxOnCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tr.TaxOnCapitalGainLoss,0)` |
