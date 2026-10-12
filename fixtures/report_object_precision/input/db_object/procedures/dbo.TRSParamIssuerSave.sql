CREATE OR ALTER PROCEDURE dbo.TRSParamIssuerSave @pcXML varchar(max)
	, @pnNIK int
AS
SET NOCOUNT ON

DECLARE
	@nOK int
	, @nErrNo int
	, @cErrMsg varchar(1000)
	, @nDocHandle int
	, @nApprId bigint

CREATE TABLE #tmp_Issuer (
	IssuerId			bigint
	,IssuerName			varchar(250)
	,Limit				money

	,[Notification]		decimal(12,2)
	,IssuerCode varchar(10)

)

CREATE TABLE #tmp_Changes (
	IssuerId			bigint
	,IssuerName			varchar(250)
	,Limit				money

	,[Notification]		decimal(12,2)

	,ProcessCode		char(1)
	,ChangeNameBit		bit			default(0)
	,IssuerCode varchar(10)
)

EXEC sp_xml_preparedocument @nDocHandle output, @pcXML

INSERT #tmp_Issuer (
	IssuerId,IssuerName,Limit

	,[Notification]
	,IssuerCode

)
SELECT
	IssuerId,IssuerName,Limit

	,ISNULL([Notification],0)
	,ISNULL(IssuerCode, '') AS IssuerCode
FROM openxml(@nDocHandle, N'/Data/Issuer', 2)
WITH (
IssuerId			bigint
,IssuerName			varchar(100)
,Limit				money

,[Notification]		decimal(12,2)
,IssuerCode varchar(10)

)

EXEC sp_xml_removedocument @nDocHandle

INSERT #tmp_Changes ( IssuerId,IssuerName,Limit,ProcessCode

	,[Notification]
	,IssuerCode

)
SELECT IssuerId,IssuerName,Limit,'I'

	,[Notification]
	,ISNULL(IssuerCode, '') AS IssuerCode
FROM #tmp_Issuer
WHERE IssuerId = 0

INSERT #tmp_Changes ( IssuerId,IssuerName,Limit,ProcessCode

	,[Notification]
	,IssuerCode

)
SELECT par.IssuerId,par.IssuerName,par.Limit,'O'

	,par.[Notification]
	,par.IssuerCode
FROM #tmp_Issuer tmp
JOIN dbo.TRSParamIssuer_TR par
	ON tmp.IssuerId			= par.IssuerId
WHERE tmp.IssuerName		<> par.IssuerName
OR tmp.Limit			<> par.Limit
OR tmp.[Notification]	<> par.[Notification]

INSERT #tmp_Changes ( IssuerId,IssuerName,Limit,ProcessCode,ChangeNameBit

	,[Notification]
	,IssuerCode

)
SELECT tmp.IssuerId,tmp.IssuerName,tmp.Limit,'N',
	case
		when tmp.IssuerName <> par.IssuerName then 1
		else 0
	end

	,tmp.[Notification]
	,tmp.IssuerCode
FROM #tmp_Issuer tmp
JOIN dbo.TRSParamIssuer_TR par
	ON tmp.IssuerId			= par.IssuerId
WHERE tmp.IssuerName		<> par.IssuerName
OR tmp.Limit			<> par.Limit
OR tmp.[Notification]	<> par.[Notification]

INSERT #tmp_Changes ( IssuerId,IssuerName,Limit,ProcessCode

	,tmp.[Notification]
	,tmp.IssuerCode

)
SELECT par.IssuerId,par.IssuerName,par.Limit,'D'

	,par.[Notification]
	,par.IssuerCode
FROM dbo.TRSParamIssuer_TR par
LEFT JOIN #tmp_Issuer tmp
	ON tmp.IssuerId			= par.IssuerId
WHERE tmp.IssuerId		is null

IF EXISTS	(
  SELECT TOP 1 1 FROM #tmp_Changes tmp
	JOIN dbo.SecurityMaster_TM sec
		ON tmp.IssuerName = sec.IssuerName
WHERE tmp.ProcessCode  = 'D'
)
BEGIN
  SET @cErrMsg = 'Tidak bisa delete data yang sudah di tag ke parameter surat berharga !'
  	GOTO ERROR
END

IF EXISTS	(
  SELECT TOP 1 1 FROM #tmp_Changes tmp
	JOIN dbo.SecurityMaster_TM sec
		ON tmp.IssuerName = sec.IssuerName
WHERE tmp.ProcessCode		= 'N'
	AND tmp.ChangeNameBit	= 1
)
BEGIN
  SET @cErrMsg = 'Tidak bisa ganti nama issuer dari data yang sudah di tag ke parameter surat berharga !'
  	GOTO ERROR
END

IF EXISTS (
  SELECT TOP 1 1 FROM #tmp_Changes tmp
	JOIN dbo.TRSParamIssuer_TH hist
		ON tmp.IssuerName  = hist.IssuerName
	JOIN dbo.TRSAuthParamIssuer_TH auth
		ON hist.ApprId = auth.ApprId
WHERE tmp.ProcessCode	= 'I'
	AND auth.AuthStatus	= 'P'
)
BEGIN
  SET @cErrMsg = 'Masih ada pending approval !'
  	GOTO ERROR
END

IF EXISTS (
  SELECT TOP 1 1 FROM #tmp_Changes tmp
	JOIN dbo.TRSParamIssuer_TH hist
		ON tmp.IssuerId  = hist.IssuerId
	JOIN dbo.TRSAuthParamIssuer_TH auth
		ON hist.ApprId = auth.ApprId
WHERE tmp.ProcessCode	in ('N', 'D')
	AND auth.AuthStatus	= 'P'
)
BEGIN
  SET @cErrMsg = 'Masih ada pending approval !'
  	GOTO ERROR
END

BEGIN TRAN
BEGIN TRY
  INSERT dbo.TRSAuthParamIssuer_TH ( IssueDate,IssueNIK,AuthDate,AuthNIK,AuthStatus )
  SELECT getdate(), @pnNIK, null, null, 'P'

  SET @nApprId = @@identity

  INSERT dbo.TRSParamIssuer_TH (
  	ApprId
  	,IssuerId
  	,IssuerName
  	,Limit

  	,[Notification]

  	,ProcessCode
  	,LastModifyDate
  	,LastModifyNIK
  	,IssuerCode
  )
  SELECT
  	@nApprId
  	,IssuerId
  	,IssuerName
  	,Limit

  	,[Notification]

  	,ProcessCode
  	,getdate()
  	,@pnNIK
  	,IssuerCode
  FROM #tmp_Changes

END TRY
BEGIN CATCH
  SET @cErrMsg	= error_message()
  	GOTO ERROR
END CATCH
COMMIT TRAN

IF object_id('tempdb..#tmp_Issuer') is not null
  DROP TABLE #tmp_Issuer

IF object_id('tempdb..#tmp_Changes') is not null
  DROP TABLE #tmp_Changes

RETURN 0

ERROR:
IF object_id('tempdb..#tmp_Issuer') is not null
  DROP TABLE #tmp_Issuer

IF object_id('tempdb..#tmp_Changes') is not null
  DROP TABLE #tmp_Changes

IF @@trancount > 0
  ROLLBACK TRAN

IF @cErrMsg is null
  SET @cErrMsg = 'Unknown error at sp TRSParamIssuerSave !'

IF @@error <> 0
  RETURN 1

RAISERROR (@cErrMsg, 16, 1)
RETURN 1
