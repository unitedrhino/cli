// view_asset_test.go — 大屏资源库扩展名类型推断与上传响应 URL 提取单元测试
package shared

import "testing"

// TestExtractViewUploadedURL 表驱动校验 upload-file 响应 URL 提取（兼容各字段名与字符串形态）
func TestExtractViewUploadedURL(t *testing.T) {
	cases := []struct {
		name string
		data any
		want string
	}{
		{"字符串形态", "https://oss.example.com/a.png", "https://oss.example.com/a.png"},
		{"url 字段", map[string]any{"url": "https://oss.example.com/a.png"}, "https://oss.example.com/a.png"},
		{"fileUrl 字段", map[string]any{"fileUrl": "https://oss.example.com/a.png"}, "https://oss.example.com/a.png"},
		{"fileUri 字段", map[string]any{"fileUri": "oss/ub.png"}, "oss/ub.png"},
		{"filePath 字段", map[string]any{"filePath": "oss/ub.png"}, "oss/ub.png"},
		{"path 字段", map[string]any{"path": "oss/ub.png"}, "oss/ub.png"},
		{"空字段跳过", map[string]any{"url": "", "fileUri": "oss/ub.png"}, "oss/ub.png"},
		{"无有效字段", map[string]any{"code": 200}, ""},
		{"非对象非字符串", []any{"a"}, ""},
		{"nil", nil, ""},
	}
	for _, tc := range cases {
		if got := extractViewUploadedURL(tc.data); got != tc.want {
			t.Errorf("%s: extractViewUploadedURL(%v) = %q, want %q", tc.name, tc.data, got, tc.want)
		}
	}
}

// TestInferViewAssetType 表驱动校验扩展名 → 资源类型推断
func TestInferViewAssetType(t *testing.T) {
	cases := []struct {
		ext  string
		want string
	}{
		{".jpg", "image"},
		{".jpeg", "image"},
		{".png", "image"},
		{".gif", "image"},
		{".svg", "image"},
		{".webp", "image"},
		{".bmp", "image"},
		{".ico", "image"},
		{".PNG", "image"}, // 大小写不敏感
		{".mp4", "video"},
		{".webm", "video"},
		{".mov", "video"},
		{".mp3", "audio"},
		{".wav", "audio"},
		{".ogg", "audio"},
		{".json", "other"},
		{".txt", "other"},
		{"", "other"},
		{".unknown", "other"},
	}
	for _, tc := range cases {
		if got := inferViewAssetType(tc.ext); got != tc.want {
			t.Errorf("inferViewAssetType(%q) = %q, want %q", tc.ext, got, tc.want)
		}
	}
}
