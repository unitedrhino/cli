// github_alias_test.go 覆盖版本无关别名包相关行为：FindAsset 在别名包与
// 版本化资产并存时必须精确选中版本化资产，latestAliasDownload 仅对已上传
// 别名包的主流平台给出直链。
package upgrade

import (
	"net/url"
	"strings"
	"testing"
)

// TestFindAssetPrefersVersionedAsset 别名包(ur-cli-<平台>.tar.gz,无版本号)
// 与版本化资产并存时,Contains 匹配会同时命中,必须精确选中版本化资产——
// 别名仅供 latest 直链下载使用。
func TestFindAssetPrefersVersionedAsset(t *testing.T) {
	release := &Release{
		TagName: "v0.7.0",
		Assets: []Asset{
			{Name: "ur-cli-Linux-x86_64.tar.gz"},        // 别名包(内容相同)
			{Name: "ur-cli-v0.7.0-Linux-x86_64.tar.gz"}, // 版本化资产
			{Name: "ur-cli-v0.7.0-Windows-x86_64.zip"},  // 其他平台
		},
	}
	asset := release.FindAsset()
	if asset == nil {
		t.Fatal("FindAsset returned nil")
	}
	if asset.Name != "ur-cli-v0.7.0-Linux-x86_64.tar.gz" {
		t.Fatalf("selected %q, want versioned asset", asset.Name)
	}
}

// TestFindAssetFallsBackWhenOnlyAliasExists 仅有别名包时仍可回退选中
// (别名与版本化包内容一致,升级功能不受影响)。
func TestFindAssetFallsBackWhenOnlyAliasExists(t *testing.T) {
	release := &Release{
		TagName: "v0.7.0",
		Assets:  []Asset{{Name: "ur-cli-Linux-x86_64.tar.gz"}},
	}
	if asset := release.FindAsset(); asset == nil {
		t.Fatal("FindAsset returned nil for alias-only assets")
	}
}

// TestLatestAliasDownload 直链构造:协议/主机/路径合法,固定名含当前平台
// 友好名且不含版本号。
func TestLatestAliasDownload(t *testing.T) {
	got, name := latestAliasDownload()
	platform := platformReleaseName()
	if !aliasablePlatforms[platform] {
		if got != "" || name != "" {
			t.Fatalf("platform %q not aliasable, expect empty, got %q", platform, got)
		}
		return
	}
	if got == "" || name == "" {
		t.Fatalf("platform %q aliasable but got empty", platform)
	}
	if !strings.HasPrefix(got, GitHubLatestDownloadURL+"/") {
		t.Fatalf("url=%q", got)
	}
	if _, err := url.Parse(got); err != nil {
		t.Fatalf("parse url: %v", err)
	}
	if strings.Contains(name, "v0.") {
		t.Fatalf("alias name should be version-free: %q", name)
	}
	if !strings.HasSuffix(name, platform) && !strings.Contains(name, platform) {
		t.Fatalf("name %q missing platform %q", name, platform)
	}
}
