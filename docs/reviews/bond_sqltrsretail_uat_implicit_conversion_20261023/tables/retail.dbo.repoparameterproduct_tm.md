# Retail UAT.dbo.REPOParameterProduct_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 55 lokasi source, 7 procedure, 9 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 17 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-admfee"></a>
## AdmFee

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOParameterProduct_TM.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001963 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.RepoUploadValidation:189](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoUploadValidation.sql#L189) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tm.AdmFee` : `MONEY` | `AdmnFee` : `DECIMAL(20,2)` | `tm.AdmFee` → `DECIMAL(20,2)` | `AdmnFee = tm.AdmFee` |
| IC-008063 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveData:526](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L526) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.AdmFee` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tm.AdmFee,0)` |
| IC-008084 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveData:610](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L610) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.AdmFee` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tm.AdmFee,0)` |

<a id="kolom-batasbawahspread"></a>
## BatasBawahSpread

**Tipe dasar:** `DECIMAL(15,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOParameterProduct_TM.sql#L28). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008028 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRepoSaveData:278](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L278) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.BatasBawahSpread` : `DECIMAL(15,2)` | `#tmp_final.BatasBawahSpread` : `MONEY` | `tm.BatasBawahSpread` → `MONEY` | `#tmp_final.BatasBawahSpread <- tm.BatasBawahSpread` |
| IC-008035 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRepoSaveData:360](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L360) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.BatasBawahSpread` : `DECIMAL(15,2)` | `#tmp_final.BatasBawahSpread` : `MONEY` | `tm.BatasBawahSpread` → `MONEY` | `#tmp_final.BatasBawahSpread <- tm.BatasBawahSpread` |
| IC-008042 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRepoSaveData:439](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L439) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.BatasBawahSpread` : `DECIMAL(15,2)` | `#tmp_final.BatasBawahSpread` : `MONEY` | `tm.BatasBawahSpread` → `MONEY` | `#tmp_final.BatasBawahSpread <- tm.BatasBawahSpread` |
| IC-008058 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveData:517](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L517) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.BatasBawahSpread` : `DECIMAL(15,2)` | `0` : `int` | `0` → `DECIMAL(15,2)` | `ISNULL(tm.BatasBawahSpread,0)` |
| IC-008059 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveData:517](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L517) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.BatasBawahSpread` : `DECIMAL(15,2)` | `0` : `int` | `0` → `DECIMAL(15,2)` | `ISNULL(tm.BatasBawahSpread,0)` |
| IC-008078 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRepoSaveData:601](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L601) | WHERE / P11 / COMPARISON | `ISNULL(tm.BatasBawahSpread,0)` : `DECIMAL(15,2)` | `ISNULL(tmp.BatasBawahSpread,0)` : `MONEY` | `ISNULL(tmp.BatasBawahSpread,0)` → `DECIMAL` | `ISNULL(tm.BatasBawahSpread,0) <> ISNULL(tmp.BatasBawahSpread,0)` |
| IC-008079 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveData:601](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L601) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.BatasBawahSpread` : `DECIMAL(15,2)` | `0` : `int` | `0` → `DECIMAL(15,2)` | `ISNULL(tm.BatasBawahSpread,0)` |

<a id="kolom-haircut"></a>
## Haircut

**Tipe dasar:** `DECIMAL(15,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOParameterProduct_TM.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008029 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRepoSaveData:278](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L278) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.Haircut` : `DECIMAL(15,2)` | `#tmp_final.Haircut` : `MONEY` | `tm.Haircut` → `MONEY` | `#tmp_final.Haircut <- tm.Haircut` |
| IC-008036 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRepoSaveData:360](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L360) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.Haircut` : `DECIMAL(15,2)` | `#tmp_final.Haircut` : `MONEY` | `tm.Haircut` → `MONEY` | `#tmp_final.Haircut <- tm.Haircut` |
| IC-008043 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRepoSaveData:439](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L439) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.Haircut` : `DECIMAL(15,2)` | `#tmp_final.Haircut` : `MONEY` | `tm.Haircut` → `MONEY` | `#tmp_final.Haircut <- tm.Haircut` |
| IC-008060 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRepoSaveData:518](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L518) | WHERE / P11 / COMPARISON | `ISNULL(tm.Haircut,0)` : `DECIMAL(15,2)` | `ISNULL(tmp.Haircut,0)` : `MONEY` | `ISNULL(tmp.Haircut,0)` → `DECIMAL` | `ISNULL(tm.Haircut,0) <> ISNULL(tmp.Haircut,0)` |
| IC-008061 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveData:518](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L518) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.Haircut` : `DECIMAL(15,2)` | `0` : `int` | `0` → `DECIMAL(15,2)` | `ISNULL(tm.Haircut,0)` |
| IC-008081 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRepoSaveData:602](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L602) | WHERE / P11 / COMPARISON | `ISNULL(tm.Haircut,0)` : `DECIMAL(15,2)` | `ISNULL(tmp.Haircut,0)` : `MONEY` | `ISNULL(tmp.Haircut,0)` → `DECIMAL` | `ISNULL(tm.Haircut,0) <> ISNULL(tmp.Haircut,0)` |
| IC-008082 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveData:602](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L602) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.Haircut` : `DECIMAL(15,2)` | `0` : `int` | `0` → `DECIMAL(15,2)` | `ISNULL(tm.Haircut,0)` |

