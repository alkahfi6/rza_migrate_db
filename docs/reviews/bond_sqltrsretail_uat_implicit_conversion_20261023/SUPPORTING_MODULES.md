# Function dan view pendukung

Temuan pada body function/view dicatat terpisah dari lokasi di procedure. Pemanggilan UDF langsung dengan argumen kolom ada dalam daftar procedure; pemanggilan transitif tidak dianggap otomatis mengeksekusi setiap branch function. Semua detail tetap tersedia di conversion_sites.csv.

| Database / module | Lokasi beda tipe | Context / form |
|---|---:|---|
| retail / [dbo.CalculateSafeKeepingAfterTaxSettBuyBack](../../../bond-sqltrsretail-uat/db_object/functions/dbo.CalculateSafeKeepingAfterTaxSettBuyBack.sql) | 1 | COMPARISON |
| retail / [dbo.ProcessTableSelect_FN](../../../bond-sqltrsretail-uat/db_object/functions/dbo.ProcessTableSelect_FN.sql) | 3 | COMPARISON, UDF_ARGUMENT |
| retail / [dbo.SIBSProcessCekReady_FN](../../../bond-sqltrsretail-uat/db_object/functions/dbo.SIBSProcessCekReady_FN.sql) | 15 | COMPARISON, ISNULL_REPLACEMENT, SELECTSETVARIABLE |
| retail / [dbo.fnCalculateAccruedInterest](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnCalculateAccruedInterest.sql) | 1 | SELECTSETVARIABLE |
| retail / [dbo.fnCalculateAccruedInterestBUYBACK](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnCalculateAccruedInterestBUYBACK.sql) | 1 | SELECTSETVARIABLE |
| retail / [dbo.fnCheckRecordingDate2](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnCheckRecordingDate2.sql) | 2 | COMPARISON |
| retail / [dbo.fnGetBranchCode](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnGetBranchCode.sql) | 1 | COMPARISON |
| retail / [dbo.fnGetCurrentDate](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnGetCurrentDate.sql) | 1 | SELECTSETVARIABLE |
| retail / [dbo.fnGetEmailUser](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnGetEmailUser.sql) | 1 | COMPARISON |
| retail / [dbo.fnGetMxWorkingDate](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnGetMxWorkingDate.sql) | 1 | COMPARISON |
| retail / [dbo.fnGetMxWorkingDateValasIncUSD](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnGetMxWorkingDateValasIncUSD.sql) | 2 | IN_VALUE_TYPE |
| retail / [dbo.fnGetNameByNIK](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnGetNameByNIK.sql) | 1 | COMPARISON |
| retail / [dbo.fnGetNameEmployee](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnGetNameEmployee.sql) | 1 | COMPARISON |
| retail / [dbo.fnGetRecordingDate](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnGetRecordingDate.sql) | 2 | ASSIGNMENTSETCLAUSE, INSERT_SELECT |
| retail / [dbo.fnGetRiskProfilSecurity](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnGetRiskProfilSecurity.sql) | 1 | ISNULL_REPLACEMENT |
| retail / [dbo.fnGetSettlementDate](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnGetSettlementDate.sql) | 7 | COMPARISON, UDF_ARGUMENT |
| retail / [dbo.fnGetTerbilang](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnGetTerbilang.sql) | 2 | COMPARISON |
| retail / [dbo.fnGetTerbilangCurrency](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnGetTerbilangCurrency.sql) | 2 | COMPARISON |
| retail / [dbo.fnGetTerbilangEnglish](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnGetTerbilangEnglish.sql) | 2 | COMPARISON |
| retail / [dbo.fnGetTrxnAcctStatus](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnGetTrxnAcctStatus.sql) | 1 | COMPARISON |
| retail / [dbo.fnMoveDate](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnMoveDate.sql) | 1 | COMPARISON |
| retail / [dbo.fnMoveDatePerCCY](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnMoveDatePerCCY.sql) | 2 | COMPARISON |
| retail / [dbo.fnOBLGetCurrentRiskProfile](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnOBLGetCurrentRiskProfile.sql) | 1 | COMPARISON |
| retail / [dbo.fnOBLGetORIOrderStatus](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnOBLGetORIOrderStatus.sql) | 1 | COMPARISON |
| retail / [dbo.fnOBLIsHoliday](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnOBLIsHoliday.sql) | 2 | COMPARISON, IN_VALUE_TYPE |
| retail / [dbo.fnOBLIsRiskProfileExpired](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnOBLIsRiskProfileExpired.sql) | 1 | COMPARISON |
| retail / [dbo.fnTrsGenerateDealNumber](../../../bond-sqltrsretail-uat/db_object/functions/dbo.fnTrsGenerateDealNumber.sql) | 2 | SELECTSETVARIABLE |
| retail / [dbo.select_process_table_fn](../../../bond-sqltrsretail-uat/db_object/functions/dbo.select_process_table_fn.sql) | 8 | COMPARISON, UDF_ARGUMENT |
| retail / [dbo.sibs_select_process_table_fn](../../../bond-sqltrsretail-uat/db_object/functions/dbo.sibs_select_process_table_fn.sql) | 3 | COMPARISON, UDF_ARGUMENT |
| retail / [dbo.AccountSegment_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.AccountSegment_v.sql) | 1 | IN_VALUE_TYPE |
| retail / [dbo.BIS4Instrument_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.BIS4Instrument_v.sql) | 4 | LIKE |
| retail / [dbo.BIS4Participant_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.BIS4Participant_v.sql) | 4 | UNION_OUTPUT |
| retail / [dbo.CDMASTALL_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.CDMASTALL_v.sql) | 2 | COMPARISON |
| retail / [dbo.DDHISTALL_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.DDHISTALL_v.sql) | 3 | UNION_OUTPUT |
| retail / [dbo.OMTreasuryCustomerBalance_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.OMTreasuryCustomerBalance_v.sql) | 25 | COMPARISON, IN_VALUE_TYPE, ISNULL_REPLACEMENT, SIMPLE_CASE_MATCH |
| retail / [dbo.PSEmailAddress_REP_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.PSEmailAddress_REP_v.sql) | 3 | COMPARISON, UNION_OUTPUT |
| retail / [dbo.RPATreasuryCustomerSaldo_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.RPATreasuryCustomerSaldo_v.sql) | 2 | BINARY_Divide, BINARY_Multiply |
| retail / [dbo.RepoHargaModal_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.RepoHargaModal_v.sql) | 8 | COMPARISON, ISNULL_REPLACEMENT, SIMPLE_CASE_MATCH |
| retail / [dbo.RepoTenor_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.RepoTenor_v.sql) | 11 | COMPARISON, ISNULL_REPLACEMENT, SIMPLE_CASE_MATCH |
| retail / [dbo.TELGRP_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.TELGRP_v.sql) | 1 | CASE_BRANCH |
| retail / [dbo.TRSOutstandingPerAccount_TM_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.TRSOutstandingPerAccount_TM_v.sql) | 2 | BINARY_Multiply, COMPARISON |
| retail / [dbo.TRSRetailTrx_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.TRSRetailTrx_v.sql) | 52 | COMPARISON, IN_VALUE_TYPE, ISNULL_REPLACEMENT, SIMPLE_CASE_MATCH |
| retail / [dbo.TreasuryCustomerBalanceExcludePTBC_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.TreasuryCustomerBalanceExcludePTBC_v.sql) | 27 | COMPARISON, IN_VALUE_TYPE, ISNULL_REPLACEMENT, SIMPLE_CASE_MATCH |
| retail / [dbo.TreasuryCustomerBalancePTBCOnly_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.TreasuryCustomerBalancePTBCOnly_v.sql) | 27 | COMPARISON, IN_VALUE_TYPE, ISNULL_REPLACEMENT, SIMPLE_CASE_MATCH |
| retail / [dbo.TreasuryCustomerBalanceRepEOD_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.TreasuryCustomerBalanceRepEOD_v.sql) | 41 | COMPARISON, IN_VALUE_TYPE, ISNULL_REPLACEMENT, SIMPLE_CASE_MATCH |
| retail / [dbo.TreasuryCustomerBalanceTaxAmnesty_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.TreasuryCustomerBalanceTaxAmnesty_v.sql) | 32 | BINARY_Multiply, COMPARISON, IN_VALUE_TYPE, ISNULL_REPLACEMENT, SIMPLE_CASE_MATCH |
| retail / [dbo.TreasuryCustomerBalanceUnblock_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.TreasuryCustomerBalanceUnblock_v.sql) | 25 | COMPARISON, IN_VALUE_TYPE, ISNULL_REPLACEMENT, SIMPLE_CASE_MATCH |
| retail / [dbo.TreasuryCustomerBalance_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.TreasuryCustomerBalance_v.sql) | 27 | COMPARISON, IN_VALUE_TYPE, ISNULL_REPLACEMENT, SIMPLE_CASE_MATCH |
| retail / [dbo.TreasuryCustomerBalance_v_forPTBC](../../../bond-sqltrsretail-uat/db_object/views/dbo.TreasuryCustomerBalance_v_forPTBC.sql) | 26 | COMPARISON, IN_VALUE_TYPE, ISNULL_REPLACEMENT, SIMPLE_CASE_MATCH |
| retail / [dbo.TreasuryCustomerSaldoExcludePTBC_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.TreasuryCustomerSaldoExcludePTBC_v.sql) | 25 | COMPARISON, IN_VALUE_TYPE, ISNULL_REPLACEMENT, SIMPLE_CASE_MATCH |
| retail / [dbo.TreasuryCustomerSaldoLite_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.TreasuryCustomerSaldoLite_v.sql) | 21 | IN_VALUE_TYPE, ISNULL_REPLACEMENT, SIMPLE_CASE_MATCH |
| retail / [dbo.TreasuryCustomerSaldoRepEOD_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.TreasuryCustomerSaldoRepEOD_v.sql) | 37 | COMPARISON, IN_VALUE_TYPE, ISNULL_REPLACEMENT, SIMPLE_CASE_MATCH |
| retail / [dbo.TreasuryCustomerSaldoRepo_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.TreasuryCustomerSaldoRepo_v.sql) | 5 | COMPARISON, IN_VALUE_TYPE, ISNULL_REPLACEMENT |
| retail / [dbo.TreasuryCustomerSaldoSyariah_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.TreasuryCustomerSaldoSyariah_v.sql) | 31 | BINARY_Multiply, CASE_BRANCH, COMPARISON, IN_VALUE_TYPE, ISNULL_REPLACEMENT, SIMPLE_CASE_MATCH |
| retail / [dbo.TreasuryCustomerSaldoTaxAmnesty_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.TreasuryCustomerSaldoTaxAmnesty_v.sql) | 31 | BINARY_Multiply, CASE_BRANCH, COMPARISON, IN_VALUE_TYPE, ISNULL_REPLACEMENT, SIMPLE_CASE_MATCH |
| retail / [dbo.TreasuryCustomerSaldo_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.TreasuryCustomerSaldo_v.sql) | 25 | COMPARISON, IN_VALUE_TYPE, ISNULL_REPLACEMENT, SIMPLE_CASE_MATCH |
| retail / [dbo.TreasuryCustomerSaldo_v_PTBCOnly](../../../bond-sqltrsretail-uat/db_object/views/dbo.TreasuryCustomerSaldo_v_PTBCOnly.sql) | 25 | COMPARISON, IN_VALUE_TYPE, ISNULL_REPLACEMENT, SIMPLE_CASE_MATCH |
| retail / [dbo.office_information_sibs_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.office_information_sibs_v.sql) | 1 | IN_VALUE_TYPE |
| retail / [dbo.office_information_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.office_information_v.sql) | 1 | IN_VALUE_TYPE |
| retail / [dbo.user_nisp_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.user_nisp_v.sql) | 5 | COMPARISON, ISNULL_REPLACEMENT, UNION_OUTPUT |
| retail / [dbo.user_table_classification_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.user_table_classification_v.sql) | 3 | COMPARISON, UNION_OUTPUT |
