// 本文件验证设备领域命令的请求合同，确保 AI 可优先使用命令而不是通用 API。
package shared

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gitee.com/unitedrhino/cli/internal/config"
)

// TestDomainCommandHelpUsesNestedPaths 验证 AI 读取帮助时得到可直接执行的完整命令路径。
func TestDomainCommandHelpUsesNestedPaths(t *testing.T) {
	tests := []struct {
		name string
		run  func(io.Writer) int
		want string
	}{
		{
			name: "设备列表",
			run: func(stdout io.Writer) int {
				return runDevice(context.Background(), []string{"info", "get-list", "--help"}, stdout, io.Discard)
			},
			want: "Usage: ur things device info get-list",
		},
		{
			name: "产品物模型",
			run: func(stdout io.Writer) int {
				return runSchema(config.AppIoT, []string{"get-list", "--help"}, stdout, io.Discard)
			},
			want: "Usage: ur things schema get-list",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var stdout bytes.Buffer
			if exitCode := tc.run(&stdout); exitCode != 0 {
				t.Fatalf("exit=%d", exitCode)
			}
			if !strings.Contains(stdout.String(), tc.want) {
				t.Fatalf("help=%q, want %q", stdout.String(), tc.want)
			}
		})
	}
}

// TestDeviceInfoGetListProjectHeader 验证设备列表命令将字符串项目 ID 放入请求头。
func TestDeviceInfoGetListProjectHeader(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("project-id"); got != "9007199254740993" {
			t.Fatalf("project-id=%q", got)
		}
		_, _ = io.WriteString(w, `{"code":200,"data":{"list":[],"total":0}}`)
	}))
	defer server.Close()
	setDeviceCommandTestEnv(t, server.URL)

	exitCode := runDevice(context.Background(), []string{
		"info", "get-list", "--project-id", "9007199254740993", "--json",
	}, io.Discard, io.Discard)
	if exitCode != 0 {
		t.Fatalf("exit=%d", exitCode)
	}
}

// TestSchemaGetListProjectHeader 验证物模型列表命令将超大字符串项目 ID 放入请求头。
func TestSchemaGetListProjectHeader(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/things/product/schema/get-list" {
			t.Fatalf("path=%q", r.URL.Path)
		}
		if got := r.Header.Get("project-id"); got != "9223372036854775807" {
			t.Fatalf("project-id=%q", got)
		}
		_, _ = io.WriteString(w, `{"code":200,"data":{"list":[]}}`)
	}))
	defer server.Close()
	setDeviceCommandTestEnv(t, server.URL)

	exitCode := runSchema(config.AppIoT, []string{
		"get-list", "-p", "product-a",
		"--project-id", "9223372036854775807", "--json",
	}, io.Discard, io.Discard)
	if exitCode != 0 {
		t.Fatalf("exit=%d", exitCode)
	}
}

// TestSchemaGetListRejectsEmptyProjectID 验证显式空项目 ID 不会回退环境并发出歧义请求。
func TestSchemaGetListRejectsEmptyProjectID(t *testing.T) {
	requestSeen := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requestSeen = true
		_, _ = io.WriteString(w, `{"code":200,"data":{"list":[]}}`)
	}))
	defer server.Close()
	setDeviceCommandTestEnv(t, server.URL)
	t.Setenv("UR_PROJECT_ID", "12")

	exitCode := runSchema(config.AppIoT, []string{
		"get-list", "-p", "product-a", "--project-id=", "--json",
	}, io.Discard, io.Discard)
	if exitCode == 0 || requestSeen {
		t.Fatalf("exit=%d requestSeen=%v", exitCode, requestSeen)
	}
}

// TestDeviceControlCloudOnly 验证云端属性修改使用字符串 data、模式 4 与字符串项目头。
func TestDeviceControlCloudOnly(t *testing.T) {
	requestSeen := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestSeen = true
		if r.URL.Path != "/api/v1/things/device/interact/property-control-send" {
			t.Fatalf("path=%q", r.URL.Path)
		}
		if got := r.Header.Get("project-id"); got != "9007199254740993" {
			t.Fatalf("project-id=%q", got)
		}
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if body["data"] != `{"temperature":25.5}` {
			t.Fatalf("data=%#v", body["data"])
		}
		if body["shadowControl"] != float64(4) {
			t.Fatalf("shadowControl=%#v", body["shadowControl"])
		}
		_, _ = io.WriteString(w, `{"code":200,"data":{"code":200}}`)
	}))
	defer server.Close()
	setDeviceCommandTestEnv(t, server.URL)

	exitCode := runDeviceControl(context.Background(), []string{
		"-p", "product-a", "-d", "device-a",
		"--data", `{"temperature":25.5}`,
		"--shadow-control", "4",
		"--project-id", "9007199254740993",
		"--json",
	}, io.Discard, io.Discard)
	if exitCode != 0 || !requestSeen {
		t.Fatalf("exit=%d requestSeen=%v", exitCode, requestSeen)
	}
}

// TestDeviceControlRejectsInnerFailure 验证外层成功但控制业务失败时仍返回非零退出码。
func TestDeviceControlRejectsInnerFailure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, `{"code":200,"data":{"code":500,"msg":"failed"}}`)
	}))
	defer server.Close()
	setDeviceCommandTestEnv(t, server.URL)

	exitCode := runDeviceControl(context.Background(), []string{
		"-p", "product-a", "-d", "device-a",
		"--data", `{"temperature":25.5}`,
		"--shadow-control", "4",
		"--project-id", "123",
		"--json",
	}, io.Discard, io.Discard)
	if exitCode == 0 {
		t.Fatal("内层业务失败不应返回成功退出码")
	}
}

// TestDeviceSimulateReport 验证管理员模拟上报命令使用独立接口和字符串属性值。
func TestDeviceSimulateReport(t *testing.T) {
	requestSeen := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestSeen = true
		if r.URL.Path != "/api/v1/things/device/simulate/report" {
			t.Fatalf("path=%q", r.URL.Path)
		}
		if got := r.Header.Get("project-id"); got != "9223372036854775807" {
			t.Fatalf("project-id=%q", got)
		}
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		data, ok := body["data"].(map[string]any)
		if !ok || data["temperature"] != "25.5" {
			t.Fatalf("data=%#v", body["data"])
		}
		_, _ = io.WriteString(w, `{"code":200,"data":{}}`)
	}))
	defer server.Close()
	setDeviceCommandTestEnv(t, server.URL)

	exitCode := runDevice(context.Background(), []string{
		"simulate-report",
		"-p", "product-a", "-d", "device-a",
		"--data", `{"temperature":"25.5"}`,
		"--project-id", "9223372036854775807",
		"--json",
	}, io.Discard, io.Discard)
	if exitCode != 0 || !requestSeen {
		t.Fatalf("exit=%d requestSeen=%v", exitCode, requestSeen)
	}
}

// setDeviceCommandTestEnv 设置设备命令测试所需的最小认证上下文。
func setDeviceCommandTestEnv(t *testing.T, baseURL string) {
	t.Helper()
	t.Setenv("UR_BASE_URL", baseURL)
	t.Setenv("UR_APP_ID", "200")
	t.Setenv("UR_TENANT_CODE", "test")
	t.Setenv("UR_TOKEN", "test-token")
}
