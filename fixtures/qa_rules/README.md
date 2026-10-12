# QA Fixtures - Engine Rules

Direktori ini berisi scaffold fixture untuk validasi QA-01 s.d. QA-26 sesuai `docs/qa/qa_gate_engine_rules.md`.
Catatan: QA-17B (Patch Immutability) diuji terpisah oleh `tests/engine/patch_immutability_test.go`.

| QA | Scenario Folder | Expected Outcome |
| --- | --- | --- |
| QA-01 | `01_hard_fail_rule` | fail |
| QA-02 | `02_folder_contract` | fail |
| QA-03 | `03_forbidden_sql` | fail |
| QA-04 | `04_pairing_missing_down` | fail |
| QA-05 | `05_version_format_invalid` | fail |
| QA-06 | `06_ordering_determinism` | pass |
| QA-07 | `07_vk_whitespace_only` | pass |
| QA-08 | `08_reconstruction_unsupported_sql` | fail |
| QA-09 | `09_file_object_mismatch` | fail |
| QA-10 | `10_inline_default_not_allowed` | fail |
| QA-11 | `11_destructive_drop_table_manual` | pass-with-warn |
| QA-12 | `12_transaction_error_throw` | fail |
| QA-13 | `13_audit_wrapper_contract` | pass |
| QA-14 | `14_down_data_restore_warning` | pass-with-warn |
| QA-15 | `15_patch_canonicalization` | pass |
| QA-16 | `16_patch_duplication_once` | pass-skip |
| QA-17 | `17_patch_source_integrity` | fail |
| QA-18 | `18_output_json_contract` | pass |
| QA-19 | `19_repo_hygiene_constraints` | pass |
| QA-20 | `20_cli_exit_code_contract` | pass-fail-matrix |
| QA-21 | `21_wrapper_action_whitelist` | fail |
| QA-22 | `22_wrapper_reason_whitelist` | fail |
| QA-23 | `23_wrapper_run_key_newid` | fail |
| QA-24 | `24_constraint_unnamed` | fail |
| QA-25 | `25_wrapper_precheck_skip` | pass |
| QA-26 | `26_alter_column_precheck_map` | pass |
