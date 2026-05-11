package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"gitee.com/unitedrhino/cli/internal/client"
	"gitee.com/unitedrhino/cli/internal/response"
)

func runAPI(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "usage: ur api <path> [--body JSON] [--body-file FILE] [--header KEY:VALUE] [--fields SELECTORS] [--summarize]")
		return 2
	}
	path := args[0]
	body := map[string]any{}
	headers := map[string]string{}
	fields := ""
	summarize := false

	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "--body":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--body requires JSON")
				return 2
			}
			parsed, err := parseBodyArg(args[i+1])
			if err != nil {
				fmt.Fprintln(stderr, err.Error())
				return 2
			}
			body = parsed
			i++
		case "--body-file":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--body-file requires file path")
				return 2
			}
			raw, err := os.ReadFile(args[i+1])
			if err != nil {
				fmt.Fprintf(stderr, "read body file: %v\n", err)
				return 1
			}
			parsed, err := parseBodyArg(string(raw))
			if err != nil {
				fmt.Fprintln(stderr, err.Error())
				return 2
			}
			body = parsed
			i++
		case "--header", "-H":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--header requires KEY:VALUE")
				return 2
			}
			parts := strings.SplitN(args[i+1], ":", 2)
			if len(parts) != 2 {
				fmt.Fprintf(stderr, "invalid header %q\n", args[i+1])
				return 2
			}
			headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
			i++
		case "--fields":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "--fields requires selectors")
				return 2
			}
			fields = args[i+1]
			i++
		case "--summarize":
			summarize = true
		default:
			fmt.Fprintf(stderr, "unknown api option: %s\n", args[i])
			return 2
		}
	}

	if fields != "" && summarize {
		fmt.Fprintln(stderr, "--fields and --summarize are mutually exclusive")
		return 2
	}

	resp, err := client.DoAPI(ctx, client.APIRequest{Path: path, Body: body, Headers: headers})
	if err != nil {
		fmt.Fprintln(stderr, err.Error())
		return 1
	}

	// 将 resp 转为 map[string]any 以便过滤
	var out any = resp
	if fields != "" {
		respMap, err := toMapAny(resp)
		if err != nil {
			fmt.Fprintf(stderr, "convert response: %v\n", err)
			return 1
		}
		selectors := strings.Split(fields, ",")
		filtered, err := response.FilterFields(respMap, selectors)
		if err != nil {
			fmt.Fprintf(stderr, "filter fields: %v\n", err)
			return 1
		}
		out = filtered
	} else if summarize {
		respMap, err := toMapAny(resp)
		if err != nil {
			fmt.Fprintf(stderr, "convert response: %v\n", err)
			return 1
		}
		out = response.Summarize(respMap)
	}

	raw, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		fmt.Fprintln(stderr, err.Error())
		return 1
	}
	_, _ = fmt.Fprintln(stdout, string(raw))
	return 0
}

func toMapAny(v any) (map[string]any, error) {
	raw, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	var m map[string]any
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, err
	}
	return m, nil
}
