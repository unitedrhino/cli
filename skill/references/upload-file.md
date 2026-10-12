# 文件上传（ur system upload-file）

本文件是 `SKILL.md`（ur-api 主入口）的子文档。主文档「生成文件交付」定义 AI 产文件的交付约定；本文件收录命令语法、参数与示例，以及临时桶/公开桶边界。

## 命令语法
```
ur system upload-file [选项]
```

内置 multipart 直传，默认上传到 OSS 临时桶；服务端返回 `fileUri`（形如 `/oss/<临时桶名>/xxx`），可直接用于 markdown 链接交给页面下载（页面经 file-proxy 登录态代理访问）。

## 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --file | -f | 是 | string | 要上传的本地文件路径 |
| --name | -n | 否 | string | 远端文件名（默认取本地文件名） |
| --public | - | 否 | bool | 上传到公开桶（默认临时桶）；需 --business/--use-by |
| --business | - | --public 时必填 | string | 业务标识（公开桶路径分段） |
| --scene | - | 否 | string | 场景标识（公开桶用，默认 file） |
| --use-by | - | 否 | string | 公开桶归属：user（默认）/tenant/platform |
| --json | -j | 否 | bool | 输出完整 JSON 响应 |

## 使用示例

### 示例 1：上传临时桶（AI 生成文件交付的默认姿势）
```bash
ur system upload-file --file /path/to/report.xlsx

# 输出：
# fileUri:  /oss/ithings-temporary/xxx.xlsx
# filePath: xxx.xlsx
# markdown: [report.xlsx](/oss/ithings-temporary/xxx.xlsx)
```

### 示例 2：指定远端文件名
```bash
ur system upload-file --file /tmp/data.csv --name 统计数据.csv
```

### 示例 3：输出完整 JSON
```bash
ur system upload-file --file /path/to/document.pdf --json
```

### 示例 4：上传公开桶（页面可直接访问的永久 URL）
```bash
ur system upload-file --file /path/to/logo.png --public --business view --use-by user
```

## AI 生成文件交付约定

1. 沙箱内生成文件（Excel/报表等）后用本命令上传，取输出中的 `filePath` 行
2. **必须在回复文本中以 markdown 链接形式给出**：`[文件名.xlsx](/oss/temporary/<filePath>)`——页面据此渲染文件卡片并提供预览/下载
3. 边界：只允许 `/oss/` 开头相对路径（临时桶 `/oss/temporary/`、公开桶 `/oss/public/`）；不要使用 `http(s)://` 完整签名地址（有时效且不渲染卡片）；临时桶对象 7 天自动过期；链接需登录态（页面经 file-proxy 访问）

## 对应 API

- `POST /api/v1/system/common/upload-file`（multipart/form-data，字段名 `file`）
