# Retail UAT.dbo.ORI_Order_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 204 lokasi source, 51 procedure, 20 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 16 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cifid"></a>
## CIFId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_TT.sql#L3). **Existing leading index:** ORI_Order_TT_NN1.

**Frekuensi source:** 6 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004117 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGetCustOrderORIBuy:9](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetCustOrderORIBuy.sql#L9) | WHERE / P11 / COMPARISON | `CIFId` : `BIGINT` | `@pnCIFId` : `int` | `@pnCIFId` → `BIGINT` | `CIFId = @pnCIFId` |
| IC-012938 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_CancelOriOrder_TT:101](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_CancelOriOrder_TT.sql#L101) | JOIN_ON / P11 / COMPARISON | `o.CIFId` : `BIGINT` | `t.CIFId` : `int` | `t.CIFId` → `BIGINT` | `o.CIFId = t.CIFId` |
| IC-012941 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_CancelOriOrder_TT:161](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_CancelOriOrder_TT.sql#L161) | JOIN_ON / P11 / COMPARISON | `Trx.CIFId` : `BIGINT` | `t.CIFId` : `int` | `t.CIFId` → `BIGINT` | `Trx.CIFId = t.CIFId` |
| IC-012943 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_CancelOriOrder_TT:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_CancelOriOrder_TT.sql#L169) | JOIN_ON / P11 / COMPARISON | `Trx.CIFId` : `BIGINT` | `t.CIFId` : `int` | `t.CIFId` → `BIGINT` | `Trx.CIFId = t.CIFId` |
| IC-012945 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_CancelOriOrder_TT:229](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_CancelOriOrder_TT.sql#L229) | JOIN_ON / P11 / COMPARISON | `oott.CIFId` : `BIGINT` | `temp.CIFId` : `int` | `temp.CIFId` → `BIGINT` | `oott.CIFId = temp.CIFId` |
| IC-012948 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_CancelOriOrder_Temp:95](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_CancelOriOrder_Temp.sql#L95) | JOIN_ON / P11 / COMPARISON | `Trx.CIFId` : `BIGINT` | `t.CIFId` : `int` | `t.CIFId` → `BIGINT` | `Trx.CIFId = t.CIFId` |

<a id="kolom-custodianbit"></a>
## CustodianBit

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_TT.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009092 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailPopulateORIOrder:308](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulateORIOrder.sql#L308) | WHERE / P11 / COMPARISON | `ot.CustodianBit` : `BIT` | `0` : `int` | `ot.CustodianBit` → `INT` | `ot.CustodianBit = 0` |

<a id="kolom-insertedby"></a>
## InsertedBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_TT.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013427 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_ListStatusORI_Order_TT:120](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L120) | JOIN_ON / P01 / COMPARISON | `ot.InsertedBy` : `INT` | `inputer.EMPLID` : `VARCHAR(11)` | `inputer.EMPLID` → `INT` | `ot.InsertedBy = inputer.EMPLID` |
| IC-013434 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_ListStatusORI_Order_TT:197](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L197) | JOIN_ON / P01 / COMPARISON | `ot.InsertedBy` : `INT` | `inputer.EMPLID` : `VARCHAR(11)` | `inputer.EMPLID` → `INT` | `ot.InsertedBy = inputer.EMPLID` |

<a id="kolom-instruction"></a>
## Instruction

**Tipe dasar:** `CHAR(2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_TT.sql#L36). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 15 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004826 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSORIOnlinePopulateHist:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ori.[Instruction]` : `CHAR(2)` | `#tmp_Detail.Instruction` : `varchar(2)` | `ori.[Instruction]` → `varchar(2)` | `#tmp_Detail.Instruction <- ori.[Instruction]` |
| IC-004832 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSORIOnlinePopulateHist:134](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L134) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ori.[Instruction]` : `CHAR(2)` | `#tmp_Detail.Instruction` : `varchar(2)` | `ori.[Instruction]` → `varchar(2)` | `#tmp_Detail.Instruction <- ori.[Instruction]` |
| IC-004839 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSORIOnlinePopulateHist:163](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L163) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ori.[Instruction]` : `CHAR(2)` | `#tmp_Detail.Instruction` : `varchar(2)` | `ori.[Instruction]` → `varchar(2)` | `#tmp_Detail.Instruction <- ori.[Instruction]` |
| IC-004850 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSORIOnlinePopulateHist:220](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L220) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ori.[Instruction]` : `CHAR(2)` | `#tmp_Detail.Instruction` : `varchar(2)` | `ori.[Instruction]` → `varchar(2)` | `#tmp_Detail.Instruction <- ori.[Instruction]` |
| IC-003352 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSAutoSettlementOnlineORI:74](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAutoSettlementOnlineORI.sql#L74) | WHERE / P12 / COMPARISON | `ori.Instruction` : `CHAR(2)` | `'OM'` : `varchar` | `ori.Instruction` → `VARCHAR` | `ori.Instruction = 'OM'` |
| IC-004829 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSORIOnlinePopulateHist:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L130) | WHERE / P12 / COMPARISON | `ori.[Instruction]` : `CHAR(2)` | `'CB'` : `varchar` | `ori.[Instruction]` → `VARCHAR` | `ori.[Instruction] = 'CB'` |
| IC-004836 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSORIOnlinePopulateHist:159](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L159) | WHERE / P12 / COMPARISON | `ori.[Instruction]` : `CHAR(2)` | `'CB'` : `varchar` | `ori.[Instruction]` → `VARCHAR` | `ori.[Instruction] = 'CB'` |
| IC-004842 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSORIOnlinePopulateHist:188](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L188) | WHERE / P12 / COMPARISON | `ori.[Instruction]` : `CHAR(2)` | `'CB'` : `varchar` | `ori.[Instruction]` → `VARCHAR` | `ori.[Instruction] = 'CB'` |
| IC-004855 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSORIOnlinePopulateHist:246](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L246) | WHERE / P12 / IN_VALUE_TYPE | `ori.[Instruction]` : `CHAR(2)` | `'OM'` : `varchar` | `ori.[Instruction]` → `VARCHAR` | `ori.[Instruction] IN ('OM','MB')` |
| IC-005365 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDeleteORIOrder:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDeleteORIOrder.sql#L114) | WHERE / P12 / COMPARISON | `oott.Instruction` : `CHAR(2)` | `'OM'` : `varchar` | `oott.Instruction` → `VARCHAR` | `oott.Instruction <> 'OM'` |
| IC-006309 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementOrder:84](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementOrder.sql#L84) | WHERE / P12 / COMPARISON | `oott.Instruction` : `CHAR(2)` | `'OM'` : `varchar` | `oott.Instruction` → `VARCHAR` | `oott.Instruction <> 'OM'` |
| IC-006314 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementOrder:157](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementOrder.sql#L157) | WHERE / P12 / COMPARISON | `oott.Instruction` : `CHAR(2)` | `'OM'` : `varchar` | `oott.Instruction` → `VARCHAR` | `oott.Instruction <> 'OM'` |
| IC-012563 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdatePostingResult:448](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L448) | WHERE / P12 / COMPARISON | `oott.Instruction` : `CHAR(2)` | `'OM'` : `varchar` | `oott.Instruction` → `VARCHAR` | `oott.Instruction <> 'OM'` |
| IC-012565 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdatePostingResult:514](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L514) | WHERE / P12 / COMPARISON | `oott.Instruction` : `CHAR(2)` | `'OM'` : `varchar` | `oott.Instruction` → `VARCHAR` | `oott.Instruction = 'OM'` |
| IC-004818 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSORIOnlineNewTransaction:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlineNewTransaction.sql#L323) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_order_tt.Instruction` : `CHAR(2)` | `'OM'` : `varchar` | `'OM'` → `CHAR(2)` | `retail.dbo.ori_order_tt.Instruction <- 'OM'` |

