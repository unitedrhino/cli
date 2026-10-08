#!/usr/bin/env python3
"""执行真实发版脚本的隔离夹具，验证 Harbor 分发先于清理、失败保留安装包；不访问网络。"""

import os
from pathlib import Path
import shutil
import subprocess
import tempfile
import unittest

# ROOT 只用于读取正式发版脚本，所有替身和产物均放在 .temp 中。
ROOT = Path(__file__).resolve().parents[1]


class ReleaseDistributionTest(unittest.TestCase):
    """替换编译与外部平台命令，保持正式发版流程及文件操作不变。"""

    def setUp(self):
        """创建五平台编译夹具及只在本地返回响应的分发工具。"""
        (ROOT / '.temp').mkdir(exist_ok=True)
        self.temporary = tempfile.TemporaryDirectory(dir=ROOT / '.temp')
        self.addCleanup(self.temporary.cleanup)
        self.root = Path(self.temporary.name)
        (self.root / 'scripts').mkdir()
        shutil.copy2(ROOT / 'scripts/release.sh', self.root / 'scripts/release.sh')
        (self.root / 'skill').mkdir()
        (self.root / 'skill/SKILL.md').write_text('# 技能\n')
        self.tools = self.root / 'tools'
        self.tools.mkdir()
        self.write_tool('go', '''#!/usr/bin/env python3
import pathlib,sys
args=sys.argv[1:]
if args == ['tool', 'dist', 'list']:
    print('linux/amd64\\nlinux/arm64\\ndarwin/amd64\\ndarwin/arm64\\nwindows/amd64')
elif args[0] == 'build':
    target=pathlib.Path(args[args.index('-o')+1]);target.parent.mkdir(parents=True,exist_ok=True)
    target.write_text('binary')
else:
    raise SystemExit('unexpected go command')
''')
        self.write_tool('curl', '''#!/usr/bin/env python3
import sys
args=sys.argv[1:]
if any('api.github.com/repos/' in arg for arg in args):
    print('{"upload_url": "https://uploads.example.invalid/release"}')
elif any('gitee.com/api/' in arg and arg.endswith('/releases') for arg in args):
    print('{"id":1}')
else:
    print('{}')
''')
        self.write_tool('docker', '''#!/usr/bin/env python3
import os,pathlib,sys
context=pathlib.Path(sys.argv[-1])
packages=list(context.glob('*.zip'))+list(context.glob('*.tar.gz'))
if not packages or not packages[0].is_file():
    raise SystemExit(2)
with pathlib.Path(os.environ['RELEASE_TEST_LOG']).open('a') as stream:
    stream.write(packages[0].name+'\\n')
raise SystemExit(1 if os.environ.get('RELEASE_TEST_HARBOR_FAIL')=='1' else 0)
''')

    def write_tool(self, name, content):
        """创建只在测试 PATH 中生效的可执行命令，内容由测试固定定义。"""
        target = self.tools / name
        target.write_text(content)
        target.chmod(0o755)

    def execute_release(self, fail=False):
        """执行真实发版脚本；fail 模拟 Harbor 失败，返回进程结果和上传记录。"""
        log = self.root / 'harbor.log'
        environment = dict(os.environ, PATH=str(self.tools) + ':' + os.environ['PATH'],
                           GITHUB_TOKEN='fake', GITEE_TOKEN='fake', PARALLEL='1',
                           KEEP_RELEASES='0', SKIP_HARBOR='0', TMPDIR=str(self.root),
                           RELEASE_TEST_LOG=str(log), RELEASE_TEST_HARBOR_FAIL='1' if fail else '0')
        result = subprocess.run(['bash', str(self.root / 'scripts/release.sh'), 'v0.8.6'],
                                env=environment, capture_output=True, text=True)
        return result, log

    def test_harbor_before_cleanup(self):
        """默认清理策略下五平台均先取得安装包，全部分发完成后才删除构建目录。"""
        result, log = self.execute_release()
        self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
        self.assertEqual(len(log.read_text().splitlines()), 5)
        self.assertFalse((self.root / 'dist/release-v0.8.6').exists())

    def test_harbor_failure_preserves_packages(self):
        """Harbor 失败必须返回非零状态并保留包与校验文件，便于原节点补传。"""
        result, log = self.execute_release(fail=True)
        self.assertNotEqual(result.returncode, 0)
        self.assertEqual(len(log.read_text().splitlines()), 5)
        self.assertTrue((self.root / 'dist/release-v0.8.6/packages/sha256sums.txt').is_file())
        self.assertEqual(len(list((self.root / 'dist/release-v0.8.6/packages').glob('ur-cli-*'))), 5)


if __name__ == '__main__':
    unittest.main()
