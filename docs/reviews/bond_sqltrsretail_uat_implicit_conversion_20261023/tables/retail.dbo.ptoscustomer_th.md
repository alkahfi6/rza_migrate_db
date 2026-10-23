# Retail UAT.dbo.PTOSCustomer_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 12 lokasi source, 3 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-actionflag"></a>
## ActionFlag

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_TH.sql#L30). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001193 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:172](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L172) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_th.ActionFlag` : `CHAR(1)` | `'U'` : `varchar` | `'U'` → `CHAR(1)` | `retail.dbo.ptoscustomer_th.ActionFlag <- 'U'` |
| IC-001209 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:203](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L203) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_th.ActionFlag` : `CHAR(1)` | `'U'` : `varchar` | `'U'` → `CHAR(1)` | `retail.dbo.ptoscustomer_th.ActionFlag <- 'U'` |
| IC-001228 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:224](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L224) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_th.ActionFlag` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.ptoscustomer_th.ActionFlag <- 'I'` |
| IC-007085 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSCustomerAuthen:120](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSCustomerAuthen.sql#L120) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_th.ActionFlag` : `CHAR(1)` | `'D'` : `varchar` | `'D'` → `CHAR(1)` | `retail.dbo.ptoscustomer_th.ActionFlag <- 'D'` |
| IC-007087 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSCustomerAuthen:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSCustomerAuthen.sql#L180) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_th.ActionFlag` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.ptoscustomer_th.ActionFlag <- 'O'` |
| IC-007089 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSCustomerAuthen:244](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSCustomerAuthen.sql#L244) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_th.ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.ptoscustomer_th.ActionFlag <- 'N'` |
| IC-007091 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSCustomerAuthen:312](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSCustomerAuthen.sql#L312) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_th.ActionFlag` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.ptoscustomer_th.ActionFlag <- 'I'` |
| IC-007212 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUploadAuthen:128](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUploadAuthen.sql#L128) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_th.ActionFlag` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.ptoscustomer_th.ActionFlag <- 'I'` |

<a id="kolom-approveflag"></a>
## ApproveFlag

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_TH.sql#L31). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001194 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:172](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L172) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_th.ApproveFlag` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.ptoscustomer_th.ApproveFlag <- 'A'` |
| IC-001210 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:203](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L203) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_th.ApproveFlag` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.ptoscustomer_th.ApproveFlag <- 'A'` |
| IC-001229 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:224](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L224) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_th.ApproveFlag` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.ptoscustomer_th.ApproveFlag <- 'A'` |

<a id="kolom-jeniskelamin"></a>
## JenisKelamin

**Tipe dasar:** `VARCHAR(6)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_TH.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007213 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUploadAuthen:128](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUploadAuthen.sql#L128) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_th.JenisKelamin` : `VARCHAR(6)` | `JenisKelamin` : `char(1)` | `JenisKelamin` → `VARCHAR(6)` | `retail.dbo.ptoscustomer_th.JenisKelamin <- JenisKelamin` |
