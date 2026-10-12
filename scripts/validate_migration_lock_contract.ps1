param(
  [string]$RepoRoot = "."
)

$ErrorActionPreference = "Stop"
Set-Location $RepoRoot

$templatePath = ".gitlab/migration_jobs/migrate_with_lock.yml"
if (-not (Test-Path $templatePath)) {
  throw "Missing migration lock template: $templatePath"
}

$content = Get-Content -Raw $templatePath
$requiredTokens = @(
  "resource_group:",
  'test -n "$MSSQL_DSN"',
  'migrate -path ./migrations -database "$MSSQL_DSN" up'
)

$missing = @()
foreach ($token in $requiredTokens) {
  if (-not $content.Contains($token)) {
    $missing += $token
  }
}

if ($missing.Count -gt 0) {
  throw ("Migration lock template is missing token(s): " + ($missing -join ", "))
}

$ciPath = ".gitlab-ci.yml"
if (-not (Test-Path $ciPath)) {
  throw "Missing CI pipeline file: $ciPath"
}
$ciContent = Get-Content -Raw $ciPath
if ($ciContent.Contains("migrate_with_lock.yml")) {
  throw "Authoring checker pipeline must not include the optional migration execution template"
}

Write-Host "Migration lock contract validation passed." -ForegroundColor Green
exit 0
