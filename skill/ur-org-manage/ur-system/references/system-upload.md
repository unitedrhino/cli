# CLI 命令 — 文件上传

## 命令语法
```
ur system upload-file [选项]
```

## 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --file | -f | 是 | string | 要上传的文件路径 |
| --json | -j | 否 | bool | 输出 JSON 格式 |

## 使用示例

### 示例 1：上传文件
```bash
ur system upload-file --file /path/to/file.jpg
```

### 示例 2：上传文件并输出 JSON
```bash
ur system upload-file --file /path/to/document.pdf --json
```

## 对应 API

- `POST /api/v1/system/common/upload-file`
