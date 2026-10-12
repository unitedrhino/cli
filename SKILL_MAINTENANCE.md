# Skill 维护指南

本文档说明如何保持 CLI skill 与后端 API 同步。

## 技能流向与上游镜像（可选）

本仓 `skill/` 是 skills 的发布源：`scripts/release.sh` 发版时自动生成 `skill/_meta.json` 版本号并打出
`ur-api-skills-<版本>.zip` 资产；客户端 AI 工具通过 `ur upgrade` / `ur skills install` 拿到发布版本，
不直接感知任何源仓库。`--source` / `UR_SKILLS_SOURCE` 未配置或探测不到时，镜像步骤自动跳过，
swagger 生成与 skills 仓库同步不受影响。

维护方若在内部仓库维护 skills 的编辑副本，可让脚本在更新前先从上游镜像（默认 dry-run 仅列差异，
`--apply` 执行）：

```bash
UR_SKILLS_SOURCE=/path/to/upstream/skills bash scripts/update-skills.sh             # 镜像预览 + 生成 + 同步
UR_SKILLS_SOURCE=/path/to/upstream/skills bash scripts/update-skills.sh --apply    # 确认预览后执行
bash scripts/update-skills.sh --mirror-only [--apply]                              # 只跑镜像步骤
```

注意：镜像以源为准（`rsync --delete`），本仓 `skill/` 多出的文件会被删除；若 `skill/` 有尚未合入
上游的内容（如未合并 PR 的技能），先更新上游再 `--apply`。

## 核心问题

后端 `.api` 文件更新后，skill 中的 API 端点列表不会自动同步。手动维护 500+ 个端点不可持续。

## 解决方案：混合模式

- **手写内容**（不变）：核心概念、工作流、权限说明、注意事项 — 变化慢，人工维护质量高
- **自动生成**（从 swagger 提取）：API 端点列表 — 按 domain 分组，一键刷新

## 文件结构

```
cli/
├── scripts/
│   ├── generate-api-lists.py   # 核心：从 swagger 生成 API 列表
│   └── update-skills.sh        # 包装：一键更新 + 同步到 skills 仓库
├── skill/
│   ├── SKILL.md                # 主 skill（手写）
│   ├── ur-device/
│   │   └── SKILL.md            # 子 skill（手写骨架 + 自动生成端点列表）
│   ├── ur-product/
│   │   └── SKILL.md
│   ├── device-firmware/       # 纯手写设备端技能，不由 Swagger 生成
│   │   ├── SKILL.md
│   │   └── references/
│   ├── ur-ota/                # 手写 OTA 工作流 + 自动生成 API 参考
│   │   ├── SKILL.md
│   │   └── references/api/
│   └── ...
└── SKILL_MAINTENANCE.md        # 本文档
```

## 文件体积规范

客户端 AI 工具（`skill_view`）读取单个技能文件有约 **12000 字节**的读取预算（约 4000 汉字），
超出部分会被截断。此前 `skill/SKILL.md` 就因超长导致截断线以下的内容 AI 读不到（用户投诉），
必须靠拆分解决，不能依赖客户端调大预算。

规则：

- **单文件上限 12000 字节**：`skill/` 下手写文档一律不得超过；`SKILL.md` 是客户端首选加载
  入口，建议控制在 10000 字节以内。
- **超限处理**：把超限内容拆到同目录 `references/` 子文档（文件名小写连字符），主文档只保留
  核心结论、决策表与导航链接；`references/` 下已有相关文档时优先合并，不要另建重复主题。
- **校验入口**：`python3 scripts/check-skill-size.py`。手写文档 >12000 字节失败（退出码 1）、
  >10000 字节警告；已接入 `scripts/update-skills.sh` 的 `[5/5] 检查变更` 步骤与
  `scripts/release.sh` 打包 skills 资产之前，超限会直接中止流程。
- **生成物说明**：`references/api/`、`*-all-endpoints.md`、`swagger-index.md`、`assets/` 及
  非 `.md` 文件等生成/资产类文件当前超限属已知遗留（由 `generate-api-lists.py` /
  `gen-view-components.py` 等生成器改造解决），豁免于本校验，检查时只汇总提示数量。

## 快速开始

### 一键更新（推荐）

后端 `.api` 文件更新后，运行：

```bash
cd /path/to/cli
bash scripts/update-skills.sh
```

脚本自动完成：
1. 读取 `backend/.swagger/{core-api.json, things-api.json}`
2. 按 domain 分组生成 API 端点列表
3. 插入到各 skill 文件的 `<!-- API_LIST:domain -->` 和 `<!-- END_API_LIST -->` 标记之间
4. 同步到 `unitedrhino/skills` 仓库（如果找到）

