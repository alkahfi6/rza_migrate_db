$ErrorActionPreference = "Stop"

$root = Join-Path (Get-Location) "fixtures/qa_rules"

if (-not (Test-Path $root)) {
  throw "Missing fixtures root: $root"
}

$scenarios = @(
  "01_hard_fail_rule",
  "02_folder_contract",
  "03_forbidden_sql",
  "04_pairing_missing_down",
  "05_version_format_invalid",
  "06_ordering_determinism",
  "07_vk_whitespace_only",
  "08_reconstruction_unsupported_sql",
  "09_file_object_mismatch",
  "10_inline_default_not_allowed",
  "11_destructive_drop_table_manual",
  "12_transaction_error_throw",
  "13_audit_wrapper_contract",
  "14_down_data_restore_warning",
  "15_patch_canonicalization",
  "16_patch_duplication_once",
  "17_patch_source_integrity",
  "18_output_json_contract",
  "19_repo_hygiene_constraints",
  "20_cli_exit_code_contract",
  "21_wrapper_action_whitelist",
  "22_wrapper_reason_whitelist",
  "23_wrapper_run_key_newid",
  "24_constraint_unnamed",
  "25_wrapper_precheck_skip",
  "26_alter_column_precheck_map"
)

$errors = @()

function Has-Prop {
  param(
    [Parameter(Mandatory = $true)]$Obj,
    [Parameter(Mandatory = $true)][string]$Name
  )
  return $Obj.PSObject.Properties.Name -contains $Name
}

