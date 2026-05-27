#!/usr/bin/env python3
"""
从 backend/.swagger/*.json 读取 swagger，按 domain 分组生成 API 端点列表。

生成内容：
1. 为每个子域生成 references/api/*.md（按 x-group 分组，含详细参数说明、请求/响应示例）
2. 更新 SKILL.md 的 API_LIST 区域为功能索引表（链接到 references/api/）

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
    "ur-product": ["/api/v1/things/product/"],
    "ur-ota": ["/api/v1/things/ota/"],
    "ur-protocol": ["/api/v1/things/protocol/"],
    "ur-rule": ["/api/v1/things/rule/"],
    "ur-schema": ["/api/v1/things/schema/"],
    "ur-project": ["/api/v1/things/project/", "/api/v1/things/area/", "/api/v1/things/group/", "/api/v1/things/data/"],
    "ur-iot-user": ["/api/v1/things/user/"],
    "ur-iot-config": ["/api/v1/things/config/"],
    "ur-iot-hook": ["/api/v1/things/hook/"],
    "ur-user": ["/api/v1/system/user/", "/api/v1/system/role/", "/api/v1/system/dept/", "/api/v1/system/dict/", "/api/v1/system/notify/", "/api/v1/system/log/"],
    "ur-tenant": ["/api/v1/system/tenant/"],
    "ur-system": ["/api/v1/system/"],
    "ur-ai": ["/api/v1/ai/", "/api/v1/things/ai/", "/api/v1/things/alarm/", "/api/v1/things/scene/"],
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

# 生成示例 JSON 时的深度限制，防止循环引用
MAX_EXAMPLE_DEPTH = 4


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


def load_swagger(swagger_dir):
    """加载并合并所有 swagger 数据"""
    all_data = {}
    for filename in ["core-api.json", "things-api.json", "core-ai.json", "things-ai.json"]:
        path = Path(swagger_dir) / filename
        if not path.exists():
            print(f"警告: 找不到 {path}", file=sys.stderr)
            continue
        with open(path, "r", encoding="utf-8") as f:
            data = json.load(f)
        all_data[filename] = data
    return all_data


def load_endpoints(swagger_dir):
    """加载并合并所有 swagger 端点"""
    endpoints = []
    for filename in ["core-api.json", "things-api.json", "core-ai.json", "things-ai.json"]:
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
                    "operationId": op.get("operationId", ""),
                    "requestBody": op.get("requestBody", {}),
                    "responses": op.get("responses", {}),
                    "parameters": op.get("parameters", []),
                    "_swagger_file": filename,
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


def get_schemas(swagger_data):
    """从所有 swagger 文件中提取 schemas"""
    schemas = {}
    for data in swagger_data.values():
        comps = data.get("components", {})
        sc = comps.get("schemas", {})
        schemas.update(sc)
    return schemas


def resolve_schema(schemas, schema, depth=0):
    """解析 schema 引用，返回展开后的 schema 字典"""
    if depth > 10:
        return {"type": "object", "description": "(嵌套过深)"}
    if not schema:
        return {}
    if "$ref" in schema:
        ref = schema["$ref"]
        name = ref.split("/")[-1]
        if name in schemas:
            resolved = dict(schemas[name])
            resolved["_resolved_name"] = name
            return resolve_schema(schemas, resolved, depth + 1)
        return {"type": "object", "description": f"(未解析: {name})"}
    return schema


def extract_properties(schema, schemas, prefix="", depth=0):
    """提取 schema 的所有属性为扁平列表，用于表格展示"""
    if depth > MAX_EXAMPLE_DEPTH:
        return []

    resolved = resolve_schema(schemas, schema, depth)
    props = resolved.get("properties", {})
    required = set(resolved.get("required", []))

    rows = []
    for name, prop in sorted(props.items()):
        prop = resolve_schema(schemas, prop, depth + 1)
        ptype = prop.get("type", "object")
        if ptype == "array":
            items = resolve_schema(schemas, prop.get("items", {}), depth + 1)
            item_type = items.get("type", "object")
            if "$ref" in prop.get("items", {}):
                item_type = prop["items"]["$ref"].split("/")[-1]
            ptype = f"array[{item_type}]"
        desc = prop.get("description", "")
        fmt = prop.get("format", "")
        if fmt:
            desc = f"{desc} (格式: {fmt})" if desc else f"格式: {fmt}"
        enum = prop.get("enum", [])
        if enum:
            desc = f"{desc} 可选: {enum}" if desc else f"可选: {enum}"

        is_required = "是" if name in required else "否"
        rows.append({
            "name": f"{prefix}{name}",
            "type": ptype,
            "required": is_required,
            "description": desc,
        })

        # 嵌套对象展开一层
        if ptype == "object" and depth < 2 and "properties" in prop:
            nested = extract_properties(prop, schemas, prefix=f"{prefix}{name}.", depth=depth + 1)
            rows.extend(nested)

    return rows


def generate_example_value(prop, schemas, depth=0):
    """根据 schema 属性生成示例值"""
    if depth > MAX_EXAMPLE_DEPTH:
        return "..."

    prop = resolve_schema(schemas, prop, depth)
    ptype = prop.get("type", "")

    if ptype == "string":
        fmt = prop.get("format", "")
        enum = prop.get("enum", [])
        if enum:
            return enum[0]
        if fmt == "int64":
            return "1"
        if "时间" in prop.get("description", "") or fmt == "date-time":
            return "2026-01-01T00:00:00Z"
        if "ID" in prop.get("description", "") or name_ends_with_id(prop):
            return "string"
        if "名称" in prop.get("description", "") or "别名" in prop.get("description", ""):
            return "示例名称"
        if "密码" in prop.get("description", ""):
            return "******"
        return "string"

    if ptype == "integer" or ptype == "number":
        enum = prop.get("enum", [])
        if enum:
            return enum[0]
        minimum = prop.get("minimum", 0)
        return minimum if minimum != 0 else 1

    if ptype == "boolean":
        return True

    if ptype == "array":
        items = prop.get("items", {})
        item_example = generate_example_value(items, schemas, depth + 1)
        if item_example == "...":
            return []
        return [item_example]

    if ptype == "object":
        props = prop.get("properties", {})
        if not props:
            return {}
        result = {}
        for k, v in props.items():
            result[k] = generate_example_value(v, schemas, depth + 1)
        return result

    return None


def name_ends_with_id(prop):
    """判断属性名是否以 ID/id/Id 结尾"""
    # 通过 _resolved_name 判断比较麻烦，简单处理
    return False


def generate_request_body_example(schema, schemas):
    """生成请求体示例 JSON"""
    resolved = resolve_schema(schemas, schema)
    props = resolved.get("properties", {})
    if not props:
        return {}
    result = {}
    for k, v in props.items():
        val = generate_example_value(v, schemas)
        if val is not None:
            result[k] = val
    return result


def generate_response_example(responses, schemas):
    """生成响应示例 JSON"""
    resp_200 = responses.get("200", {})
    content = resp_200.get("content", {})
    app_json = content.get("application/json", {})
    schema = app_json.get("schema", {})
    if not schema:
        return {"code": 200, "msg": "success", "data": {}}

    resolved = resolve_schema(schemas, schema)
    props = resolved.get("properties", {})
    if not props:
        return {"code": 200, "msg": "success", "data": {}}

    result = {}
    for k, v in props.items():
        if k == "code":
            result[k] = 200
        elif k == "msg":
            result[k] = "success"
        elif k == "data":
            # 尝试展开 data 结构
            data_schema = resolve_schema(schemas, v)
            if data_schema.get("type") == "array":
                items = data_schema.get("items", {})
                item_ex = generate_request_body_example(items, schemas)
                result[k] = [item_ex] if item_ex else []
            elif "properties" in data_schema:
                result[k] = generate_request_body_example(v, schemas)
            else:
                result[k] = generate_example_value(v, schemas)
        else:
            result[k] = generate_example_value(v, schemas)
    return result


def sanitize_filename(name):
    """将 group 名转换为合法的文件名"""
    return re.sub(r'[^a-zA-Z0-9_-]', '-', name).strip('-').lower()


def generate_endpoint_detail(ep, schemas):
    """生成单个端点的详细说明块"""
    lines = []

    # 端点标题
    lines.append(f"### {ep['method']} `{ep['path']}`")
    lines.append("")

    # 说明
    summary = ep["summary"] or ep["description"] or "无说明"
    lines.append(f"**说明**: {summary}")
    lines.append("")

    # 权限
    auth = ep["auth_type"] or "-"
    lines.append(f"**权限**: {auth}")
    lines.append("")

    # 请求参数
    params = ep.get("parameters", [])
    if params:
        lines.append("**路径参数**:")
        lines.append("")
        lines.append("| 参数 | 位置 | 类型 | 必填 | 说明 |")
        lines.append("|------|------|------|------|------|")
        for p in params:
            pname = p.get("name", "")
            p_in = p.get("in", "")
            ptype = p.get("schema", {}).get("type", "string")
            required = "是" if p.get("required") else "否"
            desc = p.get("description", "")
            lines.append(f"| `{pname}` | {p_in} | {ptype} | {required} | {desc} |")
        lines.append("")

    # 请求体
    rb = ep.get("requestBody", {})
    if rb:
        content = rb.get("content", {})
        app_json = content.get("application/json", {})
        schema = app_json.get("schema", {})
        if schema:
            lines.append("**请求体字段**:")
            lines.append("")
            lines.append("| 字段 | 类型 | 必填 | 说明 |")
            lines.append("|------|------|------|------|")
            props = extract_properties(schema, schemas)
            for p in props:
                desc = p["description"].replace("|", "\\|")
                lines.append(f"| `{p['name']}` | {p['type']} | {p['required']} | {desc} |")
            lines.append("")

            # 请求体示例
            example = generate_request_body_example(schema, schemas)
            if example:
                lines.append("**请求示例**:")
                lines.append("```json")
                lines.append(json.dumps(example, indent=2, ensure_ascii=False))
                lines.append("```")
                lines.append("")

    # 响应示例
    responses = ep.get("responses", {})
    if responses:
        example = generate_response_example(responses, schemas)
        lines.append("**响应示例**:")
        lines.append("```json")
        lines.append(json.dumps(example, indent=2, ensure_ascii=False))
        lines.append("```")
        lines.append("")

    # 使用示例
    lines.append("**调用示例**:")
    lines.append("```bash")
    # 生成 ur api 命令
    path = ep["path"]
    # 如果有请求体示例，包含 --body
    rb = ep.get("requestBody", {})
    if rb:
        content = rb.get("content", {})
        app_json = content.get("application/json", {})
        schema = app_json.get("schema", {})
        if schema:
            example = generate_request_body_example(schema, schemas)
            body_json = json.dumps(example, indent=2, ensure_ascii=False)
            # 压缩为一行用于命令行
            body_oneline = json.dumps(example, ensure_ascii=False)
            lines.append(f"ur api {path} \\")
            lines.append(f"  --body '{body_oneline}'")
        else:
            lines.append(f"ur api {path} \\")
            lines.append(f"  --body '{{}}'")
    else:
        lines.append(f"ur api {path} \\")
        lines.append(f"  --body '{{}}'")
    lines.append("```")
    lines.append("")

    return "\n".join(lines)


def generate_reference_file(domain, group, endpoints, schemas):
    """生成单个 reference 文件的内容"""
    lines = [f"# {domain} {group}", ""]

    # 说明
    summaries = [ep["summary"] for ep in endpoints if ep["summary"]]
    if summaries:
        desc = summaries[0] if len(set(summaries)) == 1 else f"{summaries[0]} 等"
        lines.append(f"{desc}")
        lines.append("")

    # 端点概览表格
    lines.append("## 端点概览")
    lines.append("")
    lines.append("| 方法 | 端点 | 说明 | 权限 |")
    lines.append("|------|------|------|------|")
    for ep in endpoints:
        auth = ep["auth_type"] or "-"
        summary = ep["summary"] or ep["description"] or "-"
        summary = summary.replace("|", "\\|").replace("\n", " ")
        lines.append(f"| {ep['method']} | `{ep['path']}` | {summary} | {auth} |")
    lines.append("")

    # 每个端点的详细说明
    lines.append("## 详细说明")
    lines.append("")
    for ep in endpoints:
        detail = generate_endpoint_detail(ep, schemas)
        lines.append(detail)

    return "\n".join(lines)


def generate_index_table(domain, group_endpoints):
    """生成功能索引表（用于 SKILL.md 的 API_LIST 区域）"""
    lines = ["| 功能组 | 说明 | 参考文档 |", "|--------|------|---------|"]
    for group, endpoints in sorted(group_endpoints.items()):
        # 获取该组的主要说明
        summaries = [ep["summary"] for ep in endpoints if ep["summary"]]
        desc = summaries[0] if summaries else "-"
        if len(desc) > 30:
            desc = desc[:27] + "..."
        desc = desc.replace("|", "\\|").replace("\n", " ")

        # 统计端点数
        count = len(endpoints)
        count_label = f"{count} 个端点" if count > 1 else "1 个端点"

        ref_name = sanitize_filename(group)
        ref_file = f"{domain}-{ref_name}.md"
        lines.append(f"| `{group}` | {desc} ({count_label}) | [{ref_file}](references/api/{ref_file}) |")

    lines.append("")
    # 添加所有端点汇总链接
    lines.append(f"[查看全部端点](references/api/{domain}-all-endpoints.md)")
    lines.append("")
    return "\n".join(lines)


def generate_all_endpoints_table(endpoints, schemas):
    """生成所有端点的完整表格（速查用）"""
    lines = ["# 全部端点速查", ""]
    lines.append("| 方法 | 端点 | 说明 | 权限 |")
    lines.append("|------|------|------|------|")
    for ep in endpoints:
        auth = ep["auth_type"] or "-"
        summary = ep["summary"] or ep["description"] or "-"
        summary = summary.replace("|", "\\|").replace("\n", " ")
        lines.append(f"| {ep['method']} | `{ep['path']}` | {summary} | {auth} |")
    lines.append("")

    # 添加每个端点的简要说明
    lines.append("## 端点详情")
    lines.append("")
    for ep in endpoints:
        detail = generate_endpoint_detail(ep, schemas)
        lines.append(detail)

    return "\n".join(lines)


def ensure_references_dir(skill_dir, domain):
    """确保 references/api/ 目录存在，并清理旧文件"""
    ref_dir = Path(skill_dir) / domain / "references" / "api"
    if ref_dir.exists():
        # 清理旧的 .md 文件
        for f in ref_dir.glob("*.md"):
            f.unlink()
    else:
        ref_dir.mkdir(parents=True, exist_ok=True)
    return ref_dir


def write_reference_files(ref_dir, domain, group_endpoints, schemas):
    """写入所有 reference 文件到 references/api/"""
    all_endpoints = []
    for group, endpoints in group_endpoints.items():
        all_endpoints.extend(endpoints)
        ref_name = sanitize_filename(group)
        ref_file = ref_dir / f"{domain}-{ref_name}.md"
        content = generate_reference_file(domain, group, endpoints, schemas)
        with open(ref_file, "w", encoding="utf-8") as f:
            f.write(content)
        print(f"  生成 {ref_file.name} ({len(endpoints)} 个端点)")

    # 生成全部端点速查
    all_ref_file = ref_dir / f"{domain}-all-endpoints.md"
    content = generate_all_endpoints_table(all_endpoints, schemas)
    with open(all_ref_file, "w", encoding="utf-8") as f:
        f.write(content)
    print(f"  生成 {all_ref_file.name} ({len(all_endpoints)} 个端点)")


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

    print(f"已更新 {skill_path}")
    return True


def main():
    swagger_dir = find_swagger_dir()
    if not swagger_dir:
        print("错误: 找不到 swagger 文件目录。请设置 UR_SWAGGER_DIR 或在 backend/.swagger 附近运行。", file=sys.stderr)
        sys.exit(1)

    print(f"使用 swagger 目录: {swagger_dir}")

    cli_dir = Path(__file__).parent.parent
    skill_dir = cli_dir / "skill"

    # 加载 swagger 数据和 schemas
    swagger_data = load_swagger(swagger_dir)
    schemas = get_schemas(swagger_data)
    print(f"加载了 {len(schemas)} 个 schema 定义")

    endpoints = load_endpoints(swagger_dir)
    print(f"加载了 {len(endpoints)} 个端点")

    groups = group_by_domain(endpoints)

    updated = 0
    for domain, eps in sorted(groups.items()):
        if not eps:
            continue

        # 按 x-group 分组
        group_endpoints = {}
        for ep in eps:
            g = ep["group"] or "未分组"
            if g not in group_endpoints:
                group_endpoints[g] = []
            group_endpoints[g].append(ep)

        print(f"\n{domain}: {len(eps)} 个端点, {len(group_endpoints)} 个分组")

        # 生成 references/ 目录和文件
        ref_dir = ensure_references_dir(skill_dir, domain)
        write_reference_files(ref_dir, domain, group_endpoints, schemas)

        # 生成索引表并更新 SKILL.md
        index_table = generate_index_table(domain, group_endpoints)
        if update_skill_file(skill_dir, domain, index_table):
            updated += 1

    print(f"\n共更新 {updated} 个 skill 文件，生成了 references/api/ 目录")


if __name__ == "__main__":
    main()
