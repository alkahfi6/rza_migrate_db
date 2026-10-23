# Retail UAT.dbo.TrsSyncMasterIdAndFuncTemp — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 7 lokasi source, 1 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cifno"></a>
## CIFNo

**Tipe dasar:** `VARCHAR(19)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TrsSyncMasterIdAndFuncTemp.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012402 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSyncMasterIdAndFunc:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterIdAndFunc.sql#L22) | JOIN_ON / P12 / COMPARISON | `tmp.CIFNo` : `VARCHAR(19)` | `mast.CFCIF` : `CHAR(19)` | `mast.CFCIF` → `VARCHAR` | `tmp.CIFNo = mast.CFCIF` |

<a id="kolom-identitasexpireddate"></a>
## IdentitasExpiredDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TrsSyncMasterIdAndFuncTemp.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012407 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSyncMasterIdAndFunc:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterIdAndFunc.sql#L30) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tmp.IdentitasExpiredDate` : `DATETIME` | `'19450817'` : `varchar` | `'19450817'` → `DATETIME` | `isnull(tmp.IdentitasExpiredDate,'19450817')` |
| IC-012414 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSyncMasterIdAndFunc:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterIdAndFunc.sql#L65) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tmp.IdentitasExpiredDate` : `DATETIME` | `'19450817'` : `varchar` | `'19450817'` → `DATETIME` | `isnull(tmp.IdentitasExpiredDate,'19450817')` |

<a id="kolom-jenisidentitas"></a>
## JenisIdentitas

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TrsSyncMasterIdAndFuncTemp.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN_EXPRESSION, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012403 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSSyncMasterIdAndFunc:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterIdAndFunc.sql#L23) | WHERE / P12 / COMPARISON | `isnull(tmp.JenisIdentitas,'')` : `CHAR(1)` | `CASE WHEN mast.CFSSCD like 'KTP%' THEN '1' WHEN mast.CFSSCD = 'SIM' THEN '2' WHEN mast.CFSSCD = 'PP' THEN '4' ELSE '3' END` : `VARCHAR (combined result; facets unknown)` | `isnull(tmp.JenisIdentitas,'')` → `VARCHAR` | `isnull(tmp.JenisIdentitas,'') <> CASE WHEN mast.CFSSCD like 'KTP%' THEN '1' WHEN mast.CFSSCD = 'SIM' THEN '2' WHEN mast.CFSSCD = 'PP' THEN '4' ELSE '3' END` |
| IC-012396 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSyncMasterIdAndFunc:7](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterIdAndFunc.sql#L7) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tmp.JenisIdentitas` : `CHAR(1)` | `CASE WHEN mast.CFSSCD like 'KTP%' THEN '1' WHEN mast.CFSSCD = 'SIM' THEN '2' WHEN mast.CFSSCD = 'PP' THEN '4' ELSE '3' END` : `VARCHAR (combined result; facets unknown)` | `CASE WHEN mast.CFSSCD like 'KTP%' THEN '1' WHEN mast.CFSSCD = 'SIM' THEN '2' WHEN mast.CFSSCD = 'PP' THEN '4' ELSE '3' END` → `CHAR(1)` | `tmp.JenisIdentitas = CASE WHEN mast.CFSSCD like 'KTP%' THEN '1' WHEN mast.CFSSCD = 'SIM' THEN '2' WHEN mast.CFSSCD = 'PP' THEN '4' ELSE '3' END` |
| IC-012404 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSyncMasterIdAndFunc:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterIdAndFunc.sql#L23) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tmp.JenisIdentitas` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(tmp.JenisIdentitas,'')` |
| IC-012412 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSyncMasterIdAndFunc:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterIdAndFunc.sql#L64) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tmp.JenisIdentitas` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(tmp.JenisIdentitas,'')` |
