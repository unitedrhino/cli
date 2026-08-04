// view.go — 大屏可视化（view）screen 域命令实现
//
// 本文件实现 `ur view screen` 的全部子命令：
//   - 项目 CRUD：get-list / get-one / create / update / delete
//   - 发布管理：publish / unpublish（内部复用 project/update 的 status 字段，无独立发布接口）
//   - 画布迁移：pull（拉取编辑态 content 落盘）/ push（本地 content 校验后推送回远端）
//   - 本地工具：validate（结构 + IoT 绑定校验）/ describe（组件状态摘要）/ screenshot（浏览器截图）
//
// API 前缀为 /api/v1/view/project/，全部 POST，鉴权 app-id=200。
// 画布 content 通过 detail/get-one（forView=false 取编辑态）与 detail/update 读写。
//
// 请求体字段名以后端 viewsvr types.go 为准（backend/things/service/viewsvr/internal/types/types.go）：
//   - project/get-one、project/update、project/delete、project/detail/get-one 的主键字段均为 id
//     （ProjectInfoGetOneReq / ProjectInfoUpdateReq / ProjectInfoDeleteReq，json tag `id,string`）；
//   - project/detail/update 的入参为 ProjectDetail{id, content}；
//   - 仅 project/create 使用 projectID（归属物联网项目 ID，ProjectInfoCreateReq.projectID）。
package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	viewdata "gitee.com/unitedrhino/cli/cmd/view/data"
	"gitee.com/unitedrhino/cli/internal/client"
)

// 大屏项目相关 API 路径常量
const (
	viewProjectCreatePath       = "/api/v1/view/project/create"
	viewProjectUpdatePath       = "/api/v1/view/project/update"
	viewProjectDeletePath       = "/api/v1/view/project/delete"
	viewProjectGetListPath      = "/api/v1/view/project/get-list"
	viewProjectGetOnePath       = "/api/v1/view/project/get-one"
	viewProjectDetailGetOnePath = "/api/v1/view/project/detail/get-one"
	viewProjectDetailUpdatePath = "/api/v1/view/project/detail/update"
)

// 大屏项目发布状态常量（与后端 project/update 的 status 字段约定一致）
const (
	viewScreenStatusPublished   = 1 // 已发布
	viewScreenStatusUnpublished = 2 // 取消发布
)

// 大屏截图默认 URL 模板。
//
// 口径说明（2026-08 核实的仓库事实 + 115 实测）：
//   - goview 包内路由（apps/web/packages/bigscreen/src/goview/router/index.ts）
//     使用 createMemoryHistory，/chart/preview/<id> 与 /chart/home/<id> 是内嵌
//     挂载时的内存路由，不构成外部可访问 URL；
//   - 外部宿主应用（如 apps/web/apps/iot）通过 createBigscreenRoutes 暴露大屏页面，
//     发布页路由为 basePath/published/:id，编辑页为 basePath/editor/:id，
//     同时兼容顶层路由 /view/:id（mount.goview.ts 注释：/view/ 路由始终使用已发布快照）；
//   - iot 应用生产环境 .env.production 配置 VITE_ROUTER_HISTORY=hash，且 115 同域
//     部署的应用路径前缀为 /app/iot/（见 e2e/src/specs/web/iot/auth.setup.ts）；
//   - 115 实测：/iot/big-bigscreen/published/:id 在 iot 应用下会落入应用外壳（不渲染画布），
//     而 /view/:id 兼容路由可正常渲染发布快照，故发布态默认用 /view/:id，
//     编辑态用 /iot/big-bigscreen/editor/:id。
//
// 若部署形态不同，可用 --url-template 覆盖（支持 {base} 与 {id} 占位符）。
const (
	viewScreenPreviewURLTemplate = "{base}/app/iot/#/view/{id}"
	viewScreenEditURLTemplate    = "{base}/app/iot/#/iot/big-bigscreen/editor/{id}"
)

// runViewScreen 执行大屏项目管理命令
func runViewScreen(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printViewScreenHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runViewScreenGetList(ctx, args[1:], stdout, stderr)
	case "get-one":
		return runViewScreenGetOne(ctx, args[1:], stdout, stderr)
	case "create":
		return runViewScreenCreate(ctx, args[1:], stdout, stderr)
	case "update":
		return runViewScreenUpdate(ctx, args[1:], stdout, stderr)
	case "delete":
		return runViewScreenDelete(ctx, args[1:], stdout, stderr)
	case "publish":
		return runViewScreenPublish(ctx, args[1:], stdout, stderr)
	case "unpublish":
		return runViewScreenUnpublish(ctx, args[1:], stdout, stderr)
	case "pull":
		return runViewScreenPull(ctx, args[1:], stdout, stderr)
	case "push":
		return runViewScreenPush(ctx, args[1:], stdout, stderr)
	case "validate":
		return runViewScreenValidate(ctx, args[1:], stdout, stderr)
	case "describe":
		return runViewScreenDescribe(ctx, args[1:], stdout, stderr)
	case "screenshot":
		return runViewScreenScreenshot(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printViewScreenHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown view screen subcommand: %s\n", args[0])
		printViewScreenHelp(stderr)
		return 2
	}
}

