# SQL_REPLICATE.dbo.CFAIDN — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 37 lokasi source, 9 procedure, 7 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 13 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cfiepl"></a>
## CFIEPL

**Tipe dasar:** `CHAR(30)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFAIDN.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004767 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSOBLExpiredKitasRPT:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOBLExpiredKitasRPT.sql#L61) | EXPRESSION / P07 / COALESCE_BRANCH | `CA.CFIEPL` : `CHAR(30)` | `result type` : `VARCHAR (combined result; facets unknown)` | `CA.CFIEPL` → `VARCHAR` | `COALESCE(CA.CFIEPL, '')` |

<a id="kolom-cfiex6"></a>
## CFIEX6

**Tipe dasar:** `CHAR(8)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFAIDN.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004432 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSJobSyncKITAS:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSJobSyncKITAS.sql#L37) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cf.[CFIEX6]` : `CHAR(8)` | `tc.[KITASExpDate]` : `DATETIME` | `cf.[CFIEX6]` → `DATETIME` | `tc.[KITASExpDate] = cf.[CFIEX6]` |

<a id="kolom-cfiex7"></a>
## CFIEX7

**Tipe dasar:** `CHAR(7)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFAIDN.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004764 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSOBLExpiredKitasRPT:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOBLExpiredKitasRPT.sql#L51) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CA.CFIEX7` : `CHAR(7)` | `@tmpCif.CFIEX7` : `VARCHAR(7)` | `CA.CFIEX7` → `VARCHAR(7)` | `@tmpCif.CFIEX7 <- CA.CFIEX7` |

<a id="kolom-cfisd6"></a>
## CFISD6

**Tipe dasar:** `CHAR(8)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFAIDN.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN_EXPRESSION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003806 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSDataNasabahRetailIndividuOrchid:274](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L274) | WHERE / P04 / COMPARISON | `ISNULL (cfaidn.CFISD6, 0)` : `CHAR(8)` | `0` : `int` | `ISNULL (cfaidn.CFISD6, 0)` → `INT` | `ISNULL (cfaidn.CFISD6, 0) <> 0` |
| IC-003891 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSDataNasabahRetailInstitusi:360](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L360) | WHERE / P04 / COMPARISON | `ISNULL (cfaidn.CFISD6, 0)` : `CHAR(8)` | `0` : `int` | `ISNULL (cfaidn.CFISD6, 0)` → `INT` | `ISNULL (cfaidn.CFISD6, 0) <> 0` |
| IC-003941 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSDataNasabahRetailInstitusiOrchid:220](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L220) | WHERE / P04 / COMPARISON | `ISNULL (cfaidn.CFISD6, 0)` : `CHAR(8)` | `0` : `int` | `ISNULL (cfaidn.CFISD6, 0)` → `INT` | `ISNULL (cfaidn.CFISD6, 0) <> 0` |
| IC-003807 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSDataNasabahRetailIndividuOrchid:274](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L274) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cfaidn.CFISD6` : `CHAR(8)` | `0` : `int` | `0` → `CHAR(8)` | `ISNULL (cfaidn.CFISD6, 0)` |
| IC-003892 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSDataNasabahRetailInstitusi:360](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L360) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cfaidn.CFISD6` : `CHAR(8)` | `0` : `int` | `0` → `CHAR(8)` | `ISNULL (cfaidn.CFISD6, 0)` |
| IC-003942 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSDataNasabahRetailInstitusiOrchid:220](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L220) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cfaidn.CFISD6` : `CHAR(8)` | `0` : `int` | `0` → `CHAR(8)` | `ISNULL (cfaidn.CFISD6, 0)` |

<a id="kolom-cfisd7"></a>
## CFISD7

**Tipe dasar:** `CHAR(7)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFAIDN.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004765 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSOBLExpiredKitasRPT:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOBLExpiredKitasRPT.sql#L51) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CA.CFISD7` : `CHAR(7)` | `@tmpCif.CFISD7` : `VARCHAR(7)` | `CA.CFISD7` → `VARCHAR(7)` | `@tmpCif.CFISD7 <- CA.CFISD7` |

<a id="kolom-cfsscd"></a>
## CFSSCD

