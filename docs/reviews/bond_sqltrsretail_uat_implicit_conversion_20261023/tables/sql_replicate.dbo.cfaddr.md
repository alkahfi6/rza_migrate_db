# SQL_REPLICATE.dbo.CFADDR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 184 lokasi source, 16 procedure, 17 kolom. **A:** 3 lokasi pada 3 procedure. **B:** 63 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cfaseq"></a>
## CFASEQ

**Tipe dasar:** `CHAR(3)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFADDR.sql#L4). **Existing leading index:** idx_cfaddr_cfaseq.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 3 lokasi / 3 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008147 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSReportFormBuySellObli:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportFormBuySellObli.sql#L111) | JOIN_ON / P01 / COMPARISON | `addr.CFASEQ` : `CHAR(3)` | `tctm.AlamatSuratSequence` : `INT` | `addr.CFASEQ` → `INT` | `tctm.AlamatSuratSequence = addr.CFASEQ` |
| IC-010324 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptObligasiKonfirmasiBond:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptObligasiKonfirmasiBond.sql#L46) | JOIN_ON / P01 / COMPARISON | `addr.CFASEQ` : `CHAR(3)` | `tctm.AlamatSuratSequence` : `INT` | `addr.CFASEQ` → `INT` | `tctm.AlamatSuratSequence = addr.CFASEQ` |
| IC-010608 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptPTOSNasabah:187](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPTOSNasabah.sql#L187) | JOIN_ON / P01 / COMPARISON | `addr.CFASEQ` : `CHAR(3)` | `tc.AlamatSuratSequence` : `INT` | `addr.CFASEQ` → `INT` | `tc.AlamatSuratSequence = addr.CFASEQ` |

<a id="kolom-cf2dt2"></a>
## CF2DT2

**Tipe dasar:** `CHAR(40)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFADDR.sql#L31). **Existing leading index:** idx_cfaddr_cf2dt2.

**Frekuensi source:** 17 lokasi / 8 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000832 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L108) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CF2DT2` : `CHAR(40)` | `@cCF2DT2` : `varchar(40)` | `CF2DT2` → `varchar(40)` | `@cCF2DT2 = CF2DT2` |
| IC-000857 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:154](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L154) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CF2DT2` : `CHAR(40)` | `@cCF2DT2` : `varchar(40)` | `CF2DT2` → `varchar(40)` | `@cCF2DT2 = CF2DT2` |
| IC-003708 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailIndividu:306](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L306) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `cfaddr.CF2DT2` : `CHAR(40)` | `#xtmpAddress.CF2DT2` : `VARCHAR(40)` | `cfaddr.CF2DT2` → `VARCHAR(40)` | `#xtmpAddress.CF2DT2 <- cfaddr.CF2DT2` |
| IC-003810 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailIndividuOrchid:284](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L284) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `cfaddr.CF2DT2` : `CHAR(40)` | `#xtmpAddress.CF2DT2` : `VARCHAR(40)` | `cfaddr.CF2DT2` → `VARCHAR(40)` | `#xtmpAddress.CF2DT2 <- cfaddr.CF2DT2` |
| IC-003861 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailInstitusi:249](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L249) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `cfaddr.CF2DT2` : `CHAR(40)` | `#xtmpAddress.CF2DT2` : `VARCHAR(40)` | `cfaddr.CF2DT2` → `VARCHAR(40)` | `#xtmpAddress.CF2DT2 <- cfaddr.CF2DT2` |
| IC-003945 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailInstitusiOrchid:230](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L230) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `cfaddr.CF2DT2` : `CHAR(40)` | `#xtmpAddress.CF2DT2` : `VARCHAR(40)` | `cfaddr.CF2DT2` → `VARCHAR(40)` | `#xtmpAddress.CF2DT2 <- cfaddr.CF2DT2` |
| IC-009192 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailSyncDataCore:464](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L464) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ca.CF2DT2` : `CHAR(40)` | `#xtmpAddress.CF2DT2` : `VARCHAR(40)` | `ca.CF2DT2` → `VARCHAR(40)` | `#xtmpAddress.CF2DT2 <- ca.CF2DT2` |
| IC-009361 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailSyncDataCoreWorkingHours:361](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L361) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ca.CF2DT2` : `CHAR(40)` | `#xtmpAddress.CF2DT2` : `VARCHAR(40)` | `ca.CF2DT2` → `VARCHAR(40)` | `#xtmpAddress.CF2DT2 <- ca.CF2DT2` |
| IC-012493 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSUpdateDataMasterNasabah:248](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateDataMasterNasabah.sql#L248) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `addr.CF2DT2` : `CHAR(40)` | `tmp.NamaDatiII` : `varchar(40)` | `addr.CF2DT2` → `varchar(40)` | `tmp.NamaDatiII = addr.CF2DT2` |
| IC-012497 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSUpdateDataMasterNasabah:257](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateDataMasterNasabah.sql#L257) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `addr.CF2DT2` : `CHAR(40)` | `tmp.NamaDatiII` : `varchar(40)` | `addr.CF2DT2` → `varchar(40)` | `tmp.NamaDatiII = addr.CF2DT2` |
| IC-012501 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSUpdateDataMasterNasabah:266](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateDataMasterNasabah.sql#L266) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `addr.CF2DT2` : `CHAR(40)` | `tmp.NamaDatiII` : `varchar(40)` | `addr.CF2DT2` → `varchar(40)` | `tmp.NamaDatiII = addr.CF2DT2` |
| IC-003717 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:352](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L352) | JOIN_ON / P12 / COMPARISON | `cfaddr.CF2DT2` : `CHAR(40)` | `kk.Nama` : `VARCHAR(50)` | `cfaddr.CF2DT2` → `VARCHAR` | `cfaddr.CF2DT2 = kk.Nama` |
| IC-003819 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:330](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L330) | JOIN_ON / P12 / COMPARISON | `cfaddr.CF2DT2` : `CHAR(40)` | `kk.Nama` : `VARCHAR(50)` | `cfaddr.CF2DT2` → `VARCHAR` | `cfaddr.CF2DT2 = kk.Nama` |
| IC-003870 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:295](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L295) | JOIN_ON / P12 / COMPARISON | `cfaddr.CF2DT2` : `CHAR(40)` | `kk.Nama` : `VARCHAR(50)` | `cfaddr.CF2DT2` → `VARCHAR` | `cfaddr.CF2DT2 = kk.Nama` |
| IC-003954 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusiOrchid:276](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L276) | JOIN_ON / P12 / COMPARISON | `cfaddr.CF2DT2` : `CHAR(40)` | `kk.Nama` : `VARCHAR(50)` | `cfaddr.CF2DT2` → `VARCHAR` | `cfaddr.CF2DT2 = kk.Nama` |
| IC-009215 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:554](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L554) | JOIN_ON / P12 / COMPARISON | `ca.CF2DT2` : `CHAR(40)` | `kk.Nama` : `VARCHAR(50)` | `ca.CF2DT2` → `VARCHAR` | `ca.CF2DT2 = kk.Nama` |
| IC-009384 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:451](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L451) | JOIN_ON / P12 / COMPARISON | `ca.CF2DT2` : `CHAR(40)` | `kk.Nama` : `VARCHAR(50)` | `ca.CF2DT2` → `VARCHAR` | `ca.CF2DT2 = kk.Nama` |

