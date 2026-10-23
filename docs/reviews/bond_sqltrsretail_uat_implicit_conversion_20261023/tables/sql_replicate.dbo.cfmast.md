# SQL_REPLICATE.dbo.CFMAST — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 216 lokasi source, 32 procedure, 21 kolom. **A:** 6 lokasi pada 4 procedure. **B:** 50 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cfcif"></a>
## CFCIF

**Tipe dasar:** `CHAR(19)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFMAST.sql#L6). **Existing leading index:** idx_cfmast_cfcif.

**Frekuensi source:** 57 lokasi / 29 SP. **A:** 4 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007897 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahPerubahanObligasi2:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql#L45) | JOIN_ON / P01 / COMPARISON | `c.CFCIF` : `CHAR(19)` | `a.CIFNumber` : `BIGINT` | `c.CFCIF` → `BIGINT` | `c.CFCIF=a.CIFNumber` |
| IC-007898 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahPerubahanObligasi2:57](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql#L57) | JOIN_ON / P01 / COMPARISON | `c.CFCIF` : `CHAR(19)` | `a.CIFNumber` : `BIGINT` | `c.CFCIF` → `BIGINT` | `c.CFCIF=a.CIFNumber` |
| IC-007899 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahPerubahanObligasi2:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql#L69) | JOIN_ON / P01 / COMPARISON | `c.CFCIF` : `CHAR(19)` | `a.CIFNumber` : `BIGINT` | `c.CFCIF` → `BIGINT` | `c.CFCIF=a.CIFNumber` |
| IC-012473 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSUpdateDataMasterNasabah:121](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateDataMasterNasabah.sql#L121) | JOIN_ON / P01 / COMPARISON | `c.CFCIF` : `CHAR(19)` | `t.CFCIF` : `BIGINT` | `c.CFCIF` → `BIGINT` | `t.CFCIF = c.CFCIF` |
| IC-001144 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSProsesDataJobNasabah:91](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L91) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `v.CFCIF` : `CHAR(19)` | `tmp.CIFNo` : `varchar(20)` | `v.CFCIF` → `varchar(20)` | `tmp.CIFNo = v.CFCIF` |
| IC-004757 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSOBLExpiredKitasRPT:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOBLExpiredKitasRPT.sql#L34) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CM.CFCIF` : `CHAR(19)` | `@tmpCif.CFCIF` : `VARCHAR(19)` | `CM.CFCIF` → `VARCHAR(19)` | `@tmpCif.CFCIF <- CM.CFCIF` |
| IC-012471 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSUpdateDataMasterNasabah:112](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateDataMasterNasabah.sql#L112) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `c.CFCIF` : `CHAR(19)` | `#tmpCFMAST.CFCIF` : `VARCHAR(19)` | `c.CFCIF` → `VARCHAR(19)` | `#tmpCFMAST.CFCIF <- c.CFCIF` |
| IC-000755 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMTrxSecPreBook:412](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMTrxSecPreBook.sql#L412) | WHERE / P12 / COMPARISON | `CFCIF` : `CHAR(19)` | `@cCIFNo` : `varchar(30)` | `CFCIF` → `VARCHAR` | `CFCIF = @cCIFNo` |
| IC-001056 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSPopulateConsentNasabahPTOS:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSPopulateConsentNasabahPTOS.sql#L85) | JOIN_ON / P12 / COMPARISON | `tm.CFCIF` : `CHAR(19)` | `inf.CIFNo` : `VARCHAR(19)` | `tm.CFCIF` → `VARCHAR` | `inf.CIFNo = tm.CFCIF` |
| IC-001762 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.ProTRSReportCOD067:10](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ProTRSReportCOD067.sql#L10) | WHERE / P12 / COMPARISON | `CFCIF` : `CHAR(19)` | `RIGHT('0000000000000000000'+@pvcCIF,19)` : `VARCHAR (string function; length unknown)` | `CFCIF` → `VARCHAR` | `CFCIF=RIGHT('0000000000000000000'+@pvcCIF,19)` |
| IC-003707 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:304](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L304) | JOIN_ON / P12 / COMPARISON | `cfmast.CFCIF` : `CHAR(19)` | `mnc._CIFNo` : `VARCHAR(19)` | `cfmast.CFCIF` → `VARCHAR` | `mnc._CIFNo = cfmast.CFCIF` |
| IC-003722 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:391](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L391) | JOIN_ON / P12 / COMPARISON | `cfmast.CFCIF` : `CHAR(19)` | `mnc._CIFNo` : `VARCHAR(19)` | `cfmast.CFCIF` → `VARCHAR` | `mnc._CIFNo = cfmast.CFCIF` |
| IC-003803 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:267](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L267) | JOIN_ON / P12 / COMPARISON | `cfmast.CFCIF` : `CHAR(19)` | `mnc._CIFNo` : `VARCHAR(19)` | `cfmast.CFCIF` → `VARCHAR` | `mnc._CIFNo = cfmast.CFCIF` |
| IC-003808 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:280](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L280) | JOIN_ON / P12 / COMPARISON | `cfmast.CFCIF` : `CHAR(19)` | `mnc._CIFNo` : `VARCHAR(19)` | `cfmast.CFCIF` → `VARCHAR` | `mnc._CIFNo = cfmast.CFCIF` |
| IC-003856 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:235](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L235) | JOIN_ON / P12 / COMPARISON | `cfmast.CFCIF` : `CHAR(19)` | `mnc._CIFNo` : `VARCHAR(19)` | `cfmast.CFCIF` → `VARCHAR` | `mnc._CIFNo = cfmast.CFCIF` |
| IC-003874 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:327](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L327) | JOIN_ON / P12 / COMPARISON | `cfmast.CFCIF` : `CHAR(19)` | `mnc._CIFNo` : `VARCHAR(19)` | `cfmast.CFCIF` → `VARCHAR` | `mnc._CIFNo = cfmast.CFCIF` |
| IC-003877 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:333](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L333) | JOIN_ON / P12 / COMPARISON | `cfmast.CFCIF` : `CHAR(19)` | `mnc._CIFNo` : `VARCHAR(19)` | `cfmast.CFCIF` → `VARCHAR` | `mnc._CIFNo = cfmast.CFCIF` |
| IC-003882 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:341](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L341) | JOIN_ON / P12 / COMPARISON | `cfmast.CFCIF` : `CHAR(19)` | `mnc._CIFNo` : `VARCHAR(19)` | `cfmast.CFCIF` → `VARCHAR` | `mnc._CIFNo = cfmast.CFCIF` |
| IC-003885 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:347](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L347) | JOIN_ON / P12 / COMPARISON | `cfmast.CFCIF` : `CHAR(19)` | `mnc._CIFNo` : `VARCHAR(19)` | `cfmast.CFCIF` → `VARCHAR` | `mnc._CIFNo = cfmast.CFCIF` |
| IC-003888 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:353](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L353) | JOIN_ON / P12 / COMPARISON | `cfmast.CFCIF` : `CHAR(19)` | `mnc._CIFNo` : `VARCHAR(19)` | `cfmast.CFCIF` → `VARCHAR` | `mnc._CIFNo = cfmast.CFCIF` |
| IC-003938 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusiOrchid:213](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L213) | JOIN_ON / P12 / COMPARISON | `cfmast.CFCIF` : `CHAR(19)` | `mnc._CIFNo` : `VARCHAR(19)` | `cfmast.CFCIF` → `VARCHAR` | `mnc._CIFNo = cfmast.CFCIF` |
| IC-003943 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusiOrchid:226](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L226) | JOIN_ON / P12 / COMPARISON | `cfmast.CFCIF` : `CHAR(19)` | `mnc._CIFNo` : `VARCHAR(19)` | `cfmast.CFCIF` → `VARCHAR` | `mnc._CIFNo = cfmast.CFCIF` |
| IC-004127 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSGetJenisNasabahCIF:6](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetJenisNasabahCIF.sql#L6) | WHERE / P12 / COMPARISON | `cf.CFCIF` : `CHAR(19)` | `right(replicate('0',19)+ @pcCIF, 19)` : `VARCHAR (string function; length unknown)` | `cf.CFCIF` → `VARCHAR` | `cf.CFCIF = right(replicate('0',19)+ @pcCIF, 19)` |
| IC-004427 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSJobSyncKITAS:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSJobSyncKITAS.sql#L25) | JOIN_ON / P12 / COMPARISON | `cf.[CFCIF]` : `CHAR(19)` | `tc.[CIFNo]` : `VARCHAR(19)` | `cf.[CFCIF]` → `VARCHAR` | `tc.[CIFNo] = cf.[CFCIF]` |
| IC-004434 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSJobSyncKITAS:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSJobSyncKITAS.sql#L47) | JOIN_ON / P12 / COMPARISON | `cf.[CFCIF]` : `CHAR(19)` | `tc.[CIFNo]` : `VARCHAR(19)` | `cf.[CFCIF]` → `VARCHAR` | `tc.[CIFNo] = cf.[CFCIF]` |
| IC-005561 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateMasterNasabahForMurex:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMasterNasabahForMurex.sql#L36) | WHERE / P12 / COMPARISON | `CFCIF` : `CHAR(19)` | `@cCIFNo` : `VARCHAR(19)` | `CFCIF` → `VARCHAR` | `CFCIF = @cCIFNo` |
| IC-005992 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateRiskProfileNasabah:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateRiskProfileNasabah.sql#L18) | WHERE / P12 / COMPARISON | `CFCIF` : `CHAR(19)` | `@pcCIF` : `varchar(19)` | `CFCIF` → `VARCHAR` | `CFCIF = @pcCIF` |
| IC-007848 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSREPOReportBiayaMaterai:308](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L308) | JOIN_ON / P12 / COMPARISON | `cfm.CFCIF` : `CHAR(19)` | `tmp.CIFNo` : `VARCHAR(19)` | `cfm.CFCIF` → `VARCHAR` | `tmp.CIFNo = cfm.CFCIF` |
| IC-007852 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSREPOReportBiayaMaterai:326](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L326) | JOIN_ON / P12 / COMPARISON | `cfm.CFCIF` : `CHAR(19)` | `tmp.CIFNo` : `VARCHAR(19)` | `cfm.CFCIF` → `VARCHAR` | `tmp.CIFNo = cfm.CFCIF` |
| IC-007875 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahObligasi:441](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahObligasi.sql#L441) | JOIN_ON / P12 / COMPARISON | `v.CFCIF` : `CHAR(19)` | `rs.CIF` : `VARCHAR(19)` | `v.CFCIF` → `VARCHAR` | `rs.CIF = v.CFCIF` |
| IC-007900 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahPerubahanObligasi2:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql#L81) | JOIN_ON / P12 / COMPARISON | `c.CFCIF` : `CHAR(19)` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19)` : `VARCHAR (string function; length unknown)` | `c.CFCIF` → `VARCHAR` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19) = c.CFCIF` |
| IC-007901 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahPerubahanObligasi2:93](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql#L93) | JOIN_ON / P12 / COMPARISON | `c.CFCIF` : `CHAR(19)` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19)` : `VARCHAR (string function; length unknown)` | `c.CFCIF` → `VARCHAR` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19) = c.CFCIF` |
| IC-007902 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahPerubahanObligasi2:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql#L105) | JOIN_ON / P12 / COMPARISON | `c.CFCIF` : `CHAR(19)` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19)` : `VARCHAR (string function; length unknown)` | `c.CFCIF` → `VARCHAR` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19) = c.CFCIF` |
| IC-007903 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahPerubahanObligasi2:117](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql#L117) | JOIN_ON / P12 / COMPARISON | `c.CFCIF` : `CHAR(19)` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19)` : `VARCHAR (string function; length unknown)` | `c.CFCIF` → `VARCHAR` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19) = c.CFCIF` |
| IC-007904 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahPerubahanObligasi2:129](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql#L129) | JOIN_ON / P12 / COMPARISON | `c.CFCIF` : `CHAR(19)` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19)` : `VARCHAR (string function; length unknown)` | `c.CFCIF` → `VARCHAR` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19) = c.CFCIF` |
| IC-007905 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahPerubahanObligasi2:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql#L141) | JOIN_ON / P12 / COMPARISON | `c.CFCIF` : `CHAR(19)` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19)` : `VARCHAR (string function; length unknown)` | `c.CFCIF` → `VARCHAR` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19) = c.CFCIF` |
| IC-007906 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahPerubahanObligasi2:153](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql#L153) | JOIN_ON / P12 / COMPARISON | `c.CFCIF` : `CHAR(19)` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19)` : `VARCHAR (string function; length unknown)` | `c.CFCIF` → `VARCHAR` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19) = c.CFCIF` |
| IC-007907 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahPerubahanObligasi2:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql#L165) | JOIN_ON / P12 / COMPARISON | `c.CFCIF` : `CHAR(19)` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19)` : `VARCHAR (string function; length unknown)` | `c.CFCIF` → `VARCHAR` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19) = c.CFCIF` |
| IC-007908 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahPerubahanObligasi2:177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql#L177) | JOIN_ON / P12 / COMPARISON | `c.CFCIF` : `CHAR(19)` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19)` : `VARCHAR (string function; length unknown)` | `c.CFCIF` → `VARCHAR` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19) = c.CFCIF` |
| IC-007909 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahPerubahanObligasi2:189](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql#L189) | JOIN_ON / P12 / COMPARISON | `c.CFCIF` : `CHAR(19)` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19)` : `VARCHAR (string function; length unknown)` | `c.CFCIF` → `VARCHAR` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19) = c.CFCIF` |
| IC-007910 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahPerubahanObligasi2:201](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql#L201) | JOIN_ON / P12 / COMPARISON | `c.CFCIF` : `CHAR(19)` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19)` : `VARCHAR (string function; length unknown)` | `c.CFCIF` → `VARCHAR` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19) = c.CFCIF` |
| IC-007911 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahPerubahanObligasi2:213](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql#L213) | JOIN_ON / P12 / COMPARISON | `c.CFCIF` : `CHAR(19)` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19)` : `VARCHAR (string function; length unknown)` | `c.CFCIF` → `VARCHAR` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19) = c.CFCIF` |
| IC-007912 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahPerubahanObligasi2:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql#L225) | JOIN_ON / P12 / COMPARISON | `c.CFCIF` : `CHAR(19)` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19)` : `VARCHAR (string function; length unknown)` | `c.CFCIF` → `VARCHAR` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19) = c.CFCIF` |
| IC-007913 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahPerubahanObligasi2:237](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql#L237) | JOIN_ON / P12 / COMPARISON | `c.CFCIF` : `CHAR(19)` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19)` : `VARCHAR (string function; length unknown)` | `c.CFCIF` → `VARCHAR` | `right('0000000000000000000'+convert(varchar(19), a.CIFNumber), 19) = c.CFCIF` |
| IC-008479 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTemplateSID:139](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSID.sql#L139) | JOIN_ON / P12 / COMPARISON | `cf.CFCIF` : `CHAR(19)` | `tmp.CIFNo` : `VARCHAR(19)` | `cf.CFCIF` → `VARCHAR` | `tmp.CIFNo = cf.CFCIF` |
| IC-008518 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTrxCorporateBOND:202](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L202) | JOIN_ON / P12 / COMPARISON | `cfmast.CFCIF` : `CHAR(19)` | `tr.CIFNo` : `varchar(20)` | `cfmast.CFCIF` → `VARCHAR` | `cfmast.CFCIF = tr.CIFNo` |
| IC-008916 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailCalculateFee2:184](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee2.sql#L184) | JOIN_ON / P12 / COMPARISON | `cf.CFCIF` : `CHAR(19)` | `tmp.[No CIF]` : `VARCHAR(19)` | `cf.CFCIF` → `VARCHAR` | `tmp.[No CIF] = cf.CFCIF` |
| IC-009178 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:404](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L404) | JOIN_ON / P12 / COMPARISON | `cf.CFCIF` : `CHAR(19)` | `t.CIFNo` : `VARCHAR(19)` | `cf.CFCIF` → `VARCHAR` | `t.CIFNo = cf.CFCIF` |
| IC-009345 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:302](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L302) | JOIN_ON / P12 / COMPARISON | `cf.CFCIF` : `CHAR(19)` | `t.CIFNo` : `VARCHAR(19)` | `cf.CFCIF` → `VARCHAR` | `t.CIFNo = cf.CFCIF` |
| IC-009591 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptCustomerPortfolio:88](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerPortfolio.sql#L88) | JOIN_ON / P12 / COMPARISON | `CFMAST.CFCIF` : `CHAR(19)` | `tctm.CIFNo` : `VARCHAR(19)` | `CFMAST.CFCIF` → `VARCHAR` | `CFMAST.CFCIF = tctm.CIFNo` |
| IC-010307 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptMonitoringDocAcc:186](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptMonitoringDocAcc.sql#L186) | JOIN_ON / P12 / COMPARISON | `cf.[CFCIF]` : `CHAR(19)` | `right('0000000000000000000' + tr.[NoCIF], 19)` : `VARCHAR (string function; length unknown)` | `cf.[CFCIF]` → `VARCHAR` | `right('0000000000000000000' + tr.[NoCIF], 19) = cf.[CFCIF]` |
| IC-012402 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncMasterIdAndFunc:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterIdAndFunc.sql#L22) | JOIN_ON / P12 / COMPARISON | `mast.CFCIF` : `CHAR(19)` | `tmp.CIFNo` : `VARCHAR(19)` | `mast.CFCIF` → `VARCHAR` | `tmp.CIFNo = mast.CFCIF` |
| IC-012415 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncMasterRiskProfileData:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterRiskProfileData.sql#L21) | JOIN_ON / P12 / COMPARISON | `CFMAST.CFCIF` : `CHAR(19)` | `tc.CIFNo` : `VARCHAR(19)` | `CFMAST.CFCIF` → `VARCHAR` | `tc.CIFNo = CFMAST.CFCIF` |
| IC-012719 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSValidateSecAccNo:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateSecAccNo.sql#L30) | WHERE / P12 / COMPARISON | `CFCIF` : `CHAR(19)` | `right('0000000000000000000' + convert(varchar,@nCIFNo), 19)` : `VARCHAR (string function; length unknown)` | `CFCIF` → `VARCHAR` | `CFCIF = right('0000000000000000000' + convert(varchar,@nCIFNo), 19)` |
| IC-013559 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListTreasuryCustomer_TM_Original:236](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListTreasuryCustomer_TM_Original.sql#L236) | JOIN_ON / P12 / COMPARISON | `cfm.CFCIF` : `CHAR(19)` | `tctm.CIFNo` : `VARCHAR(19)` | `cfm.CFCIF` → `VARCHAR` | `tctm.CIFNo = cfm.CFCIF` |
| IC-013565 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListTreasuryCustomer_TM_Original:402](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListTreasuryCustomer_TM_Original.sql#L402) | JOIN_ON / P12 / COMPARISON | `cfm.CFCIF` : `CHAR(19)` | `tctm.CIFNo` : `VARCHAR(19)` | `cfm.CFCIF` → `VARCHAR` | `tctm.CIFNo = cfm.CFCIF` |
| IC-013966 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_getExpiredKTPNasabah:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_getExpiredKTPNasabah.sql#L25) | WHERE / P12 / COMPARISON | `CFCIF` : `CHAR(19)` | `@CIFNo` : `VARCHAR(19)` | `CFCIF` → `VARCHAR` | `CFCIF = @CIFNo` |

<a id="kolom-cfbird"></a>
## CFBIRD

**Tipe dasar:** `CHAR(7)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFMAST.sql#L46). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 7 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_ARGUMENT, CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002290 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L108) | JOIN_ON / P01 / COMPARISON | `dt.CFBIRD` : `CHAR(7)` | `tmpCus.TglLahirJul` : `INT` | `dt.CFBIRD` → `INT` | `dt.CFBIRD = tmpCus.TglLahirJul` |
| IC-001152 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.PTOSProsesDataJobNasabah:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L98) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `CFBIRD` : `CHAR(7)` | `retail.dbo.fnjuliantodatetime.@nDate` : `int` | `CFBIRD` → `int` | `dbo.fnJulianToDatetime(CFBIRD)` |
| IC-003700 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSDataNasabahRetailIndividu:291](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L291) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `cfmast.CFBIRD` : `CHAR(7)` | `retail.dbo.fnjuliantodatetime.@nDate` : `int` | `cfmast.CFBIRD` → `int` | `dbo.fnJulianToDatetime(cfmast.CFBIRD)` |
| IC-003793 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSDataNasabahRetailIndividuOrchid:249](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L249) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `cfmast.CFBIRD` : `CHAR(7)` | `retail.dbo.fnjuliantodatetime.@nDate` : `int` | `cfmast.CFBIRD` → `int` | `dbo.fnJulianToDatetime(cfmast.CFBIRD)` |
| IC-003854 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSDataNasabahRetailInstitusi:229](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L229) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `cfmast.CFBIRD` : `CHAR(7)` | `retail.dbo.fnjuliantodatetime.@nDate` : `int` | `cfmast.CFBIRD` → `int` | `dbo.fnJulianToDatetime(cfmast.CFBIRD)` |
| IC-003937 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSDataNasabahRetailInstitusiOrchid:210](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L210) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `cfmast.CFBIRD` : `CHAR(7)` | `retail.dbo.fnjuliantodatetime.@nDate` : `int` | `cfmast.CFBIRD` → `int` | `dbo.fnJulianToDatetime(cfmast.CFBIRD)` |
| IC-005815 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUpload:260](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUpload.sql#L260) | JOIN_ON / P12 / COMPARISON | `dt.CFBIRD` : `CHAR(7)` | `tmpCus.TglLahirJul` : `varchar(8)` | `dt.CFBIRD` → `VARCHAR` | `dt.CFBIRD = tmpCus.TglLahirJul` |

<a id="kolom-cfssno"></a>
## CFSSNO

**Tipe dasar:** `CHAR(40)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFMAST.sql#L43). **Existing leading index:** idx_cfmast_cfssno.

**Frekuensi source:** 33 lokasi / 12 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.
- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001146 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobNasabah:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L94) | JOIN_ON / P02 / COMPARISON | `v.CFSSNO` : `CHAR(40)` | `tmp.[Nomor KTP]` : `NVARCHAR(255)` | `v.CFSSNO` → `NVARCHAR` | `tmp.[Nomor KTP] = v.CFSSNO` |
| IC-003880 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailInstitusi:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L337) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfmast.CFSSNO` : `CHAR(40)` | `SKD` : `VARCHAR(120)` | `cfmast.CFSSNO` → `VARCHAR(120)` | `SKD = cfmast.CFSSNO` |
| IC-003883 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailInstitusi:344](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L344) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfmast.CFSSNO` : `CHAR(40)` | `ARTICLE_OF_ASSC` : `VARCHAR(120)` | `cfmast.CFSSNO` → `VARCHAR(120)` | `ARTICLE_OF_ASSC = cfmast.CFSSNO` |
| IC-003886 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailInstitusi:350](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L350) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfmast.CFSSNO` : `CHAR(40)` | `BUSINESS_REG_CERT` : `VARCHAR(120)` | `cfmast.CFSSNO` → `VARCHAR(120)` | `BUSINESS_REG_CERT = cfmast.CFSSNO` |
| IC-004425 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSJobSyncKITAS:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSJobSyncKITAS.sql#L21) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cf.[CFSSNO]` : `CHAR(40)` | `tc.[KITASNo]` : `VARCHAR(40)` | `cf.[CFSSNO]` → `VARCHAR(40)` | `tc.[KITASNo] = cf.[CFSSNO]` |
| IC-004760 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSOBLExpiredKitasRPT:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOBLExpiredKitasRPT.sql#L34) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CM.CFSSNO` : `CHAR(40)` | `@tmpCif.CFSSNO` : `VARCHAR(40)` | `CM.CFSSNO` → `VARCHAR(40)` | `@tmpCif.CFSSNO <- CM.CFSSNO` |
| IC-007847 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSREPOReportBiayaMaterai:305](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L305) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfm.CFSSNO` : `CHAR(40)` | `NomorIdentitasNIK` : `VARCHAR(50)` | `cfm.CFSSNO` → `VARCHAR(50)` | `NomorIdentitasNIK = cfm.CFSSNO` |
| IC-007851 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSREPOReportBiayaMaterai:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L323) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfm.CFSSNO` : `CHAR(40)` | `NomorIdentitas` : `VARCHAR(50)` | `cfm.CFSSNO` → `VARCHAR(50)` | `NomorIdentitas = cfm.CFSSNO` |
| IC-009161 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailSyncDataCore:368](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L368) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cf.CFSSNO` : `CHAR(40)` | `newNoIdentitas` : `VARCHAR(40)` | `cf.CFSSNO` → `VARCHAR(40)` | `newNoIdentitas = cf.CFSSNO` |
| IC-009328 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailSyncDataCoreWorkingHours:266](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L266) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cf.CFSSNO` : `CHAR(40)` | `newNoIdentitas` : `VARCHAR(40)` | `cf.CFSSNO` → `VARCHAR(40)` | `newNoIdentitas = cf.CFSSNO` |
| IC-002291 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:109](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L109) | JOIN_ON / P12 / COMPARISON | `dt.CFSSNO` : `CHAR(40)` | `tmpCus.NoIdentitas` : `VARCHAR(40)` | `dt.CFSSNO` → `VARCHAR` | `dt.CFSSNO = tmpCus.NoIdentitas` |
| IC-003685 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:263](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L263) | EXPRESSION / P07 / CASE_BRANCH | `cfmast.CFSSNO` : `CHAR(40)` | `result type` : `VARCHAR (combined result; facets unknown)` | `cfmast.CFSSNO` → `VARCHAR` | `CASE WHEN cfmast.CFSSCD = 'KTP' OR cfmast.CFSSCD = 'KTPS' THEN cfmast.CFSSNO ELSE '' END` |
| IC-003688 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSDataNasabahRetailIndividu:271](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L271) | CASE / P12 / COMPARISON | `ISNULL(cfmast.CFSSNO, '')` : `CHAR(40)` | `''` : `varchar` | `ISNULL(cfmast.CFSSNO, '')` → `VARCHAR` | `ISNULL(cfmast.CFSSNO, '') <> ''` |
| IC-003694 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:278](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L278) | EXPRESSION / P07 / CASE_BRANCH | `cfmast.CFSSNO` : `CHAR(40)` | `result type` : `VARCHAR (combined result; facets unknown)` | `cfmast.CFSSNO` → `VARCHAR` | `CASE WHEN cfmast.CFSSCD = 'PP' THEN cfmast.CFSSNO ELSE '' END` |
| IC-003696 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSDataNasabahRetailIndividu:286](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L286) | CASE / P12 / COMPARISON | `ISNULL(cfmast.CFSSNO,'')` : `CHAR(40)` | `''` : `varchar` | `ISNULL(cfmast.CFSSNO,'')` → `VARCHAR` | `ISNULL(cfmast.CFSSNO,'') <> ''` |
| IC-003786 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:237](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L237) | EXPRESSION / P07 / CASE_BRANCH | `cfmast.CFSSNO` : `CHAR(40)` | `result type` : `VARCHAR (combined result; facets unknown)` | `cfmast.CFSSNO` → `VARCHAR` | `CASE WHEN mnc.NATIONALITYOBL <> '000' AND cfmast.CFSSCD = 'PP' THEN cfmast.CFSSNO ELSE '' END` |
| IC-003789 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSDataNasabahRetailIndividuOrchid:244](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L244) | CASE / P12 / COMPARISON | `ISNULL(cfmast.CFSSNO,'')` : `CHAR(40)` | `''` : `varchar` | `ISNULL(cfmast.CFSSNO,'')` → `VARCHAR` | `ISNULL(cfmast.CFSSNO,'') <> ''` |
| IC-003876 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:333](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L333) | JOIN_ON / P12 / COMPARISON | `cfmast.CFSSNO` : `CHAR(40)` | `mnc.NPWP` : `VARCHAR(90)` | `cfmast.CFSSNO` → `VARCHAR` | `mnc.NPWP = cfmast.CFSSNO` |
| IC-004763 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSOBLExpiredKitasRPT:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOBLExpiredKitasRPT.sql#L49) | EXPRESSION / P07 / COALESCE_BRANCH | `CM.CFSSNO` : `CHAR(40)` | `result type` : `VARCHAR (combined result; facets unknown)` | `CM.CFSSNO` → `VARCHAR` | `COALESCE(CM.CFSSNO,'')` |
| IC-005816 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUpload:261](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUpload.sql#L261) | JOIN_ON / P12 / COMPARISON | `dt.CFSSNO` : `CHAR(40)` | `tmpCus.NoIdentitas` : `varchar(40)` | `dt.CFSSNO` → `VARCHAR` | `dt.CFSSNO = tmpCus.NoIdentitas` |
| IC-008494 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTrxCorporateBOND:150](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L150) | EXPRESSION / P07 / CASE_BRANCH | `cfmast.CFSSNO` : `CHAR(40)` | `result type` : `VARCHAR (combined result; facets unknown)` | `cfmast.CFSSNO` → `VARCHAR` | `case cfmast.CFSSCD when 'KTP' then cfmast.CFSSNO when 'KTPS' then cfmast.CFSSNO else '' end` |
| IC-008495 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTrxCorporateBOND:150](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L150) | EXPRESSION / P07 / CASE_BRANCH | `cfmast.CFSSNO` : `CHAR(40)` | `result type` : `VARCHAR (combined result; facets unknown)` | `cfmast.CFSSNO` → `VARCHAR` | `case cfmast.CFSSCD when 'KTP' then cfmast.CFSSNO when 'KTPS' then cfmast.CFSSNO else '' end` |
| IC-008498 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTrxCorporateBOND:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L156) | EXPRESSION / P07 / CASE_BRANCH | `cfmast.CFSSNO` : `CHAR(40)` | `result type` : `VARCHAR (combined result; facets unknown)` | `cfmast.CFSSNO` → `VARCHAR` | `case cfmast.CFSSCD when 'PP' then cfmast.CFSSNO else '' end` |
| IC-008510 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSReportTrxCorporateBOND:186](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L186) | CASE / P12 / COMPARISON | `ISNULL(cfmast.CFSSNO,'')` : `CHAR(40)` | `''` : `varchar` | `ISNULL(cfmast.CFSSNO,'')` → `VARCHAR` | `ISNULL(cfmast.CFSSNO,'') <> ''` |
| IC-008514 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSReportTrxCorporateBOND:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L193) | CASE / P12 / COMPARISON | `ISNULL(cfmast.CFSSNO,'')` : `CHAR(40)` | `''` : `varchar` | `ISNULL(cfmast.CFSSNO,'')` → `VARCHAR` | `ISNULL(cfmast.CFSSNO,'') <> ''` |
| IC-009177 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:388](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L388) | CASE / P12 / COMPARISON | `cf.CFSSNO` : `CHAR(40)` | `t.NoIdentitas` : `VARCHAR(40)` | `cf.CFSSNO` → `VARCHAR` | `t.NoIdentitas <> cf.CFSSNO` |
| IC-009344 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:286](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L286) | CASE / P12 / COMPARISON | `cf.CFSSNO` : `CHAR(40)` | `t.NoIdentitas` : `VARCHAR(40)` | `cf.CFSSNO` → `VARCHAR` | `t.NoIdentitas <> cf.CFSSNO` |
| IC-003692 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSDataNasabahRetailIndividu:271](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L271) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cfmast.CFSSNO` : `CHAR(40)` | `''` : `varchar` | `''` → `CHAR(40)` | `ISNULL(cfmast.CFSSNO, '')` |
| IC-003699 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSDataNasabahRetailIndividu:286](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L286) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cfmast.CFSSNO` : `CHAR(40)` | `''` : `varchar` | `''` → `CHAR(40)` | `ISNULL(cfmast.CFSSNO,'')` |
| IC-003792 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSDataNasabahRetailIndividuOrchid:244](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L244) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cfmast.CFSSNO` : `CHAR(40)` | `''` : `varchar` | `''` → `CHAR(40)` | `ISNULL(cfmast.CFSSNO,'')` |
| IC-004424 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSJobSyncKITAS:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSJobSyncKITAS.sql#L16) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cf.[CFSSNO]` : `CHAR(40)` | `''` : `varchar` | `''` → `CHAR(40)` | `isnull(cf.[CFSSNO], '')` |
| IC-008512 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSReportTrxCorporateBOND:186](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L186) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cfmast.CFSSNO` : `CHAR(40)` | `''` : `varchar` | `''` → `CHAR(40)` | `ISNULL(cfmast.CFSSNO,'')` |
| IC-008517 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSReportTrxCorporateBOND:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L193) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cfmast.CFSSNO` : `CHAR(40)` | `''` : `varchar` | `''` → `CHAR(40)` | `ISNULL(cfmast.CFSSNO,'')` |

