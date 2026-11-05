# Extractor YAML configuration review — 2026-11-03

Scope: external credentials for `rza_get_object_db`, not a SQL object or migration-rule change. The working-directory local `config.yaml` is ignored by Git and contains blank credential placeholders; `config.example.yaml` is the distributable template. No real credentials are compiled or added to docs. Missing credentials fail before connection.

## Cross-tool boundaries

| Component | Effect / proof |
| --- | --- |
| Extractor | `LoadAppConfig` resolves YAML/legacy JSON to existing `AppConfig`; `DatabaseConfig.ConnectionString` encodes literal user/password/catalog. CLI tests capture the real exporter boundary in export and dry-run, assert summary excludes credentials, and compare config/history after repeated runs. Actual driver-parser tests cover reserved characters, literal percent sequences, whitespace, Unicode, and IPv6. |
| Interactive checker | Does not read extractor YAML. Existing apply/no-migration tests continue to govern; no new safe-fixer is introduced. |
| Checker `-check` / Linux MR | Does not read extractor YAML or need a DB. Existing no-write fixture/contract tests and released MR command remain the gate. |
| Formatter / verifier / generator | No parser, canonicalization, delta, up/down, or history changes. Existing fixture suite remains the applicable proof; no artificial migration fixture is added for credentials. |
| `build_all.bat` | Existing complete Windows/Linux authoring payload, license/checksum/identity/source-digest publication is retained. Rebuild required because shared `internal/` changes affect the release digest. It does not package extractor; compile `cmd/rza_get_object_db` separately or use the existing Windows bundle builder. |

Config order: explicit `-config` > working-directory `config.yaml` > `MSSQL_DSN` > `RZA_GET_OBJECT_DSN` > defaults. Invalid selected files fail without fallback. Preview without explicit config skips auto-load. Relative repo/object paths remain relative to the working directory. Structured `host` takes a hostname/IP without instance or port; omitted/zero port resolves to 1433. Legacy DSN remains available for other existing driver options. YAML credentials are literal, not pre-encoded, with single-quote escaping described in README.

## Focused checks

- `internal/getobject/database_config_test.go`: actual `msdsn.Parse` consumer checks, YAML defaults/options, invalid/ambiguous config without secret disclosure, YAML template round trip.
- `internal/getobject/app_config_test.go`: legacy JSON loading/validation and existing defaults.
- `cmd/rza_get_object_db/config_test.go`: file/env priority, invalid local and missing explicit files, offline preview, actual CLI export/dry-run config routing, no config/history mutation, no summary DSN/password.

No live SQL Server authentication is claimed from these offline checks. Credential syntax/driver parsing and a successful real network login are distinct validation levels.

## Executed validation

- `go test ./... -count=1`: passed after rebuilding the release. The first run correctly blocked on the old release source digest; no test was weakened to bypass it.
- `scripts/run_final_hardening_gate.ps1 -SkipGoTest`: passed; full Go tests ran separately. Fixture contract, migration-lock/runner-stop checks, generated coverage, repeatability, and benchmark passed. Learning-matrix steps were skipped automatically because that optional source tree is absent.
- `build_all.bat 9 3`: published the complete six-binary authoring release, checksums, metadata/manifest, and generated license using the existing owner and three-month term. Version `20261103.043446`; source SHA-256 `1577459139067bcf7738d4fd42d4750f54001846056e64c0191832a378c76563`.
- WSL Ubuntu 24.04: complete `release.sha256`, embedded checker `-build-info`, and `verify_release_source.sh` passed. The exact GitLab checker flags ran successfully on a temporary copy of QA-07 after interactive safe formatting; all `db_object` and migration file hashes were identical after `-check`. A local extractor YAML with blank credentials did not affect either checker mode. The initial unformatted fixture was correctly blocked by read-only `FMT_CHECK_CHANGES_REQUIRED` before preparation.
- Separately compiled `out/rza_get_object_db.exe`: explicit YAML with synthetic special-character credentials and implicit offline preview both processed one procedure successfully; summary contained no password or DSN.

Local execution logs live under ignored `out/extractor_yaml_*.log`; they are not committed artifacts. These checks validate native implementation and released consumers, not an independent agent evaluation.

## Documentation review

Updated: README (configuration/use), spec (requirements), FSD (flow), plan (scope/gates), walkthrough (selection), engine reference/call tree/package catalog (navigation), bootstrap (per-catalog setup), capability matrix (configuration-only classification), traceability and QA gate (tests), and this review.

Reviewed and intentionally unchanged: `rules.md` (SQL hard locks are unchanged); desired-state lifecycle and guarded migrations (no interpretation/delta/up/down changes); constraint naming (no constraint work); fixture README/goldens and learning backlog (no SQL rule/output scenario or unsupported SQL feature is changed); production readiness, engine docs guide, GitLab checker job and MR template (existing split authoring/extractor packaging and release procedure still apply). Generated QA coverage is regenerated using its script, rather than edited manually. No skill maintenance is needed: URL-encoded SQL Server credentials follow an established driver pattern rather than a new reusable MSSQL object-engine lesson.
