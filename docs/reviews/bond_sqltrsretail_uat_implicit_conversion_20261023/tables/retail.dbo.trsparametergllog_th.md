# Retail UAT.dbo.TRSParameterGLLog_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 32 lokasi source, 3 procedure, 6 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-actionflag"></a>
## ACTIONFLAG

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParameterGLLog_TH.sql#L26). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009541 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParameterGL:190](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParameterGL.sql#L190) | JOIN_ON / P12 / COMPARISON | `tpgl.ACTIONFLAG` : `CHAR(1)` | `'U'` : `varchar` | `tpgl.ACTIONFLAG` → `VARCHAR` | `tpgl.ACTIONFLAG = 'U'` |
| IC-003264 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L49) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.ACTIONFLAG` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.trsparametergllog_th.ACTIONFLAG <- 'I'` |
| IC-003270 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L89) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.ACTIONFLAG` : `CHAR(1)` | `'U'` : `varchar` | `'U'` → `CHAR(1)` | `retail.dbo.trsparametergllog_th.ACTIONFLAG <- 'U'` |
| IC-003277 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.ACTIONFLAG` : `CHAR(1)` | `'D'` : `varchar` | `'D'` → `CHAR(1)` | `retail.dbo.trsparametergllog_th.ACTIONFLAG <- 'D'` |
| IC-003296 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:171](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L171) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.ACTIONFLAG` : `CHAR(1)` | `'D'` : `varchar` | `'D'` → `CHAR(1)` | `retail.dbo.trsparametergllog_th.ACTIONFLAG <- 'D'` |
| IC-003301 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:189](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L189) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.ACTIONFLAG` : `CHAR(1)` | `'U'` : `varchar` | `'U'` → `CHAR(1)` | `retail.dbo.trsparametergllog_th.ACTIONFLAG <- 'U'` |

