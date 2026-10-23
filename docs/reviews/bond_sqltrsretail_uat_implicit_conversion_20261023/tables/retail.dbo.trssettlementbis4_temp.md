# Retail UAT.dbo.TRSSettlementBIS4_temp — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 47 lokasi source, 3 procedure, 19 kolom. **A:** 8 lokasi pada 2 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-buyercashaccount"></a>
## BuyerCashAccount

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4_temp.sql#L28). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 2 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006181 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4Detail:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4Detail.sql#L64) | JOIN_ON / P02 / COMPARISON | `bistemp.BuyerCashAccount` : `VARCHAR(500)` | `tmpTCash.SubAccCode` : `NVARCHAR(35)` | `bistemp.BuyerCashAccount` → `NVARCHAR` | `tmpTCash.SubAccCode = bistemp.BuyerCashAccount` |
| IC-006187 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4RepoDetail:75](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4RepoDetail.sql#L75) | JOIN_ON / P02 / COMPARISON | `bistemp.BuyerCashAccount` : `VARCHAR(500)` | `tmpTCash.SubAccCode` : `NVARCHAR(35)` | `bistemp.BuyerCashAccount` → `NVARCHAR` | `tmpTCash.SubAccCode = bistemp.BuyerCashAccount` |

<a id="kolom-participantbuyer"></a>
## ParticipantBuyer

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4_temp.sql#L24). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 3 SP. **A:** 2 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006182 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4Detail:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4Detail.sql#L66) | JOIN_ON / P02 / COMPARISON | `bistemp.ParticipantBuyer` : `VARCHAR(500)` | `particBuy.ParticipantCode` : `NVARCHAR (combined result; facets unknown)` | `bistemp.ParticipantBuyer` → `NVARCHAR` | `particBuy.ParticipantCode = bistemp.ParticipantBuyer` |
| IC-006188 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4RepoDetail:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4RepoDetail.sql#L77) | JOIN_ON / P02 / COMPARISON | `bistemp.ParticipantBuyer` : `VARCHAR(500)` | `particBuy.ParticipantCode` : `NVARCHAR (combined result; facets unknown)` | `bistemp.ParticipantBuyer` → `NVARCHAR` | `particBuy.ParticipantCode = bistemp.ParticipantBuyer` |
| IC-006035 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L165) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.ParticipantBuyer` : `VARCHAR(500)` | `case when b.ParticipantBuyer is not null then b.ParticipantBuyer else case when tt.TrxType = 4 then bpcs.ParticipantCode else @pcParticipantCode end end` : `NVARCHAR (combined result; facets unknown)` | `case when b.ParticipantBuyer is not null then b.ParticipantBuyer else case when tt.TrxType = 4 then bpcs.ParticipantCode else @pcParticipantCode end end` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.ParticipantBuyer <- case when b.ParticipantBuyer is not null then b.ParticipantBuyer else case when tt.TrxType = 4 then bpcs.ParticipantCode else @pcParticipantCode end end` |
| IC-006063 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L337) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.ParticipantBuyer` : `VARCHAR(500)` | `CASE WHEN b.ParticipantBuyer IS NOT NULL THEN b.ParticipantBuyer ELSE CASE WHEN tt.TrxType = 4 THEN bpcs.ParticipantCode ELSE @pcParticipantCode END END` : `NVARCHAR (combined result; facets unknown)` | `CASE WHEN b.ParticipantBuyer IS NOT NULL THEN b.ParticipantBuyer ELSE CASE WHEN tt.TrxType = 4 THEN bpcs.ParticipantCode ELSE @pcParticipantCode END END` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.ParticipantBuyer <- CASE WHEN b.ParticipantBuyer IS NOT NULL THEN b.ParticipantBuyer ELSE CASE WHEN tt.TrxType = 4 THEN bpcs.ParticipantCode ELSE @pcParticipantCode END END` |

