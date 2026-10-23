# SQL_REPLICATE.dbo.PROOBL_O — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 5 lokasi source, 4 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 5 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-oblrem3"></a>
## OBLREM3

**Tipe dasar:** `VARCHAR(40)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.PROOBL_O.sql#L21). **Existing leading index:** PROOBL_O_NN2&#124;PROOBL_O_NN3.

**Frekuensi source:** 5 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009731 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:630](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L630) | JOIN_ON / P13 / COMPARISON | `o.OBLREM3` : `VARCHAR(40)` | `th.UniqueId` : `UNIQUEIDENTIFIER` | `o.OBLREM3` → `UNIQUEIDENTIFIER` | `th.UniqueId = o.OBLREM3` |
| IC-009735 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:642](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L642) | JOIN_ON / P13 / COMPARISON | `o.OBLREM3` : `VARCHAR(40)` | `th.UniqueId` : `UNIQUEIDENTIFIER` | `o.OBLREM3` → `UNIQUEIDENTIFIER` | `th.UniqueId = o.OBLREM3` |
| IC-011473 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSRptRefundSafeKeepingFee:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptRefundSafeKeepingFee.sql#L44) | JOIN_ON / P13 / COMPARISON | `po.OBLREM3` : `VARCHAR(40)` | `tSafeKeepingFee.UniqueId` : `UNIQUEIDENTIFIER` | `po.OBLREM3` → `UNIQUEIDENTIFIER` | `tSafeKeepingFee.UniqueId = po.OBLREM3` |
| IC-011513 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSRptSettledCashBack:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettledCashBack.sql#L53) | JOIN_ON / P13 / COMPARISON | `po.OBLREM3` : `VARCHAR(40)` | `scblth.UniqueId` : `UNIQUEIDENTIFIER` | `po.OBLREM3` → `UNIQUEIDENTIFIER` | `scblth.UniqueId = po.OBLREM3` |
| IC-011519 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSRptSettledOrder:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettledOrder.sql#L63) | JOIN_ON / P13 / COMPARISON | `po.OBLREM3` : `VARCHAR(40)` | `solth.UniqueId` : `UNIQUEIDENTIFIER` | `po.OBLREM3` → `UNIQUEIDENTIFIER` | `solth.UniqueId = po.OBLREM3` |
