# Retail UAT.dbo.gwp_process_data_tt — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 6 lokasi source, 3 procedure, 2 kolom. **A:** 3 lokasi pada 3 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-process-key"></a>
## process_key

**Tipe dasar:** `VARCHAR(100)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.gwp_process_data_tt.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 3 lokasi / 3 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002715 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.Rvmp_trs_ListSecurityMaster_TM:188](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L188) | WHERE / P01 / COMPARISON | `g.process_key` : `VARCHAR(100)` | `tm.SecId` : `INT` | `g.process_key` → `INT` | `g.process_key = tm.SecId` |
| IC-006009 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPopulateSecurityMasterAuthorization:181](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSecurityMasterAuthorization.sql#L181) | WHERE / P01 / COMPARISON | `g.process_key` : `VARCHAR(100)` | `smptm.SecId` : `INT` | `g.process_key` → `INT` | `g.process_key = smptm.SecId` |
| IC-013302 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.trs_ListSecurityMasterPending_TM:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L141) | WHERE / P01 / COMPARISON | `g.process_key` : `VARCHAR(100)` | `smp.SecId` : `INT` | `g.process_key` → `INT` | `g.process_key = smp.SecId` |

<a id="kolom-process-action"></a>
## process_action

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.gwp_process_data_tt.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002714 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_trs_ListSecurityMaster_TM:177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L177) | CASE / P12 / COMPARISON | `gwp.process_action` : `CHAR(1)` | `'U'` : `varchar` | `gwp.process_action` → `VARCHAR` | `gwp.process_action = 'U'` |
| IC-006008 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSecurityMasterAuthorization:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSecurityMasterAuthorization.sql#L169) | CASE / P12 / COMPARISON | `gwp.process_action` : `CHAR(1)` | `'U'` : `varchar` | `gwp.process_action` → `VARCHAR` | `gwp.process_action = 'U'` |
| IC-013301 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListSecurityMasterPending_TM:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L130) | CASE / P12 / COMPARISON | `gwp.process_action` : `CHAR(1)` | `'U'` : `varchar` | `gwp.process_action` → `VARCHAR` | `gwp.process_action = 'U'` |