<a id="kolom-participantseller"></a>
## ParticipantSeller

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4_temp.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 3 SP. **A:** 2 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006179 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4Detail:62](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4Detail.sql#L62) | JOIN_ON / P02 / COMPARISON | `bistemp.ParticipantSeller` : `VARCHAR(500)` | `particSell.ParticipantCode` : `NVARCHAR (combined result; facets unknown)` | `bistemp.ParticipantSeller` → `NVARCHAR` | `particSell.ParticipantCode = bistemp.ParticipantSeller` |
| IC-006185 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4RepoDetail:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4RepoDetail.sql#L71) | JOIN_ON / P02 / COMPARISON | `bistemp.ParticipantSeller` : `VARCHAR(500)` | `particSell.ParticipantCode` : `NVARCHAR (combined result; facets unknown)` | `bistemp.ParticipantSeller` → `NVARCHAR` | `particSell.ParticipantCode = bistemp.ParticipantSeller` |
| IC-006036 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L165) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.ParticipantSeller` : `VARCHAR(500)` | `case when b.ParticipantSeller is not null then b.ParticipantSeller else case when tt.TrxType = 4 then @pcParticipantCode else bpcs.ParticipantCode end end` : `NVARCHAR (combined result; facets unknown)` | `case when b.ParticipantSeller is not null then b.ParticipantSeller else case when tt.TrxType = 4 then @pcParticipantCode else bpcs.ParticipantCode end end` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.ParticipantSeller <- case when b.ParticipantSeller is not null then b.ParticipantSeller else case when tt.TrxType = 4 then @pcParticipantCode else bpcs.ParticipantCode end end` |
| IC-006064 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L337) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.ParticipantSeller` : `VARCHAR(500)` | `CASE WHEN b.ParticipantSeller IS NOT NULL THEN b.ParticipantSeller ELSE CASE WHEN tt.TrxType = 4 THEN @pcParticipantCode ELSE bpcs.ParticipantCode END END` : `NVARCHAR (combined result; facets unknown)` | `CASE WHEN b.ParticipantSeller IS NOT NULL THEN b.ParticipantSeller ELSE CASE WHEN tt.TrxType = 4 THEN @pcParticipantCode ELSE bpcs.ParticipantCode END END` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.ParticipantSeller <- CASE WHEN b.ParticipantSeller IS NOT NULL THEN b.ParticipantSeller ELSE CASE WHEN tt.TrxType = 4 THEN @pcParticipantCode ELSE bpcs.ParticipantCode END END` |

<a id="kolom-sellercashaccount"></a>
## SellerCashAccount

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4_temp.sql#L26). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 2 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006180 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4Detail:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4Detail.sql#L63) | JOIN_ON / P02 / COMPARISON | `bistemp.SellerCashAccount` : `VARCHAR(500)` | `tmpFCash.SubAccCode` : `NVARCHAR(35)` | `bistemp.SellerCashAccount` → `NVARCHAR` | `tmpFCash.SubAccCode = bistemp.SellerCashAccount` |
| IC-006186 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4RepoDetail:73](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4RepoDetail.sql#L73) | JOIN_ON / P02 / COMPARISON | `bistemp.SellerCashAccount` : `VARCHAR(500)` | `tmpFCash.SubAccCode` : `NVARCHAR(35)` | `bistemp.SellerCashAccount` → `NVARCHAR` | `tmpFCash.SubAccCode = bistemp.SellerCashAccount` |

<a id="kolom-accountbuyer"></a>
## AccountBuyer

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4_temp.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006031 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L165) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.AccountBuyer` : `VARCHAR(500)` | `case when b.AccountBuyer is not null then b.AccountBuyer else case when tt.TrxType = 4 then bpcs.ParticipantCode else @pcAccount end end` : `NVARCHAR (combined result; facets unknown)` | `case when b.AccountBuyer is not null then b.AccountBuyer else case when tt.TrxType = 4 then bpcs.ParticipantCode else @pcAccount end end` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.AccountBuyer <- case when b.AccountBuyer is not null then b.AccountBuyer else case when tt.TrxType = 4 then bpcs.ParticipantCode else @pcAccount end end` |
| IC-006058 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L337) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.AccountBuyer` : `VARCHAR(500)` | `CASE WHEN b.AccountBuyer IS NOT NULL THEN b.AccountBuyer ELSE CASE WHEN tt.TrxType = 4 THEN bpcs.ParticipantCode ELSE @pcAccount END END` : `NVARCHAR (combined result; facets unknown)` | `CASE WHEN b.AccountBuyer IS NOT NULL THEN b.AccountBuyer ELSE CASE WHEN tt.TrxType = 4 THEN bpcs.ParticipantCode ELSE @pcAccount END END` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.AccountBuyer <- CASE WHEN b.AccountBuyer IS NOT NULL THEN b.AccountBuyer ELSE CASE WHEN tt.TrxType = 4 THEN bpcs.ParticipantCode ELSE @pcAccount END END` |

