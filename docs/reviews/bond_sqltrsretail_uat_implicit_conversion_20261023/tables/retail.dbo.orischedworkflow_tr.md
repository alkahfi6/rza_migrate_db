# Retail UAT.dbo.ORISchedWorkflow_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 10 lokasi source, 3 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-dateinsert"></a>
## DateInsert

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISchedWorkflow_TR.sql#L8). **Existing leading index:** ORISchedWorkflow_TR_NN2.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005472 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateGagalBayarOri:480](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L480) | WHERE / P04 / COMPARISON | `sched.DateInsert` : `DATETIME` | `convert(varchar, getdate(), 112)` : `varchar` | `convert(varchar, getdate(), 112)` → `DATETIME` | `sched.DateInsert > convert(varchar, getdate(), 112)` |

<a id="kolom-guidprocess"></a>
## GuidProcess

**Tipe dasar:** `UNIQUEIDENTIFIER`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISchedWorkflow_TR.sql#L2). **Existing leading index:** ORISchedWorkflow_TR_CN1&#124;ORISchedWorkflow_TR_NN3.

**Frekuensi source:** 8 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007981 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRePushGagalBayarOri:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRePushGagalBayarOri.sql#L49) | WHERE / P13 / COMPARISON | `GuidProcess` : `UNIQUEIDENTIFIER` | `@pcGuidProcess` : `NVARCHAR(50)` | `@pcGuidProcess` → `UNIQUEIDENTIFIER` | `GuidProcess = @pcGuidProcess` |
| IC-007982 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRePushGagalBayarOri:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRePushGagalBayarOri.sql#L59) | WHERE / P13 / COMPARISON | `GuidProcess` : `UNIQUEIDENTIFIER` | `@pcGuidProcess` : `NVARCHAR(50)` | `@pcGuidProcess` → `UNIQUEIDENTIFIER` | `GuidProcess = @pcGuidProcess` |
| IC-007983 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRePushGagalBayarOri:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRePushGagalBayarOri.sql#L96) | WHERE / P13 / COMPARISON | `[GuidProcess]` : `UNIQUEIDENTIFIER` | `@pcGuidProcess` : `NVARCHAR(50)` | `@pcGuidProcess` → `UNIQUEIDENTIFIER` | `[GuidProcess] = @pcGuidProcess` |
| IC-007984 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRePushGagalBayarOri:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRePushGagalBayarOri.sql#L105) | WHERE / P13 / COMPARISON | `GuidProcess` : `UNIQUEIDENTIFIER` | `@pcGuidProcess` : `NVARCHAR(50)` | `@pcGuidProcess` → `UNIQUEIDENTIFIER` | `GuidProcess = @pcGuidProcess` |
| IC-007985 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRePushGagalBayarOri:134](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRePushGagalBayarOri.sql#L134) | WHERE / P13 / COMPARISON | `[GuidProcess]` : `UNIQUEIDENTIFIER` | `@pcGuidProcess` : `NVARCHAR(50)` | `@pcGuidProcess` → `UNIQUEIDENTIFIER` | `[GuidProcess] = @pcGuidProcess` |
| IC-007986 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRePushGagalBayarOri:143](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRePushGagalBayarOri.sql#L143) | WHERE / P13 / COMPARISON | `GuidProcess` : `UNIQUEIDENTIFIER` | `@pcGuidProcess` : `NVARCHAR(50)` | `@pcGuidProcess` → `UNIQUEIDENTIFIER` | `GuidProcess = @pcGuidProcess` |
| IC-007987 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRePushGagalBayarOri:172](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRePushGagalBayarOri.sql#L172) | WHERE / P13 / COMPARISON | `[GuidProcess]` : `UNIQUEIDENTIFIER` | `@pcGuidProcess` : `NVARCHAR(50)` | `@pcGuidProcess` → `UNIQUEIDENTIFIER` | `[GuidProcess] = @pcGuidProcess` |
| IC-007988 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRePushGagalBayarOri:181](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRePushGagalBayarOri.sql#L181) | WHERE / P13 / COMPARISON | `GuidProcess` : `UNIQUEIDENTIFIER` | `@pcGuidProcess` : `NVARCHAR(50)` | `@pcGuidProcess` → `UNIQUEIDENTIFIER` | `GuidProcess = @pcGuidProcess` |

<a id="kolom-workflowstatus"></a>
## WorkflowStatus

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISchedWorkflow_TR.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000947 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.ORIOnlineReleaseFailedUnblock:10](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORIOnlineReleaseFailedUnblock.sql#L10) | WHERE / P04 / COMPARISON | `WorkflowStatus` : `INT` | `'99'` : `varchar` | `'99'` → `INT` | `WorkflowStatus = '99'` |
