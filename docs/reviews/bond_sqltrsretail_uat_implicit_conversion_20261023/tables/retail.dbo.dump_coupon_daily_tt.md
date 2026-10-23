# Retail UAT.dbo.dump_coupon_daily_tt — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 3 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-inserted-date"></a>
## inserted_date

**Tipe dasar:** `DATETIME2(7)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.dump_coupon_daily_tt.sql#L24). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002129 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSGenerateDumpOBL084:167](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGenerateDumpOBL084.sql#L167) | WHERE / P13 / COMPARISON | `inserted_date` : `DATETIME2(7)` | `@dCurrentDate` : `datetime` | `@dCurrentDate` → `DATETIME2` | `inserted_date >= @dCurrentDate` |
| IC-002181 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSKuponDumpOBL084:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponDumpOBL084.sql#L70) | WHERE / P13 / COMPARISON | `inserted_date` : `DATETIME2(7)` | `DATEADD(MINUTE,-5,GETDATE())` : `datetime` | `DATEADD(MINUTE,-5,GETDATE())` → `DATETIME2` | `inserted_date >= DATEADD(MINUTE,-5,GETDATE())` |
| IC-002182 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSKuponDumpOBL084:79](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponDumpOBL084.sql#L79) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.dump_coupon_daily_tt.inserted_date` : `DATETIME2(7)` | `InsertedDate` : `datetime` | `InsertedDate` → `DATETIME2(7)` | `retail.dbo.dump_coupon_daily_tt.inserted_date <- InsertedDate` |

<a id="kolom-ppn"></a>
## ppn

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.dump_coupon_daily_tt.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006964 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessKuponUpload:281](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessKuponUpload.sql#L281) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `d.ppn` : `MONEY` | `#tFromDatabase.PPN` : `DECIMAL(19,2)` | `d.ppn` → `DECIMAL(19,2)` | `#tFromDatabase.PPN <- d.ppn` |
