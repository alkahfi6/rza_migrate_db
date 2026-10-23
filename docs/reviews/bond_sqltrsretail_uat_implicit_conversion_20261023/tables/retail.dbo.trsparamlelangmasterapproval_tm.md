# Retail UAT.dbo.TRSParamLelangMasterApproval_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 5 lokasi source, 2 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-apprstatus"></a>
## ApprStatus

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamLelangMasterApproval_TM.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005055 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamLelangPopulateApproval:15](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangPopulateApproval.sql#L15) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ApprStatus` : `CHAR(1)` | `'P'` : `varchar` | `ApprStatus` → `VARCHAR` | `when 'P' then 'Pending'` |
| IC-005056 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamLelangPopulateApproval:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangPopulateApproval.sql#L16) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `ApprStatus` : `CHAR(1)` | `'A'` : `varchar` | `ApprStatus` → `VARCHAR` | `when 'A' then 'Approved'` |
| IC-005057 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamLelangPopulateApproval:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangPopulateApproval.sql#L21) | WHERE / P12 / COMPARISON | `ApprStatus` : `CHAR(1)` | `'P'` : `varchar` | `ApprStatus` → `VARCHAR` | `ApprStatus = 'P'` |
| IC-005121 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamLelangSave:249](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangSave.sql#L249) | WHERE / P12 / COMPARISON | `ma.ApprStatus` : `CHAR(1)` | `'P'` : `varchar` | `ma.ApprStatus` → `VARCHAR` | `ma.ApprStatus = 'P'` |
| IC-005123 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSParamLelangSave:277](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamLelangSave.sql#L277) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamlelangmasterapproval_tm.ApprStatus` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trsparamlelangmasterapproval_tm.ApprStatus <- 'P'` |
