package shared

import (
	"fmt"
	"go/parser"
	"go/token"
	"io"
	"os"
	"strings"
)

// runScript 协议脚本相关命令
//   script validate <file>                校验协议脚本
//   script template [up-before|up-after|down-before|down-after]  生成脚本模板
func runScript(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "用法: script validate <file>                    # 校验协议脚本")
		fmt.Fprintln(stderr, "       script template [up-before|up-after|down-before|down-after]  # 生成脚本模板")
		return 2
	}

	subCmd := args[0]
	subArgs := args[1:]

	switch subCmd {
	case "validate":
		return runScriptValidate(subArgs, stdout, stderr)
	case "template":
		return runScriptTemplate(subArgs, stdout, stderr)
	default:
		fmt.Fprintf(stderr, "未知的 script 子命令: %s\n", subCmd)
		return 2
	}
}

// 允许导入的包名（yaegi 解释器白名单）
var allowedScriptImports = map[string]bool{
	"log":       true,
	"context":   true,
	"strings":   true,
	"json":      true,
	"gjson":     true,
	"utils":     true,
	"deviceMsg": true,
	"dm":        true,
	"schema":    true,
}

// Handle 函数签名模板（用于校验提示）
var handleSignatures = map[string]string{
	"up-before":    "func Handle(ctx context.Context, req *deviceMsg.PublishMsg) *deviceMsg.PublishMsg",
	"up-after":     "func Handle(ctx context.Context, req *deviceMsg.PublishMsg, resp *deviceMsg.PublishMsg)",
	"down-before":  "func Handle(ctx context.Context, req *deviceMsg.PublishMsg) *deviceMsg.PublishMsg",
	"down-after":   "func Handle(ctx context.Context, req *deviceMsg.PublishMsg)",
}

func runScriptValidate(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "用法: script validate <file>")
		return 2
	}

	filePath := args[0]
	var content []byte
	var err error

	if filePath == "-" {
		content, err = io.ReadAll(os.Stdin)
	} else {
		content, err = os.ReadFile(filePath)
	}
	if err != nil {
		fmt.Fprintf(stderr, "读取失败: %v\n", err)
		return 1
	}

	src := string(content)
	var errors []string
	var warnings []string

	// 1. 基础 Go 语法解析（yaegi 脚本无 package 声明，需补包装）
	fset := token.NewFileSet()
	_, err = parser.ParseFile(fset, "script.go", "package main\n"+src, parser.AllErrors)
	if err != nil {
		errors = append(errors, fmt.Sprintf("语法错误: %v", err))
	}

	// 2. 检查 Handle 函数是否存在
	if !strings.Contains(src, "func Handle(") {
		errors = append(errors, "未找到 Handle 函数定义，脚本入口必须为 func Handle(...)")
	}

	// 3. 检查 import 包名是否在白名单
	lines := strings.Split(src, "\n")
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "import ") {
			// 提取 import 的包名
			pkg := extractImportPkg(trimmed)
			if pkg != "" && !allowedScriptImports[pkg] {
				warnings = append(warnings, fmt.Sprintf("第%d行: import '%s' 不在 yaegi 白名单中，运行时可能不可用", i+1, pkg))
			}
		}
	}

	// 4. 提示可用的 Handle 签名
	if len(errors) == 0 {
		fmt.Fprintln(stdout, "✅ 基础校验通过")
		fmt.Fprintln(stdout)
		fmt.Fprintln(stdout, "Handle 函数签名参考:")
		for typ, sig := range handleSignatures {
			fmt.Fprintf(stdout, "  %-12s %s\n", typ+":", sig)
		}
	} else {
		fmt.Fprintf(stdout, "❌ 校验失败，共 %d 个错误:\n", len(errors))
		for _, e := range errors {
			fmt.Fprintf(stdout, "  %s\n", e)
		}
	}

	if len(warnings) > 0 {
		fmt.Fprintf(stdout, "\n⚠️  共 %d 个警告:\n", len(warnings))
		for _, w := range warnings {
			fmt.Fprintf(stdout, "  %s\n", w)
		}
	}

	fmt.Fprintln(stdout)
	fmt.Fprintln(stdout, "提示:")
	fmt.Fprintln(stdout, "  - Before 脚本返回 nil 会丢弃消息")
	fmt.Fprintln(stdout, "  - dm/schema 包函数需要真实服务运行，本地不可用")
	fmt.Fprintln(stdout, "  - 完整测试请使用: go test -v -run TestScript ./...")

	if len(errors) > 0 {
		return 1
	}
	return 0
}

