# Retail UAT.dbo.TRSParamIssuer_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 6 lokasi source, 2 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-limit"></a>
## Limit

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamIssuer_TR.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003477 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSCalculateCorporateBond:87](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateCorporateBond.sql#L87) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `isnull(tpi.Limit, 0)` : `MONEY` | `tmp.IssuerLimit` : `decimal(20,4)` | `isnull(tpi.Limit, 0)` → `decimal(20,4)` | `tmp.IssuerLimit = isnull(tpi.Limit, 0)` |
| IC-003479 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSCalculateCorporateBond:88](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateCorporateBond.sql#L88) | DML_CONDITION / P05 / BINARY_Add | `isnull(tpi.Limit, 0)` : `MONEY` | `isnull(TotalAmount,0)` : `decimal(20,4)` | `isnull(tpi.Limit, 0)` → `DECIMAL` | `isnull(tpi.Limit, 0) + isnull(TotalAmount,0)` |
| IC-005012 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamIssuerPopulate:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamIssuerPopulate.sql#L21) | EXPRESSION / P07 / CASE_BRANCH | `tpi.Limit` : `MONEY` | `result type` : `DECIMAL (combined result; facets unknown)` | `tpi.Limit` → `DECIMAL` | `case when tmp.IssuerName is null then tpi.Limit else tmp.SisaAmount end` |
| IC-003478 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSCalculateCorporateBond:87](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateCorporateBond.sql#L87) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tpi.Limit` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(tpi.Limit, 0)` |
| IC-003481 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSCalculateCorporateBond:88](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateCorporateBond.sql#L88) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tpi.Limit` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(tpi.Limit, 0)` |

<a id="kolom-notification"></a>
## Notification

**Tipe dasar:** `DECIMAL(12,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamIssuer_TR.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005013 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamIssuerPopulate:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamIssuerPopulate.sql#L26) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[Notification]` : `DECIMAL(12,2)` | `0` : `int` | `0` → `DECIMAL(12,2)` | `ISNULL([Notification],0)` |
