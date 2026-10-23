# Retail UAT.dbo.coupon_source_tt — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 34 lokasi source, 4 procedure, 5 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-coupon-date"></a>
## coupon_date

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.coupon_source_tt.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002185 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSKuponProcess:112](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L112) | WHERE / P13 / COMPARISON | `cs.coupon_date` : `DATETIME` | `@dWorkingDate` : `DATE` | `@dWorkingDate` → `DATETIME` | `cs.coupon_date >= @dWorkingDate` |
| IC-002186 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSKuponProcess:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L113) | WHERE / P13 / COMPARISON | `cs.coupon_date` : `DATETIME` | `DATEADD(DAY, 1, @dWorkingDate)` : `DATE` | `DATEADD(DAY, 1, @dWorkingDate)` → `DATETIME` | `cs.coupon_date < DATEADD(DAY, 1, @dWorkingDate)` |
| IC-002187 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSKuponProcess:143](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L143) | WHERE / P13 / COMPARISON | `cs.coupon_date` : `DATETIME` | `@dWorkingDate` : `DATE` | `@dWorkingDate` → `DATETIME` | `cs.coupon_date >= @dWorkingDate` |
| IC-002188 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSKuponProcess:144](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L144) | WHERE / P13 / COMPARISON | `cs.coupon_date` : `DATETIME` | `DATEADD(DAY, 1, @dWorkingDate)` : `DATE` | `DATEADD(DAY, 1, @dWorkingDate)` → `DATETIME` | `cs.coupon_date < DATEADD(DAY, 1, @dWorkingDate)` |
| IC-002192 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSKuponProcess:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L169) | WHERE / P13 / COMPARISON | `cs.coupon_date` : `DATETIME` | `@dWorkingDate` : `DATE` | `@dWorkingDate` → `DATETIME` | `cs.coupon_date >= @dWorkingDate` |
| IC-002193 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSKuponProcess:170](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L170) | WHERE / P13 / COMPARISON | `cs.coupon_date` : `DATETIME` | `DATEADD(DAY, 1, @dWorkingDate)` : `DATE` | `DATEADD(DAY, 1, @dWorkingDate)` → `DATETIME` | `cs.coupon_date < DATEADD(DAY, 1, @dWorkingDate)` |

