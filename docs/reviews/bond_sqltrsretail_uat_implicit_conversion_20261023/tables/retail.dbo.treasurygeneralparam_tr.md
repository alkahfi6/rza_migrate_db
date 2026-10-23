# Retail UAT.dbo.TreasuryGeneralParam_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 8 lokasi source, 6 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 2 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-groupid"></a>
## GroupId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryGeneralParam_TR.sql#L2). **Existing leading index:** TreasuryGeneralParam_TR_CU1.

**Frekuensi source:** 5 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003665 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCustodyUploadJobs:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCustodyUploadJobs.sql#L25) | WHERE / P11 / COMPARISON | `GroupId` : `BIGINT` | `5000` : `int` | `5000` → `BIGINT` | `GroupId = 5000` |
| IC-003674 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCustodyUploadValidateCutoff:15](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCustodyUploadValidateCutoff.sql#L15) | WHERE / P11 / COMPARISON | `GroupId` : `BIGINT` | `4000` : `int` | `4000` → `BIGINT` | `GroupId = 4000` |
| IC-008152 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportParamReportType:6](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportParamReportType.sql#L6) | WHERE / P11 / COMPARISON | `GroupId` : `BIGINT` | `1000` : `int` | `1000` → `BIGINT` | `GroupId = 1000` |
| IC-010316 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptMonitoringDocTrx:264](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptMonitoringDocTrx.sql#L264) | WHERE / P11 / COMPARISON | `GroupId` : `BIGINT` | `3000` : `int` | `3000` → `BIGINT` | `GroupId = 3000` |
| IC-012723 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSValidateSecAccNo:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateSecAccNo.sql#L34) | WHERE / P11 / COMPARISON | `GroupId` : `BIGINT` | `2000` : `int` | `2000` → `BIGINT` | `GroupId = 2000` |

<a id="kolom-paramdesc"></a>
## ParamDesc

**Tipe dasar:** `VARCHAR(100)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryGeneralParam_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003664 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSCustodyUploadJobs:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCustodyUploadJobs.sql#L23) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `ParamDesc` : `VARCHAR(100)` | `@nMonthKeep` : `int` | `ParamDesc` → `int` | `@nMonthKeep = ParamDesc` |
| IC-012722 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSValidateSecAccNo:32](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateSecAccNo.sql#L32) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `ParamDesc` : `VARCHAR(100)` | `@dLastUpdateRisk` : `datetime` | `ParamDesc` → `datetime` | `@dLastUpdateRisk = ParamDesc` |

<a id="kolom-paramid"></a>
## ParamId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryGeneralParam_TR.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-010306 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptMonitoringDocAcc:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptMonitoringDocAcc.sql#L180) | WHERE / P11 / COMPARISON | `ParamId` : `BIGINT` | `2001` : `int` | `2001` → `BIGINT` | `ParamId = 2001` |
