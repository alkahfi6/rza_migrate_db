# Retail UAT.dbo.TRSProcessTransaksiDetail_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 3 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-processnik"></a>
## ProcessNIK

**Tipe dasar:** `VARCHAR(10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSProcessTransaksiDetail_TR.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004233 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertProcessTransaksi:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertProcessTransaksi.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsprocesstransaksidetail_tr.ProcessNIK` : `VARCHAR(10)` | `@pnNIK` : `INT` | `@pnNIK` → `VARCHAR(10)` | `retail.dbo.trsprocesstransaksidetail_tr.ProcessNIK <- @pnNIK` |
| IC-012648 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateStatusProcessTransaksi:100](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateStatusProcessTransaksi.sql#L100) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ProcessNIK` : `VARCHAR(10)` | `@pnNIK` : `INT` | `@pnNIK` → `VARCHAR(10)` | `ProcessNIK = @pnNIK` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSProcessTransaksiDetail_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007990 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRepoInsertEmailProcess:60](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoInsertEmailProcess.sql#L60) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `ISNULL(pr.[Status], CAST(0 AS BIT))` : `INT` | `@bIsPrevSeqOK` : `BIT` | `ISNULL(pr.[Status], CAST(0 AS BIT))` → `BIT` | `@bIsPrevSeqOK = ISNULL(pr.[Status], CAST(0 AS BIT))` |
| IC-007989 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRepoInsertEmailProcess:60](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoInsertEmailProcess.sql#L60) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pr.[Status]` : `INT` | `CAST(0 AS BIT)` : `BIT` | `CAST(0 AS BIT)` → `INT` | `ISNULL(pr.[Status], CAST(0 AS BIT))` |