func extractImportPkg(line string) string {
	// 匹配 import "pkg" 或 import 'pkg'
	line = strings.TrimPrefix(line, "import ")
	line = strings.TrimSpace(line)
	if (strings.HasPrefix(line, "\"") || strings.HasPrefix(line, "'")) &&
		(len(line) >= 3) {
		quote := line[0:1]
		end := strings.Index(line[1:], quote)
		if end >= 0 {
			return line[1 : 1+end]
		}
	}
	return ""
}

func runScriptTemplate(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "用法: script template [up-before|up-after|down-before|down-after]")
		return 2
	}

	scriptType := args[0]
	var tmpl string
	switch scriptType {
	case "up-before":
		tmpl = scriptUpBeforeTemplate()
	case "up-after":
		tmpl = scriptUpAfterTemplate()
	case "down-before":
		tmpl = scriptDownBeforeTemplate()
	case "down-after":
		tmpl = scriptDownAfterTemplate()
	default:
		fmt.Fprintf(stderr, "错误: 脚本类型应为 up-before/up-after/down-before/down-after，实际: %s\n", scriptType)
		return 2
	}

	fmt.Fprintln(stdout, tmpl)
	return 0
}

func scriptUpBeforeTemplate() string {
	return `// 上行前处理 (UpBefore) — 拦截设备上报消息，可修改或丢弃
// TriggerDir=1, TriggerTimer=1
// 返回 nil 表示丢弃消息

import "log"
import "context"
import "deviceMsg"
import "json"

func Handle(ctx context.Context, req *deviceMsg.PublishMsg) *deviceMsg.PublishMsg {
    log.Printf("收到消息: %s", string(req.Payload))

    var data map[string]any
    err := json.Unmarshal(req.Payload, &data)
    if err != nil {
        return req
    }

    // TODO: 修改 data...

    req.Payload, _ = json.Marshal(data)
    return req
}
`
}

func scriptUpAfterTemplate() string {
	return `// 上行后处理 (UpAfter) — 设备消息处理完成后执行，用于联动、记录等
// TriggerDir=1, TriggerTimer=2
// 无返回值

import "log"
import "context"
import "deviceMsg"

func Handle(ctx context.Context, req *deviceMsg.PublishMsg, resp *deviceMsg.PublishMsg) {
    log.Printf("请求: %s", string(req.Payload))
    if resp != nil {
        log.Printf("响应: %s", string(resp.Payload))
    }
    return
}
`
}

func scriptDownBeforeTemplate() string {
	return `// 下行前处理 (DownBefore) — 拦截平台下发指令，可修改或丢弃
// TriggerDir=2, TriggerTimer=1
// 返回 nil 表示丢弃消息

import "log"
import "context"
import "deviceMsg"
import "json"

func Handle(ctx context.Context, req *deviceMsg.PublishMsg) *deviceMsg.PublishMsg {
    log.Printf("下发指令: %s", string(req.Payload))

    // TODO: 修改指令...

    return req
}
`
}

func scriptDownAfterTemplate() string {
	return `// 下行后处理 (DownAfter) — 指令下发后执行，用于记录、联动等
// TriggerDir=2, TriggerTimer=2
// 无返回值

import "log"
import "context"
import "deviceMsg"

func Handle(ctx context.Context, req *deviceMsg.PublishMsg) {
    log.Printf("指令已下发: %s", string(req.Payload))
    return
}
`
}
