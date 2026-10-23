# Retail UAT.dbo.SecurityCouponDatePending_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 2 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-periode"></a>
## Periode

**Tipe dasar:** `VARCHAR(8)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityCouponDatePending_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P03 — Tanggal teks bertemu date/time:** Tentukan kontrak YYYYMMDD/hari/timestamp. Bentuk nilai pencarian dalam tipe kolom atau gunakan rentang native date; jangan konversi kolom seluruh baris tanpa alasan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007353 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1121](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1121) | CASE / P03 / COMPARISON | `cdp.Periode` : `VARCHAR(8)` | `@nLastCouponDate` : `DATETIME` | `cdp.Periode` → `DATETIME` | `cdp.Periode <= @nLastCouponDate` |
| IC-007365 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1621](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1621) | CASE / P03 / COMPARISON | `cdp.Periode` : `VARCHAR(8)` | `@nLastCouponDate` : `DATETIME` | `cdp.Periode` → `DATETIME` | `cdp.Periode <= @nLastCouponDate` |
