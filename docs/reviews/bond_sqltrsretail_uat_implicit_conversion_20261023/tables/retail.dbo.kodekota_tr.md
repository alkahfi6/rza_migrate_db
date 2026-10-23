# Retail UAT.dbo.KodeKota_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 6 lokasi source, 6 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-nama"></a>
## Nama

**Tipe dasar:** `VARCHAR(50)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.KodeKota_TR.sql#L3). **Existing leading index:** KodeKota_TR_NN2.

**Frekuensi source:** 6 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003717 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailIndividu:352](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L352) | JOIN_ON / P12 / COMPARISON | `kk.Nama` : `VARCHAR(50)` | `cfaddr.CF2DT2` : `CHAR(40)` | `cfaddr.CF2DT2` → `VARCHAR` | `cfaddr.CF2DT2 = kk.Nama` |
| IC-003819 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailIndividuOrchid:330](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L330) | JOIN_ON / P12 / COMPARISON | `kk.Nama` : `VARCHAR(50)` | `cfaddr.CF2DT2` : `CHAR(40)` | `cfaddr.CF2DT2` → `VARCHAR` | `cfaddr.CF2DT2 = kk.Nama` |
| IC-003870 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailInstitusi:295](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L295) | JOIN_ON / P12 / COMPARISON | `kk.Nama` : `VARCHAR(50)` | `cfaddr.CF2DT2` : `CHAR(40)` | `cfaddr.CF2DT2` → `VARCHAR` | `cfaddr.CF2DT2 = kk.Nama` |
| IC-003954 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailInstitusiOrchid:276](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L276) | JOIN_ON / P12 / COMPARISON | `kk.Nama` : `VARCHAR(50)` | `cfaddr.CF2DT2` : `CHAR(40)` | `cfaddr.CF2DT2` → `VARCHAR` | `cfaddr.CF2DT2 = kk.Nama` |
| IC-009215 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailSyncDataCore:554](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L554) | JOIN_ON / P12 / COMPARISON | `kk.Nama` : `VARCHAR(50)` | `ca.CF2DT2` : `CHAR(40)` | `ca.CF2DT2` → `VARCHAR` | `ca.CF2DT2 = kk.Nama` |
| IC-009384 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailSyncDataCoreWorkingHours:451](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L451) | JOIN_ON / P12 / COMPARISON | `kk.Nama` : `VARCHAR(50)` | `ca.CF2DT2` : `CHAR(40)` | `ca.CF2DT2` → `VARCHAR` | `ca.CF2DT2 = kk.Nama` |