<a id="kolom-match-flag"></a>
## match_flag

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.coupon_source_tt.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 19 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002177 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSKuponCopyFromCustody:102](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponCopyFromCustody.sql#L102) | WHERE / P12 / IN_VALUE_TYPE | `cs.match_flag` : `CHAR(1)` | `'U'` : `varchar` | `cs.match_flag` → `VARCHAR` | `cs.match_flag NOT IN ('U','R')` |
| IC-002179 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSKuponCopyFromCustody:164](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponCopyFromCustody.sql#L164) | WHERE / P12 / IN_VALUE_TYPE | `cs.match_flag` : `CHAR(1)` | `'U'` : `varchar` | `cs.match_flag` → `VARCHAR` | `cs.match_flag in ('U', 'R')` |
| IC-002189 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.Rvmp_TRSKuponProcess:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L145) | WHERE / P12 / IN_VALUE_TYPE | `ISNULL(cs.match_flag, '')` : `CHAR(1)` | `'M'` : `varchar` | `ISNULL(cs.match_flag, '')` → `VARCHAR` | `ISNULL(cs.match_flag, '') NOT IN ('M','C')` |
| IC-002191 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSKuponProcess:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L168) | WHERE / P12 / IN_VALUE_TYPE | `cs.match_flag` : `CHAR(1)` | `'M'` : `varchar` | `cs.match_flag` → `VARCHAR` | `cs.match_flag IN ('M','C')` |
| IC-002195 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSKuponProcess:232](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L232) | WHERE / P12 / IN_VALUE_TYPE | `cs.match_flag` : `CHAR(1)` | `'M'` : `varchar` | `cs.match_flag` → `VARCHAR` | `cs.match_flag IN ('M','C')` |
| IC-005499 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateKuponUpload:587](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUpload.sql#L587) | CASE / P12 / COMPARISON | `ISNULL(cust.match_flag,'')` : `CHAR(1)` | `'A'` : `varchar` | `ISNULL(cust.match_flag,'')` → `VARCHAR` | `ISNULL(cust.match_flag,'')='A'` |
| IC-005501 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateKuponUpload:590](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUpload.sql#L590) | CASE / P12 / COMPARISON | `ISNULL(cust.match_flag,'')` : `CHAR(1)` | `'R'` : `varchar` | `ISNULL(cust.match_flag,'')` → `VARCHAR` | `ISNULL(cust.match_flag,'')='R'` |
| IC-005503 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateKuponUpload:593](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUpload.sql#L593) | CASE / P12 / COMPARISON | `ISNULL(cust.match_flag,'')` : `CHAR(1)` | `'U'` : `varchar` | `ISNULL(cust.match_flag,'')` → `VARCHAR` | `ISNULL(cust.match_flag,'')='U'` |
| IC-005505 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateKuponUpload:596](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUpload.sql#L596) | CASE / P12 / COMPARISON | `ISNULL(cust.match_flag,'')` : `CHAR(1)` | `'C'` : `varchar` | `ISNULL(cust.match_flag,'')` → `VARCHAR` | `ISNULL(cust.match_flag,'')='C'` |
| IC-005507 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateKuponUpload:599](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUpload.sql#L599) | CASE / P12 / COMPARISON | `ISNULL(cust.match_flag,'')` : `CHAR(1)` | `'M'` : `varchar` | `ISNULL(cust.match_flag,'')` → `VARCHAR` | `ISNULL(cust.match_flag,'')='M'` |
| IC-002178 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSKuponCopyFromCustody:158](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponCopyFromCustody.sql#L158) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `cs.match_flag` : `CHAR(1)` | `'C'` : `varchar` | `'C'` → `CHAR(1)` | `cs.match_flag = 'C'` |
| IC-002190 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSKuponProcess:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L145) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cs.match_flag` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `ISNULL(cs.match_flag, '')` |
| IC-005500 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateKuponUpload:587](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUpload.sql#L587) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cust.match_flag` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `ISNULL(cust.match_flag,'')` |
| IC-005502 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateKuponUpload:590](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUpload.sql#L590) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cust.match_flag` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `ISNULL(cust.match_flag,'')` |
| IC-005504 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateKuponUpload:593](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUpload.sql#L593) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cust.match_flag` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `ISNULL(cust.match_flag,'')` |
| IC-005506 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateKuponUpload:596](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUpload.sql#L596) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cust.match_flag` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `ISNULL(cust.match_flag,'')` |
| IC-005508 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateKuponUpload:599](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUpload.sql#L599) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cust.match_flag` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `ISNULL(cust.match_flag,'')` |
| IC-006974 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessKuponUploadAuthorization:197](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessKuponUploadAuthorization.sql#L197) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mst.match_flag` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `mst.match_flag = 'A'` |
| IC-006979 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessKuponUploadAuthorization:292](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessKuponUploadAuthorization.sql#L292) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `mst.match_flag` : `CHAR(1)` | `'R'` : `varchar` | `'R'` → `CHAR(1)` | `mst.match_flag = 'R'` |

<a id="kolom-net-amount"></a>
## net_amount

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.coupon_source_tt.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005493 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateKuponUpload:549](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUpload.sql#L549) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cust.net_amount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(cust.net_amount, 0)` |
| IC-005494 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateKuponUpload:556](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUpload.sql#L556) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cust.net_amount` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(cust.net_amount, 0)` |

<a id="kolom-safe-keeping-actual"></a>
## safe_keeping_actual

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.coupon_source_tt.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005492 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateKuponUpload:544](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUpload.sql#L544) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cust.safe_keeping_actual` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(cust.safe_keeping_actual, 0)` |

<a id="kolom-sec-id"></a>
## sec_id

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.coupon_source_tt.sql#L2). **Existing leading index:** ix_coupon_source_tt_01.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005495 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateKuponUpload:563](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUpload.sql#L563) | JOIN_ON / P11 / COMPARISON | `cust.sec_id` : `BIGINT` | `obl.SecId` : `int` | `obl.SecId` → `BIGINT` | `obl.SecId = cust.sec_id` |
| IC-005509 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateKuponUpload:607](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateKuponUpload.sql#L607) | JOIN_ON / P11 / COMPARISON | `cust.sec_id` : `BIGINT` | `temp.secId` : `int` | `temp.secId` → `BIGINT` | `temp.secId = cust.sec_id` |
| IC-006973 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessKuponUploadAuthorization:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessKuponUploadAuthorization.sql#L192) | JOIN_ON / P11 / COMPARISON | `mst.sec_id` : `BIGINT` | `src.SecId` : `INT` | `src.SecId` → `BIGINT` | `src.SecId = mst.sec_id` |
| IC-006975 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessKuponUploadAuthorization:205](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessKuponUploadAuthorization.sql#L205) | JOIN_ON / P11 / COMPARISON | `mst.sec_id` : `BIGINT` | `src.SecId` : `INT` | `src.SecId` → `BIGINT` | `src.SecId = mst.sec_id` |
| IC-006978 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessKuponUploadAuthorization:287](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessKuponUploadAuthorization.sql#L287) | JOIN_ON / P11 / COMPARISON | `mst.sec_id` : `BIGINT` | `src.SecId` : `INT` | `src.SecId` → `BIGINT` | `src.SecId = mst.sec_id` |
| IC-006980 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessKuponUploadAuthorization:300](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessKuponUploadAuthorization.sql#L300) | JOIN_ON / P11 / COMPARISON | `mst.sec_id` : `BIGINT` | `src.SecId` : `INT` | `src.SecId` → `BIGINT` | `src.SecId = mst.sec_id` |
