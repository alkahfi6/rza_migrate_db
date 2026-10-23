# Retail UAT.dbo.TRSSellerPending_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 11 lokasi source, 3 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-lastprocesscode"></a>
## LastProcessCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSellerPending_TM.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007371 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSeller:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSeller.sql#L108) | JOIN_ON / P12 / COMPARISON | `sptm.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `sptm.LastProcessCode` → `VARCHAR` | `sptm.LastProcessCode = 'P'` |
| IC-007372 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSeller:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSeller.sql#L131) | JOIN_ON / P12 / COMPARISON | `sptm.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `sptm.LastProcessCode` → `VARCHAR` | `sptm.LastProcessCode = 'P'` |
| IC-007377 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSeller:157](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSeller.sql#L157) | JOIN_ON / P12 / COMPARISON | `sptm.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `sptm.LastProcessCode` → `VARCHAR` | `sptm.LastProcessCode = 'P'` |
| IC-007382 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSeller:183](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSeller.sql#L183) | JOIN_ON / P12 / COMPARISON | `sptm.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `sptm.LastProcessCode` → `VARCHAR` | `sptm.LastProcessCode = 'P'` |
| IC-007384 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSeller:197](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSeller.sql#L197) | JOIN_ON / P12 / COMPARISON | `sptm.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `sptm.LastProcessCode` → `VARCHAR` | `sptm.LastProcessCode = 'P'` |
| IC-012675 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUploadDataSeller:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUploadDataSeller.sql#L130) | JOIN_ON / P12 / COMPARISON | `p.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `p.LastProcessCode` → `VARCHAR` | `p.LastProcessCode = 'P'` |
| IC-007389 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSeller:249](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSeller.sql#L249) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssellerpending_tm.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trssellerpending_tm.LastProcessCode <- 'P'` |
| IC-007398 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSeller:428](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSeller.sql#L428) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `LastProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `LastProcessCode = 'A'` |
| IC-007400 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSeller:479](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSeller.sql#L479) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `LastProcessCode` : `CHAR(1)` | `'R'` : `varchar` | `'R'` → `CHAR(1)` | `LastProcessCode = 'R'` |

<a id="kolom-perdanaflag"></a>
## PerdanaFlag

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSellerPending_TM.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006016 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateSeller:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSeller.sql#L86) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sptm.PerdanaFlag` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(sptm.PerdanaFlag, 0)` |

<a id="kolom-sekunderflag"></a>
## SekunderFlag

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSellerPending_TM.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006017 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateSeller:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSeller.sql#L89) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sptm.SekunderFlag` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(sptm.SekunderFlag, 0)` |
