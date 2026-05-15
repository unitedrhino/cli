package ur

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestExecuteTokenRawUsesEnvToken(t *testing.T) {
	t.Setenv("UR_BASE_URL", "http://127.0.0.1:7777")
	t.Setenv("UR_APP_ID", "100")
	t.Setenv("UR_TENANT_CODE", "platform")
	t.Setenv("UR_TOKEN", "runtime-token")

	var stdout, stderr bytes.Buffer
	code := Execute(context.Background(), "test", []string{"token", "--raw"}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("exit code = %d, stderr=%s", code, stderr.String())
	}
	if strings.TrimSpace(stdout.String()) != "runtime-token" {
		t.Fatalf("stdout = %q", stdout.String())
	}
}
