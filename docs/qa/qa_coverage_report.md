# QA Coverage Report

Generated at (UTC): 2026-10-12 00:57:07

## Summary

- Total QA scenarios: 26
- Expected pass: 10
- Expected pass-with-warn: 2
- Expected fail: 14
- Total mapped rules in traceability matrix: 26

## QA Scenario Coverage

| QA | Scenario | Command | Expected | Expected Fail Rules | Expected Warn Reasons |
| --- | --- | --- | --- | --- | --- |
| QA-01 | `01_hard_fail_rule` | `verify` | fail | FORBIDDEN_SQL_TOKEN | - |
| QA-02 | `02_folder_contract` | `verify` | fail | MIGRATIONS_NOT_FLAT | - |
| QA-03 | `03_forbidden_sql` | `verify` | fail | FORBIDDEN_SQL_TOKEN | - |
| QA-04 | `04_pairing_missing_down` | `verify` | fail | MISSING_MIGRATION_PAIR | - |
| QA-05 | `05_version_format_invalid` | `verify` | fail | INVALID_VERSION_FORMAT | - |
| QA-06 | `06_ordering_determinism` | `gen` | pass | - | - |
| QA-07 | `07_vk_whitespace_only` | `gen` | pass | - | - |
| QA-08 | `08_reconstruction_unsupported_sql` | `gen` | fail | UNSUPPORTED_SQL_SUBSET | - |
| QA-09 | `09_file_object_mismatch` | `verify` | fail | FILE_OBJECT_MISMATCH | - |
| QA-10 | `10_inline_default_not_allowed` | `verify` | fail | INLINE_DEFAULT_NOT_ALLOWED | - |
| QA-11 | `11_destructive_drop_table_manual` | `gen` | pass-with-warn | - | DESTRUCTIVE_CHANGE |
| QA-12 | `12_transaction_error_throw` | `verify` | fail | WRAPPER_THROW_REQUIRED | - |
| QA-13 | `13_audit_wrapper_contract` | `verify` | pass | - | - |
| QA-14 | `14_down_data_restore_warning` | `gen` | pass-with-warn | - | DATA_RESTORE_NOT_GUARANTEED |
| QA-15 | `15_patch_canonicalization` | `gen` | pass | - | - |
| QA-16 | `16_patch_duplication_once` | `gen` | pass | - | - |
| QA-17 | `17_patch_source_integrity` | `verify` | fail | PATCH_SOURCE_MISSING | - |
| QA-18 | `18_output_json_contract` | `verify` | pass | - | - |
| QA-19 | `19_repo_hygiene_constraints` | `verify` | pass | - | - |
| QA-20 | `20_cli_exit_code_contract` | `fmt|verify|gen` | pass | - | - |
| QA-21 | `21_wrapper_action_whitelist` | `verify` | fail | WRAPPER_AUDIT_ACTION_INVALID | - |
| QA-22 | `22_wrapper_reason_whitelist` | `verify` | fail | WRAPPER_AUDIT_REASON_INVALID | - |
| QA-23 | `23_wrapper_run_key_newid` | `verify` | fail | WRAPPER_RUN_KEY_INVALID | - |
| QA-24 | `24_constraint_unnamed` | `verify` | fail | UNNAMED_CONSTRAINT_NOT_ALLOWED | - |
| QA-25 | `25_wrapper_precheck_skip` | `gen` | pass | - | - |
| QA-26 | `26_alter_column_precheck_map` | `gen` | pass | - | - |

## Rule Mapping Coverage

