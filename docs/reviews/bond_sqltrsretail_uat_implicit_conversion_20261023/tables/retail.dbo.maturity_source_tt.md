# Retail UAT.dbo.maturity_source_tt — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 57 lokasi source, 4 procedure, 20 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-acquisition-price"></a>
## acquisition_price

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_source_tt.sql#L32). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005570 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:310](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L310) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.acquisition_price` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(src.acquisition_price, 0)` |

<a id="kolom-bonus"></a>
## bonus

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_source_tt.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005568 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:307](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L307) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.bonus` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(src.bonus, 0)` |

<a id="kolom-capital-gain"></a>
## capital_gain

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_source_tt.sql#L26). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005579 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:322](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L322) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.capital_gain` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(src.capital_gain, 0)` |

<a id="kolom-capital-gain-loss-after-tax"></a>
## capital_gain_loss_after_tax

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_source_tt.sql#L28). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005581 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:324](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L324) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.capital_gain_loss_after_tax` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(src.capital_gain_loss_after_tax, 0)` |

<a id="kolom-capital-gain-loss-after-tax-spt"></a>
## capital_gain_loss_after_tax_spt

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_source_tt.sql#L30). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005584 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:327](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L327) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.capital_gain_loss_after_tax_spt` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(src.capital_gain_loss_after_tax_spt, 0)` |

<a id="kolom-capital-gain-spt"></a>
## capital_gain_spt

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_source_tt.sql#L29). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005583 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:326](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L326) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.capital_gain_spt` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(src.capital_gain_spt, 0)` |

<a id="kolom-coupon-after-tax"></a>
## coupon_after_tax

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_source_tt.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005573 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:314](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L314) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.coupon_after_tax` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(src.coupon_after_tax, 0)` |

