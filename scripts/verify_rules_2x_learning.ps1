param(
  [string]$RepoRoot = "."
)

$ErrorActionPreference = "Stop"
Set-Location $RepoRoot

function Has-Prop {
  param(
    [Parameter(Mandatory = $true)]$Obj,
    [Parameter(Mandatory = $true)][string]$Name
  )
  if ($null -eq $Obj) {
    return $false
  }
  return $Obj.PSObject.Properties.Name -contains $Name
}

function To-StringArray {
  param([object]$Value)
  if ($null -eq $Value) {
    return @()
  }
  return @($Value | ForEach-Object { [string]$_ } | Where-Object { $_.Trim() -ne "" })
}

function Resolve-ScenarioSteps {
  param([object]$Scenario)

  if ((Has-Prop -Obj $Scenario -Name "steps") -and $null -ne $Scenario.steps -and @($Scenario.steps).Count -gt 0) {
    return @($Scenario.steps)
  }

  $expectedExit = 0
  if (Has-Prop -Obj $Scenario -Name "expected") {
    $expected = [string]$Scenario.expected
    if ($expected -eq "FAIL") {
      $expectedExit = 1
    }
  }

  return @([PSCustomObject]@{
      command       = "verify"
      expected_exit = $expectedExit
    })
}

function Extract-Categories {
  param([string[]]$FileNames)
  $categories = @()
  foreach ($name in $FileNames) {
    if ($name -match "^\d{17}_([^_]+)_.+_v\d+\.(up|down)\.sql$") {
      $categories += $Matches[1].ToLowerInvariant()
    }
  }
  return @($categories | Sort-Object -Unique)
}

$matrixPath = "examples/rules_2x_learning/learning_matrix.json"
if (-not (Test-Path $matrixPath)) {
  Write-Error "matrix not found: $matrixPath"
  exit 1
}

$matrix = Get-Content -Raw $matrixPath | ConvertFrom-Json
$rows = @()
$errors = @()

if (Get-Variable -Name PSNativeCommandUseErrorActionPreference -ErrorAction SilentlyContinue) {
  $PSNativeCommandUseErrorActionPreference = $false
}