<a id="kolom-cfbrnn"></a>
## CFBRNN

**Tipe dasar:** `CHAR(5)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFMAST.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009153 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailSyncDataCore:347](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L347) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cf.CFBRNN` : `CHAR(5)` | `newBranchCode` : `VARCHAR(5)` | `cf.CFBRNN` → `VARCHAR(5)` | `newBranchCode = cf.CFBRNN` |
| IC-009320 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailSyncDataCoreWorkingHours:245](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L245) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cf.CFBRNN` : `CHAR(5)` | `newBranchCode` : `VARCHAR(5)` | `cf.CFBRNN` → `VARCHAR(5)` | `newBranchCode = cf.CFBRNN` |
| IC-009164 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:380](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L380) | CASE / P12 / COMPARISON | `cf.CFBRNN` : `CHAR(5)` | `t.BranchCode` : `VARCHAR(5)` | `cf.CFBRNN` → `VARCHAR` | `t.BranchCode <> cf.CFBRNN` |
| IC-009331 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:278](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L278) | CASE / P12 / COMPARISON | `cf.CFBRNN` : `CHAR(5)` | `t.BranchCode` : `VARCHAR(5)` | `cf.CFBRNN` → `VARCHAR` | `t.BranchCode <> cf.CFBRNN` |

<a id="kolom-cfcif"></a>
## CFCIF#

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFMAST.sql#L7). **Existing leading index:** idx_cfmast_cfcif#.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006537 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateWarganegaraNasabah:7](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateWarganegaraNasabah.sql#L7) | WHERE / P05 / BINARY_Add | `cf.CFCIF#` : `BIGINT` | `replicate('0',19)` : `VARCHAR (expression; length unknown)` | `replicate('0',19)` → `BIGINT` | `replicate('0',19)+ cf.CFCIF#` |
| IC-012417 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSyncMasterRiskProfileData:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterRiskProfileData.sql#L24) | JOIN_ON / P01 / COMPARISON | `CFMAST.CFCIF#` : `BIGINT` | `so.CIFNumber` : `VARCHAR(19)` | `so.CIFNumber` → `BIGINT` | `CFMAST.CFCIF# = so.CIFNumber` |

