# Retail UAT.dbo.TRSPajakBeliLog_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 11 lokasi source, 4 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSPajakBeliLog_TH.sql#L11). **Existing leading index:** TRSPajakBeliLog_TH_NN2.

**Frekuensi source:** 11 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002469 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli:258](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli.sql#L258) | CASE / P12 / COMPARISON | `stbjlth.[Status]` : `CHAR(1)` | `'P'` : `varchar` | `stbjlth.[Status]` → `VARCHAR` | `stbjlth.[Status] = 'P'` |
| IC-002470 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli:259](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli.sql#L259) | CASE / P12 / COMPARISON | `stbjlth.[Status]` : `CHAR(1)` | `'S'` : `varchar` | `stbjlth.[Status]` → `VARCHAR` | `stbjlth.[Status] = 'S'` |
| IC-002481 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli:321](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli.sql#L321) | JOIN_ON / P12 / COMPARISON | `stbjlth.Status` : `CHAR(1)` | `'F'` : `varchar` | `stbjlth.Status` → `VARCHAR` | `stbjlth.Status <> 'F'` |
| IC-002502 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli:479](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli.sql#L479) | JOIN_ON / P12 / COMPARISON | `tpblth.Status` : `CHAR(1)` | `'S'` : `varchar` | `tpblth.Status` → `VARCHAR` | `tpblth.Status = 'S'` |
| IC-006345 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementTransaksiBankBeli:164](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementTransaksiBankBeli.sql#L164) | CASE / P12 / COMPARISON | `stbjlth.[Status]` : `CHAR(1)` | `'P'` : `varchar` | `stbjlth.[Status]` → `VARCHAR` | `stbjlth.[Status] = 'P'` |
| IC-006346 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementTransaksiBankBeli:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementTransaksiBankBeli.sql#L165) | CASE / P12 / COMPARISON | `stbjlth.[Status]` : `CHAR(1)` | `'S'` : `varchar` | `stbjlth.[Status]` → `VARCHAR` | `stbjlth.[Status] = 'S'` |
| IC-006357 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementTransaksiBankBeli:234](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementTransaksiBankBeli.sql#L234) | JOIN_ON / P12 / COMPARISON | `stbjlth.Status` : `CHAR(1)` | `'F'` : `varchar` | `stbjlth.Status` → `VARCHAR` | `stbjlth.Status <> 'F'` |
| IC-006378 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementTransaksiBankBeli:344](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementTransaksiBankBeli.sql#L344) | JOIN_ON / P12 / COMPARISON | `tpblth.Status` : `CHAR(1)` | `'S'` : `varchar` | `tpblth.Status` → `VARCHAR` | `tpblth.Status = 'S'` |
| IC-007681 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSettlementTransaksiBankBeliPajak:205](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeliPajak.sql#L205) | JOIN_ON / P12 / COMPARISON | `stbjlth.Status` : `CHAR(1)` | `'F'` : `varchar` | `stbjlth.Status` → `VARCHAR` | `stbjlth.Status <> 'F'` |
| IC-012552 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdatePostingResult:154](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L154) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'P'` : `varchar` | `Status` → `VARCHAR` | `Status = 'P'` |
| IC-007682 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementTransaksiBankBeliPajak:214](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeliPajak.sql#L214) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trspajakbelilog_th.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trspajakbelilog_th.Status <- 'P'` |
