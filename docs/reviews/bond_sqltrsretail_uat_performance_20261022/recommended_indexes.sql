-- PROPOSAL ONLY, generated 2026-10-22. No statement has been applied.
-- Choose ONE candidate and its database after reviewing REPORT.md and live index definitions.
-- A/B/C are review order, not measured performance benefit.
-- SQL Server version/edition, replication policy, build window, log/tempdb space remain unverified.
DECLARE @CandidateId varchar(10) = NULL; -- e.g. 'IX01'; IX17/IX18 require SQL_REPLICATE context
DECLARE @ExpectedRetailDatabase sysname = NULL; -- actual retail UAT DB name, required for IX01..IX16
IF @CandidateId IS NULL THROW 51000, 'Set one reviewed CandidateId; this script defaults to no DDL.', 1;
IF @CandidateId NOT IN ('IX01', 'IX02', 'IX03', 'IX04', 'IX05', 'IX06', 'IX07', 'IX08', 'IX09', 'IX10', 'IX11', 'IX12', 'IX13', 'IX14', 'IX15', 'IX16', 'IX17', 'IX18') THROW 51000, 'Unknown CandidateId.', 1;

-- IX01: Lookup SecurityNo lintas workload. IU_SecurityMaster_TM berawalan SecId sehingga bukan seek key untuk SecurityNo saja.
-- Prerequisite: Pastikan input SecurityNo tetap VARCHAR, bukan ID numerik; ukur ukuran tabel dan lookup yang tersisa.
IF @CandidateId = 'IX01'
BEGIN
    IF @ExpectedRetailDatabase IS NULL OR DB_NAME() <> @ExpectedRetailDatabase THROW 51000, 'Set and select the verified retail UAT database.', 1;
    IF OBJECT_ID(N'dbo.SecurityMaster_TM', N'U') IS NULL THROW 51000, 'Expected base table missing.', 1;
    IF EXISTS (SELECT 1 FROM sys.indexes WHERE object_id = OBJECT_ID(N'dbo.SecurityMaster_TM') AND name = N'IX_Audit20261022_SecurityMaster_TM_SecurityNo') THROW 51000, 'Candidate name already exists; compare definitions.', 1;
    -- Compare equivalent/prefix keys and INCLUDE in catalog_validation.sql before proceeding.
    EXEC sys.sp_executesql N'CREATE NONCLUSTERED INDEX [IX_Audit20261022_SecurityMaster_TM_SecurityNo]
ON [dbo].[SecurityMaster_TM] ([SecurityNo] ASC)
INCLUDE ([SecCcy], [SecType], [ISINCode], [InterestTypeID]);';
END;

-- IX02: Lookup IdSeri pada flow ORI/SBN. Belum ada index berawalan IdSeri.
-- Prerequisite: Perbaiki pembandingan IdSeri VARCHAR(100) dengan INT dahulu; definisikan aturan leading zero dan nilai nonnumerik.
IF @CandidateId = 'IX02'
BEGIN
    IF @ExpectedRetailDatabase IS NULL OR DB_NAME() <> @ExpectedRetailDatabase THROW 51000, 'Set and select the verified retail UAT database.', 1;
    IF OBJECT_ID(N'dbo.SecurityMaster_TM', N'U') IS NULL THROW 51000, 'Expected base table missing.', 1;
    IF EXISTS (SELECT 1 FROM sys.indexes WHERE object_id = OBJECT_ID(N'dbo.SecurityMaster_TM') AND name = N'IX_Audit20261022_SecurityMaster_TM_IdSeri') THROW 51000, 'Candidate name already exists; compare definitions.', 1;
    -- Compare equivalent/prefix keys and INCLUDE in catalog_validation.sql before proceeding.
    EXEC sys.sp_executesql N'CREATE NONCLUSTERED INDEX [IX_Audit20261022_SecurityMaster_TM_IdSeri]
ON [dbo].[SecurityMaster_TM] ([IdSeri] ASC)
INCLUDE ([SecurityNo], [SecCcy]);';
END;