<a id="kolom-cfcitz"></a>
## CFCITZ

**Tipe dasar:** `CHAR(3)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFMAST.sql#L34). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 10 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009160 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailSyncDataCore:367](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L367) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cf.CFCITZ` : `CHAR(3)` | `newCitizenship` : `VARCHAR(20)` | `cf.CFCITZ` → `VARCHAR(20)` | `newCitizenship = cf.CFCITZ` |
| IC-009327 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailSyncDataCoreWorkingHours:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L265) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cf.CFCITZ` : `CHAR(3)` | `newCitizenship` : `VARCHAR(20)` | `cf.CFCITZ` → `VARCHAR(20)` | `newCitizenship = cf.CFCITZ` |
| IC-004429 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSJobSyncKITAS:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSJobSyncKITAS.sql#L26) | WHERE / P12 / COMPARISON | `cf.[CFCITZ]` : `CHAR(3)` | `'000'` : `varchar` | `cf.[CFCITZ]` → `VARCHAR` | `cf.[CFCITZ] <> '000'` |
| IC-004435 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSJobSyncKITAS:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSJobSyncKITAS.sql#L48) | JOIN_ON / P12 / COMPARISON | `cf.[CFCITZ]` : `CHAR(3)` | `'000'` : `varchar` | `cf.[CFCITZ]` → `VARCHAR` | `cf.[CFCITZ] <> '000'` |
| IC-006536 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateWarganegaraNasabah:6](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateWarganegaraNasabah.sql#L6) | JOIN_ON / P12 / COMPARISON | `cf.CFCITZ` : `CHAR(3)` | `jh.JHCCOC` : `VARCHAR(3)` | `cf.CFCITZ` → `VARCHAR` | `jh.JHCCOC = cf.CFCITZ` |
| IC-007876 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahObligasi:443](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahObligasi.sql#L443) | JOIN_ON / P12 / COMPARISON | `v.CFCITZ` : `CHAR(3)` | `jc.JHCCOC` : `VARCHAR(3)` | `v.CFCITZ` → `VARCHAR` | `v.CFCITZ = jc.JHCCOC` |
| IC-009175 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRetailSyncDataCore:387](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L387) | CASE / P12 / COMPARISON | `ISNULL(cf.CFCITZ,'')` : `CHAR(3)` | `ISNULL(t.Citizenship,'')` : `VARCHAR(20)` | `ISNULL(cf.CFCITZ,'')` → `VARCHAR` | `ISNULL(t.Citizenship,'') <> ISNULL(cf.CFCITZ,'')` |
| IC-009342 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRetailSyncDataCoreWorkingHours:285](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L285) | CASE / P12 / COMPARISON | `ISNULL(cf.CFCITZ,'')` : `CHAR(3)` | `ISNULL(t.Citizenship,'')` : `VARCHAR(20)` | `ISNULL(cf.CFCITZ,'')` → `VARCHAR` | `ISNULL(t.Citizenship,'') <> ISNULL(cf.CFCITZ,'')` |
| IC-009176 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRetailSyncDataCore:387](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L387) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cf.CFCITZ` : `CHAR(3)` | `''` : `varchar` | `''` → `CHAR(3)` | `ISNULL(cf.CFCITZ,'')` |
| IC-009343 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRetailSyncDataCoreWorkingHours:285](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L285) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cf.CFCITZ` : `CHAR(3)` | `''` : `varchar` | `''` → `CHAR(3)` | `ISNULL(cf.CFCITZ,'')` |

<a id="kolom-cfclas"></a>
## CFCLAS

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFMAST.sql#L49). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003855 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailInstitusi:231](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L231) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfmast.CFCLAS` : `CHAR(1)` | `TAX_ID_CORP` : `VARCHAR(4)` | `cfmast.CFCLAS` → `VARCHAR(4)` | `TAX_ID_CORP = cfmast.CFCLAS` |
| IC-005560 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateMasterNasabahForMurex:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMasterNasabahForMurex.sql#L34) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFCLAS` : `CHAR(1)` | `@cCFCLAS` : `varchar(1)` | `CFCLAS` → `varchar(1)` | `@cCFCLAS = CFCLAS` |
| IC-003809 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:282](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L282) | JOIN_ON / P12 / COMPARISON | `cfmast.CFCLAS` : `CHAR(1)` | `ti.CFCLAS` : `VARCHAR(1)` | `cfmast.CFCLAS` → `VARCHAR` | `cfmast.CFCLAS = ti.CFCLAS` |
| IC-003857 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:237](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L237) | JOIN_ON / P12 / COMPARISON | `cfmast.CFCLAS` : `CHAR(1)` | `ti.CFCLAS` : `VARCHAR(1)` | `cfmast.CFCLAS` → `VARCHAR` | `cfmast.CFCLAS = ti.CFCLAS` |
| IC-003944 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusiOrchid:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L228) | JOIN_ON / P12 / COMPARISON | `cfmast.CFCLAS` : `CHAR(1)` | `ti.CFCLAS` : `VARCHAR(1)` | `cfmast.CFCLAS` → `VARCHAR` | `cfmast.CFCLAS = ti.CFCLAS` |
| IC-008480 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTemplateSID:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSID.sql#L141) | JOIN_ON / P12 / COMPARISON | `cf.CFCLAS` : `CHAR(1)` | `ti.CFCLAS` : `VARCHAR(1)` | `cf.CFCLAS` → `VARCHAR` | `cf.CFCLAS = ti.CFCLAS` |
| IC-008508 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTrxCorporateBOND:178](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L178) | CASE / P12 / COMPARISON | `cfmast.CFCLAS` : `CHAR(1)` | `'A'` : `varchar` | `cfmast.CFCLAS` → `VARCHAR` | `cfmast.CFCLAS = 'A'` |
| IC-008509 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTrxCorporateBOND:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L180) | CASE / P12 / COMPARISON | `cfmast.CFCLAS` : `CHAR(1)` | `'A'` : `varchar` | `cfmast.CFCLAS` → `VARCHAR` | `cfmast.CFCLAS <> 'A'` |