### 只生成不同步

```bash
python3 scripts/generate-api-lists.py
```

### 手动同步到 skills 仓库

```bash
# 复制到 skills 仓库
cp -r skill/* /path/to/skills/
```

## Domain 分组规则

| Domain | 路径前缀 | Skill 目录 |
|--------|---------|-----------|
| ur-device | `/api/v1/things/device/`（排除 msg/ 和 interact/） | `skill/ur-device/` |
| ur-device-analytics | `/api/v1/things/device/msg/` | `skill/ur-device-analytics/` |
| ur-device-debug | `/api/v1/things/device/interact/` | `skill/ur-device-debug/` |
| ur-product | `/api/v1/things/product/`、`/api/v1/things/device/ota/` | `skill/ur-product/` |
| ur-project | `/api/v1/things/project/`、`/api/v1/things/area/`、`/api/v1/things/group/`、`/api/v1/things/data/` | `skill/ur-project/` |
| ur-user | `/api/v1/system/user/`、`/api/v1/system/role/`、`/api/v1/system/dept/`、`/api/v1/system/dict/`、`/api/v1/system/notify/`、`/api/v1/system/log/` | `skill/ur-user/` |
| ur-tenant | `/api/v1/system/tenant/` | `skill/ur-tenant/` |
| ur-system | `/api/v1/system/`（排除 user/tenant/ 等） | `skill/ur-system/` |
| ur-ai | `/api/v1/ai/`、`/api/v1/things/alarm/`、`/api/v1/things/scene/` | `skill/ur-ai/` |
| scene-linkage | `/api/v1/things/scene/` | `skill/scene-linkage/` |
| device-firmware | 无 API 前缀，纯手写设备端构建与接入知识 | `skill/device-firmware/` |
| ur-ota | `/api/v1/things/ota/`，入口工作流手写、接口参考自动生成 | `skill/ur-ota/` |

如需调整分组规则，修改 `scripts/generate-api-lists.py` 中的 `DOMAIN_PREFIXES`。
`device-firmware` 不参与 Swagger 端点生成，`ur-ota/SKILL.md` 也不会由生成器创建；修改后需同时验证
`scripts/package-skill.sh` 的统一技能包、npm 包和独立发布包分发结果。

## Skill 文件标记

在 API 参考部分插入以下标记，脚本会自动替换标记之间的内容：

```markdown
## API 参考

<!-- API_LIST:ur-device -->

（API 端点列表由脚本从 swagger 自动生成）

<!-- END_API_LIST -->
```

**注意**：标记之间的内容会被脚本完全覆盖，不要在此区域手动编辑。

## 完整更新流程

```bash
# 1. 后端更新 .api 文件后，重新生成 swagger
cd backend/core/service/apisvr
bash build.sh

cd backend/things/service/apisvr
bash build.sh

# 2. 更新 skill
cd .gits/cli
bash scripts/update-skills.sh

# 3. 检查变更
git diff --stat skill/

# 4. 提交 cli 仓库
git add skill/ scripts/
git commit -m "chore(skill): 同步后端 API 变更

- 更新 API 端点列表（共 XX 个接口）
- 新增/修改/删除的接口见 diff"
git push

# 5. 提交 skills 仓库
cd ../skills
git add -A
git commit -m "chore(skill): 同步 API 端点列表"
git push origin main && git push gitee main
```

## 排除的接口

以下接口不会被纳入 skill（业务不常用或内部使用）：
- `/api/v1/system/init/*` — 系统初始化
- `/api/v1/system/checkIn/*` — 签到
- `/api/v1/system/ops/*` — 运维工单
- `/api/v1/system/mall/*` — 商城授权

如需纳入，在 `generate-api-lists.py` 的 `DOMAIN_PREFIXES` 中添加对应前缀。

## 常见问题

### Q: 为什么有些接口不在 skill 中？
检查 `DOMAIN_PREFIXES` 映射。如果路径前缀不在任何 domain 中，该接口会被忽略。

### Q: 接口分组错了怎么办？
`DOMAIN_PREFIXES` 按**最长前缀优先**匹配。调整前缀顺序或添加更精确的前缀即可。

### Q: 如何验证更新后的 skill？
```bash
# 查看某个 skill 的 API 列表
sed -n '/API_LIST:ur-device/,/END_API_LIST/p' skill/ur-device/SILL.md

# 统计端点数量
python3 scripts/generate-api-lists.py | grep "个端点"
```
