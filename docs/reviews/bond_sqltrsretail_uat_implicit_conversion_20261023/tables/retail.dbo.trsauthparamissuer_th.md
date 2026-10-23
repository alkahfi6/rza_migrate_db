# Retail UAT.dbo.TRSAuthParamIssuer_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 7 lokasi source, 4 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-authstatus"></a>
## AuthStatus

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSAuthParamIssuer_TH.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004929 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSObligasiHousekeepProcess:121](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSObligasiHousekeepProcess.sql#L121) | WHERE / P12 / IN_VALUE_TYPE | `mst.AuthStatus` : `CHAR(1)` | `'A'` : `varchar` | `mst.AuthStatus` → `VARCHAR` | `mst.AuthStatus in ('A','R')` |
| IC-004930 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSObligasiHousekeepProcess:123](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSObligasiHousekeepProcess.sql#L123) | WHERE / P12 / IN_VALUE_TYPE | `AuthStatus` : `CHAR(1)` | `'A'` : `varchar` | `AuthStatus` → `VARCHAR` | `AuthStatus in ('A','R')` |
| IC-005002 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamIssuerAuthProcess:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamIssuerAuthProcess.sql#L26) | WHERE / P12 / COMPARISON | `AuthStatus` : `CHAR(1)` | `'P'` : `varchar` | `AuthStatus` → `VARCHAR` | `AuthStatus = 'P'` |
| IC-005010 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamIssuerGetAuthMaster:12](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamIssuerGetAuthMaster.sql#L12) | WHERE / P12 / COMPARISON | `[AuthStatus]` : `CHAR(1)` | `'P'` : `varchar` | `[AuthStatus]` → `VARCHAR` | `[AuthStatus] = 'P'` |
| IC-005024 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamIssuerSave:158](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamIssuerSave.sql#L158) | WHERE / P12 / COMPARISON | `auth.AuthStatus` : `CHAR(1)` | `'P'` : `varchar` | `auth.AuthStatus` → `VARCHAR` | `auth.AuthStatus = 'P'` |
| IC-005026 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSParamIssuerSave:172](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamIssuerSave.sql#L172) | WHERE / P12 / COMPARISON | `auth.AuthStatus` : `CHAR(1)` | `'P'` : `varchar` | `auth.AuthStatus` → `VARCHAR` | `auth.AuthStatus = 'P'` |
| IC-005027 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSParamIssuerSave:181](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSParamIssuerSave.sql#L181) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsauthparamissuer_th.AuthStatus` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trsauthparamissuer_th.AuthStatus <- 'P'` |
