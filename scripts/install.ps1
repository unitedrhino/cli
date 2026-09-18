# ur CLI 一键安装脚本(Windows PowerShell 5.1+)。对齐 urops install.ps1 体验:
# 自动查最新版本、SHA256 校验、国内默认 Gitee 源。
#
# 用法(PowerShell):
#   $script = iwr -UseBasicParsing <脚本地址>.ps1 | Select-Object -ExpandProperty Content
#   Invoke-Expression $script
#   指定版本: 环境变量 UR_VERSION=v0.7.0 UR_SOURCE=gitee|github
#
# 卸载: Remove-Item -Recurse -Force "$env:USERPROFILE\.local\lib\ur"

$ErrorActionPreference = "Stop"

$Version = if ($env:UR_VERSION) { $env:UR_VERSION } else { "latest" }
$Source  = if ($env:UR_SOURCE)  { $env:UR_SOURCE }  else { "gitee" }
$InstallDir = Join-Path $env:USERPROFILE ".local\lib\ur"
$BinDir     = Join-Path $env:USERPROFILE ".local\bin"
$RepoGithub = "unitedrhino/cli"
$RepoGitee  = "unitedrhino/cli"

# ── 平台检测 ─────────────────────────────────────────────────────────────
$Arch = switch ($env:PROCESSOR_ARCHITECTURE) {
    "AMD64" { "x86_64" }
    "ARM64" { "arm64" }
    default { throw "不支持的架构: $env:PROCESSOR_ARCHITECTURE" }
}
$Platform = "Windows-$Arch"

# ── 版本解析 ─────────────────────────────────────────────────────────────
if ($Version -eq "latest") {
    if ($Source -eq "gitee") {
        $rel = Invoke-RestMethod -Uri "https://gitee.com/api/v5/repos/$RepoGitee/releases?per_page=1"
    } else {
        $rel = Invoke-RestMethod -Uri "https://api.github.com/repos/$RepoGithub/releases/latest"
    }
    $Version = $rel[0].tag_name
}
$Asset = "ur-cli-$Version-$Platform.zip"
if ($Source -eq "gitee") {
    $Url = "https://gitee.com/$RepoGitee/releases/download/$Version/$Asset"
    $SumsUrl = "https://gitee.com/$RepoGitee/releases/download/$Version/sha256sums.txt"
} else {
    $Url = "https://github.com/$RepoGithub/releases/download/$Version/$Asset"
    $SumsUrl = "https://github.com/$RepoGithub/releases/download/$Version/sha256sums.txt"
}

Write-Host "[ur-install] 平台: $Platform  版本: $Version  来源: $Source"

# ── 下载 + SHA256 校验 ───────────────────────────────────────────────────
$Tmp = Join-Path $env:TEMP ("ur-install-" + [guid]::NewGuid().ToString("N"))
New-Item -ItemType Directory -Force -Path $Tmp | Out-Null
$PkgPath = Join-Path $Tmp $Asset
Write-Host "[ur-install] 下载 $Asset ..."
Invoke-WebRequest -UseBasicParsing -Uri $Url -OutFile $PkgPath

try {
    $SumsPath = Join-Path $Tmp "sha256sums.txt"
    Invoke-WebRequest -UseBasicParsing -Uri $SumsUrl -OutFile $SumsPath
    $Expected = (Select-String -Path $SumsPath -Pattern [regex]::Escape($Asset) |
        Select-Object -First 1).Line -split "\s+" | Select-Object -First 1
    if ($Expected) {
        $Actual = (Get-FileHash -Algorithm SHA256 $PkgPath).Hash.ToLower()
        if ($Actual -ne $Expected.ToLower()) { throw "SHA256 校验失败" }
        Write-Host "[ur-install] SHA256 校验通过"
    }
} catch { Write-Host "[ur-install] 跳过校验: $($_.Exception.Message)" }

# ── 安装:ur.exe 与 skill/ 同级,所在目录加入用户 PATH ────────────────────
New-Item -ItemType Directory -Force -Path $InstallDir, $BinDir | Out-Null
Expand-Archive -Force -Path $PkgPath -DestinationPath $InstallDir

$UserPath = [Environment]::GetEnvironmentVariable("Path", "User")
if (($UserPath -split ";") -notcontains $BinDir) {
    [Environment]::SetEnvironmentVariable("Path", "$UserPath;$BinDir", "User")
    Write-Host "[ur-install] 已把 $BinDir 加入用户 PATH(新开终端生效)"
}

& (Join-Path $InstallDir "ur.exe") --version
Write-Host "[ur-install] 完成。下一步: ur check --json(检查认证)"
Write-Host "[ur-install] 可选: ur skills install 安装 ur-api skills 到本地 AI 工具目录"

Remove-Item -Recurse -Force $Tmp -ErrorAction SilentlyContinue