-- IX03: Lookup DealNo; IU_SecurityTransaction_TT berawalan DealId. PK clustered DealId sudah menjadi row locator implisit.
-- Prerequisite: Periksa panjang parameter dan query aktual; jangan jadikan UNIQUE tanpa bukti data.
IF @CandidateId = 'IX03'
BEGIN
    IF @ExpectedRetailDatabase IS NULL OR DB_NAME() <> @ExpectedRetailDatabase THROW 51000, 'Set and select the verified retail UAT database.', 1;
    IF OBJECT_ID(N'dbo.SecurityTransaction_TT', N'U') IS NULL THROW 51000, 'Expected base table missing.', 1;
    IF EXISTS (SELECT 1 FROM sys.indexes WHERE object_id = OBJECT_ID(N'dbo.SecurityTransaction_TT') AND name = N'IX_Audit20261022_SecurityTransaction_TT_DealNo') THROW 51000, 'Candidate name already exists; compare definitions.', 1;
    -- Compare equivalent/prefix keys and INCLUDE in catalog_validation.sql before proceeding.
    EXEC sys.sp_executesql N'CREATE NONCLUSTERED INDEX [IX_Audit20261022_SecurityTransaction_TT_DealNo]
ON [dbo].[SecurityTransaction_TT] ([DealNo] ASC)
INCLUDE ([CIFId], [SecId], [TrxType], [TrxStatus], [SettlementDate]);';
END;

-- IX04: Pola saldo/coupon: equality CIFId,SecId,TrxType,TrxStatus lalu rentang SettlementDate. NN3 hanya CIFId,SecId INCLUDE TrxType.
-- Prerequisite: Bandingkan dengan NN3/NN8 dan index TrxType,TrxStatus. Tidak semua query mengisi seluruh equality prefix; jangan otomatis drop index lama.
IF @CandidateId = 'IX04'
BEGIN
    IF @ExpectedRetailDatabase IS NULL OR DB_NAME() <> @ExpectedRetailDatabase THROW 51000, 'Set and select the verified retail UAT database.', 1;
    IF OBJECT_ID(N'dbo.SecurityTransaction_TT', N'U') IS NULL THROW 51000, 'Expected base table missing.', 1;
    IF EXISTS (SELECT 1 FROM sys.indexes WHERE object_id = OBJECT_ID(N'dbo.SecurityTransaction_TT') AND name = N'IX_Audit20261022_SecurityTransaction_TT_CIFId_SecId_TrxType_TrxStatus_SettlementDate') THROW 51000, 'Candidate name already exists; compare definitions.', 1;
    -- Compare equivalent/prefix keys and INCLUDE in catalog_validation.sql before proceeding.
    EXEC sys.sp_executesql N'CREATE NONCLUSTERED INDEX [IX_Audit20261022_SecurityTransaction_TT_CIFId_SecId_TrxType_TrxStatus_SettlementDate]
ON [dbo].[SecurityTransaction_TT] ([CIFId] ASC, [SecId] ASC, [TrxType] ASC, [TrxStatus] ASC, [SettlementDate] ASC)
INCLUDE ([FaceValue], [InsertedBy]);';
END;

-- IX05: Lookup order per surat berharga/nasabah; SecId hanya INCLUDE pada IU_ORI_Order_TT berawalan OrderId.
-- Prerequisite: Pilih urutan SecId,CIFId dari selektivitas workload; query dengan OrderId tunggal sudah terlayani PK.
IF @CandidateId = 'IX05'
BEGIN
    IF @ExpectedRetailDatabase IS NULL OR DB_NAME() <> @ExpectedRetailDatabase THROW 51000, 'Set and select the verified retail UAT database.', 1;
    IF OBJECT_ID(N'dbo.ORI_Order_TT', N'U') IS NULL THROW 51000, 'Expected base table missing.', 1;
    IF EXISTS (SELECT 1 FROM sys.indexes WHERE object_id = OBJECT_ID(N'dbo.ORI_Order_TT') AND name = N'IX_Audit20261022_ORI_Order_TT_SecId_CIFId') THROW 51000, 'Candidate name already exists; compare definitions.', 1;
    -- Compare equivalent/prefix keys and INCLUDE in catalog_validation.sql before proceeding.
    EXEC sys.sp_executesql N'CREATE NONCLUSTERED INDEX [IX_Audit20261022_ORI_Order_TT_SecId_CIFId]
