# Retail UAT.dbo.TRSMapPekerjaan_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 1 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-occp-id"></a>
## OCCP_ID

**Tipe dasar:** `CHAR(5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMapPekerjaan_TR.sql#L4). **Existing leading index:** TRSMapPekerjaan_TR_NN2.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000826 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMValidateCustomer:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L98) | WHERE / P12 / COMPARISON | `OCCP_ID` : `CHAR(5)` | `'LAIN'` : `varchar` | `OCCP_ID` → `VARCHAR` | `OCCP_ID = 'LAIN'` |
