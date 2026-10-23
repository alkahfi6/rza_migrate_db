# Retail UAT.dbo.TRSTransaksiBuyBack_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 6 lokasi source, 3 procedure, 4 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-approvalnik"></a>
## ApprovalNIK

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSTransaksiBuyBack_TT.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003322 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalProcessTransaksiBuyBack:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalProcessTransaksiBuyBack.sql#L22) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ApprovalNIK` : `BIGINT` | `@pnProcessNIK` : `int` | `@pnProcessNIK` → `BIGINT` | `ApprovalNIK = @pnProcessNIK` |
| IC-012075 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveTransaksiBuyBack:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveTransaksiBuyBack.sql#L86) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trstransaksibuyback_tt.ApprovalNIK` : `BIGINT` | `''` : `varchar` | `''` → `BIGINT` | `retail.dbo.trstransaksibuyback_tt.ApprovalNIK <- ''` |

<a id="kolom-ctpid"></a>
## CTPId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSTransaksiBuyBack_TT.sql#L3). **Existing leading index:** ix_TRSTransaksiBuyBack_TT_1.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003653 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSCreateMT541Content:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCreateMT541Content.sql#L63) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CTPId` : `BIGINT` | `@cCTPId` : `int` | `CTPId` → `int` | `@cCTPId=CTPId` |

<a id="kolom-sumbiternik"></a>
## SumbiterNIK

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSTransaksiBuyBack_TT.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003320 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSApprovalProcessTransaksiBuyBack:14](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalProcessTransaksiBuyBack.sql#L14) | WHERE / P11 / COMPARISON | `SumbiterNIK` : `BIGINT` | `@pnProcessNIK` : `int` | `@pnProcessNIK` → `BIGINT` | `SumbiterNIK = @pnProcessNIK` |

<a id="kolom-transaksiid"></a>
## TransaksiID

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSTransaksiBuyBack_TT.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003321 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSApprovalProcessTransaksiBuyBack:14](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalProcessTransaksiBuyBack.sql#L14) | WHERE / P11 / COMPARISON | `TransaksiID` : `BIGINT` | `@pnTemplateId` : `int` | `@pnTemplateId` → `BIGINT` | `TransaksiID = @pnTemplateId` |
| IC-003323 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSApprovalProcessTransaksiBuyBack:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalProcessTransaksiBuyBack.sql#L28) | WHERE / P11 / COMPARISON | `TransaksiID` : `BIGINT` | `@pnTemplateId` : `int` | `@pnTemplateId` → `BIGINT` | `TransaksiID = @pnTemplateId` |
