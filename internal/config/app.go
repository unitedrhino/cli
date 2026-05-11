package config

import "fmt"

// CLIApp 前端应用标识，每个对应一个 CLI 二进制
type CLIApp string

const (
	AppPlatformManage CLIApp = "platform-manage" // 平台管理
	AppIoT            CLIApp = "iot"             // 物联网
	AppOrgManage      CLIApp = "org-manage"      // 组织管理
	AppOrgEnergy      CLIApp = "org-energy"      // 能源管理
	AppConsole        CLIApp = "console"         // 控制台
)

// Feature 功能模块定义（从前端页面提取，与代码放在一起维护）
type Feature struct {
	Name        string    // 功能名称
	Description string    // 功能说明
	Authority   []string  // 可见角色，空表示所有角色可见；如 ["platform"] 表示仅平台管理员
	APIs        []string  // 涉及的 API 路径
	SubFeatures []Feature // 子功能
}

// AllCLIApps 返回所有可用的 CLI 应用列表
func AllCLIApps() []CLIApp {
	return []CLIApp{AppPlatformManage, AppIoT, AppOrgManage, AppOrgEnergy, AppConsole}
}

// AppID 返回应用 ID
func (a CLIApp) AppID() string {
	switch a {
	case AppPlatformManage:
		return "100"
	case AppIoT:
		return "200"
	case AppOrgManage:
		return "300"
	case AppOrgEnergy:
		return "1000"
	case AppConsole:
		return "77"
	default:
		return ""
	}
}

// DefaultTenantCode 返回默认租户代码，空字符串表示需要用户输入
func (a CLIApp) DefaultTenantCode() string {
	switch a {
	case AppPlatformManage, AppIoT, AppConsole:
		return "platform"
	default:
		return "" // 组织类应用由用户输入
	}
}

// AllowedAuthTypes 返回该应用可调用的 API 权限类型
func (a CLIApp) AllowedAuthTypes() []string {
	switch a {
	case AppPlatformManage, AppIoT:
		return []string{"platform", "admin", "all", "public"}
	case AppOrgManage, AppOrgEnergy:
		return []string{"admin", "all", "public"}
	case AppConsole:
		return []string{"all", "public"}
	default:
		return []string{"all", "public"}
	}
}

// DisplayName 返回中文显示名称
func (a CLIApp) DisplayName() string {
	switch a {
	case AppPlatformManage:
		return "平台管理"
	case AppIoT:
		return "物联网"
	case AppOrgManage:
		return "组织管理"
	case AppOrgEnergy:
		return "能源管理"
	case AppConsole:
		return "控制台"
	default:
		return string(a)
	}
}

// BinaryName 返回 CLI 二进制名称
func (a CLIApp) BinaryName() string {
	return "ur-" + string(a)
}

// ParseCLIApp 解析应用标识
func ParseCLIApp(s string) (CLIApp, error) {
	switch CLIApp(s) {
	case AppPlatformManage, AppIoT, AppOrgManage, AppOrgEnergy, AppConsole:
		return CLIApp(s), nil
	default:
		return "", fmt.Errorf("unknown CLI app: %q, valid: platform-manage, iot, org-manage, org-energy, console", s)
	}
}

// Features 返回该应用的功能指引（从前端页面提取，单一事实来源）
func (a CLIApp) Features() []Feature {
	switch a {
	case AppPlatformManage:
		return platformManageFeatures()
	case AppIoT:
		return iotFeatures()
	case AppOrgManage:
		return orgManageFeatures()
	case AppOrgEnergy:
		return orgEnergyFeatures()
	case AppConsole:
		return consoleFeatures()
	default:
		return nil
	}
}

