param(
  [string]$RepoRoot = "."
)

$ErrorActionPreference = "Stop"
Set-Location $RepoRoot

$templatePath = ".gitlab/migration_jobs/migrate_with_lock.yml"
if (-not (Test-Path $templatePath)) {
  throw "Missing migration template: $templatePath"
}

$content = Get-Content -Raw $templatePath
$normalized = $content -replace "`r`n", "`n"
$migratePattern = 'migrate\s+-path\s+\./migrations\s+-database\s+"\$MSSQL_DSN"\s+up'

if ($normalized -notmatch $migratePattern) {
  throw "Template must execute golang-migrate up command."
}

if ($normalized -match '(?m)\bset\s+\+e\b') {
  throw "Runner stop behavior is unsafe: shell must not use `set +e` around migrate execution."
}

if ($normalized -match ($migratePattern + '\s*(\|\||;)\s*true(\s|$)')) {
  throw "Runner stop behavior is unsafe: migrate command must not swallow errors (`|| true` / `; true`)."
}
if ($normalized -match ($migratePattern + '\s*\|\|\s*exit\s+0(\s|$)')) {
  throw "Runner stop behavior is unsafe: migrate command must not swallow errors (`|| exit 0`)."
}
if ($normalized -match ($migratePattern + '\s*\|\|\s*:\s*(\s|$)')) {
  throw "Runner stop behavior is unsafe: migrate command must not swallow errors (`|| :`)."
}

$lockPos = $normalized.IndexOf("resource_group:")
$migrateMatch = [regex]::Match($normalized, $migratePattern)
$migratePos = -1
if ($migrateMatch.Success) {
  $migratePos = $migrateMatch.Index
}
if ($lockPos -lt 0) {
  throw "Template must serialize migration jobs using resource_group."
}
if ($migratePos -lt 0) {
  throw "Template must contain migrate up command."
}
if ($lockPos -gt $migratePos) {
  throw "Migration job serialization must be declared before migrate up."
}

Write-Host "Runner stop-behavior contract validation passed." -ForegroundColor Green
exit 0
