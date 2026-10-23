# Retail UAT.dbo.TRSCTP_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 19 lokasi source, 5 procedure, 3 kolom. **A:** 1 lokasi pada 1 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-ctpno"></a>
## CTPNo

**Tipe dasar:** `VARCHAR(30)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSCTP_TM.sql#L3). **Existing leading index:** TRSCTP_TM_NN1&#124;UQ__TRSCTP_T__E9E1B8CCFFCC5E07.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012466 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSUpdateCTPMurex:20](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateCTPMurex.sql#L20) | WHERE / P01 / COMPARISON | `ctp.CTPNo` : `VARCHAR(30)` | `@pnCTPNo` : `int` | `ctp.CTPNo` → `INT` | `ctp.CTPNo = @pnCTPNo` |

<a id="kolom-ctpid"></a>
## CTPId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSCTP_TM.sql#L2). **Existing leading index:** TRSCTP_TM_CN1.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012092 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSaveUploadCTPMurex:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveUploadCTPMurex.sql#L81) | JOIN_ON / P01 / COMPARISON | `ctp.CTPId` : `BIGINT` | `tmp.CTPId` : `VARCHAR(50)` | `tmp.CTPId` → `BIGINT` | `ctp.CTPId != tmp.CTPId` |
| IC-012094 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSaveUploadCTPMurex:87](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveUploadCTPMurex.sql#L87) | JOIN_ON / P01 / COMPARISON | `ctp.CTPId` : `BIGINT` | `tmp.CTPId` : `VARCHAR(50)` | `tmp.CTPId` → `BIGINT` | `tmp.CTPId = ctp.CTPId` |
| IC-012095 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSaveUploadCTPMurex:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveUploadCTPMurex.sql#L89) | WHERE / P01 / COMPARISON | `ctp.CTPId` : `BIGINT` | `tmp.CTPId` : `VARCHAR(50)` | `tmp.CTPId` → `BIGINT` | `ctp.CTPId = tmp.CTPId` |
| IC-012100 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSaveUploadCTPMurex:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveUploadCTPMurex.sql#L110) | JOIN_ON / P01 / COMPARISON | `ctp.CTPId` : `BIGINT` | `tmp.CTPId` : `VARCHAR(50)` | `tmp.CTPId` → `BIGINT` | `ctp.CTPId = tmp.CTPId` |

<a id="kolom-murexno"></a>
## MurexNo

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSCTP_TM.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 14 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003647 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCopyDataTrxRepoFromMx:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxRepoFromMx.sql#L81) | JOIN_ON / P11 / COMPARISON | `ctp.MurexNo` : `BIGINT` | `tmp.M_CREATOR` : `DECIMAL(10, 0)` | `ctp.MurexNo` → `DECIMAL` | `tmp.M_CREATOR=ctp.MurexNo` |
| IC-003649 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCopyDataTrxRepoFromMx:87](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxRepoFromMx.sql#L87) | JOIN_ON / P11 / COMPARISON | `tm.MurexNo` : `BIGINT` | `td.M_CREATOR` : `DECIMAL(10, 0)` | `tm.MurexNo` → `DECIMAL` | `tm.MurexNo = td.M_CREATOR` |
| IC-003650 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCopyDataTrxRepoFromMx:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxRepoFromMx.sql#L89) | JOIN_ON / P11 / COMPARISON | `ctp.MurexNo` : `BIGINT` | `td.M_NB` : `DECIMAL(10, 0)` | `ctp.MurexNo` → `DECIMAL` | `td.M_NB = ctp.MurexNo` |
| IC-000245 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.CustodyReportSettlementInstruction:351](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql#L351) | JOIN_ON / P01 / COMPARISON | `ctp.MurexNo` : `BIGINT` | `rp.NoMurexExternal` : `VARCHAR(20)` | `rp.NoMurexExternal` → `BIGINT` | `rp.NoMurexExternal = ctp.MurexNo` |
| IC-003648 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSCopyDataTrxRepoFromMx:84](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxRepoFromMx.sql#L84) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tm.MurexNo` : `BIGINT` | `td.M_NB` : `DECIMAL(10, 0)` | `td.M_NB` → `BIGINT` | `tm.MurexNo = td.M_NB` |
| IC-008352 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportReportSettlementInstruction:524](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L524) | JOIN_ON / P01 / COMPARISON | `ctp.MurexNo` : `BIGINT` | `tt.NoMurexExternal` : `VARCHAR(20)` | `tt.NoMurexExternal` → `BIGINT` | `tt.NoMurexExternal = ctp.MurexNo` |
| IC-012090 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSaveUploadCTPMurex:78](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveUploadCTPMurex.sql#L78) | JOIN_ON / P01 / COMPARISON | `ctp.MurexNo` : `BIGINT` | `tmp.MurexNo` : `VARCHAR(50)` | `tmp.MurexNo` → `BIGINT` | `ctp.MurexNo = tmp.MurexNo` |
| IC-012091 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSaveUploadCTPMurex:79](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveUploadCTPMurex.sql#L79) | JOIN_ON / P01 / COMPARISON | `ctp.MurexNo` : `BIGINT` | `tmp.MurexNo` : `VARCHAR(50)` | `tmp.MurexNo` → `BIGINT` | `ctp.MurexNo != tmp.MurexNo` |
| IC-012093 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSaveUploadCTPMurex:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveUploadCTPMurex.sql#L81) | JOIN_ON / P01 / COMPARISON | `ctp.MurexNo` : `BIGINT` | `tmp.MurexNo` : `VARCHAR(50)` | `tmp.MurexNo` → `BIGINT` | `ctp.MurexNo = tmp.MurexNo` |
| IC-012096 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSaveUploadCTPMurex:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveUploadCTPMurex.sql#L89) | WHERE / P01 / COMPARISON | `ctp.MurexNo` : `BIGINT` | `tmp.MurexNo` : `VARCHAR(50)` | `tmp.MurexNo` → `BIGINT` | `ctp.MurexNo != tmp.MurexNo` |
| IC-012097 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveUploadCTPMurex:92](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveUploadCTPMurex.sql#L92) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsctp_tm.MurexNo` : `BIGINT` | `tmp.MurexNo` : `VARCHAR(50)` | `tmp.MurexNo` → `BIGINT` | `retail.dbo.trsctp_tm.MurexNo <- tmp.MurexNo` |
| IC-012098 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSaveUploadCTPMurex:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveUploadCTPMurex.sql#L96) | JOIN_ON / P01 / COMPARISON | `ctp.MurexNo` : `BIGINT` | `tmp.MurexNo` : `VARCHAR(50)` | `tmp.MurexNo` → `BIGINT` | `ctp.MurexNo = tmp.MurexNo` |
| IC-012099 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveUploadCTPMurex:102](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveUploadCTPMurex.sql#L102) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ctp.MurexNo` : `BIGINT` | `tmp.MurexNo` : `VARCHAR(50)` | `tmp.MurexNo` → `BIGINT` | `ctp.MurexNo = tmp.MurexNo` |
| IC-012465 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateCTPMurex:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateCTPMurex.sql#L18) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ctp.MurexNo` : `BIGINT` | `@pnMurexNo` : `int` | `@pnMurexNo` → `BIGINT` | `ctp.MurexNo = @pnMurexNo` |
