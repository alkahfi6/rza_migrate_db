# SQL_REPLICATE.dbo.JHDATA — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 1 lokasi source, 1 procedure, 1 kolom. **A:** 1 lokasi pada 1 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-jdbr"></a>
## JDBR

**Tipe dasar:** `VARCHAR(5)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.JHDATA.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004776 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSOBLExpiredKitasRPT:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOBLExpiredKitasRPT.sql#L131) | JOIN_ON / P01 / COMPARISON | `JH.JDBR` : `VARCHAR(5)` | `MOI1.office_id` : `INT` | `JH.JDBR` → `INT` | `JH.JDBR = MOI1.office_id` |
