// 文件说明：通过真实 Cobra 入口和隔离 HTTP 服务验证组件、图元查询及错误输出。
package view

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gitee.com/unitedrhino/cli/internal/client"
	"github.com/spf13/cobra"
)

// executeSearch 运行独立命令树，返回标准输出与执行错误。
func executeSearch(command *cobra.Command, args ...string) (string, error) {
	var output bytes.Buffer
	command.SetOut(&output)
	command.SetErr(&bytes.Buffer{})
	command.SilenceErrors = true
	command.SilenceUsage = true
	command.SetArgs(args)
	err := command.Execute()
	return output.String(), err
}

// TestComponentSearch 验证中英文、注册 key、空结果、完整目录与 IoT 支持信息。
func TestComponentSearch(t *testing.T) {
	t.Setenv("UR_BASE_URL", "http://127.0.0.1:1")
	for _, test := range []struct {
		keyword string
		wantKey string
	}{
		{"连线", "ConnectLine"}, {"connectline", "ConnectLine"}, {"VTextCommon", "TextCommon"},
		{"图片", "Image"}, {"没有这个组件xyz", ""},
	} {
		t.Run(test.keyword, func(t *testing.T) {
			output, err := executeSearch(newComponentCommand(), "search", "--keyword", test.keyword, "--json")
			if err != nil {
				t.Fatal(err)
			}
			var response struct {
				Data struct {
					List  []struct{ Key string }
					Total int
				}
			}
			if err := json.Unmarshal([]byte(output), &response); err != nil {
				t.Fatal(err)
			}
			if test.wantKey == "" {
				if response.Data.Total != 0 || response.Data.List == nil {
					t.Fatalf("空结果应为 []，得到 %s", output)
				}
				return
			}
			for _, component := range response.Data.List {
				if component.Key == test.wantKey {
					return
				}
			}
			t.Fatalf("缺少 %s: %s", test.wantKey, output)
		})
	}
	output, err := executeSearch(newComponentCommand(), "get-one", "--key", "TextCommon", "--json")
	if err != nil || !strings.Contains(output, `"property"`) {
		t.Fatalf("IoT 支持信息缺失: %v %s", err, output)
	}
	output, err = executeSearch(newComponentCommand(), "search")
	if err != nil || !strings.Contains(output, "ConnectLine") || !strings.Contains(output, "snapshotDate") {
		t.Fatalf("完整目录查询失败: %v %s", err, output)
	}
}

// TestSearchInvalidArguments 验证参数错误在认证或联网之前返回。
func TestSearchInvalidArguments(t *testing.T) {
	for _, args := range [][]string{
		{"search", "--page", "0"}, {"search", "--size", "101"}, {"search", "--size", "bad"},
		{"search", "--status", "unknown"}, {"search", "--format", "exe"},
		{"search", "--missing"}, {"search", "extra"}, {"get-one"}, {"get-one", "--id", " "},
	} {
		if _, err := executeSearch(newBuiltinImageCommand(), args...); err == nil {
			t.Fatalf("应拒绝参数 %v", args)
		}
	}
	for _, args := range [][]string{{"get-one"}, {"get-one", "--key", " "}, {"get-one", "--key", "unknown", "--json"}} {
		output, err := executeSearch(newComponentCommand(), args...)
		if err == nil {
			t.Fatalf("应拒绝参数 %v", args)
		}
		if strings.Contains(strings.Join(args, " "), "unknown") && !strings.Contains(output, `"code": 404`) {
			t.Fatalf("JSON 必须保留失败码: %s", output)
		}
	}
}

// TestBuiltinImageRequests 验证筛选、分页、应用上下文、详情与业务失败的完整链路。
func TestBuiltinImageRequests(t *testing.T) {
	var gotPath, gotApp, gotToken string
	var gotBody map[string]any
	response := client.APIResponse{Code: 200, Msg: "success", Data: map[string]any{
		"list": []any{map[string]any{"id": "bi-electric-0952", "name": "互感器", "url": "/app/iot/builtin-images/ct.svg", "width": 64, "height": 64}}, "total": 1,
	}}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath, gotApp, gotToken = r.URL.Path, r.Header.Get("app-id"), r.Header.Get("token")
		if r.Method != http.MethodPost {
			t.Error("图元接口必须使用 POST")
		}
		if err := json.NewDecoder(r.Body).Decode(&gotBody); err != nil {
			t.Error(err)
		}
		_ = json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()
	t.Setenv("UR_BASE_URL", server.URL)
	t.Setenv("UR_APP_ID", "1000")
	t.Setenv("UR_TENANT_CODE", "demo")
	t.Setenv("UR_TOKEN", "test-token")
	t.Setenv("UR_ACCOUNT", "")
	t.Setenv("UR_PASSWORD", "")
	t.Setenv("UR_ACCESS_KEY", "")
	t.Setenv("UR_ACCESS_SECRET", "")
	output, err := executeSearch(newBuiltinImageCommand(), "search", "--keyword", "互感器", "--category", "electric", "--sub-category", "电气", "--tags", "计量,电流", "--format", "svg", "--status", "confirmed", "--page", "2", "--size", "10", "--json")
	if err != nil {
		t.Fatal(err)
	}
	if gotPath != "/api/v1/view/builtin-image/get-list" || gotApp != "200" || gotToken != "test-token" {
		t.Fatalf("请求上下文不符: %s %s %s", gotPath, gotApp, gotToken)
	}
	if gotBody["keyword"] != "互感器" || gotBody["category"] != "electric" || gotBody["subCategory"] != "电气" || gotBody["status"] != "confirmed" || gotBody["format"] != "svg" {
		t.Fatalf("筛选丢失: %+v", gotBody)
	}
	page := gotBody["page"].(map[string]any)
	if page["page"] != float64(2) || page["size"] != float64(10) || len(gotBody["tags"].([]any)) != 2 || !strings.Contains(output, "bi-electric-0952") {
		t.Fatalf("分页、标签或结果不符: %+v %s", gotBody, output)
	}
	response.Data = map[string]any{"id": "bi-electric-0952", "description": "互感器", "url": "/ct.svg", "width": 64, "height": 64}
	output, err = executeSearch(newBuiltinImageCommand(), "get-one", "--id", "bi-electric-0952", "--json")
	if err != nil || gotPath != "/api/v1/view/builtin-image/get-one" || gotBody["id"] != "bi-electric-0952" || !strings.Contains(output, "width") {
		t.Fatalf("详情查询不符: %v %s %+v", err, gotPath, gotBody)
	}
	response.Data = map[string]any{"list": []any{}, "total": 0}
	output, err = executeSearch(newBuiltinImageCommand(), "search", "--json")
	if err != nil || !strings.Contains(output, `"list": []`) {
		t.Fatalf("空查询失败: %v %s", err, output)
	}
	for _, code := range []int{401, 403, 404} {
		response = client.APIResponse{Code: code, Msg: "查询失败"}
		output, err = executeSearch(newBuiltinImageCommand(), "get-one", "--id", "unknown", "--json")
		var actual client.APIResponse
		if decodeErr := json.Unmarshal([]byte(output), &actual); err == nil || decodeErr != nil || actual.Code != code {
			t.Fatalf("业务错误必须保留 JSON 且失败: %v %s", err, output)
		}
	}
}
