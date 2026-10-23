# Retail UAT.dbo.TRSOnCallDepositListRek_TMP — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 9 lokasi source, 3 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-acctno"></a>
## AcctNo

**Tipe dasar:** `VARCHAR(14)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSOnCallDepositListRek_TMP.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004223 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertDepositAccNo:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertDepositAccNo.sql#L25) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsoncalldepositlistrek_tmp.AcctNo` : `VARCHAR(14)` | `convert(bigint, AccountNumber)` : `bigint` | `convert(bigint, AccountNumber)` → `VARCHAR(14)` | `retail.dbo.trsoncalldepositlistrek_tmp.AcctNo <- convert(bigint, AccountNumber)` |

<a id="kolom-cifno"></a>
## CIFNo

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSOnCallDepositListRek_TMP.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004224 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertDepositAccNo:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertDepositAccNo.sql#L25) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsoncalldepositlistrek_tmp.CIFNo` : `VARCHAR(20)` | `convert(bigint, CIFNumber)` : `bigint` | `convert(bigint, CIFNumber)` → `VARCHAR(20)` | `retail.dbo.trsoncalldepositlistrek_tmp.CIFNo <- convert(bigint, CIFNumber)` |

<a id="kolom-guid"></a>
## Guid

**Tipe dasar:** `UNIQUEIDENTIFIER`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSOnCallDepositListRek_TMP.sql#L2). **Existing leading index:** TRSOnCallDepositListRek_TMP_IUNC1.

**Frekuensi source:** 7 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004225 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertDepositAccNo:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertDepositAccNo.sql#L25) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsoncalldepositlistrek_tmp.Guid` : `UNIQUEIDENTIFIER` | `@vGuid` : `varchar(40)` | `@vGuid` → `UNIQUEIDENTIFIER` | `retail.dbo.trsoncalldepositlistrek_tmp.Guid <- @vGuid` |
| IC-004939 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSOnCallDepositDeleteListRek_TMP:10](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOnCallDepositDeleteListRek_TMP.sql#L10) | WHERE / P13 / COMPARISON | `Guid` : `UNIQUEIDENTIFIER` | `@vGuid` : `varchar(40)` | `@vGuid` → `UNIQUEIDENTIFIER` | `Guid = @vGuid` |
| IC-012252 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSrcDepositAccNumber:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSrcDepositAccNumber.sql#L29) | WHERE / P13 / COMPARISON | `[Guid]` : `UNIQUEIDENTIFIER` | `@cCriteria` : `varchar(40)` | `@cCriteria` → `UNIQUEIDENTIFIER` | `[Guid] = @cCriteria` |
| IC-012253 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSrcDepositAccNumber:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSrcDepositAccNumber.sql#L36) | WHERE / P13 / COMPARISON | `[Guid]` : `UNIQUEIDENTIFIER` | `@cCriteria` : `varchar(40)` | `@cCriteria` → `UNIQUEIDENTIFIER` | `[Guid] = @cCriteria` |
| IC-012254 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSrcDepositAccNumber:42](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSrcDepositAccNumber.sql#L42) | WHERE / P13 / COMPARISON | `[Guid]` : `UNIQUEIDENTIFIER` | `@cCriteria` : `varchar(40)` | `@cCriteria` → `UNIQUEIDENTIFIER` | `[Guid] = @cCriteria` |
| IC-012255 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSrcDepositAccNumber:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSrcDepositAccNumber.sql#L49) | WHERE / P13 / COMPARISON | `[Guid]` : `UNIQUEIDENTIFIER` | `@cCriteria` : `varchar(40)` | `@cCriteria` → `UNIQUEIDENTIFIER` | `[Guid] = @cCriteria` |
| IC-012256 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSrcDepositAccNumber:55](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSrcDepositAccNumber.sql#L55) | WHERE / P13 / COMPARISON | `[Guid]` : `UNIQUEIDENTIFIER` | `@cCriteria` : `varchar(40)` | `@cCriteria` → `UNIQUEIDENTIFIER` | `[Guid] = @cCriteria` |
