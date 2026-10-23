# Retail UAT.dbo.BONDSDetailTransaksiBuyBackPerdana_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 74 lokasi source, 12 procedure, 19 kolom. **A:** 7 lokasi pada 4 procedure. **B:** 16 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-dealidsekunder"></a>
## DealIdSekunder

**Tipe dasar:** `VARCHAR(100)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.BONDSDetailTransaksiBuyBackPerdana_TR.sql#L45). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 5 SP. **A:** 6 lokasi / 3 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008170 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSReportRekonMaterai:123](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L123) | JOIN_ON / P01 / COMPARISON | `bb.DealIdSekunder` : `VARCHAR(100)` | `st.DealId` : `BIGINT` | `bb.DealIdSekunder` → `BIGINT` | `st.DealId = bb.DealIdSekunder` |
| IC-010048 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiBOND:261](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L261) | JOIN_ON / P01 / COMPARISON | `bb.DealIdSekunder` : `VARCHAR(100)` | `stt.DealId` : `BIGINT` | `bb.DealIdSekunder` → `BIGINT` | `bb.DealIdSekunder = stt.DealId` |
| IC-010049 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiBOND:264](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L264) | WHERE / P01 / COMPARISON | `DealIdSekunder` : `VARCHAR(100)` | `@pnDealIdBB` : `bigint` | `DealIdSekunder` → `BIGINT` | `DealIdSekunder = @pnDealIdBB` |
| IC-010052 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiBOND:266](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L266) | WHERE / P01 / COMPARISON | `DealIdSekunder` : `VARCHAR(100)` | `@pnDealIdBB` : `bigint` | `DealIdSekunder` → `BIGINT` | `DealIdSekunder = @pnDealIdBB` |
| IC-010253 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptLaporanPajakBankBuy:205](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L205) | JOIN_ON / P01 / COMPARISON | `bb.DealIdSekunder` : `VARCHAR(100)` | `stttBuyBB.DealId` : `BIGINT` | `bb.DealIdSekunder` → `BIGINT` | `stttBuyBB.DealId = bb.DealIdSekunder` |
| IC-010292 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptLaporanPajakBankBuy:400](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L400) | JOIN_ON / P01 / COMPARISON | `bb.DealIdSekunder` : `VARCHAR(100)` | `stttBuyBB.DealId` : `BIGINT` | `bb.DealIdSekunder` → `BIGINT` | `stttBuyBB.DealId = bb.DealIdSekunder` |
| IC-010022 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptKonfirmasiBOND:136](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L136) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `bb.DealIdSekunder` : `VARCHAR(100)` | `#tmp_data.DealIdSekunder` : `bigint` | `bb.DealIdSekunder` → `bigint` | `#tmp_data.DealIdSekunder <- bb.DealIdSekunder` |
| IC-007072 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessOverBookingMateraiBuyBack:413](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql#L413) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `bb.DealIdSekunder` : `VARCHAR(100)` | `tt.DealId` : `BIGINT` | `tt.DealId` → `VARCHAR(100)` | `bb.DealIdSekunder = tt.DealId` |
| IC-009570 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptCPT:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCPT.sql#L66) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bb.DealIdSekunder` : `VARCHAR(100)` | `0` : `int` | `0` → `VARCHAR(100)` | `isnull(bb.DealIdSekunder,0)` |

<a id="kolom-norekinvestor"></a>
## NoRekInvestor

**Tipe dasar:** `VARCHAR(100)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.BONDSDetailTransaksiBuyBackPerdana_TR.sql#L43). **Existing leading index:** ix_BONDSDetailTransaksiBuyBackPerdana_TR_1&#124;ix_BONDSDetailTransaksiBuyBackPerdana_TR_2.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007070 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSProcessOverBookingMateraiBuyBack:393](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql#L393) | JOIN_ON / P01 / COMPARISON | `bb.NoRekInvestor` : `VARCHAR(100)` | `dd.ACCTNO` : `DECIMAL(19,0)` | `bb.NoRekInvestor` → `DECIMAL` | `dd.ACCTNO = bb.NoRekInvestor` |

<a id="kolom-accrueddays"></a>
## AccruedDays

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.BONDSDetailTransaksiBuyBackPerdana_TR.sql#L19). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006514 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSPopulateTrxBuyBackPerdana:210](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L210) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `AccruedDays` : `MONEY` | `retail.dbo.fncalculateaccruedinterest.@pnAccruedDays` : `int` | `AccruedDays` → `int` | `dbo.fnCalculateAccruedInterest (SecId, CouponPeriod, CouponRate, AccruedDays, LastCouponDate, NextCouponDate, SecCCY, NominalBuyBack, IsCorporateBond, SettlementDate)` |
| IC-007048 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessOverBookingMateraiBuyBack:139](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql#L139) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `bb.AccruedDays` : `MONEY` | `#tSecurityTransaction_TT.AccruedDays` : `int` | `bb.AccruedDays` → `int` | `#tSecurityTransaction_TT.AccruedDays <- bb.AccruedDays` |

<a id="kolom-accruedinterest"></a>
## AccruedInterest

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.BONDSDetailTransaksiBuyBackPerdana_TR.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT, CONVERTED_COLUMN_EXPRESSION, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-010036 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSRptKonfirmasiBOND:229](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L229) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `isnull(bb.AccruedInterest,0)` : `MONEY` | `retail.dbo.fnoblseparatorsign.@pcNominal` : `varchar(100)` | `isnull(bb.AccruedInterest,0)` → `varchar(100)` | `dbo.fnOBLSeparatorSign(isnull(bb.AccruedInterest,0))` |
| IC-012078 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSSaveUploadCTP:136](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveUploadCTP.sql#L136) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `bb.AccruedInterest` : `MONEY` | `tmp.AccruedNet` : `decimal(19,5)` | `bb.AccruedInterest` → `decimal(19,5)` | `tmp.AccruedNet = bb.AccruedInterest` |
| IC-006527 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateTrxBuyBackPerdana:223](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L223) | DML_CONDITION / P05 / BINARY_Multiply | `isnull(AccruedInterest, 0)` : `MONEY` | `( isnull(Tax, 0) / 100.00 )` : `DECIMAL (combined result; facets unknown)` | `isnull(AccruedInterest, 0)` → `DECIMAL` | `isnull(AccruedInterest, 0) * ( isnull(Tax, 0) / 100.00 )` |
| IC-006528 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateTrxBuyBackPerdana:223](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L223) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(AccruedInterest, 0)` |
| IC-010035 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptKonfirmasiBOND:229](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L229) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bb.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(bb.AccruedInterest,0)` |

