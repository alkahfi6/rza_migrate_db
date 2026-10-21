# Tutorial: Memahami Authoring Tanpa Koneksi Database

Tutorial ini membuat repo latihan baru dari satu table, menghasilkan migration awal, memeriksa source, lalu menambahkan satu kolom melalui draft/seal. Semua command dijalankan di PowerShell dari root source `rza_migrate_db`. Tidak ada command yang mengeksekusi SQL ke SQL Server.

Prasyarat: Go sesuai `go.mod`, dependency dapat diakses, dan `tools/license.key` yang masih sah. Tutorial tidak menerbitkan license baru. Jika license tidak tersedia/expired, pulihkan paket sah melalui workflow release sebelum melanjutkan. Untuk konsep baca [overview](../architecture/overview.md); kontrak command ada di [engine reference](../engine_reference.md).

## 1. Buat target latihan terpisah

```powershell
$ErrorActionPreference = 'Stop'
$sourceRepo = (Get-Location).Path
$labRepo = Join-Path $sourceRepo ('out/docs_lab_' + [Guid]::NewGuid().ToString('N'))
foreach ($folder in @('db_object/types', 'db_object/tables', 'db_object/indexes', 'db_object/views', 'db_object/functions', 'db_object/procedures', 'patch', 'migrations', 'tools')) {
    New-Item -ItemType Directory -Path (Join-Path $labRepo $folder) -Force | Out-Null
}
Copy-Item -LiteralPath (Join-Path $sourceRepo 'tools/license.key') -Destination (Join-Path $labRepo 'tools/license.key')
$tablePath = Join-Path $labRepo 'db_object/tables/dbo.DocsExample.sql'
$utf8 = [Text.UTF8Encoding]::new($false)
$tableSQL = @'
CREATE TABLE [dbo].[DocsExample]
(
    [Id] INT NOT NULL,
    CONSTRAINT [PK_DocsExample] PRIMARY KEY CLUSTERED ([Id] ASC)
);
'@
[IO.File]::WriteAllText($tablePath, $tableSQL + "`n", $utf8)
$labRepo
```

Semua `-repo` berikut harus menunjuk `$labRepo`. Folder berada di `out/` yang diabaikan Git. Migration history repository project asli tidak disentuh.

## 2. Bootstrap history awal

```powershell
go run ./cmd/rza_migrate_db bootstrap -repo $labRepo -pause-on-exit false
if ($LASTEXITCODE -ne 0) { throw 'bootstrap failed' }
Get-ChildItem -LiteralPath (Join-Path $labRepo 'migrations') -Name
```

Hasil yang diharapkan: exit `0`, core audit up/down dan satu pair table. Bootstrap menambahkan helper opsional yang belum ada tanpa overwrite. Timestamp filename mengikuti waktu run; jangan mencocokkan timestamp contoh secara literal. Baca kedua arah SQL sebelum melanjutkan.

## 3. Checker read-only dan verify

```powershell
$migrationBefore = @(Get-ChildItem -LiteralPath (Join-Path $labRepo 'migrations') -File | Sort-Object Name | ForEach-Object { $_.Name + ':' + (Get-FileHash -LiteralPath $_.FullName -Algorithm SHA256).Hash })
$sourceBefore = (Get-FileHash -LiteralPath $tablePath -Algorithm SHA256).Hash
go run ./cmd/rza_migrate_check -repo $labRepo -check -profile authoring -migration-validation integrity -pause-on-exit false
if ($LASTEXITCODE -ne 0) { throw 'checker failed' }
$migrationAfter = @(Get-ChildItem -LiteralPath (Join-Path $labRepo 'migrations') -File | Sort-Object Name | ForEach-Object { $_.Name + ':' + (Get-FileHash -LiteralPath $_.FullName -Algorithm SHA256).Hash })
if (@(Compare-Object $migrationBefore $migrationAfter).Count -ne 0) { throw 'history changed' }
if ((Get-FileHash -LiteralPath $tablePath -Algorithm SHA256).Hash -ne $sourceBefore) { throw 'source changed' }
go run ./cmd/rza_migrate_db verify -repo $labRepo -pause-on-exit false
if ($LASTEXITCODE -ne 0) { throw 'verify failed' }
```

Checker tidak mengubah table/history. `out/output.json` tetap dapat berubah karena diagnostic baru. `verify` pada tahap ini lulus karena history sudah mewakili desired state.

## 4. Tambahkan genuine delta melalui draft/seal

```powershell
$tableSQL = @'
CREATE TABLE [dbo].[DocsExample]
(
    [Id] INT NOT NULL,
    [Note] NVARCHAR(100) NULL,
    CONSTRAINT [PK_DocsExample] PRIMARY KEY CLUSTERED ([Id] ASC)
);
'@
[IO.File]::WriteAllText($tablePath, $tableSQL + "`n", $utf8)
go run ./cmd/rza_migrate_check -repo $labRepo -pause-on-exit false
if ($LASTEXITCODE -ne 0) { throw 'apply check failed' }
go run ./cmd/rza_migrate_db draft -repo $labRepo -pause-on-exit false
if ($LASTEXITCODE -ne 0) { throw 'draft failed' }
$historyAfterDraft = @(Get-ChildItem -LiteralPath (Join-Path $labRepo 'migrations') -File | Sort-Object Name | ForEach-Object { $_.Name + ':' + (Get-FileHash -LiteralPath $_.FullName -Algorithm SHA256).Hash })
if (@(Compare-Object $migrationBefore $historyAfterDraft).Count -ne 0) { throw 'apply checker or draft changed history' }
$draftFiles = @(Get-ChildItem -LiteralPath (Join-Path $labRepo 'out/draft_migrations') -Filter '*.sql')
if ($draftFiles.Count -ne 2) { throw 'expected one draft pair' }
$draftHashes = @{}
foreach ($file in $draftFiles) { $draftHashes[$file.Name] = (Get-FileHash -LiteralPath $file.FullName -Algorithm SHA256).Hash }
Get-ChildItem -LiteralPath (Join-Path $labRepo 'out/draft_migrations') -Filter '*.sql' | ForEach-Object { Get-Content -LiteralPath $_.FullName }
```

Expected: satu draft pair baru untuk perubahan `Note`; belum ada pair baru pada `migrations/`. Up menambah kolom, down menghapusnya dengan guard yang berlaku. Ini bukti struktur SQL yang dihasilkan, bukan bukti bahwa down aman untuk semua data live.

Setelah meninjau SQL draft pada repo latihan, publish reviewed bytes:

```powershell
go run ./cmd/rza_migrate_db seal -repo $labRepo -pause-on-exit false
if ($LASTEXITCODE -ne 0) { throw 'seal failed' }
foreach ($name in $draftHashes.Keys) {
    if ((Get-FileHash -LiteralPath (Join-Path $labRepo ('migrations/' + $name)) -Algorithm SHA256).Hash -ne $draftHashes[$name]) { throw 'sealed bytes differ from review' }
}
go run ./cmd/rza_migrate_db gen -repo $labRepo -pause-on-exit false
if ($LASTEXITCODE -ne 0) { throw 'second generation failed' }
$result = Get-Content -LiteralPath (Join-Path $labRepo 'out/output.json') -Raw | ConvertFrom-Json
if ($result.new_migrations.files.Count -ne 0) { throw 'unexpected second delta' }
$result.pass.items | Where-Object rule -eq 'GEN_NO_CHANGES'
```

Generation kedua harus no-op (`GEN_NO_CHANGES`, empty new migration list). History lama tetap immutable; seal hanya menambah pair baru. Pada task nyata, review draft dan scope perubahan tetap diperlukan sebelum publikasi history.

## 5. Buat report dan buka hasilnya

```powershell
go run ./cmd/rza_migrate_report -repo $labRepo -pause-on-exit false
if ($LASTEXITCODE -ne 0) { throw 'report failed' }
Get-Content -LiteralPath (Join-Path $labRepo 'docs/database/README.md')
```

Index mengarah ke current desired-state documentation dan audit history. Karena contoh hanya satu table, kosongnya procedure/relationship bukan kegagalan scanner. Hasil report tidak melakukan migration execution.

Folder latihan dapat disimpan untuk eksplorasi. Catat `$labRepo` bila akan membersihkannya; jangan menghapus folder project atau history asli. Untuk jalur live extractor/deployment lanjutkan ke [bootstrap four repos](../bootstrap_four_repos.md), [walkthrough](../walkthrough.md), dan [monitoring runbook](../operations/migration_monitoring.md).
