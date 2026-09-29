---
name: ur-doc
description: "Use when 用户上传或引用文档需要读取内容: 解析 PDF/Word/PPT/Excel/HTML/Markdown/邮件/图片为结构地图、Markdown 或无损 Docling JSON;章节提取、表格数字来源解释、excel 公式溯源、扫描件 OCR。triggers: 文档解析, 解析PDF, 解析Excel, 解析Word, 读附件, 用户上传文件, 文件内容, excel公式, 数字怎么来的, 第几章讲的什么, ur doc parse, ur doc"
---

# ur-doc — 文档解析(`ur doc`)

`ur doc` 基于 docling 库把 15 类扩展名(pdf/docx/pptx/xlsx/csv/html/md/adoc/txt/eml/png/jpg/bmp/webp)转为 AI 友好输出。命令只做**通用转换**;单元格、公式、章节等精查交给 jq(沙箱已预装)。

## 命令

```bash
ur doc parse <file|URL|-> [--format outline|md|content-list|json] \
    [--section 前缀] [--sheet 名] [--out 文件] [--ocr] [--ocr-model large] \
    [--pdf-max-file-size-mb 50] [--pdf-max-pages 2000]
ur doc formats
```

## 渐进式工作流(先小后大,避免整份文档进上下文)

1. **先看结构地图**(几百字节):`ur doc parse <f> --format outline`——标题树/表格行列/图片/工作表一览
2. **按需取内容**:
   - 通读或按章节:`--format md`(配 `--section 第四章` 只取命中章节)
   - 精查、公式、单元格坐标:`--format json --out d.json` 落盘后用 jq

## 场景 1:「xx 文档第 4 章讲的啥」

```bash
ur doc parse xx.pdf --format outline              # 1. 从标题树确认第4章标题原文
ur doc parse xx.pdf --format md --section 第四章   # 2. 只取该章(--section 为子串匹配)
```

## 场景 2:「excel 第 3 行第 4 个数字怎么来的」

```bash
ur doc parse 报表.xlsx --format outline                 # 1. sheet 与表格分布
ur doc parse 报表.xlsx --format json --out d.json       # 2. 无损 JSON 落盘
# 3. 列出全部公式单元格(值<TAB>公式):
jq -r '.texts[] | select(.meta."docling__xlsx_formula") | [.text, .meta."docling__xlsx_formula"] | @tsv' d.json
# 4. 反查公式节点(#/texts/3)被哪个表格单元格引用(0-based 行列坐标):
jq '.tables[].data.table_cells[] | select(.ref."$ref" == "#/texts/3") | {text, row: .start_row_offset_idx, col: .start_col_offset_idx}' d.json
# 5. 公式里引用的其他单元格的值,继续查同一份 d.json 或 --format md --sheet 看渲染表格
```

表格单元格 `start_row_offset_idx`/`start_col_offset_idx` 是区域内 0-based 相对坐标,配合表格 prov BBox 换算工作簿绝对坐标(如 D3 = row2/col3)。

## 大文档纪律

- md/json 先 `--out` 落盘,再用 `head`/`sed`/`jq` 局部读,不要整份打进上下文
- `--sheet 名` 过滤工作表;`--section 前缀` 过滤章节(json 格式不受过滤影响,用 jq)
- URL 输入直接下载;PDF 默认限制为 50 MiB/2000 页,可用 `--pdf-max-file-size-mb`、`--pdf-max-pages` 调高;对话中用户文件的 URL(type=file 消息里的 fileUrl)可直接传入
- 两个 PDF 限制参数必须大于 0,不能用 0 关闭保护;结构超限直接报错,单张损坏图片不影响正文

## OCR(扫描件/图片/乱码页)

- `--ocr` 默认走平台模型池(POST `/api/v1/ai/chat/completions` agentID=0,与知识库 OCR 共用租户模型池计费),`--ocr-model xlarge` 可换档
- `--ocr-provider openai` 直连 OpenAI 兼容网关(OPENAI_BASE_URL/OPENAI_API_KEY/DOCLING_OCR_MODEL),无平台环境时用
- 仅 PDF/图片生效;扫描页/乱码页/图片表格/公式密集页自动路由模型,识别失败自动回退纯 Go 结果;`--ocr-max-pages` 控制页数预算

## 沙箱备注

- jq / python3 / rg 已预装;缺工具可 `brew install`
- 解析结果与认证无关;OCR 需要平台凭据(`ur check --json` 先确认)
