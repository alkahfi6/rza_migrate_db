# Retail UAT.dbo.TRSSettlementBIS4 — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 52 lokasi source, 4 procedure, 13 kolom. **A:** 2 lokasi pada 1 procedure. **B:** 39 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-senderref"></a>
## SenderRef

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003598 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSCheckStatusBIS4:38](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCheckStatusBIS4.sql#L38) | JOIN_ON / P02 / COMPARISON | `trs.SenderRef` : `VARCHAR(500)` | `tempBIS4.trn` : `NVARCHAR(30)` | `trs.SenderRef` → `NVARCHAR` | `trs.SenderRef = tempBIS4.trn` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `VARCHAR(19)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4.sql#L32). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 12 lokasi / 3 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003599 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSCheckStatusBIS4:39](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCheckStatusBIS4.sql#L39) | WHERE / P02 / COMPARISON | `trs.Status` : `VARCHAR(19)` | `tempBIS4.step_code` : `NVARCHAR(30)` | `trs.Status` → `NVARCHAR` | `trs.Status <> tempBIS4.step_code` |
| IC-006040 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateSettlementBIS4:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L165) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `isnull(b.Status,'N')` : `VARCHAR(19)` | `retail.dbo.trssettlementbis4_temp.Status` : `CHAR(1)` | `isnull(b.Status,'N')` → `CHAR(1)` | `retail.dbo.trssettlementbis4_temp.Status <- isnull(b.Status,'N')` |
| IC-006068 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateSettlementBIS4:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L337) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ISNULL(b.Status,'N')` : `VARCHAR(19)` | `retail.dbo.trssettlementbis4_temp.Status` : `CHAR(1)` | `ISNULL(b.Status,'N')` → `CHAR(1)` | `retail.dbo.trssettlementbis4_temp.Status <- ISNULL(b.Status,'N')` |
| IC-006090 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateSettlementBIS4:502](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L502) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `isnull(b.Status,'N')` : `VARCHAR(19)` | `retail.dbo.trssettlementbis4_temp.Status` : `CHAR(1)` | `isnull(b.Status,'N')` → `CHAR(1)` | `retail.dbo.trssettlementbis4_temp.Status <- isnull(b.Status,'N')` |
| IC-006100 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateSettlementBIS4:682](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L682) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ISNULL(b.Status,'N')` : `VARCHAR(19)` | `retail.dbo.trssettlementbis4_temp.Status` : `CHAR(1)` | `ISNULL(b.Status,'N')` → `CHAR(1)` | `retail.dbo.trssettlementbis4_temp.Status <- ISNULL(b.Status,'N')` |
| IC-006115 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateSettlementBIS4:899](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L899) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `isnull(b.Status,'N')` : `VARCHAR(19)` | `#tmpDispTrx.Status` : `[char](1)` | `isnull(b.Status,'N')` → `[char](1)` | `#tmpDispTrx.Status <- isnull(b.Status,'N')` |
| IC-006140 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateSettlementBIS4:1028](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1028) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ISNULL(b.Status,'N')` : `VARCHAR(19)` | `#tmpDispTrx.Status` : `[char](1)` | `ISNULL(b.Status,'N')` → `[char](1)` | `#tmpDispTrx.Status <- ISNULL(b.Status,'N')` |
| IC-006159 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateSettlementBIS4:1151](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1151) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `isnull(b.Status,'N')` : `VARCHAR(19)` | `#tmpDispTrx.Status` : `[char](1)` | `isnull(b.Status,'N')` → `[char](1)` | `#tmpDispTrx.Status <- isnull(b.Status,'N')` |
| IC-006170 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateSettlementBIS4:1289](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1289) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ISNULL(b.Status,'N')` : `VARCHAR(19)` | `#tmpDispTrx.Status` : `[char](1)` | `ISNULL(b.Status,'N')` → `[char](1)` | `#tmpDispTrx.Status <- ISNULL(b.Status,'N')` |
| IC-006693 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessApprovalSettlementBIS4:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalSettlementBIS4.sql#L30) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `Status` : `VARCHAR(19)` | `retail.dbo.trssettlementbis4_th.Status` : `CHAR(1)` | `Status` → `CHAR(1)` | `retail.dbo.trssettlementbis4_th.Status <- Status` |
| IC-006694 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessApprovalSettlementBIS4:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalSettlementBIS4.sql#L51) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `Status` : `VARCHAR(19)` | `retail.dbo.trssettlementbis4_th.Status` : `CHAR(1)` | `Status` → `CHAR(1)` | `retail.dbo.trssettlementbis4_th.Status <- Status` |
| IC-003597 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSCheckStatusBIS4:35](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCheckStatusBIS4.sql#L35) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `trs.Status` : `VARCHAR(19)` | `tempBIS4.step_code` : `NVARCHAR(30)` | `tempBIS4.step_code` → `VARCHAR(19)` | `trs.Status = tempBIS4.step_code` |

<a id="kolom-accountbuyer"></a>
## AccountBuyer

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4.sql#L19). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006047 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:244](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L244) | EXPRESSION / P07 / CASE_BRANCH | `b.AccountBuyer` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.AccountBuyer` → `NVARCHAR` | `case when b.AccountBuyer is not null then b.AccountBuyer else case when tt.TrxType = 4 then bpcs.ParticipantCode else @pcAccount end end` |
| IC-006076 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:413](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L413) | EXPRESSION / P07 / CASE_BRANCH | `b.AccountBuyer` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.AccountBuyer` → `NVARCHAR` | `CASE WHEN b.AccountBuyer IS NOT NULL THEN b.AccountBuyer ELSE CASE WHEN tt.TrxType = 4 THEN bpcs.ParticipantCode ELSE @pcAccount END END` |

<a id="kolom-accountseller"></a>
## AccountSeller

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006048 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:252](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L252) | EXPRESSION / P07 / CASE_BRANCH | `b.AccountSeller` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.AccountSeller` → `NVARCHAR` | `case when b.AccountSeller is not null then b.AccountSeller else case when tt.TrxType = 4 then @pcAccount else bpcs.ParticipantCode end end` |
| IC-006077 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:421](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L421) | EXPRESSION / P07 / CASE_BRANCH | `b.AccountSeller` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.AccountSeller` → `NVARCHAR` | `CASE WHEN b.AccountSeller IS NOT NULL THEN b.AccountSeller ELSE CASE WHEN tt.TrxType = 4 THEN @pcAccount ELSE bpcs.ParticipantCode END END` |

<a id="kolom-deliveryagent"></a>
## DeliveryAgent

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006046 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:236](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L236) | EXPRESSION / P07 / CASE_BRANCH | `b.DeliveryAgent` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.DeliveryAgent` → `NVARCHAR` | `case when b.DeliveryAgent is not null then b.DeliveryAgent else case when tt.TrxType = 4 then @pcAccount else bpcs.PayCode end end` |
| IC-006075 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:405](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L405) | EXPRESSION / P07 / CASE_BRANCH | `b.DeliveryAgent` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.DeliveryAgent` → `NVARCHAR` | `CASE WHEN b.DeliveryAgent IS NOT NULL THEN b.DeliveryAgent ELSE CASE WHEN tt.TrxType = 4 THEN @pcAccount ELSE bpcs.PayCode END END` |
| IC-006121 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:953](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L953) | EXPRESSION / P07 / CASE_BRANCH | `b.DeliveryAgent` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.DeliveryAgent` → `NVARCHAR` | `case when b.DeliveryAgent is not null then b.DeliveryAgent else case when tt.TrxType = 4 then @pcAccount else bpcs.PayCode end end` |
| IC-006146 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:1076](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1076) | EXPRESSION / P07 / CASE_BRANCH | `b.DeliveryAgent` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.DeliveryAgent` → `NVARCHAR` | `CASE WHEN b.DeliveryAgent IS NOT NULL THEN b.DeliveryAgent ELSE CASE WHEN tt.TrxType = 4 THEN @pcAccount ELSE bpcs.PayCode END END` |

<a id="kolom-generatedate"></a>
## GenerateDate

**Tipe dasar:** `VARCHAR(100)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007414 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementBIS4:179](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBIS4.sql#L179) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `b.GenerateDate` : `VARCHAR(100)` | `tmp.GenerateDate` : `datetime` | `tmp.GenerateDate` → `VARCHAR(100)` | `b.GenerateDate = tmp.GenerateDate` |
| IC-007417 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementBIS4:222](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBIS4.sql#L222) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4.GenerateDate` : `VARCHAR(100)` | `tmp.GenerateDate` : `datetime` | `tmp.GenerateDate` → `VARCHAR(100)` | `retail.dbo.trssettlementbis4.GenerateDate <- tmp.GenerateDate` |

<a id="kolom-isinterbank"></a>
## IsInterbank

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006057 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:330](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L330) | JOIN_ON / P11 / COMPARISON | `b.IsInterbank` : `BIT` | `0` : `int` | `b.IsInterbank` → `INT` | `b.IsInterbank = 0` |
| IC-006086 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:496](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L496) | JOIN_ON / P11 / COMPARISON | `b.IsInterbank` : `BIT` | `0` : `int` | `b.IsInterbank` → `INT` | `b.IsInterbank = 0` |
| IC-006096 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:675](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L675) | JOIN_ON / P11 / COMPARISON | `b.IsInterbank` : `BIT` | `1` : `int` | `b.IsInterbank` → `INT` | `b.IsInterbank = 1` |
| IC-006106 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:854](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L854) | JOIN_ON / P11 / COMPARISON | `b.IsInterbank` : `BIT` | `1` : `int` | `b.IsInterbank` → `INT` | `b.IsInterbank = 1` |
| IC-006130 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:1021](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1021) | JOIN_ON / P11 / COMPARISON | `b.IsInterbank` : `BIT` | `0` : `int` | `b.IsInterbank` → `INT` | `b.IsInterbank = 0` |
| IC-006155 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:1145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1145) | JOIN_ON / P11 / COMPARISON | `b.IsInterbank` : `BIT` | `0` : `int` | `b.IsInterbank` → `INT` | `b.IsInterbank = 0` |
| IC-006165 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:1282](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1282) | JOIN_ON / P11 / COMPARISON | `b.IsInterbank` : `BIT` | `1` : `int` | `b.IsInterbank` → `INT` | `b.IsInterbank = 1` |
| IC-006176 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:1417](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1417) | JOIN_ON / P11 / COMPARISON | `b.IsInterbank` : `BIT` | `1` : `int` | `b.IsInterbank` → `INT` | `b.IsInterbank = 1` |

