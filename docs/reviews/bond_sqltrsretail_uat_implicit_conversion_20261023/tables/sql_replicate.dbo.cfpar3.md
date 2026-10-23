# SQL_REPLICATE.dbo.CFPAR3 — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 13 lokasi source, 8 procedure, 2 kolom. **A:** 10 lokasi pada 7 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cp3uic"></a>
## CP3UIC

**Tipe dasar:** `VARCHAR(1)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFPAR3.sql#L3). **Existing leading index:** idx_cp3uic_cp3ucd.

**Frekuensi source:** 10 lokasi / 7 SP. **A:** 10 lokasi / 7 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003724 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:405](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L405) | JOIN_ON / P01 / COMPARISON | `CP3UIC` : `VARCHAR(1)` | `3` : `int` | `CP3UIC` → `INT` | `CP3UIC = 3` |
| IC-003958 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusiOrchid:335](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L335) | JOIN_ON / P01 / COMPARISON | `CP3UIC` : `VARCHAR(1)` | `3` : `int` | `CP3UIC` → `INT` | `CP3UIC = 3` |
| IC-005994 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateRiskProfileNasabah:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateRiskProfileNasabah.sql#L21) | WHERE / P01 / COMPARISON | `CP3UIC` : `VARCHAR(1)` | `8` : `int` | `CP3UIC` → `INT` | `CP3UIC = 8` |
| IC-009829 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptDetailTrx:848](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L848) | WHERE / P01 / COMPARISON | `par.CP3UIC` : `VARCHAR(1)` | `8` : `int` | `par.CP3UIC` → `INT` | `par.CP3UIC = 8` |
| IC-009830 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptDetailTrx:855](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L855) | WHERE / P01 / COMPARISON | `par.CP3UIC` : `VARCHAR(1)` | `8` : `int` | `par.CP3UIC` → `INT` | `par.CP3UIC = 8` |
| IC-009919 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptDetailTrxTaxAmnesty:774](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L774) | WHERE / P01 / COMPARISON | `par.CP3UIC` : `VARCHAR(1)` | `8` : `int` | `par.CP3UIC` → `INT` | `par.CP3UIC = 8` |
| IC-009920 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptDetailTrxTaxAmnesty:781](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L781) | WHERE / P01 / COMPARISON | `par.CP3UIC` : `VARCHAR(1)` | `8` : `int` | `par.CP3UIC` → `INT` | `par.CP3UIC = 8` |
| IC-011630 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptTransactionDetailWMTA:788](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L788) | WHERE / P01 / COMPARISON | `par.CP3UIC` : `VARCHAR(1)` | `8` : `int` | `par.CP3UIC` → `INT` | `par.CP3UIC = 8` |
| IC-011631 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptTransactionDetailWMTA:795](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L795) | WHERE / P01 / COMPARISON | `par.CP3UIC` : `VARCHAR(1)` | `8` : `int` | `par.CP3UIC` → `INT` | `par.CP3UIC = 8` |
| IC-012418 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSSyncMasterRiskProfileData:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterRiskProfileData.sql#L25) | WHERE / P01 / COMPARISON | `CFPAR3.CP3UIC` : `VARCHAR(1)` | `8` : `int` | `CFPAR3.CP3UIC` → `INT` | `CFPAR3.CP3UIC = 8` |

<a id="kolom-cp3ucd"></a>
## CP3UCD

**Tipe dasar:** `VARCHAR(1)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.CFPAR3.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005993 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateRiskProfileNasabah:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateRiskProfileNasabah.sql#L21) | WHERE / P12 / COMPARISON | `CP3UCD` : `VARCHAR(1)` | `@cRiskProfile` : `char(1)` | `@cRiskProfile` → `VARCHAR` | `CP3UCD = @cRiskProfile` |
| IC-012416 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSyncMasterRiskProfileData:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterRiskProfileData.sql#L23) | JOIN_ON / P12 / COMPARISON | `CFPAR3.CP3UCD` : `VARCHAR(1)` | `CFMAST.CFUIC8` : `CHAR(1)` | `CFMAST.CFUIC8` → `VARCHAR` | `CFMAST.CFUIC8 = CFPAR3.CP3UCD` |
| IC-012482 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSUpdateDataMasterNasabah:149](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateDataMasterNasabah.sql#L149) | JOIN_ON / P12 / COMPARISON | `par.CP3UCD` : `VARCHAR(1)` | `t.newRiskProfileCode` : `CHAR(1)` | `t.newRiskProfileCode` → `VARCHAR` | `t.newRiskProfileCode = par.CP3UCD` |
