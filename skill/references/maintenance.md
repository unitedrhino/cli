# 新增接口维护规范

## 新增接口检查清单

1. **标注权限**：在 `.api` 文件中添加 `// x-auth-type: admin|all|platform`
2. **更新 Skill**：
   - 新增域 → 在 `backend/cli/ur/internal/config/app.go` 的 `Features()` 中添加功能模块定义
   - 新增接口 → 确认后端 `.api` 文件中已标注 `// x-auth-type: admin|all|platform`
3. **重新生成**：
   ```bash
   cd backend/cli/ur
   go run ./cmd/ur-iot generate-skills
   # 或批量生成
   for app in ur-platform-manage ur-iot ur-org-manage ur-org-energy ur-console; do
     go run ./cmd/$app generate-skills
   done
   ```

## 角色权限标注指南

| 标注值 | 适用场景 |
|--------|---------|
| `platform` | 跨企业操作、企业 CRUD、全局配置 |
| `admin` | 本企业内 CRUD、用户管理、角色分配 |
| `all` | 个人信息、设备分享、查看权限范围内数据 |
| `mixed` | 不同角色有不同权限（需在 `roleViews` 中详细说明） |
