# SQL_REPLICATE.dbo.TCLOGT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 20 lokasi source, 6 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 2 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-tcexp7"></a>
## TCEXP7

**Tipe dasar:** `DECIMAL(7,0)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.TCLOGT.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT, COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000106 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.CIFInquiryTaxAmnesty:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CIFInquiryTaxAmnesty.sql#L16) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `TCEXP7` : `DECIMAL(7,0)` | `retail.dbo.fnjuliantodatetime.@nDate` : `int` | `TCEXP7` → `int` | `dbo.fnJulianToDatetime(TCEXP7)` |
| IC-000529 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetAccountTaxAmnesty:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountTaxAmnesty.sql#L63) | JOIN_ON / P11 / COMPARISON | `pcf.TCEXP7` : `DECIMAL(7,0)` | `dbo.fnDatetimeToJulian(getdate())` : `int` | `dbo.fnDatetimeToJulian(getdate())` → `DECIMAL` | `pcf.TCEXP7 > dbo.fnDatetimeToJulian(getdate())` |
| IC-000534 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetAccountTaxAmnesty:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountTaxAmnesty.sql#L77) | JOIN_ON / P11 / COMPARISON | `pcf.TCEXP7` : `DECIMAL(7,0)` | `dbo.fnDatetimeToJulian(getdate())` : `int` | `dbo.fnDatetimeToJulian(getdate())` → `DECIMAL` | `pcf.TCEXP7 > dbo.fnDatetimeToJulian(getdate())` |
| IC-000566 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:159](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L159) | JOIN_ON / P11 / COMPARISON | `e.TCEXP7` : `DECIMAL(7,0)` | `dbo.fnDatetimeToJulian(getdate())` : `int` | `dbo.fnDatetimeToJulian(getdate())` → `DECIMAL` | `e.TCEXP7 > dbo.fnDatetimeToJulian(getdate())` |
| IC-000572 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:184](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L184) | JOIN_ON / P11 / COMPARISON | `e.TCEXP7` : `DECIMAL(7,0)` | `dbo.fnDatetimeToJulian(getdate())` : `int` | `dbo.fnDatetimeToJulian(getdate())` → `DECIMAL` | `e.TCEXP7 > dbo.fnDatetimeToJulian(getdate())` |
| IC-000579 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:219](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L219) | JOIN_ON / P11 / COMPARISON | `pcf.TCEXP7` : `DECIMAL(7,0)` | `dbo.fnDatetimeToJulian(getdate())` : `int` | `dbo.fnDatetimeToJulian(getdate())` → `DECIMAL` | `pcf.TCEXP7 > dbo.fnDatetimeToJulian(getdate())` |
| IC-000583 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:235](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L235) | JOIN_ON / P11 / COMPARISON | `pcf.TCEXP7` : `DECIMAL(7,0)` | `dbo.fnDatetimeToJulian(getdate())` : `int` | `dbo.fnDatetimeToJulian(getdate())` → `DECIMAL` | `pcf.TCEXP7 > dbo.fnDatetimeToJulian(getdate())` |
| IC-013060 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_GetRekTaxAmnesty:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_GetRekTaxAmnesty.sql#L18) | JOIN_ON / P11 / COMPARISON | `tl.TCEXP7` : `DECIMAL(7,0)` | `dbo.fnDatetimeToJulian(getdate())` : `int` | `dbo.fnDatetimeToJulian(getdate())` → `DECIMAL` | `tl.TCEXP7 > dbo.fnDatetimeToJulian(getdate())` |

<a id="kolom-tcflag"></a>
## TCFLAG

**Tipe dasar:** `DECIMAL(1,0)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.TCLOGT.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 11 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004920 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSObligasiGetTaxAmnestyTypeBySecAccNo:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSObligasiGetTaxAmnestyTypeBySecAccNo.sql#L17) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `ISNULL(TCFLAG,0)` : `DECIMAL(1,0)` | `@nTaxAmnestyType` : `INT` | `ISNULL(TCFLAG,0)` → `INT` | `@nTaxAmnestyType = ISNULL(TCFLAG,0)` |
| IC-000107 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.CIFInquiryTaxAmnesty:19](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CIFInquiryTaxAmnesty.sql#L19) | WHERE / P11 / IN_VALUE_TYPE | `TCFLAG` : `DECIMAL(1,0)` | `1` : `int` | `1` → `DECIMAL` | `TCFLAG IN (1, 2)` |
| IC-000533 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetAccountTaxAmnesty:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountTaxAmnesty.sql#L65) | WHERE / P11 / COMPARISON | `pcf.TCFLAG` : `DECIMAL(1,0)` | `1` : `int` | `1` → `DECIMAL` | `pcf.TCFLAG = 1` |
| IC-000538 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetAccountTaxAmnesty:79](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountTaxAmnesty.sql#L79) | WHERE / P11 / COMPARISON | `pcf.TCFLAG` : `DECIMAL(1,0)` | `1` : `int` | `1` → `DECIMAL` | `pcf.TCFLAG = 1` |
| IC-000565 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:158](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L158) | JOIN_ON / P11 / COMPARISON | `e.TCFLAG` : `DECIMAL(1,0)` | `1` : `int` | `1` → `DECIMAL` | `e.TCFLAG = 1` |
| IC-000571 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:183](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L183) | JOIN_ON / P11 / COMPARISON | `e.TCFLAG` : `DECIMAL(1,0)` | `1` : `int` | `1` → `DECIMAL` | `e.TCFLAG = 1` |
| IC-000581 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:222](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L222) | WHERE / P11 / COMPARISON | `pcf.TCFLAG` : `DECIMAL(1,0)` | `1` : `int` | `1` → `DECIMAL` | `pcf.TCFLAG = 1` |
| IC-000585 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:238](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L238) | WHERE / P11 / COMPARISON | `pcf.TCFLAG` : `DECIMAL(1,0)` | `1` : `int` | `1` → `DECIMAL` | `pcf.TCFLAG = 1` |
| IC-003578 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCekRekeningTaxAmnesty:14](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCekRekeningTaxAmnesty.sql#L14) | JOIN_ON / P11 / COMPARISON | `e.TCFLAG` : `DECIMAL(1,0)` | `1` : `int` | `1` → `DECIMAL` | `e.TCFLAG = 1` |
| IC-004919 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSObligasiGetTaxAmnestyTypeBySecAccNo:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSObligasiGetTaxAmnestyTypeBySecAccNo.sql#L17) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TCFLAG` : `DECIMAL(1,0)` | `0` : `int` | `0` → `DECIMAL(1,0)` | `ISNULL(TCFLAG,0)` |
| IC-004921 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSObligasiGetTaxAmnestyTypeBySecAccNo:20](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSObligasiGetTaxAmnestyTypeBySecAccNo.sql#L20) | WHERE / P11 / IN_VALUE_TYPE | `TCFLAG` : `DECIMAL(1,0)` | `1` : `int` | `1` → `DECIMAL` | `TCFLAG in (1,2)` |

<a id="kolom-tclcif"></a>
## TCLCIF

**Tipe dasar:** `DECIMAL(19,0)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.TCLOGT.sql#L2). **Existing leading index:** idx_tclcif_tcflag.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013059 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_GetRekTaxAmnesty:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_GetRekTaxAmnesty.sql#L17) | JOIN_ON / P11 / COMPARISON | `tl.TCLCIF` : `DECIMAL(19,0)` | `cast(tctm.CIFNo as bigint)` : `bigint` | `cast(tctm.CIFNo as bigint)` → `DECIMAL` | `cast(tctm.CIFNo as bigint) = tl.TCLCIF` |
