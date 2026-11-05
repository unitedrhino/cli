// 文件说明：提供离线组件目录和在线内置图元查询，统一输出结构化结果与业务错误。
package view

import (
	"encoding/json"
	"fmt"
	"strings"

	"gitee.com/unitedrhino/cli/cmd/view/data"
	"gitee.com/unitedrhino/cli/internal/client"
	"gitee.com/unitedrhino/cli/internal/notice"
	"github.com/spf13/cobra"
)

// newComponentCommand 构造离线组件查询，名称、注册 key 和分类均参与关键词匹配。
func newComponentCommand() *cobra.Command {
	root := &cobra.Command{Use: "component", Short: "离线查询随包组件目录及 IoT 支持信息"}
	search := &cobra.Command{
		Use: "search", Short: "按关键词搜索组件，省略关键词列出全部组件", Args: cobra.NoArgs,
		Example: "  ur view component search --keyword 连线 --json\n  ur view component search --keyword TextCommon --json",
		RunE: func(cmd *cobra.Command, args []string) error {
			snapshot, err := data.Load()
			if err != nil {
				return err
			}
			keyword, _ := cmd.Flags().GetString("keyword")
			keyword = strings.ToLower(strings.TrimSpace(keyword))
			components := make([]data.Component, 0)
			for _, component := range snapshot.Components {
				text := strings.Join([]string{component.Key, component.ChartKey, component.ConKey, component.Title, component.Category, component.CategoryName, component.Package}, " ")
				if strings.Contains(strings.ToLower(text), keyword) {
					components = append(components, component)
				}
			}
			return outputSearchResult(cmd, client.APIResponse{Code: 200, Msg: "success", Data: map[string]any{
				"list": components, "total": len(components), "snapshotDate": snapshot.SnapshotDate, "source": snapshot.Source,
			}})
		},
	}
	search.Flags().String("keyword", "", "匹配中文名称、key、chartKey、conKey 或分类，英文忽略大小写")
	get := &cobra.Command{
		Use: "get-one", Short: "按精确组件 key 查询注册信息及 IoT 支持矩阵", Args: cobra.NoArgs,
		Example: "  ur view component get-one --key ConnectLine --json",
		RunE: func(cmd *cobra.Command, args []string) error {
			key, _ := cmd.Flags().GetString("key")
			if strings.TrimSpace(key) == "" {
				return fmt.Errorf("--key 不能为空")
			}
			if _, err := data.Load(); err != nil {
				return err
			}
			component := data.FindComponent(key)
			if component == nil {
				return outputSearchResult(cmd, client.APIResponse{Code: 404, Msg: "组件不存在: " + key})
			}
			return outputSearchResult(cmd, client.APIResponse{Code: 200, Msg: "success", Data: component})
		},
	}
	get.Flags().String("key", "", "组件注册 key，例如 Image、ConnectLine、TextCommon")
	_ = get.MarkFlagRequired("key")
	root.PersistentFlags().Bool("json", false, "输出 code/msg/data JSON，业务失败仍返回非零退出码")
	root.AddCommand(search, get)
	return root
}

