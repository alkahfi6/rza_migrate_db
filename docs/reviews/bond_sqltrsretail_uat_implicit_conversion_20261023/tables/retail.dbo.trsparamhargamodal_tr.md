# Retail UAT.dbo.TRSParamHargaModal_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 134 lokasi source, 29 procedure, 10 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 5 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-hargamodalbeli"></a>
## HargaModalBeli

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModal_TR.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 52 lokasi / 19 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008771 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailCalculateFee:468](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L468) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `HargaModalBeli` : `DECIMAL(20,5)` | `@nDealPriceBankBuy` : `money` | `HargaModalBeli` → `money` | `@nDealPriceBankBuy = HargaModalBeli` |
| IC-008777 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailCalculateFee:547](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L547) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `HargaModalBeli` : `DECIMAL(20,5)` | `@dHargaModalBankBeli` : `money` | `HargaModalBeli` → `money` | `@dHargaModalBankBeli = HargaModalBeli` |
| IC-000680 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.OMGetListProductSekunder:416](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L416) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(sm.HargaModalBeli, 0)` |
| IC-002611 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSSaveParamHargaModalNew_Normal:457](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSaveParamHargaModalNew_Normal.sql#L457) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalBeli, 0)` |
| IC-002613 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSSaveParamHargaModalNew_Normal:459](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSaveParamHargaModalNew_Normal.sql#L459) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalBeli, 0)` |
| IC-002625 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSSaveParamHargaModalNew_Normal:558](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSaveParamHargaModalNew_Normal.sql#L558) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalBeli, 0)` |
| IC-002627 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSSaveParamHargaModalNew_Normal:560](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSaveParamHargaModalNew_Normal.sql#L560) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalBeli, 0)` |
| IC-003429 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSBondsPricePopulate:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPricePopulate.sql#L30) | CASE / P11 / COMPARISON | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalBeli = 0` |
| IC-003432 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSBondsPricePopulate:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPricePopulate.sql#L36) | CASE / P11 / COMPARISON | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalBeli = 0` |
| IC-003444 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSBondsPricePopulate:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPricePopulate.sql#L59) | CASE / P11 / COMPARISON | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalBeli = 0` |
| IC-003453 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSBondsPricePopulate:97](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPricePopulate.sql#L97) | CASE / P11 / COMPARISON | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalBeli = 0` |
| IC-003610 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCmprSecurityNo:101](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCmprSecurityNo.sql#L101) | CASE / P11 / COMPARISON | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalBeli = 0` |
| IC-003627 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCmprSecurityNoChannel:92](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCmprSecurityNoChannel.sql#L92) | CASE / P11 / COMPARISON | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalBeli = 0` |
| IC-003986 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSECBondsPricePopulate:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSECBondsPricePopulate.sql#L34) | CASE / P11 / COMPARISON | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalBeli = 0` |
| IC-004456 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSMaintenanceParamHargaModal:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamhargamodal_tr.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `retail.dbo.trsparamhargamodal_tr.HargaModalBeli <- 0` |
| IC-004472 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMaintenanceParamHargaModal:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L337) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ph.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(ph.HargaModalBeli, 0)` |
| IC-004474 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMaintenanceParamHargaModal:340](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L340) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ph.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(ph.HargaModalBeli, 0)` |
| IC-004865 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSORIOnlinePopulateHist:339](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlinePopulateHist.sql#L339) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(sm.HargaModalBeli, 0)` |
| IC-005943 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateParamSpread:68](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSpread.sql#L68) | CASE / P11 / COMPARISON | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalBeli = 0` |
| IC-005952 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateParamSpread:84](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSpread.sql#L84) | CASE / P11 / COMPARISON | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalBeli = 0` |
| IC-006578 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamHargaModal:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamHargaModal.sql#L104) | CASE / P11 / COMPARISON | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalBeli = 0` |
| IC-006581 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamHargaModal:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamHargaModal.sql#L110) | CASE / P11 / COMPARISON | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalBeli = 0` |
| IC-006608 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamSpread:208](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L208) | CASE / P11 / COMPARISON | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalBeli = 0` |
| IC-006611 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessApprParamSpread:209](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L209) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalBeli,0)` |
| IC-006613 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamSpread:214](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L214) | CASE / P11 / COMPARISON | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalBeli = 0` |
| IC-006616 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessApprParamSpread:215](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L215) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalBeli,0)` |
| IC-006630 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamSpread:275](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L275) | CASE / P11 / COMPARISON | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalBeli = 0` |
| IC-006633 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessApprParamSpread:276](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L276) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalBeli,0)` |
| IC-006635 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamSpread:281](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L281) | CASE / P11 / COMPARISON | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalBeli = 0` |
| IC-006638 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessApprParamSpread:282](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L282) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalBeli,0)` |
| IC-006671 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessApprovalHargaModal:112](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalHargaModal.sql#L112) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalBeli, 0)` |
| IC-006673 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessApprovalHargaModal:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalHargaModal.sql#L115) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalBeli, 0)` |
| IC-006681 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessApprovalHargaModal:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalHargaModal.sql#L165) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalBeli, 0)` |
| IC-006683 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessApprovalHargaModal:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalHargaModal.sql#L168) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalBeli, 0)` |
| IC-007333 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:661](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L661) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamhargamodal_tr.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `retail.dbo.trsparamhargamodal_tr.HargaModalBeli <- 0` |
| IC-011688 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModal:83](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModal.sql#L83) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(HargaModalBeli,0)` |
| IC-011738 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModal:458](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModal.sql#L458) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalBeli, 0)` |
| IC-011740 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModal:461](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModal.sql#L461) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalBeli, 0)` |
| IC-011748 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModal:510](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModal.sql#L510) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalBeli, 0)` |
| IC-011750 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModal:513](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModal.sql#L513) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalBeli, 0)` |
| IC-011771 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModalNew:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L115) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(HargaModalBeli,0)` |
| IC-011780 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModalNew:211](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L211) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `v.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(v.HargaModalBeli,0)` |
| IC-011861 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModalNew:584](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L584) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalBeli, 0)` |
| IC-011863 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModalNew:587](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L587) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalBeli, 0)` |
| IC-011871 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModalNew:638](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L638) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalBeli, 0)` |
| IC-011873 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModalNew:641](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L641) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalBeli, 0)` |
| IC-012749 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSValidateUploadParamHargaModal:268](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModal.sql#L268) | WHERE / P11 / COMPARISON | `ISNULL(HargaModalBeli,0)` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `ISNULL(HargaModalBeli,0) > 0` |
| IC-012750 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateUploadParamHargaModal:268](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModal.sql#L268) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(HargaModalBeli,0)` |
| IC-012781 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateUploadParamHargaModalNew:339](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L339) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(tr.HargaModalBeli , 0)` |
| IC-012817 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSValidateUploadParamHargaModalNew:609](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L609) | WHERE / P11 / COMPARISON | `ISNULL(HargaModalBeli,0)` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `ISNULL(HargaModalBeli,0) > 0` |
| IC-012818 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateUploadParamHargaModalNew:609](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L609) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(HargaModalBeli,0)` |
| IC-013920 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidateHargaModal:14](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateHargaModal.sql#L14) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `HargaModalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(HargaModalBeli,0)` |

<a id="kolom-hargamodaljual"></a>
## HargaModalJual

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModal_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 52 lokasi / 18 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007750 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessUploadTransaksiSuratBerharga:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessUploadTransaksiSuratBerharga.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `isnull(mdl.HargaModalJual,0)` : `DECIMAL(20,5)` | `#tempSecurityTransaction_TT.HargaModalAwal` : `MONEY` | `isnull(mdl.HargaModalJual,0)` → `MONEY` | `#tempSecurityTransaction_TT.HargaModalAwal <- isnull(mdl.HargaModalJual,0)` |
| IC-008772 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailCalculateFee:469](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L469) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `HargaModalJual` : `DECIMAL(20,5)` | `@nDealPriceBankSell` : `money` | `HargaModalJual` → `money` | `@nDealPriceBankSell = HargaModalJual` |
| IC-008773 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailCalculateFee:496](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L496) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `@nDealPrice - HargaModalJual` : `DECIMAL (combined result; facets unknown)` | `@nSpreadBankSell` : `money` | `@nDealPrice - HargaModalJual` → `money` | `@nSpreadBankSell = @nDealPrice - HargaModalJual` |
| IC-002616 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSSaveParamHargaModalNew_Normal:464](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSaveParamHargaModalNew_Normal.sql#L464) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalJual, 0)` |
| IC-002618 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSSaveParamHargaModalNew_Normal:466](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSaveParamHargaModalNew_Normal.sql#L466) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalJual, 0)` |
| IC-002633 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSSaveParamHargaModalNew_Normal:580](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSaveParamHargaModalNew_Normal.sql#L580) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalJual, 0)` |
| IC-002635 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSSaveParamHargaModalNew_Normal:582](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSaveParamHargaModalNew_Normal.sql#L582) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalJual, 0)` |
| IC-003436 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSBondsPricePopulate:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPricePopulate.sql#L45) | CASE / P11 / COMPARISON | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalJual != 0` |
| IC-003440 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSBondsPricePopulate:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPricePopulate.sql#L51) | CASE / P11 / COMPARISON | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalJual != 0` |
| IC-003447 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSBondsPricePopulate:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPricePopulate.sql#L66) | CASE / P11 / COMPARISON | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalJual != 0` |
| IC-003456 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSBondsPricePopulate:103](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPricePopulate.sql#L103) | CASE / P11 / COMPARISON | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalJual != 0` |
| IC-003613 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCmprSecurityNo:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCmprSecurityNo.sql#L107) | CASE / P11 / COMPARISON | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalJual != 0` |
| IC-003630 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCmprSecurityNoChannel:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCmprSecurityNoChannel.sql#L98) | CASE / P11 / COMPARISON | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalJual != 0` |
| IC-003989 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSECBondsPricePopulate:40](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSECBondsPricePopulate.sql#L40) | CASE / P11 / COMPARISON | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalJual != 0` |
| IC-004457 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSMaintenanceParamHargaModal:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamhargamodal_tr.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `retail.dbo.trsparamhargamodal_tr.HargaModalJual <- 0` |
| IC-004478 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMaintenanceParamHargaModal:359](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L359) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ph.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(ph.HargaModalJual, 0)` |
| IC-004480 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMaintenanceParamHargaModal:362](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L362) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ph.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(ph.HargaModalJual, 0)` |
| IC-005947 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateParamSpread:74](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSpread.sql#L74) | CASE / P11 / COMPARISON | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalJual != 0` |
| IC-005956 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateParamSpread:90](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSpread.sql#L90) | CASE / P11 / COMPARISON | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalJual != 0` |
| IC-006585 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamHargaModal:122](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamHargaModal.sql#L122) | CASE / P11 / COMPARISON | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalJual = 0` |
| IC-006588 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamHargaModal:128](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamHargaModal.sql#L128) | CASE / P11 / COMPARISON | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalJual = 0` |
| IC-006619 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamSpread:226](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L226) | CASE / P11 / COMPARISON | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalJual = 0` |
| IC-006622 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessApprParamSpread:227](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L227) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalJual,0)` |
| IC-006624 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamSpread:232](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L232) | CASE / P11 / COMPARISON | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalJual = 0` |
| IC-006627 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessApprParamSpread:233](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L233) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalJual,0)` |
| IC-006642 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamSpread:306](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L306) | CASE / P11 / COMPARISON | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalJual = 0` |
| IC-006645 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessApprParamSpread:307](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L307) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalJual,0)` |
| IC-006647 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessApprParamSpread:312](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L312) | CASE / P11 / COMPARISON | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `hm.HargaModalJual = 0` |
| IC-006650 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessApprParamSpread:313](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprParamSpread.sql#L313) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalJual,0)` |
| IC-006676 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessApprovalHargaModal:120](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalHargaModal.sql#L120) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalJual, 0)` |
| IC-006678 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessApprovalHargaModal:123](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalHargaModal.sql#L123) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalJual, 0)` |
| IC-006687 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessApprovalHargaModal:187](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalHargaModal.sql#L187) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalJual, 0)` |
| IC-006689 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessApprovalHargaModal:190](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovalHargaModal.sql#L190) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalJual, 0)` |
| IC-007334 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:661](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L661) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamhargamodal_tr.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `retail.dbo.trsparamhargamodal_tr.HargaModalJual <- 0` |
| IC-007765 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessUploadTransaksiSuratBerharga:150](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessUploadTransaksiSuratBerharga.sql#L150) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mdl.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(mdl.HargaModalJual,0)` |
| IC-011687 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModal:82](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModal.sql#L82) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(HargaModalJual,0)` |
| IC-011743 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModal:467](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModal.sql#L467) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalJual, 0)` |
| IC-011745 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModal:470](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModal.sql#L470) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalJual, 0)` |
| IC-011754 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModal:533](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModal.sql#L533) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalJual, 0)` |
| IC-011756 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModal:536](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModal.sql#L536) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalJual, 0)` |
| IC-011770 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModalNew:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L114) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(HargaModalJual,0)` |
| IC-011779 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModalNew:210](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L210) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `v.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(v.HargaModalJual,0)` |
| IC-011866 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModalNew:593](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L593) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalJual, 0)` |
| IC-011868 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModalNew:596](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L596) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalJual, 0)` |
| IC-011877 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModalNew:661](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L661) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalJual, 0)` |
| IC-011879 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModalNew:664](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L664) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `hm.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(hm.HargaModalJual, 0)` |
| IC-012747 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSValidateUploadParamHargaModal:267](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModal.sql#L267) | WHERE / P11 / COMPARISON | `ISNULL(HargaModalJual,0)` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `ISNULL(HargaModalJual,0) > 0` |
| IC-012748 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateUploadParamHargaModal:267](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModal.sql#L267) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(HargaModalJual,0)` |
| IC-012780 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateUploadParamHargaModalNew:334](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L334) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(tr.HargaModalJual, 0)` |
| IC-012815 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSValidateUploadParamHargaModalNew:608](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L608) | WHERE / P11 / COMPARISON | `ISNULL(HargaModalJual,0)` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `ISNULL(HargaModalJual,0) > 0` |
| IC-012816 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateUploadParamHargaModalNew:608](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L608) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(HargaModalJual,0)` |
| IC-013921 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ValidateHargaModal:14](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ValidateHargaModal.sql#L14) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `HargaModalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(HargaModalJual,0)` |

<a id="kolom-kuotatrxsekunder"></a>
## KuotaTrxSekunder

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModal_TR.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000814 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.OMValidateBalance:243](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateBalance.sql#L243) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `KuotaTrxSekunder` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(KuotaTrxSekunder,0)` |
| IC-000817 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.OMValidateBalance:245](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateBalance.sql#L245) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `KuotaTrxSekunder` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(KuotaTrxSekunder,0)` |
| IC-004458 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSMaintenanceParamHargaModal:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamhargamodal_tr.KuotaTrxSekunder` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `retail.dbo.trsparamhargamodal_tr.KuotaTrxSekunder <- 0` |
| IC-004462 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSMaintenanceParamHargaModal:171](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L171) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `KuotaTrxSekunder` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `KuotaTrxSekunder = 0` |
| IC-006916 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessInsertSecurityTransactionCommit:311](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertSecurityTransactionCommit.sql#L311) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `KuotaTrxSekunder` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(KuotaTrxSekunder,0)` |
| IC-012740 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateUploadParamHargaModal:139](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModal.sql#L139) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `td2.KuotaTrxSekunder` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(td2.KuotaTrxSekunder,0)` |
| IC-012770 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateUploadParamHargaModalNew:287](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L287) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `td2.KuotaTrxSekunder` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(td2.KuotaTrxSekunder,0)` |

<a id="kolom-lastupdateby"></a>
## LastUpdateBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModal_TR.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-011778 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModalNew:209](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L209) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `v.LastUpdateBy` : `INT` | `''` : `varchar` | `''` → `INT` | `ISNULL(v.LastUpdateBy, '')` |

<a id="kolom-limit"></a>
## Limit

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModal_TR.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004459 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSMaintenanceParamHargaModal:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamhargamodal_tr.Limit` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `retail.dbo.trsparamhargamodal_tr.Limit <- 0` |
| IC-004463 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSMaintenanceParamHargaModal:172](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L172) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Limit` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `Limit = 0` |

<a id="kolom-secid"></a>
## SecId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModal_TR.sql#L2). **Existing leading index:** TRSParamHargaModal_TR_CU1.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004452 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSMaintenanceParamHargaModal:121](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L121) | JOIN_ON / P11 / COMPARISON | `phm.SecId` : `INT` | `excl.SecId` : `BIGINT` | `phm.SecId` → `BIGINT` | `phm.SecId = excl.SecId` |

<a id="kolom-transactiondate"></a>
## TransactionDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModal_TR.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004460 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSMaintenanceParamHargaModal:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamhargamodal_tr.TransactionDate` : `DATETIME` | `convert(varchar, getdate(), 112)` : `varchar` | `convert(varchar, getdate(), 112)` → `DATETIME` | `retail.dbo.trsparamhargamodal_tr.TransactionDate <- convert(varchar, getdate(), 112)` |
| IC-004464 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSMaintenanceParamHargaModal:179](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L179) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `TransactionDate` : `DATETIME` | `convert(varchar, getdate(), 112)` : `varchar` | `convert(varchar, getdate(), 112)` → `DATETIME` | `TransactionDate = convert(varchar, getdate(), 112)` |
| IC-007335 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:661](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L661) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamhargamodal_tr.TransactionDate` : `DATETIME` | `convert(varchar, getdate(), 112)` : `varchar` | `convert(varchar, getdate(), 112)` → `DATETIME` | `retail.dbo.trsparamhargamodal_tr.TransactionDate <- convert(varchar, getdate(), 112)` |

<a id="kolom-yieldbeli"></a>
## YieldBeli

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModal_TR.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002424 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulateParamHargaModalNew:146](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateParamHargaModalNew.sql#L146) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr1.YieldBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(tr1.YieldBeli, 0)` |
| IC-005902 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateParamHargaModalNew:203](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamHargaModalNew.sql#L203) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr1.YieldBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(tr1.YieldBeli, 0)` |
| IC-011781 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModalNew:214](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L214) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `v.YieldBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(v.YieldBeli, 0)` |
| IC-012783 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateUploadParamHargaModalNew:351](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L351) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.YieldBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(tr.YieldBeli , 0)` |
| IC-012787 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateUploadParamHargaModalNew:384](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L384) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.YieldBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(tr.YieldBeli, 0)` |

