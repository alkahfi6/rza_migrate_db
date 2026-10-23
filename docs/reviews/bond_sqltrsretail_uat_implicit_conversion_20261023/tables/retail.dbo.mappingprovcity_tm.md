# Retail UAT.dbo.MappingProvCity_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 1 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-kodecity"></a>
## KodeCity

**Tipe dasar:** `VARCHAR(4)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.MappingProvCity_TM.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008521 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportTrxCorporateBOND:239](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L239) | JOIN_ON / P12 / COMPARISON | `pc.KodeCity` : `VARCHAR(4)` | `tr.KodeKota` : `char(4)` | `tr.KodeKota` → `VARCHAR` | `pc.KodeCity = tr.KodeKota` |

<a id="kolom-kodeprov"></a>
## KodeProv

**Tipe dasar:** `VARCHAR(4)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.MappingProvCity_TM.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008522 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportTrxCorporateBOND:239](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L239) | JOIN_ON / P12 / COMPARISON | `pc.KodeProv` : `VARCHAR(4)` | `tr.KodePropinsi` : `char(4)` | `tr.KodePropinsi` → `VARCHAR` | `pc.KodeProv = tr.KodePropinsi` |