// platformManageFeatures 平台管理功能（从前端路由 platform-manage.ts 提取）
func platformManageFeatures() []Feature {
	return []Feature{
		{
			Name:        "企业管理",
			Description: "租户（企业）的全生命周期管理",
			APIs:        []string{"/api/v1/system/tenant/info/create", "/api/v1/system/tenant/info/update", "/api/v1/system/tenant/info/delete", "/api/v1/system/tenant/info/get-one", "/api/v1/system/tenant/info/get-list"},
			SubFeatures: []Feature{
				{Name: "企业列表", Description: "查看和管理所有租户", APIs: []string{"/api/v1/system/tenant/info/get-list"}},
				{Name: "企业详情", Description: "查看租户详情、绑定应用、配置", APIs: []string{"/api/v1/system/tenant/info/get-one", "/api/v1/system/tenant/app/get-list", "/api/v1/system/tenant/app/create"}},
			},
		},
		{
			Name:        "用户管理",
			Description: "平台级用户管理（跨租户）",
			APIs:        []string{"/api/v1/system/user/info/get-list", "/api/v1/system/user/info/get-one", "/api/v1/system/user/info/create", "/api/v1/system/user/info/update", "/api/v1/system/user/info/delete"},
			SubFeatures: []Feature{
				{Name: "用户列表", Description: "查看所有租户下的用户", APIs: []string{"/api/v1/system/user/info/get-list"}},
				{Name: "用户详情", Description: "查看用户详情和租户绑定", APIs: []string{"/api/v1/system/user/info/get-one"}},
			},
		},
		{
			Name:        "应用管理",
			Description: "系统应用的创建和配置",
			APIs:        []string{"/api/v1/system/app/info/create", "/api/v1/system/app/info/update", "/api/v1/system/app/info/get-one", "/api/v1/system/app/info/get-list", "/api/v1/system/app/menu/create", "/api/v1/system/app/menu/get-list"},
			SubFeatures: []Feature{
				{Name: "应用列表", Description: "查看和管理系统应用", APIs: []string{"/api/v1/system/app/info/get-list"}},
				{Name: "应用菜单", Description: "管理应用的菜单树", APIs: []string{"/api/v1/system/app/menu/get-list", "/api/v1/system/app/menu/create"}},
			},
		},
		{
			Name:        "授权管理",
			Description: "授权码、套餐和续期管理",
			APIs:        []string{"/api/v1/system/license/package/get-list", "/api/v1/system/license/info/get-list", "/api/v1/system/license/record/get-list"},
			SubFeatures: []Feature{
				{Name: "套餐管理", Description: "管理授权套餐", APIs: []string{"/api/v1/system/license/package/get-list", "/api/v1/system/license/package/create"}},
				{Name: "授权码管理", Description: "生成和管理授权码", APIs: []string{"/api/v1/system/license/info/get-list", "/api/v1/system/license/info/create"}},
				{Name: "审计记录", Description: "查看授权操作记录", APIs: []string{"/api/v1/system/license/record/get-list"}},
			},
		},
		{
			Name:        "协议管理",
			Description: "用户协议、隐私政策等协议库管理",
			APIs:        []string{"/api/v1/system/agreement/create", "/api/v1/system/agreement/update", "/api/v1/system/agreement/get-list", "/api/v1/system/agreement/get-one"},
		},
		{
			Name:        "系统管理",
			Description: "系统设置、权限和日志",
			SubFeatures: []Feature{
				{Name: "系统设置", Description: "全局系统配置", APIs: []string{"/api/v1/system/common/config/get-one", "/api/v1/system/common/config/save"}},
				{Name: "接口管理", Description: "API 接口注册和管理", APIs: []string{"/api/v1/system/api/info/get-list", "/api/v1/system/api/info/create"}},
				{Name: "资源总览", Description: "权限资源树管理", APIs: []string{"/api/v1/system/resource/tree/get-list"}},
				{Name: "角色管理", Description: "平台角色 CRUD", APIs: []string{"/api/v1/system/role/info/get-list", "/api/v1/system/role/info/create"}},
				{Name: "登录日志", Description: "查看登录日志", APIs: []string{"/api/v1/system/log/login/get-list"}},
				{Name: "操作日志", Description: "查看操作日志", APIs: []string{"/api/v1/system/log/operation/get-list"}},
			},
		},
		{
			Name:        "系统开发",
			Description: "Hook 扩展、字典、数据源、任务管理",
			SubFeatures: []Feature{
				{Name: "Hook 扩展管理", Description: "管理 Hook 扩展点", APIs: []string{"/api/v1/system/hook/info/get-list", "/api/v1/system/hook/info/create"}},
				{Name: "字典管理", Description: "数据字典 CRUD", APIs: []string{"/api/v1/system/dict/info/get-list", "/api/v1/system/dict/info/create", "/api/v1/system/dict/data/get-list"}},
				{Name: "数据源管理", Description: "外部数据源配置", APIs: []string{"/api/v1/system/datasource/get-list", "/api/v1/system/datasource/create"}},
				{Name: "查询配置管理", Description: "数据查询配置", APIs: []string{"/api/v1/system/query/config/get-list", "/api/v1/system/query/config/create"}},
				{Name: "任务管理", Description: "定时任务管理", APIs: []string{"/api/v1/system/task/group/get-list", "/api/v1/system/task/info/get-list"}},
			},
		},
		{
			Name:        "AI 中台",
			Description: "AI 模型、智能体、知识库、技能包管理",
			SubFeatures: []Feature{
				{Name: "模型配置", Description: "LLM/TTS/ASR 模型配置", APIs: []string{"/api/v1/ai/config/get-list", "/api/v1/ai/config/create"}},
				{Name: "智能体分组", Description: "Agent 分组管理", APIs: []string{"/api/v1/ai/agent/group/get-list", "/api/v1/ai/agent/group/create"}},
				{Name: "智能体管理", Description: "Agent CRUD", APIs: []string{"/api/v1/ai/agent/get-list", "/api/v1/ai/agent/create"}},
				{Name: "MCP 服务", Description: "MCP 服务管理", APIs: []string{"/api/v1/ai/mcp/service/get-list", "/api/v1/ai/mcp/service/create"}},
				{Name: "知识库", Description: "知识库和文档管理", APIs: []string{"/api/v1/ai/knowledge/base/get-list", "/api/v1/ai/knowledge/base/create"}},
				{Name: "技能包", Description: "AI 技能包管理", APIs: []string{"/api/v1/ai/skill/get-list", "/api/v1/ai/skill/create"}},
				{Name: "记忆调试", Description: "AI 记忆系统调试", APIs: []string{"/api/v1/ai/memory/debug"}},
			},
		},
		{
			Name:        "通知管理",
			Description: "通知配置、模板、消息记录",
			SubFeatures: []Feature{
				{Name: "手动通知", Description: "手动发送通知", APIs: []string{"/api/v1/system/notify/config/get-list", "/api/v1/system/notify/send"}},
				{Name: "通知模板", Description: "通知模板管理", APIs: []string{"/api/v1/system/notify/template/get-list", "/api/v1/system/notify/template/create"}},
				{Name: "事件通知配置", Description: "事件触发通知配置", APIs: []string{"/api/v1/system/notify/config/get-list", "/api/v1/system/notify/config/create"}},
				{Name: "消息记录", Description: "消息发送记录", APIs: []string{"/api/v1/system/notify/record/get-list"}},
				{Name: "资讯管理", Description: "资讯内容管理", APIs: []string{"/api/v1/system/notify/news/get-list", "/api/v1/system/notify/news/create"}},
			},
		},
	}
}

