# Retail UAT.dbo.RepoSpreadTransaksi_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 20 lokasi source, 1 procedure, 8 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 2 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-emptrxcost"></a>
## EmpTrxCost

**Tipe dasar:** `DECIMAL(18,0)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.RepoSpreadTransaksi_TM.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008111 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveSpread:306](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveSpread.sql#L306) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.EmpTrxCost` : `DECIMAL(18,0)` | `0` : `int` | `0` → `DECIMAL(18,0)` | `ISNULL(tm.EmpTrxCost, 0)` |
| IC-008128 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveSpread:365](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveSpread.sql#L365) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.EmpTrxCost` : `DECIMAL(18,0)` | `0` : `int` | `0` → `DECIMAL(18,0)` | `ISNULL(tm.EmpTrxCost, 0)` |

<a id="kolom-maxamount"></a>
## MaxAmount

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.RepoSpreadTransaksi_TM.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008107 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveSpread:304](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveSpread.sql#L304) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.MaxAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tm.MaxAmount,0)` |
| IC-008124 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveSpread:363](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveSpread.sql#L363) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.MaxAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tm.MaxAmount,0)` |

<a id="kolom-maxmaturity"></a>
## MaxMaturity

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.RepoSpreadTransaksi_TM.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008101 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveSpread:299](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveSpread.sql#L299) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.MaxMaturity` : `INT` | `''` : `varchar` | `''` → `INT` | `ISNULL(tm.MaxMaturity,'')` |
| IC-008118 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveSpread:358](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveSpread.sql#L358) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.MaxMaturity` : `INT` | `''` : `varchar` | `''` → `INT` | `ISNULL(tm.MaxMaturity,'')` |

<a id="kolom-minamount"></a>
## MinAmount

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.RepoSpreadTransaksi_TM.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008105 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveSpread:303](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveSpread.sql#L303) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.MinAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tm.MinAmount,0)` |
| IC-008122 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveSpread:362](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveSpread.sql#L362) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.MinAmount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tm.MinAmount,0)` |

<a id="kolom-minmaturity"></a>
## MinMaturity

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.RepoSpreadTransaksi_TM.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008099 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveSpread:298](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveSpread.sql#L298) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.MinMaturity` : `INT` | `''` : `varchar` | `''` → `INT` | `ISNULL(tm.MinMaturity,'')` |
| IC-008116 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveSpread:357](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveSpread.sql#L357) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.MinMaturity` : `INT` | `''` : `varchar` | `''` → `INT` | `ISNULL(tm.MinMaturity,'')` |

<a id="kolom-mintenor"></a>
## MinTenor

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.RepoSpreadTransaksi_TM.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008103 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveSpread:301](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveSpread.sql#L301) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.MinTenor` : `INT` | `''` : `varchar` | `''` → `INT` | `ISNULL(tm.MinTenor,'')` |
| IC-008120 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveSpread:360](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveSpread.sql#L360) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.MinTenor` : `INT` | `''` : `varchar` | `''` → `INT` | `ISNULL(tm.MinTenor,'')` |

<a id="kolom-repoid"></a>
## RepoId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.RepoSpreadTransaksi_TM.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008094 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRepoSaveSpread:199](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveSpread.sql#L199) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.RepoId` : `INT` | `#tmp_final.RepoId` : `BIGINT` | `tm.RepoId` → `BIGINT` | `#tmp_final.RepoId <- tm.RepoId` |
| IC-008097 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRepoSaveSpread:249](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveSpread.sql#L249) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.RepoId` : `INT` | `#tmp_final.RepoId` : `BIGINT` | `tm.RepoId` → `BIGINT` | `#tmp_final.RepoId <- tm.RepoId` |
| IC-008092 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoSaveSpread:196](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveSpread.sql#L196) | JOIN_ON / P11 / COMPARISON | `tm.RepoId` : `INT` | `tmp.RepoId` : `BIGINT` | `tm.RepoId` → `BIGINT` | `tmp.RepoId = tm.RepoId` |
| IC-008095 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoSaveSpread:244](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveSpread.sql#L244) | JOIN_ON / P11 / COMPARISON | `tm.RepoId` : `INT` | `tmp.RepoId` : `BIGINT` | `tm.RepoId` → `BIGINT` | `tmp.RepoId = tm.RepoId` |
| IC-008098 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoSaveSpread:294](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveSpread.sql#L294) | JOIN_ON / P11 / COMPARISON | `tm.RepoId` : `INT` | `tmp.RepoId` : `BIGINT` | `tm.RepoId` → `BIGINT` | `tmp.RepoId = tm.RepoId` |
| IC-008115 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRepoSaveSpread:353](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveSpread.sql#L353) | JOIN_ON / P11 / COMPARISON | `tm.RepoId` : `INT` | `tmp.RepoId` : `BIGINT` | `tm.RepoId` → `BIGINT` | `tmp.RepoId = tm.RepoId` |

<a id="kolom-trxcost"></a>
## TrxCost

**Tipe dasar:** `DECIMAL(18,0)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.RepoSpreadTransaksi_TM.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008109 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveSpread:305](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveSpread.sql#L305) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.TrxCost` : `DECIMAL(18,0)` | `0` : `int` | `0` → `DECIMAL(18,0)` | `ISNULL(tm.TrxCost, 0)` |
| IC-008126 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoSaveSpread:364](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoSaveSpread.sql#L364) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tm.TrxCost` : `DECIMAL(18,0)` | `0` : `int` | `0` → `DECIMAL(18,0)` | `ISNULL(tm.TrxCost, 0)` |
