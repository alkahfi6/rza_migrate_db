# Kandidat index gabungan seluruh workload

Tanggal: 22 Oktober 2026. Status seluruh kandidat: **PROPOSED, belum dieksekusi/diukur**. A = tinjau awal, B = setelah rewrite/ukur kebutuhan, C = kondisional. Angka procedure adalah cakupan source, bukan frekuensi eksekusi.

Existing prefix di bawah mengabaikan INCLUDE saat menentukan search key. Daftar prefix bukan bukti coverage penuh/aktif di live DB. Predicate dalam OR tidak dianggap otomatis equality wajib.

| ID | Prioritas | Tabel dan key | Procedure dengan equality key pertama | Existing prefix terkait |
|---|---|---|---:|---|
| IX01 | A | `dbo.SecurityMaster_TM (SecurityNo)` | 148 | Tidak ada |
| IX02 | B | `dbo.SecurityMaster_TM (IdSeri)` | 9 | Tidak ada |
| IX03 | A | `dbo.SecurityTransaction_TT (DealNo)` | 22 | Tidak ada |
| IX04 | B | `dbo.SecurityTransaction_TT (CIFId, SecId, TrxType, TrxStatus, SettlementDate)` | 118 | SecurityTransaction_TT_NN3, SecurityTransaction_TT_NN8 |
| IX05 | B | `dbo.ORI_Order_TT (SecId, CIFId)` | 52 | Tidak ada |
| IX06 | A | `dbo.SecurityTransactionDetail_TT (DealId)` | 12 | Tidak ada |
| IX07 | B | `dbo.TRSECParamSpread_TR (SecId, ChannelCode, BuySell, NominalFrom)` | 17 | Tidak ada |
| IX08 | B | `dbo.TRSParameterSettlementDate_TR (SecId, ChannelCode)` | 16 | Tidak ada |
| IX09 | C | `dbo.TRSParamHargaModal_TR (SecurityNo)` | 15 | Tidak ada |
| IX10 | B | `dbo.PTOSCustomer_TM (NoRekInvestorRDN)` | 18 | Tidak ada |
| IX11 | B | `dbo.PTOSOutstanding_TR (SID, NoRekInvestorRDN, KodeProduk, PTOSType, Tanggal)` | 13 | ix_ptosoutstanding_tr_sid, ix_ptosoutstanding_tr_sidnordn |
| IX12 | A | `dbo.PTOSTransaksi_TR (DealId, PTOSType)` | 10 | Tidak ada |
| IX13 | C | `dbo.SecurityCouponDate_TR (SecId, Done, Periode)` | 11 | SecurityCouponDate_TR_CN1 |
| IX14 | B | `dbo.TRSCustAccNumber_TR (CIFId, RowId DESC)` | 68 | TRSCustAccNumber_TR_NN1 |
| IX15 | A | `dbo.TRSRetailCustBlock_TM (DealId)` | 13 | Tidak ada |
| IX16 | C | `dbo.TRSRetailCustBlock_TM (CIFId, CollateralStatus, SecId)` | 10 | Tidak ada |
| IX17 | B | `dbo.CFMAST (CFSSNO, CFSSCD, CFBIRD)` | 4 | idx_cfmast_cfssno |
| IX18 | C_AFTER_REWRITE | `dbo.ProCIFChangeMaster_TH (ChangeDate, ChangeId)` | 0 | Tidak ada |

## IX01 — dbo.SecurityMaster_TM

Lookup SecurityNo lintas workload. IU_SecurityMaster_TM berawalan SecId sehingga bukan seek key untuk SecurityNo saja.

**Prasyarat:** Pastikan input SecurityNo tetap VARCHAR, bukan ID numerik; ukur ukuran tabel dan lookup yang tersisa.

**Bukti workload:** 268 predicate kesetaraan key pertama pada 148 procedure; detail seluruh lokasi pada `index_workload.csv`.

**Contoh konsumen:** `procedures/dbo.CustodyReportSettlementInstruction.sql:178`; `procedures/dbo.CustodyReportSettlementInstruction.sql:246`; `procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql:221`; `procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql:276`; `procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql:319`

