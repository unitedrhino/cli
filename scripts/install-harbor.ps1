# =============================================================================
# ur CLI Windows one-shot install/upgrade script (provided by the official
# doc site: irm https://doc.unitedrhino.com/cli/install.ps1 | iex)
#
# English-only output on purpose: Windows PowerShell 5.1 decodes BOM-less
# UTF-8 scripts as ANSI, which garbles non-ASCII strings.
#
# Auto: query latest Harbor public artifact -> anonymous token -> download
# release zip + sha256 -> verify -> install to %LOCALAPPDATA%\Programs\ur
# (ur.exe + skill/ side by side) -> add to user PATH.
#
# Usage (PowerShell 5.1+, no admin needed):
#   irm https://doc.unitedrhino.com/cli/install.ps1 | iex
#
# Env (optional):
#   $env:UR_DEST   install dir (default $env:LOCALAPPDATA\Programs\ur)
#   $env:UR_TAG    pin a version (default: latest vX.Y.Z from Harbor)
# =============================================================================
$ErrorActionPreference = 'Stop'
$ProgressPreference = 'SilentlyContinue'
try { [Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12 } catch {}

$Registry = 'docker.unitedrhino.com'
$Project  = 'urops'
$Repo     = 'ur-cli-windows-amd64'
$Dest     = if ($env:UR_DEST) { $env:UR_DEST } else { Join-Path $env:LOCALAPPDATA 'Programs\ur' }
$ExePath  = Join-Path $Dest 'ur.exe'

function Write-Ok($msg)  { Write-Host "OK  $msg" }
function Write-Info($msg) { Write-Host "==> $msg" }
function Write-Fail($msg) { Write-Host "FAIL: $msg" -ForegroundColor Red; exit 1 }

# 1. Latest version (artifacts API is publicly readable; first vX.Y.Z tag by push time desc)
$Tag = $env:UR_TAG
if (-not $Tag) {
    try {
        $artifacts = Invoke-RestMethod -Method Get -Uri "https://${Registry}/api/v2.0/projects/${Project}/repositories/${Repo}/artifacts?page_size=50&with_tag=true&sort=push_time:desc" -TimeoutSec 15
        foreach ($a in $artifacts) {
            foreach ($t in $a.tags) {
                if ($t.name -match '^v\d+\.\d+\.\d+$') { $Tag = $t.name; break }
            }
            if ($Tag) { break }
        }
    } catch {}
    if (-not $Tag) { Write-Fail "cannot fetch latest version (check network to ${Registry})" }
}
Write-Info "Installing ur CLI $Tag"

# 2. Anonymous token (parse token service from the registry /v2/ 401 challenge)
try {
    $resp = Invoke-WebRequest -Method Get -Uri "https://${Registry}/v2/" -UseBasicParsing -TimeoutSec 15
    Write-Fail "registry did not return the expected 401 challenge"
} catch {
    $wwwAuth = $_.Exception.Response.Headers['WWW-Authenticate']
    if (-not $wwwAuth) { Write-Fail "failed to get anonymous token: no WWW-Authenticate" }
}
$realm   = [regex]::Match($wwwAuth, 'realm="([^"]+)"').Groups[1].Value
$service = [regex]::Match($wwwAuth, 'service="([^"]+)"').Groups[1].Value
if (-not $realm -or -not $service) { Write-Fail "failed to parse token service address" }
try {
    $tokenResp = Invoke-RestMethod -Method Get -Uri "${realm}?service=${service}&scope=repository:${Project}/${Repo}:pull" -TimeoutSec 15
    $Token = $tokenResp.token
} catch {}
if (-not $Token) { Write-Fail "failed to get anonymous token" }

# 3. Manifest layer digests (layer 0 = release zip, layer 1 = sha256 file)
try {
    $manifest = Invoke-RestMethod -Method Get -Uri "https://${Registry}/v2/${Project}/${Repo}/manifests/${Tag}" `
        -Headers @{ Authorization = "Bearer $Token"; Accept = 'application/vnd.oci.image.manifest.v1+json, application/vnd.docker.distribution.manifest.v2+json' } -TimeoutSec 15
    $ZipDigest = $manifest.layers[0].digest
    $ShaDigest = $manifest.layers[1].digest
} catch { Write-Fail "failed to get artifact manifest: $($_.Exception.Message)" }
if (-not $ZipDigest -or -not $ShaDigest) { Write-Fail "manifest has no artifact layers" }

# 4. Download + verify
$Work = Join-Path ([IO.Path]::GetTempPath()) ("ur-install-" + [guid]::NewGuid().ToString('N').Substring(0, 8))
New-Item -ItemType Directory -Path $Work -Force | Out-Null
$ZipPath  = Join-Path $Work 'ur-cli.zip'
$ShaPath  = Join-Path $Work 'sha256.txt'
try {
    Invoke-WebRequest -Method Get -UseBasicParsing -TimeoutSec 600 `
        -Headers @{ Authorization = "Bearer $Token" } `
        -Uri "https://${Registry}/v2/${Project}/${Repo}/blobs/${ZipDigest}" -OutFile $ZipPath
    Invoke-WebRequest -Method Get -UseBasicParsing -TimeoutSec 60 `
        -Headers @{ Authorization = "Bearer $Token" } `
        -Uri "https://${Registry}/v2/${Project}/${Repo}/blobs/${ShaDigest}" -OutFile $ShaPath
    $Expected = ((Get-Content $ShaPath -Raw).Trim() -split '\s+')[0].ToLower()
    $Actual   = (Get-FileHash $ZipPath -Algorithm SHA256).Hash.ToLower()
    if ($Expected -ne $Actual) { Write-Fail "sha256 mismatch: expected $Expected, got $Actual" }
    Write-Ok "sha256 verified"
} catch { Write-Fail "download/verify failed: $($_.Exception.Message)" }

# 5. Install (expand zip -> locate ur.exe -> move ur.exe + skill/ side by side)
New-Item -ItemType Directory -Path $Dest -Force | Out-Null
$Extract = Join-Path $Work 'extracted'
Expand-Archive -Path $ZipPath -DestinationPath $Extract -Force
$ExeSource = Get-ChildItem -Path $Extract -Recurse -Filter 'ur.exe' | Select-Object -First 1
if (-not $ExeSource) { Write-Fail "ur.exe not found inside the release zip" }
$SrcRoot = $ExeSource.Directory.FullName
Copy-Item -Force $ExeSource.FullName $ExePath
if (Test-Path (Join-Path $SrcRoot 'skill')) {
    $SkillDest = Join-Path $Dest 'skill'
    if (Test-Path $SkillDest) { Remove-Item -Recurse -Force $SkillDest }
    Copy-Item -Recurse -Force (Join-Path $SrcRoot 'skill') $SkillDest
}
Write-Ok "installed to $Dest (ur.exe + skill/)"

# 6. User PATH (session + persistent)
if (($env:Path -split ';') -notcontains $Dest) {
    $env:Path = "$env:Path;$Dest"
    $userPath = [Environment]::GetEnvironmentVariable('Path', 'User')
    if (($userPath -split ';') -notcontains $Dest) {
        [Environment]::SetEnvironmentVariable('Path', ($userPath.TrimEnd(';') + ';' + $Dest).TrimStart(';'), 'User')
    }
}

# 7. Done
Write-Ok "ur CLI $Tag installed ($ExePath)"
Write-Host ""
Write-Host "Try it: ur doc parse drawing.dwg --format outline"
Write-Host "        (CAD/DWG parsing via the bundled docling engine)"