<a id="kolom-accountseller"></a>
## AccountSeller

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4_temp.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006032 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L165) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.AccountSeller` : `VARCHAR(500)` | `case when b.AccountSeller is not null then b.AccountSeller else case when tt.TrxType = 4 then @pcAccount else bpcs.ParticipantCode end end` : `NVARCHAR (combined result; facets unknown)` | `case when b.AccountSeller is not null then b.AccountSeller else case when tt.TrxType = 4 then @pcAccount else bpcs.ParticipantCode end end` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.AccountSeller <- case when b.AccountSeller is not null then b.AccountSeller else case when tt.TrxType = 4 then @pcAccount else bpcs.ParticipantCode end end` |
| IC-006059 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L337) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.AccountSeller` : `VARCHAR(500)` | `CASE WHEN b.AccountSeller IS NOT NULL THEN b.AccountSeller ELSE CASE WHEN tt.TrxType = 4 THEN @pcAccount ELSE bpcs.ParticipantCode END END` : `NVARCHAR (combined result; facets unknown)` | `CASE WHEN b.AccountSeller IS NOT NULL THEN b.AccountSeller ELSE CASE WHEN tt.TrxType = 4 THEN @pcAccount ELSE bpcs.ParticipantCode END END` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.AccountSeller <- CASE WHEN b.AccountSeller IS NOT NULL THEN b.AccountSeller ELSE CASE WHEN tt.TrxType = 4 THEN @pcAccount ELSE bpcs.ParticipantCode END END` |

<a id="kolom-buyerpaymentagent"></a>
## BuyerPaymentAgent

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4_temp.sql#L27). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006033 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L165) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.BuyerPaymentAgent` : `VARCHAR(500)` | `CASE WHEN BankCode = '03' and tt.TrxType = 4 THEN bpcs.PayCode ELSE @pcPaymentAgent END` : `NVARCHAR (combined result; facets unknown)` | `CASE WHEN BankCode = '03' and tt.TrxType = 4 THEN bpcs.PayCode ELSE @pcPaymentAgent END` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.BuyerPaymentAgent <- CASE WHEN BankCode = '03' and tt.TrxType = 4 THEN bpcs.PayCode ELSE @pcPaymentAgent END` |
| IC-006060 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L337) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.BuyerPaymentAgent` : `VARCHAR(500)` | `CASE WHEN BankCode = '03' AND tt.TrxType = 4 THEN bpcs.PayCode ELSE @pcPaymentAgent END` : `NVARCHAR (combined result; facets unknown)` | `CASE WHEN BankCode = '03' AND tt.TrxType = 4 THEN bpcs.PayCode ELSE @pcPaymentAgent END` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.BuyerPaymentAgent <- CASE WHEN BankCode = '03' AND tt.TrxType = 4 THEN bpcs.PayCode ELSE @pcPaymentAgent END` |

<a id="kolom-codeinstrument"></a>
## CodeInstrument

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4_temp.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006087 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:502](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L502) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.CodeInstrument` : `VARCHAR(500)` | `mst.InstrumentCode` : `NVARCHAR(128)` | `mst.InstrumentCode` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.CodeInstrument <- mst.InstrumentCode` |
| IC-006097 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:682](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L682) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.CodeInstrument` : `VARCHAR(500)` | `mst.InstrumentCode` : `NVARCHAR(128)` | `mst.InstrumentCode` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.CodeInstrument <- mst.InstrumentCode` |

<a id="kolom-deliveryagent"></a>
## DeliveryAgent

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4_temp.sql#L19). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006034 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L165) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.DeliveryAgent` : `VARCHAR(500)` | `case when b.DeliveryAgent is not null then b.DeliveryAgent else case when tt.TrxType = 4 then @pcAccount else bpcs.PayCode end end` : `NVARCHAR (combined result; facets unknown)` | `case when b.DeliveryAgent is not null then b.DeliveryAgent else case when tt.TrxType = 4 then @pcAccount else bpcs.PayCode end end` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.DeliveryAgent <- case when b.DeliveryAgent is not null then b.DeliveryAgent else case when tt.TrxType = 4 then @pcAccount else bpcs.PayCode end end` |
| IC-006061 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L337) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.DeliveryAgent` : `VARCHAR(500)` | `CASE WHEN b.DeliveryAgent IS NOT NULL THEN b.DeliveryAgent ELSE CASE WHEN tt.TrxType = 4 THEN @pcAccount ELSE bpcs.PayCode END END` : `NVARCHAR (combined result; facets unknown)` | `CASE WHEN b.DeliveryAgent IS NOT NULL THEN b.DeliveryAgent ELSE CASE WHEN tt.TrxType = 4 THEN @pcAccount ELSE bpcs.PayCode END END` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.DeliveryAgent <- CASE WHEN b.DeliveryAgent IS NOT NULL THEN b.DeliveryAgent ELSE CASE WHEN tt.TrxType = 4 THEN @pcAccount ELSE bpcs.PayCode END END` |

