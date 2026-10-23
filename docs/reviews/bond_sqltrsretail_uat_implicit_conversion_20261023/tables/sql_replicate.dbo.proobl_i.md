# SQL_REPLICATE.dbo.PROOBL_I — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 22 lokasi source, 12 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-oblblatp"></a>
## OBLBLATP

**Tipe dasar:** `VARCHAR(1)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.PROOBL_I.sql#L27). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007544 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementOrder:377](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementOrder.sql#L377) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.proobl_i_v.OBLBLATP` : `VARCHAR(1)` | `DeleteBlockAccountType` : `CHAR(1)` | `DeleteBlockAccountType` → `VARCHAR(1)` | `retail.dbo.proobl_i_v.OBLBLATP <- DeleteBlockAccountType` |
| IC-007575 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementTransaksiBankBeli:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L265) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.proobl_i_v.OBLBLATP` : `VARCHAR(1)` | `DeleteBlockAccountType` : `CHAR(1)` | `DeleteBlockAccountType` → `VARCHAR(1)` | `retail.dbo.proobl_i_v.OBLBLATP <- DeleteBlockAccountType` |
| IC-007625 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementTransaksiBankBeli:568](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L568) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.proobl_i_v.OBLBLATP` : `VARCHAR(1)` | `DeleteBlockAccountType` : `CHAR(1)` | `DeleteBlockAccountType` → `VARCHAR(1)` | `retail.dbo.proobl_i_v.OBLBLATP <- DeleteBlockAccountType` |
| IC-007631 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementTransaksiBankBeli:636](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L636) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.proobl_i_v.OBLBLATP` : `VARCHAR(1)` | `DeleteBlockAccountType` : `CHAR(1)` | `DeleteBlockAccountType` → `VARCHAR(1)` | `retail.dbo.proobl_i_v.OBLBLATP <- DeleteBlockAccountType` |
| IC-007708 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementTransaksiBankJual:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankJual.sql#L228) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.proobl_i_v.OBLBLATP` : `VARCHAR(1)` | `DeleteBlockAccountType` : `CHAR(1)` | `DeleteBlockAccountType` → `VARCHAR(1)` | `retail.dbo.proobl_i_v.OBLBLATP <- DeleteBlockAccountType` |

<a id="kolom-obltltxc"></a>
## OBLTLTXC

**Tipe dasar:** `VARCHAR(4)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.PROOBL_I.sql#L24). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 17 lokasi / 12 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006767 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessBuyBack:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBuyBack.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.proobl_i_v.OBLTLTXC` : `VARCHAR(4)` | `7037` : `int` | `7037` → `VARCHAR(4)` | `retail.dbo.proobl_i_v.OBLTLTXC <- 7037` |
| IC-007451 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementBuyBack:348](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBuyBack.sql#L348) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.proobl_i_v.OBLTLTXC` : `VARCHAR(4)` | `7037` : `int` | `7037` → `VARCHAR(4)` | `retail.dbo.proobl_i_v.OBLTLTXC <- 7037` |
| IC-007452 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementBuyBack:408](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBuyBack.sql#L408) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.proobl_i_v.OBLTLTXC` : `VARCHAR(4)` | `7093` : `int` | `7093` → `VARCHAR(4)` | `retail.dbo.proobl_i_v.OBLTLTXC <- 7093` |
| IC-007466 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementCashBack:190](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementCashBack.sql#L190) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.proobl_i_v.OBLTLTXC` : `VARCHAR(4)` | `7037` : `int` | `7037` → `VARCHAR(4)` | `retail.dbo.proobl_i_v.OBLTLTXC <- 7037` |
| IC-007481 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementEarlyRedem:213](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementEarlyRedem.sql#L213) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.proobl_i_v.OBLTLTXC` : `VARCHAR(4)` | `7072` : `int` | `7072` → `VARCHAR(4)` | `retail.dbo.proobl_i_v.OBLTLTXC <- 7072` |
| IC-007500 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementKupon:281](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementKupon.sql#L281) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.proobl_i_v.OBLTLTXC` : `VARCHAR(4)` | `7049` : `int` | `7049` → `VARCHAR(4)` | `retail.dbo.proobl_i_v.OBLTLTXC <- 7049` |
| IC-007518 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementMaturity:238](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementMaturity.sql#L238) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.proobl_i_v.OBLTLTXC` : `VARCHAR(4)` | `7049` : `int` | `7049` → `VARCHAR(4)` | `retail.dbo.proobl_i_v.OBLTLTXC <- 7049` |
| IC-007519 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementMaturity:301](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementMaturity.sql#L301) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.proobl_i_v.OBLTLTXC` : `VARCHAR(4)` | `7093` : `int` | `7093` → `VARCHAR(4)` | `retail.dbo.proobl_i_v.OBLTLTXC <- 7093` |
| IC-007545 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementOrder:377](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementOrder.sql#L377) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.proobl_i_v.OBLTLTXC` : `VARCHAR(4)` | `7037` : `int` | `7037` → `VARCHAR(4)` | `retail.dbo.proobl_i_v.OBLTLTXC <- 7037` |
| IC-007576 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementTransaksiBankBeli:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L265) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.proobl_i_v.OBLTLTXC` : `VARCHAR(4)` | `7072` : `int` | `7072` → `VARCHAR(4)` | `retail.dbo.proobl_i_v.OBLTLTXC <- 7072` |
| IC-007626 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementTransaksiBankBeli:568](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L568) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.proobl_i_v.OBLTLTXC` : `VARCHAR(4)` | `7072` : `int` | `7072` → `VARCHAR(4)` | `retail.dbo.proobl_i_v.OBLTLTXC <- 7072` |
| IC-007632 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementTransaksiBankBeli:636](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L636) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.proobl_i_v.OBLTLTXC` : `VARCHAR(4)` | `7072` : `int` | `7072` → `VARCHAR(4)` | `retail.dbo.proobl_i_v.OBLTLTXC <- 7072` |
| IC-007659 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementTransaksiBankBeliFee:219](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeliFee.sql#L219) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.proobl_i_v.OBLTLTXC` : `VARCHAR(4)` | `7093` : `int` | `7093` → `VARCHAR(4)` | `retail.dbo.proobl_i_v.OBLTLTXC <- 7093` |
| IC-007683 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementTransaksiBankBeliPajak:246](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeliPajak.sql#L246) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.proobl_i_v.OBLTLTXC` : `VARCHAR(4)` | `7093` : `int` | `7093` → `VARCHAR(4)` | `retail.dbo.proobl_i_v.OBLTLTXC <- 7093` |
| IC-007709 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementTransaksiBankJual:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankJual.sql#L228) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.proobl_i_v.OBLTLTXC` : `VARCHAR(4)` | `7037` : `int` | `7037` → `VARCHAR(4)` | `retail.dbo.proobl_i_v.OBLTLTXC <- 7037` |
| IC-007712 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementTransaksiBankJual:295](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankJual.sql#L295) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.proobl_i_v.OBLTLTXC` : `VARCHAR(4)` | `7093` : `int` | `7093` → `VARCHAR(4)` | `retail.dbo.proobl_i_v.OBLTLTXC <- 7093` |
| IC-007746 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessUploadRefund:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessUploadRefund.sql#L165) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.proobl_i_v.OBLTLTXC` : `VARCHAR(4)` | `7037` : `int` | `7037` → `VARCHAR(4)` | `retail.dbo.proobl_i_v.OBLTLTXC <- 7037` |
