# Retail UAT.dbo.process_table — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 20 lokasi source, 6 procedure, 4 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 3 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-process-group"></a>
## process_group

**Tipe dasar:** `TINYINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.process_table.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008701 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailBeginOfDay:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailBeginOfDay.sql#L49) | WHERE / P11 / COMPARISON | `process_group` : `TINYINT` | `1` : `int` | `process_group` → `INT` | `process_group!=1` |
| IC-008980 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailEndOfDay:116](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailEndOfDay.sql#L116) | WHERE / P11 / COMPARISON | `process_group` : `TINYINT` | `1` : `int` | `process_group` → `INT` | `process_group=1` |
| IC-012926 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.process_check_point:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.process_check_point.sql#L22) | WHERE / P11 / IN_VALUE_TYPE | `process_group` : `TINYINT` | `7` : `int` | `process_group` → `INT` | `process_group in (7,1,2)` |

<a id="kolom-process-type"></a>
## process_type

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.process_table.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012934 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.set_process_table:109](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.set_process_table.sql#L109) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `process_type` : `SMALLINT` | `@nProcessType` : `int` | `process_type` → `int` | `@nProcessType = process_type` |
| IC-012925 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.process_check_point:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.process_check_point.sql#L21) | WHERE / P11 / COMPARISON | `process_type` : `SMALLINT` | `@nProcessType` : `int` | `process_type` → `INT` | `process_type=@nProcessType` |

<a id="kolom-product-code"></a>
## product_code

**Tipe dasar:** `CHAR(3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.process_table.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012905 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.auto_process_timer:73](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.auto_process_timer.sql#L73) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `b.product_code` : `CHAR(3)` | `''` : `varchar` | `''` → `CHAR(3)` | `isnull(b.product_code,'')` |
| IC-012907 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.auto_process_timer:79](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.auto_process_timer.sql#L79) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `b.product_code` : `CHAR(3)` | `''` : `varchar` | `''` → `CHAR(3)` | `isnull(b.product_code,'')` |
| IC-012908 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.auto_process_timer:79](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.auto_process_timer.sql#L79) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `b.product_code` : `CHAR(3)` | `''` : `varchar` | `''` → `CHAR(3)` | `isnull(b.product_code,'')` |
| IC-012935 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.set_process_table:118](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.set_process_table.sql#L118) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `product_code` : `CHAR(3)` | `'$$$'` : `varchar` | `'$$$'` → `CHAR(3)` | `isnull(product_code, '$$$')` |

<a id="kolom-success-bit"></a>
## success_bit

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.process_table.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 11 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012903 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.auto_process_timer:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.auto_process_timer.sql#L70) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `b.success_bit` : `BIT` | `ready_bit` : `int` | `b.success_bit` → `int` | `ready_bit=b.success_bit` |
| IC-012906 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.auto_process_timer:76](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.auto_process_timer.sql#L76) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `b.success_bit` : `BIT` | `ready_bit` : `int` | `b.success_bit` → `int` | `ready_bit=b.success_bit` |
| IC-012909 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.auto_process_timer:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.auto_process_timer.sql#L80) | WHERE / P11 / COMPARISON | `b.success_bit` : `BIT` | `0` : `int` | `b.success_bit` → `INT` | `b.success_bit=0` |
| IC-012917 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.job_get_current_process_group:38](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.job_get_current_process_group.sql#L38) | WHERE / P11 / COMPARISON | `success_bit` : `BIT` | `1` : `int` | `success_bit` → `INT` | `success_bit=1` |
| IC-012920 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.job_get_current_process_group:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.job_get_current_process_group.sql#L43) | WHERE / P11 / COMPARISON | `success_bit` : `BIT` | `1` : `int` | `success_bit` → `INT` | `success_bit=1` |
| IC-008700 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailBeginOfDay:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailBeginOfDay.sql#L48) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `success_bit` : `BIT` | `0` : `int` | `0` → `BIT` | `success_bit=0` |
| IC-008979 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailEndOfDay:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailEndOfDay.sql#L115) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `success_bit` : `BIT` | `0` : `int` | `0` → `BIT` | `success_bit=0` |
| IC-012930 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.set_process_table:68](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.set_process_table.sql#L68) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `success_bit` : `BIT` | `0` : `int` | `0` → `BIT` | `success_bit = 0` |
| IC-012931 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.set_process_table:75](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.set_process_table.sql#L75) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `success_bit` : `BIT` | `0` : `int` | `0` → `BIT` | `success_bit = 0` |
| IC-012932 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.set_process_table:91](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.set_process_table.sql#L91) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `success_bit` : `BIT` | `1` : `int` | `1` → `BIT` | `success_bit = 1` |
| IC-012933 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.set_process_table:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.set_process_table.sql#L98) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `success_bit` : `BIT` | `1` : `int` | `1` → `BIT` | `success_bit = 1` |