<a id="kolom-buybackid"></a>
## BuyBackID

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.BONDSDetailTransaksiBuyBackPerdana_TR.sql#L2). **Existing leading index:** PK__BONDSDet__4B63C717CD37182A.

**Frekuensi source:** 7 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007049 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessOverBookingMateraiBuyBack:139](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql#L139) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `bb.BuyBackID` : `INT` | `#tSecurityTransaction_TT.OrderId` : `BIGINT` | `bb.BuyBackID` → `BIGINT` | `#tSecurityTransaction_TT.OrderId <- bb.BuyBackID` |
| IC-000081 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.BuybackLogInsert:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BuybackLogInsert.sql#L43) | JOIN_ON / P11 / COMPARISON | `bb.BuyBackID` : `INT` | `ptt.DealId` : `BIGINT` | `bb.BuyBackID` → `BIGINT` | `ptt.DealId = bb.BuyBackID` |
| IC-007068 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessOverBookingMateraiBuyBack:390](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql#L390) | JOIN_ON / P11 / COMPARISON | `bb.BuyBackID` : `INT` | `tt.OrderId` : `BIGINT` | `bb.BuyBackID` → `BIGINT` | `tt.OrderId = bb.BuyBackID` |
| IC-007073 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessOverBookingMateraiBuyBack:417](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql#L417) | JOIN_ON / P11 / COMPARISON | `bb.BuyBackID` : `INT` | `tt.OrderId` : `BIGINT` | `bb.BuyBackID` → `BIGINT` | `bb.BuyBackID = tt.OrderId` |
| IC-010047 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptKonfirmasiBOND:261](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L261) | JOIN_ON / P11 / COMPARISON | `bb.BuyBackID` : `INT` | `stt.OrderId` : `BIGINT` | `bb.BuyBackID` → `BIGINT` | `bb.BuyBackID = stt.OrderId` |
| IC-011458 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptProcessedBuyBack:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptProcessedBuyBack.sql#L43) | JOIN_ON / P11 / COMPARISON | `bb.BuyBackID` : `INT` | `ptt.DealId` : `BIGINT` | `bb.BuyBackID` → `BIGINT` | `ptt.DealId = bb.BuyBackID` |
| IC-011509 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSettledBuyBack:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettledBuyBack.sql#L69) | JOIN_ON / P11 / COMPARISON | `bb.BuyBackID` : `INT` | `ptt.DealId` : `BIGINT` | `bb.BuyBackID` → `BIGINT` | `ptt.DealId = bb.BuyBackID` |

