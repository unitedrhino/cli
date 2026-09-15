// flow.go — 流程审批中心命令：管理侧（流程定义/表单模板/流程分类）。
// 管理操作对应后端 authType=admin 接口，仅租户管理员可执行；
// 审批动线（发起/待办/办理）见 flow_task.go。
package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"gitee.com/unitedrhino/cli/internal/client"
)

const flowAPIBase = "/api/v1/system/flow"

// runFlow 执行流程审批中心命令分发
func runFlow(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printFlowHelp(stdout)
		return 0
	}

	switch args[0] {
	case "def":
		return runFlowDef(ctx, args[1:], stdout, stderr)
	case "form":
		return runFlowForm(ctx, args[1:], stdout, stderr)
	case "category":
		return runFlowCategory(ctx, args[1:], stdout, stderr)
	case "process":
		return runFlowProcess(ctx, args[1:], stdout, stderr)
	case "task":
		return runFlowTask(ctx, args[1:], stdout, stderr)
	case "instance":
		return runFlowInstance(ctx, args[1:], stdout, stderr)
	case "cc":
		return runFlowCc(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printFlowHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown flow subcommand: %s\n", args[0])
		printFlowHelp(stderr)
		return 2
	}
}

// printFlowHelp 打印流程审批中心帮助信息
func printFlowHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur flow <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "流程审批中心 (/api/v1/system/flow/*)")
	fmt.Fprintln(w, "管理操作（def/form/category/instance 管控）仅租户管理员可执行")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  def        流程定义管理，仅管理员 (get-list, get-one, create, update, delete, deploy, unpublish)")
	fmt.Fprintln(w, "  form       表单模板管理，仅管理员 (get-list, get-one, create, update, delete)")
	fmt.Fprintln(w, "  category   流程分类管理，仅管理员 (get-list, create, update, delete)")
	fmt.Fprintln(w, "  process    流程发起（全员）(get-launch-list, launch)")
	fmt.Fprintln(w, "  task       审批任务办理（全员）(get-pending-list, get-approved-list, get-detail, consent, reject, transfer, add-sign, remove-sign)")
	fmt.Fprintln(w, "  instance   流程实例 (get-my-list, get-record, revoke 全员；get-monitor-list, terminate, resume, destroy 仅管理员)")
	fmt.Fprintln(w, "  cc         我的抄送（全员）(get-list)")
	fmt.Fprintln(w, "  help       显示本帮助")
}

// ─── 流程定义管理（仅管理员） ────────────────────────────────────────────

// runFlowDef 流程定义命令分发
func runFlowDef(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printFlowDefHelp(stdout)
		return 0
	}
	switch args[0] {
	case "get-list":
		return runFlowDefGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runFlowDefGetOne(ctx, args[1:], stdout, stderr)
	case "create":
		return runFlowDefCreate(ctx, args[1:], stdout, stderr)
	case "update":
		return runFlowDefUpdate(ctx, args[1:], stdout, stderr)
	case "delete":
		return runFlowDefDelete(ctx, args[1:], stdout, stderr)
	case "deploy":
		return runFlowDefDeploy(ctx, args[1:], stdout, stderr)
	case "unpublish":
		return runFlowDefUnpublish(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printFlowDefHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown def subcommand: %s\n", args[0])
		printFlowDefHelp(stderr)
		return 2
	}
}

// printFlowDefHelp 打印流程定义帮助
func printFlowDefHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur flow def <subcommand> [options]  （仅租户管理员）")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   查询流程定义列表 [--code] [--name] [--category] [--state] [--page] [--size]")
	fmt.Fprintln(w, "  get-one    查询流程定义详情（含模型JSON）--id <id>")
	fmt.Fprintln(w, "  create     创建流程定义（默认草稿）--code <编码> --name <名称> [--model <JSON或@file>] [--category] [--form-id] [--remark]")
	fmt.Fprintln(w, "  update     更新流程定义 --id <id> --code <编码> [--name] [--model] [--category] [--form-id] [--remark]")
	fmt.Fprintln(w, "  delete     删除流程定义（有运行实例时被拒）--id <id>")
	fmt.Fprintln(w, "  deploy     发布流程（需已绑定表单）--id <id>")
	fmt.Fprintln(w, "  unpublish  停用已发布流程 --id <id>")
}

