# Retail UAT.dbo.TRSParamDeviasiTotalSpread_TR_Temp — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 12 lokasi source, 3 procedure, 4 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 2 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-apprid"></a>
## ApprId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamDeviasiTotalSpread_TR_Temp.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006559 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamDeviasiTotalSpread:35](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamDeviasiTotalSpread.sql#L35) | WHERE / P11 / COMPARISON | `ApprId` : `BIGINT` | `@pnApprId` : `int` | `@pnApprId` → `BIGINT` | `ApprId = @pnApprId` |
| IC-006560 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamDeviasiTotalSpread:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamDeviasiTotalSpread.sql#L51) | WHERE / P11 / COMPARISON | `tmp.ApprId` : `BIGINT` | `@pnApprId` : `int` | `@pnApprId` → `BIGINT` | `tmp.ApprId = @pnApprId` |
| IC-006561 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamDeviasiTotalSpread:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamDeviasiTotalSpread.sql#L85) | WHERE / P11 / COMPARISON | `ApprId` : `BIGINT` | `@pnApprId` : `int` | `@pnApprId` → `BIGINT` | `ApprId = @pnApprId` |
| IC-006562 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamDeviasiTotalSpread:88](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamDeviasiTotalSpread.sql#L88) | WHERE / P11 / COMPARISON | `ApprId` : `BIGINT` | `@pnApprId` : `int` | `@pnApprId` → `BIGINT` | `ApprId = @pnApprId` |
| IC-006563 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamDeviasiTotalSpread:90](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamDeviasiTotalSpread.sql#L90) | WHERE / P11 / COMPARISON | `ApprId` : `BIGINT` | `@pnApprId` : `int` | `@pnApprId` → `BIGINT` | `ApprId = @pnApprId` |

<a id="kolom-id"></a>
## Id

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamDeviasiTotalSpread_TR_Temp.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003191 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSAddParamDeviasiTotalSpread:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAddParamDeviasiTotalSpread.sql#L17) | VARIABLE_ASSIGNMENT / P08 / SETVARIABLESTATEMENT | `(select max(Id) from dbo.TRSParamDeviasiTotalSpread_TR_Temp)` : `INT` | `@nMaxId` : `bigint` | `(select max(Id) from dbo.TRSParamDeviasiTotalSpread_TR_Temp)` → `bigint` | `set @nMaxId = (select max(Id) from dbo.TRSParamDeviasiTotalSpread_TR_Temp)` |
| IC-006558 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessApprParamDeviasiTotalSpread:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamDeviasiTotalSpread.sql#L33) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `Id` : `INT` | `@nId` : `bigint` | `Id` → `bigint` | `@nId = Id` |
| IC-003192 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSAddParamDeviasiTotalSpread:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAddParamDeviasiTotalSpread.sql#L21) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamdeviasitotalspread_tr_temp.Id` : `INT` | `@nMaxId + 1` : `BIGINT (combined result; facets unknown)` | `@nMaxId + 1` → `INT` | `retail.dbo.trsparamdeviasitotalspread_tr_temp.Id <- @nMaxId + 1` |

<a id="kolom-tenorjatuhtempofrom"></a>
## TenorJatuhTempoFrom

**Tipe dasar:** `DECIMAL(10,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamDeviasiTotalSpread_TR_Temp.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003193 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSAddParamDeviasiTotalSpread:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAddParamDeviasiTotalSpread.sql#L21) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamdeviasitotalspread_tr_temp.TenorJatuhTempoFrom` : `DECIMAL(10,2)` | `@pcTenorJatuhTempoFrom` : `varchar(20)` | `@pcTenorJatuhTempoFrom` → `DECIMAL(10,2)` | `retail.dbo.trsparamdeviasitotalspread_tr_temp.TenorJatuhTempoFrom <- @pcTenorJatuhTempoFrom` |
| IC-005236 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateApprParamDeviasiTotalHarga:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamDeviasiTotalHarga.sql#L59) | WHERE / P04 / COMPARISON | `tmp.TenorJatuhTempoFrom` : `DECIMAL(10,2)` | `cast(convert(varchar,tr.TenorJatuhTempoFrom, 105) as varchar)` : `varchar` | `cast(convert(varchar,tr.TenorJatuhTempoFrom, 105) as varchar)` → `DECIMAL` | `tmp.TenorJatuhTempoFrom != cast(convert(varchar,tr.TenorJatuhTempoFrom, 105) as varchar)` |

<a id="kolom-tenorjatuhtempoto"></a>
## TenorJatuhTempoTo

**Tipe dasar:** `DECIMAL(10,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamDeviasiTotalSpread_TR_Temp.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003194 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSAddParamDeviasiTotalSpread:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAddParamDeviasiTotalSpread.sql#L21) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamdeviasitotalspread_tr_temp.TenorJatuhTempoTo` : `DECIMAL(10,2)` | `@pcTenorJatuhTempoTo` : `varchar(20)` | `@pcTenorJatuhTempoTo` → `DECIMAL(10,2)` | `retail.dbo.trsparamdeviasitotalspread_tr_temp.TenorJatuhTempoTo <- @pcTenorJatuhTempoTo` |
| IC-005237 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateApprParamDeviasiTotalHarga:60](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamDeviasiTotalHarga.sql#L60) | WHERE / P04 / COMPARISON | `tmp.TenorJatuhTempoTo` : `DECIMAL(10,2)` | `cast(convert(varchar,tr.TenorJatuhTempoTo, 105) as varchar)` : `varchar` | `cast(convert(varchar,tr.TenorJatuhTempoTo, 105) as varchar)` → `DECIMAL` | `tmp.TenorJatuhTempoTo != cast(convert(varchar,tr.TenorJatuhTempoTo, 105) as varchar)` |
