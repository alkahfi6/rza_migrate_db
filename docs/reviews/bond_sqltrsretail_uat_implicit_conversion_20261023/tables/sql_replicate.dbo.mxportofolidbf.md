# SQL_REPLICATE.dbo.MXPortofoliDBF — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 9 lokasi source, 1 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-m-costcenter"></a>
## M_COSTCENTER

**Tipe dasar:** `NVARCHAR(10)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.MXPortofoliDBF.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004271 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSInsertRevenueBondJurnalGL:290](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L290) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `trdr.[M_COSTCENTER]` : `NVARCHAR(10)` | `'-'` : `varchar` | `'-'` → `NVARCHAR(10)` | `ISNULL(trdr.[M_COSTCENTER], '-')` |
| IC-004272 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSInsertRevenueBondJurnalGL:291](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L291) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ta.[M_COSTCENTER]` : `NVARCHAR(10)` | `'-'` : `varchar` | `'-'` → `NVARCHAR(10)` | `ISNULL(ta.[M_COSTCENTER], '-')` |
| IC-004274 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSInsertRevenueBondJurnalGL:294](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L294) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ta.[M_COSTCENTER]` : `NVARCHAR(10)` | `'-'` : `varchar` | `'-'` → `NVARCHAR(10)` | `ISNULL(ta.[M_COSTCENTER], '-')` |
| IC-004275 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSInsertRevenueBondJurnalGL:295](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L295) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `trdr.[M_COSTCENTER]` : `NVARCHAR(10)` | `'-'` : `varchar` | `'-'` → `NVARCHAR(10)` | `ISNULL(trdr.[M_COSTCENTER], '-')` |
| IC-004290 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSInsertRevenueBondJurnalGL:339](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L339) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ta.[M_COSTCENTER]` : `NVARCHAR(10)` | `'-'` : `varchar` | `'-'` → `NVARCHAR(10)` | `ISNULL(ta.[M_COSTCENTER], '-')` |
| IC-004292 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSInsertRevenueBondJurnalGL:352](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L352) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ta.[M_COSTCENTER]` : `NVARCHAR(10)` | `'-'` : `varchar` | `'-'` → `NVARCHAR(10)` | `ISNULL(ta.[M_COSTCENTER], '-')` |

<a id="kolom-m-label"></a>
## M_LABEL

**Tipe dasar:** `NVARCHAR(30)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.MXPortofoliDBF.sql#L2). **Existing leading index:** MXPortofoliDBF_CN1.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004277 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSInsertRevenueBondJurnalGL:308](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L308) | JOIN_ON / P02 / COMPARISON | `trdr.[M_LABEL]` : `NVARCHAR(30)` | `dd.[PortfolioInternal]` : `VARCHAR(30)` | `dd.[PortfolioInternal]` → `NVARCHAR` | `dd.[PortfolioInternal] = trdr.[M_LABEL]` |
| IC-004278 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSInsertRevenueBondJurnalGL:310](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L310) | JOIN_ON / P02 / COMPARISON | `ta.[M_LABEL]` : `NVARCHAR(30)` | `dd.[PortfolioTA]` : `VARCHAR(30)` | `dd.[PortfolioTA]` → `NVARCHAR` | `dd.[PortfolioTA] = ta.[M_LABEL]` |
| IC-004294 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSInsertRevenueBondJurnalGL:365](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L365) | JOIN_ON / P02 / COMPARISON | `ta.[M_LABEL]` : `NVARCHAR(30)` | `dd.[PortfolioTA]` : `VARCHAR(30)` | `dd.[PortfolioTA]` → `NVARCHAR` | `dd.[PortfolioTA] = ta.[M_LABEL]` |