<a id="kolom-murexno"></a>
## MurexNo

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4_temp.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006183 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBIS4Detail:67](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4Detail.sql#L67) | WHERE / P11 / COMPARISON | `bistemp.MurexNo` : `BIGINT` | `@pnMurexNo` : `int` | `@pnMurexNo` → `BIGINT` | `bistemp.MurexNo = @pnMurexNo` |
| IC-006189 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBIS4RepoDetail:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4RepoDetail.sql#L80) | WHERE / P11 / COMPARISON | `bistemp.MurexNo` : `BIGINT` | `@pnMurexNo` : `int` | `@pnMurexNo` → `BIGINT` | `bistemp.MurexNo = @pnMurexNo` |

<a id="kolom-noctp"></a>
## NoCTP

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4_temp.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006062 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L337) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.NoCTP` : `VARCHAR(500)` | `CASE WHEN b.NoCTP IS NOT NULL THEN b.NoCTP ELSE tt.DealIdExternal END` : `BIGINT (combined result; facets unknown)` | `CASE WHEN b.NoCTP IS NOT NULL THEN b.NoCTP ELSE tt.DealIdExternal END` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.NoCTP <- CASE WHEN b.NoCTP IS NOT NULL THEN b.NoCTP ELSE tt.DealIdExternal END` |

<a id="kolom-receivingagent"></a>
## ReceivingAgent

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4_temp.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006037 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L165) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.ReceivingAgent` : `VARCHAR(500)` | `case when b.ReceivingAgent is not null then b.ReceivingAgent else case when tt.TrxType = 4 then bpcs.PayCode else @pcAccount end end` : `NVARCHAR (combined result; facets unknown)` | `case when b.ReceivingAgent is not null then b.ReceivingAgent else case when tt.TrxType = 4 then bpcs.PayCode else @pcAccount end end` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.ReceivingAgent <- case when b.ReceivingAgent is not null then b.ReceivingAgent else case when tt.TrxType = 4 then bpcs.PayCode else @pcAccount end end` |
| IC-006065 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L337) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.ReceivingAgent` : `VARCHAR(500)` | `CASE WHEN b.ReceivingAgent IS NOT NULL THEN b.ReceivingAgent ELSE CASE WHEN tt.TrxType = 4 THEN bpcs.PayCode ELSE @pcAccount END END` : `NVARCHAR (combined result; facets unknown)` | `CASE WHEN b.ReceivingAgent IS NOT NULL THEN b.ReceivingAgent ELSE CASE WHEN tt.TrxType = 4 THEN bpcs.PayCode ELSE @pcAccount END END` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.ReceivingAgent <- CASE WHEN b.ReceivingAgent IS NOT NULL THEN b.ReceivingAgent ELSE CASE WHEN tt.TrxType = 4 THEN bpcs.PayCode ELSE @pcAccount END END` |

<a id="kolom-sellerpaymentagent"></a>
## SellerPaymentAgent

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4_temp.sql#L25). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006038 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L165) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.SellerPaymentAgent` : `VARCHAR(500)` | `CASE WHEN BankCode = '03' and tt.TrxType != 4 THEN bpcs.PayCode ELSE @pcPaymentAgent END` : `NVARCHAR (combined result; facets unknown)` | `CASE WHEN BankCode = '03' and tt.TrxType != 4 THEN bpcs.PayCode ELSE @pcPaymentAgent END` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.SellerPaymentAgent <- CASE WHEN BankCode = '03' and tt.TrxType != 4 THEN bpcs.PayCode ELSE @pcPaymentAgent END` |
| IC-006066 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L337) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.SellerPaymentAgent` : `VARCHAR(500)` | `CASE WHEN BankCode = '03' AND tt.TrxType != 4 THEN bpcs.PayCode ELSE @pcPaymentAgent END` : `NVARCHAR (combined result; facets unknown)` | `CASE WHEN BankCode = '03' AND tt.TrxType != 4 THEN bpcs.PayCode ELSE @pcPaymentAgent END` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.SellerPaymentAgent <- CASE WHEN BankCode = '03' AND tt.TrxType != 4 THEN bpcs.PayCode ELSE @pcPaymentAgent END` |

