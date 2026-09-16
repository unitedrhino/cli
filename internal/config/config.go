package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Profile struct {
	BaseURL      string `json:"baseURL"`
	AppID        string `json:"appID"`
	TenantCode   string `json:"tenantCode"`
	Role         string `json:"role,omitempty"`
	Token        string `json:"token,omitempty"`
	Account      string `json:"account,omitempty"`
	Password     string `json:"password,omitempty"`
	UserID       int64  `json:"userID,omitempty"`
	AccessKey    string `json:"accessKey,omitempty"`
	AccessSecret string `json:"accessSecret,omitempty"`
}

type Config struct {
	CurrentProfile string             `json:"currentProfile"`
	Profiles       map[string]Profile `json:"profiles"`
	Tokens         map[string]Token   `json:"tokens,omitempty"`
}

type Token struct {
	Token    string `json:"token"`
	ExpireAt int64  `json:"expireAt"`
}

type AuthContext struct {
	Token        string
	Account      string
	Password     string
	UserID       string
	AccessKey    string
	AccessSecret string
	Source       AuthSource
	Method       AuthMethod
}

// AuthSource 表示认证凭据来自运行时环境还是本地 profile。
type AuthSource string

const (
	// AuthSourceEnvironment 表示凭据来自 UR_* 环境变量。
	AuthSourceEnvironment AuthSource = "environment"
	// AuthSourceProfile 表示凭据来自本地 profile。
	AuthSourceProfile AuthSource = "profile"
)

// AuthMethod 表示 CLI 当前选中的认证方式。
type AuthMethod string

const (
	// AuthMethodToken 表示使用 Session Token。
	AuthMethodToken AuthMethod = "token"
	// AuthMethodPassword 表示使用账号和原始密码登录。
	AuthMethodPassword AuthMethod = "password"
	// AuthMethodAccessKey 表示使用 AK/SK 自签 JWT。
	AuthMethodAccessKey AuthMethod = "aksk"
)

func configDir() string {
	return filepath.Join(os.Getenv("HOME"), ".ur")
}

func ConfigPath() string {
	return filepath.Join(configDir(), "config.json")
}

func RuntimeProfileFromEnv() (Profile, bool) {
	baseURL := strings.TrimSpace(os.Getenv("UR_BASE_URL"))
	if baseURL == "" {
		return Profile{}, false
	}
	profile := Profile{
		BaseURL:      baseURL,
		AppID:        strings.TrimSpace(os.Getenv("UR_APP_ID")),
		TenantCode:   strings.TrimSpace(os.Getenv("UR_TENANT_CODE")),
		Role:         "admin",
		Token:        strings.TrimSpace(os.Getenv("UR_TOKEN")),
		Account:      strings.TrimSpace(os.Getenv("UR_ACCOUNT")),
		Password:     os.Getenv("UR_PASSWORD"),
		AccessKey:    strings.TrimSpace(os.Getenv("UR_ACCESS_KEY")),
		AccessSecret: os.Getenv("UR_ACCESS_SECRET"),
	}
	if rawUserID := strings.TrimSpace(os.Getenv("UR_USER_ID")); rawUserID != "" {
		if userID, err := strconv.ParseInt(rawUserID, 10, 64); err == nil {
			profile.UserID = userID
		}
	}
	return profile, true
}

func WriteConfig(cfg Config) error {
	if err := os.MkdirAll(configDir(), 0o755); err != nil {
		return fmt.Errorf("mkdir config dir: %w", err)
	}
	raw, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}
	if err := os.WriteFile(ConfigPath(), raw, 0o600); err != nil {
		return fmt.Errorf("write config: %w", err)
	}
	return nil
}

func ReadConfig() (Config, error) {
	raw, err := os.ReadFile(ConfigPath())
	if err != nil {
		return Config{}, err
	}
	var cfg Config
	if err := json.Unmarshal(raw, &cfg); err != nil {
		return Config{}, fmt.Errorf("decode config: %w", err)
	}
	if cfg.Profiles == nil {
		cfg.Profiles = map[string]Profile{}
	}
	return cfg, nil
}

func CurrentProfile() (Profile, error) {
	if runtimeProfile, ok := RuntimeProfileFromEnv(); ok {
		return runtimeProfile, nil
	}
	cfg, err := ReadConfig()
	if err != nil {
		return Profile{}, err
	}
	name := strings.TrimSpace(os.Getenv("UR_PROFILE"))
	if name == "" {
		name = cfg.CurrentProfile
	}
	if name == "" {
		name = "default"
	}
	profile, ok := cfg.Profiles[name]
	if !ok {
		return Profile{}, fmt.Errorf("profile %q not found", name)
	}
	return profile, nil
}

func GetBaseURL() (string, error) {
	if value := strings.TrimSpace(os.Getenv("UR_BASE_URL")); value != "" {
		return value, nil
	}
	profile, err := CurrentProfile()
	if err != nil {
		return "", err
	}
	if strings.TrimSpace(profile.BaseURL) == "" {
		return "", errors.New("baseURL is empty")
	}
	return strings.TrimSpace(profile.BaseURL), nil
}

