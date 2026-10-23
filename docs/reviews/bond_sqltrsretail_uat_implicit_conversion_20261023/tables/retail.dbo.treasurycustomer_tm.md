# Retail UAT.dbo.TreasuryCustomer_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 245 lokasi source, 83 procedure, 34 kolom. **A:** 7 lokasi pada 5 procedure. **B:** 43 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cifno"></a>
## CIFNo

**Tipe dasar:** `VARCHAR(19)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L4). **Existing leading index:** TreasuryCustomer_TM_NN1&#124;TreasuryCustomer_TM_NN4.

**Frekuensi source:** 16 lokasi / 11 SP. **A:** 3 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN, COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004113 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSGetAllProductName:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetAllProductName.sql#L13) | WHERE / P01 / COMPARISON | `t.CIFNo` : `VARCHAR(19)` | `@pnCIFNo` : `bigint` | `t.CIFNo` → `BIGINT` | `t.CIFNo = @pnCIFNo` |
| IC-013421 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.trs_ListStatusORI_Order_TT:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L66) | WHERE / P01 / COMPARISON | `tc.CIFNo` : `VARCHAR(19)` | `NoCIF` : `BIGINT` | `tc.CIFNo` → `BIGINT` | `NoCIF = tc.CIFNo` |
| IC-013428 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.trs_ListStatusORI_Order_TT:143](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L143) | WHERE / P01 / COMPARISON | `tc.CIFNo` : `VARCHAR(19)` | `NoCIF` : `BIGINT` | `tc.CIFNo` → `BIGINT` | `NoCIF = tc.CIFNo` |
| IC-013596 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_PopulateDataRekeningMaterai:20](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningMaterai.sql#L20) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CIFNo` : `VARCHAR(19)` | `@nCIFNo` : `BIGINT` | `CIFNo` → `BIGINT` | `@nCIFNo = CIFNo` |
| IC-013617 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_PopulateDataRekeningObligasi:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningObligasi.sql#L22) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CIFNo` : `VARCHAR(19)` | `@nCIFNo` : `bigint` | `CIFNo` → `bigint` | `@nCIFNo = CIFNo` |
| IC-000264 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.DFSListObliTransaction:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.DFSListObliTransaction.sql#L24) | WHERE / P12 / COMPARISON | `b.CIFNo` : `VARCHAR(19)` | `@cCIFNo` : `char(19)` | `@cCIFNo` → `VARCHAR` | `b.CIFNo = @cCIFNo` |
| IC-000265 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.DFSListObliTransaction:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.DFSListObliTransaction.sql#L36) | WHERE / P12 / COMPARISON | `b.CIFNo` : `VARCHAR(19)` | `@cCIFNo` : `char(19)` | `@cCIFNo` → `VARCHAR` | `b.CIFNo = @cCIFNo` |
| IC-004427 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSJobSyncKITAS:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSJobSyncKITAS.sql#L25) | JOIN_ON / P12 / COMPARISON | `tc.[CIFNo]` : `VARCHAR(19)` | `cf.[CFCIF]` : `CHAR(19)` | `cf.[CFCIF]` → `VARCHAR` | `tc.[CIFNo] = cf.[CFCIF]` |
| IC-004434 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSJobSyncKITAS:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSJobSyncKITAS.sql#L47) | JOIN_ON / P12 / COMPARISON | `tc.[CIFNo]` : `VARCHAR(19)` | `cf.[CFCIF]` : `CHAR(19)` | `cf.[CFCIF]` → `VARCHAR` | `tc.[CIFNo] = cf.[CFCIF]` |
| IC-008146 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportFormBuySellObli:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportFormBuySellObli.sql#L110) | JOIN_ON / P12 / COMPARISON | `tctm.CIFNo` : `VARCHAR(19)` | `addr.CFCIF` : `CHAR(19)` | `addr.CFCIF` → `VARCHAR` | `tctm.CIFNo = addr.CFCIF` |
| IC-009591 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptCustomerPortfolio:88](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerPortfolio.sql#L88) | JOIN_ON / P12 / COMPARISON | `tctm.CIFNo` : `VARCHAR(19)` | `CFMAST.CFCIF` : `CHAR(19)` | `CFMAST.CFCIF` → `VARCHAR` | `CFMAST.CFCIF = tctm.CIFNo` |
| IC-009592 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptCustomerPortfolio:90](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerPortfolio.sql#L90) | JOIN_ON / P12 / COMPARISON | `tctm.CIFNo` : `VARCHAR(19)` | `CFADDR.CFCIF` : `CHAR(19)` | `CFADDR.CFCIF` → `VARCHAR` | `CFADDR.CFCIF = tctm.CIFNo` |
| IC-010607 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPTOSNasabah:186](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPTOSNasabah.sql#L186) | JOIN_ON / P12 / COMPARISON | `tc.CIFNo` : `VARCHAR(19)` | `addr.CFCIF` : `CHAR(19)` | `addr.CFCIF` → `VARCHAR` | `tc.CIFNo = addr.CFCIF` |
| IC-012415 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSyncMasterRiskProfileData:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterRiskProfileData.sql#L21) | JOIN_ON / P12 / COMPARISON | `tc.CIFNo` : `VARCHAR(19)` | `CFMAST.CFCIF` : `CHAR(19)` | `CFMAST.CFCIF` → `VARCHAR` | `tc.CIFNo = CFMAST.CFCIF` |
| IC-013559 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_ListTreasuryCustomer_TM_Original:236](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListTreasuryCustomer_TM_Original.sql#L236) | JOIN_ON / P12 / COMPARISON | `tctm.CIFNo` : `VARCHAR(19)` | `cfm.CFCIF` : `CHAR(19)` | `cfm.CFCIF` → `VARCHAR` | `tctm.CIFNo = cfm.CFCIF` |
| IC-013565 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_ListTreasuryCustomer_TM_Original:402](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListTreasuryCustomer_TM_Original.sql#L402) | JOIN_ON / P12 / COMPARISON | `tctm.CIFNo` : `VARCHAR(19)` | `cfm.CFCIF` : `CHAR(19)` | `cfm.CFCIF` → `VARCHAR` | `tctm.CIFNo = cfm.CFCIF` |

<a id="kolom-kodecabangalamatsurat"></a>
## KodeCabangAlamatSurat

**Tipe dasar:** `VARCHAR(5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L51). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 2 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-010325 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptObligasiKonfirmasiBond:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptObligasiKonfirmasiBond.sql#L49) | JOIN_ON / P01 / COMPARISON | `tctm.KodeCabangAlamatSurat` : `VARCHAR(5)` | `office.office_id` : `INT` | `tctm.KodeCabangAlamatSurat` → `INT` | `tctm.KodeCabangAlamatSurat = office.office_id` |
| IC-010609 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptPTOSNasabah:189](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPTOSNasabah.sql#L189) | JOIN_ON / P01 / COMPARISON | `tc.KodeCabangAlamatSurat` : `VARCHAR(5)` | `mis.office_id` : `INT` | `tc.KodeCabangAlamatSurat` → `INT` | `tc.KodeCabangAlamatSurat = mis.office_id` |

<a id="kolom-norekinvestor"></a>
## NoRekInvestor

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 2 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000597 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabahTRS:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L51) | JOIN_ON / P01 / COMPARISON | `a.NoRekInvestor` : `VARCHAR(20)` | `c.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = c.ACCTNO` |
| IC-000617 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabahTRS:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L180) | JOIN_ON / P01 / COMPARISON | `a.NoRekInvestor` : `VARCHAR(20)` | `c.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = c.ACCTNO` |
| IC-008623 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBondFromMX:696](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L696) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tc.NoRekInvestor` : `VARCHAR(20)` | `fm.NoRekInvestor` : `BIGINT` | `fm.NoRekInvestor` → `VARCHAR(20)` | `tc.NoRekInvestor = fm.NoRekInvestor` |

<a id="kolom-alamatsuratsequence"></a>
## AlamatSuratSequence

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L49). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013590 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_PopulateAlamatNasabah:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateAlamatNasabah.sql#L13) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `AlamatSuratSequence` : `INT` | `@cAlamatSuratSequence` : `char(3)` | `AlamatSuratSequence` → `char(3)` | `@cAlamatSuratSequence = AlamatSuratSequence` |
| IC-008147 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportFormBuySellObli:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportFormBuySellObli.sql#L111) | JOIN_ON / P01 / COMPARISON | `tctm.AlamatSuratSequence` : `INT` | `addr.CFASEQ` : `CHAR(3)` | `addr.CFASEQ` → `INT` | `tctm.AlamatSuratSequence = addr.CFASEQ` |
| IC-010324 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptObligasiKonfirmasiBond:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptObligasiKonfirmasiBond.sql#L46) | JOIN_ON / P01 / COMPARISON | `tctm.AlamatSuratSequence` : `INT` | `addr.CFASEQ` : `CHAR(3)` | `addr.CFASEQ` → `INT` | `tctm.AlamatSuratSequence = addr.CFASEQ` |
| IC-010608 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPTOSNasabah:187](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPTOSNasabah.sql#L187) | JOIN_ON / P01 / COMPARISON | `tc.AlamatSuratSequence` : `INT` | `addr.CFASEQ` : `CHAR(3)` | `addr.CFASEQ` → `INT` | `tc.AlamatSuratSequence = addr.CFASEQ` |

<a id="kolom-approvalnnh"></a>
## ApprovalNNH

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L68). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007870 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahObligasi:366](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahObligasi.sql#L366) | CASE / P11 / COMPARISON | `ApprovalNNH` : `BIT` | `1` : `int` | `ApprovalNNH` → `INT` | `ApprovalNNH = 1` |
| IC-013558 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListTreasuryCustomer_TM_Original:197](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListTreasuryCustomer_TM_Original.sql#L197) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.ApprovalNNH` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(tctm.ApprovalNNH, 0)` |
| IC-013563 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListTreasuryCustomer_TM_Original:366](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListTreasuryCustomer_TM_Original.sql#L366) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.ApprovalNNH` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(tctm.ApprovalNNH, 0)` |

<a id="kolom-branchcode"></a>
## BranchCode

**Tipe dasar:** `VARCHAR(5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT.

**Pola dan arah review:**

- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005250 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSPopulateApprovalBuyBackOM:57](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalBuyBackOM.sql#L57) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `tc.BranchCode` : `VARCHAR(5)` | `retail.dbo.fngetbranchcode.@cBranchCode` : `int` | `tc.BranchCode` → `int` | `dbo.fnGetBranchCode(tc.BranchCode)` |
| IC-013422 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.trs_ListStatusORI_Order_TT:78](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L78) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `tc.BranchCode` : `VARCHAR(5)` | `retail.dbo.fngetbranchcode.@cBranchCode` : `int` | `tc.BranchCode` → `int` | `dbo.fnGetBranchCode(tc.BranchCode)` |
| IC-013429 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.trs_ListStatusORI_Order_TT:155](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L155) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `tc.BranchCode` : `VARCHAR(5)` | `retail.dbo.fngetbranchcode.@cBranchCode` : `int` | `tc.BranchCode` → `int` | `dbo.fnGetBranchCode(tc.BranchCode)` |
| IC-013455 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.trs_ListStatusORI_Order_Temp:42](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_Temp.sql#L42) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `tc.BranchCode` : `VARCHAR(5)` | `retail.dbo.fngetbranchcode.@cBranchCode` : `int` | `tc.BranchCode` → `int` | `dbo.fnGetBranchCode(tc.BranchCode)` |
| IC-013459 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.trs_ListStatusORI_Order_Temp:87](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_Temp.sql#L87) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `tc.BranchCode` : `VARCHAR(5)` | `retail.dbo.fngetbranchcode.@cBranchCode` : `int` | `tc.BranchCode` → `int` | `dbo.fnGetBranchCode(tc.BranchCode)` |

<a id="kolom-cifid"></a>
## CIFId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L2). **Existing leading index:** IU_TreasuryCustomer_TM&#124;PK__TreasuryCustomer__7A672E12.

**Frekuensi source:** 32 lokasi / 19 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000153 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.CustodyReportMaintainAcquisition:118](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportMaintainAcquisition.sql#L118) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `CIFId` : `BIGINT` | `#tmp_TreasuryCustomer_TM.CIFId` : `VARCHAR(30)` | `CIFId` → `VARCHAR(30)` | `#tmp_TreasuryCustomer_TM.CIFId <- CIFId` |
| IC-003484 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSCalculateDailyAccrued:87](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateDailyAccrued.sql#L87) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tc.CIFId` : `BIGINT` | `#tmp_TreasuryCustomerBalance.CIFId` : `int` | `tc.CIFId` → `int` | `#tmp_TreasuryCustomerBalance.CIFId <- tc.CIFId` |
| IC-005976 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateReport008:170](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateReport008.sql#L170) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CIFId` : `BIGINT` | `@nCIFId` : `int` | `CIFId` → `int` | `@nCIFId = CIFId` |
| IC-013150 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_InsertTreasuryCustomer_TM:529](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L529) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `max(CIFId)` : `BIGINT` | `@nCounter1` : `int` | `max(CIFId)` → `int` | `@nCounter1 = max(CIFId)` |
| IC-000804 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.OMValidateBalance:75](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateBalance.sql#L75) | WHERE / P11 / COMPARISON | `tb.CIFId` : `BIGINT` | `@pnCIFId` : `int` | `@pnCIFId` → `BIGINT` | `tb.CIFId = @pnCIFId` |
| IC-000805 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.OMValidateBalance:79](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateBalance.sql#L79) | WHERE / P11 / COMPARISON | `tb.CIFId` : `BIGINT` | `@pnCIFId` : `int` | `@pnCIFId` → `BIGINT` | `tb.CIFId = @pnCIFId` |
| IC-000892 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.ONEApprovalPrepareData:60](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareData.sql#L60) | JOIN_ON / P01 / COMPARISON | `obliNewMasterNasabah.CIFId` : `BIGINT` | `oap.SearchTerm` : `VARCHAR(100)` | `oap.SearchTerm` → `BIGINT` | `obliNewMasterNasabah.CIFId = oap.SearchTerm` |
| IC-000920 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.ONEApprovalPrepareDataBond:58](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareDataBond.sql#L58) | JOIN_ON / P01 / COMPARISON | `obliNewMasterNasabah.CIFId` : `BIGINT` | `oap.SearchTerm` : `VARCHAR(100)` | `oap.SearchTerm` → `BIGINT` | `obliNewMasterNasabah.CIFId = oap.SearchTerm` |
| IC-002775 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_trs_updateStatusTreasuryCustomer_TM:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_updateStatusTreasuryCustomer_TM.sql#L94) | JOIN_ON / P11 / COMPARISON | `Trx.CIFId` : `BIGINT` | `t.CIFId` : `int` | `t.CIFId` → `BIGINT` | `Trx.CIFId=t.CIFId` |
| IC-002778 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_trs_updateStatusTreasuryCustomer_TM:220](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_updateStatusTreasuryCustomer_TM.sql#L220) | JOIN_ON / P11 / COMPARISON | `tctm.CIFId` : `BIGINT` | `tt.CIFId` : `int` | `tt.CIFId` → `BIGINT` | `tctm.CIFId = tt.CIFId` |
| IC-002856 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedProcessMasterInvestor:329](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessMasterInvestor.sql#L329) | WHERE / P11 / COMPARISON | `[CIFId]` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `[CIFId] = @nCIFId` |
| IC-002864 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedProcessRekDana:117](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L117) | WHERE / P11 / COMPARISON | `tc.CIFId` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `tc.CIFId = @nCIFId` |
| IC-002872 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedProcessRekDana:384](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L384) | WHERE / P11 / COMPARISON | `CIFId` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `CIFId = @nCIFId` |
| IC-002885 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedProcessRekDana:649](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L649) | WHERE / P11 / COMPARISON | `CIFId` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `CIFId = @nCIFId` |
| IC-002889 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedProcessRekDana:847](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L847) | WHERE / P11 / COMPARISON | `CIFId` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `CIFId = @nCIFId` |
| IC-002911 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedPushMailNotifOrderCompleted:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L169) | WHERE / P11 / COMPARISON | `[CIFId]` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `[CIFId] = @nCIFId` |
| IC-002935 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedPushMailNotifOrderORI:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderORI.sql#L162) | WHERE / P11 / COMPARISON | `[CIFId]` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `[CIFId] = @nCIFId` |
| IC-002951 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedPushMailNotifRegInvestor:75](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifRegInvestor.sql#L75) | WHERE / P11 / COMPARISON | `[CIFId]` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `[CIFId] = @nCIFId` |
| IC-002960 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedPushMailNotifSettleRedempt:95](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifSettleRedempt.sql#L95) | WHERE / P11 / COMPARISON | `[CIFId]` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `[CIFId] = @nCIFId` |
| IC-009318 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailSyncDataCoreWorkingHours:235](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L235) | WHERE / P05 / BINARY_Modulo | `tc.CIFId` : `BIGINT` | `6` : `int` | `6` → `BIGINT` | `tc.CIFId % 6` |
| IC-009319 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailSyncDataCoreWorkingHours:235](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L235) | WHERE / P11 / COMPARISON | `(tc.CIFId % 6)` : `BIGINT (combined result; facets unknown)` | `@pResultMod` : `INT` | `@pResultMod` → `BIGINT` | `(tc.CIFId % 6) = @pResultMod` |
| IC-012469 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSUpdateDataMasterNasabah:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateDataMasterNasabah.sql#L110) | WHERE / P05 / BINARY_Modulo | `CIFId` : `BIGINT` | `@pPembagi` : `INT` | `@pPembagi` → `BIGINT` | `CIFId % @pPembagi` |
| IC-012470 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSUpdateDataMasterNasabah:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateDataMasterNasabah.sql#L110) | WHERE / P11 / COMPARISON | `(CIFId % @pPembagi)` : `BIGINT (combined result; facets unknown)` | `@pResultMod` : `INT` | `@pResultMod` → `BIGINT` | `(CIFId % @pPembagi) = @pResultMod` |
| IC-013147 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_InsertTreasuryCustomer_TM:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L53) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `isnull((SELECT MAX(CIFId) FROM TreasuryCustomer_TM WITH (UPDLOCK, HOLDLOCK)), 0)` : `BIGINT` | `1` : `int` | `1` → `BIGINT` | `1 + isnull((SELECT MAX(CIFId) FROM TreasuryCustomer_TM WITH (UPDLOCK, HOLDLOCK)), 0)` |
| IC-013148 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_InsertTreasuryCustomer_TM:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L53) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `(SELECT MAX(CIFId) FROM TreasuryCustomer_TM WITH (UPDLOCK, HOLDLOCK))` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `isnull((SELECT MAX(CIFId) FROM TreasuryCustomer_TM WITH (UPDLOCK, HOLDLOCK)), 0)` |
| IC-013708 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_StatusSecurityTransaction2nd_TT:93](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_StatusSecurityTransaction2nd_TT.sql#L93) | JOIN_ON / P11 / COMPARISON | `tc.CIFId` : `BIGINT` | `tt.CIFIdDebet` : `int` | `tt.CIFIdDebet` → `BIGINT` | `tt.CIFIdDebet = tc.CIFId` |
| IC-013709 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_StatusSecurityTransaction2nd_TT:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_StatusSecurityTransaction2nd_TT.sql#L96) | JOIN_ON / P11 / COMPARISON | `tc.CIFId` : `BIGINT` | `tt.CIFIdKredit` : `int` | `tt.CIFIdKredit` → `BIGINT` | `tt.CIFIdKredit = tc.CIFId` |
| IC-013718 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_StatusSecurityTransaction2nd_TT:134](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_StatusSecurityTransaction2nd_TT.sql#L134) | JOIN_ON / P11 / COMPARISON | `tc.CIFId` : `BIGINT` | `tt.CIFIdDebet` : `int` | `tt.CIFIdDebet` → `BIGINT` | `tt.CIFIdDebet = tc.CIFId` |
| IC-013719 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_StatusSecurityTransaction2nd_TT:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_StatusSecurityTransaction2nd_TT.sql#L137) | JOIN_ON / P11 / COMPARISON | `tc.CIFId` : `BIGINT` | `tt.CIFIdKredit` : `int` | `tt.CIFIdKredit` → `BIGINT` | `tt.CIFIdKredit = tc.CIFId` |
| IC-013793 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_UpdateStatusSecurityTransaction2nd_TT:151](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransaction2nd_TT.sql#L151) | WHERE / P11 / COMPARISON | `tc.CIFId` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `tc.CIFId =@nCIFId` |
| IC-013992 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_updateStatusTreasuryCustomer_TM:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateStatusTreasuryCustomer_TM.sql#L94) | JOIN_ON / P11 / COMPARISON | `Trx.CIFId` : `BIGINT` | `t.CIFId` : `int` | `t.CIFId` → `BIGINT` | `Trx.CIFId=t.CIFId` |
| IC-013995 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_updateStatusTreasuryCustomer_TM:220](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateStatusTreasuryCustomer_TM.sql#L220) | JOIN_ON / P11 / COMPARISON | `tctm.CIFId` : `BIGINT` | `tt.CIFId` : `int` | `tt.CIFId` → `BIGINT` | `tctm.CIFId = tt.CIFId` |

<a id="kolom-citizenship"></a>
## Citizenship

**Tipe dasar:** `CHAR(3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L42). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 72 lokasi / 28 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, DESTINATION_COLUMN, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009149 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailSyncDataCore:222](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L222) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tc.Citizenship` : `CHAR(3)` | `#xtmpTrsCustomer.Citizenship` : `VARCHAR(20)` | `tc.Citizenship` → `VARCHAR(20)` | `#xtmpTrsCustomer.Citizenship <- tc.Citizenship` |
| IC-009152 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailSyncDataCore:222](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L222) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tc.Citizenship` : `CHAR(3)` | `#xtmpTrsCustomer.newCitizenship` : `VARCHAR(20)` | `tc.Citizenship` → `VARCHAR(20)` | `#xtmpTrsCustomer.newCitizenship <- tc.Citizenship` |
| IC-009314 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailSyncDataCoreWorkingHours:123](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L123) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tc.Citizenship` : `CHAR(3)` | `#xtmpTrsCustomer.Citizenship` : `VARCHAR(20)` | `tc.Citizenship` → `VARCHAR(20)` | `#xtmpTrsCustomer.Citizenship <- tc.Citizenship` |
| IC-009317 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailSyncDataCoreWorkingHours:123](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L123) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tc.Citizenship` : `CHAR(3)` | `#xtmpTrsCustomer.newCitizenship` : `VARCHAR(20)` | `tc.Citizenship` → `VARCHAR(20)` | `#xtmpTrsCustomer.newCitizenship <- tc.Citizenship` |
| IC-000750 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMTrxSecPreBook:392](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMTrxSecPreBook.sql#L392) | CASE / P12 / COMPARISON | `Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `Citizenship` → `VARCHAR` | `Citizenship = '000'` |
| IC-000751 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMTrxSecPreBook:392](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMTrxSecPreBook.sql#L392) | CASE / P12 / COMPARISON | `Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `Citizenship` → `VARCHAR` | `Citizenship != '000'` |
| IC-000752 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.OMTrxSecPreBook:392](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMTrxSecPreBook.sql#L392) | CASE / P12 / COMPARISON | `ISNULL(Citizenship ,'')` : `CHAR(3)` | `''` : `varchar` | `ISNULL(Citizenship ,'')` → `VARCHAR` | `ISNULL(Citizenship ,'') = ''` |
| IC-003136 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SecurityProcessResultLelang:366](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessResultLelang.sql#L366) | CASE / P12 / COMPARISON | `Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `Citizenship` → `VARCHAR` | `Citizenship = '000'` |
| IC-003137 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SecurityProcessResultLelang:366](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessResultLelang.sql#L366) | CASE / P12 / COMPARISON | `Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `Citizenship` → `VARCHAR` | `Citizenship != '000'` |
| IC-003138 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.SecurityProcessResultLelang:366](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessResultLelang.sql#L366) | CASE / P12 / COMPARISON | `ISNULL(Citizenship ,'')` : `CHAR(3)` | `''` : `varchar` | `ISNULL(Citizenship ,'')` → `VARCHAR` | `ISNULL(Citizenship ,'') = ''` |
| IC-003176 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SecurityProcessResultLelangTA:367](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessResultLelangTA.sql#L367) | CASE / P12 / COMPARISON | `Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `Citizenship` → `VARCHAR` | `Citizenship = '000'` |
| IC-003177 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SecurityProcessResultLelangTA:367](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessResultLelangTA.sql#L367) | CASE / P12 / COMPARISON | `Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `Citizenship` → `VARCHAR` | `Citizenship != '000'` |
| IC-003178 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.SecurityProcessResultLelangTA:367](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessResultLelangTA.sql#L367) | CASE / P12 / COMPARISON | `ISNULL(Citizenship ,'')` : `CHAR(3)` | `''` : `varchar` | `ISNULL(Citizenship ,'')` → `VARCHAR` | `ISNULL(Citizenship ,'') = ''` |
| IC-003676 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:199](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L199) | CASE / P12 / COMPARISON | `tc.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tc.Citizenship` → `VARCHAR` | `tc.Citizenship = '000'` |
| IC-003677 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:200](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L200) | CASE / P12 / COMPARISON | `tc.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tc.Citizenship` → `VARCHAR` | `tc.Citizenship <> '000'` |
| IC-003681 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:242](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L242) | CASE / P12 / COMPARISON | `tc.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tc.Citizenship` → `VARCHAR` | `tc.Citizenship <> '000'` |
| IC-003682 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:251](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L251) | WHERE / P12 / COMPARISON | `tc.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tc.Citizenship` → `VARCHAR` | `tc.Citizenship <> '000'` |
| IC-003773 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L162) | CASE / P12 / COMPARISON | `tc.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tc.Citizenship` → `VARCHAR` | `tc.Citizenship = '000'` |
| IC-003774 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:163](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L163) | CASE / P12 / COMPARISON | `tc.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tc.Citizenship` → `VARCHAR` | `tc.Citizenship <> '000'` |
| IC-003778 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L193) | CASE / P12 / COMPARISON | `tc.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tc.Citizenship` → `VARCHAR` | `tc.Citizenship <> '000'` |
| IC-003779 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:202](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L202) | WHERE / P12 / COMPARISON | `tc.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tc.Citizenship` → `VARCHAR` | `tc.Citizenship <> '000'` |
| IC-003931 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusiOrchid:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L162) | CASE / P12 / COMPARISON | `tc.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tc.Citizenship` → `VARCHAR` | `tc.Citizenship = '000'` |
| IC-003932 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusiOrchid:163](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L163) | CASE / P12 / COMPARISON | `tc.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tc.Citizenship` → `VARCHAR` | `tc.Citizenship <> '000'` |
| IC-004772 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSOBLExpiredKitasRPT:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOBLExpiredKitasRPT.sql#L104) | EXPRESSION / P07 / COALESCE_BRANCH | `TC.Citizenship` : `CHAR(3)` | `result type` : `VARCHAR (combined result; facets unknown)` | `TC.Citizenship` → `VARCHAR` | `COALESCE(TC.Citizenship, '')` |
| IC-004773 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSOBLExpiredKitasRPT:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOBLExpiredKitasRPT.sql#L105) | EXPRESSION / P07 / COALESCE_BRANCH | `TC.Citizenship` : `CHAR(3)` | `result type` : `VARCHAR (combined result; facets unknown)` | `TC.Citizenship` → `VARCHAR` | `COALESCE(TC.Citizenship, '')` |
| IC-007801 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSREPOReportBiayaMaterai:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L96) | CASE / P12 / COMPARISON | `Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `Citizenship` → `VARCHAR` | `Citizenship != '000'` |
| IC-007803 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSREPOReportBiayaMaterai:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L106) | CASE / P12 / COMPARISON | `Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `Citizenship` → `VARCHAR` | `Citizenship != '000'` |
| IC-007820 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSREPOReportBiayaMaterai:183](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L183) | CASE / P12 / COMPARISON | `Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `Citizenship` → `VARCHAR` | `Citizenship != '000'` |
| IC-007822 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSREPOReportBiayaMaterai:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L193) | CASE / P12 / COMPARISON | `Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `Citizenship` → `VARCHAR` | `Citizenship != '000'` |
| IC-007835 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSREPOReportBiayaMaterai:254](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L254) | CASE / P12 / COMPARISON | `Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `Citizenship` → `VARCHAR` | `Citizenship != '000'` |
| IC-007837 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSREPOReportBiayaMaterai:264](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L264) | CASE / P12 / COMPARISON | `Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `Citizenship` → `VARCHAR` | `Citizenship != '000'` |
| IC-007980 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTUplaodCTPBySID:8](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTUplaodCTPBySID.sql#L8) | CASE / P12 / COMPARISON | `Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `Citizenship` → `VARCHAR` | `Citizenship = '000'` |
| IC-008397 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSKFBankBeliBKON:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeliBKON.sql#L98) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `tc.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tc.Citizenship` → `VARCHAR` | `WHEN '000' THEN 'WNI'` |
| IC-008470 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTemplateSID:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSID.sql#L69) | CASE / P12 / COMPARISON | `tc.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tc.Citizenship` → `VARCHAR` | `tc.Citizenship = '000'` |
| IC-008472 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTemplateSID:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSID.sql#L77) | CASE / P12 / COMPARISON | `tc.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tc.Citizenship` → `VARCHAR` | `tc.Citizenship != '000'` |
| IC-008476 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTemplateSID:120](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSID.sql#L120) | JOIN_ON / P12 / COMPARISON | `tc.Citizenship` : `CHAR(3)` | `jhc.JHCCOC` : `VARCHAR(3)` | `tc.Citizenship` → `VARCHAR` | `tc.Citizenship = jhc.JHCCOC` |
| IC-008482 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTemplateSIDBySecAccNo:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSIDBySecAccNo.sql#L29) | CASE / P12 / COMPARISON | `tc.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tc.Citizenship` → `VARCHAR` | `tc.Citizenship = '000'` |
| IC-008484 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTemplateSIDBySecAccNo:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSIDBySecAccNo.sql#L37) | CASE / P12 / COMPARISON | `tc.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tc.Citizenship` → `VARCHAR` | `tc.Citizenship != '000'` |
| IC-008488 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTemplateSIDBySecAccNo:76](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSIDBySecAccNo.sql#L76) | JOIN_ON / P12 / COMPARISON | `tc.Citizenship` : `CHAR(3)` | `jhc.JHCCOC` : `VARCHAR(3)` | `tc.Citizenship` → `VARCHAR` | `tc.Citizenship = jhc.JHCCOC` |
| IC-008736 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRetailCalculateFee:231](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L231) | CASE / P12 / COMPARISON | `isnull(Citizenship,'')` : `CHAR(3)` | `'000'` : `varchar` | `isnull(Citizenship,'')` → `VARCHAR` | `isnull(Citizenship,'') = '000'` |
| IC-008738 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRetailCalculateFee:232](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L232) | CASE / P12 / COMPARISON | `isnull(Citizenship,'')` : `CHAR(3)` | `''` : `varchar` | `isnull(Citizenship,'')` → `VARCHAR` | `isnull(Citizenship,'') = ''` |
| IC-008740 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRetailCalculateFee:233](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L233) | CASE / P12 / COMPARISON | `isnull(Citizenship,'')` : `CHAR(3)` | `'000'` : `varchar` | `isnull(Citizenship,'')` → `VARCHAR` | `isnull(Citizenship,'') != '000'` |
| IC-009077 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRetailPopulateLCNew:341](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulateLCNew.sql#L341) | CASE / P12 / COMPARISON | `isnull(tc.Citizenship,'000')` : `CHAR(3)` | `'000'` : `varchar` | `isnull(tc.Citizenship,'000')` → `VARCHAR` | `isnull(tc.Citizenship,'000') = '000'` |
| IC-009100 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRetailReCalculateAccrued:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailReCalculateAccrued.sql#L22) | CASE / P12 / COMPARISON | `isnull(tctm.Citizenship,'')` : `CHAR(3)` | `'000'` : `varchar` | `isnull(tctm.Citizenship,'')` → `VARCHAR` | `isnull(tctm.Citizenship,'') = '000'` |
| IC-009102 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRetailReCalculateAccrued:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailReCalculateAccrued.sql#L23) | CASE / P12 / COMPARISON | `isnull(tctm.Citizenship,'')` : `CHAR(3)` | `''` : `varchar` | `isnull(tctm.Citizenship,'')` → `VARCHAR` | `isnull(tctm.Citizenship,'') = ''` |
| IC-009104 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRetailReCalculateAccrued:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailReCalculateAccrued.sql#L24) | CASE / P12 / COMPARISON | `isnull(tctm.Citizenship,'')` : `CHAR(3)` | `'000'` : `varchar` | `isnull(tctm.Citizenship,'')` → `VARCHAR` | `isnull(tctm.Citizenship,'') != '000'` |
| IC-010646 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPajakKupon:235](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPajakKupon.sql#L235) | CASE / P12 / COMPARISON | `tctm.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tctm.Citizenship` → `VARCHAR` | `tctm.Citizenship != '000'` |
| IC-010740 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORI:252](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORI.sql#L252) | CASE / P12 / COMPARISON | `tctm.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tctm.Citizenship` → `VARCHAR` | `tctm.Citizenship != '000'` |
| IC-010858 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORIDetail:497](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L497) | CASE / P12 / COMPARISON | `tctm.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tctm.Citizenship` → `VARCHAR` | `tctm.Citizenship != '000'` |
| IC-010869 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORIDetail:581](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L581) | CASE / P12 / COMPARISON | `tctm.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tctm.Citizenship` → `VARCHAR` | `tctm.Citizenship != '000'` |
| IC-010986 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnesty:417](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L417) | CASE / P12 / COMPARISON | `tctm.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tctm.Citizenship` → `VARCHAR` | `tctm.Citizenship != '000'` |
| IC-010999 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnesty:505](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L505) | CASE / P12 / COMPARISON | `tctm.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tctm.Citizenship` → `VARCHAR` | `tctm.Citizenship != '000'` |
| IC-011070 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnestyEOM:241](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyEOM.sql#L241) | CASE / P12 / COMPARISON | `tctm.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tctm.Citizenship` → `VARCHAR` | `tctm.Citizenship != '000'` |
| IC-011174 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:380](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L380) | CASE / P12 / COMPARISON | `tctm.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tctm.Citizenship` → `VARCHAR` | `tctm.Citizenship != '000'` |
| IC-011282 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnestySPT:232](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestySPT.sql#L232) | CASE / P12 / COMPARISON | `tctm.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tctm.Citizenship` → `VARCHAR` | `tctm.Citizenship != '000'` |
| IC-011379 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPayCouponORITaxAmnesty_Backdated:321](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty_Backdated.sql#L321) | CASE / P12 / COMPARISON | `tctm.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `tctm.Citizenship` → `VARCHAR` | `tctm.Citizenship != '000'` |
| IC-013474 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListStatusTreasuryCustomer_TM:84](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L84) | JOIN_ON / P12 / COMPARISON | `tc.Citizenship` : `CHAR(3)` | `jh.JHCCOC` : `VARCHAR(3)` | `tc.Citizenship` → `VARCHAR` | `jh.JHCCOC = tc.Citizenship` |
| IC-013485 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListStatusTreasuryCustomer_TM:144](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L144) | JOIN_ON / P12 / COMPARISON | `tc.Citizenship` : `CHAR(3)` | `jh.JHCCOC` : `VARCHAR(3)` | `tc.Citizenship` → `VARCHAR` | `jh.JHCCOC = tc.Citizenship` |
| IC-000753 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.OMTrxSecPreBook:392](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMTrxSecPreBook.sql#L392) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `Citizenship` : `CHAR(3)` | `''` : `varchar` | `''` → `CHAR(3)` | `ISNULL(Citizenship ,'')` |
| IC-003139 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SecurityProcessResultLelang:366](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessResultLelang.sql#L366) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `Citizenship` : `CHAR(3)` | `''` : `varchar` | `''` → `CHAR(3)` | `ISNULL(Citizenship ,'')` |
| IC-003179 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SecurityProcessResultLelangTA:367](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessResultLelangTA.sql#L367) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `Citizenship` : `CHAR(3)` | `''` : `varchar` | `''` → `CHAR(3)` | `ISNULL(Citizenship ,'')` |
| IC-005565 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMasterNasabahForMurex:144](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMasterNasabahForMurex.sql#L144) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.Citizenship` : `CHAR(3)` | `''` : `varchar` | `''` → `CHAR(3)` | `ISNULL(tctm.Citizenship, '')` |
| IC-008737 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRetailCalculateFee:231](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L231) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `Citizenship` : `CHAR(3)` | `''` : `varchar` | `''` → `CHAR(3)` | `isnull(Citizenship,'')` |
| IC-008739 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRetailCalculateFee:232](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L232) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `Citizenship` : `CHAR(3)` | `''` : `varchar` | `''` → `CHAR(3)` | `isnull(Citizenship,'')` |
| IC-008741 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRetailCalculateFee:233](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L233) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `Citizenship` : `CHAR(3)` | `''` : `varchar` | `''` → `CHAR(3)` | `isnull(Citizenship,'')` |
| IC-009078 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRetailPopulateLCNew:341](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailPopulateLCNew.sql#L341) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tc.Citizenship` : `CHAR(3)` | `'000'` : `varchar` | `'000'` → `CHAR(3)` | `isnull(tc.Citizenship,'000')` |
| IC-009101 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRetailReCalculateAccrued:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailReCalculateAccrued.sql#L22) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.Citizenship` : `CHAR(3)` | `''` : `varchar` | `''` → `CHAR(3)` | `isnull(tctm.Citizenship,'')` |
| IC-009103 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRetailReCalculateAccrued:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailReCalculateAccrued.sql#L23) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.Citizenship` : `CHAR(3)` | `''` : `varchar` | `''` → `CHAR(3)` | `isnull(tctm.Citizenship,'')` |
| IC-009105 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRetailReCalculateAccrued:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailReCalculateAccrued.sql#L24) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.Citizenship` : `CHAR(3)` | `''` : `varchar` | `''` → `CHAR(3)` | `isnull(tctm.Citizenship,'')` |
| IC-009309 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailSyncDataCore:831](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L831) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Citizenship` : `CHAR(3)` | `t.newCitizenship` : `VARCHAR(20)` | `t.newCitizenship` → `CHAR(3)` | `Citizenship = t.newCitizenship` |
| IC-009474 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:703](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L703) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Citizenship` : `CHAR(3)` | `t.newCitizenship` : `VARCHAR(20)` | `t.newCitizenship` → `CHAR(3)` | `Citizenship = t.newCitizenship` |
| IC-013165 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:669](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L669) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomer_tm.Citizenship` : `CHAR(3)` | `Citizenship` : `varchar(3)` | `Citizenship` → `CHAR(3)` | `retail.dbo.treasurycustomer_tm.Citizenship <- Citizenship` |

<a id="kolom-flagkaryawan"></a>
## FlagKaryawan

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L47). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 12 lokasi / 9 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009150 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailSyncDataCore:222](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L222) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tc.FlagKaryawan` : `BIT` | `#xtmpTrsCustomer.FlagKaryawan` : `CHAR(1)` | `tc.FlagKaryawan` → `CHAR(1)` | `#xtmpTrsCustomer.FlagKaryawan <- tc.FlagKaryawan` |
| IC-009315 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailSyncDataCoreWorkingHours:123](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L123) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tc.FlagKaryawan` : `BIT` | `#xtmpTrsCustomer.FlagKaryawan` : `CHAR(1)` | `tc.FlagKaryawan` → `CHAR(1)` | `#xtmpTrsCustomer.FlagKaryawan <- tc.FlagKaryawan` |
| IC-003203 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSAmendSecurityTransaction_TT:605](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L605) | CASE / P11 / COMPARISON | `tc.FlagKaryawan` : `BIT` | `1` : `int` | `tc.FlagKaryawan` → `INT` | `tc.FlagKaryawan = 1` |
| IC-004313 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSInsertSecurityTransaction_TT_Original:771](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L771) | CASE / P11 / COMPARISON | `tc.FlagKaryawan` : `BIT` | `1` : `int` | `tc.FlagKaryawan` → `INT` | `tc.FlagKaryawan = 1` |
| IC-007864 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahObligasi:315](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahObligasi.sql#L315) | CASE / P11 / COMPARISON | `FlagKaryawan` : `BIT` | `1` : `int` | `FlagKaryawan` → `INT` | `FlagKaryawan = 1` |
| IC-013555 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListTreasuryCustomer_TM_NikNameRefKar:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListTreasuryCustomer_TM_NikNameRefKar.sql#L13) | CASE / P11 / COMPARISON | `tctm.FlagKaryawan` : `BIT` | `1` : `int` | `tctm.FlagKaryawan` → `INT` | `tctm.FlagKaryawan = 1` |
| IC-000648 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.OMGetListProductSekunder:144](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L144) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagKaryawan` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(FlagKaryawan,0)` |
| IC-008742 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRetailCalculateFee:249](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L249) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagKaryawan` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(FlagKaryawan,0)` |
| IC-009310 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailSyncDataCore:847](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L847) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `FlagKaryawan` : `BIT` | `t.FlagKaryawan` : `CHAR(1)` | `t.FlagKaryawan` → `BIT` | `FlagKaryawan = t.FlagKaryawan` |
| IC-009475 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:715](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L715) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `FlagKaryawan` : `BIT` | `t.FlagKaryawan` : `CHAR(1)` | `t.FlagKaryawan` → `BIT` | `FlagKaryawan = t.FlagKaryawan` |
| IC-013734 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_StatusSecurityTransaction_TT:159](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_StatusSecurityTransaction_TT.sql#L159) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tc.FlagKaryawan` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(tc.FlagKaryawan,0)` |
| IC-013752 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_StatusSecurityTransaction_TT:313](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_StatusSecurityTransaction_TT.sql#L313) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tc.FlagKaryawan` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(tc.FlagKaryawan,0)` |

<a id="kolom-idreviewbit"></a>
## IDReviewBit

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L72). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000802 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.OMUpdateReviewedCustomerID:8](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMUpdateReviewedCustomerID.sql#L8) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[IDReviewBit]` : `BIT` | `1` : `int` | `1` → `BIT` | `[IDReviewBit] = 1` |
| IC-013166 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:669](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L669) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomer_tm.IDReviewBit` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.treasurycustomer_tm.IDReviewBit <- 0` |

