# SQL_REPLICATE.dbo.ProCIFChangeMaster_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 3 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cifnumber"></a>
## CIFNumber

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.ProCIFChangeMaster_TH.sql#L4). **Existing leading index:** ProCIFChangeMaster_TH_NN_4&#124;ProCIFChangeMaster_TH_NN_5.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007897 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRPTMasterNasabahPerubahanObligasi2:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql#L45) | JOIN_ON / P01 / COMPARISON | `a.CIFNumber` : `BIGINT` | `c.CFCIF` : `CHAR(19)` | `c.CFCIF` → `BIGINT` | `c.CFCIF=a.CIFNumber` |
| IC-007898 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRPTMasterNasabahPerubahanObligasi2:57](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql#L57) | JOIN_ON / P01 / COMPARISON | `a.CIFNumber` : `BIGINT` | `c.CFCIF` : `CHAR(19)` | `c.CFCIF` → `BIGINT` | `c.CFCIF=a.CIFNumber` |
| IC-007899 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRPTMasterNasabahPerubahanObligasi2:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql#L69) | JOIN_ON / P01 / COMPARISON | `a.CIFNumber` : `BIGINT` | `c.CFCIF` : `CHAR(19)` | `c.CFCIF` → `BIGINT` | `c.CFCIF=a.CIFNumber` |
