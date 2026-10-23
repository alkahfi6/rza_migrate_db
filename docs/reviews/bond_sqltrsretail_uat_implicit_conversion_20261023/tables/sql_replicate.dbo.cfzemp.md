# SQL_REPLICATE.dbo.CFZEMP — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 15 lokasi source, 7 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cfcif"></a>
## CFCIF

**Tipe dasar:** `CHAR(19)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFZEMP.sql#L2). **Existing leading index:** idx_cfzemp_cfcif.

**Frekuensi source:** 8 lokasi / 7 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000825 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMValidateCustomer:92](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L92) | WHERE / P12 / COMPARISON | `a.CFCIF` : `CHAR(19)` | `@pcCIFNo` : `varchar(19)` | `a.CFCIF` → `VARCHAR` | `a.CFCIF = @pcCIFNo` |
| IC-003728 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:419](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L419) | JOIN_ON / P12 / COMPARISON | `cfz.CFCIF` : `CHAR(19)` | `mnc._CIFNo` : `VARCHAR(19)` | `cfz.CFCIF` → `VARCHAR` | `mnc._CIFNo = cfz.CFCIF` |
| IC-003736 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:462](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L462) | JOIN_ON / P12 / COMPARISON | `cfz.CFCIF` : `CHAR(19)` | `mnc._CIFNo` : `VARCHAR(19)` | `cfz.CFCIF` → `VARCHAR` | `mnc._CIFNo = cfz.CFCIF` |
| IC-003828 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:406](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L406) | JOIN_ON / P12 / COMPARISON | `cfz.CFCIF` : `CHAR(19)` | `mnc._CIFNo` : `VARCHAR(19)` | `cfz.CFCIF` → `VARCHAR` | `mnc._CIFNo = cfz.CFCIF` |
| IC-003859 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:243](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L243) | JOIN_ON / P12 / COMPARISON | `cfz.CFCIF` : `CHAR(19)` | `mnc._CIFNo` : `VARCHAR(19)` | `cfz.CFCIF` → `VARCHAR` | `mnc._CIFNo = cfz.CFCIF` |
| IC-008526 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTrxCorporateBOND:252](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L252) | JOIN_ON / P12 / COMPARISON | `data.CFCIF` : `CHAR(19)` | `tr.CIFNo` : `varchar(20)` | `data.CFCIF` → `VARCHAR` | `data.CFCIF = tr.CIFNo` |
| IC-009182 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:416](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L416) | JOIN_ON / P12 / COMPARISON | `cz.CFCIF` : `CHAR(19)` | `t.CIFNo` : `VARCHAR(19)` | `cz.CFCIF` → `VARCHAR` | `t.CIFNo = cz.CFCIF` |
| IC-009349 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:314](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L314) | JOIN_ON / P12 / COMPARISON | `cz.CFCIF` : `CHAR(19)` | `t.CIFNo` : `VARCHAR(19)` | `cz.CFCIF` → `VARCHAR` | `t.CIFNo = cz.CFCIF` |

<a id="kolom-cfenin"></a>
## CFENIN

**Tipe dasar:** `CHAR(5)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFZEMP.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003725 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:412](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L412) | EXPRESSION / P07 / CASE_BRANCH | `cfz.CFENIN` : `CHAR(5)` | `result type` : `VARCHAR (combined result; facets unknown)` | `cfz.CFENIN` → `VARCHAR` | `CASE WHEN mo.KodeOccupation = 5 THEN cfz.CFENIN ELSE mo.Keterangan END` |
| IC-003825 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:399](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L399) | EXPRESSION / P07 / CASE_BRANCH | `cfz.CFENIN` : `CHAR(5)` | `result type` : `VARCHAR (combined result; facets unknown)` | `cfz.CFENIN` → `VARCHAR` | `CASE WHEN mo.KodeOccupation = 5 THEN cfz.CFENIN ELSE mo.Keterangan END` |

<a id="kolom-cfoccd"></a>
## CFOCCD

**Tipe dasar:** `CHAR(5)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFZEMP.sql#L19). **Existing leading index:** idx_cfzemp_cfoccd.

**Frekuensi source:** 5 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003729 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:421](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L421) | JOIN_ON / P12 / COMPARISON | `cfz.CFOCCD` : `CHAR(5)` | `mo.KodePekerjaan` : `VARCHAR(20)` | `cfz.CFOCCD` → `VARCHAR` | `cfz.CFOCCD = mo.KodePekerjaan` |
| IC-003737 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:464](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L464) | JOIN_ON / P12 / COMPARISON | `cfz.CFOCCD` : `CHAR(5)` | `mo.KodePekerjaan` : `VARCHAR(20)` | `cfz.CFOCCD` → `VARCHAR` | `cfz.CFOCCD = mo.KodePekerjaan` |
| IC-003829 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:408](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L408) | JOIN_ON / P12 / COMPARISON | `cfz.CFOCCD` : `CHAR(5)` | `mo.KodePekerjaan` : `VARCHAR(20)` | `cfz.CFOCCD` → `VARCHAR` | `cfz.CFOCCD = mo.KodePekerjaan` |
| IC-003860 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:245](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L245) | JOIN_ON / P12 / COMPARISON | `cfz.CFOCCD` : `CHAR(5)` | `mo.KodePekerjaan` : `VARCHAR(20)` | `cfz.CFOCCD` → `VARCHAR` | `cfz.CFOCCD = mo.KodePekerjaan` |
| IC-008527 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTrxCorporateBOND:254](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L254) | JOIN_ON / P12 / COMPARISON | `data.CFOCCD` : `CHAR(5)` | `occ.KodePekerjaan` : `VARCHAR(20)` | `data.CFOCCD` → `VARCHAR` | `occ.KodePekerjaan = data.CFOCCD` |
