# Listing procedure dengan risiko konversi kolom pada filter/join

Daftar ini mencakup kandidat dari binding statis, termasuk temporary table; belum membuktikan scan atau biaya runtime. Perbandingan CASE/projection, perbedaan numeric family, dan CHAR/VARCHAR biasa tidak dihitung di tabel ini. Semua 1.063 procedure tercakup dalam `procedure_summary.csv`. Lokasi dan arah konversi lengkap pada `implicit_conversions.csv`.

| Procedure | Teks → numerik | Teks → tanggal | ANSI → Unicode | Kandidat index |
|---|---:|---:|---:|---|
| [dbo.TreeViewOfficeForReportingServiceDBNISP](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TreeViewOfficeForReportingServiceDBNISP.sql) | 21 | 0 | 0 |  |
| [dbo.PTOSProsesDataJobOutstandingBond](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql) | 0 | 0 | 17 | IX01|IX10|IX11 |
| [dbo.PTOSProsesDataJobOutstanding](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql) | 0 | 0 | 15 | IX01|IX10|IX11 |
| [dbo.GetRekeningMasterNasabahTRS](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql) | 14 | 0 | 0 | IX14 |
| [dbo.PTOSProsesDataJobNasabah](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql) | 2 | 0 | 10 | IX10|IX17 |
| [dbo.TRSHasilRevenueBondJurnalGL](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSHasilRevenueBondJurnalGL.sql) | 7 | 4 | 0 |  |
| [dbo.TRSRPTMasterNasabahObligasi](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahObligasi.sql) | 11 | 0 | 0 | IX14 |
| [dbo.TRSReportSeller](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSeller.sql) | 10 | 0 | 0 |  |
| [dbo.TRSRptFOPandTransferAsset](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql) | 10 | 0 | 0 |  |
| [dbo.PTOSProsesDataJobTransaksi](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql) | 0 | 0 | 9 | IX10|IX12 |
| [dbo.PTOSProsesDataJobTransaksiBond](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql) | 0 | 0 | 9 | IX10|IX12 |
| [dbo.Rvmp_TRSPopulatePTOSUploadAuthen](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql) | 9 | 0 | 0 |  |
| [dbo.TRSPopulatePTOSUploadAuthen](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql) | 9 | 0 | 0 |  |
| [dbo.TRSSaveUploadCTPMurex](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveUploadCTPMurex.sql) | 9 | 0 | 0 |  |
| [dbo.GetRekeningMasterNasabah](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql) | 8 | 0 | 0 | IX14 |
| [dbo.TRSPopulateSettlementBIS4](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql) | 0 | 0 | 8 | IX01 |
| [dbo.TRSReportReportSettlementInstruction](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql) | 3 | 0 | 5 | IX04 |
| [dbo.trs_GetDataTransaksiAktifitas](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_GetDataTransaksiAktifitas.sql) | 7 | 0 | 0 | IX04|IX14 |
| [dbo.trs_InsertTrx2ndSecurityTransaction_TT](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTrx2ndSecurityTransaction_TT.sql) | 7 | 0 | 0 | IX01|IX04|IX06|IX15 |
| [dbo.TRSAmendSecurityTransaction_TT](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSAmendSecurityTransaction_TT.sql) | 6 | 0 | 0 | IX01|IX03|IX04 |
| [dbo.TRSORIOnlineNewTransaction](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOnlineNewTransaction.sql) | 6 | 0 | 0 | IX02 |
| [dbo.TRSRptTransactionDetailWMTA](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql) | 4 | 2 | 0 | IX14 |
| [dbo.TRSUpdateTaxAmnestyExpired](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateTaxAmnestyExpired.sql) | 6 | 0 | 0 | IX04|IX14 |
| [dbo.trs_InsertORIOrderTemp](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertORIOrderTemp.sql) | 6 | 0 | 0 |  |
| [dbo.trs_ListStatusORI_Order_TT](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql) | 6 | 0 | 0 | IX05 |
| [dbo.trs_UpdateTreasuryCustomer_TM](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql) | 6 | 0 | 0 | IX04|IX05 |
| [dbo.trs_updateStatusORI_Order_Temp](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateStatusORI_Order_Temp.sql) | 6 | 0 | 0 | IX05 |
| [dbo.ONEApprovalPrepareData](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareData.sql) | 5 | 0 | 0 | IX04 |
| [dbo.ONEApprovalPrepareDataBond](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ONEApprovalPrepareDataBond.sql) | 5 | 0 | 0 | IX04 |
| [dbo.SBNSchedPopulateORIOrder](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateORIOrder.sql) | 5 | 0 | 0 | IX02 |
| [dbo.TRSInsertRevenueBondJurnalGL](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql) | 2 | 0 | 3 | IX04 |
| [dbo.TRSORIGetProductDetail](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIGetProductDetail.sql) | 5 | 0 | 0 | IX02 |
| [dbo.TRSRPTRevenueOBLforTrops](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforTrops.sql) | 5 | 0 | 0 |  |
| [dbo.TRSSaveUploadCTP](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveUploadCTP.sql) | 5 | 0 | 0 | IX01 |
| [dbo.TRSTransferAssetDealIdSrc](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransferAssetDealIdSrc.sql) | 5 | 0 | 0 |  |
| [dbo.TRSUpdateDataMasterNasabah](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateDataMasterNasabah.sql) | 5 | 0 | 0 |  |
| [dbo.RepoSaveParamHargaModal](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoSaveParamHargaModal.sql) | 0 | 4 | 0 |  |
| [dbo.TRSDataNasabahRetailIndividu](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividu.sql) | 4 | 0 | 0 | IX14 |
| [dbo.TRSDataNasabahRetailInstitusiOrchid](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailInstitusiOrchid.sql) | 4 | 0 | 0 | IX14 |
| [dbo.TRSPTOSPopulateTransaksi](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPTOSPopulateTransaksi.sql) | 4 | 0 | 0 |  |
| [dbo.TRSPopulateEarlyRedemAuthorization](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEarlyRedemAuthorization.sql) | 4 | 0 | 0 |  |
| [dbo.TRSPopulateGagalBayarOri](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql) | 4 | 0 | 0 |  |
| [dbo.TRSPopulatePTOSOuts](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSOuts.sql) | 4 | 0 | 0 |  |
| [dbo.TRSPopulatePTOSOutsAuthen](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSOutsAuthen.sql) | 4 | 0 | 0 |  |
| [dbo.TRSPopulatePTOSUTrx](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql) | 4 | 0 | 0 | IX12 |
| [dbo.TRSProcessBuyBackOM](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBuyBackOM.sql) | 4 | 0 | 0 |  |
| [dbo.TRSRptBatchTransaction](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptBatchTransaction.sql) | 4 | 0 | 0 |  |
| [dbo.TRSRptDetailTrx](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql) | 4 | 0 | 0 | IX04|IX14 |
| [dbo.TRSRptDetailTrxTaxAmnesty](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql) | 4 | 0 | 0 | IX04|IX06|IX14 |
| [dbo.TRSRptKonfirmasiBOND](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBOND.sql) | 3 | 1 | 0 | IX04|IX13 |
| [dbo.TRSRptLaporanPajakBankBuy](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql) | 4 | 0 | 0 | IX04 |
| [dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo.sql) | 4 | 0 | 0 | IX04|IX06|IX14 |
| [dbo.TRSRptPTOSNasabah](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPTOSNasabah.sql) | 4 | 0 | 0 |  |
| [dbo.TRSSentEmailTransferAsset](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailTransferAsset.sql) | 4 | 0 | 0 |  |
| [dbo.trs_StatusSecurityTransaction_TT](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_StatusSecurityTransaction_TT.sql) | 4 | 0 | 0 | IX04 |
| [dbo.trs_updateTopUpORI_Order_Temp](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateTopUpORI_Order_Temp.sql) | 4 | 0 | 0 | IX05 |
| [dbo.CustodyReportMaintainAcquisition](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportMaintainAcquisition.sql) | 3 | 0 | 0 | IX04 |
| [dbo.ORISBNIPopulateEarlyRedemption](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateEarlyRedemption.sql) | 3 | 0 | 0 | IX02|IX05 |
| [dbo.PopulateBlotterRepo](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PopulateBlotterRepo.sql) | 3 | 0 | 0 |  |
| [dbo.RPAJurnalSettlementOriva](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RPAJurnalSettlementOriva.sql) | 3 | 0 | 0 |  |
| [dbo.RepoPopulateBlotter](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoPopulateBlotter.sql) | 3 | 0 | 0 |  |
| [dbo.SBNSchedProcessRekSB](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekSB.sql) | 3 | 0 | 0 |  |
| [dbo.TRSCheckStatusBIS4](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCheckStatusBIS4.sql) | 0 | 0 | 3 |  |
| [dbo.TRSRPTMasterNasabahPerubahanObligasi2](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql) | 3 | 0 | 0 |  |
| [dbo.TRSSentEmailHasilRevenueBondGagal](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailHasilRevenueBondGagal.sql) | 1 | 2 | 0 |  |
| [dbo.trs_InsertTreasuryCustomer_TM](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql) | 3 | 0 | 0 | IX14 |
| [dbo.trs_ListStatusTreasuryCustomer_TM](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusTreasuryCustomer_TM.sql) | 3 | 0 | 0 |  |
| [dbo.BKPopulateSummaryHolding](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BKPopulateSummaryHolding.sql) | 0 | 0 | 2 |  |
| [dbo.OMSubmitTransactionSekunder](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMSubmitTransactionSekunder.sql) | 2 | 0 | 0 | IX04|IX14 |
| [dbo.OMTrxSecPreBook](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMTrxSecPreBook.sql) | 2 | 0 | 0 | IX04|IX09 |
| [dbo.OMUpdateRegisterNasabah](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMUpdateRegisterNasabah.sql) | 2 | 0 | 0 |  |
| [dbo.ORIOnlineReportDailyTransaction](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORIOnlineReportDailyTransaction.sql) | 2 | 0 | 0 | IX01 |
| [dbo.RepoRONDataMigrate](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql) | 1 | 0 | 1 |  |
| [dbo.Rvmp_TRSPTOSProcessUpload](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql) | 2 | 0 | 0 | IX10|IX11|IX12|IX17 |
| [dbo.SBNSchedCheckAlert](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedCheckAlert.sql) | 2 | 0 | 0 |  |
| [dbo.SBNSchedPopulateNotif](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateNotif.sql) | 2 | 0 | 0 |  |
| [dbo.SBNSchedPushMailNotifOrderCompleted](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql) | 2 | 0 | 0 |  |
| [dbo.SettlementPopulateHasilRevenue](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SettlementPopulateHasilRevenue.sql) | 2 | 0 | 0 |  |
| [dbo.TRSOBLExpiredKitasRPT](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOBLExpiredKitasRPT.sql) | 2 | 0 | 0 |  |
| [dbo.TRSOBLPopulateParamRegionAreaBranchRPT](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOBLPopulateParamRegionAreaBranchRPT.sql) | 2 | 0 | 0 |  |
| [dbo.TRSPopulateApprovalFOPTransferAsset](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalFOPTransferAsset.sql) | 2 | 0 | 0 |  |
| [dbo.TRSPopulateBlotterTransaksi](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateBlotterTransaksi.sql) | 2 | 0 | 0 | IX04 |
| [dbo.TRSPopulateEditCustomerAuthorization](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEditCustomerAuthorization.sql) | 2 | 0 | 0 | IX14 |
| [dbo.TRSPopulateEditHargaModalTrx](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateEditHargaModalTrx.sql) | 2 | 0 | 0 | IX04 |
| [dbo.TRSPopulateReportPermission](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateReportPermission.sql) | 2 | 0 | 0 |  |
| [dbo.TRSPopulateRiskProfileNasabah](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateRiskProfileNasabah.sql) | 2 | 0 | 0 |  |
| [dbo.TRSPopulateSecurityMasterAuthorization](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSecurityMasterAuthorization.sql) | 2 | 0 | 0 |  |
| [dbo.TRSPopulateSettlementBIS4Detail](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4Detail.sql) | 0 | 0 | 2 |  |
| [dbo.TRSPopulateSettlementBIS4RepoDetail](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4RepoDetail.sql) | 0 | 0 | 2 |  |
| [dbo.TRSProcessInsertSecurityTransactionCommit](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertSecurityTransactionCommit.sql) | 2 | 0 | 0 | IX04 |
| [dbo.TRSRPTRevenueOBLforWM](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforWM.sql) | 2 | 0 | 0 | IX07 |
| [dbo.TRSReportFormBuySellObli](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportFormBuySellObli.sql) | 2 | 0 | 0 | IX04 |
| [dbo.TRSReportSettlementBankBeli](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSettlementBankBeli.sql) | 2 | 0 | 0 | IX01|IX04 |
| [dbo.TRSRetailUpdateCouponDate](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateCouponDate.sql) | 0 | 2 | 0 | IX13 |
| [dbo.TRSRetailUpdateCouponDate_Dummy](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailUpdateCouponDate_Dummy.sql) | 0 | 2 | 0 | IX13 |
| [dbo.TRSRptAuditTrailParamSpread](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParamSpread.sql) | 2 | 0 | 0 |  |
| [dbo.TRSRptAuditTrailParamWewenangDeviasi](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParamWewenangDeviasi.sql) | 2 | 0 | 0 |  |
| [dbo.TRSRptAuditTrailParameterHargaORI](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParameterHargaORI.sql) | 2 | 0 | 0 |  |
| [dbo.TRSRptCustomerLog](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCustomerLog.sql) | 2 | 0 | 0 |  |
| [dbo.TRSRptKonfirmasiREPO](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiREPO.sql) | 2 | 0 | 0 |  |
| [dbo.TRSRptKonfirmasiSRBI](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiSRBI.sql) | 0 | 2 | 0 | IX04|IX13 |
| [dbo.TRSRptObligasiKonfirmasiBond](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptObligasiKonfirmasiBond.sql) | 2 | 0 | 0 | IX03|IX04 |
| [dbo.TRSRptOrderORIDetail](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOrderORIDetail.sql) | 2 | 0 | 0 | IX01 |
| [dbo.TRSSyncMasterRiskProfileData](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncMasterRiskProfileData.sql) | 2 | 0 | 0 |  |
| [dbo.TRSUpdateRiskProfile](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateRiskProfile.sql) | 2 | 0 | 0 |  |
| [dbo.trs_ListStatusORI_Order_Temp](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_Temp.sql) | 2 | 0 | 0 |  |
| [dbo.BuybackLogInsert](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.BuybackLogInsert.sql) | 1 | 0 | 0 |  |
| [dbo.CustodyReportSettlementInstruction](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CustodyReportSettlementInstruction.sql) | 1 | 0 | 0 | IX01|IX04|IX05 |
| [dbo.GetAccountInfo](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountInfo.sql) | 1 | 0 | 0 |  |
| [dbo.ORISBNIPopulateRegisterInvestor](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNIPopulateRegisterInvestor.sql) | 1 | 0 | 0 | IX02 |
| [dbo.ORISBNISubmitInvestorPending](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNISubmitInvestorPending.sql) | 1 | 0 | 0 | IX02 |
| [dbo.RepoReportNoticePayment](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoReportNoticePayment.sql) | 1 | 0 | 0 |  |
| [dbo.RepoUploadValidation](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoUploadValidation.sql) | 1 | 0 | 0 | IX01|IX08 |
| [dbo.Rvmp_trs_ListSecurityMaster_TM](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql) | 1 | 0 | 0 | IX01|IX05|IX13 |
| [dbo.SBNSchedPopulateRegInvestor](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestor.sql) | 1 | 0 | 0 |  |
| [dbo.SBNSchedPopulateRegInvestorForOrder](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestorForOrder.sql) | 1 | 0 | 0 | IX02 |
| [dbo.SBNSchedPushMailNotifOrderORI](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderORI.sql) | 1 | 0 | 0 |  |
| [dbo.SRBIProcessEmailNotifMaturityDateH1](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH1.sql) | 1 | 0 | 0 |  |
| [dbo.SRBIProcessEmailNotifMaturityDateH3](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH3.sql) | 1 | 0 | 0 |  |
| [dbo.SRBIProcessEmailNotifMaturityDateH7](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH7.sql) | 1 | 0 | 0 |  |
| [dbo.TRSBlotterLelangPopulate](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBlotterLelangPopulate.sql) | 1 | 0 | 0 | IX01 |
| [dbo.TRSCheckDebitAccount](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCheckDebitAccount.sql) | 1 | 0 | 0 |  |
| [dbo.TRSDataNasabahRetailIndividuOrchid](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSDataNasabahRetailIndividuOrchid.sql) | 1 | 0 | 0 |  |
| [dbo.TRSFailedTransactionPrimaryPreBook](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSFailedTransactionPrimaryPreBook.sql) | 1 | 0 | 0 |  |
| [dbo.TRSGetAllProductName](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetAllProductName.sql) | 1 | 0 | 0 |  |
| [dbo.TRSGetEmailForPushBack](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetEmailForPushBack.sql) | 1 | 0 | 0 |  |
| [dbo.TRSListEarlyRedemption](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSListEarlyRedemption.sql) | 1 | 0 | 0 | IX01|IX04|IX14 |
| [dbo.TRSOBLPopulateParamRegionRPT](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOBLPopulateParamRegionRPT.sql) | 1 | 0 | 0 |  |
| [dbo.TRSORICheckProductEligibility](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORICheckProductEligibility.sql) | 1 | 0 | 0 | IX02 |
| [dbo.TRSORIOrderSummary](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSORIOrderSummary.sql) | 1 | 0 | 0 | IX01|IX05 |
| [dbo.TRSPopulateApprovalBuyBackOM](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalBuyBackOM.sql) | 1 | 0 | 0 |  |
| [dbo.TRSPopulateCboParam](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateCboParam.sql) | 1 | 0 | 0 |  |
| [dbo.TRSPopulateMainApprovalParam](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateMainApprovalParam.sql) | 1 | 0 | 0 |  |
| [dbo.TRSPopulatePTOSCustomerAuthen](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSCustomerAuthen.sql) | 1 | 0 | 0 |  |
| [dbo.TRSPopulatePTOSSecurityMasterAuthen](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSSecurityMasterAuthen.sql) | 1 | 0 | 0 |  |
| [dbo.TRSPopulatePTOSTrxAuthen](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSTrxAuthen.sql) | 1 | 0 | 0 |  |
| [dbo.TRSPopulatePTOSUpload](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUpload.sql) | 1 | 0 | 0 | IX10|IX11|IX12|IX17 |
| [dbo.TRSPopulateParameter](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParameter.sql) | 0 | 0 | 1 |  |
| [dbo.TRSProcessOverBookingMateraiBuyBack](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessOverBookingMateraiBuyBack.sql) | 1 | 0 | 0 | IX04 |
| [dbo.TRSProcessPTOSCustomer_TM](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSCustomer_TM.sql) | 1 | 0 | 0 | IX10 |
| [dbo.TRSProcessSecurityMasterAuthorization](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql) | 0 | 1 | 0 | IX01 |
| [dbo.TRSProcessSeller](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSeller.sql) | 1 | 0 | 0 |  |
| [dbo.TRSReportFilterCabangBaseRegional](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportFilterCabangBaseRegional.sql) | 1 | 0 | 0 |  |
| [dbo.TRSReportFilterPBOBaseRegional](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportFilterPBOBaseRegional.sql) | 1 | 0 | 0 |  |
| [dbo.TRSReportRekonMaterai](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql) | 1 | 0 | 0 | IX04 |
| [dbo.TRSReportSKFBankBeli](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeli.sql) | 1 | 0 | 0 | IX04 |
| [dbo.TRSReportSKFBankBeliBKON](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeliBKON.sql) | 1 | 0 | 0 | IX04 |
| [dbo.TRSReportTrxCorporateBOND](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportTrxCorporateBOND.sql) | 1 | 0 | 0 | IX04 |
| [dbo.TRSRptKonfirmasiBONDNew](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptKonfirmasiBONDNew.sql) | 0 | 1 | 0 | IX04|IX13 |
| [dbo.TRSRptMonitoringDocAcc](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptMonitoringDocAcc.sql) | 1 | 0 | 0 |  |
| [dbo.TRSRptMonitoringDocTrx](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptMonitoringDocTrx.sql) | 1 | 0 | 0 |  |
| [dbo.TRSRptOutstandingPerSecAccNo](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingPerSecAccNo.sql) | 1 | 0 | 0 | IX04|IX14 |
| [dbo.TRSRptPTOSOutstanding](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPTOSOutstanding.sql) | 1 | 0 | 0 | IX10|IX11 |
| [dbo.TRSRptPTOSTransaksi](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPTOSTransaksi.sql) | 1 | 0 | 0 | IX10 |
| [dbo.TRSRptPajakKupon](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPajakKupon.sql) | 1 | 0 | 0 | IX04|IX14 |
| [dbo.TRSRptPayCouponORITaxAmnesty](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPayCouponORITaxAmnesty.sql) | 0 | 1 | 0 | IX01|IX04|IX13|IX14 |
| [dbo.TRSRptProductWM](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptProductWM.sql) | 1 | 0 | 0 | IX04 |
| [dbo.TRSSyncSellerDescFromPS](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSyncSellerDescFromPS.sql) | 1 | 0 | 0 |  |
| [dbo.TRSUpdateCTPMurex](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateCTPMurex.sql) | 1 | 0 | 0 |  |
| [dbo.trs_CancelTopUpOriOrder_Temp](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_CancelTopUpOriOrder_Temp.sql) | 1 | 0 | 0 |  |
| [dbo.trs_ListSecurityMasterPending_TM](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql) | 1 | 0 | 0 | IX05 |
| [dbo.trs_ListTreasuryCustomer_TM_Original](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListTreasuryCustomer_TM_Original.sql) | 1 | 0 | 0 | IX14 |
| [dbo.trs_PopulateDataRekeningMaterai](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningMaterai.sql) | 1 | 0 | 0 | IX01 |
| [dbo.trs_PopulateDataRekeningObligasi](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningObligasi.sql) | 1 | 0 | 0 | IX01 |
| [dbo.trs_StatusRetailCustBlock_TM](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_StatusRetailCustBlock_TM.sql) | 1 | 0 | 0 | IX16 |
| [dbo.trs_UpdateSecurityMaster_TM](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateSecurityMaster_TM.sql) | 1 | 0 | 0 | IX01 |
| [dbo.trs_updateStatusAccORI_Order_TT](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_updateStatusAccORI_Order_TT.sql) | 1 | 0 | 0 | IX05 |
