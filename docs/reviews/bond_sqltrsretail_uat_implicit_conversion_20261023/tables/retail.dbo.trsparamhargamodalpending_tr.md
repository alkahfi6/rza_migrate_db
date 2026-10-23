# Retail UAT.dbo.TRSParamHargaModalPending_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 10 lokasi source, 6 procedure, 6 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-approvalnik"></a>
## ApprovalNIK

**Tipe dasar:** `VARCHAR(7)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModalPending_TR.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006666 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessApprovalHargaModal:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalHargaModal.sql#L71) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ApprovalNIK` : `VARCHAR(7)` | `@pnNik` : `INT` | `@pnNik` → `VARCHAR(7)` | `ApprovalNIK = @pnNik` |

<a id="kolom-inputernik"></a>
## InputerNIK

**Tipe dasar:** `VARCHAR(7)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModalPending_TR.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005265 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSPopulateApprovalHargaModal:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalHargaModal.sql#L30) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `InputerNIK` : `VARCHAR(7)` | `retail.dbo.fngetnamebynik.@nNIK` : `int` | `InputerNIK` → `int` | `dbo.fnGetNameByNIK(InputerNIK)` |
| IC-002605 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSSaveParamHargaModalNew_Normal:383](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSaveParamHargaModalNew_Normal.sql#L383) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamhargamodalpending_tr.InputerNIK` : `VARCHAR(7)` | `@pnNik` : `int` | `@pnNik` → `VARCHAR(7)` | `retail.dbo.trsparamhargamodalpending_tr.InputerNIK <- @pnNik` |
| IC-002684 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSSaveParamHargaModalNew_Panic:283](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSaveParamHargaModalNew_Panic.sql#L283) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamhargamodalpending_tr.InputerNIK` : `VARCHAR(7)` | `@pnNik` : `int` | `@pnNik` → `VARCHAR(7)` | `retail.dbo.trsparamhargamodalpending_tr.InputerNIK <- @pnNik` |
| IC-011732 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveParamHargaModal:407](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModal.sql#L407) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamhargamodalpending_tr.InputerNIK` : `VARCHAR(7)` | `@pnNik` : `int` | `@pnNik` → `VARCHAR(7)` | `retail.dbo.trsparamhargamodalpending_tr.InputerNIK <- @pnNik` |
| IC-011855 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveParamHargaModalNew:526](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L526) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamhargamodalpending_tr.InputerNIK` : `VARCHAR(7)` | `@pnNik` : `int` | `@pnNik` → `VARCHAR(7)` | `retail.dbo.trsparamhargamodalpending_tr.InputerNIK <- @pnNik` |

<a id="kolom-kuotatrxsekunder"></a>
## KuotaTrxSekunder

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModalPending_TR.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006668 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessApprovalHargaModal:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalHargaModal.sql#L89) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hmp.KuotaTrxSekunder` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hmp.KuotaTrxSekunder,0)` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `VARCHAR(3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModalPending_TR.sql#L24). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006667 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessApprovalHargaModal:74](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalHargaModal.sql#L74) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[Status]` : `VARCHAR(3)` | `@pcProcessType` : `CHAR(1)` | `@pcProcessType` → `VARCHAR(3)` | `[Status] = @pcProcessType` |

<a id="kolom-yieldbeli"></a>
## YieldBeli

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModalPending_TR.sql#L26). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005263 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateApprovalHargaModal:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalHargaModal.sql#L17) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `YieldBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(YieldBeli, 0)` |

<a id="kolom-yieldjual"></a>
## YieldJual

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModalPending_TR.sql#L27). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005264 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateApprovalHargaModal:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalHargaModal.sql#L21) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `YieldJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(YieldJual,0)` |
