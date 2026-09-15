// flow_test.go — 流程审批中心命令测试：参数解析、请求路径与请求体契约。
// 通过 httptest 模拟后端，环境变量注入认证（与 client_test.go 同口径）。
package shared

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// newFlowTestServer 启动模拟后端并注入环境变量，捕获最近一次请求供断言。
type flowCapturedRequest struct {
	path string
	body map[string]any
}

func newFlowTestServer(t *testing.T) *flowCapturedRequest {
	t.Helper()
	captured := &flowCapturedRequest{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured.path = r.URL.Path
		captured.body = map[string]any{}
		defer r.Body.Close()
		_ = json.NewDecoder(r.Body).Decode(&captured.body)
		_, _ = w.Write([]byte(`{"code":200,"msg":"success","data":{"id":"1"}}`))
	}))
	t.Cleanup(server.Close)

	t.Setenv("UR_BASE_URL", server.URL)
	t.Setenv("UR_APP_ID", "300")
	t.Setenv("UR_TENANT_CODE", "platform")
	t.Setenv("UR_TOKEN", "runtime-token")
	return captured
}

// argsToStdio 忽略输出的辅助包装
func argsToStdio() (*os.File, *os.File) {
	return os.Stdin, os.Stdout
}

func TestRunFlowDefCreate_BuildsContractBody(t *testing.T) {
	captured := newFlowTestServer(t)
	model := `{"nodeConfig":{"nodeKey":"n1","type":"major"}}`
	code := runFlow(context.Background(), []string{
		"def", "create",
		"--code", "leave", "--name", "请假",
		"--model", model,
		"--category", "hr", "--form-id", "9",
	}, io.Discard, io.Discard)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	if captured.path != "/api/v1/system/flow/def/create" {
		t.Fatalf("path = %q", captured.path)
	}
	if captured.body["processCode"] != "leave" || captured.body["name"] != "请假" {
		t.Fatalf("body = %+v", captured.body)
	}
	// modelContent 契约为 JSON 字符串（而非对象）
	if captured.body["modelContent"] != model {
		t.Fatalf("modelContent = %v", captured.body["modelContent"])
	}
	if captured.body["formTemplateId"] != "9" {
		t.Fatalf("formTemplateId = %v（后端契约 string）", captured.body["formTemplateId"])
	}
}

func TestRunFlowDefCreate_ModelFromFile(t *testing.T) {
	captured := newFlowTestServer(t)
	dir := t.TempDir()
	file := filepath.Join(dir, "model.json")
	model := `{"nodeConfig":{"nodeKey":"start"}}`
	if err := os.WriteFile(file, []byte(model), 0o600); err != nil {
		t.Fatal(err)
	}
	code := runFlow(context.Background(), []string{
		"def", "create", "--code", "a", "--name", "b", "--model", "@" + file,
	}, io.Discard, io.Discard)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	if captured.body["modelContent"] != model {
		t.Fatalf("modelContent from file = %v", captured.body["modelContent"])
	}
}

func TestRunFlowTaskConsent_TaskIdAsString(t *testing.T) {
	captured := newFlowTestServer(t)
	code := runFlow(context.Background(), []string{
		"task", "consent", "--task-id", "123", "--comment", "同意",
	}, io.Discard, io.Discard)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	if captured.path != "/api/v1/system/flow/task/consent" {
		t.Fatalf("path = %q", captured.path)
	}
	// taskId 契约为字符串（int64,string 标签）
	if captured.body["taskId"] != "123" {
		t.Fatalf("taskId = %v", captured.body["taskId"])
	}
}

func TestRunFlowTaskReject_StrategyRequired(t *testing.T) {
	captured := newFlowTestServer(t)
	code := runFlow(context.Background(), []string{
		"task", "reject", "--task-id", "5", "--strategy", "toInitiator", "--comment", "退回发起人",
	}, io.Discard, io.Discard)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	if captured.body["strategy"] != "toInitiator" {
		t.Fatalf("strategy = %v", captured.body["strategy"])
	}
}

func TestRunFlowInstanceTerminate_InstanceIDAsString(t *testing.T) {
	captured := newFlowTestServer(t)
	code := runFlow(context.Background(), []string{
		"instance", "terminate", "--instance-id", "77", "--comment", "违规流程",
	}, io.Discard, io.Discard)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	if captured.path != "/api/v1/system/flow/instance/terminate" {
		t.Fatalf("path = %q", captured.path)
	}
	if captured.body["instanceId"] != "77" {
		t.Fatalf("instanceId = %v", captured.body["instanceId"])
	}
}

func TestRunFlowLaunch_VariableParsed(t *testing.T) {
	captured := newFlowTestServer(t)
	code := runFlow(context.Background(), []string{
		"process", "launch", "--code", "leave",
		"--variable", `{"day":2,"reason":"家事"}`, "--title", "张三的请假",
	}, io.Discard, io.Discard)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	if captured.path != "/api/v1/system/flow/process/launch" {
		t.Fatalf("path = %q", captured.path)
	}
	if captured.body["variable"] != `{"day":2,"reason":"家事"}` {
		t.Fatalf("variable = %v", captured.body["variable"])
	}
}

func TestRunFlowTaskAddSign_UserIdsCSV(t *testing.T) {
	captured := newFlowTestServer(t)
	code := runFlow(context.Background(), []string{
		"task", "add-sign", "--task-id", "9", "--user-ids", "u1,u2",
	}, io.Discard, io.Discard)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	ids, ok := captured.body["userIds"].([]any)
	if !ok || len(ids) != 2 {
		t.Fatalf("userIds = %v", captured.body["userIds"])
	}
}

func TestRunFlowDefGetList_PaginationAndFilters(t *testing.T) {
	captured := newFlowTestServer(t)
	code := runFlow(context.Background(), []string{
		"def", "get-list", "--page", "2", "--size", "50", "--state", "published", "--name", "请假",
	}, io.Discard, io.Discard)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	page, ok := captured.body["page"].(map[string]any)
	if !ok || page["page"] != float64(2) || page["size"] != float64(50) {
		t.Fatalf("page = %v", captured.body["page"])
	}
	if captured.body["state"] != "published" || captured.body["name"] != "请假" {
		t.Fatalf("filters = %+v", captured.body)
	}
}

func TestRunFlowUnknownSubcommand(t *testing.T) {
	code := runFlow(context.Background(), []string{"nope"}, io.Discard, io.Discard)
	if code != 2 {
		t.Fatalf("unknown subcommand exit = %d", code)
	}
}