ON [dbo].[ORI_Order_TT] ([SecId] ASC, [CIFId] ASC)
INCLUDE ([OrderDate], [OrderStatus], [IsBuyBack]);';
END;

-- IX06: JOIN detail pada DealId. IU_SecurityTransactionDetail_TT berawalan RowId, bukan DealId.
-- Prerequisite: Projection lebar mungkin tetap memerlukan RID lookup; tambah INCLUDE hanya berdasarkan actual plan.
IF @CandidateId = 'IX06'
BEGIN
    IF @ExpectedRetailDatabase IS NULL OR DB_NAME() <> @ExpectedRetailDatabase THROW 51000, 'Set and select the verified retail UAT database.', 1;
    IF OBJECT_ID(N'dbo.SecurityTransactionDetail_TT', N'U') IS NULL THROW 51000, 'Expected base table missing.', 1;
    IF EXISTS (SELECT 1 FROM sys.indexes WHERE object_id = OBJECT_ID(N'dbo.SecurityTransactionDetail_TT') AND name = N'IX_Audit20261022_SecurityTransactionDetail_TT_DealId') THROW 51000, 'Candidate name already exists; compare definitions.', 1;
    -- Compare equivalent/prefix keys and INCLUDE in catalog_validation.sql before proceeding.
    EXEC sys.sp_executesql N'CREATE NONCLUSTERED INDEX [IX_Audit20261022_SecurityTransactionDetail_TT_DealId]
ON [dbo].[SecurityTransactionDetail_TT] ([DealId] ASC)
INCLUDE ([TrxDetailType], [CIFIdDebet], [CIFIdKredit]);';
END;

-- IX07: Pola harga produk: SecId+ChannelCode+BuySell lalu nominal tier. Index saat ini hanya ParamId.
-- Prerequisite: Parameter table mungkin kecil; jangan deploy bila scan murah. Predicate NominalTo masih residual untuk interval overlap.
IF @CandidateId = 'IX07'
BEGIN
    IF @ExpectedRetailDatabase IS NULL OR DB_NAME() <> @ExpectedRetailDatabase THROW 51000, 'Set and select the verified retail UAT database.', 1;
    IF OBJECT_ID(N'dbo.TRSECParamSpread_TR', N'U') IS NULL THROW 51000, 'Expected base table missing.', 1;
    IF EXISTS (SELECT 1 FROM sys.indexes WHERE object_id = OBJECT_ID(N'dbo.TRSECParamSpread_TR') AND name = N'IX_Audit20261022_TRSECParamSpread_TR_SecId_ChannelCode_BuySell_NominalFrom') THROW 51000, 'Candidate name already exists; compare definitions.', 1;
    -- Compare equivalent/prefix keys and INCLUDE in catalog_validation.sql before proceeding.
    EXEC sys.sp_executesql N'CREATE NONCLUSTERED INDEX [IX_Audit20261022_TRSECParamSpread_TR_SecId_ChannelCode_BuySell_NominalFrom]
ON [dbo].[TRSECParamSpread_TR] ([SecId] ASC, [ChannelCode] ASC, [BuySell] ASC, [NominalFrom] ASC)
INCLUDE ([NominalTo], [BidSpread], [OfferSpread], [IndikasiHarga]);';
END;