// printViewScreenHelp 打印大屏项目管理帮助信息
func printViewScreenHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur view screen <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Bigscreen project management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list     Query screen project list [--body '<json>']")
	fmt.Fprintln(w, "  get-one      Query screen project detail --id <projectID>")
	fmt.Fprintln(w, "  create       Create screen project --body '<json>'")
	fmt.Fprintln(w, "  update       Update screen project --body '<json>'")
	fmt.Fprintln(w, "  delete       Delete screen project --id <projectID>")
	fmt.Fprintln(w, "  publish      Publish screen project --id <projectID>")
	fmt.Fprintln(w, "  unpublish    Unpublish screen project --id <projectID>")
	fmt.Fprintln(w, "  pull         Pull canvas content to local file --id <projectID> [-o <file>]")
	fmt.Fprintln(w, "  push         Push local canvas content to remote -f <file> [--id <projectID>] [--publish] [--force]")
	fmt.Fprintln(w, "  validate     Validate canvas content (-f <file> or --id <projectID>)")
	fmt.Fprintln(w, "  describe     Show component summary (-f <file> or --id <projectID>) [--json]")
	fmt.Fprintln(w, "  screenshot   Take screenshot of screen page --id <projectID> -o <png> [--front-base <url>] [--wait <sec>] [--edit]")
	fmt.Fprintln(w, "  help         Show this help message")
}

// runViewScreenGetList 执行查询大屏项目列表命令
func runViewScreenGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	bodyJSON := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--body":
			if i+1 < len(args) {
				bodyJSON = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", args[i])
			return 2
		}
	}

	// --body 可选：缺省时传空对象，由后端返回默认分页列表
	reqBody, err := parseBodyArg(bodyJSON)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: viewProjectGetListPath,
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runViewScreenGetOne 执行查询大屏项目详情命令
func runViewScreenGetOne(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	projectID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				projectID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", args[i])
			return 2
		}
	}

	if projectID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: viewProjectGetOnePath,
		Body: map[string]any{"id": projectID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runViewScreenCreate 执行创建大屏项目命令
func runViewScreenCreate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return runViewScreenBodyAPI(ctx, args, stdout, stderr, viewProjectCreatePath, "create")
}

// runViewScreenUpdate 执行更新大屏项目命令
func runViewScreenUpdate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return runViewScreenBodyAPI(ctx, args, stdout, stderr, viewProjectUpdatePath, "update")
}

// runViewScreenBodyAPI 是 create/update 子命令的公共实现：解析 --body 并调用指定 API
func runViewScreenBodyAPI(ctx context.Context, args []string, stdout, stderr io.Writer, path, name string) int {
	jsonOutput := false
	bodyJSON := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--body":
			if i+1 < len(args) {
				bodyJSON = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", args[i])
			return 2
		}
	}

	if bodyJSON == "" {
		fmt.Fprintf(stderr, "--body is required for %s\n", name)
		return 2
	}

	reqBody, err := parseBodyArg(bodyJSON)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: path,
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runViewScreenDelete 执行删除大屏项目命令
func runViewScreenDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	projectID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				projectID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", args[i])
			return 2
		}
	}

	if projectID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: viewProjectDeletePath,
		Body: map[string]any{"id": projectID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runViewScreenPublish 执行发布大屏项目命令
//
// 平台没有独立的发布接口，发布即 project/update 将 status 置为 1。
func runViewScreenPublish(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return runViewScreenSetStatus(ctx, args, stdout, stderr, viewScreenStatusPublished, "publish")
}

// runViewScreenUnpublish 执行取消发布大屏项目命令
//
// 取消发布即 project/update 将 status 置为 2。
func runViewScreenUnpublish(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return runViewScreenSetStatus(ctx, args, stdout, stderr, viewScreenStatusUnpublished, "unpublish")
}

// runViewScreenSetStatus 是 publish/unpublish 子命令的公共实现
func runViewScreenSetStatus(ctx context.Context, args []string, stdout, stderr io.Writer, status int, name string) int {
	jsonOutput := false
	projectID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				projectID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", args[i])
			return 2
		}
	}

	if projectID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: viewProjectUpdatePath,
		Body: map[string]any{
			"id":     projectID,
			"status": status,
		},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	if code := outputResult(resp, jsonOutput, stdout, stderr); code != 0 {
		return code
	}
	if !jsonOutput {
		fmt.Fprintf(stdout, "screen %s done: projectID=%s status=%d\n", name, projectID, status)
	}
	return 0
}

// viewScreenMeta 描述 pull 落盘时随画布一同保存的大屏项目元信息
type viewScreenMeta struct {
	// ProjectID 大屏项目 ID
	ProjectID string `json:"projectID"`
	// Name 项目名称
	Name string `json:"name,omitempty"`
	// Desc 项目描述
	Desc string `json:"desc,omitempty"`
	// IndexImage 项目封面图 URL
	IndexImage string `json:"indexImage,omitempty"`
	// Status 项目发布状态（1 已发布 / 2 未发布）
	Status any `json:"status,omitempty"`
}