// runFlowDefGetList 查询流程定义列表
func runFlowDefGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, remaining := parseInfoListParams(args)
	reqBody := map[string]any{"page": map[string]any{"page": page, "size": size}}
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--code":
			reqBody["processCode"] = nextArg(remaining, &i)
		case "--name":
			reqBody["name"] = nextArg(remaining, &i)
		case "--category":
			reqBody["category"] = nextArg(remaining, &i)
		case "--state":
			reqBody["state"] = nextArg(remaining, &i)
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + "/def/get-list", Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runFlowDefGetOne 查询流程定义详情
func runFlowDefGetOne(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, _, _, remaining := parseInfoListParams(args)
	reqBody := map[string]any{}
	for i := 0; i < len(remaining); i++ {
		if remaining[i] == "--id" {
			reqBody["id"] = nextArg(remaining, &i)
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + "/def/get-one", Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runFlowDefCreate 创建流程定义
func runFlowDefCreate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, _, _, remaining := parseInfoListParams(args)
	reqBody := map[string]any{}
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--code":
			reqBody["processCode"] = nextArg(remaining, &i)
		case "--name":
			reqBody["name"] = nextArg(remaining, &i)
		case "--model":
			v, err := readJSONArg(nextArg(remaining, &i))
			if err != nil {
				fmt.Fprintf(stderr, "Error: --model %v\n", err)
				return 1
			}
			reqBody["modelContent"] = v
		case "--category":
			reqBody["category"] = nextArg(remaining, &i)
		case "--form-id":
			reqBody["formTemplateId"] = nextArg(remaining, &i)
		case "--remark":
			reqBody["remark"] = nextArg(remaining, &i)
		case "--sort":
			reqBody["sort"] = nextArg(remaining, &i)
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + "/def/create", Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runFlowDefUpdate 更新流程定义
func runFlowDefUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, _, _, remaining := parseInfoListParams(args)
	reqBody := map[string]any{}
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--id":
			reqBody["id"] = nextArg(remaining, &i)
		case "--code":
			reqBody["processCode"] = nextArg(remaining, &i)
		case "--name":
			reqBody["name"] = nextArg(remaining, &i)
		case "--model":
			v, err := readJSONArg(nextArg(remaining, &i))
			if err != nil {
				fmt.Fprintf(stderr, "Error: --model %v\n", err)
				return 1
			}
			reqBody["modelContent"] = v
		case "--category":
			reqBody["category"] = nextArg(remaining, &i)
		case "--form-id":
			reqBody["formTemplateId"] = nextArg(remaining, &i)
		case "--remark":
			reqBody["remark"] = nextArg(remaining, &i)
		case "--sort":
			reqBody["sort"] = nextArg(remaining, &i)
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + "/def/update", Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runFlowDefDelete 删除流程定义
func runFlowDefDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return flowSimpleIDAction(ctx, "/def/delete", args, stdout, stderr)
}

// runFlowDefDeploy 发布流程定义
func runFlowDefDeploy(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return flowSimpleIDAction(ctx, "/def/deploy", args, stdout, stderr)
}

// runFlowDefUnpublish 停用已发布流程定义
func runFlowDefUnpublish(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return flowSimpleIDAction(ctx, "/def/unpublish", args, stdout, stderr)
}

// ─── 表单模板管理（仅管理员） ────────────────────────────────────────────

// runFlowForm 表单模板命令分发
func runFlowForm(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printFlowFormHelp(stdout)
		return 0
	}
	switch args[0] {
	case "get-list":
		return runFlowFormGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runFlowFormGetOne(ctx, args[1:], stdout, stderr)
	case "create":
		return runFlowFormCreate(ctx, args[1:], stdout, stderr)
	case "update":
		return runFlowFormUpdate(ctx, args[1:], stdout, stderr)
	case "delete":
		return runFlowFormDelete(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printFlowFormHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown form subcommand: %s\n", args[0])
		printFlowFormHelp(stderr)
		return 2
	}
}

// printFlowFormHelp 打印表单模板帮助
func printFlowFormHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur flow form <subcommand> [options]  （仅租户管理员）")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   查询表单模板列表 [--name] [--category] [--page] [--size]")
	fmt.Fprintln(w, "  get-one    查询表单模板详情（含表单JSON schema）--id <id>")
	fmt.Fprintln(w, "  create     创建表单模板 --name <名称> --schema <JSON或@file> [--category] [--remark]")
	fmt.Fprintln(w, "  update     更新表单模板 --id <id> --name <名称> [--schema] [--category] [--remark]")
	fmt.Fprintln(w, "  delete     删除表单模板（被流程引用时被拒）--id <id>")
}

