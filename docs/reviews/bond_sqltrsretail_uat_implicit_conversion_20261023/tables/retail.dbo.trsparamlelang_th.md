# Retail UAT.dbo.TRSParamLelang_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 3 lokasi source, 2 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-actioncode"></a>
## ActionCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamLelang_TH.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003388 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSBlotterLelangSummaryPopulate:14](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBlotterLelangSummaryPopulate.sql#L14) | WHERE / P12 / IN_VALUE_TYPE | `ActionCode` : `CHAR(1)` | `'I'` : `varchar` | `ActionCode` → `VARCHAR` | `ActionCode IN ('I','N')` |
| IC-005053 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSParamLelangJobReset:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangJobReset.sql#L26) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamlelang_th.ActionCode` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.trsparamlelang_th.ActionCode <- 'O'` |
| IC-005054 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSParamLelangJobReset:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangJobReset.sql#L49) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamlelang_th.ActionCode` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.trsparamlelang_th.ActionCode <- 'N'` |