func GetAppID() (string, error) {
	if value := strings.TrimSpace(os.Getenv("UR_APP_ID")); value != "" {
		return value, nil
	}
	if isRuntimeEnvironment() {
		return "", errors.New("UR_APP_ID is required in runtime environment")
	}
	profile, err := CurrentProfile()
	if err != nil {
		return "", err
	}
	return fallback(strings.TrimSpace(profile.AppID), "77"), nil
}

func GetTenantCode() (string, error) {
	if value := strings.TrimSpace(os.Getenv("UR_TENANT_CODE")); value != "" {
		return value, nil
	}
	if isRuntimeEnvironment() {
		return "", errors.New("UR_TENANT_CODE is required in runtime environment")
	}
	profile, err := CurrentProfile()
	if err != nil {
		return "", err
	}
	return fallback(strings.TrimSpace(profile.TenantCode), "default"), nil
}

func ResolveAuthContext() (AuthContext, error) {
	profile, err := CurrentProfile()
	if err != nil {
		return AuthContext{}, err
	}
	if isRuntimeEnvironment() {
		return resolveEnvironmentAuth(profile)
	}
	if hasCredentialEnvironment() {
		envProfile := Profile{
			Token:        strings.TrimSpace(os.Getenv("UR_TOKEN")),
			Account:      strings.TrimSpace(os.Getenv("UR_ACCOUNT")),
			Password:     os.Getenv("UR_PASSWORD"),
			AccessKey:    strings.TrimSpace(os.Getenv("UR_ACCESS_KEY")),
			AccessSecret: os.Getenv("UR_ACCESS_SECRET"),
		}
		if rawUserID := strings.TrimSpace(os.Getenv("UR_USER_ID")); rawUserID != "" {
			envProfile.UserID, _ = strconv.ParseInt(rawUserID, 10, 64)
		}
		return resolveEnvironmentAuth(envProfile)
	}
	return resolveProfileAuth(profile)
}

// resolveEnvironmentAuth 按 Token、AK/SK、账号密码顺序解析原子化的 Sandbox 凭据组。
func resolveEnvironmentAuth(profile Profile) (AuthContext, error) {
	ctx := AuthContext{Source: AuthSourceEnvironment}
	if profile.Token != "" {
		ctx.Token = profile.Token
		ctx.Method = AuthMethodToken
		return ctx, nil
	}
	if profile.AccessKey != "" || profile.AccessSecret != "" {
		if profile.AccessKey == "" || profile.AccessSecret == "" {
			return AuthContext{}, errors.New("UR_ACCESS_KEY and UR_ACCESS_SECRET must be set together")
		}
		ctx.AccessKey = profile.AccessKey
		ctx.AccessSecret = profile.AccessSecret
		if profile.UserID > 0 {
			ctx.UserID = strconv.FormatInt(profile.UserID, 10)
		}
		ctx.Method = AuthMethodAccessKey
		return ctx, nil
	}
	if profile.Account != "" || profile.Password != "" {
		if profile.Account == "" || profile.Password == "" {
			return AuthContext{}, errors.New("UR_ACCOUNT and UR_PASSWORD must be set together")
		}
		ctx.Account = profile.Account
		ctx.Password = profile.Password
		ctx.Method = AuthMethodPassword
		return ctx, nil
	}
	return AuthContext{}, errors.New("missing auth context in environment")
}

// resolveProfileAuth 保留历史 profile 中的候选凭据，并按 Token、账号密码、AK/SK 选择首选方式。
func resolveProfileAuth(profile Profile) (AuthContext, error) {
	ctx := AuthContext{
		Token:        profile.Token,
		Account:      profile.Account,
		Password:     profile.Password,
		AccessKey:    profile.AccessKey,
		AccessSecret: profile.AccessSecret,
		Source:       AuthSourceProfile,
	}
	if profile.UserID > 0 {
		ctx.UserID = strconv.FormatInt(profile.UserID, 10)
	}
	if profile.Token != "" {
		ctx.Method = AuthMethodToken
		return ctx, nil
	}
	if profile.Account != "" || profile.Password != "" {
		if profile.Account == "" || profile.Password == "" {
			return AuthContext{}, errors.New("profile account and password must be set together")
		}
		ctx.Method = AuthMethodPassword
		return ctx, nil
	}
	if profile.AccessKey != "" || profile.AccessSecret != "" {
		if profile.AccessKey == "" || profile.AccessSecret == "" {
			return AuthContext{}, errors.New("profile accessKey and accessSecret must be set together")
		}
		ctx.Method = AuthMethodAccessKey
		return ctx, nil
	}
	return AuthContext{}, errors.New("missing auth context")
}

