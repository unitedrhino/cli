// api_project_test.go 验证真实命令解析后的项目请求头、精度与冲突拒绝。
package cmd

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/pflag"
)

// TestAPIProjectID 使用本地 HTTP 服务验证项目参数，不访问业务环境。
func TestAPIProjectID(t *testing.T) {
	// cases 覆盖显式参数、环境回退、大小写头冲突以及缺省行为。
	cases := []struct {
		name string
		args []string
		env  string
		want string
		fail bool
	}{
		{"超大字符串", []string{"--project-id", "9223372036854775807"}, "", "9223372036854775807", false},
		{"环境回退", nil, "9007199254740993", "9007199254740993", false},
		{"参数覆盖环境", []string{"--project-id=9007199254740993"}, "12", "9007199254740993", false},
		{"请求头覆盖环境", []string{"--header", "Project-ID:9007199254740993"}, "12", "9007199254740993", false},
		{"相同头允许", []string{"--project-id", "12", "--header", "Project-ID:12"}, "13", "12", false},
		{"参数头冲突", []string{"--project-id", "12", "--header", "Project-ID:13"}, "", "", true},
		{"大小写头冲突", []string{"--header", "project-id:12", "--header", "Project-ID:13"}, "", "", true},
		{"重复头冲突", []string{"--header", "project-id:12", "--header", "project-id:13"}, "", "", true},
		{"空参数拒绝", []string{"--project-id="}, "12", "", true},
		{"缺参数拒绝", []string{"--project-id"}, "", "", true},
		{"空白参数拒绝", []string{"--project-id", " 12 "}, "", "", true},
		{"未指定不注入", nil, "", "", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// received 捕获 HTTP 请求，错误场景必须在发请求前停止。
			received := make(chan string, 1)
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				received <- r.Header.Get("project-id")
				_, _ = io.WriteString(w, `{"code":200,"msg":"success","data":{}}`)
			}))
			defer server.Close()
			t.Setenv("UR_BASE_URL", server.URL)
			t.Setenv("UR_APP_ID", "200")
			t.Setenv("UR_TENANT_CODE", "test")
			t.Setenv("UR_TOKEN", "test-token")
			t.Setenv("UR_PROJECT_ID", tc.env)
			// 还原全局 Cobra flag，避免不同用例共享参数状态。
			apiCmd.Flags().VisitAll(func(f *pflag.Flag) {
				if value, ok := f.Value.(pflag.SliceValue); ok {
					_ = value.Replace(nil)
				} else {
					_ = f.Value.Set(f.DefValue)
				}
				f.Changed = false
			})
			RootCmd.SetOut(io.Discard)
			RootCmd.SetErr(io.Discard)
			RootCmd.SetArgs(append([]string{"api", "/api/v1/things/device/info/get-list"}, tc.args...))
			err := RootCmd.Execute()
			if (err != nil) != tc.fail {
				t.Fatalf("error=%v, want failure=%v", err, tc.fail)
			}
			select {
			case got := <-received:
				if tc.fail || got != tc.want {
					t.Fatalf("project=%q, want=%q, failure=%v", got, tc.want, tc.fail)
				}
			default:
				if !tc.fail {
					t.Fatal("未发送预期请求")
				}
			}
		})
	}
}
