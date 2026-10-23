# Retail UAT.dbo.TRSParamAccruedDaysAdjustment_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 12 lokasi source, 3 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-actionflag"></a>
## ActionFlag

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamAccruedDaysAdjustment_TH.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004994 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSParamAccruedDaysSave:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamAccruedDaysSave.sql#L49) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamaccrueddaysadjustment_th.ActionFlag` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.trsparamaccrueddaysadjustment_th.ActionFlag <- 'A'` |
| IC-004996 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSParamAccruedDaysSave:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamAccruedDaysSave.sql#L59) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamaccrueddaysadjustment_th.ActionFlag` : `CHAR(1)` | `'D'` : `varchar` | `'D'` → `CHAR(1)` | `retail.dbo.trsparamaccrueddaysadjustment_th.ActionFlag <- 'D'` |
| IC-004998 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSParamAccruedDaysSave:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamAccruedDaysSave.sql#L69) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamaccrueddaysadjustment_th.ActionFlag` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.trsparamaccrueddaysadjustment_th.ActionFlag <- 'O'` |
| IC-005000 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSParamAccruedDaysSave:79](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamAccruedDaysSave.sql#L79) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamaccrueddaysadjustment_th.ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.trsparamaccrueddaysadjustment_th.ActionFlag <- 'N'` |

<a id="kolom-authstatus"></a>
## AuthStatus

**Tipe dasar:** `TINYINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamAccruedDaysAdjustment_TH.sql#L10). **Existing leading index:** TRSParamAccruedDaysAdjustment_TH_CN1.

**Frekuensi source:** 8 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004980 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamAccruedDaysAuthPopulate:20](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamAccruedDaysAuthPopulate.sql#L20) | WHERE / P11 / COMPARISON | `AuthStatus` : `TINYINT` | `0` : `int` | `AuthStatus` → `INT` | `AuthStatus = 0` |
| IC-004988 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamAccruedDaysAuthProcess:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamAccruedDaysAuthProcess.sql#L17) | WHERE / P11 / COMPARISON | `AuthStatus` : `TINYINT` | `0` : `int` | `AuthStatus` → `INT` | `AuthStatus = 0` |
| IC-004993 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamAccruedDaysSave:40](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamAccruedDaysSave.sql#L40) | WHERE / P11 / COMPARISON | `AuthStatus` : `TINYINT` | `0` : `int` | `AuthStatus` → `INT` | `AuthStatus = 0` |
| IC-004992 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSParamAccruedDaysAuthProcess:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamAccruedDaysAuthProcess.sql#L56) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `AuthStatus` : `TINYINT` | `case when @pbApproveBit = 1 then 1 else 99 end` : `INT (combined result; facets unknown)` | `case when @pbApproveBit = 1 then 1 else 99 end` → `TINYINT` | `AuthStatus = case when @pbApproveBit = 1 then 1 else 99 end` |
| IC-004995 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSParamAccruedDaysSave:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamAccruedDaysSave.sql#L49) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamaccrueddaysadjustment_th.AuthStatus` : `TINYINT` | `0` : `int` | `0` → `TINYINT` | `retail.dbo.trsparamaccrueddaysadjustment_th.AuthStatus <- 0` |
| IC-004997 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSParamAccruedDaysSave:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamAccruedDaysSave.sql#L59) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamaccrueddaysadjustment_th.AuthStatus` : `TINYINT` | `0` : `int` | `0` → `TINYINT` | `retail.dbo.trsparamaccrueddaysadjustment_th.AuthStatus <- 0` |
| IC-004999 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSParamAccruedDaysSave:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamAccruedDaysSave.sql#L69) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamaccrueddaysadjustment_th.AuthStatus` : `TINYINT` | `0` : `int` | `0` → `TINYINT` | `retail.dbo.trsparamaccrueddaysadjustment_th.AuthStatus <- 0` |
| IC-005001 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSParamAccruedDaysSave:79](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamAccruedDaysSave.sql#L79) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamaccrueddaysadjustment_th.AuthStatus` : `TINYINT` | `0` : `int` | `0` → `TINYINT` | `retail.dbo.trsparamaccrueddaysadjustment_th.AuthStatus <- 0` |
