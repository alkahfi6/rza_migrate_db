# Retail UAT.dbo.TRSProcessTransaksiDetail_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 1 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-processdate"></a>
## ProcessDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSProcessTransaksiDetail_TH.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000297 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.EOD_HousekeepTableTrsRetail:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.EOD_HousekeepTableTrsRetail.sql#L17) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ProcessDate` : `DATETIME` | `'1900-01-01'` : `varchar` | `'1900-01-01'` → `DATETIME` | `ISNULL(ProcessDate, '1900-01-01')` |
