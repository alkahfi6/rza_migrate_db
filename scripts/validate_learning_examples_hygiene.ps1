param(
  [string]$RepoRoot = ".",
  [switch]$Fix
)

$ErrorActionPreference = "Stop"
Set-Location $RepoRoot

$matches = Get-ChildItem -Path "examples/rules_2x_learning" -Recurse -File -Filter "output.json" |
  Where-Object { ($_.FullName -replace "\\", "/") -match "/out/output\.json$" }

if ($matches.Count -gt 0) {
  if ($Fix) {
    foreach ($item in $matches) {
      Remove-Item -Path $item.FullName -Force
    }
    $matches = Get-ChildItem -Path "examples/rules_2x_learning" -Recurse -File -Filter "output.json" |
      Where-Object { ($_.FullName -replace "\\", "/") -match "/out/output\.json$" }
  }
}

if ($matches.Count -gt 0) {
  Write-Host "Found tracked-like example output artifacts (should not exist):" -ForegroundColor Red
  $prefix = (Get-Location).Path + [System.IO.Path]::DirectorySeparatorChar
  foreach ($item in $matches) {
    Write-Host ("- " + $item.FullName.Replace($prefix, "")) -ForegroundColor Red
  }
  throw "Learning examples hygiene failed: remove examples/**/out/output.json artifacts."
}

Write-Host "Learning examples hygiene passed." -ForegroundColor Green
exit 0