<a id="kolom-couponrate"></a>
## CouponRate

**Tipe dasar:** `DECIMAL(10,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.BONDSDetailTransaksiBuyBackPerdana_TR.sql#L39). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-010028 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSRptKonfirmasiBOND:215](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L215) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `isnull(bb.CouponRate,0)` : `DECIMAL(10,5)` | `retail.dbo.fnoblseparatorsign.@pcNominal` : `varchar(100)` | `isnull(bb.CouponRate,0)` → `varchar(100)` | `dbo.fnOBLSeparatorSign(isnull(bb.CouponRate,0))` |
| IC-010027 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptKonfirmasiBOND:215](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L215) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bb.CouponRate` : `DECIMAL(10,5)` | `0` : `int` | `0` → `DECIMAL(10,5)` | `isnull(bb.CouponRate,0)` |

<a id="kolom-hargabuyback"></a>
## HargaBuyBack

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.BONDSDetailTransaksiBuyBackPerdana_TR.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 10 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT, CONVERTED_COLUMN, COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-010032 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSRptKonfirmasiBOND:220](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L220) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `isnull(bb.HargaBuyBack,0)` : `MONEY` | `retail.dbo.fnoblseparatorsign.@pcNominal` : `varchar(100)` | `isnull(bb.HargaBuyBack,0)` → `varchar(100)` | `dbo.fnOBLSeparatorSign(isnull(bb.HargaBuyBack,0))` |
| IC-006515 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateTrxBuyBackPerdana:217](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L217) | DML_CONDITION / P05 / BINARY_Divide | `HargaBuyBack` : `MONEY` | `100.00` : `decimal` | `HargaBuyBack` → `DECIMAL` | `HargaBuyBack/100.00` |
| IC-006520 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateTrxBuyBackPerdana:219](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L219) | DML_CONDITION / P05 / BINARY_Divide | `HargaBuyBack` : `MONEY` | `100.00` : `decimal` | `HargaBuyBack` → `DECIMAL` | `HargaBuyBack/100.00` |
| IC-005965 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateParameterBIS4:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParameterBIS4.sql#L43) | WHERE / P04 / COMPARISON | `bb.HargaBuyBack` : `MONEY` | `@pcData2` : `VARCHAR(100)` | `@pcData2` → `MONEY` | `bb.HargaBuyBack = @pcData2` |
| IC-006516 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateTrxBuyBackPerdana:217](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L217) | DML_CONDITION / P05 / BINARY_Multiply | `(HargaBuyBack/100.00)` : `DECIMAL (combined result; facets unknown)` | `NominalBuyBack` : `MONEY` | `NominalBuyBack` → `DECIMAL` | `(HargaBuyBack/100.00) * NominalBuyBack` |
| IC-006518 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateTrxBuyBackPerdana:218](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L218) | DML_CONDITION / P05 / BINARY_Divide | `HargaBuyBack` : `MONEY` | `100` : `int` | `100` → `MONEY` | `HargaBuyBack/100` |
| IC-006521 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateTrxBuyBackPerdana:219](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L219) | DML_CONDITION / P05 / BINARY_Multiply | `(HargaBuyBack/100.00)` : `DECIMAL (combined result; facets unknown)` | `isnull(NominalBuyBack, 0)` : `MONEY` | `isnull(NominalBuyBack, 0)` → `DECIMAL` | `isnull(NominalBuyBack, 0) * (HargaBuyBack/100.00)` |
| IC-010031 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptKonfirmasiBOND:220](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L220) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bb.HargaBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(bb.HargaBuyBack,0)` |
| IC-011506 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptSettledBuyBack:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettledBuyBack.sql#L27) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `bb.HargaBuyBack` : `MONEY` | `100` : `int` | `100` → `MONEY` | `bb.HargaBuyBack/100` |
| IC-011507 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptSettledBuyBack:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettledBuyBack.sql#L27) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `bb.HargaBuyBack` : `MONEY` | `100` : `int` | `100` → `MONEY` | `bb.HargaBuyBack/100` |

<a id="kolom-iscorporatebond"></a>
## IsCorporateBond

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.BONDSDetailTransaksiBuyBackPerdana_TR.sql#L35). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000225 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CustodyReportSettlementInstruction:262](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L262) | CASE / P11 / COMPARISON | `bb.IsCorporateBond` : `BIT` | `1` : `int` | `bb.IsCorporateBond` → `INT` | `bb.IsCorporateBond = 1` |

<a id="kolom-materai"></a>
## Materai

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.BONDSDetailTransaksiBuyBackPerdana_TR.sql#L29). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT, COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007065 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessOverBookingMateraiBuyBack:330](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql#L330) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `bb.Materai` : `MONEY` | `retail.dbo.trsmateraitransaction_tt.MateraiCost` : `DECIMAL(25,2)` | `bb.Materai` → `DECIMAL(25,2)` | `retail.dbo.trsmateraitransaction_tt.MateraiCost <- bb.Materai` |
| IC-010046 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSRptKonfirmasiBOND:255](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L255) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `bb.Materai` : `MONEY` | `retail.dbo.fnoblseparatorsign.@pcNominal` : `varchar(100)` | `bb.Materai` → `varchar(100)` | `dbo.fnOBLSeparatorSign(bb.Materai)` |
| IC-006530 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateTrxBuyBackPerdana:239](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L239) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Materai` : `MONEY` | `case when TotalProceed >= (select top 1 StartLimitNominalTransaksi from TRSParamMaterai_TR where ProductCode = 'BOND') then (select top 1 NominalMaterai from TRSParamMaterai_TR where ProductCode = 'BOND') else 0 end` : `DECIMAL (combined result; facets unknown)` | `case when TotalProceed >= (select top 1 StartLimitNominalTransaksi from TRSParamMaterai_TR where ProductCode = 'BOND') then (select top 1 NominalMaterai from TRSParamMaterai_TR where ProductCode = 'BOND') else 0 end` → `MONEY` | `Materai = case when TotalProceed >= (select top 1 StartLimitNominalTransaksi from TRSParamMaterai_TR where ProductCode = 'BOND') then (select top 1 NominalMaterai from TRSParamMaterai_TR where ProductCode = 'BOND') else 0 end` |
| IC-007071 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessOverBookingMateraiBuyBack:394](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql#L394) | WHERE / P11 / COMPARISON | `bb.Materai` : `MONEY` | `0` : `int` | `0` → `MONEY` | `bb.Materai > 0` |
| IC-008182 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:143](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L143) | WHERE / P11 / COMPARISON | `bb.Materai` : `MONEY` | `0` : `int` | `0` → `MONEY` | `bb.Materai >0` |
| IC-008185 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:154](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L154) | WHERE / P11 / COMPARISON | `bb.Materai` : `MONEY` | `0` : `int` | `0` → `MONEY` | `bb.Materai >0` |
| IC-008187 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L156) | WHERE / P11 / COMPARISON | `bb.Materai` : `MONEY` | `0` : `int` | `0` → `MONEY` | `bb.Materai >0` |
| IC-010044 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptKonfirmasiBOND:252](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L252) | CASE / P11 / COMPARISON | `ISNULL(bb.Materai, 0)` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(bb.Materai, 0) > 0` |
| IC-010045 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptKonfirmasiBOND:252](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L252) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bb.Materai` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(bb.Materai, 0)` |

<a id="kolom-nominalbuyback"></a>
## NominalBuyBack

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.BONDSDetailTransaksiBuyBackPerdana_TR.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT, CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006513 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSPopulateTrxBuyBackPerdana:210](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L210) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `NominalBuyBack` : `MONEY` | `retail.dbo.fncalculateaccruedinterest.@pmFaceValue` : `decimal (20,5)` | `NominalBuyBack` → `decimal (20,5)` | `dbo.fnCalculateAccruedInterest (SecId, CouponPeriod, CouponRate, AccruedDays, LastCouponDate, NextCouponDate, SecCCY, NominalBuyBack, IsCorporateBond, SettlementDate)` |
| IC-009569 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptCPT:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCPT.sql#L51) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `bb.NominalBuyBack` : `MONEY` | `#Temp_list.NominalBuyBack` : `decimal(18,5)` | `bb.NominalBuyBack` → `decimal(18,5)` | `#Temp_list.NominalBuyBack <- bb.NominalBuyBack` |
| IC-006516 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateTrxBuyBackPerdana:217](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L217) | DML_CONDITION / P05 / BINARY_Multiply | `NominalBuyBack` : `MONEY` | `(HargaBuyBack/100.00)` : `DECIMAL (combined result; facets unknown)` | `NominalBuyBack` → `DECIMAL` | `(HargaBuyBack/100.00) * NominalBuyBack` |
| IC-006517 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateTrxBuyBackPerdana:217](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L217) | DML_CONDITION / P05 / BINARY_Subtract | `NominalBuyBack` : `MONEY` | `((HargaBuyBack/100.00) * NominalBuyBack)` : `DECIMAL (combined result; facets unknown)` | `NominalBuyBack` → `DECIMAL` | `((HargaBuyBack/100.00) * NominalBuyBack) - NominalBuyBack` |
| IC-006521 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateTrxBuyBackPerdana:219](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L219) | DML_CONDITION / P05 / BINARY_Multiply | `isnull(NominalBuyBack, 0)` : `MONEY` | `(HargaBuyBack/100.00)` : `DECIMAL (combined result; facets unknown)` | `isnull(NominalBuyBack, 0)` → `DECIMAL` | `isnull(NominalBuyBack, 0) * (HargaBuyBack/100.00)` |
| IC-006522 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateTrxBuyBackPerdana:219](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L219) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NominalBuyBack` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(NominalBuyBack, 0)` |
| IC-012077 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSaveUploadCTP:117](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveUploadCTP.sql#L117) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `sum(bb.NominalBuyBack)` : `MONEY (aggregate; facets unknown)` | `1000000000` : `int` | `1000000000` → `MONEY` | `sum(bb.NominalBuyBack)/1000000000` |

<a id="kolom-priority"></a>
## Priority

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.BONDSDetailTransaksiBuyBackPerdana_TR.sql#L25). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006505 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateTrxBuyBackPerdana:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.bondsdetailtransaksibuybackperdana_tr.Priority` : `INT` | `oott.Priority` : `BIT` | `oott.Priority` → `INT` | `retail.dbo.bondsdetailtransaksibuybackperdana_tr.Priority <- oott.Priority` |

<a id="kolom-proceed"></a>
## Proceed

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.BONDSDetailTransaksiBuyBackPerdana_TR.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-010034 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSRptKonfirmasiBOND:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L225) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `isnull(bb.Proceed,0)` : `MONEY` | `retail.dbo.fnoblseparatorsign.@pcNominal` : `varchar(100)` | `isnull(bb.Proceed,0)` → `varchar(100)` | `dbo.fnOBLSeparatorSign(isnull(bb.Proceed,0))` |
| IC-010033 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptKonfirmasiBOND:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L225) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bb.Proceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(bb.Proceed,0)` |

<a id="kolom-safekeepingfee"></a>
## SafeKeepingFee

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.BONDSDetailTransaksiBuyBackPerdana_TR.sql#L27). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000235 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.CustodyReportSettlementInstruction:283](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L283) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bb.SafeKeepingFee` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(bb.SafeKeepingFee,0)` |
| IC-006524 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateTrxBuyBackPerdana:221](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L221) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `SafeKeepingFee` : `MONEY` | `0` : `int` | `0` → `MONEY` | `SafeKeepingFee = 0` |

<a id="kolom-tax"></a>
## Tax

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.BONDSDetailTransaksiBuyBackPerdana_TR.sql#L30). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN_EXPRESSION, COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006526 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateTrxBuyBackPerdana:223](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L223) | DML_CONDITION / P05 / BINARY_Divide | `isnull(Tax, 0)` : `MONEY` | `100.00` : `decimal` | `isnull(Tax, 0)` → `DECIMAL` | `isnull(Tax, 0) / 100.00` |
| IC-006519 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateTrxBuyBackPerdana:218](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L218) | DML_CONDITION / P05 / BINARY_Divide | `Tax` : `MONEY` | `100` : `int` | `100` → `MONEY` | `Tax / 100` |
| IC-006527 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateTrxBuyBackPerdana:223](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L223) | DML_CONDITION / P05 / BINARY_Multiply | `( isnull(Tax, 0) / 100.00 )` : `DECIMAL (combined result; facets unknown)` | `isnull(AccruedInterest, 0)` : `MONEY` | `isnull(AccruedInterest, 0)` → `DECIMAL` | `isnull(AccruedInterest, 0) * ( isnull(Tax, 0) / 100.00 )` |
| IC-006529 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateTrxBuyBackPerdana:223](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L223) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `Tax` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(Tax, 0)` |
| IC-011508 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptSettledBuyBack:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettledBuyBack.sql#L27) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `bb.Tax` : `MONEY` | `100` : `int` | `100` → `MONEY` | `bb.Tax / 100` |

