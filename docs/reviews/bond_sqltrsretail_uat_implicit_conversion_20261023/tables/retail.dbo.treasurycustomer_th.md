# Retail UAT.dbo.TreasuryCustomer_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 2 procedure, 4 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cifid"></a>
## CIFId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TH.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013151 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_InsertTreasuryCustomer_TM:531](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L531) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `max(CIFId)` : `BIGINT` | `@nCounter2` : `int` | `max(CIFId)` → `int` | `@nCounter2 = max(CIFId)` |

<a id="kolom-citizenship"></a>
## Citizenship

**Tipe dasar:** `CHAR(3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TH.sql#L43). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006835 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessEditCustomerAuthorization:422](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessEditCustomerAuthorization.sql#L422) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomer_th.Citizenship` : `CHAR(3)` | `Citizenship` : `VARCHAR(3)` | `Citizenship` → `CHAR(3)` | `retail.dbo.treasurycustomer_th.Citizenship <- Citizenship` |

<a id="kolom-signature"></a>
## Signature

**Tipe dasar:** `VARBINARY(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TH.sql#L45). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006836 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessEditCustomerAuthorization:422](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessEditCustomerAuthorization.sql#L422) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomer_th.Signature` : `VARBINARY(MAX)` | `[Signature]` : `image` | `[Signature]` → `VARBINARY(MAX)` | `retail.dbo.treasurycustomer_th.Signature <- [Signature]` |

<a id="kolom-summary"></a>
## Summary

**Tipe dasar:** `VARBINARY(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomer_TH.sql#L47). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006837 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessEditCustomerAuthorization:422](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessEditCustomerAuthorization.sql#L422) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomer_th.Summary` : `VARBINARY(MAX)` | `Summary` : `image` | `Summary` → `VARBINARY(MAX)` | `retail.dbo.treasurycustomer_th.Summary <- Summary` |
