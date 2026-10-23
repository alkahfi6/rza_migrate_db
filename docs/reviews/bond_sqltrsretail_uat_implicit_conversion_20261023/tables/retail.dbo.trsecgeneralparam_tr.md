# Retail UAT.dbo.TRSECGeneralParam_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 2 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-paramvalue"></a>
## ParamValue

**Tipe dasar:** `VARCHAR(100)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECGeneralParam_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000718 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.OMGetSystemNIK:6](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetSystemNIK.sql#L6) | EXPRESSION / P07 / CASE_BRANCH | `ParamValue` : `VARCHAR(100)` | `result type` : `INT (combined result; facets unknown)` | `ParamValue` → `INT` | `case when isnumeric(ParamValue) = 0 then 0 else ParamValue end` |
| IC-006600 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessApprParamSpread:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L147) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ParamValue` : `VARCHAR(100)` | `@nNewLockTime` : `decimal(10,2)` | `@nNewLockTime` → `VARCHAR(100)` | `ParamValue = @nNewLockTime` |
