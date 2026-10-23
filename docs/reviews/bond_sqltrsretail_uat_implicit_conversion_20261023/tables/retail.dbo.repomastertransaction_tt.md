# Retail UAT.dbo.REPOMasterTransaction_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 98 lokasi source, 21 procedure, 23 kolom. **A:** 14 lokasi pada 8 procedure. **B:** 23 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-nomurexexternal"></a>
## NoMurexExternal

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L41). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 6 SP. **A:** 6 lokasi / 5 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000245 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:351](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L351) | JOIN_ON / P01 / COMPARISON | `rp.NoMurexExternal` : `VARCHAR(20)` | `ctp.MurexNo` : `BIGINT` | `rp.NoMurexExternal` → `BIGINT` | `rp.NoMurexExternal = ctp.MurexNo` |
| IC-001881 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.RepoRONDataMigrate:73](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L73) | JOIN_ON / P01 / COMPARISON | `rmtt.NoMurexExternal` : `VARCHAR(20)` | `rmt.NoMurexExternal` : `NUMERIC(10,0)` | `rmtt.NoMurexExternal` → `NUMERIC` | `rmt.NoMurexExternal = rmtt.NoMurexExternal` |
| IC-001904 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.RepoReportNoticePayment:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoReportNoticePayment.sql#L37) | WHERE / P01 / COMPARISON | `NoMurexExternal` : `VARCHAR(20)` | `@pnNoRef` : `bigint` | `NoMurexExternal` → `BIGINT` | `NoMurexExternal = @pnNoRef` |
| IC-008352 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:524](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L524) | JOIN_ON / P01 / COMPARISON | `tt.NoMurexExternal` : `VARCHAR(20)` | `ctp.MurexNo` : `BIGINT` | `tt.NoMurexExternal` → `BIGINT` | `tt.NoMurexExternal = ctp.MurexNo` |
| IC-010160 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiREPO:10](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiREPO.sql#L10) | WHERE / P01 / COMPARISON | `NoMurexExternal` : `VARCHAR(20)` | `@pnNoRef` : `BIGINT` | `NoMurexExternal` → `BIGINT` | `NoMurexExternal = @pnNoRef` |
| IC-010169 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiREPO:60](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiREPO.sql#L60) | WHERE / P01 / COMPARISON | `tt.NoMurexExternal` : `VARCHAR(20)` | `@pnNoRef` : `BIGINT` | `tt.NoMurexExternal` → `BIGINT` | `tt.NoMurexExternal = @pnNoRef` |
| IC-001873 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repomastertransaction_tt.NoMurexExternal` : `VARCHAR(20)` | `rmt.NoMurexExternal` : `NUMERIC(10,0)` | `rmt.NoMurexExternal` → `VARCHAR(20)` | `retail.dbo.repomastertransaction_tt.NoMurexExternal <- rmt.NoMurexExternal` |
| IC-007991 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoInsertEmailProcess:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoInsertEmailProcess.sql#L77) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tt.NoMurexExternal` : `VARCHAR(20)` | `0` : `int` | `0` → `VARCHAR(20)` | `ISNULL(tt.NoMurexExternal, 0)` |

<a id="kolom-nikinputer"></a>
## NIKInputer

**Tipe dasar:** `VARCHAR(10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L34). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 3 SP. **A:** 4 lokasi / 3 SP. **Peran:** CONVERTED_ARGUMENT, CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001752 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PopulateBlotterRepo:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PopulateBlotterRepo.sql#L228) | WHERE / P01 / COMPARISON | `tt.NIKInputer` : `VARCHAR(10)` | `@pnInputter` : `INT` | `tt.NIKInputer` → `INT` | `tt.NIKInputer = @pnInputter` |
| IC-001774 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.REPOPopulateHapusTransaksi:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.REPOPopulateHapusTransaksi.sql#L56) | JOIN_ON / P01 / COMPARISON | `sttt.NIKInputer` : `VARCHAR(10)` | `unInput.nik` : `DECIMAL (combined result; facets unknown)` | `sttt.NIKInputer` → `DECIMAL` | `sttt.NIKInputer = unInput.nik` |
| IC-001777 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.REPOPopulateHapusTransaksi:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.REPOPopulateHapusTransaksi.sql#L94) | JOIN_ON / P01 / COMPARISON | `sttt.NIKInputer` : `VARCHAR(10)` | `unInput.nik` : `DECIMAL (combined result; facets unknown)` | `sttt.NIKInputer` → `DECIMAL` | `sttt.NIKInputer = unInput.nik` |
| IC-001839 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.RepoPopulateBlotter:241](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoPopulateBlotter.sql#L241) | WHERE / P01 / COMPARISON | `tt.NIKInputer` : `VARCHAR(10)` | `@pnInputter` : `INT` | `tt.NIKInputer` → `INT` | `tt.NIKInputer = @pnInputter` |
| IC-001747 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.PopulateBlotterRepo:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PopulateBlotterRepo.sql#L193) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `tt.NIKInputer` : `VARCHAR(10)` | `retail.dbo.fngetnamebynik.@nNIK` : `int` | `tt.NIKInputer` → `int` | `dbo.fnGetNameByNIK(tt.NIKInputer)` |
| IC-001833 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.RepoPopulateBlotter:195](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoPopulateBlotter.sql#L195) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `tt.NIKInputer` : `VARCHAR(10)` | `retail.dbo.fngetnamebynik.@nNIK` : `int` | `tt.NIKInputer` → `int` | `dbo.fnGetNameByNIK(tt.NIKInputer)` |

<a id="kolom-nikseller"></a>
## NIKSeller

**Tipe dasar:** `VARCHAR(10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L32). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 3 SP. **A:** 4 lokasi / 3 SP. **Peran:** CONVERTED_ARGUMENT, CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001753 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PopulateBlotterRepo:229](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PopulateBlotterRepo.sql#L229) | WHERE / P01 / COMPARISON | `tt.NIKSeller` : `VARCHAR(10)` | `@pnSeller` : `INT` | `tt.NIKSeller` → `INT` | `tt.NIKSeller = @pnSeller` |
| IC-001773 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.REPOPopulateHapusTransaksi:55](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.REPOPopulateHapusTransaksi.sql#L55) | JOIN_ON / P01 / COMPARISON | `sttt.NIKSeller` : `VARCHAR(10)` | `unSeller.nik` : `DECIMAL (combined result; facets unknown)` | `sttt.NIKSeller` → `DECIMAL` | `sttt.NIKSeller = unSeller.nik` |
| IC-001776 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.REPOPopulateHapusTransaksi:93](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.REPOPopulateHapusTransaksi.sql#L93) | JOIN_ON / P01 / COMPARISON | `sttt.NIKSeller` : `VARCHAR(10)` | `unSeller.nik` : `DECIMAL (combined result; facets unknown)` | `sttt.NIKSeller` → `DECIMAL` | `sttt.NIKSeller = unSeller.nik` |
| IC-001840 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.RepoPopulateBlotter:242](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoPopulateBlotter.sql#L242) | WHERE / P01 / COMPARISON | `tt.NIKSeller` : `VARCHAR(10)` | `@pnSeller` : `INT` | `tt.NIKSeller` → `INT` | `tt.NIKSeller = @pnSeller` |
| IC-001746 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.PopulateBlotterRepo:190](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PopulateBlotterRepo.sql#L190) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `tt.NIKSeller` : `VARCHAR(10)` | `retail.dbo.fngetnamebynik.@nNIK` : `int` | `tt.NIKSeller` → `int` | `dbo.fnGetNameByNIK(tt.NIKSeller)` |
| IC-001832 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.RepoPopulateBlotter:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoPopulateBlotter.sql#L193) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `tt.NIKSeller` : `VARCHAR(10)` | `retail.dbo.fngetnamebynik.@nNIK` : `int` | `tt.NIKSeller` → `int` | `dbo.fnGetNameByNIK(tt.NIKSeller)` |

<a id="kolom-admnfee"></a>
## AdmnFee

**Tipe dasar:** `DECIMAL(20,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001867 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repomastertransaction_tt.AdmnFee` : `DECIMAL(20,2)` | `128000` : `int` | `128000` → `DECIMAL(20,2)` | `retail.dbo.repomastertransaction_tt.AdmnFee <- 128000` |

<a id="kolom-biayaprovisi"></a>
## BiayaProvisi

**Tipe dasar:** `DECIMAL(25,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT.

**Pola dan arah review:**

- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-010166 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSRptKonfirmasiREPO:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiREPO.sql#L47) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `tt.BiayaProvisi` : `DECIMAL(25,2)` | `retail.dbo.fnformatstringthseparator.@cValue` : `varchar(100)` | `tt.BiayaProvisi` → `varchar(100)` | `dbo.fnFormatStringThSeparator(tt.BiayaProvisi, 0)` |

<a id="kolom-dirtyprice"></a>
## DirtyPrice

**Tipe dasar:** `DECIMAL(15,4)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L31). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-010164 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSRptKonfirmasiREPO:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiREPO.sql#L31) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `tt.DirtyPrice` : `DECIMAL(15,4)` | `retail.dbo.fnformatstringthseparator.@cValue` : `varchar(100)` | `tt.DirtyPrice` → `varchar(100)` | `dbo.fnFormatStringThSeparator(tt.DirtyPrice, 2)` |
| IC-001868 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repomastertransaction_tt.DirtyPrice` : `DECIMAL(15,4)` | `rmt.DirtyPrice` : `NUMERIC(24,14)` | `rmt.DirtyPrice` → `DECIMAL(15,4)` | `retail.dbo.repomastertransaction_tt.DirtyPrice <- rmt.DirtyPrice` |

<a id="kolom-facevalue"></a>
## FaceValue

**Tipe dasar:** `DECIMAL(25,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008346 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSReportReportSettlementInstruction:504](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L504) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `tt.FaceValue` : `DECIMAL(25,2)` | `retail.dbo.fnformatstringthseparator.@cValue` : `varchar(100)` | `tt.FaceValue` → `varchar(100)` | `dbo.fnFormatStringThSeparator(tt.FaceValue, 2)` |
| IC-010163 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSRptKonfirmasiREPO:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiREPO.sql#L30) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `tt.FaceValue` : `DECIMAL(25,2)` | `retail.dbo.fnformatstringthseparator.@cValue` : `varchar(100)` | `tt.FaceValue` → `varchar(100)` | `dbo.fnFormatStringThSeparator(tt.FaceValue, 0)` |
| IC-001869 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repomastertransaction_tt.FaceValue` : `DECIMAL(25,2)` | `rmt.FaceValue` : `NVARCHAR(38)` | `rmt.FaceValue` → `DECIMAL(25,2)` | `retail.dbo.repomastertransaction_tt.FaceValue <- rmt.FaceValue` |

<a id="kolom-haircut"></a>
## Haircut

**Tipe dasar:** `DECIMAL(10,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L28). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001816 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RepoGetXMLToMurex:221](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoGetXMLToMurex.sql#L221) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `Haircut` : `DECIMAL(10,2)` | `-1` : `int` | `-1` → `DECIMAL` | `Haircut*-1` |
| IC-001817 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RepoGetXMLToMurex:245](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoGetXMLToMurex.sql#L245) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `Haircut` : `DECIMAL(10,2)` | `-1` : `int` | `-1` → `DECIMAL` | `Haircut*-1` |
| IC-001818 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RepoGetXMLToMurex:271](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoGetXMLToMurex.sql#L271) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `Haircut` : `DECIMAL(10,2)` | `-1` : `int` | `-1` → `DECIMAL` | `Haircut*-1` |
| IC-001819 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RepoGetXMLToMurex:294](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoGetXMLToMurex.sql#L294) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `Haircut` : `DECIMAL(10,2)` | `-1` : `int` | `-1` → `DECIMAL` | `Haircut*-1` |
| IC-008347 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportReportSettlementInstruction:505](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L505) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `tt.Haircut` : `DECIMAL(10,2)` | `100` : `int` | `100` → `DECIMAL` | `tt.Haircut / 100` |
| IC-008348 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportReportSettlementInstruction:505](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L505) | EXPRESSION_OR_CONTROL / P05 / BINARY_Subtract | `(tt.Haircut / 100)` : `DECIMAL (combined result; facets unknown)` | `1` : `int` | `1` → `DECIMAL` | `1 - (tt.Haircut / 100)` |

