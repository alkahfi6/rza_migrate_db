# Retail UAT.dbo.TRSParamCounter_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 5 lokasi source, 3 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-paramcounter"></a>
## ParamCounter

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamCounter_TM.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003652 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSCreateMT541Content:42](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCreateMT541Content.sql#L42) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `ParamCounter` : `BIGINT` | `@nCounter` : `int` | `ParamCounter` → `int` | `@nCounter = ParamCounter` |
| IC-003651 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSCounterResetMTHeaderJob:19](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCounterResetMTHeaderJob.sql#L19) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ParamCounter` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `ParamCounter = 0` |
| IC-003655 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSCreateMT541Content:185](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCreateMT541Content.sql#L185) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ParamCounter` : `BIGINT` | `@nCounter` : `int` | `@nCounter` → `BIGINT` | `ParamCounter = @nCounter` |
| IC-003656 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSCreateMTContentForBIS4:9](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCreateMTContentForBIS4.sql#L9) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ParamCounter` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `isnull(ParamCounter,0)` |
| IC-003659 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSCreateMTContentForBIS4:143](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCreateMTContentForBIS4.sql#L143) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ParamCounter` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `isnull(ParamCounter,0)` |
