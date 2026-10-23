# Retail UAT.dbo.SecurityTrxLelang_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 19 lokasi source, 8 procedure, 7 kolom. **A:** 1 lokasi pada 1 procedure. **B:** 8 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-trxbranch"></a>
## TrxBranch

**Tipe dasar:** `VARCHAR(5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTrxLelang_TT.sql#L25). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005325 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateBlotterTransaksi:177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateBlotterTransaksi.sql#L177) | JOIN_ON / P01 / COMPARISON | `stl.TrxBranch` : `VARCHAR(5)` | `mie.office_id` : `INT` | `stl.TrxBranch` → `INT` | `stl.TrxBranch = mie.office_id` |

<a id="kolom-accountblockactype"></a>
## AccountBlockACTYPE

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTrxLelang_TT.sql#L51). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003100 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SecurityProcessLelangToTransaction:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessLelangToTransaction.sql#L70) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tt.AccountBlockACTYPE` : `CHAR(1)` | `AccountBlockACTYPE` : `VARCHAR(5)` | `tt.AccountBlockACTYPE` → `VARCHAR(5)` | `AccountBlockACTYPE = tt.AccountBlockACTYPE` |
| IC-003102 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SecurityProcessLelangToTransaction:153](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessLelangToTransaction.sql#L153) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `AccountBlockACTYPE` : `CHAR(1)` | `tmp.AccountBlockACTYPE` : `VARCHAR(5)` | `tmp.AccountBlockACTYPE` → `CHAR(1)` | `AccountBlockACTYPE = tmp.AccountBlockACTYPE` |

<a id="kolom-facevalue"></a>
## FaceValue

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTrxLelang_TT.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003390 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSBlotterLelangSummaryPopulate:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBlotterLelangSummaryPopulate.sql#L37) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `isnull(trx.FaceValue,0)` : `MONEY` | `#tmp_data.Nominal` : `decimal(20,4)` | `isnull(trx.FaceValue,0)` → `decimal(20,4)` | `#tmp_data.Nominal <- isnull(trx.FaceValue,0)` |
| IC-003394 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSBlotterLelangSummaryPopulate:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBlotterLelangSummaryPopulate.sql#L59) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `isnull(trx.FaceValue,0)` : `MONEY` | `#tmp_data.Nominal` : `decimal(20,4)` | `isnull(trx.FaceValue,0)` → `decimal(20,4)` | `#tmp_data.Nominal <- isnull(trx.FaceValue,0)` |
| IC-003392 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSBlotterLelangSummaryPopulate:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBlotterLelangSummaryPopulate.sql#L45) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `trx.FaceValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(trx.FaceValue,0)` |
| IC-003396 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSBlotterLelangSummaryPopulate:67](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBlotterLelangSummaryPopulate.sql#L67) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `trx.FaceValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(trx.FaceValue,0)` |

<a id="kolom-flagother"></a>
## FlagOther

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTrxLelang_TT.sql#L40). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003034 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SecurityPopulateDetailReplaceTrxLelang:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateDetailReplaceTrxLelang.sql#L28) | CASE / P11 / COMPARISON | `FlagOther` : `BIT` | `0` : `int` | `FlagOther` → `INT` | `FlagOther = 0` |
| IC-003093 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SecurityPopulateTrxLelang:75](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateTrxLelang.sql#L75) | CASE / P11 / COMPARISON | `FlagOther` : `BIT` | `0` : `int` | `FlagOther` → `INT` | `FlagOther = 0` |

<a id="kolom-hargamodal"></a>
## HargaModal

**Tipe dasar:** `DECIMAL(19,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTrxLelang_TT.sql#L37). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003391 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSBlotterLelangSummaryPopulate:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBlotterLelangSummaryPopulate.sql#L44) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `trx.HargaModal` : `DECIMAL(19,5)` | `0` : `int` | `0` → `DECIMAL(19,5)` | `isnull(trx.HargaModal,0)` |
| IC-003395 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSBlotterLelangSummaryPopulate:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBlotterLelangSummaryPopulate.sql#L66) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `trx.HargaModal` : `DECIMAL(19,5)` | `0` : `int` | `0` → `DECIMAL(19,5)` | `isnull(trx.HargaModal,0)` |

<a id="kolom-insertedby"></a>
## InsertedBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTrxLelang_TT.sql#L26). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003047 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SecurityPopulateResultLelangEmail:208](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateResultLelangEmail.sql#L208) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `InsertedBy` : `INT` | `@nInsertedNIK` : `VARCHAR(11)` | `InsertedBy` → `VARCHAR(11)` | `@nInsertedNIK = InsertedBy` |
| IC-005069 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSParamLelangProcessApproval:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangProcessApproval.sql#L193) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `trx.InsertedBy` : `INT` | `retail.dbo.fngetnikspv.@pcNIK` : `varchar(11)` | `trx.InsertedBy` → `varchar(11)` | `dbo.fnGetNIKSpv(trx.InsertedBy)` |
| IC-005070 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSParamLelangProcessApproval:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangProcessApproval.sql#L216) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `trx.InsertedBy` : `INT` | `retail.dbo.fngetnikspv.@pcNIK` : `varchar(11)` | `trx.InsertedBy` → `varchar(11)` | `dbo.fnGetNIKSpv(trx.InsertedBy)` |
| IC-005071 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSParamLelangProcessApproval:242](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangProcessApproval.sql#L242) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `trx.InsertedBy` : `INT` | `retail.dbo.fngetnikspv.@pcNIK` : `varchar(11)` | `trx.InsertedBy` → `varchar(11)` | `dbo.fnGetNIKSpv(trx.InsertedBy)` |

<a id="kolom-secid"></a>
## SecId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTrxLelang_TT.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003124 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SecurityProcessResultLelang:200](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessResultLelang.sql#L200) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tt.SecId` : `INT` | `#tmpTrxLelang.SecId` : `BIGINT` | `tt.SecId` → `BIGINT` | `#tmpTrxLelang.SecId <- tt.SecId` |
| IC-003042 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SecurityPopulateResultLelangEmail:136](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateResultLelangEmail.sql#L136) | WHERE / P11 / COMPARISON | `tl.SecId` : `INT` | `@nSecId` : `BIGINT` | `tl.SecId` → `BIGINT` | `tl.SecId = @nSecId` |
| IC-003044 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SecurityPopulateResultLelangEmail:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateResultLelangEmail.sql#L169) | WHERE / P11 / COMPARISON | `tl.SecId` : `INT` | `@nSecId` : `BIGINT` | `tl.SecId` → `BIGINT` | `tl.SecId = @nSecId` |
| IC-003046 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SecurityPopulateResultLelangEmail:194](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateResultLelangEmail.sql#L194) | WHERE / P11 / COMPARISON | `tl.SecId` : `INT` | `@nSecId` : `BIGINT` | `tl.SecId` → `BIGINT` | `tl.SecId = @nSecId` |