<a id="kolom-haircutmethod"></a>
## HaircutMethod

**Tipe dasar:** `VARCHAR(15)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L27). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001870 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repomastertransaction_tt.HaircutMethod` : `VARCHAR(15)` | `rmt.HaircutMethod` : `NVARCHAR(5)` | `rmt.HaircutMethod` → `VARCHAR(15)` | `retail.dbo.repomastertransaction_tt.HaircutMethod <- rmt.HaircutMethod` |

<a id="kolom-ischeckabsorbedbybank"></a>
## IsCheckAbsorbedByBank

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L55). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001859 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoProsesCancelTrx:299](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoProsesCancelTrx.sql#L299) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `rm.IsCheckAbsorbedByBank` : `BIGINT` | `@pbIsAbsorbedByBank` : `BIT` | `@pbIsAbsorbedByBank` → `BIGINT` | `rm.IsCheckAbsorbedByBank = @pbIsAbsorbedByBank` |

<a id="kolom-lastupdatenik"></a>
## LastUpdateNIK

**Tipe dasar:** `VARCHAR(10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L37). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001853 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoProsesCancelTrx:116](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoProsesCancelTrx.sql#L116) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `sttt.LastUpdateNIK` : `VARCHAR(10)` | `@pnUserNik` : `BIGINT` | `@pnUserNik` → `VARCHAR(10)` | `sttt.LastUpdateNIK = @pnUserNik` |
| IC-001854 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoProsesCancelTrx:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoProsesCancelTrx.sql#L125) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `sttt.LastUpdateNIK` : `VARCHAR(10)` | `@pnUserNik` : `BIGINT` | `@pnUserNik` → `VARCHAR(10)` | `sttt.LastUpdateNIK = @pnUserNik` |
| IC-001855 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoProsesCancelTrx:138](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoProsesCancelTrx.sql#L138) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `sttt.LastUpdateNIK` : `VARCHAR(10)` | `@pnUserNik` : `BIGINT` | `@pnUserNik` → `VARCHAR(10)` | `sttt.LastUpdateNIK = @pnUserNik` |
| IC-001856 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoProsesCancelTrx:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoProsesCancelTrx.sql#L147) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `sttt.LastUpdateNIK` : `VARCHAR(10)` | `@pnUserNik` : `BIGINT` | `@pnUserNik` → `VARCHAR(10)` | `sttt.LastUpdateNIK = @pnUserNik` |
| IC-001857 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoProsesCancelTrx:158](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoProsesCancelTrx.sql#L158) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `sttt.LastUpdateNIK` : `VARCHAR(10)` | `@pnUserNik` : `BIGINT` | `@pnUserNik` → `VARCHAR(10)` | `sttt.LastUpdateNIK = @pnUserNik` |
| IC-001858 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoProsesCancelTrx:167](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoProsesCancelTrx.sql#L167) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `sttt.LastUpdateNIK` : `VARCHAR(10)` | `@pnUserNik` : `BIGINT` | `@pnUserNik` → `VARCHAR(10)` | `sttt.LastUpdateNIK = @pnUserNik` |
| IC-001944 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSubmitTransaction:527](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSubmitTransaction.sql#L527) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `LastUpdateNIK` : `VARCHAR(10)` | `@pnNIK` : `INT` | `@pnNIK` → `VARCHAR(10)` | `LastUpdateNIK = @pnNIK` |
| IC-001945 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSubmitTransaction:561](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSubmitTransaction.sql#L561) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `LastUpdateNIK` : `VARCHAR(10)` | `@pnNIK` : `INT` | `@pnNIK` → `VARCHAR(10)` | `LastUpdateNIK = @pnNIK` |

<a id="kolom-maturitydate"></a>
## MaturityDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L14). **Existing leading index:** REPOMasterTransaction_TT_N2.

**Frekuensi source:** 5 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001871 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repomastertransaction_tt.MaturityDate` : `DATETIME` | `rmt.MaturityDate` : `NVARCHAR(8)` | `rmt.MaturityDate` → `DATETIME` | `retail.dbo.repomastertransaction_tt.MaturityDate <- rmt.MaturityDate` |
| IC-004755 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSNotifRepo:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSNotifRepo.sql#L46) | WHERE / P04 / COMPARISON | `tt.MaturityDate` : `DATETIME` | `@cDateH5` : `VARCHAR(8)` | `@cDateH5` → `DATETIME` | `tt.MaturityDate <= @cDateH5` |
| IC-004756 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSNotifRepo:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSNotifRepo.sql#L47) | WHERE / P04 / COMPARISON | `tt.MaturityDate` : `DATETIME` | `@cDateH1` : `VARCHAR(8)` | `@cDateH1` → `DATETIME` | `tt.MaturityDate >= @cDateH1` |
| IC-004951 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPDFDefaultREPO:54](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPDFDefaultREPO.sql#L54) | WHERE / P04 / COMPARISON | `tt.MaturityDate` : `DATETIME` | `@cDate` : `varchar(8)` | `@cDate` → `DATETIME` | `tt.MaturityDate = @cDate` |
| IC-004956 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPDFNoticeREPO:54](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPDFNoticeREPO.sql#L54) | WHERE / P04 / COMPARISON | `tt.MaturityDate` : `DATETIME` | `@cDate` : `varchar(8)` | `@cDate` → `DATETIME` | `tt.MaturityDate = @cDate` |

<a id="kolom-nikreferrer"></a>
## NIKReferrer

**Tipe dasar:** `VARCHAR(10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L33). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT.

**Pola dan arah review:**

- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001834 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.RepoPopulateBlotter:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoPopulateBlotter.sql#L228) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `tt.NIKReferrer` : `VARCHAR(10)` | `retail.dbo.fngetnamebynik.@nNIK` : `int` | `tt.NIKReferrer` → `int` | `dbo.fnGetNameByNIK(tt.NIKReferrer)` |

<a id="kolom-nodays"></a>
## NoDays

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001872 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repomastertransaction_tt.NoDays` : `INT` | `rmt.NoDays` : `NVARCHAR(38)` | `rmt.NoDays` → `INT` | `retail.dbo.repomastertransaction_tt.NoDays <- rmt.NoDays` |

<a id="kolom-phoneorderbit"></a>
## PhoneOrderBit

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L56). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001943 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.RepoSubmitTransaction:422](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSubmitTransaction.sql#L422) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tt.PhoneOrderBit` : `BIGINT` | `retail.dbo.repomastertransaction_th.PhoneOrderBit` : `BIT` | `tt.PhoneOrderBit` → `BIT` | `retail.dbo.repomastertransaction_th.PhoneOrderBit <- tt.PhoneOrderBit` |
| IC-001948 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.RepoSubmitTransaction:633](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSubmitTransaction.sql#L633) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tt.PhoneOrderBit` : `BIGINT` | `retail.dbo.repomastertransaction_th.PhoneOrderBit` : `BIT` | `tt.PhoneOrderBit` → `BIT` | `retail.dbo.repomastertransaction_th.PhoneOrderBit <- tt.PhoneOrderBit` |
| IC-001745 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PopulateBlotterRepo:167](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PopulateBlotterRepo.sql#L167) | CASE / P11 / COMPARISON | `tt.PhoneOrderBit` : `BIGINT` | `1` : `int` | `1` → `BIGINT` | `tt.PhoneOrderBit = 1` |
| IC-001844 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RepoPopulateDataApproval:87](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoPopulateDataApproval.sql#L87) | CASE / P11 / COMPARISON | `tt.PhoneOrderBit` : `BIGINT` | `1` : `int` | `1` → `BIGINT` | `tt.PhoneOrderBit = 1` |
| IC-001846 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RepoPopulateDataCancelApproval:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoPopulateDataCancelApproval.sql#L56) | CASE / P11 / COMPARISON | `tt.PhoneOrderBit` : `BIGINT` | `1` : `int` | `1` → `BIGINT` | `tt.PhoneOrderBit = 1` |
| IC-001938 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSubmitTransaction:231](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSubmitTransaction.sql#L231) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repomastertransaction_tt.PhoneOrderBit` : `BIGINT` | `PhoneOrderBit` : `BIT` | `PhoneOrderBit` → `BIGINT` | `retail.dbo.repomastertransaction_tt.PhoneOrderBit <- PhoneOrderBit` |
| IC-001946 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSubmitTransaction:577](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSubmitTransaction.sql#L577) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `PhoneOrderBit` : `BIGINT` | `tmp.PhoneOrderBit` : `BIT` | `tmp.PhoneOrderBit` → `BIGINT` | `PhoneOrderBit = tmp.PhoneOrderBit` |
| IC-001947 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RepoSubmitTransaction:618](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSubmitTransaction.sql#L618) | WHERE / P11 / COMPARISON | `tt.PhoneOrderBit` : `BIGINT` | `tmp.PhoneOrderBit` : `BIT` | `tmp.PhoneOrderBit` → `BIGINT` | `tt.PhoneOrderBit <> tmp.PhoneOrderBit` |

<a id="kolom-rateconvention"></a>
## RateConvention

**Tipe dasar:** `VARCHAR(25)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L30). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001874 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repomastertransaction_tt.RateConvention` : `VARCHAR(25)` | `rmt.RateConvention` : `NVARCHAR(15)` | `rmt.RateConvention` → `VARCHAR(25)` | `retail.dbo.repomastertransaction_tt.RateConvention <- rmt.RateConvention` |

<a id="kolom-repo1stleg"></a>
## Repo1stLeg

**Tipe dasar:** `DECIMAL(25,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT, COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008196 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportRekonMaterai:202](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L202) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `rmt.Repo1stLeg` : `DECIMAL(25,2)` | `#tempResult.TotalProceed` : `MONEY` | `rmt.Repo1stLeg` → `MONEY` | `#tempResult.TotalProceed <- rmt.Repo1stLeg` |
| IC-008349 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSReportReportSettlementInstruction:506](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L506) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `tt.Repo1stLeg` : `DECIMAL(25,2)` | `retail.dbo.fnformatstringthseparator.@cValue` : `varchar(100)` | `tt.Repo1stLeg` → `varchar(100)` | `dbo.fnFormatStringThSeparator(tt.Repo1stLeg, 2)` |
| IC-010165 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSRptKonfirmasiREPO:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiREPO.sql#L46) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `tt.Repo1stLeg` : `DECIMAL(25,2)` | `retail.dbo.fnformatstringthseparator.@cValue` : `varchar(100)` | `tt.Repo1stLeg` → `varchar(100)` | `dbo.fnFormatStringThSeparator(tt.Repo1stLeg, 0)` |
| IC-001875 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repomastertransaction_tt.Repo1stLeg` : `DECIMAL(25,2)` | `rmt.Repo1stLeg` : `NVARCHAR(38)` | `rmt.Repo1stLeg` → `DECIMAL(25,2)` | `retail.dbo.repomastertransaction_tt.Repo1stLeg <- rmt.Repo1stLeg` |
| IC-008207 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:254](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L254) | WHERE / P11 / COMPARISON | `rmt.Repo1stLeg` : `DECIMAL(25,2)` | `0` : `int` | `0` → `DECIMAL` | `rmt.Repo1stLeg != 0` |

<a id="kolom-repo2ndleg"></a>
## Repo2ndLeg

**Tipe dasar:** `DECIMAL(25,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001901 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.RepoReportNoticePayment:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoReportNoticePayment.sql#L22) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `Repo2ndLeg` : `DECIMAL(25,2)` | `retail.dbo.fngetterbilang.@pcNominal` : `[nvarchar](4000)` | `Repo2ndLeg` → `[nvarchar](4000)` | `dbo.fnGetTerbilang(Repo2ndLeg, sm.SecCcy)` |
| IC-001903 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.RepoReportNoticePayment:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoReportNoticePayment.sql#L23) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `Repo2ndLeg` : `DECIMAL(25,2)` | `retail.dbo.fngetterbilangenglish.@pcNominal` : `[nvarchar](4000)` | `Repo2ndLeg` → `[nvarchar](4000)` | `dbo.fnGetTerbilangEnglish(Repo2ndLeg, sm.SecCcy)` |
| IC-008350 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSReportReportSettlementInstruction:511](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L511) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `tt.Repo2ndLeg` : `DECIMAL(25,2)` | `retail.dbo.fnformatstringthseparator.@cValue` : `varchar(100)` | `tt.Repo2ndLeg` → `varchar(100)` | `dbo.fnFormatStringThSeparator(tt.Repo2ndLeg, 2)` |
| IC-010168 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSRptKonfirmasiREPO:50](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiREPO.sql#L50) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `tt.Repo2ndLeg` : `DECIMAL(25,2)` | `retail.dbo.fnformatstringthseparator.@cValue` : `varchar(100)` | `tt.Repo2ndLeg` → `varchar(100)` | `dbo.fnFormatStringThSeparator(tt.Repo2ndLeg, 0)` |
| IC-001876 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repomastertransaction_tt.Repo2ndLeg` : `DECIMAL(25,2)` | `rmt.Repo2ndLeg` : `NVARCHAR(38)` | `rmt.Repo2ndLeg` → `DECIMAL(25,2)` | `retail.dbo.repomastertransaction_tt.Repo2ndLeg <- rmt.Repo2ndLeg` |

<a id="kolom-repodealid"></a>
## RepoDealId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L2). **Existing leading index:** REPOMasterTransaction_TT_CN1.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001883 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.RepoRONDataMigrate:78](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L78) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `max(RepoDealId)` : `BIGINT` | `@nMax` : `int` | `max(RepoDealId)` → `int` | `@nMax = max(RepoDealId)` |

<a id="kolom-repointerest"></a>
## RepoInterest

**Tipe dasar:** `DECIMAL(15,4)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L25). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-010167 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSRptKonfirmasiREPO:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiREPO.sql#L49) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `tt.RepoInterest` : `DECIMAL(15,4)` | `retail.dbo.fnformatstringthseparator.@cValue` : `varchar(100)` | `tt.RepoInterest` → `varchar(100)` | `dbo.fnFormatStringThSeparator(tt.RepoInterest, 0)` |
| IC-001877 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repomastertransaction_tt.RepoInterest` : `DECIMAL(15,4)` | `rmt.RepoInterest` : `NVARCHAR(38)` | `rmt.RepoInterest` → `DECIMAL(15,4)` | `retail.dbo.repomastertransaction_tt.RepoInterest <- rmt.RepoInterest` |

<a id="kolom-secid"></a>
## SecId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001811 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.RepoGetXMLToMurex:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoGetXMLToMurex.sql#L86) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `SecId` : `INT` | `@nSecId` : `BIGINT` | `SecId` → `BIGINT` | `@nSecId = SecId` |

<a id="kolom-startdate"></a>
## StartDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L12). **Existing leading index:** REPOMasterTransaction_TT_N3.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001750 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PopulateBlotterRepo:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PopulateBlotterRepo.sql#L225) | WHERE / P04 / COMPARISON | `tt.StartDate` : `DATETIME` | `@pdTglBeliEnd` : `VARCHAR(10)` | `@pdTglBeliEnd` → `DATETIME` | `tt.StartDate <= @pdTglBeliEnd` |
| IC-001751 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PopulateBlotterRepo:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PopulateBlotterRepo.sql#L225) | WHERE / P04 / COMPARISON | `tt.StartDate` : `DATETIME` | `@pdTglBeliStart` : `VARCHAR(10)` | `@pdTglBeliStart` → `DATETIME` | `tt.StartDate >= @pdTglBeliStart` |
| IC-001837 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RepoPopulateBlotter:238](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoPopulateBlotter.sql#L238) | WHERE / P04 / COMPARISON | `tt.StartDate` : `DATETIME` | `@pdTglBeliEnd` : `VARCHAR(10)` | `@pdTglBeliEnd` → `DATETIME` | `tt.StartDate <= @pdTglBeliEnd` |
| IC-001838 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RepoPopulateBlotter:238](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoPopulateBlotter.sql#L238) | WHERE / P04 / COMPARISON | `tt.StartDate` : `DATETIME` | `@pdTglBeliStart` : `VARCHAR(10)` | `@pdTglBeliStart` → `DATETIME` | `tt.StartDate >= @pdTglBeliStart` |
| IC-001878 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repomastertransaction_tt.StartDate` : `DATETIME` | `rmt.StartDate` : `NVARCHAR(8)` | `rmt.StartDate` → `DATETIME` | `retail.dbo.repomastertransaction_tt.StartDate <- rmt.StartDate` |

