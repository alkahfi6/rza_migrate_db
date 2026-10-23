# Retail UAT.dbo.job_general_control_table_tm — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 8 lokasi source, 1 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-last-finish-date"></a>
## last_finish_date

**Tipe dasar:** `DATETIME2(7)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.job_general_control_table_tm.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000641 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.JobCachingAccountSegment:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.JobCachingAccountSegment.sql#L86) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[last_finish_date]` : `DATETIME2(7)` | `getdate()` : `datetime` | `getdate()` → `DATETIME2(7)` | `[last_finish_date] = getdate()` |
| IC-000643 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.JobCachingAccountSegment:95](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.JobCachingAccountSegment.sql#L95) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[last_finish_date]` : `DATETIME2(7)` | `getdate()` : `datetime` | `getdate()` → `DATETIME2(7)` | `[last_finish_date] = getdate()` |

<a id="kolom-last-start-date"></a>
## last_start_date

**Tipe dasar:** `DATETIME2(7)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.job_general_control_table_tm.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000635 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.JobCachingAccountSegment:9](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.JobCachingAccountSegment.sql#L9) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.job_general_control_table_tm.last_start_date` : `DATETIME2(7)` | `getdate()` : `datetime` | `getdate()` → `DATETIME2(7)` | `retail.dbo.job_general_control_table_tm.last_start_date <- getdate()` |
| IC-000638 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.JobCachingAccountSegment:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.JobCachingAccountSegment.sql#L26) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[last_start_date]` : `DATETIME2(7)` | `getdate()` : `datetime` | `getdate()` → `DATETIME2(7)` | `[last_start_date] = getdate()` |

<a id="kolom-status"></a>
## status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.job_general_control_table_tm.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000636 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.JobCachingAccountSegment:9](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.JobCachingAccountSegment.sql#L9) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.job_general_control_table_tm.status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.job_general_control_table_tm.status <- 'P'` |
| IC-000637 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.JobCachingAccountSegment:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.JobCachingAccountSegment.sql#L25) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[status]` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `[status] = 'P'` |
| IC-000640 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.JobCachingAccountSegment:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.JobCachingAccountSegment.sql#L85) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[status]` : `CHAR(1)` | `'S'` : `varchar` | `'S'` → `CHAR(1)` | `[status] = 'S'` |
| IC-000642 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.JobCachingAccountSegment:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.JobCachingAccountSegment.sql#L94) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[status]` : `CHAR(1)` | `'F'` : `varchar` | `'F'` → `CHAR(1)` | `[status] = 'F'` |