<a id="kolom-settlementdate"></a>
## SettlementDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4_temp.sql#L32). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006088 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:502](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L502) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.SettlementDate` : `DATETIME` | `convert(varchar(8),bondt.M_SE_PAY_D,112)` : `varchar(8)` | `convert(varchar(8),bondt.M_SE_PAY_D,112)` → `DATETIME` | `retail.dbo.trssettlementbis4_temp.SettlementDate <- convert(varchar(8),bondt.M_SE_PAY_D,112)` |
| IC-006098 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:682](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L682) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.SettlementDate` : `DATETIME` | `CONVERT(VARCHAR(8),bondt.M_SE_PAY_D,112)` : `VARCHAR(8)` | `CONVERT(VARCHAR(8),bondt.M_SE_PAY_D,112)` → `DATETIME` | `retail.dbo.trssettlementbis4_temp.SettlementDate <- CONVERT(VARCHAR(8),bondt.M_SE_PAY_D,112)` |

<a id="kolom-shortcodeinstrument"></a>
## ShortCodeInstrument

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4_temp.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006039 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L165) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.ShortCodeInstrument` : `VARCHAR(500)` | `ins.TypeCode` : `NVARCHAR(100)` | `ins.TypeCode` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.ShortCodeInstrument <- ins.TypeCode` |
| IC-006067 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L337) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.ShortCodeInstrument` : `VARCHAR(500)` | `ins.TypeCode` : `NVARCHAR(100)` | `ins.TypeCode` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.ShortCodeInstrument <- ins.TypeCode` |
| IC-006089 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:502](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L502) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.ShortCodeInstrument` : `VARCHAR(500)` | `mst.TypeCode` : `NVARCHAR(100)` | `mst.TypeCode` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.ShortCodeInstrument <- mst.TypeCode` |
| IC-006099 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:682](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L682) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.ShortCodeInstrument` : `VARCHAR(500)` | `mst.TypeCode` : `NVARCHAR(100)` | `mst.TypeCode` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.ShortCodeInstrument <- mst.TypeCode` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4_temp.sql#L33). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006040 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L165) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.Status` : `CHAR(1)` | `isnull(b.Status,'N')` : `VARCHAR(19)` | `isnull(b.Status,'N')` → `CHAR(1)` | `retail.dbo.trssettlementbis4_temp.Status <- isnull(b.Status,'N')` |
| IC-006068 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L337) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.Status` : `CHAR(1)` | `ISNULL(b.Status,'N')` : `VARCHAR(19)` | `ISNULL(b.Status,'N')` → `CHAR(1)` | `retail.dbo.trssettlementbis4_temp.Status <- ISNULL(b.Status,'N')` |
| IC-006090 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:502](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L502) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.Status` : `CHAR(1)` | `isnull(b.Status,'N')` : `VARCHAR(19)` | `isnull(b.Status,'N')` → `CHAR(1)` | `retail.dbo.trssettlementbis4_temp.Status <- isnull(b.Status,'N')` |
| IC-006100 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:682](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L682) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.Status` : `CHAR(1)` | `ISNULL(b.Status,'N')` : `VARCHAR(19)` | `ISNULL(b.Status,'N')` → `CHAR(1)` | `retail.dbo.trssettlementbis4_temp.Status <- ISNULL(b.Status,'N')` |

<a id="kolom-subaccountbuyer"></a>
## SubAccountBuyer

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4_temp.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006041 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L165) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.SubAccountBuyer` : `VARCHAR(500)` | `case when b.SubAccountBuyer is not null then b.SubAccountBuyer else case when tt.TrxType = 4 then bsubacc.SubAccCode else @pcSubAccount end end` : `NVARCHAR (combined result; facets unknown)` | `case when b.SubAccountBuyer is not null then b.SubAccountBuyer else case when tt.TrxType = 4 then bsubacc.SubAccCode else @pcSubAccount end end` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.SubAccountBuyer <- case when b.SubAccountBuyer is not null then b.SubAccountBuyer else case when tt.TrxType = 4 then bsubacc.SubAccCode else @pcSubAccount end end` |
| IC-006069 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L337) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.SubAccountBuyer` : `VARCHAR(500)` | `CASE WHEN b.SubAccountBuyer IS NOT NULL THEN b.SubAccountBuyer ELSE CASE WHEN tt.TrxType = 4 THEN bsubacc.SubAccCode ELSE @pcSubAccount END END` : `NVARCHAR (combined result; facets unknown)` | `CASE WHEN b.SubAccountBuyer IS NOT NULL THEN b.SubAccountBuyer ELSE CASE WHEN tt.TrxType = 4 THEN bsubacc.SubAccCode ELSE @pcSubAccount END END` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.SubAccountBuyer <- CASE WHEN b.SubAccountBuyer IS NOT NULL THEN b.SubAccountBuyer ELSE CASE WHEN tt.TrxType = 4 THEN bsubacc.SubAccCode ELSE @pcSubAccount END END` |

<a id="kolom-subaccountseller"></a>
## SubAccountSeller

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4_temp.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006042 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L165) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.SubAccountSeller` : `VARCHAR(500)` | `case when b.SubAccountSeller is not null then b.SubAccountSeller else case when tt.TrxType = 4 then @pcSubAccount else bsubacc.SubAccCode end end` : `NVARCHAR (combined result; facets unknown)` | `case when b.SubAccountSeller is not null then b.SubAccountSeller else case when tt.TrxType = 4 then @pcSubAccount else bsubacc.SubAccCode end end` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.SubAccountSeller <- case when b.SubAccountSeller is not null then b.SubAccountSeller else case when tt.TrxType = 4 then @pcSubAccount else bsubacc.SubAccCode end end` |
| IC-006070 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L337) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.SubAccountSeller` : `VARCHAR(500)` | `CASE WHEN b.SubAccountSeller IS NOT NULL THEN b.SubAccountSeller ELSE CASE WHEN tt.TrxType = 4 THEN @pcSubAccount ELSE bsubacc.SubAccCode END END` : `NVARCHAR (combined result; facets unknown)` | `CASE WHEN b.SubAccountSeller IS NOT NULL THEN b.SubAccountSeller ELSE CASE WHEN tt.TrxType = 4 THEN @pcSubAccount ELSE bsubacc.SubAccCode END END` → `VARCHAR(500)` | `retail.dbo.trssettlementbis4_temp.SubAccountSeller <- CASE WHEN b.SubAccountSeller IS NOT NULL THEN b.SubAccountSeller ELSE CASE WHEN tt.TrxType = 4 THEN @pcSubAccount ELSE bsubacc.SubAccCode END END` |

