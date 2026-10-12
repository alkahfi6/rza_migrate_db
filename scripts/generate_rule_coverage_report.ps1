param(
  [string]$RepoRoot = ".",
  [string]$OutputPath = "docs/qa/qa_coverage_report.md"
)

$ErrorActionPreference = "Stop"

Set-Location $RepoRoot

$enc = New-Object System.Text.UTF8Encoding($false)

function Write-LFFile {
  param(
    [string]$Path,
    [string]$Content
  )
  $dir = Split-Path -Parent $Path
  if ($dir -and -not (Test-Path $dir)) {
    New-Item -ItemType Directory -Force -Path $dir | Out-Null
  }
  $normalized = $Content -replace "`r`n", "`n"
  [System.IO.File]::WriteAllText($Path, $normalized, $enc)
}

function Join-ListOrDash {
  param([object[]]$Items)
  if ($null -eq $Items) {
    return "-"
  }
  $list = @($Items | ForEach-Object { [string]$_ } | Where-Object { $_.Trim() -ne "" })
  if ($list.Count -eq 0) {
    return "-"
  }
  return ($list -join ", ")
}

function Read-ResidualRiskLines {
  param([string]$MatrixPath)

  if (-not (Test-Path $MatrixPath)) {
    return @()
  }

  $lines = Get-Content $MatrixPath
  $start = -1
  for ($i = 0; $i -lt $lines.Count; $i++) {
    if ($lines[$i].Trim() -eq "Residual risk (tetap berlaku):") {
      $start = $i + 1
      break
    }
  }
  if ($start -lt 0) {
    return @()
  }

  $out = @()
  for ($i = $start; $i -lt $lines.Count; $i++) {
    $line = $lines[$i].Trim()
    if ($line -eq "") {
      if ($out.Count -gt 0) {
        break
      }
      continue
    }
    if ($line.StartsWith("- ")) {
      $out += $line.Substring(2).Trim()
      continue
    }
    if ($out.Count -gt 0) {
      break
    }
  }
  return $out
}

function Read-RuleRowsFromMatrix {
  param([string]$MatrixPath)

  if (-not (Test-Path $MatrixPath)) {
    return @()
  }

  $lines = Get-Content $MatrixPath
  $rows = @()
  foreach ($line in $lines) {
    $trimmed = $line.Trim()
    if (-not $trimmed.StartsWith("|")) {
      continue
    }
    if ($trimmed -match "^\|\s*Rule No\s*\|") {
      continue
    }
    if ($trimmed -match "^\|\s*---") {
      continue
    }

    $parts = @($trimmed.Split("|") | ForEach-Object { $_.Trim() } | Where-Object { $_ -ne "" })
    if ($parts.Count -lt 6) {
      continue
    }
    if ($parts[0] -notmatch "^\d+$") {
      continue
    }

    $rows += [PSCustomObject]@{
      RuleNo      = [int]$parts[0]
      RuleSection = [string]$parts[1]
      ControlGate = [string]$parts[5]
    }
  }

  return @($rows | Sort-Object RuleNo)
}

$fixtureRoot = "fixtures/qa_rules"
$matrixPath = "docs/traceability/rules_traceability_matrix.md"

if (-not (Test-Path $fixtureRoot)) {
  throw "Missing fixtures root: $fixtureRoot"
}

$scenarioDirs = Get-ChildItem $fixtureRoot -Directory |
  Where-Object { $_.Name -match "^\d{2}_" } |
  Sort-Object Name

$scenarioRows = @()
foreach ($dir in $scenarioDirs) {
  $expectPath = Join-Path $dir.FullName "expected/expectation.json"
  if (-not (Test-Path $expectPath)) {
    throw "Missing expectation file: $expectPath"
  }
  $expect = Get-Content $expectPath -Raw | ConvertFrom-Json

  $expectedOutcome = "pass"
  if ([int]$expect.expected_exit_code -eq 1) {
    $expectedOutcome = "fail"
  } elseif (@($expect.expected_warn_reasons).Count -gt 0) {
    $expectedOutcome = "pass-with-warn"
  }

  $scenarioRows += [PSCustomObject]@{
    QAID       = [string]$expect.qa_id
    Scenario   = [string]$dir.Name
    Command    = [string]$expect.command
    Outcome    = $expectedOutcome
    FailRules  = Join-ListOrDash -Items @($expect.expected_fail_rules)
    WarnReason = Join-ListOrDash -Items @($expect.expected_warn_reasons)
  }
}

