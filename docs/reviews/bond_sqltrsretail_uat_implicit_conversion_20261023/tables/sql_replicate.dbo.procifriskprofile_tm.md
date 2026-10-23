# SQL_REPLICATE.dbo.ProCIFRiskProfile_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 14 lokasi source, 6 procedure, 3 kolom. **A:** 6 lokasi pada 5 procedure. **B:** 2 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cifnumber"></a>
## CIFNumber

**Tipe dasar:** `VARCHAR(19)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.ProCIFRiskProfile_TM.sql#L2). **Existing leading index:** ProCIFRiskProfile_TM_idx1.

**Frekuensi source:** 8 lokasi / 5 SP. **A:** 6 lokasi / 5 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005995 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateRiskProfileNasabah:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateRiskProfileNasabah.sql#L28) | WHERE / P01 / COMPARISON | `CIFNumber` : `VARCHAR(19)` | `cast(@pcCIF as bigint)` : `bigint` | `CIFNumber` → `BIGINT` | `CIFNumber = cast(@pcCIF as bigint)` |
| IC-012417 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSSyncMasterRiskProfileData:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterRiskProfileData.sql#L24) | JOIN_ON / P01 / COMPARISON | `so.CIFNumber` : `VARCHAR(19)` | `CFMAST.CFCIF#` : `BIGINT` | `so.CIFNumber` → `BIGINT` | `CFMAST.CFCIF# = so.CIFNumber` |
| IC-012483 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSUpdateDataMasterNasabah:163](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateDataMasterNasabah.sql#L163) | JOIN_ON / P01 / COMPARISON | `cm.CIFNumber` : `VARCHAR(19)` | `t.CFCIF` : `BIGINT` | `cm.CIFNumber` → `BIGINT` | `t.CFCIF = cm.CIFNumber` |
| IC-012630 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSUpdateRiskProfile:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateRiskProfile.sql#L25) | WHERE / P01 / COMPARISON | `CIFNumber` : `VARCHAR(19)` | `@nCIFNo` : `bigint` | `CIFNumber` → `BIGINT` | `CIFNumber = @nCIFNo` |
| IC-012631 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSUpdateRiskProfile:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateRiskProfile.sql#L33) | WHERE / P01 / COMPARISON | `CIFNumber` : `VARCHAR(19)` | `@nCIFNo` : `bigint` | `CIFNumber` → `BIGINT` | `CIFNumber = @nCIFNo` |
| IC-013178 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:967](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L967) | WHERE / P01 / COMPARISON | `CIFNumber` : `VARCHAR(19)` | `@nCIFNo` : `bigint` | `CIFNumber` → `BIGINT` | `CIFNumber = @nCIFNo` |
| IC-012632 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateRiskProfile:35](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateRiskProfile.sql#L35) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sql_replicate.dbo.procifriskprofile_tm.CIFNumber` : `VARCHAR(19)` | `@nCIFNo` : `bigint` | `@nCIFNo` → `VARCHAR(19)` | `sql_replicate.dbo.procifriskprofile_tm.CIFNumber <- @nCIFNo` |
| IC-013179 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:969](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L969) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sql_replicate.dbo.procifriskprofile_tm.CIFNumber` : `VARCHAR(19)` | `@nCIFNo` : `bigint` | `@nCIFNo` → `VARCHAR(19)` | `sql_replicate.dbo.procifriskprofile_tm.CIFNumber <- @nCIFNo` |

<a id="kolom-expireddate"></a>
## ExpiredDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.ProCIFRiskProfile_TM.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007888 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRPTMasterNasabahObligasi:781](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahObligasi.sql#L781) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `rp.ExpiredDate` : `DATETIME` | `rs.TanggalExpiredProfilRisiko` : `VARCHAR(55)` | `rp.ExpiredDate` → `VARCHAR(55)` | `rs.TanggalExpiredProfilRisiko = rp.ExpiredDate` |
| IC-012628 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateRiskProfile:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateRiskProfile.sql#L24) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ExpiredDate` : `DATETIME` | `@pdExpiredDate` : `varchar(8)` | `@pdExpiredDate` → `DATETIME` | `ExpiredDate = @pdExpiredDate` |
| IC-012633 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateRiskProfile:35](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateRiskProfile.sql#L35) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sql_replicate.dbo.procifriskprofile_tm.ExpiredDate` : `DATETIME` | `@pdExpiredDate` : `varchar(8)` | `@pdExpiredDate` → `DATETIME` | `sql_replicate.dbo.procifriskprofile_tm.ExpiredDate <- @pdExpiredDate` |

<a id="kolom-lastupdated"></a>
## LastUpdated

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.ProCIFRiskProfile_TM.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007887 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRPTMasterNasabahObligasi:780](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahObligasi.sql#L780) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `rp.LastUpdated` : `DATETIME` | `rs.TanggalLastUpdateProfilRisiko` : `VARCHAR(55)` | `rp.LastUpdated` → `VARCHAR(55)` | `rs.TanggalLastUpdateProfilRisiko = rp.LastUpdated` |
| IC-012629 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateRiskProfile:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateRiskProfile.sql#L24) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `LastUpdated` : `DATETIME` | `@pdLastUpdateDate` : `varchar(8)` | `@pdLastUpdateDate` → `DATETIME` | `LastUpdated = @pdLastUpdateDate` |
| IC-012634 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateRiskProfile:35](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateRiskProfile.sql#L35) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sql_replicate.dbo.procifriskprofile_tm.LastUpdated` : `DATETIME` | `@pdLastUpdateDate` : `varchar(8)` | `@pdLastUpdateDate` → `DATETIME` | `sql_replicate.dbo.procifriskprofile_tm.LastUpdated <- @pdLastUpdateDate` |
