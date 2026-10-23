# Retail UAT.dbo.MappingKewarganegaraan_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 2 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 2 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-citizenship"></a>
## Citizenship

**Tipe dasar:** `VARCHAR(5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.MappingKewarganegaraan_TM.sql#L3). **Existing leading index:** MappingKewarganegaraan_TM_NN1.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003781 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailIndividuOrchid:208](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L208) | JOIN_ON / P12 / COMPARISON | `jhc.Citizenship` : `VARCHAR(5)` | `mnc.NATIONALITYOBL` : `CHAR(10)` | `mnc.NATIONALITYOBL` → `VARCHAR` | `mnc.NATIONALITYOBL = jhc.Citizenship` |
| IC-003936 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDataNasabahRetailInstitusiOrchid:198](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L198) | JOIN_ON / P12 / COMPARISON | `mk.Citizenship` : `VARCHAR(5)` | `mnc.LEGAL_DOMICILE` : `CHAR(10)` | `mnc.LEGAL_DOMICILE` → `VARCHAR` | `mnc.LEGAL_DOMICILE = mk.Citizenship` |

<a id="kolom-kodeorchid"></a>
## KodeOrchid

**Tipe dasar:** `VARCHAR(5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.MappingKewarganegaraan_TM.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003780 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailIndividuOrchid:205](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql#L205) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `jhc.KodeOrchid` : `VARCHAR(5)` | `mnc.NATIONALITY` : `CHAR(10)` | `jhc.KodeOrchid` → `CHAR(10)` | `mnc.NATIONALITY = jhc.KodeOrchid` |
| IC-003935 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSDataNasabahRetailInstitusiOrchid:195](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql#L195) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mk.KodeOrchid` : `VARCHAR(5)` | `mnc.LEGAL_DOMICILE` : `CHAR(10)` | `mk.KodeOrchid` → `CHAR(10)` | `mnc.LEGAL_DOMICILE = mk.KodeOrchid` |
