# Retail UAT.dbo.TRSDDMAST_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-acctno"></a>
## ACCTNO

**Tipe dasar:** `DECIMAL(19,0)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSDDMAST_TM.sql#L3). **Existing leading index:** TRSDDMAST_TM_CN1&#124;TRSDDMAST_TM_NN3.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007047 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessOverBookingMateraiBuyBack:128](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql#L128) | WHERE / P04 / COMPARISON | `ACCTNO` : `DECIMAL(19,0)` | `@GNCTWO` : `varchar (50)` | `@GNCTWO` → `DECIMAL` | `ACCTNO = @GNCTWO` |
| IC-007070 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessOverBookingMateraiBuyBack:393](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql#L393) | JOIN_ON / P01 / COMPARISON | `dd.ACCTNO` : `DECIMAL(19,0)` | `bb.NoRekInvestor` : `VARCHAR(100)` | `bb.NoRekInvestor` → `DECIMAL` | `dd.ACCTNO = bb.NoRekInvestor` |
