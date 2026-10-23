# Retail UAT.dbo.TRSParameter_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 26 lokasi source, 23 procedure, 2 kolom. **A:** 8 lokasi pada 8 procedure. **B:** 17 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-value"></a>
## Value

**Tipe dasar:** `VARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParameter_TR.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 25 lokasi / 22 SP. **A:** 7 lokasi / 7 SP. **Peran:** CONVERTED_COLUMN, COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002970 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.SRBIProcessEmailNotifMaturityDateH1:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH1.sql#L61) | JOIN_ON / P01 / COMPARISON | `b.Value` : `VARCHAR(255)` | `a.SecType` : `SMALLINT` | `b.Value` → `SMALLINT` | `a.SecType = b.Value` |
| IC-002981 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.SRBIProcessEmailNotifMaturityDateH3:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH3.sql#L61) | JOIN_ON / P01 / COMPARISON | `b.Value` : `VARCHAR(255)` | `a.SecType` : `SMALLINT` | `b.Value` → `SMALLINT` | `a.SecType = b.Value` |
| IC-002992 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.SRBIProcessEmailNotifMaturityDateH7:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH7.sql#L61) | JOIN_ON / P01 / COMPARISON | `b.Value` : `VARCHAR(255)` | `a.SecType` : `SMALLINT` | `b.Value` → `SMALLINT` | `a.SecType = b.Value` |
| IC-006010 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateSecurityMasterAuthorization:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSecurityMasterAuthorization.sql#L192) | JOIN_ON / P01 / COMPARISON | `pr.Value` : `VARCHAR(255)` | `smptm.ImbalHasil` : `SMALLINT` | `pr.Value` → `SMALLINT` | `smptm.ImbalHasil = pr.Value` |
| IC-009762 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptDetailTrx:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L37) | WHERE / P01 / COMPARISON | `Value` : `VARCHAR(255)` | `ISNULL(@pnTransactionInstruction,0)` : `INT` | `Value` → `INT` | `Value = ISNULL(@pnTransactionInstruction,0)` |
| IC-009862 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptDetailTrxTaxAmnesty:32](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L32) | WHERE / P01 / COMPARISON | `Value` : `VARCHAR(255)` | `ISNULL(@pnTransactionInstruction,0)` : `INT` | `Value` → `INT` | `Value = ISNULL(@pnTransactionInstruction,0)` |
| IC-011595 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptTransactionDetailWMTA:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L28) | WHERE / P01 / COMPARISON | `Value` : `VARCHAR(255)` | `ISNULL(@pnTransactionInstruction,0)` : `INT` | `Value` → `INT` | `Value = ISNULL(@pnTransactionInstruction,0)` |
| IC-000266 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.EBWInvestmentInquiry#ORI:32](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.EBWInvestmentInquiry%23ORI.sql#L32) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[Value]` : `VARCHAR(255)` | `@dEffDateOD1` : `datetime` | `[Value]` → `datetime` | `@dEffDateOD1 = [Value]` |
| IC-000284 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.EBWVeloInvestmentInquiry#ORI:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.EBWVeloInvestmentInquiry%23ORI.sql#L111) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[Value]` : `VARCHAR(255)` | `@dEffDateOD1` : `datetime` | `[Value]` → `datetime` | `@dEffDateOD1 = [Value]` |
| IC-002701 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.Rvmp_trs_ListSecurityMaster_TM:42](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L42) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[Value]` : `VARCHAR(255)` | `@pnBatasExp` : `int` | `[Value]` → `int` | `@pnBatasExp = [Value]` |
| IC-004135 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSGetParamDummyForUnblock:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetParamDummyForUnblock.sql#L36) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `Value` : `VARCHAR(255)` | `@dDeployDate` : `DATETIME` | `Value` → `DATETIME` | `@dDeployDate = Value` |
| IC-004137 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSGetRiskProfileProductAndAgeCustomer:14](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetRiskProfileProductAndAgeCustomer.sql#L14) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `Value` : `VARCHAR(255)` | `@pnBatasExp` : `int` | `Value` → `int` | `@pnBatasExp = Value` |
| IC-004449 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSMaintenanceParamHargaModal:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L51) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `Value` : `VARCHAR(255)` | `@nAddSettlementDay` : `int` | `Value` → `int` | `@nAddSettlementDay = Value` |
| IC-004450 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSMaintenanceParamHargaModal:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L52) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `Value` : `VARCHAR(255)` | `@nAddSettlementDay2` : `int` | `Value` → `int` | `@nAddSettlementDay2 = Value` |
| IC-007046 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSProcessOverBookingMateraiBuyBack:128](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql#L128) | EXPRESSION / P07 / COALESCE_BRANCH | `p.Value` : `VARCHAR(255)` | `result type` : `INT (combined result; facets unknown)` | `p.Value` → `INT` | `coalesce(d.BRANCH, p.Value)` |
| IC-008599 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSResponseBondFromMX:519](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L519) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `Value` : `VARCHAR(255)` | `@pnBatasExp` : `int` | `Value` → `int` | `@pnBatasExp = Value` |
| IC-009764 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptDetailTrx:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L45) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `Value` : `VARCHAR(255)` | `@nSecType` : `INT` | `Value` → `INT` | `@nSecType = Value` |
| IC-011597 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptTransactionDetailWMTA:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L36) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `Value` : `VARCHAR(255)` | `@nSecType` : `INT` | `Value` → `INT` | `@nSecType = Value` |
| IC-012724 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSValidateSecAccNo:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateSecAccNo.sql#L52) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `Value` : `VARCHAR(255)` | `@nYearRiskProfile` : `int` | `Value` → `int` | `@nYearRiskProfile = Value` |
| IC-012725 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSValidateSecAccNoOrderSuratBerharga:20](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateSecAccNoOrderSuratBerharga.sql#L20) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `Value` : `VARCHAR(255)` | `@pnTahun` : `int` | `Value` → `int` | `@pnTahun = Value` |
| IC-013291 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_ListSecurityMasterPending_TM:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L34) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `Value` : `VARCHAR(255)` | `@pnBatasExp` : `int` | `Value` → `int` | `@pnBatasExp = Value` |
| IC-013333 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_ListSecurityMaster_TM:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L26) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `Value` : `VARCHAR(255)` | `@pnBatasExp` : `int` | `Value` → `int` | `@pnBatasExp = Value` |
| IC-013839 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_UpdateStatusSecurityTransaction_TT:266](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransaction_TT.sql#L266) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `Value` : `VARCHAR(255)` | `@nBatasExp` : `INT` | `Value` → `INT` | `@nBatasExp = Value` |
| IC-013841 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.trs_UpdateTreasuryCustomerTMCauseRiskProfile:14](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomerTMCauseRiskProfile.sql#L14) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `Value` : `VARCHAR(255)` | `12` : `int` | `Value` → `INT` | `Value*12` |
| IC-005740 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateMinimumRiskProfile:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMinimumRiskProfile.sql#L25) | WHERE / P12 / COMPARISON | `prm.Value` : `VARCHAR(255)` | `@cRiskProfileCode` : `char(1)` | `@cRiskProfileCode` → `VARCHAR` | `prm.Value = @cRiskProfileCode` |

<a id="kolom-code"></a>
## Code

**Tipe dasar:** `VARCHAR(50)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParameter_TR.sql#L2). **Existing leading index:** TRSParameter_TR_CN1&#124;TRSParameter_TR_NC1.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005964 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateParameter:5](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParameter.sql#L5) | WHERE / P02 / COMPARISON | `Code` : `VARCHAR(50)` | `@pcParamCode` : `nvarchar(max)` | `Code` → `NVARCHAR` | `Code = @pcParamCode` |
