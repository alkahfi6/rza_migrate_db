# Retail UAT.dbo.TRSSecurityCurrency_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 11 lokasi source, 6 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-ccycode"></a>
## CcyCode

**Tipe dasar:** `VARCHAR(4)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSecurityCurrency_TR.sql#L3). **Existing leading index:** PK_TRSSecurityCurrency_TR.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000521 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetAccountSyariah:102](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountSyariah.sql#L102) | JOIN_ON / P12 / COMPARISON | `b.CcyCode` : `VARCHAR(4)` | `a.CurrencyCode` : `char(3)` | `a.CurrencyCode` → `VARCHAR` | `a.CurrencyCode = b.CcyCode` |
| IC-000523 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetAccountSyariah:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountSyariah.sql#L108) | JOIN_ON / P12 / COMPARISON | `b.CcyCode` : `VARCHAR(4)` | `a.CurrencyCode` : `char(3)` | `a.CurrencyCode` → `VARCHAR` | `a.CurrencyCode = b.CcyCode` |

<a id="kolom-ccyisuse"></a>
## CcyIsUse

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSecurityCurrency_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000117 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CheckRekeningIsMcy:50](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CheckRekeningIsMcy.sql#L50) | WHERE / P11 / COMPARISON | `c.CcyIsUse` : `BIT` | `1` : `int` | `c.CcyIsUse` → `INT` | `c.CcyIsUse = 1` |
| IC-000123 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.CheckRekeningIsMcy:67](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CheckRekeningIsMcy.sql#L67) | WHERE / P11 / COMPARISON | `c.CcyIsUse` : `BIT` | `1` : `int` | `c.CcyIsUse` → `INT` | `c.CcyIsUse = 1` |
| IC-000522 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.GetAccountSyariah:103](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountSyariah.sql#L103) | JOIN_ON / P11 / COMPARISON | `b.CcyIsUse` : `BIT` | `1` : `int` | `b.CcyIsUse` → `INT` | `b.CcyIsUse = 1` |
| IC-000524 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.GetAccountSyariah:109](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountSyariah.sql#L109) | JOIN_ON / P11 / COMPARISON | `b.CcyIsUse` : `BIT` | `1` : `int` | `b.CcyIsUse` → `INT` | `b.CcyIsUse = 1` |
| IC-000563 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabah:149](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L149) | JOIN_ON / P11 / COMPARISON | `b.CcyIsUse` : `BIT` | `1` : `int` | `b.CcyIsUse` → `INT` | `b.CcyIsUse = 1` |
| IC-000569 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabah:174](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L174) | JOIN_ON / P11 / COMPARISON | `b.CcyIsUse` : `BIT` | `1` : `int` | `b.CcyIsUse` → `INT` | `b.CcyIsUse = 1` |
| IC-004138 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSGetSecurityCurrency:9](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetSecurityCurrency.sql#L9) | WHERE / P11 / COMPARISON | `CcyIsUse` : `BIT` | `1` : `int` | `CcyIsUse` → `INT` | `CcyIsUse = 1` |
| IC-005352 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDataParamDeviasiTotalSpread:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDataParamDeviasiTotalSpread.sql#L51) | WHERE / P11 / COMPARISON | `CcyIsUse` : `BIT` | `1` : `int` | `CcyIsUse` → `INT` | `CcyIsUse = 1` |
| IC-008545 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSResponseBONDCPYFromMX:283](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBONDCPYFromMX.sql#L283) | WHERE / P11 / COMPARISON | `sc.CcyIsUse` : `BIT` | `1` : `int` | `sc.CcyIsUse` → `INT` | `sc.CcyIsUse = 1` |