-- IX08: Lookup settlement per SecId+ChannelCode. Index saat ini berawalan Id.
-- Prerequisite: Channel INT dan ChannelCode VARCHAR(15) adalah field berbeda; kandidat ini tidak menutup lookup SecId+Channel secara penuh.
IF @CandidateId = 'IX08'
BEGIN
    IF @ExpectedRetailDatabase IS NULL OR DB_NAME() <> @ExpectedRetailDatabase THROW 51000, 'Set and select the verified retail UAT database.', 1;
    IF OBJECT_ID(N'dbo.TRSParameterSettlementDate_TR', N'U') IS NULL THROW 51000, 'Expected base table missing.', 1;
    IF EXISTS (SELECT 1 FROM sys.indexes WHERE object_id = OBJECT_ID(N'dbo.TRSParameterSettlementDate_TR') AND name = N'IX_Audit20261022_TRSParameterSettlementDate_TR_SecId_ChannelCode') THROW 51000, 'Candidate name already exists; compare definitions.', 1;
    -- Compare equivalent/prefix keys and INCLUDE in catalog_validation.sql before proceeding.
    EXEC sys.sp_executesql N'CREATE NONCLUSTERED INDEX [IX_Audit20261022_TRSParameterSettlementDate_TR_SecId_ChannelCode]
ON [dbo].[TRSParameterSettlementDate_TR] ([SecId] ASC, [ChannelCode] ASC)
INCLUDE ([SettlementDateBankBuy], [SettlementDateBankSell]);';
END;

-- IX09: Cabang lookup berdasarkan SecurityNo; clustered unik saat ini SecId.
-- Prerequisite: Tabel parameter bisa kecil; ukur lebih dahulu. Tidak perlu index SecId tambahan.
IF @CandidateId = 'IX09'
BEGIN
    IF @ExpectedRetailDatabase IS NULL OR DB_NAME() <> @ExpectedRetailDatabase THROW 51000, 'Set and select the verified retail UAT database.', 1;
    IF OBJECT_ID(N'dbo.TRSParamHargaModal_TR', N'U') IS NULL THROW 51000, 'Expected base table missing.', 1;
    IF EXISTS (SELECT 1 FROM sys.indexes WHERE object_id = OBJECT_ID(N'dbo.TRSParamHargaModal_TR') AND name = N'IX_Audit20261022_TRSParamHargaModal_TR_SecurityNo') THROW 51000, 'Candidate name already exists; compare definitions.', 1;
    -- Compare equivalent/prefix keys and INCLUDE in catalog_validation.sql before proceeding.
    EXEC sys.sp_executesql N'CREATE NONCLUSTERED INDEX [IX_Audit20261022_TRSParamHargaModal_TR_SecurityNo]
ON [dbo].[TRSParamHargaModal_TR] ([SecurityNo] ASC)
INCLUDE ([HargaModalBeli], [HargaModalJual], [SecCcy]);';
END;

-- IX10: Validasi rekening existing/waiting tanpa SID. Index SID,NoRekInvestorRDN tidak melayani seek rekening saja.
-- Prerequisite: Samakan tipe/kolasi input; pertahankan leading zero sesuai kontrak rekening.
IF @CandidateId = 'IX10'
BEGIN
    IF @ExpectedRetailDatabase IS NULL OR DB_NAME() <> @ExpectedRetailDatabase THROW 51000, 'Set and select the verified retail UAT database.', 1;
    IF OBJECT_ID(N'dbo.PTOSCustomer_TM', N'U') IS NULL THROW 51000, 'Expected base table missing.', 1;
    IF EXISTS (SELECT 1 FROM sys.indexes WHERE object_id = OBJECT_ID(N'dbo.PTOSCustomer_TM') AND name = N'IX_Audit20261022_PTOSCustomer_TM_NoRekInvestorRDN') THROW 51000, 'Candidate name already exists; compare definitions.', 1;
    -- Compare equivalent/prefix keys and INCLUDE in catalog_validation.sql before proceeding.
    EXEC sys.sp_executesql N'CREATE NONCLUSTERED INDEX [IX_Audit20261022_PTOSCustomer_TM_NoRekInvestorRDN]
ON [dbo].[PTOSCustomer_TM] ([NoRekInvestorRDN] ASC)
INCLUDE ([CIFNo], [SID], [Status]);';
END;

