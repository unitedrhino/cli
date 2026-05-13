package response

import (
	"strings"
	"testing"
)

func TestFormatOutput_JSON(t *testing.T) {
	data := map[string]any{"code": 200, "msg": "ok"}
	out, err := FormatOutput(data, FormatOptions{Format: "json"})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), "\"code\": 200") {
		t.Fatalf("unexpected output: %s", out)
	}
}

func TestFormatOutput_Raw(t *testing.T) {
	data := map[string]any{"code": 200, "msg": "ok"}
	out, err := FormatOutput(data, FormatOptions{Format: "raw"})
	if err != nil {
		t.Fatal(err)
	}
	// raw 不应包含换行缩进
	if strings.Contains(string(out), "\n") {
		t.Fatalf("raw should not contain newlines: %s", out)
	}
}

func TestFormatOutput_YAML(t *testing.T) {
	data := map[string]any{"code": 200, "msg": "ok", "data": map[string]any{"name": "test"}}
	out, err := FormatOutput(data, FormatOptions{Format: "yaml"})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), "code: 200") {
		t.Fatalf("unexpected yaml: %s", out)
	}
}

func TestFormatOutput_Transform(t *testing.T) {
	data := map[string]any{
		"code": 200,
		"data": map[string]any{"total": 5, "list": []any{map[string]any{"id": 1}}},
	}
	out, err := FormatOutput(data, FormatOptions{Format: "json", Transform: "data.total"})
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "5" {
		t.Fatalf("expected 5, got: %s", out)
	}
}

func TestFormatOutput_TransformArray(t *testing.T) {
	data := map[string]any{
		"data": map[string]any{"list": []any{map[string]any{"name": "A"}, map[string]any{"name": "B"}}},
	}
	out, err := FormatOutput(data, FormatOptions{Format: "json", Transform: "data.list.#.name"})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), "A") || !strings.Contains(string(out), "B") {
		t.Fatalf("expected [A B], got: %s", out)
	}
}
