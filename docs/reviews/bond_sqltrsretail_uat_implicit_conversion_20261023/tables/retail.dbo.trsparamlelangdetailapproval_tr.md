# Retail UAT.dbo.TRSParamLelangDetailApproval_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 11 lokasi source, 2 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-actioncode"></a>
## ActionCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamLelangDetailApproval_TR.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 10 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005058 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamLelangPopulateApproval:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangPopulateApproval.sql#L43) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ActionCode` : `CHAR(1)` | `'I'` : `varchar` | `ActionCode` → `VARCHAR` | `when 'I' then 'Insert'` |
| IC-005059 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamLelangPopulateApproval:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangPopulateApproval.sql#L44) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ActionCode` : `CHAR(1)` | `'D'` : `varchar` | `ActionCode` → `VARCHAR` | `when 'D' then 'Delete'` |
| IC-005060 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamLelangPopulateApproval:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangPopulateApproval.sql#L45) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ActionCode` : `CHAR(1)` | `'O'` : `varchar` | `ActionCode` → `VARCHAR` | `when 'O' then 'Update - Old Data'` |
| IC-005061 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamLelangPopulateApproval:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangPopulateApproval.sql#L46) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ActionCode` : `CHAR(1)` | `'N'` : `varchar` | `ActionCode` → `VARCHAR` | `when 'N' then 'Update - New Data'` |
| IC-005062 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamLelangProcessApproval:132](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangProcessApproval.sql#L132) | WHERE / P12 / COMPARISON | `ActionCode` : `CHAR(1)` | `'I'` : `varchar` | `ActionCode` → `VARCHAR` | `ActionCode = 'I'` |
| IC-005063 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamLelangProcessApproval:139](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangProcessApproval.sql#L139) | WHERE / P12 / COMPARISON | `da.ActionCode` : `CHAR(1)` | `'D'` : `varchar` | `da.ActionCode` → `VARCHAR` | `da.ActionCode = 'D'` |
| IC-005064 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamLelangProcessApproval:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangProcessApproval.sql#L147) | WHERE / P12 / COMPARISON | `da.ActionCode` : `CHAR(1)` | `'N'` : `varchar` | `da.ActionCode` → `VARCHAR` | `da.ActionCode = 'N'` |
| IC-005067 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamLelangProcessApproval:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangProcessApproval.sql#L156) | WHERE / P12 / COMPARISON | `da.ActionCode` : `CHAR(1)` | `'N'` : `varchar` | `da.ActionCode` → `VARCHAR` | `da.ActionCode = 'N'` |
| IC-005068 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamLelangProcessApproval:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangProcessApproval.sql#L180) | WHERE / P12 / COMPARISON | `da.ActionCode` : `CHAR(1)` | `'N'` : `varchar` | `da.ActionCode` → `VARCHAR` | `da.ActionCode = 'N'` |
| IC-005072 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamLelangProcessApproval:316](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangProcessApproval.sql#L316) | JOIN_ON / P12 / COMPARISON | `pl.ActionCode` : `CHAR(1)` | `'O'` : `varchar` | `pl.ActionCode` → `VARCHAR` | `pl.ActionCode = 'O'` |

<a id="kolom-settlementdate"></a>
## SettlementDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamLelangDetailApproval_TR.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005065 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSParamLelangProcessApproval:151](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangProcessApproval.sql#L151) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `da.SettlementDate` : `DATETIME` | `0` : `int` | `0` → `DATETIME` | `isnull(da.SettlementDate,0)` |