**Existing index pada tabel:** IU_SecurityMaster_TM (SecId|SecurityNo|InterestTypeID|SecCcy) INCLUDE (); PK__SecurityMaster_T__7E37BEF6 (SecId) INCLUDE ()

```sql
CREATE NONCLUSTERED INDEX [IX_Audit20261022_SecurityMaster_TM_SecurityNo]
ON [dbo].[SecurityMaster_TM] ([SecurityNo] ASC)
INCLUDE ([SecCcy], [SecType], [ISINCode], [InterestTypeID]);
```

## IX02 — dbo.SecurityMaster_TM

Lookup IdSeri pada flow ORI/SBN. Belum ada index berawalan IdSeri.

**Prasyarat:** Perbaiki pembandingan IdSeri VARCHAR(100) dengan INT dahulu; definisikan aturan leading zero dan nilai nonnumerik.

**Bukti workload:** 23 predicate kesetaraan key pertama pada 9 procedure; detail seluruh lokasi pada `index_workload.csv`.

**Contoh konsumen:** `procedures/dbo.ORISBNIPopulateEarlyRedemption.sql:33`; `procedures/dbo.ORISBNIPopulateEarlyRedemption.sql:61`; `procedures/dbo.ORISBNIPopulateEarlyRedemption.sql:117`; `procedures/dbo.ORISBNIPopulateEarlyRedemption.sql:277`; `procedures/dbo.ORISBNIPopulateRegisterInvestor.sql:73`

**Existing index pada tabel:** IU_SecurityMaster_TM (SecId|SecurityNo|InterestTypeID|SecCcy) INCLUDE (); PK__SecurityMaster_T__7E37BEF6 (SecId) INCLUDE ()

```sql
CREATE NONCLUSTERED INDEX [IX_Audit20261022_SecurityMaster_TM_IdSeri]
ON [dbo].[SecurityMaster_TM] ([IdSeri] ASC)
INCLUDE ([SecurityNo], [SecCcy]);
```

## IX03 — dbo.SecurityTransaction_TT

Lookup DealNo; IU_SecurityTransaction_TT berawalan DealId. PK clustered DealId sudah menjadi row locator implisit.

**Prasyarat:** Periksa panjang parameter dan query aktual; jangan jadikan UNIQUE tanpa bukti data.

**Bukti workload:** 38 predicate kesetaraan key pertama pada 22 procedure; detail seluruh lokasi pada `index_workload.csv`.

**Contoh konsumen:** `procedures/dbo.DFSListObliTransaction.sql:37`; `procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql:365`; `procedures/dbo.GenBatchTxtFile2CMS_Urgent.sql:452`; `procedures/dbo.GenBatchTxtFile2CMS.sql:365`; `procedures/dbo.GenBatchTxtFile2CMS.sql:452`

**Existing index pada tabel:** IU_SecurityTransaction_TT (DealId|DealNo|CIFId|SecId|TrxType|TrxDate) INCLUDE (); SecurityTransaction_TT_NN1 (TrxType|TrxStatus) INCLUDE (SecId|FaceValue|InsertedDate|DeletedDate|IsLelang); SecurityTransaction_TT_NN10 (SecId) INCLUDE (); SecurityTransaction_TT_NN11 (InsertedDate) INCLUDE (); SecurityTransaction_TT_NN12 (DeletedDate) INCLUDE (); SecurityTransaction_TT_NN13 (TrxGuid) INCLUDE (); SecurityTransaction_TT_NN14 (TrxType|SettlementDate|TrxStatus) INCLUDE (NoRekInvestor|Instruction|ORISettlmntType); SecurityTransaction_TT_NN15 (TrxType|TrxStatus) INCLUDE (CIFId|SecId|FaceValue); SecurityTransaction_TT_NN2 (TrxType|OriginalAmount) INCLUDE (CIFId|SecId); SecurityTransaction_TT_NN3 (CIFId|SecId) INCLUDE (TrxType); SecurityTransaction_TT_NN4 (TrxDate|TrxStatus) INCLUDE (); SecurityTransaction_TT_NN5 (OrderId) INCLUDE (); SecurityTransaction_TT_NN6 (TrxDate) INCLUDE (); SecurityTransaction_TT_NN7 (IdLelang) INCLUDE (); SecurityTransaction_TT_NN8 (CIFId) INCLUDE (); SecurityTransaction_TT_NN9 (DealIdExternal) INCLUDE (); idx_SecurityTransaction_TT (DealIdExternal|SettlementDate|TrxType|TrxStatus) INCLUDE (); PK__SecurityTransact__04E4BC85 (DealId) INCLUDE ()

