# Call Tree Engine Suite

Call tree ini menampilkan jalur utama yang ditelusuri dari source. Ini bukan dump seluruh helper atau runtime trace. Indentasi berarti pemanggilan dalam jalur tersebut; cabang bertanda kondisi hanya berjalan pada mode yang sesuai. Callback dari `verify` kembali ke `engine` dijelaskan eksplisit. Gunakan nama simbol dengan `rg -n` untuk menemukan lokasi terbaru, karena nomor baris mudah berubah.

## 1. Extractor

Source: [CLI](../../cmd/rza_get_object_db/main.go), [export](../../internal/getobject/export.go), [batch metadata](../../internal/getobject/export_batch.go), [preview](../../internal/getobject/preview.go).

```text
main -> run
  ValidateExecutableName + parse flags
  loadCompiledGetObjectConfig / getobject.LoadAppConfig (-config)
  [preview-input] previewFromDir -> getobject.PreviewModulesFromDir
  [export] exportFromDatabase -> getobject.ExportFromDatabaseWithOptions
    Normalize / Validate config
    loadInitialMigrationObjects
    sql.Open -> PingContext
    loadTableTypeFiles
    loadTables -> loadTableSnapshotCache
    buildTableAndIndexFiles
    loadProgrammabilityFiles (views, functions, procedures)
    [strict, unsafe object] return error sebelum apply
    applySQLCategoryPreserving (setiap kategori, dengan DryRun)
    [bukan dry-run] management.Write
  finishRun -> out/get_object_summary.json + exit
```

Metadata dikumpulkan sebelum apply. Preserve set menjaga snapshot unsupported saat skip diizinkan. Apply memakai atomic replacement per file; hal itu bukan transaksi atomik seluruh export lintas kategori. Preview tidak membuka DB, tetapi menulis hasil formatting ke direktori preview.

## 2. Checker apply dan read-only

Source: [entrypoint](../../cmd/rza_migrate_check/main.go), [CLI checker](../../internal/cmd/migrate_check.go), [engine](../../internal/engine/engine.go), [constraint reconciliation](../../internal/engine/constraint_name_reconcile.go).

```text
main -> cmd.RunMigrateCheck
  ValidateExecutableName / parse flags / resolveCheckerProfile
  config.DiscoverRepoRoot -> config.Default -> validateCheckerLicence
  engine.New(cfg).Run("check")
    cfg.Validate
    runDeveloperCheck
      runFmt
        encoding + sqlfix.RewriteSafe + canonical formatting
        index filename correction jika target tunggal terbukti
        [EnableFmtCheck] laporkan perubahan; jangan tulis authored SQL
        [apply] tulis correction yang aman
      [tanpa fail] runConstraintNameFix
        [migration-validation=off] skip reconciliation
        reconstructHistory + collectDesiredObjects
        planDesiredTableConstraintNames
        [check] laporkan pending correction
        [apply] applyConstraintNameWrites
      [tanpa fail] resetRepositorySnapshot
      runVerifyChecks (wrapper hanya untuk migration-validation=full)
      recordCheckAttestation
    Report.WriteWithOptions
  output.PrintFailureDetails -> exit 0/1
```

`-check` mengaktifkan `EnableFmtCheck`, bukan command generator lain. Jika format/correction sudah gagal, checker dapat berhenti sebelum fase berikutnya. Setelah diperbaiki, jalankan ulang agar validasi berikutnya terjangkau. Default checker memakai `authoring/integrity`; profile `full` mewajibkan migration validation `full`. Keduanya tidak menghasilkan pair.

## 3. Generator, formatter, verifier, draft/seal

Source: [CLI](../../cmd/rza_migrate_db/main.go), [engine routing](../../internal/engine/engine.go), [planner](../../internal/engine/gen_emit.go), [publication](../../internal/engine/gen_emit_commit.go).