<a id="kolom-taxonaccrued"></a>
## TaxOnAccrued

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.BONDSDetailTransaksiBuyBackPerdana_TR.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000234 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.CustodyReportSettlementInstruction:282](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L282) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bb.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(bb.TaxOnAccrued,0)` |
| IC-010038 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptKonfirmasiBOND:235](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L235) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bb.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(bb.TaxOnAccrued,0)` |
| IC-010040 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptKonfirmasiBOND:236](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L236) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bb.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(bb.TaxOnAccrued,0)` |

<a id="kolom-taxoncapitalgain"></a>
## TaxOnCapitalGain

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.BONDSDetailTransaksiBuyBackPerdana_TR.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000233 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.CustodyReportSettlementInstruction:281](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L281) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bb.TaxOnCapitalGain` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(bb.TaxOnCapitalGain,0)` |
| IC-010039 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptKonfirmasiBOND:235](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L235) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bb.TaxOnCapitalGain` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(bb.TaxOnCapitalGain,0)` |
| IC-010041 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptKonfirmasiBOND:236](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L236) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bb.TaxOnCapitalGain` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(bb.TaxOnCapitalGain,0)` |

<a id="kolom-totalfee"></a>
## TotalFee

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.BONDSDetailTransaksiBuyBackPerdana_TR.sql#L40). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-010025 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptKonfirmasiBOND:136](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L136) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `bb.TotalFee` : `MONEY` | `#tmp_data.TotalFee` : `varchar(40)` | `bb.TotalFee` → `varchar(40)` | `#tmp_data.TotalFee <- bb.TotalFee` |
| IC-006523 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateTrxBuyBackPerdana:220](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L220) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `TotalFee` : `MONEY` | `0` : `int` | `0` → `MONEY` | `TotalFee = 0` |

