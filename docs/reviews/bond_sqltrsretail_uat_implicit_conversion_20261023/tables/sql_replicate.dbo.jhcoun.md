# SQL_REPLICATE.dbo.JHCOUN — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 30 lokasi source, 12 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 3 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-jhccoc"></a>
## JHCCOC

**Tipe dasar:** `VARCHAR(3)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.JHCOUN.sql#L2). **Existing leading index:** idx_jhccoc.

**Frekuensi source:** 13 lokasi / 11 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008477 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportTemplateSID:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSID.sql#L130) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `JHCCOC` : `VARCHAR(3)` | `Nationality` : `CHAR(20)` | `JHCCOC` → `CHAR(20)` | `Nationality = JHCCOC` |
| IC-003683 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailIndividu:257](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L257) | JOIN_ON / P12 / COMPARISON | `jhc.JHCCOC` : `VARCHAR(3)` | `mnc.NATIONALITY` : `CHAR(10)` | `mnc.NATIONALITY` → `VARCHAR` | `mnc.NATIONALITY = jhc.JHCCOC` |
| IC-003782 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailIndividuOrchid:226](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L226) | JOIN_ON / P12 / COMPARISON | `jhc.JHCCOC` : `VARCHAR(3)` | `mnc.NATIONALITYOBL` : `CHAR(10)` | `mnc.NATIONALITYOBL` → `VARCHAR` | `mnc.NATIONALITYOBL = jhc.JHCCOC` |
| IC-003853 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailInstitusi:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L225) | JOIN_ON / P12 / COMPARISON | `jhc.JHCCOC` : `VARCHAR(3)` | `mnc.LEGAL_DOMICILE` : `CHAR(10)` | `mnc.LEGAL_DOMICILE` → `VARCHAR` | `mnc.LEGAL_DOMICILE = jhc.JHCCOC` |
| IC-005402 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateEditCustomerAuthorization:476](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEditCustomerAuthorization.sql#L476) | JOIN_ON / P12 / COMPARISON | `jh.JHCCOC` : `VARCHAR(3)` | `tcptm.Citizenship` : `CHAR(3)` | `tcptm.Citizenship` → `VARCHAR` | `jh.JHCCOC = tcptm.Citizenship` |
| IC-006536 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateWarganegaraNasabah:6](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateWarganegaraNasabah.sql#L6) | JOIN_ON / P12 / COMPARISON | `jh.JHCCOC` : `VARCHAR(3)` | `cf.CFCITZ` : `CHAR(3)` | `cf.CFCITZ` → `VARCHAR` | `jh.JHCCOC = cf.CFCITZ` |
| IC-007876 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRPTMasterNasabahObligasi:443](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahObligasi.sql#L443) | JOIN_ON / P12 / COMPARISON | `jc.JHCCOC` : `VARCHAR(3)` | `v.CFCITZ` : `CHAR(3)` | `v.CFCITZ` → `VARCHAR` | `v.CFCITZ = jc.JHCCOC` |
| IC-008476 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportTemplateSID:120](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSID.sql#L120) | JOIN_ON / P12 / COMPARISON | `jhc.JHCCOC` : `VARCHAR(3)` | `tc.Citizenship` : `CHAR(3)` | `tc.Citizenship` → `VARCHAR` | `tc.Citizenship = jhc.JHCCOC` |
| IC-008488 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportTemplateSIDBySecAccNo:76](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSIDBySecAccNo.sql#L76) | JOIN_ON / P12 / COMPARISON | `jhc.JHCCOC` : `VARCHAR(3)` | `tc.Citizenship` : `CHAR(3)` | `tc.Citizenship` → `VARCHAR` | `tc.Citizenship = jhc.JHCCOC` |
| IC-008520 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportTrxCorporateBOND:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L216) | JOIN_ON / P12 / COMPARISON | `jhc.JHCCOC` : `VARCHAR(3)` | `tr.CFCOUN` : `char(3)` | `tr.CFCOUN` → `VARCHAR` | `jhc.JHCCOC = tr.CFCOUN` |
| IC-009816 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptDetailTrx:749](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L749) | JOIN_ON / P12 / COMPARISON | `jhc.JHCCOC` : `VARCHAR(3)` | `t.CitizenshipCode` : `CHAR(3)` | `t.CitizenshipCode` → `VARCHAR` | `t.CitizenshipCode = jhc.JHCCOC` |
| IC-013474 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_ListStatusTreasuryCustomer_TM:84](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L84) | JOIN_ON / P12 / COMPARISON | `jh.JHCCOC` : `VARCHAR(3)` | `tc.Citizenship` : `CHAR(3)` | `tc.Citizenship` → `VARCHAR` | `jh.JHCCOC = tc.Citizenship` |
| IC-013485 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_ListStatusTreasuryCustomer_TM:144](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L144) | JOIN_ON / P12 / COMPARISON | `jh.JHCCOC` : `VARCHAR(3)` | `tc.Citizenship` : `CHAR(3)` | `tc.Citizenship` → `VARCHAR` | `jh.JHCCOC = tc.Citizenship` |

<a id="kolom-jhcnam"></a>
## JHCNAM

**Tipe dasar:** `CHAR(20)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.JHCOUN.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 7 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008519 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportTrxCorporateBOND:212](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L212) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `jhc.JHCNAM` : `CHAR(20)` | `tr.InvestorCountry` : `varchar(50)` | `jhc.JHCNAM` → `varchar(50)` | `tr.InvestorCountry = jhc.JHCNAM` |
| IC-009815 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptDetailTrx:746](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L746) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `jhc.JHCNAM` : `CHAR(20)` | `Citizenship` : `VARCHAR(20)` | `jhc.JHCNAM` → `VARCHAR(20)` | `Citizenship = jhc.JHCNAM` |
| IC-005342 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateCountry:7](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateCountry.sql#L7) | SET_OUTPUT / P07 / UNION_OUTPUT | `right output #2` : `CHAR(20)` | `left output #2` : `varchar` | `right output #2` → `VARCHAR` | `select ''as 'Value', '' as 'Name' union all select JHCCOC, JHCNAM from JHCOUN_v where JHCOD = 'CT'` |
| IC-005396 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateEditCustomerAuthorization:440](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEditCustomerAuthorization.sql#L440) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `jh.JHCNAM` : `CHAR(20)` | `''` : `varchar` | `''` → `CHAR(20)` | `isnull(jh.JHCNAM,'')` |
| IC-008474 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSReportTemplateSID:87](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSID.sql#L87) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `jhc.JHCNAM` : `CHAR(20)` | `''` : `varchar` | `''` → `CHAR(20)` | `isnull(jhc.JHCNAM,'')` |
| IC-008486 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSReportTemplateSIDBySecAccNo:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSIDBySecAccNo.sql#L47) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `jhc.JHCNAM` : `CHAR(20)` | `''` : `varchar` | `''` → `CHAR(20)` | `isnull(jhc.JHCNAM,'')` |
| IC-013469 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusTreasuryCustomer_TM:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L61) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `jh.JHCNAM` : `CHAR(20)` | `''` : `varchar` | `''` → `CHAR(20)` | `isnull(jh.JHCNAM,'')` |
| IC-013480 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusTreasuryCustomer_TM:119](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L119) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `jh.JHCNAM` : `CHAR(20)` | `''` : `varchar` | `''` → `CHAR(20)` | `isnull(jh.JHCNAM,'')` |