// iotFeatures 物联网功能（从前端路由 iot.ts 提取）
func iotFeatures() []Feature {
	return []Feature{
		{
			Name:        "信息面板",
			Description: "IoT Dashboard，展示设备和产品统计",
			APIs:        []string{"/api/v1/things/device/info/get-list", "/api/v1/things/product/info/get-list"},
		},
		{
			Name:        "设备地图",
			Description: "全局设备地图展示",
			APIs:        []string{"/api/v1/things/device/info/get-list"},
		},
		{
			Name:        "设备管理",
			Description: "产品和设备的全生命周期管理",
			SubFeatures: []Feature{
				{
					Name:        "产品管理",
					Description: "产品 CRUD、导入导出",
					APIs:        []string{"/api/v1/things/product/info/create", "/api/v1/things/product/info/update", "/api/v1/things/product/info/delete", "/api/v1/things/product/info/get-one", "/api/v1/things/product/info/get-list"},
					SubFeatures: []Feature{
						{Name: "产品列表", Description: "查看和管理产品", APIs: []string{"/api/v1/things/product/info/get-list"}},
						{Name: "产品详情", Description: "查看产品详情和物模型", APIs: []string{"/api/v1/things/product/info/get-one", "/api/v1/things/product/schema/get-list"}},
						{Name: "添加产品", Description: "创建新产品", APIs: []string{"/api/v1/things/product/info/create"}},
					},
				},
				{
					Name:        "设备管理",
					Description: "设备 CRUD、状态查看、属性控制",
					APIs:        []string{"/api/v1/things/device/info/create", "/api/v1/things/device/info/update", "/api/v1/things/device/info/delete", "/api/v1/things/device/info/get-one", "/api/v1/things/device/info/get-list"},
					SubFeatures: []Feature{
						{Name: "设备列表", Description: "查看和管理设备", APIs: []string{"/api/v1/things/device/info/get-list"}},
						{Name: "设备详情", Description: "查看设备详情、属性、事件", APIs: []string{"/api/v1/things/device/info/get-one", "/api/v1/things/device/msg/property-latest/get-list"}},
						{Name: "属性控制", Description: "向设备发送属性控制指令", APIs: []string{"/api/v1/things/device/interact/property-control-send"}},
					},
				},
				{
					Name:        "通用物模型",
					Description: "平台级通用物模型管理（平台专属）",
					Authority:   []string{"platform"},
					APIs:        []string{"/api/v1/things/schema/create", "/api/v1/things/schema/update", "/api/v1/things/schema/get-list"},
				},
				{
					Name:        "产品品类",
					Description: "产品品类管理（平台专属）",
					Authority:   []string{"platform"},
					APIs:        []string{"/api/v1/things/category/create", "/api/v1/things/category/get-list"},
				},
			},
		},
		{
			Name:        "项目管理",
			Description: "项目 CRUD、场景联动",
			APIs:        []string{"/api/v1/things/project/crud/create", "/api/v1/things/project/crud/update", "/api/v1/things/project/crud/delete", "/api/v1/things/project/crud/get-one", "/api/v1/things/project/crud/get-list"},
			SubFeatures: []Feature{
				{Name: "项目列表", Description: "查看和管理项目", APIs: []string{"/api/v1/things/project/crud/get-list"}},
				{Name: "项目详情", Description: "查看项目详情和关联设备", APIs: []string{"/api/v1/things/project/crud/get-one"}},
				{Name: "场景编辑", Description: "编辑项目场景联动规则", APIs: []string{"/api/v1/things/scene/info/create", "/api/v1/things/scene/info/update", "/api/v1/things/scene/info/get-list"}},
			},
		},
		{
			Name:        "区域管理",
			Description: "区域 CRUD 和设备分配",
			APIs:        []string{"/api/v1/things/area/create", "/api/v1/things/area/update", "/api/v1/things/area/delete", "/api/v1/things/area/get-list"},
		},
		{
			Name:        "OTA 升级",
			Description: "固件升级包管理和批量升级",
			SubFeatures: []Feature{
				{Name: "升级包列表", Description: "管理固件升级包", APIs: []string{"/api/v1/things/ota/firmware/create", "/api/v1/things/ota/firmware/get-list"}},
				{Name: "模块列表", Description: "管理 OTA 模块", APIs: []string{"/api/v1/things/ota/module/get-list"}},
				{Name: "批量升级", Description: "创建批量升级任务", APIs: []string{"/api/v1/things/ota/job/create", "/api/v1/things/ota/job/get-list"}},
			},
		},
		{
			Name:        "数据流转",
			Description: "协议网关和协议脚本管理",
			SubFeatures: []Feature{
				{Name: "协议网关", Description: "协议网关管理（平台专属）", Authority: []string{"platform"}, APIs: []string{"/api/v1/things/protocol/gateway/get-list", "/api/v1/things/protocol/gateway/create"}},
				{Name: "协议脚本", Description: "协议脚本管理", APIs: []string{"/api/v1/things/protocol/script/get-list", "/api/v1/things/protocol/script/create"}},
			},
		},
	}
}

