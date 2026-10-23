# Retail UAT.dbo.TreasuryCustomerChange_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 15 lokasi source, 3 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-changeprofileresiko"></a>
## ChangeProfileResiko

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerChange_TH.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-010301 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptMonitoringDocAcc:142](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptMonitoringDocAcc.sql#L142) | CASE / P11 / COMPARISON | `cc.[ChangeProfileResiko]` : `BIT` | `1` : `int` | `cc.[ChangeProfileResiko]` → `INT` | `cc.[ChangeProfileResiko] = 1` |
| IC-008536 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBONDCPYFromMX:218](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L218) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerchange_th.ChangeProfileResiko` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.treasurycustomerchange_th.ChangeProfileResiko <- 0` |
| IC-013173 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:847](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L847) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerchange_th.ChangeProfileResiko` : `BIT` | `1` : `int` | `1` → `BIT` | `retail.dbo.treasurycustomerchange_th.ChangeProfileResiko <- 1` |

<a id="kolom-changerekrelasi"></a>
## ChangeRekRelasi

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerChange_TH.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008542 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSResponseBONDCPYFromMX:233](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L233) | EXPRESSION / P07 / CASE_BRANCH | `tc.[ChangeRekRelasi]` : `BIT` | `result type` : `INT (combined result; facets unknown)` | `tc.[ChangeRekRelasi]` → `INT` | `CASE WHEN tcm.[NoRekInvestor] <> tmp.[NoRekInvestor] THEN 1 ELSE tc.[ChangeRekRelasi] END` |
| IC-010303 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptMonitoringDocAcc:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptMonitoringDocAcc.sql#L147) | CASE / P11 / COMPARISON | `cc.[ChangeRekRelasi]` : `BIT` | `1` : `int` | `cc.[ChangeRekRelasi]` → `INT` | `cc.[ChangeRekRelasi] = 1` |
| IC-008537 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBONDCPYFromMX:218](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L218) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerchange_th.ChangeRekRelasi` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.treasurycustomerchange_th.ChangeRekRelasi <- 0` |
| IC-008541 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBONDCPYFromMX:232](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L232) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tc.[ChangeRekRelasi]` : `BIT` | `CASE WHEN tcm.[NoRekInvestor] <> tmp.[NoRekInvestor] THEN 1 ELSE tc.[ChangeRekRelasi] END` : `INT (combined result; facets unknown)` | `CASE WHEN tcm.[NoRekInvestor] <> tmp.[NoRekInvestor] THEN 1 ELSE tc.[ChangeRekRelasi] END` → `BIT` | `tc.[ChangeRekRelasi] = CASE WHEN tcm.[NoRekInvestor] <> tmp.[NoRekInvestor] THEN 1 ELSE tc.[ChangeRekRelasi] END` |
| IC-008543 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBONDCPYFromMX:245](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L245) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tc.[ChangeRekRelasi]` : `BIT` | `1` : `int` | `1` → `BIT` | `tc.[ChangeRekRelasi] = 1` |
| IC-008544 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBONDCPYFromMX:255](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L255) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tc.[ChangeRekRelasi]` : `BIT` | `1` : `int` | `1` → `BIT` | `tc.[ChangeRekRelasi] = 1` |
| IC-013174 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:847](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L847) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerchange_th.ChangeRekRelasi` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.treasurycustomerchange_th.ChangeRekRelasi <- 0` |

<a id="kolom-changesyaratketentuan"></a>
## ChangeSyaratKetentuan

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerChange_TH.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008540 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSResponseBONDCPYFromMX:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L228) | EXPRESSION / P07 / CASE_BRANCH | `tc.[ChangeSyaratKetentuan]` : `BIT` | `result type` : `INT (combined result; facets unknown)` | `tc.[ChangeSyaratKetentuan]` → `INT` | `CASE WHEN tcm.[SyaratKetentuanBit] <> tmp.[SyaratKetentuanBit] THEN 1 ELSE tc.[ChangeSyaratKetentuan] END` |
| IC-010299 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptMonitoringDocAcc:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptMonitoringDocAcc.sql#L137) | CASE / P11 / COMPARISON | `cc.[ChangeSyaratKetentuan]` : `BIT` | `1` : `int` | `cc.[ChangeSyaratKetentuan]` → `INT` | `cc.[ChangeSyaratKetentuan] = 1` |
| IC-008538 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBONDCPYFromMX:218](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L218) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerchange_th.ChangeSyaratKetentuan` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.treasurycustomerchange_th.ChangeSyaratKetentuan <- 0` |
| IC-008539 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBONDCPYFromMX:227](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L227) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tc.[ChangeSyaratKetentuan]` : `BIT` | `CASE WHEN tcm.[SyaratKetentuanBit] <> tmp.[SyaratKetentuanBit] THEN 1 ELSE tc.[ChangeSyaratKetentuan] END` : `INT (combined result; facets unknown)` | `CASE WHEN tcm.[SyaratKetentuanBit] <> tmp.[SyaratKetentuanBit] THEN 1 ELSE tc.[ChangeSyaratKetentuan] END` → `BIT` | `tc.[ChangeSyaratKetentuan] = CASE WHEN tcm.[SyaratKetentuanBit] <> tmp.[SyaratKetentuanBit] THEN 1 ELSE tc.[ChangeSyaratKetentuan] END` |
| IC-013175 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:847](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L847) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerchange_th.ChangeSyaratKetentuan` : `BIT` | `1` : `int` | `1` → `BIT` | `retail.dbo.treasurycustomerchange_th.ChangeSyaratKetentuan <- 1` |
