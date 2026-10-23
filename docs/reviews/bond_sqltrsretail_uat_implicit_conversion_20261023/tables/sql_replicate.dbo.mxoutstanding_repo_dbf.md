# SQL_REPLICATE.dbo.MXOUTSTANDING_REPO_DBF — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-m-costcenter"></a>
## M_COSTCENTER

**Tipe dasar:** `NVARCHAR(5)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.MXOUTSTANDING_REPO_DBF.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000072 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.BKPopulateMXRepo:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BKPopulateMXRepo.sql#L37) | WHERE / P04 / COMPARISON | `M_COSTCENTER` : `NVARCHAR(5)` | `@cCostSyar` : `VARCHAR(20)` | `@cCostSyar` → `NVARCHAR` | `M_COSTCENTER = @cCostSyar` |
| IC-000073 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.BKPopulateMXRepo:39](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BKPopulateMXRepo.sql#L39) | WHERE / P04 / COMPARISON | `M_COSTCENTER` : `NVARCHAR(5)` | `@cCostSyar` : `VARCHAR(20)` | `@cCostSyar` → `NVARCHAR` | `M_COSTCENTER != @cCostSyar` |