<a id="kolom-cf2kec"></a>
## CF2KEC

**Tipe dasar:** `CHAR(40)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFADDR.sql#L30). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000844 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:122](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L122) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CF2KEC` : `CHAR(40)` | `@cCF2KEC` : `varchar(40)` | `CF2KEC` → `varchar(40)` | `@cCF2KEC = CF2KEC` |
| IC-000869 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L168) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CF2KEC` : `CHAR(40)` | `@cCF2KEC` : `varchar(40)` | `CF2KEC` → `varchar(40)` | `@cCF2KEC = CF2KEC` |
| IC-000874 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:201](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L201) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CF2KEC` : `CHAR(40)` | `@tmp_address.CF2KEC` : `varchar(40)` | `CF2KEC` → `varchar(40)` | `@tmp_address.CF2KEC <- CF2KEC` |

<a id="kolom-cf2kel"></a>
## CF2KEL

**Tipe dasar:** `CHAR(40)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFADDR.sql#L29). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000843 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:121](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L121) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CF2KEL` : `CHAR(40)` | `@cCF2KEL` : `varchar(40)` | `CF2KEL` → `varchar(40)` | `@cCF2KEL = CF2KEL` |
| IC-000868 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:167](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L167) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CF2KEL` : `CHAR(40)` | `@cCF2KEL` : `varchar(40)` | `CF2KEL` → `varchar(40)` | `@cCF2KEL = CF2KEL` |
| IC-000875 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:201](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L201) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CF2KEL` : `CHAR(40)` | `@tmp_address.CF2KEL` : `varchar(40)` | `CF2KEL` → `varchar(40)` | `@tmp_address.CF2KEL <- CF2KEL` |

<a id="kolom-cf2prv"></a>
## CF2PRV

**Tipe dasar:** `CHAR(40)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFADDR.sql#L32). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000845 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:123](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L123) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CF2PRV` : `CHAR(40)` | `@cCF2PRV` : `varchar(40)` | `CF2PRV` → `varchar(40)` | `@cCF2PRV = CF2PRV` |
| IC-000870 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L169) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CF2PRV` : `CHAR(40)` | `@cCF2PRV` : `varchar(40)` | `CF2PRV` → `varchar(40)` | `@cCF2PRV = CF2PRV` |
| IC-000876 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:201](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L201) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CF2PRV` : `CHAR(40)` | `@tmp_address.CF2PRV` : `varchar(40)` | `CF2PRV` → `varchar(40)` | `@tmp_address.CF2PRV <- CF2PRV` |

<a id="kolom-cfapty"></a>
## CFAPTY

**Tipe dasar:** `CHAR(5)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFADDR.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000841 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:119](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L119) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFAPTY` : `CHAR(5)` | `@cCFAPTY` : `varchar(5)` | `CFAPTY` → `varchar(5)` | `@cCFAPTY = CFAPTY` |
| IC-000866 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L165) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFAPTY` : `CHAR(5)` | `@cCFAPTY` : `varchar(5)` | `CFAPTY` → `varchar(5)` | `@cCFAPTY = CFAPTY` |
| IC-000877 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:201](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L201) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFAPTY` : `CHAR(5)` | `@tmp_address.CFAPTY` : `varchar(5)` | `CFAPTY` → `varchar(5)` | `@tmp_address.CFAPTY <- CFAPTY` |

<a id="kolom-cfcif"></a>
## CFCIF

