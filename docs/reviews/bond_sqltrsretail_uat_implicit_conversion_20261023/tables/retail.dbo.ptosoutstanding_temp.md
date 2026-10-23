# Retail UAT.dbo.PTOSOutstanding_Temp — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 23 lokasi source, 7 procedure, 4 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-accruedinterest"></a>
## AccruedInterest

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_Temp.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002376 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:183](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L183) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `potrtmp.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(potrtmp.AccruedInterest, 0)` |
| IC-002384 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:315](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L315) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `potrtmp.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(potrtmp.AccruedInterest, 0)` |
| IC-005835 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUploadAuthen:182](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L182) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `potrtmp.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(potrtmp.AccruedInterest, 0)` |
| IC-005843 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUploadAuthen:314](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L314) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `potrtmp.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(potrtmp.AccruedInterest, 0)` |

<a id="kolom-insertedby"></a>
## InsertedBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_Temp.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002377 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:191](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L191) | JOIN_ON / P01 / COMPARISON | `potrtmp.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `potrtmp.InsertedBy = ndme.EMPLID` |
| IC-002385 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L323) | JOIN_ON / P01 / COMPARISON | `potrtmp.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `potrtmp.InsertedBy = ndme.EMPLID` |
| IC-005760 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSOutsAuthen:42](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSOutsAuthen.sql#L42) | JOIN_ON / P01 / COMPARISON | `ptt.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `ptt.InsertedBy = ndme.EMPLID` |
| IC-005761 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSOutsAuthen:50](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSOutsAuthen.sql#L50) | JOIN_ON / P01 / COMPARISON | `ptt.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `ptt.InsertedBy = ndme.EMPLID` |
| IC-005762 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSOutsAuthen:75](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSOutsAuthen.sql#L75) | JOIN_ON / P01 / COMPARISON | `ptt.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `ptt.InsertedBy = ndme.EMPLID` |
| IC-005764 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSOutsAuthen:84](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSOutsAuthen.sql#L84) | JOIN_ON / P01 / COMPARISON | `ptt.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `ptt.InsertedBy = ndme.EMPLID` |
| IC-005836 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSUploadAuthen:190](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L190) | JOIN_ON / P01 / COMPARISON | `potrtmp.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `potrtmp.InsertedBy = ndme.EMPLID` |
| IC-005844 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSUploadAuthen:322](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L322) | JOIN_ON / P01 / COMPARISON | `potrtmp.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `potrtmp.InsertedBy = ndme.EMPLID` |

<a id="kolom-ptostype"></a>
## PTOSType

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_Temp.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002386 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:324](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L324) | WHERE / P11 / COMPARISON | `potrtmp.PTOSType` : `SMALLINT` | `@pnType` : `INT` | `potrtmp.PTOSType` → `INT` | `potrtmp.PTOSType = @pnType` |
| IC-005763 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSOutsAuthen:76](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSOutsAuthen.sql#L76) | WHERE / P11 / COMPARISON | `ptt.PTOSType` : `SMALLINT` | `@pnType` : `INT` | `ptt.PTOSType` → `INT` | `ptt.PTOSType = @pnType` |
| IC-005765 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSOutsAuthen:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSOutsAuthen.sql#L85) | WHERE / P11 / COMPARISON | `ptt.PTOSType` : `SMALLINT` | `@pnType` : `INT` | `ptt.PTOSType` → `INT` | `ptt.PTOSType = @pnType` |
| IC-005845 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUploadAuthen:323](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L323) | WHERE / P11 / COMPARISON | `potrtmp.PTOSType` : `SMALLINT` | `@pnType` : `INT` | `potrtmp.PTOSType` → `INT` | `potrtmp.PTOSType = @pnType` |
| IC-002337 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L96) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_temp.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `INT` | `tmp.PTOSType` → `SMALLINT` | `retail.dbo.ptosoutstanding_temp.PTOSType <- tmp.PTOSType` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_Temp.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007227 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSUploadAuthen:516](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUploadAuthen.sql#L516) | WHERE / P11 / COMPARISON | `potmp.Status` : `SMALLINT` | `1` : `int` | `potmp.Status` → `INT` | `potmp.Status = 1` |
| IC-007229 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSUploadAuthen:537](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUploadAuthen.sql#L537) | WHERE / P11 / COMPARISON | `potmp.Status` : `SMALLINT` | `1` : `int` | `potmp.Status` → `INT` | `potmp.Status = 1` |
| IC-002338 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L96) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_temp.Status` : `SMALLINT` | `1` : `int` | `1` → `SMALLINT` | `retail.dbo.ptosoutstanding_temp.Status <- 1` |
| IC-007118 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSOuts:152](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSOuts.sql#L152) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_temp.Status` : `SMALLINT` | `@pnStatus` : `INT` | `@pnStatus` → `SMALLINT` | `retail.dbo.ptosoutstanding_temp.Status <- @pnStatus` |
| IC-007197 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUpload:589](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUpload.sql#L589) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_temp.Status` : `SMALLINT` | `1` : `int` | `1` → `SMALLINT` | `retail.dbo.ptosoutstanding_temp.Status <- 1` |
| IC-007201 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUpload:891](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUpload.sql#L891) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_temp.Status` : `SMALLINT` | `1` : `int` | `1` → `SMALLINT` | `retail.dbo.ptosoutstanding_temp.Status <- 1` |
