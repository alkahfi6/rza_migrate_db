# Retail UAT.dbo.PTOSSecurityMaster_Temp — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 20 lokasi source, 3 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-insertedby"></a>
## InsertedBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSSecurityMaster_Temp.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005766 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSSecurityMasterAuthen:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSSecurityMasterAuthen.sql#L17) | JOIN_ON / P01 / COMPARISON | `pstmp.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `pstmp.InsertedBy = ndme.EMPLID` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSSecurityMaster_Temp.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 19 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007144 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSSecurityMasterAuthen:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSSecurityMasterAuthen.sql#L85) | CASE / P11 / COMPARISON | `smptm.Status` : `SMALLINT` | `0` : `int` | `smptm.Status` → `INT` | `smptm.Status = 0` |
| IC-007145 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSSecurityMasterAuthen:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSSecurityMasterAuthen.sql#L89) | CASE / P11 / COMPARISON | `smptm.Status` : `SMALLINT` | `0` : `int` | `smptm.Status` → `INT` | `smptm.Status = 0` |
| IC-007146 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSSecurityMasterAuthen:93](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSSecurityMasterAuthen.sql#L93) | CASE / P11 / COMPARISON | `smptm.Status` : `SMALLINT` | `0` : `int` | `smptm.Status` → `INT` | `smptm.Status = 0` |
| IC-007147 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSSecurityMasterAuthen:97](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSSecurityMasterAuthen.sql#L97) | CASE / P11 / COMPARISON | `smptm.Status` : `SMALLINT` | `0` : `int` | `smptm.Status` → `INT` | `smptm.Status = 0` |
| IC-007148 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSSecurityMasterAuthen:103](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSSecurityMasterAuthen.sql#L103) | CASE / P11 / COMPARISON | `smptm.Status` : `SMALLINT` | `0` : `int` | `smptm.Status` → `INT` | `smptm.Status = 0` |
| IC-007149 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSSecurityMasterAuthen:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSSecurityMasterAuthen.sql#L104) | CASE / P11 / COMPARISON | `smptm.Status` : `SMALLINT` | `1` : `int` | `smptm.Status` → `INT` | `smptm.Status = 1` |
| IC-007150 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSSecurityMasterAuthen:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSSecurityMasterAuthen.sql#L105) | CASE / P11 / COMPARISON | `smptm.Status` : `SMALLINT` | `2` : `int` | `smptm.Status` → `INT` | `smptm.Status = 2` |
| IC-007152 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSSecurityMasterAuthen:144](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSSecurityMasterAuthen.sql#L144) | WHERE / P11 / COMPARISON | `smptm.Status` : `SMALLINT` | `0` : `int` | `smptm.Status` → `INT` | `smptm.Status = 0` |
| IC-007154 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSSecurityMasterAuthen:174](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSSecurityMasterAuthen.sql#L174) | WHERE / P11 / COMPARISON | `smptm.Status` : `SMALLINT` | `0` : `int` | `smptm.Status` → `INT` | `smptm.Status = 0` |
| IC-007155 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSSecurityMasterAuthen:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSSecurityMasterAuthen.sql#L180) | WHERE / P11 / COMPARISON | `smptm.Status` : `SMALLINT` | `0` : `int` | `smptm.Status` → `INT` | `smptm.Status = 0` |
| IC-007157 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSSecurityMasterAuthen:222](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSSecurityMasterAuthen.sql#L222) | WHERE / P11 / COMPARISON | `smptm.Status` : `SMALLINT` | `1` : `int` | `smptm.Status` → `INT` | `smptm.Status = 1` |
| IC-007159 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSSecurityMasterAuthen:258](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSSecurityMasterAuthen.sql#L258) | WHERE / P11 / COMPARISON | `smptm.Status` : `SMALLINT` | `1` : `int` | `smptm.Status` → `INT` | `smptm.Status = 1` |
| IC-007160 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSSecurityMasterAuthen:264](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSSecurityMasterAuthen.sql#L264) | WHERE / P11 / COMPARISON | `smptm.Status` : `SMALLINT` | `1` : `int` | `smptm.Status` → `INT` | `smptm.Status = 1` |
| IC-007162 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSSecurityMasterAuthen:300](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSSecurityMasterAuthen.sql#L300) | WHERE / P11 / COMPARISON | `smptm.Status` : `SMALLINT` | `2` : `int` | `smptm.Status` → `INT` | `smptm.Status = 2` |
| IC-007163 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSSecurityMasterAuthen:308](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSSecurityMasterAuthen.sql#L308) | WHERE / P11 / COMPARISON | `smptm.Status` : `SMALLINT` | `2` : `int` | `smptm.Status` → `INT` | `smptm.Status = 2` |
| IC-007164 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSSecurityMasterAuthen:314](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSSecurityMasterAuthen.sql#L314) | WHERE / P11 / COMPARISON | `smptm.Status` : `SMALLINT` | `2` : `int` | `smptm.Status` → `INT` | `smptm.Status = 2` |
| IC-007250 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTSOSecurityMaster:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTSOSecurityMaster.sql#L66) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptossecuritymaster_temp.Status` : `SMALLINT` | `0` : `int` | `0` → `SMALLINT` | `retail.dbo.ptossecuritymaster_temp.Status <- 0` |
| IC-007251 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTSOSecurityMaster:84](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTSOSecurityMaster.sql#L84) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptossecuritymaster_temp.Status` : `SMALLINT` | `1` : `int` | `1` → `SMALLINT` | `retail.dbo.ptossecuritymaster_temp.Status <- 1` |
| IC-007252 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTSOSecurityMaster:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTSOSecurityMaster.sql#L104) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptossecuritymaster_temp.Status` : `SMALLINT` | `2` : `int` | `2` → `SMALLINT` | `retail.dbo.ptossecuritymaster_temp.Status <- 2` |
