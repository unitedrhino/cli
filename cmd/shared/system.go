package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

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
	fmt.Fprintln(w, "  upload-file    Upload file to OSS")
	fmt.Fprintln(w, "  batch-agg      Batch aggregate API requests")
	fmt.Fprintln(w, "  help           Show this help message")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  # Upload file")
	fmt.Fprintln(w, "  ur system upload-file --file /path/to/file.jpg")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  # Batch aggregate requests")
	fmt.Fprintln(w, "  ur system batch-agg --apis '[{\"path\":\"/api/v1/system/user/self/get-one\",\"body\":{}}]'")
}

// runSystemUploadFile 执行文件上传命令
func runSystemUploadFile(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	filePath := ""
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
	if _, err := os.Stat(filePath); err != nil {
		fmt.Fprintf(stderr, "file not found: %v\n", err)
		return 2
	}

	// 读取文件内容
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(stderr, "read file error: %v\n", err)
		return 1
	}

	// 调用上传 API
	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/common/upload-file",
		Body: map[string]any{
			"fileName": filePath,
			"fileData": fileData,
		},
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
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
