# Retail UAT.dbo.TRSSettlementBIS4_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 2 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-generatedate"></a>
## GenerateDate

**Tipe dasar:** `VARCHAR(100)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4_TH.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007410 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementBIS4:139](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBIS4.sql#L139) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_th.GenerateDate` : `VARCHAR(100)` | `tmp.GenerateDate` : `datetime` | `tmp.GenerateDate` → `VARCHAR(100)` | `retail.dbo.trssettlementbis4_th.GenerateDate <- tmp.GenerateDate` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementBIS4_TH.sql#L32). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006693 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessApprovalSettlementBIS4:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalSettlementBIS4.sql#L30) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_th.Status` : `CHAR(1)` | `Status` : `VARCHAR(19)` | `Status` → `CHAR(1)` | `retail.dbo.trssettlementbis4_th.Status <- Status` |
| IC-006694 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessApprovalSettlementBIS4:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalSettlementBIS4.sql#L51) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_th.Status` : `CHAR(1)` | `Status` : `VARCHAR(19)` | `Status` → `CHAR(1)` | `retail.dbo.trssettlementbis4_th.Status <- Status` |
| IC-007411 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementBIS4:139](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBIS4.sql#L139) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementbis4_th.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trssettlementbis4_th.Status <- 'P'` |
