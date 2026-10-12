$ErrorActionPreference = "Stop"

$root = Join-Path (Get-Location) "fixtures/qa_rules"
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

if (-not (Test-Path $root)) {
  throw "Missing fixtures root: $root"
}

$pattern = "YYYYMMDDHHMMSS<seq>_<category>_<object>_v<k>.(up|down).sql"

$scenarioDirs = Get-ChildItem $root -Directory | Sort-Object Name

foreach ($scenario in $scenarioDirs) {
  Write-Host "Generating golden output for $($scenario.Name)..."
  $expectedDir = Join-Path $scenario.FullName "expected"
  $expectPath = Join-Path $expectedDir "expectation.json"
  if (-not (Test-Path $expectPath)) { continue }

  $expect = Get-Content $expectPath -Raw | ConvertFrom-Json

  $failRules = @()
  if ($null -ne $expect.expected_fail_rules) {
    $failRules = @($expect.expected_fail_rules)
  }
  $failItems = @()
  foreach ($rule in $failRules) {
    $reasonText = "SQL_ERROR"
    if ($rule -eq "INLINE_DEFAULT_NOT_ALLOWED") {
      $reasonText = "INLINE_DEFAULT_NOT_ALLOWED"
    }
    $failItems += [ordered]@{
      rule = [string]$rule
      severity = 3
      file = "input/fixture.sql"
      message = "Expected fail rule triggered: $rule"
      reason = $reasonText
      action = "FAIL"
    }
  }

  $warnReasons = @()
  if ($null -ne $expect.expected_warn_reasons) {
    $warnReasons = @($expect.expected_warn_reasons)
  }
  $passItems = @()
  foreach ($reason in $warnReasons) {
    $passItems += [ordered]@{
      rule = "EXPECTED_WARN_REASON"
      severity = 2
      file = "input/fixture.sql"
      message = "Expected warn reason: $reason"
      reason = [string]$reason
      action = "warn"
    }
  }

  if (($passItems.Count -eq 0) -and ($expect.expected_exit_code -eq 0)) {
    $passItems += [ordered]@{
      rule = "SCENARIO_BASELINE"
      severity = 0
      file = "input"
      message = "Scenario baseline pass for $($expect.qa_id)"
      reason = "CREATED"
      action = "ok"
    }
  }

  $newFiles = @()
  $orderPath = Join-Path $expectedDir "generated_order.txt"
  $zeroPath = Join-Path $expectedDir "expected_new_migrations.txt"
  $patchNamePath = Join-Path $expectedDir "expected_patch_name.txt"

  if (Test-Path $orderPath) {
    $newFiles = @(
      Get-Content $orderPath |
      ForEach-Object { [string]$_ } |
      Where-Object { $_.Trim() -ne "" }
    )
  } elseif ((Test-Path $zeroPath) -and ((Get-Content $zeroPath -Raw).Trim() -eq "0")) {
    $newFiles = @()
  } elseif (Test-Path $patchNamePath) {
    $patchName = (Get-Content $patchNamePath -Raw).Trim()
    if ($patchName -ne "") {
      $newFiles = @(
        "<ts>000_patch_${patchName}_v1.up.sql",
        "<ts>000_patch_${patchName}_v1.down.sql"
      )
    }
  }

  $golden = [ordered]@{
    fail = [ordered]@{
      count = $failItems.Count
      items = $failItems
    }
    pass = [ordered]@{
      count = $passItems.Count
      items = $passItems
    }
    new_migrations = [ordered]@{
      naming = [ordered]@{
        pattern = $pattern
        line_endings = "LF"
      }
      files = $newFiles
    }
    notes = [ordered]@{
      qa_id = [string]$expect.qa_id
      command = [string]$expect.command
      expected_exit_code = [int]$expect.expected_exit_code
      summary = [string]$expect.summary
    }
  }

  $json = $golden | ConvertTo-Json -Depth 12
  Write-LFFile -Path (Join-Path $expectedDir "output.golden.json") -Content ($json + "`n")
}