foreach ($item in $matrix) {
  $scenario = [string]$item.scenario
  $scenarioPath = [string]$item.path
  if (-not (Test-Path $scenarioPath)) {
    $errors += "[$scenario] scenario path not found: $scenarioPath"
    continue
  }

  $tmpRoot = Join-Path $env:TEMP ("rza_rules_learning_" + $scenario + "_" + [guid]::NewGuid().ToString("N"))
  New-Item -ItemType Directory -Path $tmpRoot | Out-Null
  $scenarioTemp = Join-Path $tmpRoot $scenario
  Copy-Item -Recurse -Path $scenarioPath -Destination $scenarioTemp

  $steps = Resolve-ScenarioSteps -Scenario $item
  $stepIndex = 0
  foreach ($step in $steps) {
    $command = ""
    if (Has-Prop -Obj $step -Name "command") {
      $command = [string]$step.command
    }
    if ($command -notin @("fmt", "gen", "verify")) {
      $errors += "[$scenario step#$($stepIndex+1)] invalid command '$command' (allowed: fmt|gen|verify)"
      $stepIndex++
      continue
    }

    $expectedExit = 0
    if (Has-Prop -Obj $step -Name "expected_exit") {
      $expectedExit = [int]$step.expected_exit
    }

    $prevErr = $ErrorActionPreference
    $ErrorActionPreference = "Continue"
    & go run -buildvcs=false ./cmd/rza_migrate_db $command -repo $scenarioTemp *> $null
    $ErrorActionPreference = $prevErr
    $actualExit = $LASTEXITCODE

    $stepErrors = @()
    if ($actualExit -ne $expectedExit) {
      $stepErrors += "exit mismatch expected=$expectedExit actual=$actualExit"
    }

    $outputPath = Join-Path $scenarioTemp "out/output.json"
    if (-not (Test-Path $outputPath)) {
      $stepErrors += "output json missing: $outputPath"
      $rows += [PSCustomObject]@{
        scenario      = $scenario
        step          = "$($stepIndex + 1).$command"
        expected_exit = $expectedExit
        actual_exit   = $actualExit
        match         = $false
      }
      foreach ($msg in $stepErrors) {
        $errors += "[$scenario step#$($stepIndex+1) $command] $msg"
      }
      $stepIndex++
      continue
    }

    $report = $null
    try {
      $report = Get-Content -Raw $outputPath | ConvertFrom-Json
    } catch {
      $stepErrors += "invalid output json: $outputPath"
    }

    if ($null -ne $report) {
      $actualFailRules = @()
      if ($null -ne $report.fail -and $null -ne $report.fail.items) {
        $actualFailRules = To-StringArray -Value @($report.fail.items | ForEach-Object { $_.rule })
      }

      $actualPassReasons = @()
      if ($null -ne $report.pass -and $null -ne $report.pass.items) {
        $actualPassReasons = To-StringArray -Value @($report.pass.items | ForEach-Object { $_.reason })
      }

      $expectedFailRules = @()
      if (Has-Prop -Obj $step -Name "expected_fail_rules") {
        $expectedFailRules = To-StringArray -Value $step.expected_fail_rules
      }
      foreach ($expectedRule in $expectedFailRules) {
        if (-not ($actualFailRules -contains $expectedRule)) {
          $stepErrors += "expected fail.rule missing: $expectedRule"
        }
      }

      $expectedPassReasons = @()
      if (Has-Prop -Obj $step -Name "expected_pass_reasons") {
        $expectedPassReasons = To-StringArray -Value $step.expected_pass_reasons
      }
      foreach ($expectedReason in $expectedPassReasons) {
        if (-not ($actualPassReasons -contains $expectedReason)) {
          $stepErrors += "expected pass.reason missing: $expectedReason"
        }
      }

      $generatedFiles = @()
      if ($null -ne $report.new_migrations -and $null -ne $report.new_migrations.files) {
        $generatedFiles = To-StringArray -Value $report.new_migrations.files
      }

      if (Has-Prop -Obj $step -Name "expected_generated_count") {
        $expectedGeneratedCount = [int]$step.expected_generated_count
        if ($generatedFiles.Count -ne $expectedGeneratedCount) {
          $stepErrors += "expected generated file count mismatch: expected=$expectedGeneratedCount actual=$($generatedFiles.Count)"
        }
      }

      $expectedGeneratedCategories = @()
      if (Has-Prop -Obj $step -Name "expected_generated_categories") {
        $expectedGeneratedCategories = To-StringArray -Value $step.expected_generated_categories
      }
      if ($expectedGeneratedCategories.Count -gt 0) {
        $actualCategories = Extract-Categories -FileNames $generatedFiles
        foreach ($expectedCategory in $expectedGeneratedCategories) {
          if (-not ($actualCategories -contains $expectedCategory)) {
            $stepErrors += "expected generated category missing: $expectedCategory"
          }
        }
      }

      $expectedNameContains = @()
      if (Has-Prop -Obj $step -Name "expected_generated_name_contains") {
        $expectedNameContains = To-StringArray -Value $step.expected_generated_name_contains
      }
      foreach ($needle in $expectedNameContains) {
        $found = $false
        foreach ($fileName in $generatedFiles) {
          if ($fileName.Contains($needle)) {
            $found = $true
            break
          }
        }
        if (-not $found) {
          $stepErrors += "expected generated filename pattern missing: $needle"
        }
      }

      $expectedContentContains = @()
      if (Has-Prop -Obj $step -Name "expected_generated_contains") {
        $expectedContentContains = To-StringArray -Value $step.expected_generated_contains
      }
      if ($expectedContentContains.Count -gt 0) {
        if ($generatedFiles.Count -eq 0) {
          $stepErrors += "expected generated content assertion exists but no generated files found"
        } else {
          $allGeneratedContent = ""
          foreach ($fileName in $generatedFiles) {
            $filePath = Join-Path $scenarioTemp ("migrations/" + $fileName)
            if (-not (Test-Path $filePath)) {
              $stepErrors += "generated file referenced in output not found: $fileName"
              continue
            }
            $allGeneratedContent += "`n" + (Get-Content -Raw $filePath)
          }

          foreach ($snippet in $expectedContentContains) {
            if (-not $allGeneratedContent.Contains($snippet)) {
              $stepErrors += "expected generated content missing: $snippet"
            }
          }
        }
      }
    }

    $rows += [PSCustomObject]@{
      scenario      = $scenario
      step          = "$($stepIndex + 1).$command"
      expected_exit = $expectedExit
      actual_exit   = $actualExit
      match         = ($stepErrors.Count -eq 0)
    }

    foreach ($msg in $stepErrors) {
      $errors += "[$scenario step#$($stepIndex+1) $command] $msg"
    }

    $stepIndex++
  }
}

$rows | Format-Table -AutoSize

if ($errors.Count -gt 0) {
  foreach ($msg in $errors) {
    Write-Host $msg -ForegroundColor Red
  }
  Write-Error "rules_2x_learning mismatch detected. Check matrix assertions and scenario contents."
  exit 1
}

Write-Output "rules_2x_learning verification matched matrix."
exit 0
