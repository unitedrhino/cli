// view_asset.go — 大屏可视化（view）asset 资源库域命令实现
//
// 本文件实现 `ur view asset` 的全部子命令：
//   - get-list：查询资源列表（name/type/groupId/format 筛选 + 分页）
//   - upload：两步链路上传资源——先经 /api/v1/system/common/upload-file 拿到文件 URL，
//     再调 /api/v1/view/asset/upload 登记资源记录
//   - delete：按 id 删除资源
//
// API 前缀为 /api/v1/view/asset/，全部 POST，鉴权 app-id=200。
package shared

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gitee.com/unitedrhino/cli/internal/client"
)

// 大屏资源库相关 API 路径常量
const (
	viewAssetGetListPath = "/api/v1/view/asset/get-list"
	viewAssetUploadPath  = "/api/v1/view/asset/upload"
	viewAssetDeletePath  = "/api/v1/view/asset/delete"

	// 通用文件上传接口（asset upload 第一步：拿文件 URL）
	systemCommonUploadFilePath = "/api/v1/system/common/upload-file"
)

// runViewAsset 执行大屏资源库管理命令
func runViewAsset(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printViewAssetHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runViewAssetGetList(ctx, args[1:], stdout, stderr)
	case "upload":
		return runViewAssetUpload(ctx, args[1:], stdout, stderr)
	case "delete":
		return runViewAssetDelete(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printViewAssetHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown view asset subcommand: %s\n", args[0])
		printViewAssetHelp(stderr)
		return 2
	}
}

// printViewAssetHelp 打印大屏资源库管理帮助信息
func printViewAssetHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur view asset <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Bigscreen asset library management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list   Query asset list [--name <name>] [--type <image|video|audio|other>] [--group-id <id>] [--format <ext>] [--page <n>] [--size <n>]")
	fmt.Fprintln(w, "  upload     Upload asset -f <file> [--name <name>] [--group-id <id>] [--tags <a,b>]")
	fmt.Fprintln(w, "  delete     Delete asset --id <assetID>")
	fmt.Fprintln(w, "  help       Show this help message")
}

// runViewAssetGetList 执行查询大屏资源列表命令
func runViewAssetGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	page, size := 1, 10
	reqBody := map[string]any{}
	for i := 0; i < len(args); i++ {
		next := func() string {
			if i+1 < len(args) {
				i++
				return args[i]
			}
			return ""
		}
		switch args[i] {
		case "--name":
			if v := next(); v != "" {
				reqBody["name"] = v
			}
		case "--type":
			if v := next(); v != "" {
				reqBody["type"] = v
			}
		case "--group-id":
			if v := next(); v != "" {
				reqBody["groupId"] = v
			}
		case "--format":
			if v := next(); v != "" {
				reqBody["format"] = v
			}
		case "--page":
			if v, err := strconv.Atoi(next()); err == nil && v > 0 {
				page = v
			}
		case "--size":
			if v, err := strconv.Atoi(next()); err == nil && v > 0 {
				size = v
			}
		case "--json", "-j":
			jsonOutput = true
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", args[i])
			return 2
		}
	}
	reqBody["page"] = map[string]any{"page": page, "size": size}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: viewAssetGetListPath,
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runViewAssetUpload 执行上传大屏资源命令
//
// 两步链路：
//  1. 以 multipart/form-data 调 /api/v1/system/common/upload-file（表单字段 file）拿到文件 URL；
//  2. 按扩展名推断资源类型（image/video/audio/other），调 /api/v1/view/asset/upload
//     登记资源记录（name/url/size/type/groupId/tags，tags 为逗号分隔字符串）。
func runViewAssetUpload(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	file := ""
	name := ""
	groupID := ""
	tags := ""
	jsonOutput := false
	for i := 0; i < len(args); i++ {
		next := func() string {
			if i+1 < len(args) {
				i++
				return args[i]
			}
			return ""
		}
		switch args[i] {
		case "--file", "-f":
			file = next()
		case "--name":
			name = next()
		case "--group-id":
			groupID = next()
		case "--tags":
			tags = next()
		case "--json", "-j":
			jsonOutput = true
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", args[i])
			return 2
		}
	}

	if file == "" {
		fmt.Fprintln(stderr, "-f/--file is required")
		return 2
	}

	// 检查文件并读取内容
	fi, err := os.Stat(file)
	if err != nil {
		fmt.Fprintf(stderr, "file not found: %v\n", err)
		return 2
	}
	fileData, err := os.ReadFile(file)
	if err != nil {
		fmt.Fprintf(stderr, "read file error: %v\n", err)
		return 1
	}

	// 第一步：以 multipart/form-data 上传文件拿 URL（接口字段名为 file）。
	// 大屏素材需公开桶永久 URL（前端封面/素材上传口径一致）：isPublic=true +
	// business=view + scene=goView/xxx + useBy=user
	uploadResp, err := client.UploadFileMultipart(ctx, systemCommonUploadFilePath, "file", filepath.Base(file), fileData, map[string]string{
		"isPublic": "true",
		"business": "view",
		"scene":    "goView/asset",
		"useBy":    "user",
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	if uploadResp.Code != 200 {
		fmt.Fprintf(stderr, "Error: upload-file failed: %s\n", uploadResp.Msg)
		return 1
	}
	fileURL := extractViewUploadedURL(uploadResp.Data)
	if fileURL == "" {
		fmt.Fprintf(stderr, "Error: upload-file response missing url\n")
		return 1
	}

	// 第二步：登记资源记录
	if name == "" {
		name = filepath.Base(file)
	}
	reqBody := map[string]any{
		"name": name,
		"url":  fileURL,
		"size": fi.Size(),
		"type": inferViewAssetType(filepath.Ext(file)),
	}
	if groupID != "" {
		reqBody["groupId"] = groupID
	}
	if tags != "" {
		// 后端 AssetUploadReq.tags 为逗号分隔字符串
		reqBody["tags"] = tags
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: viewAssetUploadPath,
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// extractViewUploadedURL 从 upload-file 响应中提取文件 URL；
// 兼容 data 为字符串或对象两种形态；对象兼容 url/fileUrl/fileUri/filePath/path 各字段名
func extractViewUploadedURL(data any) string {
	if s, ok := data.(string); ok {
		return s
	}
	if m, ok := data.(map[string]any); ok {
		for _, key := range []string{"url", "fileUrl", "fileUri", "filePath", "path"} {
			if s, ok := m[key].(string); ok && s != "" {
				return s
			}
		}
	}
	return ""
}

// 资源类型推断表：扩展名（小写、含点）→ 资源类型
var viewAssetTypeByExt = map[string]string{
	".jpg": "image", ".jpeg": "image", ".png": "image", ".gif": "image",
	".svg": "image", ".webp": "image", ".bmp": "image", ".ico": "image",
	".mp4": "video", ".webm": "video", ".mov": "video",
	".mp3": "audio", ".wav": "audio", ".ogg": "audio",
}

// inferViewAssetType 按文件扩展名推断大屏资源类型，未知扩展名归为 other
func inferViewAssetType(ext string) string {
	if t, ok := viewAssetTypeByExt[strings.ToLower(ext)]; ok {
		return t
	}
	return "other"
}

// runViewAssetDelete 执行删除大屏资源命令
func runViewAssetDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput := false
	assetID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				assetID = args[i+1]
				i++
			}
		case "--json", "-j":
			jsonOutput = true
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", args[i])
			return 2
		}
	}

	if assetID == "" {
		fmt.Fprintln(stderr, "--id is required")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: viewAssetDeletePath,
		Body: map[string]any{"id": assetID},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}
