# Retail UAT.dbo.TRSPajakkLog_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 7 lokasi source, 3 procedure, 4 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-orderid"></a>
## OrderId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSPajakkLog_TH.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000099 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.BuybackLogInsert:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BuybackLogInsert.sql#L125) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trspajakklog_th.OrderId` : `BIGINT` | `t.DealIdSekunder` : `VARCHAR(100)` | `t.DealIdSekunder` → `BIGINT` | `retail.dbo.trspajakklog_th.OrderId <- t.DealIdSekunder` |

<a id="kolom-processusernik"></a>
## ProcessUserNik

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSPajakkLog_TH.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000100 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.BuybackLogInsert:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BuybackLogInsert.sql#L125) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trspajakklog_th.ProcessUserNik` : `INT` | `t.SupervisorId` : `VARCHAR(10)` | `t.SupervisorId` → `INT` | `retail.dbo.trspajakklog_th.ProcessUserNik <- t.SupervisorId` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSPajakkLog_TH.sql#L12). **Existing leading index:** TRSPajakkLog_TH_NN1.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012550 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdatePostingResult:142](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L142) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'P'` : `varchar` | `Status` → `VARCHAR` | `Status = 'P'` |
| IC-000101 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.BuybackLogInsert:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BuybackLogInsert.sql#L125) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trspajakklog_th.Status` : `CHAR(1)` | `case when t.Status = 3 then 'S' when t.Status = 2 then 'F' end` : `VARCHAR (combined result; facets unknown)` | `case when t.Status = 3 then 'S' when t.Status = 2 then 'F' end` → `CHAR(1)` | `retail.dbo.trspajakklog_th.Status <- case when t.Status = 3 then 'S' when t.Status = 2 then 'F' end` |
| IC-007450 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementBuyBack:321](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBuyBack.sql#L321) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trspajakklog_th.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trspajakklog_th.Status <- 'P'` |

<a id="kolom-uniqueid"></a>
## UniqueId

**Tipe dasar:** `UNIQUEIDENTIFIER`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSPajakkLog_TH.sql#L11). **Existing leading index:** TRSPajakkLog_TH_NN2.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000102 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.BuybackLogInsert:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BuybackLogInsert.sql#L125) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trspajakklog_th.UniqueId` : `UNIQUEIDENTIFIER` | `t.Remark3` : `VARCHAR(40)` | `t.Remark3` → `UNIQUEIDENTIFIER` | `retail.dbo.trspajakklog_th.UniqueId <- t.Remark3` |
| IC-000105 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.BuybackLogInsert:135](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BuybackLogInsert.sql#L135) | JOIN_ON / P13 / COMPARISON | `pl.UniqueId` : `UNIQUEIDENTIFIER` | `t.Remark3` : `VARCHAR(40)` | `t.Remark3` → `UNIQUEIDENTIFIER` | `t.Remark3 = pl.UniqueId` |
