package shared

import (
	"fmt"
	"io"
	"os"

	"gitee.com/unitedrhino/cli/internal/skillinstall"
	"gitee.com/unitedrhino/cli/internal/upgrade"
	"gitee.com/unitedrhino/cli/internal/version"
)

// 升级依赖通过变量保留测试替换点，生产始终使用真实实现。
var (
	checkUpgrade            = upgrade.Check
	performUpgrade          = upgrade.Perform
	deployEmbeddedSkillsNow = deployEmbeddedSkills
)

func runUpgrade(args []string, stdout, stderr io.Writer) int {
	opts := upgrade.Options{}
	showHelp := false
	jsonOutput := false
	installSkills := true

	// 解析参数
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dry-run", "--check":
			opts.DryRun = true
		case "--install-skills":
			// v0.6.2 起默认同步客户端 Skills；保留该参数兼容已有脚本。
			installSkills = true
		case "--no-skills":
			installSkills = false
		case "--force":
			opts.Force = true
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

	if opts.DryRun {
		if version.IsDev() && opts.TargetVersion == "" {
			if jsonOutput {
				result := &upgrade.Result{
					CurrentVersion: version.BuildVersion,
					ErrorMessage:   "当前为开发版本，无法自动检查更新。请使用 --version 指定目标版本。",
				}
				_ = writeJSON(stdout, result)
				return 1
			}
			fmt.Fprintln(stdout, "当前为开发版本 (dev)，无法自动升级。请使用 --version 指定目标版本。")
			fmt.Fprintln(stdout, "示例: ur upgrade --version v0.3.5")
			return 1
		}
		result, err := checkUpgrade()
		if err != nil {
			fmt.Fprintf(stderr, "检查更新失败: %v\n", err)
			return 1
		}
		if result.ErrorMessage != "" {
			if jsonOutput {
				_ = writeJSON(stdout, result)
			} else {
				fmt.Fprintf(stderr, "检查更新失败: %s\n", result.ErrorMessage)
			}
			return 1
		}
		if jsonOutput {
			_ = writeJSON(stdout, result)
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

	if !jsonOutput {
		fmt.Fprintf(stdout, "当前版本: %s\n", version.BuildVersion)
		fmt.Fprintln(stdout, "正在检查更新...")
	}

	result, err := performUpgrade(opts)
	if err != nil {
		if jsonOutput {
			_ = writeJSON(stdout, result)
		} else {
			fmt.Fprintf(stderr, "升级失败: %v\n", err)
			if result.ErrorMessage != "" {
				fmt.Fprintf(stderr, "%s\n", result.ErrorMessage)
			}
		}
		return 1
	}

	if installSkills {
		installResult, installErr := deployEmbeddedSkillsNow()
		if installResult != nil {
			for _, target := range installResult.Targets {
				if target.Installed && target.Error == "" {
					result.SkillsInstallTargets++
				}
			}
			result.SkillsInstalled = result.SkillsInstallTargets > 0 && !skillinstall.HasErrors(installResult)
		}
		if installErr != nil {
			result.SkillsInstallError = installErr.Error()
			if jsonOutput {
				_ = writeJSON(stdout, result)
			} else {
				fmt.Fprintf(stderr, "部署 Skills 到 AI 工具失败: %v\n", installErr)
			}
			return 1
		}
	}

	if jsonOutput {
		_ = writeJSON(stdout, result)
		return 0
	}
	if result.UpToDate {
		fmt.Fprintf(stdout, "已是最新版本 %s\n", result.LatestVersion)
	} else {
		fmt.Fprintf(stdout, "升级成功! %s → %s\n", result.CurrentVersion, result.LatestVersion)
	}
	backupDir := upgrade.BackupDir()
	if backupDir == "" {
		home, _ := os.UserHomeDir()
		if home != "" {
			backupDir = home + "/.ur/backup"
		}
	}
	if !result.UpToDate {
		fmt.Fprintf(stdout, "旧版本已备份到: %s\n", backupDir)
	}

	// 输出内置 skills 同步结果
	if result.SkillsSynced {
		fmt.Fprintf(stdout, "内置 skills 已随升级同步（%d 个目录）\n", result.SkillsUpdated)
	} else if result.SkillsMessage != "" {
		fmt.Fprintf(stderr, "警告: %s\n", result.SkillsMessage)
	}

	if installSkills {
		fmt.Fprintf(stdout, "客户端 Skills 已同步（%d 个目标）\n", result.SkillsInstallTargets)
	} else if result.SkillsSynced {
		fmt.Fprintln(stdout, "已按 --no-skills 跳过客户端 Skills 部署")
	}
	return 0
}

// deployEmbeddedSkills 把当前内置 Skills 部署到自动发现和用户登记的全部目标。
func deployEmbeddedSkills() (*skillinstall.Result, error) {
	src := upgrade.GetDefaultSkillsDir()
	targets, err := skillinstall.ResolveTargets(cwdOr("."))
	if err != nil {
		return nil, fmt.Errorf("探测 AI Skills 目录失败: %w", err)
	}
	if len(targets) == 0 {
		return &skillinstall.Result{Source: src}, nil
	}
	result, err := skillinstall.Install(src, targets, false)
	if err != nil {
		return result, err
	}
	if skillinstall.HasErrors(result) {
		return result, fmt.Errorf("部分 AI Skills 目标同步失败")
	}
	return result, nil
}

// installEmbeddedSkills 把当前内置 Skills 部署到自动发现和用户登记的全部目标。
// 即使 CLI 已是最新版也会执行，避免二进制已升级但客户端仍保留旧副本。
func installEmbeddedSkills(stdout, stderr io.Writer) int {
	installResult, err := deployEmbeddedSkills()
	if err != nil {
		fmt.Fprintf(stderr, "部署 Skills 到 AI 工具失败: %v\n", err)
		return 1
	}
	if len(installResult.Targets) == 0 {
		fmt.Fprintln(stdout, "未发现可部署目标；可用 ur skills target add 登记目录")
		return 0
	}
	fmt.Fprintln(stdout, skillinstall.Summary(installResult))
	if skillinstall.HasErrors(installResult) {
		return 1
	}
	return 0
}

func printUpgradeHelp(w io.Writer) {
	fmt.Fprintln(w, "用法: ur upgrade [选项]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "选项:")
	fmt.Fprintln(w, "  --check, --dry-run 只检查更新，不安装")
	fmt.Fprintln(w, "  --install-skills   兼容参数；升级默认同步全部客户端 Skills")
	fmt.Fprintln(w, "  --no-skills        只升级 CLI，不部署客户端 Skills")
	fmt.Fprintln(w, "  --force            即使已是最新版也重新安装")
	fmt.Fprintln(w, "  --version <tag>    升级到指定版本（如 --version v0.3.3）")
	fmt.Fprintln(w, "  --json             以 JSON 格式输出结果")
	fmt.Fprintln(w, "  -h, --help         显示帮助信息")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "示例:")
	fmt.Fprintln(w, "  ur upgrade                  升级到最新版本")
	fmt.Fprintln(w, "  ur upgrade --check          检查是否有新版本")
	fmt.Fprintln(w, "  ur upgrade                  升级并部署 Skills 到各 AI 工具")
	fmt.Fprintln(w, "  ur upgrade --version v0.3.3 降级到指定版本")
}
