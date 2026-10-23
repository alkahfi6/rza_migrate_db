# Retail UAT.dbo.TRSBatchTransaction_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 13 lokasi source, 5 procedure, 6 kolom. **A:** 2 lokasi pada 1 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-account-id-credit"></a>
## account_id_credit

**Tipe dasar:** `VARCHAR(19)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBatchTransaction_TM.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009556 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptBatchTransaction:35](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptBatchTransaction.sql#L35) | JOIN_ON / P01 / COMPARISON | `tm.account_id_credit` : `VARCHAR(19)` | `dvc.ACCTNO` : `DECIMAL(19,0)` | `tm.account_id_credit` → `DECIMAL` | `tm.account_id_credit = dvc.ACCTNO` |

<a id="kolom-account-id-debit"></a>
## account_id_debit

**Tipe dasar:** `VARCHAR(19)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBatchTransaction_TM.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009555 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptBatchTransaction:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptBatchTransaction.sql#L33) | JOIN_ON / P01 / COMPARISON | `tm.account_id_debit` : `VARCHAR(19)` | `dvd.ACCTNO` : `DECIMAL(19,0)` | `tm.account_id_debit` → `DECIMAL` | `tm.account_id_debit = dvd.ACCTNO` |

<a id="kolom-credit-currency"></a>
## credit_currency

**Tipe dasar:** `CHAR(3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBatchTransaction_TM.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009576 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptCorporateAction:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCorporateAction.sql#L130) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `credit_currency` : `CHAR(3)` | `#tmpPooling.Currency` : `varchar(3)` | `credit_currency` → `varchar(3)` | `#tmpPooling.Currency <- credit_currency` |

<a id="kolom-guid"></a>
## guid

**Tipe dasar:** `UNIQUEIDENTIFIER`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBatchTransaction_TM.sql#L2). **Existing leading index:** TRSBatchTransaction_TM_IDX.

**Frekuensi source:** 4 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003367 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBatchTransactionInsertNew:40](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBatchTransactionInsertNew.sql#L40) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchtransaction_tm.guid` : `UNIQUEIDENTIFIER` | `@pvGuid` : `varchar(40)` | `@pvGuid` → `UNIQUEIDENTIFIER` | `retail.dbo.trsbatchtransaction_tm.guid <- @pvGuid` |
| IC-003371 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSBatchTransactionStatusByFileName:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBatchTransactionStatusByFileName.sql#L16) | WHERE / P13 / IN_SUBQUERY | `[guid]` : `UNIQUEIDENTIFIER` | `(select [Value] from [dbo].[SplitString](@pcGuid, '&#124;'))` : `NVARCHAR(MAX)` | `(select [Value] from [dbo].[SplitString](@pcGuid, '&#124;'))` → `UNIQUEIDENTIFIER` | `[guid] IN (select [Value] from [dbo].[SplitString](@pcGuid, '&#124;'))` |
| IC-003373 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSBatchTransactionStatusNew:20](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBatchTransactionStatusNew.sql#L20) | WHERE / P13 / COMPARISON | `guid` : `UNIQUEIDENTIFIER` | `@pvGuid` : `varchar(40)` | `@pvGuid` → `UNIQUEIDENTIFIER` | `guid = @pvGuid` |
| IC-003375 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSBatchTransactionStatusNew:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBatchTransactionStatusNew.sql#L31) | WHERE / P13 / COMPARISON | `guid` : `UNIQUEIDENTIFIER` | `@pvGuid` : `varchar(40)` | `@pvGuid` → `UNIQUEIDENTIFIER` | `guid = @pvGuid` |

<a id="kolom-sibs-user-suid"></a>
## sibs_user_suid

**Tipe dasar:** `CHAR(7)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBatchTransaction_TM.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003368 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBatchTransactionInsertNew:40](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBatchTransactionInsertNew.sql#L40) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchtransaction_tm.sibs_user_suid` : `CHAR(7)` | `@pcTellerId` : `varchar(7)` | `@pcTellerId` → `CHAR(7)` | `retail.dbo.trsbatchtransaction_tm.sibs_user_suid <- @pcTellerId` |

<a id="kolom-transaction-status"></a>
## transaction_status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBatchTransaction_TM.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009579 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptCorporateAction:146](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCorporateAction.sql#L146) | WHERE / P12 / COMPARISON | `[transaction_status]` : `CHAR(1)` | `'B'` : `varchar` | `[transaction_status]` → `VARCHAR` | `[transaction_status] = 'B'` |
| IC-003369 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBatchTransactionInsertNew:40](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBatchTransactionInsertNew.sql#L40) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbatchtransaction_tm.transaction_status` : `CHAR(1)` | `@pcStatus` : `varchar(3)` | `@pcStatus` → `CHAR(1)` | `retail.dbo.trsbatchtransaction_tm.transaction_status <- @pcStatus` |
| IC-003370 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBatchTransactionStatusByFileName:11](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBatchTransactionStatusByFileName.sql#L11) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `transaction_status` : `CHAR(1)` | `@pcStatus` : `varchar(10)` | `@pcStatus` → `CHAR(1)` | `transaction_status = @pcStatus` |
| IC-003372 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBatchTransactionStatusNew:14](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBatchTransactionStatusNew.sql#L14) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `transaction_status` : `CHAR(1)` | `@pcStatus` : `varchar(1)` | `@pcStatus` → `CHAR(1)` | `transaction_status = @pcStatus` |
| IC-003374 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBatchTransactionStatusNew:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBatchTransactionStatusNew.sql#L25) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `transaction_status` : `CHAR(1)` | `@pcStatus` : `varchar(1)` | `@pcStatus` → `CHAR(1)` | `transaction_status = @pcStatus` |