**Tipe dasar:** `CHAR(19)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFADDR.sql#L2). **Existing leading index:** idx_cfaddr_cfcif.

**Frekuensi source:** 29 lokasi / 15 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000827 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMValidateCustomer:101](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L101) | WHERE / P12 / COMPARISON | `CFCIF` : `CHAR(19)` | `@pcCIFNo` : `varchar(19)` | `CFCIF` → `VARCHAR` | `CFCIF = @pcCIFNo` |
| IC-000846 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMValidateCustomer:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L125) | WHERE / P12 / COMPARISON | `CFCIF` : `CHAR(19)` | `@pcCIFNo` : `varchar(19)` | `CFCIF` → `VARCHAR` | `CFCIF = @pcCIFNo` |
| IC-000848 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMValidateCustomer:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L131) | WHERE / P12 / COMPARISON | `CFCIF` : `CHAR(19)` | `@pcCIFNo` : `varchar(19)` | `CFCIF` → `VARCHAR` | `CFCIF = @pcCIFNo` |
| IC-000850 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMValidateCustomer:136](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L136) | WHERE / P12 / COMPARISON | `CFCIF` : `CHAR(19)` | `@pcCIFNo` : `varchar(19)` | `CFCIF` → `VARCHAR` | `CFCIF = @pcCIFNo` |
| IC-000852 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMValidateCustomer:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L141) | WHERE / P12 / COMPARISON | `CFCIF` : `CHAR(19)` | `@pcCIFNo` : `varchar(19)` | `CFCIF` → `VARCHAR` | `CFCIF = @pcCIFNo` |
| IC-000871 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMValidateCustomer:171](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L171) | WHERE / P12 / COMPARISON | `CFCIF` : `CHAR(19)` | `@pcCIFNo` : `varchar(19)` | `CFCIF` → `VARCHAR` | `CFCIF = @pcCIFNo` |
| IC-000887 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMValidateCustomer:206](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L206) | WHERE / P12 / COMPARISON | `CFCIF` : `CHAR(19)` | `@pcCIFNo` : `varchar(19)` | `CFCIF` → `VARCHAR` | `CFCIF = @pcCIFNo` |
| IC-003716 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:350](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L350) | JOIN_ON / P12 / COMPARISON | `cfaddr.CFCIF` : `CHAR(19)` | `mnc._CIFNo` : `VARCHAR(19)` | `cfaddr.CFCIF` → `VARCHAR` | `mnc._CIFNo = cfaddr.CFCIF` |
| IC-003818 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:328](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L328) | JOIN_ON / P12 / COMPARISON | `cfaddr.CFCIF` : `CHAR(19)` | `mnc._CIFNo` : `VARCHAR(19)` | `cfaddr.CFCIF` → `VARCHAR` | `mnc._CIFNo = cfaddr.CFCIF` |
| IC-003869 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:293](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L293) | JOIN_ON / P12 / COMPARISON | `cfaddr.CFCIF` : `CHAR(19)` | `mnc._CIFNo` : `VARCHAR(19)` | `cfaddr.CFCIF` → `VARCHAR` | `mnc._CIFNo = cfaddr.CFCIF` |
| IC-003953 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusiOrchid:274](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L274) | JOIN_ON / P12 / COMPARISON | `cfaddr.CFCIF` : `CHAR(19)` | `mnc._CIFNo` : `VARCHAR(19)` | `cfaddr.CFCIF` → `VARCHAR` | `mnc._CIFNo = cfaddr.CFCIF` |
| IC-005344 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDataAlamatKonfirmasi:12](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDataAlamatKonfirmasi.sql#L12) | WHERE / P12 / COMPARISON | `CFCIF` : `CHAR(19)` | `@pcCIF` : `varchar(19)` | `CFCIF` → `VARCHAR` | `CFCIF = @pcCIF` |
| IC-006413 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateStatusEmailKonfirmasiProduct:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateStatusEmailKonfirmasiProduct.sql#L162) | JOIN_ON / P12 / COMPARISON | `cfa.CFCIF` : `CHAR(19)` | `right(replicate('0', 19) + tmp.CIF, 19)` : `VARCHAR (string function; length unknown)` | `cfa.CFCIF` → `VARCHAR` | `right(replicate('0', 19) + tmp.CIF, 19) = cfa.CFCIF` |
| IC-008146 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportFormBuySellObli:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportFormBuySellObli.sql#L110) | JOIN_ON / P12 / COMPARISON | `addr.CFCIF` : `CHAR(19)` | `tctm.CIFNo` : `VARCHAR(19)` | `addr.CFCIF` → `VARCHAR` | `tctm.CIFNo = addr.CFCIF` |
| IC-009214 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:552](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L552) | JOIN_ON / P12 / COMPARISON | `ca.CFCIF` : `CHAR(19)` | `t.CIFNo` : `VARCHAR(19)` | `ca.CFCIF` → `VARCHAR` | `t.CIFNo = ca.CFCIF` |
| IC-009383 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:449](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L449) | JOIN_ON / P12 / COMPARISON | `ca.CFCIF` : `CHAR(19)` | `t.CIFNo` : `VARCHAR(19)` | `ca.CFCIF` → `VARCHAR` | `t.CIFNo = ca.CFCIF` |
| IC-009592 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptCustomerPortfolio:90](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerPortfolio.sql#L90) | JOIN_ON / P12 / COMPARISON | `CFADDR.CFCIF` : `CHAR(19)` | `tctm.CIFNo` : `VARCHAR(19)` | `CFADDR.CFCIF` → `VARCHAR` | `CFADDR.CFCIF = tctm.CIFNo` |
| IC-010607 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPTOSNasabah:186](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPTOSNasabah.sql#L186) | JOIN_ON / P12 / COMPARISON | `addr.CFCIF` : `CHAR(19)` | `tc.CIFNo` : `VARCHAR(19)` | `addr.CFCIF` → `VARCHAR` | `tc.CIFNo = addr.CFCIF` |
| IC-012377 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncAlamatKorespodensi:20](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncAlamatKorespodensi.sql#L20) | JOIN_ON / P12 / COMPARISON | `cf.CFCIF` : `CHAR(19)` | `tmp.CIFNo` : `varchar(19)` | `cf.CFCIF` → `VARCHAR` | `tmp.CIFNo = cf.CFCIF` |
| IC-012379 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncAlamatKorespodensi:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncAlamatKorespodensi.sql#L27) | JOIN_ON / P12 / COMPARISON | `cf.CFCIF` : `CHAR(19)` | `tmp.CIFNo` : `varchar(19)` | `cf.CFCIF` → `VARCHAR` | `tmp.CIFNo = cf.CFCIF` |
| IC-012382 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncAlamatKorespodensi:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncAlamatKorespodensi.sql#L34) | JOIN_ON / P12 / COMPARISON | `cf.CFCIF` : `CHAR(19)` | `tmp.CIFNo` : `varchar(19)` | `cf.CFCIF` → `VARCHAR` | `tmp.CIFNo = cf.CFCIF` |
| IC-012385 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncAlamatKorespodensi:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncAlamatKorespodensi.sql#L41) | JOIN_ON / P12 / COMPARISON | `cf.CFCIF` : `CHAR(19)` | `tmp.CIFNo` : `varchar(19)` | `cf.CFCIF` → `VARCHAR` | `tmp.CIFNo = cf.CFCIF` |
| IC-012387 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncAlamatKorespodensi:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncAlamatKorespodensi.sql#L48) | JOIN_ON / P12 / COMPARISON | `cf.CFCIF` : `CHAR(19)` | `tmp.CIFNo` : `varchar(19)` | `cf.CFCIF` → `VARCHAR` | `tmp.CIFNo = cf.CFCIF` |
| IC-012389 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncAlamatKorespodensi:55](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncAlamatKorespodensi.sql#L55) | JOIN_ON / P12 / COMPARISON | `cf.CFCIF` : `CHAR(19)` | `tmp.CIFNo` : `varchar(19)` | `cf.CFCIF` → `VARCHAR` | `tmp.CIFNo = cf.CFCIF` |
| IC-012391 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncAlamatKorespodensi:62](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncAlamatKorespodensi.sql#L62) | JOIN_ON / P12 / COMPARISON | `cf.CFCIF` : `CHAR(19)` | `tmp.CIFNo` : `varchar(19)` | `cf.CFCIF` → `VARCHAR` | `tmp.CIFNo = cf.CFCIF` |
| IC-012495 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdateDataMasterNasabah:252](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateDataMasterNasabah.sql#L252) | JOIN_ON / P12 / COMPARISON | `addr.CFCIF` : `CHAR(19)` | `tmp.CIFNo` : `varchar(19)` | `addr.CFCIF` → `VARCHAR` | `tmp.CIFNo = addr.CFCIF` |
| IC-012499 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdateDataMasterNasabah:261](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateDataMasterNasabah.sql#L261) | JOIN_ON / P12 / COMPARISON | `addr.CFCIF` : `CHAR(19)` | `tmp.CIFNo` : `varchar(19)` | `addr.CFCIF` → `VARCHAR` | `tmp.CIFNo = addr.CFCIF` |
| IC-012503 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdateDataMasterNasabah:270](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateDataMasterNasabah.sql#L270) | JOIN_ON / P12 / COMPARISON | `addr.CFCIF` : `CHAR(19)` | `tmp.CIFNo` : `varchar(19)` | `addr.CFCIF` → `VARCHAR` | `tmp.CIFNo = addr.CFCIF` |
| IC-013595 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_PopulateAlamatNasabah:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateAlamatNasabah.sql#L45) | WHERE / P12 / COMPARISON | `CFCIF` : `CHAR(19)` | `@pcCIFNo` : `varchar(19)` | `CFCIF` → `VARCHAR` | `CFCIF = @pcCIFNo` |

<a id="kolom-cfforn"></a>
## CFFORN

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFADDR.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000838 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L115) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFFORN` : `CHAR(1)` | `@cCFFORN` : `varchar(1)` | `CFFORN` → `varchar(1)` | `@cCFFORN = CFFORN` |
| IC-000863 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:161](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L161) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFFORN` : `CHAR(1)` | `@cCFFORN` : `varchar(1)` | `CFFORN` → `varchar(1)` | `@cCFFORN = CFFORN` |
| IC-000878 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:201](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L201) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFFORN` : `CHAR(1)` | `@tmp_address.CFFORN` : `varchar(1)` | `CFFORN` → `varchar(1)` | `@tmp_address.CFFORN <- CFFORN` |