```sql
CREATE NONCLUSTERED INDEX [IX_Audit20261022_SecurityTransaction_TT_DealNo]
ON [dbo].[SecurityTransaction_TT] ([DealNo] ASC)
INCLUDE ([CIFId], [SecId], [TrxType], [TrxStatus], [SettlementDate]);
```

## IX04 — dbo.SecurityTransaction_TT

Pola saldo/coupon: equality CIFId,SecId,TrxType,TrxStatus lalu rentang SettlementDate. NN3 hanya CIFId,SecId INCLUDE TrxType.

**Prasyarat:** Bandingkan dengan NN3/NN8 dan index TrxType,TrxStatus. Tidak semua query mengisi seluruh equality prefix; jangan otomatis drop index lama.

**Bukti workload:** 249 predicate kesetaraan key pertama pada 118 procedure; detail seluruh lokasi pada `index_workload.csv`.

**Contoh konsumen:** `procedures/dbo.CustodyReportMaintainAcquisition.sql:96`; `procedures/dbo.CustodyReportSettlementInstruction.sql:170`; `procedures/dbo.DFSListObliTransaction.sql:19`; `procedures/dbo.DFSListObliTransaction.sql:31`; `procedures/dbo.DFSListSecNumber.sql:20`

**Existing index pada tabel:** IU_SecurityTransaction_TT (DealId|DealNo|CIFId|SecId|TrxType|TrxDate) INCLUDE (); SecurityTransaction_TT_NN1 (TrxType|TrxStatus) INCLUDE (SecId|FaceValue|InsertedDate|DeletedDate|IsLelang); SecurityTransaction_TT_NN10 (SecId) INCLUDE (); SecurityTransaction_TT_NN11 (InsertedDate) INCLUDE (); SecurityTransaction_TT_NN12 (DeletedDate) INCLUDE (); SecurityTransaction_TT_NN13 (TrxGuid) INCLUDE (); SecurityTransaction_TT_NN14 (TrxType|SettlementDate|TrxStatus) INCLUDE (NoRekInvestor|Instruction|ORISettlmntType); SecurityTransaction_TT_NN15 (TrxType|TrxStatus) INCLUDE (CIFId|SecId|FaceValue); SecurityTransaction_TT_NN2 (TrxType|OriginalAmount) INCLUDE (CIFId|SecId); SecurityTransaction_TT_NN3 (CIFId|SecId) INCLUDE (TrxType); SecurityTransaction_TT_NN4 (TrxDate|TrxStatus) INCLUDE (); SecurityTransaction_TT_NN5 (OrderId) INCLUDE (); SecurityTransaction_TT_NN6 (TrxDate) INCLUDE (); SecurityTransaction_TT_NN7 (IdLelang) INCLUDE (); SecurityTransaction_TT_NN8 (CIFId) INCLUDE (); SecurityTransaction_TT_NN9 (DealIdExternal) INCLUDE (); idx_SecurityTransaction_TT (DealIdExternal|SettlementDate|TrxType|TrxStatus) INCLUDE (); PK__SecurityTransact__04E4BC85 (DealId) INCLUDE ()

```sql
CREATE NONCLUSTERED INDEX [IX_Audit20261022_SecurityTransaction_TT_CIFId_SecId_TrxType_TrxStatus_SettlementDate]
ON [dbo].[SecurityTransaction_TT] ([CIFId] ASC, [SecId] ASC, [TrxType] ASC, [TrxStatus] ASC, [SettlementDate] ASC)
INCLUDE ([FaceValue], [InsertedBy]);
```

## IX05 — dbo.ORI_Order_TT

Lookup order per surat berharga/nasabah; SecId hanya INCLUDE pada IU_ORI_Order_TT berawalan OrderId.

**Prasyarat:** Pilih urutan SecId,CIFId dari selektivitas workload; query dengan OrderId tunggal sudah terlayani PK.

**Bukti workload:** 102 predicate kesetaraan key pertama pada 52 procedure; detail seluruh lokasi pada `index_workload.csv`.