<a id="kolom-cfcoun"></a>
## CFCOUN

**Tipe dasar:** `CHAR(3)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFMAST.sql#L33). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003720 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailIndividu:387](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L387) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfmast.CFCOUN` : `CHAR(3)` | `ID_COUNTRY` : `VARCHAR(34)` | `cfmast.CFCOUN` → `VARCHAR(34)` | `ID_COUNTRY = cfmast.CFCOUN` |
| IC-003721 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailIndividu:388](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L388) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfmast.CFCOUN` : `CHAR(3)` | `CORR_COUNTRY` : `VARCHAR(34)` | `cfmast.CFCOUN` → `VARCHAR(34)` | `CORR_COUNTRY = cfmast.CFCOUN` |
| IC-003873 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailInstitusi:324](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L324) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfmast.CFCOUN` : `CHAR(3)` | `CORR_COUNTRY` : `VARCHAR(34)` | `cfmast.CFCOUN` → `VARCHAR(34)` | `CORR_COUNTRY = cfmast.CFCOUN` |

<a id="kolom-cfiex6"></a>
## CFIEX6

**Tipe dasar:** `CHAR(8)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFMAST.sql#L71). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004426 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSJobSyncKITAS:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSJobSyncKITAS.sql#L22) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cf.[CFIEX6]` : `CHAR(8)` | `tc.[KITASExpDate]` : `DATETIME` | `cf.[CFIEX6]` → `DATETIME` | `tc.[KITASExpDate] = cf.[CFIEX6]` |

