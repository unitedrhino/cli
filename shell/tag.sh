#!/bin/bash
set -e

# 打标签并推送到 origin 和 gitee
if [ $# -eq 0 ]; then
    echo "用法: ./tag.sh <tag名称>"
    exit 1
fi

tag="$1"
git tag "$tag"
git push origin "$tag"
git push gitee "$tag"
echo "标签 $tag 已推送至 origin 和 gitee"