<a id="kolom-gross-amount"></a>
## gross_amount

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_source_tt.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005571 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:312](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L312) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.gross_amount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(src.gross_amount, 0)` |

<a id="kolom-match-flag"></a>
## match_flag

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_source_tt.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 24 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002226 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.Rvmp_TRSMaturityCopyFromCustody:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityCopyFromCustody.sql#L162) | WHERE / P12 / IN_VALUE_TYPE | `ISNULL(cs.match_flag, '')` : `CHAR(1)` | `'U'` : `varchar` | `ISNULL(cs.match_flag, '')` → `VARCHAR` | `ISNULL(cs.match_flag, '') NOT IN ('U', 'R')` |
| IC-002229 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSMaturityCopyFromCustody:253](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityCopyFromCustody.sql#L253) | WHERE / P12 / IN_VALUE_TYPE | `cs.match_flag` : `CHAR(1)` | `'U'` : `varchar` | `cs.match_flag` → `VARCHAR` | `cs.match_flag IN ('U', 'R')` |
| IC-002232 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSMaturityCopyFromCustody:271](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityCopyFromCustody.sql#L271) | WHERE / P12 / IN_VALUE_TYPE | `cs.match_flag` : `CHAR(1)` | `'U'` : `varchar` | `cs.match_flag` → `VARCHAR` | `cs.match_flag IN ('U', 'R')` |
| IC-002234 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSMaturityCopyFromCustody:347](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityCopyFromCustody.sql#L347) | WHERE / P12 / COMPARISON | `cs.match_flag` : `CHAR(1)` | `'C'` : `varchar` | `cs.match_flag` → `VARCHAR` | `cs.match_flag = 'C'` |
| IC-002257 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.Rvmp_TRSMaturityProcess:250](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityProcess.sql#L250) | WHERE / P12 / IN_VALUE_TYPE | `ISNULL(ms.match_flag, '')` : `CHAR(1)` | `'M'` : `varchar` | `ISNULL(ms.match_flag, '')` → `VARCHAR` | `ISNULL(ms.match_flag, '') NOT IN ('M', 'C')` |
| IC-002259 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSMaturityProcess:285](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityProcess.sql#L285) | WHERE / P12 / IN_VALUE_TYPE | `ms.match_flag` : `CHAR(1)` | `'M'` : `varchar` | `ms.match_flag` → `VARCHAR` | `ms.match_flag IN ('M', 'C')` |
| IC-002262 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSMaturityProcess:376](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityProcess.sql#L376) | WHERE / P12 / IN_VALUE_TYPE | `ms.match_flag` : `CHAR(1)` | `'M'` : `varchar` | `ms.match_flag` → `VARCHAR` | `ms.match_flag IN ('M', 'C')` |
| IC-005585 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateMaturityUpload:328](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L328) | EXPRESSION / P07 / CASE_BRANCH | `MAX(ISNULL(src.match_flag, ''))` : `CHAR(1)` | `result type` : `VARCHAR (combined result; facets unknown)` | `MAX(ISNULL(src.match_flag, ''))` → `VARCHAR` | `CASE WHEN SUM( CASE WHEN ISNULL(src.match_flag, '') = 'C' THEN 1 ELSE 0 END ) > 0 THEN 'C' WHEN SUM( CASE WHEN ISNULL(src.match_flag, '') = 'R' THEN 1 ELSE 0 END ) > 0 THEN 'R' WHEN SUM( CASE WHEN ISNULL(src.match_flag, '') = 'U' THEN 1 ELSE 0 END ) > 0 THEN 'U' WHEN SUM( CASE WHEN ISNULL(src.match_flag, '') = 'A' THEN 1 ELSE 0 END ) > 0 THEN 'A' WHEN SUM( CASE WHEN ISNULL(src.match_flag, '') = 'M' THEN 1 ELSE 0 END ) > 0 THEN 'M' ELSE MAX(ISNULL(src.match_flag, '')) END` |
| IC-005586 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateMaturityUpload:331](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L331) | CASE / P12 / COMPARISON | `ISNULL(src.match_flag, '')` : `CHAR(1)` | `'C'` : `varchar` | `ISNULL(src.match_flag, '')` → `VARCHAR` | `ISNULL(src.match_flag, '') = 'C'` |
| IC-005588 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateMaturityUpload:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L337) | CASE / P12 / COMPARISON | `ISNULL(src.match_flag, '')` : `CHAR(1)` | `'R'` : `varchar` | `ISNULL(src.match_flag, '')` → `VARCHAR` | `ISNULL(src.match_flag, '') = 'R'` |
| IC-005590 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateMaturityUpload:343](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L343) | CASE / P12 / COMPARISON | `ISNULL(src.match_flag, '')` : `CHAR(1)` | `'U'` : `varchar` | `ISNULL(src.match_flag, '')` → `VARCHAR` | `ISNULL(src.match_flag, '') = 'U'` |
| IC-005592 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateMaturityUpload:349](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L349) | CASE / P12 / COMPARISON | `ISNULL(src.match_flag, '')` : `CHAR(1)` | `'A'` : `varchar` | `ISNULL(src.match_flag, '')` → `VARCHAR` | `ISNULL(src.match_flag, '') = 'A'` |
| IC-005594 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateMaturityUpload:355](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L355) | CASE / P12 / COMPARISON | `ISNULL(src.match_flag, '')` : `CHAR(1)` | `'M'` : `varchar` | `ISNULL(src.match_flag, '')` → `VARCHAR` | `ISNULL(src.match_flag, '') = 'M'` |
| IC-002227 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSMaturityCopyFromCustody:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityCopyFromCustody.sql#L162) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cs.match_flag` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `ISNULL(cs.match_flag, '')` |
| IC-002230 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSMaturityCopyFromCustody:262](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityCopyFromCustody.sql#L262) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cs.match_flag` : `CHAR(1)` | `'C'` : `varchar` | `'C'` → `CHAR(1)` | `cs.match_flag = 'C'` |
| IC-002258 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSMaturityProcess:250](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityProcess.sql#L250) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ms.match_flag` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `ISNULL(ms.match_flag, '')` |
| IC-005587 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:331](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L331) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.match_flag` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `ISNULL(src.match_flag, '')` |
| IC-005589 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L337) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.match_flag` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `ISNULL(src.match_flag, '')` |
| IC-005591 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:343](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L343) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.match_flag` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `ISNULL(src.match_flag, '')` |
| IC-005593 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:349](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L349) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.match_flag` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `ISNULL(src.match_flag, '')` |
| IC-005595 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:355](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L355) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.match_flag` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `ISNULL(src.match_flag, '')` |
| IC-005596 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:359](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L359) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.match_flag` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `ISNULL(src.match_flag, '')` |
| IC-007028 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessMaturityUploadAuthorization:274](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessMaturityUploadAuthorization.sql#L274) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mst.match_flag` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `mst.match_flag = 'A'` |
| IC-007033 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessMaturityUploadAuthorization:395](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessMaturityUploadAuthorization.sql#L395) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mst.match_flag` : `CHAR(1)` | `'R'` : `varchar` | `'R'` → `CHAR(1)` | `mst.match_flag = 'R'` |

<a id="kolom-maturity-date"></a>
## maturity_date

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_source_tt.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002255 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSMaturityProcess:248](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityProcess.sql#L248) | WHERE / P13 / COMPARISON | `ms.maturity_date` : `DATETIME` | `@dWorkingDate` : `DATE` | `@dWorkingDate` → `DATETIME` | `ms.maturity_date >= @dWorkingDate` |
| IC-002256 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSMaturityProcess:249](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityProcess.sql#L249) | WHERE / P13 / COMPARISON | `ms.maturity_date` : `DATETIME` | `DATEADD(DAY, 1, @dWorkingDate)` : `DATE` | `DATEADD(DAY, 1, @dWorkingDate)` → `DATETIME` | `ms.maturity_date < DATEADD(DAY, 1, @dWorkingDate)` |
| IC-002260 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSMaturityProcess:286](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityProcess.sql#L286) | WHERE / P13 / COMPARISON | `ms.maturity_date` : `DATETIME` | `@dWorkingDate` : `DATE` | `@dWorkingDate` → `DATETIME` | `ms.maturity_date >= @dWorkingDate` |
| IC-002261 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSMaturityProcess:287](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityProcess.sql#L287) | WHERE / P13 / COMPARISON | `ms.maturity_date` : `DATETIME` | `DATEADD(DAY, 1, @dWorkingDate)` : `DATE` | `DATEADD(DAY, 1, @dWorkingDate)` → `DATETIME` | `ms.maturity_date < DATEADD(DAY, 1, @dWorkingDate)` |

<a id="kolom-net-amount"></a>
## net_amount

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_source_tt.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005578 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:321](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L321) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.net_amount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(src.net_amount, 0)` |
| IC-005582 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:325](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L325) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.net_amount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(src.net_amount, 0)` |

