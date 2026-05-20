package shared

import (
	"fmt"
	"io"
	"strings"

	"gitee.com/unitedrhino/cli/internal/config"
)

func printHelp(app config.CLIApp, w io.Writer) {
	bin := app.BinaryName()
	features := app.Features()
	featureNames := make([]string, 0, len(features))
	for _, f := range features {
		authNote := ""
		if len(f.Authority) > 0 {
			authNote = fmt.Sprintf("[%s]", strings.Join(f.Authority, "/"))
		}
		featureNames = append(featureNames, f.Name+authNote)
	}

	fmt.Fprintf(w, `%s — 联犀 SaaS 平台 API 命令行工具（%s）

用法:
  %s setup
  %s login [options]
  %s api <path> [options]
  %s token [--decode|--raw]
  %s check
  %s config [--list|--use NAME|tenant [--list|--use CODE]]
  %s schema [path] [--json] [--auth-type CODE]
  %s generate-skills [--output DIR]
  %s scene validate <file>
  %s scene template [auto|manual]
  %s script validate <file>
  %s script template [up-before|up-after|down-before|down-after]
  %s model template [property|event|action|full] [--json|--yaml] [--output file]
  %s model validate <file>
  %s model generate-script <model-file> [--mode type] [--output file]
  %s completion bash|zsh|fish
  %s upgrade [--dry-run] [--version TAG] [--json]
  %s skills list|update|version [--json] [--dry-run]
%s ai-tool artifact get --id <id> [--output-dir <dir>]
%s ai-tool artifact save --id <id> --dir <dir>
%s ai-tool validate --id <id>
%s ai-tool run --id <id> --inputs <json> [--timeout <seconds>]
%s ai-tool edit --id <id> --instruction <text>
%s ai-tool render --id <id> [--output <file>]
%s agg -p PRODUCT_ID -i DATA_ID -f FUNCS [-d DEVICE_NAME] [-j]
%s device <subcommand> [options]
%s user <subcommand> [options]
%s tenant <subcommand> [options]
%s dept <subcommand> [options]
%s alarm <subcommand> [options]
%s area <subcommand> [options]
%s project <subcommand> [options]
%s ota <subcommand> [options]

全局选项:
  --app <name>           指定应用上下文（iot, platform-manage, org-manage, org-energy, console）
                         也可通过 UR_APP 环境变量设置
  --version, -v          显示 CLI 版本号

api 选项:
  --body JSON            请求体 JSON
  --body-file FILE       从文件读取请求体
  --header, -H KEY:VALUE 自定义请求头
  --fields SELECTORS     字段筛选（逗号分隔）
  --summarize            摘要模式（列表只保留前 5 条）
  --format FORMAT        输出格式：json（默认）/ raw / yaml
  --transform PATH       GJSON 路径提取
  --output FILE          将输出保存到文件
  --debug                打印 HTTP 请求/响应详情（敏感头已脱敏）

login 选项:
  --no-wait              请求授权后返回 URL 和 setupCode，不阻塞轮询（AI 模式第 1 步）
  --setup-code <CODE>    用之前的 setupCode 恢复轮询完成授权（AI 模式第 2 步）
  --json                 输出结构化 JSON（配合 --no-wait / --setup-code 使用）
  --base-url <URL>       指定平台地址，跳过交互选择

login 示例:
  # AI 模式：分步授权
  %s login --no-wait --json
  %s login --setup-code ABC123 --json

  # 人类模式：一键阻塞授权
  %s login
  %s login --base-url https://api.example.com

agg 选项:
  -p, --product-id string    产品 ID（必需）
  -d, --device-name string   设备名称（可选，不填则查询产品下所有设备）
  -i, --data-id string       属性标识符（必需，从物模型获取）
  -f, --funcs string         聚合函数，逗号分隔（必需）
                             支持：avg, first, last, count, twa, max, min, sum
      --fill string          数据缺失时的填充模式
      --no-first-ts          不填充最早的时间戳
  -j, --json                 JSON 格式输出

agg 示例:
  # 查询设备 CPU 使用率的平均值
  %s agg -p p_smartswitch_001 -d switch-001 -i CpuUsage -f avg

  # 查询设备温度的最大值和最小值
  %s agg -p p_smartswitch_001 -d switch-001 -i Temperature -f max,min

  # 查询产品下所有设备的平均温度
  %s agg -p p_smartswitch_001 -i Temperature -f avg

  # JSON 格式输出
  %s agg -p p_smartswitch_001 -d switch-001 -i CpuUsage -f avg -j

device 子命令:
  log        查询设备日志（属性、事件、命令、上下线、诊断、异常、SDK）
  control    向设备发送属性控制指令
  action     调用设备行为（send、get、resp）
  mock       生成物模型模拟数据
  report     模拟设备上报消息（通过 HTTP 协议）

device log 子命令:
  property   查询属性日志（最新值、历史记录）
  event      查询事件日志
  send       查询命令日志
  status     查询上下线日志
  hub        查询诊断日志（MQTT 通信）
  abnormal   查询异常日志
  sdk        查询 SDK 日志

device 示例:
  # 查询设备最新属性值
  %s device log property -p p_smartswitch_001 -d switch-001

  # 查询设备温度历史记录
  %s device log property -p p_smartswitch_001 -d switch-001 --data-id Temperature --arg-func avg

  # 控制设备属性
  %s device control -p p_smartswitch_001 -d switch-001 --data '{"PowerSwitch": 1}'

  # 调用设备行为
  %s device action send -p p_smartswitch_001 -d switch-001 --data-id OpenValve --input '{"Duration": 30}'

  # 生成 Mock 数据
  %s device mock -p p_smartswitch_001 -d switch-001 --data-id Temperature --num 5

  # 模拟设备上报
  %s device report -p p_smartswitch_001 -d switch-001 --params '{"Temperature": 25.3}'

应用信息:
  AppID:      %s
  TenantCode: %s
  可调用权限:  %s
  功能模块:    %s

运行时认证环境变量:
  UR_BASE_URL, UR_APP_ID, UR_TENANT_CODE, UR_TOKEN, UR_USER_ID, UR_ACCESS_KEY, UR_ACCESS_SECRET
`, bin, app.DisplayName(),
		bin, bin, bin, bin, bin, bin, bin, bin, bin, bin, bin, bin,
		bin, bin, bin, bin, bin, bin, bin, bin, bin, bin, bin, bin, bin, bin, bin, bin,
		bin, bin, bin, bin, bin, bin, bin,
		app.AppID(),
		func() string {
			if tc := app.DefaultTenantCode(); tc != "" {
				return tc + " (默认)"
			}
			return "用户输入"
		}(),
		strings.Join(app.AllowedAuthTypes(), ", "),
		strings.Join(featureNames, "、"),
	)
}
