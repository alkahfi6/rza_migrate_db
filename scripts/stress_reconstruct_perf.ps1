param(
  [string]$RepoRoot = ".",
  [int]$ObjectCount = 10000,
  [int]$BudgetMs = 120000,
  [switch]$KeepTemp
)

$ErrorActionPreference = "Stop"
Set-Location $RepoRoot

if ($ObjectCount -lt 1) {
  throw "ObjectCount must be >= 1"
}

$seed = "examples/rules_2x_learning/01_pass_minimal"
if (-not (Test-Path $seed)) {
  throw "seed scenario not found: $seed"
}

$tmpRoot = Join-Path $env:TEMP ("rza_reconstruct_perf_" + [guid]::NewGuid().ToString("N"))
New-Item -ItemType Directory -Path $tmpRoot -Force | Out-Null
$repo = Join-Path $tmpRoot "repo"
Copy-Item -Recurse -Path $seed -Destination $repo
$migrationsDir = Join-Path $repo "migrations"

Get-ChildItem -Path $migrationsDir -File -Filter "*.sql" |
  Where-Object { $_.Name -notin @("000000__core__mig__audit_init.up.sql", "000000__core__mig__audit_init.down.sql") } |
  Remove-Item -Force

$baseVersion = [int64]20270101000000000
for ($i = 0; $i -lt $ObjectCount; $i++) {
  $version = "{0:D17}" -f ($baseVersion + $i)
  $object = "dbo.ViewPerf{0:D5}" -f $i
  $upName = "{0}_views_{1}_v1.up.sql" -f $version, $object
  $downName = "{0}_views_{1}_v1.down.sql" -f $version, $object
  $upSql = "CREATE OR ALTER VIEW $object AS SELECT $i AS id;`n"
  $downSql = "CREATE OR ALTER VIEW $object AS SELECT ($i-1) AS id;`n"
  Set-Content -Path (Join-Path $migrationsDir $upName) -Value $upSql
  Set-Content -Path (Join-Path $migrationsDir $downName) -Value $downSql
}

$sw = [System.Diagnostics.Stopwatch]::StartNew()
& go run -buildvcs=false ./cmd/rza_migrate_db verify -repo $repo *> $null
$exit = $LASTEXITCODE
$sw.Stop()
$elapsed = [int]$sw.ElapsedMilliseconds

$outPath = Join-Path $repo "out/output.json"
$failCount = -1
if (Test-Path $outPath) {
  $report = Get-Content -Raw $outPath | ConvertFrom-Json
  if ($null -ne $report.fail -and $null -ne $report.fail.count) {
    $failCount = [int]$report.fail.count
  }
}

$summary = [PSCustomObject]@{
  object_count = $ObjectCount
  elapsed_ms = $elapsed
  budget_ms = $BudgetMs
  command_exit = $exit
  fail_count = $failCount
  temp_repo = $repo
  pass_budget = ($elapsed -le $BudgetMs)
}

$summaryPath = Join-Path (Get-Location) "out/reconstruct_perf_summary.json"
New-Item -ItemType Directory -Path (Split-Path -Parent $summaryPath) -Force | Out-Null
$summary | ConvertTo-Json -Depth 4 | Set-Content -Path $summaryPath

Write-Host "Reconstruct perf summary:" -ForegroundColor Cyan
Write-Host ($summary | ConvertTo-Json -Depth 4)
Write-Host "Saved: $summaryPath" -ForegroundColor Cyan

if (-not $KeepTemp) {
  Remove-Item -Recurse -Force $tmpRoot
}

if ($exit -ne 0) {
  throw "verify command failed with exit code $exit"
}
if ($failCount -gt 0) {
  throw "verify produced fail_count=$failCount"
}
if ($elapsed -gt $BudgetMs) {
  throw "reconstruct performance budget exceeded: ${elapsed}ms > ${BudgetMs}ms"
}

exit 0
