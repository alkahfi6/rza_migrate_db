# Retail UAT.dbo.TRSPendingAccruedInterest_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 8 lokasi source, 4 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `TINYINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSPendingAccruedInterest_TH.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008581 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSResponseBondFromMX:453](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L453) | WHERE / P11 / COMPARISON | `th.[Status]` : `TINYINT` | `0` : `int` | `th.[Status]` → `INT` | `th.[Status] = 0` |
| IC-008587 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSResponseBondFromMX:470](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L470) | WHERE / P11 / COMPARISON | `aih.[Status]` : `TINYINT` | `0` : `int` | `aih.[Status]` → `INT` | `aih.[Status] = 0` |
| IC-013815 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_UpdateStatusSecurityTransactionAfterMurex:297](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransactionAfterMurex.sql#L297) | WHERE / P11 / COMPARISON | `p.[Status]` : `TINYINT` | `0` : `int` | `p.[Status]` → `INT` | `p.[Status] = 0` |
| IC-013836 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_UpdateStatusSecurityTransaction_TT:206](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransaction_TT.sql#L206) | WHERE / P11 / COMPARISON | `p.[Status]` : `TINYINT` | `0` : `int` | `p.[Status]` → `INT` | `p.[Status] = 0` |
| IC-013838 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_UpdateStatusSecurityTransaction_TT:221](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransaction_TT.sql#L221) | WHERE / P11 / COMPARISON | `p.[Status]` : `TINYINT` | `0` : `int` | `p.[Status]` → `INT` | `p.[Status] = 0` |
| IC-008586 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBondFromMX:464](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L464) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[Status]` : `TINYINT` | `3` : `int` | `3` → `TINYINT` | `[Status] = 3` |
| IC-009477 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRetailUpdateAccruedInterest:35](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateAccruedInterest.sql#L35) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trspendingaccruedinterest_th.Status` : `TINYINT` | `0` : `int` | `0` → `TINYINT` | `retail.dbo.trspendingaccruedinterest_th.Status <- 0` |
| IC-013837 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_UpdateStatusSecurityTransaction_TT:209](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransaction_TT.sql#L209) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `p.[Status]` : `TINYINT` | `CASE WHEN @nCurTrxStatus = 2 AND @inpStatus = 3 THEN 0 ELSE @inpStatus END` : `INT (combined result; facets unknown)` | `CASE WHEN @nCurTrxStatus = 2 AND @inpStatus = 3 THEN 0 ELSE @inpStatus END` → `TINYINT` | `p.[Status] = CASE WHEN @nCurTrxStatus = 2 AND @inpStatus = 3 THEN 0 ELSE @inpStatus END` |
