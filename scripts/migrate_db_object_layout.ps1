param(
  [string]$RepoRoot = ".",
  [switch]$DryRun
)

Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"

function Resolve-Root([string]$Path) {
  return (Resolve-Path -LiteralPath $Path).Path
}

function Ensure-Dir([string]$Path) {
  if (-not (Test-Path -LiteralPath $Path)) {
    New-Item -ItemType Directory -Path $Path | Out-Null
  }
}

$root = Resolve-Root $RepoRoot
$dbObject = Join-Path $root "db_object"
$legacyRoot = Join-Path $dbObject "schema"

if (-not (Test-Path -LiteralPath $legacyRoot)) {
  Write-Host "[INFO] legacy path not found: $legacyRoot"
  Write-Host "[INFO] nothing to migrate"
  exit 0
}

$categories = @("types", "tables", "indexes", "views", "functions", "procedures")
$moved = 0
$skipped = 0
$overwritten = 0

foreach ($cat in $categories) {
  $legacyDir = Join-Path $legacyRoot $cat
  $canonicalDir = Join-Path $dbObject $cat

  if (-not (Test-Path -LiteralPath $legacyDir)) {
    continue
  }

  Ensure-Dir $canonicalDir
  $files = Get-ChildItem -LiteralPath $legacyDir -File -Filter *.sql -ErrorAction SilentlyContinue
  foreach ($f in $files) {
    $dest = Join-Path $canonicalDir $f.Name
    $exists = Test-Path -LiteralPath $dest
    if ($exists) {
      $overwritten++
    }
    if ($DryRun) {
      if ($exists) {
        Write-Host "[DRYRUN] overwrite: $($f.FullName) -> $dest"
      } else {
        Write-Host "[DRYRUN] move:      $($f.FullName) -> $dest"
      }
      $skipped++
      continue
    }
    Copy-Item -LiteralPath $f.FullName -Destination $dest -Force
    $moved++
  }
}

if ($DryRun) {
  Write-Host "[INFO] dry-run complete"
  Write-Host "[INFO] planned files: $skipped (overwrite: $overwritten)"
  exit 0
}

# Remove legacy tree after copy succeeds.
Remove-Item -LiteralPath $legacyRoot -Recurse -Force

Write-Host "[INFO] migration complete"
Write-Host "[INFO] copied files: $moved (overwrite: $overwritten)"
Write-Host "[INFO] removed legacy folder: $legacyRoot"
