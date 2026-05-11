# PicoClaw Skill Wrapper — 联犀 SaaS (PowerShell)
param(
    [string]$Query,
    [string]$Section
)

$SkillDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$SkillMd = Join-Path $SkillDir "SKILL.md"

function Get-SectionContent {
    param([string]$FilePath, [string]$StartPattern, [string]$EndPattern)
    $lines = Get-Content $FilePath
    $capture = $false
    $result = @()
    foreach ($line in $lines) {
        if ($line -match $StartPattern) { $capture = $true }
        if ($capture) { $result += $line }
        if ($capture -and $EndPattern -and ($line -match $EndPattern) -and ($line -notmatch $StartPattern)) { break }
    }
    return $result -join "`n"
}

if ($Query) {
    $q = $Query.ToLower()
    if ($q -match "登录|login|认证|auth|setup") {
        Write-Output "⚠️ 约束：本文档是 CLI 使用的唯一事实来源。严格基于文档回答，禁止添加文档外信息。"
        $sec1 = Get-SectionContent $SkillMd "^## 快速开始" "^## CLI 用法"
        $sec2 = Get-SectionContent $SkillMd "^## 认证原理" "^## API 通用约定"
        if ($sec1) { Write-Output $sec1 } else { Write-Output "---"; Get-Content $SkillMd -Raw }
        if ($sec2) { Write-Output "---"; Write-Output $sec2 }
    } elseif ($q -match "api|调用|schema|接口|endpoint") {
        Write-Output "⚠️ 约束：本文档是 CLI 使用的唯一事实来源。严格基于文档回答。"
        $sec = Get-SectionContent $SkillMd "^## CLI 用法" "^## API 通用约定"
        if ($sec) { Write-Output $sec } else { Get-Content $SkillMd -Raw }
    } elseif ($q -match "错误|排查|401|403|404|troubleshoot") {
        Write-Output "⚠️ 约束：本文档是 CLI 使用的唯一事实来源。严格基于文档回答。"
        $sec = Get-SectionContent $SkillMd "^## 常见错误排查" "^## 各域 API 概览"
        if ($sec) { Write-Output $sec } else { Get-Content $SkillMd -Raw }
    } else {
        Get-Content $SkillMd -Raw
    }
} else {
    Get-Content $SkillMd -Raw
}
