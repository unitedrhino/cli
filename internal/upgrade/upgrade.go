package upgrade

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"gitee.com/unitedrhino/cli/internal/version"
)

// Options 升级选项
type Options struct {
	TargetVersion string // 指定升级到的版本（为空则升级到最新）
	DryRun        bool   // 只检查不安装
	InstallSkills bool   // 升级成功后自动部署 skills 到各 AI 工具（ur skills install）
}

// Result 升级结果
type Result struct {
	CurrentVersion string `json:"currentVersion"`
	LatestVersion  string `json:"latestVersion"`
	UpToDate       bool   `json:"upToDate"`
	DownloadURL    string `json:"downloadUrl,omitempty"`
	ErrorMessage   string `json:"error,omitempty"`
	// SkillsSynced 升级后内置 skills 是否已同步（从安装包内提取）
	SkillsSynced bool `json:"skillsSynced,omitempty"`
	// SkillsUpdated 同步的 skill 目录数量
	SkillsUpdated int `json:"skillsUpdated,omitempty"`
	// SkillsMessage skills 同步的提示或警告信息
	SkillsMessage string `json:"skillsMessage,omitempty"`
}

// Check 检查是否有新版本可用
func Check() (*Result, error) {
	latestRelease, err := FetchLatestRelease()
	if err != nil {
		return &Result{
			CurrentVersion: version.BuildVersion,
			ErrorMessage:   err.Error(),
		}, nil
	}

	upToDate := !IsNewer(version.BuildVersion, latestRelease.TagName)

	result := &Result{
		CurrentVersion: version.BuildVersion,
		LatestVersion:  latestRelease.TagName,
		UpToDate:       upToDate,
	}
	if !upToDate {
		result.DownloadURL = latestRelease.HTMLURL
	}
	return result, nil
}

// Perform 执行升级
func Perform(opts Options) (*Result, error) {
	// 1. 获取目标 Release
	var release *Release
	var err error
	if opts.TargetVersion != "" {
		release, err = FetchRelease(opts.TargetVersion)
	} else {
		release, err = FetchLatestRelease()
	}
	if err != nil {
		return &Result{CurrentVersion: version.BuildVersion, ErrorMessage: err.Error()}, err
	}

	// 2. 检查是否需要升级
	if opts.TargetVersion == "" && !IsNewer(version.BuildVersion, release.TagName) {
		return &Result{
			CurrentVersion: version.BuildVersion,
			LatestVersion:  release.TagName,
			UpToDate:       true,
		}, nil
	}

	// 3. 查找匹配当前平台的资源
	asset := release.FindAsset()
	if asset == nil {
		err := fmt.Errorf("未找到适配当前平台 (%s) 的安装包", platformReleaseName())
		return &Result{
			CurrentVersion: version.BuildVersion,
			LatestVersion:  release.TagName,
			ErrorMessage:   err.Error(),
		}, err
	}

	result := &Result{
		CurrentVersion: version.BuildVersion,
		LatestVersion:  release.TagName,
		DownloadURL:    asset.BrowserDownloadURL,
	}

	if opts.DryRun {
		return result, nil
	}

	// 4. 获取当前二进制路径
	binaryPath, err := os.Executable()
	if err != nil {
		result.ErrorMessage = fmt.Sprintf("无法获取可执行文件路径: %v", err)
		return result, err
	}
	binaryPath, err = filepath.EvalSymlinks(binaryPath)
	if err != nil {
		result.ErrorMessage = fmt.Sprintf("无法解析可执行文件路径: %v", err)
		return result, err
	}

	// 5. 备份当前二进制
	if err := backupBinary(binaryPath, version.BuildVersion); err != nil {
		result.ErrorMessage = fmt.Sprintf("备份失败: %v", err)
		return result, err
	}

	// 6. 下载新版本
	tmpDir, err := os.MkdirTemp("", "ur-upgrade-*")
	if err != nil {
		result.ErrorMessage = fmt.Sprintf("创建临时目录失败: %v", err)
		return result, err
	}
	defer os.RemoveAll(tmpDir)

	archivePath := filepath.Join(tmpDir, asset.Name)
	if err := downloadFile(asset.BrowserDownloadURL, archivePath); err != nil {
		result.ErrorMessage = fmt.Sprintf("下载失败: %v", err)
		return result, err
	}

	// 7. 解压
	extractDir := filepath.Join(tmpDir, "extracted")
	if err := os.MkdirAll(extractDir, 0o755); err != nil {
		result.ErrorMessage = fmt.Sprintf("创建解压目录失败: %v", err)
		return result, err
	}

	if err := extract(archivePath, extractDir); err != nil {
		result.ErrorMessage = fmt.Sprintf("解压失败: %v", err)
		return result, err
	}

	// 8. 查找新二进制
	newBinary, err := findBinary(extractDir)
	if err != nil {
		result.ErrorMessage = fmt.Sprintf("查找新版本二进制失败: %v", err)
		return result, err
	}

	// 9. 替换当前二进制
	if err := replaceBinary(newBinary, binaryPath); err != nil {
		result.ErrorMessage = fmt.Sprintf("替换二进制失败: %v（旧版本已备份到 ~/.ur/backup/）", err)
		return result, err
	}

	// 10. 同步内置 skills：复用本次已解压的包内 skill/ 目录，避免二次下载。
	//     失败只记录警告，不阻断升级成功
	if skillsDir := GetDefaultSkillsDir(); skillsDir != "" {
		count, syncErr := SyncSkillsFromExtract(extractDir, skillsDir)
		if syncErr != nil {
			result.SkillsMessage = fmt.Sprintf("内置 skills 同步失败（可稍后运行 ur skills update）: %v", syncErr)
		} else {
			result.SkillsSynced = true
			result.SkillsUpdated = count
			result.SkillsMessage = "内置 skills 已随升级同步"
		}
	}

	return result, nil
}

