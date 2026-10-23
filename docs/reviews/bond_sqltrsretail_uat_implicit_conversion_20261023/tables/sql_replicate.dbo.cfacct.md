# SQL_REPLICATE.dbo.CFACCT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 40 lokasi source, 9 procedure, 5 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 15 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cfacc"></a>
## CFACC#

**Tipe dasar:** `DECIMAL(19,0)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFACCT.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 12 lokasi / 7 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000542 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.GetRekeningMasterNasabah:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L29) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFACC#` : `DECIMAL(19,0)` | `#tmpRek.AccountNo` : `VARCHAR(50)` | `CFACC#` → `VARCHAR(50)` | `#tmpRek.AccountNo <- CFACC#` |
| IC-013598 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_PopulateDataRekeningMaterai:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningMaterai.sql#L37) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFACC#` : `DECIMAL(19,0)` | `#tempRekMaterai.NoRekInvestor` : `BIGINT` | `CFACC#` → `BIGINT` | `#tempRekMaterai.NoRekInvestor <- CFACC#` |
| IC-013602 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_PopulateDataRekeningMaterai:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningMaterai.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFACC#` : `DECIMAL(19,0)` | `#tempRekMaterai.NoRekInvestor` : `BIGINT` | `CFACC#` → `BIGINT` | `#tempRekMaterai.NoRekInvestor <- CFACC#` |
| IC-013619 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_PopulateDataRekeningObligasi:42](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningObligasi.sql#L42) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFACC#` : `DECIMAL(19,0)` | `#temp.NoRekInvestor` : `bigint` | `CFACC#` → `bigint` | `#temp.NoRekInvestor <- CFACC#` |
| IC-013623 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_PopulateDataRekeningObligasi:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningObligasi.sql#L48) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFACC#` : `DECIMAL(19,0)` | `#temp.NoRekInvestor` : `bigint` | `CFACC#` → `bigint` | `#temp.NoRekInvestor <- CFACC#` |
| IC-003211 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSAmendSecurityTransaction_TT:694](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L694) | WHERE / P04 / COMPARISON | `CFACC#` : `DECIMAL(19,0)` | `@cNoRekInvestor` : `varchar(20)` | `@cNoRekInvestor` → `DECIMAL` | `CFACC# = @cNoRekInvestor` |
| IC-003213 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSAmendSecurityTransaction_TT:703](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L703) | JOIN_ON / P01 / COMPARISON | `a.CFACC#` : `DECIMAL(19,0)` | `temp.NoRekInvestor` : `varchar(20)` | `temp.NoRekInvestor` → `DECIMAL` | `a.CFACC# = temp.NoRekInvestor` |
| IC-004329 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSInsertSecurityTransaction_TT_Original:1212](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L1212) | WHERE / P04 / COMPARISON | `CFACC#` : `DECIMAL(19,0)` | `@cNoRekInvestor` : `varchar(20)` | `@cNoRekInvestor` → `DECIMAL` | `CFACC# = @cNoRekInvestor` |
| IC-012670 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSUpdateTaxAmnestyExpired:281](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateTaxAmnestyExpired.sql#L281) | JOIN_ON / P01 / COMPARISON | `c.CFACC#` : `DECIMAL(19,0)` | `a.NoRekInvestor` : `VARCHAR(20)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = c.CFACC#` |
| IC-013095 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_InsertORI_Order_TT_Original:458](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORI_Order_TT_Original.sql#L458) | WHERE / P04 / COMPARISON | `CFACC#` : `DECIMAL(19,0)` | `@cRekeningRelasi` : `VARCHAR(20)` | `@cRekeningRelasi` → `DECIMAL` | `CFACC# = @cRekeningRelasi` |
| IC-013603 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_PopulateDataRekeningMaterai:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningMaterai.sql#L47) | JOIN_ON / P11 / COMPARISON | `CFACC#` : `DECIMAL(19,0)` | `ta.NoRekInvestor` : `BIGINT` | `ta.NoRekInvestor` → `DECIMAL` | `CFACC# = ta.NoRekInvestor` |
| IC-013624 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_PopulateDataRekeningObligasi:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningObligasi.sql#L52) | JOIN_ON / P11 / COMPARISON | `CFACC#` : `DECIMAL(19,0)` | `ta.NoRekInvestor` : `bigint` | `ta.NoRekInvestor` → `DECIMAL` | `CFACC# = ta.NoRekInvestor` |

<a id="kolom-cfatyp"></a>
## CFATYP

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFACCT.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000543 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.GetRekeningMasterNasabah:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L29) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFATYP` : `CHAR(1)` | `#tmpRek.AccountType` : `VARCHAR(5)` | `CFATYP` → `VARCHAR(5)` | `#tmpRek.AccountType <- CFATYP` |
| IC-000482 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.GetAccountInfo:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountInfo.sql#L43) | WHERE / P12 / IN_VALUE_TYPE | `CFATYP` : `CHAR(1)` | `'S'` : `varchar` | `CFATYP` → `VARCHAR` | `CFATYP IN ('S','D')` |
| IC-000487 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.GetAccountInfo:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountInfo.sql#L51) | WHERE / P12 / IN_VALUE_TYPE | `CFATYP` : `CHAR(1)` | `'S'` : `varchar` | `CFATYP` → `VARCHAR` | `CFATYP IN ('S','D')` |
| IC-000545 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabah:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L37) | WHERE / P12 / IN_VALUE_TYPE | `CFATYP` : `CHAR(1)` | `'S'` : `varchar` | `CFATYP` → `VARCHAR` | `CFATYP IN ('S','D')` |
| IC-000546 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabah:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L45) | WHERE / P12 / IN_VALUE_TYPE | `CFATYP` : `CHAR(1)` | `'S'` : `varchar` | `CFATYP` → `VARCHAR` | `CFATYP IN ('S','D')` |
| IC-013600 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_PopulateDataRekeningMaterai:40](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningMaterai.sql#L40) | WHERE / P12 / IN_VALUE_TYPE | `CFATYP` : `CHAR(1)` | `'S'` : `varchar` | `CFATYP` → `VARCHAR` | `CFATYP IN ('S','D')` |
| IC-013605 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_PopulateDataRekeningMaterai:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningMaterai.sql#L48) | WHERE / P12 / IN_VALUE_TYPE | `CFATYP` : `CHAR(1)` | `'S'` : `varchar` | `CFATYP` → `VARCHAR` | `CFATYP IN ('S','D')` |
| IC-013621 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_PopulateDataRekeningObligasi:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningObligasi.sql#L45) | WHERE / P12 / IN_VALUE_TYPE | `CFATYP` : `CHAR(1)` | `'S'` : `varchar` | `CFATYP` → `VARCHAR` | `CFATYP IN ('S','D')` |
| IC-013626 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_PopulateDataRekeningObligasi:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningObligasi.sql#L53) | WHERE / P12 / IN_VALUE_TYPE | `CFATYP` : `CHAR(1)` | `'S'` : `varchar` | `CFATYP` → `VARCHAR` | `CFATYP IN ('S','D')` |

<a id="kolom-cfcif"></a>
## CFCIF#

**Tipe dasar:** `DECIMAL(19,0)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFACCT.sql#L3). **Existing leading index:** idx_cfcif#_cfatype_cfacc#_cfsnme_cfrela.

**Frekuensi source:** 12 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000479 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.GetAccountInfo:40](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountInfo.sql#L40) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFCIF#` : `DECIMAL(19,0)` | `@t_Account.NoCIF` : `VARCHAR(19)` | `CFCIF#` → `VARCHAR(19)` | `@t_Account.NoCIF <- CFCIF#` |
| IC-000484 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.GetAccountInfo:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountInfo.sql#L46) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFCIF#` : `DECIMAL(19,0)` | `@t_Account.NoCIF` : `VARCHAR(19)` | `CFCIF#` → `VARCHAR(19)` | `@t_Account.NoCIF <- CFCIF#` |
| IC-013597 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_PopulateDataRekeningMaterai:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningMaterai.sql#L37) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFCIF#` : `DECIMAL(19,0)` | `#tempRekMaterai.CIFNo` : `BIGINT` | `CFCIF#` → `BIGINT` | `#tempRekMaterai.CIFNo <- CFCIF#` |
| IC-013601 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_PopulateDataRekeningMaterai:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningMaterai.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFCIF#` : `DECIMAL(19,0)` | `#tempRekMaterai.CIFNo` : `BIGINT` | `CFCIF#` → `BIGINT` | `#tempRekMaterai.CIFNo <- CFCIF#` |
| IC-013618 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_PopulateDataRekeningObligasi:42](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningObligasi.sql#L42) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFCIF#` : `DECIMAL(19,0)` | `#temp.CIFNo` : `bigint` | `CFCIF#` → `bigint` | `#temp.CIFNo <- CFCIF#` |
| IC-013622 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_PopulateDataRekeningObligasi:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningObligasi.sql#L48) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFCIF#` : `DECIMAL(19,0)` | `#temp.CIFNo` : `bigint` | `CFCIF#` → `bigint` | `#temp.CIFNo <- CFCIF#` |
| IC-000481 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetAccountInfo:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountInfo.sql#L43) | WHERE / P04 / COMPARISON | `CFCIF#` : `DECIMAL(19,0)` | `@pcCIF` : `VARCHAR(19)` | `@pcCIF` → `DECIMAL` | `CFCIF# = @pcCIF` |
| IC-000486 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetAccountInfo:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountInfo.sql#L51) | WHERE / P04 / COMPARISON | `CFCIF#` : `DECIMAL(19,0)` | `@pcCIF` : `VARCHAR(19)` | `@pcCIF` → `DECIMAL` | `CFCIF# = @pcCIF` |
| IC-013599 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_PopulateDataRekeningMaterai:40](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningMaterai.sql#L40) | WHERE / P11 / COMPARISON | `CFCIF#` : `DECIMAL(19,0)` | `@nCIFNo` : `BIGINT` | `@nCIFNo` → `DECIMAL` | `CFCIF# = @nCIFNo` |
| IC-013604 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_PopulateDataRekeningMaterai:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningMaterai.sql#L48) | WHERE / P11 / COMPARISON | `CFCIF#` : `DECIMAL(19,0)` | `@nCIFNo` : `BIGINT` | `@nCIFNo` → `DECIMAL` | `CFCIF# = @nCIFNo` |
| IC-013620 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_PopulateDataRekeningObligasi:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningObligasi.sql#L45) | WHERE / P11 / COMPARISON | `CFCIF#` : `DECIMAL(19,0)` | `@nCIFNo` : `bigint` | `@nCIFNo` → `DECIMAL` | `CFCIF# = @nCIFNo` |
| IC-013625 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_PopulateDataRekeningObligasi:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningObligasi.sql#L53) | WHERE / P11 / COMPARISON | `CFCIF#` : `DECIMAL(19,0)` | `@nCIFNo` : `bigint` | `@nCIFNo` → `DECIMAL` | `CFCIF# = @nCIFNo` |

<a id="kolom-cfrela"></a>
## CFRELA

**Tipe dasar:** `CHAR(2)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFACCT.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002868 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedProcessRekDana:190](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L190) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `v.CFRELA` : `CHAR(2)` | `Relasi` : `VARCHAR(2)` | `v.CFRELA` → `VARCHAR(2)` | `Relasi = v.CFRELA` |
| IC-012667 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSUpdateTaxAmnestyExpired:274](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateTaxAmnestyExpired.sql#L274) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `isnull(c.CFRELA, '')` : `CHAR(2)` | `retail.dbo.trscustaccnumber_tr.Relasi` : `VARCHAR(2)` | `isnull(c.CFRELA, '')` → `VARCHAR(2)` | `retail.dbo.trscustaccnumber_tr.Relasi <- isnull(c.CFRELA, '')` |
| IC-003212 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSAmendSecurityTransaction_TT:695](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L695) | WHERE / P12 / COMPARISON | `CFRELA` : `CHAR(2)` | `'B'` : `varchar` | `CFRELA` → `VARCHAR` | `CFRELA = 'B'` |
| IC-003214 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSAmendSecurityTransaction_TT:705](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L705) | WHERE / P12 / COMPARISON | `a.CFRELA` : `CHAR(2)` | `'B'` : `varchar` | `a.CFRELA` → `VARCHAR` | `a.CFRELA = 'B'` |
| IC-013096 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_InsertORI_Order_TT_Original:459](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORI_Order_TT_Original.sql#L459) | WHERE / P12 / IN_VALUE_TYPE | `CFRELA` : `CHAR(2)` | `'P'` : `varchar` | `CFRELA` → `VARCHAR` | `CFRELA in ('P', 'B', 'C', 'D', 'E', 'F')` |
| IC-012668 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUpdateTaxAmnestyExpired:276](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateTaxAmnestyExpired.sql#L276) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `c.CFRELA` : `CHAR(2)` | `''` : `varchar` | `''` → `CHAR(2)` | `isnull(c.CFRELA, '')` |

<a id="kolom-cfsnme"></a>
## CFSNME

**Tipe dasar:** `CHAR(20)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFACCT.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000544 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.GetRekeningMasterNasabah:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L29) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CFSNME` : `CHAR(20)` | `#tmpRek.Nama` : `VARCHAR(200)` | `CFSNME` → `VARCHAR(200)` | `#tmpRek.Nama <- CFSNME` |
