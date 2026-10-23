# Retail UAT.dbo.TRSParamaterReport_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 2 procedure, 1 kolom. **A:** 2 lokasi pada 2 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-value"></a>
## Value

**Tipe dasar:** `VARCHAR(100)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamaterReport_TR.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 2 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009763 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptDetailTrx:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L43) | WHERE / P01 / COMPARISON | `Value` : `VARCHAR(100)` | `ISNULL(@pnTransactionType,0)` : `INT` | `Value` → `INT` | `Value = ISNULL(@pnTransactionType,0)` |
| IC-011596 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptTransactionDetailWMTA:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L34) | WHERE / P01 / COMPARISON | `Value` : `VARCHAR(100)` | `ISNULL(@pnTransactionType,0)` : `INT` | `Value` → `INT` | `Value = ISNULL(@pnTransactionType,0)` |
