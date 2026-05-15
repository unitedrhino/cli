package shared

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"gitee.com/unitedrhino/cli/internal/upgrade"
	"gitee.com/unitedrhino/cli/internal/version"
)

func runUpgrade(args []string, stdout, stderr io.Writer) int {
	opts := upgrade.Options{}
	showHelp := false
	jsonOutput := false

	// 解析参数
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dry-run":
			opts.DryRun = true
		case "--version":
			if i+1 < len(args) {
				opts.TargetVersion = args[i+1]
				i++
			} else {
				fmt.Fprintln(stderr, "错误: --version 需要指定版本号")
				return 2
			}
		case "--json":
			jsonOutput = true
		case "-h", "--help":
			showHelp = true
		default:
			if args[i] != "" {
				fmt.Fprintf(stderr, "未知参数: %s\n", args[i])
				showHelp = true
			}
		}
	}

	if showHelp {
		printUpgradeHelp(stdout)
		return 0
	}

	if opts.DryRun || jsonOutput {
		if version.IsDev() && opts.TargetVersion == "" {
			if jsonOutput {
				result := &upgrade.Result{
					CurrentVersion: version.BuildVersion,
					ErrorMessage:   "当前为开发版本，无法自动检查更新。请使用 --version 指定目标版本。",
				}
				output, _ := json.MarshalIndent(result, "", "  ")
				fmt.Fprintln(stdout, string(output))
				return 1
			}
			fmt.Fprintln(stdout, "当前为开发版本 (dev)，无法自动升级。请使用 --version 指定目标版本。")
			fmt.Fprintln(stdout, "示例: ur upgrade --version v0.3.5")
			return 1
		}
		result, err := upgrade.Check()
		if err != nil {
			fmt.Fprintf(stderr, "检查更新失败: %v\n", err)
			return 1
		}
		if jsonOutput {
			output, _ := json.MarshalIndent(result, "", "  ")
			fmt.Fprintln(stdout, string(output))
			return 0
		}
		if result.UpToDate {
			fmt.Fprintf(stdout, "已是最新版本 %s\n", result.CurrentVersion)
		} else {
			fmt.Fprintf(stdout, "发现新版本 %s（当前 %s）\n", result.LatestVersion, result.CurrentVersion)
			fmt.Fprintf(stdout, "下载地址: %s\n", result.DownloadURL)
			fmt.Fprintln(stdout, "运行 ur upgrade 执行升级")
		}
		return 0
	}

	// 执行升级
	if version.IsDev() && opts.TargetVersion == "" {
		fmt.Fprintln(stdout, "当前为开发版本 (dev)，无法自动升级。请使用 --version 指定目标版本。")
		fmt.Fprintln(stdout, "示例: ur upgrade --version v0.3.5")
		return 1
	}

	fmt.Fprintf(stdout, "当前版本: %s\n", version.BuildVersion)
	fmt.Fprintln(stdout, "正在检查更新...")

	result, err := upgrade.Perform(opts)
	if jsonOutput {
		output, _ := json.MarshalIndent(result, "", "  ")
		fmt.Fprintln(stdout, string(output))
		return 0
	}

	if err != nil {
		if result.UpToDate {
			fmt.Fprintf(stdout, "已是最新版本 %s\n", result.LatestVersion)
			return 0
		}
		fmt.Fprintf(stderr, "升级失败: %v\n", err)
		if result.ErrorMessage != "" {
			fmt.Fprintf(stderr, "%s\n", result.ErrorMessage)
		}
		return 1
	}

	if result.UpToDate {
		fmt.Fprintf(stdout, "已是最新版本 %s\n", result.LatestVersion)
		return 0
	}

	fmt.Fprintf(stdout, "升级成功! %s → %s\n", result.CurrentVersion, result.LatestVersion)
	backupDir := upgrade.BackupDir()
	if backupDir == "" {
		home, _ := os.UserHomeDir()
		if home != "" {
			backupDir = home + "/.ur/backup"
		}
	}
	fmt.Fprintf(stdout, "旧版本已备份到: %s\n", backupDir)
	return 0
}

func printUpgradeHelp(w io.Writer) {
	fmt.Fprintln(w, "用法: ur upgrade [选项]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "选项:")
	fmt.Fprintln(w, "  --dry-run          只检查更新，不安装")
	fmt.Fprintln(w, "  --version <tag>    升级到指定版本（如 --version v0.3.3）")
	fmt.Fprintln(w, "  --json             以 JSON 格式输出结果")
	fmt.Fprintln(w, "  -h, --help         显示帮助信息")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "示例:")
	fmt.Fprintln(w, "  ur upgrade                  升级到最新版本")
	fmt.Fprintln(w, "  ur upgrade --dry-run        检查是否有新版本")
	fmt.Fprintln(w, "  ur upgrade --version v0.3.3 降级到指定版本")
}