<a id="kolom-cfiex7"></a>
## CFIEX7

**Tipe dasar:** `CHAR(7)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFMAST.sql#L72). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 10 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT, CONVERTED_COLUMN_EXPRESSION, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004758 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSOBLExpiredKitasRPT:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOBLExpiredKitasRPT.sql#L34) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CM.CFIEX7` : `CHAR(7)` | `@tmpCif.CFIEX7` : `VARCHAR(7)` | `CM.CFIEX7` → `VARCHAR(7)` | `@tmpCif.CFIEX7 <- CM.CFIEX7` |
| IC-012399 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSSyncMasterIdAndFunc:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterIdAndFunc.sql#L16) | CASE / P04 / COMPARISON | `isnull(mast.CFIEX7,0)` : `CHAR(7)` | `0` : `int` | `isnull(mast.CFIEX7,0)` → `INT` | `isnull(mast.CFIEX7,0) != 0` |
| IC-012401 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSSyncMasterIdAndFunc:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterIdAndFunc.sql#L16) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `CFIEX7` : `CHAR(7)` | `retail.dbo.fnjuliantodatetime.@nDate` : `int` | `CFIEX7` → `int` | `dbo.fnJulianToDatetime(CFIEX7)` |
| IC-012408 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSSyncMasterIdAndFunc:32](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterIdAndFunc.sql#L32) | CASE / P04 / COMPARISON | `isnull(mast.CFIEX7,0)` : `CHAR(7)` | `0` : `int` | `isnull(mast.CFIEX7,0)` → `INT` | `isnull(mast.CFIEX7,0) != 0` |
| IC-012410 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSSyncMasterIdAndFunc:32](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterIdAndFunc.sql#L32) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `CFIEX7` : `CHAR(7)` | `retail.dbo.fnjuliantodatetime.@nDate` : `int` | `CFIEX7` → `int` | `dbo.fnJulianToDatetime(CFIEX7)` |
| IC-013965 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.trs_getExpiredKTPNasabah:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_getExpiredKTPNasabah.sql#L22) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `CFIEX7` : `CHAR(7)` | `retail.dbo.fnjuliantodatetime.@nDate` : `int` | `CFIEX7` → `int` | `dbo.fnJulianToDatetime(CFIEX7)` |
| IC-013963 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_getExpiredKTPNasabah:20](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_getExpiredKTPNasabah.sql#L20) | CASE / P12 / COMPARISON | `ISNULL(CFIEX7, '0')` : `CHAR(7)` | `'0'` : `varchar` | `ISNULL(CFIEX7, '0')` → `VARCHAR` | `ISNULL(CFIEX7, '0') = '0'` |
| IC-012400 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSyncMasterIdAndFunc:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterIdAndFunc.sql#L16) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mast.CFIEX7` : `CHAR(7)` | `0` : `int` | `0` → `CHAR(7)` | `isnull(mast.CFIEX7,0)` |
| IC-012409 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSyncMasterIdAndFunc:32](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterIdAndFunc.sql#L32) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mast.CFIEX7` : `CHAR(7)` | `0` : `int` | `0` → `CHAR(7)` | `isnull(mast.CFIEX7,0)` |
| IC-013964 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_getExpiredKTPNasabah:20](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_getExpiredKTPNasabah.sql#L20) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `CFIEX7` : `CHAR(7)` | `'0'` : `varchar` | `'0'` → `CHAR(7)` | `ISNULL(CFIEX7, '0')` |

<a id="kolom-cfisd6"></a>
## CFISD6

**Tipe dasar:** `CHAR(8)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFMAST.sql#L92). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN_EXPRESSION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003878 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSDataNasabahRetailInstitusi:334](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L334) | WHERE / P04 / COMPARISON | `ISNULL (cfmast.CFISD6, 0)` : `CHAR(8)` | `0` : `int` | `ISNULL (cfmast.CFISD6, 0)` → `INT` | `ISNULL (cfmast.CFISD6, 0) <> 0` |
| IC-003879 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSDataNasabahRetailInstitusi:334](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L334) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cfmast.CFISD6` : `CHAR(8)` | `0` : `int` | `0` → `CHAR(8)` | `ISNULL (cfmast.CFISD6, 0)` |

<a id="kolom-cfisd7"></a>
## CFISD7

**Tipe dasar:** `CHAR(7)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFMAST.sql#L93). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004759 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSOBLExpiredKitasRPT:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOBLExpiredKitasRPT.sql#L34) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CM.CFISD7` : `CHAR(7)` | `@tmpCif.CFISD7` : `VARCHAR(7)` | `CM.CFISD7` → `VARCHAR(7)` | `@tmpCif.CFISD7 <- CM.CFISD7` |

<a id="kolom-cfmotn"></a>
## CFMOTN

**Tipe dasar:** `CHAR(40)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFMAST.sql#L90). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003706 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailIndividu:301](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L301) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfmast.CFMOTN` : `CHAR(40)` | `MOTHER_MAIDEN_NAME` : `VARCHAR(200)` | `cfmast.CFMOTN` → `VARCHAR(200)` | `MOTHER_MAIDEN_NAME = cfmast.CFMOTN` |
| IC-003802 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailIndividuOrchid:264](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L264) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfmast.CFMOTN` : `CHAR(40)` | `MOTHER_MAIDEN_NAME` : `VARCHAR(200)` | `cfmast.CFMOTN` → `VARCHAR(200)` | `MOTHER_MAIDEN_NAME = cfmast.CFMOTN` |
| IC-008507 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportTrxCorporateBOND:175](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L175) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfmast.CFMOTN` : `CHAR(40)` | `tr.investorMothersMaidenName` : `varchar(40)` | `cfmast.CFMOTN` → `varchar(40)` | `tr.investorMothersMaidenName = cfmast.CFMOTN` |

