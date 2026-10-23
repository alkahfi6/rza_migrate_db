# Retail UAT.dbo.TRSTransactionFeeLogBeli_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 9 lokasi source, 4 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSTransactionFeeLogBeli_TR.sql#L12). **Existing leading index:** TRSTransactionFeeLogBeli_TR_NN2.

**Frekuensi source:** 9 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002484 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli:415](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli.sql#L415) | CASE / P12 / COMPARISON | `stbjlth.[Status]` : `CHAR(1)` | `'P'` : `varchar` | `stbjlth.[Status]` → `VARCHAR` | `stbjlth.[Status] = 'P'` |
| IC-002485 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli:416](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli.sql#L416) | CASE / P12 / COMPARISON | `stbjlth.[Status]` : `CHAR(1)` | `'S'` : `varchar` | `stbjlth.[Status]` → `VARCHAR` | `stbjlth.[Status] = 'S'` |
| IC-002501 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli:476](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateSettlementTransaksiBankBeli.sql#L476) | JOIN_ON / P12 / COMPARISON | `stbjlth.Status` : `CHAR(1)` | `'F'` : `varchar` | `stbjlth.Status` → `VARCHAR` | `stbjlth.Status <> 'F'` |
| IC-006360 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementTransaksiBankBeli:275](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementTransaksiBankBeli.sql#L275) | CASE / P12 / COMPARISON | `stbjlth.[Status]` : `CHAR(1)` | `'P'` : `varchar` | `stbjlth.[Status]` → `VARCHAR` | `stbjlth.[Status] = 'P'` |
| IC-006361 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementTransaksiBankBeli:276](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementTransaksiBankBeli.sql#L276) | CASE / P12 / COMPARISON | `stbjlth.[Status]` : `CHAR(1)` | `'S'` : `varchar` | `stbjlth.[Status]` → `VARCHAR` | `stbjlth.[Status] = 'S'` |
| IC-006377 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementTransaksiBankBeli:341](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementTransaksiBankBeli.sql#L341) | JOIN_ON / P12 / COMPARISON | `stbjlth.Status` : `CHAR(1)` | `'F'` : `varchar` | `stbjlth.Status` → `VARCHAR` | `stbjlth.Status <> 'F'` |
| IC-007657 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSettlementTransaksiBankBeliFee:186](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeliFee.sql#L186) | JOIN_ON / P12 / COMPARISON | `stbjlth.Status` : `CHAR(1)` | `'F'` : `varchar` | `stbjlth.Status` → `VARCHAR` | `stbjlth.Status <> 'F'` |
| IC-012553 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdatePostingResult:160](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L160) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'P'` : `varchar` | `Status` → `VARCHAR` | `Status = 'P'` |
| IC-007658 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementTransaksiBankBeliFee:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankBeliFee.sql#L192) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trstransactionfeelogbeli_tr.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trstransactionfeelogbeli_tr.Status <- 'P'` |
