CREATE NONCLUSTERED INDEX IX_MXHolidayCurrency_Currency
ON MXHolidayCurrency (Currency);
CREATE CLUSTERED INDEX IX_MXHolidayCurrency_M_HOLIDAY
ON MXHolidayCurrency (M_HOLIDAY); 