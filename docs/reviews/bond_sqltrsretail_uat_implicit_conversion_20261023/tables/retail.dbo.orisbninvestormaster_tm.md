# Retail UAT.dbo.ORISBNInvestorMaster_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 7 lokasi source, 2 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cifid"></a>
## CIFId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISBNInvestorMaster_TM.sql#L2). **Existing leading index:** ORISBNInvestorMaster_TM_CU1.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002850 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedProcessMasterInvestor:167](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessMasterInvestor.sql#L167) | WHERE / P11 / COMPARISON | `[CIFId]` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `[CIFId] = @nCIFId` |
| IC-002855 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedProcessMasterInvestor:324](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessMasterInvestor.sql#L324) | WHERE / P11 / COMPARISON | `[CIFId]` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `[CIFId] = @nCIFId` |
| IC-002952 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedPushMailNotifRegInvestor:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifRegInvestor.sql#L94) | WHERE / P11 / COMPARISON | `CIFId` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `CIFId = @nCIFId` |

<a id="kolom-sbncreationdate"></a>
## SBNCreationDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISBNInvestorMaster_TM.sql#L24). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002851 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessMasterInvestor:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessMasterInvestor.sql#L169) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.orisbninvestormaster_tm.SBNCreationDate` : `DATETIME` | `[SBNCreationDate]` : `varchar(40)` | `[SBNCreationDate]` → `DATETIME` | `retail.dbo.orisbninvestormaster_tm.SBNCreationDate <- [SBNCreationDate]` |
| IC-002853 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessMasterInvestor:252](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessMasterInvestor.sql#L252) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mst.[SBNCreationDate]` : `DATETIME` | `td.[SBNCreationDate]` : `varchar(40)` | `td.[SBNCreationDate]` → `DATETIME` | `mst.[SBNCreationDate] = td.[SBNCreationDate]` |

<a id="kolom-sbnlastmodifdate"></a>
## SBNLastModifDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISBNInvestorMaster_TM.sql#L26). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002852 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessMasterInvestor:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessMasterInvestor.sql#L169) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.orisbninvestormaster_tm.SBNLastModifDate` : `DATETIME` | `[SBNLastModifDate]` : `varchar(40)` | `[SBNLastModifDate]` → `DATETIME` | `retail.dbo.orisbninvestormaster_tm.SBNLastModifDate <- [SBNLastModifDate]` |
| IC-002854 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessMasterInvestor:254](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessMasterInvestor.sql#L254) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mst.[SBNLastModifDate]` : `DATETIME` | `td.[SBNLastModifDate]` : `varchar(40)` | `td.[SBNLastModifDate]` → `DATETIME` | `mst.[SBNLastModifDate] = td.[SBNLastModifDate]` |