-- IX11: JOIN outstanding berulang pada SID+rekening+produk+tipe. Existing SID/rekening berhenti sebelum produk/tipe.
-- Prerequisite: Normalisasi Excel NVARCHAR ke typed staging lebih dahulu; bandingkan sebagai konsolidasi ix_ptosoutstanding_tr_sidnordn. Query tanggal-first tetap butuh index tanggal yang ada.
IF @CandidateId = 'IX11'
BEGIN
    IF @ExpectedRetailDatabase IS NULL OR DB_NAME() <> @ExpectedRetailDatabase THROW 51000, 'Set and select the verified retail UAT database.', 1;
    IF OBJECT_ID(N'dbo.PTOSOutstanding_TR', N'U') IS NULL THROW 51000, 'Expected base table missing.', 1;
    IF EXISTS (SELECT 1 FROM sys.indexes WHERE object_id = OBJECT_ID(N'dbo.PTOSOutstanding_TR') AND name = N'IX_Audit20261022_PTOSOutstanding_TR_SID_NoRekInvestorRDN_KodeProduk_PTOSType_Tanggal') THROW 51000, 'Candidate name already exists; compare definitions.', 1;
    -- Compare equivalent/prefix keys and INCLUDE in catalog_validation.sql before proceeding.
    EXEC sys.sp_executesql N'CREATE NONCLUSTERED INDEX [IX_Audit20261022_PTOSOutstanding_TR_SID_NoRekInvestorRDN_KodeProduk_PTOSType_Tanggal]
ON [dbo].[PTOSOutstanding_TR] ([SID] ASC, [NoRekInvestorRDN] ASC, [KodeProduk] ASC, [PTOSType] ASC, [Tanggal] ASC)
INCLUDE ([HargaPerolehanNasabah], [Status]);';
END;

-- IX12: JOIN/upsert berdasarkan DealId+PTOSType; existing hanya SID.
-- Prerequisite: Ubah staging Deal ID dari NVARCHAR(255) ke VARCHAR(20) setelah validasi tanpa truncation; pertimbangkan cost update.
IF @CandidateId = 'IX12'
BEGIN
    IF @ExpectedRetailDatabase IS NULL OR DB_NAME() <> @ExpectedRetailDatabase THROW 51000, 'Set and select the verified retail UAT database.', 1;
    IF OBJECT_ID(N'dbo.PTOSTransaksi_TR', N'U') IS NULL THROW 51000, 'Expected base table missing.', 1;
    IF EXISTS (SELECT 1 FROM sys.indexes WHERE object_id = OBJECT_ID(N'dbo.PTOSTransaksi_TR') AND name = N'IX_Audit20261022_PTOSTransaksi_TR_DealId_PTOSType') THROW 51000, 'Candidate name already exists; compare definitions.', 1;
    -- Compare equivalent/prefix keys and INCLUDE in catalog_validation.sql before proceeding.
    EXEC sys.sp_executesql N'CREATE NONCLUSTERED INDEX [IX_Audit20261022_PTOSTransaksi_TR_DealId_PTOSType]
ON [dbo].[PTOSTransaksi_TR] ([DealId] ASC, [PTOSType] ASC)
INCLUDE ([SID], [KodeProduk], [NoRekInvestorRDN], [Status]);';
END;

-- IX13: Coupon schedule per SecId+Done+rentang Periode. Existing clustered hanya SecId.
-- Prerequisite: Perbaiki Periode VARCHAR(8) vs DATETIME; deploy hanya jika jumlah schedule per SecId cukup besar.
IF @CandidateId = 'IX13'
BEGIN
    IF @ExpectedRetailDatabase IS NULL OR DB_NAME() <> @ExpectedRetailDatabase THROW 51000, 'Set and select the verified retail UAT database.', 1;
    IF OBJECT_ID(N'dbo.SecurityCouponDate_TR', N'U') IS NULL THROW 51000, 'Expected base table missing.', 1;
    IF EXISTS (SELECT 1 FROM sys.indexes WHERE object_id = OBJECT_ID(N'dbo.SecurityCouponDate_TR') AND name = N'IX_Audit20261022_SecurityCouponDate_TR_SecId_Done_Periode') THROW 51000, 'Candidate name already exists; compare definitions.', 1;
    -- Compare equivalent/prefix keys and INCLUDE in catalog_validation.sql before proceeding.
    EXEC sys.sp_executesql N'CREATE NONCLUSTERED INDEX [IX_Audit20261022_SecurityCouponDate_TR_SecId_Done_Periode]
