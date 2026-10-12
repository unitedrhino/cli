#!/usr/bin/env python3
"""检查 skill/ 下手写文档体积，防止超出 AI 客户端读取预算被截断。

客户端 AI 工具的 skill_view 读取单个技能文件有约 12000 字节的读取预算，超出部分会被
截断（此前 SKILL.md 内容位于截断线以下导致 AI 读不到，用户投诉）。本脚本遍历 skill/
下所有文件，对手写文档（.md 且不命中生成物/资产豁免规则）检查：

1. >12000 字节：报错，退出码 1，必须拆分到同目录 references/ 子文档；
2. >10000 字节：警告，不失败，留出安全余量；SKILL.md 是客户端首选加载入口，建议
   控制在 10000 字节以内。

生成器产物与资产（references/api/**、*-all-endpoints.md、swagger-index.md、assets/**、
非 .md 文件、_meta.json）豁免于本校验，仅在汇总中提示超限文件数量，拆分改造由生成器
侧单独跟进。

用法:
    cd /path/to/cli
    python3 scripts/check-skill-size.py
"""

import sys
from pathlib import Path

# 客户端 skill_view 单文件读取预算（字节），超出会被截断
ERROR_BYTES = 12000
# 手写文档警告线（字节），为截断线留出安全余量
WARN_BYTES = 10000
# 生成物超限汇总时最多逐条列出的样例数，避免刷屏
SAMPLE_LIMIT = 5


def is_exempt(relative_path):
    """判断相对 skill/ 的路径是否属于生成物/资产豁免范围。

    豁免对象：assets/ 目录、references/api/ 生成目录、*-all-endpoints.md、
    swagger-index.md、非 .md 文件（调用方保证）及 _meta.json（版本元数据）。
    """
    parts = relative_path.parts
    if "assets" in parts:
        return True
    for index in range(len(parts) - 1):
        if parts[index] == "references" and parts[index + 1] == "api":
            return True
    if relative_path.name.endswith("-all-endpoints.md"):
        return True
    if relative_path.name == "swagger-index.md":
        return True
    if relative_path.name == "_meta.json":
        return True
    return False


def scan(skill_dir):
    """遍历 skill/，返回 (手写文档数, 超限文档, 警告文档, 超限生成物)。

    手写文档指 .md 且不命中豁免规则的文档，按字节数降序返回；超限生成物只做
    数量汇总，不参与失败判定，元素为 (字节数, 相对路径)。
    """
    checked = 0
    errors = []
    warnings = []
    exempted = []
    for path in sorted(skill_dir.rglob("*")):
        if not path.is_file():
            continue
        relative = path.relative_to(skill_dir)
        size = path.stat().st_size
        if relative.suffix != ".md" or is_exempt(relative):
            if size > ERROR_BYTES:
                exempted.append((size, relative))
            continue
        checked += 1
        if size > ERROR_BYTES:
            errors.append((size, relative))
        elif size > WARN_BYTES:
            warnings.append((size, relative))
    errors.sort(reverse=True)
    warnings.sort(reverse=True)
    exempted.sort(reverse=True)
    return checked, errors, warnings, exempted


def main():
    """命令行入口：输出体积检查报告，存在超限手写文档时返回非零退出码。"""
    cli_dir = Path(__file__).resolve().parents[1]
    skill_dir = cli_dir / "skill"
    if not skill_dir.is_dir():
        print(f"错误: 未找到 skill 目录: {skill_dir}", file=sys.stderr)
        return 2

    checked, errors, warnings, exempted = scan(skill_dir)

    print("=" * 60)
    print("  Skill 文件体积检查")
    print("=" * 60)
    print(f"  手写文档数: {checked}")
    print(f"  截断线:     {ERROR_BYTES} 字节（超出即失败）")
    print(f"  警告线:     {WARN_BYTES} 字节（超出仅提示）")
    print("=" * 60)

    if errors:
        print("\n❌ 超限手写文档（必须拆分，否则 AI 客户端读取会被截断）:")
        for size, relative in errors:
            print(f"    {size:>7} 字节  skill/{relative}")

    if warnings:
        print("\n⚠️  接近上限的手写文档（建议拆分，留出安全余量）:")
        for size, relative in warnings:
            note = "（SKILL.md 入口，建议 ≤10000 字节）" if relative.name == "SKILL.md" else ""
            print(f"    {size:>7} 字节  skill/{relative}{note}")

    if exempted:
        print(f"\nℹ️  另检出 {len(exempted)} 个生成物文件超限（属生成器已知遗留，豁免），如:")
        for size, relative in exempted[:SAMPLE_LIMIT]:
            print(f"    {size:>7} 字节  skill/{relative}")
        if len(exempted) > SAMPLE_LIMIT:
            print(f"    ... 其余 {len(exempted) - SAMPLE_LIMIT} 个不再逐条列出")

    print()
    if errors:
        print(f"结果: ❌ {len(errors)} 个超限，{len(warnings)} 个警告")
        print("拆分方法: 将超限内容移到同目录 references/ 子文档（文件名小写连字符），")
        print("主文档保留核心结论与导航链接；references/ 已有相关文档时优先合并。")
        return 1
    print(f"结果: ✅ 通过（{len(warnings)} 个警告）")
    return 0


if __name__ == "__main__":
    exit(main())