// orgManageFeatures 组织管理功能（从前端路由 org-manage/app.ts 提取）
func orgManageFeatures() []Feature {
	return []Feature{
		{
			Name:        "企业信息",
			Description: "当前租户的企业信息管理",
			APIs:        []string{"/api/v1/system/tenant/info/get-one", "/api/v1/system/tenant/info/update"},
		},
		{
			Name:        "消费记录",
			Description: "授权消费记录查看",
			APIs:        []string{"/api/v1/system/license/record/get-list"},
		},
		{
			Name:        "AI 管理",
			Description: "智能体、数字分身、会话管理",
			SubFeatures: []Feature{
				{Name: "智能体管理", Description: "租户级 Agent 管理", APIs: []string{"/api/v1/ai/agent/get-list", "/api/v1/ai/agent/create", "/api/v1/ai/agent/update"}},
				{Name: "数字分身", Description: "AI 数字分身管理", APIs: []string{"/api/v1/ai/clone/get-list", "/api/v1/ai/clone/create"}},
				{Name: "会话管理", Description: "AI 会话记录查看", APIs: []string{"/api/v1/ai/session/get-list"}},
			},
		},
		{
			Name:        "系统管理",
			Description: "租户内系统管理",
			SubFeatures: []Feature{
				{Name: "用户管理", Description: "租户内用户 CRUD", APIs: []string{"/api/v1/system/user/info/get-list", "/api/v1/system/user/info/create", "/api/v1/system/user/info/update"}},
				{Name: "角色管理", Description: "租户内角色 CRUD", APIs: []string{"/api/v1/system/role/info/get-list", "/api/v1/system/role/info/create"}},
				{Name: "菜单管理", Description: "租户菜单管理", APIs: []string{"/api/v1/system/app/menu/get-list"}},
				{Name: "智能体配置", Description: "租户级 Agent 配置", APIs: []string{"/api/v1/system/tenant/agent/get-list"}},
				{Name: "续期管理", Description: "授权续期", APIs: []string{"/api/v1/system/tenant/renewal"}},
				{Name: "企业设置", Description: "租户级设置", APIs: []string{"/api/v1/system/tenant/info/update"}},
			},
		},
		{
			Name:        "通知管理",
			Description: "消息配置和模板管理",
			SubFeatures: []Feature{
				{Name: "消息配置", Description: "通知配置管理", APIs: []string{"/api/v1/system/notify/config/get-list"}},
				{Name: "模版配置", Description: "通知模板管理", APIs: []string{"/api/v1/system/notify/template/get-list"}},
				{Name: "站内信", Description: "站内消息", APIs: []string{"/api/v1/system/notify/record/get-list"}},
			},
		},
		{
			Name:        "日志管理",
			Description: "操作日志查看",
			SubFeatures: []Feature{
				{Name: "操作日志", Description: "查看操作日志", APIs: []string{"/api/v1/system/log/operation/get-list"}},
			},
		},
	}
}

