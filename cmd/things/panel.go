// 文件说明：注册设备面板离线开发与在线资产/绑定命令；本地命令不初始化平台认证。
package things

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"gitee.com/unitedrhino/cli/internal/client"
	"gitee.com/unitedrhino/cli/internal/paneldev"
	"github.com/spf13/cobra"
)

// newPanelCommand 构造独立命令树，离线工程与在线管理各自按需读取参数。
func newPanelCommand() *cobra.Command {
	root := &cobra.Command{Use: "panel", Short: "HTML 设备面板本地开发、模拟与管理"}
	initCommand := &cobra.Command{Use: "init <directory>", Short: "离线生成面板、SDK、合同及模拟示例", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
		template, _ := cmd.Flags().GetString("template")
		return paneldev.InitTemplate(args[0], template)
	}}
	initCommand.Flags().String("template", "switch", "面板样板：switch / air-conditioner")
	validate := &cobra.Command{Use: "validate <directory|html|zip>", Short: "离线校验包资源、合同与内容摘要", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
		p, e := paneldev.Load(args[0])
		if e != nil {
			return e
		}
		return json.NewEncoder(cmd.OutOrStdout()).Encode(map[string]any{"valid": true, "packageHash": p.Hash, "manifest": p.Manifest, "files": p.Inventory})
	}}
	pack := &cobra.Command{Use: "pack <directory|html|zip>", Short: "离线校验并输出标准 ZIP（不覆盖文件）", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
		output, _ := cmd.Flags().GetString("out")
		p, e := paneldev.Pack(args[0], output)
		if e != nil {
			return e
		}
		return json.NewEncoder(cmd.OutOrStdout()).Encode(map[string]any{"file": output, "packageHash": p.Hash})
	}}
	pack.Flags().String("out", "panel.zip", "输出路径，建议放在工程目录外")
	serve := &cobra.Command{Use: "serve <directory|html|zip>", Short: "本机离线模拟运行，不连接真实设备", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
		port, _ := cmd.Flags().GetInt("port")
		return paneldev.Serve(cmd.Context(), args[0], port, cmd.OutOrStdout())
	}}
	serve.Flags().Int("port", 0, "本机端口，0 自动分配")
	imp := &cobra.Command{Use: "import <directory|html|zip>", Short: "校验上传并导入不可变资产，不自动绑定", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
		// 导入元数据只有显式名称时才覆盖，未提供时保留工程清单。
		name, _ := cmd.Flags().GetString("name")
		desc, _ := cmd.Flags().GetString("desc")
		if (cmd.Flags().Changed("name") && strings.TrimSpace(name) == "") || (desc != "" && strings.TrimSpace(name) == "") {
			return fmt.Errorf("覆盖名称或备注时必须提供非空 --name")
		}
		p, e := paneldev.Load(args[0])
		if e != nil {
			return e
		}
		data, e := p.ZIP()
		if e != nil {
			return e
		}
		upload, e := client.UploadFileMultipart(cmd.Context(), "/api/v1/system/common/upload-file", "file", "panel.zip", data, nil)
		if e != nil {
			return e
		}
		if upload.Code != 200 {
			return fmt.Errorf("上传失败 (%d)：%s", upload.Code, upload.Msg)
		}
		object, ok := upload.Data.(map[string]any)
		if !ok {
			return fmt.Errorf("上传结果无效")
		}
		path, ok := object["filePath"].(string)
		if !ok || path == "" {
			return fmt.Errorf("上传未返回 filePath")
		}
		scope, _ := cmd.Flags().GetString("scope")
		body := map[string]any{"filePath": path}
		if scope != "" {
			body["ownerScope"] = scope
		}
		if name != "" {
			body["name"] = strings.TrimSpace(name)
			body["desc"] = strings.TrimSpace(desc)
		}
		return panelAPI(cmd, "package/import", body)
	}}
	imp.Flags().String("scope", "", "兼容旧调用：平台自动公共，其他企业自动当前企业")
	imp.Flags().String("name", "", "覆盖导入清单名称；省略时保留原名称")
	imp.Flags().String("desc", "", "配合 --name 覆盖备注，空值清除旧备注")
	root.AddCommand(initCommand, validate, pack, serve, imp)
	for _, item := range []struct{ name, path, short string }{{"info", "info/get-one", "查询面板详情"}, {"list", "info/get-list", "查询面板列表"}, {"resolve", "resolve", "解析当前设备面板"}, {"delete", "info/delete", "删除无直接引用的资产"}, {"export", "package/export", "导出标准面板包"}, {"binding", "binding/get-one", "读取本层配置及继承结果"}, {"bind", "binding/update", "按期望版本更新绑定，必要时显式确认影响"}, {"bindings", "binding/get-list", "查询资产直接引用"}} {
		item := item
		command := &cobra.Command{Use: item.name, Short: item.short, Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
			raw, _ := cmd.Flags().GetString("json")
			var body map[string]any
			if e := json.Unmarshal([]byte(raw), &body); e != nil {
				return e
			}
			if body == nil {
				return fmt.Errorf("--json 必须是对象")
			}
			if item.name == "export" {
				response, e := panelRequest(cmd, item.path, body)
				if e != nil {
					return e
				}
				result, ok := response.Data.(map[string]any)
				if !ok {
					return fmt.Errorf("导出结果无效")
				}
				output, _ := cmd.Flags().GetString("out")
				if output != "" {
					url, _ := result["downloadURL"].(string)
					return downloadPanel(cmd, url, output)
				}
				return json.NewEncoder(cmd.OutOrStdout()).Encode(response)
			}
			return panelAPI(cmd, item.path, body)
		}}
		command.Flags().String("json", "{}", "API 请求 JSON；绑定包含 targetType/targetID/mode/expectedRevision")
		if item.name == "export" {
			command.Flags().String("out", "", "下载到本地新 ZIP 文件，不覆盖；省略仅返回短期下载信息")
		}
		root.AddCommand(command)
	}
	return root
}