**Contoh konsumen:** `procedures/dbo.CustodyReportSettlementInstruction.sql:237`; `procedures/dbo.ORISBNIPopulateEarlyRedemption.sql:159`; `procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql:277`; `procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql:304`; `procedures/dbo.SBNSchedUpdateStatusTimeOut.sql:55`

**Existing index pada tabel:** IU_ORI_Order_TT (OrderId) INCLUDE (CIFId|SecId|OrderDate|OrderBranch|NominalBuyBack|isTaxAmnesty); ORI_Order_TT_NN1 (CIFId) INCLUDE (); ORI_Order_TT_NN2 (InsertedDate) INCLUDE (); PK__ORI_Order_TT__01142BA1 (OrderId) INCLUDE ()

```sql
CREATE NONCLUSTERED INDEX [IX_Audit20261022_ORI_Order_TT_SecId_CIFId]
ON [dbo].[ORI_Order_TT] ([SecId] ASC, [CIFId] ASC)
INCLUDE ([OrderDate], [OrderStatus], [IsBuyBack]);
```

## IX06 — dbo.SecurityTransactionDetail_TT

JOIN detail pada DealId. IU_SecurityTransactionDetail_TT berawalan RowId, bukan DealId.

**Prasyarat:** Projection lebar mungkin tetap memerlukan RID lookup; tambah INCLUDE hanya berdasarkan actual plan.

**Bukti workload:** 36 predicate kesetaraan key pertama pada 12 procedure; detail seluruh lokasi pada `index_workload.csv`.

**Contoh konsumen:** `procedures/dbo.Rvmp_trs_ListSecurityTransactionForInsertEditBlockTrans.sql:54`; `procedures/dbo.trs_InsertTrx2ndSecurityTransaction_TT.sql:280`; `procedures/dbo.trs_InsertTrx2ndSecurityTransaction_TT.sql:309`; `procedures/dbo.trs_InsertTrx2ndSecurityTransaction_TT.sql:314`; `procedures/dbo.trs_InsertTrx2ndSecurityTransaction_TT.sql:356`

**Existing index pada tabel:** IU_SecurityTransactionDetail_TT (RowId|DealId|CIFIdDebet|CIFIdKredit|TrxDetailType|D/K|CustodyName) INCLUDE ()

```sql
CREATE NONCLUSTERED INDEX [IX_Audit20261022_SecurityTransactionDetail_TT_DealId]
ON [dbo].[SecurityTransactionDetail_TT] ([DealId] ASC)
INCLUDE ([TrxDetailType], [CIFIdDebet], [CIFIdKredit]);
```

## IX07 — dbo.TRSECParamSpread_TR

Pola harga produk: SecId+ChannelCode+BuySell lalu nominal tier. Index saat ini hanya ParamId.

**Prasyarat:** Parameter table mungkin kecil; jangan deploy bila scan murah. Predicate NominalTo masih residual untuk interval overlap.

**Bukti workload:** 44 predicate kesetaraan key pertama pada 17 procedure; detail seluruh lokasi pada `index_workload.csv`.

**Contoh konsumen:** `procedures/dbo.OMGetListProductSekunder.sql:271`; `procedures/dbo.OMGetListProductSekunder.sql:285`; `procedures/dbo.OMGetListProductSekunder.sql:304`; `procedures/dbo.OMGetListProductSekunder.sql:302`; `procedures/dbo.OMGetListProductSekunder.sql:314`

**Existing index pada tabel:** TRSECParamSpread_TR_CU1 (ParamId) INCLUDE ()

```sql
CREATE NONCLUSTERED INDEX [IX_Audit20261022_TRSECParamSpread_TR_SecId_ChannelCode_BuySell_NominalFrom]
ON [dbo].[TRSECParamSpread_TR] ([SecId] ASC, [ChannelCode] ASC, [BuySell] ASC, [NominalFrom] ASC)
INCLUDE ([NominalTo], [BidSpread], [OfferSpread], [IndikasiHarga]);
```

## IX08 — dbo.TRSParameterSettlementDate_TR

Lookup settlement per SecId+ChannelCode. Index saat ini berawalan Id.

