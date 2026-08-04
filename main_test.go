// main_test.go — main 包版本参数判断单元测试
package main

import "testing"

// TestArgsHasVersion 校验仅顶层 --version/-v 命中，子命令参数不误判
func TestArgsHasVersion(t *testing.T) {
	cases := []struct {
		args []string
		want bool
	}{
		{nil, false},
		{[]string{"--version"}, true},
		{[]string{"-v"}, true},
		{[]string{"--check-latest"}, false},
		{[]string{"upgrade", "--version", "v0.3.5"}, false},
		{[]string{"view", "screen", "list"}, false},
		{[]string{"--version", "--json"}, true},
	}
	for _, tc := range cases {
		if got := argsHasVersion(tc.args); got != tc.want {
			t.Errorf("argsHasVersion(%v) = %v, want %v", tc.args, got, tc.want)
		}
	}
}
