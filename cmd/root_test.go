// root_test.go 验证根命令的自动提示启用范围。
package cmd

import "testing"

// TestIsReadOnlyCommand 校验信息命令跳过检查，带全局参数的业务命令仍会检查。
func TestIsReadOnlyCommand(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want bool
	}{
		{name: "无参数", args: nil, want: true},
		{name: "帮助", args: []string{"--help"}, want: true},
		{name: "升级", args: []string{"upgrade"}, want: true},
		{name: "业务命令", args: []string{"things", "device", "info", "get-list"}, want: false},
		{name: "全局应用参数", args: []string{"--app", "iot", "things", "device", "info", "get-list"}, want: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := isReadOnlyCommand(test.args); got != test.want {
				t.Fatalf("isReadOnlyCommand(%v)=%v, want %v", test.args, got, test.want)
			}
		})
	}
}
