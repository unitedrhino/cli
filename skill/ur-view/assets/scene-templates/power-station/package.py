#!/usr/bin/env python3
"""文件说明：可重复打包配电站运行资源，校验体积并按需输出离线单文件预览。
测试脚本、开发说明与截图不进入托管 ZIP；仅使用 Python 标准库。
"""
import argparse
from pathlib import Path
import re
from zipfile import ZIP_DEFLATED, ZipFile, ZipInfo


def build(output: Path, preview: Path | None) -> None:
    """从脚本目录打包白名单资源；output 为 ZIP 路径，preview 为可选 HTML 路径。"""
    root = Path(__file__).resolve().parent
    files = [root / "index.html", root / "LICENSE.txt"]
    for directory in ("css", "js", "libs"):
        files.extend(sorted((root / directory).glob("*")))
    output.parent.mkdir(parents=True, exist_ok=True)
    with ZipFile(output, "w", ZIP_DEFLATED, compresslevel=9) as archive:
        for file in files:
            if not file.is_file() or file.suffix not in (".html", ".css", ".js", ".txt"):
                raise ValueError(f"非运行资源禁止进入 ZIP：{file}")
            item = ZipInfo(file.relative_to(root).as_posix(), date_time=(2026, 1, 1, 0, 0, 0))
            item.compress_type = ZIP_DEFLATED
            item.external_attr = 0o100644 << 16
            archive.writestr(item, file.read_bytes(), compresslevel=9)
    if output.stat().st_size > 2_000_000:
        raise ValueError("场景 ZIP 超出 2 MB 加载预算")
    if preview:
        html = (root / "index.html").read_text()
        html = re.sub(r'<link rel="stylesheet" href="([^"]+)">',
                      lambda match: "<style>" + (root / match[1]).read_text() + "</style>", html)
        html = re.sub(r'<script src="([^"]+)"></script>',
                      lambda match: "<script>" + (root / match[1]).read_text().replace("</script", "<\\/script") + "</script>", html)
        html = html.replace("</head>", "<script>" + (root / "js/environment.js").read_text() + "</script></head>")
        preview.parent.mkdir(parents=True, exist_ok=True)
        preview.write_text(html)
    print(f"ZIP：{output}（{output.stat().st_size:,} 字节，{len(files)} 个文件）")


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--output", type=Path, default=Path(__file__).resolve().parent / ".temp" / "power-station.zip")
    parser.add_argument("--preview", type=Path)
    arguments = parser.parse_args()
    build(arguments.output, arguments.preview)
