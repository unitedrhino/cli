#!/usr/bin/env python3
"""
验证 skills 准确率：
1. 检查 skill references 中的路径是否都在 swagger 中
2. 检查 swagger 中的路径是否都被 skill 覆盖
3. 检查手写命令中的路径是否都在 swagger 中

用法:
    cd /path/to/cli
    python3 scripts/verify-skills.py
"""

import json
import re
from pathlib import Path


def load_swagger_paths(swagger_dir):
    """加载所有 swagger 中的 API 路径"""
    all_paths = set()
    files = ['core-api.json', 'things-api.json', 'core-ai.json', 'things-ai.json']
    for filename in files:
        p = Path(swagger_dir) / filename
        if not p.exists():
            continue
        with open(p, 'r', encoding='utf-8') as f:
            data = json.load(f)
        for path in data.get('paths', {}).keys():
            all_paths.add(path)
    return all_paths


def load_skill_paths(skill_dir):
    """加载所有 skill 文件中引用的 API 路径"""
    skill_paths = set()
    skill_dir = Path(skill_dir)
    for md_file in skill_dir.rglob('*.md'):
        with open(md_file, 'r', encoding='utf-8') as f:
            content = f.read()
        # 匹配 `POST /api/v1/xxx/xxx` 或 `/api/v1/xxx/xxx`
        for match in re.finditer(r'`(?:[A-Z]+\s+)?(/api/v1/[^`\s]+)`', content):
            path = match.group(1)
            if '...' not in path and path != '/api/v1/':
                skill_paths.add(path)
    return skill_paths


def load_handwritten_paths(shared_dir):
    """加载手写命令中硬编码的 API 路径"""
    handwritten_paths = set()
    shared_dir = Path(shared_dir)
    for go_file in shared_dir.glob('*.go'):
        with open(go_file, 'r', encoding='utf-8') as f:
            content = f.read()
        for m in re.finditer(r'Path:\s*"(/api/v1/[^"]+)"', content):
            handwritten_paths.add(m.group(1))
    return handwritten_paths


def main():
    cli_dir = Path(__file__).parent.parent.resolve()
    swagger_dir = cli_dir / 'backend' / '.swagger'

    # 如果 cli 内部没有 swagger，尝试从上级目录找
    if not swagger_dir.exists():
        swagger_dir = cli_dir.parent.parent / 'backend' / '.swagger'

    # 加载路径
    swagger_paths = load_swagger_paths(swagger_dir)
    skill_paths = load_skill_paths(cli_dir / 'skill')
    handwritten_paths = load_handwritten_paths(cli_dir / 'cmd' / 'shared')

    # 计算指标
    wrong = sorted([p for p in skill_paths if p not in swagger_paths])
    missing = sorted([p for p in swagger_paths if p not in skill_paths])
    handwritten_wrong = sorted([p for p in handwritten_paths if p not in swagger_paths])

    total_skill = len(skill_paths)
    accuracy = (total_skill - len(wrong)) / total_skill * 100 if total_skill else 0
    coverage = len([p for p in swagger_paths if p in skill_paths]) / len(swagger_paths) * 100 if swagger_paths else 0

    print("=" * 60)
    print("  Skills 验证报告")
    print("=" * 60)
    print(f"  Swagger 总路径数:     {len(swagger_paths)}")
    print(f"  Skill 覆盖路径数:     {total_skill}")
    print(f"  Skill 准确率:         {accuracy:.1f}%")
    print(f"  Skill 覆盖率:         {coverage:.1f}%")
    print(f"  错误路径数:           {len(wrong)}")
    print(f"  缺失路径数:           {len(missing)}")
    print(f"  手写命令错误路径:     {len(handwritten_wrong)}")
    print("=" * 60)

    if wrong:
        print(f"\n❌ Skill 中的错误路径（不在 swagger 中）:")
        for p in wrong[:20]:
            print(f"    {p}")
        if len(wrong) > 20:
            print(f"    ... 还有 {len(wrong) - 20} 个")

    if missing:
        print(f"\n⚠️  Swagger 中缺失的路径（未被 skill 覆盖）:")
        for p in missing[:20]:
            print(f"    {p}")
        if len(missing) > 20:
            print(f"    ... 还有 {len(missing) - 20} 个")

    if handwritten_wrong:
        print(f"\n❌ 手写命令中的错误路径:")
        for p in handwritten_wrong:
            print(f"    {p}")

    if not wrong and not missing and not handwritten_wrong:
        print("\n✅ 全部通过！")

    # 返回退出码
    return 0 if len(wrong) == 0 and len(missing) == 0 and len(handwritten_wrong) == 0 else 1


if __name__ == '__main__':
    exit(main())