// newBuiltinImageCommand 构造在线图元查询，保留服务端检索排序、分页与原始来源字段。
func newBuiltinImageCommand() *cobra.Command {
	root := &cobra.Command{Use: "builtin-image", Short: "查询平台内置图元，需要已有物联网应用认证"}
	search := &cobra.Command{
		Use: "search", Short: "按关键词、分类及标签分页搜索内置图元", Args: cobra.NoArgs,
		Example: "  ur view builtin-image search --keyword 互感器 --category electric --size 10 --json",
		RunE: func(cmd *cobra.Command, args []string) error {
			page, _ := cmd.Flags().GetInt("page")
			size, _ := cmd.Flags().GetInt("size")
			if page < 1 || size < 1 || size > 100 {
				return fmt.Errorf("--page 必须大于零，--size 必须在 1 到 100 之间")
			}
			body := map[string]any{"page": map[string]any{"page": page, "size": size}}
			for _, field := range []struct{ flag, key string }{
				{"keyword", "keyword"}, {"category", "category"}, {"sub-category", "subCategory"}, {"format", "format"}, {"status", "status"},
			} {
				value, _ := cmd.Flags().GetString(field.flag)
				if value != "" {
					body[field.key] = value
				}
			}
			status, _ := cmd.Flags().GetString("status")
			if status != "" && status != "confirmed" && status != "unconfirmed" {
				return fmt.Errorf("--status 仅支持 confirmed 或 unconfirmed")
			}
			format, _ := cmd.Flags().GetString("format")
			if format != "" && format != "svg" && format != "png" && format != "jpg" && format != "gif" {
				return fmt.Errorf("--format 仅支持 svg、png、jpg 或 gif")
			}
			tags, _ := cmd.Flags().GetStringSlice("tags")
			if len(tags) > 0 {
				body["tags"] = tags
			}
			return requestBuiltinImage(cmd, "get-list", body)
		},
	}
	search.Flags().String("keyword", "", "匹配图元名称、别名、标签和说明")
	search.Flags().String("category", "", "一级分类 ID，例如 electric、symbol、industrial")
	search.Flags().String("sub-category", "", "二级分类名称")
	search.Flags().StringSlice("tags", nil, "标签，逗号分隔，命中任一标签")
	search.Flags().String("format", "", "图片格式：svg、png、jpg、gif")
	search.Flags().String("status", "", "识别状态：confirmed、unconfirmed，使用前仍须看图核对")
	search.Flags().Int("page", 1, "页码，从 1 开始")
	search.Flags().Int("size", 48, "每页数量，1 到 100")
	get := &cobra.Command{
		Use: "get-one", Short: "按稳定图元 ID 获取完整说明、URL 和尺寸", Args: cobra.NoArgs,
		Example: "  ur view builtin-image get-one --id bi-electric-0952 --json",
		RunE: func(cmd *cobra.Command, args []string) error {
			id, _ := cmd.Flags().GetString("id")
			if strings.TrimSpace(id) == "" {
				return fmt.Errorf("--id 不能为空")
			}
			return requestBuiltinImage(cmd, "get-one", map[string]any{"id": id})
		},
	}
	get.Flags().String("id", "", "图库稳定 ID，例如 bi-electric-0952")
	_ = get.MarkFlagRequired("id")
	root.PersistentFlags().Bool("json", false, "输出 code/msg/data JSON，业务失败仍返回非零退出码")
	root.AddCommand(search, get)
	return root
}

// requestBuiltinImage 在物联网应用上下文查询图元；企业与认证继续复用当前配置。
func requestBuiltinImage(cmd *cobra.Command, action string, body map[string]any) error {
	response, err := client.DoAPI(cmd.Context(), client.APIRequest{
		Path: "/api/v1/view/builtin-image/" + action, Body: body, Headers: map[string]string{"app-id": "200"},
	})
	if err != nil {
		return err
	}
	return outputSearchResult(cmd, response)
}

// outputSearchResult 输出结果，JSON 失败保留服务端错误码供 AI 和脚本判断。
func outputSearchResult(cmd *cobra.Command, response client.APIResponse) error {
	jsonOutput, _ := cmd.Flags().GetBool("json")
	if jsonOutput {
		raw, err := notice.MarshalJSON(response, true)
		if err != nil {
			return err
		}
		if _, err := fmt.Fprintln(cmd.OutOrStdout(), string(raw)); err != nil {
			return err
		}
	}
	if response.Code != 200 {
		return fmt.Errorf("查询失败 (%d) %s", response.Code, response.Msg)
	}
	if jsonOutput {
		return nil
	}
	encoder := json.NewEncoder(cmd.OutOrStdout())
	encoder.SetIndent("", "  ")
	return encoder.Encode(response.Data)
}
