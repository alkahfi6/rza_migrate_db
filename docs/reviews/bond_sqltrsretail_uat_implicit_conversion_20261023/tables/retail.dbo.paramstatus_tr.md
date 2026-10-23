# Retail UAT.dbo.ParamStatus_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 7 lokasi source, 4 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-trxstatusid"></a>
## TrxStatusID

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ParamStatus_TR.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009805 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptDetailTrx:718](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L718) | JOIN_ON / P11 / COMPARISON | `p.TrxStatusID` : `INT` | `t.nTrxStatus` : `SMALLINT` | `t.nTrxStatus` → `INT` | `t.nTrxStatus = p.TrxStatusID` |
| IC-009902 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptDetailTrxTaxAmnesty:685](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L685) | JOIN_ON / P11 / COMPARISON | `p.TrxStatusID` : `INT` | `t.nTrxStatus` : `SMALLINT` | `t.nTrxStatus` → `INT` | `t.nTrxStatus = p.TrxStatusID` |

<a id="kolom-trxtypeid"></a>
## TrxTypeID

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ParamStatus_TR.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009804 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDetailTrx:713](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L713) | JOIN_ON / P11 / COMPARISON | `p.TrxTypeID` : `SMALLINT` | `1` : `int` | `p.TrxTypeID` → `INT` | `p.TrxTypeID = 1` |
| IC-009901 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDetailTrxTaxAmnesty:679](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L679) | JOIN_ON / P11 / COMPARISON | `p.TrxTypeID` : `SMALLINT` | `1` : `int` | `p.TrxTypeID` → `INT` | `p.TrxTypeID = 1` |
| IC-013289 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_ListParamStatus_TR:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListParamStatus_TR.sql#L17) | WHERE / P11 / COMPARISON | `TrxTypeID` : `SMALLINT` | `@TrxTypeID` : `int` | `TrxTypeID` → `INT` | `TrxTypeID= @TrxTypeID` |
| IC-013682 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_SearchTrxIdParamStatus:10](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_SearchTrxIdParamStatus.sql#L10) | WHERE / P11 / COMPARISON | `TrxTypeID` : `SMALLINT` | `@TrxType` : `int` | `TrxTypeID` → `INT` | `TrxTypeID=@TrxType` |
| IC-013683 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_SearchTrxIdParamStatus:12](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_SearchTrxIdParamStatus.sql#L12) | WHERE / P11 / COMPARISON | `TrxTypeID` : `SMALLINT` | `@TrxType` : `int` | `TrxTypeID` → `INT` | `TrxTypeID=@TrxType` |
