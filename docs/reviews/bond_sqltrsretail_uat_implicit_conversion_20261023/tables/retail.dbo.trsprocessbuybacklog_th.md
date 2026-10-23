# Retail UAT.dbo.TRSProcessBuyBackLog_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 15 lokasi source, 5 procedure, 5 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-dealid"></a>
## DealId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSProcessBuyBackLog_TH.sql#L3). **Existing leading index:** TRSProcessBuyBackLog_TH_NN3.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000089 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.BuybackLogInsert:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BuybackLogInsert.sql#L86) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsprocessbuybacklog_th.DealId` : `BIGINT` | `t.DealIdSekunder` : `VARCHAR(100)` | `t.DealIdSekunder` → `BIGINT` | `retail.dbo.trsprocessbuybacklog_th.DealId <- t.DealIdSekunder` |
| IC-000097 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.BuybackLogInsert:119](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BuybackLogInsert.sql#L119) | JOIN_ON / P01 / COMPARISON | `pb.DealId` : `BIGINT` | `t.DealIdSekunder` : `VARCHAR(100)` | `t.DealIdSekunder` → `BIGINT` | `t.DealIdSekunder = pb.DealId` |

<a id="kolom-orderid"></a>
## OrderId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSProcessBuyBackLog_TH.sql#L2). **Existing leading index:** TRSProcessBuyBackLog_TH_CN1.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000090 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.BuybackLogInsert:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BuybackLogInsert.sql#L86) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsprocessbuybacklog_th.OrderId` : `BIGINT` | `t.DealIdSekunder` : `VARCHAR(100)` | `t.DealIdSekunder` → `BIGINT` | `retail.dbo.trsprocessbuybacklog_th.OrderId <- t.DealIdSekunder` |
| IC-000096 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.BuybackLogInsert:118](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BuybackLogInsert.sql#L118) | JOIN_ON / P11 / COMPARISON | `pb.OrderId` : `BIGINT` | `t.BuyBackID` : `INT` | `t.BuyBackID` → `BIGINT` | `t.BuyBackID = pb.OrderId` |

<a id="kolom-processusernik"></a>
## ProcessUserNik

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSProcessBuyBackLog_TH.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000091 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.BuybackLogInsert:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BuybackLogInsert.sql#L86) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsprocessbuybacklog_th.ProcessUserNik` : `INT` | `t.SupervisorId` : `VARCHAR(10)` | `t.SupervisorId` → `INT` | `retail.dbo.trsprocessbuybacklog_th.ProcessUserNik <- t.SupervisorId` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSProcessBuyBackLog_TH.sql#L13). **Existing leading index:** TRSProcessBuyBackLog_TH_NN1.

**Frekuensi source:** 8 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005332 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateBuyBack:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateBuyBack.sql#L18) | CASE / P12 / COMPARISON | `pbblth.Status` : `CHAR(1)` | `'S'` : `varchar` | `pbblth.Status` → `VARCHAR` | `pbblth.Status = 'S'` |
| IC-005333 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateBuyBack:19](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateBuyBack.sql#L19) | CASE / P12 / COMPARISON | `pbblth.Status` : `CHAR(1)` | `'P'` : `varchar` | `pbblth.Status` → `VARCHAR` | `pbblth.Status = 'P'` |
| IC-005334 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateBuyBack:68](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateBuyBack.sql#L68) | JOIN_ON / P12 / COMPARISON | `pbblth.Status` : `CHAR(1)` | `'F'` : `varchar` | `pbblth.Status` → `VARCHAR` | `pbblth.Status <> 'F'` |
| IC-006764 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBuyBack:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBuyBack.sql#L98) | JOIN_ON / P12 / COMPARISON | `sbblth.Status` : `CHAR(1)` | `'F'` : `varchar` | `sbblth.Status` → `VARCHAR` | `sbblth.Status <> 'F'` |
| IC-007446 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSettlementBuyBack:256](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBuyBack.sql#L256) | JOIN_ON / P12 / COMPARISON | `sbblth.Status` : `CHAR(1)` | `'S'` : `varchar` | `sbblth.Status` → `VARCHAR` | `sbblth.Status = 'S'` |
| IC-012549 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdatePostingResult:136](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L136) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'P'` : `varchar` | `Status` → `VARCHAR` | `Status = 'P'` |
| IC-000092 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.BuybackLogInsert:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BuybackLogInsert.sql#L86) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsprocessbuybacklog_th.Status` : `CHAR(1)` | `case when t.Status = 3 then 'S' when t.Status = 2 then 'F' end` : `VARCHAR (combined result; facets unknown)` | `case when t.Status = 3 then 'S' when t.Status = 2 then 'F' end` → `CHAR(1)` | `retail.dbo.trsprocessbuybacklog_th.Status <- case when t.Status = 3 then 'S' when t.Status = 2 then 'F' end` |
| IC-006766 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessBuyBack:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBuyBack.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsprocessbuybacklog_th.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trsprocessbuybacklog_th.Status <- 'P'` |

<a id="kolom-uniqueid"></a>
## UniqueId

**Tipe dasar:** `UNIQUEIDENTIFIER`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSProcessBuyBackLog_TH.sql#L12). **Existing leading index:** TRSProcessBuyBackLog_TH_NN2.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000093 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.BuybackLogInsert:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BuybackLogInsert.sql#L86) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsprocessbuybacklog_th.UniqueId` : `UNIQUEIDENTIFIER` | `t.Remark3` : `VARCHAR(40)` | `t.Remark3` → `UNIQUEIDENTIFIER` | `retail.dbo.trsprocessbuybacklog_th.UniqueId <- t.Remark3` |
| IC-000098 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.BuybackLogInsert:120](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BuybackLogInsert.sql#L120) | JOIN_ON / P13 / COMPARISON | `pb.UniqueId` : `UNIQUEIDENTIFIER` | `t.Remark3` : `VARCHAR(40)` | `t.Remark3` → `UNIQUEIDENTIFIER` | `t.Remark3 = pb.UniqueId` |
