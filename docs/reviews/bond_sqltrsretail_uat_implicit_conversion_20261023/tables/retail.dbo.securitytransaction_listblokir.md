# Retail UAT.dbo.SecurityTransaction_ListBlokir — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-accountblockactype"></a>
## AccountBlockACTYPE

**Tipe dasar:** `CHAR(10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransaction_ListBlokir.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN_EXPRESSION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005225 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateAmendAccruedAuthorization:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateAmendAccruedAuthorization.sql#L21) | CASE / P12 / COMPARISON | `ISNULL(stlb.AccountBlockACTYPE, '')` : `CHAR(10)` | `''` : `varchar` | `ISNULL(stlb.AccountBlockACTYPE, '')` → `VARCHAR` | `ISNULL(stlb.AccountBlockACTYPE, '') <> ''` |
| IC-005226 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateAmendAccruedAuthorization:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateAmendAccruedAuthorization.sql#L21) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `stlb.AccountBlockACTYPE` : `CHAR(10)` | `''` : `varchar` | `''` → `CHAR(10)` | `ISNULL(stlb.AccountBlockACTYPE, '')` |