<a id="kolom-cfsex"></a>
## CFSEX

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFMAST.sql#L48). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 13 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003795 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:252](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L252) | CASE / P12 / COMPARISON | `cfmast.CFSEX` : `CHAR(1)` | `'M'` : `varchar` | `cfmast.CFSEX` → `VARCHAR` | `cfmast.CFSEX = 'M'` |
| IC-003796 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:253](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L253) | CASE / P12 / COMPARISON | `cfmast.CFSEX` : `CHAR(1)` | `'F'` : `varchar` | `cfmast.CFSEX` → `VARCHAR` | `cfmast.CFSEX = 'F'` |
| IC-008502 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTrxCorporateBOND:164](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L164) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `cfmast.CFSEX` : `CHAR(1)` | `'M'` : `varchar` | `cfmast.CFSEX` → `VARCHAR` | `when 'M' then '1'` |
| IC-009156 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:357](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L357) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `cf.CFSEX` : `CHAR(1)` | `'M'` : `varchar` | `cf.CFSEX` → `VARCHAR` | `WHEN 'M' THEN '1'` |
| IC-009157 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:358](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L358) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `cf.CFSEX` : `CHAR(1)` | `'F'` : `varchar` | `cf.CFSEX` → `VARCHAR` | `WHEN 'F' THEN '2'` |
| IC-009166 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:382](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L382) | CASE / P12 / COMPARISON | `cf.CFSEX` : `CHAR(1)` | `'M'` : `varchar` | `cf.CFSEX` → `VARCHAR` | `cf.CFSEX <> 'M'` |
| IC-009168 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:383](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L383) | CASE / P12 / COMPARISON | `cf.CFSEX` : `CHAR(1)` | `'F'` : `varchar` | `cf.CFSEX` → `VARCHAR` | `cf.CFSEX <> 'F'` |
| IC-009171 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:384](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L384) | CASE / P12 / IN_VALUE_TYPE | `cf.CFSEX` : `CHAR(1)` | `'M'` : `varchar` | `cf.CFSEX` → `VARCHAR` | `cf.CFSEX IN ('M', 'F')` |
| IC-009323 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:255](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L255) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `cf.CFSEX` : `CHAR(1)` | `'M'` : `varchar` | `cf.CFSEX` → `VARCHAR` | `WHEN 'M' THEN '1'` |
| IC-009324 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:256](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L256) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `cf.CFSEX` : `CHAR(1)` | `'F'` : `varchar` | `cf.CFSEX` → `VARCHAR` | `WHEN 'F' THEN '2'` |
| IC-009333 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:280](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L280) | CASE / P12 / COMPARISON | `cf.CFSEX` : `CHAR(1)` | `'M'` : `varchar` | `cf.CFSEX` → `VARCHAR` | `cf.CFSEX <> 'M'` |
| IC-009335 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:281](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L281) | CASE / P12 / COMPARISON | `cf.CFSEX` : `CHAR(1)` | `'F'` : `varchar` | `cf.CFSEX` → `VARCHAR` | `cf.CFSEX <> 'F'` |
| IC-009338 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:282](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L282) | CASE / P12 / IN_VALUE_TYPE | `cf.CFSEX` : `CHAR(1)` | `'M'` : `varchar` | `cf.CFSEX` → `VARCHAR` | `cf.CFSEX IN ('M', 'F')` |

<a id="kolom-cfsnme"></a>
## CFSNME

