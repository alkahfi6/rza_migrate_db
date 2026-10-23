# SQL_REPLICATE.dbo.CFALTN — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 18 lokasi source, 4 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 9 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cfaact"></a>
## CFAACT

**Tipe dasar:** `DECIMAL(19,0)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFALTN.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001016 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.ORISBNISubmitEarlyRedemption:408](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNISubmitEarlyRedemption.sql#L408) | JOIN_ON / P11 / COMPARISON | `smt.CFAACT` : `DECIMAL(19,0)` | `CAST(tdn.RekeningRelasi AS NUMERIC)` : `NUMERIC` | `CAST(tdn.RekeningRelasi AS NUMERIC)` → `DECIMAL` | `CAST(tdn.RekeningRelasi AS NUMERIC) = smt.CFAACT` |
| IC-003739 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailIndividu:530](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L530) | JOIN_ON / P11 / COMPARISON | `cfa.CFAACT` : `DECIMAL(19,0)` | `CAST(mnc.IDR_BANK_ACCOUNT_NO AS BIGINT)` : `BIGINT` | `CAST(mnc.IDR_BANK_ACCOUNT_NO AS BIGINT)` → `DECIMAL` | `CAST(mnc.IDR_BANK_ACCOUNT_NO AS BIGINT) = cfa.CFAACT` |
| IC-003741 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailIndividu:538](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L538) | JOIN_ON / P11 / COMPARISON | `cfa.CFAACT` : `DECIMAL(19,0)` | `CAST(mnc.USD_BANK_ACCOUNT_NO AS BIGINT)` : `BIGINT` | `CAST(mnc.USD_BANK_ACCOUNT_NO AS BIGINT)` → `DECIMAL` | `CAST(mnc.USD_BANK_ACCOUNT_NO AS BIGINT) = cfa.CFAACT` |
| IC-003743 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailIndividu:546](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L546) | JOIN_ON / P11 / COMPARISON | `cfa.CFAACT` : `DECIMAL(19,0)` | `CAST(mnc.EUR_BANK_ACCOUNT_NO AS BIGINT)` : `BIGINT` | `CAST(mnc.EUR_BANK_ACCOUNT_NO AS BIGINT)` → `DECIMAL` | `CAST(mnc.EUR_BANK_ACCOUNT_NO AS BIGINT) = cfa.CFAACT` |
| IC-003901 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailInstitusi:454](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L454) | JOIN_ON / P11 / COMPARISON | `cfa.CFAACT` : `DECIMAL(19,0)` | `CAST(mnc.IDR_BANK_ACCOUNT_NO AS BIGINT)` : `BIGINT` | `CAST(mnc.IDR_BANK_ACCOUNT_NO AS BIGINT)` → `DECIMAL` | `CAST(mnc.IDR_BANK_ACCOUNT_NO AS BIGINT) = cfa.CFAACT` |
| IC-003903 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailInstitusi:462](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L462) | JOIN_ON / P11 / COMPARISON | `cfa.CFAACT` : `DECIMAL(19,0)` | `CAST(mnc.USD_BANK_ACCOUNT_NO AS BIGINT)` : `BIGINT` | `CAST(mnc.USD_BANK_ACCOUNT_NO AS BIGINT)` → `DECIMAL` | `CAST(mnc.USD_BANK_ACCOUNT_NO AS BIGINT) = cfa.CFAACT` |
| IC-003905 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailInstitusi:470](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L470) | JOIN_ON / P11 / COMPARISON | `cfa.CFAACT` : `DECIMAL(19,0)` | `CAST(mnc.EUR_BANK_ACCOUNT_NO AS BIGINT)` : `BIGINT` | `CAST(mnc.EUR_BANK_ACCOUNT_NO AS BIGINT)` → `DECIMAL` | `CAST(mnc.EUR_BANK_ACCOUNT_NO AS BIGINT) = cfa.CFAACT` |
| IC-003965 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailInstitusiOrchid:380](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L380) | JOIN_ON / P11 / COMPARISON | `cfa.CFAACT` : `DECIMAL(19,0)` | `CAST(mnc.IDR_BANK_ACCOUNT_NO AS BIGINT)` : `BIGINT` | `CAST(mnc.IDR_BANK_ACCOUNT_NO AS BIGINT)` → `DECIMAL` | `CAST(mnc.IDR_BANK_ACCOUNT_NO AS BIGINT) = cfa.CFAACT` |
| IC-003967 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailInstitusiOrchid:388](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L388) | JOIN_ON / P11 / COMPARISON | `cfa.CFAACT` : `DECIMAL(19,0)` | `CAST(mnc.USD_BANK_ACCOUNT_NO AS BIGINT)` : `BIGINT` | `CAST(mnc.USD_BANK_ACCOUNT_NO AS BIGINT)` → `DECIMAL` | `CAST(mnc.USD_BANK_ACCOUNT_NO AS BIGINT) = cfa.CFAACT` |

<a id="kolom-cfaal1"></a>
## CFAAL1

**Tipe dasar:** `CHAR(40)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFALTN.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001015 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.ORISBNISubmitEarlyRedemption:405](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNISubmitEarlyRedemption.sql#L405) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `smt.CFAAL1` : `CHAR(40)` | `TOACCTNAME` : `VARCHAR(100)` | `smt.CFAAL1` → `VARCHAR(100)` | `TOACCTNAME = smt.CFAAL1` |
| IC-003738 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailIndividu:527](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L527) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfa.CFAAL1` : `CHAR(40)` | `IDR_BENEF_NAME` : `VARCHAR(200)` | `cfa.CFAAL1` → `VARCHAR(200)` | `IDR_BENEF_NAME = cfa.CFAAL1` |
| IC-003740 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailIndividu:535](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L535) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfa.CFAAL1` : `CHAR(40)` | `USD_BENEF_NAME` : `VARCHAR(300)` | `cfa.CFAAL1` → `VARCHAR(300)` | `USD_BENEF_NAME = cfa.CFAAL1` |
| IC-003742 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailIndividu:543](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L543) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfa.CFAAL1` : `CHAR(40)` | `EUR_BENEF_NAME` : `VARCHAR(300)` | `cfa.CFAAL1` → `VARCHAR(300)` | `EUR_BENEF_NAME = cfa.CFAAL1` |
| IC-003900 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailInstitusi:451](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L451) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfa.CFAAL1` : `CHAR(40)` | `IDR_BENEF_NAME` : `VARCHAR(200)` | `cfa.CFAAL1` → `VARCHAR(200)` | `IDR_BENEF_NAME = cfa.CFAAL1` |
| IC-003902 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailInstitusi:459](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L459) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfa.CFAAL1` : `CHAR(40)` | `USD_BENEF_NAME` : `VARCHAR(300)` | `cfa.CFAAL1` → `VARCHAR(300)` | `USD_BENEF_NAME = cfa.CFAAL1` |
| IC-003904 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailInstitusi:467](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L467) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfa.CFAAL1` : `CHAR(40)` | `EUR_BENEF_NAME` : `VARCHAR(300)` | `cfa.CFAAL1` → `VARCHAR(300)` | `EUR_BENEF_NAME = cfa.CFAAL1` |
| IC-003964 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailInstitusiOrchid:377](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L377) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfa.CFAAL1` : `CHAR(40)` | `IDR_BENEF_NAME` : `VARCHAR(200)` | `cfa.CFAAL1` → `VARCHAR(200)` | `IDR_BENEF_NAME = cfa.CFAAL1` |
| IC-003966 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailInstitusiOrchid:385](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L385) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cfa.CFAAL1` : `CHAR(40)` | `USD_BENEF_NAME` : `VARCHAR(300)` | `cfa.CFAAL1` → `VARCHAR(300)` | `USD_BENEF_NAME = cfa.CFAAL1` |
