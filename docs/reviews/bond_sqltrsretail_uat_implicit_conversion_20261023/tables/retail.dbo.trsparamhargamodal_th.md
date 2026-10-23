# Retail UAT.dbo.TRSParamHargaModal_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 12 lokasi source, 5 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-hargamodalbeli"></a>
## HargaModalBeli

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModal_TH.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003425 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSBondsPricePopulate:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPricePopulate.sql#L18) | WHERE / P11 / COMPARISON | `HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `HargaModalBeli > 0` |
| IC-003603 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCmprSecurityNo:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCmprSecurityNo.sql#L30) | WHERE / P11 / COMPARISON | `HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `HargaModalBeli > 0` |
| IC-003620 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCmprSecurityNoChannel:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCmprSecurityNoChannel.sql#L30) | WHERE / P11 / COMPARISON | `HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `HargaModalBeli > 0` |
| IC-003983 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSECBondsPricePopulate:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSECBondsPricePopulate.sql#L21) | WHERE / P11 / COMPARISON | `HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `HargaModalBeli > 0` |
| IC-005939 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateParamSpread:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSpread.sql#L41) | WHERE / P11 / COMPARISON | `HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `HargaModalBeli > 0` |
| IC-005941 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateParamSpread:60](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSpread.sql#L60) | WHERE / P11 / COMPARISON | `a.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `a.HargaModalBeli > 0` |

<a id="kolom-hargamodaljual"></a>
## HargaModalJual

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModal_TH.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003424 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSBondsPricePopulate:12](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPricePopulate.sql#L12) | WHERE / P11 / COMPARISON | `HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `HargaModalJual > 0` |
| IC-003602 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCmprSecurityNo:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCmprSecurityNo.sql#L21) | WHERE / P11 / COMPARISON | `HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `HargaModalJual > 0` |
| IC-003619 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCmprSecurityNoChannel:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCmprSecurityNoChannel.sql#L21) | WHERE / P11 / COMPARISON | `HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `HargaModalJual > 0` |
| IC-003982 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSECBondsPricePopulate:15](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSECBondsPricePopulate.sql#L15) | WHERE / P11 / COMPARISON | `HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `HargaModalJual > 0` |
| IC-005938 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateParamSpread:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSpread.sql#L34) | WHERE / P11 / COMPARISON | `HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `HargaModalJual > 0` |
| IC-005940 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateParamSpread:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSpread.sql#L51) | WHERE / P11 / COMPARISON | `a.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `a.HargaModalJual > 0` |