func SaveProfile(profile Profile) error {
	cfg, err := ReadConfig()
	if err != nil {
		if os.IsNotExist(err) {
			cfg = Config{
				CurrentProfile: "default",
				Profiles:       map[string]Profile{},
			}
		} else {
			return err
		}
	}
	name := strings.TrimSpace(os.Getenv("UR_PROFILE"))
	if name == "" {
		name = cfg.CurrentProfile
	}
	if name == "" {
		name = "default"
		cfg.CurrentProfile = "default"
	}
	// merge: 逐字段合并，只覆盖非零值字段，保留已有 profile 中的其他字段
	if existing, ok := cfg.Profiles[name]; ok {
		if profile.BaseURL != "" {
			existing.BaseURL = profile.BaseURL
		}
		if profile.AppID != "" {
			existing.AppID = profile.AppID
		}
		if profile.TenantCode != "" {
			existing.TenantCode = profile.TenantCode
		}
		if profile.AccessKey != "" {
			existing.AccessKey = profile.AccessKey
		}
		if profile.AccessSecret != "" {
			existing.AccessSecret = profile.AccessSecret
		}
		if profile.UserID != 0 {
			existing.UserID = profile.UserID
		}
		if profile.Account != "" {
			existing.Account = profile.Account
		}
		if profile.Password != "" {
			existing.Password = profile.Password
		}
		if profile.Role != "" {
			existing.Role = profile.Role
		}
		if profile.Token != "" {
			existing.Token = profile.Token
		}
		cfg.Profiles[name] = existing
	} else {
		cfg.Profiles[name] = profile
	}
	return WriteConfig(cfg)
}

// ReplaceProfileAuth 显式切换当前 profile 的认证方式，并清除其他方式遗留的冲突字段。
func ReplaceProfileAuth(profile Profile, method AuthMethod) error {
	if isRuntimeEnvironment() {
		return nil
	}
	cfg, name, existing, err := loadCurrentStoredProfile()
	if err != nil {
		return err
	}
	if profile.BaseURL != "" {
		existing.BaseURL = profile.BaseURL
	}
	if profile.AppID != "" {
		existing.AppID = profile.AppID
	}
	if profile.TenantCode != "" {
		existing.TenantCode = profile.TenantCode
	}
	if profile.Role != "" {
		existing.Role = profile.Role
	}
	switch method {
	case AuthMethodPassword:
		existing.Token = profile.Token
		existing.Account = profile.Account
		existing.Password = profile.Password
		existing.UserID = 0
		existing.AccessKey = ""
		existing.AccessSecret = ""
	case AuthMethodAccessKey:
		existing.Token = ""
		existing.Account = ""
		existing.Password = ""
		existing.UserID = profile.UserID
		existing.AccessKey = profile.AccessKey
		existing.AccessSecret = profile.AccessSecret
	case AuthMethodToken:
		existing.Token = profile.Token
		existing.Account = ""
		existing.Password = ""
		existing.UserID = 0
		existing.AccessKey = ""
		existing.AccessSecret = ""
	default:
		return fmt.Errorf("unsupported auth method %q", method)
	}
	cfg.Profiles[name] = existing
	return WriteConfig(cfg)
}

// UpdateSessionToken 仅更新当前磁盘 profile 的 Session Token，保留历史账号密码用于刷新。
func UpdateSessionToken(token string) error {
	if isRuntimeEnvironment() {
		return nil
	}
	cfg, name, existing, err := loadCurrentStoredProfile()
	if err != nil {
		return err
	}
	existing.Token = token
	cfg.Profiles[name] = existing
	return WriteConfig(cfg)
}

// loadCurrentStoredProfile 读取当前磁盘 profile；配置不存在时创建兼容的默认容器。
func loadCurrentStoredProfile() (Config, string, Profile, error) {
	cfg, err := ReadConfig()
	if err != nil {
		if !os.IsNotExist(err) {
			return Config{}, "", Profile{}, err
		}
		cfg = Config{CurrentProfile: "default", Profiles: map[string]Profile{}}
	}
	name := strings.TrimSpace(os.Getenv("UR_PROFILE"))
	if name == "" {
		name = cfg.CurrentProfile
	}
	if name == "" {
		name = "default"
		cfg.CurrentProfile = name
	}
	return cfg, name, cfg.Profiles[name], nil
}

// isRuntimeEnvironment 判断当前是否由 Sandbox 等运行时通过 UR_BASE_URL 提供完整上下文。
func isRuntimeEnvironment() bool {
	return strings.TrimSpace(os.Getenv("UR_BASE_URL")) != ""
}

// hasCredentialEnvironment 判断是否显式提供了任意 UR_* 认证变量。
func hasCredentialEnvironment() bool {
	for _, key := range []string{"UR_TOKEN", "UR_ACCOUNT", "UR_PASSWORD", "UR_ACCESS_KEY", "UR_ACCESS_SECRET", "UR_USER_ID"} {
		if os.Getenv(key) != "" {
			return true
		}
	}
	return false
}

func fallback(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
