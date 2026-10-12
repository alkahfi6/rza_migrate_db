# Private GitHub repository

Repository preparation excludes local trial copies (`App-DB/`, `all_object/`,
`real-case-pipeline-repo/`, `test/`), temporary files, generated `out/` reports,
root ZIP archives, local logs, and `mssql_credential.nfm` / `.env` credentials.
These files remain on the local computer; `.gitignore` prevents accidental staging.

Keep `tests/`, `fixtures/`, and Go `_test.go` files: these are the maintained
regression suite, required by the engine QA gates. Keep `migrations/` and the
complete `tools/` release, including its license, checksums, and build identity.
Linux executables and Windows `.exe` files are marked binary in `.gitattributes`
so Git never changes their bytes during checkout.

This publication does not change extraction, interactive checker corrections,
read-only `-check`, report inference, migration generation, or the existing
GitLab pipeline. GitHub hosts the source and release artifacts privately; the
GitLab configuration remains available for repositories that run that pipeline.

Reviewed documentation: engine docs guide, README, spec, FSD, rules, plan,
walkthrough, and release source verification contract. Engine behavior and
requirements remain unchanged, so their normative contracts require no change.