<a id="kolom-noctp"></a>
## NoCTP

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006072 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:376](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L376) | EXPRESSION / P07 / CASE_BRANCH | `b.NoCTP` : `VARCHAR(500)` | `result type` : `BIGINT (combined result; facets unknown)` | `b.NoCTP` → `BIGINT` | `CASE WHEN b.NoCTP IS NOT NULL THEN b.NoCTP ELSE tt.DealIdExternal END` |

<a id="kolom-participantbuyer"></a>
## ParticipantBuyer

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006051 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:276](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L276) | EXPRESSION / P07 / CASE_BRANCH | `b.ParticipantBuyer` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.ParticipantBuyer` → `NVARCHAR` | `case when b.ParticipantBuyer is not null then b.ParticipantBuyer else case when tt.TrxType = 4 then bpcs.ParticipantCode else @pcParticipantCode end end` |
| IC-006080 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:445](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L445) | EXPRESSION / P07 / CASE_BRANCH | `b.ParticipantBuyer` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.ParticipantBuyer` → `NVARCHAR` | `CASE WHEN b.ParticipantBuyer IS NOT NULL THEN b.ParticipantBuyer ELSE CASE WHEN tt.TrxType = 4 THEN bpcs.ParticipantCode ELSE @pcParticipantCode END END` |
| IC-006122 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:969](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L969) | EXPRESSION / P07 / CASE_BRANCH | `b.ParticipantBuyer` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.ParticipantBuyer` → `NVARCHAR` | `case when b.ParticipantBuyer is not null then b.ParticipantBuyer else case when tt.TrxType = 4 then bpcs.ParticipantCode else @pcParticipantCode end end` |
| IC-006147 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:1092](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1092) | EXPRESSION / P07 / CASE_BRANCH | `b.ParticipantBuyer` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.ParticipantBuyer` → `NVARCHAR` | `CASE WHEN b.ParticipantBuyer IS NOT NULL THEN b.ParticipantBuyer ELSE CASE WHEN tt.TrxType = 4 THEN bpcs.ParticipantCode ELSE @pcParticipantCode END END` |