// orgEnergyFeatures 能源管理功能（从前端路由 org-energy/ 提取）
func orgEnergyFeatures() []Feature {
	return []Feature{
		{
			Name:        "大屏",
			Description: "能源数据大屏展示",
			APIs:        []string{"/api/v1/things/device/msg/property-latest/get-list"},
		},
		{
			Name:        "设备空间",
			Description: "设备控制台、分组和区域管理",
			SubFeatures: []Feature{
				{Name: "控制台", Description: "设备控制台", APIs: []string{"/api/v1/things/device/info/get-list", "/api/v1/things/device/interact/property-control-send"}},
				{Name: "设备分组", Description: "设备分组管理", APIs: []string{"/api/v1/things/group/create", "/api/v1/things/group/get-list"}},
				{Name: "区域管理", Description: "区域 CRUD", APIs: []string{"/api/v1/things/area/create", "/api/v1/things/area/get-list"}},
			},
		},
		{
			Name:        "能耗分析",
			Description: "多维度能耗分析报表",
			SubFeatures: []Feature{
				{Name: "用能概况", Description: "总览用能数据", APIs: []string{"/api/v1/things/device/msg/property-latest/get-list"}},
				{Name: "同比环比分析", Description: "用能同比环比对比", APIs: []string{"/api/v1/things/device/msg/property-latest/get-list"}},
				{Name: "能耗趋势", Description: "用能趋势图表", APIs: []string{"/api/v1/things/device/msg/property-history/get-list"}},
				{Name: "能耗排名", Description: "按区域/设备排名", APIs: []string{"/api/v1/things/device/msg/property-latest/get-list"}},
				{Name: "损耗分析", Description: "能源损耗分析", APIs: []string{"/api/v1/things/device/msg/property-history/get-list"}},
				{Name: "用能报表", Description: "用能数据报表导出", APIs: []string{"/api/v1/things/device/msg/property-history/get-list"}},
				{Name: "复费率报表", Description: "分时电价复费率统计", APIs: []string{"/api/v1/things/device/msg/property-history/get-list"}},
				{Name: "集抄明细报表", Description: "电表集抄数据明细", APIs: []string{"/api/v1/things/device/msg/property-history/get-list"}},
				{Name: "一次图", Description: "电力一次接线图展示", APIs: []string{"/api/v1/things/device/msg/property-latest/get-list"}},
			},
		},
		{
			Name:        "电力集抄",
			Description: "电力数据采集和报表",
			SubFeatures: []Feature{
				{Name: "数据监控", Description: "实时电力数据监控", APIs: []string{"/api/v1/things/device/msg/property-latest/get-list"}},
				{Name: "电力数据", Description: "电力历史数据查询", APIs: []string{"/api/v1/things/device/msg/property-history/get-list"}},
				{Name: "极值报表", Description: "电力极值统计", APIs: []string{"/api/v1/things/device/msg/property-history/get-list"}},
				{Name: "电力报表", Description: "电力数据报表", APIs: []string{"/api/v1/things/device/msg/property-history/get-list"}},
				{Name: "配电图", Description: "电力一次接线图展示", APIs: []string{"/api/v1/things/device/msg/property-latest/get-list"}},
			},
		},
		{
			Name:        "场景联动",
			Description: "自动化和一键场景",
			SubFeatures: []Feature{
				{Name: "自动化", Description: "自动化场景管理", APIs: []string{"/api/v1/things/scene/info/get-list", "/api/v1/things/scene/info/create", "/api/v1/things/scene/info/manually-trigger"}},
				{Name: "一键场景", Description: "手动触发场景", APIs: []string{"/api/v1/things/scene/info/get-list", "/api/v1/things/scene/info/manually-trigger"}},
			},
		},
		{
			Name:        "预付费管理",
			Description: "房间、租客、充值和消费管理",
			SubFeatures: []Feature{
				{Name: "房间管理", Description: "房间 CRUD", APIs: []string{"/api/v1/things/room/get-list", "/api/v1/things/room/create"}},
				{Name: "租客管理", Description: "租客 CRUD", APIs: []string{"/api/v1/things/tenant/get-list", "/api/v1/things/tenant/create"}},
				{Name: "充值明细", Description: "充值记录查询", APIs: []string{"/api/v1/things/prepaid/recharge/get-list"}},
				{Name: "消费明细", Description: "消费记录查询", APIs: []string{"/api/v1/things/prepaid/consumption/get-list"}},
				{Name: "系统设置", Description: "预付费系统参数配置", APIs: []string{"/api/v1/things/prepaid/config/get-one", "/api/v1/things/prepaid/config/save"}},
			},
		},
		{
			Name:        "系统管理",
			Description: "场景日志、权限区域、角色管理",
			SubFeatures: []Feature{
				{Name: "场景日志", Description: "场景执行日志", APIs: []string{"/api/v1/things/scene/log/get-list"}},
				{Name: "权限区域", Description: "数据权限区域配置", APIs: []string{"/api/v1/things/area/get-list"}},
				{Name: "角色管理", Description: "租户角色管理", APIs: []string{"/api/v1/system/role/info/get-list", "/api/v1/system/role/info/create"}},
			},
		},
	}
}

