# Retail UAT.dbo.TRSParameterGL_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 30 lokasi source, 4 procedure, 4 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-approvalnik"></a>
## ApprovalNIK

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParameterGL_TH.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003284 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:140](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L140) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ApprovalNIK` : `BIGINT` | `@pnProcessNIK` : `int` | `@pnProcessNIK` → `BIGINT` | `ApprovalNIK = @pnProcessNIK` |
| IC-003287 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L147) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ApprovalNIK` : `BIGINT` | `@pnProcessNIK` : `int` | `@pnProcessNIK` → `BIGINT` | `ApprovalNIK = @pnProcessNIK` |
| IC-003290 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:154](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L154) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ApprovalNIK` : `BIGINT` | `@pnProcessNIK` : `int` | `@pnProcessNIK` → `BIGINT` | `ApprovalNIK = @pnProcessNIK` |
| IC-003293 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L165) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ApprovalNIK` : `BIGINT` | `@pnProcessNIK` : `int` | `@pnProcessNIK` → `BIGINT` | `ApprovalNIK = @pnProcessNIK` |

<a id="kolom-inputernik"></a>
## InputerNIK

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParameterGL_TH.sql#L19). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007255 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessParamGL:95](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessParamGL.sql#L95) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergl_th.InputerNIK` : `BIGINT` | `@pnNik` : `int` | `@pnNik` → `BIGINT` | `retail.dbo.trsparametergl_th.InputerNIK <- @pnNik` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParameterGL_TH.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 24 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003263 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamGL:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L48) | WHERE / P12 / COMPARISON | `th.Status` : `CHAR(1)` | `'I'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status = 'I'` |
| IC-003268 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamGL:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L65) | WHERE / P12 / COMPARISON | `th.Status` : `CHAR(1)` | `'I'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status = 'I'` |
| IC-003269 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamGL:88](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L88) | WHERE / P12 / COMPARISON | `th.Status` : `CHAR(1)` | `'U'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status = 'U'` |
| IC-003275 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamGL:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L105) | WHERE / P12 / COMPARISON | `th.Status` : `CHAR(1)` | `'U'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status = 'U'` |
| IC-003276 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamGL:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L113) | WHERE / P12 / COMPARISON | `th.Status` : `CHAR(1)` | `'D'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status = 'D'` |
| IC-003281 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamGL:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L130) | WHERE / P12 / COMPARISON | `th.Status` : `CHAR(1)` | `'D'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status = 'D'` |
| IC-003283 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamGL:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L137) | WHERE / P12 / COMPARISON | `th.Status` : `CHAR(1)` | `'N'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status = 'N'` |
| IC-003286 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamGL:144](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L144) | WHERE / P12 / COMPARISON | `th.Status` : `CHAR(1)` | `'I'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status = 'I'` |
| IC-003289 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamGL:151](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L151) | WHERE / P12 / COMPARISON | `th.Status` : `CHAR(1)` | `'U'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status = 'U'` |
| IC-003292 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamGL:158](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L158) | WHERE / P12 / COMPARISON | `th.Status` : `CHAR(1)` | `'D'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status='D'` |
| IC-003295 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamGL:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L169) | WHERE / P12 / IN_VALUE_TYPE | `th.Status` : `CHAR(1)` | `'I'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status in ('I','U','D')` |
| IC-003300 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamGL:187](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L187) | WHERE / P12 / IN_VALUE_TYPE | `th.Status` : `CHAR(1)` | `'I'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status in ('I','D')` |
| IC-003306 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamGL:205](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L205) | WHERE / P12 / COMPARISON | `th.Status` : `CHAR(1)` | `'U'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status = 'U'` |
| IC-005266 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprovalParameterGL:14](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalParameterGL.sql#L14) | WHERE / P12 / IN_VALUE_TYPE | `Status` : `CHAR(1)` | `'O'` : `varchar` | `Status` → `VARCHAR` | `Status not in ('O','N','R')` |
| IC-005966 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateParameterGL:14](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParameterGL.sql#L14) | WHERE / P12 / IN_VALUE_TYPE | `Status` : `CHAR(1)` | `'O'` : `varchar` | `Status` → `VARCHAR` | `Status not in ('O','R')` |
| IC-007254 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessParamGL:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessParamGL.sql#L89) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'D'` : `varchar` | `Status` → `VARCHAR` | `Status = 'D'` |
| IC-007261 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessParamGL:157](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessParamGL.sql#L157) | WHERE / P12 / COMPARISON | `th.Status` : `CHAR(1)` | `'N'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status = 'N'` |
| IC-007263 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessParamGL:166](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessParamGL.sql#L166) | WHERE / P12 / COMPARISON | `th.Status` : `CHAR(1)` | `'N'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status = 'N'` |
| IC-003282 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:133](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L133) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `th.Status` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `th.Status = 'O'` |
| IC-003285 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:140](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L140) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `th.Status` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `th.Status = 'N'` |
| IC-003288 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L147) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `th.Status` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `th.Status = 'N'` |
| IC-003291 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:154](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L154) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `th.Status` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `th.Status = 'O'` |
| IC-003294 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L165) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `th.Status` : `CHAR(1)` | `'R'` : `varchar` | `'R'` → `CHAR(1)` | `th.Status = 'R'` |
| IC-007262 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessParamGL:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessParamGL.sql#L162) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Status` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `Status = 'O'` |

<a id="kolom-updaternik"></a>
## UpdaterNIK

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParameterGL_TH.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007264 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessParamGL:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessParamGL.sql#L168) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergl_th.UpdaterNIK` : `BIGINT` | `@pnNik` : `int` | `@pnNik` → `BIGINT` | `retail.dbo.trsparametergl_th.UpdaterNIK <- @pnNik` |
