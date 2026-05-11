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
	UserID       string
	AccessKey    string
	AccessSecret string
}

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
		AppID:        fallback(strings.TrimSpace(os.Getenv("UR_APP_ID")), "77"),
		TenantCode:   fallback(strings.TrimSpace(os.Getenv("UR_TENANT_CODE")), "default"),
		Role:         "admin",
		Token:        strings.TrimSpace(os.Getenv("UR_TOKEN")),
		Account:      strings.TrimSpace(os.Getenv("UR_ACCOUNT")),
		Password:     strings.TrimSpace(os.Getenv("UR_PASSWORD")),
		AccessKey:    strings.TrimSpace(os.Getenv("UR_ACCESS_KEY")),
		AccessSecret: strings.TrimSpace(os.Getenv("UR_ACCESS_SECRET")),
	}
	if rawUserID := strings.TrimSpace(os.Getenv("UR_USER_ID")); rawUserID != "" {
		if userID, err := strconv.ParseInt(rawUserID, 10, 64); err == nil {
			profile.UserID = userID
		}
	}
	if profile.Token != "" || (profile.AccessKey != "" && profile.AccessSecret != "" && profile.UserID > 0) || (profile.Account != "" && profile.Password != "") {
		return profile, true
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
	if token := strings.TrimSpace(os.Getenv("UR_TOKEN")); token != "" {
		return AuthContext{Token: token}, nil
	}
	if profile.Token != "" {
		return AuthContext{Token: profile.Token}, nil
	}
	if userID := strings.TrimSpace(os.Getenv("UR_USER_ID")); userID != "" &&
		strings.TrimSpace(os.Getenv("UR_ACCESS_KEY")) != "" &&
		strings.TrimSpace(os.Getenv("UR_ACCESS_SECRET")) != "" {
		return AuthContext{
			UserID:       userID,
			AccessKey:    strings.TrimSpace(os.Getenv("UR_ACCESS_KEY")),
			AccessSecret: strings.TrimSpace(os.Getenv("UR_ACCESS_SECRET")),
		}, nil
	}
	if profile.UserID > 0 && profile.AccessKey != "" && profile.AccessSecret != "" {
		return AuthContext{
			UserID:       strconv.FormatInt(profile.UserID, 10),
			AccessKey:    profile.AccessKey,
			AccessSecret: profile.AccessSecret,
		}, nil
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
	cfg.Profiles[name] = profile
	return WriteConfig(cfg)
}

func fallback(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