<a id="kolom-participantseller"></a>
## ParticipantSeller

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006044 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:220](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L220) | EXPRESSION / P07 / CASE_BRANCH | `b.ParticipantSeller` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.ParticipantSeller` → `NVARCHAR` | `case when b.ParticipantSeller is not null then b.ParticipantSeller else case when tt.TrxType = 4 then @pcParticipantCode else bpcs.ParticipantCode end end` |
| IC-006073 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:389](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L389) | EXPRESSION / P07 / CASE_BRANCH | `b.ParticipantSeller` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.ParticipantSeller` → `NVARCHAR` | `CASE WHEN b.ParticipantSeller IS NOT NULL THEN b.ParticipantSeller ELSE CASE WHEN tt.TrxType = 4 THEN @pcParticipantCode ELSE bpcs.ParticipantCode END END` |
| IC-006119 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:937](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L937) | EXPRESSION / P07 / CASE_BRANCH | `b.ParticipantSeller` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.ParticipantSeller` → `NVARCHAR` | `case when b.ParticipantSeller is not null then b.ParticipantSeller else case when tt.TrxType = 4 then @pcParticipantCode else bpcs.ParticipantCode end end` |
| IC-006144 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:1060](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1060) | EXPRESSION / P07 / CASE_BRANCH | `b.ParticipantSeller` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.ParticipantSeller` → `NVARCHAR` | `CASE WHEN b.ParticipantSeller IS NOT NULL THEN b.ParticipantSeller ELSE CASE WHEN tt.TrxType = 4 THEN @pcParticipantCode ELSE bpcs.ParticipantCode END END` |