| Rule No | Rule Section | Control/Gate | Status |
| --- | --- | --- | --- |
| 1 | Istilah & Konsep | Definisi konsisten di semua dokumen | Mapped |
| 2 | Foldering Repo | Verify folder contract | Mapped |
| 3 | Tooling & Pipeline | Command flow contract | Mapped |
| 4 | Larangan & Safety Guard | Forbidden token checks | Mapped |
| 5 | Pairing & Filename | Pair-safe validation | Mapped |
| 6 | Ordering, `v{k}`, Reconstruction | Deterministic emit gate | Mapped |
| 7 | Data Master Validation | Source schema contract checks | Mapped |
| 8 | Destructive & Decision Semantics | Policy mapping checks | Mapped |
| 9 | Transaction & Error Strategy | Wrapper + THROW checks | Mapped |
| 10 | Core Audit Init | Audit init pair checks | Mapped |
| 11 | Audit Logging | Logger contract checks | Mapped |
| 12 | Wrapper Standar | Wrapper minimum validation | Mapped |
| 13 | Rules Down Migration | Down behavior checks | Mapped |
| 14 | Patch Policy | Patch lifecycle + immutability hash checks + cross-platform case-insensitive actual-entry resolution | Mapped |
| 15 | Output `out/output.json` | Output schema checks + complete console failure diagnostics | Mapped |
| 16 | Locking 1 Run per DB | Lock readiness check | Mapped |
| 17 | MVP Supported Subset | Unsupported feature fail-fast | Mapped |
| 18 | ALTER COLUMN Precheck | Precheck templates + parameterized/bracketed-type parser/hash/diff regression (`TestColumnTypeParametersWithWhitespaceRemainSemantic`, `TestRiskClassifiesBracketedDecimalScaleChangeAsGuarded`, `TestBracketedAndBareBuiltInTypesAreSemanticallyEquivalent`, `TestGenDetectsParameterizedColumnChangesWithWhitespace`) | Mapped |
| 19 | constraints_naming.md | Optional guideline + deterministic unnamed constraint safe-fix | Mapped |
| 20 | Repo Hygiene | Optional warning + bootstrap seed tanpa overwrite | Mapped |
| 21 | CLI Contract | Command and exit checks | Mapped |
| 22 | Engine Config & Verify Rules | Config + verify gates + modular package contract | Mapped |
| 23 | Final Locks / Gaps Closed | Hard lock governance | Mapped |
| 24 | Fail/Pass Output Patterns | Output pattern checks | Mapped |
| 25 | Runner Cheatsheet | Runner behavior assumptions | Mapped |
| 26 | Migration History Reconstruction & Report (`rza_migrate_report`) | QA-30/QA-31; shared strict filename parser, source-positioned token access scanner (`internal/sqlparse`) with multiline identity, statement-scoped CTE/alias resolution and semantic client-result classification, excluded canonical core-audit bootstrap with explicit analyzed/excluded counters, ordered dense multi-item table reducer, wrapper-safe exactly-once `CREATE TABLE` scanner, uniqueness enforcement-vs-signal inventory, declared-FK/observed-equality graph with exact columns, Indonesian glossary, single-file connected Relational Maps bounded at 20 objects/24 lines/35 mappings/9.000 Mermaid characters without artificial Hub, node key/index context, object index/shared navigation, exact mapping edge labels/full pair+mapping coverage, separate self-join review/no raw HTML outside Mermaid code fences, separate module lineage/dynamic review with case-insensitive dependency dedup and static return-assignment recognition, catalog-backed procedure profiles with URL-safe links and source-preserved guards, top-down branch-aware local execution flow (transaction scope, conditional routing/rejoin, child calls, error transfer, Mermaid-rendered multi-line guards, diagram legend), separate local object inventory, full source sequence, exact write mapping, complete split recursive maps, cycle-safe recursive impact, dynamic resolution codes, repo-relative output resolution, concrete JSON artifact manifest, no same-column/operator/alias-scope/runtime-target inference, output-path immutability guard, unit/CLI/real-repo tests di `internal/sqlparse`, `internal/migreport`, dan `cmd/rza_migrate_report` | Mapped |

## Residual Risk Checklist

- Tidak ada residual risk yang terdeteksi dari matrix.

## Operational Commands

1. `go test ./... -count=1`
2. `powershell -ExecutionPolicy Bypass -File scripts/validate_qa_fixture_contract.ps1`
3. `powershell -ExecutionPolicy Bypass -File scripts/verify_rules_2x_learning.ps1`
4. `powershell -ExecutionPolicy Bypass -File scripts/validate_learning_examples_hygiene.ps1`
5. `powershell -ExecutionPolicy Bypass -File scripts/validate_migration_lock_contract.ps1`
6. `powershell -ExecutionPolicy Bypass -File scripts/validate_runner_stop_behavior_contract.ps1`
7. `powershell -ExecutionPolicy Bypass -File scripts/generate_rule_coverage_report.ps1`
