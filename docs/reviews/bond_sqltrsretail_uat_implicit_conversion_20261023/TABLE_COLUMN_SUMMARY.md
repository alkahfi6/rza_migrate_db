# Ringkasan seluruh tabel dan kolom: konversi implisit

Sumber 23 Oktober 2026. Unit frekuensi adalah **site pasangan operand/ordinal** dan **procedure berbeda**, bukan execution count. Satu statement dapat memiliki beberapa site konversi. Tabel sumber: 676 Retail UAT + 165 SQL_REPLICATE. Semua tabel tercatat pada `table_summary.csv`, termasuk yang tidak mempunyai match.

A = konversi pada kolom fisik dalam filter/join/operasi predicate yang berpotensi mengganggu akses index. B = expression/assignment/function boundary. C = perbedaan family rutin. Partner/destination dicatat untuk lineage dan **tidak berarti kolom tersebut ikut dikonversi**. Tidak ada runtime performance certification.

## Kolom yang paling perlu ditinjau untuk akses index

Tabel ini memuat **seluruh** kolom A, diurutkan jumlah procedure. Klik nama tabel untuk semua procedure dan query pada kolom tersebut.

| Tabel | Kolom (tipe source) | Lokasi A | SP A | Pola | Existing leading index |
|---|---|---:|---:|---|---|
| [SQL_REPLICATE.dbo.PS_NDM_EMPLOYEES](tables/sql_replicate.dbo.ps_ndm_employees.md#kolom-emplid) | `EMPLID` (`CHAR(11)`) | 59 | 20 | P01, P12 | EMPLID_EFFDT_ON_PS_NDM_EMPLOYEES&#124;idx_PS_NDM_EMPLOYEES_1&#124;ix_PS_NDM_EMPLOYEES |
| [SQL_REPLICATE.dbo.PPLMAPMasterWorkflow_TM](tables/sql_replicate.dbo.pplmapmasterworkflow_tm.md#kolom-emplid) | `EMPLID` (`VARCHAR(11)`) | 40 | 18 | P01, P08, P12 | idx_PPLMAPMasterWorkflow_TM_1&#124;idx_PPLMAPMasterWorkflow_TM_2 |
| [SQL_REPLICATE.dbo.MISOfficeInformationEODFull](tables/sql_replicate.dbo.misofficeinformationeodfull.md#kolom-regional-office-id) | `regional_office_id` (`VARCHAR(5)`) | 41 | 11 | P01 | MISOfficeInformationEODFull_NC1 |
| [Retail UAT.dbo.ORI_Online_TR](tables/retail.dbo.ori_online_tr.md#kolom-kemenkeustatusid) | `KemenkeuStatusId` (`VARCHAR(20)`) | 18 | 10 | P01, P07, P08, P10 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.SecurityMaster_TM](tables/retail.dbo.securitymaster_tm.md#kolom-idseri) | `IdSeri` (`VARCHAR(100)`) | 18 | 8 | P01, P08 | Tidak ada key pertama di snapshot |
| [SQL_REPLICATE.dbo.CFPAR3](tables/sql_replicate.dbo.cfpar3.md#kolom-cp3uic) | `CP3UIC` (`VARCHAR(1)`) | 10 | 7 | P01 | idx_cp3uic_cp3ucd |
| [Retail UAT.dbo.SecurityCouponDate_TR](tables/retail.dbo.securitycoupondate_tr.md#kolom-periode) | `Periode` (`VARCHAR(8)`) | 9 | 7 | P03, P08 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TRSParameter_TR](tables/retail.dbo.trsparameter_tr.md#kolom-value) | `Value` (`VARCHAR(255)`) | 7 | 7 | P01, P05, P07, P08, P12 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.PTOSCustomer_TM](tables/retail.dbo.ptoscustomer_tm.md#kolom-sid) | `SID` (`VARCHAR(100)`) | 21 | 5 | P02, P08 | ix_ptoscustomer_tm_sidnordn |
| [Retail UAT.dbo.REPOMasterTransaction_TT](tables/retail.dbo.repomastertransaction_tt.md#kolom-nomurexexternal) | `NoMurexExternal` (`VARCHAR(20)`) | 6 | 5 | P01, P08, P09 | Tidak ada key pertama di snapshot |
| [SQL_REPLICATE.dbo.ProCIFRiskProfile_TM](tables/sql_replicate.dbo.procifriskprofile_tm.md#kolom-cifnumber) | `CIFNumber` (`VARCHAR(19)`) | 6 | 5 | P01, P08 | ProCIFRiskProfile_TM_idx1 |
| [Retail UAT.dbo.SecuritySysParameter_TR](tables/retail.dbo.securitysysparameter_tr.md#kolom-paramvalue) | `ParamValue` (`VARCHAR(30)`) | 14 | 4 | P01 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TRSPostingGLBond_TH](tables/retail.dbo.trspostingglbond_th.md#kolom-tranid) | `TranId` (`VARCHAR(20)`) | 14 | 4 | P01 | Tidak ada key pertama di snapshot |
| [SQL_REPLICATE.dbo.DDMAST](tables/sql_replicate.dbo.ddmast.md#kolom-branch) | `BRANCH` (`INT`) | 6 | 4 | P04, P06, P08, P09 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.ORISBNInvestorRekSB_TR](tables/retail.dbo.orisbninvestorreksb_tr.md#kolom-idreksb) | `IdRekSB` (`VARCHAR(20)`) | 6 | 4 | P01, P08, P09 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.SecurityTransaction_TT](tables/retail.dbo.securitytransaction_tt.md#kolom-trxbranch) | `TrxBranch` (`VARCHAR(5)`) | 5 | 4 | P01, P10 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.BONDSDetailTransaksiBuyBackPerdana_TR](tables/retail.dbo.bondsdetailtransaksibuybackperdana_tr.md#kolom-dealidsekunder) | `DealIdSekunder` (`VARCHAR(100)`) | 6 | 3 | P01, P08, P09 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TRSCustAccNumber_TR](tables/retail.dbo.trscustaccnumber_tr.md#kolom-norekinvestor) | `NoRekInvestor` (`VARCHAR(20)`) | 6 | 3 | P01, P08 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.REPOMasterTransaction_TT](tables/retail.dbo.repomastertransaction_tt.md#kolom-nikinputer) | `NIKInputer` (`VARCHAR(10)`) | 4 | 3 | P01, P10 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.REPOMasterTransaction_TT](tables/retail.dbo.repomastertransaction_tt.md#kolom-nikseller) | `NIKSeller` (`VARCHAR(10)`) | 4 | 3 | P01, P10 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.gwp_process_data_tt](tables/retail.dbo.gwp_process_data_tt.md#kolom-process-key) | `process_key` (`VARCHAR(100)`) | 3 | 3 | P01 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.SBNMappingSubReg_TR](tables/retail.dbo.sbnmappingsubreg_tr.md#kolom-idsubreg) | `IdSubReg` (`VARCHAR(15)`) | 3 | 3 | P01, P08 | SBNMappingSubReg_TR_CN1 |
| [Retail UAT.dbo.TRSBankCustody_TM](tables/retail.dbo.trsbankcustody_tm.md#kolom-idsubreg) | `IdSubReg` (`VARCHAR(10)`) | 3 | 3 | P01 | Tidak ada key pertama di snapshot |
| [SQL_REPLICATE.dbo.CFADDR](tables/sql_replicate.dbo.cfaddr.md#kolom-cfaseq) | `CFASEQ` (`CHAR(3)`) | 3 | 3 | P01 | idx_cfaddr_cfaseq |
| [Retail UAT.dbo.PTOSOutstanding_TR](tables/retail.dbo.ptosoutstanding_tr.md#kolom-kodeproduk) | `KodeProduk` (`VARCHAR(20)`) | 10 | 2 | P02, P08 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.PTOSOutstanding_TR](tables/retail.dbo.ptosoutstanding_tr.md#kolom-sid) | `SID` (`VARCHAR(100)`) | 10 | 2 | P02, P08 | ix_ptosoutstanding_tr_sid&#124;ix_ptosoutstanding_tr_sidnordn |
| [Retail UAT.dbo.PTOSTransaksi_TR](tables/retail.dbo.ptostransaksi_tr.md#kolom-dealid) | `DealId` (`VARCHAR(20)`) | 8 | 2 | P02, P08 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TRSForONEApprovalTemp_TM](tables/retail.dbo.trsforoneapprovaltemp_tm.md#kolom-searchterm) | `SearchTerm` (`VARCHAR(100)`) | 8 | 2 | P01 | TRSForONEApprovalTemp_TM_CN1 |
| [Retail UAT.dbo.TRSHasilRevenue_TM](tables/retail.dbo.trshasilrevenue_tm.md#kolom-tanggaltransaksi) | `TanggalTransaksi` (`VARCHAR(8)`) | 6 | 2 | P03, P08 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.SecurityMaster_TM](tables/retail.dbo.securitymaster_tm.md#kolom-isincode) | `ISINCode` (`VARCHAR(20)`) | 6 | 2 | P02 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TRSRekTaxAmnesty_TR](tables/retail.dbo.trsrektaxamnesty_tr.md#kolom-norekinvestor) | `NoRekInvestor` (`VARCHAR(20)`) | 5 | 2 | P01 | Tidak ada key pertama di snapshot |
| [SQL_REPLICATE.dbo.CFMAST](tables/sql_replicate.dbo.cfmast.md#kolom-cfcif) | `CFCIF` (`CHAR(19)`) | 4 | 2 | P01, P08, P12 | idx_cfmast_cfcif |
| [Retail UAT.dbo.TRSEmailOneNotifParameter_TR](tables/retail.dbo.trsemailonenotifparameter_tr.md#kolom-value) | `Value` (`VARCHAR(MAX)`) | 4 | 2 | P01, P08 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TreasuryCustomer_TM](tables/retail.dbo.treasurycustomer_tm.md#kolom-cifno) | `CIFNo` (`VARCHAR(19)`) | 3 | 2 | P01, P08, P12 | TreasuryCustomer_TM_NN1&#124;TreasuryCustomer_TM_NN4 |
| [Retail UAT.dbo.TRSGeneralParam_TR](tables/retail.dbo.trsgeneralparam_tr.md#kolom-paramid) | `ParamId` (`VARCHAR(20)`) | 2 | 2 | P01, P07, P12 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.MappingOccupation_TM](tables/retail.dbo.mappingoccupation_tm.md#kolom-kodeoccupation) | `KodeOccupation` (`CHAR(1)`) | 2 | 2 | P01, P07, P08 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TRSProcessOverBooking_TR](tables/retail.dbo.trsprocessoverbooking_tr.md#kolom-status) | `Status` (`VARCHAR(3)`) | 2 | 2 | P01, P08 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TRSSettlementBIS4_temp](tables/retail.dbo.trssettlementbis4_temp.md#kolom-participantbuyer) | `ParticipantBuyer` (`VARCHAR(500)`) | 2 | 2 | P02, P08 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TRSSettlementBIS4_temp](tables/retail.dbo.trssettlementbis4_temp.md#kolom-participantseller) | `ParticipantSeller` (`VARCHAR(500)`) | 2 | 2 | P02, P08 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.PTOSOutstanding_Error](tables/retail.dbo.ptosoutstanding_error.md#kolom-kodeproduk) | `KodeProduk` (`VARCHAR(20)`) | 2 | 2 | P02, P08 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.PTOSOutstanding_Error](tables/retail.dbo.ptosoutstanding_error.md#kolom-sid) | `SID` (`VARCHAR(100)`) | 2 | 2 | P02, P08 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.PTOSTransaksi_TR](tables/retail.dbo.ptostransaksi_tr.md#kolom-kodeproduk) | `KodeProduk` (`VARCHAR(20)`) | 2 | 2 | P02, P08 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.PTOSTransaksi_TR](tables/retail.dbo.ptostransaksi_tr.md#kolom-sid) | `SID` (`VARCHAR(100)`) | 2 | 2 | P02, P08 | ix_ptostransaksi_tr_sid |
| [Retail UAT.dbo.TreasuryCustomer_TM](tables/retail.dbo.treasurycustomer_tm.md#kolom-kodecabangalamatsurat) | `KodeCabangAlamatSurat` (`VARCHAR(5)`) | 2 | 2 | P01 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TRSParamaterReport_TR](tables/retail.dbo.trsparamaterreport_tr.md#kolom-value) | `Value` (`VARCHAR(100)`) | 2 | 2 | P01 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TRSSettlementBIS4_temp](tables/retail.dbo.trssettlementbis4_temp.md#kolom-buyercashaccount) | `BuyerCashAccount` (`VARCHAR(500)`) | 2 | 2 | P02 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TRSSettlementBIS4_temp](tables/retail.dbo.trssettlementbis4_temp.md#kolom-sellercashaccount) | `SellerCashAccount` (`VARCHAR(500)`) | 2 | 2 | P02 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TRSMappingMurexBIS4](tables/retail.dbo.trsmappingmurexbis4.md#kolom-bis4code) | `BIS4Code` (`VARCHAR(50)`) | 8 | 1 | P02 | TRSMappingMurexBIS4_NN2 |
| [Retail UAT.dbo.PTOSTransaksi_TR](tables/retail.dbo.ptostransaksi_tr.md#kolom-trxtype) | `TrxType` (`VARCHAR(8)`) | 4 | 1 | P01, P08 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TRSRekSyariah_TR](tables/retail.dbo.trsreksyariah_tr.md#kolom-norekinvestor) | `NoRekInvestor` (`VARCHAR(20)`) | 4 | 1 | P01, P08 | Tidak ada key pertama di snapshot |
| [SQL_REPLICATE.dbo.MISOfficeInformationEODFull](tables/sql_replicate.dbo.misofficeinformationeodfull.md#kolom-office-id) | `office_id` (`INT`) | 3 | 1 | P01, P06, P11 | PK_MISOfficeInformationEODFull |
| [Retail UAT.dbo.SecurityTransferAset_TM](tables/retail.dbo.securitytransferaset_tm.md#kolom-dealid) | `DealId` (`INT`) | 2 | 1 | P04, P06, P08 | SecurityTransferAset_TM_CN1 |
| [Retail UAT.dbo.TreasuryCustomer_TM](tables/retail.dbo.treasurycustomer_tm.md#kolom-norekinvestor) | `NoRekInvestor` (`VARCHAR(20)`) | 2 | 1 | P01, P08 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.BKMappingSec_TM](tables/retail.dbo.bkmappingsec_tm.md#kolom-murexlabel) | `MurexLabel` (`VARCHAR(40)`) | 2 | 1 | P02 | BKMappingSec_TM_CN1 |
| [Retail UAT.dbo.EarlyRedemption_TT](tables/retail.dbo.earlyredemption_tt.md#kolom-earlyid) | `EarlyId` (`BIGINT`) | 2 | 1 | P06 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.Report_Permission](tables/retail.dbo.report_permission.md#kolom-orderbranch) | `OrderBranch` (`VARCHAR(5)`) | 2 | 1 | P01, P10 | Tidak ada key pertama di snapshot |
| [SQL_REPLICATE.dbo.CFMAST](tables/sql_replicate.dbo.cfmast.md#kolom-cfssno) | `CFSSNO` (`CHAR(40)`) | 1 | 1 | P02, P07, P08, P09, P12 | idx_cfmast_cfssno |
| [SQL_REPLICATE.dbo.CFMAST](tables/sql_replicate.dbo.cfmast.md#kolom-cfbird) | `CFBIRD` (`CHAR(7)`) | 1 | 1 | P01, P10, P12 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.REPOParamTenor_TR](tables/retail.dbo.repoparamtenor_tr.md#kolom-lastupdatenik) | `LastUpdateNIK` (`VARCHAR(10)`) | 1 | 1 | P01, P08 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TRSSettlementBIS4](tables/retail.dbo.trssettlementbis4.md#kolom-status) | `Status` (`VARCHAR(19)`) | 1 | 1 | P02, P08 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.SecurityMaster_TM](tables/retail.dbo.securitymaster_tm.md#kolom-securityno) | `SecurityNo` (`VARCHAR(20)`) | 1 | 1 | P04, P05, P06, P07 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TRSPendingRevenue_TT](tables/retail.dbo.trspendingrevenue_tt.md#kolom-kodecabangtrx) | `KodeCabangTRX` (`VARCHAR(5)`) | 1 | 1 | P01, P08 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.BONDSDetailTransaksiBuyBackPerdana_TR](tables/retail.dbo.bondsdetailtransaksibuybackperdana_tr.md#kolom-norekinvestor) | `NoRekInvestor` (`VARCHAR(100)`) | 1 | 1 | P01 | ix_BONDSDetailTransaksiBuyBackPerdana_TR_1&#124;ix_BONDSDetailTransaksiBuyBackPerdana_TR_2 |
| [Retail UAT.dbo.PTOSCustomer_TM](tables/retail.dbo.ptoscustomer_tm.md#kolom-norekinvestorrdn) | `NoRekInvestorRDN` (`VARCHAR(20)`) | 1 | 1 | P01 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.PTOSMappingProductCode_TM](tables/retail.dbo.ptosmappingproductcode_tm.md#kolom-produkptos) | `ProdukPTOS` (`VARCHAR(30)`) | 1 | 1 | P02 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.PTOSOutstanding_TR](tables/retail.dbo.ptosoutstanding_tr.md#kolom-norekinvestorrdn) | `NoRekInvestorRDN` (`VARCHAR(20)`) | 1 | 1 | P02 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.SecurityTrxLelang_TT](tables/retail.dbo.securitytrxlelang_tt.md#kolom-trxbranch) | `TrxBranch` (`VARCHAR(5)`) | 1 | 1 | P01 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.tmp_ptos_customer_upload](tables/retail.dbo.tmp_ptos_customer_upload.md#kolom-norekinvestorrdn) | `NoRekInvestorRDN` (`VARCHAR(20)`) | 1 | 1 | P01 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TreasuryCustomerLog_TH](tables/retail.dbo.treasurycustomerlog_th.md#kolom-jenisidentitas) | `JenisIdentitas` (`CHAR(1)`) | 1 | 1 | P01 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TreasuryCustomerLog_TH](tables/retail.dbo.treasurycustomerlog_th.md#kolom-jeniskelamin) | `JenisKelamin` (`CHAR(1)`) | 1 | 1 | P01 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.trs_getaccount_tmp](tables/retail.dbo.trs_getaccount_tmp.md#kolom-acctno) | `ACCTNO` (`VARCHAR(20)`) | 1 | 1 | P01 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TRSBatchTransaction_TM](tables/retail.dbo.trsbatchtransaction_tm.md#kolom-account-id-credit) | `account_id_credit` (`VARCHAR(19)`) | 1 | 1 | P01 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TRSBatchTransaction_TM](tables/retail.dbo.trsbatchtransaction_tm.md#kolom-account-id-debit) | `account_id_debit` (`VARCHAR(19)`) | 1 | 1 | P01 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TRSCTP_TM](tables/retail.dbo.trsctp_tm.md#kolom-ctpno) | `CTPNo` (`VARCHAR(30)`) | 1 | 1 | P01 | TRSCTP_TM_NN1&#124;UQ__TRSCTP_T__E9E1B8CCFFCC5E07 |
| [Retail UAT.dbo.TRSDebitAccount_TR](tables/retail.dbo.trsdebitaccount_tr.md#kolom-accountdebit) | `AccountDebit` (`VARCHAR(12)`) | 1 | 1 | P01, P08 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TRSMatchingCoupon_TT](tables/retail.dbo.trsmatchingcoupon_tt.md#kolom-norekening) | `NoRekening` (`VARCHAR(20)`) | 1 | 1 | P01 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TRSParameter_TR](tables/retail.dbo.trsparameter_tr.md#kolom-code) | `Code` (`VARCHAR(50)`) | 1 | 1 | P02 | TRSParameter_TR_CN1&#124;TRSParameter_TR_NC1 |
| [Retail UAT.dbo.TRSPostingGLBond_TT](tables/retail.dbo.trspostingglbond_tt.md#kolom-tranid) | `TranId` (`VARCHAR(20)`) | 1 | 1 | P01, P08 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TRSSettlementBIS4](tables/retail.dbo.trssettlementbis4.md#kolom-senderref) | `SenderRef` (`VARCHAR(500)`) | 1 | 1 | P02 | Tidak ada key pertama di snapshot |
| [Retail UAT.dbo.TRSSettlementKuponLog_TH](tables/retail.dbo.trssettlementkuponlog_th.md#kolom-debitrekeninggnc) | `DebitRekeningGNC` (`VARCHAR(12)`) | 1 | 1 | P01 | Tidak ada key pertama di snapshot |
| [SQL_REPLICATE.dbo.JHDATA](tables/sql_replicate.dbo.jhdata.md#kolom-jdbr) | `JDBR` (`VARCHAR(5)`) | 1 | 1 | P01 | Tidak ada key pertama di snapshot |

## Semua tabel yang muncul dalam conversion context

Bagian ini juga mencakup destination INSERT/UPDATE dan kolom pasangan yang **tidak dikonversi**. Gunakan angka A untuk memprioritaskan risiko index; jangan mengurutkan pekerjaan hanya dari total semua context.

| Tabel | Kolom terkait | Seluruh lokasi | Seluruh SP | Lokasi A | SP A | Lokasi B | Lokasi C | Partner/destination |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| [SQL_REPLICATE.dbo.PS_NDM_EMPLOYEES](tables/sql_replicate.dbo.ps_ndm_employees.md) | 5 | 129 | 26 | 59 | 20 | 43 | 10 | 17 |
| [SQL_REPLICATE.dbo.PPLMAPMasterWorkflow_TM](tables/sql_replicate.dbo.pplmapmasterworkflow_tm.md) | 2 | 45 | 21 | 40 | 18 | 3 | 0 | 2 |
| [SQL_REPLICATE.dbo.MISOfficeInformationEODFull](tables/sql_replicate.dbo.misofficeinformationeodfull.md) | 3 | 74 | 21 | 44 | 12 | 2 | 2 | 45 |
| [Retail UAT.dbo.SecurityMaster_TM](tables/retail.dbo.securitymaster_tm.md) | 61 | 788 | 176 | 25 | 11 | 107 | 355 | 306 |
| [Retail UAT.dbo.ORI_Online_TR](tables/retail.dbo.ori_online_tr.md) | 8 | 69 | 24 | 18 | 10 | 29 | 12 | 10 |
| [Retail UAT.dbo.REPOMasterTransaction_TT](tables/retail.dbo.repomastertransaction_tt.md) | 23 | 98 | 21 | 14 | 8 | 23 | 0 | 61 |
| [Retail UAT.dbo.TRSParameter_TR](tables/retail.dbo.trsparameter_tr.md) | 2 | 26 | 23 | 8 | 8 | 17 | 0 | 1 |
| [SQL_REPLICATE.dbo.CFPAR3](tables/sql_replicate.dbo.cfpar3.md) | 2 | 13 | 8 | 10 | 7 | 0 | 0 | 3 |
| [Retail UAT.dbo.SecurityCouponDate_TR](tables/retail.dbo.securitycoupondate_tr.md) | 3 | 32 | 12 | 9 | 7 | 3 | 13 | 7 |
| [Retail UAT.dbo.PTOSCustomer_TM](tables/retail.dbo.ptoscustomer_tm.md) | 18 | 64 | 13 | 22 | 5 | 0 | 3 | 39 |
| [Retail UAT.dbo.TreasuryCustomer_TM](tables/retail.dbo.treasurycustomer_tm.md) | 34 | 245 | 83 | 7 | 5 | 43 | 85 | 110 |
| [SQL_REPLICATE.dbo.ProCIFRiskProfile_TM](tables/sql_replicate.dbo.procifriskprofile_tm.md) | 3 | 14 | 6 | 6 | 5 | 2 | 0 | 6 |
| [Retail UAT.dbo.SecuritySysParameter_TR](tables/retail.dbo.securitysysparameter_tr.md) | 1 | 14 | 4 | 14 | 4 | 0 | 0 | 0 |
| [Retail UAT.dbo.TRSPostingGLBond_TH](tables/retail.dbo.trspostingglbond_th.md) | 1 | 14 | 4 | 14 | 4 | 0 | 0 | 0 |
| [Retail UAT.dbo.BONDSDetailTransaksiBuyBackPerdana_TR](tables/retail.dbo.bondsdetailtransaksibuybackperdana_tr.md) | 19 | 74 | 12 | 7 | 4 | 16 | 15 | 39 |
| [SQL_REPLICATE.dbo.DDMAST](tables/sql_replicate.dbo.ddmast.md) | 9 | 185 | 52 | 6 | 4 | 36 | 57 | 86 |
| [SQL_REPLICATE.dbo.CFMAST](tables/sql_replicate.dbo.cfmast.md) | 21 | 216 | 32 | 6 | 4 | 50 | 144 | 16 |
| [Retail UAT.dbo.ORISBNInvestorRekSB_TR](tables/retail.dbo.orisbninvestorreksb_tr.md) | 3 | 37 | 10 | 6 | 4 | 10 | 0 | 21 |
| [Retail UAT.dbo.SecurityTransaction_TT](tables/retail.dbo.securitytransaction_tt.md) | 59 | 1694 | 147 | 5 | 4 | 159 | 754 | 781 |
| [Retail UAT.dbo.PTOSTransaksi_TR](tables/retail.dbo.ptostransaksi_tr.md) | 21 | 197 | 14 | 16 | 3 | 4 | 39 | 138 |
| [Retail UAT.dbo.TRSCustAccNumber_TR](tables/retail.dbo.trscustaccnumber_tr.md) | 5 | 17 | 8 | 6 | 3 | 0 | 0 | 11 |
| [SQL_REPLICATE.dbo.CFADDR](tables/sql_replicate.dbo.cfaddr.md) | 17 | 184 | 16 | 3 | 3 | 63 | 114 | 4 |
| [Retail UAT.dbo.TRSBankCustody_TM](tables/retail.dbo.trsbankcustody_tm.md) | 2 | 7 | 4 | 3 | 3 | 4 | 0 | 0 |
| [Retail UAT.dbo.gwp_process_data_tt](tables/retail.dbo.gwp_process_data_tt.md) | 2 | 6 | 3 | 3 | 3 | 0 | 3 | 0 |
| [Retail UAT.dbo.SBNMappingSubReg_TR](tables/retail.dbo.sbnmappingsubreg_tr.md) | 2 | 13 | 3 | 3 | 3 | 1 | 7 | 2 |
| [Retail UAT.dbo.PTOSOutstanding_TR](tables/retail.dbo.ptosoutstanding_tr.md) | 14 | 156 | 16 | 21 | 2 | 10 | 51 | 74 |
| [Retail UAT.dbo.TRSSettlementBIS4_temp](tables/retail.dbo.trssettlementbis4_temp.md) | 19 | 47 | 3 | 8 | 2 | 0 | 0 | 39 |
| [Retail UAT.dbo.TRSForONEApprovalTemp_TM](tables/retail.dbo.trsforoneapprovaltemp_tm.md) | 2 | 10 | 2 | 8 | 2 | 0 | 0 | 2 |
| [Retail UAT.dbo.TRSHasilRevenue_TM](tables/retail.dbo.trshasilrevenue_tm.md) | 16 | 73 | 6 | 6 | 2 | 7 | 19 | 41 |
| [Retail UAT.dbo.TRSRekTaxAmnesty_TR](tables/retail.dbo.trsrektaxamnesty_tr.md) | 1 | 5 | 2 | 5 | 2 | 0 | 0 | 0 |
| [Retail UAT.dbo.TRSEmailOneNotifParameter_TR](tables/retail.dbo.trsemailonenotifparameter_tr.md) | 1 | 14 | 5 | 4 | 2 | 10 | 0 | 0 |
| [Retail UAT.dbo.PTOSOutstanding_Error](tables/retail.dbo.ptosoutstanding_error.md) | 9 | 49 | 2 | 4 | 2 | 0 | 0 | 45 |
| [Retail UAT.dbo.TRSGeneralParam_TR](tables/retail.dbo.trsgeneralparam_tr.md) | 2 | 15 | 8 | 2 | 2 | 1 | 0 | 12 |
| [Retail UAT.dbo.MappingOccupation_TM](tables/retail.dbo.mappingoccupation_tm.md) | 3 | 12 | 5 | 2 | 2 | 4 | 1 | 5 |
| [Retail UAT.dbo.TRSProcessOverBooking_TR](tables/retail.dbo.trsprocessoverbooking_tr.md) | 3 | 15 | 5 | 2 | 2 | 2 | 0 | 11 |
| [Retail UAT.dbo.TRSParamaterReport_TR](tables/retail.dbo.trsparamaterreport_tr.md) | 1 | 2 | 2 | 2 | 2 | 0 | 0 | 0 |
| [Retail UAT.dbo.TRSMappingMurexBIS4](tables/retail.dbo.trsmappingmurexbis4.md) | 1 | 8 | 1 | 8 | 1 | 0 | 0 | 0 |
| [Retail UAT.dbo.TRSRekSyariah_TR](tables/retail.dbo.trsreksyariah_tr.md) | 2 | 11 | 3 | 4 | 1 | 0 | 0 | 7 |
| [Retail UAT.dbo.TreasuryCustomerLog_TH](tables/retail.dbo.treasurycustomerlog_th.md) | 13 | 29 | 13 | 2 | 1 | 0 | 2 | 25 |
| [Retail UAT.dbo.EarlyRedemption_TT](tables/retail.dbo.earlyredemption_tt.md) | 11 | 36 | 12 | 2 | 1 | 5 | 13 | 16 |
| [Retail UAT.dbo.SecurityTransferAset_TM](tables/retail.dbo.securitytransferaset_tm.md) | 9 | 56 | 8 | 2 | 1 | 22 | 6 | 26 |
| [Retail UAT.dbo.Report_Permission](tables/retail.dbo.report_permission.md) | 2 | 19 | 6 | 2 | 1 | 8 | 4 | 5 |
| [Retail UAT.dbo.TRSBatchTransaction_TM](tables/retail.dbo.trsbatchtransaction_tm.md) | 6 | 13 | 5 | 2 | 1 | 1 | 1 | 9 |
| [Retail UAT.dbo.TRSSettlementBIS4](tables/retail.dbo.trssettlementbis4.md) | 13 | 52 | 4 | 2 | 1 | 39 | 8 | 3 |
| [Retail UAT.dbo.BKMappingSec_TM](tables/retail.dbo.bkmappingsec_tm.md) | 1 | 2 | 1 | 2 | 1 | 0 | 0 | 0 |
| [Retail UAT.dbo.TRSMatchingCoupon_TT](tables/retail.dbo.trsmatchingcoupon_tt.md) | 9 | 44 | 9 | 1 | 1 | 2 | 28 | 13 |
| [Retail UAT.dbo.SecurityTrxLelang_TT](tables/retail.dbo.securitytrxlelang_tt.md) | 7 | 19 | 8 | 1 | 1 | 8 | 5 | 5 |
| [Retail UAT.dbo.TRSSettlementKuponLog_TH](tables/retail.dbo.trssettlementkuponlog_th.md) | 9 | 36 | 8 | 1 | 1 | 2 | 23 | 10 |
| [Retail UAT.dbo.REPOParamTenor_TR](tables/retail.dbo.repoparamtenor_tr.md) | 3 | 19 | 6 | 1 | 1 | 1 | 4 | 13 |
| [Retail UAT.dbo.TRSCTP_TM](tables/retail.dbo.trsctp_tm.md) | 3 | 19 | 5 | 1 | 1 | 0 | 3 | 15 |
| [Retail UAT.dbo.PTOSMappingProductCode_TM](tables/retail.dbo.ptosmappingproductcode_tm.md) | 2 | 2 | 2 | 1 | 1 | 0 | 0 | 1 |
| [Retail UAT.dbo.tmp_ptos_customer_upload](tables/retail.dbo.tmp_ptos_customer_upload.md) | 7 | 23 | 2 | 1 | 1 | 1 | 9 | 12 |
| [Retail UAT.dbo.TRSPendingRevenue_TT](tables/retail.dbo.trspendingrevenue_tt.md) | 3 | 10 | 2 | 1 | 1 | 1 | 4 | 4 |
| [Retail UAT.dbo.trs_getaccount_tmp](tables/retail.dbo.trs_getaccount_tmp.md) | 3 | 3 | 1 | 1 | 1 | 0 | 0 | 2 |
| [Retail UAT.dbo.TRSDebitAccount_TR](tables/retail.dbo.trsdebitaccount_tr.md) | 3 | 7 | 1 | 1 | 1 | 1 | 1 | 5 |
| [Retail UAT.dbo.TRSPostingGLBond_TT](tables/retail.dbo.trspostingglbond_tt.md) | 2 | 3 | 1 | 1 | 1 | 0 | 0 | 2 |
| [SQL_REPLICATE.dbo.JHDATA](tables/sql_replicate.dbo.jhdata.md) | 1 | 1 | 1 | 1 | 1 | 0 | 0 | 0 |
| [Retail UAT.dbo.control_table](tables/retail.dbo.control_table.md) | 7 | 120 | 79 | 0 | 0 | 80 | 35 | 5 |
| [Retail UAT.dbo.ORI_Order_TT](tables/retail.dbo.ori_order_tt.md) | 20 | 204 | 51 | 0 | 0 | 16 | 75 | 113 |
| [Retail UAT.dbo.TRSParamHargaModal_TR](tables/retail.dbo.trsparamhargamodal_tr.md) | 10 | 134 | 29 | 0 | 0 | 5 | 7 | 122 |
| [Retail UAT.dbo.TRSRetailCustBlock_TM](tables/retail.dbo.trsretailcustblock_tm.md) | 4 | 90 | 22 | 0 | 0 | 2 | 62 | 26 |
| [Retail UAT.dbo.RiskProfileTRS_TM](tables/retail.dbo.riskprofiletrs_tm.md) | 5 | 119 | 21 | 0 | 0 | 0 | 96 | 23 |
| [Retail UAT.dbo.TRSMateraiTransaction_TT](tables/retail.dbo.trsmateraitransaction_tt.md) | 12 | 85 | 20 | 0 | 0 | 6 | 14 | 65 |
| [Retail UAT.dbo.ORIOnlineGeneralParam_TR](tables/retail.dbo.orionlinegeneralparam_tr.md) | 2 | 31 | 17 | 0 | 0 | 25 | 0 | 6 |
| [SQL_REPLICATE.dbo.SIBSCurrencyTable_TM](tables/sql_replicate.dbo.sibscurrencytable_tm.md) | 6 | 127 | 17 | 0 | 0 | 61 | 2 | 64 |
| [SQL_REPLICATE.dbo.TAMAST](tables/sql_replicate.dbo.tamast.md) | 2 | 41 | 17 | 0 | 0 | 0 | 0 | 41 |
| [Retail UAT.dbo.ORI_Order_Temp](tables/retail.dbo.ori_order_temp.md) | 11 | 54 | 16 | 0 | 0 | 7 | 9 | 38 |
| [SQL_REPLICATE.dbo.DDPAR2](tables/sql_replicate.dbo.ddpar2.md) | 5 | 86 | 15 | 0 | 0 | 2 | 47 | 37 |
| [Retail UAT.dbo.SecurityTransactionLink_TR](tables/retail.dbo.securitytransactionlink_tr.md) | 6 | 54 | 14 | 0 | 0 | 6 | 17 | 31 |
| [SQL_REPLICATE.dbo.DDTNEW](tables/sql_replicate.dbo.ddtnew.md) | 7 | 39 | 14 | 0 | 0 | 17 | 0 | 22 |
| [Retail UAT.dbo.TRSSafeKeepingFeeLog_TH](tables/retail.dbo.trssafekeepingfeelog_th.md) | 3 | 25 | 13 | 0 | 0 | 0 | 19 | 6 |
| [Retail UAT.dbo.TRSECParamSpread_TR](tables/retail.dbo.trsecparamspread_tr.md) | 11 | 62 | 12 | 0 | 0 | 0 | 22 | 40 |
| [Retail UAT.dbo.TRSParamBookBld_TR](tables/retail.dbo.trsparambookbld_tr.md) | 6 | 52 | 12 | 0 | 0 | 0 | 24 | 28 |
| [SQL_REPLICATE.dbo.JHCOUN](tables/sql_replicate.dbo.jhcoun.md) | 3 | 30 | 12 | 0 | 0 | 3 | 10 | 17 |
| [SQL_REPLICATE.dbo.PROOBL_I](tables/sql_replicate.dbo.proobl_i.md) | 2 | 22 | 12 | 0 | 0 | 0 | 0 | 22 |
| [Retail UAT.dbo.ORI_Order_Log_TH](tables/retail.dbo.ori_order_log_th.md) | 2 | 16 | 10 | 0 | 0 | 0 | 4 | 12 |
| [Retail UAT.dbo.SecurityTransaction_TR](tables/retail.dbo.securitytransaction_tr.md) | 1 | 25 | 10 | 0 | 0 | 3 | 21 | 1 |
| [Retail UAT.dbo.TRSParamSpread_TR](tables/retail.dbo.trsparamspread_tr.md) | 2 | 49 | 10 | 0 | 0 | 0 | 0 | 49 |
| [Retail UAT.dbo.ORISBNOrderMainSched_TM](tables/retail.dbo.orisbnordermainsched_tm.md) | 6 | 12 | 9 | 0 | 0 | 4 | 0 | 8 |
| [SQL_REPLICATE.dbo.BIS4_SUB_ACCOUNTS](tables/sql_replicate.dbo.bis4_sub_accounts.md) | 3 | 30 | 9 | 0 | 0 | 3 | 0 | 27 |
| [SQL_REPLICATE.dbo.CFACCT](tables/sql_replicate.dbo.cfacct.md) | 5 | 40 | 9 | 0 | 0 | 15 | 11 | 14 |
| [SQL_REPLICATE.dbo.CFAIDN](tables/sql_replicate.dbo.cfaidn.md) | 7 | 37 | 9 | 0 | 0 | 13 | 20 | 4 |
| [Retail UAT.dbo.SecurityMasterLog_TH](tables/retail.dbo.securitymasterlog_th.md) | 2 | 23 | 8 | 0 | 0 | 0 | 13 | 10 |
| [Retail UAT.dbo.SecurityTransaction_TH](tables/retail.dbo.securitytransaction_th.md) | 17 | 23 | 8 | 0 | 0 | 0 | 1 | 22 |
| [Retail UAT.dbo.SecurityTransactionBuy_TT](tables/retail.dbo.securitytransactionbuy_tt.md) | 12 | 56 | 8 | 0 | 0 | 0 | 4 | 52 |
| [Retail UAT.dbo.TRSBatchHeaderFile2CMS_TM](tables/retail.dbo.trsbatchheaderfile2cms_tm.md) | 3 | 38 | 8 | 0 | 0 | 0 | 18 | 20 |
| [Retail UAT.dbo.TRSEditHargaModal_TH](tables/retail.dbo.trsedithargamodal_th.md) | 3 | 21 | 8 | 0 | 0 | 0 | 10 | 11 |
| [Retail UAT.dbo.TRSJobSendEmail_TT](tables/retail.dbo.trsjobsendemail_tt.md) | 2 | 15 | 8 | 0 | 0 | 0 | 0 | 15 |
| [Retail UAT.dbo.TRSSeller_TM](tables/retail.dbo.trsseller_tm.md) | 3 | 29 | 8 | 0 | 0 | 0 | 22 | 7 |
| [Retail UAT.dbo.CFADDRKodeKotaMapping_TR](tables/retail.dbo.cfaddrkodekotamapping_tr.md) | 1 | 7 | 7 | 0 | 0 | 0 | 7 | 0 |
| [Retail UAT.dbo.ORISBNInvestorRekDana_TR](tables/retail.dbo.orisbninvestorrekdana_tr.md) | 3 | 15 | 7 | 0 | 0 | 4 | 0 | 11 |
| [Retail UAT.dbo.PTOSOutstanding_Temp](tables/retail.dbo.ptosoutstanding_temp.md) | 4 | 23 | 7 | 0 | 0 | 0 | 6 | 17 |
| [Retail UAT.dbo.REPOParameterProduct_TM](tables/retail.dbo.repoparameterproduct_tm.md) | 9 | 55 | 7 | 0 | 0 | 17 | 16 | 22 |
| [Retail UAT.dbo.SBNSchedMail_TT](tables/retail.dbo.sbnschedmail_tt.md) | 2 | 10 | 7 | 0 | 0 | 0 | 3 | 7 |
| [Retail UAT.dbo.SBNSchedNotification_TT](tables/retail.dbo.sbnschednotification_tt.md) | 2 | 10 | 7 | 0 | 0 | 0 | 3 | 7 |
| [Retail UAT.dbo.TreasuryCustomerPending_TM](tables/retail.dbo.treasurycustomerpending_tm.md) | 11 | 23 | 7 | 0 | 0 | 3 | 5 | 15 |
| [Retail UAT.dbo.TRSMaturity_TT](tables/retail.dbo.trsmaturity_tt.md) | 17 | 36 | 7 | 0 | 0 | 2 | 18 | 16 |
| [Retail UAT.dbo.TRSStatusCorporateBond_TR](tables/retail.dbo.trsstatuscorporatebond_tr.md) | 1 | 9 | 7 | 0 | 0 | 0 | 4 | 5 |
| [SQL_REPLICATE.dbo.CFZEMP](tables/sql_replicate.dbo.cfzemp.md) | 3 | 15 | 7 | 0 | 0 | 0 | 15 | 0 |
| [Retail UAT.dbo.KodeKota_TR](tables/retail.dbo.kodekota_tr.md) | 1 | 6 | 6 | 0 | 0 | 0 | 0 | 6 |
| [Retail UAT.dbo.process_table](tables/retail.dbo.process_table.md) | 4 | 20 | 6 | 0 | 0 | 3 | 7 | 10 |
| [Retail UAT.dbo.PTOSLogHasilBMW_TH](tables/retail.dbo.ptosloghasilbmw_th.md) | 2 | 12 | 6 | 0 | 0 | 2 | 0 | 10 |
| [Retail UAT.dbo.PTOSTransaksi_Temp](tables/retail.dbo.ptostransaksi_temp.md) | 7 | 38 | 6 | 0 | 0 | 0 | 8 | 30 |
| [Retail UAT.dbo.REPOParamHargaModal_TR](tables/retail.dbo.repoparamhargamodal_tr.md) | 3 | 10 | 6 | 0 | 0 | 1 | 0 | 9 |
| [Retail UAT.dbo.SecurityMasterExtMig_TM](tables/retail.dbo.securitymasterextmig_tm.md) | 2 | 26 | 6 | 0 | 0 | 0 | 13 | 13 |
| [Retail UAT.dbo.TreasuryGeneralParam_TR](tables/retail.dbo.treasurygeneralparam_tr.md) | 3 | 8 | 6 | 0 | 0 | 2 | 0 | 6 |
| [Retail UAT.dbo.TreasuryStoreProcedureLog_TL](tables/retail.dbo.treasurystoreprocedurelog_tl.md) | 4 | 13 | 6 | 0 | 0 | 0 | 0 | 13 |
| [Retail UAT.dbo.TRSMaturityTaxAmnesty_TT](tables/retail.dbo.trsmaturitytaxamnesty_tt.md) | 3 | 17 | 6 | 0 | 0 | 0 | 9 | 8 |
| [Retail UAT.dbo.TRSParamHargaModal_TL](tables/retail.dbo.trsparamhargamodal_tl.md) | 5 | 18 | 6 | 0 | 0 | 0 | 5 | 13 |
| [Retail UAT.dbo.TRSParamHargaModalPending_TR](tables/retail.dbo.trsparamhargamodalpending_tr.md) | 6 | 10 | 6 | 0 | 0 | 1 | 0 | 9 |
| [Retail UAT.dbo.TRSSecurityCurrency_TR](tables/retail.dbo.trssecuritycurrency_tr.md) | 2 | 11 | 6 | 0 | 0 | 0 | 9 | 2 |
| [Retail UAT.dbo.TRSSettlementMaturityLog_TH](tables/retail.dbo.trssettlementmaturitylog_th.md) | 4 | 31 | 6 | 0 | 0 | 1 | 28 | 2 |
| [Retail UAT.dbo.TRSSettlementOrderLog_TH](tables/retail.dbo.trssettlementorderlog_th.md) | 4 | 14 | 6 | 0 | 0 | 0 | 9 | 5 |
| [SQL_REPLICATE.dbo.TCLOGT](tables/sql_replicate.dbo.tclogt.md) | 3 | 20 | 6 | 0 | 0 | 2 | 0 | 18 |
| [Retail UAT.dbo.ORI_Order_Pending_TT](tables/retail.dbo.ori_order_pending_tt.md) | 2 | 13 | 5 | 0 | 0 | 0 | 8 | 5 |
| [Retail UAT.dbo.SecurityFeeNasabah_TR](tables/retail.dbo.securityfeenasabah_tr.md) | 10 | 39 | 5 | 0 | 0 | 3 | 2 | 34 |
| [Retail UAT.dbo.SecuritySafeKeepingFee_TR](tables/retail.dbo.securitysafekeepingfee_tr.md) | 14 | 94 | 5 | 0 | 0 | 6 | 14 | 78 |
| [Retail UAT.dbo.SecurityTransferAsetLink_TR](tables/retail.dbo.securitytransferasetlink_tr.md) | 3 | 21 | 5 | 0 | 0 | 9 | 8 | 4 |
| [Retail UAT.dbo.TRSBondsPrice_TM](tables/retail.dbo.trsbondsprice_tm.md) | 6 | 14 | 5 | 0 | 0 | 2 | 0 | 12 |
| [Retail UAT.dbo.TRSMainApproval_TM](tables/retail.dbo.trsmainapproval_tm.md) | 2 | 8 | 5 | 0 | 0 | 1 | 0 | 7 |
| [Retail UAT.dbo.TRSParamHargaModal_TH](tables/retail.dbo.trsparamhargamodal_th.md) | 2 | 12 | 5 | 0 | 0 | 0 | 0 | 12 |
| [Retail UAT.dbo.TRSProcessBuyBackLog_TH](tables/retail.dbo.trsprocessbuybacklog_th.md) | 5 | 15 | 5 | 0 | 0 | 0 | 6 | 9 |
| [Retail UAT.dbo.TRSSeller_TH](tables/retail.dbo.trsseller_th.md) | 2 | 32 | 5 | 0 | 0 | 0 | 26 | 6 |
| [Retail UAT.dbo.TRSSettlementTransaksiBankBeliLog_TH](tables/retail.dbo.trssettlementtransaksibankbelilog_th.md) | 1 | 26 | 5 | 0 | 0 | 0 | 24 | 2 |
| [Retail UAT.dbo.TRSSettlementTransaksiBankJualLog_TH](tables/retail.dbo.trssettlementtransaksibankjuallog_th.md) | 1 | 15 | 5 | 0 | 0 | 0 | 14 | 1 |
| [Retail UAT.dbo.coupon_source_tt](tables/retail.dbo.coupon_source_tt.md) | 5 | 34 | 4 | 0 | 0 | 0 | 10 | 24 |
| [Retail UAT.dbo.CustodySfeeKupon_TR](tables/retail.dbo.custodysfeekupon_tr.md) | 6 | 22 | 4 | 0 | 0 | 6 | 14 | 4 |
| [Retail UAT.dbo.CustodyUploadSfeeKupon_TR](tables/retail.dbo.custodyuploadsfeekupon_tr.md) | 2 | 6 | 4 | 0 | 0 | 1 | 1 | 4 |
| [Retail UAT.dbo.MappingFundSource_TM](tables/retail.dbo.mappingfundsource_tm.md) | 2 | 4 | 4 | 0 | 0 | 3 | 1 | 0 |
| [Retail UAT.dbo.maturity_source_tt](tables/retail.dbo.maturity_source_tt.md) | 20 | 57 | 4 | 0 | 0 | 1 | 13 | 43 |
| [Retail UAT.dbo.ParamStatus_TR](tables/retail.dbo.paramstatus_tr.md) | 2 | 7 | 4 | 0 | 0 | 0 | 5 | 2 |
| [Retail UAT.dbo.PTOSOutstanding_TH](tables/retail.dbo.ptosoutstanding_th.md) | 13 | 39 | 4 | 0 | 0 | 0 | 0 | 39 |
| [Retail UAT.dbo.PTOSTransaksi_TH](tables/retail.dbo.ptostransaksi_th.md) | 21 | 53 | 4 | 0 | 0 | 0 | 0 | 53 |
| [Retail UAT.dbo.SecurityFeeNasabah_TH](tables/retail.dbo.securityfeenasabah_th.md) | 2 | 11 | 4 | 0 | 0 | 0 | 5 | 6 |
| [Retail UAT.dbo.SecurityMasterPending_TM](tables/retail.dbo.securitymasterpending_tm.md) | 19 | 57 | 4 | 0 | 0 | 5 | 32 | 20 |
| [Retail UAT.dbo.SecurityTransactionDetail_TT](tables/retail.dbo.securitytransactiondetail_tt.md) | 5 | 8 | 4 | 0 | 0 | 6 | 2 | 0 |
| [Retail UAT.dbo.SecurityTrxResultLelang_TT](tables/retail.dbo.securitytrxresultlelang_tt.md) | 2 | 6 | 4 | 0 | 0 | 0 | 0 | 6 |
| [Retail UAT.dbo.TipeInvestor_TM](tables/retail.dbo.tipeinvestor_tm.md) | 1 | 4 | 4 | 0 | 0 | 0 | 0 | 4 |
| [Retail UAT.dbo.tmp_ptos_outstanding_upload](tables/retail.dbo.tmp_ptos_outstanding_upload.md) | 5 | 30 | 4 | 0 | 0 | 2 | 14 | 14 |
| [Retail UAT.dbo.TRSAuthParamIssuer_TH](tables/retail.dbo.trsauthparamissuer_th.md) | 1 | 7 | 4 | 0 | 0 | 0 | 6 | 1 |
| [Retail UAT.dbo.TRSBankCustody_TH](tables/retail.dbo.trsbankcustody_th.md) | 3 | 11 | 4 | 0 | 0 | 0 | 5 | 6 |
| [Retail UAT.dbo.TRSBankCustodyPending_TM](tables/retail.dbo.trsbankcustodypending_tm.md) | 3 | 42 | 4 | 0 | 0 | 0 | 36 | 6 |
| [Retail UAT.dbo.TRSCPT_TM](tables/retail.dbo.trscpt_tm.md) | 3 | 4 | 4 | 0 | 0 | 0 | 0 | 4 |
| [Retail UAT.dbo.TRSCustomerRepoFunded_TR](tables/retail.dbo.trscustomerrepofunded_tr.md) | 1 | 10 | 4 | 0 | 0 | 0 | 3 | 7 |
| [Retail UAT.dbo.TRSECBondsPrice_TM](tables/retail.dbo.trsecbondsprice_tm.md) | 10 | 24 | 4 | 0 | 0 | 0 | 0 | 24 |
| [Retail UAT.dbo.TRSPajakBeliLog_TH](tables/retail.dbo.trspajakbelilog_th.md) | 1 | 11 | 4 | 0 | 0 | 0 | 10 | 1 |
| [Retail UAT.dbo.TRSParameterGL_TH](tables/retail.dbo.trsparametergl_th.md) | 4 | 30 | 4 | 0 | 0 | 0 | 18 | 12 |
| [Retail UAT.dbo.TRSParamLelang_TR](tables/retail.dbo.trsparamlelang_tr.md) | 11 | 23 | 4 | 0 | 0 | 1 | 0 | 22 |
| [Retail UAT.dbo.TRSPendingAccruedInterest_TH](tables/retail.dbo.trspendingaccruedinterest_th.md) | 1 | 8 | 4 | 0 | 0 | 0 | 5 | 3 |
| [Retail UAT.dbo.TRSProcessTransaksi_TT](tables/retail.dbo.trsprocesstransaksi_tt.md) | 3 | 7 | 4 | 0 | 0 | 0 | 0 | 7 |
| [Retail UAT.dbo.TRSRefundLog_TH](tables/retail.dbo.trsrefundlog_th.md) | 1 | 8 | 4 | 0 | 0 | 0 | 8 | 0 |
| [Retail UAT.dbo.TRSSettlementBuyBackLog_TH](tables/retail.dbo.trssettlementbuybacklog_th.md) | 4 | 11 | 4 | 0 | 0 | 0 | 5 | 6 |
| [Retail UAT.dbo.TRSSettlementCashBackLog_TH](tables/retail.dbo.trssettlementcashbacklog_th.md) | 2 | 10 | 4 | 0 | 0 | 0 | 8 | 2 |
| [Retail UAT.dbo.TRSSettlementEarlyLog_TH](tables/retail.dbo.trssettlementearlylog_th.md) | 1 | 10 | 4 | 0 | 0 | 0 | 9 | 1 |
| [Retail UAT.dbo.TRSSettlementTransaksiBankBeliTALog_TH](tables/retail.dbo.trssettlementtransaksibankbelitalog_th.md) | 1 | 14 | 4 | 0 | 0 | 0 | 13 | 1 |
| [Retail UAT.dbo.TRSTransactionFeeLogBeli_TR](tables/retail.dbo.trstransactionfeelogbeli_tr.md) | 1 | 9 | 4 | 0 | 0 | 0 | 8 | 1 |
| [SQL_REPLICATE.dbo.CFALTN](tables/sql_replicate.dbo.cfaltn.md) | 2 | 18 | 4 | 0 | 0 | 9 | 0 | 9 |
| [SQL_REPLICATE.dbo.JHFXDT](tables/sql_replicate.dbo.jhfxdt.md) | 2 | 23 | 4 | 0 | 0 | 3 | 0 | 20 |
| [SQL_REPLICATE.dbo.MXHolidayCurrency](tables/sql_replicate.dbo.mxholidaycurrency.md) | 1 | 15 | 4 | 0 | 0 | 0 | 0 | 15 |
| [SQL_REPLICATE.dbo.PROOBL_O](tables/sql_replicate.dbo.proobl_o.md) | 1 | 5 | 4 | 0 | 0 | 5 | 0 | 0 |
| [Retail UAT.dbo.dump_coupon_daily_tt](tables/retail.dbo.dump_coupon_daily_tt.md) | 2 | 4 | 3 | 0 | 0 | 1 | 0 | 3 |
| [Retail UAT.dbo.dump_maturity_daily_tt](tables/retail.dbo.dump_maturity_daily_tt.md) | 8 | 15 | 3 | 0 | 0 | 2 | 0 | 13 |
| [Retail UAT.dbo.MappingIncome_TM](tables/retail.dbo.mappingincome_tm.md) | 1 | 3 | 3 | 0 | 0 | 0 | 0 | 3 |
| [Retail UAT.dbo.ORISchedWorkflow_TR](tables/retail.dbo.orischedworkflow_tr.md) | 3 | 10 | 3 | 0 | 0 | 0 | 0 | 10 |
| [Retail UAT.dbo.oriva_security_master](tables/retail.dbo.oriva_security_master.md) | 2 | 4 | 3 | 0 | 0 | 0 | 2 | 2 |
| [Retail UAT.dbo.oriva_settlement_log](tables/retail.dbo.oriva_settlement_log.md) | 2 | 17 | 3 | 0 | 0 | 0 | 0 | 17 |
| [Retail UAT.dbo.PTOSCustomer_TH](tables/retail.dbo.ptoscustomer_th.md) | 3 | 12 | 3 | 0 | 0 | 0 | 0 | 12 |
| [Retail UAT.dbo.PTOSSecurityMaster_Temp](tables/retail.dbo.ptossecuritymaster_temp.md) | 2 | 20 | 3 | 0 | 0 | 0 | 16 | 4 |
| [Retail UAT.dbo.REPOApprovalParameterDetail_TR](tables/retail.dbo.repoapprovalparameterdetail_tr.md) | 6 | 18 | 3 | 0 | 0 | 0 | 11 | 7 |
| [Retail UAT.dbo.REPOParamMaturityDate_TR](tables/retail.dbo.repoparammaturitydate_tr.md) | 4 | 8 | 3 | 0 | 0 | 1 | 0 | 7 |
| [Retail UAT.dbo.REPOParamTenor_TH](tables/retail.dbo.repoparamtenor_th.md) | 3 | 5 | 3 | 0 | 0 | 1 | 0 | 4 |
| [Retail UAT.dbo.REPOProductType_TM](tables/retail.dbo.repoproducttype_tm.md) | 3 | 5 | 3 | 0 | 0 | 0 | 4 | 1 |
| [Retail UAT.dbo.SecurityFeeNasabahDetail_TH](tables/retail.dbo.securityfeenasabahdetail_th.md) | 1 | 4 | 3 | 0 | 0 | 0 | 2 | 2 |
| [Retail UAT.dbo.SecurityNoBankCustody_TH](tables/retail.dbo.securitynobankcustody_th.md) | 2 | 11 | 3 | 0 | 0 | 0 | 3 | 8 |
| [Retail UAT.dbo.SecurityTransferAset_TH](tables/retail.dbo.securitytransferaset_th.md) | 1 | 6 | 3 | 0 | 0 | 0 | 0 | 6 |
| [Retail UAT.dbo.TreasuryCustomerChange_TH](tables/retail.dbo.treasurycustomerchange_th.md) | 3 | 15 | 3 | 0 | 0 | 0 | 5 | 10 |
| [Retail UAT.dbo.TRSBatchDetailFile2CMS_TM](tables/retail.dbo.trsbatchdetailfile2cms_tm.md) | 11 | 86 | 3 | 0 | 0 | 0 | 0 | 86 |
| [Retail UAT.dbo.TRSBondsPrice_TH](tables/retail.dbo.trsbondsprice_th.md) | 5 | 8 | 3 | 0 | 0 | 0 | 0 | 8 |
| [Retail UAT.dbo.TRSCustAccNumberLog_TL](tables/retail.dbo.trscustaccnumberlog_tl.md) | 2 | 16 | 3 | 0 | 0 | 0 | 0 | 16 |
| [Retail UAT.dbo.TRSCustodyUploadMonitor_TR](tables/retail.dbo.trscustodyuploadmonitor_tr.md) | 3 | 7 | 3 | 0 | 0 | 0 | 6 | 1 |
| [Retail UAT.dbo.TRSExcludeProductCode_TR](tables/retail.dbo.trsexcludeproductcode_tr.md) | 1 | 3 | 3 | 0 | 0 | 0 | 0 | 3 |
| [Retail UAT.dbo.TRSJobSendEmailRecipientDetail_TR](tables/retail.dbo.trsjobsendemailrecipientdetail_tr.md) | 1 | 3 | 3 | 0 | 0 | 0 | 0 | 3 |
| [Retail UAT.dbo.TRSListChannelBonds_TR](tables/retail.dbo.trslistchannelbonds_tr.md) | 3 | 3 | 3 | 0 | 0 | 0 | 2 | 1 |
| [Retail UAT.dbo.TRSMateraiTransaction_TH](tables/retail.dbo.trsmateraitransaction_th.md) | 2 | 4 | 3 | 0 | 0 | 0 | 0 | 4 |
| [Retail UAT.dbo.TRSOnCallDepositListRek_TMP](tables/retail.dbo.trsoncalldepositlistrek_tmp.md) | 3 | 9 | 3 | 0 | 0 | 0 | 0 | 9 |
| [Retail UAT.dbo.TRSPajakkLog_TH](tables/retail.dbo.trspajakklog_th.md) | 4 | 7 | 3 | 0 | 0 | 0 | 1 | 6 |
| [Retail UAT.dbo.TRSParamAccruedDaysAdjustment_TH](tables/retail.dbo.trsparamaccrueddaysadjustment_th.md) | 2 | 12 | 3 | 0 | 0 | 0 | 3 | 9 |
| [Retail UAT.dbo.TRSParamCounter_TM](tables/retail.dbo.trsparamcounter_tm.md) | 1 | 5 | 3 | 0 | 0 | 1 | 0 | 4 |
| [Retail UAT.dbo.TRSParamDeviasiHarga_TL](tables/retail.dbo.trsparamdeviasiharga_tl.md) | 2 | 14 | 3 | 0 | 0 | 0 | 14 | 0 |
| [Retail UAT.dbo.TRSParamDeviasiTotalSpread_TR_Temp](tables/retail.dbo.trsparamdeviasitotalspread_tr_temp.md) | 4 | 12 | 3 | 0 | 0 | 2 | 0 | 10 |
| [Retail UAT.dbo.TRSParameterGLLog_TH](tables/retail.dbo.trsparametergllog_th.md) | 6 | 32 | 3 | 0 | 0 | 0 | 3 | 29 |
| [Retail UAT.dbo.TRSParamSpread_TL](tables/retail.dbo.trsparamspread_tl.md) | 2 | 18 | 3 | 0 | 0 | 0 | 15 | 3 |
| [Retail UAT.dbo.TRSParamSpreadHargaModal_TR](tables/retail.dbo.trsparamspreadhargamodal_tr.md) | 3 | 11 | 3 | 0 | 0 | 0 | 0 | 11 |
| [Retail UAT.dbo.TRSProcessTransaksiDetail_TR](tables/retail.dbo.trsprocesstransaksidetail_tr.md) | 2 | 4 | 3 | 0 | 0 | 1 | 0 | 3 |
| [Retail UAT.dbo.TRSSecurityTransactionPreBook_TT](tables/retail.dbo.trssecuritytransactionprebook_tt.md) | 2 | 6 | 3 | 0 | 0 | 3 | 2 | 1 |
| [Retail UAT.dbo.TRSSellerPending_TM](tables/retail.dbo.trssellerpending_tm.md) | 3 | 11 | 3 | 0 | 0 | 0 | 6 | 5 |
| [Retail UAT.dbo.TRSSimanisTransaction_TT](tables/retail.dbo.trssimanistransaction_tt.md) | 1 | 6 | 3 | 0 | 0 | 0 | 0 | 6 |
| [Retail UAT.dbo.TRSTransaksiBuyBack_TT](tables/retail.dbo.trstransaksibuyback_tt.md) | 4 | 6 | 3 | 0 | 0 | 1 | 0 | 5 |
| [Retail UAT.dbo.TRSUntagSBNAccApproval_TT](tables/retail.dbo.trsuntagsbnaccapproval_tt.md) | 2 | 6 | 3 | 0 | 0 | 0 | 5 | 1 |
| [Retail UAT.dbo.TRSUploadSID_TH](tables/retail.dbo.trsuploadsid_th.md) | 2 | 18 | 3 | 0 | 0 | 0 | 8 | 10 |
| [SQL_REPLICATE.dbo.CFCONN](tables/sql_replicate.dbo.cfconn.md) | 3 | 67 | 3 | 0 | 0 | 3 | 64 | 0 |
| [SQL_REPLICATE.dbo.EACustomerFlagging_TM](tables/sql_replicate.dbo.eacustomerflagging_tm.md) | 1 | 4 | 3 | 0 | 0 | 1 | 0 | 3 |
| [SQL_REPLICATE.dbo.PBFlagDesc_TR](tables/sql_replicate.dbo.pbflagdesc_tr.md) | 1 | 3 | 3 | 0 | 0 | 0 | 3 | 0 |
| [Retail UAT.dbo.acf_data_tt](tables/retail.dbo.acf_data_tt.md) | 1 | 2 | 2 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.CIFReplicateLog_TH](tables/retail.dbo.cifreplicatelog_th.md) | 3 | 3 | 2 | 0 | 0 | 0 | 0 | 3 |
| [Retail UAT.dbo.CIFReplicatePending_TT](tables/retail.dbo.cifreplicatepending_tt.md) | 3 | 4 | 2 | 0 | 0 | 0 | 1 | 3 |
| [Retail UAT.dbo.CouponUploadTemp](tables/retail.dbo.couponuploadtemp.md) | 11 | 22 | 2 | 0 | 0 | 0 | 0 | 22 |
| [Retail UAT.dbo.EarlyRedemptionOnline_TR](tables/retail.dbo.earlyredemptiononline_tr.md) | 1 | 5 | 2 | 0 | 0 | 0 | 0 | 5 |
| [Retail UAT.dbo.gwp_sched_process_log_th](tables/retail.dbo.gwp_sched_process_log_th.md) | 1 | 2 | 2 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.MappingKewarganegaraan_TM](tables/retail.dbo.mappingkewarganegaraan_tm.md) | 2 | 4 | 2 | 0 | 0 | 2 | 0 | 2 |
| [Retail UAT.dbo.maturity_upload_temp](tables/retail.dbo.maturity_upload_temp.md) | 15 | 30 | 2 | 0 | 0 | 0 | 0 | 30 |
| [Retail UAT.dbo.MxBondTran](tables/retail.dbo.mxbondtran.md) | 4 | 5 | 2 | 0 | 0 | 0 | 0 | 5 |
| [Retail UAT.dbo.ORISBNInvestorMaster_TM](tables/retail.dbo.orisbninvestormaster_tm.md) | 3 | 7 | 2 | 0 | 0 | 0 | 0 | 7 |
| [Retail UAT.dbo.ORISBNInvestorPending_TT](tables/retail.dbo.orisbninvestorpending_tt.md) | 2 | 4 | 2 | 0 | 0 | 4 | 0 | 0 |
| [Retail UAT.dbo.ORISBNInvestorRekDana_TH](tables/retail.dbo.orisbninvestorrekdana_th.md) | 2 | 5 | 2 | 0 | 0 | 0 | 0 | 5 |
| [Retail UAT.dbo.PTBCExcludeProductHargaModal_TT](tables/retail.dbo.ptbcexcludeproducthargamodal_tt.md) | 1 | 3 | 2 | 0 | 0 | 0 | 0 | 3 |
| [Retail UAT.dbo.PTOSCustomer_Temp](tables/retail.dbo.ptoscustomer_temp.md) | 2 | 2 | 2 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.PTOSNasabah_Excel](tables/retail.dbo.ptosnasabah_excel.md) | 16 | 43 | 2 | 0 | 0 | 15 | 0 | 28 |
| [Retail UAT.dbo.PTOSOutstandingBond_Excel](tables/retail.dbo.ptosoutstandingbond_excel.md) | 6 | 26 | 2 | 0 | 0 | 0 | 0 | 26 |
| [Retail UAT.dbo.PTOSOutstandingEquity_Excel](tables/retail.dbo.ptosoutstandingequity_excel.md) | 6 | 24 | 2 | 0 | 0 | 0 | 0 | 24 |
| [Retail UAT.dbo.PTOSOutstandingMaster_TM](tables/retail.dbo.ptosoutstandingmaster_tm.md) | 2 | 4 | 2 | 0 | 0 | 0 | 0 | 4 |
| [Retail UAT.dbo.PTOSTransaksi_Error](tables/retail.dbo.ptostransaksi_error.md) | 17 | 68 | 2 | 0 | 0 | 0 | 0 | 68 |
| [Retail UAT.dbo.PTOSTransaksiBond_Excel](tables/retail.dbo.ptostransaksibond_excel.md) | 7 | 25 | 2 | 0 | 0 | 0 | 0 | 25 |
| [Retail UAT.dbo.PTOSTransaksiEquity_Excel](tables/retail.dbo.ptostransaksiequity_excel.md) | 8 | 27 | 2 | 0 | 0 | 0 | 0 | 27 |
| [Retail UAT.dbo.REPOApprovalParameterProduct_TM](tables/retail.dbo.repoapprovalparameterproduct_tm.md) | 2 | 2 | 2 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.REPOLogSentMailKonfirmasi_TH](tables/retail.dbo.repologsentmailkonfirmasi_th.md) | 3 | 5 | 2 | 0 | 0 | 1 | 0 | 4 |
| [Retail UAT.dbo.REPOLogSentMailNotice_TH](tables/retail.dbo.repologsentmailnotice_th.md) | 4 | 6 | 2 | 0 | 0 | 1 | 0 | 5 |
| [Retail UAT.dbo.RMMSecurityTransaction_TH](tables/retail.dbo.rmmsecuritytransaction_th.md) | 1 | 3 | 2 | 0 | 0 | 0 | 0 | 3 |
| [Retail UAT.dbo.SBNMailNotifParam_TR](tables/retail.dbo.sbnmailnotifparam_tr.md) | 1 | 12 | 2 | 0 | 0 | 0 | 12 | 0 |
| [Retail UAT.dbo.SecurityFeeNasabahPending_TM](tables/retail.dbo.securityfeenasabahpending_tm.md) | 2 | 6 | 2 | 0 | 0 | 0 | 2 | 4 |
| [Retail UAT.dbo.SecurityTransaction_TEMP_ForMurex](tables/retail.dbo.securitytransaction_temp_formurex.md) | 6 | 7 | 2 | 0 | 0 | 0 | 0 | 7 |
| [Retail UAT.dbo.SecurityTransactionExtMig_TT](tables/retail.dbo.securitytransactionextmig_tt.md) | 1 | 2 | 2 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.SecurityTrxResultLelang_TH](tables/retail.dbo.securitytrxresultlelang_th.md) | 2 | 4 | 2 | 0 | 0 | 0 | 0 | 4 |
| [Retail UAT.dbo.tmp_ptos_transaction_upload](tables/retail.dbo.tmp_ptos_transaction_upload.md) | 5 | 23 | 2 | 0 | 0 | 2 | 8 | 13 |
| [Retail UAT.dbo.TreasuryCustomer_TH](tables/retail.dbo.treasurycustomer_th.md) | 4 | 4 | 2 | 0 | 0 | 1 | 0 | 3 |
| [Retail UAT.dbo.TreasuryCustomerBalance_TH](tables/retail.dbo.treasurycustomerbalance_th.md) | 2 | 5 | 2 | 0 | 0 | 1 | 0 | 4 |
| [Retail UAT.dbo.TRSCustodianUploadID_TR](tables/retail.dbo.trscustodianuploadid_tr.md) | 3 | 10 | 2 | 0 | 0 | 1 | 9 | 0 |
| [Retail UAT.dbo.TRSCustomerRepoFundedPending_TR](tables/retail.dbo.trscustomerrepofundedpending_tr.md) | 1 | 6 | 2 | 0 | 0 | 0 | 5 | 1 |
| [Retail UAT.dbo.TRSECBondsPrice_TH](tables/retail.dbo.trsecbondsprice_th.md) | 5 | 5 | 2 | 0 | 0 | 0 | 0 | 5 |
| [Retail UAT.dbo.TRSECGeneralParam_TR](tables/retail.dbo.trsecgeneralparam_tr.md) | 1 | 2 | 2 | 0 | 0 | 1 | 0 | 1 |
| [Retail UAT.dbo.TRSHasilRevenue_TH](tables/retail.dbo.trshasilrevenue_th.md) | 4 | 4 | 2 | 0 | 0 | 0 | 1 | 3 |
| [Retail UAT.dbo.TRSHasilRevenueBulkLog_TM](tables/retail.dbo.trshasilrevenuebulklog_tm.md) | 1 | 4 | 2 | 0 | 0 | 0 | 0 | 4 |
| [Retail UAT.dbo.TRSListNumber_TH](tables/retail.dbo.trslistnumber_th.md) | 1 | 3 | 2 | 0 | 0 | 1 | 0 | 2 |
| [Retail UAT.dbo.TRSMappingMurexBIS4Temp](tables/retail.dbo.trsmappingmurexbis4temp.md) | 2 | 13 | 2 | 0 | 0 | 0 | 5 | 8 |
| [Retail UAT.dbo.TRSNTIParameter_TR](tables/retail.dbo.trsntiparameter_tr.md) | 1 | 4 | 2 | 0 | 0 | 4 | 0 | 0 |
| [Retail UAT.dbo.TRSOnlineDocsLinkAppDetail_TH](tables/retail.dbo.trsonlinedocslinkappdetail_th.md) | 1 | 11 | 2 | 0 | 0 | 0 | 11 | 0 |
| [Retail UAT.dbo.TRSParamBookBld_TL](tables/retail.dbo.trsparambookbld_tl.md) | 1 | 7 | 2 | 0 | 0 | 0 | 7 | 0 |
| [Retail UAT.dbo.TRSParamDeviasiTotalSpread_TR](tables/retail.dbo.trsparamdeviasitotalspread_tr.md) | 5 | 11 | 2 | 0 | 0 | 4 | 0 | 7 |
| [Retail UAT.dbo.TRSParameterRevenueCabang_TH](tables/retail.dbo.trsparameterrevenuecabang_th.md) | 1 | 3 | 2 | 0 | 0 | 0 | 0 | 3 |
| [Retail UAT.dbo.TRSParameterRevenueInternal_TH](tables/retail.dbo.trsparameterrevenueinternal_th.md) | 1 | 3 | 2 | 0 | 0 | 0 | 0 | 3 |
| [Retail UAT.dbo.TRSParamIssuer_TR](tables/retail.dbo.trsparamissuer_tr.md) | 2 | 6 | 2 | 0 | 0 | 1 | 2 | 3 |
| [Retail UAT.dbo.TRSParamLelang_TH](tables/retail.dbo.trsparamlelang_th.md) | 1 | 3 | 2 | 0 | 0 | 0 | 1 | 2 |
| [Retail UAT.dbo.TRSParamLelangDetailApproval_TR](tables/retail.dbo.trsparamlelangdetailapproval_tr.md) | 2 | 11 | 2 | 0 | 0 | 0 | 10 | 1 |
| [Retail UAT.dbo.TRSParamLelangMasterApproval_TM](tables/retail.dbo.trsparamlelangmasterapproval_tm.md) | 1 | 5 | 2 | 0 | 0 | 0 | 4 | 1 |
| [Retail UAT.dbo.TRSParamMaterai_TR](tables/retail.dbo.trsparammaterai_tr.md) | 2 | 10 | 2 | 0 | 0 | 0 | 0 | 10 |
| [Retail UAT.dbo.TRSRetailAutoBatchErrLog](tables/retail.dbo.trsretailautobatcherrlog.md) | 1 | 2 | 2 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.TRSSafeKeepingFeeCustody_TH](tables/retail.dbo.trssafekeepingfeecustody_th.md) | 3 | 8 | 2 | 0 | 0 | 0 | 8 | 1 |
| [Retail UAT.dbo.TRSSettlementBIS4_TH](tables/retail.dbo.trssettlementbis4_th.md) | 2 | 4 | 2 | 0 | 0 | 0 | 0 | 4 |
| [Retail UAT.dbo.TRSSimanisTransaction_TH](tables/retail.dbo.trssimanistransaction_th.md) | 1 | 3 | 2 | 0 | 0 | 0 | 0 | 3 |
| [Retail UAT.dbo.TRSSPDebugDetail_TT](tables/retail.dbo.trsspdebugdetail_tt.md) | 1 | 2 | 2 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.TRSStaticCustomerOutstanding_TM](tables/retail.dbo.trsstaticcustomeroutstanding_tm.md) | 3 | 7 | 2 | 0 | 0 | 1 | 2 | 6 |
| [Retail UAT.dbo.TRSTransactionFeeLog_TR](tables/retail.dbo.trstransactionfeelog_tr.md) | 1 | 3 | 2 | 0 | 0 | 0 | 2 | 1 |
| [SQL_REPLICATE.dbo.BIS4_INSTRUMENT](tables/sql_replicate.dbo.bis4_instrument.md) | 3 | 22 | 2 | 0 | 0 | 12 | 0 | 10 |
| [SQL_REPLICATE.dbo.MXCUSTODY](tables/sql_replicate.dbo.mxcustody.md) | 3 | 7 | 2 | 0 | 0 | 0 | 0 | 7 |
| [SQL_REPLICATE.dbo.MXHolidayDate](tables/sql_replicate.dbo.mxholidaydate.md) | 1 | 2 | 2 | 0 | 0 | 0 | 0 | 2 |
| [SQL_REPLICATE.dbo.office_information](tables/sql_replicate.dbo.office_information.md) | 2 | 4 | 2 | 0 | 0 | 1 | 3 | 0 |
| [SQL_REPLICATE.dbo.PBCustomerFlagging_TM](tables/sql_replicate.dbo.pbcustomerflagging_tm.md) | 1 | 3 | 2 | 0 | 0 | 1 | 0 | 2 |
| [SQL_REPLICATE.dbo.ProCIFLastUpdateRiskProfile_TT](tables/sql_replicate.dbo.prociflastupdateriskprofile_tt.md) | 2 | 3 | 2 | 0 | 0 | 1 | 0 | 2 |
| [SQL_REPLICATE.dbo.TLTEL](tables/sql_replicate.dbo.tltel.md) | 2 | 4 | 2 | 0 | 0 | 0 | 0 | 4 |
| [SQL_REPLICATE.dbo.VLSRate_TM](tables/sql_replicate.dbo.vlsrate_tm.md) | 3 | 10 | 2 | 0 | 0 | 2 | 0 | 8 |
| [Retail UAT.dbo.AmendAccrued_TM](tables/retail.dbo.amendaccrued_tm.md) | 1 | 1 | 1 | 0 | 0 | 0 | 1 | 0 |
| [Retail UAT.dbo.BKHoldingData_TM](tables/retail.dbo.bkholdingdata_tm.md) | 1 | 2 | 1 | 0 | 0 | 0 | 2 | 0 |
| [Retail UAT.dbo.BONDLogSentMailConfirmation_TH](tables/retail.dbo.bondlogsentmailconfirmation_th.md) | 1 | 2 | 1 | 0 | 0 | 0 | 2 | 0 |
| [Retail UAT.dbo.BONDSDetailTransaksiBuyBack_TR](tables/retail.dbo.bondsdetailtransaksibuyback_tr.md) | 2 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.BONDTrxBeliLogSentMailConfirmation_TH](tables/retail.dbo.bondtrxbelilogsentmailconfirmation_th.md) | 1 | 2 | 1 | 0 | 0 | 0 | 2 | 0 |
| [Retail UAT.dbo.CIFReplicateFinal_TT](tables/retail.dbo.cifreplicatefinal_tt.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.control_office_table](tables/retail.dbo.control_office_table.md) | 1 | 1 | 1 | 0 | 0 | 0 | 1 | 0 |
| [Retail UAT.dbo.CorporateBondLogSentMailKonfirmasi_TH](tables/retail.dbo.corporatebondlogsentmailkonfirmasi_th.md) | 1 | 2 | 1 | 0 | 0 | 1 | 0 | 1 |
| [Retail UAT.dbo.EarlyRedemption_Temp](tables/retail.dbo.earlyredemption_temp.md) | 3 | 4 | 1 | 0 | 0 | 0 | 0 | 4 |
| [Retail UAT.dbo.EarlyRedemption_TH](tables/retail.dbo.earlyredemption_th.md) | 1 | 10 | 1 | 0 | 0 | 0 | 0 | 10 |
| [Retail UAT.dbo.EarlyRedemptionOnline_TH](tables/retail.dbo.earlyredemptiononline_th.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.GLONTHOblRevenueLog_TH](tables/retail.dbo.glonthoblrevenuelog_th.md) | 1 | 8 | 1 | 0 | 0 | 0 | 0 | 8 |
| [Retail UAT.dbo.gwp_general_param_tr](tables/retail.dbo.gwp_general_param_tr.md) | 2 | 2 | 1 | 0 | 0 | 1 | 0 | 1 |
| [Retail UAT.dbo.job_general_control_table_tm](tables/retail.dbo.job_general_control_table_tm.md) | 3 | 8 | 1 | 0 | 0 | 0 | 0 | 8 |
| [Retail UAT.dbo.MappingProvCity_TM](tables/retail.dbo.mappingprovcity_tm.md) | 2 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.MxRepoTran](tables/retail.dbo.mxrepotran.md) | 5 | 7 | 1 | 0 | 0 | 0 | 2 | 5 |
| [Retail UAT.dbo.ORA_TRN_HDR_DBF](tables/retail.dbo.ora_trn_hdr_dbf.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.ORI_Payment_TR](tables/retail.dbo.ori_payment_tr.md) | 2 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.ORISBNAlertLog_TH](tables/retail.dbo.orisbnalertlog_th.md) | 1 | 5 | 1 | 0 | 0 | 0 | 2 | 3 |
| [Retail UAT.dbo.oriva_parameter_gnc](tables/retail.dbo.oriva_parameter_gnc.md) | 1 | 1 | 1 | 0 | 0 | 1 | 0 | 0 |
| [Retail UAT.dbo.ProACCDDMUL01H_TH](tables/retail.dbo.proaccddmul01h_th.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.process_log](tables/retail.dbo.process_log.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.process_usage_locking_temp](tables/retail.dbo.process_usage_locking_temp.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.process_usage_locking_th](tables/retail.dbo.process_usage_locking_th.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.PTOSCustomer_Error](tables/retail.dbo.ptoscustomer_error.md) | 15 | 67 | 1 | 0 | 0 | 0 | 0 | 67 |
| [Retail UAT.dbo.PTOSCustomerInfo_TM](tables/retail.dbo.ptoscustomerinfo_tm.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.PTOSMappingProductCode_TH](tables/retail.dbo.ptosmappingproductcode_th.md) | 1 | 4 | 1 | 0 | 0 | 0 | 0 | 4 |
| [Retail UAT.dbo.PTOSOutstandingMaster_TH](tables/retail.dbo.ptosoutstandingmaster_th.md) | 2 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.PTOSParamEmail_TH](tables/retail.dbo.ptosparamemail_th.md) | 1 | 4 | 1 | 0 | 0 | 0 | 0 | 4 |
| [Retail UAT.dbo.PTOSSecurityMaster_TH](tables/retail.dbo.ptossecuritymaster_th.md) | 1 | 5 | 1 | 0 | 0 | 0 | 0 | 5 |
| [Retail UAT.dbo.PTOSSecurityMaster_TM](tables/retail.dbo.ptossecuritymaster_tm.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.REPOApprovalSpreadDetail_TR](tables/retail.dbo.repoapprovalspreaddetail_tr.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.REPOLogSentMailDefault_TH](tables/retail.dbo.repologsentmaildefault_th.md) | 2 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.REPOMasterTransaction_TH](tables/retail.dbo.repomastertransaction_th.md) | 1 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.REPOParamHargaModal_TH](tables/retail.dbo.repoparamhargamodal_th.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.RepoSpreadTransaksi_TM](tables/retail.dbo.repospreadtransaksi_tm.md) | 8 | 20 | 1 | 0 | 0 | 2 | 4 | 14 |
| [Retail UAT.dbo.REPOSysParameter_TR](tables/retail.dbo.reposysparameter_tr.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.SecurityBatchResult_TH](tables/retail.dbo.securitybatchresult_th.md) | 1 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.SecurityCouponDatePending_TR](tables/retail.dbo.securitycoupondatepending_tr.md) | 1 | 2 | 1 | 0 | 0 | 2 | 0 | 0 |
| [Retail UAT.dbo.SecurityResultForTA_TT](tables/retail.dbo.securityresultforta_tt.md) | 1 | 2 | 1 | 0 | 0 | 0 | 2 | 0 |
| [Retail UAT.dbo.SecurityResultLelang_TM](tables/retail.dbo.securityresultlelang_tm.md) | 1 | 5 | 1 | 0 | 0 | 0 | 3 | 2 |
| [Retail UAT.dbo.SecurityTransaction_ListBlokir](tables/retail.dbo.securitytransaction_listblokir.md) | 1 | 2 | 1 | 0 | 0 | 0 | 1 | 1 |
| [Retail UAT.dbo.SecurityTransactionCoupon_TH](tables/retail.dbo.securitytransactioncoupon_th.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.SecurityTransactionDetail_TR](tables/retail.dbo.securitytransactiondetail_tr.md) | 1 | 1 | 1 | 0 | 0 | 0 | 1 | 0 |
| [Retail UAT.dbo.SecurityTransactionSpread_TR](tables/retail.dbo.securitytransactionspread_tr.md) | 2 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.SecurityTransactionStatus_TM](tables/retail.dbo.securitytransactionstatus_tm.md) | 1 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.SecurityTranscationLink_TL](tables/retail.dbo.securitytranscationlink_tl.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.SRBILogSentMailConfirmation_TH](tables/retail.dbo.srbilogsentmailconfirmation_th.md) | 1 | 2 | 1 | 0 | 0 | 0 | 2 | 0 |
| [Retail UAT.dbo.SRBITrxBeliLogSentMailConfirmation_TH](tables/retail.dbo.srbitrxbelilogsentmailconfirmation_th.md) | 1 | 2 | 1 | 0 | 0 | 0 | 2 | 0 |
| [Retail UAT.dbo.SREPendingNotification_TR](tables/retail.dbo.srependingnotification_tr.md) | 1 | 1 | 1 | 0 | 0 | 0 | 1 | 0 |
| [Retail UAT.dbo.TempKursTgl18](tables/retail.dbo.tempkurstgl18.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.TibcoTellerResource_TR](tables/retail.dbo.tibcotellerresource_tr.md) | 1 | 3 | 1 | 0 | 0 | 0 | 1 | 2 |
| [Retail UAT.dbo.trs_getaccount_session](tables/retail.dbo.trs_getaccount_session.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.trs_list_source_reference_lelang](tables/retail.dbo.trs_list_source_reference_lelang.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.TRSAlamatSuratObligasiLog_TL](tables/retail.dbo.trsalamatsuratobligasilog_tl.md) | 1 | 1 | 1 | 0 | 0 | 1 | 0 | 0 |
| [Retail UAT.dbo.TRSAllowViewAllBranch](tables/retail.dbo.trsallowviewallbranch.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.TRSBungaFTPParameter_TH](tables/retail.dbo.trsbungaftpparameter_th.md) | 2 | 5 | 1 | 0 | 0 | 0 | 0 | 5 |
| [Retail UAT.dbo.TRSCalculateFeeLog_TH](tables/retail.dbo.trscalculatefeelog_th.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.TRSCallNotifAPILog_TL](tables/retail.dbo.trscallnotifapilog_tl.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.TRSCustAccNumberPending_TR](tables/retail.dbo.trscustaccnumberpending_tr.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.TRSCustodyMaster_TH](tables/retail.dbo.trscustodymaster_th.md) | 1 | 2 | 1 | 0 | 0 | 0 | 2 | 0 |
| [Retail UAT.dbo.TRSCustomerRepoFunded_TH](tables/retail.dbo.trscustomerrepofunded_th.md) | 2 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.TRSDataUploadMaturity_TT](tables/retail.dbo.trsdatauploadmaturity_tt.md) | 2 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.TRSDDMAST_TM](tables/retail.dbo.trsddmast_tm.md) | 1 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.TRSFileFormat_TM](tables/retail.dbo.trsfileformat_tm.md) | 2 | 10 | 1 | 0 | 0 | 1 | 1 | 8 |
| [Retail UAT.dbo.TRSGNCParameter_TH](tables/retail.dbo.trsgncparameter_th.md) | 1 | 1 | 1 | 0 | 0 | 0 | 1 | 0 |
| [Retail UAT.dbo.TRSInputSID_TL](tables/retail.dbo.trsinputsid_tl.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.TRSJobPushNotifOM_TT](tables/retail.dbo.trsjobpushnotifom_tt.md) | 2 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.TRSJobSendEmailRecipientParam_TR](tables/retail.dbo.trsjobsendemailrecipientparam_tr.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.TRSKuotaReserved_TMP](tables/retail.dbo.trskuotareserved_tmp.md) | 1 | 2 | 1 | 0 | 0 | 1 | 0 | 1 |
| [Retail UAT.dbo.TRSKuponCustomerSaldo_TT](tables/retail.dbo.trskuponcustomersaldo_tt.md) | 2 | 4 | 1 | 0 | 0 | 0 | 1 | 3 |
| [Retail UAT.dbo.TRSLastCounterParameter_TR](tables/retail.dbo.trslastcounterparameter_tr.md) | 1 | 4 | 1 | 0 | 0 | 2 | 0 | 2 |
| [Retail UAT.dbo.TRSMapPekerjaan_TR](tables/retail.dbo.trsmappekerjaan_tr.md) | 1 | 1 | 1 | 0 | 0 | 0 | 1 | 0 |
| [Retail UAT.dbo.TRSMappingMatchRiskProfile_TR](tables/retail.dbo.trsmappingmatchriskprofile_tr.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.TRSMappingMurexBIS4_TH](tables/retail.dbo.trsmappingmurexbis4_th.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.TRSOnCallDepositListRek_TH](tables/retail.dbo.trsoncalldepositlistrek_th.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.TRSOnCallDepositListRek_TT](tables/retail.dbo.trsoncalldepositlistrek_tt.md) | 5 | 10 | 1 | 0 | 0 | 0 | 0 | 10 |
| [Retail UAT.dbo.TRSORIMDB](tables/retail.dbo.trsorimdb.md) | 3 | 5 | 1 | 0 | 0 | 0 | 0 | 5 |
| [Retail UAT.dbo.TRSParamBookBld_TH](tables/retail.dbo.trsparambookbld_th.md) | 2 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.TRSParamClassificationDivision](tables/retail.dbo.trsparamclassificationdivision.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.TRSParamDeviasiHarga_TR](tables/retail.dbo.trsparamdeviasiharga_tr.md) | 1 | 2 | 1 | 0 | 0 | 0 | 2 | 0 |
| [Retail UAT.dbo.TRSParameterJob_TM](tables/retail.dbo.trsparameterjob_tm.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.TRSParameterRevenueCabangLog_TM](tables/retail.dbo.trsparameterrevenuecabanglog_tm.md) | 1 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.TRSParameterRevenueInternalLog_TM](tables/retail.dbo.trsparameterrevenueinternallog_tm.md) | 1 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.TRSParameterRiskProfile_TM](tables/retail.dbo.trsparameterriskprofile_tm.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.TRSParameterSettlementDate_TL](tables/retail.dbo.trsparametersettlementdate_tl.md) | 2 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.TRSParameterSettlementDate_TR](tables/retail.dbo.trsparametersettlementdate_tr.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.TRSParamHargaModalSRBI_TL](tables/retail.dbo.trsparamhargamodalsrbi_tl.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.TRSParamInterfaceMurex_TH](tables/retail.dbo.trsparaminterfacemurex_th.md) | 1 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.TRSParamIssuer_TH](tables/retail.dbo.trsparamissuer_th.md) | 1 | 4 | 1 | 0 | 0 | 0 | 4 | 0 |
| [Retail UAT.dbo.TRSParamOfftoleranceHargaModal_TH](tables/retail.dbo.trsparamofftolerancehargamodal_th.md) | 1 | 4 | 1 | 0 | 0 | 0 | 0 | 4 |
| [Retail UAT.dbo.TRSParamOfftoleranceHargaModal_TR](tables/retail.dbo.trsparamofftolerancehargamodal_tr.md) | 4 | 10 | 1 | 0 | 0 | 0 | 0 | 10 |
| [Retail UAT.dbo.TRSParamSpreadHargaModal_TH](tables/retail.dbo.trsparamspreadhargamodal_th.md) | 1 | 3 | 1 | 0 | 0 | 0 | 0 | 3 |
| [Retail UAT.dbo.TRSParamThreshold_TH](tables/retail.dbo.trsparamthreshold_th.md) | 2 | 5 | 1 | 0 | 0 | 0 | 0 | 5 |
| [Retail UAT.dbo.TRSParamThreshold_TM](tables/retail.dbo.trsparamthreshold_tm.md) | 1 | 4 | 1 | 0 | 0 | 0 | 2 | 2 |
| [Retail UAT.dbo.TRSParamTradeSettleDate_TR](tables/retail.dbo.trsparamtradesettledate_tr.md) | 2 | 4 | 1 | 0 | 0 | 0 | 2 | 2 |
| [Retail UAT.dbo.TRSPekerjaan_TM](tables/retail.dbo.trspekerjaan_tm.md) | 1 | 1 | 1 | 0 | 0 | 0 | 1 | 0 |
| [Retail UAT.dbo.TRSProcessDetail_TR](tables/retail.dbo.trsprocessdetail_tr.md) | 1 | 1 | 1 | 0 | 0 | 0 | 1 | 0 |
| [Retail UAT.dbo.TRSProcessTransaksi_TH](tables/retail.dbo.trsprocesstransaksi_th.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.TRSProcessTransaksiDetail_TH](tables/retail.dbo.trsprocesstransaksidetail_th.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.TRSRekSyariahLog_TH](tables/retail.dbo.trsreksyariahlog_th.md) | 2 | 4 | 1 | 0 | 0 | 0 | 0 | 4 |
| [Retail UAT.dbo.TRSRepoNotif_TL](tables/retail.dbo.trsreponotif_tl.md) | 1 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.TRSReportEmailStatus_TL](tables/retail.dbo.trsreportemailstatus_tl.md) | 4 | 4 | 1 | 0 | 0 | 2 | 2 | 0 |
| [Retail UAT.dbo.TRSReportEmailStatusMap_TR](tables/retail.dbo.trsreportemailstatusmap_tr.md) | 3 | 6 | 1 | 0 | 0 | 2 | 0 | 4 |
| [Retail UAT.dbo.TRSRetailCustBlock_SequenceNumber](tables/retail.dbo.trsretailcustblock_sequencenumber.md) | 1 | 4 | 1 | 0 | 0 | 0 | 0 | 4 |
| [Retail UAT.dbo.TRSRetailCustBlock_TH](tables/retail.dbo.trsretailcustblock_th.md) | 1 | 1 | 1 | 0 | 0 | 0 | 1 | 0 |
| [Retail UAT.dbo.TRSRetailGlobalLog_TL](tables/retail.dbo.trsretailgloballog_tl.md) | 1 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [Retail UAT.dbo.TRSRetailUserClassificationHistory](tables/retail.dbo.trsretailuserclassificationhistory.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.TRSSecurityTransactionPrimaryPreBook_TH](tables/retail.dbo.trssecuritytransactionprimaryprebook_th.md) | 1 | 1 | 1 | 0 | 0 | 1 | 0 | 0 |
| [Retail UAT.dbo.TRSSettleParamRevenue_TM](tables/retail.dbo.trssettleparamrevenue_tm.md) | 1 | 1 | 1 | 0 | 0 | 0 | 1 | 0 |
| [Retail UAT.dbo.TRSStatusCorporateBond_TH](tables/retail.dbo.trsstatuscorporatebond_th.md) | 1 | 3 | 1 | 0 | 0 | 0 | 0 | 3 |
| [Retail UAT.dbo.TrsSyncMasterIdAndFuncTemp](tables/retail.dbo.trssyncmasteridandfunctemp.md) | 3 | 7 | 1 | 0 | 0 | 0 | 1 | 6 |
| [Retail UAT.dbo.TRSTenorParameter_TH](tables/retail.dbo.trstenorparameter_th.md) | 2 | 5 | 1 | 0 | 0 | 0 | 0 | 5 |
| [Retail UAT.dbo.TRSValidationSecLog_TH](tables/retail.dbo.trsvalidationseclog_th.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [Retail UAT.dbo.upload_sid_sre_temp](tables/retail.dbo.upload_sid_sre_temp.md) | 3 | 7 | 1 | 0 | 0 | 0 | 6 | 1 |
| [SQL_REPLICATE.dbo.AccountSegment](tables/sql_replicate.dbo.accountsegment.md) | 1 | 1 | 1 | 0 | 0 | 0 | 1 | 0 |
| [SQL_REPLICATE.dbo.BIS4_SecurityTrx](tables/sql_replicate.dbo.bis4_securitytrx.md) | 2 | 6 | 1 | 0 | 0 | 1 | 0 | 5 |
| [SQL_REPLICATE.dbo.CFBICD](tables/sql_replicate.dbo.cfbicd.md) | 1 | 1 | 1 | 0 | 0 | 0 | 1 | 0 |
| [SQL_REPLICATE.dbo.CFFINCE](tables/sql_replicate.dbo.cffince.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [SQL_REPLICATE.dbo.CFPAR2](tables/sql_replicate.dbo.cfpar2.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [SQL_REPLICATE.dbo.CFYMSC](tables/sql_replicate.dbo.cfymsc.md) | 1 | 1 | 1 | 0 | 0 | 0 | 1 | 0 |
| [SQL_REPLICATE.dbo.CFYPIN](tables/sql_replicate.dbo.cfypin.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [SQL_REPLICATE.dbo.counterp_dbf](tables/sql_replicate.dbo.counterp_dbf.md) | 1 | 1 | 1 | 0 | 0 | 0 | 1 | 0 |
| [SQL_REPLICATE.dbo.DDHIST](tables/sql_replicate.dbo.ddhist.md) | 2 | 12 | 1 | 0 | 0 | 0 | 0 | 12 |
| [SQL_REPLICATE.dbo.DDPAR2EXT](tables/sql_replicate.dbo.ddpar2ext.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [SQL_REPLICATE.dbo.DDPAR2MC](tables/sql_replicate.dbo.ddpar2mc.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [SQL_REPLICATE.dbo.DDSTOP](tables/sql_replicate.dbo.ddstop.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [SQL_REPLICATE.dbo.DDTRN2](tables/sql_replicate.dbo.ddtrn2.md) | 3 | 12 | 1 | 0 | 0 | 8 | 0 | 4 |
| [SQL_REPLICATE.dbo.GLACCT](tables/sql_replicate.dbo.glacct.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [SQL_REPLICATE.dbo.GLMAST](tables/sql_replicate.dbo.glmast.md) | 1 | 1 | 1 | 0 | 0 | 0 | 0 | 1 |
| [SQL_REPLICATE.dbo.MISCurrencyRateHistory](tables/sql_replicate.dbo.miscurrencyratehistory.md) | 1 | 2 | 1 | 0 | 0 | 1 | 0 | 1 |
| [SQL_REPLICATE.dbo.Mx_BondTran_TMP](tables/sql_replicate.dbo.mx_bondtran_tmp.md) | 3 | 3 | 1 | 0 | 0 | 2 | 0 | 1 |
| [SQL_REPLICATE.dbo.Mx_RepoTran_TMP](tables/sql_replicate.dbo.mx_repotran_tmp.md) | 5 | 9 | 1 | 0 | 0 | 5 | 0 | 4 |
| [SQL_REPLICATE.dbo.MX_SECURITY](tables/sql_replicate.dbo.mx_security.md) | 2 | 25 | 1 | 0 | 0 | 0 | 0 | 25 |
| [SQL_REPLICATE.dbo.MXOUTSTANDING_REPO_DBF](tables/sql_replicate.dbo.mxoutstanding_repo_dbf.md) | 1 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [SQL_REPLICATE.dbo.MXPergerakanDana](tables/sql_replicate.dbo.mxpergerakandana.md) | 1 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [SQL_REPLICATE.dbo.MXPortofoliDBF](tables/sql_replicate.dbo.mxportofolidbf.md) | 2 | 9 | 1 | 0 | 0 | 0 | 0 | 9 |
| [SQL_REPLICATE.dbo.MXPurpose](tables/sql_replicate.dbo.mxpurpose.md) | 1 | 2 | 1 | 0 | 0 | 0 | 0 | 2 |
| [SQL_REPLICATE.dbo.MXReasonReverse](tables/sql_replicate.dbo.mxreasonreverse.md) | 2 | 5 | 1 | 0 | 0 | 0 | 0 | 5 |
| [SQL_REPLICATE.dbo.ORA_TRN_HDR_DBF_REP](tables/sql_replicate.dbo.ora_trn_hdr_dbf_rep.md) | 1 | 1 | 1 | 0 | 0 | 1 | 0 | 0 |
| [SQL_REPLICATE.dbo.ProCIFChangeMaster_TH](tables/sql_replicate.dbo.procifchangemaster_th.md) | 1 | 3 | 1 | 0 | 0 | 0 | 0 | 3 |
| [SQL_REPLICATE.dbo.RepoMasterTransaction_Tmp](tables/sql_replicate.dbo.repomastertransaction_tmp.md) | 14 | 15 | 1 | 0 | 0 | 12 | 0 | 3 |

## Tabel tanpa bound match

Daftar lengkap berada di `table_summary.csv`, status NO_BOUND_MATCH_NOT_CERTIFIED. Ini mencakup tabel tanpa pemakaian terkait, tipe sama, maupun jalur yang belum berhasil di-bind; bukan bukti semua query terhadap tabel tersebut bebas conversion.
