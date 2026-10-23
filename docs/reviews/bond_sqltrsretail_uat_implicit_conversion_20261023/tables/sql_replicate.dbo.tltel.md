# SQL_REPLICATE.dbo.TLTEL — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 2 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-tlnum"></a>
## TLNUM

**Tipe dasar:** `DECIMAL(7,0)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.TLTEL.sql#L2). **Existing leading index:** idx_tlnum.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009768 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptDetailTrx:101](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L101) | WHERE / P11 / COMPARISON | `TLNUM` : `DECIMAL(7,0)` | `@nNIK` : `INT` | `@nNIK` → `DECIMAL` | `TLNUM = @nNIK` |
| IC-011454 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptPopulateParamBranch:58](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPopulateParamBranch.sql#L58) | WHERE / P11 / COMPARISON | `TLNUM` : `DECIMAL(7,0)` | `@nNIK` : `INT` | `@nNIK` → `DECIMAL` | `TLNUM = @nNIK` |

<a id="kolom-tlttyp"></a>
## TLTTYP

**Tipe dasar:** `VARCHAR(3)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.TLTEL.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009767 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptDetailTrx:99](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L99) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TLTTYP` : `VARCHAR(3)` | `0` : `int` | `0` → `VARCHAR(3)` | `ISNULL(TLTTYP,0)` |
| IC-011453 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPopulateParamBranch:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPopulateParamBranch.sql#L56) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TLTTYP` : `VARCHAR(3)` | `0` : `int` | `0` → `VARCHAR(3)` | `ISNULL(TLTTYP,0)` |
