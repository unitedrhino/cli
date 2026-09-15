#!/usr/bin/env python3
"""文件说明：离开业务仓库验证两套模板完整性、运行包白名单和体积；产物只写指定临时目录。"""
import argparse
import re
import shutil
import subprocess
import tempfile
from pathlib import Path
from zipfile import ZipFile


def validate(root: Path, output: Path, forbidden: list[str]) -> None:
    """独立复制模板后打包；forbidden 为调用者临时提供的去标识扫描词，不在源码保存项目词库。"""
    output.mkdir(parents=True, exist_ok=True)
    for name in ('building', 'power-station'):
        source = root / name
        target = output / name
        shutil.copytree(source, target, dirs_exist_ok=True,
                        ignore=shutil.ignore_patterns('.temp', '__pycache__', '*.zip', 'node_modules'))
        assert (target / 'index.html').is_file() and (target / 'LICENSE.txt').is_file(), name
        for file in target.rglob('*'):
            if file.is_file():
                raw = file.read_bytes()
                for word in forbidden:
                    assert word.encode() not in raw, f'未去除指定标识：{file.relative_to(target)}'
        entry = (target / 'index.html').read_text()
        for resource in re.findall(r'(?:src|href)="([^"]+)"', entry):
            if resource.startswith(('data:', '#')):
                continue
            assert not re.match(r'(?:https?:)?//', resource), '运行依赖必须本地化'
            assert (target / resource).is_file(), resource
        archive = output / (name + '.zip')
        subprocess.run(['python3', str(target / 'package.py'), '--output', str(archive)], check=True, cwd=target)
        assert archive.stat().st_size <= 2_000_000, '运行包超出2 MB'
        with ZipFile(archive) as package:
            assert 'index.html' in package.namelist()
            for member in package.namelist():
                assert not any(part in member for part in ('..', '.temp', 'node_modules', '.test.', '.cjs', 'README')), member
                assert package.read(member) == (target / member).read_bytes(), member
        print(f'{name}: 独立打包通过，{archive.stat().st_size:,} bytes')


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--out', type=Path, help='临时输出目录；省略则测试结束自动删除')
    parser.add_argument('--forbid', action='append', default=[], help='临时额外扫描词，可重复；不要提交业务词库')
    args = parser.parse_args()
    root = Path(__file__).resolve().parent
    if args.out:
        validate(root, args.out.resolve(), args.forbid)
    else:
        with tempfile.TemporaryDirectory(prefix='scene-template-') as directory:
            validate(root, Path(directory), args.forbid)