<a id="kolom-isallotted"></a>
## IsAllotted

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_TT.sql#L19). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 16 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004830 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSORIOnlinePopulateHist:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L131) | WHERE / P11 / COMPARISON | `isnull(ori.[IsAllotted], 0)` : `BIT` | `0` : `int` | `isnull(ori.[IsAllotted], 0)` → `INT` | `isnull(ori.[IsAllotted], 0) = 0` |
| IC-004837 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSORIOnlinePopulateHist:160](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L160) | WHERE / P11 / COMPARISON | `isnull(ori.[IsAllotted], 0)` : `BIT` | `1` : `int` | `isnull(ori.[IsAllotted], 0)` → `INT` | `isnull(ori.[IsAllotted], 0) = 1` |
| IC-004843 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSORIOnlinePopulateHist:189](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L189) | WHERE / P11 / COMPARISON | `isnull(ori.[IsAllotted], 0)` : `BIT` | `1` : `int` | `isnull(ori.[IsAllotted], 0)` → `INT` | `isnull(ori.[IsAllotted], 0) = 1` |
| IC-005359 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateDeleteORIOrder:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDeleteORIOrder.sql#L81) | JOIN_ON / P11 / COMPARISON | `ISNULL(oott.IsAllotted, 0)` : `BIT` | `0` : `int` | `ISNULL(oott.IsAllotted, 0)` → `INT` | `ISNULL(oott.IsAllotted, 0) = 0` |
| IC-005362 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateDeleteORIOrder:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDeleteORIOrder.sql#L111) | WHERE / P11 / COMPARISON | `ISNULL(oott.IsAllotted, 0)` : `BIT` | `0` : `int` | `ISNULL(oott.IsAllotted, 0)` → `INT` | `ISNULL(oott.IsAllotted, 0) = 0` |
| IC-006308 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementOrder:83](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementOrder.sql#L83) | WHERE / P11 / COMPARISON | `oott.IsAllotted` : `BIT` | `1` : `int` | `oott.IsAllotted` → `INT` | `oott.IsAllotted = 1` |
| IC-006313 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementOrder:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementOrder.sql#L156) | WHERE / P11 / COMPARISON | `oott.IsAllotted` : `BIT` | `1` : `int` | `oott.IsAllotted` → `INT` | `oott.IsAllotted = 1` |
| IC-006539 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessAllotment:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessAllotment.sql#L65) | WHERE / P11 / COMPARISON | `oott.IsAllotted` : `BIT` | `1` : `int` | `oott.IsAllotted` → `INT` | `oott.IsAllotted = 1` |
| IC-008868 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailCalculateFee:1002](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L1002) | WHERE / P11 / COMPARISON | `oott.IsAllotted` : `BIT` | `1` : `int` | `oott.IsAllotted` → `INT` | `oott.IsAllotted = 1` |
| IC-011520 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSettledOrder:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettledOrder.sql#L71) | WHERE / P11 / COMPARISON | `oott.IsAllotted` : `BIT` | `1` : `int` | `oott.IsAllotted` → `INT` | `oott.IsAllotted = 1` |
| IC-004831 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSORIOnlinePopulateHist:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L131) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ori.[IsAllotted]` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(ori.[IsAllotted], 0)` |
| IC-004838 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSORIOnlinePopulateHist:160](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L160) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ori.[IsAllotted]` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(ori.[IsAllotted], 0)` |
| IC-004844 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSORIOnlinePopulateHist:189](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L189) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ori.[IsAllotted]` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(ori.[IsAllotted], 0)` |
| IC-005360 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateDeleteORIOrder:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDeleteORIOrder.sql#L81) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.IsAllotted` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(oott.IsAllotted, 0)` |
| IC-005363 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateDeleteORIOrder:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDeleteORIOrder.sql#L111) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.IsAllotted` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(oott.IsAllotted, 0)` |
| IC-006548 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessAllotment:136](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessAllotment.sql#L136) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `IsAllotted` : `BIT` | `1` : `int` | `1` → `BIT` | `IsAllotted = 1` |

<a id="kolom-isbuyback"></a>
## IsBuyBack

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_TT.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 14 lokasi / 13 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000974 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.ORISBNIPopulateEarlyRedemption:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateEarlyRedemption.sql#L147) | CASE / P11 / COMPARISON | `a.IsBuyBack` : `BIT` | `1` : `int` | `a.IsBuyBack` → `INT` | `a.IsBuyBack = 1` |
| IC-001008 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.ORISBNISubmitEarlyRedemption:276](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNISubmitEarlyRedemption.sql#L276) | CASE / P11 / COMPARISON | `tm.IsBuyBack` : `BIT` | `1` : `int` | `tm.IsBuyBack` → `INT` | `tm.IsBuyBack = 1` |
| IC-002758 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_trs_ListSecurityTransactionForInsertEditBlockTrans:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityTransactionForInsertEditBlockTrans.sql#L59) | JOIN_ON / P11 / COMPARISON | `ord.IsBuyBack` : `BIT` | `1` : `int` | `ord.IsBuyBack` → `INT` | `ord.IsBuyBack = 1` |
| IC-005335 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateBuyBack:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateBuyBack.sql#L70) | WHERE / P11 / COMPARISON | `oott.IsBuyBack` : `BIT` | `1` : `int` | `oott.IsBuyBack` → `INT` | `oott.IsBuyBack = 1` |
| IC-006214 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBuyBack:160](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L160) | WHERE / P11 / COMPARISON | `oott.IsBuyBack` : `BIT` | `1` : `int` | `oott.IsBuyBack` → `INT` | `oott.IsBuyBack = 1` |
| IC-006499 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateTransaksiBuyBack:203](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTransaksiBuyBack.sql#L203) | WHERE / P11 / COMPARISON | `oott.IsBuyBack` : `BIT` | `1` : `int` | `oott.IsBuyBack` → `INT` | `oott.IsBuyBack = 1` |
| IC-009572 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptCPT:87](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCPT.sql#L87) | WHERE / P11 / COMPARISON | `oott.IsBuyBack` : `BIT` | `1` : `int` | `oott.IsBuyBack` → `INT` | `oott.IsBuyBack = 1` |
| IC-013376 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListSecurityTransactionForInsertEditBlockTrans:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityTransactionForInsertEditBlockTrans.sql#L69) | JOIN_ON / P11 / COMPARISON | `ord.IsBuyBack` : `BIT` | `1` : `int` | `ord.IsBuyBack` → `INT` | `ord.IsBuyBack = 1` |
| IC-013399 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListSecurityTransactionForInsertEditBlockTrans_REPO:83](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityTransactionForInsertEditBlockTrans_REPO.sql#L83) | JOIN_ON / P11 / COMPARISON | `ord.IsBuyBack` : `BIT` | `1` : `int` | `ord.IsBuyBack` → `INT` | `ord.IsBuyBack = 1` |
| IC-013960 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ValidatePerubahanOrderSuratBerharga:126](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidatePerubahanOrderSuratBerharga.sql#L126) | WHERE / P11 / COMPARISON | `IsBuyBack` : `BIT` | `1` : `int` | `IsBuyBack` → `INT` | `IsBuyBack=1` |
| IC-006502 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateTrxBuyBackPerdana:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L96) | WHERE / P04 / COMPARISON | `oott.IsBuyBack` : `BIT` | `'1'` : `varchar` | `'1'` → `BIT` | `oott.IsBuyBack = '1'` |
| IC-006511 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateTrxBuyBackPerdana:157](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L157) | WHERE / P04 / COMPARISON | `oott.IsBuyBack` : `BIT` | `'1'` : `varchar` | `'1'` → `BIT` | `oott.IsBuyBack = '1'` |
| IC-006771 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessBuyBackOM:128](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBuyBackOM.sql#L128) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mst.IsBuyBack` : `BIT` | `1` : `int` | `1` → `BIT` | `mst.IsBuyBack = 1` |
| IC-012515 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUpdateOrderPriority:82](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateOrderPriority.sql#L82) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `IsBuyBack` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(IsBuyBack, 0)` |

<a id="kolom-iscashback"></a>
## IsCashBack

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_TT.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006245 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementCashBack:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementCashBack.sql#L71) | WHERE / P11 / COMPARISON | `oott.IsCashBack` : `BIT` | `1` : `int` | `oott.IsCashBack` → `INT` | `oott.IsCashBack = 1` |
| IC-011514 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSettledCashBack:58](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettledCashBack.sql#L58) | WHERE / P11 / COMPARISON | `oott.IsCashBack` : `BIT` | `1` : `int` | `oott.IsCashBack` → `INT` | `oott.IsCashBack = 1` |
| IC-013961 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ValidatePerubahanOrderSuratBerharga:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidatePerubahanOrderSuratBerharga.sql#L127) | WHERE / P11 / COMPARISON | `IsCashBack` : `BIT` | `1` : `int` | `IsCashBack` → `INT` | `IsCashBack=1` |
| IC-012516 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUpdateOrderPriority:83](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateOrderPriority.sql#L83) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `IsCashBack` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(IsCashBack, 0)` |

<a id="kolom-issplit"></a>
## IsSplit

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_TT.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006903 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessInsertORIOrderCommit:154](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertORIOrderCommit.sql#L154) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_order_tt.IsSplit` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.ori_order_tt.IsSplit <- 0` |
| IC-013105 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertORI_Order_TT_Original:563](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORI_Order_TT_Original.sql#L563) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_order_tt.IsSplit` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.ori_order_tt.IsSplit <- 0` |
| IC-013980 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_updateStatusORI_Order_Temp:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateStatusORI_Order_Temp.sql#L141) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_order_tt.IsSplit` : `BIT` | `1` : `int` | `1` → `BIT` | `retail.dbo.ori_order_tt.IsSplit <- 1` |
| IC-014004 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_updateTopUpORI_Order_Temp:91](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateTopUpORI_Order_Temp.sql#L91) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_order_tt.IsSplit` : `BIT` | `1` : `int` | `1` → `BIT` | `retail.dbo.ori_order_tt.IsSplit <- 1` |

<a id="kolom-nik-cs"></a>
## NIK_CS

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_TT.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013426 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_ListStatusORI_Order_TT:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L114) | JOIN_ON / P01 / COMPARISON | `ot.NIK_CS` : `INT` | `un.EMPLID` : `VARCHAR(11)` | `un.EMPLID` → `INT` | `ot.NIK_CS = un.EMPLID` |
| IC-013433 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_ListStatusORI_Order_TT:191](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L191) | JOIN_ON / P01 / COMPARISON | `ot.NIK_CS` : `INT` | `un.EMPLID` : `VARCHAR(11)` | `un.EMPLID` → `INT` | `ot.NIK_CS = un.EMPLID` |

