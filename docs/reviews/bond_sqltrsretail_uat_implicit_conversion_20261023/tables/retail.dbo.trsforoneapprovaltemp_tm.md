# Retail UAT.dbo.TRSForONEApprovalTemp_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 10 lokasi source, 2 procedure, 2 kolom. **A:** 8 lokasi pada 2 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-searchterm"></a>
## SearchTerm

**Tipe dasar:** `VARCHAR(100)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSForONEApprovalTemp_TM.sql#L13). **Existing leading index:** TRSForONEApprovalTemp_TM_CN1.

**Frekuensi source:** 8 lokasi / 2 SP. **A:** 8 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000892 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.ONEApprovalPrepareData:60](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareData.sql#L60) | JOIN_ON / P01 / COMPARISON | `oap.SearchTerm` : `VARCHAR(100)` | `obliNewMasterNasabah.CIFId` : `BIGINT` | `oap.SearchTerm` → `BIGINT` | `obliNewMasterNasabah.CIFId = oap.SearchTerm` |
| IC-000897 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.ONEApprovalPrepareData:90](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareData.sql#L90) | JOIN_ON / P01 / COMPARISON | `oap.SearchTerm` : `VARCHAR(100)` | `obliEditMasterNasabah.CIFId` : `BIGINT` | `oap.SearchTerm` → `BIGINT` | `obliEditMasterNasabah.CIFId = oap.SearchTerm` |
| IC-000904 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.ONEApprovalPrepareData:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareData.sql#L127) | JOIN_ON / P01 / COMPARISON | `oap.SearchTerm` : `VARCHAR(100)` | `st.DealId` : `BIGINT` | `oap.SearchTerm` → `BIGINT` | `st.DealId = oap.SearchTerm` |
| IC-000914 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.ONEApprovalPrepareData:174](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareData.sql#L174) | JOIN_ON / P01 / COMPARISON | `oap.SearchTerm` : `VARCHAR(100)` | `ot.OrderId` : `BIGINT` | `oap.SearchTerm` → `BIGINT` | `ot.OrderId = oap.SearchTerm` |
| IC-000920 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.ONEApprovalPrepareDataBond:58](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareDataBond.sql#L58) | JOIN_ON / P01 / COMPARISON | `oap.SearchTerm` : `VARCHAR(100)` | `obliNewMasterNasabah.CIFId` : `BIGINT` | `oap.SearchTerm` → `BIGINT` | `obliNewMasterNasabah.CIFId = oap.SearchTerm` |
| IC-000925 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.ONEApprovalPrepareDataBond:88](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareDataBond.sql#L88) | JOIN_ON / P01 / COMPARISON | `oap.SearchTerm` : `VARCHAR(100)` | `obliEditMasterNasabah.CIFId` : `BIGINT` | `oap.SearchTerm` → `BIGINT` | `obliEditMasterNasabah.CIFId = oap.SearchTerm` |
| IC-000932 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.ONEApprovalPrepareDataBond:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareDataBond.sql#L125) | JOIN_ON / P01 / COMPARISON | `oap.SearchTerm` : `VARCHAR(100)` | `st.DealId` : `BIGINT` | `oap.SearchTerm` → `BIGINT` | `st.DealId = oap.SearchTerm` |
| IC-000942 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.ONEApprovalPrepareDataBond:172](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareDataBond.sql#L172) | JOIN_ON / P01 / COMPARISON | `oap.SearchTerm` : `VARCHAR(100)` | `ot.OrderId` : `BIGINT` | `oap.SearchTerm` → `BIGINT` | `ot.OrderId = oap.SearchTerm` |

<a id="kolom-trxguid"></a>
## TrxGUID

**Tipe dasar:** `UNIQUEIDENTIFIER`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSForONEApprovalTemp_TM.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000915 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.ONEApprovalPrepareData:184](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareData.sql#L184) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsforoneapprovaltemp_tm.TrxGUID` : `UNIQUEIDENTIFIER` | `TrxGUID` : `varchar(100)` | `TrxGUID` → `UNIQUEIDENTIFIER` | `retail.dbo.trsforoneapprovaltemp_tm.TrxGUID <- TrxGUID` |
| IC-000943 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.ONEApprovalPrepareDataBond:182](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareDataBond.sql#L182) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsforoneapprovaltemp_tm.TrxGUID` : `UNIQUEIDENTIFIER` | `TrxGUID` : `varchar(100)` | `TrxGUID` → `UNIQUEIDENTIFIER` | `retail.dbo.trsforoneapprovaltemp_tm.TrxGUID <- TrxGUID` |
