package shared

import (
	"context"
	"fmt"
	"io"

	"gitee.com/unitedrhino/cli/internal/client"
)

// runDept 执行部门管理命令
func runDept(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printDeptHelp(stdout)
		return 0
	}

	switch args[0] {
	case "user":
		return runDeptUser(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printDeptHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown dept subcommand: %s\n", args[0])
		printDeptHelp(stderr)
		return 2
	}
}

// printDeptHelp 打印部门管理帮助信息
func printDeptHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur dept <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Department management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  user       Department user management (get-list, batch-create, batch-delete)")
	fmt.Fprintln(w, "  help       Show this help message")
}

// runDeptUser 执行部门用户管理命令
func runDeptUser(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printDeptUserHelp(stdout)
		return 0
	}

	switch args[0] {
	case "get-list":
		return runDeptUserGetList(ctx, args[1:], stdout, stderr)
	case "batch-create":
		return runDeptUserBatchCreate(ctx, args[1:], stdout, stderr)
	case "batch-delete":
		return runDeptUserBatchDelete(ctx, args[1:], stdout, stderr)
	case "help", "--help", "-h":
		printDeptUserHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown dept user subcommand: %s\n", args[0])
		printDeptUserHelp(stderr)
		return 2
	}
}

// printDeptUserHelp 打印部门用户帮助信息
func printDeptUserHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage: ur dept user <subcommand> [options]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Department user management")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Subcommands:")
	fmt.Fprintln(w, "  get-list       Query department user list")
	fmt.Fprintln(w, "  batch-create   Batch create department users")
	fmt.Fprintln(w, "  batch-delete   Batch delete department users")
}

// runDeptUserGetList 执行查询部门用户列表命令
func runDeptUserGetList(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	jsonOutput, page, size, remaining := parseInfoListParams(args)
	reqBody := map[string]any{
		"page": map[string]any{"page": page, "size": size},
	}

	for i := 0; i < len(remaining); i++ {
		switch remaining[i] {
		case "--dept-id":
			if i+1 < len(remaining) {
				reqBody["deptID"] = remaining[i+1]
				i++
			}
		}
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/dept/user/get-list",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runDeptUserBatchCreate 执行批量创建部门用户命令
func runDeptUserBatchCreate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		}
	}

	if bodyJSON == "" {
		fmt.Fprintln(stderr, "--body is required")
		return 2
	}

	reqBody, err := parseBodyArg(bodyJSON)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/dept/user/batch-create",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}

// runDeptUserBatchDelete 执行批量删除部门用户命令
func runDeptUserBatchDelete(ctx context.Context, args []string, stdout, stderr io.Writer) int {
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
		}
	}

	if bodyJSON == "" {
		fmt.Fprintln(stderr, "--body is required")
		return 2
	}

	reqBody, err := parseBodyArg(bodyJSON)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{
		Path: "/api/v1/system/dept/user/batch-delete",
		Body: reqBody,
	})
	if err != nil {
		fmt.Fprintf(stderr, "API error: %v\n", err)
		return 1
	}

	return outputResult(resp, jsonOutput, stdout, stderr)
}