<a id="kolom-cfinvc"></a>
## CFINVC

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFADDR.sql#L13). **Existing leading index:** idx_cfaddr_cfinvc.

**Frekuensi source:** 56 lokasi / 12 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000849 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMValidateCustomer:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L131) | WHERE / P12 / COMPARISON | `CFINVC` : `CHAR(1)` | `'S'` : `varchar` | `CFINVC` → `VARCHAR` | `CFINVC = 'S'` |
| IC-000851 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMValidateCustomer:136](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L136) | WHERE / P12 / COMPARISON | `CFINVC` : `CHAR(1)` | `'G'` : `varchar` | `CFINVC` → `VARCHAR` | `CFINVC = 'G'` |
| IC-000853 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMValidateCustomer:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L141) | WHERE / P12 / COMPARISON | `CFINVC` : `CHAR(1)` | `'I'` : `varchar` | `CFINVC` → `VARCHAR` | `CFINVC = 'I'` |
| IC-003712 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:341](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L341) | CASE / P12 / COMPARISON | `cfaddr.CFINVC` : `CHAR(1)` | `'S'` : `varchar` | `cfaddr.CFINVC` → `VARCHAR` | `cfaddr.CFINVC = 'S'` |
| IC-003713 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:342](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L342) | CASE / P12 / COMPARISON | `cfaddr.CFINVC` : `CHAR(1)` | `'G'` : `varchar` | `cfaddr.CFINVC` → `VARCHAR` | `cfaddr.CFINVC = 'G'` |
| IC-003714 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:343](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L343) | CASE / P12 / COMPARISON | `cfaddr.CFINVC` : `CHAR(1)` | `'I'` : `varchar` | `cfaddr.CFINVC` → `VARCHAR` | `cfaddr.CFINVC = 'I'` |
| IC-003715 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:344](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L344) | CASE / P12 / COMPARISON | `cfaddr.CFINVC` : `CHAR(1)` | `'J'` : `varchar` | `cfaddr.CFINVC` → `VARCHAR` | `cfaddr.CFINVC = 'J'` |
| IC-003814 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:319](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L319) | CASE / P12 / COMPARISON | `cfaddr.CFINVC` : `CHAR(1)` | `'S'` : `varchar` | `cfaddr.CFINVC` → `VARCHAR` | `cfaddr.CFINVC = 'S'` |
| IC-003815 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:320](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L320) | CASE / P12 / COMPARISON | `cfaddr.CFINVC` : `CHAR(1)` | `'G'` : `varchar` | `cfaddr.CFINVC` → `VARCHAR` | `cfaddr.CFINVC = 'G'` |
| IC-003816 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:321](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L321) | CASE / P12 / COMPARISON | `cfaddr.CFINVC` : `CHAR(1)` | `'I'` : `varchar` | `cfaddr.CFINVC` → `VARCHAR` | `cfaddr.CFINVC = 'I'` |
| IC-003817 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:322](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L322) | CASE / P12 / COMPARISON | `cfaddr.CFINVC` : `CHAR(1)` | `'J'` : `varchar` | `cfaddr.CFINVC` → `VARCHAR` | `cfaddr.CFINVC = 'J'` |
| IC-003865 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:284](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L284) | CASE / P12 / COMPARISON | `cfaddr.CFINVC` : `CHAR(1)` | `'S'` : `varchar` | `cfaddr.CFINVC` → `VARCHAR` | `cfaddr.CFINVC = 'S'` |
| IC-003866 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:285](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L285) | CASE / P12 / COMPARISON | `cfaddr.CFINVC` : `CHAR(1)` | `'G'` : `varchar` | `cfaddr.CFINVC` → `VARCHAR` | `cfaddr.CFINVC = 'G'` |
| IC-003867 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:286](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L286) | CASE / P12 / COMPARISON | `cfaddr.CFINVC` : `CHAR(1)` | `'I'` : `varchar` | `cfaddr.CFINVC` → `VARCHAR` | `cfaddr.CFINVC = 'I'` |
| IC-003868 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:287](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L287) | CASE / P12 / COMPARISON | `cfaddr.CFINVC` : `CHAR(1)` | `'J'` : `varchar` | `cfaddr.CFINVC` → `VARCHAR` | `cfaddr.CFINVC = 'J'` |
| IC-003949 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusiOrchid:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L265) | CASE / P12 / COMPARISON | `cfaddr.CFINVC` : `CHAR(1)` | `'S'` : `varchar` | `cfaddr.CFINVC` → `VARCHAR` | `cfaddr.CFINVC = 'S'` |
| IC-003950 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusiOrchid:266](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L266) | CASE / P12 / COMPARISON | `cfaddr.CFINVC` : `CHAR(1)` | `'G'` : `varchar` | `cfaddr.CFINVC` → `VARCHAR` | `cfaddr.CFINVC = 'G'` |
| IC-003951 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusiOrchid:267](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L267) | CASE / P12 / COMPARISON | `cfaddr.CFINVC` : `CHAR(1)` | `'I'` : `varchar` | `cfaddr.CFINVC` → `VARCHAR` | `cfaddr.CFINVC = 'I'` |
| IC-003952 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusiOrchid:268](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L268) | CASE / P12 / COMPARISON | `cfaddr.CFINVC` : `CHAR(1)` | `'J'` : `varchar` | `cfaddr.CFINVC` → `VARCHAR` | `cfaddr.CFINVC = 'J'` |
| IC-005345 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDataAlamatKonfirmasi:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDataAlamatKonfirmasi.sql#L13) | WHERE / P12 / COMPARISON | `CFINVC` : `CHAR(1)` | `'X'` : `varchar` | `CFINVC` → `VARCHAR` | `CFINVC = 'X'` |
| IC-006414 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateStatusEmailKonfirmasiProduct:163](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateStatusEmailKonfirmasiProduct.sql#L163) | WHERE / P12 / COMPARISON | `cfa.CFINVC` : `CHAR(1)` | `'X'` : `varchar` | `cfa.CFINVC` → `VARCHAR` | `cfa.CFINVC = 'X'` |
| IC-009195 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:505](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L505) | CASE / P12 / COMPARISON | `ca.CFINVC` : `CHAR(1)` | `'T'` : `varchar` | `ca.CFINVC` → `VARCHAR` | `ca.CFINVC = 'T'` |
| IC-009196 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:506](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L506) | CASE / P12 / COMPARISON | `ca.CFINVC` : `CHAR(1)` | `'J'` : `varchar` | `ca.CFINVC` → `VARCHAR` | `ca.CFINVC = 'J'` |
| IC-009198 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:507](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L507) | CASE / P12 / COMPARISON | `ca.CFINVC` : `CHAR(1)` | `'S'` : `varchar` | `ca.CFINVC` → `VARCHAR` | `ca.CFINVC = 'S'` |
| IC-009201 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:508](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L508) | CASE / P12 / COMPARISON | `ca.CFINVC` : `CHAR(1)` | `'J'` : `varchar` | `ca.CFINVC` → `VARCHAR` | `ca.CFINVC = 'J'` |
| IC-009204 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:509](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L509) | CASE / P12 / COMPARISON | `ca.CFINVC` : `CHAR(1)` | `'S'` : `varchar` | `ca.CFINVC` → `VARCHAR` | `ca.CFINVC = 'S'` |
| IC-009206 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:510](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L510) | CASE / P12 / COMPARISON | `ca.CFINVC` : `CHAR(1)` | `'G'` : `varchar` | `ca.CFINVC` → `VARCHAR` | `ca.CFINVC = 'G'` |
| IC-009207 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:511](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L511) | CASE / P12 / COMPARISON | `ca.CFINVC` : `CHAR(1)` | `'I'` : `varchar` | `ca.CFINVC` → `VARCHAR` | `ca.CFINVC = 'I'` |
| IC-009210 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:541](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L541) | CASE / P12 / COMPARISON | `ca.CFINVC` : `CHAR(1)` | `'S'` : `varchar` | `ca.CFINVC` → `VARCHAR` | `ca.CFINVC = 'S'` |
| IC-009211 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:542](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L542) | CASE / P12 / COMPARISON | `ca.CFINVC` : `CHAR(1)` | `'G'` : `varchar` | `ca.CFINVC` → `VARCHAR` | `ca.CFINVC = 'G'` |
| IC-009212 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:543](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L543) | CASE / P12 / COMPARISON | `ca.CFINVC` : `CHAR(1)` | `'I'` : `varchar` | `ca.CFINVC` → `VARCHAR` | `ca.CFINVC = 'I'` |
| IC-009213 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:544](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L544) | CASE / P12 / COMPARISON | `ca.CFINVC` : `CHAR(1)` | `'J'` : `varchar` | `ca.CFINVC` → `VARCHAR` | `ca.CFINVC = 'J'` |
| IC-009364 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:402](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L402) | CASE / P12 / COMPARISON | `ca.CFINVC` : `CHAR(1)` | `'T'` : `varchar` | `ca.CFINVC` → `VARCHAR` | `ca.CFINVC = 'T'` |
| IC-009365 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:403](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L403) | CASE / P12 / COMPARISON | `ca.CFINVC` : `CHAR(1)` | `'J'` : `varchar` | `ca.CFINVC` → `VARCHAR` | `ca.CFINVC = 'J'` |
| IC-009367 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:404](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L404) | CASE / P12 / COMPARISON | `ca.CFINVC` : `CHAR(1)` | `'S'` : `varchar` | `ca.CFINVC` → `VARCHAR` | `ca.CFINVC = 'S'` |
| IC-009370 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:405](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L405) | CASE / P12 / COMPARISON | `ca.CFINVC` : `CHAR(1)` | `'J'` : `varchar` | `ca.CFINVC` → `VARCHAR` | `ca.CFINVC = 'J'` |
| IC-009373 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:406](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L406) | CASE / P12 / COMPARISON | `ca.CFINVC` : `CHAR(1)` | `'S'` : `varchar` | `ca.CFINVC` → `VARCHAR` | `ca.CFINVC = 'S'` |
| IC-009375 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:407](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L407) | CASE / P12 / COMPARISON | `ca.CFINVC` : `CHAR(1)` | `'G'` : `varchar` | `ca.CFINVC` → `VARCHAR` | `ca.CFINVC = 'G'` |
| IC-009376 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:408](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L408) | CASE / P12 / COMPARISON | `ca.CFINVC` : `CHAR(1)` | `'I'` : `varchar` | `ca.CFINVC` → `VARCHAR` | `ca.CFINVC = 'I'` |
| IC-009379 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:438](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L438) | CASE / P12 / COMPARISON | `ca.CFINVC` : `CHAR(1)` | `'S'` : `varchar` | `ca.CFINVC` → `VARCHAR` | `ca.CFINVC = 'S'` |
| IC-009380 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:439](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L439) | CASE / P12 / COMPARISON | `ca.CFINVC` : `CHAR(1)` | `'G'` : `varchar` | `ca.CFINVC` → `VARCHAR` | `ca.CFINVC = 'G'` |
| IC-009381 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:440](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L440) | CASE / P12 / COMPARISON | `ca.CFINVC` : `CHAR(1)` | `'I'` : `varchar` | `ca.CFINVC` → `VARCHAR` | `ca.CFINVC = 'I'` |
| IC-009382 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:441](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L441) | CASE / P12 / COMPARISON | `ca.CFINVC` : `CHAR(1)` | `'J'` : `varchar` | `ca.CFINVC` → `VARCHAR` | `ca.CFINVC = 'J'` |
| IC-012378 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncAlamatKorespodensi:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncAlamatKorespodensi.sql#L21) | WHERE / P12 / COMPARISON | `CFINVC` : `CHAR(1)` | `'T'` : `varchar` | `CFINVC` → `VARCHAR` | `CFINVC = 'T'` |
| IC-012380 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncAlamatKorespodensi:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncAlamatKorespodensi.sql#L28) | WHERE / P12 / COMPARISON | `CFINVC` : `CHAR(1)` | `'J'` : `varchar` | `CFINVC` → `VARCHAR` | `CFINVC = 'J'` |
| IC-012383 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncAlamatKorespodensi:35](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncAlamatKorespodensi.sql#L35) | WHERE / P12 / COMPARISON | `CFINVC` : `CHAR(1)` | `'S'` : `varchar` | `CFINVC` → `VARCHAR` | `CFINVC = 'S'` |
| IC-012386 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncAlamatKorespodensi:42](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncAlamatKorespodensi.sql#L42) | WHERE / P12 / COMPARISON | `CFINVC` : `CHAR(1)` | `'J'` : `varchar` | `CFINVC` → `VARCHAR` | `CFINVC = 'J'` |
| IC-012388 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncAlamatKorespodensi:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncAlamatKorespodensi.sql#L49) | WHERE / P12 / COMPARISON | `CFINVC` : `CHAR(1)` | `'S'` : `varchar` | `CFINVC` → `VARCHAR` | `CFINVC = 'S'` |
| IC-012390 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncAlamatKorespodensi:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncAlamatKorespodensi.sql#L56) | WHERE / P12 / COMPARISON | `CFINVC` : `CHAR(1)` | `'G'` : `varchar` | `CFINVC` → `VARCHAR` | `CFINVC = 'G'` |
| IC-012392 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncAlamatKorespodensi:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncAlamatKorespodensi.sql#L63) | WHERE / P12 / COMPARISON | `CFINVC` : `CHAR(1)` | `'I'` : `varchar` | `CFINVC` → `VARCHAR` | `CFINVC = 'I'` |
| IC-012500 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdateDataMasterNasabah:262](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateDataMasterNasabah.sql#L262) | JOIN_ON / P12 / COMPARISON | `addr.CFINVC` : `CHAR(1)` | `'S'` : `varchar` | `addr.CFINVC` → `VARCHAR` | `addr.CFINVC = 'S'` |
| IC-012504 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdateDataMasterNasabah:272](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateDataMasterNasabah.sql#L272) | WHERE / P12 / IN_VALUE_TYPE | `addr.CFINVC` : `CHAR(1)` | `'G'` : `varchar` | `addr.CFINVC` → `VARCHAR` | `addr.CFINVC in ( 'G', 'I', 'J' )` |
| IC-013591 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_PopulateAlamatNasabah:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateAlamatNasabah.sql#L31) | CASE / P12 / COMPARISON | `CFINVC` : `CHAR(1)` | `'S'` : `varchar` | `CFINVC` → `VARCHAR` | `CFINVC = 'S'` |
| IC-013592 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_PopulateAlamatNasabah:32](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateAlamatNasabah.sql#L32) | CASE / P12 / COMPARISON | `CFINVC` : `CHAR(1)` | `'G'` : `varchar` | `CFINVC` → `VARCHAR` | `CFINVC = 'G'` |
| IC-013593 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_PopulateAlamatNasabah:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateAlamatNasabah.sql#L33) | CASE / P12 / COMPARISON | `CFINVC` : `CHAR(1)` | `'I'` : `varchar` | `CFINVC` → `VARCHAR` | `CFINVC = 'I'` |
| IC-013594 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_PopulateAlamatNasabah:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateAlamatNasabah.sql#L34) | CASE / P12 / COMPARISON | `CFINVC` : `CHAR(1)` | `'J'` : `varchar` | `CFINVC` → `VARCHAR` | `CFINVC = 'J'` |

<a id="kolom-cfmail"></a>
## CFMAIL

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFADDR.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000840 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:118](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L118) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFMAIL` : `CHAR(1)` | `@cCFMAIL` : `varchar(1)` | `CFMAIL` → `varchar(1)` | `@cCFMAIL = CFMAIL` |
| IC-000865 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:164](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L164) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFMAIL` : `CHAR(1)` | `@cCFMAIL` : `varchar(1)` | `CFMAIL` → `varchar(1)` | `@cCFMAIL = CFMAIL` |
| IC-000879 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:201](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L201) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFMAIL` : `CHAR(1)` | `@tmp_address.CFMAIL` : `varchar(1)` | `CFMAIL` → `varchar(1)` | `@tmp_address.CFMAIL <- CFMAIL` |

<a id="kolom-cfna2"></a>
## CFNA2

**Tipe dasar:** `CHAR(40)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFADDR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000829 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L105) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFNA2` : `CHAR(40)` | `@cAlamatIdentitas1` : `varchar(40)` | `CFNA2` → `varchar(40)` | `@cAlamatIdentitas1 = CFNA2` |
| IC-000833 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:109](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L109) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFNA2` : `CHAR(40)` | `@cCFNA2` : `varchar(40)` | `CFNA2` → `varchar(40)` | `@cCFNA2 = CFNA2` |
| IC-000854 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:151](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L151) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFNA2` : `CHAR(40)` | `@cAlamatIdentitas1` : `varchar(40)` | `CFNA2` → `varchar(40)` | `@cAlamatIdentitas1 = CFNA2` |
| IC-000858 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:155](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L155) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFNA2` : `CHAR(40)` | `@cCFNA2` : `varchar(40)` | `CFNA2` → `varchar(40)` | `@cCFNA2 = CFNA2` |
| IC-000880 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:201](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L201) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFNA2` : `CHAR(40)` | `@tmp_address.CFNA2` : `varchar(40)` | `CFNA2` → `varchar(40)` | `@tmp_address.CFNA2 <- CFNA2` |
| IC-006412 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateStatusEmailKonfirmasiProduct:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateStatusEmailKonfirmasiProduct.sql#L156) | DML_CONDITION / P05 / BINARY_Add | `cfa.CFNA2` : `CHAR(40)` | `' '` : `varchar` | `cfa.CFNA2` → `VARCHAR` | `cfa.CFNA2 +' '` |

<a id="kolom-cfna3"></a>
## CFNA3

**Tipe dasar:** `CHAR(40)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFADDR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000830 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L106) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFNA3` : `CHAR(40)` | `@cAlamatIdentitas2` : `varchar(40)` | `CFNA3` → `varchar(40)` | `@cAlamatIdentitas2 = CFNA3` |
| IC-000834 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L110) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFNA3` : `CHAR(40)` | `@cCFNA3` : `varchar(40)` | `CFNA3` → `varchar(40)` | `@cCFNA3 = CFNA3` |
| IC-000855 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:152](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L152) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFNA3` : `CHAR(40)` | `@cAlamatIdentitas2` : `varchar(40)` | `CFNA3` → `varchar(40)` | `@cAlamatIdentitas2 = CFNA3` |
| IC-000859 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L156) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFNA3` : `CHAR(40)` | `@cCFNA3` : `varchar(40)` | `CFNA3` → `varchar(40)` | `@cCFNA3 = CFNA3` |
| IC-000881 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:201](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L201) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFNA3` : `CHAR(40)` | `@tmp_address.CFNA3` : `varchar(40)` | `CFNA3` → `varchar(40)` | `@tmp_address.CFNA3 <- CFNA3` |

<a id="kolom-cfna4"></a>
## CFNA4

**Tipe dasar:** `CHAR(40)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFADDR.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000835 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L111) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFNA4` : `CHAR(40)` | `@cCFNA4` : `varchar(40)` | `CFNA4` → `varchar(40)` | `@cCFNA4 = CFNA4` |
| IC-000860 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:157](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L157) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFNA4` : `CHAR(40)` | `@cCFNA4` : `varchar(40)` | `CFNA4` → `varchar(40)` | `@cCFNA4 = CFNA4` |
| IC-000882 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:201](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L201) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFNA4` : `CHAR(40)` | `@tmp_address.CFNA4` : `varchar(40)` | `CFNA4` → `varchar(40)` | `@tmp_address.CFNA4 <- CFNA4` |

<a id="kolom-cfna5"></a>
## CFNA5

**Tipe dasar:** `CHAR(40)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFADDR.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000831 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L107) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFNA5` : `CHAR(40)` | `@cAlamatIdentitas3` : `varchar(40)` | `CFNA5` → `varchar(40)` | `@cAlamatIdentitas3 = CFNA5` |
| IC-000836 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:112](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L112) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFNA5` : `CHAR(40)` | `@cCFNA5` : `varchar(40)` | `CFNA5` → `varchar(40)` | `@cCFNA5 = CFNA5` |
| IC-000856 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:153](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L153) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFNA5` : `CHAR(40)` | `@cAlamatIdentitas3` : `varchar(40)` | `CFNA5` → `varchar(40)` | `@cAlamatIdentitas3 = CFNA5` |
| IC-000861 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:158](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L158) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFNA5` : `CHAR(40)` | `@cCFNA5` : `varchar(40)` | `CFNA5` → `varchar(40)` | `@cCFNA5 = CFNA5` |
| IC-000883 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:201](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L201) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFNA5` : `CHAR(40)` | `@tmp_address.CFNA5` : `varchar(40)` | `CFNA5` → `varchar(40)` | `@tmp_address.CFNA5 <- CFNA5` |

<a id="kolom-cfstsn"></a>
## CFSTSN

**Tipe dasar:** `CHAR(5)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFADDR.sql#L28). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000842 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:120](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L120) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFSTSN` : `CHAR(5)` | `@cCFSTSN` : `varchar(5)` | `CFSTSN` → `varchar(5)` | `@cCFSTSN = CFSTSN` |
| IC-000867 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:166](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L166) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFSTSN` : `CHAR(5)` | `@cCFSTSN` : `varchar(5)` | `CFSTSN` → `varchar(5)` | `@cCFSTSN = CFSTSN` |
| IC-000884 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:201](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L201) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFSTSN` : `CHAR(5)` | `@tmp_address.CFSTSN` : `varchar(5)` | `CFSTSN` → `varchar(5)` | `@tmp_address.CFSTSN <- CFSTSN` |

