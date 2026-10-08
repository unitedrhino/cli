# 全部端点速查

| 方法 | 端点 | 说明 | 权限 |
|------|------|------|------|
| POST | `/api/v1/things/hook/` | Hook扩展统一入口 | public |

## 端点详情

### POST `/api/v1/things/hook/`

**说明**: Hook扩展统一入口

**权限**: public

**响应示例**:
```json
{
  "code": 200,
  "msg": "success"
}
```

**调用示例**:
```bash
ur api /api/v1/things/hook/ \
  --body '{}'
```
