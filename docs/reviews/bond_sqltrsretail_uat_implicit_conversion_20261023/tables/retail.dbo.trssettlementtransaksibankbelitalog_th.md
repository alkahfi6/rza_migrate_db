# Retail UAT.dbo.TRSSettlementTransaksiBankBeliTALog_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 14 lokasi source, 4 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementTransaksiBankBeliTALog_TH.sql#L11). **Existing leading index:** TRSSettlementTransaksiBankBeliTALog_TH_NN2.

**Frekuensi source:** 14 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002441 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli:76](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli.sql#L76) | CASE / P12 / COMPARISON | `stbjlthta.[Status]` : `CHAR(1)` | `'P'` : `varchar` | `stbjlthta.[Status]` → `VARCHAR` | `stbjlthta.[Status] = 'P'` |
| IC-002443 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli.sql#L77) | CASE / P12 / COMPARISON | `stbjlthta.[Status]` : `CHAR(1)` | `'S'` : `varchar` | `stbjlthta.[Status]` → `VARCHAR` | `stbjlthta.[Status] = 'S'` |
| IC-002445 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli:78](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli.sql#L78) | CASE / P12 / COMPARISON | `stbjlthta.[Status]` : `CHAR(1)` | `'P'` : `varchar` | `stbjlthta.[Status]` → `VARCHAR` | `stbjlthta.[Status] = 'P'` |
| IC-002448 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli.sql#L80) | CASE / P12 / COMPARISON | `stbjlthta.[Status]` : `CHAR(1)` | `'S'` : `varchar` | `stbjlthta.[Status]` → `VARCHAR` | `stbjlthta.[Status] = 'S'` |
| IC-002465 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli:144](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli.sql#L144) | JOIN_ON / P12 / COMPARISON | `stbjlthta.Status` : `CHAR(1)` | `'F'` : `varchar` | `stbjlthta.Status` → `VARCHAR` | `stbjlthta.Status <> 'F'` |
| IC-006317 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementTransaksiBankBeli:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementTransaksiBankBeli.sql#L23) | CASE / P12 / COMPARISON | `stbjlthta.[Status]` : `CHAR(1)` | `'P'` : `varchar` | `stbjlthta.[Status]` → `VARCHAR` | `stbjlthta.[Status] = 'P'` |
| IC-006319 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementTransaksiBankBeli:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementTransaksiBankBeli.sql#L24) | CASE / P12 / COMPARISON | `stbjlthta.[Status]` : `CHAR(1)` | `'S'` : `varchar` | `stbjlthta.[Status]` → `VARCHAR` | `stbjlthta.[Status] = 'S'` |
| IC-006321 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementTransaksiBankBeli:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementTransaksiBankBeli.sql#L25) | CASE / P12 / COMPARISON | `stbjlthta.[Status]` : `CHAR(1)` | `'P'` : `varchar` | `stbjlthta.[Status]` → `VARCHAR` | `stbjlthta.[Status] = 'P'` |
| IC-006324 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementTransaksiBankBeli:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementTransaksiBankBeli.sql#L27) | CASE / P12 / COMPARISON | `stbjlthta.[Status]` : `CHAR(1)` | `'S'` : `varchar` | `stbjlthta.[Status]` → `VARCHAR` | `stbjlthta.[Status] = 'S'` |
| IC-006341 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementTransaksiBankBeli:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementTransaksiBankBeli.sql#L96) | JOIN_ON / P12 / COMPARISON | `stbjlthta.Status` : `CHAR(1)` | `'F'` : `varchar` | `stbjlthta.Status` → `VARCHAR` | `stbjlthta.Status <> 'F'` |
| IC-007555 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSettlementTransaksiBankBeli:140](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L140) | JOIN_ON / P12 / COMPARISON | `stbjlthta.Status` : `CHAR(1)` | `'F'` : `varchar` | `stbjlthta.Status` → `VARCHAR` | `stbjlthta.Status <> 'F'` |
| IC-007617 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSettlementTransaksiBankBeli:473](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L473) | CASE / P12 / COMPARISON | `tstlbbta.Status` : `CHAR(1)` | `'F'` : `varchar` | `tstlbbta.Status` → `VARCHAR` | `tstlbbta.Status = 'F'` |
| IC-012554 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdatePostingResult:166](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L166) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'P'` : `varchar` | `Status` → `VARCHAR` | `Status = 'P'` |
| IC-007622 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementTransaksiBankBeli:541](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeli.sql#L541) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementtransaksibankbelitalog_th.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trssettlementtransaksibankbelitalog_th.Status <- 'P'` |
