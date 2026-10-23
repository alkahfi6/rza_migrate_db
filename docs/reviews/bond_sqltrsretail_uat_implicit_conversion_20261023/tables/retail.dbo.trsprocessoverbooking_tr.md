# Retail UAT.dbo.TRSProcessOverBooking_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 15 lokasi source, 5 procedure, 3 kolom. **A:** 2 lokasi pada 2 procedure. **B:** 2 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `VARCHAR(3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSProcessOverBooking_TR.sql#L32). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 2 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008389 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSReportSKFBankBeli:93](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeli.sql#L93) | WHERE / P01 / COMPARISON | `ob.Status` : `VARCHAR(3)` | `3` : `int` | `ob.Status` → `INT` | `ob.Status = 3` |
| IC-008405 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSReportSKFBankBeliBKON:132](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeliBKON.sql#L132) | WHERE / P01 / COMPARISON | `ob.Status` : `VARCHAR(3)` | `3` : `int` | `ob.Status` → `INT` | `ob.Status = 3` |
| IC-012647 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateStatusProcessTransaksi:50](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateStatusProcessTransaksi.sql#L50) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[Status]` : `VARCHAR(3)` | `CASE @pbIsSuccess WHEN 1 THEN 3 WHEN 0 THEN 2 END` : `INT (combined result; facets unknown)` | `CASE @pbIsSuccess WHEN 1 THEN 3 WHEN 0 THEN 2 END` → `VARCHAR(3)` | `[Status] = CASE @pbIsSuccess WHEN 1 THEN 3 WHEN 0 THEN 2 END` |

<a id="kolom-creditamount"></a>
## CreditAmount

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSProcessOverBooking_TR.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008390 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportSKFBankBeli:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeli.sql#L98) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ob.CreditAmount` : `MONEY` | `DPP` : `DECIMAL(25,2)` | `ob.CreditAmount` → `DECIMAL(25,2)` | `DPP = ob.CreditAmount` |
| IC-008391 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportSKFBankBeli:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeli.sql#L105) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ob.CreditAmount` : `MONEY` | `PPN` : `DECIMAL(25,2)` | `ob.CreditAmount` → `DECIMAL(25,2)` | `PPN = ob.CreditAmount` |

<a id="kolom-trxid"></a>
## TrxId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSProcessOverBooking_TR.sql#L3). **Existing leading index:** TRSProcessOverBooking_TR_CN1.

**Frekuensi source:** 10 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007040 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessOverBooking:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBooking.sql#L156) | WHERE / P11 / COMPARISON | `TrxId` : `BIGINT` | `@TrxId` : `INT` | `@TrxId` → `BIGINT` | `TrxId = @TrxId` |
| IC-007041 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessOverBooking:194](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBooking.sql#L194) | WHERE / P11 / COMPARISON | `TrxId` : `BIGINT` | `@TrxId` : `INT` | `@TrxId` → `BIGINT` | `TrxId = @TrxId` |
| IC-007042 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessOverBooking:199](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBooking.sql#L199) | WHERE / P11 / COMPARISON | `TrxId` : `BIGINT` | `@TrxId` : `INT` | `@TrxId` → `BIGINT` | `TrxId = @TrxId` |
| IC-007044 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessOverBooking:232](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBooking.sql#L232) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsprocessoverbooking_tr.TrxId` : `BIGINT` | `@TrxId` : `INT` | `@TrxId` → `BIGINT` | `retail.dbo.trsprocessoverbooking_tr.TrxId <- @TrxId` |
| IC-007045 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessOverBooking:291](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBooking.sql#L291) | WHERE / P11 / COMPARISON | `TrxId` : `BIGINT` | `@TrxId` : `INT` | `@TrxId` → `BIGINT` | `TrxId = @TrxId` |
| IC-007079 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessOverBookingWrapper:155](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingWrapper.sql#L155) | WHERE / P11 / COMPARISON | `TrxId` : `BIGINT` | `@TrxId` : `INT` | `@TrxId` → `BIGINT` | `TrxId = @TrxId` |
| IC-007080 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessOverBookingWrapper:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingWrapper.sql#L193) | WHERE / P11 / COMPARISON | `TrxId` : `BIGINT` | `@TrxId` : `INT` | `@TrxId` → `BIGINT` | `TrxId = @TrxId` |
| IC-007081 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessOverBookingWrapper:198](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingWrapper.sql#L198) | WHERE / P11 / COMPARISON | `TrxId` : `BIGINT` | `@TrxId` : `INT` | `@TrxId` → `BIGINT` | `TrxId = @TrxId` |
| IC-007083 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessOverBookingWrapper:232](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingWrapper.sql#L232) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsprocessoverbooking_tr.TrxId` : `BIGINT` | `@TrxId` : `INT` | `@TrxId` → `BIGINT` | `retail.dbo.trsprocessoverbooking_tr.TrxId <- @TrxId` |
| IC-007084 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessOverBookingWrapper:291](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingWrapper.sql#L291) | WHERE / P11 / COMPARISON | `TrxId` : `BIGINT` | `@TrxId` : `INT` | `@TrxId` → `BIGINT` | `TrxId = @TrxId` |
