# Retail UAT.dbo.TRSSettlementEarlyLog_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 10 lokasi source, 4 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementEarlyLog_TH.sql#L12). **Existing leading index:** TRSSettlementEarlyLog_TH_NN2.

**Frekuensi source:** 10 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006247 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementEarlyRedem:20](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementEarlyRedem.sql#L20) | CASE / P12 / COMPARISON | `selth.[Status]` : `CHAR(1)` | `'P'` : `varchar` | `selth.[Status]` → `VARCHAR` | `selth.[Status] = 'P'` |
| IC-006248 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementEarlyRedem:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementEarlyRedem.sql#L21) | CASE / P12 / COMPARISON | `selth.[Status]` : `CHAR(1)` | `'S'` : `varchar` | `selth.[Status]` → `VARCHAR` | `selth.[Status] = 'S'` |
| IC-006249 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementEarlyRedem:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementEarlyRedem.sql#L59) | JOIN_ON / P12 / COMPARISON | `selth.Status` : `CHAR(1)` | `'F'` : `varchar` | `selth.Status` → `VARCHAR` | `selth.Status <> 'F'` |
| IC-006265 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementEarlyRedemByStatus:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementEarlyRedemByStatus.sql#L18) | CASE / P12 / COMPARISON | `selth.[Status]` : `CHAR(1)` | `'P'` : `varchar` | `selth.[Status]` → `VARCHAR` | `selth.[Status] = 'P'` |
| IC-006266 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementEarlyRedemByStatus:19](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementEarlyRedemByStatus.sql#L19) | CASE / P12 / COMPARISON | `selth.[Status]` : `CHAR(1)` | `'S'` : `varchar` | `selth.[Status]` → `VARCHAR` | `selth.[Status] = 'S'` |
| IC-006267 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementEarlyRedemByStatus:50](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementEarlyRedemByStatus.sql#L50) | JOIN_ON / P12 / COMPARISON | `selth.Status` : `CHAR(1)` | `'F'` : `varchar` | `selth.Status` → `VARCHAR` | `selth.Status <> 'F'` |
| IC-007479 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSettlementEarlyRedem:157](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementEarlyRedem.sql#L157) | JOIN_ON / P12 / COMPARISON | `selth.Status` : `CHAR(1)` | `'F'` : `varchar` | `selth.Status` → `VARCHAR` | `selth.Status <> 'F'` |
| IC-012555 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdatePostingResult:172](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L172) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'P'` : `varchar` | `Status` → `VARCHAR` | `Status = 'P'` |
| IC-012599 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdatePostingResult:893](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L893) | JOIN_ON / P12 / COMPARISON | `selth.Status` : `CHAR(1)` | `'S'` : `varchar` | `selth.Status` → `VARCHAR` | `selth.Status = 'S'` |
| IC-007480 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementEarlyRedem:184](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementEarlyRedem.sql#L184) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementearlylog_th.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trssettlementearlylog_th.Status <- 'P'` |