<a id="kolom-identitasexpireddate"></a>
## IdentitasExpiredDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L61). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012413 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSyncMasterIdAndFunc:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterIdAndFunc.sql#L65) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tc.IdentitasExpiredDate` : `DATETIME` | `'19450817'` : `varchar` | `'19450817'` → `DATETIME` | `isnull(tc.IdentitasExpiredDate,'19450817')` |

<a id="kolom-insertedby"></a>
## InsertedBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L25). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007859 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRPTMasterNasabahObligasi:240](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahObligasi.sql#L240) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `InsertedBy` : `INT` | `#tmpRESULT.NIKInputer` : `VARCHAR(40)` | `InsertedBy` → `VARCHAR(40)` | `#tmpRESULT.NIKInputer <- InsertedBy` |
| IC-013167 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:669](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L669) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomer_tm.InsertedBy` : `INT` | `InsertedBy` : `varchar(20)` | `InsertedBy` → `INT` | `retail.dbo.treasurycustomer_tm.InsertedBy <- InsertedBy` |

<a id="kolom-issuratuserumah"></a>
## IsSuratUseRumah

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L48). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008141 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportFormBuySellObli:58](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportFormBuySellObli.sql#L58) | CASE / P11 / COMPARISON | `tctm.IsSuratUseRumah` : `BIT` | `0` : `int` | `tctm.IsSuratUseRumah` → `INT` | `tctm.IsSuratUseRumah = 0` |
| IC-008142 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportFormBuySellObli:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportFormBuySellObli.sql#L59) | CASE / P11 / COMPARISON | `tctm.IsSuratUseRumah` : `BIT` | `1` : `int` | `tctm.IsSuratUseRumah` → `INT` | `tctm.IsSuratUseRumah = 1` |
| IC-010323 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptObligasiKonfirmasiBond:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptObligasiKonfirmasiBond.sql#L34) | CASE / P11 / COMPARISON | `IsSuratUseRumah` : `BIT` | `1` : `int` | `IsSuratUseRumah` → `INT` | `IsSuratUseRumah = 1` |
| IC-010606 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPTOSNasabah:176](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPTOSNasabah.sql#L176) | CASE / P11 / COMPARISON | `tc.IsSuratUseRumah` : `BIT` | `1` : `int` | `tc.IsSuratUseRumah` → `INT` | `tc.IsSuratUseRumah = 1` |

<a id="kolom-jenisidentitas"></a>
## JenisIdentitas

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 20 lokasi / 8 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007802 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSREPOReportBiayaMaterai:101](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L101) | CASE / P01 / COMPARISON | `JenisIdentitas` : `CHAR(1)` | `1` : `int` | `JenisIdentitas` → `INT` | `JenisIdentitas = 1` |
| IC-007804 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSREPOReportBiayaMaterai:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L111) | CASE / P01 / COMPARISON | `JenisIdentitas` : `CHAR(1)` | `1` : `int` | `JenisIdentitas` → `INT` | `JenisIdentitas = 1` |
| IC-007821 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSREPOReportBiayaMaterai:188](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L188) | CASE / P01 / COMPARISON | `JenisIdentitas` : `CHAR(1)` | `1` : `int` | `JenisIdentitas` → `INT` | `JenisIdentitas = 1` |
| IC-007823 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSREPOReportBiayaMaterai:198](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L198) | CASE / P01 / COMPARISON | `JenisIdentitas` : `CHAR(1)` | `1` : `int` | `JenisIdentitas` → `INT` | `JenisIdentitas = 1` |
| IC-007836 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSREPOReportBiayaMaterai:259](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L259) | CASE / P01 / COMPARISON | `JenisIdentitas` : `CHAR(1)` | `1` : `int` | `JenisIdentitas` → `INT` | `JenisIdentitas = 1` |
| IC-007838 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSREPOReportBiayaMaterai:269](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L269) | CASE / P01 / COMPARISON | `JenisIdentitas` : `CHAR(1)` | `1` : `int` | `JenisIdentitas` → `INT` | `JenisIdentitas = 1` |
| IC-007865 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahObligasi:340](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahObligasi.sql#L340) | CASE / P01 / COMPARISON | `JenisIdentitas` : `CHAR(1)` | `1` : `int` | `JenisIdentitas` → `INT` | `JenisIdentitas = 1` |
| IC-007866 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahObligasi:341](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahObligasi.sql#L341) | CASE / P01 / COMPARISON | `JenisIdentitas` : `CHAR(1)` | `2` : `int` | `JenisIdentitas` → `INT` | `JenisIdentitas = 2` |
| IC-007867 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahObligasi:342](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahObligasi.sql#L342) | CASE / P01 / COMPARISON | `JenisIdentitas` : `CHAR(1)` | `3` : `int` | `JenisIdentitas` → `INT` | `JenisIdentitas = 3` |
| IC-008377 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSReportSKFBankBeli:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeli.sql#L59) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `tc.JenisIdentitas` : `CHAR(1)` | `1` : `int` | `tc.JenisIdentitas` → `INT` | `WHEN 1 THEN tc.NoIdentitas` |
| IC-013967 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.trs_getExpiredKTPNasabah:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_getExpiredKTPNasabah.sql#L33) | CASE / P01 / COMPARISON | `JenisIdentitas` : `CHAR(1)` | `1` : `int` | `JenisIdentitas` → `INT` | `JenisIdentitas = 1` |
| IC-013968 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.trs_getExpiredKTPNasabah:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_getExpiredKTPNasabah.sql#L34) | CASE / P01 / COMPARISON | `JenisIdentitas` : `CHAR(1)` | `2` : `int` | `JenisIdentitas` → `INT` | `JenisIdentitas = 2` |
| IC-008471 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTemplateSID:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSID.sql#L69) | CASE / P12 / COMPARISON | `tc.JenisIdentitas` : `CHAR(1)` | `'1'` : `varchar` | `tc.JenisIdentitas` → `VARCHAR` | `tc.JenisIdentitas = '1'` |
| IC-008473 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTemplateSID:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSID.sql#L77) | CASE / P12 / COMPARISON | `tc.JenisIdentitas` : `CHAR(1)` | `'3'` : `varchar` | `tc.JenisIdentitas` → `VARCHAR` | `tc.JenisIdentitas = '3'` |
| IC-008483 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTemplateSIDBySecAccNo:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSIDBySecAccNo.sql#L29) | CASE / P12 / COMPARISON | `tc.JenisIdentitas` : `CHAR(1)` | `'1'` : `varchar` | `tc.JenisIdentitas` → `VARCHAR` | `tc.JenisIdentitas = '1'` |
| IC-008485 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTemplateSIDBySecAccNo:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSIDBySecAccNo.sql#L37) | CASE / P12 / COMPARISON | `tc.JenisIdentitas` : `CHAR(1)` | `'3'` : `varchar` | `tc.JenisIdentitas` → `VARCHAR` | `tc.JenisIdentitas = '3'` |
| IC-013465 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListStatusTreasuryCustomer_TM:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L41) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `tc.JenisIdentitas` : `CHAR(1)` | `'1'` : `varchar` | `tc.JenisIdentitas` → `VARCHAR` | `when '1' then 'KTP'` |
| IC-013476 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListStatusTreasuryCustomer_TM:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L98) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `tc.JenisIdentitas` : `CHAR(1)` | `'1'` : `varchar` | `tc.JenisIdentitas` → `VARCHAR` | `when '1' then 'KTP'` |
| IC-013488 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListStatusTreasuryCustomer_TM:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L156) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `tc.JenisIdentitas` : `CHAR(1)` | `'1'` : `varchar` | `tc.JenisIdentitas` → `VARCHAR` | `when '1' then 'KTP'` |
| IC-012411 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSyncMasterIdAndFunc:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterIdAndFunc.sql#L64) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tc.JenisIdentitas` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(tc.JenisIdentitas,'')` |

<a id="kolom-jeniskelamin"></a>
## JenisKelamin

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000989 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.ORISBNIPopulateRegisterInvestor:95](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateRegisterInvestor.sql#L95) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `JenisKelamin` : `CHAR(1)` | `#tmpDataInvestor.KodeJenisKelamin` : `varchar(10)` | `JenisKelamin` → `varchar(10)` | `#tmpDataInvestor.KodeJenisKelamin <- JenisKelamin` |
| IC-007868 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahObligasi:347](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahObligasi.sql#L347) | CASE / P01 / COMPARISON | `JenisKelamin` : `CHAR(1)` | `1` : `int` | `JenisKelamin` → `INT` | `JenisKelamin = 1` |
| IC-007869 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahObligasi:348](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahObligasi.sql#L348) | CASE / P01 / COMPARISON | `JenisKelamin` : `CHAR(1)` | `2` : `int` | `JenisKelamin` → `INT` | `JenisKelamin = 2` |
| IC-013466 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListStatusTreasuryCustomer_TM:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L47) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `tc.JenisKelamin` : `CHAR(1)` | `'1'` : `varchar` | `tc.JenisKelamin` → `VARCHAR` | `when '1' then 'Laki-Laki'` |
| IC-013467 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListStatusTreasuryCustomer_TM:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L48) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `tc.JenisKelamin` : `CHAR(1)` | `'2'` : `varchar` | `tc.JenisKelamin` → `VARCHAR` | `when '2' then 'Wanita'` |
| IC-013477 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListStatusTreasuryCustomer_TM:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L104) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `tc.JenisKelamin` : `CHAR(1)` | `'1'` : `varchar` | `tc.JenisKelamin` → `VARCHAR` | `when '1' then 'Laki-Laki'` |
| IC-013478 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListStatusTreasuryCustomer_TM:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L105) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `tc.JenisKelamin` : `CHAR(1)` | `'2'` : `varchar` | `tc.JenisKelamin` → `VARCHAR` | `when '2' then 'Wanita'` |
| IC-013489 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListStatusTreasuryCustomer_TM:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L162) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `tc.JenisKelamin` : `CHAR(1)` | `'1'` : `varchar` | `tc.JenisKelamin` → `VARCHAR` | `when '1' then 'Laki-Laki'` |
| IC-013490 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListStatusTreasuryCustomer_TM:163](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L163) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `tc.JenisKelamin` : `CHAR(1)` | `'2'` : `varchar` | `tc.JenisKelamin` → `VARCHAR` | `when '2' then 'Wanita'` |

