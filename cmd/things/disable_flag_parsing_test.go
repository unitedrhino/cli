// disable_flag_parsing_test.go — things 组命令 DisableFlagParsing 防回归测试
//
// 背景：agg/area/device/model/ota/project/scene/script 八个命令走 cmd/shared
// 的手工参数解析（wrapOldCommand 桥接旧命令实现），自身不注册任何 cobra flag。
// 若 cobra.Command 缺少 DisableFlagParsing: true，cobra 会把 -p/-d/--data-id 等
// 参数当未注册 flag 拦截（unknown shorthand flag），导致这些命令完全不可用
// （参考 cmd/alarm.go 中的同款做法注释）。
package things

import (
	"reflect"
	"testing"

	"github.com/spf13/cobra"
)

// thingsManualParseCmds 依赖手工参数解析、必须设置 DisableFlagParsing 的八个命令名
var thingsManualParseCmds = []string{"agg", "area", "device", "model", "ota", "project", "scene", "script"}

// TestThingsCmdsDisableFlagParsing 定义级断言：ThingsCmd 下八个命令均设置了
// DisableFlagParsing: true，防止后续新增/重构时遗漏导致命令被 cobra 拦截
func TestThingsCmdsDisableFlagParsing(t *testing.T) {
	byName := map[string]*cobra.Command{}
	for _, sub := range ThingsCmd.Commands() {
		byName[sub.Name()] = sub
	}
	for _, name := range thingsManualParseCmds {
		t.Run(name, func(t *testing.T) {
			sub, ok := byName[name]
			if !ok {
				t.Fatalf("ThingsCmd 下缺少子命令 %q", name)
			}
			if !sub.DisableFlagParsing {
				t.Errorf("命令 %q 缺少 DisableFlagParsing: true，cobra 会把 -p/-d 等参数当未知 flag 拦截", name)
			}
		})
	}
}

// TestDeviceCmdUnknownFlagsReachRunE 行为级断言：device 命令收到未注册的
// -p/-d 等参数时，cobra 不应以 unknown flag 报错拦截，而应把原始参数原样
// 交给 RunE（手工参数解析入口）。
// 只验证 flag 解析层面：临时把 RunE 替换为记录型 stub，不触发真实网络请求。
func TestDeviceCmdUnknownFlagsReachRunE(t *testing.T) {
	origRunE := deviceCmd.RunE
	var gotArgs []string
	deviceCmd.RunE = func(cmd *cobra.Command, args []string) error {
		gotArgs = args
		return nil
	}
	// 还原被替换的 RunE，避免污染同包其他测试
	defer func() { deviceCmd.RunE = origRunE }()

	// 重置 SetArgs，避免测试结束后残留参数影响后续 Execute 调用
	defer ThingsCmd.SetArgs(nil)

	// Execute 过程中 cobra 会向根命令自动注入 help/completion 内置命令，测试后移除，
	// 保持 ThingsCmd 与测试前一致
	defer func() {
		var injected []*cobra.Command
		for _, sub := range ThingsCmd.Commands() {
			if sub.Name() == "help" || sub.Name() == "completion" {
				injected = append(injected, sub)
			}
		}
		ThingsCmd.RemoveCommand(injected...)
	}()

	cases := []struct {
		name string
		args []string
		want []string
	}{
		{
			// 修复前的真实故障场景：ur things device log property -p xxx
			name: "子命令后的短 flag 不被拦截",
			args: []string{"device", "log", "property", "-p", "p_smartswitch_001"},
			want: []string{"log", "property", "-p", "p_smartswitch_001"},
		},
		{
			name: "短 flag 组合不被拦截",
			args: []string{"device", "-p", "xxx", "-d", "yyy"},
			want: []string{"-p", "xxx", "-d", "yyy"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotArgs = nil
			// 以 ThingsCmd 为根执行，模拟真实入口 ur things device ... 的完整路由路径
			ThingsCmd.SetArgs(tc.args)
			if err := ThingsCmd.Execute(); err != nil {
				t.Fatalf("执行 %v 不应报错（报 unknown flag 即回归）: %v", tc.args, err)
			}
			if !reflect.DeepEqual(gotArgs, tc.want) {
				t.Fatalf("RunE 收到参数 %v, 期望原始参数 %v", gotArgs, tc.want)
			}
		})
	}
}
