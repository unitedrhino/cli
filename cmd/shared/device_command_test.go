// 本文件验证设备领域命令的请求合同，确保 AI 可优先使用命令而不是通用 API。
package shared

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

// TestDeviceMockProductDefaults 验证产品级 Mock 的默认参数、项目头与 AI 友好输出。
func TestDeviceMockProductDefaults(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		if r.URL.Path != "/api/v1/things/device/interact/schema-mock-gen" {
			t.Fatalf("path=%q", r.URL.Path)
		}
		if got := r.Header.Get("project-id"); got != "9007199254740993" {
			t.Fatalf("project-id=%q", got)
		}
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if body["productID"] != "2D" || body["type"] != float64(1) {
			t.Fatalf("body=%#v", body)
		}
		if _, exists := body["deviceName"]; exists {
			t.Fatalf("product scope must omit deviceName: %#v", body)
		}
		dataIDs, ok := body["dataIDs"].([]any)
		if !ok || len(dataIDs) != 0 {
			t.Fatalf("dataIDs=%#v", body["dataIDs"])
		}
		_, _ = io.WriteString(w, `{"code":200,"msg":"success","data":{"params":"{\"temperature\":21.5}"}}`)
	}))
	defer server.Close()
	setDeviceCommandTestEnv(t, server.URL)
	t.Setenv("UR_PROJECT_ID", "9007199254740993")

	var stdout bytes.Buffer
	exitCode := runDevice(context.Background(), []string{"mock", "-p", "2D", "-j"}, &stdout, io.Discard)
	if exitCode != 0 || requestCount != 1 {
		t.Fatalf("exit=%d requests=%d", exitCode, requestCount)
	}
	var output struct {
		Code int `json:"code"`
		Data struct {
			Scope     string           `json:"scope"`
			ProductID string           `json:"productID"`
			Type      string           `json:"type"`
			Count     int              `json:"count"`
			Items     []map[string]any `json:"items"`
		} `json:"data"`
	}
	if err := json.Unmarshal(stdout.Bytes(), &output); err != nil {
		t.Fatal(err)
	}
	if output.Code != 200 || output.Data.Scope != "product" || output.Data.ProductID != "2D" || output.Data.Type != "property" || output.Data.Count != 1 || output.Data.Items[0]["temperature"] != 21.5 {
		t.Fatalf("output=%s", stdout.String())
	}
}

