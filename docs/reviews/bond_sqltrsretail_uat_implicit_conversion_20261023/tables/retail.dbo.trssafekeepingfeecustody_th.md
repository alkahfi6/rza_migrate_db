# Retail UAT.dbo.TRSSafeKeepingFeeCustody_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 8 lokasi source, 2 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-outstandingproduct"></a>
## OutstandingProduct

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSafeKeepingFeeCustody_TH.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-011493 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSafeKeepingFeeCustody:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSafeKeepingFeeCustody.sql#L31) | CASE / P05 / BINARY_Divide | `skfcth.OutstandingProduct` : `MONEY` | `convert (decimal(30,8), skfcth.Outstanding)` : `decimal(30,8)` | `skfcth.OutstandingProduct` → `DECIMAL` | `convert (decimal(30,8), skfcth.Outstanding) / skfcth.OutstandingProduct` |
| IC-011494 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptSafeKeepingFeeCustody:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSafeKeepingFeeCustody.sql#L31) | CASE / P05 / BINARY_Multiply | `(convert (decimal(30,8), skfcth.Outstanding) / skfcth.OutstandingProduct)` : `DECIMAL (combined result; facets unknown)` | `skfcth.SKFProduct` : `MONEY` | `skfcth.SKFProduct` → `DECIMAL` | `(convert (decimal(30,8), skfcth.Outstanding) / skfcth.OutstandingProduct) * skfcth.SKFProduct` |

<a id="kolom-skfproduct"></a>
## SKFProduct

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSafeKeepingFeeCustody_TH.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-011494 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSafeKeepingFeeCustody:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSafeKeepingFeeCustody.sql#L31) | CASE / P05 / BINARY_Multiply | `skfcth.SKFProduct` : `MONEY` | `(convert (decimal(30,8), skfcth.Outstanding) / skfcth.OutstandingProduct)` : `DECIMAL (combined result; facets unknown)` | `skfcth.SKFProduct` → `DECIMAL` | `(convert (decimal(30,8), skfcth.Outstanding) / skfcth.OutstandingProduct) * skfcth.SKFProduct` |
| IC-011495 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSafeKeepingFeeCustodySummary:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSafeKeepingFeeCustodySummary.sql#L26) | EXPRESSION / P07 / CASE_BRANCH | `skfcth.SKFProduct` : `MONEY` | `result type` : `DECIMAL (combined result; facets unknown)` | `skfcth.SKFProduct` → `DECIMAL` | `case when skfcth.SafeKeepingFeeType = 1 then skfcth.SKFProduct when skfcth.SafeKeepingFeeType = 2 then skfcth.SKFProduct when (skfcth.SafeKeepingFeeType = 3) and skfcth.SKFProduct < smtm.MinSafeKeepingFee then smtm.MinSafeKeepingFee when (skfcth.SafeKeepingFeeType = 3) and skfcth.SKFProduct > smtm.MinSafeKeepingFee then skfcth.SKFProduct end` |
| IC-011496 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSafeKeepingFeeCustodySummary:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSafeKeepingFeeCustodySummary.sql#L26) | EXPRESSION / P07 / CASE_BRANCH | `skfcth.SKFProduct` : `MONEY` | `result type` : `DECIMAL (combined result; facets unknown)` | `skfcth.SKFProduct` → `DECIMAL` | `case when skfcth.SafeKeepingFeeType = 1 then skfcth.SKFProduct when skfcth.SafeKeepingFeeType = 2 then skfcth.SKFProduct when (skfcth.SafeKeepingFeeType = 3) and skfcth.SKFProduct < smtm.MinSafeKeepingFee then smtm.MinSafeKeepingFee when (skfcth.SafeKeepingFeeType = 3) and skfcth.SKFProduct > smtm.MinSafeKeepingFee then skfcth.SKFProduct end` |
| IC-011497 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSafeKeepingFeeCustodySummary:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSafeKeepingFeeCustodySummary.sql#L26) | EXPRESSION / P07 / CASE_BRANCH | `skfcth.SKFProduct` : `MONEY` | `result type` : `DECIMAL (combined result; facets unknown)` | `skfcth.SKFProduct` → `DECIMAL` | `case when skfcth.SafeKeepingFeeType = 1 then skfcth.SKFProduct when skfcth.SafeKeepingFeeType = 2 then skfcth.SKFProduct when (skfcth.SafeKeepingFeeType = 3) and skfcth.SKFProduct < smtm.MinSafeKeepingFee then smtm.MinSafeKeepingFee when (skfcth.SafeKeepingFeeType = 3) and skfcth.SKFProduct > smtm.MinSafeKeepingFee then skfcth.SKFProduct end` |
| IC-011498 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSafeKeepingFeeCustodySummary:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSafeKeepingFeeCustodySummary.sql#L29) | CASE / P11 / COMPARISON | `skfcth.SKFProduct` : `MONEY` | `smtm.MinSafeKeepingFee` : `DECIMAL(20,5)` | `skfcth.SKFProduct` → `DECIMAL` | `skfcth.SKFProduct < smtm.MinSafeKeepingFee` |
| IC-011499 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSafeKeepingFeeCustodySummary:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSafeKeepingFeeCustodySummary.sql#L30) | CASE / P11 / COMPARISON | `skfcth.SKFProduct` : `MONEY` | `smtm.MinSafeKeepingFee` : `DECIMAL(20,5)` | `skfcth.SKFProduct` → `DECIMAL` | `skfcth.SKFProduct > smtm.MinSafeKeepingFee` |

<a id="kolom-safekeepingfeeamount"></a>
## SafeKeepingFeeAmount

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSafeKeepingFeeCustody_TH.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-011492 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSafeKeepingFeeCustody:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSafeKeepingFeeCustody.sql#L29) | EXPRESSION / P07 / CASE_BRANCH | `skfcth.SafeKeepingFeeAmount` : `MONEY` | `result type` : `DECIMAL (combined result; facets unknown)` | `skfcth.SafeKeepingFeeAmount` → `DECIMAL` | `case when skfcth.SafeKeepingFeeType = 1 then skfcth.SafeKeepingFeeAmount when skfcth.SafeKeepingFeeType in (2, 3) then (convert (decimal(30,8), skfcth.Outstanding) / skfcth.OutstandingProduct) * skfcth.SKFProduct end` |
