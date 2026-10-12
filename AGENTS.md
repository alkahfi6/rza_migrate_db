# Repository Guidelines

## Project Structure & Module Organization

This Go repository contains SQL Server extraction, authoring, migration, and service tools. The primary `db_object` entrypoints are `cmd/rza_get_object_db`, `cmd/rza_migrate_check`, and `cmd/rza_migrate_db`; REST/service and test entrypoints also live under `cmd/`. Core packages are in `internal/`, with engine, verification, REST migration, SQL parsing, formatting, output, and filesystem helpers split by domain. SQL desired-state files live in `db_object/types`, `db_object/tables`, `db_object/indexes`, `db_object/views`, `db_object/functions`, and `db_object/procedures`. Generated migration pairs are stored flat in `migrations/`. Test fixtures and golden contracts are in `fixtures/` and `tests/`. Project rules and design references are in `rules.md`, `spec.md`, `fsd.md`, `plan.md`, and `docs/`.

## DB Object Cross-Tool Linearity (Mandatory)

Treat every feature that reads, writes, normalizes, validates, or interprets `db_object` as one cross-tool contract. A change is incomplete if it works only in source code or in one executable while another relevant path retains different behavior.

The components have different responsibilities, but their understanding of a database object must remain linear:

| Component | Required responsibility |
| --- | --- |
| `cmd/rza_get_object_db` | Extract supported SQL Server objects and emit canonical `db_object` source. It must not generate migrations. |
| `cmd/rza_migrate_check` | Apply proven-safe object corrections in interactive/double-click mode, then validate. It must never generate or modify migrations. With `-check`, it must detect the same pending corrections without writing and return a blocking failure. |
| `cmd/rza_migrate_db` | Use the same object parsing, canonicalization, and safe-correction rules before verification and migration generation. It alone owns the authoring flow that generates migrations. |
| `tools/linux-amd64/*` | Carry the same compiled behavior into the MR pipeline. These binaries are release artifacts, not independent implementations. |
| `build_all.bat` | Build and publish the complete Windows/Linux artifact set, license, checksums, build identity, and release manifest from one source revision. Partial or mixed-version publication is forbidden. |

The following behavior must be shared wherever applicable: supported object categories, filename/schema identity, parsing and AST semantics, encoding normalization, removal of extractor-only catalog wrappers, canonical SQL form, constraint interpretation, safe auto-corrections, and ambiguity/failure policy. Implement such behavior in a shared `internal/` package or shared engine path; do not copy slightly different logic into individual commands.

For an auto-correction feature, all of these outcomes are required:

1. The extractor emits the corrected/canonical representation for newly extracted objects whenever it has enough metadata to do so.
2. Interactive `rza_migrate_check` and the formatting/generation paths of `rza_migrate_db` apply the same safe correction to existing files.
3. `rza_migrate_check -check` and the Linux MR artifact plan the same correction read-only, do not modify the checkout, and block until the user runs the interactive tool.
4. Real semantic changes remain visible to migration generation. An auto-correction must not erase, reinterpret, or hide a business/schema delta.
5. Ambiguous or business-logic-sensitive transformations remain manual and produce an actionable report instead of a guessed rewrite.

Constraint-name reconciliation is an example: reuse a historical name only when type, columns and order, clustering, expression, references, and other semantic attributes are identical. Encoding normalization is another example: extractor output and safe authoring fixes must converge to UTF-8 without BOM, while CI only reports a required correction.

Do not force identical actions where component responsibilities differ. In particular:

- `rza_migrate_check` may correct `db_object` but must never create migration files.
- `rza_get_object_db` may refresh `db_object` but must never mutate migration history.
- CI must be read-only for authored SQL and must fail rather than silently repair the MR checkout.
- Existing migration history is immutable except through an explicit, user-authorized history repair workflow.

When adding a new object kind or object rule, audit every row in the table above. Document any intentionally non-applicable component and add a test proving the boundary. Silence or omission is not an acceptable exception.

## Cross-Tool Definition of Done

Before considering a `db_object` feature complete, an agent must:

1. Trace extraction, interactive check, CI check, format/verify/generate, and release-artifact call paths.
2. Add focused unit tests for the shared parser/fixer and integration tests for both apply and read-only plan modes.
3. Prove that checker execution leaves `migrations/` unchanged.
4. Prove that generator output contains only the genuine object delta and retains valid up/down behavior when migration generation is affected.
5. Run `go test ./...` and the relevant hardening/fixture tests.
6. Run `build_all.bat` after the final source change; never treat source-only tests as sufficient when committed binaries are used by CI.
7. Verify `tools/linux-amd64/release.sha256`, embedded build identity, and source digest, then smoke-test the same Linux checker command used by GitLab.
8. Review for stale artifacts, partial publication, write behavior in `-check`, case-sensitive/case-insensitive identifier edge cases, ambiguous AST matches, and idempotence on a second run.