<a id="kolom-nominalbuyback"></a>
## NominalBuyBack

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_TT.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 41 lokasi / 22 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002745 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.Rvmp_trs_ListSecurityMaster_TM:295](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L295) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sum(isnull(ori.NominalBuyBack, 0))` : `MONEY (aggregate; facets unknown)` | `#tmp_BuyBackCurrent.NominalBuyBack` : `decimal(24, 2)` | `sum(isnull(ori.NominalBuyBack, 0))` → `decimal(24, 2)` | `#tmp_BuyBackCurrent.NominalBuyBack <- sum(isnull(ori.NominalBuyBack, 0))` |
| IC-006191 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateSettlementBuyBack:50](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L50) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `((smtm.HargaBuyBack/100.00) * oott.NominalBuyBack) - oott.NominalBuyBack` : `DECIMAL (combined result; facets unknown)` | `#tmp_Result.CapitalGain` : `money` | `((smtm.HargaBuyBack/100.00) * oott.NominalBuyBack) - oott.NominalBuyBack` → `money` | `#tmp_Result.CapitalGain <- ((smtm.HargaBuyBack/100.00) * oott.NominalBuyBack) - oott.NominalBuyBack` |
| IC-013326 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_ListSecurityMasterPending_TM:371](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L371) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sum(isnull(ori.NominalBuyBack, 0))` : `MONEY (aggregate; facets unknown)` | `#tmp_BuyBackCurrent.NominalBuyBack` : `decimal(24,2)` | `sum(isnull(ori.NominalBuyBack, 0))` → `decimal(24,2)` | `#tmp_BuyBackCurrent.NominalBuyBack <- sum(isnull(ori.NominalBuyBack, 0))` |
| IC-006199 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBuyBack:99](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L99) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `oott.NominalBuyBack` : `MONEY` | `(smtm.HargaBuyBack/100.00)` : `DECIMAL (combined result; facets unknown)` | `oott.NominalBuyBack` → `DECIMAL` | `(smtm.HargaBuyBack/100.00) * oott.NominalBuyBack` |
| IC-006200 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBuyBack:99](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L99) | EXPRESSION_OR_CONTROL / P05 / BINARY_Subtract | `oott.NominalBuyBack` : `MONEY` | `((smtm.HargaBuyBack/100.00) * oott.NominalBuyBack)` : `DECIMAL (combined result; facets unknown)` | `oott.NominalBuyBack` → `DECIMAL` | `((smtm.HargaBuyBack/100.00) * oott.NominalBuyBack) - oott.NominalBuyBack` |
| IC-006202 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBuyBack:100](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L100) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `oott.NominalBuyBack` : `MONEY` | `(smtm.HargaBuyBack/100.00)` : `DECIMAL (combined result; facets unknown)` | `oott.NominalBuyBack` → `DECIMAL` | `(smtm.HargaBuyBack/100.00) * oott.NominalBuyBack` |
| IC-006204 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBuyBack:101](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L101) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `oott.NominalBuyBack` : `MONEY` | `(smtm.HargaBuyBack/100.00)` : `DECIMAL (combined result; facets unknown)` | `oott.NominalBuyBack` → `DECIMAL` | `(smtm.HargaBuyBack/100.00) * oott.NominalBuyBack` |
| IC-006205 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBuyBack:101](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L101) | EXPRESSION_OR_CONTROL / P05 / BINARY_Subtract | `oott.NominalBuyBack` : `MONEY` | `((smtm.HargaBuyBack/100.00) * oott.NominalBuyBack)` : `DECIMAL (combined result; facets unknown)` | `oott.NominalBuyBack` → `DECIMAL` | `((smtm.HargaBuyBack/100.00) * oott.NominalBuyBack) - oott.NominalBuyBack` |
| IC-006208 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateSettlementBuyBack:102](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L102) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `( ((smtm.HargaBuyBack/100) * oott.NominalBuyBack) - oott.NominalBuyBack )` : `MONEY (combined result; facets unknown)` | `(smtm.Tax / 100)` : `DECIMAL (combined result; facets unknown)` | `( ((smtm.HargaBuyBack/100) * oott.NominalBuyBack) - oott.NominalBuyBack )` → `DECIMAL` | `(smtm.Tax / 100) * ( ((smtm.HargaBuyBack/100) * oott.NominalBuyBack) - oott.NominalBuyBack )` |
| IC-006495 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateTransaksiBuyBack:148](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTransaksiBuyBack.sql#L148) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `(sum(NominalBuyBack) * (HargaBuyBack/100))` : `MONEY (combined result; facets unknown)` | `Accrued` : `DECIMAL(19,5)` | `(sum(NominalBuyBack) * (HargaBuyBack/100))` → `DECIMAL` | `(sum(NominalBuyBack) * (HargaBuyBack/100)) + Accrued` |
| IC-006497 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateTransaksiBuyBack:153](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTransaksiBuyBack.sql#L153) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `(((HargaBuyBack-100)/100) * sum(NominalBuyBack))` : `MONEY (combined result; facets unknown)` | `(Tax / 100)` : `DECIMAL (combined result; facets unknown)` | `(((HargaBuyBack-100)/100) * sum(NominalBuyBack))` → `DECIMAL` | `(((HargaBuyBack-100)/100) * sum(NominalBuyBack)) * (Tax / 100)` |
| IC-007429 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSProcessSettlementBuyBack:161](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBuyBack.sql#L161) | DML_CONDITION / P05 / BINARY_Multiply | `(((smtm.HargaBuyBack/100) * oott.NominalBuyBack) - oott.NominalBuyBack)` : `MONEY (combined result; facets unknown)` | `(smtm.Tax / 100)` : `DECIMAL (combined result; facets unknown)` | `(((smtm.HargaBuyBack/100) * oott.NominalBuyBack) - oott.NominalBuyBack)` → `DECIMAL` | `(smtm.Tax / 100) * (((smtm.HargaBuyBack/100) * oott.NominalBuyBack) - oott.NominalBuyBack)` |
| IC-007437 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSProcessSettlementBuyBack:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBuyBack.sql#L168) | CASE / P05 / BINARY_Multiply | `( ((smtm.HargaBuyBack/100) * oott.NominalBuyBack) - oott.NominalBuyBack )` : `MONEY (combined result; facets unknown)` | `(smtm.Tax / 100)` : `DECIMAL (combined result; facets unknown)` | `( ((smtm.HargaBuyBack/100) * oott.NominalBuyBack) - oott.NominalBuyBack )` → `DECIMAL` | `(smtm.Tax / 100) * ( ((smtm.HargaBuyBack/100) * oott.NominalBuyBack) - oott.NominalBuyBack )` |
| IC-002741 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:273](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L273) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(oott.NominalBuyBack, 0)` |
| IC-002746 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:301](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L301) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ori.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(ori.NominalBuyBack, 0)` |
| IC-002756 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityTransactionForInsertEditBlockTrans:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityTransactionForInsertEditBlockTrans.sql#L33) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ord.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ord.NominalBuyBack, 0)` |
| IC-004869 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSORIOrderSummary:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOrderSummary.sql#L77) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ot.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(ot.NominalBuyBack, 0)` |
| IC-004870 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSORIOrderSummary:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOrderSummary.sql#L80) | CASE / P11 / COMPARISON | `ot.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ot.NominalBuyBack > 0` |
| IC-005336 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateBuyBack:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateBuyBack.sql#L71) | WHERE / P11 / COMPARISON | `oott.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `oott.NominalBuyBack > 0` |
| IC-006500 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateTransaksiBuyBack:203](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTransaksiBuyBack.sql#L203) | WHERE / P11 / COMPARISON | `oott.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `oott.NominalBuyBack > 0` |
| IC-006503 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateTrxBuyBackPerdana:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L96) | WHERE / P11 / COMPARISON | `oott.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `oott.NominalBuyBack >0` |
| IC-006542 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessAllotment:91](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessAllotment.sql#L91) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(oott.NominalBuyBack, 0)` |
| IC-006545 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessAllotment:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessAllotment.sql#L98) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(oott.NominalBuyBack, 0)` |
| IC-006901 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessInsertORIOrderCommit:120](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertORIOrderCommit.sql#L120) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `SUM(NominalBuyBack)` : `MONEY (aggregate; facets unknown)` | `0` : `int` | `0` → `MONEY (aggregate; facets unknown)` | `ISNULL( SUM(NominalBuyBack), 0 )` |
| IC-009571 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptCPT:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCPT.sql#L86) | WHERE / P11 / COMPARISON | `oott.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `oott.NominalBuyBack > 0` |
| IC-012237 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSplitOrder:32](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSplitOrder.sql#L32) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(NominalBuyBack, 0)` |
| IC-012513 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUpdateOrderPriority:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateOrderPriority.sql#L80) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(NominalBuyBack, 0)` |
| IC-013065 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_InsertORIOrderTemp:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORIOrderTemp.sql#L156) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oot.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(oot.NominalBuyBack,0)` |
| IC-013074 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_InsertORIOrderTemp:189](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORIOrderTemp.sql#L189) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oot.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(oot.NominalBuyBack,0)` |
| IC-013098 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_InsertORI_Order_TT_Original:492](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORI_Order_TT_Original.sql#L492) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `SUM(NominalBuyBack)` : `MONEY (aggregate; facets unknown)` | `0` : `int` | `0` → `MONEY (aggregate; facets unknown)` | `ISNULL( SUM(NominalBuyBack), 0)` |
| IC-013322 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:354](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L354) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(oott.NominalBuyBack, 0)` |
| IC-013327 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:372](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L372) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ori.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(ori.NominalBuyBack, 0)` |
| IC-013363 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:379](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L379) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(oott.NominalBuyBack, 0)` |
| IC-013374 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityTransactionForInsertEditBlockTrans:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityTransactionForInsertEditBlockTrans.sql#L48) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ord.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ord.NominalBuyBack,0)` |
| IC-013397 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityTransactionForInsertEditBlockTrans_REPO:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityTransactionForInsertEditBlockTrans_REPO.sql#L61) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ord.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ord.NominalBuyBack,0)` |
| IC-013445 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusORI_Order_TT:250](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L250) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(oott.NominalBuyBack, 0)` |
| IC-013451 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusORI_Order_TT:289](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L289) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(oott.NominalBuyBack, 0)` |
| IC-013934 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidateInsertBuyBackOM:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateInsertBuyBackOM.sql#L51) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(NominalBuyBack,0)` |
| IC-013935 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidateInsertBuyBackOM:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateInsertBuyBackOM.sql#L61) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tableAggregate.SumNominalBuyBack` : `MONEY (aggregate; facets unknown)` | `0` : `int` | `0` → `MONEY (aggregate; facets unknown)` | `isnull(tableAggregate.SumNominalBuyBack,0)` |
| IC-013940 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidateInsertBuyBackOM:73](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateInsertBuyBackOM.sql#L73) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(oott.NominalBuyBack, 0)` |
| IC-013957 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidatePerubahanOrderSuratBerharga:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidatePerubahanOrderSuratBerharga.sql#L44) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(NominalBuyBack,0)` |

<a id="kolom-nominalcashback"></a>
## NominalCashBack

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_TT.sql#L24). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012238 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSplitOrder:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSplitOrder.sql#L33) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NominalCashBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(NominalCashBack, 0)` |
| IC-013998 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_updateTopUpORI_Order_Temp:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateTopUpORI_Order_Temp.sql#L80) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mst.NominalCashBack` : `MONEY` | `(mst.NominalCashBackValue + Trx.NominalCashBack) * sm.PersentaseCashBack / 100` : `DECIMAL (combined result; facets unknown)` | `(mst.NominalCashBackValue + Trx.NominalCashBack) * sm.PersentaseCashBack / 100` → `MONEY` | `mst.NominalCashBack = (mst.NominalCashBackValue + Trx.NominalCashBack) * sm.PersentaseCashBack / 100` |

<a id="kolom-nominalcashbackvalue"></a>
## NominalCashBackValue

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_TT.sql#L29). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 14 lokasi / 11 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002742 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:274](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L274) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.NominalCashBackValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(oott.NominalCashBackValue, 0)` |
| IC-006902 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessInsertORIOrderCommit:121](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertORIOrderCommit.sql#L121) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `SUM(NominalCashBackValue)` : `MONEY (aggregate; facets unknown)` | `0` : `int` | `0` → `MONEY (aggregate; facets unknown)` | `ISNULL( SUM(NominalCashBackValue), 0 )` |
| IC-012239 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSplitOrder:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSplitOrder.sql#L34) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NominalCashBackValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(NominalCashBackValue, 0)` |
| IC-012514 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUpdateOrderPriority:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateOrderPriority.sql#L81) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NominalCashBackValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(NominalCashBackValue, 0)` |
| IC-013066 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_InsertORIOrderTemp:157](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORIOrderTemp.sql#L157) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oot.NominalCashBackValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(oot.NominalCashBackValue,0)` |
| IC-013075 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_InsertORIOrderTemp:190](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORIOrderTemp.sql#L190) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oot.NominalCashBackValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(oot.NominalCashBackValue,0)` |
| IC-013099 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_InsertORI_Order_TT_Original:493](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORI_Order_TT_Original.sql#L493) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `SUM(NominalCashBackValue)` : `MONEY (aggregate; facets unknown)` | `0` : `int` | `0` → `MONEY (aggregate; facets unknown)` | `ISNULL( SUM(NominalCashBackValue), 0)` |
| IC-013323 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:355](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L355) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.NominalCashBackValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(oott.NominalCashBackValue, 0)` |
| IC-013364 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:381](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L381) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.NominalCashBackValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(oott.NominalCashBackValue, 0)` |
| IC-013446 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusORI_Order_TT:252](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L252) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.NominalCashBackValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(oott.NominalCashBackValue, 0)` |
| IC-013452 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusORI_Order_TT:291](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L291) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.NominalCashBackValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(oott.NominalCashBackValue, 0)` |
| IC-013933 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidateInsertBuyBackOM:50](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateInsertBuyBackOM.sql#L50) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NominalCashBackValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(NominalCashBackValue,0)` |
| IC-013941 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidateInsertBuyBackOM:74](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateInsertBuyBackOM.sql#L74) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.NominalCashBackValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(oott.NominalCashBackValue, 0)` |
| IC-013955 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidatePerubahanOrderSuratBerharga:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidatePerubahanOrderSuratBerharga.sql#L43) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NominalCashBackValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(NominalCashBackValue,0)` |

<a id="kolom-orderbranch"></a>
## OrderBranch

**Tipe dasar:** `VARCHAR(5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_TT.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT.

**Pola dan arah review:**

- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013423 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.trs_ListStatusORI_Order_TT:79](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L79) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `ot.OrderBranch` : `VARCHAR(5)` | `retail.dbo.fngetbranchcode.@cBranchCode` : `int` | `ot.OrderBranch` → `int` | `dbo.fnGetBranchCode(ot.OrderBranch)` |
| IC-013430 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.trs_ListStatusORI_Order_TT:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L156) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `ot.OrderBranch` : `VARCHAR(5)` | `retail.dbo.fngetbranchcode.@cBranchCode` : `int` | `ot.OrderBranch` → `int` | `dbo.fnGetBranchCode(ot.OrderBranch)` |

<a id="kolom-orderid"></a>
## OrderId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_TT.sql#L2). **Existing leading index:** IU_ORI_Order_TT&#124;PK__ORI_Order_TT__01142BA1.

**Frekuensi source:** 12 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000963 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.ORISBNIDeleteTransactionData:14](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIDeleteTransactionData.sql#L14) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `OrderId` : `BIGINT` | `@nOrderId` : `int` | `OrderId` → `int` | `@nOrderId = OrderId` |
| IC-006907 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessInsertORIOrderCommit:319](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertORIOrderCommit.sql#L319) | WHERE / P11 / COMPARISON | `OrderId` : `BIGINT` | `@nOrderId` : `DECIMAL` | `OrderId` → `DECIMAL` | `OrderId = @nOrderId` |
| IC-006911 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessInsertORIOrderCommit:335](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertORIOrderCommit.sql#L335) | WHERE / P11 / COMPARISON | `st.OrderId` : `BIGINT` | `@nOrderId` : `DECIMAL` | `st.OrderId` → `DECIMAL` | `st.OrderId = @nOrderId` |
| IC-000965 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.ORISBNIDeleteTransactionData:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIDeleteTransactionData.sql#L34) | WHERE / P11 / COMPARISON | `OrderId` : `BIGINT` | `@nOrderId` : `int` | `@nOrderId` → `BIGINT` | `OrderId = @nOrderId` |
| IC-012464 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSUpdateBlokirOrderORITransaction:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateBlokirOrderORITransaction.sql#L17) | WHERE / P11 / COMPARISON | `OrderId` : `BIGINT` | `@pnOrderId` : `int` | `@pnOrderId` → `BIGINT` | `OrderId = @pnOrderId` |
| IC-013063 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_InsertORIOrderTemp:154](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORIOrderTemp.sql#L154) | JOIN_ON / P01 / COMPARISON | `oot.OrderId` : `BIGINT` | `tmp.OrderIdReal` : `VARCHAR(25)` | `tmp.OrderIdReal` → `BIGINT` | `tmp.OrderIdReal = oot.OrderId` |
| IC-013067 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_InsertORIOrderTemp:163](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORIOrderTemp.sql#L163) | JOIN_ON / P01 / COMPARISON | `oot.OrderId` : `BIGINT` | `tmp.OrderIdReal` : `VARCHAR(25)` | `tmp.OrderIdReal` → `BIGINT` | `tmp.OrderIdReal = oot.OrderId` |
| IC-013070 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_InsertORIOrderTemp:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORIOrderTemp.sql#L180) | JOIN_ON / P01 / COMPARISON | `oot.OrderId` : `BIGINT` | `tmp.OrderIdReal` : `VARCHAR(25)` | `tmp.OrderIdReal` → `BIGINT` | `tmp.OrderIdReal = oot.OrderId` |
| IC-013072 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_InsertORIOrderTemp:187](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORIOrderTemp.sql#L187) | JOIN_ON / P01 / COMPARISON | `oot.OrderId` : `BIGINT` | `tmp.OrderIdReal` : `VARCHAR(25)` | `tmp.OrderIdReal` → `BIGINT` | `tmp.OrderIdReal = oot.OrderId` |
| IC-013085 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_InsertORIOrderTemp:305](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORIOrderTemp.sql#L305) | JOIN_ON / P01 / COMPARISON | `oot.OrderId` : `BIGINT` | `tmp.OrderIdReal` : `VARCHAR(25)` | `tmp.OrderIdReal` → `BIGINT` | `tmp.OrderIdReal = oot.OrderId` |
| IC-013107 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_InsertORI_Order_TT_Original:709](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORI_Order_TT_Original.sql#L709) | WHERE / P11 / COMPARISON | `OrderId` : `BIGINT` | `@OrderId` : `INT` | `@OrderId` → `BIGINT` | `OrderId = @OrderId` |
| IC-013111 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_InsertORI_Order_TT_Original:727](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORI_Order_TT_Original.sql#L727) | WHERE / P11 / COMPARISON | `st.OrderId` : `BIGINT` | `@OrderId` : `INT` | `@OrderId` → `BIGINT` | `st.OrderId = @OrderId` |

<a id="kolom-ordernominal"></a>
## OrderNominal

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_TT.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 22 lokasi / 16 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004827 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSORIOnlinePopulateHist:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ori.[OrderNominal]` : `MONEY` | `#tmp_Detail.Nominal` : `decimal(21,0)` | `ori.[OrderNominal]` → `decimal(21,0)` | `#tmp_Detail.Nominal <- ori.[OrderNominal]` |
| IC-004840 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSORIOnlinePopulateHist:163](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L163) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ori.[OrderNominal]` : `MONEY` | `#tmp_Detail.Nominal` : `decimal(21,0)` | `ori.[OrderNominal]` → `decimal(21,0)` | `#tmp_Detail.Nominal <- ori.[OrderNominal]` |
| IC-004851 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSORIOnlinePopulateHist:220](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L220) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ori.[OrderNominal]` : `MONEY` | `#tmp_Detail.Nominal` : `decimal(21,0)` | `ori.[OrderNominal]` → `decimal(21,0)` | `#tmp_Detail.Nominal <- ori.[OrderNominal]` |
| IC-002738 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L265) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.OrderNominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(oott.OrderNominal, 0)` |
| IC-005254 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateApprovalBuyBackOM:82](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalBuyBackOM.sql#L82) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tt.OrderNominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(tt.OrderNominal, 0)` |
| IC-006546 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessAllotment:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessAllotment.sql#L107) | WHERE / P11 / COMPARISON | `oott.OrderNominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `oott.OrderNominal > 0` |
| IC-006547 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessAllotment:123](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessAllotment.sql#L123) | WHERE / P11 / COMPARISON | `oott.OrderNominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `oott.OrderNominal > 0` |
| IC-006900 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessInsertORIOrderCommit:119](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertORIOrderCommit.sql#L119) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `SUM(OrderNominal)` : `MONEY (aggregate; facets unknown)` | `0` : `int` | `0` → `MONEY (aggregate; facets unknown)` | `ISNULL( SUM(OrderNominal), 0 )` |
| IC-012244 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSplitOrder:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSplitOrder.sql#L52) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `SUM(OrderNominal)` : `MONEY (aggregate; facets unknown)` | `0` : `int` | `0` → `MONEY (aggregate; facets unknown)` | `ISNULL(SUM(OrderNominal), 0)` |
| IC-012463 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateBlokirOrderORITransaction:14](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateBlokirOrderORITransaction.sql#L14) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `OrderNominal` : `MONEY` | `@pnOrderNominal` : `decimal` | `@pnOrderNominal` → `MONEY` | `OrderNominal = @pnOrderNominal` |
| IC-012511 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUpdateOrderPriority:32](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateOrderPriority.sql#L32) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `SUM(OrderNominal)` : `MONEY (aggregate; facets unknown)` | `0` : `int` | `0` → `MONEY (aggregate; facets unknown)` | `ISNULL( SUM(OrderNominal), 0 )` |
| IC-012958 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_CheckIsThereAnyOrderCorrection:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_CheckIsThereAnyOrderCorrection.sql#L13) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oot.OrderNominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(oot.OrderNominal,0)` |
| IC-013064 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_InsertORIOrderTemp:155](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORIOrderTemp.sql#L155) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oot.OrderNominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(oot.OrderNominal,0)` |
| IC-013073 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_InsertORIOrderTemp:188](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORIOrderTemp.sql#L188) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oot.OrderNominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(oot.OrderNominal,0)` |
| IC-013097 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_InsertORI_Order_TT_Original:491](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORI_Order_TT_Original.sql#L491) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `SUM(OrderNominal)` : `MONEY (aggregate; facets unknown)` | `0` : `int` | `0` → `MONEY (aggregate; facets unknown)` | `ISNULL( SUM(OrderNominal), 0)` |
| IC-013319 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:347](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L347) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.OrderNominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(oott.OrderNominal, 0)` |
| IC-013360 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:372](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L372) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.OrderNominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(oott.OrderNominal, 0)` |
| IC-013442 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusORI_Order_TT:243](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L243) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.OrderNominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(oott.OrderNominal, 0)` |
| IC-013448 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusORI_Order_TT:282](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L282) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.OrderNominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(oott.OrderNominal, 0)` |
| IC-013932 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidateInsertBuyBackOM:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateInsertBuyBackOM.sql#L49) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `OrderNominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(OrderNominal,0)` |
| IC-013937 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidateInsertBuyBackOM:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateInsertBuyBackOM.sql#L66) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.OrderNominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(oott.OrderNominal, 0)` |
| IC-013956 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidatePerubahanOrderSuratBerharga:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidatePerubahanOrderSuratBerharga.sql#L43) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `OrderNominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(OrderNominal,0)` |