<a id="kolom-maxspread"></a>
## MaxSpread

**Tipe dasar:** `DECIMAL(15,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOParameterProduct_TM.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008030 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRepoSaveData:278](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L278) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.MaxSpread` : `DECIMAL(15,2)` | `#tmp_final.MaxSpread` : `MONEY` | `tm.MaxSpread` → `MONEY` | `#tmp_final.MaxSpread <- tm.MaxSpread` |
| IC-008037 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRepoSaveData:360](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L360) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.MaxSpread` : `DECIMAL(15,2)` | `#tmp_final.MaxSpread` : `MONEY` | `tm.MaxSpread` → `MONEY` | `#tmp_final.MaxSpread <- tm.MaxSpread` |
| IC-008044 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRepoSaveData:439](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L439) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.MaxSpread` : `DECIMAL(15,2)` | `#tmp_final.MaxSpread` : `MONEY` | `tm.MaxSpread` → `MONEY` | `#tmp_final.MaxSpread <- tm.MaxSpread` |
| IC-008055 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRepoSaveData:516](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L516) | WHERE / P11 / COMPARISON | `ISNULL(tm.MaxSpread,0)` : `DECIMAL(15,2)` | `ISNULL(tmp.MaxSpread,0)` : `MONEY` | `ISNULL(tmp.MaxSpread,0)` → `DECIMAL` | `ISNULL(tm.MaxSpread,0) <> ISNULL(tmp.MaxSpread,0)` |
| IC-008056 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveData:516](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L516) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.MaxSpread` : `DECIMAL(15,2)` | `0` : `int` | `0` → `DECIMAL(15,2)` | `ISNULL(tm.MaxSpread,0)` |
| IC-008075 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRepoSaveData:600](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L600) | WHERE / P11 / COMPARISON | `ISNULL(tm.MaxSpread,0)` : `DECIMAL(15,2)` | `ISNULL(tmp.MaxSpread,0)` : `MONEY` | `ISNULL(tmp.MaxSpread,0)` → `DECIMAL` | `ISNULL(tm.MaxSpread,0) <> ISNULL(tmp.MaxSpread,0)` |
| IC-008076 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveData:600](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L600) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.MaxSpread` : `DECIMAL(15,2)` | `0` : `int` | `0` → `DECIMAL(15,2)` | `ISNULL(tm.MaxSpread,0)` |

<a id="kolom-minspread"></a>
## MinSpread

