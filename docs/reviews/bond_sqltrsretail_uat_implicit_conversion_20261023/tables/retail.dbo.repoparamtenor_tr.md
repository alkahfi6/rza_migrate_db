# Retail UAT.dbo.REPOParamTenor_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 19 lokasi source, 6 procedure, 3 kolom. **A:** 1 lokasi pada 1 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-lastupdatenik"></a>
## LastUpdateNIK

**Tipe dasar:** `VARCHAR(10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOParamTenor_TR.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 3 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001850 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.RepoPopulateParamHargaModal:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoPopulateParamHargaModal.sql#L64) | JOIN_ON / P01 / COMPARISON | `tnr.LastUpdateNIK` : `VARCHAR(10)` | `kry.nik` : `DECIMAL (combined result; facets unknown)` | `tnr.LastUpdateNIK` → `DECIMAL` | `tnr.LastUpdateNIK = kry.nik` |
| IC-001825 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoMaintenanceParamHargaModal:67](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoMaintenanceParamHargaModal.sql#L67) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `LastUpdateNIK` : `VARCHAR(10)` | `70219` : `int` | `70219` → `VARCHAR(10)` | `LastUpdateNIK = 70219` |
| IC-001827 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoMaintenanceParamHargaModal:79](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoMaintenanceParamHargaModal.sql#L79) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `LastUpdateNIK` : `VARCHAR(10)` | `70219` : `int` | `70219` → `VARCHAR(10)` | `LastUpdateNIK = 70219` |
| IC-001919 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSaveParamHargaModal:200](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSaveParamHargaModal.sql#L200) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `LastUpdateNIK` : `VARCHAR(10)` | `@pnNik` : `INT` | `@pnNik` → `VARCHAR(10)` | `LastUpdateNIK = @pnNik` |
| IC-001924 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSaveParamHargaModal:220](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSaveParamHargaModal.sql#L220) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repoparamtenor_tr.LastUpdateNIK` : `VARCHAR(10)` | `@pnNik` : `INT` | `@pnNik` → `VARCHAR(10)` | `retail.dbo.repoparamtenor_tr.LastUpdateNIK <- @pnNik` |
| IC-001926 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSaveParamHargaModal:288](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSaveParamHargaModal.sql#L288) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `LastUpdateNIK` : `VARCHAR(10)` | `@pnNik` : `INT` | `@pnNik` → `VARCHAR(10)` | `LastUpdateNIK = @pnNik` |

<a id="kolom-active"></a>
## Active

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOParamTenor_TR.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001823 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.RepoMaintenanceParamHargaModal:42](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoMaintenanceParamHargaModal.sql#L42) | WHERE / P11 / COMPARISON | `Active` : `BIT` | `1` : `int` | `Active` → `INT` | `Active = 1` |
| IC-001831 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.RepoMaintenanceParamHargaModal:97](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoMaintenanceParamHargaModal.sql#L97) | WHERE / P11 / COMPARISON | `Active` : `BIT` | `1` : `int` | `Active` → `INT` | `Active = 1` |
| IC-001981 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.RepoValidateMaturityProduct:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoValidateMaturityProduct.sql#L28) | WHERE / P11 / COMPARISON | `IsEnable` : `BIT` | `1` : `int` | `IsEnable` → `INT` | `IsEnable = 1` |
| IC-001982 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.RepoValidateParamHargaModal:8](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoValidateParamHargaModal.sql#L8) | WHERE / P11 / COMPARISON | `Active` : `BIT` | `1` : `int` | `Active` → `INT` | `Active = 1` |
| IC-001824 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoMaintenanceParamHargaModal:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoMaintenanceParamHargaModal.sql#L63) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Active` : `BIT` | `0` : `int` | `0` → `BIT` | `Active = 0` |
| IC-001826 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoMaintenanceParamHargaModal:76](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoMaintenanceParamHargaModal.sql#L76) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Active` : `BIT` | `0` : `int` | `0` → `BIT` | `Active = 0` |
| IC-001830 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoMaintenanceParamHargaModal:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoMaintenanceParamHargaModal.sql#L96) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Active` : `BIT` | `0` : `int` | `0` → `BIT` | `Active = 0` |

<a id="kolom-maturitydate"></a>
## MaturityDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOParamTenor_TR.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P03 — Tanggal teks bertemu date/time:** Tentukan kontrak YYYYMMDD/hari/timestamp. Bentuk nilai pencarian dalam tipe kolom atau gunakan rentang native date; jangan konversi kolom seluruh baris tanpa alasan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012195 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSSentEmailHargaRepo:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailHargaRepo.sql#L21) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `MaturityDate` : `DATETIME` | `@MATURITYDATENEW` : `VARCHAR(100)` | `MaturityDate` → `VARCHAR(100)` | `@MATURITYDATENEW = MaturityDate` |
| IC-001917 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RepoSaveParamHargaModal:194](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSaveParamHargaModal.sql#L194) | WHERE / P03 / COMPARISON | `target.MaturityDate` : `DATETIME` | `source.MaturityDate` : `varchar(200)` | `source.MaturityDate` → `DATETIME` | `target.MaturityDate != source.MaturityDate` |
| IC-001918 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSaveParamHargaModal:198](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSaveParamHargaModal.sql#L198) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `MaturityDate` : `DATETIME` | `source.MaturityDate` : `varchar(200)` | `source.MaturityDate` → `DATETIME` | `MaturityDate = source.MaturityDate` |
| IC-001920 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RepoSaveParamHargaModal:204](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSaveParamHargaModal.sql#L204) | WHERE / P03 / COMPARISON | `target.MaturityDate` : `DATETIME` | `source.MaturityDate` : `varchar(200)` | `source.MaturityDate` → `DATETIME` | `target.MaturityDate != source.MaturityDate` |
| IC-001925 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSaveParamHargaModal:220](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSaveParamHargaModal.sql#L220) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repoparamtenor_tr.MaturityDate` : `DATETIME` | `MaturityDate` : `varchar(200)` | `MaturityDate` → `DATETIME` | `retail.dbo.repoparamtenor_tr.MaturityDate <- MaturityDate` |
| IC-001927 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoSaveParamHargaModal:291](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSaveParamHargaModal.sql#L291) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `MaturityDate` : `DATETIME` | `tmp.MaturityDate` : `varchar(200)` | `tmp.MaturityDate` → `DATETIME` | `MaturityDate = tmp.MaturityDate` |
