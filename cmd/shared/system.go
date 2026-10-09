package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gitee.com/unitedrhino/cli/internal/client"
)

// runSystem 执行系统管理命令
func runSystem(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printSystemHelp(stdout)
		return 0
	}

	switch args[0] {
	case "upload-file":
		return runSystemUploadFile(ctx, args[1:], stdout, stderr)
	case "batch-agg":
		return runSystemBatchAgg(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printSystemHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown system subcommand: %s\n", args[0])
		printSystemHelp(stderr)
		return 2
	}
}

// printSystemHelp 打印系统管理帮助信息
func printSystemHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur system <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "System management commands")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  upload-file    Upload file to OSS (temporary bucket by default)")
	fmt.Fprintln(w, "  batch-agg      Batch aggregate API requests")
	fmt.Fprintln(w, "  help           Show this help message")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Upload file to temporary bucket, print fileUri and ready-to-use markdown link")
	fmt.Fprintln(w, "  ur system upload-file --file /path/to/report.xlsx")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Upload with custom remote file name")
	fmt.Fprintln(w, "  ur system upload-file --file /tmp/data.csv --name 统计数据.csv")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Upload to public bucket (business/useBy required by server)")
	fmt.Fprintln(w, "  ur system upload-file --file /path/to/logo.png --public --business view --use-by user")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Batch aggregate requests")
	fmt.Fprintln(w, "  ur system batch-agg --apis '[{\"path\":\"/api/v1/system/user/self/get-one\",\"body\":{}}]'")
}

// runSystemUploadFile 执行文件上传命令：内置 multipart 直传 /api/v1/system/common/upload-file，
// 默认上传到 OSS 临时桶并输出 fileUri（页面经 file-proxy 登录态下载），供 AI 沙箱交付生成文件使用。
func runSystemUploadFile(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	filePath := ""
	remoteName := ""
	isPublic := false
	business := ""
	scene := ""
	useBy := ""
	jsonOutput := false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--file", "-f":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--file requires value")
				return 2
			}
			filePath = args[i+1]
			i++
		case "--name", "-n":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--name requires value")
				return 2
			}
			remoteName = args[i+1]
			i++
		case "--public":
			isPublic = true
		case "--business":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--business requires value")
				return 2
			}
			business = args[i+1]
			i++
		case "--scene":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--scene requires value")
				return 2
			}
			scene = args[i+1]
			i++
		case "--use-by":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--use-by requires value")
				return 2
			}
			useBy = args[i+1]
			i++
		case "--json", "-j":
			jsonOutput = true
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", args[i])
			return 2
		}
	}

	if filePath == "" {
		fmt.Fprintln(stderr, "--file is required")
		return 2
	}

	// 检查文件是否存在
	fi, err := os.Stat(filePath)
	if err != nil {
		fmt.Fprintf(stderr, "file not found: %v\n", err)
		return 2
	}
	if fi.IsDir() {
		fmt.Fprintf(stderr, "path is a directory, not a file: %s\n", filePath)
		return 2
	}

	// 读取文件内容
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(stderr, "read file error: %v\n", err)
		return 1
	}
	if remoteName == "" {
		remoteName = filepath.Base(filePath)
	}

	// 组装表单：默认临时桶；公开桶时 business/useBy 为服务端必填项
	form := map[string]string{}
	if isPublic {
		if business == "" {
			fmt.Fprintln(stderr, "--business is required when --public is set")
			return 2
		}
		if useBy == "" {
			useBy = "user"
		}
		form["isPublic"] = "true"
		form["business"] = business
		form["scene"] = scene
		form["useBy"] = useBy
	}

	// multipart 直传（复用客户端统一鉴权与 401 重试）
	resp, err := client.UploadFileMultipart(ctx, systemCommonUploadFilePath, "file", remoteName, fileData, form)
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}
	if resp.Code != 200 {
		fmt.Fprintf(stderr, "Error: upload-file failed: %s\n", resp.Msg)
		return 1
	}

	// 人类可读输出 fileUri（供拼 markdown 链接给页面）；--json 输出完整响应
	if jsonOutput {
		return outputResult(resp, true, stdout, stderr)
	}
	fileUri, filePathOut := extractUploadFields(resp.Data)
	if fileUri == "" {
		fmt.Fprintf(stderr, "Error: upload-file response missing fileUri\n")
		return 1
	}
	fmt.Fprintf(stdout, "fileUri:  %s\n", fileUri)
	if filePathOut != "" {
		fmt.Fprintf(stdout, "filePath: %s\n", filePathOut)
	}
	fmt.Fprintln(stdout, "markdown: ["+remoteName+"]("+fileUri+")")
	return 0
}

// extractUploadFields 从上传接口响应 data 中提取 fileUri 与 filePath 字段。
func extractUploadFields(data any) (fileUri, filePath string) {
	m, ok := data.(map[string]any)
	if !ok {
		return "", ""
	}
	if v, ok := m["fileUri"].(string); ok {
		fileUri = v
	}
	if v, ok := m["filePath"].(string); ok {
		filePath = v
	}
	return fileUri, filePath
}

// runSystemBatchAgg 执行批量聚合接口命令
func runSystemBatchAgg(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	apis := ""
	jsonOutput := false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--apis":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--apis requires value")
				return 2
			}
			apis = args[i+1]
			i++
		case "--json", "-j":
			jsonOutput = true
		default:
			fmt.Fprintf(stderr, "unknown option: %s\n", args[i])
			return 2
		}
	}

	if apis == "" {
		fmt.Fprintln(stderr, "--apis is required")
		return 2
	}

	var apiList []map[string]any
	if err := json.Unmarshal([]byte(apis), &apiList); err != nil {
		fmt.Fprintf(stderr, "Error: --apis must be JSON array: %v\n", err)
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/common/api/batch-agg",
		Body: map[string]any{
			"apis": apiList,
		},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}