// runViewScreenPull 执行拉取大屏画布内容到本地文件命令
//
// 通过 detail/get-one（forView=false）取编辑态 content（JSON 字符串），
// 解析后美化缩进写盘；同时把项目元信息（name/desc/indexImage/status）打印到
// stdout 并写入 <output>.meta.json，便于本地识别与 push 时回取 projectID。
func runViewScreenPull(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	projectID := ""
	output := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				projectID = args[i+1]
				i++
			}
		case "--output", "-o":
			if i+1 < len(args) {
				output = args[i+1]
				i++
			}
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", args[i])
			return 2
		}
	}

	if projectID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}
	if output == "" {
		output = fmt.Sprintf("./bigscreen-%s.json", projectID)
	}

	content, code := fetchViewScreenContent(ctx, projectID, stderr)
	if code != 0 {
		return code
	}

	if err := writeViewScreenContentFile(output, content); err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 1
	}
	fmt.Fprintf(stdout, "canvas saved: %s\n", output)

	// 元信息单独从 project/get-one 获取（detail/get-one 响应仅含 content，无项目元信息），
	// 落盘 <output>.meta.json 便于本地识别与 push 时回取项目 ID
	meta := fetchViewScreenMeta(ctx, projectID)
	meta.ProjectID = projectID
	metaPath := output + ".meta.json"
	metaRaw, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		fmt.Fprintf(stderr, "Error: marshal meta: %v\n", err)
		return 1
	}
	if err := os.WriteFile(metaPath, append(metaRaw, '\n'), 0o644); err != nil {
		fmt.Fprintf(stderr, "Error: write meta file: %v\n", err)
		return 1
	}
	fmt.Fprintf(stdout, "meta saved: %s (name=%s status=%v)\n", metaPath, meta.Name, meta.Status)
	return 0
}

// fetchViewScreenMeta 调用 project/get-one 获取大屏项目元信息；
// 失败时仅告警并返回空元信息，不阻断画布落盘
func fetchViewScreenMeta(ctx context.Context, projectID string) viewScreenMeta {
	var meta viewScreenMeta
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: viewProjectGetOnePath,
		Body: map[string]any{"id": projectID},
	})
	if err != nil || resp.Code != 200 {
		return meta
	}
	info, ok := resp.Data.(map[string]any)
	if !ok {
		return meta
	}
	meta.Name, _ = info["name"].(string)
	meta.Desc, _ = info["desc"].(string)
	meta.IndexImage, _ = info["indexImage"].(string)
	meta.Status = info["status"]
	return meta
}

// fetchViewScreenContent 调用 detail/get-one 拉取编辑态画布 content；
// 出错时已向 stderr 打印原因，返回非 0 退出码。
// 响应结构为 ProjectDetail{id, content, publishedContent}，content 为字符串化 JSON。
func fetchViewScreenContent(ctx context.Context, projectID string, stderr io.Writer) (map[string]any, int) {
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: viewProjectDetailGetOnePath,
		Body: map[string]any{
			"id":      projectID,
			"forView": false,
		},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return nil, 1
	}
	if resp.Code != 200 {
		fmt.Fprintf(stderr, "Error: %s\n", resp.Msg)
		return nil, 1
	}

	data, ok := resp.Data.(map[string]any)
	if !ok {
		fmt.Fprintf(stderr, "Error: unexpected detail/get-one response shape\n")
		return nil, 1
	}

	contentStr, _ := data["content"].(string)
	if contentStr == "" {
		fmt.Fprintf(stderr, "Error: detail/get-one response missing content\n")
		return nil, 1
	}

	var content map[string]any
	if err := json.Unmarshal([]byte(contentStr), &content); err != nil {
		fmt.Fprintf(stderr, "Error: content is not valid JSON: %v\n", err)
		return nil, 1
	}
	return content, 0
}

// writeViewScreenContentFile 将画布 content 以美化缩进的 JSON 写入指定文件
func writeViewScreenContentFile(path string, content map[string]any) error {
	raw, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal content: %w", err)
	}
	if err := os.WriteFile(path, append(raw, '\n'), 0o644); err != nil {
		return fmt.Errorf("write file %s: %w", path, err)
	}
	return nil
}

