# SQL_REPLICATE.dbo.MXPergerakanDana — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-m-code"></a>
## M_CODE

**Tipe dasar:** `NVARCHAR(50)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.MXPergerakanDana.sql#L2). **Existing leading index:** MXPergerakanDana_CN1.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005927 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateParamInterfaceMurex:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamInterfaceMurex.sql#L18) | SET_OUTPUT / P07 / UNION_OUTPUT | `left output #1` : `NVARCHAR(50)` | `right output #1` : `varchar` | `right output #1` → `NVARCHAR` | `SELECT M_CODE as Value, M_CODE as Display FROM dbo.MXPergerakanDana_v union select '' as Value, '' as Display` |
| IC-005928 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateParamInterfaceMurex:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamInterfaceMurex.sql#L18) | SET_OUTPUT / P07 / UNION_OUTPUT | `left output #2` : `NVARCHAR(50)` | `right output #2` : `varchar` | `right output #2` → `NVARCHAR` | `SELECT M_CODE as Value, M_CODE as Display FROM dbo.MXPergerakanDana_v union select '' as Value, '' as Display` |
