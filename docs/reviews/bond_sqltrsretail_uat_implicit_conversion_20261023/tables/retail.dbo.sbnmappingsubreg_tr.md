# Retail UAT.dbo.SBNMappingSubReg_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 13 lokasi source, 3 procedure, 2 kolom. **A:** 3 lokasi pada 3 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-idsubreg"></a>
## IdSubReg

**Tipe dasar:** `VARCHAR(15)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SBNMappingSubReg_TR.sql#L3). **Existing leading index:** SBNMappingSubReg_TR_CN1.

**Frekuensi source:** 4 lokasi / 3 SP. **A:** 3 lokasi / 3 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002820 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.SBNSchedPopulateORIOrder:411](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateORIOrder.sql#L411) | JOIN_ON / P01 / COMPARISON | `map.IdSubReg` : `VARCHAR(15)` | `rk.IdSubregistry` : `INT` | `map.IdSubReg` → `INT` | `rk.IdSubregistry = map.IdSubReg` |
| IC-002832 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.SBNSchedPopulateRegInvestor:234](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestor.sql#L234) | JOIN_ON / P01 / COMPARISON | `map.IdSubReg` : `VARCHAR(15)` | `rksb.IdSubregistry` : `INT` | `map.IdSubReg` → `INT` | `map.IdSubReg = rksb.IdSubregistry` |
| IC-004800 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSORIOnlineNewTransaction:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlineNewTransaction.sql#L180) | JOIN_ON / P01 / COMPARISON | `map.IdSubReg` : `VARCHAR(15)` | `rk.IdSubregistry` : `INT` | `map.IdSubReg` → `INT` | `rk.IdSubregistry = map.IdSubReg` |
| IC-004801 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSORIOnlineNewTransaction:189](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlineNewTransaction.sql#L189) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `map.IdSubReg` : `VARCHAR(15)` | `@nIdSubReg` : `INT` | `map.IdSubReg` → `INT` | `@nIdSubReg = map.IdSubReg` |

<a id="kolom-isreksbusingsre"></a>
## IsRekSBUsingSRE

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SBNMappingSubReg_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002819 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedPopulateORIOrder:407](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateORIOrder.sql#L407) | JOIN_ON / P11 / COMPARISON | `map.IsRekSBUsingSRE` : `BIT` | `0` : `int` | `map.IsRekSBUsingSRE` → `INT` | `map.IsRekSBUsingSRE = 0` |
| IC-002828 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.SBNSchedPopulateRegInvestor:208](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestor.sql#L208) | CASE / P11 / COMPARISON | `ISNULL(map.IsRekSBUsingSRE, 0)` : `BIT` | `0` : `int` | `ISNULL(map.IsRekSBUsingSRE, 0)` → `INT` | `ISNULL(map.IsRekSBUsingSRE, 0) = 0` |
| IC-002830 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.SBNSchedPopulateRegInvestor:210](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestor.sql#L210) | CASE / P11 / COMPARISON | `ISNULL(map.IsRekSBUsingSRE, 0)` : `BIT` | `1` : `int` | `ISNULL(map.IsRekSBUsingSRE, 0)` → `INT` | `ISNULL(map.IsRekSBUsingSRE, 0) = 1` |
| IC-004798 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSORIOnlineNewTransaction:176](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlineNewTransaction.sql#L176) | JOIN_ON / P11 / COMPARISON | `map.IsRekSBUsingSRE` : `BIT` | `0` : `int` | `map.IsRekSBUsingSRE` → `INT` | `map.IsRekSBUsingSRE = 0` |
| IC-004799 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSORIOnlineNewTransaction:178](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlineNewTransaction.sql#L178) | JOIN_ON / P11 / COMPARISON | `map.IsRekSBUsingSRE` : `BIT` | `1` : `int` | `map.IsRekSBUsingSRE` → `INT` | `map.IsRekSBUsingSRE = 1` |
| IC-004822 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSORIOnlineNewTransaction:543](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlineNewTransaction.sql#L543) | JOIN_ON / P11 / COMPARISON | `bk.IsRekSBUsingSRE` : `BIT` | `0` : `int` | `bk.IsRekSBUsingSRE` → `INT` | `bk.IsRekSBUsingSRE = 0` |
| IC-004823 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSORIOnlineNewTransaction:545](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlineNewTransaction.sql#L545) | JOIN_ON / P11 / COMPARISON | `bk.IsRekSBUsingSRE` : `BIT` | `1` : `int` | `bk.IsRekSBUsingSRE` → `INT` | `bk.IsRekSBUsingSRE = 1` |
| IC-002829 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SBNSchedPopulateRegInvestor:208](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestor.sql#L208) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `map.IsRekSBUsingSRE` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(map.IsRekSBUsingSRE, 0)` |
| IC-002831 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.SBNSchedPopulateRegInvestor:210](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestor.sql#L210) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `map.IsRekSBUsingSRE` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(map.IsRekSBUsingSRE, 0)` |