**Tipe dasar:** `CHAR(20)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFMAST.sql#L9). **Existing leading index:** idx_cfmast_cfsnme.

**Frekuensi source:** 4 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001055 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSPopulateConsentNasabahPTOS:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSPopulateConsentNasabahPTOS.sql#L69) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.CFSNME` : `CHAR(20)` | `#tmpDataNasabah.NamaNasabah` : `VARCHAR(50)` | `tm.CFSNME` → `VARCHAR(50)` | `#tmpDataNasabah.NamaNasabah <- tm.CFSNME` |
| IC-001145 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSProsesDataJobNasabah:92](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L92) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `v.CFSNME` : `CHAR(20)` | `tmp.Name` : `varchar(50)` | `v.CFSNME` → `varchar(50)` | `tmp.Name = v.CFSNME` |
| IC-002289 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.Rvmp_TRSPTOSProcessUpload:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L106) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `dt.CFSNME` : `CHAR(20)` | `tmpCus.Name` : `VARCHAR(95)` | `dt.CFSNME` → `VARCHAR(95)` | `tmpCus.Name = dt.CFSNME` |
| IC-005814 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulatePTOSUpload:257](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUpload.sql#L257) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `dt.CFSNME` : `CHAR(20)` | `tmpCus.Name` : `varchar(95)` | `dt.CFSNME` → `varchar(95)` | `tmpCus.Name = dt.CFSNME` |

<a id="kolom-cfsscd"></a>
## CFSSCD

**Tipe dasar:** `CHAR(4)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFMAST.sql#L44). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 33 lokasi / 12 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001147 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobNasabah:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L94) | JOIN_ON / P12 / IN_VALUE_TYPE | `v.CFSSCD` : `CHAR(4)` | `'KTP'` : `varchar` | `v.CFSSCD` → `VARCHAR` | `v.CFSSCD IN ('KTP', 'KTPS')` |
| IC-002292 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:109](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L109) | JOIN_ON / P12 / IN_VALUE_TYPE | `dt.CFSSCD` : `CHAR(4)` | `'KTP'` : `varchar` | `dt.CFSSCD` → `VARCHAR` | `dt.CFSSCD IN ('KTP', 'KTPS')` |
| IC-003686 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:264](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L264) | CASE / P12 / COMPARISON | `cfmast.CFSSCD` : `CHAR(4)` | `'KTP'` : `varchar` | `cfmast.CFSSCD` → `VARCHAR` | `cfmast.CFSSCD = 'KTP'` |
| IC-003687 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:264](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L264) | CASE / P12 / COMPARISON | `cfmast.CFSSCD` : `CHAR(4)` | `'KTPS'` : `varchar` | `cfmast.CFSSCD` → `VARCHAR` | `cfmast.CFSSCD = 'KTPS'` |
| IC-003689 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:271](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L271) | CASE / P12 / COMPARISON | `cfmast.CFSSCD` : `CHAR(4)` | `'KTP'` : `varchar` | `cfmast.CFSSCD` → `VARCHAR` | `cfmast.CFSSCD = 'KTP'` |
| IC-003690 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:271](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L271) | CASE / P12 / COMPARISON | `cfmast.CFSSCD` : `CHAR(4)` | `'KTPS'` : `varchar` | `cfmast.CFSSCD` → `VARCHAR` | `cfmast.CFSSCD = 'KTPS'` |
| IC-003695 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:279](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L279) | CASE / P12 / COMPARISON | `cfmast.CFSSCD` : `CHAR(4)` | `'PP'` : `varchar` | `cfmast.CFSSCD` → `VARCHAR` | `cfmast.CFSSCD = 'PP'` |
| IC-003697 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:286](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L286) | CASE / P12 / COMPARISON | `cfmast.CFSSCD` : `CHAR(4)` | `'PP'` : `varchar` | `cfmast.CFSSCD` → `VARCHAR` | `cfmast.CFSSCD = 'PP'` |
| IC-003783 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:231](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L231) | CASE / P12 / COMPARISON | `cfmast.CFSSCD` : `CHAR(4)` | `'KTP'` : `varchar` | `cfmast.CFSSCD` → `VARCHAR` | `cfmast.CFSSCD = 'KTP'` |
| IC-003784 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:231](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L231) | CASE / P12 / COMPARISON | `cfmast.CFSSCD` : `CHAR(4)` | `'KTPS'` : `varchar` | `cfmast.CFSSCD` → `VARCHAR` | `cfmast.CFSSCD = 'KTPS'` |
| IC-003787 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:238](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L238) | CASE / P12 / COMPARISON | `cfmast.CFSSCD` : `CHAR(4)` | `'PP'` : `varchar` | `cfmast.CFSSCD` → `VARCHAR` | `cfmast.CFSSCD = 'PP'` |
| IC-003790 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:244](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L244) | CASE / P12 / COMPARISON | `cfmast.CFSSCD` : `CHAR(4)` | `'PP'` : `varchar` | `cfmast.CFSSCD` → `VARCHAR` | `cfmast.CFSSCD = 'PP'` |
| IC-003875 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:333](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L333) | JOIN_ON / P12 / COMPARISON | `CFSSCD` : `CHAR(4)` | `'NPWP'` : `varchar` | `CFSSCD` → `VARCHAR` | `CFSSCD = 'NPWP'` |
| IC-003881 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:341](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L341) | JOIN_ON / P12 / COMPARISON | `CFSSCD` : `CHAR(4)` | `'SKD'` : `varchar` | `CFSSCD` → `VARCHAR` | `CFSSCD = 'SKD'` |
| IC-003884 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:347](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L347) | JOIN_ON / P12 / COMPARISON | `CFSSCD` : `CHAR(4)` | `'APP'` : `varchar` | `CFSSCD` → `VARCHAR` | `CFSSCD = 'APP'` |
| IC-003887 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:353](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L353) | JOIN_ON / P12 / COMPARISON | `CFSSCD` : `CHAR(4)` | `'SIUP'` : `varchar` | `CFSSCD` → `VARCHAR` | `CFSSCD = 'SIUP'` |
| IC-004428 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSJobSyncKITAS:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSJobSyncKITAS.sql#L26) | WHERE / P12 / COMPARISON | `[CFSSCD]` : `CHAR(4)` | `'KITS'` : `varchar` | `[CFSSCD]` → `VARCHAR` | `[CFSSCD] = 'KITS'` |
| IC-004762 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSOBLExpiredKitasRPT:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOBLExpiredKitasRPT.sql#L48) | WHERE / P12 / COMPARISON | `CM.CFSSCD` : `CHAR(4)` | `'KITS'` : `varchar` | `CM.CFSSCD` → `VARCHAR` | `CM.CFSSCD = 'KITS'` |
| IC-005817 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUpload:261](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUpload.sql#L261) | JOIN_ON / P12 / IN_VALUE_TYPE | `dt.CFSSCD` : `CHAR(4)` | `'KTP'` : `varchar` | `dt.CFSSCD` → `VARCHAR` | `dt.CFSSCD IN ('KTP', 'KTPS')` |
| IC-008496 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTrxCorporateBOND:151](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L151) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `cfmast.CFSSCD` : `CHAR(4)` | `'KTP'` : `varchar` | `cfmast.CFSSCD` → `VARCHAR` | `when 'KTP' then cfmast.CFSSNO` |
| IC-008497 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTrxCorporateBOND:152](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L152) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `cfmast.CFSSCD` : `CHAR(4)` | `'KTPS'` : `varchar` | `cfmast.CFSSCD` → `VARCHAR` | `when 'KTPS' then cfmast.CFSSNO` |
| IC-008499 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTrxCorporateBOND:157](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L157) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `cfmast.CFSSCD` : `CHAR(4)` | `'PP'` : `varchar` | `cfmast.CFSSCD` → `VARCHAR` | `when 'PP' then cfmast.CFSSNO` |
| IC-008511 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTrxCorporateBOND:186](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L186) | CASE / P12 / COMPARISON | `cfmast.CFSSCD` : `CHAR(4)` | `'PP'` : `varchar` | `cfmast.CFSSCD` → `VARCHAR` | `cfmast.CFSSCD = 'PP'` |
| IC-008515 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTrxCorporateBOND:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L193) | CASE / P12 / COMPARISON | `cfmast.CFSSCD` : `CHAR(4)` | `'KTP'` : `varchar` | `cfmast.CFSSCD` → `VARCHAR` | `cfmast.CFSSCD = 'KTP'` |
| IC-008516 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTrxCorporateBOND:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L193) | CASE / P12 / COMPARISON | `cfmast.CFSSCD` : `CHAR(4)` | `'KTPS'` : `varchar` | `cfmast.CFSSCD` → `VARCHAR` | `cfmast.CFSSCD = 'KTPS'` |
| IC-009159 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:364](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L364) | CASE / P12 / COMPARISON | `cf.CFSSCD` : `CHAR(4)` | `'SIM'` : `varchar` | `cf.CFSSCD` → `VARCHAR` | `cf.CFSSCD = 'SIM'` |
| IC-009173 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:386](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L386) | CASE / P12 / COMPARISON | `cf.CFSSCD` : `CHAR(4)` | `'SIM'` : `varchar` | `cf.CFSSCD` → `VARCHAR` | `cf.CFSSCD <> 'SIM'` |
| IC-009326 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:262](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L262) | CASE / P12 / COMPARISON | `cf.CFSSCD` : `CHAR(4)` | `'SIM'` : `varchar` | `cf.CFSSCD` → `VARCHAR` | `cf.CFSSCD = 'SIM'` |
| IC-009340 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:284](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L284) | CASE / P12 / COMPARISON | `cf.CFSSCD` : `CHAR(4)` | `'SIM'` : `varchar` | `cf.CFSSCD` → `VARCHAR` | `cf.CFSSCD <> 'SIM'` |
| IC-012397 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncMasterIdAndFunc:10](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterIdAndFunc.sql#L10) | CASE / P12 / COMPARISON | `mast.CFSSCD` : `CHAR(4)` | `'SIM'` : `varchar` | `mast.CFSSCD` → `VARCHAR` | `mast.CFSSCD = 'SIM'` |
| IC-012398 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncMasterIdAndFunc:11](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterIdAndFunc.sql#L11) | CASE / P12 / COMPARISON | `mast.CFSSCD` : `CHAR(4)` | `'PP'` : `varchar` | `mast.CFSSCD` → `VARCHAR` | `mast.CFSSCD = 'PP'` |
| IC-012405 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncMasterIdAndFunc:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterIdAndFunc.sql#L26) | CASE / P12 / COMPARISON | `mast.CFSSCD` : `CHAR(4)` | `'SIM'` : `varchar` | `mast.CFSSCD` → `VARCHAR` | `mast.CFSSCD = 'SIM'` |
| IC-012406 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncMasterIdAndFunc:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterIdAndFunc.sql#L27) | CASE / P12 / COMPARISON | `mast.CFSSCD` : `CHAR(4)` | `'PP'` : `varchar` | `mast.CFSSCD` → `VARCHAR` | `mast.CFSSCD = 'PP'` |

<a id="kolom-cfuic2"></a>
## CFUIC2

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFMAST.sql#L63). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 11 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000754 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMTrxSecPreBook:410](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMTrxSecPreBook.sql#L410) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CFUIC2` : `CHAR(1)` | `@cMaritalStatus` : `VARCHAR(10)` | `CFUIC2` → `VARCHAR(10)` | `@cMaritalStatus = CFUIC2` |
| IC-008915 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailCalculateFee2:181](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee2.sql#L181) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cf.[CFUIC2]` : `CHAR(1)` | `tmp.[CFUIC2]` : `VARCHAR(1)` | `cf.[CFUIC2]` → `VARCHAR(1)` | `tmp.[CFUIC2] = cf.[CFUIC2]` |
| IC-003702 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:295](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L295) | CASE / P12 / COMPARISON | `cfmast.CFUIC2` : `CHAR(1)` | `'B'` : `varchar` | `cfmast.CFUIC2` → `VARCHAR` | `cfmast.CFUIC2 = 'B'` |
| IC-003703 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:296](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L296) | CASE / P12 / COMPARISON | `cfmast.CFUIC2` : `CHAR(1)` | `'A'` : `varchar` | `cfmast.CFUIC2` → `VARCHAR` | `cfmast.CFUIC2 = 'A'` |
| IC-003704 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:297](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L297) | CASE / P12 / COMPARISON | `cfmast.CFUIC2` : `CHAR(1)` | `'C'` : `varchar` | `cfmast.CFUIC2` → `VARCHAR` | `cfmast.CFUIC2 = 'C'` |
| IC-003798 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:258](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L258) | CASE / P12 / COMPARISON | `cfmast.CFUIC2` : `CHAR(1)` | `'B'` : `varchar` | `cfmast.CFUIC2` → `VARCHAR` | `cfmast.CFUIC2 = 'B'` |
| IC-003799 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:259](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L259) | CASE / P12 / COMPARISON | `cfmast.CFUIC2` : `CHAR(1)` | `'A'` : `varchar` | `cfmast.CFUIC2` → `VARCHAR` | `cfmast.CFUIC2 = 'A'` |
| IC-003800 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:260](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L260) | CASE / P12 / COMPARISON | `cfmast.CFUIC2` : `CHAR(1)` | `'C'` : `varchar` | `cfmast.CFUIC2` → `VARCHAR` | `cfmast.CFUIC2 = 'C'` |
| IC-008504 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTrxCorporateBOND:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L169) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `cfmast.CFUIC2` : `CHAR(1)` | `'B'` : `varchar` | `cfmast.CFUIC2` → `VARCHAR` | `when 'B' then '1'` |
| IC-008505 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTrxCorporateBOND:170](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L170) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `cfmast.CFUIC2` : `CHAR(1)` | `'A'` : `varchar` | `cfmast.CFUIC2` → `VARCHAR` | `when 'A' then '2'` |
| IC-008506 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTrxCorporateBOND:171](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L171) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `cfmast.CFUIC2` : `CHAR(1)` | `'C'` : `varchar` | `cfmast.CFUIC2` → `VARCHAR` | `when 'C' then '3'` |

