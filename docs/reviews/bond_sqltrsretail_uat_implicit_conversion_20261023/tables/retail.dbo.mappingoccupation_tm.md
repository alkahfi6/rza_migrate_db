# Retail UAT.dbo.MappingOccupation_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 12 lokasi source, 5 procedure, 3 kolom. **A:** 2 lokasi pada 2 procedure. **B:** 4 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-kodeoccupation"></a>
## KodeOccupation

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.MappingOccupation_TM.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 4 SP. **A:** 2 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003730 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:430](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L430) | WHERE / P01 / COMPARISON | `mo.KodeOccupation` : `CHAR(1)` | `5` : `int` | `mo.KodeOccupation` → `INT` | `mo.KodeOccupation = 5` |
| IC-003959 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailInstitusiOrchid:344](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L344) | WHERE / P01 / COMPARISON | `mo.KodeOccupation` : `CHAR(1)` | `5` : `int` | `mo.KodeOccupation` → `INT` | `mo.KodeOccupation = 5` |
| IC-003726 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividu:413](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L413) | CASE / P01 / COMPARISON | `mo.KodeOccupation` : `CHAR(1)` | `5` : `int` | `mo.KodeOccupation` → `INT` | `mo.KodeOccupation = 5` |
| IC-003826 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSDataNasabahRetailIndividuOrchid:400](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L400) | CASE / P01 / COMPARISON | `mo.KodeOccupation` : `CHAR(1)` | `5` : `int` | `mo.KodeOccupation` → `INT` | `mo.KodeOccupation = 5` |
| IC-008523 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportTrxCorporateBOND:242](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L242) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `occ.KodeOccupation` : `CHAR(1)` | `tr.investorOccupation` : `varchar(10)` | `occ.KodeOccupation` → `varchar(10)` | `tr.investorOccupation = occ.KodeOccupation` |
| IC-008524 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSReportTrxCorporateBOND:245](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L245) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `occ.KodeOccupation` : `CHAR(1)` | `'1'` : `varchar` | `occ.KodeOccupation` → `VARCHAR` | `when '1' then occ.Keterangan` |

<a id="kolom-kodeorchid"></a>
## KodeOrchid

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.MappingOccupation_TM.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003824 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailIndividuOrchid:395](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L395) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mo.KodeOrchid` : `INT` | `OCCUPATION` : `VARCHAR(30)` | `mo.KodeOrchid` → `VARCHAR(30)` | `OCCUPATION = mo.KodeOrchid` |

<a id="kolom-kodepekerjaan"></a>
## KodePekerjaan

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.MappingOccupation_TM.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003729 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailIndividu:421](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L421) | JOIN_ON / P12 / COMPARISON | `mo.KodePekerjaan` : `VARCHAR(20)` | `cfz.CFOCCD` : `CHAR(5)` | `cfz.CFOCCD` → `VARCHAR` | `cfz.CFOCCD = mo.KodePekerjaan` |
| IC-003737 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailIndividu:464](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql#L464) | JOIN_ON / P12 / COMPARISON | `mo.KodePekerjaan` : `VARCHAR(20)` | `cfz.CFOCCD` : `CHAR(5)` | `cfz.CFOCCD` → `VARCHAR` | `cfz.CFOCCD = mo.KodePekerjaan` |
| IC-003829 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailIndividuOrchid:408](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L408) | JOIN_ON / P12 / COMPARISON | `mo.KodePekerjaan` : `VARCHAR(20)` | `cfz.CFOCCD` : `CHAR(5)` | `cfz.CFOCCD` → `VARCHAR` | `cfz.CFOCCD = mo.KodePekerjaan` |
| IC-003860 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailInstitusi:245](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusi.sql#L245) | JOIN_ON / P12 / COMPARISON | `mo.KodePekerjaan` : `VARCHAR(20)` | `cfz.CFOCCD` : `CHAR(5)` | `cfz.CFOCCD` → `VARCHAR` | `cfz.CFOCCD = mo.KodePekerjaan` |
| IC-008527 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportTrxCorporateBOND:254](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql#L254) | JOIN_ON / P12 / COMPARISON | `occ.KodePekerjaan` : `VARCHAR(20)` | `data.CFOCCD` : `CHAR(5)` | `data.CFOCCD` → `VARCHAR` | `occ.KodePekerjaan = data.CFOCCD` |
