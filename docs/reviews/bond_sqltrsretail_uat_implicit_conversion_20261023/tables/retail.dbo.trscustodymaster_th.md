# Retail UAT.dbo.TRSCustodyMaster_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-periodtype"></a>
## PeriodType

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSCustodyMaster_TH.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003670 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCustodyUploadPopulatePendingAuthMain:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCustodyUploadPopulatePendingAuthMain.sql#L25) | CASE / P12 / COMPARISON | `a.PeriodType` : `CHAR(1)` | `'M'` : `varchar` | `a.PeriodType` → `VARCHAR` | `a.PeriodType = 'M'` |
| IC-003671 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCustodyUploadPopulatePendingAuthMain:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCustodyUploadPopulatePendingAuthMain.sql#L47) | CASE / P12 / COMPARISON | `a.PeriodType` : `CHAR(1)` | `'M'` : `varchar` | `a.PeriodType` → `VARCHAR` | `a.PeriodType = 'M'` |
