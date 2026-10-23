# Retail UAT.dbo.dump_maturity_daily_tt — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 15 lokasi source, 3 procedure, 8 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 2 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-account-no"></a>
## account_no

**Tipe dasar:** `CHAR(50)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.dump_maturity_daily_tt.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002236 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSMaturityDumpOBL058:218](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityDumpOBL058.sql#L218) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.dump_maturity_daily_tt.account_no` : `CHAR(50)` | `LEFT(ISNULL(AccountNo, ''), 100)` : `VARCHAR (string function; length unknown)` | `LEFT(ISNULL(AccountNo, ''), 100)` → `CHAR(50)` | `retail.dbo.dump_maturity_daily_tt.account_no <- LEFT(ISNULL(AccountNo, ''), 100)` |

<a id="kolom-capital-gain"></a>
## capital_gain

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.dump_maturity_daily_tt.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002059 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSGenerateDumpOBL058_New:174](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGenerateDumpOBL058_New.sql#L174) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `capital_gain` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(capital_gain, 0)` |
| IC-002237 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSMaturityDumpOBL058:218](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityDumpOBL058.sql#L218) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.dump_maturity_daily_tt.capital_gain` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.dump_maturity_daily_tt.capital_gain <- 0` |

<a id="kolom-capital-gain-loss-after-tax"></a>
## capital_gain_loss_after_tax

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.dump_maturity_daily_tt.sql#L25). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002061 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSGenerateDumpOBL058_New:176](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGenerateDumpOBL058_New.sql#L176) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `capital_gain_loss_after_tax` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(capital_gain_loss_after_tax, 0)` |
| IC-002238 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSMaturityDumpOBL058:218](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityDumpOBL058.sql#L218) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.dump_maturity_daily_tt.capital_gain_loss_after_tax` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.dump_maturity_daily_tt.capital_gain_loss_after_tax <- 0` |

<a id="kolom-capital-gain-loss-after-tax-spt"></a>
## capital_gain_loss_after_tax_spt

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.dump_maturity_daily_tt.sql#L30). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002063 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSGenerateDumpOBL058_New:181](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGenerateDumpOBL058_New.sql#L181) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `capital_gain_loss_after_tax_spt` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(capital_gain_loss_after_tax_spt, 0)` |
| IC-002239 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSMaturityDumpOBL058:218](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityDumpOBL058.sql#L218) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.dump_maturity_daily_tt.capital_gain_loss_after_tax_spt` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.dump_maturity_daily_tt.capital_gain_loss_after_tax_spt <- 0` |

<a id="kolom-capital-gain-spt"></a>
## capital_gain_spt

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.dump_maturity_daily_tt.sql#L29). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002062 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSGenerateDumpOBL058_New:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGenerateDumpOBL058_New.sql#L180) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `capital_gain_spt` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(capital_gain_spt, 0)` |
| IC-002240 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSMaturityDumpOBL058:218](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityDumpOBL058.sql#L218) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.dump_maturity_daily_tt.capital_gain_spt` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.dump_maturity_daily_tt.capital_gain_spt <- 0` |

<a id="kolom-inserted-date"></a>
## inserted_date

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.dump_maturity_daily_tt.sql#L31). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002064 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSGenerateDumpOBL058_New:184](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGenerateDumpOBL058_New.sql#L184) | WHERE / P13 / COMPARISON | `inserted_date` : `DATETIME` | `CONVERT(date, @dCurrentDate)` : `date` | `CONVERT(date, @dCurrentDate)` → `DATETIME` | `inserted_date >= CONVERT(date, @dCurrentDate)` |
| IC-002065 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSGenerateDumpOBL058_New:185](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGenerateDumpOBL058_New.sql#L185) | WHERE / P13 / COMPARISON | `inserted_date` : `DATETIME` | `DATEADD(day, 1, CONVERT(date, @dCurrentDate))` : `date` | `DATEADD(day, 1, CONVERT(date, @dCurrentDate))` → `DATETIME` | `inserted_date < DATEADD(day, 1, CONVERT(date, @dCurrentDate))` |

<a id="kolom-kitas-last-update-date"></a>
## kitas_last_update_date

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.dump_maturity_daily_tt.sql#L28). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002006 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.Rvmp_TRSGenerateDumpOBL058:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGenerateDumpOBL058.sql#L107) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `kitas_last_update_date` : `DATETIME` | `@tORIDetail.KITASLastUpdateDate` : `varchar(10)` | `kitas_last_update_date` → `varchar(10)` | `@tORIDetail.KITASLastUpdateDate <- kitas_last_update_date` |
| IC-002058 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.Rvmp_TRSGenerateDumpOBL058_New:120](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGenerateDumpOBL058_New.sql#L120) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `kitas_last_update_date` : `DATETIME` | `@tORIDetail.KITASLastUpdateDate` : `VARCHAR(10)` | `kitas_last_update_date` → `VARCHAR(10)` | `@tORIDetail.KITASLastUpdateDate <- kitas_last_update_date` |

<a id="kolom-tax-on-capital-gain"></a>
## tax_on_capital_gain

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.dump_maturity_daily_tt.sql#L24). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002060 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSGenerateDumpOBL058_New:175](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGenerateDumpOBL058_New.sql#L175) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tax_on_capital_gain` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tax_on_capital_gain, 0)` |
| IC-002241 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSMaturityDumpOBL058:218](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityDumpOBL058.sql#L218) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.dump_maturity_daily_tt.tax_on_capital_gain` : `MONEY` | `0` : `int` | `0` → `MONEY` | `retail.dbo.dump_maturity_daily_tt.tax_on_capital_gain <- 0` |