// runViewScreenPush 执行推送本地画布内容到远端命令
//
// 默认先对本地文件执行 validate，存在 error 级问题时拒绝上传（退出码非 0）；
// --force 可跳过校验。projectID 解析顺序：--id 参数 > <file>.meta.json > 文件名
// bigscreen-<id>.json。推送成功且指定 --publish 时，再执行发布（status=1）。
func runViewScreenPush(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	file := ""
	projectID := ""
	publish := false
	force := false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--file", "-f":
			if i+1 < len(args) {
				file = args[i+1]
				i++
			}
		case "--id":
			if i+1 < len(args) {
				projectID = args[i+1]
				i++
			}
		case "--publish":
			publish = true
		case "--force":
			force = true
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", args[i])
			return 2
		}
	}

	if file == "" {
		fmt.Fprintln(stderr, "-f/--file is required")
		return 2
	}

	content, err := readViewScreenContentFile(file)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	// 默认先校验，防止把结构非法的画布推到远端导致发布页渲染失败
	if !force {
		issues := validateScreenContent(content)
		errCount := printViewIssues(issues, stderr)
		if errCount > 0 {
			fmt.Fprintf(stderr, "validate failed: %d error(s), push rejected (use --force to skip validation)\n", errCount)
			return 1
		}
	}

	// projectID 解析：--id > meta 文件 > 文件名
	if projectID == "" {
		projectID = resolveViewScreenProjectID(file)
	}
	if projectID == "" {
		fmt.Fprintln(stderr, "cannot resolve projectID: pass --id, or keep <file>.meta.json alongside, or name the file bigscreen-<id>.json")
		return 2
	}

	// content 需以紧凑字符串化 JSON 提交
	raw, err := json.Marshal(content)
	if err != nil {
		fmt.Fprintf(stderr, "Error: marshal content: %v\n", err)
		return 1
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: viewProjectDetailUpdatePath,
		Body: map[string]any{
			"id":      projectID,
			"content": string(raw),
		},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	if resp.Code != 200 {
		fmt.Fprintf(stderr, "Error: %s\n", resp.Msg)
		return 1
	}
	fmt.Fprintf(stdout, "canvas pushed: projectID=%s file=%s\n", projectID, file)

	if publish {
		resp, err := client.DoAPI(ctx, client.APIRequest{
			Path: viewProjectUpdatePath,
			Body: map[string]any{
				"id":     projectID,
				"status": viewScreenStatusPublished,
			},
		})
		if err != nil {
			fmt.Fprintf(stderr, "API error: %v\n", err)
			return 1
		}
		if resp.Code != 200 {
			fmt.Fprintf(stderr, "Error: publish failed: %s\n", resp.Msg)
			return 1
		}
		fmt.Fprintf(stdout, "screen published: projectID=%s\n", projectID)
	}
	return 0
}

// viewScreenFileNamePattern 匹配 pull 默认输出文件名 bigscreen-<id>.json
var viewScreenFileNamePattern = regexp.MustCompile(`^bigscreen-(.+)\.json$`)

// resolveViewScreenProjectID 从 meta 文件或文件名推断 projectID
func resolveViewScreenProjectID(file string) string {
	// 优先读 <file>.meta.json（pull 落盘的元信息）
	if raw, err := os.ReadFile(file + ".meta.json"); err == nil {
		var meta viewScreenMeta
		if err := json.Unmarshal(raw, &meta); err == nil && meta.ProjectID != "" {
			return meta.ProjectID
		}
	}
	// 其次从 bigscreen-<id>.json 文件名解析
	base := filepath.Base(file)
	if m := viewScreenFileNamePattern.FindStringSubmatch(base); len(m) == 2 {
		return m[1]
	}
	return ""
}

// readViewScreenContentFile 读取本地画布 content 文件并解析为 JSON 对象
func readViewScreenContentFile(path string) (map[string]any, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file %s: %w", path, err)
	}
	var content map[string]any
	if err := json.Unmarshal(raw, &content); err != nil {
		return nil, fmt.Errorf("file %s is not valid JSON object: %w", path, err)
	}
	return content, nil
}

// runViewScreenValidate 执行画布内容校验命令
//
// 校验来源二选一：-f 本地文件，或 --id 拉取远端编辑态 content。
// 存在 error 级问题时退出码非 0；warning 不影响退出码。
func runViewScreenValidate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	file := ""
	projectID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--file", "-f":
			if i+1 < len(args) {
				file = args[i+1]
				i++
			}
		case "--id":
			if i+1 < len(args) {
				projectID = args[i+1]
				i++
			}
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", args[i])
			return 2
		}
	}

	var content map[string]any
	source := ""
	switch {
	case file != "":
		var err error
		content, err = readViewScreenContentFile(file)
		if err != nil {
			fmt.Fprintf(stderr, "Error: %v\n", err)
			return 2
		}
		source = file
	case projectID != "":
		var code int
		content, code = fetchViewScreenContent(ctx, projectID, stderr)
		if code != 0 {
			return code
		}
		source = "remote projectID=" + projectID
	default:
		fmt.Fprintln(stderr, "-f/--file or --id is required")
		return 2
	}

	issues := validateScreenContent(content)
	errCount := printViewIssues(issues, stdout)
	if errCount > 0 {
		fmt.Fprintf(stdout, "validate %s: FAILED (%d error(s))\n", source, errCount)
		return 1
	}
	fmt.Fprintf(stdout, "validate %s: OK (%d warning(s))\n", source, len(issues))
	return 0
}

// printViewIssues 逐条打印校验问题，返回 error 级问题数量
func printViewIssues(issues []viewIssue, w io.Writer) int {
	errCount := 0
	for _, issue := range issues {
		if issue.Level == "error" {
			errCount++
		}
		fmt.Fprintf(w, "[%s] 组件 %s → %s → %s\n", issue.Level, issue.Component, issue.Field, issue.Reason)
	}
	return errCount
}

