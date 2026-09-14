#!/usr/bin/env python3
"""文件说明：可重复打包建筑能耗模板运行资源，并生成双击可用的离线 HTML；无需前端构建依赖。"""
import argparse
import re
from pathlib import Path
from zipfile import ZipFile, ZipInfo, ZIP_DEFLATED


def build(output: Path, preview: Path | None) -> None:
    """按运行资源白名单生成 ZIP；preview 指定可选的离线单文件，禁止将文档和测试打进场景包。"""
    root = Path(__file__).resolve().parent
    files = [root / 'index.html', root / 'LICENSE.txt']
    for directory in ('css', 'js', 'libs'):
        files.extend(sorted((root / directory).iterdir()))
    output.parent.mkdir(parents=True, exist_ok=True)
    with ZipFile(output, 'w', ZIP_DEFLATED, compresslevel=9) as archive:
        for file in files:
            if not file.is_file() or file.suffix not in ('.html', '.css', '.js', '.txt'):
                raise ValueError(f'非运行资源：{file}')
            item = ZipInfo(file.relative_to(root).as_posix(), date_time=(2026, 1, 1, 0, 0, 0))
            item.compress_type = ZIP_DEFLATED
            item.external_attr = 0o100644 << 16
            archive.writestr(item, file.read_bytes(), compresslevel=9)
    if output.stat().st_size > 2_000_000:
        raise ValueError('ZIP 超出 2 MB 预算')
    if preview:
        html = (root / 'index.html').read_text()
        html = re.sub(r'<link rel="stylesheet" href="([^"]+)">', lambda match: '<style>' + (root / match[1]).read_text() + '</style>', html)
        html = re.sub(r'<script src="([^"]+)"></script>', lambda match: '<script>' + (root / match[1]).read_text().replace('</script', '<\\/script') + '</script>', html)
        preview.parent.mkdir(parents=True, exist_ok=True)
        preview.write_text(html)
    print(f'{output}: {output.stat().st_size:,} bytes / {len(files)} files')


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--output', type=Path, default=Path(__file__).resolve().parent / '.temp' / 'building.zip')
    parser.add_argument('--preview', type=Path)
    args = parser.parse_args()
    build(args.output, args.preview)