<a id="kolom-jhcod"></a>
## JHCOD

**Tipe dasar:** `CHAR(2)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.JHCOUN.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 7 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005343 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateCountry:10](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateCountry.sql#L10) | WHERE / P12 / COMPARISON | `JHCOD` : `CHAR(2)` | `'CT'` : `varchar` | `JHCOD` → `VARCHAR` | `JHCOD = 'CT'` |
| IC-005403 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateEditCustomerAuthorization:476](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEditCustomerAuthorization.sql#L476) | JOIN_ON / P12 / COMPARISON | `jh.JHCOD` : `CHAR(2)` | `'CT'` : `varchar` | `jh.JHCOD` → `VARCHAR` | `jh.JHCOD = 'CT'` |
| IC-006538 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateWarganegaraNasabah:8](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateWarganegaraNasabah.sql#L8) | WHERE / P12 / COMPARISON | `jh.JHCOD` : `CHAR(2)` | `'CT'` : `varchar` | `jh.JHCOD` → `VARCHAR` | `jh.JHCOD = 'CT'` |
| IC-008475 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTemplateSID:120](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSID.sql#L120) | JOIN_ON / P12 / COMPARISON | `jhc.JHCOD` : `CHAR(2)` | `'CT'` : `varchar` | `jhc.JHCOD` → `VARCHAR` | `jhc.JHCOD = 'CT'` |
| IC-008478 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTemplateSID:134](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSID.sql#L134) | JOIN_ON / P12 / COMPARISON | `coun.JHCOD` : `CHAR(2)` | `'SW'` : `varchar` | `coun.JHCOD` → `VARCHAR` | `coun.JHCOD = 'SW'` |
| IC-008487 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTemplateSIDBySecAccNo:76](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSIDBySecAccNo.sql#L76) | JOIN_ON / P12 / COMPARISON | `jhc.JHCOD` : `CHAR(2)` | `'CT'` : `varchar` | `jhc.JHCOD` → `VARCHAR` | `jhc.JHCOD = 'CT'` |
| IC-009817 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDetailTrx:750](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L750) | JOIN_ON / P12 / COMPARISON | `jhc.JHCOD` : `CHAR(2)` | `'CT'` : `varchar` | `jhc.JHCOD` → `VARCHAR` | `jhc.JHCOD = 'CT'` |
| IC-013475 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListStatusTreasuryCustomer_TM:84](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L84) | JOIN_ON / P12 / COMPARISON | `jh.JHCOD` : `CHAR(2)` | `'CT'` : `varchar` | `jh.JHCOD` → `VARCHAR` | `jh.JHCOD = 'CT'` |
| IC-013486 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListStatusTreasuryCustomer_TM:144](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L144) | JOIN_ON / P12 / COMPARISON | `jh.JHCOD` : `CHAR(2)` | `'CT'` : `varchar` | `jh.JHCOD` → `VARCHAR` | `jh.JHCOD = 'CT'` |