// BackupDir 返回备份目录
func BackupDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return filepath.Join(os.TempDir(), ".ur", "backup")
	}
	return filepath.Join(home, ".ur", "backup")
}

// backupBinary 备份当前二进制文件
func backupBinary(binaryPath, oldVersion string) error {
	dir := BackupDir()
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("创建备份目录失败: %w", err)
	}

	backupName := filepath.Base(binaryPath) + "-" + oldVersion
	if runtime.GOOS == "windows" {
		backupName += ".exe"
	}
	backupPath := filepath.Join(dir, backupName)

	src, err := os.Open(binaryPath)
	if err != nil {
		return fmt.Errorf("读取当前二进制失败: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(backupPath)
	if err != nil {
		return fmt.Errorf("创建备份文件失败: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return fmt.Errorf("写入备份文件失败: %w", err)
	}

	// 保留执行权限
	if runtime.GOOS != "windows" {
		os.Chmod(backupPath, 0o755)
	}

	return nil
}

// downloadFile 下载文件到指定路径
func downloadFile(url, dest string) error {
	client := &http.Client{Timeout: 5 * time.Minute}
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("下载请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("下载返回状态码 %d", resp.StatusCode)
	}

	f, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("创建下载文件失败: %w", err)
	}
	defer f.Close()

	if _, err := io.Copy(f, resp.Body); err != nil {
		return fmt.Errorf("写入下载文件失败: %w", err)
	}

	return nil
}

// extract 解压 tar.gz 或 zip 文件
func extract(archivePath, destDir string) error {
	if strings.HasSuffix(archivePath, ".tar.gz") {
		return extractTarGz(archivePath, destDir)
	}
	if strings.HasSuffix(archivePath, ".zip") {
		return extractZip(archivePath, destDir)
	}
	return fmt.Errorf("不支持的压缩格式: %s", filepath.Base(archivePath))
}

func extractTarGz(archivePath, destDir string) error {
	f, err := os.Open(archivePath)
	if err != nil {
		return err
	}
	defer f.Close()

	gzReader, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gzReader.Close()

	tarReader := tar.NewReader(gzReader)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// 安全检查：防止路径穿越
		target := filepath.Join(destDir, header.Name)
		if !strings.HasPrefix(target, filepath.Clean(destDir)+string(os.PathSeparator)) {
			continue
		}

		switch header.Typeflag {
		case tar.TypeDir:
			os.MkdirAll(target, 0o755)
		case tar.TypeReg:
			os.MkdirAll(filepath.Dir(target), 0o755)
			outFile, err := os.Create(target)
			if err != nil {
				return err
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				outFile.Close()
				return err
			}
			outFile.Close()
			os.Chmod(target, os.FileMode(header.Mode))
		}
	}
	return nil
}

func extractZip(archivePath, destDir string) error {
	reader, err := zip.OpenReader(archivePath)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		target := filepath.Join(destDir, file.Name)
		if !strings.HasPrefix(target, filepath.Clean(destDir)+string(os.PathSeparator)) {
			continue
		}

		if file.FileInfo().IsDir() {
			os.MkdirAll(target, 0o755)
			continue
		}

		os.MkdirAll(filepath.Dir(target), 0o755)
		src, err := file.Open()
		if err != nil {
			return err
		}
		dst, err := os.Create(target)
		if err != nil {
			src.Close()
			return err
		}
		_, err = io.Copy(dst, src)
		src.Close()
		dst.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

// findBinary 在解压目录中查找 ur 二进制文件
func findBinary(dir string) (string, error) {
	// 在目录树中查找 ur 二进制
	var found string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		name := info.Name()
		if runtime.GOOS == "windows" {
			if name == "ur.exe" {
				found = path
				return filepath.SkipAll
			}
		} else {
			if name == "ur" {
				found = path
				return filepath.SkipAll
			}
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if found == "" {
		return "", fmt.Errorf("未在安装包中找到 ur 二进制文件")
	}
	return found, nil
}

// replaceBinary 原子替换二进制文件
func replaceBinary(newPath, oldPath string) error {
	// Unix: 直接 rename 即可（同一文件系统下是原子的）
	// Windows: 需要先删除旧文件再重命名（exe 文件在运行时无法删除）
	if runtime.GOOS == "windows" {
		oldBak := oldPath + ".old"
		os.Rename(oldPath, oldBak)
		if err := os.Rename(newPath, oldPath); err != nil {
			os.Rename(oldBak, oldPath) // 回滚
			return err
		}
		os.Remove(oldBak)
		return nil
	}

	// Unix: 确保新文件有执行权限
	if err := os.Chmod(newPath, 0o755); err != nil {
		return err
	}

	if err := os.Rename(newPath, oldPath); err != nil {
		return err
	}

	return nil
}
