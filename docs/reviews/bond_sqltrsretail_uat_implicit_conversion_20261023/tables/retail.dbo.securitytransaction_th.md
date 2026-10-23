# Retail UAT.dbo.SecurityTransaction_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 23 lokasi source, 8 procedure, 17 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-actionid"></a>
## ActionID

**Tipe dasar:** `VARCHAR(15)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TH.sql#L35). **Existing leading index:** SecurityTransaction_TH_NN3.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013662 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_RejectStatusSecurityTransaction_TT:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_RejectStatusSecurityTransaction_TT.sql#L66) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_th.ActionID` : `VARCHAR(15)` | `@nNIK` : `int` | `@nNIK` → `VARCHAR(15)` | `retail.dbo.securitytransaction_th.ActionID <- @nNIK` |

<a id="kolom-changeby"></a>
## ChangeBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TH.sql#L36). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008576 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBondFromMX:295](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L295) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_th.ChangeBy` : `INT` | `'77777'` : `varchar` | `'77777'` → `INT` | `retail.dbo.securitytransaction_th.ChangeBy <- '77777'` |
| IC-012661 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateTaxAmnestyExpired:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateTaxAmnestyExpired.sql#L47) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_th.ChangeBy` : `INT` | `'77777'` : `varchar` | `'77777'` → `INT` | `retail.dbo.securitytransaction_th.ChangeBy <- '77777'` |

<a id="kolom-changedate"></a>
## ChangeDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TH.sql#L37). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013779 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateStatusSecurityMaturity:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityMaturity.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_th.ChangeDate` : `DATETIME` | `@dcurrent_working_date` : `smalldatetime` | `@dcurrent_working_date` → `DATETIME` | `retail.dbo.securitytransaction_th.ChangeDate <- @dcurrent_working_date` |

