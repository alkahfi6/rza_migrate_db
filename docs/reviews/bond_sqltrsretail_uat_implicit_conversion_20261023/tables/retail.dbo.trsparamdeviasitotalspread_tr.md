# Retail UAT.dbo.TRSParamDeviasiTotalSpread_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 11 lokasi source, 2 procedure, 5 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 4 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-apprid"></a>
## ApprId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamDeviasiTotalSpread_TR.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006564 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessApprParamDeviasiTotalSpread:92](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamDeviasiTotalSpread.sql#L92) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ApprId` : `BIGINT` | `#tempTableReseter.ApprId` : `int` | `ApprId` → `int` | `#tempTableReseter.ApprId <- ApprId` |
| IC-006567 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessApprParamDeviasiTotalSpread:126](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamDeviasiTotalSpread.sql#L126) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamdeviasitotalspread_tr.ApprId` : `BIGINT` | `ApprId` : `int` | `ApprId` → `BIGINT` | `retail.dbo.trsparamdeviasitotalspread_tr.ApprId <- ApprId` |

<a id="kolom-inserteddate"></a>
## InsertedDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamDeviasiTotalSpread_TR.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006565 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessApprParamDeviasiTotalSpread:92](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamDeviasiTotalSpread.sql#L92) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `InsertedDate` : `DATETIME` | `#tempTableReseter.InsertedDate` : `varchar(50)` | `InsertedDate` → `varchar(50)` | `#tempTableReseter.InsertedDate <- InsertedDate` |
| IC-006568 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessApprParamDeviasiTotalSpread:126](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamDeviasiTotalSpread.sql#L126) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamdeviasitotalspread_tr.InsertedDate` : `DATETIME` | `InsertedDate` : `varchar(50)` | `InsertedDate` → `DATETIME` | `retail.dbo.trsparamdeviasitotalspread_tr.InsertedDate <- InsertedDate` |

<a id="kolom-lastupdatedate"></a>
## LastUpdateDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamDeviasiTotalSpread_TR.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006566 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessApprParamDeviasiTotalSpread:92](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamDeviasiTotalSpread.sql#L92) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `LastUpdateDate` : `DATETIME` | `#tempTableReseter.LastUpdateDate` : `varchar(50)` | `LastUpdateDate` → `varchar(50)` | `#tempTableReseter.LastUpdateDate <- LastUpdateDate` |
| IC-006569 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessApprParamDeviasiTotalSpread:126](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamDeviasiTotalSpread.sql#L126) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamdeviasitotalspread_tr.LastUpdateDate` : `DATETIME` | `LastUpdateDate` : `varchar(50)` | `LastUpdateDate` → `DATETIME` | `retail.dbo.trsparamdeviasitotalspread_tr.LastUpdateDate <- LastUpdateDate` |

<a id="kolom-maxspreadbanksell"></a>
## MaxSpreadBankSell

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamDeviasiTotalSpread_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008775 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailCalculateFee:512](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L512) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `MaxSpreadBankSell` : `DECIMAL(20,5)` | `@nMaxSpreadBankSell` : `money` | `MaxSpreadBankSell` → `money` | `@nMaxSpreadBankSell = MaxSpreadBankSell` |

<a id="kolom-sectype"></a>
## SecType

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamDeviasiTotalSpread_TR.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008776 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:517](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L517) | WHERE / P11 / COMPARISON | `SecType` : `INT` | `@nSecType` : `SMALLINT` | `@nSecType` → `INT` | `SecType = @nSecType` |
| IC-008806 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:622](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L622) | WHERE / P11 / COMPARISON | `SecType` : `INT` | `@nSecType` : `SMALLINT` | `@nSecType` → `INT` | `SecType = @nSecType` |
| IC-008811 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:640](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L640) | WHERE / P11 / COMPARISON | `pdts.SecType` : `INT` | `@nSecType` : `SMALLINT` | `@nSecType` → `INT` | `pdts.SecType = @nSecType` |
| IC-008812 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:645](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L645) | WHERE / P11 / COMPARISON | `pdts.SecType` : `INT` | `@nSecType` : `SMALLINT` | `@nSecType` → `INT` | `pdts.SecType = @nSecType` |
