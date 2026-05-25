#!/bin/bash
set -e

# 打标签、推送，并同时在 GitHub 和 Gitee 发布 Release
if [ $# -eq 0 ]; then
    echo "用法: ./tag.sh <tag名称> [Release标题]"
    exit 1
fi

tag="$1"
name="${2:-$tag}"

git tag "$tag"
git push origin "$tag"
git push gitee "$tag"
echo "标签 $tag 已推送至 origin 和 gitee"

# 自动生成 Release 说明（取最近 20 条提交）
changelog=$(git log --pretty=format:"- %s" -20)
body="## 变更内容

${changelog}
"

# 发布 GitHub Release
if command -v gh &>/dev/null; then
    gh release create "$tag" --title "$name" --notes "$body" || echo "GitHub Release 创建失败，请检查 gh CLI 登录状态"
    echo "GitHub Release 已发布"
else
    echo "gh CLI 未安装，跳过 GitHub Release"
fi

# 发布 Gitee Release
if [ -n "$GITEE_TOKEN" ]; then
    repo=$(git remote get-url gitee 2>/dev/null | sed 's/.*gitee.com[:/]\([^/]*\)\/\(.*\)\.git/\1\/\2/')
    if [ -n "$repo" ]; then
        curl -s -X POST "https://gitee.com/api/v5/repos/${repo}/releases" \
            -H "Content-Type: application/json" \
            -d "{
                \"access_token\": \"$GITEE_TOKEN\",
                \"tag_name\": \"$tag\",
                \"name\": \"$name\",
                \"body\": \"$body\",
                \"prerelease\": false,
                \"target_commitish\": \"$(git symbolic-ref --short HEAD)\"
            }" | grep -q '"id":' && echo "Gitee Release 已发布" || echo "Gitee Release 创建失败"
    else
        echo "无法解析 Gitee 仓库地址，跳过 Gitee Release"
    fi
else
    echo "环境变量 GITEE_TOKEN 未设置，跳过 Gitee Release"
fi
