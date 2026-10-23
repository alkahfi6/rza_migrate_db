# Retail UAT.dbo.TipeInvestor_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 4 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cfclas"></a>
## CFCLAS

**Tipe dasar:** `VARCHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TipeInvestor_TM.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003809 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailIndividuOrchid:282](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L282) | JOIN_ON / P12 / COMPARISON | `ti.CFCLAS` : `VARCHAR(1)` | `cfmast.CFCLAS` : `CHAR(1)` | `cfmast.CFCLAS` → `VARCHAR` | `cfmast.CFCLAS = ti.CFCLAS` |
| IC-003857 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailInstitusi:237](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L237) | JOIN_ON / P12 / COMPARISON | `ti.CFCLAS` : `VARCHAR(1)` | `cfmast.CFCLAS` : `CHAR(1)` | `cfmast.CFCLAS` → `VARCHAR` | `cfmast.CFCLAS = ti.CFCLAS` |
| IC-003944 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailInstitusiOrchid:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L228) | JOIN_ON / P12 / COMPARISON | `ti.CFCLAS` : `VARCHAR(1)` | `cfmast.CFCLAS` : `CHAR(1)` | `cfmast.CFCLAS` → `VARCHAR` | `cfmast.CFCLAS = ti.CFCLAS` |
| IC-008480 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportTemplateSID:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTemplateSID.sql#L141) | JOIN_ON / P12 / COMPARISON | `ti.CFCLAS` : `VARCHAR(1)` | `cf.CFCLAS` : `CHAR(1)` | `cf.CFCLAS` → `VARCHAR` | `cf.CFCLAS = ti.CFCLAS` |
