# Retail UAT.dbo.TRSParamLelang_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 23 lokasi source, 4 procedure, 11 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-batchminbranch"></a>
## BatchMinBranch

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamLelang_TR.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005095 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamLelangSave:202](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangSave.sql#L202) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pl.BatchMinBranch` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(pl.BatchMinBranch,0)` |
| IC-005116 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamLelangSave:231](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangSave.sql#L231) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pl.BatchMinBranch` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(pl.BatchMinBranch,0)` |

<a id="kolom-batchminta"></a>
## BatchMinTA

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamLelang_TR.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005093 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamLelangSave:201](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangSave.sql#L201) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pl.BatchMinTA` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(pl.BatchMinTA,0)` |
| IC-005114 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamLelangSave:230](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangSave.sql#L230) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pl.BatchMinTA` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(pl.BatchMinTA,0)` |

<a id="kolom-hargamodal"></a>
## HargaModal

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamLelang_TR.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005097 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamLelangSave:203](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangSave.sql#L203) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pl.HargaModal` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(pl.HargaModal,0)` |
| IC-005118 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamLelangSave:232](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangSave.sql#L232) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pl.HargaModal` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(pl.HargaModal,0)` |

<a id="kolom-maxspread"></a>
## MaxSpread

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamLelang_TR.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005091 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamLelangSave:200](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangSave.sql#L200) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pl.MaxSpread` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(pl.MaxSpread,0)` |
| IC-005112 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamLelangSave:229](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangSave.sql#L229) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pl.MaxSpread` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(pl.MaxSpread,0)` |

<a id="kolom-minspread"></a>
## MinSpread

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamLelang_TR.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005089 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamLelangSave:199](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangSave.sql#L199) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pl.MinSpread` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(pl.MinSpread,0)` |
| IC-005110 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamLelangSave:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangSave.sql#L228) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pl.MinSpread` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(pl.MinSpread,0)` |

<a id="kolom-nominalkelipatan"></a>
## NominalKelipatan

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamLelang_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005083 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamLelangSave:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangSave.sql#L193) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pl.NominalKelipatan` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(pl.NominalKelipatan,0)` |
| IC-005104 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamLelangSave:222](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangSave.sql#L222) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pl.NominalKelipatan` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(pl.NominalKelipatan,0)` |

<a id="kolom-nominalminimum"></a>
## NominalMinimum

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamLelang_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005081 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamLelangSave:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangSave.sql#L192) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pl.NominalMinimum` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(pl.NominalMinimum,0)` |
| IC-005102 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamLelangSave:221](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangSave.sql#L221) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pl.NominalMinimum` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(pl.NominalMinimum,0)` |

<a id="kolom-secid"></a>
## SecId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamLelang_TR.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003040 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SecurityPopulateResultLelangEmail:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateResultLelangEmail.sql#L51) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `SecId` : `INT` | `@nSecId` : `BIGINT` | `SecId` → `BIGINT` | `@nSecId = SecId` |

<a id="kolom-settlementdate"></a>
## SettlementDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamLelang_TR.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005066 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamLelangProcessApproval:151](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangProcessApproval.sql#L151) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pl.SettlementDate` : `DATETIME` | `0` : `int` | `0` → `DATETIME` | `isnull(pl.SettlementDate,0)` |
| IC-005099 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamLelangSave:204](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangSave.sql#L204) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pl.SettlementDate` : `DATETIME` | `'19000101'` : `varchar` | `'19000101'` → `DATETIME` | `isnull(pl.SettlementDate,'19000101')` |
| IC-005120 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamLelangSave:233](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangSave.sql#L233) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pl.SettlementDate` : `DATETIME` | `'19000101'` : `varchar` | `'19000101'` → `DATETIME` | `isnull(pl.SettlementDate,'19000101')` |

<a id="kolom-tglakhir"></a>
## TglAkhir

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamLelang_TR.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003060 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SecurityPopulateResultLelangProduct:58](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateResultLelangProduct.sql#L58) | JOIN_ON / P04 / COMPARISON | `tr.TglAkhir` : `DATETIME` | `CONVERT(VARCHAR(10),tt.LastUpdateDate,112)` : `VARCHAR(10)` | `CONVERT(VARCHAR(10),tt.LastUpdateDate,112)` → `DATETIME` | `CONVERT(VARCHAR(10),tt.LastUpdateDate,112) = tr.TglAkhir` |
| IC-005087 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamLelangSave:196](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangSave.sql#L196) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pl.TglAkhir` : `DATETIME` | `'19000101'` : `varchar` | `'19000101'` → `DATETIME` | `isnull(pl.TglAkhir,'19000101')` |
| IC-005108 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamLelangSave:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangSave.sql#L225) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pl.TglAkhir` : `DATETIME` | `'19000101'` : `varchar` | `'19000101'` → `DATETIME` | `isnull(pl.TglAkhir,'19000101')` |

<a id="kolom-tglmulai"></a>
## TglMulai

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamLelang_TR.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005085 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamLelangSave:194](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangSave.sql#L194) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pl.TglMulai` : `DATETIME` | `'19000101'` : `varchar` | `'19000101'` → `DATETIME` | `isnull(pl.TglMulai,'19000101')` |
| IC-005106 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamLelangSave:223](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangSave.sql#L223) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pl.TglMulai` : `DATETIME` | `'19000101'` : `varchar` | `'19000101'` → `DATETIME` | `isnull(pl.TglMulai,'19000101')` |
