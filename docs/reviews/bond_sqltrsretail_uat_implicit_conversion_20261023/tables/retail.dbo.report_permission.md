# Retail UAT.dbo.Report_Permission — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 19 lokasi source, 6 procedure, 2 kolom. **A:** 2 lokasi pada 1 procedure. **B:** 8 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-orderbranch"></a>
## OrderBranch

**Tipe dasar:** `VARCHAR(5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.Report_Permission.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 1 SP. **A:** 2 lokasi / 1 SP. **Peran:** CONVERTED_ARGUMENT, CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005986 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateReportPermission:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateReportPermission.sql#L46) | WHERE / P01 / COMPARISON | `OrderBranch` : `VARCHAR(5)` | `@nOrderBranch` : `int` | `OrderBranch` → `INT` | `OrderBranch=@nOrderBranch` |
| IC-005989 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateReportPermission:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateReportPermission.sql#L65) | WHERE / P01 / COMPARISON | `OrderBranch` : `VARCHAR(5)` | `@nOrderBranch` : `int` | `OrderBranch` → `INT` | `OrderBranch=@nOrderBranch` |
| IC-005982 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSPopulateReportPermission:19](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateReportPermission.sql#L19) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `OrderBranch` : `VARCHAR(5)` | `retail.dbo.fngetbranchcode.@cBranchCode` : `int` | `OrderBranch` → `int` | `dbo.fnGetBranchCode(OrderBranch)` |
| IC-005984 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSPopulateReportPermission:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateReportPermission.sql#L37) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `OrderBranch` : `VARCHAR(5)` | `retail.dbo.fngetbranchcode.@cBranchCode` : `int` | `OrderBranch` → `int` | `dbo.fnGetBranchCode(OrderBranch)` |
| IC-005987 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSPopulateReportPermission:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateReportPermission.sql#L56) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `OrderBranch` : `VARCHAR(5)` | `retail.dbo.fngetbranchcode.@cBranchCode` : `int` | `OrderBranch` → `int` | `dbo.fnGetBranchCode(OrderBranch)` |
| IC-005990 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSPopulateReportPermission:74](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateReportPermission.sql#L74) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `OrderBranch` : `VARCHAR(5)` | `retail.dbo.fngetbranchcode.@cBranchCode` : `int` | `OrderBranch` → `int` | `dbo.fnGetBranchCode(OrderBranch)` |

<a id="kolom-bitexec"></a>
## BitExec

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.Report_Permission.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 13 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009032 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailPopulateCA:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulateCA.sql#L56) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `BitExec` : `BIT` | `@iBitExec` : `bigint` | `BitExec` → `bigint` | `@iBitExec=BitExec` |
| IC-009036 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailPopulateCI:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulateCI.sql#L52) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `BitExec` : `BIT` | `@iBitExec` : `bigint` | `BitExec` → `bigint` | `@iBitExec=BitExec` |
| IC-009098 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailPopulatePC:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulatePC.sql#L52) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `BitExec` : `BIT` | `@iBitExec` : `bigint` | `BitExec` → `bigint` | `@iBitExec=BitExec` |
| IC-011562 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptSummaryPerBranch:103](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSummaryPerBranch.sql#L103) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `BitExec` : `BIT` | `@iBitExec` : `bigint` | `BitExec` → `bigint` | `@iBitExec=BitExec` |
| IC-005983 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateReportPermission:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateReportPermission.sql#L28) | WHERE / P11 / COMPARISON | `BitExec` : `BIT` | `1` : `int` | `BitExec` → `INT` | `BitExec=1` |
| IC-005985 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateReportPermission:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateReportPermission.sql#L46) | WHERE / P11 / COMPARISON | `BitExec` : `BIT` | `1` : `int` | `BitExec` → `INT` | `BitExec=1` |
| IC-005988 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateReportPermission:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateReportPermission.sql#L65) | WHERE / P11 / COMPARISON | `BitExec` : `BIT` | `1` : `int` | `BitExec` → `INT` | `BitExec=1` |
| IC-005991 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateReportPermission:83](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateReportPermission.sql#L83) | WHERE / P11 / COMPARISON | `BitExec` : `BIT` | `1` : `int` | `BitExec` → `INT` | `BitExec=1` |
| IC-009034 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailPopulateCA:84](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulateCA.sql#L84) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `BitExec` : `BIT` | `1` : `int` | `1` → `BIT` | `BitExec=1` |
| IC-009039 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailPopulateCI:90](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulateCI.sql#L90) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `BitExec` : `BIT` | `1` : `int` | `1` → `BIT` | `BitExec=1` |
| IC-009099 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailPopulatePC:73](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulatePC.sql#L73) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `BitExec` : `BIT` | `1` : `int` | `1` → `BIT` | `BitExec=1` |
| IC-011578 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRptSummaryPerBranch:134](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSummaryPerBranch.sql#L134) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `BitExec` : `BIT` | `1` : `int` | `1` → `BIT` | `BitExec=1` |
| IC-012462 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateBitReportPermission:32](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateBitReportPermission.sql#L32) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `BitExec` : `BIT` | `0` : `int` | `0` → `BIT` | `BitExec=0` |
