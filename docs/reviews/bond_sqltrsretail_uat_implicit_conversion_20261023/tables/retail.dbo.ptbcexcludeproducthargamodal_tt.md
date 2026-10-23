# Retail UAT.dbo.PTBCExcludeProductHargaModal_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 3 lokasi source, 2 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-secid"></a>
## SecId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTBCExcludeProductHargaModal_TT.sql#L2). **Existing leading index:** PTBCExcludeProductHargaModal_TT_CN1.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004452 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSMaintenanceParamHargaModal:121](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L121) | JOIN_ON / P11 / COMPARISON | `excl.SecId` : `BIGINT` | `phm.SecId` : `INT` | `phm.SecId` → `BIGINT` | `phm.SecId = excl.SecId` |
| IC-004461 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSMaintenanceParamHargaModal:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L165) | JOIN_ON / P11 / COMPARISON | `excl.SecId` : `BIGINT` | `mst.SecId` : `INT` | `mst.SecId` → `BIGINT` | `mst.SecId = excl.SecId` |
| IC-006013 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSecurityParameter:10](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSecurityParameter.sql#L10) | JOIN_ON / P11 / COMPARISON | `ex.SecId` : `BIGINT` | `tm.SecId` : `INT` | `tm.SecId` → `BIGINT` | `tm.SecId = ex.SecId` |
