// api_project_test.go 验证兼容 API 入口与 Cobra 入口使用同一项目上下文规则。
package shared

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestAPIProjectID 覆盖兼容入口参数、环境回退与冲突安全失败。
func TestAPIProjectID(t *testing.T) {
	// cases 的期望项目 ID 使用超过 JavaScript 安全整数的字符串。
	cases := []struct {
		name string
		args []string
		want string
		fail bool
	}{
		{"环境回退", nil, "9007199254740993", false},
		{"参数覆盖", []string{"--project-id", "9223372036854775807"}, "9223372036854775807", false},
		{"等号参数", []string{"--project-id=123"}, "123", false},
		{"参数头冲突", []string{"--project-id", "123", "--header", "Project-ID:124"}, "", true},
		{"空参数", []string{"--project-id="}, "", true},
		{"缺值", []string{"--project-id"}, "", true},
		{"下一选项不是值", []string{"--project-id", "--debug"}, "", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// received 在错误场景应保持为空，保证没有意外发出跨项目请求。
			received := make(chan string, 1)
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				received <- r.Header.Get("project-id")
				_, _ = io.WriteString(w, `{"code":200,"data":{}}`)
			}))
			defer server.Close()
			t.Setenv("UR_BASE_URL", server.URL)
			t.Setenv("UR_APP_ID", "200")
			t.Setenv("UR_TENANT_CODE", "test")
			t.Setenv("UR_TOKEN", "test-token")
			t.Setenv("UR_PROJECT_ID", "9007199254740993")
			code := runAPI(context.Background(), append([]string{"/api/v1/things/device/info/get-list"}, tc.args...), io.Discard, io.Discard)
			if (code != 0) != tc.fail {
				t.Fatalf("exit=%d, want failure=%v", code, tc.fail)
			}
			select {
			case got := <-received:
				if tc.fail || got != tc.want {
					t.Fatalf("project=%q, want=%q", got, tc.want)
				}
			default:
				if !tc.fail {
					t.Fatal("未发出预期请求")
				}
			}
		})
	}
}
