# Retail UAT.dbo.EarlyRedemption_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 36 lokasi source, 12 procedure, 11 kolom. **A:** 2 lokasi pada 1 procedure. **B:** 5 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-earlyid"></a>
## EarlyId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.EarlyRedemption_TT.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 2 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P06 — LIKE dan representasi karakter:** LIKE pada key numerik memerlukan representasi karakter. Equality numeric yang tervalidasi cocok untuk lookup tepat; prefix/substring membutuhkan requirement pencarian tersendiri, bukan cast kolom sebagai solusi index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012173 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSSearchSecIdEarlyRedem:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchSecIdEarlyRedem.sql#L44) | WHERE / P06 / LIKE | `EarlyId` : `BIGINT` | `rtrim(ltrim(@cCol1)) + '%'` : `VARCHAR (combined result; facets unknown)` | `EarlyId` → `VARCHAR (combined result; facets unknown)` | `EarlyId like rtrim(ltrim(@cCol1)) + '%'` |
| IC-012174 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSSearchSecIdEarlyRedem:60](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchSecIdEarlyRedem.sql#L60) | WHERE / P06 / LIKE | `EarlyId` : `BIGINT` | `rtrim(ltrim(@cCol1)) + '%'` : `VARCHAR (combined result; facets unknown)` | `EarlyId` → `VARCHAR (combined result; facets unknown)` | `EarlyId like rtrim(ltrim(@cCol1)) + '%'` |

<a id="kolom-cifid"></a>
## CIFId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.EarlyRedemption_TT.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002959 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPushMailNotifSettleRedempt:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifSettleRedempt.sql#L89) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[CIFId]` : `BIGINT` | `@nCIFId` : `int` | `[CIFId]` → `int` | `@nCIFId = [CIFId]` |

<a id="kolom-custodianbit"></a>
## CustodianBit

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.EarlyRedemption_TT.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN_EXPRESSION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009093 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRetailPopulateORIOrder:381](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulateORIOrder.sql#L381) | WHERE / P11 / COMPARISON | `isnull(ot.CustodianBit,0)` : `BIT` | `0` : `int` | `isnull(ot.CustodianBit,0)` → `INT` | `isnull(ot.CustodianBit,0) = 0` |
| IC-009094 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRetailPopulateORIOrder:381](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulateORIOrder.sql#L381) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ot.CustodianBit` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(ot.CustodianBit,0)` |

<a id="kolom-instruction"></a>
## Instruction

**Tipe dasar:** `CHAR(2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.EarlyRedemption_TT.sql#L25). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, DESTINATION_COLUMN, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004845 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSORIOnlinePopulateHist:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L192) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `erl.[Instruction]` : `CHAR(2)` | `#tmp_Detail.Instruction` : `varchar(2)` | `erl.[Instruction]` → `varchar(2)` | `#tmp_Detail.Instruction <- erl.[Instruction]` |
| IC-004856 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSORIOnlinePopulateHist:248](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L248) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `erl.[Instruction]` : `CHAR(2)` | `#tmp_Detail.Instruction` : `varchar(2)` | `erl.[Instruction]` → `varchar(2)` | `#tmp_Detail.Instruction <- erl.[Instruction]` |
| IC-000954 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.ORIOnlineReportDailyTransaction:278](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORIOnlineReportDailyTransaction.sql#L278) | WHERE / P12 / IN_VALUE_TYPE | `trx.Instruction` : `CHAR(2)` | `'OM'` : `varchar` | `trx.Instruction` → `VARCHAR` | `trx.Instruction in ('OM','MB')` |
| IC-004849 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSORIOnlinePopulateHist:217](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L217) | WHERE / P12 / COMPARISON | `erl.[Instruction]` : `CHAR(2)` | `'CB'` : `varchar` | `erl.[Instruction]` → `VARCHAR` | `erl.[Instruction] = 'CB'` |
| IC-004860 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSORIOnlinePopulateHist:276](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L276) | WHERE / P12 / IN_VALUE_TYPE | `erl.[Instruction]` : `CHAR(2)` | `'OM'` : `varchar` | `erl.[Instruction]` → `VARCHAR` | `erl.[Instruction] IN ('OM','MB')` |
| IC-012597 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSUpdatePostingResult:861](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L861) | WHERE / P12 / IN_VALUE_TYPE | `ISNULL(ertt.Instruction,'')` : `CHAR(2)` | `'OM'` : `varchar` | `ISNULL(ertt.Instruction,'')` → `VARCHAR` | `ISNULL(ertt.Instruction,'') NOT IN ('OM', 'MB')` |
| IC-012602 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdatePostingResult:904](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L904) | WHERE / P12 / IN_VALUE_TYPE | `ertt.Instruction` : `CHAR(2)` | `'OM'` : `varchar` | `ertt.Instruction` → `VARCHAR` | `ertt.Instruction IN ('OM', 'MB')` |
| IC-001019 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.ORISBNISubmitEarlyRedemption:436](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNISubmitEarlyRedemption.sql#L436) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.earlyredemption_tt.Instruction` : `CHAR(2)` | `'OM'` : `varchar` | `'OM'` → `CHAR(2)` | `retail.dbo.earlyredemption_tt.Instruction <- 'OM'` |
| IC-012598 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUpdatePostingResult:861](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L861) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ertt.Instruction` : `CHAR(2)` | `''` : `varchar` | `''` → `CHAR(2)` | `ISNULL(ertt.Instruction,'')` |

<a id="kolom-isallotted"></a>
## IsAllotted

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.EarlyRedemption_TT.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006550 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessAllotment:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessAllotment.sql#L216) | WHERE / P11 / COMPARISON | `oott.IsAllotted` : `BIT` | `1` : `int` | `oott.IsAllotted` → `INT` | `oott.IsAllotted = 1` |
| IC-006551 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessAllotment:223](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessAllotment.sql#L223) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `oott.IsAllotted` : `BIT` | `1` : `int` | `1` → `BIT` | `oott.IsAllotted = 1` |

<a id="kolom-isptbc"></a>
## IsPTBC

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.EarlyRedemption_TT.sql#L28). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN_EXPRESSION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006250 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateSettlementEarlyRedem:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementEarlyRedem.sql#L63) | CASE / P11 / COMPARISON | `ISNULL(ertt.IsPTBC, 0)` : `BIT` | `1` : `int` | `ISNULL(ertt.IsPTBC, 0)` → `INT` | `ISNULL(ertt.IsPTBC, 0) = 1` |
| IC-007475 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSProcessSettlementEarlyRedem:135](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementEarlyRedem.sql#L135) | CASE / P11 / COMPARISON | `ISNULL(ertt.IsPTBC, 0)` : `BIT` | `1` : `int` | `ISNULL(ertt.IsPTBC, 0)` → `INT` | `ISNULL(ertt.IsPTBC, 0) = 1` |
| IC-010341 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptOrderORIDetail:825](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOrderORIDetail.sql#L825) | WHERE / P11 / COMPARISON | `isnull(early.IsPTBC, 0)` : `BIT` | `1` : `int` | `isnull(early.IsPTBC, 0)` → `INT` | `isnull(early.IsPTBC, 0) = 1` |
| IC-012610 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSUpdatePostingResult:1033](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L1033) | WHERE / P11 / COMPARISON | `ISNULL(er.IsPTBC, 0)` : `BIT` | `1` : `int` | `ISNULL(er.IsPTBC, 0)` → `INT` | `ISNULL(er.IsPTBC, 0) = 1` |
| IC-000953 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.ORIOnlineReportDailyTransaction:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORIOnlineReportDailyTransaction.sql#L265) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `trx.IsPTBC` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(trx.IsPTBC, 0)` |
| IC-006251 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateSettlementEarlyRedem:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementEarlyRedem.sql#L63) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ertt.IsPTBC` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(ertt.IsPTBC, 0)` |
| IC-007476 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSettlementEarlyRedem:135](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementEarlyRedem.sql#L135) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ertt.IsPTBC` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(ertt.IsPTBC, 0)` |
| IC-010342 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptOrderORIDetail:825](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOrderORIDetail.sql#L825) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `early.IsPTBC` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(early.IsPTBC, 0)` |
| IC-012611 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUpdatePostingResult:1033](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L1033) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `er.IsPTBC` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(er.IsPTBC, 0)` |

<a id="kolom-istaxamnesty"></a>
## IsTaxAmnesty

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.EarlyRedemption_TT.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009084 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailPopulateORIOrder:138](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulateORIOrder.sql#L138) | CASE / P11 / COMPARISON | `ot.IsTaxAmnesty` : `BIT` | `0` : `int` | `ot.IsTaxAmnesty` → `INT` | `ot.IsTaxAmnesty = 0` |
| IC-009085 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailPopulateORIOrder:152](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulateORIOrder.sql#L152) | JOIN_ON / P11 / COMPARISON | `ot.IsTaxAmnesty` : `BIT` | `1` : `int` | `ot.IsTaxAmnesty` → `INT` | `ot.IsTaxAmnesty = 1` |

<a id="kolom-nikreff"></a>
## NikReff

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.EarlyRedemption_TT.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005382 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateEarlyRedemAuthorization:144](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEarlyRedemAuthorization.sql#L144) | JOIN_ON / P01 / COMPARISON | `ertt.NikReff` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `ertt.NikReff = ndme.EMPLID` |
| IC-005387 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateEarlyRedemAuthorization:204](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEarlyRedemAuthorization.sql#L204) | JOIN_ON / P01 / COMPARISON | `ertemp.NikReff` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `ertemp.NikReff = ndme.EMPLID` |

<a id="kolom-nominalredemption"></a>
## NominalRedemption

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.EarlyRedemption_TT.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004846 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSORIOnlinePopulateHist:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L192) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `erl.[NominalRedemption] * -1` : `MONEY (combined result; facets unknown)` | `#tmp_Detail.Nominal` : `decimal(21,0)` | `erl.[NominalRedemption] * -1` → `decimal(21,0)` | `#tmp_Detail.Nominal <- erl.[NominalRedemption] * -1` |
| IC-004857 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSORIOnlinePopulateHist:248](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L248) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `erl.[NominalRedemption] * -1` : `MONEY (combined result; facets unknown)` | `#tmp_Detail.Nominal` : `decimal(21,0)` | `erl.[NominalRedemption] * -1` → `decimal(21,0)` | `#tmp_Detail.Nominal <- erl.[NominalRedemption] * -1` |
| IC-004848 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSORIOnlinePopulateHist:198](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L198) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `erl.[NominalRedemption]` : `MONEY` | `-1` : `int` | `-1` → `MONEY` | `erl.[NominalRedemption] * -1` |
| IC-004859 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSORIOnlinePopulateHist:254](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L254) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `erl.[NominalRedemption]` : `MONEY` | `-1` : `int` | `-1` → `MONEY` | `erl.[NominalRedemption] * -1` |
| IC-005385 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateEarlyRedemAuthorization:189](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEarlyRedemAuthorization.sql#L189) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ertemp.NominalRedemption` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(ertemp.NominalRedemption,0)` |

<a id="kolom-outstanding"></a>
## Outstanding

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.EarlyRedemption_TT.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005386 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateEarlyRedemAuthorization:190](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEarlyRedemAuthorization.sql#L190) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ertemp.Outstanding` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(ertemp.Outstanding,0)` |

<a id="kolom-trxdate"></a>
## TrxDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.EarlyRedemption_TT.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001020 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.ORISBNISubmitEarlyRedemption:436](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNISubmitEarlyRedemption.sql#L436) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.earlyredemption_tt.TrxDate` : `DATETIME` | `convert(varchar(15), getdate(), 112)` : `varchar(15)` | `convert(varchar(15), getdate(), 112)` → `DATETIME` | `retail.dbo.earlyredemption_tt.TrxDate <- convert(varchar(15), getdate(), 112)` |
