# Retail UAT.dbo.REPOApprovalParameterDetail_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 18 lokasi source, 3 procedure, 6 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-actioncode"></a>
## ActionCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOApprovalParameterDetail_TR.sql#L28). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 11 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003310 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamRepo:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamRepo.sql#L165) | WHERE / P12 / COMPARISON | `tr.ActionCode` : `CHAR(1)` | `'I'` : `varchar` | `tr.ActionCode` → `VARCHAR` | `tr.ActionCode = 'I'` |
| IC-003312 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamRepo:235](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamRepo.sql#L235) | WHERE / P12 / COMPARISON | `tr.ActionCode` : `CHAR(1)` | `'D'` : `varchar` | `tr.ActionCode` → `VARCHAR` | `tr.ActionCode = 'D'` |
| IC-003314 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamRepo:269](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamRepo.sql#L269) | WHERE / P12 / COMPARISON | `tr.ActionCode` : `CHAR(1)` | `'N'` : `varchar` | `tr.ActionCode` → `VARCHAR` | `tr.ActionCode = 'N'` |
| IC-003315 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamRepo:329](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamRepo.sql#L329) | CASE / P12 / COMPARISON | `tr.ActionCode` : `CHAR(1)` | `'N'` : `varchar` | `tr.ActionCode` → `VARCHAR` | `tr.ActionCode = 'N'` |
| IC-003316 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamRepo:330](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamRepo.sql#L330) | CASE / P12 / COMPARISON | `tr.ActionCode` : `CHAR(1)` | `'D'` : `varchar` | `tr.ActionCode` → `VARCHAR` | `tr.ActionCode = 'D'` |
| IC-003318 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamRepo:338](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamRepo.sql#L338) | WHERE / P12 / COMPARISON | `tr.ActionCode` : `CHAR(1)` | `'D'` : `varchar` | `tr.ActionCode` → `VARCHAR` | `tr.ActionCode = 'D'` |
| IC-003319 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamRepo:338](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamRepo.sql#L338) | WHERE / P12 / COMPARISON | `tr.ActionCode` : `CHAR(1)` | `'N'` : `varchar` | `tr.ActionCode` → `VARCHAR` | `tr.ActionCode = 'N'` |
| IC-005932 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateParamRepo:178](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamRepo.sql#L178) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ActionCode` : `CHAR(1)` | `'I'` : `varchar` | `ActionCode` → `VARCHAR` | `WHEN 'I' THEN 'Insert'` |
| IC-005933 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateParamRepo:179](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamRepo.sql#L179) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ActionCode` : `CHAR(1)` | `'D'` : `varchar` | `ActionCode` → `VARCHAR` | `WHEN 'D' THEN 'Delete'` |
| IC-005934 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateParamRepo:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamRepo.sql#L180) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ActionCode` : `CHAR(1)` | `'O'` : `varchar` | `ActionCode` → `VARCHAR` | `WHEN 'O' THEN 'Update - Old Data'` |
| IC-005935 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateParamRepo:181](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamRepo.sql#L181) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ActionCode` : `CHAR(1)` | `'N'` : `varchar` | `ActionCode` → `VARCHAR` | `WHEN 'N' THEN 'Update - New Data'` |

<a id="kolom-batasbawahspread"></a>
## BatasBawahSpread

**Tipe dasar:** `DECIMAL(15,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOApprovalParameterDetail_TR.sql#L29). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008086 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRepoSaveData:642](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L642) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repoapprovalparameterdetail_tr.BatasBawahSpread` : `DECIMAL(15,2)` | `BatasBawahSpread` : `MONEY` | `BatasBawahSpread` → `DECIMAL(15,2)` | `retail.dbo.repoapprovalparameterdetail_tr.BatasBawahSpread <- BatasBawahSpread` |

<a id="kolom-haircut"></a>
## Haircut

**Tipe dasar:** `DECIMAL(15,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOApprovalParameterDetail_TR.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008087 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRepoSaveData:642](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L642) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repoapprovalparameterdetail_tr.Haircut` : `DECIMAL(15,2)` | `Haircut` : `MONEY` | `Haircut` → `DECIMAL(15,2)` | `retail.dbo.repoapprovalparameterdetail_tr.Haircut <- Haircut` |

<a id="kolom-maxspread"></a>
## MaxSpread

**Tipe dasar:** `DECIMAL(15,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOApprovalParameterDetail_TR.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008088 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRepoSaveData:642](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L642) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repoapprovalparameterdetail_tr.MaxSpread` : `DECIMAL(15,2)` | `MaxSpread` : `MONEY` | `MaxSpread` → `DECIMAL(15,2)` | `retail.dbo.repoapprovalparameterdetail_tr.MaxSpread <- MaxSpread` |

<a id="kolom-minspread"></a>
## MinSpread

**Tipe dasar:** `DECIMAL(15,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOApprovalParameterDetail_TR.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008089 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRepoSaveData:642](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L642) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repoapprovalparameterdetail_tr.MinSpread` : `DECIMAL(15,2)` | `MinSpread` : `MONEY` | `MinSpread` → `DECIMAL(15,2)` | `retail.dbo.repoapprovalparameterdetail_tr.MinSpread <- MinSpread` |

<a id="kolom-repoprodid"></a>
## RepoProdId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOApprovalParameterDetail_TR.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003311 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSApprovalParamRepo:232](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamRepo.sql#L232) | JOIN_ON / P11 / COMPARISON | `tr.RepoProdId` : `BIGINT` | `tm.RepoProdId` : `INT` | `tm.RepoProdId` → `BIGINT` | `tm.RepoProdId = tr.RepoProdId` |
| IC-003313 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSApprovalParamRepo:266](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamRepo.sql#L266) | JOIN_ON / P11 / COMPARISON | `tr.RepoProdId` : `BIGINT` | `tm.RepoProdId` : `INT` | `tm.RepoProdId` → `BIGINT` | `tm.RepoProdId = tr.RepoProdId` |
| IC-003317 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSApprovalParamRepo:335](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamRepo.sql#L335) | JOIN_ON / P11 / COMPARISON | `tr.RepoProdId` : `BIGINT` | `tm.RepoProdId` : `INT` | `tm.RepoProdId` → `BIGINT` | `tm.RepoProdId = tr.RepoProdId` |