```text
main -> run
  parse command (default verify_then_gen), flags/env, discover repo, license
  runCommandWithProgress -> engine.New(cfg).Run(command)
    cfg.Validate
    [BootstrapMode] seedOptionalRepoFilesForBootstrap + seedCoreAuditForBootstrap
    switch command:
      fmt -> runFmt
      verify -> runVerify -> runVerifyChecks(includeWrapper=true)
      gen / draft -> runGen
        runFmt -> runGenPrecheck -> runGenCore
      verify_then_gen / bootstrap -> runVerifyThenGen
        runFmt -> runGenPrecheck -> runGenCore
      seal -> runSealDrafts
    Report.WriteWithOptions
  [auto-batch + combined/bootstrap + pending] ulang dengan gen
  exit dari error atau report.Fail.Count

runGenCore
  [fail sebelumnya] return tanpa emit
  reuse precheck cache / reconstructHistory
  reuse desired cache / collectDesiredObjects
  reconcileDesiredTableConstraintNames
  [source berubah] reset snapshot + recollectDesiredObjects
  collectPatchSources
  [EnableGenAutoIndent] autoIndentChangedModuleSources
  [patch mode] autoIndentPendingPatchSources
  [source berubah] reset snapshot + recollect
  buildMigrationPlan -> planGenerationWithInputs
  [tanpa delta] GEN_NO_CHANGES
  validatePlannedRoundTrip
  emitPlannedMigrations
    allocateMigrationBatch -> renderAllocatedMigrationFiles
    writeEmitStaging -> verifyStagedMigrations
    [valid] commitStagedFiles ke migrations atau out/draft_migrations
    [gagal] diagnostic + tidak publish hasil final

runSealDrafts
  validasi combined history + reviewed draft
  publishDraftFiles (exclusive create, bytes hasil review)
  [sukses] hapus draft yang sudah dipublikasikan
```

Nama `commitStagedFiles` merujuk publication file migration, bukan Git commit. Reuse cache adalah optimasi; ia tidak boleh mengubah hasil. `verify_then_gen` memakai combined precheck supaya reconstruction dapat dipakai ulang, bukan dua proses CLI terpisah. `fmt` sendiri tidak menjalankan history-name reconciliation checker/gen; jangan menyamakan seluruh command authoring sebagai satu aksi.

## 4. Shared verify orchestration

Source: [engine bridge](../../internal/engine/verify_callbacks.go), [orchestrator](../../internal/verify/verify.go), [engine runVerifyChecks](../../internal/engine/engine.go).

```text
engine.runVerifyChecks
  ensureRepositorySnapshot
  verify.Orchestrator.Run(report, engine.verifyCallbacks(cache))
    runMigrationsDomain
      [off] skip domain
      ValidateMigrations -> engine.validateMigrations -> verify validator
      [integrity] return metadata tanpa semantic history reconstruction
      [full, belum fail] ReconstructHistory -> engine.reconstructHistory
      [full] ValidateCoreAuditContracts + ValidateForbiddenSQL
    runPatchDomain
      ValidatePatchFolders
      [migration mode bukan off] ValidatePatchSourceIntegrity
    runDbObjectDomain (serial / bounded parallel dengan stable merge)
      ValidateTableInlineDefault / ValidateRepoHygiene
      ValidateDbObjectConsistency / CollectDesiredObjects
      [full] ValidateDesiredHistoryDrift
      ValidateProgrammabilityHeaders
    [IncludeWrapper] runWrapperDomain -> ValidateWrapperSubset
  [mode bukan full] validateForbiddenSQL
```

Callback menyimpan model privat engine; paket `verify` tidak memiliki planner. Skip reconstruction pada domain integrity di atas hanya berlaku di orchestrator: checker masih dapat merekonstruksi history lebih awal untuk constraint-name reconciliation. Profil integrity tidak berarti seluruh SQL historis sudah diaudit secara semantik. Forbidden SQL checks tetap berjalan melalui jalur di luar semantic history validation. `FailFast` dapat menghentikan fase lanjutan.

## 5. Report offline

Source: [CLI](../../cmd/rza_migrate_report/main.go), [Run](../../internal/migreport/report.go), [desired analysis](../../internal/migreport/desired_state.go), [publisher](../../internal/migreport/layout.go).

```text
main -> run
  flags / config.DiscoverRepoRoot
  migreport.Run(opts)
    DiscoverRepository + resolve output/docs directories
    validateDirectoryOutsideProtected
    ScanMigrations -> excludeInitialCoreMigration
    Analyze (history reconciliation)
    BuildDesiredStateDocs
      scanSQLFiles tables -> parseTableDefinition
      scanSQLFiles indexes -> parseIndexes
      scanModules views/procedures/functions
      analyzeModuleRelationships
      buildProcedureProfiles
    reportArtifacts
    publishReportBundle (manifest ownership, replace, stale cleanup)
    PrintConsoleSummary -> status/exit
```

