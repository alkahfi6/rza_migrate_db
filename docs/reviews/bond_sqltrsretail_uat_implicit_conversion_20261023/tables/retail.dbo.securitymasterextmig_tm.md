# Retail UAT.dbo.SecurityMasterExtMig_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 26 lokasi source, 6 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-secid"></a>
## SecId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterExtMig_TM.sql#L3). **Existing leading index:** SecurityMasterExtMig_TM_CN1.

**Frekuensi source:** 13 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008858 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:954](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L954) | JOIN_ON / P11 / COMPARISON | `ext.SecId` : `BIGINT` | `tmp.SecId` : `int` | `tmp.SecId` → `BIGINT` | `tmp.SecId = ext.SecId` |
| IC-008878 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:1029](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L1029) | JOIN_ON / P11 / COMPARISON | `ext.SecId` : `BIGINT` | `tmp.SecId` : `int` | `tmp.SecId` → `BIGINT` | `tmp.SecId = ext.SecId` |
| IC-008895 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:1150](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L1150) | JOIN_ON / P11 / COMPARISON | `ext.SecId` : `BIGINT` | `tmp.SecId` : `int` | `tmp.SecId` → `BIGINT` | `tmp.SecId = ext.SecId` |
| IC-010769 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORI:284](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORI.sql#L284) | JOIN_ON / P11 / COMPARISON | `ext.SecId` : `BIGINT` | `tmp._SecId` : `int` | `tmp._SecId` → `BIGINT` | `tmp._SecId = ext.SecId` |
| IC-010772 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORI:297](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORI.sql#L297) | JOIN_ON / P11 / COMPARISON | `ext.SecId` : `BIGINT` | `tmp._SecId` : `int` | `tmp._SecId` → `BIGINT` | `tmp._SecId = ext.SecId` |
| IC-010880 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORIDetail:662](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L662) | JOIN_ON / P11 / COMPARISON | `ext.SecId` : `BIGINT` | `tmp._SecId` : `int` | `tmp._SecId` → `BIGINT` | `tmp._SecId = ext.SecId` |
| IC-010883 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORIDetail:676](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L676) | JOIN_ON / P11 / COMPARISON | `ext.SecId` : `BIGINT` | `tmp._SecId` : `int` | `tmp._SecId` → `BIGINT` | `tmp._SecId = ext.SecId` |
| IC-011010 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnesty:586](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L586) | JOIN_ON / P11 / COMPARISON | `ext.SecId` : `BIGINT` | `tmp._SecId` : `int` | `tmp._SecId` → `BIGINT` | `tmp._SecId = ext.SecId` |
| IC-011013 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnesty:600](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L600) | JOIN_ON / P11 / COMPARISON | `ext.SecId` : `BIGINT` | `tmp._SecId` : `int` | `tmp._SecId` → `BIGINT` | `tmp._SecId = ext.SecId` |
| IC-011221 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:470](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L470) | JOIN_ON / P11 / COMPARISON | `ext.SecId` : `BIGINT` | `tmp._SecId` : `INT` | `tmp._SecId` → `BIGINT` | `tmp._SecId = ext.SecId` |
| IC-011224 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:483](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L483) | JOIN_ON / P11 / COMPARISON | `ext.SecId` : `BIGINT` | `tmp._SecId` : `INT` | `tmp._SecId` → `BIGINT` | `tmp._SecId = ext.SecId` |
| IC-011421 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnesty_Backdated:407](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty_Backdated.sql#L407) | JOIN_ON / P11 / COMPARISON | `ext.SecId` : `BIGINT` | `tmp._SecId` : `int` | `tmp._SecId` → `BIGINT` | `tmp._SecId = ext.SecId` |
| IC-011424 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPayCouponORITaxAmnesty_Backdated:421](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty_Backdated.sql#L421) | JOIN_ON / P11 / COMPARISON | `ext.SecId` : `BIGINT` | `tmp._SecId` : `int` | `tmp._SecId` → `BIGINT` | `tmp._SecId = ext.SecId` |

<a id="kolom-waiveskf"></a>
## WaiveSKF

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterExtMig_TM.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 13 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008859 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailCalculateFee:959](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L959) | WHERE / P11 / COMPARISON | `ext.WaiveSKF` : `BIT` | `1` : `int` | `ext.WaiveSKF` → `INT` | `ext.WaiveSKF = 1` |
| IC-008879 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailCalculateFee:1034](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L1034) | WHERE / P11 / COMPARISON | `ext.WaiveSKF` : `BIT` | `1` : `int` | `ext.WaiveSKF` → `INT` | `ext.WaiveSKF = 1` |
| IC-008896 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailCalculateFee:1155](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L1155) | WHERE / P11 / COMPARISON | `ext.WaiveSKF` : `BIT` | `1` : `int` | `ext.WaiveSKF` → `INT` | `ext.WaiveSKF = 1` |
| IC-010770 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORI:289](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORI.sql#L289) | WHERE / P11 / COMPARISON | `ext.WaiveSKF` : `BIT` | `1` : `int` | `ext.WaiveSKF` → `INT` | `ext.WaiveSKF = 1` |
| IC-010773 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORI:302](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORI.sql#L302) | WHERE / P11 / COMPARISON | `ext.WaiveSKF` : `BIT` | `1` : `int` | `ext.WaiveSKF` → `INT` | `ext.WaiveSKF = 1` |
| IC-010881 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORIDetail:667](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L667) | WHERE / P11 / COMPARISON | `ext.WaiveSKF` : `BIT` | `1` : `int` | `ext.WaiveSKF` → `INT` | `ext.WaiveSKF = 1` |
| IC-010884 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORIDetail:681](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L681) | WHERE / P11 / COMPARISON | `ext.WaiveSKF` : `BIT` | `1` : `int` | `ext.WaiveSKF` → `INT` | `ext.WaiveSKF = 1` |
| IC-011011 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnesty:591](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L591) | WHERE / P11 / COMPARISON | `ext.WaiveSKF` : `BIT` | `1` : `int` | `ext.WaiveSKF` → `INT` | `ext.WaiveSKF = 1` |
| IC-011014 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnesty:605](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L605) | WHERE / P11 / COMPARISON | `ext.WaiveSKF` : `BIT` | `1` : `int` | `ext.WaiveSKF` → `INT` | `ext.WaiveSKF = 1` |
| IC-011222 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:475](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L475) | WHERE / P11 / COMPARISON | `ext.WaiveSKF` : `BIT` | `1` : `int` | `ext.WaiveSKF` → `INT` | `ext.WaiveSKF = 1` |
| IC-011225 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:488](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L488) | WHERE / P11 / COMPARISON | `ext.WaiveSKF` : `BIT` | `1` : `int` | `ext.WaiveSKF` → `INT` | `ext.WaiveSKF = 1` |
| IC-011422 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnesty_Backdated:412](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty_Backdated.sql#L412) | WHERE / P11 / COMPARISON | `ext.WaiveSKF` : `BIT` | `1` : `int` | `ext.WaiveSKF` → `INT` | `ext.WaiveSKF = 1` |
| IC-011425 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnesty_Backdated:426](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty_Backdated.sql#L426) | WHERE / P11 / COMPARISON | `ext.WaiveSKF` : `BIT` | `1` : `int` | `ext.WaiveSKF` → `INT` | `ext.WaiveSKF = 1` |