// runViewScreenDescribe 执行组件状态摘要命令
//
// 每个组件输出一行摘要（index/title/key/chartFrame/坐标尺寸/层级/显隐/锁定/
// 数据绑定类型/IoT 摘要）；--json 输出结构化数组。
func runViewScreenDescribe(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	file := ""
	projectID := ""
	jsonOutput := false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--file", "-f":
			if i+1 < len(args) {
				file = args[i+1]
				i++
			}
		case "--id":
			if i+1 < len(args) {
				projectID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", args[i])
			return 2
		}
	}

	var content map[string]any
	switch {
	case file != "":
		var err error
		content, err = readViewScreenContentFile(file)
		if err != nil {
			fmt.Fprintf(stderr, "Error: %v\n", err)
			return 2
		}
	case projectID != "":
		var code int
		content, code = fetchViewScreenContent(ctx, projectID, stderr)
		if code != 0 {
			return code
		}
	default:
		fmt.Fprintln(stderr, "-f/--file or --id is required")
		return 2
	}

	summaries := describeScreenComponents(content)
	if jsonOutput {
		raw, err := json.MarshalIndent(summaries, "", "  ")
		if err != nil {
			fmt.Fprintf(stderr, "Error: %v\n", err)
			return 1
		}
		fmt.Fprintln(stdout, string(raw))
		return 0
	}

	fmt.Fprintf(stdout, "%-5s %-16s %-14s %-8s %-18s %-7s %-5s %-5s %-6s %s\n",
		"index", "title", "key", "frame", "x,y,w,h", "zIndex", "hide", "lock", "bind", "iot")
	for _, s := range summaries {
		fmt.Fprintf(stdout, "%-5d %-16s %-14s %-8s %-18s %-7s %-5t %-5t %-6s %s\n",
			s.Index, truncateViewString(s.Title, 16), s.Key, s.ChartFrame,
			fmt.Sprintf("%v,%v,%v,%v", s.X, s.Y, s.W, s.H), formatViewNumber(s.ZIndex),
			s.Hide, s.Lock, s.BindType, s.IoTSummary)
	}
	fmt.Fprintf(stdout, "total: %d component(s)\n", len(summaries))
	return 0
}

// truncateViewString 截断字符串到指定长度（按 rune 计算，适配中文标题）
func truncateViewString(s string, n int) string {
	r := []rune(s)
	if len(r) <= n {
		return s
	}
	return string(r[:n-1]) + "…"
}

// formatViewNumber 将 JSON 数值格式化为紧凑字符串（整数去掉小数点）
func formatViewNumber(v any) string {
	f, ok := asViewFloat(v)
	if !ok {
		return "-"
	}
	if f == float64(int64(f)) {
		return strconv.FormatInt(int64(f), 10)
	}
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// runViewScreenScreenshot 执行大屏页面截图命令
//
// 通过本机 agent-browser 打开大屏页面并截图：
//   - 默认截发布态页面，--edit 截编辑态页面（URL 模板见文件头常量说明）；
//   - 前端地址取 --front-base 或环境变量 UR_FRONT_BASE_URL；
//   - agent-browser 不存在时打印 URL 与手动命令提示并返回非 0。
func runViewScreenScreenshot(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	projectID := ""
	output := ""
	frontBase := ""
	urlTemplate := ""
	waitSec := 5
	edit := false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				projectID = args[i+1]
				i++
			}
		case "--output", "-o":
			if i+1 < len(args) {
				output = args[i+1]
				i++
			}
		case "--front-base":
			if i+1 < len(args) {
				frontBase = args[i+1]
				i++
			}
		case "--url-template":
			if i+1 < len(args) {
				urlTemplate = args[i+1]
				i++
			}
		case "--wait":
			if i+1 < len(args) {
				if v, err := strconv.Atoi(args[i+1]); err == nil && v > 0 {
					waitSec = v
				}
				i++
			}
		case "--edit":
			edit = true
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", args[i])
			return 2
		}
	}

	if projectID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}
	if output == "" {
		fmt.Fprintln(stderr, "-o/--output is required")
		return 2
	}
	if frontBase == "" {
		frontBase = os.Getenv("UR_FRONT_BASE_URL")
	}
	if frontBase == "" {
		fmt.Fprintln(stderr, "--front-base or env UR_FRONT_BASE_URL is required")
		return 2
	}

	pageURL := buildViewScreenPageURL(frontBase, projectID, edit, urlTemplate)

	browser, err := exec.LookPath("agent-browser")
	if err != nil {
		fmt.Fprintf(stderr, "agent-browser not found in PATH; please open the URL manually:\n")
		fmt.Fprintf(stderr, "  url: %s\n", pageURL)
		fmt.Fprintf(stderr, "  manual commands (after installing agent-browser):\n")
		fmt.Fprintf(stderr, "    agent-browser open %q\n", pageURL)
		fmt.Fprintf(stderr, "    agent-browser wait %d\n", waitSec*1000)
		fmt.Fprintf(stderr, "    agent-browser screenshot %q\n", output)
		return 1
	}

	// 依次执行：打开页面 → 等待渲染 → 截图
	steps := [][]string{
		{"open", pageURL},
		{"wait", strconv.Itoa(waitSec * 1000)},
		{"screenshot", output},
	}
	for _, step := range steps {
		cmd := exec.CommandContext(ctx, browser, step...)
		cmd.Stdout = stdout
		cmd.Stderr = stderr
		if err := cmd.Run(); err != nil {
			fmt.Fprintf(stderr, "Error: agent-browser %s failed: %v\n", step[0], err)
			return 1
		}
	}
	fmt.Fprintf(stdout, "screenshot saved: %s (url=%s)\n", output, pageURL)
	return 0
}

