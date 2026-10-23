# Retail UAT.dbo.SecurityTrxResultLelang_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 6 lokasi source, 4 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-facevalueresult"></a>
## FaceValueResult

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTrxResultLelang_TT.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003045 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SecurityPopulateResultLelangEmail:172](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateResultLelangEmail.sql#L172) | WHERE / P11 / COMPARISON | `tr.FaceValueResult` : `MONEY` | `0` : `int` | `0` → `MONEY` | `tr.FaceValueResult > 0` |
| IC-003080 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SecurityPopulateResultLelangTA:212](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateResultLelangTA.sql#L212) | CASE / P11 / COMPARISON | `srl.FaceValueResult` : `MONEY` | `0` : `int` | `0` → `MONEY` | `srl.FaceValueResult > 0` |
| IC-003083 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SecurityPopulateResultLelangTA:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateResultLelangTA.sql#L225) | CASE / P11 / COMPARISON | `srl.FaceValueResult` : `MONEY` | `0` : `int` | `0` → `MONEY` | `srl.FaceValueResult > 0` |
| IC-003084 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SecurityPopulateResultLelangTA:227](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityPopulateResultLelangTA.sql#L227) | CASE / P11 / COMPARISON | `srl.FaceValueResult` : `MONEY` | `0` : `int` | `0` → `MONEY` | `srl.FaceValueResult <= 0` |

<a id="kolom-secid"></a>
## SecId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTrxResultLelang_TT.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003156 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SecurityProcessResultLelang:701](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessResultLelang.sql#L701) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytrxresultlelang_tt.SecId` : `INT` | `tt.SecId` : `BIGINT` | `tt.SecId` → `INT` | `retail.dbo.securitytrxresultlelang_tt.SecId <- tt.SecId` |
| IC-003186 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SecurityProcessResultLelangTA:538](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessResultLelangTA.sql#L538) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytrxresultlelang_tt.SecId` : `INT` | `tt.SecId` : `BIGINT` | `tt.SecId` → `INT` | `retail.dbo.securitytrxresultlelang_tt.SecId <- tt.SecId` |