<a id="kolom-yieldjual"></a>
## YieldJual

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModal_TR.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002425 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulateParamHargaModalNew:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulateParamHargaModalNew.sql#L147) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr1.YieldJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(tr1.YieldJual, 0)` |
| IC-005903 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateParamHargaModalNew:204](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamHargaModalNew.sql#L204) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr1.YieldJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(tr1.YieldJual, 0)` |
| IC-011782 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamHargaModalNew:215](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L215) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `v.YieldJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(v.YieldJual, 0)` |
| IC-012782 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateUploadParamHargaModalNew:345](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L345) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.YieldJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(tr.YieldJual, 0)` |
| IC-012790 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSValidateUploadParamHargaModalNew:392](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadParamHargaModalNew.sql#L392) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.YieldJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(tr.YieldJual, 0)` |

<a id="kolom-ispanicon"></a>
## isPanicON

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamHargaModal_TR.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000720 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMSubmitTransactionSekunder:153](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMSubmitTransactionSekunder.sql#L153) | WHERE / P11 / COMPARISON | `tr.isPanicON` : `BIT` | `1` : `int` | `tr.isPanicON` → `INT` | `tr.isPanicON = 1` |
| IC-000756 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMTrxSecPreBook:470](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMTrxSecPreBook.sql#L470) | WHERE / P11 / COMPARISON | `tl.isPanicON` : `BIT` | `1` : `int` | `tl.isPanicON` → `INT` | `tl.isPanicON = 1` |
| IC-002670 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSSaveParamHargaModalNew_Panic:161](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSSaveParamHargaModalNew_Panic.sql#L161) | CASE / P11 / COMPARISON | `tr.isPanicON` : `BIT` | `1` : `int` | `tr.isPanicON` → `INT` | `tr.isPanicON = 1` |
| IC-004325 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSInsertSecurityTransaction_TT_Original:916](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L916) | WHERE / P11 / COMPARISON | `isPanicON` : `BIT` | `1` : `int` | `isPanicON` → `INT` | `isPanicON = 1` |
| IC-005355 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDataTransaksiSuratBerharga:74](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDataTransaksiSuratBerharga.sql#L74) | WHERE / P11 / COMPARISON | `tl.isPanicON` : `BIT` | `1` : `int` | `tl.isPanicON` → `INT` | `tl.isPanicON = 1` |
| IC-011786 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveParamHargaModalNew:238](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamHargaModalNew.sql#L238) | CASE / P11 / COMPARISON | `tr.isPanicON` : `BIT` | `1` : `int` | `tr.isPanicON` → `INT` | `tr.isPanicON = 1` |