// runFlowFormGetList 查询表单模板列表
func runFlowFormGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, remaining := parseInfoListParams(args)
	reqBody := map[string]any{"page": map[string]any{"page": page, "size": size}}
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--name":
			reqBody["name"] = nextArg(remaining, &i)
		case "--category":
			reqBody["category"] = nextArg(remaining, &i)
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + "/form/get-list", Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runFlowFormGetOne 查询表单模板详情
func runFlowFormGetOne(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return flowGetOneByID(ctx, "/form/get-one", args, stdout, stderr)
}

// runFlowFormCreate 创建表单模板
func runFlowFormCreate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, _, _, remaining := parseInfoListParams(args)
	reqBody := map[string]any{}
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--name":
			reqBody["name"] = nextArg(remaining, &i)
		case "--schema":
			v, err := readJSONArg(nextArg(remaining, &i))
			if err != nil {
				fmt.Fprintf(stderr, "Error: --schema %v\n", err)
				return 1
			}
			reqBody["formSchema"] = v
		case "--category":
			reqBody["category"] = nextArg(remaining, &i)
		case "--remark":
			reqBody["remark"] = nextArg(remaining, &i)
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + "/form/create", Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runFlowFormUpdate 更新表单模板
func runFlowFormUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, _, _, remaining := parseInfoListParams(args)
	reqBody := map[string]any{}
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--id":
			reqBody["id"] = nextArg(remaining, &i)
		case "--name":
			reqBody["name"] = nextArg(remaining, &i)
		case "--schema":
			v, err := readJSONArg(nextArg(remaining, &i))
			if err != nil {
				fmt.Fprintf(stderr, "Error: --schema %v\n", err)
				return 1
			}
			reqBody["formSchema"] = v
		case "--category":
			reqBody["category"] = nextArg(remaining, &i)
		case "--remark":
			reqBody["remark"] = nextArg(remaining, &i)
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + "/form/update", Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runFlowFormDelete 删除表单模板
func runFlowFormDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return flowSimpleIDAction(ctx, "/form/delete", args, stdout, stderr)
}

// ─── 流程分类管理（仅管理员） ────────────────────────────────────────────

// runFlowCategory 流程分类命令分发
func runFlowCategory(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printFlowCategoryHelp(stdout)
		return 0
	}
	switch args[0] {
	case "get-list":
		return runFlowCategoryGetList(ctx, args[1:], stdout, stderr)
	case "create":
		return runFlowCategoryCreate(ctx, args[1:], stdout, stderr)
	case "update":
		return runFlowCategoryUpdate(ctx, args[1:], stdout, stderr)
	case "delete":
		return runFlowCategoryDelete(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printFlowCategoryHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown category subcommand: %s\n", args[0])
		printFlowCategoryHelp(stderr)
		return 2
	}
}

