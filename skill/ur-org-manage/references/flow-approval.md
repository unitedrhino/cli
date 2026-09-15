# 流程审批中心操作指南（ur-org-manage / flow 命令）

> 手写维护文件，`generate-skills` 不会覆盖本文件（生成器只写 `groups/`、`*-index.md` 与 SKILL.md）。
> 对应后端 TASK-114 流程审批中心（`/api/v1/system/flow/*`）。

## 权限口径（先读）

| 操作集 | 权限 | 说明 |
|-------|------|------|
| 流程发起 / 待办办理 / 我的申请 / 抄送 | 全员 | 普通租户成员可用 |
| 流程定义 / 表单模板 / 流程分类 管理与实例管控（监控/终止/恢复/作废） | 仅管理员 | 后端 `ctxs.IsAdmin` 强校验；普通用户调用返回权限错误码 100012「只允许管理员操作」 |

管理员判定：租户管理员（admin 角色）、租户所有者、平台管理员。普通成员前端「审批管理」菜单不可见，直调 API 被拒属预期，不是故障。

## 专用命令（ur ≥ 支持 flow 的版本）

所有命令在 org-manage 应用下执行（`--app org-manage` 或 `ur-org-manage` 二进制）。管理操作需管理员身份登录。

### 管理侧（仅管理员）

```bash
# 流程分类
ur flow category get-list --size 200          # 分类多时后端缺省 20/页，用 --size 拉全
ur flow category create --name "人事" --code hr --sort 1
ur flow category update --id 12 --name "人事类" --sort 2
ur flow category delete --id 12               # 分类下有流程定义时被拒

# 表单模板
ur flow form get-list --name 请假
ur flow form get-one --id 5                   # 返回 formSchema（JSON 字符串）
ur flow form create --name "请假表单" --schema '{"name":"请假表单","fields":[...]}'
ur flow form create --name "请假表单" --schema @form.json   # @file 传 schema
ur flow form update --id 5 --name "请假表单V2" --schema @form.json
ur flow form delete --id 5                    # 被流程定义引用时被拒

# 流程定义
ur flow def get-list --state published --name 请假
ur flow def get-one --id 9                    # 返回 modelContent（流程模型 JSON 字符串）
ur flow def create --code leave --name "请假审批" \
  --model @model.json --category hr --form-id 5
ur flow def update --id 9 --code leave --name "请假审批V2" --model @model.json
ur flow def deploy --id 9                     # 发布（强校验：必须已绑定表单模板）
ur flow def unpublish --id 9                  # 停用已发布流程
ur flow def delete --id 9                     # 有运行实例时被拒

# 实例管控
ur flow instance get-monitor-list --size 50
ur flow instance terminate --instance-id 77 --comment "违规流程"
ur flow instance resume --instance-id 77
ur flow instance destroy --instance-id 77 --comment "作废"
```

### 审批动线（全员）

```bash
# 发起
ur flow process get-launch-list               # 可发起流程（published，含分类分组）
ur flow process launch --code leave \
  --variable '{"day":3,"reason":"家事"}' --title "张三的请假"

# 办理
ur flow task get-pending-list                 # 我的待办
ur flow task get-detail --task-id 546         # 表单+操作权限+记录聚合视图
ur flow task consent --task-id 546 --comment "同意"
ur flow task reject --task-id 546 --strategy toInitiator --comment "信息不全"
#   strategy: toInitiator / toPreviousNode / toSpecifiedNode(需 --target-node-key) / terminateApproval
ur flow task transfer --task-id 546 --to-user-id 366810892901328 --comment "转办"
ur flow task add-sign --task-id 546 --user-ids "id1,id2"    # 加签
ur flow task remove-sign --task-id 546 --user-ids "id1,id2" # 减签

# 轨迹
ur flow task get-approved-list                # 我的已办
ur flow instance get-my-list                  # 我的申请
ur flow instance get-record --instance-id 452 # 审批记录时间线
ur flow instance revoke --instance-id 452     # 撤销（下一节点未办理前）
ur flow cc get-list                           # 我的抄送
```

## 兜底：通用 api 命令

专用命令未覆盖的端点，用通用调用（路径契约不变）：

```bash
ur api /api/v1/system/flow/task/reclaim --body '{"taskId":"546"}'
```

## JSON 字符串契约要点（易踩坑）

- `modelContent` / `formSchema` / `variable` 后端契约是 **JSON 字符串**（非对象）。
  专用命令已内置处理（`--model`/`--schema`/`--variable` 传内联 JSON 或 `@file`）；
  用通用 `api` 命令时必须自行序列化为字符串。
- `taskId` / `instanceId` / `id` 等后端 `int64,string` 字段，请求体传字符串（如 `{"taskId":"546"}`）。
- 发布强校验：未绑定业务表单的流程定义 `deploy` 会被拒绝，先 `form/create` 再 `def/create --form-id`。
- 流程模型 JSON 规范（节点类型/枚举/结构规则）见主仓
  `docs/中台/任务/进行中/中-26-9-3-TASK-114-流程审批中心/模型JSON规范.md`，AI 生成/修改模型 JSON 前先读该文档。

## 典型任务链路

**管理员建一条可用的审批流**：
`category create` → `form create`（记 formId）→ `def create --form-id`（记 defId）→ `def deploy`。

**用户完成一次审批**：
`process get-launch-list` 选流程 → `process launch`（记 instanceId）→ 审批人 `task get-pending-list`（记 taskId）
→ `task consent` / `task reject` → `instance get-record` 核对时间线。
