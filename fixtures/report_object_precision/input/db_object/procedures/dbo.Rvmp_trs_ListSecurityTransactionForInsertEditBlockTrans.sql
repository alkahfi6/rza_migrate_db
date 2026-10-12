CREATE OR ALTER PROCEDURE dbo.Rvmp_trs_ListSecurityTransactionForInsertEditBlockTrans @pvSecurityNo VARCHAR (20), @pvSecAccNo VARCHAR (10)
AS
SET NOCOUNT ON;
DECLARE @message AS VARCHAR (80);
DECLARE @inErrNo AS INT;
DECLARE @dToday AS DATETIME, @dMaturityDate AS DATETIME;
SELECT @dToday = current_working_date
FROM   dbo.control_table
WHERE  office_id = 'TRS';
SELECT @dMaturityDate = MaturityDate
FROM   dbo.SecurityMaster_TM
WHERE  SecurityNo = @pvSecurityNo;
CREATE TABLE #tSumLinkTrx (
	DealId     BIGINT,
	SumNominal MONEY
);
CREATE TABLE #tCustBlock (
	DealId  BIGINT,
	Nominal MONEY
);
SELECT CAST (0 AS BIT) AS 'SelectBit',
	tctm.SecAccNo,
	CASE
		WHEN isnull(sttt.isTaxAmnesty, 0) = 0 THEN 'Non Tax Amnesty'
		ELSE 'Tax Amnesty'
	END AS isTaxAmnesty,
	sttt.DealId,
	CAST (0 AS MONEY) AS 'OutStandindBlokir',
	sttt.TrxDate,
	sttt.SettlementDate,
	sttt.DealPrice,
	CASE
		WHEN pbbl.DealId IS NULL THEN sttt.FaceValue - ISNULL(ord.NominalBuyBack, 0)
		ELSE sttt.FaceValue
	END AS FaceValue,
	CAST (0 AS MONEY) AS Blocked,
	tctm.CIFNo,
	sttt.TrxType,
	stdtt.AcquitionDate,
	smtm.NominalMinimum,
	sttt.CIFId,
	sttt.SecId,
	sttt.InsertedDate
INTO   #tSecurityTransaction
FROM   dbo.SecurityTransaction_TT AS sttt
INNER JOIN
dbo.TreasuryCustomer_TM AS tctm
	ON sttt.CIFId = tctm.CIFId
INNER JOIN
dbo.SecurityMaster_TM AS smtm
	ON sttt.SecId = smtm.SecId
LEFT OUTER JOIN
dbo.SecurityTransactionDetail_TT AS stdtt
	ON sttt.DealId = stdtt.DealId
		AND sttt.TrxType = 10
LEFT OUTER JOIN
dbo.ORI_Order_TT AS ord
	ON sttt.OrderId = ord.OrderId
		AND ord.IsBuyBack = 1
LEFT OUTER JOIN
TRSProcessBuyBackLog_TH AS pbbl
	ON ord.OrderId = pbbl.OrderId
		AND pbbl.DealId IS NOT NULL
WHERE  tctm.SecAccNo = @pvSecAccNo
	AND smtm.SecurityNo = @pvSecurityNo
	AND sttt.TrxType IN (4, 6, 9, 10)
	AND sttt.TrxStatus = 3
	AND sttt.FaceValue >= 0;
IF @dToday <= @dMaturityDate
BEGIN
  INSERT INTO #tCustBlock (DealId, Nominal)
  SELECT   cb.DealId,
  	SUM(cb.Nominal) AS Nominal
  FROM     #tSecurityTransaction AS t
  INNER JOIN
  dbo.TRSRetailCustBlock_TM AS cb
  	ON t.DealId = cb.DealId
  WHERE    cb.CollateralStatus IN (6)
  GROUP BY cb.DealId;
  INSERT INTO #tSumLinkTrx (DealId, SumNominal)
  SELECT   t.DealId,
  	SUM(stl.Nominal) AS SumNominal
  FROM     #tSecurityTransaction AS t
  INNER JOIN
  SecurityTransactionLink_TR AS stl
  	ON t.DealId = stl.DealIdChild
  INNER JOIN
  SecurityTransaction_TT AS st
  	ON stl.DealIdParent = st.DealId
  INNER JOIN
  TRSRetailCustBlock_TM AS cb
  	ON t.SecId = cb.SecId
  WHERE    cb.CollateralStatus = 7
  GROUP BY t.DealId;
  UPDATE t
  SET    FaceValue = (t.FaceValue - (ISNULL(lt.SumNominal, 0) + ISNULL(cb.Nominal, 0))),
  	Blocked   = ISNULL(cb.Nominal, 0)
  FROM   #tSecurityTransaction AS t
  LEFT OUTER JOIN
  #tSumLinkTrx AS lt
  	ON t.DealId = lt.DealId
  LEFT OUTER JOIN
  #tCustBlock AS cb
  	ON t.DealId = cb.DealId;
  DELETE t1
  FROM   #tSecurityTransaction AS t1
  INNER JOIN
  dbo.SecurityTransaction_TT AS t2
  	ON t1.CIFId = t2.CIFId
  		AND t1.SecId = t2.SecId
  		AND t2.TrxType = 9
  		AND t1.TrxType <> 9
  		AND t1.InsertedDate < t2.InsertedDate
  		AND t2.FaceValue < 0;
  UPDATE tr1
  SET    FaceValue = tr1.FaceValue + tr2.FaceValue
  FROM   #tSecurityTransaction AS tr1
  INNER JOIN
  dbo.SecurityTransaction_TT AS tr2
  	ON tr2.CIFId = tr1.CIFId
  		AND tr2.SecId = tr1.SecId
  		AND tr1.DealId = tr2.FOPOriDealId
  		AND tr2.TrxType IN (9, 10)
  WHERE  tr2.FaceValue < 0;
  ALTER TABLE #tSecurityTransaction DROP COLUMN CIFId;
  ALTER TABLE #tSecurityTransaction DROP COLUMN SecId;
  ALTER TABLE #tSecurityTransaction DROP COLUMN InsertedDate;
END
IF @dToday > @dMaturityDate
BEGIN
  UPDATE #tSecurityTransaction
  SET    FaceValue = 0.0;
END
UPDATE #tSecurityTransaction
SET    FaceValue = 0
WHERE  FaceValue < 0;
SELECT   SelectBit,
	SecAccNo,
	isTaxAmnesty,
	DealId,
	OutStandindBlokir,
	TrxDate,
	SettlementDate,
	DealPrice,
	FaceValue,
	Blocked,
	CIFNo,
	TrxType,
	AcquitionDate
FROM     #tSecurityTransaction
ORDER BY CONVERT (VARCHAR (8), SettlementDate, 112), TrxDate ASC;
RETURN 0;
ERR_HANDLER:
IF @@Trancount > 0
  ROLLBACK;
RAISERROR (@message, 16, 1);
RETURN 1;