**Prasyarat:** Channel INT dan ChannelCode VARCHAR(15) adalah field berbeda; kandidat ini tidak menutup lookup SecId+Channel secara penuh.

**Bukti workload:** 39 predicate kesetaraan key pertama pada 16 procedure; detail seluruh lokasi pada `index_workload.csv`.

**Contoh konsumen:** `procedures/dbo.GetParamHargaModalSekunder.sql:8`; `procedures/dbo.OMGetListProductSekunder.sql:371`; `procedures/dbo.OMValidateBalance.sql:68`; `procedures/dbo.RepoSaveParamSecurityDirtyPrice.sql:112`; `procedures/dbo.RepoUploadValidation.sql:238`

**Existing index pada tabel:** TRSParameterSettlementDate_TR_CN1 (Id) INCLUDE ()

```sql
CREATE NONCLUSTERED INDEX [IX_Audit20261022_TRSParameterSettlementDate_TR_SecId_ChannelCode]
ON [dbo].[TRSParameterSettlementDate_TR] ([SecId] ASC, [ChannelCode] ASC)
INCLUDE ([SettlementDateBankBuy], [SettlementDateBankSell]);
```

## IX09 — dbo.TRSParamHargaModal_TR

Cabang lookup berdasarkan SecurityNo; clustered unik saat ini SecId.

**Prasyarat:** Tabel parameter bisa kecil; ukur lebih dahulu. Tidak perlu index SecId tambahan.

**Bukti workload:** 22 predicate kesetaraan key pertama pada 15 procedure; detail seluruh lokasi pada `index_workload.csv`.

**Contoh konsumen:** `procedures/dbo.OMTrxSecPreBook.sql:471`; `procedures/dbo.trs_DefaultSettlementDate.sql:23`; `procedures/dbo.trs_ValidateHargaModal.sql:15`; `procedures/dbo.TRSCekTanggalSuratBerharga.sql:10`; `procedures/dbo.TRSCmprSecurityNo.sql:13`

**Existing index pada tabel:** TRSParamHargaModal_TR_CU1 (SecId) INCLUDE ()

```sql
CREATE NONCLUSTERED INDEX [IX_Audit20261022_TRSParamHargaModal_TR_SecurityNo]
ON [dbo].[TRSParamHargaModal_TR] ([SecurityNo] ASC)
INCLUDE ([HargaModalBeli], [HargaModalJual], [SecCcy]);
```

## IX10 — dbo.PTOSCustomer_TM

Validasi rekening existing/waiting tanpa SID. Index SID,NoRekInvestorRDN tidak melayani seek rekening saja.

**Prasyarat:** Samakan tipe/kolasi input; pertahankan leading zero sesuai kontrak rekening.

**Bukti workload:** 44 predicate kesetaraan key pertama pada 18 procedure; detail seluruh lokasi pada `index_workload.csv`.

**Contoh konsumen:** `procedures/dbo.PTOSJobMasterOutstanding.sql:64`; `procedures/dbo.PTOSJobMasterOutstanding.sql:77`; `procedures/dbo.PTOSProsesDataJobNasabah.sql:127`; `procedures/dbo.PTOSProsesDataJobNasabah.sql:141`; `procedures/dbo.PTOSProsesDataJobNasabah.sql:168`

**Existing index pada tabel:** ix_PTOSCustomer_TM_1 (CIFNo) INCLUDE (); ix_ptoscustomer_tm_obl114 (CIFNo|NamaPerusahaanEfek) INCLUDE (); ix_ptoscustomer_tm_sidnordn (SID|NoRekInvestorRDN) INCLUDE ()

```sql
CREATE NONCLUSTERED INDEX [IX_Audit20261022_PTOSCustomer_TM_NoRekInvestorRDN]
ON [dbo].[PTOSCustomer_TM] ([NoRekInvestorRDN] ASC)
INCLUDE ([CIFNo], [SID], [Status]);
```

## IX11 — dbo.PTOSOutstanding_TR

JOIN outstanding berulang pada SID+rekening+produk+tipe. Existing SID/rekening berhenti sebelum produk/tipe.

**Prasyarat:** Normalisasi Excel NVARCHAR ke typed staging lebih dahulu; bandingkan sebagai konsolidasi ix_ptosoutstanding_tr_sidnordn. Query tanggal-first tetap butuh index tanggal yang ada.

