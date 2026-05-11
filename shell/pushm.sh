#!/bin/bash
set -e

# 强制推送当前分支到两个远程的指定分支（默认 master）
target="${1:-master}"
branch=$(git symbolic-ref --short HEAD)
echo "强制推送: $branch → $target (origin + gitee)"
git push -f origin "$branch:$target"
git push -f gitee "$branch:$target"
echo "完成"
