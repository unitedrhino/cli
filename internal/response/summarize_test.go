package response

import (
	"reflect"
	"testing"
)

func TestSummarize(t *testing.T) {
	t.Run("business error preserved", func(t *testing.T) {
		resp := map[string]any{"code": float64(400), "msg": "参数错误", "data": map[string]any{"detail": "xxx"}}
		out := Summarize(resp)
		if !reflect.DeepEqual(out, resp) {
			t.Errorf("got %+v, want %+v", out, resp)
		}
	})

	t.Run("null data", func(t *testing.T) {
		resp := map[string]any{"code": float64(200), "msg": "ok", "data": nil}
		out := Summarize(resp)
		want := map[string]any{"code": float64(200), "msg": "ok"}
		if !reflect.DeepEqual(out, want) {
			t.Errorf("got %+v, want %+v", out, want)
		}
	})

	t.Run("non-object data", func(t *testing.T) {
		resp := map[string]any{"code": float64(200), "msg": "ok", "data": "hello"}
		out := Summarize(resp)
		want := map[string]any{"code": float64(200), "msg": "ok", "data": "hello"}
		if !reflect.DeepEqual(out, want) {
			t.Errorf("got %+v, want %+v", out, want)
		}
	})

	t.Run("paged list", func(t *testing.T) {
		resp := map[string]any{
			"code": float64(200),
			"msg":  "ok",
			"data": map[string]any{
				"total": float64(100),
				"list": []any{
					map[string]any{"id": 1, "name": "a", "status": 1, "extra": "x"},
					map[string]any{"id": 2, "name": "b", "status": 2, "extra": "y"},
					map[string]any{"id": 3, "name": "c", "status": 3, "extra": "z"},
					map[string]any{"id": 4, "name": "d", "status": 4, "extra": "w"},
					map[string]any{"id": 5, "name": "e", "status": 5, "extra": "v"},
					map[string]any{"id": 6, "name": "f", "status": 6, "extra": "u"},
				},
			},
		}
		out := Summarize(resp)
		data, ok := out["data"].(map[string]any)
		if !ok {
			t.Fatalf("data is not map: %T", out["data"])
		}
		if data["total"] != int64(100) {
			t.Errorf("total = %v, want %v", data["total"], int64(100))
		}
		list, ok := data["list"].([]any)
		if !ok {
			t.Fatalf("list is not array: %T", data["list"])
		}
		if len(list) != 5 {
			t.Errorf("list len = %d, want 5", len(list))
		}
		note, ok := data["_note"].(string)
		if !ok || note == "" {
			t.Errorf("missing _note")
		}
		first, ok := list[0].(map[string]any)
		if !ok {
			t.Fatalf("first item is not map: %T", list[0])
		}
		if first["id"] != 1 || first["name"] != "a" {
			t.Errorf("first item = %+v", first)
		}
		if _, hasExtra := first["extra"]; hasExtra {
			t.Errorf("first item should not have extra field")
		}
	})

	t.Run("plain object", func(t *testing.T) {
		resp := map[string]any{
			"code": float64(200),
			"msg":  "ok",
			"data": map[string]any{
				"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6,
			},
		}
		out := Summarize(resp)
		data, ok := out["data"].(map[string]any)
		if !ok {
			t.Fatalf("data is not map: %T", out["data"])
		}
		if len(data) != 6 { // 5 fields + _note
			t.Errorf("data len = %d, want 6", len(data))
		}
		if _, hasNote := data["_note"]; !hasNote {
			t.Errorf("missing _note")
		}
	})

	t.Run("array field truncated", func(t *testing.T) {
		resp := map[string]any{
			"code": float64(200),
			"msg":  "ok",
			"data": map[string]any{
				"items": []any{1, 2, 3, 4, 5, 6, 7},
			},
		}
		out := Summarize(resp)
		data, ok := out["data"].(map[string]any)
		if !ok {
			t.Fatalf("data is not map: %T", out["data"])
		}
		arr, ok := data["items"].([]any)
		if !ok {
			t.Fatalf("items is not array: %T", data["items"])
		}
		if len(arr) != 5 {
			t.Errorf("items len = %d, want 5", len(arr))
		}
		if _, hasNote := data["items_note"]; !hasNote {
			t.Errorf("missing items_note")
		}
	})
}
