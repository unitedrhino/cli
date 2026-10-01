#!/usr/bin/env python3
"""验证手写技能经平台打包、独立发布和 npm 白名单后保持完整；不编译或发布。

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
    """以固件、OTA 和嵌套场景夹具检查真实分发脚本。"""

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
        # 设备固件是非 Swagger 手写技能，入口与全部参考资料都必须逐字节分发。
        firmware_root = ROOT / 'skill/device-firmware'
        self.firmware_payloads = {
            str(path.relative_to(firmware_root)): path.read_bytes()
            for path in firmware_root.rglob('*') if path.is_file()
        }
        for name, content in self.firmware_payloads.items():
            self.write(self.root / 'skill/device-firmware' / name, content)
        # OTA 的 API 参考可生成，但平台工作流入口是手写内容，整个目录必须完整分发。
        ota_root = ROOT / 'skill/ur-ota'
        self.ota_payloads = {
            str(path.relative_to(ota_root)): path.read_bytes()
            for path in ota_root.rglob('*') if path.is_file()
        }
        for name, content in self.ota_payloads.items():
            self.write(self.root / 'skill/ur-ota' / name, content)
        self.write(self.root / 'skill/SKILL.md', '# 统一技能\n')
        self.write(self.root / 'references/README.md', '# API 参考\n')
        # 真实手写指南使用 persona 约定的域级路径，不能只依赖扁平兼容副本。
        self.guide = (ROOT / 'skill/ur-device/references/device-control.md').read_bytes()
        self.write(self.root / 'skill/ur-device/references/device-control.md', self.guide)
        self.ai_voice_guide = (ROOT / 'skill/ur-ai/references/device-voice.md').read_bytes()
        self.write(self.root / 'skill/ur-ai/references/device-voice.md', self.ai_voice_guide)
        self.photo_guide = (ROOT / 'skill/device-firmware/references/photo-vision.md').read_bytes()
        self.write(self.root / 'skill/ur-device/references/shared.md', '设备同名指南\n')
        self.write(self.root / 'skill/ur-product/references/shared.md', '产品同名指南\n')
        self.write(self.root / 'references/quick-reference.md', (ROOT / 'references/quick-reference.md').read_bytes())
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

    def assert_firmware_tree(self, destination):
        """逐字节核对设备固件技能入口和参考资料。"""
        for name, expected in self.firmware_payloads.items():
            actual = destination / name
            self.assertTrue(actual.is_file(), f'分发缺少 {actual}')
            self.assertEqual(actual.read_bytes(), expected)

    def assert_ota_tree(self, destination):
        """逐字节核对 OTA 手写入口和 API 参考。"""
        for name, expected in self.ota_payloads.items():
            actual = destination / name
            self.assertTrue(actual.is_file(), f'分发缺少 {actual}')
            self.assertEqual(actual.read_bytes(), expected)

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
    (target / 'references').mkdir(exist_ok=True)
    (target / 'references/generated-index.md').write_text('generated')
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
            api_root = output / 'x64-linux/skill/ur-api'
            self.assert_firmware_tree(api_root / 'device-firmware')
            self.assert_ota_tree(api_root / 'ur-ota')
            self.assertEqual((api_root / 'ur-device/references/device-control.md').read_bytes(), self.guide)
            self.assertEqual((api_root / 'ur-ai/references/device-voice.md').read_bytes(),
                             self.ai_voice_guide)
            self.assertEqual((api_root / 'device-firmware/references/photo-vision.md').read_bytes(),
                             self.photo_guide)
            self.assertEqual((api_root / 'ur-device/references/shared.md').read_text(), '设备同名指南\n')
            self.assertEqual((api_root / 'ur-product/references/shared.md').read_text(), '产品同名指南\n')
            self.assertEqual((api_root / 'SKILL.md').read_text().count('(ur-device/references/device-control.md)'), 1)
            self.assertFalse((api_root / 'references/references').exists())
            self.assertEqual((api_root / 'references/quick-reference.md').read_bytes(),
                             (ROOT / 'references/quick-reference.md').read_bytes())
            self.assertFalse((output / 'x64-linux/skill/ur-api/ur-view/ur-view').exists())
            self.assertEqual((output / 'x64-linux/skill/ur-api/SKILL.md').read_text()
                             .count('[大屏技能](ur-view/SKILL.md)'), 1)
            self.assertEqual((api_root / 'SKILL.md').read_text()
                             .count('[设备固件技能](device-firmware/SKILL.md)'), 1)
            self.assertEqual((api_root / 'SKILL.md').read_text()
                             .count('(ur-ai/references/device-voice.md)'), 1)
            self.assertEqual((api_root / 'SKILL.md').read_text()
                             .count('[OTA 管理技能](ur-ota/SKILL.md)'), 1)
        # --dry-run 不生成归档或上传；--ignore-scripts 避免运行 npm 发布构建。
        result = subprocess.run(['npm', 'pack', '--dry-run', '--json', '--ignore-scripts'],
                                cwd=self.root / 'npm-package', check=True, capture_output=True, text=True)
        files = {entry['path'] for entry in json.loads(result.stdout)[0]['files']}
        self.assertIn('ur-api/ur-device/references/device-control.md', files)
        self.assertIn('ur-api/ur-ai/references/device-voice.md', files)
        self.assertIn('ur-api/device-firmware/references/photo-vision.md', files)
        self.assertIn('ur-api/ur-product/references/shared.md', files)
        for name in self.firmware_payloads:
            self.assertIn('ur-api/device-firmware/' + name, files, f'npm 包漏掉 {name}')
        for name in self.ota_payloads:
            self.assertIn('ur-api/ur-ota/' + name, files, f'npm 包漏掉 OTA 文件 {name}')
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

    def test_existing_reference_directory(self):
        """已有导出目录重复同步时保留生成索引，且路径不嵌套、导航不重复。"""
        script = (ROOT / 'scripts/package-skill.sh').read_text()
        begin = script.index('  mkdir -p "${api_skill_dir}/references"')
        end = script.index('  # Swagger 导出不包含手写场景模板', begin)
        destination = self.root / 'existing-references/ur-api'
        self.write(destination / 'SKILL.md', '# API 索引\n')
        self.write(destination / 'references/generated-index.md', '保留生成索引\n')
        command = 'ROOT="$1"; api_skill_dir="$2"\ncopy_references() {\n' + script[begin:end] + '\n}\ncopy_references'
        for _ in range(2):
            subprocess.run(['bash', '-c', command, 'reference-distribution', str(self.root), str(destination)],
                           check=True, capture_output=True, text=True)
            self.assertEqual((destination / 'ur-device/references/device-control.md').read_bytes(), self.guide)
            self.assertEqual((destination / 'ur-ai/references/device-voice.md').read_bytes(),
                             self.ai_voice_guide)
            self.assertEqual((destination / 'references/generated-index.md').read_text(), '保留生成索引\n')
            self.assertFalse((destination / 'references/references').exists())
            self.assertEqual((destination / 'SKILL.md').read_text().count('(ur-device/references/device-control.md)'), 1)
            self.assert_firmware_tree(destination / 'device-firmware')
            self.assert_ota_tree(destination / 'ur-ota')
            self.assertEqual((destination / 'SKILL.md').read_text()
                             .count('[设备固件技能](device-firmware/SKILL.md)'), 1)
            self.assertEqual((destination / 'SKILL.md').read_text()
                             .count('(ur-ai/references/device-voice.md)'), 1)
            self.assertEqual((destination / 'SKILL.md').read_text()
                             .count('[OTA 管理技能](ur-ota/SKILL.md)'), 1)

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
        self.assert_firmware_tree(destination / 'device-firmware')
        self.assert_ota_tree(destination / 'ur-ota')
        self.assertEqual((destination / 'ur-device/references/device-control.md').read_bytes(), self.guide)
        self.assertEqual((destination / 'ur-ai/references/device-voice.md').read_bytes(),
                         self.ai_voice_guide)

    def test_release_alias_names(self):
        """验证 CLI 与 Skills 的 latest 别名均移除版本号且不会覆盖版本资产。"""
        script = (ROOT / 'scripts/release.sh').read_text()
        begin = script.index('release_alias_name() {')
        end = script.index('\n}\n', begin) + 3
        command = ('VERSION=v0.8.0\n' + script[begin:end] +
                   '\nrelease_alias_name ur-cli-v0.8.0-Linux-x86_64.tar.gz\n' +
                   'release_alias_name ur-api-skills-v0.8.0.zip')
        result = subprocess.run(['bash', '-c', command], check=True, capture_output=True, text=True)
        self.assertEqual(result.stdout.splitlines(), [
            'ur-cli-Linux-x86_64.tar.gz',
            'ur-api-skills.zip',
        ])

    def test_device_intent_reference_contract(self):
        """验证两种发行入口的导航和模拟控制合同没有回退到旧样例。"""
        guide = (ROOT / 'skill/ur-device/references/device-control.md').read_text()
        for expected in ('--shadow-control 4', '/api/v1/things/device/simulate/report',
                         '`--data` 的属性值必须是字符串', 'proc.exited', '--project-id',
                         '先按意图区分'):
            self.assertIn(expected, guide)
        for relative in ('references/quick-reference.md', 'skill/references/quick-reference.md'):
            self.assertIn('../ur-device/references/device-control.md', (ROOT / relative).read_text())
        self.assertIn('(ur-device/references/device-control.md)', (ROOT / 'skill/SKILL.md').read_text())

    def test_manual_device_ai_navigation_generator(self):
        """生成型技能的语音和视觉指南入口必须可重入，且链接指向真实文件。"""
        skill_root = self.root / 'generated-navigation'
        for domain in ('ur-ai', 'ur-device-debug', 'ur-product'):
            self.write(skill_root / domain / 'SKILL.md', '# 测试技能\n\n## 典型业务场景\n')
        generator = ROOT / 'scripts/generate-api-lists.py'
        command = ['python3', str(generator), '--manual-guides-only', '--skill-dir', str(skill_root)]
        subprocess.run(command, check=True, capture_output=True, text=True)
        subprocess.run(command, check=True, capture_output=True, text=True)

        expected_links = {
            'ur-ai': ('(references/device-voice.md)',
                      '(references/device-voice.md#拍照识图与图片输入)'),
            'ur-device-debug': ('(../ur-ai/references/device-voice.md)',
                                '(../device-firmware/references/photo-vision.md)'),
            'ur-product': ('(../device-firmware/references/voice-ai.md)',
                           '(../device-firmware/references/photo-vision.md)'),
        }
        for domain, links in expected_links.items():
            content = (skill_root / domain / 'SKILL.md').read_text()
            self.assertEqual(content.count(f'<!-- MANUAL_GUIDES:{domain} -->'), 1)
            for expected_link in links:
                self.assertEqual(content.count(expected_link), 1)

        self.assertTrue((ROOT / 'skill/ur-ai/references/device-voice.md').is_file())
        self.assertTrue((ROOT / 'skill/device-firmware/references/voice-ai.md').is_file())
        self.assertTrue((ROOT / 'skill/device-firmware/references/photo-vision.md').is_file())

    def test_repeatable_photo_vision_reference_contract(self):
        """视觉技能必须保留两条图片入口、凭据边界和三层可复测门禁。"""
        firmware_guide = (
            ROOT / 'skill/device-firmware/references/photo-vision.md'
        ).read_text()
        platform_guide = (ROOT / 'skill/ur-ai/references/device-voice.md').read_text()
        for expected in (
            'Test(TakePhotoEndToEnd|ImageInputEndToEnd|ImageInputDuringVoiceAudioStop|EmojiEmotionText|EmojiNotSentForPlainQuestion)',
            '"actionID": "takePhoto"',
            '"type": "image_url"',
            '上传成功、失败或 MQTT 断线后都立即清零副本',
            '关键纯逻辑用例至少连续 10 次',
            '真实按键、镜头画面、屏幕表情、扬声器和断电必须由现场观察',
        ):
            self.assertIn(expected, firmware_guide)
        for expected in (
            '`text`、`voice`、`image_input`',
            '同步生成',
            '硬编码历史 ID',
            '`deviceTakePhoto`',
            '`inputSend`',
        ):
            self.assertIn(expected, platform_guide)

    def test_repeatable_voice_e2e_reference_contract(self):
        """语音技能必须保留 16kHz 协议基准、分层解码和可循环的真机 runner。"""
        firmware_guide = (
            ROOT / 'skill/device-firmware/references/voice-ai.md'
        ).read_text()
        platform_guide = (ROOT / 'skill/ur-ai/references/device-voice.md').read_text()
        for expected in (
            'DEVICESIM_AUDIO_SAMPLE_RATE=16000',
            'tools/devicesim/cmd/opusfixture',
            'ENABLE_UR_AI_E2E_TEST=1',
            '`VOICE-HW-002` 重复稳定性',
            'STT 必须命中固定语料关键词',
            'run_ur_ai_e2e.py',
            '--repeat 5',
            'Opus 解码器可从同一 16 kHz 码流原生输出',
            '不能靠事后检查或把协议改成 24 kHz 掩盖',
            'Failed to resample output audio',
            'AudioStop→首帧小于 8 秒',
            '不能相减计时原点不同的 `elapsed`',
        ):
            self.assertIn(expected, firmware_guide)
        for expected in (
            'DEVICESIM_AUDIO_SAMPLE_RATE=16000',
            'run_ur_ai_e2e.py --port <serial-port> --repeat 5',
            'Watcher 保持 16 kHz 会话并由 Opus 解码器直接输出 24 kHz PCM',
            'device-firmware/references/voice-ai.md',
        ):
            self.assertIn(expected, platform_guide)


if __name__ == '__main__':
    unittest.main()