<a id="kolom-cfuse"></a>
## CFUSE

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFADDR.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 29 lokasi / 10 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000839 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:116](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L116) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFUSE` : `CHAR(1)` | `@cCFUSE` : `varchar(1)` | `CFUSE` → `varchar(1)` | `@cCFUSE = CFUSE` |
| IC-000864 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L162) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFUSE` : `CHAR(1)` | `@cCFUSE` : `varchar(1)` | `CFUSE` → `varchar(1)` | `@cCFUSE = CFUSE` |
| IC-000885 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:201](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L201) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFUSE` : `CHAR(1)` | `@tmp_address.CFUSE` : `varchar(1)` | `CFUSE` → `varchar(1)` | `@tmp_address.CFUSE <- CFUSE` |
| IC-000828 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMValidateCustomer:101](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L101) | WHERE / P12 / COMPARISON | `CFUSE` : `CHAR(1)` | `'Y'` : `varchar` | `CFUSE` → `VARCHAR` | `CFUSE = 'Y'` |
| IC-000847 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMValidateCustomer:126](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L126) | WHERE / P12 / COMPARISON | `CFUSE` : `CHAR(1)` | `'Y'` : `varchar` | `CFUSE` → `VARCHAR` | `CFUSE = 'Y'` |
| IC-003711 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L337) | CASE / P12 / COMPARISON | `cfaddr.CFUSE` : `CHAR(1)` | `'Y'` : `varchar` | `cfaddr.CFUSE` → `VARCHAR` | `cfaddr.CFUSE = 'Y'` |
| IC-003813 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:315](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L315) | CASE / P12 / COMPARISON | `cfaddr.CFUSE` : `CHAR(1)` | `'Y'` : `varchar` | `cfaddr.CFUSE` → `VARCHAR` | `cfaddr.CFUSE = 'Y'` |
| IC-003864 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:280](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L280) | CASE / P12 / COMPARISON | `cfaddr.CFUSE` : `CHAR(1)` | `'Y'` : `varchar` | `cfaddr.CFUSE` → `VARCHAR` | `cfaddr.CFUSE = 'Y'` |
| IC-003948 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusiOrchid:261](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L261) | CASE / P12 / COMPARISON | `cfaddr.CFUSE` : `CHAR(1)` | `'Y'` : `varchar` | `cfaddr.CFUSE` → `VARCHAR` | `cfaddr.CFUSE = 'Y'` |
| IC-009197 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:506](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L506) | CASE / P12 / COMPARISON | `ca.CFUSE` : `CHAR(1)` | `'Y'` : `varchar` | `ca.CFUSE` → `VARCHAR` | `ca.CFUSE = 'Y'` |
| IC-009199 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:507](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L507) | CASE / P12 / COMPARISON | `ca.CFUSE` : `CHAR(1)` | `'Y'` : `varchar` | `ca.CFUSE` → `VARCHAR` | `ca.CFUSE = 'Y'` |
| IC-009200 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRetailSyncDataCore:508](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L508) | CASE / P12 / COMPARISON | `ISNULL(ca.CFUSE,'')` : `CHAR(1)` | `'Y'` : `varchar` | `ISNULL(ca.CFUSE,'')` → `VARCHAR` | `ISNULL(ca.CFUSE,'') <> 'Y'` |
| IC-009203 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRetailSyncDataCore:509](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L509) | CASE / P12 / COMPARISON | `ISNULL(ca.CFUSE,'')` : `CHAR(1)` | `'Y'` : `varchar` | `ISNULL(ca.CFUSE,'')` → `VARCHAR` | `ISNULL(ca.CFUSE,'') <> 'Y'` |
| IC-009208 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:515](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L515) | CASE / P12 / COMPARISON | `ca.CFUSE` : `CHAR(1)` | `'Y'` : `varchar` | `ca.CFUSE` → `VARCHAR` | `ca.CFUSE = 'Y'` |
| IC-009209 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:537](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L537) | CASE / P12 / COMPARISON | `ca.CFUSE` : `CHAR(1)` | `'Y'` : `varchar` | `ca.CFUSE` → `VARCHAR` | `ca.CFUSE = 'Y'` |
| IC-009366 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:403](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L403) | CASE / P12 / COMPARISON | `ca.CFUSE` : `CHAR(1)` | `'Y'` : `varchar` | `ca.CFUSE` → `VARCHAR` | `ca.CFUSE = 'Y'` |
| IC-009368 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:404](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L404) | CASE / P12 / COMPARISON | `ca.CFUSE` : `CHAR(1)` | `'Y'` : `varchar` | `ca.CFUSE` → `VARCHAR` | `ca.CFUSE = 'Y'` |
| IC-009369 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRetailSyncDataCoreWorkingHours:405](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L405) | CASE / P12 / COMPARISON | `ISNULL(ca.CFUSE,'')` : `CHAR(1)` | `'Y'` : `varchar` | `ISNULL(ca.CFUSE,'')` → `VARCHAR` | `ISNULL(ca.CFUSE,'') <> 'Y'` |
| IC-009372 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRetailSyncDataCoreWorkingHours:406](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L406) | CASE / P12 / COMPARISON | `ISNULL(ca.CFUSE,'')` : `CHAR(1)` | `'Y'` : `varchar` | `ISNULL(ca.CFUSE,'')` → `VARCHAR` | `ISNULL(ca.CFUSE,'') <> 'Y'` |
| IC-009377 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:412](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L412) | CASE / P12 / COMPARISON | `ca.CFUSE` : `CHAR(1)` | `'Y'` : `varchar` | `ca.CFUSE` → `VARCHAR` | `ca.CFUSE = 'Y'` |
| IC-009378 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:434](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L434) | CASE / P12 / COMPARISON | `ca.CFUSE` : `CHAR(1)` | `'Y'` : `varchar` | `ca.CFUSE` → `VARCHAR` | `ca.CFUSE = 'Y'` |
| IC-009593 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptCustomerPortfolio:91](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerPortfolio.sql#L91) | JOIN_ON / P12 / COMPARISON | `CFADDR.CFUSE` : `CHAR(1)` | `'Y'` : `varchar` | `CFADDR.CFUSE` → `VARCHAR` | `CFADDR.CFUSE = 'Y'` |
| IC-012381 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncAlamatKorespodensi:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncAlamatKorespodensi.sql#L28) | WHERE / P12 / COMPARISON | `CFUSE` : `CHAR(1)` | `'Y'` : `varchar` | `CFUSE` → `VARCHAR` | `CFUSE = 'Y'` |
| IC-012384 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncAlamatKorespodensi:35](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncAlamatKorespodensi.sql#L35) | WHERE / P12 / COMPARISON | `CFUSE` : `CHAR(1)` | `'Y'` : `varchar` | `CFUSE` → `VARCHAR` | `CFUSE = 'Y'` |
| IC-012496 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdateDataMasterNasabah:253](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateDataMasterNasabah.sql#L253) | JOIN_ON / P12 / COMPARISON | `addr.CFUSE` : `CHAR(1)` | `'Y'` : `varchar` | `addr.CFUSE` → `VARCHAR` | `addr.CFUSE = 'Y'` |
| IC-009202 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRetailSyncDataCore:508](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L508) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ca.CFUSE` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `ISNULL(ca.CFUSE,'')` |
| IC-009205 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRetailSyncDataCore:509](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L509) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ca.CFUSE` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `ISNULL(ca.CFUSE,'')` |
| IC-009371 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRetailSyncDataCoreWorkingHours:405](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L405) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ca.CFUSE` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `ISNULL(ca.CFUSE,'')` |
| IC-009374 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRetailSyncDataCoreWorkingHours:406](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L406) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ca.CFUSE` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `ISNULL(ca.CFUSE,'')` |

<a id="kolom-cfzip"></a>
## CFZIP

**Tipe dasar:** `CHAR(9)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFADDR.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 10 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000837 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L114) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFZIP` : `CHAR(9)` | `@cCFZIP` : `varchar(9)` | `CFZIP` → `varchar(9)` | `@cCFZIP = CFZIP` |
| IC-000862 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:160](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L160) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFZIP` : `CHAR(9)` | `@cCFZIP` : `varchar(9)` | `CFZIP` → `varchar(9)` | `@cCFZIP = CFZIP` |
| IC-000886 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMValidateCustomer:201](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L201) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFZIP` : `CHAR(9)` | `@tmp_address.CFZIP` : `varchar(9)` | `CFZIP` → `varchar(9)` | `@tmp_address.CFZIP <- CFZIP` |
| IC-003709 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailIndividu:306](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L306) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `cfaddr.CFZIP` : `CHAR(9)` | `#xtmpAddress.KodePos` : `VARCHAR(6)` | `cfaddr.CFZIP` → `VARCHAR(6)` | `#xtmpAddress.KodePos <- cfaddr.CFZIP` |
| IC-003811 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailIndividuOrchid:284](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L284) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `cfaddr.CFZIP` : `CHAR(9)` | `#xtmpAddress.KodePos` : `VARCHAR(5)` | `cfaddr.CFZIP` → `VARCHAR(5)` | `#xtmpAddress.KodePos <- cfaddr.CFZIP` |
| IC-003862 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailInstitusi:249](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L249) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `cfaddr.CFZIP` : `CHAR(9)` | `#xtmpAddress.KodePos` : `VARCHAR(5)` | `cfaddr.CFZIP` → `VARCHAR(5)` | `#xtmpAddress.KodePos <- cfaddr.CFZIP` |
| IC-003946 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailInstitusiOrchid:230](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L230) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `cfaddr.CFZIP` : `CHAR(9)` | `#xtmpAddress.KodePos` : `VARCHAR(5)` | `cfaddr.CFZIP` → `VARCHAR(5)` | `#xtmpAddress.KodePos <- cfaddr.CFZIP` |
| IC-012494 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSUpdateDataMasterNasabah:249](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateDataMasterNasabah.sql#L249) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `addr.CFZIP` : `CHAR(9)` | `tmp.ZipCode` : `varchar(10)` | `addr.CFZIP` → `varchar(10)` | `tmp.ZipCode = addr.CFZIP` |
| IC-012498 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSUpdateDataMasterNasabah:258](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateDataMasterNasabah.sql#L258) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `addr.CFZIP` : `CHAR(9)` | `tmp.ZipCode` : `varchar(10)` | `addr.CFZIP` → `varchar(10)` | `tmp.ZipCode = addr.CFZIP` |
| IC-012502 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSUpdateDataMasterNasabah:267](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateDataMasterNasabah.sql#L267) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `addr.CFZIP` : `CHAR(9)` | `tmp.ZipCode` : `varchar(10)` | `addr.CFZIP` → `varchar(10)` | `tmp.ZipCode = addr.CFZIP` |
