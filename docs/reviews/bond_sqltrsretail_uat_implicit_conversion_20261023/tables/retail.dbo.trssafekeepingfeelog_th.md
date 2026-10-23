# Retail UAT.dbo.TRSSafeKeepingFeeLog_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 25 lokasi source, 13 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-refundamount"></a>
## RefundAmount

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSafeKeepingFeeLog_TH.sql#L7). **Existing leading index:** TRSSafeKeepingFeeLog_TH_NN2.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-011475 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptRefundSafeKeepingFee:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptRefundSafeKeepingFee.sql#L46) | WHERE / P11 / COMPARISON | `tSafeKeepingFee.RefundAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `tSafeKeepingFee.RefundAmount > 0` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSafeKeepingFeeLog_TH.sql#L15). **Existing leading index:** TRSSafeKeepingFeeLog_TH_NN1.

**Frekuensi source:** 23 lokasi / 13 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003223 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSAmendSecurityTransaction_TT:919](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L919) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'P'` : `varchar` | `Status` → `VARCHAR` | `Status = 'P'` |
| IC-004079 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSGenerateSafeKeepingFeeNasabah:161](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeNasabah.sql#L161) | WHERE / P12 / COMPARISON | `tRankedTable.Status` : `CHAR(1)` | `'F'` : `varchar` | `tRankedTable.Status` → `VARCHAR` | `tRankedTable.Status = 'F'` |
| IC-004337 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSInsertSecurityTransaction_TT_Original:1429](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L1429) | WHERE / P12 / COMPARISON | `[Status]` : `CHAR(1)` | `'P'` : `varchar` | `[Status]` → `VARCHAR` | `[Status] = 'P'` |
| IC-004339 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSInsertSecurityTransaction_TT_Original:1457](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L1457) | WHERE / P12 / COMPARISON | `[Status]` : `CHAR(1)` | `'F'` : `varchar` | `[Status]` → `VARCHAR` | `[Status] = 'F'` |
| IC-006819 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessDeleteSecurityTransactionCommit:139](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessDeleteSecurityTransactionCommit.sql#L139) | JOIN_ON / P12 / COMPARISON | `tskflth.Status` : `CHAR(1)` | `'P'` : `varchar` | `tskflth.Status` → `VARCHAR` | `tskflth.Status = 'P'` |
| IC-006913 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessInsertSecurityTransactionCommit:288](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertSecurityTransactionCommit.sql#L288) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'F'` : `varchar` | `Status` → `VARCHAR` | `Status = 'F'` |
| IC-008580 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSResponseBondFromMX:444](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L444) | WHERE / P12 / COMPARISON | `tskflth.[Status]` : `CHAR(1)` | `'P'` : `varchar` | `tskflth.[Status]` → `VARCHAR` | `tskflth.[Status] = 'P'` |
| IC-008704 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailCalculateCoupon:120](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateCoupon.sql#L120) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'P'` : `varchar` | `Status` → `VARCHAR` | `Status = 'P'` |
| IC-008706 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailCalculateCoupon:135](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateCoupon.sql#L135) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'F'` : `varchar` | `Status` → `VARCHAR` | `Status = 'F'` |
| IC-011470 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptRefundSafeKeepingFee:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptRefundSafeKeepingFee.sql#L23) | CASE / P12 / COMPARISON | `tSafeKeepingFee.Status` : `CHAR(1)` | `'P'` : `varchar` | `tSafeKeepingFee.Status` → `VARCHAR` | `tSafeKeepingFee.Status = 'P'` |
| IC-011471 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptRefundSafeKeepingFee:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptRefundSafeKeepingFee.sql#L24) | CASE / P12 / COMPARISON | `tSafeKeepingFee.Status` : `CHAR(1)` | `'S'` : `varchar` | `tSafeKeepingFee.Status` → `VARCHAR` | `tSafeKeepingFee.Status = 'S'` |
| IC-011472 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptRefundSafeKeepingFee:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptRefundSafeKeepingFee.sql#L25) | CASE / P12 / COMPARISON | `tSafeKeepingFee.Status` : `CHAR(1)` | `'F'` : `varchar` | `tSafeKeepingFee.Status` → `VARCHAR` | `tSafeKeepingFee.Status = 'F'` |
| IC-011500 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSafeKeepingFeeNasabah:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSafeKeepingFeeNasabah.sql#L13) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `skflth.Status` : `CHAR(1)` | `'P'` : `varchar` | `skflth.Status` → `VARCHAR` | `WHEN 'P' THEN 'Menunggu Response Core Bank'` |
| IC-011501 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSafeKeepingFeeNasabah:14](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSafeKeepingFeeNasabah.sql#L14) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `skflth.Status` : `CHAR(1)` | `'F'` : `varchar` | `skflth.Status` → `VARCHAR` | `WHEN 'F' THEN 'Proses Gagal'` |
| IC-011502 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSafeKeepingFeeNasabah:15](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSafeKeepingFeeNasabah.sql#L15) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `skflth.Status` : `CHAR(1)` | `'S'` : `varchar` | `skflth.Status` → `VARCHAR` | `WHEN 'S' THEN 'Proses Sukses'` |
| IC-012547 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdatePostingResult:124](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L124) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'P'` : `varchar` | `Status` → `VARCHAR` | `Status = 'P'` |
| IC-012980 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_DeleteStatusSecurityTransaction_TT:83](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_DeleteStatusSecurityTransaction_TT.sql#L83) | JOIN_ON / P12 / COMPARISON | `tskflth.Status` : `CHAR(1)` | `'P'` : `varchar` | `tskflth.Status` → `VARCHAR` | `tskflth.Status = 'P'` |
| IC-013813 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_UpdateStatusSecurityTransactionAfterMurex:282](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransactionAfterMurex.sql#L282) | JOIN_ON / P12 / COMPARISON | `skf.Status` : `CHAR(1)` | `'P'` : `varchar` | `skf.Status` → `VARCHAR` | `skf.Status = 'P'` |
| IC-013835 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_UpdateStatusSecurityTransaction_TT:196](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransaction_TT.sql#L196) | JOIN_ON / P12 / COMPARISON | `skf.Status` : `CHAR(1)` | `'P'` : `varchar` | `skf.Status` → `VARCHAR` | `skf.Status = 'P'` |
| IC-004110 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSGenerateSafeKeepingFeeNasabah:390](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeNasabah.sql#L390) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssafekeepingfeelog_th.Status` : `CHAR(1)` | `CASE WHEN AmountFeeType = 1 THEN 'P' WHEN AmountFeeType = 2 THEN 'S' END` : `VARCHAR (combined result; facets unknown)` | `CASE WHEN AmountFeeType = 1 THEN 'P' WHEN AmountFeeType = 2 THEN 'S' END` → `CHAR(1)` | `retail.dbo.trssafekeepingfeelog_th.Status <- CASE WHEN AmountFeeType = 1 THEN 'P' WHEN AmountFeeType = 2 THEN 'S' END` |
| IC-004399 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertSecurityTransaction_TT_Original:2302](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L2302) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssafekeepingfeelog_th.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trssafekeepingfeelog_th.Status <- 'P'` |
| IC-006931 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessInsertSecurityTransactionCommit:528](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertSecurityTransactionCommit.sql#L528) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssafekeepingfeelog_th.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trssafekeepingfeelog_th.Status <- 'P'` |
| IC-013834 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateStatusSecurityTransaction_TT:191](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransaction_TT.sql#L191) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `skf.Status` : `CHAR(1)` | `'S'` : `varchar` | `'S'` → `CHAR(1)` | `skf.Status = 'S'` |

<a id="kolom-uniqueid"></a>
## UniqueId

**Tipe dasar:** `UNIQUEIDENTIFIER`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSafeKeepingFeeLog_TH.sql#L14). **Existing leading index:** TRSSafeKeepingFeeLog_TH_NN4.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-011473 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptRefundSafeKeepingFee:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptRefundSafeKeepingFee.sql#L44) | JOIN_ON / P13 / COMPARISON | `tSafeKeepingFee.UniqueId` : `UNIQUEIDENTIFIER` | `po.OBLREM3` : `VARCHAR(40)` | `po.OBLREM3` → `UNIQUEIDENTIFIER` | `tSafeKeepingFee.UniqueId = po.OBLREM3` |
