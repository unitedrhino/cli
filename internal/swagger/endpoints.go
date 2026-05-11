package swagger

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Endpoint struct {
	Path        string `json:"path"`
	Method      string `json:"method"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	AuthType    string `json:"authType"`
	Group       string `json:"group"`
	Resource    string `json:"resource"`
	Action      string `json:"action"`
}

type doc struct {
	Paths map[string]map[string]operation `json:"paths"`
}

type operation struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
	AuthType    string `json:"x-auth-type"`
	Group       string `json:"x-group"`
	Resource    string `json:"x-resource"`
	Action      string `json:"x-action"`
}

func LoadEndpoints() ([]Endpoint, error) {
	files, err := ResolveFiles()
	if err != nil {
		return nil, err
	}
	var endpoints []Endpoint
	for _, file := range files {
		raw, err := os.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("read swagger %s: %w", file, err)
		}
		var d doc
		if err := json.Unmarshal(raw, &d); err != nil {
			return nil, fmt.Errorf("decode swagger %s: %w", file, err)
		}
		for path, methods := range d.Paths {
			for method, op := range methods {
				endpoints = append(endpoints, Endpoint{
					Path:        path,
					Method:      strings.ToUpper(method),
					Summary:     op.Summary,
					Description: op.Description,
					AuthType:    op.AuthType,
					Group:       op.Group,
					Resource:    op.Resource,
					Action:      op.Action,
				})
			}
		}
	}
	sort.Slice(endpoints, func(i, j int) bool {
		if endpoints[i].Path == endpoints[j].Path {
			return endpoints[i].Method < endpoints[j].Method
		}
		return endpoints[i].Path < endpoints[j].Path
	})
	return endpoints, nil
}

// FilterEndpointsByApp 按应用的 AllowedAuthTypes 过滤端点
func FilterEndpointsByApp(all []Endpoint, allowedAuthTypes []string) []Endpoint {
	allowed := make(map[string]bool, len(allowedAuthTypes))
	for _, at := range allowedAuthTypes {
		allowed[at] = true
	}
	out := make([]Endpoint, 0, len(all))
	for _, item := range all {
		if allowed[item.AuthType] {
			out = append(out, item)
		}
	}
	return out
}

func FilterEndpoints(all []Endpoint, targetPath, authType string) []Endpoint {
	targetPath = strings.TrimSpace(targetPath)
	authType = strings.TrimSpace(authType)
	out := make([]Endpoint, 0, len(all))
	for _, item := range all {
		if targetPath != "" && !strings.Contains(item.Path, targetPath) {
			continue
		}
		if authType != "" && item.AuthType != authType {
			continue
		}
		out = append(out, item)
	}
	return out
}
