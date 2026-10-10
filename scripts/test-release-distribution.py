#!/usr/bin/env python3
"""执行真实发版脚本的隔离夹具，验证 Harbor 必需检查、分发及失败保留；不访问网络。"""

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
import os,pathlib,sys
pathlib.Path(os.environ['RELEASE_TEST_GO_LOG']).write_text('called')
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
import os,pathlib,sys
pathlib.Path(os.environ['RELEASE_TEST_CURL_LOG']).write_text('called')
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
if sys.argv[1:] == ['buildx', 'version']:
    raise SystemExit(1 if os.environ.get('RELEASE_TEST_BUILDX_FAIL')=='1' else 0)
if sys.argv[1:] == ['info']:
    raise SystemExit(1 if os.environ.get('RELEASE_TEST_DAEMON_FAIL')=='1' else 0)
if sys.argv[1:3] != ['buildx', 'build']:
    raise SystemExit('unexpected docker command')
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

    def execute_release(self, fail=False, overrides=None):
        """执行真实脚本；fail 模拟推送失败，overrides 覆盖前置环境，返回结果和上传记录。"""
        log = self.root / 'harbor.log'
        environment = dict(os.environ, PATH=str(self.tools) + ':' + os.environ['PATH'],
                           GITHUB_TOKEN='fake', GITEE_TOKEN='fake', PARALLEL='1',
                           KEEP_RELEASES='0', SKIP_HARBOR='0', TMPDIR=str(self.root),
                           RELEASE_TEST_LOG=str(log), RELEASE_TEST_HARBOR_FAIL='1' if fail else '0',
                           RELEASE_TEST_GO_LOG=str(self.root / 'go.log'),
                           RELEASE_TEST_CURL_LOG=str(self.root / 'curl.log'),
                           RELEASE_TEST_BUILDX_FAIL='0', RELEASE_TEST_DAEMON_FAIL='0')
        environment.update(overrides or {})
        result = subprocess.run(['bash', str(self.root / 'scripts/release.sh'), 'v0.8.6'],
                                env=environment, capture_output=True, text=True)
        return result, log

    def assert_preflight_failure(self, overrides, message):
        """前置检查失败须保留旧产物，并在编译或访问发布平台前停止。"""
        marker = self.root / 'dist/release-v0.8.6/old-package'
        marker.parent.mkdir(parents=True)
        marker.write_text('retained')
        result, log = self.execute_release(overrides=overrides)
        self.assertNotEqual(result.returncode, 0)
        self.assertIn(message, result.stdout + result.stderr)
        self.assertEqual(marker.read_text(), 'retained')
        for path in [log, self.root / 'go.log', self.root / 'curl.log']:
            self.assertFalse(path.exists(), path)

    def test_skip_harbor_is_rejected_before_build(self):
        """正式发布不能通过跳过 Harbor 返回成功，也不能提前清理旧产物。"""
        self.assert_preflight_failure({'SKIP_HARBOR': '1'}, '禁止跳过 Harbor')

    def test_missing_docker_is_rejected_before_build(self):
        """隔离 PATH 中不提供 Docker，确认脚本明确报错并在构建前停止。"""
        minimal_tools = self.root / 'minimal-tools'
        minimal_tools.mkdir()
        for name in ['bash', 'dirname']:
            (minimal_tools / name).symlink_to(shutil.which(name))
        self.assert_preflight_failure({'PATH': str(minimal_tools)}, '未安装 Docker')

    def test_unavailable_buildx_is_rejected_before_build(self):
        """缺少 buildx 插件时停止发布，不产生 Release 或丢弃已有包。"""
        self.assert_preflight_failure({'RELEASE_TEST_BUILDX_FAIL': '1'}, 'buildx 不可用')

    def test_unavailable_daemon_is_rejected_before_build(self):
        """Docker daemon 不可访问时停止发布，避免构建完成后才发现无法分发。"""
        self.assert_preflight_failure({'RELEASE_TEST_DAEMON_FAIL': '1'}, 'Docker daemon 不可用')

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