foreach ($scenario in $scenarios) {
  $base = Join-Path $root $scenario
  $readme = Join-Path $base "README.md"
  $input = Join-Path $base "input"
  $expected = Join-Path $base "expected"
  $expectJson = Join-Path $expected "expectation.json"
  $goldenJson = Join-Path $expected "output.golden.json"

  if (-not (Test-Path $base)) { $errors += "Missing scenario dir: $base"; continue }
  if (-not (Test-Path $readme)) { $errors += "Missing README: $readme" }
  if (-not (Test-Path $input)) { $errors += "Missing input dir: $input" }
  if (-not (Test-Path $expected)) { $errors += "Missing expected dir: $expected" }
  if (-not (Test-Path $expectJson)) { $errors += "Missing expectation contract: $expectJson" }
  if (-not (Test-Path $goldenJson)) { $errors += "Missing output golden: $goldenJson" }

  $inputFileCount = (Get-ChildItem $input -Recurse -File -ErrorAction SilentlyContinue | Measure-Object).Count
  if ($inputFileCount -lt 1) { $errors += "Input fixture has no files: $input" }

  $expectObj = $null
  if (Test-Path $expectJson) {
    try {
      $obj = Get-Content $expectJson -Raw | ConvertFrom-Json
      $expectObj = $obj
      foreach ($prop in @("qa_id", "command", "expected_exit_code", "expected_fail_rules", "expected_warn_reasons", "summary")) {
        if (-not (Has-Prop -Obj $obj -Name $prop)) {
          $errors += "Expectation missing property '$prop': $expectJson"
        }
      }
    } catch {
      $errors += "Invalid JSON in expectation: $expectJson"
    }
  }

  if (Test-Path $goldenJson) {
    try {
      $gold = Get-Content $goldenJson -Raw | ConvertFrom-Json
      foreach ($rootProp in @("fail", "pass", "new_migrations", "notes")) {
        if (-not (Has-Prop -Obj $gold -Name $rootProp)) {
          $errors += "Golden missing root property '$rootProp': $goldenJson"
        }
      }

      if (Has-Prop -Obj $gold -Name "fail") {
        if (-not (Has-Prop -Obj $gold.fail -Name "count")) { $errors += "Golden fail.count missing: $goldenJson" }
        if (-not (Has-Prop -Obj $gold.fail -Name "items")) { $errors += "Golden fail.items missing: $goldenJson" }
      }
      if (Has-Prop -Obj $gold -Name "pass") {
        if (-not (Has-Prop -Obj $gold.pass -Name "count")) { $errors += "Golden pass.count missing: $goldenJson" }
        if (-not (Has-Prop -Obj $gold.pass -Name "items")) { $errors += "Golden pass.items missing: $goldenJson" }
      }
      if (Has-Prop -Obj $gold -Name "new_migrations") {
        if (-not (Has-Prop -Obj $gold.new_migrations -Name "naming")) { $errors += "Golden new_migrations.naming missing: $goldenJson" }
        if (-not (Has-Prop -Obj $gold.new_migrations -Name "files")) { $errors += "Golden new_migrations.files missing: $goldenJson" }
        if (Has-Prop -Obj $gold.new_migrations -Name "naming") {
          if (-not (Has-Prop -Obj $gold.new_migrations.naming -Name "pattern")) { $errors += "Golden naming.pattern missing: $goldenJson" }
          if (-not (Has-Prop -Obj $gold.new_migrations.naming -Name "line_endings")) { $errors += "Golden naming.line_endings missing: $goldenJson" }
          if ((Has-Prop -Obj $gold.new_migrations.naming -Name "line_endings") -and ($gold.new_migrations.naming.line_endings -ne "LF")) {
            $errors += "Golden naming.line_endings must be LF: $goldenJson"
          }
        }
      }

      $failItems = @()
      if ($null -ne $gold.fail.items) { $failItems = @($gold.fail.items) }
      $passItems = @()
      if ($null -ne $gold.pass.items) { $passItems = @($gold.pass.items) }

      if ([int]$gold.fail.count -ne $failItems.Count) {
        $errors += "Golden fail.count mismatch with fail.items length: $goldenJson"
      }
      if ([int]$gold.pass.count -ne $passItems.Count) {
        $errors += "Golden pass.count mismatch with pass.items length: $goldenJson"
      }

      if ($null -ne $expectObj) {
        $expectedFailRules = @()
        if ($null -ne $expectObj.expected_fail_rules) {
          $expectedFailRules = @($expectObj.expected_fail_rules | ForEach-Object { [string]$_ })
        }
        $expectedWarnReasons = @()
        if ($null -ne $expectObj.expected_warn_reasons) {
          $expectedWarnReasons = @($expectObj.expected_warn_reasons | ForEach-Object { [string]$_ })
        }

        $actualFailRules = @($failItems | ForEach-Object { [string]$_.rule })
        foreach ($expectedRule in $expectedFailRules) {
          if (-not ($actualFailRules -contains $expectedRule)) {
            $errors += "Expected fail rule '$expectedRule' missing in golden: $goldenJson"
          }
        }

        $actualWarnReasons = @($passItems | ForEach-Object { [string]$_.reason })
        foreach ($expectedReason in $expectedWarnReasons) {
          if (-not ($actualWarnReasons -contains $expectedReason)) {
            $errors += "Expected warn reason '$expectedReason' missing in golden: $goldenJson"
          }
        }

        if ((Has-Prop -Obj $gold -Name "notes") -and (Has-Prop -Obj $gold.notes -Name "qa_id")) {
          if ([string]$gold.notes.qa_id -ne [string]$expectObj.qa_id) {
            $errors += "Golden notes.qa_id mismatch: $goldenJson"
          }
        } else {
          $errors += "Golden notes.qa_id missing: $goldenJson"
        }

        if (([int]$expectObj.expected_exit_code -eq 1) -and ([int]$gold.fail.count -eq 0)) {
          $errors += "Expected exit_code=1 but golden fail.count=0: $goldenJson"
        }
        if (([int]$expectObj.expected_exit_code -eq 0) -and ([int]$gold.fail.count -gt 0)) {
          $errors += "Expected exit_code=0 but golden fail.count>0: $goldenJson"
        }
      }
    } catch {
      $errors += "Invalid JSON in golden: $goldenJson"
    }
  }
}

if ($errors.Count -gt 0) {
  $errors | ForEach-Object { Write-Host $_ -ForegroundColor Red }
  throw "Fixture validation failed with $($errors.Count) error(s)."
}

Write-Host "Fixture contract validation passed." -ForegroundColor Green
