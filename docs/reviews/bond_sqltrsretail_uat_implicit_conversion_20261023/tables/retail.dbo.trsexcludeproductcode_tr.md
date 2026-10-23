# Retail UAT.dbo.TRSExcludeProductCode_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 3 lokasi source, 3 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-productcode"></a>
## ProductCode

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSExcludeProductCode_TR.sql#L2). **Existing leading index:** TRSExcludeProductCode_TR_CN1.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000515 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetAccountSyariah:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountSyariah.sql#L52) | WHERE / P12 / IN_SUBQUERY | `(SELECT ProductCode FROM dbo.TRSExcludeProductCode_TR)` : `VARCHAR(20)` | `dd.SCCODE` : `CHAR(10)` | `dd.SCCODE` → `VARCHAR` | `dd.SCCODE IN (SELECT ProductCode FROM dbo.TRSExcludeProductCode_TR)` |
| IC-000528 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetAccountTaxAmnesty:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountTaxAmnesty.sql#L51) | WHERE / P12 / IN_SUBQUERY | `(SELECT ProductCode FROM dbo.TRSExcludeProductCode_TR)` : `VARCHAR(20)` | `dd.SCCODE` : `CHAR(10)` | `dd.SCCODE` → `VARCHAR` | `dd.SCCODE IN (SELECT ProductCode FROM dbo.TRSExcludeProductCode_TR)` |
| IC-000577 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabah:206](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L206) | WHERE / P12 / IN_SUBQUERY | `(SELECT ProductCode FROM dbo.TRSExcludeProductCode_TR)` : `VARCHAR(20)` | `dd.SCCODE` : `CHAR(10)` | `dd.SCCODE` → `VARCHAR` | `dd.SCCODE IN (SELECT ProductCode FROM dbo.TRSExcludeProductCode_TR)` |
