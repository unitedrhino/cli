#!/usr/bin/env python3
"""验证五组技能元信息、链接、生成路径及公开同步边界；全部写操作使用隔离夹具。"""

import importlib.util
import contextlib
import io
import os
import subprocess
import sys
import shutil
from pathlib import Path
import re
import tempfile
import unittest
from unittest.mock import patch
from urllib.parse import unquote, urlsplit

import yaml

# 测试导入正式脚本时禁止在源码目录生成字节码缓存。
sys.dont_write_bytecode = True

# ROOT 是 CLI 仓库位置，SKILLS 是本次需要校验的分发源。
ROOT = Path(__file__).resolve().parents[1]
SKILLS = ROOT / 'skill'


def load_script(name):
    """按脚本名加载正式实现供夹具调用，返回模块，不执行命令行入口。"""
    spec = importlib.util.spec_from_file_location(name, ROOT / 'scripts' / f'{name}.py')
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)
    return module


class SkillsLayoutTest(unittest.TestCase):
    """覆盖结构、公开边界与可重入生成，防止工具链重新生成平铺目录。"""

    def setUp(self):
        """为生成和同步测试准备任务临时目录。"""
        (ROOT / '.temp').mkdir(exist_ok=True)
        self.temporary = tempfile.TemporaryDirectory(dir=ROOT / '.temp')
        self.addCleanup(self.temporary.cleanup)
        self.fixture = Path(self.temporary.name)

    def test_entries_and_links(self):
        """五组是唯一业务目录，各层声明名称唯一，链接存在且没有 GUIDE 入口。"""
        self.assertEqual({p.name for p in SKILLS.iterdir() if p.is_dir()},
                         {'references', 'ur-iot', 'ur-org-manage', 'ur-ai', 'ur-view', 'ur-doc'})
        self.assertFalse(list(SKILLS.rglob('GUIDE.md')))
        names = set()
        for entry in SKILLS.rglob('SKILL.md'):
            text = entry.read_text()
            match = re.match(r'^---\n(.*?)\n---', text, re.S)
            self.assertIsNotNone(match, str(entry))
            metadata = yaml.safe_load(match[1])
            name = metadata['name']
            self.assertNotIn(name, names, str(entry))
            self.assertTrue(metadata['description'])
            names.add(name)
        for entry in SKILLS.rglob('*.md'):
            for raw in re.findall(r'!?\[[^\]\n]*\]\(\s*(<[^>]+>|[^\s)]+)', entry.read_text()):
                parsed = urlsplit(raw.strip('<>'))
                if not parsed.scheme and not parsed.netloc and parsed.path:
                    self.assertTrue((entry.parent / unquote(parsed.path)).exists(), f'{entry}: {raw}')

    def test_generated_domains_stay_grouped(self):
        """所有 API 域输出归入五组；重复生成保留手写资料并且不生成顶层子模块。"""
        generator = load_script('generate-api-lists')
        for _ in range(2):
            for domain in generator.DOMAIN_PREFIXES:
                target = generator.ensure_references_dir(self.fixture, domain)
                self.assertTrue(target.is_relative_to(self.fixture / 'ur-iot') or
                                target.is_relative_to(self.fixture / 'ur-org-manage') or
                                target.is_relative_to(self.fixture / 'ur-ai') or
                                target.is_relative_to(self.fixture / 'ur-view'))
                with contextlib.redirect_stdout(io.StringIO()):
                    generator.write_reference_files(target, domain, {}, {})
        self.assertEqual({p.name for p in self.fixture.iterdir()},
                         {'ur-iot', 'ur-org-manage', 'ur-ai', 'ur-view'})

    def test_public_sync_and_repeat(self):
        """公开同步排除三个专题及其导航，迁移旧入口并保留仓库文件和其他内容。"""
        sync = load_script('sync-skills')
        for name in sync.GROUPS:
            if name != 'references':
                entry = self.fixture / name / 'SKILL.md'
                entry.parent.mkdir(parents=True, exist_ok=True)
                entry.write_text('# 业务组\n')
        (self.fixture / 'SKILL.md').write_text('# 总入口\n')
        for name in sync.PRIVATE_TOPICS:
            entry = self.fixture / 'ur-iot' / name / 'SKILL.md'
            entry.parent.mkdir()
            entry.write_text('# 内部指南\n')
            with (self.fixture / 'ur-iot/SKILL.md').open('a') as stream:
                stream.write(f'- [{name}]({name}/SKILL.md)\n')
        target = self.fixture.parent / (self.fixture.name + '-public')
        target.mkdir()
        self.addCleanup(shutil.rmtree, target)
        (target / '.git').write_text('gitdir: repository')
        (target / 'keep.txt').write_text('其他内容')
        (target / 'ur-device').mkdir()
        (target / 'ur-device/GUIDE.md').write_text('旧入口')
        for _ in range(2):
            sync.synchronize(self.fixture, target, public=True)
            self.assertFalse((target / 'ur-device').exists())
            for name in sync.PRIVATE_TOPICS:
                self.assertFalse((target / 'ur-iot' / name).exists())
                self.assertNotIn(name, (target / 'ur-iot/SKILL.md').read_text())
            self.assertEqual((target / 'keep.txt').read_text(), '其他内容')
            self.assertEqual((target / '.git').read_text(), 'gitdir: repository')

    def test_private_guides_survive_source_mirror(self):
        """公开源更新 CLI 时保留客户端专题，并在组入口恢复一次有效导航。"""
        sync = load_script('sync-skills')
        source = self.fixture / 'source'
        target = self.fixture / 'cli'
        source.mkdir()
        (source / 'SKILL.md').write_text('# 总入口\n')
        for name in sync.GROUPS[:-1]:
            entry = source / name / 'SKILL.md'
            entry.parent.mkdir()
            entry.write_text('# 业务组\n')
        for name in sync.PRIVATE_TOPICS:
            entry = target / 'ur-iot' / name / 'SKILL.md'
            entry.parent.mkdir(parents=True)
            entry.write_text('内部资料')
        for _ in range(2):
            sync.synchronize(source, target, preserve_client_guides=True)
            for name in sync.PRIVATE_TOPICS:
                self.assertEqual((target / 'ur-iot' / name / 'SKILL.md').read_text(), '内部资料')
                self.assertEqual((target / 'ur-iot/SKILL.md').read_text().count(f']({name}/SKILL.md)'), 1)

    @unittest.skipUnless(shutil.which('rsync'), '同步脚本预览依赖 rsync')
    def test_update_script_uses_grouped_source(self):
        """实际更新入口在隔离源、目标和 Swagger 中执行两次，验证公开边界及旧路径清理。"""
        cli = self.fixture / 'cli'
        source = self.fixture / 'source'
        public = self.fixture / 'public'
        scripts = cli / 'scripts'
        scripts.mkdir(parents=True)
        shutil.copytree(SKILLS, source)
        shutil.copytree(SKILLS, cli / 'skill')
        for name in ('update-skills.sh', 'sync-skills.py', 'generate-api-lists.py'):
            shutil.copy2(ROOT / 'scripts' / name, scripts / name)
        swagger = self.fixture / 'swagger'
        swagger.mkdir()
        for name in ('core-api.json', 'things-api.json'):
            (swagger / name).write_text('{"paths": {}}')
        (cli / 'skill/ur-device').mkdir()
        (cli / 'skill/ur-device/GUIDE.md').write_text('旧入口')
        environment = dict(os.environ, UR_SKILLS_REPO=str(public),
                           UR_SKILLS_SOURCE=str(source), UR_SWAGGER_DIR=str(swagger))
        for _ in range(2):
            subprocess.run(['bash', str(scripts / 'update-skills.sh'), '--apply'],
                           check=True, capture_output=True, text=True, env=environment)
            self.assertFalse((cli / 'skill/ur-device').exists())
            self.assertFalse((cli / 'skill/.git').exists())
            self.assertEqual((cli / 'skill/SKILL.md').read_bytes(), (source / 'SKILL.md').read_bytes())
            for name in ('ur-iot-client', 'ur-iot-context', 'ur-iot-device'):
                self.assertTrue((cli / 'skill/ur-iot' / name / 'SKILL.md').is_file())
                self.assertFalse((public / 'ur-iot' / name).exists())

    def test_sync_rolls_back_failed_replacement(self):
        """替换中途失败时恢复全部原目录，包括尚未发布的客户端资料。"""
        sync = load_script('sync-skills')
        source = self.fixture / 'source'
        target = self.fixture / 'target'
        shutil.copytree(SKILLS, source)
        shutil.copytree(SKILLS, target)
        original = {str(p.relative_to(target)): p.read_bytes() for p in target.rglob('*') if p.is_file()}
        (source / 'SKILL.md').write_text('新的总入口')
        rename = Path.rename

        def fail_group(path, destination):
            """模拟 stage 内 AI 目录替换失败，备份恢复调用仍使用真实 rename。"""
            if path.parent.name == 'stage' and path.name == 'ur-ai':
                raise OSError('模拟磁盘替换失败')
            return rename(path, destination)

        with patch.object(Path, 'rename', fail_group):
            with self.assertRaises(OSError):
                sync.synchronize(source, target, preserve_client_guides=True)
        actual = {str(p.relative_to(target)): p.read_bytes() for p in target.rglob('*') if p.is_file()}
        self.assertEqual(actual, original)


if __name__ == '__main__':
    unittest.main()
