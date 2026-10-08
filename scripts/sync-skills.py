#!/usr/bin/env python3
"""同步五组技能树到公开原型仓或整理副本；清除已迁移的旧入口，保留仓库元数据。"""

import argparse
import shutil
import tempfile
from pathlib import Path

# GROUPS 是新结构唯一的业务分组；references 为通用参考目录。
GROUPS = ('ur-iot', 'ur-org-manage', 'ur-ai', 'ur-view', 'ur-doc', 'references')
# LEGACY_DIRECTORIES 只列本次迁移前受管理的目录，禁止据此清理其他任务资源。
LEGACY_DIRECTORIES = (
    'ur-device', 'ur-product', 'ur-project', 'device-firmware', 'ur-ota', 'thing-model',
    'protocol-script', 'scene-linkage', 'ur-device-analytics', 'ur-device-debug',
    'ur-iot-device', 'ur-iot-context', 'ur-iot-client', 'ur-protocol', 'ur-rule',
    'ur-schema', 'ur-iot-user', 'ur-iot-config', 'ur-iot-hook', 'ur-tenant', 'ur-user',
    'ur-system', 'ai-tool',
)
# PRIVATE_TOPICS 是只分发到 CLI 和 SaaS 整理副本的客户端指南。
PRIVATE_TOPICS = ('ur-iot-device', 'ur-iot-context', 'ur-iot-client')


def synchronize(source, destination, public=False, preserve_client_guides=False):
    """同步源目录到目标；public 排除内部专题，preserve_client_guides 保留目标内的内部指南；失败抛出异常。"""
    source, destination = source.resolve(), destination.resolve()
    if source == destination or source in destination.parents or destination in source.parents:
        raise ValueError('源和目标不得相同或相互包含')
    for name in GROUPS[:-1]:
        if not (source / name / 'SKILL.md').is_file():
            raise ValueError(f'源缺少业务组入口: {name}/SKILL.md')
    if not (source / 'SKILL.md').is_file():
        raise ValueError('源缺少总入口 SKILL.md')
    destination.mkdir(parents=True, exist_ok=True)
    temporary_root = destination.parent / '.temp'
    temporary_root.mkdir(exist_ok=True)
    # 完整准备新树后再替换目标；复制失败不会破坏旧树或内部客户端资料。
    with tempfile.TemporaryDirectory(prefix='skills-sync-', dir=temporary_root) as temporary:
        stage = Path(temporary) / 'stage'
        backup = Path(temporary) / 'backup'
        stage.mkdir()
        backup.mkdir()
        shutil.copy2(source / 'SKILL.md', stage / 'SKILL.md')
        for name in GROUPS:
            if (source / name).is_dir():
                shutil.copytree(source / name, stage / name)
        if preserve_client_guides:
            for name in PRIVATE_TOPICS:
                existing = destination / 'ur-iot' / name
                if existing.is_dir():
                    target = stage / 'ur-iot' / name
                    if target.exists():
                        shutil.rmtree(target)
                    shutil.copytree(existing, target)
            entry = stage / 'ur-iot/SKILL.md'
            text = entry.read_text(encoding='utf-8')
            for name in PRIVATE_TOPICS:
                if (stage / 'ur-iot' / name / 'SKILL.md').is_file() and f']({name}/SKILL.md)' not in text:
                    text = text.rstrip() + f'\n- [{name}]({name}/SKILL.md)\n'
            entry.write_text(text, encoding='utf-8')
        if public:
            for name in PRIVATE_TOPICS:
                target = stage / 'ur-iot' / name
                if target.is_dir():
                    shutil.rmtree(target)
            entry = stage / 'ur-iot/SKILL.md'
            lines = entry.read_text(encoding='utf-8').splitlines(keepends=True)
            entry.write_text(''.join(line for line in lines if not any(
                f']({name}/SKILL.md)' in line for name in PRIVATE_TOPICS
            )), encoding='utf-8')
        saved, installed = [], []
        try:
            # 只替换受管理的新旧路径，仓库元数据和其他文件不在操作集合内。
            for name in ('SKILL.md',) + GROUPS + LEGACY_DIRECTORIES:
                target = destination / name
                if target.exists() or target.is_symlink():
                    target.rename(backup / name)
                    saved.append(name)
                if (stage / name).exists():
                    (stage / name).rename(target)
                    installed.append(name)
        except OSError:
            for name in reversed(installed):
                target = destination / name
                if target.is_dir() and not target.is_symlink():
                    shutil.rmtree(target)
                else:
                    target.unlink()
            for name in reversed(saved):
                (backup / name).rename(destination / name)
            raise


def main():
    """解析显式源、目标与公开分发选项；成功返回零，失败保留错误供上层中止。"""
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('source', type=Path)
    parser.add_argument('destination', type=Path)
    modes = parser.add_mutually_exclusive_group()
    modes.add_argument('--public', action='store_true')
    modes.add_argument('--preserve-client-guides', action='store_true')
    args = parser.parse_args()
    synchronize(args.source, args.destination, args.public, args.preserve_client_guides)


if __name__ == '__main__':
    main()