**Bukti workload:** 50 predicate kesetaraan key pertama pada 13 procedure; detail seluruh lokasi pada `index_workload.csv`.

**Contoh konsumen:** `procedures/dbo.PTOSJobMasterOutstanding.sql:64`; `procedures/dbo.PTOSJobMasterOutstanding.sql:78`; `procedures/dbo.PTOSJobMasterOutstanding.sql:78`; `procedures/dbo.PTOSJobMasterOutstanding.sql:77`; `procedures/dbo.PTOSProsesDataJobOutstanding.sql:21`

**Existing index pada tabel:** ix_ptosoutstanding_tr_obl114group (Tanggal|SID|NoRekInvestorRDN|NamaNasabah|KodeProduk|NamaProduk|Ccy|HargaPerolehanNasabah|IndikasiHargaBid|AccruedInterest|PTOSType) INCLUDE (); ix_ptosoutstanding_tr_obl114where (Tanggal|PTOSType|KodeProduk) INCLUDE (); ix_ptosoutstanding_tr_sid (SID) INCLUDE (); ix_ptosoutstanding_tr_sidnordn (SID|NoRekInvestorRDN) INCLUDE (); ix_ptosoutstanding_tr_tanggal (Tanggal) INCLUDE ()

```sql
CREATE NONCLUSTERED INDEX [IX_Audit20261022_PTOSOutstanding_TR_SID_NoRekInvestorRDN_KodeProduk_PTOSType_Tanggal]
ON [dbo].[PTOSOutstanding_TR] ([SID] ASC, [NoRekInvestorRDN] ASC, [KodeProduk] ASC, [PTOSType] ASC, [Tanggal] ASC)
INCLUDE ([HargaPerolehanNasabah], [Status]);
```

## IX12 — dbo.PTOSTransaksi_TR

JOIN/upsert berdasarkan DealId+PTOSType; existing hanya SID.

**Prasyarat:** Ubah staging Deal ID dari NVARCHAR(255) ke VARCHAR(20) setelah validasi tanpa truncation; pertimbangkan cost update.

**Bukti workload:** 38 predicate kesetaraan key pertama pada 10 procedure; detail seluruh lokasi pada `index_workload.csv`.

**Contoh konsumen:** `procedures/dbo.PTOSProsesDataJobTransaksi.sql:31`; `procedures/dbo.PTOSProsesDataJobTransaksi.sql:102`; `procedures/dbo.PTOSProsesDataJobTransaksi.sql:143`; `procedures/dbo.PTOSProsesDataJobTransaksi.sql:167`; `procedures/dbo.PTOSProsesDataJobTransaksiBond.sql:28`

**Existing index pada tabel:** ix_ptostransaksi_tr_sid (SID) INCLUDE ()

```sql
CREATE NONCLUSTERED INDEX [IX_Audit20261022_PTOSTransaksi_TR_DealId_PTOSType]
ON [dbo].[PTOSTransaksi_TR] ([DealId] ASC, [PTOSType] ASC)
INCLUDE ([SID], [KodeProduk], [NoRekInvestorRDN], [Status]);
```

## IX13 — dbo.SecurityCouponDate_TR

Coupon schedule per SecId+Done+rentang Periode. Existing clustered hanya SecId.

**Prasyarat:** Perbaiki Periode VARCHAR(8) vs DATETIME; deploy hanya jika jumlah schedule per SecId cukup besar.

**Bukti workload:** 35 predicate kesetaraan key pertama pada 11 procedure; detail seluruh lokasi pada `index_workload.csv`.

**Contoh konsumen:** `procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql:335`; `procedures/dbo.trs_ListSecurityMaster_TM.sql:414`; `procedures/dbo.TRSRetailCalculateFee.sql:759`; `procedures/dbo.TRSRetailUpdateCouponDate_Dummy.sql:82`; `procedures/dbo.TRSRetailUpdateCouponDate_Dummy.sql:88`

**Existing index pada tabel:** SecurityCouponDate_TR_CN1 (SecId) INCLUDE ()

