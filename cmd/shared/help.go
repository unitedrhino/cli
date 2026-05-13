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
  %s config [--list|--use NAME]
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

应用信息:
  AppID:      %s
  TenantCode: %s
  可调用权限:  %s
  功能模块:    %s

运行时认证环境变量:
  UR_BASE_URL, UR_APP_ID, UR_TENANT_CODE, UR_TOKEN, UR_USER_ID, UR_ACCESS_KEY, UR_ACCESS_SECRET
`, bin, app.DisplayName(),
		bin, bin, bin, bin, bin, bin, bin, bin, bin, bin, bin, bin,
		bin, bin, bin, bin, bin, bin, bin, bin, bin, bin,
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
