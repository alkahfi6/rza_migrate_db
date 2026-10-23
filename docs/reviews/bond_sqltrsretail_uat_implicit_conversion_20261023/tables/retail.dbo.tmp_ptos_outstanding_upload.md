# Retail UAT.dbo.tmp_ptos_outstanding_upload — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 30 lokasi source, 4 procedure, 5 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 2 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-accruedinterest"></a>
## AccruedInterest

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.tmp_ptos_outstanding_upload.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002298 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L168) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.tmp_ptos_outstanding_upload.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.tmp_ptos_outstanding_upload.AccruedInterest <- 0` |

<a id="kolom-hargaperolehannasabah"></a>
## HargaPerolehanNasabah

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.tmp_ptos_outstanding_upload.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002328 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPTOSSubmitFromTemp:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L81) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tmp.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(tmp.HargaPerolehanNasabah, 0)` |
| IC-002335 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPTOSSubmitFromTemp:93](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L93) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tmp.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(tmp.HargaPerolehanNasabah, 0)` |
| IC-002341 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPTOSSubmitFromTemp:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L111) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tmp.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(tmp.HargaPerolehanNasabah, 0)` |
| IC-002347 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPTOSSubmitFromTemp:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L130) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tmp.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(tmp.HargaPerolehanNasabah, 0)` |

<a id="kolom-indikasihargabid"></a>
## IndikasiHargaBid

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.tmp_ptos_outstanding_upload.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002303 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSPTOSProcessUpload:204](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L204) | WHERE / P11 / COMPARISON | `tmp.IndikasiHargaBid` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `tmp.IndikasiHargaBid = 0` |

<a id="kolom-iserror"></a>
## IsError

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.tmp_ptos_outstanding_upload.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 18 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001993 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_PopulatePTOSUploadOutstandingCorporate:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_PopulatePTOSUploadOutstandingCorporate.sql#L43) | WHERE / P11 / COMPARISON | `IsError` : `BIT` | `0` : `int` | `IsError` → `INT` | `IsError = 0` |
| IC-001994 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_PopulatePTOSUploadOutstandingCorporate:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_PopulatePTOSUploadOutstandingCorporate.sql#L52) | WHERE / P11 / COMPARISON | `IsError` : `BIT` | `0` : `int` | `IsError` → `INT` | `IsError = 0` |
| IC-001995 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_PopulatePTOSUploadOutstandingCorporate:72](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_PopulatePTOSUploadOutstandingCorporate.sql#L72) | WHERE / P11 / COMPARISON | `IsError` : `BIT` | `1` : `int` | `IsError` → `INT` | `IsError = 1` |
| IC-001996 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_PopulatePTOSUploadOutstandingCorporate:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_PopulatePTOSUploadOutstandingCorporate.sql#L81) | WHERE / P11 / COMPARISON | `IsError` : `BIT` | `1` : `int` | `IsError` → `INT` | `IsError = 1` |
| IC-001997 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_PopulatePTOSUploadOutstandingEquity:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_PopulatePTOSUploadOutstandingEquity.sql#L43) | WHERE / P11 / COMPARISON | `IsError` : `BIT` | `0` : `int` | `IsError` → `INT` | `IsError = 0` |
| IC-001998 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_PopulatePTOSUploadOutstandingEquity:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_PopulatePTOSUploadOutstandingEquity.sql#L52) | WHERE / P11 / COMPARISON | `IsError` : `BIT` | `0` : `int` | `IsError` → `INT` | `IsError = 0` |
| IC-001999 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_PopulatePTOSUploadOutstandingEquity:72](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_PopulatePTOSUploadOutstandingEquity.sql#L72) | WHERE / P11 / COMPARISON | `IsError` : `BIT` | `1` : `int` | `IsError` → `INT` | `IsError = 1` |
| IC-002000 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_PopulatePTOSUploadOutstandingEquity:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_PopulatePTOSUploadOutstandingEquity.sql#L81) | WHERE / P11 / COMPARISON | `IsError` : `BIT` | `1` : `int` | `IsError` → `INT` | `IsError = 1` |
| IC-002330 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:82](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L82) | WHERE / P11 / COMPARISON | `tmp.IsError` : `BIT` | `0` : `int` | `tmp.IsError` → `INT` | `tmp.IsError = 0` |
| IC-002336 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L94) | WHERE / P11 / COMPARISON | `tmp.IsError` : `BIT` | `0` : `int` | `tmp.IsError` → `INT` | `tmp.IsError = 0` |
| IC-002342 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:112](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L112) | WHERE / P11 / COMPARISON | `tmp.IsError` : `BIT` | `0` : `int` | `tmp.IsError` → `INT` | `tmp.IsError = 0` |
| IC-002348 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:132](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L132) | WHERE / P11 / COMPARISON | `tmp.IsError` : `BIT` | `0` : `int` | `tmp.IsError` → `INT` | `tmp.IsError = 0` |
| IC-002349 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:144](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L144) | WHERE / P11 / COMPARISON | `IsError` : `BIT` | `1` : `int` | `IsError` → `INT` | `IsError = 1` |
| IC-002350 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:148](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L148) | WHERE / P11 / COMPARISON | `IsError` : `BIT` | `1` : `int` | `IsError` → `INT` | `IsError = 1` |
| IC-002297 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.tmp_ptos_outstanding_upload.IsError` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.tmp_ptos_outstanding_upload.IsError <- 0` |
| IC-002299 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L168) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.tmp_ptos_outstanding_upload.IsError` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.tmp_ptos_outstanding_upload.IsError <- 0` |
| IC-002300 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:191](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L191) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tmp.IsError` : `BIT` | `1` : `int` | `1` → `BIT` | `tmp.IsError = 1` |
| IC-002325 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:74](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L74) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tmp.IsError` : `BIT` | `1` : `int` | `1` → `BIT` | `tmp.IsError = 1` |

<a id="kolom-ptostype"></a>
## PTOSType

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.tmp_ptos_outstanding_upload.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002337 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.Rvmp_TRSPTOSSubmitFromTemp:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L96) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tmp.PTOSType` : `INT` | `retail.dbo.ptosoutstanding_temp.PTOSType` : `SMALLINT` | `tmp.PTOSType` → `SMALLINT` | `retail.dbo.ptosoutstanding_temp.PTOSType <- tmp.PTOSType` |
| IC-002343 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.Rvmp_TRSPTOSSubmitFromTemp:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tmp.PTOSType` : `INT` | `retail.dbo.ptosoutstanding_tr.PTOSType` : `SMALLINT` | `tmp.PTOSType` → `SMALLINT` | `retail.dbo.ptosoutstanding_tr.PTOSType <- tmp.PTOSType` |
| IC-002326 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSPTOSSubmitFromTemp:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L80) | JOIN_ON / P11 / COMPARISON | `tmp.PTOSType` : `INT` | `potr.PTOSType` : `SMALLINT` | `potr.PTOSType` → `INT` | `tmp.PTOSType = potr.PTOSType` |
| IC-002333 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSPTOSSubmitFromTemp:92](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L92) | JOIN_ON / P11 / COMPARISON | `tmp.PTOSType` : `INT` | `potr.PTOSType` : `SMALLINT` | `potr.PTOSType` → `INT` | `potr.PTOSType = tmp.PTOSType` |
| IC-002339 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSPTOSSubmitFromTemp:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L110) | JOIN_ON / P11 / COMPARISON | `tmp.PTOSType` : `INT` | `potr.PTOSType` : `SMALLINT` | `potr.PTOSType` → `INT` | `tmp.PTOSType = potr.PTOSType` |
| IC-002345 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSPTOSSubmitFromTemp:129](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L129) | JOIN_ON / P11 / COMPARISON | `tmp.PTOSType` : `INT` | `potr.PTOSType` : `SMALLINT` | `potr.PTOSType` → `INT` | `tmp.PTOSType = potr.PTOSType` |
