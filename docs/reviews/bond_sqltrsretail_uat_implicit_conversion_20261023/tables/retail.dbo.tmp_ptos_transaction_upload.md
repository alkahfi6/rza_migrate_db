# Retail UAT.dbo.tmp_ptos_transaction_upload — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 23 lokasi source, 2 procedure, 5 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 2 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-accruedinterest"></a>
## AccruedInterest

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.tmp_ptos_transaction_upload.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002305 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:251](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L251) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.tmp_ptos_transaction_upload.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.tmp_ptos_transaction_upload.AccruedInterest <- 0` |

<a id="kolom-iserror"></a>
## IsError

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.tmp_ptos_transaction_upload.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 14 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002313 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:298](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L298) | WHERE / P11 / COMPARISON | `tmp.IsError` : `BIT` | `0` : `int` | `tmp.IsError` → `INT` | `tmp.IsError = 0` |
| IC-002315 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:304](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L304) | WHERE / P11 / COMPARISON | `IsError` : `BIT` | `0` : `int` | `IsError` → `INT` | `IsError = 0` |
| IC-002354 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:163](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L163) | WHERE / P11 / COMPARISON | `tmp.IsError` : `BIT` | `0` : `int` | `tmp.IsError` → `INT` | `tmp.IsError = 0` |
| IC-002358 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:171](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L171) | WHERE / P11 / COMPARISON | `tmp.IsError` : `BIT` | `0` : `int` | `tmp.IsError` → `INT` | `tmp.IsError = 0` |
| IC-002362 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:187](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L187) | WHERE / P11 / COMPARISON | `tmp.IsError` : `BIT` | `0` : `int` | `tmp.IsError` → `INT` | `tmp.IsError = 0` |
| IC-002366 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:203](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L203) | WHERE / P11 / COMPARISON | `tmp.IsError` : `BIT` | `0` : `int` | `tmp.IsError` → `INT` | `tmp.IsError = 0` |
| IC-002367 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:217](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L217) | WHERE / P11 / COMPARISON | `IsError` : `BIT` | `1` : `int` | `IsError` → `INT` | `IsError = 1` |
| IC-002368 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:221](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L221) | WHERE / P11 / COMPARISON | `IsError` : `BIT` | `1` : `int` | `IsError` → `INT` | `IsError = 1` |
| IC-002304 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:217](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L217) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.tmp_ptos_transaction_upload.IsError` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.tmp_ptos_transaction_upload.IsError <- 0` |
| IC-002306 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:251](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L251) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.tmp_ptos_transaction_upload.IsError` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.tmp_ptos_transaction_upload.IsError <- 0` |
| IC-002309 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:285](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L285) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `t1.IsError` : `BIT` | `1` : `int` | `1` → `BIT` | `t1.IsError = 1` |
| IC-002310 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:294](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L294) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tmp.IsError` : `BIT` | `1` : `int` | `1` → `BIT` | `tmp.IsError = 1` |
| IC-002314 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:302](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L302) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `IsError` : `BIT` | `1` : `int` | `1` → `BIT` | `IsError = 1` |
| IC-002351 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:159](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L159) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tmp.IsError` : `BIT` | `1` : `int` | `1` → `BIT` | `tmp.IsError = 1` |

<a id="kolom-ptostype"></a>
## PTOSType

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.tmp_ptos_transaction_upload.sql#L28). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002359 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.Rvmp_TRSPTOSSubmitFromTemp:173](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L173) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tmp.PTOSType` : `INT` | `retail.dbo.ptostransaksi_temp.PTOSType` : `SMALLINT` | `tmp.PTOSType` → `SMALLINT` | `retail.dbo.ptostransaksi_temp.PTOSType <- tmp.PTOSType` |
| IC-002363 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.Rvmp_TRSPTOSSubmitFromTemp:189](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L189) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tmp.PTOSType` : `INT` | `retail.dbo.ptostransaksi_tr.PTOSType` : `SMALLINT` | `tmp.PTOSType` → `SMALLINT` | `retail.dbo.ptostransaksi_tr.PTOSType <- tmp.PTOSType` |
| IC-002352 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSPTOSSubmitFromTemp:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L162) | JOIN_ON / P11 / COMPARISON | `tmp.PTOSType` : `INT` | `pttr.PTOSType` : `SMALLINT` | `pttr.PTOSType` → `INT` | `tmp.PTOSType = pttr.PTOSType` |
| IC-002357 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSPTOSSubmitFromTemp:170](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L170) | JOIN_ON / P11 / COMPARISON | `tmp.PTOSType` : `INT` | `pttr.PTOSType` : `SMALLINT` | `pttr.PTOSType` → `INT` | `pttr.PTOSType = tmp.PTOSType` |
| IC-002361 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSPTOSSubmitFromTemp:186](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L186) | JOIN_ON / P11 / COMPARISON | `tmp.PTOSType` : `INT` | `pttr.PTOSType` : `SMALLINT` | `pttr.PTOSType` → `INT` | `tmp.PTOSType = pttr.PTOSType` |
| IC-002365 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSPTOSSubmitFromTemp:202](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L202) | JOIN_ON / P11 / COMPARISON | `tmp.PTOSType` : `INT` | `pttr.PTOSType` : `SMALLINT` | `pttr.PTOSType` → `INT` | `tmp.PTOSType = pttr.PTOSType` |

<a id="kolom-taxcapitalgainloss"></a>
## TaxCapitalGainLoss

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.tmp_ptos_transaction_upload.sql#L26). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002307 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:251](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L251) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.tmp_ptos_transaction_upload.TaxCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.tmp_ptos_transaction_upload.TaxCapitalGainLoss <- 0` |

<a id="kolom-taxonaccrued"></a>
## TaxOnAccrued

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.tmp_ptos_transaction_upload.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002308 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:251](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L251) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.tmp_ptos_transaction_upload.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.tmp_ptos_transaction_upload.TaxOnAccrued <- 0` |