**Tipe dasar:** `DECIMAL(15,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOParameterProduct_TM.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008031 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRepoSaveData:278](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L278) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.MinSpread` : `DECIMAL(15,2)` | `#tmp_final.MinSpread` : `MONEY` | `tm.MinSpread` → `MONEY` | `#tmp_final.MinSpread <- tm.MinSpread` |
| IC-008038 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRepoSaveData:360](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L360) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.MinSpread` : `DECIMAL(15,2)` | `#tmp_final.MinSpread` : `MONEY` | `tm.MinSpread` → `MONEY` | `#tmp_final.MinSpread <- tm.MinSpread` |
| IC-008045 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRepoSaveData:439](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L439) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.MinSpread` : `DECIMAL(15,2)` | `#tmp_final.MinSpread` : `MONEY` | `tm.MinSpread` → `MONEY` | `#tmp_final.MinSpread <- tm.MinSpread` |
| IC-008052 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRepoSaveData:515](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L515) | WHERE / P11 / COMPARISON | `ISNULL(tm.MinSpread,0)` : `DECIMAL(15,2)` | `ISNULL(tmp.MinSpread,0)` : `MONEY` | `ISNULL(tmp.MinSpread,0)` → `DECIMAL` | `ISNULL(tm.MinSpread,0) <> ISNULL(tmp.MinSpread,0)` |
| IC-008053 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveData:515](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L515) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.MinSpread` : `DECIMAL(15,2)` | `0` : `int` | `0` → `DECIMAL(15,2)` | `ISNULL(tm.MinSpread,0)` |
| IC-008072 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRepoSaveData:599](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L599) | WHERE / P11 / COMPARISON | `ISNULL(tm.MinSpread,0)` : `DECIMAL(15,2)` | `ISNULL(tmp.MinSpread,0)` : `MONEY` | `ISNULL(tmp.MinSpread,0)` → `DECIMAL` | `ISNULL(tm.MinSpread,0) <> ISNULL(tmp.MinSpread,0)` |
| IC-008073 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveData:599](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L599) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.MinSpread` : `DECIMAL(15,2)` | `0` : `int` | `0` → `DECIMAL(15,2)` | `ISNULL(tm.MinSpread,0)` |

<a id="kolom-nominalkelipatan"></a>
## NominalKelipatan

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOParameterProduct_TM.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008050 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveData:512](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L512) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.NominalKelipatan` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tm.NominalKelipatan,0)` |
| IC-008070 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveData:596](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L596) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.NominalKelipatan` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tm.NominalKelipatan,0)` |

<a id="kolom-nominalminimum"></a>
## NominalMinimum

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOParameterProduct_TM.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008048 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveData:511](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L511) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.NominalMinimum` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tm.NominalMinimum,0)` |
| IC-008068 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveData:595](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L595) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.NominalMinimum` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tm.NominalMinimum,0)` |

<a id="kolom-repoprodid"></a>
## RepoProdId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOParameterProduct_TM.sql#L2). **Existing leading index:** REPOParameterProduct_TM_CN1.

