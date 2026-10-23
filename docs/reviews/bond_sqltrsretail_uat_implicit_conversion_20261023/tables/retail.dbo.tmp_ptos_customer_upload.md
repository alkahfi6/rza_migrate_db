# Retail UAT.dbo.tmp_ptos_customer_upload — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 23 lokasi source, 2 procedure, 7 kolom. **A:** 1 lokasi pada 1 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-norekinvestorrdn"></a>
## NoRekInvestorRDN

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.tmp_ptos_customer_upload.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002287 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:95](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L95) | JOIN_ON / P01 / COMPARISON | `tmpCus.NoRekInvestorRDN` : `VARCHAR(20)` | `dt.ACCTNO` : `DECIMAL(19,0)` | `tmpCus.NoRekInvestorRDN` → `DECIMAL` | `tmpCus.NoRekInvestorRDN = dt.ACCTNO` |

<a id="kolom-ccy"></a>
## Ccy

**Tipe dasar:** `VARCHAR(3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.tmp_ptos_customer_upload.sql#L27). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002286 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:93](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L93) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tmpCus.Ccy` : `VARCHAR(3)` | `dt.CURR` : `CHAR(4)` | `dt.CURR` → `VARCHAR(3)` | `tmpCus.Ccy = dt.CURR` |

<a id="kolom-iserror"></a>
## IsError

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.tmp_ptos_customer_upload.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 17 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002283 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:79](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L79) | WHERE / P11 / COMPARISON | `tmp.IsError` : `BIT` | `0` : `int` | `tmp.IsError` → `INT` | `tmp.IsError = 0` |
| IC-002285 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:87](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L87) | WHERE / P11 / COMPARISON | `tmp.IsError` : `BIT` | `0` : `int` | `tmp.IsError` → `INT` | `tmp.IsError = 0` |
| IC-002294 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:116](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L116) | WHERE / P11 / COMPARISON | `IsError` : `BIT` | `0` : `int` | `IsError` → `INT` | `IsError = 0` |
| IC-002296 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:122](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L122) | WHERE / P11 / COMPARISON | `IsError` : `BIT` | `0` : `int` | `IsError` → `INT` | `IsError = 0` |
| IC-002317 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L22) | WHERE / P11 / COMPARISON | `tmp.IsError` : `BIT` | `0` : `int` | `tmp.IsError` → `INT` | `tmp.IsError = 0` |
| IC-002319 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L29) | WHERE / P11 / COMPARISON | `tmp.IsError` : `BIT` | `0` : `int` | `tmp.IsError` → `INT` | `tmp.IsError = 0` |
| IC-002322 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L46) | WHERE / P11 / COMPARISON | `IsError` : `BIT` | `0` : `int` | `IsError` → `INT` | `IsError = 0` |
| IC-002323 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L59) | WHERE / P11 / COMPARISON | `IsError` : `BIT` | `1` : `int` | `IsError` → `INT` | `IsError = 1` |
| IC-002324 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L63) | WHERE / P11 / COMPARISON | `IsError` : `BIT` | `1` : `int` | `IsError` → `INT` | `IsError = 1` |
| IC-002280 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L28) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.tmp_ptos_customer_upload.IsError` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.tmp_ptos_customer_upload.IsError <- 0` |
| IC-002281 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L64) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `t1.IsError` : `BIT` | `1` : `int` | `1` → `BIT` | `t1.IsError = 1` |
| IC-002282 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:75](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L75) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tmp.IsError` : `BIT` | `1` : `int` | `1` → `BIT` | `tmp.IsError = 1` |
| IC-002284 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:83](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L83) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tmp.IsError` : `BIT` | `1` : `int` | `1` → `BIT` | `tmp.IsError = 1` |
| IC-002293 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L114) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `IsError` : `BIT` | `1` : `int` | `1` → `BIT` | `IsError = 1` |
| IC-002295 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:120](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L120) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `IsError` : `BIT` | `1` : `int` | `1` → `BIT` | `IsError = 1` |
| IC-002316 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L18) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tmp.IsError` : `BIT` | `1` : `int` | `1` → `BIT` | `tmp.IsError = 1` |
| IC-002318 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L25) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tmp.IsError` : `BIT` | `1` : `int` | `1` → `BIT` | `tmp.IsError = 1` |

<a id="kolom-name"></a>
## Name

**Tipe dasar:** `VARCHAR(95)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.tmp_ptos_customer_upload.sql#L26). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002289 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L106) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tmpCus.Name` : `VARCHAR(95)` | `dt.CFSNME` : `CHAR(20)` | `dt.CFSNME` → `VARCHAR(95)` | `tmpCus.Name = dt.CFSNME` |

<a id="kolom-noidentitas"></a>
## NoIdentitas

**Tipe dasar:** `VARCHAR(40)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.tmp_ptos_customer_upload.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002291 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSPTOSProcessUpload:109](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L109) | JOIN_ON / P12 / COMPARISON | `tmpCus.NoIdentitas` : `VARCHAR(40)` | `dt.CFSSNO` : `CHAR(40)` | `dt.CFSSNO` → `VARCHAR` | `dt.CFSSNO = tmpCus.NoIdentitas` |

<a id="kolom-ptostype"></a>
## PTOSType

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.tmp_ptos_customer_upload.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002320 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.Rvmp_TRSPTOSSubmitFromTemp:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L33) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `PTOSType` : `INT` | `retail.dbo.ptoscustomer_tm.PTOSType` : `SMALLINT` | `PTOSType` → `SMALLINT` | `retail.dbo.ptoscustomer_tm.PTOSType <- PTOSType` |

<a id="kolom-tgllahirjul"></a>
## TglLahirJul

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.tmp_ptos_customer_upload.sql#L28). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002290 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSPTOSProcessUpload:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L108) | JOIN_ON / P01 / COMPARISON | `tmpCus.TglLahirJul` : `INT` | `dt.CFBIRD` : `CHAR(7)` | `dt.CFBIRD` → `INT` | `dt.CFBIRD = tmpCus.TglLahirJul` |
