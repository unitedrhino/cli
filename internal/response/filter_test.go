package response

import (
	"reflect"
	"testing"
)

func TestFilterFields(t *testing.T) {
	src := map[string]any{
		"code": 200,
		"msg":  "ok",
		"data": map[string]any{
			"list": []any{
				map[string]any{"id": 1, "name": "a", "status": 1, "extra": "x"},
				map[string]any{"id": 2, "name": "b", "status": 2, "extra": "y"},
			},
			"total": 100,
			"info": map[string]any{
				"nickName":   "admin",
				"tenantCode": "default",
				"age":        30,
			},
		},
	}

	t.Run("single path", func(t *testing.T) {
		out, err := FilterFields(src, []string{"code"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := map[string]any{"code": 200}
		if !reflect.DeepEqual(out, want) {
			t.Errorf("got %+v, want %+v", out, want)
		}
	})

	t.Run("nested path", func(t *testing.T) {
		out, err := FilterFields(src, []string{"data.total"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := map[string]any{"data": map[string]any{"total": 100}}
		if !reflect.DeepEqual(out, want) {
			t.Errorf("got %+v, want %+v", out, want)
		}
	})

	t.Run("field selection", func(t *testing.T) {
		out, err := FilterFields(src, []string{"data.info.{nickName,tenantCode}"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := map[string]any{
			"data": map[string]any{
				"info": map[string]any{"nickName": "admin", "tenantCode": "default"},
			},
		}
		if !reflect.DeepEqual(out, want) {
			t.Errorf("got %+v, want %+v", out, want)
		}
	})

	t.Run("multiple selectors", func(t *testing.T) {
		out, err := FilterFields(src, []string{"code", "data.total"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := map[string]any{
			"code": 200,
			"data": map[string]any{"total": 100},
		}
		if !reflect.DeepEqual(out, want) {
			t.Errorf("got %+v, want %+v", out, want)
		}
	})

	t.Run("preserve array", func(t *testing.T) {
		out, err := FilterFields(src, []string{"data.list"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := map[string]any{
			"data": map[string]any{"list": src["data"].(map[string]any)["list"]},
		}
		if !reflect.DeepEqual(out, want) {
			t.Errorf("got %+v, want %+v", out, want)
		}
	})

	t.Run("empty selectors", func(t *testing.T) {
		out, err := FilterFields(src, []string{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(out, src) {
			t.Errorf("got %+v, want %+v", out, src)
		}
	})

	t.Run("missing key ignored", func(t *testing.T) {
		out, err := FilterFields(src, []string{"data.missing"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		// data 存在但 missing 不存在，保留空的 data map
		want := map[string]any{"data": map[string]any{}}
		if !reflect.DeepEqual(out, want) {
			t.Errorf("got %+v, want %+v", out, want)
		}
	})
}
