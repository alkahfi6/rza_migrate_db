# Retail UAT.dbo.ORI_Order_Log_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 16 lokasi source, 10 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-operationcode"></a>
## OperationCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_Log_TH.sql#L27). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 8 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005361 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDeleteORIOrder:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDeleteORIOrder.sql#L98) | WHERE / P12 / COMPARISON | `oolth.OperationCode` : `CHAR(1)` | `'D'` : `varchar` | `oolth.OperationCode` → `VARCHAR` | `oolth.OperationCode = 'D'` |
| IC-004824 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSORIOnlineNewTransaction:550](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlineNewTransaction.sql#L550) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_order_log_th.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.ori_order_log_th.OperationCode <- 'I'` |
| IC-006549 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessAllotment:143](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessAllotment.sql#L143) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_order_log_th.OperationCode` : `CHAR(1)` | `'L'` : `varchar` | `'L'` → `CHAR(1)` | `retail.dbo.ori_order_log_th.OperationCode <- 'L'` |
| IC-006807 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessDeleteORIOrder:148](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessDeleteORIOrder.sql#L148) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_order_log_th.OperationCode` : `CHAR(1)` | `'D'` : `varchar` | `'D'` → `CHAR(1)` | `retail.dbo.ori_order_log_th.OperationCode <- 'D'` |
| IC-006906 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessInsertORIOrderCommit:241](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertORIOrderCommit.sql#L241) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_order_log_th.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.ori_order_log_th.OperationCode <- 'I'` |
| IC-012248 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSplitOrder:233](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSplitOrder.sql#L233) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_order_log_th.OperationCode` : `CHAR(1)` | `'S'` : `varchar` | `'S'` → `CHAR(1)` | `retail.dbo.ori_order_log_th.OperationCode <- 'S'` |
| IC-012523 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateOrderPriority:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateOrderPriority.sql#L162) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_order_log_th.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `'U'` → `CHAR(1)` | `retail.dbo.ori_order_log_th.OperationCode <- 'U'` |
| IC-013106 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertORI_Order_TT_Original:631](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORI_Order_TT_Original.sql#L631) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_order_log_th.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.ori_order_log_th.OperationCode <- 'I'` |

<a id="kolom-processcode"></a>
## ProcessCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_Log_TH.sql#L28). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005364 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDeleteORIOrder:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDeleteORIOrder.sql#L113) | WHERE / P12 / COMPARISON | `oolth.ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `oolth.ProcessCode` → `VARCHAR` | `oolth.ProcessCode <> 'P'` |
| IC-005367 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDeleteORIOrder:190](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDeleteORIOrder.sql#L190) | WHERE / P12 / COMPARISON | `ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `ProcessCode` → `VARCHAR` | `ProcessCode = 'P'` |
| IC-012957 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_CheckDiffNIKOrderHapusOriInpAppv:15](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_CheckDiffNIKOrderHapusOriInpAppv.sql#L15) | WHERE / P12 / COMPARISON | `oolth.ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `oolth.ProcessCode` → `VARCHAR` | `oolth.ProcessCode = 'P'` |
| IC-006808 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessDeleteORIOrder:148](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessDeleteORIOrder.sql#L148) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_order_log_th.ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.ori_order_log_th.ProcessCode <- 'P'` |
| IC-006812 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessDeleteORIOrder:230](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessDeleteORIOrder.sql#L230) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_order_log_th.ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.ori_order_log_th.ProcessCode <- 'A'` |
| IC-006814 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessDeleteORIOrder:305](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessDeleteORIOrder.sql#L305) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_order_log_th.ProcessCode` : `CHAR(1)` | `'R'` : `varchar` | `'R'` → `CHAR(1)` | `retail.dbo.ori_order_log_th.ProcessCode <- 'R'` |
| IC-006818 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessDeleteORIOrderCommit:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessDeleteORIOrderCommit.sql#L64) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_order_log_th.ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.ori_order_log_th.ProcessCode <- 'A'` |
| IC-012524 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateOrderPriority:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateOrderPriority.sql#L162) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_order_log_th.ProcessCode` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.ori_order_log_th.ProcessCode <- 'O'` |