<a id="kolom-trxdate"></a>
## TrxDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOMasterTransaction_TT.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 20 lokasi / 9 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001835 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.RepoPopulateBlotter:237](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoPopulateBlotter.sql#L237) | WHERE / P13 / COMPARISON | `tt.TrxDate` : `DATETIME` | `@pdTglTrxEnd` : `DATETIME2` | `tt.TrxDate` → `DATETIME2` | `tt.TrxDate <= @pdTglTrxEnd` |
| IC-001836 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.RepoPopulateBlotter:237](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoPopulateBlotter.sql#L237) | WHERE / P13 / COMPARISON | `tt.TrxDate` : `DATETIME` | `@pdTglTrxStart` : `DATETIME2` | `tt.TrxDate` → `DATETIME2` | `tt.TrxDate >= @pdTglTrxStart` |
| IC-001748 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PopulateBlotterRepo:224](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PopulateBlotterRepo.sql#L224) | WHERE / P04 / COMPARISON | `tt.TrxDate` : `DATETIME` | `@pdTglTrxEnd` : `VARCHAR(10)` | `@pdTglTrxEnd` → `DATETIME` | `tt.TrxDate <= @pdTglTrxEnd` |
| IC-001749 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PopulateBlotterRepo:224](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PopulateBlotterRepo.sql#L224) | WHERE / P04 / COMPARISON | `tt.TrxDate` : `DATETIME` | `@pdTglTrxStart` : `VARCHAR(10)` | `@pdTglTrxStart` → `DATETIME` | `tt.TrxDate >= @pdTglTrxStart` |
| IC-001879 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repomastertransaction_tt.TrxDate` : `DATETIME` | `rmt.TrxDate` : `NVARCHAR(8)` | `rmt.TrxDate` → `DATETIME` | `retail.dbo.repomastertransaction_tt.TrxDate <- rmt.TrxDate` |
| IC-001898 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RepoReportNoticePayment:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoReportNoticePayment.sql#L17) | CASE / P04 / COMPARISON | `repo.TrxDate` : `DATETIME` | `'2023-05-02'` : `varchar` | `'2023-05-02'` → `DATETIME` | `repo.TrxDate < '2023-05-02'` |
| IC-001899 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RepoReportNoticePayment:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoReportNoticePayment.sql#L18) | CASE / P04 / BETWEEN_BOUND_TYPE | `repo.TrxDate` : `DATETIME` | `'2023-05-02'` : `varchar` | `'2023-05-02'` → `DATETIME` | `repo.TrxDate between '2023-05-02' AND '2023-11-13'` |
| IC-001905 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RepoReportTermSheet:7](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoReportTermSheet.sql#L7) | CASE / P04 / COMPARISON | `repo.TrxDate` : `DATETIME` | `'2023-05-02'` : `varchar` | `'2023-05-02'` → `DATETIME` | `repo.TrxDate < '2023-05-02'` |
| IC-001906 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RepoReportTermSheet:8](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoReportTermSheet.sql#L8) | CASE / P04 / BETWEEN_BOUND_TYPE | `repo.TrxDate` : `DATETIME` | `'2023-05-02'` : `varchar` | `'2023-05-02'` → `DATETIME` | `repo.TrxDate between '2023-05-02' AND '2023-11-13'` |
| IC-004954 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPDFKonfirmasiREPO:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPDFKonfirmasiREPO.sql#L52) | WHERE / P04 / COMPARISON | `tt.TrxDate` : `DATETIME` | `@cDate` : `varchar(8)` | `@cDate` → `DATETIME` | `tt.TrxDate >= @cDate` |
| IC-007832 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSREPOReportBiayaMaterai:244](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L244) | CASE / P04 / COMPARISON | `st.TrxDate` : `DATETIME` | `'20230502'` : `varchar` | `'20230502'` → `DATETIME` | `st.TrxDate < '20230502'` |
| IC-007833 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSREPOReportBiayaMaterai:245](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L245) | CASE / P04 / COMPARISON | `st.TrxDate` : `DATETIME` | `'20230502'` : `varchar` | `'20230502'` → `DATETIME` | `st.TrxDate >= '20230502'` |
| IC-007834 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSREPOReportBiayaMaterai:245](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L245) | CASE / P04 / COMPARISON | `st.TrxDate` : `DATETIME` | `'20231114'` : `varchar` | `'20231114'` → `DATETIME` | `st.TrxDate < '20231114'` |
| IC-007843 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSREPOReportBiayaMaterai:295](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L295) | WHERE / P04 / COMPARISON | `TrxDate` : `DATETIME` | `@cTanggalAwal` : `VARCHAR(15)` | `@cTanggalAwal` → `DATETIME` | `TrxDate >= @cTanggalAwal` |
| IC-007844 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSREPOReportBiayaMaterai:296](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L296) | WHERE / P04 / COMPARISON | `TrxDate` : `DATETIME` | `@cTanggalAkhir` : `VARCHAR(15)` | `@cTanggalAkhir` → `DATETIME` | `TrxDate < @cTanggalAkhir` |
| IC-007845 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSREPOReportBiayaMaterai:297](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L297) | WHERE / P04 / COMPARISON | `TrxDate` : `DATETIME` | `'20230711'` : `varchar` | `'20230711'` → `DATETIME` | `TrxDate > '20230711'` |
| IC-008344 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportReportSettlementInstruction:492](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L492) | CASE / P04 / COMPARISON | `tt.TrxDate` : `DATETIME` | `'2023-05-02'` : `varchar` | `'2023-05-02'` → `DATETIME` | `tt.TrxDate < '2023-05-02'` |
| IC-008345 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportReportSettlementInstruction:493](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L493) | CASE / P04 / BETWEEN_BOUND_TYPE | `tt.TrxDate` : `DATETIME` | `'2023-05-02'` : `varchar` | `'2023-05-02'` → `DATETIME` | `tt.TrxDate between '2023-05-02' AND '2023-11-13'` |
| IC-010161 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptKonfirmasiREPO:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiREPO.sql#L22) | CASE / P04 / COMPARISON | `tt.TrxDate` : `DATETIME` | `'2023-05-02'` : `varchar` | `'2023-05-02'` → `DATETIME` | `tt.TrxDate < '2023-05-02'` |
| IC-010162 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptKonfirmasiREPO:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiREPO.sql#L23) | CASE / P04 / BETWEEN_BOUND_TYPE | `tt.TrxDate` : `DATETIME` | `'2023-05-02'` : `varchar` | `'2023-05-02'` → `DATETIME` | `tt.TrxDate between '2023-05-02' AND '2023-11-13'` |
