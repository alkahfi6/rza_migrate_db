# Retail UAT.dbo.BONDSDetailTransaksiBuyBack_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 1 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-taxonaccrued"></a>
## TaxOnAccrued

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.BONDSDetailTransaksiBuyBack_TR.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006211 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateSettlementBuyBack:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L137) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `dbb.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(dbb.TaxOnAccrued,0)` |

<a id="kolom-taxoncapitalgain"></a>
## TaxOnCapitalGain

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.BONDSDetailTransaksiBuyBack_TR.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006212 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateSettlementBuyBack:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBuyBack.sql#L137) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `dbb.TaxOnCapitalGain` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(dbb.TaxOnCapitalGain,0)` |