```sql
CREATE NONCLUSTERED INDEX [IX_Audit20261022_SecurityCouponDate_TR_SecId_Done_Periode]
ON [dbo].[SecurityCouponDate_TR] ([SecId] ASC, [Done] ASC, [Periode] ASC)
INCLUDE ([Sequence], [SecurityNo]);
```

## IX14 — dbo.TRSCustAccNumber_TR

fnGetOBLAccByPriority: TOP(1) WHERE CIFId ORDER BY RowId DESC. NN1 hanya CIFId.

**Prasyarat:** Bandingkan dengan NN1 sebagai alternatif konsolidasi; manfaat bergantung apakah UDF di-inline atau dipanggil per baris.

**Bukti workload:** 129 predicate kesetaraan key pertama pada 68 procedure; detail seluruh lokasi pada `index_workload.csv`.

**Contoh konsumen:** `procedures/dbo.GetRekeningMasterNasabah.sql:153`; `procedures/dbo.GetRekeningMasterNasabah.sql:178`; `procedures/dbo.GetRekeningMasterNasabahTRS.sql:79`; `procedures/dbo.GetRekeningMasterNasabahTRS.sql:208`; `procedures/dbo.OMSubmitTransactionSekunder.sql:243`

**Existing index pada tabel:** IUNC_TRSCustAccNumber_TR1 (CcyCode|NoRekInvestor|NamaPemilik|IsJoin|Relasi|ProductCode) INCLUDE (); TRSCustAccNumber_TR_NN1 (CIFId) INCLUDE (); PK_TRSCustAccNumber_TR_RowId (RowId) INCLUDE (); PK_TRSCustAccNumber_TR (CIFId|CcyCode|NoRekInvestor) INCLUDE ()

```sql
CREATE NONCLUSTERED INDEX [IX_Audit20261022_TRSCustAccNumber_TR_CIFId_RowId]
ON [dbo].[TRSCustAccNumber_TR] ([CIFId] ASC, [RowId] DESC)
INCLUDE ([NoRekInvestor], [CcyCode]);
```

## IX15 — dbo.TRSRetailCustBlock_TM

JOIN block/release ke transaksi berdasarkan DealId; unique existing berawalan CollateralId.

**Prasyarat:** Ukur update cost pada status/block lifecycle; tidak mengasumsikan satu DealId hanya satu collateral.

**Bukti workload:** 19 predicate kesetaraan key pertama pada 13 procedure; detail seluruh lokasi pada `index_workload.csv`.

**Contoh konsumen:** `procedures/dbo.Rvmp_trs_ListSecurityTransactionForInsertEditBlockTrans.sql:77`; `procedures/dbo.trs_InsertReleaseTRSRetailCustBlock_TM.sql:78`; `procedures/dbo.trs_InsertReleaseTRSRetailCustBlock_TM.sql:100`; `procedures/dbo.trs_InsertReleaseTRSRetailCustBlock_TM.sql:106`; `procedures/dbo.trs_InsertTrx2ndSecurityTransaction_TT.sql:390`

**Existing index pada tabel:** IU_TRSRetailCustBlock_TM (CollateralId|CIFId|SecId|CollateralNo|ExpiryDate) INCLUDE ()

```sql
CREATE NONCLUSTERED INDEX [IX_Audit20261022_TRSRetailCustBlock_TM_DealId]
ON [dbo].[TRSRetailCustBlock_TM] ([DealId] ASC)
INCLUDE ([CollateralStatus], [CIFId], [SecId], [CollateralNo]);
```

## IX16 — dbo.TRSRetailCustBlock_TM

Pencarian collateral nasabah+status; tidak sama dengan lookup DealId pada IX15.

**Prasyarat:** Pilih dari Query Store; jangan otomatis pasang IX15 dan IX16 sekaligus pada tabel write-heavy.

**Bukti workload:** 28 predicate kesetaraan key pertama pada 10 procedure; detail seluruh lokasi pada `index_workload.csv`.

**Contoh konsumen:** `procedures/dbo.trs_ListSecurityTransactionForInsertEditBlockTrans_REPO.sql:112`; `procedures/dbo.trs_ListTRSRetailCustBlock_TM.sql:10`; `procedures/dbo.trs_ListTRSRetailCustBlock_TM.sql:34`; `procedures/dbo.trs_ListTRSRetailCustBlockForRelease_TM.sql:10`; `procedures/dbo.trs_ListTRSRetailCustBlockForRelease_TM.sql:34`

