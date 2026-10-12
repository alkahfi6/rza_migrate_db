CREATE OR ALTER FUNCTION dbo.fnReksaGoodFund (@nProdId int,@dInputTgl datetime,@nTipe int)
returns datetime
as

BEGIN
  declare @dGoodFund int
  	, @dTglAkhir datetime
  	, @dTglAwal datetime
  	, @nNomor int
  	, @nNDayBefore int
  	, @bFixNAVDate bit
  	, @nMonthPeriod tinyint

  declare @temp table
  	(dif int default 1,tgl datetime,nomor int)
  declare @temp1 table
  	(nomor int identity,tglbaru datetime)

  	select @dGoodFund = GoodFundPeriod
  	from ReksaProduct_TM
  	where ProdId=@nProdId

  set @dTglAwal = @dInputTgl

  set @dTglAkhir=dateadd(m,2,@dInputTgl)

  while (@dTglAwal<@dTglAkhir)

  begin
    insert into @temp(tgl) select @dTglAwal
    set @dTglAwal=dateadd(day,1,@dTglAwal)
  end

  delete a from @temp a join ReksaHolidayTable_TM b
  		on (a.tgl = b.ValueDate)

  delete from @temp
  	where DATEPART(dw, tgl) in (1,7)

  insert into @temp1(tglbaru) select tgl from @temp
  if(@nTipe=1)
  begin
    	select @dTglAkhir=tglbaru from @temp1
    	where nomor = (select nomor+0 from @temp1 where tglbaru=@dInputTgl)

  end

  else if(@nTipe=2)

  begin

    set @dTglAkhir = dbo.fnReksaGetEffectiveDate(@dInputTgl, @dGoodFund)

  end

  else if(@nTipe=3)
  begin
    	select @nNDayBefore = NDayBefore, @bFixNAVDate = FixNAVDate, @dTglAkhir = ReferDate, @nMonthPeriod = MonthPeriod
    	from ReksaWindowPeriod_TM
    	where ProdId = @nProdId

    If @bFixNAVDate = 0
    Begin
      	select @dTglAkhir=tglbaru from @temp1
      	where nomor = (select nomor+@nNDayBefore from @temp1 where tglbaru=@dInputTgl)
    End
    else
    Begin
      while @dTglAkhir < @dInputTgl
      Begin
        	select @dTglAkhir = dateadd(mm, @nMonthPeriod, @dTglAkhir)
      End
      while (DATEPART(dw, @dTglAkhir) in (1,7)
        	or exists(select top 1 1 from ReksaHolidayTable_TM where ValueDate = @dTglAkhir))
      	select @dTglAkhir = dateadd(dd, 1, @dTglAkhir)
    End
  end
  else if(@nTipe=4)
  begin
    	select @dTglAkhir=tglbaru from @temp1
    	where nomor = (select nomor+3 from @temp1 where tglbaru=@dInputTgl)
  end

  return CAST(CONVERT(char(8),@dTglAkhir,112) as datetime)
END