// buildViewScreenPageURL 按模板拼接大屏页面 URL；模板支持 {base} 与 {id} 占位符
func buildViewScreenPageURL(frontBase, projectID string, edit bool, urlTemplate string) string {
	if urlTemplate == "" {
		if edit {
			urlTemplate = viewScreenEditURLTemplate
		} else {
			urlTemplate = viewScreenPreviewURLTemplate
		}
	}
	url := strings.ReplaceAll(urlTemplate, "{base}", strings.TrimRight(frontBase, "/"))
	url = strings.ReplaceAll(url, "{id}", projectID)
	return url
}

// ────────────────────────────────────────────────────────────
// 以下为画布 content 校验（validate）与摘要（describe）的纯函数实现，供单元测试直接调用
// ────────────────────────────────────────────────────────────

// viewIssue 描述一条画布校验问题
type viewIssue struct {
	// Level 问题级别：error（阻断）/ warning（提示）
	Level string `json:"level"`
	// Component 组件定位（index/id），顶层问题为 "-"
	Component string `json:"component"`
	// Field 出问题的字段路径
	Field string `json:"field"`
	// Reason 问题原因
	Reason string `json:"reason"`
}

// validateScreenContent 校验大屏画布 content，返回全部 error/warning 问题列表。
//
// 校验规则（与任务约定一致）：
//   - 顶层：editCanvasConfig 存在且 width/height 为正数；componentList 为数组（可为空）；
//     requestGlobalConfig 缺失仅给 warning；
//   - 组件级：id 非空且全局唯一；chartConfig.key 非空且存在于组件 key 表；
//     chartKey == "V"+key、conKey == "VC"+key；attr.w/attr.h 为正数，x/y/zIndex 为数值；
//     status.lock/status.hide、events 三件套缺失仅给 warning；
//   - IoT 级：request.requestDataType == 5 时 requestIoTDeviceConfig 必填，其 queryType
//     必须在该组件 iotQueryTypes 支持矩阵内（矩阵为空表示组件不支持 IoT），且
//     productID 非空、deviceNames/dataIDs 为非空数组。
func validateScreenContent(content map[string]any) []viewIssue {
	var issues []viewIssue
	add := func(level, component, field, reason string) {
		issues = append(issues, viewIssue{Level: level, Component: component, Field: field, Reason: reason})
	}

	// ── 顶层校验 ──
	canvas, ok := asViewMap(content["editCanvasConfig"])
	if !ok {
		add("error", "-", "editCanvasConfig", "缺失或不是对象")
	} else {
		for _, dim := range []string{"width", "height"} {
			v, ok := asViewFloat(canvas[dim])
			if !ok {
				add("error", "-", "editCanvasConfig."+dim, "缺失或不是数值")
			} else if v <= 0 {
				add("error", "-", "editCanvasConfig."+dim, "必须为正数")
			}
		}
	}

	components, ok := asViewSlice(content["componentList"])
	if !ok {
		add("error", "-", "componentList", "缺失或不是数组")
	}

	if _, ok := content["requestGlobalConfig"]; !ok {
		add("warning", "-", "requestGlobalConfig", "缺失，组件将使用默认请求配置")
	}

	// ── 组件级校验 ──
	seenIDs := map[string]int{}
	var walk func(list []any)
	walk = func(list []any) {
		for i, item := range list {
			comp, ok := asViewMap(item)
			if !ok {
				add("error", fmt.Sprintf("[%d]", i), "-", "组件项不是对象")
				continue
			}
			id, _ := comp["id"].(string)
			locator := fmt.Sprintf("[%d]/%s", i, id)

			// Group 组件（isGroup=true）没有独立的 chartKey/conKey，
			// 跳过 key 体系校验，递归校验其 groupList 子组件
			if isGroup, _ := comp["isGroup"].(bool); isGroup {
				if id != "" {
					if prev, dup := seenIDs[id]; dup {
						add("error", locator, "id", fmt.Sprintf("与组件 [%d] 重复", prev))
					} else {
						seenIDs[id] = i
					}
				}
				if groupList, ok := asViewSlice(comp["groupList"]); ok {
					walk(groupList)
				}
				continue
			}

			if id == "" {
				add("error", fmt.Sprintf("[%d]", i), "id", "缺失或为空")
			} else if prev, dup := seenIDs[id]; dup {
				add("error", locator, "id", fmt.Sprintf("与组件 [%d] 重复", prev))
			} else {
				seenIDs[id] = i
			}

			validateViewComponent(comp, locator, add)
		}
	}
	walk(components)

	return issues
}

