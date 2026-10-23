# SQL_REPLICATE.dbo.office_information — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 2 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-office-id"></a>
## office_id

**Tipe dasar:** `CHAR(3)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.office_information.sql#L2). **Existing leading index:** office_information_pk.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005331 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateBranch:38](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateBranch.sql#L38) | WHERE / P12 / COMPARISON | `office_id` : `CHAR(3)` | `@cOfficeId` : `varchar(10)` | `office_id` → `VARCHAR` | `office_id = @cOfficeId` |
| IC-005974 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateRegional:19](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateRegional.sql#L19) | WHERE / P12 / COMPARISON | `office_id` : `CHAR(3)` | `@cOfficeId` : `varchar(10)` | `office_id` → `VARCHAR` | `office_id = @cOfficeId` |

<a id="kolom-regional-office-id"></a>
## regional_office_id

**Tipe dasar:** `CHAR(3)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.office_information.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005973 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateRegional:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateRegional.sql#L17) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `regional_office_id` : `CHAR(3)` | `@cRegional` : `varchar(3)` | `regional_office_id` → `varchar(3)` | `@cRegional = regional_office_id` |
| IC-005330 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateBranch:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateBranch.sql#L30) | WHERE / P12 / COMPARISON | `regional_office_id` : `CHAR(3)` | `@cRegionalOfficeId` : `varchar(3)` | `regional_office_id` → `VARCHAR` | `regional_office_id = @cRegionalOfficeId` |