<a id="kolom-totalproceed"></a>
## TotalProceed

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.BONDSDetailTransaksiBuyBackPerdana_TR.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT, CONVERTED_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-010030 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSRptKonfirmasiBOND:219](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L219) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `isnull(bb.TotalProceed,0)` : `MONEY` | `retail.dbo.fnoblseparatorsign.@pcNominal` : `varchar(100)` | `isnull(bb.TotalProceed,0)` → `varchar(100)` | `dbo.fnOBLSeparatorSign(isnull(bb.TotalProceed,0))` |
| IC-010043 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSRptKonfirmasiBOND:241](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L241) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `isnull(bb.TotalProceed,0)` : `MONEY` | `retail.dbo.fnoblseparatorsign.@pcNominal` : `varchar(100)` | `isnull(bb.TotalProceed,0)` → `varchar(100)` | `dbo.fnOBLSeparatorSign(isnull(bb.TotalProceed,0))` |
| IC-006531 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateTrxBuyBackPerdana:240](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L240) | CASE / P11 / COMPARISON | `TotalProceed` : `MONEY` | `(select top 1 StartLimitNominalTransaksi from TRSParamMaterai_TR where ProductCode = 'BOND')` : `DECIMAL(25,2)` | `TotalProceed` → `DECIMAL` | `TotalProceed >= (select top 1 StartLimitNominalTransaksi from TRSParamMaterai_TR where ProductCode = 'BOND')` |
| IC-010029 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptKonfirmasiBOND:219](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L219) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bb.TotalProceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(bb.TotalProceed,0)` |
| IC-010042 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptKonfirmasiBOND:241](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql#L241) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bb.TotalProceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(bb.TotalProceed,0)` |

<a id="kolom-transactionfee"></a>
## TransactionFee

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.BONDSDetailTransaksiBuyBackPerdana_TR.sql#L28). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006525 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateTrxBuyBackPerdana:222](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L222) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `TransactionFee` : `MONEY` | `0` : `int` | `0` → `MONEY` | `TransactionFee = 0` |
