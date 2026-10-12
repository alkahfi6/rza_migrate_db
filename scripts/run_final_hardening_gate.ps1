param(
  [string]$RepoRoot = ".",
  [switch]$SkipGoTest,
  [switch]$SkipFixtureContract,
  [switch]$SkipLearningMatrix,
  [switch]$SkipParallelRepeatability,
  [switch]$SkipParallelBenchmark
)

$ErrorActionPreference = "Stop"
Set-Location $RepoRoot
$learningRootExists = Test-Path "examples/rules_2x_learning"
$learningMatrixExists = Test-Path "examples/rules_2x_learning/learning_matrix.json"

$steps = @(
  [PSCustomObject]@{
    Name = "go test -buildvcs=false ./... -count=1"
    Skip = $SkipGoTest
    Run  = { & go test -buildvcs=false ./... -count=1 }
  },
  [PSCustomObject]@{
    Name = "validate_qa_fixture_contract.ps1"
    Skip = $SkipFixtureContract
    Run  = { & powershell -ExecutionPolicy Bypass -File scripts/validate_qa_fixture_contract.ps1 }
  },
  [PSCustomObject]@{
    Name = "verify_rules_2x_learning.ps1"
    Skip = $SkipLearningMatrix -or -not $learningMatrixExists
    Run  = { & powershell -ExecutionPolicy Bypass -File scripts/verify_rules_2x_learning.ps1 }
  },
  [PSCustomObject]@{
    Name = "validate_learning_examples_hygiene.ps1 -Fix"
    Skip = $SkipLearningMatrix -or -not $learningRootExists
    Run  = { & powershell -ExecutionPolicy Bypass -File scripts/validate_learning_examples_hygiene.ps1 -Fix }
  },
  [PSCustomObject]@{
    Name = "validate_learning_examples_hygiene.ps1"
    Skip = $SkipLearningMatrix -or -not $learningRootExists
    Run  = { & powershell -ExecutionPolicy Bypass -File scripts/validate_learning_examples_hygiene.ps1 }
  },
  [PSCustomObject]@{
    Name = "validate_migration_lock_contract.ps1"
    Skip = $false
    Run  = { & powershell -ExecutionPolicy Bypass -File scripts/validate_migration_lock_contract.ps1 }
  },
  [PSCustomObject]@{
    Name = "validate_runner_stop_behavior_contract.ps1"
    Skip = $false
    Run  = { & powershell -ExecutionPolicy Bypass -File scripts/validate_runner_stop_behavior_contract.ps1 }
  },
  [PSCustomObject]@{
    Name = "generate_rule_coverage_report.ps1"
    Skip = $false
    Run  = { & powershell -ExecutionPolicy Bypass -File scripts/generate_rule_coverage_report.ps1 }
  },
  [PSCustomObject]@{
    Name = "validate_parallel_repeatability.ps1"
    Skip = $SkipParallelRepeatability
    Run  = { & powershell -ExecutionPolicy Bypass -File scripts/validate_parallel_repeatability.ps1 -Iterations 6 -Command verify }
  },
  [PSCustomObject]@{
    Name = "benchmark_parallel_read.ps1"
    Skip = $SkipParallelBenchmark
    Run  = { & powershell -ExecutionPolicy Bypass -File scripts/benchmark_parallel_read.ps1 -Iterations 3 -Command gen }
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
  Write-Host "Final hardening gate failed on $($failed.Count) step(s):" -ForegroundColor Red
  foreach ($name in $failed) {
    Write-Host "- $name" -ForegroundColor Red
  }
  exit 1
}

Write-Host ""
Write-Host "Final hardening gate passed." -ForegroundColor Green
exit 0
