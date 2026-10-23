# SQL_REPLICATE.dbo.PBFlagDesc_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 3 lokasi source, 3 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-flagtype"></a>
## FlagType

**Tipe dasar:** `CHAR(2)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.PBFlagDesc_TR.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007853 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSREPOReportOutstandingNonPVB:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportOutstandingNonPVB.sql#L21) | WHERE / P12 / COMPARISON | `cf.FlagType` : `CHAR(2)` | `'PV'` : `varchar` | `cf.FlagType` → `VARCHAR` | `cf.FlagType = 'PV'` |
| IC-009587 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptCustomerLog:103](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerLog.sql#L103) | JOIN_ON / P12 / COMPARISON | `pb.FlagType` : `CHAR(2)` | `'PV'` : `varchar` | `pb.FlagType` → `VARCHAR` | `pb.FlagType = 'PV'` |
| IC-012954 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_CekNasabahPremier:11](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_CekNasabahPremier.sql#L11) | WHERE / P12 / COMPARISON | `fd.FlagType` : `CHAR(2)` | `'PB'` : `varchar` | `fd.FlagType` → `VARCHAR` | `fd.FlagType = 'PB'` |
