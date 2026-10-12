# Database report output layout

`rza_migrate_report` publishes layout V2 in one stable documentation root.
Default CLI, discovery from a subfolder, and double-click use the same layout:

```text
docs/database/
  README.md
  manifest.json
  identifiers/table_unique_ids.md
  relationships/relational_erd.md
  relationships/relationship_catalog.md
  relationships/module_lineage.md
  procedures/README.md
  procedures/procedure_impact_catalog.md
  procedures/profiles/dbo.<Procedure>.md
  migrations/README.md
  migrations/report.json
  migrations/datasets/report_summary.csv
  migrations/datasets/migration_summary.csv
  migrations/datasets/object_history.csv
  migrations/datasets/table_changes.csv
  migrations/datasets/index_changes.csv
  migrations/datasets/programmable_changes.csv
  migrations/datasets/type_changes.csv
```

Start with `README.md`. Each document links back to that page. Current-state
documentation reads `db_object/`; migration datasets read `migrations/`.
These represent different sources and may legitimately disagree. Warnings and
errors from history analysis also appear in `migrations/README.md`.

## CLI and compatibility

```powershell
tools/rza_migrate_report.exe -repo . -pause-on-exit false
tools/rza_migrate_report.exe -repo . -docs documentation/database -pause-on-exit false
tools/rza_migrate_report.exe -repo . -docs documentation/database -out audit/history -pause-on-exit false
```

Relative paths are always resolved from the discovered repository root.
`-docs` controls the documentation root. Without `-out`, migration output follows
`<docs>/migrations/`. An explicit `-out` selects an audit directory; CSVs are
still grouped in its `datasets/` child. Navigation supports spaces and `#` in
paths and procedure identities. Roots may be separate, but artifact paths may
not collide. Targets inside `db_object/` or source `migrations/`, including
symlink aliases, are rejected before publication.

This intentionally changes output paths: integrations reading `out/report/*.csv`
must switch to the new paths or use the concrete manifest. Existing CSV columns,
JSON history schema/summary semantics, and warning exit behavior remain unchanged.
`reports[]` now lists paths relative to `documentation_directory` (absolute paths
only when a cross-volume path cannot be relative); `artifacts[]` identifies each
file with `root` and `path`. `documentation_index` locates the root README from the audit directory. No wildcard inventory is emitted. The legacy
`table_relationship_tree.md` pointer is no longer generated; use the root index
and `relationships/` documents.

## Manifest and file ownership

`manifest.json` has schema `RZA_DATABASE_DOCUMENTATION_V2`, repository/source
metadata, creation time, both resolved roots, relative `root_paths`, and every concrete artifact. Relative root locations preserve ownership after cloning to another directory or switching operating systems; absolute roots remain provenance.
`sha256` hashes exact output bytes. The manifest lists itself without a hash,
avoiding a recursive self-hash. Names remain stable between runs. Only time
metadata in manifest/history JSON and the summary CSV changes routinely.

All content is staged before validating existing targets. A manual file at a
reserved output path, or an edited generated file whose checksum no longer
matches, blocks publication before any published file changes. Move annotations
to another filename or use another output root. Symlinks in managed child paths
are rejected. Individual replacements are atomic; the inventory is published
last. Whole-folder publication is not transactional: an unexpected I/O failure
after validation can leave a partial update and returns an error.

A deleted procedure's stale profile is removed only when its previous manifest
path is allowed and its SHA-256 still matches. Edited stale profiles and files
absent from the ownership inventory remain untouched. Changing an explicit
output root retains output at the previous root; it is not recursively erased.

Legacy flat Markdown files and default `out/report/` exports are removed only
when regenerated content is equivalent. Only `run_at` / generated timestamp and
layout metadata are ignored when comparing summary/history metadata. Modified,
unknown, or no-longer-reconstructible legacy files are preserved with a console
notice. Unowned split diagrams are retained for manual review. No documentation
directory receives recursive cleanup. Temporary staging is tool-owned.

## Cross-tool boundary and release

| Component | Behavior |
| --- | --- |
| Extractor | Existing canonical `db_object` extraction and operational `out/` summary unchanged. |
| Interactive checker | Existing safe corrections/validation unchanged; no report generation or migration generation. |
| Checker `-check` / Linux MR gate | Read-only source validation and existing `out/output.json` unchanged; no documentation rewrite. |
| Migration generator | Canonicalization, diff planning, and genuine up/down SQL unchanged. |
| Report | New grouped layout, navigation, inventory, ownership checks, and conservative cleanup; source/history remain read-only. |
| Release | Complete Windows/Linux suite must be rebuilt atomically; checksum, source digest, and Linux gate must pass. |

Proof: `internal/migreport/layout_test.go`, CLI discovery/override tests,
QA-31 golden/layout assertions, existing source/history immutability tests, and
the full fixture/hardening gates. Layout tests cover navigation, hash inventory,
rerun stability, repository clone/regeneration, retained audit roots after changing -docs, manual collisions, stale ownership, nested symlinks, and overlap.

Updated documents: README, rules, spec, FSD, plan, walkthrough, capability matrix,
traceability, QA gate, QA-31 fixture/golden, coverage (generated), edge-case backlog,
production readiness, engine docs guide, JOIN listing guide, and this contract.
Reviewed but unchanged: desired-state lifecycle, guarded migrations, bootstrap
four repositories, constraint naming, parallel safety, REST hardening, migration
monitoring, rebaseline, legacy audit policy, GitLab checker/job/MR templates, and
GitHub publication scope. This change has no schema, SQL, execution, parallelism,
CI command, or repository publication policy delta.

## Verified delivery

2026-10-12: complete build `20261012.020616`; full Go suite, final hardening gate, Linux report package tests (including symlinks/clone), release SHA-256/source identity, and the exact Linux MR checker command passed. Released Windows and Linux report binaries regenerated the same checkout successfully in both directions. Current workspace inventory contains 17 generated files, 16 verified file hashes (self manifest excluded), and 22 valid Markdown file links. Source/history hashes stayed unchanged. Learning examples remain unavailable and the existing gate reports that directory as skipped. Proof logs are local `out/report-layout-*`; they are not repository artifacts.

The workspace legacy history snapshot included the bootstrap pair in application counts. It was reviewed and preserved manually under `docs/database/migrations/legacy/`; the three empty legacy relationship/lineage documents were reviewed against their unchanged Git baseline and removed after their grouped replacements were generated. Automatic cleanup remains conservative for unverified files in other repositories.
