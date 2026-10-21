-- READ-ONLY follow-up for DBA. Not executed during this offline audit.
-- Run in the intended database; repeat for SQL_REPLICATE.
-- Permissions/metadata visibility can limit results; missing rows do not prove absence.
SELECT @@VERSION AS sql_version, DB_NAME() AS database_name,
       d.compatibility_level, d.collation_name
FROM sys.databases AS d WHERE d.database_id = DB_ID();

SELECT s.name AS schema_name, t.name AS table_name,
       c.column_id, c.name AS column_name, ty.name AS type_name,
       c.max_length, c.precision, c.scale, c.collation_name,
       c.is_nullable, c.is_computed
FROM sys.tables AS t
JOIN sys.schemas AS s ON s.schema_id = t.schema_id
JOIN sys.columns AS c ON c.object_id = t.object_id
JOIN sys.types AS ty ON ty.user_type_id = c.user_type_id
WHERE t.name IN ('CFMAST','ProCIFRiskProfile_TM','ProCIFChangeMaster_TH',
 'SecurityMaster_TM','SecurityTransaction_TT','SecurityTransactionDetail_TT',
 'ORI_Order_TT','TRSECParamSpread_TR','TRSParameterSettlementDate_TR',
 'TRSParamHargaModal_TR','PTOSCustomer_TM','PTOSOutstanding_TR',
 'PTOSTransaksi_TR','SecurityCouponDate_TR','TRSCustAccNumber_TR','TRSRetailCustBlock_TM')
ORDER BY s.name, t.name, c.column_id;

SELECT s.name AS schema_name, t.name AS table_name, i.name AS index_name,
       i.type_desc, i.is_unique, i.is_disabled, i.is_hypothetical,
       i.has_filter, i.filter_definition,
       ic.key_ordinal, ic.is_included_column, ic.is_descending_key,
       c.name AS column_name
FROM sys.tables AS t
JOIN sys.schemas AS s ON s.schema_id = t.schema_id
JOIN sys.indexes AS i ON i.object_id = t.object_id
JOIN sys.index_columns AS ic ON ic.object_id = i.object_id AND ic.index_id = i.index_id
JOIN sys.columns AS c ON c.object_id = ic.object_id AND c.column_id = ic.column_id
WHERE i.index_id > 0
ORDER BY s.name, t.name, i.index_id, ic.is_included_column, ic.key_ordinal, ic.index_column_id;

-- DMV reads need version-appropriate VIEW DATABASE/VIEW SERVER permissions.
SELECT OBJECT_SCHEMA_NAME(ps.object_id) AS schema_name,
       OBJECT_NAME(ps.object_id) AS table_name, ps.index_id,
       SUM(ps.row_count) AS row_count, SUM(ps.reserved_page_count)*8 AS reserved_kb
FROM sys.dm_db_partition_stats AS ps
GROUP BY ps.object_id, ps.index_id
ORDER BY reserved_kb DESC;

SELECT OBJECT_SCHEMA_NAME(us.object_id) AS schema_name,
       OBJECT_NAME(us.object_id) AS table_name, i.name,
       us.user_seeks, us.user_scans, us.user_lookups, us.user_updates,
       us.last_user_seek, us.last_user_scan
FROM sys.dm_db_index_usage_stats AS us
JOIN sys.indexes AS i ON i.object_id = us.object_id AND i.index_id = us.index_id
WHERE us.database_id = DB_ID()
ORDER BY us.user_updates DESC;

-- Cache evidence is transient; this is not actual row-count/runtime plan evidence.
-- CONVERT_IMPLICIT can be harmless; inspect which operand and which operator.
SELECT TOP (50) qs.execution_count, qs.total_logical_reads, qs.total_worker_time,
       qs.total_elapsed_time, qs.plan_handle, st.text, qp.query_plan
FROM sys.dm_exec_query_stats AS qs
CROSS APPLY sys.dm_exec_sql_text(qs.sql_handle) AS st
CROSS APPLY sys.dm_exec_query_plan(qs.plan_handle) AS qp
WHERE st.dbid = DB_ID()
  AND CONVERT(nvarchar(max), qp.query_plan) LIKE N'%CONVERT_IMPLICIT%'
ORDER BY qs.total_logical_reads DESC;

-- For representative approved executions collect actual plan plus STATISTICS IO/TIME,
-- before/after one rewrite or index at a time. No business procedure is executed here.