// panelRequest 复用 CLI 认证和应用上下文，业务失败保留平台原因，不擅自二次确认。
func panelRequest(cmd *cobra.Command, path string, body map[string]any) (client.APIResponse, error) {
	r, e := client.DoAPI(cmd.Context(), client.APIRequest{Path: "/api/v1/things/panel/" + path, Body: body})
	if e == nil && r.Code != 200 {
		e = fmt.Errorf("面板接口失败 (%d)：%s", r.Code, r.Msg)
	}
	return r, e
}

// panelAPI 输出平台完整响应，applied=false 时保留 impactToken 供显式二次提交。
func panelAPI(cmd *cobra.Command, path string, body map[string]any) error {
	r, e := panelRequest(cmd, path, body)
	if e != nil {
		return e
	}
	return json.NewEncoder(cmd.OutOrStdout()).Encode(r)
}

// downloadPanel 下载服务端签发的完整包，限制体积并校验后再原子创建本地文件。
func downloadPanel(cmd *cobra.Command, url, output string) error {
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		return fmt.Errorf("下载地址无效")
	}
	request, e := http.NewRequestWithContext(cmd.Context(), http.MethodGet, url, nil)
	if e != nil {
		return e
	}
	r, e := (&http.Client{Timeout: 30 * time.Second}).Do(request)
	if e != nil {
		return e
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		return fmt.Errorf("下载失败 HTTP %d", r.StatusCode)
	}
	data, e := io.ReadAll(io.LimitReader(r.Body, 10<<20+1))
	if e != nil {
		return e
	}
	temporary, e := os.CreateTemp("", "ur-panel-*.zip")
	if e != nil {
		return e
	}
	name := temporary.Name()
	defer os.Remove(name)
	if _, e = temporary.Write(data); e != nil {
		temporary.Close()
		return e
	}
	temporary.Close()
	p, e := paneldev.Pack(name, output)
	if e != nil {
		return e
	}
	return json.NewEncoder(cmd.OutOrStdout()).Encode(map[string]any{"file": output, "packageHash": p.Hash})
}

// init 注册时不触发网络请求，父命令仅在执行在线子命令时读取认证。
func init() { ThingsCmd.AddCommand(newPanelCommand()) }