**Tipe dasar:** `CHAR(4)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFAIDN.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 14 lokasi / 9 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003804 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:273](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L273) | JOIN_ON / P12 / COMPARISON | `CFSSCD` : `CHAR(4)` | `'NPWP'` : `varchar` | `CFSSCD` → `VARCHAR` | `CFSSCD = 'NPWP'` |
| IC-003889 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:359](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L359) | JOIN_ON / P12 / COMPARISON | `CFSSCD` : `CHAR(4)` | `'NPWP'` : `varchar` | `CFSSCD` → `VARCHAR` | `CFSSCD = 'NPWP'` |
| IC-003894 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:368](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L368) | JOIN_ON / P12 / COMPARISON | `CFSSCD` : `CHAR(4)` | `'SKD'` : `varchar` | `CFSSCD` → `VARCHAR` | `CFSSCD = 'SKD'` |
| IC-003896 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:375](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L375) | JOIN_ON / P12 / COMPARISON | `CFSSCD` : `CHAR(4)` | `'APP'` : `varchar` | `CFSSCD` → `VARCHAR` | `CFSSCD = 'APP'` |
| IC-003898 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:382](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L382) | JOIN_ON / P12 / COMPARISON | `CFSSCD` : `CHAR(4)` | `'SIUP'` : `varchar` | `CFSSCD` → `VARCHAR` | `CFSSCD = 'SIUP'` |
| IC-003939 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusiOrchid:219](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L219) | JOIN_ON / P12 / COMPARISON | `CFSSCD` : `CHAR(4)` | `'NPWP'` : `varchar` | `CFSSCD` → `VARCHAR` | `CFSSCD = 'NPWP'` |
| IC-003961 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusiOrchid:350](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L350) | JOIN_ON / P12 / COMPARISON | `CFSSCD` : `CHAR(4)` | `'SIUP'` : `varchar` | `CFSSCD` → `VARCHAR` | `CFSSCD = 'SIUP'` |
| IC-004433 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSJobSyncKITAS:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSJobSyncKITAS.sql#L41) | WHERE / P12 / COMPARISON | `cf.[CFSSCD]` : `CHAR(4)` | `'KITS'` : `varchar` | `cf.[CFSSCD]` → `VARCHAR` | `cf.[CFSSCD] = 'KITS'` |
| IC-004436 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSJobSyncKITAS:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSJobSyncKITAS.sql#L52) | JOIN_ON / P12 / COMPARISON | `cfi.[CFSSCD]` : `CHAR(4)` | `'KITS'` : `varchar` | `cfi.[CFSSCD]` → `VARCHAR` | `cfi.[CFSSCD] = 'KITS'` |
| IC-004768 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSOBLExpiredKitasRPT:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOBLExpiredKitasRPT.sql#L66) | WHERE / P12 / COMPARISON | `CA.CFSSCD` : `CHAR(4)` | `'KITS'` : `varchar` | `CA.CFSSCD` → `VARCHAR` | `CA.CFSSCD = 'KITS'` |
| IC-007850 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSREPOReportBiayaMaterai:320](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L320) | WHERE / P12 / COMPARISON | `CFSSCD` : `CHAR(4)` | `'PP'` : `varchar` | `CFSSCD` → `VARCHAR` | `CFSSCD = 'PP'` |
| IC-008148 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportMatchingSID:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportMatchingSID.sql#L28) | JOIN_ON / P12 / COMPARISON | `cf.CFSSCD` : `CHAR(4)` | `'SII'` : `varchar` | `cf.CFSSCD` → `VARCHAR` | `cf.CFSSCD = 'SII'` |
| IC-009183 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:429](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L429) | WHERE / P12 / COMPARISON | `CFSSCD` : `CHAR(4)` | `'NPWP'` : `varchar` | `CFSSCD` → `VARCHAR` | `CFSSCD = 'NPWP'` |
| IC-009350 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:327](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L327) | WHERE / P12 / COMPARISON | `CFSSCD` : `CHAR(4)` | `'NPWP'` : `varchar` | `CFSSCD` → `VARCHAR` | `CFSSCD = 'NPWP'` |

<a id="kolom-cfssno"></a>
## CFSSNO

**Tipe dasar:** `CHAR(40)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFAIDN.sql#L3). **Existing leading index:** idx_cfssno_cfsscd.