ON [dbo].[SecurityCouponDate_TR] ([SecId] ASC, [Done] ASC, [Periode] ASC)
INCLUDE ([Sequence], [SecurityNo]);';
END;

-- IX14: fnGetOBLAccByPriority: TOP(1) WHERE CIFId ORDER BY RowId DESC. NN1 hanya CIFId.
-- Prerequisite: Bandingkan dengan NN1 sebagai alternatif konsolidasi; manfaat bergantung apakah UDF di-inline atau dipanggil per baris.
IF @CandidateId = 'IX14'
BEGIN
    IF @ExpectedRetailDatabase IS NULL OR DB_NAME() <> @ExpectedRetailDatabase THROW 51000, 'Set and select the verified retail UAT database.', 1;
    IF OBJECT_ID(N'dbo.TRSCustAccNumber_TR', N'U') IS NULL THROW 51000, 'Expected base table missing.', 1;
    IF EXISTS (SELECT 1 FROM sys.indexes WHERE object_id = OBJECT_ID(N'dbo.TRSCustAccNumber_TR') AND name = N'IX_Audit20261022_TRSCustAccNumber_TR_CIFId_RowId') THROW 51000, 'Candidate name already exists; compare definitions.', 1;
    -- Compare equivalent/prefix keys and INCLUDE in catalog_validation.sql before proceeding.
    EXEC sys.sp_executesql N'CREATE NONCLUSTERED INDEX [IX_Audit20261022_TRSCustAccNumber_TR_CIFId_RowId]
ON [dbo].[TRSCustAccNumber_TR] ([CIFId] ASC, [RowId] DESC)
INCLUDE ([NoRekInvestor], [CcyCode]);';
END;

-- IX15: JOIN block/release ke transaksi berdasarkan DealId; unique existing berawalan CollateralId.
-- Prerequisite: Ukur update cost pada status/block lifecycle; tidak mengasumsikan satu DealId hanya satu collateral.
IF @CandidateId = 'IX15'
BEGIN
    IF @ExpectedRetailDatabase IS NULL OR DB_NAME() <> @ExpectedRetailDatabase THROW 51000, 'Set and select the verified retail UAT database.', 1;
    IF OBJECT_ID(N'dbo.TRSRetailCustBlock_TM', N'U') IS NULL THROW 51000, 'Expected base table missing.', 1;
    IF EXISTS (SELECT 1 FROM sys.indexes WHERE object_id = OBJECT_ID(N'dbo.TRSRetailCustBlock_TM') AND name = N'IX_Audit20261022_TRSRetailCustBlock_TM_DealId') THROW 51000, 'Candidate name already exists; compare definitions.', 1;
    -- Compare equivalent/prefix keys and INCLUDE in catalog_validation.sql before proceeding.
    EXEC sys.sp_executesql N'CREATE NONCLUSTERED INDEX [IX_Audit20261022_TRSRetailCustBlock_TM_DealId]
ON [dbo].[TRSRetailCustBlock_TM] ([DealId] ASC)
INCLUDE ([CollateralStatus], [CIFId], [SecId], [CollateralNo]);';
END;

