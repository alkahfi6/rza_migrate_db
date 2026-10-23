# Retail UAT.dbo.TRSSettlementCashBackLog_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 10 lokasi source, 4 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementCashBackLog_TH.sql#L11). **Existing leading index:** TRSSettlementCashBackLog_TH_NN1.

**Frekuensi source:** 9 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006242 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementCashBack:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementCashBack.sql#L18) | CASE / P12 / COMPARISON | `scblth.[Status]` : `CHAR(1)` | `'P'` : `varchar` | `scblth.[Status]` → `VARCHAR` | `scblth.[Status] = 'P'` |
| IC-006243 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementCashBack:19](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementCashBack.sql#L19) | CASE / P12 / COMPARISON | `scblth.[Status]` : `CHAR(1)` | `'S'` : `varchar` | `scblth.[Status]` → `VARCHAR` | `scblth.[Status] = 'S'` |
| IC-006244 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementCashBack:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementCashBack.sql#L66) | JOIN_ON / P12 / COMPARISON | `scblth.Status` : `CHAR(1)` | `'F'` : `varchar` | `scblth.Status` → `VARCHAR` | `scblth.Status <> 'F'` |
| IC-007464 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSettlementCashBack:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementCashBack.sql#L145) | JOIN_ON / P12 / COMPARISON | `scblth.Status` : `CHAR(1)` | `'F'` : `varchar` | `scblth.Status` → `VARCHAR` | `scblth.Status <> 'F'` |
| IC-011510 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSettledCashBack:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettledCashBack.sql#L29) | CASE / P12 / COMPARISON | `scblth.Status` : `CHAR(1)` | `'P'` : `varchar` | `scblth.Status` → `VARCHAR` | `scblth.Status = 'P'` |
| IC-011511 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSettledCashBack:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettledCashBack.sql#L30) | CASE / P12 / COMPARISON | `scblth.Status` : `CHAR(1)` | `'S'` : `varchar` | `scblth.Status` → `VARCHAR` | `scblth.Status = 'S'` |
| IC-011512 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSettledCashBack:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettledCashBack.sql#L31) | CASE / P12 / COMPARISON | `scblth.Status` : `CHAR(1)` | `'F'` : `varchar` | `scblth.Status` → `VARCHAR` | `scblth.Status = 'F'` |
| IC-012542 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdatePostingResult:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L94) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'P'` : `varchar` | `Status` → `VARCHAR` | `Status = 'P'` |
| IC-007465 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementCashBack:163](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementCashBack.sql#L163) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementcashbacklog_th.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trssettlementcashbacklog_th.Status <- 'P'` |

<a id="kolom-uniqueid"></a>
## UniqueId

**Tipe dasar:** `UNIQUEIDENTIFIER`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementCashBackLog_TH.sql#L10). **Existing leading index:** TRSSettlementCashBackLog_TH_NN2.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-011513 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptSettledCashBack:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettledCashBack.sql#L53) | JOIN_ON / P13 / COMPARISON | `scblth.UniqueId` : `UNIQUEIDENTIFIER` | `po.OBLREM3` : `VARCHAR(40)` | `po.OBLREM3` → `UNIQUEIDENTIFIER` | `scblth.UniqueId = po.OBLREM3` |
