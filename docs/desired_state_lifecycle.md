# Desired-State Lifecycle Contract

`db_object` is the current desired state. `migrations` is immutable history.
Generation must reconcile every semantic difference between the two; an
unhandled difference is an error, never `GEN_NO_CHANGES`.

## Object removal

An object present in active migration history but absent from `db_object`
produces a real removal migration for the existing categories only:

- `tables`: guarded `DROP TABLE IF EXISTS`; blocked when rows, incoming foreign
  keys, SQL expression dependencies, or triggers remain. The emptiness check
  uses `TABLOCKX, HOLDLOCK` inside the migration transaction so concurrent
  inserts cannot race the emptiness check.
- `indexes`: drop each historical index and restore its prior definition in
  `down`.
- `types`: blocked while a column or parameter uses the type.
- `views`, `functions`, and `procedures`: blocked while another SQL object
  references the module.

Removal order is the reverse of creation dependency order. Explicit indexes are
removed before their table. `down` is executed by the migration runner in the
reverse version order, restoring dependencies safely.

The reconstructed history records a deleted state. Re-running `gen` is
idempotent, while adding the same object back to `db_object` creates the next
object version instead of resetting its history.

## Precheck policy

Prechecks follow three outcomes:

1. Object/state absent and the requested operation is already satisfied: skip.
2. Existing `CREATE TABLE` target satisfies the bootstrap-compatible column
   contract: skip.
3. Existing state differs, data is incompatible, a dependency remains, or
   equivalence cannot be proven: fail and roll back.

Name-only table existence is never sufficient. Every requested column must be
found using exact-case `Latin1_General_100_BIN2` comparison and must have a
compatible base type, identity-kind, and computed-kind.

For an initial `CREATE TABLE` whose target already exists, the
`bootstrap-compatible` profile intentionally does not block on:

- extra columns or column order;
- length, precision/scale, nullability, or collation;
- default, key, CHECK, or FOREIGN KEY definitions and runtime state;
- physical index options managed by DBAs.

This relaxation applies only to the existing-target guard around an initial
`CREATE TABLE`. Explicit `ALTER COLUMN`, `DROP COLUMN`, `ADD/DROP CONSTRAINT`,
and `CREATE/DROP INDEX` migrations retain operation-specific strict prechecks.
For explicit evolution, type parameters are part of desired state: changed
precision/scale, length, or date-time precision must remain visible even when
the source uses whitespace around `(` or `,`. A single delimited built-in type
name such as `[DECIMAL]` is semantically identical to `DECIMAL`; this
normalization does not erase changes to its precision or scale.
Existing rowstore index guards still compare key order/direction, include set,
filter, uniqueness, type, enabled state, and persisted rowstore options.

## Historical hardening boundary

Migration files are not rewritten. The acknowledged legacy batch ends before
version `20260726000000000`. Strict generated-precheck verification applies to
every later migration. This preserves immutable history while preventing new
unsafe scripts. Starting at `20260728000000000`, newly generated initial-table
wrappers must declare and satisfy the `bootstrap-compatible` guard contract.

## Extractor behavior

The extractor no longer embeds a database credential. Use `-config` with a JSON
extractor config, `MSSQL_DSN`, or the compatibility environment variable
`RZA_GET_OBJECT_DSN`.

The default extractor policy is complete-snapshot mode:
`skip_unsupported_tables=false`. If any supported-category object cannot be
rendered safely, extraction fails before applying files. Writes use a synced
temporary file followed by an atomic replace.

Disabled or untrusted CHECK/FOREIGN KEY constraints do not remove the table
from the logical snapshot. The extractor emits their definitions and records
structured `CONSTRAINT_STATE_REVIEW_REQUIRED` health findings without changing
the database. Non-default physical metadata on PK/UQ backing indexes is
reported as `KEY_PHYSICAL_METADATA_NOT_CAPTURED`. Use
`-require-healthy-constraints` to promote disabled/untrusted findings to a
non-zero exit after snapshots and evidence have been written.

For first-time adoption, `rza_get_object_db -profile bootstrap` enables
lossless-skip behavior. Every skipped table, index, type, view, function, or
procedure is written to `out/object_management_manifest.json`. Generation
excludes only those explicit category/object keys from desired-state removal;
an absent object without manifest evidence still produces a guarded removal.

## Patch formatting policy

Patch formatting is disabled by default because patch SQL may contain dynamic
SQL or literals whose bytes are semantically significant. `-format-patch safe`
formats only files that pass the semantic safety check and reports
`PATCH_FORMAT_SKIPPED_UNSAFE` for the rest. `strict` turns that warning into
`FORMAT_SAFETY_BLOCKED`. No mode rewrites an unsafe source.

## Draft and seal

`rza_migrate_db draft` generates and verifies migrations under
`out/draft_migrations`. Review those files, then run `rza_migrate_db seal`.
Seal verifies the combined immutable history, creates targets exclusively (it
never overwrites), syncs reviewed bytes, and only then removes draft sources.
If publication fails, already published files are rolled back and all draft
sources remain available. At most 999 pairs are generated per batch; seal one
batch before drafting the next.

## Applied history failures

Applied history remains a hard safety boundary. The production App-DB service
returns HTTP 409 with an actionable code:

- `applied_migration_file_missing`: deployment artifact omitted an applied file;
- `applied_migration_hash_mismatch`: an applied migration file changed;
- `applied_migration_evidence_invalid`: journal evidence is malformed or contradictory;
- `applied_migration_body_file_missing`: audited generated file is absent;
- `applied_migration_body_hash_mismatch`: audited migration body changed.

Do not edit a checksum to bypass these errors. Restore the exact reviewed file
or redeploy the complete immutable migration history. Only a never-applied,
complete up/down pair may be discarded and regenerated through draft/seal.

## Explicit non-goals

- Table rename inference is not implemented.
- No SQL object category was added.
- Migration history is not rewritten to match current generator formatting.