-- IX16: Pencarian collateral nasabah+status; tidak sama dengan lookup DealId pada IX15.
-- Prerequisite: Pilih dari Query Store; jangan otomatis pasang IX15 dan IX16 sekaligus pada tabel write-heavy.
IF @CandidateId = 'IX16'
BEGIN
    IF @ExpectedRetailDatabase IS NULL OR DB_NAME() <> @ExpectedRetailDatabase THROW 51000, 'Set and select the verified retail UAT database.', 1;
    IF OBJECT_ID(N'dbo.TRSRetailCustBlock_TM', N'U') IS NULL THROW 51000, 'Expected base table missing.', 1;
    IF EXISTS (SELECT 1 FROM sys.indexes WHERE object_id = OBJECT_ID(N'dbo.TRSRetailCustBlock_TM') AND name = N'IX_Audit20261022_TRSRetailCustBlock_TM_CIFId_CollateralStatus_SecId') THROW 51000, 'Candidate name already exists; compare definitions.', 1;
    -- Compare equivalent/prefix keys and INCLUDE in catalog_validation.sql before proceeding.
    EXEC sys.sp_executesql N'CREATE NONCLUSTERED INDEX [IX_Audit20261022_TRSRetailCustBlock_TM_CIFId_CollateralStatus_SecId]
ON [dbo].[TRSRetailCustBlock_TM] ([CIFId] ASC, [CollateralStatus] ASC, [SecId] ASC)
INCLUDE ([DealId], [ExpiryDate], [NIK_Inputer]);';
END;

-- IX17: Pencarian CIF berdasarkan identitas+jenis dokumen+tanggal lahir. Perluas/alternatif idx_cfmast_cfssno yang sudah ada.
-- Prerequisite: Index di SQL_REPLICATE, bukan di CFMAST_v. Betulkan Unicode KTP dan CFBIRD CHAR(7) vs INT dahulu. Evaluasi replication/DML dan jangan drop idx_cfmast_cfssno tanpa evidence.
IF @CandidateId = 'IX17'
BEGIN
    IF DB_NAME() <> N'SQL_REPLICATE' THROW 51000, 'Use SQL_REPLICATE database for this candidate.', 1;
    IF OBJECT_ID(N'dbo.CFMAST', N'U') IS NULL THROW 51000, 'Expected base table missing.', 1;
    IF EXISTS (SELECT 1 FROM sys.indexes WHERE object_id = OBJECT_ID(N'dbo.CFMAST') AND name = N'IX_Audit20261022_CFMAST_CFSSNO_CFSSCD_CFBIRD') THROW 51000, 'Candidate name already exists; compare definitions.', 1;
    -- Compare equivalent/prefix keys and INCLUDE in catalog_validation.sql before proceeding.
    EXEC sys.sp_executesql N'CREATE NONCLUSTERED INDEX [IX_Audit20261022_CFMAST_CFSSNO_CFSSCD_CFBIRD]
ON [dbo].[CFMAST] ([CFSSNO] ASC, [CFSSCD] ASC, [CFBIRD] ASC)
INCLUDE ([CFCIF], [CFSNME]);';
END;

-- IX18: Laporan perubahan saat ini memfilter tanggal setelah UNION dimaterialisasi. Setelah date predicate didorong ke source, kandidat date-range ini relevan.
-- Prerequisite: Wajib pushdown date filter terlebih dahulu dan verifikasi semantik rentang hari. Existing index CIFNumber-first tetap diperlukan workload lain.
IF @CandidateId = 'IX18'
BEGIN
    IF DB_NAME() <> N'SQL_REPLICATE' THROW 51000, 'Use SQL_REPLICATE database for this candidate.', 1;
    IF OBJECT_ID(N'dbo.ProCIFChangeMaster_TH', N'U') IS NULL THROW 51000, 'Expected base table missing.', 1;
    IF EXISTS (SELECT 1 FROM sys.indexes WHERE object_id = OBJECT_ID(N'dbo.ProCIFChangeMaster_TH') AND name = N'IX_Audit20261022_ProCIFChangeMaster_TH_ChangeDate_ChangeId') THROW 51000, 'Candidate name already exists; compare definitions.', 1;
    -- Compare equivalent/prefix keys and INCLUDE in catalog_validation.sql before proceeding.
    EXEC sys.sp_executesql N'CREATE NONCLUSTERED INDEX [IX_Audit20261022_ProCIFChangeMaster_TH_ChangeDate_ChangeId]
ON [dbo].[ProCIFChangeMaster_TH] ([ChangeDate] ASC, [ChangeId] ASC)
INCLUDE ([CIFNumber], [CIFName], [Branch], [TellerId]);';
END;
