# Retail UAT.dbo.SecuritySysParameter_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 14 lokasi source, 4 procedure, 1 kolom. **A:** 14 lokasi pada 4 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-paramvalue"></a>
## ParamValue

**Tipe dasar:** `VARCHAR(30)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecuritySysParameter_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 14 lokasi / 4 SP. **A:** 14 lokasi / 4 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005259 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateApprovalFOPTransferAsset:79](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalFOPTransferAsset.sql#L79) | JOIN_ON / P01 / COMPARISON | `par1.ParamValue` : `VARCHAR(30)` | `tm.Purpose` : `INT` | `par1.ParamValue` → `INT` | `tm.Purpose = par1.ParamValue` |
| IC-005262 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateApprovalFOPTransferAsset:144](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalFOPTransferAsset.sql#L144) | JOIN_ON / P01 / COMPARISON | `par1.ParamValue` : `VARCHAR(30)` | `tm.Purpose` : `INT` | `par1.ParamValue` → `INT` | `tm.Purpose = par1.ParamValue` |
| IC-009954 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptFOPandTransferAsset:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L105) | JOIN_ON / P01 / COMPARISON | `par.ParamValue` : `VARCHAR(30)` | `sta.TrxStatus` : `INT` | `par.ParamValue` → `INT` | `sta.TrxStatus = par.ParamValue` |
| IC-009955 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptFOPandTransferAsset:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L108) | JOIN_ON / P01 / COMPARISON | `par2.ParamValue` : `VARCHAR(30)` | `sta.JenisTransaksi` : `INT` | `par2.ParamValue` → `INT` | `sta.JenisTransaksi = par2.ParamValue` |
| IC-009956 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptFOPandTransferAsset:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L111) | JOIN_ON / P01 / COMPARISON | `par3.ParamValue` : `VARCHAR(30)` | `sta.Purpose` : `INT` | `par3.ParamValue` → `INT` | `sta.Purpose = par3.ParamValue` |
| IC-009964 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptFOPandTransferAsset:177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L177) | JOIN_ON / P01 / COMPARISON | `par.ParamValue` : `VARCHAR(30)` | `sta.TrxStatus` : `INT` | `par.ParamValue` → `INT` | `sta.TrxStatus = par.ParamValue` |
| IC-009965 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptFOPandTransferAsset:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L180) | JOIN_ON / P01 / COMPARISON | `par2.ParamValue` : `VARCHAR(30)` | `sta.JenisTransaksi` : `INT` | `par2.ParamValue` → `INT` | `sta.JenisTransaksi = par2.ParamValue` |
| IC-009966 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptFOPandTransferAsset:183](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L183) | JOIN_ON / P01 / COMPARISON | `par3.ParamValue` : `VARCHAR(30)` | `sta.Purpose` : `INT` | `par3.ParamValue` → `INT` | `sta.Purpose = par3.ParamValue` |
| IC-012438 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSTransferAssetDealIdSrc:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransferAssetDealIdSrc.sql#L41) | JOIN_ON / P01 / COMPARISON | `tr1.ParamValue` : `VARCHAR(30)` | `tm.JenisTransaksi` : `INT` | `tr1.ParamValue` → `INT` | `tm.JenisTransaksi = tr1.ParamValue` |
| IC-012440 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSTransferAssetDealIdSrc:76](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransferAssetDealIdSrc.sql#L76) | JOIN_ON / P01 / COMPARISON | `tr1.ParamValue` : `VARCHAR(30)` | `tm.JenisTransaksi` : `INT` | `tr1.ParamValue` → `INT` | `tm.JenisTransaksi = tr1.ParamValue` |
| IC-012442 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSTransferAssetDealIdSrc:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransferAssetDealIdSrc.sql#L107) | JOIN_ON / P01 / COMPARISON | `tr1.ParamValue` : `VARCHAR(30)` | `tm.JenisTransaksi` : `INT` | `tr1.ParamValue` → `INT` | `tm.JenisTransaksi = tr1.ParamValue` |
| IC-012443 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSTransferAssetDealIdSrc:138](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransferAssetDealIdSrc.sql#L138) | JOIN_ON / P01 / COMPARISON | `tr1.ParamValue` : `VARCHAR(30)` | `tm.JenisTransaksi` : `INT` | `tr1.ParamValue` → `INT` | `tm.JenisTransaksi = tr1.ParamValue` |
| IC-012445 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSTransferAssetDealIdSrc:170](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransferAssetDealIdSrc.sql#L170) | JOIN_ON / P01 / COMPARISON | `tr1.ParamValue` : `VARCHAR(30)` | `tm.JenisTransaksi` : `INT` | `tr1.ParamValue` → `INT` | `tm.JenisTransaksi = tr1.ParamValue` |
| IC-013052 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.trs_GetDataTransaksiAktifitas:336](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_GetDataTransaksiAktifitas.sql#L336) | JOIN_ON / P01 / COMPARISON | `par.ParamValue` : `VARCHAR(30)` | `tm.JenisTransaksi` : `INT` | `par.ParamValue` → `INT` | `tm.JenisTransaksi = par.ParamValue` |
