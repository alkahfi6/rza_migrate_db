# Retail UAT.dbo.ORI_Order_Pending_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 13 lokasi source, 5 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-lastprocesscode"></a>
## LastProcessCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_Pending_TT.sql#L26). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006801 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessDeleteORIOrder:58](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessDeleteORIOrder.sql#L58) | JOIN_ON / P12 / COMPARISON | `ooptt.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `ooptt.LastProcessCode` → `VARCHAR` | `ooptt.LastProcessCode = 'P'` |
| IC-006803 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessDeleteORIOrder:72](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessDeleteORIOrder.sql#L72) | JOIN_ON / P12 / COMPARISON | `ooptt.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `ooptt.LastProcessCode` → `VARCHAR` | `ooptt.LastProcessCode = 'P'` |
| IC-012956 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_CheckDiffNIKOrderHapusOriInpAppv:14](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_CheckDiffNIKOrderHapusOriInpAppv.sql#L14) | WHERE / P12 / COMPARISON | `ooptt.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `ooptt.LastProcessCode` → `VARCHAR` | `ooptt.LastProcessCode = 'P'` |
| IC-013942 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ValidateInsertBuyBackOM:88](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateInsertBuyBackOM.sql#L88) | WHERE / P12 / COMPARISON | `LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `LastProcessCode` → `VARCHAR` | `LastProcessCode = 'P'` |
| IC-013958 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ValidatePerubahanOrderSuratBerharga:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidatePerubahanOrderSuratBerharga.sql#L49) | WHERE / P12 / COMPARISON | `LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `LastProcessCode` → `VARCHAR` | `LastProcessCode = 'P'` |
| IC-006804 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessDeleteORIOrder:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessDeleteORIOrder.sql#L89) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_order_pending_tt.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.ori_order_pending_tt.LastProcessCode <- 'P'` |
| IC-006811 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessDeleteORIOrder:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessDeleteORIOrder.sql#L225) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `LastProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `LastProcessCode = 'A'` |
| IC-006813 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessDeleteORIOrder:299](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessDeleteORIOrder.sql#L299) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `LastProcessCode` : `CHAR(1)` | `'R'` : `varchar` | `'R'` → `CHAR(1)` | `LastProcessCode = 'R'` |
| IC-006817 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessDeleteORIOrderCommit:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessDeleteORIOrderCommit.sql#L59) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `LastProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `LastProcessCode = 'A'` |

<a id="kolom-operationcode"></a>
## OperationCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Order_Pending_TT.sql#L25). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012955 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_CheckDiffNIKOrderHapusOriInpAppv:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_CheckDiffNIKOrderHapusOriInpAppv.sql#L13) | WHERE / P12 / COMPARISON | `ooptt.OperationCode` : `CHAR(1)` | `'D'` : `varchar` | `ooptt.OperationCode` → `VARCHAR` | `ooptt.OperationCode = 'D'` |
| IC-013943 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ValidateInsertBuyBackOM:88](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateInsertBuyBackOM.sql#L88) | WHERE / P12 / COMPARISON | `OperationCode` : `CHAR(1)` | `'D'` : `varchar` | `OperationCode` → `VARCHAR` | `OperationCode = 'D'` |
| IC-013959 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ValidatePerubahanOrderSuratBerharga:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidatePerubahanOrderSuratBerharga.sql#L49) | WHERE / P12 / COMPARISON | `OperationCode` : `CHAR(1)` | `'D'` : `varchar` | `OperationCode` → `VARCHAR` | `OperationCode = 'D'` |
| IC-006805 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessDeleteORIOrder:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessDeleteORIOrder.sql#L89) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_order_pending_tt.OperationCode` : `CHAR(1)` | `'D'` : `varchar` | `'D'` → `CHAR(1)` | `retail.dbo.ori_order_pending_tt.OperationCode <- 'D'` |