<a id="kolom-originalordernominal"></a>
## OriginalOrderNominal

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_TT.sql#L28). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004833 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSORIOnlinePopulateHist:134](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L134) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `isnull(ori.[OriginalOrderNominal],0)` : `MONEY` | `#tmp_Detail.Nominal` : `decimal(21,0)` | `isnull(ori.[OriginalOrderNominal],0)` → `decimal(21,0)` | `#tmp_Detail.Nominal <- isnull(ori.[OriginalOrderNominal],0)` |
| IC-004835 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSORIOnlinePopulateHist:140](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L140) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ori.[OriginalOrderNominal]` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(ori.[OriginalOrderNominal],0)` |

<a id="kolom-phoneorderbit"></a>
## PhoneOrderBit

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_TT.sql#L30). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013424 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListStatusORI_Order_TT:91](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L91) | CASE / P11 / COMPARISON | `ot.PhoneOrderBit` : `BIT` | `1` : `int` | `ot.PhoneOrderBit` → `INT` | `ot.PhoneOrderBit = 1` |
| IC-013431 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListStatusORI_Order_TT:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L168) | CASE / P11 / COMPARISON | `ot.PhoneOrderBit` : `BIT` | `1` : `int` | `ot.PhoneOrderBit` → `INT` | `ot.PhoneOrderBit = 1` |
| IC-006772 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessBuyBackOM:135](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBuyBackOM.sql#L135) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mst.PhoneOrderBit` : `BIT` | `CASE WHEN Trx.PhoneOrderBit = 'Y' THEN 1 WHEN Trx.PhoneOrderBit = 'N' THEN 0 END` : `INT (combined result; facets unknown)` | `CASE WHEN Trx.PhoneOrderBit = 'Y' THEN 1 WHEN Trx.PhoneOrderBit = 'N' THEN 0 END` → `BIT` | `mst.PhoneOrderBit = CASE WHEN Trx.PhoneOrderBit = 'Y' THEN 1 WHEN Trx.PhoneOrderBit = 'N' THEN 0 END` |

<a id="kolom-priority"></a>
## Priority

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_TT.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 24 lokasi / 13 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, DESTINATION_COLUMN, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005218 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateAllotment:103](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateAllotment.sql#L103) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `oott.Priority` : `BIT` | `@tmp_AllocatedAlloted.Priority` : `int` | `oott.Priority` → `int` | `@tmp_AllocatedAlloted.Priority <- oott.Priority` |
| IC-006505 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateTrxBuyBackPerdana:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `oott.Priority` : `BIT` | `retail.dbo.bondsdetailtransaksibuybackperdana_tr.Priority` : `INT` | `oott.Priority` → `INT` | `retail.dbo.bondsdetailtransaksibuybackperdana_tr.Priority <- oott.Priority` |
| IC-002740 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.Rvmp_trs_ListSecurityMaster_TM:267](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L267) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ISNULL(oott.Priority, 0)` : `BIT` | `0` : `int` | `ISNULL(oott.Priority, 0)` → `INT` | `WHEN 0 THEN 0` |
| IC-006799 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessCancelORIOrderCommit:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessCancelORIOrderCommit.sql#L77) | WHERE / P11 / COMPARISON | `Priority` : `BIT` | `1` : `int` | `Priority` → `INT` | `Priority = 1` |
| IC-009080 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailPopulateORIOrder:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulateORIOrder.sql#L86) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ot.Priority` : `BIT` | `0` : `int` | `ot.Priority` → `INT` | `WHEN 0 THEN '0'` |
| IC-012245 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSplitOrder:55](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSplitOrder.sql#L55) | WHERE / P11 / COMPARISON | `Priority` : `BIT` | `0` : `int` | `Priority` → `INT` | `Priority = 0` |
| IC-012246 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSplitOrder:155](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSplitOrder.sql#L155) | SET_OUTPUT / P07 / UNION_OUTPUT | `right output #8` : `BIT` | `left output #8` : `INT (combined result; facets unknown)` | `right output #8` → `INT` | `SELECT CIFId ,SecId ,ExistingCustomer ,OrderDate ,OrderBranch ,@pnOrderNominal AS OrderNominal ,CustodianBit ,CASE @pbUpdateToPriority0 WHEN 1 THEN 0 ELSE Priority END AS Priority ,OrderStatus ,NIK_CS ,InsertedBy ,InsertedDate ,LastUpdateBy ,LastUpdateDate ,AccountBlockACTYPE ,AccountBlockSequence ,IsAllotted ,IsBuyBack ,IsCashBack ,1 AS IsSplit ,@pnNominalBuyBack AS NominalBuyBack ,@nNominalCashBack AS NominalCashBack ,NominalBuyBack AS OriginalNominalBuyBack ,NominalCashBack AS OriginalNominalCashBack ,@pnOrderId AS OriginalOrderId ,OrderNominal AS OriginalOrderNominal ,@pnNominalCashBack AS NominalCashBackValue ,PhoneOrderBit ,NIK_Ref ,NoRekInvestor ,isTaxAmnesty ,NominalCashBackValue AS OriginalCashBackValue FROM ORI_Order_TT WHERE OrderId = @pnOrderId UNION ALL SELECT CIFId ,SecId ,ExistingCustomer ,OrderDate ,OrderBranch ,OrderNominal - @pnOrderNominal AS OrderNominal ,CustodianBit ,Priority ,OrderStatus ,NIK_CS ,InsertedBy ,InsertedDate ,LastUpdateBy ,LastUpdateDate ,AccountBlockACTYPE ,AccountBlockSequence ,IsAllotted ,IsBuyBack ,IsCashBack ,1 AS IsSplit ,NominalBuyBack - @pnNominalBuyBack AS NominalBuyBack ,NominalCashBack - @nNominalCashBack AS NominalCashBack ,NominalBuyBack AS OriginalNominalBuyBack ,NominalCashBack AS OriginalNominalCashBack ,@pnOrderId AS OriginalOrderId ,OrderNominal AS OriginalOrderNominal ,NominalCashBackValue - @pnNominalCashBack AS NominalCashBackValue ,PhoneOrderBit ,NIK_Ref ,NoRekInvestor ,isTaxAmnesty ,NominalCashBackValue AS OriginalCashBackValue FROM ORI_Order_TT WHERE OrderId = @pnOrderId` |
| IC-012247 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSplitOrder:163](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSplitOrder.sql#L163) | EXPRESSION / P07 / CASE_BRANCH | `Priority` : `BIT` | `result type` : `INT (combined result; facets unknown)` | `Priority` → `INT` | `CASE @pbUpdateToPriority0 WHEN 1 THEN 0 ELSE Priority END` |
| IC-012512 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdateOrderPriority:35](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateOrderPriority.sql#L35) | WHERE / P11 / COMPARISON | `Priority` : `BIT` | `0` : `int` | `Priority` → `INT` | `Priority = 0` |
| IC-012517 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdateOrderPriority:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateOrderPriority.sql#L86) | WHERE / P11 / COMPARISON | `Priority` : `BIT` | `1` : `int` | `Priority` → `INT` | `Priority = 1` |
| IC-012940 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_CancelOriOrder_TT:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_CancelOriOrder_TT.sql#L108) | WHERE / P11 / COMPARISON | `Priority` : `BIT` | `1` : `int` | `Priority` → `INT` | `Priority = 1` |
| IC-013321 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_ListSecurityMasterPending_TM:350](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L350) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ISNULL(oott.Priority, 0)` : `BIT` | `0` : `int` | `ISNULL(oott.Priority, 0)` → `INT` | `WHEN 0 THEN 0` |
| IC-013362 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_ListSecurityMaster_TM:375](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L375) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ISNULL(oott.Priority, 0)` : `BIT` | `0` : `int` | `ISNULL(oott.Priority, 0)` → `INT` | `WHEN 0 THEN 0` |
| IC-013444 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_ListStatusORI_Order_TT:246](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L246) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ISNULL(oott.Priority, 0)` : `BIT` | `0` : `int` | `ISNULL(oott.Priority, 0)` → `INT` | `WHEN 0 THEN 0` |
| IC-013450 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_ListStatusORI_Order_TT:285](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L285) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ISNULL(oott.Priority, 0)` : `BIT` | `0` : `int` | `ISNULL(oott.Priority, 0)` → `INT` | `WHEN 0 THEN 0` |
| IC-013939 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_ValidateInsertBuyBackOM:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateInsertBuyBackOM.sql#L69) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ISNULL(oott.Priority, 0)` : `BIT` | `0` : `int` | `ISNULL(oott.Priority, 0)` → `INT` | `WHEN 0 THEN 0` |
| IC-002739 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:266](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L266) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.Priority` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(oott.Priority, 0)` |
| IC-006904 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessInsertORIOrderCommit:154](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertORIOrderCommit.sql#L154) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_order_tt.Priority` : `BIT` | `CASE WHEN @nSisaKuota > 0 THEN Priority ELSE 1 END` : `INT (combined result; facets unknown)` | `CASE WHEN @nSisaKuota > 0 THEN Priority ELSE 1 END` → `BIT` | `retail.dbo.ori_order_tt.Priority <- CASE WHEN @nSisaKuota > 0 THEN Priority ELSE 1 END` |
| IC-012521 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateOrderPriority:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateOrderPriority.sql#L156) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Priority` : `BIT` | `0` : `int` | `0` → `BIT` | `Priority = 0` |
| IC-013320 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:349](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L349) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.Priority` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(oott.Priority, 0)` |
| IC-013361 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:374](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L374) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.Priority` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(oott.Priority, 0)` |
| IC-013443 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusORI_Order_TT:245](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L245) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.Priority` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(oott.Priority, 0)` |
| IC-013449 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusORI_Order_TT:284](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L284) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.Priority` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(oott.Priority, 0)` |
| IC-013938 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidateInsertBuyBackOM:68](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateInsertBuyBackOM.sql#L68) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.Priority` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(oott.Priority, 0)` |

<a id="kolom-secid"></a>
## SecId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_TT.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013970 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_updateStatusAccORI_Order_TT:124](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateStatusAccORI_Order_TT.sql#L124) | JOIN_ON / P01 / COMPARISON | `Trx.SecId` : `INT` | `t.SecId` : `varchar(20)` | `t.SecId` → `INT` | `Trx.SecId=t.SecId` |

<a id="kolom-istaxamnesty"></a>
## isTaxAmnesty

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_TT.sql#L33). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 17 lokasi / 7 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006209 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateSettlementBuyBack:121](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L121) | CASE / P11 / COMPARISON | `isnull(oott.isTaxAmnesty,0)` : `BIT` | `1` : `int` | `isnull(oott.isTaxAmnesty,0)` → `INT` | `isnull(oott.isTaxAmnesty,0) = 1` |
| IC-006306 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateSettlementOrder:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementOrder.sql#L49) | CASE / P11 / COMPARISON | `isnull(oott.isTaxAmnesty,0)` : `BIT` | `1` : `int` | `isnull(oott.isTaxAmnesty,0)` → `INT` | `isnull(oott.isTaxAmnesty,0) = 1` |
| IC-007462 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSProcessSettlementCashBack:102](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementCashBack.sql#L102) | CASE / P11 / COMPARISON | `isnull(oott.isTaxAmnesty,0)` : `BIT` | `1` : `int` | `isnull(oott.isTaxAmnesty,0)` → `INT` | `isnull(oott.isTaxAmnesty,0) = 1` |
| IC-007524 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSProcessSettlementOrder:134](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementOrder.sql#L134) | CASE / P11 / COMPARISON | `isnull(oott.isTaxAmnesty,0)` : `BIT` | `1` : `int` | `isnull(oott.isTaxAmnesty,0)` → `INT` | `isnull(oott.isTaxAmnesty,0) = 1` |
| IC-009079 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailPopulateORIOrder:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulateORIOrder.sql#L81) | CASE / P11 / COMPARISON | `ot.isTaxAmnesty` : `BIT` | `0` : `int` | `ot.isTaxAmnesty` → `INT` | `ot.isTaxAmnesty = 0` |
| IC-009081 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailPopulateORIOrder:101](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulateORIOrder.sql#L101) | JOIN_ON / P11 / COMPARISON | `ot.isTaxAmnesty` : `BIT` | `1` : `int` | `ot.isTaxAmnesty` → `INT` | `ot.isTaxAmnesty = 1` |
| IC-013425 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListStatusORI_Order_TT:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L106) | CASE / P11 / COMPARISON | `ot.isTaxAmnesty` : `BIT` | `1` : `int` | `ot.isTaxAmnesty` → `INT` | `ot.isTaxAmnesty = 1` |
| IC-013432 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListStatusORI_Order_TT:183](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L183) | CASE / P11 / COMPARISON | `ot.isTaxAmnesty` : `BIT` | `1` : `int` | `ot.isTaxAmnesty` → `INT` | `ot.isTaxAmnesty = 1` |
| IC-013860 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_UpdateTreasuryCustomer_TM:586](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L586) | WHERE / P11 / COMPARISON | `oott.isTaxAmnesty` : `BIT` | `1` : `int` | `oott.isTaxAmnesty` → `INT` | `oott.isTaxAmnesty = 1` |
| IC-013865 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_UpdateTreasuryCustomer_TM:632](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L632) | WHERE / P11 / COMPARISON | `oott.isTaxAmnesty` : `BIT` | `0` : `int` | `oott.isTaxAmnesty` → `INT` | `oott.isTaxAmnesty = 0` |
| IC-013877 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_UpdateTreasuryCustomer_TM:689](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L689) | WHERE / P11 / COMPARISON | `oott.isTaxAmnesty` : `BIT` | `0` : `int` | `oott.isTaxAmnesty` → `INT` | `oott.isTaxAmnesty = 0` |
| IC-013888 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_UpdateTreasuryCustomer_TM:750](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L750) | WHERE / P11 / COMPARISON | `oott.isTaxAmnesty` : `BIT` | `0` : `int` | `oott.isTaxAmnesty` → `INT` | `oott.isTaxAmnesty = 0` |
| IC-013900 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_UpdateTreasuryCustomer_TM:803](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L803) | WHERE / P11 / COMPARISON | `oott.isTaxAmnesty` : `BIT` | `0` : `int` | `oott.isTaxAmnesty` → `INT` | `oott.isTaxAmnesty = 0` |
| IC-006210 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateSettlementBuyBack:121](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L121) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(oott.isTaxAmnesty,0)` |
| IC-006307 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateSettlementOrder:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementOrder.sql#L49) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(oott.isTaxAmnesty,0)` |
| IC-007463 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSettlementCashBack:102](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementCashBack.sql#L102) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(oott.isTaxAmnesty,0)` |
| IC-007525 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSettlementOrder:134](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementOrder.sql#L134) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `oott.isTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(oott.isTaxAmnesty,0)` |
