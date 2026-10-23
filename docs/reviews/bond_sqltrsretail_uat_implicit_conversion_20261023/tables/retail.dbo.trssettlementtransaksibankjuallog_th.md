# Retail UAT.dbo.TRSSettlementTransaksiBankJualLog_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 15 lokasi source, 5 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementTransaksiBankJualLog_TH.sql#L11). **Existing leading index:** TRSSettlementTransaksiBankJualLog_TH_NN1.

**Frekuensi source:** 15 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002507 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateSettlementTransaksiBankJual:50](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateSettlementTransaksiBankJual.sql#L50) | JOIN_ON / P12 / COMPARISON | `stbjlth.Status` : `CHAR(1)` | `'F'` : `varchar` | `stbjlth.Status` → `VARCHAR` | `stbjlth.Status <> 'F'` |
| IC-002508 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateSettlementTransaksiBankJual:72](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateSettlementTransaksiBankJual.sql#L72) | CASE / P12 / COMPARISON | `stbjlth.[Status]` : `CHAR(1)` | `'P'` : `varchar` | `stbjlth.[Status]` → `VARCHAR` | `stbjlth.[Status] = 'P'` |
| IC-002509 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateSettlementTransaksiBankJual:73](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateSettlementTransaksiBankJual.sql#L73) | CASE / P12 / COMPARISON | `stbjlth.[Status]` : `CHAR(1)` | `'S'` : `varchar` | `stbjlth.[Status]` → `VARCHAR` | `stbjlth.[Status] = 'S'` |
| IC-002510 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateSettlementTransaksiBankJual:84](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateSettlementTransaksiBankJual.sql#L84) | CASE / P12 / COMPARISON | `stbjlth.[Status]` : `CHAR(1)` | `'P'` : `varchar` | `stbjlth.[Status]` → `VARCHAR` | `stbjlth.[Status] = 'P'` |
| IC-002511 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateSettlementTransaksiBankJual:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateSettlementTransaksiBankJual.sql#L85) | CASE / P12 / COMPARISON | `stbjlth.[Status]` : `CHAR(1)` | `'S'` : `varchar` | `stbjlth.[Status]` → `VARCHAR` | `stbjlth.[Status] = 'S'` |
| IC-002520 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateSettlementTransaksiBankJual:139](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateSettlementTransaksiBankJual.sql#L139) | JOIN_ON / P12 / COMPARISON | `stbjlth.Status` : `CHAR(1)` | `'F'` : `varchar` | `stbjlth.Status` → `VARCHAR` | `stbjlth.Status <> 'F'` |
| IC-002521 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateSettlementTransaksiBankJual:161](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateSettlementTransaksiBankJual.sql#L161) | CASE / P12 / COMPARISON | `stbjlth.[Status]` : `CHAR(1)` | `'P'` : `varchar` | `stbjlth.[Status]` → `VARCHAR` | `stbjlth.[Status] = 'P'` |
| IC-002522 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateSettlementTransaksiBankJual:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateSettlementTransaksiBankJual.sql#L162) | CASE / P12 / COMPARISON | `stbjlth.[Status]` : `CHAR(1)` | `'S'` : `varchar` | `stbjlth.[Status]` → `VARCHAR` | `stbjlth.[Status] = 'S'` |
| IC-005296 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateBlokirSwitchingTransaction:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateBlokirSwitchingTransaction.sql#L113) | JOIN_ON / P12 / COMPARISON | `th.[Status]` : `CHAR(1)` | `'S'` : `varchar` | `th.[Status]` → `VARCHAR` | `th.[Status] = 'S'` |
| IC-006381 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementTransaksiBankJual:19](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementTransaksiBankJual.sql#L19) | CASE / P12 / COMPARISON | `stbjlth.[Status]` : `CHAR(1)` | `'P'` : `varchar` | `stbjlth.[Status]` → `VARCHAR` | `stbjlth.[Status] = 'P'` |
| IC-006382 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementTransaksiBankJual:20](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementTransaksiBankJual.sql#L20) | CASE / P12 / COMPARISON | `stbjlth.[Status]` : `CHAR(1)` | `'S'` : `varchar` | `stbjlth.[Status]` → `VARCHAR` | `stbjlth.[Status] = 'S'` |
| IC-006391 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementTransaksiBankJual:83](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementTransaksiBankJual.sql#L83) | JOIN_ON / P12 / COMPARISON | `stbjlth.Status` : `CHAR(1)` | `'F'` : `varchar` | `stbjlth.Status` → `VARCHAR` | `stbjlth.Status <> 'F'` |
| IC-007704 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSettlementTransaksiBankJual:171](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankJual.sql#L171) | JOIN_ON / P12 / COMPARISON | `stbjlth.Status` : `CHAR(1)` | `'F'` : `varchar` | `stbjlth.Status` → `VARCHAR` | `stbjlth.Status <> 'F'` |
| IC-012544 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdatePostingResult:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L106) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'P'` : `varchar` | `Status` → `VARCHAR` | `Status = 'P'` |
| IC-007705 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementTransaksiBankJual:176](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankJual.sql#L176) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementtransaksibankjuallog_th.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trssettlementtransaksibankjuallog_th.Status <- 'P'` |
