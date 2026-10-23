# Retail UAT.dbo.TRSMaturityTaxAmnesty_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 17 lokasi source, 6 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-inserteddate"></a>
## InsertedDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMaturityTaxAmnesty_TT.sql#L21). **Existing leading index:** TRSMaturityTaxAmnesty_TT_NN2.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007021 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessMaturityUploadAuthorization:122](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessMaturityUploadAuthorization.sql#L122) | WHERE / P13 / COMPARISON | `mtt.InsertedDate` : `DATETIME` | `@dProcessDate` : `DATE` | `@dProcessDate` → `DATETIME` | `mtt.InsertedDate >= @dProcessDate` |
| IC-007022 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessMaturityUploadAuthorization:123](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessMaturityUploadAuthorization.sql#L123) | WHERE / P13 / COMPARISON | `mtt.InsertedDate` : `DATETIME` | `DATEADD(DAY, 1, @dCurrentDate)` : `DATE` | `DATEADD(DAY, 1, @dCurrentDate)` → `DATETIME` | `mtt.InsertedDate < DATEADD(DAY, 1, @dCurrentDate)` |

<a id="kolom-istaxamnesty"></a>
## IsTaxAmnesty

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMaturityTaxAmnesty_TT.sql#L25). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007501 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSettlementMaturity:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementMaturity.sql#L111) | CASE / P11 / COMPARISON | `mtt.IsTaxAmnesty` : `BIT` | `0` : `int` | `mtt.IsTaxAmnesty` → `INT` | `mtt.IsTaxAmnesty = 0` |
| IC-002276 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSMaturityProcess:707](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityProcess.sql#L707) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmaturitytaxamnesty_tt.IsTaxAmnesty` : `BIT` | `CASE WHEN tmpSec.SecAccNo IS NULL THEN 0 ELSE 1 END` : `INT (combined result; facets unknown)` | `CASE WHEN tmpSec.SecAccNo IS NULL THEN 0 ELSE 1 END` → `BIT` | `retail.dbo.trsmaturitytaxamnesty_tt.IsTaxAmnesty <- CASE WHEN tmpSec.SecAccNo IS NULL THEN 0 ELSE 1 END` |
| IC-002277 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSMaturityProcess:768](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityProcess.sql#L768) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmaturitytaxamnesty_tt.IsTaxAmnesty` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.trsmaturitytaxamnesty_tt.IsTaxAmnesty <- 0` |
| IC-002278 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSMaturityProcess:823](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityProcess.sql#L823) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmaturitytaxamnesty_tt.IsTaxAmnesty` : `BIT` | `1` : `int` | `1` → `BIT` | `retail.dbo.trsmaturitytaxamnesty_tt.IsTaxAmnesty <- 1` |
| IC-002279 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSMaturityProcess:877](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityProcess.sql#L877) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmaturitytaxamnesty_tt.IsTaxAmnesty` : `BIT` | `CASE WHEN tmpSec.SecAccNo IS NULL THEN 0 ELSE 1 END` : `INT (combined result; facets unknown)` | `CASE WHEN tmpSec.SecAccNo IS NULL THEN 0 ELSE 1 END` → `BIT` | `retail.dbo.trsmaturitytaxamnesty_tt.IsTaxAmnesty <- CASE WHEN tmpSec.SecAccNo IS NULL THEN 0 ELSE 1 END` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMaturityTaxAmnesty_TT.sql#L19). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 10 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002266 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSMaturityProcess:621](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityProcess.sql#L621) | WHERE / P12 / IN_VALUE_TYPE | `mttta.Status` : `CHAR(1)` | `'P'` : `varchar` | `mttta.Status` → `VARCHAR` | `mttta.Status in ('P','R')` |
| IC-005625 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateMaturityUpload:617](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L617) | CASE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'R'` : `varchar` | `Status` → `VARCHAR` | `Status = 'R'` |
| IC-005626 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateMaturityUpload:618](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L618) | CASE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'A'` : `varchar` | `Status` → `VARCHAR` | `Status = 'A'` |
| IC-005627 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateMaturityUpload:619](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L619) | CASE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'P'` : `varchar` | `Status` → `VARCHAR` | `Status = 'P'` |
| IC-005737 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateMaturityUploadAuthorization:38](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUploadAuthorization.sql#L38) | WHERE / P12 / COMPARISON | `mtt.Status` : `CHAR(1)` | `'P'` : `varchar` | `mtt.Status` → `VARCHAR` | `mtt.Status = 'P'` |
| IC-006305 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementMaturity:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementMaturity.sql#L141) | WHERE / P12 / COMPARISON | `mtt.[Status]` : `CHAR(1)` | `'A'` : `varchar` | `mtt.[Status]` → `VARCHAR` | `mtt.[Status] = 'A'` |
| IC-007023 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessMaturityUploadAuthorization:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessMaturityUploadAuthorization.sql#L125) | WHERE / P12 / COMPARISON | `mtt.Status` : `CHAR(1)` | `'P'` : `varchar` | `mtt.Status` → `VARCHAR` | `mtt.Status = 'P'` |
| IC-007024 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessMaturityUploadAuthorization:163](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessMaturityUploadAuthorization.sql#L163) | WHERE / P12 / COMPARISON | `mtt.Status` : `CHAR(1)` | `'P'` : `varchar` | `mtt.Status` → `VARCHAR` | `mtt.Status <> 'P'` |
| IC-007026 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessMaturityUploadAuthorization:185](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessMaturityUploadAuthorization.sql#L185) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mtt.Status` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `mtt.Status = 'A'` |
| IC-007031 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessMaturityUploadAuthorization:306](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessMaturityUploadAuthorization.sql#L306) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mtt.Status` : `CHAR(1)` | `'R'` : `varchar` | `'R'` → `CHAR(1)` | `mtt.Status = 'R'` |