Static access/call parsing menggunakan helper `sqlparse`; analysis report dan migration reconstruction mempunyai tujuan berbeda. `sqlparse.ExtractObjectReferences` mengabaikan dynamic SQL untuk migration dependency; report dapat menganalisis literal dynamic tertentu dengan batas yang dinyatakan. Hasil report tidak dimasukkan kembali ke planner sebagai fakta schema. Output dapat ditulis meski summary analysis berstatus failed; periksa exit dan isi report.

## 6. REST execution

Source: [bootstrap](../../internal/cmd/rest.go), [handler](../../internal/handler/migration_handler.go), [service](../../internal/migrest/service.go), [lock](../../internal/migrest/lock.go), [journal](../../internal/migrest/journal.go).

```text
cmd/rza_migrate_rest.main -> cmd.RunREST
  migrest.LoadConfig -> logger + telemetry.Init
  migrest.NewService(NewServiceConfigFromRuntime(cfg))
  handler.NewMigrationHandler(...).Router
  WithTraceIDHeader(WithAccessLog(router)) -> otelhttp.NewHandler
  http.Server.ListenAndServe
  [signal] server.Shutdown -> deferred service/telemetry cleanup

HTTP request
  middleware -> withAuth -> endpoint handler -> service interface
    CurrentVersion / Pending -> migrator + state/listing
    Up / Down / RunSpecificFile -> runWithMigrationLockWithOptions
      mutex + operation context + configured DB lock
      migrator + before state + immutable applied-history evidence
      operation callback -> golang-migrate Up / Steps
      after state + operation report/journal + release lock
    AcknowledgeManualFile -> controlled recovery path (config + reason)
```

Down memerlukan `allow_down` dan steps positif; run-file up harus next version, down harus current version. Ack-manual adalah operasi recovery yang dapat mengubah version evidence tanpa mengeksekusi file tersebut, bukan pengganti normal migration. Detail error/status dan deployment ada pada [Swagger](../../internal/handler/swagger.yaml), [REST hardening](../policy/rest_api_hardening.md), dan [monitoring](../operations/migration_monitoring.md).

## 7. Utilitas pengujian database

| Tool | Jalur utama | Batas |
| --- | --- | --- |
| `rza_migrate_test` | `main -> run -> loadNFMCredentials -> command`; `up/down -> prepareCatalogMappedMigrations -> runUp/runDown -> migrest.Service` | `create-db`, `status`, `up`, `down`, `reset`; catalog-map diterapkan pada temporary copy, bukan history asli. Reset/create-db memakai confirmation target. |
| `rza_mssql_test` | `main -> run -> loadCredentialConfig -> sql.Open/Ping -> currentDatabase -> runStatus/runMigration/runReset` | Command default status; `down -all`/reset memerlukan allow-destructive. Flag dan exit invalid-input berbeda dari tool lain. |

Source: [migrate test](../../cmd/rza_migrate_test/main.go), [MSSQL test](../../cmd/rza_mssql_test/main.go). Keduanya membuka DB nyata; jangan dipakai sebagai smoke onboarding offline.

## 8. Build dan MR gate

```text
build_all.bat
  release_artifacts.ps1 -Mode Prepare (normalized source digest + build identity)
  build 3 Windows + 3 Linux binaries
  license_gen
  release_artifacts.ps1 -Mode Finalize (checksum + release manifest)
  release_artifacts.ps1 -Mode Publish (verified publication dengan rollback)

GitLab rza_migrate_check_gate
  complete release SHA-256 + checker SHA-256
  checker -build-info == build-info.txt
  verify_release_source.sh
  checker -check -profile authoring -migration-validation integrity
  out/output.json sebagai artifact; exit nonzero memblokir pipeline
```

Source: [build](../../build_all.bat), [release script](../../scripts/release_artifacts.ps1), [CI job](../../.gitlab/rza_migrate_check.gitlab-ci.yml). Wrapper `scripts/build_migrate_check_linux.bat` mendelegasikan ke complete build. Lihat [production readiness](../policy/production_readiness.md) untuk bundle extractor/REST yang terpisah.
