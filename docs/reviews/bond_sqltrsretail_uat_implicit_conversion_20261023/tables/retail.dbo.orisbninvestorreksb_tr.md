# Retail UAT.dbo.ORISBNInvestorRekSB_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 37 lokasi source, 10 procedure, 3 kolom. **A:** 6 lokasi pada 4 procedure. **B:** 10 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-idreksb"></a>
## IdRekSB

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISBNInvestorRekSB_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 21 lokasi / 9 SP. **A:** 6 lokasi / 4 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000959 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.ORIOnlineReportDailyTransaction:364](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORIOnlineReportDailyTransaction.sql#L364) | JOIN_ON / P01 / COMPARISON | `sb.IdRekSB` : `VARCHAR(20)` | `tc.IdRekSB` : `INT` | `sb.IdRekSB` → `INT` | `tc.IdRekSB = sb.IdRekSB` |
| IC-002898 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.SBNSchedProcessRekSB:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekSB.sql#L98) | WHERE / P01 / COMPARISON | `[IdRekSB]` : `VARCHAR(20)` | `@nIdRekSB` : `int` | `[IdRekSB]` → `INT` | `[IdRekSB] = @nIdRekSB` |
| IC-002901 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.SBNSchedProcessRekSB:157](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekSB.sql#L157) | JOIN_ON / P01 / COMPARISON | `rek.IdRekSB` : `VARCHAR(20)` | `td.IdRekSB` : `int` | `rek.IdRekSB` → `INT` | `rek.IdRekSB = td.IdRekSB` |
| IC-002903 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.SBNSchedProcessRekSB:198](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekSB.sql#L198) | WHERE / P01 / COMPARISON | `[IdRekSB]` : `VARCHAR(20)` | `@nIdRekSB` : `int` | `[IdRekSB]` → `INT` | `[IdRekSB] = @nIdRekSB` |
| IC-008371 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:619](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L619) | JOIN_ON / P01 / COMPARISON | `sb.IdRekSB` : `VARCHAR(20)` | `scd.IdRekSB` : `INT` | `sb.IdRekSB` → `INT` | `scd.IdRekSB = sb.IdRekSB` |
| IC-010343 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptOrderORIDetail:836](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOrderORIDetail.sql#L836) | JOIN_ON / P01 / COMPARISON | `sb.IdRekSB` : `VARCHAR(20)` | `tc.IdRekSB` : `INT` | `sb.IdRekSB` → `INT` | `tc.IdRekSB = sb.IdRekSB` |
| IC-001001 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.ORISBNIPopulateRegisterInvestor:254](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateRegisterInvestor.sql#L254) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `rd.IdRekSB` : `VARCHAR(20)` | `IdRekSB` : `int` | `rd.IdRekSB` → `int` | `IdRekSB = rd.IdRekSB` |
| IC-002814 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPopulateORIOrder:361](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateORIOrder.sql#L361) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `rek.IdRekSB` : `VARCHAR(20)` | `tmp.[IdRekSB]` : `int` | `rek.IdRekSB` → `int` | `tmp.[IdRekSB] = rek.IdRekSB` |
| IC-002815 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPopulateORIOrder:387](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateORIOrder.sql#L387) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `rk.IdRekSB` : `VARCHAR(20)` | `IdRekSB` : `int` | `rk.IdRekSB` → `int` | `IdRekSB = rk.IdRekSB` |
| IC-002817 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPopulateORIOrder:397](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateORIOrder.sql#L397) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `rk.IdRekSB` : `VARCHAR(20)` | `IdRekSB` : `int` | `rk.IdRekSB` → `int` | `IdRekSB = rk.IdRekSB` |
| IC-002824 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPopulateRegInvestor:172](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestor.sql#L172) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ISNULL(rksb.IdRekSB,0)` : `VARCHAR(20)` | `@tmp_Data.IdRekSB` : `int` | `ISNULL(rksb.IdRekSB,0)` → `int` | `@tmp_Data.IdRekSB <- ISNULL(rksb.IdRekSB,0)` |
| IC-002836 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPopulateRegInvestorForOrder:90](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestorForOrder.sql#L90) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ISNULL(rd.IdRekSB,0)` : `VARCHAR(20)` | `IdRekSB` : `int` | `ISNULL(rd.IdRekSB,0)` → `int` | `IdRekSB = ISNULL(rd.IdRekSB,0)` |
| IC-002839 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPopulateRegInvestorForOrder:112](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestorForOrder.sql#L112) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ISNULL(rd.IdRekSB,0)` : `VARCHAR(20)` | `IdRekSB` : `int` | `ISNULL(rd.IdRekSB,0)` → `int` | `IdRekSB = ISNULL(rd.IdRekSB,0)` |
| IC-002843 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPopulateRegInvestorForOrder:128](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestorForOrder.sql#L128) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ISNULL(rd.IdRekSB,0)` : `VARCHAR(20)` | `IdRekSB` : `int` | `ISNULL(rd.IdRekSB,0)` → `int` | `IdRekSB = ISNULL(rd.IdRekSB,0)` |
| IC-004796 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSORIOnlineNewTransaction:167](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlineNewTransaction.sql#L167) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `rk.[IdRekSB]` : `VARCHAR(20)` | `tmp.[IdRekSB]` : `int` | `rk.[IdRekSB]` → `int` | `tmp.[IdRekSB] = rk.[IdRekSB]` |
| IC-002827 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SBNSchedPopulateRegInvestor:199](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestor.sql#L199) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `rksb.IdRekSB` : `VARCHAR(20)` | `0` : `int` | `0` → `VARCHAR(20)` | `ISNULL(rksb.IdRekSB,0)` |
| IC-002837 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SBNSchedPopulateRegInvestorForOrder:90](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestorForOrder.sql#L90) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `rd.IdRekSB` : `VARCHAR(20)` | `0` : `int` | `0` → `VARCHAR(20)` | `ISNULL(rd.IdRekSB,0)` |
| IC-002840 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SBNSchedPopulateRegInvestorForOrder:112](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestorForOrder.sql#L112) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `rd.IdRekSB` : `VARCHAR(20)` | `0` : `int` | `0` → `VARCHAR(20)` | `ISNULL(rd.IdRekSB,0)` |
| IC-002844 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SBNSchedPopulateRegInvestorForOrder:128](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestorForOrder.sql#L128) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `rd.IdRekSB` : `VARCHAR(20)` | `0` : `int` | `0` → `VARCHAR(20)` | `ISNULL(rd.IdRekSB,0)` |
| IC-002899 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessRekSB:102](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekSB.sql#L102) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.orisbninvestorreksb_tr.IdRekSB` : `VARCHAR(20)` | `[IdRekSB]` : `int` | `[IdRekSB]` → `VARCHAR(20)` | `retail.dbo.orisbninvestorreksb_tr.IdRekSB <- [IdRekSB]` |
| IC-002900 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessRekSB:142](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekSB.sql#L142) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `rek.[IdRekSB]` : `VARCHAR(20)` | `td.[IdRekSB]` : `int` | `td.[IdRekSB]` → `VARCHAR(20)` | `rek.[IdRekSB] = td.[IdRekSB]` |

<a id="kolom-cifid"></a>
## CIFId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISBNInvestorRekSB_TR.sql#L2). **Existing leading index:** ORISBNInvestorRekSB_TR_CU1.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002897 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedProcessRekSB:97](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekSB.sql#L97) | WHERE / P11 / COMPARISON | `[CIFId]` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `[CIFId] = @nCIFId` |
| IC-002902 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedProcessRekSB:197](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekSB.sql#L197) | WHERE / P11 / COMPARISON | `[CIFId]` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `[CIFId] = @nCIFId` |
| IC-002955 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedPushMailNotifRegInvestor:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifRegInvestor.sql#L108) | WHERE / P11 / COMPARISON | `CIFId` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `CIFId = @nCIFId` |

<a id="kolom-idsubregistry"></a>
## IdSubregistry

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISBNInvestorRekSB_TR.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 13 lokasi / 8 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000998 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.ORISBNIPopulateRegisterInvestor:218](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateRegisterInvestor.sql#L218) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `rs.IdSubregistry` : `INT` | `iv.IdSubRegistry` : `varchar(10)` | `rs.IdSubregistry` → `varchar(10)` | `iv.IdSubRegistry = rs.IdSubregistry` |
| IC-000960 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.ORIOnlineReportDailyTransaction:366](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORIOnlineReportDailyTransaction.sql#L366) | JOIN_ON / P01 / COMPARISON | `sb.IdSubregistry` : `INT` | `cusd.IdSubReg` : `VARCHAR(10)` | `cusd.IdSubReg` → `INT` | `sb.IdSubregistry = cusd.IdSubReg` |
| IC-002816 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedPopulateORIOrder:391](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateORIOrder.sql#L391) | WHERE / P04 / COMPARISON | `rk.IdSubregistry` : `INT` | `'384'` : `varchar` | `'384'` → `INT` | `rk.IdSubregistry = '384'` |
| IC-002820 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedPopulateORIOrder:411](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateORIOrder.sql#L411) | JOIN_ON / P01 / COMPARISON | `rk.IdSubregistry` : `INT` | `map.IdSubReg` : `VARCHAR(15)` | `map.IdSubReg` → `INT` | `rk.IdSubregistry = map.IdSubReg` |
| IC-002832 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedPopulateRegInvestor:234](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestor.sql#L234) | JOIN_ON / P01 / COMPARISON | `rksb.IdSubregistry` : `INT` | `map.IdSubReg` : `VARCHAR(15)` | `map.IdSubReg` → `INT` | `map.IdSubReg = rksb.IdSubregistry` |
| IC-002838 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SBNSchedPopulateRegInvestorForOrder:92](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestorForOrder.sql#L92) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `rd.IdSubregistry` : `INT` | `'384'` : `varchar` | `'384'` → `INT` | `ISNULL(rd.IdSubregistry,'384')` |
| IC-002841 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SBNSchedPopulateRegInvestorForOrder:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestorForOrder.sql#L114) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `rd.IdSubregistry` : `INT` | `'384'` : `varchar` | `'384'` → `INT` | `ISNULL(rd.IdSubregistry,'384')` |
| IC-002842 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedPopulateRegInvestorForOrder:122](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestorForOrder.sql#L122) | JOIN_ON / P04 / COMPARISON | `rd.IdSubregistry` : `INT` | `'384'` : `varchar` | `'384'` → `INT` | `rd.IdSubregistry = '384'` |
| IC-002845 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SBNSchedPopulateRegInvestorForOrder:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestorForOrder.sql#L130) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `rd.IdSubregistry` : `INT` | `'1335'` : `varchar` | `'1335'` → `INT` | `ISNULL(rd.IdSubregistry,'1335')` |
| IC-002846 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedPopulateRegInvestorForOrder:138](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestorForOrder.sql#L138) | JOIN_ON / P04 / COMPARISON | `rd.IdSubregistry` : `INT` | `'1335'` : `varchar` | `'1335'` → `INT` | `rd.IdSubregistry = '1335'` |
| IC-004800 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSORIOnlineNewTransaction:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlineNewTransaction.sql#L180) | JOIN_ON / P01 / COMPARISON | `rk.IdSubregistry` : `INT` | `map.IdSubReg` : `VARCHAR(15)` | `map.IdSubReg` → `INT` | `rk.IdSubregistry = map.IdSubReg` |
| IC-008372 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportReportSettlementInstruction:621](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L621) | JOIN_ON / P01 / COMPARISON | `sb.IdSubregistry` : `INT` | `tbctm.IdSubReg` : `VARCHAR(10)` | `tbctm.IdSubReg` → `INT` | `sb.IdSubregistry = tbctm.IdSubReg` |
| IC-010344 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptOrderORIDetail:838](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOrderORIDetail.sql#L838) | JOIN_ON / P01 / COMPARISON | `sb.IdSubregistry` : `INT` | `cusd.IdSubReg` : `VARCHAR(10)` | `cusd.IdSubReg` → `INT` | `sb.IdSubregistry = cusd.IdSubReg` |