<a id="kolom-apprstatus"></a>
## APPRSTATUS

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParameterGLLog_TH.sql#L27). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 10 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009539 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParameterGL:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParameterGL.sql#L127) | WHERE / P12 / IN_VALUE_TYPE | `APPRSTATUS` : `CHAR(1)` | `'A'` : `varchar` | `APPRSTATUS` → `VARCHAR` | `APPRSTATUS in ('A','R')` |
| IC-009540 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParameterGL:189](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParameterGL.sql#L189) | JOIN_ON / P12 / COMPARISON | `tpgl.APPRSTATUS` : `CHAR(1)` | `'P'` : `varchar` | `tpgl.APPRSTATUS` → `VARCHAR` | `tpgl.APPRSTATUS = 'P'` |
| IC-003265 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L49) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.APPRSTATUS` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.trsparametergllog_th.APPRSTATUS <- 'A'` |
| IC-003271 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L89) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.APPRSTATUS` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.trsparametergllog_th.APPRSTATUS <- 'A'` |
| IC-003278 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.APPRSTATUS` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.trsparametergllog_th.APPRSTATUS <- 'A'` |
| IC-003297 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:171](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L171) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.APPRSTATUS` : `CHAR(1)` | `'R'` : `varchar` | `'R'` → `CHAR(1)` | `retail.dbo.trsparametergllog_th.APPRSTATUS <- 'R'` |
| IC-003302 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:189](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L189) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.APPRSTATUS` : `CHAR(1)` | `'R'` : `varchar` | `'R'` → `CHAR(1)` | `retail.dbo.trsparametergllog_th.APPRSTATUS <- 'R'` |
| IC-007256 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessParamGL:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessParamGL.sql#L110) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.APPRSTATUS` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trsparametergllog_th.APPRSTATUS <- 'P'` |
| IC-007259 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessParamGL:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessParamGL.sql#L130) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.APPRSTATUS` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trsparametergllog_th.APPRSTATUS <- 'P'` |
| IC-007265 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessParamGL:185](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessParamGL.sql#L185) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.APPRSTATUS` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trsparametergllog_th.APPRSTATUS <- 'P'` |

<a id="kolom-approvalnik"></a>
## ApprovalNIK

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParameterGLLog_TH.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003266 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L49) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.ApprovalNIK` : `BIGINT` | `@pnProcessNIK` : `int` | `@pnProcessNIK` → `BIGINT` | `retail.dbo.trsparametergllog_th.ApprovalNIK <- @pnProcessNIK` |
| IC-003272 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L89) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.ApprovalNIK` : `BIGINT` | `@pnProcessNIK` : `int` | `@pnProcessNIK` → `BIGINT` | `retail.dbo.trsparametergllog_th.ApprovalNIK <- @pnProcessNIK` |
| IC-003279 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.ApprovalNIK` : `BIGINT` | `@pnProcessNIK` : `int` | `@pnProcessNIK` → `BIGINT` | `retail.dbo.trsparametergllog_th.ApprovalNIK <- @pnProcessNIK` |
| IC-003298 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:171](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L171) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.ApprovalNIK` : `BIGINT` | `@pnProcessNIK` : `int` | `@pnProcessNIK` → `BIGINT` | `retail.dbo.trsparametergllog_th.ApprovalNIK <- @pnProcessNIK` |
| IC-003303 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:189](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L189) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.ApprovalNIK` : `BIGINT` | `@pnProcessNIK` : `int` | `@pnProcessNIK` → `BIGINT` | `retail.dbo.trsparametergllog_th.ApprovalNIK <- @pnProcessNIK` |

<a id="kolom-detailupdate"></a>
## DETAILUPDATE

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParameterGLLog_TH.sql#L28). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003273 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L89) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.DETAILUPDATE` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.trsparametergllog_th.DETAILUPDATE <- 'N'` |
| IC-003304 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:189](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L189) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.DETAILUPDATE` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.trsparametergllog_th.DETAILUPDATE <- 'N'` |
| IC-007257 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessParamGL:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessParamGL.sql#L110) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.DETAILUPDATE` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.trsparametergllog_th.DETAILUPDATE <- 'O'` |

<a id="kolom-inputernik"></a>
## InputerNIK

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParameterGLLog_TH.sql#L19). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007258 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessParamGL:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessParamGL.sql#L110) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.InputerNIK` : `BIGINT` | `@pnNik` : `int` | `@pnNik` → `BIGINT` | `retail.dbo.trsparametergllog_th.InputerNIK <- @pnNik` |
| IC-007260 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessParamGL:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessParamGL.sql#L130) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.InputerNIK` : `BIGINT` | `@pnNik` : `int` | `@pnNik` → `BIGINT` | `retail.dbo.trsparametergllog_th.InputerNIK <- @pnNik` |
| IC-007266 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessParamGL:185](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessParamGL.sql#L185) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.InputerNIK` : `BIGINT` | `@pnNik` : `int` | `@pnNik` → `BIGINT` | `retail.dbo.trsparametergllog_th.InputerNIK <- @pnNik` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParameterGLLog_TH.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003267 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L49) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.Status` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.trsparametergllog_th.Status <- 'I'` |
| IC-003274 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L89) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.Status` : `CHAR(1)` | `'U'` : `varchar` | `'U'` → `CHAR(1)` | `retail.dbo.trsparametergllog_th.Status <- 'U'` |
| IC-003280 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.Status` : `CHAR(1)` | `'D'` : `varchar` | `'D'` → `CHAR(1)` | `retail.dbo.trsparametergllog_th.Status <- 'D'` |
| IC-003299 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:171](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L171) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.Status` : `CHAR(1)` | `'D'` : `varchar` | `'D'` → `CHAR(1)` | `retail.dbo.trsparametergllog_th.Status <- 'D'` |
| IC-003305 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamGL:189](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamGL.sql#L189) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparametergllog_th.Status` : `CHAR(1)` | `'U'` : `varchar` | `'U'` → `CHAR(1)` | `retail.dbo.trsparametergllog_th.Status <- 'U'` |