<a id="kolom-jenispekerjaan"></a>
## JenisPekerjaan

**Tipe dasar:** `CHAR(2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000990 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.ORISBNIPopulateRegisterInvestor:95](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateRegisterInvestor.sql#L95) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `JenisPekerjaan` : `CHAR(2)` | `#tmpDataInvestor.KodePekerjaan` : `varchar(10)` | `JenisPekerjaan` → `varchar(10)` | `#tmpDataInvestor.KodePekerjaan <- JenisPekerjaan` |
| IC-007862 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRPTMasterNasabahObligasi:240](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahObligasi.sql#L240) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `JenisPekerjaan` : `CHAR(2)` | `#tmpRESULT.Pekerjaan` : `VARCHAR(60)` | `JenisPekerjaan` → `VARCHAR(60)` | `#tmpRESULT.Pekerjaan <- JenisPekerjaan` |

<a id="kolom-kitasexpdate"></a>
## KITASExpDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L74). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 7 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004426 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSJobSyncKITAS:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSJobSyncKITAS.sql#L22) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tc.[KITASExpDate]` : `DATETIME` | `cf.[CFIEX6]` : `CHAR(8)` | `cf.[CFIEX6]` → `DATETIME` | `tc.[KITASExpDate] = cf.[CFIEX6]` |
| IC-004432 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSJobSyncKITAS:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSJobSyncKITAS.sql#L37) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tc.[KITASExpDate]` : `DATETIME` | `cf.[CFIEX6]` : `CHAR(8)` | `cf.[CFIEX6]` → `DATETIME` | `tc.[KITASExpDate] = cf.[CFIEX6]` |
| IC-010870 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORIDetail:585](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L585) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.KITASExpDate` : `DATETIME` | `'1900-01-01'` : `varchar` | `'1900-01-01'` → `DATETIME` | `isnull(tctm.KITASExpDate,'1900-01-01')` |
| IC-010987 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnesty:422](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L422) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.KITASExpDate` : `DATETIME` | `'1900-01-01'` : `varchar` | `'1900-01-01'` → `DATETIME` | `isnull(tctm.KITASExpDate,'1900-01-01')` |
| IC-011000 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnesty:509](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L509) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.KITASExpDate` : `DATETIME` | `'1900-01-01'` : `varchar` | `'1900-01-01'` → `DATETIME` | `isnull(tctm.KITASExpDate,'1900-01-01')` |
| IC-011175 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:384](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L384) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.KITASExpDate` : `DATETIME` | `'1900-01-01'` : `varchar` | `'1900-01-01'` → `DATETIME` | `ISNULL(tctm.KITASExpDate,'1900-01-01')` |
| IC-011283 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnestySPT:236](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestySPT.sql#L236) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.KITASExpDate` : `DATETIME` | `'1900-01-01'` : `varchar` | `'1900-01-01'` → `DATETIME` | `isnull(tctm.KITASExpDate,'1900-01-01')` |
| IC-011380 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnesty_Backdated:325](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty_Backdated.sql#L325) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.KITASExpDate` : `DATETIME` | `'1900-01-01'` : `varchar` | `'1900-01-01'` → `DATETIME` | `isnull(tctm.KITASExpDate,'1900-01-01')` |
| IC-013168 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:669](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L669) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomer_tm.KITASExpDate` : `DATETIME` | `[KITASExpDate]` : `varchar(10)` | `[KITASExpDate]` → `DATETIME` | `retail.dbo.treasurycustomer_tm.KITASExpDate <- [KITASExpDate]` |

<a id="kolom-kitaslastupdatedate"></a>
## KITASLastUpdateDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L75). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-010871 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORIDetail:586](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORIDetail.sql#L586) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.KITASLastUpdateDate` : `DATETIME` | `'1900-01-01'` : `varchar` | `'1900-01-01'` → `DATETIME` | `isnull(tctm.KITASLastUpdateDate,'1900-01-01')` |
| IC-010988 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnesty:424](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L424) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.KITASLastUpdateDate` : `DATETIME` | `'1900-01-01'` : `varchar` | `'1900-01-01'` → `DATETIME` | `isnull(tctm.KITASLastUpdateDate,'1900-01-01')` |
| IC-011001 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnesty:510](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql#L510) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.KITASLastUpdateDate` : `DATETIME` | `'1900-01-01'` : `varchar` | `'1900-01-01'` → `DATETIME` | `isnull(tctm.KITASLastUpdateDate,'1900-01-01')` |
| IC-011176 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnestyPTBC:385](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestyPTBC.sql#L385) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.KITASLastUpdateDate` : `DATETIME` | `'1900-01-01'` : `varchar` | `'1900-01-01'` → `DATETIME` | `ISNULL(tctm.KITASLastUpdateDate,'1900-01-01')` |
| IC-011284 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnestySPT:237](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnestySPT.sql#L237) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.KITASLastUpdateDate` : `DATETIME` | `'1900-01-01'` : `varchar` | `'1900-01-01'` → `DATETIME` | `isnull(tctm.KITASLastUpdateDate,'1900-01-01')` |
| IC-011381 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPayCouponORITaxAmnesty_Backdated:326](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty_Backdated.sql#L326) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.KITASLastUpdateDate` : `DATETIME` | `'1900-01-01'` : `varchar` | `'1900-01-01'` → `DATETIME` | `isnull(tctm.KITASLastUpdateDate,'1900-01-01')` |
| IC-013169 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:669](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L669) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomer_tm.KITASLastUpdateDate` : `DATETIME` | `[KITASLastUpdateDate]` : `varchar(10)` | `[KITASLastUpdateDate]` → `DATETIME` | `retail.dbo.treasurycustomer_tm.KITASLastUpdateDate <- [KITASLastUpdateDate]` |

<a id="kolom-kitasno"></a>
## KITASNo

**Tipe dasar:** `VARCHAR(40)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L73). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004425 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSJobSyncKITAS:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSJobSyncKITAS.sql#L21) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tc.[KITASNo]` : `VARCHAR(40)` | `cf.[CFSSNO]` : `CHAR(40)` | `cf.[CFSSNO]` → `VARCHAR(40)` | `tc.[KITASNo] = cf.[CFSSNO]` |
| IC-004431 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSJobSyncKITAS:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSJobSyncKITAS.sql#L36) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tc.[KITASNo]` : `VARCHAR(40)` | `cf.[CFSSNO]` : `CHAR(40)` | `cf.[CFSSNO]` → `VARCHAR(40)` | `tc.[KITASNo] = cf.[CFSSNO]` |

