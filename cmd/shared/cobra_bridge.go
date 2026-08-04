// cobra_bridge.go — 为旧命令提供 Cobra 兼容的导出入口
package shared

import (
	"context"
	"io"

	"gitee.com/unitedrhino/cli/internal/config"
)

// CobraBridge 将旧的手动解析命令桥接到 Cobra 框架
type CobraBridge struct{}

// RunUser 桥接 user 命令
func (b CobraBridge) RunUser(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return runUser(ctx, args, stdout, stderr)
}

// RunTenant 桥接 tenant 命令
func (b CobraBridge) RunTenant(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return runTenant(ctx, args, stdout, stderr)
}

// RunDept 桥接 dept 命令
func (b CobraBridge) RunDept(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return runDept(ctx, args, stdout, stderr)
}

// RunAlarm 桥接 alarm 命令
func (b CobraBridge) RunAlarm(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return runAlarm(ctx, args, stdout, stderr)
}

// RunDevice 桥接 device 命令
func (b CobraBridge) RunDevice(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return runDevice(ctx, args, stdout, stderr)
}

// RunSchema 桥接 schema 命令（需要 app 上下文）
func (b CobraBridge) RunSchema(app config.CLIApp, args []string, stdout, stderr io.Writer) int {
	return runSchema(app, args, stdout, stderr)
}

// RunScene 桥接 scene 命令
func (b CobraBridge) RunScene(args []string, stdout, stderr io.Writer) int {
	return runScene(args, stdout, stderr)
}

// RunScript 桥接 script 命令
func (b CobraBridge) RunScript(args []string, stdout, stderr io.Writer) int {
	return runScript(args, stdout, stderr)
}

// RunModel 桥接 model 命令
func (b CobraBridge) RunModel(args []string, stdout, stderr io.Writer) int {
	return runModel(args, stdout, stderr)
}

// RunOta 桥接 ota 命令
func (b CobraBridge) RunOta(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return runOta(ctx, args, stdout, stderr)
}

// RunProject 桥接 project 命令
func (b CobraBridge) RunProject(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return runProject(ctx, args, stdout, stderr)
}

// RunArea 桥接 area 命令
func (b CobraBridge) RunArea(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return runArea(ctx, args, stdout, stderr)
}

// RunAgg 桥接 agg 命令
func (b CobraBridge) RunAgg(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return runAgg(ctx, args, stdout, stderr)
}

// RunAiTool 桥接 ai-tool 命令
func (b CobraBridge) RunAiTool(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return runAiTool(ctx, args, stdout, stderr)
}

// RunUpgrade 桥接 upgrade 命令
func (b CobraBridge) RunUpgrade(args []string, stdout, stderr io.Writer) int {
	return runUpgrade(args, stdout, stderr)
}

// RunGenerateSkills 桥接 generate-skills 命令（需要 app 上下文）
func (b CobraBridge) RunGenerateSkills(app config.CLIApp, args []string, stdout, stderr io.Writer) int {
	return runGenerateSkills(app, args, stdout, stderr)
}

// RunSkills 桥接 skills 命令
func (b CobraBridge) RunSkills(args []string, stdout, stderr io.Writer) int {
	return runSkills(args, stdout, stderr)
}

// RunSystem 桥接 system 命令
func (b CobraBridge) RunSystem(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return runSystem(ctx, args, stdout, stderr)
}

// RunViewScreen 桥接 view screen 命令
func (b CobraBridge) RunViewScreen(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return runViewScreen(ctx, args, stdout, stderr)
}

// RunViewAsset 桥接 view asset 命令
func (b CobraBridge) RunViewAsset(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return runViewAsset(ctx, args, stdout, stderr)
}
