# Retail UAT.dbo.TRSSecurityTransactionPreBook_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 6 lokasi source, 3 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 3 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-facevalue"></a>
## FaceValue

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSecurityTransactionPreBook_TT.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003471 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSCalculateCorporateBond:57](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateCorporateBond.sql#L57) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sum(ISNULL(-st.FaceValue,0))` : `MONEY (aggregate; facets unknown)` | `@tmp_totalCorpBond_OM.TotalAmount` : `decimal(20,4)` | `sum(ISNULL(-st.FaceValue,0))` → `decimal(20,4)` | `@tmp_totalCorpBond_OM.TotalAmount <- sum(ISNULL(-st.FaceValue,0))` |
| IC-003572 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSCalculateTrxSecCorporate:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateTrxSecCorporate.sql#L33) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `SUM(st.FaceValue)` : `MONEY (aggregate; facets unknown)` | `@tmp_total.TotalAmount` : `decimal(20,4)` | `SUM(st.FaceValue)` → `decimal(20,4)` | `@tmp_total.TotalAmount <- SUM(st.FaceValue)` |
| IC-003472 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSCalculateCorporateBond:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateCorporateBond.sql#L61) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `-st.FaceValue` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(-st.FaceValue,0)` |

<a id="kolom-trxtype"></a>
## TrxType

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSecurityTransactionPreBook_TT.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000719 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.OMSubmitTransactionSekunder:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMSubmitTransactionSekunder.sql#L127) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `TrxType` : `SMALLINT` | `@nTrxType` : `int` | `TrxType` → `int` | `@nTrxType = TrxType` |
| IC-003473 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCalculateCorporateBond:68](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateCorporateBond.sql#L68) | JOIN_ON / P11 / COMPARISON | `st.TrxType` : `SMALLINT` | `3` : `int` | `st.TrxType` → `INT` | `st.TrxType = 3` |
| IC-003573 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCalculateTrxSecCorporate:39](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateTrxSecCorporate.sql#L39) | WHERE / P11 / IN_VALUE_TYPE | `st.TrxType` : `SMALLINT` | `3` : `int` | `st.TrxType` → `INT` | `st.TrxType IN (3, 4)` |