// validateViewComponent 校验单个非分组组件的 key 体系、布局属性与 IoT 绑定
func validateViewComponent(comp map[string]any, locator string, add func(level, component, field, reason string)) {
	// key 取自 chartConfig.key（GoView 组件的真实定义位置）
	chartConfig, _ := asViewMap(comp["chartConfig"])
	key, _ := chartConfig["key"].(string)
	if key == "" {
		add("error", locator, "chartConfig.key", "缺失或为空")
	} else if viewdata.FindComponent(key) == nil {
		add("error", locator, "chartConfig.key", fmt.Sprintf("未知组件 key: %s（不在组件 key 表内）", key))
	}

	// chartKey/conKey 与 key 的前缀一致性：chartKey == "V"+key，conKey == "VC"+key
	// 真实画布内容中二者位于 chartConfig 内（编辑器顶层字段不一定持久化），优先取 chartConfig
	chartKey, _ := comp["chartKey"].(string)
	conKey, _ := comp["conKey"].(string)
	if chartKey == "" {
		chartKey, _ = chartConfig["chartKey"].(string)
	}
	if conKey == "" {
		conKey, _ = chartConfig["conKey"].(string)
	}
	if key != "" {
		if chartKey != "V"+key {
			add("error", locator, "chartKey", fmt.Sprintf("应为 %q，实际 %q", "V"+key, chartKey))
		}
		if conKey != "VC"+key {
			add("error", locator, "conKey", fmt.Sprintf("应为 %q，实际 %q", "VC"+key, conKey))
		}
	}

	// 布局属性：w/h 必须为正数，x/y/zIndex 必须为数值
	attr, ok := asViewMap(comp["attr"])
	if !ok {
		add("error", locator, "attr", "缺失或不是对象")
	} else {
		for _, dim := range []string{"w", "h"} {
			v, ok := asViewFloat(attr[dim])
			if !ok {
				add("error", locator, "attr."+dim, "缺失或不是数值")
			} else if v <= 0 {
				add("error", locator, "attr."+dim, "必须为正数")
			}
		}
		for _, field := range []string{"x", "y", "zIndex"} {
			if _, ok := asViewFloat(attr[field]); !ok {
				add("error", locator, "attr."+field, "缺失或不是数值")
			}
		}
	}

	// status / events 三件套缺失仅告警，不阻断
	status, ok := asViewMap(comp["status"])
	if !ok {
		add("warning", locator, "status", "缺失，默认按未锁定未隐藏处理")
	} else {
		for _, field := range []string{"lock", "hide"} {
			if _, ok := status[field].(bool); !ok {
				add("warning", locator, "status."+field, "缺失或不是布尔值")
			}
		}
	}
	events, ok := asViewMap(comp["events"])
	if !ok {
		add("warning", locator, "events", "缺失，组件无事件配置")
	} else {
		for _, field := range []string{"baseEvent", "advancedEvents", "interactEvents"} {
			if _, ok := events[field]; !ok {
				add("warning", locator, "events."+field, "缺失")
			}
		}
	}

	// 渲染器硬要求：styles.animations 缺失会导致页面渲染崩溃（animationsClass 报错）
	styles, ok := asViewMap(comp["styles"])
	if !ok {
		add("error", locator, "styles", "缺失，渲染器要求 styles 对象（含 animations 数组）")
	} else if _, ok := asViewSlice(styles["animations"]); !ok {
		add("error", locator, "styles.animations", "缺失或不是数组，缺失会导致页面渲染崩溃")
	}

	// ECharts 组件 option 过简时图表无法出图（只画边框），仅告警提示补齐完整配置
	if chartFrame, _ := chartConfig["chartFrame"].(string); chartFrame == "echarts" {
		if option, ok := asViewMap(comp["option"]); ok && len(option) < 2 {
			add("warning", locator, "option",
				"ECharts 组件 option 仅含少量字段，可能无法正常出图，建议从组件默认配置拷贝完整 option")
		}
	}

	// IoT 绑定校验：requestDataType == 5 表示 IoT 设备数据源
	request, _ := asViewMap(comp["request"])
	dataType, _ := asViewFloat(request["requestDataType"])
	if int(dataType) != 5 {
		return
	}
	iotCfg, ok := asViewMap(request["requestIoTDeviceConfig"])
	if !ok {
		add("error", locator, "request.requestIoTDeviceConfig", "requestDataType=5（IoT）时必须配置")
		return
	}
	queryType, _ := iotCfg["queryType"].(string)
	if key != "" {
		if meta := viewdata.FindComponent(key); meta != nil {
			if len(meta.IoTQueryTypes) == 0 {
				add("error", locator, "request.requestIoTDeviceConfig.queryType",
					fmt.Sprintf("组件 %s 不支持 IoT 数据绑定", key))
			} else if !viewdata.SupportsIoTQueryType(key, queryType) {
				add("error", locator, "request.requestIoTDeviceConfig.queryType",
					fmt.Sprintf("组件 %s 不支持 queryType=%q（支持: %s）", key, queryType, strings.Join(meta.IoTQueryTypes, ",")))
			}
		}
	}
	if productID, _ := iotCfg["productID"].(string); productID == "" {
		add("error", locator, "request.requestIoTDeviceConfig.productID", "缺失或为空")
	}
	// deviceStatus 查询按产品维度统计在线状态，不需要 deviceNames/dataIDs（真实画布即如此）
	if queryType != "deviceStatus" {
		if deviceNames, ok := asViewSlice(iotCfg["deviceNames"]); !ok || len(deviceNames) == 0 {
			add("error", locator, "request.requestIoTDeviceConfig.deviceNames", "必须为非空数组")
		}
		if dataIDs, ok := asViewSlice(iotCfg["dataIDs"]); !ok || len(dataIDs) == 0 {
			add("error", locator, "request.requestIoTDeviceConfig.dataIDs", "必须为非空数组")
		}
	}
}

