# Retail UAT.dbo.TRSEditHargaModal_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 21 lokasi source, 8 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-approvalnik"></a>
## ApprovalNIK

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSEditHargaModal_TH.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006659 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessApprovalEditHargaModal:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalEditHargaModal.sql#L41) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `aatm.ApprovalNIK` : `BIGINT` | `@pnUserNik` : `INT` | `@pnUserNik` → `BIGINT` | `aatm.ApprovalNIK = @pnUserNik` |
| IC-006661 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessApprovalEditHargaModal:62](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalEditHargaModal.sql#L62) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `aatm.ApprovalNIK` : `BIGINT` | `@pnUserNik` : `INT` | `@pnUserNik` → `BIGINT` | `aatm.ApprovalNIK = @pnUserNik` |

<a id="kolom-inputernik"></a>
## InputerNIK

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSEditHargaModal_TH.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006839 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessEditHargaModalTrx:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessEditHargaModalTrx.sql#L65) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsedithargamodal_th.InputerNIK` : `BIGINT` | `@pnNik` : `int` | `@pnNik` → `BIGINT` | `retail.dbo.trsedithargamodal_th.InputerNIK <- @pnNik` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSEditHargaModal_TH.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 18 lokasi / 8 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000743 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.OMSubmitTransactionSekunder:446](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMSubmitTransactionSekunder.sql#L446) | WHERE / P12 / IN_VALUE_TYPE | `isnull(tehm.Status,'')` : `CHAR(1)` | `'N'` : `varchar` | `isnull(tehm.Status,'')` → `VARCHAR` | `isnull(tehm.Status,'') not in ('N','O')` |
| IC-000907 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.ONEApprovalPrepareData:132](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareData.sql#L132) | WHERE / P12 / IN_VALUE_TYPE | `ISNULL(tehm.[Status],'')` : `CHAR(1)` | `'N'` : `varchar` | `ISNULL(tehm.[Status],'')` → `VARCHAR` | `ISNULL(tehm.[Status],'') not in ('N','O')` |
| IC-000935 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.ONEApprovalPrepareDataBond:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareDataBond.sql#L130) | WHERE / P12 / IN_VALUE_TYPE | `ISNULL(tehm.[Status],'')` : `CHAR(1)` | `'N'` : `varchar` | `ISNULL(tehm.[Status],'')` → `VARCHAR` | `ISNULL(tehm.[Status],'') not in ('N','O')` |
| IC-005406 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateEditHargaModalAuthorization:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEditHargaModalAuthorization.sql#L18) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'N'` : `varchar` | `Status` → `VARCHAR` | `Status = 'N'` |
| IC-006658 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprovalEditHargaModal:38](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalEditHargaModal.sql#L38) | WHERE / P12 / COMPARISON | `aatm.Status` : `CHAR(1)` | `'N'` : `varchar` | `aatm.Status` → `VARCHAR` | `aatm.Status = 'N'` |
| IC-006660 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprovalEditHargaModal:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalEditHargaModal.sql#L47) | WHERE / P12 / COMPARISON | `aatm.Status` : `CHAR(1)` | `'N'` : `varchar` | `aatm.Status` → `VARCHAR` | `aatm.Status = 'N'` |
| IC-006662 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprovalEditHargaModal:68](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalEditHargaModal.sql#L68) | WHERE / P12 / COMPARISON | `aatm.Status` : `CHAR(1)` | `'N'` : `varchar` | `aatm.Status` → `VARCHAR` | `aatm.Status = 'N'` |
| IC-006948 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSProcessInsertSecurityTransactionCommit:717](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertSecurityTransactionCommit.sql#L717) | WHERE / P12 / IN_VALUE_TYPE | `isnull(tehm.Status,'')` : `CHAR(1)` | `'N'` : `varchar` | `isnull(tehm.Status,'')` → `VARCHAR` | `isnull(tehm.Status,'') not in ('N','O')` |
| IC-013739 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_StatusSecurityTransaction_TT:198](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_StatusSecurityTransaction_TT.sql#L198) | WHERE / P12 / IN_VALUE_TYPE | `isnull(tehm.Status,'')` : `CHAR(1)` | `'N'` : `varchar` | `isnull(tehm.Status,'')` → `VARCHAR` | `isnull(tehm.Status,'') not in ('N','O')` |
| IC-013757 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_StatusSecurityTransaction_TT:350](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_StatusSecurityTransaction_TT.sql#L350) | WHERE / P12 / IN_VALUE_TYPE | `isnull(tehm.Status,'')` : `CHAR(1)` | `'N'` : `varchar` | `isnull(tehm.Status,'')` → `VARCHAR` | `isnull(tehm.Status,'') not in ('N','O')` |
| IC-000744 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.OMSubmitTransactionSekunder:446](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMSubmitTransactionSekunder.sql#L446) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tehm.Status` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(tehm.Status,'')` |
| IC-000908 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.ONEApprovalPrepareData:132](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareData.sql#L132) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tehm.[Status]` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `ISNULL(tehm.[Status],'')` |
| IC-000936 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.ONEApprovalPrepareDataBond:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareDataBond.sql#L130) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tehm.[Status]` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `ISNULL(tehm.[Status],'')` |
| IC-006838 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessEditHargaModalTrx:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessEditHargaModalTrx.sql#L61) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Status` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `Status = 'O'` |
| IC-006840 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessEditHargaModalTrx:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessEditHargaModalTrx.sql#L65) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsedithargamodal_th.Status` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.trsedithargamodal_th.Status <- 'N'` |
| IC-006949 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessInsertSecurityTransactionCommit:717](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertSecurityTransactionCommit.sql#L717) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tehm.Status` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(tehm.Status,'')` |
| IC-013740 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_StatusSecurityTransaction_TT:198](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_StatusSecurityTransaction_TT.sql#L198) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tehm.Status` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(tehm.Status,'')` |
| IC-013758 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_StatusSecurityTransaction_TT:350](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_StatusSecurityTransaction_TT.sql#L350) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tehm.Status` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(tehm.Status,'')` |
