#!/usr/bin/env python3
"""
从 backend/.swagger/*.json 读取 swagger，按 domain 分组生成 API 端点列表，
插入到 cli/skill/*/SKILL.md 的 AUTO-GENERATED 标记之间。

用法:
    cd /path/to/cli
    python3 scripts/generate-api-lists.py

需要环境变量或相对路径找到 backend/.swagger/:
    - $UR_SWAGGER_DIR
    - ../backend/.swagger/
    - ../../backend/.swagger/
"""

import json
import os
import re
import sys
from pathlib import Path

# domain -> 路径前缀列表（按优先级排序，长的在前）
DOMAIN_PREFIXES = {
    "ur-device-analytics": ["/api/v1/things/device/msg/"],
    "ur-device-debug": ["/api/v1/things/device/interact/"],
    "ur-device": ["/api/v1/things/device/"],
    "ur-product": ["/api/v1/things/product/", "/api/v1/things/device/ota/"],
    "ur-project": ["/api/v1/things/project/", "/api/v1/things/area/", "/api/v1/things/group/", "/api/v1/things/data/"],
    "ur-user": ["/api/v1/system/user/", "/api/v1/system/role/", "/api/v1/system/dept/", "/api/v1/system/dict/", "/api/v1/system/notify/", "/api/v1/system/log/"],
    "ur-tenant": ["/api/v1/system/tenant/"],
    "ur-system": ["/api/v1/system/"],
    "ur-ai": ["/api/v1/ai/", "/api/v1/things/alarm/", "/api/v1/things/scene/"],
}

# 反向映射：路径前缀 -> domain
PREFIX_TO_DOMAIN = []
for domain, prefixes in DOMAIN_PREFIXES.items():
    for prefix in prefixes:
        PREFIX_TO_DOMAIN.append((prefix, domain))
# 按前缀长度降序，保证最长前缀优先匹配
PREFIX_TO_DOMAIN.sort(key=lambda x: len(x[0]), reverse=True)

# 特殊处理：scene-linkage 只取 scene 相关
SCENE_LINKAGE_PREFIX = "/api/v1/things/scene/"


def find_swagger_dir():
    """查找 swagger 文件目录"""
    candidates = []
    if os.environ.get("UR_SWAGGER_DIR"):
        candidates.append(os.environ["UR_SWAGGER_DIR"])

    # 从当前目录向上查找 backend/.swagger
    cwd = Path.cwd()
    for _ in range(5):
        candidates.append(str(cwd / "backend" / ".swagger"))
        candidates.append(str(cwd.parent / "backend" / ".swagger"))
        candidates.append(str(cwd.parent.parent / "backend" / ".swagger"))
        cwd = cwd.parent

    for candidate in candidates:
        d = Path(candidate)
        core = d / "core-api.json"
        things = d / "things-api.json"
        if core.exists() and things.exists():
            return str(d)

    return None


def load_endpoints(swagger_dir):
    """加载并合并所有 swagger 端点"""
    endpoints = []
    for filename in ["core-api.json", "things-api.json"]:
        path = Path(swagger_dir) / filename
        if not path.exists():
            print(f"警告: 找不到 {path}", file=sys.stderr)
            continue
        with open(path, "r", encoding="utf-8") as f:
            data = json.load(f)
        for api_path, methods in data.get("paths", {}).items():
            for method, op in methods.items():
                endpoints.append({
                    "path": api_path,
                    "method": method.upper(),
                    "summary": op.get("summary", ""),
                    "description": op.get("description", ""),
                    "auth_type": op.get("x-auth-type", ""),
                    "group": op.get("x-group", ""),
                })
    # 排序
    endpoints.sort(key=lambda e: (e["path"], e["method"]))
    return endpoints


def classify_endpoint(endpoint):
    """将端点分类到 domain"""
    path = endpoint["path"]

    # 特殊：scene-linkage 只取 scene
    if path.startswith(SCENE_LINKAGE_PREFIX):
        return "scene-linkage"

    # 其他按前缀匹配
    for prefix, domain in PREFIX_TO_DOMAIN:
        if path.startswith(prefix):
            return domain

    return None


def group_by_domain(endpoints):
    """按 domain 分组端点"""
    groups = {}
    for ep in endpoints:
        domain = classify_endpoint(ep)
        if domain is None:
            continue
        if domain not in groups:
            groups[domain] = []
        groups[domain].append(ep)
    return groups


def generate_markdown_table(endpoints):
    """生成 Markdown 表格"""
    lines = ["| 方法 | 端点 | 说明 | 权限 |", "|------|------|------|------|"]
    for ep in endpoints:
        auth = ep["auth_type"] or "-"
        summary = ep["summary"] or ep["description"] or "-"
        # 清理 markdown 特殊字符
        summary = summary.replace("|", "\\|").replace("\n", " ")
        lines.append(f"| {ep['method']} | `{ep['path']}` | {summary} | {auth} |")
    return "\n".join(lines)


def update_skill_file(skill_dir, domain, new_content):
    """更新 skill 文件中的 AUTO-GENERATED 区域"""
    skill_path = Path(skill_dir) / domain / "SKILL.md"
    if not skill_path.exists():
        print(f"警告: 找不到 {skill_path}", file=sys.stderr)
        return False

    with open(skill_path, "r", encoding="utf-8") as f:
        content = f.read()

    marker_start = f"<!-- API_LIST:{domain} -->"
    marker_end = "<!-- END_API_LIST -->"

    if marker_start not in content:
        print(f"警告: {skill_path} 中找不到标记 {marker_start}，跳过", file=sys.stderr)
        return False

    # 替换标记之间的内容
    pattern = re.compile(
        re.escape(marker_start) + ".*?" + re.escape(marker_end),
        re.DOTALL,
    )
    replacement = f"{marker_start}\n\n{new_content}\n\n{marker_end}"
    new_content_full = pattern.sub(replacement, content)

    with open(skill_path, "w", encoding="utf-8") as f:
        f.write(new_content_full)

    print(f"已更新 {skill_path} ({len(new_content.split(chr(10)))} 行)")
    return True


def main():
    swagger_dir = find_swagger_dir()
    if not swagger_dir:
        print("错误: 找不到 swagger 文件目录。请设置 UR_SWAGGER_DIR 或在 backend/.swagger 附近运行。", file=sys.stderr)
        sys.exit(1)

    print(f"使用 swagger 目录: {swagger_dir}")

    cli_dir = Path(__file__).parent.parent
    skill_dir = cli_dir / "skill"

    endpoints = load_endpoints(swagger_dir)
    print(f"加载了 {len(endpoints)} 个端点")

    groups = group_by_domain(endpoints)

    updated = 0
    for domain, eps in sorted(groups.items()):
        if not eps:
            continue
        md = generate_markdown_table(eps)
        if update_skill_file(skill_dir, domain, md):
            updated += 1
        print(f"  {domain}: {len(eps)} 个端点")

    print(f"\n共更新 {updated} 个 skill 文件")


if __name__ == "__main__":
    main()
