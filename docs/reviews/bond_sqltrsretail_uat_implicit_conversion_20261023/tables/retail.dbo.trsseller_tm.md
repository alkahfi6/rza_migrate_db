# Retail UAT.dbo.TRSSeller_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 29 lokasi source, 8 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-nik"></a>
## Nik

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSeller_TM.sql#L2). **Existing leading index:** TRSSeller_TM_CU1.

**Frekuensi source:** 16 lokasi / 7 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008408 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L31) | JOIN_ON / P11 / COMPARISON | `tsrtm.Nik` : `INT` | `unv.nik` : `DECIMAL (combined result; facets unknown)` | `tsrtm.Nik` → `DECIMAL` | `tsrtm.Nik = unv.nik` |
| IC-008413 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L56) | JOIN_ON / P11 / COMPARISON | `tsrtm.Nik` : `INT` | `unv.nik` : `DECIMAL (combined result; facets unknown)` | `tsrtm.Nik` → `DECIMAL` | `tsrtm.Nik = unv.nik` |
| IC-008418 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L81) | JOIN_ON / P11 / COMPARISON | `tsrtm.Nik` : `INT` | `unv.nik` : `DECIMAL (combined result; facets unknown)` | `tsrtm.Nik` → `DECIMAL` | `tsrtm.Nik = unv.nik` |
| IC-008423 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L107) | JOIN_ON / P11 / COMPARISON | `tsrtm.Nik` : `INT` | `unv.nik` : `DECIMAL (combined result; facets unknown)` | `tsrtm.Nik` → `DECIMAL` | `tsrtm.Nik = unv.nik` |
| IC-008428 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:133](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L133) | JOIN_ON / P11 / COMPARISON | `tsrtm.Nik` : `INT` | `unv.nik` : `DECIMAL (combined result; facets unknown)` | `tsrtm.Nik` → `DECIMAL` | `tsrtm.Nik = unv.nik` |
| IC-008433 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:200](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L200) | JOIN_ON / P11 / COMPARISON | `tsrtm.Nik` : `INT` | `unv.nik` : `DECIMAL (combined result; facets unknown)` | `tsrtm.Nik` → `DECIMAL` | `tsrtm.Nik = unv.nik` |
| IC-008438 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:219](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L219) | JOIN_ON / P11 / COMPARISON | `tsrtm.Nik` : `INT` | `unv.nik` : `DECIMAL (combined result; facets unknown)` | `tsrtm.Nik` → `DECIMAL` | `tsrtm.Nik = unv.nik` |
| IC-008443 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:238](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L238) | JOIN_ON / P11 / COMPARISON | `tsrtm.Nik` : `INT` | `unv.nik` : `DECIMAL (combined result; facets unknown)` | `tsrtm.Nik` → `DECIMAL` | `tsrtm.Nik = unv.nik` |
| IC-008448 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:258](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L258) | JOIN_ON / P11 / COMPARISON | `tsrtm.Nik` : `INT` | `unv.nik` : `DECIMAL (combined result; facets unknown)` | `tsrtm.Nik` → `DECIMAL` | `tsrtm.Nik = unv.nik` |
| IC-008453 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportSeller:278](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql#L278) | JOIN_ON / P11 / COMPARISON | `tsrtm.Nik` : `INT` | `unv.nik` : `DECIMAL (combined result; facets unknown)` | `tsrtm.Nik` → `DECIMAL` | `tsrtm.Nik = unv.nik` |
| IC-008941 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailCalculateFee2:273](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee2.sql#L273) | JOIN_ON / P11 / COMPARISON | `sl.Nik` : `INT` | `tmp.[NIK Seller]` : `BIGINT` | `sl.Nik` → `BIGINT` | `sl.Nik = tmp.[NIK Seller]` |
| IC-002695 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSSearchSellerSecondary:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSearchSellerSecondary.sql#L17) | WHERE / P04 / COMPARISON | `slr.Nik` : `INT` | `LTRIM(RTRIM(@pcCol1))` : `VARCHAR (expression; length unknown)` | `LTRIM(RTRIM(@pcCol1))` → `INT` | `slr.Nik = LTRIM(RTRIM(@pcCol1))` |
| IC-003381 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSBlotterLelangPopulate:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBlotterLelangPopulate.sql#L180) | JOIN_ON / P01 / COMPARISON | `tm.Nik` : `INT` | `td.Seller` : `varchar(100)` | `td.Seller` → `INT` | `td.Seller = tm.Nik` |
| IC-012185 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSearchSeller:15](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchSeller.sql#L15) | WHERE / P04 / COMPARISON | `Nik` : `INT` | `LTRIM(RTRIM(@pcCol1))` : `VARCHAR (expression; length unknown)` | `LTRIM(RTRIM(@pcCol1))` → `INT` | `Nik = LTRIM(RTRIM(@pcCol1))` |
| IC-012186 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSearchSellerNew:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchSellerNew.sql#L17) | WHERE / P04 / COMPARISON | `Nik` : `INT` | `LTRIM(RTRIM(@pcCol1))` : `VARCHAR (expression; length unknown)` | `LTRIM(RTRIM(@pcCol1))` → `INT` | `Nik = LTRIM(RTRIM(@pcCol1))` |
| IC-012187 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSearchSellerSecondary:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchSellerSecondary.sql#L17) | WHERE / P04 / COMPARISON | `slr.Nik` : `INT` | `LTRIM(RTRIM(@pcCol1))` : `VARCHAR (expression; length unknown)` | `LTRIM(RTRIM(@pcCol1))` → `INT` | `slr.Nik = LTRIM(RTRIM(@pcCol1))` |

<a id="kolom-perdanaflag"></a>
## PerdanaFlag

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSeller_TM.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006014 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateSeller:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSeller.sql#L44) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `PerdanaFlag` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(PerdanaFlag, 0)` |

<a id="kolom-sekunderflag"></a>
## SekunderFlag

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSeller_TM.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 12 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002696 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSSearchSellerSecondary:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSearchSellerSecondary.sql#L18) | WHERE / P11 / COMPARISON | `slr.SekunderFlag` : `BIT` | `1` : `int` | `slr.SekunderFlag` → `INT` | `slr.SekunderFlag = 1` |
| IC-002697 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSSearchSellerSecondary:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSearchSellerSecondary.sql#L28) | WHERE / P11 / COMPARISON | `slr.SekunderFlag` : `BIT` | `1` : `int` | `slr.SekunderFlag` → `INT` | `slr.SekunderFlag = 1` |
| IC-002698 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSSearchSellerSecondary:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSearchSellerSecondary.sql#L37) | WHERE / P11 / COMPARISON | `slr.SekunderFlag` : `BIT` | `1` : `int` | `slr.SekunderFlag` → `INT` | `slr.SekunderFlag = 1` |
| IC-002699 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSSearchSellerSecondary:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSearchSellerSecondary.sql#L47) | WHERE / P11 / COMPARISON | `slr.SekunderFlag` : `BIT` | `1` : `int` | `slr.SekunderFlag` → `INT` | `slr.SekunderFlag = 1` |
| IC-002700 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSSearchSellerSecondary:55](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSearchSellerSecondary.sql#L55) | WHERE / P11 / COMPARISON | `slr.SekunderFlag` : `BIT` | `1` : `int` | `slr.SekunderFlag` → `INT` | `slr.SekunderFlag = 1` |
| IC-008942 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailCalculateFee2:274](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee2.sql#L274) | JOIN_ON / P11 / COMPARISON | `sl.SekunderFlag` : `BIT` | `1` : `int` | `sl.SekunderFlag` → `INT` | `sl.SekunderFlag = 1` |
| IC-012188 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchSellerSecondary:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchSellerSecondary.sql#L18) | WHERE / P11 / COMPARISON | `slr.SekunderFlag` : `BIT` | `1` : `int` | `slr.SekunderFlag` → `INT` | `slr.SekunderFlag = 1` |
| IC-012189 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchSellerSecondary:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchSellerSecondary.sql#L28) | WHERE / P11 / COMPARISON | `slr.SekunderFlag` : `BIT` | `1` : `int` | `slr.SekunderFlag` → `INT` | `slr.SekunderFlag = 1` |
| IC-012190 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchSellerSecondary:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchSellerSecondary.sql#L37) | WHERE / P11 / COMPARISON | `slr.SekunderFlag` : `BIT` | `1` : `int` | `slr.SekunderFlag` → `INT` | `slr.SekunderFlag = 1` |
| IC-012191 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchSellerSecondary:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchSellerSecondary.sql#L47) | WHERE / P11 / COMPARISON | `slr.SekunderFlag` : `BIT` | `1` : `int` | `slr.SekunderFlag` → `INT` | `slr.SekunderFlag = 1` |
| IC-012192 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchSellerSecondary:55](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchSellerSecondary.sql#L55) | WHERE / P11 / COMPARISON | `slr.SekunderFlag` : `BIT` | `1` : `int` | `slr.SekunderFlag` → `INT` | `slr.SekunderFlag = 1` |
| IC-006015 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateSeller:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSeller.sql#L47) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `SekunderFlag` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(SekunderFlag,0)` |