<a id="kolom-receivingagent"></a>
## ReceivingAgent

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006045 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L228) | EXPRESSION / P07 / CASE_BRANCH | `b.ReceivingAgent` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.ReceivingAgent` → `NVARCHAR` | `case when b.ReceivingAgent is not null then b.ReceivingAgent else case when tt.TrxType = 4 then bpcs.PayCode else @pcAccount end end` |
| IC-006074 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:397](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L397) | EXPRESSION / P07 / CASE_BRANCH | `b.ReceivingAgent` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.ReceivingAgent` → `NVARCHAR` | `CASE WHEN b.ReceivingAgent IS NOT NULL THEN b.ReceivingAgent ELSE CASE WHEN tt.TrxType = 4 THEN bpcs.PayCode ELSE @pcAccount END END` |
| IC-006124 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:985](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L985) | EXPRESSION / P07 / CASE_BRANCH | `b.ReceivingAgent` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.ReceivingAgent` → `NVARCHAR` | `case when b.ReceivingAgent is not null then b.ReceivingAgent else case when tt.TrxType = 4 then bpcs.PayCode else @pcAccount end end` |
| IC-006149 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:1108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1108) | EXPRESSION / P07 / CASE_BRANCH | `b.ReceivingAgent` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.ReceivingAgent` → `NVARCHAR` | `CASE WHEN b.ReceivingAgent IS NOT NULL THEN b.ReceivingAgent ELSE CASE WHEN tt.TrxType = 4 THEN bpcs.PayCode ELSE @pcAccount END END` |

<a id="kolom-subaccountbuyer"></a>
## SubAccountBuyer

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006049 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:260](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L260) | EXPRESSION / P07 / CASE_BRANCH | `b.SubAccountBuyer` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.SubAccountBuyer` → `NVARCHAR` | `case when b.SubAccountBuyer is not null then b.SubAccountBuyer else case when tt.TrxType = 4 then bsubacc.SubAccCode else @pcSubAccount end end` |
| IC-006078 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:429](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L429) | EXPRESSION / P07 / CASE_BRANCH | `b.SubAccountBuyer` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.SubAccountBuyer` → `NVARCHAR` | `CASE WHEN b.SubAccountBuyer IS NOT NULL THEN b.SubAccountBuyer ELSE CASE WHEN tt.TrxType = 4 THEN bsubacc.SubAccCode ELSE @pcSubAccount END END` |
| IC-006123 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:977](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L977) | EXPRESSION / P07 / CASE_BRANCH | `b.SubAccountBuyer` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.SubAccountBuyer` → `NVARCHAR` | `case when b.SubAccountBuyer is not null then b.SubAccountBuyer else case when tt.TrxType = 4 then bsubacc.SubAccCode else @pcSubAccount end end` |
| IC-006148 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:1100](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1100) | EXPRESSION / P07 / CASE_BRANCH | `b.SubAccountBuyer` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.SubAccountBuyer` → `NVARCHAR` | `CASE WHEN b.SubAccountBuyer IS NOT NULL THEN b.SubAccountBuyer ELSE CASE WHEN tt.TrxType = 4 THEN bsubacc.SubAccCode ELSE @pcSubAccount END END` |

<a id="kolom-subaccountseller"></a>
## SubAccountSeller

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006050 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:268](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L268) | EXPRESSION / P07 / CASE_BRANCH | `b.SubAccountSeller` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.SubAccountSeller` → `NVARCHAR` | `case when b.SubAccountSeller is not null then b.SubAccountSeller else case when tt.TrxType = 4 then @pcSubAccount else bsubacc.SubAccCode end end` |
| IC-006079 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:437](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L437) | EXPRESSION / P07 / CASE_BRANCH | `b.SubAccountSeller` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.SubAccountSeller` → `NVARCHAR` | `CASE WHEN b.SubAccountSeller IS NOT NULL THEN b.SubAccountSeller ELSE CASE WHEN tt.TrxType = 4 THEN @pcSubAccount ELSE bsubacc.SubAccCode END END` |
| IC-006120 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:945](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L945) | EXPRESSION / P07 / CASE_BRANCH | `b.SubAccountSeller` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.SubAccountSeller` → `NVARCHAR` | `case when b.SubAccountSeller is not null then b.SubAccountSeller else case when tt.TrxType = 4 then @pcSubAccount else bsubacc.SubAccCode end end` |
| IC-006145 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementBIS4:1068](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1068) | EXPRESSION / P07 / CASE_BRANCH | `b.SubAccountSeller` : `VARCHAR(500)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `b.SubAccountSeller` → `NVARCHAR` | `CASE WHEN b.SubAccountSeller IS NOT NULL THEN b.SubAccountSeller ELSE CASE WHEN tt.TrxType = 4 THEN @pcSubAccount ELSE bsubacc.SubAccCode END END` |
