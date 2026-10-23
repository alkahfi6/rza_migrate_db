# Retail UAT.dbo.SRBILogSentMailConfirmation_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-haveemail"></a>
## HaveEmail

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SRBILogSentMailConfirmation_TH.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006405 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateStatusEmailKonfirmasiProduct:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateStatusEmailKonfirmasiProduct.sql#L107) | CASE / P11 / COMPARISON | `HaveEmail` : `BIT` | `1` : `int` | `HaveEmail` → `INT` | `HaveEmail = 1` |
| IC-006406 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateStatusEmailKonfirmasiProduct:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateStatusEmailKonfirmasiProduct.sql#L108) | CASE / P11 / COMPARISON | `HaveEmail` : `BIT` | `0` : `int` | `HaveEmail` → `INT` | `HaveEmail = 0` |
