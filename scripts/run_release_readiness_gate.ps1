param(
  [string]$RepoRoot = ".",
  [switch]$SkipEngineFixtures,
  [switch]$SkipRepeatability,
  [switch]$SkipBenchmark,
  [switch]$SkipBuild
)

$ErrorActionPreference = "Stop"
Set-Location $RepoRoot

$steps = @(
  [PSCustomObject]@{
    Name = "fast unit packages"
    Skip = $false
    Run  = { & go test -buildvcs=false ./internal/... ./cmd/... -count=1 }
  },
  [PSCustomObject]@{
    Name = "fixture golden contract"
    Skip = $false
    Run  = { & go test -buildvcs=false ./tests/fixtures -count=1 }
  },
  [PSCustomObject]@{
    Name = "engine fixture expectations"
    Skip = $SkipEngineFixtures
    Run  = { & go test -buildvcs=false ./tests/engine -count=1 }
  },
  [PSCustomObject]@{
    Name = "parallel repeatability"
    Skip = $SkipRepeatability
    Run  = { & powershell -ExecutionPolicy Bypass -File scripts/validate_parallel_repeatability.ps1 -Iterations 6 -Command verify }
  },
  [PSCustomObject]@{
    Name = "parallel benchmark"
    Skip = $SkipBenchmark
    Run  = { & powershell -ExecutionPolicy Bypass -File scripts/benchmark_parallel_read.ps1 -Iterations 3 -Command gen }
  },
  [PSCustomObject]@{
    Name = "windows distribution build"
    Skip = $SkipBuild
    Run  = { & powershell -ExecutionPolicy Bypass -File scripts/build_dist_windows.ps1 -SkipTests }
  }
)

$failed = @()
foreach ($step in $steps) {
  if ($step.Skip) {
    Write-Host "[SKIP] $($step.Name)" -ForegroundColor Yellow
    continue
  }

  Write-Host "[RUN ] $($step.Name)" -ForegroundColor Cyan
  try {
    & $step.Run
    if ($LASTEXITCODE -ne $null -and $LASTEXITCODE -ne 0) {
      throw "command exited with code $LASTEXITCODE"
    }
    Write-Host "[ OK ] $($step.Name)" -ForegroundColor Green
  } catch {
    $failed += $step.Name
    Write-Host "[FAIL] $($step.Name)" -ForegroundColor Red
    Write-Host $_.Exception.Message -ForegroundColor Red
  }
}

if ($failed.Count -gt 0) {
  Write-Host ""
  Write-Host "Release readiness gate failed on $($failed.Count) step(s):" -ForegroundColor Red
  foreach ($name in $failed) {
    Write-Host "- $name" -ForegroundColor Red
  }
  exit 1
}

Write-Host ""
Write-Host "Release readiness gate passed." -ForegroundColor Green
exit 0
