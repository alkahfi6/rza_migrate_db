# Retail UAT.dbo.ORISBNInvestorRekDana_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 5 lokasi source, 2 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-actioncode"></a>
## ActionCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISBNInvestorRekDana_TH.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002869 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessRekDana:197](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L197) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.orisbninvestorrekdana_th.ActionCode` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.orisbninvestorrekdana_th.ActionCode <- 'O'` |
| IC-002880 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessRekDana:495](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L495) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.orisbninvestorrekdana_th.ActionCode` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.orisbninvestorrekdana_th.ActionCode <- 'I'` |
| IC-002886 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessRekDana:660](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L660) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.orisbninvestorrekdana_th.ActionCode` : `CHAR(1)` | `'U'` : `varchar` | `'U'` → `CHAR(1)` | `retail.dbo.orisbninvestorrekdana_th.ActionCode <- 'U'` |

<a id="kolom-modifyby"></a>
## ModifyBy

**Tipe dasar:** `VARCHAR(200)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISBNInvestorRekDana_TH.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012457 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUntagRekSBNProcessAuth:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUntagRekSBNProcessAuth.sql#L49) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.orisbninvestorrekdana_th.ModifyBy` : `VARCHAR(200)` | `tmp.IssueNIK` : `int` | `tmp.IssueNIK` → `VARCHAR(200)` | `retail.dbo.orisbninvestorrekdana_th.ModifyBy <- tmp.IssueNIK` |
| IC-012458 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUntagRekSBNProcessAuth:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUntagRekSBNProcessAuth.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.orisbninvestorrekdana_th.ModifyBy` : `VARCHAR(200)` | `tmp.IssueNIK` : `int` | `tmp.IssueNIK` → `VARCHAR(200)` | `retail.dbo.orisbninvestorrekdana_th.ModifyBy <- tmp.IssueNIK` |
