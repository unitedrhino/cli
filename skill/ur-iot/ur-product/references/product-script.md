# CLI 命令 — 协议脚本（script）

协议脚本相关操作：验证脚本语法、生成脚本模板。

协议脚本使用 yaegi Go 解释器执行，支持以下包：
- `log` / `context` / `strings` / `json` / `gjson` / `utils` / `deviceMsg` / `dm` / `schema`

---

## script validate — 校验协议脚本

### 命令语法

```
ur things script validate <文件路径>
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| file |  | 是 | string | 脚本文件路径，`-`表示从标准输入读取 |

### 校验内容

1. 基础 Go 语法解析
2. Handle 函数是否存在
3. import 包名是否在 yaegi 白名单中
4. 提示可用的 Handle 函数签名

### Handle 函数签名参考

| 脚本类型 | 签名 |
|----------|------|
| up-before | `func Handle(ctx context.Context, req *deviceMsg.PublishMsg) *deviceMsg.PublishMsg` |
| up-after | `func Handle(ctx context.Context, req *deviceMsg.PublishMsg, resp *deviceMsg.PublishMsg)` |
| down-before | `func Handle(ctx context.Context, req *deviceMsg.PublishMsg) *deviceMsg.PublishMsg` |
| down-after | `func Handle(ctx context.Context, req *deviceMsg.PublishMsg)` |

### 使用示例

#### 示例1：校验脚本文件
```bash
ur things script validate script.go
```

#### 示例2：从标准输入校验
```bash
cat script.go | ur things script validate -
```

### 注意事项

- Before 脚本返回 nil 会丢弃消息
- `dm`/`schema` 包函数需要真实服务运行，本地不可用
- 完整测试请使用：`go test -v -run TestScript ./...`

---

## script template — 生成脚本模板

### 命令语法

```
ur things script template <脚本类型>
```

### 参数说明

| 参数 | 简写 | 必填 | 类型 | 说明 |
|------|------|------|------|------|
| script-type |  | 是 | string | 脚本类型：up-before/up-after/down-before/down-after |

### 使用示例

#### 示例1：生成上行前处理模板
```bash
ur things script template up-before
```

#### 示例2：生成上行后处理模板
```bash
ur things script template up-after
```

#### 示例3：生成下行前处理模板
```bash
ur things script template down-before
```

#### 示例4：生成下行后处理模板
```bash
ur things script template down-after
```

### 脚本类型说明

| 类型 | 触发方向 | 触发时机 | 返回值 | 用途 |
|------|----------|----------|--------|------|
| up-before | 上行（设备→平台） | 处理前 | `*deviceMsg.PublishMsg` | 拦截/修改设备上报消息 |
| up-after | 上行（设备→平台） | 处理后 | 无 | 联动、记录 |
| down-before | 下行（平台→设备） | 处理前 | `*deviceMsg.PublishMsg` | 拦截/修改下发指令 |
| down-after | 下行（平台→设备） | 处理后 | 无 | 记录、联动 |
