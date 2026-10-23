# Retail UAT.dbo.TRSUploadSID_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 18 lokasi source, 3 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-createdby"></a>
## CreatedBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSUploadSID_TH.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005268 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprovalSID:11](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalSID.sql#L11) | JOIN_ON / P11 / COMPARISON | `mst.CreatedBy` : `INT` | `emp.nik` : `DECIMAL (combined result; facets unknown)` | `mst.CreatedBy` → `DECIMAL` | `emp.nik = mst.CreatedBy` |

<a id="kolom-transactionstatus"></a>
## TransactionStatus

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSUploadSID_TH.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 17 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005269 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprovalSID:12](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalSID.sql#L12) | WHERE / P12 / COMPARISON | `mst.TransactionStatus` : `CHAR(1)` | `'A'` : `varchar` | `mst.TransactionStatus` → `VARCHAR` | `mst.TransactionStatus = 'A'` |
| IC-012064 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveSID:60](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveSID.sql#L60) | JOIN_ON / P12 / COMPARISON | `th.TransactionStatus` : `CHAR(1)` | `'A'` : `varchar` | `th.TransactionStatus` → `VARCHAR` | `th.TransactionStatus = 'A'` |
| IC-012067 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveSID:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveSID.sql#L70) | JOIN_ON / P12 / COMPARISON | `th.TransactionStatus` : `CHAR(1)` | `'A'` : `varchar` | `th.TransactionStatus` → `VARCHAR` | `th.TransactionStatus = 'A'` |
| IC-012069 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveSID:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveSID.sql#L77) | JOIN_ON / P12 / COMPARISON | `th.TransactionStatus` : `CHAR(1)` | `'A'` : `varchar` | `th.TransactionStatus` → `VARCHAR` | `th.TransactionStatus = 'A'` |
| IC-012070 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveSID:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveSID.sql#L80) | WHERE / P12 / COMPARISON | `b.TransactionStatus` : `CHAR(1)` | `'A'` : `varchar` | `b.TransactionStatus` → `VARCHAR` | `b.TransactionStatus = 'A'` |
| IC-012073 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveSID:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveSID.sql#L96) | WHERE / P12 / COMPARISON | `TransactionStatus` : `CHAR(1)` | `'A'` : `varchar` | `TransactionStatus` → `VARCHAR` | `TransactionStatus = 'A'` |
| IC-012074 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveSID:100](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveSID.sql#L100) | WHERE / P12 / COMPARISON | `TransactionStatus` : `CHAR(1)` | `'F'` : `varchar` | `TransactionStatus` → `VARCHAR` | `TransactionStatus = 'F'` |
| IC-003346 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApproveSID:102](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveSID.sql#L102) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mst.TransactionStatus` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `mst.TransactionStatus = 'N'` |
| IC-003351 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApproveSID:283](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveSID.sql#L283) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mst.TransactionStatus` : `CHAR(1)` | `'R'` : `varchar` | `'R'` → `CHAR(1)` | `mst.TransactionStatus = 'R'` |
| IC-012061 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveSID:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveSID.sql#L46) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsuploadsid_th.TransactionStatus` : `CHAR(1)` | `'F'` : `varchar` | `'F'` → `CHAR(1)` | `retail.dbo.trsuploadsid_th.TransactionStatus <- 'F'` |
| IC-012062 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveSID:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveSID.sql#L52) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsuploadsid_th.TransactionStatus` : `CHAR(1)` | `'F'` : `varchar` | `'F'` → `CHAR(1)` | `retail.dbo.trsuploadsid_th.TransactionStatus <- 'F'` |
| IC-012063 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveSID:57](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveSID.sql#L57) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsuploadsid_th.TransactionStatus` : `CHAR(1)` | `'F'` : `varchar` | `'F'` → `CHAR(1)` | `retail.dbo.trsuploadsid_th.TransactionStatus <- 'F'` |
| IC-012065 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveSID:62](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveSID.sql#L62) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsuploadsid_th.TransactionStatus` : `CHAR(1)` | `'F'` : `varchar` | `'F'` → `CHAR(1)` | `retail.dbo.trsuploadsid_th.TransactionStatus <- 'F'` |
| IC-012066 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveSID:67](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveSID.sql#L67) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsuploadsid_th.TransactionStatus` : `CHAR(1)` | `'F'` : `varchar` | `'F'` → `CHAR(1)` | `retail.dbo.trsuploadsid_th.TransactionStatus <- 'F'` |
| IC-012068 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveSID:72](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveSID.sql#L72) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsuploadsid_th.TransactionStatus` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.trsuploadsid_th.TransactionStatus <- 'A'` |
| IC-012071 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveSID:82](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveSID.sql#L82) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsuploadsid_th.TransactionStatus` : `CHAR(1)` | `'F'` : `varchar` | `'F'` → `CHAR(1)` | `retail.dbo.trsuploadsid_th.TransactionStatus <- 'F'` |
| IC-012072 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveSID:88](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveSID.sql#L88) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsuploadsid_th.TransactionStatus` : `CHAR(1)` | `'F'` : `varchar` | `'F'` → `CHAR(1)` | `retail.dbo.trsuploadsid_th.TransactionStatus <- 'F'` |
