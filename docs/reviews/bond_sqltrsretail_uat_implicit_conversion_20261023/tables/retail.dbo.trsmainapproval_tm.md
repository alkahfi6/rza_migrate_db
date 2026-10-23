# Retail UAT.dbo.TRSMainApproval_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 8 lokasi source, 5 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-apprid"></a>
## ApprId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMainApproval_TM.sql#L2). **Existing leading index:** TRSMainApproval_TM_CU1.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003979 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSDeleteParamDeviasiTotalSpread:9](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDeleteParamDeviasiTotalSpread.sql#L9) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `max(ApprId)` : `BIGINT` | `1` : `int` | `1` → `BIGINT` | `max(ApprId)+1` |
| IC-011684 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSaveParamDeviasiTotalSpread:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamDeviasiTotalSpread.sql#L59) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `max(ApprId)` : `BIGINT` | `1` : `int` | `1` → `BIGINT` | `max(ApprId)+1` |

<a id="kolom-initnik"></a>
## InitNik

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMainApproval_TM.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006591 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessApprParamSpread:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L41) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `InitNik` : `INT` | `@pnNik` : `bigint` | `InitNik` → `bigint` | `@pnNik = InitNik` |
| IC-003980 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSDeleteParamDeviasiTotalSpread:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDeleteParamDeviasiTotalSpread.sql#L44) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmainapproval_tm.InitNik` : `INT` | `@pnNIK` : `bigint` | `@pnNIK` → `INT` | `retail.dbo.trsmainapproval_tm.InitNik <- @pnNIK` |
| IC-009518 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptAuditTrailParamSpread:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParamSpread.sql#L37) | JOIN_ON / P01 / COMPARISON | `tmatm.InitNik` : `INT` | `nikIns.EMPLID` : `CHAR(11)` | `nikIns.EMPLID` → `INT` | `tmatm.InitNik = nikIns.EMPLID` |
| IC-009519 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptAuditTrailParamSpread:39](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParamSpread.sql#L39) | JOIN_ON / P01 / COMPARISON | `tmatm.InitNik` : `INT` | `nikAut.EMPLID` : `CHAR(11)` | `nikAut.EMPLID` → `INT` | `tmatm.InitNik = nikAut.EMPLID` |
| IC-009527 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptAuditTrailParamWewenangDeviasi:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParamWewenangDeviasi.sql#L47) | JOIN_ON / P01 / COMPARISON | `tmatm.InitNik` : `INT` | `nikIns.EMPLID` : `CHAR(11)` | `nikIns.EMPLID` → `INT` | `tmatm.InitNik = nikIns.EMPLID` |
| IC-009528 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptAuditTrailParamWewenangDeviasi:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParamWewenangDeviasi.sql#L49) | JOIN_ON / P01 / COMPARISON | `tmatm.InitNik` : `INT` | `nikAut.EMPLID` : `CHAR(11)` | `nikAut.EMPLID` → `INT` | `tmatm.InitNik = nikAut.EMPLID` |