$ruleRows = Read-RuleRowsFromMatrix -MatrixPath $matrixPath
$residualRisks = Read-ResidualRiskLines -MatrixPath $matrixPath

$totalScenarios = $scenarioRows.Count
$failScenarios = @($scenarioRows | Where-Object { $_.Outcome -eq "fail" }).Count
$warnScenarios = @($scenarioRows | Where-Object { $_.Outcome -eq "pass-with-warn" }).Count
$passScenarios = @($scenarioRows | Where-Object { $_.Outcome -eq "pass" }).Count
$totalRules = $ruleRows.Count
$generatedAtUTC = (Get-Date).ToUniversalTime().ToString("yyyy-MM-dd HH:mm:ss")

$lines = @()
$lines += "# QA Coverage Report"
$lines += ""
$lines += "Generated at (UTC): $generatedAtUTC"
$lines += ""
$lines += "## Summary"
$lines += ""
$lines += "- Total QA scenarios: $totalScenarios"
$lines += "- Expected pass: $passScenarios"
$lines += "- Expected pass-with-warn: $warnScenarios"
$lines += "- Expected fail: $failScenarios"
$lines += "- Total mapped rules in traceability matrix: $totalRules"
$lines += ""
$lines += "## QA Scenario Coverage"
$lines += ""
$lines += "| QA | Scenario | Command | Expected | Expected Fail Rules | Expected Warn Reasons |"
$lines += "| --- | --- | --- | --- | --- | --- |"
foreach ($row in $scenarioRows) {
  $lines += "| $($row.QAID) | ``$($row.Scenario)`` | ``$($row.Command)`` | $($row.Outcome) | $($row.FailRules) | $($row.WarnReason) |"
}

$lines += ""
$lines += "## Rule Mapping Coverage"
$lines += ""
$lines += "| Rule No | Rule Section | Control/Gate | Status |"
$lines += "| --- | --- | --- | --- |"
foreach ($row in $ruleRows) {
  $status = "Mapped"
  if ($row.ControlGate.Trim() -eq "") {
    $status = "Needs Update"
  }
  $lines += "| $($row.RuleNo) | $($row.RuleSection) | $($row.ControlGate) | $status |"
}

$lines += ""
$lines += "## Residual Risk Checklist"
$lines += ""
if ($residualRisks.Count -eq 0) {
  $lines += "- Tidak ada residual risk yang terdeteksi dari matrix."
} else {
  foreach ($risk in $residualRisks) {
    $lines += "- [ ] $risk"
  }
}
$lines += ""
$lines += "## Operational Commands"
$lines += ""
$lines += "1. ``go test ./... -count=1``"
$lines += "2. ``powershell -ExecutionPolicy Bypass -File scripts/validate_qa_fixture_contract.ps1``"
$lines += "3. ``powershell -ExecutionPolicy Bypass -File scripts/verify_rules_2x_learning.ps1``"
$lines += "4. ``powershell -ExecutionPolicy Bypass -File scripts/validate_learning_examples_hygiene.ps1``"
$lines += "5. ``powershell -ExecutionPolicy Bypass -File scripts/validate_migration_lock_contract.ps1``"
$lines += "6. ``powershell -ExecutionPolicy Bypass -File scripts/validate_runner_stop_behavior_contract.ps1``"
$lines += "7. ``powershell -ExecutionPolicy Bypass -File scripts/generate_rule_coverage_report.ps1``"

Write-LFFile -Path $OutputPath -Content (($lines -join "`n") + "`n")
Write-Host "Coverage report generated: $OutputPath" -ForegroundColor Green