// printFlowCategoryHelp 打印流程分类帮助
func printFlowCategoryHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur flow category <subcommand> [options]  （仅租户管理员）")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   查询分类列表 [--page] [--size]（后端缺省每页20条，分类多时用 --size 拉全）")
	fmt.Fprintln(w, "  create     创建分类 --name <名称> --code <编码> [--sort] [--remark]")
	fmt.Fprintln(w, "  update     更新分类 --id <id> --name <名称> [--code] [--sort]")
	fmt.Fprintln(w, "  delete     删除分类 --id <id>")
}

// runFlowCategoryGetList 查询流程分类列表
func runFlowCategoryGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, _ := parseInfoListParams(args)
	reqBody := map[string]any{"page": map[string]any{"page": page, "size": size}}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + "/category/get-list", Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runFlowCategoryCreate 创建流程分类
func runFlowCategoryCreate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, _, _, remaining := parseInfoListParams(args)
	reqBody := map[string]any{}
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--name":
			reqBody["name"] = nextArg(remaining, &i)
		case "--code":
			reqBody["code"] = nextArg(remaining, &i)
		case "--sort":
			reqBody["sort"] = nextArg(remaining, &i)
		case "--remark":
			reqBody["remark"] = nextArg(remaining, &i)
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + "/category/create", Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runFlowCategoryUpdate 更新流程分类
func runFlowCategoryUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, _, _, remaining := parseInfoListParams(args)
	reqBody := map[string]any{}
	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--id":
			reqBody["id"] = nextArg(remaining, &i)
		case "--name":
			reqBody["name"] = nextArg(remaining, &i)
		case "--code":
			reqBody["code"] = nextArg(remaining, &i)
		case "--sort":
			reqBody["sort"] = nextArg(remaining, &i)
		case "--remark":
			reqBody["remark"] = nextArg(remaining, &i)
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + "/category/update", Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runFlowCategoryDelete 删除流程分类
func runFlowCategoryDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return flowSimpleIDAction(ctx, "/category/delete", args, stdout, stderr)
}

// ─── 共享小工具 ──────────────────────────────────────────────────────────

// readJSONArg 解析 JSON 参数值：@开头视为文件路径读取文件内容，否则视为内联 JSON 字符串。
// 统一校验为合法 JSON 对象后返回原文（后端契约接收 JSON 字符串而非对象）。
func readJSONArg(v string) (string, error) {
	var raw []byte
	if strings.HasPrefix(v, "@") {
		var err error
		raw, err = os.ReadFile(strings.TrimPrefix(v, "@"))
		if err != nil {
			return "", fmt.Errorf("读取文件失败: %w", err)
		}
	} else {
		raw = []byte(v)
	}
	var obj map[string]any
	if err := json.Unmarshal(raw, &obj); err != nil {
		return "", fmt.Errorf("不是合法的JSON对象: %w", err)
	}
	return string(raw), nil
}

// nextArg 取当前位置的下一个参数值（未提供时返回空串）
func nextArg(args []string, i *int) string {
	if *i+1 < len(args) {
		*i++
		return args[*i]
	}
	return ""
}

// flowSimpleIDAction 处理仅需 --id 的简单动作（delete/deploy/unpublish 等）
func flowSimpleIDAction(ctx context.Context, path string, args []string, stdout, stderr io.Writer) int {
	jsonOutput, _, _, remaining := parseInfoListParams(args)
	reqBody := map[string]any{}
	for i := 0; i < len(remaining); i++ {
		if remaining[i] == "--id" {
			reqBody["id"] = nextArg(remaining, &i)
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + path, Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}

// flowGetOneByID 处理仅需 --id 的详情查询
func flowGetOneByID(ctx context.Context, path string, args []string, stdout, stderr io.Writer) int {
	jsonOutput, _, _, remaining := parseInfoListParams(args)
	reqBody := map[string]any{}
	for i := 0; i < len(remaining); i++ {
		if remaining[i] == "--id" {
			reqBody["id"] = nextArg(remaining, &i)
		}
	}
	resp, err := client.DoAPI(ctx, client.APIRequest{Path: flowAPIBase + path, Body: reqBody})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	return outputResult(resp, jsonOutput, stdout, stderr)
}