// viewComponentSummary 描述单个组件的状态摘要（describe 输出用）
type viewComponentSummary struct {
	// Index 组件在 componentList 中的下标
	Index int `json:"index"`
	// ID 组件 ID
	ID string `json:"id"`
	// Title 组件标题
	Title string `json:"title"`
	// Key 组件 key
	Key string `json:"key"`
	// ChartFrame 渲染框架（来自组件 key 表）
	ChartFrame string `json:"chartFrame"`
	// X/Y/W/H 组件坐标与尺寸
	X any `json:"x"`
	Y any `json:"y"`
	W any `json:"w"`
	H any `json:"h"`
	// ZIndex 层级
	ZIndex any `json:"zIndex"`
	// Hide 是否隐藏
	Hide bool `json:"hide"`
	// Lock 是否锁定
	Lock bool `json:"lock"`
	// BindType 数据绑定类型（0 静态 / 1 AJAX / 2 Pond / 5 IoT）
	BindType string `json:"bindType"`
	// IoTSummary IoT 绑定摘要（queryType/productID/deviceNames/dataIDs/chartMode）
	IoTSummary string `json:"iotSummary,omitempty"`
}

// 数据绑定类型码表（request.requestDataType）
var viewBindTypeNames = map[int]string{
	0: "0静态",
	1: "1AJAX",
	2: "2Pond",
	5: "5IoT",
}

// describeScreenComponents 生成画布全部组件的状态摘要列表
func describeScreenComponents(content map[string]any) []viewComponentSummary {
	var out []viewComponentSummary
	components, _ := asViewSlice(content["componentList"])

	var walk func(list []any)
	walk = func(list []any) {
		for _, item := range list {
			comp, ok := asViewMap(item)
			if !ok {
				continue
			}
			// Group 组件只展开其子组件，自身不作为独立摘要行
			if isGroup, _ := comp["isGroup"].(bool); isGroup {
				if groupList, ok := asViewSlice(comp["groupList"]); ok {
					walk(groupList)
				}
				continue
			}
			out = append(out, summarizeViewComponent(comp, len(out)))
		}
	}
	walk(components)
	return out
}

// summarizeViewComponent 提取单个组件的摘要信息
func summarizeViewComponent(comp map[string]any, index int) viewComponentSummary {
	s := viewComponentSummary{Index: index}
	s.ID, _ = comp["id"].(string)
	s.Title, _ = comp["title"].(string)

	chartConfig, _ := asViewMap(comp["chartConfig"])
	s.Key, _ = chartConfig["key"].(string)
	if meta := viewdata.FindComponent(s.Key); meta != nil {
		s.ChartFrame = meta.ChartFrame
		if s.Title == "" {
			s.Title = meta.Title
		}
	}

	attr, _ := asViewMap(comp["attr"])
	s.X, s.Y, s.W, s.H = attr["x"], attr["y"], attr["w"], attr["h"]
	s.ZIndex = attr["zIndex"]

	status, _ := asViewMap(comp["status"])
	s.Lock, _ = status["lock"].(bool)
	s.Hide, _ = status["hide"].(bool)

	request, _ := asViewMap(comp["request"])
	dataType, _ := asViewFloat(request["requestDataType"])
	if name, ok := viewBindTypeNames[int(dataType)]; ok {
		s.BindType = name
	} else {
		s.BindType = fmt.Sprintf("%d", int(dataType))
	}

	// IoT 绑定组件输出配置摘要，便于快速核对绑定关系
	if int(dataType) == 5 {
		if iotCfg, ok := asViewMap(request["requestIoTDeviceConfig"]); ok {
			queryType, _ := iotCfg["queryType"].(string)
			productID, _ := iotCfg["productID"].(string)
			chartMode, _ := iotCfg["chartMode"].(string)
			s.IoTSummary = fmt.Sprintf("queryType=%s productID=%s deviceNames=%s dataIDs=%s chartMode=%s",
				queryType, productID,
				formatViewStringSlice(iotCfg["deviceNames"]),
				formatViewStringSlice(iotCfg["dataIDs"]),
				chartMode)
		}
	}
	return s
}

// formatViewStringSlice 将 JSON 数组格式化为 [a,b] 形式的紧凑字符串
func formatViewStringSlice(v any) string {
	list, ok := asViewSlice(v)
	if !ok || len(list) == 0 {
		return "[]"
	}
	parts := make([]string, 0, len(list))
	for _, item := range list {
		if s, ok := item.(string); ok {
			parts = append(parts, s)
		}
	}
	return "[" + strings.Join(parts, ",") + "]"
}

// asViewMap 将 JSON 值断言为对象
func asViewMap(v any) (map[string]any, bool) {
	m, ok := v.(map[string]any)
	return m, ok
}

// asViewSlice 将 JSON 值断言为数组
func asViewSlice(v any) ([]any, bool) {
	s, ok := v.([]any)
	return s, ok
}

// asViewFloat 将 JSON 值断言为数值（encoding/json 默认解析为 float64）
func asViewFloat(v any) (float64, bool) {
	f, ok := v.(float64)
	return f, ok
}
