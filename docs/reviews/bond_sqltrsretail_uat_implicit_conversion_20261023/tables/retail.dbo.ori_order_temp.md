# Retail UAT.dbo.ORI_Order_Temp — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 54 lokasi source, 16 procedure, 11 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 7 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cifid"></a>
## CIFId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_Temp.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012950 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_CancelOriOrder_Temp:99](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_CancelOriOrder_Temp.sql#L99) | JOIN_ON / P11 / COMPARISON | `Trx.CIFId` : `BIGINT` | `t.CIFId` : `int` | `t.CIFId` → `BIGINT` | `Trx.CIFId=t.CIFId` |

<a id="kolom-materai"></a>
## Materai

**Tipe dasar:** `DECIMAL(19,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_Temp.sql#L26). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012566 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSUpdatePostingResult:520](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L520) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tmp.Materai` : `DECIMAL(19,2)` | `Materai` : `MONEY` | `tmp.Materai` → `MONEY` | `Materai = tmp.Materai` |
| IC-004871 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSORIOrderSummary:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOrderSummary.sql#L85) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tp.Materai` : `DECIMAL(19,2)` | `0` : `int` | `0` → `DECIMAL(19,2)` | `ISNULL(tp.Materai,0)` |
| IC-012567 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSUpdatePostingResult:525](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L525) | WHERE / P11 / COMPARISON | `tmp.Materai` : `DECIMAL(19,2)` | `0` : `int` | `0` → `DECIMAL` | `tmp.Materai > 0` |

<a id="kolom-nik-cs"></a>
## NIK_CS

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_Temp.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000913 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.ONEApprovalPrepareData:167](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareData.sql#L167) | JOIN_ON / P01 / COMPARISON | `ot.NIK_CS` : `INT` | `un.EMPLID` : `VARCHAR(11)` | `un.EMPLID` → `INT` | `ot.NIK_CS = un.EMPLID` |
| IC-000941 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.ONEApprovalPrepareDataBond:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareDataBond.sql#L165) | JOIN_ON / P01 / COMPARISON | `ot.NIK_CS` : `INT` | `un.EMPLID` : `VARCHAR(11)` | `un.EMPLID` → `INT` | `ot.NIK_CS = un.EMPLID` |
| IC-005256 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateApprovalBuyBackOM:90](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalBuyBackOM.sql#L90) | JOIN_ON / P01 / COMPARISON | `ot.NIK_CS` : `INT` | `un.EMPLID` : `VARCHAR(11)` | `un.EMPLID` → `INT` | `ot.NIK_CS = un.EMPLID` |
| IC-013458 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_ListStatusORI_Order_Temp:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_Temp.sql#L64) | JOIN_ON / P01 / COMPARISON | `ot.NIK_CS` : `INT` | `un.EMPLID` : `VARCHAR(11)` | `un.EMPLID` → `INT` | `ot.NIK_CS = un.EMPLID` |
| IC-013462 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_ListStatusORI_Order_Temp:109](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_Temp.sql#L109) | JOIN_ON / P01 / COMPARISON | `ot.NIK_CS` : `INT` | `un.EMPLID` : `VARCHAR(11)` | `un.EMPLID` → `INT` | `ot.NIK_CS = un.EMPLID` |

<a id="kolom-nominalbuyback"></a>
## NominalBuyBack

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_Temp.sql#L19). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002743 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.Rvmp_trs_ListSecurityMaster_TM:282](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L282) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sum(isnull(ori.NominalBuyBack, 0))` : `MONEY (aggregate; facets unknown)` | `#tmp_BuyBackPending.NominalBuyBack` : `decimal(24, 2)` | `sum(isnull(ori.NominalBuyBack, 0))` → `decimal(24, 2)` | `#tmp_BuyBackPending.NominalBuyBack <- sum(isnull(ori.NominalBuyBack, 0))` |
| IC-013324 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_ListSecurityMasterPending_TM:363](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L363) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sum(isnull(ori.NominalBuyBack, 0))` : `MONEY (aggregate; facets unknown)` | `#tmp_BuyBackPending.NominalBuyBack` : `decimal(24,2)` | `sum(isnull(ori.NominalBuyBack, 0))` → `decimal(24,2)` | `#tmp_BuyBackPending.NominalBuyBack <- sum(isnull(ori.NominalBuyBack, 0))` |
| IC-013365 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_ListSecurityMaster_TM:389](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L389) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sum(isnull(ori.NominalBuyBack, 0))` : `MONEY (aggregate; facets unknown)` | `#tmp_BuyBackPending.NominalBuyBack` : `decimal(24,2)` | `sum(isnull(ori.NominalBuyBack, 0))` → `decimal(24,2)` | `#tmp_BuyBackPending.NominalBuyBack <- sum(isnull(ori.NominalBuyBack, 0))` |
| IC-002744 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:288](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L288) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ori.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(ori.NominalBuyBack, 0)` |
| IC-005253 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateApprovalBuyBackOM:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalBuyBackOM.sql#L65) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ot.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(ot.NominalBuyBack,0)` |
| IC-009573 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptChangeORIOrder:42](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptChangeORIOrder.sql#L42) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ott.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(ott.NominalBuyBack,0)` |
| IC-013076 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_InsertORIOrderTemp:213](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORIOrderTemp.sql#L213) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sum(NominalBuyBack)` : `MONEY (aggregate; facets unknown)` | `0` : `int` | `0` → `MONEY (aggregate; facets unknown)` | `isnull(sum(NominalBuyBack), 0)` |
| IC-013325 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:364](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L364) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ori.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(ori.NominalBuyBack, 0)` |
| IC-013366 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:390](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L390) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ori.NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(ori.NominalBuyBack, 0)` |

