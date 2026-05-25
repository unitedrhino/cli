# CLI 命令 — 批量聚合

## 命令语法
```
ur system batch-agg [选项]
```

## 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| --apis | | 是 | string | 批量请求列表 JSON 数组 |
| --json | -j | 否 | bool | 输出 JSON 格式 |

## 使用示例

### 示例 1：批量查询多个接口
```bash
ur system batch-agg \
  --apis '[{"path":"/api/v1/system/user/self/get-one","body":{}},{"path":"/api/v1/system/user/self/tenant/get-list","body":{"page":{"page":1,"size":10}}}]'
```

### 示例 2：批量聚合并输出 JSON
```bash
ur system batch-agg \
  --apis '[{"path":"/api/v1/system/user/self/get-one","body":{}}]' \
  --json
```

## 对应 API

- `POST /api/v1/system/common/api/batch-agg`
