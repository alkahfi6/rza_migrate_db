# Retail UAT.dbo.TRSUntagSBNAccApproval_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 6 lokasi source, 3 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-authstatus"></a>
## AuthStatus

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSUntagSBNAccApproval_TT.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012452 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUntagRekSBNPopulate:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUntagRekSBNPopulate.sql#L52) | JOIN_ON / P12 / COMPARISON | `unt.AuthStatus` : `CHAR(1)` | `'P'` : `varchar` | `unt.AuthStatus` → `VARCHAR` | `unt.AuthStatus = 'P'` |
| IC-012453 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUntagRekSBNPopulate:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUntagRekSBNPopulate.sql#L66) | JOIN_ON / P12 / COMPARISON | `unt.AuthStatus` : `CHAR(1)` | `'P'` : `varchar` | `unt.AuthStatus` → `VARCHAR` | `unt.AuthStatus = 'P'` |
| IC-012455 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUntagRekSBNPopulateApproval:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUntagRekSBNPopulateApproval.sql#L27) | WHERE / P12 / COMPARISON | `app.AuthStatus` : `CHAR(1)` | `'P'` : `varchar` | `app.AuthStatus` → `VARCHAR` | `app.AuthStatus = 'P'` |
| IC-012460 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUntagRekSBNSave:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUntagRekSBNSave.sql#L64) | WHERE / P12 / COMPARISON | `sbn.AuthStatus` : `CHAR(1)` | `'P'` : `varchar` | `sbn.AuthStatus` → `VARCHAR` | `sbn.AuthStatus = 'P'` |
| IC-012461 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUntagRekSBNSave:75](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUntagRekSBNSave.sql#L75) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsuntagsbnaccapproval_tt.AuthStatus` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trsuntagsbnaccapproval_tt.AuthStatus <- 'P'` |

<a id="kolom-issuenik"></a>
## IssueNIK

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSUntagSBNAccApproval_TT.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012454 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUntagRekSBNPopulateApproval:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUntagRekSBNPopulateApproval.sql#L26) | JOIN_ON / P11 / COMPARISON | `app.IssueNIK` : `INT` | `un.nik` : `DECIMAL (combined result; facets unknown)` | `app.IssueNIK` → `DECIMAL` | `app.IssueNIK = un.nik` |