<a id="kolom-cfuic3"></a>
## CFUIC3

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFMAST.sql#L64). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003705 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailIndividu:300](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L300) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfmast.CFUIC3` : `CHAR(1)` | `EDUCATION` : `VARCHAR(43)` | `cfmast.CFUIC3` → `VARCHAR(43)` | `EDUCATION = cfmast.CFUIC3` |
| IC-003801 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailIndividuOrchid:263](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L263) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfmast.CFUIC3` : `CHAR(1)` | `EDUCATION` : `VARCHAR(43)` | `cfmast.CFUIC3` → `VARCHAR(43)` | `EDUCATION = cfmast.CFUIC3` |

<a id="kolom-cfuic8"></a>
## CFUIC8

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFMAST.sql#L69). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012472 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSUpdateDataMasterNasabah:112](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateDataMasterNasabah.sql#L112) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `c.CFUIC8` : `CHAR(1)` | `#tmpCFMAST.CFUIC8` : `VARCHAR(1)` | `c.CFUIC8` → `VARCHAR(1)` | `#tmpCFMAST.CFUIC8 <- c.CFUIC8` |
| IC-010308 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptMonitoringDocAcc:187](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptMonitoringDocAcc.sql#L187) | WHERE / P12 / COMPARISON | `isnull(cf.[CFUIC8], '')` : `CHAR(1)` | `''` : `varchar` | `isnull(cf.[CFUIC8], '')` → `VARCHAR` | `isnull(cf.[CFUIC8], '') <> ''` |
| IC-012416 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSyncMasterRiskProfileData:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterRiskProfileData.sql#L23) | JOIN_ON / P12 / COMPARISON | `CFMAST.CFUIC8` : `CHAR(1)` | `CFPAR3.CP3UCD` : `VARCHAR(1)` | `CFMAST.CFUIC8` → `VARCHAR` | `CFMAST.CFUIC8 = CFPAR3.CP3UCD` |
| IC-012720 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSValidateSecAccNo:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateSecAccNo.sql#L30) | WHERE / P12 / COMPARISON | `isnull(CFUIC8, '')` : `CHAR(1)` | `''` : `varchar` | `isnull(CFUIC8, '')` → `VARCHAR` | `isnull(CFUIC8, '') <> ''` |
| IC-010309 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptMonitoringDocAcc:187](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptMonitoringDocAcc.sql#L187) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cf.[CFUIC8]` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(cf.[CFUIC8], '')` |
| IC-012721 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateSecAccNo:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateSecAccNo.sql#L30) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `CFUIC8` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(CFUIC8, '')` |

<a id="kolom-cfybip"></a>
## CFYBIP

**Tipe dasar:** `CHAR(30)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFMAST.sql#L75). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008500 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportTrxCorporateBOND:160](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L160) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfmast.CFYBIP` : `CHAR(30)` | `tr.investorBirthPlace` : `varchar(30)` | `cfmast.CFYBIP` → `varchar(30)` | `tr.investorBirthPlace = cfmast.CFYBIP` |
| IC-009154 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailSyncDataCore:348](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L348) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cf.CFYBIP` : `CHAR(30)` | `newTempatLahir` : `VARCHAR(50)` | `cf.CFYBIP` → `VARCHAR(50)` | `newTempatLahir = cf.CFYBIP` |
| IC-009321 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailSyncDataCoreWorkingHours:246](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L246) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cf.CFYBIP` : `CHAR(30)` | `newTempatLahir` : `VARCHAR(50)` | `cf.CFYBIP` → `VARCHAR(50)` | `newTempatLahir = cf.CFYBIP` |
| IC-009165 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:381](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L381) | CASE / P12 / COMPARISON | `cf.CFYBIP` : `CHAR(30)` | `t.TempatLahir` : `VARCHAR(50)` | `cf.CFYBIP` → `VARCHAR` | `t.TempatLahir <> cf.CFYBIP` |
| IC-009332 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:279](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L279) | CASE / P12 / COMPARISON | `cf.CFYBIP` : `CHAR(30)` | `t.TempatLahir` : `VARCHAR(50)` | `cf.CFYBIP` → `VARCHAR` | `t.TempatLahir <> cf.CFYBIP` |

<a id="kolom-cfyidp"></a>
## CFYIDP

**Tipe dasar:** `CHAR(30)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFMAST.sql#L73). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004761 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSOBLExpiredKitasRPT:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOBLExpiredKitasRPT.sql#L44) | EXPRESSION / P07 / COALESCE_BRANCH | `CM.CFYIDP` : `CHAR(30)` | `result type` : `VARCHAR (combined result; facets unknown)` | `CM.CFYIDP` → `VARCHAR` | `COALESCE(CM.CFYIDP, '')` |