**Existing index pada tabel:** IU_TRSRetailCustBlock_TM (CollateralId|CIFId|SecId|CollateralNo|ExpiryDate) INCLUDE ()

```sql
CREATE NONCLUSTERED INDEX [IX_Audit20261022_TRSRetailCustBlock_TM_CIFId_CollateralStatus_SecId]
ON [dbo].[TRSRetailCustBlock_TM] ([CIFId] ASC, [CollateralStatus] ASC, [SecId] ASC)
INCLUDE ([DealId], [ExpiryDate], [NIK_Inputer]);
```

## IX17 — dbo.CFMAST

Pencarian CIF berdasarkan identitas+jenis dokumen+tanggal lahir. Perluas/alternatif idx_cfmast_cfssno yang sudah ada.

**Prasyarat:** Index di SQL_REPLICATE, bukan di CFMAST_v. Betulkan Unicode KTP dan CFBIRD CHAR(7) vs INT dahulu. Evaluasi replication/DML dan jangan drop idx_cfmast_cfssno tanpa evidence.

**Bukti workload:** 4 predicate kesetaraan key pertama pada 4 procedure; detail seluruh lokasi pada `index_workload.csv`.

**Contoh konsumen:** `procedures/dbo.PTOSProsesDataJobNasabah.sql:94`; `procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql:109`; `procedures/dbo.TRSDataNasabahRetailInstitusi.sql:333`; `procedures/dbo.TRSPopulatePTOSUpload.sql:261`

**Existing index pada tabel:** idx_cfmast_cfcif# (CFCIF#) INCLUDE (); idx_cfmast_cfcif (CFCIF) INCLUDE (); idx_cfmast_cfssno (CFSSNO) INCLUDE (); idx_cfmast_cfsnme (CFSNME) INCLUDE ()

```sql
CREATE NONCLUSTERED INDEX [IX_Audit20261022_CFMAST_CFSSNO_CFSSCD_CFBIRD]
ON [dbo].[CFMAST] ([CFSSNO] ASC, [CFSSCD] ASC, [CFBIRD] ASC)
INCLUDE ([CFCIF], [CFSNME]);
```

## IX18 — dbo.ProCIFChangeMaster_TH

Laporan perubahan saat ini memfilter tanggal setelah UNION dimaterialisasi. Setelah date predicate didorong ke source, kandidat date-range ini relevan.

**Prasyarat:** Wajib pushdown date filter terlebih dahulu dan verifikasi semantik rentang hari. Existing index CIFNumber-first tetap diperlukan workload lain.

**Bukti workload:** 0 predicate kesetaraan key pertama pada 0 procedure; detail seluruh lokasi pada `index_workload.csv`.

**Contoh konsumen:** `procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql:251` (date pushdown yang diusulkan).

**Existing index pada tabel:** IX_ProCIFChangeMaster_TH_New_Housekeep_1 (ChangeId) INCLUDE (); ProCIFChangeMaster_TH_NN_1 (TableName|CIFNumber) INCLUDE (ChangeId|ChangeDate); ProCIFChangeMaster_TH_NN_2 (TellerId) INCLUDE (ChangeId|Branch|CIFNumber|CIFName|Status|ChangeType|ChangeDate|CheckerSuid|ApprovedDate); ProCIFChangeMaster_TH_NN_3 (Status) INCLUDE (ChangeId|Branch|CIFNumber|CIFName|ChangeType|ChangeDate|TellerId|CheckerSuid|ApprovedDate); ProCIFChangeMaster_TH_NN_4 (CIFNumber|TableName|Status) INCLUDE (ChangeId|ApprovedDate); ProCIFChangeMaster_TH_NN_5 (CIFNumber|TableName|Status) INCLUDE (ChangeId|ApprovedDate|Branch|TellerId)

```sql
CREATE NONCLUSTERED INDEX [IX_Audit20261022_ProCIFChangeMaster_TH_ChangeDate_ChangeId]
ON [dbo].[ProCIFChangeMaster_TH] ([ChangeDate] ASC, [ChangeId] ASC)
INCLUDE ([CIFNumber], [CIFName], [Branch], [TellerId]);
```
