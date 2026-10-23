# Retail UAT.dbo.TreasuryCustomerBalance_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 5 lokasi source, 2 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-balance"></a>
## Balance

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerBalance_TH.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-010454 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptOutstandingNasabahBySettDate:67](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingNasabahBySettDate.sql#L67) | WHERE / P11 / COMPARISON | `tcb.Balance` : `MONEY` | `0` : `int` | `0` → `MONEY` | `tcb.Balance > 0` |

<a id="kolom-secid"></a>
## SecId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TreasuryCustomerBalance_TH.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT, COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-010453 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSRptOutstandingNasabahBySettDate:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingNasabahBySettDate.sql#L65) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tcb.SecId` : `BIGINT` | `@pnSecId` : `int` | `tcb.SecId` → `int` | `isnull (@pnSecId, tcb.SecId)` |
| IC-010451 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptOutstandingNasabahBySettDate:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingNasabahBySettDate.sql#L61) | JOIN_ON / P11 / COMPARISON | `tcb.SecId` : `BIGINT` | `sm.SecId` : `INT` | `sm.SecId` → `BIGINT` | `sm.SecId = tcb.SecId` |
| IC-010452 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptOutstandingNasabahBySettDate:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingNasabahBySettDate.sql#L65) | WHERE / P11 / COMPARISON | `tcb.SecId` : `BIGINT` | `isnull (@pnSecId, tcb.SecId)` : `int` | `isnull (@pnSecId, tcb.SecId)` → `BIGINT` | `tcb.SecId = isnull (@pnSecId, tcb.SecId)` |
| IC-011661 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveBalanceHistory:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveBalanceHistory.sql#L29) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.treasurycustomerbalance_th.SecId` : `BIGINT` | `tcb.SecId` : `INT` | `tcb.SecId` → `BIGINT` | `retail.dbo.treasurycustomerbalance_th.SecId <- tcb.SecId` |
