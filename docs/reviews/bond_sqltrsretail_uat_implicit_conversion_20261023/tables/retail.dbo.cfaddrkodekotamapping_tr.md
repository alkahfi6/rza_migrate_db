# Retail UAT.dbo.CFADDRKodeKotaMapping_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 7 lokasi source, 7 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cf2dt2"></a>
## CF2DT2

**Tipe dasar:** `CHAR(50)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CFADDRKodeKotaMapping_TR.sql#L2). **Existing leading index:** PK__CFADDRKodeKotaMa__4297D63B.

**Frekuensi source:** 7 lokasi / 7 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000872 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMValidateCustomer:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateCustomer.sql#L192) | WHERE / P12 / COMPARISON | `akmap.CF2DT2` : `CHAR(50)` | `@cCF2DT2` : `varchar(40)` | `akmap.CF2DT2` → `VARCHAR` | `akmap.CF2DT2 = @cCF2DT2` |
| IC-003719 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:362](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L362) | JOIN_ON / P12 / COMPARISON | `akmap.CF2DT2` : `CHAR(50)` | `ta.CF2DT2` : `VARCHAR(40)` | `akmap.CF2DT2` → `VARCHAR` | `ta.CF2DT2 = akmap.CF2DT2` |
| IC-003821 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:340](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L340) | JOIN_ON / P12 / COMPARISON | `akmap.CF2DT2` : `CHAR(50)` | `ta.CF2DT2` : `VARCHAR(40)` | `akmap.CF2DT2` → `VARCHAR` | `ta.CF2DT2 = akmap.CF2DT2` |
| IC-003872 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusi:305](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L305) | JOIN_ON / P12 / COMPARISON | `akmap.CF2DT2` : `CHAR(50)` | `ta.CF2DT2` : `VARCHAR(40)` | `akmap.CF2DT2` → `VARCHAR` | `ta.CF2DT2 = akmap.CF2DT2` |
| IC-003956 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusiOrchid:286](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L286) | JOIN_ON / P12 / COMPARISON | `akmap.CF2DT2` : `CHAR(50)` | `ta.CF2DT2` : `VARCHAR(40)` | `akmap.CF2DT2` → `VARCHAR` | `ta.CF2DT2 = akmap.CF2DT2` |
| IC-009218 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCore:569](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCore.sql#L569) | JOIN_ON / P12 / COMPARISON | `akmap.CF2DT2` : `CHAR(50)` | `ta.CF2DT2` : `VARCHAR(40)` | `akmap.CF2DT2` → `VARCHAR` | `ta.CF2DT2 = akmap.CF2DT2` |
| IC-009387 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailSyncDataCoreWorkingHours:466](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailSyncDataCoreWorkingHours.sql#L466) | JOIN_ON / P12 / COMPARISON | `akmap.CF2DT2` : `CHAR(50)` | `ta.CF2DT2` : `VARCHAR(40)` | `akmap.CF2DT2` → `VARCHAR` | `ta.CF2DT2 = akmap.CF2DT2` |