<a id="kolom-nominalcashback"></a>
## NominalCashBack

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_Temp.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009574 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptChangeORIOrder:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptChangeORIOrder.sql#L44) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ott.NominalCashBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(ott.NominalCashBack,0)` |

<a id="kolom-orderbranch"></a>
## OrderBranch

**Tipe dasar:** `VARCHAR(5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_Temp.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT.

**Pola dan arah review:**

- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005251 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSPopulateApprovalBuyBackOM:58](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalBuyBackOM.sql#L58) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `ot.OrderBranch` : `VARCHAR(5)` | `retail.dbo.fngetbranchcode.@cBranchCode` : `int` | `ot.OrderBranch` → `int` | `dbo.fnGetBranchCode(ot.OrderBranch)` |
| IC-013456 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.trs_ListStatusORI_Order_Temp:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_Temp.sql#L43) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `ot.OrderBranch` : `VARCHAR(5)` | `retail.dbo.fngetbranchcode.@cBranchCode` : `int` | `ot.OrderBranch` → `int` | `dbo.fnGetBranchCode(ot.OrderBranch)` |
| IC-013460 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.trs_ListStatusORI_Order_Temp:88](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_Temp.sql#L88) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `ot.OrderBranch` : `VARCHAR(5)` | `retail.dbo.fngetbranchcode.@cBranchCode` : `int` | `ot.OrderBranch` → `int` | `dbo.fnGetBranchCode(ot.OrderBranch)` |

<a id="kolom-orderid"></a>
## OrderId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_Temp.sql#L2). **Existing leading index:** PK__ORI_Order_Temp__74A4FE72.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000914 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.ONEApprovalPrepareData:174](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareData.sql#L174) | JOIN_ON / P01 / COMPARISON | `ot.OrderId` : `BIGINT` | `oap.SearchTerm` : `VARCHAR(100)` | `oap.SearchTerm` → `BIGINT` | `ot.OrderId = oap.SearchTerm` |
| IC-000942 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.ONEApprovalPrepareDataBond:172](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareDataBond.sql#L172) | JOIN_ON / P01 / COMPARISON | `ot.OrderId` : `BIGINT` | `oap.SearchTerm` : `VARCHAR(100)` | `oap.SearchTerm` → `BIGINT` | `ot.OrderId = oap.SearchTerm` |

<a id="kolom-orderidreal"></a>
## OrderIdReal

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_Temp.sql#L21). **Existing leading index:** ORI_Order_Temp_NN2.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013068 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_InsertORIOrderTemp:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORIOrderTemp.sql#L165) | JOIN_ON / P01 / COMPARISON | `ori.OrderIdReal` : `BIGINT` | `tmp.OrderIdReal` : `VARCHAR(25)` | `tmp.OrderIdReal` → `BIGINT` | `tmp.OrderIdReal = ori.OrderIdReal` |
| IC-013077 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_InsertORIOrderTemp:217](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORIOrderTemp.sql#L217) | WHERE / P04 / IN_SUBQUERY | `OrderIdReal` : `BIGINT` | `(select OrderIdReal from #tempORI_Order_Temp)` : `VARCHAR(25)` | `(select OrderIdReal from #tempORI_Order_Temp)` → `BIGINT` | `OrderIdReal not in (select OrderIdReal from #tempORI_Order_Temp)` |
| IC-013079 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertORIOrderTemp:258](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORIOrderTemp.sql#L258) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_order_temp.OrderIdReal` : `BIGINT` | `tmp.OrderIdReal` : `VARCHAR(25)` | `tmp.OrderIdReal` → `BIGINT` | `retail.dbo.ori_order_temp.OrderIdReal <- tmp.OrderIdReal` |

<a id="kolom-ordernominal"></a>
## OrderNominal

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_Temp.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005252 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateApprovalBuyBackOM:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalBuyBackOM.sql#L64) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ot.OrderNominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(ot.OrderNominal,0)` |
| IC-013457 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusORI_Order_Temp:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_Temp.sql#L45) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ot.OrderNominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(ot.OrderNominal,0)` |
| IC-013461 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusORI_Order_Temp:90](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_Temp.sql#L90) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ot.OrderNominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(ot.OrderNominal,0)` |

<a id="kolom-priority"></a>
## Priority

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_Temp.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006774 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBuyBackOM:148](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBuyBackOM.sql#L148) | WHERE / P11 / COMPARISON | `Trx.Priority` : `BIT` | `0` : `int` | `Trx.Priority` → `INT` | `Trx.Priority = 0` |
| IC-006778 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBuyBackOM:157](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBuyBackOM.sql#L157) | WHERE / P11 / COMPARISON | `Trx.Priority` : `BIT` | `0` : `int` | `Trx.Priority` → `INT` | `Trx.Priority = 0` |
| IC-013975 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_updateStatusORI_Order_Temp:122](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateStatusORI_Order_Temp.sql#L122) | WHERE / P11 / COMPARISON | `Trx.Priority` : `BIT` | `0` : `int` | `Trx.Priority` → `INT` | `Trx.Priority = 0` |
| IC-013978 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_updateStatusORI_Order_Temp:139](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateStatusORI_Order_Temp.sql#L139) | WHERE / P11 / COMPARISON | `Trx.Priority` : `BIT` | `0` : `int` | `Trx.Priority` → `INT` | `Trx.Priority = 0` |
| IC-013982 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_updateStatusORI_Order_Temp:179](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateStatusORI_Order_Temp.sql#L179) | WHERE / P11 / COMPARISON | `Trx.Priority` : `BIT` | `1` : `int` | `Trx.Priority` → `INT` | `Trx.Priority = 1` |
| IC-013990 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_updateStatusORI_Order_Temp:207](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateStatusORI_Order_Temp.sql#L207) | WHERE / P11 / COMPARISON | `st.Priority` : `BIT` | `1` : `int` | `st.Priority` → `INT` | `st.Priority = 1` |
| IC-014002 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_updateTopUpORI_Order_Temp:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateTopUpORI_Order_Temp.sql#L89) | WHERE / P11 / COMPARISON | `Trx.Priority` : `BIT` | `0` : `int` | `Trx.Priority` → `INT` | `Trx.Priority = 0` |
| IC-014006 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_updateTopUpORI_Order_Temp:135](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateTopUpORI_Order_Temp.sql#L135) | WHERE / P11 / COMPARISON | `Trx.Priority` : `BIT` | `1` : `int` | `Trx.Priority` → `INT` | `Trx.Priority = 1` |
| IC-014012 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_updateTopUpORI_Order_Temp:153](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateTopUpORI_Order_Temp.sql#L153) | WHERE / P11 / COMPARISON | `st.Priority` : `BIT` | `1` : `int` | `st.Priority` → `INT` | `st.Priority = 1` |

<a id="kolom-secid"></a>
## SecId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_Temp.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 15 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006769 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessBuyBackOM:121](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBuyBackOM.sql#L121) | JOIN_ON / P01 / COMPARISON | `Trx.SecId` : `INT` | `t.SecId` : `varchar(20)` | `t.SecId` → `INT` | `Trx.SecId = t.SecId` |
| IC-006773 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessBuyBackOM:146](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBuyBackOM.sql#L146) | JOIN_ON / P01 / COMPARISON | `Trx.SecId` : `INT` | `t.SecId` : `varchar(20)` | `t.SecId` → `INT` | `Trx.SecId = t.SecId` |
| IC-006777 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessBuyBackOM:155](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBuyBackOM.sql#L155) | JOIN_ON / P01 / COMPARISON | `Trx.SecId` : `INT` | `t.SecId` : `varchar(20)` | `t.SecId` → `INT` | `Trx.SecId = t.SecId` |
| IC-006797 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessBuyBackOM:330](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBuyBackOM.sql#L330) | JOIN_ON / P01 / COMPARISON | `Trx.SecId` : `INT` | `t.SecId` : `varchar(20)` | `t.SecId` → `INT` | `Trx.SecId=t.SecId` |
| IC-012952 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_CancelTopUpOriOrder_Temp:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_CancelTopUpOriOrder_Temp.sql#L64) | JOIN_ON / P01 / COMPARISON | `Trx.SecId` : `INT` | `t.SecId` : `varchar(20)` | `t.SecId` → `INT` | `Trx.SecId = t.SecId` |
| IC-013972 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_updateStatusORI_Order_Temp:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateStatusORI_Order_Temp.sql#L106) | JOIN_ON / P01 / COMPARISON | `Trx.SecId` : `INT` | `t.SecId` : `varchar(20)` | `t.SecId` → `INT` | `Trx.SecId = t.SecId` |
| IC-013974 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_updateStatusORI_Order_Temp:120](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateStatusORI_Order_Temp.sql#L120) | JOIN_ON / P01 / COMPARISON | `Trx.SecId` : `INT` | `t.SecId` : `varchar(20)` | `t.SecId` → `INT` | `Trx.SecId = t.SecId` |
| IC-013977 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_updateStatusORI_Order_Temp:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateStatusORI_Order_Temp.sql#L137) | JOIN_ON / P01 / COMPARISON | `Trx.SecId` : `INT` | `t.SecId` : `varchar(20)` | `t.SecId` → `INT` | `Trx.SecId = t.SecId` |
| IC-013981 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_updateStatusORI_Order_Temp:176](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateStatusORI_Order_Temp.sql#L176) | JOIN_ON / P01 / COMPARISON | `Trx.SecId` : `INT` | `t.SecId` : `varchar(20)` | `t.SecId` → `INT` | `Trx.SecId=t.SecId` |
| IC-013984 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_updateStatusORI_Order_Temp:186](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateStatusORI_Order_Temp.sql#L186) | JOIN_ON / P01 / COMPARISON | `Trx.SecId` : `INT` | `t.SecId` : `varchar(20)` | `t.SecId` → `INT` | `Trx.SecId=t.SecId` |
| IC-013989 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_updateStatusORI_Order_Temp:199](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateStatusORI_Order_Temp.sql#L199) | JOIN_ON / P01 / COMPARISON | `st.SecId` : `INT` | `t.SecId` : `varchar(20)` | `t.SecId` → `INT` | `st.SecId=t.SecId` |
| IC-013996 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_updateTopUpORI_Order_Temp:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateTopUpORI_Order_Temp.sql#L71) | JOIN_ON / P01 / COMPARISON | `Trx.SecId` : `INT` | `t.SecId` : `varchar(20)` | `t.SecId` → `INT` | `Trx.SecId = t.SecId` |
| IC-014001 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_updateTopUpORI_Order_Temp:87](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateTopUpORI_Order_Temp.sql#L87) | JOIN_ON / P01 / COMPARISON | `Trx.SecId` : `INT` | `t.SecId` : `varchar(20)` | `t.SecId` → `INT` | `Trx.SecId = t.SecId` |
| IC-014005 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_updateTopUpORI_Order_Temp:132](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateTopUpORI_Order_Temp.sql#L132) | JOIN_ON / P01 / COMPARISON | `Trx.SecId` : `INT` | `t.SecId` : `varchar(20)` | `t.SecId` → `INT` | `Trx.SecId=t.SecId` |
| IC-014011 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_updateTopUpORI_Order_Temp:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateTopUpORI_Order_Temp.sql#L145) | JOIN_ON / P01 / COMPARISON | `st.SecId` : `INT` | `t.SecId` : `varchar(20)` | `t.SecId` → `INT` | `st.SecId=t.SecId` |
