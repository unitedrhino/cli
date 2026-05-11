package ur

import (
	"context"
	"io"

	"gitee.com/unitedrhino/cli/cmd/shared"
	"gitee.com/unitedrhino/cli/internal/config"
)

// Execute 保持向后兼容，默认使用 org-manage 应用
func Execute(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	return shared.Execute(config.AppOrgManage, ctx, args, stdout, stderr)
}
