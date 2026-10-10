---
name: ur-doc
description: "使用联犀 CLI 解析文档、CAD 图纸和 GLB 三维模型，提取章节、图框、文本、模型属性与来源；适用于读附件、图纸分析、模型分析、知识库入库和 ur doc 命令。"
---

# ur-doc — 文档解析(`ur doc`)

`ur doc` 基于 docling 将 PDF、Office、图片、邮件等文档，以及 DWG/DXF/DXFB 图纸、GLB 2.0 三维模型转为结构地图、Markdown、content-list 或 Docling JSON。用 `ur doc formats` 查看当前 CLI 支持的扩展名。

CAD 图纸自动按识别到的图框组织图名、渲染图和文本。GLB 提取场景、节点关系、网格、渲染材质及已有属性；CLI 只做元数据转换，不渲染三维几何。

## 命令

```bash
ur doc parse <file|URL|-> [--format outline|md|content-list|json] \
    [--section 前缀] [--sheet 名] [--out 文件] [--ocr] [--ocr-model large]
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

## 场景 3:「这份施工图里都有什么」(CAD)

```bash
ur doc parse 施工图.dwg --format outline   # 1. 图框清单:每张图的图名/页号(渐进式披露)
ur doc parse 施工图.dwg --format md        # 2. 全文:每图框一节(图名标题+图框内文本),含标注数值/图例表
ur doc parse 施工图.dwg --format json --out d.json   # 3. 无损 JSON:文本带页号与 BBox,图片项内嵌渲染 PNG
```

- DWG 自动按图框拆分,图名从标题栏提取(如"RD-31-十三层弱电平面图");设计说明/图例表等无框内容自动兜底切分
- 中文标注、φ/°/± 符号、尺寸标注数值直接可读,适合图纸知识库入库与多模态问答(文本+渲染图一起给模型)

## 场景 4：读取 GLB 三维模型

先查看结构，再只读取相关节点的属性，保留对象索引与来源：

```bash
ur doc parse 模型.glb --format outline
ur doc parse 模型.glb --format md --section 泵
ur doc parse 模型.glb --format content-list --out model-items.json
ur doc parse 模型.glb --format json --out model.json
# 查看对象来源；pointer 如 #/nodes/1，可回到原模型 JSON 定位。
jq '.texts[] | select(.meta["glb:source"]) | {text, source: .meta["glb:source"]}' model.json
```

在 WorkBuddy / CodeBuddy 中，指定可访问的文件后，可请求：

> 用联犀 CLI 分析工作目录里的“模型.glb”，整理场景、节点关系、设备名称和已有属性，注明节点索引及来源；模型没有写明的参数单独列出。

- GLB 支持来自 docling v1.5.0；使用前核对 `ur doc formats` 包含 `glb`。老版本使用 `ur upgrade` 更新 CLI 与客户端 Skills，重新加载会话后再解析。
- `extras` / `extensions` 中已有属性进入正文和 JSON；设备编号等大整数保持原始精度。content-list 保留节点章节与“模型来源”正文，JSON 另有 `glb:source` 元数据。
- 渲染材质名称不等于工程材料。没有写入文件的尺寸、数量、BOM 和设备参数不能凭外观或节点名称推测。
- 不请求外部纹理或缓冲，不解压 Draco 几何或解码 KTX2 纹理，也不把 `--ocr` 用作三维视觉分析。模型结构与属性可接入知识库；三维预览由知识库前端提供。
- 文件上限 128 MiB、JSON 块上限 16 MiB、节点上限 10,000；损坏容器或非法节点关系直接报错。

## 大文档纪律

- md/json 先 `--out` 落盘,再用 `head`/`sed`/`jq` 局部读,不要整份打进上下文
- `--sheet 名` 过滤工作表;`--section 前缀` 过滤章节(json 格式不受过滤影响,用 jq)
- URL 输入直接下载(≤200MB);对话中用户文件的 URL(type=file 消息里的 fileUrl)可直接传入

## OCR(扫描件/图片/乱码页)

- `--ocr` 默认走平台模型池(POST `/api/v1/ai/chat/completions` agentID=0,与知识库 OCR 共用租户模型池计费),`--ocr-model xlarge` 可换档
- `--ocr-provider openai` 直连 OpenAI 兼容网关(OPENAI_BASE_URL/OPENAI_API_KEY/DOCLING_OCR_MODEL),无平台环境时用
- 仅 PDF/图片生效;扫描页/乱码页/图片表格/公式密集页自动路由模型,识别失败自动回退纯 Go 结果;`--ocr-max-pages` 控制页数预算

## 沙箱备注

- jq / python3 / rg 已预装;缺工具可 `brew install`
- 解析结果与认证无关;OCR 需要平台凭据(`ur check --json` 先确认)
