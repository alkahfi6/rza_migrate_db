# Retail UAT.dbo.EarlyRedemption_Temp — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 1 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-nikreff"></a>
## NikReff

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.EarlyRedemption_Temp.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005380 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateEarlyRedemAuthorization:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEarlyRedemAuthorization.sql#L69) | JOIN_ON / P01 / COMPARISON | `ertemp.NikReff` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `ertemp.NikReff = ndme.EMPLID` |
| IC-005381 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateEarlyRedemAuthorization:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEarlyRedemAuthorization.sql#L107) | JOIN_ON / P01 / COMPARISON | `ertemp.NikReff` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `ertemp.NikReff = ndme.EMPLID` |

<a id="kolom-nominalredemption"></a>
## NominalRedemption

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.EarlyRedemption_Temp.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005378 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateEarlyRedemAuthorization:54](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEarlyRedemAuthorization.sql#L54) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ertemp.NominalRedemption` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(ertemp.NominalRedemption,0)` |

<a id="kolom-outstanding"></a>
## Outstanding

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.EarlyRedemption_Temp.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005379 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateEarlyRedemAuthorization:55](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEarlyRedemAuthorization.sql#L55) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ertemp.Outstanding` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(ertemp.Outstanding,0)` |
