# SQL_REPLICATE.dbo.MXReasonReverse — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 5 lokasi source, 1 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-m-cddesc"></a>
## M_CDDESC

**Tipe dasar:** `NVARCHAR(60)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.MXReasonReverse.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P06 — LIKE dan representasi karakter:** LIKE pada key numerik memerlukan representasi karakter. Equality numeric yang tervalidasi cocok untuk lookup tepat; prefix/substring membutuhkan requirement pencarian tersendiri, bukan cast kolom sebagai solusi index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012146 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSearchReasonMX:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchReasonMX.sql#L25) | WHERE / P06 / LIKE | `upper(M_CDDESC)` : `NVARCHAR (expression; length unknown)` | `@cCol2 + '%'` : `VARCHAR (combined result; facets unknown)` | `@cCol2 + '%'` → `NVARCHAR` | `upper(M_CDDESC) like @cCol2 + '%'` |
| IC-012147 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSearchReasonMX:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchReasonMX.sql#L31) | WHERE / P06 / LIKE | `upper(M_CDDESC)` : `NVARCHAR (expression; length unknown)` | `@cCol2 + '%'` : `VARCHAR (combined result; facets unknown)` | `@cCol2 + '%'` → `NVARCHAR` | `upper(M_CDDESC) like @cCol2 + '%'` |

<a id="kolom-m-code"></a>
## M_CODE

**Tipe dasar:** `NVARCHAR(10)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.MXReasonReverse.sql#L6). **Existing leading index:** MXReasonReverse_CN1.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P06 — LIKE dan representasi karakter:** LIKE pada key numerik memerlukan representasi karakter. Equality numeric yang tervalidasi cocok untuk lookup tepat; prefix/substring membutuhkan requirement pencarian tersendiri, bukan cast kolom sebagai solusi index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012144 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSearchReasonMX:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchReasonMX.sql#L17) | WHERE / P04 / COMPARISON | `M_CODE` : `NVARCHAR(10)` | `@cCol1` : `varchar(40)` | `@cCol1` → `NVARCHAR` | `M_CODE = @cCol1` |
| IC-012145 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSearchReasonMX:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchReasonMX.sql#L24) | WHERE / P06 / LIKE | `M_CODE` : `NVARCHAR(10)` | `@cCol1 + '%'` : `VARCHAR (combined result; facets unknown)` | `@cCol1 + '%'` → `NVARCHAR` | `M_CODE like @cCol1 + '%'` |
| IC-012148 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSearchReasonMX:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchReasonMX.sql#L37) | WHERE / P06 / LIKE | `M_CODE` : `NVARCHAR(10)` | `@cCol1 + '%'` : `VARCHAR (combined result; facets unknown)` | `@cCol1 + '%'` → `NVARCHAR` | `M_CODE like @cCol1 + '%'` |