<a id="kolom-nominal"></a>
## nominal

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_source_tt.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005567 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:306](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L306) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.nominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(src.nominal, 0)` |

<a id="kolom-safe-keeping"></a>
## safe_keeping

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_source_tt.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005574 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:316](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L316) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.safe_keeping` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(src.safe_keeping, 0)` |

<a id="kolom-safe-keeping-actual"></a>
## safe_keeping_actual

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_source_tt.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005575 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:317](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L317) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.safe_keeping_actual` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(src.safe_keeping_actual, 0)` |

<a id="kolom-safe-keeping-after-tax"></a>
## safe_keeping_after_tax

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_source_tt.sql#L19). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005577 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:319](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L319) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.safe_keeping_after_tax` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(src.safe_keeping_after_tax, 0)` |

<a id="kolom-safe-keeping-ppn"></a>
## safe_keeping_ppn

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_source_tt.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005576 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:318](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L318) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.safe_keeping_ppn` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(src.safe_keeping_ppn, 0)` |

<a id="kolom-sec-id"></a>
## sec_id

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_source_tt.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 11 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005598 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateMaturityUpload:371](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L371) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sec_id` : `BIGINT` | `#tCustody.sec_id` : `int` | `sec_id` → `int` | `#tCustody.sec_id <- sec_id` |
| IC-002224 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSMaturityCopyFromCustody:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityCopyFromCustody.sql#L137) | WHERE / P11 / COMPARISON | `cs.sec_id` : `BIGINT` | `@nSecId` : `int` | `@nSecId` → `BIGINT` | `cs.sec_id = @nSecId` |
| IC-002225 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSMaturityCopyFromCustody:159](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityCopyFromCustody.sql#L159) | WHERE / P11 / COMPARISON | `cs.sec_id` : `BIGINT` | `@nSecId` : `int` | `@nSecId` → `BIGINT` | `cs.sec_id = @nSecId` |
| IC-002228 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSMaturityCopyFromCustody:250](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityCopyFromCustody.sql#L250) | WHERE / P11 / COMPARISON | `cs.sec_id` : `BIGINT` | `@nSecId` : `int` | `@nSecId` → `BIGINT` | `cs.sec_id = @nSecId` |
| IC-002231 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSMaturityCopyFromCustody:268](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityCopyFromCustody.sql#L268) | WHERE / P11 / COMPARISON | `cs.sec_id` : `BIGINT` | `@nSecId` : `int` | `@nSecId` → `BIGINT` | `cs.sec_id = @nSecId` |
| IC-002233 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSMaturityCopyFromCustody:344](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSMaturityCopyFromCustody.sql#L344) | WHERE / P11 / COMPARISON | `cs.sec_id` : `BIGINT` | `@nSecId` : `int` | `@nSecId` → `BIGINT` | `cs.sec_id = @nSecId` |
| IC-005597 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateMaturityUpload:362](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L362) | WHERE / P11 / COMPARISON | `src.sec_id` : `BIGINT` | `@nSecId` : `int` | `@nSecId` → `BIGINT` | `src.sec_id = @nSecId` |
| IC-007027 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessMaturityUploadAuthorization:270](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessMaturityUploadAuthorization.sql#L270) | JOIN_ON / P11 / COMPARISON | `mst.sec_id` : `BIGINT` | `src.SecId` : `INT` | `src.SecId` → `BIGINT` | `src.SecId = mst.sec_id` |
| IC-007029 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessMaturityUploadAuthorization:282](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessMaturityUploadAuthorization.sql#L282) | JOIN_ON / P11 / COMPARISON | `mst.sec_id` : `BIGINT` | `src.SecId` : `INT` | `src.SecId` → `BIGINT` | `src.SecId = mst.sec_id` |
| IC-007032 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessMaturityUploadAuthorization:391](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessMaturityUploadAuthorization.sql#L391) | JOIN_ON / P11 / COMPARISON | `mst.sec_id` : `BIGINT` | `src.SecId` : `INT` | `src.SecId` → `BIGINT` | `src.SecId = mst.sec_id` |
| IC-007034 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessMaturityUploadAuthorization:403](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessMaturityUploadAuthorization.sql#L403) | JOIN_ON / P11 / COMPARISON | `mst.sec_id` : `BIGINT` | `src.SecId` : `INT` | `src.SecId` → `BIGINT` | `src.SecId = mst.sec_id` |

<a id="kolom-sold-price"></a>
## sold_price

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_source_tt.sql#L31). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005569 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:309](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L309) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.sold_price` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(src.sold_price, 0)` |

<a id="kolom-tax-interest"></a>
## tax_interest

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_source_tt.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005572 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:313](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L313) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.tax_interest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(src.tax_interest, 0)` |

<a id="kolom-tax-on-capital-gain"></a>
## tax_on_capital_gain

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_source_tt.sql#L27). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005580 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateMaturityUpload:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMaturityUpload.sql#L323) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `src.tax_on_capital_gain` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(src.tax_on_capital_gain, 0)` |
