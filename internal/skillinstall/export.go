// export.go — 将内置 ur-api 导出为遵循标准 Skills 目录结构的 ZIP 包。
package skillinstall

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// ExportResult 描述标准 Skills ZIP 的导出结果。
type ExportResult struct {
	// Format 是导出格式，当前固定为 zip。
	Format string `json:"format"`
	// Version 是技能包版本。
	Version string `json:"version"`
	// Path 是生成文件的绝对路径。
	Path string `json:"path"`
	// Files 是归档内的文件数量。
	Files int `json:"files"`
	// Bytes 是 ZIP 文件大小。
	Bytes int64 `json:"bytes"`
}

// DefaultExportPath 返回默认的标准 Skills ZIP 输出路径。
func DefaultExportPath(src string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("无法获取用户目录: %w", err)
	}
	version := readVersion(src)
	return filepath.Join(home, ".ur", "exports", "ur-api-skills-"+version+".zip"), nil
}

// ExportZIP 把 src 内容放入 ZIP 的 ur-api/ 根目录，供支持标准技能包的平台导入。
func ExportZIP(src, output string) (*ExportResult, error) {
	if info, err := os.Stat(src); err != nil || !info.IsDir() {
		if err != nil {
			return nil, fmt.Errorf("读取内置 Skills 失败: %w", err)
		}
		return nil, fmt.Errorf("内置 Skills 路径不是目录: %s", src)
	}
	if strings.TrimSpace(output) == "" {
		var err error
		output, err = DefaultExportPath(src)
		if err != nil {
			return nil, err
		}
	} else if !strings.EqualFold(filepath.Ext(output), ".zip") {
		output = filepath.Join(output, "ur-api-skills-"+readVersion(src)+".zip")
	}
	output = normalizePath(output)
	sourcePath := normalizePath(src)
	if relative, err := filepath.Rel(sourcePath, output); err == nil && relative != ".." && !strings.HasPrefix(relative, ".."+string(filepath.Separator)) {
		return nil, errorsOutputInsideSource(output)
	}
	if err := os.MkdirAll(filepath.Dir(output), 0o755); err != nil {
		return nil, fmt.Errorf("创建导出目录失败: %w", err)
	}
	temp, err := os.CreateTemp(filepath.Dir(output), ".ur-api-skills-*.zip")
	if err != nil {
		return nil, fmt.Errorf("创建临时 ZIP 失败: %w", err)
	}
	tempPath := temp.Name()
	defer os.Remove(tempPath)
	archive := zip.NewWriter(temp)
	files := 0
	walkErr := filepath.Walk(sourcePath, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if info.IsDir() {
			return nil
		}
		if !info.Mode().IsRegular() {
			return fmt.Errorf("不支持导出特殊文件: %s", path)
		}
		relative, err := filepath.Rel(sourcePath, path)
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = "ur-api/" + filepath.ToSlash(relative)
		header.Method = zip.Deflate
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		_, copyErr := io.Copy(writer, file)
		closeErr := file.Close()
		if copyErr != nil {
			return copyErr
		}
		if closeErr != nil {
			return closeErr
		}
		files++
		return nil
	})
	if walkErr != nil {
		archive.Close()
		temp.Close()
		return nil, fmt.Errorf("导出 Skills ZIP 失败: %w", walkErr)
	}
	if err := archive.Close(); err != nil {
		temp.Close()
		return nil, fmt.Errorf("完成 Skills ZIP 失败: %w", err)
	}
	if err := temp.Close(); err != nil {
		return nil, fmt.Errorf("关闭 Skills ZIP 失败: %w", err)
	}
	if err := replaceFileAtomically(tempPath, output); err != nil {
		return nil, fmt.Errorf("保存 Skills ZIP 失败: %w", err)
	}
	info, err := os.Stat(output)
	if err != nil {
		return nil, fmt.Errorf("读取 Skills ZIP 信息失败: %w", err)
	}
	return &ExportResult{Format: "zip", Version: readVersion(src), Path: output, Files: files, Bytes: info.Size()}, nil
}

// errorsOutputInsideSource 返回输出路径位于源目录内部时的明确错误。
func errorsOutputInsideSource(output string) error {
	return fmt.Errorf("导出路径不能位于内置 Skills 目录内部: %s", output)
}
