package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"gitee.com/unitedrhino/cli/internal/client"
)

func runAiToolArtifact(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "用法: ai-tool artifact get --id <id> [--output-dir <dir>]")
		fmt.Fprintln(stderr, "      ai-tool artifact save --id <id> --dir <dir>")
		return 2
	}

	switch args[0] {
	case "get":
		return runAiToolArtifactGet(ctx, args[1:], stdout, stderr)
	case "save":
		return runAiToolArtifactSave(ctx, args[1:], stdout, stderr)
	default:
		fmt.Fprintf(stderr, "未知的 artifact 子命令: %s\n", args[0])
		return 2
	}
}

func runAiToolArtifactGet(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	var id int64
	outputDir := ""

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--id 需要参数")
				return 2
			}
			var err error
			id, err = strconv.ParseInt(args[i+1], 10, 64)
			if err != nil {
				fmt.Fprintf(stderr, "--id 格式无效: %v\n", err)
				return 2
			}
			i++
		case "--output-dir":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--output-dir 需要参数")
				return 2
			}
			outputDir = args[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "未知选项: %s\n", args[i])
			return 2
		}
	}

	if id == 0 {
		fmt.Fprintln(stderr, "必须提供 --id")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/ai/tool/get-one",
		Body: map[string]any{"id": strconv.FormatInt(id, 10)},
	})
	if err != nil {
		fmt.Fprintf(stderr, "请求失败: %v\n", err)
		return 1
	}
	if resp.Code != 200 {
		fmt.Fprintf(stderr, "API 错误 code=%d: %s\n", resp.Code, resp.Msg)
		return 1
	}

	dataMap, ok := resp.Data.(map[string]any)
	if !ok {
		fmt.Fprintln(stderr, "响应 data 格式异常")
		return 1
	}

	artifact, ok := dataMap["artifact"].(map[string]any)
	if !ok {
		fmt.Fprintln(stderr, "响应中没有 artifact")
		return 1
	}

	executorJs, _ := artifact["executorJs"].(string)
	skillMd, _ := artifact["skillMd"].(string)
	manifestJson, _ := artifact["manifestJson"].(string)

	if outputDir != "" {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			fmt.Fprintf(stderr, "创建目录失败: %v\n", err)
			return 1
		}
		writeFile(filepath.Join(outputDir, "executor.js"), executorJs, stderr)
		writeFile(filepath.Join(outputDir, "skill.md"), skillMd, stderr)
		writeFile(filepath.Join(outputDir, "manifest.json"), manifestJson, stderr)
		fmt.Fprintf(stdout, "已保存到: %s\n", outputDir)
	} else {
		out := map[string]string{
			"executorJs":   executorJs,
			"skillMd":   skillMd,
			"manifestJson": manifestJson,
		}
		raw, _ := json.MarshalIndent(out, "", "  ")
		fmt.Fprintln(stdout, string(raw))
	}
	return 0
}

func runAiToolArtifactSave(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	var id int64
	dir := ""

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--id 需要参数")
				return 2
			}
			var err error
			id, err = strconv.ParseInt(args[i+1], 10, 64)
			if err != nil {
				fmt.Fprintf(stderr, "--id 格式无效: %v\n", err)
				return 2
			}
			i++
		case "--dir":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--dir 需要参数")
				return 2
			}
			dir = args[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "未知选项: %s\n", args[i])
			return 2
		}
	}

	if id == 0 {
		fmt.Fprintln(stderr, "必须提供 --id")
		return 2
	}
	if dir == "" {
		fmt.Fprintln(stderr, "必须提供 --dir")
		return 2
	}

	executorJs, err := os.ReadFile(filepath.Join(dir, "executor.js"))
	if err != nil {
		fmt.Fprintf(stderr, "读取 executor.js 失败: %v\n", err)
		return 1
	}
	skillMd, _ := os.ReadFile(filepath.Join(dir, "skill.md"))
	manifestJson, _ := os.ReadFile(filepath.Join(dir, "manifest.json"))

	body := map[string]any{
		"id":         strconv.FormatInt(id, 10),
		"executorJs": string(executorJs),
	}
	if len(skillMd) > 0 {
		body["skillMd"] = string(skillMd)
	}
	if len(manifestJson) > 0 {
		body["manifestJson"] = string(manifestJson)
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/ai/tool/save-artifact",
		Body: body,
	})
	if err != nil {
		fmt.Fprintf(stderr, "请求失败: %v\n", err)
		return 1
	}
	if resp.Code != 200 {
		fmt.Fprintf(stderr, "API 错误 code=%d: %s\n", resp.Code, resp.Msg)
		return 1
	}

	fmt.Fprintln(stdout, "保存成功")
	return 0
}

func writeFile(path string, content string, stderr io.Writer) {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		fmt.Fprintf(stderr, "写入 %s 失败: %v\n", path, err)
	}
}
