// 文件说明：验证面板导入元数据与自动归属参数，不访问真实平台或读取用户账号。
package things

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"gitee.com/unitedrhino/cli/internal/paneldev"
)

// TestPanelImportMetadata 验证 CLI 保留清单模式与显式覆盖模式，默认不再指定资源归属。
func TestPanelImportMetadata(t *testing.T) {
	for _, override := range []bool{false, true} {
		name := "保留清单"
		if override {
			name = "覆盖元数据"
		}
		t.Run(name, func(t *testing.T) {
			var body map[string]any
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if r.URL.Path == "/api/v1/system/common/upload-file" {
					io.WriteString(w, `{"code":200,"data":{"filePath":"test/owned/panel.zip"}}`)
					return
				}
				if r.URL.Path != "/api/v1/things/panel/package/import" {
					t.Errorf("意外请求 %s", r.URL.Path)
					w.WriteHeader(400)
					return
				}
				if e := json.NewDecoder(r.Body).Decode(&body); e != nil {
					t.Error(e)
				}
				io.WriteString(w, `{"code":200,"data":{"panelID":"new-panel"}}`)
			}))
			defer server.Close()
			t.Setenv("UR_BASE_URL", server.URL)
			t.Setenv("UR_APP_ID", "200")
			t.Setenv("UR_TENANT_CODE", "platform")
			t.Setenv("UR_TOKEN", "test-panel-token")
			directory := filepath.Join(t.TempDir(), "panel")
			if e := paneldev.Init(directory); e != nil {
				t.Fatal(e)
			}
			args := []string{"import", directory}
			if override {
				args = append(args, "--name", " 新面板 ", "--desc", " 新备注 ")
			}
			command := newPanelCommand()
			command.SetOut(io.Discard)
			command.SetErr(io.Discard)
			command.SetArgs(args)
			if e := command.Execute(); e != nil {
				t.Fatal(e)
			}
			if _, ok := body["ownerScope"]; ok {
				t.Fatal("默认归属应交由服务端按身份决定")
			}
			if override {
				if body["name"] != "新面板" || body["desc"] != "新备注" {
					t.Fatal("导入元数据丢失", body)
				}
			} else if _, ok := body["name"]; ok {
				t.Fatal("未指定元数据却覆盖了清单")
			}
		})
	}
}
