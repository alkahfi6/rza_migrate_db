# Retail UAT.dbo.ORISBNInvestorRekDana_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 15 lokasi source, 7 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 4 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cifid"></a>
## CIFId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISBNInvestorRekDana_TR.sql#L2). **Existing leading index:** ORISBNInvestorRekDana_TR_NN1.

**Frekuensi source:** 9 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002870 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedProcessRekDana:238](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L238) | WHERE / P11 / COMPARISON | `[CIFId]` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `[CIFId] = @nCIFId` |
| IC-002878 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedProcessRekDana:454](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L454) | WHERE / P11 / COMPARISON | `[CIFId]` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `[CIFId] = @nCIFId` |
| IC-002881 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedProcessRekDana:536](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L536) | WHERE / P11 / COMPARISON | `[CIFId]` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `[CIFId] = @nCIFId` |
| IC-002887 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedProcessRekDana:701](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L701) | WHERE / P11 / COMPARISON | `[CIFId]` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `[CIFId] = @nCIFId` |
| IC-002906 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedPushMailNotifOrderCompleted:133](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L133) | WHERE / P11 / COMPARISON | `[CIFId]` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `[CIFId] = @nCIFId` |
| IC-002907 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedPushMailNotifOrderCompleted:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L141) | WHERE / P11 / COMPARISON | `[CIFId]` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `[CIFId] = @nCIFId` |
| IC-002934 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedPushMailNotifOrderORI:150](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderORI.sql#L150) | WHERE / P11 / COMPARISON | `[CIFId]` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `[CIFId] = @nCIFId` |
| IC-002953 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedPushMailNotifRegInvestor:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifRegInvestor.sql#L98) | WHERE / P11 / COMPARISON | `CIFId` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `CIFId = @nCIFId` |
| IC-002954 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedPushMailNotifRegInvestor:103](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifRegInvestor.sql#L103) | WHERE / P11 / COMPARISON | `CIFId` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `CIFId = @nCIFId` |

<a id="kolom-idbank"></a>
## IdBank

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISBNInvestorRekDana_TR.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000999 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.ORISBNIPopulateRegisterInvestor:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateRegisterInvestor.sql#L228) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `rd.IdBank` : `INT` | `IdBank` : `varchar(20)` | `rd.IdBank` → `varchar(20)` | `IdBank = rd.IdBank` |

<a id="kolom-idrekdana"></a>
## IdRekDana

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISBNInvestorRekDana_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001000 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.ORISBNIPopulateRegisterInvestor:229](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateRegisterInvestor.sql#L229) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `rd.IdRekDana` : `VARCHAR(20)` | `IdRekDana` : `int` | `rd.IdRekDana` → `int` | `IdRekDana = rd.IdRekDana` |
| IC-002834 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPopulateRegInvestorForOrder:73](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestorForOrder.sql#L73) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `rd.IdRekDana` : `VARCHAR(20)` | `IdRekDana` : `int` | `rd.IdRekDana` → `int` | `IdRekDana = rd.IdRekDana` |
| IC-004795 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSORIOnlineNewTransaction:160](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlineNewTransaction.sql#L160) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `rk.[IdRekDana]` : `VARCHAR(20)` | `tmp.[IdRekDana]` : `int` | `rk.[IdRekDana]` → `int` | `tmp.[IdRekDana] = rk.[IdRekDana]` |
| IC-002879 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessRekDana:457](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L457) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.orisbninvestorrekdana_tr.IdRekDana` : `VARCHAR(20)` | `[IdRekDana]` : `int` | `[IdRekDana]` → `VARCHAR(20)` | `retail.dbo.orisbninvestorrekdana_tr.IdRekDana <- [IdRekDana]` |
| IC-002882 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessRekDana:544](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L544) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `rek.[IdRekDana]` : `VARCHAR(20)` | `td.[IdRekDana]` : `int` | `td.[IdRekDana]` → `VARCHAR(20)` | `rek.[IdRekDana] = td.[IdRekDana]` |