<a id="kolom-dealprice"></a>
## DealPrice

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TH.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012963 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_DeleteSecurityTransaction_TT:124](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_DeleteSecurityTransaction_TT.sql#L124) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_th.DealPrice` : `DECIMAL(20,5)` | `td.DealPrice` : `money` | `td.DealPrice` → `DECIMAL(20,5)` | `retail.dbo.securitytransaction_th.DealPrice <- td.DealPrice` |

<a id="kolom-dealyield"></a>
## DealYield

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TH.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012964 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_DeleteSecurityTransaction_TT:124](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_DeleteSecurityTransaction_TT.sql#L124) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_th.DealYield` : `DECIMAL(20,5)` | `td.DealYield` : `numeric` | `td.DealYield` → `DECIMAL(20,5)` | `retail.dbo.securitytransaction_th.DealYield <- td.DealYield` |

<a id="kolom-hargamodal"></a>
## HargaModal

**Tipe dasar:** `DECIMAL(19,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TH.sql#L52). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003240 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSAmendSecurityTransaction_TT:1263](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L1263) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_th.HargaModal` : `DECIMAL(19,5)` | `t.HargaModal` : `MONEY` | `t.HargaModal` → `DECIMAL(19,5)` | `retail.dbo.securitytransaction_th.HargaModal <- t.HargaModal` |
| IC-003257 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSAmendSecurityTransaction_TT_Commit:237](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT_Commit.sql#L237) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_th.HargaModal` : `DECIMAL(19,5)` | `t.HargaModal` : `MONEY` | `t.HargaModal` → `DECIMAL(19,5)` | `retail.dbo.securitytransaction_th.HargaModal <- t.HargaModal` |

<a id="kolom-hargaori"></a>
## HargaORI

**Tipe dasar:** `DECIMAL(19,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TH.sql#L51). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003241 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSAmendSecurityTransaction_TT:1263](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql#L1263) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_th.HargaORI` : `DECIMAL(19,5)` | `t.HargaORI` : `MONEY` | `t.HargaORI` → `DECIMAL(19,5)` | `retail.dbo.securitytransaction_th.HargaORI <- t.HargaORI` |
| IC-003258 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSAmendSecurityTransaction_TT_Commit:237](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT_Commit.sql#L237) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_th.HargaORI` : `DECIMAL(19,5)` | `t.HargaORI` : `MONEY` | `t.HargaORI` → `DECIMAL(19,5)` | `retail.dbo.securitytransaction_th.HargaORI <- t.HargaORI` |

<a id="kolom-insertedby"></a>
## InsertedBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TH.sql#L28). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012965 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_DeleteSecurityTransaction_TT:124](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_DeleteSecurityTransaction_TT.sql#L124) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_th.InsertedBy` : `INT` | `td.InsertedBy` : `varchar(20)` | `td.InsertedBy` → `INT` | `retail.dbo.securitytransaction_th.InsertedBy <- td.InsertedBy` |
| IC-013663 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_RejectStatusSecurityTransaction_TT:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_RejectStatusSecurityTransaction_TT.sql#L66) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_th.InsertedBy` : `INT` | `Trx.InsertedDate` : `DATETIME` | `Trx.InsertedDate` → `INT` | `retail.dbo.securitytransaction_th.InsertedBy <- Trx.InsertedDate` |

<a id="kolom-inserteddate"></a>
## InsertedDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TH.sql#L29). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012966 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_DeleteSecurityTransaction_TT:124](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_DeleteSecurityTransaction_TT.sql#L124) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_th.InsertedDate` : `DATETIME` | `td.InsertedDate` : `varchar(10)` | `td.InsertedDate` → `DATETIME` | `retail.dbo.securitytransaction_th.InsertedDate <- td.InsertedDate` |
| IC-013664 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_RejectStatusSecurityTransaction_TT:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_RejectStatusSecurityTransaction_TT.sql#L66) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_th.InsertedDate` : `DATETIME` | `Trx.LastUpdateBy` : `INT` | `Trx.LastUpdateBy` → `DATETIME` | `retail.dbo.securitytransaction_th.InsertedDate <- Trx.LastUpdateBy` |

<a id="kolom-lastupdateby"></a>
## LastUpdateBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TH.sql#L30). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012967 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_DeleteSecurityTransaction_TT:124](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_DeleteSecurityTransaction_TT.sql#L124) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_th.LastUpdateBy` : `INT` | `td.LastUpdateBy` : `varchar(20)` | `td.LastUpdateBy` → `INT` | `retail.dbo.securitytransaction_th.LastUpdateBy <- td.LastUpdateBy` |
| IC-013665 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_RejectStatusSecurityTransaction_TT:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_RejectStatusSecurityTransaction_TT.sql#L66) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_th.LastUpdateBy` : `INT` | `Trx.LastUpdateDate` : `DATETIME` | `Trx.LastUpdateDate` → `INT` | `retail.dbo.securitytransaction_th.LastUpdateBy <- Trx.LastUpdateDate` |

<a id="kolom-lastupdatedate"></a>
## LastUpdateDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TH.sql#L31). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013666 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_RejectStatusSecurityTransaction_TT:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_RejectStatusSecurityTransaction_TT.sql#L66) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_th.LastUpdateDate` : `DATETIME` | `'D'` : `varchar` | `'D'` → `DATETIME` | `retail.dbo.securitytransaction_th.LastUpdateDate <- 'D'` |

<a id="kolom-secid"></a>
## SecId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TH.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012968 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_DeleteSecurityTransaction_TT:124](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_DeleteSecurityTransaction_TT.sql#L124) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_th.SecId` : `INT` | `td.SecId` : `varchar(20)` | `td.SecId` → `INT` | `retail.dbo.securitytransaction_th.SecId <- td.SecId` |

<a id="kolom-settlementdate"></a>
## SettlementDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TH.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012969 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_DeleteSecurityTransaction_TT:124](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_DeleteSecurityTransaction_TT.sql#L124) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_th.SettlementDate` : `DATETIME` | `td.SettlementDate` : `varchar(10)` | `td.SettlementDate` → `DATETIME` | `retail.dbo.securitytransaction_th.SettlementDate <- td.SettlementDate` |

<a id="kolom-spreadyield"></a>
## SpreadYield

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TH.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012970 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_DeleteSecurityTransaction_TT:124](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_DeleteSecurityTransaction_TT.sql#L124) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_th.SpreadYield` : `DECIMAL(20,5)` | `td.SpreadYield` : `numeric` | `td.SpreadYield` → `DECIMAL(20,5)` | `retail.dbo.securitytransaction_th.SpreadYield <- td.SpreadYield` |

<a id="kolom-trxdate"></a>
## TrxDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TH.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012971 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_DeleteSecurityTransaction_TT:124](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_DeleteSecurityTransaction_TT.sql#L124) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_th.TrxDate` : `DATETIME` | `td.TrxDate` : `varchar(10)` | `td.TrxDate` → `DATETIME` | `retail.dbo.securitytransaction_th.TrxDate <- td.TrxDate` |

<a id="kolom-trxstatus"></a>
## TrxStatus

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TH.sql#L25). **Existing leading index:** SecurityTransaction_TH_NN2.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004240 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSInsertRevenueBondJurnalGL:122](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L122) | JOIN_ON / P11 / COMPARISON | `th3.TrxStatus` : `SMALLINT` | `3` : `int` | `th3.TrxStatus` → `INT` | `th3.TrxStatus = 3` |

<a id="kolom-yield"></a>
## Yield

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_TH.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012972 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_DeleteSecurityTransaction_TT:124](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_DeleteSecurityTransaction_TT.sql#L124) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytransaction_th.Yield` : `DECIMAL(20,5)` | `td.Yield` : `numeric` | `td.Yield` → `DECIMAL(20,5)` | `retail.dbo.securitytransaction_th.Yield <- td.Yield` |
