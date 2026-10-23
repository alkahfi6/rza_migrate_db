# Retail UAT.dbo.TRSOnCallDepositListRek_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 10 lokasi source, 1 procedure, 5 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-isprinted"></a>
## IsPrinted

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSOnCallDepositListRek_TT.sql#L27). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004941 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSOnCallDepositPopulate:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOnCallDepositPopulate.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsoncalldepositlistrek_tt.IsPrinted` : `BIT` | `'1'` : `varchar` | `'1'` → `BIT` | `retail.dbo.trsoncalldepositlistrek_tt.IsPrinted <- '1'` |
| IC-004946 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSOnCallDepositPopulate:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOnCallDepositPopulate.sql#L86) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsoncalldepositlistrek_tt.IsPrinted` : `BIT` | `'1'` : `varchar` | `'1'` → `BIT` | `retail.dbo.trsoncalldepositlistrek_tt.IsPrinted <- '1'` |

<a id="kolom-isreprinted"></a>
## IsReprinted

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSOnCallDepositListRek_TT.sql#L28). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004942 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSOnCallDepositPopulate:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOnCallDepositPopulate.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsoncalldepositlistrek_tt.IsReprinted` : `BIT` | `'1'` : `varchar` | `'1'` → `BIT` | `retail.dbo.trsoncalldepositlistrek_tt.IsReprinted <- '1'` |
| IC-004947 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSOnCallDepositPopulate:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOnCallDepositPopulate.sql#L86) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsoncalldepositlistrek_tt.IsReprinted` : `BIT` | `'0'` : `varchar` | `'0'` → `BIT` | `retail.dbo.trsoncalldepositlistrek_tt.IsReprinted <- '0'` |

<a id="kolom-matdate"></a>
## MatDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSOnCallDepositListRek_TT.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004943 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSOnCallDepositPopulate:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOnCallDepositPopulate.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsoncalldepositlistrek_tt.MatDate` : `DATETIME` | `@MatDate` : `varchar(8)` | `@MatDate` → `DATETIME` | `retail.dbo.trsoncalldepositlistrek_tt.MatDate <- @MatDate` |
| IC-004948 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSOnCallDepositPopulate:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOnCallDepositPopulate.sql#L86) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsoncalldepositlistrek_tt.MatDate` : `DATETIME` | `@MatDate` : `varchar(8)` | `@MatDate` → `DATETIME` | `retail.dbo.trsoncalldepositlistrek_tt.MatDate <- @MatDate` |

<a id="kolom-trxdate"></a>
## TrxDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSOnCallDepositListRek_TT.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004944 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSOnCallDepositPopulate:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOnCallDepositPopulate.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsoncalldepositlistrek_tt.TrxDate` : `DATETIME` | `@TrxDate` : `varchar(8)` | `@TrxDate` → `DATETIME` | `retail.dbo.trsoncalldepositlistrek_tt.TrxDate <- @TrxDate` |
| IC-004949 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSOnCallDepositPopulate:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOnCallDepositPopulate.sql#L86) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsoncalldepositlistrek_tt.TrxDate` : `DATETIME` | `@TrxDate` : `varchar(8)` | `@TrxDate` → `DATETIME` | `retail.dbo.trsoncalldepositlistrek_tt.TrxDate <- @TrxDate` |

<a id="kolom-valuedate"></a>
## ValueDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSOnCallDepositListRek_TT.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004945 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSOnCallDepositPopulate:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOnCallDepositPopulate.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsoncalldepositlistrek_tt.ValueDate` : `DATETIME` | `@ValueDate` : `varchar(8)` | `@ValueDate` → `DATETIME` | `retail.dbo.trsoncalldepositlistrek_tt.ValueDate <- @ValueDate` |
| IC-004950 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSOnCallDepositPopulate:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOnCallDepositPopulate.sql#L86) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsoncalldepositlistrek_tt.ValueDate` : `DATETIME` | `@ValueDate` : `varchar(8)` | `@ValueDate` → `DATETIME` | `retail.dbo.trsoncalldepositlistrek_tt.ValueDate <- @ValueDate` |
