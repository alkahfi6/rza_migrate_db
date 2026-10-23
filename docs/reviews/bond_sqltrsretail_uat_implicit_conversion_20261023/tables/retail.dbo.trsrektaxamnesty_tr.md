# Retail UAT.dbo.TRSRekTaxAmnesty_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 5 lokasi source, 2 procedure, 1 kolom. **A:** 5 lokasi pada 2 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-norekinvestor"></a>
## NoRekInvestor

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSRekTaxAmnesty_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 2 SP. **A:** 5 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000606 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabahTRS:109](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L109) | JOIN_ON / P01 / COMPARISON | `a.NoRekInvestor` : `VARCHAR(20)` | `b.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = b.ACCTNO` |
| IC-000607 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabahTRS:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L111) | JOIN_ON / P01 / COMPARISON | `a.NoRekInvestor` : `VARCHAR(20)` | `c.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = c.ACCTNO` |
| IC-000626 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabahTRS:232](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L232) | JOIN_ON / P01 / COMPARISON | `a.NoRekInvestor` : `VARCHAR(20)` | `b.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = b.ACCTNO` |
| IC-000627 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabahTRS:234](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L234) | JOIN_ON / P01 / COMPARISON | `a.NoRekInvestor` : `VARCHAR(20)` | `c.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = c.ACCTNO` |
| IC-012672 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSUpdateTaxAmnestyExpired:295](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateTaxAmnestyExpired.sql#L295) | JOIN_ON / P01 / COMPARISON | `a.NoRekInvestor` : `VARCHAR(20)` | `b.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = b.ACCTNO` |