// TestDeviceMockDeviceBatch 验证设备合并物模型、重复标识符、数字类型别名与多份聚合。
func TestDeviceMockDeviceBatch(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		if got := r.Header.Get("project-id"); got != "9223372036854775807" {
			t.Fatalf("project-id=%q", got)
		}
		var body struct {
			ProductID  string   `json:"productID"`
			DeviceName string   `json:"deviceName"`
			Type       int64    `json:"type"`
			DataIDs    []string `json:"dataIDs"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if body.ProductID != "2D" || body.DeviceName != "yanshi-dev01" || body.Type != 2 || strings.Join(body.DataIDs, ",") != "temperature,humidity" {
			t.Fatalf("body=%#v", body)
		}
		fmt.Fprintf(w, `{"code":200,"msg":"success","data":{"params":"{\"sequence\":%d}"}}`, requestCount)
	}))
	defer server.Close()
	setDeviceCommandTestEnv(t, server.URL)
	t.Setenv("UR_PROJECT_ID", "stale-project")

	var stdout bytes.Buffer
	exitCode := runDevice(context.Background(), []string{
		"mock", "-p", "2D", "-d", "yanshi-dev01", "-t", "2",
		"--data-id", "temperature", "--data-id=humidity", "-n", "2",
		"--project-id", "9223372036854775807", "-j",
	}, &stdout, io.Discard)
	if exitCode != 0 || requestCount != 2 {
		t.Fatalf("exit=%d requests=%d", exitCode, requestCount)
	}
	var output struct {
		Data struct {
			Scope      string           `json:"scope"`
			DeviceName string           `json:"deviceName"`
			Type       string           `json:"type"`
			Count      int              `json:"count"`
			Items      []map[string]any `json:"items"`
		} `json:"data"`
	}
	if err := json.Unmarshal(stdout.Bytes(), &output); err != nil {
		t.Fatal(err)
	}
	if output.Data.Scope != "device" || output.Data.DeviceName != "yanshi-dev01" || output.Data.Type != "event" || output.Data.Count != 2 || output.Data.Items[1]["sequence"] != float64(2) {
		t.Fatalf("output=%s", stdout.String())
	}
}

// TestParseDeviceMockTypeAliases 验证三种物模型类型及数字别名保持一致。
func TestParseDeviceMockTypeAliases(t *testing.T) {
	tests := []struct {
		input    string
		wantType int64
		wantName string
	}{
		{"property", 1, "property"}, {"1", 1, "property"},
		{"event", 2, "event"}, {"2", 2, "event"},
		{"action", 3, "action"}, {"3", 3, "action"},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			opts, err := parseDeviceMockOptions([]string{"-p", "p1", "--type", tc.input})
			if err != nil || opts.mockType != tc.wantType || opts.typeName != tc.wantName {
				t.Fatalf("opts=%#v err=%v", opts, err)
			}
		})
	}
}

// TestDeviceMockRejectsInvalidParameters 验证缺失上下文与非法参数不会发送请求。
func TestDeviceMockRejectsInvalidParameters(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requestCount++
		_, _ = io.WriteString(w, `{"code":200,"data":{"params":"{}"}}`)
	}))
	defer server.Close()
	tests := []struct {
		name      string
		args      []string
		projectID string
	}{
		{"缺少产品", []string{"mock", "-j"}, "1"},
		{"缺少项目", []string{"mock", "-p", "p1", "-j"}, ""},
		{"数量过小", []string{"mock", "-p", "p1", "-n", "0", "-j"}, "1"},
		{"数量过大", []string{"mock", "-p", "p1", "--num=101", "-j"}, "1"},
		{"数量非法", []string{"mock", "-p", "p1", "-n", "many", "-j"}, "1"},
		{"类型非法", []string{"mock", "-p", "p1", "-t", "service", "-j"}, "1"},
		{"标识符为空", []string{"mock", "-p", "p1", "--data-id=", "-j"}, "1"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			setDeviceCommandTestEnv(t, server.URL)
			t.Setenv("UR_PROJECT_ID", tc.projectID)
			if exitCode := runDevice(context.Background(), tc.args, io.Discard, io.Discard); exitCode != 2 {
				t.Fatalf("exit=%d", exitCode)
			}
		})
	}
	if requestCount != 0 {
		t.Fatalf("requests=%d", requestCount)
	}
}

// TestDeviceMockDoesNotOutputPartialResults 验证业务失败或非法 params 时不输出残缺聚合结果。
func TestDeviceMockDoesNotOutputPartialResults(t *testing.T) {
	tests := []struct {
		name   string
		second string
	}{
		{"业务失败", `{"code":500,"msg":"failed","data":{}}`},
		{"非法参数", `{"code":200,"msg":"success","data":{"params":"not-json"}}`},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			requestCount := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				requestCount++
				if requestCount == 1 {
					_, _ = io.WriteString(w, `{"code":200,"msg":"success","data":{"params":"{\"ok\":true}"}}`)
					return
				}
				_, _ = io.WriteString(w, tc.second)
			}))
			defer server.Close()
			setDeviceCommandTestEnv(t, server.URL)
			t.Setenv("UR_PROJECT_ID", "1")

			var stdout bytes.Buffer
			exitCode := runDevice(context.Background(), []string{"mock", "-p", "p1", "-n", "3", "-j"}, &stdout, io.Discard)
			if exitCode != 1 || requestCount != 2 || stdout.Len() != 0 {
				t.Fatalf("exit=%d requests=%d stdout=%q", exitCode, requestCount, stdout.String())
			}
		})
	}
}

// TestDeviceMockHelpUsesNestedPath 验证帮助文本提供完整领域命令而非通用 API。
func TestDeviceMockHelpUsesNestedPath(t *testing.T) {
	var stdout bytes.Buffer
	if exitCode := runDevice(context.Background(), []string{"mock", "--help"}, &stdout, io.Discard); exitCode != 0 {
		t.Fatalf("exit=%d", exitCode)
	}
	if !strings.Contains(stdout.String(), "Usage: ur things device mock") || strings.Contains(stdout.String(), "ur api ") {
		t.Fatalf("help=%q", stdout.String())
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
