# Retail UAT.dbo.TRSBankCustody_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 11 lokasi source, 4 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-iscommitted"></a>
## IsCommitted

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBankCustody_TH.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006737 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessBankCustody:520](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustody.sql#L520) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbankcustody_th.IsCommitted` : `BIT` | `CASE WHEN OperationCode IN ('U', 'D') THEN 0 ELSE 1 END` : `INT (combined result; facets unknown)` | `CASE WHEN OperationCode IN ('U', 'D') THEN 0 ELSE 1 END` → `BIT` | `retail.dbo.trsbankcustody_th.IsCommitted <- CASE WHEN OperationCode IN ('U', 'D') THEN 0 ELSE 1 END` |
| IC-006748 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessBankCustodyEOD:72](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustodyEOD.sql#L72) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbankcustody_th.IsCommitted` : `BIT` | `1` : `int` | `1` → `BIT` | `retail.dbo.trsbankcustody_th.IsCommitted <- 1` |

<a id="kolom-operationcode"></a>
## OperationCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBankCustody_TH.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009531 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParameterBankKustodi:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParameterBankKustodi.sql#L162) | JOIN_ON / P12 / COMPARISON | `tbcth.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `tbcth.OperationCode` → `VARCHAR` | `tbcth.OperationCode = 'U'` |

<a id="kolom-processcode"></a>
## ProcessCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBankCustody_TH.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005290 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateBankCustody:300](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateBankCustody.sql#L300) | WHERE / P12 / COMPARISON | `ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `ProcessCode` → `VARCHAR` | `ProcessCode = 'P'` |
| IC-009529 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParameterBankKustodi:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParameterBankKustodi.sql#L111) | WHERE / P12 / IN_VALUE_TYPE | `ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `ProcessCode` → `VARCHAR` | `ProcessCode in ('A','R')` |
| IC-009530 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParameterBankKustodi:161](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParameterBankKustodi.sql#L161) | JOIN_ON / P12 / COMPARISON | `tbcth.ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `tbcth.ProcessCode` → `VARCHAR` | `tbcth.ProcessCode = 'P'` |
| IC-009532 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParameterBankKustodi:171](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParameterBankKustodi.sql#L171) | JOIN_ON / P12 / COMPARISON | `tbcth.ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `tbcth.ProcessCode` → `VARCHAR` | `tbcth.ProcessCode = 'P'` |
| IC-006727 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessBankCustody:299](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustody.sql#L299) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbankcustody_th.ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trsbankcustody_th.ProcessCode <- 'P'` |
| IC-006729 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessBankCustody:363](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustody.sql#L363) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbankcustody_th.ProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trsbankcustody_th.ProcessCode <- 'P'` |
| IC-006738 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessBankCustody:520](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustody.sql#L520) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbankcustody_th.ProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.trsbankcustody_th.ProcessCode <- 'A'` |
| IC-006741 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessBankCustody:602](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustody.sql#L602) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbankcustody_th.ProcessCode` : `CHAR(1)` | `'R'` : `varchar` | `'R'` → `CHAR(1)` | `retail.dbo.trsbankcustody_th.ProcessCode <- 'R'` |
