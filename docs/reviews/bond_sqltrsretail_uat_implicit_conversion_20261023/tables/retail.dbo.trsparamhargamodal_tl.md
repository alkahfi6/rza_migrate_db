# Retail UAT.dbo.TRSParamHargaModal_TL — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 18 lokasi source, 6 procedure, 5 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-actionflag"></a>
## ActionFlag

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModal_TL.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005238 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprParamHargaModal:6](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamHargaModal.sql#L6) | WHERE / P12 / COMPARISON | `ActionFlag` : `CHAR(1)` | `ParamId` : `VARCHAR(20)` | `ActionFlag` → `VARCHAR` | `ParamId = ActionFlag` |
| IC-005239 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprParamHargaModal:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamHargaModal.sql#L16) | WHERE / P12 / COMPARISON | `ActionFlag` : `CHAR(1)` | `'O'` : `varchar` | `ActionFlag` → `VARCHAR` | `ActionFlag = 'O'` |
| IC-005240 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprParamHargaModal:19](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamHargaModal.sql#L19) | WHERE / P12 / COMPARISON | `ActionFlag` : `CHAR(1)` | `ParamId` : `VARCHAR(20)` | `ActionFlag` → `VARCHAR` | `ParamId = ActionFlag` |
| IC-005241 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprParamHargaModal:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprParamHargaModal.sql#L29) | WHERE / P12 / COMPARISON | `ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `ActionFlag` → `VARCHAR` | `ActionFlag = 'N'` |
| IC-006570 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprParamHargaModal:20](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamHargaModal.sql#L20) | WHERE / P12 / COMPARISON | `pl.ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `pl.ActionFlag` → `VARCHAR` | `pl.ActionFlag = 'N'` |

<a id="kolom-actiontype"></a>
## ActionType

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModal_TL.sql#L19). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006664 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessApprovalHargaModal:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalHargaModal.sql#L46) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamhargamodal_tl.ActionType` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.trsparamhargamodal_tl.ActionType <- 'O'` |
| IC-006692 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessApprovalHargaModal:238](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalHargaModal.sql#L238) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamhargamodal_tl.ActionType` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.trsparamhargamodal_tl.ActionType <- 'N'` |

<a id="kolom-hargamodalbeli"></a>
## HargaModalBeli

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModal_TL.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006571 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamHargaModal:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamHargaModal.sql#L48) | WHERE / P11 / COMPARISON | `HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `HargaModalBeli != 0` |
| IC-006573 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamHargaModal:67](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamHargaModal.sql#L67) | WHERE / P11 / COMPARISON | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalBeli != 0` |
| IC-006601 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamSpread:155](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L155) | WHERE / P11 / COMPARISON | `HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `HargaModalBeli != 0` |
| IC-006603 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamSpread:172](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L172) | WHERE / P11 / COMPARISON | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalBeli != 0` |

<a id="kolom-hargamodaljual"></a>
## HargaModalJual

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModal_TL.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006572 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamHargaModal:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamHargaModal.sql#L56) | WHERE / P11 / COMPARISON | `HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `HargaModalJual != 0` |
| IC-006574 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamHargaModal:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamHargaModal.sql#L77) | WHERE / P11 / COMPARISON | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalJual!= 0` |
| IC-006602 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamSpread:161](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L161) | WHERE / P11 / COMPARISON | `HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `HargaModalJual != 0` |
| IC-006604 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamSpread:182](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L182) | WHERE / P11 / COMPARISON | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalJual!= 0` |

<a id="kolom-processtype"></a>
## ProcessType

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModal_TL.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002596 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSSaveParamHargaModalNew_Normal:309](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSaveParamHargaModalNew_Normal.sql#L309) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamhargamodal_tl.ProcessType` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.trsparamhargamodal_tl.ProcessType <- 'N'` |
| IC-002681 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSSaveParamHargaModalNew_Panic:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSaveParamHargaModalNew_Panic.sql#L225) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamhargamodal_tl.ProcessType` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trsparamhargamodal_tl.ProcessType <- 'P'` |
| IC-006665 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessApprovalHargaModal:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalHargaModal.sql#L46) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamhargamodal_tl.ProcessType` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.trsparamhargamodal_tl.ProcessType <- 'A'` |
