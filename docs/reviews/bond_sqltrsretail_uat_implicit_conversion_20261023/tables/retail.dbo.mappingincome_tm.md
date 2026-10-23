# Retail UAT.dbo.MappingIncome_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 3 lokasi source, 3 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-batasawal"></a>
## BatasAwal

**Tipe dasar:** `DECIMAL(18,0)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.MappingIncome_TM.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003731 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailIndividu:437](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L437) | JOIN_ON / P01 / BETWEEN_BOUND_TYPE | `mi.BatasAwal` : `DECIMAL(18,0)` | `mnc._INCOME_PER_ANNUM` : `VARCHAR(90)` | `mnc._INCOME_PER_ANNUM` → `DECIMAL` | `mnc._INCOME_PER_ANNUM BETWEEN mi.BatasAwal AND mi.BatasAkhir` |
| IC-003962 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailInstitusiOrchid:357](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L357) | JOIN_ON / P01 / BETWEEN_BOUND_TYPE | `mi.BatasAwal` : `DECIMAL(18,0)` | `mnc._INCOME_PER_ANNUM` : `VARCHAR(90)` | `mnc._INCOME_PER_ANNUM` → `DECIMAL` | `mnc._INCOME_PER_ANNUM BETWEEN mi.BatasAwal AND mi.BatasAkhir` |
| IC-008530 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportTrxCorporateBOND:276](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L276) | JOIN_ON / P01 / BETWEEN_BOUND_TYPE | `mi.BatasAwal` : `DECIMAL(18,0)` | `tr.CFECUR` : `char(17)` | `tr.CFECUR` → `DECIMAL` | `tr.CFECUR between mi.BatasAwal and mi.BatasAkhir` |