If any source covered by the release digest changes after building, rebuild all artifacts again. Never update a binary checksum alone to make a pipeline pass.

## Documentation Map and Update Contract (Mandatory)

Start with the [Engine Docs Guide](docs/engine_docs_guide.md). It defines the documentation hierarchy, reading order, fixture references, and final validation workflow. Do not implement a feature by reading code alone.

Use this repository documentation map:

| Document | What it owns | Update requirement |
| --- | --- | --- |
| [README.md](README.md) | User-facing behavior, daily workflow, CLI examples, supported folders, and operational expectations | Update for every user-visible feature or command/flag behavior change. |
| [rules.md](rules.md) | Normative rules and hard locks | Update when adding, removing, relaxing, or strengthening a validation/safety rule. Do not silently change code behavior against this document. |
| [spec.md](spec.md) | Requirement baseline and acceptance criteria | Update for every new engine capability or changed requirement. |
| [fsd.md](fsd.md) | Technical design, module boundaries, data flow, parser/fixer/generator behavior | Update when implementation flow, shared packages, AST behavior, report fields, or component responsibilities change. |
| [plan.md](plan.md) | Delivery phases, roadmap, readiness gates, and completion status | Update when the feature changes scope, milestone status, residual work, or release gate. |
| [docs/walkthrough.md](docs/walkthrough.md) | Actual end-to-end flow across extractor, checker, generator, and runner | Update for every change that users or operators will observe in the cross-tool workflow. |
| [docs/desired_state_lifecycle.md](docs/desired_state_lifecycle.md) | Contract between current `db_object`, immutable migration history, deletion, and reintroduction | Update when desired-state/history interpretation changes. |
| [docs/engine_guarded_migrations.md](docs/engine_guarded_migrations.md) | Automatic/guarded/manual migration behavior and up/down examples | Update when diff planning, migration SQL, reversibility, or manual fallback changes. |
| [docs/qa/sql_server_capability_matrix.md](docs/qa/sql_server_capability_matrix.md) | Supported, guarded, manual-only, fail-hard, and skipped SQL Server features | Update for every added object feature or changed support level. A feature is not “supported” until all relevant cross-tool paths and tests pass. |
| [docs/traceability/rules_traceability_matrix.md](docs/traceability/rules_traceability_matrix.md) | Mapping from rules and requirements to implementation controls/tests | Update whenever a rule, report code, control, or QA coverage mapping changes. |
| [docs/qa/qa_gate_engine_rules.md](docs/qa/qa_gate_engine_rules.md) | QA scenarios and release checks for engine rules | Add or update scenarios for every behavior change, including positive, negative, and no-write cases. |
| [fixtures/qa_rules/README.md](fixtures/qa_rules/README.md) and `fixtures/qa_rules/<scenario>/` | Executable QA fixtures, expectations, and golden reports | Add/update when the feature creates a new rule, output item, migration shape, edge case, or acceptance scenario. |
| [docs/qa/qa_coverage_report.md](docs/qa/qa_coverage_report.md) | Generated QA coverage and residual-risk view | Regenerate with `scripts/generate_rule_coverage_report.ps1`; do not hand-edit generated facts. |
| [docs/qa/learning_edge_case_backlog.md](docs/qa/learning_edge_case_backlog.md) | Known unsupported or not-yet-covered edge cases | Update when discovering, resolving, accepting, or deferring an edge case. |
| [docs/bootstrap_four_repos.md](docs/bootstrap_four_repos.md) | Bootstrap workflow for the four Bond/Reksa repository/catalog combinations | Update when extraction, initial migration, catalog handling, or bootstrap behavior changes. |
| [constraints_naming.md](constraints_naming.md) | Optional constraint naming guidance | Update when constraint naming, unnamed-constraint correction, or history-name reconciliation changes; it must remain optional unless `rules.md` explicitly changes that policy. |

Review these conditional documents when the affected area matches:

| Area changed | Documents to review/update |
| --- | --- |
| Packaging, Linux tools, checksum, license, MR gate, or release identity | [production readiness](docs/policy/production_readiness.md), [engine docs guide](docs/engine_docs_guide.md), [.gitlab checker job](.gitlab/rza_migrate_check.gitlab-ci.yml), and relevant MR templates under `.gitlab/merge_request_templates/` |
| Extractor or initial four-repository workflow | [bootstrap four repos](docs/bootstrap_four_repos.md), [walkthrough](docs/walkthrough.md), capability matrix, and `out/get_object_summary.json` contract documentation |
| Migration execution/runtime behavior | [migration monitoring runbook](docs/operations/migration_monitoring.md), [production readiness](docs/policy/production_readiness.md), and migration job templates under `.gitlab/migration_jobs/` |
| REST migration service behavior | [REST API hardening policy](docs/policy/rest_api_hardening.md), [production readiness](docs/policy/production_readiness.md), README, spec, and FSD |
| Parallelism, caching, ordering, or performance | [parallel-safe design](docs/policy/parallel_safe_design.md), production readiness, FSD, plan, and benchmark/repeatability documentation |
| Rebaseline or historical migration policy | [rebaseline playbook](docs/policy/rebaseline_playbook.md), desired-state lifecycle, guarded migrations, rules, spec, and FSD |
| Legacy wrapper/audit behavior | [legacy wrapper cutover policy](docs/policy/legacy_wrapper_cutover.md), rules, traceability matrix, and QA gate |

For every new `db_object` feature, the minimum documentation set is:

1. `README.md` for the user-visible behavior.
2. `spec.md` for requirements and acceptance criteria.
3. `fsd.md` for the shared cross-tool implementation and boundaries.
4. `docs/walkthrough.md` for extractor → checker → generator → pipeline behavior.
5. `docs/qa/sql_server_capability_matrix.md` for the support classification.
6. `docs/traceability/rules_traceability_matrix.md` for rule-to-control/test mapping.
7. `docs/qa/qa_gate_engine_rules.md` plus a focused fixture/test for proof.
8. `rules.md`, `plan.md`, lifecycle, guarded-migration, and conditional policy/runbook documents whenever their owned contract is affected.

Documentation must explicitly state the behavior of each relevant component: extractor output, interactive checker apply mode, checker `-check` plan mode, migration generator impact, pipeline behavior, and released binary availability. “Implemented in the engine” is not sufficient documentation.

Before finishing, report which documents were updated and which were reviewed but intentionally unchanged, with a short reason. A behavior change with stale or contradictory documentation is incomplete even when tests pass.

## Build, Test, and Development Commands

- `go test ./...`: runs all unit, integration, and fixture contract tests.
- `go run ./cmd/rza_migrate_db verify -repo .`: validates folder contracts, migration pairs, wrappers, and SQL rules.
- `go run ./cmd/rza_migrate_db gen -repo .`: reconstructs history and emits deterministic migration files.
- `go run ./cmd/rza_migrate_check -repo .`: applies safe `db_object` corrections and validates without generating migrations.
- `go run ./cmd/rza_migrate_check -repo . -check`: performs the pipeline-safe, read-only form of the same authoring checks.
- `go run ./cmd/rza_get_object_db -dry-run`: checks extractor deltas without writing snapshots.
- `build_all.bat <license_owner> <license_months>`: atomically rebuilds the complete Windows/Linux release and license.
- `powershell -ExecutionPolicy Bypass -File scripts/run_final_hardening_gate.ps1`: runs the full local hardening gate.
- `powershell -ExecutionPolicy Bypass -File scripts/validate_parallel_repeatability.ps1`: checks deterministic parallel behavior.

## Coding Style & Naming Conventions

Use standard Go formatting (`gofmt`) and keep packages focused by domain. Test files use Go's `_test.go` suffix and descriptive `Test...` names. SQL object files should follow the repository pattern `dbo.ObjectName.sql` inside the correct object-type directory. Migration files must keep paired `.up.sql` and `.down.sql` files with deterministic timestamp prefixes and object-type names, for example `20260217073015000_tables_dbo.TableA_v1.up.sql`.

## Testing Guidelines

Add targeted tests beside changed Go packages, and update fixture goldens only when behavior intentionally changes. Fixture scenarios in `fixtures/qa_rules/<scenario>/` must include `README.md`, `input/`, `expected/expectation.json`, and `expected/output.golden.json`. Every safe fixer needs idempotence coverage, an interactive write test, a CI no-write test, and a negative ambiguity/semantic-change test. Run `go test ./...` before submitting changes; use the PowerShell gates for migration-rule or parallel-read changes.

## Commit & Pull Request Guidelines

This checkout does not include Git history, so use concise, imperative commit messages such as `verify wrapper action whitelist` or `fix migration filename parser`. Pull requests should explain the rule or workflow affected, list validation commands run, link related issues, and call out changes to generated migrations, fixtures, or security-sensitive REST settings.

## Security & Configuration Tips

Do not commit secrets, SQL Server credentials, API keys, or local `out/` artifacts. Extractor configuration is currently compile-time in `cmd/rza_get_object_db/config.go`; review connection strings and repository paths before building distributable binaries.
