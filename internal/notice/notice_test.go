// notice_test.go 验证结构化提示合并和人类提示去重。
package notice

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

// TestMarshalJSON 验证提示仅合并到 JSON 对象且不会重复写入 stderr。
func TestMarshalJSON(t *testing.T) {
	Reset()
	SetUpdate(Update{Current: "v0.6.1", Latest: "v0.6.2", Message: "可升级", Command: "ur upgrade"})
	raw, err := MarshalJSON(map[string]any{"code": 200, "data": map[string]any{"ok": true}}, true)
	if err != nil {
		t.Fatalf("MarshalJSON: %v", err)
	}
	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		t.Fatalf("解析结果: %v", err)
	}
	if result["code"] != float64(200) || result["_notice"] == nil {
		t.Fatalf("业务字段或 _notice 缺失: %s", raw)
	}
	var stderr bytes.Buffer
	WriteHuman(&stderr)
	if stderr.Len() != 0 {
		t.Fatalf("结构化提示后不应重复输出 stderr: %q", stderr.String())
	}
}

// TestWriteHuman 验证 CLI 与 Skills 提示只输出一次。
func TestWriteHuman(t *testing.T) {
	Reset()
	SetUpdate(Update{Current: "v0.6.1", Latest: "v0.6.2", Command: "ur upgrade"})
	SetSkills(Skills{Current: "v0.6.0", Target: "v0.6.2", Targets: []string{"codex-user"}, Command: "ur upgrade"})
	var stderr bytes.Buffer
	WriteHuman(&stderr)
	WriteHuman(&stderr)
	output := stderr.String()
	if strings.Count(output, "发现 ur CLI 新版本") != 1 || strings.Count(output, "AI Skills") != 1 {
		t.Fatalf("提示内容或去重错误: %q", output)
	}
}

// TestMarshalJSONArrayKeepsShape 验证数组输出不会被包装，终端仍可兜底提醒。
func TestMarshalJSONArrayKeepsShape(t *testing.T) {
	Reset()
	SetUpdate(Update{Current: "v0.6.1", Latest: "v0.6.2", Command: "ur upgrade"})
	raw, err := MarshalJSON([]string{"a", "b"}, false)
	if err != nil {
		t.Fatalf("MarshalJSON: %v", err)
	}
	if string(raw) != `["a","b"]` {
		t.Fatalf("数组结构被改变: %s", raw)
	}
	var stderr bytes.Buffer
	WriteHuman(&stderr)
	if !strings.Contains(stderr.String(), "v0.6.2") {
		t.Fatalf("数组输出后应由 stderr 提示: %q", stderr.String())
	}
}
