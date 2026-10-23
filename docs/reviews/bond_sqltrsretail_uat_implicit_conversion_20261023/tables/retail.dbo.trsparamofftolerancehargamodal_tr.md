# Retail UAT.dbo.TRSParamOfftoleranceHargaModal_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 10 lokasi source, 1 procedure, 4 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-bounds"></a>
## Bounds

**Tipe dasar:** `DECIMAL(10,4)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamOfftoleranceHargaModal_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012318 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitParamOfftoleranceHargaModal:140](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitParamOfftoleranceHargaModal.sql#L140) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `par.Bounds` : `DECIMAL(10,4)` | `0` : `int` | `0` → `DECIMAL(10,4)` | `ISNULL(par.Bounds,0)` |
| IC-012325 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitParamOfftoleranceHargaModal:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitParamOfftoleranceHargaModal.sql#L156) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `par.Bounds` : `DECIMAL(10,4)` | `0` : `int` | `0` → `DECIMAL(10,4)` | `ISNULL(par.Bounds,0)` |
| IC-012332 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitParamOfftoleranceHargaModal:167](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitParamOfftoleranceHargaModal.sql#L167) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `par.Bounds` : `DECIMAL(10,4)` | `0` : `int` | `0` → `DECIMAL(10,4)` | `ISNULL(par.Bounds,0)` |

<a id="kolom-endperiode"></a>
## EndPeriode

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamOfftoleranceHargaModal_TR.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012316 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitParamOfftoleranceHargaModal:139](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitParamOfftoleranceHargaModal.sql#L139) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `par.EndPeriode` : `DATETIME` | `'19920706'` : `varchar` | `'19920706'` → `DATETIME` | `ISNULL(par.EndPeriode,'19920706')` |
| IC-012323 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitParamOfftoleranceHargaModal:155](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitParamOfftoleranceHargaModal.sql#L155) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `par.EndPeriode` : `DATETIME` | `'19920706'` : `varchar` | `'19920706'` → `DATETIME` | `ISNULL(par.EndPeriode,'19920706')` |
| IC-012330 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitParamOfftoleranceHargaModal:166](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitParamOfftoleranceHargaModal.sql#L166) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `par.EndPeriode` : `DATETIME` | `'19920706'` : `varchar` | `'19920706'` → `DATETIME` | `ISNULL(par.EndPeriode,'19920706')` |

<a id="kolom-lastupdatenik"></a>
## LastUpdateNIK

**Tipe dasar:** `VARCHAR(10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamOfftoleranceHargaModal_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012312 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSubmitParamOfftoleranceHargaModal:116](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitParamOfftoleranceHargaModal.sql#L116) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamofftolerancehargamodal_tr.LastUpdateNIK` : `VARCHAR(10)` | `@pnNIK` : `INT` | `@pnNIK` → `VARCHAR(10)` | `retail.dbo.trsparamofftolerancehargamodal_tr.LastUpdateNIK <- @pnNIK` |

<a id="kolom-startperiode"></a>
## StartPeriode

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamOfftoleranceHargaModal_TR.sql#L3). **Existing leading index:** TRSParamOfftoleranceHargaModal_TR_CN1.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012314 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitParamOfftoleranceHargaModal:138](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitParamOfftoleranceHargaModal.sql#L138) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `par.StartPeriode` : `DATETIME` | `'19920706'` : `varchar` | `'19920706'` → `DATETIME` | `ISNULL(par.StartPeriode,'19920706')` |
| IC-012321 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitParamOfftoleranceHargaModal:154](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitParamOfftoleranceHargaModal.sql#L154) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `par.StartPeriode` : `DATETIME` | `'19920706'` : `varchar` | `'19920706'` → `DATETIME` | `ISNULL(par.StartPeriode,'19920706')` |
| IC-012328 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitParamOfftoleranceHargaModal:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitParamOfftoleranceHargaModal.sql#L165) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `par.StartPeriode` : `DATETIME` | `'19920706'` : `varchar` | `'19920706'` → `DATETIME` | `ISNULL(par.StartPeriode,'19920706')` |