<a id="kolom-tradedate"></a>
## TradeDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4_temp.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006043 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L165) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.TradeDate` : `DATETIME` | `convert(varchar(8),tt.TrxDate,112)` : `varchar(8)` | `convert(varchar(8),tt.TrxDate,112)` → `DATETIME` | `retail.dbo.trssettlementbis4_temp.TradeDate <- convert(varchar(8),tt.TrxDate,112)` |
| IC-006071 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L337) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.TradeDate` : `DATETIME` | `CONVERT(VARCHAR(8),tt.TrxDate,112)` : `VARCHAR(8)` | `CONVERT(VARCHAR(8),tt.TrxDate,112)` → `DATETIME` | `retail.dbo.trssettlementbis4_temp.TradeDate <- CONVERT(VARCHAR(8),tt.TrxDate,112)` |
| IC-006091 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:502](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L502) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.TradeDate` : `DATETIME` | `convert(varchar(8),bondt.M_TRN_DATE,112)` : `varchar(8)` | `convert(varchar(8),bondt.M_TRN_DATE,112)` → `DATETIME` | `retail.dbo.trssettlementbis4_temp.TradeDate <- convert(varchar(8),bondt.M_TRN_DATE,112)` |
| IC-006101 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPopulateSettlementBIS4:682](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L682) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_temp.TradeDate` : `DATETIME` | `CONVERT(VARCHAR(8),bondt.M_TRN_DATE,112)` : `VARCHAR(8)` | `CONVERT(VARCHAR(8),bondt.M_TRN_DATE,112)` → `DATETIME` | `retail.dbo.trssettlementbis4_temp.TradeDate <- CONVERT(VARCHAR(8),bondt.M_TRN_DATE,112)` |
