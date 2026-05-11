#!/bin/bash
set -e

# 同时推送到 origin(GitHub) 和 gitee
branch=$(git symbolic-ref --short HEAD)
echo "推送分支: $branch"
git push origin "$branch"
git push gitee "$branch"
echo "完成: $branch 已推送至 origin 和 gitee"
