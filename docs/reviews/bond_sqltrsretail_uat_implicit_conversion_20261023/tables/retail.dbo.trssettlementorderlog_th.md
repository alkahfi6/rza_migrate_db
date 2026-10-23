# Retail UAT.dbo.TRSSettlementOrderLog_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 14 lokasi source, 6 procedure, 4 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-instruction"></a>
## Instruction

**Tipe dasar:** `CHAR(2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementOrderLog_TH.sql#L15). **Existing leading index:** TRSSettlementOrderLog_TH_NN4.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003361 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSAutoSettlementOnlineORI:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAutoSettlementOnlineORI.sql#L137) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementorderlog_th.Instruction` : `CHAR(2)` | `'OM'` : `varchar` | `'OM'` → `CHAR(2)` | `retail.dbo.trssettlementorderlog_th.Instruction <- 'OM'` |

<a id="kolom-processusernik"></a>
## ProcessUserNik

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementOrderLog_TH.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003362 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSAutoSettlementOnlineORI:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAutoSettlementOnlineORI.sql#L137) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementorderlog_th.ProcessUserNik` : `INT` | `'77777'` : `varchar` | `'77777'` → `INT` | `retail.dbo.trssettlementorderlog_th.ProcessUserNik <- '77777'` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementOrderLog_TH.sql#L13). **Existing leading index:** TRSSettlementOrderLog_NN1.

**Frekuensi source:** 11 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006310 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementOrder:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementOrder.sql#L113) | CASE / P12 / COMPARISON | `solth.[Status]` : `CHAR(1)` | `'P'` : `varchar` | `solth.[Status]` → `VARCHAR` | `solth.[Status] = 'P'` |
| IC-006311 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementOrder:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementOrder.sql#L114) | CASE / P12 / COMPARISON | `solth.[Status]` : `CHAR(1)` | `'S'` : `varchar` | `solth.[Status]` → `VARCHAR` | `solth.[Status] = 'S'` |
| IC-006312 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementOrder:152](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementOrder.sql#L152) | JOIN_ON / P12 / COMPARISON | `solth.Status` : `CHAR(1)` | `'F'` : `varchar` | `solth.Status` → `VARCHAR` | `solth.Status <> 'F'` |
| IC-006765 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBuyBack:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBuyBack.sql#L108) | JOIN_ON / P12 / COMPARISON | `solth.Status` : `CHAR(1)` | `'S'` : `varchar` | `solth.Status` → `VARCHAR` | `solth.Status = 'S'` |
| IC-007532 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSettlementOrder:209](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementOrder.sql#L209) | JOIN_ON / P12 / COMPARISON | `solth.Status` : `CHAR(1)` | `'F'` : `varchar` | `solth.Status` → `VARCHAR` | `solth.Status <> 'F'` |
| IC-011516 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSettledOrder:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettledOrder.sql#L28) | CASE / P12 / COMPARISON | `solth.Status` : `CHAR(1)` | `'P'` : `varchar` | `solth.Status` → `VARCHAR` | `solth.Status = 'P'` |
| IC-011517 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSettledOrder:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettledOrder.sql#L29) | CASE / P12 / COMPARISON | `solth.Status` : `CHAR(1)` | `'S'` : `varchar` | `solth.Status` → `VARCHAR` | `solth.Status = 'S'` |
| IC-011518 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSettledOrder:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettledOrder.sql#L30) | CASE / P12 / COMPARISON | `solth.Status` : `CHAR(1)` | `'F'` : `varchar` | `solth.Status` → `VARCHAR` | `solth.Status = 'F'` |
| IC-012543 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdatePostingResult:100](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L100) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'P'` : `varchar` | `Status` → `VARCHAR` | `Status = 'P'` |
| IC-003363 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSAutoSettlementOnlineORI:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAutoSettlementOnlineORI.sql#L137) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementorderlog_th.Status` : `CHAR(1)` | `'S'` : `varchar` | `'S'` → `CHAR(1)` | `retail.dbo.trssettlementorderlog_th.Status <- 'S'` |
| IC-007543 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementOrder:341](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementOrder.sql#L341) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementorderlog_th.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trssettlementorderlog_th.Status <- 'P'` |

<a id="kolom-uniqueid"></a>
## UniqueId

**Tipe dasar:** `UNIQUEIDENTIFIER`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementOrderLog_TH.sql#L12). **Existing leading index:** TRSSettlementOrderLog_TH_NN2.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-011519 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptSettledOrder:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettledOrder.sql#L63) | JOIN_ON / P13 / COMPARISON | `solth.UniqueId` : `UNIQUEIDENTIFIER` | `po.OBLREM3` : `VARCHAR(40)` | `po.OBLREM3` → `UNIQUEIDENTIFIER` | `solth.UniqueId = po.OBLREM3` |