**Frekuensi source:** 15 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005195 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSParamRepoValidate:841](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamRepoValidate.sql#L841) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ISNULL(tm.RepoProdId, 0)` : `INT` | `tmp.RepoProdId` : `BIGINT` | `ISNULL(tm.RepoProdId, 0)` → `BIGINT` | `tmp.RepoProdId = ISNULL(tm.RepoProdId, 0)` |
| IC-008032 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRepoSaveData:278](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L278) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.RepoProdId` : `INT` | `#tmp_final.RepoProdId` : `BIGINT` | `tm.RepoProdId` → `BIGINT` | `#tmp_final.RepoProdId <- tm.RepoProdId` |
| IC-008039 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRepoSaveData:360](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L360) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.RepoProdId` : `INT` | `#tmp_final.RepoProdId` : `BIGINT` | `tm.RepoProdId` → `BIGINT` | `#tmp_final.RepoProdId <- tm.RepoProdId` |
| IC-008046 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRepoSaveData:439](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L439) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.RepoProdId` : `INT` | `#tmp_final.RepoProdId` : `BIGINT` | `tm.RepoProdId` → `BIGINT` | `#tmp_final.RepoProdId <- tm.RepoProdId` |
| IC-001798 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.RepoCalculateProduct:55](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoCalculateProduct.sql#L55) | WHERE / P11 / COMPARISON | `RepoProdId` : `INT` | `@pnRepoId` : `BIGINT` | `RepoProdId` → `BIGINT` | `RepoProdId = @pnRepoId` |
| IC-001799 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.RepoCalculateProduct:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoCalculateProduct.sql#L59) | WHERE / P11 / COMPARISON | `RepoProdId` : `INT` | `@pnRepoId` : `BIGINT` | `RepoProdId` → `BIGINT` | `RepoProdId = @pnRepoId` |
| IC-001848 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.RepoPopulateDetailProduct:62](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoPopulateDetailProduct.sql#L62) | WHERE / P11 / COMPARISON | `tm.RepoProdId` : `INT` | `@pnRepoId` : `BIGINT` | `tm.RepoProdId` → `BIGINT` | `tm.RepoProdId = @pnRepoId` |
| IC-003311 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamRepo:232](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamRepo.sql#L232) | JOIN_ON / P11 / COMPARISON | `tm.RepoProdId` : `INT` | `tr.RepoProdId` : `BIGINT` | `tm.RepoProdId` → `BIGINT` | `tm.RepoProdId = tr.RepoProdId` |
| IC-003313 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamRepo:266](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamRepo.sql#L266) | JOIN_ON / P11 / COMPARISON | `tm.RepoProdId` : `INT` | `tr.RepoProdId` : `BIGINT` | `tm.RepoProdId` → `BIGINT` | `tm.RepoProdId = tr.RepoProdId` |
| IC-003317 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApprovalParamRepo:335](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamRepo.sql#L335) | JOIN_ON / P11 / COMPARISON | `tm.RepoProdId` : `INT` | `tr.RepoProdId` : `BIGINT` | `tm.RepoProdId` → `BIGINT` | `tm.RepoProdId = tr.RepoProdId` |
| IC-008026 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoSaveData:272](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L272) | JOIN_ON / P11 / COMPARISON | `tm.RepoProdId` : `INT` | `tmp.RepoProdId` : `BIGINT` | `tm.RepoProdId` → `BIGINT` | `tmp.RepoProdId = tm.RepoProdId` |
| IC-008033 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoSaveData:346](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L346) | JOIN_ON / P11 / COMPARISON | `tm.RepoProdId` : `INT` | `tmp.RepoProdId` : `BIGINT` | `tm.RepoProdId` → `BIGINT` | `tmp.RepoProdId = tm.RepoProdId` |
| IC-008040 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoSaveData:428](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L428) | JOIN_ON / P11 / COMPARISON | `tm.RepoProdId` : `INT` | `tmp.RepoProdId` : `BIGINT` | `tm.RepoProdId` → `BIGINT` | `tmp.RepoProdId = tm.RepoProdId` |
| IC-008047 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoSaveData:507](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L507) | JOIN_ON / P11 / COMPARISON | `tm.RepoProdId` : `INT` | `tmp.RepoProdId` : `BIGINT` | `tm.RepoProdId` → `BIGINT` | `tmp.RepoProdId = tm.RepoProdId` |
| IC-008067 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoSaveData:591](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveData.sql#L591) | JOIN_ON / P11 / COMPARISON | `tm.RepoProdId` : `INT` | `tmp.RepoProdId` : `BIGINT` | `tm.RepoProdId` → `BIGINT` | `tmp.RepoProdId = tm.RepoProdId` |

<a id="kolom-isactive"></a>
## isActive

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOParameterProduct_TM.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008011 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductSrc:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductSrc.sql#L77) | WHERE / P11 / COMPARISON | `rep.isActive` : `BIT` | `1` : `int` | `rep.isActive` → `INT` | `rep.isActive = 1` |
| IC-008014 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductSrc:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductSrc.sql#L106) | WHERE / P11 / COMPARISON | `rep.isActive` : `BIT` | `1` : `int` | `rep.isActive` → `INT` | `rep.isActive = 1` |
| IC-008017 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductSrc:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductSrc.sql#L131) | WHERE / P11 / COMPARISON | `rep.isActive` : `BIT` | `1` : `int` | `rep.isActive` → `INT` | `rep.isActive = 1` |
| IC-008020 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductSrc:157](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductSrc.sql#L157) | WHERE / P11 / COMPARISON | `rep.isActive` : `BIT` | `1` : `int` | `rep.isActive` → `INT` | `rep.isActive = 1` |
| IC-008023 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoProductSrc:183](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoProductSrc.sql#L183) | WHERE / P11 / COMPARISON | `rep.isActive` : `BIT` | `1` : `int` | `rep.isActive` → `INT` | `rep.isActive = 1` |
