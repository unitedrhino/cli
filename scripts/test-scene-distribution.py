#!/usr/bin/env python3
"""验证场景源码经平台打包、独立技能发布和 npm 白名单后保持完整；不编译或发布。

默认使用小型夹具；设置 SCENE_DISTRIBUTION_SOURCE 为 ur-view 目录可逐字节验证完整真实模板。
"""

import json
import os
from pathlib import Path
import subprocess
import tempfile
import unittest

# 仓库路径仅用于读取正式分发入口，所有夹具与输出均位于隔离临时目录。
ROOT = Path(__file__).resolve().parents[1]


class SceneDistributionTest(unittest.TestCase):
    """以嵌套资源夹具检查真实分发脚本，避免只检查扩展名字符串。"""

    def setUp(self):
        """创建含脚本、纹理、许可证和两套场景的最小技能仓库。"""
        (ROOT / '.temp').mkdir(exist_ok=True)
        self.temporary = tempfile.TemporaryDirectory(prefix='scene-distribution-', dir=ROOT / '.temp')
        self.addCleanup(self.temporary.cleanup)
        self.root = Path(self.temporary.name) / 'cli'
        self.root.mkdir()
        self.payloads = {'SKILL.md': '# 大屏技能\n'}
        for scene in ('building', 'power-station'):
            for asset in ('index.html', 'js/data.js', 'css/style.css', 'libs/three.min.js',
                          'libs/LICENSE', 'assets/textures/panel.png', 'package.py', 'tests/data.test.cjs'):
                self.payloads[f'assets/scene-templates/{scene}/{asset}'] = f'{scene}/{asset}\n'
        # 接收实际 ur-view 源目录时，全部文件参与发布检查，避免只测代表性后缀。
        source = os.environ.get('SCENE_DISTRIBUTION_SOURCE')
        if source:
            source_path = Path(source).resolve()
            for scene in ('building', 'power-station'):
                self.assertTrue((source_path / 'assets/scene-templates' / scene / 'index.html').is_file())
            self.payloads = {str(path.relative_to(source_path)): path.read_bytes()
                             for path in source_path.rglob('*') if path.is_file()}
        for name, content in self.payloads.items():
            self.write(self.root / 'skill/ur-view' / name, content)
        self.write(self.root / 'skill/SKILL.md', '# 统一技能\n')
        self.write(self.root / 'references/README.md', '# API 参考\n')
        self.write(self.root / 'npm-package/package.json', (ROOT / 'npm-package/package.json').read_text())
        self.write(self.root / 'scripts/package-skill.sh', (ROOT / 'scripts/package-skill.sh').read_text())

    def write(self, path, content):
        """写入隔离夹具文件并自动创建父目录。"""
        path.parent.mkdir(parents=True, exist_ok=True)
        path.write_bytes(content if isinstance(content, bytes) else content.encode())

    def assert_tree(self, destination):
        """逐字节核对整套场景文件，目录存在不能替代内容完整性。"""
        for name, expected in self.payloads.items():
            actual = destination / name
            self.assertTrue(actual.is_file(), f'分发缺少 {actual}')
            self.assertEqual(actual.read_bytes(), expected if isinstance(expected, bytes) else expected.encode())

    def test_platform_package_and_npm(self):
        """替代耗时的 Go 编译与 Swagger 导出，仅执行正式打包和 npm 文件选择。"""
        # 编译替身只生成占位产物；资源复制、包装脚本和 npm 打包仍走正式入口。
        fake_go = self.root / 'tools/go'
        self.write(fake_go, '''#!/usr/bin/env python3
import pathlib, sys
args = sys.argv[1:]
if args[0] == 'build':
    target = pathlib.Path(args[args.index('-o') + 1])
    target.parent.mkdir(parents=True, exist_ok=True)
    target.write_text('test binary')
elif args[0] == 'run':
    target = pathlib.Path(args[args.index('--output') + 1])
    target.mkdir(parents=True, exist_ok=True)
    (target / 'SKILL.md').write_text('# API 索引\\n')
else:
    raise SystemExit('unexpected go invocation')
''')
        fake_go.chmod(0o755)
        environment = dict(os.environ, PATH=f'{fake_go.parent}{os.pathsep}{os.environ["PATH"]}')
        output = self.root / 'output'
        # 同一输出位置连续打包，入口与目录必须保持幂等。
        for _ in range(2):
            subprocess.run(['bash', str(self.root / 'scripts/package-skill.sh'), str(output),
                            '--arch', 'linux-amd64'], env=environment, check=True, capture_output=True, text=True)
            self.assert_tree(output / 'x64-linux/skill/ur-api/ur-view')
            self.assertFalse((output / 'x64-linux/skill/ur-api/ur-view/ur-view').exists())
            self.assertEqual((output / 'x64-linux/skill/ur-api/SKILL.md').read_text()
                             .count('[大屏技能](ur-view/SKILL.md)'), 1)
        # --dry-run 不生成归档或上传；--ignore-scripts 避免运行 npm 发布构建。
        result = subprocess.run(['npm', 'pack', '--dry-run', '--json', '--ignore-scripts'],
                                cwd=self.root / 'npm-package', check=True, capture_output=True, text=True)
        files = {entry['path'] for entry in json.loads(result.stdout)[0]['files']}
        for name in self.payloads:
            # npm 固定排除 Git 忽略规则；它不属于运行、复制或打包所需源码。
            if Path(name).name == '.gitignore':
                continue
            self.assertIn('ur-api/ur-view/' + name, files, f'npm 包漏掉 {name}')

    def test_existing_scene_directory(self):
        """对已存在的技能输出连续执行实际复制段，防止目录嵌套和重复导航。"""
        script = (ROOT / 'scripts/package-skill.sh').read_text()
        begin = script.index('  # Swagger 导出不包含手写场景模板')
        end = script.index('  # 保留顶层 SKILL.md', begin)
        destination = self.root / 'existing/ur-api'
        self.write(destination / 'SKILL.md', '# API 索引\n')
        self.write(destination / 'ur-view/preserved.txt', '已有其它内容\n')
        command = 'ROOT="$1"; api_skill_dir="$2"\n' + script[begin:end]
        for _ in range(2):
            subprocess.run(['bash', '-c', command, 'scene-distribution', str(self.root), str(destination)],
                           check=True, capture_output=True, text=True)
            self.assert_tree(destination / 'ur-view')
            self.assertFalse((destination / 'ur-view/ur-view').exists())
            self.assertEqual((destination / 'SKILL.md').read_text().count('[大屏技能](ur-view/SKILL.md)'), 1)
            self.assertEqual((destination / 'ur-view/preserved.txt').read_text(), '已有其它内容\n')

    def test_release_copy(self):
        """执行 release.sh 的实际共享复制函数，覆盖平台包和独立 skills ZIP 的资源来源。"""
        script = (ROOT / 'scripts/release.sh').read_text()
        begin = script.index('copy_skills_into() {')
        end = script.index('\n}\n', begin) + 3
        destination = self.root / 'release/ur-api'
        subprocess.run(['bash', '-c', script[begin:end] + '\nROOT="$1"; copy_skills_into "$2"',
                        'scene-distribution', str(self.root), str(destination)],
                       check=True, capture_output=True, text=True)
        self.assert_tree(destination / 'ur-view')


if __name__ == '__main__':
    unittest.main()
