package shared

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"strings"

	"gitee.com/unitedrhino/cli/internal/auth"
)

func runToken(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	rawOutput := false
	decode := false
	for _, arg := range args {
		switch arg {
		case "--raw", "-r":
			rawOutput = true
		case "--decode", "-d":
			decode = true
		default:
			fmt.Fprintf(stderr, "unknown token option: %s\n", arg)
			return 2
		}
	}
	token, err := auth.ResolveToken(ctx)
	if err != nil {
		fmt.Fprintln(stderr, err.Error())
		return 1
	}
	if rawOutput {
		_, _ = fmt.Fprintln(stdout, token)
		return 0
	}
	if decode {
		parts := strings.Split(token, ".")
		if len(parts) != 3 {
			fmt.Fprintln(stderr, "token is not a JWT")
			return 1
		}
		payload, err := base64.RawURLEncoding.DecodeString(parts[1])
		if err != nil {
			fmt.Fprintln(stderr, err.Error())
			return 1
		}
		_, _ = fmt.Fprintln(stdout, string(payload))
		return 0
	}
	_, _ = fmt.Fprintln(stdout, token)
	return 0
}
