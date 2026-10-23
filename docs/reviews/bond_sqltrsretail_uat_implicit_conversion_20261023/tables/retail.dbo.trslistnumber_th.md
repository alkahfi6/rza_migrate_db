# Retail UAT.dbo.TRSListNumber_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 3 lokasi source, 2 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-number"></a>
## Number

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSListNumber_TH.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001866 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.RepoRONDataMigrate:9](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L9) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `Number` : `BIGINT` | `@nlastRepoDealId` : `int` | `Number` → `int` | `@nlastRepoDealId = Number` |
| IC-001884 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoRONDataMigrate:83](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L83) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Number` : `BIGINT` | `@nMax` : `int` | `@nMax` → `BIGINT` | `Number = @nMax` |
| IC-003007 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SecurityGetXMLLelang:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityGetXMLLelang.sql#L51) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `Number` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `ISNULL(Number, 0)` |