// consoleFeatures 控制台功能（从前端路由 console.ts 提取）
func consoleFeatures() []Feature {
	return []Feature{
		{
			Name:        "控制台",
			Description: "应用入口和租户切换",
			APIs:        []string{"/api/v1/system/user/self/app/get-list", "/api/v1/system/tenant/info/get-list"},
		},
		{
			Name:        "个人信息",
			Description: "用户个人设置",
			SubFeatures: []Feature{
				{Name: "修改昵称", Description: "修改用户昵称", APIs: []string{"/api/v1/system/user/self/update"}},
				{Name: "修改密码", Description: "修改登录密码", APIs: []string{"/api/v1/system/user/self/update-password"}},
				{Name: "绑定账号", Description: "绑定第三方账号", APIs: []string{"/api/v1/system/user/self/bind-account"}},
				{Name: "我的消息", Description: "查看站内消息", APIs: []string{"/api/v1/system/notify/record/get-list"}},
			},
		},
		{
			Name:        "访问令牌",
			Description: "API 访问令牌管理",
			SubFeatures: []Feature{
				{Name: "创建令牌", Description: "创建 AccessKey/Secret", APIs: []string{"/api/v1/system/user/self/access-token/create"}},
				{Name: "查看令牌", Description: "查看已有令牌", APIs: []string{"/api/v1/system/user/self/access-token/get-list"}},
				{Name: "删除令牌", Description: "删除令牌", APIs: []string{"/api/v1/system/user/self/access-token/delete"}},
			},
		},
		{
			Name:        "续期管理",
			Description: "授权续期和充值",
			APIs:        []string{"/api/v1/system/tenant/renewal"},
		},
	}
}
