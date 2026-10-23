# Retail UAT.dbo.TRSParamDeviasiHarga_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-secccy"></a>
## SecCcy

**Tipe dasar:** `CHAR(3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamDeviasiHarga_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013572 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_PengecekanWewenangDanDevasi:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PengecekanWewenangDanDevasi.sql#L89) | WHERE / P12 / COMPARISON | `SecCcy` : `CHAR(3)` | `@pcCurrency` : `varchar(3)` | `SecCcy` → `VARCHAR` | `SecCcy = @pcCurrency` |
| IC-013573 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_PengecekanWewenangDanDevasi:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PengecekanWewenangDanDevasi.sql#L107) | WHERE / P12 / COMPARISON | `SecCcy` : `CHAR(3)` | `@pcCurrency` : `varchar(3)` | `SecCcy` → `VARCHAR` | `SecCcy = @pcCurrency` |
