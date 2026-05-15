package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"gitee.com/unitedrhino/cli/cmd/ur"
	"gitee.com/unitedrhino/cli/internal/version"
)

// version 由 release.sh 在构建时通过 -ldflags -X main.version=VERSION 注入
// commit 和 buildDate 同样通过 ldflags 注入
var versionVar = "dev"

func main() {
	// 初始化版本信息
	version.BuildVersion = versionVar

	// 检查是否需要输出结构化版本信息
	if argsHasVersion(os.Args[1:]) {
		if argsHasJSON(os.Args[1:]) || argsHasCheckLatest(os.Args[1:]) {
			// 获取二进制路径用于查找 skills 目录
			binaryPath, _ := os.Executable()
			// 如果无法获取可执行文件路径，尝试用当前工作目录
			if binaryPath == "" {
				binaryPath, _ = os.Getwd()
				if binaryPath != "" {
					binaryPath = filepath.Join(binaryPath, "ur")
				}
			}
			fmt.Println(version.FormatVersionJSON(binaryPath))
			os.Exit(0)
		}
		fmt.Println(version.BuildVersion)
		os.Exit(0)
	}

	os.Exit(ur.Execute(context.Background(), version.BuildVersion, os.Args[1:], os.Stdout, os.Stderr))
}

func argsHasVersion(args []string) bool {
	for _, arg := range args {
		if arg == "--version" || arg == "-v" {
			return true
		}
	}
	return false
}

func argsHasJSON(args []string) bool {
	for _, arg := range args {
		if arg == "--json" {
			return true
		}
	}
	return false
}

func argsHasCheckLatest(args []string) bool {
	for _, arg := range args {
		if arg == "--check-latest" {
			return true
		}
	}
	return false
}
