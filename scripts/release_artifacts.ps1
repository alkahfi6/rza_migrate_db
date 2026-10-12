param(
    [Parameter(Mandatory = $true)]
    [ValidateSet("Prepare", "Finalize", "Publish")]
    [string]$Mode,
    [Parameter(Mandatory = $true)]
    [string]$RepoRoot,
    [Parameter(Mandatory = $true)]
    [string]$StageRoot,
    [string]$OutputRoot
)

Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"

$RepoRoot = [IO.Path]::GetFullPath($RepoRoot).TrimEnd('\', '/')
$StageRoot = [IO.Path]::GetFullPath($StageRoot).TrimEnd('\', '/')

function Get-SHA256Hex([string]$Path) {
    $sha = [Security.Cryptography.SHA256]::Create()
    try {
        $stream = [IO.File]::OpenRead($Path)
        try { $bytes = $sha.ComputeHash($stream) } finally { $stream.Dispose() }
    } finally {
        $sha.Dispose()
    }
    return ([BitConverter]::ToString($bytes)).Replace("-", "").ToLowerInvariant()
}

function Get-NormalizedSourceSHA256Hex([string]$Path) {
    # Git may materialize .bat files as CRLF on Windows and LF in Linux. Source
    # identity is intentionally independent of that checkout-only difference.
    $text = [IO.File]::ReadAllText($Path, [Text.Encoding]::UTF8).Replace("`r`n", "`n")
    $sha = [Security.Cryptography.SHA256]::Create()
    try { $bytes = $sha.ComputeHash([Text.Encoding]::UTF8.GetBytes($text)) } finally { $sha.Dispose() }
    return ([BitConverter]::ToString($bytes)).Replace("-", "").ToLowerInvariant()
}

function Write-UTF8([string]$Path, [string]$Text) {
    $parent = [IO.Path]::GetDirectoryName($Path)
    if ($parent) { [IO.Directory]::CreateDirectory($parent) | Out-Null }
    [IO.File]::WriteAllText($Path, $Text, [Text.UTF8Encoding]::new($false))
}

function Get-SourceFiles {
    $paths = [Collections.Generic.List[string]]::new()
    foreach ($relativeRoot in @('cmd\rza_migrate_db', 'cmd\rza_migrate_check', 'cmd\rza_migrate_report', 'internal')) {
        $root = Join-Path $RepoRoot $relativeRoot
        Get-ChildItem -LiteralPath $root -Recurse -File -Filter '*.go' | ForEach-Object { $paths.Add($_.FullName) }
    }
    foreach ($relative in @('go.mod', 'go.sum', 'build_all.bat', 'scripts\release_artifacts.ps1', 'scripts\verify_release_source.sh', '.gitlab\rza_migrate_check.gitlab-ci.yml')) {
        $path = Join-Path $RepoRoot $relative
        if (-not [IO.File]::Exists($path)) { throw "release source input is missing: $relative" }
        $paths.Add($path)
    }
    return $paths
}

function Get-SourceDigest {
    $lines = [Text.StringBuilder]::new()
    [string[]]$sourceFiles = @(Get-SourceFiles)
    [Array]::Sort($sourceFiles, [StringComparer]::Ordinal)
    foreach ($path in $sourceFiles) {
        $relative = $path.Substring($RepoRoot.Length + 1).Replace('\', '/')
        [void]$lines.Append((Get-NormalizedSourceSHA256Hex $path)).Append('  ').Append($relative).Append("`n")
    }
    $sha = [Security.Cryptography.SHA256]::Create()
    try { $digest = $sha.ComputeHash([Text.Encoding]::UTF8.GetBytes($lines.ToString())) } finally { $sha.Dispose() }
    return ([BitConverter]::ToString($digest)).Replace("-", "").ToLowerInvariant()
}

$metadataPath = Join-Path $StageRoot 'build.env'

if ($Mode -eq 'Prepare') {
    [IO.Directory]::CreateDirectory($StageRoot) | Out-Null
    $now = [DateTime]::UtcNow
    $version = $now.ToString('yyyyMMdd.HHmmss')
    $buildTime = $now.ToString('yyyy-MM-ddTHH:mm:ssZ')
    $sourceDigest = Get-SourceDigest
    Write-UTF8 $metadataPath "RZA_BUILD_VERSION=$version`nRZA_BUILD_TIME=$buildTime`nRZA_SOURCE_SHA256=$sourceDigest`n"
    exit 0
}

if (-not [IO.File]::Exists($metadataPath)) { throw "build metadata is missing: $metadataPath" }
$metadata = @{}
foreach ($line in [IO.File]::ReadAllLines($metadataPath)) {
    $parts = $line.Split('=', 2)
    if ($parts.Count -eq 2) { $metadata[$parts[0]] = $parts[1] }
}
foreach ($key in @('RZA_BUILD_VERSION', 'RZA_BUILD_TIME', 'RZA_SOURCE_SHA256')) {
    if (-not $metadata.ContainsKey($key) -or [string]::IsNullOrWhiteSpace($metadata[$key])) { throw "build metadata $key is missing" }
}
if ($metadata.RZA_SOURCE_SHA256 -notmatch '^[0-9a-f]{64}$') { throw 'source digest is invalid' }

$payloads = @(
    'tools/rza_migrate_db.exe',
    'tools/rza_migrate_check.exe',
    'tools/rza_migrate_report.exe',
    'tools/linux-amd64/rza_migrate_db',
    'tools/linux-amd64/rza_migrate_check',
    'tools/linux-amd64/rza_migrate_report',
    'tools/license.key'
)

if ($Mode -eq 'Finalize') {
    foreach ($relative in $payloads) {
        $path = Join-Path $StageRoot ($relative.Replace('/', '\'))
        if (-not [IO.File]::Exists($path) -or ([IO.FileInfo]$path).Length -eq 0) { throw "release payload is missing or empty: $relative" }
    }
    $licenseHeader = [IO.File]::ReadLines((Join-Path $StageRoot 'tools\license.key')) | Select-Object -First 1
    if ($licenseHeader -ne 'RZA_LICENSE_V1') {
        throw 'generated license has an invalid header'
    }

    $artifacts = [Collections.Generic.List[object]]::new()
    foreach ($relative in $payloads) {
        $path = Join-Path $StageRoot ($relative.Replace('/', '\'))
        $hash = Get-SHA256Hex $path
        if ($hash -notmatch '^[0-9a-f]{64}$') { throw "invalid SHA-256 for $relative" }
        $artifacts.Add([ordered]@{ path = $relative; sha256 = $hash; size = ([IO.FileInfo]$path).Length })
        if ($relative -ne 'tools/license.key') { Write-UTF8 ($path + '.sha256') "$hash  $relative`n" }
    }

    $buildInfoRelative = 'tools/linux-amd64/build-info.txt'
    $buildInfo = "version=$($metadata.RZA_BUILD_VERSION)`nbuild_time=$($metadata.RZA_BUILD_TIME)`nsource_sha256=$($metadata.RZA_SOURCE_SHA256)`n"
    Write-UTF8 (Join-Path $StageRoot ($buildInfoRelative.Replace('/', '\'))) $buildInfo

    $manifestRelative = 'tools/linux-amd64/release-manifest.json'
    $manifest = [ordered]@{
        schema = 'RZA_RELEASE_V1'
        version = $metadata.RZA_BUILD_VERSION
        build_time = $metadata.RZA_BUILD_TIME
        source_sha256 = $metadata.RZA_SOURCE_SHA256
        artifacts = $artifacts
    } | ConvertTo-Json -Depth 5
    Write-UTF8 (Join-Path $StageRoot ($manifestRelative.Replace('/', '\'))) ($manifest + "`n")

    $releaseEntries = [Collections.Generic.List[string]]::new()
    foreach ($relative in $payloads) {
        $path = Join-Path $StageRoot ($relative.Replace('/', '\'))
        $releaseEntries.Add("$(Get-SHA256Hex $path)  $relative")
        if ($relative -ne 'tools/license.key') {
            $shaRelative = $relative + '.sha256'
            $releaseEntries.Add("$(Get-SHA256Hex ($path + '.sha256'))  $shaRelative")
        }
    }
    foreach ($relative in @($buildInfoRelative, $manifestRelative)) {
        $releaseEntries.Add("$(Get-SHA256Hex (Join-Path $StageRoot ($relative.Replace('/', '\'))))  $relative")
    }
    Write-UTF8 (Join-Path $StageRoot 'tools\linux-amd64\release.sha256') (($releaseEntries -join "`n") + "`n")
    exit 0
}

if ([string]::IsNullOrWhiteSpace($OutputRoot)) { throw 'OutputRoot is required for Publish' }
$OutputRoot = [IO.Path]::GetFullPath($OutputRoot).TrimEnd('\', '/')
$releaseList = Join-Path $StageRoot 'tools\linux-amd64\release.sha256'
if (-not [IO.File]::Exists($releaseList)) { throw 'release checksum manifest is missing' }

$publishFiles = [Collections.Generic.List[string]]::new()
foreach ($line in [IO.File]::ReadAllLines($releaseList)) {
    if ($line -notmatch '^([0-9a-f]{64})  (.+)$') { throw "invalid release checksum entry: $line" }
    $relative = $Matches[2]
    if ([IO.Path]::IsPathRooted($relative) -or $relative -match '(^|/|\\)\.\.($|/|\\)') {
        throw "unsafe release path: $relative"
    }
    $source = Join-Path $StageRoot ($relative.Replace('/', '\'))
    if ((Get-SHA256Hex $source) -ne $Matches[1]) { throw "staged checksum mismatch: $relative" }
    $publishFiles.Add($relative)
}
$publishFiles.Add('tools/linux-amd64/release.sha256')

$published = [Collections.Generic.List[object]]::new()
try {
    foreach ($relative in $publishFiles) {
        $source = Join-Path $StageRoot ($relative.Replace('/', '\'))
        $destination = Join-Path $OutputRoot ($relative.Replace('/', '\'))
        [IO.Directory]::CreateDirectory([IO.Path]::GetDirectoryName($destination)) | Out-Null
        $existed = [IO.File]::Exists($destination)
        $backup = ''
        if ($existed) { $backup = $destination + '.rza-backup-' + [Guid]::NewGuid().ToString('N') + '.tmp' }
        $temporary = $destination + '.rza-publish-' + [Guid]::NewGuid().ToString('N') + '.tmp'
        [IO.File]::Copy($source, $temporary, $true)
        if ((Get-SHA256Hex $temporary) -ne (Get-SHA256Hex $source)) { throw "publish copy verification failed: $relative" }
        if ($existed) { [IO.File]::Replace($temporary, $destination, $backup, $true) }
        else { [IO.File]::Move($temporary, $destination) }
        $published.Add([pscustomobject]@{ Destination = $destination; Existed = $existed; Backup = $backup })
    }
} catch {
    for ($index = $published.Count - 1; $index -ge 0; $index--) {
        $item = $published[$index]
        if ($item.Existed) {
            [IO.File]::Copy($item.Backup, $item.Destination, $true)
            [IO.File]::Delete($item.Backup)
        }
        elseif ([IO.File]::Exists($item.Destination)) { [IO.File]::Delete($item.Destination) }
    }
    throw
}
foreach ($item in $published) {
    if ($item.Existed -and [IO.File]::Exists($item.Backup)) { [IO.File]::Delete($item.Backup) }
}
exit 0
