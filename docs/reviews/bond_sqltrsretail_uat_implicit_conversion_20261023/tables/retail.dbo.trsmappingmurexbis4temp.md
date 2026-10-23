# Retail UAT.dbo.TRSMappingMurexBIS4Temp — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 13 lokasi source, 2 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMappingMurexBIS4Temp.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003333 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApproveMappingMurexBIS4:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveMappingMurexBIS4.sql#L36) | WHERE / P12 / COMPARISON | `tmp.Status` : `CHAR(1)` | `'U'` : `varchar` | `tmp.Status` → `VARCHAR` | `tmp.Status = 'U'` |
| IC-003334 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApproveMappingMurexBIS4:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveMappingMurexBIS4.sql#L48) | CASE / P12 / COMPARISON | `tmp.Status` : `CHAR(1)` | `'U'` : `varchar` | `tmp.Status` → `VARCHAR` | `tmp.Status = 'U'` |
| IC-003336 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApproveMappingMurexBIS4:57](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveMappingMurexBIS4.sql#L57) | WHERE / P12 / IN_VALUE_TYPE | `tmp.Status` : `CHAR(1)` | `'U'` : `varchar` | `tmp.Status` → `VARCHAR` | `tmp.Status in ( 'U', 'D' )` |
| IC-003338 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApproveMappingMurexBIS4:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveMappingMurexBIS4.sql#L63) | WHERE / P12 / COMPARISON | `tmp.Status` : `CHAR(1)` | `'D'` : `varchar` | `tmp.Status` → `VARCHAR` | `tmp.Status = 'D'` |
| IC-003340 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSApproveMappingMurexBIS4:73](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveMappingMurexBIS4.sql#L73) | WHERE / P12 / COMPARISON | `tmp.Status` : `CHAR(1)` | `'C'` : `varchar` | `tmp.Status` → `VARCHAR` | `tmp.Status = 'C'` |
| IC-011663 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveMappingMurexBIS4:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveMappingMurexBIS4.sql#L111) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmappingmurexbis4temp.Status` : `CHAR(1)` | `@pcStatus` : `varchar(1)` | `@pcStatus` → `CHAR(1)` | `retail.dbo.trsmappingmurexbis4temp.Status <- @pcStatus` |
| IC-011664 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveMappingMurexBIS4:116](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveMappingMurexBIS4.sql#L116) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmappingmurexbis4temp.Status` : `CHAR(1)` | `@pcStatus` : `varchar(1)` | `@pcStatus` → `CHAR(1)` | `retail.dbo.trsmappingmurexbis4temp.Status <- @pcStatus` |

<a id="kolom-tempid"></a>
## TempId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMappingMurexBIS4Temp.sql#L2). **Existing leading index:** PK__TRSMappingMurexB__3F47339B.

**Frekuensi source:** 6 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003332 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSApproveMappingMurexBIS4:35](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveMappingMurexBIS4.sql#L35) | JOIN_ON / P11 / COMPARISON | `tmp.TempId` : `BIGINT` | `s.TempId` : `int` | `s.TempId` → `BIGINT` | `s.TempId = tmp.TempId` |
| IC-003335 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSApproveMappingMurexBIS4:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveMappingMurexBIS4.sql#L56) | JOIN_ON / P11 / COMPARISON | `tmp.TempId` : `BIGINT` | `s.TempId` : `int` | `s.TempId` → `BIGINT` | `tmp.TempId = s.TempId` |
| IC-003337 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSApproveMappingMurexBIS4:62](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveMappingMurexBIS4.sql#L62) | JOIN_ON / P11 / COMPARISON | `tmp.TempId` : `BIGINT` | `s.TempId` : `int` | `s.TempId` → `BIGINT` | `s.TempId = tmp.TempId` |
| IC-003339 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSApproveMappingMurexBIS4:72](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveMappingMurexBIS4.sql#L72) | JOIN_ON / P11 / COMPARISON | `tmp.TempId` : `BIGINT` | `s.TempId` : `int` | `s.TempId` → `BIGINT` | `s.TempId = tmp.TempId` |
| IC-003342 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSApproveMappingMurexBIS4:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveMappingMurexBIS4.sql#L89) | JOIN_ON / P11 / COMPARISON | `tmp.TempId` : `BIGINT` | `s.TempId` : `int` | `s.TempId` → `BIGINT` | `s.TempId = tmp.TempId` |
| IC-003343 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSApproveMappingMurexBIS4:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveMappingMurexBIS4.sql#L96) | JOIN_ON / P11 / COMPARISON | `tmp.TempId` : `BIGINT` | `s.TempId` : `int` | `s.TempId` → `BIGINT` | `s.TempId = tmp.TempId` |