<a id="kolom-kodekota"></a>
## KodeKota

**Tipe dasar:** `VARCHAR(4)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008492 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportTrxCorporateBOND:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L141) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tc.KodeKota` : `VARCHAR(4)` | `#tmpResult.KodeKota` : `char(4)` | `tc.KodeKota` → `char(4)` | `#tmpResult.KodeKota <- tc.KodeKota` |

<a id="kolom-kodepropinsi"></a>
## KodePropinsi

**Tipe dasar:** `VARCHAR(4)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008493 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportTrxCorporateBOND:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L141) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tc.KodePropinsi` : `VARCHAR(4)` | `#tmpResult.KodePropinsi` : `char(4)` | `tc.KodePropinsi` → `char(4)` | `#tmpResult.KodePropinsi <- tc.KodePropinsi` |

<a id="kolom-lastupdateby"></a>
## LastUpdateBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L27). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007860 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRPTMasterNasabahObligasi:240](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahObligasi.sql#L240) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `LastUpdateBy` : `INT` | `#tmpRESULT.NIKOtorisasi` : `VARCHAR(40)` | `LastUpdateBy` → `VARCHAR(40)` | `#tmpRESULT.NIKOtorisasi <- LastUpdateBy` |
| IC-013170 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:669](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L669) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomer_tm.LastUpdateBy` : `INT` | `LastUpdateBy` : `varchar(20)` | `LastUpdateBy` → `INT` | `retail.dbo.treasurycustomer_tm.LastUpdateBy <- LastUpdateBy` |

<a id="kolom-mxlabel"></a>
## MxLabel

**Tipe dasar:** `CHAR(15)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L67). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, DESTINATION_COLUMN, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005559 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateMasterNasabahForMurex:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMasterNasabahForMurex.sql#L26) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `MxLabel` : `CHAR(15)` | `@cMxLabel` : `VARCHAR(16)` | `MxLabel` → `VARCHAR(16)` | `@cMxLabel = MxLabel` |
| IC-008554 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSResponseBONDCPYFromMX:753](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L753) | EXPRESSION / P07 / CASE_BRANCH | `tctm.MxLabel` : `CHAR(15)` | `result type` : `VARCHAR (combined result; facets unknown)` | `tctm.MxLabel` → `VARCHAR` | `CASE WHEN isnull(tctm.MxLabel,'') = '' THEN @pcMXLABEL ELSE tctm.MxLabel END` |
| IC-008555 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSResponseBONDCPYFromMX:754](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L754) | CASE / P12 / COMPARISON | `isnull(tctm.MxLabel,'')` : `CHAR(15)` | `''` : `varchar` | `isnull(tctm.MxLabel,'')` → `VARCHAR` | `isnull(tctm.MxLabel,'') = ''` |
| IC-008553 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBONDCPYFromMX:752](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L752) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `MxLabel` : `CHAR(15)` | `CASE WHEN isnull(tctm.MxLabel,'') = '' THEN @pcMXLABEL ELSE tctm.MxLabel END` : `VARCHAR (combined result; facets unknown)` | `CASE WHEN isnull(tctm.MxLabel,'') = '' THEN @pcMXLABEL ELSE tctm.MxLabel END` → `CHAR(15)` | `MxLabel = CASE WHEN isnull(tctm.MxLabel,'') = '' THEN @pcMXLABEL ELSE tctm.MxLabel END` |
| IC-008556 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSResponseBONDCPYFromMX:754](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L754) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.MxLabel` : `CHAR(15)` | `''` : `varchar` | `''` → `CHAR(15)` | `isnull(tctm.MxLabel,'')` |

<a id="kolom-nikuploadorichid"></a>
## NIKUploadOrichid

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L79). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000255 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.CustodySaveUploadOrchid:50](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodySaveUploadOrchid.sql#L50) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `NIKUploadOrichid` : `BIGINT` | `@pnNik` : `int` | `@pnNik` → `BIGINT` | `NIKUploadOrichid = @pnNik` |

<a id="kolom-nik-cs"></a>
## NIK_CS

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013468 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_ListStatusTreasuryCustomer_TM:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L56) | CASE / P04 / COMPARISON | `tc.NIK_CS` : `INT` | `'77777'` : `varchar` | `'77777'` → `INT` | `tc.NIK_CS = '77777'` |
| IC-013473 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_ListStatusTreasuryCustomer_TM:79](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L79) | JOIN_ON / P01 / COMPARISON | `tc.NIK_CS` : `INT` | `un.EMPLID` : `VARCHAR(11)` | `un.EMPLID` → `INT` | `tc.NIK_CS = un.EMPLID` |
| IC-013479 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_ListStatusTreasuryCustomer_TM:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L113) | CASE / P04 / COMPARISON | `tc.NIK_CS` : `INT` | `'77777'` : `varchar` | `'77777'` → `INT` | `tc.NIK_CS = '77777'` |
| IC-013484 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_ListStatusTreasuryCustomer_TM:139](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L139) | JOIN_ON / P01 / COMPARISON | `tc.NIK_CS` : `INT` | `un.EMPLID` : `VARCHAR(11)` | `un.EMPLID` → `INT` | `tc.NIK_CS = un.EMPLID` |
| IC-013491 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_ListStatusTreasuryCustomer_TM:171](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L171) | CASE / P04 / COMPARISON | `tc.NIK_CS` : `INT` | `'77777'` : `varchar` | `'77777'` → `INT` | `tc.NIK_CS = '77777'` |
| IC-013495 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_ListStatusTreasuryCustomer_TM:194](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L194) | JOIN_ON / P01 / COMPARISON | `tc.NIK_CS` : `INT` | `un.EMPLID` : `VARCHAR(11)` | `un.EMPLID` → `INT` | `tc.NIK_CS = un.EMPLID` |

<a id="kolom-nikreferentor"></a>
## NikReferentor

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L57). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007861 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRPTMasterNasabahObligasi:240](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahObligasi.sql#L240) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `NikReferentor` : `INT` | `#tmpRESULT.NIKReferantor` : `VARCHAR(50)` | `NikReferentor` → `VARCHAR(50)` | `#tmpRESULT.NIKReferantor <- NikReferentor` |

