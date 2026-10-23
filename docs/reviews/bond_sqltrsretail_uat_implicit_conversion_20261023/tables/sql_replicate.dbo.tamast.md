# SQL_REPLICATE.dbo.TAMAST — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 41 lokasi source, 17 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cfacc"></a>
## CFACC#

**Tipe dasar:** `DECIMAL(19,0)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.TAMAST.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 13 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000507 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetAccountInfo:172](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountInfo.sql#L172) | JOIN_ON / P01 / COMPARISON | `tms.[CFACC#]` : `DECIMAL(19,0)` | `tmp.ACCTNO` : `VARCHAR(20)` | `tmp.ACCTNO` → `DECIMAL` | `tmp.ACCTNO = tms.[CFACC#]` |
| IC-000557 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:117](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L117) | JOIN_ON / P01 / COMPARISON | `tms.[CFACC#]` : `DECIMAL(19,0)` | `ta.AccountNo` : `VARCHAR(50)` | `ta.AccountNo` → `DECIMAL` | `ta.AccountNo = tms.[CFACC#]` |
| IC-000560 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:133](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L133) | JOIN_ON / P01 / COMPARISON | `tms.[CFACC#]` : `DECIMAL(19,0)` | `ta.AccountNo` : `VARCHAR(50)` | `ta.AccountNo` → `DECIMAL` | `ta.AccountNo = tms.[CFACC#]` |
| IC-000564 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:155](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L155) | JOIN_ON / P01 / COMPARISON | `d.CFACC#` : `DECIMAL(19,0)` | `a.AccountNo` : `VARCHAR(50)` | `a.AccountNo` → `DECIMAL` | `a.AccountNo = d.CFACC#` |
| IC-000570 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L180) | JOIN_ON / P01 / COMPARISON | `d.CFACC#` : `DECIMAL(19,0)` | `a.AccountNo` : `VARCHAR(50)` | `a.AccountNo` → `DECIMAL` | `a.AccountNo = d.CFACC#` |
| IC-000578 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L216) | JOIN_ON / P01 / COMPARISON | `b.CFACC#` : `DECIMAL(19,0)` | `a.AccountNo` : `VARCHAR(50)` | `a.AccountNo` → `DECIMAL` | `a.AccountNo = b.CFACC#` |
| IC-000582 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:232](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L232) | JOIN_ON / P01 / COMPARISON | `b.CFACC#` : `DECIMAL(19,0)` | `a.AccountNo` : `VARCHAR(50)` | `a.AccountNo` → `DECIMAL` | `a.AccountNo = b.CFACC#` |
| IC-000591 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:310](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L310) | JOIN_ON / P01 / COMPARISON | `tms.[CFACC#]` : `DECIMAL(19,0)` | `tmp.AccountNo` : `VARCHAR(50)` | `tmp.AccountNo` → `DECIMAL` | `tmp.AccountNo = tms.[CFACC#]` |
| IC-003579 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCekRekeningTaxAmnesty:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCekRekeningTaxAmnesty.sql#L16) | WHERE / P04 / COMPARISON | `d.CFACC#` : `DECIMAL(19,0)` | `@pcAccountRelasi` : `varchar(100)` | `@pcAccountRelasi` → `DECIMAL` | `d.CFACC# = @pcAccountRelasi` |
| IC-012657 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSUpdateTaxAmnestyExpired:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateTaxAmnestyExpired.sql#L30) | JOIN_ON / P01 / COMPARISON | `b.CFACC#` : `DECIMAL(19,0)` | `a.NoRekInvestor` : `VARCHAR(20)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = b.CFACC#` |
| IC-012662 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSUpdateTaxAmnestyExpired:251](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateTaxAmnestyExpired.sql#L251) | JOIN_ON / P01 / COMPARISON | `b.CFACC#` : `DECIMAL(19,0)` | `a.NoRekInvestor` : `VARCHAR(20)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = b.CFACC#` |
| IC-013615 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_PopulateDataRekeningMaterai:93](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningMaterai.sql#L93) | JOIN_ON / P11 / COMPARISON | `tms.[CFACC#]` : `DECIMAL(19,0)` | `dtc.NoRekInvestor` : `BIGINT` | `dtc.NoRekInvestor` → `DECIMAL` | `tms.[CFACC#] = dtc.NoRekInvestor` |
| IC-013636 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_PopulateDataRekeningObligasi:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningObligasi.sql#L98) | JOIN_ON / P11 / COMPARISON | `tms.[CFACC#]` : `DECIMAL(19,0)` | `dtc.NoRekInvestor` : `bigint` | `dtc.NoRekInvestor` → `DECIMAL` | `tms.[CFACC#] = dtc.NoRekInvestor` |

<a id="kolom-taflag"></a>
## TAFLAG

**Tipe dasar:** `DECIMAL(1,0)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.TAMAST.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 28 lokasi / 16 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000501 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetAccountInfo:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountInfo.sql#L108) | WHERE / P11 / COMPARISON | `tms.TAFLAG` : `DECIMAL(1,0)` | `7` : `int` | `7` → `DECIMAL` | `tms.TAFLAG = 7` |
| IC-000508 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetAccountInfo:173](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountInfo.sql#L173) | WHERE / P11 / COMPARISON | `tms.TAFLAG` : `DECIMAL(1,0)` | `7` : `int` | `7` → `DECIMAL` | `tms.TAFLAG = 7` |
| IC-000516 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetAccountSyariah:58](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountSyariah.sql#L58) | WHERE / P11 / COMPARISON | `tms.TAFLAG` : `DECIMAL(1,0)` | `7` : `int` | `7` → `DECIMAL` | `tms.TAFLAG = 7` |
| IC-000532 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetAccountTaxAmnesty:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountTaxAmnesty.sql#L65) | WHERE / P11 / COMPARISON | `b.TAFLAG` : `DECIMAL(1,0)` | `1` : `int` | `1` → `DECIMAL` | `b.TAFLAG = 1` |
| IC-000537 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetAccountTaxAmnesty:79](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountTaxAmnesty.sql#L79) | WHERE / P11 / COMPARISON | `b.TAFLAG` : `DECIMAL(1,0)` | `1` : `int` | `1` → `DECIMAL` | `b.TAFLAG = 1` |
| IC-000558 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:118](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L118) | WHERE / P11 / COMPARISON | `tms.TAFLAG` : `DECIMAL(1,0)` | `7` : `int` | `7` → `DECIMAL` | `tms.TAFLAG = 7` |
| IC-000561 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:134](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L134) | WHERE / P11 / COMPARISON | `tms.TAFLAG` : `DECIMAL(1,0)` | `7` : `int` | `7` → `DECIMAL` | `tms.TAFLAG = 7` |
| IC-000567 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:161](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L161) | WHERE / P11 / COMPARISON | `d.TAFLAG` : `DECIMAL(1,0)` | `0` : `int` | `0` → `DECIMAL` | `d.TAFLAG = 0` |
| IC-000568 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L162) | WHERE / P11 / COMPARISON | `d.TAFLAG` : `DECIMAL(1,0)` | `1` : `int` | `1` → `DECIMAL` | `d.TAFLAG = 1` |
| IC-000573 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:186](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L186) | WHERE / P11 / COMPARISON | `d.TAFLAG` : `DECIMAL(1,0)` | `0` : `int` | `0` → `DECIMAL` | `d.TAFLAG = 0` |
| IC-000574 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:187](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L187) | WHERE / P11 / COMPARISON | `d.TAFLAG` : `DECIMAL(1,0)` | `1` : `int` | `1` → `DECIMAL` | `d.TAFLAG = 1` |
| IC-000580 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:222](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L222) | WHERE / P11 / COMPARISON | `b.TAFLAG` : `DECIMAL(1,0)` | `1` : `int` | `1` → `DECIMAL` | `b.TAFLAG = 1` |
| IC-000584 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:238](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L238) | WHERE / P11 / COMPARISON | `b.TAFLAG` : `DECIMAL(1,0)` | `1` : `int` | `1` → `DECIMAL` | `b.TAFLAG = 1` |
| IC-000592 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:311](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L311) | WHERE / P11 / COMPARISON | `tms.TAFLAG` : `DECIMAL(1,0)` | `7` : `int` | `7` → `DECIMAL` | `tms.TAFLAG = 7` |
| IC-000759 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.OMTrxSecPreBook:504](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMTrxSecPreBook.sql#L504) | WHERE / P11 / IN_VALUE_TYPE | `TAFLAG` : `DECIMAL(1,0)` | `1` : `int` | `1` → `DECIMAL` | `TAFLAG IN (1,2)` |
| IC-001003 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.ORISBNISubmitEarlyRedemption:262](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNISubmitEarlyRedemption.sql#L262) | WHERE / P11 / IN_VALUE_TYPE | `tm.TAFLAG` : `DECIMAL(1,0)` | `1` : `int` | `1` → `DECIMAL` | `tm.TAFLAG IN (1,2)` |
| IC-001143 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobNasabah:88](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L88) | JOIN_ON / P11 / COMPARISON | `ta.TAFLAG` : `DECIMAL(1,0)` | `7` : `int` | `7` → `DECIMAL` | `ta.TAFLAG = 7` |
| IC-001192 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobNasabah:170](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L170) | JOIN_ON / P11 / COMPARISON | `ta.TAFLAG` : `DECIMAL(1,0)` | `7` : `int` | `7` → `DECIMAL` | `ta.TAFLAG = 7` |
| IC-001763 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.ProTRSReportCOD067:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ProTRSReportCOD067.sql#L113) | WHERE / P11 / IN_VALUE_TYPE | `c.TAFLAG` : `DECIMAL(1,0)` | `1` : `int` | `1` → `DECIMAL` | `c.TAFLAG IN (1)` |
| IC-002288 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSPTOSProcessUpload:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L96) | JOIN_ON / P11 / COMPARISON | `ta.TAFLAG` : `DECIMAL(1,0)` | `7` : `int` | `7` → `DECIMAL` | `ta.TAFLAG = 7` |
| IC-004966 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPTOSPopulateRekeningRDN:12](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPTOSPopulateRekeningRDN.sql#L12) | JOIN_ON / P11 / COMPARISON | `ta.TAFLAG` : `DECIMAL(1,0)` | `7` : `int` | `7` → `DECIMAL` | `ta.TAFLAG = 7` |
| IC-005812 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSUpload:250](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUpload.sql#L250) | JOIN_ON / P11 / COMPARISON | `ta.TAFLAG` : `DECIMAL(1,0)` | `7` : `int` | `7` → `DECIMAL` | `ta.TAFLAG = 7` |
| IC-007111 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessPTOSCustomer_TM:128](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSCustomer_TM.sql#L128) | JOIN_ON / P11 / COMPARISON | `ta.TAFLAG` : `DECIMAL(1,0)` | `7` : `int` | `7` → `DECIMAL` | `ta.TAFLAG = 7` |
| IC-010605 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPTOSNasabah:166](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPTOSNasabah.sql#L166) | JOIN_ON / P11 / COMPARISON | `ta.TAFLAG` : `DECIMAL(1,0)` | `7` : `int` | `7` → `DECIMAL` | `ta.TAFLAG = 7` |
| IC-012658 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSUpdateTaxAmnestyExpired:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateTaxAmnestyExpired.sql#L31) | WHERE / P11 / IN_VALUE_TYPE | `b.TAFLAG` : `DECIMAL(1,0)` | `1` : `int` | `1` → `DECIMAL` | `b.TAFLAG IN (1, 2)` |
| IC-012663 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSUpdateTaxAmnestyExpired:252](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateTaxAmnestyExpired.sql#L252) | WHERE / P11 / IN_VALUE_TYPE | `b.TAFLAG` : `DECIMAL(1,0)` | `1` : `int` | `1` → `DECIMAL` | `b.TAFLAG IN (1, 2)` |
| IC-013616 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_PopulateDataRekeningMaterai:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningMaterai.sql#L94) | WHERE / P11 / COMPARISON | `tms.TAFLAG` : `DECIMAL(1,0)` | `7` : `int` | `7` → `DECIMAL` | `tms.TAFLAG = 7` |
| IC-013637 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_PopulateDataRekeningObligasi:99](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningObligasi.sql#L99) | WHERE / P11 / COMPARISON | `tms.TAFLAG` : `DECIMAL(1,0)` | `7` : `int` | `7` → `DECIMAL` | `tms.TAFLAG = 7` |