**Frekuensi source:** 13 lokasi / 7 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003893 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailInstitusi:364](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L364) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfaidn.CFSSNO` : `CHAR(40)` | `SKD` : `VARCHAR(120)` | `cfaidn.CFSSNO` → `VARCHAR(120)` | `SKD = cfaidn.CFSSNO` |
| IC-003895 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailInstitusi:372](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L372) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfaidn.CFSSNO` : `CHAR(40)` | `ARTICLE_OF_ASSC` : `VARCHAR(120)` | `cfaidn.CFSSNO` → `VARCHAR(120)` | `ARTICLE_OF_ASSC = cfaidn.CFSSNO` |
| IC-003897 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailInstitusi:379](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L379) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfaidn.CFSSNO` : `CHAR(40)` | `BUSINESS_REG_CERT` : `VARCHAR(120)` | `cfaidn.CFSSNO` → `VARCHAR(120)` | `BUSINESS_REG_CERT = cfaidn.CFSSNO` |
| IC-003960 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailInstitusiOrchid:347](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L347) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfaidn.CFSSNO` : `CHAR(40)` | `BusinessRegistrationNumber` : `VARCHAR(90)` | `cfaidn.CFSSNO` → `VARCHAR(90)` | `BusinessRegistrationNumber = cfaidn.CFSSNO` |
| IC-004431 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSJobSyncKITAS:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSJobSyncKITAS.sql#L36) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cf.[CFSSNO]` : `CHAR(40)` | `tc.[KITASNo]` : `VARCHAR(40)` | `cf.[CFSSNO]` → `VARCHAR(40)` | `tc.[KITASNo] = cf.[CFSSNO]` |
| IC-004766 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSOBLExpiredKitasRPT:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOBLExpiredKitasRPT.sql#L51) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CA.CFSSNO` : `CHAR(40)` | `@tmpCif.CFSSNO` : `VARCHAR(40)` | `CA.CFSSNO` → `VARCHAR(40)` | `@tmpCif.CFSSNO <- CA.CFSSNO` |
| IC-007849 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSREPOReportBiayaMaterai:315](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L315) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `CFSSNO` : `CHAR(40)` | `NomorIdentitas` : `VARCHAR(50)` | `CFSSNO` → `VARCHAR(50)` | `NomorIdentitas = CFSSNO` |
| IC-003805 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:273](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L273) | JOIN_ON / P12 / COMPARISON | `cfaidn.CFSSNO` : `CHAR(40)` | `mnc.NPWP` : `VARCHAR(90)` | `cfaidn.CFSSNO` → `VARCHAR` | `mnc.NPWP = cfaidn.CFSSNO` |
| IC-003890 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:359](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L359) | JOIN_ON / P12 / COMPARISON | `cfaidn.CFSSNO` : `CHAR(40)` | `mnc.NPWP` : `VARCHAR(90)` | `cfaidn.CFSSNO` → `VARCHAR` | `mnc.NPWP = cfaidn.CFSSNO` |
| IC-003940 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusiOrchid:219](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L219) | JOIN_ON / P12 / COMPARISON | `cfaidn.CFSSNO` : `CHAR(40)` | `mnc.NPWP` : `VARCHAR(90)` | `cfaidn.CFSSNO` → `VARCHAR` | `mnc.NPWP = cfaidn.CFSSNO` |
| IC-004769 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSOBLExpiredKitasRPT:67](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOBLExpiredKitasRPT.sql#L67) | EXPRESSION / P07 / COALESCE_BRANCH | `CA.CFSSNO` : `CHAR(40)` | `result type` : `VARCHAR (combined result; facets unknown)` | `CA.CFSSNO` → `VARCHAR` | `COALESCE(CA.CFSSNO,'')` |
| IC-008149 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportMatchingSID:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportMatchingSID.sql#L30) | WHERE / P12 / COMPARISON | `cf.CFSSNO` : `CHAR(40)` | `isnull(mst.SID,'')` : `VARCHAR(100)` | `cf.CFSSNO` → `VARCHAR` | `cf.CFSSNO != isnull(mst.SID,'')` |
| IC-004430 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSJobSyncKITAS:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSJobSyncKITAS.sql#L31) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cf.[CFSSNO]` : `CHAR(40)` | `''` : `varchar` | `''` → `CHAR(40)` | `isnull(cf.[CFSSNO], '')` |