<a id="kolom-regviaonemobile"></a>
## RegViaONEMobile

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L70). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003345 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApproveSID:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveSID.sql#L81) | WHERE / P12 / COMPARISON | `tm.RegViaONEMobile` : `CHAR(1)` | `'Y'` : `varchar` | `tm.RegViaONEMobile` → `VARCHAR` | `tm.RegViaONEMobile = 'Y'` |
| IC-007863 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahObligasi:303](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahObligasi.sql#L303) | CASE / P12 / COMPARISON | `RegViaONEMobile` : `CHAR(1)` | `'Y'` : `varchar` | `RegViaONEMobile` → `VARCHAR` | `RegViaONEMobile = 'Y'` |
| IC-012492 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdateDataMasterNasabah:244](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateDataMasterNasabah.sql#L244) | WHERE / P12 / COMPARISON | `RegViaONEMobile` : `CHAR(1)` | `'Y'` : `varchar` | `RegViaONEMobile` → `VARCHAR` | `RegViaONEMobile = 'Y'` |

<a id="kolom-riskprofilecode"></a>
## RiskProfileCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L43). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000991 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.ORISBNIPopulateRegisterInvestor:95](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateRegisterInvestor.sql#L95) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `RiskProfileCode` : `CHAR(1)` | `#tmpDataInvestor.RiskProfileCode` : `varchar(2)` | `RiskProfileCode` → `varchar(2)` | `#tmpDataInvestor.RiskProfileCode <- RiskProfileCode` |

<a id="kolom-sid"></a>
## SID

**Tipe dasar:** `VARCHAR(100)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L41). **Existing leading index:** TreasuryCustomer_TM_NN5.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008149 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportMatchingSID:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportMatchingSID.sql#L30) | WHERE / P12 / COMPARISON | `isnull(mst.SID,'')` : `VARCHAR(100)` | `cf.CFSSNO` : `CHAR(40)` | `cf.CFSSNO` → `VARCHAR` | `cf.CFSSNO != isnull(mst.SID,'')` |

<a id="kolom-signature"></a>
## Signature

**Tipe dasar:** `VARBINARY(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L44). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013171 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:669](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L669) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomer_tm.Signature` : `VARBINARY(MAX)` | `[Signature]` : `image` | `[Signature]` → `VARBINARY(MAX)` | `retail.dbo.treasurycustomer_tm.Signature <- [Signature]` |

<a id="kolom-summary"></a>
## Summary

**Tipe dasar:** `VARBINARY(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L46). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013172 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:669](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L669) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomer_tm.Summary` : `VARBINARY(MAX)` | `Summary` : `image` | `Summary` → `VARBINARY(MAX)` | `retail.dbo.treasurycustomer_tm.Summary <- Summary` |

<a id="kolom-syaratketentuanbit"></a>
## SyaratKetentuanBit

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L38). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007871 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTMasterNasabahObligasi:374](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahObligasi.sql#L374) | CASE / P11 / COMPARISON | `SyaratKetentuanBit` : `BIT` | `1` : `int` | `SyaratKetentuanBit` → `INT` | `SyaratKetentuanBit = 1` |
| IC-012717 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateSecAccNo:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateSecAccNo.sql#L16) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `SyaratKetentuanBit` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(SyaratKetentuanBit, 0)` |
| IC-012726 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateSecAccNoOrderSuratBerharga:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateSecAccNoOrderSuratBerharga.sql#L25) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.SyaratKetentuanBit` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(tctm.SyaratKetentuanBit, 0)` |
| IC-012727 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateSecAccNoRiskProfile:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateSecAccNoRiskProfile.sql#L16) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `SyaratKetentuanBit` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(SyaratKetentuanBit, 0)` |
| IC-013557 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListTreasuryCustomer_TM_Original:155](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListTreasuryCustomer_TM_Original.sql#L155) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.SyaratKetentuanBit` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(tctm.SyaratKetentuanBit, 0)` |
| IC-013562 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListTreasuryCustomer_TM_Original:324](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListTreasuryCustomer_TM_Original.sql#L324) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.SyaratKetentuanBit` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(tctm.SyaratKetentuanBit, 0)` |

<a id="kolom-tanggallahir"></a>
## TanggalLahir

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005275 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateApprovalSRESID:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalSRESID.sql#L110) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `LTRIM(RTRIM(tc.TanggalLahir))` : `VARCHAR (expression; length unknown)` | `TanggalLahir` : `DATETIME` | `LTRIM(RTRIM(tc.TanggalLahir))` → `DATETIME` | `TanggalLahir = LTRIM(RTRIM(tc.TanggalLahir))` |

<a id="kolom-isonlineacc"></a>
## isOnlineAcc

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TM.sql#L45). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN_EXPRESSION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013470 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_ListStatusTreasuryCustomer_TM:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L63) | CASE / P11 / COMPARISON | `isnull(tc.isOnlineAcc,0)` : `BIT` | `1` : `int` | `isnull(tc.isOnlineAcc,0)` → `INT` | `isnull(tc.isOnlineAcc,0) = 1` |
| IC-013481 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_ListStatusTreasuryCustomer_TM:122](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L122) | CASE / P11 / COMPARISON | `isnull(tc.isOnlineAcc,0)` : `BIT` | `1` : `int` | `isnull(tc.isOnlineAcc,0)` → `INT` | `isnull(tc.isOnlineAcc,0) = 1` |
| IC-013492 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.trs_ListStatusTreasuryCustomer_TM:177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L177) | CASE / P11 / COMPARISON | `isnull(tc.isOnlineAcc,0)` : `BIT` | `1` : `int` | `isnull(tc.isOnlineAcc,0)` → `INT` | `isnull(tc.isOnlineAcc,0) = 1` |
| IC-005397 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateEditCustomerAuthorization:441](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEditCustomerAuthorization.sql#L441) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tctm.isOnlineAcc` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(tctm.isOnlineAcc,0)` |
| IC-013471 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusTreasuryCustomer_TM:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L63) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tc.isOnlineAcc` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(tc.isOnlineAcc,0)` |
| IC-013482 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusTreasuryCustomer_TM:122](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L122) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tc.isOnlineAcc` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(tc.isOnlineAcc,0)` |
| IC-013493 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListStatusTreasuryCustomer_TM:177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql#L177) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tc.isOnlineAcc` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(tc.isOnlineAcc,0)` |
